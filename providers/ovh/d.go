package ovh

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "ovh_cloud_project_region",
			Category:         "Data Sources",
			ShortDescription: `Get information & status of a region associated with a public cloud project.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Optional) Deprecated. The id of the public cloud project. If omitted, the ` + "`" + `OVH_PROJECT_ID` + "`" + ` environment variable is used. One of ` + "`" + `service_name` + "`" + ` or ` + "`" + `project_id` + "`" + ` is required. Conflits with ` + "`" + `service_name` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `(Optional) The id of the public cloud project. If omitted, the ` + "`" + `OVH_CLOUD_PROJECT_SERVICE` + "`" + ` environment variable is used. One of ` + "`" + `service_name` + "`" + ` or ` + "`" + `project_id` + "`" + ` is required. Conflits with ` + "`" + `project_id` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the region associated with the public cloud project. ## Attributes Reference ` + "`" + `id` + "`" + ` is set to the ID of the project concatenated with the name of the region. In addition, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "continent_code",
					Description: `the code of the geographic continent the region is running. E.g.: EU for Europe, US for America...`,
				},
				resource.Attribute{
					Name:        "datacenter_location",
					Description: `The location code of the datacenter. E.g.: "GRA", meaning Gravelines, for region "GRA1"`,
				},
				resource.Attribute{
					Name:        "continentCode",
					Description: `(Deprecated) Use ` + "`" + `continent_code` + "`" + ` instead.`,
				},
				resource.Attribute{
					Name:        "datacenterLocation",
					Description: `(Deprecated) Use ` + "`" + `datacenter_location` + "`" + ` instead.`,
				},
				resource.Attribute{
					Name:        "services",
					Description: `The list of public cloud services running within the region`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `the name of the public cloud service`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `the status of the service`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "continent_code",
					Description: `the code of the geographic continent the region is running. E.g.: EU for Europe, US for America...`,
				},
				resource.Attribute{
					Name:        "datacenter_location",
					Description: `The location code of the datacenter. E.g.: "GRA", meaning Gravelines, for region "GRA1"`,
				},
				resource.Attribute{
					Name:        "continentCode",
					Description: `(Deprecated) Use ` + "`" + `continent_code` + "`" + ` instead.`,
				},
				resource.Attribute{
					Name:        "datacenterLocation",
					Description: `(Deprecated) Use ` + "`" + `datacenter_location` + "`" + ` instead.`,
				},
				resource.Attribute{
					Name:        "services",
					Description: `The list of public cloud services running within the region`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `the name of the public cloud service`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `the status of the service`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_cloud_project_regions",
			Category:         "Data Sources",
			ShortDescription: `Get the list of regions associated with a public cloud project.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Optional) Deprecated. The id of the public cloud project. If omitted, the ` + "`" + `OVH_PROJECT_ID` + "`" + ` environment variable is used. One of ` + "`" + `service_name` + "`" + ` or ` + "`" + `project_id` + "`" + ` is required. Conflits with ` + "`" + `service_name` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `(Optional) The id of the public cloud project. If omitted, the ` + "`" + `OVH_CLOUD_PROJECT_SERVICE` + "`" + ` environment variable is used. One of ` + "`" + `service_name` + "`" + ` or ` + "`" + `project_id` + "`" + ` is required. Conflits with ` + "`" + `project_id` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "has_services_up",
					Description: `(Optional) List of services which has to be UP in regions. Example: "image", "instance", "network", "storage", "volume", "workflow", ... If left blank, returns all regions associated with the project_id. ## Attributes Reference ` + "`" + `id` + "`" + ` is set to the ID of the project. In addition, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `The list of regions associated with the project, filtered by services UP.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "names",
					Description: `The list of regions associated with the project, filtered by services UP.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_cloud_region (deprecated)",
			Category:         "Data Sources",
			ShortDescription: `Get information & status of a region associated with a public cloud project.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The id of the public cloud project. If omitted, the ` + "`" + `OVH_PROJECT_ID` + "`" + ` environment variable is used.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Required) The name of the region associated with the public cloud project. ## Attributes Reference ` + "`" + `id` + "`" + ` is set to the ID of the project concatenated with the name of the region. In addition, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "continent_code",
					Description: `the code of the geographic continent the region is running. E.g.: EU for Europe, US for America...`,
				},
				resource.Attribute{
					Name:        "datacenter_location",
					Description: `The location code of the datacenter. E.g.: "GRA", meaning Gravelines, for region "GRA1"`,
				},
				resource.Attribute{
					Name:        "continentCode",
					Description: `(Deprecated) Use ` + "`" + `continent_code` + "`" + ` instead.`,
				},
				resource.Attribute{
					Name:        "datacenterLocation",
					Description: `(Deprecated) Use ` + "`" + `datacenter_location` + "`" + ` instead.`,
				},
				resource.Attribute{
					Name:        "services",
					Description: `The list of public cloud services running within the region`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `the name of the public cloud service`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `the status of the service`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "continent_code",
					Description: `the code of the geographic continent the region is running. E.g.: EU for Europe, US for America...`,
				},
				resource.Attribute{
					Name:        "datacenter_location",
					Description: `The location code of the datacenter. E.g.: "GRA", meaning Gravelines, for region "GRA1"`,
				},
				resource.Attribute{
					Name:        "continentCode",
					Description: `(Deprecated) Use ` + "`" + `continent_code` + "`" + ` instead.`,
				},
				resource.Attribute{
					Name:        "datacenterLocation",
					Description: `(Deprecated) Use ` + "`" + `datacenter_location` + "`" + ` instead.`,
				},
				resource.Attribute{
					Name:        "services",
					Description: `The list of public cloud services running within the region`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `the name of the public cloud service`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `the status of the service`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_cloud_regions (deprecated)",
			Category:         "Data Sources",
			ShortDescription: `Get the list of regions associated with a public cloud project.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The id of the public cloud project. If omitted, the ` + "`" + `OVH_PROJECT_ID` + "`" + ` environment variable is used.`,
				},
				resource.Attribute{
					Name:        "has_services_up",
					Description: `(Optional) List of services which has to be UP in regions. Example: "image", "instance", "network", "storage", "volume", "workflow", ... If left blank, returns all regions associated with the project_id. ## Attributes Reference ` + "`" + `id` + "`" + ` is set to the ID of the project. In addition, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `The list of regions associated with the project, filtered by services UP.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "names",
					Description: `The list of regions associated with the project, filtered by services UP.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_dedicated_ceph",
			Category:         "Data Sources",
			ShortDescription: `Get information & status of a dedicated CEPH instance.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "service_name",
					Description: `(Required) The service name of the dedicated CEPH cluster. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "ceph_mons",
					Description: `list of CEPH monitors IPs`,
				},
				resource.Attribute{
					Name:        "ceph_version",
					Description: `CEPH cluster version`,
				},
				resource.Attribute{
					Name:        "crush_tunables",
					Description: `CRUSH algorithm settings. Possible values`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `CEPH cluster label`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `cluster region`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `Cluster size in TB`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `the state of the cluster`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `the status of the service`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ceph_mons",
					Description: `list of CEPH monitors IPs`,
				},
				resource.Attribute{
					Name:        "ceph_version",
					Description: `CEPH cluster version`,
				},
				resource.Attribute{
					Name:        "crush_tunables",
					Description: `CRUSH algorithm settings. Possible values`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `CEPH cluster label`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `cluster region`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `Cluster size in TB`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `the state of the cluster`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `the status of the service`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_dedicated_installation_templates",
			Category:         "Data Sources",
			ShortDescription: `Get the list of installation templates available for dedicated servers.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "result",
					Description: `The list of installation templates IDs available for dedicated servers.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "result",
					Description: `The list of installation templates IDs available for dedicated servers.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_dedicated_server",
			Category:         "Data Sources",
			ShortDescription: `Get information of a dedicated server associated with your OVH Account.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "service_name",
					Description: `(Required) The service_name of your dedicated server. ## Attributes Reference ` + "`" + `id` + "`" + ` is set with the service_name of the dedicated server. In addition, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "boot_id",
					Description: `boot id of the server`,
				},
				resource.Attribute{
					Name:        "commercial_range",
					Description: `dedicater server commercial range`,
				},
				resource.Attribute{
					Name:        "datacenter",
					Description: `dedicated datacenter localisation (bhs1,bhs2,...)`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `dedicated server ip (IPv4)`,
				},
				resource.Attribute{
					Name:        "ips",
					Description: `dedicated server ip blocks`,
				},
				resource.Attribute{
					Name:        "link_speed",
					Description: `link speed of the server`,
				},
				resource.Attribute{
					Name:        "monitoring",
					Description: `Icmp monitoring state`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `dedicated server name`,
				},
				resource.Attribute{
					Name:        "os",
					Description: `Operating system`,
				},
				resource.Attribute{
					Name:        "professional_use",
					Description: `Does this server have professional use option`,
				},
				resource.Attribute{
					Name:        "rack",
					Description: `rack id of the server`,
				},
				resource.Attribute{
					Name:        "rescue_mail",
					Description: `rescue mail of the server`,
				},
				resource.Attribute{
					Name:        "reverse",
					Description: `dedicated server reverse`,
				},
				resource.Attribute{
					Name:        "root_device",
					Description: `root device of the server`,
				},
				resource.Attribute{
					Name:        "server_id",
					Description: `your server id`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `error, hacked, hackedBlocked, ok`,
				},
				resource.Attribute{
					Name:        "support_level",
					Description: `Dedicated server support level (critical, fastpath, gs, pro)`,
				},
				resource.Attribute{
					Name:        "vnis",
					Description: `the list of Virtualnetworkinterface assiociated with this server`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `VirtualNetworkInterface activation state`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `VirtualNetworkInterface mode (public,vrack,vrack_aggregation)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `User defined VirtualNetworkInterface name`,
				},
				resource.Attribute{
					Name:        "server_name",
					Description: `Server bound to this VirtualNetworkInterface`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `VirtualNetworkInterface unique id`,
				},
				resource.Attribute{
					Name:        "vrack",
					Description: `vRack name`,
				},
				resource.Attribute{
					Name:        "ncis",
					Description: `NetworkInterfaceControllers bound to this VirtualNetworkInterface`,
				},
				resource.Attribute{
					Name:        "enabled_vrack_vnis",
					Description: `List of enabled vrack VNI uuids`,
				},
				resource.Attribute{
					Name:        "enabled_vrack_aggregation_vnis",
					Description: `List of enabled vrack_aggregation VNI uuids`,
				},
				resource.Attribute{
					Name:        "enabled_public_vnis",
					Description: `List of enabled public VNI uuids`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "boot_id",
					Description: `boot id of the server`,
				},
				resource.Attribute{
					Name:        "commercial_range",
					Description: `dedicater server commercial range`,
				},
				resource.Attribute{
					Name:        "datacenter",
					Description: `dedicated datacenter localisation (bhs1,bhs2,...)`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `dedicated server ip (IPv4)`,
				},
				resource.Attribute{
					Name:        "ips",
					Description: `dedicated server ip blocks`,
				},
				resource.Attribute{
					Name:        "link_speed",
					Description: `link speed of the server`,
				},
				resource.Attribute{
					Name:        "monitoring",
					Description: `Icmp monitoring state`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `dedicated server name`,
				},
				resource.Attribute{
					Name:        "os",
					Description: `Operating system`,
				},
				resource.Attribute{
					Name:        "professional_use",
					Description: `Does this server have professional use option`,
				},
				resource.Attribute{
					Name:        "rack",
					Description: `rack id of the server`,
				},
				resource.Attribute{
					Name:        "rescue_mail",
					Description: `rescue mail of the server`,
				},
				resource.Attribute{
					Name:        "reverse",
					Description: `dedicated server reverse`,
				},
				resource.Attribute{
					Name:        "root_device",
					Description: `root device of the server`,
				},
				resource.Attribute{
					Name:        "server_id",
					Description: `your server id`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `error, hacked, hackedBlocked, ok`,
				},
				resource.Attribute{
					Name:        "support_level",
					Description: `Dedicated server support level (critical, fastpath, gs, pro)`,
				},
				resource.Attribute{
					Name:        "vnis",
					Description: `the list of Virtualnetworkinterface assiociated with this server`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `VirtualNetworkInterface activation state`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `VirtualNetworkInterface mode (public,vrack,vrack_aggregation)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `User defined VirtualNetworkInterface name`,
				},
				resource.Attribute{
					Name:        "server_name",
					Description: `Server bound to this VirtualNetworkInterface`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `VirtualNetworkInterface unique id`,
				},
				resource.Attribute{
					Name:        "vrack",
					Description: `vRack name`,
				},
				resource.Attribute{
					Name:        "ncis",
					Description: `NetworkInterfaceControllers bound to this VirtualNetworkInterface`,
				},
				resource.Attribute{
					Name:        "enabled_vrack_vnis",
					Description: `List of enabled vrack VNI uuids`,
				},
				resource.Attribute{
					Name:        "enabled_vrack_aggregation_vnis",
					Description: `List of enabled vrack_aggregation VNI uuids`,
				},
				resource.Attribute{
					Name:        "enabled_public_vnis",
					Description: `List of enabled public VNI uuids`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_dedicated_server_boots",
			Category:         "Data Sources",
			ShortDescription: `Get the list of compatible netboots for a dedicated server associated with your OVH Account.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "service_name",
					Description: `(Required) The internal name of your dedicated server.`,
				},
				resource.Attribute{
					Name:        "boot_type",
					Description: `(Optional) Filter the value of bootType property (harddisk, rescue, ipxeCustomerScript, internal, network) ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "result",
					Description: `The list of dedicated server netboots.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "result",
					Description: `The list of dedicated server netboots.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_dedicated_servers",
			Category:         "Data Sources",
			ShortDescription: `Get the list of dedicated servers associated with your OVH Account.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "result",
					Description: `The list of dedicated servers IDs associated with your OVH Account.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "result",
					Description: `The list of dedicated servers IDs associated with your OVH Account.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_domain_zone",
			Category:         "Data Sources",
			ShortDescription: `Get information & status of a domain zone.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the domain zone. ## Attributes Reference ` + "`" + `id` + "`" + ` is set to the domain zone name. In addition, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "last_update",
					Description: `Last update date of the DNS zone`,
				},
				resource.Attribute{
					Name:        "has_dns_anycast",
					Description: `hasDnsAnycast flag of the DNS zone`,
				},
				resource.Attribute{
					Name:        "name_servers",
					Description: `Name servers that host the DNS zone`,
				},
				resource.Attribute{
					Name:        "dnssec_supported",
					Description: `Is DNSSEC supported by this zone`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "last_update",
					Description: `Last update date of the DNS zone`,
				},
				resource.Attribute{
					Name:        "has_dns_anycast",
					Description: `hasDnsAnycast flag of the DNS zone`,
				},
				resource.Attribute{
					Name:        "name_servers",
					Description: `Name servers that host the DNS zone`,
				},
				resource.Attribute{
					Name:        "dnssec_supported",
					Description: `Is DNSSEC supported by this zone`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_iploadbalancing",
			Category:         "Data Sources",
			ShortDescription: `Get information & status of an IP Load Balancing product.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ipv6",
					Description: `The IPV6 associated to your IP load balancing`,
				},
				resource.Attribute{
					Name:        "ipv4",
					Description: `The IPV4 associated to your IP load balancing`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `Location where your service is. This takes an array of values.`,
				},
				resource.Attribute{
					Name:        "offer",
					Description: `The offer of your IP load balancing`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `The internal name of your IP load balancing`,
				},
				resource.Attribute{
					Name:        "ip_loadbalancing",
					Description: `Your IP load balancing`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `Current state of your IP. Can take any of the following value: "blacklisted", "deleted", "free", "ok", "quarantined", "suspended"`,
				},
				resource.Attribute{
					Name:        "vrack_eligibility",
					Description: `Vrack eligibility. Takes a boolean value.`,
				},
				resource.Attribute{
					Name:        "vrack_name",
					Description: `Name of the vRack on which the current Load Balancer is attached to, as it is named on vRack product`,
				},
				resource.Attribute{
					Name:        "ssl_configuration",
					Description: `Modern oldest compatible clients : Firefox 27, Chrome 30, IE 11 on Windows 7, Edge, Opera 17, Safari 9, Android 5.0, and Java 8. Intermediate oldest compatible clients : Firefox 1, Chrome 1, IE 7, Opera 5, Safari 1, Windows XP IE8, Android 2.3, Java 7. Can take any of the following value: "intermediate", "modern"`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `the name displayed in ManagerV6 for your iplb (max 50 chars) ## Attributes Reference ` + "`" + `id` + "`" + ` is set to the service_name of your IP load balancing In addition, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "metrics_token",
					Description: `The metrics token associated with your IP load balancing This attribute is sensitive.`,
				},
				resource.Attribute{
					Name:        "orderable_zone",
					Description: `Available additional zone for your Load Balancer`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The zone three letter code`,
				},
				resource.Attribute{
					Name:        "plan_code",
					Description: `The billing planCode for this zone`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "metrics_token",
					Description: `The metrics token associated with your IP load balancing This attribute is sensitive.`,
				},
				resource.Attribute{
					Name:        "orderable_zone",
					Description: `Available additional zone for your Load Balancer`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The zone three letter code`,
				},
				resource.Attribute{
					Name:        "plan_code",
					Description: `The billing planCode for this zone`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_iploadbalancing_vrack_network",
			Category:         "Data Sources",
			ShortDescription: `Get the details of Vrack network available for your IPLoadbalancer associated with your OVH account.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "service_name",
					Description: `(Required) The internal name of your IP load balancing`,
				},
				resource.Attribute{
					Name:        "vrack_network_id",
					Description: `(Required) Internal Load Balancer identifier of the vRack private network ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `Human readable name for your vrack network`,
				},
				resource.Attribute{
					Name:        "nat_ip",
					Description: `An IP block used as a pool of IPs by this Load Balancer to connect to the servers in this private network. The blck must be in the private network and reserved for the Load Balancer`,
				},
				resource.Attribute{
					Name:        "subnet",
					Description: `IP block of the private network in the vRack`,
				},
				resource.Attribute{
					Name:        "vlan",
					Description: `VLAN of the private network in the vRack. 0 if the private network is not in a VLAN`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "display_name",
					Description: `Human readable name for your vrack network`,
				},
				resource.Attribute{
					Name:        "nat_ip",
					Description: `An IP block used as a pool of IPs by this Load Balancer to connect to the servers in this private network. The blck must be in the private network and reserved for the Load Balancer`,
				},
				resource.Attribute{
					Name:        "subnet",
					Description: `IP block of the private network in the vRack`,
				},
				resource.Attribute{
					Name:        "vlan",
					Description: `VLAN of the private network in the vRack. 0 if the private network is not in a VLAN`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_iploadbalancing_vrack_networks",
			Category:         "Data Sources",
			ShortDescription: `Get the list of Vrack network ids available for your IPLoadbalancer associated with your OVH account.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "service_name",
					Description: `(Required) The internal name of your IP load balancing`,
				},
				resource.Attribute{
					Name:        "subnet",
					Description: `Filters networks on the subnet.`,
				},
				resource.Attribute{
					Name:        "vlan_id",
					Description: `Filters networks on the vlan id. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "result",
					Description: `The list of vrack network ids.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "result",
					Description: `The list of vrack network ids.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_me_identity_user",
			Category:         "Data Sources",
			ShortDescription: `Get information about identity User.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "user",
					Description: `(Required) User's login. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "login",
					Description: `User's login suffix.`,
				},
				resource.Attribute{
					Name:        "creation",
					Description: `Creation date of this user.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `User description.`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `User's email.`,
				},
				resource.Attribute{
					Name:        "group",
					Description: `User's group.`,
				},
				resource.Attribute{
					Name:        "last_update",
					Description: `Last update of this user.`,
				},
				resource.Attribute{
					Name:        "password_last_update",
					Description: `When the user changed his password for the last time.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Current user's status.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "login",
					Description: `User's login suffix.`,
				},
				resource.Attribute{
					Name:        "creation",
					Description: `Creation date of this user.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `User description.`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `User's email.`,
				},
				resource.Attribute{
					Name:        "group",
					Description: `User's group.`,
				},
				resource.Attribute{
					Name:        "last_update",
					Description: `Last update of this user.`,
				},
				resource.Attribute{
					Name:        "password_last_update",
					Description: `When the user changed his password for the last time.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Current user's status.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_me_identity_users",
			Category:         "Data Sources",
			ShortDescription: `Get the list of the identity users for the account.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "users",
					Description: `The list of the user's logins of all the identity users.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "users",
					Description: `The list of the user's logins of all the identity users.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_me_installation_template",
			Category:         "Data Sources",
			ShortDescription: `Get a custom installation template available for dedicated servers.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_me_installation_templates",
			Category:         "Data Sources",
			ShortDescription: `Get the list of custom installation templates available for dedicated servers.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "result",
					Description: `The list of custom installation templates IDs available for dedicated servers.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "result",
					Description: `The list of custom installation templates IDs available for dedicated servers.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_me_ipxe_script",
			Category:         "Data Sources",
			ShortDescription: `Get information & status of an IPXE Script.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the IPXE Script. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "script",
					Description: `The content of the script.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "script",
					Description: `The content of the script.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_me_ipxe_scripts",
			Category:         "Data Sources",
			ShortDescription: `Get the list of the IPXE Scripts of the account.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "result",
					Description: `The list of the names of all the IPXE Scripts.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "result",
					Description: `The list of the names of all the IPXE Scripts.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_me_paymentmean_bankaccount",
			Category:         "Data Sources",
			ShortDescription: `Get information & status of an ovh bank account payment mean`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "description_regexp",
					Description: `(Optional) a regexp used to filter bank accounts on their ` + "`" + `description` + "`" + ` attributes.`,
				},
				resource.Attribute{
					Name:        "use_default",
					Description: `(Optional) Retrieve bank account marked as default payment mean.`,
				},
				resource.Attribute{
					Name:        "use_oldest",
					Description: `(Optional) Retrieve oldest bank account. project.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `(Optional) Filter bank accounts on their ` + "`" + `state` + "`" + ` attribute. Can be "blockedForIncidents", "valid", "pendingValidation" ## Attributes Reference ` + "`" + `id` + "`" + ` is set to the ID of the bank account payment mean`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `the description attribute of the bank account`,
				},
				resource.Attribute{
					Name:        "default",
					Description: `a boolean which tells if the retrieved bank account is marked as the default payment mean`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `the description attribute of the bank account`,
				},
				resource.Attribute{
					Name:        "default",
					Description: `a boolean which tells if the retrieved bank account is marked as the default payment mean`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_me_paymentmean_creditcard",
			Category:         "Data Sources",
			ShortDescription: `Get information & status of an ovh credit card payment mean`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "description_regexp",
					Description: `(Optional) a regexp used to filter credit cards on their ` + "`" + `description` + "`" + ` attributes.`,
				},
				resource.Attribute{
					Name:        "use_default",
					Description: `(Optional) Retrieve credit card marked as default payment mean.`,
				},
				resource.Attribute{
					Name:        "use_last_to_expire",
					Description: `(Optional) Retrieve the credit card that will be the last to expire according to its expiration date.`,
				},
				resource.Attribute{
					Name:        "states",
					Description: `(Optional) Filter credit cards on their ` + "`" + `state` + "`" + ` attribute. Can be "expired", "valid", "tooManyFailures" ## Attributes Reference ` + "`" + `id` + "`" + ` is set to the ID of the credit card payment mean`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `the description attribute of the credit card`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `the state attribute of the credit card`,
				},
				resource.Attribute{
					Name:        "default",
					Description: `a boolean which tells if the retrieved credit card is marked as the default payment mean`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `the description attribute of the credit card`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `the state attribute of the credit card`,
				},
				resource.Attribute{
					Name:        "default",
					Description: `a boolean which tells if the retrieved credit card is marked as the default payment mean`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_me_ssh_key",
			Category:         "Data Sources",
			ShortDescription: `Get information & status of an SSH key.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key_name",
					Description: `(Required) The name of the SSH key. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "key_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The content of the public key. E.g.: "ssh-ed25519 AAAAC3..."`,
				},
				resource.Attribute{
					Name:        "default",
					Description: `True when this public SSH key is used for rescue mode and reinstallations.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "key_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The content of the public key. E.g.: "ssh-ed25519 AAAAC3..."`,
				},
				resource.Attribute{
					Name:        "default",
					Description: `True when this public SSH key is used for rescue mode and reinstallations.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_me_ssh_keys",
			Category:         "Data Sources",
			ShortDescription: `Get the list of the SSH keys of the account.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "names",
					Description: `The list of the names of all the SSH keys.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "names",
					Description: `The list of the names of all the SSH keys.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_vracks",
			Category:         "Data Sources",
			ShortDescription: `Get the list of Vrack ids available for your OVH account.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
	}

	dataSourcesMap = map[string]int{

		"ovh_cloud_project_region":             0,
		"ovh_cloud_project_regions":            1,
		"ovh_cloud_region (deprecated)":        2,
		"ovh_cloud_regions (deprecated)":       3,
		"ovh_dedicated_ceph":                   4,
		"ovh_dedicated_installation_templates": 5,
		"ovh_dedicated_server":                 6,
		"ovh_dedicated_server_boots":           7,
		"ovh_dedicated_servers":                8,
		"ovh_domain_zone":                      9,
		"ovh_iploadbalancing":                  10,
		"ovh_iploadbalancing_vrack_network":    11,
		"ovh_iploadbalancing_vrack_networks":   12,
		"ovh_me_identity_user":                 13,
		"ovh_me_identity_users":                14,
		"ovh_me_installation_template":         15,
		"ovh_me_installation_templates":        16,
		"ovh_me_ipxe_script":                   17,
		"ovh_me_ipxe_scripts":                  18,
		"ovh_me_paymentmean_bankaccount":       19,
		"ovh_me_paymentmean_creditcard":        20,
		"ovh_me_ssh_key":                       21,
		"ovh_me_ssh_keys":                      22,
		"ovh_vracks":                           23,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
