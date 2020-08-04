package docker

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "docker_config",
			Category:         "Swarm Resources",
			ShortDescription: `Manages the configs of a Docker service in a swarm.`,
			Description:      ``,
			Keywords: []string{
				"swarm",
				"config",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required, string) The name of the Docker config.`,
				},
				resource.Attribute{
					Name:        "data",
					Description: `(Required, string) The base64 encoded data of the config. ## Attributes Reference The following attributes are exported in addition to the above configuration:`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "docker_container",
			Category:         "Resources",
			ShortDescription: `Manages the lifecycle of a Docker container.`,
			Description:      ``,
			Keywords: []string{
				"container",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required, string) The name of the Docker container.`,
				},
				resource.Attribute{
					Name:        "image",
					Description: `(Required, string) The ID of the image to back this container. The easiest way to get this value is to use the ` + "`" + `docker_image` + "`" + ` resource as is shown in the example above.`,
				},
				resource.Attribute{
					Name:        "command",
					Description: `(Optional, list of strings) The command to use to start the container. For example, to run ` + "`" + `/usr/bin/myprogram -f baz.conf` + "`" + ` set the command to be ` + "`" + `["/usr/bin/myprogram", "-f", "baz.conf"]` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "entrypoint",
					Description: `(Optional, list of strings) The command to use as the Entrypoint for the container. The Entrypoint allows you to configure a container to run as an executable. For example, to run ` + "`" + `/usr/bin/myprogram` + "`" + ` when starting a container, set the entrypoint to be ` + "`" + `["/usr/bin/myprogram"]` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "user",
					Description: `(Optional, string) User used for run the first process. Format is ` + "`" + `user` + "`" + ` or ` + "`" + `user:group` + "`" + ` which user and group can be passed literraly or by name.`,
				},
				resource.Attribute{
					Name:        "dns",
					Description: `(Optional, set of strings) Set of DNS servers.`,
				},
				resource.Attribute{
					Name:        "dns_opts",
					Description: `(Optional, set of strings) Set of DNS options used by the DNS provider(s), see ` + "`" + `resolv.conf` + "`" + ` documentation for valid list of options.`,
				},
				resource.Attribute{
					Name:        "dns_search",
					Description: `(Optional, set of strings) Set of DNS search domains that are used when bare unqualified hostnames are used inside of the container.`,
				},
				resource.Attribute{
					Name:        "env",
					Description: `(Optional, set of strings) Environment variables to set.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional, block) See [Labels](#labels-1) below for details.`,
				},
				resource.Attribute{
					Name:        "links",
					Description: `(Optional, set of strings) Set of links for link based connectivity between containers that are running on the same host. ~>`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `(Optional, string) Hostname of the container.`,
				},
				resource.Attribute{
					Name:        "domainname",
					Description: `(Optional, string) Domain name of the container.`,
				},
				resource.Attribute{
					Name:        "restart",
					Description: `(Optional, string) The restart policy for the container. Must be one of "no", "on-failure", "always", "unless-stopped".`,
				},
				resource.Attribute{
					Name:        "max_retry_count",
					Description: `(Optional, int) The maximum amount of times to an attempt a restart when ` + "`" + `restart` + "`" + ` is set to "on-failure"`,
				},
				resource.Attribute{
					Name:        "rm",
					Description: `(Optional, bool) If true, then the container will be automatically removed after his execution. Terraform won't check this container after creation.`,
				},
				resource.Attribute{
					Name:        "read_only",
					Description: `(Optional, bool) If true, the container will be started as readonly.`,
				},
				resource.Attribute{
					Name:        "start",
					Description: `(Optional, bool) If true, then the Docker container will be started after creation. If false, then the container is only created.`,
				},
				resource.Attribute{
					Name:        "attach",
					Description: `(Optional, bool) If true attach to the container after its creation and waits the end of his execution.`,
				},
				resource.Attribute{
					Name:        "logs",
					Description: `(Optional, bool) Save the container logs (` + "`" + `attach` + "`" + ` must be enabled).`,
				},
				resource.Attribute{
					Name:        "must_run",
					Description: `(Optional, bool) If true, then the Docker container will be kept running. If false, then as long as the container exists, Terraform assumes it is successful.`,
				},
				resource.Attribute{
					Name:        "capabilities",
					Description: `(Optional, block) See [Capabilities](#capabilities-1) below for details.`,
				},
				resource.Attribute{
					Name:        "mounts",
					Description: `(Optional, set of blocks) See [Mounts](#mounts-1) below for details.`,
				},
				resource.Attribute{
					Name:        "tmpfs",
					Description: `(Optional, map) A map of container directories which should be replaced by ` + "`" + `tmpfs mounts` + "`" + `, and their corresponding mount options.`,
				},
				resource.Attribute{
					Name:        "ports",
					Description: `(Optional, block) See [Ports](#ports-1) below for details.`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `(Optional, block) See [Extra Hosts](#extra_hosts-1) below for details.`,
				},
				resource.Attribute{
					Name:        "privileged",
					Description: `(Optional, bool) Run container in privileged mode.`,
				},
				resource.Attribute{
					Name:        "devices",
					Description: `(Optional, bool) See [Devices](#devices-1) below for details.`,
				},
				resource.Attribute{
					Name:        "publish_all_ports",
					Description: `(Optional, bool) Publish all ports of the container.`,
				},
				resource.Attribute{
					Name:        "volumes",
					Description: `(Optional, block) See [Volumes](#volumes-1) below for details.`,
				},
				resource.Attribute{
					Name:        "memory",
					Description: `(Optional, int) The memory limit for the container in MBs.`,
				},
				resource.Attribute{
					Name:        "memory_swap",
					Description: `(Optional, int) The total memory limit (memory + swap) for the container in MBs. This setting may compute to ` + "`" + `-1` + "`" + ` after ` + "`" + `terraform apply` + "`" + ` if the target host doesn't support memory swap, when that is the case docker will use a soft limitation.`,
				},
				resource.Attribute{
					Name:        "shm_size",
					Description: `(Optional, int) Size of ` + "`" + `/dev/shm` + "`" + ` in MBs.`,
				},
				resource.Attribute{
					Name:        "cpu_shares",
					Description: `(Optional, int) CPU shares (relative weight) for the container.`,
				},
				resource.Attribute{
					Name:        "cpu_set",
					Description: `(Optional, string) A comma-separated list or hyphen-separated range of CPUs a container can use, e.g. ` + "`" + `0-1` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "log_driver",
					Description: `(Optional, string) The logging driver to use for the container. Defaults to "json-file".`,
				},
				resource.Attribute{
					Name:        "log_opts",
					Description: `(Optional, map of strings) Key/value pairs to use as options for the logging driver.`,
				},
				resource.Attribute{
					Name:        "network_alias",
					Description: `(Optional, set of strings) Network aliases of the container for user-defined networks only.`,
				},
				resource.Attribute{
					Name:        "network_mode",
					Description: `(Optional, string) Network mode of the container.`,
				},
				resource.Attribute{
					Name:        "networks",
					Description: `(Optional, set of strings) Id of the networks in which the container is.`,
				},
				resource.Attribute{
					Name:        "networks_advanced",
					Description: `(Optional, block) See [Networks Advanced](#networks_advanced-1) below for details. If this block has priority to the deprecated ` + "`" + `network_alias` + "`" + ` and ` + "`" + `network` + "`" + ` properties.`,
				},
				resource.Attribute{
					Name:        "destroy_grace_seconds",
					Description: `(Optional, int) If defined will attempt to stop the container before destroying. Container will be destroyed after ` + "`" + `n` + "`" + ` seconds or on successful stop.`,
				},
				resource.Attribute{
					Name:        "upload",
					Description: `(Optional, block) See [File Upload](#upload-1) below for details.`,
				},
				resource.Attribute{
					Name:        "ulimit",
					Description: `(Optional, block) See [Ulimits](#ulimits-1) below for details.`,
				},
				resource.Attribute{
					Name:        "pid_mode",
					Description: `(Optional, string) The PID (Process) Namespace mode for the container. Either ` + "`" + `container:<name|id>` + "`" + ` or ` + "`" + `host` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "userns_mode",
					Description: `(Optional, string) Sets the usernamespace mode for the container when usernamespace remapping option is enabled.`,
				},
				resource.Attribute{
					Name:        "healthcheck",
					Description: `(Optional, block) See [Healthcheck](#healthcheck-1) below for details.`,
				},
				resource.Attribute{
					Name:        "sysctls",
					Description: `(Optional, map) A map of kernel parameters (sysctls) to set in the container.`,
				},
				resource.Attribute{
					Name:        "ipc_mode",
					Description: `(Optional, string) IPC sharing mode for the container. Possible values are: ` + "`" + `none` + "`" + `, ` + "`" + `private` + "`" + `, ` + "`" + `shareable` + "`" + `, ` + "`" + `container:<name|id>` + "`" + ` or ` + "`" + `host` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "group_add",
					Description: `(Optional, set of strings) Add additional groups to run as. <a id="labels-1"></a> #### Labels ` + "`" + `labels` + "`" + ` is a block within the configuration that can be repeated to specify additional label name and value data to the container. Each ` + "`" + `labels` + "`" + ` block supports the following:`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `(Required, string) Name of the label`,
				},
				resource.Attribute{
					Name:        "add",
					Description: `(Optional, set of strings) list of linux capabilities to add.`,
				},
				resource.Attribute{
					Name:        "drop",
					Description: `(Optional, set of strings) list of linux capabilities to drop. Example: ` + "`" + `` + "`" + `` + "`" + `hcl resource "docker_container" "ubuntu" { name = "foo" image = "${docker_image.ubuntu.latest}" capabilities { add = ["ALL"] drop = ["SYS_ADMIN"] } } ` + "`" + `` + "`" + `` + "`" + ` <a id="mounts-1"></a> ### Mounts ` + "`" + `mounts` + "`" + ` is a block within the configuration that can be repeated to specify the extra mount mappings for the container. Each ` + "`" + `mounts` + "`" + ` block is the Specification for mounts to be added to container and supports the following:`,
				},
				resource.Attribute{
					Name:        "target",
					Description: `(Required, string) The container path.`,
				},
				resource.Attribute{
					Name:        "source",
					Description: `(Optional, string) The mount source (e.g., a volume name, a host path)`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required, string) The mount type: valid values are ` + "`" + `bind|volume|tmpfs` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "read_only",
					Description: `(Optional, string) Whether the mount should be read-only`,
				},
				resource.Attribute{
					Name:        "bind_options",
					Description: `(Optional, map) Optional configuration for the ` + "`" + `bind` + "`" + ` type.`,
				},
				resource.Attribute{
					Name:        "propagation",
					Description: `(Optional, string) A propagation mode with the value.`,
				},
				resource.Attribute{
					Name:        "volume_options",
					Description: `(Optional, map) Optional configuration for the ` + "`" + `volume` + "`" + ` type.`,
				},
				resource.Attribute{
					Name:        "no_copy",
					Description: `(Optional, string) Whether to populate volume with data from the target.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional, map of key/value pairs) Adding labels.`,
				},
				resource.Attribute{
					Name:        "driver_options",
					Description: `(Optional, map of key/value pairs) Options for the driver.`,
				},
				resource.Attribute{
					Name:        "tmpfs_options",
					Description: `(Optional, map) Optional configuration for the ` + "`" + `tmpf` + "`" + ` type.`,
				},
				resource.Attribute{
					Name:        "size_bytes",
					Description: `(Optional, int) The size for the tmpfs mount in bytes.`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `(Optional, int) The permission mode for the tmpfs mount in an integer. <a id="ports-1"></a> ### Ports ` + "`" + `ports` + "`" + ` is a block within the configuration that can be repeated to specify the port mappings of the container. Each ` + "`" + `ports` + "`" + ` block supports the following:`,
				},
				resource.Attribute{
					Name:        "internal",
					Description: `(Required, int) Port within the container.`,
				},
				resource.Attribute{
					Name:        "external",
					Description: `(Optional, int) Port exposed out of the container. If not given a free random port ` + "`" + `>= 32768` + "`" + ` will be used.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `(Optional, string) IP address/mask that can access this port, default to ` + "`" + `0.0.0.0` + "`" + ``,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Optional, string) Protocol that can be used over this port, defaults to ` + "`" + `tcp` + "`" + `. <a id="extra_hosts"></a> ### Extra Hosts ` + "`" + `host` + "`" + ` is a block within the configuration that can be repeated to specify the extra host mappings for the container. Each ` + "`" + `host` + "`" + ` block supports the following:`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `(Required, string) Hostname to add.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `(Required, string) IP address this hostname should resolve to. This is equivalent to using the ` + "`" + `--add-host` + "`" + ` option when using the ` + "`" + `run` + "`" + ` command of the Docker CLI. <a id="volumes-1"></a> ### Volumes ` + "`" + `volumes` + "`" + ` is a block within the configuration that can be repeated to specify the volumes attached to a container. Each ` + "`" + `volumes` + "`" + ` block supports the following:`,
				},
				resource.Attribute{
					Name:        "from_container",
					Description: `(Optional, string) The container where the volume is coming from.`,
				},
				resource.Attribute{
					Name:        "host_path",
					Description: `(Optional, string) The path on the host where the volume is coming from.`,
				},
				resource.Attribute{
					Name:        "volume_name",
					Description: `(Optional, string) The name of the docker volume which should be mounted.`,
				},
				resource.Attribute{
					Name:        "container_path",
					Description: `(Optional, string) The path in the container where the volume will be mounted.`,
				},
				resource.Attribute{
					Name:        "read_only",
					Description: `(Optional, bool) If true, this volume will be readonly. Defaults to false. One of ` + "`" + `from_container` + "`" + `, ` + "`" + `host_path` + "`" + ` or ` + "`" + `volume_name` + "`" + ` must be set. <a id="upload-1"></a> ### File Upload ` + "`" + `upload` + "`" + ` is a block within the configuration that can be repeated to specify files to upload to the container before starting it. Only one of ` + "`" + `content` + "`" + ` or ` + "`" + `content_base64` + "`" + ` can be set and at least one of them hast to be set. Each ` + "`" + `upload` + "`" + ` supports the following`,
				},
				resource.Attribute{
					Name:        "content",
					Description: `(Optional, string, conflicts with ` + "`" + `content_base64` + "`" + ` & ` + "`" + `source` + "`" + `) Literal string value to use as the object content, which will be uploaded as UTF-8-encoded text.`,
				},
				resource.Attribute{
					Name:        "content_base64",
					Description: `(Optional, string, conflicts with ` + "`" + `content` + "`" + ` & ` + "`" + `source` + "`" + `) Base64-encoded data that will be decoded and uploaded as raw bytes for the object content. This allows safely uploading non-UTF8 binary data, but is recommended only for larger binary content such as the result of the ` + "`" + `base64encode` + "`" + ` interpolation function. See [here](https://github.com/terraform-providers/terraform-provider-docker/issues/48#issuecomment-374174588) for the reason.`,
				},
				resource.Attribute{
					Name:        "source",
					Description: `(Optional, string, conflicts with ` + "`" + `content` + "`" + ` & ` + "`" + `content_base64` + "`" + `) A filename that references a file which will be uploaded as the object content. This allows for large file uploads that do not get stored in state.`,
				},
				resource.Attribute{
					Name:        "source_hash",
					Description: `(Optional, string) If using ` + "`" + `source` + "`" + `, this will force an update if the file content has updated but the filename has not.`,
				},
				resource.Attribute{
					Name:        "file",
					Description: `(Required, string) path to a file in the container.`,
				},
				resource.Attribute{
					Name:        "executable",
					Description: `(Optional, bool) If true, the file will be uploaded with user executable permission. Defaults to false. <a id="networks_advanced-1"></a> ### Networks advanced ` + "`" + `networks_advanced` + "`" + ` is a block within the configuration that can be repeated to specify advanced options for the container in a specific network. Each ` + "`" + `networks_advanced` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required, string) The name of the network.`,
				},
				resource.Attribute{
					Name:        "aliases",
					Description: `(Optional, set of strings) The network aliases of the container in the specific network.`,
				},
				resource.Attribute{
					Name:        "ipv4_address",
					Description: `(Optional, string) The IPV4 address of the container in the specific network.`,
				},
				resource.Attribute{
					Name:        "ipv6_address",
					Description: `(Optional, string) The IPV6 address of the container in the specific network. <a id="devices-1"></a> ### Devices ` + "`" + `devices` + "`" + ` is a block within the configuration that can be repeated to specify the devices exposed to a container. Each ` + "`" + `devices` + "`" + ` block supports the following:`,
				},
				resource.Attribute{
					Name:        "host_path",
					Description: `(Required, string) The path on the host where the device is located.`,
				},
				resource.Attribute{
					Name:        "container_path",
					Description: `(Optional, string) The path in the container where the device will be binded.`,
				},
				resource.Attribute{
					Name:        "permissions",
					Description: `(Optional, string) The cgroup permissions given to the container to access the device. Defaults to ` + "`" + `rwm` + "`" + `. <a id="ulimits-1"></a> ### Ulimits ` + "`" + `ulimit` + "`" + ` is a block within the configuration that can be repeated to specify the extra ulimits for the container. Each ` + "`" + `ulimit` + "`" + ` block supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required, string)`,
				},
				resource.Attribute{
					Name:        "soft",
					Description: `(Required, int)`,
				},
				resource.Attribute{
					Name:        "hard",
					Description: `(Required, int) <a id="healthcheck-1"></a> ### Healthcheck ` + "`" + `healthcheck` + "`" + ` is a block within the configuration that can be repeated only`,
				},
				resource.Attribute{
					Name:        "test",
					Description: `(Required, list of strings) Command to run to check health. For example, to run ` + "`" + `curl -f http://localhost/health` + "`" + ` set the command to be ` + "`" + `["CMD", "curl", "-f", "http://localhost/health"]` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "interval",
					Description: `(Optional, string) Time between running the check ` + "`" + `(ms|s|m|h)` + "`" + `. Default: ` + "`" + `0s` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `(Optional, string) Maximum time to allow one check to run ` + "`" + `(ms|s|m|h)` + "`" + `. Default: ` + "`" + `0s` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "start_period",
					Description: `(Optional, string) Start period for the container to initialize before counting retries towards unstable ` + "`" + `(ms|s|m|h)` + "`" + `. Default: ` + "`" + `0s` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "retries",
					Description: `(Optional, int) Consecutive failures needed to report unhealthy. Default: ` + "`" + `0` + "`" + `. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "exit_code",
					Description: `The exit code of the container if its execution is done (` + "`" + `must_run` + "`" + ` must be disabled).`,
				},
				resource.Attribute{
					Name:        "container_logs",
					Description: `The logs of the container if its execution is done (` + "`" + `attach` + "`" + ` must be disabled).`,
				},
				resource.Attribute{
					Name:        "network_data",
					Description: `(Map of a block) The IP addresses of the container on each network. Key are the network names, values are the IP addresses.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `The IP address of the container.`,
				},
				resource.Attribute{
					Name:        "ip_prefix_length",
					Description: `The IP prefix length of the container.`,
				},
				resource.Attribute{
					Name:        "gateway",
					Description: `The network gateway of the container.`,
				},
				resource.Attribute{
					Name:        "bridge",
					Description: `The network bridge of the container as read from its NetworkSettings.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "exit_code",
					Description: `The exit code of the container if its execution is done (` + "`" + `must_run` + "`" + ` must be disabled).`,
				},
				resource.Attribute{
					Name:        "container_logs",
					Description: `The logs of the container if its execution is done (` + "`" + `attach` + "`" + ` must be disabled).`,
				},
				resource.Attribute{
					Name:        "network_data",
					Description: `(Map of a block) The IP addresses of the container on each network. Key are the network names, values are the IP addresses.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `The IP address of the container.`,
				},
				resource.Attribute{
					Name:        "ip_prefix_length",
					Description: `The IP prefix length of the container.`,
				},
				resource.Attribute{
					Name:        "gateway",
					Description: `The network gateway of the container.`,
				},
				resource.Attribute{
					Name:        "bridge",
					Description: `The network bridge of the container as read from its NetworkSettings.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "docker_image",
			Category:         "Resources",
			ShortDescription: `Pulls a Docker image to a given Docker host.`,
			Description:      ``,
			Keywords: []string{
				"image",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required, string) The name of the Docker image, including any tags or SHA256 repo digests.`,
				},
				resource.Attribute{
					Name:        "keep_locally",
					Description: `(Optional, boolean) If true, then the Docker image won't be deleted on destroy operation. If this is false, it will delete the image from the docker local storage on destroy operation.`,
				},
				resource.Attribute{
					Name:        "pull_triggers",
					Description: `(Optional, list of strings) List of values which cause an image pull when changed. This is used to store the image digest from the registry when using the ` + "`" + `docker_registry_image` + "`" + ` [data source](/docs/providers/docker/d/registry_image.html) to trigger an image update.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "docker_network",
			Category:         "Resources",
			ShortDescription: `Manages a Docker Network.`,
			Description:      ``,
			Keywords: []string{
				"network",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required, string) The name of the Docker network.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional, block) See [Labels](#labels-1) below for details.`,
				},
				resource.Attribute{
					Name:        "check_duplicate",
					Description: `(Optional, boolean) Requests daemon to check for networks with same name.`,
				},
				resource.Attribute{
					Name:        "driver",
					Description: `(Optional, string) Name of the network driver to use. Defaults to ` + "`" + `bridge` + "`" + ` driver.`,
				},
				resource.Attribute{
					Name:        "options",
					Description: `(Optional, map of strings) Network specific options to be used by the drivers.`,
				},
				resource.Attribute{
					Name:        "internal",
					Description: `(Optional, boolean) Restrict external access to the network. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "attachable",
					Description: `(Optional, boolean) Enable manual container attachment to the network. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ingress",
					Description: `(Optional, boolean) Create swarm routing-mesh network. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ipv6",
					Description: `(Optional, boolean) Enable IPv6 networking. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ipam_driver",
					Description: `(Optional, string) Driver used by the custom IP scheme of the network.`,
				},
				resource.Attribute{
					Name:        "ipam_config",
					Description: `(Optional, block) See [IPAM config](#ipam_config-1) below for details. <a id="labels-1"></a> #### Labels ` + "`" + `labels` + "`" + ` is a block within the configuration that can be repeated to specify additional label name and value data to the container. Each ` + "`" + `labels` + "`" + ` block supports the following:`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `(Required, string) Name of the label`,
				},
				resource.Attribute{
					Name:        "subnet",
					Description: `(Optional, string)`,
				},
				resource.Attribute{
					Name:        "ip_range",
					Description: `(Optional, string)`,
				},
				resource.Attribute{
					Name:        "gateway",
					Description: `(Optional, string)`,
				},
				resource.Attribute{
					Name:        "aux_address",
					Description: `(Optional, map of string) ## Attributes Reference The following attributes are exported in addition to the above configuration:`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "docker_secret",
			Category:         "Swarm Resources",
			ShortDescription: `Manages the secrets of a Docker service in a swarm.`,
			Description:      ``,
			Keywords: []string{
				"swarm",
				"secret",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required, string) The name of the Docker secret.`,
				},
				resource.Attribute{
					Name:        "data",
					Description: `(Required, string) The base64 encoded data of the secret.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional, block) See [Labels](#labels-1) below for details. <a id="labels-1"></a> #### Labels ` + "`" + `labels` + "`" + ` is a block within the configuration that can be repeated to specify additional label name and value data to the container. Each ` + "`" + `labels` + "`" + ` block supports the following:`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `(Required, string) Name of the label`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "docker_service",
			Category:         "Swarm Resources",
			ShortDescription: `Manages the lifecycle of a Docker service.`,
			Description:      ``,
			Keywords: []string{
				"swarm",
				"service",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "auth",
					Description: `(Optional, block) See [Auth](#auth-1) below for details.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required, string) The name of the Docker service.`,
				},
				resource.Attribute{
					Name:        "task_spec",
					Description: `(Required, block) See [TaskSpec](#task-spec-1) below for details.`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `(Optional, block) See [Mode](#mode-1) below for details.`,
				},
				resource.Attribute{
					Name:        "update_config",
					Description: `(Optional, block) See [UpdateConfig](#update-rollback-config-1) below for details.`,
				},
				resource.Attribute{
					Name:        "rollback_config",
					Description: `(Optional, block) See [RollbackConfig](#update-rollback-config-1) below for details.`,
				},
				resource.Attribute{
					Name:        "endpoint_spec",
					Description: `(Optional, block) See [EndpointSpec](#endpoint-spec-1) below for details.`,
				},
				resource.Attribute{
					Name:        "converge_config",
					Description: `(Optional, block) See [Converge Config](#converge-config-1) below for details. <a id="auth-1"></a> ### Auth ` + "`" + `auth` + "`" + ` can be used additionally to the ` + "`" + `registry_auth` + "`" + `. If both properties are given the ` + "`" + `auth` + "`" + ` wins and overwrites the auth of the provider.`,
				},
				resource.Attribute{
					Name:        "server_address",
					Description: `(Required, string) The address of the registry server`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Optional, string) The username to use for authenticating to the registry. If this is blank, the ` + "`" + `DOCKER_REGISTRY_USER` + "`" + ` is also be checked.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional, string) The password to use for authenticating to the registry. If this is blank, the ` + "`" + `DOCKER_REGISTRY_PASS` + "`" + ` is also be checked. <!-- start task-spec --> <a id="task-spec-1"></a> ### TaskSpec ` + "`" + `task_spec` + "`" + ` is a block within the configuration that can be repeated only`,
				},
				resource.Attribute{
					Name:        "networks",
					Description: `(Optional, set of strings) Ids of the networks in which the container will be put in.`,
				},
				resource.Attribute{
					Name:        "log_driver",
					Description: `(Optional, block) See [Log Driver](#log-driver-1) below for details. <!-- start task-container-spec --> <a id="container-spec-1"></a> #### ContainerSpec ` + "`" + `container_spec` + "`" + ` is a block within the configuration that can be repeated only`,
				},
				resource.Attribute{
					Name:        "image",
					Description: `(Required, string) The image used to create the Docker service.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional, block) See [Labels](#labels-1) below for details.`,
				},
				resource.Attribute{
					Name:        "command",
					Description: `(Optional, list of strings) The command to be run in the image.`,
				},
				resource.Attribute{
					Name:        "args",
					Description: `(Optional, list of strings) Arguments to the command.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `(Optional, string) The hostname to use for the container, as a valid RFC 1123 hostname.`,
				},
				resource.Attribute{
					Name:        "env",
					Description: `(Optional, map of string/string) A list of environment variables in the form VAR=value.`,
				},
				resource.Attribute{
					Name:        "dir",
					Description: `(Optional, string) The working directory for commands to run in.`,
				},
				resource.Attribute{
					Name:        "user",
					Description: `(Optional, string) The user inside the container.`,
				},
				resource.Attribute{
					Name:        "groups",
					Description: `(Optional, list of strings) A list of additional groups that the container process will run as.`,
				},
				resource.Attribute{
					Name:        "read_only",
					Description: `(Optional, bool) Mount the container's root filesystem as read only.`,
				},
				resource.Attribute{
					Name:        "mounts",
					Description: `(Optional, set of blocks) See [Mounts](#mounts-1) below for details.`,
				},
				resource.Attribute{
					Name:        "stop_signal",
					Description: `(Optional, string) Signal to stop the container.`,
				},
				resource.Attribute{
					Name:        "stop_grace_period",
					Description: `(Optional, string) Amount of time to wait for the container to terminate before forcefully removing it ` + "`" + `(ms|s|m|h)` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "healthcheck",
					Description: `(Optional, block) See [Healthcheck](#healthcheck-1) below for details.`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `(Optional, map of string/string) A list of hostname/IP mappings to add to the container's hosts file.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `(Required string) The ip`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `(Required string) The hostname`,
				},
				resource.Attribute{
					Name:        "dns_config",
					Description: `(Optional, block) See [DNS Config](#dnsconfig-1) below for details.`,
				},
				resource.Attribute{
					Name:        "secrets",
					Description: `(Optional, set of blocks) See [Secrets](#secrets-1) below for details.`,
				},
				resource.Attribute{
					Name:        "configs",
					Description: `(Optional, set of blocks) See [Configs](#configs-1) below for details.`,
				},
				resource.Attribute{
					Name:        "isolation",
					Description: `(Optional, string) Isolation technology of the containers running the service. (Windows only). Valid values are: ` + "`" + `default|process|hyperv` + "`" + ` <a id="labels-1"></a> #### Labels ` + "`" + `labels` + "`" + ` is a block within the configuration that can be repeated to specify additional label name and value data to the container. Each ` + "`" + `labels` + "`" + ` block supports the following:`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `(Required, string) Name of the label`,
				},
				resource.Attribute{
					Name:        "credential_spec",
					Description: `(Optional, block) For managed service account (Windows only)`,
				},
				resource.Attribute{
					Name:        "file",
					Description: `(Optional, string) Load credential spec from this file.`,
				},
				resource.Attribute{
					Name:        "registry",
					Description: `(Optional, string) Load credential spec from this value in the Windows registry.`,
				},
				resource.Attribute{
					Name:        "se_linux_context",
					Description: `(Optional, block) SELinux labels of the container`,
				},
				resource.Attribute{
					Name:        "disable",
					Description: `(Optional, bool) Disable SELinux`,
				},
				resource.Attribute{
					Name:        "user",
					Description: `(Optional, string) SELinux user label`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `(Optional, string) SELinux role label`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional, string) SELinux type label`,
				},
				resource.Attribute{
					Name:        "level",
					Description: `(Optional, string) SELinux level label <a id="mounts-1"></a> #### Mounts ` + "`" + `mounts` + "`" + ` is a block within the configuration that can be repeated to specify the extra mount mappings for the container. Each ` + "`" + `mounts` + "`" + ` block is the Specification for mounts to be added to containers created as part of the service and supports the following:`,
				},
				resource.Attribute{
					Name:        "target",
					Description: `(Required, string) The container path.`,
				},
				resource.Attribute{
					Name:        "source",
					Description: `(Optional, string) The mount source (e.g., a volume name, a host path)`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required, string) The mount type: valid values are ` + "`" + `bind|volume|tmpfs` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "read_only",
					Description: `(Optional, string) Whether the mount should be read-only`,
				},
				resource.Attribute{
					Name:        "bind_options",
					Description: `(Optional, map) Optional configuration for the ` + "`" + `bind` + "`" + ` type.`,
				},
				resource.Attribute{
					Name:        "propagation",
					Description: `(Optional, string) A propagation mode with the value.`,
				},
				resource.Attribute{
					Name:        "volume_options",
					Description: `(Optional, map) Optional configuration for the ` + "`" + `volume` + "`" + ` type.`,
				},
				resource.Attribute{
					Name:        "no_copy",
					Description: `(Optional, string) Whether to populate volume with data from the target.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional, block) See [Labels](#labels-1) above for details.`,
				},
				resource.Attribute{
					Name:        "driver_config",
					Description: `(Optional, map) The name of the driver to create the volume.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional, string) The name of the driver to create the volume.`,
				},
				resource.Attribute{
					Name:        "options",
					Description: `(Optional, map of key/value pairs) Options for the driver.`,
				},
				resource.Attribute{
					Name:        "tmpfs_options",
					Description: `(Optional, map) Optional configuration for the ` + "`" + `tmpf` + "`" + ` type.`,
				},
				resource.Attribute{
					Name:        "size_bytes",
					Description: `(Optional, int) The size for the tmpfs mount in bytes.`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `(Optional, int) The permission mode for the tmpfs mount in an integer. <a id="healthcheck-1"></a> #### Healthcheck ` + "`" + `healthcheck` + "`" + ` is a block within the configuration that can be repeated only`,
				},
				resource.Attribute{
					Name:        "test",
					Description: `(Required, list of strings) Command to run to check health. For example, to run ` + "`" + `curl -f http://localhost/health` + "`" + ` set the command to be ` + "`" + `["CMD", "curl", "-f", "http://localhost/health"]` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "interval",
					Description: `(Optional, string) Time between running the check ` + "`" + `(ms|s|m|h)` + "`" + `. Default: ` + "`" + `0s` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `(Optional, string) Maximum time to allow one check to run ` + "`" + `(ms|s|m|h)` + "`" + `. Default: ` + "`" + `0s` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "start_period",
					Description: `(Optional, string) Start period for the container to initialize before counting retries towards unstable ` + "`" + `(ms|s|m|h)` + "`" + `. Default: ` + "`" + `0s` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "retries",
					Description: `(Optional, int) Consecutive failures needed to report unhealthy. Default: ` + "`" + `0` + "`" + `. <a id="dnsconfig-1"></a> ### DNS Config ` + "`" + `dns_config` + "`" + ` is a block within the configuration that can be repeated only`,
				},
				resource.Attribute{
					Name:        "nameservers",
					Description: `(Required, list of strings) The IP addresses of the name servers, for example, ` + "`" + `8.8.8.8` + "`" + ``,
				},
				resource.Attribute{
					Name:        "search",
					Description: `(Optional, list of strings)A search list for host-name lookup.`,
				},
				resource.Attribute{
					Name:        "options",
					Description: `(Optional, list of strings) A list of internal resolver variables to be modified, for example, ` + "`" + `debug` + "`" + `, ` + "`" + `ndots:3` + "`" + ` <a id="secrets-1"></a> ### Secrets ` + "`" + `secrets` + "`" + ` is a block within the configuration that can be repeated to specify the extra mount mappings for the container. Each ` + "`" + `secrets` + "`" + ` block is a reference to a secret that will be exposed to the service and supports the following:`,
				},
				resource.Attribute{
					Name:        "secret_id",
					Description: `(Required, string) ConfigID represents the ID of the specific secret.`,
				},
				resource.Attribute{
					Name:        "secret_name",
					Description: `(Optional, string) The name of the secret that this references, but internally it is just provided for lookup/display purposes`,
				},
				resource.Attribute{
					Name:        "file_name",
					Description: `(Required, string) Represents the final filename in the filesystem. The specific target file that the secret data is written within the docker container, e.g. ` + "`" + `/root/secret/secret.json` + "`" + ``,
				},
				resource.Attribute{
					Name:        "file_uid",
					Description: `(Optional, string) Represents the file UID. Defaults: ` + "`" + `0` + "`" + ``,
				},
				resource.Attribute{
					Name:        "file_gid",
					Description: `(Optional, string) Represents the file GID. Defaults: ` + "`" + `0` + "`" + ``,
				},
				resource.Attribute{
					Name:        "file_mode",
					Description: `(Optional, int) Represents the FileMode of the file. Defaults: ` + "`" + `0444` + "`" + ` <a id="configs-1"></a> ### Configs ` + "`" + `configs` + "`" + ` is a block within the configuration that can be repeated to specify the extra mount mappings for the container. Each ` + "`" + `configs` + "`" + ` is a reference to a secret that is exposed to the service and supports the following:`,
				},
				resource.Attribute{
					Name:        "config_id",
					Description: `(Required, string) ConfigID represents the ID of the specific config.`,
				},
				resource.Attribute{
					Name:        "config_name",
					Description: `(Optional, string) The name of the config that this references, but internally it is just provided for lookup/display purposes`,
				},
				resource.Attribute{
					Name:        "file_name",
					Description: `(Required, string) Represents the final filename in the filesystem. The specific target file that the config data is written within the docker container, e.g. ` + "`" + `/root/config/config.json` + "`" + ``,
				},
				resource.Attribute{
					Name:        "file_uid",
					Description: `(Optional, string) Represents the file UID. Defaults: ` + "`" + `0` + "`" + ``,
				},
				resource.Attribute{
					Name:        "file_gid",
					Description: `(Optional, string) Represents the file GID. Defaults: ` + "`" + `0` + "`" + ``,
				},
				resource.Attribute{
					Name:        "file_mode",
					Description: `(Optional, int) Represents the FileMode of the file. Defaults: ` + "`" + `0444` + "`" + ` <!-- end task-container-spec --> <!-- start task-resources-spec --> <a id="resources-1"></a> #### Resources ` + "`" + `resources` + "`" + ` is a block within the configuration that can be repeated only`,
				},
				resource.Attribute{
					Name:        "limits",
					Description: `(Optional, list of strings) Describes the resources which can be advertised by a node and requested by a task.`,
				},
				resource.Attribute{
					Name:        "reservation",
					Description: `(Optional, list of strings) An object describing the resources which can be advertised by a node and requested by a task.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required, string) The logging driver to use. Either ` + "`" + `(none|json-file|syslog|journald|gelf|fluentd|awslogs|splunk|etwlogs|gcplogs)` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "options",
					Description: `(Optional, a map of strings and strings) The options for the logging driver, e.g. ` + "`" + `` + "`" + `` + "`" + `hcl options { awslogs-region = "us-west-2" awslogs-group = "dev/foo-service" } ` + "`" + `` + "`" + `` + "`" + ` <!-- end log-driver-spec --> <!-- end task-spec --> <a id="mode-1"></a> ### Mode ` + "`" + `mode` + "`" + ` is a block within the configuration that can be repeated only`,
				},
				resource.Attribute{
					Name:        "global",
					Description: `(Optional, bool) set it to ` + "`" + `true` + "`" + ` to run the service in the global mode ` + "`" + `` + "`" + `` + "`" + `hcl resource "docker_service" "foo" { ... mode { global = true } ... } ` + "`" + `` + "`" + `` + "`" + ``,
				},
				resource.Attribute{
					Name:        "replicated",
					Description: `(Optional, map), which contains atm only the amount of ` + "`" + `replicas` + "`" + ` ` + "`" + `` + "`" + `` + "`" + `hcl resource "docker_service" "foo" { ... mode { replicated { replicas = 2 } } ... } ` + "`" + `` + "`" + `` + "`" + ` ~>`,
				},
				resource.Attribute{
					Name:        "parallelism",
					Description: `(Optional, int) The maximum number of tasks to be updated in one iteration simultaneously (0 to update all at once).`,
				},
				resource.Attribute{
					Name:        "delay",
					Description: `(Optional, int) Delay between updates ` + "`" + `(ns|us|ms|s|m|h)` + "`" + `, e.g. ` + "`" + `5s` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "failure_action",
					Description: `(Optional, int) Action on update failure: ` + "`" + `pause|continue|rollback` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "monitor",
					Description: `(Optional, int) Duration after each task update to monitor for failure ` + "`" + `(ns|us|ms|s|m|h)` + "`" + ``,
				},
				resource.Attribute{
					Name:        "max_failure_ratio",
					Description: `(Optional, string) The failure rate to tolerate during an update as ` + "`" + `float` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "order",
					Description: `(Optional, int) Update order either 'stop-first' or 'start-first'. <a id="endpoint-spec-1"></a> ### EndpointSpec ` + "`" + `endpoint_spec` + "`" + ` is a block within the configuration that can be repeated only`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `(Optional, string) The mode of resolution to use for internal load balancing between tasks. ` + "`" + `(vip|dnsrr)` + "`" + `. Default: ` + "`" + `vip` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ports",
					Description: `(Optional, block) See [Ports](#ports-1) below for details. <a id="ports-1"></a> #### Ports ` + "`" + `ports` + "`" + ` is a block within the configuration that can be repeated to specify the port mappings of the container. Each ` + "`" + `ports` + "`" + ` block supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional, string) A random name for the port.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Optional, string) Protocol that can be used over this port: ` + "`" + `tcp|udp|sctp` + "`" + `. Default: ` + "`" + `tcp` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "target_port",
					Description: `(Required, int) Port inside the container.`,
				},
				resource.Attribute{
					Name:        "published_port",
					Description: `(Required, int) The port on the swarm hosts. If not set the value of ` + "`" + `target_port` + "`" + ` will be used.`,
				},
				resource.Attribute{
					Name:        "publish_mode",
					Description: `(Optional, string) Represents the mode in which the port is to be published: ` + "`" + `ingress|host` + "`" + ` <a id="converge-config-1"></a> ### Converge Config ` + "`" + `converge_config` + "`" + ` is a block within the configuration that can be repeated only`,
				},
				resource.Attribute{
					Name:        "delay",
					Description: `(Optional, string) Time between each the check to check docker endpoint ` + "`" + `(ms|s|m|h)` + "`" + `. For example, to check if all tasks are up when a service is created, or to check if all tasks are successfully updated on an update. Default: ` + "`" + `7s` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `(Optional, string) The timeout of the service to reach the desired state ` + "`" + `(s|m)` + "`" + `. Default: ` + "`" + `3m` + "`" + `. ## Attributes Reference The following attributes are exported in addition to the above configuration:`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "docker_volume",
			Category:         "Resources",
			ShortDescription: `Creates and destroys docker volumes.`,
			Description:      ``,
			Keywords: []string{
				"volume",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional, string) The name of the Docker volume (generated if not provided).`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional, map of string/string key/value pairs) User-defined key/value metadata.`,
				},
				resource.Attribute{
					Name:        "driver",
					Description: `(Optional, string) Driver type for the volume (defaults to local).`,
				},
				resource.Attribute{
					Name:        "driver_opts",
					Description: `(Optional, map of strings) Options specific to the driver. ## Attributes Reference The following attributes are exported in addition to the above configuration:`,
				},
			},
			Attributes: []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"docker_config":    0,
		"docker_container": 1,
		"docker_image":     2,
		"docker_network":   3,
		"docker_secret":    4,
		"docker_service":   5,
		"docker_volume":    6,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
