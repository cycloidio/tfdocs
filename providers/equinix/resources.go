package equinix

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "equinix_equinix_ecx_l2_connection",
			Category:         "Fabric",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"fabric",
				"ecx",
				"l2",
				"connection",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Connection name. An alpha-numeric 24 characters string which can include only hyphens and underscores`,
				},
				resource.Attribute{
					Name:        "profile_uuid",
					Description: `(Required) Unique identifier of the service provider's profile.`,
				},
				resource.Attribute{
					Name:        "speed",
					Description: `(Required) Speed/Bandwidth to be allocated to the connection.`,
				},
				resource.Attribute{
					Name:        "speed_unit",
					Description: `(Required) Unit of the speed/bandwidth to be allocated to the connection.`,
				},
				resource.Attribute{
					Name:        "notifications",
					Description: `(Required) A list of email addresses used for sending connection update notifications.`,
				},
				resource.Attribute{
					Name:        "purchase_order_number",
					Description: `(Optional) Connection's purchase order number to reflect on the invoice`,
				},
				resource.Attribute{
					Name:        "port_uuid",
					Description: `(Required when ` + "`" + `device_uuid` + "`" + ` or ` + "`" + `service_token` + "`" + ` are not set) Unique identifier of the Equinix Fabric Port from which the connection would originate.`,
				},
				resource.Attribute{
					Name:        "device_uuid",
					Description: `(Required when ` + "`" + `port_uuid` + "`" + ` or ` + "`" + `service_token` + "`" + ` are not set) Unique identifier of the Network Edge virtual device from which the connection would originate.`,
				},
				resource.Attribute{
					Name:        "device_interface_id",
					Description: `(Optional) Applicable with ` + "`" + `device_uuid` + "`" + `, identifier of network interface on a given device, used for a connection. If not specified then first available interface will be selected.`,
				},
				resource.Attribute{
					Name:        "vlan_stag",
					Description: `(Required when port_uuid is set) S-Tag/Outer-Tag of the connection - a numeric character ranging from 2 - 4094.`,
				},
				resource.Attribute{
					Name:        "vlan_ctag",
					Description: `(Optional) C-Tag/Inner-Tag of the connection - a numeric character ranging from 2 \- 4094.`,
				},
				resource.Attribute{
					Name:        "named_tag",
					Description: `(Optional) The type of peering to set up when connecting to Azure Express Route. Valid values: ` + "`" + `PRIVATE` + "`" + `, ` + "`" + `MICROSOFT` + "`" + `, ` + "`" + `MANUAL` + "`" + `\`,
				},
				resource.Attribute{
					Name:        "additional_info",
					Description: `(Optional) one or more additional information key-value objects`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) additional information key`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) additional information value`,
				},
				resource.Attribute{
					Name:        "zside_port_uuid",
					Description: `(Optional) Unique identifier of the port on the remote/destination side (z-side). Allows you to connect between your own ports or virtual devices across your company's Equinix Fabric deployment, with no need for a private service profile.`,
				},
				resource.Attribute{
					Name:        "zside_vlan_stag",
					Description: `(Optional) S-Tag/Outer-Tag of the connection on the remote/destination side (z-side) - a numeric character ranging from 2 - 4094.`,
				},
				resource.Attribute{
					Name:        "zside_vlan_ctag",
					Description: `(Optional) C-Tag/Inner-Tag of the connection on the remote/destination side (z-side) - a numeric character ranging from 2 - 4094. ` + "`" + `secondary_connection` + "`" + ` is defined it will internally use same ` + "`" + `zside_vlan_ctag` + "`" + ` for the secondary connection.`,
				},
				resource.Attribute{
					Name:        "seller_region",
					Description: `(Optional) The region in which the seller port resides.`,
				},
				resource.Attribute{
					Name:        "seller_metro_code",
					Description: `(Optional) The metro code that denotes the connection’s remote/destination side (z-side).`,
				},
				resource.Attribute{
					Name:        "authorization_key",
					Description: `(Optional) Unique identifier authorizing Equinix to provision a connection towards a cloud service provider. At Equinix, an ` + "`" + `Authorization Key` + "`" + ` is a generic term and is NOT encrypted on Equinix Fabric. Cloud Service Providers might use a different name to refer to this key such as ` + "`" + `Service Key` + "`" + ` or ` + "`" + `Authentication Key` + "`" + `. Value depends on a provider service profile, more information on [Equinix Fabric how to guide](https://developer.equinix.com/docs/ecx-how-to-guide).`,
				},
				resource.Attribute{
					Name:        "secondary_connection",
					Description: `(Optional) Definition of secondary connection for redundant, HA connectivity. See [Secondary Connection](#secondary-connection) below for more details. ### Secondary Connection ->`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) secondary connection name`,
				},
				resource.Attribute{
					Name:        "speed",
					Description: `(Optional) Speed/Bandwidth to be allocated to the secondary connection. If not specified primary ` + "`" + `speed` + "`" + ` will be used.`,
				},
				resource.Attribute{
					Name:        "speed_unit",
					Description: `(Optional) Unit of the speed/bandwidth to be allocated to the secondary connection. If not specified primary ` + "`" + `speed_unit` + "`" + ` will be used.`,
				},
				resource.Attribute{
					Name:        "port_uuid",
					Description: `(Optional) Applicable with primary ` + "`" + `port_uuid` + "`" + `. Identifier of the Equinix Fabric Port from which the secondary connection would originate. If not specified primary ` + "`" + `port_uuid` + "`" + ` will be used.`,
				},
				resource.Attribute{
					Name:        "device_uuid",
					Description: `(Optional) Applicable with primary ` + "`" + `device_uuid` + "`" + `. Identifier of the Network Edge virtual device from which the secondary connection would originate. If not specified primary ` + "`" + `device_uuid` + "`" + ` will be used.`,
				},
				resource.Attribute{
					Name:        "device_interface_id",
					Description: `(Optional) Applicable with ` + "`" + `device_uuid` + "`" + `, identifier of network interface on a given device. If not specified then first available interface will be selected.`,
				},
				resource.Attribute{
					Name:        "vlan_stag",
					Description: `(Required when ` + "`" + `port_uuid` + "`" + ` is set) S-Tag/Outer-Tag of the secondary connection, a numeric character ranging from 2 - 4094.`,
				},
				resource.Attribute{
					Name:        "vlan_ctag",
					Description: `(Optional) Applicable with ` + "`" + `port_uuid` + "`" + `. C-Tag/Inner-Tag of the secondary connection, a numeric character ranging from 2 - 4094.`,
				},
				resource.Attribute{
					Name:        "seller_metro_code",
					Description: `(Optional) The metro code that denotes the secondary connection’s destination (Z side). .`,
				},
				resource.Attribute{
					Name:        "seller_region",
					Description: `(Optional) The region in which the seller port resides. If not specified primary ` + "`" + `seller_region` + "`" + ` will be used.`,
				},
				resource.Attribute{
					Name:        "authorization_key",
					Description: `(Optional) Unique identifier authorizing Equinix to provision a connection towards a cloud service provider. If not specified primary ` + "`" + `authorization_key` + "`" + ` will be used. However, some service providers may require different keys for each connection. More information on [Equinix Fabric how to guide](https://developer.equinix.com/docs/ecx-how-to-guide). ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `Unique identifier of the connection.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Connection provisioning status on Equinix Fabric side.`,
				},
				resource.Attribute{
					Name:        "provider_status",
					Description: `Connection provisioning status on service provider's side.`,
				},
				resource.Attribute{
					Name:        "redundant_uuid",
					Description: `Unique identifier of the redundant connection, applicable for HA connections.`,
				},
				resource.Attribute{
					Name:        "redundancy_type",
					Description: `Connection redundancy type, applicable for HA connections. Valid values are ` + "`" + `PRIMARY` + "`" + `, ` + "`" + `SECONDARY` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "redundancy_group",
					Description: `Unique identifier of group containing a primary and secondary connection.`,
				},
				resource.Attribute{
					Name:        "zside_port_uuid",
					Description: `When not provided as an argument, it is identifier of the z-side port, assigned by the Fabric.`,
				},
				resource.Attribute{
					Name:        "zside_vlan_stag",
					Description: `When not provided as an argument, it is S-Tag/Outer-Tag of the connection on the Z side, assigned by the Fabric.`,
				},
				resource.Attribute{
					Name:        "actions",
					Description: `One or more pending actions to complete connection provisioning.`,
				},
				resource.Attribute{
					Name:        "vendor_token",
					Description: `The Equinix Fabric Token the connection was created with. Applicable if the connection was created with a ` + "`" + `service_token` + "`" + ` (a-side) or ` + "`" + `zside_service_token` + "`" + ` (z-side).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "uuid",
					Description: `Unique identifier of the connection.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Connection provisioning status on Equinix Fabric side.`,
				},
				resource.Attribute{
					Name:        "provider_status",
					Description: `Connection provisioning status on service provider's side.`,
				},
				resource.Attribute{
					Name:        "redundant_uuid",
					Description: `Unique identifier of the redundant connection, applicable for HA connections.`,
				},
				resource.Attribute{
					Name:        "redundancy_type",
					Description: `Connection redundancy type, applicable for HA connections. Valid values are ` + "`" + `PRIMARY` + "`" + `, ` + "`" + `SECONDARY` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "redundancy_group",
					Description: `Unique identifier of group containing a primary and secondary connection.`,
				},
				resource.Attribute{
					Name:        "zside_port_uuid",
					Description: `When not provided as an argument, it is identifier of the z-side port, assigned by the Fabric.`,
				},
				resource.Attribute{
					Name:        "zside_vlan_stag",
					Description: `When not provided as an argument, it is S-Tag/Outer-Tag of the connection on the Z side, assigned by the Fabric.`,
				},
				resource.Attribute{
					Name:        "actions",
					Description: `One or more pending actions to complete connection provisioning.`,
				},
				resource.Attribute{
					Name:        "vendor_token",
					Description: `The Equinix Fabric Token the connection was created with. Applicable if the connection was created with a ` + "`" + `service_token` + "`" + ` (a-side) or ` + "`" + `zside_service_token` + "`" + ` (z-side).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "equinix_equinix_ecx_l2_connection_accepter",
			Category:         "Fabric",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"fabric",
				"ecx",
				"l2",
				"connection",
				"accepter",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "connection_id",
					Description: `(Required) Identifier of Layer 2 connection that will be accepted.`,
				},
				resource.Attribute{
					Name:        "access_key",
					Description: `(Optional) Access Key used to accept connection on provider side.`,
				},
				resource.Attribute{
					Name:        "secret_key",
					Description: `(Optional) Secret Key used to accept connection on provider side.`,
				},
				resource.Attribute{
					Name:        "aws_profile",
					Description: `(Optional) AWS Profile Name for retrieving credentials from. shared credentials file ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "aws_connection_id",
					Description: `Identifier of a hosted Direct Connect connection on AWS side, applicable for accepter resource with connections to AWS only. ## Import This resource can be imported using an existing ID: ` + "`" + `` + "`" + `` + "`" + `sh terraform import equinix_ecx_l2_connection_accepter.example {existing_id} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "aws_connection_id",
					Description: `Identifier of a hosted Direct Connect connection on AWS side, applicable for accepter resource with connections to AWS only. ## Import This resource can be imported using an existing ID: ` + "`" + `` + "`" + `` + "`" + `sh terraform import equinix_ecx_l2_connection_accepter.example {existing_id} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "equinix_equinix_ecx_l2_serviceprofile",
			Category:         "Fabric",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"fabric",
				"ecx",
				"l2",
				"serviceprofile",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the service profile. An alpha-numeric 50 characters string which can include only hyphens and underscores.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the service profile.`,
				},
				resource.Attribute{
					Name:        "bandwidth_alert_threshold",
					Description: `(Optional) Specifies the port bandwidth threshold percentage. If the bandwidth limit is met or exceeded, an alert is sent to the seller.`,
				},
				resource.Attribute{
					Name:        "speed_customization_allowed",
					Description: `(Optional) Boolean value that determines if customer is allowed to enter a custom connection speed.`,
				},
				resource.Attribute{
					Name:        "oversubscription_allowed",
					Description: `(Optional) Boolean value that determines if, regardless of the utilization, Equinix Fabric will continue to add connections to your links until we reach the oversubscription limit. By selecting this service, you acknowledge that you will manage decisions on when to increase capacity on these link.`,
				},
				resource.Attribute{
					Name:        "api_integration",
					Description: `(Optional) Boolean value that determines if API integration is enabled. It allows you to complete connection provisioning in less than five minutes. Without API Integration, additional manual steps will be required and the provisioning will likely take longer.`,
				},
				resource.Attribute{
					Name:        "authkey_label",
					Description: `(Optional) Name of the authentication key label to be used by the Authentication Key service. It allows Service Providers with QinQ ports to accept groups of connections or VLANs from Dot1q customers. This is similar to S-Tag/C-Tag capabilities.`,
				},
				resource.Attribute{
					Name:        "connection_name_label",
					Description: `(Optional) Custom name used for calling a connections e.g. ` + "`" + `circuit` + "`" + `. Defaults to ` + "`" + `Connection` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ctag_label",
					Description: `(Optional) C-Tag/Inner-Tag label name for the connections.`,
				},
				resource.Attribute{
					Name:        "servicekey_autogenerated",
					Description: `(Optional) Boolean value that indicates whether multiple connections can be created with the same authorization key to connect to this service profile after the first connection has been approved by the seller.`,
				},
				resource.Attribute{
					Name:        "equinix_managed_port_vlan",
					Description: `(Optional) Applicable when ` + "`" + `api_integration` + "`" + ` is set to ` + "`" + `true` + "`" + `. It indicates whether the port and VLAN details are managed by Equinix.`,
				},
				resource.Attribute{
					Name:        "integration_id",
					Description: `(Optional) Specifies the API integration ID that was provided to the customer during onboarding. You can validate your API integration ID using the validateIntegrationId API.`,
				},
				resource.Attribute{
					Name:        "bandwidth_threshold_notifications",
					Description: `(Optional) A list of email addresses that will receive notifications about bandwidth thresholds.`,
				},
				resource.Attribute{
					Name:        "profile_statuschange_notifications",
					Description: `(Required) A list of email addresses that will receive notifications about profile status changes.`,
				},
				resource.Attribute{
					Name:        "vc_statuschange_notifications",
					Description: `(Required) A list of email addresses that will receive notifications about connections approvals and rejections.`,
				},
				resource.Attribute{
					Name:        "oversubscription",
					Description: `(Optional) You can set an alert for when a percentage of your profile has been sold. Service providers like to use this functionality to alert them when they need to add more ports or when they need to create a new service profile. Required with ` + "`" + `oversubscription_allowed` + "`" + `, defaults to ` + "`" + `1x` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "private",
					Description: `(Optional) Boolean value that indicates whether or not this is a private profile, i.e. not public like AWS/Azure/Oracle/Google, etc. If private, it can only be available for creating connections if correct permissions are granted.`,
				},
				resource.Attribute{
					Name:        "private_user_emails",
					Description: `(Optional) An array of users email ids who have permission to access this service profile. Argument is required when profile is set as private.`,
				},
				resource.Attribute{
					Name:        "redundancy_required",
					Description: `(Optional) Boolean value that determines if your connections will require redundancy. if yes, then users need to create a secondary redundant connection.`,
				},
				resource.Attribute{
					Name:        "speed_from_api",
					Description: `(Optional) Boolean valuta that determines if connection speed will be derived from an API call. Argument has to be specified when ` + "`" + `api_integration` + "`" + ` is enabled.`,
				},
				resource.Attribute{
					Name:        "tag_type",
					Description: `(Optional) Specifies additional tagging information required by the seller profile for Dot1Q to QinQ translation. See [Enhance Dot1q to QinQ translation support](https://docs.equinix.com/es/Content/Interconnection/Fabric/layer-2/Fabric-Create-Layer2-Service-Profile.htm#:~:text=Enhance%20Dot1q%20to%20QinQ%20translation%20support) for additional information. Valid values are:`,
				},
				resource.Attribute{
					Name:        "CTAGED",
					Description: `When seller side VLAN C-Tag has to be provided _(Default)_.`,
				},
				resource.Attribute{
					Name:        "NAMED",
					Description: `When application named tag has to be provided.`,
				},
				resource.Attribute{
					Name:        "BOTH",
					Description: `When both, application tag or seller side VLAN C-Tag can be provided.`,
				},
				resource.Attribute{
					Name:        "secondary_vlan_from_primary",
					Description: `(Optional) Indicates whether the VLAN ID of. the secondary connection is the same as the primary connection.`,
				},
				resource.Attribute{
					Name:        "features",
					Description: `(Required) Block of profile features configuration. See [Features](#features) below for more details.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Required) One or more definitions of ports residing in locations, from which your customers will be able to access services using this service profile. See [Port](#port) below for more details.`,
				},
				resource.Attribute{
					Name:        "speed_band",
					Description: `(Optional) One or more definitions of supported speed/bandwidth. Argument is required when ` + "`" + `speed_from_api` + "`" + ` is set to ` + "`" + `false` + "`" + `. See [Speed Band](#speed-band) below for more details. ### Features The ` + "`" + `features` + "`" + ` block has below fields:`,
				},
				resource.Attribute{
					Name:        "allow_remote_connections",
					Description: `(Required) Indicates whether or not connections to this profile can be created from remote metro locations.`,
				},
				resource.Attribute{
					Name:        "test_profile",
					Description: `(Deprecated) Indicates whether or not this profile can be used for test connections. ### Port Each ` + "`" + `port` + "`" + ` block has below fields:`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `(Required) Unique identifier of the port.`,
				},
				resource.Attribute{
					Name:        "metro_code",
					Description: `(Required) The metro code of location where the port resides. ### Speed Band Each ` + "`" + `speed_band` + "`" + ` block has below fields:`,
				},
				resource.Attribute{
					Name:        "speed",
					Description: `(Required) Speed/bandwidth supported by this service profile.`,
				},
				resource.Attribute{
					Name:        "speed_unit",
					Description: `(Required) Unit of the speed/bandwidth supported by this service profile. One of ` + "`" + `MB` + "`" + `, ` + "`" + `GB` + "`" + `. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `Unique identifier of the service profile.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `Service profile provisioning status. ## Import This resource can be imported using an existing ID: ` + "`" + `` + "`" + `` + "`" + `sh terraform import equinix_ecx_l2_serviceprofile.example {existing_id} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "uuid",
					Description: `Unique identifier of the service profile.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `Service profile provisioning status. ## Import This resource can be imported using an existing ID: ` + "`" + `` + "`" + `` + "`" + `sh terraform import equinix_ecx_l2_serviceprofile.example {existing_id} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "equinix_equinix_fabric_connection",
			Category:         "Fabric",
			ShortDescription: `Fabric V4 API compatible resource allows creation and management of Equinix Fabric connection ~> Note Equinix Fabric v4 resources and datasources are currently in Beta. The interfaces related to equinix_fabric_ resources and datasources may change ahead of general availability. Please, do not hesitate to report any problems that you experience by opening a new issue https://github.com/equinix/terraform-provider-equinix/issues/new?template=bug.md`,
			Description:      ``,
			Keywords: []string{
				"fabric",
				"connection",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "equinix_equinix_fabric_service_profile",
			Category:         "Fabric",
			ShortDescription: `Fabric V4 API compatible resource allows creation and management of Equinix Fabric Service Profile ~> Note Equinix Fabric v4 resources and datasources are currently in Beta. The interfaces related to equinix_fabric_ resources and datasources may change ahead of general availability. Please, do not hesitate to report any problems that you experience by opening a new issue https://github.com/equinix/terraform-provider-equinix/issues/new?template=bug.md`,
			Description:      ``,
			Keywords: []string{
				"fabric",
				"service",
				"profile",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "equinix_equinix_metal_bgp_session",
			Category:         "Metal",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"metal",
				"bgp",
				"session",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "device_id",
					Description: `(Required) ID of device.`,
				},
				resource.Attribute{
					Name:        "address_family",
					Description: `(Required) ` + "`" + `ipv4` + "`" + ` or ` + "`" + `ipv6` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "default_route",
					Description: `(Optional) Boolean flag to set the default route policy. False by default. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "equinix_equinix_metal_connection",
			Category:         "Metal",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"metal",
				"connection",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the connection resource`,
				},
				resource.Attribute{
					Name:        "metro",
					Description: `(Optional) Metro where the connection will be created.`,
				},
				resource.Attribute{
					Name:        "facility",
					Description: `(`,
				},
				resource.Attribute{
					Name:        "redundancy",
					Description: `(Required) Connection redundancy - redundant or primary.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) Connection type - dedicated or shared.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Optional) ID of the project where the connection is scoped to, must be set for.`,
				},
				resource.Attribute{
					Name:        "speed",
					Description: `(Required) Connection speed - one of 50Mbps, 200Mbps, 500Mbps, 1Gbps, 2Gbps, 5Gbps, 10Gbps.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description for the connection resource.`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `(Optional) Mode for connections in IBX facilities with the dedicated type - standard or tunnel. Default is standard.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) String list of tags.`,
				},
				resource.Attribute{
					Name:        "vlans",
					Description: `(Optional) Only used with shared connection. Vlans to attach. Pass one vlan for Primary/Single connection and two vlans for Redundant connection.`,
				},
				resource.Attribute{
					Name:        "service_token_type",
					Description: `(Optional) Only used with shared connection. Type of service token to use for the connection, a_side or z_side. (`,
				},
				resource.Attribute{
					Name:        "organization_id",
					Description: `ID of the organization where the connection is scoped to.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the connection resource.`,
				},
				resource.Attribute{
					Name:        "ports",
					Description: `List of connection ports - primary (` + "`" + `ports[0]` + "`" + `) and secondary (` + "`" + `ports[1]` + "`" + `). Schema of port is described in documentation of the [equinix_metal_connection datasource](../data-sources/equinix_metal_connection.md).`,
				},
				resource.Attribute{
					Name:        "service_tokens",
					Description: `List of connection service tokens with attributes required to configure the connection in Equinix Fabric with the [equinix_ecx_l2_connection](./equinix_ecx_l2_connection.md) resource or from the [Equinix Fabric Portal](https://ecxfabric.equinix.com/dashboard). Scehma of service_token is described in documentation of the [equinix_metal_connection datasource](../data-sources/equinix_metal_connection.md).`,
				},
				resource.Attribute{
					Name:        "token",
					Description: `(Deprecated) Fabric Token required to configure the connection in Equinix Fabric with the [equinix_ecx_l2_connection](./equinix_ecx_l2_connection.md) resource or from the [Equinix Fabric Portal](https://ecxfabric.equinix.com/dashboard). If your organization already has connection service tokens enabled, use ` + "`" + `service_tokens` + "`" + ` instead.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "organization_id",
					Description: `ID of the organization where the connection is scoped to.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the connection resource.`,
				},
				resource.Attribute{
					Name:        "ports",
					Description: `List of connection ports - primary (` + "`" + `ports[0]` + "`" + `) and secondary (` + "`" + `ports[1]` + "`" + `). Schema of port is described in documentation of the [equinix_metal_connection datasource](../data-sources/equinix_metal_connection.md).`,
				},
				resource.Attribute{
					Name:        "service_tokens",
					Description: `List of connection service tokens with attributes required to configure the connection in Equinix Fabric with the [equinix_ecx_l2_connection](./equinix_ecx_l2_connection.md) resource or from the [Equinix Fabric Portal](https://ecxfabric.equinix.com/dashboard). Scehma of service_token is described in documentation of the [equinix_metal_connection datasource](../data-sources/equinix_metal_connection.md).`,
				},
				resource.Attribute{
					Name:        "token",
					Description: `(Deprecated) Fabric Token required to configure the connection in Equinix Fabric with the [equinix_ecx_l2_connection](./equinix_ecx_l2_connection.md) resource or from the [Equinix Fabric Portal](https://ecxfabric.equinix.com/dashboard). If your organization already has connection service tokens enabled, use ` + "`" + `service_tokens` + "`" + ` instead.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "equinix_equinix_metal_device",
			Category:         "Metal",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"metal",
				"device",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "always_pxe",
					Description: `(Optional) If true, a device with OS ` + "`" + `custom_ipxe` + "`" + ` will continue to boot via iPXE on reboots.`,
				},
				resource.Attribute{
					Name:        "behavior",
					Description: `(Optional) Behavioral overrides that change how the resource handles certain attribute updates. See [Behavior](#behavior) below for more details.`,
				},
				resource.Attribute{
					Name:        "billing_cycle",
					Description: `(Optional) monthly or hourly`,
				},
				resource.Attribute{
					Name:        "custom_data",
					Description: `(Optional) A string of the desired Custom Data for the device. By default, changing this attribute will cause the provider to destroy and recreate your device. If ` + "`" + `reinstall` + "`" + ` is specified or ` + "`" + `behavior.allow_changes` + "`" + ` includes ` + "`" + `"custom_data"` + "`" + `, the device will be updated in-place instead of recreated.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The device description.`,
				},
				resource.Attribute{
					Name:        "facilities",
					Description: `(`,
				},
				resource.Attribute{
					Name:        "force_detach_volumes",
					Description: `(Optional) Delete device even if it has volumes attached. Only applies for destroy action.`,
				},
				resource.Attribute{
					Name:        "hardware_reservation_id",
					Description: `(Optional) The UUID of the hardware reservation where you want this device deployed, or ` + "`" + `next-available` + "`" + ` if you want to pick your next available reservation automatically. Changing this from a reservation UUID to ` + "`" + `next-available` + "`" + ` will re-create the device in another reservation. Please be careful when using hardware reservation UUID and ` + "`" + `next-available` + "`" + ` together for the same pool of reservations. It might happen that the reservation which Equinix Metal API will pick as ` + "`" + `next-available` + "`" + ` is the reservation which you refer with UUID in another equinix_metal_device resource. If that happens, and the equinix_metal_device with the UUID is created later, resource creation will fail because the reservation is already in use (by the resource created with ` + "`" + `next-available` + "`" + `). To workaround this, have the ` + "`" + `next-available` + "`" + ` resource [explicitly depend_on](https://learn.hashicorp.com/terraform/getting-started/dependencies.html#implicit-and-explicit-dependencies) the resource with hardware reservation UUID, so that the latter is created first. For more details, see [issue #176](https://github.com/packethost/terraform-provider-packet/issues/176).`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `(Optional) The device hostname used in deployments taking advantage of Layer3 DHCP or metadata service configuration.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `(Optional) A list of IP address types for the device. See [IP address](#ip-address) below for more details.`,
				},
				resource.Attribute{
					Name:        "ipxe_script_url",
					Description: `(Optional) URL pointing to a hosted iPXE script. More information is in the [Custom iPXE](https://metal.equinix.com/developers/docs/servers/custom-ipxe/) doc.`,
				},
				resource.Attribute{
					Name:        "metro",
					Description: `(Optional) Metro area for the new device. Conflicts with ` + "`" + `facilities` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "operating_system",
					Description: `(Required) The operating system slug. To find the slug, or visit [Operating Systems API docs](https://metal.equinix.com/developers/api/operatingsystems), set your API auth token in the top of the page and see JSON from the API response.`,
				},
				resource.Attribute{
					Name:        "plan",
					Description: `(Required) The device plan slug. To find the plan slug, visit [Device plans API docs](https://metal.equinix.com/developers/api/plans), set your auth token in the top of the page and see JSON from the API response.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The ID of the project in which to create the device`,
				},
				resource.Attribute{
					Name:        "project_ssh_key_ids",
					Description: `(Optional) Array of IDs of the project SSH keys which should be added to the device. If you omit this, SSH keys of all the members of the parent project will be added to the device. If you specify this array, only the listed project SSH keys will be added. Project SSH keys can be created with the [equinix_metal_project_ssh_key](equinix_metal_project_ssh_key.md) resource.`,
				},
				resource.Attribute{
					Name:        "user_ssh_key_ids",
					Description: `(Optional) Array of IDs of the user SSH keys which should be added to the device. If you omit this, SSH keys of all the members of the parent project will be added to the device. If you specify this array, only the listed user SSH keys (and any project_ssh_key_ids) will be added. User SSH keys can be created with the [equinix_metal_ssh_key](equinix_metal_ssh_key.md) resource`,
				},
				resource.Attribute{
					Name:        "reinstall",
					Description: `(Optional) Whether the device should be reinstalled instead of destroyed when modifying user_data, custom_data, or operating system. See [Reinstall](#reinstall) below for more details.`,
				},
				resource.Attribute{
					Name:        "storage",
					Description: `(Optional) JSON for custom partitioning. Only usable on reserved hardware. More information in in the [Custom Partitioning and RAID](https://metal.equinix.com/developers/docs/servers/custom-partitioning-raid/) doc. Please note that the disks.partitions.size attribute must be a string, not an integer. It can be a number string, or size notation string, e.g. "4G" or "8M" (for gigabytes and megabytes).`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags attached to the device.`,
				},
				resource.Attribute{
					Name:        "termination_time",
					Description: `(Optional) Timestamp for device termination. For example ` + "`" + `2021-09-03T16:32:00+03:00` + "`" + `. If you don't supply timezone info, timestamp is assumed to be in UTC.`,
				},
				resource.Attribute{
					Name:        "user_data",
					Description: `(Optional) A string of the desired User Data for the device. By default, changing this attribute will cause the provider to destroy and recreate your device. If ` + "`" + `reinstall` + "`" + ` is specified or ` + "`" + `behavior.allow_changes` + "`" + ` includes ` + "`" + `"user_data"` + "`" + `, the device will be updated in-place instead of recreated.`,
				},
				resource.Attribute{
					Name:        "wait_for_reservation_deprovision",
					Description: `(Optional) Only used for devices in reserved hardware. If set, the deletion of this device will block until the hardware reservation is marked provisionable (about 4 minutes in August 2019). ### Behavior The ` + "`" + `behavior` + "`" + ` block has below fields:`,
				},
				resource.Attribute{
					Name:        "allow_changes",
					Description: `(Optional) List of attributes that are allowed to change without recreating the instance. Supported attributes: ` + "`" + `custom_data` + "`" + `, ` + "`" + `user_data` + "`" + `" ### IP address The ` + "`" + `ip_address` + "`" + ` block has below fields:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) One of ` + "`" + `private_ipv4` + "`" + `, ` + "`" + `public_ipv4` + "`" + `, ` + "`" + `public_ipv6` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "cidr",
					Description: `(Optional) CIDR suffix for IP address block to be assigned, i.e. amount of addresses.`,
				},
				resource.Attribute{
					Name:        "reservation_ids",
					Description: `(Optional) List of UUIDs of [IP block reservations](metal_reserved_ip_block.md) from which the public IPv4 address should be taken. You can supply one ` + "`" + `ip_address` + "`" + ` block per IP address type. If you use the ` + "`" + `ip_address` + "`" + ` you must always pass a block for ` + "`" + `private_ipv4` + "`" + `. To learn more about using the reserved IP addresses for new devices, see the examples in the [equinix_metal_reserved_ip_block](metal_reserved_ip_block.md) documentation. ### Reinstall The ` + "`" + `reinstall` + "`" + ` block has below fields:`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) Whether the provider should favour reinstall over destroy and create. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "preserve_data",
					Description: `(Optional) Whether the non-OS disks should be kept or wiped during reinstall. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "deprovision_fast",
					Description: `(Optional) Whether the OS disk should be filled with ` + "`" + `00h` + "`" + ` bytes before reinstall. Defaults to ` + "`" + `false` + "`" + `. ### Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/configuration/resources#operation-timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Defaults to 20 mins) Used when creating the Device. This includes the time to provision the OS.`,
				},
				resource.Attribute{
					Name:        "update",
					Description: `(Defaults to 20 mins) Used when updating the Device. This includes the time needed to reprovision instances when ` + "`" + `reinstall` + "`" + ` arguments are used.`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Defaults to 20 mins) Used when deleting the Device. This includes the time to deprovision a hardware reservation when ` + "`" + `wait_for_reservation_deprovision` + "`" + ` is enabled. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "access_private_ipv4",
					Description: `The ipv4 private IP assigned to the device.`,
				},
				resource.Attribute{
					Name:        "access_public_ipv4",
					Description: `The ipv4 maintenance IP assigned to the device.`,
				},
				resource.Attribute{
					Name:        "access_public_ipv6",
					Description: `The ipv6 maintenance IP assigned to the device.`,
				},
				resource.Attribute{
					Name:        "billing_cycle",
					Description: `The billing cycle of the device (monthly or hourly).`,
				},
				resource.Attribute{
					Name:        "created",
					Description: `The timestamp for when the device was created.`,
				},
				resource.Attribute{
					Name:        "deployed_facility",
					Description: `(`,
				},
				resource.Attribute{
					Name:        "deployed_hardware_reservation_id",
					Description: `ID of hardware reservation where this device was deployed. It is useful when using the ` + "`" + `next-available` + "`" + ` hardware reservation.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description string for the device.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `The hostname of the device.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the device.`,
				},
				resource.Attribute{
					Name:        "locked",
					Description: `Whether the device is locked.`,
				},
				resource.Attribute{
					Name:        "metro",
					Description: `The metro area where the device is deployed.`,
				},
				resource.Attribute{
					Name:        "network",
					Description: `The device's private and public IP (v4 and v6) network details. See [Network Attribute](#network-attribute) below for more details.`,
				},
				resource.Attribute{
					Name:        "network_type",
					Description: `(Deprecated) Network type of a device, used in [Layer 2 networking](https://metal.equinix.com/developers/docs/networking/layer2/). Since this attribute is deprecated you should handle Network Type with one of [equinix_metal_port](equinix_metal_port.md), [equinix_metal_device_network_type](equinix_metal_device_network_type.md) resources or [equinix_metal_port](../data-sources/equinix_metal_port.md) datasource. See [network_types guide](../guides/network_types.md) for more info.`,
				},
				resource.Attribute{
					Name:        "operating_system",
					Description: `The operating system running on the device.`,
				},
				resource.Attribute{
					Name:        "plan",
					Description: `The hardware config of the device.`,
				},
				resource.Attribute{
					Name:        "ports",
					Description: `List of ports assigned to the device. See [Ports Attribute](#ports-attribute) below for more details.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `The ID of the project the device belongs to.`,
				},
				resource.Attribute{
					Name:        "root_password",
					Description: `Root password to the server (disabled after 24 hours).`,
				},
				resource.Attribute{
					Name:        "ssh_key_ids",
					Description: `List of IDs of SSH keys deployed in the device, can be both user and project SSH keys.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The status of the device.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags attached to the device.`,
				},
				resource.Attribute{
					Name:        "updated",
					Description: `The timestamp for the last time the device was updated. ### Network Attribute When a device is run without any special network, it will have 3 networks:`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `IPv4 or IPv6 address string.`,
				},
				resource.Attribute{
					Name:        "cidr",
					Description: `Bit length of the network mask of the address.`,
				},
				resource.Attribute{
					Name:        "gateway",
					Description: `Address of router.`,
				},
				resource.Attribute{
					Name:        "public",
					Description: `Whether the address is routable from the Internet.`,
				},
				resource.Attribute{
					Name:        "family",
					Description: `IP version. One of ` + "`" + `4` + "`" + `, ` + "`" + `6` + "`" + `. ### Ports Attribute Each element in the ` + "`" + `ports` + "`" + ` list exports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the port (e.g. ` + "`" + `eth0` + "`" + `, or ` + "`" + `bond0` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the port.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type of the port (e.g. ` + "`" + `NetworkPort` + "`" + ` or ` + "`" + `NetworkBondPort` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "mac",
					Description: `MAC address assigned to the port.`,
				},
				resource.Attribute{
					Name:        "bonded",
					Description: `Whether this port is part of a bond in bonded network setup. ## Import This resource can be imported using an existing device ID: ` + "`" + `` + "`" + `` + "`" + `sh terraform import equinix_metal_device {existing_device_id} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "access_private_ipv4",
					Description: `The ipv4 private IP assigned to the device.`,
				},
				resource.Attribute{
					Name:        "access_public_ipv4",
					Description: `The ipv4 maintenance IP assigned to the device.`,
				},
				resource.Attribute{
					Name:        "access_public_ipv6",
					Description: `The ipv6 maintenance IP assigned to the device.`,
				},
				resource.Attribute{
					Name:        "billing_cycle",
					Description: `The billing cycle of the device (monthly or hourly).`,
				},
				resource.Attribute{
					Name:        "created",
					Description: `The timestamp for when the device was created.`,
				},
				resource.Attribute{
					Name:        "deployed_facility",
					Description: `(`,
				},
				resource.Attribute{
					Name:        "deployed_hardware_reservation_id",
					Description: `ID of hardware reservation where this device was deployed. It is useful when using the ` + "`" + `next-available` + "`" + ` hardware reservation.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description string for the device.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `The hostname of the device.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the device.`,
				},
				resource.Attribute{
					Name:        "locked",
					Description: `Whether the device is locked.`,
				},
				resource.Attribute{
					Name:        "metro",
					Description: `The metro area where the device is deployed.`,
				},
				resource.Attribute{
					Name:        "network",
					Description: `The device's private and public IP (v4 and v6) network details. See [Network Attribute](#network-attribute) below for more details.`,
				},
				resource.Attribute{
					Name:        "network_type",
					Description: `(Deprecated) Network type of a device, used in [Layer 2 networking](https://metal.equinix.com/developers/docs/networking/layer2/). Since this attribute is deprecated you should handle Network Type with one of [equinix_metal_port](equinix_metal_port.md), [equinix_metal_device_network_type](equinix_metal_device_network_type.md) resources or [equinix_metal_port](../data-sources/equinix_metal_port.md) datasource. See [network_types guide](../guides/network_types.md) for more info.`,
				},
				resource.Attribute{
					Name:        "operating_system",
					Description: `The operating system running on the device.`,
				},
				resource.Attribute{
					Name:        "plan",
					Description: `The hardware config of the device.`,
				},
				resource.Attribute{
					Name:        "ports",
					Description: `List of ports assigned to the device. See [Ports Attribute](#ports-attribute) below for more details.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `The ID of the project the device belongs to.`,
				},
				resource.Attribute{
					Name:        "root_password",
					Description: `Root password to the server (disabled after 24 hours).`,
				},
				resource.Attribute{
					Name:        "ssh_key_ids",
					Description: `List of IDs of SSH keys deployed in the device, can be both user and project SSH keys.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The status of the device.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags attached to the device.`,
				},
				resource.Attribute{
					Name:        "updated",
					Description: `The timestamp for the last time the device was updated. ### Network Attribute When a device is run without any special network, it will have 3 networks:`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `IPv4 or IPv6 address string.`,
				},
				resource.Attribute{
					Name:        "cidr",
					Description: `Bit length of the network mask of the address.`,
				},
				resource.Attribute{
					Name:        "gateway",
					Description: `Address of router.`,
				},
				resource.Attribute{
					Name:        "public",
					Description: `Whether the address is routable from the Internet.`,
				},
				resource.Attribute{
					Name:        "family",
					Description: `IP version. One of ` + "`" + `4` + "`" + `, ` + "`" + `6` + "`" + `. ### Ports Attribute Each element in the ` + "`" + `ports` + "`" + ` list exports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the port (e.g. ` + "`" + `eth0` + "`" + `, or ` + "`" + `bond0` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the port.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type of the port (e.g. ` + "`" + `NetworkPort` + "`" + ` or ` + "`" + `NetworkBondPort` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "mac",
					Description: `MAC address assigned to the port.`,
				},
				resource.Attribute{
					Name:        "bonded",
					Description: `Whether this port is part of a bond in bonded network setup. ## Import This resource can be imported using an existing device ID: ` + "`" + `` + "`" + `` + "`" + `sh terraform import equinix_metal_device {existing_device_id} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "equinix_equinix_metal_device_network_type",
			Category:         "Metal",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"metal",
				"device",
				"network",
				"type",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "device_id",
					Description: `(Required) The ID of the device on which the network type should be set.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) Network type to set. Must be one of ` + "`" + `layer3` + "`" + `, ` + "`" + `hybrid` + "`" + `, ` + "`" + `hybrid-bonded` + "`" + `, ` + "`" + `layer2-individual` + "`" + ` and ` + "`" + `layer2-bonded` + "`" + `. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the controlled device. Use this in linked resources, if you need to wait for the network type change. It is the same as ` + "`" + `device_id` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the controlled device. Use this in linked resources, if you need to wait for the network type change. It is the same as ` + "`" + `device_id` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "equinix_equinix_metal_gateway",
			Category:         "Metal",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"metal",
				"gateway",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) UUID of the project where the gateway is scoped to.`,
				},
				resource.Attribute{
					Name:        "vlan_id",
					Description: `(Required) UUID of the VLAN where the gateway is scoped to.`,
				},
				resource.Attribute{
					Name:        "ip_reservation_id",
					Description: `(Optional) UUID of Public or VRF IP Reservation to associate with the gateway, the reservation must be in the same metro as the VLAN, conflicts with ` + "`" + `private_ipv4_subnet_size` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "private_ipv4_subnet_size",
					Description: `(Optional) Size of the private IPv4 subnet to create for this metal gateway, must be one of ` + "`" + `8` + "`" + `, ` + "`" + `16` + "`" + `, ` + "`" + `32` + "`" + `, ` + "`" + `64` + "`" + `, ` + "`" + `128` + "`" + `. Conflicts with ` + "`" + `ip_reservation_id` + "`" + `. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `Status of the gateway resource.`,
				},
				resource.Attribute{
					Name:        "vrf_id",
					Description: `UUID of the VRF associated with the IP Reservation`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "state",
					Description: `Status of the gateway resource.`,
				},
				resource.Attribute{
					Name:        "vrf_id",
					Description: `UUID of the VRF associated with the IP Reservation`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "equinix_equinix_metal_ip_attachment",
			Category:         "Metal",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"metal",
				"ip",
				"attachment",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "device_id",
					Description: `(Required) ID of device to which to assign the subnet.`,
				},
				resource.Attribute{
					Name:        "cidr_notation",
					Description: `(Required) CIDR notation of subnet from block reserved in the same project and metro as the device. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique ID of the assignment.`,
				},
				resource.Attribute{
					Name:        "device_id",
					Description: `ID of device to which subnet is assigned.`,
				},
				resource.Attribute{
					Name:        "cidr_notation",
					Description: `Assigned subnet in CIDR notation, e.g., ` + "`" + `147.229.15.30/31` + "`" + ``,
				},
				resource.Attribute{
					Name:        "gateway",
					Description: `IP address of gateway for the subnet.`,
				},
				resource.Attribute{
					Name:        "network",
					Description: `Subnet network address.`,
				},
				resource.Attribute{
					Name:        "netmask",
					Description: `Subnet mask in decimal notation, e.g., ` + "`" + `255.255.255.0` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "cidr",
					Description: `Length of CIDR prefix of the subnet as integer.`,
				},
				resource.Attribute{
					Name:        "address_family",
					Description: `Address family as integer. One of ` + "`" + `4` + "`" + ` or ` + "`" + `6` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "public",
					Description: `Boolean flag whether subnet is reachable from the Internet.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The unique ID of the assignment.`,
				},
				resource.Attribute{
					Name:        "device_id",
					Description: `ID of device to which subnet is assigned.`,
				},
				resource.Attribute{
					Name:        "cidr_notation",
					Description: `Assigned subnet in CIDR notation, e.g., ` + "`" + `147.229.15.30/31` + "`" + ``,
				},
				resource.Attribute{
					Name:        "gateway",
					Description: `IP address of gateway for the subnet.`,
				},
				resource.Attribute{
					Name:        "network",
					Description: `Subnet network address.`,
				},
				resource.Attribute{
					Name:        "netmask",
					Description: `Subnet mask in decimal notation, e.g., ` + "`" + `255.255.255.0` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "cidr",
					Description: `Length of CIDR prefix of the subnet as integer.`,
				},
				resource.Attribute{
					Name:        "address_family",
					Description: `Address family as integer. One of ` + "`" + `4` + "`" + ` or ` + "`" + `6` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "public",
					Description: `Boolean flag whether subnet is reachable from the Internet.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "equinix_equinix_metal_organization",
			Category:         "Metal",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"metal",
				"organization",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the Organization.`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `(Required) An object that has the address information. See [Address](#address) below for more details.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description string.`,
				},
				resource.Attribute{
					Name:        "website",
					Description: `(Optional) Website link.`,
				},
				resource.Attribute{
					Name:        "twitter",
					Description: `(Optional) Twitter handle.`,
				},
				resource.Attribute{
					Name:        "logo",
					Description: `(Optional) Logo URL. ### Address The ` + "`" + `address` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `(Required) Postal address.`,
				},
				resource.Attribute{
					Name:        "city",
					Description: `(Required) City name.`,
				},
				resource.Attribute{
					Name:        "country",
					Description: `(Required) Two letter country code (ISO 3166-1 alpha-2), e.g. US.`,
				},
				resource.Attribute{
					Name:        "zip_code",
					Description: `(Required) Zip Code.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `(Optional) State name. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique ID of the organization.`,
				},
				resource.Attribute{
					Name:        "created",
					Description: `The timestamp for when the organization was created.`,
				},
				resource.Attribute{
					Name:        "updated",
					Description: `The timestamp for the last time the organization was updated. ## Import This resource can be imported using an existing organization ID: ` + "`" + `` + "`" + `` + "`" + `sh terraform import equinix_metal_organization {existing_organization_id} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The unique ID of the organization.`,
				},
				resource.Attribute{
					Name:        "created",
					Description: `The timestamp for when the organization was created.`,
				},
				resource.Attribute{
					Name:        "updated",
					Description: `The timestamp for the last time the organization was updated. ## Import This resource can be imported using an existing organization ID: ` + "`" + `` + "`" + `` + "`" + `sh terraform import equinix_metal_organization {existing_organization_id} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "equinix_equinix_metal_organization_member",
			Category:         "Metal",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"metal",
				"organization",
				"member",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "invitee",
					Description: `(Required) The email address of the user to invite`,
				},
				resource.Attribute{
					Name:        "organization_id",
					Description: `(Required) The organization to invite the user to`,
				},
				resource.Attribute{
					Name:        "projects_ids",
					Description: `(Required) Project IDs the member has access to within the organization. If the member is an 'admin', the projects list should be empty.`,
				},
				resource.Attribute{
					Name:        "roles",
					Description: `(Required) Organization roles (admin, collaborator, limited_collaborator, billing)`,
				},
				resource.Attribute{
					Name:        "message",
					Description: `A message to include in the emailed invitation. ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique ID of the membership.`,
				},
				resource.Attribute{
					Name:        "nonce",
					Description: `The nonce for the invitation (only known in the invitation stage)`,
				},
				resource.Attribute{
					Name:        "invited_by",
					Description: `The user_id of the user that sent the invitation (only known in the invitation stage)`,
				},
				resource.Attribute{
					Name:        "created",
					Description: `When the invitation was created (only known in the invitation stage)`,
				},
				resource.Attribute{
					Name:        "updated",
					Description: `When the invitation was updated (only known in the invitation stage)`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the membership ('invited' when an invitation is open, 'active' when the user is an organization member) ## Import This resource can be imported using the ` + "`" + `invitee` + "`" + ` and ` + "`" + `organization_id` + "`" + ` as colon separated arguments: ` + "`" + `` + "`" + `` + "`" + `sh terraform import equinix_metal_organization_member.resource_name {invitee}:{organization_id} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The unique ID of the membership.`,
				},
				resource.Attribute{
					Name:        "nonce",
					Description: `The nonce for the invitation (only known in the invitation stage)`,
				},
				resource.Attribute{
					Name:        "invited_by",
					Description: `The user_id of the user that sent the invitation (only known in the invitation stage)`,
				},
				resource.Attribute{
					Name:        "created",
					Description: `When the invitation was created (only known in the invitation stage)`,
				},
				resource.Attribute{
					Name:        "updated",
					Description: `When the invitation was updated (only known in the invitation stage)`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the membership ('invited' when an invitation is open, 'active' when the user is an organization member) ## Import This resource can be imported using the ` + "`" + `invitee` + "`" + ` and ` + "`" + `organization_id` + "`" + ` as colon separated arguments: ` + "`" + `` + "`" + `` + "`" + `sh terraform import equinix_metal_organization_member.resource_name {invitee}:{organization_id} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "equinix_equinix_metal_port",
			Category:         "Metal",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"metal",
				"port",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "port_id",
					Description: `(Required) ID of the port to read.`,
				},
				resource.Attribute{
					Name:        "bonded",
					Description: `(Required) Whether the port should be bonded.`,
				},
				resource.Attribute{
					Name:        "layer2",
					Description: `(Optional) Whether to put the port to Layer 2 mode, valid only for bond ports.`,
				},
				resource.Attribute{
					Name:        "vlan_ids",
					Description: `(Optional) List of VLAN UUIDs to attach to the port, valid only for L2 and Hybrid ports.`,
				},
				resource.Attribute{
					Name:        "vxlan_ids",
					Description: `(Optional) List of VXLAN IDs to attach to the port, valid only for L2 and Hybrid ports.`,
				},
				resource.Attribute{
					Name:        "native_vlan_id",
					Description: `(Optional) UUID of a VLAN to assign as a native VLAN. It must be one of attached VLANs (from ` + "`" + `vlan_ids` + "`" + ` parameter).`,
				},
				resource.Attribute{
					Name:        "reset_on_delete",
					Description: `(Optional) Behavioral setting to reset the port to default settings (layer3 bonded mode without any vlan attached) before delete/destroy. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the port, e.g. ` + "`" + `bond0` + "`" + ` or ` + "`" + `eth0` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "network_type",
					Description: `One of layer2-bonded, layer2-individual, layer3, hybrid and hybrid-bonded. This attribute is only set on bond ports.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type is either "NetworkBondPort" for bond ports or "NetworkPort" for bondable ethernet ports.`,
				},
				resource.Attribute{
					Name:        "mac",
					Description: `MAC address of the port.`,
				},
				resource.Attribute{
					Name:        "bond_id",
					Description: `UUID of the bond port.`,
				},
				resource.Attribute{
					Name:        "bond_name",
					Description: `Name of the bond port.`,
				},
				resource.Attribute{
					Name:        "disbond_supported",
					Description: `Flag indicating whether the port can be removed from a bond.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `Name of the port, e.g. ` + "`" + `bond0` + "`" + ` or ` + "`" + `eth0` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "network_type",
					Description: `One of layer2-bonded, layer2-individual, layer3, hybrid and hybrid-bonded. This attribute is only set on bond ports.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type is either "NetworkBondPort" for bond ports or "NetworkPort" for bondable ethernet ports.`,
				},
				resource.Attribute{
					Name:        "mac",
					Description: `MAC address of the port.`,
				},
				resource.Attribute{
					Name:        "bond_id",
					Description: `UUID of the bond port.`,
				},
				resource.Attribute{
					Name:        "bond_name",
					Description: `Name of the bond port.`,
				},
				resource.Attribute{
					Name:        "disbond_supported",
					Description: `Flag indicating whether the port can be removed from a bond.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "equinix_equinix_metal_port_vlan_attachment",
			Category:         "Metal",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"metal",
				"port",
				"vlan",
				"attachment",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "device_id",
					Description: `(Required) ID of device to be assigned to the VLAN.`,
				},
				resource.Attribute{
					Name:        "port_name",
					Description: `(Required) Name of network port to be assigned to the VLAN.`,
				},
				resource.Attribute{
					Name:        "vlan_vnid",
					Description: `(Required) VXLAN Network Identifier.`,
				},
				resource.Attribute{
					Name:        "force_bond",
					Description: `(Optional) Add port back to the bond when this resource is removed. Default is ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "native",
					Description: `(Optional) Mark this VLAN a native VLAN on the port. This can be used only if this assignment assigns second or further VLAN to the port. To ensure that this attachment is not first on a port, you can use ` + "`" + `depends_on` + "`" + ` pointing to another ` + "`" + `equinix_metal_port_vlan_attachment` + "`" + `, just like in the layer2-individual example above. ## Attribute Referece In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `UUID of device port used in the assignment.`,
				},
				resource.Attribute{
					Name:        "vlan_id",
					Description: `UUID of VLAN API resource.`,
				},
				resource.Attribute{
					Name:        "port_id",
					Description: `UUID of device port.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "equinix_equinix_metal_project",
			Category:         "Metal",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"metal",
				"project",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the project.`,
				},
				resource.Attribute{
					Name:        "organization_id",
					Description: `(Required) The UUID of organization under which you want to create the project. If you leave it out, the project will be create under your the default organization of your account.`,
				},
				resource.Attribute{
					Name:        "payment_method_id",
					Description: `The UUID of payment method for this project. The payment method and the project need to belong to the same organization (passed with ` + "`" + `organization_id` + "`" + `, or default).`,
				},
				resource.Attribute{
					Name:        "backend_transfer",
					Description: `Enable or disable [Backend Transfer](https://metal.equinix.com/developers/docs/networking/backend-transfer/), default is ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "bgp_config",
					Description: `Optional BGP settings. Refer to [Equinix Metal guide for BGP](https://metal.equinix.com/developers/docs/networking/local-global-bgp/). ->`,
				},
				resource.Attribute{
					Name:        "asn",
					Description: `(Required) Autonomous System Number for local BGP deployment.`,
				},
				resource.Attribute{
					Name:        "deployment_type",
					Description: `(Required) ` + "`" + `local` + "`" + ` or ` + "`" + `global` + "`" + `, the ` + "`" + `local` + "`" + ` is likely to be usable immediately, the ` + "`" + `global` + "`" + ` will need to be reviewed by Equinix Metal engineers.`,
				},
				resource.Attribute{
					Name:        "md5",
					Description: `(Optional) Password for BGP session in plaintext (not a checksum). ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique ID of the project.`,
				},
				resource.Attribute{
					Name:        "created",
					Description: `The timestamp for when the project was created.`,
				},
				resource.Attribute{
					Name:        "updated",
					Description: `The timestamp for the last time the project was updated. The ` + "`" + `bgp_config` + "`" + ` block additionally exports:`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `status of BGP configuration in the project.`,
				},
				resource.Attribute{
					Name:        "max_prefix",
					Description: `The maximum number of route filters allowed per server. ## Import This resource can be imported using an existing project ID: ` + "`" + `` + "`" + `` + "`" + `sh terraform import equinix_metal_project {existing_project_id} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The unique ID of the project.`,
				},
				resource.Attribute{
					Name:        "created",
					Description: `The timestamp for when the project was created.`,
				},
				resource.Attribute{
					Name:        "updated",
					Description: `The timestamp for the last time the project was updated. The ` + "`" + `bgp_config` + "`" + ` block additionally exports:`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `status of BGP configuration in the project.`,
				},
				resource.Attribute{
					Name:        "max_prefix",
					Description: `The maximum number of route filters allowed per server. ## Import This resource can be imported using an existing project ID: ` + "`" + `` + "`" + `` + "`" + `sh terraform import equinix_metal_project {existing_project_id} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "equinix_equinix_metal_project_api_key",
			Category:         "Metal",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"metal",
				"project",
				"api",
				"key",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) UUID of the project where the API key is scoped to.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Required) Description string for the Project API Key resource.`,
				},
				resource.Attribute{
					Name:        "read-only",
					Description: `(Optional) Flag indicating whether the API key shoud be read-only. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "token",
					Description: `API token which can be used in Equinix Metal API clients`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "token",
					Description: `API token which can be used in Equinix Metal API clients`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "equinix_equinix_metal_project_ssh_key",
			Category:         "Metal",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"metal",
				"project",
				"ssh",
				"key",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the SSH key for identification.`,
				},
				resource.Attribute{
					Name:        "public_key",
					Description: `(Required) The public key. If this is a file, it can be read using the file interpolation function.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The ID of parent project. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique ID of the key.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the SSH key.`,
				},
				resource.Attribute{
					Name:        "owner_id",
					Description: `The ID of parent project (same as project_id).`,
				},
				resource.Attribute{
					Name:        "fingerprint",
					Description: `The fingerprint of the SSH key.`,
				},
				resource.Attribute{
					Name:        "created",
					Description: `The timestamp for when the SSH key was created.`,
				},
				resource.Attribute{
					Name:        "updated",
					Description: `The timestamp for the last time the SSH key was updated.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The unique ID of the key.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the SSH key.`,
				},
				resource.Attribute{
					Name:        "owner_id",
					Description: `The ID of parent project (same as project_id).`,
				},
				resource.Attribute{
					Name:        "fingerprint",
					Description: `The fingerprint of the SSH key.`,
				},
				resource.Attribute{
					Name:        "created",
					Description: `The timestamp for when the SSH key was created.`,
				},
				resource.Attribute{
					Name:        "updated",
					Description: `The timestamp for the last time the SSH key was updated.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "equinix_equinix_metal_reserved_ip_block",
			Category:         "Metal",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"metal",
				"reserved",
				"ip",
				"block",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The metal project ID where to allocate the address block.`,
				},
				resource.Attribute{
					Name:        "quantity",
					Description: `(Optional) The number of allocated ` + "`" + `/32` + "`" + ` addresses, a power of 2. Required when ` + "`" + `type` + "`" + ` is not ` + "`" + `vrf` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) One of ` + "`" + `global_ipv4` + "`" + `, ` + "`" + `public_ipv4` + "`" + `, or ` + "`" + `vrf` + "`" + `. Defaults to ` + "`" + `public_ipv4` + "`" + ` for backward compatibility.`,
				},
				resource.Attribute{
					Name:        "facility",
					Description: `(`,
				},
				resource.Attribute{
					Name:        "metro",
					Description: `(Optional) Metro where to allocate the public IP address block, makes sense only if type is ` + "`" + `public_ipv4` + "`" + ` and must be empty if type is ` + "`" + `global_ipv4` + "`" + `. Conflicts with ` + "`" + `facility` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Arbitrary description.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) String list of tags.`,
				},
				resource.Attribute{
					Name:        "vrf_id",
					Description: `(Optional) Only valid and required when ` + "`" + `type` + "`" + ` is ` + "`" + `vrf` + "`" + `. VRF ID for type=vrf reservations.`,
				},
				resource.Attribute{
					Name:        "wait_for_state",
					Description: `(Optional) Wait for the IP reservation block to reach a desired state on resource creation. One of: ` + "`" + `pending` + "`" + `, ` + "`" + `created` + "`" + `. The ` + "`" + `created` + "`" + ` state is default and recommended if the addresses are needed within the configuration. An error will be returned if a timeout or the ` + "`" + `denied` + "`" + ` state is encountered.`,
				},
				resource.Attribute{
					Name:        "custom_data",
					Description: `(Optional) Custom Data is an arbitrary object (submitted in Terraform as serialized JSON) to assign to the IP Reservation. This may be helpful for self-managed IPAM. The object must be valid JSON.`,
				},
				resource.Attribute{
					Name:        "network",
					Description: `(Optional) Only valid as an argument and required when ` + "`" + `type` + "`" + ` is ` + "`" + `vrf` + "`" + `. An unreserved network address from an existing ` + "`" + `ip_range` + "`" + ` in the specified VRF.`,
				},
				resource.Attribute{
					Name:        "cidr",
					Description: `(Optional) Only valid as an argument and required when ` + "`" + `type` + "`" + ` is ` + "`" + `vrf` + "`" + `. The size of the network to reserve from an existing VRF ip_range. ` + "`" + `cidr` + "`" + ` can only be specified with ` + "`" + `vrf_id` + "`" + `. Range is 22-31. Virtual Circuits require 30-31. Other VRF resources must use a CIDR in the 22-29 range. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique ID of the block.`,
				},
				resource.Attribute{
					Name:        "cidr_notation",
					Description: `Address and mask in CIDR notation, e.g. ` + "`" + `147.229.15.30/31` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "network",
					Description: `Network IP address portion of the block specification.`,
				},
				resource.Attribute{
					Name:        "netmask",
					Description: `Mask in decimal notation, e.g. ` + "`" + `255.255.255.0` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "cidr",
					Description: `length of CIDR prefix of the block as integer.`,
				},
				resource.Attribute{
					Name:        "address_family",
					Description: `Address family as integer. One of ` + "`" + `4` + "`" + ` or ` + "`" + `6` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "public",
					Description: `Boolean flag whether addresses from a block are public.`,
				},
				resource.Attribute{
					Name:        "global",
					Description: `Boolean flag whether addresses from a block are global (i.e. can be assigned in any metro).`,
				},
				resource.Attribute{
					Name:        "vrf_id",
					Description: `VRF ID of the block when type=vrf ->`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The unique ID of the block.`,
				},
				resource.Attribute{
					Name:        "cidr_notation",
					Description: `Address and mask in CIDR notation, e.g. ` + "`" + `147.229.15.30/31` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "network",
					Description: `Network IP address portion of the block specification.`,
				},
				resource.Attribute{
					Name:        "netmask",
					Description: `Mask in decimal notation, e.g. ` + "`" + `255.255.255.0` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "cidr",
					Description: `length of CIDR prefix of the block as integer.`,
				},
				resource.Attribute{
					Name:        "address_family",
					Description: `Address family as integer. One of ` + "`" + `4` + "`" + ` or ` + "`" + `6` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "public",
					Description: `Boolean flag whether addresses from a block are public.`,
				},
				resource.Attribute{
					Name:        "global",
					Description: `Boolean flag whether addresses from a block are global (i.e. can be assigned in any metro).`,
				},
				resource.Attribute{
					Name:        "vrf_id",
					Description: `VRF ID of the block when type=vrf ->`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "equinix_equinix_metal_spot_market_request",
			Category:         "Metal",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"metal",
				"spot",
				"market",
				"request",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "devices_max",
					Description: `(Required) Maximum number devices to be created.`,
				},
				resource.Attribute{
					Name:        "devices_min",
					Description: `(Required) Miniumum number devices to be created.`,
				},
				resource.Attribute{
					Name:        "max_bid_price",
					Description: `(Required) Maximum price user is willing to pay per hour per device.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) Project ID.`,
				},
				resource.Attribute{
					Name:        "wait_for_devices",
					Description: `(Optional) On resource creation wait until all desired devices are active. On resource destruction wait until devices are removed.`,
				},
				resource.Attribute{
					Name:        "facilities",
					Description: `(`,
				},
				resource.Attribute{
					Name:        "metro",
					Description: `(Optional) Metro where devices should be created.`,
				},
				resource.Attribute{
					Name:        "locked",
					Description: `(Optional) Blocks deletion of the SpotMarketRequest device until the lock is disabled.`,
				},
				resource.Attribute{
					Name:        "instance_parameters",
					Description: `(Required) Key/Value pairs of parameters for devices provisioned from this request. Valid keys are: ` + "`" + `billing_cycle` + "`" + `, ` + "`" + `plan` + "`" + `, ` + "`" + `operating_system` + "`" + `, ` + "`" + `hostname` + "`" + `, ` + "`" + `termintation_time` + "`" + `, ` + "`" + `always_pxe` + "`" + `, ` + "`" + `description` + "`" + `, ` + "`" + `features` + "`" + `, ` + "`" + `locked` + "`" + `, ` + "`" + `project_ssh_keys` + "`" + `, ` + "`" + `user_ssh_keys` + "`" + `, ` + "`" + `userdata` + "`" + `, ` + "`" + `customdata` + "`" + `, ` + "`" + `ipxe_script_url` + "`" + `, ` + "`" + `tags` + "`" + `. You can find each parameter description in [equinix_metal_device](equinix_metal_device.md) docs. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Spot Market Request. ### Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/configuration/resources#operation-timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Defaults to 60 mins) Used when creating the Spot Market Request and ` + "`" + `wait_for_devices` + "`" + ` is set to ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Defaults to 60 mins) Used when destroying the Spot Market Request and ` + "`" + `wait_for_devices` + "`" + ` is set to ` + "`" + `true` + "`" + `. ## Import This resource can be imported using an existing spot market request ID: ` + "`" + `` + "`" + `` + "`" + `sh terraform import equinix_metal_spot_market_request {existing_spot_market_request_id} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Spot Market Request. ### Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/configuration/resources#operation-timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Defaults to 60 mins) Used when creating the Spot Market Request and ` + "`" + `wait_for_devices` + "`" + ` is set to ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Defaults to 60 mins) Used when destroying the Spot Market Request and ` + "`" + `wait_for_devices` + "`" + ` is set to ` + "`" + `true` + "`" + `. ## Import This resource can be imported using an existing spot market request ID: ` + "`" + `` + "`" + `` + "`" + `sh terraform import equinix_metal_spot_market_request {existing_spot_market_request_id} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "equinix_equinix_metal_ssh_key",
			Category:         "Metal",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"metal",
				"ssh",
				"key",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the SSH key for identification`,
				},
				resource.Attribute{
					Name:        "public_key",
					Description: `(Required) The public key. If this is a file, it can be read using the file interpolation function ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique ID of the key.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the SSH key.`,
				},
				resource.Attribute{
					Name:        "public_key",
					Description: `The text of the public key.`,
				},
				resource.Attribute{
					Name:        "fingerprint",
					Description: `The fingerprint of the SSH key.`,
				},
				resource.Attribute{
					Name:        "owner_id",
					Description: `The UUID of the Equinix Metal API User who owns this key.`,
				},
				resource.Attribute{
					Name:        "created",
					Description: `The timestamp for when the SSH key was created.`,
				},
				resource.Attribute{
					Name:        "updated",
					Description: `The timestamp for the last time the SSH key was updated. ## Import This resource can be imported using an existing SSH Key ID: ` + "`" + `` + "`" + `` + "`" + `sh terraform import equinix_metal_ssh_key {existing_sshkey_id} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The unique ID of the key.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the SSH key.`,
				},
				resource.Attribute{
					Name:        "public_key",
					Description: `The text of the public key.`,
				},
				resource.Attribute{
					Name:        "fingerprint",
					Description: `The fingerprint of the SSH key.`,
				},
				resource.Attribute{
					Name:        "owner_id",
					Description: `The UUID of the Equinix Metal API User who owns this key.`,
				},
				resource.Attribute{
					Name:        "created",
					Description: `The timestamp for when the SSH key was created.`,
				},
				resource.Attribute{
					Name:        "updated",
					Description: `The timestamp for the last time the SSH key was updated. ## Import This resource can be imported using an existing SSH Key ID: ` + "`" + `` + "`" + `` + "`" + `sh terraform import equinix_metal_ssh_key {existing_sshkey_id} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "equinix_equinix_metal_user_api_key",
			Category:         "Metal",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"metal",
				"user",
				"api",
				"key",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `(Required) Description string for the User API Key resource.`,
				},
				resource.Attribute{
					Name:        "read-only",
					Description: `(Required) Flag indicating whether the API key shoud be read-only. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "user_id",
					Description: `UUID of the owner of the API key.`,
				},
				resource.Attribute{
					Name:        "token",
					Description: `API token which can be used in Equinix Metal API clients.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "user_id",
					Description: `UUID of the owner of the API key.`,
				},
				resource.Attribute{
					Name:        "token",
					Description: `API token which can be used in Equinix Metal API clients.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "equinix_equinix_metal_virtual_circuit",
			Category:         "Metal",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"metal",
				"virtual",
				"circuit",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "connection_id",
					Description: `(Required) UUID of Connection where the VC is scoped to.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) UUID of the Project where the VC is scoped to.`,
				},
				resource.Attribute{
					Name:        "port_id",
					Description: `(Required) UUID of the Connection Port where the VC is scoped to.`,
				},
				resource.Attribute{
					Name:        "nni_vlan",
					Description: `(Required) Equinix Metal network-to-network VLAN ID.`,
				},
				resource.Attribute{
					Name:        "vlan_id",
					Description: `(Required) UUID of the VLAN to associate.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the Virtual Circuit resource.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description for the Virtual Circuit resource.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags for the Virtual Circuit resource.`,
				},
				resource.Attribute{
					Name:        "speed",
					Description: `(Optional) Speed of the Virtual Circuit resource.`,
				},
				resource.Attribute{
					Name:        "vrf_id",
					Description: `(Optional) UUID of the VRF to associate.`,
				},
				resource.Attribute{
					Name:        "peer_asn",
					Description: `(Optional, required with ` + "`" + `vrf_id` + "`" + `) The BGP ASN of the peer. The same ASN may be the used across several VCs, but it cannot be the same as the local_asn of the VRF.`,
				},
				resource.Attribute{
					Name:        "subnet",
					Description: `(Optional, required with ` + "`" + `vrf_id` + "`" + `) A subnet from one of the IP blocks associated with the VRF that we will help create an IP reservation for. Can only be either a /30 or /31.`,
				},
				resource.Attribute{
					Name:        "metal_ip",
					Description: `(Optional, required with ` + "`" + `vrf_id` + "`" + `) The Metal IP address for the SVI (Switch Virtual Interface) of the VirtualCircuit. Will default to the first usable IP in the subnet.`,
				},
				resource.Attribute{
					Name:        "customer_ip",
					Description: `(Optional, required with ` + "`" + `vrf_id` + "`" + `) The Customer IP address which the CSR switch will peer with. Will default to the other usable IP in the subnet.`,
				},
				resource.Attribute{
					Name:        "md5",
					Description: `(Optional, only valid with ` + "`" + `vrf_id` + "`" + `) The password that can be set for the VRF BGP peer ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the virtal circuit.`,
				},
				resource.Attribute{
					Name:        "vnid",
					Description: `VNID VLAN parameter, see the [documentation for Equinix Fabric](https://metal.equinix.com/developers/docs/networking/fabric/).`,
				},
				resource.Attribute{
					Name:        "nni_vnid",
					Description: `NNI VLAN parameters, see the [documentation for Equinix Fabric](https://metal.equinix.com/developers/docs/networking/fabric/). ## Import This resource can be imported using an existing Virtual Circuit ID: ` + "`" + `` + "`" + `` + "`" + `sh terraform import equinix_metal_virtual_circuit {existing_id} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "status",
					Description: `Status of the virtal circuit.`,
				},
				resource.Attribute{
					Name:        "vnid",
					Description: `VNID VLAN parameter, see the [documentation for Equinix Fabric](https://metal.equinix.com/developers/docs/networking/fabric/).`,
				},
				resource.Attribute{
					Name:        "nni_vnid",
					Description: `NNI VLAN parameters, see the [documentation for Equinix Fabric](https://metal.equinix.com/developers/docs/networking/fabric/). ## Import This resource can be imported using an existing Virtual Circuit ID: ` + "`" + `` + "`" + `` + "`" + `sh terraform import equinix_metal_virtual_circuit {existing_id} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "equinix_equinix_metal_vlan",
			Category:         "Metal",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"metal",
				"vlan",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) ID of parent project.`,
				},
				resource.Attribute{
					Name:        "metro",
					Description: `(Optional) Metro in which to create the VLAN`,
				},
				resource.Attribute{
					Name:        "facility",
					Description: `(`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description string.`,
				},
				resource.Attribute{
					Name:        "vxlan",
					Description: `(Optional) VLAN ID, must be unique in metro. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the virtual network. ## Import This resource can be imported using an existing VLAN ID (UUID): ` + "`" + `` + "`" + `` + "`" + `sh terraform import equinix_metal_vlan {existing_vlan_id} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the virtual network. ## Import This resource can be imported using an existing VLAN ID (UUID): ` + "`" + `` + "`" + `` + "`" + `sh terraform import equinix_metal_vlan {existing_vlan_id} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "equinix_equinix_metal_vrf",
			Category:         "Metal",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"metal",
				"vrf",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) User-supplied name of the VRF, unique to the project`,
				},
				resource.Attribute{
					Name:        "metro",
					Description: `(Required) Metro ID or Code where the VRF will be deployed.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) Project ID where the VRF will be deployed.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the VRF.`,
				},
				resource.Attribute{
					Name:        "local_asn",
					Description: `(Optional) The 4-byte ASN set on the VRF.`,
				},
				resource.Attribute{
					Name:        "ip_ranges",
					Description: `(Optional) All IPv4 and IPv6 Ranges that will be available to BGP Peers. IPv4 addresses must be /8 or smaller with a minimum size of /29. IPv6 must be /56 or smaller with a minimum size of /64. Ranges must not overlap other ranges within the VRF. ## Attributes Reference No additional attributes are exported. ## Import This resource can be imported using an existing VRF ID: ` + "`" + `` + "`" + `` + "`" + `sh terraform import equinix_metal_vrf {existing_id} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "equinix_equinix_network_acl_template",
			Category:         "Network Edge",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"network",
				"edge",
				"acl",
				"template",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) ACL template name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) ACL template description, up to 200 characters.`,
				},
				resource.Attribute{
					Name:        "metro_code",
					Description: `(Deprecated) ACL template location metro code.`,
				},
				resource.Attribute{
					Name:        "inbound_rule",
					Description: `(Required) One or more rules to specify allowed inbound traffic. Rules are ordered, matching traffic rule stops processing subsequent ones. The ` + "`" + `inbound_rule` + "`" + ` block has below fields:`,
				},
				resource.Attribute{
					Name:        "subnets",
					Description: `(Deprecated) Inbound traffic source IP subnets in CIDR format.`,
				},
				resource.Attribute{
					Name:        "subnet",
					Description: `(Required) Inbound traffic source IP subnet in CIDR format.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Required) Inbound traffic protocol. One of ` + "`" + `IP` + "`" + `, ` + "`" + `TCP` + "`" + `, ` + "`" + `UDP` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "src_port",
					Description: `(Required) Inbound traffic source ports. Allowed values are a comma separated list of ports, e.g., ` + "`" + `20,22,23` + "`" + `, port range, e.g., ` + "`" + `1023-1040` + "`" + ` or word ` + "`" + `any` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "dst_port",
					Description: `(Required) Inbound traffic destination ports. Allowed values are a comma separated list of ports, e.g., ` + "`" + `20,22,23` + "`" + `, port range, e.g., ` + "`" + `1023-1040` + "`" + ` or word ` + "`" + `any` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Inbound rule description, up to 200 characters. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `Unique identifier of ACL template resource.`,
				},
				resource.Attribute{
					Name:        "device_id",
					Description: `(Deprecated) Identifier of a network device where template was applied.`,
				},
				resource.Attribute{
					Name:        "device_acl_status",
					Description: `Status of ACL template provisioning process, where template was applied. One of ` + "`" + `PROVISIONING` + "`" + `, ` + "`" + `PROVISIONED` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "device_details",
					Description: `List of the devices where the ACL template is applied. The ` + "`" + `device_details` + "`" + ` block has below fields:`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `Device uuid.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Device name.`,
				},
				resource.Attribute{
					Name:        "acl_status",
					Description: `Device ACL provisioning status where template was applied. One of ` + "`" + `PROVISIONING` + "`" + `, ` + "`" + `PROVISIONED` + "`" + `. ## Import This resource can be imported using an existing ID: ` + "`" + `` + "`" + `` + "`" + `sh terraform import equinix_network_acl_template.example {existing_id} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "uuid",
					Description: `Unique identifier of ACL template resource.`,
				},
				resource.Attribute{
					Name:        "device_id",
					Description: `(Deprecated) Identifier of a network device where template was applied.`,
				},
				resource.Attribute{
					Name:        "device_acl_status",
					Description: `Status of ACL template provisioning process, where template was applied. One of ` + "`" + `PROVISIONING` + "`" + `, ` + "`" + `PROVISIONED` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "device_details",
					Description: `List of the devices where the ACL template is applied. The ` + "`" + `device_details` + "`" + ` block has below fields:`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `Device uuid.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Device name.`,
				},
				resource.Attribute{
					Name:        "acl_status",
					Description: `Device ACL provisioning status where template was applied. One of ` + "`" + `PROVISIONING` + "`" + `, ` + "`" + `PROVISIONED` + "`" + `. ## Import This resource can be imported using an existing ID: ` + "`" + `` + "`" + `` + "`" + `sh terraform import equinix_network_acl_template.example {existing_id} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "equinix_equinix_network_bgp",
			Category:         "Network Edge",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"network",
				"edge",
				"bgp",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "connection_id",
					Description: `(Required) identifier of a connection established between. network device and remote service provider that will be used for peering.`,
				},
				resource.Attribute{
					Name:        "local_ip_address",
					Description: `(Required) IP address in CIDR format of a local device.`,
				},
				resource.Attribute{
					Name:        "local_asn",
					Description: `(Required) Local ASN number.`,
				},
				resource.Attribute{
					Name:        "remote_ip_address",
					Description: `(Required) IP address of remote peer.`,
				},
				resource.Attribute{
					Name:        "remote_asn",
					Description: `(Required) Remote ASN number.`,
				},
				resource.Attribute{
					Name:        "authentication_key",
					Description: `(Optional) shared key used for BGP peer authentication. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `BGP peering configuration unique identifier.`,
				},
				resource.Attribute{
					Name:        "device_id",
					Description: `unique identifier of a network device that is a local peer in a given BGP peering configuration.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `BGP peer state, one of ` + "`" + `Idle` + "`" + `, ` + "`" + `Connect` + "`" + `, ` + "`" + `Active` + "`" + `, ` + "`" + `OpenSent` + "`" + `, ` + "`" + `OpenConfirm` + "`" + `, ` + "`" + `Established` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "provisioning_status",
					Description: `BGP peering configuration provisioning status, one of ` + "`" + `PROVISIONING` + "`" + `, ` + "`" + `PENDING_UPDATE` + "`" + `, ` + "`" + `PROVISIONED` + "`" + `, ` + "`" + `FAILED` + "`" + `. ## Import This resource can be imported using an existing ID: ` + "`" + `` + "`" + `` + "`" + `sh terraform import equinix_network_bgp.example {existing_id} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "uuid",
					Description: `BGP peering configuration unique identifier.`,
				},
				resource.Attribute{
					Name:        "device_id",
					Description: `unique identifier of a network device that is a local peer in a given BGP peering configuration.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `BGP peer state, one of ` + "`" + `Idle` + "`" + `, ` + "`" + `Connect` + "`" + `, ` + "`" + `Active` + "`" + `, ` + "`" + `OpenSent` + "`" + `, ` + "`" + `OpenConfirm` + "`" + `, ` + "`" + `Established` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "provisioning_status",
					Description: `BGP peering configuration provisioning status, one of ` + "`" + `PROVISIONING` + "`" + `, ` + "`" + `PENDING_UPDATE` + "`" + `, ` + "`" + `PROVISIONED` + "`" + `, ` + "`" + `FAILED` + "`" + `. ## Import This resource can be imported using an existing ID: ` + "`" + `` + "`" + `` + "`" + `sh terraform import equinix_network_bgp.example {existing_id} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "equinix_equinix_network_device",
			Category:         "Network Edge",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"network",
				"edge",
				"device",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Device name.`,
				},
				resource.Attribute{
					Name:        "type_code",
					Description: `(Required) Device type code.`,
				},
				resource.Attribute{
					Name:        "metro_code",
					Description: `(Required) Device location metro code.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `(Optional) Device hostname prefix.`,
				},
				resource.Attribute{
					Name:        "package_code",
					Description: `(Required) Device software package code.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Required) Device software software version.`,
				},
				resource.Attribute{
					Name:        "core_count",
					Description: `(Required) Number of CPU cores used by device.`,
				},
				resource.Attribute{
					Name:        "term_length",
					Description: `(Required) Device term length.`,
				},
				resource.Attribute{
					Name:        "self_managed",
					Description: `(Optional) Boolean value that determines device management mode, i.e., ` + "`" + `self-managed` + "`" + ` or ` + "`" + `Equinix-managed` + "`" + ` (default).`,
				},
				resource.Attribute{
					Name:        "byol",
					Description: `(Optional) Boolean value that determines device licensing mode, i.e., ` + "`" + `bring your own license` + "`" + ` or ` + "`" + `subscription` + "`" + ` (default).`,
				},
				resource.Attribute{
					Name:        "license_token",
					Description: `(Optional, conflicts with ` + "`" + `license_file` + "`" + `) License Token applicable for some device types in BYOL licensing mode.`,
				},
				resource.Attribute{
					Name:        "license_file",
					Description: `(Optional) Path to the license file that will be uploaded and applied on a device. Applicable for some device types in BYOL licensing mode.`,
				},
				resource.Attribute{
					Name:        "license_file_id",
					Description: `(Optional, conflicts with ` + "`" + `license_file` + "`" + `) Identifier of a license file that will be applied on the device.`,
				},
				resource.Attribute{
					Name:        "cloud_init_file_id",
					Description: `(Optional) Identifier of a cloud init file that will be applied on the device.`,
				},
				resource.Attribute{
					Name:        "throughput",
					Description: `(Optional) Device license throughput.`,
				},
				resource.Attribute{
					Name:        "throughput_unit",
					Description: `(Optional) License throughput unit. One of ` + "`" + `Mbps` + "`" + ` or ` + "`" + `Gbps` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "account_number",
					Description: `(Required) Billing account number for a device.`,
				},
				resource.Attribute{
					Name:        "notifications",
					Description: `(Required) List of email addresses that will receive device status notifications.`,
				},
				resource.Attribute{
					Name:        "purchase_order_number",
					Description: `(Optional) Purchase order number associated with a device order.`,
				},
				resource.Attribute{
					Name:        "order_reference",
					Description: `(Optional) Name/number used to identify device order on the invoice.`,
				},
				resource.Attribute{
					Name:        "acl_template_id",
					Description: `(Optional) Identifier of a WAN interface ACL template that will be applied on the device.`,
				},
				resource.Attribute{
					Name:        "mgmt_acl_template_uuid",
					Description: `(Optional) Identifier of an MGMT interface ACL template that will be applied on the device.`,
				},
				resource.Attribute{
					Name:        "additional_bandwidth",
					Description: `(Optional) Additional Internet bandwidth, in Mbps, that will be allocated to the device (in addition to default 15Mbps).`,
				},
				resource.Attribute{
					Name:        "interface_count",
					Description: `(Optional) Number of network interfaces on a device. If not specified, default number for a given device type will be used.`,
				},
				resource.Attribute{
					Name:        "wan_interafce_id",
					Description: `(Optional) Specify the WAN/SSH interface id. If not specified, default WAN/SSH interface for a given device type will be used.`,
				},
				resource.Attribute{
					Name:        "vendor_configuration",
					Description: `(Optional) Map of vendor specific configuration parameters for a device (controller1, activationKey, managementType, siteId, systemIpAddress)`,
				},
				resource.Attribute{
					Name:        "ssh-key",
					Description: `(Optional) Definition of SSH key that will be provisioned on a device (max one key). See [SSH Key](#ssh-key) below for more details.`,
				},
				resource.Attribute{
					Name:        "secondary_device",
					Description: `(Optional) Definition of secondary device for redundant device configurations. See [Secondary Device](#secondary-device) below for more details.`,
				},
				resource.Attribute{
					Name:        "cluster_details",
					Description: `(Optional) An object that has the cluster details. See [Cluster Details](#cluster-details) below for more details. ### Secondary Device ->`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Secondary device name.`,
				},
				resource.Attribute{
					Name:        "metro_code",
					Description: `(Required) Metro location of a secondary device.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `(Optional) Secondary device hostname.`,
				},
				resource.Attribute{
					Name:        "license_token",
					Description: `(Optional, conflicts with ` + "`" + `license_file` + "`" + `) License Token can be provided for some device types o the device.`,
				},
				resource.Attribute{
					Name:        "license_file",
					Description: `(Optional) Path to the license file that will be uploaded and applied on a secondary device. Applicable for some device types in BYOL licensing mode.`,
				},
				resource.Attribute{
					Name:        "license_file_id",
					Description: `(Optional, conflicts with ` + "`" + `license_file` + "`" + `) Identifier of a license file that will be applied on a secondary device.`,
				},
				resource.Attribute{
					Name:        "cloud_init_file_id",
					Description: `(Optional) Identifier of a cloud init file that will be applied on a secondary device.`,
				},
				resource.Attribute{
					Name:        "account_number",
					Description: `(Required) Billing account number for secondary device.`,
				},
				resource.Attribute{
					Name:        "notifications",
					Description: `(Required) List of email addresses that will receive notifications about secondary device.`,
				},
				resource.Attribute{
					Name:        "additional_bandwidth",
					Description: `(Optional) Additional Internet bandwidth, in Mbps, for a secondary device.`,
				},
				resource.Attribute{
					Name:        "vendor_configuration",
					Description: `(Optional) Key/Value pairs of vendor specific configuration parameters for a secondary device. Key values are ` + "`" + `controller1` + "`" + `, ` + "`" + `activationKey` + "`" + `, ` + "`" + `managementType` + "`" + `, ` + "`" + `siteId` + "`" + `, ` + "`" + `systemIpAddress` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "acl_template_id",
					Description: `(Optional) Identifier of a WAN interface ACL template that will be applied on a secondary device.`,
				},
				resource.Attribute{
					Name:        "mgmt_acl_template_uuid",
					Description: `(Optional) Identifier of an MGMT interface ACL template that will be applied on a secondary device.`,
				},
				resource.Attribute{
					Name:        "ssh-key",
					Description: `(Optional) Up to one definition of SSH key that will be provisioned on a secondary device. ### SSH Key The ` + "`" + `ssh_key` + "`" + ` block supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Required) username associated with given key.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) reference by name to previously provisioned public SSH key. ### Cluster Details ->`,
				},
				resource.Attribute{
					Name:        "cluster_name",
					Description: `(Required) The name of the cluster device`,
				},
				resource.Attribute{
					Name:        "node0",
					Description: `(Required) An object that has ` + "`" + `node0` + "`" + ` configuration. See [Cluster Details - Nodes](#cluster-details---nodes) below for more details.`,
				},
				resource.Attribute{
					Name:        "node1",
					Description: `(Required) An object that has ` + "`" + `node1` + "`" + ` configuration. See [Cluster Details - Nodes](#cluster-details---nodes) below for more details. ### Cluster Details - Nodes The ` + "`" + `node0` + "`" + ` and ` + "`" + `node1` + "`" + ` blocks supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "vendor_configuration",
					Description: `(Optional) An object that has fields relevant to the vendor of the cluster device. See [Cluster Details - Nodes - Vendor Configuration](#cluster-details---nodes---vendor-configuration) below for more details.`,
				},
				resource.Attribute{
					Name:        "license_file_id",
					Description: `(Optional) License file id. This is necessary for Fortinet and Juniper clusters.`,
				},
				resource.Attribute{
					Name:        "license_token",
					Description: `(Optional) License token. This is necessary for Palo Alto clusters. ### Cluster Details - Nodes - Vendor Configuration The ` + "`" + `vendor_configuration` + "`" + ` block supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `(Optional) Hostname. This is necessary for Palo Alto, Juniper, and Fortinet clusters.`,
				},
				resource.Attribute{
					Name:        "admin_password",
					Description: `(Optional) The administrative password of the device. You can use it to log in to the console. This field is not available for all device types.`,
				},
				resource.Attribute{
					Name:        "controller1",
					Description: `(Optional) System IP Address. Mandatory for the Fortinet SDWAN cluster device.`,
				},
				resource.Attribute{
					Name:        "activation_key",
					Description: `(Optional) Activation key. This is required for Velocloud clusters.`,
				},
				resource.Attribute{
					Name:        "controller_fqdn",
					Description: `(Optional) Controller fqdn. This is required for Velocloud clusters.`,
				},
				resource.Attribute{
					Name:        "root_password",
					Description: `(Optional) The CLI password of the device. This field is relevant only for the Velocloud SDWAN cluster. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `Device unique identifier.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Device provisioning status. Possible values are ` + "`" + `INITIALIZING` + "`" + `, ` + "`" + `PROVISIONING` + "`" + `, ` + "`" + `WAITING_FOR_PRIMARY` + "`" + `, ` + "`" + `WAITING_FOR_SECONDARY` + "`" + `, ` + "`" + `WAITING_FOR_REPLICA_CLUSTER_NODES` + "`" + `, ` + "`" + `CLUSTER_SETUP_IN_PROGRESS` + "`" + `, ` + "`" + `FAILED` + "`" + `, ` + "`" + `PROVISIONED` + "`" + `, ` + "`" + `DEPROVISIONING` + "`" + `, ` + "`" + `DEPROVISIONED` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "license_status",
					Description: `Device license registration status. Possible values are ` + "`" + `APPLYING_LICENSE` + "`" + `, ` + "`" + `REGISTERED` + "`" + `, ` + "`" + `APPLIED` + "`" + `, ` + "`" + `WAITING_FOR_CLUSTER_SETUP` + "`" + `, ` + "`" + `REGISTRATION_FAILED` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "license_file_id",
					Description: `Unique identifier of applied license file.`,
				},
				resource.Attribute{
					Name:        "ibx",
					Description: `Device location Equinix Business Exchange name.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `Device location region.`,
				},
				resource.Attribute{
					Name:        "acl_template_id",
					Description: `Unique identifier of applied ACL template.`,
				},
				resource.Attribute{
					Name:        "ssh_ip_address",
					Description: `IP address of SSH enabled interface on the device.`,
				},
				resource.Attribute{
					Name:        "ssh_ip_fqdn",
					Description: `FQDN of SSH enabled interface on the device.`,
				},
				resource.Attribute{
					Name:        "redundancy_type",
					Description: `Device redundancy type applicable for HA devices, either primary or secondary.`,
				},
				resource.Attribute{
					Name:        "redundant_id",
					Description: `Unique identifier for a redundant device applicable for HA devices.`,
				},
				resource.Attribute{
					Name:        "interface",
					Description: `List of device interfaces. See [Interface Attribute](#interface-attribute) below for more details.`,
				},
				resource.Attribute{
					Name:        "asn",
					Description: `(Autonomous System Number) Unique identifier for a network on the internet.`,
				},
				resource.Attribute{
					Name:        "zone_code",
					Description: `Device location zone code.`,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `The ID of the cluster.`,
				},
				resource.Attribute{
					Name:        "num_of_nodes",
					Description: `The number of nodes in the cluster. ### Interface Attribute Each interface attribute has below fields:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `interface identifier.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `interface name.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `interface status. One of ` + "`" + `AVAILABLE` + "`" + `, ` + "`" + `RESERVED` + "`" + `, ` + "`" + `ASSIGNED` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "operational_status",
					Description: `interface operational status. One of ` + "`" + `up` + "`" + `, ` + "`" + `down` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "mac_address",
					Description: `interface MAC address.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `interface IP address.`,
				},
				resource.Attribute{
					Name:        "assigned_type",
					Description: `interface management type (Equinix Managed or empty).`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `interface type. ## Timeouts This resource provides the following [Timeouts configuration](https://www.terraform.io/language/resources/syntax#operation-timeouts) options:`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "uuid",
					Description: `Device unique identifier.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Device provisioning status. Possible values are ` + "`" + `INITIALIZING` + "`" + `, ` + "`" + `PROVISIONING` + "`" + `, ` + "`" + `WAITING_FOR_PRIMARY` + "`" + `, ` + "`" + `WAITING_FOR_SECONDARY` + "`" + `, ` + "`" + `WAITING_FOR_REPLICA_CLUSTER_NODES` + "`" + `, ` + "`" + `CLUSTER_SETUP_IN_PROGRESS` + "`" + `, ` + "`" + `FAILED` + "`" + `, ` + "`" + `PROVISIONED` + "`" + `, ` + "`" + `DEPROVISIONING` + "`" + `, ` + "`" + `DEPROVISIONED` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "license_status",
					Description: `Device license registration status. Possible values are ` + "`" + `APPLYING_LICENSE` + "`" + `, ` + "`" + `REGISTERED` + "`" + `, ` + "`" + `APPLIED` + "`" + `, ` + "`" + `WAITING_FOR_CLUSTER_SETUP` + "`" + `, ` + "`" + `REGISTRATION_FAILED` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "license_file_id",
					Description: `Unique identifier of applied license file.`,
				},
				resource.Attribute{
					Name:        "ibx",
					Description: `Device location Equinix Business Exchange name.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `Device location region.`,
				},
				resource.Attribute{
					Name:        "acl_template_id",
					Description: `Unique identifier of applied ACL template.`,
				},
				resource.Attribute{
					Name:        "ssh_ip_address",
					Description: `IP address of SSH enabled interface on the device.`,
				},
				resource.Attribute{
					Name:        "ssh_ip_fqdn",
					Description: `FQDN of SSH enabled interface on the device.`,
				},
				resource.Attribute{
					Name:        "redundancy_type",
					Description: `Device redundancy type applicable for HA devices, either primary or secondary.`,
				},
				resource.Attribute{
					Name:        "redundant_id",
					Description: `Unique identifier for a redundant device applicable for HA devices.`,
				},
				resource.Attribute{
					Name:        "interface",
					Description: `List of device interfaces. See [Interface Attribute](#interface-attribute) below for more details.`,
				},
				resource.Attribute{
					Name:        "asn",
					Description: `(Autonomous System Number) Unique identifier for a network on the internet.`,
				},
				resource.Attribute{
					Name:        "zone_code",
					Description: `Device location zone code.`,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `The ID of the cluster.`,
				},
				resource.Attribute{
					Name:        "num_of_nodes",
					Description: `The number of nodes in the cluster. ### Interface Attribute Each interface attribute has below fields:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `interface identifier.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `interface name.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `interface status. One of ` + "`" + `AVAILABLE` + "`" + `, ` + "`" + `RESERVED` + "`" + `, ` + "`" + `ASSIGNED` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "operational_status",
					Description: `interface operational status. One of ` + "`" + `up` + "`" + `, ` + "`" + `down` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "mac_address",
					Description: `interface MAC address.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `interface IP address.`,
				},
				resource.Attribute{
					Name:        "assigned_type",
					Description: `interface management type (Equinix Managed or empty).`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `interface type. ## Timeouts This resource provides the following [Timeouts configuration](https://www.terraform.io/language/resources/syntax#operation-timeouts) options:`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "equinix_equinix_network_device_link",
			Category:         "Network Edge",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"network",
				"edge",
				"device",
				"link",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) device link name.`,
				},
				resource.Attribute{
					Name:        "subnet",
					Description: `(Optional) device link subnet in CIDR format. Not required for link between self configured devices.`,
				},
				resource.Attribute{
					Name:        "device",
					Description: `(Required) definition of one or more devices belonging to the device link. See [Device](#device) section below for more details.`,
				},
				resource.Attribute{
					Name:        "link",
					Description: `(Optional) definition of one or more, inter metro, connections belonging to the device link. See [Link](#link) section below for more details. ### Device The ` + "`" + `device` + "`" + ` block supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Required) Device identifier.`,
				},
				resource.Attribute{
					Name:        "asn",
					Description: `(Optional) Device ASN number. Not required for self configured devices.`,
				},
				resource.Attribute{
					Name:        "interface_id",
					Description: `(Optional) Device network interface identifier to use for device link connection. ### Link The ` + "`" + `link` + "`" + ` block supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "account_number",
					Description: `(Required) billing account number to be used for connection charges`,
				},
				resource.Attribute{
					Name:        "throughput",
					Description: `(Required) connection throughput.`,
				},
				resource.Attribute{
					Name:        "throughput_unit",
					Description: `(Required) connection throughput unit (Mbps or Gbps).`,
				},
				resource.Attribute{
					Name:        "src_metro_code",
					Description: `(Required) connection source metro code.`,
				},
				resource.Attribute{
					Name:        "dst_metro_code",
					Description: `(Required) connection destination metro code.`,
				},
				resource.Attribute{
					Name:        "src_zone_code",
					Description: `(Deprecated) connection source zone code is not required.`,
				},
				resource.Attribute{
					Name:        "dst_zone_code",
					Description: `(Deprecated) connection destination zone code is not required. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `Device link unique identifier.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Device link provisioning status. One of ` + "`" + `PROVISIONING` + "`" + `, ` + "`" + `PROVISIONED` + "`" + `, ` + "`" + `DEPROVISIONING` + "`" + `, ` + "`" + `DEPROVISIONED` + "`" + `, ` + "`" + `FAILED` + "`" + `. The ` + "`" + `device` + "`" + ` block attributes:`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `IP address from device link subnet that was assigned to the device`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `device link provisioning status on a given device. One of ` + "`" + `PROVISIONING` + "`" + `, ` + "`" + `PROVISIONED` + "`" + `, ` + "`" + `DEPROVISIONING` + "`" + `, ` + "`" + `DEPROVISIONED` + "`" + `, ` + "`" + `FAILED` + "`" + `. ## Timeouts This resource provides the following [Timeouts configuration](https://www.terraform.io/language/resources/syntax#operation-timeouts) options:`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "uuid",
					Description: `Device link unique identifier.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Device link provisioning status. One of ` + "`" + `PROVISIONING` + "`" + `, ` + "`" + `PROVISIONED` + "`" + `, ` + "`" + `DEPROVISIONING` + "`" + `, ` + "`" + `DEPROVISIONED` + "`" + `, ` + "`" + `FAILED` + "`" + `. The ` + "`" + `device` + "`" + ` block attributes:`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `IP address from device link subnet that was assigned to the device`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `device link provisioning status on a given device. One of ` + "`" + `PROVISIONING` + "`" + `, ` + "`" + `PROVISIONED` + "`" + `, ` + "`" + `DEPROVISIONING` + "`" + `, ` + "`" + `DEPROVISIONED` + "`" + `, ` + "`" + `FAILED` + "`" + `. ## Timeouts This resource provides the following [Timeouts configuration](https://www.terraform.io/language/resources/syntax#operation-timeouts) options:`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "equinix_equinix_network_file",
			Category:         "Network Edge",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"network",
				"edge",
				"file",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "file_name",
					Description: `(Required) File name.`,
				},
				resource.Attribute{
					Name:        "content",
					Description: `(Required) Uploaded file content, expected to be a UTF-8 encoded string.`,
				},
				resource.Attribute{
					Name:        "metro_code",
					Description: `(Required) File upload location metro code. It should match the device location metro code.`,
				},
				resource.Attribute{
					Name:        "type_code",
					Description: `(Required) Device type code.`,
				},
				resource.Attribute{
					Name:        "process_type",
					Description: `(Required) File process type (LICENSE or CLOUD_INIT).`,
				},
				resource.Attribute{
					Name:        "self_managed",
					Description: `(Required) Boolean value that determines device management mode, i.e., ` + "`" + `self-managed` + "`" + ` or ` + "`" + `Equinix-managed` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "byol",
					Description: `(Required) Boolean value that determines device licensing mode, i.e., ` + "`" + `bring your own license` + "`" + ` or ` + "`" + `subscription` + "`" + `. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `Unique identifier of file resource.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `File upload status. ## Import This resource can be imported using an existing ID: ` + "`" + `` + "`" + `` + "`" + `sh terraform import equinix_network_file.example {existing_id} ` + "`" + `` + "`" + `` + "`" + ` The ` + "`" + `content` + "`" + `, ` + "`" + `self_managed` + "`" + ` and ` + "`" + `byol` + "`" + ` fields can not be imported.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "uuid",
					Description: `Unique identifier of file resource.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `File upload status. ## Import This resource can be imported using an existing ID: ` + "`" + `` + "`" + `` + "`" + `sh terraform import equinix_network_file.example {existing_id} ` + "`" + `` + "`" + `` + "`" + ` The ` + "`" + `content` + "`" + `, ` + "`" + `self_managed` + "`" + ` and ` + "`" + `byol` + "`" + ` fields can not be imported.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "equinix_equinix_network_ssh_key",
			Category:         "Network Edge",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"network",
				"edge",
				"ssh",
				"key",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of SSH key used for identification.`,
				},
				resource.Attribute{
					Name:        "public_key",
					Description: `(Required) The SSH public key. If this is a file, it can be read using the file interpolation function.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) The type of SSH key: ` + "`" + `RSA` + "`" + ` (default) or ` + "`" + `DSA` + "`" + `. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `The unique identifier of the key ## Import This resource can be imported using an existing ID: ` + "`" + `` + "`" + `` + "`" + `sh terraform import equinix_network_ssh_key.example {existing_id} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "uuid",
					Description: `The unique identifier of the key ## Import This resource can be imported using an existing ID: ` + "`" + `` + "`" + `` + "`" + `sh terraform import equinix_network_ssh_key.example {existing_id} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "equinix_equinix_network_ssh_user",
			Category:         "Network Edge",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"network",
				"edge",
				"ssh",
				"user",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "username",
					Description: `(Required) SSH user login name.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Required) SSH user password.`,
				},
				resource.Attribute{
					Name:        "device_ids",
					Description: `(Required) list of device identifiers to which user will have access. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `SSH user unique identifier. ## Import This resource can be imported using an existing ID: ` + "`" + `` + "`" + `` + "`" + `sh terraform import equinix_network_ssh_user.example {existing_id} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "uuid",
					Description: `SSH user unique identifier. ## Import This resource can be imported using an existing ID: ` + "`" + `` + "`" + `` + "`" + `sh terraform import equinix_network_ssh_user.example {existing_id} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
	}

	resourcesMap = map[string]int{

		"equinix_equinix_ecx_l2_connection":          0,
		"equinix_equinix_ecx_l2_connection_accepter": 1,
		"equinix_equinix_ecx_l2_serviceprofile":      2,
		"equinix_equinix_fabric_connection":          3,
		"equinix_equinix_fabric_service_profile":     4,
		"equinix_equinix_metal_bgp_session":          5,
		"equinix_equinix_metal_connection":           6,
		"equinix_equinix_metal_device":               7,
		"equinix_equinix_metal_device_network_type":  8,
		"equinix_equinix_metal_gateway":              9,
		"equinix_equinix_metal_ip_attachment":        10,
		"equinix_equinix_metal_organization":         11,
		"equinix_equinix_metal_organization_member":  12,
		"equinix_equinix_metal_port":                 13,
		"equinix_equinix_metal_port_vlan_attachment": 14,
		"equinix_equinix_metal_project":              15,
		"equinix_equinix_metal_project_api_key":      16,
		"equinix_equinix_metal_project_ssh_key":      17,
		"equinix_equinix_metal_reserved_ip_block":    18,
		"equinix_equinix_metal_spot_market_request":  19,
		"equinix_equinix_metal_ssh_key":              20,
		"equinix_equinix_metal_user_api_key":         21,
		"equinix_equinix_metal_virtual_circuit":      22,
		"equinix_equinix_metal_vlan":                 23,
		"equinix_equinix_metal_vrf":                  24,
		"equinix_equinix_network_acl_template":       25,
		"equinix_equinix_network_bgp":                26,
		"equinix_equinix_network_device":             27,
		"equinix_equinix_network_device_link":        28,
		"equinix_equinix_network_file":               29,
		"equinix_equinix_network_ssh_key":            30,
		"equinix_equinix_network_ssh_user":           31,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
