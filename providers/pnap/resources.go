package pnap

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "pnap_ip_block",
			Category:         "Resources",
			ShortDescription: `Provides a phoenixNAP IP Block resource. This can be used to create, modify and delete IP Blocks.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "location",
					Description: `(Required) IP Block location ID. Currently this field should be set to ` + "`" + `PHX` + "`" + `, ` + "`" + `ASH` + "`" + `, ` + "`" + `SGP` + "`" + `, ` + "`" + `NLD` + "`" + `, ` + "`" + `CHI` + "`" + `, ` + "`" + `SEA` + "`" + ` or ` + "`" + `AUS` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "cidr_block_size",
					Description: `(Required) CIDR IP Block Size. Currently this field should be set to either ` + "`" + `/31` + "`" + `, ` + "`" + `/30` + "`" + `, ` + "`" + `/29` + "`" + ` or ` + "`" + `/28` + "`" + `. For a larger Block Size contact support.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the IP Block.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags to set to IP Block, if any.`,
				},
				resource.Attribute{
					Name:        "tag_assignment",
					Description: `Tag request to assign to the IP Block.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the tag.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag assigned to the IP Block. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `IP Block identifier.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `IP Block location ID.`,
				},
				resource.Attribute{
					Name:        "cidr_block_size",
					Description: `CIDR IP Block Size.`,
				},
				resource.Attribute{
					Name:        "cidr",
					Description: `The IP Block in CIDR notation.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the IP Block.`,
				},
				resource.Attribute{
					Name:        "assigned_resource_id",
					Description: `ID of the resource assigned to the IP Block.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the IP Block.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `The tags assigned to the IP Block.`,
				},
				resource.Attribute{
					Name:        "tag_assignment",
					Description: `Tag assigned to the IP Block.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique id of the tag.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the tag.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag assigned to the IP Block.`,
				},
				resource.Attribute{
					Name:        "is_billing_tag",
					Description: `Whether or not to show the tag as part of billing and invoices.`,
				},
				resource.Attribute{
					Name:        "created_by",
					Description: `Who the tag was created by.`,
				},
				resource.Attribute{
					Name:        "is_bring_your_own",
					Description: `True if the IP Block is a "bring your own" block.`,
				},
				resource.Attribute{
					Name:        "created_on",
					Description: `Date and time when the IP Block was created.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `IP Block identifier.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `IP Block location ID.`,
				},
				resource.Attribute{
					Name:        "cidr_block_size",
					Description: `CIDR IP Block Size.`,
				},
				resource.Attribute{
					Name:        "cidr",
					Description: `The IP Block in CIDR notation.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the IP Block.`,
				},
				resource.Attribute{
					Name:        "assigned_resource_id",
					Description: `ID of the resource assigned to the IP Block.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the IP Block.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `The tags assigned to the IP Block.`,
				},
				resource.Attribute{
					Name:        "tag_assignment",
					Description: `Tag assigned to the IP Block.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique id of the tag.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the tag.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag assigned to the IP Block.`,
				},
				resource.Attribute{
					Name:        "is_billing_tag",
					Description: `Whether or not to show the tag as part of billing and invoices.`,
				},
				resource.Attribute{
					Name:        "created_by",
					Description: `Who the tag was created by.`,
				},
				resource.Attribute{
					Name:        "is_bring_your_own",
					Description: `True if the IP Block is a "bring your own" block.`,
				},
				resource.Attribute{
					Name:        "created_on",
					Description: `Date and time when the IP Block was created.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "pnap_private_network",
			Category:         "Resources",
			ShortDescription: `Provides a phoenixNAP Private Network resource. This can be used to create, modify, and delete private networks.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The friendly name of this private network. This name should be unique.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of this private network.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `(Required) The location of this private network. Supported values are ` + "`" + `PHX` + "`" + `, ` + "`" + `ASH` + "`" + `, ` + "`" + `SGP` + "`" + `, ` + "`" + `NLD` + "`" + `, ` + "`" + `CHI` + "`" + `, ` + "`" + `SEA` + "`" + ` and ` + "`" + `AUS` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "location_default",
					Description: `Identifies network as the default private network for the specified location. Default value is ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "cidr",
					Description: `(Required) IP range associated with this private network in CIDR notation. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The private network identifier.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The friendly name of this private network. This name should be unique.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of this private network.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The location of this private network.`,
				},
				resource.Attribute{
					Name:        "location_default",
					Description: `Identifies network as the default private network for the specified location. Default value is ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "cidr",
					Description: `IP range associated with this private network in CIDR notation.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of the private network.`,
				},
				resource.Attribute{
					Name:        "servers",
					Description: `(Deprecated) List of server details linked to the private network.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The server identifier.`,
				},
				resource.Attribute{
					Name:        "ips",
					Description: `List of private IPs associated to the server.`,
				},
				resource.Attribute{
					Name:        "memberships",
					Description: `A list of resources that are members of this private network.`,
				},
				resource.Attribute{
					Name:        "resource_id",
					Description: `The resource identifier.`,
				},
				resource.Attribute{
					Name:        "resource_type",
					Description: `The resource's type.`,
				},
				resource.Attribute{
					Name:        "ips",
					Description: `List of public IPs associated to the resource.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the private network.`,
				},
				resource.Attribute{
					Name:        "created_on",
					Description: `Date and time when this private network was created.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The private network identifier.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The friendly name of this private network. This name should be unique.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of this private network.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The location of this private network.`,
				},
				resource.Attribute{
					Name:        "location_default",
					Description: `Identifies network as the default private network for the specified location. Default value is ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "cidr",
					Description: `IP range associated with this private network in CIDR notation.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of the private network.`,
				},
				resource.Attribute{
					Name:        "servers",
					Description: `(Deprecated) List of server details linked to the private network.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The server identifier.`,
				},
				resource.Attribute{
					Name:        "ips",
					Description: `List of private IPs associated to the server.`,
				},
				resource.Attribute{
					Name:        "memberships",
					Description: `A list of resources that are members of this private network.`,
				},
				resource.Attribute{
					Name:        "resource_id",
					Description: `The resource identifier.`,
				},
				resource.Attribute{
					Name:        "resource_type",
					Description: `The resource's type.`,
				},
				resource.Attribute{
					Name:        "ips",
					Description: `List of public IPs associated to the resource.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the private network.`,
				},
				resource.Attribute{
					Name:        "created_on",
					Description: `Date and time when this private network was created.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "pnap_public_network",
			Category:         "Resources",
			ShortDescription: `Provides a phoenixNAP Public Network resource. This can be used to create, modify, and delete public networks.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The friendly name of this public network. This name should be unique.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of this public network.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `(Required) The location of this public network. Supported values are ` + "`" + `PHX` + "`" + `, ` + "`" + `ASH` + "`" + `, ` + "`" + `SGP` + "`" + `, ` + "`" + `NLD` + "`" + `, ` + "`" + `CHI` + "`" + `, ` + "`" + `SEA` + "`" + ` and ` + "`" + `AUS` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ip_blocks",
					Description: `A list of IP Blocks that will be associated with this public network (10 items at most).`,
				},
				resource.Attribute{
					Name:        "public_network_ip_block",
					Description: `The assigned IP Block to the public network.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The IP Block identifier. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The public network identifier.`,
				},
				resource.Attribute{
					Name:        "memberships",
					Description: `A list of resources that are members of this public network.`,
				},
				resource.Attribute{
					Name:        "resource_id",
					Description: `The resource identifier.`,
				},
				resource.Attribute{
					Name:        "resource_type",
					Description: `The resource's type.`,
				},
				resource.Attribute{
					Name:        "ips",
					Description: `List of public IPs associated to the resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The friendly name of this public network.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The location of this public network.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of this public network.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the public network.`,
				},
				resource.Attribute{
					Name:        "created_on",
					Description: `Date and time when this public network was created.`,
				},
				resource.Attribute{
					Name:        "ip_blocks",
					Description: `A list of IP Blocks that are associated with this public network.`,
				},
				resource.Attribute{
					Name:        "public_network_ip_block",
					Description: `The assigned IP Block to the public network.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The IP Block identifier.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The public network identifier.`,
				},
				resource.Attribute{
					Name:        "memberships",
					Description: `A list of resources that are members of this public network.`,
				},
				resource.Attribute{
					Name:        "resource_id",
					Description: `The resource identifier.`,
				},
				resource.Attribute{
					Name:        "resource_type",
					Description: `The resource's type.`,
				},
				resource.Attribute{
					Name:        "ips",
					Description: `List of public IPs associated to the resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The friendly name of this public network.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The location of this public network.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of this public network.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the public network.`,
				},
				resource.Attribute{
					Name:        "created_on",
					Description: `Date and time when this public network was created.`,
				},
				resource.Attribute{
					Name:        "ip_blocks",
					Description: `A list of IP Blocks that are associated with this public network.`,
				},
				resource.Attribute{
					Name:        "public_network_ip_block",
					Description: `The assigned IP Block to the public network.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The IP Block identifier.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "pnap_rancher_cluster",
			Category:         "Resources",
			ShortDescription: `Provides a phoenixNAP Rancher Cluster resource. This can be used to create and delete Rancher Server deployments.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `Cluster (Rancher Cluster) name. This field is autogenerated if not provided.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Cluster description.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `(Required) Deployment location. Cannot be changed once the cluster is created. For a full list of allowed locations visit [API docs](https://developers.phoenixnap.com/docs/rancher/1)`,
				},
				resource.Attribute{
					Name:        "node_pools",
					Description: `The node pools associated with the cluster (must contain exactly one item). The ` + "`" + `node_pools` + "`" + ` block has 4 fields.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the node pool.`,
				},
				resource.Attribute{
					Name:        "node_count",
					Description: `Number of configured nodes. Currently only node counts of 1 and 3 are possible.`,
				},
				resource.Attribute{
					Name:        "server_type",
					Description: `Node server type. Default value is "s0.d1.small". For a full list of allowed values visit [API docs](https://developers.phoenixnap.com/docs/rancher/1)`,
				},
				resource.Attribute{
					Name:        "ssh_config",
					Description: `(Write-only) Configuration defining which public SSH keys are pre-installed as authorized on the server. The ` + "`" + `ssh_config` + "`" + ` block has 3 fields.`,
				},
				resource.Attribute{
					Name:        "install_default_keys",
					Description: `Define whether public keys marked as default should be installed on this node. Default value is true.`,
				},
				resource.Attribute{
					Name:        "keys",
					Description: `List of public SSH keys.`,
				},
				resource.Attribute{
					Name:        "key_ids",
					Description: `List of public SSH key identifiers.`,
				},
				resource.Attribute{
					Name:        "configuration",
					Description: `(Write-only) Rancher configuration parameters. The ` + "`" + `configuration` + "`" + ` block has 7 fields.`,
				},
				resource.Attribute{
					Name:        "token",
					Description: `Shared secret used to join a server or agent to a cluster.`,
				},
				resource.Attribute{
					Name:        "tls_san",
					Description: `This maps to rancher's tls-san. Add additional hostname or IP as a Subject Alternative Name in the TLS cert.`,
				},
				resource.Attribute{
					Name:        "etcd_snapshot_schedule_cron",
					Description: `This maps to rancher's etcd-snapshot-schedule-cron. Snapshot interval time in cron spec.`,
				},
				resource.Attribute{
					Name:        "etcd_snapshot_retention",
					Description: `This maps to rancher's etcd-snapshot-retention. Number of snapshots to retain. Default value is 5.`,
				},
				resource.Attribute{
					Name:        "node_taint",
					Description: `This maps to rancher's node-taint. Registering kubelet with set of taints.`,
				},
				resource.Attribute{
					Name:        "cluster_domain",
					Description: `This maps to rancher's cluster-domain. Cluster Domain.`,
				},
				resource.Attribute{
					Name:        "certificates",
					Description: `Define the custom SSL certificates to be used instead of defaults. The ` + "`" + `certificates` + "`" + ` block has 3 fields.`,
				},
				resource.Attribute{
					Name:        "ca_certificate",
					Description: `The SSL CA certificate to be used for rancher admin.`,
				},
				resource.Attribute{
					Name:        "certificate",
					Description: `The SSL certificate to be used for rancher admin.`,
				},
				resource.Attribute{
					Name:        "certificate_key",
					Description: `The SSL certificate key to be used for rancher admin.`,
				},
				resource.Attribute{
					Name:        "workload_configuration",
					Description: `(Write-only) Workload cluster configuration parameters. The ` + "`" + `workload_configuration` + "`" + ` block has 4 fields.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the workload cluster. This field is autogenerated if not provided.`,
				},
				resource.Attribute{
					Name:        "server_count",
					Description: `Number of configured servers. Currently only server counts of 1 and 3 are possible. Default value is 1.`,
				},
				resource.Attribute{
					Name:        "server_type",
					Description: `(Required) Node server type. Cannot be changed once the cluster is created. Default value is "s0.d1.small". For a full list of allowed values visit [API docs](https://developers.phoenixnap.com/docs/rancher/1)`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `(Required) Workload cluster location. Cannot be changed once the cluster is created. For a full list of allowed locations visit [API docs](https://developers.phoenixnap.com/docs/rancher/1) ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The cluster identifier.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Cluster name. This field is autogenerated if not provided.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Cluster description.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `Deployment location.`,
				},
				resource.Attribute{
					Name:        "initial_cluster_version",
					Description: `The Rancher version that was installed on the cluster during the first creation process.`,
				},
				resource.Attribute{
					Name:        "node_pools",
					Description: `The node pools associated with the cluster (must contain exactly one item).`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the node pool.`,
				},
				resource.Attribute{
					Name:        "node_count",
					Description: `Number of configured nodes.`,
				},
				resource.Attribute{
					Name:        "server_type",
					Description: `Node server type. Default value is "s0.d1.small".`,
				},
				resource.Attribute{
					Name:        "ssh_config",
					Description: `Configuration defining which public SSH keys are pre-installed as authorized on the server.`,
				},
				resource.Attribute{
					Name:        "install_default_keys",
					Description: `Define whether public keys marked as default should be installed on this node. Default value is true.`,
				},
				resource.Attribute{
					Name:        "keys",
					Description: `List of public SSH keys.`,
				},
				resource.Attribute{
					Name:        "key_ids",
					Description: `List of public SSH key identifiers.`,
				},
				resource.Attribute{
					Name:        "nodes",
					Description: `The nodes associated with this node pool.`,
				},
				resource.Attribute{
					Name:        "server_id",
					Description: `The server identifier.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `Connection parameters to use to connect to the Rancher Server Administrative GUI.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `The Rancher Server URL.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `The username to use to login to the Rancher Server. This field is returned only as a response to the create cluster request. Make sure to take note or you will not be able to access the server.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `This is the password to be used to login to the Rancher Server. This field is returned only as a response to the create cluster request. Make sure to take note or you will not be able to access the server.`,
				},
				resource.Attribute{
					Name:        "status_description",
					Description: `The cluster status.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The cluster identifier.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Cluster name. This field is autogenerated if not provided.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Cluster description.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `Deployment location.`,
				},
				resource.Attribute{
					Name:        "initial_cluster_version",
					Description: `The Rancher version that was installed on the cluster during the first creation process.`,
				},
				resource.Attribute{
					Name:        "node_pools",
					Description: `The node pools associated with the cluster (must contain exactly one item).`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the node pool.`,
				},
				resource.Attribute{
					Name:        "node_count",
					Description: `Number of configured nodes.`,
				},
				resource.Attribute{
					Name:        "server_type",
					Description: `Node server type. Default value is "s0.d1.small".`,
				},
				resource.Attribute{
					Name:        "ssh_config",
					Description: `Configuration defining which public SSH keys are pre-installed as authorized on the server.`,
				},
				resource.Attribute{
					Name:        "install_default_keys",
					Description: `Define whether public keys marked as default should be installed on this node. Default value is true.`,
				},
				resource.Attribute{
					Name:        "keys",
					Description: `List of public SSH keys.`,
				},
				resource.Attribute{
					Name:        "key_ids",
					Description: `List of public SSH key identifiers.`,
				},
				resource.Attribute{
					Name:        "nodes",
					Description: `The nodes associated with this node pool.`,
				},
				resource.Attribute{
					Name:        "server_id",
					Description: `The server identifier.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `Connection parameters to use to connect to the Rancher Server Administrative GUI.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `The Rancher Server URL.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `The username to use to login to the Rancher Server. This field is returned only as a response to the create cluster request. Make sure to take note or you will not be able to access the server.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `This is the password to be used to login to the Rancher Server. This field is returned only as a response to the create cluster request. Make sure to take note or you will not be able to access the server.`,
				},
				resource.Attribute{
					Name:        "status_description",
					Description: `The cluster status.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "pnap_reservation",
			Category:         "Resources",
			ShortDescription: `Provides a phoenixNAP reservation resource. This can be used to create and modify reservations.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "sku",
					Description: `(Required) The SKU code of product pricing plan.`,
				},
				resource.Attribute{
					Name:        "auto_renew",
					Description: `A flag indicating whether the reservation will auto-renew (default is true, it can only be modified after the creation of resource).`,
				},
				resource.Attribute{
					Name:        "auto_renew_disable_reason",
					Description: `The reason for disabling auto-renewal. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The reservation identifier.`,
				},
				resource.Attribute{
					Name:        "product_code",
					Description: `The code identifying the product. This code has significance across all locations.`,
				},
				resource.Attribute{
					Name:        "product_category",
					Description: `The product category.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The location code.`,
				},
				resource.Attribute{
					Name:        "reservation_model",
					Description: `The reservation model.`,
				},
				resource.Attribute{
					Name:        "initial_invoice_model",
					Description: `Reservations created with initial invoice model ON_CREATION will be invoiced on same date when reservation is created. Reservation created with CALENDAR_MONTH initial invoice model will be invoiced at the begining of next month.`,
				},
				resource.Attribute{
					Name:        "start_date_time",
					Description: `The point in time (in UTC) when the reservation starts.`,
				},
				resource.Attribute{
					Name:        "end_date_time",
					Description: `The point in time (in UTC) when the reservation ends.`,
				},
				resource.Attribute{
					Name:        "last_renewal_date_time",
					Description: `The point in time (in UTC) when the reservation was renewed last.`,
				},
				resource.Attribute{
					Name:        "next_renewal_date_time",
					Description: `The point in time (in UTC) when the reservation will be renewed if auto renew is set to true.`,
				},
				resource.Attribute{
					Name:        "auto_renew",
					Description: `A flag indicating whether the reservation will auto-renew (default is true, it can only be modified after the creation of resource).`,
				},
				resource.Attribute{
					Name:        "sku",
					Description: `The SKU that will be applied to this reservation.`,
				},
				resource.Attribute{
					Name:        "price",
					Description: `Reservation price.`,
				},
				resource.Attribute{
					Name:        "price_unit",
					Description: `The unit to which the price applies.`,
				},
				resource.Attribute{
					Name:        "assigned_resource_id",
					Description: `The resource ID currently being assigned to reservation.`,
				},
				resource.Attribute{
					Name:        "next_billing_date",
					Description: `Next billing date for reservation.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The reservation identifier.`,
				},
				resource.Attribute{
					Name:        "product_code",
					Description: `The code identifying the product. This code has significance across all locations.`,
				},
				resource.Attribute{
					Name:        "product_category",
					Description: `The product category.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The location code.`,
				},
				resource.Attribute{
					Name:        "reservation_model",
					Description: `The reservation model.`,
				},
				resource.Attribute{
					Name:        "initial_invoice_model",
					Description: `Reservations created with initial invoice model ON_CREATION will be invoiced on same date when reservation is created. Reservation created with CALENDAR_MONTH initial invoice model will be invoiced at the begining of next month.`,
				},
				resource.Attribute{
					Name:        "start_date_time",
					Description: `The point in time (in UTC) when the reservation starts.`,
				},
				resource.Attribute{
					Name:        "end_date_time",
					Description: `The point in time (in UTC) when the reservation ends.`,
				},
				resource.Attribute{
					Name:        "last_renewal_date_time",
					Description: `The point in time (in UTC) when the reservation was renewed last.`,
				},
				resource.Attribute{
					Name:        "next_renewal_date_time",
					Description: `The point in time (in UTC) when the reservation will be renewed if auto renew is set to true.`,
				},
				resource.Attribute{
					Name:        "auto_renew",
					Description: `A flag indicating whether the reservation will auto-renew (default is true, it can only be modified after the creation of resource).`,
				},
				resource.Attribute{
					Name:        "sku",
					Description: `The SKU that will be applied to this reservation.`,
				},
				resource.Attribute{
					Name:        "price",
					Description: `Reservation price.`,
				},
				resource.Attribute{
					Name:        "price_unit",
					Description: `The unit to which the price applies.`,
				},
				resource.Attribute{
					Name:        "assigned_resource_id",
					Description: `The resource ID currently being assigned to reservation.`,
				},
				resource.Attribute{
					Name:        "next_billing_date",
					Description: `Next billing date for reservation.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "pnap_server",
			Category:         "Resources",
			ShortDescription: `Provides a phoenixNAP server resource. This can be used to create, modify, and delete servers.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "hostname",
					Description: `(Required) Server hostname.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Server description.`,
				},
				resource.Attribute{
					Name:        "os",
					Description: `(Required) The server’s OS ID used when the server was created (e.g., ubuntu/bionic, centos/centos7). For a full list of available operating systems visit [API docs](https://developers.phoenixnap.com/docs/bmc/1).`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) Server type ID. Cannot be changed once a server is created (e.g., s1.c1.small, s1.c1.medium). For a full list of available types visit [API docs](https://developers.phoenixnap.com/docs/bmc/1).`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `(Required) Server Location ID. Cannot be changed once a server is created (e.g., PHX). For a full list of available locations visit [API docs](https://developers.phoenixnap.com/docs/bmc/1)`,
				},
				resource.Attribute{
					Name:        "installDefaultSshKeys",
					Description: `Whether or not to install SSH keys marked as default in addition to any SSH keys specified in this request.`,
				},
				resource.Attribute{
					Name:        "ssh_keys",
					Description: `A list of SSH Keys that will be installed on the server.`,
				},
				resource.Attribute{
					Name:        "ssh_key_ids",
					Description: `A list of SSH key IDs that will be installed on the server in addition to any SSH keys specified in this request.`,
				},
				resource.Attribute{
					Name:        "reservation_id",
					Description: `Server reservation ID.`,
				},
				resource.Attribute{
					Name:        "pricing_model",
					Description: `Server pricing model. Currently this field should be set to HOURLY, ONE_MONTH_RESERVATION, TWELVE_MONTHS_RESERVATION, TWENTY_FOUR_MONTHS_RESERVATION or THIRTY_SIX_MONTHS_RESERVATION.`,
				},
				resource.Attribute{
					Name:        "network_type",
					Description: `The type of network configuration for this server. Currently this field should be set to PUBLIC_AND_PRIVATE or PRIVATE_ONLY.`,
				},
				resource.Attribute{
					Name:        "rdp_allowed_ips",
					Description: `List of IPs allowed for RDP access to Windows OS. Supported in single IP, CIDR and range format. When undefined, RDP is disabled. To allow RDP access from any IP use 0.0.0.0/0. Must contain at least 1 item.`,
				},
				resource.Attribute{
					Name:        "management_access_allowed_ips",
					Description: `Define list of IPs allowed to access the Management UI. Supported in single IP, CIDR and range format. When undefined, Management UI is disabled.Must contain at least 1 item.`,
				},
				resource.Attribute{
					Name:        "install_os_to_ram",
					Description: `If true, OS will be installed to and booted from the server's RAM. On restart RAM OS will be lost and the server will not be reachable unless a custom bootable OS has been deployed. Only supported for ubuntu/focal. Default value is ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "cloud_init",
					Description: `Cloud-init configuration details.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags to set to server, if any.`,
				},
				resource.Attribute{
					Name:        "network_configuration",
					Description: `Entire network details of bare metal server. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `Action to perform on server. Allowed actions are: reboot, reset (deprecated), powered-on, powered-off, shutdown.`,
				},
				resource.Attribute{
					Name:        "force",
					Description: `Query parameter controlling advanced features availability. Currently applicable for networking. It is advised to use with caution since it might lead to unhealthy setups. The ` + "`" + `cloud_init` + "`" + ` block has one field:`,
				},
				resource.Attribute{
					Name:        "user_data",
					Description: `User data for the [cloud-init](https://cloudinit.readthedocs.io/en/latest/) configuration in base64 encoding. NoCloud format is supported. Follow the [instructions](https://phoenixnap.com/kb/bmc-cloud-init) on how to provision a server using cloud-init. Only ubuntu/bionic and ubuntu/focal are supported. The ` + "`" + `tags` + "`" + ` block has field ` + "`" + `tag_assignment` + "`" + `. The ` + "`" + `tag_assignment` + "`" + ` block has 2 fields:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the tag.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag assigned to the IP Block. The ` + "`" + `network_configuration` + "`" + ` block has 4 fields: ` + "`" + `gateway_address` + "`" + `, ` + "`" + `private_network_configuration` + "`" + `, ` + "`" + `ip_blocks_configuration` + "`" + ` and ` + "`" + `public_network_configuration` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "gateway_address",
					Description: `(Deprecated) The address of the gateway assigned / to assign to the server. When used as part of request body, it has to match one of the IP addresses used in the existing assigned private networks for the relevant location. Deprecated in favour of a common gateway address across all networks available under ` + "`" + `network_configuration` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "configuration_type",
					Description: `Determines the approach for configuring IP blocks for the server being provisioned. Currently this field should be set to ` + "`" + `USE_OR_CREATE_DEFAULT` + "`" + ` or ` + "`" + `USER_DEFINED` + "`" + `. Default value is ` + "`" + `USE_OR_CREATE_DEFAULT` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "private_networks",
					Description: `The list of private networks this server is member of. When this field is part of request body, it'll be used to specify the private networks to assign to this server upon provisioning. Used alongside the ` + "`" + `USER_DEFINED` + "`" + ` configurationType. The ` + "`" + `private_networks` + "`" + ` block has field ` + "`" + `server_private_network` + "`" + `. The ` + "`" + `server_private_network` + "`" + ` block has 3 fields:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Required) The network identifier.`,
				},
				resource.Attribute{
					Name:        "ips",
					Description: `IPs to configure/configured on the server. Should be null or empty list if DHCP is true. Must contain at most 10 items.`,
				},
				resource.Attribute{
					Name:        "dhcp",
					Description: `Determines whether DHCP is enabled for this server. Should be false if ips is not an empty list. Not supported for proxmox OS. Default value is ` + "`" + `false` + "`" + `. The ` + "`" + `ip_blocks_configuration` + "`" + ` is the third field of the ` + "`" + `network_configuration` + "`" + ` block. The ` + "`" + `ip_blocks_configuration` + "`" + ` block has 2 fields:`,
				},
				resource.Attribute{
					Name:        "configuration_type",
					Description: `Determines the approach for configuring IP blocks for the server being provisioned. If ` + "`" + `PURCHASE_NEW` + "`" + ` is selected, the smallest supported range, depending on the operating system, is allocated to the server. The following values are allowed: ` + "`" + `PURCHASE_NEW` + "`" + `, ` + "`" + `USER_DEFINED` + "`" + `, ` + "`" + `NONE` + "`" + `. Default value is ` + "`" + `PURCHASE_NEW` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ip_blocks",
					Description: `Used to specify the previously purchased IP blocks to assign to this server upon provisioning. Used alongside the ` + "`" + `USER_DEFINED` + "`" + ` configurationType. Must contain at most 1 item. The ` + "`" + `ip_blocks` + "`" + ` block has field ` + "`" + `server_ip_block` + "`" + `. The ` + "`" + `server_ip_block` + "`" + ` block has 2 fields:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Required) The IP Block's ID.`,
				},
				resource.Attribute{
					Name:        "vlan_id",
					Description: `The VLAN on which this IP block has been configured within the network switch. The ` + "`" + `public_network_configuration` + "`" + ` is the fourth field of the ` + "`" + `network_configuration` + "`" + ` block. The ` + "`" + `public_network_configuration` + "`" + ` block has field ` + "`" + `public_networks` + "`" + `: The ` + "`" + `public_networks` + "`" + ` block has field ` + "`" + `server_public_network` + "`" + `. The ` + "`" + `server_public_network` + "`" + ` block has 2 fields:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Required) The network identifier.`,
				},
				resource.Attribute{
					Name:        "ips",
					Description: `(Required) IPs to configure on the server. IPs must be within the network's range. Must contain at least 1 item. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "cpu",
					Description: `A description of the machine CPU.`,
				},
				resource.Attribute{
					Name:        "cpu_count",
					Description: `The number of CPUs available in the system.`,
				},
				resource.Attribute{
					Name:        "cores_per_cpu",
					Description: `The number of physical cores present on each CPU.`,
				},
				resource.Attribute{
					Name:        "cpu_frequency_in_ghz",
					Description: `The CPU frequency in GHz.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Server description.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `Server hostname.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique identifier of the server.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `Server Location ID. Cannot be changed once a server is created.`,
				},
				resource.Attribute{
					Name:        "os",
					Description: `The server’s OS ID used when the server was created.`,
				},
				resource.Attribute{
					Name:        "ram",
					Description: `A description of the machine RAM.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the server.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Server type ID. Cannot be changed once a server is created.`,
				},
				resource.Attribute{
					Name:        "private_ip_addresses",
					Description: `Private IP Addresses assigned to server. Must contain at least 1 item.`,
				},
				resource.Attribute{
					Name:        "public_ip_addresses",
					Description: `Public IP Addresses assigned to server. Must contain at least 1 item.`,
				},
				resource.Attribute{
					Name:        "reservation_id",
					Description: `The reservation reference id if any.`,
				},
				resource.Attribute{
					Name:        "pricing_model",
					Description: `The pricing model this server is being billed.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `Password set for user Admin on Windows server which will only be returned in response to provisioning a server.`,
				},
				resource.Attribute{
					Name:        "network_type",
					Description: `The type of network configuration for this server.`,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `The cluster reference id if any.`,
				},
				resource.Attribute{
					Name:        "management_ui_url",
					Description: `The URL of the management UI which will only be returned in response to provisioning a server.`,
				},
				resource.Attribute{
					Name:        "root_password",
					Description: `Password set for user root on an ESXi server which will only be returned in response to provisioning a server.`,
				},
				resource.Attribute{
					Name:        "management_access_allowed_ips",
					Description: `A list of IPs allowed to access the Management UI. Supported in single IP, CIDR and range format. When undefined, Management UI is disabled.`,
				},
				resource.Attribute{
					Name:        "install_os_to_ram",
					Description: `If true, OS will be installed to and booted from the server's RAM. On restart RAM OS will be lost and the server will not be reachable unless a custom bootable OS has been deployed. Only supported for ubuntu/focal. Default value is ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "cloud_init",
					Description: `Cloud-init configuration details.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `The tags assigned if any.`,
				},
				resource.Attribute{
					Name:        "network_configuration",
					Description: `Entire network details of bare metal server.`,
				},
				resource.Attribute{
					Name:        "provisioned_on",
					Description: `Date and time when server was provisioned. The ` + "`" + `cloud_init` + "`" + ` block has one field:`,
				},
				resource.Attribute{
					Name:        "user_data",
					Description: `User data for the cloud-init configuration in base64 encoding. The ` + "`" + `tags` + "`" + ` block has field ` + "`" + `tag_assignment` + "`" + `. The ` + "`" + `tag_assignment` + "`" + ` block has 5 fields:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique id of the tag.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the tag.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag assigned to the server.`,
				},
				resource.Attribute{
					Name:        "is_billing_tag",
					Description: `Whether or not to show the tag as part of billing and invoices.`,
				},
				resource.Attribute{
					Name:        "created_by",
					Description: `Who the tag was created by.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "cpu",
					Description: `A description of the machine CPU.`,
				},
				resource.Attribute{
					Name:        "cpu_count",
					Description: `The number of CPUs available in the system.`,
				},
				resource.Attribute{
					Name:        "cores_per_cpu",
					Description: `The number of physical cores present on each CPU.`,
				},
				resource.Attribute{
					Name:        "cpu_frequency_in_ghz",
					Description: `The CPU frequency in GHz.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Server description.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `Server hostname.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique identifier of the server.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `Server Location ID. Cannot be changed once a server is created.`,
				},
				resource.Attribute{
					Name:        "os",
					Description: `The server’s OS ID used when the server was created.`,
				},
				resource.Attribute{
					Name:        "ram",
					Description: `A description of the machine RAM.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the server.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Server type ID. Cannot be changed once a server is created.`,
				},
				resource.Attribute{
					Name:        "private_ip_addresses",
					Description: `Private IP Addresses assigned to server. Must contain at least 1 item.`,
				},
				resource.Attribute{
					Name:        "public_ip_addresses",
					Description: `Public IP Addresses assigned to server. Must contain at least 1 item.`,
				},
				resource.Attribute{
					Name:        "reservation_id",
					Description: `The reservation reference id if any.`,
				},
				resource.Attribute{
					Name:        "pricing_model",
					Description: `The pricing model this server is being billed.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `Password set for user Admin on Windows server which will only be returned in response to provisioning a server.`,
				},
				resource.Attribute{
					Name:        "network_type",
					Description: `The type of network configuration for this server.`,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `The cluster reference id if any.`,
				},
				resource.Attribute{
					Name:        "management_ui_url",
					Description: `The URL of the management UI which will only be returned in response to provisioning a server.`,
				},
				resource.Attribute{
					Name:        "root_password",
					Description: `Password set for user root on an ESXi server which will only be returned in response to provisioning a server.`,
				},
				resource.Attribute{
					Name:        "management_access_allowed_ips",
					Description: `A list of IPs allowed to access the Management UI. Supported in single IP, CIDR and range format. When undefined, Management UI is disabled.`,
				},
				resource.Attribute{
					Name:        "install_os_to_ram",
					Description: `If true, OS will be installed to and booted from the server's RAM. On restart RAM OS will be lost and the server will not be reachable unless a custom bootable OS has been deployed. Only supported for ubuntu/focal. Default value is ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "cloud_init",
					Description: `Cloud-init configuration details.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `The tags assigned if any.`,
				},
				resource.Attribute{
					Name:        "network_configuration",
					Description: `Entire network details of bare metal server.`,
				},
				resource.Attribute{
					Name:        "provisioned_on",
					Description: `Date and time when server was provisioned. The ` + "`" + `cloud_init` + "`" + ` block has one field:`,
				},
				resource.Attribute{
					Name:        "user_data",
					Description: `User data for the cloud-init configuration in base64 encoding. The ` + "`" + `tags` + "`" + ` block has field ` + "`" + `tag_assignment` + "`" + `. The ` + "`" + `tag_assignment` + "`" + ` block has 5 fields:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique id of the tag.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the tag.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag assigned to the server.`,
				},
				resource.Attribute{
					Name:        "is_billing_tag",
					Description: `Whether or not to show the tag as part of billing and invoices.`,
				},
				resource.Attribute{
					Name:        "created_by",
					Description: `Who the tag was created by.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "pnap_ssh_key",
			Category:         "Resources",
			ShortDescription: `Provides a phoenixNAP SSH key resource. This can be used to create, modify, and delete SSH keys.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "default",
					Description: `(Required) Keys marked as default are always included on server creation and reset unless toggled off in creation/reset request.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Friendly SSH key name to represent an SSH key.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) SSH key actual key value. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique identifier of the SSH Key.`,
				},
				resource.Attribute{
					Name:        "default",
					Description: `Keys marked as default are always included on server creation and reset unless toggled off in creation/reset request.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Friendly SSH key name to represent an SSH key.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `SSH Key value.`,
				},
				resource.Attribute{
					Name:        "fingerprint",
					Description: `SSH key auto-generated SHA-256 fingerprint.`,
				},
				resource.Attribute{
					Name:        "lastUpdatedOn",
					Description: `Date and time of last update.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The unique identifier of the SSH Key.`,
				},
				resource.Attribute{
					Name:        "default",
					Description: `Keys marked as default are always included on server creation and reset unless toggled off in creation/reset request.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Friendly SSH key name to represent an SSH key.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `SSH Key value.`,
				},
				resource.Attribute{
					Name:        "fingerprint",
					Description: `SSH key auto-generated SHA-256 fingerprint.`,
				},
				resource.Attribute{
					Name:        "lastUpdatedOn",
					Description: `Date and time of last update.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "pnap_storage_network",
			Category:         "Resources",
			ShortDescription: `Provides a phoenixNAP Storage Network resource. This can be used to create, modify and delete storage networks.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The friendly name of this storage network. This name should be unique.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of this storage network.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `(Required) The location of this storage network. Currently this field should be set to ` + "`" + `PHX` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "volumes",
					Description: `(Required) Volumes to be created alongside storage. Currently only 1 volume is supported (must contain exactly one item).`,
				},
				resource.Attribute{
					Name:        "volume",
					Description: `(Required) Volume to be created alongside storage.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Volume friendly name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Volume description.`,
				},
				resource.Attribute{
					Name:        "path_suffix",
					Description: `Last part of volume's path.`,
				},
				resource.Attribute{
					Name:        "capacity_in_gb",
					Description: `(Required) Capacity of volume in GB. Currently only whole numbers and multiples of 1000 GB are supported. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The storage network identifier.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The friendly name of this storage network.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of this storage network.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Storage network's status.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The location of this storage network.`,
				},
				resource.Attribute{
					Name:        "ips",
					Description: `IP of the storage network`,
				},
				resource.Attribute{
					Name:        "created_on",
					Description: `Date and time when this storage network was created.`,
				},
				resource.Attribute{
					Name:        "volumes",
					Description: `Volumes for the storage network.`,
				},
				resource.Attribute{
					Name:        "volume",
					Description: `Volume for the storage network.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Volume ID.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Volume friendly name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Volume description.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `Volume's full path. It is in form of ` + "`" + `/{volumeId}/pathSuffix` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "path_suffix",
					Description: `Last part of volume's path.`,
				},
				resource.Attribute{
					Name:        "capacity_in_gb",
					Description: `Maximum capacity in GB.`,
				},
				resource.Attribute{
					Name:        "used_capacity_in_gb",
					Description: `Used capacity in GB, updated periodically.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `File system protocol.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Volume's status.`,
				},
				resource.Attribute{
					Name:        "created_on",
					Description: `Date and time when this volume was created.`,
				},
				resource.Attribute{
					Name:        "permissions",
					Description: `Permissions for the volume.`,
				},
				resource.Attribute{
					Name:        "nfs",
					Description: `NFS specific permissions on the volume.`,
				},
				resource.Attribute{
					Name:        "read_write",
					Description: `Read/Write access.`,
				},
				resource.Attribute{
					Name:        "read_only",
					Description: `Read only access.`,
				},
				resource.Attribute{
					Name:        "root_squash",
					Description: `Root squash permission.`,
				},
				resource.Attribute{
					Name:        "no_squash",
					Description: `No squash permission.`,
				},
				resource.Attribute{
					Name:        "all_squash",
					Description: `All squash permission.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The storage network identifier.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The friendly name of this storage network.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of this storage network.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Storage network's status.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The location of this storage network.`,
				},
				resource.Attribute{
					Name:        "ips",
					Description: `IP of the storage network`,
				},
				resource.Attribute{
					Name:        "created_on",
					Description: `Date and time when this storage network was created.`,
				},
				resource.Attribute{
					Name:        "volumes",
					Description: `Volumes for the storage network.`,
				},
				resource.Attribute{
					Name:        "volume",
					Description: `Volume for the storage network.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Volume ID.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Volume friendly name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Volume description.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `Volume's full path. It is in form of ` + "`" + `/{volumeId}/pathSuffix` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "path_suffix",
					Description: `Last part of volume's path.`,
				},
				resource.Attribute{
					Name:        "capacity_in_gb",
					Description: `Maximum capacity in GB.`,
				},
				resource.Attribute{
					Name:        "used_capacity_in_gb",
					Description: `Used capacity in GB, updated periodically.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `File system protocol.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Volume's status.`,
				},
				resource.Attribute{
					Name:        "created_on",
					Description: `Date and time when this volume was created.`,
				},
				resource.Attribute{
					Name:        "permissions",
					Description: `Permissions for the volume.`,
				},
				resource.Attribute{
					Name:        "nfs",
					Description: `NFS specific permissions on the volume.`,
				},
				resource.Attribute{
					Name:        "read_write",
					Description: `Read/Write access.`,
				},
				resource.Attribute{
					Name:        "read_only",
					Description: `Read only access.`,
				},
				resource.Attribute{
					Name:        "root_squash",
					Description: `Root squash permission.`,
				},
				resource.Attribute{
					Name:        "no_squash",
					Description: `No squash permission.`,
				},
				resource.Attribute{
					Name:        "all_squash",
					Description: `All squash permission.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "pnap_tag",
			Category:         "Resources",
			ShortDescription: `Provides a phoenixNAP tag resource. This can be used to create, modify, and delete tags.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The unique name of the tag.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the tag.`,
				},
				resource.Attribute{
					Name:        "is_billing_tag",
					Description: `(Required) Whether or not to show the tag as part of billing and invoices. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique identifier of the tag.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the tag.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `The optional values of the tag..`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the tag.`,
				},
				resource.Attribute{
					Name:        "resource_assignments",
					Description: `The tag's assigned resources.`,
				},
				resource.Attribute{
					Name:        "resource_name",
					Description: `The resource name.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag assigned to the resource.`,
				},
				resource.Attribute{
					Name:        "created_by",
					Description: `The tag's creator.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The unique identifier of the tag.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the tag.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `The optional values of the tag..`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the tag.`,
				},
				resource.Attribute{
					Name:        "resource_assignments",
					Description: `The tag's assigned resources.`,
				},
				resource.Attribute{
					Name:        "resource_name",
					Description: `The resource name.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the tag assigned to the resource.`,
				},
				resource.Attribute{
					Name:        "created_by",
					Description: `The tag's creator.`,
				},
			},
		},
	}

	resourcesMap = map[string]int{

		"pnap_ip_block":        0,
		"pnap_private_network": 1,
		"pnap_public_network":  2,
		"pnap_rancher_cluster": 3,
		"pnap_reservation":     4,
		"pnap_server":          5,
		"pnap_ssh_key":         6,
		"pnap_storage_network": 7,
		"pnap_tag":             8,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
