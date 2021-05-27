package vultr

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "vultr_account",
			Category:         "Data Sources",
			ShortDescription: `Get information about your Vultr account.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `The name on your Vultr account.`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `The email address on your Vultr account.`,
				},
				resource.Attribute{
					Name:        "acl",
					Description: `The access control list on your Vultr account.`,
				},
				resource.Attribute{
					Name:        "balance",
					Description: `The current balance on your Vultr account.`,
				},
				resource.Attribute{
					Name:        "pending_charges",
					Description: `The pending charges on your Vultr account.`,
				},
				resource.Attribute{
					Name:        "last_payment_date",
					Description: `The date of the last payment made on your Vultr account.`,
				},
				resource.Attribute{
					Name:        "last_payment_amount",
					Description: `The amount of the last payment made on your Vultr account.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `The name on your Vultr account.`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `The email address on your Vultr account.`,
				},
				resource.Attribute{
					Name:        "acl",
					Description: `The access control list on your Vultr account.`,
				},
				resource.Attribute{
					Name:        "balance",
					Description: `The current balance on your Vultr account.`,
				},
				resource.Attribute{
					Name:        "pending_charges",
					Description: `The pending charges on your Vultr account.`,
				},
				resource.Attribute{
					Name:        "last_payment_date",
					Description: `The date of the last payment made on your Vultr account.`,
				},
				resource.Attribute{
					Name:        "last_payment_amount",
					Description: `The amount of the last payment made on your Vultr account.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vultr_application",
			Category:         "Data Sources",
			ShortDescription: `Get information about applications that can be launched when creating a Vultr VPS.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `(Required) Query parameters for finding applications. The ` + "`" + `filter` + "`" + ` block supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Attribute name to filter with.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `One or more values filter with. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the application.`,
				},
				resource.Attribute{
					Name:        "deploy_name",
					Description: `The deploy name of the application.`,
				},
				resource.Attribute{
					Name:        "short_name",
					Description: `The short name of the application.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `The name of the application.`,
				},
				resource.Attribute{
					Name:        "deploy_name",
					Description: `The deploy name of the application.`,
				},
				resource.Attribute{
					Name:        "short_name",
					Description: `The short name of the application.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vultr_backup",
			Category:         "Data Sources",
			ShortDescription: `Get information about a Vultr backup.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `(Required) Query parameters for finding backups. The ` + "`" + `filter` + "`" + ` block supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Attribute name to filter with.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `One or more values filter with. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the backup`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the backup.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The size of the backup in bytes.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the backup.`,
				},
				resource.Attribute{
					Name:        "date_created",
					Description: `The date the backup was added to your Vultr account.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the backup`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the backup.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The size of the backup in bytes.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the backup.`,
				},
				resource.Attribute{
					Name:        "date_created",
					Description: `The date the backup was added to your Vultr account.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vultr_bare_metal_plan",
			Category:         "Data Sources",
			ShortDescription: `Get information about a Vultr bare metal server plan.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `(Required) Query parameters for finding plans. The ` + "`" + `filter` + "`" + ` block supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Attribute name to filter with.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `One or more values filter with. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "cpu_count",
					Description: `The number of CPUs available on the plan.`,
				},
				resource.Attribute{
					Name:        "cpu_model",
					Description: `The CPU model of the plan.`,
				},
				resource.Attribute{
					Name:        "cpu_threads",
					Description: `The number of CPU threads.`,
				},
				resource.Attribute{
					Name:        "ram",
					Description: `The amount of memory available on the plan in MB.`,
				},
				resource.Attribute{
					Name:        "disk",
					Description: `The description of the disk(s) on the plan.`,
				},
				resource.Attribute{
					Name:        "bandwidth",
					Description: `The bandwidth available on the plan.`,
				},
				resource.Attribute{
					Name:        "monthly_cost",
					Description: `The price per month of the plan in USD.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of plan it is.`,
				},
				resource.Attribute{
					Name:        "locations",
					Description: `A list of DCIDs (used as ` + "`" + `region` + "`" + ` in Terraform) where the plan can be deployed.`,
				},
				resource.Attribute{
					Name:        "disk_count",
					Description: `The number of disks that this plan offers.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "cpu_count",
					Description: `The number of CPUs available on the plan.`,
				},
				resource.Attribute{
					Name:        "cpu_model",
					Description: `The CPU model of the plan.`,
				},
				resource.Attribute{
					Name:        "cpu_threads",
					Description: `The number of CPU threads.`,
				},
				resource.Attribute{
					Name:        "ram",
					Description: `The amount of memory available on the plan in MB.`,
				},
				resource.Attribute{
					Name:        "disk",
					Description: `The description of the disk(s) on the plan.`,
				},
				resource.Attribute{
					Name:        "bandwidth",
					Description: `The bandwidth available on the plan.`,
				},
				resource.Attribute{
					Name:        "monthly_cost",
					Description: `The price per month of the plan in USD.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of plan it is.`,
				},
				resource.Attribute{
					Name:        "locations",
					Description: `A list of DCIDs (used as ` + "`" + `region` + "`" + ` in Terraform) where the plan can be deployed.`,
				},
				resource.Attribute{
					Name:        "disk_count",
					Description: `The number of disks that this plan offers.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vultr_bare_metal_server",
			Category:         "Data Sources",
			ShortDescription: `Get information about a Vultr bare metal server.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `(Required) Query parameters for finding servers. The ` + "`" + `filter` + "`" + ` block supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Attribute name to filter with.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `One or more values filter with. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "os",
					Description: `The operating system of the server.`,
				},
				resource.Attribute{
					Name:        "ram",
					Description: `The amount of memory available on the server in MB.`,
				},
				resource.Attribute{
					Name:        "disk",
					Description: `The description of the disk(s) on the server.`,
				},
				resource.Attribute{
					Name:        "main_ip",
					Description: `The server's main IP address.`,
				},
				resource.Attribute{
					Name:        "cpu_count",
					Description: `The number of CPUs available on the server.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The location of the server.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The region ID of the server.`,
				},
				resource.Attribute{
					Name:        "default_password",
					Description: `The server's default password.`,
				},
				resource.Attribute{
					Name:        "date_created",
					Description: `The date the server was added to your Vultr account.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the server's subscription.`,
				},
				resource.Attribute{
					Name:        "netmask_v4",
					Description: `The server's IPv4 netmask.`,
				},
				resource.Attribute{
					Name:        "gateway_v4",
					Description: `The server's IPv4 gateway.`,
				},
				resource.Attribute{
					Name:        "plan",
					Description: `The server's plan ID.`,
				},
				resource.Attribute{
					Name:        "v6_networks",
					Description: `A list of the server's IPv6 networks.`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `The server's label.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `The server's tag.`,
				},
				resource.Attribute{
					Name:        "os_id",
					Description: `The server's operating system ID.`,
				},
				resource.Attribute{
					Name:        "app_id",
					Description: `The server's application ID.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "os",
					Description: `The operating system of the server.`,
				},
				resource.Attribute{
					Name:        "ram",
					Description: `The amount of memory available on the server in MB.`,
				},
				resource.Attribute{
					Name:        "disk",
					Description: `The description of the disk(s) on the server.`,
				},
				resource.Attribute{
					Name:        "main_ip",
					Description: `The server's main IP address.`,
				},
				resource.Attribute{
					Name:        "cpu_count",
					Description: `The number of CPUs available on the server.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The location of the server.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The region ID of the server.`,
				},
				resource.Attribute{
					Name:        "default_password",
					Description: `The server's default password.`,
				},
				resource.Attribute{
					Name:        "date_created",
					Description: `The date the server was added to your Vultr account.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the server's subscription.`,
				},
				resource.Attribute{
					Name:        "netmask_v4",
					Description: `The server's IPv4 netmask.`,
				},
				resource.Attribute{
					Name:        "gateway_v4",
					Description: `The server's IPv4 gateway.`,
				},
				resource.Attribute{
					Name:        "plan",
					Description: `The server's plan ID.`,
				},
				resource.Attribute{
					Name:        "v6_networks",
					Description: `A list of the server's IPv6 networks.`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `The server's label.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `The server's tag.`,
				},
				resource.Attribute{
					Name:        "os_id",
					Description: `The server's operating system ID.`,
				},
				resource.Attribute{
					Name:        "app_id",
					Description: `The server's application ID.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vultr_block_storage",
			Category:         "Data Sources",
			ShortDescription: `Get information about a Vultr block storage subscription.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `(Required) Query parameters for finding block storage subscriptions. The ` + "`" + `filter` + "`" + ` block supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Attribute name to filter with.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `One or more values filter with. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `The label of the block storage subscription.`,
				},
				resource.Attribute{
					Name:        "cost",
					Description: `The cost per month of the block storage subscription in USD.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the block storage subscription.`,
				},
				resource.Attribute{
					Name:        "size_gb",
					Description: `The size of the block storage subscription in GB.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The region ID of the block storage subscription.`,
				},
				resource.Attribute{
					Name:        "attached_to_instance",
					Description: `The ID of the VPS the block storage subscription is attached to.`,
				},
				resource.Attribute{
					Name:        "date_created",
					Description: `The date the block storage subscription was added to your Vultr account.`,
				},
				resource.Attribute{
					Name:        "mount_id",
					Description: `An ID associated with the instance, when mounted the ID can be found in /dev/disk/by-id prefixed with virtio.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "label",
					Description: `The label of the block storage subscription.`,
				},
				resource.Attribute{
					Name:        "cost",
					Description: `The cost per month of the block storage subscription in USD.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the block storage subscription.`,
				},
				resource.Attribute{
					Name:        "size_gb",
					Description: `The size of the block storage subscription in GB.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The region ID of the block storage subscription.`,
				},
				resource.Attribute{
					Name:        "attached_to_instance",
					Description: `The ID of the VPS the block storage subscription is attached to.`,
				},
				resource.Attribute{
					Name:        "date_created",
					Description: `The date the block storage subscription was added to your Vultr account.`,
				},
				resource.Attribute{
					Name:        "mount_id",
					Description: `An ID associated with the instance, when mounted the ID can be found in /dev/disk/by-id prefixed with virtio.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vultr_dns_domain",
			Category:         "Data Sources",
			ShortDescription: `Get information about a DNS domain associated with your Vultr account.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "domain",
					Description: `(Required) The name you're searching for. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `Name of domain.`,
				},
				resource.Attribute{
					Name:        "date_created",
					Description: `The date the DNS domain was added to your Vultr account.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "domain",
					Description: `Name of domain.`,
				},
				resource.Attribute{
					Name:        "date_created",
					Description: `The date the DNS domain was added to your Vultr account.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vultr_firewall_group",
			Category:         "Data Sources",
			ShortDescription: `Get information about a firewall group on your Vultr account.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `(Required) Query parameters for finding firewall groups. The ` + "`" + `filter` + "`" + ` block supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Attribute name to filter with.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `One or more values filter with. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the firewall group.`,
				},
				resource.Attribute{
					Name:        "date_created",
					Description: `The date the firewall group was added to your Vultr account.`,
				},
				resource.Attribute{
					Name:        "date_modified",
					Description: `The date the firewall group was last modified.`,
				},
				resource.Attribute{
					Name:        "instance_count",
					Description: `The number of instances this firewall group is applied to.`,
				},
				resource.Attribute{
					Name:        "rule_count",
					Description: `The number of rules added to this firewall group.`,
				},
				resource.Attribute{
					Name:        "max_rule_count",
					Description: `The maximum number of rules this firewall group can have.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `The description of the firewall group.`,
				},
				resource.Attribute{
					Name:        "date_created",
					Description: `The date the firewall group was added to your Vultr account.`,
				},
				resource.Attribute{
					Name:        "date_modified",
					Description: `The date the firewall group was last modified.`,
				},
				resource.Attribute{
					Name:        "instance_count",
					Description: `The number of instances this firewall group is applied to.`,
				},
				resource.Attribute{
					Name:        "rule_count",
					Description: `The number of rules added to this firewall group.`,
				},
				resource.Attribute{
					Name:        "max_rule_count",
					Description: `The maximum number of rules this firewall group can have.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vultr_instance",
			Category:         "Data Sources",
			ShortDescription: `Get information about a Vultr instance.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `(Required) Query parameters for finding instances. The ` + "`" + `filter` + "`" + ` block supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Attribute name to filter with.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `One or more values filter with. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "os",
					Description: `The operating system of the instance.`,
				},
				resource.Attribute{
					Name:        "ram",
					Description: `The amount of memory available on the instance in MB.`,
				},
				resource.Attribute{
					Name:        "disk",
					Description: `The description of the disk(s) on the server.`,
				},
				resource.Attribute{
					Name:        "main_ip",
					Description: `The server's main IP address.`,
				},
				resource.Attribute{
					Name:        "vcpu_count",
					Description: `The number of virtual CPUs available on the server.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The region ID of the server.`,
				},
				resource.Attribute{
					Name:        "date_created",
					Description: `The date the server was added to your Vultr account.`,
				},
				resource.Attribute{
					Name:        "allowed_bandwidth",
					Description: `The server's allowed bandwidth usage in GB.`,
				},
				resource.Attribute{
					Name:        "netmask_v4",
					Description: `The server's IPv4 netmask.`,
				},
				resource.Attribute{
					Name:        "gateway_v4",
					Description: `The server's IPv4 gateway.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the server's subscription.`,
				},
				resource.Attribute{
					Name:        "power_status",
					Description: `Whether the server is powered on or not.`,
				},
				resource.Attribute{
					Name:        "server_status",
					Description: `A more detailed server status (none, locked, installingbooting, isomounting, ok).`,
				},
				resource.Attribute{
					Name:        "plan",
					Description: `The server's plan ID.`,
				},
				resource.Attribute{
					Name:        "v6_network",
					Description: `The IPv6 subnet.`,
				},
				resource.Attribute{
					Name:        "v6_main_ip",
					Description: `The main IPv6 network address.`,
				},
				resource.Attribute{
					Name:        "v6_network_size",
					Description: `The IPv6 network size in bits.`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `The server's label.`,
				},
				resource.Attribute{
					Name:        "internal_ip",
					Description: `The server's internal IP address.`,
				},
				resource.Attribute{
					Name:        "kvm",
					Description: `The server's current KVM URL. This URL will change periodically. It is not advised to cache this value.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `The server's tag.`,
				},
				resource.Attribute{
					Name:        "os_id",
					Description: `The server's operating system ID.`,
				},
				resource.Attribute{
					Name:        "app_id",
					Description: `The server's application ID.`,
				},
				resource.Attribute{
					Name:        "firewall_group_id",
					Description: `The ID of the firewall group applied to this server.`,
				},
				resource.Attribute{
					Name:        "features",
					Description: `Array of which features are enabled.`,
				},
				resource.Attribute{
					Name:        "backups_schedule",
					Description: `The current configuration for backups`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "os",
					Description: `The operating system of the instance.`,
				},
				resource.Attribute{
					Name:        "ram",
					Description: `The amount of memory available on the instance in MB.`,
				},
				resource.Attribute{
					Name:        "disk",
					Description: `The description of the disk(s) on the server.`,
				},
				resource.Attribute{
					Name:        "main_ip",
					Description: `The server's main IP address.`,
				},
				resource.Attribute{
					Name:        "vcpu_count",
					Description: `The number of virtual CPUs available on the server.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The region ID of the server.`,
				},
				resource.Attribute{
					Name:        "date_created",
					Description: `The date the server was added to your Vultr account.`,
				},
				resource.Attribute{
					Name:        "allowed_bandwidth",
					Description: `The server's allowed bandwidth usage in GB.`,
				},
				resource.Attribute{
					Name:        "netmask_v4",
					Description: `The server's IPv4 netmask.`,
				},
				resource.Attribute{
					Name:        "gateway_v4",
					Description: `The server's IPv4 gateway.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the server's subscription.`,
				},
				resource.Attribute{
					Name:        "power_status",
					Description: `Whether the server is powered on or not.`,
				},
				resource.Attribute{
					Name:        "server_status",
					Description: `A more detailed server status (none, locked, installingbooting, isomounting, ok).`,
				},
				resource.Attribute{
					Name:        "plan",
					Description: `The server's plan ID.`,
				},
				resource.Attribute{
					Name:        "v6_network",
					Description: `The IPv6 subnet.`,
				},
				resource.Attribute{
					Name:        "v6_main_ip",
					Description: `The main IPv6 network address.`,
				},
				resource.Attribute{
					Name:        "v6_network_size",
					Description: `The IPv6 network size in bits.`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `The server's label.`,
				},
				resource.Attribute{
					Name:        "internal_ip",
					Description: `The server's internal IP address.`,
				},
				resource.Attribute{
					Name:        "kvm",
					Description: `The server's current KVM URL. This URL will change periodically. It is not advised to cache this value.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `The server's tag.`,
				},
				resource.Attribute{
					Name:        "os_id",
					Description: `The server's operating system ID.`,
				},
				resource.Attribute{
					Name:        "app_id",
					Description: `The server's application ID.`,
				},
				resource.Attribute{
					Name:        "firewall_group_id",
					Description: `The ID of the firewall group applied to this server.`,
				},
				resource.Attribute{
					Name:        "features",
					Description: `Array of which features are enabled.`,
				},
				resource.Attribute{
					Name:        "backups_schedule",
					Description: `The current configuration for backups`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vultr_instance_ipv4",
			Category:         "Data Sources",
			ShortDescription: `Get information about a Vultr instance IPv4.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `(Required) Query parameters for finding IPv4 address. The ` + "`" + `filter` + "`" + ` block supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Attribute name to filter with.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `One or more values to filter with. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `The ID of the instance the IPv4 address.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `The IPv4 address in canonical format.`,
				},
				resource.Attribute{
					Name:        "gateway",
					Description: `The gateway IP address.`,
				},
				resource.Attribute{
					Name:        "netmask",
					Description: `The IPv4 netmask in dot-decimal notation.`,
				},
				resource.Attribute{
					Name:        "reverse",
					Description: `The reverse DNS information for this IP address.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "instance_id",
					Description: `The ID of the instance the IPv4 address.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `The IPv4 address in canonical format.`,
				},
				resource.Attribute{
					Name:        "gateway",
					Description: `The gateway IP address.`,
				},
				resource.Attribute{
					Name:        "netmask",
					Description: `The IPv4 netmask in dot-decimal notation.`,
				},
				resource.Attribute{
					Name:        "reverse",
					Description: `The reverse DNS information for this IP address.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vultr_iso_private",
			Category:         "Data Sources",
			ShortDescription: `Get information about an ISO file uploaded to your Vultr account.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `(Required) Query parameters for finding ISO files. The ` + "`" + `filter` + "`" + ` block supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Attribute name to filter with.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `One or more values filter with. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "filename",
					Description: `The ISO file's filename.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the ISO file.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The size of the ISO file in bytes.`,
				},
				resource.Attribute{
					Name:        "md5sum",
					Description: `The md5 hash of the ISO file.`,
				},
				resource.Attribute{
					Name:        "sha512sum",
					Description: `The sha512 hash of the ISO file.`,
				},
				resource.Attribute{
					Name:        "date_created",
					Description: `The date the ISO file was added to your Vultr account.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "filename",
					Description: `The ISO file's filename.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the ISO file.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The size of the ISO file in bytes.`,
				},
				resource.Attribute{
					Name:        "md5sum",
					Description: `The md5 hash of the ISO file.`,
				},
				resource.Attribute{
					Name:        "sha512sum",
					Description: `The sha512 hash of the ISO file.`,
				},
				resource.Attribute{
					Name:        "date_created",
					Description: `The date the ISO file was added to your Vultr account.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vultr_iso_public",
			Category:         "Data Sources",
			ShortDescription: `Get information about a public ISO file offered in the Vultr ISO library.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `(Required) Query parameters for finding ISO files. The ` + "`" + `filter` + "`" + ` block supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Attribute name to filter with.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `One or more values filter with. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The ISO file's name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the ISO file.`,
				},
				resource.Attribute{
					Name:        "md5sum",
					Description: `The MD5Sum of the ISO file.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `The ISO file's name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the ISO file.`,
				},
				resource.Attribute{
					Name:        "md5sum",
					Description: `The MD5Sum of the ISO file.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vultr_load_balancer",
			Category:         "Data Sources",
			ShortDescription: `Get information about a Vultr Load Balancer.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `(Required) Query parameters for finding load balancers. The ` + "`" + `filter` + "`" + ` block supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Attribute name to filter with.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `One or more values filter with. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The region your load balancer is deployed in.`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `The load balancers label.`,
				},
				resource.Attribute{
					Name:        "balancing_algorithm",
					Description: `The balancing algorithm for your load balancer.`,
				},
				resource.Attribute{
					Name:        "proxy_protocol",
					Description: `Boolean value that indicates if Proxy Protocol is enabled.`,
				},
				resource.Attribute{
					Name:        "cookie_name",
					Description: `Name for your given sticky session.`,
				},
				resource.Attribute{
					Name:        "ssl_redirect",
					Description: `Boolean value that indicates if HTTP calls will be redirected to HTTPS.`,
				},
				resource.Attribute{
					Name:        "has_ssl",
					Description: `Boolean value that indicates if SSL is enabled.`,
				},
				resource.Attribute{
					Name:        "attached_instances",
					Description: `Array of instances that are currently attached to the load balancer.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Current status for the load balancer`,
				},
				resource.Attribute{
					Name:        "ipv4",
					Description: `IPv4 address for your load balancer.`,
				},
				resource.Attribute{
					Name:        "ipv6",
					Description: `IPv6 address for your load balancer.`,
				},
				resource.Attribute{
					Name:        "health_check",
					Description: `Defines the way load balancers should check for health. The configuration of a ` + "`" + `health_check` + "`" + ` is listed below.`,
				},
				resource.Attribute{
					Name:        "forwarding_rules",
					Description: `Defines the forwarding rules for a load balancer. The configuration of a ` + "`" + `forwarding_rules` + "`" + ` is listened below.`,
				},
				resource.Attribute{
					Name:        "private_network",
					Description: `Defines the private network the load balancer is attached to. ` + "`" + `health_check` + "`" + ` supports the following`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `The protocol used to traffic requests to the load balancer. Possible values are "http", "https", or "tcp".`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `The path on the attached instances that the load balancer should check against.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `The assigned port (integer) on the attached instances that the load balancer should check against.`,
				},
				resource.Attribute{
					Name:        "check_interval",
					Description: `Time in seconds to perform health check. Default value is 15.`,
				},
				resource.Attribute{
					Name:        "response_timeout",
					Description: `Time in seconds to wait for a health check response. Default value is 5.`,
				},
				resource.Attribute{
					Name:        "unhealthy_threshold",
					Description: `Number of failed attempts encountered before failover. Default value is 5.`,
				},
				resource.Attribute{
					Name:        "healthy_threshold",
					Description: `Number of failed attempts encountered before failover. Default value is 5. ` + "`" + `forwarding_rules` + "`" + ` supports the following`,
				},
				resource.Attribute{
					Name:        "frontend_protocol",
					Description: `Protocol on load balancer side. Possible values: "http", "https", "tcp".`,
				},
				resource.Attribute{
					Name:        "frontend_port",
					Description: `Port on load balancer side.`,
				},
				resource.Attribute{
					Name:        "backend_protocol",
					Description: `Protocol on instance side. Possible values: "http", "https", "tcp".`,
				},
				resource.Attribute{
					Name:        "target_port",
					Description: `Port on instance side. ` + "`" + `firewall_rules` + "`" + ` supports the following`,
				},
				resource.Attribute{
					Name:        "frontend_port",
					Description: `(Required) Port on load balancer side.`,
				},
				resource.Attribute{
					Name:        "ip_type",
					Description: `(Required) The type of ip this rule is - may be either v4 or v6.`,
				},
				resource.Attribute{
					Name:        "source",
					Description: `(Required) IP address with subnet that is allowed through the firewall. You may also pass in ` + "`" + `cloudflare` + "`" + ` which will allow only CloudFlares IP range.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `The region your load balancer is deployed in.`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `The load balancers label.`,
				},
				resource.Attribute{
					Name:        "balancing_algorithm",
					Description: `The balancing algorithm for your load balancer.`,
				},
				resource.Attribute{
					Name:        "proxy_protocol",
					Description: `Boolean value that indicates if Proxy Protocol is enabled.`,
				},
				resource.Attribute{
					Name:        "cookie_name",
					Description: `Name for your given sticky session.`,
				},
				resource.Attribute{
					Name:        "ssl_redirect",
					Description: `Boolean value that indicates if HTTP calls will be redirected to HTTPS.`,
				},
				resource.Attribute{
					Name:        "has_ssl",
					Description: `Boolean value that indicates if SSL is enabled.`,
				},
				resource.Attribute{
					Name:        "attached_instances",
					Description: `Array of instances that are currently attached to the load balancer.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Current status for the load balancer`,
				},
				resource.Attribute{
					Name:        "ipv4",
					Description: `IPv4 address for your load balancer.`,
				},
				resource.Attribute{
					Name:        "ipv6",
					Description: `IPv6 address for your load balancer.`,
				},
				resource.Attribute{
					Name:        "health_check",
					Description: `Defines the way load balancers should check for health. The configuration of a ` + "`" + `health_check` + "`" + ` is listed below.`,
				},
				resource.Attribute{
					Name:        "forwarding_rules",
					Description: `Defines the forwarding rules for a load balancer. The configuration of a ` + "`" + `forwarding_rules` + "`" + ` is listened below.`,
				},
				resource.Attribute{
					Name:        "private_network",
					Description: `Defines the private network the load balancer is attached to. ` + "`" + `health_check` + "`" + ` supports the following`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `The protocol used to traffic requests to the load balancer. Possible values are "http", "https", or "tcp".`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `The path on the attached instances that the load balancer should check against.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `The assigned port (integer) on the attached instances that the load balancer should check against.`,
				},
				resource.Attribute{
					Name:        "check_interval",
					Description: `Time in seconds to perform health check. Default value is 15.`,
				},
				resource.Attribute{
					Name:        "response_timeout",
					Description: `Time in seconds to wait for a health check response. Default value is 5.`,
				},
				resource.Attribute{
					Name:        "unhealthy_threshold",
					Description: `Number of failed attempts encountered before failover. Default value is 5.`,
				},
				resource.Attribute{
					Name:        "healthy_threshold",
					Description: `Number of failed attempts encountered before failover. Default value is 5. ` + "`" + `forwarding_rules` + "`" + ` supports the following`,
				},
				resource.Attribute{
					Name:        "frontend_protocol",
					Description: `Protocol on load balancer side. Possible values: "http", "https", "tcp".`,
				},
				resource.Attribute{
					Name:        "frontend_port",
					Description: `Port on load balancer side.`,
				},
				resource.Attribute{
					Name:        "backend_protocol",
					Description: `Protocol on instance side. Possible values: "http", "https", "tcp".`,
				},
				resource.Attribute{
					Name:        "target_port",
					Description: `Port on instance side. ` + "`" + `firewall_rules` + "`" + ` supports the following`,
				},
				resource.Attribute{
					Name:        "frontend_port",
					Description: `(Required) Port on load balancer side.`,
				},
				resource.Attribute{
					Name:        "ip_type",
					Description: `(Required) The type of ip this rule is - may be either v4 or v6.`,
				},
				resource.Attribute{
					Name:        "source",
					Description: `(Required) IP address with subnet that is allowed through the firewall. You may also pass in ` + "`" + `cloudflare` + "`" + ` which will allow only CloudFlares IP range.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vultr_object_storage",
			Category:         "Data Sources",
			ShortDescription: `Get information about a Object Storage subscription on Vultr.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `(Required) Query parameters for finding operating systems. The ` + "`" + `filter` + "`" + ` block supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Attribute name to filter with.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `One or more values filter with. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `The label of the object storage subscription.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The location which this subscription resides in.`,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `The identifying cluster ID.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The region ID of the object storage subscription.`,
				},
				resource.Attribute{
					Name:        "s3_access_key",
					Description: `Your access key.`,
				},
				resource.Attribute{
					Name:        "s3_hostname",
					Description: `The hostname for this subscription.`,
				},
				resource.Attribute{
					Name:        "s3_secret_key",
					Description: `Your secret key.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Current status of this object storage subscription.`,
				},
				resource.Attribute{
					Name:        "date_created",
					Description: `Date of creation for the object storage subscription.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "label",
					Description: `The label of the object storage subscription.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The location which this subscription resides in.`,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `The identifying cluster ID.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The region ID of the object storage subscription.`,
				},
				resource.Attribute{
					Name:        "s3_access_key",
					Description: `Your access key.`,
				},
				resource.Attribute{
					Name:        "s3_hostname",
					Description: `The hostname for this subscription.`,
				},
				resource.Attribute{
					Name:        "s3_secret_key",
					Description: `Your secret key.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Current status of this object storage subscription.`,
				},
				resource.Attribute{
					Name:        "date_created",
					Description: `Date of creation for the object storage subscription.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vultr_os",
			Category:         "Data Sources",
			ShortDescription: `Get information about operating systems that can be launched when creating a Vultr VPS.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `(Required) Query parameters for finding operating systems. The ` + "`" + `filter` + "`" + ` block supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Attribute name to filter with.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `One or more values filter with. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the operating system.`,
				},
				resource.Attribute{
					Name:        "arch",
					Description: `The architecture of the operating system.`,
				},
				resource.Attribute{
					Name:        "family",
					Description: `The family of the operating system.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `The name of the operating system.`,
				},
				resource.Attribute{
					Name:        "arch",
					Description: `The architecture of the operating system.`,
				},
				resource.Attribute{
					Name:        "family",
					Description: `The family of the operating system.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vultr_plan",
			Category:         "Data Sources",
			ShortDescription: `Get information about a Vultr plan.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `(Required) Query parameters for finding plans. The ` + "`" + `filter` + "`" + ` block supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Attribute name to filter with.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `One or more values filter with. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "vcpu_count",
					Description: `The number of virtual CPUs available on the plan.`,
				},
				resource.Attribute{
					Name:        "ram",
					Description: `The amount of memory available on the plan in MB.`,
				},
				resource.Attribute{
					Name:        "disk",
					Description: `The amount of disk space in GB available on the plan.`,
				},
				resource.Attribute{
					Name:        "bandwidth",
					Description: `The bandwidth available on the plan in GB.`,
				},
				resource.Attribute{
					Name:        "monthly_cost",
					Description: `The price per month of the plan in USD.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of plan it is.`,
				},
				resource.Attribute{
					Name:        "locations",
					Description: `A list of DCIDs (used as ` + "`" + `region` + "`" + ` in Terraform) where the plan can be deployed.`,
				},
				resource.Attribute{
					Name:        "disk_count",
					Description: `The number of disks that this plan offers.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "vcpu_count",
					Description: `The number of virtual CPUs available on the plan.`,
				},
				resource.Attribute{
					Name:        "ram",
					Description: `The amount of memory available on the plan in MB.`,
				},
				resource.Attribute{
					Name:        "disk",
					Description: `The amount of disk space in GB available on the plan.`,
				},
				resource.Attribute{
					Name:        "bandwidth",
					Description: `The bandwidth available on the plan in GB.`,
				},
				resource.Attribute{
					Name:        "monthly_cost",
					Description: `The price per month of the plan in USD.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of plan it is.`,
				},
				resource.Attribute{
					Name:        "locations",
					Description: `A list of DCIDs (used as ` + "`" + `region` + "`" + ` in Terraform) where the plan can be deployed.`,
				},
				resource.Attribute{
					Name:        "disk_count",
					Description: `The number of disks that this plan offers.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vultr_private_network",
			Category:         "Data Sources",
			ShortDescription: `Get information about a Vultr private network.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `(Required) Query parameters for finding private networks. The ` + "`" + `filter` + "`" + ` block supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Attribute name to filter with.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `One or more values filter with. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The ID of the region that the private network is in.`,
				},
				resource.Attribute{
					Name:        "v4_subnet",
					Description: `The IPv4 network address. For example: 10.1.1.0.`,
				},
				resource.Attribute{
					Name:        "v4_subnet_mask",
					Description: `The number of bits for the netmask in CIDR notation. Example: 20`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The private network's description.`,
				},
				resource.Attribute{
					Name:        "date_created",
					Description: `The date the private network was added to your Vultr account.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `The ID of the region that the private network is in.`,
				},
				resource.Attribute{
					Name:        "v4_subnet",
					Description: `The IPv4 network address. For example: 10.1.1.0.`,
				},
				resource.Attribute{
					Name:        "v4_subnet_mask",
					Description: `The number of bits for the netmask in CIDR notation. Example: 20`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The private network's description.`,
				},
				resource.Attribute{
					Name:        "date_created",
					Description: `The date the private network was added to your Vultr account.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vultr_region",
			Category:         "Data Sources",
			ShortDescription: `Get information about a Vultr region.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `(Required) Query parameters for finding regions. The ` + "`" + `filter` + "`" + ` block supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Attribute name to filter with.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `One or more values filter with. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "continent",
					Description: `The continent the region is in.`,
				},
				resource.Attribute{
					Name:        "country",
					Description: `The country the region is in.`,
				},
				resource.Attribute{
					Name:        "city",
					Description: `The city the region is in.`,
				},
				resource.Attribute{
					Name:        "options",
					Description: `Shows whether options like ddos protection or block storage are available in the region.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "continent",
					Description: `The continent the region is in.`,
				},
				resource.Attribute{
					Name:        "country",
					Description: `The country the region is in.`,
				},
				resource.Attribute{
					Name:        "city",
					Description: `The city the region is in.`,
				},
				resource.Attribute{
					Name:        "options",
					Description: `Shows whether options like ddos protection or block storage are available in the region.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vultr_reserved_ip",
			Category:         "Data Sources",
			ShortDescription: `Get information about a Vultr reserved IP address.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `(Required) Query parameters for finding reserved IP addresses. The ` + "`" + `filter` + "`" + ` block supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Attribute name to filter with.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `One or more values filter with. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The ID of the region that the reserved IP is in.`,
				},
				resource.Attribute{
					Name:        "ip_type",
					Description: `The IP type of the reserved IP.`,
				},
				resource.Attribute{
					Name:        "subnet",
					Description: `The subnet of the reserved IP.`,
				},
				resource.Attribute{
					Name:        "subnet_size",
					Description: `The subnet size of the reserved IP.`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `The label of the reserved IP.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `The ID of the VPS the reserved IP is attached to.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `The ID of the region that the reserved IP is in.`,
				},
				resource.Attribute{
					Name:        "ip_type",
					Description: `The IP type of the reserved IP.`,
				},
				resource.Attribute{
					Name:        "subnet",
					Description: `The subnet of the reserved IP.`,
				},
				resource.Attribute{
					Name:        "subnet_size",
					Description: `The subnet size of the reserved IP.`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `The label of the reserved IP.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `The ID of the VPS the reserved IP is attached to.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vultr_reverse_ipv4",
			Category:         "Data Sources",
			ShortDescription: `Get information about a Vultr Reverse IPv4.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `(Required) Query parameters for finding IPv4 reverse DNS records. The ` + "`" + `filter` + "`" + ` block supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Attribute name to filter with.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `One or more values to filter with. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `The ID of the instance the IPv4 reverse DNS record was set for.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `The IPv4 address in canonical format used in the reverse DNS record.`,
				},
				resource.Attribute{
					Name:        "reverse",
					Description: `The hostname used in the IPv4 reverse DNS record.`,
				},
				resource.Attribute{
					Name:        "gateway",
					Description: `The gateway IP address.`,
				},
				resource.Attribute{
					Name:        "netmask",
					Description: `The IPv4 netmask in dot-decimal notation.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "instance_id",
					Description: `The ID of the instance the IPv4 reverse DNS record was set for.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `The IPv4 address in canonical format used in the reverse DNS record.`,
				},
				resource.Attribute{
					Name:        "reverse",
					Description: `The hostname used in the IPv4 reverse DNS record.`,
				},
				resource.Attribute{
					Name:        "gateway",
					Description: `The gateway IP address.`,
				},
				resource.Attribute{
					Name:        "netmask",
					Description: `The IPv4 netmask in dot-decimal notation.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vultr_reverse_ipv6",
			Category:         "Data Sources",
			ShortDescription: `Get information about a Vultr Reverse IPv6.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `(Required) Query parameters for finding IPv6 reverse DNS records. The ` + "`" + `filter` + "`" + ` block supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Attribute name to filter with.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `One or more values to filter with. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `The ID of the instance the IPv6 reverse DNS record was set for.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `The IPv6 address in canonical format used in the reverse DNS record.`,
				},
				resource.Attribute{
					Name:        "reverse",
					Description: `The hostname used in the IPv6 reverse DNS record.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "instance_id",
					Description: `The ID of the instance the IPv6 reverse DNS record was set for.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `The IPv6 address in canonical format used in the reverse DNS record.`,
				},
				resource.Attribute{
					Name:        "reverse",
					Description: `The hostname used in the IPv6 reverse DNS record.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vultr_snapshot",
			Category:         "Data Sources",
			ShortDescription: `Get information about a Vultr snapshot.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `(Required) Query parameters for finding snapshots. The ` + "`" + `filter` + "`" + ` block supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Attribute name to filter with.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `One or more values filter with. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID for the given snapshot.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the snapshot.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The size of the snapshot in bytes.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the snapshot.`,
				},
				resource.Attribute{
					Name:        "date_created",
					Description: `The date the snapshot was added to your Vultr account.`,
				},
				resource.Attribute{
					Name:        "os_id",
					Description: `The operating system ID of the snapshot.`,
				},
				resource.Attribute{
					Name:        "app_id",
					Description: `The application ID of the snapshot.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID for the given snapshot.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the snapshot.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The size of the snapshot in bytes.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the snapshot.`,
				},
				resource.Attribute{
					Name:        "date_created",
					Description: `The date the snapshot was added to your Vultr account.`,
				},
				resource.Attribute{
					Name:        "os_id",
					Description: `The operating system ID of the snapshot.`,
				},
				resource.Attribute{
					Name:        "app_id",
					Description: `The application ID of the snapshot.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vultr_ssh_key",
			Category:         "Data Sources",
			ShortDescription: `Get information about a Vultr SSH key.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `(Required) Query parameters for finding SSH keys. The ` + "`" + `filter` + "`" + ` block supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Attribute name to filter with.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `One or more values filter with. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the SSH key.`,
				},
				resource.Attribute{
					Name:        "ssh_key",
					Description: `The public SSH key.`,
				},
				resource.Attribute{
					Name:        "date_created",
					Description: `The date the SSH key was added to your Vultr account.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `The name of the SSH key.`,
				},
				resource.Attribute{
					Name:        "ssh_key",
					Description: `The public SSH key.`,
				},
				resource.Attribute{
					Name:        "date_created",
					Description: `The date the SSH key was added to your Vultr account.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vultr_startup_script",
			Category:         "Data Sources",
			ShortDescription: `Get information about a Vultr startup script.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `(Required) Query parameters for finding startup scripts. The ` + "`" + `filter` + "`" + ` block supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Attribute name to filter with.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `One or more values filter with. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the startup script.`,
				},
				resource.Attribute{
					Name:        "script",
					Description: `The contents of the startup script base64 encoded.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of the startup script.`,
				},
				resource.Attribute{
					Name:        "date_created",
					Description: `The date the startup script was added to your Vultr account.`,
				},
				resource.Attribute{
					Name:        "date_modified",
					Description: `The date the startup script was last modified.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `The name of the startup script.`,
				},
				resource.Attribute{
					Name:        "script",
					Description: `The contents of the startup script base64 encoded.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of the startup script.`,
				},
				resource.Attribute{
					Name:        "date_created",
					Description: `The date the startup script was added to your Vultr account.`,
				},
				resource.Attribute{
					Name:        "date_modified",
					Description: `The date the startup script was last modified.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vultr_user",
			Category:         "Data Sources",
			ShortDescription: `Get information about a Vultr user associated with your account.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `(Required) Query parameters for finding users. The ` + "`" + `filter` + "`" + ` block supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Attribute name to filter with.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `One or more values filter with. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the user.`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `The email of the user.`,
				},
				resource.Attribute{
					Name:        "api_enabled",
					Description: `Whether API is enabled for the user.`,
				},
				resource.Attribute{
					Name:        "acl",
					Description: `The access control list for the user.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `The name of the user.`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `The email of the user.`,
				},
				resource.Attribute{
					Name:        "api_enabled",
					Description: `Whether API is enabled for the user.`,
				},
				resource.Attribute{
					Name:        "acl",
					Description: `The access control list for the user.`,
				},
			},
		},
	}

	dataSourcesMap = map[string]int{

		"vultr_account":           0,
		"vultr_application":       1,
		"vultr_backup":            2,
		"vultr_bare_metal_plan":   3,
		"vultr_bare_metal_server": 4,
		"vultr_block_storage":     5,
		"vultr_dns_domain":        6,
		"vultr_firewall_group":    7,
		"vultr_instance":          8,
		"vultr_instance_ipv4":     9,
		"vultr_iso_private":       10,
		"vultr_iso_public":        11,
		"vultr_load_balancer":     12,
		"vultr_object_storage":    13,
		"vultr_os":                14,
		"vultr_plan":              15,
		"vultr_private_network":   16,
		"vultr_region":            17,
		"vultr_reserved_ip":       18,
		"vultr_reverse_ipv4":      19,
		"vultr_reverse_ipv6":      20,
		"vultr_snapshot":          21,
		"vultr_ssh_key":           22,
		"vultr_startup_script":    23,
		"vultr_user":              24,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
