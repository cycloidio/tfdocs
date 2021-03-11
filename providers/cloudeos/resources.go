package cloudeos

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "cloudeos_cloudeos_aws_vpn",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cgw_id",
					Description: `(Required) AWS Customer Gateway ID`,
				},
				resource.Attribute{
					Name:        "cnps",
					Description: `(Required) VRF Segment in which the Ipsec VPN is created.`,
				},
				resource.Attribute{
					Name:        "router_id",
					Description: `(Required) CloudEOS Router to which the AWS Ipsec VPN terminates.`,
				},
				resource.Attribute{
					Name:        "vpn_connection_id",
					Description: `(Required) AWS Site-to-Site VPN Connection ID`,
				},
				resource.Attribute{
					Name:        "tunnel1_aws_endpoint_ip",
					Description: `(Required) AWS Tunnel1 Underlay IP Address`,
				},
				resource.Attribute{
					Name:        "tunnel1_aws_overlay_ip",
					Description: `(Required) VPN Tunnel1 IP address`,
				},
				resource.Attribute{
					Name:        "tunnel1_router_overlay_ip",
					Description: `(Required) CloudEOS Router Tunnel1 IP address`,
				},
				resource.Attribute{
					Name:        "tunnel1_bgp_asn",
					Description: `(Required) AWS VPN Tunnel1 BGP ASN`,
				},
				resource.Attribute{
					Name:        "tunnel1_bgp_holdtime",
					Description: `(Required) VPN Tunnel1 BGP Hold time`,
				},
				resource.Attribute{
					Name:        "tunnel1_preshared_key",
					Description: `(Required) VPN Tunnel1 Ipsec Preshared key`,
				},
				resource.Attribute{
					Name:        "tunnel2_aws_endpoint_ip",
					Description: `(Required) AWS VPN Tunnel2 Underlay IP Address`,
				},
				resource.Attribute{
					Name:        "tunnel2_aws_overlay_ip",
					Description: `(Required) AWS VPN Tunnel2 IP address`,
				},
				resource.Attribute{
					Name:        "tunnel2_router_overlay_ip",
					Description: `(Required) CloudEOS Router Tunnel2 IP address`,
				},
				resource.Attribute{
					Name:        "tunnel2_bgp_asn",
					Description: `(Required) AWS VPN Tunnel2 BGP ASN`,
				},
				resource.Attribute{
					Name:        "tunnel2_bgp_holdtime",
					Description: `(Required) VPN Tunnel2 BGP Hold time`,
				},
				resource.Attribute{
					Name:        "tunnel2_preshared_key",
					Description: `(Required) VPN Tunnel2 Ipsec Preshared key`,
				},
				resource.Attribute{
					Name:        "tgw_id",
					Description: `(Optional) AWS Transit Gateway ID, if the AWS Site-to-Site connection terminates on a TGW.`,
				},
				resource.Attribute{
					Name:        "vpn_gateway_id",
					Description: `(Optional) AWS VPN Gateway ID, if the AWS Site-to-Site connection terminates on a VPN Gateway.`,
				},
				resource.Attribute{
					Name:        "vpn_tgw_attachment_id",
					Description: `(Optional) AWS VPN Transit Gateway Attachment ID ## Attributes Reference In addition to Arguments listed above - the following Attributes are exported`,
				},
				resource.Attribute{
					Name:        "tf_id",
					Description: `The ID of cloudeos_aws_vpn Resource.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "tf_id",
					Description: `The ID of cloudeos_aws_vpn Resource.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudeos_cloudeos_clos",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) CLOS resource name.`,
				},
				resource.Attribute{
					Name:        "topology_name",
					Description: `(Required) Topology name that this clos resource depends on.`,
				},
				resource.Attribute{
					Name:        "cv_container_name",
					Description: `(Required) CVaaS Configlet Container Name to which the CloudLeaf Routers will be added to.`,
				},
				resource.Attribute{
					Name:        "fabric",
					Description: `(Optional) full_mesh or hub_spoke, default value is ` + "`" + `hub_spoke` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "leaf_to_edge_peering",
					Description: `(Optional) Leaf to edge VPC peering, default is ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "leaf_to_edge_igw",
					Description: `(Optional) Leaf to edge VPC connection through Internet Gateway, default is ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "leaf_encryption",
					Description: `(Optional) Support encryption using Ipsec between Leaf and Edge. Default is ` + "`" + `false` + "`" + `. ## Attributes Reference In addition to the Arguments listed above - the following Attributes are exported`,
				},
				resource.Attribute{
					Name:        "ID",
					Description: `The ID of the cloudeos_clos Resource. ## Timeouts`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Defaults to 5 minutes) Used when deleting the cloudeos_clos Resource.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ID",
					Description: `The ID of the cloudeos_clos Resource. ## Timeouts`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Defaults to 5 minutes) Used when deleting the cloudeos_clos Resource.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudeos_cloudeos_router_config",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cloud_provider",
					Description: `(Required) Cloud Provider for this deployment. Supports only aws or azure.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required) VPC/VNET ID in which this CloudEOS is deployed.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Required) Region of deployment.`,
				},
				resource.Attribute{
					Name:        "topology_name",
					Description: `(Required) Name of the topology in which this CloudEOS router is deployed in.`,
				},
				resource.Attribute{
					Name:        "intf_name",
					Description: `(Required) List of interface names.`,
				},
				resource.Attribute{
					Name:        "intf_private_ip",
					Description: `(Required) List of interface private IPs. Currently, only supports 1 IP address per interface.`,
				},
				resource.Attribute{
					Name:        "intf_type",
					Description: `(Required) List of Interface type (public, private, internal). A ` + "`" + `public` + "`" + ` interface has a public IP associated with it. An ` + "`" + `internal` + "`" + ` interface is the interface which connects the Leaf and Edge routers. And a ` + "`" + `private` + "`" + ` interface is the default GW interface for all host traffic.`,
				},
				resource.Attribute{
					Name:        "cnps",
					Description: `(Optional) Cloud Network Private Segments Name. ( VRF name )`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A mapping of tags to assign to the resource.`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `(Optional) CloudEdge or CloudLeaf (Same as VPC role).`,
				},
				resource.Attribute{
					Name:        "is_rr",
					Description: `(Optional) true if this CloudEOS acts as a Route Reflector.`,
				},
				resource.Attribute{
					Name:        "ami",
					Description: `(Optional) CloudEOS image. ( AWS only )`,
				},
				resource.Attribute{
					Name:        "key_name",
					Description: `(Optional) keypair name ( AWS only )`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `(Optional) Availability Zone of VPC. ## Attributes Reference In addition to Arguments listed above - the following Attributes are exported`,
				},
				resource.Attribute{
					Name:        "ID",
					Description: `The ID of cloudeos_router_config Resource.`,
				},
				resource.Attribute{
					Name:        "bootstrap_cfg",
					Description: `Bootstrap configuration for the CloudEOS router.`,
				},
				resource.Attribute{
					Name:        "peer_routetable_id",
					Description: `Router table ID of peer. ## Timeouts`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Default of 5 minute) Used when creating the cloudeos_config Resource.`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Defaults to 10 minutes) Used when deleting the cloudeos_config Resource.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ID",
					Description: `The ID of cloudeos_router_config Resource.`,
				},
				resource.Attribute{
					Name:        "bootstrap_cfg",
					Description: `Bootstrap configuration for the CloudEOS router.`,
				},
				resource.Attribute{
					Name:        "peer_routetable_id",
					Description: `Router table ID of peer. ## Timeouts`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Default of 5 minute) Used when creating the cloudeos_config Resource.`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Defaults to 10 minutes) Used when deleting the cloudeos_config Resource.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudeos_cloudeos_router_status",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cloud_provider",
					Description: `(Required) CloudProvider type. Supports only aws or azure.`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `(Required) Instance ID of deployed CloudEOS.`,
				},
				resource.Attribute{
					Name:        "intf_name",
					Description: `(Required) List of interface names for the routers.`,
				},
				resource.Attribute{
					Name:        "intf_id",
					Description: `(Required) List of interface IDs attached to the routers.`,
				},
				resource.Attribute{
					Name:        "intf_private_ip",
					Description: `(Required) List of private IPs attached to the interfaces.`,
				},
				resource.Attribute{
					Name:        "intf_subnet_id",
					Description: `(Required) List of subnet IDs of interfaces.`,
				},
				resource.Attribute{
					Name:        "intf_type",
					Description: `(Required) List of interface types. Values supported : public, internal, private.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Required) Region of deployment.`,
				},
				resource.Attribute{
					Name:        "cv_container",
					Description: `(Optional) Container in CVaaS to which the router will be added to.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Optional) VPC/VNET ID of the VPC in which the CloudEOS is deployed in.`,
				},
				resource.Attribute{
					Name:        "rg_name",
					Description: `(Optional) Resource group name, only for Azure.`,
				},
				resource.Attribute{
					Name:        "rg_location",
					Description: `(Optional) Resource group location, only for Azure.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A mapping of tags to assign to the resource.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `(Optional) Availability Zone in which the router is deployed in.`,
				},
				resource.Attribute{
					Name:        "primary_network_interface_id",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "availability_set_id",
					Description: `(Optional) Availability Set.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `(Optional) Public IP of interface.`,
				},
				resource.Attribute{
					Name:        "private_rt_table_ids",
					Description: `(Optional) List of private interface route table IDs.`,
				},
				resource.Attribute{
					Name:        "internal_rt_table_ids",
					Description: `(Optional) List of internal interface route table IDs.`,
				},
				resource.Attribute{
					Name:        "public_rt_table_ids",
					Description: `(Optional) List of public route table IDs.`,
				},
				resource.Attribute{
					Name:        "ha_name",
					Description: `(Optional) Cloud HA pair name.`,
				},
				resource.Attribute{
					Name:        "cnps",
					Description: `(Optional) Cloud Network Private Segments ( VRF name )`,
				},
				resource.Attribute{
					Name:        "is_rr",
					Description: `(Optional) true if this CloudEOS acts as a Route Reflector. ## Attributes Reference In addition to Arguments listed above - the following Attributes are exported`,
				},
				resource.Attribute{
					Name:        "ID",
					Description: `The ID of cloudeos_router_status Resource. ## Timeouts`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Defaults to 10 minutes) Used when deleting the cloudeos_status Resource.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ID",
					Description: `The ID of cloudeos_router_status Resource. ## Timeouts`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Defaults to 10 minutes) Used when deleting the cloudeos_status Resource.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudeos_cloudeos_subnet",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cloud_provider",
					Description: `(Required) Cloud Provider in which the subnet is being deployed. Supported: aws or azure.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required) VPC ID in which this subnet is created, equivalent to rg_name in Azure.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Required) ID of subnet deployed in AWS/Azure.`,
				},
				resource.Attribute{
					Name:        "cidr_block",
					Description: `(Required) CIDR of the subnet.`,
				},
				resource.Attribute{
					Name:        "subnet_name",
					Description: `(Required) Name of the subnet.`,
				},
				resource.Attribute{
					Name:        "vnet_name",
					Description: `(Optional) VNET name, only needed in Azure.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `(Optional) Availability zone. ## Attributes Reference In addition to the Arguments listed above - the following Attributes are exported:`,
				},
				resource.Attribute{
					Name:        "ID",
					Description: `The ID of the cloudeos_subnet resource.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ID",
					Description: `The ID of the cloudeos_subnet resource.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudeos_cloudeos_topology",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "topology_name",
					Description: `(Required) Name of the topology.`,
				},
				resource.Attribute{
					Name:        "bgp_asn",
					Description: `(Required) A range of BGP ASN’s which would be used to configure CloudEOS instances, based on the role and region in which they are being deployed. For example, a CloudEdge and CloudLeaf instance in the same region and CLOS will use iBGP and will have the same ASN. Whereas 2 CloudEdge’s in different regions use eBGP and will have different ASNs.`,
				},
				resource.Attribute{
					Name:        "vtep_ip_cidr",
					Description: `(Required) CIDR block for VTEP IPs for CloudEOS Routers`,
				},
				resource.Attribute{
					Name:        "terminattr_ip_cidr",
					Description: `(Required) TerminAttr is used by Arista devices to stream Telemetry to CVaaS. Every CloudEOS Router needs a unique TerminAttr local IP.`,
				},
				resource.Attribute{
					Name:        "dps_controlplane_cidr",
					Description: `(Required) Each CloudEOS router needs a unique IP for Dynamic Path Selection.`,
				},
				resource.Attribute{
					Name:        "eos_managed",
					Description: `(Optional) List of CloudEOS devices already deployed. ## Attributes Reference In addition to the Arguments listed above - the following Attributes are exported:`,
				},
				resource.Attribute{
					Name:        "ID",
					Description: `The ID of the cloudeos_topology Resource. ## Timeouts`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Defaults to 5 minutes) Used when deleting the Topology Resource.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ID",
					Description: `The ID of the cloudeos_topology Resource. ## Timeouts`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Defaults to 5 minutes) Used when deleting the Topology Resource.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudeos_cloudeos_vpc_config",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "topology_name",
					Description: `(Required) Name of topology resource.`,
				},
				resource.Attribute{
					Name:        "clos_name",
					Description: `(Optional) CLOS Name this VPC refers to for attributes.`,
				},
				resource.Attribute{
					Name:        "wan_name",
					Description: `(Optional) WAN Name this VPC refers to for attributes.`,
				},
				resource.Attribute{
					Name:        "rg_name",
					Description: `(Optional) Resource group name, only valid for Azure.`,
				},
				resource.Attribute{
					Name:        "vnet_name",
					Description: `(Optional) VNET name, only valid for Azure.`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `(Required) CloudEdge or CloudLeaf.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A mapping of tags to assign to the resource. ## Attributes Reference In addition to Arguments listed above - the following Attributes are exported`,
				},
				resource.Attribute{
					Name:        "ID",
					Description: `The ID of cloudeos_vpc_config Resource. A CloudLeaf VPC peers with the CloudEdge VPC to enable communication between instances between them. The following Attributes are exported in CloudLeaf VPC that provides information about the peer CloudEdge VPC.`,
				},
				resource.Attribute{
					Name:        "peer_vpc_id",
					Description: `ID of the CloudEdge peer VPC, only valid for AWS.`,
				},
				resource.Attribute{
					Name:        "peer_vpc_cidr",
					Description: `CIDR of the CloudEdge peer VPC, only valid for AWS.`,
				},
				resource.Attribute{
					Name:        "peer_vnet_id",
					Description: `ID of the CloudEdge peer VNET, only valid for Azure.`,
				},
				resource.Attribute{
					Name:        "peer_rg_name",
					Description: `Resource Group name of the peer CloudEdge, only valid for Azure.`,
				},
				resource.Attribute{
					Name:        "peer_vnet_name",
					Description: `VNET name of the peer CloudEdge, only valid for Azure. ## Timeouts`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Default of 3 minute) Used when creating the cloudeos_vpc_config Resource.`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Defaults to 5 minutes) Used when deleting the cloudeos_vpc_config Resource.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ID",
					Description: `The ID of cloudeos_vpc_config Resource. A CloudLeaf VPC peers with the CloudEdge VPC to enable communication between instances between them. The following Attributes are exported in CloudLeaf VPC that provides information about the peer CloudEdge VPC.`,
				},
				resource.Attribute{
					Name:        "peer_vpc_id",
					Description: `ID of the CloudEdge peer VPC, only valid for AWS.`,
				},
				resource.Attribute{
					Name:        "peer_vpc_cidr",
					Description: `CIDR of the CloudEdge peer VPC, only valid for AWS.`,
				},
				resource.Attribute{
					Name:        "peer_vnet_id",
					Description: `ID of the CloudEdge peer VNET, only valid for Azure.`,
				},
				resource.Attribute{
					Name:        "peer_rg_name",
					Description: `Resource Group name of the peer CloudEdge, only valid for Azure.`,
				},
				resource.Attribute{
					Name:        "peer_vnet_name",
					Description: `VNET name of the peer CloudEdge, only valid for Azure. ## Timeouts`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Default of 3 minute) Used when creating the cloudeos_vpc_config Resource.`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Defaults to 5 minutes) Used when deleting the cloudeos_vpc_config Resource.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudeos_cloudeos_vpc_status",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cloud_provider",
					Description: `(Required) The Cloud Provider in which the VPC/VNET is deployed.`,
				},
				resource.Attribute{
					Name:        "cnps",
					Description: `(Required) Cloud Network Private Segments Name. ( VRF Name )`,
				},
				resource.Attribute{
					Name:        "topology_name",
					Description: `(Required) Name of topology resource.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Required) Region of deployment.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required) VPC ID, this is equiv to vnet_id in Azure.`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `(Required) CloudEdge or CloudLeaf.`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `(Required) VPC role, CloudEdge/CloudLeaf.`,
				},
				resource.Attribute{
					Name:        "account",
					Description: `(Required) The unique identifier of the account.`,
				},
				resource.Attribute{
					Name:        "rg_name",
					Description: `(Optional) Resource group name, only valid for Azure.`,
				},
				resource.Attribute{
					Name:        "vnet_name",
					Description: `(Optional) VNET name, only valid for Azure.`,
				},
				resource.Attribute{
					Name:        "clos_name",
					Description: `(Optional) Clos Name this VPC refers to for attributes.`,
				},
				resource.Attribute{
					Name:        "wan_name",
					Description: `(Optional) Wan Name this VPC refers to for attributes.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A mapping of tags to assign to the resource.`,
				},
				resource.Attribute{
					Name:        "cidr_block",
					Description: `(Optional) CIDR Block for VPC.`,
				},
				resource.Attribute{
					Name:        "ID",
					Description: `The ID of cloudeos_vpc_status Resource. ## Timeouts`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Defaults to 5 minutes) Used when deleting the cloudeos_vpc_status Resource.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ID",
					Description: `The ID of cloudeos_vpc_status Resource. ## Timeouts`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Defaults to 5 minutes) Used when deleting the cloudeos_vpc_status Resource.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudeos_cloudeos_wan",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the wan resource.`,
				},
				resource.Attribute{
					Name:        "topology_name",
					Description: `(Required) Name of the topology this wan resource depends on.`,
				},
				resource.Attribute{
					Name:        "cv_container_name",
					Description: `(Required) CVaaS Container Name to which the CloudEdge Routers will be added to.`,
				},
				resource.Attribute{
					Name:        "edge_to_edge_igw",
					Description: `(Optional) Edge to Edge Connectivity through the Internet Gateway.`,
				},
				resource.Attribute{
					Name:        "edge_to_edge_peering",
					Description: `(Optional) Peering across Edge VPC's, default is false. ( Not supported yet )`,
				},
				resource.Attribute{
					Name:        "edge_to_edge_dedicated_connect",
					Description: `(Optional) Dedicated connection between two Edge VPC, default is false. ( Not Supported yet ) ## Attributes Reference In addition to the Arguments listed above - the following Attributes are exported`,
				},
				resource.Attribute{
					Name:        "ID",
					Description: `The ID of the cloudeos_wan resource. ## Timeouts`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Defaults to 5 minutes) Used when deleting the Wan Resource.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ID",
					Description: `The ID of the cloudeos_wan resource. ## Timeouts`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Defaults to 5 minutes) Used when deleting the Wan Resource.`,
				},
			},
		},
	}

	resourcesMap = map[string]int{

		"cloudeos_cloudeos_aws_vpn":       0,
		"cloudeos_cloudeos_clos":          1,
		"cloudeos_cloudeos_router_config": 2,
		"cloudeos_cloudeos_router_status": 3,
		"cloudeos_cloudeos_subnet":        4,
		"cloudeos_cloudeos_topology":      5,
		"cloudeos_cloudeos_vpc_config":    6,
		"cloudeos_cloudeos_vpc_status":    7,
		"cloudeos_cloudeos_wan":           8,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
