package aws

import (
	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*Resource{

		&resource.Resource{
			Name:             "",
			Type:             "opentelekomcloud_antiddos_v1",
			Category:         "Antiddos Resources",
			ShortDescription: `Anti-DDoS Defends resources on the public cloud against network and monitors the service traffic from the Internet to ECSs, ELB instances, and BMSs to detect attack traffic in real time.`,
			Description:      ``,
			Keywords: []string{
				"antiddos",
				"v1",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "enable_l7",
					Description: `(Required) Specifies whether to enable L7 defense.`,
				},
				resource.Attribute{
					Name:        "traffic_pos_id",
					Description: `(Required) The position ID of traffic. The value ranges from 1 to 9.`,
				},
				resource.Attribute{
					Name:        "http_request_pos_id",
					Description: `(Required) The position ID of number of HTTP requests. The value ranges from 1 to 15.`,
				},
				resource.Attribute{
					Name:        "cleaning_access_pos_id",
					Description: `(Required)The position ID of access limit during cleaning. The value ranges from 1 to 8.`,
				},
				resource.Attribute{
					Name:        "app_type_id",
					Description: `(Required) The application type ID.`,
				},
				resource.Attribute{
					Name:        "floating_ip_id",
					Description: `(Required) The ID corresponding to the Elastic IP Address (EIP) of a user. ## Attributes Reference All above argument parameters can be exported as attribute parameters. ## Import Antiddos can be imported using the floating_ip_id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import opentelekomcloud_antiddos_v1.myantiddos c1881895-cdcb-4d23-96cb-032e6a3ee667 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opentelekomcloud_as_configuration_v1",
			Category:         "Auto Scaling Resources",
			ShortDescription: `Manages a V1 AS Configuration resource within OpenTelekomCloud.`,
			Description:      ``,
			Keywords: []string{
				"auto",
				"scaling",
				"as",
				"configuration",
				"v1",
			},
			Arguments: []resource.Argument{
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
					Description: `(Required) The disk type, which must be the same as the disk type available in the system. The options include ` + "`" + `SATA` + "`" + ` (common I/O disk type) and ` + "`" + `SSD` + "`" + ` (ultra-high I/O disk type).`,
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
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opentelekomcloud_as_group_v1",
			Category:         "Auto Scaling Resources",
			ShortDescription: `Manages a V1 Autoscaling Group resource within OpenTelekomCloud.`,
			Description:      ``,
			Keywords: []string{
				"auto",
				"scaling",
				"as",
				"group",
				"v1",
			},
			Arguments: []resource.Argument{
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
					Description: `(Optional) The ELB listener IDs. The system supports up to three ELB listeners, the IDs of which are separated using a comma (,).`,
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
					Description: `(Required) The UUID of the security group. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "scaling_group_name",
					Description: `See Argument Reference above.`,
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
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "scaling_group_name",
					Description: `See Argument Reference above.`,
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
			Type:             "opentelekomcloud_as_policy_v1",
			Category:         "Auto Scaling Resources",
			ShortDescription: `Manages a V1 AS Policy resource within OpenTelekomCloud.`,
			Description:      ``,
			Keywords: []string{
				"auto",
				"scaling",
				"as",
				"policy",
				"v1",
			},
			Arguments: []resource.Argument{
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
			Attributes: []resource.Argument{
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
			Type:             "opentelekomcloud_blockstorage_volume_v2",
			Category:         "Block Storage Resources",
			ShortDescription: `Manages a V2 volume resource within OpenTelekomCloud.`,
			Description:      ``,
			Keywords: []string{
				"block",
				"storage",
				"blockstorage",
				"volume",
				"v2",
			},
			Arguments: []resource.Argument{
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
					Description: `(Optional) Metadata key/value pairs to associate with the volume. Changing this updates the existing volume metadata. The EVS encryption capability with KMS key can be set with the following parameters:`,
				},
				resource.Attribute{
					Name:        "__system__encrypted",
					Description: `The default value is set to '0', which means the volume is not encrypted, the value '1' indicates volume is encrypted.`,
				},
				resource.Attribute{
					Name:        "__system__cmkid",
					Description: `(Optional) The ID of the kms key.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags key/value pairs to associate with the volume. Changing this updates the existing volume tags.`,
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
					Description: `(Optional) The type of volume to create. Changing this creates a new volume.`,
				},
				resource.Attribute{
					Name:        "cascade",
					Description: `(Optional, Default:false) Specifies to delete all snapshots associated with the EVS disk. ## Attributes Reference The following attributes are exported:`,
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
					Description: `If a volume is attached to an instance, this attribute will display the Attachment ID, Instance ID, and the Device as the Instance sees it. ## Import Volumes can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import opentelekomcloud_blockstorage_volume_v2.volume_1 ea257959-eeb1-4c10-8d33-26f0409a755d ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{
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
					Description: `If a volume is attached to an instance, this attribute will display the Attachment ID, Instance ID, and the Device as the Instance sees it. ## Import Volumes can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import opentelekomcloud_blockstorage_volume_v2.volume_1 ea257959-eeb1-4c10-8d33-26f0409a755d ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opentelekomcloud_cce_cluster_v3",
			Category:         "CCE Resources",
			ShortDescription: `Provides Cloud Container Engine(CCE) resource.`,
			Description:      ``,
			Keywords: []string{
				"cce",
				"cluster",
				"v3",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Cluster name. Changing this parameter will create a new cluster resource.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) Cluster tag, key/value pair format. Changing this parameter will create a new cluster resource.`,
				},
				resource.Attribute{
					Name:        "annotations",
					Description: `(Optional) Cluster annotation, key/value pair format. Changing this parameter will create a new cluster resource.`,
				},
				resource.Attribute{
					Name:        "flavor_id",
					Description: `(Required) Cluster specifications. Changing this parameter will create a new cluster resource.`,
				},
				resource.Attribute{
					Name:        "cce.s1.small",
					Description: `small-scale single cluster (up to 50 nodes).`,
				},
				resource.Attribute{
					Name:        "cce.s1.medium",
					Description: `medium-scale single cluster (up to 200 nodes).`,
				},
				resource.Attribute{
					Name:        "cce.s1.large",
					Description: `large-scale single cluster (up to 1000 nodes).`,
				},
				resource.Attribute{
					Name:        "cce.s2.small",
					Description: `small-scale HA cluster (up to 50 nodes).`,
				},
				resource.Attribute{
					Name:        "cce.s2.medium",
					Description: `medium-scale HA cluster (up to 200 nodes).`,
				},
				resource.Attribute{
					Name:        "cce.s2.large",
					Description: `large-scale HA cluster (up to 1000 nodes).`,
				},
				resource.Attribute{
					Name:        "cce.t1.small",
					Description: `small-scale single physical machine cluster (up to 10 nodes).`,
				},
				resource.Attribute{
					Name:        "cce.t1.medium",
					Description: `medium-scale single physical machine cluster (up to 100 nodes).`,
				},
				resource.Attribute{
					Name:        "cce.t1.large",
					Description: `large-scale single physical machine cluster (up to 500 nodes).`,
				},
				resource.Attribute{
					Name:        "cce.t2.small",
					Description: `small-scale HA physical machine cluster (up to 10 nodes).`,
				},
				resource.Attribute{
					Name:        "cce.t2.medium",
					Description: `medium-scale HA physical machine cluster (up to 100 nodes).`,
				},
				resource.Attribute{
					Name:        "cce.t2.large",
					Description: `large-scale HA physical machine cluster (up to 500 nodes).`,
				},
				resource.Attribute{
					Name:        "cluster_version",
					Description: `(Optional) For the cluster version, possible values are v1.9.2-r2 or v1.11.3-r1. Changing this parameter will create a new cluster resource. [OTC-API](https://docs.otc.t-systems.com/en-us/api2/cce/cce_02_0236.html)`,
				},
				resource.Attribute{
					Name:        "cluster_type",
					Description: `(Required) Cluster Type, possible values are VirtualMachine and BareMetal. Changing this parameter will create a new cluster resource.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Cluster description.`,
				},
				resource.Attribute{
					Name:        "billing_mode",
					Description: `(Optional) Charging mode of the cluster, which is 0 (on demand). Changing this parameter will create a new cluster resource.`,
				},
				resource.Attribute{
					Name:        "extend_param",
					Description: `(Optional) Extended parameter. Changing this parameter will create a new cluster resource.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required) The ID of the VPC used to create the node. Changing this parameter will create a new cluster resource.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Required) The ID of the subnet used to create the node. Changing this parameter will create a new cluster resource.`,
				},
				resource.Attribute{
					Name:        "highway_subnet_id",
					Description: `(Optional) The ID of the high speed network used to create bare metal nodes. Changing this parameter will create a new cluster resource.`,
				},
				resource.Attribute{
					Name:        "container_network_type",
					Description: `(Required) Container network type.`,
				},
				resource.Attribute{
					Name:        "overlay_l2",
					Description: `An overlay_l2 network built for containers by using Open vSwitch(OVS)`,
				},
				resource.Attribute{
					Name:        "underlay_ipvlan",
					Description: `An underlay_ipvlan network built for bare metal servers by using ipvlan.`,
				},
				resource.Attribute{
					Name:        "vpc-router",
					Description: `An vpc-router network built for containers by using ipvlan and custom VPC routes.`,
				},
				resource.Attribute{
					Name:        "container_network_cidr",
					Description: `(Optional) Container network segment. Changing this parameter will create a new cluster resource.`,
				},
				resource.Attribute{
					Name:        "authentication_mode",
					Description: `(Optional) Authentication mode of the cluster, possible values are x509 and rbac. Defaults to x509. Changing this parameter will create a new cluster resource. ## Attributes Reference All above argument parameters can be exported as attribute parameters along with attribute reference.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Id of the cluster resource.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Cluster status information.`,
				},
				resource.Attribute{
					Name:        "internal",
					Description: `The internal network address.`,
				},
				resource.Attribute{
					Name:        "external",
					Description: `The external network address.`,
				},
				resource.Attribute{
					Name:        "external_otc",
					Description: `The endpoint of the cluster to be accessed through API Gateway. ## Import Cluster can be imported using the cluster id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import opentelekomcloud_cce_cluster_v3.cluster_1 4779ab1c-7c1a-44b1-a02e-93dfc361b32d ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "id",
					Description: `Id of the cluster resource.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Cluster status information.`,
				},
				resource.Attribute{
					Name:        "internal",
					Description: `The internal network address.`,
				},
				resource.Attribute{
					Name:        "external",
					Description: `The external network address.`,
				},
				resource.Attribute{
					Name:        "external_otc",
					Description: `The endpoint of the cluster to be accessed through API Gateway. ## Import Cluster can be imported using the cluster id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import opentelekomcloud_cce_cluster_v3.cluster_1 4779ab1c-7c1a-44b1-a02e-93dfc361b32d ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opentelekomcloud_cce_node_v3",
			Category:         "CCE Resources",
			ShortDescription: `Add a node to a container cluster.`,
			Description:      ``,
			Keywords: []string{
				"cce",
				"node",
				"v3",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Required) ID of the cluster. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "billing_mode",
					Description: `(Optional) Node's billing mode: The value is 0 (on demand). Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Node Name.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) Node tag, key/value pair format. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "annotations",
					Description: `(Optional) Node annotation, key/value pair format. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "flavor_id",
					Description: `(Required) Specifies the flavor id. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `(Required) specify the name of the available partition (AZ). Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "key_pair",
					Description: `(Required) Key pair name when logging in to select the key pair mode. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "eip_ids",
					Description: `(Optional) List of existing elastic IP IDs. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "eip_count",
					Description: `(Optional) Number of elastic IPs to be dynamically created. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "iptype",
					Description: `(Required) Elastic IP type.`,
				},
				resource.Attribute{
					Name:        "bandwidth_charge_mode",
					Description: `(Optional) Bandwidth billing type. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "sharetype",
					Description: `(Required) Bandwidth sharing type. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "bandwidth_size",
					Description: `(Required) Bandwidth size. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "extend_param_charging_mode",
					Description: `(Optional) Node charging mode, 0 is on-demand charging. Changing this parameter will create a new cluster resource.`,
				},
				resource.Attribute{
					Name:        "ecs_performance_type",
					Description: `(Optional) Classification of cloud server specifications. Changing this parameter will create a new cluster resource.`,
				},
				resource.Attribute{
					Name:        "order_id",
					Description: `(Optional) Order ID, mandatory when the node payment type is the automatic payment package period type. Changing this parameter will create a new cluster resource.`,
				},
				resource.Attribute{
					Name:        "product_id",
					Description: `(Optional) The Product ID. Changing this parameter will create a new cluster resource.`,
				},
				resource.Attribute{
					Name:        "max_pods",
					Description: `(Optional) The maximum number of instances a node is allowed to create. Changing this parameter will create a new cluster resource.`,
				},
				resource.Attribute{
					Name:        "public_key",
					Description: `(Optional) The Public key. Changing this parameter will create a new cluster resource.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `(Required) Disk size in GB.`,
				},
				resource.Attribute{
					Name:        "volumetype",
					Description: `(Required) Disk type.`,
				},
				resource.Attribute{
					Name:        "extend_param",
					Description: `(Optional) Disk expansion parameters.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `(Required) Disk size in GB.`,
				},
				resource.Attribute{
					Name:        "volumetype",
					Description: `(Required) Disk type.`,
				},
				resource.Attribute{
					Name:        "extend_param",
					Description: `(Optional) Disk expansion parameters. ## Attributes Reference All above argument parameters can be exported as attribute parameters along with attribute reference.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Node status information.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `Private IP of the CCE node.`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "status",
					Description: `Node status information.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `Private IP of the CCE node.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opentelekomcloud_ces_alarmrule",
			Category:         "Ces Alarm",
			ShortDescription: `Manages a V2 topic resource within OpenTelekomCloud.`,
			Description:      ``,
			Keywords: []string{
				"ces",
				"alarm",
				"alarmrule",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "alarm_name",
					Description: `(Required) Specifies the name of an alarm rule. The value can be a string of 1 to 128 characters that can consist of numbers, lowercase letters, uppercase letters, underscores (_), or hyphens (-).`,
				},
				resource.Attribute{
					Name:        "alarm_description",
					Description: `(Optional) The value can be a string of 0 to 256 characters.`,
				},
				resource.Attribute{
					Name:        "metric",
					Description: `(Required) Specifies the alarm metrics. The structure is described below.`,
				},
				resource.Attribute{
					Name:        "condition",
					Description: `(Required) Specifies the alarm triggering condition. The structure is described below.`,
				},
				resource.Attribute{
					Name:        "alarm_actions",
					Description: `(Optional) Specifies the actions list triggered by an alarm. The structure is described below.`,
				},
				resource.Attribute{
					Name:        "insufficientdata_actions",
					Description: `(Optional) Specifies the actions list triggered by data insufficiency. The structure is described below.`,
				},
				resource.Attribute{
					Name:        "ok_actions",
					Description: `(Optional) Specifies the actions list triggered by the clearing of an alarm. The structure is described below.`,
				},
				resource.Attribute{
					Name:        "alarm_enabled",
					Description: `(Optional) Specifies whether to enable the alarm. The default value is true.`,
				},
				resource.Attribute{
					Name:        "alarm_action_enabled",
					Description: `(Optional) Specifies whether to enable the action to be triggered by an alarm. The default value is true. Note: If alarm_action_enabled is set to true, at least one of the following parameters alarm_actions, insufficientdata_actions, and ok_actions cannot be empty. If alarm_actions, insufficientdata_actions, and ok_actions coexist, their corresponding notification_list must be of the same value. The ` + "`" + `metric` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "namespace",
					Description: `(Required) Specifies the namespace in service.item format. service.item can be a string of 3 to 32 characters that must start with a letter and can consists of uppercase letters, lowercase letters, numbers, or underscores (_).`,
				},
				resource.Attribute{
					Name:        "metric_name",
					Description: `(Required) Specifies the metric name. The value can be a string of 1 to 64 characters that must start with a letter and can consists of uppercase letters, lowercase letters, numbers, or underscores (_).`,
				},
				resource.Attribute{
					Name:        "dimensions",
					Description: `(Required) Specifies the list of metric dimensions. Currently, the maximum length of the dimesion list that are supported is 3. The structure is described below. The ` + "`" + `dimensions` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Specifies the dimension name. The value can be a string of 1 to 32 characters that must start with a letter and can consists of uppercase letters, lowercase letters, numbers, underscores (_), or hyphens (-).`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) Specifies the dimension value. The value can be a string of 1 to 64 characters that must start with a letter or a number and can consists of uppercase letters, lowercase letters, numbers, underscores (_), or hyphens (-). The ` + "`" + `condition` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "period",
					Description: `(Required) Specifies the alarm checking period in seconds. The value can be 1, 300, 1200, 3600, 14400, and 86400. Note: If period is set to 1, the raw metric data is used to determine whether to generate an alarm.`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Required) Specifies the data rollup methods. The value can be max, min, average, sum, and vaiance.`,
				},
				resource.Attribute{
					Name:        "comparison_operator",
					Description: `(Required) Specifies the comparison condition of alarm thresholds. The value can be >, =, <, >=, or <=.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) Specifies the alarm threshold. The value ranges from 0 to Number of 1.7976931348623157e+308.`,
				},
				resource.Attribute{
					Name:        "unit",
					Description: `(Optional) Specifies the data unit.`,
				},
				resource.Attribute{
					Name:        "count",
					Description: `(Required) Specifies the number of consecutive occurrence times. The value ranges from 1 to 5. the ` + "`" + `alarm_actions` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) specifies the type of action triggered by an alarm. the value can be notification or autoscaling. notification: indicates that a notification will be sent to the user. autoscaling: indicates that a scaling action will be triggered.`,
				},
				resource.Attribute{
					Name:        "notification_list",
					Description: `(Required) specifies the topic urn list of the target notification objects. the maximum length is 5. the topic urn list can be obtained from simple message notification (smn) and in the following format: urn: smn:([a-z]|[a-z]|[0-9]|\-){1,32}:([a-z]|[a-z]|[0-9]){32}:([a-z]|[a-z]|[0-9]|\-|\_){1,256}. if type is set to notification, the value of notification_list cannot be empty. if type is set to autoscaling, the value of notification_list must be [] and the value of namespace must be sys.as. Note: to enable the as alarm rules take effect, you must bind scaling policies. for details, see the auto scaling api reference. the ` + "`" + `insufficientdata_actions` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) specifies the type of action triggered by an alarm. the value is notification. notification: indicates that a notification will be sent to the user.`,
				},
				resource.Attribute{
					Name:        "notification_list",
					Description: `(Optional) indicates the list of objects to be notified if the alarm status changes. the maximum length is 5. the ` + "`" + `ok_actions` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) specifies the type of action triggered by an alarm. the value is notification. notification: indicates that a notification will be sent to the user.`,
				},
				resource.Attribute{
					Name:        "notification_list",
					Description: `(Optional) indicates the list of objects to be notified if the alarm status changes. the maximum length is 5. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "alarm_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "alarm_description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "metric",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "condition",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "alarm_actions",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "insufficientdata_actions",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "ok_actions",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "alarm_enabled",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "alarm_action_enabled",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Specifies the alarm rule ID.`,
				},
				resource.Attribute{
					Name:        "update_time",
					Description: `Specifies the time when the alarm status changed. The value is a UNIX timestamp and the unit is ms.`,
				},
				resource.Attribute{
					Name:        "alarm_state",
					Description: `Specifies the alarm status. The value can be: ok: The alarm status is normal, alarm: An alarm is generated, insufficient_data: The required data is insufficient.`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "alarm_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "alarm_description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "metric",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "condition",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "alarm_actions",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "insufficientdata_actions",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "ok_actions",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "alarm_enabled",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "alarm_action_enabled",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Specifies the alarm rule ID.`,
				},
				resource.Attribute{
					Name:        "update_time",
					Description: `Specifies the time when the alarm status changed. The value is a UNIX timestamp and the unit is ms.`,
				},
				resource.Attribute{
					Name:        "alarm_state",
					Description: `Specifies the alarm status. The value can be: ok: The alarm status is normal, alarm: An alarm is generated, insufficient_data: The required data is insufficient.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opentelekomcloud_compute_server_v2",
			Category:         "Compute Resources",
			ShortDescription: `Manages a BMS server resource within OpenTelekomCloud.`,
			Description:      ``,
			Keywords: []string{
				"compute",
				"server",
				"v2",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the BMS.`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `(Optional; Required if ` + "`" + `image_name` + "`" + ` is empty.) Changing this creates a new bms server.`,
				},
				resource.Attribute{
					Name:        "image_name",
					Description: `(Optional; Required if ` + "`" + `image_id` + "`" + ` is empty.) The name of the desired image for the bms server. Changing this creates a new bms server.`,
				},
				resource.Attribute{
					Name:        "flavor_id",
					Description: `(Optional; Required if ` + "`" + `flavor_name` + "`" + ` is empty) The flavor ID of the desired flavor for the bms server. Changing this resizes the existing bms server.`,
				},
				resource.Attribute{
					Name:        "flavor_name",
					Description: `(Optional; Required if ` + "`" + `flavor_id` + "`" + ` is empty) The name of the desired flavor for the bms server. Changing this resizes the existing bms server.`,
				},
				resource.Attribute{
					Name:        "user_data",
					Description: `(Optional) The user data to provide when launching the instance. Changing this creates a new bms server.`,
				},
				resource.Attribute{
					Name:        "security_groups",
					Description: `(Optional) An array of one or more security group names to associate with the bms server. Changing this results in adding/removing security groups from the existing bms server.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `(Required) The availability zone in which to create the bms server.`,
				},
				resource.Attribute{
					Name:        "network",
					Description: `(Optional) An array of one or more networks to attach to the bms instance. Changing this creates a new bms server.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `(Optional) Metadata key/value pairs to make available from within the instance. Changing this updates the existing bms server metadata.`,
				},
				resource.Attribute{
					Name:        "admin_pass",
					Description: `(Optional) The administrative password to assign to the bms server. Changing this changes the root password on the existing server.`,
				},
				resource.Attribute{
					Name:        "key_pair",
					Description: `(Optional) The name of a key pair to put on the bms server. The key pair must already be created and associated with the tenant's account. Changing this creates a new bms server.`,
				},
				resource.Attribute{
					Name:        "stop_before_destroy",
					Description: `(Optional) Whether to try stop instance gracefully before destroying it, thus giving chance for guest OS daemons to stop correctly. If instance doesn't stop within timeout, it will be destroyed anyway. The ` + "`" + `network` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `(Required unless ` + "`" + `port` + "`" + ` or ` + "`" + `name` + "`" + ` is provided) The network UUID to attach to the bms server. Changing this creates a new bms server.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required unless ` + "`" + `uuid` + "`" + ` or ` + "`" + `port` + "`" + ` is provided) The human-readable name of the network. Changing this creates a new bms server.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Required unless ` + "`" + `uuid` + "`" + ` or ` + "`" + `name` + "`" + ` is provided) The port UUID of a network to attach to the bms server. Changing this creates a new server.`,
				},
				resource.Attribute{
					Name:        "fixed_ip_v4",
					Description: `(Optional) Specifies a fixed IPv4 address to be used on this network. Changing this creates a new bms server.`,
				},
				resource.Attribute{
					Name:        "fixed_ip_v6",
					Description: `(Optional) Specifies a fixed IPv6 address to be used on this network. Changing this creates a new bms server.`,
				},
				resource.Attribute{
					Name:        "access_network",
					Description: `(Optional) Specifies if this network should be used for provisioning access. Accepts true or false. Defaults to false. # Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of the bms server.`,
				},
				resource.Attribute{
					Name:        "config_drive",
					Description: `Whether to use the config_drive feature to configure the instance.`,
				},
				resource.Attribute{
					Name:        "kernel_id",
					Description: `The UUID of the kernel image when the AMI image is used.`,
				},
				resource.Attribute{
					Name:        "user_id",
					Description: `The ID of the user to which the BMS belongs.`,
				},
				resource.Attribute{
					Name:        "host_status",
					Description: `The nova-compute status:`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opentelekomcloud_compute_floatingip_associate_v2",
			Category:         "Compute Resources",
			ShortDescription: `Associate a floating IP to an instance`,
			Description:      ``,
			Keywords: []string{
				"compute",
				"floatingip",
				"associate",
				"v2",
			},
			Arguments: []resource.Argument{
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
					Name:        "floating_ip",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "fixed_ip",
					Description: `See Argument Reference above. ## Import This resource can be imported by specifying all three arguments, separated by a forward slash: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import opentelekomcloud_compute_floatingip_associate_v2.fip_1 <floating_ip>/<instance_id>/<fixed_ip> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{
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
					Description: `See Argument Reference above. ## Import This resource can be imported by specifying all three arguments, separated by a forward slash: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import opentelekomcloud_compute_floatingip_associate_v2.fip_1 <floating_ip>/<instance_id>/<fixed_ip> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opentelekomcloud_compute_floatingip_v2",
			Category:         "Compute Resources",
			ShortDescription: `Manages a V2 floating IP resource within OpenTelekomCloud Nova (compute).`,
			Description:      ``,
			Keywords: []string{
				"compute",
				"floatingip",
				"v2",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "pool",
					Description: `(Required) The name of the pool from which to obtain the floating IP. Changing this creates a new floating IP. ## Attributes Reference The following attributes are exported:`,
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
					Name:        "fixed_ip",
					Description: `The fixed IP address corresponding to the floating IP.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `UUID of the compute instance associated with the floating IP. ## Import Floating IPs can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import opentelekomcloud_compute_floatingip_v2.floatip_1 89c60255-9bd6-460c-822a-e2b959ede9d2 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "pool",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `The actual floating IP address itself.`,
				},
				resource.Attribute{
					Name:        "fixed_ip",
					Description: `The fixed IP address corresponding to the floating IP.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `UUID of the compute instance associated with the floating IP. ## Import Floating IPs can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import opentelekomcloud_compute_floatingip_v2.floatip_1 89c60255-9bd6-460c-822a-e2b959ede9d2 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opentelekomcloud_compute_instance_v2",
			Category:         "Compute Resources",
			ShortDescription: `Manages a V2 VM instance resource within OpenTelekomCloud.`,
			Description:      ``,
			Keywords: []string{
				"compute",
				"instance",
				"v2",
			},
			Arguments: []resource.Argument{
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
					Description: `(Optional) An array of one or more networks to attach to the instance. Required when there are multiple networks defined for the tenant. The network object structure is documented below. Changing this creates a new server.`,
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
					Name:        "personality",
					Description: `(Optional) Customize the personality of an instance by defining one or more files and their contents. The personality structure is described below.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `(Optional) Tags key/value pairs to associate with the instance.`,
				},
				resource.Attribute{
					Name:        "stop_before_destroy",
					Description: `(Optional) Whether to try stop instance gracefully before destroying it, thus giving chance for guest OS daemons to stop correctly. If instance doesn't stop within timeout, it will be destroyed anyway.`,
				},
				resource.Attribute{
					Name:        "force_delete",
					Description: `(Optional) Whether to force the OpenTelekomCloud instance to be forcefully deleted. This is useful for environments that have reclaim / soft deletion enabled.`,
				},
				resource.Attribute{
					Name:        "auto_recovery",
					Description: `(Optional) Configures or deletes automatic recovery of an instance The ` + "`" + `network` + "`" + ` block supports:`,
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
					Description: `The size of the volume to create (in gigabytes). Required in the following combinations: source=image and destination=volume, and source=blank and destination=volume. Changing this creates a new server.`,
				},
				resource.Attribute{
					Name:        "volume_type",
					Description: `(Optional) Currently, the value can be ` + "`" + `SSD` + "`" + ` (ultra-I/O disk type), ` + "`" + `SAS` + "`" + ` (high I/O disk type), or ` + "`" + `SATA` + "`" + ` (common I/O disk type) [OTC-API](https://docs.otc.t-systems.com/en-us/api/ecs/en-us_topic_0065817708.html)`,
				},
				resource.Attribute{
					Name:        "boot_index",
					Description: `(Optional) The boot index of the volume. It defaults to 0. Changing this creates a new server.`,
				},
				resource.Attribute{
					Name:        "destination_type",
					Description: `(Optional) The type that gets created. Currently only support "volume". Changing this creates a new server.`,
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
					Description: `(Optional) An IP Address in CIDR form. The instance will be placed on a compute node that is in the same subnet.`,
				},
				resource.Attribute{
					Name:        "tenancy",
					Description: `(Optional) The tenancy specifies whether the ECS is to be created on a Dedicated Host (DeH) or in a shared pool.`,
				},
				resource.Attribute{
					Name:        "deh_id",
					Description: `(Optional) The ID of DeH. This parameter takes effect only when the value of tenancy is dedicated. The ` + "`" + `personality` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "file",
					Description: `(Required) The absolute path of the destination file.`,
				},
				resource.Attribute{
					Name:        "contents",
					Description: `(Required) The contents of the file. Limited to 255 bytes. ## Attributes Reference The following attributes are exported:`,
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
					Description: `Contains all instance metadata, even metadata not set by Terraform.`,
				},
				resource.Attribute{
					Name:        "auto_recovery",
					Description: `See Argument Reference above. ## Notes ### Multiple Ephemeral Disks It's possible to specify multiple ` + "`" + `block_device` + "`" + ` entries to create an instance with multiple ephemeral (local) disks. In order to create multiple ephemeral disks, the sum of the total amount of ephemeral space must be less than or equal to what the chosen flavor supports. The following example shows how to create an instance with multiple ephemeral disks: ` + "`" + `` + "`" + `` + "`" + ` resource "opentelekomcloud_compute_instance_v2" "foo" { name = "terraform-test" security_groups = ["default"] block_device { boot_index = 0 delete_on_termination = true destination_type = "volume" source_type = "image" uuid = "<image uuid>" } block_device { boot_index = -1 delete_on_termination = true destination_type = "volume" source_type = "blank" volume_size = 1 } block_device { boot_index = -1 delete_on_termination = true destination_type = "volume" source_type = "blank" volume_size = 1 } } ` + "`" + `` + "`" + `` + "`" + ` ### Instances and Ports Neutron Ports are a great feature and provide a lot of functionality. However, there are some notes to be aware of when mixing Instances and Ports:`,
				},
			},
			Attributes: []resource.Argument{
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
					Description: `Contains all instance metadata, even metadata not set by Terraform.`,
				},
				resource.Attribute{
					Name:        "auto_recovery",
					Description: `See Argument Reference above. ## Notes ### Multiple Ephemeral Disks It's possible to specify multiple ` + "`" + `block_device` + "`" + ` entries to create an instance with multiple ephemeral (local) disks. In order to create multiple ephemeral disks, the sum of the total amount of ephemeral space must be less than or equal to what the chosen flavor supports. The following example shows how to create an instance with multiple ephemeral disks: ` + "`" + `` + "`" + `` + "`" + ` resource "opentelekomcloud_compute_instance_v2" "foo" { name = "terraform-test" security_groups = ["default"] block_device { boot_index = 0 delete_on_termination = true destination_type = "volume" source_type = "image" uuid = "<image uuid>" } block_device { boot_index = -1 delete_on_termination = true destination_type = "volume" source_type = "blank" volume_size = 1 } block_device { boot_index = -1 delete_on_termination = true destination_type = "volume" source_type = "blank" volume_size = 1 } } ` + "`" + `` + "`" + `` + "`" + ` ### Instances and Ports Neutron Ports are a great feature and provide a lot of functionality. However, there are some notes to be aware of when mixing Instances and Ports:`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opentelekomcloud_compute_keypair_v2",
			Category:         "Compute Resources",
			ShortDescription: `Manages a V2 keypair resource within OpenTelekomCloud.`,
			Description:      ``,
			Keywords: []string{
				"compute",
				"keypair",
				"v2",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A unique name for the keypair. Changing this creates a new keypair.`,
				},
				resource.Attribute{
					Name:        "public_key",
					Description: `(Required) A pregenerated OpenSSH-formatted public key. Changing this creates a new keypair.`,
				},
				resource.Attribute{
					Name:        "value_specs",
					Description: `(Optional) Map of additional options. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "public_key",
					Description: `See Argument Reference above. ## Import Keypairs can be imported using the ` + "`" + `name` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import opentelekomcloud_compute_keypair_v2.my-keypair test-keypair ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "public_key",
					Description: `See Argument Reference above. ## Import Keypairs can be imported using the ` + "`" + `name` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import opentelekomcloud_compute_keypair_v2.my-keypair test-keypair ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opentelekomcloud_compute_secgroup_v2",
			Category:         "Compute Resources",
			ShortDescription: `Manages a V2 security group resource within OpenTelekomCloud.`,
			Description:      ``,
			Keywords: []string{
				"compute",
				"secgroup",
				"v2",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A unique name for the security group. Changing this updates the ` + "`" + `name` + "`" + ` of an existing security group.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Required) A description for the security group. Changing this updates the ` + "`" + `description` + "`" + ` of an existing security group.`,
				},
				resource.Attribute{
					Name:        "rule",
					Description: `(Optional) A rule describing how the security group operates. The rule object structure is documented below. Changing this updates the security group rules. As shown in the example above, multiple rule blocks may be used. The ` + "`" + `rule` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "from_port",
					Description: `(Required) An integer representing the lower bound of the port range to open. Changing this creates a new security group rule.`,
				},
				resource.Attribute{
					Name:        "to_port",
					Description: `(Required) An integer representing the upper bound of the port range to open. Changing this creates a new security group rule.`,
				},
				resource.Attribute{
					Name:        "ip_protocol",
					Description: `(Required) The protocol type that will be allowed. Changing this creates a new security group rule.`,
				},
				resource.Attribute{
					Name:        "cidr",
					Description: `(Optional) Required if ` + "`" + `from_group_id` + "`" + ` or ` + "`" + `self` + "`" + ` is empty. The IP range that will be the source of network traffic to the security group. Use 0.0.0.0/0 to allow all IP addresses. Changing this creates a new security group rule. Cannot be combined with ` + "`" + `from_group_id` + "`" + ` or ` + "`" + `self` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "from_group_id",
					Description: `(Optional) Required if ` + "`" + `cidr` + "`" + ` or ` + "`" + `self` + "`" + ` is empty. The ID of a group from which to forward traffic to the parent group. Changing this creates a new security group rule. Cannot be combined with ` + "`" + `cidr` + "`" + ` or ` + "`" + `self` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "self",
					Description: `(Optional) Required if ` + "`" + `cidr` + "`" + ` and ` + "`" + `from_group_id` + "`" + ` is empty. If true, the security group itself will be added as a source to this ingress rule. Cannot be combined with ` + "`" + `cidr` + "`" + ` or ` + "`" + `from_group_id` + "`" + `. ## Attributes Reference The following attributes are exported:`,
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
					Name:        "rule",
					Description: `See Argument Reference above. ## Notes ### ICMP Rules When using ICMP as the ` + "`" + `ip_protocol` + "`" + `, the ` + "`" + `from_port` + "`" + ` sets the ICMP _type_ and the ` + "`" + `to_port` + "`" + ` sets the ICMP _code_. To allow all ICMP types, set each value to ` + "`" + `-1` + "`" + `, like so: ` + "`" + `` + "`" + `` + "`" + `hcl rule { from_port = -1 to_port = -1 ip_protocol = "icmp" cidr = "0.0.0.0/0" } ` + "`" + `` + "`" + `` + "`" + ` A list of ICMP types and codes can be found [here](https://en.wikipedia.org/wiki/Internet_Control_Message_Protocol#Control_messages). ### Referencing Security Groups When referencing a security group in a configuration (for example, a configuration creates a new security group and then needs to apply it to an instance being created in the same configuration), it is currently recommended to reference the security group by name and not by ID, like this: ` + "`" + `` + "`" + `` + "`" + `hcl resource "opentelekomcloud_compute_instance_v2" "test-server" { name = "tf-test" image_id = "ad091b52-742f-469e-8f3c-fd81cadf0743" flavor_id = "3" key_pair = "my_key_pair_name" security_groups = ["${opentelekomcloud_compute_secgroup_v2.secgroup_1.name}"] } ` + "`" + `` + "`" + `` + "`" + ` ## Import Security Groups can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import opentelekomcloud_compute_secgroup_v2.my_secgroup 1bc30ee9-9d5b-4c30-bdd5-7f1e663f5edf ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "rule",
					Description: `See Argument Reference above. ## Notes ### ICMP Rules When using ICMP as the ` + "`" + `ip_protocol` + "`" + `, the ` + "`" + `from_port` + "`" + ` sets the ICMP _type_ and the ` + "`" + `to_port` + "`" + ` sets the ICMP _code_. To allow all ICMP types, set each value to ` + "`" + `-1` + "`" + `, like so: ` + "`" + `` + "`" + `` + "`" + `hcl rule { from_port = -1 to_port = -1 ip_protocol = "icmp" cidr = "0.0.0.0/0" } ` + "`" + `` + "`" + `` + "`" + ` A list of ICMP types and codes can be found [here](https://en.wikipedia.org/wiki/Internet_Control_Message_Protocol#Control_messages). ### Referencing Security Groups When referencing a security group in a configuration (for example, a configuration creates a new security group and then needs to apply it to an instance being created in the same configuration), it is currently recommended to reference the security group by name and not by ID, like this: ` + "`" + `` + "`" + `` + "`" + `hcl resource "opentelekomcloud_compute_instance_v2" "test-server" { name = "tf-test" image_id = "ad091b52-742f-469e-8f3c-fd81cadf0743" flavor_id = "3" key_pair = "my_key_pair_name" security_groups = ["${opentelekomcloud_compute_secgroup_v2.secgroup_1.name}"] } ` + "`" + `` + "`" + `` + "`" + ` ## Import Security Groups can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import opentelekomcloud_compute_secgroup_v2.my_secgroup 1bc30ee9-9d5b-4c30-bdd5-7f1e663f5edf ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opentelekomcloud_compute_servergroup_v2",
			Category:         "Compute Resources",
			ShortDescription: `Manages a V2 Server Group resource within OpenTelekomCloud.`,
			Description:      ``,
			Keywords: []string{
				"compute",
				"servergroup",
				"v2",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A unique name for the server group. Changing this creates a new server group.`,
				},
				resource.Attribute{
					Name:        "policies",
					Description: `(Required) The set of policies for the server group. Only two two policies are available right now, and both are mutually exclusive. See the Policies section for more information. Changing this creates a new server group.`,
				},
				resource.Attribute{
					Name:        "value_specs",
					Description: `(Optional) Map of additional options. ## Policies`,
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
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "policies",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "members",
					Description: `The instances that are part of this server group.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the server group. ## Import Server Groups can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import opentelekomcloud_compute_servergroup_v2.test-sg 1bc30ee9-9d5b-4c30-bdd5-7f1e663f5edf ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{
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
					Description: `The instances that are part of this server group.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the server group. ## Import Server Groups can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import opentelekomcloud_compute_servergroup_v2.test-sg 1bc30ee9-9d5b-4c30-bdd5-7f1e663f5edf ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opentelekomcloud_compute_volume_attach_v2",
			Category:         "Compute Resources",
			ShortDescription: `Attaches a Block Storage Volume to an Instance.`,
			Description:      ``,
			Keywords: []string{
				"compute",
				"volume",
				"attach",
				"v2",
			},
			Arguments: []resource.Argument{
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
					Name:        "instance_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "volume_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "device",
					Description: `See Argument Reference above. _NOTE_: The correctness of this information is dependent upon the hypervisor in use. In some cases, this should not be used as an authoritative piece of information. ## Import Volume Attachments can be imported using the Instance ID and Volume ID separated by a slash, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import opentelekomcloud_compute_volume_attach_v2.va_1 89c60255-9bd6-460c-822a-e2b959ede9d2/45670584-225f-46c3-b33e-6707b589b666 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{
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
					Description: `See Argument Reference above. _NOTE_: The correctness of this information is dependent upon the hypervisor in use. In some cases, this should not be used as an authoritative piece of information. ## Import Volume Attachments can be imported using the Instance ID and Volume ID separated by a slash, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import opentelekomcloud_compute_volume_attach_v2.va_1 89c60255-9bd6-460c-822a-e2b959ede9d2/45670584-225f-46c3-b33e-6707b589b666 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opentelekomcloud_csbs_backup_policy_v1",
			Category:         "CSBS Resources",
			ShortDescription: `Provides an OpenTelekomCloud Backup Policy of Resource.`,
			Description:      ``,
			Keywords: []string{
				"csbs",
				"backup",
				"policy",
				"v1",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Specifies the name of backup policy. The value consists of 1 to 255 characters and can contain only letters, digits, underscores (_), and hyphens (-).`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Backup policy description. The value consists of 0 to 255 characters and must not contain a greater-than sign (>) or less-than sign (<).`,
				},
				resource.Attribute{
					Name:        "provider_id",
					Description: `(Required) Specifies backup provider ID. Default value is`,
				},
				resource.Attribute{
					Name:        "common",
					Description: `(Optional) General backup policy parameters, which are blank by default.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Specifies Scheduling period name.The value consists of 1 to 255 characters and can contain only letters, digits, underscores (_), and hyphens (-).`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Specifies Scheduling period description.The value consists of 0 to 255 characters and must not contain a greater-than sign (>) or less-than sign (<).`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) Specifies whether the scheduling period is enabled. Default value is`,
				},
				resource.Attribute{
					Name:        "max_backups",
					Description: `(Optional) Specifies maximum number of backups that can be automatically created for a backup object.`,
				},
				resource.Attribute{
					Name:        "retention_duration_days",
					Description: `(Optional) Specifies duration of retaining a backup, in days.`,
				},
				resource.Attribute{
					Name:        "permanent",
					Description: `(Optional) Specifies whether backups are permanently retained.`,
				},
				resource.Attribute{
					Name:        "trigger_pattern",
					Description: `(Required) Specifies Scheduling policy of the scheduler.`,
				},
				resource.Attribute{
					Name:        "operation_type",
					Description: `(Required) Specifies Operation type, which can be backup.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Required) Specifies the ID of the object to be backed up.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) Entity object type of the backup object. If the type is VMs, the value is`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Specifies backup object name.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) Tag key. It cannot be an empty string.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) Tag value. It can be an empty string. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of Backup Policy.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Backup Policy ID.`,
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
					Description: `Specifies Scheduler type. ## Import Backup Policy can be imported using ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import opentelekomcloud_csbs_backup_policy_v1.backup_policy_v1 7056d636-ac60-4663-8a6c-82d3c32c1c64 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "status",
					Description: `Status of Backup Policy.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Backup Policy ID.`,
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
					Description: `Specifies Scheduler type. ## Import Backup Policy can be imported using ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import opentelekomcloud_csbs_backup_policy_v1.backup_policy_v1 7056d636-ac60-4663-8a6c-82d3c32c1c64 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opentelekomcloud_csbs_backup_v1",
			Category:         "CSBS Resources",
			ShortDescription: `Provides an OpenTelekomCloud Backup of Resources.`,
			Description:      ``,
			Keywords: []string{
				"csbs",
				"backup",
				"v1",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "backup_name",
					Description: `(Optional) Name for the backup. The value consists of 1 to 255 characters and can contain only letters, digits, underscores (_), and hyphens (-). Changing backup_name creates a new backup.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Backup description. The value consists of 0 to 255 characters and must not contain a greater-than sign (>) or less-than sign (<). Changing description creates a new backup.`,
				},
				resource.Attribute{
					Name:        "resource_id",
					Description: `(Required) ID of the target to which the backup is restored. Changing this creates a new backup.`,
				},
				resource.Attribute{
					Name:        "resource_type",
					Description: `(Optional) Type of the target to which the backup is restored. The default value is`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) block supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) Tag key. It cannot be an empty string.Changing key creates a new backup.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) Tag value. It can be an empty string.Changing value creates a new backup. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `It specifies the status of backup.`,
				},
				resource.Attribute{
					Name:        "backup_record_id",
					Description: `Specifies backup record ID.`,
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
					Description: `Specifies image type. ## Import Backup can be imported using ` + "`" + `backup_record_id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import opentelekomcloud_csbs_backup_v1.backup_v1.backup_v1 7056d636-ac60-4663-8a6c-82d3c32c1c64 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "status",
					Description: `It specifies the status of backup.`,
				},
				resource.Attribute{
					Name:        "backup_record_id",
					Description: `Specifies backup record ID.`,
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
					Description: `Specifies image type. ## Import Backup can be imported using ` + "`" + `backup_record_id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import opentelekomcloud_csbs_backup_v1.backup_v1.backup_v1 7056d636-ac60-4663-8a6c-82d3c32c1c64 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opentelekomcloud_css_cluster_v1",
			Category:         "Cloud Search Service Resources",
			ShortDescription: `cluster management`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"search",
				"service",
				"css",
				"cluster",
				"v1",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Cluster name. It contains 4 to 32 characters. Only letters, digits, hyphens (-), and underscores (_) are allowed. The value must start with a letter. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "node_config",
					Description: `(Required) Instance object. Structure is documented below. Changing this parameter will create a new resource. The ` + "`" + `node_config` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `(Optional) Availability zone (AZ). Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "flavor",
					Description: `(Required) Instance flavor name. Value range of flavor css.medium.8: 40 GB to 640 GB Value range of flavor css.large.8: 40 GB to 1280 GB Value range of flavor css.xlarge.8: 40 GB to 2560 GB Value range of flavor css.2xlarge.8: 80 GB to 5120 GB Value range of flavor css.4xlarge.8: 160 GB to 10240 GB Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "network_info",
					Description: `(Required) Network information. Structure is documented below. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "volume",
					Description: `(Required) Information about the volume. Structure is documented below. Changing this parameter will create a new resource. The ` + "`" + `network_info` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `(Required) Network ID. All instances in a cluster must have the same networks and security groups. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `(Required) Security group ID. All instances in a cluster must have the same subnets and security groups. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required) VPC ID, which is used for configuring cluster network. Changing this parameter will create a new resource. The ` + "`" + `volume` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "encryption_key",
					Description: `(Optional) Key ID. The Default Master Keys cannot be used to create grants. Specifically, you cannot use Default Master Keys whose aliases end with /default in KMS to create clusters. After a cluster is created, do not delete the key used by the cluster. Otherwise, the cluster will become unavailable. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `(Required) Volume size, which must be a multiple of 4 and 10. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "volume_type",
					Description: `(Required) COMMON: Common I/O. The SATA disk is used. HIGH: High I/O. The SAS disk is used. ULTRAHIGH: Ultra-high I/O. The solid-state drive (SSD) is used. Changing this parameter will create a new resource. - - -`,
				},
				resource.Attribute{
					Name:        "enable_https",
					Description: `(Optional) Whether communication encryption is performed on the cluster. Available values include true and false. By default, communication encryption is enabled. Value true indicates that communication encryption is performed on the cluster. Value false indicates that communication encryption is not performed on the cluster. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "expect_node_num",
					Description: `(Optional) Number of cluster instances. The value range is 1 to 32. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "created",
					Description: `Time when a cluster is created. The format is ISO8601: CCYY-MM-DDThh:mm:ss.`,
				},
				resource.Attribute{
					Name:        "datastore",
					Description: `Type of the data search engine. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "endpoint",
					Description: `Indicates the IP address and port number of the user used to access the VPC.`,
				},
				resource.Attribute{
					Name:        "nodes",
					Description: `List of node objects. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "updated",
					Description: `Last modification time of a cluster. The format is ISO8601: CCYY-MM-DDThh:mm:ss. The ` + "`" + `datastore` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Supported type: elasticsearch`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `Engine version number. The ` + "`" + `nodes` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Instance ID.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Instance name.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Supported type: ess (indicating the Elasticsearch node) ## Timeouts This resource provides the following timeouts configuration options: - ` + "`" + `create` + "`" + ` - Default is 15 minute. - ` + "`" + `update` + "`" + ` - Default is 30 minute.`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "created",
					Description: `Time when a cluster is created. The format is ISO8601: CCYY-MM-DDThh:mm:ss.`,
				},
				resource.Attribute{
					Name:        "datastore",
					Description: `Type of the data search engine. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "endpoint",
					Description: `Indicates the IP address and port number of the user used to access the VPC.`,
				},
				resource.Attribute{
					Name:        "nodes",
					Description: `List of node objects. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "updated",
					Description: `Last modification time of a cluster. The format is ISO8601: CCYY-MM-DDThh:mm:ss. The ` + "`" + `datastore` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Supported type: elasticsearch`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `Engine version number. The ` + "`" + `nodes` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Instance ID.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Instance name.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Supported type: ess (indicating the Elasticsearch node) ## Timeouts This resource provides the following timeouts configuration options: - ` + "`" + `create` + "`" + ` - Default is 15 minute. - ` + "`" + `update` + "`" + ` - Default is 30 minute.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opentelekomcloud_cts_tracker_v1",
			Category:         "CTS Resources",
			ShortDescription: `CTS tracker allows you to collect, store, and query cloud resource operation records and use these records for security analysis, compliance auditing, resource tracking, and fault locating.`,
			Description:      ``,
			Keywords: []string{
				"cts",
				"tracker",
				"v1",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "bucket_name",
					Description: `(Required) The OBS bucket name for a tracker.`,
				},
				resource.Attribute{
					Name:        "file_prefix_name",
					Description: `(Optional) The prefix of a log that needs to be stored in an OBS bucket.`,
				},
				resource.Attribute{
					Name:        "is_support_smn",
					Description: `(Required) Specifies whether SMN is supported. When the value is false, topic_id and operations can be left empty.`,
				},
				resource.Attribute{
					Name:        "topic_id",
					Description: `(Required)The theme of the SMN service, Is obtained from SMN and in the format of`,
				},
				resource.Attribute{
					Name:        "operations",
					Description: `(Required) Trigger conditions for sending a notification.`,
				},
				resource.Attribute{
					Name:        "is_send_all_key_operation",
					Description: `(Required) When the value is`,
				},
				resource.Attribute{
					Name:        "need_notify_user_list",
					Description: `(Optional) The users using the login function. When these users log in, notifications will be sent.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Optional) The status of a tracker. The value should be`,
				},
				resource.Attribute{
					Name:        "tracker_name",
					Description: `The tracker name. Currently, only tracker`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "tracker_name",
					Description: `The tracker name. Currently, only tracker`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opentelekomcloud_dcs_instance_v1",
			Category:         "DCS Resources",
			ShortDescription: `Manages a DCS instance in the opentelekomcloud DCS Service`,
			Description:      ``,
			Keywords: []string{
				"dcs",
				"instance",
				"v1",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Indicates the name of an instance. An instance name starts with a letter, consists of 4 to 64 characters, and supports only letters, digits, and hyphens (-).`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Indicates the description of an instance. It is a character string containing not more than 1024 characters.`,
				},
				resource.Attribute{
					Name:        "engine",
					Description: `(Required) Indicates a cache engine. Only Redis is supported. Changing this creates a new instance.`,
				},
				resource.Attribute{
					Name:        "engine_version",
					Description: `(Required) Indicates the version of a cache engine, which is 3.0.7. Changing this creates a new instance.`,
				},
				resource.Attribute{
					Name:        "capacity",
					Description: `(Required) Indicates the Cache capacity. Unit: GB. For a DCS Redis or Memcached instance in single-node or master/standby mode, the cache capacity can be 2 GB, 4 GB, 8 GB, 16 GB, 32 GB, or 64 GB. For a DCS Redis instance in cluster mode, the cache capacity can be 64, 128, 256, 512GB. Changing this creates a new instance.`,
				},
				resource.Attribute{
					Name:        "access_user",
					Description: `(Optional) Username used for accessing a DCS instance after password authentication. A username starts with a letter, consists of 1 to 64 characters, and supports only letters, digits, and hyphens (-). Changing this creates a new instance.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Required) Password of a DCS instance. The password of a DCS Redis instance must meet the following complexity requirements: Changing this creates a new instance.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required) Tenant's VPC ID. For details on how to create VPCs, see the Virtual Private Cloud API Reference. Changing this creates a new instance.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `(Required) Tenant's security group ID. For details on how to create security groups, see the Virtual Private Cloud API Reference.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Required) Subnet ID. For details on how to create subnets, see the Virtual Private Cloud API Reference. Changing this creates a new instance.`,
				},
				resource.Attribute{
					Name:        "available_zones",
					Description: `(Required) IDs of the AZs where cache nodes reside. For details on how to query AZs, see Querying AZ Information. Changing this creates a new instance.`,
				},
				resource.Attribute{
					Name:        "product_id",
					Description: `(Required) Product ID used to differentiate DCS instance types. Changing this creates a new instance.`,
				},
				resource.Attribute{
					Name:        "maintain_begin",
					Description: `(Optional) Indicates the time at which a maintenance time window starts. Format: HH:mm:ss. The start time and end time of a maintenance time window must indicate the time segment of a supported maintenance time window. For details, see section Querying Maintenance Time Windows. The start time must be set to 22:00, 02:00, 06:00, 10:00, 14:00, or 18:00. Parameters maintain_begin and maintain_end must be set in pairs. If parameter maintain_begin is left blank, parameter maintain_end is also blank. In this case, the system automatically allocates the default start time 02:00.`,
				},
				resource.Attribute{
					Name:        "maintain_end",
					Description: `(Optional) Indicates the time at which a maintenance time window ends. Format: HH:mm:ss. The start time and end time of a maintenance time window must indicate the time segment of a supported maintenance time window. For details, see section Querying Maintenance Time Windows. The end time is four hours later than the start time. For example, if the start time is 22:00, the end time is 02:00. Parameters maintain_begin and maintain_end must be set in pairs. If parameter maintain_end is left blank, parameter maintain_begin is also blank. In this case, the system automatically allocates the default end time 06:00.`,
				},
				resource.Attribute{
					Name:        "save_days",
					Description: `(Required) Retention time. Unit: day. Range: 1–7. Changing this creates a new instance.`,
				},
				resource.Attribute{
					Name:        "backup_type",
					Description: `(Required) Backup type. Options: auto: automatic backup. manual: manual backup. Changing this creates a new instance.`,
				},
				resource.Attribute{
					Name:        "begin_at",
					Description: `(Required) Time at which backup starts. "00:00-01:00" indicates that backup starts at 00:00:00. Changing this creates a new instance.`,
				},
				resource.Attribute{
					Name:        "period_type",
					Description: `(Required) Interval at which backup is performed. Currently, only weekly backup is supported. Changing this creates a new instance.`,
				},
				resource.Attribute{
					Name:        "backup_at",
					Description: `(Required) Day in a week on which backup starts. Range: 1–7. Where: 1 indicates Monday; 7 indicates Sunday. Changing this creates a new instance. ## Attributes Reference The following attributes are exported:`,
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
					Name:        "engine",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "engine_version",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "capacity",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "access_user",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "vpc_name",
					Description: `Indicates the name of a vpc.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "security_group_name",
					Description: `Indicates the name of a security group.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "subnet_name",
					Description: `Indicates the name of a subnet.`,
				},
				resource.Attribute{
					Name:        "available_zones",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "product_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "maintain_begin",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "maintain_end",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "save_days",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "backup_type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "begin_at",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "period_type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "backup_at",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "order_id",
					Description: `An order ID is generated only in the monthly or yearly billing mode. In other billing modes, no value is returned for this parameter.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Port of the cache node.`,
				},
				resource.Attribute{
					Name:        "resource_spec_code",
					Description: `Resource specifications. dcs.single_node: indicates a DCS instance in single-node mode. dcs.master_standby: indicates a DCS instance in master/standby mode. dcs.cluster: indicates a DCS instance in cluster mode.`,
				},
				resource.Attribute{
					Name:        "used_memory",
					Description: `Size of the used memory. Unit: MB.`,
				},
				resource.Attribute{
					Name:        "internal_version",
					Description: `Internal DCS version.`,
				},
				resource.Attribute{
					Name:        "max_memory",
					Description: `Overall memory size. Unit: MB.`,
				},
				resource.Attribute{
					Name:        "user_id",
					Description: `Indicates a user ID.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `Cache node's IP address in tenant's VPC.`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "engine",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "engine_version",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "capacity",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "access_user",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "vpc_name",
					Description: `Indicates the name of a vpc.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "security_group_name",
					Description: `Indicates the name of a security group.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "subnet_name",
					Description: `Indicates the name of a subnet.`,
				},
				resource.Attribute{
					Name:        "available_zones",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "product_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "maintain_begin",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "maintain_end",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "save_days",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "backup_type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "begin_at",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "period_type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "backup_at",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "order_id",
					Description: `An order ID is generated only in the monthly or yearly billing mode. In other billing modes, no value is returned for this parameter.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Port of the cache node.`,
				},
				resource.Attribute{
					Name:        "resource_spec_code",
					Description: `Resource specifications. dcs.single_node: indicates a DCS instance in single-node mode. dcs.master_standby: indicates a DCS instance in master/standby mode. dcs.cluster: indicates a DCS instance in cluster mode.`,
				},
				resource.Attribute{
					Name:        "used_memory",
					Description: `Size of the used memory. Unit: MB.`,
				},
				resource.Attribute{
					Name:        "internal_version",
					Description: `Internal DCS version.`,
				},
				resource.Attribute{
					Name:        "max_memory",
					Description: `Overall memory size. Unit: MB.`,
				},
				resource.Attribute{
					Name:        "user_id",
					Description: `Indicates a user ID.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `Cache node's IP address in tenant's VPC.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opentelekomcloud_deh_host_v1",
			Category:         "DeH Resources",
			ShortDescription: `Allocates a Dedicated Host to a tenant and set minimum required parameters for this Dedicated Host.`,
			Description: `

Allocates a Dedicated Host to a tenant and set minimum required parameters for this Dedicated Host.

`,
			Keywords: []string{
				"deh",
				"host",
				"v1",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "status",
					Description: `Specifies the Dedicated Host status.`,
				},
				resource.Attribute{
					Name:        "available_vcpus",
					Description: `The number of available vCPUs for the Dedicated Host.`,
				},
				resource.Attribute{
					Name:        "available_memory",
					Description: `The size of available memory for the Dedicated Host.`,
				},
				resource.Attribute{
					Name:        "instance_total",
					Description: `The number of the placed VMs.`,
				},
				resource.Attribute{
					Name:        "instance_uuids",
					Description: `The VMs started on the Dedicated Host.`,
				},
				resource.Attribute{
					Name:        "host_type_name",
					Description: `The name of the Dedicated Host type.`,
				},
				resource.Attribute{
					Name:        "vcpus",
					Description: `The number of host vCPUs.`,
				},
				resource.Attribute{
					Name:        "cores",
					Description: `The number of host physical cores.`,
				},
				resource.Attribute{
					Name:        "sockets",
					Description: `The number of host physical sockets.`,
				},
				resource.Attribute{
					Name:        "memory",
					Description: `The size of host physical memory (MB).`,
				},
				resource.Attribute{
					Name:        "available_instance_capacities",
					Description: `The VM flavors placed on the Dedicated Host. ## Import DeH can be imported using the ` + "`" + `dedicated_host_id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import opentelekomcloud_deh_host_v1.deh_host 4779ab1c-7c1a-44b1-a02e-93dfc361b32d ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "status",
					Description: `Specifies the Dedicated Host status.`,
				},
				resource.Attribute{
					Name:        "available_vcpus",
					Description: `The number of available vCPUs for the Dedicated Host.`,
				},
				resource.Attribute{
					Name:        "available_memory",
					Description: `The size of available memory for the Dedicated Host.`,
				},
				resource.Attribute{
					Name:        "instance_total",
					Description: `The number of the placed VMs.`,
				},
				resource.Attribute{
					Name:        "instance_uuids",
					Description: `The VMs started on the Dedicated Host.`,
				},
				resource.Attribute{
					Name:        "host_type_name",
					Description: `The name of the Dedicated Host type.`,
				},
				resource.Attribute{
					Name:        "vcpus",
					Description: `The number of host vCPUs.`,
				},
				resource.Attribute{
					Name:        "cores",
					Description: `The number of host physical cores.`,
				},
				resource.Attribute{
					Name:        "sockets",
					Description: `The number of host physical sockets.`,
				},
				resource.Attribute{
					Name:        "memory",
					Description: `The size of host physical memory (MB).`,
				},
				resource.Attribute{
					Name:        "available_instance_capacities",
					Description: `The VM flavors placed on the Dedicated Host. ## Import DeH can be imported using the ` + "`" + `dedicated_host_id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import opentelekomcloud_deh_host_v1.deh_host 4779ab1c-7c1a-44b1-a02e-93dfc361b32d ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opentelekomcloud_dms_group_v1",
			Category:         "DMS Resources",
			ShortDescription: `Manages a DMS group in the opentelekomcloud DMS Service`,
			Description:      ``,
			Keywords: []string{
				"dms",
				"group",
				"v1",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Indicates the unique name of a group. A string of 1 to 64 characters that contain a-z, A-Z, 0-9, hyphens (-), and underscores (_). The name cannot be modified once specified.`,
				},
				resource.Attribute{
					Name:        "queue_id",
					Description: `(Required) Indicates the ID of a specified queue. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "queue_id",
					Description: `Indicates the ID of a queue.`,
				},
				resource.Attribute{
					Name:        "redrive_policy",
					Description: `Indicates whether to enable dead letter messages.`,
				},
				resource.Attribute{
					Name:        "produced_messages",
					Description: `Indicates the total number of messages (not including the messages that have expired and been deleted) in a queue.`,
				},
				resource.Attribute{
					Name:        "consumed_messages",
					Description: `Indicates the total number of messages that are successfully consumed.`,
				},
				resource.Attribute{
					Name:        "available_messages",
					Description: `Indicates the accumulated number of messages that can be consumed.`,
				},
				resource.Attribute{
					Name:        "produced_deadletters",
					Description: `Indicates the total number of dead letter messages generated by the consumer group.`,
				},
				resource.Attribute{
					Name:        "available_deadletters",
					Description: `Indicates the accumulated number of dead letter messages that have not been consumed.`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "queue_id",
					Description: `Indicates the ID of a queue.`,
				},
				resource.Attribute{
					Name:        "redrive_policy",
					Description: `Indicates whether to enable dead letter messages.`,
				},
				resource.Attribute{
					Name:        "produced_messages",
					Description: `Indicates the total number of messages (not including the messages that have expired and been deleted) in a queue.`,
				},
				resource.Attribute{
					Name:        "consumed_messages",
					Description: `Indicates the total number of messages that are successfully consumed.`,
				},
				resource.Attribute{
					Name:        "available_messages",
					Description: `Indicates the accumulated number of messages that can be consumed.`,
				},
				resource.Attribute{
					Name:        "produced_deadletters",
					Description: `Indicates the total number of dead letter messages generated by the consumer group.`,
				},
				resource.Attribute{
					Name:        "available_deadletters",
					Description: `Indicates the accumulated number of dead letter messages that have not been consumed.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opentelekomcloud_dms_queue_v1",
			Category:         "DMS Resources",
			ShortDescription: `Manages a DMS queue in the opentelekomcloud DMS Service`,
			Description:      ``,
			Keywords: []string{
				"dms",
				"queue",
				"v1",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Indicates the unique name of a queue. A string of 1 to 64 characters that contain a-z, A-Z, 0-9, hyphens (-), and underscores (_). The name cannot be modified once specified.`,
				},
				resource.Attribute{
					Name:        "queue_mode",
					Description: `(Optional) Indicates the queue type. It only support 'NORMAL' and 'FIFO'. NORMAL: Standard queue. Best-effort ordering. Messages might be retrieved in an order different from which they were sent. Select standard queues when throughput is important. FIFO: First-ln-First-out (FIFO) queue. FIFO delivery. Messages are retrieved in the order they were sent. Select FIFO queues when the order of messages is important. Default value: NORMAL.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Indicates the basic information about a queue. The queue description must be 0 to 160 characters in length, and does not contain angle brackets (<) and (>).`,
				},
				resource.Attribute{
					Name:        "redrive_policy",
					Description: `(Optional) Indicates whether to enable dead letter messages. Dead letter messages indicate messages that cannot be normally consumed. The redrive_policy should be set to 'enable' or 'disable'. The default value is 'disable'.`,
				},
				resource.Attribute{
					Name:        "max_consume_count",
					Description: `(Optional) This parameter is mandatory only when redrive_policy is set to enable. This parameter indicates the maximum number of allowed message consumption failures. When a message fails to be consumed after the number of consumption attempts of this message reaches this value, DMS stores this message into the dead letter queue. The max_consume_count value range is 1–100. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "queue_mode",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "redrive_policy",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "max_consume_count",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "created",
					Description: `Indicates the time when a queue is created.`,
				},
				resource.Attribute{
					Name:        "reservation",
					Description: `Indicates the retention period (unit: min) of a message in a queue.`,
				},
				resource.Attribute{
					Name:        "max_msg_size_byte",
					Description: `Indicates the maximum message size (unit: byte) that is allowed in queue.`,
				},
				resource.Attribute{
					Name:        "produced_messages",
					Description: `Indicates the total number of messages (not including the messages that have expired and been deleted) in a queue.`,
				},
				resource.Attribute{
					Name:        "group_count",
					Description: `Indicates the total number of consumer groups in a queue.`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "queue_mode",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "redrive_policy",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "max_consume_count",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "created",
					Description: `Indicates the time when a queue is created.`,
				},
				resource.Attribute{
					Name:        "reservation",
					Description: `Indicates the retention period (unit: min) of a message in a queue.`,
				},
				resource.Attribute{
					Name:        "max_msg_size_byte",
					Description: `Indicates the maximum message size (unit: byte) that is allowed in queue.`,
				},
				resource.Attribute{
					Name:        "produced_messages",
					Description: `Indicates the total number of messages (not including the messages that have expired and been deleted) in a queue.`,
				},
				resource.Attribute{
					Name:        "group_count",
					Description: `Indicates the total number of consumer groups in a queue.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opentelekomcloud_dns_recordset_v2",
			Category:         "DNS Resources",
			ShortDescription: `Manages a DNS record set in the OpenTelekomCloud DNS Service`,
			Description:      ``,
			Keywords: []string{
				"dns",
				"recordset",
				"v2",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "zone_id",
					Description: `(Required) The ID of the zone in which to create the record set. Changing this creates a new DNS record set.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the record set. Note the ` + "`" + `.` + "`" + ` at the end of the name. Changing this creates a new DNS record set.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) The type of record set. Examples: "A", "MX". Changing this creates a new DNS record set.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `(Optional) The time to live (TTL) of the record set.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A description of the record set.`,
				},
				resource.Attribute{
					Name:        "records",
					Description: `(Optional) An array of DNS records.`,
				},
				resource.Attribute{
					Name:        "value_specs",
					Description: `(Optional) Map of additional options. Changing this creates a new record set. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "records",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "zone_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "value_specs",
					Description: `See Argument Reference above. ## Import This resource can be imported by specifying the zone ID and recordset ID, separated by a forward slash. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import opentelekomcloud_dns_recordset_v2.recordset_1 <zone_id>/<recordset_id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "records",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "zone_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "value_specs",
					Description: `See Argument Reference above. ## Import This resource can be imported by specifying the zone ID and recordset ID, separated by a forward slash. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import opentelekomcloud_dns_recordset_v2.recordset_1 <zone_id>/<recordset_id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opentelekomcloud_dns_zone_v2",
			Category:         "DNS Resources",
			ShortDescription: `Manages a DNS zone in the OpenTelekomCloud DNS Service`,
			Description:      ``,
			Keywords: []string{
				"dns",
				"zone",
				"v2",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the zone. Note the ` + "`" + `.` + "`" + ` at the end of the name. Changing this creates a new DNS zone.`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `(Optional) The email contact for the zone record.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) The type of zone. Can either be ` + "`" + `public` + "`" + ` or ` + "`" + `private` + "`" + `. Changing this creates a new zone.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `(Optional) The time to live (TTL) of the zone.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A description of the zone.`,
				},
				resource.Attribute{
					Name:        "masters",
					Description: `(Optional) An array of master DNS servers.`,
				},
				resource.Attribute{
					Name:        "value_specs",
					Description: `(Optional) Map of additional options. Changing this creates a new zone. The ` + "`" + `router` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "masters",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "value_specs",
					Description: `See Argument Reference above. ## Import This resource can be imported by specifying the zone ID: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import opentelekomcloud_dns_zone_v2.zone_1 <zone_id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "masters",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "value_specs",
					Description: `See Argument Reference above. ## Import This resource can be imported by specifying the zone ID: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import opentelekomcloud_dns_zone_v2.zone_1 <zone_id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opentelekomcloud_elb_backend",
			Category:         "Elastic Load Balancer Resources",
			ShortDescription: `Manages an elastic loadbalancer backend resource within OpentelekomCloud.`,
			Description:      ``,
			Keywords: []string{
				"elastic",
				"load",
				"balancer",
				"elb",
				"backend",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "listener_id",
					Description: `(Required) Specifies the listener ID. Changing this creates a new elb backend.`,
				},
				resource.Attribute{
					Name:        "server_id",
					Description: `(Required) Specifies the backend member ID. Changing this creates a new elb backend.`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `(Required) Specifies the private IP address of the backend member. Changing this creates a new elb backend. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "listener_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "server_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "server_address",
					Description: `Specifies the floating IP address assigned to the backend member.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Specifies the backend member ID.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Specifies the backend ECS status. The value is ACTIVE, PENDING, or ERROR.`,
				},
				resource.Attribute{
					Name:        "health_status",
					Description: `Specifies the health check status. The value is NORMAL, ABNORMAL, or UNAVAILABLE.`,
				},
				resource.Attribute{
					Name:        "update_time",
					Description: `Specifies the time when information about the backend member was updated.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Specifies the time when the backend member was created.`,
				},
				resource.Attribute{
					Name:        "server_name",
					Description: `Specifies the backend member name.`,
				},
				resource.Attribute{
					Name:        "listeners",
					Description: `Specifies the listener to which the backend member belongs.`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "listener_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "server_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "server_address",
					Description: `Specifies the floating IP address assigned to the backend member.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Specifies the backend member ID.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Specifies the backend ECS status. The value is ACTIVE, PENDING, or ERROR.`,
				},
				resource.Attribute{
					Name:        "health_status",
					Description: `Specifies the health check status. The value is NORMAL, ABNORMAL, or UNAVAILABLE.`,
				},
				resource.Attribute{
					Name:        "update_time",
					Description: `Specifies the time when information about the backend member was updated.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Specifies the time when the backend member was created.`,
				},
				resource.Attribute{
					Name:        "server_name",
					Description: `Specifies the backend member name.`,
				},
				resource.Attribute{
					Name:        "listeners",
					Description: `Specifies the listener to which the backend member belongs.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opentelekomcloud_elb_health",
			Category:         "Elastic Load Balancer Resources",
			ShortDescription: `Manages an elastic loadbalancer health resource within OpentelekomCloud.`,
			Description:      ``,
			Keywords: []string{
				"elastic",
				"load",
				"balancer",
				"elb",
				"health",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "listener_id",
					Description: `(Required) Specifies the ID of the listener to which the health check task belongs. Changing this creates a new elb health.`,
				},
				resource.Attribute{
					Name:        "healthcheck_protocol",
					Description: `(Optional) Specifies the protocol used for the health check. The value can be HTTP or TCP (case-insensitive).`,
				},
				resource.Attribute{
					Name:        "healthcheck_uri",
					Description: `(Optional) Specifies the URI for health check. This parameter is valid when healthcheck_ protocol is HTTP. The value is a string of 1 to 80 characters that must start with a slash (/) and can only contain letters, digits, and special characters, such as -/.%?#&.`,
				},
				resource.Attribute{
					Name:        "healthcheck_connect_port",
					Description: `(Optional) Specifies the port used for the health check. The value ranges from 1 to 65535.`,
				},
				resource.Attribute{
					Name:        "healthy_threshold",
					Description: `(Optional) Specifies the threshold at which the health check result is success, that is, the number of consecutive successful health checks when the health check result of the backend server changes from fail to success. The value ranges from 1 to 10.`,
				},
				resource.Attribute{
					Name:        "unhealthy_threshold",
					Description: `(Optional) Specifies the threshold at which the health check result is fail, that is, the number of consecutive failed health checks when the health check result of the backend server changes from success to fail. The value ranges from 1 to 10.`,
				},
				resource.Attribute{
					Name:        "healthcheck_timeout",
					Description: `(Optional) Specifies the maximum timeout duration (s) for the health check. The value ranges from 1 to 50.`,
				},
				resource.Attribute{
					Name:        "healthcheck_interval",
					Description: `(Optional) Specifies the maximum interval (s) for health check. The value ranges from 1 to 5. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "listener_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "healthcheck_protocol",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "healthcheck_uri",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "healthcheck_connect_port",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "healthy_threshold",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "unhealthy_threshold",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "healthcheck_timeout",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "healthcheck_interval",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Specifies the health check task ID.`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "listener_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "healthcheck_protocol",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "healthcheck_uri",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "healthcheck_connect_port",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "healthy_threshold",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "unhealthy_threshold",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "healthcheck_timeout",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "healthcheck_interval",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Specifies the health check task ID.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opentelekomcloud_elb_listener",
			Category:         "Elastic Load Balancer Resources",
			ShortDescription: `Manages an elastic loadbalancer listener resource within OpentelekomCloud.`,
			Description:      ``,
			Keywords: []string{
				"elastic",
				"load",
				"balancer",
				"elb",
				"listener",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Specifies the load balancer name. The name is a string of 1 to 64 characters that consist of letters, digits, underscores (_), and hyphens (-).`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Provides supplementary information about the listener. The value is a string of 0 to 128 characters and cannot be <>.`,
				},
				resource.Attribute{
					Name:        "loadbalancer_id",
					Description: `(Required) Specifies the ID of the load balancer to which the listener belongs. Changing this creates a new elb listener.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Required) Specifies the listening protocol used for layer 4 or 7. The value can be HTTP, TCP, HTTPS, or UDP. Changing this creates a new elb listener.`,
				},
				resource.Attribute{
					Name:        "protocol_port",
					Description: `(Required) Specifies the listening port. The value ranges from 1 to 65535.`,
				},
				resource.Attribute{
					Name:        "backend_protocol",
					Description: `(Required) Specifies the backend protocol. If the value of protocol is UDP, the value of this parameter can only be UDP. The value can be HTTP, TCP, or UDP. Changing this creates a new elb listener.`,
				},
				resource.Attribute{
					Name:        "backend_port",
					Description: `(Required) Specifies the backend port. The value ranges from 1 to 65535.`,
				},
				resource.Attribute{
					Name:        "lb_algorithm",
					Description: `(Required) Specifies the load balancing algorithm for the listener. The value can be roundrobin, leastconn, or source.`,
				},
				resource.Attribute{
					Name:        "session_sticky",
					Description: `(Optional) Specifies whether to enable sticky session. The value can be true or false. The Sticky session is enabled when the value is true, and is disabled when the value is false. If the value of protocol is HTTP, HTTPS, or TCP, and the value of lb_algorithm is not roundrobin, the value of this parameter can only be false. Changing this creates a new elb listener.`,
				},
				resource.Attribute{
					Name:        "session_sticky_type",
					Description: `(Optional) Specifies the cookie processing method. The value is insert. insert indicates that the cookie is inserted by the load balancer. This parameter is valid when protocol is set to HTTP, and session_sticky to true. The default value is insert. This parameter is invalid when protocol is set to TCP or UDP, which means the parameter is empty. Changing this creates a new elb listener.`,
				},
				resource.Attribute{
					Name:        "cookie_timeout",
					Description: `(Optional) Specifies the cookie timeout period (minutes). This parameter is valid when protocol is set to HTTP, session_sticky to true, and session_sticky_type to insert. This parameter is invalid when protocol is set to TCP or UDP. The value ranges from 1 to 1440. Changing this creates a new elb listener.`,
				},
				resource.Attribute{
					Name:        "tcp_timeout",
					Description: `(Optional) Specifies the TCP timeout period (minutes). This parameter is valid when protocol is set to TCP. The value ranges from 1 to 5.`,
				},
				resource.Attribute{
					Name:        "tcp_draining",
					Description: `(Optional) Specifies whether to maintain the TCP connection to the backend ECS after the ECS is deleted. This parameter is valid when protocol is set to TCP. The value can be true or false.`,
				},
				resource.Attribute{
					Name:        "tcp_draining_timeout",
					Description: `(Optional) Specifies the timeout duration (minutes) for the TCP connection to the backend ECS after the ECS is deleted. This parameter is valid when protocol is set to TCP, and tcp_draining to true. The value ranges from 0 to 60.`,
				},
				resource.Attribute{
					Name:        "certificate_id",
					Description: `(Optional) Specifies the ID of the SSL certificate used for security authentication when HTTPS is used to make API calls. This parameter is mandatory if the value of protocol is HTTPS. The value can be obtained by viewing the details of the SSL certificate. Changing this creates a new elb listener.`,
				},
				resource.Attribute{
					Name:        "udp_timeout",
					Description: `(Optional) Specifies the UDP timeout duration (minutes). This parameter is valid when protocol is set to UDP. The value ranges from 1 to 1440.`,
				},
				resource.Attribute{
					Name:        "ssl_protocols",
					Description: `(Optional) Specifies the SSL protocol standard supported by a tracker, which is used for enabling specified encryption protocols. This parameter is valid only when the value of protocol is set to HTTPS. The value is TLSv1.2 or TLSv1.2 TLSv1.1 TLSv1. The default value is TLSv1.2. Changing this creates a new elb listener.`,
				},
				resource.Attribute{
					Name:        "ssl_ciphers",
					Description: `(Optional) Specifies the cipher suite of an encryption protocol. This parameter is valid only when the value of protocol is set to HTTPS. The value is Default, Extended, or Strict. The default value is Default. The value can only be set to Extended if the value of ssl_protocols is set to TLSv1.2 TLSv1.1 TLSv1. ## Attributes Reference The following attributes are exported:`,
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
					Name:        "loadbalancer_id",
					Description: `See Argument Reference above.`,
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
					Name:        "backend_protocol",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "backend_port",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "lb_algorithm",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "session_sticky",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "session_sticky_type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "cookie_timeout",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tcp_timeout",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tcp_draining",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tcp_draining_timeout",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "certificate_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "udp_timeout",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "ssl_protocols",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "ssl_ciphers",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Specifies the listener ID.`,
				},
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `Specifies the status of the load balancer. Value range: false: The load balancer is disabled. true: The load balancer runs properly.`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "loadbalancer_id",
					Description: `See Argument Reference above.`,
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
					Name:        "backend_protocol",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "backend_port",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "lb_algorithm",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "session_sticky",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "session_sticky_type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "cookie_timeout",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tcp_timeout",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tcp_draining",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tcp_draining_timeout",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "certificate_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "udp_timeout",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "ssl_protocols",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "ssl_ciphers",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Specifies the listener ID.`,
				},
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `Specifies the status of the load balancer. Value range: false: The load balancer is disabled. true: The load balancer runs properly.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opentelekomcloud_elb_loadbalancer",
			Category:         "Elastic Load Balancer Resources",
			ShortDescription: `Manages an elastic loadbalancer resource within OpentelekomCloud.`,
			Description:      ``,
			Keywords: []string{
				"elastic",
				"load",
				"balancer",
				"elb",
				"loadbalancer",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Specifies the load balancer name. The name is a string of 1 to 64 characters that consist of letters, digits, underscores (_), and hyphens (-).`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Provides supplementary information about the listener. The value is a string of 0 to 128 characters and cannot be <>.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required) Specifies the VPC ID. Changing this creates a new elb loadbalancer.`,
				},
				resource.Attribute{
					Name:        "bandwidth",
					Description: `(Optional) Specifies the bandwidth (Mbit/s). This parameter is mandatory when type is set to External, and it is invalid when type is set to Internal. The value ranges from 1 to 300.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) Specifies the load balancer type. The value can be Internal or External. Changing this creates a new elb loadbalancer.`,
				},
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `(Required) Specifies the status of the load balancer. Value range: 0 or false: indicates that the load balancer is stopped. Only tenants are allowed to enter these two values. 1 or true: indicates that the load balancer is running properly. 2 or false: indicates that the load balancer is frozen. Only tenants are allowed to enter these two values.`,
				},
				resource.Attribute{
					Name:        "vip_subnet_id",
					Description: `(Optional) Specifies the ID of the private network to be added. This parameter is mandatory when type is set to Internal, and it is invalid when type is set to External. Changing this creates a new elb loadbalancer.`,
				},
				resource.Attribute{
					Name:        "az",
					Description: `(Optional) Specifies the ID of the availability zone (AZ). This parameter is mandatory when type is set to Internal, and it is invalid when type is set to External. Changing this creates a new elb loadbalancer.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `(Optional) Specifies the security group ID. The value is a string of 1 to 200 characters that consists of uppercase and lowercase letters, digits, and hyphens (-). This parameter is mandatory only when type is set to Internal. Changing this creates a new elb loadbalancer.`,
				},
				resource.Attribute{
					Name:        "vip_address",
					Description: `(Optional) Specifies the IP address provided by ELB. When type is set to External, the value of this parameter is the elastic IP address. When type is set to Internal, the value of this parameter is the private network IP address. You can select an existing elastic IP address and create a public network load balancer. When this parameter is configured, parameter bandwidth is invalid. Changing this creates a new elb loadbalancer.`,
				},
				resource.Attribute{
					Name:        "tenantid",
					Description: `(Optional) Specifies the tenant ID. This parameter is mandatory only when type is set to Internal. ## Attributes Reference The following attributes are exported:`,
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
					Name:        "vpc_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "bandwidth",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "vip_subnet_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "az",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "vip_address",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tenantid",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Specifies the load balancer ID.`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "bandwidth",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "vip_subnet_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "az",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "vip_address",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tenantid",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Specifies the load balancer ID.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opentelekomcloud_fw_firewall_group_v2",
			Category:         "Firewall Resources",
			ShortDescription: `Manages a v1 firewall group resource within OpenTelekomCloud.`,
			Description:      ``,
			Keywords: []string{
				"firewall",
				"fw",
				"group",
				"v2",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "ingress_policy_id",
					Description: `The ingress policy resource id for the firewall group. Changing this updates the ` + "`" + `ingress_policy_id` + "`" + ` of an existing firewall group.`,
				},
				resource.Attribute{
					Name:        "egress_policy_id",
					Description: `The egress policy resource id for the firewall group. Changing this updates the ` + "`" + `egress_policy_id` + "`" + ` of an existing firewall group.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) A name for the firewall group. Changing this updates the ` + "`" + `name` + "`" + ` of an existing firewall group.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Required) A description for the firewall group. Changing this updates the ` + "`" + `description` + "`" + ` of an existing firewall group.`,
				},
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `(Optional) Administrative up/down status for the firewall group (must be "true" or "false" if provided - defaults to "true"). Changing this updates the ` + "`" + `admin_state_up` + "`" + ` of an existing firewall group.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Optional) The owner of the floating IP. Required if admin wants to create a firewall group for another tenant. Changing this creates a new firewall group.`,
				},
				resource.Attribute{
					Name:        "ports",
					Description: `(Optional) Port(s) to associate this firewall group instance with. Must be a list of strings. Changing this updates the associated routers of an existing firewall group.`,
				},
				resource.Attribute{
					Name:        "value_specs",
					Description: `(Optional) Map of additional options. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "policy_id",
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
					Name:        "admin_state_up",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "ports",
					Description: `See Argument Reference above. ## Import Firewall Groups can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import opentelekomcloud_fw_firewall_group_v2.firewall_group_1 c9e39fb2-ce20-46c8-a964-25f3898c7a97 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "policy_id",
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
					Name:        "admin_state_up",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "ports",
					Description: `See Argument Reference above. ## Import Firewall Groups can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import opentelekomcloud_fw_firewall_group_v2.firewall_group_1 c9e39fb2-ce20-46c8-a964-25f3898c7a97 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opentelekomcloud_fw_policy_v2",
			Category:         "Firewall Resources",
			ShortDescription: `Manages a v1 firewall policy resource within OpenTelekomCloud.`,
			Description:      ``,
			Keywords: []string{
				"firewall",
				"fw",
				"policy",
				"v2",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) A name for the firewall policy. Changing this updates the ` + "`" + `name` + "`" + ` of an existing firewall policy.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A description for the firewall policy. Changing this updates the ` + "`" + `description` + "`" + ` of an existing firewall policy.`,
				},
				resource.Attribute{
					Name:        "rules",
					Description: `(Optional) An array of one or more firewall rules that comprise the policy. Changing this results in adding/removing rules from the existing firewall policy.`,
				},
				resource.Attribute{
					Name:        "audited",
					Description: `(Optional) Audit status of the firewall policy (must be "true" or "false" if provided - defaults to "false"). This status is set to "false" whenever the firewall policy or any of its rules are changed. Changing this updates the ` + "`" + `audited` + "`" + ` status of an existing firewall policy.`,
				},
				resource.Attribute{
					Name:        "shared",
					Description: `(Optional) Sharing status of the firewall policy (must be "true" or "false" if provided). If this is "true" the policy is visible to, and can be used in, firewalls in other tenants. Changing this updates the ` + "`" + `shared` + "`" + ` status of an existing firewall policy. Only administrative users can specify if the policy should be shared.`,
				},
				resource.Attribute{
					Name:        "value_specs",
					Description: `(Optional) Map of additional options. ## Attributes Reference The following attributes are exported:`,
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
					Name:        "audited",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "shared",
					Description: `See Argument Reference above. ## Import Firewall Policies can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import opentelekomcloud_fw_policy_v2.policy_1 07f422e6-c596-474b-8b94-fe2c12506ce0 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "audited",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "shared",
					Description: `See Argument Reference above. ## Import Firewall Policies can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import opentelekomcloud_fw_policy_v2.policy_1 07f422e6-c596-474b-8b94-fe2c12506ce0 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opentelekomcloud_fw_rule_v2",
			Category:         "Firewall Resources",
			ShortDescription: `Manages a v1 firewall group rule resource within OpenTelekomCloud.`,
			Description:      ``,
			Keywords: []string{
				"firewall",
				"fw",
				"rule",
				"v2",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) A unique name for the firewall rule. Changing this updates the ` + "`" + `name` + "`" + ` of an existing firewall rule.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A description for the firewall rule. Changing this updates the ` + "`" + `description` + "`" + ` of an existing firewall rule.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Required) The protocol type on which the firewall rule operates. Valid values are: ` + "`" + `tcp` + "`" + `, ` + "`" + `udp` + "`" + `, ` + "`" + `icmp` + "`" + `, and ` + "`" + `any` + "`" + `. Changing this updates the ` + "`" + `protocol` + "`" + ` of an existing firewall rule.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Required) Action to be taken ( must be "allow" or "deny") when the firewall rule matches. Changing this updates the ` + "`" + `action` + "`" + ` of an existing firewall rule.`,
				},
				resource.Attribute{
					Name:        "ip_version",
					Description: `(Optional) IP version, either 4 (default) or 6. Changing this updates the ` + "`" + `ip_version` + "`" + ` of an existing firewall rule.`,
				},
				resource.Attribute{
					Name:        "source_ip_address",
					Description: `(Optional) The source IP address on which the firewall rule operates. Changing this updates the ` + "`" + `source_ip_address` + "`" + ` of an existing firewall rule.`,
				},
				resource.Attribute{
					Name:        "destination_ip_address",
					Description: `(Optional) The destination IP address on which the firewall rule operates. Changing this updates the ` + "`" + `destination_ip_address` + "`" + ` of an existing firewall rule.`,
				},
				resource.Attribute{
					Name:        "source_port",
					Description: `(Optional) The source port on which the firewall rule operates. Changing this updates the ` + "`" + `source_port` + "`" + ` of an existing firewall rule.`,
				},
				resource.Attribute{
					Name:        "destination_port",
					Description: `(Optional) The destination port on which the firewall rule operates. Changing this updates the ` + "`" + `destination_port` + "`" + ` of an existing firewall rule.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) Enabled status for the firewall rule (must be "true" or "false" if provided - defaults to "true"). Changing this updates the ` + "`" + `enabled` + "`" + ` status of an existing firewall rule.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Optional) The owner of the firewall rule. Required if admin wants to create a firewall rule for another tenant. Changing this creates a new firewall rule.`,
				},
				resource.Attribute{
					Name:        "value_specs",
					Description: `(Optional) Map of additional options. ## Attributes Reference The following attributes are exported:`,
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
					Name:        "action",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "ip_version",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "source_ip_address",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "destination_ip_address",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "source_port",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "destination_port",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above. ## Import Firewall Rules can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import opentelekomcloud_fw_rule_v2.rule_1 8dbc0c28-e49c-463f-b712-5c5d1bbac327 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{
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
					Name:        "action",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "ip_version",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "source_ip_address",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "destination_ip_address",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "source_port",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "destination_port",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above. ## Import Firewall Rules can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import opentelekomcloud_fw_rule_v2.rule_1 8dbc0c28-e49c-463f-b712-5c5d1bbac327 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opentelekomcloud_identity_agency_v3",
			Category:         "Identity Resources",
			ShortDescription: `Manages an agency resource within Opentelekomcloud.`,
			Description:      ``,
			Keywords: []string{
				"identity",
				"agency",
				"v3",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of agency. The name is a string of 1 to 64 characters.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Provides supplementary information about the agency. The value is a string of 0 to 255 characters.`,
				},
				resource.Attribute{
					Name:        "delegated_domain_name",
					Description: `(Required) The name of delegated domain.`,
				},
				resource.Attribute{
					Name:        "project_role",
					Description: `(Optional) An array of roles and projects which are used to grant permissions to agency on project. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "domain_roles",
					Description: `(optional) An array of role names which stand for the permissionis to be granted to agency on domain. The ` + "`" + `project_role` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Required) The name of project`,
				},
				resource.Attribute{
					Name:        "roles",
					Description: `(Required) An array of role names`,
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
					Name:        "delegated_domain_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "project_role",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "domain_roles",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "duration",
					Description: `Validity period of an agency. The default value is null, indicating that the agency is permanently valid.`,
				},
				resource.Attribute{
					Name:        "expire_time",
					Description: `The expiration time of agency`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The time when the agency was created.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The agency ID.`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "delegated_domain_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "project_role",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "domain_roles",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "duration",
					Description: `Validity period of an agency. The default value is null, indicating that the agency is permanently valid.`,
				},
				resource.Attribute{
					Name:        "expire_time",
					Description: `The expiration time of agency`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The time when the agency was created.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The agency ID.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opentelekomcloud_identity_group_membership_v3",
			Category:         "Identity Resources",
			ShortDescription: `Manages the membership combine User Group resource and User resource within OpentelekomCloud IAM service.`,
			Description:      ``,
			Keywords: []string{
				"identity",
				"group",
				"membership",
				"v3",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "group",
					Description: `(Required) The group ID of this membership.`,
				},
				resource.Attribute{
					Name:        "users",
					Description: `(Required) A List of user IDs to associate to the group. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "group",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "users",
					Description: `See Argument Reference above.`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "group",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "users",
					Description: `See Argument Reference above.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opentelekomcloud_identity_group_v3",
			Category:         "Identity Resources",
			ShortDescription: `Manages a User Group resource within OpentelekomCloud IAM service.`,
			Description:      ``,
			Keywords: []string{
				"identity",
				"group",
				"v3",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the group.The length is less than or equal to 64 bytes`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A description of the group.`,
				},
				resource.Attribute{
					Name:        "domain_id",
					Description: `(Optional) The domain this group belongs to. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "domain_id",
					Description: `See Argument Reference above. ## Import Groups can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import opentelekomcloud_identity_group_v3.group_1 89c60255-9bd6-460c-822a-e2b959ede9d2 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "domain_id",
					Description: `See Argument Reference above. ## Import Groups can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import opentelekomcloud_identity_group_v3.group_1 89c60255-9bd6-460c-822a-e2b959ede9d2 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opentelekomcloud_identity_project_v3",
			Category:         "Identity Resources",
			ShortDescription: `Manages a Project resource within OpentelekomCloud Keystone.`,
			Description:      ``,
			Keywords: []string{
				"identity",
				"project",
				"v3",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the project. it must start with ID of an existing region_ and be less than or equal to 64 characters. Example: eu-de_project1.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A description of the project.`,
				},
				resource.Attribute{
					Name:        "domain_id",
					Description: `(Optional) The domain this project belongs to. Changing this creates a new Project.`,
				},
				resource.Attribute{
					Name:        "parent_id",
					Description: `(Optional) The parent of this project. Changing this creates a new Project. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "domain_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "parent_id",
					Description: `See Argument Reference above. ## Import Projects can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import opentelekomcloud_identity_project_v3.project_1 89c60255-9bd6-460c-822a-e2b959ede9d2 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "domain_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "parent_id",
					Description: `See Argument Reference above. ## Import Projects can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import opentelekomcloud_identity_project_v3.project_1 89c60255-9bd6-460c-822a-e2b959ede9d2 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opentelekomcloud_identity_role_assignment_v3",
			Category:         "Identity Resources",
			ShortDescription: `Manages a V3 Policy assignment within OpentelekomCloud IAM Service.`,
			Description:      ``,
			Keywords: []string{
				"identity",
				"role",
				"assignment",
				"v3",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "domain_id",
					Description: `(Optional; Required if ` + "`" + `project_id` + "`" + ` is empty) The domain to assign the role in.`,
				},
				resource.Attribute{
					Name:        "group_id",
					Description: `(Optional; Required if ` + "`" + `user_id` + "`" + ` is empty) The group to assign the role to.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Optional; Required if ` + "`" + `domain_id` + "`" + ` is empty) The project to assign the role in.`,
				},
				resource.Attribute{
					Name:        "user_id",
					Description: `(Optional; Required if ` + "`" + `group_id` + "`" + ` is empty) The user to assign the role in.`,
				},
				resource.Attribute{
					Name:        "role_id",
					Description: `(Required) The role to assign. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "domain_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "group_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "user_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "role_id",
					Description: `See Argument Reference above.`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "domain_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "group_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "user_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "role_id",
					Description: `See Argument Reference above.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opentelekomcloud_identity_role_v3",
			Category:         "Identity Resources",
			ShortDescription: `custom role management`,
			Description:      ``,
			Keywords: []string{
				"identity",
				"role",
				"v3",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "description",
					Description: `(Required) Description of a role. The value cannot exceed 256 characters.`,
				},
				resource.Attribute{
					Name:        "display_layer",
					Description: `(Required) Display layer of a role.\ndomain - A role is displayed at the domain layer.\nproject - A role is displayed at the project layer.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Required) Displayed name of a role. The value cannot exceed 64 characters.`,
				},
				resource.Attribute{
					Name:        "statement",
					Description: `(Required) Statement: The Statement field contains the Effect and Action elements. Effect indicates whether the policy allows or denies access. Action indicates authorization items. The number of statements cannot exceed 8. Structure is documented below. The ` + "`" + `statement` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Required) Permission set, which specifies the operation permissions on resources. The number of permission sets cannot exceed 100. Format: The value format is Service name:Resource type:Action, for example, vpc:ports:create. Service name: indicates the product name, such as ecs, evs, or vpc. Only lowercase letters are allowed. Resource type and Action: The values are case-insensitive, and the wildcard (`,
				},
				resource.Attribute{
					Name:        "effect",
					Description: `(Required) The value can be Allow and Deny. If both Allow and Deny are found in statements, the policy evaluation starts with Deny. - - - ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "catalog",
					Description: `Directory where a role locates`,
				},
				resource.Attribute{
					Name:        "domain_id",
					Description: `ID of the domain to which a role belongs`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of a role ## Import Role can be imported using the following format: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import opentelekomcloud_identity_role_v3.default {{ resource id}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "catalog",
					Description: `Directory where a role locates`,
				},
				resource.Attribute{
					Name:        "domain_id",
					Description: `ID of the domain to which a role belongs`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of a role ## Import Role can be imported using the following format: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import opentelekomcloud_identity_role_v3.default {{ resource id}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opentelekomcloud_identity_user_v3",
			Category:         "Identity Resources",
			ShortDescription: `Manages a User resource within OpentelekomCloud IAM service.`,
			Description:      ``,
			Keywords: []string{
				"identity",
				"user",
				"v3",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the user. The user name consists of 5 to 32 characters. It can contain only uppercase letters, lowercase letters, digits, spaces, and special characters (-_) and cannot start with a digit.`,
				},
				resource.Attribute{
					Name:        "default_project_id",
					Description: `(Optional) The default project this user belongs to.`,
				},
				resource.Attribute{
					Name:        "domain_id",
					Description: `(Optional) The domain this user belongs to.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) Whether the user is enabled or disabled. Valid values are ` + "`" + `true` + "`" + ` and ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional) The password for the user. It must contain at least two of the following character types: uppercase letters, lowercase letters, digits, and special characters. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "domain_id",
					Description: `See Argument Reference above. ## Import Users can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import opentelekomcloud_identity_user_v3.user_1 89c60255-9bd6-460c-822a-e2b959ede9d2 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "domain_id",
					Description: `See Argument Reference above. ## Import Users can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import opentelekomcloud_identity_user_v3.user_1 89c60255-9bd6-460c-822a-e2b959ede9d2 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opentelekomcloud_images_image_v2",
			Category:         "Images Resources",
			ShortDescription: `Manages a V2 Image resource within OpenTelekomCloud Glance.`,
			Description:      ``,
			Keywords: []string{
				"images",
				"image",
				"v2",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "container_format",
					Description: `(Required) The container format. Must be one of "ami", "ari", "aki", "bare", "ovf".`,
				},
				resource.Attribute{
					Name:        "disk_format",
					Description: `(Required) The disk format. Must be one of "ami", "ari", "aki", "vhd", "vmdk", "raw", "qcow2", "vdi", "iso".`,
				},
				resource.Attribute{
					Name:        "local_file_path",
					Description: `(Optional) This is the filepath of the raw image file that will be uploaded to Glance. Conflicts with ` + "`" + `image_source_url` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "image_cache_path",
					Description: `(Optional) This is the directory where the images will be downloaded. Images will be stored with a filename corresponding to the url's md5 hash. Defaults to "$HOME/.terraform/image_cache"`,
				},
				resource.Attribute{
					Name:        "image_source_url",
					Description: `(Optional) This is the url of the raw image that will be downloaded in the ` + "`" + `image_cache_path` + "`" + ` before being uploaded to Glance. Glance is able to download image from internet but the ` + "`" + `gophercloud` + "`" + ` library does not yet provide a way to do so. Conflicts with ` + "`" + `local_file_path` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "min_disk_gb",
					Description: `(Optional) Amount of disk space (in GB) required to boot image. Defaults to 0.`,
				},
				resource.Attribute{
					Name:        "min_ram_mb",
					Description: `(Optional) Amount of ram (in MB) required to boot image. Defauts to 0.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the image.`,
				},
				resource.Attribute{
					Name:        "protected",
					Description: `(Optional) If true, image will not be deletable. Defaults to false.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) The tags of the image. It must be a list of strings. At this time, it is not possible to delete all tags of an image.`,
				},
				resource.Attribute{
					Name:        "visibility",
					Description: `(Optional) The visibility of the image. Must be one of "public", "private", "community", or "shared". The ability to set the visibility depends upon the configuration of the OpenTelekomCloud cloud. Note: The ` + "`" + `properties` + "`" + ` attribute handling in the gophercloud library is currently buggy and needs to be fixed before being implemented in this resource. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "checksum",
					Description: `The checksum of the data associated with the image.`,
				},
				resource.Attribute{
					Name:        "container_format",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `The date the image was created.`,
				},
				resource.Attribute{
					Name:        "disk_format",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "file",
					Description: `the trailing path after the glance endpoint that represent the location of the image or the path to retrieve it.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `A unique ID assigned by Glance.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `The metadata associated with the image. Image metadata allow for meaningfully define the image properties and tags. See http://docs.openstack.org/developer/glance/metadefs-concepts.html.`,
				},
				resource.Attribute{
					Name:        "min_disk_gb",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "min_ram_mb",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "owner",
					Description: `The id of the opentelekomcloud user who owns the image.`,
				},
				resource.Attribute{
					Name:        "protected",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "schema",
					Description: `The path to the JSON-schema that represent the image or image`,
				},
				resource.Attribute{
					Name:        "size_bytes",
					Description: `The size in bytes of the data associated with the image.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the image. It can be "queued", "active" or "saving".`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "update_at",
					Description: `The date the image was last updated.`,
				},
				resource.Attribute{
					Name:        "visibility",
					Description: `See Argument Reference above. ## Import Images can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import opentelekomcloud_images_image_v2.rancheros 89c60255-9bd6-460c-822a-e2b959ede9d2 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "checksum",
					Description: `The checksum of the data associated with the image.`,
				},
				resource.Attribute{
					Name:        "container_format",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `The date the image was created.`,
				},
				resource.Attribute{
					Name:        "disk_format",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "file",
					Description: `the trailing path after the glance endpoint that represent the location of the image or the path to retrieve it.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `A unique ID assigned by Glance.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `The metadata associated with the image. Image metadata allow for meaningfully define the image properties and tags. See http://docs.openstack.org/developer/glance/metadefs-concepts.html.`,
				},
				resource.Attribute{
					Name:        "min_disk_gb",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "min_ram_mb",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "owner",
					Description: `The id of the opentelekomcloud user who owns the image.`,
				},
				resource.Attribute{
					Name:        "protected",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "schema",
					Description: `The path to the JSON-schema that represent the image or image`,
				},
				resource.Attribute{
					Name:        "size_bytes",
					Description: `The size in bytes of the data associated with the image.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the image. It can be "queued", "active" or "saving".`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "update_at",
					Description: `The date the image was last updated.`,
				},
				resource.Attribute{
					Name:        "visibility",
					Description: `See Argument Reference above. ## Import Images can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import opentelekomcloud_images_image_v2.rancheros 89c60255-9bd6-460c-822a-e2b959ede9d2 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opentelekomcloud_kms_key_v1",
			Category:         "Kms Resources",
			ShortDescription: `Manages a V1 key resource within KMS.`,
			Description:      ``,
			Keywords: []string{
				"kms",
				"key",
				"v1",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "key_alias",
					Description: `(Required) The alias in which to create the key. It is required when we create a new key. Changing this updates the alias of key.`,
				},
				resource.Attribute{
					Name:        "key_description",
					Description: `(Optional) The description of the key as viewed in OpenTelekomCloud console. Changing this updates the description of key.`,
				},
				resource.Attribute{
					Name:        "realm",
					Description: `(Optional) Region where a key resides. Changing this creates a new key.`,
				},
				resource.Attribute{
					Name:        "pending_days",
					Description: `(Optional) Duration in days after which the key is deleted after destruction of the resource, must be between 7 and 1096 days. Defaults to 7. It only be used when delete a key.`,
				},
				resource.Attribute{
					Name:        "is_enabled",
					Description: `(Optional) Specifies whether the key is enabled. Defaults to true. Changing this updates the state of existing key. ## Attributes Reference The following attributes are exported:`,
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
					Description: `The globally unique identifier for the key.`,
				},
				resource.Attribute{
					Name:        "default_key_flag",
					Description: `Identification of a Master Key. The value 1 indicates a Default Master Key, and the value 0 indicates a key.`,
				},
				resource.Attribute{
					Name:        "origin",
					Description: `Origin of a key. The default value is kms.`,
				},
				resource.Attribute{
					Name:        "scheduled_deletion_date",
					Description: `Scheduled deletion time (time stamp) of a key.`,
				},
				resource.Attribute{
					Name:        "domain_id",
					Description: `ID of a user domain for the key.`,
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
					Name:        "is_enabled",
					Description: `See Argument Reference above. ## Import KMS Keys can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import opentelekomcloud_kms_key_v1.key_1 7056d636-ac60-4663-8a6c-82d3c32c1c64 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{
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
					Description: `The globally unique identifier for the key.`,
				},
				resource.Attribute{
					Name:        "default_key_flag",
					Description: `Identification of a Master Key. The value 1 indicates a Default Master Key, and the value 0 indicates a key.`,
				},
				resource.Attribute{
					Name:        "origin",
					Description: `Origin of a key. The default value is kms.`,
				},
				resource.Attribute{
					Name:        "scheduled_deletion_date",
					Description: `Scheduled deletion time (time stamp) of a key.`,
				},
				resource.Attribute{
					Name:        "domain_id",
					Description: `ID of a user domain for the key.`,
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
					Name:        "is_enabled",
					Description: `See Argument Reference above. ## Import KMS Keys can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import opentelekomcloud_kms_key_v1.key_1 7056d636-ac60-4663-8a6c-82d3c32c1c64 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opentelekomcloud_lb_listener_v2",
			Category:         "Load Balancer Resources",
			ShortDescription: `Manages a V2 listener resource within OpenTelekomCloud.`,
			Description:      ``,
			Keywords: []string{
				"load",
				"balancer",
				"lb",
				"listener",
				"v2",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "protocol",
					Description: `(Required) The protocol - can either be TCP, HTTP, HTTPS or TERMINATED_HTTPS. Changing this creates a new Listener.`,
				},
				resource.Attribute{
					Name:        "protocol_port",
					Description: `(Required) The port on which to listen for client traffic. Changing this creates a new Listener.`,
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
					Name:        "default_tls_container_ref",
					Description: `(Optional) A reference to a Barbican Secrets container which stores TLS information. This is required if the protocol is ` + "`" + `TERMINATED_HTTPS` + "`" + `. See [here](https://wiki.openstack.org/wiki/Network/LBaaS/docs/how-to-create-tls-loadbalancer) for more information.`,
				},
				resource.Attribute{
					Name:        "sni_container_refs",
					Description: `(Optional) A list of references to Barbican Secrets containers which store SNI information. See [here](https://wiki.openstack.org/wiki/Network/LBaaS/docs/how-to-create-tls-loadbalancer) for more information.`,
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
					Name:        "default_tls_container_ref",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "sni_container_refs",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `See Argument Reference above.`,
				},
			},
			Attributes: []resource.Argument{
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
					Name:        "default_tls_container_ref",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "sni_container_refs",
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
			Type:             "opentelekomcloud_lb_loadbalancer_v2",
			Category:         "Load Balancer Resources",
			ShortDescription: `Manages a V2 loadbalancer resource within OpenTelekomCloud.`,
			Description:      ``,
			Keywords: []string{
				"load",
				"balancer",
				"lb",
				"loadbalancer",
				"v2",
			},
			Arguments: []resource.Argument{
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
					Description: `(Optional) The administrative state of the Loadbalancer. A valid value is only true (UP).`,
				},
				resource.Attribute{
					Name:        "loadbalancer_provider",
					Description: `(Optional) The name of the provider. Changing this creates a new loadbalancer.`,
				},
				resource.Attribute{
					Name:        "security_group_ids",
					Description: `(Optional) A list of security group IDs to apply to the loadbalancer. The security groups must be specified by ID and not name (as opposed to how they are configured with the Compute Instance). ## Attributes Reference The following attributes are exported:`,
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
					Name:        "security_group_ids",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "vip_port_id",
					Description: `The Port ID of the Load Balancer IP.`,
				},
			},
			Attributes: []resource.Argument{
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
					Name:        "security_group_ids",
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
			Type:             "opentelekomcloud_lb_member_v2",
			Category:         "Load Balancer Resources",
			ShortDescription: `Manages a V2 member resource within OpenTelekomCloud.`,
			Description:      ``,
			Keywords: []string{
				"load",
				"balancer",
				"lb",
				"member",
				"v2",
			},
			Arguments: []resource.Argument{
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
					Name:        "tenant_id",
					Description: `(Optional) Required for admins. The UUID of the tenant who owns the member. Only administrative users can specify a tenant UUID other than their own. Changing this creates a new member.`,
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
					Name:        "tenant_id",
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
			Attributes: []resource.Argument{
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
					Name:        "tenant_id",
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
			Type:             "opentelekomcloud_lb_monitor_v2",
			Category:         "Load Balancer Resources",
			ShortDescription: `Manages a V2 monitor resource within OpenTelekomCloud.`,
			Description:      ``,
			Keywords: []string{
				"load",
				"balancer",
				"lb",
				"monitor",
				"v2",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "pool_id",
					Description: `(Required) The id of the pool that this monitor will be assigned to.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The Name of the Monitor.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Optional) Required for admins. The UUID of the tenant who owns the monitor. Only administrative users can specify a tenant UUID other than their own. Changing this creates a new monitor.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) The type of probe, which is PING, TCP, HTTP, or HTTPS, that is sent by the load balancer to verify the member state. Changing this creates a new monitor.`,
				},
				resource.Attribute{
					Name:        "delay",
					Description: `(Required) The time, in seconds, between sending probes to members.`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `(Required) Maximum number of seconds for a monitor to wait for a ping reply before it times out. The value must be less than the delay value.`,
				},
				resource.Attribute{
					Name:        "max_retries",
					Description: `(Required) Number of permissible ping failures before changing the member's status to INACTIVE. Must be a number between 1 and 10..`,
				},
				resource.Attribute{
					Name:        "url_path",
					Description: `(Optional) Required for HTTP(S) types. URI path that will be accessed if monitor type is HTTP or HTTPS.`,
				},
				resource.Attribute{
					Name:        "expected_codes",
					Description: `(Optional) Required for HTTP(S) types. Expected HTTP codes for a passing HTTP(S) monitor. You can either specify a single status like "200", or a range like "200-202".`,
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
					Name:        "tenant_id",
					Description: `See Argument Reference above.`,
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
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "id",
					Description: `The unique ID for the monitor.`,
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
			Type:             "opentelekomcloud_lb_pool_v2",
			Category:         "Load Balancer Resources",
			ShortDescription: `Manages a V2 pool resource within OpenTelekomCloud.`,
			Description:      ``,
			Keywords: []string{
				"load",
				"balancer",
				"lb",
				"pool",
				"v2",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Optional) Required for admins. The UUID of the tenant who owns the pool. Only administrative users can specify a tenant UUID other than their own. Changing this creates a new pool.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Human-readable name for the pool.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Human-readable description for the pool.`,
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
					Description: `(Optional) The name of the cookie if persistence mode is set appropriately. Required if ` + "`" + `type = APP_COOKIE` + "`" + `. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique ID for the pool.`,
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
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "id",
					Description: `The unique ID for the pool.`,
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
			Type:             "opentelekomcloud_maas_task_v1",
			Category:         "MAAS Resources",
			ShortDescription: `Manages resource task within OpenTelekomCloud MAAS.`,
			Description:      ``,
			Keywords: []string{
				"maas",
				"task",
				"v1",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "src_node",
					Description: `(Required) Specifies the source node information.`,
				},
				resource.Attribute{
					Name:        "dst_node",
					Description: `(Required) Specifies the destination node information.`,
				},
				resource.Attribute{
					Name:        "enable_kms",
					Description: `(Required) Specifies whether to use KMS encryption.`,
				},
				resource.Attribute{
					Name:        "thread_num",
					Description: `(Required) Specifies the number of threads used by the migration task. The value cannot exceed 50.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Specifies tasks description, which cannot exceed 255 characters. The following special characters are not allowed: <>()"&`,
				},
				resource.Attribute{
					Name:        "smn_info",
					Description: `(Optional) Specifies the field used for sending messages using the Simple Message Notification (SMN) service. The ` + "`" + `src_node` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "ak",
					Description: `(Required) Specifies the source bucket Access Key.`,
				},
				resource.Attribute{
					Name:        "sk",
					Description: `(Required) Specifies the source bucket Secret Key.`,
				},
				resource.Attribute{
					Name:        "object_key",
					Description: `(Required) Specifies the name of the object to be selected in the source bucket.`,
				},
				resource.Attribute{
					Name:        "bucket",
					Description: `(Required) Specifies the name of the source bucket.`,
				},
				resource.Attribute{
					Name:        "cloud_type",
					Description: `(Optional) Specifies the source cloud vendor. The default value is AWS and only AWS is supported now. The ` + "`" + `dst_node` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "ak",
					Description: `(Required) Specifies the destination bucket Access Key.`,
				},
				resource.Attribute{
					Name:        "sk",
					Description: `(Required) Specifies the destination bucket Secret Key.`,
				},
				resource.Attribute{
					Name:        "object_key",
					Description: `(Optional) Specifies the name of the object to be selected in the destination bucket.`,
				},
				resource.Attribute{
					Name:        "bucket",
					Description: `(Required) Specifies the name of the destination bucket. The ` + "`" + `smn_info` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "topic_urn",
					Description: `(Required) Specifies the SMN message topic URN bound to a migration task.`,
				},
				resource.Attribute{
					Name:        "language",
					Description: `(Optional) Specifies the management console language used by the current users. Users can select en-us.`,
				},
				resource.Attribute{
					Name:        "trigger_conditions",
					Description: `(Required) Specifies the trigger conditions of sending messages using SMN. The value depending on the state of a migration task. The migration task status can be SUCCESS or FAIL. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "src_node",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "dst_node",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "enable_kms",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "thread_num",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "smn_info",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Specifies the name for a task.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Specifies the task status as follows: 0: Not started, 1: Waiting to migrate, 2: Migrating, 3: Migration paused, 4: Migration failed, 5: Migration succeeded.`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "src_node",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "dst_node",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "enable_kms",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "thread_num",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "smn_info",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Specifies the name for a task.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Specifies the task status as follows: 0: Not started, 1: Waiting to migrate, 2: Migrating, 3: Migration paused, 4: Migration failed, 5: Migration succeeded.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opentelekomcloud_mrs_cluster_v1",
			Category:         "MRS Resources",
			ShortDescription: `Manages resource cluster within OpenTelekomCloud MRS.`,
			Description:      ``,
			Keywords: []string{
				"mrs",
				"cluster",
				"v1",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "billing_type",
					Description: `(Required) The value is 12, indicating on-demand payment.`,
				},
				resource.Attribute{
					Name:        "master_node_num",
					Description: `(Required) Number of Master nodes The value is 2.`,
				},
				resource.Attribute{
					Name:        "master_node_size",
					Description: `(Required) Best match based on several years of commissioning experience. MRS supports specifications of hosts, and host specifications are determined by CPUs, memory, and disks space. Master nodes support h1.2xlarge.linux.mrs h1.4xlarge.linux.mrs, h1.8xlarge.linux.mrs, s1.4xlarge.linux.mrs, and s1.8xlarge.linux.mrs. Core nodes of a streaming cluster support all specifications c2.2xlarge.linux.mrs, c2.4xlarge.linux.mrs, s1.xlarge.linux.mrs, s1.4xlarge.linux.mrs, s1.8xlarge.linux.mrs, d1.xlarge.linux.mrs, d1.2xlarge.linux.mrs, d1.4xlarge.linux.mrs, d1.8xlarge.linux.mrs, h1.2xlarge.linux.mrs, h1.4xlarge.linux.mrs and h1.8xlarge.linux.mrs. Task nodes support c2.2xlarge.linux.mrs, c2.4xlarge.linux.mrs, s1.xlarge.linux.mrs, s1.4xlarge.linux.mrs, s1.8xlarge.linux.mrs, h1.2xlarge.linux.mrs, h1.4xlarge.linux.mrs, and h1.8xlarge.linux.mrs.`,
				},
				resource.Attribute{
					Name:        "core_node_num",
					Description: `(Required) Number of Core nodes Value range: 1 to 500 A maximum of 500 Core nodes are supported by default. If more than 500 Core nodes are required, contact technical support engineers or invoke background APIs to modify the database.`,
				},
				resource.Attribute{
					Name:        "core_node_size",
					Description: `(Required) Instance specification of a Core node Configuration method of this parameter is identical to that of master_node_size.`,
				},
				resource.Attribute{
					Name:        "available_zone_id",
					Description: `(Required) ID of an available zone. Obtain the value from Regions and Endpoints.`,
				},
				resource.Attribute{
					Name:        "cluster_name",
					Description: `(Required) Cluster name, which is globally unique and contains only 1 to 64 letters, digits, hyphens (-), and underscores (_).`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required) ID of the VPC where the subnet locates Obtain the VPC ID from the management console as follows: Register an account and log in to the management console. Click Virtual Private Cloud and select Virtual Private Cloud from the left list. On the Virtual Private Cloud page, obtain the VPC ID from the list.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Required) Subnet ID Obtain the subnet ID from the management console as follows: Register an account and log in to the management console. Click Virtual Private Cloud and select Virtual Private Cloud from the left list. On the Virtual Private Cloud page, obtain the subnet ID from the list.`,
				},
				resource.Attribute{
					Name:        "cluster_version",
					Description: `(Optional) Version of the clusters Currently, MRS 1.2, MRS 1.3.0, MRS 1.5.0, MRS 1.6.0 and MRS 1.6.3 are supported. The latest version of MRS is used by default. Currently, the latest version is MRS 1.6.3.`,
				},
				resource.Attribute{
					Name:        "cluster_type",
					Description: `(Optional) Type of clusters 0: analysis cluster 1: streaming cluster The default value is 0.`,
				},
				resource.Attribute{
					Name:        "volume_type",
					Description: `(Optional) Type of disks SATA, SAS and SSD are supported. SATA: common I/O, SAS: High I/O, SSD: Ultra-high I/O.`,
				},
				resource.Attribute{
					Name:        "volume_size",
					Description: `(Optional) Data disk storage space of a Core node Users can add disks to expand storage capacity when creating a cluster. There are the following scenarios: Separation of data storage and computing: Data is stored in the OBS system. Costs of clusters are relatively low but computing performance is poor. The clusters can be deleted at any time. It is recommended when data computing is not frequently performed. Integration of data storage and computing: Data is stored in the HDFS system. Costs of clusters are relatively high but computing performance is good. The clusters cannot be deleted in a short term. It is recommended when data computing is frequently performed. Value range: 100 GB to 32000 GB`,
				},
				resource.Attribute{
					Name:        "master_data_volume_type",
					Description: `(Optional) Data disk storage type of the Master node, supporting SATA, SAS and SSD. SATA: Common I/O, SAS: High I/O, SSD: Ultra-high I/O. This parameter is a multi-disk parameter available for MRS 1.6.0 or later versions.`,
				},
				resource.Attribute{
					Name:        "master_data_volume_size",
					Description: `(Optional) Data disk size of the Master node. Value range: 100 GB to 32000 GB. This parameter is a multi-disk parameter available for MRS 1.6.0 or later versions.`,
				},
				resource.Attribute{
					Name:        "master_data_volume_count",
					Description: `(Optional) Number of data disks of the Master node. The value can be set to 1 only. This parameter is a multi-disk parameter available for MRS 1.6.0 or later versions.`,
				},
				resource.Attribute{
					Name:        "core_data_volume_type",
					Description: `(Optional) Data disk storage type of the Core node, supporting SATA, SAS and SSD. SATA: Common I/O, SAS: High I/O, SSD: Ultra-high I/O. This parameter is a multi-disk parameter available for MRS 1.6.0 or later versions.`,
				},
				resource.Attribute{
					Name:        "core_data_volume_size",
					Description: `(Optional) Data disk size of the Core node. Value range: 100 GB to 32000 GB. This parameter is a multi-disk parameter available for MRS 1.6.0 or later versions.`,
				},
				resource.Attribute{
					Name:        "core_data_volume_count",
					Description: `(Optional) Number of data disks of the Core node. Value range: 1 to 10. This parameter is a multi-disk parameter available for MRS 1.6.0 or later versions.`,
				},
				resource.Attribute{
					Name:        "node_public_cert_name",
					Description: `(Required) Name of a key pair You can use a key to log in to the Master node in the cluster.`,
				},
				resource.Attribute{
					Name:        "safe_mode",
					Description: `(Required) MRS cluster running mode 0: common mode The value indicates that the Kerberos authentication is disabled. Users can use all functions provided by the cluster. 1: safe mode The value indicates that the Kerberos authentication is enabled. Common users cannot use the file management or job management functions of an MRS cluster and cannot view cluster resource usage or the job records of Hadoop and Spark. To use these functions, the users must obtain the relevant permissions from the MRS Manager administrator. The request has the cluster_admin_secret parameter only when safe_mode is set to 1.`,
				},
				resource.Attribute{
					Name:        "cluster_admin_secret",
					Description: `(Optional) Indicates the password of the MRS Manager administrator. The password for MRS 1.5.0: Must contain 6 to 32 characters. Must contain at least two types of the following: Lowercase letters Uppercase letters Digits Special characters of ` + "`" + `~!@#$%^&`,
				},
				resource.Attribute{
					Name:        "log_collection",
					Description: `(Optional) Indicates whether logs are collected when cluster installation fails. 0: not collected 1: collected The default value is 0. If log_collection is set to 1, OBS buckets will be created to collect the MRS logs. These buckets will be charged.`,
				},
				resource.Attribute{
					Name:        "component_list",
					Description: `(Required) Service component list.`,
				},
				resource.Attribute{
					Name:        "add_jobs",
					Description: `(Optional) You can submit a job when you create a cluster to save time and use MRS easily. Only one job can be added. The ` + "`" + `component_list` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "component_name",
					Description: `(Required) Component name Currently, Hadoop, Spark, HBase, Hive, Hue, Loader, Flume, Kafka and Storm are supported. Loader and Flume are not supported by MRS 1.3.0. Kafka and Storm are not supported by MRS 1.2.`,
				},
				resource.Attribute{
					Name:        "componentId",
					Description: `Component ID Component IDs supported by MRS 1.5.0 include: MRS 1.5.0_001: Hadoop MRS 1.5.0_002: Spark MRS 1.5.0_003: HBase MRS 1.5.0_004: Hive MRS 1.5.0_005: Hue MRS 1.5.0_006: Kafka MRS 1.5.0_007: Storm MRS 1.5.0_008: Loader MRS 1.5.0_009: Flume Component IDs supported by MRS 1.3.0 include: MRS 1.3.0_001: Hadoop MRS 1.3.0_002: Spark MRS 1.3.0_003: HBase MRS 1.3.0_004: Hive MRS 1.3.0_005: Hue MRS 1.3.0_006: Kafka MRS 1.3.0_007: Storm For example, the component ID of Hadoop is MRS 1.5.0_001, or MRS 1.3.0_001.`,
				},
				resource.Attribute{
					Name:        "componentName",
					Description: `Component name Currently, Hadoop, Spark, HBase, Hive, Hue, Loader, Flume, Kafka and Storm are supported. Loader and Flume are not supported by MRS 1.3.0.`,
				},
				resource.Attribute{
					Name:        "componentVersion",
					Description: `Component version MRS 1.5.0 supports the following component version: Component version of an analysis cluster: Hadoop: 2.7.2 Spark: 2.1.0 HBase: 1.0.2 Hive: 1.2.1 Hue: 3.11.0 Loader: 2.0.0 Component version of a streaming cluster: Kafka: 0.10.0.0 Storm: 1.0.2 Flume: 1.6.0 MRS 1.3.0 supports the following component version: Component version of an analysis cluster: Hadoop: 2.7.2 Spark: 1.5.1 HBase: 1.0.2 Hive: 1.2.1 Hue: 3.11.0 Component version of a streaming cluster: Kafka: 0.10.0.0 Storm: 1.0.2`,
				},
				resource.Attribute{
					Name:        "componentDesc",
					Description: `Component description The ` + "`" + `add_jobs` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "job_type",
					Description: `(Required) Job type 1: MapReduce 2: Spark 3: Hive Script 4: HiveQL (not supported currently) 5: DistCp, importing and exporting data (not supported in this API currently). 6: Spark Script 7: Spark SQL, submitting Spark SQL statements (not supported in this API currently). NOTE: Spark and Hive jobs can be added to only clusters including Spark and Hive components.`,
				},
				resource.Attribute{
					Name:        "job_name",
					Description: `(Required) Job name It contains only 1 to 64 letters, digits, hyphens (-), and underscores (_). NOTE: Identical job names are allowed but not recommended.`,
				},
				resource.Attribute{
					Name:        "jar_path",
					Description: `(Required) Path of the .jar file or .sql file for program execution The parameter must meet the following requirements: Contains a maximum of 1023 characters, excluding special characters such as ;|&><'$. The address cannot be empty or full of spaces. Starts with / or s3a://. Spark Script must end with .sql; while MapReduce and Spark Jar must end with .jar. sql and jar are case-insensitive.`,
				},
				resource.Attribute{
					Name:        "arguments",
					Description: `(Optional) Key parameter for program execution The parameter is specified by the function of the user's program. MRS is only responsible for loading the parameter. The parameter contains a maximum of 2047 characters, excluding special characters such as ;|&>'<$, and can be empty.`,
				},
				resource.Attribute{
					Name:        "input",
					Description: `(Optional) Path for inputting data, which must start with / or s3a://. A correct OBS path is required. The parameter contains a maximum of 1023 characters, excluding special characters such as ;|&>'<$, and can be empty.`,
				},
				resource.Attribute{
					Name:        "output",
					Description: `(Optional) Path for outputting data, which must start with / or s3a://. A correct OBS path is required. If the path does not exist, the system automatically creates it. The parameter contains a maximum of 1023 characters, excluding special characters such as ;|&>'<$, and can be empty.`,
				},
				resource.Attribute{
					Name:        "job_log",
					Description: `(Optional) Path for storing job logs that record job running status. This path must start with / or s3a://. A correct OBS path is required. The parameter contains a maximum of 1023 characters, excluding special characters such as ;|&>'<$, and can be empty.`,
				},
				resource.Attribute{
					Name:        "shutdown_cluster",
					Description: `(Optional) Whether to delete the cluster after the jobs are complete true: Yes false: No`,
				},
				resource.Attribute{
					Name:        "file_action",
					Description: `(Optional) Data import and export import export`,
				},
				resource.Attribute{
					Name:        "submit_job_once_cluster_run",
					Description: `(Required) true: A job is submitted when a cluster is created. false: A job is submitted separately. The parameter is set to true in this example.`,
				},
				resource.Attribute{
					Name:        "hql",
					Description: `(Optional) HiveQL statement`,
				},
				resource.Attribute{
					Name:        "hive_script_path",
					Description: `(Optional) SQL program path This parameter is needed by Spark Script and Hive Script jobs only and must meet the following requirements: Contains a maximum of 1023 characters, excluding special characters such as ;|&><'$. The address cannot be empty or full of spaces. Starts with / or s3a://. Ends with .sql. sql is case-insensitive. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "billing_type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "data_center",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "master_node_num",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "master_node_size",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "core_node_num",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "core_node_size",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "available_zone_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "cluster_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "cluster_version",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "cluster_type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "volume_type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "volume_size",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "node_public_cert_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "safe_mode",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "cluster_admin_secret",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "log_collection",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "component_list",
					Description: `See Argument Reference below.`,
				},
				resource.Attribute{
					Name:        "add_jobs",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "order_id",
					Description: `Order ID for creating clusters.`,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `Cluster ID.`,
				},
				resource.Attribute{
					Name:        "available_zone_name",
					Description: `Name of an availability zone.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `Instance ID.`,
				},
				resource.Attribute{
					Name:        "hadoop_version",
					Description: `Hadoop version.`,
				},
				resource.Attribute{
					Name:        "master_node_ip",
					Description: `IP address of a Master node.`,
				},
				resource.Attribute{
					Name:        "externalIp",
					Description: `Internal IP address.`,
				},
				resource.Attribute{
					Name:        "private_ip_first",
					Description: `Primary private IP address.`,
				},
				resource.Attribute{
					Name:        "external_ip",
					Description: `External IP address.`,
				},
				resource.Attribute{
					Name:        "slave_security_groups_id",
					Description: `Standby security group ID.`,
				},
				resource.Attribute{
					Name:        "security_groups_id",
					Description: `Security group ID.`,
				},
				resource.Attribute{
					Name:        "external_alternate_ip",
					Description: `Backup external IP address.`,
				},
				resource.Attribute{
					Name:        "master_node_spec_id",
					Description: `Specification ID of a Master node.`,
				},
				resource.Attribute{
					Name:        "core_node_spec_id",
					Description: `Specification ID of a Core node.`,
				},
				resource.Attribute{
					Name:        "master_node_product_id",
					Description: `Product ID of a Master node.`,
				},
				resource.Attribute{
					Name:        "core_node_product_id",
					Description: `Product ID of a Core node.`,
				},
				resource.Attribute{
					Name:        "duration",
					Description: `Cluster subscription duration.`,
				},
				resource.Attribute{
					Name:        "vnc",
					Description: `URI address for remote login of the elastic cloud server.`,
				},
				resource.Attribute{
					Name:        "fee",
					Description: `Cluster creation fee, which is automatically calculated.`,
				},
				resource.Attribute{
					Name:        "deployment_id",
					Description: `Deployment ID of a cluster.`,
				},
				resource.Attribute{
					Name:        "cluster_state",
					Description: `Cluster status Valid values include: existing history starting running terminated failed abnormal terminating rebooting shutdown frozen scaling-out scaling-in scaling-error.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `Project ID.`,
				},
				resource.Attribute{
					Name:        "create_at",
					Description: `Cluster creation time.`,
				},
				resource.Attribute{
					Name:        "update_at",
					Description: `Cluster update time.`,
				},
				resource.Attribute{
					Name:        "error_info",
					Description: `Error information.`,
				},
				resource.Attribute{
					Name:        "charging_start_time",
					Description: `Time when charging starts.`,
				},
				resource.Attribute{
					Name:        "remark",
					Description: `Remarks of a cluster. The component_list attributes:`,
				},
				resource.Attribute{
					Name:        "component_name",
					Description: `(Required) Component name Currently, Hadoop, Spark, HBase, Hive, Hue, Loader, Flume, Kafka and Storm are supported. Loader and Flume are not supported by MRS 1.3.0.`,
				},
				resource.Attribute{
					Name:        "component_id",
					Description: `Component ID Component IDs supported by MRS 1.5.0 include: MRS 1.5.0_001: Hadoop MRS 1.5.0_002: Spark MRS 1.5.0_003: HBase MRS 1.5.0_004: Hive MRS 1.5.0_005: Hue MRS 1.5.0_006: Kafka MRS 1.5.0_007: Storm MRS 1.5.0_008: Loader MRS 1.5.0_009: Flume Component IDs supported by MRS 1.3.0 include: MRS 1.3.0_001: Hadoop MRS 1.3.0_002: Spark MRS 1.3.0_003: HBase MRS 1.3.0_004: Hive MRS 1.3.0_005: Hue MRS 1.3.0_006: Kafka MRS 1.3.0_007: Storm For example, the component ID of Hadoop is MRS 1.5.0_001, or MRS 1.3.0_001.`,
				},
				resource.Attribute{
					Name:        "component_version",
					Description: `Component version MRS 1.5.0 supports the following component version: Component version of an analysis cluster: Hadoop: 2.7.2 Spark: 2.1.0 HBase: 1.0.2 Hive: 1.2.1 Hue: 3.11.0 Loader: 2.0.0 Component version of a streaming cluster: Kafka: 0.10.0.0 Storm: 1.0.2 Flume: 1.6.0 MRS 1.3.0 supports the following component version: Component version of an analysis cluster: Hadoop: 2.7.2 Spark: 1.5.1 HBase: 1.0.2 Hive: 1.2.1 Hue: 3.11.0 Component version of a streaming cluster: Kafka: 0.10.0.0 Storm: 1.0.2`,
				},
				resource.Attribute{
					Name:        "component_desc",
					Description: `Component description`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "billing_type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "data_center",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "master_node_num",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "master_node_size",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "core_node_num",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "core_node_size",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "available_zone_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "cluster_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "cluster_version",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "cluster_type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "volume_type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "volume_size",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "node_public_cert_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "safe_mode",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "cluster_admin_secret",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "log_collection",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "component_list",
					Description: `See Argument Reference below.`,
				},
				resource.Attribute{
					Name:        "add_jobs",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "order_id",
					Description: `Order ID for creating clusters.`,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `Cluster ID.`,
				},
				resource.Attribute{
					Name:        "available_zone_name",
					Description: `Name of an availability zone.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `Instance ID.`,
				},
				resource.Attribute{
					Name:        "hadoop_version",
					Description: `Hadoop version.`,
				},
				resource.Attribute{
					Name:        "master_node_ip",
					Description: `IP address of a Master node.`,
				},
				resource.Attribute{
					Name:        "externalIp",
					Description: `Internal IP address.`,
				},
				resource.Attribute{
					Name:        "private_ip_first",
					Description: `Primary private IP address.`,
				},
				resource.Attribute{
					Name:        "external_ip",
					Description: `External IP address.`,
				},
				resource.Attribute{
					Name:        "slave_security_groups_id",
					Description: `Standby security group ID.`,
				},
				resource.Attribute{
					Name:        "security_groups_id",
					Description: `Security group ID.`,
				},
				resource.Attribute{
					Name:        "external_alternate_ip",
					Description: `Backup external IP address.`,
				},
				resource.Attribute{
					Name:        "master_node_spec_id",
					Description: `Specification ID of a Master node.`,
				},
				resource.Attribute{
					Name:        "core_node_spec_id",
					Description: `Specification ID of a Core node.`,
				},
				resource.Attribute{
					Name:        "master_node_product_id",
					Description: `Product ID of a Master node.`,
				},
				resource.Attribute{
					Name:        "core_node_product_id",
					Description: `Product ID of a Core node.`,
				},
				resource.Attribute{
					Name:        "duration",
					Description: `Cluster subscription duration.`,
				},
				resource.Attribute{
					Name:        "vnc",
					Description: `URI address for remote login of the elastic cloud server.`,
				},
				resource.Attribute{
					Name:        "fee",
					Description: `Cluster creation fee, which is automatically calculated.`,
				},
				resource.Attribute{
					Name:        "deployment_id",
					Description: `Deployment ID of a cluster.`,
				},
				resource.Attribute{
					Name:        "cluster_state",
					Description: `Cluster status Valid values include: existing history starting running terminated failed abnormal terminating rebooting shutdown frozen scaling-out scaling-in scaling-error.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `Project ID.`,
				},
				resource.Attribute{
					Name:        "create_at",
					Description: `Cluster creation time.`,
				},
				resource.Attribute{
					Name:        "update_at",
					Description: `Cluster update time.`,
				},
				resource.Attribute{
					Name:        "error_info",
					Description: `Error information.`,
				},
				resource.Attribute{
					Name:        "charging_start_time",
					Description: `Time when charging starts.`,
				},
				resource.Attribute{
					Name:        "remark",
					Description: `Remarks of a cluster. The component_list attributes:`,
				},
				resource.Attribute{
					Name:        "component_name",
					Description: `(Required) Component name Currently, Hadoop, Spark, HBase, Hive, Hue, Loader, Flume, Kafka and Storm are supported. Loader and Flume are not supported by MRS 1.3.0.`,
				},
				resource.Attribute{
					Name:        "component_id",
					Description: `Component ID Component IDs supported by MRS 1.5.0 include: MRS 1.5.0_001: Hadoop MRS 1.5.0_002: Spark MRS 1.5.0_003: HBase MRS 1.5.0_004: Hive MRS 1.5.0_005: Hue MRS 1.5.0_006: Kafka MRS 1.5.0_007: Storm MRS 1.5.0_008: Loader MRS 1.5.0_009: Flume Component IDs supported by MRS 1.3.0 include: MRS 1.3.0_001: Hadoop MRS 1.3.0_002: Spark MRS 1.3.0_003: HBase MRS 1.3.0_004: Hive MRS 1.3.0_005: Hue MRS 1.3.0_006: Kafka MRS 1.3.0_007: Storm For example, the component ID of Hadoop is MRS 1.5.0_001, or MRS 1.3.0_001.`,
				},
				resource.Attribute{
					Name:        "component_version",
					Description: `Component version MRS 1.5.0 supports the following component version: Component version of an analysis cluster: Hadoop: 2.7.2 Spark: 2.1.0 HBase: 1.0.2 Hive: 1.2.1 Hue: 3.11.0 Loader: 2.0.0 Component version of a streaming cluster: Kafka: 0.10.0.0 Storm: 1.0.2 Flume: 1.6.0 MRS 1.3.0 supports the following component version: Component version of an analysis cluster: Hadoop: 2.7.2 Spark: 1.5.1 HBase: 1.0.2 Hive: 1.2.1 Hue: 3.11.0 Component version of a streaming cluster: Kafka: 0.10.0.0 Storm: 1.0.2`,
				},
				resource.Attribute{
					Name:        "component_desc",
					Description: `Component description`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opentelekomcloud_mrs_job_v1",
			Category:         "MRS Resources",
			ShortDescription: `Manages resource job within OpenTelekomCloud MRS.`,
			Description:      ``,
			Keywords: []string{
				"mrs",
				"job",
				"v1",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "job_type",
					Description: `(Required) Job type 1: MapReduce 2: Spark 3: Hive Script 4: HiveQL (not supported currently) 5: DistCp, importing and exporting data. 6: Spark Script 7: Spark SQL, submitting Spark SQL statements. (not supported in this APIcurrently) NOTE: Spark and Hive jobs can be added to only clusters including Spark and Hive components.`,
				},
				resource.Attribute{
					Name:        "job_name",
					Description: `(Required) Job name Contains only 1 to 64 letters, digits, hyphens (-), and underscores (_). NOTE: Identical job names are allowed but not recommended.`,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Required) Cluster ID`,
				},
				resource.Attribute{
					Name:        "jar_path",
					Description: `(Required) Path of the .jar package or .sql file for program execution The parameter must meet the following requirements: Contains a maximum of 1023 characters, excluding special characters such as ;|&><'$. The address cannot be empty or full of spaces. Starts with / or s3a://. Spark Script must end with .sql; while MapReduce and Spark Jar must end with .jar. sql and jar are case-insensitive.`,
				},
				resource.Attribute{
					Name:        "arguments",
					Description: `(Optional) Key parameter for program execution. The parameter is specified by the function of the user's program. MRS is only responsible for loading the parameter. The parameter contains a maximum of 2047 characters, excluding special characters such as ;|&>'<$, and can be empty.`,
				},
				resource.Attribute{
					Name:        "input",
					Description: `(Optional) Path for inputting data, which must start with / or s3a://. A correct OBS path is required. The parameter contains a maximum of 1023 characters, excluding special characters such as ;|&>'<$, and can be empty.`,
				},
				resource.Attribute{
					Name:        "output",
					Description: `(Optional) Path for outputting data, which must start with / or s3a://. A correct OBS path is required. If the path does not exist, the system automatically creates it. The parameter contains a maximum of 1023 characters, excluding special characters such as ;|&>'<$, and can be empty.`,
				},
				resource.Attribute{
					Name:        "job_log",
					Description: `(Optional) Path for storing job logs that record job running status. This path must start with / or s3a://. A correct OBS path is required. The parameter contains a maximum of 1023 characters, excluding special characters such as ;|&>'<$, and can be empty.`,
				},
				resource.Attribute{
					Name:        "hive_script_path",
					Description: `(Optional) SQL program path This parameter is needed by Spark Script and Hive Script jobs only and must meet the following requirements: Contains a maximum of 1023 characters, excluding special characters such as ;|&><'$. The address cannot be empty or full of spaces. Starts with / or s3a://. Ends with .sql. sql is case-insensitive.`,
				},
				resource.Attribute{
					Name:        "is_protected",
					Description: `(Optional) Whether a job is protected true false The current version does not support this function.`,
				},
				resource.Attribute{
					Name:        "is_public",
					Description: `(Optional) Whether a job is public true false The current version does not support this function. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "job_type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "job_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "jar_path",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "arguments",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "input",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "output",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "job_log",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "hive_script_path",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "is_protected",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "is_public",
					Description: `See Argument Reference above.`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "job_type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "job_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "jar_path",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "arguments",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "input",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "output",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "job_log",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "hive_script_path",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "is_protected",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "is_public",
					Description: `See Argument Reference above.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opentelekomcloud_nat_gateway_v2",
			Category:         "Nat Resources",
			ShortDescription: `Manages a V2 nat gateway resource within OpenTelekomCloud Nat.`,
			Description:      ``,
			Keywords: []string{
				"nat",
				"gateway",
				"v2",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the nat gateway.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description of the nat gateway.`,
				},
				resource.Attribute{
					Name:        "spec",
					Description: `(Required) The specification of the nat gateway, valid values are "1", "2", "3", "4".`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Optional) The target tenant ID in which to allocate the nat gateway. Changing this creates a new nat gateway.`,
				},
				resource.Attribute{
					Name:        "router_id",
					Description: `(Required) ID of the router this nat gateway belongs to. Changing this creates a new nat gateway.`,
				},
				resource.Attribute{
					Name:        "internal_network_id",
					Description: `(Required) ID of the network this nat gateway connects to. Changing this creates a new nat gateway. ## Attributes Reference The following attributes are exported:`,
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
					Name:        "spec",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "router_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "internal_network_id",
					Description: `See Argument Reference above.`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "spec",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "router_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "internal_network_id",
					Description: `See Argument Reference above.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opentelekomcloud_nat_snat_rule_v2",
			Category:         "Nat Resources",
			ShortDescription: `Manages a V2 snat rule resource within OpenTelekomCloud Nat.`,
			Description:      ``,
			Keywords: []string{
				"nat",
				"snat",
				"rule",
				"v2",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "nat_gateway_id",
					Description: `(Required) ID of the nat gateway this snat rule belongs to. Changing this creates a new snat rule.`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `(Optional) ID of the network this snat rule connects to. This parameter and cidr are alternative. Changing this creates a new snat rule.`,
				},
				resource.Attribute{
					Name:        "source_type",
					Description: `(Optional) 0: Either network_id or cidr can be specified in a VPC. 1: Only cidr can be specified over a dedicated network. Changing this creates a new snat rule.`,
				},
				resource.Attribute{
					Name:        "cidr",
					Description: `(Optional) Specifies CIDR, which can be in the format of a network segment or a host IP address. This parameter and network_id are alternative. If the value of source_type is 0, the CIDR block must be a subset of the VPC subnet CIDR block. If the value of source_type is 1, the CIDR block must be a CIDR block of Direct Connect and cannot conflict with the VPC CIDR blocks. Changing this creates a new snat rule.`,
				},
				resource.Attribute{
					Name:        "floating_ip_id",
					Description: `(Required) ID of the floating ip this snat rule connets to. Changing this creates a new snat rule. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "nat_gateway_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "floating_ip_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "source_type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "cidr",
					Description: `See Argument Reference above.`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "nat_gateway_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "floating_ip_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "source_type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "cidr",
					Description: `See Argument Reference above.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opentelekomcloud_networking_floatingip_associate_v2",
			Category:         "Networking Resources",
			ShortDescription: `Associates a Floating IP to a Port`,
			Description:      ``,
			Keywords: []string{
				"networking",
				"floatingip",
				"associate",
				"v2",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "floating_ip",
					Description: `(Required) IP Address of an existing floating IP.`,
				},
				resource.Attribute{
					Name:        "port_id",
					Description: `(Required) ID of an existing port with at least one IP address to associate with this floating IP. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "floating_ip",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "port_id",
					Description: `See Argument Reference above. ## Import Floating IP associations can be imported using the ` + "`" + `id` + "`" + ` of the floating IP, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import opentelekomcloud_networking_floatingip_associate_v2.fip 2c7f39f3-702b-48d1-940c-b50384177ee1 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "floating_ip",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "port_id",
					Description: `See Argument Reference above. ## Import Floating IP associations can be imported using the ` + "`" + `id` + "`" + ` of the floating IP, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import opentelekomcloud_networking_floatingip_associate_v2.fip 2c7f39f3-702b-48d1-940c-b50384177ee1 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opentelekomcloud_networking_floatingip_v2",
			Category:         "Networking Resources",
			ShortDescription: `Manages a V2 floating IP resource within OpenTelekomCloud Neutron (networking).`,
			Description:      ``,
			Keywords: []string{
				"networking",
				"floatingip",
				"v2",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "pool",
					Description: `(Required) The name of the pool from which to obtain the floating IP. Changing this creates a new floating IP.`,
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
					Description: `Fixed IP of the port to associate with this floating IP. Required if the port has multiple fixed IPs.`,
				},
				resource.Attribute{
					Name:        "value_specs",
					Description: `(Optional) Map of additional options. ## Attributes Reference The following attributes are exported:`,
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
					Description: `The fixed IP which the floating IP maps to. ## Import Floating IPs can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import opentelekomcloud_networking_floatingip_v2.floatip_1 2c7f39f3-702b-48d1-940c-b50384177ee1 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{
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
					Description: `The fixed IP which the floating IP maps to. ## Import Floating IPs can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import opentelekomcloud_networking_floatingip_v2.floatip_1 2c7f39f3-702b-48d1-940c-b50384177ee1 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opentelekomcloud_networking_network_v2",
			Category:         "Networking Resources",
			ShortDescription: `Manages a V2 Neutron network resource within OpenTelekomCloud.`,
			Description:      ``,
			Keywords: []string{
				"networking",
				"network",
				"v2",
			},
			Arguments: []resource.Argument{
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
					Description: `(Optional) The administrative state of the network. Acceptable values are "true" and "false". Changing this value updates the state of the existing network.`,
				},
				resource.Attribute{
					Name:        "segments",
					Description: `(Optional) An array of one or more provider segment objects.`,
				},
				resource.Attribute{
					Name:        "value_specs",
					Description: `(Optional) Map of additional options. The ` + "`" + `segments` + "`" + ` block supports:`,
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
					Description: `See Argument Reference above. ## Import Networks can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import opentelekomcloud_networking_network_v2.network_1 d90ce693-5ccf-4136-a0ed-152ce412b6b9 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{
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
					Description: `See Argument Reference above. ## Import Networks can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import opentelekomcloud_networking_network_v2.network_1 d90ce693-5ccf-4136-a0ed-152ce412b6b9 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opentelekomcloud_networking_port_v2",
			Category:         "Networking Resources",
			ShortDescription: `Manages a V2 port resource within OpenTelekomCloud.`,
			Description:      ``,
			Keywords: []string{
				"networking",
				"port",
				"v2",
			},
			Arguments: []resource.Argument{
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
					Description: `(Optional) A list of security group IDs to apply to the port. The security groups must be specified by ID and not name (as opposed to how they are configured with the Compute Instance).`,
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
					Description: `(Optional) An IP/MAC Address pair of additional IP addresses that can be active on this port. The structure is described below.`,
				},
				resource.Attribute{
					Name:        "value_specs",
					Description: `(Optional) Map of additional options. The ` + "`" + `fixed_ip` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Required) Subnet in which to allocate IP address for this port.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `(Optional) IP address desired in the subnet for this port. If you don't specify ` + "`" + `ip_address` + "`" + `, an available IP address from the specified subnet will be allocated to this port. The ` + "`" + `allowed_address_pairs` + "`" + ` block supports:`,
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
					Name:        "all fixed_ips",
					Description: `The collection of Fixed IP addresses on the port in the order returned by the Network v2 API. ## Import Ports can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import opentelekomcloud_networking_port_v2.port_1 eae26a3e-1c33-4cc1-9c31-0cd729c438a1 ` + "`" + `` + "`" + `` + "`" + ` ## Notes ### Ports and Instances There are some notes to consider when connecting Instances to networks using Ports. Please see the ` + "`" + `opentelekomcloud_compute_instance_v2` + "`" + ` documentation for further documentation.`,
				},
			},
			Attributes: []resource.Argument{
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
					Name:        "all fixed_ips",
					Description: `The collection of Fixed IP addresses on the port in the order returned by the Network v2 API. ## Import Ports can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import opentelekomcloud_networking_port_v2.port_1 eae26a3e-1c33-4cc1-9c31-0cd729c438a1 ` + "`" + `` + "`" + `` + "`" + ` ## Notes ### Ports and Instances There are some notes to consider when connecting Instances to networks using Ports. Please see the ` + "`" + `opentelekomcloud_compute_instance_v2` + "`" + ` documentation for further documentation.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opentelekomcloud_networking_router_interface_v2",
			Category:         "Networking Resources",
			ShortDescription: `Manages a V2 router interface resource within OpenTelekomCloud.`,
			Description:      ``,
			Keywords: []string{
				"networking",
				"router",
				"interface",
				"v2",
			},
			Arguments: []resource.Argument{
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
					Name:        "router_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "port_id",
					Description: `See Argument Reference above.`,
				},
			},
			Attributes: []resource.Argument{
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
					Description: `See Argument Reference above.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opentelekomcloud_networking_router_route_v2",
			Category:         "Networking Resources",
			ShortDescription: `Creates a routing entry on a OpenTelekomCloud V2 router.`,
			Description:      ``,
			Keywords: []string{
				"networking",
				"router",
				"route",
				"v2",
			},
			Arguments: []resource.Argument{
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
					Name:        "router_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "destination_cidr",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "next_hop",
					Description: `See Argument Reference above. ## Notes The ` + "`" + `next_hop` + "`" + ` IP address must be directly reachable from the router at the ` + "`" + `` + "`" + `opentelekomcloud_networking_router_route_v2` + "`" + `` + "`" + ` resource creation time. You can ensure that by explicitly specifying a dependency on the ` + "`" + `` + "`" + `opentelekomcloud_networking_router_interface_v2` + "`" + `` + "`" + ` resource that connects the next hop to the router, as in the example above.`,
				},
			},
			Attributes: []resource.Argument{
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
					Description: `See Argument Reference above. ## Notes The ` + "`" + `next_hop` + "`" + ` IP address must be directly reachable from the router at the ` + "`" + `` + "`" + `opentelekomcloud_networking_router_route_v2` + "`" + `` + "`" + ` resource creation time. You can ensure that by explicitly specifying a dependency on the ` + "`" + `` + "`" + `opentelekomcloud_networking_router_interface_v2` + "`" + `` + "`" + ` resource that connects the next hop to the router, as in the example above.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opentelekomcloud_networking_router_v2",
			Category:         "Networking Resources",
			ShortDescription: `Manages a V2 router resource within OpenTelekomCloud. The router is the top-level resource for the VPC within OpenTelekomCloud.`,
			Description:      ``,
			Keywords: []string{
				"networking",
				"router",
				"v2",
			},
			Arguments: []resource.Argument{
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
					Name:        "external_gateway",
					Description: `(Optional) The network UUID of an external gateway for the router. A router with an external gateway is required if any compute instances or load balancers will be using floating IPs. Changing this updates the ` + "`" + `external_gateway` + "`" + ` of an existing router.`,
				},
				resource.Attribute{
					Name:        "enable_snat",
					Description: `(Optional) Enable Source NAT for the router. Valid values are "true" or "false". An ` + "`" + `external_gateway` + "`" + ` has to be set in order to set this property. Changing this updates the ` + "`" + `enable_snat` + "`" + ` of the router.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Optional) The owner of the floating IP. Required if admin wants to create a router for another tenant. Changing this creates a new router.`,
				},
				resource.Attribute{
					Name:        "value_specs",
					Description: `(Optional) Map of additional driver-specific options. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the router.`,
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
					Name:        "external_gateway",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "enable_snat",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "value_specs",
					Description: `See Argument Reference above.`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the router.`,
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
					Name:        "external_gateway",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "enable_snat",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "value_specs",
					Description: `See Argument Reference above.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opentelekomcloud_networking_secgroup_rule_v2",
			Category:         "Networking Resources",
			ShortDescription: `Manages a V2 Neutron security group rule resource within OpenTelekomCloud.`,
			Description:      ``,
			Keywords: []string{
				"networking",
				"secgroup",
				"rule",
				"v2",
			},
			Arguments: []resource.Argument{
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
					Description: `(Optional) The remote group id, the value needs to be an OpenTelekomCloud ID of a security group in the same tenant. Changing this creates a new security group rule.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `(Required) The security group id the rule should belong to, the value needs to be an OpenTelekomCloud ID of a security group in the same tenant. Changing this creates a new security group rule.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Optional) The owner of the security group. Required if admin wants to create a port for another tenant. Changing this creates a new security group rule. ## Attributes Reference The following attributes are exported:`,
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
					Description: `See Argument Reference above. ## Import Security Group Rules can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import opentelekomcloud_networking_secgroup_rule_v2.secgroup_rule_1 aeb68ee3-6e9d-4256-955c-9584a6212745 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{
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
					Description: `See Argument Reference above. ## Import Security Group Rules can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import opentelekomcloud_networking_secgroup_rule_v2.secgroup_rule_1 aeb68ee3-6e9d-4256-955c-9584a6212745 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opentelekomcloud_networking_secgroup_v2",
			Category:         "Networking Resources",
			ShortDescription: `Manages a V2 Neutron security group resource within OpenTelekomCloud.`,
			Description:      ``,
			Keywords: []string{
				"networking",
				"secgroup",
				"v2",
			},
			Arguments: []resource.Argument{
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
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above. ## Default Security Group Rules In most cases, OpenTelekomCloud will create some egress security group rules for each new security group. These security group rules will not be managed by Terraform, so if you prefer to have`,
				},
			},
			Attributes: []resource.Argument{
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
					Description: `See Argument Reference above. ## Default Security Group Rules In most cases, OpenTelekomCloud will create some egress security group rules for each new security group. These security group rules will not be managed by Terraform, so if you prefer to have`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opentelekomcloud_networking_subnet_v2",
			Category:         "Networking Resources",
			ShortDescription: `Manages a V2 Neutron subnet resource within OpenTelekomCloud.`,
			Description:      ``,
			Keywords: []string{
				"networking",
				"subnet",
				"v2",
			},
			Arguments: []resource.Argument{
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
					Description: `(Optional) The administrative state of the network. Acceptable values are "true" and "false". Changing this value enables or disables the DHCP capabilities of the existing subnet. Defaults to true.`,
				},
				resource.Attribute{
					Name:        "dns_nameservers",
					Description: `(Optional) An array of DNS name server names used by hosts in this subnet. Changing this updates the DNS name servers for the existing subnet.`,
				},
				resource.Attribute{
					Name:        "host_routes",
					Description: `(Optional) An array of routes that should be used by devices with IPs from this subnet (not including local subnet route). The host_route object structure is documented below. Changing this updates the host routes for the existing subnet.`,
				},
				resource.Attribute{
					Name:        "value_specs",
					Description: `(Optional) Map of additional options. The ` + "`" + `allocation_pools` + "`" + ` block supports:`,
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
					Description: `See Argument Reference above. ## Import Subnets can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import opentelekomcloud_networking_subnet_v2.subnet_1 da4faf16-5546-41e4-8330-4d0002b74048 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{
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
					Description: `See Argument Reference above. ## Import Subnets can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import opentelekomcloud_networking_subnet_v2.subnet_1 da4faf16-5546-41e4-8330-4d0002b74048 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opentelekomcloud_networking_vip_associate_v2",
			Category:         "Networking Resources",
			ShortDescription: `Manages a V2 vip associate resource within OpenTelekomCloud.`,
			Description:      ``,
			Keywords: []string{
				"networking",
				"vip",
				"associate",
				"v2",
			},
			Arguments: []resource.Argument{
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
			Attributes: []resource.Argument{
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
			Type:             "opentelekomcloud_networking_vip_v2",
			Category:         "Networking Resources",
			ShortDescription: `Manages a V2 vip resource within OpenTelekomCloud.`,
			Description:      ``,
			Keywords: []string{
				"networking",
				"vip",
				"v2",
			},
			Arguments: []resource.Argument{
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
			Attributes: []resource.Argument{
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
		&resource.Resource{
			Name:             "",
			Type:             "opentelekomcloud_rds_instance_v1",
			Category:         "DB Instance Resources",
			ShortDescription: `Manages rds instance resource within OpenTelekomCloud`,
			Description:      ``,
			Keywords: []string{
				"db",
				"instance",
				"rds",
				"v1",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Specifies the DB instance name. The DB instance name of the same type is unique in the same tenant. The changes of the instance name will be suppressed in HA scenario.`,
				},
				resource.Attribute{
					Name:        "datastore",
					Description: `(Required) Specifies database information. The structure is described below.`,
				},
				resource.Attribute{
					Name:        "flavorref",
					Description: `(Required) Specifies the specification ID (flavors.id in the response message in Obtaining All DB Instance Specifications). If you want to enable ha for the rds instance, a flavor with ha speccode is required.`,
				},
				resource.Attribute{
					Name:        "volume",
					Description: `(Required) Specifies the volume information. The structure is described below.`,
				},
				resource.Attribute{
					Name:        "availabilityzone",
					Description: `(Required) Specifies the ID of the AZ.`,
				},
				resource.Attribute{
					Name:        "vpc",
					Description: `(Required) Specifies the VPC ID. For details about how to obtain this parameter value, see section "Virtual Private Cloud" in the Virtual Private Cloud API Reference.`,
				},
				resource.Attribute{
					Name:        "nics",
					Description: `(Required) Specifies the nics information. For details about how to obtain this parameter value, see section "Subnet" in the Virtual Private Cloud API Reference. The structure is described below.`,
				},
				resource.Attribute{
					Name:        "securitygroup",
					Description: `(Required) Specifies the security group which the RDS DB instance belongs to. The structure is described below.`,
				},
				resource.Attribute{
					Name:        "dbport",
					Description: `(Optional) Specifies the database port number.`,
				},
				resource.Attribute{
					Name:        "backupstrategy",
					Description: `(Optional) Specifies the advanced backup policy. The structure is described below.`,
				},
				resource.Attribute{
					Name:        "dbrtpd",
					Description: `(Required) Specifies the password for user root of the database.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `(Optional) Tags key/value pairs to associate with the instance.`,
				},
				resource.Attribute{
					Name:        "ha",
					Description: `(Optional) Specifies the parameters configured on HA and is used when creating HA DB instances. The structure is described below. NOTICE: RDS for Microsoft SQL Server does not support creating HA DB instances and this parameter is not involved. The ` + "`" + `datastore` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) Specifies the DB engine. Currently, MySQL, PostgreSQL, and Microsoft SQL Server are supported. The value is MySQL, PostgreSQL, or SQLServer.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Required) Specifies the DB instance version.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) Specifies the volume type. Valid value: It must be COMMON (SATA) or ULTRAHIGH (SSD) and is case-sensitive.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `(Required) Specifies the volume size. Its value must be a multiple of 10 and the value range is 100 GB to 2000 GB. The ` + "`" + `nics` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "subnetId",
					Description: `(Required) Specifies the subnet ID obtained from the VPC. The ` + "`" + `securitygroup ` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Required) Specifies the ID obtained from the securitygroup. The ` + "`" + `backupstrategy ` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "starttime",
					Description: `(Optional) Indicates the backup start time that has been set. The backup task will be triggered within one hour after the backup start time. Valid value: The value cannot be empty. It must use the hh:mm:ss format and must be valid. The current time is the UTC time.`,
				},
				resource.Attribute{
					Name:        "keepdays",
					Description: `(Optional) Specifies the number of days to retain the generated backup files. Its value range is 0 to 35. If this parameter is not specified or set to 0, the automated backup policy is disabled. The ` + "`" + `ha` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "enable",
					Description: `(Optional) Specifies the configured parameters on the HA. Valid value: The value is true or false. The value true indicates creating HA DB instances. The value false indicates creating a single DB instance.`,
				},
				resource.Attribute{
					Name:        "replicationmode",
					Description: `(Optional) Specifies the replication mode for the standby DB instance. The value cannot be empty. For MySQL, the value is async or semisync. For PostgreSQL, the value is async or sync. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "flavorref",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "volume",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "availabilityzone",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "vpc",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "nics",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "securitygroup",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "dbport",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "backupstrategy",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "dbrtpd",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "ha",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Indicates the DB instance status.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `Indicates the instance connection address. It is a blank string.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Indicates the DB instance type, which can be master or readreplica.`,
				},
				resource.Attribute{
					Name:        "created",
					Description: `Indicates the creation time in the following format: yyyy-mm-dd Thh:mm:ssZ.`,
				},
				resource.Attribute{
					Name:        "updated",
					Description: `Indicates the update time in the following format: yyyy-mm-dd Thh:mm:ssZ. ## Attributes Reference The following attributes can be updated:`,
				},
				resource.Attribute{
					Name:        "volume.size",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "flavorref",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "backupstrategy",
					Description: `See Argument Reference above.`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "flavorref",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "volume",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "availabilityzone",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "vpc",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "nics",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "securitygroup",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "dbport",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "backupstrategy",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "dbrtpd",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "ha",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Indicates the DB instance status.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `Indicates the instance connection address. It is a blank string.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Indicates the DB instance type, which can be master or readreplica.`,
				},
				resource.Attribute{
					Name:        "created",
					Description: `Indicates the creation time in the following format: yyyy-mm-dd Thh:mm:ssZ.`,
				},
				resource.Attribute{
					Name:        "updated",
					Description: `Indicates the update time in the following format: yyyy-mm-dd Thh:mm:ssZ. ## Attributes Reference The following attributes can be updated:`,
				},
				resource.Attribute{
					Name:        "volume.size",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "flavorref",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "backupstrategy",
					Description: `See Argument Reference above.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opentelekomcloud_rds_instance_v3",
			Category:         "DB Instance Resources",
			ShortDescription: `instance management`,
			Description:      ``,
			Keywords: []string{
				"db",
				"instance",
				"rds",
				"v3",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "availability_zone",
					Description: `(Required) Specifies the AZ name. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "db",
					Description: `(Required) Specifies the database information. Structure is documented below. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "flavor",
					Description: `(Required) Specifies the specification code. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Specifies the DB instance name. The DB instance name of the same type must be unique for the same tenant. The value must be 4 to 64 characters in length and start with a letter. It is case-sensitive and can contain only letters, digits, hyphens (-), and underscores (_). Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `(Required) Specifies the security group which the RDS DB instance belongs to. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Required) Specifies the subnet id. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "volume",
					Description: `(Required) Specifies the volume information. Structure is documented below. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required) Specifies the VPC ID. Changing this parameter will create a new resource. The ` + "`" + `db` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Required) Specifies the database password. The value cannot be empty and should contain 8 to 32 characters, including uppercase and lowercase letters, digits, and the following special characters: ~!@#%^`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional) Specifies the database port information. The MySQL database port ranges from 1024 to 65535 (excluding 12017 and 33071, which are occupied by the RDS system and cannot be used). The PostgreSQL database port ranges from 2100 to 9500. The Microsoft SQL Server database port can be 1433 or ranges from 2100 to 9500, excluding 5355 and 5985. If this parameter is not set, the default value is as follows: For MySQL, the default value is 3306. For PostgreSQL, the default value is 5432. For Microsoft SQL Server, the default value is 1433. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) Specifies the DB engine. Value: MySQL, PostgreSQL, SQLServer. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "user_name",
					Description: `Indicates the default user name of database.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Required) Specifies the database version. MySQL databases support MySQL 5.6 and 5.7. PostgreSQL databases support PostgreSQL 9.5 and 9.6. Microsoft SQL Server databases support 2014 SE, 2016 SE, and 2016 EE. Changing this parameter will create a new resource. The ` + "`" + `volume` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "disk_encryption_id",
					Description: `(Optional) Specifies the key ID for disk encryption. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `(Required) Specifies the volume size. Its value range is from 40 GB to 4000 GB. The value must be a multiple of 10. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) Specifies the volume type. Its value can be any of the following and is case-sensitive: COMMON: indicates the SATA type. ULTRAHIGH: indicates the SSD type. Changing this parameter will create a new resource. - - -`,
				},
				resource.Attribute{
					Name:        "backup_strategy",
					Description: `(Optional) Specifies the advanced backup policy. Structure is documented below. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "ha_replication_mode",
					Description: `(Optional) Specifies the replication mode for the standby DB instance. For MySQL, the value is async or semisync. For PostgreSQL, the value is async or sync. For Microsoft SQL Server, the value is sync. NOTE: async indicates the asynchronous replication mode. semisync indicates the semi-synchronous replication mode. sync indicates the synchronous replication mode. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "param_group_id",
					Description: `(Optional) Specifies the parameter group ID. Changing this parameter will create a new resource. The ` + "`" + `backup_strategy` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "keep_days",
					Description: `(Optional) Specifies the retention days for specific backup files. The value range is from 0 to 732. If this parameter is not specified or set to 0, the automated backup policy is disabled. NOTICE: Primary/standby DB instances of Microsoft SQL Server do not support disabling the automated backup policy. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "start_time",
					Description: `(Required) Specifies the backup time window. Automated backups will be triggered during the backup time window. It must be a valid value in the &quot;hh:mm-HH:MM&quot; format. The current time is in the UTC format. The HH value must be 1 greater than the hh value. The values of mm and MM must be the same and must be set to any of the following: 00, 15, 30, or 45. Example value: 08:15-09:15 23:00-00:00. Changing this parameter will create a new resource. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "created",
					Description: `Indicates the creation time.`,
				},
				resource.Attribute{
					Name:        "nodes",
					Description: `Indicates the instance nodes information. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "private_ips",
					Description: `Indicates the private IP address list. It is a blank string until an ECS is created.`,
				},
				resource.Attribute{
					Name:        "public_ips",
					Description: `Indicates the public IP address list. The ` + "`" + `nodes` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `Indicates the AZ.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Indicates the node ID.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Indicates the node name.`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `Indicates the node type. The value can be master or slave, indicating the primary node or standby node respectively.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Indicates the node status. ## Timeouts This resource provides the following timeouts configuration options: - ` + "`" + `create` + "`" + ` - Default is 30 minute. ## Import RDS instance can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import opentelekomcloud_rds_instance_v3.instance_1 7117d38e-4c8f-4624-a505-bd96b97d024c ` + "`" + `` + "`" + `` + "`" + ` But due to some attrubutes missing from the API response, it's required to ignore changes as below. ` + "`" + `` + "`" + `` + "`" + ` resource "opentelekomcloud_rds_instance_v3" "instance_1" { ... lifecycle { ignore_changes = [ "db", ] } } ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "created",
					Description: `Indicates the creation time.`,
				},
				resource.Attribute{
					Name:        "nodes",
					Description: `Indicates the instance nodes information. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "private_ips",
					Description: `Indicates the private IP address list. It is a blank string until an ECS is created.`,
				},
				resource.Attribute{
					Name:        "public_ips",
					Description: `Indicates the public IP address list. The ` + "`" + `nodes` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `Indicates the AZ.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Indicates the node ID.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Indicates the node name.`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `Indicates the node type. The value can be master or slave, indicating the primary node or standby node respectively.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Indicates the node status. ## Timeouts This resource provides the following timeouts configuration options: - ` + "`" + `create` + "`" + ` - Default is 30 minute. ## Import RDS instance can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import opentelekomcloud_rds_instance_v3.instance_1 7117d38e-4c8f-4624-a505-bd96b97d024c ` + "`" + `` + "`" + `` + "`" + ` But due to some attrubutes missing from the API response, it's required to ignore changes as below. ` + "`" + `` + "`" + `` + "`" + ` resource "opentelekomcloud_rds_instance_v3" "instance_1" { ... lifecycle { ignore_changes = [ "db", ] } } ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opentelekomcloud_rds_parametergroup_v3",
			Category:         "DB Instance Resources",
			ShortDescription: `Manages a V3 RDS parametergroup resource within OpenTelekomCloud.`,
			Description:      ``,
			Keywords: []string{
				"db",
				"instance",
				"rds",
				"parametergroup",
				"v3",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The parameter group name. It contains a maximum of 64 characters.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The parameter group description. It contains a maximum of 256 characters and cannot contain the following special characters:>!<"&'= the value is left blank by default.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Optional) Parameter group values key/value pairs defined by users based on the default parameter groups.`,
				},
				resource.Attribute{
					Name:        "datastore",
					Description: `(Required) Database object. The database object structure is documented below. Changing this creates a new parameter group. The ` + "`" + `datastore` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) The DB engine. Currently, MySQL, PostgreSQL, and Microsoft SQL Server are supported. The value is case-insensitive and can be mysql, postgresql, or sqlserver.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Required) Specifies the database version.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the parameter group.`,
				},
				resource.Attribute{
					Name:        "configuration_parameters",
					Description: `Indicates the parameter configuration defined by users based on the default parameters groups.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Indicates the parameter name.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `Indicates the parameter value.`,
				},
				resource.Attribute{
					Name:        "restart_required",
					Description: `Indicates whether a restart is required.`,
				},
				resource.Attribute{
					Name:        "readonly",
					Description: `Indicates whether the parameter is read-only.`,
				},
				resource.Attribute{
					Name:        "value_range",
					Description: `Indicates the parameter value range.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Indicates the parameter type.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Indicates the parameter description. ## Import Parameter groups can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import opentelekomcloud_rds_parametergroup_v3.pg_1 7117d38e-4c8f-4624-a505-bd96b97d024c ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the parameter group.`,
				},
				resource.Attribute{
					Name:        "configuration_parameters",
					Description: `Indicates the parameter configuration defined by users based on the default parameters groups.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Indicates the parameter name.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `Indicates the parameter value.`,
				},
				resource.Attribute{
					Name:        "restart_required",
					Description: `Indicates whether a restart is required.`,
				},
				resource.Attribute{
					Name:        "readonly",
					Description: `Indicates whether the parameter is read-only.`,
				},
				resource.Attribute{
					Name:        "value_range",
					Description: `Indicates the parameter value range.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Indicates the parameter type.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Indicates the parameter description. ## Import Parameter groups can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import opentelekomcloud_rds_parametergroup_v3.pg_1 7117d38e-4c8f-4624-a505-bd96b97d024c ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opentelekomcloud_rts_software_config_v1",
			Category:         "RTS Resources",
			ShortDescription: `Provides an RTS software config resource.`,
			Description:      ``,
			Keywords: []string{
				"rts",
				"software",
				"config",
				"v1",
			},
			Arguments:  []resource.Argument{},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opentelekomcloud_rts_software_deployment_v1",
			Category:         "RTS Resources",
			ShortDescription: `Provides an RTS software deployment resource.`,
			Description:      ``,
			Keywords: []string{
				"rts",
				"software",
				"deployment",
				"v1",
			},
			Arguments:  []resource.Argument{},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opentelekomcloud-resource-rts-stack-v1",
			Category:         "RTS Resources",
			ShortDescription: `Provides an OpenTelekomCloud RTS Stack.`,
			Description:      ``,
			Keywords: []string{
				"rts",
				"opentelekomcloud-resource-rts-stack-v1",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A unique name for the stack. The value must meet the regular expression rule (` + "`" + `^[a-zA-Z][a-zA-Z0-9_.-]{0,254}$` + "`" + `). Changing this creates a new stack.`,
				},
				resource.Attribute{
					Name:        "template_body",
					Description: `(Optional; Required if ` + "`" + `template_url` + "`" + ` is empty) Structure containing the template body. The template content must use the yaml syntax.`,
				},
				resource.Attribute{
					Name:        "template_url",
					Description: `(Optional; Required if ` + "`" + `template_body` + "`" + ` is empty) Location of a file containing the template body.`,
				},
				resource.Attribute{
					Name:        "environment",
					Description: `(Optional) Tthe environment information about the stack.`,
				},
				resource.Attribute{
					Name:        "files",
					Description: `(Optional) Files used in the environment.`,
				},
				resource.Attribute{
					Name:        "parameters",
					Description: `(Optional) A list of Parameter structures that specify input parameters for the stack.`,
				},
				resource.Attribute{
					Name:        "disable_rollback",
					Description: `(Optional) Set to true to disable rollback of the stack if stack creation failed.`,
				},
				resource.Attribute{
					Name:        "timeout_mins",
					Description: `(Optional) Specifies the timeout duration. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "outputs",
					Description: `A map of outputs from the stack.`,
				},
				resource.Attribute{
					Name:        "capabilities",
					Description: `List of stack capabilities for stack.`,
				},
				resource.Attribute{
					Name:        "notification_topics",
					Description: `List of notification topics for stack.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Specifies the stack status. ## Import RTS Stacks can be imported using the ` + "`" + `name` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import opentelekomcloud_rts_stack_v1.mystack rts-stack ` + "`" + `` + "`" + `` + "`" + ` <a id="timeouts"></a> ## Timeouts ` + "`" + `opentelekomcloud_rts_stack_v1` + "`" + ` provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - (Default ` + "`" + `30 minutes` + "`" + `) Used for Creating Stacks - ` + "`" + `update` + "`" + ` - (Default ` + "`" + `30 minutes` + "`" + `) Used for Stack modifications - ` + "`" + `delete` + "`" + ` - (Default ` + "`" + `30 minutes` + "`" + `) Used for destroying stacks.`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "outputs",
					Description: `A map of outputs from the stack.`,
				},
				resource.Attribute{
					Name:        "capabilities",
					Description: `List of stack capabilities for stack.`,
				},
				resource.Attribute{
					Name:        "notification_topics",
					Description: `List of notification topics for stack.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Specifies the stack status. ## Import RTS Stacks can be imported using the ` + "`" + `name` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import opentelekomcloud_rts_stack_v1.mystack rts-stack ` + "`" + `` + "`" + `` + "`" + ` <a id="timeouts"></a> ## Timeouts ` + "`" + `opentelekomcloud_rts_stack_v1` + "`" + ` provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - (Default ` + "`" + `30 minutes` + "`" + `) Used for Creating Stacks - ` + "`" + `update` + "`" + ` - (Default ` + "`" + `30 minutes` + "`" + `) Used for Stack modifications - ` + "`" + `delete` + "`" + ` - (Default ` + "`" + `30 minutes` + "`" + `) Used for destroying stacks.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opentelekomcloud_s3_bucket",
			Category:         "S3 Resource",
			ShortDescription: `Provides a S3 bucket resource.`,
			Description:      ``,
			Keywords: []string{
				"s3",
				"resource",
				"bucket",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "bucket",
					Description: `(Optional, Forces new resource) The name of the bucket. If omitted, Terraform will assign a random, unique name.`,
				},
				resource.Attribute{
					Name:        "bucket_prefix",
					Description: `(Optional, Forces new resource) Creates a unique bucket name beginning with the specified prefix. Conflicts with ` + "`" + `bucket` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "acl",
					Description: `(Optional) The [canned ACL](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl) to apply. Defaults to "private".`,
				},
				resource.Attribute{
					Name:        "policy",
					Description: `(Optional) A valid [bucket policy](https://docs.aws.amazon.com/AmazonS3/latest/dev/example-bucket-policies.html) JSON document. Note that if the policy document is not specific enough (but still valid), Terraform may view the policy as constantly changing in a ` + "`" + `terraform plan` + "`" + `. In this case, please make sure you use the verbose/specific version of the policy.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A mapping of tags to assign to the bucket.`,
				},
				resource.Attribute{
					Name:        "force_destroy",
					Description: `(Optional, Default:false ) A boolean that indicates all objects should be deleted from the bucket so that the bucket can be destroyed without error. These objects are`,
				},
				resource.Attribute{
					Name:        "website",
					Description: `(Optional) A website object (documented below).`,
				},
				resource.Attribute{
					Name:        "cors_rule",
					Description: `(Optional) A rule of [Cross-Origin Resource Sharing](https://docs.aws.amazon.com/AmazonS3/latest/dev/cors.html) (documented below).`,
				},
				resource.Attribute{
					Name:        "versioning",
					Description: `(Optional) A state of [versioning](https://docs.aws.amazon.com/AmazonS3/latest/dev/Versioning.html) (documented below)`,
				},
				resource.Attribute{
					Name:        "logging",
					Description: `(Optional) A settings of [bucket logging](https://docs.aws.amazon.com/AmazonS3/latest/UG/ManagingBucketLogging.html) (documented below).`,
				},
				resource.Attribute{
					Name:        "lifecycle_rule",
					Description: `(Optional) A configuration of [object lifecycle management](http://docs.aws.amazon.com/AmazonS3/latest/dev/object-lifecycle-mgmt.html) (documented below). The ` + "`" + `website` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "index_document",
					Description: `(Required, unless using ` + "`" + `redirect_all_requests_to` + "`" + `) Amazon S3 returns this index document when requests are made to the root domain or any of the subfolders.`,
				},
				resource.Attribute{
					Name:        "error_document",
					Description: `(Optional) An absolute path to the document to return in case of a 4XX error.`,
				},
				resource.Attribute{
					Name:        "redirect_all_requests_to",
					Description: `(Optional) A hostname to redirect all website requests for this bucket to. Hostname can optionally be prefixed with a protocol (` + "`" + `http://` + "`" + ` or ` + "`" + `https://` + "`" + `) to use when redirecting requests. The default is the protocol that is used in the original request.`,
				},
				resource.Attribute{
					Name:        "routing_rules",
					Description: `(Optional) A json array containing [routing rules](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-websiteconfiguration-routingrules.html) describing redirect behavior and when redirects are applied. The ` + "`" + `cors_rule` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) Enable versioning. Once you version-enable a bucket, it can never return to an unversioned state. You can, however, suspend versioning on that bucket.`,
				},
				resource.Attribute{
					Name:        "mfa_delete",
					Description: `(Optional) Enable MFA delete for either ` + "`" + `Change the versioning state of your bucket` + "`" + ` or ` + "`" + `Permanently delete an object version` + "`" + `. Default is ` + "`" + `false` + "`" + `. The ` + "`" + `logging` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "target_bucket",
					Description: `(Required) The name of the bucket that will receive the log objects.`,
				},
				resource.Attribute{
					Name:        "target_prefix",
					Description: `(Optional) To specify a key prefix for log objects. The ` + "`" + `lifecycle_rule` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) Unique identifier for the rule.`,
				},
				resource.Attribute{
					Name:        "prefix",
					Description: `(Optional) Object key prefix identifying one or more objects to which the rule applies.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Required) Specifies lifecycle rule status.`,
				},
				resource.Attribute{
					Name:        "expiration",
					Description: `(Optional) Specifies a period in the object's expire (documented below).`,
				},
				resource.Attribute{
					Name:        "noncurrent_version_expiration",
					Description: `(Optional) Specifies when noncurrent object versions expire (documented below). At least one of ` + "`" + `expiration` + "`" + `, ` + "`" + `noncurrent_version_expiration` + "`" + ` must be specified. The ` + "`" + `expiration` + "`" + ` object supports the following`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The name of the bucket.`,
				},
				resource.Attribute{
					Name:        "arn",
					Description: `The ARN of the bucket. Will be of format ` + "`" + `arn:aws:s3:::bucketname` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "bucket_domain_name",
					Description: `The bucket domain name. Will be of format ` + "`" + `bucketname.s3.amazonaws.com` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "hosted_zone_id",
					Description: `The [Route 53 Hosted Zone ID](https://docs.aws.amazon.com/general/latest/gr/rande.html#s3_website_region_endpoints) for this bucket's region.`,
				},
				resource.Attribute{
					Name:        "website_endpoint",
					Description: `The website endpoint, if the bucket is configured with a website. If not, this will be an empty string.`,
				},
				resource.Attribute{
					Name:        "website_domain",
					Description: `The domain of the website endpoint, if the bucket is configured with a website. If not, this will be an empty string. This is used to create Route 53 alias records. ## Import S3 bucket can be imported using the ` + "`" + `bucket` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import opentelekomcloud_s3_bucket.bucket bucket-name ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "id",
					Description: `The name of the bucket.`,
				},
				resource.Attribute{
					Name:        "arn",
					Description: `The ARN of the bucket. Will be of format ` + "`" + `arn:aws:s3:::bucketname` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "bucket_domain_name",
					Description: `The bucket domain name. Will be of format ` + "`" + `bucketname.s3.amazonaws.com` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "hosted_zone_id",
					Description: `The [Route 53 Hosted Zone ID](https://docs.aws.amazon.com/general/latest/gr/rande.html#s3_website_region_endpoints) for this bucket's region.`,
				},
				resource.Attribute{
					Name:        "website_endpoint",
					Description: `The website endpoint, if the bucket is configured with a website. If not, this will be an empty string.`,
				},
				resource.Attribute{
					Name:        "website_domain",
					Description: `The domain of the website endpoint, if the bucket is configured with a website. If not, this will be an empty string. This is used to create Route 53 alias records. ## Import S3 bucket can be imported using the ` + "`" + `bucket` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import opentelekomcloud_s3_bucket.bucket bucket-name ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opentelekomcloud_s3-bucket-object",
			Category:         "S3 Resource",
			ShortDescription: `Provides a S3 bucket object resource.`,
			Description:      ``,
			Keywords: []string{
				"s3",
				"resource",
				"s3-bucket-object",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "bucket",
					Description: `(Required) The name of the bucket to put the file in.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) The name of the object once it is in the bucket.`,
				},
				resource.Attribute{
					Name:        "source",
					Description: `(Required) The path to the source file being uploaded to the bucket.`,
				},
				resource.Attribute{
					Name:        "content",
					Description: `(Required unless ` + "`" + `source` + "`" + ` given) The literal content being uploaded to the bucket.`,
				},
				resource.Attribute{
					Name:        "acl",
					Description: `(Optional) The [canned ACL](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl) to apply. Defaults to "private".`,
				},
				resource.Attribute{
					Name:        "cache_control",
					Description: `(Optional) Specifies caching behavior along the request/reply chain Read [w3c cache_control](http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.9) for further details.`,
				},
				resource.Attribute{
					Name:        "content_disposition",
					Description: `(Optional) Specifies presentational information for the object. Read [wc3 content_disposition](http://www.w3.org/Protocols/rfc2616/rfc2616-sec19.html#sec19.5.1) for further information.`,
				},
				resource.Attribute{
					Name:        "content_encoding",
					Description: `(Optional) Specifies what content encodings have been applied to the object and thus what decoding mechanisms must be applied to obtain the media-type referenced by the Content-Type header field. Read [w3c content encoding](http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.11) for further information.`,
				},
				resource.Attribute{
					Name:        "content_language",
					Description: `(Optional) The language the content is in e.g. en-US or en-GB.`,
				},
				resource.Attribute{
					Name:        "content_type",
					Description: `(Optional) A standard MIME type describing the format of the object data, e.g. application/octet-stream. All Valid MIME Types are valid for this input.`,
				},
				resource.Attribute{
					Name:        "website_redirect",
					Description: `(Optional) Specifies a target URL for [website redirect](http://docs.aws.amazon.com/AmazonS3/latest/dev/how-to-page-redirect.html).`,
				},
				resource.Attribute{
					Name:        "etag",
					Description: `(Optional) Used to trigger updates. The only meaningful value is ` + "`" + `${md5(file("path/to/file"))}` + "`" + `. This attribute is not compatible with ` + "`" + `kms_key_id` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "server_side_encryption",
					Description: `(Optional) Specifies server-side encryption of the object in S3. Valid values are "` + "`" + `AES256` + "`" + `" and "` + "`" + `aws:kms` + "`" + `". Either ` + "`" + `source` + "`" + ` or ` + "`" + `content` + "`" + ` must be provided to specify the bucket content. These two arguments are mutually-exclusive. ## Attributes Reference The following attributes are exported`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `the ` + "`" + `key` + "`" + ` of the resource supplied above`,
				},
				resource.Attribute{
					Name:        "etag",
					Description: `the ETag generated for the object (an MD5 sum of the object content).`,
				},
				resource.Attribute{
					Name:        "version_id",
					Description: `A unique version ID value for the object, if bucket versioning is enabled.`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "id",
					Description: `the ` + "`" + `key` + "`" + ` of the resource supplied above`,
				},
				resource.Attribute{
					Name:        "etag",
					Description: `the ETag generated for the object (an MD5 sum of the object content).`,
				},
				resource.Attribute{
					Name:        "version_id",
					Description: `A unique version ID value for the object, if bucket versioning is enabled.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opentelekomcloud_s3_bucket_policy",
			Category:         "S3 Resource",
			ShortDescription: `Attaches a policy to an S3 bucket resource.`,
			Description:      ``,
			Keywords: []string{
				"s3",
				"resource",
				"bucket",
				"policy",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "bucket",
					Description: `(Required) The name of the bucket to which to apply the policy.`,
				},
				resource.Attribute{
					Name:        "policy",
					Description: `(Required) The text of the policy.`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opentelekomcloud_sfs_file_system_v2",
			Category:         "SFS Resources",
			ShortDescription: `Provides an Scalable File Resource (SFS) resource.`,
			Description:      ``,
			Keywords: []string{
				"sfs",
				"file",
				"system",
				"v2",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "size",
					Description: `(Required) The size (GB) of the shared file system.`,
				},
				resource.Attribute{
					Name:        "share_proto",
					Description: `(Optional) The protocol for sharing file systems. The default value is NFS.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the shared file system.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Describes the shared file system.`,
				},
				resource.Attribute{
					Name:        "is_public",
					Description: `(Optional) The level of visibility for the shared file system.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `(Optional) Metadata key and value pairs as a dictionary of strings.Changing this will create a new resource.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `(Optional) The availability zone name.Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "access_level",
					Description: `(Required) The access level of the shared file system. Changing this will create a new access rule.`,
				},
				resource.Attribute{
					Name:        "access_type",
					Description: `(Optional) The type of the share access rule. Changing this will create a new access rule.`,
				},
				resource.Attribute{
					Name:        "access_to",
					Description: `(Required) The access that the back end grants or denies. Changing this will create a new access rule ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The UUID of the shared file system.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the shared file system.`,
				},
				resource.Attribute{
					Name:        "share_type",
					Description: `The storage service type assigned for the shared file system, such as high-performance storage (composed of SSDs) and large-capacity storage (composed of SATA disks).`,
				},
				resource.Attribute{
					Name:        "volume_type",
					Description: `The volume type.`,
				},
				resource.Attribute{
					Name:        "export_location",
					Description: `The address for accessing the shared file system.`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `The host name of the shared file system.`,
				},
				resource.Attribute{
					Name:        "share_access_id",
					Description: `The UUID of the share access rule.`,
				},
				resource.Attribute{
					Name:        "access_rules_status",
					Description: `The status of the share access rule. ## Import SFS can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import opentelekomcloud_sfs_file_system_v2 4779ab1c-7c1a-44b1-a02e-93dfc361b32d ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "id",
					Description: `The UUID of the shared file system.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the shared file system.`,
				},
				resource.Attribute{
					Name:        "share_type",
					Description: `The storage service type assigned for the shared file system, such as high-performance storage (composed of SSDs) and large-capacity storage (composed of SATA disks).`,
				},
				resource.Attribute{
					Name:        "volume_type",
					Description: `The volume type.`,
				},
				resource.Attribute{
					Name:        "export_location",
					Description: `The address for accessing the shared file system.`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `The host name of the shared file system.`,
				},
				resource.Attribute{
					Name:        "share_access_id",
					Description: `The UUID of the share access rule.`,
				},
				resource.Attribute{
					Name:        "access_rules_status",
					Description: `The status of the share access rule. ## Import SFS can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import opentelekomcloud_sfs_file_system_v2 4779ab1c-7c1a-44b1-a02e-93dfc361b32d ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opentelekomcloud_smn_subscription_v2",
			Category:         "SMN Resource",
			ShortDescription: `Manages a V2 subscription resource within OpenTelekomCloud.`,
			Description:      ``,
			Keywords: []string{
				"smn",
				"resource",
				"subscription",
				"v2",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "topic_urn",
					Description: `(Required) Resource identifier of a topic, which is unique.`,
				},
				resource.Attribute{
					Name:        "endpoint",
					Description: `(Required) Message endpoint. For an HTTP subscription, the endpoint starts with http\://. For an HTTPS subscription, the endpoint starts with https\://. For an email subscription, the endpoint is a mail address. For an SMS message subscription, the endpoint is a phone number.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Required) Protocol of the message endpoint. Currently, email, sms, http, and https are supported.`,
				},
				resource.Attribute{
					Name:        "remark",
					Description: `(Optional) Remark information. The remarks must be a UTF-8-coded character string containing 128 bytes.`,
				},
				resource.Attribute{
					Name:        "subscription_urn",
					Description: `(Optional) Resource identifier of a subscription, which is unique.`,
				},
				resource.Attribute{
					Name:        "owner",
					Description: `(Optional) Project ID of the topic creator.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Optional) Subscription status. 0 indicates that the subscription is not confirmed. 1 indicates that the subscription is confirmed. 3 indicates that the subscription is canceled. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "topic_urn",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "endpoint",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "remark",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "subscription_urn",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "owner",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `See Argument Reference above.`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "topic_urn",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "endpoint",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "remark",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "subscription_urn",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "owner",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `See Argument Reference above.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opentelekomcloud_smn_topic_v2",
			Category:         "SMN Resource",
			ShortDescription: `Manages a V2 topic resource within OpenTelekomCloud.`,
			Description:      ``,
			Keywords: []string{
				"smn",
				"resource",
				"topic",
				"v2",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the topic to be created.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) Topic display name, which is presented as the name of the email sender in an email message. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "topic_urn",
					Description: `Resource identifier of a topic, which is unique.`,
				},
				resource.Attribute{
					Name:        "push_policy",
					Description: `Message pushing policy. 0 indicates that the message sending fails and the message is cached in the queue. 1 indicates that the failed message is discarded.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Time when the topic was created.`,
				},
				resource.Attribute{
					Name:        "update_time",
					Description: `Time when the topic was updated.`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "topic_urn",
					Description: `Resource identifier of a topic, which is unique.`,
				},
				resource.Attribute{
					Name:        "push_policy",
					Description: `Message pushing policy. 0 indicates that the message sending fails and the message is cached in the queue. 1 indicates that the failed message is discarded.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Time when the topic was created.`,
				},
				resource.Attribute{
					Name:        "update_time",
					Description: `Time when the topic was updated.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opentelekomcloud-vbs-backup-policy-v2",
			Category:         "VBS Resources",
			ShortDescription: `Provides an VBS Backup Policy resource.`,
			Description:      ``,
			Keywords: []string{
				"vbs",
				"opentelekomcloud-vbs-backup-policy-v2",
			},
			Arguments:  []resource.Argument{},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opentelekomcloud-vbs-backup-share-v2",
			Category:         "VBS Resources",
			ShortDescription: `Provides an VBS Backup Share resource.`,
			Description:      ``,
			Keywords: []string{
				"vbs",
				"opentelekomcloud-vbs-backup-share-v2",
			},
			Arguments:  []resource.Argument{},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opentelekomcloud-vbs-backup-v2",
			Category:         "VBS Resources",
			ShortDescription: `Provides an VBS Backup resource.`,
			Description:      ``,
			Keywords: []string{
				"vbs",
				"opentelekomcloud-vbs-backup-v2",
			},
			Arguments:  []resource.Argument{},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opentelekomcloud_vpc_eip_v1",
			Category:         "EIP Resources",
			ShortDescription: `Manages a V1 EIP resource within OpenTelekomCloud VPC.`,
			Description:      ``,
			Keywords: []string{
				"eip",
				"vpc",
				"v1",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "publicip",
					Description: `(Required) The elastic IP address object.`,
				},
				resource.Attribute{
					Name:        "bandwidth",
					Description: `(Required) The bandwidth object. The ` + "`" + `publicip` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) The value must be a type supported by the system. Only ` + "`" + `5_bgp` + "`" + ` supported now. Changing this creates a new eip.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `(Optional) The value must be a valid IP address in the available IP address segment. Changing this creates a new eip.`,
				},
				resource.Attribute{
					Name:        "port_id",
					Description: `(Optional) The port id which this eip will associate with. If the value is "" or this not specified, the eip will be in unbind state. The ` + "`" + `bandwidth` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The bandwidth name, which is a string of 1 to 64 characters that contain letters, digits, underscores (_), and hyphens (-).`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `(Required) The bandwidth size. The value ranges from 1 to 300 Mbit/s.`,
				},
				resource.Attribute{
					Name:        "share_type",
					Description: `(Required) Whether the bandwidth is shared or exclusive. Changing this creates a new eip.`,
				},
				resource.Attribute{
					Name:        "charge_mode",
					Description: `(Optional) This is a reserved field. If the system supports charging by traffic and this field is specified, then you are charged by traffic for elastic IP addresses. Changing this creates a new eip. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "publicip/type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "publicip/ip_address",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "publicip/port_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "bandwidth/name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "bandwidth/size",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "bandwidth/share_type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "bandwidth/charge_mode",
					Description: `See Argument Reference above. ## Import EIPs can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import opentelekomcloud_vpc_eip_v1.eip_1 2c7f39f3-702b-48d1-940c-b50384177ee1 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "publicip/type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "publicip/ip_address",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "publicip/port_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "bandwidth/name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "bandwidth/size",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "bandwidth/share_type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "bandwidth/charge_mode",
					Description: `See Argument Reference above. ## Import EIPs can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import opentelekomcloud_vpc_eip_v1.eip_1 2c7f39f3-702b-48d1-940c-b50384177ee1 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opentelekomcloud_vpc_peering_connection_accepter_v2",
			Category:         "Networking Resources",
			ShortDescription: `Manage the accepter's side of a VPC Peering Connection.`,
			Description:      ``,
			Keywords: []string{
				"networking",
				"vpc",
				"peering",
				"connection",
				"accepter",
				"v2",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `The VPC peering connection name.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The VPC peering connection ID.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The VPC peering connection status.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `The ID of requester VPC involved in a VPC peering connection.`,
				},
				resource.Attribute{
					Name:        "peer_vpc_id",
					Description: `The VPC ID of the accepter tenant.`,
				},
				resource.Attribute{
					Name:        "peer_tenant_id",
					Description: `The Tenant Id of the accepter tenant.`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `The VPC peering connection name.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The VPC peering connection ID.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The VPC peering connection status.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `The ID of requester VPC involved in a VPC peering connection.`,
				},
				resource.Attribute{
					Name:        "peer_vpc_id",
					Description: `The VPC ID of the accepter tenant.`,
				},
				resource.Attribute{
					Name:        "peer_tenant_id",
					Description: `The Tenant Id of the accepter tenant.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opentelekomcloud_vpc_peering_connection_v2",
			Category:         "Networking Resources",
			ShortDescription: `Manage a VPC Peering Connection resource.`,
			Description:      ``,
			Keywords: []string{
				"networking",
				"vpc",
				"peering",
				"connection",
				"v2",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "id",
					Description: `The VPC peering connection ID.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The VPC peering connection status. The value can be PENDING_ACCEPTANCE, REJECTED, EXPIRED, DELETED, or ACTIVE. ## Notes If you create a VPC peering connection with another VPC of your own, the connection is created without the need for you to accept the connection. ## Import VPC Peering resources can be imported using the ` + "`" + `vpc peering id` + "`" + `, e.g. > $ terraform import opentelekomcloud_vpc_peering_connection_v2.test_connection 22b76469-08e3-4937-8c1d-7aad34892be1`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "id",
					Description: `The VPC peering connection ID.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The VPC peering connection status. The value can be PENDING_ACCEPTANCE, REJECTED, EXPIRED, DELETED, or ACTIVE. ## Notes If you create a VPC peering connection with another VPC of your own, the connection is created without the need for you to accept the connection. ## Import VPC Peering resources can be imported using the ` + "`" + `vpc peering id` + "`" + `, e.g. > $ terraform import opentelekomcloud_vpc_peering_connection_v2.test_connection 22b76469-08e3-4937-8c1d-7aad34892be1`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opentelekomcloud_vpc_route_v2",
			Category:         "Networking Resources",
			ShortDescription: `Provides an VPC route resource.`,
			Description:      ``,
			Keywords: []string{
				"networking",
				"vpc",
				"route",
				"v2",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "id",
					Description: `The route ID.`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "id",
					Description: `The route ID.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opentelekomcloud_vpc_subnet_v1",
			Category:         "Networking Resources",
			ShortDescription: `Provides an VPC subnet resource.`,
			Description:      ``,
			Keywords: []string{
				"networking",
				"vpc",
				"subnet",
				"v1",
			},
			Arguments:  []resource.Argument{},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opentelekomcloud_vpc_v1",
			Category:         "Networking Resources",
			ShortDescription: `Manages a V1 VPC resource within OpenTelekomCloud.`,
			Description:      ``,
			Keywords: []string{
				"networking",
				"vpc",
				"v1",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "cidr",
					Description: `(Required) The range of available subnets in the VPC. The value ranges from 10.0.0.0/8 to 10.255.255.0/24, 172.16.0.0/12 to 172.31.255.0/24, or 192.168.0.0/16 to 192.168.255.0/24.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the VPC. The name must be unique for a tenant. The value is a string of no more than 64 characters and can contain digits, letters, underscores (_), and hyphens (-). Changing this updates the name of the existing VPC. ## Attributes Reference The following attributes are exported:`,
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
					Name:        "cidr",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The current status of the desired VPC. Can be either CREATING, OK, DOWN, PENDING_UPDATE, PENDING_DELETE, or ERROR.`,
				},
				resource.Attribute{
					Name:        "shared",
					Description: `Specifies whether the cross-tenant sharing is supported. ## Import VPCs can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import opentelekomcloud_vpc_v1.vpc_v1 7117d38e-4c8f-4624-a505-bd96b97d024c ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the VPC.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "cidr",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The current status of the desired VPC. Can be either CREATING, OK, DOWN, PENDING_UPDATE, PENDING_DELETE, or ERROR.`,
				},
				resource.Attribute{
					Name:        "shared",
					Description: `Specifies whether the cross-tenant sharing is supported. ## Import VPCs can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import opentelekomcloud_vpc_v1.vpc_v1 7117d38e-4c8f-4624-a505-bd96b97d024c ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opentelekomcloud_waf_ccattackprotection_rule_v1",
			Category:         "WAF Resources",
			ShortDescription: `Manages a V1 WAF CC Attack Protection Rule resource within OpenTelekomCloud.`,
			Description:      ``,
			Keywords: []string{
				"waf",
				"ccattackprotection",
				"rule",
				"v1",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "policy_id",
					Description: `(Required) The WAF policy ID. Changing this creates a new rule.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Required) Specifies a misreported URL excluding a domain name. Changing this creates a new rule.`,
				},
				resource.Attribute{
					Name:        "limit_num",
					Description: `(Required) Specifies the number of requests allowed from a web visitor in a rate limiting period. Changing this creates a new rule.`,
				},
				resource.Attribute{
					Name:        "limit_period",
					Description: `(Required) Specifies the rate limiting period. Changing this creates a new rule.`,
				},
				resource.Attribute{
					Name:        "lock_time",
					Description: `(Optional) Specifies the lock duration. The value ranges from 0 seconds to 2^32 seconds. Changing this creates a new rule.`,
				},
				resource.Attribute{
					Name:        "tag_type",
					Description: `(Required) Specifies the rate limit mode. Changing this creates a new rule. Valid Options are:`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `A web visitor is identified by the IP address.`,
				},
				resource.Attribute{
					Name:        "cookie",
					Description: `A web visitor is identified by the cookie key value.`,
				},
				resource.Attribute{
					Name:        "other",
					Description: `A web visitor is identified by the Referer field(user-defined request source).`,
				},
				resource.Attribute{
					Name:        "tag_index",
					Description: `(Optional) If ` + "`" + `tag_type` + "`" + ` is set to ` + "`" + `cookie` + "`" + `, this parameter indicates cookie name. Changing this creates a new rule.`,
				},
				resource.Attribute{
					Name:        "tag_category",
					Description: `(Optional) Specifies the category. The value is ` + "`" + `referer` + "`" + `. Changing this creates a new rule.`,
				},
				resource.Attribute{
					Name:        "tag_contents",
					Description: `(Optional) Specifies the category content. Changing this creates a new rule.`,
				},
				resource.Attribute{
					Name:        "action_category",
					Description: `(Required) Specifies the action. Changing this creates a new rule. Valid Options are:`,
				},
				resource.Attribute{
					Name:        "block",
					Description: `block the requests.`,
				},
				resource.Attribute{
					Name:        "captcha",
					Description: `Verification code. The user needs to enter the correct verification code after blocking to restore the correct access page.`,
				},
				resource.Attribute{
					Name:        "block_content_type",
					Description: `(Optional) Specifies the type of the returned page. The options are ` + "`" + `application/json` + "`" + `, ` + "`" + `text/html` + "`" + `, and ` + "`" + `text/xml` + "`" + `. Changing this creates a new rule.`,
				},
				resource.Attribute{
					Name:        "block_content",
					Description: `(Optional) Specifies the content of the returned page. Changing this creates a new rule. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the rule.`,
				},
				resource.Attribute{
					Name:        "default",
					Description: `Specifies whether the rule is the default CC attack protection rule. ## Import CC Attack Protection Rules can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import opentelekomcloud_waf_ccattackprotection_rule_v1.rule_1 7117d38e-4c8f-4624-a505-bd96b97d024c ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the rule.`,
				},
				resource.Attribute{
					Name:        "default",
					Description: `Specifies whether the rule is the default CC attack protection rule. ## Import CC Attack Protection Rules can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import opentelekomcloud_waf_ccattackprotection_rule_v1.rule_1 7117d38e-4c8f-4624-a505-bd96b97d024c ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opentelekomcloud_waf_certificate_v1",
			Category:         "WAF Resources",
			ShortDescription: `Manages a V1 WAF certificate resource within OpenTelekomCloud.`,
			Description:      ``,
			Keywords: []string{
				"waf",
				"certificate",
				"v1",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The certificate name. The maximum length is 256 characters. Only digits, letters, underscores(_), and hyphens(-) are allowed.`,
				},
				resource.Attribute{
					Name:        "content",
					Description: `(Optional) The certificate content. Changing this creates a new certificate.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Optional) The private key. Changing this creates a new certificate. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the certificate.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "content",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `See Argument Reference above. ## Import Certificates can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import opentelekomcloud_waf_certificate_v1.cert_1 7117d38e-4c8f-4624-a505-bd96b97d024c ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the certificate.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "content",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `See Argument Reference above. ## Import Certificates can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import opentelekomcloud_waf_certificate_v1.cert_1 7117d38e-4c8f-4624-a505-bd96b97d024c ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opentelekomcloud_waf_datamasking_rule_v1",
			Category:         "WAF Resources",
			ShortDescription: `Manages a V1 WAF Data Masking Rule resource within OpenTelekomCloud.`,
			Description:      ``,
			Keywords: []string{
				"waf",
				"datamasking",
				"rule",
				"v1",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "policy_id",
					Description: `(Required) The WAF policy ID. Changing this creates a new rule.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Required) Specifies the URL to which the data masking rule applies.`,
				},
				resource.Attribute{
					Name:        "category",
					Description: `(Required) Specifies the masked field. The options are params and header.`,
				},
				resource.Attribute{
					Name:        "index",
					Description: `(Required) Specifies the masked subfield. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the rule. ## Import Data Masking Rules can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import opentelekomcloud_waf_datamasking_rule_v1.rule_1 7117d38e-4c8f-4624-a505-bd96b97d024c ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the rule. ## Import Data Masking Rules can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import opentelekomcloud_waf_datamasking_rule_v1.rule_1 7117d38e-4c8f-4624-a505-bd96b97d024c ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opentelekomcloud_waf_domain_v1",
			Category:         "WAF Resources",
			ShortDescription: `Manages a V1 WAF domain resource within OpenTelekomCloud.`,
			Description:      ``,
			Keywords: []string{
				"waf",
				"domain",
				"v1",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "hostname",
					Description: `(Required) The domain name. For example, www.example.com or`,
				},
				resource.Attribute{
					Name:        "certificate_id",
					Description: `(Optional) The certificate ID. This parameter is mandatory when front_protocol is set to HTTPS.`,
				},
				resource.Attribute{
					Name:        "server",
					Description: `(Required) Array of server object. The server object structure is documented below.`,
				},
				resource.Attribute{
					Name:        "proxy",
					Description: `(Required) Specifies whether a proxy is configured.`,
				},
				resource.Attribute{
					Name:        "sip_header_name",
					Description: `(Optional) The type of the source IP header. This parameter is required only when proxy is set to true. The options are as follows: default, cloudflare, akamai, and custom.`,
				},
				resource.Attribute{
					Name:        "sip_header_list",
					Description: `(Optional) Array of HTTP request header for identifying the real source IP address. This parameter is required only when proxy is set to true.`,
				},
				resource.Attribute{
					Name:        "front_protocol",
					Description: `(Required) Protocol type of the client. The options are HTTP and HTTPS.`,
				},
				resource.Attribute{
					Name:        "back_protocol",
					Description: `(Required) Protocl used by WAF to forward client requests to the server. The options are HTTP and HTTPS.`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `(Required) IP address or domain name of the web server that the client accesses. For example, 192.168.1.1 or www.a.com.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Required) Port number used by the web server. The value ranges from 0 to 65535, for example, 8080. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the domain.`,
				},
				resource.Attribute{
					Name:        "access_code",
					Description: `The acccess code.`,
				},
				resource.Attribute{
					Name:        "cname",
					Description: `The CNAME value.`,
				},
				resource.Attribute{
					Name:        "txt_code",
					Description: `The TXT record. This attribute is returned only when proxy is set to true.`,
				},
				resource.Attribute{
					Name:        "sub_domain",
					Description: `The subdomain name. This attribute is returned only when proxy is set to true.`,
				},
				resource.Attribute{
					Name:        "policy_id",
					Description: `The proxy ID.`,
				},
				resource.Attribute{
					Name:        "protect_status",
					Description: `The WAF mode. -1: bypassed, 0: disabled, 1: enabled.`,
				},
				resource.Attribute{
					Name:        "access_status",
					Description: `Whether a domain name is connected to WAF. 0: The domain name is not connected to WAF, 1: The domain name is connected to WAF.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `The protocol type of the client. The options are HTTP, HTTPS, and HTTP&HTTPS. ## Import Domains can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import opentelekomcloud_waf_domain_v1.dom_1 7117d38e-4c8f-4624-a505-bd96b97d024c ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the domain.`,
				},
				resource.Attribute{
					Name:        "access_code",
					Description: `The acccess code.`,
				},
				resource.Attribute{
					Name:        "cname",
					Description: `The CNAME value.`,
				},
				resource.Attribute{
					Name:        "txt_code",
					Description: `The TXT record. This attribute is returned only when proxy is set to true.`,
				},
				resource.Attribute{
					Name:        "sub_domain",
					Description: `The subdomain name. This attribute is returned only when proxy is set to true.`,
				},
				resource.Attribute{
					Name:        "policy_id",
					Description: `The proxy ID.`,
				},
				resource.Attribute{
					Name:        "protect_status",
					Description: `The WAF mode. -1: bypassed, 0: disabled, 1: enabled.`,
				},
				resource.Attribute{
					Name:        "access_status",
					Description: `Whether a domain name is connected to WAF. 0: The domain name is not connected to WAF, 1: The domain name is connected to WAF.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `The protocol type of the client. The options are HTTP, HTTPS, and HTTP&HTTPS. ## Import Domains can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import opentelekomcloud_waf_domain_v1.dom_1 7117d38e-4c8f-4624-a505-bd96b97d024c ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opentelekomcloud_waf_falsealarmmasking_rule_v1",
			Category:         "WAF Resources",
			ShortDescription: `Manages a V1 WAF False Alarm Masking Rule resource within OpenTelekomCloud.`,
			Description:      ``,
			Keywords: []string{
				"waf",
				"falsealarmmasking",
				"rule",
				"v1",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "policy_id",
					Description: `(Required) The WAF policy ID. Changing this creates a new rule.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Required) Specifies a misreported URL excluding a domain name. Changing this creates a new rule.`,
				},
				resource.Attribute{
					Name:        "rule",
					Description: `(Required) Specifies the rule ID, which consists of six digits and cannot be empty. Changing this creates a new rule. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the rule. ## Import False Alarm Masking Rules can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import opentelekomcloud_waf_falsealarmmasking_rule_v1.rule_1 7117d38e-4c8f-4624-a505-bd96b97d024c ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the rule. ## Import False Alarm Masking Rules can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import opentelekomcloud_waf_falsealarmmasking_rule_v1.rule_1 7117d38e-4c8f-4624-a505-bd96b97d024c ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opentelekomcloud_waf_policy_v1",
			Category:         "WAF Resources",
			ShortDescription: `Manages a V1 WAF policy resource within OpenTelekomCloud.`,
			Description:      ``,
			Keywords: []string{
				"waf",
				"policy",
				"v1",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The policy name. The maximum length is 256 characters. Only digits, letters, underscores(_), and hyphens(-) are allowed.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Optional) Specifies the protective action after a rule is matched. The action object structure is documented below.`,
				},
				resource.Attribute{
					Name:        "options",
					Description: `(Optional) Specifies the protection switches. The options object structure is documented below.`,
				},
				resource.Attribute{
					Name:        "level",
					Description: `(Optional) Specifies the protection level.`,
				},
				resource.Attribute{
					Name:        "full_detection",
					Description: `(Optional) Specifies the detection mode in Precise Protection.`,
				},
				resource.Attribute{
					Name:        "hosts",
					Description: `(Optional) An array of the domain IDs. The ` + "`" + `action` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "category",
					Description: `(Required) Specifies the protective action.`,
				},
				resource.Attribute{
					Name:        "webattack",
					Description: `(Optional) Specifies whether Basic Web Protection is enabled.`,
				},
				resource.Attribute{
					Name:        "common",
					Description: `(Optional) Specifies whether General Check in Basic Web Protection is enabled.`,
				},
				resource.Attribute{
					Name:        "crawler",
					Description: `(Optional) Specifies whether the master crawler detection switch in Basic Web Protection is enabled.`,
				},
				resource.Attribute{
					Name:        "crawler_engine",
					Description: `(Optional) Specifies whether the Search Engine switch in Basic Web Protection is enabled.`,
				},
				resource.Attribute{
					Name:        "crawler_scanner",
					Description: `(Optional) Specifies whether the Scanner switch in Basic Web Protection is enabled.`,
				},
				resource.Attribute{
					Name:        "crawler_script",
					Description: `(Optional) Specifies whether the Script Tool switch in Basic Web Protection is enabled.`,
				},
				resource.Attribute{
					Name:        "crawler_other",
					Description: `(Optional) Specifies whether detection of other crawlers in Basic Web Protection is enabled.`,
				},
				resource.Attribute{
					Name:        "webshell",
					Description: `(Optional) Specifies whether webshell detection in Basic Web Protection is enabled.`,
				},
				resource.Attribute{
					Name:        "cc",
					Description: `(Optional) Specifies whether CC Attack Protection is enabled.`,
				},
				resource.Attribute{
					Name:        "custom",
					Description: `(Optional) Specifies whether Precise Protection is enabled.`,
				},
				resource.Attribute{
					Name:        "whiteblackip",
					Description: `(Optional) Specifies whether Blacklist and Whitelist is enabled.`,
				},
				resource.Attribute{
					Name:        "privacy",
					Description: `(Optional) Specifies whether Data Masking is enabled.`,
				},
				resource.Attribute{
					Name:        "ignore",
					Description: `(Optional) Specifies whether False Alarm Masking is enabled.`,
				},
				resource.Attribute{
					Name:        "antitamper",
					Description: `(Optional) Specifies whether Web Tamper Protection is enabled. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the policy.`,
				},
				resource.Attribute{
					Name:        "hosts",
					Description: `Specifies the domain IDs. ## Import Policies can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import opentelekomcloud_waf_policy_v1.policy_1 7117d38e-4c8f-4624-a505-bd96b97d024c ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the policy.`,
				},
				resource.Attribute{
					Name:        "hosts",
					Description: `Specifies the domain IDs. ## Import Policies can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import opentelekomcloud_waf_policy_v1.policy_1 7117d38e-4c8f-4624-a505-bd96b97d024c ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opentelekomcloud_waf_preciseprotection_rule_v1",
			Category:         "WAF Resources",
			ShortDescription: `Manages a V1 WAF Precise Protection Rule resource within OpenTelekomCloud.`,
			Description:      ``,
			Keywords: []string{
				"waf",
				"preciseprotection",
				"rule",
				"v1",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "policy_id",
					Description: `(Required) The WAF policy ID. Changing this creates a new rule.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Specifies the name of a precise protection rule. Changing this creates a new rule.`,
				},
				resource.Attribute{
					Name:        "time",
					Description: `(Optional) Specifies the effect time of the precise protection rule. Changing this creates a new rule.`,
				},
				resource.Attribute{
					Name:        "false",
					Description: `The rule takes effect immediately.`,
				},
				resource.Attribute{
					Name:        "true",
					Description: `The rule takes effect at the scheduled time.`,
				},
				resource.Attribute{
					Name:        "start",
					Description: `(Optional) Specifies the time when the precise protection rule takes effect. If time is set to true, either the start time or the end time must be set. Changing this creates a new rule.`,
				},
				resource.Attribute{
					Name:        "end",
					Description: `(Optional) Specifies the time when the precise protection rule expires. If time is set to true, either the start time or the end time must be set. Changing this creates a new rule.`,
				},
				resource.Attribute{
					Name:        "conditions",
					Description: `(Required) Specifies the condition parameters. Changing this creates a new rule. The conditions object structure is documented below.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Required) Specifies the protective action after the precise protection rule is matched. Changing this creates a new rule. The action object structure is documented below.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `(Optional) Specifies the priority of a rule being executed. Smaller values correspond to higher priorities. If two rules are assigned with the same priority, the rule added earlier has higher priority, the rule added earlier has higher priority. The value ranges from 0 to 65535. Changing this creates a new rule. The ` + "`" + `conditions` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "category",
					Description: `(Required) Specifies the condition type. The value can be url, user-agent, ip, params, cookie, referer, or header.`,
				},
				resource.Attribute{
					Name:        "index",
					Description: `(Optional) If ` + "`" + `category` + "`" + ` is set to cookie, index indicates cookie name, if set to params, index indicates param name, if set to header, index indicates an option in the header.`,
				},
				resource.Attribute{
					Name:        "logic",
					Description: `(Required) 1,2,3,4,5,6,7, and 8 indicate include, exclude, equal to, not equal to, prefix is, prefix is not, suffix is, and suffix is not, respectively. If ` + "`" + `category` + "`" + ` is set to ip, logic can only be 3 or 4.`,
				},
				resource.Attribute{
					Name:        "contents",
					Description: `(Required) Specifies a list of content matching the condition. Currently, only one value is accepted. The ` + "`" + `action` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "category",
					Description: `(Required) Specifies the protective action. The value can be block or pass. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the rule. ## Import Precise Protection Rules can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import opentelekomcloud_waf_preciseprotection_rule_v1.rule_1 7117d38e-4c8f-4624-a505-bd96b97d024c ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the rule. ## Import Precise Protection Rules can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import opentelekomcloud_waf_preciseprotection_rule_v1.rule_1 7117d38e-4c8f-4624-a505-bd96b97d024c ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opentelekomcloud_waf_webtamperprotection_rule_v1",
			Category:         "WAF Resources",
			ShortDescription: `Manages a V1 WAF Web Tamper Protection Rule resource within OpenTelekomCloud.`,
			Description:      ``,
			Keywords: []string{
				"waf",
				"webtamperprotection",
				"rule",
				"v1",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "policy_id",
					Description: `(Required) The WAF policy ID. Changing this creates a new rule.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `(Required) Specifies the domain name. Changing this creates a new rule.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Required) Specifies the URL protected by the web tamper protection rule, excluding a domain name. Changing this creates a new rule. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the rule. ## Import Web Tamper Protection Rules can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import opentelekomcloud_waf_webtamperprotection_rule_v1.rule_1 7117d38e-4c8f-4624-a505-bd96b97d024c ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the rule. ## Import Web Tamper Protection Rules can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import opentelekomcloud_waf_webtamperprotection_rule_v1.rule_1 7117d38e-4c8f-4624-a505-bd96b97d024c ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opentelekomcloud_waf_whiteblackip_rule_v1",
			Category:         "WAF Resources",
			ShortDescription: `Manages a V1 WAF WhiteBlackIP Rule resource within OpenTelekomCloud.`,
			Description:      ``,
			Keywords: []string{
				"waf",
				"whiteblackip",
				"rule",
				"v1",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "policy_id",
					Description: `(Required) The WAF policy ID. Changing this creates a new rule.`,
				},
				resource.Attribute{
					Name:        "addr",
					Description: `(Required) Specifies the IP address or range. For example, 192.168.0.125 or 192.168.0.0/24.`,
				},
				resource.Attribute{
					Name:        "white",
					Description: `(Optional) Specifies the IP address type. 1: Whitelist, 0: Blacklist. If you do not configure the white parameter, the value is Blacklist by default. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the rule. ## Import WhiteBlackIP Rules can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import opentelekomcloud_waf_whiteblackip_rule_v1.rule_1 7117d38e-4c8f-4624-a505-bd96b97d024c ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the rule. ## Import WhiteBlackIP Rules can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import opentelekomcloud_waf_whiteblackip_rule_v1.rule_1 7117d38e-4c8f-4624-a505-bd96b97d024c ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
	}

	resourcesMap = map[string]int{

		"opentelekomcloud_antiddos_v1":                        0,
		"opentelekomcloud_as_configuration_v1":                1,
		"opentelekomcloud_as_group_v1":                        2,
		"opentelekomcloud_as_policy_v1":                       3,
		"opentelekomcloud_blockstorage_volume_v2":             4,
		"opentelekomcloud_cce_cluster_v3":                     5,
		"opentelekomcloud_cce_node_v3":                        6,
		"opentelekomcloud_ces_alarmrule":                      7,
		"opentelekomcloud_compute_server_v2":                  8,
		"opentelekomcloud_compute_floatingip_associate_v2":    9,
		"opentelekomcloud_compute_floatingip_v2":              10,
		"opentelekomcloud_compute_instance_v2":                11,
		"opentelekomcloud_compute_keypair_v2":                 12,
		"opentelekomcloud_compute_secgroup_v2":                13,
		"opentelekomcloud_compute_servergroup_v2":             14,
		"opentelekomcloud_compute_volume_attach_v2":           15,
		"opentelekomcloud_csbs_backup_policy_v1":              16,
		"opentelekomcloud_csbs_backup_v1":                     17,
		"opentelekomcloud_css_cluster_v1":                     18,
		"opentelekomcloud_cts_tracker_v1":                     19,
		"opentelekomcloud_dcs_instance_v1":                    20,
		"opentelekomcloud_deh_host_v1":                        21,
		"opentelekomcloud_dms_group_v1":                       22,
		"opentelekomcloud_dms_queue_v1":                       23,
		"opentelekomcloud_dns_recordset_v2":                   24,
		"opentelekomcloud_dns_zone_v2":                        25,
		"opentelekomcloud_elb_backend":                        26,
		"opentelekomcloud_elb_health":                         27,
		"opentelekomcloud_elb_listener":                       28,
		"opentelekomcloud_elb_loadbalancer":                   29,
		"opentelekomcloud_fw_firewall_group_v2":               30,
		"opentelekomcloud_fw_policy_v2":                       31,
		"opentelekomcloud_fw_rule_v2":                         32,
		"opentelekomcloud_identity_agency_v3":                 33,
		"opentelekomcloud_identity_group_membership_v3":       34,
		"opentelekomcloud_identity_group_v3":                  35,
		"opentelekomcloud_identity_project_v3":                36,
		"opentelekomcloud_identity_role_assignment_v3":        37,
		"opentelekomcloud_identity_role_v3":                   38,
		"opentelekomcloud_identity_user_v3":                   39,
		"opentelekomcloud_images_image_v2":                    40,
		"opentelekomcloud_kms_key_v1":                         41,
		"opentelekomcloud_lb_listener_v2":                     42,
		"opentelekomcloud_lb_loadbalancer_v2":                 43,
		"opentelekomcloud_lb_member_v2":                       44,
		"opentelekomcloud_lb_monitor_v2":                      45,
		"opentelekomcloud_lb_pool_v2":                         46,
		"opentelekomcloud_maas_task_v1":                       47,
		"opentelekomcloud_mrs_cluster_v1":                     48,
		"opentelekomcloud_mrs_job_v1":                         49,
		"opentelekomcloud_nat_gateway_v2":                     50,
		"opentelekomcloud_nat_snat_rule_v2":                   51,
		"opentelekomcloud_networking_floatingip_associate_v2": 52,
		"opentelekomcloud_networking_floatingip_v2":           53,
		"opentelekomcloud_networking_network_v2":              54,
		"opentelekomcloud_networking_port_v2":                 55,
		"opentelekomcloud_networking_router_interface_v2":     56,
		"opentelekomcloud_networking_router_route_v2":         57,
		"opentelekomcloud_networking_router_v2":               58,
		"opentelekomcloud_networking_secgroup_rule_v2":        59,
		"opentelekomcloud_networking_secgroup_v2":             60,
		"opentelekomcloud_networking_subnet_v2":               61,
		"opentelekomcloud_networking_vip_associate_v2":        62,
		"opentelekomcloud_networking_vip_v2":                  63,
		"opentelekomcloud_rds_instance_v1":                    64,
		"opentelekomcloud_rds_instance_v3":                    65,
		"opentelekomcloud_rds_parametergroup_v3":              66,
		"opentelekomcloud_rts_software_config_v1":             67,
		"opentelekomcloud_rts_software_deployment_v1":         68,
		"opentelekomcloud-resource-rts-stack-v1":              69,
		"opentelekomcloud_s3_bucket":                          70,
		"opentelekomcloud_s3-bucket-object":                   71,
		"opentelekomcloud_s3_bucket_policy":                   72,
		"opentelekomcloud_sfs_file_system_v2":                 73,
		"opentelekomcloud_smn_subscription_v2":                74,
		"opentelekomcloud_smn_topic_v2":                       75,
		"opentelekomcloud-vbs-backup-policy-v2":               76,
		"opentelekomcloud-vbs-backup-share-v2":                77,
		"opentelekomcloud-vbs-backup-v2":                      78,
		"opentelekomcloud_vpc_eip_v1":                         79,
		"opentelekomcloud_vpc_peering_connection_accepter_v2": 80,
		"opentelekomcloud_vpc_peering_connection_v2":          81,
		"opentelekomcloud_vpc_route_v2":                       82,
		"opentelekomcloud_vpc_subnet_v1":                      83,
		"opentelekomcloud_vpc_v1":                             84,
		"opentelekomcloud_waf_ccattackprotection_rule_v1":     85,
		"opentelekomcloud_waf_certificate_v1":                 86,
		"opentelekomcloud_waf_datamasking_rule_v1":            87,
		"opentelekomcloud_waf_domain_v1":                      88,
		"opentelekomcloud_waf_falsealarmmasking_rule_v1":      89,
		"opentelekomcloud_waf_policy_v1":                      90,
		"opentelekomcloud_waf_preciseprotection_rule_v1":      91,
		"opentelekomcloud_waf_webtamperprotection_rule_v1":    92,
		"opentelekomcloud_waf_whiteblackip_rule_v1":           93,
	}
)

func GetResource(r string) (*resouce.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs]
}
