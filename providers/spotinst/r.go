package spotinst

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "elastigroup_aws",
			Category:         "Spotinst Resources",
			ShortDescription: `Provides a Spotinst AWS group resource.`,
			Description:      ``,
			Keywords: []string{
				"elastigroup",
				"aws",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The group name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The group description.`,
				},
				resource.Attribute{
					Name:        "product",
					Description: `(Required) Operation system type. Valid values: ` + "`" + `"Linux/UNIX"` + "`" + `, ` + "`" + `"SUSE Linux"` + "`" + `, ` + "`" + `"Windows"` + "`" + `. For EC2 Classic instances: ` + "`" + `"Linux/UNIX (Amazon VPC)"` + "`" + `, ` + "`" + `"SUSE Linux (Amazon VPC)"` + "`" + `, ` + "`" + `"Windows (Amazon VPC)"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "availability_zones",
					Description: `(Optional) List of Strings of availability zones. When this parameter is set, ` + "`" + `subnet_ids` + "`" + ` should be left unused. Note: ` + "`" + `availability_zones` + "`" + ` naming syntax follows the convention ` + "`" + `availability-zone:subnet:placement-group-name` + "`" + `. For example, to set an AZ in ` + "`" + `us-east-1` + "`" + ` with subnet ` + "`" + `subnet-123456` + "`" + ` and placement group ` + "`" + `ClusterI03` + "`" + `, you would set: ` + "`" + `availability_zones = ["us-east-1a:subnet-123456:ClusterI03"]` + "`" + ``,
				},
				resource.Attribute{
					Name:        "subnet_ids",
					Description: `(Optional) List of Strings of subnet identifiers. Note: When this parameter is set, ` + "`" + `availability_zones` + "`" + ` should be left unused.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The AWS region your group will be created in. Note: This parameter is required if you specify subnets (through subnet_ids). This parameter is optional if you specify Availability Zones (through availability_zones).`,
				},
				resource.Attribute{
					Name:        "preferred_availability_zones",
					Description: `The AZs to prioritize when launching Spot instances. If no markets are available in the Preferred AZs, Spot instances are launched in the non-preferred AZs. Note: Must be a sublist of ` + "`" + `availability_zones` + "`" + ` and ` + "`" + `orientation` + "`" + ` value must not be ` + "`" + `"equalAzDistribution"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "max_size",
					Description: `(Optional; Required if using scaling policies) The maximum number of instances the group should have at any time.`,
				},
				resource.Attribute{
					Name:        "min_size",
					Description: `(Optional; Required if using scaling policies) The minimum number of instances the group should have at any time.`,
				},
				resource.Attribute{
					Name:        "desired_capacity",
					Description: `(Optional) The desired number of instances the group should have at any time.`,
				},
				resource.Attribute{
					Name:        "capacity_unit",
					Description: `(Optional, Default: ` + "`" + `"instance"` + "`" + `) The capacity unit to launch instances by. If not specified, when choosing the weight unit, each instance will weight as the number of its vCPUs.`,
				},
				resource.Attribute{
					Name:        "security_groups",
					Description: `(Required) A list of associated security group IDS.`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `(Optional) The ID of the AMI used to launch the instance.`,
				},
				resource.Attribute{
					Name:        "iam_instance_profile",
					Description: `(Optional) The ARN or name of an IAM instance profile to associate with launched instances.`,
				},
				resource.Attribute{
					Name:        "key_name",
					Description: `(Optional) The key name that should be used for the instance.`,
				},
				resource.Attribute{
					Name:        "enable_monitoring",
					Description: `(Optional) Indicates whether monitoring is enabled for the instance.`,
				},
				resource.Attribute{
					Name:        "user_data",
					Description: `(Optional) The user data to provide when launching the instance.`,
				},
				resource.Attribute{
					Name:        "shutdown_script",
					Description: `(Optional) The Base64-encoded shutdown script that executes prior to instance termination, for more information please see: [Shutdown Script](https://api.spotinst.com/integration-docs/elastigroup/concepts/compute-concepts/shutdown-scripts/)`,
				},
				resource.Attribute{
					Name:        "ebs_optimized",
					Description: `(Optional) Enable high bandwidth connectivity between instances and AWS’s Elastic Block Store (EBS). For instance types that are EBS-optimized by default this parameter will be ignored.`,
				},
				resource.Attribute{
					Name:        "placement_tenancy",
					Description: `(Optional) Enable dedicated tenancy. Note: There is a flat hourly fee for each region in which dedicated tenancy is used.`,
				},
				resource.Attribute{
					Name:        "instance_types_ondemand",
					Description: `(Required) The type of instance determines your instance's CPU capacity, memory and storage (e.g., m1.small, c1.xlarge).`,
				},
				resource.Attribute{
					Name:        "instance_types_spot",
					Description: `(Required) One or more instance types.`,
				},
				resource.Attribute{
					Name:        "instance_types_preferred_spot",
					Description: `(Optional) Prioritize a subset of spot instance types. Must be a subset of the selected spot instance types.`,
				},
				resource.Attribute{
					Name:        "instance_types_weights",
					Description: `(Optional) List of weights per instance type for weighted groups. Each object in the list should have the following attributes:`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `(Required) Weight per instance type (Integer).`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `(Required) Name of instance type (String).`,
				},
				resource.Attribute{
					Name:        "cpu_credits",
					Description: `(Optional) Controls how T3 instances are launched. Valid values: ` + "`" + `standard` + "`" + `, ` + "`" + `unlimited` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "fallback_to_ondemand",
					Description: `(Required) In a case of no Spot instances available, Elastigroup will launch on-demand instances instead.`,
				},
				resource.Attribute{
					Name:        "wait_for_capacity",
					Description: `(Optional) Minimum number of instances in a 'HEALTHY' status that is required before continuing. This is ignored when updating with blue/green deployment. Cannot exceed ` + "`" + `desired_capacity` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "wait_for_capacity_timeout",
					Description: `(Optional) Time (seconds) to wait for instances to report a 'HEALTHY' status. Useful for plans with multiple dependencies that take some time to initialize. Leave undefined or set to ` + "`" + `0` + "`" + ` to indicate no wait. This is ignored when updating with blue/green deployment.`,
				},
				resource.Attribute{
					Name:        "orientation",
					Description: `(Required, Default: ` + "`" + `"balanced"` + "`" + `) Select a prediction strategy. Valid values: ` + "`" + `"balanced"` + "`" + `, ` + "`" + `"costOriented"` + "`" + `, ` + "`" + `"equalAzDistribution"` + "`" + `, ` + "`" + `"availabilityOriented"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "spot_percentage",
					Description: `(Optional; Required if not using ` + "`" + `ondemand_count` + "`" + `) The percentage of Spot instances that would spin up from the ` + "`" + `desired_capacity` + "`" + ` number.`,
				},
				resource.Attribute{
					Name:        "ondemand_count",
					Description: `(Optional; Required if not using ` + "`" + `spot_percentage` + "`" + `) Number of on demand instances to launch in the group. All other instances will be spot instances. When this parameter is set the ` + "`" + `spot_percentage` + "`" + ` parameter is being ignored.`,
				},
				resource.Attribute{
					Name:        "draining_timeout",
					Description: `(Optional) The time in seconds, the instance is allowed to run while detached from the ELB. This is to allow the instance time to be drained from incoming TCP connections before terminating it, during a scale down operation.`,
				},
				resource.Attribute{
					Name:        "utilize_reserved_instances",
					Description: `(Optional) In a case of any available reserved instances, Elastigroup will utilize them first before purchasing Spot instances.`,
				},
				resource.Attribute{
					Name:        "scaling_strategy",
					Description: `(Optional) Set termination policy.`,
				},
				resource.Attribute{
					Name:        "terminate_at_end_of_billing_hour",
					Description: `(Optional) Specify whether to terminate instances at the end of each billing hour.`,
				},
				resource.Attribute{
					Name:        "termination_policy",
					Description: `(Optional) - Determines whether to terminate the newest instances when performing a scaling action. Valid values: ` + "`" + `"default"` + "`" + `, ` + "`" + `"newestInstance"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "health_check_type",
					Description: `(Optional) The service that will perform health checks for the instance. Valid values: ` + "`" + `"ELB"` + "`" + `, ` + "`" + `"HCS"` + "`" + `, ` + "`" + `"TARGET_GROUP"` + "`" + `, ` + "`" + `"MLB"` + "`" + `, ` + "`" + `"EC2"` + "`" + `, ` + "`" + `"MULTAI_TARGET_SET"` + "`" + `, ` + "`" + `"MLB_RUNTIME"` + "`" + `, ` + "`" + `"K8S_NODE"` + "`" + `, ` + "`" + `"NOMAD_NODE"` + "`" + `, ` + "`" + `"ECS_CLUSTER_INSTANCE"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "health_check_grace_period",
					Description: `(Optional) The amount of time, in seconds, after the instance has launched to starts and check its health.`,
				},
				resource.Attribute{
					Name:        "health_check_unhealthy_duration_before_replacement",
					Description: `(Optional) The amount of time, in seconds, that we will wait before replacing an instance that is running and became unhealthy (this is only applicable for instances that were once healthy).`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A key/value mapping of tags to assign to the resource.`,
				},
				resource.Attribute{
					Name:        "elastic_ips",
					Description: `(Optional) A list of [AWS Elastic IP](http://docs.aws.amazon.com/AWSEC2/latest/UserGuide/elastic-ip-addresses-eip.html) allocation IDs to associate to the group instances.`,
				},
				resource.Attribute{
					Name:        "revert_to_spot",
					Description: `(Optional) Hold settings for strategy correction – replacing On-Demand for Spot instances. Supported Values: ` + "`" + `"never"` + "`" + `, ` + "`" + `"always"` + "`" + `, ` + "`" + `"timeWindow"` + "`" + ``,
				},
				resource.Attribute{
					Name:        "perform_at",
					Description: `(Required) In the event of a fallback to On-Demand instances, select the time period to revert back to Spot. Supported Arguments – always (default), timeWindow, never. For timeWindow or never to be valid the group must have availabilityOriented OR persistence defined.`,
				},
				resource.Attribute{
					Name:        "time_windows",
					Description: `(Optional) Specify a list of time windows for to execute revertToSpot strategy. Time window format: ` + "`" + `ddd:hh:mm-ddd:hh:mm` + "`" + `. Example: ` + "`" + `Mon:03:00-Wed:02:30` + "`" + ` <a id="load-balancers"></a> ## Load Balancers`,
				},
				resource.Attribute{
					Name:        "elastic_load_balancers",
					Description: `(Optional) List of Elastic Load Balancers names (ELB).`,
				},
				resource.Attribute{
					Name:        "target_group_arns",
					Description: `(Optional) List of Target Group ARNs to register the instances to.`,
				},
				resource.Attribute{
					Name:        "multai_target_sets",
					Description: `(Optional) Set of targets to register.`,
				},
				resource.Attribute{
					Name:        "target_set_id",
					Description: `(Required) ID of Multai target set.`,
				},
				resource.Attribute{
					Name:        "balancer_id",
					Description: `(Required) ID of Multai Load Balancer. Usage: ` + "`" + `` + "`" + `` + "`" + `hcl elastic_load_balancers = ["bal5", "bal2"] target_group_arns = ["tg-arn"] multai_target_sets = [{ target_set_id = "ts-123", balancer_id = "bal-123" }, { target_set_id = "ts-234", balancer_id = "bal-234" }] ` + "`" + `` + "`" + `` + "`" + ` <a id="signal"></a> ## Signals Each ` + "`" + `signal` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the signal defined for the group. Valid Values: ` + "`" + `"INSTANCE_READY"` + "`" + `, ` + "`" + `"INSTANCE_READY_TO_SHUTDOWN"` + "`" + ``,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `(Optional) The signals defined timeout- default is 40 minutes (1800 seconds). Usage: ` + "`" + `` + "`" + `` + "`" + `hcl signal = { name = "INSTANCE_READY_TO_SHUTDOWN" timeout = 100 } ` + "`" + `` + "`" + `` + "`" + ` <a id="scheduled-task"></a> ## Scheduled Tasks Each ` + "`" + `scheduled_task` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "task_type",
					Description: `(Required) The task type to run. Supported task types are: ` + "`" + `"scale"` + "`" + `, ` + "`" + `"backup_ami"` + "`" + `, ` + "`" + `"roll"` + "`" + `, ` + "`" + `"scaleUp"` + "`" + `, ` + "`" + `"percentageScaleUp"` + "`" + `, ` + "`" + `"scaleDown"` + "`" + `, ` + "`" + `"percentageScaleDown"` + "`" + `, ` + "`" + `"statefulUpdateCapacity"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "cron_expression",
					Description: `(Optional; Required if not using ` + "`" + `frequency` + "`" + `) A valid cron expression. The cron is running in UTC time zone and is in [Unix cron format](https://en.wikipedia.org/wiki/Cron).`,
				},
				resource.Attribute{
					Name:        "start_time",
					Description: `(Optional; Format: ISO 8601) Set a start time for one time tasks.`,
				},
				resource.Attribute{
					Name:        "frequency",
					Description: `(Optional; Required if not using ` + "`" + `cron_expression` + "`" + `) The recurrence frequency to run this task. Supported values are ` + "`" + `"hourly"` + "`" + `, ` + "`" + `"daily"` + "`" + `, ` + "`" + `"weekly"` + "`" + ` and ` + "`" + `"continuous"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "scale_target_capacity",
					Description: `(Optional) The desired number of instances the group should have.`,
				},
				resource.Attribute{
					Name:        "scale_min_capacity",
					Description: `(Optional) The minimum number of instances the group should have.`,
				},
				resource.Attribute{
					Name:        "scale_max_capacity",
					Description: `(Optional) The maximum number of instances the group should have.`,
				},
				resource.Attribute{
					Name:        "is_enabled",
					Description: `(Optional, Default: ` + "`" + `true` + "`" + `) Setting the task to being enabled or disabled.`,
				},
				resource.Attribute{
					Name:        "target_capacity",
					Description: `(Optional; Only valid for statefulUpdateCapacity) The desired number of instances the group should have.`,
				},
				resource.Attribute{
					Name:        "min_capacity",
					Description: `(Optional; Only valid for statefulUpdateCapacity) The minimum number of instances the group should have.`,
				},
				resource.Attribute{
					Name:        "max_capacity",
					Description: `(Optional; Only valid for statefulUpdateCapacity) The maximum number of instances the group should have.`,
				},
				resource.Attribute{
					Name:        "batch_size_percentage",
					Description: `(Optional; Required when the ` + "`" + `task_type` + "`" + ` is ` + "`" + `"roll"` + "`" + `.) The percentage size of each batch in the scheduled deployment roll.`,
				},
				resource.Attribute{
					Name:        "grace_period",
					Description: `(Optional) The period of time (seconds) to wait before checking a batch's health after it's deployment.`,
				},
				resource.Attribute{
					Name:        "adjustment",
					Description: `(Optional; Min 1) The number of instances to add or remove.`,
				},
				resource.Attribute{
					Name:        "adjustment_percentage",
					Description: `(Optional; Min 1) The percentage of instances to add or remove. Usage: ` + "`" + `` + "`" + `` + "`" + `hcl scheduled_task = [{ task_type = "backup_ami" cron_expression = "" start_time = "1970-01-01T01:00:00Z" frequency = "hourly" scale_target_capacity = 5 scale_min_capacity = 0 scale_max_capacity = 10 is_enabled = false target_capacity = 5 min_capacity = 0 max_capacity = 10 batch_size_percentage = 33 grace_period = 300 }] ` + "`" + `` + "`" + `` + "`" + ` <a id="scaling-policy"></a> ## Scaling Policies Each ` + "`" + `scaling_`,
				},
				resource.Attribute{
					Name:        "namespace",
					Description: `(Required) The namespace for the alarm's associated metric.`,
				},
				resource.Attribute{
					Name:        "metric_name",
					Description: `(Required) The name of the metric, with or without spaces.`,
				},
				resource.Attribute{
					Name:        "threshold",
					Description: `(Required) The value against which the specified statistic is compared.`,
				},
				resource.Attribute{
					Name:        "policy_name",
					Description: `(Required) The name of the policy.`,
				},
				resource.Attribute{
					Name:        "statistic",
					Description: `(Optional, Default: ` + "`" + `"average"` + "`" + `) The metric statistics to return. For information about specific statistics go to [Statistics](http://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/index.html?CHAP_TerminologyandKeyConcepts.html#Statistic) in the Amazon CloudWatch Developer Guide.`,
				},
				resource.Attribute{
					Name:        "unit",
					Description: `(Required) The unit for the alarm's associated metric. Valid values: ` + "`" + `"percent` + "`" + `, ` + "`" + `"seconds"` + "`" + `, ` + "`" + `"microseconds"` + "`" + `, ` + "`" + `"milliseconds"` + "`" + `, ` + "`" + `"bytes"` + "`" + `, ` + "`" + `"kilobytes"` + "`" + `, ` + "`" + `"megabytes"` + "`" + `, ` + "`" + `"gigabytes"` + "`" + `, ` + "`" + `"terabytes"` + "`" + `, ` + "`" + `"bits"` + "`" + `, ` + "`" + `"kilobits"` + "`" + `, ` + "`" + `"megabits"` + "`" + `, ` + "`" + `"gigabits"` + "`" + `, ` + "`" + `"terabits"` + "`" + `, ` + "`" + `"count"` + "`" + `, ` + "`" + `"bytes/second"` + "`" + `, ` + "`" + `"kilobytes/second"` + "`" + `, ` + "`" + `"megabytes/second"` + "`" + `, ` + "`" + `"gigabytes/second"` + "`" + `, ` + "`" + `"terabytes/second"` + "`" + `, ` + "`" + `"bits/second"` + "`" + `, ` + "`" + `"kilobits/second"` + "`" + `, ` + "`" + `"megabits/second"` + "`" + `, ` + "`" + `"gigabits/second"` + "`" + `, ` + "`" + `"terabits/second"` + "`" + `, ` + "`" + `"count/second"` + "`" + `, ` + "`" + `"none"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "is_enabled",
					Description: `(Optional, Default: ` + "`" + `true` + "`" + `) Specifies whether the scaling policy described in this block is enabled.`,
				},
				resource.Attribute{
					Name:        "period",
					Description: `(Optional, Default: ` + "`" + `300` + "`" + `) The granularity, in seconds, of the returned datapoints. Period must be at least 60 seconds and must be a multiple of 60.`,
				},
				resource.Attribute{
					Name:        "evaluation_periods",
					Description: `(Optional, Default: ` + "`" + `1` + "`" + `) The number of periods over which data is compared to the specified threshold.`,
				},
				resource.Attribute{
					Name:        "cooldown",
					Description: `(Optional, Default: ` + "`" + `300` + "`" + `) The amount of time, in seconds, after a scaling activity completes and before the next scaling activity can start. If this parameter is not specified, the default cooldown period for the group applies.`,
				},
				resource.Attribute{
					Name:        "dimensions",
					Description: `(Optional) A list of dimensions describing qualities of the metric.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The dimension name.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) The dimension value.`,
				},
				resource.Attribute{
					Name:        "operator",
					Description: `(Optional, Scale Up Default: ` + "`" + `gte` + "`" + `, Scale Down Default: ` + "`" + `lte` + "`" + `) The operator to use in order to determine if the scaling policy is applicable. Valid values: ` + "`" + `"gt"` + "`" + `, ` + "`" + `"gte"` + "`" + `, ` + "`" + `"lt"` + "`" + `, ` + "`" + `"lte"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "source",
					Description: `(Optional) The source of the metric. Valid values: ` + "`" + `"cloudWatch"` + "`" + `, ` + "`" + `"spectrum"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "action_type",
					Description: `(Optional; if not using ` + "`" + `min_target_capacity` + "`" + ` or ` + "`" + `max_target_capacity` + "`" + `) The type of action to perform for scaling. Valid values: ` + "`" + `"adjustment"` + "`" + `, ` + "`" + `"percentageAdjustment"` + "`" + `, ` + "`" + `"setMaxTarget"` + "`" + `, ` + "`" + `"setMinTarget"` + "`" + `, ` + "`" + `"updateCapacity"` + "`" + `. If you do not specify an action type, you can only use – ` + "`" + `adjustment` + "`" + `, ` + "`" + `minTargetCapacity` + "`" + `, ` + "`" + `maxTargetCapacity` + "`" + `. While using action_type, please also set the following: When using ` + "`" + `adjustment` + "`" + ` – set the field ` + "`" + `adjustment` + "`" + ` When using ` + "`" + `percentageAdjustment` + "`" + ` - set the field ` + "`" + `adjustment` + "`" + ` When using ` + "`" + `setMaxTarget` + "`" + ` – set the field ` + "`" + `max_target_capacity` + "`" + ` When using ` + "`" + `setMinTarget` + "`" + ` – set the field ` + "`" + `min_target_capacity` + "`" + ` When using ` + "`" + `updateCapacity` + "`" + ` – set the fields ` + "`" + `minimum` + "`" + `, ` + "`" + `maximum` + "`" + `, and ` + "`" + `target` + "`" + ``,
				},
				resource.Attribute{
					Name:        "adjustment",
					Description: `(Optional; if not using ` + "`" + `min_target_capacity` + "`" + ` or ` + "`" + `max_target_capacity` + "`" + `;) The number of instances to add/remove to/from the target capacity when scale is needed. Can be used as advanced expression for scaling of instances to add/remove to/from the target capacity when scale is needed. You can see more information here: Advanced expression. Example value: ` + "`" + `"MAX(currCapacity / 5, value`,
				},
				resource.Attribute{
					Name:        "min_target_capacity",
					Description: `(Optional; if not using ` + "`" + `adjustment` + "`" + `; available only for scale up). The number of the desired target (and minimum) capacity`,
				},
				resource.Attribute{
					Name:        "max_target_capacity",
					Description: `(Optional; if not using ` + "`" + `adjustment` + "`" + `; available only for scale down). The number of the desired target (and maximum) capacity`,
				},
				resource.Attribute{
					Name:        "minimum",
					Description: `(Optional; if using ` + "`" + `updateCapacity` + "`" + `) The minimal number of instances to have in the group.`,
				},
				resource.Attribute{
					Name:        "maximum",
					Description: `(Optional; if using ` + "`" + `updateCapacity` + "`" + `) The maximal number of instances to have in the group.`,
				},
				resource.Attribute{
					Name:        "target",
					Description: `(Optional; if using ` + "`" + `updateCapacity` + "`" + `) The target number of instances to have in the group. ` + "`" + `scaling_target_policies` + "`" + ` support predictive scaling:`,
				},
				resource.Attribute{
					Name:        "predictive_mode",
					Description: `(Optional) Start a metric prediction process to determine the expected target metric value within the next two days. See [Predictive Autoscaling](https://api.spotinst.com/elastigroup-for-aws/concepts/scaling-concepts/predictive-autoscaling/) documentation for more info. Valid values: ` + "`" + `FORECAST_AND_SCALE` + "`" + `, ` + "`" + `FORECAST_ONLY` + "`" + `. Usage: ` + "`" + `` + "`" + `` + "`" + `hcl scaling_up_policy = [{ policy_name = "policy-name" metric_name = "CPUUtilization" namespace = "AWS/EC2" source = "" statistic = "average" unit = "" cooldown = 60 is_enabled = false dimensions = { name = "name-1" value = "value-1" } threshold = 10 operator = "gt" evaluation_periods = 10 period = 60 // === MIN TARGET =================== action_type = "setMinTarget" min_target_capacity = 1 // ================================== // === ADJUSTMENT =================== # action_type = "adjustment" # action_type = "percentageAdjustment" # adjustment = "MAX(5,10)" // ================================== // === UPDATE CAPACITY ============== # action_type = "updateCapacity" # minimum = 0 # maximum = 10 # target = 5 // ================================== }] ` + "`" + `` + "`" + `` + "`" + ` ` + "`" + `` + "`" + `` + "`" + `hcl scaling_target_policy = [{ policy_name = "" metric_name = "" namespace = "" source = "" statistic = "" unit = "" cooldown = 10 target = 1 predictive_mode = "" dimensions = [{ name = "" value = "" }] }] ` + "`" + `` + "`" + `` + "`" + ` <a id="network-interface"></a> ## Network Interfaces Each of the ` + "`" + `network_interface` + "`" + ` attributes controls a portion of the AWS Instance's "Elastic Network Interfaces". It's a good idea to familiarize yourself with [AWS's Elastic Network Interfaces docs](http://docs.aws.amazon.com/AWSEC2/latest/UserGuide/using-eni.html) to understand the implications of using these attributes.`,
				},
				resource.Attribute{
					Name:        "network_interface_id",
					Description: `(Optional) The ID of the network interface.`,
				},
				resource.Attribute{
					Name:        "device_index",
					Description: `(Required) The index of the device on the instance for the network interface attachment.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Required) The description of the network interface.`,
				},
				resource.Attribute{
					Name:        "private_ip_address",
					Description: `(Optional) The private IP address of the network interface.`,
				},
				resource.Attribute{
					Name:        "delete_on_termination",
					Description: `(Optional) If set to true, the interface is deleted when the instance is terminated.`,
				},
				resource.Attribute{
					Name:        "secondary_private_ip_address_count",
					Description: `(Optional) The number of secondary private IP addresses.`,
				},
				resource.Attribute{
					Name:        "associate_public_ip_address",
					Description: `(Optional) Indicates whether to assign a public IP address to an instance you launch in a VPC. The public IP address can only be assigned to a network interface for eth0, and can only be assigned to a new network interface, not an existing one.`,
				},
				resource.Attribute{
					Name:        "associate_ipv6_address",
					Description: `(Optional) Indicates whether to assign IPV6 addresses to your instance. Requires a subnet with IPV6 CIDR block ranges. Usage: ` + "`" + `` + "`" + `` + "`" + `hcl network_interface = [{ network_interface_id = "" device_index = 1 description = "nic description in here" private_ip_address = "1.1.1.1" delete_on_termination = false secondary_private_ip_address_count = 1 associate_public_ip_address = true }] ` + "`" + `` + "`" + `` + "`" + ` <a id="block-devices"></a> ## Block Devices Each of the ` + "`" + ``,
				},
				resource.Attribute{
					Name:        "device_name",
					Description: `(Required) The name of the device to mount.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `(Optional) The Snapshot ID to mount.`,
				},
				resource.Attribute{
					Name:        "volume_type",
					Description: `(Optional, Default: ` + "`" + `"standard"` + "`" + `) The type of volume. Can be ` + "`" + `"standard"` + "`" + `, ` + "`" + `"gp2"` + "`" + `, ` + "`" + `"io1"` + "`" + `, ` + "`" + `"st1"` + "`" + ` or ` + "`" + `"sc1"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "volume_size",
					Description: `(Optional) The size of the volume in gigabytes.`,
				},
				resource.Attribute{
					Name:        "iops",
					Description: `(Optional) The amount of provisioned [IOPS](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ebs-io-characteristics.html). This must be set with a ` + "`" + `volume_type` + "`" + ` of ` + "`" + `"io1"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "delete_on_termination",
					Description: `(Optional) Whether the volume should be destroyed on instance termination.`,
				},
				resource.Attribute{
					Name:        "encrypted",
					Description: `(Optional) Enables [EBS encryption](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSEncryption.html) on the volume.`,
				},
				resource.Attribute{
					Name:        "kms_key_id",
					Description: `(Optional) ID for a user managed CMK under which the EBS Volume is encrypted Modifying any ` + "`" + `ebs_block_device` + "`" + ` currently requires resource replacement. Usage: ` + "`" + `` + "`" + `` + "`" + `hcl ebs_block_device = [{ device_name = "/dev/sdb" snapshot_id = "" volume_type = "gp2" volume_size = 8 iops = 1 delete_on_termination = true encrypted = false kms_key_id = "kms-key-01" }, { device_name = "/dev/sdc" snapshot_id = "" volume_type = "gp2" volume_size = 8 iops = 1 delete_on_termination = true encrypted = true kms_key_id = "kms-key-02" }] ` + "`" + `` + "`" + `` + "`" + ` Each ` + "`" + `ephemeral_block_device` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "device_name",
					Description: `(Required) The name of the block device to mount on the instance.`,
				},
				resource.Attribute{
					Name:        "virtual_name",
					Description: `(Required) The [Instance Store Device Name](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/InstanceStorage.html#InstanceStoreDeviceNames) (e.g. ` + "`" + `"ephemeral0"` + "`" + `). Usage: ` + "`" + `` + "`" + `` + "`" + `hcl ephemeral_block_device = [{ device_name = "/dev/xvdc" virtual_name = "ephemeral0" }] ` + "`" + `` + "`" + `` + "`" + ` <a id="stateful"></a> ## Stateful We support instance persistence via the following configurations. all values are boolean. For more information on instance persistence please see: [Stateful configuration](https://api.spotinst.com/integration-docs/elastigroup/concepts/stateful-concepts/introduction/)`,
				},
				resource.Attribute{
					Name:        "persist_root_device",
					Description: `(Optional) Boolean, should the instance maintain its root device volumes.`,
				},
				resource.Attribute{
					Name:        "persist_block_devices",
					Description: `(Optional) Boolean, should the instance maintain its Data volumes.`,
				},
				resource.Attribute{
					Name:        "persist_private_ip",
					Description: `(Optional) Boolean, should the instance maintain its private IP.`,
				},
				resource.Attribute{
					Name:        "block_devices_mode",
					Description: `(Optional) String, determine the way we attach the data volumes to the data devices, possible values: ` + "`" + `"reattach"` + "`" + ` and ` + "`" + `"onLaunch"` + "`" + ` (default is onLaunch).`,
				},
				resource.Attribute{
					Name:        "private_ips",
					Description: `(Optional) List of Private IPs to associate to the group instances.(e.g. "172.1.1.0"). Please note: This setting will only apply if persistence.persist_private_ip is set to true. Usage: ` + "`" + `` + "`" + `` + "`" + `hcl persist_root_device = false persist_block_devices = false persist_private_ip = true block_devices_mode = "onLaunch" private_ips = ["1.1.1.1", "2.2.2.2"] ` + "`" + `` + "`" + `` + "`" + ` <a id="stateful-deallocation"></a> ## Stateful Deallocation`,
				},
				resource.Attribute{
					Name:        "stateful_deallocation",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "should_delete_images",
					Description: `(Optional) For stateful groups: remove persistent images.`,
				},
				resource.Attribute{
					Name:        "should_delete_network_interfaces",
					Description: `(Optional) For stateful groups: remove network interfaces.`,
				},
				resource.Attribute{
					Name:        "should_delete_volumes",
					Description: `(Optional) For stateful groups: remove persistent volumes.`,
				},
				resource.Attribute{
					Name:        "should_delete_snapshots",
					Description: `(Optional) For stateful groups: remove snapshots. Usage: ` + "`" + `` + "`" + `` + "`" + `hcl stateful_deallocation = { should_delete_images = false should_delete_network_interfaces = false should_delete_volumes = false should_delete_snapshots = false } ` + "`" + `` + "`" + `` + "`" + ` <a id="health-check"></a> ## Health Check`,
				},
				resource.Attribute{
					Name:        "health_check_type",
					Description: `(Optional) The service that will perform health checks for the instance. Supported values : ` + "`" + `"ELB"` + "`" + `, ` + "`" + `"HCS"` + "`" + `, ` + "`" + `"TARGET_GROUP"` + "`" + `, ` + "`" + `"CUSTOM"` + "`" + `, ` + "`" + `"K8S_NODE"` + "`" + `, ` + "`" + `"MLB"` + "`" + `, ` + "`" + `"EC2"` + "`" + `, ` + "`" + `"MULTAI_TARGET_SET"` + "`" + `, ` + "`" + `"MLB_RUNTIME"` + "`" + `, ` + "`" + `"K8S_NODE"` + "`" + `, ` + "`" + `"NOMAD_NODE"` + "`" + `, ` + "`" + `"ECS_CLUSTER_INSTANCE"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "health_check_grace_period",
					Description: `(Optional) The amount of time, in seconds, after the instance has launched to starts and check its health`,
				},
				resource.Attribute{
					Name:        "health_check_unhealthy_duration_before_replacement",
					Description: `(Optional) The amount of time, in seconds, that we will wait before replacing an instance that is running and became unhealthy (this is only applicable for instances that were once healthy) Usage: ` + "`" + `` + "`" + `` + "`" + `hcl health_check_type = "ELB" health_check_grace_period = 100 health_check_unhealthy_duration_before_replacement = 120 ` + "`" + `` + "`" + `` + "`" + ` <a id="third-party-integrations"></a> ## Third-Party Integrations`,
				},
				resource.Attribute{
					Name:        "integration_rancher",
					Description: `(Optional) Describes the [Rancher](http://rancherlabs.com/) integration.`,
				},
				resource.Attribute{
					Name:        "master_host",
					Description: `(Required) The URL of the Rancher Master host.`,
				},
				resource.Attribute{
					Name:        "access_key",
					Description: `(Required) The access key of the Rancher API.`,
				},
				resource.Attribute{
					Name:        "secret_key",
					Description: `(Required) The secret key of the Rancher API.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Optional) The Rancher version. Must be ` + "`" + `"1"` + "`" + ` or ` + "`" + `"2"` + "`" + `. If this field is omitted, it’s assumed that the Rancher cluster is version 1. Note that Kubernetes is required when using Rancher version 2^. Usage: ` + "`" + `` + "`" + `` + "`" + `hcl integration_rancher = { master_host = "master_host" access_key = "access_key" secret_key = "secret_key" version = "2" } ` + "`" + `` + "`" + `` + "`" + ``,
				},
				resource.Attribute{
					Name:        "integration_ecs",
					Description: `(Optional) Describes the [EC2 Container Service](https://aws.amazon.com/documentation/ecs/?id=docs_gateway) integration.`,
				},
				resource.Attribute{
					Name:        "cluster_name",
					Description: `(Required) The name of the EC2 Container Service cluster.`,
				},
				resource.Attribute{
					Name:        "autoscale_is_enabled",
					Description: `(Optional, Default: ` + "`" + `false` + "`" + `) Specifies whether the auto scaling feature is enabled.`,
				},
				resource.Attribute{
					Name:        "autoscale_cooldown",
					Description: `(Optional, Default: ` + "`" + `300` + "`" + `) The amount of time, in seconds, after a scaling activity completes before any further trigger-related scaling activities can start.`,
				},
				resource.Attribute{
					Name:        "autoscale_is_auto_config",
					Description: `(Optional, Default: ` + "`" + `false` + "`" + `) Enabling the automatic auto-scaler functionality. For more information please see: [ECS auto scaler](https://api.spotinst.com/container-management/amazon-ecs/elastigroup-for-ecs-concepts/autoscaling/).`,
				},
				resource.Attribute{
					Name:        "autoscale_scale_down_non_service_tasks",
					Description: `(Optional) Determines whether to scale down non-service tasks.`,
				},
				resource.Attribute{
					Name:        "autoscale_headroom",
					Description: `(Optional) Headroom for the cluster.`,
				},
				resource.Attribute{
					Name:        "cpu_per_unit",
					Description: `(Optional, Default: ` + "`" + `0` + "`" + `) Cpu units for compute.`,
				},
				resource.Attribute{
					Name:        "memory_per_unit",
					Description: `(Optional, Default: ` + "`" + `0` + "`" + `) RAM units for compute.`,
				},
				resource.Attribute{
					Name:        "num_of_units",
					Description: `(Optional, Default: ` + "`" + `0` + "`" + `) Amount of units for compute.`,
				},
				resource.Attribute{
					Name:        "autoscale_down",
					Description: `(Optional) Enabling scale down.`,
				},
				resource.Attribute{
					Name:        "evaluation_periods",
					Description: `(Optional, Default: ` + "`" + `5` + "`" + `) Amount of cooldown evaluation periods for scale down.`,
				},
				resource.Attribute{
					Name:        "max_scale_down_percentage",
					Description: `(Optional) Represents the maximum percent to scale-down. Number between 1-100.`,
				},
				resource.Attribute{
					Name:        "autoscale_attributes",
					Description: `(Optional) A key/value mapping of tags to assign to the resource. Usage: ` + "`" + `` + "`" + `` + "`" + `hcl integration_ecs = { cluster_name = "ecs-cluster" autoscale_is_enabled = false autoscale_cooldown = 300 autoscale_scale_down_non_service_tasks = false autoscale_headroom = { cpu_per_unit = 1024 memory_per_unit = 512 num_of_units = 2 } autoscale_down = { evaluation_periods = 300 max_scale_down_percentage = 70 } autoscale_attributes = [{ key = "test.ecs.key" value = "test.ecs.value" }] } ` + "`" + `` + "`" + `` + "`" + ``,
				},
				resource.Attribute{
					Name:        "integration_codedeploy",
					Description: `(Optional) Describes the [Code Deploy](https://aws.amazon.com/documentation/codedeploy/?id=docs_gateway) integration.`,
				},
				resource.Attribute{
					Name:        "cleanup_on_failure",
					Description: `(Optional) Cleanup automatically after a failed deploy.`,
				},
				resource.Attribute{
					Name:        "terminate_instance_on_failure",
					Description: `(Optional) Terminate the instance automatically after a failed deploy.`,
				},
				resource.Attribute{
					Name:        "deployment_groups",
					Description: `(Optional) Specify the deployment groups details.`,
				},
				resource.Attribute{
					Name:        "application_name",
					Description: `(Optional) The application name.`,
				},
				resource.Attribute{
					Name:        "deployment_group_name",
					Description: `(Optional) The deployment group name. Usage: ` + "`" + `` + "`" + `` + "`" + `hcl integration_codedeploy = { cleanup_on_failure = false terminate_instance_on_failure = false deployment_groups = { application_name = "my-app" deployment_group_name = "my-group" } } ` + "`" + `` + "`" + `` + "`" + ``,
				},
				resource.Attribute{
					Name:        "integration_route53",
					Description: `(Optional) Describes the [Route53](https://aws.amazon.com/documentation/route53/?id=docs_gateway) integration.`,
				},
				resource.Attribute{
					Name:        "domains",
					Description: `(Required) Collection of one or more domains to register.`,
				},
				resource.Attribute{
					Name:        "hosted_zone_id",
					Description: `(Required) The id associated with a hosted zone.`,
				},
				resource.Attribute{
					Name:        "spotinst_acct_id",
					Description: `(Optional) The Spotinst account ID that is linked to the AWS account that holds the Route 53 hosted Zone ID. The default is the user Spotinst account provided as a URL parameter.`,
				},
				resource.Attribute{
					Name:        "record_sets",
					Description: `(Required) Collection of records containing authoritative DNS information for the specified domain name.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The record set name.`,
				},
				resource.Attribute{
					Name:        "use_public_ip",
					Description: `(Optional, Default: ` + "`" + `false` + "`" + `) - Designates if the IP address should be exposed to connections outside the VPC. Usage: ` + "`" + `` + "`" + `` + "`" + `hcl integration_route53 = { domains = { hosted_zone_id = "zone-id" spotinst_acct_id = "act-123456" record_sets = { name = "foo.example.com" use_public_ip = true } } } ` + "`" + `` + "`" + `` + "`" + ``,
				},
				resource.Attribute{
					Name:        "integration_docker_swarm",
					Description: `(Optional) Describes the [Docker Swarm](https://api.spotinst.com/integration-docs/elastigroup/container-management/docker-swarm/docker-swarm-integration/) integration.`,
				},
				resource.Attribute{
					Name:        "master_host",
					Description: `(Required) IP or FQDN of one of your swarm managers.`,
				},
				resource.Attribute{
					Name:        "master_port",
					Description: `(Required) Network port used by your swarm.`,
				},
				resource.Attribute{
					Name:        "autoscale_is_enabled",
					Description: `(Optional, Default: ` + "`" + `false` + "`" + `) Specifies whether the auto scaling feature is enabled.`,
				},
				resource.Attribute{
					Name:        "autoscale_cooldown",
					Description: `(Optional, Default: ` + "`" + `300` + "`" + `) The amount of time, in seconds, after a scaling activity completes before any further trigger-related scaling activities can start. Minimum 180, must be a multiple of 60.`,
				},
				resource.Attribute{
					Name:        "autoscale_headroom",
					Description: `(Optional) An option to set compute reserve for the cluster.`,
				},
				resource.Attribute{
					Name:        "cpu_per_unit",
					Description: `(Optional, Default: ` + "`" + `0` + "`" + `) How much CPU to allocate for headroom unit.`,
				},
				resource.Attribute{
					Name:        "memory_per_unit",
					Description: `(Optional, Default: ` + "`" + `0` + "`" + `) The amount of memory in each headroom unit. Measured in MiB.`,
				},
				resource.Attribute{
					Name:        "num_of_units",
					Description: `(Optional, Default: ` + "`" + `0` + "`" + `) How many units to allocate for headroom unit.`,
				},
				resource.Attribute{
					Name:        "autoscale_down",
					Description: `(Optional) Setting for scale down actions.`,
				},
				resource.Attribute{
					Name:        "evaluation_periods",
					Description: `(Optional, Default: ` + "`" + `5` + "`" + `) Number of periods over which data is compared. Minimum 3, Measured in consecutive minutes. Usage: ` + "`" + `` + "`" + `` + "`" + `hcl integration_docker_swarm = { master_host = "10.10.10.10" master_port = 2376 autoscale_is_enabled = true autoscale_cooldown = 180 autoscale_headroom = { cpu_per_unit = 2048 memory_per_unit = 2048 num_of_units = 1 } autoscale_down = { evaluation_periods = 3 } } ` + "`" + `` + "`" + `` + "`" + ``,
				},
				resource.Attribute{
					Name:        "integration_kubernetes",
					Description: `(Optional) Describes the [Kubernetes](https://kubernetes.io/) integration.`,
				},
				resource.Attribute{
					Name:        "integration_mode",
					Description: `(Required) Valid values: ` + "`" + `"saas"` + "`" + `, ` + "`" + `"pod"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "cluster_identifier",
					Description: `(Required; if using integration_mode as pod)`,
				},
				resource.Attribute{
					Name:        "api_server",
					Description: `(Required; if using integration_mode as saas)`,
				},
				resource.Attribute{
					Name:        "token",
					Description: `(Required; if using integration_mode as saas) Kubernetes Token`,
				},
				resource.Attribute{
					Name:        "autoscale_is_enabled",
					Description: `(Optional, Default: ` + "`" + `false` + "`" + `) Specifies whether the auto scaling feature is enabled.`,
				},
				resource.Attribute{
					Name:        "autoscale_is_auto_config",
					Description: `(Optional, Default: ` + "`" + `false` + "`" + `) Enabling the automatic k8s auto-scaler functionality. For more information please see: [Kubernetes auto scaler](https://api.spotinst.com/integration-docs/elastigroup/container-management/kubernetes/autoscaler/).`,
				},
				resource.Attribute{
					Name:        "autoscale_cooldown",
					Description: `(Optional, Default: ` + "`" + `300` + "`" + `) The amount of time, in seconds, after a scaling activity completes before any further trigger-related scaling activities can start.`,
				},
				resource.Attribute{
					Name:        "autoscale_headroom",
					Description: `(Optional) An option to set compute reserve for the cluster.`,
				},
				resource.Attribute{
					Name:        "cpu_per_unit",
					Description: `(Optional, Default: ` + "`" + `0` + "`" + `) How much CPU to allocate for headroom unit.`,
				},
				resource.Attribute{
					Name:        "memory_per_unit",
					Description: `(Optional, Default: ` + "`" + `0` + "`" + `) How much Memory allocate for headroom unit.`,
				},
				resource.Attribute{
					Name:        "num_of_units",
					Description: `(Optional, Default: ` + "`" + `0` + "`" + `) How many units to allocate for headroom unit.`,
				},
				resource.Attribute{
					Name:        "autoscale_down",
					Description: `(Optional) Setting for scale down actions.`,
				},
				resource.Attribute{
					Name:        "evaluation_periods",
					Description: `(Optional, Default: ` + "`" + `5` + "`" + `) How many evaluation periods should accumulate before a scale down action takes place.`,
				},
				resource.Attribute{
					Name:        "autoscale_labels",
					Description: `(Optional) A key/value mapping of tags to assign to the resource. Usage: ` + "`" + `` + "`" + `` + "`" + `hcl integration_kubernetes = { integration_mode = "pod" cluster_identifier = "my-identifier.ek8s.com" // === SAAS =================== # integration_mode = "saas" # api_server = "https://api.my-identifier.ek8s.com/api/v1/namespaces/kube-system/services/..." # token = "top-secret" // ============================ autoscale_is_enabled = false autoscale_is_auto_config = false autoscale_cooldown = 300 autoscale_headroom = { cpu_per_unit = 1024 memory_per_unit = 512 num_of_units = 1 } autoscale_down = { evaluation_periods = 300 } autoscale_labels = [{ key = "test.k8s.key" value = "test.k8s.value" }] } ` + "`" + `` + "`" + `` + "`" + ``,
				},
				resource.Attribute{
					Name:        "integration_nomad",
					Description: `(Optional) Describes the [Nomad](https://www.nomadproject.io/) integration.`,
				},
				resource.Attribute{
					Name:        "master_host",
					Description: `(Required) The URL for the Nomad master host.`,
				},
				resource.Attribute{
					Name:        "master_port",
					Description: `(Required) The network port for the master host.`,
				},
				resource.Attribute{
					Name:        "acl_token",
					Description: `(Required) Nomad ACL Token`,
				},
				resource.Attribute{
					Name:        "autoscale_is_enabled",
					Description: `(Optional, Default: ` + "`" + `false` + "`" + `) Specifies whether the auto scaling feature is enabled.`,
				},
				resource.Attribute{
					Name:        "autoscale_cooldown",
					Description: `(Optional, Default: ` + "`" + `300` + "`" + `) The amount of time, in seconds, after a scaling activity completes before any further trigger-related scaling activities can start.`,
				},
				resource.Attribute{
					Name:        "autoscale_headroom",
					Description: `(Optional) An option to set compute reserve for the cluster.`,
				},
				resource.Attribute{
					Name:        "cpu_per_unit",
					Description: `(Optional, Default: ` + "`" + `0` + "`" + `) How much CPU (MHz) to allocate for headroom unit.`,
				},
				resource.Attribute{
					Name:        "memory_per_unit",
					Description: `(Optional, Default: ` + "`" + `0` + "`" + `) How much Memory allocate for headroom unit.`,
				},
				resource.Attribute{
					Name:        "num_of_units",
					Description: `(Optional, Default: ` + "`" + `0` + "`" + `) How many units of headroom to allocate.`,
				},
				resource.Attribute{
					Name:        "autoscale_down",
					Description: `(Optional) Settings for scale down actions.`,
				},
				resource.Attribute{
					Name:        "evaluation_periods",
					Description: `(Optional, Default: ` + "`" + `5` + "`" + `) How many evaluation periods should accumulate before a scale down action takes place.`,
				},
				resource.Attribute{
					Name:        "autoscale_constraints",
					Description: `(Optional) A key/value mapping of tags to assign to the resource. Usage: ` + "`" + `` + "`" + `` + "`" + `hcl integration_nomad = { master_host = "my-nomad-host" master_port = 9000 acl_token = "top-secret" autoscale_is_enabled = false autoscale_cooldown = 300 autoscale_headroom = { cpu_per_unit = 1024 memory_per_unit = 512 num_of_units = 2 } autoscale_down = { evaluation_periods = 300 } autoscale_constraints = [{ key = "test.nomad.key" value = "test.nomad.value" }] } ` + "`" + `` + "`" + `` + "`" + ``,
				},
				resource.Attribute{
					Name:        "integration_mesosphere",
					Description: `(Optional) Describes the [Mesosphere](https://mesosphere.com/) integration.`,
				},
				resource.Attribute{
					Name:        "api_server",
					Description: `(Optional) The public IP of the DC/OS Master. Usage: ` + "`" + `` + "`" + `` + "`" + `hcl integration_mesosphere = { api_server = "" } ` + "`" + `` + "`" + `` + "`" + ``,
				},
				resource.Attribute{
					Name:        "integration_multai_runtime",
					Description: `(Optional) Describes the [Multai Runtime](https://spotinst.com/) integration.`,
				},
				resource.Attribute{
					Name:        "deployment_id",
					Description: `(Optional) The deployment id you want to get Usage: ` + "`" + `` + "`" + `` + "`" + `hcl integration_multai_runtime = { deployment_id = "" } ` + "`" + `` + "`" + `` + "`" + ``,
				},
				resource.Attribute{
					Name:        "integration_gitlab",
					Description: `(Optional) Describes the [Gitlab](https://api.spotinst.com/integration-docs/gitlab/) integration.`,
				},
				resource.Attribute{
					Name:        "runner",
					Description: `(Optional) Settings for Gitlab runner.`,
				},
				resource.Attribute{
					Name:        "is_enabled",
					Description: `(Optional, Default: ` + "`" + `false` + "`" + `) Specifies whether the integration is enabled. Usage: ` + "`" + `` + "`" + `` + "`" + `hcl integration_gitlab = { runner = { is_enabled = true } } ` + "`" + `` + "`" + `` + "`" + ``,
				},
				resource.Attribute{
					Name:        "integration_beanstalk",
					Description: `(Optional) Describes the [Beanstalk](https://api.spotinst.com/provisioning-ci-cd-sdk/provisioning-tools/terraform/resources/terraform-v-2/elastic-beanstalk/) integration.`,
				},
				resource.Attribute{
					Name:        "deployment_preferences",
					Description: `(Optional) Preferences when performing a roll`,
				},
				resource.Attribute{
					Name:        "automatic_roll",
					Description: `(Required) Should roll perform automatically`,
				},
				resource.Attribute{
					Name:        "batch_size_percentage",
					Description: `(Required) Percent size of each batch`,
				},
				resource.Attribute{
					Name:        "grace_period",
					Description: `(Required) Amount of time to wait between batches`,
				},
				resource.Attribute{
					Name:        "strategy",
					Description: `(Optional) Strategy parameters`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Required) Action to take`,
				},
				resource.Attribute{
					Name:        "should_drain_instances",
					Description: `(Required) Bool value if to wait to drain instance`,
				},
				resource.Attribute{
					Name:        "managed_actions",
					Description: `(Optional) Managed Actions parameters`,
				},
				resource.Attribute{
					Name:        "platform_update",
					Description: `(Optional) Platform Update parameters`,
				},
				resource.Attribute{
					Name:        "perform_at",
					Description: `(Required) Actions to perform (options: timeWindow, never)`,
				},
				resource.Attribute{
					Name:        "time_window",
					Description: `(Required) Time Window for when action occurs ex. Mon:23:50-Tue:00:20`,
				},
				resource.Attribute{
					Name:        "update_level",
					Description: `(Required) - Level to update Usage: ` + "`" + `` + "`" + `` + "`" + `hcl integration_beanstalk = { environment_id = "e-3tkmbj7hzc" deployment_preferences = { automatic_roll = true batch_size_percentage = 100 grace_period = 90 strategy = { action = "REPLACE_SERVER" should_drain_instance = true } } managed_actions = { platform_update = { perform_at = "timeWindow" field_name = "Mon:23:50-Tue:00:20" update_level = "minorAndPatch" } } } ` + "`" + `` + "`" + `` + "`" + ` <a id="update-policy"></a> ## Update Policy`,
				},
				resource.Attribute{
					Name:        "update_policy",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "should_resume_stateful",
					Description: `(Required) This will apply resuming action for Stateful instances in the Elastigroup upon scale up or capacity changes. Example usage will be for Elastigroups that will have scheduling rules to set a target capacity of 0 instances in the night and automatically restore the same state of the instances in the morning.`,
				},
				resource.Attribute{
					Name:        "auto_apply_tags",
					Description: `(Optional) Enables updates to tags without rolling the group when set to ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "should_roll",
					Description: `(Required) Sets the enablement of the roll option.`,
				},
				resource.Attribute{
					Name:        "roll_config",
					Description: `(Required) While used, you can control whether the group should perform a deployment after an update to the configuration.`,
				},
				resource.Attribute{
					Name:        "batch_size_percentage",
					Description: `(Required) Sets the percentage of the instances to deploy in each batch.`,
				},
				resource.Attribute{
					Name:        "health_check_type",
					Description: `(Optional) Sets the health check type to use. Valid values: ` + "`" + `"EC2"` + "`" + `, ` + "`" + `"ECS_CLUSTER_INSTANCE"` + "`" + `, ` + "`" + `"ELB"` + "`" + `, ` + "`" + `"HCS"` + "`" + `, ` + "`" + `"MLB"` + "`" + `, ` + "`" + `"TARGET_GROUP"` + "`" + `, ` + "`" + `"MULTAI_TARGET_SET"` + "`" + `, ` + "`" + `"NONE"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "grace_period",
					Description: `(Optional) Sets the grace period for new instances to become healthy.`,
				},
				resource.Attribute{
					Name:        "wait_for_roll_percentage",
					Description: `(Optional) For use with ` + "`" + `should_roll` + "`" + `. Sets minimum % of roll required to complete before continuing the plan. Required if ` + "`" + `wait_for_roll_timeout` + "`" + ` is set.`,
				},
				resource.Attribute{
					Name:        "wait_for_roll_timeout",
					Description: `(Optional) For use with ` + "`" + `should_roll` + "`" + `. Sets how long to wait for the deployed % of a roll to exceed ` + "`" + `wait_for_roll_percentage` + "`" + ` before continuing the plan. Required if ` + "`" + `wait_for_roll_percentage` + "`" + ` is set.`,
				},
				resource.Attribute{
					Name:        "strategy",
					Description: `(Optional) Strategy parameters`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Required) Action to take. Valid values: ` + "`" + `REPLACE_SERVER` + "`" + `, ` + "`" + `RESTART_SERVER` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "should_drain_instances",
					Description: `(Optional) Specify whether to drain incoming TCP connections before terminating a server.`,
				},
				resource.Attribute{
					Name:        "batch_min_healthy_percentage",
					Description: `(Optional, Default ` + "`" + `50` + "`" + `) Indicates the threshold of minimum healthy instances in single batch. If the amount of healthy instances in single batch is under the threshold, the deployment will fail. Range ` + "`" + `1` + "`" + ` - ` + "`" + `100` + "`" + `. ` + "`" + `` + "`" + `` + "`" + `hcl update_policy = { should_resume_stateful = false should_roll = false auto_apply_tags = false roll_config = { batch_size_percentage = 33 health_check_type = "ELB" grace_period = 300 wait_for_roll_percentage = 10 wait_for_roll_timeout = 1500 strategy = { action = "REPLACE_SERVER" should_drain_instances = false batch_min_healthy_percentage = 10 } } } ` + "`" + `` + "`" + `` + "`" + ` ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The group ID.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The group ID.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "elastigroup_aws_beanstalk",
			Category:         "Spotinst Resources",
			ShortDescription: `Provides a Spotinst AWS group resource using Elastic Beanstalk.`,
			Description:      ``,
			Keywords: []string{
				"elastigroup",
				"aws",
				"beanstalk",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The group name.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Required) The AWS region your group will be created in. Cannot be changed after the group has been created.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The group description.`,
				},
				resource.Attribute{
					Name:        "product",
					Description: `(Required) Operation system type. Valid values: ` + "`" + `"Linux/UNIX"` + "`" + `, ` + "`" + `"SUSE Linux"` + "`" + `, ` + "`" + `"Windows"` + "`" + `. For EC2 Classic instances: ` + "`" + `"Linux/UNIX (Amazon VPC)"` + "`" + `, ` + "`" + `"SUSE Linux (Amazon VPC)"` + "`" + `, ` + "`" + `"Windows (Amazon VPC)"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "max_size",
					Description: `(Required) The maximum number of instances the group should have at any time.`,
				},
				resource.Attribute{
					Name:        "min_size",
					Description: `(Required) The minimum number of instances the group should have at any time.`,
				},
				resource.Attribute{
					Name:        "desired_capacity",
					Description: `(Required) The desired number of instances the group should have at any time.`,
				},
				resource.Attribute{
					Name:        "beanstalk_environment_name",
					Description: `(Optional) The name of an existing Beanstalk environment.`,
				},
				resource.Attribute{
					Name:        "beanstalk_environment_id",
					Description: `(Optional) The id of an existing Beanstalk environment.`,
				},
				resource.Attribute{
					Name:        "instance_types_spot",
					Description: `(Required) One or more instance types. To maximize the availability of Spot instances, select as many instance types as possible.`,
				},
				resource.Attribute{
					Name:        "deployment_preferences",
					Description: `(Optional) Preferences when performing a roll`,
				},
				resource.Attribute{
					Name:        "automatic_roll",
					Description: `(Required) Should roll perform automatically`,
				},
				resource.Attribute{
					Name:        "batch_size_percentage",
					Description: `(Required) Percent size of each batch`,
				},
				resource.Attribute{
					Name:        "grace_period",
					Description: `(Required) Amount of time to wait between batches`,
				},
				resource.Attribute{
					Name:        "strategy",
					Description: `(Optional) Strategy parameters`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Required) Action to take`,
				},
				resource.Attribute{
					Name:        "should_drain_instances",
					Description: `(Required) Bool value if to wait to drain instance`,
				},
				resource.Attribute{
					Name:        "managed_actions",
					Description: `(Optional) Managed Actions parameters`,
				},
				resource.Attribute{
					Name:        "platform_update",
					Description: `(Optional) Platform Update parameters`,
				},
				resource.Attribute{
					Name:        "perform_at",
					Description: `(Required) Actions to perform (options: timeWindow, never)`,
				},
				resource.Attribute{
					Name:        "time_window",
					Description: `(Required) Time Window for when action occurs ex. Mon:23:50-Tue:00:20`,
				},
				resource.Attribute{
					Name:        "update_level",
					Description: `(Required) - Level to update`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "elastigroup_azure",
			Category:         "Spotinst Resources",
			ShortDescription: `Provides a Spotinst elastigroup resource for Microsoft Azure.`,
			Description:      ``,
			Keywords: []string{
				"elastigroup",
				"azure",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The group name.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Required) The region your Azure group will be created in.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `(Required) Name of the Resource Group for Elastigroup.`,
				},
				resource.Attribute{
					Name:        "product",
					Description: `(Required) Operation system type. Valid values: ` + "`" + `"Linux"` + "`" + `, ` + "`" + `"Windows"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "max_size",
					Description: `(Required) The maximum number of instances the group should have at any time.`,
				},
				resource.Attribute{
					Name:        "min_size",
					Description: `(Required) The minimum number of instances the group should have at any time.`,
				},
				resource.Attribute{
					Name:        "desired_capacity",
					Description: `(Required) The desired number of instances the group should have at any time.`,
				},
				resource.Attribute{
					Name:        "od_sizes",
					Description: `(Required) Available On-Demand sizes`,
				},
				resource.Attribute{
					Name:        "low_priority_sizes",
					Description: `(Required) Available Low-Priority sizes.`,
				},
				resource.Attribute{
					Name:        "user_data",
					Description: `(Optional) Base64-encoded MIME user data to make available to the instances.`,
				},
				resource.Attribute{
					Name:        "shutdown_script",
					Description: `(Optional) Shutdown script for the group. Value should be passed as a string encoded at Base64 only.`,
				},
				resource.Attribute{
					Name:        "managed_service_identity",
					Description: `(Optional) Add a user-assigned managed identity to the VMs in the cluster.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `(Required) The Resource Group that the user-assigned managed identity resides in.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the managed identity.`,
				},
				resource.Attribute{
					Name:        "strategy",
					Description: `(Required) Describes the deployment strategy.`,
				},
				resource.Attribute{
					Name:        "low_priority_percentage",
					Description: `(Optional, Default ` + "`" + `100` + "`" + `) Percentage of Low Priority instances to maintain. Required if ` + "`" + `od_count` + "`" + ` is not specified.`,
				},
				resource.Attribute{
					Name:        "od_count",
					Description: `(Optional) Number of On-Demand instances to maintain. Required if low_priority_percentage is not specified.`,
				},
				resource.Attribute{
					Name:        "draining_timeout",
					Description: `(Optional, Default ` + "`" + `120` + "`" + `) Time (seconds) to allow the instance to be drained from incoming TCP connections and detached from MLB before terminating it during a scale-down operation. <a id="load-balancers"></a> ## Load Balancers`,
				},
				resource.Attribute{
					Name:        "load_balancers",
					Description: `(Required) Describes a set of one or more classic load balancer target groups and/or Multai load balancer target sets.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) The resource type. Valid values: CLASSIC, TARGET_GROUP, MULTAI_TARGET_SET.`,
				},
				resource.Attribute{
					Name:        "balancer_id",
					Description: `(Required) The balancer ID.`,
				},
				resource.Attribute{
					Name:        "target_set_id",
					Description: `(Required) The scale set ID associated with the load balancer.`,
				},
				resource.Attribute{
					Name:        "auto_weight",
					Description: `(Optional, Default: ` + "`" + `false` + "`" + `) ` + "`" + `` + "`" + `` + "`" + `hcl load_balancers = [{ type = "MULTAI_TARGET_SET" balancer_id = "lb-1ee2e3q" target_set_id = "ts-3eq" auto_weight = true }] ` + "`" + `` + "`" + `` + "`" + ` <a id="image"></a> ## Image`,
				},
				resource.Attribute{
					Name:        "image",
					Description: `(Required) Image of a VM. An image is a template for creating new VMs. Choose from Azure image catalogue (marketplace) or use a custom image.`,
				},
				resource.Attribute{
					Name:        "publisher",
					Description: `(Optional) Image publisher. Required if resource_group_name is not specified.`,
				},
				resource.Attribute{
					Name:        "offer",
					Description: `(Optional) Name of the image to use. Required if publisher is specified.`,
				},
				resource.Attribute{
					Name:        "sku",
					Description: `(Optional) Image's Stock Keeping Unit, which is the specific version of the image. Required if publisher is specified.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `(Optional) Name of Resource Group for custom image. Required if publisher not specified.`,
				},
				resource.Attribute{
					Name:        "image_name",
					Description: `(Optional) Name of the custom image. Required if resource_group_name is specified. ` + "`" + `` + "`" + `` + "`" + `hcl // market image image = { marketplace = { publisher = "Canonical" offer = "UbuntuServer" sku = "16.04-LTS" } } // custom image image = { custom = { image_name = "customImage" resource_group_name = "resourceGroup" } } ` + "`" + `` + "`" + `` + "`" + ` <a id="health-check"></a> ## Health Check`,
				},
				resource.Attribute{
					Name:        "health_check",
					Description: `(Optional) Describes the health check configuration.`,
				},
				resource.Attribute{
					Name:        "health_check_type",
					Description: `(Optional) Health check used to validate VM health. Valid values: “INSTANCE_STATE”.`,
				},
				resource.Attribute{
					Name:        "grace_period",
					Description: `(Optional) Period of time (seconds) to wait for VM to reach healthiness before monitoring for unhealthiness.`,
				},
				resource.Attribute{
					Name:        "auto_healing",
					Description: `(Optional) Enable auto-healing of unhealthy VMs. ` + "`" + `` + "`" + `` + "`" + `hcl health_check = { health_check_type = "INSTANCE_STATE" grace_period = 120 auto_healing = true } ` + "`" + `` + "`" + `` + "`" + ` <a id="network"></a> ## Network`,
				},
				resource.Attribute{
					Name:        "network",
					Description: `(Required) Defines the Virtual Network and Subnet for your Elastigroup.`,
				},
				resource.Attribute{
					Name:        "virtual_network_name",
					Description: `(Required) Name of Vnet.`,
				},
				resource.Attribute{
					Name:        "subnet_name",
					Description: `(Required) ID of subnet.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `(Required) Vnet Resource Group Name.`,
				},
				resource.Attribute{
					Name:        "assign_public_up",
					Description: `(Optional, Default: ` + "`" + `false` + "`" + `) Assign a public IP to each VM in the Elastigroup.`,
				},
				resource.Attribute{
					Name:        "additional_ip_configs",
					Description: `(Optional) Array of additional IP configuration objects.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The IP configuration name.`,
				},
				resource.Attribute{
					Name:        "private_ip_version",
					Description: `(Optional) Available from Azure Api-Version 2017-03-30 onwards, it represents whether the specific ipconfiguration is IPv4 or IPv6. Valid values: ` + "`" + `IPv4` + "`" + `, ` + "`" + `IPv6` + "`" + `. ` + "`" + `` + "`" + `` + "`" + `hcl network = { virtual_network_name = "vname" subnet_name = "my-subnet-name" resource_group_name = "subnetResourceGroup" assign_public_ip = true additional_ip_configs = [{ name = "test" private_ip_version = "IPv4" }] } ` + "`" + `` + "`" + `` + "`" + ` <a id="login"></a> ## Login ` + "`" + `` + "`" + `` + "`" + `hcl network = { virtual_network_name = "vname" subnet_name = "my-subnet-name" resource_group_name = "subnetResourceGroup" assign_public_ip = true } ` + "`" + `` + "`" + `` + "`" + ` <a id="login"></a> ## Login`,
				},
				resource.Attribute{
					Name:        "login",
					Description: `(Required) Describes the login configuration.`,
				},
				resource.Attribute{
					Name:        "user_name",
					Description: `(Required) Set admin access for accessing your VMs.`,
				},
				resource.Attribute{
					Name:        "ssh_public_key",
					Description: `(Optional) SSH for admin access to Linux VMs. Required for Linux product types.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional) Password for admin access to Windows VMs. Required for Windows product types. ` + "`" + `` + "`" + `` + "`" + `hcl login = { user_name = "admin" ssh_public_key = "33a2s1f3g5a1df5g1ad21651sag56dfg==" } ` + "`" + `` + "`" + `` + "`" + ` <a id="scaling-policy"></a> ## Scaling Policies Each ` + "`" + `scaling_`,
				},
				resource.Attribute{
					Name:        "policy_name",
					Description: `(Optional) The name of the policy.`,
				},
				resource.Attribute{
					Name:        "metric_name",
					Description: `(Required) Metric to monitor by Azure metric display name.`,
				},
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional, Default: ` + "`" + `“Microsoft.Compute”` + "`" + `) The namespace for the alarm's associated metric. Valid values: ` + "`" + `` + "`" + `` + "`" + `text Microsoft.AnalysisServices/servers Microsoft.ApiManagement/service Microsoft.Automation/automationAccounts Microsoft.Batch/batchAccounts Microsoft.Cache/redis Microsoft.CognitiveServices/accounts Microsoft.Compute Microsoft.ContainerInstance/containerGroups Microsoft.ContainerService/managedClusters Microsoft.CustomerInsights/hubs Microsoft.DataFactory/datafactories Microsoft.DataFactory/factories Microsoft.DataLakeAnalytics/accounts Microsoft.DataLakeStore/accounts Microsoft.DBforMariaDB/servers Microsoft.DBforMySQL/servers Microsoft.DBforPostgreSQL/servers Microsoft.Devices/IotHubs Microsoft.Devices/provisioningServices Microsoft.DocumentDB/databaseAccounts Microsoft.EventGrid/eventSubscriptions Microsoft.EventGrid/extensionTopics Microsoft.EventGrid/topics Microsoft.EventHub/clusters Microsoft.EventHub/namespaces Microsoft.HDInsight/clusters Microsoft.Insights/AutoscaleSettings Microsoft.Insights/Components Microsoft.KeyVault/vaults Microsoft.Kusto/Clusters Microsoft.LocationBasedServices/accounts Microsoft.Logic/workflows Microsoft.NetApp/netAppAccounts/capacityPools/Volumes Microsoft.NetApp/netAppAccounts/capacityPools Microsoft.Network/applicationGateways Microsoft.Network/dnszones Microsoft.Network/connections Microsoft.Network/expressRouteCircuits Microsoft.Network/expressRouteCircuits/peerings Microsoft.Network/frontdoors Microsoft.Network/loadBalancers Microsoft.Network/networkInterfaces Microsoft.Network/networkWatchers/connectionMonitors Microsoft.Network/publicIPAddresses Microsoft.Network/trafficManagerProfiles Microsoft.Network/virtualNetworkGateways Microsoft.NotificationHubs/Namespaces/NotificationHubs Microsoft.OperationalInsights/workspaces Microsoft.PowerBIDedicated/capacities Microsoft.Relay/namespaces Microsoft.Search/searchServices Microsoft.ServiceBus/namespaces Microsoft.SignalRService/SignalR Microsoft.Sql/managedInstances Microsoft.Sql/servers/databases Microsoft.Sql/servers/elasticPools Microsoft.Storage/storageAccounts Microsoft.Storage/storageAccounts/blobServices Microsoft.Storage/storageAccounts/fileServices Microsoft.Storage/storageAccounts/queueServices Microsoft.Storage/storageAccounts/tableServices Microsoft.StreamAnalytics/streamingjobs Microsoft.TimeSeriesInsights/environments Microsoft.TimeSeriesInsights/environments/eventsources Microsoft.Web/hostingEnvironments/multiRolePools Microsoft.Web/hostingEnvironments/workerPools Microsoft.Web/serverfarms Microsoft.Web/sites (excluding functions) Microsoft.Web/sites (functions) Microsoft.Web/sites/slots ` + "`" + `` + "`" + `` + "`" + ``,
				},
				resource.Attribute{
					Name:        "statistic",
					Description: `(Optional) The metric statistics to return. Valid values: ` + "`" + `average` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "threshold",
					Description: `(Required) The value against which the specified statistic is compared.`,
				},
				resource.Attribute{
					Name:        "unit",
					Description: `(Required) The unit for the alarm's associated metric. Valid values: ` + "`" + `"percent` + "`" + `, ` + "`" + `"seconds"` + "`" + `, ` + "`" + `"microseconds"` + "`" + `, ` + "`" + `"milliseconds"` + "`" + `, ` + "`" + `"bytes"` + "`" + `, ` + "`" + `"kilobytes"` + "`" + `, ` + "`" + `"megabytes"` + "`" + `, ` + "`" + `"gigabytes"` + "`" + `, ` + "`" + `"terabytes"` + "`" + `, ` + "`" + `"bits"` + "`" + `, ` + "`" + `"kilobits"` + "`" + `, ` + "`" + `"megabits"` + "`" + `, ` + "`" + `"gigabits"` + "`" + `, ` + "`" + `"terabits"` + "`" + `, ` + "`" + `"count"` + "`" + `, ` + "`" + `"bytes/second"` + "`" + `, ` + "`" + `"kilobytes/second"` + "`" + `, ` + "`" + `"megabytes/second"` + "`" + `, ` + "`" + `"gigabytes/second"` + "`" + `, ` + "`" + `"terabytes/second"` + "`" + `, ` + "`" + `"bits/second"` + "`" + `, ` + "`" + `"kilobits/second"` + "`" + `, ` + "`" + `"megabits/second"` + "`" + `, ` + "`" + `"gigabits/second"` + "`" + `, ` + "`" + `"terabits/second"` + "`" + `, ` + "`" + `"count/second"` + "`" + `, ` + "`" + `"none"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "cooldown",
					Description: `(Optional, Default: ` + "`" + `300` + "`" + `) The amount of time, in seconds, after a scaling activity completes and before the next scaling activity can start. If this parameter is not specified, the default cooldown period for the group applies.`,
				},
				resource.Attribute{
					Name:        "operator",
					Description: `(Optional, Scale Up Default: ` + "`" + `gte` + "`" + `, Scale Down Default: ` + "`" + `lte` + "`" + `) The operator to use in order to determine if the scaling policy is applicable. Valid values: ` + "`" + `"gt"` + "`" + `, ` + "`" + `"gte"` + "`" + `, ` + "`" + `"lt"` + "`" + `, ` + "`" + `"lte"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "evaluation_periods",
					Description: `(Optional, Default: ` + "`" + `1` + "`" + `) The number of periods over which data is compared to the specified threshold.`,
				},
				resource.Attribute{
					Name:        "period",
					Description: `(Optional, Default: ` + "`" + `300` + "`" + `) The granularity, in seconds, of the returned datapoints. Period must be at least 60 seconds and must be a multiple of 60.`,
				},
				resource.Attribute{
					Name:        "dimensions",
					Description: `(Optional) A list of dimensions describing qualities of the metric. Required when ` + "`" + `namespace` + "`" + ` is defined AND not ` + "`" + `"Microsoft.Compute"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The dimension name.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Optional) The dimension value. When ` + "`" + `namespace` + "`" + ` is defined and is not ` + "`" + `"Microsoft.Compute"` + "`" + ` the list of dimensions must contain the following: ` + "`" + `` + "`" + `` + "`" + `hcl dimensions = [ { name = "resourceName" value = "example-resource-name" }, { name = "resourceGroupName" value = "example-resource-group-name" }, ] ` + "`" + `` + "`" + `` + "`" + ``,
				},
				resource.Attribute{
					Name:        "action_type",
					Description: `(Optional; if not using ` + "`" + `min_target_capacity` + "`" + ` or ` + "`" + `max_target_capacity` + "`" + `) The type of action to perform for scaling. Valid values: ` + "`" + `"adjustment"` + "`" + `, ` + "`" + `"percentageAdjustment"` + "`" + `, ` + "`" + `"setMaxTarget"` + "`" + `, ` + "`" + `"setMinTarget"` + "`" + `, ` + "`" + `"updateCapacity"` + "`" + `. If you do not specify an action type, you can only use – ` + "`" + `adjustment` + "`" + `, ` + "`" + `min_target_capacity` + "`" + `, ` + "`" + `max_target_capacity` + "`" + `. While using action_type, please also set the following: When using ` + "`" + `adjustment` + "`" + ` – set the field ` + "`" + `adjustment` + "`" + ` When using ` + "`" + `percentageAdjustment` + "`" + ` - set the field ` + "`" + `adjustment` + "`" + ` When using ` + "`" + `setMaxTarget` + "`" + ` – set the field ` + "`" + `max_target_capacity` + "`" + ` When using ` + "`" + `setMinTarget` + "`" + ` – set the field ` + "`" + `min_target_capacity` + "`" + ` When using ` + "`" + `updateCapacity` + "`" + ` – set the fields ` + "`" + `minimum` + "`" + `, ` + "`" + `maximum` + "`" + `, and ` + "`" + `target` + "`" + ``,
				},
				resource.Attribute{
					Name:        "adjustment",
					Description: `(Optional) Value to which the action type will be adjusted. Required if using ` + "`" + `numeric` + "`" + ` or ` + "`" + `percentage_adjustment` + "`" + ` action types.`,
				},
				resource.Attribute{
					Name:        "min_target_capacity",
					Description: `(Optional; if not using ` + "`" + `adjustment` + "`" + `; available only for scale up). The number of the desired target (and minimum) capacity`,
				},
				resource.Attribute{
					Name:        "max_target_capacity",
					Description: `(Optional; if not using ` + "`" + `adjustment` + "`" + `; available only for scale down). The number of the desired target (and maximum) capacity`,
				},
				resource.Attribute{
					Name:        "minimum",
					Description: `(Optional; if using ` + "`" + `updateCapacity` + "`" + `) The minimal number of instances to have in the group.`,
				},
				resource.Attribute{
					Name:        "maximum",
					Description: `(Optional; if using ` + "`" + `updateCapacity` + "`" + `) The maximal number of instances to have in the group.`,
				},
				resource.Attribute{
					Name:        "target",
					Description: `(Optional; if using ` + "`" + `updateCapacity` + "`" + `) The target number of instances to have in the group. Usage: ` + "`" + `` + "`" + `` + "`" + `hcl // --- SCALE DOWN POLICY ------------------ scaling_down_policy = [{ policy_name = "policy-name" metric_name = "CPUUtilization" namespace = "Microsoft.Compute" statistic = "average" threshold = 10 unit = "percent" cooldown = 60 dimensions = { name = "name-1" value = "value-1" } operator = "gt" evaluation_periods = "10" period = "60" // === MIN TARGET =================== # action_type = "setMinTarget" # min_target_capacity = 1 // ================================== // === ADJUSTMENT =================== action_type = "adjustment" # action_type = "percentageAdjustment" adjustment = "MIN(5,10)" // ================================== // === UPDATE CAPACITY ============== # action_type = "updateCapacity" # minimum = 0 # maximum = 10 # target = 5 // ================================== }] // ---------------------------------------- // --- SCALE DOWN POLICY ------------------ scaling_down_policy = [{ policy_name = "policy-name-update" metric_name = "CPUUtilization" namespace = "Microsoft.Compute" statistic = "sum" threshold = 5 unit = "bytes" cooldown = 120 dimensions = { name = "name-1-update" value = "value-1-update" } operator = "lt" evaluation_periods = 5 period = 120 //// === MIN TARGET =================== # action_type = "setMinTarget" # min_target_capacity = 1 //// ================================== // === ADJUSTMENT =================== # action_type = "percentageAdjustment" # action_type = "adjustment" # adjustment = "MAX(5,10)" // ================================== // === UPDATE CAPACITY ============== action_type = "updateCapacity" minimum = 0 maximum = 10 target = 5 // ================================== }] // ---------------------------------------- ` + "`" + `` + "`" + `` + "`" + ` <a id="scheduling"></a> ## Scheduling`,
				},
				resource.Attribute{
					Name:        "scheduled_task",
					Description: `(Optional) Describes the configuration of one or more scheduled tasks.`,
				},
				resource.Attribute{
					Name:        "is_enabled",
					Description: `(Optional, Default: ` + "`" + `true` + "`" + `) Describes whether the task is enabled. When true the task should run when false it should not run.`,
				},
				resource.Attribute{
					Name:        "cron_expression",
					Description: `(Required) A valid cron expression (` + "`" + ``,
				},
				resource.Attribute{
					Name:        "task_type",
					Description: `(Required) The task type to run. Valid Values: ` + "`" + `backup_ami` + "`" + `, ` + "`" + `scale` + "`" + `, ` + "`" + `scaleUp` + "`" + `, ` + "`" + `roll` + "`" + `, ` + "`" + `statefulUpdateCapacity` + "`" + `, ` + "`" + `statefulRecycle` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "scale_min_capacity",
					Description: `(Optional) The min capacity of the group. Should be used when choosing ‘task_type' of ‘scale'.`,
				},
				resource.Attribute{
					Name:        "scale_max_capacity",
					Description: `(Optional) The max capacity of the group. Required when ‘task_type' is ‘scale'.`,
				},
				resource.Attribute{
					Name:        "scale_target_capacity",
					Description: `(Optional) The target capacity of the group. Should be used when choosing ‘task_type' of ‘scale'.`,
				},
				resource.Attribute{
					Name:        "adjustment",
					Description: `(Optional) The number of instances to add/remove to/from the target capacity when scale is needed.`,
				},
				resource.Attribute{
					Name:        "adjustment_percentage",
					Description: `(Optional) The percent of instances to add/remove to/from the target capacity when scale is needed.`,
				},
				resource.Attribute{
					Name:        "batch_size_percentage",
					Description: `(Optional) The percentage size of each batch in the scheduled deployment roll. Required when the 'task_type' is 'roll'.`,
				},
				resource.Attribute{
					Name:        "grace_period",
					Description: `(Optional) The time to allow instances to become healthy. ` + "`" + `` + "`" + `` + "`" + `hcl scheduled_task = [{ is_enabled = true cron_expression = "`,
				},
				resource.Attribute{
					Name:        "update_policy",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "should_roll",
					Description: `(Required) Sets the enablement of the roll option.`,
				},
				resource.Attribute{
					Name:        "roll_config",
					Description: `(Required) While used, you can control whether the group should perform a deployment after an update to the configuration.`,
				},
				resource.Attribute{
					Name:        "batch_size_percentage",
					Description: `(Required) Sets the percentage of the instances to deploy in each batch.`,
				},
				resource.Attribute{
					Name:        "health_check_type",
					Description: `(Optional) Sets the health check type to use. Valid values: ` + "`" + `"INSTANCE_STATE"` + "`" + `, ` + "`" + `"NONE"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "grace_period",
					Description: `(Optional) Sets the grace period for new instances to become healthy. ` + "`" + `` + "`" + `` + "`" + `hcl update_policy = { should_roll = false roll_config = { batch_size_percentage = 33 health_check_type = "INSTANCE_STATE" grace_period = 300 } } ` + "`" + `` + "`" + `` + "`" + ` <a id="third-party-integrations"></a> ## Third-Party Integrations`,
				},
				resource.Attribute{
					Name:        "integration_kubernetes",
					Description: `(Optional) Describes the [Kubernetes](https://kubernetes.io/) integration.`,
				},
				resource.Attribute{
					Name:        "cluster_identifier",
					Description: `(Required) The cluster ID. Usage: ` + "`" + `` + "`" + `` + "`" + `hcl integration_kubernetes = { cluster_identifier = "k8s-cluster-id" } ` + "`" + `` + "`" + `` + "`" + ``,
				},
				resource.Attribute{
					Name:        "integration_multai_runtime",
					Description: `(Optional) Describes the [Multai Runtime](https://spotinst.com/) integration.`,
				},
				resource.Attribute{
					Name:        "deployment_id",
					Description: `(Optional) The deployment id you want to get Usage: ` + "`" + `` + "`" + `` + "`" + `hcl integration_multai_runtime = { deployment_id = "" } ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "elastigroup_gcp",
			Category:         "Spotinst Resources",
			ShortDescription: `Provides a Spotinst elastigroup resource for Google Cloud.`,
			Description:      ``,
			Keywords: []string{
				"elastigroup",
				"gcp",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The group name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The region your GCP group will be created in.`,
				},
				resource.Attribute{
					Name:        "startup_script",
					Description: `(Optional) Create and run your own startup scripts on your virtual machines to perform automated tasks every time your instance boots up.`,
				},
				resource.Attribute{
					Name:        "shutdown_script",
					Description: `(Optional) The Base64-encoded shutdown script that executes prior to instance termination, for more information please see: [Shutdown Script](https://api.spotinst.com/integration-docs/elastigroup/concepts/compute-concepts/shutdown-scripts/)`,
				},
				resource.Attribute{
					Name:        "service_account",
					Description: `(Optional) The email of the service account in which the group instances will be launched.`,
				},
				resource.Attribute{
					Name:        "max_size",
					Description: `(Required) The maximum number of instances the group should have at any time.`,
				},
				resource.Attribute{
					Name:        "min_size",
					Description: `(Required) The minimum number of instances the group should have at any time.`,
				},
				resource.Attribute{
					Name:        "desired_capacity",
					Description: `(Required) The desired number of instances the group should have at any time.`,
				},
				resource.Attribute{
					Name:        "availability_zones",
					Description: `(Required) List of availability zones for the group.`,
				},
				resource.Attribute{
					Name:        "subnets",
					Description: `(Optional) A list of regions and subnets.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Required) The region for the group of subnets.`,
				},
				resource.Attribute{
					Name:        "subnet_names",
					Description: `(Required) The names of the subnets in the region.`,
				},
				resource.Attribute{
					Name:        "instance_types_preemptible",
					Description: `(Required) The preemptible VMs instance type. To maximize cost savings and market availability, select as many types as possible. Required if instance_types_ondemand is not set.`,
				},
				resource.Attribute{
					Name:        "instance_types_ondemand",
					Description: `(Required) The regular VM instance type to use for mixed-type groups and when falling back to on-demand. Required if instance_types_preemptible is not set.`,
				},
				resource.Attribute{
					Name:        "instance_types_custom",
					Description: `(Required) Defines a set of custom instance types. Required if instance_types_preemptible and instance_types_ondemand are not set.`,
				},
				resource.Attribute{
					Name:        "vCPU",
					Description: `(Optional) The number of vCPUs in the custom instance type. GCP has a number of limitations on accepted vCPU values. For more information, see the GCP documentation (here.)[https://cloud.google.com/compute/docs/instances/creating-instance-with-custom-machine-type#specifications]`,
				},
				resource.Attribute{
					Name:        "memory_gib",
					Description: `(Optional) The memory (in GiB) in the custom instance types. GCP has a number of limitations on accepted memory values.For more information, see the GCP documentation (here.)[https://cloud.google.com/compute/docs/instances/creating-instance-with-custom-machine-type#specifications]`,
				},
				resource.Attribute{
					Name:        "preemptible_percentage",
					Description: `(Optional) Percentage of Preemptible VMs to spin up from the "desired_capacity".`,
				},
				resource.Attribute{
					Name:        "on_demand_count",
					Description: `(Optional) Number of regular VMs to launch in the group. The rest will be Preemptible VMs. When this parameter is specified, the preemptible_percentage parameter is being ignored.`,
				},
				resource.Attribute{
					Name:        "fallback_to_ondemand",
					Description: `(Optional) Activate fallback-to-on-demand. When provisioning an instance, if no Preemptible market is available, fallback-to-on-demand will provision an On-Demand instance to maintain the group capacity.`,
				},
				resource.Attribute{
					Name:        "draining_timeout",
					Description: `(Optional) Time (seconds) the instance is allowed to run after it is detached from the group. This is to allow the instance time to drain all the current TCP connections before terminating it.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `(Optional) Array of objects with key-value pairs.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Optional) Metadata key.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Optional) Metadata value.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) Array of objects with key-value pairs.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Optional) Labels key.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Optional) Labels value.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags to mark created instances. <a id="GPU"></a> ## GPU`,
				},
				resource.Attribute{
					Name:        "gpu",
					Description: `(Optional) Defines the GPU configuration.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) The type of GPU instance. Valid values: ` + "`" + `nvidia-tesla-v100` + "`" + `, ` + "`" + `nvidia-tesla-p100` + "`" + `, ` + "`" + `nvidia-tesla-k80` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "count",
					Description: `(Required) The number of GPUs. Must be 0, 2, 4, 6, 8. Usage: ` + "`" + `` + "`" + `` + "`" + `hcl gpu = { count = 2 type = "nvidia-tesla-p100" } ` + "`" + `` + "`" + `` + "`" + ` <a id="health-check"></a> ## Health Check`,
				},
				resource.Attribute{
					Name:        "auto_healing",
					Description: `(Optional) Enable auto-replacement of unhealthy instances.`,
				},
				resource.Attribute{
					Name:        "health_check_grace_period",
					Description: `(Optional) Period of time (seconds) to wait for VM to reach healthiness before monitoring for unhealthiness.`,
				},
				resource.Attribute{
					Name:        "health_check_type",
					Description: `(Optional) The kind of health check to perform when monitoring for unhealthiness.`,
				},
				resource.Attribute{
					Name:        "unhealthy_duration",
					Description: `(Optional) Period of time (seconds) to remain in an unhealthy status before a replacement is triggered. ` + "`" + `` + "`" + `` + "`" + `hcl auto_health = true health_check_grace_period = 100 health_check_type = "" unhealthy_duration = "" ` + "`" + `` + "`" + `` + "`" + ` <a id="backend-services"></a> ## Backend Services`,
				},
				resource.Attribute{
					Name:        "backend_services",
					Description: `(Optional) Describes the backend service configurations.`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `(Required) The name of the backend service.`,
				},
				resource.Attribute{
					Name:        "location_type",
					Description: `(Optional) Sets which location the backend services will be active. Valid values: ` + "`" + `regional` + "`" + `, ` + "`" + `global` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "scheme",
					Description: `(Optional) Use when ` + "`" + `location_type` + "`" + ` is "regional". Set the traffic for the backend service to either between the instances in the vpc or to traffic from the internet. Valid values: ` + "`" + `INTERNAL` + "`" + `, ` + "`" + `EXTERNAL` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "named_port",
					Description: `(Optional) Describes a named port and a list of ports.`,
				},
				resource.Attribute{
					Name:        "port_name",
					Description: `(Required) The name of the port.`,
				},
				resource.Attribute{
					Name:        "ports",
					Description: `(Required) A list of ports. Usage: ` + "`" + `` + "`" + `` + "`" + `hcl backend_services_config = [ { service_name = "spotinst-elb-backend-service" locationType = "regional" scheme = "INTERNAL" ports = { port_name = "port-name" ports = [8000, 6000] } }, ] ` + "`" + `` + "`" + `` + "`" + ` <a id="disks"></a> ## Disks`,
				},
				resource.Attribute{
					Name:        "disks",
					Description: `(Optional) Array of disks associated with this instance. Persistent disks must be created before you can assign them.`,
				},
				resource.Attribute{
					Name:        "auto_delete",
					Description: `(Optional) Specifies whether the disk will be auto-deleted when the instance is deleted.`,
				},
				resource.Attribute{
					Name:        "boot",
					Description: `(Optional) Indicates that this is a boot disk. The virtual machine will use the first partition of the disk for its root filesystem.`,
				},
				resource.Attribute{
					Name:        "device_name",
					Description: `(Optional) Specifies a unique device name of your choice.`,
				},
				resource.Attribute{
					Name:        "interface",
					Description: `(Optional, Default: ` + "`" + `SCSI` + "`" + `) Specifies the disk interface to use for attaching this disk, which is either SCSI or NVME.`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `(Optional, Default: ` + "`" + `READ_WRITE` + "`" + `) The mode in which to attach this disk, either READ_WRITE or READ_ONLY.`,
				},
				resource.Attribute{
					Name:        "source",
					Description: `(Optional) Specifies a valid partial or full URL to an existing Persistent Disk resource. This field is only applicable for persistent disks.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional, Default: ` + "`" + `PERSISTENT` + "`" + `) Specifies the type of disk, either SCRATCH or PERSISTENT.`,
				},
				resource.Attribute{
					Name:        "initialize_params",
					Description: `(Optional) Specifies the parameters for a new disk that will be created alongside the new instance. Use initialization parameters to create boot disks or local SSDs attached to the new instance.`,
				},
				resource.Attribute{
					Name:        "disk_size_gb",
					Description: `(Optional) Specifies disk size in gigabytes. Must be in increments of 2.`,
				},
				resource.Attribute{
					Name:        "disk_type",
					Description: `(Optional, Default" ` + "`" + `pd-standard` + "`" + `) Specifies the disk type to use to create the instance. Valid values: pd-ssd, local-ssd.`,
				},
				resource.Attribute{
					Name:        "source_image",
					Description: `(Optional) A source image used to create the disk. You can provide a private (custom) image, and Compute Engine will use the corresponding image from your project. Usage: ` + "`" + `` + "`" + `` + "`" + `hcl disks = [ { device_name = "device" mode = "READ_WRITE" type = "PERSISTENT" auto_delete = true boot = true interface = "SCSI" initialize_params = { disk_size_gb = 10 disk_type = "pd-standard" source_image = "" } } ] ` + "`" + `` + "`" + `` + "`" + ` <a id="network-interface"></a> ## Network Interfaces Each of the ` + "`" + `network_interface` + "`" + ` attributes controls a portion of the GCP Instance's "Network Interfaces". It's a good idea to familiarize yourself with [GCP's Network Interfaces docs](https://cloud.google.com/vpc/docs/multiple-interfaces-concepts) to understand the implications of using these attributes.`,
				},
				resource.Attribute{
					Name:        "network_interface",
					Description: `(Required, minimum 1) Array of objects representing the network configuration for the elastigroup.`,
				},
				resource.Attribute{
					Name:        "network",
					Description: `(Required) Network resource for this group.`,
				},
				resource.Attribute{
					Name:        "access_configs",
					Description: `(Optional) Array of configurations.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of this access configuration.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) Array of configurations for this interface. Currently, only ONE_TO_ONE_NAT is supported. ` + "`" + `` + "`" + `` + "`" + `hcl network_interface = [{ network = "default" access_configs = { name = "config1" type = "ONE_TO_ONE_NAT" } alias_ip_ranges = { subnetwork_range_name = "range-name-1" ip_cidr_range = "10.128.0.0/20" } }] ` + "`" + `` + "`" + `` + "`" + ` <a id="scaling-policy"></a> ## Scaling Policies`,
				},
				resource.Attribute{
					Name:        "scaling_up_policy",
					Description: `(Optional) Contains scaling policies for scaling the Elastigroup up.`,
				},
				resource.Attribute{
					Name:        "scaling_down_policy",
					Description: `(Optional) Contains scaling policies for scaling the Elastigroup down. Each ` + "`" + `scaling_`,
				},
				resource.Attribute{
					Name:        "policy_name",
					Description: `(Optional) Name of scaling policy.`,
				},
				resource.Attribute{
					Name:        "metric_name",
					Description: `(Optional) Metric to monitor. Valid values: "Percentage CPU", "Network In", "Network Out", "Disk Read Bytes", "Disk Write Bytes", "Disk Write Operations/Sec", "Disk Read Operations/Sec".`,
				},
				resource.Attribute{
					Name:        "statistic",
					Description: `(Optional) Statistic by which to evaluate the selected metric. Valid values: "AVERAGE", "SAMPLE_COUNT", "SUM", "MINIMUM", "MAXIMUM", "PERCENTILE", "COUNT".`,
				},
				resource.Attribute{
					Name:        "threshold",
					Description: `(Optional) The value at which the scaling action is triggered.`,
				},
				resource.Attribute{
					Name:        "period",
					Description: `(Optional) Amount of time (seconds) for which the threshold must be met in order to trigger the scaling action.`,
				},
				resource.Attribute{
					Name:        "evaluation_periods",
					Description: `(Optional) Number of consecutive periods in which the threshold must be met in order to trigger a scaling action.`,
				},
				resource.Attribute{
					Name:        "cooldown",
					Description: `(Optional) Time (seconds) to wait after a scaling action before resuming monitoring.`,
				},
				resource.Attribute{
					Name:        "operator",
					Description: `(Optional) The operator used to evaluate the threshold against the current metric value. Valid values: "gt" (greater than), "get" (greater-than or equal), "lt" (less than), "lte" (less than or equal).`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Optional) Scaling action to take when the policy is triggered.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) Type of scaling action to take when the scaling policy is triggered. Valid values: "adjustment", "setMinTarget", "updateCapacity", "percentageAdjustment"`,
				},
				resource.Attribute{
					Name:        "adjustment",
					Description: `(Optional) Value to which the action type will be adjusted. Required if using "numeric" or "percentageAdjustment" action types.`,
				},
				resource.Attribute{
					Name:        "dimensions",
					Description: `(Optional) A list of dimensions describing qualities of the metric.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The dimension name.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) The dimension value. Usage: ` + "`" + `` + "`" + `` + "`" + `hcl scaling = { up = { policy_name = "scale_up_1" source = "stackdriver" metric_name = "instance/disk/read_ops_count" namespace = "compute" statistic = "average" unit = "percent" threshold = 10000 period = 300 cooldown = 300 operator = "gte" evaluation_periods = 1 action = { type = "adjustment" adjustment = 1 } dimensions = [ { name = "storage_type" value = "pd-ssd" } ] } } ` + "`" + `` + "`" + `` + "`" + ` <a id="third-party-integrations"></a> ## Third-Party Integrations`,
				},
				resource.Attribute{
					Name:        "integration_docker_swarm",
					Description: `(Optional) Describes the [Docker Swarm](https://api.spotinst.com/integration-docs/elastigroup/container-management/docker-swarm/docker-swarm-integration/) integration.`,
				},
				resource.Attribute{
					Name:        "master_host",
					Description: `(Required) IP or FQDN of one of your swarm managers.`,
				},
				resource.Attribute{
					Name:        "master_port",
					Description: `(Required) Network port used by your swarm. Usage: ` + "`" + `` + "`" + `` + "`" + `hcl integration_docker_swarm = { master_host = "10.10.10.10" master_port = 2376 } ` + "`" + `` + "`" + `` + "`" + ` <a id="scheduled-task"></a> ## Scheduled Tasks Each ` + "`" + `scheduled_task` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "task_type",
					Description: `(Required) The task type to run. Valid values: ` + "`" + `"setCapacity"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "cron_expression",
					Description: `(Optional) A valid cron expression. The cron is running in UTC time zone and is in [Unix cron format](https://en.wikipedia.org/wiki/Cron).`,
				},
				resource.Attribute{
					Name:        "is_enabled",
					Description: `(Optional, Default: ` + "`" + `true` + "`" + `) Setting the task to being enabled or disabled.`,
				},
				resource.Attribute{
					Name:        "target_capacity",
					Description: `(Optional) The desired number of instances the group should have.`,
				},
				resource.Attribute{
					Name:        "min_capacity",
					Description: `(Optional) The minimum number of instances the group should have.`,
				},
				resource.Attribute{
					Name:        "max_capacity",
					Description: `(Optional) The maximum number of instances the group should have. Usage: ` + "`" + `` + "`" + `` + "`" + `hcl scheduled_task = [{ task_type = "setCapacity" cron_expression = "" is_enabled = false target_capacity = 5 min_capacity = 0 max_capacity = 10 }] ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "elastigroup_gke",
			Category:         "Spotinst Resources",
			ShortDescription: `Provides a Spotinst elastigroup resource for Google Cloud using the Google Kubernetes Engine.`,
			Description:      ``,
			Keywords: []string{
				"elastigroup",
				"gke",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cluster_zone_name",
					Description: `(Required) The zone where the cluster is hosted.`,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Required) The name of the GKE cluster you wish to import.`,
				},
				resource.Attribute{
					Name:        "node_image",
					Description: `(Optional, Default: ` + "`" + `COS` + "`" + `) The image that will be used for the node VMs. Possible values: COS, UBUNTU. <a id="third-party-integrations"></a> ## Third-Party Integrations`,
				},
				resource.Attribute{
					Name:        "integration_gke",
					Description: `(Required) Describes the [GKE]() integration.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `(Optional) The location of your GKE cluster.`,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Optional) The GKE cluster ID you wish to import.`,
				},
				resource.Attribute{
					Name:        "autoscale_is_enabled",
					Description: `(Optional, Default: ` + "`" + `false` + "`" + `) Specifies whether the auto scaling feature is enabled.`,
				},
				resource.Attribute{
					Name:        "autoscale_is_autoconfig",
					Description: `(Optional, Default: ` + "`" + `false` + "`" + `) Enabling the automatic auto-scaler functionality. For more information please see: []().`,
				},
				resource.Attribute{
					Name:        "autoscale_cooldown",
					Description: `(Optional, Default: ` + "`" + `300` + "`" + `) The amount of time, in seconds, after a scaling activity completes before any further trigger-related scaling activities can start.`,
				},
				resource.Attribute{
					Name:        "autoscale_headroom",
					Description: `(Optional) Headroom for the cluster.`,
				},
				resource.Attribute{
					Name:        "cpu_per_unit",
					Description: `(Optional, Default: ` + "`" + `0` + "`" + `) Cpu units for compute.`,
				},
				resource.Attribute{
					Name:        "memory_per_unit",
					Description: `(Optional, Default: ` + "`" + `0` + "`" + `) RAM units for compute.`,
				},
				resource.Attribute{
					Name:        "num_of_units",
					Description: `(Optional, Default: ` + "`" + `0` + "`" + `) Amount of units for compute.`,
				},
				resource.Attribute{
					Name:        "autoscale_down",
					Description: `(Optional) Enabling scale down.`,
				},
				resource.Attribute{
					Name:        "evaluation_periods",
					Description: `(Optional, Default: ` + "`" + `5` + "`" + `) Amount of cooldown evaluation periods for scale down.`,
				},
				resource.Attribute{
					Name:        "autoscale_labels",
					Description: `(Optional) Labels to assign to the resource.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Optional) The label name.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Optional) The label value. Usage: ` + "`" + `` + "`" + `` + "`" + `hcl integration_gke = { location = "us-central1-a" cluster_id = "terraform-acc-test-cluster" autoscale_is_enabled = true autoscale_is_auto_config = false autoscale_cooldown = 300 autoscale_headroom = { cpu_per_unit = 1024 memory_per_unit = 512 num_of_units = 2 } autoscale_down = { evaluation_periods = 300 } autoscale_labels = { key = "label_key" value = "label_value" } } ` + "`" + `` + "`" + `` + "`" + ` <a id="diff-suppressed-parameters"></a> ## Diff-suppressed Parameters The following parameters are created remotely and imported. The diffs have been suppressed in order to maintain plan legibility. You may update the values of these imported parameters by defining them in your template with your desired new value (including null values).`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mrscaler_aws",
			Category:         "Spotinst Resources",
			ShortDescription: `Provides a Spotinst MrScaler resource.`,
			Description:      ``,
			Keywords: []string{
				"mrscaler",
				"aws",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The MrScaler name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The MrScaler description.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Required) The MrScaler region.`,
				},
				resource.Attribute{
					Name:        "strategy",
					Description: `(Required) The MrScaler strategy. Allowed values are ` + "`" + `new` + "`" + ` ` + "`" + `clone` + "`" + ` and ` + "`" + `wrap` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Optional) The MrScaler cluster id.`,
				},
				resource.Attribute{
					Name:        "expose_cluster_id",
					Description: `(Optional) Allow the ` + "`" + `cluster_id` + "`" + ` to set a Terraform output variable. <a id="provisioning-timeout"></a> ## Provisioning Timeout (Clone, New strategies)`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `(Optional) The amount of time (minutes) after which the cluster is automatically terminated if it's still in provisioning status. Minimum: '15'.`,
				},
				resource.Attribute{
					Name:        "timeout_action",
					Description: `(Optional) The action to take if the timeout is exceeded. Valid values: ` + "`" + `terminate` + "`" + `, ` + "`" + `terminateAndRetry` + "`" + `. <a id="cluster-config"></a> ## Cluster Configuration (New strategy only)`,
				},
				resource.Attribute{
					Name:        "log_uri",
					Description: `(Optional) The path to the Amazon S3 location where logs for this cluster are stored.`,
				},
				resource.Attribute{
					Name:        "additional_info",
					Description: `(Optional) This is meta information about third-party applications that third-party vendors use for testing purposes.`,
				},
				resource.Attribute{
					Name:        "security_config",
					Description: `(Optional) The name of the security configuration applied to the cluster.`,
				},
				resource.Attribute{
					Name:        "service_role",
					Description: `(Optional) The IAM role that will be assumed by the Amazon EMR service to access AWS resources on your behalf.`,
				},
				resource.Attribute{
					Name:        "job_flow_role",
					Description: `(Optional) The IAM role that was specified when the job flow was launched. The EC2 instances of the job flow assume this role.`,
				},
				resource.Attribute{
					Name:        "termination_protected",
					Description: `(Optional) Specifies whether the Amazon EC2 instances in the cluster are protected from termination by API calls, user intervention, or in the event of a job-flow error.`,
				},
				resource.Attribute{
					Name:        "keep_job_flow_alive",
					Description: `(Optional) Specifies whether the cluster should remain available after completing all steps. <a id="task-group"></a> ## Task Group (Wrap, Clone, and New strategies)`,
				},
				resource.Attribute{
					Name:        "task_instance_types",
					Description: `(Required) The MrScaler instance types for the task nodes.`,
				},
				resource.Attribute{
					Name:        "task_target",
					Description: `(Required) amount of instances in task group.`,
				},
				resource.Attribute{
					Name:        "task_maximum",
					Description: `(Optional) maximal amount of instances in task group.`,
				},
				resource.Attribute{
					Name:        "task_minimum",
					Description: `(Optional) The minimal amount of instances in task group.`,
				},
				resource.Attribute{
					Name:        "task_lifecycle",
					Description: `(Required) The MrScaler lifecycle for instances in task group. Allowed values are 'SPOT' and 'ON_DEMAND'.`,
				},
				resource.Attribute{
					Name:        "task_ebs_optimized",
					Description: `(Optional) EBS Optimization setting for instances in group.`,
				},
				resource.Attribute{
					Name:        "task_ebs_block_device",
					Description: `(Required) This determines the ebs configuration for your task group instances. Only a single block is allowed.`,
				},
				resource.Attribute{
					Name:        "volumes_per_instance",
					Description: `(Optional; Default 1) Amount of volumes per instance in the task group.`,
				},
				resource.Attribute{
					Name:        "volume_type",
					Description: `(Required) volume type. Allowed values are 'gp2', 'io1' and others.`,
				},
				resource.Attribute{
					Name:        "size_in_gb",
					Description: `(Required) Size of the volume, in GBs.`,
				},
				resource.Attribute{
					Name:        "iops",
					Description: `(Optional) IOPS for the volume. Required in some volume types, such as io1. <a id="core-group"></a> ## Core Group (Clone, New strategies)`,
				},
				resource.Attribute{
					Name:        "core_instance_types",
					Description: `(Required) The MrScaler instance types for the core nodes.`,
				},
				resource.Attribute{
					Name:        "core_target",
					Description: `(Required) amount of instances in core group.`,
				},
				resource.Attribute{
					Name:        "core_maximum",
					Description: `(Optional) maximal amount of instances in core group.`,
				},
				resource.Attribute{
					Name:        "core_minimum",
					Description: `(Optional) The minimal amount of instances in core group.`,
				},
				resource.Attribute{
					Name:        "core_lifecycle",
					Description: `(Required) The MrScaler lifecycle for instances in core group. Allowed values are 'SPOT' and 'ON_DEMAND'.`,
				},
				resource.Attribute{
					Name:        "core_ebs_optimized",
					Description: `(Optional) EBS Optimization setting for instances in group.`,
				},
				resource.Attribute{
					Name:        "core_ebs_block_device",
					Description: `(Required) This determines the ebs configuration for your core group instances. Only a single block is allowed.`,
				},
				resource.Attribute{
					Name:        "volumes_per_instance",
					Description: `(Optional; Default 1) Amount of volumes per instance in the core group.`,
				},
				resource.Attribute{
					Name:        "volume_type",
					Description: `(Required) volume type. Allowed values are 'gp2', 'io1' and others.`,
				},
				resource.Attribute{
					Name:        "size_in_gb",
					Description: `(Required) Size of the volume, in GBs.`,
				},
				resource.Attribute{
					Name:        "iops",
					Description: `(Optional) IOPS for the volume. Required in some volume types, such as io1. <a id="master-group"></a> ## Master Group (Clone, New strategies)`,
				},
				resource.Attribute{
					Name:        "master_instance_types",
					Description: `(Required) The MrScaler instance types for the master nodes.`,
				},
				resource.Attribute{
					Name:        "master_lifecycle",
					Description: `(Required) The MrScaler lifecycle for instances in master group. Allowed values are 'SPOT' and 'ON_DEMAND'.`,
				},
				resource.Attribute{
					Name:        "master_ebs_optimized",
					Description: `(Optional) EBS Optimization setting for instances in group.`,
				},
				resource.Attribute{
					Name:        "master_ebs_block_device",
					Description: `(Required) This determines the ebs configuration for your master group instances. Only a single block is allowed.`,
				},
				resource.Attribute{
					Name:        "volumes_per_instance",
					Description: `(Optional; Default 1) Amount of volumes per instance in the master group.`,
				},
				resource.Attribute{
					Name:        "volume_type",
					Description: `(Required) volume type. Allowed values are 'gp2', 'io1' and others.`,
				},
				resource.Attribute{
					Name:        "size_in_gb",
					Description: `(Required) Size of the volume, in GBs.`,
				},
				resource.Attribute{
					Name:        "iops",
					Description: `(Optional) IOPS for the volume. Required in some volume types, such as io1. <a id="tags"></a> ## Tags (Clone, New strategies)`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A list of tags to assign to the resource. You may define multiple tags.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) Tag key.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) Tag value. <a id="Optional Compute Parameters"></a> ## Optional Compute Parameters (New strategy)`,
				},
				resource.Attribute{
					Name:        "managed_primary_security_group",
					Description: `(Optional) EMR Managed Security group that will be set to the primary instance group.`,
				},
				resource.Attribute{
					Name:        "managed_replica_security_group",
					Description: `(Optional) EMR Managed Security group that will be set to the replica instance group.`,
				},
				resource.Attribute{
					Name:        "service_access_security_group",
					Description: `(Optional) The identifier of the Amazon EC2 security group for the Amazon EMR service to access clusters in VPC private subnets.`,
				},
				resource.Attribute{
					Name:        "additional_primary_security_groups",
					Description: `(Optional) A list of additional Amazon EC2 security group IDs for the master node.`,
				},
				resource.Attribute{
					Name:        "additional_replica_security_groups",
					Description: `(Optional) A list of additional Amazon EC2 security group IDs for the core and task nodes.`,
				},
				resource.Attribute{
					Name:        "custom_ami_id",
					Description: `(Optional) The ID of a custom Amazon EBS-backed Linux AMI if the cluster uses a custom AMI.`,
				},
				resource.Attribute{
					Name:        "repo_upgrade_on_boot",
					Description: `(Optional) Applies only when ` + "`" + `custom_ami_id` + "`" + ` is used. Specifies the type of updates that are applied from the Amazon Linux AMI package repositories when an instance boots using the AMI. Possible values include: ` + "`" + `SECURITY` + "`" + `, ` + "`" + `NONE` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ec2_key_name",
					Description: `(Optional) The name of an Amazon EC2 key pair that can be used to ssh to the master node.`,
				},
				resource.Attribute{
					Name:        "applications",
					Description: `(Optional) A case-insensitive list of applications for Amazon EMR to install and configure when launching the cluster`,
				},
				resource.Attribute{
					Name:        "args",
					Description: `(Optional) Arguments for EMR to pass to the application.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The application name.`,
				},
				resource.Attribute{
					Name:        "instance_weights",
					Description: `(Optional) Describes the instance and weights. Check out [Elastigroup Weighted Instances](https://api.spotinst.com/elastigroup-for-aws/concepts/general-concepts/elastigroup-capacity-instances-or-weighted) for more info.`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `(Required) The type of the instance.`,
				},
				resource.Attribute{
					Name:        "weighted_capacity",
					Description: `(Required) The weight given to the associated instance type. <a id="availability-zone"></a> ## Availability Zones (Clone, New strategies)`,
				},
				resource.Attribute{
					Name:        "availability_zones",
					Description: `(Required in Clone) List of AZs and their subnet Ids. See example above for usage. <a id="configurations"></a> ## Configurations (Clone, New strategies)`,
				},
				resource.Attribute{
					Name:        "configurations_file",
					Description: `(Optional) Describes path to S3 file containing description of configurations. [More Information](https://api.spotinst.com/elastigroup-for-aws/services-integrations/elastic-mapreduce/import-an-emr-cluster/advanced/)`,
				},
				resource.Attribute{
					Name:        "bucket",
					Description: `(Required) S3 Bucket name for configurations.`,
				},
				resource.Attribute{
					Name:        "steps_file",
					Description: `(Optional) Steps from S3.`,
				},
				resource.Attribute{
					Name:        "bucket",
					Description: `(Required) S3 Bucket name for steps.`,
				},
				resource.Attribute{
					Name:        "bootstrap_actions_file",
					Description: `(Optional) Describes path to S3 file containing description of bootstrap actions. [More Information](https://api.spotinst.com/elastigroup-for-aws/services-integrations/elastic-mapreduce/import-an-emr-cluster/advanced/)`,
				},
				resource.Attribute{
					Name:        "bucket",
					Description: `(Required) S3 Bucket name for bootstrap actions.`,
				},
				resource.Attribute{
					Name:        "policy_name",
					Description: `(Required) The name of the policy.`,
				},
				resource.Attribute{
					Name:        "metric_name",
					Description: `(Required) The name of the metric, with or without spaces.`,
				},
				resource.Attribute{
					Name:        "statistic",
					Description: `(Required) The metric statistics to return. For information about specific statistics go to [Statistics](http://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/index.html?CHAP_TerminologyandKeyConcepts.html#Statistic) in the Amazon CloudWatch Developer Guide.`,
				},
				resource.Attribute{
					Name:        "unit",
					Description: `(Required) The unit for the metric.`,
				},
				resource.Attribute{
					Name:        "threshold",
					Description: `(Required) The value against which the specified statistic is compared.`,
				},
				resource.Attribute{
					Name:        "adjustment",
					Description: `(Optional) The number of instances to add/remove to/from the target capacity when scale is needed.`,
				},
				resource.Attribute{
					Name:        "min_target_capacity",
					Description: `(Optional) Min target capacity for scale up.`,
				},
				resource.Attribute{
					Name:        "max_target_capacity",
					Description: `(Optional) Max target capacity for scale down.`,
				},
				resource.Attribute{
					Name:        "namespace",
					Description: `(Required) The namespace for the metric.`,
				},
				resource.Attribute{
					Name:        "operator",
					Description: `(Required) The operator to use. Allowed values are : 'gt', 'gte', 'lt' , 'lte'.`,
				},
				resource.Attribute{
					Name:        "evaluation_periods",
					Description: `(Required) The number of periods over which data is compared to the specified threshold.`,
				},
				resource.Attribute{
					Name:        "period",
					Description: `(Required) The granularity, in seconds, of the returned datapoints. Period must be at least 60 seconds and must be a multiple of 60.`,
				},
				resource.Attribute{
					Name:        "cooldown",
					Description: `(Required) The amount of time, in seconds, after a scaling activity completes and before the next scaling activity can start.`,
				},
				resource.Attribute{
					Name:        "dimensions",
					Description: `(Optional) A mapping of dimensions describing qualities of the metric.`,
				},
				resource.Attribute{
					Name:        "minimum",
					Description: `(Optional) The minimum to set when scale is needed.`,
				},
				resource.Attribute{
					Name:        "maximum",
					Description: `(Optional) The maximum to set when scale is needed.`,
				},
				resource.Attribute{
					Name:        "target",
					Description: `(Optional) The number of instances to set when scale is needed.`,
				},
				resource.Attribute{
					Name:        "action_type",
					Description: `(Required) The type of action to perform. Allowed values are : 'adjustment', 'setMinTarget', 'setMaxTarget', 'updateCapacity', 'percentageAdjustment' <a id="scheduled-task"></a> ## Scheduled Tasks`,
				},
				resource.Attribute{
					Name:        "scheduled_task",
					Description: `(Optional) An array of scheduled tasks.`,
				},
				resource.Attribute{
					Name:        "is_enabled",
					Description: `(Optional) Enable/Disable the specified scheduling task.`,
				},
				resource.Attribute{
					Name:        "task_type",
					Description: `(Required) The type of task to be scheduled. Valid values: ` + "`" + `setCapacity` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "instance_group_type",
					Description: `(Required) Select the EMR instance groups to execute the scheduled task on. Valid values: ` + "`" + `task` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "cron",
					Description: `(Required) A cron expression representing the schedule for the task.`,
				},
				resource.Attribute{
					Name:        "desired_capacity",
					Description: `(Optional) New desired capacity for the elastigroup.`,
				},
				resource.Attribute{
					Name:        "min_capacity",
					Description: `(Optional) New min capacity for the elastigroup.`,
				},
				resource.Attribute{
					Name:        "max_capacity",
					Description: `(Optional) New max capacity for the elastigroup. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The scaler ID.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The scaler ID.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "multai_balancer",
			Category:         "Spotinst Resources",
			ShortDescription: `Provides a Spotinst Multai Balancer.`,
			Description:      ``,
			Keywords: []string{
				"multai",
				"balancer",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The balancer name. May contain only alphanumeric characters or hyphens, and must not begin or end with a hyphen.`,
				},
				resource.Attribute{
					Name:        "scheme",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "dns_cname_aliases",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "connection_timeouts",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "idle",
					Description: `(Optional) The idle timeout value, in seconds. (range: 1 - 3600).`,
				},
				resource.Attribute{
					Name:        "draining",
					Description: `(Optional) The time for the load balancer to keep connections alive before reporting the target as de-registered, in seconds (range: 1 - 3600).`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A list of key:value paired tags.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) The tag's key.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) The tag's value.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "multai_deployment",
			Category:         "Spotinst Resources",
			ShortDescription: `Provides a Spotinst Multai Deployment.`,
			Description:      ``,
			Keywords: []string{
				"multai",
				"deployment",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The deployment name.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "multai_listener",
			Category:         "Spotinst Resources",
			ShortDescription: `Provides a Spotinst Multai Listener.`,
			Description:      ``,
			Keywords: []string{
				"multai",
				"listener",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "balancer_id",
					Description: `(Required) The ID of the balancer.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Required) The protocol to allow connections to the load balancer.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Required) The port on which the load balancer is listening.`,
				},
				resource.Attribute{
					Name:        "tls_config",
					Description: `(Optional) Describes the TLSConfig configuration.`,
				},
				resource.Attribute{
					Name:        "min_version",
					Description: `(Required) MinVersion contains the minimum SSL/TLS version that is acceptable (1.0 is the minimum).`,
				},
				resource.Attribute{
					Name:        "max_version",
					Description: `(Required) MaxVersion contains the maximum SSL/TLS version that is acceptable.`,
				},
				resource.Attribute{
					Name:        "certificate_ids",
					Description: `(Optional) Contains one or more certificate chains to present to the other side of the connection.`,
				},
				resource.Attribute{
					Name:        "cipher_suites",
					Description: `(Optional) List of supported cipher suites. If cipherSuites is nil, TLS uses a list of suites supported by the implementation.`,
				},
				resource.Attribute{
					Name:        "prefer_server_cipher_suites",
					Description: `(Optional) Controls whether the server selects the client’s most preferred ciphersuite, or the server’s most preferred ciphersuite.`,
				},
				resource.Attribute{
					Name:        "session_tickets_disabled",
					Description: `(Optional) May be set to true to disable session ticket (resumption) support.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A list of key:value paired tags.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) The tag's key.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) The tag's value.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "multai_routing_rule",
			Category:         "Spotinst Resources",
			ShortDescription: `Provides a Spotinst Multai Balancer.`,
			Description:      ``,
			Keywords: []string{
				"multai",
				"routing",
				"rule",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "balancer_id",
					Description: `(Required) The ID of the balancer.`,
				},
				resource.Attribute{
					Name:        "listener_id",
					Description: `(Required) The ID of the listener.`,
				},
				resource.Attribute{
					Name:        "route",
					Description: `(Required) Route defines a simple language for matching HTTP requests and route the traffic accordingly. Route provides series of matchers that follow the syntax: Path matcher: — Path("/foo/bar") // trie-based PathRegexp(“/foo/.`,
				},
				resource.Attribute{
					Name:        "strategy",
					Description: `(Optional) Balancing strategy. Valid values: ` + "`" + `ROUNDROBIN` + "`" + `, ` + "`" + `RANDOM` + "`" + `, ` + "`" + `LEASTCONN` + "`" + `, ` + "`" + `IPHASH` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A list of key:value paired tags.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) The tag's key.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) The tag's value.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "multai_target",
			Category:         "Spotinst Resources",
			ShortDescription: `Provides a Spotinst Multai target.`,
			Description:      ``,
			Keywords: []string{
				"multai",
				"target",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "balancer_id",
					Description: `(Required) The ID of the balancer.`,
				},
				resource.Attribute{
					Name:        "target_set_id",
					Description: `(Required) The ID of the target set.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the Target . Must contain only alphanumeric characters or hyphens, and must not begin or end with a hyphen.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Required) The port the target will register to.`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `(Required) The address (IP or URL) of the targets to register`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `(Required) Defines how traffic is distributed between targets.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A list of key:value paired tags.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) The tag's key.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) The tag's value.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "multai_target_set",
			Category:         "Spotinst Resources",
			ShortDescription: `Provides a Spotinst Multai Target Set.`,
			Description:      ``,
			Keywords: []string{
				"multai",
				"target",
				"set",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the Target Set. Must contain only alphanumeric characters or hyphens, and must not begin or end with a hyphen.`,
				},
				resource.Attribute{
					Name:        "balancer_id",
					Description: `(Required) The id of the balancer.`,
				},
				resource.Attribute{
					Name:        "deployment_id",
					Description: `(Required) The id of the deployment.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Required) The protocol to allow connections to the target.`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `(Required) Defines how traffic is distributed between the Target Set.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Required) The protocol to allow connections to the target for the health check.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Required) The path to perform the health check.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Required) The port on which the load balancer is listening.`,
				},
				resource.Attribute{
					Name:        "interval",
					Description: `(Required) The interval for the health check.`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `(Required) The time out for the health check.`,
				},
				resource.Attribute{
					Name:        "healthy_threshold",
					Description: `(Required) Total number of allowed healthy Targets.`,
				},
				resource.Attribute{
					Name:        "unhealthy_threshold",
					Description: `(Required) Total number of allowed unhealthy Targets.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A list of key:value paired tags.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) The tag's key.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) The tag's value.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ocean_aws",
			Category:         "Spotinst Resources",
			ShortDescription: `Provides a Spotinst Ocean resource using AWS.`,
			Description:      ``,
			Keywords: []string{
				"ocean",
				"aws",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The cluster name.`,
				},
				resource.Attribute{
					Name:        "controller_id",
					Description: `(Required) The ocean cluster identifier. Example: ` + "`" + `ocean.k8s` + "`" + ``,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Required) The region the cluster will run in.`,
				},
				resource.Attribute{
					Name:        "max_size",
					Description: `(Optional, Default: ` + "`" + `1000` + "`" + `) The upper limit of instances the cluster can scale up to.`,
				},
				resource.Attribute{
					Name:        "min_size",
					Description: `(Optional) The lower limit of instances the cluster can scale down to.`,
				},
				resource.Attribute{
					Name:        "desired_capacity",
					Description: `(Optional) The number of instances to launch and maintain in the cluster.`,
				},
				resource.Attribute{
					Name:        "subnet_ids",
					Description: `(Required) A comma-separated list of subnet identifiers for the Ocean cluster. Subnet IDs should be configured with auto assign public ip. ` + "`" + `` + "`" + `` + "`" + `hcl name = "demo" controller_id = "fakeClusterId" region = "us-west-2" max_size = 2 min_size = 1 desired_capacity = 2 ` + "`" + `` + "`" + `` + "`" + ``,
				},
				resource.Attribute{
					Name:        "whitelist",
					Description: `(Optional) Instance types allowed in the Ocean cluster. Cannot be configured if ` + "`" + `blacklist` + "`" + ` is configured.`,
				},
				resource.Attribute{
					Name:        "blacklist",
					Description: `(Optional) Instance types not allowed in the Ocean cluster. Cannot be configured if ` + "`" + `whitelist` + "`" + ` is configured. ` + "`" + `` + "`" + `` + "`" + `hcl whitelist = ["t1.micro", "m1.small"] // blacklist = ["t1.micro", "m1.small"] ` + "`" + `` + "`" + `` + "`" + ``,
				},
				resource.Attribute{
					Name:        "user_data",
					Description: `(Optional) Base64-encoded MIME user data to make available to the instances.`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `(Required) ID of the image used to launch the instances.`,
				},
				resource.Attribute{
					Name:        "security_groups",
					Description: `(Required) One or more security group ids.`,
				},
				resource.Attribute{
					Name:        "key_name",
					Description: `(Optional) The key pair to attach the instances.`,
				},
				resource.Attribute{
					Name:        "iam_instance_profile",
					Description: `(Optional) The instance profile iam role.`,
				},
				resource.Attribute{
					Name:        "associate_public_ip_address",
					Description: `(Optional, Default: ` + "`" + `false` + "`" + `) Configure public IP address allocation.`,
				},
				resource.Attribute{
					Name:        "root_volume_size",
					Description: `(Optional) The size (in Gb) to allocate for the root volume. Minimum ` + "`" + `20` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "monitoring",
					Description: `(Optional) Enable detailed monitoring for cluster. Flag will enable Cloud Watch detailed detailed monitoring (one minute increments). Note: there are additional hourly costs for this service based on the region used.`,
				},
				resource.Attribute{
					Name:        "ebs_optimized",
					Description: `(Optional) Enable EBS optimized for cluster. Flag will enable optimized capacity for high bandwidth connectivity to the EB service for non EBS optimized instance types. For instances that are EBS optimized this flag will be ignored.`,
				},
				resource.Attribute{
					Name:        "load_balancers",
					Description: `(Optional) - Array of load balancer objects to add to ocean cluster`,
				},
				resource.Attribute{
					Name:        "arn",
					Description: `(Optional) Required if type is set to TARGET_GROUP`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Required if type is set to CLASSIC`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) Can be set to CLASSIC or TARGET_GROUP ` + "`" + `` + "`" + `` + "`" + `hcl image_id = "ami-79826301" security_groups = ["sg-042d658b3ee907848"] key_name = "fake key" user_data = "echo hello world" iam_instance_profile = "iam-profile" root_volume_size = 20 monitoring = true ebs_optimized = true associate_public_ip_address = true load_balancers = [ { arn = "arn:aws:elasticloadbalancing:us-west-2:fake-arn" type = "TARGET_GROUP" }, { name = "AntonK" type = "CLASSIC" } ] ` + "`" + `` + "`" + `` + "`" + ``,
				},
				resource.Attribute{
					Name:        "fallback_to_ondemand",
					Description: `(Optional, Default: ` + "`" + `true` + "`" + `) If not Spot instance markets are available, enable Ocean to launch On-Demand instances instead.`,
				},
				resource.Attribute{
					Name:        "spot_percentage",
					Description: `(Optional, Default: ` + "`" + `100` + "`" + `) The percentage of Spot instances the cluster should maintain. Min 0, max 100.`,
				},
				resource.Attribute{
					Name:        "utilize_reserved_instances",
					Description: `(Optional, Default ` + "`" + `false` + "`" + `) If Reserved instances exist, OCean will utilize them before launching Spot instances. ` + "`" + `` + "`" + `` + "`" + `hcl fallback_to_ondemand = true spot_percentage = 100 utilize_reserved_instances = false ` + "`" + `` + "`" + `` + "`" + ``,
				},
				resource.Attribute{
					Name:        "autoscaler",
					Description: `(Optional) Describes the Ocean Kubernetes autoscaler.`,
				},
				resource.Attribute{
					Name:        "autoscale_is_enabled",
					Description: `(Optional, Default: ` + "`" + `true` + "`" + `) Enable the Ocean Kubernetes autoscaler.`,
				},
				resource.Attribute{
					Name:        "autoscale_is_auto_config",
					Description: `(Optional, Default: ` + "`" + `true` + "`" + `) Automatically configure and optimize headroom resources.`,
				},
				resource.Attribute{
					Name:        "autoscale_cooldown",
					Description: `(Optional, Default: ` + "`" + `null` + "`" + `) Cooldown period between scaling actions.`,
				},
				resource.Attribute{
					Name:        "autoscale_headroom",
					Description: `(Optional) Spare resource capacity management enabling fast assignment of Pods without waiting for new resources to launch.`,
				},
				resource.Attribute{
					Name:        "cpu_per_unit",
					Description: `(Optional) Optionally configure the number of CPUs to allocate the headroom. CPUs are denoted in millicores, where 1000 millicores = 1 vCPU.`,
				},
				resource.Attribute{
					Name:        "gpu_per_unit",
					Description: `(Optional) Optionally configure the number of GPUS to allocate the headroom.`,
				},
				resource.Attribute{
					Name:        "memory_per_unit",
					Description: `(Optional) Optionally configure the amount of memory (MB) to allocate the headroom.`,
				},
				resource.Attribute{
					Name:        "num_of_units",
					Description: `(Optional) The number of units to retain as headroom, where each unit has the defined headroom CPU and memory.`,
				},
				resource.Attribute{
					Name:        "autoscale_down",
					Description: `(Optional) Auto Scaling scale down operations.`,
				},
				resource.Attribute{
					Name:        "evaluation_periods",
					Description: `(Optional, Default: ` + "`" + `null` + "`" + `) The number of evaluation periods that should accumulate before a scale down action takes place.`,
				},
				resource.Attribute{
					Name:        "resource_limits",
					Description: `(Optional) Optionally set upper and lower bounds on the resource usage of the cluster.`,
				},
				resource.Attribute{
					Name:        "max_vcpu",
					Description: `(Optional) The maximum cpu in vCPU units that can be allocated to the cluster.`,
				},
				resource.Attribute{
					Name:        "max_memory_gib",
					Description: `(Optional) The maximum memory in GiB units that can be allocated to the cluster. ` + "`" + `` + "`" + `` + "`" + `hcl autoscaler = { autoscale_is_enabled = false autoscale_is_auto_config = false autoscale_cooldown = 300 autoscale_headroom = { cpu_per_unit = 1024 gpu_per_unit = 1 memory_per_unit = 512 num_of_units = 2 } autoscale_down = { evaluation_periods = 300 } resource_limits = { max_vcpu = 1024 max_memory_gib = 20 } } ` + "`" + `` + "`" + `` + "`" + ``,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Optionally adds tags to instances launched in an Ocean cluster.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Optional) The tag key.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Optional) The tag value. ` + "`" + `` + "`" + `` + "`" + `hcl tags = [{ key = "fakeKey" value = "fakeValue" }] ` + "`" + `` + "`" + `` + "`" + ` <a id="update-policy"></a> ## Update Policy`,
				},
				resource.Attribute{
					Name:        "update_policy",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "should_roll",
					Description: `(Required) Enables the roll.`,
				},
				resource.Attribute{
					Name:        "roll_config",
					Description: `(Required) While used, you can control whether the group should perform a deployment after an update to the configuration.`,
				},
				resource.Attribute{
					Name:        "batch_size_percentage",
					Description: `(Required) Sets the percentage of the instances to deploy in each batch. ` + "`" + `` + "`" + `` + "`" + `hcl update_policy = { should_roll = false roll_config = { batch_size_percentage = 33 } } ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ocean_aws_launch_spec",
			Category:         "Spotinst Resources",
			ShortDescription: `Provides a Spotinst Ocean Launch Spec resource using AWS.`,
			Description:      ``,
			Keywords: []string{
				"ocean",
				"aws",
				"launch",
				"spec",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ocean_id",
					Description: `(Required) The ocean cluster you wish to`,
				},
				resource.Attribute{
					Name:        "user_data",
					Description: `(Optional) Base64-encoded MIME user data to make available to the instances.`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `(Optional) ID of the image used to launch the instances.`,
				},
				resource.Attribute{
					Name:        "iam_instance_profile",
					Description: `(Optional) The ARN or name of an IAM instance profile to associate with launched instances.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) Optionally adds labels to instances launched in an Ocean cluster.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) The tag key.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) The tag value.`,
				},
				resource.Attribute{
					Name:        "taints",
					Description: `(Optional) Optionally adds labels to instances launched in an Ocean cluster.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) The tag key.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) The tag value.`,
				},
				resource.Attribute{
					Name:        "effect",
					Description: `(Required) The effect of the taint. Valid values: ` + "`" + `"NoSchedule"` + "`" + `, ` + "`" + `"PreferNoSchedule"` + "`" + `, ` + "`" + `"NoExecute"` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ocean_gke",
			Category:         "Spotinst Resources",
			ShortDescription: `Provides a Spotinst Ocean resource using gke.`,
			Description:      ``,
			Keywords: []string{
				"ocean",
				"gke",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The cluster name.`,
				},
				resource.Attribute{
					Name:        "controller_id",
					Description: `(Required) The ocean cluster identifier. Example: ` + "`" + `ocean.k8s` + "`" + ``,
				},
				resource.Attribute{
					Name:        "cluster_name",
					Description: `(Required) The GKE cluster name.`,
				},
				resource.Attribute{
					Name:        "master_location",
					Description: `(Required) The zone the master cluster is located in.`,
				},
				resource.Attribute{
					Name:        "subnet_name",
					Description: `(Required) Subnet identifier for the Ocean cluster.`,
				},
				resource.Attribute{
					Name:        "availability_zones",
					Description: `(Required) List of availability zones available to the cluster.`,
				},
				resource.Attribute{
					Name:        "whitelist",
					Description: `(Optional) Instance types allowed in the Ocean cluster.`,
				},
				resource.Attribute{
					Name:        "max_size",
					Description: `(Optional, Default: ` + "`" + `1000` + "`" + `) The upper limit of instances the cluster can scale up to.`,
				},
				resource.Attribute{
					Name:        "min_size",
					Description: `(Optional) The lower limit of instances the cluster can scale down to.`,
				},
				resource.Attribute{
					Name:        "desired_capacity",
					Description: `(Optional) The number of instances to launch and maintain in the cluster. Usage: ` + "`" + `` + "`" + `` + "`" + `hcl name = "example-ocean-cluster-name" controller_id = "example-cluster-id" cluster_name = "example-cluster-name" master_location = "us-central1-a" subnet_name = "example-subnet-1" availability_zones = ["us-central1-a"] whitelist = ["n1-standard-1", "n1-standard-2"] max_size = 1000 min_size = 0 desired_capacity = 500 ` + "`" + `` + "`" + `` + "`" + ` <a id="launch-configuration"></a> ## Launch Configuration Note: label, metadata, and tag keys are required, and depend on your GKE cluster. Please modify the values to match your configuration. You may also add additional key/value pairs. This resource is intended to be used as part of a Module.`,
				},
				resource.Attribute{
					Name:        "source_image",
					Description: `(Optional) A source image used to create the disk. You can provide a private (custom) image, and Compute Engine will use the corresponding image from your project.`,
				},
				resource.Attribute{
					Name:        "service_account",
					Description: `(Optional) The email of the service account in which the group instances will be launched.`,
				},
				resource.Attribute{
					Name:        "root_volume_size_in_gb",
					Description: `(Optional) The size (in Gb) to allocate for the root volume. Minimum ` + "`" + `100` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ip_forwarding",
					Description: `(Optional) Enables the transfer IP packets from one network to another.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) Array of objects with key-value pairs.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Optional) Labels key.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Optional) Labels value.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `(Optional) Array of objects with key-value pairs.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Optional) Metadata key.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Optional) Metadata value.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags to mark created instances. Minimum 1. Usage: ` + "`" + `` + "`" + `` + "`" + `hcl source_image = "https://www.googleapis.com/compute/v1/projects/my-project/global/examples/example-image-1" service_account = "example-account@my-account.iam.gserviceaccount.com" root_volume_size_in_gb = 100 ip_forwarding = true labels = [{ key = "spotinst-gke-original-node-pool", value = "example-cluster-name__default-pool" }] metadata = [{ key = "cluster-name" value = "example-cluster" }] tags = ["gke-example-vpc-1234567-node"] ` + "`" + `` + "`" + `` + "`" + ` <a id="backend-services"></a> ## Backend Services`,
				},
				resource.Attribute{
					Name:        "backend_services",
					Description: `(Optional) Describes the backend service configurations.`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `(Required) The name of the backend service.`,
				},
				resource.Attribute{
					Name:        "location_type",
					Description: `(Optional) Sets which location the backend services will be active. Valid values: ` + "`" + `regional` + "`" + `, ` + "`" + `global` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "scheme",
					Description: `(Optional) Use when ` + "`" + `location_type` + "`" + ` is "regional". Set the traffic for the backend service to either between the instances in the vpc or to traffic from the internet. Valid values: ` + "`" + `INTERNAL` + "`" + `, ` + "`" + `EXTERNAL` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "named_port",
					Description: `(Optional) Describes a named port and a list of ports.`,
				},
				resource.Attribute{
					Name:        "port_name",
					Description: `(Required) The name of the port.`,
				},
				resource.Attribute{
					Name:        "ports",
					Description: `(Required) A list of ports. Usage: ` + "`" + `` + "`" + `` + "`" + `hcl backend_services = [{ service_name = "example-backend-service" location_type = "global" scheme = "INTERNAL" named_ports = { name = "http" ports = [80, 8080] } }] ` + "`" + `` + "`" + `` + "`" + ` <a id="autoscaler"></a> ## Autoscaler`,
				},
				resource.Attribute{
					Name:        "autoscaler",
					Description: `(Optional) Describes the Ocean Kubernetes autoscaler.`,
				},
				resource.Attribute{
					Name:        "autoscale_is_enabled",
					Description: `(Optional, Default: ` + "`" + `true` + "`" + `) Enable the Ocean Kubernetes autoscaler.`,
				},
				resource.Attribute{
					Name:        "autoscale_is_auto_config",
					Description: `(Optional, Default: ` + "`" + `true` + "`" + `) Automatically configure and optimize headroom resources.`,
				},
				resource.Attribute{
					Name:        "autoscale_cooldown",
					Description: `(Optional, Default: ` + "`" + `null` + "`" + `) Cooldown period between scaling actions.`,
				},
				resource.Attribute{
					Name:        "autoscale_headroom",
					Description: `(Optional) Spare resource capacity management enabling fast assignment of Pods without waiting for new resources to launch.`,
				},
				resource.Attribute{
					Name:        "cpu_per_unit",
					Description: `(Optional) Optionally configure the number of CPUs to allocate the headroom. CPUs are denoted in millicores, where 1000 millicores = 1 vCPU.`,
				},
				resource.Attribute{
					Name:        "gpu_per_unit",
					Description: `(Optional) Optionally configure the number of GPUS to allocate the headroom.`,
				},
				resource.Attribute{
					Name:        "memory_per_unit",
					Description: `(Optional) Optionally configure the amount of memory (MB) to allocate the headroom.`,
				},
				resource.Attribute{
					Name:        "num_of_units",
					Description: `(Optional) The number of units to retain as headroom, where each unit has the defined headroom CPU and memory.`,
				},
				resource.Attribute{
					Name:        "autoscale_down",
					Description: `(Optional) Auto Scaling scale down operations.`,
				},
				resource.Attribute{
					Name:        "evaluation_periods",
					Description: `(Optional, Default: ` + "`" + `null` + "`" + `) The number of evaluation periods that should accumulate before a scale down action takes place.`,
				},
				resource.Attribute{
					Name:        "resource_limits",
					Description: `(Optional) Optionally set upper and lower bounds on the resource usage of the cluster.`,
				},
				resource.Attribute{
					Name:        "max_vcpu",
					Description: `(Optional) The maximum cpu in vCPU units that can be allocated to the cluster.`,
				},
				resource.Attribute{
					Name:        "max_memory_gib",
					Description: `(Optional) The maximum memory in GiB units that can be allocated to the cluster. Usage: ` + "`" + `` + "`" + `` + "`" + `hcl autoscaler = { autoscale_is_enabled = false autoscale_is_auto_config = false autoscale_cooldown = 300 autoscale_headroom = { cpu_per_unit = 1024 gpu_per_unit = 1 memory_per_unit = 512 num_of_units = 2 } autoscale_down = { evaluation_periods = 300 } resource_limits = { max_vcpu = 1024 max_memory_gib = 20 } } ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "subscription",
			Category:         "Spotinst Resources",
			ShortDescription: `Provides a Spotinst subscription resource.`,
			Description:      ``,
			Keywords: []string{
				"subscription",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "resource_id",
					Description: `(Required) Spotinst Resource ID (Elastigroup ID).`,
				},
				resource.Attribute{
					Name:        "event_type",
					Description: `(Required) The event to send the notification when triggered. Valid values: ` + "`" + `"AWS_EC2_INSTANCE_TERMINATE"` + "`" + `, ` + "`" + `"AWS_EC2_INSTANCE_TERMINATED"` + "`" + `, ` + "`" + `"AWS_EC2_INSTANCE_LAUNCH"` + "`" + `, ` + "`" + `"AWS_EC2_INSTANCE_UNHEALTHY_IN_ELB"` + "`" + `, ` + "`" + `"GROUP_ROLL_FAILED"` + "`" + `, ` + "`" + `"GROUP_ROLL_FINISHED"` + "`" + `, ` + "`" + `"CANT_SCALE_UP_GROUP_MAX_CAPACITY"` + "`" + `, ` + "`" + `"GROUP_UPDATED"` + "`" + `, ` + "`" + `"AWS_EC2_CANT_SPIN_OD"` + "`" + `, ` + "`" + `"AWS_EMR_PROVISION_TIMEOUT"` + "`" + `, ` + "`" + `"AWS_EC2_INSTANCE_READY_SIGNAL_TIMEOUT"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Required) The protocol to send the notification. Valid values: ` + "`" + `"http"` + "`" + `, ` + "`" + `"https"` + "`" + `, ` + "`" + `"email"` + "`" + `, ` + "`" + `"email-json"` + "`" + `, ` + "`" + `"aws-sns"` + "`" + `, ` + "`" + `"web"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "endpoint",
					Description: `(Required) The endpoint the notification will be sent to: url in case of ` + "`" + `"http"` + "`" + `/` + "`" + `"https"` + "`" + `, email address in case of ` + "`" + `"email"` + "`" + `/` + "`" + `"email-json"` + "`" + `, sns-topic-arn in case of ` + "`" + `"aws-sns"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "format",
					Description: `(Optional) The format of the notification content (JSON Format - Key+Value). Valid values: ` + "`" + `"%instance-id%"` + "`" + `, ` + "`" + `"%event%"` + "`" + `, ` + "`" + `"%resource-id%"` + "`" + `, ` + "`" + `"%resource-name%"` + "`" + `. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The subscription ID.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The subscription ID.`,
				},
			},
		},
	}

	resourcesMap = map[string]int{

		"elastigroup_aws":           0,
		"elastigroup_aws_beanstalk": 1,
		"elastigroup_azure":         2,
		"elastigroup_gcp":           3,
		"elastigroup_gke":           4,
		"mrscaler_aws":              5,
		"multai_balancer":           6,
		"multai_deployment":         7,
		"multai_listener":           8,
		"multai_routing_rule":       9,
		"multai_target":             10,
		"multai_target_set":         11,
		"ocean_aws":                 12,
		"ocean_aws_launch_spec":     13,
		"ocean_gke":                 14,
		"subscription":              15,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
