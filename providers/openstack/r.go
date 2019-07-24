package aws

import (
	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*Resource{

		&resource.Resource{
			Name:             "",
			Type:             "openstack_blockstorage_volume_attach_v2",
			Category:         "Block Storage Resources",
			ShortDescription: `Creates an attachment connection to a Block Storage volume`,
			Description:      ``,
			Keywords: []string{
				"block",
				"storage",
				"blockstorage",
				"volume",
				"attach",
				"v2",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 Block Storage client. A Block Storage client is needed to create a volume attachment. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new volume attachment.`,
				},
				resource.Attribute{
					Name:        "attach_mode",
					Description: `(Optional) Specify whether to attach the volume as Read-Only (` + "`" + `ro` + "`" + `) or Read-Write (` + "`" + `rw` + "`" + `). Only values of ` + "`" + `ro` + "`" + ` and ` + "`" + `rw` + "`" + ` are accepted. If left unspecified, the Block Storage API will apply a default of ` + "`" + `rw` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "device",
					Description: `(Optional) The device to tell the Block Storage service this volume will be attached as. This is purely for informational purposes. You can specify ` + "`" + `auto` + "`" + ` or a device such as ` + "`" + `/dev/vdc` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "host_name",
					Description: `(Required) The host to attach the volume to.`,
				},
				resource.Attribute{
					Name:        "initiator",
					Description: `(Optional) The iSCSI initiator string to make the connection.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `(Optional) The IP address of the ` + "`" + `host_name` + "`" + ` above.`,
				},
				resource.Attribute{
					Name:        "multipath",
					Description: `(Optional) Whether to connect to this volume via multipath.`,
				},
				resource.Attribute{
					Name:        "os_type",
					Description: `(Optional) The iSCSI initiator OS type.`,
				},
				resource.Attribute{
					Name:        "platform",
					Description: `(Optional) The iSCSI initiator platform.`,
				},
				resource.Attribute{
					Name:        "volume_id",
					Description: `(Required) The ID of the Volume to attach to an Instance.`,
				},
				resource.Attribute{
					Name:        "wwpn",
					Description: `(Optional) An array of wwpn strings. Used for Fibre Channel connections.`,
				},
				resource.Attribute{
					Name:        "wwnn",
					Description: `(Optional) A wwnn name. Used for Fibre Channel connections. ## Attributes Reference In addition to the above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "data",
					Description: `This is a map of key/value pairs that contain the connection information. You will want to pass this information to a provisioner script to finalize the connection. See below for more information.`,
				},
				resource.Attribute{
					Name:        "driver_volume_type",
					Description: `The storage driver that the volume is based on.`,
				},
				resource.Attribute{
					Name:        "mount_point_base",
					Description: `A mount point base name for shared storage. ## Volume Connection Data Upon creation of this resource, a ` + "`" + `data` + "`" + ` exported attribute will be available. This attribute is a set of key/value pairs that contains the information required to complete the block storage connection. As an example, creating an iSCSI-based volume will return the following: ` + "`" + `` + "`" + `` + "`" + ` data.access_mode = rw data.auth_method = CHAP data.auth_password = xUhbGKQ8QCwKmHQ2 data.auth_username = Sphn5X4EoyFUUMYVYSA4 data.target_iqn = iqn.2010-10.org.openstack:volume-2d87ed25-c312-4f42-be1d-3b36b014561d data.target_portal = 192.168.255.10:3260 data.volume_id = 2d87ed25-c312-4f42-be1d-3b36b014561d ` + "`" + `` + "`" + `` + "`" + ` This information can then be fed into a provisioner or a template shell script, where the final result would look something like: ` + "`" + `` + "`" + `` + "`" + ` iscsiadm -m node -T ${self.data.target_iqn} -p ${self.data.target_portal} --interface default --op new iscsiadm -m node -T ${self.data.target_iqn} -p ${self.data.target_portal} --op update -n node.session.auth.authmethod -v ${self.data.auth_method} iscsiadm -m node -T ${self.data.target_iqn} -p ${self.data.target_portal} --op update -n node.session.auth.username -v ${self.data.auth_username} iscsiadm -m node -T ${self.data.target_iqn} -p ${self.data.target_portal} --op update -n node.session.auth.password -v ${self.data.auth_password} iscsiadm -m node -T ${self.data.target_iqn} -p ${self.data.target_portal} --login iscsiadm -m node -T ${self.data.target_iqn} -p ${self.data.target_portal} --op update -n node.startup -v automatic iscsiadm -m node -T ${self.data.target_iqn} -p ${self.data.target_portal} --rescan ` + "`" + `` + "`" + `` + "`" + ` The contents of ` + "`" + `data` + "`" + ` will vary from each Block Storage service. You must have a good understanding of how the service is configured and how to make the appropriate final connection. However, if used correctly, this has the flexibility to be able to attach OpenStack Block Storage volumes to non-OpenStack resources. ## Import It is not possible to import this resource.`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "data",
					Description: `This is a map of key/value pairs that contain the connection information. You will want to pass this information to a provisioner script to finalize the connection. See below for more information.`,
				},
				resource.Attribute{
					Name:        "driver_volume_type",
					Description: `The storage driver that the volume is based on.`,
				},
				resource.Attribute{
					Name:        "mount_point_base",
					Description: `A mount point base name for shared storage. ## Volume Connection Data Upon creation of this resource, a ` + "`" + `data` + "`" + ` exported attribute will be available. This attribute is a set of key/value pairs that contains the information required to complete the block storage connection. As an example, creating an iSCSI-based volume will return the following: ` + "`" + `` + "`" + `` + "`" + ` data.access_mode = rw data.auth_method = CHAP data.auth_password = xUhbGKQ8QCwKmHQ2 data.auth_username = Sphn5X4EoyFUUMYVYSA4 data.target_iqn = iqn.2010-10.org.openstack:volume-2d87ed25-c312-4f42-be1d-3b36b014561d data.target_portal = 192.168.255.10:3260 data.volume_id = 2d87ed25-c312-4f42-be1d-3b36b014561d ` + "`" + `` + "`" + `` + "`" + ` This information can then be fed into a provisioner or a template shell script, where the final result would look something like: ` + "`" + `` + "`" + `` + "`" + ` iscsiadm -m node -T ${self.data.target_iqn} -p ${self.data.target_portal} --interface default --op new iscsiadm -m node -T ${self.data.target_iqn} -p ${self.data.target_portal} --op update -n node.session.auth.authmethod -v ${self.data.auth_method} iscsiadm -m node -T ${self.data.target_iqn} -p ${self.data.target_portal} --op update -n node.session.auth.username -v ${self.data.auth_username} iscsiadm -m node -T ${self.data.target_iqn} -p ${self.data.target_portal} --op update -n node.session.auth.password -v ${self.data.auth_password} iscsiadm -m node -T ${self.data.target_iqn} -p ${self.data.target_portal} --login iscsiadm -m node -T ${self.data.target_iqn} -p ${self.data.target_portal} --op update -n node.startup -v automatic iscsiadm -m node -T ${self.data.target_iqn} -p ${self.data.target_portal} --rescan ` + "`" + `` + "`" + `` + "`" + ` The contents of ` + "`" + `data` + "`" + ` will vary from each Block Storage service. You must have a good understanding of how the service is configured and how to make the appropriate final connection. However, if used correctly, this has the flexibility to be able to attach OpenStack Block Storage volumes to non-OpenStack resources. ## Import It is not possible to import this resource.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "openstack_blockstorage_volume_attach_v3",
			Category:         "Block Storage Resources",
			ShortDescription: `Creates an attachment connection to a Block Storage volume`,
			Description:      ``,
			Keywords: []string{
				"block",
				"storage",
				"blockstorage",
				"volume",
				"attach",
				"v3",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V3 Block Storage client. A Block Storage client is needed to create a volume attachment. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new volume attachment.`,
				},
				resource.Attribute{
					Name:        "attach_mode",
					Description: `(Optional) Specify whether to attach the volume as Read-Only (` + "`" + `ro` + "`" + `) or Read-Write (` + "`" + `rw` + "`" + `). Only values of ` + "`" + `ro` + "`" + ` and ` + "`" + `rw` + "`" + ` are accepted. If left unspecified, the Block Storage API will apply a default of ` + "`" + `rw` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "device",
					Description: `(Optional) The device to tell the Block Storage service this volume will be attached as. This is purely for informational purposes. You can specify ` + "`" + `auto` + "`" + ` or a device such as ` + "`" + `/dev/vdc` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "host_name",
					Description: `(Required) The host to attach the volume to.`,
				},
				resource.Attribute{
					Name:        "initiator",
					Description: `(Optional) The iSCSI initiator string to make the connection.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `(Optional) The IP address of the ` + "`" + `host_name` + "`" + ` above.`,
				},
				resource.Attribute{
					Name:        "multipath",
					Description: `(Optional) Whether to connect to this volume via multipath.`,
				},
				resource.Attribute{
					Name:        "os_type",
					Description: `(Optional) The iSCSI initiator OS type.`,
				},
				resource.Attribute{
					Name:        "platform",
					Description: `(Optional) The iSCSI initiator platform.`,
				},
				resource.Attribute{
					Name:        "volume_id",
					Description: `(Required) The ID of the Volume to attach to an Instance.`,
				},
				resource.Attribute{
					Name:        "wwpn",
					Description: `(Optional) An array of wwpn strings. Used for Fibre Channel connections.`,
				},
				resource.Attribute{
					Name:        "wwnn",
					Description: `(Optional) A wwnn name. Used for Fibre Channel connections. ## Attributes Reference In addition to the above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "data",
					Description: `This is a map of key/value pairs that contain the connection information. You will want to pass this information to a provisioner script to finalize the connection. See below for more information.`,
				},
				resource.Attribute{
					Name:        "driver_volume_type",
					Description: `The storage driver that the volume is based on.`,
				},
				resource.Attribute{
					Name:        "mount_point_base",
					Description: `A mount point base name for shared storage. ## Volume Connection Data Upon creation of this resource, a ` + "`" + `data` + "`" + ` exported attribute will be available. This attribute is a set of key/value pairs that contains the information required to complete the block storage connection. As an example, creating an iSCSI-based volume will return the following: ` + "`" + `` + "`" + `` + "`" + ` data.access_mode = rw data.auth_method = CHAP data.auth_password = xUhbGKQ8QCwKmHQ2 data.auth_username = Sphn5X4EoyFUUMYVYSA4 data.target_iqn = iqn.2010-10.org.openstack:volume-2d87ed25-c312-4f42-be1d-3b36b014561d data.target_portal = 192.168.255.10:3260 data.volume_id = 2d87ed25-c312-4f42-be1d-3b36b014561d ` + "`" + `` + "`" + `` + "`" + ` This information can then be fed into a provisioner or a template shell script, where the final result would look something like: ` + "`" + `` + "`" + `` + "`" + ` iscsiadm -m node -T ${self.data.target_iqn} -p ${self.data.target_portal} --interface default --op new iscsiadm -m node -T ${self.data.target_iqn} -p ${self.data.target_portal} --op update -n node.session.auth.authmethod -v ${self.data.auth_method} iscsiadm -m node -T ${self.data.target_iqn} -p ${self.data.target_portal} --op update -n node.session.auth.username -v ${self.data.auth_username} iscsiadm -m node -T ${self.data.target_iqn} -p ${self.data.target_portal} --op update -n node.session.auth.password -v ${self.data.auth_password} iscsiadm -m node -T ${self.data.target_iqn} -p ${self.data.target_portal} --login iscsiadm -m node -T ${self.data.target_iqn} -p ${self.data.target_portal} --op update -n node.startup -v automatic iscsiadm -m node -T ${self.data.target_iqn} -p ${self.data.target_portal} --rescan ` + "`" + `` + "`" + `` + "`" + ` The contents of ` + "`" + `data` + "`" + ` will vary from each Block Storage service. You must have a good understanding of how the service is configured and how to make the appropriate final connection. However, if used correctly, this has the flexibility to be able to attach OpenStack Block Storage volumes to non-OpenStack resources. ## Import It is not possible to import this resource.`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "data",
					Description: `This is a map of key/value pairs that contain the connection information. You will want to pass this information to a provisioner script to finalize the connection. See below for more information.`,
				},
				resource.Attribute{
					Name:        "driver_volume_type",
					Description: `The storage driver that the volume is based on.`,
				},
				resource.Attribute{
					Name:        "mount_point_base",
					Description: `A mount point base name for shared storage. ## Volume Connection Data Upon creation of this resource, a ` + "`" + `data` + "`" + ` exported attribute will be available. This attribute is a set of key/value pairs that contains the information required to complete the block storage connection. As an example, creating an iSCSI-based volume will return the following: ` + "`" + `` + "`" + `` + "`" + ` data.access_mode = rw data.auth_method = CHAP data.auth_password = xUhbGKQ8QCwKmHQ2 data.auth_username = Sphn5X4EoyFUUMYVYSA4 data.target_iqn = iqn.2010-10.org.openstack:volume-2d87ed25-c312-4f42-be1d-3b36b014561d data.target_portal = 192.168.255.10:3260 data.volume_id = 2d87ed25-c312-4f42-be1d-3b36b014561d ` + "`" + `` + "`" + `` + "`" + ` This information can then be fed into a provisioner or a template shell script, where the final result would look something like: ` + "`" + `` + "`" + `` + "`" + ` iscsiadm -m node -T ${self.data.target_iqn} -p ${self.data.target_portal} --interface default --op new iscsiadm -m node -T ${self.data.target_iqn} -p ${self.data.target_portal} --op update -n node.session.auth.authmethod -v ${self.data.auth_method} iscsiadm -m node -T ${self.data.target_iqn} -p ${self.data.target_portal} --op update -n node.session.auth.username -v ${self.data.auth_username} iscsiadm -m node -T ${self.data.target_iqn} -p ${self.data.target_portal} --op update -n node.session.auth.password -v ${self.data.auth_password} iscsiadm -m node -T ${self.data.target_iqn} -p ${self.data.target_portal} --login iscsiadm -m node -T ${self.data.target_iqn} -p ${self.data.target_portal} --op update -n node.startup -v automatic iscsiadm -m node -T ${self.data.target_iqn} -p ${self.data.target_portal} --rescan ` + "`" + `` + "`" + `` + "`" + ` The contents of ` + "`" + `data` + "`" + ` will vary from each Block Storage service. You must have a good understanding of how the service is configured and how to make the appropriate final connection. However, if used correctly, this has the flexibility to be able to attach OpenStack Block Storage volumes to non-OpenStack resources. ## Import It is not possible to import this resource.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "openstack_blockstorage_volume_v1",
			Category:         "Block Storage Resources",
			ShortDescription: `Manages a V1 volume resource within OpenStack.`,
			Description:      ``,
			Keywords: []string{
				"block",
				"storage",
				"blockstorage",
				"volume",
				"v1",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to create the volume. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new volume.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `(Required) The size of the volume to create (in gigabytes). Changing this creates a new volume.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) A unique name for the volume. Changing this updates the volume's name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A description of the volume. Changing this updates the volume's description.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `(Optional) The availability zone for the volume. Changing this creates a new volume.`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `(Optional) The image ID from which to create the volume. Changing this creates a new volume.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `(Optional) The snapshot ID from which to create the volume. Changing this creates a new volume.`,
				},
				resource.Attribute{
					Name:        "source_vol_id",
					Description: `(Optional) The volume ID from which to create the volume. Changing this creates a new volume.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `(Optional) Metadata key/value pairs to associate with the volume. Changing this updates the existing volume metadata.`,
				},
				resource.Attribute{
					Name:        "volume_type",
					Description: `(Optional) The type of volume to create. Changing this creates a new volume. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "source_vol_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "volume_type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "attachment",
					Description: `If a volume is attached to an instance, this attribute will display the Attachment ID, Instance ID, and the Device as the Instance sees it. ## Import Volumes can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import openstack_blockstorage_volume_v1.volume_1 ea257959-eeb1-4c10-8d33-26f0409a755d ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "source_vol_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "volume_type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "attachment",
					Description: `If a volume is attached to an instance, this attribute will display the Attachment ID, Instance ID, and the Device as the Instance sees it. ## Import Volumes can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import openstack_blockstorage_volume_v1.volume_1 ea257959-eeb1-4c10-8d33-26f0409a755d ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "openstack_blockstorage_volume_v2",
			Category:         "Block Storage Resources",
			ShortDescription: `Manages a V2 volume resource within OpenStack.`,
			Description:      ``,
			Keywords: []string{
				"block",
				"storage",
				"blockstorage",
				"volume",
				"v2",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to create the volume. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new volume.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `(Required) The size of the volume to create (in gigabytes). Changing this creates a new volume.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `(Optional) The availability zone for the volume. Changing this creates a new volume.`,
				},
				resource.Attribute{
					Name:        "consistency_group_id",
					Description: `(Optional) The consistency group to place the volume in.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A description of the volume. Changing this updates the volume's description.`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `(Optional) The image ID from which to create the volume. Changing this creates a new volume.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `(Optional) Metadata key/value pairs to associate with the volume. Changing this updates the existing volume metadata.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) A unique name for the volume. Changing this updates the volume's name.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `(Optional) The snapshot ID from which to create the volume. Changing this creates a new volume.`,
				},
				resource.Attribute{
					Name:        "source_replica",
					Description: `(Optional) The volume ID to replicate with.`,
				},
				resource.Attribute{
					Name:        "source_vol_id",
					Description: `(Optional) The volume ID from which to create the volume. Changing this creates a new volume.`,
				},
				resource.Attribute{
					Name:        "volume_type",
					Description: `(Optional) The type of volume to create. Changing this creates a new volume. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "source_vol_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "volume_type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "attachment",
					Description: `If a volume is attached to an instance, this attribute will display the Attachment ID, Instance ID, and the Device as the Instance sees it. ## Import Volumes can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import openstack_blockstorage_volume_v2.volume_1 ea257959-eeb1-4c10-8d33-26f0409a755d ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "source_vol_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "volume_type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "attachment",
					Description: `If a volume is attached to an instance, this attribute will display the Attachment ID, Instance ID, and the Device as the Instance sees it. ## Import Volumes can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import openstack_blockstorage_volume_v2.volume_1 ea257959-eeb1-4c10-8d33-26f0409a755d ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "openstack_blockstorage_volume_v3",
			Category:         "Block Storage Resources",
			ShortDescription: `Manages a V3 volume resource within OpenStack.`,
			Description:      ``,
			Keywords: []string{
				"block",
				"storage",
				"blockstorage",
				"volume",
				"v3",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to create the volume. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new volume.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `(Required) The size of the volume to create (in gigabytes).`,
				},
				resource.Attribute{
					Name:        "enable_online_resize",
					Description: `(Optional) When this option is set it allows extending attached volumes. Note: updating size of an attached volume requires Cinder support for version 3.42 and a compatible storage driver.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `(Optional) The availability zone for the volume. Changing this creates a new volume.`,
				},
				resource.Attribute{
					Name:        "consistency_group_id",
					Description: `(Optional) The consistency group to place the volume in.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A description of the volume. Changing this updates the volume's description.`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `(Optional) The image ID from which to create the volume. Changing this creates a new volume.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `(Optional) Metadata key/value pairs to associate with the volume. Changing this updates the existing volume metadata.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) A unique name for the volume. Changing this updates the volume's name.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `(Optional) The snapshot ID from which to create the volume. Changing this creates a new volume.`,
				},
				resource.Attribute{
					Name:        "source_replica",
					Description: `(Optional) The volume ID to replicate with.`,
				},
				resource.Attribute{
					Name:        "source_vol_id",
					Description: `(Optional) The volume ID from which to create the volume. Changing this creates a new volume.`,
				},
				resource.Attribute{
					Name:        "volume_type",
					Description: `(Optional) The type of volume to create. Changing this creates a new volume.`,
				},
				resource.Attribute{
					Name:        "multiattach",
					Description: `(Optional) Allow the volume to be attached to more than one Compute instance. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "source_vol_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "volume_type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "attachment",
					Description: `If a volume is attached to an instance, this attribute will display the Attachment ID, Instance ID, and the Device as the Instance sees it.`,
				},
				resource.Attribute{
					Name:        "multiattach",
					Description: `See Argument Reference above. ## Import Volumes can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import openstack_blockstorage_volume_v3.volume_1 ea257959-eeb1-4c10-8d33-26f0409a755d ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "source_vol_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "volume_type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "attachment",
					Description: `If a volume is attached to an instance, this attribute will display the Attachment ID, Instance ID, and the Device as the Instance sees it.`,
				},
				resource.Attribute{
					Name:        "multiattach",
					Description: `See Argument Reference above. ## Import Volumes can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import openstack_blockstorage_volume_v3.volume_1 ea257959-eeb1-4c10-8d33-26f0409a755d ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "openstack_compute_flavor_access_v2",
			Category:         "Compute Resources",
			ShortDescription: `Manages a project access for flavor V2 resource within OpenStack.`,
			Description:      ``,
			Keywords: []string{
				"compute",
				"flavor",
				"access",
				"v2",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 Compute client. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new flavor access.`,
				},
				resource.Attribute{
					Name:        "flavor_id",
					Description: `(Required) The UUID of flavor to use. Changing this creates a new flavor access.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Required) The UUID of tenant which is allowed to use the flavor. Changing this creates a new flavor access. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "flavor_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above. ## Import This resource can be imported by specifying all two arguments, separated by a forward slash: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import openstack_compute_flavor_access_v2.access_1 <flavor_id>/<tenant_id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "flavor_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above. ## Import This resource can be imported by specifying all two arguments, separated by a forward slash: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import openstack_compute_flavor_access_v2.access_1 <flavor_id>/<tenant_id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "openstack_compute_flavor_v2",
			Category:         "Compute Resources",
			ShortDescription: `Manages a V2 flavor resource within OpenStack.`,
			Description:      ``,
			Keywords: []string{
				"compute",
				"flavor",
				"v2",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 Compute client. Flavors are associated with accounts, but a Compute client is needed to create one. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new flavor.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A unique name for the flavor. Changing this creates a new flavor.`,
				},
				resource.Attribute{
					Name:        "ram",
					Description: `(Required) The amount of RAM to use, in megabytes. Changing this creates a new flavor.`,
				},
				resource.Attribute{
					Name:        "vcpus",
					Description: `(Required) The number of virtual CPUs to use. Changing this creates a new flavor.`,
				},
				resource.Attribute{
					Name:        "disk",
					Description: `(Required) The amount of disk space in gigabytes to use for the root (/) partition. Changing this creates a new flavor.`,
				},
				resource.Attribute{
					Name:        "swap",
					Description: `(Optional) The amount of disk space in megabytes to use. If unspecified, the default is 0. Changing this creates a new flavor.`,
				},
				resource.Attribute{
					Name:        "rx_tx_factor",
					Description: `(Optional) RX/TX bandwith factor. The default is 1. Changing this creates a new flavor.`,
				},
				resource.Attribute{
					Name:        "is_public",
					Description: `(Optional) Whether the flavor is public. Changing this creates a new flavor.`,
				},
				resource.Attribute{
					Name:        "extra_specs",
					Description: `(Optional) Key/Value pairs of metadata for the flavor. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "ram",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "vcpus",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "disk",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "swap",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "rx_tx_factor",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "is_public",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "extra_specs",
					Description: `See Argument Reference above. ## Import Flavors can be imported using the ` + "`" + `ID` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import openstack_compute_flavor_v2.my-flavor 4142e64b-1b35-44a0-9b1e-5affc7af1106 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "ram",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "vcpus",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "disk",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "swap",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "rx_tx_factor",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "is_public",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "extra_specs",
					Description: `See Argument Reference above. ## Import Flavors can be imported using the ` + "`" + `ID` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import openstack_compute_flavor_v2.my-flavor 4142e64b-1b35-44a0-9b1e-5affc7af1106 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "openstack_compute_floatingip_associate_v2",
			Category:         "Compute Resources",
			ShortDescription: `Associate a floating IP to an instance`,
			Description:      ``,
			Keywords: []string{
				"compute",
				"floatingip",
				"associate",
				"v2",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 Compute client. Keypairs are associated with accounts, but a Compute client is needed to create one. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new floatingip_associate.`,
				},
				resource.Attribute{
					Name:        "floating_ip",
					Description: `(Required) The floating IP to associate.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Required) The instance to associte the floating IP with.`,
				},
				resource.Attribute{
					Name:        "fixed_ip",
					Description: `(Optional) The specific IP address to direct traffic to.`,
				},
				resource.Attribute{
					Name:        "wait_until_associated",
					Description: `(Optional) In cases where the OpenStack environment does not automatically wait until the association has finished, set this option to have Terraform poll the instance until the floating IP has been associated. Defaults to false. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "floating_ip",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "fixed_ip",
					Description: `See Argument Reference above. ## Import This resource can be imported by specifying all three arguments, separated by a forward slash: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import openstack_compute_floatingip_associate_v2.fip_1 <floating_ip>/<instance_id>/<fixed_ip> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "floating_ip",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "fixed_ip",
					Description: `See Argument Reference above. ## Import This resource can be imported by specifying all three arguments, separated by a forward slash: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import openstack_compute_floatingip_associate_v2.fip_1 <floating_ip>/<instance_id>/<fixed_ip> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "openstack_compute_floatingip_v2",
			Category:         "Compute Resources",
			ShortDescription: `Manages a V2 floating IP resource within OpenStack Nova (compute).`,
			Description:      ``,
			Keywords: []string{
				"compute",
				"floatingip",
				"v2",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 Compute client. A Compute client is needed to create a floating IP that can be used with a compute instance. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new floating IP (which may or may not have a different address).`,
				},
				resource.Attribute{
					Name:        "pool",
					Description: `(Required) The name of the pool from which to obtain the floating IP. Changing this creates a new floating IP. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "pool",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `The actual floating IP address itself.`,
				},
				resource.Attribute{
					Name:        "fixed_ip",
					Description: `The fixed IP address corresponding to the floating IP.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `UUID of the compute instance associated with the floating IP. ## Import Floating IPs can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import openstack_compute_floatingip_v2.floatip_1 89c60255-9bd6-460c-822a-e2b959ede9d2 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "pool",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `The actual floating IP address itself.`,
				},
				resource.Attribute{
					Name:        "fixed_ip",
					Description: `The fixed IP address corresponding to the floating IP.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `UUID of the compute instance associated with the floating IP. ## Import Floating IPs can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import openstack_compute_floatingip_v2.floatip_1 89c60255-9bd6-460c-822a-e2b959ede9d2 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "openstack_compute_instance_v2",
			Category:         "Compute Resources",
			ShortDescription: `Manages a V2 VM instance resource within OpenStack.`,
			Description:      ``,
			Keywords: []string{
				"compute",
				"instance",
				"v2",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to create the server instance. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new server.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A unique name for the resource.`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `(Optional; Required if ` + "`" + `image_name` + "`" + ` is empty and not booting from a volume. Do not specify if booting from a volume.) The image ID of the desired image for the server. Changing this creates a new server.`,
				},
				resource.Attribute{
					Name:        "image_name",
					Description: `(Optional; Required if ` + "`" + `image_id` + "`" + ` is empty and not booting from a volume. Do not specify if booting from a volume.) The name of the desired image for the server. Changing this creates a new server.`,
				},
				resource.Attribute{
					Name:        "flavor_id",
					Description: `(Optional; Required if ` + "`" + `flavor_name` + "`" + ` is empty) The flavor ID of the desired flavor for the server. Changing this resizes the existing server.`,
				},
				resource.Attribute{
					Name:        "flavor_name",
					Description: `(Optional; Required if ` + "`" + `flavor_id` + "`" + ` is empty) The name of the desired flavor for the server. Changing this resizes the existing server.`,
				},
				resource.Attribute{
					Name:        "user_data",
					Description: `(Optional) The user data to provide when launching the instance. Changing this creates a new server.`,
				},
				resource.Attribute{
					Name:        "security_groups",
					Description: `(Optional) An array of one or more security group names to associate with the server. Changing this results in adding/removing security groups from the existing server.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `(Optional) The availability zone in which to create the server. Changing this creates a new server.`,
				},
				resource.Attribute{
					Name:        "network",
					Description: `(Optional) An array of one or more networks to attach to the instance. The network object structure is documented below. Changing this creates a new server.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `(Optional) Metadata key/value pairs to make available from within the instance. Changing this updates the existing server metadata.`,
				},
				resource.Attribute{
					Name:        "config_drive",
					Description: `(Optional) Whether to use the config_drive feature to configure the instance. Changing this creates a new server.`,
				},
				resource.Attribute{
					Name:        "admin_pass",
					Description: `(Optional) The administrative password to assign to the server. Changing this changes the root password on the existing server.`,
				},
				resource.Attribute{
					Name:        "key_pair",
					Description: `(Optional) The name of a key pair to put on the server. The key pair must already be created and associated with the tenant's account. Changing this creates a new server.`,
				},
				resource.Attribute{
					Name:        "block_device",
					Description: `(Optional) Configuration of block devices. The block_device structure is documented below. Changing this creates a new server. You can specify multiple block devices which will create an instance with multiple disks. This configuration is very flexible, so please see the following [reference](https://docs.openstack.org/nova/latest/user/block-device-mapping.html) for more information.`,
				},
				resource.Attribute{
					Name:        "scheduler_hints",
					Description: `(Optional) Provide the Nova scheduler with hints on how the instance should be launched. The available hints are described below.`,
				},
				resource.Attribute{
					Name:        "personality",
					Description: `(Optional) Customize the personality of an instance by defining one or more files and their contents. The personality structure is described below.`,
				},
				resource.Attribute{
					Name:        "stop_before_destroy",
					Description: `(Optional) Whether to try stop instance gracefully before destroying it, thus giving chance for guest OS daemons to stop correctly. If instance doesn't stop within timeout, it will be destroyed anyway.`,
				},
				resource.Attribute{
					Name:        "force_delete",
					Description: `(Optional) Whether to force the OpenStack instance to be forcefully deleted. This is useful for environments that have reclaim / soft deletion enabled.`,
				},
				resource.Attribute{
					Name:        "power_state",
					Description: `(Optional) Provide the VM state. Only 'active' and 'shutoff' are supported values.`,
				},
				resource.Attribute{
					Name:        "vendor_options",
					Description: `(Optional) Map of additional vendor-specific options. Supported options are described below. The ` + "`" + `network` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `(Required unless ` + "`" + `port` + "`" + ` or ` + "`" + `name` + "`" + ` is provided) The network UUID to attach to the server. Changing this creates a new server.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required unless ` + "`" + `uuid` + "`" + ` or ` + "`" + `port` + "`" + ` is provided) The human-readable name of the network. Changing this creates a new server.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Required unless ` + "`" + `uuid` + "`" + ` or ` + "`" + `name` + "`" + ` is provided) The port UUID of a network to attach to the server. Changing this creates a new server.`,
				},
				resource.Attribute{
					Name:        "fixed_ip_v4",
					Description: `(Optional) Specifies a fixed IPv4 address to be used on this network. Changing this creates a new server.`,
				},
				resource.Attribute{
					Name:        "access_network",
					Description: `(Optional) Specifies if this network should be used for provisioning access. Accepts true or false. Defaults to false. The ` + "`" + `block_device` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `(Required unless ` + "`" + `source_type` + "`" + ` is set to ` + "`" + `"blank"` + "`" + ` ) The UUID of the image, volume, or snapshot. Changing this creates a new server.`,
				},
				resource.Attribute{
					Name:        "source_type",
					Description: `(Required) The source type of the device. Must be one of "blank", "image", "volume", or "snapshot". Changing this creates a new server.`,
				},
				resource.Attribute{
					Name:        "volume_size",
					Description: `The size of the volume to create (in gigabytes). Required in the following combinations: source=image and destination=volume, source=blank and destination=local, and source=blank and destination=volume. Changing this creates a new server.`,
				},
				resource.Attribute{
					Name:        "boot_index",
					Description: `(Optional) The boot index of the volume. It defaults to 0. Changing this creates a new server.`,
				},
				resource.Attribute{
					Name:        "destination_type",
					Description: `(Optional) The type that gets created. Possible values are "volume" and "local". Changing this creates a new server.`,
				},
				resource.Attribute{
					Name:        "delete_on_termination",
					Description: `(Optional) Delete the volume / block device upon termination of the instance. Defaults to false. Changing this creates a new server.`,
				},
				resource.Attribute{
					Name:        "device_type",
					Description: `(Optional) The low-level device type that will be used. Most common thing is to leave this empty. Changing this creates a new server.`,
				},
				resource.Attribute{
					Name:        "disk_bus",
					Description: `(Optional) The low-level disk bus that will be used. Most common thing is to leave this empty. Changing this creates a new server. The ` + "`" + `scheduler_hints` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "group",
					Description: `(Optional) A UUID of a Server Group. The instance will be placed into that group.`,
				},
				resource.Attribute{
					Name:        "different_host",
					Description: `(Optional) A list of instance UUIDs. The instance will be scheduled on a different host than all other instances.`,
				},
				resource.Attribute{
					Name:        "same_host",
					Description: `(Optional) A list of instance UUIDs. The instance will be scheduled on the same host of those specified.`,
				},
				resource.Attribute{
					Name:        "query",
					Description: `(Optional) A conditional query that a compute node must pass in order to host an instance. The query must use the ` + "`" + `JsonFilter` + "`" + ` syntax which is described [here](https://docs.openstack.org/nova/latest/admin/configuration/schedulers.html#jsonfilter). At this time, only simple queries are supported. Compound queries using ` + "`" + `and` + "`" + `, ` + "`" + `or` + "`" + `, or ` + "`" + `not` + "`" + ` are not supported. An example of a simple query is: ` + "`" + `` + "`" + `` + "`" + ` [">=", "$free_ram_mb", "1024"] ` + "`" + `` + "`" + `` + "`" + ``,
				},
				resource.Attribute{
					Name:        "target_cell",
					Description: `(Optional) The name of a cell to host the instance.`,
				},
				resource.Attribute{
					Name:        "build_near_host_ip",
					Description: `(Optional) An IP Address in CIDR form. The instance will be placed on a compute node that is in the same subnet.`,
				},
				resource.Attribute{
					Name:        "additional_properties",
					Description: `(Optional) Arbitrary key/value pairs of additional properties to pass to the scheduler. The ` + "`" + `personality` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "file",
					Description: `(Required) The absolute path of the destination file.`,
				},
				resource.Attribute{
					Name:        "content",
					Description: `(Required) The contents of the file. Limited to 255 bytes. The ` + "`" + `vendor_options` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "ignore_resize_confirmation",
					Description: `(Optional) Boolean to control whether to ignore manual confirmation of the instance resizing. This can be helpful to work with some OpenStack clouds which automatically confirm resizing of instances after some timeout. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "access_ip_v4",
					Description: `The first detected Fixed IPv4 address.`,
				},
				resource.Attribute{
					Name:        "access_ip_v6",
					Description: `The first detected Fixed IPv6 address.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "security_groups",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "flavor_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "flavor_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "network/uuid",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "network/name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "network/port",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "network/fixed_ip_v4",
					Description: `The Fixed IPv4 address of the Instance on that network.`,
				},
				resource.Attribute{
					Name:        "network/fixed_ip_v6",
					Description: `The Fixed IPv6 address of the Instance on that network.`,
				},
				resource.Attribute{
					Name:        "network/mac",
					Description: `The MAC address of the NIC on that network.`,
				},
				resource.Attribute{
					Name:        "all_metadata",
					Description: `Contains all instance metadata, even metadata not set by Terraform. ## Notes ### Multiple Ephemeral Disks It's possible to specify multiple ` + "`" + `block_device` + "`" + ` entries to create an instance with multiple ephemeral (local) disks. In order to create multiple ephemeral disks, the sum of the total amount of ephemeral space must be less than or equal to what the chosen flavor supports. The following example shows how to create an instance with multiple ephemeral disks: ` + "`" + `` + "`" + `` + "`" + `hcl resource "openstack_compute_instance_v2" "foo" { name = "terraform-test" security_groups = ["default"] block_device { boot_index = 0 delete_on_termination = true destination_type = "local" source_type = "image" uuid = "<image uuid>" } block_device { boot_index = -1 delete_on_termination = true destination_type = "local" source_type = "blank" volume_size = 1 } block_device { boot_index = -1 delete_on_termination = true destination_type = "local" source_type = "blank" volume_size = 1 } } ` + "`" + `` + "`" + `` + "`" + ` ### Instances and Security Groups When referencing a security group resource in an instance resource, always use the _name_ of the security group. If you specify the ID of the security group, Terraform will remove and reapply the security group upon each call. This is because the OpenStack Compute API returns the names of the associated security groups and not their IDs. Note the following example: ` + "`" + `` + "`" + `` + "`" + `hcl resource "openstack_networking_secgroup_v2" "sg_1" { name = "sg_1" } resource "openstack_compute_instance_v2" "foo" { name = "terraform-test" security_groups = ["${openstack_networking_secgroup_v2.sg_1.name}"] } ` + "`" + `` + "`" + `` + "`" + ` ### Instances and Ports Neutron Ports are a great feature and provide a lot of functionality. However, there are some notes to be aware of when mixing Instances and Ports:`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "access_ip_v4",
					Description: `The first detected Fixed IPv4 address.`,
				},
				resource.Attribute{
					Name:        "access_ip_v6",
					Description: `The first detected Fixed IPv6 address.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "security_groups",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "flavor_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "flavor_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "network/uuid",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "network/name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "network/port",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "network/fixed_ip_v4",
					Description: `The Fixed IPv4 address of the Instance on that network.`,
				},
				resource.Attribute{
					Name:        "network/fixed_ip_v6",
					Description: `The Fixed IPv6 address of the Instance on that network.`,
				},
				resource.Attribute{
					Name:        "network/mac",
					Description: `The MAC address of the NIC on that network.`,
				},
				resource.Attribute{
					Name:        "all_metadata",
					Description: `Contains all instance metadata, even metadata not set by Terraform. ## Notes ### Multiple Ephemeral Disks It's possible to specify multiple ` + "`" + `block_device` + "`" + ` entries to create an instance with multiple ephemeral (local) disks. In order to create multiple ephemeral disks, the sum of the total amount of ephemeral space must be less than or equal to what the chosen flavor supports. The following example shows how to create an instance with multiple ephemeral disks: ` + "`" + `` + "`" + `` + "`" + `hcl resource "openstack_compute_instance_v2" "foo" { name = "terraform-test" security_groups = ["default"] block_device { boot_index = 0 delete_on_termination = true destination_type = "local" source_type = "image" uuid = "<image uuid>" } block_device { boot_index = -1 delete_on_termination = true destination_type = "local" source_type = "blank" volume_size = 1 } block_device { boot_index = -1 delete_on_termination = true destination_type = "local" source_type = "blank" volume_size = 1 } } ` + "`" + `` + "`" + `` + "`" + ` ### Instances and Security Groups When referencing a security group resource in an instance resource, always use the _name_ of the security group. If you specify the ID of the security group, Terraform will remove and reapply the security group upon each call. This is because the OpenStack Compute API returns the names of the associated security groups and not their IDs. Note the following example: ` + "`" + `` + "`" + `` + "`" + `hcl resource "openstack_networking_secgroup_v2" "sg_1" { name = "sg_1" } resource "openstack_compute_instance_v2" "foo" { name = "terraform-test" security_groups = ["${openstack_networking_secgroup_v2.sg_1.name}"] } ` + "`" + `` + "`" + `` + "`" + ` ### Instances and Ports Neutron Ports are a great feature and provide a lot of functionality. However, there are some notes to be aware of when mixing Instances and Ports:`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "openstack_compute_interface_attach_v2",
			Category:         "Compute Resources",
			ShortDescription: `Attaches a Network Interface to an Instance.`,
			Description:      ``,
			Keywords: []string{
				"compute",
				"interface",
				"attach",
				"v2",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to create the interface attachment. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new attachment.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Required) The ID of the Instance to attach the Port or Network to.`,
				},
				resource.Attribute{
					Name:        "port_id",
					Description: `(Optional) The ID of the Port to attach to an Instance. _NOTE_: This option and ` + "`" + `network_id` + "`" + ` are mutually exclusive.`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `(Optional) The ID of the Network to attach to an Instance. A port will be created automatically. _NOTE_: This option and ` + "`" + `port_id` + "`" + ` are mutually exclusive.`,
				},
				resource.Attribute{
					Name:        "fixed_ip",
					Description: `(Optional) An IP address to assosciate with the port. _NOTE_: This option cannot be used with port_id. You must specifiy a network_id. The IP address must lie in a range on the supplied network. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "port_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `See Argument Reference above.`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "port_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `See Argument Reference above.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "openstack_compute_keypair_v2",
			Category:         "Compute Resources",
			ShortDescription: `Manages a V2 keypair resource within OpenStack.`,
			Description:      ``,
			Keywords: []string{
				"compute",
				"keypair",
				"v2",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 Compute client. Keypairs are associated with accounts, but a Compute client is needed to create one. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new keypair.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A unique name for the keypair. Changing this creates a new keypair.`,
				},
				resource.Attribute{
					Name:        "public_key",
					Description: `(Optional) A pregenerated OpenSSH-formatted public key. Changing this creates a new keypair. If a public key is not specified, then a public/private key pair will be automatically generated. If a pair is created, then destroying this resource means you will lose access to that keypair forever.`,
				},
				resource.Attribute{
					Name:        "value_specs",
					Description: `(Optional) Map of additional options. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "public_key",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "fingerprint",
					Description: `The fingerprint of the public key.`,
				},
				resource.Attribute{
					Name:        "private_key",
					Description: `The generated private key when no public key is specified. ## Import Keypairs can be imported using the ` + "`" + `name` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import openstack_compute_keypair_v2.my-keypair test-keypair ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "public_key",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "fingerprint",
					Description: `The fingerprint of the public key.`,
				},
				resource.Attribute{
					Name:        "private_key",
					Description: `The generated private key when no public key is specified. ## Import Keypairs can be imported using the ` + "`" + `name` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import openstack_compute_keypair_v2.my-keypair test-keypair ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "openstack_compute_secgroup_v2",
			Category:         "Compute Resources",
			ShortDescription: `Manages a V2 security group resource within OpenStack.`,
			Description:      ``,
			Keywords: []string{
				"compute",
				"secgroup",
				"v2",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 Compute client. A Compute client is needed to create a security group. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new security group.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A unique name for the security group. Changing this updates the ` + "`" + `name` + "`" + ` of an existing security group.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Required) A description for the security group. Changing this updates the ` + "`" + `description` + "`" + ` of an existing security group.`,
				},
				resource.Attribute{
					Name:        "rule",
					Description: `(Optional) A rule describing how the security group operates. The rule object structure is documented below. Changing this updates the security group rules. As shown in the example above, multiple rule blocks may be used. The ` + "`" + `rule` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "from_port",
					Description: `(Required) An integer representing the lower bound of the port range to open. Changing this creates a new security group rule.`,
				},
				resource.Attribute{
					Name:        "to_port",
					Description: `(Required) An integer representing the upper bound of the port range to open. Changing this creates a new security group rule.`,
				},
				resource.Attribute{
					Name:        "ip_protocol",
					Description: `(Required) The protocol type that will be allowed. Changing this creates a new security group rule.`,
				},
				resource.Attribute{
					Name:        "cidr",
					Description: `(Optional) Required if ` + "`" + `from_group_id` + "`" + ` or ` + "`" + `self` + "`" + ` is empty. The IP range that will be the source of network traffic to the security group. Use 0.0.0.0/0 to allow all IP addresses. Changing this creates a new security group rule. Cannot be combined with ` + "`" + `from_group_id` + "`" + ` or ` + "`" + `self` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "from_group_id",
					Description: `(Optional) Required if ` + "`" + `cidr` + "`" + ` or ` + "`" + `self` + "`" + ` is empty. The ID of a group from which to forward traffic to the parent group. Changing this creates a new security group rule. Cannot be combined with ` + "`" + `cidr` + "`" + ` or ` + "`" + `self` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "self",
					Description: `(Optional) Required if ` + "`" + `cidr` + "`" + ` and ` + "`" + `from_group_id` + "`" + ` is empty. If true, the security group itself will be added as a source to this ingress rule. Cannot be combined with ` + "`" + `cidr` + "`" + ` or ` + "`" + `from_group_id` + "`" + `. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "rule",
					Description: `See Argument Reference above. ## Notes ### ICMP Rules When using ICMP as the ` + "`" + `ip_protocol` + "`" + `, the ` + "`" + `from_port` + "`" + ` sets the ICMP _type_ and the ` + "`" + `to_port` + "`" + ` sets the ICMP _code_. To allow all ICMP types, set each value to ` + "`" + `-1` + "`" + `, like so: ` + "`" + `` + "`" + `` + "`" + `hcl rule { from_port = -1 to_port = -1 ip_protocol = "icmp" cidr = "0.0.0.0/0" } ` + "`" + `` + "`" + `` + "`" + ` A list of ICMP types and codes can be found [here](https://en.wikipedia.org/wiki/Internet_Control_Message_Protocol#Control_messages). ### Referencing Security Groups When referencing a security group in a configuration (for example, a configuration creates a new security group and then needs to apply it to an instance being created in the same configuration), it is currently recommended to reference the security group by name and not by ID, like this: ` + "`" + `` + "`" + `` + "`" + `hcl resource "openstack_compute_instance_v2" "test-server" { name = "tf-test" image_id = "ad091b52-742f-469e-8f3c-fd81cadf0743" flavor_id = "3" key_pair = "my_key_pair_name" security_groups = ["${openstack_compute_secgroup_v2.secgroup_1.name}"] } ` + "`" + `` + "`" + `` + "`" + ` ## Import Security Groups can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import openstack_compute_secgroup_v2.my_secgroup 1bc30ee9-9d5b-4c30-bdd5-7f1e663f5edf ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "rule",
					Description: `See Argument Reference above. ## Notes ### ICMP Rules When using ICMP as the ` + "`" + `ip_protocol` + "`" + `, the ` + "`" + `from_port` + "`" + ` sets the ICMP _type_ and the ` + "`" + `to_port` + "`" + ` sets the ICMP _code_. To allow all ICMP types, set each value to ` + "`" + `-1` + "`" + `, like so: ` + "`" + `` + "`" + `` + "`" + `hcl rule { from_port = -1 to_port = -1 ip_protocol = "icmp" cidr = "0.0.0.0/0" } ` + "`" + `` + "`" + `` + "`" + ` A list of ICMP types and codes can be found [here](https://en.wikipedia.org/wiki/Internet_Control_Message_Protocol#Control_messages). ### Referencing Security Groups When referencing a security group in a configuration (for example, a configuration creates a new security group and then needs to apply it to an instance being created in the same configuration), it is currently recommended to reference the security group by name and not by ID, like this: ` + "`" + `` + "`" + `` + "`" + `hcl resource "openstack_compute_instance_v2" "test-server" { name = "tf-test" image_id = "ad091b52-742f-469e-8f3c-fd81cadf0743" flavor_id = "3" key_pair = "my_key_pair_name" security_groups = ["${openstack_compute_secgroup_v2.secgroup_1.name}"] } ` + "`" + `` + "`" + `` + "`" + ` ## Import Security Groups can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import openstack_compute_secgroup_v2.my_secgroup 1bc30ee9-9d5b-4c30-bdd5-7f1e663f5edf ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "openstack_compute_servergroup_v2",
			Category:         "Compute Resources",
			ShortDescription: `Manages a V2 Server Group resource within OpenStack.`,
			Description:      ``,
			Keywords: []string{
				"compute",
				"servergroup",
				"v2",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 Compute client. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new server group.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A unique name for the server group. Changing this creates a new server group.`,
				},
				resource.Attribute{
					Name:        "policies",
					Description: `(Required) The set of policies for the server group. All policies are mutually exclusive. See the Policies section for more information. Changing this creates a new server group.`,
				},
				resource.Attribute{
					Name:        "value_specs",
					Description: `(Optional) Map of additional options. ## Policies`,
				},
				resource.Attribute{
					Name:        "affinity",
					Description: `All instances/servers launched in this group will be hosted on the same compute node.`,
				},
				resource.Attribute{
					Name:        "anti-affinity",
					Description: `All instances/servers launched in this group will be hosted on different compute nodes.`,
				},
				resource.Attribute{
					Name:        "soft-affinity",
					Description: `All instances/servers launched in this group will be hosted on the same compute node if possible, but if not possible they still will be scheduled instead of failure. To use this policy your OpenStack environment should support Compute service API 2.15 or above.`,
				},
				resource.Attribute{
					Name:        "soft-anti-affinity",
					Description: `All instances/servers launched in this group will be hosted on different compute nodes if possible, but if not possible they still will be scheduled instead of failure. To use this policy your OpenStack environment should support Compute service API 2.15 or above. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "policies",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "members",
					Description: `The instances that are part of this server group. ## Import Server Groups can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import openstack_compute_servergroup_v2.test-sg 1bc30ee9-9d5b-4c30-bdd5-7f1e663f5edf ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "policies",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "members",
					Description: `The instances that are part of this server group. ## Import Server Groups can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import openstack_compute_servergroup_v2.test-sg 1bc30ee9-9d5b-4c30-bdd5-7f1e663f5edf ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "openstack_compute_volume_attach_v2",
			Category:         "Compute Resources",
			ShortDescription: `Attaches a Block Storage Volume to an Instance.`,
			Description:      ``,
			Keywords: []string{
				"compute",
				"volume",
				"attach",
				"v2",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 Compute client. A Compute client is needed to create a volume attachment. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new volume attachment.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Required) The ID of the Instance to attach the Volume to.`,
				},
				resource.Attribute{
					Name:        "volume_id",
					Description: `(Required) The ID of the Volume to attach to an Instance.`,
				},
				resource.Attribute{
					Name:        "device",
					Description: `(Optional) The device of the volume attachment (ex: ` + "`" + `/dev/vdc` + "`" + `). _NOTE_: Being able to specify a device is dependent upon the hypervisor in use. There is a chance that the device specified in Terraform will not be the same device the hypervisor chose. If this happens, Terraform will wish to update the device upon subsequent applying which will cause the volume to be detached and reattached indefinitely. Please use with caution.`,
				},
				resource.Attribute{
					Name:        "multiattach",
					Description: `(Optional) Enable attachment of multiattach-capable volumes. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "volume_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "device",
					Description: `See Argument Reference above. _NOTE_: The correctness of this information is dependent upon the hypervisor in use. In some cases, this should not be used as an authoritative piece of information.`,
				},
				resource.Attribute{
					Name:        "multiattach",
					Description: `See Argument Reference above. ## Import Volume Attachments can be imported using the Instance ID and Volume ID separated by a slash, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import openstack_compute_volume_attach_v2.va_1 89c60255-9bd6-460c-822a-e2b959ede9d2/45670584-225f-46c3-b33e-6707b589b666 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "volume_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "device",
					Description: `See Argument Reference above. _NOTE_: The correctness of this information is dependent upon the hypervisor in use. In some cases, this should not be used as an authoritative piece of information.`,
				},
				resource.Attribute{
					Name:        "multiattach",
					Description: `See Argument Reference above. ## Import Volume Attachments can be imported using the Instance ID and Volume ID separated by a slash, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import openstack_compute_volume_attach_v2.va_1 89c60255-9bd6-460c-822a-e2b959ede9d2/45670584-225f-46c3-b33e-6707b589b666 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "openstack_containerinfra_cluster_v1",
			Category:         "Container Infra Resources",
			ShortDescription: `Manages a V1 Magnum cluster resource within OpenStack.`,
			Description:      ``,
			Keywords: []string{
				"container",
				"infra",
				"containerinfra",
				"cluster",
				"v1",
			},
			Arguments:  []resource.Argument{},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "openstack_containerinfra_clustertemplate_v1",
			Category:         "Container Infra Resources",
			ShortDescription: `Manages a V1 Magnum cluster template resource within OpenStack.`,
			Description:      ``,
			Keywords: []string{
				"container",
				"infra",
				"containerinfra",
				"clustertemplate",
				"v1",
			},
			Arguments:  []resource.Argument{},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "openstack_db_configuration_v1",
			Category:         "Database Resources",
			ShortDescription: `Manages a V1 DB configuration resource within OpenStack.`,
			Description:      ``,
			Keywords: []string{
				"database",
				"db",
				"configuration",
				"v1",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "region",
					Description: `(Required) The region in which to create the db instance. Changing this creates a new instance.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A unique name for the resource.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the resource.`,
				},
				resource.Attribute{
					Name:        "datastore",
					Description: `(Required) An array of database engine type and version. The datastore object structure is documented below. Changing this creates resource.`,
				},
				resource.Attribute{
					Name:        "configuration",
					Description: `(Optional) An array of configuration parameter name and value. Can be specified multiple times. The configuration object structure is documented below. The ` + "`" + `datastore` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) Database engine type to be used with this configuration. Changing this creates a new resource.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Required) Version of database engine type to be used with this configuration. Changing this creates a new resource. The ` + "`" + `configuration` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Configuration parameter name. Changing this creates a new resource.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Optional) Configuration parameter value. Changing this creates a new resource. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "datastore/type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "datastore/version",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "configuration/name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "configuration/value",
					Description: `See Argument Reference above.`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "datastore/type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "datastore/version",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "configuration/name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "configuration/value",
					Description: `See Argument Reference above.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "openstack_db_database_v1",
			Category:         "Database Resources",
			ShortDescription: `Manages a V1 database resource within OpenStack.`,
			Description:      ``,
			Keywords: []string{
				"database",
				"db",
				"v1",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A unique name for the resource.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Required) The ID for the database instance. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `Openstack region resource is created in.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `See Argument Reference above. ## Import Databases can be imported by using ` + "`" + `instance-id/db-name` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import openstack_db_database_v1.mydb 7b9e3cd3-00d9-449c-b074-8439f8e274fa/mydb ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "region",
					Description: `Openstack region resource is created in.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `See Argument Reference above. ## Import Databases can be imported by using ` + "`" + `instance-id/db-name` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import openstack_db_database_v1.mydb 7b9e3cd3-00d9-449c-b074-8439f8e274fa/mydb ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "openstack_db_instance_v1",
			Category:         "Database Resources",
			ShortDescription: `Manages a V1 DB instance resource within OpenStack.`,
			Description:      ``,
			Keywords: []string{
				"database",
				"db",
				"instance",
				"v1",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "region",
					Description: `(Required) The region in which to create the db instance. Changing this creates a new instance.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A unique name for the resource.`,
				},
				resource.Attribute{
					Name:        "flavor_id",
					Description: `(Required) The flavor ID of the desired flavor for the instance. Changing this creates new instance.`,
				},
				resource.Attribute{
					Name:        "configuration_id",
					Description: `(Optional) Configuration ID to be attached to the instance. Database instance will be rebooted when configuration is detached.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `(Required) Specifies the volume size in GB. Changing this creates new instance.`,
				},
				resource.Attribute{
					Name:        "datastore",
					Description: `(Required) An array of database engine type and version. The datastore object structure is documented below. Changing this creates a new instance.`,
				},
				resource.Attribute{
					Name:        "network",
					Description: `(Optional) An array of one or more networks to attach to the instance. The network object structure is documented below. Changing this creates a new instance.`,
				},
				resource.Attribute{
					Name:        "user",
					Description: `(Optional) An array of username, password, host and databases. The user object structure is documented below.`,
				},
				resource.Attribute{
					Name:        "database",
					Description: `(Optional) An array of database name, charset and collate. The database object structure is documented below. The ` + "`" + `datastore` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) Database engine type to be used in new instance. Changing this creates a new instance.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Required) Version of database engine type to be used in new instance. Changing this creates a new instance. The ` + "`" + `network` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `(Required unless ` + "`" + `port` + "`" + ` is provided) The network UUID to attach to the instance. Changing this creates a new instance.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Required unless ` + "`" + `uuid` + "`" + ` is provided) The port UUID of a network to attach to the instance. Changing this creates a new instance.`,
				},
				resource.Attribute{
					Name:        "fixed_ip_v4",
					Description: `(Optional) Specifies a fixed IPv4 address to be used on this network. Changing this creates a new instance.`,
				},
				resource.Attribute{
					Name:        "fixed_ip_v6",
					Description: `(Optional) Specifies a fixed IPv6 address to be used on this network. Changing this creates a new instance. The ` + "`" + `user` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Username to be created on new instance. Changing this creates a new instance.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional) User's password. Changing this creates a new instance.`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `(Optional) An ip address or % sign indicating what ip addresses can connect with this user credentials. Changing this creates a new instance.`,
				},
				resource.Attribute{
					Name:        "databases",
					Description: `(Optional) A list of databases that user will have access to. If not specified, user has access to all databases on th einstance. Changing this creates a new instance. The ` + "`" + `database` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Database to be created on new instance. Changing this creates a new instance.`,
				},
				resource.Attribute{
					Name:        "collate",
					Description: `(Optional) Database collation. Changing this creates a new instance.`,
				},
				resource.Attribute{
					Name:        "charset",
					Description: `(Optional) Database character set. Changing this creates a new instance. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "flavor_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "configuration_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "datastore/type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "datastore/version",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "network/uuid",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "network/port",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "network/fixed_ip_v4",
					Description: `The Fixed IPv4 address of the Instance on that network.`,
				},
				resource.Attribute{
					Name:        "network/fixed_ip_v6",
					Description: `The Fixed IPv6 address of the Instance on that`,
				},
				resource.Attribute{
					Name:        "database/name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "database/collate",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "database/charset",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "user/name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "user/password",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "user/databases",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "user/host",
					Description: `See Argument Reference above.`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "flavor_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "configuration_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "datastore/type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "datastore/version",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "network/uuid",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "network/port",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "network/fixed_ip_v4",
					Description: `The Fixed IPv4 address of the Instance on that network.`,
				},
				resource.Attribute{
					Name:        "network/fixed_ip_v6",
					Description: `The Fixed IPv6 address of the Instance on that`,
				},
				resource.Attribute{
					Name:        "database/name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "database/collate",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "database/charset",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "user/name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "user/password",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "user/databases",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "user/host",
					Description: `See Argument Reference above.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "openstack_db_user_v1",
			Category:         "Database Resources",
			ShortDescription: `Manages a V1 database user resource within OpenStack.`,
			Description:      ``,
			Keywords: []string{
				"database",
				"db",
				"user",
				"v1",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A unique name for the resource.`,
				},
				resource.Attribute{
					Name:        "instance",
					Description: `(Required) The ID for the database instance.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Required) User's password.`,
				},
				resource.Attribute{
					Name:        "databases",
					Description: `(Optional) A list of database user should have access to. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `Openstack region resource is created in.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "instance",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "databases",
					Description: `See Argument Reference above.`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "region",
					Description: `Openstack region resource is created in.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "instance",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "databases",
					Description: `See Argument Reference above.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "openstack_dns_recordset_v2",
			Category:         "DNS Resources",
			ShortDescription: `Manages a DNS record set in the OpenStack DNS Service`,
			Description:      ``,
			Keywords: []string{
				"dns",
				"recordset",
				"v2",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 DNS client. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new DNS record set.`,
				},
				resource.Attribute{
					Name:        "zone_id",
					Description: `(Required) The ID of the zone in which to create the record set. Changing this creates a new DNS record set.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the record set. Note the ` + "`" + `.` + "`" + ` at the end of the name. Changing this creates a new DNS record set.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) The type of record set. Examples: "A", "MX". Changing this creates a new DNS record set.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `(Optional) The time to live (TTL) of the record set.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A description of the record set.`,
				},
				resource.Attribute{
					Name:        "records",
					Description: `(Optional) An array of DNS records. _Note:_ if an IPv6 address contains brackets (` + "`" + `[ ]` + "`" + `), the brackets will be stripped and the modified address will be recorded in the state.`,
				},
				resource.Attribute{
					Name:        "value_specs",
					Description: `(Optional) Map of additional options. Changing this creates a new record set. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "records",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "zone_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "value_specs",
					Description: `See Argument Reference above. ## Import This resource can be imported by specifying the zone ID and recordset ID, separated by a forward slash. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import openstack_dns_recordset_v2.recordset_1 <zone_id>/<recordset_id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "records",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "zone_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "value_specs",
					Description: `See Argument Reference above. ## Import This resource can be imported by specifying the zone ID and recordset ID, separated by a forward slash. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import openstack_dns_recordset_v2.recordset_1 <zone_id>/<recordset_id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "openstack_dns_zone_v2",
			Category:         "DNS Resources",
			ShortDescription: `Manages a DNS zone in the OpenStack DNS Service`,
			Description:      ``,
			Keywords: []string{
				"dns",
				"zone",
				"v2",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 Compute client. Keypairs are associated with accounts, but a Compute client is needed to create one. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new DNS zone.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the zone. Note the ` + "`" + `.` + "`" + ` at the end of the name. Changing this creates a new DNS zone.`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `(Optional) The email contact for the zone record.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) The type of zone. Can either be ` + "`" + `PRIMARY` + "`" + ` or ` + "`" + `SECONDARY` + "`" + `. Changing this creates a new zone.`,
				},
				resource.Attribute{
					Name:        "attributes",
					Description: `(Optional) Attributes for the DNS Service scheduler. Changing this creates a new zone.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `(Optional) The time to live (TTL) of the zone.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A description of the zone.`,
				},
				resource.Attribute{
					Name:        "masters",
					Description: `(Optional) An array of master DNS servers. For when ` + "`" + `type` + "`" + ` is ` + "`" + `SECONDARY` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "value_specs",
					Description: `(Optional) Map of additional options. Changing this creates a new zone. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "attributes",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "masters",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "value_specs",
					Description: `See Argument Reference above. ## Import This resource can be imported by specifying the zone ID: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import openstack_dns_zone_v2.zone_1 <zone_id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "attributes",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "masters",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "value_specs",
					Description: `See Argument Reference above. ## Import This resource can be imported by specifying the zone ID: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import openstack_dns_zone_v2.zone_1 <zone_id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "openstack_fw_firewall_v1",
			Category:         "Firewall Resources",
			ShortDescription: `Manages a v1 firewall resource within OpenStack.`,
			Description:      ``,
			Keywords: []string{
				"firewall",
				"fw",
				"v1",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the v1 networking client. A networking client is needed to create a firewall. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new firewall.`,
				},
				resource.Attribute{
					Name:        "policy_id",
					Description: `(Required) The policy resource id for the firewall. Changing this updates the ` + "`" + `policy_id` + "`" + ` of an existing firewall.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) A name for the firewall. Changing this updates the ` + "`" + `name` + "`" + ` of an existing firewall.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Required) A description for the firewall. Changing this updates the ` + "`" + `description` + "`" + ` of an existing firewall.`,
				},
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `(Optional) Administrative up/down status for the firewall (must be "true" or "false" if provided - defaults to "true"). Changing this updates the ` + "`" + `admin_state_up` + "`" + ` of an existing firewall.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Optional) The owner of the floating IP. Required if admin wants to create a firewall for another tenant. Changing this creates a new firewall.`,
				},
				resource.Attribute{
					Name:        "associated_routers",
					Description: `(Optional) Router(s) to associate this firewall instance with. Must be a list of strings. Changing this updates the associated routers of an existing firewall. Conflicts with ` + "`" + `no_routers` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "no_routers",
					Description: `(Optional) Should this firewall not be associated with any routers (must be "true" or "false" if provide - defaults to "false"). Conflicts with ` + "`" + `associated_routers` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "value_specs",
					Description: `(Optional) Map of additional options. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "policy_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "associated_routers",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "no_routers",
					Description: `See Argument Reference above. ## Import Firewalls can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import openstack_fw_firewall_v1.firewall_1 c9e39fb2-ce20-46c8-a964-25f3898c7a97 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "policy_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "associated_routers",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "no_routers",
					Description: `See Argument Reference above. ## Import Firewalls can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import openstack_fw_firewall_v1.firewall_1 c9e39fb2-ce20-46c8-a964-25f3898c7a97 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "openstack_fw_policy_v1",
			Category:         "Firewall Resources",
			ShortDescription: `Manages a v1 firewall policy resource within OpenStack.`,
			Description:      ``,
			Keywords: []string{
				"firewall",
				"fw",
				"policy",
				"v1",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the v1 networking client. A networking client is needed to create a firewall policy. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new firewall policy.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) A name for the firewall policy. Changing this updates the ` + "`" + `name` + "`" + ` of an existing firewall policy.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A description for the firewall policy. Changing this updates the ` + "`" + `description` + "`" + ` of an existing firewall policy.`,
				},
				resource.Attribute{
					Name:        "rules",
					Description: `(Optional) An array of one or more firewall rules that comprise the policy. Changing this results in adding/removing rules from the existing firewall policy.`,
				},
				resource.Attribute{
					Name:        "audited",
					Description: `(Optional) Audit status of the firewall policy (must be "true" or "false" if provided - defaults to "false"). This status is set to "false" whenever the firewall policy or any of its rules are changed. Changing this updates the ` + "`" + `audited` + "`" + ` status of an existing firewall policy.`,
				},
				resource.Attribute{
					Name:        "shared",
					Description: `(Optional) Sharing status of the firewall policy (must be "true" or "false" if provided). If this is "true" the policy is visible to, and can be used in, firewalls in other tenants. Changing this updates the ` + "`" + `shared` + "`" + ` status of an existing firewall policy. Only administrative users can specify if the policy should be shared.`,
				},
				resource.Attribute{
					Name:        "value_specs",
					Description: `(Optional) Map of additional options. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "audited",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "shared",
					Description: `See Argument Reference above. ## Import Firewall Policies can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import openstack_fw_policy_v1.policy_1 07f422e6-c596-474b-8b94-fe2c12506ce0 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "audited",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "shared",
					Description: `See Argument Reference above. ## Import Firewall Policies can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import openstack_fw_policy_v1.policy_1 07f422e6-c596-474b-8b94-fe2c12506ce0 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "openstack_fw_rule_v1",
			Category:         "Firewall Resources",
			ShortDescription: `Manages a v1 firewall rule resource within OpenStack.`,
			Description:      ``,
			Keywords: []string{
				"firewall",
				"fw",
				"rule",
				"v1",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the v1 Compute client. A Compute client is needed to create a firewall rule. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new firewall rule.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) A unique name for the firewall rule. Changing this updates the ` + "`" + `name` + "`" + ` of an existing firewall rule.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A description for the firewall rule. Changing this updates the ` + "`" + `description` + "`" + ` of an existing firewall rule.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Required) The protocol type on which the firewall rule operates. Valid values are: ` + "`" + `tcp` + "`" + `, ` + "`" + `udp` + "`" + `, ` + "`" + `icmp` + "`" + `, and ` + "`" + `any` + "`" + `. Changing this updates the ` + "`" + `protocol` + "`" + ` of an existing firewall rule.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Required) Action to be taken ( must be "allow" or "deny") when the firewall rule matches. Changing this updates the ` + "`" + `action` + "`" + ` of an existing firewall rule.`,
				},
				resource.Attribute{
					Name:        "ip_version",
					Description: `(Optional) IP version, either 4 (default) or 6. Changing this updates the ` + "`" + `ip_version` + "`" + ` of an existing firewall rule.`,
				},
				resource.Attribute{
					Name:        "source_ip_address",
					Description: `(Optional) The source IP address on which the firewall rule operates. Changing this updates the ` + "`" + `source_ip_address` + "`" + ` of an existing firewall rule.`,
				},
				resource.Attribute{
					Name:        "destination_ip_address",
					Description: `(Optional) The destination IP address on which the firewall rule operates. Changing this updates the ` + "`" + `destination_ip_address` + "`" + ` of an existing firewall rule.`,
				},
				resource.Attribute{
					Name:        "source_port",
					Description: `(Optional) The source port on which the firewall rule operates. Changing this updates the ` + "`" + `source_port` + "`" + ` of an existing firewall rule.`,
				},
				resource.Attribute{
					Name:        "destination_port",
					Description: `(Optional) The destination port on which the firewall rule operates. Changing this updates the ` + "`" + `destination_port` + "`" + ` of an existing firewall rule.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) Enabled status for the firewall rule (must be "true" or "false" if provided - defaults to "true"). Changing this updates the ` + "`" + `enabled` + "`" + ` status of an existing firewall rule.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Optional) The owner of the firewall rule. Required if admin wants to create a firewall rule for another tenant. Changing this creates a new firewall rule.`,
				},
				resource.Attribute{
					Name:        "value_specs",
					Description: `(Optional) Map of additional options. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "ip_version",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "source_ip_address",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "destination_ip_address",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "source_port",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "destination_port",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above. ## Import Firewall Rules can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import openstack_fw_rule_v1.rule_1 8dbc0c28-e49c-463f-b712-5c5d1bbac327 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "ip_version",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "source_ip_address",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "destination_ip_address",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "source_port",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "destination_port",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above. ## Import Firewall Rules can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import openstack_fw_rule_v1.rule_1 8dbc0c28-e49c-463f-b712-5c5d1bbac327 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "openstack_identity_application_credential_v3",
			Category:         "Identity Resources",
			ShortDescription: `Manages a V3 Application Credential resource within OpenStack Keystone.`,
			Description:      ``,
			Keywords: []string{
				"identity",
				"application",
				"credential",
				"v3",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V3 Keystone client. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new application credential.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A name of the application credential. Changing this creates a new application credential.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A description of the application credential. Changing this creates a new application credential.`,
				},
				resource.Attribute{
					Name:        "unrestricted",
					Description: `(Optional) A flag indicating whether the application credential may be used for creation or destruction of other application credentials or trusts. Changing this creates a new application credential.`,
				},
				resource.Attribute{
					Name:        "secret",
					Description: `(Optional) The secret for the application credential. If omitted, it will be generated by the server. Changing this creates a new application credential.`,
				},
				resource.Attribute{
					Name:        "roles",
					Description: `(Optional) A collection of one or more role names, which this application credential has to be associated with its project. If omitted, all the current user's roles within the scoped project will be inherited by a new application credential. Changing this creates a new application credential.`,
				},
				resource.Attribute{
					Name:        "expires_at",
					Description: `(Optional) The expiration time of the application credential in the RFC3339 timestamp format (e.g. ` + "`" + `2019-03-09T12:58:49Z` + "`" + `). If omitted, an application credential will never expire. Changing this creates a new application credential. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "unrestricted",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "secret",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "roles",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "expires_at",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `The ID of the project the application credential was created for and that authentication requests using this application credential will be scoped to. ## Import Application Credentials can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import openstack_identity_application_credential_v3.application_credential_1 c17304b7-0953-4738-abb0-67005882b0a0 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "unrestricted",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "secret",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "roles",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "expires_at",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `The ID of the project the application credential was created for and that authentication requests using this application credential will be scoped to. ## Import Application Credentials can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import openstack_identity_application_credential_v3.application_credential_1 c17304b7-0953-4738-abb0-67005882b0a0 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "openstack_identity_project_v3",
			Category:         "Identity Resources",
			ShortDescription: `Manages a V3 Project resource within OpenStack Keystone.`,
			Description:      ``,
			Keywords: []string{
				"identity",
				"project",
				"v3",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A description of the project.`,
				},
				resource.Attribute{
					Name:        "domain_id",
					Description: `(Optional) The domain this project belongs to.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) Whether the project is enabled or disabled. Valid values are ` + "`" + `true` + "`" + ` and ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "is_domain",
					Description: `(Optional) Whether this project is a domain. Valid values are ` + "`" + `true` + "`" + ` and ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the project.`,
				},
				resource.Attribute{
					Name:        "parent_id",
					Description: `(Optional) The parent of this project.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V3 Keystone client. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new User. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "domain_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "parent_id",
					Description: `See Argument Reference above. ## Import Projects can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import openstack_identity_project_v3.project_1 89c60255-9bd6-460c-822a-e2b959ede9d2 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "domain_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "parent_id",
					Description: `See Argument Reference above. ## Import Projects can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import openstack_identity_project_v3.project_1 89c60255-9bd6-460c-822a-e2b959ede9d2 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "openstack_identity_role_assignment_v3",
			Category:         "Identity Resources",
			ShortDescription: `Manages a V3 Role assignment within OpenStack Keystone.`,
			Description:      ``,
			Keywords: []string{
				"identity",
				"role",
				"assignment",
				"v3",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "domain_id",
					Description: `(Optional; Required if ` + "`" + `project_id` + "`" + ` is empty) The domain to assign the role in.`,
				},
				resource.Attribute{
					Name:        "group_id",
					Description: `(Optional; Required if ` + "`" + `user_id` + "`" + ` is empty) The group to assign the role to.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Optional; Required if ` + "`" + `domain_id` + "`" + ` is empty) The project to assign the role in.`,
				},
				resource.Attribute{
					Name:        "user_id",
					Description: `(Optional; Required if ` + "`" + `group_id` + "`" + ` is empty) The user to assign the role to.`,
				},
				resource.Attribute{
					Name:        "role_id",
					Description: `(Required) The role to assign. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "domain_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "group_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "user_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "role_id",
					Description: `See Argument Reference above.`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "domain_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "group_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "user_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "role_id",
					Description: `See Argument Reference above.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "openstack_identity_role_v3",
			Category:         "Identity Resources",
			ShortDescription: `Manages a V3 Role resource within OpenStack Keystone.`,
			Description:      ``,
			Keywords: []string{
				"identity",
				"role",
				"v3",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `The name of the role.`,
				},
				resource.Attribute{
					Name:        "domain_id",
					Description: `(Optional) The domain the role belongs to.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V3 Keystone client. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new Role. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "domain_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above. ## Import Roles can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import openstack_identity_role_v3.role_1 89c60255-9bd6-460c-822a-e2b959ede9d2 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "domain_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above. ## Import Roles can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import openstack_identity_role_v3.role_1 89c60255-9bd6-460c-822a-e2b959ede9d2 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "openstack_identity_user_v3",
			Category:         "Identity Resources",
			ShortDescription: `Manages a V3 User resource within OpenStack Keystone.`,
			Description:      ``,
			Keywords: []string{
				"identity",
				"user",
				"v3",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A description of the user.`,
				},
				resource.Attribute{
					Name:        "default_project_id",
					Description: `(Optional) The default project this user belongs to.`,
				},
				resource.Attribute{
					Name:        "domain_id",
					Description: `(Optional) The domain this user belongs to.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) Whether the user is enabled or disabled. Valid values are ` + "`" + `true` + "`" + ` and ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "extra",
					Description: `(Optional) Free-form key/value pairs of extra information.`,
				},
				resource.Attribute{
					Name:        "ignore_change_password_upon_first_use",
					Description: `(Optional) User will not have to change their password upon first use. Valid values are ` + "`" + `true` + "`" + ` and ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ignore_password_expiry",
					Description: `(Optional) User's password will not expire. Valid values are ` + "`" + `true` + "`" + ` and ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ignore_lockout_failure_attempts",
					Description: `(Optional) User will not have a failure lockout placed on their account. Valid values are ` + "`" + `true` + "`" + ` and ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "multi_factor_auth_enabled",
					Description: `(Optional) Whether to enable multi-factor authentication. Valid values are ` + "`" + `true` + "`" + ` and ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "multi_factor_auth_rule",
					Description: `(Optional) A multi-factor authentication rule. The structure is documented below. Please see the [Ocata release notes](https://docs.openstack.org/releasenotes/keystone/ocata.html) for more information on how to use mulit-factor rules.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the user.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional) The password for the user.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V3 Keystone client. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new User. The ` + "`" + `multi_factor_auth_rule` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "rule",
					Description: `(Required) A list of authentication plugins that the user must authenticate with. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "domain_id",
					Description: `See Argument Reference above. ## Import Users can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import openstack_identity_user_v3.user_1 89c60255-9bd6-460c-822a-e2b959ede9d2 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "domain_id",
					Description: `See Argument Reference above. ## Import Users can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import openstack_identity_user_v3.user_1 89c60255-9bd6-460c-822a-e2b959ede9d2 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "openstack_images_image_v2",
			Category:         "Images Resources",
			ShortDescription: `Manages a V2 Image resource within OpenStack Glance.`,
			Description:      ``,
			Keywords: []string{
				"images",
				"image",
				"v2",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "container_format",
					Description: `(Required) The container format. Must be one of "ami", "ari", "aki", "bare", "ovf".`,
				},
				resource.Attribute{
					Name:        "disk_format",
					Description: `(Required) The disk format. Must be one of "ami", "ari", "aki", "vhd", "vmdk", "raw", "qcow2", "vdi", "iso".`,
				},
				resource.Attribute{
					Name:        "local_file_path",
					Description: `(Optional) This is the filepath of the raw image file that will be uploaded to Glance. Conflicts with ` + "`" + `image_source_url` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "image_cache_path",
					Description: `(Optional) This is the directory where the images will be downloaded. Images will be stored with a filename corresponding to the url's md5 hash. Defaults to "$HOME/.terraform/image_cache"`,
				},
				resource.Attribute{
					Name:        "image_source_url",
					Description: `(Optional) This is the url of the raw image that will be downloaded in the ` + "`" + `image_cache_path` + "`" + ` before being uploaded to Glance. Glance is able to download image from internet but the ` + "`" + `gophercloud` + "`" + ` library does not yet provide a way to do so. Conflicts with ` + "`" + `local_file_path` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "min_disk_gb",
					Description: `(Optional) Amount of disk space (in GB) required to boot image. Defaults to 0.`,
				},
				resource.Attribute{
					Name:        "min_ram_mb",
					Description: `(Optional) Amount of ram (in MB) required to boot image. Defauts to 0.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the image.`,
				},
				resource.Attribute{
					Name:        "properties",
					Description: `(Optional) A map of key/value pairs to set freeform information about an image. See the "Notes" section for further information about properties.`,
				},
				resource.Attribute{
					Name:        "protected",
					Description: `(Optional) If true, image will not be deletable. Defaults to false.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 Glance client. A Glance client is needed to create an Image that can be used with a compute instance. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new Image.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) The tags of the image. It must be a list of strings. At this time, it is not possible to delete all tags of an image.`,
				},
				resource.Attribute{
					Name:        "verify_checksum",
					Description: `(Optional) If false, the checksum will not be verified once the image is finished uploading. Defaults to true.`,
				},
				resource.Attribute{
					Name:        "visibility",
					Description: `(Optional) The visibility of the image. Must be one of "public", "private", "community", or "shared". The ability to set the visibility depends upon the configuration of the OpenStack cloud. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "checksum",
					Description: `The checksum of the data associated with the image.`,
				},
				resource.Attribute{
					Name:        "container_format",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `The date the image was created.`,
				},
				resource.Attribute{
					Name:        "disk_format",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "file",
					Description: `the trailing path after the glance endpoint that represent the location of the image or the path to retrieve it.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `A unique ID assigned by Glance.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `The metadata associated with the image. Image metadata allow for meaningfully define the image properties and tags. See https://docs.openstack.org/glance/latest/user/metadefs-concepts.html.`,
				},
				resource.Attribute{
					Name:        "min_disk_gb",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "min_ram_mb",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "owner",
					Description: `The id of the openstack user who owns the image.`,
				},
				resource.Attribute{
					Name:        "properties",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "protected",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "schema",
					Description: `The path to the JSON-schema that represent the image or image`,
				},
				resource.Attribute{
					Name:        "size_bytes",
					Description: `The size in bytes of the data associated with the image.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the image. It can be "queued", "active" or "saving".`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `The date the image was last updated.`,
				},
				resource.Attribute{
					Name:        "update_at",
					Description: `(`,
				},
				resource.Attribute{
					Name:        "visibility",
					Description: `See Argument Reference above. ## Notes ### Properties This resource supports the ability to add properties to a resource during creation as well as add, update, and delete properties during an update of this resource. Newer versions of OpenStack are adding some read-only properties to each image. These properties start with the prefix ` + "`" + `os_` + "`" + `. If these properties are detected, this resource will automatically reconcile these with the user-provided properties. In addition, the ` + "`" + `direct_url` + "`" + ` property is also automatically reconciled if the Image Service set it. ## Import Images can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import openstack_images_image_v2.rancheros 89c60255-9bd6-460c-822a-e2b959ede9d2 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "checksum",
					Description: `The checksum of the data associated with the image.`,
				},
				resource.Attribute{
					Name:        "container_format",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `The date the image was created.`,
				},
				resource.Attribute{
					Name:        "disk_format",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "file",
					Description: `the trailing path after the glance endpoint that represent the location of the image or the path to retrieve it.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `A unique ID assigned by Glance.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `The metadata associated with the image. Image metadata allow for meaningfully define the image properties and tags. See https://docs.openstack.org/glance/latest/user/metadefs-concepts.html.`,
				},
				resource.Attribute{
					Name:        "min_disk_gb",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "min_ram_mb",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "owner",
					Description: `The id of the openstack user who owns the image.`,
				},
				resource.Attribute{
					Name:        "properties",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "protected",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "schema",
					Description: `The path to the JSON-schema that represent the image or image`,
				},
				resource.Attribute{
					Name:        "size_bytes",
					Description: `The size in bytes of the data associated with the image.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the image. It can be "queued", "active" or "saving".`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `The date the image was last updated.`,
				},
				resource.Attribute{
					Name:        "update_at",
					Description: `(`,
				},
				resource.Attribute{
					Name:        "visibility",
					Description: `See Argument Reference above. ## Notes ### Properties This resource supports the ability to add properties to a resource during creation as well as add, update, and delete properties during an update of this resource. Newer versions of OpenStack are adding some read-only properties to each image. These properties start with the prefix ` + "`" + `os_` + "`" + `. If these properties are detected, this resource will automatically reconcile these with the user-provided properties. In addition, the ` + "`" + `direct_url` + "`" + ` property is also automatically reconciled if the Image Service set it. ## Import Images can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import openstack_images_image_v2.rancheros 89c60255-9bd6-460c-822a-e2b959ede9d2 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "openstack_lb_l7policy_v2",
			Category:         "Load Balancer Resources",
			ShortDescription: `Manages a V2 L7 Policy resource within OpenStack.`,
			Description:      ``,
			Keywords: []string{
				"load",
				"balancer",
				"lb",
				"l7policy",
				"v2",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 Networking client. A Networking client is needed to create an . If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new L7 Policy.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Optional) Required for admins. The UUID of the tenant who owns the L7 Policy. Only administrative users can specify a tenant UUID other than their own. Changing this creates a new L7 Policy.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Human-readable name for the L7 Policy. Does not have to be unique.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Human-readable description for the L7 Policy.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Required) The L7 Policy action - can either be REDIRECT\_TO\_POOL, REDIRECT\_TO\_URL or REJECT.`,
				},
				resource.Attribute{
					Name:        "listener_id",
					Description: `(Required) The Listener on which the L7 Policy will be associated with. Changing this creates a new L7 Policy.`,
				},
				resource.Attribute{
					Name:        "position",
					Description: `(Optional) The position of this policy on the listener. Positions start at 1.`,
				},
				resource.Attribute{
					Name:        "redirect_pool_id",
					Description: `(Optional) Requests matching this policy will be redirected to the pool with this ID. Only valid if action is REDIRECT\_TO\_POOL.`,
				},
				resource.Attribute{
					Name:        "redirect_url",
					Description: `(Optional) Requests matching this policy will be redirected to this URL. Only valid if action is REDIRECT\_TO\_URL.`,
				},
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `(Optional) The administrative state of the L7 Policy. A valid value is true (UP) or false (DOWN). ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique ID for the L7 {olicy.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "listener_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "position",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "redirect_pool_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "redirect_url",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `See Argument Reference above. ## Import Load Balancer L7 Policy can be imported using the L7 Policy ID, e.g.: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import openstack_lb_l7policy_v2.l7policy_1 8a7a79c2-cf17-4e65-b2ae-ddc8bfcf6c74 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "id",
					Description: `The unique ID for the L7 {olicy.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "listener_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "position",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "redirect_pool_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "redirect_url",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `See Argument Reference above. ## Import Load Balancer L7 Policy can be imported using the L7 Policy ID, e.g.: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import openstack_lb_l7policy_v2.l7policy_1 8a7a79c2-cf17-4e65-b2ae-ddc8bfcf6c74 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "openstack_lb_l7rule_v2",
			Category:         "Load Balancer Resources",
			ShortDescription: `Manages a V2 l7rule resource within OpenStack.`,
			Description:      ``,
			Keywords: []string{
				"load",
				"balancer",
				"lb",
				"l7rule",
				"v2",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 Networking client. A Networking client is needed to create an . If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new L7 Rule.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Optional) Required for admins. The UUID of the tenant who owns the L7 Rule. Only administrative users can specify a tenant UUID other than their own. Changing this creates a new L7 Rule.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Human-readable description for the L7 Rule.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) The L7 Rule type - can either be COOKIE, FILE\_TYPE, HEADER, HOST\_NAME or PATH.`,
				},
				resource.Attribute{
					Name:        "compare_type",
					Description: `(Required) The comparison type for the L7 rule - can either be CONTAINS, STARTS\_WITH, ENDS_WITH, EQUAL_TO or REGEX`,
				},
				resource.Attribute{
					Name:        "l7policy_id",
					Description: `(Required) The ID of the L7 Policy to query. Changing this creates a new L7 Rule.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) The value to use for the comparison. For example, the file type to compare.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Optional) The key to use for the comparison. For example, the name of the cookie to evaluate. Valid when ` + "`" + `type` + "`" + ` is set to COOKIE or HEADER.`,
				},
				resource.Attribute{
					Name:        "invert",
					Description: `(Optional) When true the logic of the rule is inverted. For example, with invert true, equal to would become not equal to. Default is false.`,
				},
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `(Optional) The administrative state of the L7 Rule. A valid value is true (UP) or false (DOWN). ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique ID for the L7 Rule.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "compare_type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "l7policy_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "invert",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "listener_id",
					Description: `The ID of the Listener owning this resource. ## Import Load Balancer L7 Rule can be imported using the L7 Policy ID and L7 Rule ID separated by a slash, e.g.: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import openstack_lb_l7rule_v2.l7rule_1 e0bd694a-abbe-450e-b329-0931fd1cc5eb/4086b0c9-b18c-4d1c-b6b8-4c56c3ad2a9e ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "id",
					Description: `The unique ID for the L7 Rule.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "compare_type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "l7policy_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "invert",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "listener_id",
					Description: `The ID of the Listener owning this resource. ## Import Load Balancer L7 Rule can be imported using the L7 Policy ID and L7 Rule ID separated by a slash, e.g.: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import openstack_lb_l7rule_v2.l7rule_1 e0bd694a-abbe-450e-b329-0931fd1cc5eb/4086b0c9-b18c-4d1c-b6b8-4c56c3ad2a9e ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "openstack_lb_listener_v2",
			Category:         "Load Balancer Resources",
			ShortDescription: `Manages a V2 listener resource within OpenStack.`,
			Description:      ``,
			Keywords: []string{
				"load",
				"balancer",
				"lb",
				"listener",
				"v2",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 Networking client. A Networking client is needed to create an . If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new Listener.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Required) The protocol - can either be TCP, HTTP, HTTPS or TERMINATED_HTTPS. Changing this creates a new Listener.`,
				},
				resource.Attribute{
					Name:        "protocol_port",
					Description: `(Required) The port on which to listen for client traffic. Changing this creates a new Listener.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Optional) Required for admins. The UUID of the tenant who owns the Listener. Only administrative users can specify a tenant UUID other than their own. Changing this creates a new Listener.`,
				},
				resource.Attribute{
					Name:        "loadbalancer_id",
					Description: `(Required) The load balancer on which to provision this Listener. Changing this creates a new Listener.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Human-readable name for the Listener. Does not have to be unique.`,
				},
				resource.Attribute{
					Name:        "default_pool_id",
					Description: `(Optional) The ID of the default pool with which the Listener is associated.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Human-readable description for the Listener.`,
				},
				resource.Attribute{
					Name:        "connection_limit",
					Description: `(Optional) The maximum number of connections allowed for the Listener.`,
				},
				resource.Attribute{
					Name:        "default_tls_container_ref",
					Description: `(Optional) A reference to a Barbican Secrets container which stores TLS information. This is required if the protocol is ` + "`" + `TERMINATED_HTTPS` + "`" + `. See [here](https://wiki.openstack.org/wiki/Network/LBaaS/docs/how-to-create-tls-loadbalancer) for more information.`,
				},
				resource.Attribute{
					Name:        "sni_container_refs",
					Description: `(Optional) A list of references to Barbican Secrets containers which store SNI information. See [here](https://wiki.openstack.org/wiki/Network/LBaaS/docs/how-to-create-tls-loadbalancer) for more information.`,
				},
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `(Optional) The administrative state of the Listener. A valid value is true (UP) or false (DOWN). ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique ID for the Listener.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "protocol_port",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "default_port_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "connection_limit",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "default_tls_container_ref",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "sni_container_refs",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `See Argument Reference above. ## Import Load Balancer Listener can be imported using the Listener ID, e.g.: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import openstack_lb_listener_v2.listener_1 b67ce64e-8b26-405d-afeb-4a078901f15a ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "id",
					Description: `The unique ID for the Listener.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "protocol_port",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "default_port_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "connection_limit",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "default_tls_container_ref",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "sni_container_refs",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `See Argument Reference above. ## Import Load Balancer Listener can be imported using the Listener ID, e.g.: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import openstack_lb_listener_v2.listener_1 b67ce64e-8b26-405d-afeb-4a078901f15a ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "openstack_lb_loadbalancer_v2",
			Category:         "Load Balancer Resources",
			ShortDescription: `Manages a V2 loadbalancer resource within OpenStack.`,
			Description:      ``,
			Keywords: []string{
				"load",
				"balancer",
				"lb",
				"loadbalancer",
				"v2",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 Networking client. A Networking client is needed to create an LB member. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new LB member.`,
				},
				resource.Attribute{
					Name:        "vip_subnet_id",
					Description: `(Required) The network on which to allocate the Loadbalancer's address. A tenant can only create Loadbalancers on networks authorized by policy (e.g. networks that belong to them or networks that are shared). Changing this creates a new loadbalancer.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Human-readable name for the Loadbalancer. Does not have to be unique.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Human-readable description for the Loadbalancer.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Optional) Required for admins. The UUID of the tenant who owns the Loadbalancer. Only administrative users can specify a tenant UUID other than their own. Changing this creates a new loadbalancer.`,
				},
				resource.Attribute{
					Name:        "vip_address",
					Description: `(Optional) The ip address of the load balancer. Changing this creates a new loadbalancer.`,
				},
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `(Optional) The administrative state of the Loadbalancer. A valid value is true (UP) or false (DOWN).`,
				},
				resource.Attribute{
					Name:        "flavor",
					Description: `(Optional) The UUID of a flavor. Changing this creates a new loadbalancer.`,
				},
				resource.Attribute{
					Name:        "loadbalancer_provider",
					Description: `(Optional) The name of the provider. Changing this creates a new loadbalancer.`,
				},
				resource.Attribute{
					Name:        "security_group_ids",
					Description: `(Optional) A list of security group IDs to apply to the loadbalancer. The security groups must be specified by ID and not name (as opposed to how they are configured with the Compute Instance). ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "vip_subnet_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "vip_address",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "flavor",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "loadbalancer_provider",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "security_group_ids",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "vip_port_id",
					Description: `The Port ID of the Load Balancer IP. ## Import Load Balancer can be imported using the Load Balancer ID, e.g.: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import openstack_lb_loadbalancer_v2.loadbalancer_1 19bcfdc7-c521-4a7e-9459-6750bd16df76 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "vip_subnet_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "vip_address",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "flavor",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "loadbalancer_provider",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "security_group_ids",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "vip_port_id",
					Description: `The Port ID of the Load Balancer IP. ## Import Load Balancer can be imported using the Load Balancer ID, e.g.: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import openstack_lb_loadbalancer_v2.loadbalancer_1 19bcfdc7-c521-4a7e-9459-6750bd16df76 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "openstack_lb_member_v1",
			Category:         "Load Balancer Resources",
			ShortDescription: `Manages a V1 load balancer member resource within OpenStack.`,
			Description:      ``,
			Keywords: []string{
				"load",
				"balancer",
				"lb",
				"member",
				"v1",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 Networking client. A Networking client is needed to create an LB member. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new LB member.`,
				},
				resource.Attribute{
					Name:        "pool_id",
					Description: `(Required) The ID of the LB pool. Changing this creates a new member.`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `(Required) The IP address of the member. Changing this creates a new member.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Required) An integer representing the port on which the member is hosted. Changing this creates a new member.`,
				},
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `(Optional) The administrative state of the member. Acceptable values are 'true' and 'false'. Changing this value updates the state of the existing member.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Optional) The owner of the member. Required if admin wants to create a member for another tenant. Changing this creates a new member. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "pool_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `The load balancing weight of the member. This is currently unable to be set through Terraform. ## Import Load Balancer Members can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import openstack_lb_member_v1.member_1 a7498676-4fe4-4243-a864-2eaaf18c73df ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "pool_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `The load balancing weight of the member. This is currently unable to be set through Terraform. ## Import Load Balancer Members can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import openstack_lb_member_v1.member_1 a7498676-4fe4-4243-a864-2eaaf18c73df ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "openstack_lb_member_v2",
			Category:         "Load Balancer Resources",
			ShortDescription: `Manages a V2 member resource within OpenStack.`,
			Description:      ``,
			Keywords: []string{
				"load",
				"balancer",
				"lb",
				"member",
				"v2",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 Networking client. A Networking client is needed to create an . If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new member.`,
				},
				resource.Attribute{
					Name:        "pool_id",
					Description: `(Required) The id of the pool that this member will be assigned to.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Optional) The subnet in which to access the member`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Human-readable name for the member.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Optional) Required for admins. The UUID of the tenant who owns the member. Only administrative users can specify a tenant UUID other than their own. Changing this creates a new member.`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `(Required) The IP address of the member to receive traffic from the load balancer. Changing this creates a new member.`,
				},
				resource.Attribute{
					Name:        "protocol_port",
					Description: `(Required) The port on which to listen for client traffic. Changing this creates a new member.`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `(Optional) A positive integer value that indicates the relative portion of traffic that this member should receive from the pool. For example, a member with a weight of 10 receives five times as much traffic as a member with a weight of 2.`,
				},
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `(Optional) The administrative state of the member. A valid value is true (UP) or false (DOWN). ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique ID for the member.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "pool_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "protocol_port",
					Description: `See Argument Reference above. ## Import Load Balancer Pool Member can be imported using the Pool ID and Member ID separated by a slash. e.g.: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import openstack_lb_member_v2.member_1 c22974d2-4c95-4bcb-9819-0afc5ed303d5/9563b79c-8460-47da-8a95-2711b746510f ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "id",
					Description: `The unique ID for the member.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "pool_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "protocol_port",
					Description: `See Argument Reference above. ## Import Load Balancer Pool Member can be imported using the Pool ID and Member ID separated by a slash. e.g.: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import openstack_lb_member_v2.member_1 c22974d2-4c95-4bcb-9819-0afc5ed303d5/9563b79c-8460-47da-8a95-2711b746510f ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "openstack_lb_monitor_v1",
			Category:         "Load Balancer Resources",
			ShortDescription: `Manages a V1 load balancer monitor resource within OpenStack.`,
			Description:      ``,
			Keywords: []string{
				"load",
				"balancer",
				"lb",
				"monitor",
				"v1",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 Networking client. A Networking client is needed to create an LB monitor. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new LB monitor.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) The type of probe, which is PING, TCP, HTTP, or HTTPS, that is sent by the monitor to verify the member state. Changing this creates a new monitor.`,
				},
				resource.Attribute{
					Name:        "delay",
					Description: `(Required) The time, in seconds, between sending probes to members. Changing this creates a new monitor.`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `(Required) Maximum number of seconds for a monitor to wait for a ping reply before it times out. The value must be less than the delay value. Changing this updates the timeout of the existing monitor.`,
				},
				resource.Attribute{
					Name:        "max_retries",
					Description: `(Required) Number of permissible ping failures before changing the member's status to INACTIVE. Must be a number between 1 and 10. Changing this updates the max_retries of the existing monitor.`,
				},
				resource.Attribute{
					Name:        "url_path",
					Description: `(Optional) Required for HTTP(S) types. URI path that will be accessed if monitor type is HTTP or HTTPS. Changing this updates the url_path of the existing monitor.`,
				},
				resource.Attribute{
					Name:        "http_method",
					Description: `(Optional) Required for HTTP(S) types. The HTTP method used for requests by the monitor. If this attribute is not specified, it defaults to "GET". Changing this updates the http_method of the existing monitor.`,
				},
				resource.Attribute{
					Name:        "expected_codes",
					Description: `(Optional) Required for HTTP(S) types. Expected HTTP codes for a passing HTTP(S) monitor. You can either specify a single status like "200", or a range like "200-202". Changing this updates the expected_codes of the existing monitor.`,
				},
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `(Optional) The administrative state of the monitor. Acceptable values are "true" and "false". Changing this value updates the state of the existing monitor.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Optional) The owner of the monitor. Required if admin wants to create a monitor for another tenant. Changing this creates a new monitor. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "delay",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "max_retries",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "url_path",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "http_method",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "expected_codes",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above. ## Import Load Balancer Members can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import openstack_lb_monitor_v1.monitor_1 119d7530-72e9-449a-aa97-124a5ef1992c ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "delay",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "max_retries",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "url_path",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "http_method",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "expected_codes",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above. ## Import Load Balancer Members can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import openstack_lb_monitor_v1.monitor_1 119d7530-72e9-449a-aa97-124a5ef1992c ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "openstack_lb_monitor_v2",
			Category:         "Load Balancer Resources",
			ShortDescription: `Manages a V2 monitor resource within OpenStack.`,
			Description:      ``,
			Keywords: []string{
				"load",
				"balancer",
				"lb",
				"monitor",
				"v2",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 Networking client. A Networking client is needed to create an . If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new monitor.`,
				},
				resource.Attribute{
					Name:        "pool_id",
					Description: `(Required) The id of the pool that this monitor will be assigned to.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The Name of the Monitor.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Optional) Required for admins. The UUID of the tenant who owns the monitor. Only administrative users can specify a tenant UUID other than their own. Changing this creates a new monitor.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) The type of probe, which is PING, TCP, HTTP, or HTTPS, that is sent by the load balancer to verify the member state. Changing this creates a new monitor.`,
				},
				resource.Attribute{
					Name:        "delay",
					Description: `(Required) The time, in seconds, between sending probes to members.`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `(Required) Maximum number of seconds for a monitor to wait for a ping reply before it times out. The value must be less than the delay value.`,
				},
				resource.Attribute{
					Name:        "max_retries",
					Description: `(Required) Number of permissible ping failures before changing the member's status to INACTIVE. Must be a number between 1 and 10..`,
				},
				resource.Attribute{
					Name:        "url_path",
					Description: `(Optional) Required for HTTP(S) types. URI path that will be accessed if monitor type is HTTP or HTTPS.`,
				},
				resource.Attribute{
					Name:        "expected_codes",
					Description: `(Optional) Required for HTTP(S) types. Expected HTTP codes for a passing HTTP(S) monitor. You can either specify a single status like "200", or a range like "200-202".`,
				},
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `(Optional) The administrative state of the monitor. A valid value is true (UP) or false (DOWN). ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique ID for the monitor.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "delay",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "max_retries",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "url_path",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "http_method",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "expected_codes",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `See Argument Reference above. ## Import Load Balancer Pool Monitor can be imported using the Monitor ID, e.g.: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import openstack_lb_monitor_v2.monitor_1 47c26fc3-2403-427a-8c79-1589bd0533c2 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "id",
					Description: `The unique ID for the monitor.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "delay",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "max_retries",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "url_path",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "http_method",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "expected_codes",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `See Argument Reference above. ## Import Load Balancer Pool Monitor can be imported using the Monitor ID, e.g.: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import openstack_lb_monitor_v2.monitor_1 47c26fc3-2403-427a-8c79-1589bd0533c2 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "openstack_lb_pool_v1",
			Category:         "Load Balancer Resources",
			ShortDescription: `Manages a V1 load balancer pool resource within OpenStack.`,
			Description:      ``,
			Keywords: []string{
				"load",
				"balancer",
				"lb",
				"pool",
				"v1",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 Networking client. A Networking client is needed to create an LB pool. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new LB pool.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the pool. Changing this updates the name of the existing pool.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Required) The protocol used by the pool members, you can use either 'TCP, 'HTTP', or 'HTTPS'. Changing this creates a new pool.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Required) The network on which the members of the pool will be located. Only members that are on this network can be added to the pool. Changing this creates a new pool.`,
				},
				resource.Attribute{
					Name:        "lb_method",
					Description: `(Required) The algorithm used to distribute load between the members of the pool. The current specification supports 'ROUND_ROBIN' and 'LEAST_CONNECTIONS' as valid values for this attribute.`,
				},
				resource.Attribute{
					Name:        "lb_provider",
					Description: `(Optional) The backend load balancing provider. For example: ` + "`" + `haproxy` + "`" + `, ` + "`" + `F5` + "`" + `, etc.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Optional) The owner of the pool. Required if admin wants to create a pool member for another tenant. Changing this creates a new pool.`,
				},
				resource.Attribute{
					Name:        "monitor_ids",
					Description: `(Optional) A list of IDs of monitors to associate with the pool.`,
				},
				resource.Attribute{
					Name:        "member",
					Description: `(Optional) An existing node to add to the pool. Changing this updates the members of the pool. The member object structure is documented below. Please note that the ` + "`" + `member` + "`" + ` block is deprecated in favor of the ` + "`" + `openstack_lb_member_v1` + "`" + ` resource. The ` + "`" + `member` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `(Required) The IP address of the member. Changing this creates a new member.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Required) An integer representing the port on which the member is hosted. Changing this creates a new member.`,
				},
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `(Required) The administrative state of the member. Acceptable values are 'true' and 'false'. Changing this value updates the state of the existing member.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Optional) The owner of the member. Required if admin wants to create a pool member for another tenant. Changing this creates a new member. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "lb_method",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "lb_provider",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "monitor_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "member",
					Description: `See Argument Reference above. ## Notes The ` + "`" + `member` + "`" + ` block is deprecated in favor of the ` + "`" + `openstack_lb_member_v1` + "`" + ` resource. ## Import Load Balancer Pools can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import openstack_lb_pool_v1.pool_1 b255e6ba-02ad-43e6-8951-3428ca26b713 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "lb_method",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "lb_provider",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "monitor_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "member",
					Description: `See Argument Reference above. ## Notes The ` + "`" + `member` + "`" + ` block is deprecated in favor of the ` + "`" + `openstack_lb_member_v1` + "`" + ` resource. ## Import Load Balancer Pools can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import openstack_lb_pool_v1.pool_1 b255e6ba-02ad-43e6-8951-3428ca26b713 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "openstack_lb_pool_v2",
			Category:         "Load Balancer Resources",
			ShortDescription: `Manages a V2 pool resource within OpenStack.`,
			Description:      ``,
			Keywords: []string{
				"load",
				"balancer",
				"lb",
				"pool",
				"v2",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 Networking client. A Networking client is needed to create an . If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new pool.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Optional) Required for admins. The UUID of the tenant who owns the pool. Only administrative users can specify a tenant UUID other than their own. Changing this creates a new pool.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Human-readable name for the pool.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Human-readable description for the pool.`,
				},
				resource.Attribute{
					Name:        "loadbalancer_id",
					Description: `(Optional) The load balancer on which to provision this pool. Changing this creates a new pool. Note: One of LoadbalancerID or ListenerID must be provided.`,
				},
				resource.Attribute{
					Name:        "listener_id",
					Description: `(Optional) The Listener on which the members of the pool will be associated with. Changing this creates a new pool. Note: One of LoadbalancerID or ListenerID must be provided.`,
				},
				resource.Attribute{
					Name:        "lb_method",
					Description: `(Required) The load balancing algorithm to distribute traffic to the pool's members. Must be one of ROUND_ROBIN, LEAST_CONNECTIONS, or SOURCE_IP.`,
				},
				resource.Attribute{
					Name:        "persistence",
					Description: `Omit this field to prevent session persistence. Indicates whether connections in the same session will be processed by the same Pool member or not. Changing this creates a new pool.`,
				},
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `(Optional) The administrative state of the pool. A valid value is true (UP) or false (DOWN). The ` + "`" + `persistence` + "`" + ` argument supports:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) The type of persistence mode. The current specification supports SOURCE_IP, HTTP_COOKIE, and APP_COOKIE.`,
				},
				resource.Attribute{
					Name:        "cookie_name",
					Description: `(Optional) The name of the cookie if persistence mode is set appropriately. Required if ` + "`" + `type = APP_COOKIE` + "`" + `. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique ID for the pool.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "lb_method",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "persistence",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `See Argument Reference above. ## Import Load Balancer Pool can be imported using the Pool ID, e.g.: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import openstack_lb_pool_v2.pool_1 60ad9ee4-249a-4d60-a45b-aa60e046c513 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "id",
					Description: `The unique ID for the pool.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "lb_method",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "persistence",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `See Argument Reference above. ## Import Load Balancer Pool can be imported using the Pool ID, e.g.: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import openstack_lb_pool_v2.pool_1 60ad9ee4-249a-4d60-a45b-aa60e046c513 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "openstack_lb_vip_v1",
			Category:         "Load Balancer Resources",
			ShortDescription: `Manages a V1 load balancer vip resource within OpenStack.`,
			Description:      ``,
			Keywords: []string{
				"load",
				"balancer",
				"lb",
				"vip",
				"v1",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 Networking client. A Networking client is needed to create a VIP. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new VIP.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the vip. Changing this updates the name of the existing vip.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Required) The network on which to allocate the vip's address. A tenant can only create vips on networks authorized by policy (e.g. networks that belong to them or networks that are shared). Changing this creates a new vip.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Required) The protocol - can be either 'TCP, 'HTTP', or HTTPS'. Changing this creates a new vip.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Required) The port on which to listen for client traffic. Changing this creates a new vip.`,
				},
				resource.Attribute{
					Name:        "pool_id",
					Description: `(Required) The ID of the pool with which the vip is associated. Changing this updates the pool_id of the existing vip.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Optional) The owner of the vip. Required if admin wants to create a vip member for another tenant. Changing this creates a new vip.`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `(Optional) The IP address of the vip. Changing this creates a new vip.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Human-readable description for the vip. Changing this updates the description of the existing vip.`,
				},
				resource.Attribute{
					Name:        "persistence",
					Description: `(Optional) Omit this field to prevent session persistence. The persistence object structure is documented below. Changing this updates the persistence of the existing vip.`,
				},
				resource.Attribute{
					Name:        "conn_limit",
					Description: `(Optional) The maximum number of connections allowed for the vip. Default is -1, meaning no limit. Changing this updates the conn_limit of the existing vip.`,
				},
				resource.Attribute{
					Name:        "floating_ip",
					Description: `(Optional) A`,
				},
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `(Optional) The administrative state of the vip. Acceptable values are "true" and "false". Changing this value updates the state of the existing vip. The ` + "`" + `persistence` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) The type of persistence mode. Valid values are "SOURCE_IP", "HTTP_COOKIE", or "APP_COOKIE".`,
				},
				resource.Attribute{
					Name:        "cookie_name",
					Description: `(Optional) The name of the cookie if persistence mode is set appropriately. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "pool_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "persistence",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "conn_limit",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "floating_ip",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "port_id",
					Description: `Port UUID for this VIP at associated floating IP (if any). ## Import Load Balancer VIPs can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import openstack_lb_vip_v1.vip_1 50e16b26-89c1-475e-a492-76167182511e ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "pool_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "persistence",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "conn_limit",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "floating_ip",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "port_id",
					Description: `Port UUID for this VIP at associated floating IP (if any). ## Import Load Balancer VIPs can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import openstack_lb_vip_v1.vip_1 50e16b26-89c1-475e-a492-76167182511e ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "openstack_networking_addressscope_v2",
			Category:         "Networking Resources",
			ShortDescription: `Manages a V2 Neutron addressscope resource within OpenStack.`,
			Description:      ``,
			Keywords: []string{
				"networking",
				"addressscope",
				"v2",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 Networking client. A Networking client is needed to create a Neutron address-scope. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new address-scope.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the address-scope. Changing this updates the name of the existing address-scope.`,
				},
				resource.Attribute{
					Name:        "ip_version",
					Description: `(Optional) IP version, either 4 (default) or 6. Changing this creates a new address-scope.`,
				},
				resource.Attribute{
					Name:        "shared",
					Description: `(Optional) Indicates whether this address-scope is shared across all projects. Changing this updates the shared status of the existing address-scope.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Optional) The owner of the address-scope. Required if admin wants to create a address-scope for another project. Changing this creates a new address-scope. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "ip_version",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "shared",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `See Argument Reference above. ## Import Address-scopes can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import openstack_networking_addressscope_v2.addressscope_1 9cc35860-522a-4d35-974d-51d4b011801e ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "ip_version",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "shared",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `See Argument Reference above. ## Import Address-scopes can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import openstack_networking_addressscope_v2.addressscope_1 9cc35860-522a-4d35-974d-51d4b011801e ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "openstack_networking_floatingip_associate_v2",
			Category:         "Networking Resources",
			ShortDescription: `Associates a Floating IP to a Port`,
			Description:      ``,
			Keywords: []string{
				"networking",
				"floatingip",
				"associate",
				"v2",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 Networking client. A Networking client is needed to create a floating IP that can be used with another networking resource, such as a load balancer. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new floating IP (which may or may not have a different address).`,
				},
				resource.Attribute{
					Name:        "floating_ip",
					Description: `(Required) IP Address of an existing floating IP.`,
				},
				resource.Attribute{
					Name:        "port_id",
					Description: `(Required) ID of an existing port with at least one IP address to associate with this floating IP. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "floating_ip",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "port_id",
					Description: `See Argument Reference above. ## Import Floating IP associations can be imported using the ` + "`" + `id` + "`" + ` of the floating IP, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import openstack_networking_floatingip_associate_v2.fip 2c7f39f3-702b-48d1-940c-b50384177ee1 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "floating_ip",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "port_id",
					Description: `See Argument Reference above. ## Import Floating IP associations can be imported using the ` + "`" + `id` + "`" + ` of the floating IP, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import openstack_networking_floatingip_associate_v2.fip 2c7f39f3-702b-48d1-940c-b50384177ee1 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "openstack_networking_floatingip_v2",
			Category:         "Networking Resources",
			ShortDescription: `Manages a V2 floating IP resource within OpenStack Neutron (networking).`,
			Description:      ``,
			Keywords: []string{
				"networking",
				"floatingip",
				"v2",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 Networking client. A Networking client is needed to create a floating IP that can be used with another networking resource, such as a load balancer. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new floating IP (which may or may not have a different address).`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Human-readable description for the floating IP.`,
				},
				resource.Attribute{
					Name:        "pool",
					Description: `(Required) The name of the pool from which to obtain the floating IP. Changing this creates a new floating IP.`,
				},
				resource.Attribute{
					Name:        "port_id",
					Description: `(Optional) ID of an existing port with at least one IP address to associate with this floating IP.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Optional) The target tenant ID in which to allocate the floating IP, if you specify this together with a port_id, make sure the target port belongs to the same tenant. Changing this creates a new floating IP (which may or may not have a different address)`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `(Optional) The actual/specific floating IP to obtain. By default, non-admin users are not able to specify a floating IP, so you must either be an admin user or have had a custom policy or role applied to your OpenStack user or project.`,
				},
				resource.Attribute{
					Name:        "fixed_ip",
					Description: `Fixed IP of the port to associate with this floating IP. Required if the port has multiple fixed IPs.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Optional) The subnet ID of the floating IP pool. Specify this if the floating IP network has multiple subnets.`,
				},
				resource.Attribute{
					Name:        "value_specs",
					Description: `(Optional) Map of additional options.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A set of string tags for the floating IP.`,
				},
				resource.Attribute{
					Name:        "dns_name",
					Description: `(Optional) The floating IP DNS name. Available, when Neutron DNS extension is enabled. The data in this attribute will be published in an external DNS service when Neutron is configured to integrate with such a service. Changing this creates a new floating IP.`,
				},
				resource.Attribute{
					Name:        "dns_domain",
					Description: `(Optional) The floating IP DNS domain. Available, when Neutron DNS extension is enabled. The data in this attribute will be published in an external DNS service when Neutron is configured to integrate with such a service. Changing this creates a new floating IP. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "pool",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `The actual floating IP address itself.`,
				},
				resource.Attribute{
					Name:        "port_id",
					Description: `ID of associated port.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `the ID of the tenant in which to create the floating IP.`,
				},
				resource.Attribute{
					Name:        "fixed_ip",
					Description: `The fixed IP which the floating IP maps to.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "all_tags",
					Description: `The collection of tags assigned on the floating IP, which have been explicitly and implicitly added.`,
				},
				resource.Attribute{
					Name:        "dns_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "dns_domain",
					Description: `See Argument Reference above. ## Import Floating IPs can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import openstack_networking_floatingip_v2.floatip_1 2c7f39f3-702b-48d1-940c-b50384177ee1 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "pool",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `The actual floating IP address itself.`,
				},
				resource.Attribute{
					Name:        "port_id",
					Description: `ID of associated port.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `the ID of the tenant in which to create the floating IP.`,
				},
				resource.Attribute{
					Name:        "fixed_ip",
					Description: `The fixed IP which the floating IP maps to.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "all_tags",
					Description: `The collection of tags assigned on the floating IP, which have been explicitly and implicitly added.`,
				},
				resource.Attribute{
					Name:        "dns_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "dns_domain",
					Description: `See Argument Reference above. ## Import Floating IPs can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import openstack_networking_floatingip_v2.floatip_1 2c7f39f3-702b-48d1-940c-b50384177ee1 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "openstack_networking_network_v2",
			Category:         "Networking Resources",
			ShortDescription: `Manages a V2 Neutron network resource within OpenStack.`,
			Description:      ``,
			Keywords: []string{
				"networking",
				"network",
				"v2",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 Networking client. A Networking client is needed to create a Neutron network. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new network.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the network. Changing this updates the name of the existing network.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Human-readable description of the network. Changing this updates the name of the existing network.`,
				},
				resource.Attribute{
					Name:        "shared",
					Description: `(Optional) Specifies whether the network resource can be accessed by any tenant or not. Changing this updates the sharing capabilities of the existing network.`,
				},
				resource.Attribute{
					Name:        "external",
					Description: `(Optional) Specifies whether the network resource has the external routing facility. Valid values are true and false. Defaults to false. Changing this updates the external attribute of the existing network.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Optional) The owner of the network. Required if admin wants to create a network for another tenant. Changing this creates a new network.`,
				},
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `(Optional) The administrative state of the network. Acceptable values are "true" and "false". Changing this value updates the state of the existing network.`,
				},
				resource.Attribute{
					Name:        "segments",
					Description: `(Optional) An array of one or more provider segment objects.`,
				},
				resource.Attribute{
					Name:        "value_specs",
					Description: `(Optional) Map of additional options.`,
				},
				resource.Attribute{
					Name:        "availability_zone_hints",
					Description: `(Optional) An availability zone is used to make network resources highly available. Used for resources with high availability so that they are scheduled on different availability zones. Changing this creates a new network.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A set of string tags for the network.`,
				},
				resource.Attribute{
					Name:        "transparent_vlan",
					Description: `(Optional) Specifies whether the network resource has the VLAN transparent attribute set. Valid values are true and false. Defaults to false. Changing this updates the ` + "`" + `transparent_vlan` + "`" + ` attribute of the existing network.`,
				},
				resource.Attribute{
					Name:        "port_security_enabled",
					Description: `(Optional) Whether to explicitly enable or disable port security on the network. Port Security is usually enabled by default, so omitting this argument will usually result in a value of "true". Setting this explicitly to ` + "`" + `false` + "`" + ` will disable port security. Valid values are ` + "`" + `true` + "`" + ` and ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "mtu",
					Description: `(Optional) The network MTU. Available for read-only, when Neutron ` + "`" + `net-mtu` + "`" + ` extension is enabled. Available for the modification, when Neutron ` + "`" + `net-mtu-writable` + "`" + ` extension is enabled.`,
				},
				resource.Attribute{
					Name:        "dns_domain",
					Description: `(Optional) The network DNS domain. Available, when Neutron DNS extension is enabled. The ` + "`" + `dns_domain` + "`" + ` of a network in conjunction with the ` + "`" + `dns_name` + "`" + ` attribute of its ports will be published in an external DNS service when Neutron is configured to integrate with such a service.`,
				},
				resource.Attribute{
					Name:        "qos_policy_id",
					Description: `(Optional) Reference to the associated QoS policy. The ` + "`" + `segments` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "physical_network",
					Description: `The physical network where this network is implemented.`,
				},
				resource.Attribute{
					Name:        "segmentation_id",
					Description: `An isolated segment on the physical network.`,
				},
				resource.Attribute{
					Name:        "network_type",
					Description: `The type of physical network. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "shared",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "external",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "availability_zone_hints",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "all_tags",
					Description: `The collection of tags assigned on the network, which have been explicitly and implicitly added.`,
				},
				resource.Attribute{
					Name:        "transparent_vlan",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "port_security_enabled",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "mtu",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "dns_domain",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "qos_policy_id",
					Description: `See Argument Reference above. ## Import Networks can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import openstack_networking_network_v2.network_1 d90ce693-5ccf-4136-a0ed-152ce412b6b9 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "shared",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "external",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "availability_zone_hints",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "all_tags",
					Description: `The collection of tags assigned on the network, which have been explicitly and implicitly added.`,
				},
				resource.Attribute{
					Name:        "transparent_vlan",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "port_security_enabled",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "mtu",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "dns_domain",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "qos_policy_id",
					Description: `See Argument Reference above. ## Import Networks can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import openstack_networking_network_v2.network_1 d90ce693-5ccf-4136-a0ed-152ce412b6b9 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "openstack_networking_port_secgroup_associate_v2",
			Category:         "Networking Resources",
			ShortDescription: `Manages a V2 port's security groups within OpenStack.`,
			Description:      ``,
			Keywords: []string{
				"networking",
				"port",
				"secgroup",
				"associate",
				"v2",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 networking client. A networking client is needed to manage a port. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new resource.`,
				},
				resource.Attribute{
					Name:        "port_id",
					Description: `(Required) An UUID of the port to apply security groups to.`,
				},
				resource.Attribute{
					Name:        "security_group_ids",
					Description: `(Required) A list of security group IDs to apply to the port. The security groups must be specified by ID and not name (as opposed to how they are configured with the Compute Instance).`,
				},
				resource.Attribute{
					Name:        "enforce",
					Description: `(Optional) Whether to replace or append the list of security groups, specified in the ` + "`" + `security_group_ids` + "`" + `. Defaults to ` + "`" + `false` + "`" + `. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "port_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "security_group_ids",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "all_security_group_ids",
					Description: `The collection of Security Group IDs on the port which have been explicitly and implicitly added.`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "port_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "security_group_ids",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "all_security_group_ids",
					Description: `The collection of Security Group IDs on the port which have been explicitly and implicitly added.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "openstack_networking_port_v2",
			Category:         "Networking Resources",
			ShortDescription: `Manages a V2 port resource within OpenStack.`,
			Description:      ``,
			Keywords: []string{
				"networking",
				"port",
				"v2",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 networking client. A networking client is needed to create a port. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new port.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) A unique name for the port. Changing this updates the ` + "`" + `name` + "`" + ` of an existing port.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Human-readable description of the floating IP. Changing this updates the ` + "`" + `description` + "`" + ` of an existing port.`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `(Required) The ID of the network to attach the port to. Changing this creates a new port.`,
				},
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `(Optional) Administrative up/down status for the port (must be "true" or "false" if provided). Changing this updates the ` + "`" + `admin_state_up` + "`" + ` of an existing port.`,
				},
				resource.Attribute{
					Name:        "mac_address",
					Description: `(Optional) Specify a specific MAC address for the port. Changing this creates a new port.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Optional) The owner of the Port. Required if admin wants to create a port for another tenant. Changing this creates a new port.`,
				},
				resource.Attribute{
					Name:        "device_owner",
					Description: `(Optional) The device owner of the Port. Changing this creates a new port.`,
				},
				resource.Attribute{
					Name:        "security_group_ids",
					Description: `(Optional - Conflicts with ` + "`" + `no_security_groups` + "`" + `) A list of security group IDs to apply to the port. The security groups must be specified by ID and not name (as opposed to how they are configured with the Compute Instance).`,
				},
				resource.Attribute{
					Name:        "no_security_groups",
					Description: `(Optional - Conflicts with ` + "`" + `security_group_ids` + "`" + `) If set to ` + "`" + `true` + "`" + `, then no security groups are applied to the port. If set to ` + "`" + `false` + "`" + ` and no ` + "`" + `security_group_ids` + "`" + ` are specified, then the Port will yield to the default behavior of the Networking service, which is to usually apply the "default" security group.`,
				},
				resource.Attribute{
					Name:        "device_id",
					Description: `(Optional) The ID of the device attached to the port. Changing this creates a new port.`,
				},
				resource.Attribute{
					Name:        "fixed_ip",
					Description: `(Optional - Conflicts with ` + "`" + `no_fixed_ip` + "`" + `) An array of desired IPs for this port. The structure is described below.`,
				},
				resource.Attribute{
					Name:        "no_fixed_ip",
					Description: `(Optional - Conflicts with ` + "`" + `fixed_ip` + "`" + `) Create a port with no fixed IP address. This will also remove any fixed IPs previously set on a port. ` + "`" + `true` + "`" + ` is the only valid value for this argument.`,
				},
				resource.Attribute{
					Name:        "allowed_address_pairs",
					Description: `(Optional) An IP/MAC Address pair of additional IP addresses that can be active on this port. The structure is described below.`,
				},
				resource.Attribute{
					Name:        "extra_dhcp_option",
					Description: `(Optional) An extra DHCP option that needs to be configured on the port. The structure is described below. Can be specified multiple times.`,
				},
				resource.Attribute{
					Name:        "port_security_enabled",
					Description: `(Optional) Whether to explicitly enable or disable port security on the port. Port Security is usually enabled by default, so omitting argument will usually result in a value of "true". Setting this explicitly to ` + "`" + `false` + "`" + ` will disable port security. In order to disable port security, the port must not have any security groups. Valid values are ` + "`" + `true` + "`" + ` and ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "value_specs",
					Description: `(Optional) Map of additional options.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A set of string tags for the port.`,
				},
				resource.Attribute{
					Name:        "binding",
					Description: `(Optional) The port binding allows to specify binding information for the port. The structure is described below.`,
				},
				resource.Attribute{
					Name:        "dns_name",
					Description: `(Optional) The port DNS name. Available, when Neutron DNS extension is enabled.`,
				},
				resource.Attribute{
					Name:        "qos_policy_id",
					Description: `(Optional) Reference to the associated QoS policy. The ` + "`" + `fixed_ip` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Required) Subnet in which to allocate IP address for this port.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `(Optional) IP address desired in the subnet for this port. If you don't specify ` + "`" + `ip_address` + "`" + `, an available IP address from the specified subnet will be allocated to this port. This field will not be populated if it is left blank or omitted. To retrieve the assigned IP address, use the ` + "`" + `all_fixed_ips` + "`" + ` attribute. The ` + "`" + `allowed_address_pairs` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `(Required) The additional IP address.`,
				},
				resource.Attribute{
					Name:        "mac_address",
					Description: `(Optional) The additional MAC address. The ` + "`" + `extra_dhcp_option` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the DHCP option.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) Value of the DHCP option.`,
				},
				resource.Attribute{
					Name:        "ip_version",
					Description: `(Optional) IP protocol version. Defaults to 4. The ` + "`" + `binding` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "host_id",
					Description: `(Optional) The ID of the host to allocate port on.`,
				},
				resource.Attribute{
					Name:        "profile",
					Description: `(Optional) Custom data to be passed as ` + "`" + `binding:profile` + "`" + `. Data must be passed as JSON.`,
				},
				resource.Attribute{
					Name:        "vnic_type",
					Description: `(Optional) VNIC type for the port. Can either be ` + "`" + `direct` + "`" + `, ` + "`" + `direct-physical` + "`" + `, ` + "`" + `macvtap` + "`" + `, ` + "`" + `normal` + "`" + `, ` + "`" + `baremetal` + "`" + ` or ` + "`" + `virtio-forwarder` + "`" + `. Default value is ` + "`" + `normal` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "vif_details",
					Description: `(Computed) A map of JSON strings containing additional details for this specific binding.`,
				},
				resource.Attribute{
					Name:        "vif_type",
					Description: `(Computed) The VNIC type of the port binding. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "mac_address",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "device_owner",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "security_group_ids",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "device_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "fixed_ip",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "all_fixed_ips",
					Description: `The collection of Fixed IP addresses on the port in the order returned by the Network v2 API.`,
				},
				resource.Attribute{
					Name:        "all_security_group_ids",
					Description: `The collection of Security Group IDs on the port which have been explicitly and implicitly added.`,
				},
				resource.Attribute{
					Name:        "extra_dhcp_option",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "all_tags",
					Description: `The collection of tags assigned on the port, which have been explicitly and implicitly added.`,
				},
				resource.Attribute{
					Name:        "binding",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "dns_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "dns_assignment",
					Description: `The list of maps representing port DNS assignments.`,
				},
				resource.Attribute{
					Name:        "qos_policy_id",
					Description: `See Argument Reference above. ## Import Ports can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import openstack_networking_port_v2.port_1 eae26a3e-1c33-4cc1-9c31-0cd729c438a1 ` + "`" + `` + "`" + `` + "`" + ` ## Notes ### Ports and Instances There are some notes to consider when connecting Instances to networks using Ports. Please see the ` + "`" + `openstack_compute_instance_v2` + "`" + ` documentation for further documentation.`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "mac_address",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "device_owner",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "security_group_ids",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "device_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "fixed_ip",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "all_fixed_ips",
					Description: `The collection of Fixed IP addresses on the port in the order returned by the Network v2 API.`,
				},
				resource.Attribute{
					Name:        "all_security_group_ids",
					Description: `The collection of Security Group IDs on the port which have been explicitly and implicitly added.`,
				},
				resource.Attribute{
					Name:        "extra_dhcp_option",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "all_tags",
					Description: `The collection of tags assigned on the port, which have been explicitly and implicitly added.`,
				},
				resource.Attribute{
					Name:        "binding",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "dns_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "dns_assignment",
					Description: `The list of maps representing port DNS assignments.`,
				},
				resource.Attribute{
					Name:        "qos_policy_id",
					Description: `See Argument Reference above. ## Import Ports can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import openstack_networking_port_v2.port_1 eae26a3e-1c33-4cc1-9c31-0cd729c438a1 ` + "`" + `` + "`" + `` + "`" + ` ## Notes ### Ports and Instances There are some notes to consider when connecting Instances to networks using Ports. Please see the ` + "`" + `openstack_compute_instance_v2` + "`" + ` documentation for further documentation.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "openstack_networking_qos_bandwidth_limit_rule_v2",
			Category:         "Networking Resources",
			ShortDescription: `Manages a V2 Neutron QoS bandwidth limit rule resource within OpenStack.`,
			Description:      ``,
			Keywords: []string{
				"networking",
				"qos",
				"bandwidth",
				"limit",
				"rule",
				"v2",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 Networking client. A Networking client is needed to create a Neutron QoS bandwidth limit rule. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new QoS bandwidth limit rule.`,
				},
				resource.Attribute{
					Name:        "qos_policy_id",
					Description: `(Required) The QoS policy reference. Changing this creates a new QoS bandwidth limit rule.`,
				},
				resource.Attribute{
					Name:        "max_kbps",
					Description: `(Required) The maximum kilobits per second of a QoS bandwidth limit rule. Changing this updates the maximum kilobits per second of the existing QoS bandwidth limit rule.`,
				},
				resource.Attribute{
					Name:        "max_burst_kbps",
					Description: `(Optional) The maximum burst size in kilobits of a QoS bandwidth limit rule. Changing this updates the maximum burst size in kilobits of the existing QoS bandwidth limit rule.`,
				},
				resource.Attribute{
					Name:        "direction",
					Description: `(Optional) The direction of traffic. Defaults to "egress". Changing this updates the direction of the existing QoS bandwidth limit rule. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "qos_policy_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "max_kbps",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "max_burst_kbps",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "direction",
					Description: `See Argument Reference above. ## Import QoS bandwidth limit rules can be imported using the ` + "`" + `qos_policy_id/bandwidth_limit_rule` + "`" + ` format, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import openstack_networking_qos_bandwidth_limit_rule_v2.bw_limit_rule_1 d6ae28ce-fcb5-4180-aa62-d260a27e09ae/46dfb556-b92f-48ce-94c5-9a9e2140de94 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "qos_policy_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "max_kbps",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "max_burst_kbps",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "direction",
					Description: `See Argument Reference above. ## Import QoS bandwidth limit rules can be imported using the ` + "`" + `qos_policy_id/bandwidth_limit_rule` + "`" + ` format, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import openstack_networking_qos_bandwidth_limit_rule_v2.bw_limit_rule_1 d6ae28ce-fcb5-4180-aa62-d260a27e09ae/46dfb556-b92f-48ce-94c5-9a9e2140de94 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "openstack_networking_qos_dscp_marking_rule_v2",
			Category:         "Networking Resources",
			ShortDescription: `Manages a V2 Neutron QoS DSCP marking rule resource within OpenStack.`,
			Description:      ``,
			Keywords: []string{
				"networking",
				"qos",
				"dscp",
				"marking",
				"rule",
				"v2",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 Networking client. A Networking client is needed to create a Neutron QoS DSCP marking rule. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new QoS DSCP marking rule.`,
				},
				resource.Attribute{
					Name:        "qos_policy_id",
					Description: `(Required) The QoS policy reference. Changing this creates a new QoS DSCP marking rule.`,
				},
				resource.Attribute{
					Name:        "dscp_mark",
					Description: `(Required) The value of DSCP mark. Changing this updates the DSCP mark value existing QoS DSCP marking rule. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "qos_policy_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "dscp_mark",
					Description: `See Argument Reference above. ## Import QoS DSCP marking rules can be imported using the ` + "`" + `qos_policy_id/dscp_marking_rule_id` + "`" + ` format, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import openstack_networking_qos_dscp_marking_rule_v2.dscp_marking_rule_1 d6ae28ce-fcb5-4180-aa62-d260a27e09ae/46dfb556-b92f-48ce-94c5-9a9e2140de94 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "qos_policy_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "dscp_mark",
					Description: `See Argument Reference above. ## Import QoS DSCP marking rules can be imported using the ` + "`" + `qos_policy_id/dscp_marking_rule_id` + "`" + ` format, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import openstack_networking_qos_dscp_marking_rule_v2.dscp_marking_rule_1 d6ae28ce-fcb5-4180-aa62-d260a27e09ae/46dfb556-b92f-48ce-94c5-9a9e2140de94 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "openstack_networking_qos_minimum_bandwidth_rule_v2",
			Category:         "Networking Resources",
			ShortDescription: `Manages a V2 Neutron QoS minimum bandwidth rule resource within OpenStack.`,
			Description:      ``,
			Keywords: []string{
				"networking",
				"qos",
				"minimum",
				"bandwidth",
				"rule",
				"v2",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 Networking client. A Networking client is needed to create a Neutron QoS minimum bandwidth rule. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new QoS minimum bandwidth rule.`,
				},
				resource.Attribute{
					Name:        "qos_policy_id",
					Description: `(Required) The QoS policy reference. Changing this creates a new QoS minimum bandwidth rule.`,
				},
				resource.Attribute{
					Name:        "min_kbps",
					Description: `(Required) The minimum kilobits per second. Changing this updates the min kbps value of the existing QoS minimum bandwidth rule.`,
				},
				resource.Attribute{
					Name:        "direction",
					Description: `(Optional) The direction of traffic. Defaults to "egress". Changing this updates the direction of the existing QoS minimum bandwidth rule. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "qos_policy_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "min_kbps",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "direction",
					Description: `See Argument Reference above. ## Import QoS minimum bandwidth rules can be imported using the ` + "`" + `qos_policy_id/minimum_bandwidth_rule_id` + "`" + ` format, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import openstack_networking_qos_minimum_bandwidth_rule_v2.minimum_bandwidth_rule_1 d6ae28ce-fcb5-4180-aa62-d260a27e09ae/46dfb556-b92f-48ce-94c5-9a9e2140de94 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "qos_policy_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "min_kbps",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "direction",
					Description: `See Argument Reference above. ## Import QoS minimum bandwidth rules can be imported using the ` + "`" + `qos_policy_id/minimum_bandwidth_rule_id` + "`" + ` format, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import openstack_networking_qos_minimum_bandwidth_rule_v2.minimum_bandwidth_rule_1 d6ae28ce-fcb5-4180-aa62-d260a27e09ae/46dfb556-b92f-48ce-94c5-9a9e2140de94 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "openstack_networking_qos_policy_v2",
			Category:         "Networking Resources",
			ShortDescription: `Manages a V2 Neutron QoS policy resource within OpenStack.`,
			Description:      ``,
			Keywords: []string{
				"networking",
				"qos",
				"policy",
				"v2",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 Networking client. A Networking client is needed to create a Neutron Qos policy. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new QoS policy.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the QoS policy. Changing this updates the name of the existing QoS policy.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Optional) The owner of the QoS policy. Required if admin wants to create a QoS policy for another project. Changing this creates a new QoS policy.`,
				},
				resource.Attribute{
					Name:        "shared",
					Description: `(Optional) Indicates whether this QoS policy is shared across all projects. Changing this updates the shared status of the existing QoS policy.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The human-readable description for the QoS policy. Changing this updates the description of the existing QoS policy.`,
				},
				resource.Attribute{
					Name:        "is_default",
					Description: `(Optional) Indicates whether the QoS policy is default QoS policy or not. Changing this updates the default status of the existing QoS policy.`,
				},
				resource.Attribute{
					Name:        "value_specs",
					Description: `(Optional) Map of additional options.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A set of string tags for the QoS policy. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `The time at which QoS policy was created.`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `The time at which QoS policy was created.`,
				},
				resource.Attribute{
					Name:        "shared",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "is_default",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "revision_number",
					Description: `The revision number of the QoS policy.`,
				},
				resource.Attribute{
					Name:        "value_specs",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "all_tags",
					Description: `The collection of tags assigned on the QoS policy, which have been explicitly and implicitly added. ## Import QoS Policies can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import openstack_networking_qos_policy_v2.qos_policy_1 d6ae28ce-fcb5-4180-aa62-d260a27e09ae ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `The time at which QoS policy was created.`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `The time at which QoS policy was created.`,
				},
				resource.Attribute{
					Name:        "shared",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "is_default",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "revision_number",
					Description: `The revision number of the QoS policy.`,
				},
				resource.Attribute{
					Name:        "value_specs",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "all_tags",
					Description: `The collection of tags assigned on the QoS policy, which have been explicitly and implicitly added. ## Import QoS Policies can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import openstack_networking_qos_policy_v2.qos_policy_1 d6ae28ce-fcb5-4180-aa62-d260a27e09ae ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "openstack_networking_router_interface_v2",
			Category:         "Networking Resources",
			ShortDescription: `Manages a V2 router interface resource within OpenStack.`,
			Description:      ``,
			Keywords: []string{
				"networking",
				"router",
				"interface",
				"v2",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 networking client. A networking client is needed to create a router. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new router interface.`,
				},
				resource.Attribute{
					Name:        "router_id",
					Description: `(Required) ID of the router this interface belongs to. Changing this creates a new router interface.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `ID of the subnet this interface connects to. Changing this creates a new router interface.`,
				},
				resource.Attribute{
					Name:        "port_id",
					Description: `ID of the port this interface connects to. Changing this creates a new router interface. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "router_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "port_id",
					Description: `See Argument Reference above. ## Import Router Interfaces can be imported using the port ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ openstack port list --router <router name or id> $ terraform import openstack_networking_router_interface_v2.int_1 <port id from above output> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "router_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "port_id",
					Description: `See Argument Reference above. ## Import Router Interfaces can be imported using the port ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ openstack port list --router <router name or id> $ terraform import openstack_networking_router_interface_v2.int_1 <port id from above output> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "openstack_networking_router_route_v2",
			Category:         "Networking Resources",
			ShortDescription: `Creates a routing entry on a OpenStack V2 router.`,
			Description:      ``,
			Keywords: []string{
				"networking",
				"router",
				"route",
				"v2",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 networking client. A networking client is needed to configure a routing entry on a router. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new routing entry.`,
				},
				resource.Attribute{
					Name:        "router_id",
					Description: `(Required) ID of the router this routing entry belongs to. Changing this creates a new routing entry.`,
				},
				resource.Attribute{
					Name:        "destination_cidr",
					Description: `(Required) CIDR block to match on the packet’s destination IP. Changing this creates a new routing entry.`,
				},
				resource.Attribute{
					Name:        "next_hop",
					Description: `(Required) IP address of the next hop gateway. Changing this creates a new routing entry. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "router_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "destination_cidr",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "next_hop",
					Description: `See Argument Reference above. ## Notes The ` + "`" + `next_hop` + "`" + ` IP address must be directly reachable from the router at the ` + "`" + `` + "`" + `openstack_networking_router_route_v2` + "`" + `` + "`" + ` resource creation time. You can ensure that by explicitly specifying a dependency on the ` + "`" + `` + "`" + `openstack_networking_router_interface_v2` + "`" + `` + "`" + ` resource that connects the next hop to the router, as in the example above. ## Import Routing entries can be imported using a combined ID using the following format: ` + "`" + `` + "`" + `<router_id>-route-<destination_cidr>-<next_hop>` + "`" + `` + "`" + ` ` + "`" + `` + "`" + `` + "`" + ` $ terraform import openstack_networking_router_route_v2.router_route_1 686fe248-386c-4f70-9f6c-281607dad079-route-10.0.1.0/24-192.168.199.25 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "router_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "destination_cidr",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "next_hop",
					Description: `See Argument Reference above. ## Notes The ` + "`" + `next_hop` + "`" + ` IP address must be directly reachable from the router at the ` + "`" + `` + "`" + `openstack_networking_router_route_v2` + "`" + `` + "`" + ` resource creation time. You can ensure that by explicitly specifying a dependency on the ` + "`" + `` + "`" + `openstack_networking_router_interface_v2` + "`" + `` + "`" + ` resource that connects the next hop to the router, as in the example above. ## Import Routing entries can be imported using a combined ID using the following format: ` + "`" + `` + "`" + `<router_id>-route-<destination_cidr>-<next_hop>` + "`" + `` + "`" + ` ` + "`" + `` + "`" + `` + "`" + ` $ terraform import openstack_networking_router_route_v2.router_route_1 686fe248-386c-4f70-9f6c-281607dad079-route-10.0.1.0/24-192.168.199.25 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "openstack_networking_router_v2",
			Category:         "Networking Resources",
			ShortDescription: `Manages a V2 router resource within OpenStack.`,
			Description:      ``,
			Keywords: []string{
				"networking",
				"router",
				"v2",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 networking client. A networking client is needed to create a router. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new router.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) A unique name for the router. Changing this updates the ` + "`" + `name` + "`" + ` of an existing router.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Human-readable description for the router.`,
				},
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `(Optional) Administrative up/down status for the router (must be "true" or "false" if provided). Changing this updates the ` + "`" + `admin_state_up` + "`" + ` of an existing router.`,
				},
				resource.Attribute{
					Name:        "distributed",
					Description: `(Optional) Indicates whether or not to create a distributed router. The default policy setting in Neutron restricts usage of this property to administrative users only.`,
				},
				resource.Attribute{
					Name:        "external_gateway",
					Description: `(`,
				},
				resource.Attribute{
					Name:        "external_network_id",
					Description: `(Optional) The network UUID of an external gateway for the router. A router with an external gateway is required if any compute instances or load balancers will be using floating IPs. Changing this updates the external gateway of the router.`,
				},
				resource.Attribute{
					Name:        "enable_snat",
					Description: `(Optional) Enable Source NAT for the router. Valid values are "true" or "false". An ` + "`" + `external_network_id` + "`" + ` has to be set in order to set this property. Changing this updates the ` + "`" + `enable_snat` + "`" + ` of the router.`,
				},
				resource.Attribute{
					Name:        "external_fixed_ip",
					Description: `(Optional) An external fixed IP for the router. This can be repeated. The structure is described below. An ` + "`" + `external_network_id` + "`" + ` has to be set in order to set this property. Changing this updates the external fixed IPs of the router.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Optional) The owner of the floating IP. Required if admin wants to create a router for another tenant. Changing this creates a new router.`,
				},
				resource.Attribute{
					Name:        "value_specs",
					Description: `(Optional) Map of additional driver-specific options.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A set of string tags for the router.`,
				},
				resource.Attribute{
					Name:        "vendor_options",
					Description: `(Optional) Map of additional vendor-specific options. Supported options are described below.`,
				},
				resource.Attribute{
					Name:        "availability_zone_hints",
					Description: `(Optional) An availability zone is used to make network resources highly available. Used for resources with high availability so that they are scheduled on different availability zones. Changing this creates a new router. The ` + "`" + `external_fixed_ip` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Optional) Subnet in which the fixed IP belongs to.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `(Optional) The IP address to set on the router. The ` + "`" + `vendor_options` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "set_router_gateway_after_create",
					Description: `(Optional) Boolean to control whether the Router gateway is assigned during creation or updated after creation. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the router.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "external_gateway",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "external_network_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "enable_snat",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "external_fixed_ip",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "value_specs",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "availability_zone_hints",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "all_tags",
					Description: `The collection of tags assigned on the router, which have been explicitly and implicitly added. ## Import Routers can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import openstack_networking_router_v2.router_1 014395cd-89fc-4c9b-96b7-13d1ee79dad2 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the router.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "external_gateway",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "external_network_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "enable_snat",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "external_fixed_ip",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "value_specs",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "availability_zone_hints",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "all_tags",
					Description: `The collection of tags assigned on the router, which have been explicitly and implicitly added. ## Import Routers can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import openstack_networking_router_v2.router_1 014395cd-89fc-4c9b-96b7-13d1ee79dad2 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "openstack_networking_secgroup_rule_v2",
			Category:         "Networking Resources",
			ShortDescription: `Manages a V2 Neutron security group rule resource within OpenStack.`,
			Description:      ``,
			Keywords: []string{
				"networking",
				"secgroup",
				"rule",
				"v2",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 networking client. A networking client is needed to create a port. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new security group rule.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A description of the rule. Changing this creates a new security group rule.`,
				},
				resource.Attribute{
					Name:        "direction",
					Description: `(Required) The direction of the rule, valid values are __ingress__ or __egress__. Changing this creates a new security group rule.`,
				},
				resource.Attribute{
					Name:        "ethertype",
					Description: `(Required) The layer 3 protocol type, valid values are __IPv4__ or __IPv6__. Changing this creates a new security group rule.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Optional) The layer 4 protocol type, valid values are following. Changing this creates a new security group rule. This is required if you want to specify a port range.`,
				},
				resource.Attribute{
					Name:        "port_range_min",
					Description: `(Optional) The lower part of the allowed port range, valid integer value needs to be between 1 and 65535. Changing this creates a new security group rule.`,
				},
				resource.Attribute{
					Name:        "port_range_max",
					Description: `(Optional) The higher part of the allowed port range, valid integer value needs to be between 1 and 65535. Changing this creates a new security group rule.`,
				},
				resource.Attribute{
					Name:        "remote_ip_prefix",
					Description: `(Optional) The remote CIDR, the value needs to be a valid CIDR (i.e. 192.168.0.0/16). Changing this creates a new security group rule.`,
				},
				resource.Attribute{
					Name:        "remote_group_id",
					Description: `(Optional) The remote group id, the value needs to be an Openstack ID of a security group in the same tenant. Changing this creates a new security group rule.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `(Required) The security group id the rule should belong to, the value needs to be an Openstack ID of a security group in the same tenant. Changing this creates a new security group rule.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Optional) The owner of the security group. Required if admin wants to create a port for another tenant. Changing this creates a new security group rule. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "direction",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "ethertype",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "port_range_min",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "port_range_max",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "remote_ip_prefix",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "remote_group_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above. ## Import Security Group Rules can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import openstack_networking_secgroup_rule_v2.secgroup_rule_1 aeb68ee3-6e9d-4256-955c-9584a6212745 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "direction",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "ethertype",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "port_range_min",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "port_range_max",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "remote_ip_prefix",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "remote_group_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above. ## Import Security Group Rules can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import openstack_networking_secgroup_rule_v2.secgroup_rule_1 aeb68ee3-6e9d-4256-955c-9584a6212745 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "openstack_networking_secgroup_v2",
			Category:         "Networking Resources",
			ShortDescription: `Manages a V2 Neutron security group resource within OpenStack.`,
			Description:      ``,
			Keywords: []string{
				"networking",
				"secgroup",
				"v2",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 networking client. A networking client is needed to create a port. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new security group.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A unique name for the security group.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A unique name for the security group.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Optional) The owner of the security group. Required if admin wants to create a port for another tenant. Changing this creates a new security group.`,
				},
				resource.Attribute{
					Name:        "delete_default_rules",
					Description: `(Optional) Whether or not to delete the default egress security rules. This is ` + "`" + `false` + "`" + ` by default. See the below note for more information.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A set of string tags for the security group. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "all_tags",
					Description: `The collection of tags assigned on the security group, which have been explicitly and implicitly added. ## Default Security Group Rules In most cases, OpenStack will create some egress security group rules for each new security group. These security group rules will not be managed by Terraform, so if you prefer to have`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "all_tags",
					Description: `The collection of tags assigned on the security group, which have been explicitly and implicitly added. ## Default Security Group Rules In most cases, OpenStack will create some egress security group rules for each new security group. These security group rules will not be managed by Terraform, so if you prefer to have`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "openstack_networking_subnet_route_v2",
			Category:         "Networking Resources",
			ShortDescription: `Creates a routing entry on a OpenStack V2 subnet.`,
			Description:      ``,
			Keywords: []string{
				"networking",
				"subnet",
				"route",
				"v2",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 networking client. A networking client is needed to configure a routing entry on a subnet. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new routing entry.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Required) ID of the subnet this routing entry belongs to. Changing this creates a new routing entry.`,
				},
				resource.Attribute{
					Name:        "destination_cidr",
					Description: `(Required) CIDR block to match on the packet’s destination IP. Changing this creates a new routing entry.`,
				},
				resource.Attribute{
					Name:        "next_hop",
					Description: `(Required) IP address of the next hop gateway. Changing this creates a new routing entry. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "destination_cidr",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "next_hop",
					Description: `See Argument Reference above. ## Notes ## Import Routing entries can be imported using a combined ID using the following format: ` + "`" + `` + "`" + `<subnet_id>-route-<destination_cidr>-<next_hop>` + "`" + `` + "`" + ` ` + "`" + `` + "`" + `` + "`" + ` $ terraform import openstack_networking_subnet_route_v2.subnet_route_1 686fe248-386c-4f70-9f6c-281607dad079-route-10.0.1.0/24-192.168.199.25 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "destination_cidr",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "next_hop",
					Description: `See Argument Reference above. ## Notes ## Import Routing entries can be imported using a combined ID using the following format: ` + "`" + `` + "`" + `<subnet_id>-route-<destination_cidr>-<next_hop>` + "`" + `` + "`" + ` ` + "`" + `` + "`" + `` + "`" + ` $ terraform import openstack_networking_subnet_route_v2.subnet_route_1 686fe248-386c-4f70-9f6c-281607dad079-route-10.0.1.0/24-192.168.199.25 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "openstack_networking_subnet_v2",
			Category:         "Networking Resources",
			ShortDescription: `Manages a V2 Neutron subnet resource within OpenStack.`,
			Description:      ``,
			Keywords: []string{
				"networking",
				"subnet",
				"v2",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 Networking client. A Networking client is needed to create a Neutron subnet. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new subnet.`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `(Required) The UUID of the parent network. Changing this creates a new subnet.`,
				},
				resource.Attribute{
					Name:        "cidr",
					Description: `(Optional) CIDR representing IP range for this subnet, based on IP version. You can omit this option if you are creating a subnet from a subnet pool.`,
				},
				resource.Attribute{
					Name:        "prefix_length",
					Description: `(Optional) The prefix length to use when creating a subnet from a subnet pool. The default subnet pool prefix length that was defined when creating the subnet pool will be used if not provided. Changing this creates a new subnet.`,
				},
				resource.Attribute{
					Name:        "ip_version",
					Description: `(Optional) IP version, either 4 (default) or 6. Changing this creates a new subnet.`,
				},
				resource.Attribute{
					Name:        "ipv6_address_mode",
					Description: `(Optional) The IPv6 address mode. Valid values are ` + "`" + `dhcpv6-stateful` + "`" + `, ` + "`" + `dhcpv6-stateless` + "`" + `, or ` + "`" + `slaac` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ipv6_ra_mode",
					Description: `(Optional) The IPv6 Router Advertisement mode. Valid values are ` + "`" + `dhcpv6-stateful` + "`" + `, ` + "`" + `dhcpv6-stateless` + "`" + `, or ` + "`" + `slaac` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the subnet. Changing this updates the name of the existing subnet.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Human-readable description of the subnet. Changing this updates the name of the existing subnet.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Optional) The owner of the subnet. Required if admin wants to create a subnet for another tenant. Changing this creates a new subnet.`,
				},
				resource.Attribute{
					Name:        "allocation_pools",
					Description: `(`,
				},
				resource.Attribute{
					Name:        "allocation_pool",
					Description: `(Optional) A block declaring the start and end range of the IP addresses available for use with DHCP in this subnet. Multiple ` + "`" + `allocation_pool` + "`" + ` blocks can be declared, providing the subnet with more than one range of IP addresses to use with DHCP. However, each IP range must be from the same CIDR that the subnet is part of. The ` + "`" + `allocation_pool` + "`" + ` block is documented below.`,
				},
				resource.Attribute{
					Name:        "gateway_ip",
					Description: `(Optional) Default gateway used by devices in this subnet. Leaving this blank and not setting ` + "`" + `no_gateway` + "`" + ` will cause a default gateway of ` + "`" + `.1` + "`" + ` to be used. Changing this updates the gateway IP of the existing subnet.`,
				},
				resource.Attribute{
					Name:        "no_gateway",
					Description: `(Optional) Do not set a gateway IP on this subnet. Changing this removes or adds a default gateway IP of the existing subnet.`,
				},
				resource.Attribute{
					Name:        "enable_dhcp",
					Description: `(Optional) The administrative state of the network. Acceptable values are "true" and "false". Changing this value enables or disables the DHCP capabilities of the existing subnet. Defaults to true.`,
				},
				resource.Attribute{
					Name:        "dns_nameservers",
					Description: `(Optional) An array of DNS name server names used by hosts in this subnet. Changing this updates the DNS name servers for the existing subnet.`,
				},
				resource.Attribute{
					Name:        "host_routes",
					Description: `(`,
				},
				resource.Attribute{
					Name:        "subnetpool_id",
					Description: `(Optional) The ID of the subnetpool associated with the subnet.`,
				},
				resource.Attribute{
					Name:        "value_specs",
					Description: `(Optional) Map of additional options.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A set of string tags for the subnet. The deprecated ` + "`" + `allocation_pools` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "start",
					Description: `(Required) The starting address.`,
				},
				resource.Attribute{
					Name:        "end",
					Description: `(Required) The ending address. The ` + "`" + `allocation_pool` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "start",
					Description: `(Required) The starting address.`,
				},
				resource.Attribute{
					Name:        "end",
					Description: `(Required) The ending address. The ` + "`" + `host_routes` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "destination_cidr",
					Description: `(Required) The destination CIDR.`,
				},
				resource.Attribute{
					Name:        "next_hop",
					Description: `(Required) The next hop in the route. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "cidr",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "ip_version",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "allocation_pools",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "gateway_ip",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "enable_dhcp",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "dns_nameservers",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "host_routes",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "subnetpool_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "all_tags",
					Description: `The collection of ags assigned on the subnet, which have been explicitly and implicitly added. ## Import Subnets can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import openstack_networking_subnet_v2.subnet_1 da4faf16-5546-41e4-8330-4d0002b74048 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "cidr",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "ip_version",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "allocation_pools",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "gateway_ip",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "enable_dhcp",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "dns_nameservers",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "host_routes",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "subnetpool_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "all_tags",
					Description: `The collection of ags assigned on the subnet, which have been explicitly and implicitly added. ## Import Subnets can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import openstack_networking_subnet_v2.subnet_1 da4faf16-5546-41e4-8330-4d0002b74048 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "openstack_networking_subnetpool_v2",
			Category:         "Networking Resources",
			ShortDescription: `Manages a V2 Neutron subnetpool resource within OpenStack.`,
			Description:      ``,
			Keywords: []string{
				"networking",
				"subnetpool",
				"v2",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 Networking client. A Networking client is needed to create a Neutron subnetpool. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new subnetpool.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the subnetpool. Changing this updates the name of the existing subnetpool.`,
				},
				resource.Attribute{
					Name:        "default_quota",
					Description: `(Optional) The per-project quota on the prefix space that can be allocated from the subnetpool for project subnets. Changing this updates the default quota of the existing subnetpool.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Optional) The owner of the subnetpool. Required if admin wants to create a subnetpool for another project. Changing this creates a new subnetpool.`,
				},
				resource.Attribute{
					Name:        "prefixes",
					Description: `(Required) A list of subnet prefixes to assign to the subnetpool. Neutron API merges adjacent prefixes and treats them as a single prefix. Each subnet prefix must be unique among all subnet prefixes in all subnetpools that are associated with the address scope. Changing this updates the prefixes list of the existing subnetpool.`,
				},
				resource.Attribute{
					Name:        "default_prefixlen",
					Description: `(Optional) The size of the prefix to allocate when the cidr or prefixlen attributes are omitted when you create the subnet. Defaults to the MinPrefixLen. Changing this updates the default prefixlen of the existing subnetpool.`,
				},
				resource.Attribute{
					Name:        "min_prefixlen",
					Description: `(Optional) The smallest prefix that can be allocated from a subnetpool. For IPv4 subnetpools, default is 8. For IPv6 subnetpools, default is 64. Changing this updates the min prefixlen of the existing subnetpool.`,
				},
				resource.Attribute{
					Name:        "max_prefixlen",
					Description: `(Optional) The maximum prefix size that can be allocated from the subnetpool. For IPv4 subnetpools, default is 32. For IPv6 subnetpools, default is 128. Changing this updates the max prefixlen of the existing subnetpool.`,
				},
				resource.Attribute{
					Name:        "address_scope_id",
					Description: `(Optional) The Neutron address scope to assign to the subnetpool. Changing this updates the address scope id of the existing subnetpool.`,
				},
				resource.Attribute{
					Name:        "shared",
					Description: `(Optional) Indicates whether this subnetpool is shared across all projects. Changing this updates the shared status of the existing subnetpool.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The human-readable description for the subnetpool. Changing this updates the description of the existing subnetpool.`,
				},
				resource.Attribute{
					Name:        "is_default",
					Description: `(Optional) Indicates whether the subnetpool is default subnetpool or not. Changing this updates the default status of the existing subnetpool.`,
				},
				resource.Attribute{
					Name:        "value_specs",
					Description: `(Optional) Map of additional options.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A set of string tags for the subnetpool. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "default_quota",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `The time at which subnetpool was created.`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `The time at which subnetpool was created.`,
				},
				resource.Attribute{
					Name:        "prefixes",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "default_prefixlen",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "min_prefixlen",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "max_prefixlen",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "address_scope_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "ip_version",
					Description: `The IP protocol version.`,
				},
				resource.Attribute{
					Name:        "shared",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "is_default",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "revision_number",
					Description: `The revision number of the subnetpool.`,
				},
				resource.Attribute{
					Name:        "value_specs",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "all_tags",
					Description: `The collection of tags assigned on the subnetpool, which have been explicitly and implicitly added. ## Import Subnetpools can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import openstack_networking_subnetpool_v2.subnetpool_1 832cb7f3-59fe-40cf-8f64-8350ffc03272 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "default_quota",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `The time at which subnetpool was created.`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `The time at which subnetpool was created.`,
				},
				resource.Attribute{
					Name:        "prefixes",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "default_prefixlen",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "min_prefixlen",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "max_prefixlen",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "address_scope_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "ip_version",
					Description: `The IP protocol version.`,
				},
				resource.Attribute{
					Name:        "shared",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "is_default",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "revision_number",
					Description: `The revision number of the subnetpool.`,
				},
				resource.Attribute{
					Name:        "value_specs",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "all_tags",
					Description: `The collection of tags assigned on the subnetpool, which have been explicitly and implicitly added. ## Import Subnetpools can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import openstack_networking_subnetpool_v2.subnetpool_1 832cb7f3-59fe-40cf-8f64-8350ffc03272 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "openstack_networking_trunk_v2",
			Category:         "Networking Resources",
			ShortDescription: `Manages a networking V2 trunk resource within OpenStack.`,
			Description:      ``,
			Keywords: []string{
				"networking",
				"trunk",
				"v2",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 networking client. A networking client is needed to create a trunk. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new trunk.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) A unique name for the trunk. Changing this updates the ` + "`" + `name` + "`" + ` of an existing trunk.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Human-readable description of the trunk. Changing this updates the name of the existing trunk.`,
				},
				resource.Attribute{
					Name:        "port_id",
					Description: `(Required) The ID of the port to be used as the parent port of the trunk. This is the port that should be used as the compute instance network port. Changing this creates a new trunk.`,
				},
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `(Optional) Administrative up/down status for the trunk (must be "true" or "false" if provided). Changing this updates the ` + "`" + `admin_state_up` + "`" + ` of an existing trunk.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Optional) The owner of the Trunk. Required if admin wants to create a trunk on behalf of another tenant. Changing this creates a new trunk.`,
				},
				resource.Attribute{
					Name:        "sub_port",
					Description: `(Optional) The set of ports that will be made subports of the trunk. The structure of each subport is described below.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A set of string tags for the port. The ` + "`" + `sub_port` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "port_id",
					Description: `(Required) The ID of the port to be made a subport of the trunk.`,
				},
				resource.Attribute{
					Name:        "segmentation_type",
					Description: `(Required) The segmentation technology to use, e.g., "vlan".`,
				},
				resource.Attribute{
					Name:        "segmentation_id",
					Description: `(Required) The numeric id of the subport segment. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "port_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "sub_port",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "all_tags",
					Description: `The collection of tags assigned on the trunk, which have been explicitly and implicitly added.`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "port_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "sub_port",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "all_tags",
					Description: `The collection of tags assigned on the trunk, which have been explicitly and implicitly added.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "openstack_objectstorage_container_v1",
			Category:         "Object Storage Resources",
			ShortDescription: `Manages a V1 container resource within OpenStack.`,
			Description:      ``,
			Keywords: []string{
				"object",
				"storage",
				"objectstorage",
				"container",
				"v1",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to create the container. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new container.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A unique name for the container. Changing this creates a new container.`,
				},
				resource.Attribute{
					Name:        "container_read",
					Description: `(Optional) Sets an access control list (ACL) that grants read access. This header can contain a comma-delimited list of users that can read the container (allows the GET method for all objects in the container). Changing this updates the access control list read access.`,
				},
				resource.Attribute{
					Name:        "container_sync_to",
					Description: `(Optional) The destination for container synchronization. Changing this updates container synchronization.`,
				},
				resource.Attribute{
					Name:        "container_sync_key",
					Description: `(Optional) The secret key for container synchronization. Changing this updates container synchronization.`,
				},
				resource.Attribute{
					Name:        "container_write",
					Description: `(Optional) Sets an ACL that grants write access. Changing this updates the access control list write access.`,
				},
				resource.Attribute{
					Name:        "versioning",
					Description: `(Optional) Enable object versioning. The structure is described below.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `(Optional) Custom key/value pairs to associate with the container. Changing this updates the existing container metadata.`,
				},
				resource.Attribute{
					Name:        "content_type",
					Description: `(Optional) The MIME type for the container. Changing this updates the MIME type.`,
				},
				resource.Attribute{
					Name:        "force_destroy",
					Description: `(Optional, Default:false ) A boolean that indicates all objects should be deleted from the container so that the container can be destroyed without error. These objects are not recoverable. The ` + "`" + `versioning` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) Versioning type which can be ` + "`" + `versions` + "`" + ` or ` + "`" + `history` + "`" + ` according to [Openstack documentation](https://docs.openstack.org/swift/latest/overview_object_versioning.html).`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `(Required) Container in which versions will be stored. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "container_read",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "container_sync_to",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "container_sync_key",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "container_write",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "versioning",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "content_type",
					Description: `See Argument Reference above.`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "container_read",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "container_sync_to",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "container_sync_key",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "container_write",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "versioning",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "content_type",
					Description: `See Argument Reference above.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "openstack_objectstorage_object_v1",
			Category:         "Object Storage Resources",
			ShortDescription: `Manages a V1 container object resource within OpenStack.`,
			Description:      ``,
			Keywords: []string{
				"object",
				"storage",
				"objectstorage",
				"v1",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "container_name",
					Description: `(Required) A unique (within an account) name for the container. The container name must be from 1 to 256 characters long and can start with any character and contain any pattern. Character set must be UTF-8. The container name cannot contain a slash (/) character because this character delimits the container and object name. For example, the path /v1/account/www/pages specifies the www container, not the www/pages container.`,
				},
				resource.Attribute{
					Name:        "content",
					Description: `(Optional) A string representing the content of the object. Conflicts with ` + "`" + `source` + "`" + ` and ` + "`" + `copy_from` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "content_disposition",
					Description: `(Optional) A string which specifies the override behavior for the browser. For example, this header might specify that the browser use a download program to save this file rather than show the file, which is the default.`,
				},
				resource.Attribute{
					Name:        "content_encoding",
					Description: `(Optional) A string representing the value of the Content-Encoding metadata.`,
				},
				resource.Attribute{
					Name:        "content_type",
					Description: `(Optional) A string which sets the MIME type for the object.`,
				},
				resource.Attribute{
					Name:        "copy_from",
					Description: `(Optional) A string representing the name of an object used to create the new object by copying the ` + "`" + `copy_from` + "`" + ` object. The value is in form {container}/{object}. You must UTF-8-encode and then URL-encode the names of the container and object before you include them in the header. Conflicts with ` + "`" + `source` + "`" + ` and ` + "`" + `content` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "delete_after",
					Description: `(Optional) An integer representing the number of seconds after which the system removes the object. Internally, the Object Storage system stores this value in the X-Delete-At metadata item.`,
				},
				resource.Attribute{
					Name:        "delete_at",
					Description: `(Optional) An string representing the date when the system removes the object. For example, "2015-08-26" is equivalent to Mon, Wed, 26 Aug 2015 00:00:00 GMT.`,
				},
				resource.Attribute{
					Name:        "detect_content_type",
					Description: `(Optional) If set to true, Object Storage guesses the content type based on the file extension and ignores the value sent in the Content-Type header, if present.`,
				},
				resource.Attribute{
					Name:        "etag",
					Description: `(Optional) Used to trigger updates. The only meaningful value is ${md5(file("path/to/file"))}.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A unique name for the object.`,
				},
				resource.Attribute{
					Name:        "object_manifest",
					Description: `(Optional) A string set to specify that this is a dynamic large object manifest object. The value is the container and object name prefix of the segment objects in the form container/prefix. You must UTF-8-encode and then URL-encode the names of the container and prefix before you include them in this header.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to create the container. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new container.`,
				},
				resource.Attribute{
					Name:        "source",
					Description: `(Optional) A string representing the local path of a file which will be used as the object's content. Conflicts with ` + "`" + `source` + "`" + ` and ` + "`" + `copy_from` + "`" + `. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "content_length",
					Description: `If the operation succeeds, this value is zero (0) or the length of informational or error text in the response body.`,
				},
				resource.Attribute{
					Name:        "content_type",
					Description: `If the operation succeeds, this value is the MIME type of the object. If the operation fails, this value is the MIME type of the error text in the response body.`,
				},
				resource.Attribute{
					Name:        "date",
					Description: `The date and time the system responded to the request, using the preferred format of RFC 7231 as shown in this example Thu, 16 Jun 2016 15:10:38 GMT. The time is always in UTC.`,
				},
				resource.Attribute{
					Name:        "etag",
					Description: `Whatever the value given in argument, will be overriden by the MD5 checksum of the uploaded object content. The value is not quoted. If it is an SLO, it would be MD5 checksum of the segments’ etags.`,
				},
				resource.Attribute{
					Name:        "last_modified",
					Description: `The date and time when the object was last modified. The date and time stamp format is ISO 8601: CCYY-MM-DDThh:mm:ss±hh:mm For example, 2015-08-27T09:49:58-05:00. The ±hh:mm value, if included, is the time zone as an offset from UTC. In the previous example, the offset value is -05:00.`,
				},
				resource.Attribute{
					Name:        "static_large_object",
					Description: `True if object is a multipart_manifest.`,
				},
				resource.Attribute{
					Name:        "trans_id",
					Description: `A unique transaction ID for this request. Your service provider might need this value if you report a problem.`,
				},
				resource.Attribute{
					Name:        "container_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "content",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "content_disposition",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "content_encoding",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "copy_from",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "delete_after",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "delete_at",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "detect_content_type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "object_manifest",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "source",
					Description: `See Argument Reference above.`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "content_length",
					Description: `If the operation succeeds, this value is zero (0) or the length of informational or error text in the response body.`,
				},
				resource.Attribute{
					Name:        "content_type",
					Description: `If the operation succeeds, this value is the MIME type of the object. If the operation fails, this value is the MIME type of the error text in the response body.`,
				},
				resource.Attribute{
					Name:        "date",
					Description: `The date and time the system responded to the request, using the preferred format of RFC 7231 as shown in this example Thu, 16 Jun 2016 15:10:38 GMT. The time is always in UTC.`,
				},
				resource.Attribute{
					Name:        "etag",
					Description: `Whatever the value given in argument, will be overriden by the MD5 checksum of the uploaded object content. The value is not quoted. If it is an SLO, it would be MD5 checksum of the segments’ etags.`,
				},
				resource.Attribute{
					Name:        "last_modified",
					Description: `The date and time when the object was last modified. The date and time stamp format is ISO 8601: CCYY-MM-DDThh:mm:ss±hh:mm For example, 2015-08-27T09:49:58-05:00. The ±hh:mm value, if included, is the time zone as an offset from UTC. In the previous example, the offset value is -05:00.`,
				},
				resource.Attribute{
					Name:        "static_large_object",
					Description: `True if object is a multipart_manifest.`,
				},
				resource.Attribute{
					Name:        "trans_id",
					Description: `A unique transaction ID for this request. Your service provider might need this value if you report a problem.`,
				},
				resource.Attribute{
					Name:        "container_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "content",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "content_disposition",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "content_encoding",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "copy_from",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "delete_after",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "delete_at",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "detect_content_type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "object_manifest",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "source",
					Description: `See Argument Reference above.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "openstack_objectstorage_tempurl_v1",
			Category:         "Object Storage Resources",
			ShortDescription: `Generate a TempURL for a Swift container and object.`,
			Description:      ``,
			Keywords: []string{
				"object",
				"storage",
				"objectstorage",
				"tempurl",
				"v1",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region the tempurl is located in.`,
				},
				resource.Attribute{
					Name:        "container",
					Description: `(Required) The container name the object belongs to.`,
				},
				resource.Attribute{
					Name:        "object",
					Description: `(Required) The object name the tempurl is for.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `(Required) The TTL, in seconds, for the URL. For how long it should be valid.`,
				},
				resource.Attribute{
					Name:        "method",
					Description: `(Optional) The method allowed when accessing this URL. Valid values are ` + "`" + `GET` + "`" + `, and ` + "`" + `POST` + "`" + `. Default is ` + "`" + `GET` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "regenerate",
					Description: `(Optional) Whether to automatically regenerate the URL when it has expired. If set to true, this will create a new resource with a new ID and new URL. Defaults to false. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Computed md5 hash based on the generated url`,
				},
				resource.Attribute{
					Name:        "container",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "object",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "method",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `The URL`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The region the endpoint is located in.`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "id",
					Description: `Computed md5 hash based on the generated url`,
				},
				resource.Attribute{
					Name:        "container",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "object",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "method",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `The URL`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The region the endpoint is located in.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "openstack_sharedfilesystem_securityservice_v2",
			Category:         "Shared File System",
			ShortDescription: `Configure a Shared File System security service.`,
			Description:      ``,
			Keywords: []string{
				"shared",
				"file",
				"system",
				"sharedfilesystem",
				"securityservice",
				"v2",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 Shared File System client. A Shared File System client is needed to create a security service. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new security service.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the security service. Changing this updates the name of the existing security service.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The human-readable description for the security service. Changing this updates the description of the existing security service.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) The security service type - can either be active\_directory, kerberos or ldap. Changing this updates the existing security service.`,
				},
				resource.Attribute{
					Name:        "dns_ip",
					Description: `(Optional) The security service DNS IP address that is used inside the tenant network.`,
				},
				resource.Attribute{
					Name:        "ou",
					Description: `(Optional) The security service ou. An organizational unit can be added to specify where the share ends up. New in Manila microversion 2.44.`,
				},
				resource.Attribute{
					Name:        "user",
					Description: `(Optional) The security service user or group name that is used by the tenant.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional) The user password, if you specify a user.`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `(Optional) The security service domain.`,
				},
				resource.Attribute{
					Name:        "server",
					Description: `(Optional) The security service host name or IP address. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique ID for the Security Service.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `The owner of the Security Service.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "dns_ip",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "ou",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "user",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "server",
					Description: `See Argument Reference above. ## Import This resource can be imported by specifying the ID of the security service: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import openstack_sharedfilesystem_securityservice_v2.securityservice_1 <id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "id",
					Description: `The unique ID for the Security Service.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `The owner of the Security Service.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "dns_ip",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "ou",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "user",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "server",
					Description: `See Argument Reference above. ## Import This resource can be imported by specifying the ID of the security service: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import openstack_sharedfilesystem_securityservice_v2.securityservice_1 <id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "openstack_sharedfilesystem_share_access_v2",
			Category:         "Shared File System",
			ShortDescription: `Configure a Shared File System share access list.`,
			Description:      ``,
			Keywords: []string{
				"shared",
				"file",
				"system",
				"sharedfilesystem",
				"share",
				"access",
				"v2",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "region",
					Description: `The region in which to obtain the V2 Shared File System client. A Shared File System client is needed to create a share access. Changing this creates a new share access.`,
				},
				resource.Attribute{
					Name:        "share_id",
					Description: `(Required) The UUID of the share to which you are granted access.`,
				},
				resource.Attribute{
					Name:        "access_type",
					Description: `(Required) The access rule type. Can either be an ip, user, cert, or cephx. cephx support requires an OpenStack environment that supports Shared Filesystem microversion 2.13 (Mitaka) or later.`,
				},
				resource.Attribute{
					Name:        "access_to",
					Description: `(Required) The value that defines the access. Can either be an IP address or a username verified by configured Security Service of the Share Network.`,
				},
				resource.Attribute{
					Name:        "access_level",
					Description: `(Required) The access level to the share. Can either be ` + "`" + `rw` + "`" + ` or ` + "`" + `ro` + "`" + `. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique ID for the Share Access.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "share_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "access_type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "access_to",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "access_level",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "access_key",
					Description: `The access credential of the entity granted access. ## Import This resource can be imported by specifying the ID of the share and the ID of the share access, separated by a slash, e.g.: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import openstack_sharedfilesystem_share_access_v2.share_access_1 <share id>/<share access id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "id",
					Description: `The unique ID for the Share Access.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "share_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "access_type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "access_to",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "access_level",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "access_key",
					Description: `The access credential of the entity granted access. ## Import This resource can be imported by specifying the ID of the share and the ID of the share access, separated by a slash, e.g.: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import openstack_sharedfilesystem_share_access_v2.share_access_1 <share id>/<share access id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "openstack_sharedfilesystem_share_v2",
			Category:         "Shared File System",
			ShortDescription: `Configure a Shared File System share.`,
			Description:      ``,
			Keywords: []string{
				"shared",
				"file",
				"system",
				"sharedfilesystem",
				"share",
				"v2",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "region",
					Description: `The region in which to obtain the V2 Shared File System client. A Shared File System client is needed to create a share. Changing this creates a new share.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the share. Changing this updates the name of the existing share.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The human-readable description for the share. Changing this updates the description of the existing share.`,
				},
				resource.Attribute{
					Name:        "share_proto",
					Description: `(Required) The share protocol - can either be NFS, CIFS, CEPHFS, GLUSTERFS, HDFS or MAPRFS. Changing this creates a new share.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `(Required) The share size, in GBs. The requested share size cannot be greater than the allowed GB quota. Changing this resizes the existing share.`,
				},
				resource.Attribute{
					Name:        "share_type",
					Description: `(Optional) The share type name. If you omit this parameter, the default share type is used.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `(Optional) The UUID of the share's base snapshot. Changing this creates a new share.`,
				},
				resource.Attribute{
					Name:        "is_public",
					Description: `(Optional) The level of visibility for the share. Set to true to make share public. Set to false to make it private. Default value is false. Changing this updates the existing share.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `(Optional) One or more metadata key and value pairs as a dictionary of strings.`,
				},
				resource.Attribute{
					Name:        "share_network_id",
					Description: `(Optional) The UUID of a share network where the share server exists or will be created. If ` + "`" + `share_network_id` + "`" + ` is not set and you provide a ` + "`" + `snapshot_id` + "`" + `, the share_network_id value from the snapshot is used. Changing this creates a new share.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `(Optional) The share availability zone. Changing this creates a new share. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique ID for the Share.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `The owner of the Share.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "share_proto",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "share_type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "is_public",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "share_network_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "export_locations",
					Description: `A list of export locations. For example, when a share server has more than one network interface, it can have multiple export locations.`,
				},
				resource.Attribute{
					Name:        "has_replicas",
					Description: `Indicates whether a share has replicas or not.`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `The share host name.`,
				},
				resource.Attribute{
					Name:        "replication_type",
					Description: `The share replication type.`,
				},
				resource.Attribute{
					Name:        "share_server_id",
					Description: `The UUID of the share server. ## Import This resource can be imported by specifying the ID of the share: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import openstack_sharedfilesystem_share_v2.share_1 <id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "id",
					Description: `The unique ID for the Share.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `The owner of the Share.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "share_proto",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "share_type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "is_public",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "share_network_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "export_locations",
					Description: `A list of export locations. For example, when a share server has more than one network interface, it can have multiple export locations.`,
				},
				resource.Attribute{
					Name:        "has_replicas",
					Description: `Indicates whether a share has replicas or not.`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `The share host name.`,
				},
				resource.Attribute{
					Name:        "replication_type",
					Description: `The share replication type.`,
				},
				resource.Attribute{
					Name:        "share_server_id",
					Description: `The UUID of the share server. ## Import This resource can be imported by specifying the ID of the share: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import openstack_sharedfilesystem_share_v2.share_1 <id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "openstack_sharedfilesystem_sharenetwork_v2",
			Category:         "Shared File System",
			ShortDescription: `Configure a Shared File System share network.`,
			Description:      ``,
			Keywords: []string{
				"shared",
				"file",
				"system",
				"sharedfilesystem",
				"sharenetwork",
				"v2",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 Shared File System client. A Shared File System client is needed to create a share network. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new share network.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name for the share network. Changing this updates the name of the existing share network.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The human-readable description for the share network. Changing this updates the description of the existing share network.`,
				},
				resource.Attribute{
					Name:        "neutron_net_id",
					Description: `(Required) The UUID of a neutron network when setting up or updating a share network. Changing this updates the existing share network if it's not used by shares.`,
				},
				resource.Attribute{
					Name:        "neutron_subnet_id",
					Description: `(Required) The UUID of the neutron subnet when setting up or updating a share network. Changing this updates the existing share network if it's not used by shares.`,
				},
				resource.Attribute{
					Name:        "security_service_ids",
					Description: `(Optional) The list of security service IDs to associate with the share network. The security service must be specified by ID and not name. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique ID for the Share Network.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `The owner of the Share Network.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "neutron_net_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "neutron_subnet_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "security_service_ids",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "network_type",
					Description: `The share network type. Can either be VLAN, VXLAN, GRE, or flat.`,
				},
				resource.Attribute{
					Name:        "segmentation_id",
					Description: `The share network segmentation ID.`,
				},
				resource.Attribute{
					Name:        "cidr",
					Description: `The share network CIDR.`,
				},
				resource.Attribute{
					Name:        "ip_version",
					Description: `The IP version of the share network. Can either be 4 or 6. ## Import This resource can be imported by specifying the ID of the share network: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import openstack_sharedfilesystem_sharenetwork_v2.sharenetwork_1 <id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "id",
					Description: `The unique ID for the Share Network.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `The owner of the Share Network.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "neutron_net_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "neutron_subnet_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "security_service_ids",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "network_type",
					Description: `The share network type. Can either be VLAN, VXLAN, GRE, or flat.`,
				},
				resource.Attribute{
					Name:        "segmentation_id",
					Description: `The share network segmentation ID.`,
				},
				resource.Attribute{
					Name:        "cidr",
					Description: `The share network CIDR.`,
				},
				resource.Attribute{
					Name:        "ip_version",
					Description: `The IP version of the share network. Can either be 4 or 6. ## Import This resource can be imported by specifying the ID of the share network: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import openstack_sharedfilesystem_sharenetwork_v2.sharenetwork_1 <id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "openstack_vpnaas_endpoint_group_v2",
			Category:         "VPNaaS Resources",
			ShortDescription: `Manages a V2 Neutron Endpoint Group resource within OpenStack.`,
			Description:      ``,
			Keywords: []string{
				"vpnaas",
				"endpoint",
				"group",
				"v2",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 Networking client. A Networking client is needed to create an endpoint group. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new group.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the group. Changing this updates the name of the existing group.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Optional) The owner of the group. Required if admin wants to create an endpoint group for another project. Changing this creates a new group.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The human-readable description for the group. Changing this updates the description of the existing group.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of the endpoints in the group. A valid value is subnet, cidr, network, router, or vlan. Changing this creates a new group.`,
				},
				resource.Attribute{
					Name:        "endpoints",
					Description: `List of endpoints of the same type, for the endpoint group. The values will depend on the type. Changing this creates a new group.`,
				},
				resource.Attribute{
					Name:        "value_specs",
					Description: `(Optional) Map of additional options. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "endpoints",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "value_specs",
					Description: `See Argument Reference above. ## Import Groups can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import openstack_vpnaas_endpoint_group_v2.group_1 832cb7f3-59fe-40cf-8f64-8350ffc03272 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "endpoints",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "value_specs",
					Description: `See Argument Reference above. ## Import Groups can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import openstack_vpnaas_endpoint_group_v2.group_1 832cb7f3-59fe-40cf-8f64-8350ffc03272 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "openstack_vpnaas_ike_policy_v2",
			Category:         "VPNaaS Resources",
			ShortDescription: `Manages a V2 Neutron IKE policy resource within OpenStack.`,
			Description:      ``,
			Keywords: []string{
				"vpnaas",
				"ike",
				"policy",
				"v2",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 Networking client. A Networking client is needed to create a VPN service. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new service.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the policy. Changing this updates the name of the existing policy.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Optional) The owner of the policy. Required if admin wants to create a service for another policy. Changing this creates a new policy.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The human-readable description for the policy. Changing this updates the description of the existing policy.`,
				},
				resource.Attribute{
					Name:        "auth_algorithm",
					Description: `(Optional) The authentication hash algorithm. Valid values are sha1, sha256, sha384, sha512. Default is sha1. Changing this updates the algorithm of the existing policy.`,
				},
				resource.Attribute{
					Name:        "encryption_algorithm",
					Description: `(Optional) The encryption algorithm. Valid values are 3des, aes-128, aes-192 and so on. The default value is aes-128. Changing this updates the existing policy.`,
				},
				resource.Attribute{
					Name:        "pfs",
					Description: `(Optional) The perfect forward secrecy mode. Valid values are Group2, Group5 and Group14. Default is Group5. Changing this updates the existing policy.`,
				},
				resource.Attribute{
					Name:        "phase1_negotiation_mode",
					Description: `(Optional) The IKE mode. A valid value is main, which is the default. Changing this updates the existing policy.`,
				},
				resource.Attribute{
					Name:        "ike_version",
					Description: `(Optional) The IKE mode. A valid value is v1 or v2. Default is v1. Changing this updates the existing policy.`,
				},
				resource.Attribute{
					Name:        "lifetime",
					Description: `(Optional) The lifetime of the security association. Consists of Unit and Value. - ` + "`" + `unit` + "`" + ` - (Optional) The units for the lifetime of the security association. Can be either seconds or kilobytes. Default is seconds. - ` + "`" + `value` + "`" + ` - (Optional) The value for the lifetime of the security association. Must be a positive integer. Default is 3600.`,
				},
				resource.Attribute{
					Name:        "value_specs",
					Description: `(Optional) Map of additional options. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "auth_algorithm",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "encapsulation_mode",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "encryption_algorithm",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "pfs",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "transform_protocol",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "lifetime",
					Description: `See Argument Reference above. - ` + "`" + `unit` + "`" + ` - See Argument Reference above. - ` + "`" + `value` + "`" + ` - See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "value_specs",
					Description: `See Argument Reference above. ## Import Services can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import openstack_vpnaas_ike_policy_v2.policy_1 832cb7f3-59fe-40cf-8f64-8350ffc03272 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "auth_algorithm",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "encapsulation_mode",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "encryption_algorithm",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "pfs",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "transform_protocol",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "lifetime",
					Description: `See Argument Reference above. - ` + "`" + `unit` + "`" + ` - See Argument Reference above. - ` + "`" + `value` + "`" + ` - See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "value_specs",
					Description: `See Argument Reference above. ## Import Services can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import openstack_vpnaas_ike_policy_v2.policy_1 832cb7f3-59fe-40cf-8f64-8350ffc03272 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "openstack_vpnaas_ipsec_policy_v2",
			Category:         "VPNaaS Resources",
			ShortDescription: `Manages a V2 Neutron IPSec policy resource within OpenStack.`,
			Description:      ``,
			Keywords: []string{
				"vpnaas",
				"ipsec",
				"policy",
				"v2",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 Networking client. A Networking client is needed to create an IPSec policy. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new policy.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the policy. Changing this updates the name of the existing policy.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Optional) The owner of the policy. Required if admin wants to create a policy for another project. Changing this creates a new policy.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The human-readable description for the policy. Changing this updates the description of the existing policy.`,
				},
				resource.Attribute{
					Name:        "auth_algorithm",
					Description: `(Optional) The authentication hash algorithm. Valid values are sha1, sha256, sha384, sha512. Default is sha1. Changing this updates the algorithm of the existing policy.`,
				},
				resource.Attribute{
					Name:        "encapsulation_mode",
					Description: `(Optional) The encapsulation mode. Valid values are tunnel and transport. Default is tunnel. Changing this updates the existing policy.`,
				},
				resource.Attribute{
					Name:        "encryption_algorithm",
					Description: `(Optional) The encryption algorithm. Valid values are 3des, aes-128, aes-192 and so on. The default value is aes-128. Changing this updates the existing policy.`,
				},
				resource.Attribute{
					Name:        "pfs",
					Description: `(Optional) The perfect forward secrecy mode. Valid values are Group2, Group5 and Group14. Default is Group5. Changing this updates the existing policy.`,
				},
				resource.Attribute{
					Name:        "transform_protocol",
					Description: `(Optional) The transform protocol. Valid values are ESP, AH and AH-ESP. Changing this updates the existing policy. Default is ESP.`,
				},
				resource.Attribute{
					Name:        "lifetime",
					Description: `(Optional) The lifetime of the security association. Consists of Unit and Value. - ` + "`" + `unit` + "`" + ` - (Optional) The units for the lifetime of the security association. Can be either seconds or kilobytes. Default is seconds. - ` + "`" + `value` + "`" + ` - (Optional) The value for the lifetime of the security association. Must be a positive integer. Default is 3600.`,
				},
				resource.Attribute{
					Name:        "value_specs",
					Description: `(Optional) Map of additional options. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "auth_algorithm",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "encapsulation_mode",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "encryption_algorithm",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "pfs",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "transform_protocol",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "lifetime",
					Description: `See Argument Reference above. - ` + "`" + `unit` + "`" + ` - See Argument Reference above. - ` + "`" + `value` + "`" + ` - See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "value_specs",
					Description: `See Argument Reference above. ## Import Policies can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import openstack_vpnaas_ipsec_policy_v2.policy_1 832cb7f3-59fe-40cf-8f64-8350ffc03272 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "auth_algorithm",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "encapsulation_mode",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "encryption_algorithm",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "pfs",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "transform_protocol",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "lifetime",
					Description: `See Argument Reference above. - ` + "`" + `unit` + "`" + ` - See Argument Reference above. - ` + "`" + `value` + "`" + ` - See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "value_specs",
					Description: `See Argument Reference above. ## Import Policies can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import openstack_vpnaas_ipsec_policy_v2.policy_1 832cb7f3-59fe-40cf-8f64-8350ffc03272 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "openstack_vpnaas_service_v2",
			Category:         "VPNaaS Resources",
			ShortDescription: `Manages a V2 Neutron VPN service resource within OpenStack.`,
			Description:      ``,
			Keywords: []string{
				"vpnaas",
				"service",
				"v2",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 Networking client. A Networking client is needed to create a VPN service. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new service.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the service. Changing this updates the name of the existing service.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Optional) The owner of the service. Required if admin wants to create a service for another project. Changing this creates a new service.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The human-readable description for the service. Changing this updates the description of the existing service.`,
				},
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `(Optional) The administrative state of the resource. Can either be up(true) or down(false). Changing this updates the administrative state of the existing service.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Optional) SubnetID is the ID of the subnet. Default is null.`,
				},
				resource.Attribute{
					Name:        "router_id",
					Description: `(Required) The ID of the router. Changing this creates a new service.`,
				},
				resource.Attribute{
					Name:        "value_specs",
					Description: `(Optional) Map of additional options. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "router_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Indicates whether IPsec VPN service is currently operational. Values are ACTIVE, DOWN, BUILD, ERROR, PENDING_CREATE, PENDING_UPDATE, or PENDING_DELETE.`,
				},
				resource.Attribute{
					Name:        "external_v6_ip",
					Description: `The read-only external (public) IPv6 address that is used for the VPN service.`,
				},
				resource.Attribute{
					Name:        "external_v4_ip",
					Description: `The read-only external (public) IPv4 address that is used for the VPN service.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "value_specs",
					Description: `See Argument Reference above. ## Import Services can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import openstack_vpnaas_service_v2.service_1 832cb7f3-59fe-40cf-8f64-8350ffc03272 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "router_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Indicates whether IPsec VPN service is currently operational. Values are ACTIVE, DOWN, BUILD, ERROR, PENDING_CREATE, PENDING_UPDATE, or PENDING_DELETE.`,
				},
				resource.Attribute{
					Name:        "external_v6_ip",
					Description: `The read-only external (public) IPv6 address that is used for the VPN service.`,
				},
				resource.Attribute{
					Name:        "external_v4_ip",
					Description: `The read-only external (public) IPv4 address that is used for the VPN service.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "value_specs",
					Description: `See Argument Reference above. ## Import Services can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import openstack_vpnaas_service_v2.service_1 832cb7f3-59fe-40cf-8f64-8350ffc03272 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "openstack_vpnaas_site_connection_v2",
			Category:         "VPNaaS Resources",
			ShortDescription: `Manages a V2 Neutron IPSec site connection resource within OpenStack.`,
			Description:      ``,
			Keywords: []string{
				"vpnaas",
				"site",
				"connection",
				"v2",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 Networking client. A Networking client is needed to create an IPSec site connection. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used. Changing this creates a new site connection.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the connection. Changing this updates the name of the existing connection.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Optional) The owner of the connection. Required if admin wants to create a connection for another project. Changing this creates a new connection.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The human-readable description for the connection. Changing this updates the description of the existing connection.`,
				},
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `(Optional) The administrative state of the resource. Can either be up(true) or down(false). Changing this updates the administrative state of the existing connection.`,
				},
				resource.Attribute{
					Name:        "ikepolicy_id",
					Description: `(Required) The ID of the IKE policy. Changing this creates a new connection.`,
				},
				resource.Attribute{
					Name:        "vpnservice_id",
					Description: `(Required) The ID of the VPN service. Changing this creates a new connection.`,
				},
				resource.Attribute{
					Name:        "local_ep_group_id",
					Description: `(Optional) The ID for the endpoint group that contains private subnets for the local side of the connection. You must specify this parameter with the peer_ep_group_id parameter unless in backward- compatible mode where peer_cidrs is provided with a subnet_id for the VPN service. Changing this updates the existing connection.`,
				},
				resource.Attribute{
					Name:        "ipsecpolicy_id",
					Description: `(Required) The ID of the IPsec policy. Changing this creates a new connection.`,
				},
				resource.Attribute{
					Name:        "peer_id",
					Description: `(Required) The peer router identity for authentication. A valid value is an IPv4 address, IPv6 address, e-mail address, key ID, or FQDN. Typically, this value matches the peer_address value. Changing this updates the existing policy.`,
				},
				resource.Attribute{
					Name:        "peer_ep_group_id",
					Description: `(Optional) The ID for the endpoint group that contains private CIDRs in the form < net_address > / < prefix > for the peer side of the connection. You must specify this parameter with the local_ep_group_id parameter unless in backward-compatible mode where peer_cidrs is provided with a subnet_id for the VPN service.`,
				},
				resource.Attribute{
					Name:        "local_id",
					Description: `(Optional) An ID to be used instead of the external IP address for a virtual router used in traffic between instances on different networks in east-west traffic. Most often, local ID would be domain name, email address, etc. If this is not configured then the external IP address will be used as the ID.`,
				},
				resource.Attribute{
					Name:        "peer_address",
					Description: `(Required) The peer gateway public IPv4 or IPv6 address or FQDN.`,
				},
				resource.Attribute{
					Name:        "psk",
					Description: `(Required) The pre-shared key. A valid value is any string.`,
				},
				resource.Attribute{
					Name:        "initiator",
					Description: `(Optional) A valid value is response-only or bi-directional. Default is bi-directional.`,
				},
				resource.Attribute{
					Name:        "peer_cidrs",
					Description: `(Optional) Unique list of valid peer private CIDRs in the form < net_address > / < prefix > .`,
				},
				resource.Attribute{
					Name:        "dpd",
					Description: `(Optional) A dictionary with dead peer detection (DPD) protocol controls. - ` + "`" + `action` + "`" + ` - (Optional) The dead peer detection (DPD) action. A valid value is clear, hold, restart, disabled, or restart-by-peer. Default value is hold. - ` + "`" + `timeout` + "`" + ` - (Optional) The dead peer detection (DPD) timeout in seconds. A valid value is a positive integer that is greater than the DPD interval value. Default is 120. - ` + "`" + `interval` + "`" + ` - (Optional) The dead peer detection (DPD) interval, in seconds. A valid value is a positive integer. Default is 30.`,
				},
				resource.Attribute{
					Name:        "mtu",
					Description: `(Optional) The maximum transmission unit (MTU) value to address fragmentation. Minimum value is 68 for IPv4, and 1280 for IPv6.`,
				},
				resource.Attribute{
					Name:        "value_specs",
					Description: `(Optional) Map of additional options. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "dpd",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "psk",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "initiator",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "peer_address",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "peer_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "peer_cidrs",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "mtu",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "local_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "peer_ep_group_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "ipsecpolicy_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "vpnservice_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "ikepolicy_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "value_specs",
					Description: `See Argument Reference above. ## Import Site Connections can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import openstack_vpnaas_site_connection_v2.conn_1 832cb7f3-59fe-40cf-8f64-8350ffc03272 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "dpd",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "psk",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "initiator",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "peer_address",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "peer_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "peer_cidrs",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "mtu",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "local_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "peer_ep_group_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "ipsecpolicy_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "vpnservice_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "ikepolicy_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "value_specs",
					Description: `See Argument Reference above. ## Import Site Connections can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import openstack_vpnaas_site_connection_v2.conn_1 832cb7f3-59fe-40cf-8f64-8350ffc03272 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
	}

	resourcesMap = map[string]int{

		"openstack_blockstorage_volume_attach_v2":            0,
		"openstack_blockstorage_volume_attach_v3":            1,
		"openstack_blockstorage_volume_v1":                   2,
		"openstack_blockstorage_volume_v2":                   3,
		"openstack_blockstorage_volume_v3":                   4,
		"openstack_compute_flavor_access_v2":                 5,
		"openstack_compute_flavor_v2":                        6,
		"openstack_compute_floatingip_associate_v2":          7,
		"openstack_compute_floatingip_v2":                    8,
		"openstack_compute_instance_v2":                      9,
		"openstack_compute_interface_attach_v2":              10,
		"openstack_compute_keypair_v2":                       11,
		"openstack_compute_secgroup_v2":                      12,
		"openstack_compute_servergroup_v2":                   13,
		"openstack_compute_volume_attach_v2":                 14,
		"openstack_containerinfra_cluster_v1":                15,
		"openstack_containerinfra_clustertemplate_v1":        16,
		"openstack_db_configuration_v1":                      17,
		"openstack_db_database_v1":                           18,
		"openstack_db_instance_v1":                           19,
		"openstack_db_user_v1":                               20,
		"openstack_dns_recordset_v2":                         21,
		"openstack_dns_zone_v2":                              22,
		"openstack_fw_firewall_v1":                           23,
		"openstack_fw_policy_v1":                             24,
		"openstack_fw_rule_v1":                               25,
		"openstack_identity_application_credential_v3":       26,
		"openstack_identity_project_v3":                      27,
		"openstack_identity_role_assignment_v3":              28,
		"openstack_identity_role_v3":                         29,
		"openstack_identity_user_v3":                         30,
		"openstack_images_image_v2":                          31,
		"openstack_lb_l7policy_v2":                           32,
		"openstack_lb_l7rule_v2":                             33,
		"openstack_lb_listener_v2":                           34,
		"openstack_lb_loadbalancer_v2":                       35,
		"openstack_lb_member_v1":                             36,
		"openstack_lb_member_v2":                             37,
		"openstack_lb_monitor_v1":                            38,
		"openstack_lb_monitor_v2":                            39,
		"openstack_lb_pool_v1":                               40,
		"openstack_lb_pool_v2":                               41,
		"openstack_lb_vip_v1":                                42,
		"openstack_networking_addressscope_v2":               43,
		"openstack_networking_floatingip_associate_v2":       44,
		"openstack_networking_floatingip_v2":                 45,
		"openstack_networking_network_v2":                    46,
		"openstack_networking_port_secgroup_associate_v2":    47,
		"openstack_networking_port_v2":                       48,
		"openstack_networking_qos_bandwidth_limit_rule_v2":   49,
		"openstack_networking_qos_dscp_marking_rule_v2":      50,
		"openstack_networking_qos_minimum_bandwidth_rule_v2": 51,
		"openstack_networking_qos_policy_v2":                 52,
		"openstack_networking_router_interface_v2":           53,
		"openstack_networking_router_route_v2":               54,
		"openstack_networking_router_v2":                     55,
		"openstack_networking_secgroup_rule_v2":              56,
		"openstack_networking_secgroup_v2":                   57,
		"openstack_networking_subnet_route_v2":               58,
		"openstack_networking_subnet_v2":                     59,
		"openstack_networking_subnetpool_v2":                 60,
		"openstack_networking_trunk_v2":                      61,
		"openstack_objectstorage_container_v1":               62,
		"openstack_objectstorage_object_v1":                  63,
		"openstack_objectstorage_tempurl_v1":                 64,
		"openstack_sharedfilesystem_securityservice_v2":      65,
		"openstack_sharedfilesystem_share_access_v2":         66,
		"openstack_sharedfilesystem_share_v2":                67,
		"openstack_sharedfilesystem_sharenetwork_v2":         68,
		"openstack_vpnaas_endpoint_group_v2":                 69,
		"openstack_vpnaas_ike_policy_v2":                     70,
		"openstack_vpnaas_ipsec_policy_v2":                   71,
		"openstack_vpnaas_service_v2":                        72,
		"openstack_vpnaas_site_connection_v2":                73,
	}
)

func GetResource(r string) (*resouce.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs]
}
