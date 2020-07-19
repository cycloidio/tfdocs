package ksyun

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "ksyun_certificate",
			Category:         "KKCM Resources",
			ShortDescription: `Provides an ksyun_certificate resource.`,
			Description:      ``,
			Keywords: []string{
				"kkcm",
				"certificate",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "public_key",
					Description: `(Optional) The Public key of certificate .`,
				},
				resource.Attribute{
					Name:        "private_key",
					Description: `(Optional) The private key of certificate.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ksyun_eip",
			Category:         "KEIP Resources",
			ShortDescription: `Provides an Elastic IP resource.`,
			Description:      ``,
			Keywords: []string{
				"keip",
				"eip",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "line_id",
					Description: `(Required) The id of the line.`,
				},
				resource.Attribute{
					Name:        "band_width",
					Description: `(Required) The band width of the public address.`,
				},
				resource.Attribute{
					Name:        "charge_type",
					Description: `(Required) The charge type of the Elastic IP address.Valid Values:'PrePaidByMonth', 'PostPaidByPeak', 'PostPaidByDay', 'PostPaidByTransfer', 'PostPaidByHour', 'HourlyInstantSettlement'.`,
				},
				resource.Attribute{
					Name:        "purchase_time",
					Description: `(Required) Purchase time.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Optional) The id of the project. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `The Elastic IP address.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Elastic IP .`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "public_ip",
					Description: `The Elastic IP address.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Elastic IP .`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ksyun_eip_associate",
			Category:         "KEIP Resources",
			ShortDescription: `Provides an EIP Association resource for associating Elastic IP to UHost Instance, Load Balancer, etc..`,
			Description:      ``,
			Keywords: []string{
				"keip",
				"eip",
				"associate",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "allocation_id",
					Description: `(Required) The ID of EIP.`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `(Required) The type of the instance.Valid Values:'Ipfwd', 'Slb'.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Required) The id of the instance.`,
				},
				resource.Attribute{
					Name:        "network_interface_id",
					Description: `(Optional) The id of the network interface.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ksyun_healthcheck",
			Category:         "KSLB Resources",
			ShortDescription: `Provides an Health Check resource.`,
			Description:      ``,
			Keywords: []string{
				"kslb",
				"healthcheck",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "listener_id",
					Description: `(Required) The id of the listener.`,
				},
				resource.Attribute{
					Name:        "health_check_state",
					Description: `(Required) Status maintained by health examination.Valid Values:'start', 'stop'.`,
				},
				resource.Attribute{
					Name:        "healthy_threshold",
					Description: `(Required) Health threshold.Valid Values:1-10.`,
				},
				resource.Attribute{
					Name:        "interval",
					Description: `(Required) Interval of health examination.Valid Values:1-3600.`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `(Required) Health check timeout.Valid Values:1-3600.`,
				},
				resource.Attribute{
					Name:        "unhealthy_threshold",
					Description: `(Required) Unhealthy threshold.Valid Values:1-10.`,
				},
				resource.Attribute{
					Name:        "url_path",
					Description: `(Optional) Link to HTTP type listener health check.`,
				},
				resource.Attribute{
					Name:        "host_name",
					Description: `(Optional) Domain name of HTTP type health check.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ksyun_instance",
			Category:         "KKEC Resources",
			ShortDescription: `Provides an Host Instance resource.`,
			Description:      ``,
			Keywords: []string{
				"kkec",
				"instance",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "image_id",
					Description: `(Required) The ID for the image to use for the instance.`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `(Required) The type of instance to start.`,
				},
				resource.Attribute{
					Name:        "system_disk",
					Description: `(Required) System disk parameters. - ` + "`" + `disk_type` + "`" + ` - System disk type. ` + "`" + `Local_SSD` + "`" + `, Local SSD disk. ` + "`" + `SSD3.0` + "`" + `, The SSD cloud disk. ` + "`" + `EHDD` + "`" + `, The EHDD cloud disk. - ` + "`" + `disk_size` + "`" + ` - The size of the data disk.`,
				},
				resource.Attribute{
					Name:        "data_disk_gb",
					Description: `(Optional) The local SSD disk.`,
				},
				resource.Attribute{
					Name:        "data_disk",
					Description: `(Optional) The list of data disks created with instance. - ` + "`" + `type` + "`" + ` - Data disk type. ` + "`" + `SSD3.0` + "`" + `, The SSD cloud disk. ` + "`" + `EHDD` + "`" + `, The EHDD cloud disk. - ` + "`" + `size` + "`" + ` - Data disk type size. - ` + "`" + `delete_with_instance` + "`" + ` - Delete this data disk when the instance is destroyed. It only works on SSD3.0, EHDD, disk.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Required) The ID of subnet. the instance will use the subnet in the current region.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `(Required) Security Group to associate with.`,
				},
				resource.Attribute{
					Name:        "instance_password",
					Description: `(Optional) Password to an instance is a string of 8 to 32 characters.`,
				},
				resource.Attribute{
					Name:        "instance_name",
					Description: `(Optional) The name of instance, which contains 2-64 characters and only support Chinese, English, numbers.`,
				},
				resource.Attribute{
					Name:        "keep_image_login",
					Description: `(Optional) Keep the initial settings of the custom image.`,
				},
				resource.Attribute{
					Name:        "charge_type",
					Description: `(Required) Valid values are Monthly, Daily, HourlyInstantSettlement.`,
				},
				resource.Attribute{
					Name:        "purchase_time",
					Description: `(Optional) The duration that you will buy the resource.`,
				},
				resource.Attribute{
					Name:        "private_ip_address",
					Description: `(Optional) Instance private IP address can be specified when you creating new instance.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Optional) The project instance belongs to.`,
				},
				resource.Attribute{
					Name:        "user_data",
					Description: `(Optional) The user data to be specified into this instance. Must be encrypted in base64 format and limited in 16 KB. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "creation_date",
					Description: `The time of creation for instance, formatted in ISO8601 time string.`,
				},
				resource.Attribute{
					Name:        "instance_state",
					Description: `Instance current status. Possible values are ` + "`" + `active` + "`" + `, ` + "`" + `building` + "`" + `, ` + "`" + `stopped` + "`" + `, ` + "`" + `deleting` + "`" + `. ## Import Instance can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "creation_date",
					Description: `The time of creation for instance, formatted in ISO8601 time string.`,
				},
				resource.Attribute{
					Name:        "instance_state",
					Description: `Instance current status. Possible values are ` + "`" + `active` + "`" + `, ` + "`" + `building` + "`" + `, ` + "`" + `stopped` + "`" + `, ` + "`" + `deleting` + "`" + `. ## Import Instance can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ksyun_krds",
			Category:         "KKRDS Resources",
			ShortDescription: `Provides an KRDS resource.`,
			Description:      ``,
			Keywords: []string{
				"kkrds",
				"krds",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ksyun_krds_read_replica",
			Category:         "KKRDS Resources",
			ShortDescription: `Provides an KRDS readonly instance resource.`,
			Description:      ``,
			Keywords: []string{
				"kkrds",
				"krds",
				"read",
				"replica",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ksyun_krds_security_group",
			Category:         "KKRDS Resources",
			ShortDescription: `Provides an KRDS security group resource.`,
			Description:      ``,
			Keywords: []string{
				"kkrds",
				"krds",
				"security",
				"group",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ksyun_lb",
			Category:         "KSLB Resources",
			ShortDescription: `Provides a Load Balancer resource.`,
			Description:      ``,
			Keywords: []string{
				"kslb",
				"lb",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "load_balancer_name",
					Description: `(Optional) The name of the load balancer.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required) The ID of the VPC linked to the Load Balancers.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) The type of load balancer.Valid Values:'public', 'internal'.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Optional) The id of the subnet.only Internal type is Required. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The time of creation for load balancer, formatted in RFC3339 time string.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `The IP address of intranet IP. It is ` + "`" + `""` + "`" + ` if ` + "`" + `internal` + "`" + ` is ` + "`" + `false` + "`" + `. ## Import LB can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import ksyun_lb.example fdeba8ca-8aa6-4cd0-8ffa-52ca9e9fef42 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "create_time",
					Description: `The time of creation for load balancer, formatted in RFC3339 time string.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `The IP address of intranet IP. It is ` + "`" + `""` + "`" + ` if ` + "`" + `internal` + "`" + ` is ` + "`" + `false` + "`" + `. ## Import LB can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import ksyun_lb.example fdeba8ca-8aa6-4cd0-8ffa-52ca9e9fef42 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ksyun_lb_acl",
			Category:         "KSLB Resources",
			ShortDescription: `Provides a Load Balancer acl resource to add content forwarding policies for Load Balancer backend resource.`,
			Description:      ``,
			Keywords: []string{
				"kslb",
				"lb",
				"acl",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "load_balancer_acl_name",
					Description: `(Required) The name of a load balancer acl.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ksyun_lb_acl_entry",
			Category:         "KSLB Resources",
			ShortDescription: `Provides a Load Balancer acl entry resource to add content forwarding policies for Load Balancer backend resource.`,
			Description:      ``,
			Keywords: []string{
				"kslb",
				"lb",
				"acl",
				"entry",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "load_balancer_acl_id",
					Description: `(Required) The ID of a load balancer acl.`,
				},
				resource.Attribute{
					Name:        "load_balancer_acl_id",
					Description: `(Required) The id of the load balancer acl.`,
				},
				resource.Attribute{
					Name:        "cidr_block",
					Description: `(Required) The information of load balancer Acl's cidr block.`,
				},
				resource.Attribute{
					Name:        "rule_number",
					Description: `(Required) The information of load balancer Acl's rule priority.Valid Values:1-32766.`,
				},
				resource.Attribute{
					Name:        "rule_action",
					Description: `(Required) The action of load balancer Acl rule.Valid Values:'allow', 'deny'.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Required) protocol.Valid Values:'ip'.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ksyun_lb_listener",
			Category:         "KSLB Resources",
			ShortDescription: `Provides a Load Balancer Listener resource.`,
			Description:      ``,
			Keywords: []string{
				"kslb",
				"lb",
				"listener",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "listener_name",
					Description: `(Optional) The name of listener.`,
				},
				resource.Attribute{
					Name:        "listener_port",
					Description: `(Required) The protocol port of listener.`,
				},
				resource.Attribute{
					Name:        "listener_protocol",
					Description: `(Required) The protocol of listener.Valid Values:'TCP', 'UDP', 'HTTP', 'HTTPS'.`,
				},
				resource.Attribute{
					Name:        "listener_state",
					Description: `(Required) The state of listener.Valid Values:'start', 'stop'.`,
				},
				resource.Attribute{
					Name:        "load_balancer_id",
					Description: `(Required) The ID of load balancer instance.`,
				},
				resource.Attribute{
					Name:        "method",
					Description: `(Required) Forwarding mode of listener.Valid Values:'RoundRobin', 'LeastConnections'.`,
				},
				resource.Attribute{
					Name:        "certificate_id",
					Description: `(Optional) The ID of certificate.`,
				},
				resource.Attribute{
					Name:        "session_state",
					Description: `(Required) The state of session.Valid Values:'start', 'stop'.`,
				},
				resource.Attribute{
					Name:        "session_persistence_period",
					Description: `(Optional) Session hold timeout.Valid Values:1-86400`,
				},
				resource.Attribute{
					Name:        "cookie_type",
					Description: `(Optional) The type of the cookie.Valid Values:'ImplantCookie', 'RewriteCookie'.`,
				},
				resource.Attribute{
					Name:        "cookie_name",
					Description: `(Optional) The name of the cookie. ## Import LB Listener can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import ksyun_lb_listener.example vserver-abcdefg ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ksyun_lb_listener_associate_acl",
			Category:         "KSLB Resources",
			ShortDescription: `Associate a Load Balancer Listener resource with acl.`,
			Description:      ``,
			Keywords: []string{
				"kslb",
				"lb",
				"listener",
				"associate",
				"acl",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "load_balancer_id",
					Description: `(Required) The ID of load balancer instance.`,
				},
				resource.Attribute{
					Name:        "listener_name",
					Description: `(Optional) The id of the listener.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ksyun_lb_listener_server",
			Category:         "KSLB Resources",
			ShortDescription: `Provides a Load Balancer Listener server resource.`,
			Description:      ``,
			Keywords: []string{
				"kslb",
				"lb",
				"listener",
				"server",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "listener_id",
					Description: `(Required) The id of the listener.`,
				},
				resource.Attribute{
					Name:        "real_server_type",
					Description: `(Required) The type of real server.Valid Values:'Host', 'DirectConnectGateway', 'VpnTunnel'.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Optional) The ID of instance.`,
				},
				resource.Attribute{
					Name:        "real_server_ip",
					Description: `(Required) The IP of real server.`,
				},
				resource.Attribute{
					Name:        "real_server_port",
					Description: `(Required) The port of real server.Valid Values:1-65535`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `(Optional) The weight of backend service.Valid Values:1-255 ## Import LB Listener can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import ksyun_lb_listener.example vserver-abcdefg ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ksyun_mongodb_instance",
			Category:         "KMONGODB Resources",
			ShortDescription: `Provides an replica set MongoDB resource.`,
			Description:      ``,
			Keywords: []string{
				"kmongodb",
				"mongodb",
				"instance",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of instance, which contains 6-64 characters and only support Chinese, English, numbers, '-', '_'.`,
				},
				resource.Attribute{
					Name:        "instance_account",
					Description: `(Optional) The administrator name of instance, if not defined ` + "`" + `instance_account` + "`" + `, the instance will use ` + "`" + `root` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "instance_password",
					Description: `(Required) The administrator password of instance.`,
				},
				resource.Attribute{
					Name:        "instance_class",
					Description: `(Required) The class of instance cpu and memory.`,
				},
				resource.Attribute{
					Name:        "storage",
					Description: `(Required) The size of instance disk, measured in GB (GigaByte).`,
				},
				resource.Attribute{
					Name:        "node_num",
					Description: `(Required) The num of instance node.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required) The id of VPC linked to the instance.`,
				},
				resource.Attribute{
					Name:        "vnet_id",
					Description: `(Required) The id of subnet linked to the instance.`,
				},
				resource.Attribute{
					Name:        "db_version",
					Description: `(Required) The version of instance engine, and support ` + "`" + `3.2` + "`" + ` and ` + "`" + `3.6` + "`" + ``,
				},
				resource.Attribute{
					Name:        "pay_type",
					Description: `(Optional) Instance charge type, if not defined ` + "`" + `pay_type` + "`" + `, the instance will use ` + "`" + `byMonth` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "duration",
					Description: `(Optional) The duration of instance use, if ` + "`" + `pay_type` + "`" + ` is ` + "`" + `byMonth` + "`" + `, the duration is required.`,
				},
				resource.Attribute{
					Name:        "iam_project_id",
					Description: `(Optional) The project id of instance belong, if not defined ` + "`" + `iam_project_id` + "`" + `, the instance will use ` + "`" + `0` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `(Required) Availability zone where instance is located.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ksyun_mongodb_security_rule",
			Category:         "KMONGODB Resources",
			ShortDescription: `Provides an MongoDB Security Rule resource.`,
			Description:      ``,
			Keywords: []string{
				"kmongodb",
				"mongodb",
				"security",
				"rule",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Required) The id of instance, .`,
				},
				resource.Attribute{
					Name:        "cidrs",
					Description: `(Required) The cidr block of source for the instance, multiple cidr separated by comma.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ksyun_mongodb_shard_instance",
			Category:         "KMONGODB Resources",
			ShortDescription: `Provides an shard MongoDB resource.`,
			Description:      ``,
			Keywords: []string{
				"kmongodb",
				"mongodb",
				"shard",
				"instance",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of instance, which contains 6-64 characters and only support Chinese, English, numbers, '-', '_'.`,
				},
				resource.Attribute{
					Name:        "instance_account",
					Description: `(Optional) The administrator name of instance, if not defined ` + "`" + `instance_account` + "`" + `, the instance will use ` + "`" + `root` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "instance_password",
					Description: `(Required) The administrator password of instance.`,
				},
				resource.Attribute{
					Name:        "mongos_class",
					Description: `(Required) The class of instance mongo node cpu and memory.`,
				},
				resource.Attribute{
					Name:        "mongos_num",
					Description: `(Required) The num of instance mongo node.`,
				},
				resource.Attribute{
					Name:        "shard_class",
					Description: `(Required) The class of instance shard node cpu and memory.`,
				},
				resource.Attribute{
					Name:        "shard_num",
					Description: `(Required) The num of instance shard node.`,
				},
				resource.Attribute{
					Name:        "storage",
					Description: `(Required) The size of instance disk, measured in GB (GigaByte).`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required) The id of VPC linked to the instance.`,
				},
				resource.Attribute{
					Name:        "vnet_id",
					Description: `(Required) The id of subnet linked to the instance.`,
				},
				resource.Attribute{
					Name:        "db_version",
					Description: `(Required) The version of instance engine, and support ` + "`" + `3.2` + "`" + ` and ` + "`" + `3.6` + "`" + ``,
				},
				resource.Attribute{
					Name:        "pay_type",
					Description: `(Optional) Instance charge type, if not defined ` + "`" + `pay_type` + "`" + `, the instance will use ` + "`" + `byMonth` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "duration",
					Description: `(Optional) The duration of instance use, if ` + "`" + `pay_type` + "`" + ` is ` + "`" + `byMonth` + "`" + `, the duration is required.`,
				},
				resource.Attribute{
					Name:        "iam_project_id",
					Description: `(Optional) The project id of instance belong, if not defined ` + "`" + `iam_project_id` + "`" + `, the instance will use ` + "`" + `0` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `(Required) Availability zone where instance is located.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ksyun_redis_instance",
			Category:         "KKCS Resources",
			ShortDescription: `Provides an Redis instance resource.`,
			Description:      ``,
			Keywords: []string{
				"kkcs",
				"redis",
				"instance",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "available_zone",
					Description: `(Optional) The Zone to launch the DB instance.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of DB instance.`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `(Optional) The KVStore instance system architecture required by the user. Valid values: 1(cluster),2(single).`,
				},
				resource.Attribute{
					Name:        "capacity",
					Description: `(Require) The instance capacity required by the user. Valid values :{1, 2, 4, 8, 16,20,24,28, 32, 64}.`,
				},
				resource.Attribute{
					Name:        "slaveNum",
					Description: `(Optional) The readonly node num required by the user. Valid values ：{0-7}`,
				},
				resource.Attribute{
					Name:        "net_type",
					Description: `(Require) The network type. Valid values ：2(vpc).`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Require) Used to retrieve instances belong to specified VPC .`,
				},
				resource.Attribute{
					Name:        "vnet_id",
					Description: `(Require) The ID of subnet. the instance will use the subnet in the current region.`,
				},
				resource.Attribute{
					Name:        "bill_type",
					Description: `(Optional)Valid values are 1 (Monthly), 5(Daily), 87(HourlyInstantSettlement).`,
				},
				resource.Attribute{
					Name:        "duration",
					Description: `(Optional)Only meaningful if bill_type is 1。 Valid values：{1~36}.`,
				},
				resource.Attribute{
					Name:        "duration_unit",
					Description: `(Optional)Only meaningful if bill_type is 1。 Valid values：month.`,
				},
				resource.Attribute{
					Name:        "pass_word",
					Description: `(Optional)The password of the instance.The password is a string of 8 to 30 characters and must contain uppercase letters, lowercase letters, and numbers.`,
				},
				resource.Attribute{
					Name:        "iam_project_id",
					Description: `(Optional) The project instance belongs to.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `Engine version. Supported values: 2.8, 4.0 and 5.0.`,
				},
				resource.Attribute{
					Name:        "parameters",
					Description: `Set of parameters needs to be set after instance was launched. Available parameters can refer to the docs https://docs.ksyun.com/documents/1018 .`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ksyun_redis_instance_node",
			Category:         "KKCS Resources",
			ShortDescription: `Provides an redis instance node resource.`,
			Description:      ``,
			Keywords: []string{
				"kkcs",
				"redis",
				"instance",
				"node",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cache_id",
					Description: `(Optional) The ID of the intance .`,
				},
				resource.Attribute{
					Name:        "available_zone",
					Description: `(Optional) The Zone to launch the DB instance.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ksyun_redis_sec_rule",
			Category:         "KKCS Resources",
			ShortDescription: `Provides an redis security rule resource.`,
			Description:      ``,
			Keywords: []string{
				"kkcs",
				"redis",
				"sec",
				"rule",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Required) The id of instance, .`,
				},
				resource.Attribute{
					Name:        "cidrs",
					Description: `(Required) The cidr block of source for the instance, multiple cidr separated by comma.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ksyun_security_group",
			Category:         "KVPC Resources",
			ShortDescription: `Provides a Security Group resource.`,
			Description:      ``,
			Keywords: []string{
				"kvpc",
				"security",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "security_group_name",
					Description: `(Optional) The name of the security group which contains 1-63 characters and only support Chinese, English, numbers, '-', '_' and '.'.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Optional) The Id of the vpc. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The time of creation of security group, formatted in RFC3339 time string. ## Import Security Group can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import ksyun_security_group.example firewall-abc123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "create_time",
					Description: `The time of creation of security group, formatted in RFC3339 time string. ## Import Security Group can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import ksyun_security_group.example firewall-abc123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ksyun_security_group_entry",
			Category:         "KVPC Resources",
			ShortDescription: `Provides a Security Group resource.`,
			Description:      ``,
			Keywords: []string{
				"kvpc",
				"security",
				"group",
				"entry",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description of the security group .`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `(Required) The ID of the security group.`,
				},
				resource.Attribute{
					Name:        "cidr_block",
					Description: `(Required) The cidr block of security group rules.`,
				},
				resource.Attribute{
					Name:        "direction",
					Description: `(Required) .Valid Values:'in', 'out'.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Required) protocol.Valid Values:'ip', 'tcp', 'udp', 'icmp'.`,
				},
				resource.Attribute{
					Name:        "icmp_type",
					Description: `(Optional) ICMP protocol.The required if protocol type is 'icmp'.`,
				},
				resource.Attribute{
					Name:        "icmp_code",
					Description: `(Optional) ICMP protocol.The required if protocol type is 'icmp'.`,
				},
				resource.Attribute{
					Name:        "port_range_from",
					Description: `(Optional) Port rule start port for TCP or UDP protocol.The required if protocol type is 'tcp' or 'udp'.`,
				},
				resource.Attribute{
					Name:        "port_range_to",
					Description: `(Optional) Port rule start port for TCP or UDP protocol.The required if protocol type is 'tcp' or 'udp'. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The time of creation of security group, formatted in RFC3339 time string. ## Import Security Group can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import ksyun_security_group_entry.example firewall-abc123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "create_time",
					Description: `The time of creation of security group, formatted in RFC3339 time string. ## Import Security Group can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import ksyun_security_group_entry.example firewall-abc123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ksyun_sqlserver",
			Category:         "KKRDS Resources",
			ShortDescription: `Provides an sqlserver instance resource.`,
			Description:      ``,
			Keywords: []string{
				"kkrds",
				"sqlserver",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ksyun_subnet",
			Category:         "KVPC Resources",
			ShortDescription: `Provides a Subnet resource under VPC resource.`,
			Description:      ``,
			Keywords: []string{
				"kvpc",
				"subnet",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "subnet_name",
					Description: `(Required) The name of the subnet.`,
				},
				resource.Attribute{
					Name:        "cidr_block",
					Description: `(Required) The CIDR block assigned to the subnet.`,
				},
				resource.Attribute{
					Name:        "subnet_type",
					Description: `(Required) The type of subnet.Valid Values:'Reserve', 'Normal', 'Physical'.`,
				},
				resource.Attribute{
					Name:        "dhcp_ip_from",
					Description: `(Required) DHCP start IP.`,
				},
				resource.Attribute{
					Name:        "dhcp_ip_to",
					Description: `(Required) DHCP end IP.`,
				},
				resource.Attribute{
					Name:        "gateway_ip",
					Description: `(Required) The IP of gateway.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required) The id of the vpc.`,
				},
				resource.Attribute{
					Name:        "dns1",
					Description: `(Optional) The dns of the subnet.`,
				},
				resource.Attribute{
					Name:        "dns2",
					Description: `(Optional) The dns of the subnet.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `(Optional) The name of the availability zone. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The time of creation of subnet, formatted in RFC3339 time string. ## Import Subnet can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import ksyun_subnet.example subnet-abc123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "create_time",
					Description: `The time of creation of subnet, formatted in RFC3339 time string. ## Import Subnet can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import ksyun_subnet.example subnet-abc123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ksyun_vpc",
			Category:         "KVPC Resources",
			ShortDescription: `Provides a VPC resource.`,
			Description:      ``,
			Keywords: []string{
				"kvpc",
				"vpc",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cidr_block",
					Description: `(Required) The CIDR blocks of VPC.`,
				},
				resource.Attribute{
					Name:        "vpc_name",
					Description: `(Optional) The name of the vpc. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The time of creation for VPC, formatted in RFC3339 time string.`,
				},
				resource.Attribute{
					Name:        "cidr_block",
					Description: `The CIDR block of the VPC. ## Import VPC can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import ksyun_vpc.example uvnet-abc123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "create_time",
					Description: `The time of creation for VPC, formatted in RFC3339 time string.`,
				},
				resource.Attribute{
					Name:        "cidr_block",
					Description: `The CIDR block of the VPC. ## Import VPC can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import ksyun_vpc.example uvnet-abc123456 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
	}

	resourcesMap = map[string]int{

		"ksyun_certificate":               0,
		"ksyun_eip":                       1,
		"ksyun_eip_associate":             2,
		"ksyun_healthcheck":               3,
		"ksyun_instance":                  4,
		"ksyun_krds":                      5,
		"ksyun_krds_read_replica":         6,
		"ksyun_krds_security_group":       7,
		"ksyun_lb":                        8,
		"ksyun_lb_acl":                    9,
		"ksyun_lb_acl_entry":              10,
		"ksyun_lb_listener":               11,
		"ksyun_lb_listener_associate_acl": 12,
		"ksyun_lb_listener_server":        13,
		"ksyun_mongodb_instance":          14,
		"ksyun_mongodb_security_rule":     15,
		"ksyun_mongodb_shard_instance":    16,
		"ksyun_redis_instance":            17,
		"ksyun_redis_instance_node":       18,
		"ksyun_redis_sec_rule":            19,
		"ksyun_security_group":            20,
		"ksyun_security_group_entry":      21,
		"ksyun_sqlserver":                 22,
		"ksyun_subnet":                    23,
		"ksyun_vpc":                       24,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
