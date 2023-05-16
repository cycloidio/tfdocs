package rediscloud

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "rediscloud_rediscloud_cloud_account",
			Category:         "Data Sources",
			ShortDescription: `Cloud Account data source in the Terraform provider Redis Cloud.`,
			Description: `

The Cloud Account data source allows access to the ID of a Cloud Account configuration.  This ID can be 
used when creating Subscription resources. 

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "exclude_internal_account",
					Description: `(Optional) Whether to exclude the Redis Labs internal cloud account.`,
				},
				resource.Attribute{
					Name:        "provider_type",
					Description: `(Optional) The cloud provider of the cloud account, (either ` + "`" + `AWS` + "`" + ` or ` + "`" + `GCP` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) A meaningful name to identify the cloud account ## Attributes Reference ` + "`" + `id` + "`" + ` is set to the ID of the found cloud account.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "rediscloud_rediscloud_data_persistence",
			Category:         "Data Sources",
			ShortDescription: `Data Persistence data source in the Terraform provider Redis Cloud.`,
			Description: `

The data persistence data source allows access to a list of supported data persistence options.  
Each option represents the rate at which a database will persist its data to storage.

`,
			Keywords:  []string{},
			Arguments: []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `The identifier of the data persistence option.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A meaningful description of the data persistence option.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "rediscloud_rediscloud_database",
			Category:         "Data Sources",
			ShortDescription: `Database data source in the Terraform provider Redis Cloud.`,
			Description: `

The Database data source allows access to the details of an existing database within your Redis Enterprise Cloud account.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "subscription_id",
					Description: `(Required) ID of the subscription that the database belongs to`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the database to filter returned databases`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Optional) The protocol of the database to filter returned databases`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region of the database to filter returned databases ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the database`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `The protocol of the database.`,
				},
				resource.Attribute{
					Name:        "memory_limit_in_gb",
					Description: `The maximum memory usage for the database.`,
				},
				resource.Attribute{
					Name:        "support_oss_cluster_api",
					Description: `Supports the Redis open-source (OSS) Cluster API.`,
				},
				resource.Attribute{
					Name:        "replica_of",
					Description: `The set of Redis database URIs, in the format ` + "`" + `redis://user:password@host:port` + "`" + `, that this database will be a replica of.`,
				},
				resource.Attribute{
					Name:        "alert",
					Description: `Set of alerts to enable on the database, documented below.`,
				},
				resource.Attribute{
					Name:        "data_persistence",
					Description: `The rate of database data persistence (in persistent storage).`,
				},
				resource.Attribute{
					Name:        "data_eviction",
					Description: `The data items eviction policy.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `The password used to access the database - not present on ` + "`" + `memcached` + "`" + ` protocol databases.`,
				},
				resource.Attribute{
					Name:        "replication",
					Description: `Database replication.`,
				},
				resource.Attribute{
					Name:        "throughput_measurement_by",
					Description: `The throughput measurement method.`,
				},
				resource.Attribute{
					Name:        "throughput_measurement_value",
					Description: `The throughput value.`,
				},
				resource.Attribute{
					Name:        "hashing_policy",
					Description: `The list of regular expression rules the database is sharded by. See [the documentation on clustering](https://docs.redislabs.com/latest/rc/concepts/clustering/) for more information on the hashing policy.`,
				},
				resource.Attribute{
					Name:        "public_endpoint",
					Description: `Public endpoint to access the database`,
				},
				resource.Attribute{
					Name:        "private_endpoint",
					Description: `Private endpoint to access the database`,
				},
				resource.Attribute{
					Name:        "enable_tls",
					Description: `Enable TLS for database, default is ` + "`" + `false` + "`" + ` The ` + "`" + `alert` + "`" + ` block supports:`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `The name of the database`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `The protocol of the database.`,
				},
				resource.Attribute{
					Name:        "memory_limit_in_gb",
					Description: `The maximum memory usage for the database.`,
				},
				resource.Attribute{
					Name:        "support_oss_cluster_api",
					Description: `Supports the Redis open-source (OSS) Cluster API.`,
				},
				resource.Attribute{
					Name:        "replica_of",
					Description: `The set of Redis database URIs, in the format ` + "`" + `redis://user:password@host:port` + "`" + `, that this database will be a replica of.`,
				},
				resource.Attribute{
					Name:        "alert",
					Description: `Set of alerts to enable on the database, documented below.`,
				},
				resource.Attribute{
					Name:        "data_persistence",
					Description: `The rate of database data persistence (in persistent storage).`,
				},
				resource.Attribute{
					Name:        "data_eviction",
					Description: `The data items eviction policy.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `The password used to access the database - not present on ` + "`" + `memcached` + "`" + ` protocol databases.`,
				},
				resource.Attribute{
					Name:        "replication",
					Description: `Database replication.`,
				},
				resource.Attribute{
					Name:        "throughput_measurement_by",
					Description: `The throughput measurement method.`,
				},
				resource.Attribute{
					Name:        "throughput_measurement_value",
					Description: `The throughput value.`,
				},
				resource.Attribute{
					Name:        "hashing_policy",
					Description: `The list of regular expression rules the database is sharded by. See [the documentation on clustering](https://docs.redislabs.com/latest/rc/concepts/clustering/) for more information on the hashing policy.`,
				},
				resource.Attribute{
					Name:        "public_endpoint",
					Description: `Public endpoint to access the database`,
				},
				resource.Attribute{
					Name:        "private_endpoint",
					Description: `Private endpoint to access the database`,
				},
				resource.Attribute{
					Name:        "enable_tls",
					Description: `Enable TLS for database, default is ` + "`" + `false` + "`" + ` The ` + "`" + `alert` + "`" + ` block supports:`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "rediscloud_rediscloud_database_modules",
			Category:         "Data Sources",
			ShortDescription: `Database Modules data source in the Terraform provider Redis Cloud.`,
			Description: `

The database modules data source allows access to a list of supported [Redis Enterprise Cloud modules](https://redislabs.com/redis-enterprise/modules).  
Each module represents an enrichment that can be applied to a Redis database.

`,
			Keywords:   []string{},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "rediscloud_rediscloud_payment_method",
			Category:         "Data Sources",
			ShortDescription: `Payment method data source in the Terraform provider Redis Cloud.`,
			Description: `

The Payment Method data source allows access to the ID of a Payment Method configured against your Redis Enterprise Cloud account. This ID can be used when creating Subscription resources.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "card_type",
					Description: `(Optional) Type of card that the payment method should be, such as ` + "`" + `Visa` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "last_four_numbers",
					Description: `(Optional) Last four numbers of the card of the payment method.`,
				},
				resource.Attribute{
					Name:        "exclude_expired",
					Description: `(Optional) Whether to exclude any expired cards or not. Default is ` + "`" + `true` + "`" + `. ## Attributes Reference ` + "`" + `id` + "`" + ` is set to the ID of the found payment method.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "rediscloud_rediscloud_regions",
			Category:         "Data Sources",
			ShortDescription: `Regions data source in the Terraform provider Redis Cloud.`,
			Description: `

The Regions data source allows access to a list of supported cloud provider regions. These regions can be used with the subscription resource.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "provider_name",
					Description: `(Optional) The name of the cloud provider to filter returned regions, (accepted values are ` + "`" + `AWS` + "`" + ` or ` + "`" + `GCP` + "`" + `). ## Attributes Reference`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "rediscloud_rediscloud_subscription",
			Category:         "Data Sources",
			ShortDescription: `Subscription data source in the Terraform provider Redis Cloud.`,
			Description: `

The Subscription data source allows access to the details of an existing subscription within your Redis Enterprise Cloud account.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the subscription to filter returned subscriptions ## Attributes Reference ` + "`" + `id` + "`" + ` is set to the ID of the found subscription.`,
				},
				resource.Attribute{
					Name:        "payment_method_id",
					Description: `A valid payment method pre-defined in the current account`,
				},
				resource.Attribute{
					Name:        "memory_storage",
					Description: `Memory storage preference: either ‘ram’ or a combination of 'ram-and-flash’`,
				},
				resource.Attribute{
					Name:        "cloud_provider",
					Description: `A cloud provider object, documented below`,
				},
				resource.Attribute{
					Name:        "number_of_databases",
					Description: `The number of databases that are linked to this subscription.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Current status of the subscription The ` + "`" + `cloud_provider` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "provider",
					Description: `The cloud provider to use with the subscription, (either ` + "`" + `AWS` + "`" + ` or ` + "`" + `GCP` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "cloud_account_id",
					Description: `Cloud account identifier, (A Cloud Account Id = 1 implies using Redis Labs internal cloud account)`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `Cloud networking details, per region (single region or multiple regions for Active-Active cluster only), documented below The cloud_provider ` + "`" + `region` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `Deployment region as defined by cloud provider`,
				},
				resource.Attribute{
					Name:        "multiple_availability_zones",
					Description: `Support deployment on multiple availability zones within the selected region`,
				},
				resource.Attribute{
					Name:        "networking_vpc_id",
					Description: `The ID of the VPC where the Redis Cloud subscription is deployed.`,
				},
				resource.Attribute{
					Name:        "preferred_availability_zones",
					Description: `List of availability zones used`,
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
					Description: `VPC id for the generated network`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "payment_method_id",
					Description: `A valid payment method pre-defined in the current account`,
				},
				resource.Attribute{
					Name:        "memory_storage",
					Description: `Memory storage preference: either ‘ram’ or a combination of 'ram-and-flash’`,
				},
				resource.Attribute{
					Name:        "cloud_provider",
					Description: `A cloud provider object, documented below`,
				},
				resource.Attribute{
					Name:        "number_of_databases",
					Description: `The number of databases that are linked to this subscription.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Current status of the subscription The ` + "`" + `cloud_provider` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "provider",
					Description: `The cloud provider to use with the subscription, (either ` + "`" + `AWS` + "`" + ` or ` + "`" + `GCP` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "cloud_account_id",
					Description: `Cloud account identifier, (A Cloud Account Id = 1 implies using Redis Labs internal cloud account)`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `Cloud networking details, per region (single region or multiple regions for Active-Active cluster only), documented below The cloud_provider ` + "`" + `region` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `Deployment region as defined by cloud provider`,
				},
				resource.Attribute{
					Name:        "multiple_availability_zones",
					Description: `Support deployment on multiple availability zones within the selected region`,
				},
				resource.Attribute{
					Name:        "networking_vpc_id",
					Description: `The ID of the VPC where the Redis Cloud subscription is deployed.`,
				},
				resource.Attribute{
					Name:        "preferred_availability_zones",
					Description: `List of availability zones used`,
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
					Description: `VPC id for the generated network`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "rediscloud_rediscloud_subscription_peerings",
			Category:         "Data Sources",
			ShortDescription: `Subscription Peerings data source in the Terraform provider Redis Cloud.`,
			Description: `

The Subscription Peerings data source allows access to a list of VPC peerings for a particular subscription.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "subscription_id",
					Description: `(Required) ID of the subscription that the peerings belongs to`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Optional) Current status of the peering - ` + "`" + `initiating-request` + "`" + `, ` + "`" + `pending-acceptance` + "`" + `, ` + "`" + `active` + "`" + `, ` + "`" + `inactive` + "`" + ` or ` + "`" + `failed` + "`" + `. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "peering_id",
					Description: `ID of the subscription peering`,
				},
				resource.Attribute{
					Name:        "provider_name",
					Description: `The name of the cloud provider. (either ` + "`" + `AWS` + "`" + ` or ` + "`" + `GCP` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "gcp_redis_network_name",
					Description: `The name of the Redis Enterprise Cloud network to be peered`,
				},
				resource.Attribute{
					Name:        "gcp_peering_id",
					Description: `Identifier of the cloud peering`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "peering_id",
					Description: `ID of the subscription peering`,
				},
				resource.Attribute{
					Name:        "provider_name",
					Description: `The name of the cloud provider. (either ` + "`" + `AWS` + "`" + ` or ` + "`" + `GCP` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "gcp_redis_network_name",
					Description: `The name of the Redis Enterprise Cloud network to be peered`,
				},
				resource.Attribute{
					Name:        "gcp_peering_id",
					Description: `Identifier of the cloud peering`,
				},
			},
		},
	}

	dataSourcesMap = map[string]int{

		"rediscloud_rediscloud_cloud_account":         0,
		"rediscloud_rediscloud_data_persistence":      1,
		"rediscloud_rediscloud_database":              2,
		"rediscloud_rediscloud_database_modules":      3,
		"rediscloud_rediscloud_payment_method":        4,
		"rediscloud_rediscloud_regions":               5,
		"rediscloud_rediscloud_subscription":          6,
		"rediscloud_rediscloud_subscription_peerings": 7,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
