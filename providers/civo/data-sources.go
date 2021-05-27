package civo

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "civo_dns_domain_name",
			Category:         "Data Sources",
			ShortDescription: `Get information on a domain.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The id of the domain.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the domain. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `A unique ID that can be used to identify and reference a domain.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the domain.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `A unique ID that can be used to identify and reference a domain.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the domain.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "civo_dns_domain_record",
			Category:         "Data Sources",
			ShortDescription: `Get information on a DNS record.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the record.`,
				},
				resource.Attribute{
					Name:        "domain_id",
					Description: `(Required) The domain id of the record. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `A unique ID that can be used to identify and reference a Record.`,
				},
				resource.Attribute{
					Name:        "domain_id",
					Description: `The id of the domain`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The choice of record type from A, CNAME, MX, SRV or TXT`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The portion before the domain name (e.g. www) or an @ for the apex/root domain (you cannot use an A record with an amex/root domain)`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The IP address (A or MX), hostname (CNAME or MX) or text value (TXT) to serve for this record`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `The priority of the record.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `How long caching DNS servers should cache this record.`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `The id account of the domain.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `The date when it was created in UTC format`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `The date when it was updated in UTC format`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `A unique ID that can be used to identify and reference a Record.`,
				},
				resource.Attribute{
					Name:        "domain_id",
					Description: `The id of the domain`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The choice of record type from A, CNAME, MX, SRV or TXT`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The portion before the domain name (e.g. www) or an @ for the apex/root domain (you cannot use an A record with an amex/root domain)`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The IP address (A or MX), hostname (CNAME or MX) or text value (TXT) to serve for this record`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `The priority of the record.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `How long caching DNS servers should cache this record.`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `The id account of the domain.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `The date when it was created in UTC format`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `The date when it was updated in UTC format`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "civo_instance",
			Category:         "Data Sources",
			ShortDescription: `Get information on a Instance.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The ID of the Instance`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `(Optional) The hostname of the Instance.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region of an existing Instance. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Instance.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `The Instance hostname.`,
				},
				resource.Attribute{
					Name:        "reverse_dns",
					Description: `A fully qualified domain name.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The name of the size.`,
				},
				resource.Attribute{
					Name:        "cpu_cores",
					Description: `Total cpu of the inatance.`,
				},
				resource.Attribute{
					Name:        "ram_mb",
					Description: `Total ram of the instance.`,
				},
				resource.Attribute{
					Name:        "disk_gb",
					Description: `The size of the disk.`,
				},
				resource.Attribute{
					Name:        "public_ip_requiered",
					Description: `This should be either false, true or ` + "`" + `move_ip_from:intances_id` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `This will be the ID of the network.`,
				},
				resource.Attribute{
					Name:        "template",
					Description: `The ID for the template to used to build the instance.`,
				},
				resource.Attribute{
					Name:        "initial_user",
					Description: `The name of the initial user created on the server.`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `The notes of the instance.`,
				},
				resource.Attribute{
					Name:        "sshkey_id",
					Description: `The ID SSH.`,
				},
				resource.Attribute{
					Name:        "firewall_id",
					Description: `The ID of the firewall used.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `An optional list of tags`,
				},
				resource.Attribute{
					Name:        "initial_password",
					Description: `Instance initial password`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `The private ip.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `The public ip.`,
				},
				resource.Attribute{
					Name:        "pseudo_ip",
					Description: `Is the ip that is used to route the public ip from the internet to the instance using NAT`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the instance`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The region of the instance`,
				},
				resource.Attribute{
					Name:        "script",
					Description: `the contents of a script uploaded`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `The date of creation of the instance`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Instance.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `The Instance hostname.`,
				},
				resource.Attribute{
					Name:        "reverse_dns",
					Description: `A fully qualified domain name.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The name of the size.`,
				},
				resource.Attribute{
					Name:        "cpu_cores",
					Description: `Total cpu of the inatance.`,
				},
				resource.Attribute{
					Name:        "ram_mb",
					Description: `Total ram of the instance.`,
				},
				resource.Attribute{
					Name:        "disk_gb",
					Description: `The size of the disk.`,
				},
				resource.Attribute{
					Name:        "public_ip_requiered",
					Description: `This should be either false, true or ` + "`" + `move_ip_from:intances_id` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `This will be the ID of the network.`,
				},
				resource.Attribute{
					Name:        "template",
					Description: `The ID for the template to used to build the instance.`,
				},
				resource.Attribute{
					Name:        "initial_user",
					Description: `The name of the initial user created on the server.`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `The notes of the instance.`,
				},
				resource.Attribute{
					Name:        "sshkey_id",
					Description: `The ID SSH.`,
				},
				resource.Attribute{
					Name:        "firewall_id",
					Description: `The ID of the firewall used.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `An optional list of tags`,
				},
				resource.Attribute{
					Name:        "initial_password",
					Description: `Instance initial password`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `The private ip.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `The public ip.`,
				},
				resource.Attribute{
					Name:        "pseudo_ip",
					Description: `Is the ip that is used to route the public ip from the internet to the instance using NAT`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the instance`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The region of the instance`,
				},
				resource.Attribute{
					Name:        "script",
					Description: `the contents of a script uploaded`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `The date of creation of the instance`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "civo_instance_size",
			Category:         "Data Sources",
			ShortDescription: `Get information on a Civo Instance Size.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) Filter the results. The ` + "`" + `filter` + "`" + ` block is documented below.`,
				},
				resource.Attribute{
					Name:        "sort",
					Description: `(Optional) Sort the results. The ` + "`" + `sort` + "`" + ` block is documented below. ` + "`" + `filter` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) Filter the sizes by this key. This may be one of ` + "`" + `name` + "`" + `, ` + "`" + `type` + "`" + `, ` + "`" + `cpu` + "`" + `, ` + "`" + `ram` + "`" + `, ` + "`" + `disk` + "`" + `, ` + "`" + `selectable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) Only retrieves images which keys has value that matches one of the values provided here. ` + "`" + `sort` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) Sort the sizes by this key. This may be one of ` + "`" + `name` + "`" + `, ` + "`" + `type` + "`" + `, ` + "`" + `cpu` + "`" + `, ` + "`" + `ram` + "`" + `, ` + "`" + `disk` + "`" + `, ` + "`" + `selectable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "direction",
					Description: `(Required) The sort direction. This may be either ` + "`" + `asc` + "`" + ` or ` + "`" + `desc` + "`" + `. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "cpu",
					Description: `Total of CPU in the instance.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A description of the instance size.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "cpu",
					Description: `Total of CPU in the instance.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A description of the instance size.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "civo_instances",
			Category:         "Data Sources",
			ShortDescription: `Retrieve information on Instances.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) If is used, them all instances will be from that region.`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) Filter the results. The ` + "`" + `filter` + "`" + ` block is documented below.`,
				},
				resource.Attribute{
					Name:        "sort",
					Description: `(Optional) Sort the results. The ` + "`" + `sort` + "`" + ` block is documented below. ` + "`" + `filter` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) Filter the Instances by this key. This may be one of '` + "`" + `id` + "`" + `, ` + "`" + `hostname` + "`" + `, ` + "`" + `public_ip` + "`" + `, ` + "`" + `private_ip` + "`" + `, ` + "`" + `pseudo_ip` + "`" + `, ` + "`" + `size` + "`" + `, ` + "`" + `cpu_cores` + "`" + `, ` + "`" + `ram_mb` + "`" + `, ` + "`" + `disk_gb` + "`" + `, ` + "`" + `template` + "`" + ` or ` + "`" + `created_at` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) A list of values to match against the ` + "`" + `key` + "`" + ` field. Only retrieves Instances where the ` + "`" + `key` + "`" + ` field takes on one or more of the values provided here. ` + "`" + `sort` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) Sort the Instance by this key. This may be one of ` + "`" + `id` + "`" + `, ` + "`" + `hostname` + "`" + `, ` + "`" + `public_ip` + "`" + `, ` + "`" + `private_ip` + "`" + `, ` + "`" + `pseudo_ip` + "`" + `, ` + "`" + `size` + "`" + `, ` + "`" + `cpu_cores` + "`" + `, ` + "`" + `ram_mb` + "`" + `, ` + "`" + `disk_gb` + "`" + `, ` + "`" + `template` + "`" + ` or ` + "`" + `created_at` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "direction",
					Description: `(Required) The sort direction. This may be either ` + "`" + `asc` + "`" + ` or ` + "`" + `desc` + "`" + `. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "instances",
					Description: `A list of Instances satisfying any ` + "`" + `filter` + "`" + ` and ` + "`" + `sort` + "`" + ` criteria. Each instance has the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Instance.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `The Instance hostname.`,
				},
				resource.Attribute{
					Name:        "reverse_dns",
					Description: `A fully qualified domain name.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The name of the size.`,
				},
				resource.Attribute{
					Name:        "cpu_cores",
					Description: `Total cpu of the inatance.`,
				},
				resource.Attribute{
					Name:        "ram_mb",
					Description: `Total ram of the instance.`,
				},
				resource.Attribute{
					Name:        "disk_gb",
					Description: `The size of the disk.`,
				},
				resource.Attribute{
					Name:        "public_ip_requiered",
					Description: `This should be either false, true or ` + "`" + `move_ip_from:intances_id` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `This will be the ID of the network.`,
				},
				resource.Attribute{
					Name:        "template",
					Description: `The ID for the template to used to build the instance.`,
				},
				resource.Attribute{
					Name:        "initial_user",
					Description: `The name of the initial user created on the server.`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `The notes of the instance.`,
				},
				resource.Attribute{
					Name:        "sshkey_id",
					Description: `The ID SSH.`,
				},
				resource.Attribute{
					Name:        "firewall_id",
					Description: `The ID of the firewall used.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `An optional list of tags`,
				},
				resource.Attribute{
					Name:        "initial_password",
					Description: `Instance initial password`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `The private ip.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `The public ip.`,
				},
				resource.Attribute{
					Name:        "pseudo_ip",
					Description: `Is the ip that is used to route the public ip from the internet to the instance using NAT`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the instance`,
				},
				resource.Attribute{
					Name:        "script",
					Description: `the contents of a script uploaded`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `The date of creation of the instance`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "instances",
					Description: `A list of Instances satisfying any ` + "`" + `filter` + "`" + ` and ` + "`" + `sort` + "`" + ` criteria. Each instance has the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Instance.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `The Instance hostname.`,
				},
				resource.Attribute{
					Name:        "reverse_dns",
					Description: `A fully qualified domain name.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The name of the size.`,
				},
				resource.Attribute{
					Name:        "cpu_cores",
					Description: `Total cpu of the inatance.`,
				},
				resource.Attribute{
					Name:        "ram_mb",
					Description: `Total ram of the instance.`,
				},
				resource.Attribute{
					Name:        "disk_gb",
					Description: `The size of the disk.`,
				},
				resource.Attribute{
					Name:        "public_ip_requiered",
					Description: `This should be either false, true or ` + "`" + `move_ip_from:intances_id` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `This will be the ID of the network.`,
				},
				resource.Attribute{
					Name:        "template",
					Description: `The ID for the template to used to build the instance.`,
				},
				resource.Attribute{
					Name:        "initial_user",
					Description: `The name of the initial user created on the server.`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `The notes of the instance.`,
				},
				resource.Attribute{
					Name:        "sshkey_id",
					Description: `The ID SSH.`,
				},
				resource.Attribute{
					Name:        "firewall_id",
					Description: `The ID of the firewall used.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `An optional list of tags`,
				},
				resource.Attribute{
					Name:        "initial_password",
					Description: `Instance initial password`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `The private ip.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `The public ip.`,
				},
				resource.Attribute{
					Name:        "pseudo_ip",
					Description: `Is the ip that is used to route the public ip from the internet to the instance using NAT`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the instance`,
				},
				resource.Attribute{
					Name:        "script",
					Description: `the contents of a script uploaded`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `The date of creation of the instance`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "civo_kubernetes_cluster",
			Category:         "Data Sources",
			ShortDescription: `Get a Civo Kubernetes cluster resource.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The ID of the kubernetes Cluster`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the kubernetes Cluster. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `A unique ID that can be used to identify and reference a Kubernetes cluster.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of your cluster,.`,
				},
				resource.Attribute{
					Name:        "num_target_nodes",
					Description: `The size of the Kubernetes cluster.`,
				},
				resource.Attribute{
					Name:        "target_nodes_size",
					Description: `The size of each node.`,
				},
				resource.Attribute{
					Name:        "kubernetes_version",
					Description: `The version of Kubernetes.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A list of tags.`,
				},
				resource.Attribute{
					Name:        "applications",
					Description: `A list of application installed.`,
				},
				resource.Attribute{
					Name:        "instances",
					Description: `In addition to the arguments provided, these additional attributes about the cluster's default node instance are exported: - ` + "`" + `hostname` + "`" + ` - The hostname of the instance. - ` + "`" + `cpu_cores` + "`" + ` - Total cpu of the inatance. - ` + "`" + `ram_mb` + "`" + ` - Total ram of the instance - ` + "`" + `disk_gb` + "`" + ` - The size of the disk. - ` + "`" + `status` + "`" + ` - The status of the instance. - ` + "`" + `tags` + "`" + ` - The tag of the instances`,
				},
				resource.Attribute{
					Name:        "pools",
					Description: `A list of node pools associated with the cluster. Each node pool exports the following attributes: - ` + "`" + `id` + "`" + ` - The ID of the pool - ` + "`" + `count` + "`" + ` - The size of the pool - ` + "`" + `size` + "`" + ` - The size of each node inside the pool - ` + "`" + `instance_names` + "`" + ` - A list of the instance in the pool`,
				},
				resource.Attribute{
					Name:        "instances",
					Description: `A list of instance inside the pool - ` + "`" + `hostname` + "`" + ` - The hostname of the instance. - ` + "`" + `size` + "`" + ` - The size of the instance. - ` + "`" + `cpu_cores` + "`" + ` - Total cpu of the inatance. - ` + "`" + `ram_mb` + "`" + ` - Total ram of the instance - ` + "`" + `disk_gb` + "`" + ` - The size of the disk. - ` + "`" + `status` + "`" + ` - The status of the instance. - ` + "`" + `tags` + "`" + ` - The tag of the instances`,
				},
				resource.Attribute{
					Name:        "installed_applications",
					Description: `A unique ID that can be used to identify and reference a Kubernetes cluster. - ` + "`" + `application` + "`" + ` - The name of the application - ` + "`" + `version` + "`" + ` - The version of the application - ` + "`" + `installed` + "`" + ` - if installed or not - ` + "`" + `category` + "`" + ` - The category of the application`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of Kubernetes cluster.`,
				},
				resource.Attribute{
					Name:        "kubeconfig",
					Description: `A representation of the Kubernetes cluster's kubeconfig in yaml format.`,
				},
				resource.Attribute{
					Name:        "api_endpoint",
					Description: `The base URL of the API server on the Kubernetes master node.`,
				},
				resource.Attribute{
					Name:        "master_ip",
					Description: `The Ip of the Kubernetes master node.`,
				},
				resource.Attribute{
					Name:        "dns_entry",
					Description: `The unique dns entry for the cluster in this case point to the master.`,
				},
				resource.Attribute{
					Name:        "built_at",
					Description: `The date where the Kubernetes cluster was build.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `The date where the Kubernetes cluster was create.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `A unique ID that can be used to identify and reference a Kubernetes cluster.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of your cluster,.`,
				},
				resource.Attribute{
					Name:        "num_target_nodes",
					Description: `The size of the Kubernetes cluster.`,
				},
				resource.Attribute{
					Name:        "target_nodes_size",
					Description: `The size of each node.`,
				},
				resource.Attribute{
					Name:        "kubernetes_version",
					Description: `The version of Kubernetes.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A list of tags.`,
				},
				resource.Attribute{
					Name:        "applications",
					Description: `A list of application installed.`,
				},
				resource.Attribute{
					Name:        "instances",
					Description: `In addition to the arguments provided, these additional attributes about the cluster's default node instance are exported: - ` + "`" + `hostname` + "`" + ` - The hostname of the instance. - ` + "`" + `cpu_cores` + "`" + ` - Total cpu of the inatance. - ` + "`" + `ram_mb` + "`" + ` - Total ram of the instance - ` + "`" + `disk_gb` + "`" + ` - The size of the disk. - ` + "`" + `status` + "`" + ` - The status of the instance. - ` + "`" + `tags` + "`" + ` - The tag of the instances`,
				},
				resource.Attribute{
					Name:        "pools",
					Description: `A list of node pools associated with the cluster. Each node pool exports the following attributes: - ` + "`" + `id` + "`" + ` - The ID of the pool - ` + "`" + `count` + "`" + ` - The size of the pool - ` + "`" + `size` + "`" + ` - The size of each node inside the pool - ` + "`" + `instance_names` + "`" + ` - A list of the instance in the pool`,
				},
				resource.Attribute{
					Name:        "instances",
					Description: `A list of instance inside the pool - ` + "`" + `hostname` + "`" + ` - The hostname of the instance. - ` + "`" + `size` + "`" + ` - The size of the instance. - ` + "`" + `cpu_cores` + "`" + ` - Total cpu of the inatance. - ` + "`" + `ram_mb` + "`" + ` - Total ram of the instance - ` + "`" + `disk_gb` + "`" + ` - The size of the disk. - ` + "`" + `status` + "`" + ` - The status of the instance. - ` + "`" + `tags` + "`" + ` - The tag of the instances`,
				},
				resource.Attribute{
					Name:        "installed_applications",
					Description: `A unique ID that can be used to identify and reference a Kubernetes cluster. - ` + "`" + `application` + "`" + ` - The name of the application - ` + "`" + `version` + "`" + ` - The version of the application - ` + "`" + `installed` + "`" + ` - if installed or not - ` + "`" + `category` + "`" + ` - The category of the application`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of Kubernetes cluster.`,
				},
				resource.Attribute{
					Name:        "kubeconfig",
					Description: `A representation of the Kubernetes cluster's kubeconfig in yaml format.`,
				},
				resource.Attribute{
					Name:        "api_endpoint",
					Description: `The base URL of the API server on the Kubernetes master node.`,
				},
				resource.Attribute{
					Name:        "master_ip",
					Description: `The Ip of the Kubernetes master node.`,
				},
				resource.Attribute{
					Name:        "dns_entry",
					Description: `The unique dns entry for the cluster in this case point to the master.`,
				},
				resource.Attribute{
					Name:        "built_at",
					Description: `The date where the Kubernetes cluster was build.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `The date where the Kubernetes cluster was create.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "civo_kubernetes_version",
			Category:         "Data Sources",
			ShortDescription: `Get available Civo Kubernetes versions.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) Filter the results. The ` + "`" + `filter` + "`" + ` block is documented below.`,
				},
				resource.Attribute{
					Name:        "sort",
					Description: `(Optional) Sort the results. The ` + "`" + `sort` + "`" + ` block is documented below. ` + "`" + `filter` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) Filter the sizes by this key. This may be one of ` + "`" + `version` + "`" + `, ` + "`" + `label` + "`" + `, ` + "`" + `type` + "`" + `, ` + "`" + `default` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) Only retrieves the version which keys has value that matches one of the values provided here. ` + "`" + `sort` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) Sort the sizes by this key. This may be one of ` + "`" + `version` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "direction",
					Description: `(Required) The sort direction. This may be either ` + "`" + `asc` + "`" + ` or ` + "`" + `desc` + "`" + `. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `A version of the kubernetes.`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `The label of this version.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of the version can be ` + "`" + `stable` + "`" + `, ` + "`" + `legacy` + "`" + ` etc...`,
				},
				resource.Attribute{
					Name:        "default",
					Description: `If is the default version used in all cluster.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "version",
					Description: `A version of the kubernetes.`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `The label of this version.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of the version can be ` + "`" + `stable` + "`" + `, ` + "`" + `legacy` + "`" + ` etc...`,
				},
				resource.Attribute{
					Name:        "default",
					Description: `If is the default version used in all cluster.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "civo_loadbalancer",
			Category:         "Data Sources",
			ShortDescription: `Get information on a loadbalancer.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The ID of the Load Balancer.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `(Optional) The hostname of the Load Balancer. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Load Balancer`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `The hostname of the Load Balancer`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `The protocol used`,
				},
				resource.Attribute{
					Name:        "tls_certificate",
					Description: `If is set will be returned`,
				},
				resource.Attribute{
					Name:        "tls_key",
					Description: `If is set will be returned`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `The port set in the configuration`,
				},
				resource.Attribute{
					Name:        "max_request_size",
					Description: `The max request size set in the configuration`,
				},
				resource.Attribute{
					Name:        "policy",
					Description: `The policy set in the Load Balancer`,
				},
				resource.Attribute{
					Name:        "health_check_path",
					Description: `The path to check the health of the backend`,
				},
				resource.Attribute{
					Name:        "fail_timeout",
					Description: `The wait time until the backend is marked as a failure`,
				},
				resource.Attribute{
					Name:        "max_conns",
					Description: `How many concurrent connections can each backend handle`,
				},
				resource.Attribute{
					Name:        "ignore_invalid_backend_tls",
					Description: `Should self-signed/invalid certificates be ignored from the backend servers`,
				},
				resource.Attribute{
					Name:        "backend",
					Description: `A list of backend instances - ` + "`" + `instance_id` + "`" + ` - The instance id - ` + "`" + `protocol` + "`" + ` - The protocol used in the configuration. - ` + "`" + `port` + "`" + ` - The port set in the configuration.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Load Balancer`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `The hostname of the Load Balancer`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `The protocol used`,
				},
				resource.Attribute{
					Name:        "tls_certificate",
					Description: `If is set will be returned`,
				},
				resource.Attribute{
					Name:        "tls_key",
					Description: `If is set will be returned`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `The port set in the configuration`,
				},
				resource.Attribute{
					Name:        "max_request_size",
					Description: `The max request size set in the configuration`,
				},
				resource.Attribute{
					Name:        "policy",
					Description: `The policy set in the Load Balancer`,
				},
				resource.Attribute{
					Name:        "health_check_path",
					Description: `The path to check the health of the backend`,
				},
				resource.Attribute{
					Name:        "fail_timeout",
					Description: `The wait time until the backend is marked as a failure`,
				},
				resource.Attribute{
					Name:        "max_conns",
					Description: `How many concurrent connections can each backend handle`,
				},
				resource.Attribute{
					Name:        "ignore_invalid_backend_tls",
					Description: `Should self-signed/invalid certificates be ignored from the backend servers`,
				},
				resource.Attribute{
					Name:        "backend",
					Description: `A list of backend instances - ` + "`" + `instance_id` + "`" + ` - The instance id - ` + "`" + `protocol` + "`" + ` - The protocol used in the configuration. - ` + "`" + `port` + "`" + ` - The port set in the configuration.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "civo_network",
			Category:         "Data Sources",
			ShortDescription: `Get information about a Network.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The unique identifier of an existing Network.`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `(Optional) The label of an existing Network.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region of an existing Network. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `A unique ID that can be used to identify and reference a Network.`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `The label used in the configuration.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the network.`,
				},
				resource.Attribute{
					Name:        "default",
					Description: `If is the default network.`,
				},
				resource.Attribute{
					Name:        "cidr",
					Description: `The block ip assigned to the network.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `A unique ID that can be used to identify and reference a Network.`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `The label used in the configuration.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the network.`,
				},
				resource.Attribute{
					Name:        "default",
					Description: `If is the default network.`,
				},
				resource.Attribute{
					Name:        "cidr",
					Description: `The block ip assigned to the network.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "civo_region",
			Category:         "Data Sources",
			ShortDescription: `Get information on a Civo Region.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) Filter the results. The ` + "`" + `filter` + "`" + ` block is documented below.`,
				},
				resource.Attribute{
					Name:        "sort",
					Description: `(Optional) Sort the results. The ` + "`" + `sort` + "`" + ` block is documented below. ` + "`" + `filter` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) Filter the sizes by this key. This may be one of ` + "`" + `code` + "`" + `, ` + "`" + `name` + "`" + `, ` + "`" + `country` + "`" + `, ` + "`" + `default` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) Only retrieves region which keys has value that matches one of the values provided here. ` + "`" + `sort` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) Sort the sizes by this key. This may be one of ` + "`" + `code` + "`" + `,` + "`" + `name` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "direction",
					Description: `(Required) The sort direction. This may be either ` + "`" + `asc` + "`" + ` or ` + "`" + `desc` + "`" + `. ## Attributes Reference The following attributes are exported:`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "civo_snapshot",
			Category:         "Data Sources",
			ShortDescription: `Get information about a Civo snapshot.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The ID of the snapshot.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the snapshot. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the snapshot.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `The ID of the Instance from which the snapshot was be taken.`,
				},
				resource.Attribute{
					Name:        "safe",
					Description: `If is ` + "`" + `true` + "`" + ` the instance will be shut down during the snapshot if id ` + "`" + `false` + "`" + ` them not.`,
				},
				resource.Attribute{
					Name:        "cron_timing",
					Description: `A string with the cron format.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `The hostname of the instance.`,
				},
				resource.Attribute{
					Name:        "template_id",
					Description: `The template id.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The region where the snapshot was take.`,
				},
				resource.Attribute{
					Name:        "size_gb",
					Description: `The size of the snapshot in GB.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The status of the snapshot.`,
				},
				resource.Attribute{
					Name:        "next_execution",
					Description: `if cron was define this date will be the next execution date.`,
				},
				resource.Attribute{
					Name:        "requested_at",
					Description: `The date where the snapshot was requested.`,
				},
				resource.Attribute{
					Name:        "completed_at",
					Description: `The date where the snapshot was completed.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `The name of the snapshot.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `The ID of the Instance from which the snapshot was be taken.`,
				},
				resource.Attribute{
					Name:        "safe",
					Description: `If is ` + "`" + `true` + "`" + ` the instance will be shut down during the snapshot if id ` + "`" + `false` + "`" + ` them not.`,
				},
				resource.Attribute{
					Name:        "cron_timing",
					Description: `A string with the cron format.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `The hostname of the instance.`,
				},
				resource.Attribute{
					Name:        "template_id",
					Description: `The template id.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The region where the snapshot was take.`,
				},
				resource.Attribute{
					Name:        "size_gb",
					Description: `The size of the snapshot in GB.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The status of the snapshot.`,
				},
				resource.Attribute{
					Name:        "next_execution",
					Description: `if cron was define this date will be the next execution date.`,
				},
				resource.Attribute{
					Name:        "requested_at",
					Description: `The date where the snapshot was requested.`,
				},
				resource.Attribute{
					Name:        "completed_at",
					Description: `The date where the snapshot was completed.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "civo_ssh_key",
			Category:         "Data Sources",
			ShortDescription: `Get information on a ssh key.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The ID of the ssh key.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the ssh key. ## Attributes Reference The following attributes are exported:`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "civo_template",
			Category:         "Data Sources",
			ShortDescription: `Get information on a Civo template.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) If is used, them all instances will be from that region.`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) Filter the results. The ` + "`" + `filter` + "`" + ` block is documented below.`,
				},
				resource.Attribute{
					Name:        "sort",
					Description: `(Optional) Sort the results. The ` + "`" + `sort` + "`" + ` block is documented below. ` + "`" + `filter` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) Filter the sizes by this key. This may be one of ` + "`" + `id` + "`" + `,` + "`" + `name` + "`" + `,` + "`" + `version` + "`" + `,` + "`" + `label` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) Only retrieves the template which keys has value that matches one of the values provided here. ` + "`" + `sort` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) Sort the sizes by this key. This may be one of ` + "`" + `id` + "`" + `,` + "`" + `name` + "`" + `,` + "`" + `version` + "`" + `,` + "`" + `label` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "direction",
					Description: `(Required) The sort direction. This may be either ` + "`" + `asc` + "`" + ` or ` + "`" + `desc` + "`" + `. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of the template`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `A short human readable name for the template`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `The version of the template.`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `The label of the template.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id of the template`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `A short human readable name for the template`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `The version of the template.`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `The label of the template.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "civo_volume",
			Category:         "Data Sources",
			ShortDescription: `Get information on a volume.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The unique identifier for the volume.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the volume. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique identifier for the volume.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the volume.`,
				},
				resource.Attribute{
					Name:        "size_gb",
					Description: `The size of the volume.`,
				},
				resource.Attribute{
					Name:        "bootable",
					Description: `if is bootable or not.`,
				},
				resource.Attribute{
					Name:        "mount_point",
					Description: `The mount point of the volume.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `The date of the creation of the volume.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The unique identifier for the volume.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the volume.`,
				},
				resource.Attribute{
					Name:        "size_gb",
					Description: `The size of the volume.`,
				},
				resource.Attribute{
					Name:        "bootable",
					Description: `if is bootable or not.`,
				},
				resource.Attribute{
					Name:        "mount_point",
					Description: `The mount point of the volume.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `The date of the creation of the volume.`,
				},
			},
		},
	}

	dataSourcesMap = map[string]int{

		"civo_dns_domain_name":    0,
		"civo_dns_domain_record":  1,
		"civo_instance":           2,
		"civo_instance_size":      3,
		"civo_instances":          4,
		"civo_kubernetes_cluster": 5,
		"civo_kubernetes_version": 6,
		"civo_loadbalancer":       7,
		"civo_network":            8,
		"civo_region":             9,
		"civo_snapshot":           10,
		"civo_ssh_key":            11,
		"civo_template":           12,
		"civo_volume":             13,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
