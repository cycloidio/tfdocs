package aws

import (
	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*Resource{

		&resource.Resource{
			Name:             "",
			Type:             "rightscale_credential",
			Category:         "Resources",
			ShortDescription: `Create and maintain a RightScale credential resource.`,
			Description:      ``,
			Keywords: []string{
				"credential",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the credential.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) Value of the credential.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the credential. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the credential.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the credential.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `Value of the credential.`,
				},
				resource.Attribute{
					Name:        "links",
					Description: `Hrefs of related API resources.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Datestamp of credential creation.`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `Datestamp of when credential was updated last.`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `Name of the credential.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the credential.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `Value of the credential.`,
				},
				resource.Attribute{
					Name:        "links",
					Description: `Hrefs of related API resources.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Datestamp of credential creation.`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `Datestamp of when credential was updated last.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "rightscale_deployment",
			Category:         "Resources",
			ShortDescription: `Create and maintain a RightScale deployment.`,
			Description:      ``,
			Keywords: []string{
				"deployment",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Deployment name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Deployment description.`,
				},
				resource.Attribute{
					Name:        "resource_group_href",
					Description: `(Optional) Href of the Windows Azure Resource Group attached to the deployment.`,
				},
				resource.Attribute{
					Name:        "locked",
					Description: `(Optional) Set to true to lock the deployment.`,
				},
				resource.Attribute{
					Name:        "server_tag_scope",
					Description: `(Optional) Routing scope for tags for servers in the deployment. Options are 'account' or 'deployment,' defaults to 'deployment.' ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "href",
					Description: `Href of the deployment.`,
				},
				resource.Attribute{
					Name:        "links",
					Description: `Hrefs of related API resources.`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "href",
					Description: `Href of the deployment.`,
				},
				resource.Attribute{
					Name:        "links",
					Description: `Hrefs of related API resources.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "rightscale_instance",
			Category:         "Resources",
			ShortDescription: `Create and maintain a RightScale instance.`,
			Description:      ``,
			Keywords: []string{
				"instance",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the instance.`,
				},
				resource.Attribute{
					Name:        "cloud_href",
					Description: `(Required) The cloud_href the instance belongs to.`,
				},
				resource.Attribute{
					Name:        "image_href",
					Description: `(Required) The href of the instance image.`,
				},
				resource.Attribute{
					Name:        "instance_type_href",
					Description: `(Required) The href of the instance type.`,
				},
				resource.Attribute{
					Name:        "server_template_href",
					Description: `(Optional) The href of the instance server template resource.`,
				},
				resource.Attribute{
					Name:        "inputs",
					Description: `(Optional) Inputs associated with an instance when incarnated from a [server](https://github.com/terraform-providers/terraform-provider-rightscale/blob/master/website/docs/r/cm_server.markdown) or [server_array](https://github.com/terraform-providers/terraform-provider-rightscale/blob/master/website/docs/r/cm_server_array.markdown).`,
				},
				resource.Attribute{
					Name:        "associate_public_ip_address",
					Description: `(Optional) Indicates if the instance will get a Public IP address.`,
				},
				resource.Attribute{
					Name:        "datacenter_href",
					Description: `(Optional) The href of the datacenter that holds the instance (e.g. /api/clouds/6/datacenters/6IHONC8ANOUHI).`,
				},
				resource.Attribute{
					Name:        "deployment_href",
					Description: `(Optional) The href of the deployment that contains the instance (e.g. /api/deployments/594684003).`,
				},
				resource.Attribute{
					Name:        "ip_forwarding_enabled",
					Description: `(Optional) Allows this Instance to send and receive network traffic when the source and destination IP addresses do not match the IP address of this Instance.`,
				},
				resource.Attribute{
					Name:        "private_ip_address",
					Description: `(Optional) The private ip address of this instance.`,
				},
				resource.Attribute{
					Name:        "kernel_image_href",
					Description: `(Optional) The href of the instance kernel image.`,
				},
				resource.Attribute{
					Name:        "ramdisk_image_href",
					Description: `(Optional) The href of the instance ramdisk image.`,
				},
				resource.Attribute{
					Name:        "security_group_hrefs",
					Description: `(Optional) The href of the instance security groups.`,
				},
				resource.Attribute{
					Name:        "placement_group_href",
					Description: `(Optional) The href of the [placement_group](http://docs.rightscale.com/cm/dashboard/clouds/aws/ec2_placement_groups.html) that contains the instance (e.g. /api/placement_groups/512SV3FUJA7OO).`,
				},
				resource.Attribute{
					Name:        "ssh_key_href",
					Description: `(Optional) The href of the SSH key to use.`,
				},
				resource.Attribute{
					Name:        "subnet_hrefs",
					Description: `(Optional) The hrefs of the instance subnet.`,
				},
				resource.Attribute{
					Name:        "user_data",
					Description: `(Optional) User data that RightScale automatically passes to your instance at boot time.`,
				},
				resource.Attribute{
					Name:        "locked",
					Description: `(Optional) Whether instance is locked, a locked instance cannot be terminated or deleted.`,
				},
				resource.Attribute{
					Name:        "cloud_specific_attributes",
					Description: `(Optional) Attributes specific to the cloud the instance belongs to that have no specific rightscale abstraction. This block supports:`,
				},
				resource.Attribute{
					Name:        "admin_username",
					Description: `The user that will be granted administrative privileges. Supported by AzureRM cloud only.`,
				},
				resource.Attribute{
					Name:        "automatic_instance_store_mapping",
					Description: `A flag indicating whether instance store mapping should be enabled. Only available on clouds supporting automatic instance store mapping.`,
				},
				resource.Attribute{
					Name:        "availability_set",
					Description: `Availability set for raw instance. Supported by Azure v2 cloud only.`,
				},
				resource.Attribute{
					Name:        "create_boot_volume",
					Description: `If enabled, the instance will launch into volume storage. Otherwise, it will boot to local storage. Only available on clouds supporting this option.`,
				},
				resource.Attribute{
					Name:        "create_default_port_forwarding_rules",
					Description: `Automatically create default port forwarding rules (enabled by default). Supported by Azure cloud only.`,
				},
				resource.Attribute{
					Name:        "delete_boot_volume",
					Description: `If enabled, the associated volume will be deleted when the instance is terminated. Only available on clouds supporting this option.`,
				},
				resource.Attribute{
					Name:        "disk_gb",
					Description: `The size of root disk. Supported by UCA cloud only.`,
				},
				resource.Attribute{
					Name:        "ebs_optimized",
					Description: `Whether the instance is able to connect to IOPS-enabled volumes. AWS clouds only.`,
				},
				resource.Attribute{
					Name:        "iam_instance_profile",
					Description: `The name or ARN of the IAM Instance Profile (IIP) to associate with the instance. AWS clouds only.`,
				},
				resource.Attribute{
					Name:        "keep_alive_id",
					Description: `The id of keep alive. Supported by UCA cloud only.`,
				},
				resource.Attribute{
					Name:        "local_ssd_count",
					Description: `Additional local SSDs. Supported by GCE cloud only.`,
				},
				resource.Attribute{
					Name:        "local_ssd_interface",
					Description: `The type of SSD(s) to be created. Supported by GCE cloud only.`,
				},
				resource.Attribute{
					Name:        "max_spot_price",
					Description: `Specify the max spot price you will pay for. Required when 'pricing_type' is 'spot'. Only applies to clouds which support spot-pricing and when 'spot' is chosen as the 'pricing_type'. Should be a Float value >= 0.001, eg: 0.095, 0.123, 1.23, etc... AWS clouds only.`,
				},
				resource.Attribute{
					Name:        "memory_mb",
					Description: `The size of instance memory. Supported by UCA cloud only.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `Extra data used for configuration, in query string format. AWS clouds only.`,
				},
				resource.Attribute{
					Name:        "num_cores",
					Description: `The number of instance cores. Supported by UCA cloud only.`,
				},
				resource.Attribute{
					Name:        "placement_tenancy",
					Description: `The tenancy of the server you want to launch. A server with a tenancy of dedicated runs on single-tenant hardware and can only be launched into a VPC. AWS clouds only.`,
				},
				resource.Attribute{
					Name:        "preemptible",
					Description: `Launch a preemptible instance. A preemptible instance costs much less, but lasts only 24 hours. It can be terminated sooner due to system demands. Supported by GCE cloud only.`,
				},
				resource.Attribute{
					Name:        "pricing_type",
					Description: `Specify whether or not you want to utilize 'fixed' (on-demand) or 'spot' pricing. Defaults to 'fixed' and only applies to clouds which support spot instances. Can only be set on when creating a new Instance, Server, or ServerArray, or when updating a Server or ServerArray's next_instance. AWS clouds only.`,
				},
				resource.Attribute{
					Name:        "root_volume_performance",
					Description: `The number of IOPS (I/O Operations Per Second) this root volume should support. Only available on clouds supporting performance provisioning.`,
				},
				resource.Attribute{
					Name:        "root_volume_size",
					Description: `The size for root disk. Only available on clouds supporting dynamic resizing of root volume size.`,
				},
				resource.Attribute{
					Name:        "root_volume_type_uid",
					Description: `The type of root volume for instance. Only available on clouds supporting root volume type.`,
				},
				resource.Attribute{
					Name:        "service_account",
					Description: `Email of service account for instance. Scope will default to cloud-platform. Supported by GCE cloud only. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "links",
					Description: `Hrefs of related API resources`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Datestamp of instance creation.`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `Datestamp of when instance was updated last.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the instance (operational, terminating, pending, stranded, etc.)`,
				},
				resource.Attribute{
					Name:        "href",
					Description: `Href of the instance.`,
				},
				resource.Attribute{
					Name:        "resource_uid",
					Description: `Cloud resource_uid as reported by cm platform.`,
				},
				resource.Attribute{
					Name:        "public_ip_addresses",
					Description: `List of public IP addresses associated to the instance`,
				},
				resource.Attribute{
					Name:        "private_ip_addresses",
					Description: `List of private IP addresses associated to the instance`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "links",
					Description: `Hrefs of related API resources`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Datestamp of instance creation.`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `Datestamp of when instance was updated last.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the instance (operational, terminating, pending, stranded, etc.)`,
				},
				resource.Attribute{
					Name:        "href",
					Description: `Href of the instance.`,
				},
				resource.Attribute{
					Name:        "resource_uid",
					Description: `Cloud resource_uid as reported by cm platform.`,
				},
				resource.Attribute{
					Name:        "public_ip_addresses",
					Description: `List of public IP addresses associated to the instance`,
				},
				resource.Attribute{
					Name:        "private_ip_addresses",
					Description: `List of private IP addresses associated to the instance`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "rightscale_network",
			Category:         "Resources",
			ShortDescription: `Create and maintain a RightScale network.`,
			Description:      ``,
			Keywords: []string{
				"network",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "cloud_href",
					Description: `(Required) Cloud you want to create the network in.`,
				},
				resource.Attribute{
					Name:        "cidr_block",
					Description: `(Optional`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Network name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Network description.`,
				},
				resource.Attribute{
					Name:        "instance_tenancy",
					Description: `(Optional) Launch policy for AWS instances in the network. Specify 'dedicated' to force all instances to be launched as 'dedicated'. Defaults to 'default.'`,
				},
				resource.Attribute{
					Name:        "route_table_href",
					Description: `(Optional) Sets the default route table for this network, useful if you create the route table with a different resource.`,
				},
				resource.Attribute{
					Name:        "deployment_href",
					Description: `(Optional) Href of the deployment that owns the network. If you wish to use a deployment object as top level ownership construct, perhaps allocating the new network to a single deployment, then provide this href. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "href",
					Description: `Href of the network.`,
				},
				resource.Attribute{
					Name:        "resource_uid",
					Description: `Cloud resource_uid as reported by cm platform.`,
				},
				resource.Attribute{
					Name:        "links",
					Description: `Hrefs of related API resources.`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "href",
					Description: `Href of the network.`,
				},
				resource.Attribute{
					Name:        "resource_uid",
					Description: `Cloud resource_uid as reported by cm platform.`,
				},
				resource.Attribute{
					Name:        "links",
					Description: `Hrefs of related API resources.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "rightscale_network_gateway",
			Category:         "Resources",
			ShortDescription: `Create and maintain a RightScale network gateway.`,
			Description:      ``,
			Keywords: []string{
				"network",
				"gateway",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "cloud_href",
					Description: `(Required) Cloud you want to create the network gateway in.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) Type of network gateway. Options are "internet" or "vpc".`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Network gateway name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Network gateway description.`,
				},
				resource.Attribute{
					Name:        "network_href",
					Description: `(Optional) Href of network you want to attach the network gateway to. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "href",
					Description: `Href of the network gateway.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Date the network gateway was created at.`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `Date the network gateway was updated at.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `State of the network gateway. ("available" means attached to a network)`,
				},
				resource.Attribute{
					Name:        "resource_uid",
					Description: `Cloud resource_uid.`,
				},
				resource.Attribute{
					Name:        "links",
					Description: `Hrefs of related API resources.`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "href",
					Description: `Href of the network gateway.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Date the network gateway was created at.`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `Date the network gateway was updated at.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `State of the network gateway. ("available" means attached to a network)`,
				},
				resource.Attribute{
					Name:        "resource_uid",
					Description: `Cloud resource_uid.`,
				},
				resource.Attribute{
					Name:        "links",
					Description: `Hrefs of related API resources.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "rightscale_route",
			Category:         "Resources",
			ShortDescription: `Create and maintain a RightScale route.`,
			Description:      ``,
			Keywords: []string{
				"route",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "route_table_href",
					Description: `(Required) Href of route table in which to create new route.`,
				},
				resource.Attribute{
					Name:        "destination_cidr_block",
					Description: `(Required) Destination network in CIDR nodation.`,
				},
				resource.Attribute{
					Name:        "next_hop_type",
					Description: `(Required) The route next hop type. Options are 'instance', 'network_interface', 'network_gateway', 'ip_string', and 'url'.`,
				},
				resource.Attribute{
					Name:        "next_hop_href",
					Description: `(Contextual) The href of the Route's next hop. Required if 'next_hop_type' is 'instance', 'network_interface', or 'network_gateway'.`,
				},
				resource.Attribute{
					Name:        "next_hop_ip",
					Description: `(Contextual) The IP Address of the Route's next hop. Required if 'next_hop_type' is 'ip_string'.`,
				},
				resource.Attribute{
					Name:        "next_hop_url",
					Description: `(Contextual) The URL of the Route's next hop. Required if 'next_hop_type' is 'url'.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Route description. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "href",
					Description: `Href of the route.`,
				},
				resource.Attribute{
					Name:        "resource_uid",
					Description: `Route resource_uid.`,
				},
				resource.Attribute{
					Name:        "links",
					Description: `Hrefs of related API resources.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Created at datestamp.`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `Last updated at datestamp.`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "href",
					Description: `Href of the route.`,
				},
				resource.Attribute{
					Name:        "resource_uid",
					Description: `Route resource_uid.`,
				},
				resource.Attribute{
					Name:        "links",
					Description: `Hrefs of related API resources.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Created at datestamp.`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `Last updated at datestamp.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "rightscale_route_table",
			Category:         "Resources",
			ShortDescription: `Create and maintain a RightScale route_table.`,
			Description:      ``,
			Keywords: []string{
				"route",
				"table",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "cloud_href",
					Description: `(Required) Href of the cloud you want to create the route table in.`,
				},
				resource.Attribute{
					Name:        "network_href",
					Description: `(Required) Href of the network that owns the route table.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Route table name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Route table description. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "href",
					Description: `Href of the route table.`,
				},
				resource.Attribute{
					Name:        "resource_uid",
					Description: `Cloud resource_uid.`,
				},
				resource.Attribute{
					Name:        "links",
					Description: `Hrefs of related API resources.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Created at datestamp.`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `Last updated at datestamp.`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "href",
					Description: `Href of the route table.`,
				},
				resource.Attribute{
					Name:        "resource_uid",
					Description: `Cloud resource_uid.`,
				},
				resource.Attribute{
					Name:        "links",
					Description: `Hrefs of related API resources.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Created at datestamp.`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `Last updated at datestamp.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "rightscale_security_group",
			Category:         "Resources",
			ShortDescription: `Create and maintain a RightScale security_group.`,
			Description:      ``,
			Keywords: []string{
				"security",
				"group",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "cloud_href",
					Description: `(Required) Href of the cloud you want to create the security group in.`,
				},
				resource.Attribute{
					Name:        "network_href",
					Description: `(Required) Href of the network to create the security group in.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Security group name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Security group description.`,
				},
				resource.Attribute{
					Name:        "deployment_href",
					Description: `(Optional) Href of the deployment that owns the security group. If you wish to use a deployment object as top level ownership construct, perhaps allocating the new security group to a single deployment, then provide this href. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "href",
					Description: `Href of the security group.`,
				},
				resource.Attribute{
					Name:        "resource_uid",
					Description: `Cloud resource_uid.`,
				},
				resource.Attribute{
					Name:        "links",
					Description: `Hrefs of related API resources.`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "href",
					Description: `Href of the security group.`,
				},
				resource.Attribute{
					Name:        "resource_uid",
					Description: `Cloud resource_uid.`,
				},
				resource.Attribute{
					Name:        "links",
					Description: `Hrefs of related API resources.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "rightscale_security_group_rule",
			Category:         "Resources",
			ShortDescription: `Create and maintain a RightScale security_group_rule.`,
			Description:      ``,
			Keywords: []string{
				"security",
				"group",
				"rule",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "source_type",
					Description: `(Required) Source type. May be a CIDR block or another Security Group. Options are 'cidr_ips' or 'group'.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Required) Protocol to filter on. Options are 'tcp', 'udp', 'icmp' and 'all'.`,
				},
				resource.Attribute{
					Name:        "security_group_href",
					Description: `(Required) Href of parent security group.`,
				},
				resource.Attribute{
					Name:        "protocol_details",
					Description: `(Required) Block options include:`,
				},
				resource.Attribute{
					Name:        "cidr_ips",
					Description: `(Contextual) An IP address range in CIDR notation. Required if source_type is 'cidr'. Conflicts with 'group_name' and 'group_owner'.`,
				},
				resource.Attribute{
					Name:        "group_name",
					Description: `(Contextual) Name of source Security Group. Required if source_type is 'group'. Conflicts with 'cidr_ips'.`,
				},
				resource.Attribute{
					Name:        "group_owner",
					Description: `(Contexual) Owner of source Security Group. Required if source_type is 'group'. Conflicts with 'cidr_ips'.`,
				},
				resource.Attribute{
					Name:        "direction",
					Description: `(Optional) Direction of traffic to apply rule against. Options are 'ingress' or 'egress'.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `(Optional) Lower takes precedence. Supported by security group rules created in Microsoft Azure only. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "href",
					Description: `Href of the security group rule.`,
				},
				resource.Attribute{
					Name:        "resource_uid",
					Description: `Cloud resource_uid.`,
				},
				resource.Attribute{
					Name:        "links",
					Description: `Hrefs of related API resources.`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "href",
					Description: `Href of the security group rule.`,
				},
				resource.Attribute{
					Name:        "resource_uid",
					Description: `Cloud resource_uid.`,
				},
				resource.Attribute{
					Name:        "links",
					Description: `Hrefs of related API resources.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "rightscale_server",
			Category:         "Resources",
			ShortDescription: `Create and maintain a RightScale server`,
			Description:      ``,
			Keywords: []string{
				"server",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the server`,
				},
				resource.Attribute{
					Name:        "deployment_href",
					Description: `(Required) The href of the deployment the server will be placed in.`,
				},
				resource.Attribute{
					Name:        "instance",
					Description: `(Required) See [rightscale_instance](./cm_instance.html).`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A description of the server.`,
				},
				resource.Attribute{
					Name:        "optimized",
					Description: `(Optional) A flag indicating whether instances of this server should be optimized for high-performance volumes.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Any tags you want attached to the server and any instances created from this server object. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "links",
					Description: `Hrefs of related API resources`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Datestamp of server creation.`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `Datestamp of when server was updated last.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the server (operational, terminating, pending, stranded, etc.)`,
				},
				resource.Attribute{
					Name:        "href",
					Description: `Href of the server.`,
				},
				resource.Attribute{
					Name:        "resource_uid",
					Description: `Cloud resource_uid as reported by cm platform.`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "links",
					Description: `Hrefs of related API resources`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Datestamp of server creation.`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `Datestamp of when server was updated last.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the server (operational, terminating, pending, stranded, etc.)`,
				},
				resource.Attribute{
					Name:        "href",
					Description: `Href of the server.`,
				},
				resource.Attribute{
					Name:        "resource_uid",
					Description: `Cloud resource_uid as reported by cm platform.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "rightscale_server_array",
			Category:         "Resources",
			ShortDescription: `Create and maintain a RightScale server_array`,
			Description:      ``,
			Keywords: []string{
				"server",
				"array",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the server_array`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the server_array`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `(Required) he status of the server array. If enabled, the server array is enabled for scaling actions. One of "enabled" or "disabled"`,
				},
				resource.Attribute{
					Name:        "deployment_href",
					Description: `(Required) Href of deployment in which to create server_array`,
				},
				resource.Attribute{
					Name:        "array_type",
					Description: `(Required) The type of server_array. One of "alert" or "queue"`,
				},
				resource.Attribute{
					Name:        "optimized",
					Description: `(Optional) A flag indicating whether Instances of this ServerArray should be optimized for high-performance volumes (e.g. Volumes supporting a specified number of IOPS). Not supported in all Clouds.`,
				},
				resource.Attribute{
					Name:        "datacenter_policy",
					Description: `(Required) This is an array of datacenter policies. Each one must contain:`,
				},
				resource.Attribute{
					Name:        "datacenter_href",
					Description: `(Required) The href of the server_array's datacenter / zone.`,
				},
				resource.Attribute{
					Name:        "max",
					Description: `(Required) Maximum numbers of servers that can be allocated in this datacenter (0 for unlimited).`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `(Required) Instance allocation (should total 100% accross datacenter_policies).`,
				},
				resource.Attribute{
					Name:        "elasticity_params",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "bounds",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "min_count",
					Description: `(Required) The minimum number of servers that must be operational at all times in the server array.`,
				},
				resource.Attribute{
					Name:        "max_count",
					Description: `(Required) The maximum number of servers that can be operational at the same time in the server array.`,
				},
				resource.Attribute{
					Name:        "pacing",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "resize_down_by",
					Description: `(Required) The number of servers to scale down by.`,
				},
				resource.Attribute{
					Name:        "resize_up_by",
					Description: `(Required) The number of servers to scale up by.`,
				},
				resource.Attribute{
					Name:        "resize_calm_time",
					Description: `(Optional) The time (in minutes) on how long you want to wait before you repeat another action.`,
				},
				resource.Attribute{
					Name:        "alert_specific_params",
					Description: `(Required if alert array_type specified)`,
				},
				resource.Attribute{
					Name:        "decision_threshold",
					Description: `(Required) The percentage of servers that must agree in order to trigger an alert before an action is taken.`,
				},
				resource.Attribute{
					Name:        "voters_tag_predicate",
					Description: `(Optional) The Voters Tag that RightScale will use in order to determine when to scale up/down.`,
				},
				resource.Attribute{
					Name:        "queue_specific_params",
					Description: `(Required if queue alert_type specified)`,
				},
				resource.Attribute{
					Name:        "collect_audit_entries",
					Description: `(Optional) The audit SQS queue that will store audit entries.`,
				},
				resource.Attribute{
					Name:        "item_age",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "algorithm",
					Description: `(Optional) The algorithm that defines how an item's age will be determined, either by the average age or max (oldest) age.`,
				},
				resource.Attribute{
					Name:        "max_age",
					Description: `(Optional) The threshold (in seconds) before a resize action occurs on the server array.`,
				},
				resource.Attribute{
					Name:        "regexp",
					Description: `(Optional) The regexp that helps the system determine an item's \"age\" in the queue. Example: created_at: (\\d\\d\\d\\d-\\d\\d-\\d\\d \\d\\d:\\d\\d:\\d\\d UTC)`,
				},
				resource.Attribute{
					Name:        "queue_size",
					Description: `(Required) Defines the ratio of worker instances per items in the queue. Example: If there are 50 items in the queue and \"Items per instance\" is set to 10, the server array will resize to 5 worker instances (50/10). Default = 1`,
				},
				resource.Attribute{
					Name:        "schedule",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "day",
					Description: `(Required) Specifies the day when an alert-based array resizes. One of "Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday".`,
				},
				resource.Attribute{
					Name:        "max_count",
					Description: `(Required) The maximum number of servers that must be operational at all times in the server array. NOTE: Any changes that are made to the min/max count in the server array schedule will overwrite the array's default min/max count settings.`,
				},
				resource.Attribute{
					Name:        "min_count",
					Description: `(Required) The minimum number of servers that must be operational at all times in the server array. NOTE: Any changes that are made to the min/max count in the server array schedule will overwrite the array's default min/max count settings.`,
				},
				resource.Attribute{
					Name:        "time",
					Description: `(Required) Specifies the time when an alert-based array resizes.`,
				},
				resource.Attribute{
					Name:        "instance",
					Description: `(Required) See [rightscale_instance](https://github.com/terraform-providers/terraform-provider-rightscale/blob/master/website/docs/r/cm_instance.markdown) ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "links",
					Description: `Hrefs of related API resources`,
				},
				resource.Attribute{
					Name:        "href",
					Description: `Href of the server_array.`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "links",
					Description: `Hrefs of related API resources`,
				},
				resource.Attribute{
					Name:        "href",
					Description: `Href of the server_array.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "rightscale_ssh_key",
			Category:         "Resources",
			ShortDescription: `Create and maintain an ssh key resource in a given cloud.`,
			Description:      ``,
			Keywords: []string{
				"ssh",
				"key",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "cloud_href",
					Description: `(Required) The href of the cloud with the ssh key you want.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) SSH Key name. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "resource_uid",
					Description: `Cloud resource_uid.`,
				},
				resource.Attribute{
					Name:        "links",
					Description: `Hrefs of related API resources.`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "resource_uid",
					Description: `Cloud resource_uid.`,
				},
				resource.Attribute{
					Name:        "links",
					Description: `Hrefs of related API resources.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "rightscale_subnet",
			Category:         "Resources",
			ShortDescription: `Create and maintain a RightScale subnet.`,
			Description:      ``,
			Keywords: []string{
				"subnet",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "cloud_href",
					Description: `(Required) Href of cloud you want to create the subnet in.`,
				},
				resource.Attribute{
					Name:        "network_href",
					Description: `(Required) Href of network to create subnet in.`,
				},
				resource.Attribute{
					Name:        "cidr_block",
					Description: `(Required) Subnet allocation range in CIDR notation.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Subnet name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Subnet description.`,
				},
				resource.Attribute{
					Name:        "datacenter_href",
					Description: `(Optional) Href of cloud datacenter to assign subnet to.`,
				},
				resource.Attribute{
					Name:        "route_table_href",
					Description: `(Optional) Sets the default route table for this subnet, useful if you create the route table with a different resource. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "href",
					Description: `Href of the subnet.`,
				},
				resource.Attribute{
					Name:        "resource_uid",
					Description: `Cloud resource_uid.`,
				},
				resource.Attribute{
					Name:        "is_default",
					Description: `Indicates whether the subnet is the network default subnet. (true or false)`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `Indicates whether subnet is pending, available etc.`,
				},
				resource.Attribute{
					Name:        "links",
					Description: `Hrefs of related API resources.`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "href",
					Description: `Href of the subnet.`,
				},
				resource.Attribute{
					Name:        "resource_uid",
					Description: `Cloud resource_uid.`,
				},
				resource.Attribute{
					Name:        "is_default",
					Description: `Indicates whether the subnet is the network default subnet. (true or false)`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `Indicates whether subnet is pending, available etc.`,
				},
				resource.Attribute{
					Name:        "links",
					Description: `Hrefs of related API resources.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "rightscale_cwf_process",
			Category:         "Resources",
			ShortDescription: `Create and maintain a RightScale CloudWorkFlow process.`,
			Description:      ``,
			Keywords: []string{
				"cwf",
				"process",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "source",
					Description: `(Required) Source code to be executed, written in [RCL (RightScale CloudWorkFlow Language)](http://docs.rightscale.com/ss/reference/rcl/v2/index.html). Several functions can be defined but the entry function should be called ` + "`" + `main` + "`" + `. Example: ` + "`" + `` + "`" + `` + "`" + `hcl source = <<EOF define adder($n1, $n2) return $res do $res = $n1 + $n2 end define main($a, $b) return $result do call adder($a, $b) retrieve $tmp $result = "The total is " + $tmp end EOF ` + "`" + `` + "`" + `` + "`" + ``,
				},
				resource.Attribute{
					Name:        "parameters",
					Description: `Parameters for the RCL function. It consists of an array of values corresponding to the values being passed to the function defined in the "source" field in order of declaration. The values are defined as string maps with the "kind" and "value" keys. "kind" contains the type of the value being passed, could be one of "array", "boolean", "collection", "datetime", "declaration", "null", "number", "object", "string". The "value" key contains the value For example: ` + "`" + `` + "`" + `` + "`" + `hcl parameters = [ { "kind" = "string" "value" = "db-slave-" }, { "kind" = "number" "value" = "42" } ] ` + "`" + `` + "`" + `` + "`" + ` Note that the "value" key should always be a string (regardless of the type specified in "kind"). These are several examples on how to pass arrays: ` + "`" + `` + "`" + `` + "`" + `hcl parameters = [ { "kind" = "array" "value" = "[ 22.3, 9.7, 10 ]" }, { "kind" = "array" "value" = "[ \"It\", \" works!\" ]" }, { "kind" = "array" "value" = "${jsonencode(var.zones)}" }, ] ` + "`" + `` + "`" + `` + "`" + ` ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Process status, one of "completed", "failed", "canceled" or "aborted".`,
				},
				resource.Attribute{
					Name:        "error",
					Description: `Process execution error if any.`,
				},
				resource.Attribute{
					Name:        "outputs",
					Description: `Process outputs if any. This is a TypeMap, one particular output can be accessed via ` + "`" + `outputs["$var"]` + "`" + `, see "Example Usage" section.`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "status",
					Description: `Process status, one of "completed", "failed", "canceled" or "aborted".`,
				},
				resource.Attribute{
					Name:        "error",
					Description: `Process execution error if any.`,
				},
				resource.Attribute{
					Name:        "outputs",
					Description: `Process outputs if any. This is a TypeMap, one particular output can be accessed via ` + "`" + `outputs["$var"]` + "`" + `, see "Example Usage" section.`,
				},
			},
		},
	}

	resourcesMap = map[string]int{

		"rightscale_credential":          0,
		"rightscale_deployment":          1,
		"rightscale_instance":            2,
		"rightscale_network":             3,
		"rightscale_network_gateway":     4,
		"rightscale_route":               5,
		"rightscale_route_table":         6,
		"rightscale_security_group":      7,
		"rightscale_security_group_rule": 8,
		"rightscale_server":              9,
		"rightscale_server_array":        10,
		"rightscale_ssh_key":             11,
		"rightscale_subnet":              12,
		"rightscale_cwf_process":         13,
	}
)

func GetResource(r string) (*resouce.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs]
}
