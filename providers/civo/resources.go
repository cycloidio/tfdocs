package civo

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "civo_dns_domain_name",
			Category:         "Resources",
			ShortDescription: `Provides a Civo dns domain name resource.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the domain ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `A unique ID that can be used to identify and reference a domain.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the domain.`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `The id account of the domain ## Import Domains can be imported using the ` + "`" + `domain name` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import civo_dns_domain_name.main mydomain.com ` + "`" + `` + "`" + `` + "`" + ``,
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
				resource.Attribute{
					Name:        "account_id",
					Description: `The id account of the domain ## Import Domains can be imported using the ` + "`" + `domain name` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import civo_dns_domain_name.main mydomain.com ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "civo_dns_domain_record",
			Category:         "Resources",
			ShortDescription: `Provides a Civo dns domain record resource.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "domain_id",
					Description: `(Required) The id of the domain`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) The choice of record type from A, CNAME, MX, SRV or TXT`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The portion before the domain name (e.g. www) or an @ for the apex/root domain (you cannot use an A record with an amex/root domain)`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) The IP address (A or MX), hostname (CNAME or MX) or text value (TXT) to serve for this record`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `(Optional) Useful for MX records only, the priority mail should be attempted it (defaults to 10)`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `(Required) How long caching DNS servers should cache this record for, in seconds (the minimum is 600 and the default if unspecified is 600) ## Attributes Reference The following attributes are exported including the arguments:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `A unique ID that can be used to identify and reference a Record.`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `The id account of the domain`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `The date when it was created in UTC format`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `The date when it was updated in UTC format ## Import Domains can be imported using the ` + "`" + `id_domain:id_domain_record` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import civo_dns_domain_record.www a3cd6832-9577-4017-afd7-17d239fc0bf0:c9a39d14-ee1b-4870-8fb0-a2d4f465e822 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `A unique ID that can be used to identify and reference a Record.`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `The id account of the domain`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `The date when it was created in UTC format`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `The date when it was updated in UTC format ## Import Domains can be imported using the ` + "`" + `id_domain:id_domain_record` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import civo_dns_domain_record.www a3cd6832-9577-4017-afd7-17d239fc0bf0:c9a39d14-ee1b-4870-8fb0-a2d4f465e822 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "civo_firewall",
			Category:         "Resources",
			ShortDescription: `Provides a Civo Cloud Firewall resource. This can be used to create, modify, and delete Firewalls.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The Firewall name ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `A unique ID that can be used to identify and reference a Firewall.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the Firewall.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The region where the firewall was create. ## Import Firewalls can be imported using the firewall ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import civo_firewall.www b8ecd2ab-2267-4a5e-8692-cbf1d32583e3 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `A unique ID that can be used to identify and reference a Firewall.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the Firewall.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The region where the firewall was create. ## Import Firewalls can be imported using the firewall ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import civo_firewall.www b8ecd2ab-2267-4a5e-8692-cbf1d32583e3 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "civo_firewall_rule",
			Category:         "Resources",
			ShortDescription: `Provides a Civo Cloud Firewall Rule resource. This can be used to create, modify, and delete Firewalls Rules.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "firewall_id",
					Description: `(Required) The Firewall id`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `A unique ID that can be used to identify and reference a Firewall Rule.`,
				},
				resource.Attribute{
					Name:        "firewall_id",
					Description: `The Firewall id`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `A unique ID that can be used to identify and reference a Firewall Rule.`,
				},
				resource.Attribute{
					Name:        "firewall_id",
					Description: `The Firewall id`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "civo_instance",
			Category:         "Resources",
			ShortDescription: `Provides a Civo Instance resource. This can be used to create, modify, and delete Instances.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "hostname",
					Description: `(Required) The Instance hostname.`,
				},
				resource.Attribute{
					Name:        "reverse_dns",
					Description: `(Optional) A fully qualified domain name that should be used as the instance's IP's reverse DNS (optional, uses the hostname if unspecified).`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `(Optional) The name of the size, from the current list, e.g. g2.small (required).`,
				},
				resource.Attribute{
					Name:        "public_ip_required",
					Description: `(Optional) This should be either false, true or ` + "`" + `move_ip_from:intances_id` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `(Optional) This must be the ID of the network from the network listing (optional; default network used when not specified).`,
				},
				resource.Attribute{
					Name:        "template",
					Description: `(Optional) The ID for the template to use to build the instance.`,
				},
				resource.Attribute{
					Name:        "initial_user",
					Description: `(Optional) The name of the initial user created on the server (optional; this will default to the template's default_username and fallback to civo).`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `(Optional) Add some notes to the instance.`,
				},
				resource.Attribute{
					Name:        "sshkey_id",
					Description: `(Optional) The ID of an already uploaded SSH public key to use for login to the default user (optional; if one isn't provided a random password will be set and returned in the initial_password field).`,
				},
				resource.Attribute{
					Name:        "firewall_id",
					Description: `(Optional) The ID of the firewall to use, from the current list. If left blank or not sent, the default firewall will be used (open to all).`,
				},
				resource.Attribute{
					Name:        "script",
					Description: `(Optional) the contents of a script that will be uploaded to /usr/local/bin/civo-user-init-script on your instance, read/write/executable only by root and then will be executed at the end of the cloud initialization`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) An optional list of tags, represented as a key, value pair. ## Attributes Reference The following attributes are exported:`,
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
					Description: `The date of creation of the instance ## Import Instances can be imported using the instance ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import civo_instance.myintance 18bd98ad-1b6e-4f87-b48f-e690b4fd7413 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
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
					Description: `The date of creation of the instance ## Import Instances can be imported using the instance ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import civo_instance.myintance 18bd98ad-1b6e-4f87-b48f-e690b4fd7413 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "civo_kubernetes_cluster",
			Category:         "Resources",
			ShortDescription: `Provides a Civo Kubernetes cluster resource.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A name for the Kubernetes cluster.`,
				},
				resource.Attribute{
					Name:        "num_target_nodes",
					Description: `(Optional) The number of instances to create (The default at the time of writing is 3).`,
				},
				resource.Attribute{
					Name:        "target_nodes_size",
					Description: `(Optional) The size of each node (The default is currently g2.small)`,
				},
				resource.Attribute{
					Name:        "kubernetes_version",
					Description: `(Optional) The version of k3s to install (The default is currently the latest available).`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A space separated list of tags, to be used freely as required.`,
				},
				resource.Attribute{
					Name:        "applications",
					Description: `(Optional) A comma separated list of applications to install. Spaces within application names are fine, but shouldn't be either side of the comma. If you want to remove a default installed application, prefix it with a '-', e.g. -traefik ## Attributes Reference In addition to the arguments listed above, the following additional attributes are exported:`,
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
					Description: `In addition to the arguments provided, these additional attributes about the cluster's default node instance are exported. - ` + "`" + `hostname` + "`" + ` - The hostname of the instance. - ` + "`" + `size` + "`" + ` - The size of the instance. - ` + "`" + `cpu_cores` + "`" + ` - Total cpu of the inatance. - ` + "`" + `ram_mb` + "`" + ` - Total ram of the instance. - ` + "`" + `disk_gb` + "`" + ` - The size of the disk. - ` + "`" + `region` + "`" + ` - The region where instance are. - ` + "`" + `status` + "`" + ` - The status of the instance. - ` + "`" + `created_at` + "`" + ` - The date where the instances was created. - ` + "`" + `firewall_id` + "`" + ` - The firewall id assigned to the instance - ` + "`" + `public_ip` + "`" + ` - The public ip of the instances, only available if the instances is the master - ` + "`" + `tags` + "`" + ` - The tag of the instances`,
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
					Description: `The date where the Kubernetes cluster was create. ## Import Then the Kubernetes cluster can be imported using the cluster's ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import civo_kubernetes_cluster.my-cluster 1b8b2100-0e9f-4e8f-ad78-9eb578c2a0af ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `In addition to the arguments provided, these additional attributes about the cluster's default node instance are exported. - ` + "`" + `hostname` + "`" + ` - The hostname of the instance. - ` + "`" + `size` + "`" + ` - The size of the instance. - ` + "`" + `cpu_cores` + "`" + ` - Total cpu of the inatance. - ` + "`" + `ram_mb` + "`" + ` - Total ram of the instance. - ` + "`" + `disk_gb` + "`" + ` - The size of the disk. - ` + "`" + `region` + "`" + ` - The region where instance are. - ` + "`" + `status` + "`" + ` - The status of the instance. - ` + "`" + `created_at` + "`" + ` - The date where the instances was created. - ` + "`" + `firewall_id` + "`" + ` - The firewall id assigned to the instance - ` + "`" + `public_ip` + "`" + ` - The public ip of the instances, only available if the instances is the master - ` + "`" + `tags` + "`" + ` - The tag of the instances`,
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
					Description: `The date where the Kubernetes cluster was create. ## Import Then the Kubernetes cluster can be imported using the cluster's ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import civo_kubernetes_cluster.my-cluster 1b8b2100-0e9f-4e8f-ad78-9eb578c2a0af ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "civo_loadbalance",
			Category:         "Resources",
			ShortDescription: `Provides a Civo Load Balancer resource. This can be used to create, modify, and delete Load Balancers.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "hostname",
					Description: `(Required) The hostname to receive traffic for, e.g. www.example.com (optional: sets hostname to loadbalancer-uuid.civo.com if blank)`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Required) Either http or https. If you specify https then you must also provide the next two fields, the default is http",`,
				},
				resource.Attribute{
					Name:        "tls_certificate",
					Description: `(Optional) If your protocol is https then you should send the TLS certificate in Base64-encoded PEM format`,
				},
				resource.Attribute{
					Name:        "tls_key",
					Description: `(Optional) If your protocol is https then you should send the TLS private key in Base64-encoded PEM format`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Required) You can listen on any port, the default is 80 to match the default protocol of http, if not you must specify it here (commonly 80 for HTTP or 443 for HTTPS)`,
				},
				resource.Attribute{
					Name:        "max_request_size",
					Description: `(Required) The size in megabytes of the maximum request content that will be accepted`,
				},
				resource.Attribute{
					Name:        "policy",
					Description: `(Required) One of: ` + "`" + `least_conn` + "`" + ` (sends new requests to the least busy server) ` + "`" + `random` + "`" + ` (sends new requests to a random backend), ` + "`" + `round_robin` + "`" + ` (sends new requests to the next backend in order), ` + "`" + `ip_hash` + "`" + ` (sends requests from a given IP address to the same backend), default is ` + "`" + `random` + "`" + ``,
				},
				resource.Attribute{
					Name:        "health_check_path",
					Description: `(Optional) What URL should be used on the backends to determine if it's OK (2xx/3xx status), defaults to /`,
				},
				resource.Attribute{
					Name:        "fail_timeout",
					Description: `(Required) How long to wait in seconds before determining a backend has failed, defaults to 30.`,
				},
				resource.Attribute{
					Name:        "backend",
					Description: `(Required) A list of backend instances, each containing an ` + "`" + `instance_id` + "`" + `, ` + "`" + `protocol` + "`" + ` (http or https) and ` + "`" + `port` + "`" + `. - ` + "`" + `instance_id` + "`" + ` - (Required) - The instance id - ` + "`" + `protocol` + "`" + ` - (Required) - The protocol Either http or https. - ` + "`" + `port` + "`" + ` - (Required) - You can listen on any port. ## Attributes Reference The following attributes are exported:`,
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
					Description: `A list of backend instances - ` + "`" + `instance_id` + "`" + ` - The instance id - ` + "`" + `protocol` + "`" + ` - The protocol used in the configuration. - ` + "`" + `port` + "`" + ` - The port set in the configuration. ## Import Load Balancers can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import civo_loadbalancer.myloadbalancer 4de7ac8b-495b-4884-9a69-1050c6793cd6 ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `A list of backend instances - ` + "`" + `instance_id` + "`" + ` - The instance id - ` + "`" + `protocol` + "`" + ` - The protocol used in the configuration. - ` + "`" + `port` + "`" + ` - The port set in the configuration. ## Import Load Balancers can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import civo_loadbalancer.myloadbalancer 4de7ac8b-495b-4884-9a69-1050c6793cd6 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "civo_network",
			Category:         "Resources",
			ShortDescription: `Provides a Civo Network resource. This can be used to create, modify, and delete Networks.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "label",
					Description: `(Required) The Network label ## Attributes Reference The following attributes are exported:`,
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
					Name:        "region",
					Description: `The region where the network was create.`,
				},
				resource.Attribute{
					Name:        "default",
					Description: `If is the default network.`,
				},
				resource.Attribute{
					Name:        "cidr",
					Description: `The block ip assigned to the network. ## Import Firewalls can be imported using the firewall ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import civo_network.custom_net b8ecd2ab-2267-4a5e-8692-cbf1d32583e3 ` + "`" + `` + "`" + `` + "`" + ``,
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
					Name:        "region",
					Description: `The region where the network was create.`,
				},
				resource.Attribute{
					Name:        "default",
					Description: `If is the default network.`,
				},
				resource.Attribute{
					Name:        "cidr",
					Description: `The block ip assigned to the network. ## Import Firewalls can be imported using the firewall ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import civo_network.custom_net b8ecd2ab-2267-4a5e-8692-cbf1d32583e3 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "civo_snapshot",
			Category:         "Resources",
			ShortDescription: `Provides a Civo Instance snapshot resource.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A name for the instance snapshot.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Required) The ID of the Instance from which the snapshot will be taken.`,
				},
				resource.Attribute{
					Name:        "safe",
					Description: `(Optional) If ` + "`" + `true` + "`" + ` the instance will be shut down during the snapshot to ensure all files are in a consistent state (e.g. database tables aren't in the middle of being optimised and hence risking corruption). The default is ` + "`" + `false` + "`" + ` so you experience no interruption of service, but a small risk of corruption.`,
				},
				resource.Attribute{
					Name:        "cron_timing",
					Description: `(Optional) If a valid cron string is passed, the snapshot will be saved as an automated snapshot continuing to automatically update based on the schedule of the cron sequence provided The default is nil meaning the snapshot will be saved as a one-off snapshot. ## Attributes Reference The following attributes are exported:`,
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
					Description: `The date where the snapshot was completed. ## Import Instance Snapshots can be imported using the ` + "`" + `snapshot id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import civo_snapshot.myinstance-backup 4cc87851-e1d0-4270-822a-b36d28c7a77f ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `The date where the snapshot was completed. ## Import Instance Snapshots can be imported using the ` + "`" + `snapshot id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import civo_snapshot.myinstance-backup 4cc87851-e1d0-4270-822a-b36d28c7a77f ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "civo_ssh_key",
			Category:         "Resources",
			ShortDescription: `Provides a Civo SSH key resource.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the SSH key for identification`,
				},
				resource.Attribute{
					Name:        "public_key",
					Description: `(Required) The public key. If this is a file, it can be read using the file interpolation function. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique ID of the key`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the SSH key`,
				},
				resource.Attribute{
					Name:        "public_key",
					Description: `The text of the public key`,
				},
				resource.Attribute{
					Name:        "fingerprint",
					Description: `The fingerprint of the SSH key ## Import SSH Keys can be imported using the ` + "`" + `ssh key id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import civo_ssh_key.mykey 87ca2ee4-57d3-4420-b9b6-411b0b4b2a0e ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The unique ID of the key`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the SSH key`,
				},
				resource.Attribute{
					Name:        "public_key",
					Description: `The text of the public key`,
				},
				resource.Attribute{
					Name:        "fingerprint",
					Description: `The fingerprint of the SSH key ## Import SSH Keys can be imported using the ` + "`" + `ssh key id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import civo_ssh_key.mykey 87ca2ee4-57d3-4420-b9b6-411b0b4b2a0e ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "civo_template",
			Category:         "Resources",
			ShortDescription: `Provides a Civo Template resource. This can be used to create, modify, and delete Templates.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "code",
					Description: `(Required) This is a unqiue, alphanumerical, short, human readable code for the template.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) This is a short human readable name for the template`,
				},
				resource.Attribute{
					Name:        "volume_id",
					Description: `(Optional) This is the ID of a bootable volume, either owned by you or global (optional; but must be specified if no image_id is specified)`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `(Optional) This is the Image ID of any default template or the ID of another template either owned by you or global (optional; but must be specified if no volume_id is specified).`,
				},
				resource.Attribute{
					Name:        "short_description",
					Description: `(Optional) A one line description of the template`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A multi-line description of the template, in Markdown format`,
				},
				resource.Attribute{
					Name:        "default_username",
					Description: `(Optional) The default username to suggest that the user creates`,
				},
				resource.Attribute{
					Name:        "cloud_config",
					Description: `(Optional) Commonly referred to as 'user-data', this is a customisation script that is run after the instance is first booted. We recommend using cloud-config as it's a great distribution-agnostic way of configuring cloud servers. If you put ` + "`" + `$INITIAL_USER` + "`" + ` in your script, this will automatically be replaced by the initial user chosen when creating the instance, ` + "`" + `$INITIAL_PASSWORD` + "`" + ` will be replaced with the random password generated by the system, ` + "`" + `$HOSTNAME` + "`" + ` is the fully qualified domain name of the instance and ` + "`" + `$SSH_KEY` + "`" + ` will be the content of the SSH public key. (this is technically optional, but you won't really be able to use instances without it - see our learn guide on templates for more information). ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of the template`,
				},
				resource.Attribute{
					Name:        "code",
					Description: `A unqiue, alphanumerical, short, human readable code for the template.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `A short human readable name for the template`,
				},
				resource.Attribute{
					Name:        "volume_id",
					Description: `The ID of a bootable volume, either owned by you or global.`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `The Image ID of any default template or the ID of another template.`,
				},
				resource.Attribute{
					Name:        "short_description",
					Description: `A one line description of the template`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A multi-line description of the template, in Markdown format`,
				},
				resource.Attribute{
					Name:        "default_username",
					Description: `The default username to suggest that the user creates`,
				},
				resource.Attribute{
					Name:        "cloud_config",
					Description: `Commonly referred to as 'user-data', this is a customisation script that is run after the instance is first booted. ## Import Template can be imported using the template ` + "`" + `code` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import civo_template.my-custom-template my-template-code ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id of the template`,
				},
				resource.Attribute{
					Name:        "code",
					Description: `A unqiue, alphanumerical, short, human readable code for the template.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `A short human readable name for the template`,
				},
				resource.Attribute{
					Name:        "volume_id",
					Description: `The ID of a bootable volume, either owned by you or global.`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `The Image ID of any default template or the ID of another template.`,
				},
				resource.Attribute{
					Name:        "short_description",
					Description: `A one line description of the template`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A multi-line description of the template, in Markdown format`,
				},
				resource.Attribute{
					Name:        "default_username",
					Description: `The default username to suggest that the user creates`,
				},
				resource.Attribute{
					Name:        "cloud_config",
					Description: `Commonly referred to as 'user-data', this is a customisation script that is run after the instance is first booted. ## Import Template can be imported using the template ` + "`" + `code` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import civo_template.my-custom-template my-template-code ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "civo_volume",
			Category:         "Resources",
			ShortDescription: `Provides a Civo volume resource.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A name that you wish to use to refer to this volume .`,
				},
				resource.Attribute{
					Name:        "size_gb",
					Description: `(Required) A minimum of 1 and a maximum of your available disk space from your quota specifies the size of the volume in gigabytes .`,
				},
				resource.Attribute{
					Name:        "bootable",
					Description: `(Required) Mark the volume as bootable. ## Attributes Reference The following attributes are exported:`,
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
					Description: `The date of the creation of the volume. ## Import Volumes can be imported using the ` + "`" + `volume id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import civo_volume.db 506f78a4-e098-11e5-ad9f-000f53306ae1 ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `The date of the creation of the volume. ## Import Volumes can be imported using the ` + "`" + `volume id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import civo_volume.db 506f78a4-e098-11e5-ad9f-000f53306ae1 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "civo_volume_attachment",
			Category:         "Resources",
			ShortDescription: `Provides a Civo volume attachment resource.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Required) ID of the instance to attach the volume to.`,
				},
				resource.Attribute{
					Name:        "volume_id",
					Description: `(Required) ID of the Volume to be attached to the instance. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique identifier for the volume attachment.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The unique identifier for the volume attachment.`,
				},
			},
		},
	}

	resourcesMap = map[string]int{

		"civo_dns_domain_name":    0,
		"civo_dns_domain_record":  1,
		"civo_firewall":           2,
		"civo_firewall_rule":      3,
		"civo_instance":           4,
		"civo_kubernetes_cluster": 5,
		"civo_loadbalance":        6,
		"civo_network":            7,
		"civo_snapshot":           8,
		"civo_ssh_key":            9,
		"civo_template":           10,
		"civo_volume":             11,
		"civo_volume_attachment":  12,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
