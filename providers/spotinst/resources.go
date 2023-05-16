package spotinst

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "spotinst_data_integration",
			Category:         "Data Integration",
			ShortDescription: `Manages an Data Integration resource.`,
			Description:      ``,
			Keywords: []string{
				"data",
				"integration",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "status",
					Description: `(Optional, only when update) Determines if this data integration is on or off. Valid values: ` + "`" + `"enabled"` + "`" + `, ` + "`" + `"disabled"` + "`" + ``,
				},
				resource.Attribute{
					Name:        "s3",
					Description: `(Required) When vendor value is s3, the following fields are included:`,
				},
				resource.Attribute{
					Name:        "bucketName",
					Description: `(Required) The name of the bucket to use. Your spot IAM Role policy needs to include s3:putObject permissions for this bucket. Can't be null.`,
				},
				resource.Attribute{
					Name:        "subdir",
					Description: `(Optional) The subdirectory in which your files will be stored within the bucket. Adds the prefix subdir/ to new objects' keys. Can't be null or contain '/'. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The Spotinst Data Integration ID.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The Spotinst Data Integration ID.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "spotinst_elastigroup_aws",
			Category:         "Elastigroup",
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
					Description: `(Optional, Required if using scaling policies) The maximum number of instances the group should have at any time.`,
				},
				resource.Attribute{
					Name:        "min_size",
					Description: `(Optional, Required if using scaling policies) The minimum number of instances the group should have at any time.`,
				},
				resource.Attribute{
					Name:        "desired_capacity",
					Description: `(Required) The desired number of instances the group should have at any time.`,
				},
				resource.Attribute{
					Name:        "capacity_unit",
					Description: `(Optional, Default: ` + "`" + `instance` + "`" + `) The capacity unit to launch instances by. If not specified, when choosing the weight unit, each instance will weight as the number of its vCPUs. Valid values: ` + "`" + `instance` + "`" + `, ` + "`" + `weight` + "`" + `.`,
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
					Name:        "images",
					Description: `(Optional) An array of image objects. Note: Elastigroup can be configured with either imageId or images, but not both.`,
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
					Description: `(Optional, Default: "default") Enable dedicated tenancy. Note: There is a flat hourly fee for each region in which dedicated tenancy is used. Valid values: "default", "dedicated" .`,
				},
				resource.Attribute{
					Name:        "metadata_options",
					Description: `(Optional) Data that used to configure or manage the running instances:`,
				},
				resource.Attribute{
					Name:        "http_tokens",
					Description: `(Required) The state of token usage for your instance metadata requests. Valid values: ` + "`" + `optional` + "`" + ` or ` + "`" + `required` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "http_put_response_hop_limit",
					Description: `(Optional, Default: ` + "`" + `1` + "`" + `) The desired HTTP PUT response hop limit for instance metadata requests. The larger the number, the further instance metadata requests can travel. Valid values: Integers from ` + "`" + `1` + "`" + ` to ` + "`" + `64` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "instance_metadata_tags",
					Description: `(Optional) Indicates whether access to instance tags from the instance metadata is enabled or disabled. Can’t be null.`,
				},
				resource.Attribute{
					Name:        "cpu_options",
					Description: `(Optional) The CPU options for the instances that are launched within the group:`,
				},
				resource.Attribute{
					Name:        "threads_per_core",
					Description: `(Required) The ability to define the number of threads per core in instances that allow this.`,
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
					Description: `(Required, Default: ` + "`" + `balanced` + "`" + `) Select a prediction strategy. Valid values: ` + "`" + `balanced` + "`" + `, ` + "`" + `costOriented` + "`" + `, ` + "`" + `equalAzDistribution` + "`" + `, ` + "`" + `availabilityOriented` + "`" + `. You can read more in our documentation.`,
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
					Name:        "minimum_instance_lifetime",
					Description: `(Optional) Defines the preferred minimum instance lifetime in hours. Markets which comply with this preference will be prioritized. Optional values: 1, 3, 6, 12, 24.`,
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
					Description: `(Optional) Specify a list of time windows for to execute revertToSpot strategy. Time window format: ` + "`" + `ddd:hh:mm-ddd:hh:mm` + "`" + `. Example: ` + "`" + `Mon:03:00-Wed:02:30` + "`" + ``,
				},
				resource.Attribute{
					Name:        "resource_tag_specification",
					Description: `(Optional) User will specify which resources should be tagged with group tags.`,
				},
				resource.Attribute{
					Name:        "should_tag_snapshots",
					Description: `(Optional) Tag specification for Snapshot resources.`,
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
					Description: `(Required) ID of Multai Load Balancer. Usage: ` + "`" + `` + "`" + `` + "`" + `hcl elastic_load_balancers = ["bal5", "bal2"] target_group_arns = ["tg-arn"] multai_target_sets { target_set_id = "ts-123", balancer_id = "bal-123" } multai_target_sets { target_set_id = "ts-234", balancer_id = "bal-234" } ` + "`" + `` + "`" + `` + "`" + ` <a id="signal"></a> ## Signals Each ` + "`" + `signal` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the signal defined for the group. Valid Values: ` + "`" + `"INSTANCE_READY"` + "`" + `, ` + "`" + `"INSTANCE_READY_TO_SHUTDOWN"` + "`" + ``,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `(Optional) The signals defined timeout- default is 40 minutes (1800 seconds). Usage: ` + "`" + `` + "`" + `` + "`" + `hcl signal { name = "INSTANCE_READY_TO_SHUTDOWN" timeout = 100 } ` + "`" + `` + "`" + `` + "`" + ` <a id="scheduled-task"></a> ## Scheduled Tasks Each ` + "`" + `scheduled_task` + "`" + ` supports the following:`,
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
					Description: `(Optional; Format: ISO 8601; Time Standard: UTC time) Set a start time for one time tasks.`,
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
					Description: `(Optional; Min 1) The percentage of instances to add or remove. Usage: ` + "`" + `` + "`" + `` + "`" + `hcl scheduled_task { task_type = "backup_ami" cron_expression = "" start_time = "1970-01-01T01:00:00Z" frequency = "hourly" scale_target_capacity = 5 scale_min_capacity = 0 scale_max_capacity = 10 is_enabled = false target_capacity = 5 min_capacity = 0 max_capacity = 10 batch_size_percentage = 33 grace_period = 300 } ` + "`" + `` + "`" + `` + "`" + ` <a id="scaling-policy"></a> ## Scaling Policies ` + "`" + `scaling_up_policy` + "`" + ` supports the following:`,
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
					Description: `(Optional, Default: ` + "`" + `"average"` + "`" + `) The metric statistics to return. For information about specific statistics go to [Statistics](http://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/index.html?CHAP_TerminologyandKeyConcepts.html#Statistic) in the Amazon CloudWatch Developer Guide.`,
				},
				resource.Attribute{
					Name:        "unit",
					Description: `(Required) The unit for the alarm's associated metric. Valid values: ` + "`" + `"percent` + "`" + `, ` + "`" + `"seconds"` + "`" + `, ` + "`" + `"microseconds"` + "`" + `, ` + "`" + `"milliseconds"` + "`" + `, ` + "`" + `"bytes"` + "`" + `, ` + "`" + `"kilobytes"` + "`" + `, ` + "`" + `"megabytes"` + "`" + `, ` + "`" + `"gigabytes"` + "`" + `, ` + "`" + `"terabytes"` + "`" + `, ` + "`" + `"bits"` + "`" + `, ` + "`" + `"kilobits"` + "`" + `, ` + "`" + `"megabits"` + "`" + `, ` + "`" + `"gigabits"` + "`" + `, ` + "`" + `"terabits"` + "`" + `, ` + "`" + `"count"` + "`" + `, ` + "`" + `"bytes/second"` + "`" + `, ` + "`" + `"kilobytes/second"` + "`" + `, ` + "`" + `"megabytes/second"` + "`" + `, ` + "`" + `"gigabytes/second"` + "`" + `, ` + "`" + `"terabytes/second"` + "`" + `, ` + "`" + `"bits/second"` + "`" + `, ` + "`" + `"kilobits/second"` + "`" + `, ` + "`" + `"megabits/second"` + "`" + `, ` + "`" + `"gigabits/second"` + "`" + `, ` + "`" + `"terabits/second"` + "`" + `, ` + "`" + `"count/second"` + "`" + `, ` + "`" + `"none"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "threshold",
					Description: `(Required) The value against which the specified statistic is compared. If a ` + "`" + `step_adjustment` + "`" + ` object is defined, then it cannot be specified.`,
				},
				resource.Attribute{
					Name:        "action_type",
					Description: `(Optional; if not using ` + "`" + `min_target_capacity` + "`" + ` or ` + "`" + `max_target_capacity` + "`" + `) The type of action to perform for scaling. Valid values: ` + "`" + `"adjustment"` + "`" + `, ` + "`" + `"percentageAdjustment"` + "`" + `, ` + "`" + `"setMaxTarget"` + "`" + `, ` + "`" + `"setMinTarget"` + "`" + `, ` + "`" + `"updateCapacity"` + "`" + `. If a ` + "`" + `step_adjustment` + "`" + ` object is defined, then it cannot be specified.`,
				},
				resource.Attribute{
					Name:        "namespace",
					Description: `(Required) The namespace for the alarm's associated metric.`,
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
					Name:        "step_adjustment",
					Description: `(Optional) The list of steps to define actions to take based on different thresholds. When set, policy-level ` + "`" + `threshold` + "`" + ` and ` + "`" + `action_type` + "`" + ` cannot be specified.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Required) The action to take when scale up according to step's threshold is needed.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) The type of the action to take when scale up is needed. Valid types: ` + "`" + `"adjustment"` + "`" + `, ` + "`" + `"updateCapacity"` + "`" + `, ` + "`" + `"setMinTarget"` + "`" + `, ` + "`" + `"percentageAdjustment"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "adjustment",
					Description: `(Optional) The number/percentage associated with the specified adjustment type. Required if using ` + "`" + `"adjustment"` + "`" + ` or ` + "`" + `"percentageAdjustment"` + "`" + ` as action type.`,
				},
				resource.Attribute{
					Name:        "maximum",
					Description: `(Optional) The upper limit number of instances that you can scale up to. Required if using ` + "`" + `"updateCapacity"` + "`" + ` as action type and neither ` + "`" + `"target"` + "`" + ` nor ` + "`" + `"minimum"` + "`" + ` are not defined.`,
				},
				resource.Attribute{
					Name:        "minimum",
					Description: `(Optional) The lower limit number of instances that you can scale down to. Required if using ` + "`" + `"updateCapacity"` + "`" + ` as action type and neither ` + "`" + `"target"` + "`" + ` nor ` + "`" + `"maximum"` + "`" + ` are not defined.`,
				},
				resource.Attribute{
					Name:        "min_target_capacity",
					Description: `(Optional) The desired target capacity of a group. Required if using ` + "`" + `"setMinTarget"` + "`" + ` as action type`,
				},
				resource.Attribute{
					Name:        "target",
					Description: `(Optional) The desired number of instances. Required if using ` + "`" + `"updateCapacity"` + "`" + ` as action type and neither ` + "`" + `"minimum"` + "`" + ` nor ` + "`" + `"maximum"` + "`" + ` are not defined.`,
				},
				resource.Attribute{
					Name:        "threshold",
					Description: `(Required) The value against which the specified statistic is compared in order to determine if a step should be applied. If you do not specify an action type, you can only use – ` + "`" + `adjustment` + "`" + `, ` + "`" + `minTargetCapacity` + "`" + `, ` + "`" + `maxTargetCapacity` + "`" + `. While using action_type, please also set the following: When using ` + "`" + `adjustment` + "`" + ` – set the field ` + "`" + `adjustment` + "`" + ` When using ` + "`" + `setMinTarget` + "`" + ` – set the field ` + "`" + `min_target_capacity` + "`" + ` When using ` + "`" + `updateCapacity` + "`" + ` – set the fields ` + "`" + `minimum` + "`" + `, ` + "`" + `maximum` + "`" + `, and ` + "`" + `target` + "`" + ``,
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
					Name:        "minimum",
					Description: `(Optional; if using ` + "`" + `updateCapacity` + "`" + `) The minimal number of instances to have in the group.`,
				},
				resource.Attribute{
					Name:        "maximum",
					Description: `(Optional; if using ` + "`" + `updateCapacity` + "`" + `) The maximal number of instances to have in the group.`,
				},
				resource.Attribute{
					Name:        "target",
					Description: `(Optional; if using ` + "`" + `updateCapacity` + "`" + `) The target number of instances to have in the group. ` + "`" + `scaling_down_policy` + "`" + ` supports the following:`,
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
					Description: `(Optional, Default: ` + "`" + `"average"` + "`" + `) The metric statistics to return. For information about specific statistics go to [Statistics](http://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/index.html?CHAP_TerminologyandKeyConcepts.html#Statistic) in the Amazon CloudWatch Developer Guide.`,
				},
				resource.Attribute{
					Name:        "unit",
					Description: `(Required) The unit for the alarm's associated metric. Valid values: ` + "`" + `"percent` + "`" + `, ` + "`" + `"seconds"` + "`" + `, ` + "`" + `"microseconds"` + "`" + `, ` + "`" + `"milliseconds"` + "`" + `, ` + "`" + `"bytes"` + "`" + `, ` + "`" + `"kilobytes"` + "`" + `, ` + "`" + `"megabytes"` + "`" + `, ` + "`" + `"gigabytes"` + "`" + `, ` + "`" + `"terabytes"` + "`" + `, ` + "`" + `"bits"` + "`" + `, ` + "`" + `"kilobits"` + "`" + `, ` + "`" + `"megabits"` + "`" + `, ` + "`" + `"gigabits"` + "`" + `, ` + "`" + `"terabits"` + "`" + `, ` + "`" + `"count"` + "`" + `, ` + "`" + `"bytes/second"` + "`" + `, ` + "`" + `"kilobytes/second"` + "`" + `, ` + "`" + `"megabytes/second"` + "`" + `, ` + "`" + `"gigabytes/second"` + "`" + `, ` + "`" + `"terabytes/second"` + "`" + `, ` + "`" + `"bits/second"` + "`" + `, ` + "`" + `"kilobits/second"` + "`" + `, ` + "`" + `"megabits/second"` + "`" + `, ` + "`" + `"gigabits/second"` + "`" + `, ` + "`" + `"terabits/second"` + "`" + `, ` + "`" + `"count/second"` + "`" + `, ` + "`" + `"none"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "threshold",
					Description: `(Required) The value against which the specified statistic is compared. If a ` + "`" + `step_adjustment` + "`" + ` object is defined, then it cannot be specified.`,
				},
				resource.Attribute{
					Name:        "action_type",
					Description: `(Optional; if not using ` + "`" + `min_target_capacity` + "`" + ` or ` + "`" + `max_target_capacity` + "`" + `) The type of action to perform for scaling. Valid values: ` + "`" + `"adjustment"` + "`" + `, ` + "`" + `"percentageAdjustment"` + "`" + `, ` + "`" + `"setMaxTarget"` + "`" + `, ` + "`" + `"setMinTarget"` + "`" + `, ` + "`" + `"updateCapacity"` + "`" + `. If a ` + "`" + `step_adjustment` + "`" + ` object is defined, then it cannot be specified.`,
				},
				resource.Attribute{
					Name:        "namespace",
					Description: `(Required) The namespace for the alarm's associated metric.`,
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
					Name:        "step_adjustment",
					Description: `(Optional) The list of steps to define actions to take based on different thresholds. When set, policy-level ` + "`" + `threshold` + "`" + ` and ` + "`" + `action_type` + "`" + ` cannot be specified.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Required) The action to take when scale up according to step's threshold is needed.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) The type of the action to take when scale up is needed. Valid types: ` + "`" + `"adjustment"` + "`" + `, ` + "`" + `"updateCapacity"` + "`" + `, ` + "`" + `"setMaxTarget"` + "`" + `, ` + "`" + `"percentageAdjustment"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "adjustment",
					Description: `(Optional) The number/percentage associated with the specified adjustment type. Required if using ` + "`" + `"adjustment"` + "`" + ` or ` + "`" + `"percentageAdjustment"` + "`" + ` as action type.`,
				},
				resource.Attribute{
					Name:        "maximum",
					Description: `(Optional) The upper limit number of instances that you can scale up to. Required if using ` + "`" + `"updateCapacity"` + "`" + ` as action type and neither ` + "`" + `"target"` + "`" + ` nor ` + "`" + `"minimum"` + "`" + ` are not defined.`,
				},
				resource.Attribute{
					Name:        "minimum",
					Description: `(Optional) The lower limit number of instances that you can scale down to. Required if using ` + "`" + `"updateCapacity"` + "`" + ` as action type and neither ` + "`" + `"target"` + "`" + ` nor ` + "`" + `"maximum"` + "`" + ` are not defined.`,
				},
				resource.Attribute{
					Name:        "max_target_capacity",
					Description: `(Optional) The desired target capacity of a group. Required if using ` + "`" + `"setMaxTarget"` + "`" + ` as action type`,
				},
				resource.Attribute{
					Name:        "target",
					Description: `(Optional) The desired number of instances. Required if using ` + "`" + `"updateCapacity"` + "`" + ` as action type and neither ` + "`" + `"minimum"` + "`" + ` nor ` + "`" + `"maximum"` + "`" + ` are not defined.`,
				},
				resource.Attribute{
					Name:        "threshold",
					Description: `(Required) The value against which the specified statistic is compared in order to determine if a step should be applied. If you do not specify an action type, you can only use – ` + "`" + `adjustment` + "`" + `, ` + "`" + `minTargetCapacity` + "`" + `, ` + "`" + `maxTargetCapacity` + "`" + `. While using action_type, please also set the following: When using ` + "`" + `adjustment` + "`" + ` – set the field ` + "`" + `adjustment` + "`" + ` When using ` + "`" + `updateCapacity` + "`" + ` – set the fields ` + "`" + `minimum` + "`" + `, ` + "`" + `maximum` + "`" + `, and ` + "`" + `target` + "`" + ``,
				},
				resource.Attribute{
					Name:        "adjustment",
					Description: `(Optional; if not using ` + "`" + `min_target_capacity` + "`" + ` or ` + "`" + `max_target_capacity` + "`" + `;) The number of instances to add/remove to/from the target capacity when scale is needed. Can be used as advanced expression for scaling of instances to add/remove to/from the target capacity when scale is needed. You can see more information here: Advanced expression. Example value: ` + "`" + `"MAX(currCapacity / 5, value`,
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
					Description: `(Optional; if using ` + "`" + `updateCapacity` + "`" + `) The target number of instances to have in the group. ` + "`" + `scaling_target_policy` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "policy_name",
					Description: `(Required) String, the name of the policy.`,
				},
				resource.Attribute{
					Name:        "metric_name",
					Description: `(Required) String, the name of the metric, with or without spaces.`,
				},
				resource.Attribute{
					Name:        "statistic",
					Description: `(Optional, Default: ` + "`" + `"average"` + "`" + `) String, the metric statistics to return. For information about specific statistics go to [Statistics](http://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/index.html?CHAP_TerminologyandKeyConcepts.html#Statistic) in the Amazon CloudWatch Developer Guide.`,
				},
				resource.Attribute{
					Name:        "unit",
					Description: `(Required) String, tThe unit for the alarm's associated metric. Valid values: ` + "`" + `"percent` + "`" + `, ` + "`" + `"seconds"` + "`" + `, ` + "`" + `"microseconds"` + "`" + `, ` + "`" + `"milliseconds"` + "`" + `, ` + "`" + `"bytes"` + "`" + `, ` + "`" + `"kilobytes"` + "`" + `, ` + "`" + `"megabytes"` + "`" + `, ` + "`" + `"gigabytes"` + "`" + `, ` + "`" + `"terabytes"` + "`" + `, ` + "`" + `"bits"` + "`" + `, ` + "`" + `"kilobits"` + "`" + `, ` + "`" + `"megabits"` + "`" + `, ` + "`" + `"gigabits"` + "`" + `, ` + "`" + `"terabits"` + "`" + `, ` + "`" + `"count"` + "`" + `, ` + "`" + `"bytes/second"` + "`" + `, ` + "`" + `"kilobytes/second"` + "`" + `, ` + "`" + `"megabytes/second"` + "`" + `, ` + "`" + `"gigabytes/second"` + "`" + `, ` + "`" + `"terabytes/second"` + "`" + `, ` + "`" + `"bits/second"` + "`" + `, ` + "`" + `"kilobits/second"` + "`" + `, ` + "`" + `"megabits/second"` + "`" + `, ` + "`" + `"gigabits/second"` + "`" + `, ` + "`" + `"terabits/second"` + "`" + `, ` + "`" + `"count/second"` + "`" + `, ` + "`" + `"none"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "namespace",
					Description: `(Required) String, the namespace for the alarm's associated metric.`,
				},
				resource.Attribute{
					Name:        "cooldown",
					Description: `(Optional, Default: ` + "`" + `300` + "`" + `) Integer the amount of time, in seconds, after a scaling activity completes and before the next scaling activity can start. If this parameter is not specified, the default cooldown period for the group applies.`,
				},
				resource.Attribute{
					Name:        "source",
					Description: `(Optional) String, the source of the metric. Valid values: ` + "`" + `"cloudWatch"` + "`" + `, ` + "`" + `"spectrum"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "dimensions",
					Description: `(Optional) A list of dimensions describing qualities of the metric.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) String, the dimension name.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) String, the dimension value.`,
				},
				resource.Attribute{
					Name:        "target",
					Description: `(Optional; if using ` + "`" + `updateCapacity` + "`" + `) The target number of instances to have in the group.`,
				},
				resource.Attribute{
					Name:        "max_capacity_per_scale",
					Description: `(Optional) String, restrict the maximal number of instances which can be added in each scale-up action. ` + "`" + `scaling_target_policies` + "`" + ` support predictive scaling:`,
				},
				resource.Attribute{
					Name:        "predictive_mode",
					Description: `(Optional) Start a metric prediction process to determine the expected target metric value within the next two days. See [Predictive Autoscaling](https://api.spotinst.com/elastigroup-for-aws/concepts/scaling-concepts/predictive-autoscaling/) documentation for more info. Valid values: ` + "`" + `FORECAST_AND_SCALE` + "`" + `, ` + "`" + `FORECAST_ONLY` + "`" + `. Usage: ` + "`" + `` + "`" + `` + "`" + `hcl scaling_up_policy { policy_name = "policy-name" metric_name = "CPUUtilization" namespace = "AWS/EC2" source = "" statistic = "average" unit = "" cooldown = 60 dimensions { name = "name-1" value = "value-1" } threshold = 10 operator = "gt" evaluation_periods = 10 period = 60 step_adjustments { threshold = 50 action { type = "setMinTarget" min_target_capacity = "3" } } // === MIN TARGET =================== action_type = "setMinTarget" min_target_capacity = 1 // ================================== // === ADJUSTMENT =================== # action_type = "adjustment" # action_type = "percentageAdjustment" # adjustment = "MAX(5,10)" // ================================== // === UPDATE CAPACITY ============== # action_type = "updateCapacity" # minimum = 0 # maximum = 10 # target = 5 // ================================== } ` + "`" + `` + "`" + `` + "`" + ` ` + "`" + `` + "`" + `` + "`" + `hcl scaling_target_policy { policy_name = "policy-name" metric_name = "CPUUtilization" namespace = "AWS/EC2" source = "spectrum" statistic = "average" unit = "bytes" cooldown = 120 target = 2 predictive_mode = "FORCAST_AND_SCALE" max_capacity_per_scale = "10" dimensions { name = "name-1" value = "value-1" } } ` + "`" + `` + "`" + `` + "`" + ` ` + "`" + `scaling_multiple_metrics` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "expressions",
					Description: `(Optional) Array of objects (Expression config)`,
				},
				resource.Attribute{
					Name:        "expression",
					Description: `(Required) An expression consisting of the metric names listed in the 'metrics' array.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The expression name.`,
				},
				resource.Attribute{
					Name:        "metrics",
					Description: `(Optional) Array of objects (Metric config)`,
				},
				resource.Attribute{
					Name:        "metric_name",
					Description: `(Required) The name of the source metric.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The expression name.`,
				},
				resource.Attribute{
					Name:        "name_space",
					Description: `(Required, default: ` + "`" + `AWS/EC2` + "`" + `) The namespace for the alarm's associated metric.`,
				},
				resource.Attribute{
					Name:        "statistic",
					Description: `(Optional) The metric statistics to return. Valid values: ` + "`" + `"average"` + "`" + `, ` + "`" + `"sum"` + "`" + `, ` + "`" + `"sampleCount"` + "`" + `, ` + "`" + `"maximum"` + "`" + `, ` + "`" + `"minimum"` + "`" + `, ` + "`" + `"percentile"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "extended_statistic",
					Description: `(Optional) Percentile statistic. Valid values: ` + "`" + `"p0.1"` + "`" + ` - ` + "`" + `"p100"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "unit",
					Description: `(Optional) The unit for the alarm's associated metric. Valid values: ` + "`" + `"seconds"` + "`" + `, ` + "`" + `"microseconds"` + "`" + `, ` + "`" + `"milliseconds"` + "`" + `, ` + "`" + `"bytes"` + "`" + `, ` + "`" + `"kilobytes"` + "`" + `, ` + "`" + `"megabytes"` + "`" + `, ` + "`" + `"gigabytes"` + "`" + `, ` + "`" + `"terabytes"` + "`" + `, ` + "`" + `"bits"` + "`" + `, ` + "`" + `"kilobits"` + "`" + `, ` + "`" + `"megabits"` + "`" + `, ` + "`" + `"gigabits"` + "`" + `, ` + "`" + `"terabits"` + "`" + `, ` + "`" + `"percent"` + "`" + `, ` + "`" + `"count"` + "`" + `, ` + "`" + `"bytes/second"` + "`" + `, ` + "`" + `"kilobytes/second"` + "`" + `, ` + "`" + `"megabytes/second"` + "`" + `, ` + "`" + `"gigabytes/second"` + "`" + `, ` + "`" + `"terabytes/second"` + "`" + `, ` + "`" + `"bits/second"` + "`" + `, ` + "`" + `"kilobits/second"` + "`" + `, ` + "`" + `"megabits/second"` + "`" + `, ` + "`" + `"gigabits/second"` + "`" + `, ` + "`" + `"terabits/second"` + "`" + `, ` + "`" + `"count/second"` + "`" + `, ` + "`" + `"none"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "dimensions",
					Description: `(Optional) The dimensions for the alarm's associated metric. When name is "instanceId", no value is needed.`,
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
					Description: `(Optional) Indicates whether to assign IPV6 addresses to your instance. Requires a subnet with IPV6 CIDR block ranges. Usage: ` + "`" + `` + "`" + `` + "`" + `hcl network_interface { network_interface_id = "" device_index = 1 description = "nic description in here" private_ip_address = "1.1.1.1" delete_on_termination = false secondary_private_ip_address_count = 1 associate_public_ip_address = true } ` + "`" + `` + "`" + `` + "`" + ` <a id="block-devices"></a> ## Block Devices Each of the ` + "`" + ``,
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
					Description: `(Optional) ID for a user managed CMK under which the EBS Volume is encrypted`,
				},
				resource.Attribute{
					Name:        "device_name",
					Description: `(Required) The name of the block device to mount on the instance.`,
				},
				resource.Attribute{
					Name:        "virtual_name",
					Description: `(Required) The [Instance Store Device Name](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/InstanceStorage.html#InstanceStoreDeviceNames) (e.g. ` + "`" + `"ephemeral0"` + "`" + `). Usage: ` + "`" + `` + "`" + `` + "`" + `hcl ephemeral_block_device { device_name = "/dev/xvdc" virtual_name = "ephemeral0" } ` + "`" + `` + "`" + `` + "`" + ` <a id="stateful"></a> ## Stateful We support instance persistence via the following configurations. all values are boolean. For more information on instance persistence please see: [Stateful configuration](https://docs.spot.io/elastigroup/features/stateful-instance/stateful-instances)`,
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
					Description: `(Optional) For stateful groups: remove snapshots. Usage: ` + "`" + `` + "`" + `` + "`" + `hcl stateful_deallocation { should_delete_images = false should_delete_network_interfaces = false should_delete_volumes = false should_delete_snapshots = false } ` + "`" + `` + "`" + `` + "`" + ` <a id="stateful_instance_action"></a> ## Stateful Instance Action`,
				},
				resource.Attribute{
					Name:        "stateful_instance_action",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "stateful_instance_id",
					Description: `(Required) String, Stateful Instance ID on which the action should be performed.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) String, Action type. Supported action types: ` + "`" + `pause` + "`" + `, ` + "`" + `resume` + "`" + `, ` + "`" + `recycle` + "`" + `, ` + "`" + `deallocate` + "`" + `. Usage: ` + "`" + `` + "`" + `` + "`" + `hcl stateful_instance_action { type = "pause" stateful_instance_id = "ssi-foo" } stateful_instance_action { type = "recycle" stateful_instance_id = "ssi-bar" } ` + "`" + `` + "`" + `` + "`" + ` <a id="health-check"></a> ## Health Check`,
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
					Description: `(Optional) The Rancher version. Must be ` + "`" + `"1"` + "`" + ` or ` + "`" + `"2"` + "`" + `. If this field is omitted, it’s assumed that the Rancher cluster is version 1. Note that Kubernetes is required when using Rancher version 2^. Usage: ` + "`" + `` + "`" + `` + "`" + `hcl integration_rancher { master_host = "master_host" access_key = "access_key" secret_key = "secret_key" version = "2" } ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `(Optional) A key/value mapping of tags to assign to the resource.`,
				},
				resource.Attribute{
					Name:        "batch",
					Description: `(Optional) Batch configuration object:`,
				},
				resource.Attribute{
					Name:        "job_queue_names",
					Description: `(Required) Array of strings. Usage: ` + "`" + `` + "`" + `` + "`" + `hcl integration_ecs { cluster_name = "ecs-cluster" autoscale_is_enabled = false autoscale_cooldown = 300 autoscale_scale_down_non_service_tasks = false autoscale_headroom { cpu_per_unit = 1024 memory_per_unit = 512 num_of_units = 2 } autoscale_down { evaluation_periods = 300 max_scale_down_percentage = 70 } autoscale_attributes { key = "test.ecs.key" value = "test.ecs.value" } batch { job_queue_names = [ "first-job", "second-job" ] } } ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `(Optional) The deployment group name. Usage: ` + "`" + `` + "`" + `` + "`" + `hcl integration_codedeploy { cleanup_on_failure = false terminate_instance_on_failure = false deployment_groups { application_name = "my-app" deployment_group_name = "my-group" } } ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `(Optional) The Spotinst account ID that is linked to the AWS account that holds the Route 53 Hosted Zone ID. The default is the user Spotinst account provided as a URL parameter.`,
				},
				resource.Attribute{
					Name:        "record_set_type",
					Description: `(Optional, Default: ` + "`" + `a` + "`" + `) The type of the record set. Valid values: ` + "`" + `"a"` + "`" + `, ` + "`" + `"cname"` + "`" + `.`,
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
					Description: `(Optional, Default: ` + "`" + `false` + "`" + `) - Designates whether the IP address should be exposed to connections outside the VPC.`,
				},
				resource.Attribute{
					Name:        "use_public_dns",
					Description: `(Optional, Default: ` + "`" + `false` + "`" + `) - Designates whether the DNS address should be exposed to connections outside the VPC. Usage: ` + "`" + `` + "`" + `` + "`" + `hcl integration_route53 { # Option 1: Use A records. domains { hosted_zone_id = "zone-id" spotinst_acct_id = "act-123456" record_set_type = "a" record_sets { name = "foo.example.com" use_public_ip = true } } # Option 2: Use CNAME records. domains { hosted_zone_id = "zone-id" spotinst_acct_id = "act-123456" record_set_type = "cname" record_sets { name = "foo.example.com" use_public_dns = true } } } ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `(Optional, Default: ` + "`" + `5` + "`" + `) Number of periods over which data is compared. Minimum 3, Measured in consecutive minutes.`,
				},
				resource.Attribute{
					Name:        "max_scale_down_percentage",
					Description: `(Optional) Would represent the maximum % to scale-down. Number between 1-100. Usage: ` + "`" + `` + "`" + `` + "`" + `hcl integration_docker_swarm { master_host = "10.10.10.10" master_port = 2376 autoscale_is_enabled = true autoscale_cooldown = 180 autoscale_headroom { cpu_per_unit = 2048 memory_per_unit = 2048 num_of_units = 1 } autoscale_down { evaluation_periods = 3 max_scale_down_percentage = 30 } } ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `(Optional) A key/value mapping of tags to assign to the resource. Usage: ` + "`" + `` + "`" + `` + "`" + `hcl integration_kubernetes { integration_mode = "pod" cluster_identifier = "my-identifier.ek8s.com" // === SAAS =================== # integration_mode = "saas" # api_server = "https://api.my-identifier.ek8s.com/api/v1/namespaces/kube-system/services/..." # token = "top-secret" // ============================ autoscale_is_enabled = false autoscale_is_auto_config = false autoscale_cooldown = 300 autoscale_headroom { cpu_per_unit = 1024 memory_per_unit = 512 num_of_units = 1 } autoscale_down { evaluation_periods = 300 } autoscale_labels { key = "test.k8s.key" value = "test.k8s.value" } } ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `(Optional) A key/value mapping of tags to assign to the resource. Usage: ` + "`" + `` + "`" + `` + "`" + `hcl integration_nomad { master_host = "my-nomad-host" master_port = 9000 acl_token = "top-secret" autoscale_is_enabled = false autoscale_cooldown = 300 autoscale_headroom { cpu_per_unit = 1024 memory_per_unit = 512 num_of_units = 2 } autoscale_down { evaluation_periods = 300 } autoscale_constraints { key = "test.nomad.key" value = "test.nomad.value" } } ` + "`" + `` + "`" + `` + "`" + ``,
				},
				resource.Attribute{
					Name:        "integration_mesosphere",
					Description: `(Optional) Describes the [Mesosphere](https://mesosphere.com/) integration.`,
				},
				resource.Attribute{
					Name:        "api_server",
					Description: `(Optional) The public IP of the DC/OS Master. Usage: ` + "`" + `` + "`" + `` + "`" + `hcl integration_mesosphere { api_server = "" } ` + "`" + `` + "`" + `` + "`" + ``,
				},
				resource.Attribute{
					Name:        "integration_multai_runtime",
					Description: `(Optional) Describes the [Multai Runtime](https://spotinst.com/) integration.`,
				},
				resource.Attribute{
					Name:        "deployment_id",
					Description: `(Optional) The deployment id you want to get Usage: ` + "`" + `` + "`" + `` + "`" + `hcl integration_multai_runtime { deployment_id = "" } ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `(Optional, Default: ` + "`" + `false` + "`" + `) Specifies whether the integration is enabled. Usage: ` + "`" + `` + "`" + `` + "`" + `hcl integration_gitlab { runner { is_enabled = true } } ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `(Required) - Level to update Usage: ` + "`" + `` + "`" + `` + "`" + `hcl integration_beanstalk { environment_id = "e-3tkmbj7hzc" deployment_preferences { automatic_roll = true batch_size_percentage = 100 grace_period = 90 strategy { action = "REPLACE_SERVER" should_drain_instances = true } } managed_actions { platform_update { perform_at = "timeWindow" time_window = "Mon:23:50-Tue:00:20" update_level = "minorAndPatch" } } } ` + "`" + `` + "`" + `` + "`" + ` <a id="update-policy"></a> ## Update Policy`,
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
					Description: `(Optional, Default ` + "`" + `50` + "`" + `) Indicates the threshold of minimum healthy instances in single batch. If the amount of healthy instances in single batch is under the threshold, the deployment will fail. Range ` + "`" + `1` + "`" + ` - ` + "`" + `100` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "on_failure",
					Description: `(Optional) Set detach options to the deployment.`,
				},
				resource.Attribute{
					Name:        "action_type",
					Description: `(Required) Sets the action that will take place, Accepted values are: ` + "`" + `DETACH_OLD` + "`" + `, ` + "`" + `DETACH_NEW` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "should_handle_all_batches",
					Description: `(Optional, Default: ` + "`" + `false` + "`" + `) Indicator if the action should apply to all batches of the deployment or only the latest batch.`,
				},
				resource.Attribute{
					Name:        "draining_timeout",
					Description: `(Optional, Default: The Elastigroups draining time out) Indicates (in seconds) the timeout to wait until instance are detached.`,
				},
				resource.Attribute{
					Name:        "should_decrement_target_capacity",
					Description: `(Optional, Default: ` + "`" + `true` + "`" + `) Decrementing the group target capacity after detaching the instances. ` + "`" + `` + "`" + `` + "`" + `hcl update_policy { should_resume_stateful = false should_roll = false auto_apply_tags = false roll_config { batch_size_percentage = 33 health_check_type = "ELB" grace_period = 300 wait_for_roll_percentage = 10 wait_for_roll_timeout = 1500 strategy { action = "REPLACE_SERVER" should_drain_instances = false batch_min_healthy_percentage = 10 on_failure { action_type = "DETACH_NEW" should_handle_all_batches = true draining_timeout = 600 should_decrement_target_capacity = true } } } } ` + "`" + `` + "`" + `` + "`" + ` ## Attributes Reference The following attributes are exported:`,
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
			Type:             "spotinst_elastigroup_aws_beanstalk",
			Category:         "Elastigroup",
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
					Description: `(Required) - Level to update <a id="scheduled-task"></a> ## Scheduled Tasks Each ` + "`" + `scheduled_task` + "`" + ` supports the following:`,
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
					Description: `(Optional; Min 1) The percentage of instances to add or remove. Usage: ` + "`" + `` + "`" + `` + "`" + `hcl scheduled_task { task_type = "backup_ami" cron_expression = "" start_time = "1970-01-01T01:00:00Z" frequency = "hourly" scale_target_capacity = 5 scale_min_capacity = 0 scale_max_capacity = 10 is_enabled = false target_capacity = 5 min_capacity = 0 max_capacity = 10 batch_size_percentage = 33 grace_period = 300 } ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "spotinst_elastigroup_aws_suspension",
			Category:         "Elastigroup",
			ShortDescription: `Provides a process suspension to Spotinst AWS group resources.`,
			Description:      ``,
			Keywords: []string{
				"elastigroup",
				"aws",
				"suspension",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "group_id",
					Description: `(Required; string) Elastigroup ID to apply the suspensions on.`,
				},
				resource.Attribute{
					Name:        "suspension",
					Description: `(Required; at least one block is required) block of single process to suspend.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required; string) The name of process to suspend. Valid values: ` + "`" + `"AUTO_HEALING" , "OUT_OF_STRATEGY", "PREVENTIVE_REPLACEMENT", "REVERT_PREFERRED", or "SCHEDULING"` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "spotinst_elastigroup_azure",
			Category:         "Elastigroup",
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
					Description: `(Optional, Default: ` + "`" + `false` + "`" + `) ` + "`" + `` + "`" + `` + "`" + `hcl load_balancers { type = "MULTAI_TARGET_SET" balancer_id = "lb-1ee2e3q" target_set_id = "ts-3eq" auto_weight = true } ` + "`" + `` + "`" + `` + "`" + ` <a id="image"></a> ## Image`,
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
					Description: `(Optional) Name of the custom image. Required if resource_group_name is specified. ` + "`" + `` + "`" + `` + "`" + `hcl // market image image { marketplace { publisher = "Canonical" offer = "UbuntuServer" sku = "16.04-LTS" } } // custom image image { custom { image_name = "customImage" resource_group_name = "resourceGroup" } } ` + "`" + `` + "`" + `` + "`" + ` <a id="health-check"></a> ## Health Check`,
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
					Description: `(Optional) Enable auto-healing of unhealthy VMs. ` + "`" + `` + "`" + `` + "`" + `hcl health_check { health_check_type = "INSTANCE_STATE" grace_period = 120 auto_healing = true } ` + "`" + `` + "`" + `` + "`" + ` <a id="network"></a> ## Network`,
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
					Description: `(Optional) Available from Azure Api-Version 2017-03-30 onwards, it represents whether the specific ipconfiguration is IPv4 or IPv6. Valid values: ` + "`" + `IPv4` + "`" + `, ` + "`" + `IPv6` + "`" + `. ` + "`" + `` + "`" + `` + "`" + `hcl network { virtual_network_name = "vname" subnet_name = "my-subnet-name" resource_group_name = "subnetResourceGroup" assign_public_ip = true additional_ip_configs { name = "test" private_ip_version = "IPv4" } } ` + "`" + `` + "`" + `` + "`" + ` <a id="login"></a> ## Login ` + "`" + `` + "`" + `` + "`" + `hcl network { virtual_network_name = "vname" subnet_name = "my-subnet-name" resource_group_name = "subnetResourceGroup" assign_public_ip = true } ` + "`" + `` + "`" + `` + "`" + ` <a id="login"></a> ## Login`,
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
					Description: `(Optional) Password for admin access to Windows VMs. Required for Windows product types. ` + "`" + `` + "`" + `` + "`" + `hcl login { user_name = "admin" ssh_public_key = "33a2s1f3g5a1df5g1ad21651sag56dfg==" } ` + "`" + `` + "`" + `` + "`" + ` <a id="scaling-policy"></a> ## Scaling Policies Each ` + "`" + `scaling_`,
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
					Description: `(Optional) The dimension value. When ` + "`" + `namespace` + "`" + ` is defined and is not ` + "`" + `"Microsoft.Compute"` + "`" + ` the list of dimensions must contain the following: ` + "`" + `` + "`" + `` + "`" + `hcl dimensions { name = "resourceName" value = "example-resource-name" } dimensions { name = "resourceGroupName" value = "example-resource-group-name" } ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `(Optional; if using ` + "`" + `updateCapacity` + "`" + `) The target number of instances to have in the group. Usage: ` + "`" + `` + "`" + `` + "`" + `hcl // --- SCALE DOWN POLICY ------------------ scaling_down_policy { policy_name = "policy-name" metric_name = "CPUUtilization" namespace = "Microsoft.Compute" statistic = "average" threshold = 10 unit = "percent" cooldown = 60 dimensions { name = "name-1" value = "value-1" } operator = "gt" evaluation_periods = "10" period = "60" // === MIN TARGET =================== # action_type = "setMinTarget" # min_target_capacity = 1 // ================================== // === ADJUSTMENT =================== action_type = "adjustment" # action_type = "percentageAdjustment" adjustment = "MIN(5,10)" // ================================== // === UPDATE CAPACITY ============== # action_type = "updateCapacity" # minimum = 0 # maximum = 10 # target = 5 // ================================== } // ---------------------------------------- // --- SCALE DOWN POLICY ------------------ scaling_down_policy { policy_name = "policy-name-update" metric_name = "CPUUtilization" namespace = "Microsoft.Compute" statistic = "sum" threshold = 5 unit = "bytes" cooldown = 120 dimensions { name = "name-1-update" value = "value-1-update" } operator = "lt" evaluation_periods = 5 period = 120 //// === MIN TARGET =================== # action_type = "setMinTarget" # min_target_capacity = 1 //// ================================== // === ADJUSTMENT =================== # action_type = "percentageAdjustment" # action_type = "adjustment" # adjustment = "MAX(5,10)" // ================================== // === UPDATE CAPACITY ============== action_type = "updateCapacity" minimum = 0 maximum = 10 target = 5 // ================================== } // ---------------------------------------- ` + "`" + `` + "`" + `` + "`" + ` <a id="scheduling"></a> ## Scheduling`,
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
					Description: `(Optional) The time to allow instances to become healthy. ` + "`" + `` + "`" + `` + "`" + `hcl scheduled_task { is_enabled = true cron_expression = "`,
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
					Description: `(Optional) Sets the grace period for new instances to become healthy. ` + "`" + `` + "`" + `` + "`" + `hcl update_policy { should_roll = false roll_config { batch_size_percentage = 33 health_check_type = "INSTANCE_STATE" grace_period = 300 } } ` + "`" + `` + "`" + `` + "`" + ` <a id="third-party-integrations"></a> ## Third-Party Integrations`,
				},
				resource.Attribute{
					Name:        "integration_kubernetes",
					Description: `(Optional) Describes the [Kubernetes](https://kubernetes.io/) integration.`,
				},
				resource.Attribute{
					Name:        "cluster_identifier",
					Description: `(Required) The cluster ID. Usage: ` + "`" + `` + "`" + `` + "`" + `hcl integration_kubernetes { cluster_identifier = "k8s-cluster-id" } ` + "`" + `` + "`" + `` + "`" + ``,
				},
				resource.Attribute{
					Name:        "integration_multai_runtime",
					Description: `(Optional) Describes the [Multai Runtime](https://spotinst.com/) integration.`,
				},
				resource.Attribute{
					Name:        "deployment_id",
					Description: `(Optional) The deployment id you want to get Usage: ` + "`" + `` + "`" + `` + "`" + `hcl integration_multai_runtime { deployment_id = "" } ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "spotinst_elastigroup_azure_v3",
			Category:         "Elastigroup",
			ShortDescription: `Provides a Spotinst elastigroup resource for Microsoft Azure.`,
			Description:      ``,
			Keywords: []string{
				"elastigroup",
				"azure",
				"v3",
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
					Name:        "os",
					Description: `(Required) Type of the operating system. Valid values: ` + "`" + `"Linux"` + "`" + `, ` + "`" + `"Windows"` + "`" + `.`,
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
					Name:        "custom_data",
					Description: `(Optional) Custom init script file or text in Base64 encoded format.`,
				},
				resource.Attribute{
					Name:        "managed_service_identity",
					Description: `(Optional) List of Managed Service Identity objects.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `(Required) Name of the Azure Resource Group where the Managed Service Identity is located.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the Managed Service Identity.`,
				},
				resource.Attribute{
					Name:        "od_sizes",
					Description: `(Required) Available On-Demand sizes`,
				},
				resource.Attribute{
					Name:        "spot_sizes",
					Description: `(Required) Available Low-Priority sizes. <a id="strategy"></a> ## Strategy`,
				},
				resource.Attribute{
					Name:        "spot_percentage",
					Description: `(Optional) Percentage of Spot-VMs to maintain. Required if ` + "`" + `on_demand_count` + "`" + ` is not specified.`,
				},
				resource.Attribute{
					Name:        "on_demand_count",
					Description: `(Optional) Number of On-Demand VMs to maintain. Required if ` + "`" + `spot_percentage` + "`" + ` is not specified.`,
				},
				resource.Attribute{
					Name:        "fallback_to_on_demand",
					Description: ``,
				},
				resource.Attribute{
					Name:        "draining_timeout",
					Description: `(Optional, Default ` + "`" + `120` + "`" + `) Time (seconds) to allow the instance to be drained from incoming TCP connections and detached from MLB before terminating it during a scale-down operation. <a id="image"></a> ## Image`,
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
					Name:        "version",
					Description: ``,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `(Optional) Name of Resource Group for custom image. Required if publisher not specified.`,
				},
				resource.Attribute{
					Name:        "image_name",
					Description: `(Optional) Name of the custom image. Required if resource_group_name is specified. ` + "`" + `` + "`" + `` + "`" + `hcl // market image image { marketplace { publisher = "Canonical" offer = "UbuntuServer" sku = "18.04-LTS" version = "latest" } } // custom image image { custom { image_name = "customImage" resource_group_name = "resourceGroup" } } ` + "`" + `` + "`" + `` + "`" + ` <a id="network"></a> ## Network`,
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
					Name:        "resource_group_name",
					Description: `(Required) Vnet Resource Group Name.`,
				},
				resource.Attribute{
					Name:        "network_interfaces",
					Description: ``,
				},
				resource.Attribute{
					Name:        "subnet_name",
					Description: `(Required) ID of subnet.`,
				},
				resource.Attribute{
					Name:        "assign_public_up",
					Description: `(Optional, Default: ` + "`" + `false` + "`" + `) Assign a public IP to each VM in the Elastigroup.`,
				},
				resource.Attribute{
					Name:        "is_primary",
					Description: ``,
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
					Description: `(Optional) Available from Azure Api-Version 2017-03-30 onwards, it represents whether the specific ip configuration is IPv4 or IPv6. Valid values: ` + "`" + `IPv4` + "`" + `, ` + "`" + `IPv6` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "application_security_group",
					Description: `(Optional) - List of Application Security Groups that will be associated to the primary ip configuration of the network interface.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) - The name of the Application Security group.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `(Required) - The resource group of the Application Security Group. } ` + "`" + `` + "`" + `` + "`" + `hcl network { virtual_network_name = "VirtualNetworkName" resource_group_name = "ResourceGroup" network_interfaces { subnet_name = "default" assign_public_ip = false is_primary = true additional_ip_configs { name = "SecondaryIPConfig" PrivateIPVersion = "IPv4" } application_security_group { name = "ApplicationSecurityGroupName" resource_group_name = "ResourceGroup" } } } ` + "`" + `` + "`" + `` + "`" + ` <a id="login"></a> ## Login`,
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
					Description: `(Optional) SSH for admin access to Linux VMs. Required for Linux OS types.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional) Password for admin access to Windows VMs. Required for Windows OS types. ` + "`" + `` + "`" + `` + "`" + `hcl login { user_name = "admin" ssh_public_key = "33a2s1f3g5a1df5g1ad21651sag56dfg==" } ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "spotinst_elastigroup_gcp",
			Category:         "Elastigroup",
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
					Name:        "provisioning_model",
					Description: `(Optional) Valid values: "SPOT", "PREEMPTIBLE". Define the provisioning model of the launched instances. Default value is "PREEMPTIBLE".`,
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
					Description: `(Optional) Tags to mark created instances.`,
				},
				resource.Attribute{
					Name:        "instance_name_prefix",
					Description: `(Optional) Set an instance name prefix to be used for all launched instances and their boot disk. The prefix value should comply with the following limitations:`,
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
					Description: `(Required) The number of GPUs. Must be 0, 2, 4, 6, 8. Usage: ` + "`" + `` + "`" + `` + "`" + `hcl gpu { count = 2 type = "nvidia-tesla-p100" } ` + "`" + `` + "`" + `` + "`" + ` <a id="health-check"></a> ## Health Check`,
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
					Name:        "named_ports",
					Description: `(Optional) Describes a named port and a list of ports.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the port.`,
				},
				resource.Attribute{
					Name:        "ports",
					Description: `(Required) A list of ports. Usage: ` + "`" + `` + "`" + `` + "`" + `hcl backend_services { service_name = "spotinst-elb-backend-service" location_type = "regional" scheme = "INTERNAL" named_ports { name = "port-name" ports = [8000, 6000] } } ` + "`" + `` + "`" + `` + "`" + ` <a id="disks"></a> ## Disks`,
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
					Description: `(Optional) A source image used to create the disk. You can provide a private (custom) image, and Compute Engine will use the corresponding image from your project. Usage: ` + "`" + `` + "`" + `` + "`" + `hcl disks { device_name = "device" mode = "READ_WRITE" type = "PERSISTENT" auto_delete = true boot = true interface = "SCSI" initialize_params = { disk_size_gb = 10 disk_type = "pd-standard" source_image = "" } } ` + "`" + `` + "`" + `` + "`" + ` <a id="network-interface"></a> ## Network Interfaces Each of the ` + "`" + `network_interface` + "`" + ` attributes controls a portion of the GCP Instance's "Network Interfaces". It's a good idea to familiarize yourself with [GCP's Network Interfaces docs](https://cloud.google.com/vpc/docs/multiple-interfaces-concepts) to understand the implications of using these attributes.`,
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
					Description: `(Optional) Array of configurations for this interface. Currently, only ONE_TO_ONE_NAT is supported. ` + "`" + `` + "`" + `` + "`" + `hcl network_interface { network = "default" access_configs = { name = "config1" type = "ONE_TO_ONE_NAT" } alias_ip_ranges = { subnetwork_range_name = "range-name-1" ip_cidr_range = "10.128.0.0/20" } } ` + "`" + `` + "`" + `` + "`" + ` <a id="scaling-policy"></a> ## Scaling Policies`,
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
					Name:        "action_type",
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
					Description: `(Required) The dimension value. Usage: ` + "`" + `` + "`" + `` + "`" + `hcl scaling_up_policy { policy_name = "scale_up_1" source = "stackdriver" metric_name = "instance/disk/read_ops_count" namespace = "compute" statistic = "average" unit = "percent" threshold = 10000 period = 300 cooldown = 300 operator = "gte" evaluation_periods = 1 action_type = "adjustment" adjustment = 1 dimensions { name = "storage_type" value = "pd-ssd" } } ` + "`" + `` + "`" + `` + "`" + ` <a id="third-party-integrations"></a> ## Third-Party Integrations`,
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
					Description: `(Optional) The maximum number of instances the group should have. Usage: ` + "`" + `` + "`" + `` + "`" + `hcl scheduled_task { task_type = "setCapacity" cron_expression = "" is_enabled = false target_capacity = 5 min_capacity = 0 max_capacity = 10 } ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "spotinst_elastigroup_gke",
			Category:         "Elastigroup",
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
					Description: `(Optional) The label value. Usage: ` + "`" + `` + "`" + `` + "`" + `hcl integration_gke { location = "us-central1-a" cluster_id = "terraform-acc-test-cluster" autoscale_is_enabled = true autoscale_is_auto_config = false autoscale_cooldown = 300 autoscale_headroom { cpu_per_unit = 1024 memory_per_unit = 512 num_of_units = 2 } autoscale_down { evaluation_periods = 300 } autoscale_labels { key = "label_key" value = "label_value" } } ` + "`" + `` + "`" + `` + "`" + ` <a id="diff-suppressed-parameters"></a> ## Diff-suppressed Parameters The following parameters are created remotely and imported. The diffs have been suppressed in order to maintain plan legibility. You may update the values of these imported parameters by defining them in your template with your desired new value (including null values).`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "spotinst_health_check",
			Category:         "Elastigroup",
			ShortDescription: `Provides a Spotinst Health Check resource.`,
			Description:      ``,
			Keywords: []string{
				"elastigroup",
				"health",
				"check",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the health check.`,
				},
				resource.Attribute{
					Name:        "resource_id",
					Description: `(Required) The ID of the resource to check.`,
				},
				resource.Attribute{
					Name:        "check",
					Description: `(Required) Describes the check to execute.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Required) The protocol to use to connect with the instance. Valid values: http, https.`,
				},
				resource.Attribute{
					Name:        "endpoint",
					Description: `(Required) The destination for the request.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Required) The port to use to connect with the instance.`,
				},
				resource.Attribute{
					Name:        "interval",
					Description: `(Required) The amount of time (in seconds) between each health check (minimum: 10).`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `(Required) the amount of time (in seconds) to wait when receiving a response from the health check.`,
				},
				resource.Attribute{
					Name:        "threshold",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "healthy",
					Description: `(Required) The number of consecutive successful health checks that must occur before declaring an instance healthy.`,
				},
				resource.Attribute{
					Name:        "unhealthy",
					Description: `(Required) The number of consecutive failed health checks that must occur before declaring an instance unhealthy.`,
				},
				resource.Attribute{
					Name:        "proxy",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "addr",
					Description: `(Required) The public hostname / IP where you installed the Spotinst HCS.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Required) The port of the Spotinst HCS (default: 80). ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The Health Check ID.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The Health Check ID.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "spotinst_managed_instance_aws",
			Category:         "Managed Instance",
			ShortDescription: `Provides a Spotinst AWS managed instance resource.`,
			Description:      ``,
			Keywords: []string{
				"managed",
				"instance",
				"aws",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The ManagedInstance name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The ManagedInstance description.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Required) The AWS region your group will be created in.`,
				},
				resource.Attribute{
					Name:        "life_cycle",
					Description: `(Optional) Set lifecycle, valid values: ` + "`" + `"spot"` + "`" + `, ` + "`" + `"on_demand"` + "`" + `. Default ` + "`" + `"spot"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "orientation",
					Description: `(Optional) Select a prediction strategy. Valid values: ` + "`" + `"balanced"` + "`" + `, ` + "`" + `"costOriented"` + "`" + `, ` + "`" + `"availabilityOriented"` + "`" + `, ` + "`" + `"cheapest"` + "`" + `. Default: ` + "`" + `"availabilityOriented"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "draining_timeout",
					Description: `(Optional) The time in seconds to allow the instance be drained from incoming TCP connections and detached from ELB before terminating it, during a scale down operation.`,
				},
				resource.Attribute{
					Name:        "fallback_to_ondemand",
					Description: `(Optional) In case of no spots available, Managed Instance will launch an On-demand instance instead. Default: ` + "`" + `"true"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "utilize_reserved_instances",
					Description: `(Optional) In case of any available Reserved Instances, Managed Instance will utilize them before purchasing Spot instances. Default: ` + "`" + `"false"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "optimization_windows",
					Description: `(Optional) When ` + "`" + `performAt` + "`" + ` is ` + "`" + `"timeWindow"` + "`" + `: must specify a list of ` + "`" + `"timeWindows"` + "`" + ` with at least one time window. Each string should be formatted as ` + "`" + `ddd:hh:mm-ddd:hh:mm` + "`" + ` (ddd = day of week = Sun | Mon | Tue | Wed | Thu | Fri | Sat hh = hour 24 = 0 -23 mm = minute = 0 - 59).`,
				},
				resource.Attribute{
					Name:        "perform_at",
					Description: `(Optional) Valid values: ` + "`" + `"always"` + "`" + `, ` + "`" + `"never"` + "`" + `, ` + "`" + `"timeWindow"` + "`" + `. Default ` + "`" + `"never"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "minimum_instnace_lifetime",
					Description: `(Optional) Defines the preferred minimum instance lifetime. Markets which comply with this preference will be prioritized. Optional values: ` + "`" + `1` + "`" + `, ` + "`" + `3` + "`" + `, ` + "`" + `6` + "`" + `, ` + "`" + `12` + "`" + `, ` + "`" + `24` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "persist_private_ip",
					Description: `(Optional) Should the instance maintain its private IP.`,
				},
				resource.Attribute{
					Name:        "persist_block_devices",
					Description: `(Optional) Should the instance maintain its Data volumes.`,
				},
				resource.Attribute{
					Name:        "persist_root_device",
					Description: `(Optional) Should the instance maintain its root device volumes.`,
				},
				resource.Attribute{
					Name:        "block_devices_mode",
					Description: `(Optional) Determine the way we attach the data volumes to the data devices. Valid values: ` + "`" + `"reattach"` + "`" + `, ` + "`" + `"onLaunch"` + "`" + `. Default: ` + "`" + `"onLaunch"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "health_check_type",
					Description: `(Optional) The service to use for the health check. Valid values: ` + "`" + `"EC2"` + "`" + `, ` + "`" + `"ELB"` + "`" + `, ` + "`" + `"TARGET_GROUP"` + "`" + `, ` + "`" + `"MULTAI_TARGET_SET"` + "`" + `. Default: ` + "`" + `"EC2"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "auto_healing",
					Description: `(Optional) Enable the auto healing which auto replaces the instance in case the health check fails, default: ` + "`" + `"true"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "grace_period",
					Description: `(Optional) The amount of time, in seconds, after the instance has launched to starts and check its health, default ` + "`" + `"120"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "unhealthy_duration",
					Description: `(Optional) The amount of time, in seconds, an existing instance should remain active after becoming unhealthy. After the set time out the instance will be replaced, default ` + "`" + `"120"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "subnet_ids",
					Description: `(Required) A comma-separated list of subnet identifiers for your instance.`,
				},
				resource.Attribute{
					Name:        "vpcId",
					Description: `(Required) VPC id for your instance.`,
				},
				resource.Attribute{
					Name:        "elastic_ip",
					Description: `(Optional) Elastic IP Allocation Id to associate to the instance.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `(Optional) Private IP Allocation Id to associate to the instance.`,
				},
				resource.Attribute{
					Name:        "product",
					Description: `(Required) Operation system type. Valid values: ` + "`" + `"Linux/UNIX"` + "`" + `, ` + "`" + `"SUSE Linux"` + "`" + `, ` + "`" + `"Windows"` + "`" + `, ` + "`" + `"Red Hat Enterprise Linux"` + "`" + `, ` + "`" + `"Linux/UNIX (Amazon VPC)"` + "`" + `, ` + "`" + `"SUSE Linux (Amazon VPC)"` + "`" + `, ` + "`" + `"Windows (Amazon VPC)"` + "`" + `, ` + "`" + `"Red Hat Enterprise Linux (Amazon VPC)"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "instance_types",
					Description: `(Required) Comma separated list of available instance types for instance.`,
				},
				resource.Attribute{
					Name:        "preferred_type",
					Description: `(Required) Preferred instance types for the instance. We will automatically select optional similar instance types to ensure optimized cost efficiency`,
				},
				resource.Attribute{
					Name:        "ebs_optimized",
					Description: `(Optional, Default: ` + "`" + `false` + "`" + `) Enable EBS optimization for supported instances. Note: Additional charges will be applied by the Cloud Provider. Default: false`,
				},
				resource.Attribute{
					Name:        "enable_monitoring",
					Description: `(Optional) Describes whether instance Enhanced Monitoring is enabled. Default: false`,
				},
				resource.Attribute{
					Name:        "placement_tenancy",
					Description: `(Optional) Valid values: ` + "`" + `"default"` + "`" + `, ` + "`" + `"dedicated"` + "`" + `. Default: default`,
				},
				resource.Attribute{
					Name:        "iam_instance_profile",
					Description: `(Optional) Set IAM profile to instance. Set only one of ARN or Name.`,
				},
				resource.Attribute{
					Name:        "security_group_ids",
					Description: `(Optional) One or more security group IDs.`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `(Required) The ID of the image used to launch the instance.`,
				},
				resource.Attribute{
					Name:        "key_pair",
					Description: `(Optional) Specify a Key Pair to attach to the instances.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Set tags for the instance. Items should be unique.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `Tag's key.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `Tag's name.`,
				},
				resource.Attribute{
					Name:        "user_data",
					Description: `(Optional) The Base64-encoded MIME user data to make available to the instances.`,
				},
				resource.Attribute{
					Name:        "shutdown_script",
					Description: `(Optional) The Base64-encoded shutdown script to execute prior to instance termination.`,
				},
				resource.Attribute{
					Name:        "cpu_credits",
					Description: `(Optional) cpuCredits can have one of two values: ` + "`" + `"unlimited"` + "`" + `, ` + "`" + `"standard"` + "`" + `. <a id="block-device-mapping"></a> ## Block Device Mapping`,
				},
				resource.Attribute{
					Name:        "block_device_mappings",
					Description: `(Optional) Attributes controls a portion of the AWS:`,
				},
				resource.Attribute{
					Name:        "device_name",
					Description: `(Required) The name of the device to mount.`,
				},
				resource.Attribute{
					Name:        "ebs",
					Description: `(Required) Object`,
				},
				resource.Attribute{
					Name:        "volume_type",
					Description: `(Optional, Default: ` + "`" + `"standard"` + "`" + `) The type of volume. Can be ` + "`" + `"standard"` + "`" + `, ` + "`" + `"gp2"` + "`" + `, ` + "`" + `"gp3"` + "`" + `, ` + "`" + `"io1"` + "`" + `, ` + "`" + `"st1"` + "`" + ` or ` + "`" + `"sc1"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "volume_size",
					Description: `(Optional) The size of the volume, in GiBs.`,
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
					Name:        "resource_tag_specification",
					Description: `(Optional) User will specify which resources should be tagged with group tags.`,
				},
				resource.Attribute{
					Name:        "should_tag_snapshots",
					Description: `(Optional) Tag specification for Snapshot resources.`,
				},
				resource.Attribute{
					Name:        "device_index",
					Description: `(Optional) The position of the network interface in the attachment order. A primary network interface has a device index of 0. If you specify a network interface when launching an instance, you must specify the device index.`,
				},
				resource.Attribute{
					Name:        "associate_public_ip_address",
					Description: `(Optional) Indicates whether to assign a public IPv4 address to an instance you launch in a VPC. The public IP address can only be assigned to a network interface for eth0, and can only be assigned to a new network interface, not an existing one. You cannot specify more than one network interface in the request. If launching into a default subnet, the default value is true.`,
				},
				resource.Attribute{
					Name:        "associate_ipv6_address",
					Description: `(Optional) Indicates whether to assign an IPv6 address. Amazon EC2 chooses the IPv6 addresses from the range of the subnet. Default: ` + "`" + `false` + "`" + ` Usage: ` + "`" + `` + "`" + `` + "`" + `hcl network_interface { device_index = 0 associate_public_ip_address = false associate_ipv6_address = true } ` + "`" + `` + "`" + `` + "`" + ` <a id="scheduled-task"></a> ## Scheduled Tasks Each ` + "`" + `scheduled_task` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "is_enabled",
					Description: `(Optional) Describes whether the task is enabled. When true the task should run when false it should not run.`,
				},
				resource.Attribute{
					Name:        "frequency",
					Description: `(Optional) Set frequency for the task. Valid values: "hourly", "daily", "weekly", "continuous".`,
				},
				resource.Attribute{
					Name:        "start_time",
					Description: `(Optional) DATETIME in ISO-8601 format. Sets a start time for scheduled actions. If "frequency" or "cronExpression" are not used - the task will run only once at the start time and will then be deleted from the instance configuration. Example: ` + "`" + `"2019-05-23T10:55:09Z"` + "`" + ``,
				},
				resource.Attribute{
					Name:        "cron_expression",
					Description: `(Optional) A valid cron expression. The cron is running in UTC time zone and is in Unix cron format Cron Expression Validator Script. Only one of ‘frequency’ or ‘cronExpression’ should be used at a time. Example: ` + "`" + `"0 1`,
				},
				resource.Attribute{
					Name:        "load_balancers",
					Description: `(Optional) List of load balancers configs.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The AWS resource name. Required for Classic Load Balancer. Optional for Application Load Balancer.`,
				},
				resource.Attribute{
					Name:        "arn",
					Description: `The AWS resource ARN (Required only for ALB target groups).`,
				},
				resource.Attribute{
					Name:        "balancer_id",
					Description: `The Multai load balancer ID. Example: lb-123456`,
				},
				resource.Attribute{
					Name:        "target_set_id",
					Description: `The Multai load target set ID. Example: ts-123456`,
				},
				resource.Attribute{
					Name:        "auto_weight",
					Description: `"Auto Weight" will automatically provide a higher weight for instances that are larger as appropriate. For example, if you have configured your Elastigroup with m4.large and m4.xlarge instances the m4.large will have half the weight of an m4.xlarge. This ensures that larger instances receive a higher number of MLB requests.`,
				},
				resource.Attribute{
					Name:        "az_awareness",
					Description: `"AZ Awareness" will ensure that instances within the same AZ are using the corresponding MLB runtime instance in the same AZ. This feature reduces multi-zone data transfer fees.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The resource type. Valid Values: ` + "`" + `"CLASSIC"` + "`" + `, ` + "`" + `"TARGET_GROUP"` + "`" + `, ` + "`" + `"MULTAI_TARGET_SET"` + "`" + `. Usage: ` + "`" + `` + "`" + `` + "`" + `hcl load_balancers { arn = "arn" type = "CLASSIC" balancer_id = "lb-123" target_set_id = "ts-123" auto_weight = "true" az_awareness = "true" } ` + "`" + `` + "`" + `` + "`" + ` <a id="route53"></a> ## Route53`,
				},
				resource.Attribute{
					Name:        "integration_route53",
					Description: `(Optional) Describes the [Route53](https://aws.amazon.com/documentation/route53/?id=docs_gateway) integration.`,
				},
				resource.Attribute{
					Name:        "domains",
					Description: `(Required) Route 53 Domain configurations.`,
				},
				resource.Attribute{
					Name:        "hosted_zone_id",
					Description: `(Required) The Route 53 Hosted Zone Id for the registered Domain.`,
				},
				resource.Attribute{
					Name:        "spotinst_acct_id",
					Description: `(Optional) The Spotinst account ID that is linked to the AWS account that holds the Route 53 hosted Zone Id. The default is the user Spotinst account provided as a URL parameter.`,
				},
				resource.Attribute{
					Name:        "record_set_type",
					Description: `(Optional, Default: ` + "`" + `a` + "`" + `) The type of the record set. Valid values: ` + "`" + `"a"` + "`" + `, ` + "`" + `"cname"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "record_sets",
					Description: `(Required) List of record sets`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The record set name.`,
				},
				resource.Attribute{
					Name:        "use_public_ip",
					Description: `(Optional, Default: ` + "`" + `false` + "`" + `) - Designates whether the IP address should be exposed to connections outside the VPC.`,
				},
				resource.Attribute{
					Name:        "use_public_dns",
					Description: `(Optional, Default: ` + "`" + `false` + "`" + `) - Designates whether the DNS address should be exposed to connections outside the VPC. Usage: ` + "`" + `` + "`" + `` + "`" + `hcl integration_route53 { # Option 1: Use A records. domains { hosted_zone_id = "zone-id" spotinst_acct_id = "act-123456" record_set_type = "a" record_sets { name = "foo.example.com" use_public_ip = true } } # Option 2: Use CNAME records. domains { hosted_zone_id = "zone-id" spotinst_acct_id = "act-123456" record_set_type = "cname" record_sets { name = "foo.example.com" use_public_dns = true } } } ` + "`" + `` + "`" + `` + "`" + ` <a id="managed_instance_action"></a> ## Managed Instance Action`,
				},
				resource.Attribute{
					Name:        "managed_instance_action",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) String, Action type. Supported action types: ` + "`" + `pause` + "`" + `, ` + "`" + `resume` + "`" + `, ` + "`" + `recycle` + "`" + `. Usage: ` + "`" + `` + "`" + `` + "`" + `hcl managed_instance_action { type = "pause" } ` + "`" + `` + "`" + `` + "`" + ` ## Attributes Reference The following attributes are exported:`,
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
			Type:             "spotinst_mrscaler_aws",
			Category:         "Mr Scaler",
			ShortDescription: `Provides a Spotinst MrScaler resource.`,
			Description:      ``,
			Keywords: []string{
				"mr",
				"scaler",
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
					Description: `(Optional) Specifies whether the cluster should remain available after completing all steps.`,
				},
				resource.Attribute{
					Name:        "retries",
					Description: `(Optional; Requires: ` + "`" + `timeout_action` + "`" + ` is set to ` + "`" + `terminateAndRetry` + "`" + `) Specifies the maximum number of times a capacity provisioning should be retried if the provisioning timeout is exceeded. Valid values: ` + "`" + `1-5` + "`" + `. <a id="task-group"></a> ## Task Group (Wrap, Clone, and New strategies)`,
				},
				resource.Attribute{
					Name:        "task_instance_types",
					Description: `(Required) The MrScaler instance types for the task nodes.`,
				},
				resource.Attribute{
					Name:        "task_desired_capacity",
					Description: `(Required) amount of instances in task group.`,
				},
				resource.Attribute{
					Name:        "task_max_size",
					Description: `(Optional) maximal amount of instances in task group.`,
				},
				resource.Attribute{
					Name:        "task_min_size",
					Description: `(Optional) The minimal amount of instances in task group.`,
				},
				resource.Attribute{
					Name:        "task_unit",
					Description: `(Optional, Default: ` + "`" + `instance` + "`" + `) Unit of task group for target, min and max. The unit could be ` + "`" + `instance` + "`" + ` or ` + "`" + `weight` + "`" + `. instance - amount of instances. weight - amount of vCPU.`,
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
					Name:        "core_desired_capacity",
					Description: `(Required) amount of instances in core group.`,
				},
				resource.Attribute{
					Name:        "core_max_size",
					Description: `(Optional) maximal amount of instances in core group.`,
				},
				resource.Attribute{
					Name:        "core_min_size",
					Description: `(Optional) The minimal amount of instances in core group.`,
				},
				resource.Attribute{
					Name:        "core_unit",
					Description: `(Optional, Default: ` + "`" + `instance` + "`" + `) Unit of task group for target, min and max. The unit could be ` + "`" + `instance` + "`" + ` or ` + "`" + `weight` + "`" + `. instance - amount of instances. weight - amount of vCPU.`,
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
					Name:        "master_target",
					Description: `(Optional; Default 1) Number of instances in the master group.`,
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
					Description: `(Optional) New max capacity for the elastigroup. <a id="termination-policies"></a> ## termination policies ## Example Usage ` + "`" + `` + "`" + `` + "`" + `hcl termination_policies { statements { namespace = "AWS/ElasticMapReduce" metric_name = "AppsRunning" statistic = "average" unit = "count" threshold = 1 period = 360 evaluation_periods = 4 operator = "lte" } statements { namespace = "AWS/ElasticMapReduce" metric_name = "AppsPending" statistic = "average" unit = "count" threshold = 0 period = 300 evaluation_periods = 3 operator = "lte" } } termination_policies { statements { namespace = "AWS/ElasticMapReduce" metric_name = "AppsRunning" statistic = "average" unit = "count" threshold = 0 period = 300 evaluation_periods = 3 operator = "lte" } } ` + "`" + `` + "`" + `` + "`" + ` ## Argument Reference`,
				},
				resource.Attribute{
					Name:        "termination_policies",
					Description: `(Optional) Allows defining termination policies for EMR clusters based on CloudWatch Metrics.`,
				},
				resource.Attribute{
					Name:        "statements",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "namespace",
					Description: `(Required) Must contain the value: ` + "`" + `AWS/ElasticMapReduce` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "metric_name",
					Description: `(Required) The name of the metric in CloudWatch which the statement will be based on.`,
				},
				resource.Attribute{
					Name:        "statistic",
					Description: `(Optional, Default: ` + "`" + `sum` + "`" + `) The aggregation method of the given metric. Valid Values: ` + "`" + `average` + "`" + ` | ` + "`" + `sum` + "`" + ` | ` + "`" + `sampleCount` + "`" + ` | ` + "`" + `maximum` + "`" + ` | ` + "`" + `minimum` + "`" + ``,
				},
				resource.Attribute{
					Name:        "unit",
					Description: `(Optional, Default: ` + "`" + `count` + "`" + `) The unit for a given metric. Valid Values: ` + "`" + `seconds` + "`" + ` | ` + "`" + `microseconds` + "`" + ` | ` + "`" + `milliseconds` + "`" + ` | ` + "`" + `bytes` + "`" + ` | ` + "`" + `kilobytes` + "`" + ` | ` + "`" + `megabytes` + "`" + ` | ` + "`" + `gigabytes` + "`" + ` | ` + "`" + `terabytes` + "`" + ` | ` + "`" + `bits` + "`" + ` | ` + "`" + `kilobits` + "`" + ` | ` + "`" + `megabits` + "`" + ` | ` + "`" + `gigabits` + "`" + ` | ` + "`" + `terabits` + "`" + ` | ` + "`" + `percent` + "`" + ` | ` + "`" + `count` + "`" + ` | ` + "`" + `bytes/second` + "`" + ` | ` + "`" + `kilobytes/second` + "`" + ` | ` + "`" + `megabytes/second` + "`" + ` | ` + "`" + `gigabytes/second` + "`" + ` | ` + "`" + `terabytes/second` + "`" + ` | ` + "`" + `bits/second` + "`" + ` | ` + "`" + `kilobits/second` + "`" + ` | ` + "`" + `megabits/second` + "`" + ` | ` + "`" + `gigabits/second` + "`" + ` | ` + "`" + `terabits/second` + "`" + ` | ` + "`" + `count/second` + "`" + ` | ` + "`" + `none` + "`" + ``,
				},
				resource.Attribute{
					Name:        "threshold",
					Description: `(Required) The value that the specified statistic is compared to.`,
				},
				resource.Attribute{
					Name:        "period",
					Description: `(Optional, Default: ` + "`" + `300` + "`" + `) The time window in seconds over which the statistic is applied.`,
				},
				resource.Attribute{
					Name:        "evaluation_periods",
					Description: `(Optional, Default: ` + "`" + `1` + "`" + `) The number of periods over which data is compared to the specified threshold.`,
				},
				resource.Attribute{
					Name:        "operator",
					Description: `(Optional, Default: ` + "`" + `gte` + "`" + `) The operator to use in order to determine if the policy is applicable. Valid values: ` + "`" + `gt` + "`" + ` | ` + "`" + `gte` + "`" + ` | ` + "`" + `lt` + "`" + ` | ` + "`" + `lte` + "`" + ` ## Attributes Reference The following attributes are exported:`,
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
			Type:             "spotinst_ocean_aks",
			Category:         "Ocean",
			ShortDescription: `Provides a Spotinst Ocean resource using AKS.`,
			Description:      ``,
			Keywords: []string{
				"ocean",
				"aks",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The Ocean cluster name.`,
				},
				resource.Attribute{
					Name:        "controller_cluster_id",
					Description: `(Required) A unique identifier used for connecting the Ocean SaaS platform and the Kubernetes cluster. Typically, the cluster name is used as its identifier.`,
				},
				resource.Attribute{
					Name:        "zones",
					Description: `(Optional) An Array holding default Availability Zones, this configures the availability zones the Ocean may launch instances in.`,
				},
				resource.Attribute{
					Name:        "aks_name",
					Description: `(Required) The AKS cluster name.`,
				},
				resource.Attribute{
					Name:        "acd_identifier",
					Description: `(Required) The AKS identifier. A valid identifier should be formatted as ` + "`" + `acd-nnnnnnnn` + "`" + ` and previously used identifiers cannot be reused.`,
				},
				resource.Attribute{
					Name:        "aks_resource_group_name",
					Description: `(Required) Name of the Azure Resource Group where the AKS cluster is located.`,
				},
				resource.Attribute{
					Name:        "ssh_public_key",
					Description: `(Required) SSH public key for admin access to Linux VMs.`,
				},
				resource.Attribute{
					Name:        "user_name",
					Description: `(Optional) Username for admin access to VMs.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `(Optional) Name of the Azure Resource Group into which VMs will be launched. Cannot be updated.`,
				},
				resource.Attribute{
					Name:        "custom_data",
					Description: `(Optional) Must contain a valid Base64 encoded string.`,
				},
				resource.Attribute{
					Name:        "max_pods",
					Description: `(Optional) The maximum number of pods per node in an AKS cluster.`,
				},
				resource.Attribute{
					Name:        "managed_service_identity",
					Description: `(Optional) List of Managed Service Identity objects.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `(Required) Name of the Azure Resource Group where the Managed Service Identity is located.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the Managed Service Identity.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `(Optional) Unique key-value pairs that will be used to tag VMs that are launched in the cluster.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Optional) Tag key.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Optional) Tag value.`,
				},
				resource.Attribute{
					Name:        "network",
					Description: `(Optional) Define the Virtual Network and Subnet.`,
				},
				resource.Attribute{
					Name:        "virtual_network_name",
					Description: `(Optional) Virtual network.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `(Optional) Vnet resource group name.`,
				},
				resource.Attribute{
					Name:        "network_interface",
					Description: `(Optional) A list of virtual network interfaces. The publicIpSku must be identical between all the network interfaces. One network interface must be set as the primary.`,
				},
				resource.Attribute{
					Name:        "subnet_name",
					Description: `(Optional) Subnet name.`,
				},
				resource.Attribute{
					Name:        "assign_public_ip",
					Description: `(Optional) Assign public IP.`,
				},
				resource.Attribute{
					Name:        "is_primary",
					Description: `(Optional) Defines whether the network interface is primary or not.`,
				},
				resource.Attribute{
					Name:        "additional_ip_config",
					Description: `(Optional) Additional configuration of network interface. The name fields between all the ` + "`" + `additional_ip_config` + "`" + ` must be unique.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Configuration name.`,
				},
				resource.Attribute{
					Name:        "private_ip_version",
					Description: `(Optional, Default: ` + "`" + `IPv4` + "`" + `) Supported values: ` + "`" + `IPv4` + "`" + `, ` + "`" + `IPv6` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "extension",
					Description: `(Optional) List of Azure extension objects.`,
				},
				resource.Attribute{
					Name:        "api_version",
					Description: `(Optional) API version of the extension.`,
				},
				resource.Attribute{
					Name:        "minor_version_auto_upgrade",
					Description: `(Optional) Toggles whether auto upgrades are allowed.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Extension name.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) Extension type.`,
				},
				resource.Attribute{
					Name:        "os_disk",
					Description: `(Optional) OS disk specifications.`,
				},
				resource.Attribute{
					Name:        "size_gb",
					Description: `(Optional) The size of the OS disk in GB.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) The type of the OS disk. Supported values: ` + "`" + `Standard_LRS` + "`" + `, ` + "`" + `Premium_LRS` + "`" + `, ` + "`" + `StandardSSD_LRS` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "load_balancer",
					Description: `(Optional) Configure Load Balancer.`,
				},
				resource.Attribute{
					Name:        "backend_pool_names",
					Description: `(Optional) Names of the Backend Pools to register the Cluster VMs to. Each Backend Pool is a separate load balancer.`,
				},
				resource.Attribute{
					Name:        "load_balancer_sku",
					Description: `(Optional) Supported values: ` + "`" + `Standard` + "`" + `, ` + "`" + `Basic` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the Load Balancer.`,
				},
				resource.Attribute{
					Name:        "resource_group_name",
					Description: `(Optional) The Resource Group name of the Load Balancer.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) The type of load balancer. Supported value: ` + "`" + `loadBalancer` + "`" + ``,
				},
				resource.Attribute{
					Name:        "image",
					Description: `(Optional) Image of VM. An image is a template for creating new VMs. Choose from Azure image catalogue (marketplace).`,
				},
				resource.Attribute{
					Name:        "marketplace",
					Description: `(Optional) Select an image from Azure's Marketplace image catalogue.`,
				},
				resource.Attribute{
					Name:        "publisher",
					Description: `(Optional) Image publisher.`,
				},
				resource.Attribute{
					Name:        "offer",
					Description: `(Optional) Image name.`,
				},
				resource.Attribute{
					Name:        "sku",
					Description: `(Optional) Image Stock Keeping Unit (which is the specific version of the image).`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Optional, Default: ` + "`" + `latest` + "`" + `) Image version.`,
				},
				resource.Attribute{
					Name:        "vm_sizes",
					Description: `(Optional) The types of virtual machines that may or may not be a part of the Ocean cluster.`,
				},
				resource.Attribute{
					Name:        "whitelist",
					Description: `(Optional) VM types allowed in the Ocean cluster.`,
				},
				resource.Attribute{
					Name:        "strategy",
					Description: `(Optional) The Ocean AKS strategy object.`,
				},
				resource.Attribute{
					Name:        "fallback_to_ondemand",
					Description: `(Optional) If no spot instance markets are available, enable Ocean to launch on-demand instances instead.`,
				},
				resource.Attribute{
					Name:        "spot_percentage",
					Description: `(Optional) Percentage of Spot VMs to maintain.`,
				},
				resource.Attribute{
					Name:        "health",
					Description: `(Optional) The Ocean AKS Health object.`,
				},
				resource.Attribute{
					Name:        "grace_period",
					Description: `(Optional, Default: ` + "`" + `600` + "`" + `) The amount of time to wait, in seconds, from the moment the instance has launched before monitoring its health checks.`,
				},
				resource.Attribute{
					Name:        "autoscaler",
					Description: `(Optional) The Ocean Kubernetes Autoscaler object.`,
				},
				resource.Attribute{
					Name:        "autoscale_is_enabled",
					Description: `(Optional) Enable the Ocean Kubernetes Autoscaler.`,
				},
				resource.Attribute{
					Name:        "autoscale_down",
					Description: `(Optional) Auto Scaling scale down operations.`,
				},
				resource.Attribute{
					Name:        "max_scale_down_percentage",
					Description: `(Optional) Would represent the maximum % to scale-down.`,
				},
				resource.Attribute{
					Name:        "resource_limits",
					Description: `(Optional) Optionally set upper and lower bounds on the resource usage of the cluster.`,
				},
				resource.Attribute{
					Name:        "max_vcpu",
					Description: `(Optional) The maximum cpu in vCpu units that can be allocated to the cluster.`,
				},
				resource.Attribute{
					Name:        "max_memory_gib",
					Description: `(Optional) The maximum memory in GiB units that can be allocated to the cluster.`,
				},
				resource.Attribute{
					Name:        "autoscale_headroom",
					Description: `(Optional) Spare Resource Capacity Management feature enables fast assignment of Pods without having to wait for new resources to be launched.`,
				},
				resource.Attribute{
					Name:        "automatic",
					Description: `(Optional) Automatic headroom configuration.`,
				},
				resource.Attribute{
					Name:        "is_enabled",
					Description: `(Optional) Enable automatic headroom. When set to ` + "`" + `true` + "`" + `, Ocean configures and optimizes headroom automatically.`,
				},
				resource.Attribute{
					Name:        "percentage",
					Description: `(Optional) Optionally set a number between 0-100 to control the percentage of total cluster resources dedicated to headroom. Relevant when ` + "`" + `isEnabled` + "`" + ` is toggled on.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "spotinst_ocean_aks_virtual_node_group",
			Category:         "Ocean",
			ShortDescription: `Provides a Spotinst Ocean Virtual Node Group resource using AKS.`,
			Description:      ``,
			Keywords: []string{
				"ocean",
				"aks",
				"virtual",
				"node",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ocean_id",
					Description: `(Required) The Ocean cluster ID.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Set name for the virtual node group.`,
				},
				resource.Attribute{
					Name:        "zones",
					Description: `(Optional) An Array holding Availability Zones, this configures the availability zones the Ocean may launch instances in per VNG.`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `(Optional) Additional labels for the virtual node group. Only custom user labels are allowed. Kubernetes built-in labels and Spot internal labels are not allowed.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) The label key.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Optional) The label value.`,
				},
				resource.Attribute{
					Name:        "taint",
					Description: `(Optional) Additional taints for the virtual node group. Only custom user labels are allowed. Kubernetes built-in labels and Spot internal labels are not allowed.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Optional) The taint key.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Optional) The taint value.`,
				},
				resource.Attribute{
					Name:        "effect",
					Description: `(Optional) The effect of the taint. Valid values: ` + "`" + `"NoSchedule"` + "`" + `, ` + "`" + `"PreferNoSchedule"` + "`" + `, ` + "`" + `"NoExecute"` + "`" + `, ` + "`" + `"PreferNoExecute"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "resource_limits",
					Description: `(Optional).`,
				},
				resource.Attribute{
					Name:        "max_instance_count",
					Description: `(Optional) Option to set a maximum number of instances per virtual node group. If set, value must be greater than or equal to 0.`,
				},
				resource.Attribute{
					Name:        "autoscale",
					Description: `(Optional).`,
				},
				resource.Attribute{
					Name:        "auto_headroom_percentage",
					Description: `(Optional) Number between 0-200 to control the headroom % of the specific Virtual Node Group. Effective when ` + "`" + `cluster.autoScaler.headroom.automatic.is_enabled` + "`" + ` = true is set on the Ocean cluster.`,
				},
				resource.Attribute{
					Name:        "autoscale_headroom",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "cpu_per_unit",
					Description: `(Optional) Configure the number of CPUs to allocate for the headroom. CPUs are denoted in millicores, where 1000 millicores = 1 vCPU.`,
				},
				resource.Attribute{
					Name:        "gpu_per_unit",
					Description: `(Optional) How many GPU cores should be allocated for headroom unit.`,
				},
				resource.Attribute{
					Name:        "memory_per_unit",
					Description: `(Optional) Configure the amount of memory (MiB) to allocate the headroom.`,
				},
				resource.Attribute{
					Name:        "num_of_units",
					Description: `(Required) The number of headroom units to maintain, where each unit has the defined CPU, memory and GPU.`,
				},
				resource.Attribute{
					Name:        "launch_specification",
					Description: `(Optional).`,
				},
				resource.Attribute{
					Name:        "os_disk",
					Description: `(Optional) Specify OS disk specification other than default.`,
				},
				resource.Attribute{
					Name:        "size_gb",
					Description: `(Required) The size of the OS disk in GB, Required if dataDisks is specified.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) The type of the OS disk. Valid values: ` + "`" + `"Standard_LRS"` + "`" + `, ` + "`" + `"Premium_LRS"` + "`" + `, ` + "`" + `"StandardSSD_LRS"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "utilize_ephemeral_storage",
					Description: `(Optional) Flag to enable/disable the Ephemeral OS Disk utilization.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `(Optional) Additional key-value pairs to be used to tag the VMs in the virtual node group.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Optional) Tag Key for Vms in the cluster.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Optional) Tag Value for VMs in the cluster.`,
				},
				resource.Attribute{
					Name:        "max_pods",
					Description: `(Optional) The maximum number of pods per node in an AKS cluster.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "spotinst_ocean_aws",
			Category:         "Ocean",
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
					Description: `(Required) A unique identifier used for connecting the Ocean SaaS platform and the Kubernetes cluster. Typically, the cluster name is used as its identifier.`,
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
					Description: `(Required) A comma-separated list of subnet identifiers for the Ocean cluster. Subnet IDs should be configured with auto assign public IP.`,
				},
				resource.Attribute{
					Name:        "instanceTypes",
					Description: `(Optional) The type of instances that may or may not be a part of the Ocean cluster.`,
				},
				resource.Attribute{
					Name:        "whitelist",
					Description: `(Optional) Instance types allowed in the Ocean cluster. Cannot be configured if ` + "`" + `blacklist` + "`" + ` is configured.`,
				},
				resource.Attribute{
					Name:        "blacklist",
					Description: `(Optional) Instance types not allowed in the Ocean cluster. Cannot be configured if ` + "`" + `whitelist` + "`" + ` is configured.`,
				},
				resource.Attribute{
					Name:        "filters",
					Description: `(Optional) List of filters. The Instance types that match with all filters compose the Ocean's whitelist parameter. Cannot be configured together with whitelist/blacklist.`,
				},
				resource.Attribute{
					Name:        "architectures",
					Description: `(Optional) The filtered instance types will support at least one of the architectures from this list.`,
				},
				resource.Attribute{
					Name:        "categories",
					Description: `(Optional) The filtered instance types will belong to one of the categories types from this list.`,
				},
				resource.Attribute{
					Name:        "disk_types",
					Description: `(Optional) The filtered instance types will have one of the disk type from this list.`,
				},
				resource.Attribute{
					Name:        "exclude_families",
					Description: `(Optional) Types belonging to a family from the ExcludeFamilies will not be available for scaling (asterisk wildcard is also supported). For example, C`,
				},
				resource.Attribute{
					Name:        "exclude_metal",
					Description: `(Optional, Default: false) In case excludeMetal is set to true, metal types will not be available for scaling.`,
				},
				resource.Attribute{
					Name:        "hypervisor",
					Description: `(Optional) The filtered instance types will have a hypervisor type from this list.`,
				},
				resource.Attribute{
					Name:        "include_families",
					Description: `(Optional) Types belonging to a family from the IncludeFamilies will be available for scaling (asterisk wildcard is also supported). For example, C`,
				},
				resource.Attribute{
					Name:        "is_ena_supported",
					Description: `(Optional) Ena is supported or not.`,
				},
				resource.Attribute{
					Name:        "max_gpu",
					Description: `(Optional) Maximum total number of GPUs.`,
				},
				resource.Attribute{
					Name:        "max_memory_gib",
					Description: `(Optional) Maximum amount of Memory (GiB).`,
				},
				resource.Attribute{
					Name:        "max_network_performance",
					Description: `(Optional) Maximum Bandwidth in Gib/s of network performance.`,
				},
				resource.Attribute{
					Name:        "max_vcpu",
					Description: `(Optional) Maximum number of vcpus available.`,
				},
				resource.Attribute{
					Name:        "min_enis",
					Description: `(Optional) Minimum number of network interfaces (ENIs).`,
				},
				resource.Attribute{
					Name:        "min_gpu",
					Description: `(Optional) Minimum total number of GPUs.`,
				},
				resource.Attribute{
					Name:        "min_memory_gib",
					Description: `(Optional) Minimum amount of Memory (GiB).`,
				},
				resource.Attribute{
					Name:        "min_network_performance",
					Description: `(Optional) Minimum Bandwidth in Gib/s of network performance.`,
				},
				resource.Attribute{
					Name:        "min_vcpu",
					Description: `(Optional) Minimum number of vcpus available.`,
				},
				resource.Attribute{
					Name:        "root_device_types",
					Description: `(Optional) The filtered instance types will have a root device types from this list.`,
				},
				resource.Attribute{
					Name:        "virtualization_types",
					Description: `(Optional) The filtered instance types will support at least one of the virtualization types from this list.`,
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
					Name:        "associate_ipv6_address",
					Description: `(Optional, Default: ` + "`" + `false` + "`" + `) Configure IPv6 address allocation.`,
				},
				resource.Attribute{
					Name:        "root_volume_size",
					Description: `(Optional) The size (in Gb) to allocate for the root volume. Minimum ` + "`" + `20` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "monitoring",
					Description: `(Optional) Enable detailed monitoring for cluster. Flag will enable Cloud Watch detailed monitoring (one minute increments). Note: there are additional hourly costs for this service based on the region used.`,
				},
				resource.Attribute{
					Name:        "ebs_optimized",
					Description: `(Optional) Enable EBS optimized for cluster. Flag will enable optimized capacity for high bandwidth connectivity to the EB service for non EBS optimized instance types. For instances that are EBS optimized this flag will be ignored.`,
				},
				resource.Attribute{
					Name:        "use_as_template_only",
					Description: `(Optional, Default: false) launch specification defined on the Ocean object will function only as a template for virtual node groups. When set to true, on Ocean resource creation please make sure your custom VNG has an initial_nodes parameter to create nodes for your VNG.`,
				},
				resource.Attribute{
					Name:        "load_balancers",
					Description: `(Optional) - Array of load balancer objects to add to ocean cluster`,
				},
				resource.Attribute{
					Name:        "arn",
					Description: `(Optional) Required if type is set to ` + "`" + `TARGET_GROUP` + "`" + ``,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Required if type is set to ` + "`" + `CLASSIC` + "`" + ``,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) Can be set to ` + "`" + `CLASSIC` + "`" + ` or ` + "`" + `TARGET_GROUP` + "`" + ``,
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
					Description: `(Optional) The tag value.`,
				},
				resource.Attribute{
					Name:        "fallback_to_ondemand",
					Description: `(Optional, Default: ` + "`" + `true` + "`" + `) If not Spot instance markets are available, enable Ocean to launch On-Demand instances instead.`,
				},
				resource.Attribute{
					Name:        "utilize_reserved_instances",
					Description: `(Optional, Default ` + "`" + `true` + "`" + `) If Reserved instances exist, Ocean will utilize them before launching Spot instances.`,
				},
				resource.Attribute{
					Name:        "draining_timeout",
					Description: `(Optional) The time in seconds, the instance is allowed to run while detached from the ELB. This is to allow the instance time to be drained from incoming TCP connections before terminating it, during a scale down operation.`,
				},
				resource.Attribute{
					Name:        "grace_period",
					Description: `(Optional, Default: 600) The amount of time, in seconds, after the instance has launched to start checking its health.`,
				},
				resource.Attribute{
					Name:        "spot_percentage",
					Description: `(Optional; Required if not using ` + "`" + `ondemand_count` + "`" + `) The percentage of Spot instances that would spin up from the ` + "`" + `desired_capacity` + "`" + ` number.`,
				},
				resource.Attribute{
					Name:        "utilize_commitments",
					Description: `(Optional, Default false) If savings plans exist, Ocean will utilize them before launching Spot instances.`,
				},
				resource.Attribute{
					Name:        "spread_nodes_by",
					Description: `(Optional, Default: ` + "`" + `count` + "`" + `) Ocean will spread the nodes across markets by this value. Possible values: ` + "`" + `vcpu` + "`" + ` or ` + "`" + `count` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "instance_metadata_options",
					Description: `(Optional) Ocean instance metadata options object for IMDSv2.`,
				},
				resource.Attribute{
					Name:        "http_tokens",
					Description: `(Required) Determines if a signed token is required or not. Valid values: ` + "`" + `optional` + "`" + ` or ` + "`" + `required` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "http_put_response_hop_limit",
					Description: `(Optional) An integer from 1 through 64. The desired HTTP PUT response hop limit for instance metadata requests. The larger the number, the further the instance metadata requests can travel.`,
				},
				resource.Attribute{
					Name:        "block_device_mappings",
					Description: `(Optional) Object. Array list of block devices that are exposed to the instance, specify either virtual devices and EBS volumes.`,
				},
				resource.Attribute{
					Name:        "device_name",
					Description: `(Optional) String. Set device name. (Example: ` + "`" + `/dev/xvda` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "delete_on_termination",
					Description: `(Optional) Boolean. Flag to delete the EBS on instance termination.`,
				},
				resource.Attribute{
					Name:        "encrypted",
					Description: `(Optional) Boolean. Enables [EBS encryption](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSEncryption.html) on the volume.`,
				},
				resource.Attribute{
					Name:        "iops",
					Description: `(Required for requests to create io1 volumes; it is not used in requests to create ` + "`" + `gp2` + "`" + `, ` + "`" + `st1` + "`" + `, ` + "`" + `sc1` + "`" + `, or standard volumes) Int. The number of I/O operations per second (IOPS) that the volume supports.`,
				},
				resource.Attribute{
					Name:        "kms_key_id",
					Description: `(Optional) String. Identifier (key ID, key alias, ID ARN, or alias ARN) for a customer managed CMK under which the EBS volume is encrypted.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `(Optional) (Optional) String. The Snapshot ID to mount by.`,
				},
				resource.Attribute{
					Name:        "volume_type",
					Description: `(Optional, Default: ` + "`" + `"standard"` + "`" + `) String. The type of the volume. (Example: ` + "`" + `gp2` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "volume_size",
					Description: `(Optional) Int. The size, in GB of the volume.`,
				},
				resource.Attribute{
					Name:        "throughput",
					Description: `(Optional) The amount of data transferred to or from a storage device per second, you can use this param just in a case that ` + "`" + `volume_type` + "`" + ` = ` + "`" + `gp3` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "dynamic_volume_size",
					Description: `(Optional) Object. Set dynamic volume size properties. When using this object, you cannot use volumeSize. You must use one or the other.`,
				},
				resource.Attribute{
					Name:        "availability_vs_cost",
					Description: `(Optional, Default: ` + "`" + `balanced` + "`" + `) You can control the approach that Ocean takes while launching nodes by configuring this value. Possible values: ` + "`" + `costOriented` + "`" + `,` + "`" + `balanced` + "`" + `,` + "`" + `cheapest` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "logging",
					Description: `(Optional) Logging configuration.`,
				},
				resource.Attribute{
					Name:        "export",
					Description: `(Optional) Logging Export configuration.`,
				},
				resource.Attribute{
					Name:        "s3",
					Description: `(Optional) Exports your cluster's logs to the S3 bucket and subdir configured on the S3 data integration given.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Required) The identifier of The S3 data integration to export the logs to. ## Auto Scaler`,
				},
				resource.Attribute{
					Name:        "autoscaler",
					Description: `(Optional) Describes the Ocean Kubernetes Auto Scaler.`,
				},
				resource.Attribute{
					Name:        "autoscale_is_enabled",
					Description: `(Optional, Default: ` + "`" + `true` + "`" + `) Enable the Ocean Kubernetes Auto Scaler.`,
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
					Name:        "auto_headroom_percentage",
					Description: `(Optional) Set the auto headroom percentage (a number in the range [0, 200]) which controls the percentage of headroom from the cluster. Relevant only when ` + "`" + `autoscale_is_auto_config` + "`" + ` toggled on.`,
				},
				resource.Attribute{
					Name:        "enable_automatic_and_manual_headroom",
					Description: `(Optional, Default: ` + "`" + `false` + "`" + `) enables automatic and manual headroom to work in parallel. When set to false, automatic headroom overrides all other headroom definitions manually configured, whether they are at cluster or VNG level.`,
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
					Description: `(Optional) Optionally configure the number of GPUs to allocate the headroom.`,
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
					Name:        "max_scale_down_percentage",
					Description: `(Optional) Would represent the maximum % to scale-down. Number between 1-100.`,
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
					Description: `(Optional) The maximum memory in GiB units that can be allocated to the cluster.`,
				},
				resource.Attribute{
					Name:        "extended_resource_definitions",
					Description: `(Optional) List of Ocean extended resource definitions to use in this cluster. ` + "`" + `` + "`" + `` + "`" + `hcl autoscaler { autoscale_is_enabled = true autoscale_is_auto_config = true auto_headroom_percentage = 100 autoscale_cooldown = 300 enable_automatic_and_manual_headroom = false autoscale_headroom { cpu_per_unit = 1024 gpu_per_unit = 0 memory_per_unit = 512 num_of_units = 2 } autoscale_down { max_scale_down_percentage = 60 } resource_limits { max_vcpu = 1024 max_memory_gib = 1500 } extended_resource_definitions = ["erd-abc123"] } ` + "`" + `` + "`" + `` + "`" + ` <a id="update-policy"></a> ## Update Policy`,
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
					Name:        "conditioned_roll",
					Description: `(Optional, Default: false) Spot will perform a cluster Roll in accordance with a relevant modification of the cluster’s settings. When set to true , only specific changes in the cluster’s configuration will trigger a cluster roll (such as AMI, Key Pair, user data, instance types, load balancers, etc).`,
				},
				resource.Attribute{
					Name:        "auto_apply_tags",
					Description: `(Optional, Default: false) will update instance tags on the fly without rolling the cluster.`,
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
					Name:        "launch_spec_ids",
					Description: `(Optional) List of virtual node group identifiers to be rolled.`,
				},
				resource.Attribute{
					Name:        "batch_min_healthy_percentage",
					Description: `(Optional) Default: 50. Indicates the threshold of minimum healthy instances in single batch. If the amount of healthy instances in single batch is under the threshold, the cluster roll will fail. If exists, the parameter value will be in range of 1-100. In case of null as value, the default value in the backend will be 50%. Value of param should represent the number in percentage (%) of the batch.`,
				},
				resource.Attribute{
					Name:        "respect_pdb",
					Description: `(Optional, Default: false) During the roll, if the parameter is set to True we honor PDB during the instance replacement. ` + "`" + `` + "`" + `` + "`" + `hcl update_policy { should_roll = false conditioned_roll = true auto_apply_tags = true roll_config { batch_size_percentage = 33 launch_spec_ids = ["ols-1a2b3c4d"] batch_min_healthy_percentage = 20 respect_pdb = true } } ` + "`" + `` + "`" + `` + "`" + ` <a id="scheduled-task"></a> ## Scheduled Task`,
				},
				resource.Attribute{
					Name:        "scheduled_task",
					Description: `(Optional) Set scheduling object.`,
				},
				resource.Attribute{
					Name:        "shutdown_hours",
					Description: `(Optional) Set shutdown hours for cluster object.`,
				},
				resource.Attribute{
					Name:        "is_enabled",
					Description: `(Optional) Toggle the shutdown hours task. (Example: ` + "`" + `true` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "time_windows",
					Description: `(Required) Set time windows for shutdown hours. Specify a list of ` + "`" + `timeWindows` + "`" + ` with at least one time window Each string is in the format of: ` + "`" + `ddd:hh:mm-ddd:hh:mm` + "`" + ` where ` + "`" + `ddd` + "`" + ` = day of week = Sun | Mon | Tue | Wed | Thu | Fri | Sat, ` + "`" + `hh` + "`" + ` = hour 24 = 0 -23, ` + "`" + `mm` + "`" + ` = minute = 0 - 59. Time windows should not overlap. Required if ` + "`" + `cluster.scheduling.isEnabled` + "`" + ` is ` + "`" + `true` + "`" + `. (Example: ` + "`" + `Fri:15:30-Wed:14:30` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "tasks",
					Description: `(Optional) The scheduling tasks for the cluster.`,
				},
				resource.Attribute{
					Name:        "is_enabled",
					Description: `(Required) Describes whether the task is enabled. When true the task should run when false it should not run. Required for ` + "`" + `cluster.scheduling.tasks` + "`" + ` object.`,
				},
				resource.Attribute{
					Name:        "cron_expression",
					Description: `(Required) A valid cron expression. The cron is running in UTC time zone and is in Unix cron format Cron Expression Validator Script. Only one of ` + "`" + `frequency` + "`" + ` or ` + "`" + `cronExpression` + "`" + ` should be used at a time. Required for ` + "`" + `cluster.scheduling.tasks` + "`" + ` object. (Example: ` + "`" + `0 1`,
				},
				resource.Attribute{
					Name:        "task_type",
					Description: `(Required) Valid values: ` + "`" + `clusterRoll` + "`" + `. Required for ` + "`" + `cluster.scheduling.tasks` + "`" + ` object. (Example: ` + "`" + `clusterRoll` + "`" + `). ` + "`" + `` + "`" + `` + "`" + `hcl scheduled_task { shutdown_hours { is_enabled = true time_windows = [ "Fri:15:30-Sat:13:30", "Sun:15:30-Mon:13:30", ] } tasks { is_enabled = false cron_expression = "`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The Cluster ID. <a id="import"></a> ## Import Clusters can be imported using the Ocean ` + "`" + `id` + "`" + `, e.g., ` + "`" + `` + "`" + `` + "`" + `hcl $ terraform import spotinst_ocean_aws.this o-12345678 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The Cluster ID. <a id="import"></a> ## Import Clusters can be imported using the Ocean ` + "`" + `id` + "`" + `, e.g., ` + "`" + `` + "`" + `` + "`" + `hcl $ terraform import spotinst_ocean_aws.this o-12345678 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "spotinst_ocean_aws_extended_resource_definition",
			Category:         "Ocean",
			ShortDescription: `Manages an Ocean extended resource definition resource.`,
			Description:      ``,
			Keywords: []string{
				"ocean",
				"aws",
				"extended",
				"resource",
				"definition",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The extended resource name as should be requested by your pods and registered to the nodes. Cannot be updated. The name should be a valid Kubernetes extended resource name.`,
				},
				resource.Attribute{
					Name:        "resource_mapping",
					Description: `(Required) A mapping between AWS instanceType or`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The Spotinst Extended Resource Definition ID.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The Spotinst Extended Resource Definition ID.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "spotinst_ocean_aws_launch_spec",
			Category:         "Ocean",
			ShortDescription: `Provides a Spotinst Virtual Node Group resource using AWS.`,
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
					Description: `(Required) The ID of the Ocean cluster.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the Virtual Node Group.`,
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
					Name:        "images",
					Description: `Array of objects (Image object, containing the id of the image used to launch instances.) You can configure VNG with either the imageId or images objects, but not both simultaneously. For each architecture type (amd64, arm64) only one AMI is allowed. Valid values: null, or an array with at least one element.`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `Identifier of the image in AWS. Valid values: any string which is not empty or null.`,
				},
				resource.Attribute{
					Name:        "iam_instance_profile",
					Description: `(Optional) The ARN or name of an IAM instance profile to associate with launched instances.`,
				},
				resource.Attribute{
					Name:        "security_groups",
					Description: `(Optional) Optionally adds security group IDs.`,
				},
				resource.Attribute{
					Name:        "subnet_ids",
					Description: `(Optional) A list of subnet IDs.`,
				},
				resource.Attribute{
					Name:        "instance_types",
					Description: `(Optional) A list of instance types allowed to be provisioned for pods pending under the specified launch specification. The list overrides the list defined for the cluster.`,
				},
				resource.Attribute{
					Name:        "preferred_spot_types",
					Description: `(Optional) A list of instance types. Takes the preferred types into consideration while maintaining a variety of machine types running for optimized distribution.`,
				},
				resource.Attribute{
					Name:        "root_volume_size",
					Description: `(Optional) Set root volume size (in GB).`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A key/value mapping of tags to assign to the resource.`,
				},
				resource.Attribute{
					Name:        "associate_public_ip_address",
					Description: `(Optional, Default: ` + "`" + `false` + "`" + `) Configure public IP address allocation.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) Optionally adds labels to instances launched in the cluster.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) The label key.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) The label value.`,
				},
				resource.Attribute{
					Name:        "taints",
					Description: `(Optional) Optionally adds labels to instances launched in the cluster.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) The taint key.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) The taint value.`,
				},
				resource.Attribute{
					Name:        "effect",
					Description: `(Required) The effect of the taint. Valid values: ` + "`" + `"NoSchedule"` + "`" + `, ` + "`" + `"PreferNoSchedule"` + "`" + `, ` + "`" + `"NoExecute"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "elastic_ip_pool",
					Description: `(Optional) Assign an Elastic IP to the instances spun by the Virtual Node Group. Can be null.`,
				},
				resource.Attribute{
					Name:        "tag_selector",
					Description: `(Optional) A key-value pair, which defines an Elastic IP from the customer pool. Can be null.`,
				},
				resource.Attribute{
					Name:        "tag_key",
					Description: `(Required) Elastic IP tag key. The Virtual Node Group will consider all Elastic IPs tagged with this tag as a part of the Elastic IP pool to use.`,
				},
				resource.Attribute{
					Name:        "tag_value",
					Description: `(Optional) Elastic IP tag value. Can be null.`,
				},
				resource.Attribute{
					Name:        "block_device_mappings",
					Description: `(Optional) Object. Array list of block devices that are exposed to the instance, specify either virtual devices and EBS volumes.`,
				},
				resource.Attribute{
					Name:        "device_name",
					Description: `(Optional) String. Set device name. (Example: ` + "`" + `/dev/xvda` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "delete_on_termination",
					Description: `(Optional) Boolean. Flag to delete the EBS on instance termination.`,
				},
				resource.Attribute{
					Name:        "encrypted",
					Description: `(Optional) Boolean. Enables [EBS encryption](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSEncryption.html) on the volume.`,
				},
				resource.Attribute{
					Name:        "iops",
					Description: `(Required for requests to create io1 volumes; it is not used in requests to create ` + "`" + `gp2` + "`" + `, ` + "`" + `st1` + "`" + `, ` + "`" + `sc1` + "`" + `, or standard volumes) Int. The number of I/O operations per second (IOPS) that the volume supports.`,
				},
				resource.Attribute{
					Name:        "kms_key_id",
					Description: `(Optional) String. Identifier (key ID, key alias, ID ARN, or alias ARN) for a customer managed CMK under which the EBS volume is encrypted.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `(Optional) (Optional) String. The Snapshot ID to mount by.`,
				},
				resource.Attribute{
					Name:        "volume_type",
					Description: `(Optional, Default: ` + "`" + `"standard"` + "`" + `) String. The type of the volume. (Example: ` + "`" + `gp2` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "volume_size",
					Description: `(Optional) Int. The size, in GB of the volume.`,
				},
				resource.Attribute{
					Name:        "throughput",
					Description: `(Optional) The amount of data transferred to or from a storage device per second, you can use this param just in a case that ` + "`" + `volume_type` + "`" + ` = ` + "`" + `gp3` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "dynamic_volume_size",
					Description: `(Optional) Object. Set dynamic volume size properties. When using this object, you cannot use volumeSize. You must use one or the other.`,
				},
				resource.Attribute{
					Name:        "no_device",
					Description: `(Optional) String. Suppresses the specified device included in the block device mapping of the AMI.`,
				},
				resource.Attribute{
					Name:        "autoscale_headrooms_automatic",
					Description: `(Optional) Set automatic headroom per launch spec.`,
				},
				resource.Attribute{
					Name:        "auto_headroom_percentage",
					Description: `(Optional) Number between 0-200 to control the headroom % of the specific Virtual Node Group. Effective when cluster.autoScaler.headroom.automatic.` + "`" + `is_enabled` + "`" + ` = true is set on the Ocean cluster.`,
				},
				resource.Attribute{
					Name:        "autoscale_headrooms",
					Description: `(Optional) Set custom headroom per Virtual Node Group. Provide a list of headrooms object.`,
				},
				resource.Attribute{
					Name:        "num_of_units",
					Description: `(Required) The number of units to retain as headroom, where each unit has the defined headroom CPU, memory and GPU.`,
				},
				resource.Attribute{
					Name:        "cpu_per_unit",
					Description: `(Optional) Optionally configure the number of CPUs to allocate for each headroom unit. CPUs are denoted in millicores, where 1000 millicores = 1 vCPU.`,
				},
				resource.Attribute{
					Name:        "gpu_per_unit",
					Description: `(Optional) Optionally configure the number of GPUS to allocate for each headroom unit.`,
				},
				resource.Attribute{
					Name:        "memory_per_unit",
					Description: `(Optional) Optionally configure the amount of memory (MiB) to allocate for each headroom unit.`,
				},
				resource.Attribute{
					Name:        "autoscale_down",
					Description: `(Optional) Auto Scaling scale down operations.`,
				},
				resource.Attribute{
					Name:        "max_scale_down_percentage",
					Description: `(Optional) The maximum percentage allowed to scale down in a single scaling action on the nodes running in a specific VNG. Allowed only if maxScaleDownPercentage is set to null at the cluster level. Number between [0.1-100].`,
				},
				resource.Attribute{
					Name:        "resource_limits",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "max_instance_count",
					Description: `(Optional) Set a maximum number of instances per Virtual Node Group. Can be null. If set, value must be greater than or equal to 0.`,
				},
				resource.Attribute{
					Name:        "min_instance_count",
					Description: `(Optional) Set a minimum number of instances per Virtual Node Group. Can be null. If set, value must be greater than or equal to 0.`,
				},
				resource.Attribute{
					Name:        "strategy",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "spot_percentage",
					Description: `(Optional; if not using ` + "`" + `spot_percentege` + "`" + ` under ` + "`" + `ocean strategy` + "`" + `) When set, Ocean will proactively try to maintain as close as possible to the percentage of Spot instances out of all the Virtual Node Group instances.`,
				},
				resource.Attribute{
					Name:        "create_options",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "initial_nodes",
					Description: `(Optional) When set to an integer greater than 0, a corresponding amount of nodes will be launched from the created Virtual Node Group. The parameter is recommended in case the use_as_template_only (in spotinst_ocean_aws resource) is set to true during Ocean resource creation.`,
				},
				resource.Attribute{
					Name:        "delete_options",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "force_delete",
					Description: `(Optional) When set to ` + "`" + `true` + "`" + `, delete even if it is the last Virtual Node Group (also, the default Virtual Node Group must be configured with ` + "`" + `useAsTemlateOnly = true` + "`" + `). Should be set at creation or update, but will be used only at deletion.`,
				},
				resource.Attribute{
					Name:        "delete_nodes",
					Description: `(Optional) When set to "true", all instances belonging to the deleted launch specification will be drained, detached, and terminated.`,
				},
				resource.Attribute{
					Name:        "scheduling_task",
					Description: `(Optional) Used to define scheduled tasks such as a manual headroom update.`,
				},
				resource.Attribute{
					Name:        "is_enabled",
					Description: `(Required) Describes whether the task is enabled. When True, the task runs. When False, it does not run.`,
				},
				resource.Attribute{
					Name:        "cron_expression",
					Description: `(Required) A valid cron expression. For example : "`,
				},
				resource.Attribute{
					Name:        "task_type",
					Description: `(Required) The activity that you are scheduling. Valid values: "manualHeadroomUpdate".`,
				},
				resource.Attribute{
					Name:        "task_headroom",
					Description: `(Optional) The config of this scheduled task. Depends on the value of taskType.`,
				},
				resource.Attribute{
					Name:        "num_of_units",
					Description: `(Required) The number of units to retain as headroom, where each unit has the defined headroom CPU, memory and GPU.`,
				},
				resource.Attribute{
					Name:        "cpu_per_unit",
					Description: `(Optional) Optionally configure the number of CPUs to allocate for each headroom unit. CPUs are denoted in millicores, where 1000 millicores = 1 vCPU.`,
				},
				resource.Attribute{
					Name:        "gpu_per_unit",
					Description: `(Optional) Optionally configure the number of GPUS to allocate for each headroom unit.`,
				},
				resource.Attribute{
					Name:        "memory_per_unit",
					Description: `(Optional) Optionally configure the amount of memory (MiB) to allocate for each headroom unit.`,
				},
				resource.Attribute{
					Name:        "scheduling_shutdown_hours",
					Description: `(Optional) Used to specify times that the nodes in the virtual node group will be taken down.`,
				},
				resource.Attribute{
					Name:        "time_windows",
					Description: `(Required ) The times that the shutdown hours will apply.`,
				},
				resource.Attribute{
					Name:        "is_enabled",
					Description: `(Optional) Flag to enable or disable the shutdown hours mechanism. When False, the mechanism is deactivated, and the virtual node group remains in its current state.`,
				},
				resource.Attribute{
					Name:        "instance_metadata_options",
					Description: `(Optional) Ocean instance metadata options object for IMDSv2.`,
				},
				resource.Attribute{
					Name:        "http_tokens",
					Description: `(Required) Determines if a signed token is required or not. Valid values: ` + "`" + `optional` + "`" + ` or ` + "`" + `required` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "http_put_response_hop_limit",
					Description: `(Optional) An integer from 1 through 64. The desired HTTP PUT response hop limit for instance metadata requests. The larger the number, the further the instance metadata requests can travel. <a id="update-policy"></a> ## Update Policy`,
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
					Description: `(Required) Holds the roll configuration.`,
				},
				resource.Attribute{
					Name:        "batch_size_percentage",
					Description: `(Required) Sets the percentage of the instances to deploy in each batch. ` + "`" + `` + "`" + `` + "`" + `hcl update_policy { should_roll = false roll_config { batch_size_percentage = 33 } } ` + "`" + `` + "`" + `` + "`" + ` ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The Virtual Node Group ID.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The Virtual Node Group ID.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "spotinst_ocean_ecs",
			Category:         "Ocean",
			ShortDescription: `Provides a Spotinst Ocean ECS resource using AWS.`,
			Description:      ``,
			Keywords: []string{
				"ocean",
				"ecs",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The Ocean cluster name.`,
				},
				resource.Attribute{
					Name:        "cluster_name",
					Description: `(Required) The name of the ECS cluster.`,
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
					Description: `(Required) A comma-separated list of subnet identifiers for the Ocean cluster. Subnet IDs should be configured with auto assign public ip.`,
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
					Description: `(Optional) The tag value.`,
				},
				resource.Attribute{
					Name:        "instanceTypes",
					Description: `(Optional) The type of instances that may or may not be a part of the Ocean cluster.`,
				},
				resource.Attribute{
					Name:        "whitelist",
					Description: `(Optional) Instance types allowed in the Ocean cluster. Cannot be configured if ` + "`" + `blacklist` + "`" + `/` + "`" + `filters` + "`" + ` is configured.`,
				},
				resource.Attribute{
					Name:        "blacklist",
					Description: `(Optional) Instance types not allowed in the Ocean cluster. Cannot be configured if ` + "`" + `whitelist` + "`" + `/` + "`" + `filters` + "`" + ` is configured.`,
				},
				resource.Attribute{
					Name:        "filters",
					Description: `(Optional) List of filters. The Instance types that match with all filters compose the Ocean's whitelist parameter. Cannot be configured together with ` + "`" + `whitelist` + "`" + `/` + "`" + `blacklist` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "architectures",
					Description: `(Optional) The filtered instance types will support at least one of the architectures from this list.`,
				},
				resource.Attribute{
					Name:        "categories",
					Description: `(Optional) The filtered instance types will belong to one of the categories types from this list.`,
				},
				resource.Attribute{
					Name:        "disk_types",
					Description: `(Optional) The filtered instance types will have one of the disk type from this list.`,
				},
				resource.Attribute{
					Name:        "exclude_families",
					Description: `(Optional) Types belonging to a family from the ExcludeFamilies will not be available for scaling (asterisk wildcard is also supported). For example, C`,
				},
				resource.Attribute{
					Name:        "exclude_metal",
					Description: `(Optional, Default: false) In case excludeMetal is set to true, metal types will not be available for scaling.`,
				},
				resource.Attribute{
					Name:        "hypervisor",
					Description: `(Optional) The filtered instance types will have a hypervisor type from this list.`,
				},
				resource.Attribute{
					Name:        "include_families",
					Description: `(Optional) Types belonging to a family from the IncludeFamilies will be available for scaling (asterisk wildcard is also supported). For example, C`,
				},
				resource.Attribute{
					Name:        "is_ena_supported",
					Description: `(Optional) Ena is supported or not.`,
				},
				resource.Attribute{
					Name:        "max_gpu",
					Description: `(Optional) Maximum total number of GPUs.`,
				},
				resource.Attribute{
					Name:        "max_memory_gib",
					Description: `(Optional) Maximum amount of Memory (GiB).`,
				},
				resource.Attribute{
					Name:        "max_network_performance",
					Description: `(Optional) Maximum Bandwidth in Gib/s of network performance.`,
				},
				resource.Attribute{
					Name:        "max_vcpu",
					Description: `(Optional) Maximum number of vcpus available.`,
				},
				resource.Attribute{
					Name:        "min_enis",
					Description: `(Optional) Minimum number of network interfaces (ENIs).`,
				},
				resource.Attribute{
					Name:        "min_gpu",
					Description: `(Optional) Minimum total number of GPUs.`,
				},
				resource.Attribute{
					Name:        "min_memory_gib",
					Description: `(Optional) Minimum amount of Memory (GiB).`,
				},
				resource.Attribute{
					Name:        "min_network_performance",
					Description: `(Optional) Minimum Bandwidth in Gib/s of network performance.`,
				},
				resource.Attribute{
					Name:        "min_vcpu",
					Description: `(Optional) Minimum number of vcpus available.`,
				},
				resource.Attribute{
					Name:        "root_device_types",
					Description: `(Optional) The filtered instance types will have a root device types from this list.`,
				},
				resource.Attribute{
					Name:        "virtualization_types",
					Description: `(Optional) The filtered instance types will support at least one of the virtualization types from this list.`,
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
					Name:        "security_group_ids",
					Description: `(Required) One or more security group ids.`,
				},
				resource.Attribute{
					Name:        "key_pair",
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
					Name:        "utilize_reserved_instances",
					Description: `(Optional, Default ` + "`" + `true` + "`" + `) If Reserved instances exist, Ocean will utilize them before launching Spot instances.`,
				},
				resource.Attribute{
					Name:        "draining_timeout",
					Description: `(Optional) The time in seconds, the instance is allowed to run while detached from the ELB. This is to allow the instance time to be drained from incoming TCP connections before terminating it, during a scale down operation.`,
				},
				resource.Attribute{
					Name:        "monitoring",
					Description: `(Optional) Enable detailed monitoring for cluster. Flag will enable Cloud Watch detailed monitoring (one minute increments). Note: there are additional hourly costs for this service based on the region used.`,
				},
				resource.Attribute{
					Name:        "ebs_optimized",
					Description: `(Optional) Enable EBS optimized for cluster. Flag will enable optimized capacity for high bandwidth connectivity to the EB service for non EBS optimized instance types. For instances that are EBS optimized this flag will be ignored.`,
				},
				resource.Attribute{
					Name:        "use_as_template_only",
					Description: `(Optional, Default: false) launch specification defined on the Ocean object will function only as a template for virtual node groups.`,
				},
				resource.Attribute{
					Name:        "spot_percentage",
					Description: `(Optional) The percentage of Spot instances that would spin up from the ` + "`" + `desired_capacity` + "`" + ` number.`,
				},
				resource.Attribute{
					Name:        "utilize_commitments",
					Description: `(Optional, Default false) If savings plans exist, Ocean will utilize them before launching Spot instances.`,
				},
				resource.Attribute{
					Name:        "availability_vs_cost",
					Description: `(Optional, Default: ` + "`" + `balanced` + "`" + `) You can control the approach that Ocean takes while launching nodes by configuring this value. Possible values: ` + "`" + `costOriented` + "`" + `,` + "`" + `balanced` + "`" + `,` + "`" + `cheapest` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "instance_metadata_options",
					Description: `(Optional) Ocean instance metadata options object for IMDSv2.`,
				},
				resource.Attribute{
					Name:        "http_tokens",
					Description: `(Required) Determines if a signed token is required or not. Valid values: ` + "`" + `optional` + "`" + ` or ` + "`" + `required` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "http_put_response_hop_limit",
					Description: `(Optional) An integer from 1 through 64. The desired HTTP PUT response hop limit for instance metadata requests. The larger the number, the further the instance metadata requests can travel.`,
				},
				resource.Attribute{
					Name:        "logging",
					Description: `(Optional) Logging configuration.`,
				},
				resource.Attribute{
					Name:        "export",
					Description: `(Optional) Logging Export configuration.`,
				},
				resource.Attribute{
					Name:        "s3",
					Description: `(Optional) Exports your cluster's logs to the S3 bucket and subdir configured on the S3 data integration given.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Required) The identifier of The S3 data integration to export the logs to. <a id="block-devices"></a> ## Block Devices`,
				},
				resource.Attribute{
					Name:        "block_device_mappings",
					Description: `(Optional) Object. List of block devices that are exposed to the instance, specify either virtual devices and EBS volumes.`,
				},
				resource.Attribute{
					Name:        "device_name",
					Description: `(Optional) String. Set device name. Example: ` + "`" + `/dev/xvda1` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ebs",
					Description: `(Optional) Object. Set Elastic Block Store properties.`,
				},
				resource.Attribute{
					Name:        "delete_on_termination",
					Description: `(Optional) Boolean. Toggles EBS deletion upon instance termination.`,
				},
				resource.Attribute{
					Name:        "encrypted",
					Description: `(Optional) Boolean. Enables [EBS encryption](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSEncryption.html) on the volume.`,
				},
				resource.Attribute{
					Name:        "iops",
					Description: `(Required for requests to create ` + "`" + `io1` + "`" + ` volumes; it is not used in requests to create ` + "`" + `gp2` + "`" + `, ` + "`" + `st1` + "`" + `, ` + "`" + `sc1` + "`" + `, or standard volumes) Int. The number of I/O operations per second (IOPS) that the volume supports.`,
				},
				resource.Attribute{
					Name:        "kms_key_id",
					Description: `(Optional) String. Identifier (key ID, key alias, ID ARN, or alias ARN) for a customer managed CMK under which the EBS volume is encrypted.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `(Optional) (Optional) String. The snapshot ID to mount by.`,
				},
				resource.Attribute{
					Name:        "volume_type",
					Description: `(Optional, Default: ` + "`" + `standard` + "`" + `) String. The type of the volume. Example: ` + "`" + `gp2` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "volume_size",
					Description: `(Optional) Int. The size (in GB) of the volume.`,
				},
				resource.Attribute{
					Name:        "dynamic_volume_size",
					Description: `(Optional) Object. Set dynamic volume size properties. When using this object, you cannot use volumeSize. You must use one or the other.`,
				},
				resource.Attribute{
					Name:        "base_size",
					Description: `(Required) Int. Initial size for volume. Example: ` + "`" + `50` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "resource",
					Description: `(Required) String. Resource type to increase volume size dynamically by. Valid values: ` + "`" + `CPU` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "size_per_resource_unit",
					Description: `(Required) Int. Additional size (in GB) per resource unit. Example: When the ` + "`" + `baseSize=50` + "`" + `, ` + "`" + `sizePerResourceUnit=20` + "`" + `, and instance with two CPUs is launched, its total disk size will be: 90GB.`,
				},
				resource.Attribute{
					Name:        "no_device",
					Description: `(Optional) String. Suppresses the specified device included in the block device mapping of the AMI.`,
				},
				resource.Attribute{
					Name:        "optimize_images",
					Description: `(Optional) Object. Set auto image update settings.`,
				},
				resource.Attribute{
					Name:        "perform_at",
					Description: `(Required) String. Valid values: "always" "never" "timeWindow".`,
				},
				resource.Attribute{
					Name:        "time_windows",
					Description: `(Optional; Required if not using ` + "`" + `perform_at` + "`" + ` = timeWindow) Array of strings. Set time windows for image update, at least one time window. Each string is in the format of ddd:hh:mm-ddd:hh:mm ddd. Time windows should not overlap.`,
				},
				resource.Attribute{
					Name:        "should_optimize_ecs_ami",
					Description: `(Required) Boolean. Enable auto image (AMI) update for the ECS container instances. The auto update applies for ECS-Optimized AMIs. <a id="auto-scaler"></a> ## Auto Scaler`,
				},
				resource.Attribute{
					Name:        "autoscaler",
					Description: `(Optional) Describes the Ocean ECS autoscaler.`,
				},
				resource.Attribute{
					Name:        "is_enabled",
					Description: `(Optional, Default: ` + "`" + `true` + "`" + `) Enable the Ocean ECS autoscaler.`,
				},
				resource.Attribute{
					Name:        "is_auto_config",
					Description: `(Optional, Default: ` + "`" + `true` + "`" + `) Automatically configure and optimize headroom resources.`,
				},
				resource.Attribute{
					Name:        "cooldown",
					Description: `(Optional, Default: ` + "`" + `null` + "`" + `) Cooldown period between scaling actions.`,
				},
				resource.Attribute{
					Name:        "headroom",
					Description: `(Optional) Spare resource capacity management enabling fast assignment of tasks without waiting for new resources to launch.`,
				},
				resource.Attribute{
					Name:        "cpu_per_unit",
					Description: `(Optional) Optionally configure the number of CPUs to allocate the headroom. CPUs are denoted in millicores, where 1000 millicores = 1 vCPU.`,
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
					Name:        "down",
					Description: `(Optional) Auto Scaling scale down operations.`,
				},
				resource.Attribute{
					Name:        "max_scale_down_percentage",
					Description: `(Optional) Would represent the maximum % to scale-down. Number between 1-100.`,
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
					Description: `(Optional) The maximum memory in GiB units that can be allocated to the cluster.`,
				},
				resource.Attribute{
					Name:        "auto_headroom_percentage",
					Description: `(Optional) The auto-headroom percentage. Set a number between 0-200 to control the headroom % of the cluster. Relevant when ` + "`" + `isAutoConfig` + "`" + `= true.`,
				},
				resource.Attribute{
					Name:        "should_scale_down_non_service_tasks",
					Description: `(Optional, Default: ` + "`" + `false` + "`" + `) Option to scale down non-service tasks. If not set, Ocean does not scale down standalone tasks.`,
				},
				resource.Attribute{
					Name:        "enable_automatic_and_manual_headroom",
					Description: `(Optional) When set to true, both automatic and per custom launch specification manual headroom to be saved concurrently and independently in the cluster. prerequisite: isAutoConfig must be true ` + "`" + `` + "`" + `` + "`" + `hcl autoscaler { is_enabled = false is_auto_config = true cooldown = 300 headroom { cpu_per_unit = 1024 memory_per_unit = 512 num_of_units = 2 } down { max_scale_down_percentage = 20 } resource_limits { max_vcpu = 1024 max_memory_gib = 20 } auto_headroom_percentage = 10 should_scale_down_non_service_tasks = false enable_automatic_and_manual_headroom = true } ` + "`" + `` + "`" + `` + "`" + ` <a id="update-policy"></a> ## Update Policy`,
				},
				resource.Attribute{
					Name:        "update_policy",
					Description: `(Optional) While used, you can control whether the group should perform a deployment after an update to the configuration.`,
				},
				resource.Attribute{
					Name:        "should_roll",
					Description: `(Required) Enables the roll.`,
				},
				resource.Attribute{
					Name:        "conditioned_roll",
					Description: `(Optional, Default: false) Spot will perform a cluster Roll in accordance with a relevant modification of the cluster’s settings. When set to true , only specific changes in the cluster’s configuration will trigger a cluster roll (such as AMI, Key Pair, user data, instance types, load balancers, etc).`,
				},
				resource.Attribute{
					Name:        "auto_apply_tags",
					Description: `(Optional, Default: false) will update instance tags on the fly without rolling the cluster.`,
				},
				resource.Attribute{
					Name:        "roll_config",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "batch_size_percentage",
					Description: `(Required) Sets the percentage of the instances to deploy in each batch.`,
				},
				resource.Attribute{
					Name:        "batch_min_healthy_percentage",
					Description: `(Optional) Default: 50. Indicates the threshold of minimum healthy instances in single batch. If the amount of healthy instances in single batch is under the threshold, the cluster roll will fail. If exists, the parameter value will be in range of 1-100. In case of null as value, the default value in the backend will be 50%. Value of param should represent the number in percentage (%) of the batch. ` + "`" + `` + "`" + `` + "`" + `hcl update_policy { should_roll = false conditioned_roll = true auto_apply_tags = true roll_config { batch_size_percentage = 33 batch_min_healthy_percentage = 20 } } ` + "`" + `` + "`" + `` + "`" + ` <a id="scheduled-tasks"></a> ## Scheduled Tasks`,
				},
				resource.Attribute{
					Name:        "scheduled_task",
					Description: `(Optional) While used, you can control whether the group should perform a deployment after an update to the configuration.`,
				},
				resource.Attribute{
					Name:        "shutdown_hours",
					Description: `(Optional) Set shutdown hours for cluster object.`,
				},
				resource.Attribute{
					Name:        "is_enabled",
					Description: `(Optional) Flag to enable / disable the shutdown hours.`,
				},
				resource.Attribute{
					Name:        "time_windows",
					Description: `(Required) Set time windows for shutdown hours. Specify a list of ` + "`" + `timeWindows` + "`" + ` with at least one time window Each string is in the format of ` + "`" + `ddd:hh:mm-ddd:hh:mm` + "`" + ` (ddd = day of week = Sun | Mon | Tue | Wed | Thu | Fri | Sat hh = hour 24 = 0 -23 mm = minute = 0 - 59). Time windows should not overlap. Required when ` + "`" + `cluster.scheduling.isEnabled` + "`" + ` is true. API Times are in UTC. Example: ` + "`" + `Fri:15:30-Wed:14:30` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "tasks",
					Description: `(Optional) The scheduling tasks for the cluster.`,
				},
				resource.Attribute{
					Name:        "is_enabled",
					Description: `(Required) Describes whether the task is enabled. When true the task should run when false it should not run. Required for ` + "`" + `cluster.scheduling.tasks` + "`" + ` object.`,
				},
				resource.Attribute{
					Name:        "cron_expression",
					Description: `(Required) A valid cron expression. The cron is running in UTC time zone and is in Unix cron format Cron Expression Validator Script. Only one of ` + "`" + `frequency` + "`" + ` or ` + "`" + `cronExpression` + "`" + ` should be used at a time. Required for ` + "`" + `cluster.scheduling.tasks` + "`" + ` object. Example: ` + "`" + `0 1`,
				},
				resource.Attribute{
					Name:        "task_type",
					Description: `(Required) Valid values: "clusterRoll". Required for ` + "`" + `cluster.scheduling.tasks object` + "`" + `. Example: ` + "`" + `clusterRoll` + "`" + `. ` + "`" + `` + "`" + `` + "`" + `hcl scheduled_task { shutdown_hours { is_enabled = false time_windows = ["Fri:15:30-Wed:13:30"] } tasks { is_enabled = false cron_expression = "`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The Spotinst Ocean ID. <a id="import"></a> ## Import Clusters can be imported using the Ocean ` + "`" + `id` + "`" + `, e.g., ` + "`" + `` + "`" + `` + "`" + `hcl $ terraform import spotinst_ocean_ecs.this o-12345678 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The Spotinst Ocean ID. <a id="import"></a> ## Import Clusters can be imported using the Ocean ` + "`" + `id` + "`" + `, e.g., ` + "`" + `` + "`" + `` + "`" + `hcl $ terraform import spotinst_ocean_ecs.this o-12345678 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "spotinst_ocean_ecs_launch_spec",
			Category:         "Ocean",
			ShortDescription: `Provides a Spotinst Ocean ECS Launch Spec resource using AWS.`,
			Description:      ``,
			Keywords: []string{
				"ocean",
				"ecs",
				"launch",
				"spec",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "user_data",
					Description: `(Optional) Base64-encoded MIME user data to make available to the instances.`,
				},
				resource.Attribute{
					Name:        "iam_instance_profile",
					Description: `(Optional) The ARN or name of an IAM instance profile to associate with launched instances.`,
				},
				resource.Attribute{
					Name:        "security_group_ids",
					Description: `(Optional) One or more security group ids.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A key/value mapping of tags to assign to the resource.`,
				},
				resource.Attribute{
					Name:        "instance_types",
					Description: `(Optional) A list of instance types allowed to be provisioned for pods pending under the specified launch specification. The list overrides the list defined for the Ocean cluster.`,
				},
				resource.Attribute{
					Name:        "preferred_spot_types",
					Description: `(Optional) When Ocean scales up instances, it takes your preferred types into consideration while maintaining a variety of machine types running for optimized distribution.`,
				},
				resource.Attribute{
					Name:        "subnet_ids",
					Description: `(Optional) Set subnets in launchSpec. Each element in the array should be a subnet ID.`,
				},
				resource.Attribute{
					Name:        "instance_metadata_options",
					Description: `(Optional) Ocean instance metadata options object for IMDSv2.`,
				},
				resource.Attribute{
					Name:        "http_tokens",
					Description: `(Required) Determines if a signed token is required or not. Valid values: ` + "`" + `optional` + "`" + ` or ` + "`" + `required` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "http_put_response_hop_limit",
					Description: `(Optional) An integer from 1 through 64. The desired HTTP PUT response hop limit for instance metadata requests. The larger the number, the further the instance metadata requests can travel.`,
				},
				resource.Attribute{
					Name:        "attributes",
					Description: `(Optional) Optionally adds labels to instances launched in an Ocean cluster.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) The label key.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) The label value.`,
				},
				resource.Attribute{
					Name:        "autoscale_headrooms",
					Description: `(Optional) Set custom headroom per launch spec. provide list of headrooms object.`,
				},
				resource.Attribute{
					Name:        "num_of_units",
					Description: `(Required) The number of units to retain as headroom, where each unit has the defined headroom CPU and memory.`,
				},
				resource.Attribute{
					Name:        "cpu_per_unit",
					Description: `(Optional) Optionally configure the number of CPUs to allocate for each headroom unit. CPUs are denoted in CPU units, where 1024 units = 1 vCPU.`,
				},
				resource.Attribute{
					Name:        "memory_per_unit",
					Description: `(Optional) Optionally configure the amount of memory (MiB) to allocate for each headroom unit.`,
				},
				resource.Attribute{
					Name:        "scheduling_task",
					Description: `(Optional) Used to define scheduled tasks such as a manual headroom update.`,
				},
				resource.Attribute{
					Name:        "is_enabled",
					Description: `(Required) Describes whether the task is enabled. When True, the task runs. When False, it does not run.`,
				},
				resource.Attribute{
					Name:        "cron_expression",
					Description: `(Required) A valid cron expression. For example : "`,
				},
				resource.Attribute{
					Name:        "task_type",
					Description: `(Required) The activity that you are scheduling. Valid values: "manualHeadroomUpdate".`,
				},
				resource.Attribute{
					Name:        "task_headroom",
					Description: `(Optional) The config of this scheduled task. Depends on the value of taskType.`,
				},
				resource.Attribute{
					Name:        "num_of_units",
					Description: `(Required) The number of units to retain as headroom, where each unit has the defined headroom CPU, memory and GPU.`,
				},
				resource.Attribute{
					Name:        "cpu_per_unit",
					Description: `(Optional) Optionally configure the number of CPUs to allocate for each headroom unit. CPUs are denoted in millicores, where 1000 millicores = 1 vCPU.`,
				},
				resource.Attribute{
					Name:        "memory_per_unit",
					Description: `(Optional) Optionally configure the amount of memory (MiB) to allocate for each headroom unit.`,
				},
				resource.Attribute{
					Name:        "spot_percentage",
					Description: `(Optional- if not using ` + "`" + `spot_percentege` + "`" + ` under ` + "`" + `ocean strategy` + "`" + `) When set, Ocean will proactively try to maintain as close as possible to the percentage of Spot instances out of all the Virtual Node Group instances. <a id="block-devices"></a> ## Block Devices`,
				},
				resource.Attribute{
					Name:        "device_name",
					Description: `(Optional) String. Set device name. (Example: "/dev/xvda1").`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The Spotinst LaunchSpec ID.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The Spotinst LaunchSpec ID.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "spotinst_ocean_gke_import",
			Category:         "Ocean",
			ShortDescription: `Provides a Spotinst Ocean resource using gke.`,
			Description:      ``,
			Keywords: []string{
				"ocean",
				"gke",
				"import",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cluster_name",
					Description: `(Required) The GKE cluster name.`,
				},
				resource.Attribute{
					Name:        "controller_cluster_id",
					Description: `(Required) A unique identifier used for connecting the Ocean SaaS platform and the Kubernetes cluster. Typically, the cluster name is used as its identifier.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `(Required) The zone the master cluster is located in.`,
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
					Name:        "whitelist",
					Description: `(Optional) Instance types allowed in the Ocean cluster. Cannot be configured if blacklist list is configured.`,
				},
				resource.Attribute{
					Name:        "blacklist",
					Description: `(Optional) Instance types to avoid launching in the Ocean cluster. Cannot be configured if whitelist list is configured.`,
				},
				resource.Attribute{
					Name:        "draining_timeout",
					Description: `(Optional) The draining timeout (in seconds) before terminating the instance.`,
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
					Description: `(Optional) Use when ` + "`" + `location_type` + "`" + ` is ` + "`" + `regional` + "`" + `. Set the traffic for the backend service to either between the instances in the vpc or to traffic from the internet. Valid values: ` + "`" + `INTERNAL` + "`" + `, ` + "`" + `EXTERNAL` + "`" + `.`,
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
					Description: `(Required) A list of ports.`,
				},
				resource.Attribute{
					Name:        "root_volume_type",
					Description: `(Optional) The root volume disk type.`,
				},
				resource.Attribute{
					Name:        "shielded_instance_config",
					Description: `(Optional) The Ocean shielded instance configuration object.`,
				},
				resource.Attribute{
					Name:        "enable_integrity_monitoring",
					Description: `(Optional) Boolean. Enable the integrity monitoring parameter on the GCP instances.`,
				},
				resource.Attribute{
					Name:        "enable_secure_boot",
					Description: `(Optional) Boolean. Enable the secure boot parameter on the GCP instances.`,
				},
				resource.Attribute{
					Name:        "use_as_template_only",
					Description: `(Optional, Default: false) launch specification defined on the Ocean object will function only as a template for virtual node groups. <a id="scheduled-task"></a> ## Scheduled task`,
				},
				resource.Attribute{
					Name:        "scheduled_task",
					Description: `(Optional) Set scheduling object.`,
				},
				resource.Attribute{
					Name:        "shutdown_hours",
					Description: `(Optional) Set shutdown hours for cluster object.`,
				},
				resource.Attribute{
					Name:        "is_enabled",
					Description: `(Optional) Flag to enable / disable the shutdown hours. Example: True`,
				},
				resource.Attribute{
					Name:        "time_windows",
					Description: `(Required) Set time windows for shutdown hours. specify a list of 'timeWindows' with at least one time window Each string is in the format of - ddd:hh:mm-ddd:hh:mm ddd = day of week = Sun | Mon | Tue | Wed | Thu | Fri | Sat hh = hour 24 = 0 -23 mm = minute = 0 - 59. Time windows should not overlap. required on cluster.scheduling.isEnabled = True. API Times are in UTC Example: Fri:15:30-Wed:14:30`,
				},
				resource.Attribute{
					Name:        "tasks",
					Description: `(Optional) The scheduling tasks for the cluster.`,
				},
				resource.Attribute{
					Name:        "is_enabled",
					Description: `(Required) Describes whether the task is enabled. When true the task should run when false it should not run. Required for cluster.scheduling.tasks object.`,
				},
				resource.Attribute{
					Name:        "cron_expression",
					Description: `(Required) A valid cron expression. For example : "`,
				},
				resource.Attribute{
					Name:        "task_type",
					Description: `(Required) Valid values: "clusterRoll". Required for cluster.scheduling.tasks object.`,
				},
				resource.Attribute{
					Name:        "batch_size_percentage",
					Description: `(Optional) Value in % to set size of batch in roll. Valid values are 0-100 Example: 20.`,
				},
				resource.Attribute{
					Name:        "task_parameters",
					Description: `(Optional) The scheduling parameters for the cluster.`,
				},
				resource.Attribute{
					Name:        "cluster_roll",
					Description: `(Optional) The cluster roll parameters for the cluster.`,
				},
				resource.Attribute{
					Name:        "batch_min_healthy_percentage",
					Description: `(Optional, Default: 50) Indicates the threshold of minimum healthy instances in single batch. If the amount of healthy instances in single batch is under the threshold, the cluster roll will fail. If exists, the parameter value will be in range of 1-100. In case of null as value, the default value in the backend will be 50%. Value of param should represent the number in percentage (%) of the batch.`,
				},
				resource.Attribute{
					Name:        "batch_size_percentage",
					Description: `(Optional) Value as a percent to set the size of a batch in a roll. Valid values are 0-100.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `(Optional) Add a comment description for the roll. The comment is limited to 256 chars.`,
				},
				resource.Attribute{
					Name:        "respect_pdb",
					Description: `(Optional, Default: 'false' ) During the roll, if the parameter is set to true we honor PDB during the instance replacement. ` + "`" + `` + "`" + `` + "`" + `hcl scheduled_task { shutdown_hours { is_enabled = false time_windows = ["Fri:15:30-Sat:18:30"] } tasks { is_enabled = false cron_expression = "0 1`,
				},
				resource.Attribute{
					Name:        "autoscaler",
					Description: `(Optional) The Ocean Kubernetes Autoscaler object.`,
				},
				resource.Attribute{
					Name:        "is_enabled",
					Description: `(Optional, Default: ` + "`" + `true` + "`" + `) Enable the Ocean Kubernetes Autoscaler.`,
				},
				resource.Attribute{
					Name:        "is_auto_config",
					Description: `(Optional, Default: ` + "`" + `true` + "`" + `) Automatically configure and optimize headroom resources.`,
				},
				resource.Attribute{
					Name:        "auto_headroom_percentage",
					Description: `Optionally set the auto headroom percentage, set a number between 0-200 to control the headroom % from the cluster. Relevant when isAutoConfig=true.`,
				},
				resource.Attribute{
					Name:        "cooldown",
					Description: `(Optional, Default: ` + "`" + `null` + "`" + `) Cooldown period between scaling actions.`,
				},
				resource.Attribute{
					Name:        "enable_automatic_and_manual_headroom",
					Description: `(Optional, Default: ` + "`" + `false` + "`" + `) enables automatic and manual headroom to work in parallel. When set to false, automatic headroom overrides all other headroom definitions manually configured, whether they are at cluster or VNG level.`,
				},
				resource.Attribute{
					Name:        "headroom",
					Description: `(Optional) Spare resource capacity management enabling fast assignment of Pods without waiting for new resources to launch.`,
				},
				resource.Attribute{
					Name:        "cpu_per_unit",
					Description: `(Optional) Optionally configure the number of CPUs to allocate the headroom. CPUs are denoted in millicores, where 1000 millicores = 1 vCPU.`,
				},
				resource.Attribute{
					Name:        "memory_per_unit",
					Description: `(Optional) Optionally configure the amount of memory (MiB) to allocate the headroom.`,
				},
				resource.Attribute{
					Name:        "gpu_per_unit",
					Description: `(Optional) How much GPU allocate for headroom unit.`,
				},
				resource.Attribute{
					Name:        "num_of_units",
					Description: `(Optional) The number of units to retain as headroom, where each unit has the defined headroom CPU and memory.`,
				},
				resource.Attribute{
					Name:        "down",
					Description: `(Optional) Auto Scaling scale down operations.`,
				},
				resource.Attribute{
					Name:        "evaluation_periods",
					Description: `(Optional, Default: ` + "`" + `null` + "`" + `) The number of evaluation periods that should accumulate before a scale down action takes place.`,
				},
				resource.Attribute{
					Name:        "max_scale_down_percentage",
					Description: `(Optional) Would represent the maximum % to scale-down. Number between 1-100.`,
				},
				resource.Attribute{
					Name:        "resource_limits",
					Description: `(Optional) Optionally set upper and lower bounds on the resource usage of the cluster.`,
				},
				resource.Attribute{
					Name:        "max_vcpu",
					Description: `(Optional) The maximum cpu in vCpu units that can be allocated to the cluster.`,
				},
				resource.Attribute{
					Name:        "max_memory_gib",
					Description: `(Optional) The maximum memory in GiB units that can be allocated to the cluster. ` + "`" + `` + "`" + `` + "`" + `hcl autoscaler { is_enabled = true is_auto_config = false cooldown = 30 auto_headroom_percentage = 10 enable_automatic_and_manual_headroom = false headroom { cpu_per_unit = 0 gpu_per_unit = 0 memory_per_unit = 0 num_of_units = 0 } down { evaluation_periods = 3 max_scale_down_percentage = 30 } resource_limits { max_vcpu = 1500 max_memory_gib = 750 } } ` + "`" + `` + "`" + `` + "`" + ` <a id="strategy"></a> ## Strategy`,
				},
				resource.Attribute{
					Name:        "strategy",
					Description: `(Optional) Strategy object.`,
				},
				resource.Attribute{
					Name:        "draining_timeout",
					Description: `(Optional) The draining timeout (in seconds) before terminating the instance. If no draining timeout is defined, the default draining timeout will be used.`,
				},
				resource.Attribute{
					Name:        "provisioning_model",
					Description: `(Optional) Define the provisioning model of the launched instances. Valid values: ` + "`" + `SPOT` + "`" + `, ` + "`" + `PREEMPTIBLE` + "`" + `.`,
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
					Name:        "conditioned_roll",
					Description: `(Optional, Default: false) Spot will perform a cluster Roll in accordance with a relevant modification of the cluster’s settings. When set to true , only specific changes in the cluster’s configuration will trigger a cluster roll (such as AMI, Key Pair, user data, instance types, load balancers, etc).`,
				},
				resource.Attribute{
					Name:        "roll_config",
					Description: `(Required) Holds the roll configuration.`,
				},
				resource.Attribute{
					Name:        "batch_size_percentage",
					Description: `(Required) Sets the percentage of the instances to deploy in each batch.`,
				},
				resource.Attribute{
					Name:        "launch_spec_ids",
					Description: `(Optional) List of Virtual Node Group identifiers to be rolled.`,
				},
				resource.Attribute{
					Name:        "batch_min_healthy_percentage",
					Description: `(Optional) Default: 50. Indicates the threshold of minimum healthy instances in single batch. If the amount of healthy instances in single batch is under the threshold, the cluster roll will fail. If exists, the parameter value will be in range of 1-100. In case of null as value, the default value in the backend will be 50%. Value of param should represent the number in percentage (%) of the batch.`,
				},
				resource.Attribute{
					Name:        "respect_pdb",
					Description: `(Optional) Default: False. During the roll, if the parameter is set to True we honor PDB during the instance replacement. ` + "`" + `` + "`" + `` + "`" + `hcl update_policy { should_roll = false conditioned_roll = true roll_config { batch_size_percentage = 33 launch_spec_ids = ["ols-1a2b3c4d"] batch_min_healthy_percentage = 20 respect_pdb = true } } ` + "`" + `` + "`" + `` + "`" + ` ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The Spotinst Ocean ID.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The Spotinst Ocean ID.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "spotinst_ocean_gke_launch_spec",
			Category:         "Ocean",
			ShortDescription: `Provides a Spotinst Ocean Launch Spec resource using GKE.`,
			Description:      ``,
			Keywords: []string{
				"ocean",
				"gke",
				"launch",
				"spec",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ocean_id",
					Description: `(Required) The Ocean cluster ID.`,
				},
				resource.Attribute{
					Name:        "node_pool_name",
					Description: `(Optional) The node pool you wish to use in your Launch Spec.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The launch specification name.`,
				},
				resource.Attribute{
					Name:        "source_image",
					Description: `(Required) Image URL.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `(Required only if ` + "`" + `node_pool_name` + "`" + ` is not set) Cluster's metadata.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) The metadata key.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) The metadata value.`,
				},
				resource.Attribute{
					Name:        "taints",
					Description: `(Optional) Optionally adds labels to instances launched in an Ocean cluster.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) The taint key.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) The taint value.`,
				},
				resource.Attribute{
					Name:        "effect",
					Description: `(Required) The effect of the taint. Valid values: ` + "`" + `"NoSchedule"` + "`" + `, ` + "`" + `"PreferNoSchedule"` + "`" + `, ` + "`" + `"NoExecute"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) Optionally adds labels to instances launched in an Ocean cluster.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) The label key.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) The label value.`,
				},
				resource.Attribute{
					Name:        "restrict_scale_down",
					Description: `(Optional) Boolean. When set to ` + "`" + `true` + "`" + `, VNG nodes will be treated as if all pods running have the restrict-scale-down label. Therefore, Ocean will not scale nodes down unless empty.`,
				},
				resource.Attribute{
					Name:        "root_volume_type",
					Description: `(Optional) Root volume disk type. Valid values: ` + "`" + `"pd-standard"` + "`" + `, ` + "`" + `"pd-ssd"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "root_volume_size",
					Description: `(Optional) Root volume size (in GB).`,
				},
				resource.Attribute{
					Name:        "instance_types",
					Description: `(Optional) List of supported machine types for the Launch Spec.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Every node launched from this configuration will be tagged with those tags. Note: during creation some tags are automatically imported to the state file, it is required to manually add it to the template configuration`,
				},
				resource.Attribute{
					Name:        "autoscale_headrooms_automatic",
					Description: `(Optional) Set automatic headroom per launch spec.`,
				},
				resource.Attribute{
					Name:        "auto_headroom_percentage",
					Description: `(Optional) Number between 0-200 to control the headroom % of the specific Virtual Node Group. Effective when cluster.autoScaler.headroom.automatic.` + "`" + `is_enabled` + "`" + ` = true is set on the Ocean cluster.`,
				},
				resource.Attribute{
					Name:        "autoscale_headrooms",
					Description: `(Optional) Set custom headroom per launch spec. provide list of headrooms object.`,
				},
				resource.Attribute{
					Name:        "num_of_units",
					Description: `(Required) The number of units to retain as headroom, where each unit has the defined headroom CPU, memory and GPU.`,
				},
				resource.Attribute{
					Name:        "cpu_per_unit",
					Description: `(Optional) Optionally configure the number of CPUs to allocate for each headroom unit. CPUs are denoted in millicores, where 1000 millicores = 1 vCPU.`,
				},
				resource.Attribute{
					Name:        "gpu_per_unit",
					Description: `(Optional) Optionally configure the number of GPUS to allocate for each headroom unit.`,
				},
				resource.Attribute{
					Name:        "memory_per_unit",
					Description: `(Optional) Optionally configure the amount of memory (MiB) to allocate for each headroom unit.`,
				},
				resource.Attribute{
					Name:        "strategy",
					Description: `(Optional) The Ocean Launch Spec Strategy object.`,
				},
				resource.Attribute{
					Name:        "preemptible_percentage",
					Description: `(Optional) Defines the desired preemptible percentage for this launch specification.`,
				},
				resource.Attribute{
					Name:        "shielded_instance_config",
					Description: `(Optional) The Ocean shielded instance configuration object.`,
				},
				resource.Attribute{
					Name:        "enable_integrity_monitoring",
					Description: `(Optional) Boolean. Enable the integrity monitoring parameter on the GCP instances.`,
				},
				resource.Attribute{
					Name:        "enable_secure_boot",
					Description: `(Optional) Boolean. Enable the secure boot parameter on the GCP instances.`,
				},
				resource.Attribute{
					Name:        "storage",
					Description: `(Optional) The Ocean virtual node group storage object.`,
				},
				resource.Attribute{
					Name:        "local_ssd_count",
					Description: `(Optional) Defines the number of local SSDs to be attached per node for this VNG.`,
				},
				resource.Attribute{
					Name:        "resource_limits",
					Description: `(Optional) The Ocean virtual node group resource limits object.`,
				},
				resource.Attribute{
					Name:        "max_instance_count",
					Description: `(Optional) Option to set a maximum number of instances per virtual node group. Can be null. If set, the value must be greater than or equal to 0.`,
				},
				resource.Attribute{
					Name:        "min_instance_count",
					Description: `(Optional) Option to set a minimum number of instances per virtual node group. Can be null. If set, the value must be greater than or equal to 0.`,
				},
				resource.Attribute{
					Name:        "service_account",
					Description: `(Optional) The account used by applications running on the VM to call GCP APIs.`,
				},
				resource.Attribute{
					Name:        "scheduling_task",
					Description: `(Optional) Used to define scheduled tasks such as a manual headroom update.`,
				},
				resource.Attribute{
					Name:        "is_enabled",
					Description: `(Required) Describes whether the task is enabled. When True, the task runs. When False, it does not run.`,
				},
				resource.Attribute{
					Name:        "cron_expression",
					Description: `(Required) A valid cron expression. For example : "`,
				},
				resource.Attribute{
					Name:        "task_type",
					Description: `(Required) The activity that you are scheduling. Valid values: "manualHeadroomUpdate".`,
				},
				resource.Attribute{
					Name:        "task_headroom",
					Description: `(Optional) The config of this scheduled task. Depends on the value of taskType.`,
				},
				resource.Attribute{
					Name:        "num_of_units",
					Description: `(Required) The number of units to retain as headroom, where each unit has the defined headroom CPU, memory and GPU.`,
				},
				resource.Attribute{
					Name:        "cpu_per_unit",
					Description: `(Optional) Optionally configure the number of CPUs to allocate for each headroom unit. CPUs are denoted in millicores, where 1000 millicores = 1 vCPU.`,
				},
				resource.Attribute{
					Name:        "gpu_per_unit",
					Description: `(Optional) Optionally configure the number of GPUS to allocate for each headroom unit.`,
				},
				resource.Attribute{
					Name:        "memory_per_unit",
					Description: `(Optional) Optionally configure the amount of memory (MiB) to allocate for each headroom unit.`,
				},
				resource.Attribute{
					Name:        "network_interfaces",
					Description: `(Optional) Settings for network interfaces.`,
				},
				resource.Attribute{
					Name:        "network",
					Description: `(Required) The name of the network.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Optional) Use a network resource from a different project. Set the project identifier to use its network resource. This parameter is relevant only if the network resource is in a different project.`,
				},
				resource.Attribute{
					Name:        "access_configs",
					Description: `(Optional) The network protocol of the VNG.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the access configuration.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) The type of the access configuration.`,
				},
				resource.Attribute{
					Name:        "alias_ip_ranges",
					Description: `(Optional) use the imported node pool’s associated aliasIpRange to assign secondary IP addresses to the nodes. Cannot be changed after VNG creation.`,
				},
				resource.Attribute{
					Name:        "ip_cidr_range",
					Description: `(Required) specify the IP address range in CIDR notation that can be used for the alias IP addresses associated with the imported node pool.`,
				},
				resource.Attribute{
					Name:        "subnetwork_range_name",
					Description: `(Required) specify the IP address range for the subnet secondary IP range. <a id="update-policy"></a> ## Update Policy`,
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
					Description: `(Required) Holds the roll configuration.`,
				},
				resource.Attribute{
					Name:        "batch_size_percentage",
					Description: `(Required) Sets the percentage of the instances to deploy in each batch. ` + "`" + `` + "`" + `` + "`" + `hcl update_policy { should_roll = false roll_config { batch_size_percentage = 33 } } ` + "`" + `` + "`" + `` + "`" + ` ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The Spotinst LaunchSpec ID.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The Spotinst LaunchSpec ID.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "spotinst_ocean_gke_launch_spec_import",
			Category:         "Ocean",
			ShortDescription: `Provides a Spotinst Ocean Launch Spec Import resource using GKE.`,
			Description:      ``,
			Keywords: []string{
				"ocean",
				"gke",
				"launch",
				"spec",
				"import",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "node_pool_name",
					Description: `(Required) The node pool you wish to use in your launchSpec. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The Spotinst LaunchSpec ID.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The Spotinst LaunchSpec ID.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "spotinst_ocean_spark",
			Category:         "Ocean",
			ShortDescription: `Provides a Spotinst Ocean Spark resource.`,
			Description:      ``,
			Keywords: []string{
				"ocean",
				"spark",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "spotinst_ocean_spark_virtual_node_group",
			Category:         "Ocean",
			ShortDescription: `Provides a Spotinst Ocean Spark Virtual Node Group resource`,
			Description:      ``,
			Keywords: []string{
				"ocean",
				"spark",
				"virtual",
				"node",
				"group",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "spotinst_stateful_node_azure",
			Category:         "Elastigroup",
			ShortDescription: `Provides a Spotinst Stateful Node resource using Azure.`,
			Description:      ``,
			Keywords: []string{
				"elastigroup",
				"stateful",
				"node",
				"azure",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "spotinst_subscription",
			Category:         "Subscription",
			ShortDescription: `Provides a Spotinst subscription resource.`,
			Description:      ``,
			Keywords: []string{
				"subscription",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "resource_id",
					Description: `(Required) Spotinst Resource id (Elastigroup or Ocean ID).`,
				},
				resource.Attribute{
					Name:        "event_type",
					Description: `(Required) The event to send the notification when triggered. Valid values: ` + "`" + `"AWS_EC2_INSTANCE_TERMINATE"` + "`" + `, ` + "`" + `"AWS_EC2_INSTANCE_TERMINATED"` + "`" + `, ` + "`" + `"AWS_EC2_INSTANCE_LAUNCH"` + "`" + `, ` + "`" + `"AWS_EC2_INSTANCE_READY_SIGNAL_TIMEOUT"` + "`" + `, ` + "`" + `"AWS_EC2_CANT_SPIN_OD"` + "`" + `, ` + "`" + `"AWS_EC2_INSTANCE_UNHEALTHY_IN_ELB"` + "`" + `, ` + "`" + `"GROUP_ROLL_FAILED"` + "`" + `, ` + "`" + `"GROUP_ROLL_FINISHED"` + "`" + `, ` + "`" + `"CANT_SCALE_UP_GROUP_MAX_CAPACITY"` + "`" + `, ` + "`" + `"GROUP_UPDATED"` + "`" + `, ` + "`" + `"AWS_EMR_PROVISION_TIMEOUT"` + "`" + `, ` + "`" + `"GROUP_BEANSTALK_INIT_READY"` + "`" + `, ` + "`" + `"AZURE_VM_TERMINATED"` + "`" + `, ` + "`" + `"AZURE_VM_TERMINATE"` + "`" + `, ` + "`" + `"AWS_EC2_MANAGED_INSTANCE_PAUSING"` + "`" + `, ` + "`" + `"AWS_EC2_MANAGED_INSTANCE_RESUMING"` + "`" + `, ` + "`" + `"AWS_EC2_MANAGED_INSTANCE_RECYCLING"` + "`" + `,` + "`" + `"AWS_EC2_MANAGED_INSTANCE_DELETING"` + "`" + `. Ocean Events:` + "`" + `"CLUSTER_ROLL_FINISHED"` + "`" + `,` + "`" + `"GROUP_ROLL_FAILED"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Required) The protocol to send the notification. Valid values: ` + "`" + `"email"` + "`" + `, ` + "`" + `"email-json"` + "`" + `, ` + "`" + `"aws-sns"` + "`" + `, ` + "`" + `"web"` + "`" + `. The following values are deprecated: ` + "`" + `"http"` + "`" + ` , ` + "`" + `"https"` + "`" + ` You can use the generic ` + "`" + `"web"` + "`" + ` protocol instead. ` + "`" + `"aws-sns"` + "`" + ` is only supported with AWS provider`,
				},
				resource.Attribute{
					Name:        "endpoint",
					Description: `(Required) The endpoint the notification will be sent to. url in case of ` + "`" + `"http"` + "`" + `/` + "`" + `"https"` + "`" + `/` + "`" + `"web"` + "`" + `, email address in case of ` + "`" + `"email"` + "`" + `/` + "`" + `"email-json"` + "`" + ` and sns-topic-arn in case of ` + "`" + `"aws-sns"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "format",
					Description: `(Optional) The format of the notification content (JSON Format - Key+Value). Valid Values : ` + "`" + `"instance-id"` + "`" + `, ` + "`" + `"event"` + "`" + `, ` + "`" + `"resource-id"` + "`" + `, ` + "`" + `"resource-name"` + "`" + `, ` + "`" + `"subnet-id"` + "`" + `, ` + "`" + `"availability-zone"` + "`" + `, ` + "`" + `"reason"` + "`" + `, ` + "`" + `"private-ip"` + "`" + `, ` + "`" + `"launchspec-id"` + "`" + ` Example: {"event": ` + "`" + `"event"` + "`" + `, ` + "`" + `"resourceId"` + "`" + `: ` + "`" + `"resource-id"` + "`" + `, ` + "`" + `"resourceName"` + "`" + `: ` + "`" + `"resource-name"` + "`" + `", ` + "`" + `"myCustomKey"` + "`" + `: ` + "`" + `"My content is set here"` + "`" + ` } Default: {` + "`" + `"event"` + "`" + `: ` + "`" + `"<event>"` + "`" + `, ` + "`" + `"instanceId"` + "`" + `: ` + "`" + `"<instance-id>"` + "`" + `, ` + "`" + `"resourceId"` + "`" + `: ` + "`" + `"<resource-id>"` + "`" + `, ` + "`" + `"resourceName"` + "`" + `: ` + "`" + `"<resource-name>"` + "`" + ` }. ## Attributes Reference The following attributes are exported:`,
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

		"spotinst_data_integration":                       0,
		"spotinst_elastigroup_aws":                        1,
		"spotinst_elastigroup_aws_beanstalk":              2,
		"spotinst_elastigroup_aws_suspension":             3,
		"spotinst_elastigroup_azure":                      4,
		"spotinst_elastigroup_azure_v3":                   5,
		"spotinst_elastigroup_gcp":                        6,
		"spotinst_elastigroup_gke":                        7,
		"spotinst_health_check":                           8,
		"spotinst_managed_instance_aws":                   9,
		"spotinst_mrscaler_aws":                           10,
		"spotinst_ocean_aks":                              11,
		"spotinst_ocean_aks_virtual_node_group":           12,
		"spotinst_ocean_aws":                              13,
		"spotinst_ocean_aws_extended_resource_definition": 14,
		"spotinst_ocean_aws_launch_spec":                  15,
		"spotinst_ocean_ecs":                              16,
		"spotinst_ocean_ecs_launch_spec":                  17,
		"spotinst_ocean_gke_import":                       18,
		"spotinst_ocean_gke_launch_spec":                  19,
		"spotinst_ocean_gke_launch_spec_import":           20,
		"spotinst_ocean_spark":                            21,
		"spotinst_ocean_spark_virtual_node_group":         22,
		"spotinst_stateful_node_azure":                    23,
		"spotinst_subscription":                           24,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
