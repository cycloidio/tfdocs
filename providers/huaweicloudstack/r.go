package huaweicloudstack

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "huaweicloudstack_as_configuration_v1",
			Category:         "Auto Scaling Resources",
			ShortDescription: `Manages a V1 AS Configuration resource within HuaweiCloudStack.`,
			Description:      ``,
			Keywords: []string{
				"auto",
				"scaling",
				"as",
				"configuration",
				"v1",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to create the AS configuration. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new AS configuration.`,
				},
				resource.Attribute{
					Name:        "scaling_configuration_name",
					Description: `(Required) The name of the AS configuration. The name can contain letters, digits, underscores(_), and hyphens(-), and cannot exceed 64 characters.`,
				},
				resource.Attribute{
					Name:        "instance_config",
					Description: `(Required) The information about instance configurations. The instance_config dictionary data structure is documented below. The ` + "`" + `instance_config` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Optional) When using the existing instance specifications as the template to create AS configurations, specify this argument. In this case, flavor, image, and disk arguments do not take effect. If the instance_id argument is not specified, flavor, image, and disk arguments are mandatory.`,
				},
				resource.Attribute{
					Name:        "flavor",
					Description: `(Optional) The flavor ID.`,
				},
				resource.Attribute{
					Name:        "image",
					Description: `(Optional) The image ID.`,
				},
				resource.Attribute{
					Name:        "disk",
					Description: `(Optional) The disk group information. System disks are mandatory and data disks are optional. The dick structure is described below.`,
				},
				resource.Attribute{
					Name:        "key_name",
					Description: `(Required) The name of the SSH key pair used to log in to the instance.`,
				},
				resource.Attribute{
					Name:        "user_data",
					Description: `(Optional) The user data to provide when launching the instance. The file content must be encoded with Base64.`,
				},
				resource.Attribute{
					Name:        "personality",
					Description: `(Optional) Customize the personality of an instance by defining one or more files and their contents. The personality structure is described below.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `(Optional) The elastic IP address of the instance. The public_ip structure is described below.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `(Optional) Metadata key/value pairs to make available from within the instance. The ` + "`" + `disk` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `(Required) The disk size. The unit is GB. The system disk size ranges from 40 to 32768, and the data disk size ranges from 10 to 32768.`,
				},
				resource.Attribute{
					Name:        "volume_type",
					Description: `(Required) The disk type, which must be the same as the disk type available in the system. The available types are ` + "`" + `SSD` + "`" + `, ` + "`" + `SAS` + "`" + `, ` + "`" + `SATA` + "`" + ` or other types defined in CCS.`,
				},
				resource.Attribute{
					Name:        "disk_type",
					Description: `(Required) Whether the disk is a system disk or a data disk. Option ` + "`" + `DATA` + "`" + ` indicates a data disk. option ` + "`" + `SYS` + "`" + ` indicates a system disk. The ` + "`" + `personality` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Required) The absolute path of the destination file.`,
				},
				resource.Attribute{
					Name:        "contents",
					Description: `(Required) The content of the injected file, which must be encoded with base64. The ` + "`" + `public_ip` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "eip",
					Description: `(Required) The configuration parameter for creating an elastic IP address that will be automatically assigned to the instance. The eip structure is described below. The ` + "`" + `eip` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "ip_type",
					Description: `(Required) The IP address type. The system only supports ` + "`" + `5_bgp` + "`" + ` (indicates dynamic BGP).`,
				},
				resource.Attribute{
					Name:        "bandwidth",
					Description: `(Required) The bandwidth information. The structure is described below. The ` + "`" + `bandwidth` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `(Required) The bandwidth (Mbit/s). The value range is 1 to 300.`,
				},
				resource.Attribute{
					Name:        "share_type",
					Description: `(Required) The bandwidth sharing type. The system only supports ` + "`" + `PER` + "`" + ` (indicates exclusive bandwidth).`,
				},
				resource.Attribute{
					Name:        "charging_mode",
					Description: `(Required) The bandwidth charging mode. The system only supports ` + "`" + `traffic` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloudstack_as_group_v1",
			Category:         "Auto Scaling Resources",
			ShortDescription: `Manages a V1 Autoscaling Group resource within HuaweiCloudStack.`,
			Description:      ``,
			Keywords: []string{
				"auto",
				"scaling",
				"as",
				"group",
				"v1",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to create the AS group. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new AS group.`,
				},
				resource.Attribute{
					Name:        "scaling_group_name",
					Description: `(Required) The name of the scaling group. The name can contain letters, digits, underscores(_), and hyphens(-),and cannot exceed 64 characters.`,
				},
				resource.Attribute{
					Name:        "scaling_configuration_id",
					Description: `(Optional) The configuration ID which defines configurations of instances in the AS group.`,
				},
				resource.Attribute{
					Name:        "desire_instance_number",
					Description: `(Optional) The expected number of instances. The default value is the minimum number of instances. The value ranges from the minimum number of instances to the maximum number of instances.`,
				},
				resource.Attribute{
					Name:        "min_instance_number",
					Description: `(Optional) The minimum number of instances. The default value is 0.`,
				},
				resource.Attribute{
					Name:        "max_instance_number",
					Description: `(Optional) The maximum number of instances. The default value is 0.`,
				},
				resource.Attribute{
					Name:        "cool_down_time",
					Description: `(Optional) The cooling duration (in seconds). The value ranges from 0 to 86400, and is 900 by default.`,
				},
				resource.Attribute{
					Name:        "lb_listener_id",
					Description: `(Optional) The ELB (classic) listener IDs. The system supports up to three ELB listeners, the IDs of which are separated using a comma (,). This argument is deprecated, using ` + "`" + `lbaas_listeners` + "`" + ` instead.`,
				},
				resource.Attribute{
					Name:        "lbaas_listeners",
					Description: `(Optional) An array of one or more ELB (enhanced). The system supports the binding of up to three load balancers. The field is alternative to lb_listener_id. The lbaas_listeners object structure is documented below.`,
				},
				resource.Attribute{
					Name:        "available_zones",
					Description: `(Optional) The availability zones in which to create the instances in the autoscaling group.`,
				},
				resource.Attribute{
					Name:        "networks",
					Description: `(Required) An array of one or more network IDs. The system supports up to five networks. The networks object structure is documented below.`,
				},
				resource.Attribute{
					Name:        "security_groups",
					Description: `(Required) An array of one or more security group IDs to associate with the group. The security_groups object structure is documented below.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required) The VPC ID. Changing this creates a new group.`,
				},
				resource.Attribute{
					Name:        "health_periodic_audit_method",
					Description: `(Optional) The health check method for instances in the AS group. The health check methods include ` + "`" + `ELB_AUDIT` + "`" + ` and ` + "`" + `NOVA_AUDIT` + "`" + `. If load balancing is configured, the default value of this parameter is ` + "`" + `ELB_AUDIT` + "`" + `. Otherwise, the default value is ` + "`" + `NOVA_AUDIT` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "health_periodic_audit_time",
					Description: `(Optional) The health check period for instances. The period has four options: 5 minutes (default), 15 minutes, 60 minutes, and 180 minutes.`,
				},
				resource.Attribute{
					Name:        "instance_terminate_policy",
					Description: `(Optional) The instance removal policy. The policy has four options: ` + "`" + `OLD_CONFIG_OLD_INSTANCE` + "`" + ` (default), ` + "`" + `OLD_CONFIG_NEW_INSTANCE` + "`" + `, ` + "`" + `OLD_INSTANCE` + "`" + `, and ` + "`" + `NEW_INSTANCE` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "notifications",
					Description: `(Optional) The notification mode. The system only supports ` + "`" + `EMAIL` + "`" + ` mode which refers to notification by email.`,
				},
				resource.Attribute{
					Name:        "delete_publicip",
					Description: `(Optional) Whether to delete the elastic IP address bound to the instances of AS group when deleting the instances. The options are ` + "`" + `true` + "`" + ` and ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "delete_instances",
					Description: `(Optional) Whether to delete the instances in the AS group when deleting the AS group. The options are ` + "`" + `yes` + "`" + ` and ` + "`" + `no` + "`" + `. The ` + "`" + `networks` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Required) The network UUID. The ` + "`" + `security_groups` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Required) The UUID of the security group. The ` + "`" + `lbaas_listeners` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "listener_id",
					Description: `(Required) Specifies the ELB listener ID.`,
				},
				resource.Attribute{
					Name:        "protocol_port",
					Description: `(Required) Specifies the backend protocol, which is the port on which a backend ECS listens for traffic. The number of the port ranges from 1 to 65535.`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `(Optional) Specifies the weight, which determines the portion of requests a backend ECS processes compared to other backend ECSs added to the same listener. The value of this parameter ranges from 0 to 100. The default value is 1. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "scaling_group_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "scaling_group_status",
					Description: `The status of the AS group.`,
				},
				resource.Attribute{
					Name:        "current_instance_number",
					Description: `The number of current instances in the AS group.`,
				},
				resource.Attribute{
					Name:        "desire_instance_number",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "min_instance_number",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "max_instance_number",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "cool_down_time",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "lb_listener_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "lbaas_listeners",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "health_periodic_audit_method",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "health_periodic_audit_time",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "instance_terminate_policy",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "scaling_configuration_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "delete_publicip",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "notifications",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "instances",
					Description: `The instances IDs of the AS group.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "scaling_group_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "scaling_group_status",
					Description: `The status of the AS group.`,
				},
				resource.Attribute{
					Name:        "current_instance_number",
					Description: `The number of current instances in the AS group.`,
				},
				resource.Attribute{
					Name:        "desire_instance_number",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "min_instance_number",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "max_instance_number",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "cool_down_time",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "lb_listener_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "lbaas_listeners",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "health_periodic_audit_method",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "health_periodic_audit_time",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "instance_terminate_policy",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "scaling_configuration_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "delete_publicip",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "notifications",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "instances",
					Description: `The instances IDs of the AS group.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloudstack_as_policy_v1",
			Category:         "Auto Scaling Resources",
			ShortDescription: `Manages a V1 AS Policy resource within HuaweiCloudStack.`,
			Description:      ``,
			Keywords: []string{
				"auto",
				"scaling",
				"as",
				"policy",
				"v1",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to create the AS policy. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new AS policy.`,
				},
				resource.Attribute{
					Name:        "scaling_policy_name",
					Description: `(Required) The name of the AS policy. The name can contain letters, digits, underscores(_), and hyphens(-),and cannot exceed 64 characters.`,
				},
				resource.Attribute{
					Name:        "scaling_group_id",
					Description: `(Required) The AS group ID. Changing this creates a new AS policy.`,
				},
				resource.Attribute{
					Name:        "scaling_policy_type",
					Description: `(Required) The AS policy type. The values can be ` + "`" + `ALARM` + "`" + `, ` + "`" + `SCHEDULED` + "`" + `, and ` + "`" + `RECURRENCE` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "alarm_id",
					Description: `(Optional) The alarm rule ID. This argument is mandatory when ` + "`" + `scaling_policy_type` + "`" + ` is set to ` + "`" + `ALARM` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "scheduled_policy",
					Description: `(Optional) The periodic or scheduled AS policy. This argument is mandatory when ` + "`" + `scaling_policy_type` + "`" + ` is set to ` + "`" + `SCHEDULED` + "`" + ` or ` + "`" + `RECURRENCE` + "`" + `. The scheduled_policy structure is documented below.`,
				},
				resource.Attribute{
					Name:        "scaling_policy_action",
					Description: `(Optional) The action of the AS policy. The scaling_policy_action structure is documented below.`,
				},
				resource.Attribute{
					Name:        "cool_down_time",
					Description: `(Optional) The cooling duration (in seconds), and is 900 by default. The ` + "`" + `scheduled_policy` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "launch_time",
					Description: `(Required) The time when the scaling action is triggered. If ` + "`" + `scaling_policy_type` + "`" + ` is set to ` + "`" + `SCHEDULED` + "`" + `, the time format is YYYY-MM-DDThh:mmZ. If ` + "`" + `scaling_policy_type` + "`" + ` is set to ` + "`" + `RECURRENCE` + "`" + `, the time format is hh:mm.`,
				},
				resource.Attribute{
					Name:        "recurrence_type",
					Description: `(Optional) The periodic triggering type. This argument is mandatory when ` + "`" + `scaling_policy_type` + "`" + ` is set to ` + "`" + `RECURRENCE` + "`" + `. The options include ` + "`" + `Daily` + "`" + `, ` + "`" + `Weekly` + "`" + `, and ` + "`" + `Monthly` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "recurrence_value",
					Description: `(Optional) The frequency at which scaling actions are triggered.`,
				},
				resource.Attribute{
					Name:        "start_time",
					Description: `(Optional) The start time of the scaling action triggered periodically. The time format complies with UTC. The current time is used by default. The time format is YYYY-MM-DDThh:mmZ.`,
				},
				resource.Attribute{
					Name:        "end_time",
					Description: `(Optional) The end time of the scaling action triggered periodically. The time format complies with UTC. This argument is mandatory when ` + "`" + `scaling_policy_type` + "`" + ` is set to ` + "`" + `RECURRENCE` + "`" + `. The time format is YYYY-MM-DDThh:mmZ. The ` + "`" + `scaling_policy_action` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "operation",
					Description: `(Optional) The operation to be performed. The options include ` + "`" + `ADD` + "`" + ` (default), ` + "`" + `REMOVE` + "`" + `, and ` + "`" + `SET` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "instance_number",
					Description: `(Optional) The number of instances to be operated. The default number is 1. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "scaling_policy_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "scaling_policy_type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "alarm_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "cool_down_time",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "scaling_policy_action/operation",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "scaling_policy_action/instance_number",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "scheduled_policy/launch_time",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "scheduled_policy/recurrence_type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "scheduled_policy/recurrence_value",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "scheduled_policy/start_time",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "scheduled_policy/end_time",
					Description: `See Argument Reference above.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "scaling_policy_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "scaling_policy_type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "alarm_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "cool_down_time",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "scaling_policy_action/operation",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "scaling_policy_action/instance_number",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "scheduled_policy/launch_time",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "scheduled_policy/recurrence_type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "scheduled_policy/recurrence_value",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "scheduled_policy/start_time",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "scheduled_policy/end_time",
					Description: `See Argument Reference above.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloudstack_blockstorage_volume_v2",
			Category:         "Block Storage Resources",
			ShortDescription: `Manages a V2 volume resource within HuaweiCloudStack.`,
			Description:      ``,
			Keywords: []string{
				"block",
				"storage",
				"blockstorage",
				"volume",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to create the volume. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new volume.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `(Required) The size of the volume to create (in gigabytes). Changing this creates a new volume.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `(Optional) The availability zone for the volume. Changing this creates a new volume.`,
				},
				resource.Attribute{
					Name:        "consistency_group_id",
					Description: `(Optional) The consistency group to place the volume in.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A description of the volume. Changing this updates the volume's description.`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `(Optional) The image ID from which to create the volume. Changing this creates a new volume.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `(Optional) Metadata key/value pairs to associate with the volume. Changing this updates the existing volume metadata.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) A unique name for the volume. Changing this updates the volume's name.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `(Optional) The snapshot ID from which to create the volume. Changing this creates a new volume.`,
				},
				resource.Attribute{
					Name:        "source_replica",
					Description: `(Optional) The volume ID to replicate with.`,
				},
				resource.Attribute{
					Name:        "source_vol_id",
					Description: `(Optional) The volume ID from which to create the volume. Changing this creates a new volume.`,
				},
				resource.Attribute{
					Name:        "volume_type",
					Description: `(Optional) The type of volume to create. Available types are ` + "`" + `SSD` + "`" + `, ` + "`" + `SAS` + "`" + ` and ` + "`" + `SATA` + "`" + `. Changing this creates a new volume.`,
				},
				resource.Attribute{
					Name:        "cascade",
					Description: `(Optional, Default:false) Specifies to delete all snapshots associated with the EVS disk. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "source_vol_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "volume_type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "attachment",
					Description: `If a volume is attached to an instance, this attribute will display the Attachment ID, Instance ID, and the Device as the Instance sees it. ## Import Volumes can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import huaweicloudstack_blockstorage_volume_v2.volume_1 ea257959-eeb1-4c10-8d33-26f0409a755d ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "source_vol_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "volume_type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "attachment",
					Description: `If a volume is attached to an instance, this attribute will display the Attachment ID, Instance ID, and the Device as the Instance sees it. ## Import Volumes can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import huaweicloudstack_blockstorage_volume_v2.volume_1 ea257959-eeb1-4c10-8d33-26f0409a755d ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloudstack_compute_floatingip_associate_v2",
			Category:         "Compute Resources",
			ShortDescription: `Associate a floating IP to an instance`,
			Description:      ``,
			Keywords: []string{
				"compute",
				"floatingip",
				"associate",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 Compute client. Keypairs are associated with accounts, but a Compute client is needed to create one. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new floatingip_associate.`,
				},
				resource.Attribute{
					Name:        "floating_ip",
					Description: `(Required) The floating IP to associate.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Required) The instance to associte the floating IP with.`,
				},
				resource.Attribute{
					Name:        "fixed_ip",
					Description: `(Optional) The specific IP address to direct traffic to. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "floating_ip",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "fixed_ip",
					Description: `See Argument Reference above. ## Import This resource can be imported by specifying all three arguments, separated by a forward slash: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import huaweicloudstack_compute_floatingip_associate_v2.fip_1 <floating_ip>/<instance_id>/<fixed_ip> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "floating_ip",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "fixed_ip",
					Description: `See Argument Reference above. ## Import This resource can be imported by specifying all three arguments, separated by a forward slash: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import huaweicloudstack_compute_floatingip_associate_v2.fip_1 <floating_ip>/<instance_id>/<fixed_ip> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloudstack_compute_instance_v2",
			Category:         "Compute Resources",
			ShortDescription: `Manages a V2 VM instance resource within HuaweiCloudStack.`,
			Description:      ``,
			Keywords: []string{
				"compute",
				"instance",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to create the server instance. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new server.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A unique name for the resource.`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `(Optional; Required if ` + "`" + `image_name` + "`" + ` is empty and not booting from a volume. Do not specify if booting from a volume.) The image ID of the desired image for the server. Changing this creates a new server.`,
				},
				resource.Attribute{
					Name:        "image_name",
					Description: `(Optional; Required if ` + "`" + `image_id` + "`" + ` is empty and not booting from a volume. Do not specify if booting from a volume.) The name of the desired image for the server. Changing this creates a new server.`,
				},
				resource.Attribute{
					Name:        "flavor_id",
					Description: `(Optional; Required if ` + "`" + `flavor_name` + "`" + ` is empty) The flavor ID of the desired flavor for the server. Changing this resizes the existing server.`,
				},
				resource.Attribute{
					Name:        "flavor_name",
					Description: `(Optional; Required if ` + "`" + `flavor_id` + "`" + ` is empty) The name of the desired flavor for the server. Changing this resizes the existing server.`,
				},
				resource.Attribute{
					Name:        "user_data",
					Description: `(Optional) The user data to provide when launching the instance. Changing this creates a new server.`,
				},
				resource.Attribute{
					Name:        "security_groups",
					Description: `(Optional) An array of one or more security group names to associate with the server. Changing this results in adding/removing security groups from the existing server.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `(Optional) The availability zone in which to create the server. Changing this creates a new server.`,
				},
				resource.Attribute{
					Name:        "network",
					Description: `(Optional) An array of one or more networks to attach to the instance. The network object structure is documented below. Changing this creates a new server.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `(Optional) Metadata key/value pairs to make available from within the instance. Changing this updates the existing server metadata.`,
				},
				resource.Attribute{
					Name:        "config_drive",
					Description: `(Optional) Whether to use the config_drive feature to configure the instance. Changing this creates a new server.`,
				},
				resource.Attribute{
					Name:        "admin_pass",
					Description: `(Optional) The administrative password to assign to the server. Changing this changes the root password on the existing server.`,
				},
				resource.Attribute{
					Name:        "key_pair",
					Description: `(Optional) The name of a key pair to put on the server. The key pair must already be created and associated with the tenant's account. Changing this creates a new server.`,
				},
				resource.Attribute{
					Name:        "block_device",
					Description: `(Optional) Configuration of block devices. The block_device structure is documented below. Changing this creates a new server. You can specify multiple block devices which will create an instance with multiple disks. This configuration is very flexible, so please see the following [reference](http://docs.openstack.org/developer/nova/block_device_mapping.html) for more information.`,
				},
				resource.Attribute{
					Name:        "scheduler_hints",
					Description: `(Optional) Provide the Nova scheduler with hints on how the instance should be launched. The available hints are described below.`,
				},
				resource.Attribute{
					Name:        "stop_before_destroy",
					Description: `(Optional) Whether to try stop instance gracefully before destroying it, thus giving chance for guest OS daemons to stop correctly. If instance doesn't stop within timeout, it will be destroyed anyway. The ` + "`" + `network` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `(Required unless ` + "`" + `port` + "`" + ` or ` + "`" + `name` + "`" + ` is provided) The network UUID to attach to the server. Changing this creates a new server.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required unless ` + "`" + `uuid` + "`" + ` or ` + "`" + `port` + "`" + ` is provided) The human-readable name of the network. Changing this creates a new server.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Required unless ` + "`" + `uuid` + "`" + ` or ` + "`" + `name` + "`" + ` is provided) The port UUID of a network to attach to the server. Changing this creates a new server.`,
				},
				resource.Attribute{
					Name:        "fixed_ip_v4",
					Description: `(Optional) Specifies a fixed IPv4 address to be used on this network. Changing this creates a new server.`,
				},
				resource.Attribute{
					Name:        "fixed_ip_v6",
					Description: `(Optional) Specifies a fixed IPv6 address to be used on this network. Changing this creates a new server.`,
				},
				resource.Attribute{
					Name:        "access_network",
					Description: `(Optional) Specifies if this network should be used for provisioning access. Accepts true or false. Defaults to false. The ` + "`" + `block_device` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `(Required unless ` + "`" + `source_type` + "`" + ` is set to ` + "`" + `"blank"` + "`" + ` ) The UUID of the image, volume, or snapshot. Changing this creates a new server.`,
				},
				resource.Attribute{
					Name:        "source_type",
					Description: `(Required) The source type of the device. Must be one of "blank", "image", "volume", or "snapshot". Changing this creates a new server.`,
				},
				resource.Attribute{
					Name:        "volume_size",
					Description: `The size of the volume to create (in gigabytes). Required in the following combinations: source=image and destination=volume, source=blank and destination=local, and source=blank and destination=volume. Changing this creates a new server.`,
				},
				resource.Attribute{
					Name:        "boot_index",
					Description: `(Optional) The boot index of the volume. It defaults to 0. Changing this creates a new server.`,
				},
				resource.Attribute{
					Name:        "destination_type",
					Description: `(Optional) The type that gets created. Possible values are "volume" and "local". Changing this creates a new server.`,
				},
				resource.Attribute{
					Name:        "delete_on_termination",
					Description: `(Optional) Delete the volume / block device upon termination of the instance. Defaults to false. Changing this creates a new server. The ` + "`" + `scheduler_hints` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "group",
					Description: `(Optional) A UUID of a Server Group. The instance will be placed into that group.`,
				},
				resource.Attribute{
					Name:        "different_host",
					Description: `(Optional) A list of instance UUIDs. The instance will be scheduled on a different host than all other instances.`,
				},
				resource.Attribute{
					Name:        "same_host",
					Description: `(Optional) A list of instance UUIDs. The instance will be scheduled on the same host of those specified.`,
				},
				resource.Attribute{
					Name:        "query",
					Description: `(Optional) A conditional query that a compute node must pass in order to host an instance.`,
				},
				resource.Attribute{
					Name:        "target_cell",
					Description: `(Optional) The name of a cell to host the instance.`,
				},
				resource.Attribute{
					Name:        "build_near_host_ip",
					Description: `(Optional) An IP Address in CIDR form. The instance will be placed on a compute node that is in the same subnet. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "access_ip_v4",
					Description: `The first detected Fixed IPv4 address _or_ the Floating IP.`,
				},
				resource.Attribute{
					Name:        "access_ip_v6",
					Description: `The first detected Fixed IPv6 address.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "security_groups",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "flavor_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "flavor_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "network/uuid",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "network/name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "network/port",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "network/fixed_ip_v4",
					Description: `The Fixed IPv4 address of the Instance on that network.`,
				},
				resource.Attribute{
					Name:        "network/fixed_ip_v6",
					Description: `The Fixed IPv6 address of the Instance on that network.`,
				},
				resource.Attribute{
					Name:        "network/mac",
					Description: `The MAC address of the NIC on that network.`,
				},
				resource.Attribute{
					Name:        "all_metadata",
					Description: `Contains all instance metadata, even metadata not set by Terraform. ## Notes ### Multiple Ephemeral Disks It's possible to specify multiple ` + "`" + `block_device` + "`" + ` entries to create an instance with multiple ephemeral (local) disks. In order to create multiple ephemeral disks, the sum of the total amount of ephemeral space must be less than or equal to what the chosen flavor supports. The following example shows how to create an instance with multiple ephemeral disks: ` + "`" + `` + "`" + `` + "`" + ` resource "huaweicloudstack_compute_instance_v2" "foo" { name = "terraform-test" security_groups = ["default"] block_device { boot_index = 0 delete_on_termination = true destination_type = "local" source_type = "image" uuid = "<image uuid>" } block_device { boot_index = -1 delete_on_termination = true destination_type = "local" source_type = "blank" volume_size = 1 } block_device { boot_index = -1 delete_on_termination = true destination_type = "local" source_type = "blank" volume_size = 1 } } ` + "`" + `` + "`" + `` + "`" + ` ### Instances and Ports Neutron Ports are a great feature and provide a lot of functionality. However, there are some notes to be aware of when mixing Instances and Ports:`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "access_ip_v4",
					Description: `The first detected Fixed IPv4 address _or_ the Floating IP.`,
				},
				resource.Attribute{
					Name:        "access_ip_v6",
					Description: `The first detected Fixed IPv6 address.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "security_groups",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "flavor_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "flavor_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "network/uuid",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "network/name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "network/port",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "network/fixed_ip_v4",
					Description: `The Fixed IPv4 address of the Instance on that network.`,
				},
				resource.Attribute{
					Name:        "network/fixed_ip_v6",
					Description: `The Fixed IPv6 address of the Instance on that network.`,
				},
				resource.Attribute{
					Name:        "network/mac",
					Description: `The MAC address of the NIC on that network.`,
				},
				resource.Attribute{
					Name:        "all_metadata",
					Description: `Contains all instance metadata, even metadata not set by Terraform. ## Notes ### Multiple Ephemeral Disks It's possible to specify multiple ` + "`" + `block_device` + "`" + ` entries to create an instance with multiple ephemeral (local) disks. In order to create multiple ephemeral disks, the sum of the total amount of ephemeral space must be less than or equal to what the chosen flavor supports. The following example shows how to create an instance with multiple ephemeral disks: ` + "`" + `` + "`" + `` + "`" + ` resource "huaweicloudstack_compute_instance_v2" "foo" { name = "terraform-test" security_groups = ["default"] block_device { boot_index = 0 delete_on_termination = true destination_type = "local" source_type = "image" uuid = "<image uuid>" } block_device { boot_index = -1 delete_on_termination = true destination_type = "local" source_type = "blank" volume_size = 1 } block_device { boot_index = -1 delete_on_termination = true destination_type = "local" source_type = "blank" volume_size = 1 } } ` + "`" + `` + "`" + `` + "`" + ` ### Instances and Ports Neutron Ports are a great feature and provide a lot of functionality. However, there are some notes to be aware of when mixing Instances and Ports:`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloudstack_compute_interface_attach_v2",
			Category:         "Compute Resources",
			ShortDescription: `Attaches a Network Interface to an Instance.`,
			Description:      ``,
			Keywords: []string{
				"compute",
				"interface",
				"attach",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to create the interface attachment. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new attachment.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Required) The ID of the Instance to attach the Port or Network to.`,
				},
				resource.Attribute{
					Name:        "port_id",
					Description: `(Optional) The ID of the Port to attach to an Instance. _NOTE_: This option and ` + "`" + `network_id` + "`" + ` are mutually exclusive.`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `(Optional) The ID of the Network to attach to an Instance. A port will be created automatically. _NOTE_: This option and ` + "`" + `port_id` + "`" + ` are mutually exclusive.`,
				},
				resource.Attribute{
					Name:        "fixed_ip",
					Description: `(Optional) An IP address to assosciate with the port. _NOTE_: This option cannot be used with port_id. You must specifiy a network_id. The IP address must lie in a range on the supplied network. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "port_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `See Argument Reference above.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "port_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `See Argument Reference above.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloudstack_compute_keypair_v2",
			Category:         "Compute Resources",
			ShortDescription: `Manages a V2 keypair resource within HuaweiCloudStack.`,
			Description:      ``,
			Keywords: []string{
				"compute",
				"keypair",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 Compute client. Keypairs are associated with accounts, but a Compute client is needed to create one. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new keypair.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A unique name for the keypair. Changing this creates a new keypair.`,
				},
				resource.Attribute{
					Name:        "public_key",
					Description: `(Required) A pregenerated OpenSSH-formatted public key. Changing this creates a new keypair. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "public_key",
					Description: `See Argument Reference above. ## Import Keypairs can be imported using the ` + "`" + `name` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import huaweicloudstack_compute_keypair_v2.my-keypair test-keypair ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "public_key",
					Description: `See Argument Reference above. ## Import Keypairs can be imported using the ` + "`" + `name` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import huaweicloudstack_compute_keypair_v2.my-keypair test-keypair ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloudstack_compute_servergroup_v2",
			Category:         "Compute Resources",
			ShortDescription: `Manages a V2 Server Group resource within HuaweiCloudStack.`,
			Description:      ``,
			Keywords: []string{
				"compute",
				"servergroup",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 Compute client. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new server group.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A unique name for the server group. Changing this creates a new server group.`,
				},
				resource.Attribute{
					Name:        "policies",
					Description: `(Required) The set of policies for the server group. Only two two policies are available right now, and both are mutually exclusive. See the Policies section for more information. Changing this creates a new server group. ## Policies`,
				},
				resource.Attribute{
					Name:        "affinity",
					Description: `All instances/servers launched in this group will be hosted on the same compute node.`,
				},
				resource.Attribute{
					Name:        "anti-affinity",
					Description: `All instances/servers launched in this group will be hosted on different compute nodes. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "policies",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "members",
					Description: `The instances that are part of this server group. ## Import Server Groups can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import huaweicloudstack_compute_servergroup_v2.test-sg 1bc30ee9-9d5b-4c30-bdd5-7f1e663f5edf ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "policies",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "members",
					Description: `The instances that are part of this server group. ## Import Server Groups can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import huaweicloudstack_compute_servergroup_v2.test-sg 1bc30ee9-9d5b-4c30-bdd5-7f1e663f5edf ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloudstack_compute_volume_attach_v2",
			Category:         "Compute Resources",
			ShortDescription: `Attaches a Block Storage Volume to an Instance.`,
			Description:      ``,
			Keywords: []string{
				"compute",
				"volume",
				"attach",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 Compute client. A Compute client is needed to create a volume attachment. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new volume attachment.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Required) The ID of the Instance to attach the Volume to.`,
				},
				resource.Attribute{
					Name:        "volume_id",
					Description: `(Required) The ID of the Volume to attach to an Instance.`,
				},
				resource.Attribute{
					Name:        "device",
					Description: `(Optional) The device of the volume attachment (ex: ` + "`" + `/dev/vdc` + "`" + `). _NOTE_: Being able to specify a device is dependent upon the hypervisor in use. There is a chance that the device specified in Terraform will not be the same device the hypervisor chose. If this happens, Terraform will wish to update the device upon subsequent applying which will cause the volume to be detached and reattached indefinitely. Please use with caution. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "volume_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "device",
					Description: `See Argument Reference above. _NOTE_: The correctness of this information is dependent upon the hypervisor in use. In some cases, this should not be used as an authoritative piece of information. ## Import Volume Attachments can be imported using the Instance ID and Volume ID separated by a slash, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import huaweicloudstack_compute_volume_attach_v2.va_1 89c60255-9bd6-460c-822a-e2b959ede9d2/45670584-225f-46c3-b33e-6707b589b666 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "volume_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "device",
					Description: `See Argument Reference above. _NOTE_: The correctness of this information is dependent upon the hypervisor in use. In some cases, this should not be used as an authoritative piece of information. ## Import Volume Attachments can be imported using the Instance ID and Volume ID separated by a slash, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import huaweicloudstack_compute_volume_attach_v2.va_1 89c60255-9bd6-460c-822a-e2b959ede9d2/45670584-225f-46c3-b33e-6707b589b666 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloudstack_lb_certificate_v2",
			Category:         "Elastic Load Balancer Resources",
			ShortDescription: `Manages a V2 certificate resource within HuaweiCloudStack.`,
			Description:      ``,
			Keywords: []string{
				"elastic",
				"load",
				"balancer",
				"lb",
				"certificate",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 Networking client. A Networking client is needed to create an LB certificate. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new LB certificate.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Human-readable name for the Certificate. Does not have to be unique.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Human-readable description for the Certificate.`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `(Optional) The domain of the Certificate.`,
				},
				resource.Attribute{
					Name:        "private_key",
					Description: `(Required) The private encrypted key of the Certificate, PEM format.`,
				},
				resource.Attribute{
					Name:        "certificate",
					Description: `(Required) The public encrypted key of the Certificate, PEM format. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "private_key",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "certificate",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "update_time",
					Description: `Indicates the update time.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Indicates the creation time.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "private_key",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "certificate",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "update_time",
					Description: `Indicates the update time.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Indicates the creation time.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloudstack_lb_l7policy_v2",
			Category:         "Elastic Load Balancer Resources",
			ShortDescription: `Manages a V2 L7 Policy resource within HuaweiCloudStack.`,
			Description:      ``,
			Keywords: []string{
				"elastic",
				"load",
				"balancer",
				"lb",
				"l7policy",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 Networking client. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new L7 Policy.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Optional) Required for admins. The UUID of the tenant who owns the L7 Policy. Only administrative users can specify a tenant UUID other than their own. Changing this creates a new L7 Policy.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Human-readable name for the L7 Policy. Does not have to be unique.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Human-readable description for the L7 Policy.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Optional) The L7 Policy action. The value can only be REDIRECT\_TO\_POOL.`,
				},
				resource.Attribute{
					Name:        "listener_id",
					Description: `(Required) The Listener on which the L7 Policy will be associated with. Changing this creates a new L7 Policy.`,
				},
				resource.Attribute{
					Name:        "position",
					Description: `(Optional) The position of this policy on the listener. Positions start at 1. Changing this creates a new L7 Policy.`,
				},
				resource.Attribute{
					Name:        "redirect_pool_id",
					Description: `(Required) Requests matching this policy will be redirected to the pool with this ID.`,
				},
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `(Optional) The administrative state of the L7 Policy. This value can only be true (UP). ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique ID for the L7 Policy.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "listener_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "position",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "redirect_pool_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `See Argument Reference above. ## Import Load Balancer L7 Policy can be imported using the L7 Policy ID, e.g.: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import huaweicloudstack_lb_l7policy_v2.l7policy_1 8a7a79c2-cf17-4e65-b2ae-ddc8bfcf6c74 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The unique ID for the L7 Policy.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "listener_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "position",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "redirect_pool_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `See Argument Reference above. ## Import Load Balancer L7 Policy can be imported using the L7 Policy ID, e.g.: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import huaweicloudstack_lb_l7policy_v2.l7policy_1 8a7a79c2-cf17-4e65-b2ae-ddc8bfcf6c74 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloudstack_lb_l7rule_v2",
			Category:         "Elastic Load Balancer Resources",
			ShortDescription: `Manages a V2 l7rule resource within HuaweiCloudStack.`,
			Description:      ``,
			Keywords: []string{
				"elastic",
				"load",
				"balancer",
				"lb",
				"l7rule",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 Networking client. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new L7 Rule.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Optional) Required for admins. The UUID of the tenant who owns the L7 Rule. Only administrative users can specify a tenant UUID other than their own. Changing this creates a new L7 Rule.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Human-readable description for the L7 Rule.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) The L7 Rule type - can either be HOST\_NAME or PATH. Changing this creates a new L7 Rule.`,
				},
				resource.Attribute{
					Name:        "compare_type",
					Description: `(Required) The comparison type for the L7 rule - can either be STARTS\_WITH, EQUAL_TO or REGEX`,
				},
				resource.Attribute{
					Name:        "l7policy_id",
					Description: `(Required) The ID of the L7 Policy to query. Changing this creates a new L7 Rule.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) The value to use for the comparison. For example, the file type to compare.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Optional) The key to use for the comparison. For example, the name of the cookie to evaluate. Valid when ` + "`" + `type` + "`" + ` is set to COOKIE or HEADER. Changing this creates a new L7 Rule.`,
				},
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `(Optional) The administrative state of the L7 Rule. The value can only be true (UP). ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique ID for the L7 Rule.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "compare_type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "l7policy_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "invert",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "listener_id",
					Description: `The ID of the Listener owning this resource. ## Import Load Balancer L7 Rule can be imported using the L7 Policy ID and L7 Rule ID separated by a slash, e.g.: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import huaweicloudstack_lb_l7rule_v2.l7rule_1 e0bd694a-abbe-450e-b329-0931fd1cc5eb/4086b0c9-b18c-4d1c-b6b8-4c56c3ad2a9e ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The unique ID for the L7 Rule.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "compare_type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "l7policy_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "invert",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "listener_id",
					Description: `The ID of the Listener owning this resource. ## Import Load Balancer L7 Rule can be imported using the L7 Policy ID and L7 Rule ID separated by a slash, e.g.: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import huaweicloudstack_lb_l7rule_v2.l7rule_1 e0bd694a-abbe-450e-b329-0931fd1cc5eb/4086b0c9-b18c-4d1c-b6b8-4c56c3ad2a9e ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloudstack_lb_listener_v2",
			Category:         "Elastic Load Balancer Resources",
			ShortDescription: `Manages a V2 listener resource within HuaweiCloudStack.`,
			Description:      ``,
			Keywords: []string{
				"elastic",
				"load",
				"balancer",
				"lb",
				"listener",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 Networking client. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new Listener.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Required) The listening protocol. Converged ELB in Region Type I and Region Type II supports TCP, UDP, HTTP, and TERMINATED_HTTPS. Non-converged ELB in Region Type II supports TCP and HTTP. Changing this creates a new Listener.`,
				},
				resource.Attribute{
					Name:        "protocol_port",
					Description: `(Required) The port on which to listen for client traffic. Must be an integer in the range of 1-65535. Changing this creates a new Listener.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Optional) Required for admins. The UUID of the tenant who owns the Listener. Only administrative users can specify a tenant UUID other than their own. Changing this creates a new Listener.`,
				},
				resource.Attribute{
					Name:        "loadbalancer_id",
					Description: `(Required) The load balancer on which to provision this Listener. Changing this creates a new Listener.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Human-readable name for the Listener. Does not have to be unique.`,
				},
				resource.Attribute{
					Name:        "default_pool_id",
					Description: `(Optional) The ID of the default pool with which the Listener is associated. Changing this creates a new Listener.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Human-readable description for the Listener.`,
				},
				resource.Attribute{
					Name:        "connection_limit",
					Description: `(Optional) The maximum number of connections allowed for the Listener. A valid value is from -1 to 2147483647. The default value for this attribute will be -1, indicating an infinite limit.`,
				},
				resource.Attribute{
					Name:        "default_tls_container_ref",
					Description: `(Optional) A reference to a Barbican Secrets container which stores TLS information. This is required if the protocol is ` + "`" + `TERMINATED_HTTPS` + "`" + `. For converged ELB in Region Type I and Region Type II, enter a certificate ID.`,
				},
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `(Optional) The administrative state of the Listener. A valid value is true (UP) or false (DOWN). ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique ID for the Listener.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "protocol_port",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "default_port_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "connection_limit",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "default_tls_container_ref",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `See Argument Reference above.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The unique ID for the Listener.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "protocol_port",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "default_port_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "connection_limit",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "default_tls_container_ref",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `See Argument Reference above.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloudstack_lb_loadbalancer_v2",
			Category:         "Elastic Load Balancer Resources",
			ShortDescription: `Manages a V2 loadbalancer resource within HuaweiCloudStack.`,
			Description:      ``,
			Keywords: []string{
				"elastic",
				"load",
				"balancer",
				"lb",
				"loadbalancer",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 Networking client. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new loadbalancer.`,
				},
				resource.Attribute{
					Name:        "vip_subnet_id",
					Description: `(Required) The network on which to allocate the Loadbalancer's address. A tenant can only create Loadbalancers on networks authorized by policy (e.g. networks that belong to them or networks that are shared). Changing this creates a new loadbalancer.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Human-readable name for the Loadbalancer. Does not have to be unique.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Human-readable description for the Loadbalancer.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Optional) Required for admins. The UUID of the tenant who owns the Loadbalancer. Only administrative users can specify a tenant UUID other than their own. Changing this creates a new loadbalancer.`,
				},
				resource.Attribute{
					Name:        "vip_address",
					Description: `(Optional) The ip address of the load balancer. Changing this creates a new loadbalancer.`,
				},
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `(Optional) The administrative state of the Loadbalancer. A valid value is true (UP) or false (DOWN).`,
				},
				resource.Attribute{
					Name:        "loadbalancer_provider",
					Description: `(Optional) The name of the provider. For Region Type I, only VLB is supported. For Region Type II, VLB, native load balancers, and load balancers of third-party providers are supported. Changing this creates a new loadbalancer. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique ID for the load balancer.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "vip_subnet_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "vip_address",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "loadbalancer_provider",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "vip_port_id",
					Description: `The Port ID of the Load Balancer IP.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The unique ID for the load balancer.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "vip_subnet_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "vip_address",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "loadbalancer_provider",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "vip_port_id",
					Description: `The Port ID of the Load Balancer IP.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloudstack_lb_member_v2",
			Category:         "Elastic Load Balancer Resources",
			ShortDescription: `Manages a V2 member resource within HuaweiCloudStack.`,
			Description:      ``,
			Keywords: []string{
				"elastic",
				"load",
				"balancer",
				"lb",
				"member",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "pool_id",
					Description: `(Required) The id of the pool that this member will be assigned to.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Required) The subnet in which to access the member`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Human-readable name for the member.`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `(Required) The IP address of the member to receive traffic from the load balancer. Changing this creates a new member.`,
				},
				resource.Attribute{
					Name:        "protocol_port",
					Description: `(Required) The port on which to listen for client traffic. Changing this creates a new member.`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `(Optional) A positive integer value that indicates the relative portion of traffic that this member should receive from the pool. For example, a member with a weight of 10 receives five times as much traffic as a member with a weight of 2.`,
				},
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `(Optional) The administrative state of the member. A valid value is true (UP) or false (DOWN). ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique ID for the member.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "pool_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "protocol_port",
					Description: `See Argument Reference above.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The unique ID for the member.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "pool_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "protocol_port",
					Description: `See Argument Reference above.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloudstack_lb_monitor_v2",
			Category:         "Elastic Load Balancer Resources",
			ShortDescription: `Manages a V2 monitor resource within HuaweiCloudStack.`,
			Description:      ``,
			Keywords: []string{
				"elastic",
				"load",
				"balancer",
				"lb",
				"monitor",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "pool_id",
					Description: `(Required) The id of the pool that this monitor will be assigned to.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The Name of the Monitor.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) The type of protocol. Converged ELB in Region Type I and Region Type II supports TCP, UDP_CONNECT, or HTTP. Non-converged ELB in Region Type II supports TCP, PING, or HTTP. For Region Type I, if protocol of the listener is set to UDP, type of the health check must be set to UDP_CONNECT. Changing this creates a new monitor.`,
				},
				resource.Attribute{
					Name:        "delay",
					Description: `(Required) The interval in seconds between health check. A valid value is from 1 to 50.`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `(Required) Maximum number of seconds for a monitor to wait for a ping reply before it times out. The value must be less than the delay value.`,
				},
				resource.Attribute{
					Name:        "max_retries",
					Description: `(Required) Number of permissible ping failures before changing the member's status to INACTIVE. Must be a number between 1 and 10.`,
				},
				resource.Attribute{
					Name:        "url_path",
					Description: `(Optional) Required for HTTP types. URI path that will be accessed if monitor type is HTTP.`,
				},
				resource.Attribute{
					Name:        "expected_codes",
					Description: `(Optional) Required for HTTP types. Expected HTTP codes for a passing HTTP(S) monitor. You can either specify a single status like "200", or a range like "200-202".`,
				},
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `(Optional) The administrative state of the monitor. A valid value is true (UP) or false (DOWN). ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique ID for the monitor.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "delay",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "max_retries",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "url_path",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "http_method",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "expected_codes",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `See Argument Reference above.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The unique ID for the monitor.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "delay",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "max_retries",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "url_path",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "http_method",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "expected_codes",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `See Argument Reference above.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloudstack_lb_pool_v2",
			Category:         "Elastic Load Balancer Resources",
			ShortDescription: `Manages a V2 pool resource within HuaweiCloudStack.`,
			Description:      ``,
			Keywords: []string{
				"elastic",
				"load",
				"balancer",
				"lb",
				"pool",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Human-readable name for the pool.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Human-readable description for the pool.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Required) The IP protocol, can either be TCP, HTTP or UDP. Changing this creates a new pool.`,
				},
				resource.Attribute{
					Name:        "loadbalancer_id",
					Description: `(Optional) The load balancer on which to provision this pool. Changing this creates a new pool. Note: One of LoadbalancerID or ListenerID must be provided.`,
				},
				resource.Attribute{
					Name:        "listener_id",
					Description: `(Optional) The Listener on which the members of the pool will be associated with. Changing this creates a new pool. Note: One of LoadbalancerID or ListenerID must be provided.`,
				},
				resource.Attribute{
					Name:        "lb_method",
					Description: `(Required) The load balancing algorithm to distribute traffic to the pool's members. Must be one of ROUND_ROBIN, LEAST_CONNECTIONS, or SOURCE_IP.`,
				},
				resource.Attribute{
					Name:        "persistence",
					Description: `Omit this field to prevent session persistence. Indicates whether connections in the same session will be processed by the same Pool member or not. Changing this creates a new pool.`,
				},
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `(Optional) The administrative state of the pool. A valid value is true (UP) or false (DOWN). The ` + "`" + `persistence` + "`" + ` argument supports:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) The type of persistence mode. The current specification supports SOURCE_IP, HTTP_COOKIE, and APP_COOKIE.`,
				},
				resource.Attribute{
					Name:        "cookie_name",
					Description: `(Optional) The name of the cookie if persistence mode is set appropriately. It's only supported in the ` + "`" + `APP_COOKIE` + "`" + ` type. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique ID for the pool.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "lb_method",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "persistence",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `See Argument Reference above.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The unique ID for the pool.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "lb_method",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "persistence",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `See Argument Reference above.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloudstack_lb_whitelist_v2",
			Category:         "Elastic Load Balancer Resources",
			ShortDescription: `Manages an ELB whitelist resource within HuaweiCloudStack.`,
			Description:      ``,
			Keywords: []string{
				"elastic",
				"load",
				"balancer",
				"lb",
				"whitelist",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Optional) Required for admins. The UUID of the tenant who owns the whitelist. Only administrative users can specify a tenant UUID other than their own. Changing this creates a new whitelist.`,
				},
				resource.Attribute{
					Name:        "listener_id",
					Description: `(Required) The Listener ID that the whitelist will be associated with. Changing this creates a new whitelist.`,
				},
				resource.Attribute{
					Name:        "enable_whitelist",
					Description: `(Optional) Specify whether to enable access control.`,
				},
				resource.Attribute{
					Name:        "whitelist",
					Description: `(Optional) Specifies the IP addresses in the whitelist. Use commas(,) to separate the multiple IP addresses. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique ID for the whitelist.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "listener_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "enable_whitelist",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "whitelist",
					Description: `See Argument Reference above.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The unique ID for the whitelist.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "listener_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "enable_whitelist",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "whitelist",
					Description: `See Argument Reference above.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloudstack_networking_floatingip_associate_v2",
			Category:         "Networking Resources",
			ShortDescription: `Associates a Floating IP to a Port`,
			Description:      ``,
			Keywords: []string{
				"networking",
				"floatingip",
				"associate",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 Networking client. A Networking client is needed to create a floating IP that can be used with another networking resource, such as a load balancer. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new floating IP (which may or may not have a different address).`,
				},
				resource.Attribute{
					Name:        "floating_ip",
					Description: `(Required) IP Address of an existing floating IP.`,
				},
				resource.Attribute{
					Name:        "port_id",
					Description: `(Required) ID of an existing port with at least one IP address to associate with this floating IP. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "floating_ip",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "port_id",
					Description: `See Argument Reference above. ## Import Floating IP associations can be imported using the ` + "`" + `id` + "`" + ` of the floating IP, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import huaweicloudstack_networking_floatingip_associate_v2.fip 2c7f39f3-702b-48d1-940c-b50384177ee1 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "floating_ip",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "port_id",
					Description: `See Argument Reference above. ## Import Floating IP associations can be imported using the ` + "`" + `id` + "`" + ` of the floating IP, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import huaweicloudstack_networking_floatingip_associate_v2.fip 2c7f39f3-702b-48d1-940c-b50384177ee1 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloudstack_networking_floatingip_v2",
			Category:         "Networking Resources",
			ShortDescription: `Manages a V2 floating IP resource within HuaweiCloudStack.`,
			Description:      ``,
			Keywords: []string{
				"networking",
				"floatingip",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 Networking client. A Networking client is needed to create a floating IP that can be used with another networking resource, such as a load balancer. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new floating IP (which may or may not have a different address).`,
				},
				resource.Attribute{
					Name:        "pool",
					Description: `(Optional) The name of the pool from which to obtain the floating IP. Only admin_external_net is valid. Changing this creates a new floating IP.`,
				},
				resource.Attribute{
					Name:        "port_id",
					Description: `(Optional) ID of an existing port with at least one IP address to associate with this floating IP.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Optional) The target tenant ID in which to allocate the floating IP, if you specify this together with a port_id, make sure the target port belongs to the same tenant. Changing this creates a new floating IP (which may or may not have a different address)`,
				},
				resource.Attribute{
					Name:        "fixed_ip",
					Description: `Fixed IP of the port to associate with this floating IP. Required if the port has multiple fixed IPs. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "pool",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `The actual floating IP address itself.`,
				},
				resource.Attribute{
					Name:        "port_id",
					Description: `ID of associated port.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `the ID of the tenant in which to create the floating IP.`,
				},
				resource.Attribute{
					Name:        "fixed_ip",
					Description: `The fixed IP which the floating IP maps to. ## Import Floating IPs can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import huaweicloudstack_networking_floatingip_v2.floatip_1 2c7f39f3-702b-48d1-940c-b50384177ee1 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "pool",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `The actual floating IP address itself.`,
				},
				resource.Attribute{
					Name:        "port_id",
					Description: `ID of associated port.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `the ID of the tenant in which to create the floating IP.`,
				},
				resource.Attribute{
					Name:        "fixed_ip",
					Description: `The fixed IP which the floating IP maps to. ## Import Floating IPs can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import huaweicloudstack_networking_floatingip_v2.floatip_1 2c7f39f3-702b-48d1-940c-b50384177ee1 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloudstack_networking_network_v2",
			Category:         "Networking Resources",
			ShortDescription: `Manages a V2 Neutron network resource within HuaweiCloudStack.`,
			Description:      ``,
			Keywords: []string{
				"networking",
				"network",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 Networking client. A Networking client is needed to create a Neutron network. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new network.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the network. Changing this updates the name of the existing network.`,
				},
				resource.Attribute{
					Name:        "shared",
					Description: `(Optional) Specifies whether the network resource can be accessed by any tenant or not. Changing this updates the sharing capabalities of the existing network.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Optional) The owner of the network. Required if admin wants to create a network for another tenant. Changing this creates a new network.`,
				},
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `(Optional) The administrative state of the network. The value must be "true".`,
				},
				resource.Attribute{
					Name:        "segments",
					Description: `(Optional) An array of one or more provider segment objects. The ` + "`" + `segments` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "physical_network",
					Description: `The phisical network where this network is implemented.`,
				},
				resource.Attribute{
					Name:        "segmentation_id",
					Description: `An isolated segment on the physical network.`,
				},
				resource.Attribute{
					Name:        "network_type",
					Description: `The type of physical network. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "shared",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `See Argument Reference above. ## Import Networks can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import huaweicloudstack_networking_network_v2.network_1 d90ce693-5ccf-4136-a0ed-152ce412b6b9 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "shared",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `See Argument Reference above. ## Import Networks can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import huaweicloudstack_networking_network_v2.network_1 d90ce693-5ccf-4136-a0ed-152ce412b6b9 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloudstack_networking_port_v2",
			Category:         "Networking Resources",
			ShortDescription: `Manages a V2 port resource within HuaweiCloudStack.`,
			Description:      ``,
			Keywords: []string{
				"networking",
				"port",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 networking client. A networking client is needed to create a port. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new port.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) A unique name for the port. Changing this updates the ` + "`" + `name` + "`" + ` of an existing port.`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `(Required) The ID of the network to attach the port to. Changing this creates a new port.`,
				},
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `(Optional) Administrative up/down status for the port (must be "true" or "false" if provided). Changing this updates the ` + "`" + `admin_state_up` + "`" + ` of an existing port.`,
				},
				resource.Attribute{
					Name:        "mac_address",
					Description: `(Optional) Specify a specific MAC address for the port. Changing this creates a new port.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Optional) The owner of the Port. Required if admin wants to create a port for another tenant. Changing this creates a new port.`,
				},
				resource.Attribute{
					Name:        "device_owner",
					Description: `(Optional) The device owner of the Port. Changing this creates a new port.`,
				},
				resource.Attribute{
					Name:        "security_group_ids",
					Description: `(Optional - Conflicts with ` + "`" + `no_security_groups` + "`" + `) A list of security group IDs to apply to the port. The security groups must be specified by ID and not name (as opposed to how they are configured with the Compute Instance).`,
				},
				resource.Attribute{
					Name:        "no_security_groups",
					Description: `(Optional - Conflicts with ` + "`" + `security_group_ids` + "`" + `) If set to ` + "`" + `true` + "`" + `, then no security groups are applied to the port. If set to ` + "`" + `false` + "`" + ` and no ` + "`" + `security_group_ids` + "`" + ` are specified, then the Port will yield to the default behavior of the Networking service, which is to usually apply the "default" security group.`,
				},
				resource.Attribute{
					Name:        "device_id",
					Description: `(Optional) The ID of the device attached to the port. Changing this creates a new port.`,
				},
				resource.Attribute{
					Name:        "fixed_ip",
					Description: `(Optional) An array of desired IPs for this port. The structure is described below.`,
				},
				resource.Attribute{
					Name:        "allowed_address_pairs",
					Description: `(Optional) An IP/MAC Address pair of additional IP addresses that can be active on this port. The structure is described below. The ` + "`" + `fixed_ip` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Required) Subnet in which to allocate IP address for this port.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `(Optional) IP address desired in the subnet for this port. If you don't specify ` + "`" + `ip_address` + "`" + `, an available IP address from the specified subnet will be allocated to this port. This field will not be populated if it is left blank. To retrieve the assigned IP address, use the ` + "`" + `all_fixed_ips` + "`" + ` attribute. The ` + "`" + `allowed_address_pairs` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `(Required) The additional IP address.`,
				},
				resource.Attribute{
					Name:        "mac_address",
					Description: `(Optional) The additional MAC address. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "mac_address",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "device_owner",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "security_group_ids",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "device_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "fixed_ip",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "all_fixed_ips",
					Description: `The collection of Fixed IP addresses on the port in the order returned by the Network v2 API.`,
				},
				resource.Attribute{
					Name:        "all_security_group_ids",
					Description: `The collection of Security Group IDs on the port which have been explicitly and implicitly added. ## Import Ports can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import huaweicloudstack_networking_port_v2.port_1 eae26a3e-1c33-4cc1-9c31-0cd729c438a1 ` + "`" + `` + "`" + `` + "`" + ` ## Notes ### Ports and Instances There are some notes to consider when connecting Instances to networks using Ports. Please see the ` + "`" + `huaweicloudstack_compute_instance_v2` + "`" + ` documentation for further documentation.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "mac_address",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "device_owner",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "security_group_ids",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "device_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "fixed_ip",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "all_fixed_ips",
					Description: `The collection of Fixed IP addresses on the port in the order returned by the Network v2 API.`,
				},
				resource.Attribute{
					Name:        "all_security_group_ids",
					Description: `The collection of Security Group IDs on the port which have been explicitly and implicitly added. ## Import Ports can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import huaweicloudstack_networking_port_v2.port_1 eae26a3e-1c33-4cc1-9c31-0cd729c438a1 ` + "`" + `` + "`" + `` + "`" + ` ## Notes ### Ports and Instances There are some notes to consider when connecting Instances to networks using Ports. Please see the ` + "`" + `huaweicloudstack_compute_instance_v2` + "`" + ` documentation for further documentation.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloudstack_networking_router_interface_v2",
			Category:         "Networking Resources",
			ShortDescription: `Manages a V2 router interface resource within HuaweiCloudStack.`,
			Description:      ``,
			Keywords: []string{
				"networking",
				"router",
				"interface",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 networking client. A networking client is needed to create a router. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new router interface.`,
				},
				resource.Attribute{
					Name:        "router_id",
					Description: `(Required) ID of the router this interface belongs to. Changing this creates a new router interface.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `ID of the subnet this interface connects to. Changing this creates a new router interface.`,
				},
				resource.Attribute{
					Name:        "port_id",
					Description: `ID of the port this interface connects to. Changing this creates a new router interface. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "router_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "port_id",
					Description: `See Argument Reference above. ## Import Router Interfaces can be imported using the port ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ openstack port list --router <router name or id> $ terraform import huaweicloudstack_networking_router_interface_v2.int_1 <port id from above output> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "router_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "port_id",
					Description: `See Argument Reference above. ## Import Router Interfaces can be imported using the port ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ openstack port list --router <router name or id> $ terraform import huaweicloudstack_networking_router_interface_v2.int_1 <port id from above output> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloudstack_networking_router_route_v2",
			Category:         "Networking Resources",
			ShortDescription: `Creates a routing entry on a HuaweiCloudStack V2 router.`,
			Description:      ``,
			Keywords: []string{
				"networking",
				"router",
				"route",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 networking client. A networking client is needed to configure a routing entry on a router. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new routing entry.`,
				},
				resource.Attribute{
					Name:        "router_id",
					Description: `(Required) ID of the router this routing entry belongs to. Changing this creates a new routing entry.`,
				},
				resource.Attribute{
					Name:        "destination_cidr",
					Description: `(Required) CIDR block to match on the packet’s destination IP. Changing this creates a new routing entry.`,
				},
				resource.Attribute{
					Name:        "next_hop",
					Description: `(Required) IP address of the next hop gateway. Changing this creates a new routing entry. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "router_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "destination_cidr",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "next_hop",
					Description: `See Argument Reference above. ## Notes The ` + "`" + `next_hop` + "`" + ` IP address must be directly reachable from the router at the ` + "`" + `` + "`" + `huaweicloudstack_networking_router_route_v2` + "`" + `` + "`" + ` resource creation time. You can ensure that by explicitly specifying a dependency on the ` + "`" + `` + "`" + `huaweicloudstack_networking_router_interface_v2` + "`" + `` + "`" + ` resource that connects the next hop to the router, as in the example above. ## Import Routing entries can be imported using a combined ID using the following format: ` + "`" + `` + "`" + `<router_id>-route-<destination_cidr>-<next_hop>` + "`" + `` + "`" + ` ` + "`" + `` + "`" + `` + "`" + ` $ terraform import huaweicloudstack_networking_router_route_v2.router_route_1 686fe248-386c-4f70-9f6c-281607dad079-route-10.0.1.0/24-192.168.199.25 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "router_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "destination_cidr",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "next_hop",
					Description: `See Argument Reference above. ## Notes The ` + "`" + `next_hop` + "`" + ` IP address must be directly reachable from the router at the ` + "`" + `` + "`" + `huaweicloudstack_networking_router_route_v2` + "`" + `` + "`" + ` resource creation time. You can ensure that by explicitly specifying a dependency on the ` + "`" + `` + "`" + `huaweicloudstack_networking_router_interface_v2` + "`" + `` + "`" + ` resource that connects the next hop to the router, as in the example above. ## Import Routing entries can be imported using a combined ID using the following format: ` + "`" + `` + "`" + `<router_id>-route-<destination_cidr>-<next_hop>` + "`" + `` + "`" + ` ` + "`" + `` + "`" + `` + "`" + ` $ terraform import huaweicloudstack_networking_router_route_v2.router_route_1 686fe248-386c-4f70-9f6c-281607dad079-route-10.0.1.0/24-192.168.199.25 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloudstack_networking_router_v2",
			Category:         "Networking Resources",
			ShortDescription: `Manages a V2 router resource within HuaweiCloudStack.`,
			Description:      ``,
			Keywords: []string{
				"networking",
				"router",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 networking client. A networking client is needed to create a router. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new router.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) A unique name for the router. Changing this updates the ` + "`" + `name` + "`" + ` of an existing router.`,
				},
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `(Optional) Administrative up/down status for the router (must be "true" or "false" if provided). Changing this updates the ` + "`" + `admin_state_up` + "`" + ` of an existing router.`,
				},
				resource.Attribute{
					Name:        "distributed",
					Description: `(Optional) Indicates whether or not to create a distributed router. The default policy setting in Neutron restricts usage of this property to administrative users only.`,
				},
				resource.Attribute{
					Name:        "external_network_id",
					Description: `(Optional) The network UUID of an external gateway for the router. A router with an external gateway is required if any compute instances or load balancers will be using floating IPs. Changing this updates the external gateway of the router.`,
				},
				resource.Attribute{
					Name:        "enable_snat",
					Description: `(Optional) Enable Source NAT for the router. Valid values are "true" or "false". An ` + "`" + `external_network_id` + "`" + ` has to be set in order to set this property. Changing this updates the ` + "`" + `enable_snat` + "`" + ` of the router.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Optional) The owner of the floating IP. Required if admin wants to create a router for another tenant. Changing this creates a new router. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the router.`,
				},
				resource.Attribute{
					Name:        "region",
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
					Name:        "external_network_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "enable_snat",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above. ## Import Routers can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import huaweicloudstack_networking_router_v2.router_1 014395cd-89fc-4c9b-96b7-13d1ee79dad2 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the router.`,
				},
				resource.Attribute{
					Name:        "region",
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
					Name:        "external_network_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "enable_snat",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above. ## Import Routers can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import huaweicloudstack_networking_router_v2.router_1 014395cd-89fc-4c9b-96b7-13d1ee79dad2 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloudstack_networking_secgroup_rule_v2",
			Category:         "Networking Resources",
			ShortDescription: `Manages a V2 Neutron security group rule resource within HuaweiCloudStack.`,
			Description:      ``,
			Keywords: []string{
				"networking",
				"secgroup",
				"rule",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 networking client. A networking client is needed to create a port. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new security group rule.`,
				},
				resource.Attribute{
					Name:        "direction",
					Description: `(Required) The direction of the rule, valid values are __ingress__ or __egress__. Changing this creates a new security group rule.`,
				},
				resource.Attribute{
					Name:        "ethertype",
					Description: `(Required) The layer 3 protocol type, valid values are __IPv4__ or __IPv6__. Changing this creates a new security group rule.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Optional) The layer 4 protocol type, valid values are following. Changing this creates a new security group rule. This is required if you want to specify a port range.`,
				},
				resource.Attribute{
					Name:        "port_range_min",
					Description: `(Optional) The lower part of the allowed port range, valid integer value needs to be between 1 and 65535. Changing this creates a new security group rule.`,
				},
				resource.Attribute{
					Name:        "port_range_max",
					Description: `(Optional) The higher part of the allowed port range, valid integer value needs to be between 1 and 65535. Changing this creates a new security group rule.`,
				},
				resource.Attribute{
					Name:        "remote_ip_prefix",
					Description: `(Optional) The remote CIDR, the value needs to be a valid CIDR (i.e. 192.168.0.0/16). Changing this creates a new security group rule.`,
				},
				resource.Attribute{
					Name:        "remote_group_id",
					Description: `(Optional) The remote group id, the value needs to be an Openstack ID of a security group in the same tenant. Changing this creates a new security group rule.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `(Required) The security group id the rule should belong to, the value needs to be an Openstack ID of a security group in the same tenant. Changing this creates a new security group rule.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Optional) The owner of the security group. Required if admin wants to create a port for another tenant. Changing this creates a new security group rule. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "direction",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "ethertype",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "port_range_min",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "port_range_max",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "remote_ip_prefix",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "remote_group_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above. ## Import Security Group Rules can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import huaweicloudstack_networking_secgroup_rule_v2.secgroup_rule_1 aeb68ee3-6e9d-4256-955c-9584a6212745 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "direction",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "ethertype",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "port_range_min",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "port_range_max",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "remote_ip_prefix",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "remote_group_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above. ## Import Security Group Rules can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import huaweicloudstack_networking_secgroup_rule_v2.secgroup_rule_1 aeb68ee3-6e9d-4256-955c-9584a6212745 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloudstack_networking_secgroup_v2",
			Category:         "Networking Resources",
			ShortDescription: `Manages a V2 security group resource within HuaweiCloudStack.`,
			Description:      ``,
			Keywords: []string{
				"networking",
				"secgroup",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 networking client. A networking client is needed to create a port. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new security group.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A unique name for the security group.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A unique name for the security group.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Optional) The owner of the security group. Required if admin wants to create a port for another tenant. Changing this creates a new security group.`,
				},
				resource.Attribute{
					Name:        "delete_default_rules",
					Description: `(Optional) Whether or not to delete the default egress security rules. This is ` + "`" + `false` + "`" + ` by default. See the below note for more information. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above. ## Default Security Group Rules In most cases, HuaweiCloudStack will create some egress security group rules for each new security group. These security group rules will not be managed by Terraform, so if you prefer to have`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above. ## Default Security Group Rules In most cases, HuaweiCloudStack will create some egress security group rules for each new security group. These security group rules will not be managed by Terraform, so if you prefer to have`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloudstack_networking_subnet_v2",
			Category:         "Networking Resources",
			ShortDescription: `Manages a V2 Neutron subnet resource within HuaweiCloudStack.`,
			Description:      ``,
			Keywords: []string{
				"networking",
				"subnet",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 Networking client. A Networking client is needed to create a Neutron subnet. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new subnet.`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `(Required) The UUID of the parent network. Changing this creates a new subnet.`,
				},
				resource.Attribute{
					Name:        "cidr",
					Description: `(Required) CIDR representing IP range for this subnet, based on IP version. Changing this creates a new subnet.`,
				},
				resource.Attribute{
					Name:        "ip_version",
					Description: `(Optional) IP version, either 4 (default) or 6. Changing this creates a new subnet.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the subnet. Changing this updates the name of the existing subnet.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Optional) The owner of the subnet. Required if admin wants to create a subnet for another tenant. Changing this creates a new subnet.`,
				},
				resource.Attribute{
					Name:        "allocation_pools",
					Description: `(Optional) An array of sub-ranges of CIDR available for dynamic allocation to ports. The allocation_pool object structure is documented below. Changing this creates a new subnet.`,
				},
				resource.Attribute{
					Name:        "gateway_ip",
					Description: `(Optional) Default gateway used by devices in this subnet. Leaving this blank and not setting ` + "`" + `no_gateway` + "`" + ` will cause a default gateway of ` + "`" + `.1` + "`" + ` to be used. Changing this updates the gateway IP of the existing subnet.`,
				},
				resource.Attribute{
					Name:        "no_gateway",
					Description: `(Optional) Do not set a gateway IP on this subnet. Changing this removes or adds a default gateway IP of the existing subnet.`,
				},
				resource.Attribute{
					Name:        "enable_dhcp",
					Description: `(Optional) The administrative state of the network. The value must be "true".`,
				},
				resource.Attribute{
					Name:        "dns_nameservers",
					Description: `(Optional) An array of DNS name server names used by hosts in this subnet. Changing this updates the DNS name servers for the existing subnet.`,
				},
				resource.Attribute{
					Name:        "host_routes",
					Description: `(Optional) An array of routes that should be used by devices with IPs from this subnet (not including local subnet route). The host_route object structure is documented below. Changing this updates the host routes for the existing subnet. The ` + "`" + `allocation_pools` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "start",
					Description: `(Required) The starting address.`,
				},
				resource.Attribute{
					Name:        "end",
					Description: `(Required) The ending address. The ` + "`" + `host_routes` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "destination_cidr",
					Description: `(Required) The destination CIDR.`,
				},
				resource.Attribute{
					Name:        "next_hop",
					Description: `(Required) The next hop in the route. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "cidr",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "ip_version",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "allocation_pools",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "gateway_ip",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "enable_dhcp",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "dns_nameservers",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "host_routes",
					Description: `See Argument Reference above. ## Import Subnets can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import huaweicloudstack_networking_subnet_v2.subnet_1 da4faf16-5546-41e4-8330-4d0002b74048 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "cidr",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "ip_version",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "allocation_pools",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "gateway_ip",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "enable_dhcp",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "dns_nameservers",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "host_routes",
					Description: `See Argument Reference above. ## Import Subnets can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import huaweicloudstack_networking_subnet_v2.subnet_1 da4faf16-5546-41e4-8330-4d0002b74048 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloudstack_networking_vip_associate_v2",
			Category:         "Networking Resources",
			ShortDescription: `Manages a V2 vip associate resource within HuaweiCloudStack.`,
			Description:      ``,
			Keywords: []string{
				"networking",
				"vip",
				"associate",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vip_id",
					Description: `(Required) The ID of vip to attach the port to. Changing this creates a new vip associate.`,
				},
				resource.Attribute{
					Name:        "port_ids",
					Description: `(Required) An array of one or more IDs of the ports to attach the vip to. Changing this creates a new vip associate. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "vip_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "port_ids",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "vip_subnet_id",
					Description: `The ID of the subnet this vip connects to.`,
				},
				resource.Attribute{
					Name:        "vip_ip_address",
					Description: `The IP address in the subnet for this vip.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "vip_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "port_ids",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "vip_subnet_id",
					Description: `The ID of the subnet this vip connects to.`,
				},
				resource.Attribute{
					Name:        "vip_ip_address",
					Description: `The IP address in the subnet for this vip.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloudstack_networking_vip_v2",
			Category:         "Networking Resources",
			ShortDescription: `Manages a V2 vip resource within HuaweiCloudStack.`,
			Description:      ``,
			Keywords: []string{
				"networking",
				"vip",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "network_id",
					Description: `(Required) The ID of the network to attach the vip to. Changing this creates a new vip.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Required) Subnet in which to allocate IP address for this vip. Changing this creates a new vip.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `(Optional) IP address desired in the subnet for this vip. If you don't specify ` + "`" + `ip_address` + "`" + `, an available IP address from the specified subnet will be allocated to this vip.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) A unique name for the vip. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of vip.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the vip.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `The tenant ID of the vip.`,
				},
				resource.Attribute{
					Name:        "device_owner",
					Description: `The device owner of the vip.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "network_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of vip.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the vip.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `The tenant ID of the vip.`,
				},
				resource.Attribute{
					Name:        "device_owner",
					Description: `The device owner of the vip.`,
				},
			},
		},
	}

	resourcesMap = map[string]int{

		"huaweicloudstack_as_configuration_v1":                0,
		"huaweicloudstack_as_group_v1":                        1,
		"huaweicloudstack_as_policy_v1":                       2,
		"huaweicloudstack_blockstorage_volume_v2":             3,
		"huaweicloudstack_compute_floatingip_associate_v2":    4,
		"huaweicloudstack_compute_instance_v2":                5,
		"huaweicloudstack_compute_interface_attach_v2":        6,
		"huaweicloudstack_compute_keypair_v2":                 7,
		"huaweicloudstack_compute_servergroup_v2":             8,
		"huaweicloudstack_compute_volume_attach_v2":           9,
		"huaweicloudstack_lb_certificate_v2":                  10,
		"huaweicloudstack_lb_l7policy_v2":                     11,
		"huaweicloudstack_lb_l7rule_v2":                       12,
		"huaweicloudstack_lb_listener_v2":                     13,
		"huaweicloudstack_lb_loadbalancer_v2":                 14,
		"huaweicloudstack_lb_member_v2":                       15,
		"huaweicloudstack_lb_monitor_v2":                      16,
		"huaweicloudstack_lb_pool_v2":                         17,
		"huaweicloudstack_lb_whitelist_v2":                    18,
		"huaweicloudstack_networking_floatingip_associate_v2": 19,
		"huaweicloudstack_networking_floatingip_v2":           20,
		"huaweicloudstack_networking_network_v2":              21,
		"huaweicloudstack_networking_port_v2":                 22,
		"huaweicloudstack_networking_router_interface_v2":     23,
		"huaweicloudstack_networking_router_route_v2":         24,
		"huaweicloudstack_networking_router_v2":               25,
		"huaweicloudstack_networking_secgroup_rule_v2":        26,
		"huaweicloudstack_networking_secgroup_v2":             27,
		"huaweicloudstack_networking_subnet_v2":               28,
		"huaweicloudstack_networking_vip_associate_v2":        29,
		"huaweicloudstack_networking_vip_v2":                  30,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
