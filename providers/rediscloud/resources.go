package rediscloud

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

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
This resource is responsible for creating subscriptions and the databases within that subscription. 
This allows Redis Enterprise Cloud to provision your databases in the most efficient way.

~> **Note:** The subscription resource manages changes to its databases by identifying a databases through its name.  This means **database names cannot be changed**, as this resource has no other way of being able to identify the database and would lead to the database to be destroyed.
Due to the limitations mentioned above, the differences shown by Terraform will be different from normal plan.
When an argument has been changed on a nested database - for example changing the ` + "`" + `memory_limit_in_gb` + "`" + ` from 1 to 2, Terraform
will display the resource as being modified with the database as being removed, and a new one added. As the resource
identifies the database based on the name, the only change that would happen would be to update the database to increase
the memory limit. Below is the Terraform output for changing the ` + "`" + `memory_limit_in_gb` + "`" + ` for a single database within a
subscription.

` + "`" + `` + "`" + `` + "`" + `
An execution plan has been generated and is shown below.
Resource actions are indicated with the following symbols:
  ~ update in-place

Terraform will perform the following actions:

  # rediscloud_subscription.example will be updated in-place
  ~ resource "rediscloud_subscription" "example" {

        ...

      - database {
          - average_item_size_in_bytes            = 0 -> null
          - data_persistence                      = "none" -> null
          - db_id                                 = 51040112 -> null
          - external_endpoint_for_oss_cluster_api = false -> null
          - memory_limit_in_gb                    = 1 -> null
          - name                                  = "tf-example-database" -> null
          - password                              = (sensitive value)
          - private_endpoint                      = "private.example.com" -> null
          - protocol                              = "redis" -> null
          - public_endpoint                       = "public.example.com" -> null
          - replica_of                            = [] -> null
          - replication                           = false -> null
          - source_ips                            = [] -> null
          - support_oss_cluster_api               = false -> null
          - throughput_measurement_by             = "operations-per-second" -> null
          - throughput_measurement_value          = 10000 -> null
        }
      + database {
          + average_item_size_in_bytes            = 0
          + data_persistence                      = "none"
          + db_id                                 = (known after apply)
          + external_endpoint_for_oss_cluster_api = false
          + memory_limit_in_gb                    = 2
          + name                                  = "tf-example-database"
          + password                              = (sensitive value)
          + private_endpoint                      = (known after apply)
          + protocol                              = "redis"
          + public_endpoint                       = (known after apply)
          + replica_of                            = []
          + replication                           = false
          + source_ips                            = []
          + support_oss_cluster_api               = false
          + throughput_measurement_by             = "operations-per-second"
          + throughput_measurement_value          = 10000
        }
    }

Plan: 0 to add, 1 to change, 0 to destroy.
` + "`" + `` + "`" + `` + "`" + `

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A meaningful name to identify the subscription`,
				},
				resource.Attribute{
					Name:        "payment_method_id",
					Description: `(Optional) A valid payment method pre-defined in the current account. This value is __Optional__ for AWS/GCP Marketplace accounts, but __Required__ for all other account types.`,
				},
				resource.Attribute{
					Name:        "memory_storage",
					Description: `(Optional) Memory storage preference: either ‘ram’ or a combination of 'ram-and-flash’. Default: ‘ram’`,
				},
				resource.Attribute{
					Name:        "persistent_storage_encryption",
					Description: `(Optional) Encrypt data stored in persistent storage. Required for a GCP subscription. Default: ‘true’`,
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
					Name:        "database",
					Description: `(Required) A database object, documented below The ` + "`" + `allowlist` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "cidrs",
					Description: `(Optional) Set of CIDR ranges that are allowed to access the databases associated with this subscription`,
				},
				resource.Attribute{
					Name:        "security_group_ids",
					Description: `(Optional) Set of security groups that are allowed to access the databases associated with this subscription ~>`,
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
					Description: `(Required) Cloud networking details, per region, documented below The ` + "`" + `database` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A meaningful name to identify the database. Caution should be taken when changing this value - see the top of the page for more information.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Optional) The protocol that will be used to access the database, (either ‘redis’ or 'memcached’) Default: ‘redis’`,
				},
				resource.Attribute{
					Name:        "memory_limit_in_gb",
					Description: `(Required) Maximum memory usage for this specific database`,
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
					Name:        "module",
					Description: `(Optional) A module object, documented below`,
				},
				resource.Attribute{
					Name:        "alert",
					Description: `(Optional) Set of alerts to enable on the database, documented below`,
				},
				resource.Attribute{
					Name:        "data_persistence",
					Description: `(Optional) Rate of database data persistence (in persistent storage). Default: ‘none’`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Required) Password used to access the database`,
				},
				resource.Attribute{
					Name:        "replication",
					Description: `(Optional) Databases replication. Default: ‘true’`,
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
					Name:        "average_item_size_in_bytes",
					Description: `(Optional) Relevant only to ram-and-flash clusters. Estimated average size (measured in bytes) of the items stored in the database. Default: 1000`,
				},
				resource.Attribute{
					Name:        "source_ips",
					Description: `(Optional) Set of CIDR addresses to allow access to the database. Defaults to allowing traffic.`,
				},
				resource.Attribute{
					Name:        "hashing_policy",
					Description: `(Optional) List of regular expression rules to shard the database by. See [the documentation on clustering](https://docs.redislabs.com/latest/rc/concepts/clustering/) for more information on the hashing policy. This cannot be set when ` + "`" + `support_oss_cluster_api` + "`" + ` is set to true. The cloud_provider ` + "`" + `region` + "`" + ` block supports:`,
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
					Description: `(Required) Deployment CIDR mask.`,
				},
				resource.Attribute{
					Name:        "networking_vpc_id",
					Description: `(Optional) Either an existing VPC Id (already exists in the specific region) or create a new VPC (if no VPC is specified). VPC Identifier must be in a valid format (for example: ‘vpc-0125be68a4625884ad’) and existing within the hosting account.`,
				},
				resource.Attribute{
					Name:        "preferred_availability_zones",
					Description: `(Required) Availability zones deployment preferences (for the selected provider & region). ~>`,
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
					Description: `(Defaults to 10 mins) Used when destroying the subscription ## Attribute reference The ` + "`" + `database` + "`" + ` block has these attributes:`,
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
					Description: `Private endpoint to access the database The ` + "`" + `region` + "`" + ` block has these attributes:`,
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
					Description: `VPC id for the generated network ## Import ` + "`" + `rediscloud_subscription` + "`" + ` can be imported using the ID of the subscription, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import rediscloud_subscription.example 12345678 ` + "`" + `` + "`" + `` + "`" + ``,
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

		"rediscloud_rediscloud_cloud_account":        0,
		"rediscloud_rediscloud_subscription":         1,
		"rediscloud_rediscloud_subscription_peering": 2,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
