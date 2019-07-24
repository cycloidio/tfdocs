package aws

import (
	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*Resource{

		&resource.Resource{
			Name:             "",
			Type:             "rightscale_cloud",
			Category:         "Data Sources",
			ShortDescription: `Defines a cloud datasource to operate against.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Cloud name as displayed in cm platform. Pattern match.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Cloud description as displayed in cm platform. Pattern match.`,
				},
				resource.Attribute{
					Name:        "cloud_type",
					Description: `(Optional) Cloud type as referenced in cm platform. Common types include: amazon, google, azure, and vscale. See [supportedCloudTypes](https://github.com/terraform-providers/terraform-provider-rightscale/blob/master/rightscale/data_source_cloud.go#L95) for complete list. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Official cloud name as displayed in cm platform.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `Display name for cloud as displayed in cm platform.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Cloud description as displayed in cm platform.`,
				},
				resource.Attribute{
					Name:        "cloud_type",
					Description: `Cloud type as referenced in cm platform.`,
				},
				resource.Attribute{
					Name:        "links",
					Description: `Hrefs of related API resources.`,
				},
				resource.Attribute{
					Name:        "href",
					Description: `Href of the cloud.`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `Official cloud name as displayed in cm platform.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `Display name for cloud as displayed in cm platform.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Cloud description as displayed in cm platform.`,
				},
				resource.Attribute{
					Name:        "cloud_type",
					Description: `Cloud type as referenced in cm platform.`,
				},
				resource.Attribute{
					Name:        "links",
					Description: `Hrefs of related API resources.`,
				},
				resource.Attribute{
					Name:        "href",
					Description: `Href of the cloud.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "rightscale_credential",
			Category:         "Data Sources",
			ShortDescription: `Defines a credential datasource to operate against.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "view",
					Description: `(Optional) Set this to 'default' to NOT request credential value with api response. This allows use of existing credential with other rightscale provider resources (extracting href and handing to other resources). Offereed in case user lacks rs account privs sufficient to view credential values. The ` + "`" + `filter` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Credential name. Pattern match.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of credential. Pattern match. ## Attributes Reference The following attributes are exported:`,
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
					Description: `(Contextual) Available unless if 'default' view is set. Value of the credential.`,
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
				resource.Attribute{
					Name:        "href",
					Description: `Href of the credential.`,
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
					Description: `(Contextual) Available unless if 'default' view is set. Value of the credential.`,
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
				resource.Attribute{
					Name:        "href",
					Description: `Href of the credential.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "rightscale_datacenter",
			Category:         "Data Sources",
			ShortDescription: `Defines a datacenter datasource to operate against.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `The name of the datacenter`,
				},
				resource.Attribute{
					Name:        "resource_uid",
					Description: `The resource_uid of the datacenter. If this filter option is set, additional retry logic will be enabled to wait up to 5 minutes for cloud resources to be polled and populated for use. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the datacenter`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the datacenter`,
				},
				resource.Attribute{
					Name:        "resource_uid",
					Description: `The resource_uid of the datacenter as reported by the rightscale platform`,
				},
				resource.Attribute{
					Name:        "links",
					Description: `Hrefs of related API resources`,
				},
				resource.Attribute{
					Name:        "cloud_href",
					Description: `Href of the cloud the datacenter belongs to`,
				},
				resource.Attribute{
					Name:        "href",
					Description: `Href of the datacenter`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `The name of the datacenter`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the datacenter`,
				},
				resource.Attribute{
					Name:        "resource_uid",
					Description: `The resource_uid of the datacenter as reported by the rightscale platform`,
				},
				resource.Attribute{
					Name:        "links",
					Description: `Hrefs of related API resources`,
				},
				resource.Attribute{
					Name:        "cloud_href",
					Description: `Href of the cloud the datacenter belongs to`,
				},
				resource.Attribute{
					Name:        "href",
					Description: `Href of the datacenter`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "rightscale_deployment",
			Category:         "Data Sources",
			ShortDescription: `Defines a deployment datasource to operate against.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "view",
					Description: `(Optional) Options include 'default,' 'inputs' or 'inputs_2_0.' Defaults to 'default.' Please see RightScale documentation for inputs for details on these different views.`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) Filter block to find matching deployment. The ` + "`" + `filter` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Credential name. Pattern match.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of credential. Pattern match.`,
				},
				resource.Attribute{
					Name:        "resource_group_href",
					Description: `(Optional) Resource group href to filter on.`,
				},
				resource.Attribute{
					Name:        "server_tag_scope",
					Description: `(Optional) Tag routing scope to filter on. Pattern match. ## Attributes Reference The following attributes are exported:`,
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
					Name:        "links",
					Description: `Hrefs of related API resources.`,
				},
				resource.Attribute{
					Name:        "locked",
					Description: `Displays if the deployment is locked or not.`,
				},
				resource.Attribute{
					Name:        "server_tag_scope",
					Description: `Displays what the scope of tags are in the deployment. Options are "deployment" or "account."`,
				},
				resource.Attribute{
					Name:        "href",
					Description: `Href of the deployment.`,
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
					Name:        "links",
					Description: `Hrefs of related API resources.`,
				},
				resource.Attribute{
					Name:        "locked",
					Description: `Displays if the deployment is locked or not.`,
				},
				resource.Attribute{
					Name:        "server_tag_scope",
					Description: `Displays what the scope of tags are in the deployment. Options are "deployment" or "account."`,
				},
				resource.Attribute{
					Name:        "href",
					Description: `Href of the deployment.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "rightscale_image",
			Category:         "Data Sources",
			ShortDescription: `Defines an image datasource to operate against.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "cloud_href",
					Description: `(Required) The Href of the cloud with the image you want.`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) block supports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Image name as displayed in cm platform. Pattern match.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Image description as displayed in cm platform. Pattern match.`,
				},
				resource.Attribute{
					Name:        "image_type",
					Description: `(Optional) Image type as referenced in cm platform. This will be either "machine", "machine_azure", "ramdisk" or "kernel".`,
				},
				resource.Attribute{
					Name:        "os_platform",
					Description: `(Optional) Image OS platform as referenced in cm platform. This will either be "windows" or "linux."`,
				},
				resource.Attribute{
					Name:        "cpu_architecture",
					Description: `(Optional) Image CPU architecture as referenced in cm platform. Generally "x64_64", etc. Pattern match. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "visibility",
					Description: `Image visibility as displayed in cm platform.`,
				},
				resource.Attribute{
					Name:        "resource_uid",
					Description: `Image unique resource identifier as displayed in cm platform.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Image name as displayed in cm platform.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Image description as displayed in cm platform.`,
				},
				resource.Attribute{
					Name:        "cpu_architecture",
					Description: `Image CPU architecture as referenced in cm platform.`,
				},
				resource.Attribute{
					Name:        "os_platform",
					Description: `Image OS platform as referenced in cm platform.`,
				},
				resource.Attribute{
					Name:        "root_device_storage",
					Description: `Image root device storage as reported in cm platform. Eg "volume" vs "instance", etc.`,
				},
				resource.Attribute{
					Name:        "image_type",
					Description: `Image type as referenced in cm platform.`,
				},
				resource.Attribute{
					Name:        "virtualization_type",
					Description: `Image virtualization type as referenced in cm platform. Eg "hvm" etc.`,
				},
				resource.Attribute{
					Name:        "links",
					Description: `Hrefs of related API resources.`,
				},
				resource.Attribute{
					Name:        "href",
					Description: `Href of the image.`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "visibility",
					Description: `Image visibility as displayed in cm platform.`,
				},
				resource.Attribute{
					Name:        "resource_uid",
					Description: `Image unique resource identifier as displayed in cm platform.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Image name as displayed in cm platform.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Image description as displayed in cm platform.`,
				},
				resource.Attribute{
					Name:        "cpu_architecture",
					Description: `Image CPU architecture as referenced in cm platform.`,
				},
				resource.Attribute{
					Name:        "os_platform",
					Description: `Image OS platform as referenced in cm platform.`,
				},
				resource.Attribute{
					Name:        "root_device_storage",
					Description: `Image root device storage as reported in cm platform. Eg "volume" vs "instance", etc.`,
				},
				resource.Attribute{
					Name:        "image_type",
					Description: `Image type as referenced in cm platform.`,
				},
				resource.Attribute{
					Name:        "virtualization_type",
					Description: `Image virtualization type as referenced in cm platform. Eg "hvm" etc.`,
				},
				resource.Attribute{
					Name:        "links",
					Description: `Hrefs of related API resources.`,
				},
				resource.Attribute{
					Name:        "href",
					Description: `Href of the image.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "rightscale_instance",
			Category:         "Data Sources",
			ShortDescription: `Defines an instance datasource to operate against.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `The name of the instance`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the instance (e.g.: operational, terminated, stranded, ...)`,
				},
				resource.Attribute{
					Name:        "os_platform",
					Description: `The OS platform of the instance. One of "linux" or "windows"`,
				},
				resource.Attribute{
					Name:        "parent_href",
					Description: `The Href of instance server or server array parent resource.`,
				},
				resource.Attribute{
					Name:        "server_template_href",
					Description: `The Href of the instance server template resource`,
				},
				resource.Attribute{
					Name:        "public_dns_name",
					Description: `The public DNS name of the instance`,
				},
				resource.Attribute{
					Name:        "private_dns_name",
					Description: `The private DNS name of the instance`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `The public IP of the instance`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `The private IP of the instance`,
				},
				resource.Attribute{
					Name:        "resource_uid",
					Description: `The resource_uid of the instance. If this filter option is set, additional retry logic will be enabled to wait up to 5 minutes for cloud resources to be polled and populated for use.`,
				},
				resource.Attribute{
					Name:        "deployment_href",
					Description: `The href of the [deployment](http://docs.rightscale.com/cm/dashboard/manage/deployments/) that contains the instance (e.g. /api/deployments/594684003)`,
				},
				resource.Attribute{
					Name:        "placement_group_href",
					Description: `The href of the [placement_group](http://docs.rightscale.com/cm/dashboard/clouds/aws/ec2_placement_groups.html) that contains the instance (e.g. /api/placement_groups/512SV3FUJA7OO)`,
				},
				resource.Attribute{
					Name:        "datacenter_href",
					Description: `The href of the [datacenter](http://docs.rightscale.com/cm/dashboard/clouds/generic/datacenter_zones_concepts.html) that holds the instance (e.g. /api/clouds/6/datacenters/6IHONC8ANOUHI) ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "associate_public_ip_address",
					Description: `Indicates if the instance will get a Public IP address`,
				},
				resource.Attribute{
					Name:        "cloud_href",
					Description: `The cloud_href the instance belongs to (mutually exclusive with server_array_href)`,
				},
				resource.Attribute{
					Name:        "server_array_href",
					Description: `The server_array_href the instance belongs to (mutually exclusive with cloud_href)`,
				},
				resource.Attribute{
					Name:        "cloud_specific_attributes",
					Description: `Attributes specific to the cloud the instance belongs to that have no specific rightscale abstraction. This block includes:`,
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
					Description: `Email of service account for instance. Scope will default to cloud-platform. Supported by GCE cloud only.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the instance`,
				},
				resource.Attribute{
					Name:        "pricing_type",
					Description: `Pricing type of the instance (e.g. fixed, spot)`,
				},
				resource.Attribute{
					Name:        "resource_uid",
					Description: `The resource_uid of the instance (e.g. e0bf62bc-4e35-11e8-9f1f-0242ac110003)`,
				},
				resource.Attribute{
					Name:        "links",
					Description: `Hrefs of related API resources`,
				},
				resource.Attribute{
					Name:        "locked",
					Description: `Whether instance is locked, a locked instance cannot be terminated or deleted`,
				},
				resource.Attribute{
					Name:        "private_ip_addresses",
					Description: `List of private IP addresses of the instance`,
				},
				resource.Attribute{
					Name:        "public_ip_addresses",
					Description: `List of public IP addresses of the instance`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The instance state (e.g. operational, terminated, stranded, ...)`,
				},
				resource.Attribute{
					Name:        "created_ at",
					Description: `Time of creation of the instance`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `Last update of the instance`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The instance ID (e.g. rs_cm:/api/clouds/1/instances/63NFHKF8B7RP4)`,
				},
				resource.Attribute{
					Name:        "href",
					Description: `Href of the instance (e.g. /api/clouds/1/instances/63NFHKF8B7RP4)`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "associate_public_ip_address",
					Description: `Indicates if the instance will get a Public IP address`,
				},
				resource.Attribute{
					Name:        "cloud_href",
					Description: `The cloud_href the instance belongs to (mutually exclusive with server_array_href)`,
				},
				resource.Attribute{
					Name:        "server_array_href",
					Description: `The server_array_href the instance belongs to (mutually exclusive with cloud_href)`,
				},
				resource.Attribute{
					Name:        "cloud_specific_attributes",
					Description: `Attributes specific to the cloud the instance belongs to that have no specific rightscale abstraction. This block includes:`,
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
					Description: `Email of service account for instance. Scope will default to cloud-platform. Supported by GCE cloud only.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the instance`,
				},
				resource.Attribute{
					Name:        "pricing_type",
					Description: `Pricing type of the instance (e.g. fixed, spot)`,
				},
				resource.Attribute{
					Name:        "resource_uid",
					Description: `The resource_uid of the instance (e.g. e0bf62bc-4e35-11e8-9f1f-0242ac110003)`,
				},
				resource.Attribute{
					Name:        "links",
					Description: `Hrefs of related API resources`,
				},
				resource.Attribute{
					Name:        "locked",
					Description: `Whether instance is locked, a locked instance cannot be terminated or deleted`,
				},
				resource.Attribute{
					Name:        "private_ip_addresses",
					Description: `List of private IP addresses of the instance`,
				},
				resource.Attribute{
					Name:        "public_ip_addresses",
					Description: `List of public IP addresses of the instance`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The instance state (e.g. operational, terminated, stranded, ...)`,
				},
				resource.Attribute{
					Name:        "created_ at",
					Description: `Time of creation of the instance`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `Last update of the instance`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The instance ID (e.g. rs_cm:/api/clouds/1/instances/63NFHKF8B7RP4)`,
				},
				resource.Attribute{
					Name:        "href",
					Description: `Href of the instance (e.g. /api/clouds/1/instances/63NFHKF8B7RP4)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "rightscale_instance_type",
			Category:         "Data Sources",
			ShortDescription: `Defines a instance_type datasource to operate against.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "cloud_href",
					Description: `(Required) The Href of the cloud with the instance type you want.`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) block supports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Instance type name as displayed in cm platform. Pattern match.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Instance type description as displayed in cm platform. Pattern match.`,
				},
				resource.Attribute{
					Name:        "cpu_architecture",
					Description: `(Optional) Instance type CPU architecture as referenced in cm platform. Generally "x64_64", etc. Pattern match. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "resource_uid",
					Description: `Instance type unique resource identifier as displayed in cm platform.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Instance type name as displayed in cm platform.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Instance type description as displayed in cm platform.`,
				},
				resource.Attribute{
					Name:        "cpu_architecture",
					Description: `Instance type CPU architecture as displayed in cm platform.`,
				},
				resource.Attribute{
					Name:        "cpu_count",
					Description: `Instance type CPU count as displayed in cm platform.`,
				},
				resource.Attribute{
					Name:        "cpu_speed",
					Description: `Instance type CPU speed as displayed in cm platform.`,
				},
				resource.Attribute{
					Name:        "memory",
					Description: `Instance type memory as displayed in cm platform.`,
				},
				resource.Attribute{
					Name:        "links",
					Description: `Hrefs of related API resources.`,
				},
				resource.Attribute{
					Name:        "href",
					Description: `Href of the instance type.`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "resource_uid",
					Description: `Instance type unique resource identifier as displayed in cm platform.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Instance type name as displayed in cm platform.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Instance type description as displayed in cm platform.`,
				},
				resource.Attribute{
					Name:        "cpu_architecture",
					Description: `Instance type CPU architecture as displayed in cm platform.`,
				},
				resource.Attribute{
					Name:        "cpu_count",
					Description: `Instance type CPU count as displayed in cm platform.`,
				},
				resource.Attribute{
					Name:        "cpu_speed",
					Description: `Instance type CPU speed as displayed in cm platform.`,
				},
				resource.Attribute{
					Name:        "memory",
					Description: `Instance type memory as displayed in cm platform.`,
				},
				resource.Attribute{
					Name:        "links",
					Description: `Hrefs of related API resources.`,
				},
				resource.Attribute{
					Name:        "href",
					Description: `Href of the instance type.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "rightscale_multi_cloud_image",
			Category:         "Data Sources",
			ShortDescription: `Defines a multi cloud image datasource to operate against.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `The name of the multi cloud image`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the multi cloud image`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `The revision of multi-cloud image, use 0 to match latest non-committed version ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the multi cloud image`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the multi cloud image`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `The revision of multi-cloud image, use 0 to match latest non-committed version`,
				},
				resource.Attribute{
					Name:        "links",
					Description: `Hrefs of related API resources`,
				},
				resource.Attribute{
					Name:        "href",
					Description: `Href of the multi-cloud image`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `The name of the multi cloud image`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the multi cloud image`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `The revision of multi-cloud image, use 0 to match latest non-committed version`,
				},
				resource.Attribute{
					Name:        "links",
					Description: `Hrefs of related API resources`,
				},
				resource.Attribute{
					Name:        "href",
					Description: `Href of the multi-cloud image`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "rightscale_network",
			Category:         "Data Sources",
			ShortDescription: `Defines a network datasource to operate against.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Network name. Pattern match.`,
				},
				resource.Attribute{
					Name:        "cloud_href",
					Description: `(Optional) Cloud Href of network.`,
				},
				resource.Attribute{
					Name:        "deployment_href",
					Description: `(Optional) Deployment href associated with network.`,
				},
				resource.Attribute{
					Name:        "cidr_block",
					Description: `(Optional) CIDR notation block of network.`,
				},
				resource.Attribute{
					Name:        "resource_uid",
					Description: `(Optional) The resource_uid of the network. If this filter option is set, additional retry logic will be enabled to wait up to 5 minutes for cloud resources to be polled and populated for use. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the network.`,
				},
				resource.Attribute{
					Name:        "resource_uid",
					Description: `Network resource_uid as reported by cm platform.`,
				},
				resource.Attribute{
					Name:        "cidr_block",
					Description: `Network CIDR notation block of network.`,
				},
				resource.Attribute{
					Name:        "instance_tenancy",
					Description: `Tenancy of instances on network.`,
				},
				resource.Attribute{
					Name:        "is_default",
					Description: `Reports if network is 'default' for a given cloud.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the network.`,
				},
				resource.Attribute{
					Name:        "links",
					Description: `Hrefs of related API resources.`,
				},
				resource.Attribute{
					Name:        "href",
					Description: `Href of the network.`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `Name of the network.`,
				},
				resource.Attribute{
					Name:        "resource_uid",
					Description: `Network resource_uid as reported by cm platform.`,
				},
				resource.Attribute{
					Name:        "cidr_block",
					Description: `Network CIDR notation block of network.`,
				},
				resource.Attribute{
					Name:        "instance_tenancy",
					Description: `Tenancy of instances on network.`,
				},
				resource.Attribute{
					Name:        "is_default",
					Description: `Reports if network is 'default' for a given cloud.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the network.`,
				},
				resource.Attribute{
					Name:        "links",
					Description: `Hrefs of related API resources.`,
				},
				resource.Attribute{
					Name:        "href",
					Description: `Href of the network.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "rightscale_network_gateway",
			Category:         "Data Sources",
			ShortDescription: `Defines a network gateway datasource to operate against.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Network gateway name. Pattern match.`,
				},
				resource.Attribute{
					Name:        "cloud_href",
					Description: `(Optional) Cloud href of network gateway.`,
				},
				resource.Attribute{
					Name:        "network_href",
					Description: `(Optional) Network href that network gateway is attached to. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the network gateway.`,
				},
				resource.Attribute{
					Name:        "resource_uid",
					Description: `Network gateway resource_uid from cloud.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type of network gateway. Options are "internet" or "vpc."`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `State of the network gateway. ("available" means attached to a network)`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the network gateway.`,
				},
				resource.Attribute{
					Name:        "links",
					Description: `Hrefs of related API resources.`,
				},
				resource.Attribute{
					Name:        "href",
					Description: `Href of the network gateway.`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `Name of the network gateway.`,
				},
				resource.Attribute{
					Name:        "resource_uid",
					Description: `Network gateway resource_uid from cloud.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type of network gateway. Options are "internet" or "vpc."`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `State of the network gateway. ("available" means attached to a network)`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the network gateway.`,
				},
				resource.Attribute{
					Name:        "links",
					Description: `Hrefs of related API resources.`,
				},
				resource.Attribute{
					Name:        "href",
					Description: `Href of the network gateway.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "rightscale_route_table",
			Category:         "Data Sources",
			ShortDescription: `Defines a route_table datasource to operate against.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Route table name. Pattern match.`,
				},
				resource.Attribute{
					Name:        "cloud_href",
					Description: `(Optional) Cloud href of route table.`,
				},
				resource.Attribute{
					Name:        "network_href",
					Description: `(Optional) Network href that owns the route table. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the route table.`,
				},
				resource.Attribute{
					Name:        "resource_uid",
					Description: `Cloud resource_uid.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the route table.`,
				},
				resource.Attribute{
					Name:        "routes",
					Description: `Associated routes.`,
				},
				resource.Attribute{
					Name:        "links",
					Description: `Hrefs of related API resources.`,
				},
				resource.Attribute{
					Name:        "href",
					Description: `Href of the route table.`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `Name of the route table.`,
				},
				resource.Attribute{
					Name:        "resource_uid",
					Description: `Cloud resource_uid.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the route table.`,
				},
				resource.Attribute{
					Name:        "routes",
					Description: `Associated routes.`,
				},
				resource.Attribute{
					Name:        "links",
					Description: `Hrefs of related API resources.`,
				},
				resource.Attribute{
					Name:        "href",
					Description: `Href of the route table.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "rightscale_security_group",
			Category:         "Data Sources",
			ShortDescription: `Defines a security_group datasource to operate against.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Security group name. Pattern match.`,
				},
				resource.Attribute{
					Name:        "resource_uid",
					Description: `(Optional) Cloud resource uid for security group. If this filter option is set, additional retry logic will be enabled to wait up to 5 minutes for cloud resources to be polled and populated for use.`,
				},
				resource.Attribute{
					Name:        "network_href",
					Description: `(Optional) Network href that security group is created in.`,
				},
				resource.Attribute{
					Name:        "deployment_href",
					Description: `(Optional) Href of the deployment that owns the security group. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the security group.`,
				},
				resource.Attribute{
					Name:        "resource_uid",
					Description: `The cloud resource_uid of the security group.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the security group.`,
				},
				resource.Attribute{
					Name:        "links",
					Description: `Hrefs of related API resources.`,
				},
				resource.Attribute{
					Name:        "href",
					Description: `Href of the security group.`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `Name of the security group.`,
				},
				resource.Attribute{
					Name:        "resource_uid",
					Description: `The cloud resource_uid of the security group.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the security group.`,
				},
				resource.Attribute{
					Name:        "links",
					Description: `Hrefs of related API resources.`,
				},
				resource.Attribute{
					Name:        "href",
					Description: `Href of the security group.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "rightscale_server",
			Category:         "Data Sources",
			ShortDescription: `Defines an server datasource to operate against.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "deployment_href",
					Description: `(Optional) The href of the deployment`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the server`,
				},
				resource.Attribute{
					Name:        "cloud_href",
					Description: `(Optional) The Href of the cloud with the ssh key you want ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A description of the server`,
				},
				resource.Attribute{
					Name:        "instance",
					Description: `See [rightscale_instance](https://github.com/terraform-providers/terraform-provider-rightscale/blob/master/website/docs/r/cm_server.markdown)`,
				},
				resource.Attribute{
					Name:        "optimized",
					Description: `A flag indicating whether instances of this server should be optimized for high-performance volumes`,
				},
				resource.Attribute{
					Name:        "links",
					Description: `Hrefs of related API resources`,
				},
				resource.Attribute{
					Name:        "href",
					Description: `Href of the server`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "description",
					Description: `A description of the server`,
				},
				resource.Attribute{
					Name:        "instance",
					Description: `See [rightscale_instance](https://github.com/terraform-providers/terraform-provider-rightscale/blob/master/website/docs/r/cm_server.markdown)`,
				},
				resource.Attribute{
					Name:        "optimized",
					Description: `A flag indicating whether instances of this server should be optimized for high-performance volumes`,
				},
				resource.Attribute{
					Name:        "links",
					Description: `Hrefs of related API resources`,
				},
				resource.Attribute{
					Name:        "href",
					Description: `Href of the server`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "rightscale_server_template",
			Category:         "Data Sources",
			ShortDescription: `Defines a server template datasource to operate against.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `The name of the server template`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `The revision of the server template, use 0 to match latest non-committed version`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the server template`,
				},
				resource.Attribute{
					Name:        "lineage",
					Description: `The lineage of the server template`,
				},
				resource.Attribute{
					Name:        "multi_cloud_image_href",
					Description: `The href of the server template multicloud image resource ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the server template`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the server template`,
				},
				resource.Attribute{
					Name:        "lineage",
					Description: `The lineage of the server template`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `The revision of the server template, use 0 to match latest non-committed version`,
				},
				resource.Attribute{
					Name:        "links",
					Description: `Hrefs of related API resources`,
				},
				resource.Attribute{
					Name:        "href",
					Description: `Href of the server template`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `The name of the server template`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the server template`,
				},
				resource.Attribute{
					Name:        "lineage",
					Description: `The lineage of the server template`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `The revision of the server template, use 0 to match latest non-committed version`,
				},
				resource.Attribute{
					Name:        "links",
					Description: `Hrefs of related API resources`,
				},
				resource.Attribute{
					Name:        "href",
					Description: `Href of the server template`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "rightscale_ssh_key",
			Category:         "Data Sources",
			ShortDescription: `Defines an ssh key datasource to operate against.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "cloud_href",
					Description: `(Required) The Href of the cloud with the ssh key you want.`,
				},
				resource.Attribute{
					Name:        "view",
					Description: `(Optional) Set this to 'sensitive' to request the api return 'sensitive' information (in this case the private key material) with the request. Assumes rs account privs sufficient to do this operation. The ` + "`" + `filter` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) SSH key name. Pattern match.`,
				},
				resource.Attribute{
					Name:        "resource_uid",
					Description: `(Optional) resource_uid of the SSH key. If this filter option is set, additional retry logic will be enabled to wait up to 5 minutes for cloud resources to be polled and populated for use. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Official cloud name as displayed in cm platform.`,
				},
				resource.Attribute{
					Name:        "resource_uid",
					Description: `resource_uid of the SSH key.`,
				},
				resource.Attribute{
					Name:        "links",
					Description: `Hrefs of related API resources.`,
				},
				resource.Attribute{
					Name:        "material",
					Description: `(Contextual) Available only if 'sensitive' view is set.`,
				},
				resource.Attribute{
					Name:        "href",
					Description: `Href of the SSH key.`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `Official cloud name as displayed in cm platform.`,
				},
				resource.Attribute{
					Name:        "resource_uid",
					Description: `resource_uid of the SSH key.`,
				},
				resource.Attribute{
					Name:        "links",
					Description: `Hrefs of related API resources.`,
				},
				resource.Attribute{
					Name:        "material",
					Description: `(Contextual) Available only if 'sensitive' view is set.`,
				},
				resource.Attribute{
					Name:        "href",
					Description: `Href of the SSH key.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "rightscale_subnet",
			Category:         "Data Sources",
			ShortDescription: `Defines a subnet datasource to operate against.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Subnet name. Pattern match.`,
				},
				resource.Attribute{
					Name:        "network_href",
					Description: `(Optional) Network href the the subnet exists in.`,
				},
				resource.Attribute{
					Name:        "resource_uid",
					Description: `(Optional) The resource_uid of the subnet. If this filter option is set, additional retry logic will be enabled to wait up to 5 minutes for cloud resources to be polled and populated for use.`,
				},
				resource.Attribute{
					Name:        "datacenter_href",
					Description: `(Optional) Href of the subnet datacenter resource.`,
				},
				resource.Attribute{
					Name:        "instance_href",
					Description: `(Optional) Href of instance resource attached to subnet.`,
				},
				resource.Attribute{
					Name:        "visibility",
					Description: `(Optional) Visibility of the subnet to filter by (private, shared, etc). ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the subnet.`,
				},
				resource.Attribute{
					Name:        "resource_uid",
					Description: `Subnet resource_uid.`,
				},
				resource.Attribute{
					Name:        "cidr_block",
					Description: `Subnet allocation range in CIDR notation.`,
				},
				resource.Attribute{
					Name:        "is_default",
					Description: `Reports if subnet is 'default' for a given subnet.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the subnet.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `Indicates whether subnet is pending, available etc.`,
				},
				resource.Attribute{
					Name:        "visibility",
					Description: `Visibility of the subnet.`,
				},
				resource.Attribute{
					Name:        "links",
					Description: `Hrefs of related API resources.`,
				},
				resource.Attribute{
					Name:        "href",
					Description: `Href of the subnet.`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `Name of the subnet.`,
				},
				resource.Attribute{
					Name:        "resource_uid",
					Description: `Subnet resource_uid.`,
				},
				resource.Attribute{
					Name:        "cidr_block",
					Description: `Subnet allocation range in CIDR notation.`,
				},
				resource.Attribute{
					Name:        "is_default",
					Description: `Reports if subnet is 'default' for a given subnet.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the subnet.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `Indicates whether subnet is pending, available etc.`,
				},
				resource.Attribute{
					Name:        "visibility",
					Description: `Visibility of the subnet.`,
				},
				resource.Attribute{
					Name:        "links",
					Description: `Hrefs of related API resources.`,
				},
				resource.Attribute{
					Name:        "href",
					Description: `Href of the subnet.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "rightscale_volume",
			Category:         "Data Sources",
			ShortDescription: `Defines a volume datasource to operate against.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `The name of the volume`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the volume`,
				},
				resource.Attribute{
					Name:        "resource_uid",
					Description: `The resource_uid of the volume. If this filter option is set, additional retry logic will be enabled to wait up to 5 minutes for cloud resources to be polled and populated for use.`,
				},
				resource.Attribute{
					Name:        "deployment_href",
					Description: `The href of the [deployment](http://docs.rightscale.com/cm/dashboard/manage/deployments/) that contains the volume (e.g. /api/deployments/594684003)`,
				},
				resource.Attribute{
					Name:        "datacenter_href",
					Description: `The href of the [datacenter](http://docs.rightscale.com/cm/dashboard/clouds/generic/datacenter_zones_concepts.html) that holds the volume (e.g. /api/clouds/6/datacenters/6IHONC8ANOUHI)`,
				},
				resource.Attribute{
					Name:        "parent_volume_snapshot_href",
					Description: `The href of snapshot the volume was created of ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the volume`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the volume`,
				},
				resource.Attribute{
					Name:        "resource_uid",
					Description: `The resource_uid of the volume (e.g. vol-045e33fd28a746c45)`,
				},
				resource.Attribute{
					Name:        "links",
					Description: `Hrefs of related API resources`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The volume size (in GB)`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The volume Status (e.g. available, in-use, ...)`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `Last update of the volume`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The volume ID (e.g. rs_cm:/api/clouds/1/volumes/63NFHKF8B7RP4)`,
				},
				resource.Attribute{
					Name:        "href",
					Description: `Href of the volume (e.g. /api/clouds/1/volumes/63NFHKF8B7RP4)`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `The name of the volume`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the volume`,
				},
				resource.Attribute{
					Name:        "resource_uid",
					Description: `The resource_uid of the volume (e.g. vol-045e33fd28a746c45)`,
				},
				resource.Attribute{
					Name:        "links",
					Description: `Hrefs of related API resources`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The volume size (in GB)`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The volume Status (e.g. available, in-use, ...)`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `Last update of the volume`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The volume ID (e.g. rs_cm:/api/clouds/1/volumes/63NFHKF8B7RP4)`,
				},
				resource.Attribute{
					Name:        "href",
					Description: `Href of the volume (e.g. /api/clouds/1/volumes/63NFHKF8B7RP4)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "rightscale_volume_snapshot",
			Category:         "Data Sources",
			ShortDescription: `Defines a volume snapshot datasource to operate against.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `The name of the volume snapshot`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the volume snapshot`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the volume snapshot (e.g.: available, pending, ...)`,
				},
				resource.Attribute{
					Name:        "parent_volume_href",
					Description: `The Href of the parent resource`,
				},
				resource.Attribute{
					Name:        "resource_uid",
					Description: `The resource_uid of the volume snapshot. If this filter option is set, additional retry logic will be enabled to wait up to 5 minutes for cloud resources to be polled and populated for use.`,
				},
				resource.Attribute{
					Name:        "deployment_href",
					Description: `The href of the [deployment](http://docs.rightscale.com/cm/dashboard/manage/deployments/) that contains the volume snapshot (e.g. /api/deployments/594684003) ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the volume snapshot`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the volume snapshot`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The size of the volume snapshot`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the volume snapshot (e.g.: available, pending, ...)`,
				},
				resource.Attribute{
					Name:        "resource_uid",
					Description: `The resource_uid of the volume snapshot (e.g. snap-08287ed6c8bce9ab4)`,
				},
				resource.Attribute{
					Name:        "links",
					Description: `Hrefs of related API resources`,
				},
				resource.Attribute{
					Name:        "created_ at",
					Description: `Time of creation of the volume snapshot`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `Last update of the volume snapshot`,
				},
				resource.Attribute{
					Name:        "href",
					Description: `Href of the volume snapshot (e.g. /api/clouds/1/volume_snapshots/4VODPN6TQ60RC)`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "description",
					Description: `The description of the volume snapshot`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the volume snapshot`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The size of the volume snapshot`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the volume snapshot (e.g.: available, pending, ...)`,
				},
				resource.Attribute{
					Name:        "resource_uid",
					Description: `The resource_uid of the volume snapshot (e.g. snap-08287ed6c8bce9ab4)`,
				},
				resource.Attribute{
					Name:        "links",
					Description: `Hrefs of related API resources`,
				},
				resource.Attribute{
					Name:        "created_ at",
					Description: `Time of creation of the volume snapshot`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `Last update of the volume snapshot`,
				},
				resource.Attribute{
					Name:        "href",
					Description: `Href of the volume snapshot (e.g. /api/clouds/1/volume_snapshots/4VODPN6TQ60RC)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "rightscale_volume_type",
			Category:         "Data Sources",
			ShortDescription: `Defines a volume type datasource to operate against.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `The name of the volume type as reported by the rightscale platform`,
				},
				resource.Attribute{
					Name:        "resource_uid",
					Description: `The resource_uid of the volume_type. If this filter option is set, additional retry logic will be enabled to wait up to 5 minutes for cloud resources to be polled and populated for use. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the volume type.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the volume type.`,
				},
				resource.Attribute{
					Name:        "resource_uid",
					Description: `The resource_uid of the volume type. (e.g. gp2)`,
				},
				resource.Attribute{
					Name:        "links",
					Description: `Hrefs of related API resources`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The volume size (in GB) if applicable (depends on cloud)`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Creation date of the volume type`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `Last update of the volume type`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The volume type ID (e.g. rs_cm:/api/clouds/1/volume_types/B37A8VOCJIODH)`,
				},
				resource.Attribute{
					Name:        "href",
					Description: `Href of the volume type (e.g. /api/clouds/1/volume_types/B37A8VOCJIODH)`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `The name of the volume type.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the volume type.`,
				},
				resource.Attribute{
					Name:        "resource_uid",
					Description: `The resource_uid of the volume type. (e.g. gp2)`,
				},
				resource.Attribute{
					Name:        "links",
					Description: `Hrefs of related API resources`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The volume size (in GB) if applicable (depends on cloud)`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Creation date of the volume type`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `Last update of the volume type`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The volume type ID (e.g. rs_cm:/api/clouds/1/volume_types/B37A8VOCJIODH)`,
				},
				resource.Attribute{
					Name:        "href",
					Description: `Href of the volume type (e.g. /api/clouds/1/volume_types/B37A8VOCJIODH)`,
				},
			},
		},
	}

	dataSourcesMap = map[string]Resource{

		"rightscale_cloud":             0,
		"rightscale_credential":        1,
		"rightscale_datacenter":        2,
		"rightscale_deployment":        3,
		"rightscale_image":             4,
		"rightscale_instance":          5,
		"rightscale_instance_type":     6,
		"rightscale_multi_cloud_image": 7,
		"rightscale_network":           8,
		"rightscale_network_gateway":   9,
		"rightscale_route_table":       10,
		"rightscale_security_group":    11,
		"rightscale_server":            12,
		"rightscale_server_template":   13,
		"rightscale_ssh_key":           14,
		"rightscale_subnet":            15,
		"rightscale_volume":            16,
		"rightscale_volume_snapshot":   17,
		"rightscale_volume_type":       18,
	}
)

func GetDataSource(r string) (*resouce.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs]
}
