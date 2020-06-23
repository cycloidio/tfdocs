package kubernetes

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "kubernetes_all_namespaces",
			Category:         "Data Sources",
			ShortDescription: `Lists all namespaces within a cluster.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "kubernetes_config_map",
			Category:         "Data Sources",
			ShortDescription: `This data source reads configuration data from a config map.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "metadata",
					Description: `(Required) Standard config map's metadata. For more info see [Kubernetes reference](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata) ## Nested Blocks ### ` + "`" + `metadata` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the config map, must be unique. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/identifiers#names)`,
				},
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) Namespace defines the space within which name of the config map must be unique. #### Attributes`,
				},
				resource.Attribute{
					Name:        "generation",
					Description: `A sequence number representing a specific generation of the desired state.`,
				},
				resource.Attribute{
					Name:        "resource_version",
					Description: `An opaque value that represents the internal version of this config map that can be used by clients to determine when config map has changed. For more info see [Kubernetes reference](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency)`,
				},
				resource.Attribute{
					Name:        "self_link",
					Description: `A URL representing this config map.`,
				},
				resource.Attribute{
					Name:        "uid",
					Description: `The unique in time and space value for this config map. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/identifiers#uids) ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "data",
					Description: `A map of the config map data.`,
				},
				resource.Attribute{
					Name:        "binary_data",
					Description: `A map of preserved non-UTF8 data. For more info see [Kubernetes API reference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.17/#configmap-v1-core).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "data",
					Description: `A map of the config map data.`,
				},
				resource.Attribute{
					Name:        "binary_data",
					Description: `A map of preserved non-UTF8 data. For more info see [Kubernetes API reference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.17/#configmap-v1-core).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "kubernetes_ingress",
			Category:         "Data Sources",
			ShortDescription: `Ingress is a collection of rules that allow inbound connections to reach the endpoints defined by a backend. An Ingress can be configured to give services externally-reachable urls, load balance traffic, terminate SSL, offer name based virtual hosting etc.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "metadata",
					Description: `(Required) Standard service's metadata. For more info see [Kubernetes reference](https://github.com/kubernetes/community/blob/e59e666e3464c7d4851136baa8835a311efdfb8e/contributors/devel/api-conventions.md#metadata) ## Nested Blocks ### ` + "`" + `metadata` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the service, must be unique. Cannot be updated. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/identifiers#names)`,
				},
				resource.Attribute{
					Name:        "namespace",
					Description: `(Required) Namespace defines the space within which name of the service must be unique. #### Attributes`,
				},
				resource.Attribute{
					Name:        "annotations",
					Description: `(Optional) An unstructured key value map stored with the service that may be used to store arbitrary metadata. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/annotations)`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) Map of string keys and values that can be used to organize and categorize (scope and select) the service. May match selectors of replication controllers and services. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/labels)`,
				},
				resource.Attribute{
					Name:        "generation",
					Description: `A sequence number representing a specific generation of the desired state.`,
				},
				resource.Attribute{
					Name:        "resource_version",
					Description: `An opaque value that represents the internal version of this service that can be used by clients to determine when service has changed. For more info see [Kubernetes reference](https://github.com/kubernetes/community/blob/e59e666e3464c7d4851136baa8835a311efdfb8e/contributors/devel/api-conventions.md#concurrency-control-and-consistency)`,
				},
				resource.Attribute{
					Name:        "self_link",
					Description: `A URL representing this service.`,
				},
				resource.Attribute{
					Name:        "uid",
					Description: `The unique in time and space value for this service. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/identifiers#uids) ## Attribute Reference ### ` + "`" + `spec` + "`" + ` #### Attributes`,
				},
				resource.Attribute{
					Name:        "backend",
					Description: `Backend defines the referenced service endpoint to which the traffic will be forwarded. See ` + "`" + `backend` + "`" + ` block attributes below.`,
				},
				resource.Attribute{
					Name:        "rule",
					Description: `A list of host rules used to configure the Ingress. If unspecified, or no rule matches, all traffic is sent to the default backend. See ` + "`" + `rule` + "`" + ` block attributes below.`,
				},
				resource.Attribute{
					Name:        "tls",
					Description: `TLS configuration. Currently the Ingress only supports a single TLS port, 443. If multiple members of this list specify different hosts, they will be multiplexed on the same port according to the hostname specified through the SNI TLS extension, if the ingress controller fulfilling the ingress supports SNI. See ` + "`" + `tls` + "`" + ` block attributes below. ### ` + "`" + `backend` + "`" + ` #### Attributes`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `Specifies the name of the referenced service.`,
				},
				resource.Attribute{
					Name:        "service_port",
					Description: `Specifies the port of the referenced service. ### ` + "`" + `rule` + "`" + ` #### Attributes`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `Host is the fully qualified domain name of a network host, as defined by RFC 3986. Note the following deviations from the \"host\" part of the URI as defined in the RFC: 1. IPs are not allowed. Currently an IngressRuleValue can only apply to the IP in the Spec of the parent Ingress. 2. The : delimiter is not respected because ports are not allowed. Currently the port of an Ingress is implicitly :80 for http and :443 for https. Both these may change in the future. Incoming requests are matched against the host before the IngressRuleValue. If the host is unspecified, the Ingress routes all traffic based on the specified IngressRuleValue.`,
				},
				resource.Attribute{
					Name:        "http",
					Description: `http is a list of http selectors pointing to backends. In the example: http:///? -> backend where where parts of the url correspond to RFC 3986, this resource will be used to match against everything after the last '/' and before the first '?' or '#'. See ` + "`" + `http` + "`" + ` block attributes below. #### ` + "`" + `http` + "`" + ``,
				},
				resource.Attribute{
					Name:        "path",
					Description: `Path array of path regex associated with a backend. Incoming urls matching the path are forwarded to the backend, see below for ` + "`" + `path` + "`" + ` block structure. #### ` + "`" + `path` + "`" + ``,
				},
				resource.Attribute{
					Name:        "path",
					Description: `A string or an extended POSIX regular expression as defined by IEEE Std 1003.1, (i.e this follows the egrep/unix syntax, not the perl syntax) matched against the path of an incoming request. Currently it can contain characters disallowed from the conventional \"path\" part of a URL as defined by RFC 3986. Paths must begin with a '/'. If unspecified, the path defaults to a catch all sending traffic to the backend.`,
				},
				resource.Attribute{
					Name:        "backend",
					Description: `Backend defines the referenced service endpoint to which the traffic will be forwarded to. ### ` + "`" + `tls` + "`" + ` #### Attributes`,
				},
				resource.Attribute{
					Name:        "hosts",
					Description: `Hosts are a list of hosts included in the TLS certificate. The values in this list must match the name/s used in the tlsSecret. Defaults to the wildcard host setting for the loadbalancer controller fulfilling this Ingress, if left unspecified.`,
				},
				resource.Attribute{
					Name:        "secret_name",
					Description: `SecretName is the name of the secret used to terminate SSL traffic on 443. Field is left optional to allow SSL routing based on SNI hostname alone. If the SNI host in a listener conflicts with the \"Host\" header field used by an IngressRule, the SNI host is used for termination and value of the Host header is used for routing. ## Attributes`,
				},
				resource.Attribute{
					Name:        "load_balancer_ingress",
					Description: `A list containing ingress points for the load-balancer ### ` + "`" + `load_balancer_ingress` + "`" + ` #### Attributes`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `IP which is set for load-balancer ingress points that are IP based (typically GCE or OpenStack load-balancers)`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `Hostname which is set for load-balancer ingress points that are DNS based (typically AWS load-balancers)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "backend",
					Description: `Backend defines the referenced service endpoint to which the traffic will be forwarded. See ` + "`" + `backend` + "`" + ` block attributes below.`,
				},
				resource.Attribute{
					Name:        "rule",
					Description: `A list of host rules used to configure the Ingress. If unspecified, or no rule matches, all traffic is sent to the default backend. See ` + "`" + `rule` + "`" + ` block attributes below.`,
				},
				resource.Attribute{
					Name:        "tls",
					Description: `TLS configuration. Currently the Ingress only supports a single TLS port, 443. If multiple members of this list specify different hosts, they will be multiplexed on the same port according to the hostname specified through the SNI TLS extension, if the ingress controller fulfilling the ingress supports SNI. See ` + "`" + `tls` + "`" + ` block attributes below. ### ` + "`" + `backend` + "`" + ` #### Attributes`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `Specifies the name of the referenced service.`,
				},
				resource.Attribute{
					Name:        "service_port",
					Description: `Specifies the port of the referenced service. ### ` + "`" + `rule` + "`" + ` #### Attributes`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `Host is the fully qualified domain name of a network host, as defined by RFC 3986. Note the following deviations from the \"host\" part of the URI as defined in the RFC: 1. IPs are not allowed. Currently an IngressRuleValue can only apply to the IP in the Spec of the parent Ingress. 2. The : delimiter is not respected because ports are not allowed. Currently the port of an Ingress is implicitly :80 for http and :443 for https. Both these may change in the future. Incoming requests are matched against the host before the IngressRuleValue. If the host is unspecified, the Ingress routes all traffic based on the specified IngressRuleValue.`,
				},
				resource.Attribute{
					Name:        "http",
					Description: `http is a list of http selectors pointing to backends. In the example: http:///? -> backend where where parts of the url correspond to RFC 3986, this resource will be used to match against everything after the last '/' and before the first '?' or '#'. See ` + "`" + `http` + "`" + ` block attributes below. #### ` + "`" + `http` + "`" + ``,
				},
				resource.Attribute{
					Name:        "path",
					Description: `Path array of path regex associated with a backend. Incoming urls matching the path are forwarded to the backend, see below for ` + "`" + `path` + "`" + ` block structure. #### ` + "`" + `path` + "`" + ``,
				},
				resource.Attribute{
					Name:        "path",
					Description: `A string or an extended POSIX regular expression as defined by IEEE Std 1003.1, (i.e this follows the egrep/unix syntax, not the perl syntax) matched against the path of an incoming request. Currently it can contain characters disallowed from the conventional \"path\" part of a URL as defined by RFC 3986. Paths must begin with a '/'. If unspecified, the path defaults to a catch all sending traffic to the backend.`,
				},
				resource.Attribute{
					Name:        "backend",
					Description: `Backend defines the referenced service endpoint to which the traffic will be forwarded to. ### ` + "`" + `tls` + "`" + ` #### Attributes`,
				},
				resource.Attribute{
					Name:        "hosts",
					Description: `Hosts are a list of hosts included in the TLS certificate. The values in this list must match the name/s used in the tlsSecret. Defaults to the wildcard host setting for the loadbalancer controller fulfilling this Ingress, if left unspecified.`,
				},
				resource.Attribute{
					Name:        "secret_name",
					Description: `SecretName is the name of the secret used to terminate SSL traffic on 443. Field is left optional to allow SSL routing based on SNI hostname alone. If the SNI host in a listener conflicts with the \"Host\" header field used by an IngressRule, the SNI host is used for termination and value of the Host header is used for routing. ## Attributes`,
				},
				resource.Attribute{
					Name:        "load_balancer_ingress",
					Description: `A list containing ingress points for the load-balancer ### ` + "`" + `load_balancer_ingress` + "`" + ` #### Attributes`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `IP which is set for load-balancer ingress points that are IP based (typically GCE or OpenStack load-balancers)`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `Hostname which is set for load-balancer ingress points that are DNS based (typically AWS load-balancers)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "kubernetes_namespace",
			Category:         "Data Sources",
			ShortDescription: `Queries attributes of a Namespace within the cluster.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "metadata",
					Description: `(Required) Standard object metadata. For more info see [Kubernetes reference](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata) ## Nested Blocks ### ` + "`" + `metadata` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the namespace, must be unique. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/identifiers#names) #### Attributes`,
				},
				resource.Attribute{
					Name:        "annotations",
					Description: `(Optional) An unstructured key value map stored with the namespace that may be used to store arbitrary metadata.`,
				},
				resource.Attribute{
					Name:        "generation",
					Description: `A sequence number representing a specific generation of the desired state.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) Map of string keys and values that can be used to organize and categorize (scope and select) namespaces. May match selectors of replication controllers and services.`,
				},
				resource.Attribute{
					Name:        "resource_version",
					Description: `An opaque value that represents the internal version of this namespace that can be used by clients to determine when namespaces have changed. Read more about [concurrency control and consistency](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency).`,
				},
				resource.Attribute{
					Name:        "self_link",
					Description: `A URL representing this namespace.`,
				},
				resource.Attribute{
					Name:        "uid",
					Description: `The unique in time and space value for this namespace. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/identifiers#uids) ### ` + "`" + `spec` + "`" + ` #### Attributes`,
				},
				resource.Attribute{
					Name:        "finalizers",
					Description: `An opaque list of values that must be empty to permanently remove object from storage. More info: https://kubernetes.io/docs/tasks/administer-cluster/namespaces/`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "kubernetes_secret",
			Category:         "Data Sources",
			ShortDescription: `The resource provides mechanisms to inject containers with sensitive information while keeping containers agnostic of Kubernetes.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "metadata",
					Description: `(Required) Standard secret's metadata. For more info see [Kubernetes reference](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata) ## Nested Blocks ### ` + "`" + `metadata` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the secret, must be unique. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/identifiers#names)`,
				},
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) Namespace defines the space within which name of the secret must be unique. #### Attributes`,
				},
				resource.Attribute{
					Name:        "generation",
					Description: `A sequence number representing a specific generation of the desired state.`,
				},
				resource.Attribute{
					Name:        "resource_version",
					Description: `An opaque value that represents the internal version of this secret that can be used by clients to determine when secret has changed. For more info see [Kubernetes reference](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency)`,
				},
				resource.Attribute{
					Name:        "self_link",
					Description: `A URL representing this secret.`,
				},
				resource.Attribute{
					Name:        "uid",
					Description: `The unique in time and space value for this secret. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/identifiers#uids) ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "data",
					Description: `A map of the secret data.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The secret type. Defaults to ` + "`" + `Opaque` + "`" + `. For more info see [Kubernetes reference](https://github.com/kubernetes/community/blob/c7151dd8dd7e487e96e5ce34c6a416bb3b037609/contributors/design-proposals/auth/secrets.md#proposed-design)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "data",
					Description: `A map of the secret data.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The secret type. Defaults to ` + "`" + `Opaque` + "`" + `. For more info see [Kubernetes reference](https://github.com/kubernetes/community/blob/c7151dd8dd7e487e96e5ce34c6a416bb3b037609/contributors/design-proposals/auth/secrets.md#proposed-design)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "kubernetes_service",
			Category:         "Data Sources",
			ShortDescription: `A Service is an abstraction which defines a logical set of pods and a policy by which to access them - sometimes called a micro-service.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "metadata",
					Description: `(Required) Standard service's metadata. For more info see [Kubernetes reference](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata) ## Attributes`,
				},
				resource.Attribute{
					Name:        "spec",
					Description: `Spec defines the behavior of a service. [Kubernetes reference](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#spec-and-status)`,
				},
				resource.Attribute{
					Name:        "load_balancer_ingress",
					Description: `A list containing ingress points for the load-balancer (only valid if ` + "`" + `type = "LoadBalancer"` + "`" + `) ## Nested Blocks ### ` + "`" + `metadata` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the service, must be unique. Cannot be updated. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/identifiers#names)`,
				},
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) Namespace defines the space within which name of the service must be unique. #### Attributes`,
				},
				resource.Attribute{
					Name:        "annotations",
					Description: `(Optional) An unstructured key value map stored with the service that may be used to store arbitrary metadata. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/annotations)`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) Map of string keys and values that can be used to organize and categorize (scope and select) the service. May match selectors of replication controllers and services. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/labels)`,
				},
				resource.Attribute{
					Name:        "generation",
					Description: `A sequence number representing a specific generation of the desired state.`,
				},
				resource.Attribute{
					Name:        "resource_version",
					Description: `An opaque value that represents the internal version of this service that can be used by clients to determine when service has changed. For more info see [Kubernetes reference](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency)`,
				},
				resource.Attribute{
					Name:        "self_link",
					Description: `A URL representing this service.`,
				},
				resource.Attribute{
					Name:        "uid",
					Description: `The unique in time and space value for this service. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/identifiers#uids) ### ` + "`" + `port` + "`" + ` #### Attributes`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of this port within the service. All ports within the service must have unique names. Optional if only one ServicePort is defined on this service.`,
				},
				resource.Attribute{
					Name:        "node_port",
					Description: `The port on each node on which this service is exposed when ` + "`" + `type` + "`" + ` is ` + "`" + `NodePort` + "`" + ` or ` + "`" + `LoadBalancer` + "`" + `. Usually assigned by the system. If specified, it will be allocated to the service if unused or else creation of the service will fail. Default is to auto-allocate a port if the ` + "`" + `type` + "`" + ` of this service requires one. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/services#type--nodeport)`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `The port that will be exposed by this service.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `The IP protocol for this port. Supports ` + "`" + `TCP` + "`" + ` and ` + "`" + `UDP` + "`" + `. Default is ` + "`" + `TCP` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "target_port",
					Description: `Number or name of the port to access on the pods targeted by the service. Number must be in the range 1 to 65535. This field is ignored for services with ` + "`" + `cluster_ip = "None"` + "`" + `. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/services#defining-a-service) ### ` + "`" + `spec` + "`" + ` #### Attributes`,
				},
				resource.Attribute{
					Name:        "cluster_ip",
					Description: `The IP address of the service. It is usually assigned randomly by the master. If an address is specified manually and is not in use by others, it will be allocated to the service; otherwise, creation of the service will fail. ` + "`" + `None` + "`" + ` can be specified for headless services when proxying is not required. Ignored if type is ` + "`" + `ExternalName` + "`" + `. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/services#virtual-ips-and-service-proxies)`,
				},
				resource.Attribute{
					Name:        "external_ips",
					Description: `A list of IP addresses for which nodes in the cluster will also accept traffic for this service. These IPs are not managed by Kubernetes. The user is responsible for ensuring that traffic arrives at a node with this IP. A common example is external load-balancers that are not part of the Kubernetes system.`,
				},
				resource.Attribute{
					Name:        "external_name",
					Description: `The external reference that kubedns or equivalent will return as a CNAME record for this service. No proxying will be involved. Must be a valid DNS name and requires ` + "`" + `type` + "`" + ` to be ` + "`" + `ExternalName` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "external_traffic_policy",
					Description: `(Optional) Denotes if this Service desires to route external traffic to node-local or cluster-wide endpoints. ` + "`" + `Local` + "`" + ` preserves the client source IP and avoids a second hop for LoadBalancer and Nodeport type services, but risks potentially imbalanced traffic spreading. ` + "`" + `Cluster` + "`" + ` obscures the client source IP and may cause a second hop to another node, but should have good overall load-spreading. More info: https://kubernetes.io/docs/tutorials/services/source-ip/`,
				},
				resource.Attribute{
					Name:        "load_balancer_ip",
					Description: `Only applies to ` + "`" + `type = LoadBalancer` + "`" + `. LoadBalancer will get created with the IP specified in this field. This feature depends on whether the underlying cloud-provider supports specifying this field when a load balancer is created. This field will be ignored if the cloud-provider does not support the feature.`,
				},
				resource.Attribute{
					Name:        "load_balancer_source_ranges",
					Description: `If specified and supported by the platform, this will restrict traffic through the cloud-provider load-balancer will be restricted to the specified client IPs. This field will be ignored if the cloud-provider does not support the feature. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/services-firewalls)`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `The list of ports that are exposed by this service. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/services#virtual-ips-and-service-proxies)`,
				},
				resource.Attribute{
					Name:        "selector",
					Description: `Route service traffic to pods with label keys and values matching this selector. Only applies to types ` + "`" + `ClusterIP` + "`" + `, ` + "`" + `NodePort` + "`" + `, and ` + "`" + `LoadBalancer` + "`" + `. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/services#overview)`,
				},
				resource.Attribute{
					Name:        "session_affinity",
					Description: `Used to maintain session affinity. Supports ` + "`" + `ClientIP` + "`" + ` and ` + "`" + `None` + "`" + `. Defaults to ` + "`" + `None` + "`" + `. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/services#virtual-ips-and-service-proxies)`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Determines how the service is exposed. Defaults to ` + "`" + `ClusterIP` + "`" + `. Valid options are ` + "`" + `ExternalName` + "`" + `, ` + "`" + `ClusterIP` + "`" + `, ` + "`" + `NodePort` + "`" + `, and ` + "`" + `LoadBalancer` + "`" + `. ` + "`" + `ExternalName` + "`" + ` maps to the specified ` + "`" + `external_name` + "`" + `. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/services#overview) ### ` + "`" + `load_balancer_ingress` + "`" + ` #### Attributes`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `Hostname which is set for load-balancer ingress points that are DNS based (typically AWS load-balancers)`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `IP which is set for load-balancer ingress points that are IP based (typically GCE or OpenStack load-balancers)`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "kubernetes_service_account",
			Category:         "Data Sources",
			ShortDescription: `A service account provides an identity for processes that run in a Pod.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "metadata",
					Description: `(Required) Standard service account's metadata. For more info see [Kubernetes reference](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata) ## Nested Blocks ### ` + "`" + `metadata` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the service account, must be unique. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/identifiers#names)`,
				},
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) Namespace defines the space within which name of the service account must be unique. #### Attributes`,
				},
				resource.Attribute{
					Name:        "generation",
					Description: `A sequence number representing a specific generation of the desired state.`,
				},
				resource.Attribute{
					Name:        "resource_version",
					Description: `An opaque value that represents the internal version of this service account that can be used by clients to determine when service account has changed. For more info see [Kubernetes reference](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency)`,
				},
				resource.Attribute{
					Name:        "self_link",
					Description: `A URL representing this service account.`,
				},
				resource.Attribute{
					Name:        "uid",
					Description: `The unique in time and space value for this service account. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/identifiers#uids) ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "image_pull_secret",
					Description: `A list of image pull secrets associated with the service account.`,
				},
				resource.Attribute{
					Name:        "secret",
					Description: `A list of secrets associated with the service account.`,
				},
				resource.Attribute{
					Name:        "default_secret_name",
					Description: `Name of the default secret, containing service account token, created & managed by the service. ### ` + "`" + `image_pull_secret` + "`" + ` #### Attributes`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the referent. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/identifiers#names) ### ` + "`" + `secret` + "`" + ` #### Attributes`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the referent. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/identifiers#names)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "image_pull_secret",
					Description: `A list of image pull secrets associated with the service account.`,
				},
				resource.Attribute{
					Name:        "secret",
					Description: `A list of secrets associated with the service account.`,
				},
				resource.Attribute{
					Name:        "default_secret_name",
					Description: `Name of the default secret, containing service account token, created & managed by the service. ### ` + "`" + `image_pull_secret` + "`" + ` #### Attributes`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the referent. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/identifiers#names) ### ` + "`" + `secret` + "`" + ` #### Attributes`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the referent. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/identifiers#names)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "kubernetes_storage_class",
			Category:         "Data Sources",
			ShortDescription: `Storage class is the foundation of dynamic provisioning, allowing cluster administrators to define abstractions for the underlying storage platform.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "metadata",
					Description: `(Required) Standard storage class's metadata. For more info see [Kubernetes reference](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata) ## Nested Blocks ### ` + "`" + `metadata` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the storage class, must be unique. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/identifiers#names) #### Attributes`,
				},
				resource.Attribute{
					Name:        "generation",
					Description: `A sequence number representing a specific generation of the desired state.`,
				},
				resource.Attribute{
					Name:        "resource_version",
					Description: `An opaque value that represents the internal version of this storage class that can be used by clients to determine when storage class has changed. For more info see [Kubernetes reference](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency)`,
				},
				resource.Attribute{
					Name:        "self_link",
					Description: `A URL representing this storage class.`,
				},
				resource.Attribute{
					Name:        "uid",
					Description: `The unique in time and space value for this storage class. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/identifiers#uids) ## Argument Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "parameters",
					Description: `The parameters for the provisioner that creates volume of this storage class. Read more about [available parameters](https://kubernetes.io/docs/concepts/storage/persistent-volumes/#parameters).`,
				},
				resource.Attribute{
					Name:        "storage_provisioner",
					Description: `Indicates the type of the provisioner this storage class represents`,
				},
				resource.Attribute{
					Name:        "reclaim_policy",
					Description: `Indicates the reclaim policy used.`,
				},
				resource.Attribute{
					Name:        "volume_binding_mode",
					Description: `Indicates when volume binding and dynamic provisioning should occur.`,
				},
				resource.Attribute{
					Name:        "allow_volume_expansion",
					Description: `Indicates whether the storage class allow volume expand.`,
				},
				resource.Attribute{
					Name:        "mount_options",
					Description: `Persistent Volumes that are dynamically created by a storage class will have the mount options specified.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
	}

	dataSourcesMap = map[string]int{

		"kubernetes_all_namespaces":  0,
		"kubernetes_config_map":      1,
		"kubernetes_ingress":         2,
		"kubernetes_namespace":       3,
		"kubernetes_secret":          4,
		"kubernetes_service":         5,
		"kubernetes_service_account": 6,
		"kubernetes_storage_class":   7,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
