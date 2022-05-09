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
					Description: `TLS configuration. Currently the Ingress only supports a single TLS port, 443. If multiple members of this list specify different hosts, they will be multiplexed on the same port according to the hostname specified through the SNI TLS extension, if the ingress controller fulfilling the ingress supports SNI. See ` + "`" + `tls` + "`" + ` block attributes below.`,
				},
				resource.Attribute{
					Name:        "ingress_class_name",
					Description: `The name of the IngressClass cluster resource. The associated IngressClass defines which controller will implement the resource. This replaces the deprecated ` + "`" + `kubernetes.io/ingress.class` + "`" + ` annotation. For backwards compatibility, when that annotation is set, it must be given precedence over this field. ### ` + "`" + `backend` + "`" + ` #### Attributes`,
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
					Description: `http is a list of http selectors pointing to backends. In the example: http:///? -> backend where parts of the url correspond to RFC 3986, this resource will be used to match against everything after the last '/' and before the first '?' or '#'. See ` + "`" + `http` + "`" + ` block attributes below. #### ` + "`" + `http` + "`" + ``,
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
					Description: `SecretName is the name of the secret used to terminate SSL traffic on 443. Field is left optional to allow SSL routing based on SNI hostname alone. If the SNI host in a listener conflicts with the \"Host\" header field used by an IngressRule, the SNI host is used for termination and value of the Host header is used for routing. ## Attributes ### ` + "`" + `status` + "`" + ``,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status is the current state of the Ingress. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status #### ` + "`" + `load_balancer` + "`" + ``,
				},
				resource.Attribute{
					Name:        "ingress",
					Description: `Ingress is a list containing ingress points for the load-balancer. Traffic intended for the service should be sent to these ingress points. ###### Attributes`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `IP is set for load-balancer ingress points that are IP based (typically GCE or OpenStack load-balancers).`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `Hostname is set for load-balancer ingress points that are DNS based (typically AWS load-balancers).`,
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
					Description: `TLS configuration. Currently the Ingress only supports a single TLS port, 443. If multiple members of this list specify different hosts, they will be multiplexed on the same port according to the hostname specified through the SNI TLS extension, if the ingress controller fulfilling the ingress supports SNI. See ` + "`" + `tls` + "`" + ` block attributes below.`,
				},
				resource.Attribute{
					Name:        "ingress_class_name",
					Description: `The name of the IngressClass cluster resource. The associated IngressClass defines which controller will implement the resource. This replaces the deprecated ` + "`" + `kubernetes.io/ingress.class` + "`" + ` annotation. For backwards compatibility, when that annotation is set, it must be given precedence over this field. ### ` + "`" + `backend` + "`" + ` #### Attributes`,
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
					Description: `http is a list of http selectors pointing to backends. In the example: http:///? -> backend where parts of the url correspond to RFC 3986, this resource will be used to match against everything after the last '/' and before the first '?' or '#'. See ` + "`" + `http` + "`" + ` block attributes below. #### ` + "`" + `http` + "`" + ``,
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
					Description: `SecretName is the name of the secret used to terminate SSL traffic on 443. Field is left optional to allow SSL routing based on SNI hostname alone. If the SNI host in a listener conflicts with the \"Host\" header field used by an IngressRule, the SNI host is used for termination and value of the Host header is used for routing. ## Attributes ### ` + "`" + `status` + "`" + ``,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status is the current state of the Ingress. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status #### ` + "`" + `load_balancer` + "`" + ``,
				},
				resource.Attribute{
					Name:        "ingress",
					Description: `Ingress is a list containing ingress points for the load-balancer. Traffic intended for the service should be sent to these ingress points. ###### Attributes`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `IP is set for load-balancer ingress points that are IP based (typically GCE or OpenStack load-balancers).`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `Hostname is set for load-balancer ingress points that are DNS based (typically AWS load-balancers).`,
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
					Description: `(Optional) An unstructured key value map stored with the namespace that may be used to store arbitrary metadata. ~> By default, the provider ignores any annotations whose key names end with`,
				},
				resource.Attribute{
					Name:        "generation",
					Description: `A sequence number representing a specific generation of the desired state.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) Map of string keys and values that can be used to organize and categorize (scope and select) namespaces. May match selectors of replication controllers and services. ~> By default, the provider ignores any labels whose key names end with`,
				},
				resource.Attribute{
					Name:        "resource_version",
					Description: `An opaque value that represents the internal version of this namespace that can be used by clients to determine when namespaces have changed. Read more about [concurrency control and consistency](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency).`,
				},
				resource.Attribute{
					Name:        "uid",
					Description: `The unique in time and space value for this namespace. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/identifiers#uids) ### ` + "`" + `spec` + "`" + ` #### Attributes`,
				},
				resource.Attribute{
					Name:        "finalizers",
					Description: `An opaque list of values that must be empty to permanently remove object from storage. For more info: https://kubernetes.io/docs/tasks/administer-cluster/namespaces/`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "kubernetes_persistent_volume_claim",
			Category:         "Data Sources",
			ShortDescription: `A PersistentVolumeClaim (PVC) is a request for storage by a user. This data source retrieves information about the specified PVC.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "metadata",
					Description: `(Required) Standard persistent volume claim's metadata. For more info see [Kubernetes reference](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata) ## Nested Blocks ### ` + "`" + `metadata` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the persistent volume claim, must be unique. Cannot be updated. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/identifiers#names)`,
				},
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) Namespace defines the space within which name of the persistent volume claim must be unique. #### Attributes`,
				},
				resource.Attribute{
					Name:        "generation",
					Description: `A sequence number representing a specific generation of the desired state.`,
				},
				resource.Attribute{
					Name:        "resource_version",
					Description: `An opaque value that represents the internal version of this persistent volume claim that can be used by clients to determine when persistent volume claim has changed. For more info see [Kubernetes reference](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency)`,
				},
				resource.Attribute{
					Name:        "uid",
					Description: `The unique in time and space value for this persistent volume claim. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/identifiers#uids) ### ` + "`" + `spec` + "`" + ` #### Attributes`,
				},
				resource.Attribute{
					Name:        "access_modes",
					Description: `A set of the desired access modes the volume should have. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/persistent-volumes#access-modes-1)`,
				},
				resource.Attribute{
					Name:        "selector",
					Description: `Claims can specify a label selector to further filter the set of volumes. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/persistent-volumes#selector)`,
				},
				resource.Attribute{
					Name:        "volume_name",
					Description: `The binding reference to the PersistentVolume backing this claim.`,
				},
				resource.Attribute{
					Name:        "storage_class_name",
					Description: `Name of the storage class requested by the claim. ## Import Persistent Volume Claim can be imported using its namespace and name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import kubernetes_persistent_volume_claim.example default/example-name ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "kubernetes_pod",
			Category:         "Data Sources",
			ShortDescription: `A pod is a group of one or more containers, the shared storage for those containers, and options about how to run the containers. Pods are always co-located and co-scheduled, and run in a shared context.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "metadata",
					Description: `(Required) Standard pod's metadata. For more info see [Kubernetes reference](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata) ## Nested Blocks ### ` + "`" + `metadata` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the pod, must be unique. Cannot be updated. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/identifiers#names)`,
				},
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) Namespace defines the space within which name of the pod must be unique. #### Attributes`,
				},
				resource.Attribute{
					Name:        "generation",
					Description: `A sequence number representing a specific generation of the desired state.`,
				},
				resource.Attribute{
					Name:        "resource_version",
					Description: `An opaque value that represents the internal version of this pod that can be used by clients to determine when pod has changed. For more info see [Kubernetes reference](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency)`,
				},
				resource.Attribute{
					Name:        "uid",
					Description: `The unique in time and space value for this pod. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/identifiers#uids) ### ` + "`" + `spec` + "`" + ` #### Attributes`,
				},
				resource.Attribute{
					Name:        "affinity",
					Description: `A group of affinity scheduling rules. If specified, the pod will be dispatched by specified scheduler. If not specified, the pod will be dispatched by default scheduler.`,
				},
				resource.Attribute{
					Name:        "active_deadline_seconds",
					Description: `Optional duration in seconds the pod may be active on the node relative to StartTime before the system will actively try to mark it failed and kill associated containers. Value must be a positive integer.`,
				},
				resource.Attribute{
					Name:        "automount_service_account_token",
					Description: `Indicates whether a service account token should be automatically mounted. Defaults to true for Pods.`,
				},
				resource.Attribute{
					Name:        "container",
					Description: `List of containers belonging to the pod. Containers cannot currently be added or removed. There must be at least one container in a Pod. Cannot be updated. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/containers)`,
				},
				resource.Attribute{
					Name:        "init_container",
					Description: `List of init containers belonging to the pod. Init containers always run to completion and each must complete successfully before the next is started. For more info see [Kubernetes reference](https://kubernetes.io/docs/concepts/workloads/pods/init-containers)/`,
				},
				resource.Attribute{
					Name:        "dns_policy",
					Description: `Set DNS policy for containers within the pod. Valid values are 'ClusterFirstWithHostNet', 'ClusterFirst', 'Default' or 'None'. DNS parameters given in DNSConfig will be merged with the policy selected with DNSPolicy. To have DNS options set along with hostNetwork, you have to specify DNS policy explicitly to 'ClusterFirstWithHostNet'. Optional: Defaults to 'ClusterFirst', see [Kubernetes reference](https://kubernetes.io/docs/concepts/services-networking/dns-pod-service/#pod-s-dns-policy).`,
				},
				resource.Attribute{
					Name:        "dns_config",
					Description: `Specifies the DNS parameters of a pod. Parameters specified here will be merged to the generated DNS configuration based on DNSPolicy. Defaults to empty. See ` + "`" + `dns_config` + "`" + ` block definition below.`,
				},
				resource.Attribute{
					Name:        "host_alias",
					Description: `List of hosts and IPs that will be injected into the pod's hosts file if specified. Optional: Defaults to empty. See ` + "`" + `host_alias` + "`" + ` block definition below.`,
				},
				resource.Attribute{
					Name:        "host_ipc",
					Description: `Use the host's ipc namespace. Optional: Defaults to false.`,
				},
				resource.Attribute{
					Name:        "host_network",
					Description: `Host networking requested for this pod. Use the host's network namespace. If this option is set, the ports that will be used must be specified.`,
				},
				resource.Attribute{
					Name:        "host_pid",
					Description: `Use the host's pid namespace.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `Specifies the hostname of the Pod If not specified, the pod's hostname will be set to a system-defined value.`,
				},
				resource.Attribute{
					Name:        "image_pull_secrets",
					Description: `(Optional) ImagePullSecrets is an optional list of references to secrets in the same namespace to use for pulling any of the images used by this PodSpec. If specified, these secrets will be passed to individual puller implementations for them to use. For example, in the case of docker, only DockerConfig type secrets are honored. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/images#specifying-imagepullsecrets-on-a-pod)`,
				},
				resource.Attribute{
					Name:        "node_name",
					Description: `NodeName is a request to schedule this pod onto a specific node. If it is non-empty, the scheduler simply schedules this pod onto that node, assuming that it fits resource requirements.`,
				},
				resource.Attribute{
					Name:        "node_selector",
					Description: `NodeSelector is a selector which must be true for the pod to fit on a node. Selector which must match a node's labels for the pod to be scheduled on that node. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/node-selection).`,
				},
				resource.Attribute{
					Name:        "priority_class_name",
					Description: `If specified, indicates the pod's priority. 'system-node-critical' and 'system-cluster-critical' are two special keywords which indicate the highest priorities with the formerer being the highest priority. Any other name must be defined by creating a PriorityClass object with that name. If not specified, the pod priority will be default or zero if there is no default.`,
				},
				resource.Attribute{
					Name:        "restart_policy",
					Description: `Restart policy for all containers within the pod. One of Always, OnFailure, Never. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/pod-states#restartpolicy).`,
				},
				resource.Attribute{
					Name:        "security_context",
					Description: `(SecurityContext holds pod-level security attributes and common container settings. Optional: Defaults to empty`,
				},
				resource.Attribute{
					Name:        "service_account_name",
					Description: `ServiceAccountName is the name of the ServiceAccount to use to run this pod. For more info see https://kubernetes.io/docs/tasks/configure-pod-container/configure-service-account/.`,
				},
				resource.Attribute{
					Name:        "share_process_namespace",
					Description: `Share a single process namespace between all of the containers in a pod. When this is set containers will be able to view and signal processes from other containers in the same pod, and the first process in each container will not be assigned PID 1. HostPID and ShareProcessNamespace cannot both be set.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `If specified, the fully qualified Pod hostname will be "...svc.". If not specified, the pod will not have a domainname at all..`,
				},
				resource.Attribute{
					Name:        "termination_grace_period_seconds",
					Description: `Optional duration in seconds the pod needs to terminate gracefully. May be decreased in delete request. Value must be non-negative integer. The value zero indicates delete immediately. If this value is nil, the default grace period will be used instead. The grace period is the duration in seconds after the processes running in the pod are sent a termination signal and the time when the processes are forcibly halted with a kill signal. Set this value longer than the expected cleanup time for your process.`,
				},
				resource.Attribute{
					Name:        "toleration",
					Description: `Optional pod node tolerations. For more info see [Kubernetes reference](https://kubernetes.io/docs/concepts/configuration/taint-and-toleration/)`,
				},
				resource.Attribute{
					Name:        "volume",
					Description: `(Optional) List of volumes that can be mounted by containers belonging to the pod. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/volumes) ### ` + "`" + `affinity` + "`" + ` #### Attributes`,
				},
				resource.Attribute{
					Name:        "node_affinity",
					Description: `Node affinity scheduling rules for the pod. For more info see [Kubernetes reference](https://kubernetes.io/docs/concepts/configuration/assign-pod-node/#node-affinity-beta-feature)`,
				},
				resource.Attribute{
					Name:        "pod_affinity",
					Description: `Inter-pod topological affinity. rules that specify that certain pods should be placed in the same topological domain (e.g. same node, same rack, same zone, same power domain, etc.) For more info see [Kubernetes reference](https://kubernetes.io/docs/concepts/configuration/assign-pod-node/#inter-pod-affinity-and-anti-affinity-beta-feature)`,
				},
				resource.Attribute{
					Name:        "pod_anti_affinity",
					Description: `Inter-pod topological affinity. rules that specify that certain pods should be placed in the same topological domain (e.g. same node, same rack, same zone, same power domain, etc.) For more info see [Kubernetes reference](https://kubernetes.io/docs/concepts/configuration/assign-pod-node/#inter-pod-affinity-and-anti-affinity-beta-feature) ### ` + "`" + `container` + "`" + ` #### Attributes`,
				},
				resource.Attribute{
					Name:        "args",
					Description: `Arguments to the entrypoint. The docker image's CMD is used if this is not provided. Variable references $(VAR_NAME) are expanded using the container's environment. If a variable cannot be resolved, the reference in the input string will be unchanged. The $(VAR_NAME) syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped references will never be expanded, regardless of whether the variable exists or not. Cannot be updated. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/containers#containers-and-commands)`,
				},
				resource.Attribute{
					Name:        "command",
					Description: `Entrypoint array. Not executed within a shell. The docker image's ENTRYPOINT is used if this is not provided. Variable references $(VAR_NAME) are expanded using the container's environment. If a variable cannot be resolved, the reference in the input string will be unchanged. The $(VAR_NAME) syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped references will never be expanded, regardless of whether the variable exists or not. Cannot be updated. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/containers#containers-and-commands)`,
				},
				resource.Attribute{
					Name:        "env",
					Description: `Block of string name and value pairs to set in the container's environment. May be declared multiple times. Cannot be updated.`,
				},
				resource.Attribute{
					Name:        "env_from",
					Description: `List of sources to populate environment variables in the container. The keys defined within a source must be a C_IDENTIFIER. All invalid keys will be reported as an event when the container is starting. When a key exists in multiple sources, the value associated with the last source will take precedence. Values defined by an Env with a duplicate key will take precedence. Cannot be updated.`,
				},
				resource.Attribute{
					Name:        "image",
					Description: `Docker image name. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/images)`,
				},
				resource.Attribute{
					Name:        "image_pull_policy",
					Description: `Image pull policy. One of Always, Never, IfNotPresent. Defaults to Always if :latest tag is specified, or IfNotPresent otherwise. Cannot be updated. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/images#updating-images)`,
				},
				resource.Attribute{
					Name:        "lifecycle",
					Description: `Actions that the management system should take in response to container lifecycle events`,
				},
				resource.Attribute{
					Name:        "liveness_probe",
					Description: `Periodic probe of container liveness. Container will be restarted if the probe fails. Cannot be updated. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/pod-states#container-probes)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the container specified as a DNS_LABEL. Each container in a pod must have a unique name (DNS_LABEL). Cannot be updated.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `List of ports to expose from the container. Exposing a port here gives the system additional information about the network connections a container uses, but is primarily informational. Not specifying a port here DOES NOT prevent that port from being exposed. Any port which is listening on the default "0.0.0.0" address inside a container will be accessible from the network. Cannot be updated.`,
				},
				resource.Attribute{
					Name:        "readiness_probe",
					Description: `Periodic probe of container service readiness. Container will be removed from service endpoints if the probe fails. Cannot be updated. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/pod-states#container-probes)`,
				},
				resource.Attribute{
					Name:        "resources",
					Description: `Compute Resources required by this container. Cannot be updated. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/persistent-volumes#resources)`,
				},
				resource.Attribute{
					Name:        "security_context",
					Description: `Security options the pod should run with. For more info see https://kubernetes.io/docs/tasks/configure-pod-container/security-context/.`,
				},
				resource.Attribute{
					Name:        "stdin",
					Description: `Whether this container should allocate a buffer for stdin in the container runtime. If this is not set, reads from stdin in the container will always result in EOF.`,
				},
				resource.Attribute{
					Name:        "stdin_once",
					Description: `Whether the container runtime should close the stdin channel after it has been opened by a single attach. When stdin is true the stdin stream will remain open across multiple attach sessions. If stdinOnce is set to true, stdin is opened on container start, is empty until the first client attaches to stdin, and then remains open and accepts data until the client disconnects, at which time stdin is closed and remains closed until the container is restarted. If this flag is false, a container processes that reads from stdin will never receive an EOF.`,
				},
				resource.Attribute{
					Name:        "termination_message_path",
					Description: `Optional: Path at which the file to which the container's termination message will be written is mounted into the container's filesystem. Message written is intended to be brief final status, such as an assertion failure message. Defaults to /dev/termination-log. Cannot be updated.`,
				},
				resource.Attribute{
					Name:        "tty",
					Description: `Whether this container should allocate a TTY for itself`,
				},
				resource.Attribute{
					Name:        "volume_mount",
					Description: `Pod volumes to mount into the container's filesystem. Cannot be updated.`,
				},
				resource.Attribute{
					Name:        "working_dir",
					Description: `Container's working directory. If not specified, the container runtime's default will be used, which might be configured in the container image. Cannot be updated. ### ` + "`" + `config_map` + "`" + ` #### Attributes`,
				},
				resource.Attribute{
					Name:        "default_mode",
					Description: `Optional: mode bits to use on created files by default. Must be a value between 0 and 0777. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.`,
				},
				resource.Attribute{
					Name:        "items",
					Description: `(Optional) If unspecified, each key-value pair in the Data field of the referenced ConfigMap will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the ConfigMap, the volume setup will error unless it is marked ` + "`" + `optional` + "`" + `. Paths must be relative and may not contain the '..' path or start with '..'.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the referent. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/identifiers#names) ### ` + "`" + `config_map_ref` + "`" + ` #### Attributes`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the referent. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/identifiers#names)`,
				},
				resource.Attribute{
					Name:        "optional",
					Description: `Specify whether the ConfigMap must be defined ### ` + "`" + `config_map_key_ref` + "`" + ` #### Attributes`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key to select.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the referent. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/identifiers#names) ### ` + "`" + `dns_config` + "`" + ` #### Attributes`,
				},
				resource.Attribute{
					Name:        "nameservers",
					Description: `A list of DNS name server IP addresses specified as strings. This will be appended to the base nameservers generated from DNSPolicy. Duplicated nameservers will be removed. Optional: Defaults to empty.`,
				},
				resource.Attribute{
					Name:        "option",
					Description: `A list of DNS resolver options specified as blocks with ` + "`" + `name` + "`" + `/` + "`" + `value` + "`" + ` pairs. This will be merged with the base options generated from DNSPolicy. Duplicated entries will be removed. Resolution options given in Options will override those that appear in the base DNSPolicy. Optional: Defaults to empty.`,
				},
				resource.Attribute{
					Name:        "searches",
					Description: `A list of DNS search domains for host-name lookup specified as strings. This will be appended to the base search paths generated from DNSPolicy. Duplicated search paths will be removed. Optional: Defaults to empty. The ` + "`" + `option` + "`" + ` block supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the option.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `Value of the option. Optional: Defaults to empty. ### ` + "`" + `downward_api` + "`" + ` #### Attributes`,
				},
				resource.Attribute{
					Name:        "default_mode",
					Description: `Optional: mode bits to use on created files by default. Must be a value between 0 and 0777. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.`,
				},
				resource.Attribute{
					Name:        "items",
					Description: `If unspecified, each key-value pair in the Data field of the referenced ConfigMap will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the ConfigMap, the volume setup will error. Paths must be relative and may not contain the '..' path or start with '..'. ### ` + "`" + `empty_dir` + "`" + ` #### Attributes`,
				},
				resource.Attribute{
					Name:        "medium",
					Description: `What type of storage medium should back this directory. The default is "" which means to use the node's default medium. Must be an empty string (default) or Memory. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/volumes#emptydir)`,
				},
				resource.Attribute{
					Name:        "size_limit",
					Description: `(Optional) Total amount of local storage required for this EmptyDir volume. For more info see [Kubernetes reference](https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/#resource-units-in-kubernetes) and [Kubernetes Quantity type](https://pkg.go.dev/k8s.io/apimachinery/pkg/api/resource?tab=doc#Quantity). ### ` + "`" + `env` + "`" + ` #### Attributes`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the environment variable. Must be a C_IDENTIFIER`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `Variable references $(VAR_NAME) are expanded using the previous defined environment variables in the container and any service environment variables. If a variable cannot be resolved, the reference in the input string will be unchanged. The $(VAR_NAME) syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped references will never be expanded, regardless of whether the variable exists or not. Defaults to "".`,
				},
				resource.Attribute{
					Name:        "value_from",
					Description: `Source for the environment variable's value ### ` + "`" + `env_from` + "`" + ` #### Attributes`,
				},
				resource.Attribute{
					Name:        "config_map_ref",
					Description: `The ConfigMap to select from`,
				},
				resource.Attribute{
					Name:        "prefix",
					Description: `An optional identifer to prepend to each key in the ConfigMap. Must be a C_IDENTIFIER..`,
				},
				resource.Attribute{
					Name:        "secret_ref",
					Description: `The Secret to select from ### ` + "`" + `exec` + "`" + ` #### Attributes`,
				},
				resource.Attribute{
					Name:        "command",
					Description: `Command is the command line to execute inside the container, the working directory for the command is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy. ### ` + "`" + `image_pull_secrets` + "`" + ` #### Attributes`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the referent. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/identifiers#names) ### ` + "`" + `lifecycle` + "`" + ` #### Attributes`,
				},
				resource.Attribute{
					Name:        "post_start",
					Description: `post_start is called immediately after a container is created. If the handler fails, the container is terminated and restarted according to its restart policy. Other management of the container blocks until the hook completes. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/container-environment#hook-details)`,
				},
				resource.Attribute{
					Name:        "pre_stop",
					Description: `pre_stop is called immediately before a container is terminated. The container is terminated after the handler completes. The reason for termination is passed to the handler. Regardless of the outcome of the handler, the container is eventually terminated. Other management of the container blocks until the hook completes. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/container-environment#hook-details) ### ` + "`" + `limits` + "`" + ` #### Attributes`,
				},
				resource.Attribute{
					Name:        "cpu",
					Description: `CPU`,
				},
				resource.Attribute{
					Name:        "memory",
					Description: `Memory ### ` + "`" + `liveness_probe` + "`" + ` #### Attributes`,
				},
				resource.Attribute{
					Name:        "exec",
					Description: `exec specifies the action to take.`,
				},
				resource.Attribute{
					Name:        "failure_threshold",
					Description: `Minimum consecutive failures for the probe to be considered failed after having succeeded.`,
				},
				resource.Attribute{
					Name:        "http_get",
					Description: `Specifies the http request to perform.`,
				},
				resource.Attribute{
					Name:        "initial_delay_seconds",
					Description: `Number of seconds after the container has started before liveness probes are initiated. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/pod-states#container-probes)`,
				},
				resource.Attribute{
					Name:        "period_seconds",
					Description: `How often (in seconds) to perform the probe`,
				},
				resource.Attribute{
					Name:        "success_threshold",
					Description: `Minimum consecutive successes for the probe to be considered successful after having failed.`,
				},
				resource.Attribute{
					Name:        "tcp_socket",
					Description: `TCPSocket specifies an action involving a TCP port. TCP hooks not yet supported`,
				},
				resource.Attribute{
					Name:        "timeout_seconds",
					Description: `Number of seconds after which the probe times out. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/pod-states#container-probes) ### ` + "`" + `nfs` + "`" + ` #### Attributes`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `Path that is exported by the NFS server. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/volumes#nfs)`,
				},
				resource.Attribute{
					Name:        "read_only",
					Description: `Whether to force the NFS export to be mounted with read-only permissions. Defaults to false. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/volumes#nfs)`,
				},
				resource.Attribute{
					Name:        "server",
					Description: `Server is the hostname or IP address of the NFS server. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/volumes#nfs) ### ` + "`" + `persistent_volume_claim` + "`" + ` #### Attributes`,
				},
				resource.Attribute{
					Name:        "claim_name",
					Description: `ClaimName is the name of a PersistentVolumeClaim in the same`,
				},
				resource.Attribute{
					Name:        "read_only",
					Description: `Will force the ReadOnly setting in VolumeMounts. ### ` + "`" + `photon_persistent_disk` + "`" + ` #### Attributes`,
				},
				resource.Attribute{
					Name:        "fs_type",
					Description: `Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.`,
				},
				resource.Attribute{
					Name:        "pd_id",
					Description: `ID that identifies Photon Controller persistent disk ### ` + "`" + `port` + "`" + ` #### Attributes`,
				},
				resource.Attribute{
					Name:        "container_port",
					Description: `Number of port to expose on the pod's IP address. This must be a valid port number, 0 < x < 65536.`,
				},
				resource.Attribute{
					Name:        "host_ip",
					Description: `What host IP to bind the external port to.`,
				},
				resource.Attribute{
					Name:        "host_port",
					Description: `Number of port to expose on the host. If specified, this must be a valid port number, 0 < x < 65536. If HostNetwork is specified, this must match ContainerPort. Most containers do not need this.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `If specified, this must be an IANA_SVC_NAME and unique within the pod. Each named port in a pod must have a unique name. Name for the port that can be referred to by services`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `Protocol for port. Must be UDP or TCP. Defaults to "TCP". ### ` + "`" + `post_start` + "`" + ` #### Attributes`,
				},
				resource.Attribute{
					Name:        "exec",
					Description: `exec specifies the action to take.`,
				},
				resource.Attribute{
					Name:        "http_get",
					Description: `Specifies the http request to perform.`,
				},
				resource.Attribute{
					Name:        "tcp_socket",
					Description: `TCPSocket specifies an action involving a TCP port. TCP hooks not yet supported ### ` + "`" + `pre_stop` + "`" + ` #### Attributes`,
				},
				resource.Attribute{
					Name:        "exec",
					Description: `exec specifies the action to take.`,
				},
				resource.Attribute{
					Name:        "http_get",
					Description: `Specifies the http request to perform.`,
				},
				resource.Attribute{
					Name:        "tcp_socket",
					Description: `TCPSocket specifies an action involving a TCP port. TCP hooks not yet supported ### ` + "`" + `quobyte` + "`" + ` #### Attributes`,
				},
				resource.Attribute{
					Name:        "group",
					Description: `Group to map volume access to Default is no group`,
				},
				resource.Attribute{
					Name:        "read_only",
					Description: `Whether to force the Quobyte volume to be mounted with read-only permissions. Defaults to false.`,
				},
				resource.Attribute{
					Name:        "registry",
					Description: `Registry represents a single or multiple Quobyte Registry services specified as a string as host:port pair (multiple entries are separated with commas) which acts as the central registry for volumes`,
				},
				resource.Attribute{
					Name:        "user",
					Description: `User to map volume access to Defaults to serivceaccount user`,
				},
				resource.Attribute{
					Name:        "volume",
					Description: `Volume is a string that references an already created Quobyte volume by name. ### ` + "`" + `rbd` + "`" + ` #### Attributes`,
				},
				resource.Attribute{
					Name:        "ceph_monitors",
					Description: `A collection of Ceph monitors. For more info see https://kubernetes.io/docs/concepts/storage/volumes/#cephfs.`,
				},
				resource.Attribute{
					Name:        "fs_type",
					Description: `Filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/volumes#rbd)`,
				},
				resource.Attribute{
					Name:        "keyring",
					Description: `Keyring is the path to key ring for RBDUser. Default is /etc/ceph/keyring. For more info see https://github.com/kubernetes/examples/tree/master/volumes/rbd#how-to-use-it.`,
				},
				resource.Attribute{
					Name:        "rados_user",
					Description: `The rados user name. Default is admin. For more info see https://github.com/kubernetes/examples/tree/master/volumes/rbd#how-to-use-it.`,
				},
				resource.Attribute{
					Name:        "rbd_image",
					Description: `The rados image name. For more info see https://github.com/kubernetes/examples/tree/master/volumes/rbd#how-to-use-it.`,
				},
				resource.Attribute{
					Name:        "rbd_pool",
					Description: `The rados pool name. Default is rbd. For more info see https://github.com/kubernetes/examples/tree/master/volumes/rbd#how-to-use-it.`,
				},
				resource.Attribute{
					Name:        "read_only",
					Description: `Whether to force the read-only setting in VolumeMounts. Defaults to false. For more info see https://github.com/kubernetes/examples/tree/master/volumes/rbd#how-to-use-it.`,
				},
				resource.Attribute{
					Name:        "secret_ref",
					Description: `Name of the authentication secret for RBDUser. If provided overrides keyring. Default is nil. For more info see https://github.com/kubernetes/examples/tree/master/volumes/rbd#how-to-use-it. ### ` + "`" + `readiness_probe` + "`" + ` #### Attributes`,
				},
				resource.Attribute{
					Name:        "exec",
					Description: `exec specifies the action to take.`,
				},
				resource.Attribute{
					Name:        "failure_threshold",
					Description: `Minimum consecutive failures for the probe to be considered failed after having succeeded.`,
				},
				resource.Attribute{
					Name:        "http_get",
					Description: `Specifies the http request to perform.`,
				},
				resource.Attribute{
					Name:        "initial_delay_seconds",
					Description: `Number of seconds after the container has started before liveness probes are initiated. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/pod-states#container-probes)`,
				},
				resource.Attribute{
					Name:        "period_seconds",
					Description: `How often (in seconds) to perform the probe`,
				},
				resource.Attribute{
					Name:        "success_threshold",
					Description: `Minimum consecutive successes for the probe to be considered successful after having failed.`,
				},
				resource.Attribute{
					Name:        "tcp_socket",
					Description: `TCPSocket specifies an action involving a TCP port. TCP hooks not yet supported`,
				},
				resource.Attribute{
					Name:        "timeout_seconds",
					Description: `Number of seconds after which the probe times out. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/pod-states#container-probes) ### ` + "`" + `resources` + "`" + ` #### Attributes`,
				},
				resource.Attribute{
					Name:        "limits",
					Description: `Describes the maximum amount of compute resources allowed. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/compute-resources)/`,
				},
				resource.Attribute{
					Name:        "requests",
					Description: `Describes the minimum amount of compute resources required. ### ` + "`" + `requests` + "`" + ` #### Attributes`,
				},
				resource.Attribute{
					Name:        "cpu",
					Description: `CPU`,
				},
				resource.Attribute{
					Name:        "memory",
					Description: `Memory ### ` + "`" + `resource_field_ref` + "`" + ` #### Attributes`,
				},
				resource.Attribute{
					Name:        "container_name",
					Description: `The name of the container`,
				},
				resource.Attribute{
					Name:        "resource",
					Description: `Resource to select ### ` + "`" + `seccomp_profile` + "`" + ` #### Attributes`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Indicates which kind of seccomp profile will be applied. Valid options are:`,
				},
				resource.Attribute{
					Name:        "Localhost",
					Description: `a profile defined in a file on the node should be used.`,
				},
				resource.Attribute{
					Name:        "RuntimeDefault",
					Description: `the container runtime default profile should be used.`,
				},
				resource.Attribute{
					Name:        "Unconfined",
					Description: `(Default) no profile should be applied.`,
				},
				resource.Attribute{
					Name:        "localhost_profile",
					Description: `Indicates a profile defined in a file on the node should be used. The profile must be preconfigured on the node to work. Must be a descending path, relative to the kubelet's configured seccomp profile location. Must only be set if ` + "`" + `type` + "`" + ` is ` + "`" + `Localhost` + "`" + `. ### ` + "`" + `se_linux_options` + "`" + ` #### Attributes`,
				},
				resource.Attribute{
					Name:        "level",
					Description: `Level is SELinux level label that applies to the container.`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `Role is a SELinux role label that applies to the container.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type is a SELinux type label that applies to the container.`,
				},
				resource.Attribute{
					Name:        "user",
					Description: `User is a SELinux user label that applies to the container. ### ` + "`" + `secret` + "`" + ` #### Attributes`,
				},
				resource.Attribute{
					Name:        "default_mode",
					Description: `Mode bits to use on created files by default. Must be a value between 0 and 0777. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.`,
				},
				resource.Attribute{
					Name:        "items",
					Description: `List of Secret Items to project into the volume. See ` + "`" + `items` + "`" + ` block definition below. If unspecified, each key-value pair in the Data field of the referenced Secret will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the Secret, the volume setup will error unless it is marked ` + "`" + `optional` + "`" + `. Paths must be relative and may not contain the '..' path or start with '..'.`,
				},
				resource.Attribute{
					Name:        "optional",
					Description: `Specify whether the Secret or its keys must be defined.`,
				},
				resource.Attribute{
					Name:        "secret_name",
					Description: `Name of the secret in the pod's namespace to use. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/volumes#secrets) The ` + "`" + `items` + "`" + ` block supports the following:`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key to project.`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `Mode bits to use on this file, must be a value between 0 and 0777. If not specified, the volume defaultMode will be used.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `The relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'. ### ` + "`" + `secret_ref` + "`" + ` #### Attributes`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the referent. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/identifiers#names)`,
				},
				resource.Attribute{
					Name:        "optional",
					Description: `Specify whether the Secret must be defined ### ` + "`" + `secret_key_ref` + "`" + ` #### Attributes`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the secret to select from. Must be a valid secret key.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the referent. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/identifiers#names) ### ` + "`" + `secret_ref` + "`" + ` #### Attributes`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the referent. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/identifiers#names) ### container ` + "`" + `security_context` + "`" + ` #### ArgumAttributesents`,
				},
				resource.Attribute{
					Name:        "allow_privilege_escalation",
					Description: `AllowPrivilegeEscalation controls whether a process can gain more privileges than its parent process. This bool directly controls if the no_new_privs flag will be set on the container process. AllowPrivilegeEscalation is true always when the container is: 1) run as Privileged 2) has CAP_SYS_ADMIN`,
				},
				resource.Attribute{
					Name:        "capabilities",
					Description: `The capabilities to add/drop when running containers. Defaults to the default set of capabilities granted by the container runtime.`,
				},
				resource.Attribute{
					Name:        "privileged",
					Description: `Run container in privileged mode. Processes in privileged containers are essentially equivalent to root on the host. Defaults to false.`,
				},
				resource.Attribute{
					Name:        "read_only_root_filesystem",
					Description: `Whether this container has a read-only root filesystem. Default is false.`,
				},
				resource.Attribute{
					Name:        "run_as_group",
					Description: `The GID to run the entrypoint of the container process. Uses runtime default if unset. May also be set in PodSecurityContext. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.`,
				},
				resource.Attribute{
					Name:        "run_as_non_root",
					Description: `Indicates that the container must run as a non-root user. If true, the Kubelet will validate the image at runtime to ensure that it does not run as UID 0 (root) and fail to start the container if it does. If unset or false, no such validation will be performed. May also be set in PodSecurityContext. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.`,
				},
				resource.Attribute{
					Name:        "run_as_user",
					Description: `The UID to run the entrypoint of the container process. Defaults to user specified in image metadata if unspecified. May also be set in PodSecurityContext. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.`,
				},
				resource.Attribute{
					Name:        "seccomp_profile",
					Description: `The seccomp options to use by the containers in this pod. Note that this field cannot be set when spec.os.name is windows.`,
				},
				resource.Attribute{
					Name:        "se_linux_options",
					Description: `The SELinux context to be applied to the container. If unspecified, the container runtime will allocate a random SELinux context for each container. May also be set in PodSecurityContext. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence. ### ` + "`" + `capabilities` + "`" + ` #### Attributes`,
				},
				resource.Attribute{
					Name:        "add",
					Description: `A list of added capabilities.`,
				},
				resource.Attribute{
					Name:        "drop",
					Description: `A list of removed capabilities. ### pod ` + "`" + `security_context` + "`" + ` #### Attributes`,
				},
				resource.Attribute{
					Name:        "fs_group",
					Description: `A special supplemental group that applies to all containers in a pod. Some volume types allow the Kubelet to change the ownership of that volume to be owned by the pod: 1. The owning GID will be the FSGroup 2. The setgid bit is set (new files created in the volume will be owned by FSGroup) 3. The permission bits are OR'd with rw-rw---- If unset, the Kubelet will not modify the ownership and permissions of any volume.`,
				},
				resource.Attribute{
					Name:        "run_as_group",
					Description: `The GID to run the entrypoint of the container process. Uses runtime default if unset. May also be set in SecurityContext. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence for that container.`,
				},
				resource.Attribute{
					Name:        "run_as_non_root",
					Description: `Indicates that the container must run as a non-root user. If true, the Kubelet will validate the image at runtime to ensure that it does not run as UID 0 (root) and fail to start the container if it does. If unset or false, no such validation will be performed. May also be set in SecurityContext. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.`,
				},
				resource.Attribute{
					Name:        "run_as_user",
					Description: `The UID to run the entrypoint of the container process. Defaults to user specified in image metadata if unspecified. May also be set in SecurityContext. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence for that container.`,
				},
				resource.Attribute{
					Name:        "seccomp_profile",
					Description: `The seccomp options to use by the containers in this pod. Note that this field cannot be set when spec.os.name is windows.`,
				},
				resource.Attribute{
					Name:        "se_linux_options",
					Description: `The SELinux context to be applied to all containers. If unspecified, the container runtime will allocate a random SELinux context for each container. May also be set in SecurityContext. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence for that container.`,
				},
				resource.Attribute{
					Name:        "supplemental_groups",
					Description: `A list of groups applied to the first process run in each container, in addition to the container's primary GID. If unspecified, no groups will be added to any container. ### ` + "`" + `tcp_socket` + "`" + ` #### Attributes`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME. ### ` + "`" + `value_from` + "`" + ` #### Attributes`,
				},
				resource.Attribute{
					Name:        "config_map_key_ref",
					Description: `Selects a key of a ConfigMap.`,
				},
				resource.Attribute{
					Name:        "field_ref",
					Description: `(Optional) Selects a field of the pod: supports metadata.name, metadata.namespace, metadata.labels, metadata.annotations, spec.nodeName, spec.serviceAccountName, status.podIP.`,
				},
				resource.Attribute{
					Name:        "resource_field_ref",
					Description: `(Optional) Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported.`,
				},
				resource.Attribute{
					Name:        "secret_key_ref",
					Description: `(Optional) Selects a key of a secret in the pod's namespace. ### ` + "`" + `volume_mount` + "`" + ` #### Attributes`,
				},
				resource.Attribute{
					Name:        "mount_path",
					Description: `Path within the container at which the volume should be mounted. Must not contain ':'.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `This must match the Name of a Volume.`,
				},
				resource.Attribute{
					Name:        "read_only",
					Description: `Mounted read-only if true, read-write otherwise (false or unspecified). Defaults to false.`,
				},
				resource.Attribute{
					Name:        "sub_path",
					Description: `Path within the volume from which the container's volume should be mounted. Defaults to "" (volume's root).`,
				},
				resource.Attribute{
					Name:        "mount_propagation",
					Description: `Mount propagation mode. Defaults to "None". For more info see [Kubernetes reference](https://kubernetes.io/docs/concepts/storage/volumes/#mount-propagation) ## Argument Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The current status of the pods. ## Import Pod can be imported using the namespace and name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import kubernetes_pod.example default/terraform-example ` + "`" + `` + "`" + `` + "`" + ``,
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
					Name:        "uid",
					Description: `The unique in time and space value for this secret. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/identifiers#uids) ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "data",
					Description: `A map of the secret data.`,
				},
				resource.Attribute{
					Name:        "binary_data",
					Description: `A map of the secret data with values encoded in base64 format. ~> In case the secret has been created outside terraform in order to retrieve binary data from the secret in base64 format you need to define a ` + "`" + `binary_data` + "`" + ` map with data to retrieve as key and an empty string as a value ` + "`" + `` + "`" + `` + "`" + `hcl data "kubernetes_secret" "example" { metadata { name = "example-secret" namespace = "kube-system" } binary_data = { "keystore.p12" = "" another_field = "" } } ` + "`" + `` + "`" + `` + "`" + ``,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The secret type. Defaults to ` + "`" + `Opaque` + "`" + `. For more info see [Kubernetes reference](https://github.com/kubernetes/community/blob/c7151dd8dd7e487e96e5ce34c6a416bb3b037609/contributors/design-proposals/auth/secrets.md#proposed-design)`,
				},
				resource.Attribute{
					Name:        "immutable",
					Description: `Ensures that data stored in the Secret cannot be updated (only object metadata can be modified).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "data",
					Description: `A map of the secret data.`,
				},
				resource.Attribute{
					Name:        "binary_data",
					Description: `A map of the secret data with values encoded in base64 format. ~> In case the secret has been created outside terraform in order to retrieve binary data from the secret in base64 format you need to define a ` + "`" + `binary_data` + "`" + ` map with data to retrieve as key and an empty string as a value ` + "`" + `` + "`" + `` + "`" + `hcl data "kubernetes_secret" "example" { metadata { name = "example-secret" namespace = "kube-system" } binary_data = { "keystore.p12" = "" another_field = "" } } ` + "`" + `` + "`" + `` + "`" + ``,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The secret type. Defaults to ` + "`" + `Opaque` + "`" + `. For more info see [Kubernetes reference](https://github.com/kubernetes/community/blob/c7151dd8dd7e487e96e5ce34c6a416bb3b037609/contributors/design-proposals/auth/secrets.md#proposed-design)`,
				},
				resource.Attribute{
					Name:        "immutable",
					Description: `Ensures that data stored in the Secret cannot be updated (only object metadata can be modified).`,
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
					Description: `Spec defines the behavior of a service. [Kubernetes reference](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#spec-and-status) ## Nested Blocks ### ` + "`" + `metadata` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the service, must be unique. Cannot be updated. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/identifiers#names)`,
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
					Name:        "uid",
					Description: `The unique in time and space value for this service. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/identifiers#uids) ### ` + "`" + `port` + "`" + ` #### Attributes`,
				},
				resource.Attribute{
					Name:        "app_protocol",
					Description: `(Optional) The application protocol for this port. This field follows standard Kubernetes label syntax. Un-prefixed names are reserved for IANA standard service names (as per [RFC-6335](https://datatracker.ietf.org/doc/html/rfc6335) and [IANA standard service names](http://www.iana.org/assignments/service-names)). Non-standard protocols should use prefixed names such as ` + "`" + `mycompany.com/my-custom-protocol` + "`" + `. For more info see [Kubernetes reference](https://kubernetes.io/docs/concepts/services-networking/service/#application-protocol)`,
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
					Description: `(Optional) Denotes if this Service desires to route external traffic to node-local or cluster-wide endpoints. ` + "`" + `Local` + "`" + ` preserves the client source IP and avoids a second hop for LoadBalancer and Nodeport type services, but risks potentially imbalanced traffic spreading. ` + "`" + `Cluster` + "`" + ` obscures the client source IP and may cause a second hop to another node, but should have good overall load-spreading. For more info see [Kubernetes reference](https://kubernetes.io/docs/tutorials/services/source-ip/)`,
				},
				resource.Attribute{
					Name:        "ip_families",
					Description: `(Optional) A list of IP families (e.g. IPv4, IPv6) assigned to this service. This field is usually assigned automatically based on cluster configuration and the ` + "`" + `ip_family_policy` + "`" + ` field. If this field is specified manually, the requested family is available in the cluster, and ` + "`" + `ip_family_policy` + "`" + ` allows it, it will be used; otherwise creation of the service will fail. This field is conditionally mutable: it allows for adding or removing a secondary IP family, but it does not allow changing the primary IP family of the Service. For more info see [Kubernetes reference](https://kubernetes.io/docs/concepts/services-networking/dual-stack/)`,
				},
				resource.Attribute{
					Name:        "ip_family_policy",
					Description: `(Optional) Represents the dual-stack-ness requested or required by this Service. If there is no value provided, then this field will be set to ` + "`" + `SingleStack` + "`" + `. Services can be ` + "`" + `SingleStack` + "`" + `(a single IP family), ` + "`" + `PreferDualStack` + "`" + `(two IP families on dual-stack configured clusters or a single IP family on single-stack clusters), or ` + "`" + `RequireDualStack` + "`" + `(two IP families on dual-stack configured clusters, otherwise fail). The ` + "`" + `ip_families` + "`" + ` and ` + "`" + `cluster_ip` + "`" + ` fields depend on the value of this field.`,
				},
				resource.Attribute{
					Name:        "load_balancer_ip",
					Description: `Only applies to ` + "`" + `type = LoadBalancer` + "`" + `. LoadBalancer will get created with the IP specified in this field. This feature depends on whether the underlying cloud-provider supports specifying this field when a load balancer is created. This field will be ignored if the cloud-provider does not support the feature.`,
				},
				resource.Attribute{
					Name:        "load_balancer_source_ranges",
					Description: `If specified and supported by the platform, this will restrict traffic through the cloud-provider load-balancer will be restricted to the specified client IPs. This field will be ignored if the cloud-provider does not support the feature. For more info see [Kubernetes reference](https://kubernetes.io/docs/tasks/access-application-cluster/create-external-load-balancer/).`,
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
					Description: `Determines how the service is exposed. Defaults to ` + "`" + `ClusterIP` + "`" + `. Valid options are ` + "`" + `ExternalName` + "`" + `, ` + "`" + `ClusterIP` + "`" + `, ` + "`" + `NodePort` + "`" + `, and ` + "`" + `LoadBalancer` + "`" + `. ` + "`" + `ExternalName` + "`" + ` maps to the specified ` + "`" + `external_name` + "`" + `. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/services#overview) ## Attributes`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status is a list containing the most recently observed status of the service. Populated by the system. Read-only. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status ### ` + "`" + `status` + "`" + ` #### Attributes`,
				},
				resource.Attribute{
					Name:        "load_balancer",
					Description: `a list containing the current status of the load-balancer, if one is present. ### ` + "`" + `load_balancer` + "`" + ` #### Attributes`,
				},
				resource.Attribute{
					Name:        "ingress",
					Description: `a list containing ingress points for the load-balancer. Traffic intended for the service should be sent to these ingress points. ### ` + "`" + `ingress` + "`" + ` #### Attributes`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `IP is set for load-balancer ingress points that are IP based (typically GCE or OpenStack load-balancers).`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `Hostname is set for load-balancer ingress points that are DNS based (typically AWS load-balancers).`,
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
					Description: `Name of the default secret, containing service account token, created & managed by the service. By default, the provider will try to find the secret containing the service account token that Kubernetes automatically created for the service account. Where there are multiple tokens and the provider cannot determine which was created by Kubernetes, this attribute will be empty. When only one token is associated with the service account, the provider will return this single token secret. ### ` + "`" + `image_pull_secret` + "`" + ` #### Attributes`,
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
					Description: `Name of the default secret, containing service account token, created & managed by the service. By default, the provider will try to find the secret containing the service account token that Kubernetes automatically created for the service account. Where there are multiple tokens and the provider cannot determine which was created by Kubernetes, this attribute will be empty. When only one token is associated with the service account, the provider will return this single token secret. ### ` + "`" + `image_pull_secret` + "`" + ` #### Attributes`,
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
					Description: `(Required) Name of the storage class, must be unique. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/identifiers#names) ### ` + "`" + `allowed_topologies` + "`" + ` ￼ #### Arguments ￼`,
				},
				resource.Attribute{
					Name:        "match_label_expressions",
					Description: `(Optional) A list of topology selector requirements by labels. See [match_label_expressions](#match_label_expressions) ### ` + "`" + `match_label_expressions` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Optional) The label key that the selector applies to.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Optional) An array of string values. One value must match the label to be selected. #### Attributes`,
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
				resource.Attribute{
					Name:        "allowed_topologies",
					Description: `(Optional) Restrict the node topologies where volumes can be dynamically provisioned. See [allowed_topologies](#allowed_topologies)`,
				},
			},
			Attributes: []resource.Attribute{},
		},
	}

	dataSourcesMap = map[string]int{

		"kubernetes_all_namespaces":          0,
		"kubernetes_config_map":              1,
		"kubernetes_ingress":                 2,
		"kubernetes_namespace":               3,
		"kubernetes_persistent_volume_claim": 4,
		"kubernetes_pod":                     5,
		"kubernetes_secret":                  6,
		"kubernetes_service":                 7,
		"kubernetes_service_account":         8,
		"kubernetes_storage_class":           9,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
