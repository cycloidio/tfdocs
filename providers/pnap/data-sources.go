package pnap

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "pnap_events",
			Category:         "Data Sources",
			ShortDescription: `Provides a phoenixNAP events datasource. This can be used to read event logs.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "events",
					Description: `(Required) Block ` + "`" + `events` + "`" + ` has field ` + "`" + `name` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Event name. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "events",
					Description: `The list of events recorded.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the event.`,
				},
				resource.Attribute{
					Name:        "timestamp",
					Description: `The UTC time the event initiated.`,
				},
				resource.Attribute{
					Name:        "user_info",
					Description: `Details related to the user / application.`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `The BMC account ID.`,
				},
				resource.Attribute{
					Name:        "client_id",
					Description: `The client ID of the application.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `The logged in user or owner of the client application.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "events",
					Description: `The list of events recorded.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the event.`,
				},
				resource.Attribute{
					Name:        "timestamp",
					Description: `The UTC time the event initiated.`,
				},
				resource.Attribute{
					Name:        "user_info",
					Description: `Details related to the user / application.`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `The BMC account ID.`,
				},
				resource.Attribute{
					Name:        "client_id",
					Description: `The client ID of the application.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `The logged in user or owner of the client application.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "pnap_ip_block",
			Category:         "Data Sources",
			ShortDescription: `Provides a phoenixNAP IP Block datasource. This can be used to read IP Blocks.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cidr",
					Description: `The IP Block in CIDR notation.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The IP Block identifier. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The IP Block identifier.`,
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
					Description: `The IP Block identifier.`,
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
			Category:         "Data Sources",
			ShortDescription: `Provides a phoenixNAP Private Network datasource. This can be used to read private networks.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `The friendly name of this private network.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The private network identifier. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The private network identifier.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The friendly name of this private network.`,
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
					Description: `(Deprecated) List of servers' details linked to the private network.`,
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
					Description: `The friendly name of this private network.`,
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
					Description: `(Deprecated) List of servers' details linked to the private network.`,
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
			Type:             "pnap_product_availability",
			Category:         "Data Sources",
			ShortDescription: `Provides a phoenixNAP product availability datasource. This can be used to read product availabilities.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "product_category",
					Description: `Product category. Currently only ` + "`" + `SERVER` + "`" + ` category is supported.`,
				},
				resource.Attribute{
					Name:        "product_code",
					Description: `The code identifying the product. This code has significance across all locations.`,
				},
				resource.Attribute{
					Name:        "show_only_min_quantity_available",
					Description: `Show only locations where product with requested quantity is available or all locations where product is offered. Default value is ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The location code. Currently the following values are allowed: ` + "`" + `PHX` + "`" + `, ` + "`" + `ASH` + "`" + `, ` + "`" + `NLD` + "`" + `, ` + "`" + `SGP` + "`" + `, ` + "`" + `CHI` + "`" + `, ` + "`" + `SEA` + "`" + ` and ` + "`" + `AUS` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "solution",
					Description: `Currently only the following value is allowed: ` + "`" + `SERVER_RANCHER` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "min_quantity",
					Description: `Minimal quantity of product needed. Minimum, maximum and default values might differ for different products. For servers, they are 1, 10 and 1 respectively. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "product_availabilities",
					Description: `List of product availabilities.`,
				},
				resource.Attribute{
					Name:        "product_code",
					Description: `The code identifying the product.`,
				},
				resource.Attribute{
					Name:        "product_category",
					Description: `The product category.`,
				},
				resource.Attribute{
					Name:        "location_availability_details",
					Description: `Infos about location, solutions and availability for a product.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The code identifying the location.`,
				},
				resource.Attribute{
					Name:        "min_quantity_requested",
					Description: `Requested quantity.`,
				},
				resource.Attribute{
					Name:        "min_quantity_available",
					Description: `Is product available in specific location for requested quantity.`,
				},
				resource.Attribute{
					Name:        "available_quantity",
					Description: `Total available quantity of product in specific location. Max value is 10.`,
				},
				resource.Attribute{
					Name:        "solutions",
					Description: `Solutions supported in specific location for a product.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "product_availabilities",
					Description: `List of product availabilities.`,
				},
				resource.Attribute{
					Name:        "product_code",
					Description: `The code identifying the product.`,
				},
				resource.Attribute{
					Name:        "product_category",
					Description: `The product category.`,
				},
				resource.Attribute{
					Name:        "location_availability_details",
					Description: `Infos about location, solutions and availability for a product.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The code identifying the location.`,
				},
				resource.Attribute{
					Name:        "min_quantity_requested",
					Description: `Requested quantity.`,
				},
				resource.Attribute{
					Name:        "min_quantity_available",
					Description: `Is product available in specific location for requested quantity.`,
				},
				resource.Attribute{
					Name:        "available_quantity",
					Description: `Total available quantity of product in specific location. Max value is 10.`,
				},
				resource.Attribute{
					Name:        "solutions",
					Description: `Solutions supported in specific location for a product.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "pnap_products",
			Category:         "Data Sources",
			ShortDescription: `Provides a phoenixNAP products datasource. This can be used to read products.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "product_code",
					Description: `The code identifying the product. This code has significance across all locations.`,
				},
				resource.Attribute{
					Name:        "product_category",
					Description: `The product category.`,
				},
				resource.Attribute{
					Name:        "sku_code",
					Description: `The SKU identifier.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The location code. Currently the following values are allowed: ` + "`" + `PHX` + "`" + `, ` + "`" + `ASH` + "`" + `, ` + "`" + `NLD` + "`" + `, ` + "`" + `SGP` + "`" + `, ` + "`" + `CHI` + "`" + `, ` + "`" + `SEA` + "`" + `, ` + "`" + `AUS` + "`" + ` and ` + "`" + `GLOBAL` + "`" + `. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "products",
					Description: `The list of products recorded.`,
				},
				resource.Attribute{
					Name:        "product_code",
					Description: `The code identifying the product.`,
				},
				resource.Attribute{
					Name:        "product_category",
					Description: `The product category.`,
				},
				resource.Attribute{
					Name:        "plans",
					Description: `The pricing plans available for this product.`,
				},
				resource.Attribute{
					Name:        "sku",
					Description: `The SKU identifying the pricing plan.`,
				},
				resource.Attribute{
					Name:        "sku_description",
					Description: `Description of the pricing plan.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The code identifying the location.`,
				},
				resource.Attribute{
					Name:        "pricing_model",
					Description: `The pricing model.`,
				},
				resource.Attribute{
					Name:        "price",
					Description: `Price per unit.`,
				},
				resource.Attribute{
					Name:        "price_unit",
					Description: `The unit to which the price applies.`,
				},
				resource.Attribute{
					Name:        "correlated_product_code",
					Description: `Product code of the correlated product.`,
				},
				resource.Attribute{
					Name:        "package_quantity",
					Description: `Package size per month.`,
				},
				resource.Attribute{
					Name:        "package_unit",
					Description: `Package size unit.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `Details of the server product.`,
				},
				resource.Attribute{
					Name:        "ram_in_gb",
					Description: `RAM in GB.`,
				},
				resource.Attribute{
					Name:        "cpu",
					Description: `CPU name.`,
				},
				resource.Attribute{
					Name:        "cpu_count",
					Description: `Number of CPUs.`,
				},
				resource.Attribute{
					Name:        "cores_per_cpu",
					Description: `The number of physical cores present on each CPU.`,
				},
				resource.Attribute{
					Name:        "cpu_frequency",
					Description: `CPU frequency in GHz.`,
				},
				resource.Attribute{
					Name:        "network",
					Description: `Server network.`,
				},
				resource.Attribute{
					Name:        "storage",
					Description: `Server storage.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "products",
					Description: `The list of products recorded.`,
				},
				resource.Attribute{
					Name:        "product_code",
					Description: `The code identifying the product.`,
				},
				resource.Attribute{
					Name:        "product_category",
					Description: `The product category.`,
				},
				resource.Attribute{
					Name:        "plans",
					Description: `The pricing plans available for this product.`,
				},
				resource.Attribute{
					Name:        "sku",
					Description: `The SKU identifying the pricing plan.`,
				},
				resource.Attribute{
					Name:        "sku_description",
					Description: `Description of the pricing plan.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The code identifying the location.`,
				},
				resource.Attribute{
					Name:        "pricing_model",
					Description: `The pricing model.`,
				},
				resource.Attribute{
					Name:        "price",
					Description: `Price per unit.`,
				},
				resource.Attribute{
					Name:        "price_unit",
					Description: `The unit to which the price applies.`,
				},
				resource.Attribute{
					Name:        "correlated_product_code",
					Description: `Product code of the correlated product.`,
				},
				resource.Attribute{
					Name:        "package_quantity",
					Description: `Package size per month.`,
				},
				resource.Attribute{
					Name:        "package_unit",
					Description: `Package size unit.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `Details of the server product.`,
				},
				resource.Attribute{
					Name:        "ram_in_gb",
					Description: `RAM in GB.`,
				},
				resource.Attribute{
					Name:        "cpu",
					Description: `CPU name.`,
				},
				resource.Attribute{
					Name:        "cpu_count",
					Description: `Number of CPUs.`,
				},
				resource.Attribute{
					Name:        "cores_per_cpu",
					Description: `The number of physical cores present on each CPU.`,
				},
				resource.Attribute{
					Name:        "cpu_frequency",
					Description: `CPU frequency in GHz.`,
				},
				resource.Attribute{
					Name:        "network",
					Description: `Server network.`,
				},
				resource.Attribute{
					Name:        "storage",
					Description: `Server storage.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "pnap_public_network",
			Category:         "Data Sources",
			ShortDescription: `Provides a phoenixNAP Public Network datasource. This can be used to read public networks.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `The friendly name of this public network.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The public network identifier. ## Attributes Reference The following attributes are exported:`,
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
					Name:        "id",
					Description: `The IP Block identifier.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "pnap_quota",
			Category:         "Data Sources",
			ShortDescription: `Provides a phoenixNAP Quota datasource. This can be used to read Quotas.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `The name of the Quota.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Quota. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Quota.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the Quota.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The Quota description.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the Quota.`,
				},
				resource.Attribute{
					Name:        "limit",
					Description: `The limit set for the Quota.`,
				},
				resource.Attribute{
					Name:        "used",
					Description: `The Quota used expressed as a number.`,
				},
				resource.Attribute{
					Name:        "quota_edit_limit_request_details",
					Description: `List of requests to change the limit on a Quota.`,
				},
				resource.Attribute{
					Name:        "limit",
					Description: `The new limit that is requested.`,
				},
				resource.Attribute{
					Name:        "reason",
					Description: `The reason for changing the limit.`,
				},
				resource.Attribute{
					Name:        "requested_on",
					Description: `The point in time the request was submitted.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Quota.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the Quota.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The Quota description.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the Quota.`,
				},
				resource.Attribute{
					Name:        "limit",
					Description: `The limit set for the Quota.`,
				},
				resource.Attribute{
					Name:        "used",
					Description: `The Quota used expressed as a number.`,
				},
				resource.Attribute{
					Name:        "quota_edit_limit_request_details",
					Description: `List of requests to change the limit on a Quota.`,
				},
				resource.Attribute{
					Name:        "limit",
					Description: `The new limit that is requested.`,
				},
				resource.Attribute{
					Name:        "reason",
					Description: `The reason for changing the limit.`,
				},
				resource.Attribute{
					Name:        "requested_on",
					Description: `The point in time the request was submitted.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "pnap_rancher_cluster",
			Category:         "Data Sources",
			ShortDescription: `Provides a phoenixNAP Rancher Cluster datasource. This can be used to read Rancher Server deployment details.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The cluster (Rancher Cluster) identifier.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Cluster name. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The cluster identifier.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Cluster name.`,
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
					Description: `The node pools associated with the cluster.`,
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
					Description: `Node server type.`,
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
					Description: `Cluster name.`,
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
					Description: `The node pools associated with the cluster.`,
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
					Description: `Node server type.`,
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
					Name:        "status_description",
					Description: `The cluster status.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "pnap_reservation",
			Category:         "Data Sources",
			ShortDescription: `Provides a phoenixNAP reservation datasource. This can be used to read reservation details.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The reservation identifier.`,
				},
				resource.Attribute{
					Name:        "sku",
					Description: `The SKU code of product pricing plan. ## Attributes Reference The following attributes are exported:`,
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
					Description: `The SKU applied to this reservation.`,
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
					Description: `The SKU applied to this reservation.`,
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
			Category:         "Data Sources",
			ShortDescription: `Provides a phoenixNAP server datasource. This can be used to read servers.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "hostname",
					Description: `Server hostname.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique identifier of the server. ## Attributes Reference The following attributes are exported:`,
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
					Name:        "primary_ip_address",
					Description: `First usable public IP Addresses.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `The tags assigned to the server.`,
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
				resource.Attribute{
					Name:        "network_configuration",
					Description: `Entire network details of bare metal server.`,
				},
				resource.Attribute{
					Name:        "gateway_address",
					Description: `The address of the gateway assigned to the server.`,
				},
				resource.Attribute{
					Name:        "private_network_configuration",
					Description: `Private network details of bare metal server.`,
				},
				resource.Attribute{
					Name:        "configuration_type",
					Description: `Determines the approach for configuring private network(s) for the server being provisioned.`,
				},
				resource.Attribute{
					Name:        "private_networks",
					Description: `The list of private networks this server is member of.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The network identifier.`,
				},
				resource.Attribute{
					Name:        "ips",
					Description: `IPs configured on the server.`,
				},
				resource.Attribute{
					Name:        "dhcp",
					Description: `Determines whether DHCP is enabled for this server.`,
				},
				resource.Attribute{
					Name:        "status_description",
					Description: `The status of the network.`,
				},
				resource.Attribute{
					Name:        "ip_blocks_configuration",
					Description: `IP block details of bare metal server.`,
				},
				resource.Attribute{
					Name:        "configuration_type",
					Description: `Determines the approach for configuring IP blocks for the server being provisioned.`,
				},
				resource.Attribute{
					Name:        "ip_blocks",
					Description: `The IP blocks assigned to this server.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The IP block's ID.`,
				},
				resource.Attribute{
					Name:        "ips",
					Description: `The VLAN on which this IP block has been configured within the network switch.`,
				},
				resource.Attribute{
					Name:        "public_network_configuration",
					Description: `Public network details of bare metal server.`,
				},
				resource.Attribute{
					Name:        "public_networks",
					Description: `The list of public networks this server is member of.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The network identifier.`,
				},
				resource.Attribute{
					Name:        "ips",
					Description: `IPs configured on the server.`,
				},
				resource.Attribute{
					Name:        "status_description",
					Description: `The status of the assignment to the network.`,
				},
			},
			Attributes: []resource.Attribute{
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
					Name:        "primary_ip_address",
					Description: `First usable public IP Addresses.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `The tags assigned to the server.`,
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
				resource.Attribute{
					Name:        "network_configuration",
					Description: `Entire network details of bare metal server.`,
				},
				resource.Attribute{
					Name:        "gateway_address",
					Description: `The address of the gateway assigned to the server.`,
				},
				resource.Attribute{
					Name:        "private_network_configuration",
					Description: `Private network details of bare metal server.`,
				},
				resource.Attribute{
					Name:        "configuration_type",
					Description: `Determines the approach for configuring private network(s) for the server being provisioned.`,
				},
				resource.Attribute{
					Name:        "private_networks",
					Description: `The list of private networks this server is member of.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The network identifier.`,
				},
				resource.Attribute{
					Name:        "ips",
					Description: `IPs configured on the server.`,
				},
				resource.Attribute{
					Name:        "dhcp",
					Description: `Determines whether DHCP is enabled for this server.`,
				},
				resource.Attribute{
					Name:        "status_description",
					Description: `The status of the network.`,
				},
				resource.Attribute{
					Name:        "ip_blocks_configuration",
					Description: `IP block details of bare metal server.`,
				},
				resource.Attribute{
					Name:        "configuration_type",
					Description: `Determines the approach for configuring IP blocks for the server being provisioned.`,
				},
				resource.Attribute{
					Name:        "ip_blocks",
					Description: `The IP blocks assigned to this server.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The IP block's ID.`,
				},
				resource.Attribute{
					Name:        "ips",
					Description: `The VLAN on which this IP block has been configured within the network switch.`,
				},
				resource.Attribute{
					Name:        "public_network_configuration",
					Description: `Public network details of bare metal server.`,
				},
				resource.Attribute{
					Name:        "public_networks",
					Description: `The list of public networks this server is member of.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The network identifier.`,
				},
				resource.Attribute{
					Name:        "ips",
					Description: `IPs configured on the server.`,
				},
				resource.Attribute{
					Name:        "status_description",
					Description: `The status of the assignment to the network.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "pnap_ssh_key",
			Category:         "Data Sources",
			ShortDescription: `Provides a phoenixNAP SSH key datasource. This can be used to read SSH keys.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `Friendly SSH key name to represent an SSH key.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique identifier of the SSH Key. ## Attributes Reference The following attributes are exported:`,
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
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "pnap_storage_network",
			Category:         "Data Sources",
			ShortDescription: `Provides a phoenixNAP Storage Network datasource. This can be used to read storage networks.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `The friendly name of this storage network.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The storage network identifier. ## Attributes Reference The following attributes are exported:`,
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
			Category:         "Data Sources",
			ShortDescription: `Provides a phoenixNAP tag datasource. This can be used to read tags.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `The unique name of the tag.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique identifier of the tag. ## Attributes Reference The following attributes are exported:`,
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

	dataSourcesMap = map[string]int{

		"pnap_events":               0,
		"pnap_ip_block":             1,
		"pnap_private_network":      2,
		"pnap_product_availability": 3,
		"pnap_products":             4,
		"pnap_public_network":       5,
		"pnap_quota":                6,
		"pnap_rancher_cluster":      7,
		"pnap_reservation":          8,
		"pnap_server":               9,
		"pnap_ssh_key":              10,
		"pnap_storage_network":      11,
		"pnap_tag":                  12,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
