package flexibleengine

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_antiddos_v1",
			Category:         "Antiddos Resources",
			ShortDescription: `Defends resources on the public cloud against network and monitors the service traffic from the Internet to ECSs, ELB instances, and BMSs to detect attack traffic in real time.`,
			Description:      ``,
			Icon:             "Security-Anti-DDoS.svg",
			Keywords: []string{
				"antiddos",
				"v1",
			},
			Arguments: []resource.Attribute{
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
					Description: `(Required) The ID corresponding to the Elastic IP Address (EIP) of a user. ## Attributes Reference All above argument parameters can be exported as attribute parameters. ## Import Antiddos can be imported using the floating_ip_id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_antiddos_v1.myantiddos c1881895-cdcb-4d23-96cb-032e6a3ee667 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_as_configuration_v1",
			Category:         "Auto Scaling Resources",
			ShortDescription: `Manages a V1 AS Configuration resource within flexibleengine.`,
			Description:      ``,
			Icon:             "Computing-AS.svg",
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
					Description: `(Required) The disk size. The unit is GB. The system disk size ranges from 1 to 32768, and the data disk size ranges from 10 to 32768.`,
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
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_as_group_v1",
			Category:         "Auto Scaling Resources",
			ShortDescription: `Manages a V1 Autoscaling Group resource within flexibleengine.`,
			Description:      ``,
			Icon:             "Computing-AS.svg",
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
					Description: `(Optional) The ELB listener IDs. The system supports up to three ELB listeners, the IDs of which are separated using a comma (,).`,
				},
				resource.Attribute{
					Name:        "lbaas_listeners",
					Description: `(Optional) An array of one or more enhanced load balancer. The system supports the binding of up to three load balancers. The field is alternative to lb_listener_id. The lbaas_listeners object structure is documented below.`,
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
					Description: `(Required) An array of ` + "`" + `one` + "`" + ` security group ID to associate with the group. The security_groups object structure is documented below.`,
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
					Name:        "pool_id",
					Description: `(Required) Specifies the backend ECS group ID.`,
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
			Type:             "flexibleengine_as_policy_v1",
			Category:         "Auto Scaling Resources",
			ShortDescription: `Manages a V1 AS Policy resource within flexibleengine.`,
			Description:      ``,
			Icon:             "Computing-AS.svg",
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
			Type:             "flexibleengine_blockstorage_volume_v2",
			Category:         "Block Storage Resources",
			ShortDescription: `Manages a V2 volume resource within FlexibleEngine.`,
			Description:      ``,
			Icon:             "Storage.svg",
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
					Description: `(Required) The size of the volume to create (in gigabytes).`,
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
					Description: `(Optional, Default:false) Specifies to delete all snapshots associated with the EVS disk.`,
				},
				resource.Attribute{
					Name:        "multiattach",
					Description: `(Optional) Specifies whether the EVS disk is shareable. ## Attributes Reference The following attributes are exported:`,
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
					Name:        "multiattach",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "attachment",
					Description: `If a volume is attached to an instance, this attribute will display the Attachment ID, Instance ID, and the Device as the Instance sees it. ## Import Volumes can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_blockstorage_volume_v2.volume_1 ea257959-eeb1-4c10-8d33-26f0409a755d ` + "`" + `` + "`" + `` + "`" + ``,
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
					Name:        "multiattach",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "attachment",
					Description: `If a volume is attached to an instance, this attribute will display the Attachment ID, Instance ID, and the Device as the Instance sees it. ## Import Volumes can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_blockstorage_volume_v2.volume_1 ea257959-eeb1-4c10-8d33-26f0409a755d ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_cce_cluster_v3",
			Category:         "CCE Resources",
			ShortDescription: `Provides Cloud Container Engine(CCE) resource.`,
			Description:      ``,
			Icon:             "Computing-CCE.svg",
			Keywords: []string{
				"cce",
				"cluster",
				"v3",
			},
			Arguments: []resource.Attribute{
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
					Description: `(Optional) For the cluster version, possible value is 'v1.11.7-r2','v1.13.10-r0' and 'v1.15.6-r1'. If this parameter is not set, the latest version will be used.`,
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
					Description: `(Required) The NETWORK ID of the subnet used to create the node. Changing this parameter will create a new cluster resource.`,
				},
				resource.Attribute{
					Name:        "highway_subnet_id",
					Description: `(Optional) The ID of the high speed network used to create bare metal nodes. Changing this parameter will create a new cluster resource.`,
				},
				resource.Attribute{
					Name:        "container_network_type",
					Description: `(Required) Container network parameters. Possible values:`,
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
					Description: `(Optional) Authentication mode of the cluster, possible values are x509 and rbac. Defaults to x509. Changing this parameter will create a new cluster resource.`,
				},
				resource.Attribute{
					Name:        "eip",
					Description: `(Optional) EIP address of the cluster. Changing this parameter will create a new cluster resource. ## Attributes Reference All above argument parameters can be exported as attribute parameters along with attribute reference.`,
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
					Name:        "internal_endpoint",
					Description: `The internal network address.`,
				},
				resource.Attribute{
					Name:        "external_endpoint",
					Description: `The external network address.`,
				},
				resource.Attribute{
					Name:        "external_apig_endpoint",
					Description: `The endpoint of the cluster to be accessed through API Gateway.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `Security group ID of the cluster.`,
				},
				resource.Attribute{
					Name:        "certificate_clusters.name",
					Description: `The cluster name.`,
				},
				resource.Attribute{
					Name:        "certificate_clusters.server",
					Description: `The server IP address.`,
				},
				resource.Attribute{
					Name:        "certificate_clusters.certificate_authority_data",
					Description: `The certificate data.`,
				},
				resource.Attribute{
					Name:        "certificate_users.name",
					Description: `The user name.`,
				},
				resource.Attribute{
					Name:        "certificate_users.client_certificate_data",
					Description: `The client certificate data.`,
				},
				resource.Attribute{
					Name:        "certificate_users.client_key_data",
					Description: `The client key data. ## Import Cluster can be imported using the cluster id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_cce_cluster_v3.cluster_1 4779ab1c-7c1a-44b1-a02e-93dfc361b32d ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Id of the cluster resource.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Cluster status information.`,
				},
				resource.Attribute{
					Name:        "internal_endpoint",
					Description: `The internal network address.`,
				},
				resource.Attribute{
					Name:        "external_endpoint",
					Description: `The external network address.`,
				},
				resource.Attribute{
					Name:        "external_apig_endpoint",
					Description: `The endpoint of the cluster to be accessed through API Gateway.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `Security group ID of the cluster.`,
				},
				resource.Attribute{
					Name:        "certificate_clusters.name",
					Description: `The cluster name.`,
				},
				resource.Attribute{
					Name:        "certificate_clusters.server",
					Description: `The server IP address.`,
				},
				resource.Attribute{
					Name:        "certificate_clusters.certificate_authority_data",
					Description: `The certificate data.`,
				},
				resource.Attribute{
					Name:        "certificate_users.name",
					Description: `The user name.`,
				},
				resource.Attribute{
					Name:        "certificate_users.client_certificate_data",
					Description: `The client certificate data.`,
				},
				resource.Attribute{
					Name:        "certificate_users.client_key_data",
					Description: `The client key data. ## Import Cluster can be imported using the cluster id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_cce_cluster_v3.cluster_1 4779ab1c-7c1a-44b1-a02e-93dfc361b32d ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_cce_nodes_v3",
			Category:         "CCE Resources",
			ShortDescription: `Add a node to a container cluster.`,
			Description:      ``,
			Icon:             "Computing-CCE.svg",
			Keywords: []string{
				"cce",
				"nodes",
				"v3",
			},
			Arguments: []resource.Attribute{
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
					Name:        "os",
					Description: `(Optional) Operating System of the node, possible values are EulerOS 2.2 and CentOS 7.6. Defaults to EulerOS 2.2. Changing this parameter will create a new resource.`,
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
					Name:        "preinstall",
					Description: `(Optional) Script required before installation. The input value can be a Base64 encoded string or not. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "postinstall",
					Description: `(Optional) Script required after installation. The input value can be a Base64 encoded string or not. Changing this parameter will create a new resource.`,
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
				resource.Attribute{
					Name:        "public_ip",
					Description: `Public IP of the CCE node.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "status",
					Description: `Node status information.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `Private IP of the CCE node.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `Public IP of the CCE node.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_ces-alarmrule",
			Category:         "CES Resources",
			ShortDescription: `Manages a V2 topic resource within FlexibleEngine.`,
			Description:      ``,
			Icon:             "Management&Deployment-CES.svg",
			Keywords: []string{
				"ces",
				"ces-alarmrule",
			},
			Arguments: []resource.Attribute{
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
					Description: `(Optional) Specifies the action list triggered by an alarm. The structure is described below.`,
				},
				resource.Attribute{
					Name:        "insufficientdata_actions",
					Description: `(Optional) Specifies the action list triggered by data insufficiency. The structure is described below.`,
				},
				resource.Attribute{
					Name:        "ok_actions",
					Description: `(Optional) Specifies the action list triggered by the clearing of an alarm. The structure is described below.`,
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
			Attributes: []resource.Attribute{
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
			Type:             "flexibleengine_compute_server_v2",
			Category:         "Compute Resources",
			ShortDescription: `Manages a BMS server resource within FlexibleEngine.`,
			Description:      ``,
			Icon:             "Computing-ECS.svg",
			Keywords: []string{
				"compute",
				"server",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to create the bms server instance. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new bms server.`,
				},
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
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_compute_floatingip_associate_v2",
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
					Description: `See Argument Reference above. ## Import This resource can be imported by specifying all three arguments, separated by a forward slash: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_compute_floatingip_associate_v2.fip_1 <floating_ip>/<instance_id>/<fixed_ip> ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `See Argument Reference above. ## Import This resource can be imported by specifying all three arguments, separated by a forward slash: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_compute_floatingip_associate_v2.fip_1 <floating_ip>/<instance_id>/<fixed_ip> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_compute_floatingip_v2",
			Category:         "Compute Resources",
			ShortDescription: `Manages a V2 floating IP resource within FlexibleEngine Nova (compute).`,
			Description:      ``,
			Keywords: []string{
				"compute",
				"floatingip",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 Compute client. A Compute client is needed to create a floating IP that can be used with a compute instance. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new floating IP (which may or may not have a different address).`,
				},
				resource.Attribute{
					Name:        "pool",
					Description: `(Optional) The name of the pool from which to obtain the floating IP. Default value is admin_external_net. Changing this creates a new floating IP. ## Attributes Reference The following attributes are exported:`,
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
					Name:        "fixed_ip",
					Description: `The fixed IP address corresponding to the floating IP.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `UUID of the compute instance associated with the floating IP. ## Import Floating IPs can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_compute_floatingip_v2.floatip_1 89c60255-9bd6-460c-822a-e2b959ede9d2 ` + "`" + `` + "`" + `` + "`" + ``,
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
					Name:        "fixed_ip",
					Description: `The fixed IP address corresponding to the floating IP.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `UUID of the compute instance associated with the floating IP. ## Import Floating IPs can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_compute_floatingip_v2.floatip_1 89c60255-9bd6-460c-822a-e2b959ede9d2 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_compute_instance_v2",
			Category:         "Compute Resources",
			ShortDescription: `Manages a V2 VM instance resource within FlexibleEngine.`,
			Description:      ``,
			Icon:             "Computing-ECS.svg",
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
					Description: `(Optional) Whether to try stop instance gracefully before destroying it, thus giving chance for guest OS daemons to stop correctly. If instance doesn't stop within timeout, it will be destroyed anyway.`,
				},
				resource.Attribute{
					Name:        "force_delete",
					Description: `(Optional) Whether to force the FlexibleEngine instance to be forcefully deleted. This is useful for environments that have reclaim / soft deletion enabled.`,
				},
				resource.Attribute{
					Name:        "auto_recovery",
					Description: `(Optional) Configures or deletes automatic recovery of an instance`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags key/value pairs to associate with the instance. The ` + "`" + `network` + "`" + ` block supports:`,
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
					Name:        "volume_type",
					Description: `(Optional) Currently, the value can be ` + "`" + `SSD` + "`" + ` (ultra-I/O disk type), ` + "`" + `SAS` + "`" + ` (high I/O disk type), or ` + "`" + `SATA` + "`" + ` (common I/O disk type)`,
				},
				resource.Attribute{
					Name:        "boot_index",
					Description: `(Optional) The boot index of the volume. It defaults to 0, which indicates that it's a system disk. Changing this creates a new server.`,
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
					Description: `Contains all instance metadata, even metadata not set by Terraform.`,
				},
				resource.Attribute{
					Name:        "auto_recovery",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `See Argument Reference above. ## Notes ### Multiple Ephemeral Disks It's possible to specify multiple ` + "`" + `block_device` + "`" + ` entries to create an instance with multiple ephemeral (local) disks. In order to create multiple ephemeral disks, the sum of the total amount of ephemeral space must be less than or equal to what the chosen flavor supports. The following example shows how to create an instance with multiple ephemeral disks: ` + "`" + `` + "`" + `` + "`" + ` resource "flexibleengine_compute_instance_v2" "foo" { name = "terraform-test" security_groups = ["default"] block_device { boot_index = 0 delete_on_termination = true destination_type = "local" source_type = "image" uuid = "<image uuid>" } block_device { boot_index = -1 delete_on_termination = true destination_type = "local" source_type = "blank" volume_size = 1 } block_device { boot_index = -1 delete_on_termination = true destination_type = "local" source_type = "blank" volume_size = 1 } } ` + "`" + `` + "`" + `` + "`" + ` ### Instances and Ports Neutron Ports are a great feature and provide a lot of functionality. However, there are some notes to be aware of when mixing Instances and Ports:`,
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
					Description: `Contains all instance metadata, even metadata not set by Terraform.`,
				},
				resource.Attribute{
					Name:        "auto_recovery",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `See Argument Reference above. ## Notes ### Multiple Ephemeral Disks It's possible to specify multiple ` + "`" + `block_device` + "`" + ` entries to create an instance with multiple ephemeral (local) disks. In order to create multiple ephemeral disks, the sum of the total amount of ephemeral space must be less than or equal to what the chosen flavor supports. The following example shows how to create an instance with multiple ephemeral disks: ` + "`" + `` + "`" + `` + "`" + ` resource "flexibleengine_compute_instance_v2" "foo" { name = "terraform-test" security_groups = ["default"] block_device { boot_index = 0 delete_on_termination = true destination_type = "local" source_type = "image" uuid = "<image uuid>" } block_device { boot_index = -1 delete_on_termination = true destination_type = "local" source_type = "blank" volume_size = 1 } block_device { boot_index = -1 delete_on_termination = true destination_type = "local" source_type = "blank" volume_size = 1 } } ` + "`" + `` + "`" + `` + "`" + ` ### Instances and Ports Neutron Ports are a great feature and provide a lot of functionality. However, there are some notes to be aware of when mixing Instances and Ports:`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_compute_interface_attach_v2",
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
			Type:             "flexibleengine_compute_keypair_v2",
			Category:         "Compute Resources",
			ShortDescription: `Manages a V2 keypair resource within FlexibleEngine.`,
			Description:      ``,
			Icon:             "key-pair.svg",
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
					Description: `(Required) A pregenerated OpenSSH-formatted public key. Changing this creates a new keypair.`,
				},
				resource.Attribute{
					Name:        "value_specs",
					Description: `(Optional) Map of additional options. ## Attributes Reference The following attributes are exported:`,
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
					Description: `See Argument Reference above. ## Import Keypairs can be imported using the ` + "`" + `name` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_compute_keypair_v2.my-keypair test-keypair ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `See Argument Reference above. ## Import Keypairs can be imported using the ` + "`" + `name` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_compute_keypair_v2.my-keypair test-keypair ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_compute_servergroup_v2",
			Category:         "Compute Resources",
			ShortDescription: `Manages a V2 Server Group resource within FlexibleEngine.`,
			Description:      ``,
			Icon:             "cloud-server-group.svg",
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
					Description: `The instances that are part of this server group. ## Import Server Groups can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_compute_servergroup_v2.test-sg 1bc30ee9-9d5b-4c30-bdd5-7f1e663f5edf ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `The instances that are part of this server group. ## Import Server Groups can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_compute_servergroup_v2.test-sg 1bc30ee9-9d5b-4c30-bdd5-7f1e663f5edf ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_compute_volume_attach_v2",
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
					Description: `See Argument Reference above. _NOTE_: The correctness of this information is dependent upon the hypervisor in use. In some cases, this should not be used as an authoritative piece of information. ## Import Volume Attachments can be imported using the Instance ID and Volume ID separated by a slash, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_compute_volume_attach_v2.va_1 89c60255-9bd6-460c-822a-e2b959ede9d2/45670584-225f-46c3-b33e-6707b589b666 ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `See Argument Reference above. _NOTE_: The correctness of this information is dependent upon the hypervisor in use. In some cases, this should not be used as an authoritative piece of information. ## Import Volume Attachments can be imported using the Instance ID and Volume ID separated by a slash, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_compute_volume_attach_v2.va_1 89c60255-9bd6-460c-822a-e2b959ede9d2/45670584-225f-46c3-b33e-6707b589b666 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_csbs_backup_policy_v1",
			Category:         "CSBS Resources",
			ShortDescription: `Provides an FlexibleEngine Backup Policy of Resource.`,
			Description:      ``,
			Icon:             "Storage-CSBS.svg",
			Keywords: []string{
				"csbs",
				"backup",
				"policy",
				"v1",
			},
			Arguments: []resource.Attribute{
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
					Description: `(Required) Specifies backup object name. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
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
					Description: `Specifies Scheduler type. ## Import Backup Policy can be imported using ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_csbs_backup_policy_v1.backup_policy_v1 7056d636-ac60-4663-8a6c-82d3c32c1c64 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
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
					Description: `Specifies Scheduler type. ## Import Backup Policy can be imported using ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_csbs_backup_policy_v1.backup_policy_v1 7056d636-ac60-4663-8a6c-82d3c32c1c64 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_csbs_backup_v1",
			Category:         "CSBS Resources",
			ShortDescription: `Provides an FlexibleEngine Backup of Resources.`,
			Description:      ``,
			Icon:             "Storage-CSBS.svg",
			Keywords: []string{
				"csbs",
				"backup",
				"v1",
			},
			Arguments: []resource.Attribute{
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
					Description: `Specifies image type. ## Import Backup can be imported using ` + "`" + `backup_record_id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_csbs_backup_v1.backup_v1.backup_v1 7056d636-ac60-4663-8a6c-82d3c32c1c64 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
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
					Description: `Specifies image type. ## Import Backup can be imported using ` + "`" + `backup_record_id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_csbs_backup_v1.backup_v1.backup_v1 7056d636-ac60-4663-8a6c-82d3c32c1c64 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_css_cluster_v1",
			Category:         "CSS Resources",
			ShortDescription: `CSS cluster management`,
			Description:      ``,
			Keywords: []string{
				"css",
				"cluster",
				"v1",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Cluster name. It contains 4 to 32 characters. Only letters, digits, hyphens (-), and underscores (_) are allowed. The value must start with a letter. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "engine_type",
					Description: `(Optional) Engine type. The default value is "elasticsearch". Currently, the value can only be "elasticsearch". Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "engine_version",
					Description: `(Required) Engine version. Versions 6.5.4 and 7.1.1 are supported. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "node_number",
					Description: `(Optional) Number of cluster instances. The value range is 1 to 32. Defaults to 1.`,
				},
				resource.Attribute{
					Name:        "node_config",
					Description: `(Required) Node configuration. Structure is documented below. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "backup_strategy",
					Description: `(Optional) Specifies the advanced backup policy. Structure is documented below. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) The key/value pairs to associate with the cluster. Changing this parameter will create a new resource. The ` + "`" + `node_config` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `(Optional) Availability zone (AZ). Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "flavor",
					Description: `(Required) Instance flavor name. For example: value range of flavor ess.spec-2u8g: 40 GB to 800 GB, value range of flavor ess.spec-4u16g: 40 GB to 1600 GB, value range of flavor ess.spec-8u32g: 80 GB to 3200 GB, value range of flavor ess.spec-16u64g: 100 GB to 6400 GB, value range of flavor ess.spec-32u128g: 100 GB to 10240 GB. Changing this parameter will create a new resource.`,
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
					Name:        "vpc_id",
					Description: `(Required) VPC ID, which is used for configuring cluster network. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Required) Subnet ID. All instances in a cluster must have the same subnet which should be configured with a`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `(Required) Security group ID. All instances in a cluster must have the same security group. Changing this parameter will create a new resource. The ` + "`" + `volume` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `(Required) Specifies volume size in GB, which must be a multiple of 4 and 10. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "volume_type",
					Description: `(Required) Specifies the volume type. Supported value: "COMMON": The SATA disk is used; "HIGH": The SAS disk is used; "ULTRAHIGH": The solid-state drive (SSD) is used. Changing this parameter will create a new resource. The ` + "`" + `backup_strategy` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "start_time",
					Description: `(Required) Specifies the time when a snapshot is automatically created everyday. Example value: "04:00 GMT+08:00".`,
				},
				resource.Attribute{
					Name:        "keep_days",
					Description: `(Optional) Specifies the number of days to retain the generated snapshots. Snapshots are reserved for seven days by default.`,
				},
				resource.Attribute{
					Name:        "prefix",
					Description: `(Optional) Specifies the prefix of the snapshot that is automatically created. The default value is "snapshot". ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "endpoint",
					Description: `Indicates the IP address and port number.`,
				},
				resource.Attribute{
					Name:        "created",
					Description: `Time when a cluster is created. The format is ISO8601: CCYY-MM-DDThh:mm:ss.`,
				},
				resource.Attribute{
					Name:        "nodes",
					Description: `List of node objects. Structure is documented below. The ` + "`" + `nodes` + "`" + ` block contains:`,
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
					Description: `Supported type: ess (indicating the Elasticsearch node). ## Timeouts This resource provides the following timeouts configuration options: - ` + "`" + `create` + "`" + ` - Default is 30 minute. - ` + "`" + `update` + "`" + ` - Default is 30 minute.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "endpoint",
					Description: `Indicates the IP address and port number.`,
				},
				resource.Attribute{
					Name:        "created",
					Description: `Time when a cluster is created. The format is ISO8601: CCYY-MM-DDThh:mm:ss.`,
				},
				resource.Attribute{
					Name:        "nodes",
					Description: `List of node objects. Structure is documented below. The ` + "`" + `nodes` + "`" + ` block contains:`,
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
					Description: `Supported type: ess (indicating the Elasticsearch node). ## Timeouts This resource provides the following timeouts configuration options: - ` + "`" + `create` + "`" + ` - Default is 30 minute. - ` + "`" + `update` + "`" + ` - Default is 30 minute.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_cts_tracker_v1",
			Category:         "CTS Resources",
			ShortDescription: `CTS tracker allows you to collect, store, and query cloud resource operation records and use these records for security analysis, compliance auditing, resource tracking, and fault locating.`,
			Description:      ``,
			Icon:             "Management&Deployment-CTS.svg",
			Keywords: []string{
				"cts",
				"tracker",
				"v1",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "bucket_name",
					Description: `(Required) The OBS bucket name for a tracker.`,
				},
				resource.Attribute{
					Name:        "file_prefix_name",
					Description: `(Optional) The prefix of a log that needs to be stored in an OBS bucket.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of a tracker. The value should be`,
				},
				resource.Attribute{
					Name:        "tracker_name",
					Description: `The tracker name. Currently, only tracker`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "tracker_name",
					Description: `The tracker name. Currently, only tracker`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_dcs_instance_v1",
			Category:         "DCS Resources",
			ShortDescription: `Manages a DCS instance in the flexibleengine DCS Service`,
			Description:      ``,
			Icon:             "Database-DCS.svg",
			Keywords: []string{
				"dcs",
				"instance",
				"v1",
			},
			Arguments: []resource.Attribute{
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
					Description: `(Required) Indicates the Cache capacity. Unit: GB. For a DCS Redis or Memcached instance in single-node or master/standby mode, the cache capacity can be 2 GB, 4 GB, 8 GB, 16 GB, 32 GB, or 64 GB. For a DCS Redis instance in cluster mode, the cache capacity can be 64, 128, 256, 512, or 1024 GB. Changing this creates a new instance.`,
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
					Description: `(Deprecated, Optional, conflict with ` + "`" + `network_id` + "`" + `) Network ID. Changing this creates a new instance.`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `(Optional, conflict with ` + "`" + `subnet_id` + "`" + `) Network ID. Changing this creates a new instance.`,
				},
				resource.Attribute{
					Name:        "available_zones",
					Description: `(Required) IDs or Names of the AZs where cache nodes reside. For details on how to query AZs, see Querying AZ Information. Changing this creates a new instance.`,
				},
				resource.Attribute{
					Name:        "product_id",
					Description: `(Optional) Product ID used to differentiate DCS instance types. For example now there are following values available: - dcs.memcached.master_standby-h, - dcs.memcached.master_standby-m, - dcs.memcached.single_node-h, - dcs.memcached.single_node-m, - dcs.master_standby-h, - dcs.cluster-h, - dcs.single_node-h, - redis.cluster.xu1.large.r1.4-h, - redis.cluster.xu1.large.r2.4-h, - redis.cluster.xu1.large.r1.8-h, - redis.cluster.xu1.large.r2.16-h, - redis.cluster.xu1.large.r1.16-h, - redis.cluster.xu1.large.r2.24-h, - redis.cluster.xu1.large.r2.32-h, - redis.cluster.xu1.large.r1.32-h, - redis.cluster.xu1.large.r2.48-h, - redis.cluster.xu1.large.r1.48-h - redis.cluster.xu1.large.r2.64-h - redis.cluster.xu1.large.r1.64-h - redis.cluster.xu1.large.r2.96-h - redis.cluster.xu1.large.r1.96-h - redis.cluster.xu1.large.r2.128-h - redis.cluster.xu1.large.r1.128-h - redis.cluster.xu1.large.r1.192-h - redis.cluster.xu1.large.r2.192-h - redis.cluster.xu1.large.r2.256-h - redis.cluster.xu1.large.r1.256-h - redis.cluster.xu1.large.r2.384-h - redis.cluster.xu1.large.r1.384-h - redis.cluster.xu1.large.r1.512-h - redis.cluster.xu1.large.r2.512-h ..... For the whole list and the specification of product id please check the DCS API DOC for querying: https://dcs.eu-west-0.prod-cloud-ocb.orange-business.com/v1.0/products Changing this creates a new instance.`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `(Optional) DCS instance specification code. For example now there are following values available: - dcs.single_node - dcs.master_standby - dcs.cluster`,
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
					Description: `(Optional) Retention time. Unit: day. Range: 1–7. Changing this creates a new instance.`,
				},
				resource.Attribute{
					Name:        "backup_type",
					Description: `(Optional) Backup type. Options: auto: automatic backup. manual: manual backup. Changing this creates a new instance.`,
				},
				resource.Attribute{
					Name:        "begin_at",
					Description: `(Optional) Time at which backup starts. "00:00-01:00" indicates that backup starts at 00:00:00. Changing this creates a new instance.`,
				},
				resource.Attribute{
					Name:        "period_type",
					Description: `(Optional) Interval at which backup is performed. Currently, only weekly backup is supported. Changing this creates a new instance.`,
				},
				resource.Attribute{
					Name:        "backup_at",
					Description: `(Optional) Day in a week on which backup starts. Range: 1–7. Where: 1 indicates Monday; 7 indicates Sunday. Changing this creates a new instance. ## Attributes Reference The following attributes are exported:`,
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
			Attributes: []resource.Attribute{
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
			Type:             "flexibleengine_dds_instance_v3",
			Category:         "DDS Resources",
			ShortDescription: `Manages dds instance resource within FlexibleEngine`,
			Description:      ``,
			Keywords: []string{
				"dds",
				"instance",
				"v3",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) Specifies the region of the DDS instance. Changing this creates a new instance.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Specifies the DB instance name. The DB instance name of the same type is unique in the same tenant. Changing this creates a new instance.`,
				},
				resource.Attribute{
					Name:        "datastore",
					Description: `(Required) Specifies database information. The structure is described below. Changing this creates a new instance.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `(Required) Specifies the ID of the availability zone. Changing this creates a new instance.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required) Specifies the VPC ID. For details about how to obtain this parameter value, see section "Virtual Private Cloud" in the Virtual Private Cloud API Reference. Changing this creates a new instance.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Required) Specifies the subnet Network ID. Changing this creates a new instance.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `(Required) Specifies the security group ID of the DDS instance. Changing this creates a new instance.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Required) Specifies the Administrator password of the database instance. Changing this creates a new instance.`,
				},
				resource.Attribute{
					Name:        "disk_encryption_id",
					Description: `(Required) Specifies the disk encryption ID of the instance. Changing this creates a new instance.`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `(Required) Specifies the mode of the database instance. Changing this creates a new instance.`,
				},
				resource.Attribute{
					Name:        "flavor",
					Description: `(Required) Specifies the flavors information. The structure is described below. Changing this creates a new instance.`,
				},
				resource.Attribute{
					Name:        "backup_strategy",
					Description: `(Optional) Specifies the advanced backup policy. The structure is described below. Changing this creates a new instance. The ` + "`" + `datastore` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) Specifies the DB engine. Only DDS-Community is supported now.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Required) Specifies the DB instance version. Only 3.4 is supported now.`,
				},
				resource.Attribute{
					Name:        "storage_engine",
					Description: `(Optional) Specifies the storage engine of the DB instance. Only wiredTiger is supported now. The ` + "`" + `flavor` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) Specifies the node type. Valid value: mongos, shard, config, replica.`,
				},
				resource.Attribute{
					Name:        "num",
					Description: `(Required) Specifies the node quantity. Valid value:`,
				},
				resource.Attribute{
					Name:        "storage",
					Description: `(Optional) Specifies the disk type. Valid value: ULTRAHIGH which indicates the type SSD.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `(Optional) Specifies the disk size. The value must be a multiple of 10. The unit is GB. This parameter is mandatory for nodes except mongos and invalid for mongos.`,
				},
				resource.Attribute{
					Name:        "spec_code",
					Description: `(Required) Specifies the resource specification code. Valid values: engine_name | type | vcpus | ram | speccode ---- | --- | --- DDS-Community | mongos | 1 | 4 | dds.mongodb.s3.medium.4.mongos DDS-Community | mongos | 2 | 8 | dds.mongodb.s3.large.4.mongos DDS-Community | mongos | 4 | 16 | dds.mongodb.s3.xlarge.4.mongos DDS-Community | mongos | 8 | 32 | dds.mongodb.s3.2xlarge.4.mongos DDS-Community | mongos | 16 | 64 | dds.mongodb.s3.4xlarge.4.mongos DDS-Community | shard | 1 | 4 | dds.mongodb.s3.medium.4.shard DDS-Community | shard | 2 | 8 | dds.mongodb.s3.large.4.shard DDS-Community | shard | 4 | 16 | dds.mongodb.s3.xlarge.4.shard DDS-Community | shard | 8 | 32 | dds.mongodb.s3.2xlarge.4.shard DDS-Community | shard | 16 | 64 | dds.mongodb.s3.4xlarge.4.shard DDS-Community | config | 2 | 4 | dds.mongodb.s3.large.2.config DDS-Community | replica | 1 | 4 | dds.mongodb.s3.medium.4.repset DDS-Community | replica | 2 | 8 | dds.mongodb.s3.large.4.repset DDS-Community | replica | 4 | 16 | dds.mongodb.s3.xlarge.4.repset DDS-Community | replica | 8 | 32 | dds.mongodb.s3.2xlarge.4.repset DDS-Community | replica | 16 | 64 | dds.mongodb.s3.4xlarge.4.repset The ` + "`" + `backup_strategy ` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "start_time",
					Description: `(Required) Specifies the backup time window. Automated backups will be triggered during the backup time window. The value cannot be empty. It must be a valid value in the "hh:mm-HH:MM" format. The current time is in the UTC format.`,
				},
				resource.Attribute{
					Name:        "keep_days",
					Description: `(Optional) Specifies the number of days to retain the generated backup files. The value range is from 0 to 732.`,
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
					Name:        "datastore",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
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
					Name:        "security_group_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "disk_encryption_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "flavor",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "backup_strategy",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "db_username",
					Description: `Indicates the DB Administator name.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Indicates the database port number. The port range is 2100 to 9500.`,
				},
				resource.Attribute{
					Name:        "nodes",
					Description: `Indicates the instance nodes information. Structure is documented below. The ` + "`" + `nodes` + "`" + ` block contains: - ` + "`" + `id` + "`" + ` - Indicates the node ID. - ` + "`" + `name` + "`" + ` - Indicates the node name. - ` + "`" + `role` + "`" + ` - Indicates the node role. - ` + "`" + `type` + "`" + ` - Indicates the node type. - ` + "`" + `private_ip` + "`" + ` - Indicates the private IP address of a node. This parameter is valid only for mongos nodes, replica set instances, and single node instances. - ` + "`" + `public_ip` + "`" + ` - Indicates the EIP that has been bound on a node. This parameter is valid only for mongos nodes of cluster instances, primary nodes and secondary nodes of replica set instances, and single node instances. - ` + "`" + `status` + "`" + ` - Indicates the node status.`,
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
					Name:        "datastore",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
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
					Name:        "security_group_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "disk_encryption_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "flavor",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "backup_strategy",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "db_username",
					Description: `Indicates the DB Administator name.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Indicates the database port number. The port range is 2100 to 9500.`,
				},
				resource.Attribute{
					Name:        "nodes",
					Description: `Indicates the instance nodes information. Structure is documented below. The ` + "`" + `nodes` + "`" + ` block contains: - ` + "`" + `id` + "`" + ` - Indicates the node ID. - ` + "`" + `name` + "`" + ` - Indicates the node name. - ` + "`" + `role` + "`" + ` - Indicates the node role. - ` + "`" + `type` + "`" + ` - Indicates the node type. - ` + "`" + `private_ip` + "`" + ` - Indicates the private IP address of a node. This parameter is valid only for mongos nodes, replica set instances, and single node instances. - ` + "`" + `public_ip` + "`" + ` - Indicates the EIP that has been bound on a node. This parameter is valid only for mongos nodes of cluster instances, primary nodes and secondary nodes of replica set instances, and single node instances. - ` + "`" + `status` + "`" + ` - Indicates the node status.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_dns_recordset_v2",
			Category:         "DNS Resources",
			ShortDescription: `Manages a DNS record set in the FlexibleEngine DNS Service`,
			Description:      ``,
			Icon:             "Network-DNS.svg",
			Keywords: []string{
				"dns",
				"recordset",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 DNS client. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new DNS record set.`,
				},
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
					Description: `(Optional) The type of record set. The options include ` + "`" + `A` + "`" + `, ` + "`" + `AAAA` + "`" + `, ` + "`" + `MX` + "`" + `, ` + "`" + `CNAME` + "`" + `, ` + "`" + `TXT` + "`" + `, ` + "`" + `NS` + "`" + `, ` + "`" + `SRV` + "`" + `, and ` + "`" + `PTR` + "`" + `. Changing this creates a new DNS record set.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `(Optional) The time to live (TTL) of the record set (in seconds). The value range is 300–2147483647. The default value is 300.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A description of the record set.`,
				},
				resource.Attribute{
					Name:        "records",
					Description: `(Required) An array of DNS records.`,
				},
				resource.Attribute{
					Name:        "value_specs",
					Description: `(Optional) Map of additional options. Changing this creates a new record set. ## Attributes Reference The following attributes are exported:`,
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
					Description: `See Argument Reference above. ## Import This resource can be imported by specifying the zone ID and recordset ID, separated by a forward slash. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_dns_recordset_v2.recordset_1 <zone_id>/<recordset_id> ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `See Argument Reference above. ## Import This resource can be imported by specifying the zone ID and recordset ID, separated by a forward slash. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_dns_recordset_v2.recordset_1 <zone_id>/<recordset_id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_dns_zone_v2",
			Category:         "DNS Resources",
			ShortDescription: `Manages a DNS zone in the FlexibleEngine DNS Service`,
			Description:      ``,
			Icon:             "zone.svg",
			Keywords: []string{
				"dns",
				"zone",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 Compute client. Keypairs are associated with accounts, but a Compute client is needed to create one. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new DNS zone. Changing this creates a new DNS zone.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the zone. Note the ` + "`" + `.` + "`" + ` at the end of the name. Changing this creates a new DNS zone.`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `(Optional) The email contact for the zone record.`,
				},
				resource.Attribute{
					Name:        "zone_type",
					Description: `(Optional) The type of zone. Can either be ` + "`" + `public` + "`" + ` or ` + "`" + `private` + "`" + `. Changing this creates a new DNS zone.`,
				},
				resource.Attribute{
					Name:        "router",
					Description: `(Optional) Router configuration block which is required if zone_type is private. The router structure is documented below.`,
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
					Name:        "value_specs",
					Description: `(Optional) Map of additional options. Changing this creates a new DNS zone. The ` + "`" + `router` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "router_id",
					Description: `(Required) The VPC UUID.`,
				},
				resource.Attribute{
					Name:        "router_region",
					Description: `(Required) The region of the VPC. ## Attributes Reference The following attributes are exported:`,
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
					Name:        "email",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "zone_type",
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
					Description: `An array of master DNS servers.`,
				},
				resource.Attribute{
					Name:        "value_specs",
					Description: `See Argument Reference above. ## Import This resource can be imported by specifying the zone ID: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_dns_zone_v2.zone_1 <zone_id> ` + "`" + `` + "`" + `` + "`" + ``,
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
					Name:        "email",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "zone_type",
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
					Description: `An array of master DNS servers.`,
				},
				resource.Attribute{
					Name:        "value_specs",
					Description: `See Argument Reference above. ## Import This resource can be imported by specifying the zone ID: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_dns_zone_v2.zone_1 <zone_id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_drs_replication_v2",
			Category:         "DRS Resources",
			ShortDescription: `Manages a V2 replication resource within FlexibleEngine.`,
			Description:      ``,
			Icon:             "drs.svg",
			Keywords: []string{
				"drs",
				"replication",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the EVS replication pair. The name can contain a maximum of 255 bytes.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description of the EVS replication pair. The description can contain a maximum of 255 bytes.`,
				},
				resource.Attribute{
					Name:        "volume_ids",
					Description: `(Required) An array of one or more IDs of the EVS disks used to create the EVS replication pair.`,
				},
				resource.Attribute{
					Name:        "priority_station",
					Description: `(Required) The primary AZ of the EVS replication pair. That is the AZ where the production disk belongs.`,
				},
				resource.Attribute{
					Name:        "replication_model",
					Description: `(Optional) The type of the EVS replication pair. Currently only type hypermetro is supported. ## Attributes Reference The following attributes are exported:`,
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
					Name:        "volume_ids",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "priority_station",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "replication_model",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the EVS replication pair.`,
				},
				resource.Attribute{
					Name:        "replication_consistency_group_id",
					Description: `The ID of the replication consistency group where the EVS replication pair belongs.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `The creation time of the EVS replication pair.`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `The update time of the EVS replication pair.`,
				},
				resource.Attribute{
					Name:        "replication_status",
					Description: `The replication status of the EVS replication pair.`,
				},
				resource.Attribute{
					Name:        "progress",
					Description: `The synchronization progress of the EVS replication pair. Unit: %.`,
				},
				resource.Attribute{
					Name:        "failure_detail",
					Description: `The returned error code if the EVS replication pair status is error.`,
				},
				resource.Attribute{
					Name:        "record_metadata",
					Description: `The metadata of the EVS replication pair.`,
				},
				resource.Attribute{
					Name:        "fault_level",
					Description: `The fault level of the EVS replication pair.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "volume_ids",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "priority_station",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "replication_model",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the EVS replication pair.`,
				},
				resource.Attribute{
					Name:        "replication_consistency_group_id",
					Description: `The ID of the replication consistency group where the EVS replication pair belongs.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `The creation time of the EVS replication pair.`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `The update time of the EVS replication pair.`,
				},
				resource.Attribute{
					Name:        "replication_status",
					Description: `The replication status of the EVS replication pair.`,
				},
				resource.Attribute{
					Name:        "progress",
					Description: `The synchronization progress of the EVS replication pair. Unit: %.`,
				},
				resource.Attribute{
					Name:        "failure_detail",
					Description: `The returned error code if the EVS replication pair status is error.`,
				},
				resource.Attribute{
					Name:        "record_metadata",
					Description: `The metadata of the EVS replication pair.`,
				},
				resource.Attribute{
					Name:        "fault_level",
					Description: `The fault level of the EVS replication pair.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_drs_replicationconsistencygroup_v2",
			Category:         "DRS Resources",
			ShortDescription: `Manages a V2 replicationconsistencygroup resource within FlexibleEngine.`,
			Description:      ``,
			Icon:             "drs.svg",
			Keywords: []string{
				"drs",
				"replicationconsistencygroup",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the replication consistency group. The name can contain a maximum of 255 bytes.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description of the replication consistency group. The description can contain a maximum of 255 bytes.`,
				},
				resource.Attribute{
					Name:        "replication_ids",
					Description: `(Required) An array of one or more IDs of the EVS replication pairs used to create the replication consistency group.`,
				},
				resource.Attribute{
					Name:        "priority_station",
					Description: `(Required) The primary AZ of the replication consistency group. That is the AZ where the production disk belongs.`,
				},
				resource.Attribute{
					Name:        "replication_model",
					Description: `(Optional) The type of the created replication consistency group. Currently only type hypermetro is supported. ## Attributes Reference The following attributes are exported:`,
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
					Name:        "replication_ids",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "priority_station",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "replication_model",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the replication consistency group.`,
				},
				resource.Attribute{
					Name:        "replication_status",
					Description: `The replication status of the replication consistency group.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `The creation time of the replication consistency group.`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `The update time of the replication consistency group.`,
				},
				resource.Attribute{
					Name:        "failure_detail",
					Description: `The returned error code if the replication consistency group status is error.`,
				},
				resource.Attribute{
					Name:        "fault_level",
					Description: `The fault level of the replication consistency group.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "replication_ids",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "priority_station",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "replication_model",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the replication consistency group.`,
				},
				resource.Attribute{
					Name:        "replication_status",
					Description: `The replication status of the replication consistency group.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `The creation time of the replication consistency group.`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `The update time of the replication consistency group.`,
				},
				resource.Attribute{
					Name:        "failure_detail",
					Description: `The returned error code if the replication consistency group status is error.`,
				},
				resource.Attribute{
					Name:        "fault_level",
					Description: `The fault level of the replication consistency group.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_dws_cluster_v1",
			Category:         "Data Warehouse Resources",
			ShortDescription: `Manages a DWS cluster resource within FlexibleEngine.`,
			Description:      ``,
			Icon:             "Data Analysis&AI-DWS.svg",
			Keywords: []string{
				"data",
				"warehouse",
				"dws",
				"cluster",
				"v1",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "availability_zone",
					Description: `(Optional) AZ in a cluster`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Cluster name, which must be unique and contains 4 to 64 characters, which consist of letters, digits, hyphens (-), or underscores (_) only and must start with a letter.`,
				},
				resource.Attribute{
					Name:        "node_type",
					Description: `(Required) Node type.`,
				},
				resource.Attribute{
					Name:        "number_of_node",
					Description: `(Required) Number of nodes in a cluster. The value ranges from 3 to 32.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional) Service port of a cluster (8000 to 10000). The default value is 8000.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `(Optional) Public IP address. If the value is not specified, public connection is not used by default.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `(Required) ID of a security group. The ID is used for configuring cluster network.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Required) Subnet ID, which is used for configuring cluster network.`,
				},
				resource.Attribute{
					Name:        "user_name",
					Description: `(Required) Administrator username for logging in to a data warehouse cluster The administrator username must: Consist of lowercase letters, digits, or underscores. Start with a lowercase letter or an underscore. Contain 1 to 63 characters. Cannot be a keyword of the DWS database.`,
				},
				resource.Attribute{
					Name:        "user_pwd",
					Description: `(Required) Administrator password for logging in to a data warehouse cluster A password must conform to the following rules: Contains 8 to 32 characters. Cannot be the same as the username or the username written in reverse order. Contains three types of the following: Lowercase letters Uppercase letters Digits Special characters ~!@#%^&`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required) VPC ID, which is used for configuring cluster network. The ` + "`" + `public_ip` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "eip_id",
					Description: `(Optional) EIP ID`,
				},
				resource.Attribute{
					Name:        "public_bind_type",
					Description: `(Optional) Binding type of an EIP. The value can be either of the following: auto_assign not_use bind_existing The default value is not_use. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "number_of_node",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "user_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "node_type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "created",
					Description: `Cluster creation time. The format is ISO8601:YYYY-MM-DDThh:mm:ssZ.`,
				},
				resource.Attribute{
					Name:        "endpoints",
					Description: `View the private network connection information about the cluster.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Cluster ID`,
				},
				resource.Attribute{
					Name:        "public_endpoints",
					Description: `Public network connection information about the cluster. If the value is not specified, the public network connection information is not used by default.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Cluster status, which can be one of the following: CREATING AVAILABLE UNAVAILABLE CREATION FAILED`,
				},
				resource.Attribute{
					Name:        "sub_status",
					Description: `Sub-status of clusters in the AVAILABLE state. The value can be one of the following: NORMAL READONLY REDISTRIBUTING REDISTRIBUTION-FAILURE UNBALANCED UNBALANCED | READONLY DEGRADED DEGRADED | READONLY DEGRADED | UNBALANCED UNBALANCED | REDISTRIBUTING UNBALANCED | REDISTRIBUTION-FAILURE READONLY | REDISTRIBUTION-FAILURE UNBALANCED | READONLY | REDISTRIBUTION-FAILURE DEGRADED | REDISTRIBUTION-FAILURE DEGRADED | UNBALANCED | REDISTRIBUTION-FAILURE DEGRADED | UNBALANCED | READONLY | REDISTRIBUTION-FAILURE DEGRADED | UNBALANCED | READONLY`,
				},
				resource.Attribute{
					Name:        "task_status",
					Description: `Cluster management task. The value can be one of the following: RESTORING SNAPSHOTTING GROWING REBOOTING SETTING_CONFIGURATION CONFIGURING_EXT_DATASOURCE DELETING_EXT_DATASOURCE REBOOT_FAILURE RESIZE_FAILURE`,
				},
				resource.Attribute{
					Name:        "updated",
					Description: `Last modification time of a cluster. The format is ISO8601:YYYY-MM-DDThh:mm:ssZ.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `Data warehouse version The ` + "`" + `endpoints` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "connect_info",
					Description: `Private network connection information`,
				},
				resource.Attribute{
					Name:        "jdbc_url",
					Description: `JDBC URL. The following is the default format: jdbc:postgresql://< connect_info>/<YOUR_DATABASE_NAME> The ` + "`" + `public_endpoints` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "public_connect_info",
					Description: `Public network connection information`,
				},
				resource.Attribute{
					Name:        "jdbc_url",
					Description: `JDBC URL. The following is the default format: jdbc:postgresql://< public_connect_info>/<YOUR_DATABASE_NAME>`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "number_of_node",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "user_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "node_type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "created",
					Description: `Cluster creation time. The format is ISO8601:YYYY-MM-DDThh:mm:ssZ.`,
				},
				resource.Attribute{
					Name:        "endpoints",
					Description: `View the private network connection information about the cluster.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Cluster ID`,
				},
				resource.Attribute{
					Name:        "public_endpoints",
					Description: `Public network connection information about the cluster. If the value is not specified, the public network connection information is not used by default.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Cluster status, which can be one of the following: CREATING AVAILABLE UNAVAILABLE CREATION FAILED`,
				},
				resource.Attribute{
					Name:        "sub_status",
					Description: `Sub-status of clusters in the AVAILABLE state. The value can be one of the following: NORMAL READONLY REDISTRIBUTING REDISTRIBUTION-FAILURE UNBALANCED UNBALANCED | READONLY DEGRADED DEGRADED | READONLY DEGRADED | UNBALANCED UNBALANCED | REDISTRIBUTING UNBALANCED | REDISTRIBUTION-FAILURE READONLY | REDISTRIBUTION-FAILURE UNBALANCED | READONLY | REDISTRIBUTION-FAILURE DEGRADED | REDISTRIBUTION-FAILURE DEGRADED | UNBALANCED | REDISTRIBUTION-FAILURE DEGRADED | UNBALANCED | READONLY | REDISTRIBUTION-FAILURE DEGRADED | UNBALANCED | READONLY`,
				},
				resource.Attribute{
					Name:        "task_status",
					Description: `Cluster management task. The value can be one of the following: RESTORING SNAPSHOTTING GROWING REBOOTING SETTING_CONFIGURATION CONFIGURING_EXT_DATASOURCE DELETING_EXT_DATASOURCE REBOOT_FAILURE RESIZE_FAILURE`,
				},
				resource.Attribute{
					Name:        "updated",
					Description: `Last modification time of a cluster. The format is ISO8601:YYYY-MM-DDThh:mm:ssZ.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `Data warehouse version The ` + "`" + `endpoints` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "connect_info",
					Description: `Private network connection information`,
				},
				resource.Attribute{
					Name:        "jdbc_url",
					Description: `JDBC URL. The following is the default format: jdbc:postgresql://< connect_info>/<YOUR_DATABASE_NAME> The ` + "`" + `public_endpoints` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "public_connect_info",
					Description: `Public network connection information`,
				},
				resource.Attribute{
					Name:        "jdbc_url",
					Description: `JDBC URL. The following is the default format: jdbc:postgresql://< public_connect_info>/<YOUR_DATABASE_NAME>`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_elb_backend",
			Category:         "Elastic Load Balancer Resources",
			ShortDescription: `Manages an elastic loadbalancer backend resource within FlexibleEngine.`,
			Description:      ``,
			Icon:             "Network-ELB.svg",
			Keywords: []string{
				"elastic",
				"load",
				"balancer",
				"elb",
				"backend",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "listener_id",
					Description: `(Required) Specifies the listener ID.`,
				},
				resource.Attribute{
					Name:        "server_id",
					Description: `(Required) Specifies the backend member ID.`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `(Required) Specifies the private IP address of the backend member. ## Attributes Reference The following attributes are exported:`,
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
			Attributes: []resource.Attribute{
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
			Type:             "flexibleengine_elb_health",
			Category:         "Elastic Load Balancer Resources",
			ShortDescription: `Manages an elastic loadbalancer health resource within FlexibleEngine.`,
			Description:      ``,
			Icon:             "Network-ELB.svg",
			Keywords: []string{
				"elastic",
				"load",
				"balancer",
				"elb",
				"health",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to create the elb health. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new elb health.`,
				},
				resource.Attribute{
					Name:        "listener_id",
					Description: `(Required) Specifies the ID of the listener to which the health check task belongs.`,
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
					Name:        "region",
					Description: `See Argument Reference above.`,
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
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
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
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_elb_listener",
			Category:         "Elastic Load Balancer Resources",
			ShortDescription: `Manages an elastic loadbalancer listener resource within FlexibleEngine.`,
			Description:      ``,
			Icon:             "Network-ELB.svg",
			Keywords: []string{
				"elastic",
				"load",
				"balancer",
				"elb",
				"listener",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to create the elb listener. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new elb listener.`,
				},
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
					Description: `(Required) Specifies the ID of the load balancer to which the listener belongs.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Required) Specifies the listening protocol used for layer 4 or 7. The value can be HTTP, TCP, HTTPS, or UDP.`,
				},
				resource.Attribute{
					Name:        "protocol_port",
					Description: `(Required) Specifies the listening port. The value ranges from 1 to 65535.`,
				},
				resource.Attribute{
					Name:        "backend_protocol",
					Description: `(Required) Specifies the backend protocol. If the value of protocol is UDP, the value of this parameter can only be UDP. The value can be HTTP, TCP, or UDP.`,
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
					Description: `(Optional) Specifies whether to enable sticky session. The value can be true or false. The Sticky session is enabled when the value is true, and is disabled when the value is false. If the value of protocol is HTTP, HTTPS, or TCP, and the value of lb_algorithm is not roundrobin, the value of this parameter can only be false.`,
				},
				resource.Attribute{
					Name:        "session_sticky_type",
					Description: `(Optional) Specifies the cookie processing method. The value is insert. insert indicates that the cookie is inserted by the load balancer. This parameter is valid when protocol is set to HTTP, and session_sticky to true. The default value is insert. This parameter is invalid when protocol is set to TCP or UDP, which means the parameter is empty.`,
				},
				resource.Attribute{
					Name:        "cookie_timeout",
					Description: `(Optional) Specifies the cookie timeout period (minutes). This parameter is valid when protocol is set to HTTP, session_sticky to true, and session_sticky_type to insert. This parameter is invalid when protocol is set to TCP or UDP. The value ranges from 1 to 1440.`,
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
					Description: `(Optional) Specifies the ID of the SSL certificate used for security authentication when HTTPS is used to make API calls. This parameter is mandatory if the value of protocol is HTTPS. The value can be obtained by viewing the details of the SSL certificate.`,
				},
				resource.Attribute{
					Name:        "udp_timeout",
					Description: `(Optional) Specifies the UDP timeout duration (minutes). This parameter is valid when protocol is set to UDP. The value ranges from 1 to 1440.`,
				},
				resource.Attribute{
					Name:        "ssl_protocols",
					Description: `(Optional) Specifies the SSL protocol standard supported by a tracker, which is used for enabling specified encryption protocols. This parameter is valid only when the value of protocol is set to HTTPS. The value is TLSv1.2 or TLSv1.2 TLSv1.1 TLSv1. The default value is TLSv1.2.`,
				},
				resource.Attribute{
					Name:        "ssl_ciphers",
					Description: `(Optional) Specifies the cipher suite of an encryption protocol. This parameter is valid only when the value of protocol is set to HTTPS. The value is Default, Extended, or Strict. The default value is Default. The value can only be set to Extended if the value of ssl_protocols is set to TLSv1.2 TLSv1.1 TLSv1. ## Attributes Reference The following attributes are exported:`,
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
			Type:             "flexibleengine_elb_loadbalancer",
			Category:         "Elastic Load Balancer Resources",
			ShortDescription: `Manages an elastic loadbalancer resource within FlexibleEngine.`,
			Description:      ``,
			Icon:             "Network-ELB.svg",
			Keywords: []string{
				"elastic",
				"load",
				"balancer",
				"elb",
				"loadbalancer",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to create the loadbalancer. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new loadbalancer.`,
				},
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
					Description: `(Required) Specifies the VPC ID.`,
				},
				resource.Attribute{
					Name:        "bandwidth",
					Description: `(Optional) Specifies the bandwidth (Mbit/s). This parameter is mandatory when type is set to External, and it is invalid when type is set to Internal. The value ranges from 1 to 300.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) Specifies the load balancer type. The value can be Internal or External.`,
				},
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `(Required) Specifies the status of the load balancer. Value range: 0 or false: indicates that the load balancer is stopped. Only tenants are allowed to enter these two values. 1 or true: indicates that the load balancer is running properly. 2 or false: indicates that the load balancer is frozen. Only tenants are allowed to enter these two values.`,
				},
				resource.Attribute{
					Name:        "vip_subnet_id",
					Description: `(Optional) Specifies the ID of the private network to be added. This parameter is mandatory when type is set to Internal, and it is invalid when type is set to External.`,
				},
				resource.Attribute{
					Name:        "az",
					Description: `(Optional) Specifies the ID of the availability zone (AZ). This parameter is mandatory when type is set to Internal, and it is invalid when type is set to External.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `(Optional) Specifies the security group ID. The value is a string of 1 to 200 characters that consists of uppercase and lowercase letters, digits, and hyphens (-). This parameter is mandatory only when type is set to Internal.`,
				},
				resource.Attribute{
					Name:        "vip_address",
					Description: `(Optional) Specifies the IP address provided by ELB. When type is set to External, the value of this parameter is the elastic IP address. When type is set to Internal, the value of this parameter is the private network IP address. You can select an existing elastic IP address and create a public network load balancer. When this parameter is configured, parameter bandwidth is invalid.`,
				},
				resource.Attribute{
					Name:        "tenantid",
					Description: `(Optional) Specifies the tenant ID. This parameter is mandatory only when type is set to Internal. ## Attributes Reference The following attributes are exported:`,
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
			Type:             "flexibleengine_fw_firewall_group_v2",
			Category:         "Firewall Resources",
			ShortDescription: `Manages a v2 firewall group resource within FlexibleEngine.`,
			Description:      ``,
			Icon:             "firewall.svg",
			Keywords: []string{
				"firewall",
				"fw",
				"group",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the v2 networking client. A networking client is needed to create a firewall group. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new firewall group.`,
				},
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
					Name:        "region",
					Description: `See Argument Reference above.`,
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
					Description: `See Argument Reference above. ## Import Firewall Groups can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_fw_firewall_group_v2.firewall_group_1 c9e39fb2-ce20-46c8-a964-25f3898c7a97 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
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
					Description: `See Argument Reference above. ## Import Firewall Groups can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_fw_firewall_group_v2.firewall_group_1 c9e39fb2-ce20-46c8-a964-25f3898c7a97 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_fw_policy_v2",
			Category:         "Firewall Resources",
			ShortDescription: `Manages a v2 firewall policy resource within FlexibleEngine.`,
			Description:      ``,
			Icon:             "firewall.svg",
			Keywords: []string{
				"firewall",
				"fw",
				"policy",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the v2 networking client. A networking client is needed to create a firewall policy. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new firewall policy.`,
				},
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
					Name:        "audited",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "shared",
					Description: `See Argument Reference above. ## Import Firewall Policies can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_fw_policy_v2.policy_1 07f422e6-c596-474b-8b94-fe2c12506ce0 ` + "`" + `` + "`" + `` + "`" + ``,
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
					Name:        "audited",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "shared",
					Description: `See Argument Reference above. ## Import Firewall Policies can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_fw_policy_v2.policy_1 07f422e6-c596-474b-8b94-fe2c12506ce0 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_fw_rule_v2",
			Category:         "Firewall Resources",
			ShortDescription: `Manages a v2 firewall group rule resource within FlexibleEngine.`,
			Description:      ``,
			Icon:             "firewall.svg",
			Keywords: []string{
				"firewall",
				"fw",
				"rule",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the v2 networking client. A Compute client is needed to create a firewall rule. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new firewall rule.`,
				},
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
					Description: `See Argument Reference above. ## Import Firewall Rules can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_fw_rule_v2.rule_1 8dbc0c28-e49c-463f-b712-5c5d1bbac327 ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `See Argument Reference above. ## Import Firewall Rules can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_fw_rule_v2.rule_1 8dbc0c28-e49c-463f-b712-5c5d1bbac327 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_identity_agency_v3",
			Category:         "Identity Resources",
			ShortDescription: `Manages an agency resource within FlexibleEngine.`,
			Description:      ``,
			Keywords: []string{
				"identity",
				"agency",
				"v3",
			},
			Arguments: []resource.Attribute{
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
					Name:        "id",
					Description: `The agency ID.`,
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
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The agency ID.`,
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
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_identity_group_membership_v3",
			Category:         "Identity Resources",
			ShortDescription: `Manages the membership combine User Group resource and User resource within FlexibleEngine IAM service.`,
			Description:      ``,
			Keywords: []string{
				"identity",
				"group",
				"membership",
				"v3",
			},
			Arguments: []resource.Attribute{
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
			Attributes: []resource.Attribute{
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
			Type:             "flexibleengine_identity_group_v3",
			Category:         "Identity Resources",
			ShortDescription: `Manages a User Group resource within FlexibleEngine IAM service.`,
			Description:      ``,
			Keywords: []string{
				"identity",
				"group",
				"v3",
			},
			Arguments: []resource.Attribute{
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
					Description: `See Argument Reference above. ## Import Groups can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_identity_group_v3.group_1 89c60255-9bd6-460c-822a-e2b959ede9d2 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "domain_id",
					Description: `See Argument Reference above. ## Import Groups can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_identity_group_v3.group_1 89c60255-9bd6-460c-822a-e2b959ede9d2 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_identity_role_assignment_v3",
			Category:         "Identity Resources",
			ShortDescription: `Manages a V3 Policy assignment within FlexibleEngine IAM Service.`,
			Description:      ``,
			Keywords: []string{
				"identity",
				"role",
				"assignment",
				"v3",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "role_id",
					Description: `(Required) The role to assign.`,
				},
				resource.Attribute{
					Name:        "domain_id",
					Description: `(Optional; Required if ` + "`" + `project_id` + "`" + ` is empty) The domain to assign the role in.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Optional; Required if ` + "`" + `domain_id` + "`" + ` is empty) The project to assign the role in.`,
				},
				resource.Attribute{
					Name:        "group_id",
					Description: `(Optional; Required if ` + "`" + `user_id` + "`" + ` is empty) The group to assign the role in.`,
				},
				resource.Attribute{
					Name:        "user_id",
					Description: `(Optional; Required if ` + "`" + `group_id` + "`" + ` is empty) The user to assign the role in. ## Attributes Reference The following attributes are exported:`,
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
			Attributes: []resource.Attribute{
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
			Type:             "flexibleengine_identity_role_v3",
			Category:         "Identity Resources",
			ShortDescription: `custom role management in FlexibleEngine`,
			Description:      ``,
			Keywords: []string{
				"identity",
				"role",
				"v3",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Specify the name of a role. The value cannot exceed 64 characters.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Required) Specify the description of a role. The value cannot exceed 256 characters.`,
				},
				resource.Attribute{
					Name:        "scope",
					Description: `(Required) Specify the scope layer of a role. The value supports: - domain - A role is displayed at the domain layer. - project - A role is displayed at the project layer.`,
				},
				resource.Attribute{
					Name:        "policy",
					Description: `(Required) The policy field contains the ` + "`" + `effect` + "`" + ` and ` + "`" + `action` + "`" + ` elements. Effect indicates whether the policy allows or denies access. Action indicates authorization items. The number of policy cannot exceed 8. Structure is documented below. The ` + "`" + `policy` + "`" + ` block supports:`,
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
					Description: `ID of the domain to which a role belongs ## Import Role can be imported using the following format: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_identity_role_v3.default {{ resource id}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "catalog",
					Description: `Directory where a role locates`,
				},
				resource.Attribute{
					Name:        "domain_id",
					Description: `ID of the domain to which a role belongs ## Import Role can be imported using the following format: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_identity_role_v3.default {{ resource id}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_identity_user_v3",
			Category:         "Identity Resources",
			ShortDescription: `Manages a User resource within FlexibleEngine IAM service.`,
			Description:      ``,
			Keywords: []string{
				"identity",
				"user",
				"v3",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the user. The user name consists of 5 to 32 characters. It can contain only uppercase letters, lowercase letters, digits, spaces, and special characters (-_) and cannot start with a digit.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A description of the user.`,
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
					Description: `See Argument Reference above. ## Import Users can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_identity_user_v3.user_1 89c60255-9bd6-460c-822a-e2b959ede9d2 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "domain_id",
					Description: `See Argument Reference above. ## Import Users can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_identity_user_v3.user_1 89c60255-9bd6-460c-822a-e2b959ede9d2 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_images_image_v2",
			Category:         "Images Resources",
			ShortDescription: `Manages a V2 Image resource within FlexibleEngine Glance.`,
			Description:      ``,
			Keywords: []string{
				"images",
				"image",
				"v2",
			},
			Arguments: []resource.Attribute{
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
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 Glance client. A Glance client is needed to create an Image that can be used with a compute instance. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new Image.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) The tags of the image. It must be a list of strings. At this time, it is not possible to delete all tags of an image.`,
				},
				resource.Attribute{
					Name:        "visibility",
					Description: `(Optional) The visibility of the image. Must be one of "public", "private", "community", or "shared". The ability to set the visibility depends upon the configuration of the FlexibleEngine cloud. Note: The ` + "`" + `properties` + "`" + ` attribute handling in the gophercloud library is currently buggy and needs to be fixed before being implemented in this resource. ## Attributes Reference The following attributes are exported:`,
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
					Description: `The id of the flexibleengine user who owns the image.`,
				},
				resource.Attribute{
					Name:        "protected",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "region",
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
					Description: `See Argument Reference above. ## Import Images can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_images_image_v2.rancheros 89c60255-9bd6-460c-822a-e2b959ede9d2 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
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
					Description: `The id of the flexibleengine user who owns the image.`,
				},
				resource.Attribute{
					Name:        "protected",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "region",
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
					Description: `See Argument Reference above. ## Import Images can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_images_image_v2.rancheros 89c60255-9bd6-460c-822a-e2b959ede9d2 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_kms_key_v1",
			Category:         "KMS Resources",
			ShortDescription: `Manages a V1 key resource within KMS.`,
			Description:      ``,
			Icon:             "Security-KMS.svg",
			Keywords: []string{
				"kms",
				"key",
				"v1",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key_alias",
					Description: `(Required) The alias in which to create the key. It is required when we create a new key. Changing this updates the alias of key.`,
				},
				resource.Attribute{
					Name:        "key_description",
					Description: `(Optional) The description of the key as viewed in FlexibleEngine console. Changing this updates the description of key.`,
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
					Description: `See Argument Reference above. ## Import KMS Keys can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_kms_key_v1.key_1 7056d636-ac60-4663-8a6c-82d3c32c1c64 ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `See Argument Reference above. ## Import KMS Keys can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_kms_key_v1.key_1 7056d636-ac60-4663-8a6c-82d3c32c1c64 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_lb_certificate_v2",
			Category:         "Enhanced Load Balancer Resources",
			ShortDescription: `Manages a V2 certificate resource within FlexibleEngine.`,
			Description:      ``,
			Keywords: []string{
				"enhanced",
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
			Type:             "flexibleengine_lb_l7policy_v2",
			Category:         "Enhanced Load Balancer Resources",
			ShortDescription: `Manages a V2 L7 Policy resource within FlexibleEngine.`,
			Description:      ``,
			Keywords: []string{
				"enhanced",
				"load",
				"balancer",
				"lb",
				"l7policy",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 Networking client. A Networking client is needed to create an . If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new L7 Policy.`,
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
					Description: `(Required) The L7 Policy action - can either be REDIRECT\_TO\_POOL, or REDIRECT\_TO\_LISTENER. Changing this creates a new L7 Policy.`,
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
					Description: `(Optional) Requests matching this policy will be redirected to the pool with this ID. Only valid if action is REDIRECT\_TO\_POOL.`,
				},
				resource.Attribute{
					Name:        "redirect_listener_id",
					Description: `(Optional) Requests matching this policy will be redirected to the listener with this ID. Only valid if action is REDIRECT\_TO\_LISTENER.`,
				},
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `(Optional) The administrative state of the L7 Policy. This value can only be true (UP). ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique ID for the L7 {olicy.`,
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
					Name:        "redirect_listener_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `See Argument Reference above. ## Import Load Balancer L7 Policy can be imported using the L7 Policy ID, e.g.: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_lb_l7policy_v2.l7policy_1 8a7a79c2-cf17-4e65-b2ae-ddc8bfcf6c74 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The unique ID for the L7 {olicy.`,
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
					Name:        "redirect_listener_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `See Argument Reference above. ## Import Load Balancer L7 Policy can be imported using the L7 Policy ID, e.g.: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_lb_l7policy_v2.l7policy_1 8a7a79c2-cf17-4e65-b2ae-ddc8bfcf6c74 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_lb_l7rule_v2",
			Category:         "Enhanced Load Balancer Resources",
			ShortDescription: `Manages a V2 l7rule resource within FlexibleEngine.`,
			Description:      ``,
			Keywords: []string{
				"enhanced",
				"load",
				"balancer",
				"lb",
				"l7rule",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 Networking client. A Networking client is needed to create an . If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new L7 Rule.`,
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
					Description: `The ID of the Listener owning this resource. ## Import Load Balancer L7 Rule can be imported using the L7 Policy ID and L7 Rule ID separated by a slash, e.g.: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_lb_l7rule_v2.l7rule_1 e0bd694a-abbe-450e-b329-0931fd1cc5eb/4086b0c9-b18c-4d1c-b6b8-4c56c3ad2a9e ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `The ID of the Listener owning this resource. ## Import Load Balancer L7 Rule can be imported using the L7 Policy ID and L7 Rule ID separated by a slash, e.g.: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_lb_l7rule_v2.l7rule_1 e0bd694a-abbe-450e-b329-0931fd1cc5eb/4086b0c9-b18c-4d1c-b6b8-4c56c3ad2a9e ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_lb_listener_v2",
			Category:         "Enhanced Load Balancer Resources",
			ShortDescription: `Manages a V2 listener resource within FlexibleEngine.`,
			Description:      ``,
			Keywords: []string{
				"enhanced",
				"load",
				"balancer",
				"lb",
				"listener",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 Networking client. A Networking client is needed to create an . If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new Listener.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Required) The protocol - can either be TCP, UDP, HTTP or TERMINATED_HTTPS. Changing this creates a new Listener.`,
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
					Name:        "tls_ciphers_policy",
					Description: `(Optional) Specifies the security policy used by the listener. This parameter is valid only when the load balancer protocol is set to TERMINATED_HTTPS. The value can be tls-1-0, tls-1-1, tls-1-2, or tls-1-2-strict, and the default value is tls-1-0. For details of cipher suites for each security policy, see the table below. <table> <tr> <th>Security Policy</th> <th>TLS Version</th> <th>Cipher Suite</th> </tr > <tr > <td>tls-1-0</td> <td>TLSv1.2 TLSv1.1 TLSv1</td> <td rowspan="3">ECDHE-RSA-AES256-GCM-SHA384:ECDHE-RSA-AES128-GCM-SHA256:ECDHE-ECDSA-AES256-GCM-SHA384:ECDHE-ECDSA-AES128-GCM-SHA256:AES128-GCM-SHA256:AES256-GCM-SHA384:ECDHE-ECDSA-AES128-SHA256:ECDHE-RSA-AES128-SHA256:AES128-SHA256:AES256-SHA256:ECDHE-ECDSA-AES256-SHA384:ECDHE-RSA-AES256-SHA384:ECDHE-ECDSA-AES128-SHA:ECDHE-RSA-AES128-SHA:ECDHE-RSA-AES256-SHA:ECDHE-ECDSA-AES256-SHA:AES128-SHA:AES256-SHA</td> </tr> <tr> <td>tls-1-1</td> <td>TLSv1.2 TLSv1.1</td> </tr> <tr> <td>tls-1-2</td> <td>TLSv1.2</td> </tr> <tr> <td >tls-1-2-strict</td> <td >TLSv1.2</td> <td >ECDHE-RSA-AES256-GCM-SHA384:ECDHE-RSA-AES128-GCM-SHA256:ECDHE-ECDSA-AES256-GCM-SHA384:ECDHE-ECDSA-AES128-GCM-SHA256:AES128-GCM-SHA256:AES256-GCM-SHA384:ECDHE-ECDSA-AES128-SHA256:ECDHE-RSA-AES128-SHA256:AES128-SHA256:AES256-SHA256:ECDHE-ECDSA-AES256-SHA384:ECDHE-RSA-AES256-SHA384</td> </tr> </table>`,
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
					Name:        "sni_container_refs",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tls_ciphers_policy",
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
					Name:        "sni_container_refs",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tls_ciphers_policy",
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
			Type:             "flexibleengine_lb_loadbalancer_v2",
			Category:         "Enhanced Load Balancer Resources",
			ShortDescription: `Manages a V2 loadbalancer resource within FlexibleEngine.`,
			Description:      ``,
			Icon:             "Network-ELB.svg",
			Keywords: []string{
				"enhanced",
				"load",
				"balancer",
				"lb",
				"loadbalancer",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 Networking client. A Networking client is needed to create an LB member. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new LB member.`,
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
					Name:        "flavor",
					Description: `(Optional) The UUID of a flavor. Currently, this is not supported. Changing this creates a new loadbalancer.`,
				},
				resource.Attribute{
					Name:        "loadbalancer_provider",
					Description: `(Optional) The name of the provider. Currently, only vlb is supported. Changing this creates a new loadbalancer.`,
				},
				resource.Attribute{
					Name:        "security_group_ids",
					Description: `(Optional) A list of security group IDs to apply to the loadbalancer. The security groups must be specified by ID and not name (as opposed to how they are configured with the Compute Instance). ## Attributes Reference The following attributes are exported:`,
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
					Name:        "flavor",
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
			Attributes: []resource.Attribute{
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
					Name:        "flavor",
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
			Type:             "flexibleengine_lb_member_v2",
			Category:         "Enhanced Load Balancer Resources",
			ShortDescription: `Manages a V2 member resource within FlexibleEngine.`,
			Description:      ``,
			Keywords: []string{
				"enhanced",
				"load",
				"balancer",
				"lb",
				"member",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 Networking client. A Networking client is needed to create an . If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new member.`,
				},
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
			Type:             "flexibleengine_lb_monitor_v2",
			Category:         "Enhanced Load Balancer Resources",
			ShortDescription: `Manages a V2 monitor resource within FlexibleEngine.`,
			Description:      ``,
			Keywords: []string{
				"enhanced",
				"load",
				"balancer",
				"lb",
				"monitor",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 Networking client. A Networking client is needed to create an . If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new monitor.`,
				},
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
					Description: `(Optional) The administrative state of the monitor. A valid value is true (UP) or false (DOWN).`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional) Specifies the health check port. The value ranges from 1 to 65536. ## Attributes Reference The following attributes are exported:`,
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
				resource.Attribute{
					Name:        "port",
					Description: `See Argument Reference above.`,
				},
			},
			Attributes: []resource.Attribute{
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
				resource.Attribute{
					Name:        "port",
					Description: `See Argument Reference above.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_lb_pool_v2",
			Category:         "Enhanced Load Balancer Resources",
			ShortDescription: `Manages a V2 pool resource within FlexibleEngine.`,
			Description:      ``,
			Keywords: []string{
				"enhanced",
				"load",
				"balancer",
				"lb",
				"pool",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 Networking client. A Networking client is needed to create an . If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new pool.`,
				},
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
			Attributes: []resource.Attribute{
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
			Type:             "flexibleengine_lb_whitelist_v2",
			Category:         "Enhanced Load Balancer Resources",
			ShortDescription: `Manages an Enhanced LB whitelist resource within FlexibleEngine.`,
			Description:      ``,
			Keywords: []string{
				"enhanced",
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
			Type:             "flexibleengine_mls_instance_v1",
			Category:         "MLS Resources",
			ShortDescription: `Manages mls instance resource within FlexibleEngine`,
			Description:      ``,
			Icon:             "Data Analysis&AI-MLS.svg",
			Keywords: []string{
				"mls",
				"instance",
				"v1",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to create the MLS instance. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new instance.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Specifies the MLS instance name. The DB instance name of the same type is unique in the same tenant. Changing this creates a new instance.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Required) Specifies MLS Software version, only ` + "`" + `1.2.0` + "`" + ` is supported now. Changing this creates a new instance.`,
				},
				resource.Attribute{
					Name:        "network",
					Description: `(Required) Specifies the instance network information. The structure is described below. Changing this creates a new instance.`,
				},
				resource.Attribute{
					Name:        "agency",
					Description: `(Optional) Specifies the agency name. This parameter is mandatory only when you bind an instance to an elastic IP address (EIP). An instance must be bound to an EIP to grant MLS rights to abtain a tenant's token. NOTE: The tenant must create an agency on the Identity and Access Management (IAM) interface in advance. Changing this creates a new instance.`,
				},
				resource.Attribute{
					Name:        "flavor",
					Description: `(Required) Specifies the instance flavor, only ` + "`" + `mls.c2.2xlarge.common` + "`" + ` is supported now. Changing this creates a new instance.`,
				},
				resource.Attribute{
					Name:        "mrs_cluster",
					Description: `(Required) Specifies the MRS cluster information which the instance is associated. The structure is described below. NOTE: The current MRS instance requires an MRS cluster whose version is 1.3.0 and that is configured with the Spark component. MRS clusters whose version is not 1.3.0 or that are not configured with the Spark component cannot be selected. Changing this creates a new instance. The ` + "`" + `network` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required) Specifies the ID of the virtual private cloud (VPC) where the instance resides. Changing this creates a new instance.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Required) Specifies the ID of the subnet where the instance resides. Changing this creates a new instance.`,
				},
				resource.Attribute{
					Name:        "security_group",
					Description: `(Optional) Specifies the ID of the security group of the instance. Changing this creates a new instance.`,
				},
				resource.Attribute{
					Name:        "available_zone",
					Description: `(Required) Specifies the AZ of the instance. Changing this creates a new instance.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `(Required) Specifies the IP address of the instance. The structure is described below. Changing this creates a new instance. The ` + "`" + `public_ip` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "bind_type",
					Description: `(Required) Specifies the bind type. Possible values: ` + "`" + `auto_assign` + "`" + ` and ` + "`" + `not_use` + "`" + `. Changing this creates a new instance. The ` + "`" + `mrs_cluster` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Required) Specifies the ID of the MRS cluster. Changing this creates a new instance.`,
				},
				resource.Attribute{
					Name:        "user_name",
					Description: `(Optional) Specifies the MRS cluster username. This parameter is mandatory only when the MRS cluster is in the security mode. Changing this creates a new instance.`,
				},
				resource.Attribute{
					Name:        "user_password",
					Description: `(Optional) Specifies the password of the MRS cluster user. The password and username work in a pair. Changing this creates a new instance. ## Attributes Reference The following attributes are exported:`,
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
					Name:        "version",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "agency",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "flavor",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "network/vpc_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "network/subnet_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "network/security_group",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "network/available_zone",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "network/public_ip/bind_type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "network/public_ip/eip_id",
					Description: `Indicates the EIP ID, This is returned only when bind_type is set to auto_assign.`,
				},
				resource.Attribute{
					Name:        "mrs_cluster",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Indicates the MLS instance status.`,
				},
				resource.Attribute{
					Name:        "inner_endpoint",
					Description: `Indicates the URL for accessing the instance. Only machines in the same VPC and subnet as the instance can access the URL.`,
				},
				resource.Attribute{
					Name:        "public_endpoint",
					Description: `Indicates the URL for accessing the instance. The URL can be accessed from the Internet. The URL is created only after the instance is bound to an EIP.`,
				},
				resource.Attribute{
					Name:        "created",
					Description: `Indicates the creation time in the following format: yyyy-mm-dd Thh:mm:ssZ.`,
				},
				resource.Attribute{
					Name:        "updated",
					Description: `Indicates the update time in the following format: yyyy-mm-dd Thh:mm:ssZ.`,
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
					Name:        "version",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "agency",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "flavor",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "network/vpc_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "network/subnet_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "network/security_group",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "network/available_zone",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "network/public_ip/bind_type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "network/public_ip/eip_id",
					Description: `Indicates the EIP ID, This is returned only when bind_type is set to auto_assign.`,
				},
				resource.Attribute{
					Name:        "mrs_cluster",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Indicates the MLS instance status.`,
				},
				resource.Attribute{
					Name:        "inner_endpoint",
					Description: `Indicates the URL for accessing the instance. Only machines in the same VPC and subnet as the instance can access the URL.`,
				},
				resource.Attribute{
					Name:        "public_endpoint",
					Description: `Indicates the URL for accessing the instance. The URL can be accessed from the Internet. The URL is created only after the instance is bound to an EIP.`,
				},
				resource.Attribute{
					Name:        "created",
					Description: `Indicates the creation time in the following format: yyyy-mm-dd Thh:mm:ssZ.`,
				},
				resource.Attribute{
					Name:        "updated",
					Description: `Indicates the update time in the following format: yyyy-mm-dd Thh:mm:ssZ.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_mrs_cluster_v1",
			Category:         "MRS Resources",
			ShortDescription: `Manages resource cluster within FlexibleEngine MRS.`,
			Description:      ``,
			Icon:             "Data Analysis&AI-MRS.svg",
			Keywords: []string{
				"mrs",
				"cluster",
				"v1",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "billing_type",
					Description: `(Required) The value is 12, indicating on-demand payment.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Required) Cluster region information. Obtain the value from Regions and Endpoints.`,
				},
				resource.Attribute{
					Name:        "master_node_num",
					Description: `(Required) Number of Master nodes The value is 2.`,
				},
				resource.Attribute{
					Name:        "master_node_size",
					Description: `(Required) Best match based on several years of commissioning experience. MRS supports specifications of hosts, and host specifications are determined by CPUs, memory, and disks space. - Master nodes support s1.4xlarge and s1.8xlarge, c3.2xlarge.2, c3.xlarge.4, c3.2xlarge.4, c3.4xlarge.2, c3.4xlarge.4, c3.8xlarge.4, c3.15xlarge.4. - Core nodes of a streaming cluster support s1.xlarge, c2.2xlarge, s1.2xlarge, s1.4xlarge, s1.8xlarge, d1.8xlarge, , c3.2xlarge.2, c3.xlarge.4, c3.2xlarge.4, c3.4xlarge.2, c3.4xlarge.4, c3.8xlarge.4, c3.15xlarge.4. - Core nodes of an analysis cluster support all specifications c2.2xlarge, s1.xlarge, s1.4xlarge, s1.8xlarge, d1.xlarge, d1.2xlarge, d1.4xlarge, d1.8xlarge, , c3.2xlarge.2, c3.xlarge.4, c3.2xlarge.4, c3.4xlarge.2, c3.4xlarge.4, c3.8xlarge.4, c3.15xlarge.4, d2.xlarge.8, d2.2xlarge.8, d2.4xlarge.8, d2.8xlarge.8. The following provides specification details. node_size | CPU(core) | Memory(GB) | System Disk | Data Disk | --- | --- | --- | --- | --- | c2.2xlarge.linux.mrs | 8 | 16 | 40 | - cc3.xlarge.4.linux.mrs | 4 | 16 | 40 | - cc3.2xlarge.4.linux.mrs | 8 | 32 | 40 | - cc3.4xlarge.4.linux.mrs | 16 | 64 | 40 | - cc3.8xlarge.4.linux.mrs | 32 | 128 | 40 | - s1.xlarge.linux.mrs | 4 | 16 | 40 | - s1.4xlarge.linux.mrs | 16 | 64 | 40 | - s1.8xlarge.linux.mrs | 32 | 128 | 40 | - s3.xlarge.4.linux.mrs| 4 | 16 | 40 | - s3.2xlarge.4.linux.mrs| 8 | 32 | 40 | - s3.4xlarge.4.linux.mrs| 16 | 64 | 40 | - d1.xlarge.linux.mrs | 6 | 55 | 40 | 1.8 TB x 3 HDDs d1.2xlarge.linux.mrs | 12 | 110 | 40 | 1.8 TB x 6 HDDs d1.4xlarge.linux.mrs | 24 | 220 | 40 | 1.8 TB x 12 HDDs d1.8xlarge.linux.mrs | 48 | 440 | 40 | 1.8 TB x 24 HDDs d2.xlarge.linux.mrs | 4 | 32 | 40 | - d2.2xlarge.linux.mrs | 8 | 64 | 40 | - d2.4xlarge.linux.mrs | 16 | 128 | 40 | 1.8TB`,
				},
				resource.Attribute{
					Name:        "core_node_num",
					Description: `(Required) Number of Core nodes Value range: 3 to 500. A maximum of 500 Core nodes are supported by default. If more than 500 Core nodes are required, contact technical support engineers or invoke background APIs to modify the database.`,
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
					Description: `(Optional) Version of the clusters. Currently, MRS 1.3.0, MRS 1.5.0, MRS 1.6.3, MRS 1.8.9, and MRS 2.0.1 are supported. The latest version of MRS is used by default. Currently, the latest version is MRS 2.0.1.`,
				},
				resource.Attribute{
					Name:        "cluster_type",
					Description: `(Optional) Type of clusters. 0: analysis cluster; 1: streaming cluster. The default value is 0.`,
				},
				resource.Attribute{
					Name:        "volume_type",
					Description: `(Required) Type of disks SATA and SSD are supported. SATA: common I/O SSD: super high-speed I/O`,
				},
				resource.Attribute{
					Name:        "volume_size",
					Description: `(Required) Data disk storage space of a Core node Users can add disks to expand storage capacity when creating a cluster. There are the following scenarios: Separation of data storage and computing: Data is stored in the OBS system. Costs of clusters are relatively low but computing performance is poor. The clusters can be deleted at any time. It is recommended when data computing is not frequently performed. Integration of data storage and computing: Data is stored in the HDFS system. Costs of clusters are relatively high but computing performance is good. The clusters cannot be deleted in a short term. It is recommended when data computing is frequently performed. Value range: 100 GB to 32000 GB`,
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
					Description: `(Required) Component name Currently, Hadoop, Spark, HBase, Hive, Hue, Loader, Flume, Kafka and Storm are supported. Loader and Flume are not supported by MRS 1.3.0.`,
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
			Attributes: []resource.Attribute{
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
			Type:             "flexibleengine_mrs_hybrid_cluster_v1",
			Category:         "MRS Resources",
			ShortDescription: `Manages resource cluster within FlexibleEngine MRS.`,
			Description:      ``,
			Keywords: []string{
				"mrs",
				"hybrid",
				"cluster",
				"v1",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) Cluster region information. Obtain the value from Regions and Endpoints.`,
				},
				resource.Attribute{
					Name:        "available_zone",
					Description: `(Required) ID or Name of an available zone. Obtain the value from Regions and Endpoints.`,
				},
				resource.Attribute{
					Name:        "cluster_name",
					Description: `(Required) Cluster name, which is globally unique and contains only 1 to 64 letters, digits, hyphens (-), and underscores (_).`,
				},
				resource.Attribute{
					Name:        "cluster_version",
					Description: `(Required) Version of the clusters. Currently, MRS 1.6.3, MRS 1.8.9, and MRS 2.0.1 are supported. The latest version of MRS is used by default. Currently, the latest version is MRS 2.0.1.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required) Specifies the id of the VPC.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Required) Specifies the id of the subnet.`,
				},
				resource.Attribute{
					Name:        "safe_mode",
					Description: `(Optional) MRS cluster running mode - 0: common mode The value indicates that the Kerberos authentication is disabled. Users can use all functions provided by the cluster. - 1: safe mode (by default) The value indicates that the Kerberos authentication is enabled. Common users cannot use the file management or job management functions of an MRS cluster and cannot view cluster resource usage or the job records of Hadoop and Spark. To use these functions, the users must obtain the relevant permissions from the MRS Manager administrator. The request has the cluster_admin_secret parameter only when safe_mode is set to 1.`,
				},
				resource.Attribute{
					Name:        "cluster_admin_secret",
					Description: `(Required) Indicates the password of the MRS Manager administrator. - Must contain 8 to 32 characters. - Must contain at least three types of the following: Lowercase letters, Uppercase letters, Digits, Special characters of ` + "`" + `~!@#$%^&`,
				},
				resource.Attribute{
					Name:        "master_node_key_pair",
					Description: `(Required) Name of a key pair You can use a key to log in to the Master node in the cluster.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `(Optional) Specifies the id of the security group which the cluster belongs to. If this parameter is empty, MRS automatically creates a security group, whose name starts with mrs_{cluster_name}.`,
				},
				resource.Attribute{
					Name:        "log_collection",
					Description: `(Optional) Indicates whether logs are collected when cluster installation fails. 0: not collected 1: collected The default value is 0. If log_collection is set to 1, OBS buckets will be created to collect the MRS logs. These buckets will be charged.`,
				},
				resource.Attribute{
					Name:        "component_list",
					Description: `(Required) Component name - Presto, Hadoop, Spark, HBase, Hive, Tez, Hue, Loader, Flume, Kafka and Storm are supported by MRS 2.0.1 or later. - Presto, Hadoop, Spark, HBase, Opentsdb, Hive, Hue, Loader, Flink, Flume, Kafka, KafkaManager and Storm are supported by MRS 1.8.9. - Hadoop, Spark, HBase, Hive, Hue, Loader, Flume, Kafka and Storm are supported by versions earlier than MRS 1.8.9.`,
				},
				resource.Attribute{
					Name:        "master_nodes",
					Description: `(Required) Specifies the master nodes information.`,
				},
				resource.Attribute{
					Name:        "analysis_core_nodes",
					Description: `(Required) Specifies the analysis core nodes information.`,
				},
				resource.Attribute{
					Name:        "streaming_core_nodes",
					Description: `(Required) Specifies the streaming core nodes information.`,
				},
				resource.Attribute{
					Name:        "analysis_task_nodes",
					Description: `(Optional) Specifies the analysis task nodes information.`,
				},
				resource.Attribute{
					Name:        "streaming_task_nodes",
					Description: `(Optional) Specifies the streaming task nodes information. The ` + "`" + `master_nodes` + "`" + `, ` + "`" + `analysis_core_nodes` + "`" + `, ` + "`" + `streaming_core_nodes` + "`" + `, ` + "`" + `analysis_task_nodes` + "`" + `, ` + "`" + `streaming_task_nodes` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "flavor",
					Description: `(Required) Best match based on several years of commissioning experience. MRS supports specifications of hosts, and host specifications are determined by CPUs, memory, and disks space. - Master nodes support s1.4xlarge and s1.8xlarge, c3.2xlarge.2, c3.xlarge.4, c3.2xlarge.4, c3.4xlarge.2, c3.4xlarge.4, c3.8xlarge.4, c3.15xlarge.4. - Core nodes of a streaming cluster support s1.xlarge, c2.2xlarge, s1.2xlarge, s1.4xlarge, s1.8xlarge, d1.8xlarge, , c3.2xlarge.2, c3.xlarge.4, c3.2xlarge.4, c3.4xlarge.2, c3.4xlarge.4, c3.8xlarge.4, c3.15xlarge.4. - Core nodes of an analysis cluster support all specifications c2.2xlarge, s1.xlarge, s1.4xlarge, s1.8xlarge, d1.xlarge, d1.2xlarge, d1.4xlarge, d1.8xlarge, , c3.2xlarge.2, c3.xlarge.4, c3.2xlarge.4, c3.4xlarge.2, c3.4xlarge.4, c3.8xlarge.4, c3.15xlarge.4, d2.xlarge.8, d2.2xlarge.8, d2.4xlarge.8, d2.8xlarge.8. The following provides specification details. node_size | CPU(core) | Memory(GB) | System Disk | Data Disk | --- | --- | --- | --- | --- | c2.2xlarge.linux.mrs | 8 | 16 | 40 | - cc3.xlarge.4.linux.mrs | 4 | 16 | 40 | - cc3.2xlarge.4.linux.mrs | 8 | 32 | 40 | - cc3.4xlarge.4.linux.mrs | 16 | 64 | 40 | - cc3.8xlarge.4.linux.mrs | 32 | 128 | 40 | - s1.xlarge.linux.mrs | 4 | 16 | 40 | - s1.4xlarge.linux.mrs | 16 | 64 | 40 | - s1.8xlarge.linux.mrs | 32 | 128 | 40 | - s3.xlarge.4.linux.mrs| 4 | 16 | 40 | - s3.2xlarge.4.linux.mrs| 8 | 32 | 40 | - s3.4xlarge.4.linux.mrs| 16 | 64 | 40 | - d1.xlarge.linux.mrs | 6 | 55 | 40 | 1.8 TB x 3 HDDs d1.2xlarge.linux.mrs | 12 | 110 | 40 | 1.8 TB x 6 HDDs d1.4xlarge.linux.mrs | 24 | 220 | 40 | 1.8 TB x 12 HDDs d1.8xlarge.linux.mrs | 48 | 440 | 40 | 1.8 TB x 24 HDDs d2.xlarge.linux.mrs | 4 | 32 | 40 | - d2.2xlarge.linux.mrs | 8 | 64 | 40 | - d2.4xlarge.linux.mrs | 16 | 128 | 40 | 1.8TB`,
				},
				resource.Attribute{
					Name:        "node_number",
					Description: `(Required) Number of nodes. The value ranges from 0 to 500 and the default value is 0. The total number of Core and Task nodes cannot exceed 500.`,
				},
				resource.Attribute{
					Name:        "data_volume_type",
					Description: `(Required) Data disk storage type of the node, supporting SATA and SSD currently - SATA: common I/O - SSD: Ultrahigh-speed I/O`,
				},
				resource.Attribute{
					Name:        "data_volume_size",
					Description: `(Required) Data disk size of the node Value range: 100 GB to 32000 GB`,
				},
				resource.Attribute{
					Name:        "data_volume_count",
					Description: `(Required) Number of data disks of the node Value range: 0 to 10 ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "available_zone",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "cluster_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "cluster_version",
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
					Name:        "master_node_key_pair",
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
					Name:        "security_group_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "log_collection",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "master_nodes",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "analysis_core_nodes",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "streaming_core_nodes",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "analysis_task_nodes",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "streaming_task_nodes",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "component_list",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "billing_type",
					Description: `The value is Metered, indicating on-demand payment.`,
				},
				resource.Attribute{
					Name:        "total_node_number",
					Description: `Total node number.`,
				},
				resource.Attribute{
					Name:        "master_node_ip",
					Description: `IP address of a Master node.`,
				},
				resource.Attribute{
					Name:        "internal_ip",
					Description: `Iternal IP address.`,
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
					Name:        "external_alternate_ip",
					Description: `Backup external IP address.`,
				},
				resource.Attribute{
					Name:        "vnc",
					Description: `URI address for remote login of the elastic cloud server.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `Cluster creation fee, which is automatically calculated.`,
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
					Name:        "charging_start_time",
					Description: `Time when charging starts. The components attributes:`,
				},
				resource.Attribute{
					Name:        "component_name",
					Description: `Component name`,
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
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "available_zone",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "cluster_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "cluster_version",
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
					Name:        "master_node_key_pair",
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
					Name:        "security_group_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "log_collection",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "master_nodes",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "analysis_core_nodes",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "streaming_core_nodes",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "analysis_task_nodes",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "streaming_task_nodes",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "component_list",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "billing_type",
					Description: `The value is Metered, indicating on-demand payment.`,
				},
				resource.Attribute{
					Name:        "total_node_number",
					Description: `Total node number.`,
				},
				resource.Attribute{
					Name:        "master_node_ip",
					Description: `IP address of a Master node.`,
				},
				resource.Attribute{
					Name:        "internal_ip",
					Description: `Iternal IP address.`,
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
					Name:        "external_alternate_ip",
					Description: `Backup external IP address.`,
				},
				resource.Attribute{
					Name:        "vnc",
					Description: `URI address for remote login of the elastic cloud server.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `Cluster creation fee, which is automatically calculated.`,
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
					Name:        "charging_start_time",
					Description: `Time when charging starts. The components attributes:`,
				},
				resource.Attribute{
					Name:        "component_name",
					Description: `Component name`,
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
			Type:             "flexibleengine_mrs_job_v1",
			Category:         "MRS Resources",
			ShortDescription: `Manages resource job within FlexibleEngine MRS.`,
			Description:      ``,
			Icon:             "Data Analysis&AI-MRS.svg",
			Keywords: []string{
				"mrs",
				"job",
				"v1",
			},
			Arguments: []resource.Attribute{
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
			Attributes: []resource.Attribute{
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
			Type:             "flexibleengine_nat_dnat_rule_v2",
			Category:         "NAT Resources",
			ShortDescription: `Manages a V2 dnat rule resource within FlexibleEngine Nat.`,
			Description:      ``,
			Icon:             "nat-gateway.svg",
			Keywords: []string{
				"nat",
				"dnat",
				"rule",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "floating_ip_id",
					Description: `(Required) Specifies the ID of the floating IP address. Changing this creates a new resource.`,
				},
				resource.Attribute{
					Name:        "internal_service_port",
					Description: `(Required) Specifies port used by ECSs or BMSs to provide services for external systems. Changing this creates a new resource.`,
				},
				resource.Attribute{
					Name:        "nat_gateway_id",
					Description: `(Required) ID of the nat gateway this dnat rule belongs to. Changing this creates a new dnat rule.`,
				},
				resource.Attribute{
					Name:        "port_id",
					Description: `(Optional) Specifies the port ID of an ECS or a BMS. This parameter and private_ip are alternative. Changing this creates a new dnat rule.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `(Optional) Specifies the private IP address of a user, for example, the IP address of a VPC for dedicated connection. This parameter and port_id are alternative. Changing this creates a new dnat rule.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Required) Specifies the protocol type. Currently, TCP, UDP, and ANY are supported. The protocol number of TCP, UDP, and ANY is 6, 17, and 0, respectively. Changing this creates a new dnat rule.`,
				},
				resource.Attribute{
					Name:        "external_service_port",
					Description: `(Required) Specifies port used by ECSs or BMSs to provide services for external systems. Changing this creates a new dnat rule. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Dnat rule creation time.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Dnat rule status. ## Import Dnat can be imported using the following format: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_nat_dnat_rule_v2.dnat_1 f4f783a7-b908-4215-b018-724960e5df4a ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "created_at",
					Description: `Dnat rule creation time.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Dnat rule status. ## Import Dnat can be imported using the following format: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_nat_dnat_rule_v2.dnat_1 f4f783a7-b908-4215-b018-724960e5df4a ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_nat_gateway_v2",
			Category:         "NAT Resources",
			ShortDescription: `Manages a V2 nat gateway resource within FlexibleEngine Nat.`,
			Description:      ``,
			Icon:             "nat-gateway.svg",
			Keywords: []string{
				"nat",
				"gateway",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 nat client. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new nat gateway.`,
				},
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
					Description: `(Required) The specification of the nat gateway, valid values are "1", "2", "3", "4" (for Small, Medium, Large, Extra-Large)`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Optional) The target tenant/project ID in which to allocate the nat gateway. Changing this creates a new nat gateway .`,
				},
				resource.Attribute{
					Name:        "router_id",
					Description: `(Required) ID of the router/VPC this nat gateway belongs to. Changing this creates a new nat gateway.`,
				},
				resource.Attribute{
					Name:        "internal_network_id",
					Description: `(Required) ID of the subnet (!) this nat gateway connects to. Changing this creates a new nat gateway. ## Attributes Reference The following attributes are exported:`,
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
			Type:             "flexibleengine_nat_snat_rule_v2",
			Category:         "NAT Resources",
			ShortDescription: `Manages a V2 snat rule resource within FlexibleEngine Nat.`,
			Description:      ``,
			Icon:             "nat-gateway.svg",
			Keywords: []string{
				"nat",
				"snat",
				"rule",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 nat client. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new snat rule.`,
				},
				resource.Attribute{
					Name:        "nat_gateway_id",
					Description: `(Required) ID of the nat gateway this snat rule belongs to. Changing this creates a new snat rule.`,
				},
				resource.Attribute{
					Name:        "floating_ip_id",
					Description: `(Required) ID of the floating ip this snat rule connets to. Changing this creates a new snat rule.`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `(Optional) ID of the network this snat rule connects to. This parameter and ` + "`" + `cidr` + "`" + ` are alternative. Changing this creates a new snat rule.`,
				},
				resource.Attribute{
					Name:        "cidr",
					Description: `(Optional) Specifies CIDR, which can be in the format of a network segment or a host IP address. This parameter and ` + "`" + `network_id` + "`" + ` are alternative. Changing this creates a new snat rule.`,
				},
				resource.Attribute{
					Name:        "source_type",
					Description: `(Optional) Specifies the scenario. The valid value is 0 (VPC scenario) and 1 (Direct Connect scenario). Only ` + "`" + `cidr` + "`" + ` can be specified over a Direct Connect connection. If no value is entered, the default value 0 (VPC scenario) is used. Changing this creates a new snat rule. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
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
					Name:        "cidr",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "source_type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "floating_ip_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "floating_ip_address",
					Description: `The actual floating IP address.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the SNAT rule.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
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
					Name:        "cidr",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "source_type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "floating_ip_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "floating_ip_address",
					Description: `The actual floating IP address.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the SNAT rule.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_networking_floatingip_associate_v2",
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
					Description: `See Argument Reference above. ## Import Floating IP associations can be imported using the ` + "`" + `id` + "`" + ` of the floating IP, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_networking_floatingip_associate_v2.fip 2c7f39f3-702b-48d1-940c-b50384177ee1 ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `See Argument Reference above. ## Import Floating IP associations can be imported using the ` + "`" + `id` + "`" + ` of the floating IP, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_networking_floatingip_associate_v2.fip 2c7f39f3-702b-48d1-940c-b50384177ee1 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_networking_floatingip_v2",
			Category:         "Networking Resources",
			ShortDescription: `Manages a V2 floating IP resource within FlexibleEngine Neutron (networking).`,
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
					Description: `(Optional) The name of the pool from which to obtain the floating IP. Default value is admin_external_net. Changing this creates a new floating IP.`,
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
					Description: `The fixed IP which the floating IP maps to. ## Import Floating IPs can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_networking_floatingip_v2.floatip_1 2c7f39f3-702b-48d1-940c-b50384177ee1 ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `The fixed IP which the floating IP maps to. ## Import Floating IPs can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_networking_floatingip_v2.floatip_1 2c7f39f3-702b-48d1-940c-b50384177ee1 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_networking_network_v2",
			Category:         "Networking Resources",
			ShortDescription: `Manages a V2 Neutron network resource within FlexibleEngine.`,
			Description:      ``,
			Icon:             "Network.svg",
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
					Description: `The physical network where this network is implemented.`,
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
					Description: `See Argument Reference above. ## Import Networks can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_networking_network_v2.network_1 d90ce693-5ccf-4136-a0ed-152ce412b6b9 ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `See Argument Reference above. ## Import Networks can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_networking_network_v2.network_1 d90ce693-5ccf-4136-a0ed-152ce412b6b9 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_networking_port_v2",
			Category:         "Networking Resources",
			ShortDescription: `Manages a V2 port resource within FlexibleEngine.`,
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
					Name:        "all fixed_ips",
					Description: `The collection of Fixed IP addresses on the port in the order returned by the Network v2 API. ## Import Ports can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_networking_port_v2.port_1 eae26a3e-1c33-4cc1-9c31-0cd729c438a1 ` + "`" + `` + "`" + `` + "`" + ` ## Notes ### Ports and Instances There are some notes to consider when connecting Instances to networks using Ports. Please see the ` + "`" + `flexibleengine_compute_instance_v2` + "`" + ` documentation for further documentation.`,
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
					Name:        "all fixed_ips",
					Description: `The collection of Fixed IP addresses on the port in the order returned by the Network v2 API. ## Import Ports can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_networking_port_v2.port_1 eae26a3e-1c33-4cc1-9c31-0cd729c438a1 ` + "`" + `` + "`" + `` + "`" + ` ## Notes ### Ports and Instances There are some notes to consider when connecting Instances to networks using Ports. Please see the ` + "`" + `flexibleengine_compute_instance_v2` + "`" + ` documentation for further documentation.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_networking_router_interface_v2",
			Category:         "Networking Resources",
			ShortDescription: `Manages a V2 router interface resource within FlexibleEngine.`,
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
					Description: `See Argument Reference above.`,
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
					Description: `See Argument Reference above.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_networking_router_route_v2",
			Category:         "Networking Resources",
			ShortDescription: `Creates a routing entry on a FlexibleEngine V2 router.`,
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
					Description: `See Argument Reference above. ## Notes The ` + "`" + `next_hop` + "`" + ` IP address must be directly reachable from the router at the ` + "`" + `` + "`" + `flexibleengine_networking_router_route_v2` + "`" + `` + "`" + ` resource creation time. You can ensure that by explicitly specifying a dependency on the ` + "`" + `` + "`" + `flexibleengine_networking_router_interface_v2` + "`" + `` + "`" + ` resource that connects the next hop to the router, as in the example above.`,
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
					Description: `See Argument Reference above. ## Notes The ` + "`" + `next_hop` + "`" + ` IP address must be directly reachable from the router at the ` + "`" + `` + "`" + `flexibleengine_networking_router_route_v2` + "`" + `` + "`" + ` resource creation time. You can ensure that by explicitly specifying a dependency on the ` + "`" + `` + "`" + `flexibleengine_networking_router_interface_v2` + "`" + `` + "`" + ` resource that connects the next hop to the router, as in the example above.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_networking_router_v2",
			Category:         "Networking Resources",
			ShortDescription: `Manages a V2 router resource within FlexibleEngine. The router is the top-level resource for the VPC within FlexibleEngine.`,
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
			Type:             "flexibleengine_networking_secgroup_rule_v2",
			Category:         "Networking Resources",
			ShortDescription: `Manages a V2 Neutron security group rule resource within FlexibleEngine.`,
			Description:      ``,
			Icon:             "security-group.svg",
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
					Description: `(Optional) The remote group id, the value needs to be an FlexibleEngine ID of a security group in the same tenant. Changing this creates a new security group rule.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `(Required) The security group id the rule should belong to, the value needs to be an FlexibleEngine ID of a security group in the same tenant. Changing this creates a new security group rule.`,
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
					Description: `See Argument Reference above. ## Import Security Group Rules can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_networking_secgroup_rule_v2.secgroup_rule_1 aeb68ee3-6e9d-4256-955c-9584a6212745 ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `See Argument Reference above. ## Import Security Group Rules can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_networking_secgroup_rule_v2.secgroup_rule_1 aeb68ee3-6e9d-4256-955c-9584a6212745 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_networking_secgroup_v2",
			Category:         "Networking Resources",
			ShortDescription: `Manages a V2 Neutron security group resource within FlexibleEngine.`,
			Description:      ``,
			Icon:             "security-group.svg",
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
					Description: `See Argument Reference above. ## Default Security Group Rules In most cases, FlexibleEngine will create some egress security group rules for each new security group. These security group rules will not be managed by Terraform, so if you prefer to have`,
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
					Description: `See Argument Reference above. ## Default Security Group Rules In most cases, FlexibleEngine will create some egress security group rules for each new security group. These security group rules will not be managed by Terraform, so if you prefer to have`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_networking_subnet_v2",
			Category:         "Networking Resources",
			ShortDescription: `Manages a V2 Neutron subnet resource within FlexibleEngine.`,
			Description:      ``,
			Icon:             "subnet.svg",
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
					Description: `(Required) The ending addresss. The ` + "`" + `host_routes` + "`" + ` block supports:`,
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
					Description: `See Argument Reference above. ## Import Subnets can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_networking_subnet_v2.subnet_1 da4faf16-5546-41e4-8330-4d0002b74048 ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `See Argument Reference above. ## Import Subnets can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_networking_subnet_v2.subnet_1 da4faf16-5546-41e4-8330-4d0002b74048 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_networking_vip_associate_v2",
			Category:         "Networking Resources",
			ShortDescription: `Manages a V2 vip associate resource within FlexibleEngine.`,
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
			Type:             "flexibleengine_networking_vip_v2",
			Category:         "Networking Resources",
			ShortDescription: `Manages a V2 vip resource within FlexibleEngine.`,
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
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_obs_bucket",
			Category:         "OBS Resource",
			ShortDescription: `Provides an OBS bucket resource within FlexibleEngine.`,
			Description:      ``,
			Keywords: []string{
				"obs",
				"resource",
				"bucket",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "bucket",
					Description: `(Required) Specifies the name of the bucket. Changing this parameter will create a new resource. A bucket must be named according to the globally applied DNS naming regulations as follows:`,
				},
				resource.Attribute{
					Name:        "storage_class",
					Description: `(Optional) Specifies the storage class of the bucket. OBS provides three storage classes: "STANDARD", "STANDARD_IA" (Infrequent Access) and "GLACIER" (Archive). Defaults to ` + "`" + `STANDARD` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "acl",
					Description: `(Optional) Specifies the ACL policy for a bucket. The predefined common policies are as follows: "private", "public-read", "public-read-write" and "log-delivery-write". Defaults to ` + "`" + `private` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "versioning",
					Description: `(Optional) Whether enable versioning. Once you version-enable a bucket, it can never return to an unversioned state. You can, however, suspend versioning on that bucket.`,
				},
				resource.Attribute{
					Name:        "logging",
					Description: `(Optional) A settings of bucket logging (documented below).`,
				},
				resource.Attribute{
					Name:        "website",
					Description: `(Optional) A website object (documented below).`,
				},
				resource.Attribute{
					Name:        "cors_rule",
					Description: `(Optional) A rule of Cross-Origin Resource Sharing (documented below).`,
				},
				resource.Attribute{
					Name:        "lifecycle_rule",
					Description: `(Optional) A configuration of object lifecycle management (documented below).`,
				},
				resource.Attribute{
					Name:        "force_destroy",
					Description: `(Optional) A boolean that indicates all objects should be deleted from the bucket so that the bucket can be destroyed without error. Default to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) If specified, the region this bucket should reside in. Otherwise, the region used by the provider. The ` + "`" + `logging` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "target_bucket",
					Description: `(Required) The name of the bucket that will receive the log objects. The acl policy of the target bucket should be ` + "`" + `log-delivery-write` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "target_prefix",
					Description: `(Optional) To specify a key prefix for log objects. The ` + "`" + `website` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "index_document",
					Description: `(Required, unless using ` + "`" + `redirect_all_requests_to` + "`" + `) Specifies the default homepage of the static website, only HTML web pages are supported. OBS only allows files such as ` + "`" + `index.html` + "`" + ` in the root directory of a bucket to function as the default homepage. That is to say, do not set the default homepage with a multi-level directory structure (for example, /page/index.html).`,
				},
				resource.Attribute{
					Name:        "error_document",
					Description: `(Optional) Specifies the error page returned when an error occurs during static website access. Only HTML, JPG, PNG, BMP, and WEBP files under the root directory are supported.`,
				},
				resource.Attribute{
					Name:        "redirect_all_requests_to",
					Description: `(Optional) A hostname to redirect all website requests for this bucket to. Hostname can optionally be prefixed with a protocol (` + "`" + `http://` + "`" + ` or ` + "`" + `https://` + "`" + `) to use when redirecting requests. The default is the protocol that is used in the original request.`,
				},
				resource.Attribute{
					Name:        "routing_rules",
					Description: `(Optional) A JSON or XML format containing routing rules describing redirect behavior and when redirects are applied. Each rule contains a ` + "`" + `Condition` + "`" + ` and a ` + "`" + `Redirect` + "`" + ` as shown in the following table: Parameter | Key -|- Condition | KeyPrefixEquals, HttpErrorCodeReturnedEquals Redirect | Protocol, HostName, ReplaceKeyPrefixWith, ReplaceKeyWith, HttpRedirectCode The ` + "`" + `cors_rule` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Unique identifier for lifecycle rules. The Rule Name contains a maximum of 255 characters.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Required) Specifies lifecycle rule status.`,
				},
				resource.Attribute{
					Name:        "prefix",
					Description: `(Optional) Object key prefix identifying one or more objects to which the rule applies. If omitted, all objects in the bucket will be managed by the lifecycle rule. The prefix cannot start or end with a slash (/), cannot have consecutive slashes (/), and cannot contain the following special characters: \:`,
				},
				resource.Attribute{
					Name:        "expiration",
					Description: `(Optional) Specifies a period when objects that have been last updated are automatically deleted. (documented below).`,
				},
				resource.Attribute{
					Name:        "transition",
					Description: `(Optional) Specifies a period when objects that have been last updated are automatically transitioned to ` + "`" + `STANDARD_IA` + "`" + ` or ` + "`" + `GLACIER` + "`" + ` storage class (documented below).`,
				},
				resource.Attribute{
					Name:        "noncurrent_version_expiration",
					Description: `(Optional) Specifies a period when noncurrent object versions are automatically deleted. (documented below).`,
				},
				resource.Attribute{
					Name:        "noncurrent_version_transition",
					Description: `(Optional) Specifies a period when noncurrent object versions are automatically transitioned to ` + "`" + `STANDARD_IA` + "`" + ` or ` + "`" + `GLACIER` + "`" + ` storage class (documented below). At least one of ` + "`" + `expiration` + "`" + `, ` + "`" + `transition` + "`" + `, ` + "`" + `noncurrent_version_expiration` + "`" + `, ` + "`" + `noncurrent_version_transition` + "`" + ` must be specified. The ` + "`" + `expiration` + "`" + ` object supports the following`,
				},
				resource.Attribute{
					Name:        "storage_class",
					Description: `(Required) The class of storage used to store the object. Only "STANDARD_IA" and "GLACIER" are supported. The ` + "`" + `noncurrent_version_expiration` + "`" + ` object supports the following`,
				},
				resource.Attribute{
					Name:        "storage_class",
					Description: `(Required) The class of storage used to store the object. Only "STANDARD_IA" and "GLACIER" are supported. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The name of the bucket.`,
				},
				resource.Attribute{
					Name:        "bucket_domain_name",
					Description: `The bucket domain name. Will be of format ` + "`" + `bucketname.oss.region.prod-cloud-ocb.orange-business.com` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The region this bucket resides in. ## Import OBS bucket can be imported using the ` + "`" + `bucket` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_obs_bucket.bucket bucket-name ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The name of the bucket.`,
				},
				resource.Attribute{
					Name:        "bucket_domain_name",
					Description: `The bucket domain name. Will be of format ` + "`" + `bucketname.oss.region.prod-cloud-ocb.orange-business.com` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The region this bucket resides in. ## Import OBS bucket can be imported using the ` + "`" + `bucket` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_obs_bucket.bucket bucket-name ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_obs_bucket_object",
			Category:         "OBS Resource",
			ShortDescription: `Provides an OBS bucket object resource within FlexibleEngine.`,
			Description:      ``,
			Keywords: []string{
				"obs",
				"resource",
				"bucket",
				"object",
			},
			Arguments: []resource.Attribute{
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
					Description: `(Optional) The path to the source file being uploaded to the bucket.`,
				},
				resource.Attribute{
					Name:        "content",
					Description: `(Optional) The literal content being uploaded to the bucket.`,
				},
				resource.Attribute{
					Name:        "acl",
					Description: `(Optional) The ACL policy to apply. Defaults to ` + "`" + `private` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "storage_class",
					Description: `(Optioanl) Specifies the storage class of the object. Defaults to ` + "`" + `STANDARD` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "content_type",
					Description: `(Optional) A standard MIME type describing the format of the object data, e.g. application/octet-stream. All Valid MIME Types are valid for this input.`,
				},
				resource.Attribute{
					Name:        "encryption",
					Description: `(Optional) Whether enable server-side encryption of the object in SSE-KMS mode.`,
				},
				resource.Attribute{
					Name:        "sse_kms_key_id",
					Description: `(Optional) The ID of the kms key. If omitted, the default master key will be used.`,
				},
				resource.Attribute{
					Name:        "etag",
					Description: `(Optional) Specifies the unique identifier of the object content. It can be used to trigger updates. The only meaningful value is ` + "`" + `md5(file("path_to_file"))` + "`" + `. Either ` + "`" + `source` + "`" + ` or ` + "`" + `content` + "`" + ` must be provided to specify the bucket content. These two arguments are mutually-exclusive. ## Attributes Reference The following attributes are exported`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `the ` + "`" + `key` + "`" + ` of the resource supplied above.`,
				},
				resource.Attribute{
					Name:        "etag",
					Description: `the ETag generated for the object (an MD5 sum of the object content). When the object is encrypted on the server side, the ETag value is not the MD5 value of the object, but the unique identifier calculated through the server-side encryption.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `the size of the object in bytes.`,
				},
				resource.Attribute{
					Name:        "version_id",
					Description: `A unique version ID value for the object, if bucket versioning is enabled.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `the ` + "`" + `key` + "`" + ` of the resource supplied above.`,
				},
				resource.Attribute{
					Name:        "etag",
					Description: `the ETag generated for the object (an MD5 sum of the object content). When the object is encrypted on the server side, the ETag value is not the MD5 value of the object, but the unique identifier calculated through the server-side encryption.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `the size of the object in bytes.`,
				},
				resource.Attribute{
					Name:        "version_id",
					Description: `A unique version ID value for the object, if bucket versioning is enabled.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_rds_instance_v1",
			Category:         "RDS Resources",
			ShortDescription: `Manages rds instance resource within FlexibleEngine`,
			Description:      ``,
			Icon:             "Database-RDS .svg",
			Keywords: []string{
				"rds",
				"instance",
				"v1",
			},
			Arguments: []resource.Attribute{
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
					Name:        "region",
					Description: `(Required) Specifies the region ID.`,
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
					Name:        "region",
					Description: `See Argument Reference above.`,
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
			Type:             "flexibleengine_rds_instance_v3",
			Category:         "RDS Resources",
			ShortDescription: `instance management`,
			Description:      ``,
			Icon:             "Database-RDS .svg",
			Keywords: []string{
				"rds",
				"instance",
				"v3",
			},
			Arguments: []resource.Attribute{
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
					Description: `(Required) Specifies the VPC ID. Changing this parameter will create a new resource.`,
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
					Description: `(Optional) Specifies the parameter group ID. Changing this parameter will create a new resource. The ` + "`" + `db` + "`" + ` block supports:`,
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
					Name:        "version",
					Description: `(Required) Specifies the database version. MySQL databases support MySQL 5.6 and 5.7, example values: "5.6", "5.7". PostgreSQL databases support PostgreSQL 9.5, 9.6, 10 and 11, example values: "9.5", "9.6", "10", "11". Microsoft SQL Server databases support 2014 SE and 2014 EE, example values: "2014_SE", "2014_EE". Changing this parameter will create a new resource. The ` + "`" + `volume` + "`" + ` block supports:`,
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
					Description: `(Required) Specifies the volume type. Its value can be any of the following and is case-sensitive: COMMON: indicates the SATA type. ULTRAHIGH: indicates the SSD type. Changing this parameter will create a new resource. The ` + "`" + `backup_strategy` + "`" + ` block supports:`,
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
					Description: `Indicates the public IP address list.`,
				},
				resource.Attribute{
					Name:        "db",
					Description: `See Argument Reference above. The ` + "`" + `db` + "`" + ` block also contains:`,
				},
				resource.Attribute{
					Name:        "user_name",
					Description: `Indicates the default user name of database. The ` + "`" + `nodes` + "`" + ` block contains:`,
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
					Description: `Indicates the node status. ## Timeouts This resource provides the following timeouts configuration options: - ` + "`" + `create` + "`" + ` - Default is 30 minute. ## Import RDS instance can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_rds_instance_v3.instance_1 7117d38e-4c8f-4624-a505-bd96b97d024c ` + "`" + `` + "`" + `` + "`" + ` But due to some attrubutes missing from the API response, it's required to ignore changes as below. ` + "`" + `` + "`" + `` + "`" + ` resource "flexibleengine_rds_instance_v3" "instance_1" { ... lifecycle { ignore_changes = [ "db", ] } } ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
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
					Description: `Indicates the public IP address list.`,
				},
				resource.Attribute{
					Name:        "db",
					Description: `See Argument Reference above. The ` + "`" + `db` + "`" + ` block also contains:`,
				},
				resource.Attribute{
					Name:        "user_name",
					Description: `Indicates the default user name of database. The ` + "`" + `nodes` + "`" + ` block contains:`,
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
					Description: `Indicates the node status. ## Timeouts This resource provides the following timeouts configuration options: - ` + "`" + `create` + "`" + ` - Default is 30 minute. ## Import RDS instance can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_rds_instance_v3.instance_1 7117d38e-4c8f-4624-a505-bd96b97d024c ` + "`" + `` + "`" + `` + "`" + ` But due to some attrubutes missing from the API response, it's required to ignore changes as below. ` + "`" + `` + "`" + `` + "`" + ` resource "flexibleengine_rds_instance_v3" "instance_1" { ... lifecycle { ignore_changes = [ "db", ] } } ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_rds_parametergroup_v3",
			Category:         "RDS Resources",
			ShortDescription: `Manages a V3 RDS parametergroup resource.`,
			Description:      ``,
			Keywords: []string{
				"rds",
				"parametergroup",
				"v3",
			},
			Arguments: []resource.Attribute{
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
					Description: `Indicates the parameter description. ## Import Parameter groups can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_rds_parametergroup_v3.pg_1 7117d38e-4c8f-4624-a505-bd96b97d024c ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
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
					Description: `Indicates the parameter description. ## Import Parameter groups can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_rds_parametergroup_v3.pg_1 7117d38e-4c8f-4624-a505-bd96b97d024c ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_rds_read_replica_v3",
			Category:         "RDS Resources",
			ShortDescription: `read replica management`,
			Description:      ``,
			Keywords: []string{
				"rds",
				"read",
				"replica",
				"v3",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Specifies the DB instance name. The DB instance name of the same type must be unique for the same tenant. The value must be 4 to 64 characters in length and start with a letter. It is case-sensitive and can contain only letters, digits, hyphens (-), and underscores (_). Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "flavor",
					Description: `(Required) Specifies the specification code. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "replica_of_id",
					Description: `(Required) Specifies the DB instance ID, which is used to create a read replica. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "volume",
					Description: `(Required) Specifies the volume information. Structure is documented below. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `(Required) Specifies the AZ name. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to create the instance. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Currently, read replicas can be created only in the same region as that of the promary DB instance. Changing this parameter will create a new resource. The ` + "`" + `volume` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) Specifies the volume type. Its value can be any of the following and is case-sensitive: - ULTRAHIGH: indicates the SSD type. - ULTRAHIGHPRO: indicates the ultra-high I/O. Changing this parameter will create a new resource.`,
				},
				resource.Attribute{
					Name:        "disk_encryption_id",
					Description: `(Optional) Specifies the key ID for disk encryption. Changing this parameter will create a new resource. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Indicates the instance ID.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Indicates the instance status.`,
				},
				resource.Attribute{
					Name:        "db",
					Description: `Indicates the database information. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "private_ips",
					Description: `Indicates the private IP address list.`,
				},
				resource.Attribute{
					Name:        "public_ips",
					Description: `Indicates the public IP address list.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `Indicates the security group which the RDS DB instance belongs to.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `Indicates the subnet id.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `Indicates the VPC ID. The ` + "`" + `db` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Indicates the database port information.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Indicates the DB engine. Value: MySQL, PostgreSQL, SQLServer.`,
				},
				resource.Attribute{
					Name:        "user_name",
					Description: `Indicates the default user name of database.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `Indicates the database version. ## Import RDS instance can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_rds_read_replica_v3.instance_1 7117d38e-4c8f-4624-a505-bd96b97d024c ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Indicates the instance ID.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Indicates the instance status.`,
				},
				resource.Attribute{
					Name:        "db",
					Description: `Indicates the database information. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "private_ips",
					Description: `Indicates the private IP address list.`,
				},
				resource.Attribute{
					Name:        "public_ips",
					Description: `Indicates the public IP address list.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `Indicates the security group which the RDS DB instance belongs to.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `Indicates the subnet id.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `Indicates the VPC ID. The ` + "`" + `db` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Indicates the database port information.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Indicates the DB engine. Value: MySQL, PostgreSQL, SQLServer.`,
				},
				resource.Attribute{
					Name:        "user_name",
					Description: `Indicates the default user name of database.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `Indicates the database version. ## Import RDS instance can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_rds_read_replica_v3.instance_1 7117d38e-4c8f-4624-a505-bd96b97d024c ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_rts_software_config_v1",
			Category:         "RTS Resources",
			ShortDescription: `Provides an RTS software config resource.`,
			Description:      ``,
			Icon:             "Management&Deployment-RTS.svg",
			Keywords: []string{
				"rts",
				"software",
				"config",
				"v1",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_resource_rts_stack_v1",
			Category:         "RTS Resources",
			ShortDescription: `Provides an FlexibleEngine RTS Stack.`,
			Description:      ``,
			Icon:             "Management&Deployment-RTS.svg",
			Keywords: []string{
				"rts",
				"resource",
				"stack",
				"v1",
			},
			Arguments: []resource.Attribute{
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
					Description: `Specifies the stack status. ## Import RTS Stacks can be imported using the ` + "`" + `name` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_rts_stack_v1.mystack rts-stack ` + "`" + `` + "`" + `` + "`" + ` <a id="timeouts"></a> ## Timeouts ` + "`" + `flexibleengine_rts_stack_v1` + "`" + ` provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - (Default ` + "`" + `30 minutes` + "`" + `) Used for Creating Stacks - ` + "`" + `update` + "`" + ` - (Default ` + "`" + `30 minutes` + "`" + `) Used for Stack modifications - ` + "`" + `delete` + "`" + ` - (Default ` + "`" + `30 minutes` + "`" + `) Used for destroying stacks.`,
				},
			},
			Attributes: []resource.Attribute{
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
					Description: `Specifies the stack status. ## Import RTS Stacks can be imported using the ` + "`" + `name` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_rts_stack_v1.mystack rts-stack ` + "`" + `` + "`" + `` + "`" + ` <a id="timeouts"></a> ## Timeouts ` + "`" + `flexibleengine_rts_stack_v1` + "`" + ` provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - (Default ` + "`" + `30 minutes` + "`" + `) Used for Creating Stacks - ` + "`" + `update` + "`" + ` - (Default ` + "`" + `30 minutes` + "`" + `) Used for Stack modifications - ` + "`" + `delete` + "`" + ` - (Default ` + "`" + `30 minutes` + "`" + `) Used for destroying stacks.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_s3_bucket",
			Category:         "S3 Resource",
			ShortDescription: `Provides a S3 bucket resource.`,
			Description:      ``,
			Icon:             "bucket.svg",
			Keywords: []string{
				"s3",
				"resource",
				"bucket",
			},
			Arguments: []resource.Attribute{
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
					Description: `(Optional) A configuration of [object lifecycle management](http://docs.aws.amazon.com/AmazonS3/latest/dev/object-lifecycle-mgmt.html) (documented below).`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) If specified, the region this bucket should reside in. Otherwise, the region used by the callee. The ` + "`" + `website` + "`" + ` object supports the following:`,
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
					Name:        "region",
					Description: `The region this bucket resides in.`,
				},
				resource.Attribute{
					Name:        "website_endpoint",
					Description: `The website endpoint, if the bucket is configured with a website. If not, this will be an empty string.`,
				},
				resource.Attribute{
					Name:        "website_domain",
					Description: `The domain of the website endpoint, if the bucket is configured with a website. If not, this will be an empty string. This is used to create Route 53 alias records. ## Import S3 bucket can be imported using the ` + "`" + `bucket` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_s3_bucket.bucket bucket-name ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
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
					Name:        "region",
					Description: `The region this bucket resides in.`,
				},
				resource.Attribute{
					Name:        "website_endpoint",
					Description: `The website endpoint, if the bucket is configured with a website. If not, this will be an empty string.`,
				},
				resource.Attribute{
					Name:        "website_domain",
					Description: `The domain of the website endpoint, if the bucket is configured with a website. If not, this will be an empty string. This is used to create Route 53 alias records. ## Import S3 bucket can be imported using the ` + "`" + `bucket` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_s3_bucket.bucket bucket-name ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_s3_bucket_object",
			Category:         "S3 Resource",
			ShortDescription: `Provides a S3 bucket object resource.`,
			Description:      ``,
			Icon:             "bucket.svg",
			Keywords: []string{
				"s3",
				"resource",
				"bucket",
				"object",
			},
			Arguments: []resource.Attribute{
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
			Attributes: []resource.Attribute{
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
			Type:             "flexibleengine_s3_bucket_policy",
			Category:         "S3 Resource",
			ShortDescription: `Attaches a policy to an S3 bucket resource.`,
			Description:      ``,
			Icon:             "bucket.svg",
			Keywords: []string{
				"s3",
				"resource",
				"bucket",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "bucket",
					Description: `(Required) The name of the bucket to which to apply the policy.`,
				},
				resource.Attribute{
					Name:        "policy",
					Description: `(Required) The text of the policy.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_sdrs_drill_v1",
			Category:         "SDRS Resources",
			ShortDescription: `Manages a Disaster Recovery Drill resource within FlexibleEngine.`,
			Description:      ``,
			Keywords: []string{
				"sdrs",
				"drill",
				"v1",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of a DR drill. The name can contain a maximum of 64 bytes. The value can contain only letters (a to z and A to Z), digits (0 to 9), decimal points (.), underscores (_), and hyphens (-).`,
				},
				resource.Attribute{
					Name:        "group_id",
					Description: `(Required) Specifies the ID of a protection group. Changing this creates a new drill.`,
				},
				resource.Attribute{
					Name:        "drill_vpc_id",
					Description: `(Required) Specifies the ID used for a DR drill. Changing this creates a new drill. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of a DR drill.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "group_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "drill_vpc_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of a DR drill. For details, see [DR Drill Status](https://docs.prod-cloud-ocb.orange-business.com/en-us/api/sdrs/en-us_topic_0126152933.html). ## Import DR drill can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_sdrs_drill_v1.drill_1 22fce838-4bfb-4a92-b9aa-fc80a583eb59 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of a DR drill.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "group_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "drill_vpc_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of a DR drill. For details, see [DR Drill Status](https://docs.prod-cloud-ocb.orange-business.com/en-us/api/sdrs/en-us_topic_0126152933.html). ## Import DR drill can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_sdrs_drill_v1.drill_1 22fce838-4bfb-4a92-b9aa-fc80a583eb59 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_sdrs_protectedinstance_v1",
			Category:         "SDRS Resources",
			ShortDescription: `Manages a V1 SDRS protected instance resource within FlexibleEngine.`,
			Description:      ``,
			Keywords: []string{
				"sdrs",
				"protectedinstance",
				"v1",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of a protected instance.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description of a protected instance. Changing this creates a new instance.`,
				},
				resource.Attribute{
					Name:        "group_id",
					Description: `(Required) Specifies the ID of the protection group where a protected instance is added. Changing this creates a new instance.`,
				},
				resource.Attribute{
					Name:        "server_id",
					Description: `(Required) Specifies the ID of the source server. Changing this creates a new instance.`,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Optional) Specifies the ID of a storage pool. Changing this creates a new instance.`,
				},
				resource.Attribute{
					Name:        "primary_subnet_id",
					Description: `(Optional) Specifies the subnet ID of the primary NIC on the target server. Changing this creates a new instance.`,
				},
				resource.Attribute{
					Name:        "primary_ip_address",
					Description: `(Optional) Specifies the IP address of the primary NIC on the target server. Changing this creates a new instance.`,
				},
				resource.Attribute{
					Name:        "delete_target_server",
					Description: `(Optional) Specifies whether to delete the target server. The default value is false.. Changing this creates a new instance.`,
				},
				resource.Attribute{
					Name:        "delete_target_eip",
					Description: `(Optional) Specifies whether to delete the EIP of the target server. The default value is false. Changing this creates a new instance. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the protected instance.`,
				},
				resource.Attribute{
					Name:        "target_server",
					Description: `ID of the target server.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the protected instance.`,
				},
				resource.Attribute{
					Name:        "target_server",
					Description: `ID of the target server.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_sdrs_protectiongroup_v1",
			Category:         "SDRS Resources",
			ShortDescription: `Manages a V1 SDRS protection group resource within FlexibleEngine.`,
			Description:      ``,
			Keywords: []string{
				"sdrs",
				"protectiongroup",
				"v1",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of a protection group.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description of a protection group. Changing this creates a new group.`,
				},
				resource.Attribute{
					Name:        "source_availability_zone",
					Description: `(Required) Specifies the source AZ of a protection group. Changing this creates a new group.`,
				},
				resource.Attribute{
					Name:        "target_availability_zone",
					Description: `(Required) Specifies the target AZ of a protection group. Changing this creates a new group.`,
				},
				resource.Attribute{
					Name:        "domain_id",
					Description: `(Required) Specifies the ID of an active-active domain. Changing this creates a new group.`,
				},
				resource.Attribute{
					Name:        "source_vpc_id",
					Description: `(Required) Specifies the ID of the source VPC. Changing this creates a new group.`,
				},
				resource.Attribute{
					Name:        "dr_type",
					Description: `(Optional) Specifies the deployment model. The default value is migration indicating migration within a VPC. Changing this creates a new group.`,
				},
				resource.Attribute{
					Name:        "enable",
					Description: `(Optional) Enable protection or not. It can only be set to true when there's replication pairs within the protection group. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the protection group. ## Import Protection groups can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_sdrs_protectiongroup_v1.group_1 7117d38e-4c8f-4624-a505-bd96b97d024c ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the protection group. ## Import Protection groups can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_sdrs_protectiongroup_v1.group_1 7117d38e-4c8f-4624-a505-bd96b97d024c ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_sdrs_replication_attach_v1",
			Category:         "SDRS Resources",
			ShortDescription: `Manages a V1 SDRS replication attach resource within FlexibleEngine.`,
			Description:      ``,
			Keywords: []string{
				"sdrs",
				"replication",
				"attach",
				"v1",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Required) Specifies the ID of a protected instance. Changing this creates a new replication attach.`,
				},
				resource.Attribute{
					Name:        "replication_id",
					Description: `(Required) Specifies the ID of a replication pair. Changing this creates a new replication attach.`,
				},
				resource.Attribute{
					Name:        "device",
					Description: `(Required) Specifies the device name, eg. /dev/vdb. Changing this creates a new replication attach.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_sdrs_replication_pair_v1",
			Category:         "SDRS Resources",
			ShortDescription: `Manages a V1 SDRS replication pair resource within FlexibleEngine.`,
			Description:      ``,
			Keywords: []string{
				"sdrs",
				"replication",
				"pair",
				"v1",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of a replication pair. The name can contain a maximum of 64 bytes. The value can contain only letters (a to z and A to Z), digits (0 to 9), decimal points (.), underscores (_), and hyphens (-).`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description of a replication pair. Changing this creates a new pair.`,
				},
				resource.Attribute{
					Name:        "group_id",
					Description: `(Required) Specifies the ID of a protection group. Changing this creates a new pair.`,
				},
				resource.Attribute{
					Name:        "volume_id",
					Description: `(Required) Specifies the ID of a source disk. Changing this creates a new pair.`,
				},
				resource.Attribute{
					Name:        "delete_target_volume",
					Description: `(Optional) Specifies whether to delete the target disk. The default value is ` + "`" + `false` + "`" + `. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the replication pair.`,
				},
				resource.Attribute{
					Name:        "fault_level",
					Description: `Specifies the fault level of a replication pair.`,
				},
				resource.Attribute{
					Name:        "replication_model",
					Description: `Specifies the replication mode of a replication pair. The default value is ` + "`" + `hypermetro` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Specifies the status of a replication pair.`,
				},
				resource.Attribute{
					Name:        "target_volume_id",
					Description: `Specifies the ID of the disk in the protection availability zone. ## Import Replication pairs can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_sdrs_replication_pair_v1.replication_1 43b28b66-770b-4e9e-b5c6-cfc43f0593d9 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the replication pair.`,
				},
				resource.Attribute{
					Name:        "fault_level",
					Description: `Specifies the fault level of a replication pair.`,
				},
				resource.Attribute{
					Name:        "replication_model",
					Description: `Specifies the replication mode of a replication pair. The default value is ` + "`" + `hypermetro` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Specifies the status of a replication pair.`,
				},
				resource.Attribute{
					Name:        "target_volume_id",
					Description: `Specifies the ID of the disk in the protection availability zone. ## Import Replication pairs can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_sdrs_replication_pair_v1.replication_1 43b28b66-770b-4e9e-b5c6-cfc43f0593d9 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_sfs_file_system_v2",
			Category:         "SFS Resources",
			ShortDescription: `Provides an Scalable File Resource (SFS) resource.`,
			Description:      ``,
			Icon:             "Storage-SFS.svg",
			Keywords: []string{
				"sfs",
				"file",
				"system",
				"v2",
			},
			Arguments: []resource.Attribute{
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
					Description: `The status of the share access rule. ## Import SFS can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_sfs_file_system_v2 4779ab1c-7c1a-44b1-a02e-93dfc361b32d ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
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
					Description: `The status of the share access rule. ## Import SFS can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_sfs_file_system_v2 4779ab1c-7c1a-44b1-a02e-93dfc361b32d ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_smn_subscription_v2",
			Category:         "SMN Resources",
			ShortDescription: `Manages a V2 subscription resource within FlexibleEngine.`,
			Description:      ``,
			Icon:             "Application-SMN.svg",
			Keywords: []string{
				"smn",
				"subscription",
				"v2",
			},
			Arguments: []resource.Attribute{
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
			Attributes: []resource.Attribute{
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
			Type:             "flexibleengine_smn_topic_v2",
			Category:         "SMN Resources",
			ShortDescription: `Manages a V2 topic resource within FlexibleEngine.`,
			Description:      ``,
			Icon:             "Application-SMN.svg",
			Keywords: []string{
				"smn",
				"topic",
				"v2",
			},
			Arguments: []resource.Attribute{
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
			Attributes: []resource.Attribute{
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
			Type:             "flexibleengine-vbs-backup-policy-v2",
			Category:         "VBS Resources",
			ShortDescription: `Provides an VBS Backup Policy resource.`,
			Description:      ``,
			Icon:             "Storage-VBS.svg",
			Keywords: []string{
				"vbs",
				"flexibleengine-vbs-backup-policy-v2",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine-vbs-backup-v2",
			Category:         "VBS Resources",
			ShortDescription: `Provides an VBS Backup resource.`,
			Description:      ``,
			Icon:             "Storage-VBS.svg",
			Keywords: []string{
				"vbs",
				"flexibleengine-vbs-backup-v2",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_vpc_eip_v1",
			Category:         "EIP Resources",
			ShortDescription: `Manages a V1 EIP resource within FlexibleEngine VPC.`,
			Description:      ``,
			Icon:             "Network-EIP.svg",
			Keywords: []string{
				"eip",
				"vpc",
				"v1",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to create the eip. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new eip.`,
				},
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
					Description: `(Required) The bandwidth size. The value ranges from 1 to 1000 Mbit/s.`,
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
					Name:        "region",
					Description: `See Argument Reference above.`,
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
					Description: `See Argument Reference above. ## Import EIPs can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_vpc_eip_v1.eip_1 2c7f39f3-702b-48d1-940c-b50384177ee1 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
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
					Description: `See Argument Reference above. ## Import EIPs can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_vpc_eip_v1.eip_1 2c7f39f3-702b-48d1-940c-b50384177ee1 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_vpc_peering_connection_accepter_v2",
			Category:         "Networking Resources",
			ShortDescription: `Manage the accepter's side of a VPC Peering Connection.`,
			Description:      ``,
			Icon:             "peer link.svg",
			Keywords: []string{
				"networking",
				"vpc",
				"peering",
				"connection",
				"accepter",
				"v2",
			},
			Arguments: []resource.Attribute{
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
			Attributes: []resource.Attribute{
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
			Type:             "flexibleengine_vpc_peering_connection_v2",
			Category:         "Networking Resources",
			ShortDescription: `Manage a VPC Peering Connection resource.`,
			Description:      ``,
			Icon:             "peer link.svg",
			Keywords: []string{
				"networking",
				"vpc",
				"peering",
				"connection",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The VPC peering connection ID.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The VPC peering connection status. The value can be PENDING_ACCEPTANCE, REJECTED, EXPIRED, DELETED, or ACTIVE. ## Notes If you create a VPC peering connection with another VPC of your own, the connection is created without the need for you to accept the connection. ## Import VPC Peering resources can be imported using the ` + "`" + `vpc peering id` + "`" + `, e.g. > $ terraform import flexibleengine_vpc_peering_connection_v2.test_connection 22b76469-08e3-4937-8c1d-7aad34892be1`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The VPC peering connection ID.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The VPC peering connection status. The value can be PENDING_ACCEPTANCE, REJECTED, EXPIRED, DELETED, or ACTIVE. ## Notes If you create a VPC peering connection with another VPC of your own, the connection is created without the need for you to accept the connection. ## Import VPC Peering resources can be imported using the ` + "`" + `vpc peering id` + "`" + `, e.g. > $ terraform import flexibleengine_vpc_peering_connection_v2.test_connection 22b76469-08e3-4937-8c1d-7aad34892be1`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_vpc_route_v2",
			Category:         "Networking Resources",
			ShortDescription: `Provides an VPC route resource.`,
			Description:      ``,
			Icon:             "Network-VPC.svg",
			Keywords: []string{
				"networking",
				"vpc",
				"route",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The route ID.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The route ID.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_vpc_subnet_v1",
			Category:         "Networking Resources",
			ShortDescription: `Provides an VPC subnet resource.`,
			Description:      ``,
			Icon:             "subnet.svg",
			Keywords: []string{
				"networking",
				"vpc",
				"subnet",
				"v1",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_vpc_v1",
			Category:         "Networking Resources",
			ShortDescription: `Manages a V1 VPC resource within FlexibleEngine.`,
			Description:      ``,
			Icon:             "Network-VPC.svg",
			Keywords: []string{
				"networking",
				"vpc",
				"v1",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cidr",
					Description: `(Required) The range of available subnets in the VPC. The value ranges from 10.0.0.0/8 to 10.255.255.0/24, 172.16.0.0/12 to 172.31.255.0/24, or 192.168.0.0/16 to 192.168.255.0/24.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V1 VPC client. A VPC client is needed to create a VPC. If omitted, the region argument of the provider is used. Changing this creates a new VPC.`,
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
					Description: `Specifies whether the cross-tenant sharing is supported.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above. ## Import VPCs can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_vpc_v1.vpc_v1 7117d38e-4c8f-4624-a505-bd96b97d024c ` + "`" + `` + "`" + `` + "`" + ``,
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
					Name:        "cidr",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The current status of the desired VPC. Can be either CREATING, OK, DOWN, PENDING_UPDATE, PENDING_DELETE, or ERROR.`,
				},
				resource.Attribute{
					Name:        "shared",
					Description: `Specifies whether the cross-tenant sharing is supported.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above. ## Import VPCs can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import flexibleengine_vpc_v1.vpc_v1 7117d38e-4c8f-4624-a505-bd96b97d024c ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
	}

	resourcesMap = map[string]int{

		"flexibleengine_antiddos_v1":                        0,
		"flexibleengine_as_configuration_v1":                1,
		"flexibleengine_as_group_v1":                        2,
		"flexibleengine_as_policy_v1":                       3,
		"flexibleengine_blockstorage_volume_v2":             4,
		"flexibleengine_cce_cluster_v3":                     5,
		"flexibleengine_cce_nodes_v3":                       6,
		"flexibleengine_ces-alarmrule":                      7,
		"flexibleengine_compute_server_v2":                  8,
		"flexibleengine_compute_floatingip_associate_v2":    9,
		"flexibleengine_compute_floatingip_v2":              10,
		"flexibleengine_compute_instance_v2":                11,
		"flexibleengine_compute_interface_attach_v2":        12,
		"flexibleengine_compute_keypair_v2":                 13,
		"flexibleengine_compute_servergroup_v2":             14,
		"flexibleengine_compute_volume_attach_v2":           15,
		"flexibleengine_csbs_backup_policy_v1":              16,
		"flexibleengine_csbs_backup_v1":                     17,
		"flexibleengine_css_cluster_v1":                     18,
		"flexibleengine_cts_tracker_v1":                     19,
		"flexibleengine_dcs_instance_v1":                    20,
		"flexibleengine_dds_instance_v3":                    21,
		"flexibleengine_dns_recordset_v2":                   22,
		"flexibleengine_dns_zone_v2":                        23,
		"flexibleengine_drs_replication_v2":                 24,
		"flexibleengine_drs_replicationconsistencygroup_v2": 25,
		"flexibleengine_dws_cluster_v1":                     26,
		"flexibleengine_elb_backend":                        27,
		"flexibleengine_elb_health":                         28,
		"flexibleengine_elb_listener":                       29,
		"flexibleengine_elb_loadbalancer":                   30,
		"flexibleengine_fw_firewall_group_v2":               31,
		"flexibleengine_fw_policy_v2":                       32,
		"flexibleengine_fw_rule_v2":                         33,
		"flexibleengine_identity_agency_v3":                 34,
		"flexibleengine_identity_group_membership_v3":       35,
		"flexibleengine_identity_group_v3":                  36,
		"flexibleengine_identity_role_assignment_v3":        37,
		"flexibleengine_identity_role_v3":                   38,
		"flexibleengine_identity_user_v3":                   39,
		"flexibleengine_images_image_v2":                    40,
		"flexibleengine_kms_key_v1":                         41,
		"flexibleengine_lb_certificate_v2":                  42,
		"flexibleengine_lb_l7policy_v2":                     43,
		"flexibleengine_lb_l7rule_v2":                       44,
		"flexibleengine_lb_listener_v2":                     45,
		"flexibleengine_lb_loadbalancer_v2":                 46,
		"flexibleengine_lb_member_v2":                       47,
		"flexibleengine_lb_monitor_v2":                      48,
		"flexibleengine_lb_pool_v2":                         49,
		"flexibleengine_lb_whitelist_v2":                    50,
		"flexibleengine_mls_instance_v1":                    51,
		"flexibleengine_mrs_cluster_v1":                     52,
		"flexibleengine_mrs_hybrid_cluster_v1":              53,
		"flexibleengine_mrs_job_v1":                         54,
		"flexibleengine_nat_dnat_rule_v2":                   55,
		"flexibleengine_nat_gateway_v2":                     56,
		"flexibleengine_nat_snat_rule_v2":                   57,
		"flexibleengine_networking_floatingip_associate_v2": 58,
		"flexibleengine_networking_floatingip_v2":           59,
		"flexibleengine_networking_network_v2":              60,
		"flexibleengine_networking_port_v2":                 61,
		"flexibleengine_networking_router_interface_v2":     62,
		"flexibleengine_networking_router_route_v2":         63,
		"flexibleengine_networking_router_v2":               64,
		"flexibleengine_networking_secgroup_rule_v2":        65,
		"flexibleengine_networking_secgroup_v2":             66,
		"flexibleengine_networking_subnet_v2":               67,
		"flexibleengine_networking_vip_associate_v2":        68,
		"flexibleengine_networking_vip_v2":                  69,
		"flexibleengine_obs_bucket":                         70,
		"flexibleengine_obs_bucket_object":                  71,
		"flexibleengine_rds_instance_v1":                    72,
		"flexibleengine_rds_instance_v3":                    73,
		"flexibleengine_rds_parametergroup_v3":              74,
		"flexibleengine_rds_read_replica_v3":                75,
		"flexibleengine_rts_software_config_v1":             76,
		"flexibleengine_resource_rts_stack_v1":              77,
		"flexibleengine_s3_bucket":                          78,
		"flexibleengine_s3_bucket_object":                   79,
		"flexibleengine_s3_bucket_policy":                   80,
		"flexibleengine_sdrs_drill_v1":                      81,
		"flexibleengine_sdrs_protectedinstance_v1":          82,
		"flexibleengine_sdrs_protectiongroup_v1":            83,
		"flexibleengine_sdrs_replication_attach_v1":         84,
		"flexibleengine_sdrs_replication_pair_v1":           85,
		"flexibleengine_sfs_file_system_v2":                 86,
		"flexibleengine_smn_subscription_v2":                87,
		"flexibleengine_smn_topic_v2":                       88,
		"flexibleengine-vbs-backup-policy-v2":               89,
		"flexibleengine-vbs-backup-v2":                      90,
		"flexibleengine_vpc_eip_v1":                         91,
		"flexibleengine_vpc_peering_connection_accepter_v2": 92,
		"flexibleengine_vpc_peering_connection_v2":          93,
		"flexibleengine_vpc_route_v2":                       94,
		"flexibleengine_vpc_subnet_v1":                      95,
		"flexibleengine_vpc_v1":                             96,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
