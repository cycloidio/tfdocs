package rediscloud

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "rediscloud_rediscloud_active_active_subscription",
			Category:         "Resources",
			ShortDescription: `Subscription resource in the Terraform provider Redis Cloud.`,
			Description: `

Creates an Active-Active Subscription within your Redis Enterprise Cloud Account.
This resource is responsible for creating and managing subscriptions.

~> **Note:** The creation_plan block allows the API server to create a well-optimised infrastructure for your databases in the cluster.
The attributes inside the block are used by the provider to create initial 
databases. Those databases will be deleted after provisioning a new 
subscription, then the databases defined as separate resources will be attached to 
the subscription. The creation_plan block can ONLY be used for provisioning new 
subscriptions, the block will be ignored if you make any further changes or try importing the resource (e.g. ` + "`" + `terraform import` + "`" + ` ...).  

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A meaningful name to identify the subscription`,
				},
				resource.Attribute{
					Name:        "payment_method_id",
					Description: `(Optional) A valid payment method pre-defined in the current account. This value is __Optional__ for AWS/GCP Marketplace accounts, but __Required__ for all other account types`,
				},
				resource.Attribute{
					Name:        "cloud_provider",
					Description: `(Optional) The cloud provider to use with the subscription, (either ` + "`" + `AWS` + "`" + ` or ` + "`" + `GCP` + "`" + `). Default: ‘AWS’`,
				},
				resource.Attribute{
					Name:        "creation_plan",
					Description: `(Required) A creation plan object, documented below The ` + "`" + `creation_plan` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "memory_limit_in_gb",
					Description: `(Required) Maximum memory usage that will be used for your largest planned database, including replication and other overhead`,
				},
				resource.Attribute{
					Name:        "quantity",
					Description: `(Required) The planned number of databases in the subscription.`,
				},
				resource.Attribute{
					Name:        "support_oss_cluster_api",
					Description: `(Optional) Support Redis open-source (OSS) Cluster API. Default: ‘false’ The creation_plan ` + "`" + `region` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Required) Deployment region as defined by cloud provider`,
				},
				resource.Attribute{
					Name:        "networking_deployment_cidr",
					Description: `(Required) Deployment CIDR mask. The total number of bits must be 24 (x.x.x.x/24)`,
				},
				resource.Attribute{
					Name:        "write_operations_per_second",
					Description: `(Required) Throughput measurement for an active-active subscription`,
				},
				resource.Attribute{
					Name:        "read_operations_per_second",
					Description: `(Required) Throughput measurement for an active-active subscription ~>`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Defaults to 30 mins) Used when creating the subscription`,
				},
				resource.Attribute{
					Name:        "update",
					Description: `(Defaults to 30 mins) Used when updating the subscription`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Defaults to 10 mins) Used when destroying the subscription ## Import ` + "`" + `rediscloud_active_active_subscription` + "`" + ` can be imported using the ID of the subscription, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import rediscloud_active_active_subscription.subscription-resource 12345678 ` + "`" + `` + "`" + `` + "`" + ` ~>`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "rediscloud_rediscloud_active_active_subscription_database",
			Category:         "Resources",
			ShortDescription: `Database resource for Active-Active Subscriptions in the Terraform provider Redis Cloud.`,
			Description: `

Creates a Database within a specified Active-Active Subscription in your Redis Enterprise Cloud Account.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A meaningful name to identify the database`,
				},
				resource.Attribute{
					Name:        "memory_limit_in_gb",
					Description: `(Required) Maximum memory usage for this specific database, including replication and other overhead`,
				},
				resource.Attribute{
					Name:        "support_oss_cluster_api",
					Description: `(Optional) Support Redis open-source (OSS) Cluster API. Default: ‘false’`,
				},
				resource.Attribute{
					Name:        "external_endpoint_for_oss_cluster_api",
					Description: `(Optional) Should use the external endpoint for open-source (OSS) Cluster API. Can only be enabled if OSS Cluster API support is enabled. Default: 'false'`,
				},
				resource.Attribute{
					Name:        "enable_tls",
					Description: `(Optional) Use TLS for authentication. Default: ‘false’`,
				},
				resource.Attribute{
					Name:        "client_ssl_certificate",
					Description: `(Optional) SSL certificate to authenticate user connections.`,
				},
				resource.Attribute{
					Name:        "data_eviction",
					Description: `(Optional) The data items eviction policy (either: 'allkeys-lru', 'allkeys-lfu', 'allkeys-random', 'volatile-lru', 'volatile-lfu', 'volatile-random', 'volatile-ttl' or 'noeviction'. Default: 'volatile-lru')`,
				},
				resource.Attribute{
					Name:        "global_data_persistence",
					Description: `(Optional) Global rate of database data persistence (in persistent storage) of regions that dont override global settings. Default: 'none'`,
				},
				resource.Attribute{
					Name:        "global_password",
					Description: `(Optional) Password to access the database of regions that dont override global settings. If left empty, the password will be generated automatically`,
				},
				resource.Attribute{
					Name:        "global_alert",
					Description: `(Optional) A block defining Redis database alert of regions that dont override global settings, documented below, can be specified multiple times`,
				},
				resource.Attribute{
					Name:        "global_source_ips",
					Description: `(Optional) List of source IP addresses or subnet masks of regions that dont override global settings. If specified, Redis clients will be able to connect to this database only from within the specified source IP addresses ranges (example: ['192.168.10.0/32', '192.168.12.0/24'])`,
				},
				resource.Attribute{
					Name:        "override_region",
					Description: `(Optional) Override region specific configuration, documented below The ` + "`" + `override_region` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Region name.`,
				},
				resource.Attribute{
					Name:        "override_global_alert",
					Description: `(Optional) A block defining Redis regional instance of an Active-Active database alert, documented below, can be specified multiple times`,
				},
				resource.Attribute{
					Name:        "override_global_password",
					Description: `(Optional) If specified, this regional instance of an Active-Active database password will be used to access the database`,
				},
				resource.Attribute{
					Name:        "override_global_source_ips",
					Description: `(Optional) List of regional instance of an Active-Active database source IP addresses or subnet masks. If specified, Redis clients will be able to connect to this database only from within the specified source IP addresses ranges (example: ['192.168.10.0/32', '192.168.12.0/24'] )`,
				},
				resource.Attribute{
					Name:        "override_global_data_persistence",
					Description: `(Optional) Regional instance of an Active-Active database data persistence rate (in persistent storage) The ` + "`" + `override_global_alert` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Alert name`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) Alert value ### Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/language/resources/syntax#operation-timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Defaults to 30 mins) Used when creating the database`,
				},
				resource.Attribute{
					Name:        "update",
					Description: `(Defaults to 30 mins) Used when updating the database`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Defaults to 10 mins) Used when destroying the database ## Attribute reference`,
				},
				resource.Attribute{
					Name:        "db_id",
					Description: `Identifier of the database created`,
				},
				resource.Attribute{
					Name:        "public_endpoint",
					Description: `A map of which public endpoints can to access the database per region, uses region name as key.`,
				},
				resource.Attribute{
					Name:        "private_endpoint",
					Description: `A map of which private endpoints can to access the database per region, uses region name as key. ## Import ` + "`" + `rediscloud_active_active_subscription_database` + "`" + ` can be imported using the ID of the Active-Active subscription and the ID of the database in the format {subscription ID}/{database ID}, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import rediscloud_active_active_subscription_database.database-resource 123456/12345678 ` + "`" + `` + "`" + `` + "`" + ` Note: Due to constraints in the Redis Cloud API, the import process will not import global attributes or override region attributes. If you wish to use these attributes in your Terraform configuraton, you will need to manually add them to your Terraform configuration and run ` + "`" + `terraform apply` + "`" + ` to update the database.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "rediscloud_rediscloud_active_active_subscription_peering",
			Category:         "Resources",
			ShortDescription: `Active-Active subscription VPC peering resource in the Terraform provider Redis Cloud.`,
			Description: `

Creates an AWS or GCP VPC peering for an existing Redis Enterprise Cloud Active-Active Subscription, allowing access to your subscription databases as if they were on the same network.

For AWS, peering should be accepted by the other side.
For GCP, the opposite peering request should be submitted.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "provider_name",
					Description: `(Optional) The cloud provider to use with the vpc peering, (either ` + "`" + `AWS` + "`" + ` or ` + "`" + `GCP` + "`" + `). Default: ‘AWS’`,
				},
				resource.Attribute{
					Name:        "subscription_id",
					Description: `(Required) A valid Active-Active subscription predefined in the current account`,
				},
				resource.Attribute{
					Name:        "aws_account_id",
					Description: `(Required) AWS account ID that the VPC to be peered lives in`,
				},
				resource.Attribute{
					Name:        "source_region",
					Description: `(Required) Name of the region to create the VPC peering from`,
				},
				resource.Attribute{
					Name:        "destination_region",
					Description: `(Required) Name of the region to create the VPC peering to`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required) Identifier of the VPC to be peered`,
				},
				resource.Attribute{
					Name:        "vpc_cidr",
					Description: `(Required) CIDR range of the VPC to be peered`,
				},
				resource.Attribute{
					Name:        "gcp_project_id",
					Description: `(Required) GCP project ID that the VPC to be peered lives in`,
				},
				resource.Attribute{
					Name:        "gcp_network_name",
					Description: `(Required) The name of the network to be peered ### Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Defaults to 10 mins) Used when creating the peering connection`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Defaults to 10 mins) Used when deleting the peering connection ## Attribute reference`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "rediscloud_rediscloud_active_active_subscription_regions",
			Category:         "Resources",
			ShortDescription: `Regions resource in the Terraform provider Redis Cloud.`,
			Description: `

Manages regions within your Redis Enterprise Cloud Active-Active subscription.
This resource is responsible for creating and managing regions within that subscription.
This allows Redis Enterprise Cloud to efficiently provision your cluster within each defined region in a separate block.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "subscription_id",
					Description: `(Required) ID of the subscription that the regions belong to`,
				},
				resource.Attribute{
					Name:        "delete_regions",
					Description: `(Optional) Flag required to be set when one or more regions is to be deleted, if the flag is not set an error will be thrown`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Required) Cloud networking details, per region, documented below The ` + "`" + `region` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "region_id",
					Description: `(Computed) The ID of the region, as created by the API`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Required) Region name`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Computed) Identifier of the VPC to be peered, set by the API`,
				},
				resource.Attribute{
					Name:        "networking_deployment_cidr",
					Description: `(Required) Deployment CIDR mask. The total number of bits must be 24 (x.x.x.x/24)`,
				},
				resource.Attribute{
					Name:        "recreate_region",
					Description: `(Optional) Protection flag, needs to be set if a region has to be re-created. A region will need to be re-created in the case of a change on the ` + "`" + `networking_deployment_cidr` + "`" + ` field. During re-create, the region will be deleted (so the ` + "`" + `delete_regions` + "`" + ` flag also needs to be set) and then created again. Default: 'false'`,
				},
				resource.Attribute{
					Name:        "database",
					Description: `(Required) A block defining the write and read operations in the region, per database, documented below The ` + "`" + `database` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "database_id",
					Description: `(Required) Database ID belonging to the subscription`,
				},
				resource.Attribute{
					Name:        "database_name",
					Description: `(Required) Database name belonging to the subscription`,
				},
				resource.Attribute{
					Name:        "local_write_operations_per_second",
					Description: `(Required) Local write operations per second for this active-active region`,
				},
				resource.Attribute{
					Name:        "local_read_operations_per_second",
					Description: `(Required) Local read operations per second for this active-active region ### Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Defaults to 60 mins) Used when creating the regions`,
				},
				resource.Attribute{
					Name:        "update",
					Description: `(Defaults to 60 mins) Used when updating the regions`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Defaults to 10 mins) Used when destroying the regions ## Import ` + "`" + `rediscloud_active_active_regions` + "`" + ` can be imported using the ID of the subscription, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import rediscloud_active_active_regions.regions-resource 12345678 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "rediscloud_rediscloud_cloud_account",
			Category:         "Resources",
			ShortDescription: `Cloud Account resource in the Terraform provider Redis Enterprise Cloud.`,
			Description: `

Creates a Cloud Account resource representing the access credentials to a cloud provider account, (` + "`" + `AWS` + "`" + `).
Redis Enterprise Cloud uses these credentials to provision databases within your infrastructure. 

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "access_key_id",
					Description: `(Required) Cloud provider access key.`,
				},
				resource.Attribute{
					Name:        "access_secret_key",
					Description: `(Required) Cloud provider secret key. Note that drift cannot currently be detected for this.`,
				},
				resource.Attribute{
					Name:        "console_username",
					Description: `(Required) Cloud provider management console username. Note that drift cannot currently be detected for this.`,
				},
				resource.Attribute{
					Name:        "console_password",
					Description: `(Required) Cloud provider management console password. Note that drift cannot currently be detected for this.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Display name of the account.`,
				},
				resource.Attribute{
					Name:        "provider_type",
					Description: `(Required) Cloud provider type - either ` + "`" + `AWS` + "`" + ` or ` + "`" + `GCP` + "`" + `. Note that drift cannot currently be detected for this.`,
				},
				resource.Attribute{
					Name:        "sign_in_login_url",
					Description: `(Required) Cloud provider management console login URL. Note that drift cannot currently be detected for this. ### Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Defaults to 5 mins) Used when creating the Cloud Account`,
				},
				resource.Attribute{
					Name:        "update",
					Description: `(Defaults to 5 mins) Used when updating the Cloud Account`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Defaults to 5 mins) Used when destroying the Cloud Account ## Attribute Reference ` + "`" + `status` + "`" + ` is set to the current status of the account - ` + "`" + `draft` + "`" + `, ` + "`" + `pending` + "`" + ` or ` + "`" + `active` + "`" + `. ## Import ` + "`" + `rediscloud_cloud_account` + "`" + ` can be imported using the ID of the Cloud Account, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import rediscloud_cloud_account.example 12345678 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "rediscloud_rediscloud_subscription",
			Category:         "Resources",
			ShortDescription: `Subscription resource in the Terraform provider Redis Cloud.`,
			Description: `

Creates a Subscription within your Redis Enterprise Cloud Account.
This resource is responsible for creating and managing subscriptions.

~> **Note:** The creation_plan block allows the API server to create a well-optimised infrastructure for your databases in the cluster.
The attributes inside the block are used by the provider to create initial 
databases. Those databases will be deleted after provisioning a new 
subscription, then the databases defined as separate resources will be attached to 
the subscription. The creation_plan block can ONLY be used for provisioning new 
subscriptions, the block will be ignored if you make any further changes or try importing the resource (e.g. ` + "`" + `terraform import` + "`" + ` ...).  

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A meaningful name to identify the subscription`,
				},
				resource.Attribute{
					Name:        "payment_method_id",
					Description: `(Optional) A valid payment method pre-defined in the current account. This value is __Optional__ for AWS/GCP Marketplace accounts, but __Required__ for all other account types`,
				},
				resource.Attribute{
					Name:        "memory_storage",
					Description: `(Optional) Memory storage preference: either ‘ram’ or a combination of ‘ram-and-flash’. Default: ‘ram’`,
				},
				resource.Attribute{
					Name:        "allowlist",
					Description: `(Optional) An allowlist object, documented below`,
				},
				resource.Attribute{
					Name:        "cloud_provider",
					Description: `(Required) A cloud provider object, documented below`,
				},
				resource.Attribute{
					Name:        "creation_plan",
					Description: `(Required) A creation plan object, documented below The ` + "`" + `allowlist` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "security_group_ids",
					Description: `(Required) Set of security groups that are allowed to access the databases associated with this subscription`,
				},
				resource.Attribute{
					Name:        "cidrs",
					Description: `(Optional) Set of CIDR ranges that are allowed to access the databases associated with this subscription ~>`,
				},
				resource.Attribute{
					Name:        "provider",
					Description: `(Optional) The cloud provider to use with the subscription, (either ` + "`" + `AWS` + "`" + ` or ` + "`" + `GCP` + "`" + `). Default: ‘AWS’`,
				},
				resource.Attribute{
					Name:        "cloud_account_id",
					Description: `(Optional) Cloud account identifier. Default: Redis Labs internal cloud account (using Cloud Account ID = 1 implies using Redis Labs internal cloud account). Note that a GCP subscription can be created only with Redis Labs internal cloud account`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Required) A region object, documented below The ` + "`" + `creation_plan` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "memory_limit_in_gb",
					Description: `(Required) Maximum memory usage that will be used for your largest planned database.`,
				},
				resource.Attribute{
					Name:        "modules",
					Description: `(Optional) a list of modules that will be used by the databases in this subscription. Not currently compatible with ‘ram-and-flash’ memory storage. Example: ` + "`" + `modules = ["RedisJSON", RedisBloom"]` + "`" + ``,
				},
				resource.Attribute{
					Name:        "support_oss_cluster_api",
					Description: `(Optional) Support Redis open-source (OSS) Cluster API. Default: ‘false’`,
				},
				resource.Attribute{
					Name:        "replication",
					Description: `(Required) Databases replication. Set to ` + "`" + `true` + "`" + ` if any of your databases will use replication`,
				},
				resource.Attribute{
					Name:        "quantity",
					Description: `(Required) The planned number of databases in the subscription`,
				},
				resource.Attribute{
					Name:        "throughput_measurement_by",
					Description: `(Required) Throughput measurement method that will be used by your databases, (either ‘number-of-shards’ or ‘operations-per-second’)`,
				},
				resource.Attribute{
					Name:        "throughput_measurement_value",
					Description: `(Required) Throughput value that will be used by your databases (as applies to selected measurement method). The value needs to be the maximum throughput measurement value defined in one of your databases`,
				},
				resource.Attribute{
					Name:        "average_item_size_in_bytes",
					Description: `(Optional) Relevant only to ram-and-flash clusters Estimated average size (measured in bytes) of the items stored in the database. The value needs to be the maximum average item size defined in one of your databases. Default: 1000 ~>`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Required) Deployment region as defined by cloud provider`,
				},
				resource.Attribute{
					Name:        "multiple_availability_zones",
					Description: `(Optional) Support deployment on multiple availability zones within the selected region. Default: ‘false’`,
				},
				resource.Attribute{
					Name:        "networking_deployment_cidr",
					Description: `(Required) Deployment CIDR mask. The total number of bits must be 24 (x.x.x.x/24)`,
				},
				resource.Attribute{
					Name:        "networking_vpc_id",
					Description: `(Optional) Either an existing VPC Id (already exists in the specific region) or create a new VPC (if no VPC is specified). VPC Identifier must be in a valid format (for example: ‘vpc-0125be68a4986384ad’) and existing within the hosting account.`,
				},
				resource.Attribute{
					Name:        "preferred_availability_zones",
					Description: `(Optional) Availability zones deployment preferences (for the selected provider & region). If multiple_availability_zones is set to 'true', select three availability zones from the list. If you don't want to specify preferred availability zones, set this attribute to an empty list ('[]'). ~>`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Defaults to 30 mins) Used when creating the subscription`,
				},
				resource.Attribute{
					Name:        "update",
					Description: `(Defaults to 30 mins) Used when updating the subscription`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Defaults to 10 mins) Used when destroying the subscription ## Attribute reference The ` + "`" + `region` + "`" + ` block has these attributes:`,
				},
				resource.Attribute{
					Name:        "networks",
					Description: `List of generated network configuration The ` + "`" + `networks` + "`" + ` block has these attributes:`,
				},
				resource.Attribute{
					Name:        "networking_subnet_id",
					Description: `The subnet that the subscription deploys into`,
				},
				resource.Attribute{
					Name:        "networking_deployment_cidr",
					Description: `Deployment CIDR mask for the generated`,
				},
				resource.Attribute{
					Name:        "networking_vpc_id",
					Description: `VPC id for the generated network ## Import ` + "`" + `rediscloud_subscription` + "`" + ` can be imported using the ID of the subscription, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import rediscloud_subscription.subscription-resource 12345678 ` + "`" + `` + "`" + `` + "`" + ` ~>`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "rediscloud_rediscloud_subscription_database",
			Category:         "Resources",
			ShortDescription: `Database resource in the Terraform provider Redis Cloud.`,
			Description: `

Creates a Database within a specified Subscription in your Redis Enterprise Cloud Account.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "subscription_id",
					Description: `(Required) The ID of the subscription to create the database in`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A meaningful name to identify the database`,
				},
				resource.Attribute{
					Name:        "throughput_measurement_by",
					Description: `(Required) Throughput measurement method, (either ‘number-of-shards’ or ‘operations-per-second’)`,
				},
				resource.Attribute{
					Name:        "throughput_measurement_value",
					Description: `(Required) Throughput value (as applies to selected measurement method)`,
				},
				resource.Attribute{
					Name:        "memory_limit_in_gb",
					Description: `(Required) Maximum memory usage for this specific database`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Optional) The protocol that will be used to access the database, (either ‘redis’ or 'memcached’) Default: ‘redis’`,
				},
				resource.Attribute{
					Name:        "support_oss_cluster_api",
					Description: `(Optional) Support Redis open-source (OSS) Cluster API. Default: ‘false’`,
				},
				resource.Attribute{
					Name:        "external_endpoint_for_oss_cluster_api",
					Description: `(Optional) Should use the external endpoint for open-source (OSS) Cluster API. Can only be enabled if OSS Cluster API support is enabled. Default: 'false'`,
				},
				resource.Attribute{
					Name:        "client_ssl_certificate",
					Description: `(Optional) SSL certificate to authenticate user connections`,
				},
				resource.Attribute{
					Name:        "periodic_backup_path",
					Description: `(Optional) Path that will be used to store database backup files`,
				},
				resource.Attribute{
					Name:        "replica_of",
					Description: `(Optional) Set of Redis database URIs, in the format ` + "`" + `redis://user:password@host:port` + "`" + `, that this database will be a replica of. If the URI provided is Redis Labs Cloud instance, only host and port should be provided. Cannot be enabled when ` + "`" + `support_oss_cluster_api` + "`" + ` is enabled.`,
				},
				resource.Attribute{
					Name:        "modules",
					Description: `(Optional) A list of modules objects, documented below`,
				},
				resource.Attribute{
					Name:        "alert",
					Description: `(Optional) A block defining Redis database alert, documented below, can be specified multiple times`,
				},
				resource.Attribute{
					Name:        "data_persistence",
					Description: `(Optional) Rate of database data persistence (in persistent storage). Default: ‘none’`,
				},
				resource.Attribute{
					Name:        "data_eviction",
					Description: `(Optional) The data items eviction policy (either: 'allkeys-lru', 'allkeys-lfu', 'allkeys-random', 'volatile-lru', 'volatile-lfu', 'volatile-random', 'volatile-ttl' or 'noeviction'). Default: 'volatile-lru'`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional) Password to access the database. If omitted, a random 32 character long alphanumeric password will be automatically generated`,
				},
				resource.Attribute{
					Name:        "replication",
					Description: `(Optional) Databases replication. Default: ‘true’`,
				},
				resource.Attribute{
					Name:        "average_item_size_in_bytes",
					Description: `(Optional) Relevant only to ram-and-flash clusters. Estimated average size (measured in bytes) of the items stored in the database. Default: 1000`,
				},
				resource.Attribute{
					Name:        "source_ips",
					Description: `(Optional) List of source IP addresses or subnet masks. If specified, Redis clients will be able to connect to this database only from within the specified source IP addresses ranges (example: [‘192.168.10.0/32’, ‘192.168.12.0/24’])`,
				},
				resource.Attribute{
					Name:        "hashing_policy",
					Description: `(Optional) List of regular expression rules to shard the database by. See [the documentation on clustering](https://docs.redislabs.com/latest/rc/concepts/clustering/) for more information on the hashing policy. This cannot be set when ` + "`" + `support_oss_cluster_api` + "`" + ` is set to true.`,
				},
				resource.Attribute{
					Name:        "enable_tls",
					Description: `(Optional) Use TLS for authentication. Default: ‘false’ The ` + "`" + `alert` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Defaults to 30 mins) Used when creating the database`,
				},
				resource.Attribute{
					Name:        "update",
					Description: `(Defaults to 30 mins) Used when updating the database`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Defaults to 10 mins) Used when destroying the database ## Attribute reference`,
				},
				resource.Attribute{
					Name:        "db_id",
					Description: `Identifier of the database created`,
				},
				resource.Attribute{
					Name:        "public_endpoint",
					Description: `Public endpoint to access the database`,
				},
				resource.Attribute{
					Name:        "private_endpoint",
					Description: `Private endpoint to access the database ## Import ` + "`" + `rediscloud_subscription_database` + "`" + ` can be imported using the ID of the subscription and the ID of the database in the format {subscription ID}/{database ID}, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import rediscloud_subscription_database.database-resource 123456/12345678 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "rediscloud_rediscloud_subscription_peering",
			Category:         "Resources",
			ShortDescription: `Subscription VPC peering resource in the Terraform provider Redis Cloud.`,
			Description: `

Creates an AWS or GCP VPC peering for an existing Redis Enterprise Cloud Subscription, allowing access to your subscription databases as if they were on the same network.

For AWS, peering should be accepted by the other side.
For GCP, the opposite peering request should be submitted.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "provider_name",
					Description: `(Optional) The cloud provider to use with the vpc peering, (either ` + "`" + `AWS` + "`" + ` or ` + "`" + `GCP` + "`" + `). Default: ‘AWS’`,
				},
				resource.Attribute{
					Name:        "subscription_id",
					Description: `(Required) A valid subscription predefined in the current account`,
				},
				resource.Attribute{
					Name:        "aws_account_id",
					Description: `(Required AWS) AWS account ID that the VPC to be peered lives in`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Required AWS) AWS Region that the VPC to be peered lives in`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required AWS) Identifier of the VPC to be peered`,
				},
				resource.Attribute{
					Name:        "vpc_cidr",
					Description: `(Required AWS) CIDR range of the VPC to be peered`,
				},
				resource.Attribute{
					Name:        "gcp_project_id",
					Description: `(Required GCP) GCP project ID that the VPC to be peered lives in`,
				},
				resource.Attribute{
					Name:        "gcp_network_name",
					Description: `(Required GCP) The name of the network to be peered ### Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Defaults to 10 mins) Used when creating the peering connection`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Defaults to 10 mins) Used when deleting the peering connection ## Attribute reference`,
				},
			},
			Attributes: []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"rediscloud_rediscloud_active_active_subscription":          0,
		"rediscloud_rediscloud_active_active_subscription_database": 1,
		"rediscloud_rediscloud_active_active_subscription_peering":  2,
		"rediscloud_rediscloud_active_active_subscription_regions":  3,
		"rediscloud_rediscloud_cloud_account":                       4,
		"rediscloud_rediscloud_subscription":                        5,
		"rediscloud_rediscloud_subscription_database":               6,
		"rediscloud_rediscloud_subscription_peering":                7,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
