package cobbler

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "cobbler_distro",
			Category:         "Resources",
			ShortDescription: `Manages a distribution within Cobbler.`,
			Description:      ``,
			Keywords: []string{
				"distro",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "arch",
					Description: `(Required) The architecture of the distro. Valid options are: i386, x86_64, ia64, ppc, ppc64, s390, arm.`,
				},
				resource.Attribute{
					Name:        "breed",
					Description: `(Required) The "breed" of distribution. Valid options are: redhat, fedora, centos, scientific linux, suse, debian, and ubuntu. These choices may vary depending on the version of Cobbler in use.`,
				},
				resource.Attribute{
					Name:        "boot_files",
					Description: `(Optional) Files copied into tftpboot beyond the kernel/initrd.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `(Optional) Free form text description.`,
				},
				resource.Attribute{
					Name:        "fetchable_files",
					Description: `(Optional) Templates for tftp or wget.`,
				},
				resource.Attribute{
					Name:        "kernel",
					Description: `(Required) Absolute path to kernel on filesystem. This must already exist prior to creating the distro.`,
				},
				resource.Attribute{
					Name:        "kernel_options",
					Description: `(Optional) Kernel options to use with the kernel.`,
				},
				resource.Attribute{
					Name:        "kernel_options_post",
					Description: `(Optional) Post install Kernel options to use with the kernel after installation.`,
				},
				resource.Attribute{
					Name:        "initrd",
					Description: `(Required) Absolute path to initrd on filesystem. This must already exist prior to creating the distro.`,
				},
				resource.Attribute{
					Name:        "mgmt_classes",
					Description: `(Optional) Management classes for external config management.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A name for the distro.`,
				},
				resource.Attribute{
					Name:        "os_version",
					Description: `(Required) The version of the distro you are creating. This varies with the version of Cobbler you are using. An updated signature list may need to be obtained in order to support a newer version. Example: ` + "`" + `trusty` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "owners",
					Description: `(Optional) Owners list for authz_ownership.`,
				},
				resource.Attribute{
					Name:        "redhat_management_key",
					Description: `(Optional) Red Hat Management key.`,
				},
				resource.Attribute{
					Name:        "redhat_management_server",
					Description: `(Optional) Red Hat Management server.`,
				},
				resource.Attribute{
					Name:        "template_files",
					Description: `(Optional) File mappings for built-in config management. ## Attributes Reference All of the above Optional attributes are also exported. ## Notes The path to the ` + "`" + `kernel` + "`" + ` and ` + "`" + `initrd` + "`" + ` files must exist before creating a Distro. Usually this involves running ` + "`" + `cobbler import ...` + "`" + ` prior to creating the Distro.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cobbler_kickstart_file",
			Category:         "Resources",
			ShortDescription: `Manages a Kickstart File within Cobbler.`,
			Description:      ``,
			Keywords: []string{
				"kickstart",
				"file",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "body",
					Description: `(Required) The body of the kickstart file.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the kickstart file. This must be the full path, including ` + "`" + `/var/lib/cobbler/kickstarts` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cobbler_profile",
			Category:         "Resources",
			ShortDescription: `Manages a Profile within Cobbler.`,
			Description:      ``,
			Keywords: []string{
				"profile",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "boot_files",
					Description: `(Optional) Files copied into tftpboot beyond the kernel/initrd.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `(Optional) Free form text description.`,
				},
				resource.Attribute{
					Name:        "parent",
					Description: `(Optional) The parent this profile inherits settings from.`,
				},
				resource.Attribute{
					Name:        "server",
					Description: `(Optional) The server-override for the profile.`,
				},
				resource.Attribute{
					Name:        "distro",
					Description: `(Optional) Parent distribution.`,
				},
				resource.Attribute{
					Name:        "enable_gpxe",
					Description: `(Optional) Use gPXE instead of PXELINUX for advanced booting options.`,
				},
				resource.Attribute{
					Name:        "enable_menu",
					Description: `(Optional) Enable a boot menu.`,
				},
				resource.Attribute{
					Name:        "fetchable_files",
					Description: `(Optional) Templates for tftp or wget.`,
				},
				resource.Attribute{
					Name:        "kernel_options",
					Description: `(Optional) Kernel options for the profile.`,
				},
				resource.Attribute{
					Name:        "kernel_options_post",
					Description: `(Optional) Post install kernel options.`,
				},
				resource.Attribute{
					Name:        "kickstart",
					Description: `(Optional) The kickstart file to use.`,
				},
				resource.Attribute{
					Name:        "ks_meta",
					Description: `(Optional) Kickstart metadata.`,
				},
				resource.Attribute{
					Name:        "mgmt_classes",
					Description: `(Optional) For external configuration management.`,
				},
				resource.Attribute{
					Name:        "mgmt_parameters",
					Description: `(Optional) Parameters which will be handed to your management application (Must be a valid YAML dictionary).`,
				},
				resource.Attribute{
					Name:        "name_servers_search",
					Description: `(Optional) Name server search settings.`,
				},
				resource.Attribute{
					Name:        "name_servers",
					Description: `(Optional) Name servers.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the profile.`,
				},
				resource.Attribute{
					Name:        "owners",
					Description: `(Optional) Owners list for authz_ownership.`,
				},
				resource.Attribute{
					Name:        "proxy",
					Description: `(Optional) Proxy URL.`,
				},
				resource.Attribute{
					Name:        "redhat_management_key",
					Description: `(Optional) Red Hat Management Key.`,
				},
				resource.Attribute{
					Name:        "redhat_management_server",
					Description: `(Optional) RedHat Management Server.`,
				},
				resource.Attribute{
					Name:        "repos",
					Description: `(Optional) Repos to auto-assign to this profile.`,
				},
				resource.Attribute{
					Name:        "template_files",
					Description: `(Optional) File mappings for built-in config management.`,
				},
				resource.Attribute{
					Name:        "template_remote_kickstarts",
					Description: `(Optional) remote kickstart templates.`,
				},
				resource.Attribute{
					Name:        "virt_auto_boot",
					Description: `(Optional) Auto boot virtual machines.`,
				},
				resource.Attribute{
					Name:        "virt_bridge",
					Description: `(Optional) The bridge for virtual machines.`,
				},
				resource.Attribute{
					Name:        "virt_cpus",
					Description: `(Optional) The number of virtual CPUs.`,
				},
				resource.Attribute{
					Name:        "virt_file_size",
					Description: `(Optional) The virtual machine file size.`,
				},
				resource.Attribute{
					Name:        "virt_path",
					Description: `(Optional) The virtual machine path.`,
				},
				resource.Attribute{
					Name:        "virt_ram",
					Description: `(Optional) The amount of RAM for the virtual machine.`,
				},
				resource.Attribute{
					Name:        "virt_type",
					Description: `(Optional) The type of virtual machine. Valid options are: xenpv, xenfv, qemu, kvm, vmware, openvz.`,
				},
				resource.Attribute{
					Name:        "virt_disk_driver",
					Description: `(Optional) The virtual machine disk driver. ## Attributes Reference All of the above Optional attributes are also exported.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cobbler_repo",
			Category:         "Resources",
			ShortDescription: `Manages a repo within Cobbler.`,
			Description:      ``,
			Keywords: []string{
				"repo",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "apt_components",
					Description: `(Optional) List of Apt components such as main, restricted, universe. Applicable to apt breeds only.`,
				},
				resource.Attribute{
					Name:        "apt_dists",
					Description: `(Optional) List of Apt distribution names such as trusty, trusty-updates. Applicable to apt breeds only.`,
				},
				resource.Attribute{
					Name:        "arch",
					Description: `(Optional) The architecture of the repo. Valid options are: i386, x86_64, ia64, ppc, ppc64, s390, arm.`,
				},
				resource.Attribute{
					Name:        "breed",
					Description: `(Required) The "breed" of distribution. Valid options are: rsync, rhn, yum, apt, and wget. These choices may vary depending on the version of Cobbler in use.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `(Optional) Free form text description.`,
				},
				resource.Attribute{
					Name:        "createrepo_flags",
					Description: `(Optional) Flags to use with ` + "`" + `createrepo` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "environment",
					Description: `(Optional) Environment variables to use during repo command execution.`,
				},
				resource.Attribute{
					Name:        "keep_updated",
					Description: `(Optional) Update the repo upon Cobbler sync. Valid values are true or false.`,
				},
				resource.Attribute{
					Name:        "mirror",
					Description: `(Required) Address of the repo to mirror.`,
				},
				resource.Attribute{
					Name:        "mirror_locally",
					Description: `(Required) Whether to copy the files locally or just references to the external files. Valid values are true or false.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A name for the repo.`,
				},
				resource.Attribute{
					Name:        "owners",
					Description: `(Optional) List of Owners for authz_ownership.`,
				},
				resource.Attribute{
					Name:        "proxy",
					Description: `(Optional) Proxy to use for downloading the repo. This argument does not work on older versions of Cobbler.`,
				},
				resource.Attribute{
					Name:        "rpm_list",
					Description: `(Optional) List of specific RPMs to mirror. ## Attributes Reference All of the above Optional attributes are also exported.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cobbler_snippet",
			Category:         "Resources",
			ShortDescription: `Manages a Snippet within Cobbler.`,
			Description:      ``,
			Keywords: []string{
				"snippet",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "body",
					Description: `(Required) The body of the snippet.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the snippet. This must be the full path, including ` + "`" + `/var/lib/cobbler/snippets` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cobbler_system",
			Category:         "Resources",
			ShortDescription: `Manages a System within Cobbler.`,
			Description:      ``,
			Keywords: []string{
				"system",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "boot_files",
					Description: `(Optional) TFTP boot files copied into tftpboot.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `(Optional) Free form text description`,
				},
				resource.Attribute{
					Name:        "enable_gpxe",
					Description: `(Optional) Use gPXE instead of PXELINUX.`,
				},
				resource.Attribute{
					Name:        "fetchable_files",
					Description: `(Optional) Templates for tftp or wget.`,
				},
				resource.Attribute{
					Name:        "gateway",
					Description: `(Optional) Network gateway.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `(Optional) Hostname of the system.`,
				},
				resource.Attribute{
					Name:        "image",
					Description: `(Optional) Parent image (if no profile is used).`,
				},
				resource.Attribute{
					Name:        "interface",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "ipv6_default_device",
					Description: `(Optional) IPv6 default device.`,
				},
				resource.Attribute{
					Name:        "kernel_options",
					Description: `(Optional) Kernel options. ex: selinux=permissive.`,
				},
				resource.Attribute{
					Name:        "kernel_options_post",
					Description: `(Optional) Kernel options (post install).`,
				},
				resource.Attribute{
					Name:        "kickstart",
					Description: `(Optional) Path to kickstart template.`,
				},
				resource.Attribute{
					Name:        "ks_meta",
					Description: `(Optional) Kickstart metadata.`,
				},
				resource.Attribute{
					Name:        "ldap_enabled",
					Description: `(Optional) Configure LDAP at next config update.`,
				},
				resource.Attribute{
					Name:        "ldap_type",
					Description: `(Optional) LDAP management type.`,
				},
				resource.Attribute{
					Name:        "mgmt_classes",
					Description: `(Optional) Management classes for external config management.`,
				},
				resource.Attribute{
					Name:        "mgmt_parameters",
					Description: `(Optional) Parameters which will be handed to your management application. Must be a valid YAML dictionary.`,
				},
				resource.Attribute{
					Name:        "monit_enabled",
					Description: `(Optional) Configure monit on this machine at next config update.`,
				},
				resource.Attribute{
					Name:        "name_servers_search",
					Description: `(Optional) Name servers search path.`,
				},
				resource.Attribute{
					Name:        "name_servers",
					Description: `(Optional) Name servers.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the system.`,
				},
				resource.Attribute{
					Name:        "netboot_enabled",
					Description: `(Optional) (re)Install this machine at next boot.`,
				},
				resource.Attribute{
					Name:        "owners",
					Description: `(Optional) Owners list for authz_ownership.`,
				},
				resource.Attribute{
					Name:        "power_address",
					Description: `(Optional) Power management address.`,
				},
				resource.Attribute{
					Name:        "power_id",
					Description: `(Optional) Usually a plug number or blade name if power type requires it.`,
				},
				resource.Attribute{
					Name:        "power_pass",
					Description: `(Optional) Power management password.`,
				},
				resource.Attribute{
					Name:        "power_type",
					Description: `(Optional) Power management type.`,
				},
				resource.Attribute{
					Name:        "power_user",
					Description: `(Optional) Power management user.`,
				},
				resource.Attribute{
					Name:        "profile",
					Description: `(Required) Parent profile.`,
				},
				resource.Attribute{
					Name:        "proxy",
					Description: `(Optional) Proxy URL.`,
				},
				resource.Attribute{
					Name:        "redhat_management_key",
					Description: `(Optional) Red Hat management key.`,
				},
				resource.Attribute{
					Name:        "redhat_management_server",
					Description: `(Optional) Red Hat management server.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Optional) System status (development, testing, acceptance, production).`,
				},
				resource.Attribute{
					Name:        "template_files",
					Description: `(Optional) File mappings for built-in configuration management.`,
				},
				resource.Attribute{
					Name:        "template_remote_kickstarts",
					Description: `(Optional) template remote kickstarts.`,
				},
				resource.Attribute{
					Name:        "virt_auto_boot",
					Description: `(Optional) Auto boot the VM.`,
				},
				resource.Attribute{
					Name:        "virt_cpus",
					Description: `(Optional) Number of virtual CPUs in the VM.`,
				},
				resource.Attribute{
					Name:        "virt_disk_driver",
					Description: `(Optional) The on-disk format for the virtualization disk.`,
				},
				resource.Attribute{
					Name:        "virt_file_size",
					Description: `(Optional) Virt file size.`,
				},
				resource.Attribute{
					Name:        "virt_path",
					Description: `(Optional) Path to the VM.`,
				},
				resource.Attribute{
					Name:        "virt_pxe_boot",
					Description: `(Optional) Use PXE to build this VM?`,
				},
				resource.Attribute{
					Name:        "virt_ram",
					Description: `(Optional) The amount of RAM for the VM.`,
				},
				resource.Attribute{
					Name:        "virt_type",
					Description: `(Optional) Virtualization technology to use: xenpv, xenfv, qemu, kvm, vmware, openvz. The ` + "`" + `interface` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The device name of the interface. ex: eth0.`,
				},
				resource.Attribute{
					Name:        "cnames",
					Description: `(Optional) Canonical name records.`,
				},
				resource.Attribute{
					Name:        "dhcp_tag",
					Description: `(Optional) DHCP tag.`,
				},
				resource.Attribute{
					Name:        "dns_name",
					Description: `(Optional) DNS name.`,
				},
				resource.Attribute{
					Name:        "bonding_opts",
					Description: `(Optional) Options for bonded interfaces.`,
				},
				resource.Attribute{
					Name:        "bridge_opts",
					Description: `(Optional) Options for bridge interfaces.`,
				},
				resource.Attribute{
					Name:        "gateway",
					Description: `(Optional) Per-interface gateway.`,
				},
				resource.Attribute{
					Name:        "interface_type",
					Description: `(Optional) The type of interface: na, master, slave, bond, bond_slave, bridge, bridge_slave, bonded_bridge_slave.`,
				},
				resource.Attribute{
					Name:        "interface_master",
					Description: `(Optional) The master interface when slave.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `(Optional) The IP address of the interface.`,
				},
				resource.Attribute{
					Name:        "ipv6_address",
					Description: `(Optional) The IPv6 address of the interface.`,
				},
				resource.Attribute{
					Name:        "ipv6_mtu",
					Description: `(Optional) The MTU of the IPv6 address.`,
				},
				resource.Attribute{
					Name:        "ipv6_static_routes",
					Description: `(Optional) Static routes for the IPv6 interface.`,
				},
				resource.Attribute{
					Name:        "ipv6_default_gateway",
					Description: `(Optional) The default gateawy for the IPv6 address / interface.`,
				},
				resource.Attribute{
					Name:        "mac_address",
					Description: `(Optional) The MAC address of the interface.`,
				},
				resource.Attribute{
					Name:        "management",
					Description: `(Optional) Whether this interface is a management interface.`,
				},
				resource.Attribute{
					Name:        "netmask",
					Description: `(Optional) The IPv4 netmask of the interface.`,
				},
				resource.Attribute{
					Name:        "static",
					Description: `(Optional) Whether the interface should be static or DHCP.`,
				},
				resource.Attribute{
					Name:        "static_routes",
					Description: `(Optional) Static routes for the interface.`,
				},
				resource.Attribute{
					Name:        "virt_bridge",
					Description: `(Optional) The virtual bridge to attach to. ## Attribute Reference All optional attributes listed above are also exported.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"cobbler_distro":         0,
		"cobbler_kickstart_file": 1,
		"cobbler_profile":        2,
		"cobbler_repo":           3,
		"cobbler_snippet":        4,
		"cobbler_system":         5,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
