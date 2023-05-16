package stackpath

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "stackpath_compute_network_policy",
			Category:         "Resources",
			ShortDescription: `Network ingress and egress control of StackPath computing workloads.`,
			Description:      ``,
			Keywords: []string{
				"compute",
				"network",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A human readable name.`,
				},
				resource.Attribute{
					Name:        "slug",
					Description: `(Required) A programmatic name for the network policy.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A brief description.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) Key/value pairs of arbitrary label names and values that can be referenced as selectors.`,
				},
				resource.Attribute{
					Name:        "annotations",
					Description: `(Optional) Key/value pairs that define StackPath-specific network policy configuration.`,
				},
				resource.Attribute{
					Name:        "instance_selector",
					Description: `(Optional) A compute workload instance that the network policy applies to. A network policy with no selectors applies to all networks and all instances in the stack. See [Selectors](#selectors) below for details.`,
				},
				resource.Attribute{
					Name:        "network_selector",
					Description: `(Optional) A network that the network policy applies to. A network policy with no selectors applies to all networks and all instances in the stack. See [Selectors](#selectors) below for details.`,
				},
				resource.Attribute{
					Name:        "policy_types",
					Description: `(Required) A list of network policy types, either "INGRESS" and/or "EGRESS".`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `(Required) A priority value between 1 and 65000. Higher priority network policies override lower priority policies, and priorities must be unique across the stack.`,
				},
				resource.Attribute{
					Name:        "egress",
					Description: `(Optional) Outbound networking information. See [Egress](#egress) below for details.`,
				},
				resource.Attribute{
					Name:        "ingress",
					Description: `(Optional) Inbound networking information. See [Ingress](#ingress) below for details. ### Egress ` + "`" + `egress` + "`" + ` takes the following arguments:`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Required) How a network policy treats outbound traffic, either "ALLOW" or "BLOCK".`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A brief description.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Optional) Network protocol specific information. See [Network Protocols](#network-protocols) below for details.`,
				},
				resource.Attribute{
					Name:        "to",
					Description: `(Optional) Allow or block outbound traffic to the specified targets. See [Network Selectors](#network-selectors) below for details. ### Ingress ` + "`" + `ingress` + "`" + ` takes the following arguments:`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Required) How a network policy treats outbound traffic, either "ALLOW" or "BLOCK".`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A brief description.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Optional) Network protocol specific information. See [Network Protocols](#network-protocols) below for details.`,
				},
				resource.Attribute{
					Name:        "from",
					Description: `(Optional) Allow or block inbound traffic from the specified targets. See [Network Selectors](#network-selectors) below for details. ### Network Protocols ` + "`" + `protocol` + "`" + ` takes the following arguments:`,
				},
				resource.Attribute{
					Name:        "ah",
					Description: `(Optional) Allow or block the IPSec Authentication Header protocol. This argument block has no configuration.`,
				},
				resource.Attribute{
					Name:        "esp",
					Description: `(Optional) Allow or block the IPSec Encapsulating Security Payload protocol. This argument block has no configuration.`,
				},
				resource.Attribute{
					Name:        "gre",
					Description: `(Optional) Allow or block the Generic Routing Encapsulation protocol. This argument block has no configuration.`,
				},
				resource.Attribute{
					Name:        "icmp",
					Description: `(Optional) Allow or block the Internet Control Message Protocol. This argument block has no configuration.`,
				},
				resource.Attribute{
					Name:        "tcp",
					Description: `(Optional) Allow or block Transmission Control Protocol connections. See [Network Ports](#network-ports) below for details.`,
				},
				resource.Attribute{
					Name:        "udp",
					Description: `(Optional) Allow or block User Datagram Protocol connections. See [Network Ports](#network-ports) below for details.`,
				},
				resource.Attribute{
					Name:        "tcp_udp",
					Description: `(Optional) Allow or block both TCP and UDP connections. See [Network Ports](#network-ports) below for details. ### Network Ports ` + "`" + `tcp` + "`" + `, ` + "`" + `udp` + "`" + `, and ` + "`" + `tcp_udp` + "`" + ` take the following arguments:`,
				},
				resource.Attribute{
					Name:        "source_ports",
					Description: `(Optional) A list of destination ports.`,
				},
				resource.Attribute{
					Name:        "destination_ports",
					Description: `(Optional) A list of destination ports. ### Network Selectors ` + "`" + `to` + "`" + ` and ` + "`" + `from` + "`" + ` take the following arguments:`,
				},
				resource.Attribute{
					Name:        "instance_selector",
					Description: `(Optional) Target the given compute workload instances. See [Selectors](#selectors) below for details.`,
				},
				resource.Attribute{
					Name:        "network_selector",
					Description: `(Optional) Target the given networks. See [Selectors](#selectors) below for details.`,
				},
				resource.Attribute{
					Name:        "ip_block",
					Description: `(Optional) Target the given IP address blocks. See [IP Blocks](#ip-blocks) below for details. #### IP Blocks ` + "`" + `ip_block` + "`" + ` takes the following arguments:`,
				},
				resource.Attribute{
					Name:        "cidr",
					Description: `(Required) A CIDR formatted subnet.`,
				},
				resource.Attribute{
					Name:        "except",
					Description: `(Optional) A list of CIDR formatted subnets to exclude from the ` + "`" + `cidr` + "`" + ` subnet. ### Selectors ` + "`" + `instance_selector` + "`" + ` and ` + "`" + `network_selector` + "`" + ` take the following arguments:`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) The name of the data that a selector is based on.`,
				},
				resource.Attribute{
					Name:        "operator",
					Description: `(Required) A logical operator to apply to a selector. Only the "in" operator is supported.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) Data values to look for in a label selector. ## Import StackPath compute network policies can be imported by their UUID v4 formatted id. e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import stackpath_compute_network_policy.terraform bdb77768-2938-4ad8-a736-be5290add801 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "stackpath_compute_workload",
			Category:         "Resources",
			ShortDescription: `A computing application deployed to StackPath's edge network.`,
			Description:      ``,
			Keywords: []string{
				"compute",
				"workload",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A human readable name.`,
				},
				resource.Attribute{
					Name:        "slug",
					Description: `(Required) A programmatic name for the workload. Workload slugs are used to build the workload's instance names and cannot be changed after creation.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) Key/value pairs of arbitrary label names and values that can be referenced as [selectors](#selectors) by [network policies](/docs/providers/stackpath/r/compute_network_policy.html).`,
				},
				resource.Attribute{
					Name:        "annotations",
					Description: `(Optional) Key/value pairs that define StackPath-specific workload configuration.`,
				},
				resource.Attribute{
					Name:        "network_interface",
					Description: `(Required) Networks to place the compute instance on. See [Network Interfaces](#network-interfaces) below for details.`,
				},
				resource.Attribute{
					Name:        "image_pull_credentials",
					Description: `(Optional) Credentials to pull container images with. See [Image Pull Credentials](#image-pull-credentials) below for details.`,
				},
				resource.Attribute{
					Name:        "virtual_machine",
					Description: `(Optional) Virtual machine configuration. StackPath supports a single virtual machine specification in a workload. At least one of ` + "`" + `virtual_machine` + "`" + ` or ` + "`" + `container` + "`" + ` must be provided. See [Virtual Machines](#virtual-machines) below for details.`,
				},
				resource.Attribute{
					Name:        "container",
					Description: `(Optional) Container configuration. At least one of ` + "`" + `virtual_machine` + "`" + ` or ` + "`" + `container` + "`" + ` must be provided. See [Containers](#containers) below for details.`,
				},
				resource.Attribute{
					Name:        "volume_claim",
					Description: `(Optional) Storage that is mounted to a compute workload's instances. See [Volume Claims](#volume-claims) below for details.`,
				},
				resource.Attribute{
					Name:        "target",
					Description: `(Required) How the compute workload should be deployed across the StackPath edge platform. See [Deployment Targets](#deployment-targets) below for details. ### Network Interfaces ` + "`" + `network_interfaces` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "network",
					Description: `(Required) A name that can be referenced by a [selector](#selectors) by [network policies](/docs/providers/stackpath/r/compute_network_policy.html). Both default and user created VPC networks are supported.`,
				},
				resource.Attribute{
					Name:        "enable_one_to_one_nat",
					Description: `(Optional) Boolean Flag to specify enabling of one to one nat to interface VPC IP address. Default is true.`,
				},
				resource.Attribute{
					Name:        "subnet",
					Description: `(Optional) A name of IPv4 subnet to be used for IPv4 IP allocation for interface. subnet should belong to network specified in ` + "`" + `network` + "`" + ` field. if not specified then default IPv4 subnet will be used.`,
				},
				resource.Attribute{
					Name:        "ipv6_subnet",
					Description: `(Optional) A name of IPv6 subnet to be used for IPv6 IP allocation for interface. subnet should belong to network specified in ` + "`" + `network` + "`" + ` field. if not specified then default IPv6 subnet will be used.`,
				},
				resource.Attribute{
					Name:        "ip_families",
					Description: `(Optional) List of IP Families from which IP allocation should happen to network interface. Default is [IPv4]. Currently this supports IPv4-only [IPv4] and Dual stack- [IPv4, IPv6]. It does not suport IPv6 only- [IPv6] requests. If Dual stack [IPv4, IPv6] is requested then VPC network specified in ` + "`" + `network` + "`" + ` field is expected to enabled for Dual stack networking. Bu default all Default networks are enabled for Dual stack, any user created VPC network needs to be created with [IPv4, IPv6] IPFamilies to enable it for Dual stack networking. ### Image Pull Credentials ` + "`" + `image_pull_credentials` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "docker_registry",
					Description: `(Required) Authentication configuration needed to pull images from a Docker registry. See [Docker Registry Credentials](#docker-registry-credentials) below for details. #### Docker Registry Credentials ` + "`" + `docker_registry` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "server",
					Description: `(Optional) The address of a Docker registry server. Defaults to "hub.docker.com".`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Required) A username to connect the Docker registry.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Required) A password to connect to the Docker registry.`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `(Optional) An email address to use with the Docker registry account. ### Virtual Machines ` + "`" + `virtual_machine` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A virtual machine's name.`,
				},
				resource.Attribute{
					Name:        "image",
					Description: `(Required) The disk image to run as a virtual machine.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional) Network ports to expose from the virtual machine. Ports can also be used for internal DNS-based service discovery. See [Network Ports](#network-ports) below for details.`,
				},
				resource.Attribute{
					Name:        "liveness_probe",
					Description: `(Optional) Criteria to determine if the compute workload is online. See [Probes](#probes) below for details.`,
				},
				resource.Attribute{
					Name:        "readiness_probe",
					Description: `(Optional) Criteria to determine if the compute workload is ready to serve requests. See [Probes](#probes) below for details.`,
				},
				resource.Attribute{
					Name:        "resources",
					Description: `(Required) Hardware resources required by the virtual machine. See [Resources](#resources) below for details.`,
				},
				resource.Attribute{
					Name:        "volume_mount",
					Description: `(Optional) Storage volumes to mount in the virtual machine. See [Volume Mounts](#volume-mounts) below for details.`,
				},
				resource.Attribute{
					Name:        "user_data",
					Description: `(Optional) [Cloud-init](https://cloud-init.io/) user data. ### Containers ` + "`" + `container` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A container's name.`,
				},
				resource.Attribute{
					Name:        "image",
					Description: `(Required) A container's image location.`,
				},
				resource.Attribute{
					Name:        "command",
					Description: `(Optional) The command to execute a container.`,
				},
				resource.Attribute{
					Name:        "env",
					Description: `(Optional) Environment variables to set in the container instance. See [Environment Variables](#environment-variables) below for details.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional) Networking ports to expose from the container. Ports can also be used for internal DNS-based service discovery. See [Network Ports](#network-ports) below for details.`,
				},
				resource.Attribute{
					Name:        "liveness_probe",
					Description: `(Optional) Criteria to determine if the compute workload is online. See [Probes](#probes) below for details.`,
				},
				resource.Attribute{
					Name:        "readiness_probe",
					Description: `(Optional) Criteria to determine if the compute workload is ready to serve requests. See [Probes](#probes) below for details.`,
				},
				resource.Attribute{
					Name:        "resources",
					Description: `(Required) Hardware resources required by the container. See [Resources](#resources) below for details.`,
				},
				resource.Attribute{
					Name:        "volume_mount",
					Description: `(Optional) Storage volumes to mount in the container. See [Volume Mounts](#volume-mounts) below for details. #### Environment Variables ` + "`" + `env` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) The environment variable name.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Optional) The environment variable value. One of ` + "`" + `value` + "`" + ` or ` + "`" + `secret_value` + "`" + ` must be provided.`,
				},
				resource.Attribute{
					Name:        "secret_value",
					Description: `(Optional) A sensitive environment variable value. This value cannot be read after it is set. One of ` + "`" + `value` + "`" + ` or ` + "`" + `secret_value` + "`" + ` must be provided. ### Network Ports ` + "`" + `port` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The network port's name.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Required) The network port's number.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Optional) The network port's protocol, either "tcp" or "udp". Defaults to "tcp".`,
				},
				resource.Attribute{
					Name:        "enable_implicit_network_policy",
					Description: `(Optional) Whether or not the network port is accessible from the public Internet. Defaults to ` + "`" + `false` + "`" + `. ### Volume Claims ` + "`" + `volume_claim` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A human readable name.`,
				},
				resource.Attribute{
					Name:        "slug",
					Description: `(Optional) A programmatic slug. Reference this slug when [mounting](#volume-mounts) the claim into a workload's instances.`,
				},
				resource.Attribute{
					Name:        "resources",
					Description: `(Required) Hardware resources to allocate to the volume claim. See [Resources](#resources) below for details. ### Probes ` + "`" + `liveness_probe` + "`" + ` and ` + "`" + `readiness_probe` + "`" + ` take the following arguments:`,
				},
				resource.Attribute{
					Name:        "http_get",
					Description: `(Optional) HTTP request information. One of ` + "`" + `http_get` + "`" + ` or ` + "`" + `tcp_socket` + "`" + ` must be provided. See [HTTP probes](#http-probes) below for details`,
				},
				resource.Attribute{
					Name:        "tcp_socket",
					Description: `(Optional) TCP socket information. One of ` + "`" + `http_get` + "`" + ` or ` + "`" + `tcp_socket` + "`" + ` must be provided. See [TCP probes](#tcp-probes) below for details`,
				},
				resource.Attribute{
					Name:        "initial_delay_seconds",
					Description: `(Optional) The initial delay before the probe starts. Defaults to 0.`,
				},
				resource.Attribute{
					Name:        "timeout_seconds",
					Description: `(Optional) The number of seconds before the probe times out and is considered a failure. Defaults to 10.`,
				},
				resource.Attribute{
					Name:        "period",
					Description: `(Optional) The frequency of the probe. Defaults to 60.`,
				},
				resource.Attribute{
					Name:        "success_threshold",
					Description: `(Required) The minimum consecutive successes required before a probe is considered successful after a failure. This must be 1 for liveness probes.`,
				},
				resource.Attribute{
					Name:        "failure_threshold",
					Description: `(Required) The amount of failures seen before the probe is considered a failure. #### HTTP Probes ` + "`" + `http_get` + "`" + ` takes the following arguments:`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Optional) The URL path to request from the application. Defaults to "/".`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Required) The TCP port the HTTP service listens on.`,
				},
				resource.Attribute{
					Name:        "scheme",
					Description: `(Optional) The URL scheme to query the application with. Defaults to "http".`,
				},
				resource.Attribute{
					Name:        "http_headers",
					Description: `(Optional) HTTP header names and values to send to the HTTP service. #### TCP probes ` + "`" + `tcp_socket` + "`" + ` takes the following arguments:`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Required) The TCP port number to connect to. ### Resources ` + "`" + `resources` + "`" + ` takes the following arguments:`,
				},
				resource.Attribute{
					Name:        "requests",
					Description: `(Required) Key/value pairs of hardware resource types and values. ### Volume Mounts ` + "`" + `volume-mount` + "`" + ` takes the following arguments:`,
				},
				resource.Attribute{
					Name:        "slug",
					Description: `(Required) The slug of the [volume claim](#volume-claim) to mount into the workload's instances.`,
				},
				resource.Attribute{
					Name:        "mount_path",
					Description: `(Required) The path the volume is mounted to in a workload's instances. ### Deployment Targets ` + "`" + `target` + "`" + ` takes the following arguments:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A human readable name.`,
				},
				resource.Attribute{
					Name:        "min_replicas",
					Description: `(Required) The minimum number of instances that should be deployed to a target.`,
				},
				resource.Attribute{
					Name:        "max_replicas",
					Description: `(Optional) The maximum number of instances that should be deployed to a target.`,
				},
				resource.Attribute{
					Name:        "scale_settings",
					Description: `(Optional) How to auto-scale the number of instances in the deployment target. See [Scaling Settings](#scaling-settings) below for details.`,
				},
				resource.Attribute{
					Name:        "deployment_scope",
					Description: `(Optional) Criteria that defines a deployment target. Defaults to "cityCode".`,
				},
				resource.Attribute{
					Name:        "selector",
					Description: `(Required) The value of the deployment scope to deploy to. See [Selectors](#selectors) below for details. #### Scaling Settings ` + "`" + `scale_settings` + "`" + ` takes the following arguments:`,
				},
				resource.Attribute{
					Name:        "metrics",
					Description: `(Required) Scaling metrics. See [Scaling Metrics](#scaling-metrics) below for details. ##### Scaling Metrics ` + "`" + `metrics` + "`" + ` takes the following arguments:`,
				},
				resource.Attribute{
					Name:        "metric",
					Description: `(Required) A hardware metric to use as a scaling basis. Currently, only the "cpu" metric is supported.`,
				},
				resource.Attribute{
					Name:        "average_utilization",
					Description: `(Optional) The ` + "`" + `metric` + "`" + `'s average utilization that should trigger scaling. One of ` + "`" + `average_utilization` + "`" + ` or ` + "`" + `average_value` + "`" + ` must be provided.`,
				},
				resource.Attribute{
					Name:        "average_value",
					Description: `(Optional) The ` + "`" + `metric` + "`" + `'s average value that should trigger scaling. One of ` + "`" + `average_utilization` + "`" + ` or ` + "`" + `average_value` + "`" + ` must be provided. #### Selectors ` + "`" + `selector` + "`" + ` takes the following arguments:`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) The name of the data that a selector is based on.`,
				},
				resource.Attribute{
					Name:        "operator",
					Description: `(Required) A logical operator to apply to a selector. Only the "in" operator is supported.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) Data values to look for in a label selector. ## Instances A workload instance is a collection of containers or a virtual machine created based on the template provided in a workload. Instances are accessed via a ` + "`" + `stackpath_compute_workload` + "`" + `'s computed ` + "`" + `instances` + "`" + ` field. ### Example Usage ` + "`" + `` + "`" + `` + "`" + `hcl # Output a StackPath compute workload's instances' name, internal IP addresses, # and status output "my-compute-workload-instances" { value = { for instance in stackpath_compute_workload.my-compute-workload.instances: instance.name => { ip_address = instance.external_ip_address phase = instance.phase } } } ` + "`" + `` + "`" + `` + "`" + ` ### Instance Fields`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) An instance's name. Names are generated from their corresponding workload's slug, followed by a unique hash.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `(Optional) Metadata associated with a running instance, including the workload's ` + "`" + `labels` + "`" + ` and [annotations](#annotations), both supplied by the user and generated by StackPath.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `(Optional) The instance's physical location. See [Locations](#locations) below for details.`,
				},
				resource.Attribute{
					Name:        "external_ip_address",
					Description: `(Optional) An IPv4 address bound to the instance.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `(Optional) An instance's internal IPv4 address.`,
				},
				resource.Attribute{
					Name:        "external_ipv6_address",
					Description: `(Optional) An IPv6 address bound to the instance.`,
				},
				resource.Attribute{
					Name:        "ipv6_address",
					Description: `(Optional) An instance's internal IPv6 address.`,
				},
				resource.Attribute{
					Name:        "network_interface",
					Description: `(Optional) A network interface bound to an instance. See [Instance Network Interfaces](#instance-network-interfaces) below for details.`,
				},
				resource.Attribute{
					Name:        "virtual_machine",
					Description: `(Optional) An instance's [virtual machine](#virtual-machines) specification. An instance has either a ` + "`" + `virtual_machine` + "`" + ` or ` + "`" + `container` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "container",
					Description: `(Optional) An instance's [container](#containers) specification. An instance has either a ` + "`" + `virtual_machine` + "`" + ` or ` + "`" + `container` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "phase",
					Description: `(Optional) An instance's current status, such as "STARTING", "RUNNING", "FAILED", or "STOPPED".`,
				},
				resource.Attribute{
					Name:        "reason",
					Description: `(Optional) A short reason why an instance is in its current phase.`,
				},
				resource.Attribute{
					Name:        "message",
					Description: `(Optional) A longer message with details why an instance is in its current phase. ### Locations ` + "`" + `locaton` + "`" + ` has the following fields:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) A location's name.`,
				},
				resource.Attribute{
					Name:        "city_code",
					Description: `(Optional) A city's [IATA code](https://en.wikipedia.org/wiki/IATA_airport_code).`,
				},
				resource.Attribute{
					Name:        "subdivision",
					Description: `(Optional) A location's subdivision.`,
				},
				resource.Attribute{
					Name:        "subdivision_code",
					Description: `(Optional) A subdivision's [ISO 3166-2](https://en.wikipedia.org/wiki/ISO_3166-2) code.`,
				},
				resource.Attribute{
					Name:        "country",
					Description: `(Optional) A location's country.`,
				},
				resource.Attribute{
					Name:        "country_code",
					Description: `(Optional) A country's [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) code.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) A location's region.`,
				},
				resource.Attribute{
					Name:        "region_code",
					Description: `(Optional) A region's GeoIP code.`,
				},
				resource.Attribute{
					Name:        "continent",
					Description: `(Optional) A location's continent.`,
				},
				resource.Attribute{
					Name:        "continent_code",
					Description: `(Optional) A continent's GeoIP code.`,
				},
				resource.Attribute{
					Name:        "latitude",
					Description: `(Optional) A location's latitude coordinate.`,
				},
				resource.Attribute{
					Name:        "longitude",
					Description: `(Optional) A location's longitude coordinate. ### Instance Network Interfaces ` + "`" + `network_interface` + "`" + ` has the following fields:`,
				},
				resource.Attribute{
					Name:        "network",
					Description: `(Required) The name of the [workload network interface](#network-interfaces).`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `(Required) A network interface's primary IPv4 address.`,
				},
				resource.Attribute{
					Name:        "ip_address_aliases",
					Description: `(Optional) Additional IPv4 addresses bound to a network interface.`,
				},
				resource.Attribute{
					Name:        "gateway",
					Description: `(Required) A network interface IPv4 subnet's gateway IP address.`,
				},
				resource.Attribute{
					Name:        "ipv6_address",
					Description: `(Optional) A network interface's primary IPv6 address.`,
				},
				resource.Attribute{
					Name:        "ipv6_address_aliases",
					Description: `(Optional) Additional IPv6 addresses bound to a network interface.`,
				},
				resource.Attribute{
					Name:        "ipv6_gateway",
					Description: `(Optional) A network interface IPv6 subnet's gateway IP address. ## Import StackPath compute workloads can be imported by their UUID v4 formatted id. e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import stackpath_compute_workload.terraform bdb77768-2938-4ad8-a736-be5290add801 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "stackpath_object_storage_bucket",
			Category:         "Resources",
			ShortDescription: `An S3 compatible object storage bucket deployed to StackPath's edge network.`,
			Description:      ``,
			Keywords: []string{
				"object",
				"storage",
				"bucket",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "label",
					Description: `(Required) A human readable label for the bucket. Bucket label only supports (a-z, 0-9, -) and must start/end with a letter or number.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Required) Bucket region (us-east-2 us-west-1 eu-central-1)`,
				},
				resource.Attribute{
					Name:        "visibility",
					Description: `(Optional) PRIVATE or PUBLIC, defaults to PRIVATE ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "endpoint_url",
					Description: `S3 compatible region endpoint for the bucket, e.g. https://s3.us-east-2.stackpathstorage.com ## Import StackPath object storage buckets can be imported by their UUID v4 formatted id. e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import stackpath_object_storage_bucket.bucket bdb77768-2938-4ad8-a736-be5290add801 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "endpoint_url",
					Description: `S3 compatible region endpoint for the bucket, e.g. https://s3.us-east-2.stackpathstorage.com ## Import StackPath object storage buckets can be imported by their UUID v4 formatted id. e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import stackpath_object_storage_bucket.bucket bdb77768-2938-4ad8-a736-be5290add801 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
	}

	resourcesMap = map[string]int{

		"stackpath_compute_network_policy": 0,
		"stackpath_compute_workload":       1,
		"stackpath_object_storage_bucket":  2,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
