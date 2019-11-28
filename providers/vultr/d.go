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
			Type:             "vultr_api_key",
			Category:         "Data Sources",
			ShortDescription: `Get information about your Vultr API key.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `The name associated with your Vultr API key.`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `The email associated with your Vultr API key.`,
				},
				resource.Attribute{
					Name:        "acl",
					Description: `The access control list for your Vultr API key.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `The name associated with your Vultr API key.`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `The email associated with your Vultr API key.`,
				},
				resource.Attribute{
					Name:        "acl",
					Description: `The access control list for your Vultr API key.`,
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
					Name:        "BACKUPID",
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
					Name:        "BACKUPID",
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
					Name:        "name",
					Description: `The name of the plan.`,
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
					Name:        "ram",
					Description: `The amount of memory available on the plan in MB.`,
				},
				resource.Attribute{
					Name:        "disk",
					Description: `The description of the disk(s) on the plan.`,
				},
				resource.Attribute{
					Name:        "bandwidth_tb",
					Description: `The bandwidth available on the plan in TB.`,
				},
				resource.Attribute{
					Name:        "price_per_month",
					Description: `The price per month of the plan in USD.`,
				},
				resource.Attribute{
					Name:        "plan_type",
					Description: `The type of plan it is.`,
				},
				resource.Attribute{
					Name:        "available_locations",
					Description: `A list of DCIDs (used as ` + "`" + `region_id` + "`" + ` in Terraform) where the plan can be deployed.`,
				},
				resource.Attribute{
					Name:        "deprecated",
					Description: `Indicates that the plan will be going away in the future. New deployments of it will still be accepted, but you should begin to transition away from its usage.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `The name of the plan.`,
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
					Name:        "ram",
					Description: `The amount of memory available on the plan in MB.`,
				},
				resource.Attribute{
					Name:        "disk",
					Description: `The description of the disk(s) on the plan.`,
				},
				resource.Attribute{
					Name:        "bandwidth_tb",
					Description: `The bandwidth available on the plan in TB.`,
				},
				resource.Attribute{
					Name:        "price_per_month",
					Description: `The price per month of the plan in USD.`,
				},
				resource.Attribute{
					Name:        "plan_type",
					Description: `The type of plan it is.`,
				},
				resource.Attribute{
					Name:        "available_locations",
					Description: `A list of DCIDs (used as ` + "`" + `region_id` + "`" + ` in Terraform) where the plan can be deployed.`,
				},
				resource.Attribute{
					Name:        "deprecated",
					Description: `Indicates that the plan will be going away in the future. New deployments of it will still be accepted, but you should begin to transition away from its usage.`,
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
					Name:        "region_id",
					Description: `The region ID (` + "`" + `DCID` + "`" + ` in the Vultr API) of the server.`,
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
					Name:        "plan_id",
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
					Name:        "region_id",
					Description: `The region ID (` + "`" + `DCID` + "`" + ` in the Vultr API) of the server.`,
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
					Name:        "plan_id",
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
					Name:        "cost_per_month",
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
					Name:        "region_id",
					Description: `The region ID (` + "`" + `DCID` + "`" + ` in the Vultr API) of the block storage subscription.`,
				},
				resource.Attribute{
					Name:        "attached_to_vps",
					Description: `The ID of the VPS the block storage subscription is attached to.`,
				},
				resource.Attribute{
					Name:        "date_created",
					Description: `The date the block storage subscription was added to your Vultr account.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "label",
					Description: `The label of the block storage subscription.`,
				},
				resource.Attribute{
					Name:        "cost_per_month",
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
					Name:        "region_id",
					Description: `The region ID (` + "`" + `DCID` + "`" + ` in the Vultr API) of the block storage subscription.`,
				},
				resource.Attribute{
					Name:        "attached_to_vps",
					Description: `The ID of the VPS the block storage subscription is attached to.`,
				},
				resource.Attribute{
					Name:        "date_created",
					Description: `The date the block storage subscription was added to your Vultr account.`,
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
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vultr_network",
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
					Name:        "region_id",
					Description: `The ID of the region that the private network is in.`,
				},
				resource.Attribute{
					Name:        "cidr_block",
					Description: `The CIDR block of the private network.`,
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
					Name:        "region_id",
					Description: `The ID of the region that the private network is in.`,
				},
				resource.Attribute{
					Name:        "cidr_block",
					Description: `The CIDR block of the private network.`,
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
				resource.Attribute{
					Name:        "windows",
					Description: `If true, a Windows license will be included with the instance, which will increase the cost.`,
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
				resource.Attribute{
					Name:        "windows",
					Description: `If true, a Windows license will be included with the instance, which will increase the cost.`,
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
					Name:        "name",
					Description: `The name of the plan.`,
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
					Description: `The bandwidth available on the plan in TB.`,
				},
				resource.Attribute{
					Name:        "bandwidth_gb",
					Description: `The bandwidth available on the plan in GB.`,
				},
				resource.Attribute{
					Name:        "price_per_month",
					Description: `The price per month of the plan in USD.`,
				},
				resource.Attribute{
					Name:        "plan_type",
					Description: `The type of plan it is.`,
				},
				resource.Attribute{
					Name:        "available_locations",
					Description: `A list of DCIDs (used as ` + "`" + `region_id` + "`" + ` in Terraform) where the plan can be deployed.`,
				},
				resource.Attribute{
					Name:        "deprecated",
					Description: `Indicates that the plan will be going away in the future. New deployments of it will still be accepted, but you should begin to transition away from its usage.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `The name of the plan.`,
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
					Description: `The bandwidth available on the plan in TB.`,
				},
				resource.Attribute{
					Name:        "bandwidth_gb",
					Description: `The bandwidth available on the plan in GB.`,
				},
				resource.Attribute{
					Name:        "price_per_month",
					Description: `The price per month of the plan in USD.`,
				},
				resource.Attribute{
					Name:        "plan_type",
					Description: `The type of plan it is.`,
				},
				resource.Attribute{
					Name:        "available_locations",
					Description: `A list of DCIDs (used as ` + "`" + `region_id` + "`" + ` in Terraform) where the plan can be deployed.`,
				},
				resource.Attribute{
					Name:        "deprecated",
					Description: `Indicates that the plan will be going away in the future. New deployments of it will still be accepted, but you should begin to transition away from its usage.`,
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
					Name:        "name",
					Description: `The name of the region.`,
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
					Name:        "state",
					Description: `The state the region is in.`,
				},
				resource.Attribute{
					Name:        "ddos_protection",
					Description: `Whether the region has DDoS protection.`,
				},
				resource.Attribute{
					Name:        "block_storage",
					Description: `Whether the region has block storage.`,
				},
				resource.Attribute{
					Name:        "regioncode",
					Description: `The region code of the region.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `The name of the region.`,
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
					Name:        "state",
					Description: `The state the region is in.`,
				},
				resource.Attribute{
					Name:        "ddos_protection",
					Description: `Whether the region has DDoS protection.`,
				},
				resource.Attribute{
					Name:        "block_storage",
					Description: `Whether the region has block storage.`,
				},
				resource.Attribute{
					Name:        "regioncode",
					Description: `The region code of the region.`,
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
					Name:        "region_id",
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
					Name:        "attached_to_vps",
					Description: `The ID of the VPS the reserved IP is attached to.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "region_id",
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
					Name:        "attached_to_vps",
					Description: `The ID of the VPS the reserved IP is attached to.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vultr_server",
			Category:         "Data Sources",
			ShortDescription: `Get information about a Vultr server.`,
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
					Name:        "vps_cpu_count",
					Description: `The number of virtual CPUs available on the server.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The physical location of the server.`,
				},
				resource.Attribute{
					Name:        "region_id",
					Description: `The region ID (` + "`" + `DCID` + "`" + ` in the Vultr API) of the server.`,
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
					Name:        "pending_charges",
					Description: `Charges pending for this server's subscription in USD.`,
				},
				resource.Attribute{
					Name:        "cost_per_month",
					Description: `The server's cost per month in USD.`,
				},
				resource.Attribute{
					Name:        "current_bandwidth",
					Description: `The server's current bandwidth usage in GB.`,
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
					Name:        "server_state",
					Description: `A more detailed server status (none, locked, installingbooting, isomounting, ok).`,
				},
				resource.Attribute{
					Name:        "plan_id",
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
					Name:        "internal_ip",
					Description: `The server's internal IP address.`,
				},
				resource.Attribute{
					Name:        "kvm_url",
					Description: `The server's current KVM URL. This URL will change periodically. It is not advised to cache this value.`,
				},
				resource.Attribute{
					Name:        "auto_backups",
					Description: `Whether auto backups are enabled on this server.`,
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
					Name:        "vps_cpu_count",
					Description: `The number of virtual CPUs available on the server.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The physical location of the server.`,
				},
				resource.Attribute{
					Name:        "region_id",
					Description: `The region ID (` + "`" + `DCID` + "`" + ` in the Vultr API) of the server.`,
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
					Name:        "pending_charges",
					Description: `Charges pending for this server's subscription in USD.`,
				},
				resource.Attribute{
					Name:        "cost_per_month",
					Description: `The server's cost per month in USD.`,
				},
				resource.Attribute{
					Name:        "current_bandwidth",
					Description: `The server's current bandwidth usage in GB.`,
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
					Name:        "server_state",
					Description: `A more detailed server status (none, locked, installingbooting, isomounting, ok).`,
				},
				resource.Attribute{
					Name:        "plan_id",
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
					Name:        "internal_ip",
					Description: `The server's internal IP address.`,
				},
				resource.Attribute{
					Name:        "kvm_url",
					Description: `The server's current KVM URL. This URL will change periodically. It is not advised to cache this value.`,
				},
				resource.Attribute{
					Name:        "auto_backups",
					Description: `Whether auto backups are enabled on this server.`,
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
					Description: `The contents of the startup script.`,
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
					Description: `The contents of the startup script.`,
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
		"vultr_api_key":           1,
		"vultr_application":       2,
		"vultr_backup":            3,
		"vultr_bare_metal_plan":   4,
		"vultr_bare_metal_server": 5,
		"vultr_block_storage":     6,
		"vultr_dns_domain":        7,
		"vultr_firewall_group":    8,
		"vultr_iso_private":       9,
		"vultr_iso_public":        10,
		"vultr_network":           11,
		"vultr_os":                12,
		"vultr_plan":              13,
		"vultr_region":            14,
		"vultr_reserved_ip":       15,
		"vultr_server":            16,
		"vultr_snapshot":          17,
		"vultr_ssh_key":           18,
		"vultr_startup_script":    19,
		"vultr_user":              20,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
