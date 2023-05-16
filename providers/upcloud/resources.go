package upcloud

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "upcloud_firewall_rules",
			Category:         "Resources",
			ShortDescription: `This resource represents a generated list of UpCloud firewall rules. Firewall rules are used in conjunction with UpCloud servers. Each server has its own firewall rules. The firewall is enabled on all network interfaces except ones attached to private virtual networks. The maximum number of firewall rules per server is 1000.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "upcloud_floating_ip_address",
			Category:         "Resources",
			ShortDescription: `This resource represents a UpCloud floating IP address resource.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "upcloud_gateway",
			Category:         "Resources",
			ShortDescription: `Network gateways connect SDN Private Networks to external IP networks.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "upcloud_kubernetes_cluster",
			Category:         "Resources",
			ShortDescription: `Kubernetes cluster. NOTE: this is an experimental feature in development phase, the resource definition will change in the future.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "upcloud_kubernetes_node_group",
			Category:         "Resources",
			ShortDescription: `Kubernetes node group. NOTE: this is an experimental feature in development phase, the resource definition might change in the future.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "upcloud_loadbalancer",
			Category:         "Resources",
			ShortDescription: `This resource represents load balancer service`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "upcloud_loadbalancer_backend",
			Category:         "Resources",
			ShortDescription: `This resource represents load balancer backend service`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "upcloud_loadbalancer_dynamic_backend_member",
			Category:         "Resources",
			ShortDescription: `This resource represents load balancer's dynamic backend member`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "upcloud_loadbalancer_dynamic_certificate_bundle",
			Category:         "Resources",
			ShortDescription: `This resource represents dynamic certificate bundle`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "upcloud_loadbalancer_frontend",
			Category:         "Resources",
			ShortDescription: `This resource represents load balancer frontend service`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "upcloud_loadbalancer_frontend_rule",
			Category:         "Resources",
			ShortDescription: `This resource represents load balancer frontend rule`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "upcloud_loadbalancer_frontend_tls_config",
			Category:         "Resources",
			ShortDescription: `This resource represents frontend TLS config`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "upcloud_loadbalancer_manual_certificate_bundle",
			Category:         "Resources",
			ShortDescription: `This resource represents manual certificate bundle`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "upcloud_loadbalancer_resolver",
			Category:         "Resources",
			ShortDescription: `This resource represents service's domain name resolver`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "upcloud_loadbalancer_static_backend_member",
			Category:         "Resources",
			ShortDescription: `This resource represents load balancer's static backend member`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "upcloud_managed_database_logical_database",
			Category:         "Resources",
			ShortDescription: `This resource represents a logical database in managed database`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "upcloud_managed_database_mysql",
			Category:         "Resources",
			ShortDescription: `This resource represents MySQL managed database`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "upcloud_managed_database_postgresql",
			Category:         "Resources",
			ShortDescription: `This resource represents PostgreSQL managed database`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "upcloud_managed_database_redis",
			Category:         "Resources",
			ShortDescription: `This resource represents Redis managed database`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "upcloud_managed_database_user",
			Category:         "Resources",
			ShortDescription: `This resource represents a user in managed database`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "upcloud_network",
			Category:         "Resources",
			ShortDescription: `This resource represents an SDN private network that cloud servers from the same zone can be attached to.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "upcloud_object_storage",
			Category:         "Resources",
			ShortDescription: `This resource represents an UpCloud Object Storage instance, which provides S3 compatible storage.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "upcloud_router",
			Category:         "Resources",
			ShortDescription: `This resource represents a generated UpCloud router resource. Routers can be used to connect multiple Private Networks. UpCloud Servers on any attached network can communicate directly with each other.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "upcloud_server",
			Category:         "Resources",
			ShortDescription: `The UpCloud server resource allows the creation, update and deletion of a server.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "upcloud_server_group",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "upcloud_storage",
			Category:         "Resources",
			ShortDescription: `Manages UpCloud storage block devices.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "upcloud_tag",
			Category:         "Resources",
			ShortDescription: `This resource is deprecated, use tags schema in server resource`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"upcloud_firewall_rules":                          0,
		"upcloud_floating_ip_address":                     1,
		"upcloud_gateway":                                 2,
		"upcloud_kubernetes_cluster":                      3,
		"upcloud_kubernetes_node_group":                   4,
		"upcloud_loadbalancer":                            5,
		"upcloud_loadbalancer_backend":                    6,
		"upcloud_loadbalancer_dynamic_backend_member":     7,
		"upcloud_loadbalancer_dynamic_certificate_bundle": 8,
		"upcloud_loadbalancer_frontend":                   9,
		"upcloud_loadbalancer_frontend_rule":              10,
		"upcloud_loadbalancer_frontend_tls_config":        11,
		"upcloud_loadbalancer_manual_certificate_bundle":  12,
		"upcloud_loadbalancer_resolver":                   13,
		"upcloud_loadbalancer_static_backend_member":      14,
		"upcloud_managed_database_logical_database":       15,
		"upcloud_managed_database_mysql":                  16,
		"upcloud_managed_database_postgresql":             17,
		"upcloud_managed_database_redis":                  18,
		"upcloud_managed_database_user":                   19,
		"upcloud_network":                                 20,
		"upcloud_object_storage":                          21,
		"upcloud_router":                                  22,
		"upcloud_server":                                  23,
		"upcloud_server_group":                            24,
		"upcloud_storage":                                 25,
		"upcloud_tag":                                     26,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
