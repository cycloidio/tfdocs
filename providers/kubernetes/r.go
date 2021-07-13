package kubernetes

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "kubernetes_api_service",
			Category:         "Resources",
			ShortDescription: `An API Service is an abstraction which defines for locating and communicating with servers.`,
			Description:      ``,
			Keywords: []string{
				"api",
				"service",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "metadata",
					Description: `(Required) Standard API service's metadata. For more info see [Kubernetes reference](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata)`,
				},
				resource.Attribute{
					Name:        "spec",
					Description: `(Required) Spec contains information for locating and communicating with a server. [Kubernetes reference](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#spec-and-status) ## Nested Blocks ### ` + "`" + `metadata` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "annotations",
					Description: `(Optional) An unstructured key value map stored with the API service that may be used to store arbitrary metadata. ~> By default, the provider ignores any annotations whose key names end with`,
				},
				resource.Attribute{
					Name:        "generate_name",
					Description: `(Optional) Prefix, used by the server, to generate a unique name ONLY IF the ` + "`" + `name` + "`" + ` field has not been provided. This value will also be combined with a unique suffix. For more info see [Kubernetes reference](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#idempotency)`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) Map of string keys and values that can be used to organize and categorize (scope and select) the API service. ~> By default, the provider ignores any labels whose key names end with`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the API service, must be unique. Cannot be updated. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/identifiers#names) #### Attributes`,
				},
				resource.Attribute{
					Name:        "generation",
					Description: `A sequence number representing a specific generation of the desired state.`,
				},
				resource.Attribute{
					Name:        "resource_version",
					Description: `An opaque value that represents the internal version of this API service that can be used by clients to determine when API service has changed. For more info see [Kubernetes reference](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency)`,
				},
				resource.Attribute{
					Name:        "uid",
					Description: `The unique in time and space value for this API service. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/identifiers#uids) ### ` + "`" + `spec` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "ca_bundle",
					Description: `(Optional) CABundle is a PEM encoded CA bundle which will be used to validate an API server's serving certificate. If unspecified, system trust roots on the apiserver are used.`,
				},
				resource.Attribute{
					Name:        "group",
					Description: `(Required) Group is the API group name this server hosts.`,
				},
				resource.Attribute{
					Name:        "group_priority_minimum",
					Description: `(Required) GroupPriorityMininum is the priority this group should have at least. Higher priority means that the group is preferred by clients over lower priority ones. Note that other versions of this group might specify even higher GroupPriorityMininum values such that the whole group gets a higher priority. The primary sort is based on GroupPriorityMinimum, ordered highest number to lowest (20 before 10). The secondary sort is based on the alphabetical comparison of the name of the object. (v1.bar before v1.foo) We'd recommend something like:`,
				},
				resource.Attribute{
					Name:        "insecure_skip_tls_verify",
					Description: `(Required) InsecureSkipTLSVerify disables TLS certificate verification when communicating with this server. This is strongly discouraged. You should use the CABundle instead.`,
				},
				resource.Attribute{
					Name:        "service",
					Description: `(Optional) Service is a reference to the service for this API server. It must communicate on port 443. If the Service is nil, that means the handling for the API groupversion is handled locally on this server. The call will simply delegate to the normal handler chain to be fulfilled. See ` + "`" + `service` + "`" + ` block attributes below.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Required) Version is the API version this server hosts. For example, ` + "`" + `v1` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "version_priority",
					Description: `(Required) VersionPriority controls the ordering of this API version inside of its group. Must be greater than zero. The primary sort is based on VersionPriority, ordered highest to lowest (20 before 10). Since it's inside of a group, the number can be small, probably in the 10s. In case of equal version priorities, the version string will be used to compute the order inside a group. If the version string is ` + "`" + `kube-like` + "`" + `, it will sort above non ` + "`" + `kube-like` + "`" + ` version strings, which are ordered lexicographically. ` + "`" + `Kube-like` + "`" + ` versions start with a ` + "`" + `v` + "`" + `, then are followed by a number (the major version), then optionally the string ` + "`" + `alpha` + "`" + ` or ` + "`" + `beta` + "`" + ` and another number (the minor version). These are sorted first by GA > ` + "`" + `beta` + "`" + ` > ` + "`" + `alpha` + "`" + ` (where GA is a version with no suffix such as ` + "`" + `beta` + "`" + ` or ` + "`" + `alpha` + "`" + `), and then by comparing major version, then minor version. An example sorted list of versions: ` + "`" + `v10` + "`" + `, ` + "`" + `v2` + "`" + `, ` + "`" + `v1` + "`" + `, ` + "`" + `v11beta2` + "`" + `, ` + "`" + `v10beta3` + "`" + `, ` + "`" + `v3beta1` + "`" + `, ` + "`" + `v12alpha1` + "`" + `, ` + "`" + `v11alpha2` + "`" + `, ` + "`" + `foo1` + "`" + `, ` + "`" + `foo10` + "`" + `.. ### ` + "`" + `service` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name is the name of the service.`,
				},
				resource.Attribute{
					Name:        "namespace",
					Description: `(Required) Namespace is the namespace of the service.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional) If specified, the port on the service that is hosting the service. Defaults to 443 for backward compatibility. Should be a valid port number (1-65535, inclusive). ## Import API service can be imported using its name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import kubernetes_api_service.example v1.terraform-name.k8s.io ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "kubernetes_certificate_signing_request",
			Category:         "Resources",
			ShortDescription: `Use this resource to generate TLS certificates using Kubernetes.`,
			Description:      ``,
			Keywords: []string{
				"certificate",
				"signing",
				"request",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "auto_approve",
					Description: `(Optional) Automatically approve the CertificateSigningRequest. Defaults to 'true'.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `(Required) Standard certificate signing request's metadata. For more info see [Kubernetes reference](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata)`,
				},
				resource.Attribute{
					Name:        "spec",
					Description: `(Required) Spec defines the specification of the desired behavior of the deployment. For more info see [Kubernetes reference](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#spec-and-status) ## Nested Blocks ### ` + "`" + `metadata` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "annotations",
					Description: `(Optional) An unstructured key value map stored with the certificate signing request that may be used to store arbitrary metadata. ~> By default, the provider ignores any annotations whose key names end with`,
				},
				resource.Attribute{
					Name:        "generate_name",
					Description: `(Optional) Prefix, used by the server, to generate a unique name ONLY IF the ` + "`" + `name` + "`" + ` field has not been provided. This value will also be combined with a unique suffix. For more info see [Kubernetes reference](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#idempotency)`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) Map of string keys and values that can be used to organize and categorize (scope and select) the certificate signing request. May match selectors of replication controllers and services. ~> By default, the provider ignores any labels whose key names end with`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the certificate signing request, must be unique. Cannot be updated. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/identifiers#names) #### Attributes`,
				},
				resource.Attribute{
					Name:        "certificate",
					Description: `The signed certificate PEM data.`,
				},
				resource.Attribute{
					Name:        "generation",
					Description: `A sequence number representing a specific generation of the desired state.`,
				},
				resource.Attribute{
					Name:        "resource_version",
					Description: `An opaque value that represents the internal version of this certificate signing request that can be used by clients to determine when certificate signing request has changed. For more info see [Kubernetes reference](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency)`,
				},
				resource.Attribute{
					Name:        "uid",
					Description: `The unique in time and space value for this certificate signing request. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/identifiers#uids) ### ` + "`" + `spec` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "request",
					Description: `(Required) Base64-encoded PKCS#10 CSR data.`,
				},
				resource.Attribute{
					Name:        "signer_name",
					Description: `(Optional) Requested signer for the request. It is a qualified name in the form: ` + "`" + `scope-hostname.io/name` + "`" + `. If empty, it will be defaulted: 1. If it's a kubelet client certificate, it is assigned "kubernetes.io/kube-apiserver-client-kubelet". 2. If it's a kubelet serving certificate, it is assigned "kubernetes.io/kubelet-serving". 3. Otherwise, it is assigned "kubernetes.io/legacy-unknown". Distribution of trust for signers happens out of band.`,
				},
				resource.Attribute{
					Name:        "usages",
					Description: `(Required) Specifies a set of usage contexts the key will be valid for. See https://godoc.org/k8s.io/api/certificates/v1beta1#KeyUsage ## Generating a New Certificate Since the certificate is a logical resource that lives only in the Terraform state, it will persist until it is explicitly destroyed by the user. In order to force the generation of a new certificate within an existing state, the certificate instance can be "tainted": ` + "`" + `` + "`" + `` + "`" + ` terraform taint kubernetes_certificate_signing_request.example ` + "`" + `` + "`" + `` + "`" + ` A new certificate will then be generated on the next ` + "`" + `` + "`" + `terraform apply` + "`" + `` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "kubernetes_cluster_role",
			Category:         "Resources",
			ShortDescription: `A ClusterRole creates a role at the cluster level and in all namespaces.`,
			Description:      ``,
			Keywords: []string{
				"cluster",
				"role",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "metadata",
					Description: `(Required) Standard kubernetes metadata. For more info see [Kubernetes reference](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata)`,
				},
				resource.Attribute{
					Name:        "rule",
					Description: `(Optional) The PolicyRoles for this ClusterRole. For more info see [Kubernetes reference](https://kubernetes.io/docs/reference/access-authn-authz/rbac/#role-and-clusterrole)`,
				},
				resource.Attribute{
					Name:        "aggregation_rule",
					Description: `(Optional) Describes how to build the Rules for this ClusterRole. If AggregationRule is set, then the Rules are controller managed and direct changes to Rules will be overwritten by the controller. . For more info see [Kubernetes reference](https://kubernetes.io/docs/reference/access-authn-authz/rbac/#aggregated-clusterroles) ## Nested Blocks ### ` + "`" + `metadata` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "annotations",
					Description: `(Optional) An unstructured key value map stored with the cluster role binding that may be used to store arbitrary metadata. ~> By default, the provider ignores any annotations whose key names end with`,
				},
				resource.Attribute{
					Name:        "generate_name",
					Description: `(Optional) Prefix, used by the server, to generate a unique name ONLY IF the ` + "`" + `name` + "`" + ` field has not been provided. This value will also be combined with a unique suffix. For more info see [Kubernetes reference](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#idempotency)`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) Map of string keys and values that can be used to organize and categorize (scope and select) the cluster role binding. ~> By default, the provider ignores any labels whose key names end with`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the cluster role binding, must be unique. Cannot be updated. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/identifiers#names) #### Attributes`,
				},
				resource.Attribute{
					Name:        "generation",
					Description: `A sequence number representing a specific generation of the desired state.`,
				},
				resource.Attribute{
					Name:        "resource_version",
					Description: `An opaque value that represents the internal version of this object that can be used by clients to determine when the object has changed. For more info see [Kubernetes reference](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency)`,
				},
				resource.Attribute{
					Name:        "uid",
					Description: `The unique in time and space value for this cluster role binding. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/identifiers#uids) ### ` + "`" + `rule` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "api_groups",
					Description: `(Optional) APIGroups is the name of the APIGroup that contains the resources. If multiple API groups are specified, any action requested against one of the enumerated resources in any API group will be allowed.`,
				},
				resource.Attribute{
					Name:        "non_resource_urls",
					Description: `(Optional) NonResourceURLs is a set of partial urls that a user should have access to. \`,
				},
				resource.Attribute{
					Name:        "resource_names",
					Description: `(Optional) ResourceNames is an optional white list of names that the rule applies to. An empty set means that everything is allowed.`,
				},
				resource.Attribute{
					Name:        "resources",
					Description: `(Optional) Resources is a list of resources this rule applies to. ResourceAll represents all resources.`,
				},
				resource.Attribute{
					Name:        "verbs",
					Description: `(Required) Verbs is a list of Verbs that apply to ALL the ResourceKinds and AttributeRestrictions contained in this rule. VerbAll represents all kinds. ### ` + "`" + `aggregation_rule` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "cluster_role_selectors",
					Description: `(Optional) A list of selectors which will be used to find ClusterRoles and create the rules. ### ` + "`" + `cluster_role_selectors` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "match_expressions",
					Description: `(Optional) A list of label selector requirements. The requirements are ANDed.`,
				},
				resource.Attribute{
					Name:        "match_labels",
					Description: `(Optional) A map of ` + "`" + `{key,value}` + "`" + ` pairs. A single ` + "`" + `{key,value}` + "`" + ` in the matchLabels map is equivalent to an element of matchExpressions, whose key field is "key", the operator is "In", and the values array contains only "value". The requirements are ANDed. ## Import ClusterRole can be imported using the name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import kubernetes_cluster_role.example terraform-name ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "kubernetes_cluster_role_binding",
			Category:         "Resources",
			ShortDescription: `A ClusterRoleBinding may be used to grant permission at the cluster level and in all namespaces.`,
			Description:      ``,
			Keywords: []string{
				"cluster",
				"role",
				"binding",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "metadata",
					Description: `(Required) Standard kubernetes metadata. For more info see [Kubernetes reference](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata)`,
				},
				resource.Attribute{
					Name:        "role_ref",
					Description: `(Required) The ClusterRole to bind Subjects to. For more info see [Kubernetes reference](https://kubernetes.io/docs/admin/authorization/rbac/#rolebinding-and-clusterrolebinding)`,
				},
				resource.Attribute{
					Name:        "subject",
					Description: `(Required) The Users, Groups, or ServiceAccounts to grant permissions to. For more info see [Kubernetes reference](https://kubernetes.io/docs/admin/authorization/rbac/#referring-to-subjects) ## Nested Blocks ### ` + "`" + `metadata` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "annotations",
					Description: `(Optional) An unstructured key value map stored with the cluster role binding that may be used to store arbitrary metadata. ~> By default, the provider ignores any annotations whose key names end with`,
				},
				resource.Attribute{
					Name:        "generate_name",
					Description: `(Optional) Prefix, used by the server, to generate a unique name ONLY IF the ` + "`" + `name` + "`" + ` field has not been provided. This value will also be combined with a unique suffix. For more info see [Kubernetes reference](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#idempotency)`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) Map of string keys and values that can be used to organize and categorize (scope and select) the cluster role binding. ~> By default, the provider ignores any labels whose key names end with`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the cluster role binding, must be unique. Cannot be updated. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/identifiers#names) #### Attributes`,
				},
				resource.Attribute{
					Name:        "generation",
					Description: `A sequence number representing a specific generation of the desired state.`,
				},
				resource.Attribute{
					Name:        "resource_version",
					Description: `An opaque value that represents the internal version of this object that can be used by clients to determine when the object has changed. For more info see [Kubernetes reference](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency)`,
				},
				resource.Attribute{
					Name:        "uid",
					Description: `The unique in time and space value for this cluster role binding. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/identifiers#uids) ### ` + "`" + `role_ref` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of this ClusterRole to bind Subjects to.`,
				},
				resource.Attribute{
					Name:        "kind",
					Description: `(Required) The type of binding to use. This value must be and defaults to ` + "`" + `ClusterRole` + "`" + ``,
				},
				resource.Attribute{
					Name:        "api_group",
					Description: `(Required) The API group to drive authorization decisions. This value must be and defaults to ` + "`" + `rbac.authorization.k8s.io` + "`" + ` ### ` + "`" + `subject` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of this ClusterRole to bind Subjects to.`,
				},
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) Namespace defines the namespace of the ServiceAccount to bind to. This value only applies to kind ` + "`" + `ServiceAccount` + "`" + ``,
				},
				resource.Attribute{
					Name:        "kind",
					Description: `(Required) The type of binding to use. This value must be ` + "`" + `ServiceAccount` + "`" + `, ` + "`" + `User` + "`" + ` or ` + "`" + `Group` + "`" + ``,
				},
				resource.Attribute{
					Name:        "api_group",
					Description: `(Required) The API group to drive authorization decisions. This value only applies to kind ` + "`" + `User` + "`" + ` and ` + "`" + `Group` + "`" + `. It must be ` + "`" + `rbac.authorization.k8s.io` + "`" + ` ## Import ClusterRoleBinding can be imported using the name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import kubernetes_cluster_role_binding.example terraform-name ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "kubernetes_config_map",
			Category:         "Resources",
			ShortDescription: `The resource provides mechanisms to inject containers with configuration data while keeping containers agnostic of Kubernetes.`,
			Description:      ``,
			Keywords: []string{
				"config",
				"map",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "binary_data",
					Description: `(Optional) BinaryData contains the binary data. Each key must consist of alphanumeric characters, '-', '_' or '.'. BinaryData can contain byte sequences that are not in the UTF-8 range. The keys stored in BinaryData must not overlap with the ones in the Data field, this is enforced during validation process. Using this field will require 1.10+ apiserver and kubelet. This field only accepts base64-encoded payloads that will be decoded/received before being sent/received to the apiserver.`,
				},
				resource.Attribute{
					Name:        "data",
					Description: `(Optional) Data contains the configuration data. Each key must consist of alphanumeric characters, '-', '_' or '.'. Values with non-UTF-8 byte sequences must use the BinaryData field. The keys stored in Data must not overlap with the keys in the BinaryData field, this is enforced during validation process.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `(Required) Standard config map's metadata. For more info see [Kubernetes reference](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata) ## Nested Blocks ### ` + "`" + `metadata` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "annotations",
					Description: `(Optional) An unstructured key value map stored with the config map that may be used to store arbitrary metadata. ~> By default, the provider ignores any annotations whose key names end with`,
				},
				resource.Attribute{
					Name:        "generate_name",
					Description: `(Optional) Prefix, used by the server, to generate a unique name ONLY IF the ` + "`" + `name` + "`" + ` field has not been provided. This value will also be combined with a unique suffix. For more info see [Kubernetes reference](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#idempotency)`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) Map of string keys and values that can be used to organize and categorize (scope and select) the config map. May match selectors of replication controllers and services. ~> By default, the provider ignores any labels whose key names end with`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the config map, must be unique. Cannot be updated. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/identifiers#names)`,
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
					Description: `The unique in time and space value for this config map. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/identifiers#uids) ## Import Config Map can be imported using its namespace and name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import kubernetes_config_map.example default/my-config ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "kubernetes_cron_job",
			Category:         "Resources",
			ShortDescription: `A Cron Job creates Jobs on a time-based schedule. One CronJob object is like one line of a crontab (cron table) file.`,
			Description:      ``,
			Keywords: []string{
				"cron",
				"job",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "metadata",
					Description: `(Required) Standard resource's metadata. For more info: https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#spec-and-status`,
				},
				resource.Attribute{
					Name:        "spec",
					Description: `(Required) Spec defines the behavior of a CronJob. https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#spec-and-status ## Nested Blocks ### ` + "`" + `metadata` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "annotations",
					Description: `(Optional) An unstructured key value map stored with the resource that may be used to store arbitrary metadata. ~> By default, the provider ignores any annotations whose key names end with`,
				},
				resource.Attribute{
					Name:        "generate_name",
					Description: `(Optional) Prefix, used by the server, to generate a unique name ONLY IF the ` + "`" + `name` + "`" + ` field has not been provided. This value will also be combined with a unique suffix. Read more: https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#idempotency`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) Map of string keys and values that can be used to organize and categorize (scope and select) the service. May match selectors of replication controllers and services. ~> By default, the provider ignores any labels whose key names end with`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the service, must be unique. Cannot be updated. For more info: http://kubernetes.io/docs/user-guide/identifiers#names`,
				},
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) Namespace defines the space within which name of the service must be unique. #### Attributes`,
				},
				resource.Attribute{
					Name:        "generation",
					Description: `A sequence number representing a specific generation of the desired state.`,
				},
				resource.Attribute{
					Name:        "resource_version",
					Description: `An opaque value that represents the internal version of this service that can be used by clients to determine when service has changed. Read more: https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency`,
				},
				resource.Attribute{
					Name:        "uid",
					Description: `The unique in time and space value for this service. For more info: http://kubernetes.io/docs/user-guide/identifiers#uids ### ` + "`" + `spec` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "concurrency_policy",
					Description: `(Optional) Specifies how to treat concurrent executions of a Job. Valid values are: - "Allow" (default): allows CronJobs to run concurrently; - "Forbid": forbids concurrent runs, skipping next run if previous run hasn't finished yet; - "Replace": cancels currently running job and replaces it with a new one`,
				},
				resource.Attribute{
					Name:        "failed_jobs_history_limit",
					Description: `(Optional) The number of failed finished jobs to retain. This is a pointer to distinguish between explicit zero and not specified. Defaults to 1.`,
				},
				resource.Attribute{
					Name:        "job_template",
					Description: `(Required) Specifies the job that will be created when executing a CronJob.`,
				},
				resource.Attribute{
					Name:        "schedule",
					Description: `(Required) The schedule in Cron format, see https://en.wikipedia.org/wiki/Cron.`,
				},
				resource.Attribute{
					Name:        "starting_deadline_seconds",
					Description: `(Optional) Deadline in seconds for starting the job if it misses scheduled time for any reason. Missed jobs executions will be counted as failed ones.`,
				},
				resource.Attribute{
					Name:        "successful_jobs_history_limit",
					Description: `(Optional) The number of successful finished jobs to retain. This is a pointer to distinguish between explicit zero and not specified. Defaults to 3.`,
				},
				resource.Attribute{
					Name:        "suspend",
					Description: `(Optional) This flag tells the controller to suspend subsequent executions, it does not apply to already started executions. Defaults to false. ### ` + "`" + `job_template` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `(Required) Standard object's metadata of the jobs created from this template. For more info: https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata`,
				},
				resource.Attribute{
					Name:        "spec",
					Description: `(Required) Specification of the desired behavior of the job. For more info: https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#spec-and-status ### ` + "`" + `spec` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "active_deadline_seconds",
					Description: `(Optional) Specifies the duration in seconds relative to the startTime that the job may be active before the system tries to terminate it; value must be positive integer.`,
				},
				resource.Attribute{
					Name:        "backoff_limit",
					Description: `(Optional) Specifies the number of retries before marking this job failed. Defaults to 6`,
				},
				resource.Attribute{
					Name:        "completions",
					Description: `(Optional) Specifies the desired number of successfully finished pods the job should be run with. Setting to nil means that the success of any pod signals the success of all pods, and allows parallelism to have any positive value. Setting to 1 means that parallelism is limited to 1 and the success of that pod signals the success of the job. For more info: https://kubernetes.io/docs/concepts/workloads/controllers/jobs-run-to-completion/`,
				},
				resource.Attribute{
					Name:        "manual_selector",
					Description: `(Optional) Controls generation of pod labels and pod selectors. Leave ` + "`" + `manualSelector` + "`" + ` unset unless you are certain what you are doing. When false or unset, the system pick labels unique to this job and appends those labels to the pod template. When true, the user is responsible for picking unique labels and specifying the selector. Failure to pick a unique label may cause this and other jobs to not function correctly. However, You may see ` + "`" + `manualSelector=true` + "`" + ` in jobs that were created with the old ` + "`" + `extensions/v1beta1` + "`" + ` API. For more info: https://kubernetes.io/docs/concepts/workloads/controllers/jobs-run-to-completion/#specifying-your-own-pod-selector`,
				},
				resource.Attribute{
					Name:        "parallelism",
					Description: `(Optional) Specifies the maximum desired number of pods the job should run at any given time. The actual number of pods running in steady state will be less than this number when ` + "`" + `((.spec.completions - .status.successful) < .spec.parallelism)` + "`" + `, i.e. when the work left to do is less than max parallelism. For more info: https://kubernetes.io/docs/concepts/workloads/controllers/jobs-run-to-completion/`,
				},
				resource.Attribute{
					Name:        "selector",
					Description: `(Optional) A label query over pods that should match the pod count. Normally, the system sets this field for you. For more info: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#label-selectors`,
				},
				resource.Attribute{
					Name:        "template",
					Description: `(Optional) Describes the pod that will be created when executing a job. For more info: https://kubernetes.io/docs/concepts/workloads/controllers/jobs-run-to-completion/`,
				},
				resource.Attribute{
					Name:        "ttl_seconds_after_finished",
					Description: `(Optional) ttlSecondsAfterFinished limits the lifetime of a Job that has finished execution (either Complete or Failed). If this field is set, ttlSecondsAfterFinished after the Job finishes, it is eligible to be automatically deleted. When the Job is being deleted, its lifecycle guarantees (e.g. finalizers) will be honored. If this field is unset, the Job won't be automatically deleted. If this field is set to zero, the Job becomes eligible to be deleted immediately after it finishes. ### ` + "`" + `selector` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "match_expressions",
					Description: `(Optional) A list of label selector requirements. The requirements are ANDed.`,
				},
				resource.Attribute{
					Name:        "match_labels",
					Description: `(Optional) A map of ` + "`" + `{key,value}` + "`" + ` pairs. A single ` + "`" + `{key,value}` + "`" + ` in the matchLabels map is equivalent to an element of matchExpressions, whose key field is "key", the operator is "In", and the values array contains only "value". The requirements are ANDed. ### ` + "`" + `template` + "`" + ` #### Arguments These arguments are the same as the for the ` + "`" + `spec` + "`" + ` block of a Pod. Please see the [Pod resource](pod.html#spec-1) for reference.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "kubernetes_csi_driver",
			Category:         "Resources",
			ShortDescription: `The Container Storage Interface (CSI) is a standard for exposing arbitrary block and file storage systems to containerized workloads on Container Orchestration Systems (COs) like Kubernetes.`,
			Description:      ``,
			Keywords: []string{
				"csi",
				"driver",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "metadata",
					Description: `(Required) Standard CSI driver's metadata. For more info see [Kubernetes reference](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata)`,
				},
				resource.Attribute{
					Name:        "spec",
					Description: `(Required) The Specification of the CSI Driver. ## Nested Blocks ### ` + "`" + `metadata` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "annotations",
					Description: `(Optional) An unstructured key value map stored with the csi driver that may be used to store arbitrary metadata. ~> By default, the provider ignores any annotations whose key names end with`,
				},
				resource.Attribute{
					Name:        "generate_name",
					Description: `(Optional) Prefix, used by the server, to generate a unique name ONLY IF the ` + "`" + `name` + "`" + ` field has not been provided. This value will also be combined with a unique suffix. For more info see [Kubernetes reference](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#idempotency)`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) Map of string keys and values that can be used to organize and categorize (scope and select) the csi driver. May match selectors of replication controllers and services. ~> By default, the provider ignores any labels whose key names end with`,
				},
				resource.Attribute{
					Name:        "attach_required",
					Description: `(Required) Indicates if the CSI volume driver requires an attachment operation.`,
				},
				resource.Attribute{
					Name:        "pod_info_on_mount",
					Description: `(Optional) Indicates that the CSI volume driver requires additional pod information (like podName, podUID, etc.) during mount operations.`,
				},
				resource.Attribute{
					Name:        "volume_lifecycle_modes",
					Description: `(Optional) A list of volume types the CSI volume driver supports. values can be ` + "`" + `Persistent` + "`" + ` and ` + "`" + `Ephemeral` + "`" + `. #### Attributes`,
				},
				resource.Attribute{
					Name:        "generation",
					Description: `A sequence number representing a specific generation of the desired state.`,
				},
				resource.Attribute{
					Name:        "resource_version",
					Description: `An opaque value that represents the internal version of this csi driver that can be used by clients to determine when csi driver has changed. For more info see [Kubernetes reference](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency)`,
				},
				resource.Attribute{
					Name:        "uid",
					Description: `The unique in time and space value for this csi driver. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/identifiers#uids) ## Import kubernetes_csi_driver can be imported using its name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import kubernetes_csi_driver.example terraform-example ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "kubernetes_daemonset",
			Category:         "Resources",
			ShortDescription: `A DaemonSet ensures that all (or some) Nodes run a copy of a Pod. As nodes are added to the cluster, Pods are added to them. As nodes are removed from the cluster, those Pods are garbage collected. Deleting a DaemonSet will clean up the Pods it created.`,
			Description:      ``,
			Keywords: []string{
				"daemonset",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "metadata",
					Description: `(Required) Standard daemonset's metadata. For more info see [Kubernetes reference](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata)`,
				},
				resource.Attribute{
					Name:        "spec",
					Description: `(Required) Spec defines the specification of the desired behavior of the daemonset. For more info see [Kubernetes reference](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#spec-and-status)`,
				},
				resource.Attribute{
					Name:        "wait_for_rollout",
					Description: `(Optional) Wait for the deployment to successfully roll out. Defaults to ` + "`" + `true` + "`" + `. ## Nested Blocks ### ` + "`" + `metadata` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "annotations",
					Description: `(Optional) An unstructured key value map stored with the deployment that may be used to store arbitrary metadata. ~> By default, the provider ignores any annotations whose key names end with`,
				},
				resource.Attribute{
					Name:        "generate_name",
					Description: `(Optional) Prefix, used by the server, to generate a unique name ONLY IF the ` + "`" + `name` + "`" + ` field has not been provided. This value will also be combined with a unique suffix. For more info see [Kubernetes reference](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#idempotency)`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) Map of string keys and values that can be used to organize and categorize (scope and select) the deployment. ~> By default, the provider ignores any labels whose key names end with`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the deployment, must be unique. Cannot be updated. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/identifiers#names)`,
				},
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) Namespace defines the space within which name of the deployment must be unique. #### Attributes`,
				},
				resource.Attribute{
					Name:        "generation",
					Description: `A sequence number representing a specific generation of the desired state.`,
				},
				resource.Attribute{
					Name:        "resource_version",
					Description: `An opaque value that represents the internal version of this deployment that can be used by clients to determine when deployment has changed. For more info see [Kubernetes reference](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency)`,
				},
				resource.Attribute{
					Name:        "uid",
					Description: `The unique in time and space value for this deployment. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/identifiers#uids) ### ` + "`" + `spec` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "min_ready_seconds",
					Description: `(Optional) Minimum number of seconds for which a newly created pod should be ready without any of its container crashing, for it to be considered available. Defaults to 0 (pod will be considered available as soon as it is ready)`,
				},
				resource.Attribute{
					Name:        "revision_history_limit",
					Description: `(Optional) The number of old ReplicaSets to retain to allow rollback. This is a pointer to distinguish between explicit zero and not specified. Defaults to 10.`,
				},
				resource.Attribute{
					Name:        "strategy",
					Description: `(Optional) The update strategy to use to replace existing pods with new ones.`,
				},
				resource.Attribute{
					Name:        "selector",
					Description: `(Optional) A label query over pods that should match the Replicas count. Label keys and values that must match in order to be controlled by this deployment.`,
				},
				resource.Attribute{
					Name:        "template",
					Description: `(Required) Describes the pod that will be created per Node. This takes precedence over a TemplateRef. For more info see [Kubernetes reference](https://kubernetes.io/docs/concepts/workloads/controllers/deployment/#pod-template) ### ` + "`" + `strategy` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type of daemon set update. Can be 'RollingUpdate' or 'OnDelete'. Default is 'RollingUpdate'.`,
				},
				resource.Attribute{
					Name:        "rolling_update",
					Description: `Rolling update config params. Present only if type = 'RollingUpdate'. ### ` + "`" + `rolling_update` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "max_unavailable",
					Description: `The maximum number of DaemonSet pods that can be unavailable during the update. Value can be an absolute number (ex: 5) or a percentage of total number of DaemonSet pods at the start of the update (ex: 10%). Absolute number is calculated from percentage by rounding up. This cannot be 0. Default value is 1. Example: when this is set to 30%, at most 30% of the total number of nodes that should be running the daemon pod (i.e. status.desiredNumberScheduled) can have their pods stopped for an update at any given time. The update starts by stopping at most 30% of those DaemonSet pods and then brings up new DaemonSet pods in their place. Once the new pods are available, it then proceeds onto other DaemonSet pods, thus ensuring that at least 70% of original number of DaemonSet pods are available at all times during the update. ### ` + "`" + `template` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `(Required) Standard object's metadata. For more info see https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata.`,
				},
				resource.Attribute{
					Name:        "spec",
					Description: `(Required) Specification of the desired behavior of the pod. For more info see https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#spec-and-status. ### template ` + "`" + `spec` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "affinity",
					Description: `(Optional) A group of affinity scheduling rules. If specified, the pod will be dispatched by specified scheduler. If not specified, the pod will be dispatched by default scheduler.`,
				},
				resource.Attribute{
					Name:        "active_deadline_seconds",
					Description: `(Optional) Optional duration in seconds the pod may be active on the node relative to StartTime before the system will actively try to mark it failed and kill associated containers. Value must be a positive integer.`,
				},
				resource.Attribute{
					Name:        "automount_service_account_token",
					Description: `(Optional) Indicates whether a service account token should be automatically mounted. Defaults to ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "container",
					Description: `(Optional) List of containers belonging to the pod. Containers cannot currently be added or removed. There must be at least one container in a Pod. Cannot be updated. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/containers)`,
				},
				resource.Attribute{
					Name:        "init_container",
					Description: `(Optional) List of init containers belonging to the pod. Init containers always run to completion and each must complete successfully before the next is started. For more info see [Kubernetes reference](https://kubernetes.io/docs/concepts/workloads/pods/init-containers)/`,
				},
				resource.Attribute{
					Name:        "dns_policy",
					Description: `(Optional) Set DNS policy for containers within the pod. Valid values are 'ClusterFirstWithHostNet', 'ClusterFirst', 'Default' or 'None'. DNS parameters given in DNSConfig will be merged with the policy selected with DNSPolicy. To have DNS options set along with hostNetwork, you have to specify DNS policy explicitly to 'ClusterFirstWithHostNet'. Optional: Defaults to 'ClusterFirst', see [Kubernetes reference](https://kubernetes.io/docs/concepts/services-networking/dns-pod-service/#pod-s-dns-policy).`,
				},
				resource.Attribute{
					Name:        "dns_config",
					Description: `(Optional) Specifies the DNS parameters of a pod. Parameters specified here will be merged to the generated DNS configuration based on DNSPolicy. Defaults to empty. See ` + "`" + `dns_config` + "`" + ` block definition below.`,
				},
				resource.Attribute{
					Name:        "enable_service_links",
					Description: `(Optional) Enables generating environment variables for service discovery. Optional: Defaults to true. For more info see [Kubernetes reference](https://kubernetes.io/docs/concepts/services-networking/connect-applications-service/#accessing-the-service).`,
				},
				resource.Attribute{
					Name:        "host_aliases",
					Description: `(Optional) List of hosts and IPs that will be injected into the pod's hosts file if specified. Optional: Defaults to empty. See ` + "`" + `host_aliases` + "`" + ` block definition below.`,
				},
				resource.Attribute{
					Name:        "host_ipc",
					Description: `(Optional) Use the host's ipc namespace. Optional: Defaults to false.`,
				},
				resource.Attribute{
					Name:        "host_network",
					Description: `(Optional) Host networking requested for this pod. Use the host's network namespace. If this option is set, the ports that will be used must be specified.`,
				},
				resource.Attribute{
					Name:        "host_pid",
					Description: `(Optional) Use the host's pid namespace.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `(Optional) Specifies the hostname of the Pod If not specified, the pod's hostname will be set to a system-defined value.`,
				},
				resource.Attribute{
					Name:        "image_pull_secrets",
					Description: `(Optional) ImagePullSecrets is an optional list of references to secrets in the same namespace to use for pulling any of the images used by this PodSpec. If specified, these secrets will be passed to individual puller implementations for them to use. For example, in the case of docker, only DockerConfig type secrets are honored. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/images#specifying-imagepullsecrets-on-a-pod)`,
				},
				resource.Attribute{
					Name:        "node_name",
					Description: `(Optional) NodeName is a request to schedule this pod onto a specific node. If it is non-empty, the scheduler simply schedules this pod onto that node, assuming that it fits resource requirements.`,
				},
				resource.Attribute{
					Name:        "node_selector",
					Description: `(Optional) NodeSelector is a selector which must be true for the pod to fit on a node. Selector which must match a node's labels for the pod to be scheduled on that node. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/node-selection).`,
				},
				resource.Attribute{
					Name:        "priority_class_name",
					Description: `(Optional) If specified, indicates the pod's priority. 'system-node-critical' and 'system-cluster-critical' are two special keywords which indicate the highest priorities with the formerer being the highest priority. Any other name must be defined by creating a PriorityClass object with that name. If not specified, the pod priority will be default or zero if there is no default.`,
				},
				resource.Attribute{
					Name:        "restart_policy",
					Description: `(Optional) Restart policy for all containers within the pod. One of Always, OnFailure, Never. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/pod-states#restartpolicy).`,
				},
				resource.Attribute{
					Name:        "security_context",
					Description: `(Optional) SecurityContext holds pod-level security attributes and common container settings. Optional: Defaults to empty`,
				},
				resource.Attribute{
					Name:        "service_account_name",
					Description: `(Optional) ServiceAccountName is the name of the ServiceAccount to use to run this pod. For more info see https://kubernetes.io/docs/reference/access-authn-authz/service-accounts-admin/.`,
				},
				resource.Attribute{
					Name:        "share_process_namespace",
					Description: `(Optional) Share a single process namespace between all of the containers in a pod. When this is set containers will be able to view and signal processes from other containers in the same pod, and the first process in each container will not be assigned PID 1. HostPID and ShareProcessNamespace cannot both be set.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `(Optional) If specified, the fully qualified Pod hostname will be "...svc.". If not specified, the pod will not have a domainname at all..`,
				},
				resource.Attribute{
					Name:        "termination_grace_period_seconds",
					Description: `(Optional) Optional duration in seconds the pod needs to terminate gracefully. May be decreased in delete request. Value must be non-negative integer. The value zero indicates delete immediately. If this value is nil, the default grace period will be used instead. The grace period is the duration in seconds after the processes running in the pod are sent a termination signal and the time when the processes are forcibly halted with a kill signal. Set this value longer than the expected cleanup time for your process.`,
				},
				resource.Attribute{
					Name:        "toleration",
					Description: `(Optional) Optional pod node tolerations. For more info see [Kubernetes reference](https://kubernetes.io/docs/concepts/configuration/taint-and-toleration/)`,
				},
				resource.Attribute{
					Name:        "volume",
					Description: `(Optional) List of volumes that can be mounted by containers belonging to the pod. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/volumes) ### ` + "`" + `affinity` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "node_affinity",
					Description: `(Optional) Node affinity scheduling rules for the pod. For more info see [Kubernetes reference](https://kubernetes.io/docs/concepts/configuration/assign-pod-node/#node-affinity-beta-feature)`,
				},
				resource.Attribute{
					Name:        "pod_affinity",
					Description: `(Optional) Inter-pod topological affinity. rules that specify that certain pods should be placed in the same topological domain (e.g. same node, same rack, same zone, same power domain, etc.) For more info see [Kubernetes reference](https://kubernetes.io/docs/concepts/configuration/assign-pod-node/#inter-pod-affinity-and-anti-affinity-beta-feature)`,
				},
				resource.Attribute{
					Name:        "pod_anti_affinity",
					Description: `(Optional) Inter-pod topological affinity. rules that specify that certain pods should be placed in the same topological domain (e.g. same node, same rack, same zone, same power domain, etc.) For more info see [Kubernetes reference](https://kubernetes.io/docs/concepts/configuration/assign-pod-node/#inter-pod-affinity-and-anti-affinity-beta-feature) ### ` + "`" + `node_affinity` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "required_during_scheduling_ignored_during_execution",
					Description: `(Optional) If the affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to an update), the system may or may not try to eventually evict the pod from its node.`,
				},
				resource.Attribute{
					Name:        "preferred_during_scheduling_ignored_during_execution",
					Description: `(Optional) The scheduler will prefer to schedule pods to nodes that satisfy the affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. ### ` + "`" + `required_during_scheduling_ignored_during_execution` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "node_selector_term",
					Description: `(Required) A list of node selector terms. The terms are ORed. ## ` + "`" + `node_selector_term` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "match_expressions",
					Description: `(Optional) A list of node selector requirements by node's labels.`,
				},
				resource.Attribute{
					Name:        "match_fields",
					Description: `(Optional) A list of node selector requirements by node's fields. ### ` + "`" + `match_expressions` + "`" + ` / ` + "`" + `match_fields` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) The label key that the selector applies to.`,
				},
				resource.Attribute{
					Name:        "operator",
					Description: `(Required) Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Optional) An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. ### ` + "`" + `preferred_during_scheduling_ignored_during_execution` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "preference",
					Description: `(Required) A node selector term, associated with the corresponding weight.`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `(Required) Weight associated with matching the corresponding nodeSelectorTerm, in the range 1-100. ### ` + "`" + `preference` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "match_expressions",
					Description: `(Optional) A list of node selector requirements by node's labels.`,
				},
				resource.Attribute{
					Name:        "match_fields",
					Description: `(Optional) A list of node selector requirements by node's fields. ## ` + "`" + `match_expressions` + "`" + ` / ` + "`" + `match_fields` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) The label key that the selector applies to.`,
				},
				resource.Attribute{
					Name:        "operator",
					Description: `(Required) Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Optional) An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. ### ` + "`" + `pod_affinity` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "required_during_scheduling_ignored_during_execution",
					Description: `(Optional) If the affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to a pod label update), the system may or may not try to eventually evict the pod from its node.`,
				},
				resource.Attribute{
					Name:        "preferred_during_scheduling_ignored_during_execution",
					Description: `(Optional) The scheduler will prefer to schedule pods to nodes that satisfy the affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. ### ` + "`" + `pod_anti_affinity` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "required_during_scheduling_ignored_during_execution",
					Description: `(Optional) If the anti-affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the anti-affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to a pod label update), the system may or may not try to eventually evict the pod from its node.`,
				},
				resource.Attribute{
					Name:        "preferred_during_scheduling_ignored_during_execution",
					Description: `(Optional) The scheduler will prefer to schedule pods to nodes that satisfy the anti-affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. ### ` + "`" + `required_during_scheduling_ignored_during_execution` + "`" + ` (pod_affinity_term) #### Arguments`,
				},
				resource.Attribute{
					Name:        "label_selector",
					Description: `(Optional) A label query over a set of resources, in this case pods.`,
				},
				resource.Attribute{
					Name:        "namespaces",
					Description: `(Optional) Specifies which namespaces the ` + "`" + `label_selector` + "`" + ` applies to (matches against). Null or empty list means "this pod's namespace"`,
				},
				resource.Attribute{
					Name:        "topology_key",
					Description: `(Optional) This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the ` + "`" + `label_selector` + "`" + ` in the specified namespaces, where co-located is defined as running on a node whose value of the label with key ` + "`" + `topology_key` + "`" + ` matches that of any node on which any of the selected pods is running. Empty ` + "`" + `topology_key` + "`" + ` is not allowed. ### ` + "`" + `preferred_during_scheduling_ignored_during_execution` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "pod_affinity_term",
					Description: `(Required) A pod affinity term, associated with the corresponding weight.`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `(Required) Weight associated with matching the corresponding ` + "`" + `pod_affinity_term` + "`" + `, in the range 1-100. ### ` + "`" + `container` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "args",
					Description: `(Optional) Arguments to the entrypoint. The docker image's CMD is used if this is not provided. Variable references $(VAR_NAME) are expanded using the container's environment. If a variable cannot be resolved, the reference in the input string will be unchanged. The $(VAR_NAME) syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped references will never be expanded, regardless of whether the variable exists or not. Cannot be updated. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/containers#containers-and-commands)`,
				},
				resource.Attribute{
					Name:        "command",
					Description: `(Optional) Entrypoint array. Not executed within a shell. The docker image's ENTRYPOINT is used if this is not provided. Variable references $(VAR_NAME) are expanded using the container's environment. If a variable cannot be resolved, the reference in the input string will be unchanged. The $(VAR_NAME) syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped references will never be expanded, regardless of whether the variable exists or not. Cannot be updated. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/containers#containers-and-commands)`,
				},
				resource.Attribute{
					Name:        "env",
					Description: `(Optional) Block of string name and value pairs to set in the container's environment. May be declared multiple times. Cannot be updated.`,
				},
				resource.Attribute{
					Name:        "env_from",
					Description: `(Optional) List of sources to populate environment variables in the container. The keys defined within a source must be a C_IDENTIFIER. All invalid keys will be reported as an event when the container is starting. When a key exists in multiple sources, the value associated with the last source will take precedence. Values defined by an Env with a duplicate key will take precedence. Cannot be updated.`,
				},
				resource.Attribute{
					Name:        "image",
					Description: `(Optional) Docker image name. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/images)`,
				},
				resource.Attribute{
					Name:        "image_pull_policy",
					Description: `(Optional) Image pull policy. One of Always, Never, IfNotPresent. Defaults to Always if :latest tag is specified, or IfNotPresent otherwise. Cannot be updated. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/images#updating-images)`,
				},
				resource.Attribute{
					Name:        "lifecycle",
					Description: `(Optional) Actions that the management system should take in response to container lifecycle events`,
				},
				resource.Attribute{
					Name:        "liveness_probe",
					Description: `(Optional) Periodic probe of container liveness. Container will be restarted if the probe fails. Cannot be updated. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/pod-states#container-probes)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the container specified as a DNS_LABEL. Each container in a pod must have a unique name (DNS_LABEL). Cannot be updated.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional) List of ports to expose from the container. Exposing a port here gives the system additional information about the network connections a container uses, but is primarily informational. Not specifying a port here DOES NOT prevent that port from being exposed. Any port which is listening on the default "0.0.0.0" address inside a container will be accessible from the network. Cannot be updated.`,
				},
				resource.Attribute{
					Name:        "readiness_probe",
					Description: `(Optional) Periodic probe of container service readiness. Container will be removed from service endpoints if the probe fails. Cannot be updated. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/pod-states#container-probes)`,
				},
				resource.Attribute{
					Name:        "resources",
					Description: `(Optional) Compute Resources required by this container. Cannot be updated. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/persistent-volumes#resources)`,
				},
				resource.Attribute{
					Name:        "security_context",
					Description: `(Optional) Security options the pod should run with. For more info see https://kubernetes.io/docs/tasks/configure-pod-container/security-context/.`,
				},
				resource.Attribute{
					Name:        "startup_probe",
					Description: `(Optional) StartupProbe indicates that the Pod has successfully initialized. If specified, no other probes are executed until this completes successfully. If this probe fails, the Pod will be restarted, just as if the livenessProbe failed. This can be used to provide different probe parameters at the beginning of a Pod's lifecycle, when it might take a long time to load data or warm a cache, than during steady-state operation. This cannot be updated. For more info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes`,
				},
				resource.Attribute{
					Name:        "stdin",
					Description: `(Optional) Whether this container should allocate a buffer for stdin in the container runtime. If this is not set, reads from stdin in the container will always result in EOF.`,
				},
				resource.Attribute{
					Name:        "stdin_once",
					Description: `(Optional) Whether the container runtime should close the stdin channel after it has been opened by a single attach. When stdin is true the stdin stream will remain open across multiple attach sessions. If stdinOnce is set to true, stdin is opened on container start, is empty until the first client attaches to stdin, and then remains open and accepts data until the client disconnects, at which time stdin is closed and remains closed until the container is restarted. If this flag is false, a container processes that reads from stdin will never receive an EOF.`,
				},
				resource.Attribute{
					Name:        "termination_message_path",
					Description: `(Optional) Optional: Path at which the file to which the container's termination message will be written is mounted into the container's filesystem. Message written is intended to be brief final status, such as an assertion failure message. Defaults to /dev/termination-log. Cannot be updated.`,
				},
				resource.Attribute{
					Name:        "tty",
					Description: `(Optional) Whether this container should allocate a TTY for itself`,
				},
				resource.Attribute{
					Name:        "volume_mount",
					Description: `(Optional) Pod volumes to mount into the container's filesystem. Cannot be updated.`,
				},
				resource.Attribute{
					Name:        "working_dir",
					Description: `(Optional) Container's working directory. If not specified, the container runtime's default will be used, which might be configured in the container image. Cannot be updated. ### ` + "`" + `aws_elastic_block_store` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "fs_type",
					Description: `(Optional) Filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/volumes#awselasticblockstore)`,
				},
				resource.Attribute{
					Name:        "partition",
					Description: `(Optional) The partition in the volume that you want to mount. If omitted, the default is to mount by volume name. Examples: For volume /dev/sda1, you specify the partition as "1". Similarly, the volume partition for /dev/sda is "0" (or you can leave the property empty).`,
				},
				resource.Attribute{
					Name:        "read_only",
					Description: `(Optional) Whether to set the read-only property in VolumeMounts to "true". If omitted, the default is "false". For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/volumes#awselasticblockstore)`,
				},
				resource.Attribute{
					Name:        "volume_id",
					Description: `(Required) Unique ID of the persistent disk resource in AWS (Amazon EBS volume). For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/volumes#awselasticblockstore) ### ` + "`" + `azure_disk` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "caching_mode",
					Description: `(Required) Host Caching mode: None, Read Only, Read Write.`,
				},
				resource.Attribute{
					Name:        "data_disk_uri",
					Description: `(Required) The URI the data disk in the blob storage`,
				},
				resource.Attribute{
					Name:        "disk_name",
					Description: `(Required) The Name of the data disk in the blob storage`,
				},
				resource.Attribute{
					Name:        "fs_type",
					Description: `(Optional) Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.`,
				},
				resource.Attribute{
					Name:        "read_only",
					Description: `(Optional) Whether to force the read-only setting in VolumeMounts. Defaults to false (read/write). ### ` + "`" + `azure_file` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "read_only",
					Description: `(Optional) Whether to force the read-only setting in VolumeMounts. Defaults to false (read/write).`,
				},
				resource.Attribute{
					Name:        "secret_name",
					Description: `(Required) The name of secret that contains Azure Storage Account Name and Key`,
				},
				resource.Attribute{
					Name:        "share_name",
					Description: `(Required) Share Name ### ` + "`" + `capabilities` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "add",
					Description: `(Optional) Added capabilities`,
				},
				resource.Attribute{
					Name:        "drop",
					Description: `(Optional) Removed capabilities ### ` + "`" + `ceph_fs` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "monitors",
					Description: `(Required) Monitors is a collection of Ceph monitors. For more info see https://github.com/kubernetes/examples/tree/master/volumes/cephfs/#how-to-use-it.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Optional) Used as the mounted root, rather than the full Ceph tree, default is /`,
				},
				resource.Attribute{
					Name:        "read_only",
					Description: `(Optional) Whether to force the read-only setting in VolumeMounts. Defaults to ` + "`" + `false` + "`" + ` (read/write). For more info see https://github.com/kubernetes/examples/tree/master/volumes/cephfs/#how-to-use-it.`,
				},
				resource.Attribute{
					Name:        "secret_file",
					Description: `(Optional) The path to key ring for User, default is /etc/ceph/user.secret. For more info see https://github.com/kubernetes/examples/tree/master/volumes/cephfs/#how-to-use-it.`,
				},
				resource.Attribute{
					Name:        "secret_ref",
					Description: `(Optional) Reference to the authentication secret for User, default is empty. For more info see https://github.com/kubernetes/examples/tree/master/volumes/cephfs/#how-to-use-it.`,
				},
				resource.Attribute{
					Name:        "user",
					Description: `(Optional) User is the rados user name, default is admin. For more info see https://github.com/kubernetes/examples/tree/master/volumes/cephfs/#how-to-use-it. ### ` + "`" + `cinder` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "fs_type",
					Description: `(Optional) Filesystem type to mount. Must be a filesystem type supported by the host operating system. Examples: "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified. For more info see https://github.com/kubernetes/examples/blob/master/mysql-cinder-pd/README.md#mysql-installation-with-cinder-volume-plugin.`,
				},
				resource.Attribute{
					Name:        "read_only",
					Description: `(Optional) Whether to force the read-only setting in VolumeMounts. Defaults to false (read/write). For more info see https://github.com/kubernetes/examples/blob/master/mysql-cinder-pd/README.md#mysql-installation-with-cinder-volume-plugin.`,
				},
				resource.Attribute{
					Name:        "volume_id",
					Description: `(Required) Volume ID used to identify the volume in Cinder. For more info see https://github.com/kubernetes/examples/blob/master/mysql-cinder-pd/README.md#mysql-installation-with-cinder-volume-plugin. ### ` + "`" + `config_map` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "default_mode",
					Description: `(Optional) Optional: mode bits to use on created files by default. Must be a value between 0 and 0777. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.`,
				},
				resource.Attribute{
					Name:        "items",
					Description: `(Optional) If unspecified, each key-value pair in the Data field of the referenced ConfigMap will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the ConfigMap, the volume setup will error unless it is marked ` + "`" + `optional` + "`" + `. Paths must be relative and may not contain the '..' path or start with '..'.`,
				},
				resource.Attribute{
					Name:        "optional",
					Description: `(Optional) Specify whether the ConfigMap or its keys must be defined.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the referent. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/identifiers#names) ### ` + "`" + `config_map_ref` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the referent. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/identifiers#names)`,
				},
				resource.Attribute{
					Name:        "optional",
					Description: `(Optional) Specify whether the ConfigMap must be defined ### ` + "`" + `config_map_key_ref` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Optional) The key to select.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the referent. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/identifiers#names)`,
				},
				resource.Attribute{
					Name:        "optional",
					Description: `(Optional) Specify whether the ConfigMap or its key must be defined ### ` + "`" + `dns_config` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "nameservers",
					Description: `(Optional) A list of DNS name server IP addresses specified as strings. This will be appended to the base nameservers generated from DNSPolicy. Duplicated nameservers will be removed. Optional: Defaults to empty.`,
				},
				resource.Attribute{
					Name:        "option",
					Description: `(Optional) A list of DNS resolver options specified as blocks with ` + "`" + `name` + "`" + `/` + "`" + `value` + "`" + ` pairs. This will be merged with the base options generated from DNSPolicy. Duplicated entries will be removed. Resolution options given in Options will override those that appear in the base DNSPolicy. Optional: Defaults to empty.`,
				},
				resource.Attribute{
					Name:        "searches",
					Description: `(Optional) A list of DNS search domains for host-name lookup specified as strings. This will be appended to the base search paths generated from DNSPolicy. Duplicated search paths will be removed. Optional: Defaults to empty. The ` + "`" + `option` + "`" + ` block supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the option.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Optional) Value of the option. Optional: Defaults to empty. ### ` + "`" + `downward_api` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "default_mode",
					Description: `(Optional) Optional: mode bits to use on created files by default. Must be a value between 0 and 0777. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.`,
				},
				resource.Attribute{
					Name:        "items",
					Description: `(Optional) If unspecified, each key-value pair in the Data field of the referenced ConfigMap will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the ConfigMap, the volume setup will error. Paths must be relative and may not contain the '..' path or start with '..'. ### ` + "`" + `empty_dir` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "medium",
					Description: `(Optional) What type of storage medium should back this directory. The default is "" which means to use the node's default medium. Must be an empty string (default) or Memory. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/volumes#emptydir)`,
				},
				resource.Attribute{
					Name:        "size_limit",
					Description: `(Optional) Total amount of local storage required for this EmptyDir volume. For more info see [Kubernetes reference](https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/#resource-units-in-kubernetes) and [Kubernetes Quantity type](https://pkg.go.dev/k8s.io/apimachinery/pkg/api/resource?tab=doc#Quantity). ### ` + "`" + `env` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the environment variable. Must be a C_IDENTIFIER`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Optional) Variable references $(VAR_NAME) are expanded using the previous defined environment variables in the container and any service environment variables. If a variable cannot be resolved, the reference in the input string will be unchanged. The $(VAR_NAME) syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped references will never be expanded, regardless of whether the variable exists or not. Defaults to "".`,
				},
				resource.Attribute{
					Name:        "value_from",
					Description: `(Optional) Source for the environment variable's value ### ` + "`" + `env_from` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "config_map_ref",
					Description: `(Optional) The ConfigMap to select from`,
				},
				resource.Attribute{
					Name:        "prefix",
					Description: `(Optional) An optional identifer to prepend to each key in the ConfigMap. Must be a C_IDENTIFIER..`,
				},
				resource.Attribute{
					Name:        "secret_ref",
					Description: `(Optional) The Secret to select from ### ` + "`" + `exec` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "command",
					Description: `(Optional) Command is the command line to execute inside the container, the working directory for the command is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy. ### ` + "`" + `fc` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "fs_type",
					Description: `(Optional) Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.`,
				},
				resource.Attribute{
					Name:        "lun",
					Description: `(Required) FC target lun number`,
				},
				resource.Attribute{
					Name:        "read_only",
					Description: `(Optional) Whether to force the read-only setting in VolumeMounts. Defaults to false (read/write).`,
				},
				resource.Attribute{
					Name:        "target_ww_ns",
					Description: `(Required) FC target worldwide names (WWNs) ### ` + "`" + `field_ref` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "api_version",
					Description: `(Optional) Version of the schema the FieldPath is written in terms of, defaults to "v1".`,
				},
				resource.Attribute{
					Name:        "field_path",
					Description: `(Optional) Path of the field to select in the specified API version ### ` + "`" + `flex_volume` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "driver",
					Description: `(Required) Driver is the name of the driver to use for this volume.`,
				},
				resource.Attribute{
					Name:        "fs_type",
					Description: `(Optional) Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs", "ntfs". The default filesystem depends on FlexVolume script.`,
				},
				resource.Attribute{
					Name:        "options",
					Description: `(Optional) Extra command options if any.`,
				},
				resource.Attribute{
					Name:        "read_only",
					Description: `(Optional) Whether to force the ReadOnly setting in VolumeMounts. Defaults to false (read/write).`,
				},
				resource.Attribute{
					Name:        "secret_ref",
					Description: `(Optional) Reference to the secret object containing sensitive information to pass to the plugin scripts. This may be empty if no secret object is specified. If the secret object contains more than one secret, all secrets are passed to the plugin scripts. ### ` + "`" + `flocker` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "dataset_name",
					Description: `(Optional) Name of the dataset stored as metadata -> name on the dataset for Flocker should be considered as deprecated`,
				},
				resource.Attribute{
					Name:        "dataset_uuid",
					Description: `(Optional) UUID of the dataset. This is unique identifier of a Flocker dataset ### ` + "`" + `gce_persistent_disk` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "fs_type",
					Description: `(Optional) Filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/volumes#gcepersistentdisk)`,
				},
				resource.Attribute{
					Name:        "partition",
					Description: `(Optional) The partition in the volume that you want to mount. If omitted, the default is to mount by volume name. Examples: For volume /dev/sda1, you specify the partition as "1". Similarly, the volume partition for /dev/sda is "0" (or you can leave the property empty). For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/volumes#gcepersistentdisk)`,
				},
				resource.Attribute{
					Name:        "pd_name",
					Description: `(Required) Unique name of the PD resource in GCE. Used to identify the disk in GCE. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/volumes#gcepersistentdisk)`,
				},
				resource.Attribute{
					Name:        "read_only",
					Description: `(Optional) Whether to force the ReadOnly setting in VolumeMounts. Defaults to false. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/volumes#gcepersistentdisk) ### ` + "`" + `git_repo` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "directory",
					Description: `(Optional) Target directory name. Must not contain or start with '..'. If '.' is supplied, the volume directory will be the git repository. Otherwise, if specified, the volume will contain the git repository in the subdirectory with the given name.`,
				},
				resource.Attribute{
					Name:        "repository",
					Description: `(Optional) Repository URL`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `(Optional) Commit hash for the specified revision. ### ` + "`" + `glusterfs` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "endpoints_name",
					Description: `(Required) The endpoint name that details Glusterfs topology. For more info see https://github.com/kubernetes/examples/tree/master/volumes/glusterfs#create-a-pod.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Required) The Glusterfs volume path. For more info see https://github.com/kubernetes/examples/tree/master/volumes/glusterfs#create-a-pod.`,
				},
				resource.Attribute{
					Name:        "read_only",
					Description: `(Optional) Whether to force the Glusterfs volume to be mounted with read-only permissions. Defaults to false. For more info see https://github.com/kubernetes/examples/tree/master/volumes/glusterfs#create-a-pod. ### ` + "`" + `host_aliases` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "hostnames",
					Description: `(Required) Array of hostnames for the IP address.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `(Required) IP address of the host file entry. ### ` + "`" + `host_path` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Optional) Path of the directory on the host. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/volumes#hostpath)`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) Type for HostPath volume. Defaults to "". For more info see [Kubernetes reference](https://kubernetes.io/docs/concepts/storage/volumes#hostpath) ### ` + "`" + `http_get` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `(Optional) Host name to connect to, defaults to the pod IP. You probably want to set "Host" in httpHeaders instead.`,
				},
				resource.Attribute{
					Name:        "http_header",
					Description: `(Optional) Scheme to use for connecting to the host.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Optional) Path to access on the HTTP server.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional) Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.`,
				},
				resource.Attribute{
					Name:        "scheme",
					Description: `(Optional) Scheme to use for connecting to the host. ### ` + "`" + `http_header` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The header field name`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Optional) The header field value ### ` + "`" + `image_pull_secrets` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the referent. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/identifiers#names) ### ` + "`" + `iscsi` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "fs_type",
					Description: `(Optional) Filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/volumes#iscsi)`,
				},
				resource.Attribute{
					Name:        "iqn",
					Description: `(Required) Target iSCSI Qualified Name.`,
				},
				resource.Attribute{
					Name:        "iscsi_interface",
					Description: `(Optional) iSCSI interface name that uses an iSCSI transport. Defaults to 'default' (tcp).`,
				},
				resource.Attribute{
					Name:        "lun",
					Description: `(Optional) iSCSI target lun number.`,
				},
				resource.Attribute{
					Name:        "read_only",
					Description: `(Optional) Whether to force the read-only setting in VolumeMounts. Defaults to false.`,
				},
				resource.Attribute{
					Name:        "target_portal",
					Description: `(Required) iSCSI target portal. The portal is either an IP or ip_addr:port if the port is other than default (typically TCP ports 860 and 3260). ### ` + "`" + `items` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Optional) The key to project.`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `(Optional) Optional: mode bits to use on this file, must be a value between 0 and 0777. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Optional) The relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'. ### ` + "`" + `lifecycle` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "post_start",
					Description: `(Optional) post_start is called immediately after a container is created. If the handler fails, the container is terminated and restarted according to its restart policy. Other management of the container blocks until the hook completes. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/container-environment#hook-details)`,
				},
				resource.Attribute{
					Name:        "pre_stop",
					Description: `(Optional) pre_stop is called immediately before a container is terminated. The container is terminated after the handler completes. The reason for termination is passed to the handler. Regardless of the outcome of the handler, the container is eventually terminated. Other management of the container blocks until the hook completes. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/container-environment#hook-details) ### ` + "`" + `liveness_probe` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "exec",
					Description: `(Optional) exec specifies the action to take.`,
				},
				resource.Attribute{
					Name:        "failure_threshold",
					Description: `(Optional) Minimum consecutive failures for the probe to be considered failed after having succeeded.`,
				},
				resource.Attribute{
					Name:        "http_get",
					Description: `(Optional) Specifies the http request to perform.`,
				},
				resource.Attribute{
					Name:        "initial_delay_seconds",
					Description: `(Optional) Number of seconds after the container has started before liveness probes are initiated. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/pod-states#container-probes)`,
				},
				resource.Attribute{
					Name:        "period_seconds",
					Description: `(Optional) How often (in seconds) to perform the probe`,
				},
				resource.Attribute{
					Name:        "success_threshold",
					Description: `(Optional) Minimum consecutive successes for the probe to be considered successful after having failed.`,
				},
				resource.Attribute{
					Name:        "tcp_socket",
					Description: `(Optional) TCPSocket specifies an action involving a TCP port. TCP hooks not yet supported`,
				},
				resource.Attribute{
					Name:        "timeout_seconds",
					Description: `(Optional) Number of seconds after which the probe times out. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/pod-states#container-probes) ### ` + "`" + `nfs` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Required) Path that is exported by the NFS server. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/volumes#nfs)`,
				},
				resource.Attribute{
					Name:        "read_only",
					Description: `(Optional) Whether to force the NFS export to be mounted with read-only permissions. Defaults to false. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/volumes#nfs)`,
				},
				resource.Attribute{
					Name:        "server",
					Description: `(Required) Server is the hostname or IP address of the NFS server. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/volumes#nfs) ### ` + "`" + `persistent_volume_claim` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "claim_name",
					Description: `(Optional) ClaimName is the name of a PersistentVolumeClaim in the same`,
				},
				resource.Attribute{
					Name:        "read_only",
					Description: `(Optional) Will force the ReadOnly setting in VolumeMounts. ### ` + "`" + `photon_persistent_disk` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "fs_type",
					Description: `(Optional) Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.`,
				},
				resource.Attribute{
					Name:        "pd_id",
					Description: `(Required) ID that identifies Photon Controller persistent disk ### ` + "`" + `port` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "container_port",
					Description: `(Required) Number of port to expose on the pod's IP address. This must be a valid port number, 0 < x < 65536.`,
				},
				resource.Attribute{
					Name:        "host_ip",
					Description: `(Optional) What host IP to bind the external port to.`,
				},
				resource.Attribute{
					Name:        "host_port",
					Description: `(Optional) Number of port to expose on the host. If specified, this must be a valid port number, 0 < x < 65536. If HostNetwork is specified, this must match ContainerPort. Most containers do not need this.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) If specified, this must be an IANA_SVC_NAME and unique within the pod. Each named port in a pod must have a unique name. Name for the port that can be referred to by services`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Optional) Protocol for port. Must be UDP or TCP. Defaults to "TCP". ### ` + "`" + `post_start` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "exec",
					Description: `(Optional) exec specifies the action to take.`,
				},
				resource.Attribute{
					Name:        "http_get",
					Description: `(Optional) Specifies the http request to perform.`,
				},
				resource.Attribute{
					Name:        "tcp_socket",
					Description: `(Optional) TCPSocket specifies an action involving a TCP port. TCP hooks not yet supported ### ` + "`" + `pre_stop` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "exec",
					Description: `(Optional) exec specifies the action to take.`,
				},
				resource.Attribute{
					Name:        "http_get",
					Description: `(Optional) Specifies the http request to perform.`,
				},
				resource.Attribute{
					Name:        "tcp_socket",
					Description: `(Optional) TCPSocket specifies an action involving a TCP port. TCP hooks not yet supported ### ` + "`" + `quobyte` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "group",
					Description: `(Optional) Group to map volume access to Default is no group`,
				},
				resource.Attribute{
					Name:        "read_only",
					Description: `(Optional) Whether to force the Quobyte volume to be mounted with read-only permissions. Defaults to false.`,
				},
				resource.Attribute{
					Name:        "registry",
					Description: `(Required) Registry represents a single or multiple Quobyte Registry services specified as a string as host:port pair (multiple entries are separated with commas) which acts as the central registry for volumes`,
				},
				resource.Attribute{
					Name:        "user",
					Description: `(Optional) User to map volume access to Defaults to serivceaccount user`,
				},
				resource.Attribute{
					Name:        "volume",
					Description: `(Required) Volume is a string that references an already created Quobyte volume by name. ### ` + "`" + `rbd` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "ceph_monitors",
					Description: `(Required) A collection of Ceph monitors. For more info see https://github.com/kubernetes/examples/tree/master/volumes/rbd#how-to-use-it.`,
				},
				resource.Attribute{
					Name:        "fs_type",
					Description: `(Optional) Filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/volumes#rbd)`,
				},
				resource.Attribute{
					Name:        "keyring",
					Description: `(Optional) Keyring is the path to key ring for RBDUser. Default is /etc/ceph/keyring. For more info see https://github.com/kubernetes/examples/tree/master/volumes/rbd#how-to-use-it.`,
				},
				resource.Attribute{
					Name:        "rados_user",
					Description: `(Optional) The rados user name. Default is admin. For more info see https://github.com/kubernetes/examples/tree/master/volumes/rbd#how-to-use-it.`,
				},
				resource.Attribute{
					Name:        "rbd_image",
					Description: `(Required) The rados image name. For more info see https://github.com/kubernetes/examples/tree/master/volumes/rbd#how-to-use-it.`,
				},
				resource.Attribute{
					Name:        "rbd_pool",
					Description: `(Optional) The rados pool name. Default is rbd. For more info see https://github.com/kubernetes/examples/tree/master/volumes/rbd#how-to-use-it.`,
				},
				resource.Attribute{
					Name:        "read_only",
					Description: `(Optional) Whether to force the read-only setting in VolumeMounts. Defaults to false. For more info see https://github.com/kubernetes/examples/tree/master/volumes/rbd#how-to-use-it.`,
				},
				resource.Attribute{
					Name:        "secret_ref",
					Description: `(Optional) Name of the authentication secret for RBDUser. If provided overrides keyring. Default is nil. For more info see https://github.com/kubernetes/examples/tree/master/volumes/rbd#how-to-use-it. ### ` + "`" + `readiness_probe` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "exec",
					Description: `(Optional) exec specifies the action to take.`,
				},
				resource.Attribute{
					Name:        "failure_threshold",
					Description: `(Optional) Minimum consecutive failures for the probe to be considered failed after having succeeded.`,
				},
				resource.Attribute{
					Name:        "http_get",
					Description: `(Optional) Specifies the http request to perform.`,
				},
				resource.Attribute{
					Name:        "initial_delay_seconds",
					Description: `(Optional) Number of seconds after the container has started before liveness probes are initiated. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/pod-states#container-probes)`,
				},
				resource.Attribute{
					Name:        "period_seconds",
					Description: `(Optional) How often (in seconds) to perform the probe`,
				},
				resource.Attribute{
					Name:        "success_threshold",
					Description: `(Optional) Minimum consecutive successes for the probe to be considered successful after having failed.`,
				},
				resource.Attribute{
					Name:        "tcp_socket",
					Description: `(Optional) TCPSocket specifies an action involving a TCP port. TCP hooks not yet supported`,
				},
				resource.Attribute{
					Name:        "timeout_seconds",
					Description: `(Optional) Number of seconds after which the probe times out. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/pod-states#container-probes) ### ` + "`" + `resources` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "limits",
					Description: `(Optional) Describes the maximum amount of compute resources allowed. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/compute-resources)/`,
				},
				resource.Attribute{
					Name:        "requests",
					Description: `(Optional) Describes the minimum amount of compute resources required. ### ` + "`" + `resource_field_ref` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "container_name",
					Description: `(Optional) The name of the container`,
				},
				resource.Attribute{
					Name:        "resource",
					Description: `(Required) Resource to select`,
				},
				resource.Attribute{
					Name:        "divisor",
					Description: `(Optional) Specifies the output format of the exposed resources, defaults to "1". ### ` + "`" + `se_linux_options` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "level",
					Description: `(Optional) Level is SELinux level label that applies to the container.`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `(Optional) Role is a SELinux role label that applies to the container.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) Type is a SELinux type label that applies to the container.`,
				},
				resource.Attribute{
					Name:        "user",
					Description: `(Optional) User is a SELinux user label that applies to the container. ### ` + "`" + `secret` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "default_mode",
					Description: `(Optional) Mode bits to use on created files by default. Must be a value between 0 and 0777. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.`,
				},
				resource.Attribute{
					Name:        "items",
					Description: `(Optional) List of Secret Items to project into the volume. See ` + "`" + `items` + "`" + ` block definition below. If unspecified, each key-value pair in the Data field of the referenced Secret will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the Secret, the volume setup will error unless it is marked ` + "`" + `optional` + "`" + `. Paths must be relative and may not contain the '..' path or start with '..'.`,
				},
				resource.Attribute{
					Name:        "optional",
					Description: `(Optional) Specify whether the Secret or its keys must be defined.`,
				},
				resource.Attribute{
					Name:        "secret_name",
					Description: `(Optional) Name of the secret in the pod's namespace to use. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/volumes#secrets) The ` + "`" + `items` + "`" + ` block supports the following:`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) The key to project.`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `(Optional) Mode bits to use on this file, must be a value between 0 and 0777. If not specified, the volume defaultMode will be used.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Required) The relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'. ### ` + "`" + `secret_ref` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the referent. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/identifiers#names)`,
				},
				resource.Attribute{
					Name:        "optional",
					Description: `(Optional) Specify whether the Secret must be defined ### ` + "`" + `secret_key_ref` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Optional) The key of the secret to select from. Must be a valid secret key.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the referent. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/identifiers#names)`,
				},
				resource.Attribute{
					Name:        "optional",
					Description: `(Optional) Specify whether the Secret or its key must be defined ### ` + "`" + `secret_ref` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the referent. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/identifiers#names) ### container ` + "`" + `security_context` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "allow_privilege_escalation",
					Description: `(Optional) AllowPrivilegeEscalation controls whether a process can gain more privileges than its parent process. This bool directly controls if the no_new_privs flag will be set on the container process. AllowPrivilegeEscalation is true always when the container is: 1) run as Privileged 2) has CAP_SYS_ADMIN`,
				},
				resource.Attribute{
					Name:        "capabilities",
					Description: `(Optional) The capabilities to add/drop when running containers. Defaults to the default set of capabilities granted by the container runtime.`,
				},
				resource.Attribute{
					Name:        "privileged",
					Description: `(Optional) Run container in privileged mode. Processes in privileged containers are essentially equivalent to root on the host. Defaults to false.`,
				},
				resource.Attribute{
					Name:        "read_only_root_filesystem",
					Description: `(Optional) Whether this container has a read-only root filesystem. Default is false.`,
				},
				resource.Attribute{
					Name:        "run_as_non_root",
					Description: `(Optional) Indicates that the container must run as a non-root user. If true, the Kubelet will validate the image at runtime to ensure that it does not run as UID 0 (root) and fail to start the container if it does. If unset or false, no such validation will be performed. May also be set in PodSecurityContext. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.`,
				},
				resource.Attribute{
					Name:        "run_as_user",
					Description: `(Optional) The UID to run the entrypoint of the container process. Defaults to user specified in image metadata if unspecified. May also be set in PodSecurityContext. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.`,
				},
				resource.Attribute{
					Name:        "se_linux_options",
					Description: `(Optional) The SELinux context to be applied to the container. If unspecified, the container runtime will allocate a random SELinux context for each container. May also be set in PodSecurityContext. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence. ### ` + "`" + `capabilities` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "add",
					Description: `(Optional) A list of added capabilities.`,
				},
				resource.Attribute{
					Name:        "drop",
					Description: `(Optional) A list of removed capabilities. ### pod ` + "`" + `security_context` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "fs_group",
					Description: `(Optional) A special supplemental group that applies to all containers in a pod. Some volume types allow the Kubelet to change the ownership of that volume to be owned by the pod: 1. The owning GID will be the FSGroup 2. The setgid bit is set (new files created in the volume will be owned by FSGroup) 3. The permission bits are OR'd with rw-rw---- If unset, the Kubelet will not modify the ownership and permissions of any volume.`,
				},
				resource.Attribute{
					Name:        "run_as_non_root",
					Description: `(Optional) Indicates that the container must run as a non-root user. If true, the Kubelet will validate the image at runtime to ensure that it does not run as UID 0 (root) and fail to start the container if it does. If unset or false, no such validation will be performed. May also be set in SecurityContext. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.`,
				},
				resource.Attribute{
					Name:        "run_as_user",
					Description: `(Optional) The UID to run the entrypoint of the container process. Defaults to user specified in image metadata if unspecified. May also be set in SecurityContext. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence for that container.`,
				},
				resource.Attribute{
					Name:        "se_linux_options",
					Description: `(Optional) The SELinux context to be applied to all containers. If unspecified, the container runtime will allocate a random SELinux context for each container. May also be set in SecurityContext. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence for that container.`,
				},
				resource.Attribute{
					Name:        "supplemental_groups",
					Description: `(Optional) A list of groups applied to the first process run in each container, in addition to the container's primary GID. If unspecified, no groups will be added to any container.`,
				},
				resource.Attribute{
					Name:        "sysctl",
					Description: `(Optional) holds a list of namespaced sysctls used for the pod. see [Sysctl](#sysctl) block. See [official docs](https://kubernetes.io/docs/tasks/administer-cluster/sysctl-cluster/) for more details. ##### Sysctl`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of a property to set.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) Value of a property to set. ### ` + "`" + `tcp_socket` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Required) Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME. ### ` + "`" + `value_from` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "config_map_key_ref",
					Description: `(Optional) Selects a key of a ConfigMap.`,
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
					Description: `(Optional) Selects a key of a secret in the pod's namespace. ### ` + "`" + `toleration` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "effect",
					Description: `(Optional) Effect indicates the taint effect to match. Empty means match all taint effects. When specified, allowed values are NoSchedule, PreferNoSchedule and NoExecute.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Optional) Key is the taint key that the toleration applies to. Empty means match all taint keys. If the key is empty, operator must be Exists; this combination means to match all values and all keys.`,
				},
				resource.Attribute{
					Name:        "operator",
					Description: `(Optional) Operator represents a key's relationship to the value. Valid operators are Exists and Equal. Defaults to Equal. Exists is equivalent to wildcard for value, so that a pod can tolerate all taints of a particular category.`,
				},
				resource.Attribute{
					Name:        "toleration_seconds",
					Description: `(Optional) TolerationSeconds represents the period of time the toleration (which must be of effect NoExecute, otherwise this field is ignored) tolerates the taint. By default, it is not set, which means tolerate the taint forever (do not evict). Zero and negative values will be treated as 0 (evict immediately) by the system.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Optional) Value is the taint value the toleration matches to. If the operator is Exists, the value should be empty, otherwise just a regular string. ### ` + "`" + `projected` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "default_mode",
					Description: `(Optional) Mode bits used to set permissions on created files by default. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.`,
				},
				resource.Attribute{
					Name:        "sources",
					Description: `(Required) List of volume projection sources ### ` + "`" + `sources` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "config_map",
					Description: `(Optional) Adapts a ConfigMap into a projected volume. The contents of the target ConfigMap's Data field will be presented in a projected volume as files using the keys in the Data field as the file names, unless the items element is populated with specific mappings of keys to paths. Note that this is identical to a configmap volume source without the default mode.`,
				},
				resource.Attribute{
					Name:        "downward_api",
					Description: `(Optional) Represents downward API info for projecting into a projected volume. Note that this is identical to a downward_api volume source without the default mode.`,
				},
				resource.Attribute{
					Name:        "secret",
					Description: `(Optional) Adapts a secret into a projected volume. The contents of the target Secret's Data field will be presented in a projected volume as files using the keys in the Data field as the file names. Note that this is identical to a secret volume source without the default mode.`,
				},
				resource.Attribute{
					Name:        "service_account_token",
					Description: `(Optional) Represents a projected service account token volume. This projection can be used to insert a service account token into the pods runtime filesystem for use against APIs (Kubernetes API Server or otherwise). ### ` + "`" + `service_account_token` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "audience",
					Description: `(Optional) Audience is the intended audience of the token. A recipient of a token must identify itself with an identifier specified in the audience of the token, and otherwise should reject the token. The audience defaults to the identifier of the apiserver.`,
				},
				resource.Attribute{
					Name:        "expiration_seconds",
					Description: `(Optional) The requested duration of validity of the service account token. As the token approaches expiration, the kubelet volume plugin will proactively rotate the service account token. The kubelet will start trying to rotate the token if the token is older than 80 percent of its time to live or if the token is older than 24 hours.Defaults to 1 hour and must be at least 10 minutes.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Required) Path is the path relative to the mount point of the file to project the token into. ### ` + "`" + `volume` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "aws_elastic_block_store",
					Description: `(Optional) Represents an AWS Disk resource that is attached to a kubelet's host machine and then exposed to the pod. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/volumes#awselasticblockstore)`,
				},
				resource.Attribute{
					Name:        "azure_disk",
					Description: `(Optional) Represents an Azure Data Disk mount on the host and bind mount to the pod.`,
				},
				resource.Attribute{
					Name:        "azure_file",
					Description: `(Optional) Represents an Azure File Service mount on the host and bind mount to the pod.`,
				},
				resource.Attribute{
					Name:        "ceph_fs",
					Description: `(Optional) Represents a Ceph FS mount on the host that shares a pod's lifetime`,
				},
				resource.Attribute{
					Name:        "cinder",
					Description: `(Optional) Represents a cinder volume attached and mounted on kubelets host machine. For more info see https://github.com/kubernetes/examples/blob/master/mysql-cinder-pd/README.md#mysql-installation-with-cinder-volume-plugin.`,
				},
				resource.Attribute{
					Name:        "config_map",
					Description: `(Optional) ConfigMap represents a configMap that should populate this volume`,
				},
				resource.Attribute{
					Name:        "downward_api",
					Description: `(Optional) DownwardAPI represents downward API about the pod that should populate this volume`,
				},
				resource.Attribute{
					Name:        "empty_dir",
					Description: `(Optional) EmptyDir represents a temporary directory that shares a pod's lifetime. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/volumes#emptydir)`,
				},
				resource.Attribute{
					Name:        "fc",
					Description: `(Optional) Represents a Fibre Channel resource that is attached to a kubelet's host machine and then exposed to the pod.`,
				},
				resource.Attribute{
					Name:        "flex_volume",
					Description: `(Optional) Represents a generic volume resource that is provisioned/attached using an exec based plugin. This is an alpha feature and may change in future.`,
				},
				resource.Attribute{
					Name:        "flocker",
					Description: `(Optional) Represents a Flocker volume attached to a kubelet's host machine and exposed to the pod for its usage. This depends on the Flocker control service being running`,
				},
				resource.Attribute{
					Name:        "gce_persistent_disk",
					Description: `(Optional) Represents a GCE Disk resource that is attached to a kubelet's host machine and then exposed to the pod. Provisioned by an admin. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/volumes#gcepersistentdisk)`,
				},
				resource.Attribute{
					Name:        "git_repo",
					Description: `(Optional) GitRepo represents a git repository at a particular revision.`,
				},
				resource.Attribute{
					Name:        "glusterfs",
					Description: `(Optional) Represents a Glusterfs volume that is attached to a host and exposed to the pod. Provisioned by an admin. For more info see https://github.com/kubernetes/examples/tree/master/volumes/glusterfs#glusterfs.`,
				},
				resource.Attribute{
					Name:        "host_path",
					Description: `(Optional) Represents a directory on the host. Provisioned by a developer or tester. This is useful for single-node development and testing only! On-host storage is not supported in any way and WILL NOT WORK in a multi-node cluster. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/volumes#hostpath)`,
				},
				resource.Attribute{
					Name:        "iscsi",
					Description: `(Optional) Represents an ISCSI Disk resource that is attached to a kubelet's host machine and then exposed to the pod. Provisioned by an admin.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Volume's name. Must be a DNS_LABEL and unique within the pod. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/identifiers#names)`,
				},
				resource.Attribute{
					Name:        "nfs",
					Description: `(Optional) Represents an NFS mount on the host. Provisioned by an admin. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/volumes#nfs)`,
				},
				resource.Attribute{
					Name:        "persistent_volume_claim",
					Description: `(Optional) The specification of a persistent volume.`,
				},
				resource.Attribute{
					Name:        "photon_persistent_disk",
					Description: `(Optional) Represents a PhotonController persistent disk attached and mounted on kubelets host machine`,
				},
				resource.Attribute{
					Name:        "quobyte",
					Description: `(Optional) Quobyte represents a Quobyte mount on the host that shares a pod's lifetime`,
				},
				resource.Attribute{
					Name:        "rbd",
					Description: `(Optional) Represents a Rados Block Device mount on the host that shares a pod's lifetime. For more info see https://kubernetes.io/docs/concepts/storage/volumes/#rbd.`,
				},
				resource.Attribute{
					Name:        "secret",
					Description: `(Optional) Secret represents a secret that should populate this volume. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/volumes#secrets)`,
				},
				resource.Attribute{
					Name:        "vsphere_volume",
					Description: `(Optional) Represents a vSphere volume attached and mounted on kubelets host machine ### ` + "`" + `volume_mount` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "mount_path",
					Description: `(Required) Path within the container at which the volume should be mounted. Must not contain ':'.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) This must match the Name of a Volume.`,
				},
				resource.Attribute{
					Name:        "read_only",
					Description: `(Optional) Mounted read-only if true, read-write otherwise (false or unspecified). Defaults to false.`,
				},
				resource.Attribute{
					Name:        "sub_path",
					Description: `(Optional) Path within the volume from which the container's volume should be mounted. Defaults to "" (volume's root).`,
				},
				resource.Attribute{
					Name:        "mount_propagation",
					Description: `(Optional) Mount propagation mode. Defaults to "None". For more info see [Kubernetes reference](https://kubernetes.io/docs/concepts/storage/volumes/#mount-propagation) ### ` + "`" + `vsphere_volume` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "fs_type",
					Description: `(Optional) Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.`,
				},
				resource.Attribute{
					Name:        "volume_path",
					Description: `(Required) Path that identifies vSphere volume vmdk ## Timeouts The following [Timeout](/docs/configuration/resources.html#operation-timeouts) configuration options are available for the ` + "`" + `kubernetes_daemonset` + "`" + ` resource:`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Default ` + "`" + `10 minutes` + "`" + `) Used for creating new controller`,
				},
				resource.Attribute{
					Name:        "update",
					Description: `(Default ` + "`" + `10 minutes` + "`" + `) Used for updating a controller`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Default ` + "`" + `10 minutes` + "`" + `) Used for destroying a controller ## Import DaemonSet can be imported using the namespace and name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import kubernetes_daemonset.example default/terraform-example ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "kubernetes_deployment",
			Category:         "Resources",
			ShortDescription: `A Deployment ensures that a specified number of pod “replicas” are running at any one time. In other words, a Deployment makes sure that a pod or homogeneous set of pods are always up and available. If there are too many pods, it will kill some. If there are too few, the Deployment will start more.`,
			Description:      ``,
			Keywords: []string{
				"deployment",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "metadata",
					Description: `(Required) Standard deployment's metadata. For more info see [Kubernetes reference](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata)`,
				},
				resource.Attribute{
					Name:        "spec",
					Description: `(Required) Spec defines the specification of the desired behavior of the deployment. For more info see [Kubernetes reference](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#spec-and-status)`,
				},
				resource.Attribute{
					Name:        "wait_for_rollout",
					Description: `(Optional) Wait for the deployment to successfully roll out. Defaults to ` + "`" + `true` + "`" + `. ## Nested Blocks ### ` + "`" + `metadata` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "annotations",
					Description: `(Optional) An unstructured key value map stored with the deployment that may be used to store arbitrary metadata. ~> By default, the provider ignores any annotations whose key names end with`,
				},
				resource.Attribute{
					Name:        "generate_name",
					Description: `(Optional) Prefix, used by the server, to generate a unique name ONLY IF the ` + "`" + `name` + "`" + ` field has not been provided. This value will also be combined with a unique suffix. For more info see [Kubernetes reference](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#idempotency)`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) Map of string keys and values that can be used to organize and categorize (scope and select) the deployment. ~> By default, the provider ignores any labels whose key names end with`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the deployment, must be unique. Cannot be updated. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/identifiers#names)`,
				},
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) Namespace defines the space within which name of the deployment must be unique. #### Attributes`,
				},
				resource.Attribute{
					Name:        "generation",
					Description: `A sequence number representing a specific generation of the desired state.`,
				},
				resource.Attribute{
					Name:        "resource_version",
					Description: `An opaque value that represents the internal version of this deployment that can be used by clients to determine when deployment has changed. For more info see [Kubernetes reference](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency)`,
				},
				resource.Attribute{
					Name:        "uid",
					Description: `The unique in time and space value for this deployment. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/identifiers#uids) ### ` + "`" + `spec` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "min_ready_seconds",
					Description: `(Optional) Minimum number of seconds for which a newly created pod should be ready without any of its container crashing, for it to be considered available. Defaults to 0 (pod will be considered available as soon as it is ready)`,
				},
				resource.Attribute{
					Name:        "paused",
					Description: `(Optional) Indicates that the deployment is paused.`,
				},
				resource.Attribute{
					Name:        "progress_deadline_seconds",
					Description: `(Optional) The maximum time in seconds for a deployment to make progress before it is considered to be failed. The deployment controller will continue to process failed deployments and a condition with a ProgressDeadlineExceeded reason will be surfaced in the deployment status. Note that progress will not be estimated during the time a deployment is paused. Defaults to 600s.`,
				},
				resource.Attribute{
					Name:        "replicas",
					Description: `(Optional) The number of desired replicas. This attribute is a string to be able to distinguish between explicit zero and not specified. Defaults to 1. For more info see [Kubernetes reference](https://kubernetes.io/docs/concepts/workloads/controllers/deployment/#scaling-a-deployment)`,
				},
				resource.Attribute{
					Name:        "revision_history_limit",
					Description: `(Optional) The number of old ReplicaSets to retain to allow rollback. This is a pointer to distinguish between explicit zero and not specified. Defaults to 10.`,
				},
				resource.Attribute{
					Name:        "strategy",
					Description: `(Optional) The deployment strategy to use to replace existing pods with new ones.`,
				},
				resource.Attribute{
					Name:        "selector",
					Description: `(Optional) A label query over pods that should match the Replicas count. Label keys and values that must match in order to be controlled by this deployment.`,
				},
				resource.Attribute{
					Name:        "template",
					Description: `(Required) Describes the pod that will be created if insufficient replicas are detected. This takes precedence over a TemplateRef. For more info see [Kubernetes reference](https://kubernetes.io/docs/concepts/workloads/controllers/deployment/#pod-template) ### ` + "`" + `strategy` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type of deployment. Can be 'Recreate' or 'RollingUpdate'. Default is RollingUpdate.`,
				},
				resource.Attribute{
					Name:        "rolling_update",
					Description: `Rolling update config params. Present only if type = RollingUpdate. ### ` + "`" + `rolling_update` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "max_surge",
					Description: `The maximum number of pods that can be scheduled above the desired number of pods. Value can be an absolute number (ex: 5) or a percentage of desired pods (ex: 10%). This can not be 0 if MaxUnavailable is 0. Absolute number is calculated from percentage by rounding up. Defaults to 25%. Example: when this is set to 30%, the new RC can be scaled up immediately when the rolling update starts, such that the total number of old and new pods do not exceed 130% of desired pods. Once old pods have been killed, new RC can be scaled up further, ensuring that total number of pods running at any time during the update is atmost 130% of desired pods.`,
				},
				resource.Attribute{
					Name:        "max_unavailable",
					Description: `The maximum number of pods that can be unavailable during the update. Value can be an absolute number (ex: 5) or a percentage of desired pods (ex: 10%). Absolute number is calculated from percentage by rounding down. This can not be 0 if MaxSurge is 0. Defaults to 25%. Example: when this is set to 30%, the old RC can be scaled down to 70% of desired pods immediately when the rolling update starts. Once new pods are ready, old RC can be scaled down further, followed by scaling up the new RC, ensuring that the total number of pods available at all times during the update is at least 70% of desired pods. ### ` + "`" + `template` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `(Required) Standard object's metadata. For more info see https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata.`,
				},
				resource.Attribute{
					Name:        "spec",
					Description: `(Required) Specification of the desired behavior of the pod. For more info see https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#spec-and-status. ### template ` + "`" + `spec` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "affinity",
					Description: `(Optional) A group of affinity scheduling rules. If specified, the pod will be dispatched by specified scheduler. If not specified, the pod will be dispatched by default scheduler.`,
				},
				resource.Attribute{
					Name:        "active_deadline_seconds",
					Description: `(Optional) Optional duration in seconds the pod may be active on the node relative to StartTime before the system will actively try to mark it failed and kill associated containers. Value must be a positive integer.`,
				},
				resource.Attribute{
					Name:        "automount_service_account_token",
					Description: `(Optional) Indicates whether a service account token should be automatically mounted. Defaults to ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "container",
					Description: `(Optional) List of containers belonging to the pod. Containers cannot currently be added or removed. There must be at least one container in a Pod. Cannot be updated. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/containers)`,
				},
				resource.Attribute{
					Name:        "readiness_gate",
					Description: `(Optional) If specified, all readiness gates will be evaluated for pod readiness. A pod is ready when all its containers are ready AND all conditions specified in the readiness gates have status equal to "True". [More info](https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/#pod-readiness-gate)`,
				},
				resource.Attribute{
					Name:        "init_container",
					Description: `(Optional) List of init containers belonging to the pod. Init containers always run to completion and each must complete successfully before the next is started. For more info see [Kubernetes reference](https://kubernetes.io/docs/concepts/workloads/pods/init-containers)/`,
				},
				resource.Attribute{
					Name:        "dns_policy",
					Description: `(Optional) Set DNS policy for containers within the pod. Valid values are 'ClusterFirstWithHostNet', 'ClusterFirst', 'Default' or 'None'. DNS parameters given in DNSConfig will be merged with the policy selected with DNSPolicy. To have DNS options set along with hostNetwork, you have to specify DNS policy explicitly to 'ClusterFirstWithHostNet'. Optional: Defaults to 'ClusterFirst', see [Kubernetes reference](https://kubernetes.io/docs/concepts/services-networking/dns-pod-service/#pod-s-dns-policy).`,
				},
				resource.Attribute{
					Name:        "dns_config",
					Description: `(Optional) Specifies the DNS parameters of a pod. Parameters specified here will be merged to the generated DNS configuration based on DNSPolicy. Defaults to empty. See ` + "`" + `dns_config` + "`" + ` block definition below.`,
				},
				resource.Attribute{
					Name:        "enable_service_links",
					Description: `(Optional) Enables generating environment variables for service discovery. Optional: Defaults to true. For more info see [Kubernetes reference](https://kubernetes.io/docs/concepts/services-networking/connect-applications-service/#accessing-the-service).`,
				},
				resource.Attribute{
					Name:        "host_aliases",
					Description: `(Optional) List of hosts and IPs that will be injected into the pod's hosts file if specified. Optional: Defaults to empty. See ` + "`" + `host_aliases` + "`" + ` block definition below.`,
				},
				resource.Attribute{
					Name:        "host_ipc",
					Description: `(Optional) Use the host's ipc namespace. Optional: Defaults to false.`,
				},
				resource.Attribute{
					Name:        "host_network",
					Description: `(Optional) Host networking requested for this pod. Use the host's network namespace. If this option is set, the ports that will be used must be specified.`,
				},
				resource.Attribute{
					Name:        "host_pid",
					Description: `(Optional) Use the host's pid namespace.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `(Optional) Specifies the hostname of the Pod If not specified, the pod's hostname will be set to a system-defined value.`,
				},
				resource.Attribute{
					Name:        "image_pull_secrets",
					Description: `(Optional) ImagePullSecrets is an optional list of references to secrets in the same namespace to use for pulling any of the images used by this PodSpec. If specified, these secrets will be passed to individual puller implementations for them to use. For example, in the case of docker, only DockerConfig type secrets are honored. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/images#specifying-imagepullsecrets-on-a-pod)`,
				},
				resource.Attribute{
					Name:        "node_name",
					Description: `(Optional) NodeName is a request to schedule this pod onto a specific node. If it is non-empty, the scheduler simply schedules this pod onto that node, assuming that it fits resource requirements.`,
				},
				resource.Attribute{
					Name:        "node_selector",
					Description: `(Optional) NodeSelector is a selector which must be true for the pod to fit on a node. Selector which must match a node's labels for the pod to be scheduled on that node. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/node-selection).`,
				},
				resource.Attribute{
					Name:        "priority_class_name",
					Description: `(Optional) If specified, indicates the pod's priority. 'system-node-critical' and 'system-cluster-critical' are two special keywords which indicate the highest priorities with the formerer being the highest priority. Any other name must be defined by creating a PriorityClass object with that name. If not specified, the pod priority will be default or zero if there is no default.`,
				},
				resource.Attribute{
					Name:        "restart_policy",
					Description: `(Optional) Restart policy for all containers within the pod. One of Always, OnFailure, Never. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/pod-states#restartpolicy).`,
				},
				resource.Attribute{
					Name:        "security_context",
					Description: `(Optional) SecurityContext holds pod-level security attributes and common container settings. Optional: Defaults to empty`,
				},
				resource.Attribute{
					Name:        "service_account_name",
					Description: `(Optional) ServiceAccountName is the name of the ServiceAccount to use to run this pod. For more info see https://kubernetes.io/docs/tasks/configure-pod-container/configure-service-account/.`,
				},
				resource.Attribute{
					Name:        "share_process_namespace",
					Description: `(Optional) Share a single process namespace between all of the containers in a pod. When this is set containers will be able to view and signal processes from other containers in the same pod, and the first process in each container will not be assigned PID 1. HostPID and ShareProcessNamespace cannot both be set.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `(Optional) If specified, the fully qualified Pod hostname will be "...svc.". If not specified, the pod will not have a domainname at all..`,
				},
				resource.Attribute{
					Name:        "termination_grace_period_seconds",
					Description: `(Optional) Optional duration in seconds the pod needs to terminate gracefully. May be decreased in delete request. Value must be non-negative integer. The value zero indicates delete immediately. If this value is nil, the default grace period will be used instead. The grace period is the duration in seconds after the processes running in the pod are sent a termination signal and the time when the processes are forcibly halted with a kill signal. Set this value longer than the expected cleanup time for your process.`,
				},
				resource.Attribute{
					Name:        "toleration",
					Description: `(Optional) Optional pod node tolerations. For more info see [Kubernetes reference](https://kubernetes.io/docs/concepts/configuration/taint-and-toleration/)`,
				},
				resource.Attribute{
					Name:        "volume",
					Description: `(Optional) List of volumes that can be mounted by containers belonging to the pod. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/volumes) ### ` + "`" + `affinity` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "node_affinity",
					Description: `(Optional) Node affinity scheduling rules for the pod. For more info see [Kubernetes reference](https://kubernetes.io/docs/concepts/configuration/assign-pod-node/#node-affinity-beta-feature)`,
				},
				resource.Attribute{
					Name:        "pod_affinity",
					Description: `(Optional) Inter-pod topological affinity. rules that specify that certain pods should be placed in the same topological domain (e.g. same node, same rack, same zone, same power domain, etc.) For more info see [Kubernetes reference](https://kubernetes.io/docs/concepts/configuration/assign-pod-node/#inter-pod-affinity-and-anti-affinity-beta-feature)`,
				},
				resource.Attribute{
					Name:        "pod_anti_affinity",
					Description: `(Optional) Inter-pod topological affinity. rules that specify that certain pods should be placed in the same topological domain (e.g. same node, same rack, same zone, same power domain, etc.) For more info see [Kubernetes reference](https://kubernetes.io/docs/concepts/configuration/assign-pod-node/#inter-pod-affinity-and-anti-affinity-beta-feature) ### ` + "`" + `node_affinity` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "required_during_scheduling_ignored_during_execution",
					Description: `(Optional) If the affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to an update), the system may or may not try to eventually evict the pod from its node.`,
				},
				resource.Attribute{
					Name:        "preferred_during_scheduling_ignored_during_execution",
					Description: `(Optional) The scheduler will prefer to schedule pods to nodes that satisfy the affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. ### ` + "`" + `required_during_scheduling_ignored_during_execution` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "node_selector_term",
					Description: `(Required) A list of node selector terms. The terms are ORed. ## ` + "`" + `node_selector_term` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "match_expressions",
					Description: `(Optional) A list of node selector requirements by node's labels.`,
				},
				resource.Attribute{
					Name:        "match_fields",
					Description: `(Optional) A list of node selector requirements by node's fields. ### ` + "`" + `match_expressions` + "`" + ` / ` + "`" + `match_fields` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) The label key that the selector applies to.`,
				},
				resource.Attribute{
					Name:        "operator",
					Description: `(Required) Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Optional) An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. ### ` + "`" + `preferred_during_scheduling_ignored_during_execution` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "preference",
					Description: `(Required) A node selector term, associated with the corresponding weight.`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `(Required) Weight associated with matching the corresponding nodeSelectorTerm, in the range 1-100. ### ` + "`" + `preference` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "match_expressions",
					Description: `(Optional) A list of node selector requirements by node's labels.`,
				},
				resource.Attribute{
					Name:        "match_fields",
					Description: `(Optional) A list of node selector requirements by node's fields. ## ` + "`" + `match_expressions` + "`" + ` / ` + "`" + `match_fields` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) The label key that the selector applies to.`,
				},
				resource.Attribute{
					Name:        "operator",
					Description: `(Required) Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Optional) An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. ### ` + "`" + `pod_affinity` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "required_during_scheduling_ignored_during_execution",
					Description: `(Optional) If the affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to a pod label update), the system may or may not try to eventually evict the pod from its node.`,
				},
				resource.Attribute{
					Name:        "preferred_during_scheduling_ignored_during_execution",
					Description: `(Optional) The scheduler will prefer to schedule pods to nodes that satisfy the affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. ### ` + "`" + `pod_anti_affinity` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "required_during_scheduling_ignored_during_execution",
					Description: `(Optional) If the anti-affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the anti-affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to a pod label update), the system may or may not try to eventually evict the pod from its node.`,
				},
				resource.Attribute{
					Name:        "preferred_during_scheduling_ignored_during_execution",
					Description: `(Optional) The scheduler will prefer to schedule pods to nodes that satisfy the anti-affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. ### ` + "`" + `required_during_scheduling_ignored_during_execution` + "`" + ` (pod_affinity_term) #### Arguments`,
				},
				resource.Attribute{
					Name:        "label_selector",
					Description: `(Optional) A label query over a set of resources, in this case pods.`,
				},
				resource.Attribute{
					Name:        "namespaces",
					Description: `(Optional) Specifies which namespaces the ` + "`" + `label_selector` + "`" + ` applies to (matches against). Null or empty list means "this pod's namespace"`,
				},
				resource.Attribute{
					Name:        "topology_key",
					Description: `(Optional) This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the ` + "`" + `label_selector` + "`" + ` in the specified namespaces, where co-located is defined as running on a node whose value of the label with key ` + "`" + `topology_key` + "`" + ` matches that of any node on which any of the selected pods is running. Empty ` + "`" + `topology_key` + "`" + ` is not allowed. ### ` + "`" + `preferred_during_scheduling_ignored_during_execution` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "pod_affinity_term",
					Description: `(Required) A pod affinity term, associated with the corresponding weight.`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `(Required) Weight associated with matching the corresponding ` + "`" + `pod_affinity_term` + "`" + `, in the range 1-100. ### ` + "`" + `container` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "args",
					Description: `(Optional) Arguments to the entrypoint. The docker image's CMD is used if this is not provided. Variable references $(VAR_NAME) are expanded using the container's environment. If a variable cannot be resolved, the reference in the input string will be unchanged. The $(VAR_NAME) syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped references will never be expanded, regardless of whether the variable exists or not. Cannot be updated. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/containers#containers-and-commands)`,
				},
				resource.Attribute{
					Name:        "command",
					Description: `(Optional) Entrypoint array. Not executed within a shell. The docker image's ENTRYPOINT is used if this is not provided. Variable references $(VAR_NAME) are expanded using the container's environment. If a variable cannot be resolved, the reference in the input string will be unchanged. The $(VAR_NAME) syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped references will never be expanded, regardless of whether the variable exists or not. Cannot be updated. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/containers#containers-and-commands)`,
				},
				resource.Attribute{
					Name:        "env",
					Description: `(Optional) Block of string name and value pairs to set in the container's environment. May be declared multiple times. Cannot be updated.`,
				},
				resource.Attribute{
					Name:        "env_from",
					Description: `(Optional) List of sources to populate environment variables in the container. The keys defined within a source must be a C_IDENTIFIER. All invalid keys will be reported as an event when the container is starting. When a key exists in multiple sources, the value associated with the last source will take precedence. Values defined by an Env with a duplicate key will take precedence. Cannot be updated.`,
				},
				resource.Attribute{
					Name:        "image",
					Description: `(Optional) Docker image name. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/images)`,
				},
				resource.Attribute{
					Name:        "image_pull_policy",
					Description: `(Optional) Image pull policy. One of Always, Never, IfNotPresent. Defaults to Always if :latest tag is specified, or IfNotPresent otherwise. Cannot be updated. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/images#updating-images)`,
				},
				resource.Attribute{
					Name:        "lifecycle",
					Description: `(Optional) Actions that the management system should take in response to container lifecycle events`,
				},
				resource.Attribute{
					Name:        "liveness_probe",
					Description: `(Optional) Periodic probe of container liveness. Container will be restarted if the probe fails. Cannot be updated. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/pod-states#container-probes)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the container specified as a DNS_LABEL. Each container in a pod must have a unique name (DNS_LABEL). Cannot be updated.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional) List of ports to expose from the container. Exposing a port here gives the system additional information about the network connections a container uses, but is primarily informational. Not specifying a port here DOES NOT prevent that port from being exposed. Any port which is listening on the default "0.0.0.0" address inside a container will be accessible from the network. Cannot be updated.`,
				},
				resource.Attribute{
					Name:        "readiness_probe",
					Description: `(Optional) Periodic probe of container service readiness. Container will be removed from service endpoints if the probe fails. Cannot be updated. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/pod-states#container-probes)`,
				},
				resource.Attribute{
					Name:        "resources",
					Description: `(Optional) Compute Resources required by this container. Cannot be updated. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/persistent-volumes#resources)`,
				},
				resource.Attribute{
					Name:        "security_context",
					Description: `(Optional) Security options the pod should run with. For more info see https://kubernetes.io/docs/tasks/configure-pod-container/security-context/.`,
				},
				resource.Attribute{
					Name:        "startup_probe",
					Description: `(Optional) StartupProbe indicates that the Pod has successfully initialized. If specified, no other probes are executed until this completes successfully. If this probe fails, the Pod will be restarted, just as if the livenessProbe failed. This can be used to provide different probe parameters at the beginning of a Pod's lifecycle, when it might take a long time to load data or warm a cache, than during steady-state operation. This cannot be updated. For more info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes`,
				},
				resource.Attribute{
					Name:        "stdin",
					Description: `(Optional) Whether this container should allocate a buffer for stdin in the container runtime. If this is not set, reads from stdin in the container will always result in EOF.`,
				},
				resource.Attribute{
					Name:        "stdin_once",
					Description: `(Optional) Whether the container runtime should close the stdin channel after it has been opened by a single attach. When stdin is true the stdin stream will remain open across multiple attach sessions. If stdinOnce is set to true, stdin is opened on container start, is empty until the first client attaches to stdin, and then remains open and accepts data until the client disconnects, at which time stdin is closed and remains closed until the container is restarted. If this flag is false, a container processes that reads from stdin will never receive an EOF.`,
				},
				resource.Attribute{
					Name:        "termination_message_path",
					Description: `(Optional) Optional: Path at which the file to which the container's termination message will be written is mounted into the container's filesystem. Message written is intended to be brief final status, such as an assertion failure message. Defaults to /dev/termination-log. Cannot be updated.`,
				},
				resource.Attribute{
					Name:        "tty",
					Description: `(Optional) Whether this container should allocate a TTY for itself`,
				},
				resource.Attribute{
					Name:        "volume_mount",
					Description: `(Optional) Pod volumes to mount into the container's filesystem. Cannot be updated.`,
				},
				resource.Attribute{
					Name:        "working_dir",
					Description: `(Optional) Container's working directory. If not specified, the container runtime's default will be used, which might be configured in the container image. Cannot be updated. ### ` + "`" + `readiness_gate` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "condition_type",
					Description: `(Required) refers to a condition in the pod's condition list with matching type. ### ` + "`" + `aws_elastic_block_store` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "fs_type",
					Description: `(Optional) Filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/volumes#awselasticblockstore)`,
				},
				resource.Attribute{
					Name:        "partition",
					Description: `(Optional) The partition in the volume that you want to mount. If omitted, the default is to mount by volume name. Examples: For volume /dev/sda1, you specify the partition as "1". Similarly, the volume partition for /dev/sda is "0" (or you can leave the property empty).`,
				},
				resource.Attribute{
					Name:        "read_only",
					Description: `(Optional) Whether to set the read-only property in VolumeMounts to "true". If omitted, the default is "false". For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/volumes#awselasticblockstore)`,
				},
				resource.Attribute{
					Name:        "volume_id",
					Description: `(Required) Unique ID of the persistent disk resource in AWS (Amazon EBS volume). For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/volumes#awselasticblockstore) ### ` + "`" + `azure_disk` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "caching_mode",
					Description: `(Required) Host Caching mode: None, Read Only, Read Write.`,
				},
				resource.Attribute{
					Name:        "data_disk_uri",
					Description: `(Required) The URI the data disk in the blob storage`,
				},
				resource.Attribute{
					Name:        "disk_name",
					Description: `(Required) The Name of the data disk in the blob storage`,
				},
				resource.Attribute{
					Name:        "fs_type",
					Description: `(Optional) Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.`,
				},
				resource.Attribute{
					Name:        "read_only",
					Description: `(Optional) Whether to force the read-only setting in VolumeMounts. Defaults to false (read/write). ### ` + "`" + `azure_file` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "read_only",
					Description: `(Optional) Whether to force the read-only setting in VolumeMounts. Defaults to false (read/write).`,
				},
				resource.Attribute{
					Name:        "secret_name",
					Description: `(Required) The name of secret that contains Azure Storage Account Name and Key`,
				},
				resource.Attribute{
					Name:        "share_name",
					Description: `(Required) Share Name ### ` + "`" + `capabilities` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "add",
					Description: `(Optional) Added capabilities`,
				},
				resource.Attribute{
					Name:        "drop",
					Description: `(Optional) Removed capabilities ### ` + "`" + `ceph_fs` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "monitors",
					Description: `(Required) Monitors is a collection of Ceph monitors. For more info see https://github.com/kubernetes/examples/tree/master/volumes/cephfs/#how-to-use-it.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Optional) Used as the mounted root, rather than the full Ceph tree, default is /.`,
				},
				resource.Attribute{
					Name:        "read_only",
					Description: `(Optional) Whether to force the read-only setting in VolumeMounts. Defaults to ` + "`" + `false` + "`" + ` (read/write). For more info see https://github.com/kubernetes/examples/tree/master/volumes/cephfs/#how-to-use-it.`,
				},
				resource.Attribute{
					Name:        "secret_file",
					Description: `(Optional) The path to key ring for User, default is /etc/ceph/user.secret. For more info see https://github.com/kubernetes/examples/tree/master/volumes/cephfs/#how-to-use-it.`,
				},
				resource.Attribute{
					Name:        "secret_ref",
					Description: `(Optional) Reference to the authentication secret for User, default is empty. For more info see https://github.com/kubernetes/examples/tree/master/volumes/cephfs/#how-to-use-it.`,
				},
				resource.Attribute{
					Name:        "user",
					Description: `(Optional) User is the rados user name, default is admin. For more info see https://github.com/kubernetes/examples/tree/master/volumes/cephfs/#how-to-use-it. ### ` + "`" + `cinder` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "fs_type",
					Description: `(Optional) Filesystem type to mount. Must be a filesystem type supported by the host operating system. Examples: "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified. For more info see https://github.com/kubernetes/examples/blob/master/mysql-cinder-pd/README.md#mysql-installation-with-cinder-volume-plugin.`,
				},
				resource.Attribute{
					Name:        "read_only",
					Description: `(Optional) Whether to force the read-only setting in VolumeMounts. Defaults to false (read/write). For more info see https://github.com/kubernetes/examples/blob/master/mysql-cinder-pd/README.md#mysql-installation-with-cinder-volume-plugin.`,
				},
				resource.Attribute{
					Name:        "volume_id",
					Description: `(Required) Volume ID used to identify the volume in Cinder. For more info see https://github.com/kubernetes/examples/blob/master/mysql-cinder-pd/README.md#mysql-installation-with-cinder-volume-plugin. ### ` + "`" + `config_map` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "default_mode",
					Description: `(Optional) Optional: mode bits to use on created files by default. Must be a value between 0 and 0777. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.`,
				},
				resource.Attribute{
					Name:        "items",
					Description: `(Optional) If unspecified, each key-value pair in the Data field of the referenced ConfigMap will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the ConfigMap, the volume setup will error unless it is marked ` + "`" + `optional` + "`" + `. Paths must be relative and may not contain the '..' path or start with '..'.`,
				},
				resource.Attribute{
					Name:        "optional",
					Description: `(Optional) Specify whether the ConfigMap or its keys must be defined.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the referent. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/identifiers#names) ### ` + "`" + `config_map_ref` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the referent. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/identifiers#names)`,
				},
				resource.Attribute{
					Name:        "optional",
					Description: `(Optional) Specify whether the ConfigMap must be defined ### ` + "`" + `config_map_key_ref` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Optional) The key to select.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the referent. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/identifiers#names)`,
				},
				resource.Attribute{
					Name:        "optional",
					Description: `(Optional) Specify whether the ConfigMap or its key must be defined ### ` + "`" + `dns_config` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "nameservers",
					Description: `(Optional) A list of DNS name server IP addresses specified as strings. This will be appended to the base nameservers generated from DNSPolicy. Duplicated nameservers will be removed. Optional: Defaults to empty.`,
				},
				resource.Attribute{
					Name:        "option",
					Description: `(Optional) A list of DNS resolver options specified as blocks with ` + "`" + `name` + "`" + `/` + "`" + `value` + "`" + ` pairs. This will be merged with the base options generated from DNSPolicy. Duplicated entries will be removed. Resolution options given in Options will override those that appear in the base DNSPolicy. Optional: Defaults to empty.`,
				},
				resource.Attribute{
					Name:        "searches",
					Description: `(Optional) A list of DNS search domains for host-name lookup specified as strings. This will be appended to the base search paths generated from DNSPolicy. Duplicated search paths will be removed. Optional: Defaults to empty. The ` + "`" + `option` + "`" + ` block supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the option.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Optional) Value of the option. Optional: Defaults to empty. ### ` + "`" + `downward_api` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "default_mode",
					Description: `(Optional) Optional: mode bits to use on created files by default. Must be a value between 0 and 0777. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.`,
				},
				resource.Attribute{
					Name:        "items",
					Description: `(Optional) If unspecified, each key-value pair in the Data field of the referenced ConfigMap will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the ConfigMap, the volume setup will error. Paths must be relative and may not contain the '..' path or start with '..'. ### ` + "`" + `empty_dir` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "medium",
					Description: `(Optional) What type of storage medium should back this directory. The default is "" which means to use the node's default medium. Must be an empty string (default) or Memory. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/volumes#emptydir)`,
				},
				resource.Attribute{
					Name:        "size_limit",
					Description: `(Optional) Total amount of local storage required for this EmptyDir volume. For more info see [Kubernetes reference](https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/#resource-units-in-kubernetes) and [Kubernetes Quantity type](https://pkg.go.dev/k8s.io/apimachinery/pkg/api/resource?tab=doc#Quantity). ### ` + "`" + `env` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the environment variable. Must be a C_IDENTIFIER`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Optional) Variable references $(VAR_NAME) are expanded using the previous defined environment variables in the container and any service environment variables. If a variable cannot be resolved, the reference in the input string will be unchanged. The $(VAR_NAME) syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped references will never be expanded, regardless of whether the variable exists or not. Defaults to "".`,
				},
				resource.Attribute{
					Name:        "value_from",
					Description: `(Optional) Source for the environment variable's value ### ` + "`" + `env_from` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "config_map_ref",
					Description: `(Optional) The ConfigMap to select from`,
				},
				resource.Attribute{
					Name:        "prefix",
					Description: `(Optional) An optional identifer to prepend to each key in the ConfigMap. Must be a C_IDENTIFIER..`,
				},
				resource.Attribute{
					Name:        "secret_ref",
					Description: `(Optional) The Secret to select from ### ` + "`" + `exec` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "command",
					Description: `(Optional) Command is the command line to execute inside the container, the working directory for the command is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy. ### ` + "`" + `fc` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "fs_type",
					Description: `(Optional) Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.`,
				},
				resource.Attribute{
					Name:        "lun",
					Description: `(Required) FC target lun number`,
				},
				resource.Attribute{
					Name:        "read_only",
					Description: `(Optional) Whether to force the read-only setting in VolumeMounts. Defaults to false (read/write).`,
				},
				resource.Attribute{
					Name:        "target_ww_ns",
					Description: `(Required) FC target worldwide names (WWNs) ### ` + "`" + `field_ref` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "api_version",
					Description: `(Optional) Version of the schema the FieldPath is written in terms of, defaults to "v1".`,
				},
				resource.Attribute{
					Name:        "field_path",
					Description: `(Optional) Path of the field to select in the specified API version ### ` + "`" + `flex_volume` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "driver",
					Description: `(Required) Driver is the name of the driver to use for this volume.`,
				},
				resource.Attribute{
					Name:        "fs_type",
					Description: `(Optional) Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs", "ntfs". The default filesystem depends on FlexVolume script.`,
				},
				resource.Attribute{
					Name:        "options",
					Description: `(Optional) Extra command options if any.`,
				},
				resource.Attribute{
					Name:        "read_only",
					Description: `(Optional) Whether to force the ReadOnly setting in VolumeMounts. Defaults to false (read/write).`,
				},
				resource.Attribute{
					Name:        "secret_ref",
					Description: `(Optional) Reference to the secret object containing sensitive information to pass to the plugin scripts. This may be empty if no secret object is specified. If the secret object contains more than one secret, all secrets are passed to the plugin scripts. ### ` + "`" + `flocker` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "dataset_name",
					Description: `(Optional) Name of the dataset stored as metadata -> name on the dataset for Flocker should be considered as deprecated`,
				},
				resource.Attribute{
					Name:        "dataset_uuid",
					Description: `(Optional) UUID of the dataset. This is unique identifier of a Flocker dataset ### ` + "`" + `gce_persistent_disk` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "fs_type",
					Description: `(Optional) Filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/volumes#gcepersistentdisk)`,
				},
				resource.Attribute{
					Name:        "partition",
					Description: `(Optional) The partition in the volume that you want to mount. If omitted, the default is to mount by volume name. Examples: For volume /dev/sda1, you specify the partition as "1". Similarly, the volume partition for /dev/sda is "0" (or you can leave the property empty). For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/volumes#gcepersistentdisk)`,
				},
				resource.Attribute{
					Name:        "pd_name",
					Description: `(Required) Unique name of the PD resource in GCE. Used to identify the disk in GCE. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/volumes#gcepersistentdisk)`,
				},
				resource.Attribute{
					Name:        "read_only",
					Description: `(Optional) Whether to force the ReadOnly setting in VolumeMounts. Defaults to false. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/volumes#gcepersistentdisk) ### ` + "`" + `git_repo` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "directory",
					Description: `(Optional) Target directory name. Must not contain or start with '..'. If '.' is supplied, the volume directory will be the git repository. Otherwise, if specified, the volume will contain the git repository in the subdirectory with the given name.`,
				},
				resource.Attribute{
					Name:        "repository",
					Description: `(Optional) Repository URL`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `(Optional) Commit hash for the specified revision. ### ` + "`" + `glusterfs` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "endpoints_name",
					Description: `(Required) The endpoint name that details Glusterfs topology. For more info see https://github.com/kubernetes/examples/tree/master/volumes/glusterfs#create-a-pod.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Required) The Glusterfs volume path. For more info see https://github.com/kubernetes/examples/tree/master/volumes/glusterfs#create-a-pod.`,
				},
				resource.Attribute{
					Name:        "read_only",
					Description: `(Optional) Whether to force the Glusterfs volume to be mounted with read-only permissions. Defaults to false. For more info see https://github.com/kubernetes/examples/tree/master/volumes/glusterfs#create-a-pod. ### ` + "`" + `host_aliases` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "hostnames",
					Description: `(Required) Array of hostnames for the IP address.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `(Required) IP address of the host file entry. ### ` + "`" + `host_path` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Optional) Path of the directory on the host. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/volumes#hostpath)`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) Type for HostPath volume. Defaults to "". For more info see [Kubernetes reference](https://kubernetes.io/docs/concepts/storage/volumes#hostpath) ### ` + "`" + `http_get` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `(Optional) Host name to connect to, defaults to the pod IP. You probably want to set "Host" in httpHeaders instead.`,
				},
				resource.Attribute{
					Name:        "http_header",
					Description: `(Optional) Scheme to use for connecting to the host.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Optional) Path to access on the HTTP server.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional) Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.`,
				},
				resource.Attribute{
					Name:        "scheme",
					Description: `(Optional) Scheme to use for connecting to the host. ### ` + "`" + `http_header` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The header field name`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Optional) The header field value ### ` + "`" + `image_pull_secrets` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the referent. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/identifiers#names) ### ` + "`" + `iscsi` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "fs_type",
					Description: `(Optional) Filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/volumes#iscsi)`,
				},
				resource.Attribute{
					Name:        "iqn",
					Description: `(Required) Target iSCSI Qualified Name.`,
				},
				resource.Attribute{
					Name:        "iscsi_interface",
					Description: `(Optional) iSCSI interface name that uses an iSCSI transport. Defaults to 'default' (tcp).`,
				},
				resource.Attribute{
					Name:        "lun",
					Description: `(Optional) iSCSI target lun number.`,
				},
				resource.Attribute{
					Name:        "read_only",
					Description: `(Optional) Whether to force the read-only setting in VolumeMounts. Defaults to false.`,
				},
				resource.Attribute{
					Name:        "target_portal",
					Description: `(Required) iSCSI target portal. The portal is either an IP or ip_addr:port if the port is other than default (typically TCP ports 860 and 3260). ### ` + "`" + `items` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Optional) The key to project.`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `(Optional) Optional: mode bits to use on this file, must be a value between 0 and 0777. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Optional) The relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'. ### ` + "`" + `lifecycle` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "post_start",
					Description: `(Optional) post_start is called immediately after a container is created. If the handler fails, the container is terminated and restarted according to its restart policy. Other management of the container blocks until the hook completes. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/container-environment#hook-details)`,
				},
				resource.Attribute{
					Name:        "pre_stop",
					Description: `(Optional) pre_stop is called immediately before a container is terminated. The container is terminated after the handler completes. The reason for termination is passed to the handler. Regardless of the outcome of the handler, the container is eventually terminated. Other management of the container blocks until the hook completes. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/container-environment#hook-details) ### ` + "`" + `liveness_probe` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "exec",
					Description: `(Optional) exec specifies the action to take.`,
				},
				resource.Attribute{
					Name:        "failure_threshold",
					Description: `(Optional) Minimum consecutive failures for the probe to be considered failed after having succeeded.`,
				},
				resource.Attribute{
					Name:        "http_get",
					Description: `(Optional) Specifies the http request to perform.`,
				},
				resource.Attribute{
					Name:        "initial_delay_seconds",
					Description: `(Optional) Number of seconds after the container has started before liveness probes are initiated. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/pod-states#container-probes)`,
				},
				resource.Attribute{
					Name:        "period_seconds",
					Description: `(Optional) How often (in seconds) to perform the probe`,
				},
				resource.Attribute{
					Name:        "success_threshold",
					Description: `(Optional) Minimum consecutive successes for the probe to be considered successful after having failed.`,
				},
				resource.Attribute{
					Name:        "tcp_socket",
					Description: `(Optional) TCPSocket specifies an action involving a TCP port. TCP hooks not yet supported`,
				},
				resource.Attribute{
					Name:        "timeout_seconds",
					Description: `(Optional) Number of seconds after which the probe times out. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/pod-states#container-probes) ### ` + "`" + `nfs` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Required) Path that is exported by the NFS server. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/volumes#nfs)`,
				},
				resource.Attribute{
					Name:        "read_only",
					Description: `(Optional) Whether to force the NFS export to be mounted with read-only permissions. Defaults to false. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/volumes#nfs)`,
				},
				resource.Attribute{
					Name:        "server",
					Description: `(Required) Server is the hostname or IP address of the NFS server. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/volumes#nfs) ### ` + "`" + `persistent_volume_claim` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "claim_name",
					Description: `(Optional) ClaimName is the name of a PersistentVolumeClaim in the same`,
				},
				resource.Attribute{
					Name:        "read_only",
					Description: `(Optional) Will force the ReadOnly setting in VolumeMounts. ### ` + "`" + `photon_persistent_disk` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "fs_type",
					Description: `(Optional) Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.`,
				},
				resource.Attribute{
					Name:        "pd_id",
					Description: `(Required) ID that identifies Photon Controller persistent disk ### ` + "`" + `port` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "container_port",
					Description: `(Required) Number of port to expose on the pod's IP address. This must be a valid port number, 0 < x < 65536.`,
				},
				resource.Attribute{
					Name:        "host_ip",
					Description: `(Optional) What host IP to bind the external port to.`,
				},
				resource.Attribute{
					Name:        "host_port",
					Description: `(Optional) Number of port to expose on the host. If specified, this must be a valid port number, 0 < x < 65536. If HostNetwork is specified, this must match ContainerPort. Most containers do not need this.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) If specified, this must be an IANA_SVC_NAME and unique within the pod. Each named port in a pod must have a unique name. Name for the port that can be referred to by services`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Optional) Protocol for port. Must be UDP or TCP. Defaults to "TCP". ### ` + "`" + `post_start` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "exec",
					Description: `(Optional) exec specifies the action to take.`,
				},
				resource.Attribute{
					Name:        "http_get",
					Description: `(Optional) Specifies the http request to perform.`,
				},
				resource.Attribute{
					Name:        "tcp_socket",
					Description: `(Optional) TCPSocket specifies an action involving a TCP port. TCP hooks not yet supported ### ` + "`" + `pre_stop` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "exec",
					Description: `(Optional) exec specifies the action to take.`,
				},
				resource.Attribute{
					Name:        "http_get",
					Description: `(Optional) Specifies the http request to perform.`,
				},
				resource.Attribute{
					Name:        "tcp_socket",
					Description: `(Optional) TCPSocket specifies an action involving a TCP port. TCP hooks not yet supported ### ` + "`" + `quobyte` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "group",
					Description: `(Optional) Group to map volume access to Default is no group`,
				},
				resource.Attribute{
					Name:        "read_only",
					Description: `(Optional) Whether to force the Quobyte volume to be mounted with read-only permissions. Defaults to false.`,
				},
				resource.Attribute{
					Name:        "registry",
					Description: `(Required) Registry represents a single or multiple Quobyte Registry services specified as a string as host:port pair (multiple entries are separated with commas) which acts as the central registry for volumes`,
				},
				resource.Attribute{
					Name:        "user",
					Description: `(Optional) User to map volume access to Defaults to serivceaccount user`,
				},
				resource.Attribute{
					Name:        "volume",
					Description: `(Required) Volume is a string that references an already created Quobyte volume by name. ### ` + "`" + `rbd` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "ceph_monitors",
					Description: `(Required) A collection of Ceph monitors. For more info see https://github.com/kubernetes/examples/tree/master/volumes/rbd#how-to-use-it.`,
				},
				resource.Attribute{
					Name:        "fs_type",
					Description: `(Optional) Filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/volumes#rbd)`,
				},
				resource.Attribute{
					Name:        "keyring",
					Description: `(Optional) Keyring is the path to key ring for RBDUser. Default is /etc/ceph/keyring. For more info see https://github.com/kubernetes/examples/tree/master/volumes/rbd#how-to-use-it.`,
				},
				resource.Attribute{
					Name:        "rados_user",
					Description: `(Optional) The rados user name. Default is admin. For more info see https://github.com/kubernetes/examples/tree/master/volumes/rbd#how-to-use-it.`,
				},
				resource.Attribute{
					Name:        "rbd_image",
					Description: `(Required) The rados image name. For more info see https://github.com/kubernetes/examples/tree/master/volumes/rbd#how-to-use-it.`,
				},
				resource.Attribute{
					Name:        "rbd_pool",
					Description: `(Optional) The rados pool name. Default is rbd. For more info see https://github.com/kubernetes/examples/tree/master/volumes/rbd#how-to-use-it.`,
				},
				resource.Attribute{
					Name:        "read_only",
					Description: `(Optional) Whether to force the read-only setting in VolumeMounts. Defaults to false. For more info see https://github.com/kubernetes/examples/tree/master/volumes/rbd#how-to-use-it.`,
				},
				resource.Attribute{
					Name:        "secret_ref",
					Description: `(Optional) Name of the authentication secret for RBDUser. If provided overrides keyring. Default is nil. For more info see https://github.com/kubernetes/examples/tree/master/volumes/rbd#how-to-use-it. ### ` + "`" + `readiness_probe` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "exec",
					Description: `(Optional) exec specifies the action to take.`,
				},
				resource.Attribute{
					Name:        "failure_threshold",
					Description: `(Optional) Minimum consecutive failures for the probe to be considered failed after having succeeded.`,
				},
				resource.Attribute{
					Name:        "http_get",
					Description: `(Optional) Specifies the http request to perform.`,
				},
				resource.Attribute{
					Name:        "initial_delay_seconds",
					Description: `(Optional) Number of seconds after the container has started before liveness probes are initiated. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/pod-states#container-probes)`,
				},
				resource.Attribute{
					Name:        "period_seconds",
					Description: `(Optional) How often (in seconds) to perform the probe`,
				},
				resource.Attribute{
					Name:        "success_threshold",
					Description: `(Optional) Minimum consecutive successes for the probe to be considered successful after having failed.`,
				},
				resource.Attribute{
					Name:        "tcp_socket",
					Description: `(Optional) TCPSocket specifies an action involving a TCP port. TCP hooks not yet supported`,
				},
				resource.Attribute{
					Name:        "timeout_seconds",
					Description: `(Optional) Number of seconds after which the probe times out. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/pod-states#container-probes) ### ` + "`" + `resources` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "limits",
					Description: `(Optional) Describes the maximum amount of compute resources allowed. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/compute-resources)/`,
				},
				resource.Attribute{
					Name:        "requests",
					Description: `(Optional) Describes the minimum amount of compute resources required. ### ` + "`" + `resource_field_ref` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "container_name",
					Description: `(Optional) The name of the container`,
				},
				resource.Attribute{
					Name:        "resource",
					Description: `(Required) Resource to select`,
				},
				resource.Attribute{
					Name:        "divisor",
					Description: `(Optional) Specifies the output format of the exposed resources, defaults to "1". ### ` + "`" + `se_linux_options` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "level",
					Description: `(Optional) Level is SELinux level label that applies to the container.`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `(Optional) Role is a SELinux role label that applies to the container.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) Type is a SELinux type label that applies to the container.`,
				},
				resource.Attribute{
					Name:        "user",
					Description: `(Optional) User is a SELinux user label that applies to the container. ### ` + "`" + `secret` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "default_mode",
					Description: `(Optional) Mode bits to use on created files by default. Must be a value between 0 and 0777. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.`,
				},
				resource.Attribute{
					Name:        "items",
					Description: `(Optional) List of Secret Items to project into the volume. See ` + "`" + `items` + "`" + ` block definition below. If unspecified, each key-value pair in the Data field of the referenced Secret will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the Secret, the volume setup will error unless it is marked ` + "`" + `optional` + "`" + `. Paths must be relative and may not contain the '..' path or start with '..'.`,
				},
				resource.Attribute{
					Name:        "optional",
					Description: `(Optional) Specify whether the Secret or its keys must be defined.`,
				},
				resource.Attribute{
					Name:        "secret_name",
					Description: `(Optional) Name of the secret in the pod's namespace to use. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/volumes#secrets) The ` + "`" + `items` + "`" + ` block supports the following:`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) The key to project.`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `(Optional) Mode bits to use on this file, must be a value between 0 and 0777. If not specified, the volume defaultMode will be used.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Required) The relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'. ### ` + "`" + `secret_ref` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the referent. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/identifiers#names)`,
				},
				resource.Attribute{
					Name:        "optional",
					Description: `(Optional) Specify whether the Secret must be defined ### ` + "`" + `secret_key_ref` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Optional) The key of the secret to select from. Must be a valid secret key.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the referent. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/identifiers#names)`,
				},
				resource.Attribute{
					Name:        "optional",
					Description: `(Optional) Specify whether the Secret or its key must be defined ### ` + "`" + `secret_ref` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the referent. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/identifiers#names) ### container ` + "`" + `security_context` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "allow_privilege_escalation",
					Description: `(Optional) AllowPrivilegeEscalation controls whether a process can gain more privileges than its parent process. This bool directly controls if the no_new_privs flag will be set on the container process. AllowPrivilegeEscalation is true always when the container is: 1) run as Privileged 2) has CAP_SYS_ADMIN`,
				},
				resource.Attribute{
					Name:        "capabilities",
					Description: `(Optional) The capabilities to add/drop when running containers. Defaults to the default set of capabilities granted by the container runtime.`,
				},
				resource.Attribute{
					Name:        "privileged",
					Description: `(Optional) Run container in privileged mode. Processes in privileged containers are essentially equivalent to root on the host. Defaults to false.`,
				},
				resource.Attribute{
					Name:        "read_only_root_filesystem",
					Description: `(Optional) Whether this container has a read-only root filesystem. Default is false.`,
				},
				resource.Attribute{
					Name:        "run_as_group",
					Description: `(Optional) The GID to run the entrypoint of the container process. Uses runtime default if unset. May also be set in PodSecurityContext. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.`,
				},
				resource.Attribute{
					Name:        "run_as_non_root",
					Description: `(Optional) Indicates that the container must run as a non-root user. If true, the Kubelet will validate the image at runtime to ensure that it does not run as UID 0 (root) and fail to start the container if it does. If unset or false, no such validation will be performed. May also be set in PodSecurityContext. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.`,
				},
				resource.Attribute{
					Name:        "run_as_user",
					Description: `(Optional) The UID to run the entrypoint of the container process. Defaults to user specified in image metadata if unspecified. May also be set in PodSecurityContext. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.`,
				},
				resource.Attribute{
					Name:        "se_linux_options",
					Description: `(Optional) The SELinux context to be applied to the container. If unspecified, the container runtime will allocate a random SELinux context for each container. May also be set in PodSecurityContext. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence. ### ` + "`" + `capabilities` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "add",
					Description: `(Optional) A list of added capabilities.`,
				},
				resource.Attribute{
					Name:        "drop",
					Description: `(Optional) A list of removed capabilities. ### pod ` + "`" + `security_context` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "fs_group",
					Description: `(Optional) A special supplemental group that applies to all containers in a pod. Some volume types allow the Kubelet to change the ownership of that volume to be owned by the pod: 1. The owning GID will be the FSGroup 2. The setgid bit is set (new files created in the volume will be owned by FSGroup) 3. The permission bits are OR'd with rw-rw---- If unset, the Kubelet will not modify the ownership and permissions of any volume.`,
				},
				resource.Attribute{
					Name:        "run_as_group",
					Description: `(Optional) The GID to run the entrypoint of the container process. Uses runtime default if unset. May also be set in SecurityContext. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence for that container.`,
				},
				resource.Attribute{
					Name:        "run_as_non_root",
					Description: `(Optional) Indicates that the container must run as a non-root user. If true, the Kubelet will validate the image at runtime to ensure that it does not run as UID 0 (root) and fail to start the container if it does. If unset or false, no such validation will be performed. May also be set in SecurityContext. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.`,
				},
				resource.Attribute{
					Name:        "run_as_user",
					Description: `(Optional) The UID to run the entrypoint of the container process. Defaults to user specified in image metadata if unspecified. May also be set in SecurityContext. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence for that container.`,
				},
				resource.Attribute{
					Name:        "se_linux_options",
					Description: `(Optional) The SELinux context to be applied to all containers. If unspecified, the container runtime will allocate a random SELinux context for each container. May also be set in SecurityContext. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence for that container.`,
				},
				resource.Attribute{
					Name:        "supplemental_groups",
					Description: `(Optional) A list of groups applied to the first process run in each container, in addition to the container's primary GID. If unspecified, no groups will be added to any container.`,
				},
				resource.Attribute{
					Name:        "sysctl",
					Description: `(Optional) holds a list of namespaced sysctls used for the pod. see [Sysctl](#sysctl) block. See [official docs](https://kubernetes.io/docs/tasks/administer-cluster/sysctl-cluster/) for more details. ##### Sysctl`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of a property to set.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) Value of a property to set. ### ` + "`" + `tcp_socket` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Required) Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME. ### ` + "`" + `value_from` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "config_map_key_ref",
					Description: `(Optional) Selects a key of a ConfigMap.`,
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
					Description: `(Optional) Selects a key of a secret in the pod's namespace. ### ` + "`" + `toleration` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "effect",
					Description: `(Optional) Effect indicates the taint effect to match. Empty means match all taint effects. When specified, allowed values are NoSchedule, PreferNoSchedule and NoExecute.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Optional) Key is the taint key that the toleration applies to. Empty means match all taint keys. If the key is empty, operator must be Exists; this combination means to match all values and all keys.`,
				},
				resource.Attribute{
					Name:        "operator",
					Description: `(Optional) Operator represents a key's relationship to the value. Valid operators are Exists and Equal. Defaults to Equal. Exists is equivalent to wildcard for value, so that a pod can tolerate all taints of a particular category.`,
				},
				resource.Attribute{
					Name:        "toleration_seconds",
					Description: `(Optional) TolerationSeconds represents the period of time the toleration (which must be of effect NoExecute, otherwise this field is ignored) tolerates the taint. By default, it is not set, which means tolerate the taint forever (do not evict). Zero and negative values will be treated as 0 (evict immediately) by the system.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Optional) Value is the taint value the toleration matches to. If the operator is Exists, the value should be empty, otherwise just a regular string. ### ` + "`" + `projected` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "default_mode",
					Description: `(Optional) Mode bits used to set permissions on created files by default. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.`,
				},
				resource.Attribute{
					Name:        "sources",
					Description: `(Required) List of volume projection sources ### ` + "`" + `sources` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "config_map",
					Description: `(Optional) Adapts a ConfigMap into a projected volume. The contents of the target ConfigMap's Data field will be presented in a projected volume as files using the keys in the Data field as the file names, unless the items element is populated with specific mappings of keys to paths. Note that this is identical to a configmap volume source without the default mode.`,
				},
				resource.Attribute{
					Name:        "downward_api",
					Description: `(Optional) Represents downward API info for projecting into a projected volume. Note that this is identical to a downward_api volume source without the default mode.`,
				},
				resource.Attribute{
					Name:        "secret",
					Description: `(Optional) Adapts a secret into a projected volume. The contents of the target Secret's Data field will be presented in a projected volume as files using the keys in the Data field as the file names. Note that this is identical to a secret volume source without the default mode.`,
				},
				resource.Attribute{
					Name:        "service_account_token",
					Description: `(Optional) Represents a projected service account token volume. This projection can be used to insert a service account token into the pods runtime filesystem for use against APIs (Kubernetes API Server or otherwise). ### ` + "`" + `service_account_token` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "audience",
					Description: `(Optional) Audience is the intended audience of the token. A recipient of a token must identify itself with an identifier specified in the audience of the token, and otherwise should reject the token. The audience defaults to the identifier of the apiserver.`,
				},
				resource.Attribute{
					Name:        "expiration_seconds",
					Description: `(Optional) The requested duration of validity of the service account token. As the token approaches expiration, the kubelet volume plugin will proactively rotate the service account token. The kubelet will start trying to rotate the token if the token is older than 80 percent of its time to live or if the token is older than 24 hours.Defaults to 1 hour and must be at least 10 minutes.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Required) Path is the path relative to the mount point of the file to project the token into. ### ` + "`" + `volume` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "aws_elastic_block_store",
					Description: `(Optional) Represents an AWS Disk resource that is attached to a kubelet's host machine and then exposed to the pod. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/volumes#awselasticblockstore)`,
				},
				resource.Attribute{
					Name:        "azure_disk",
					Description: `(Optional) Represents an Azure Data Disk mount on the host and bind mount to the pod.`,
				},
				resource.Attribute{
					Name:        "azure_file",
					Description: `(Optional) Represents an Azure File Service mount on the host and bind mount to the pod.`,
				},
				resource.Attribute{
					Name:        "ceph_fs",
					Description: `(Optional) Represents a Ceph FS mount on the host that shares a pod's lifetime`,
				},
				resource.Attribute{
					Name:        "cinder",
					Description: `(Optional) Represents a cinder volume attached and mounted on kubelets host machine. For more info see https://github.com/kubernetes/examples/blob/master/mysql-cinder-pd/README.md#mysql-installation-with-cinder-volume-plugin.`,
				},
				resource.Attribute{
					Name:        "config_map",
					Description: `(Optional) ConfigMap represents a configMap that should populate this volume`,
				},
				resource.Attribute{
					Name:        "downward_api",
					Description: `(Optional) DownwardAPI represents downward API about the pod that should populate this volume`,
				},
				resource.Attribute{
					Name:        "empty_dir",
					Description: `(Optional) EmptyDir represents a temporary directory that shares a pod's lifetime. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/volumes#emptydir)`,
				},
				resource.Attribute{
					Name:        "fc",
					Description: `(Optional) Represents a Fibre Channel resource that is attached to a kubelet's host machine and then exposed to the pod.`,
				},
				resource.Attribute{
					Name:        "flex_volume",
					Description: `(Optional) Represents a generic volume resource that is provisioned/attached using an exec based plugin. This is an alpha feature and may change in future.`,
				},
				resource.Attribute{
					Name:        "flocker",
					Description: `(Optional) Represents a Flocker volume attached to a kubelet's host machine and exposed to the pod for its usage. This depends on the Flocker control service being running`,
				},
				resource.Attribute{
					Name:        "gce_persistent_disk",
					Description: `(Optional) Represents a GCE Disk resource that is attached to a kubelet's host machine and then exposed to the pod. Provisioned by an admin. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/volumes#gcepersistentdisk)`,
				},
				resource.Attribute{
					Name:        "git_repo",
					Description: `(Optional) GitRepo represents a git repository at a particular revision.`,
				},
				resource.Attribute{
					Name:        "glusterfs",
					Description: `(Optional) Represents a Glusterfs volume that is attached to a host and exposed to the pod. Provisioned by an admin. For more info see https://github.com/kubernetes/examples/tree/master/volumes/glusterfs#glusterfs.`,
				},
				resource.Attribute{
					Name:        "host_path",
					Description: `(Optional) Represents a directory on the host. Provisioned by a developer or tester. This is useful for single-node development and testing only! On-host storage is not supported in any way and WILL NOT WORK in a multi-node cluster. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/volumes#hostpath)`,
				},
				resource.Attribute{
					Name:        "iscsi",
					Description: `(Optional) Represents an ISCSI Disk resource that is attached to a kubelet's host machine and then exposed to the pod. Provisioned by an admin.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Volume's name. Must be a DNS_LABEL and unique within the pod. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/identifiers#names)`,
				},
				resource.Attribute{
					Name:        "nfs",
					Description: `(Optional) Represents an NFS mount on the host. Provisioned by an admin. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/volumes#nfs)`,
				},
				resource.Attribute{
					Name:        "persistent_volume_claim",
					Description: `(Optional) The specification of a persistent volume.`,
				},
				resource.Attribute{
					Name:        "photon_persistent_disk",
					Description: `(Optional) Represents a PhotonController persistent disk attached and mounted on kubelets host machine`,
				},
				resource.Attribute{
					Name:        "quobyte",
					Description: `(Optional) Quobyte represents a Quobyte mount on the host that shares a pod's lifetime`,
				},
				resource.Attribute{
					Name:        "rbd",
					Description: `(Optional) Represents a Rados Block Device mount on the host that shares a pod's lifetime. For more info see https://kubernetes.io/docs/concepts/storage/volumes/#rbd.`,
				},
				resource.Attribute{
					Name:        "secret",
					Description: `(Optional) Secret represents a secret that should populate this volume. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/volumes#secrets)`,
				},
				resource.Attribute{
					Name:        "vsphere_volume",
					Description: `(Optional) Represents a vSphere volume attached and mounted on kubelets host machine ### ` + "`" + `volume_mount` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "mount_path",
					Description: `(Required) Path within the container at which the volume should be mounted. Must not contain ':'.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) This must match the Name of a Volume.`,
				},
				resource.Attribute{
					Name:        "read_only",
					Description: `(Optional) Mounted read-only if true, read-write otherwise (false or unspecified). Defaults to false.`,
				},
				resource.Attribute{
					Name:        "sub_path",
					Description: `(Optional) Path within the volume from which the container's volume should be mounted. Defaults to "" (volume's root).`,
				},
				resource.Attribute{
					Name:        "mount_propagation",
					Description: `(Optional) Mount propagation mode. Defaults to "None". For more info see [Kubernetes reference](https://kubernetes.io/docs/concepts/storage/volumes/#mount-propagation) ### ` + "`" + `vsphere_volume` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "fs_type",
					Description: `(Optional) Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.`,
				},
				resource.Attribute{
					Name:        "volume_path",
					Description: `(Required) Path that identifies vSphere volume vmdk ## Timeouts The following [Timeout](/docs/configuration/resources.html#operation-timeouts) configuration options are available for the ` + "`" + `kubernetes_deployment` + "`" + ` resource:`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Default ` + "`" + `10 minutes` + "`" + `) Used for creating new controller`,
				},
				resource.Attribute{
					Name:        "update",
					Description: `(Default ` + "`" + `10 minutes` + "`" + `) Used for updating a controller`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Default ` + "`" + `10 minutes` + "`" + `) Used for destroying a controller ## Import Deployment can be imported using the namespace and name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import kubernetes_deployment.example default/terraform-example ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "kubernetes_endpoints",
			Category:         "Resources",
			ShortDescription: `An Endpoints resource is an abstraction, linked to a Service, which defines the list of endpoints that actually implement the service.`,
			Description:      ``,
			Keywords: []string{
				"endpoints",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "metadata",
					Description: `(Required) Standard endpoints' metadata. For more info see [Kubernetes reference](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata)`,
				},
				resource.Attribute{
					Name:        "subset",
					Description: `(Optional) Set of addresses and ports that comprise a service. Can be repeated multiple times. ## Nested Blocks ### ` + "`" + `metadata` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "annotations",
					Description: `(Optional) An unstructured key value map stored with the endpoints resource that may be used to store arbitrary metadata. ~> By default, the provider ignores any annotations whose key names end with`,
				},
				resource.Attribute{
					Name:        "generate_name",
					Description: `(Optional) Prefix, used by the server, to generate a unique name ONLY IF the ` + "`" + `name` + "`" + ` field has not been provided. This value will also be combined with a unique suffix. For more info see [Kubernetes reference](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#idempotency)`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) Map of string keys and values that can be used to organize and categorize (scope and select) the endpoints resource. May match selectors of replication controllers and services. ~> By default, the provider ignores any labels whose key names end with`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the endpoints resource, must be unique. Cannot be updated. This name should correspond with an accompanying Service resource. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/identifiers#names)`,
				},
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) Namespace defines the space within which name of the endpoints resource must be unique. #### Attributes`,
				},
				resource.Attribute{
					Name:        "generation",
					Description: `A sequence number representing a specific generation of the desired state.`,
				},
				resource.Attribute{
					Name:        "resource_version",
					Description: `An opaque value that represents the internal version of this endpoints resource that can be used by clients to determine when endpoints resource has changed. For more info see [Kubernetes reference](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency)`,
				},
				resource.Attribute{
					Name:        "uid",
					Description: `The unique in time and space value for this endpoints resource. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/identifiers#uids) ### ` + "`" + `subset` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `(Optional) An IP address block which offers the related ports and is ready to accept traffic. These endpoints should be considered safe for load balancers and clients to utilize. Can be repeated multiple times.`,
				},
				resource.Attribute{
					Name:        "not_ready_address",
					Description: `(Optional) A IP address block which offers the related ports but is not currently marked as ready because it have not yet finished starting, have recently failed a readiness check, or have recently failed a liveness check. Can be repeated multiple times.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional) A port number block available on the related IP addresses. Can be repeated multiple times. ### ` + "`" + `address` + "`" + ` #### Attributes`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `The IP of this endpoint. May not be loopback (127.0.0.0/8), link-local (169.254.0.0/16), or link-local multicast ((224.0.0.0/24).`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `(Optional) The Hostname of this endpoint.`,
				},
				resource.Attribute{
					Name:        "node_name",
					Description: `(Optional) Node hosting this endpoint. This can be used to determine endpoints local to a node. ### ` + "`" + `not_ready_address` + "`" + ` #### Attributes`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `The IP of this endpoint. May not be loopback (127.0.0.0/8), link-local (169.254.0.0/16), or link-local multicast ((224.0.0.0/24).`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `(Optional) The Hostname of this endpoint.`,
				},
				resource.Attribute{
					Name:        "node_name",
					Description: `(Optional) Node hosting this endpoint. This can be used to determine endpoints local to a node. ### ` + "`" + `port` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of this port within the endpoint. All ports within the endpoint must have unique names. Optional if only one port is defined on this endpoint.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Required) The port that will be utilized by this endpoint.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Optional) The IP protocol for this port. Supports ` + "`" + `TCP` + "`" + ` and ` + "`" + `UDP` + "`" + `. Default is ` + "`" + `TCP` + "`" + `. ## Import An Endpoints resource can be imported using its namespace and name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import kubernetes_endpoints.example default/terraform-name ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "kubernetes_horizontal_pod_autoscaler",
			Category:         "Resources",
			ShortDescription: `Horizontal Pod Autoscaler automatically scales the number of pods in a replication controller, deployment or replica set based on observed CPU utilization.`,
			Description:      ``,
			Keywords: []string{
				"horizontal",
				"pod",
				"autoscaler",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "metadata",
					Description: `(Required) Standard horizontal pod autoscaler's metadata. For more info see [Kubernetes reference](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata)`,
				},
				resource.Attribute{
					Name:        "spec",
					Description: `(Required) Behaviour of the autoscaler. For more info see [Kubernetes reference](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#spec-and-status) ## Nested Blocks ### ` + "`" + `metadata` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "annotations",
					Description: `(Optional) An unstructured key value map stored with the horizontal pod autoscaler that may be used to store arbitrary metadata. ~> By default, the provider ignores any annotations whose key names end with`,
				},
				resource.Attribute{
					Name:        "generate_name",
					Description: `(Optional) Prefix, used by the server, to generate a unique name ONLY IF the ` + "`" + `name` + "`" + ` field has not been provided. This value will also be combined with a unique suffix. For more info see [Kubernetes reference](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#idempotency)`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) Map of string keys and values that can be used to organize and categorize (scope and select) the horizontal pod autoscaler. May match selectors of replication controllers and services. ~> By default, the provider ignores any labels whose key names end with`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the horizontal pod autoscaler, must be unique. Cannot be updated. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/identifiers#names)`,
				},
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) Namespace defines the space within which name of the horizontal pod autoscaler must be unique. #### Attributes`,
				},
				resource.Attribute{
					Name:        "generation",
					Description: `A sequence number representing a specific generation of the desired state.`,
				},
				resource.Attribute{
					Name:        "resource_version",
					Description: `An opaque value that represents the internal version of this horizontal pod autoscaler that can be used by clients to determine when horizontal pod autoscaler has changed. For more info see [Kubernetes reference](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency)`,
				},
				resource.Attribute{
					Name:        "uid",
					Description: `The unique in time and space value for this horizontal pod autoscaler. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/identifiers#uids) ### ` + "`" + `spec` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "max_replicas",
					Description: `(Required) Upper limit for the number of pods that can be set by the autoscaler.`,
				},
				resource.Attribute{
					Name:        "min_replicas",
					Description: `(Optional) Lower limit for the number of pods that can be set by the autoscaler, defaults to ` + "`" + `1` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "scale_target_ref",
					Description: `(Required) Reference to scaled resource. e.g. Replication Controller`,
				},
				resource.Attribute{
					Name:        "target_cpu_utilization_percentage",
					Description: `(Optional) Target average CPU utilization (represented as a percentage of requested CPU) over all the pods. If not specified the default autoscaling policy will be used.`,
				},
				resource.Attribute{
					Name:        "metric",
					Description: `(Optional) A metric on which to scale. ### ` + "`" + `metric` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) The type of metric. It can be one of "Object", "Pods", "Resource", or "External".`,
				},
				resource.Attribute{
					Name:        "object",
					Description: `(Optional) A metric describing a single kubernetes object (for example, hits-per-second on an Ingress object).`,
				},
				resource.Attribute{
					Name:        "pods",
					Description: `(Optional) A metric describing each pod in the current scale target (for example, transactions-processed-per-second). The values will be averaged together before being compared to the target value.`,
				},
				resource.Attribute{
					Name:        "resource",
					Description: `(Optional) A resource metric (such as those specified in requests and limits) known to Kubernetes describing each pod in the current scale target (e.g. CPU or memory). Such metrics are built in to Kubernetes, and have special scaling options on top of those available to normal per-pod metrics using the "pods" source.`,
				},
				resource.Attribute{
					Name:        "external",
					Description: `(Optional) A global metric that is not associated with any Kubernetes object. It allows autoscaling based on information coming from components running outside of cluster (for example length of queue in cloud messaging service, or QPS from loadbalancer running outside of cluster). ### Metric Type: ` + "`" + `external` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "metric",
					Description: `(Required) Identifies the target by name and selector.`,
				},
				resource.Attribute{
					Name:        "target",
					Description: `(Required) The target for the given metric. ### Metric Type: ` + "`" + `object` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "described_object",
					Description: `(Required) Reference to the object.`,
				},
				resource.Attribute{
					Name:        "metric",
					Description: `(Required) Identifies the target by name and selector.`,
				},
				resource.Attribute{
					Name:        "target",
					Description: `(Required) The target for the given metric. ### Metric Type: ` + "`" + `pods` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "metric",
					Description: `(Required) Identifies the target by name and selector.`,
				},
				resource.Attribute{
					Name:        "target",
					Description: `(Required) The target for the given metric. ### Metric Type: ` + "`" + `resource` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the resource in question.`,
				},
				resource.Attribute{
					Name:        "target",
					Description: `(Required) The target for the given metric. ### ` + "`" + `metric` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the given metric`,
				},
				resource.Attribute{
					Name:        "selector",
					Description: `(Optional) The label selector for the given metric ### ` + "`" + `target` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) Represents whether the metric type is Utilization, Value, or AverageValue.`,
				},
				resource.Attribute{
					Name:        "average_value",
					Description: `(Optional) The target value of the average of the metric across all relevant pods (as a quantity).`,
				},
				resource.Attribute{
					Name:        "average_utilization",
					Description: `(Optional) The target value of the average of the resource metric across all relevant pods, represented as a percentage of the requested value of the resource for the pods. Currently only valid for Resource metric source type.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Optional) value is the target value of the metric (as a quantity). #### Quantities See [here](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.17/#quantity-resource-core) for documentation on quantities. ### ` + "`" + `described_object` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "api_version",
					Description: `(Optional) API version of the referent`,
				},
				resource.Attribute{
					Name:        "kind",
					Description: `(Required) Kind of the referent. e.g. ` + "`" + `ReplicationController` + "`" + `. For more info see https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#types-kinds`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the referent. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/identifiers#names) ### ` + "`" + `scale_target_ref` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "api_version",
					Description: `(Optional) API version of the referent`,
				},
				resource.Attribute{
					Name:        "kind",
					Description: `(Required) Kind of the referent. e.g. ` + "`" + `ReplicationController` + "`" + `. For more info see https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#types-kinds`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the referent. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/identifiers#names) ## Import Horizontal Pod Autoscaler can be imported using the namespace and name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import kubernetes_horizontal_pod_autoscaler.example default/terraform-example ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "kubernetes_ingress",
			Category:         "Resources",
			ShortDescription: `Ingress is a collection of rules that allow inbound connections to reach the endpoints defined by a backend. An Ingress can be configured to give services externally-reachable urls, load balance traffic, terminate SSL, offer name based virtual hosting etc.`,
			Description:      ``,
			Keywords: []string{
				"ingress",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "metadata",
					Description: `(Required) Standard ingress's metadata. For more info: https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata`,
				},
				resource.Attribute{
					Name:        "spec",
					Description: `(Required) Spec defines the behavior of a ingress. https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#spec-and-status`,
				},
				resource.Attribute{
					Name:        "wait_for_load_balancer",
					Description: `(Optional) Terraform will wait for the load balancer to have at least 1 endpoint before considering the resource created. Defaults to ` + "`" + `false` + "`" + `. ## Nested Blocks ### ` + "`" + `metadata` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "annotations",
					Description: `(Optional) An unstructured key value map stored with the ingress that may be used to store arbitrary metadata. ~> By default, the provider ignores any annotations whose key names end with`,
				},
				resource.Attribute{
					Name:        "generate_name",
					Description: `(Optional) Prefix, used by the server, to generate a unique name ONLY IF the ` + "`" + `name` + "`" + ` field has not been provided. This value will also be combined with a unique suffix. Read more: https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#idempotency`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) Map of string keys and values that can be used to organize and categorize (scope and select) the service. May match selectors of replication controllers and services. ~> By default, the provider ignores any labels whose key names end with`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the service, must be unique. Cannot be updated. For more info: http://kubernetes.io/docs/user-guide/identifiers#names`,
				},
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) Namespace defines the space within which name of the service must be unique. #### Attributes`,
				},
				resource.Attribute{
					Name:        "generation",
					Description: `A sequence number representing a specific generation of the desired state.`,
				},
				resource.Attribute{
					Name:        "resource_version",
					Description: `An opaque value that represents the internal version of this service that can be used by clients to determine when service has changed. Read more: https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency`,
				},
				resource.Attribute{
					Name:        "uid",
					Description: `The unique in time and space value for this service. For more info: http://kubernetes.io/docs/user-guide/identifiers#uids ### ` + "`" + `spec` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "backend",
					Description: `(Optional) Backend defines the referenced service endpoint to which the traffic will be forwarded. See ` + "`" + `backend` + "`" + ` block attributes below.`,
				},
				resource.Attribute{
					Name:        "rule",
					Description: `(Optional) A list of host rules used to configure the Ingress. If unspecified, or no rule matches, all traffic is sent to the default backend. See ` + "`" + `rule` + "`" + ` block attributes below.`,
				},
				resource.Attribute{
					Name:        "tls",
					Description: `(Optional) TLS configuration. Currently the Ingress only supports a single TLS port, 443. If multiple members of this list specify different hosts, they will be multiplexed on the same port according to the hostname specified through the SNI TLS extension, if the ingress controller fulfilling the ingress supports SNI. See ` + "`" + `tls` + "`" + ` block attributes below. ### ` + "`" + `backend` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `(Optional) Specifies the name of the referenced service.`,
				},
				resource.Attribute{
					Name:        "service_port",
					Description: `(Optional) Specifies the port of the referenced service. ### ` + "`" + `rule` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `(Optional) Host is the fully qualified domain name of a network host, as defined by RFC 3986. Note the following deviations from the \"host\" part of the URI as defined in the RFC: 1. IPs are not allowed. Currently an IngressRuleValue can only apply to the IP in the Spec of the parent Ingress. 2. The : delimiter is not respected because ports are not allowed. Currently the port of an Ingress is implicitly :80 for http and :443 for https. Both these may change in the future. Incoming requests are matched against the host before the IngressRuleValue. If the host is unspecified, the Ingress routes all traffic based on the specified IngressRuleValue.`,
				},
				resource.Attribute{
					Name:        "http",
					Description: `(Required) http is a list of http selectors pointing to backends. In the example: http:///? -> backend where parts of the url correspond to RFC 3986, this resource will be used to match against everything after the last '/' and before the first '?' or '#'. See ` + "`" + `http` + "`" + ` block attributes below. #### ` + "`" + `http` + "`" + ``,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Required) Path array of path regex associated with a backend. Incoming urls matching the path are forwarded to the backend, see below for ` + "`" + `path` + "`" + ` block structure. #### ` + "`" + `path` + "`" + ``,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Required) A string or an extended POSIX regular expression as defined by IEEE Std 1003.1, (i.e this follows the egrep/unix syntax, not the perl syntax) matched against the path of an incoming request. Currently it can contain characters disallowed from the conventional \"path\" part of a URL as defined by RFC 3986. Paths must begin with a '/'. If unspecified, the path defaults to a catch all sending traffic to the backend.`,
				},
				resource.Attribute{
					Name:        "backend",
					Description: `(Required) Backend defines the referenced service endpoint to which the traffic will be forwarded to. ### ` + "`" + `tls` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "hosts",
					Description: `(Optional) Hosts are a list of hosts included in the TLS certificate. The values in this list must match the name/s used in the tlsSecret. Defaults to the wildcard host setting for the loadbalancer controller fulfilling this Ingress, if left unspecified.`,
				},
				resource.Attribute{
					Name:        "secret_name",
					Description: `(Optional) SecretName is the name of the secret used to terminate SSL traffic on 443. Field is left optional to allow SSL routing based on SNI hostname alone. If the SNI host in a listener conflicts with the \"Host\" header field used by an IngressRule, the SNI host is used for termination and value of the Host header is used for routing. ## Attributes ### ` + "`" + `status` + "`" + ``,
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
					Description: `Hostname is set for load-balancer ingress points that are DNS based (typically AWS load-balancers). ## Import Ingress can be imported using its namespace and name: ` + "`" + `` + "`" + `` + "`" + ` terraform import kubernetes_ingress.<TERRAFORM_RESOURCE_NAME> <KUBE_NAMESPACE>/<KUBE_INGRESS_NAME> ` + "`" + `` + "`" + `` + "`" + ` e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import kubernetes_ingress.example default/terraform-name ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "kubernetes_job",
			Category:         "Resources",
			ShortDescription: `A Job creates one or more Pods and ensures that a specified number of them successfully terminate. You can also use a Job to run multiple Pods in parallel.`,
			Description:      ``,
			Keywords: []string{
				"job",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "metadata",
					Description: `(Required) Standard resource's metadata. For more info: https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata`,
				},
				resource.Attribute{
					Name:        "spec",
					Description: `(Required) Specification of the desired behavior of a job. For more info: https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#spec-and-status`,
				},
				resource.Attribute{
					Name:        "wait_for_completion",
					Description: `(Optional) If ` + "`" + `true` + "`" + ` blocks job ` + "`" + `create` + "`" + ` or ` + "`" + `update` + "`" + ` until the status of the job has a ` + "`" + `Complete` + "`" + ` or ` + "`" + `Failed` + "`" + ` condition. Defaults to ` + "`" + `true` + "`" + `. ## Nested Blocks ### ` + "`" + `metadata` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "annotations",
					Description: `(Optional) An unstructured key value map stored with the resource that may be used to store arbitrary metadata. ~> By default, the provider ignores any annotations whose key names end with`,
				},
				resource.Attribute{
					Name:        "generate_name",
					Description: `(Optional) Prefix, used by the server, to generate a unique name ONLY IF the ` + "`" + `name` + "`" + ` field has not been provided. This value will also be combined with a unique suffix. Read more: https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#idempotency`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) Map of string keys and values that can be used to organize and categorize (scope and select) the service. May match selectors of replication controllers and services. ~> By default, the provider ignores any labels whose key names end with`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the service, must be unique. Cannot be updated. For more info: http://kubernetes.io/docs/user-guide/identifiers#names`,
				},
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) Namespace defines the space within which name of the service must be unique. #### Attributes`,
				},
				resource.Attribute{
					Name:        "generation",
					Description: `A sequence number representing a specific generation of the desired state.`,
				},
				resource.Attribute{
					Name:        "resource_version",
					Description: `An opaque value that represents the internal version of this service that can be used by clients to determine when service has changed. Read more: https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency`,
				},
				resource.Attribute{
					Name:        "uid",
					Description: `The unique in time and space value for this service. For more info: http://kubernetes.io/docs/user-guide/identifiers#uids ### ` + "`" + `spec` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "active_deadline_seconds",
					Description: `(Optional) Specifies the duration in seconds relative to the startTime that the job may be active before the system tries to terminate it; value must be positive integer.`,
				},
				resource.Attribute{
					Name:        "backoff_limit",
					Description: `(Optional) Specifies the number of retries before marking this job failed. Defaults to 6`,
				},
				resource.Attribute{
					Name:        "completions",
					Description: `(Optional) Specifies the desired number of successfully finished pods the job should be run with. Setting to nil means that the success of any pod signals the success of all pods, and allows parallelism to have any positive value. Setting to 1 means that parallelism is limited to 1 and the success of that pod signals the success of the job. For more info: https://kubernetes.io/docs/concepts/workloads/controllers/jobs-run-to-completion/`,
				},
				resource.Attribute{
					Name:        "manual_selector",
					Description: `(Optional) Controls generation of pod labels and pod selectors. Leave ` + "`" + `manualSelector` + "`" + ` unset unless you are certain what you are doing. When false or unset, the system pick labels unique to this job and appends those labels to the pod template. When true, the user is responsible for picking unique labels and specifying the selector. Failure to pick a unique label may cause this and other jobs to not function correctly. However, You may see ` + "`" + `manualSelector=true` + "`" + ` in jobs that were created with the old ` + "`" + `extensions/v1beta1` + "`" + ` API. For more info: https://kubernetes.io/docs/concepts/workloads/controllers/jobs-run-to-completion/#specifying-your-own-pod-selector`,
				},
				resource.Attribute{
					Name:        "parallelism",
					Description: `(Optional) Specifies the maximum desired number of pods the job should run at any given time. The actual number of pods running in steady state will be less than this number when ` + "`" + `((.spec.completions - .status.successful) < .spec.parallelism)` + "`" + `, i.e. when the work left to do is less than max parallelism. For more info: https://kubernetes.io/docs/concepts/workloads/controllers/jobs-run-to-completion/`,
				},
				resource.Attribute{
					Name:        "selector",
					Description: `(Optional) A label query over pods that should match the pod count. Normally, the system sets this field for you. For more info: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#label-selectors`,
				},
				resource.Attribute{
					Name:        "template",
					Description: `(Optional) Describes the pod that will be created when executing a job. For more info: https://kubernetes.io/docs/concepts/workloads/controllers/jobs-run-to-completion/`,
				},
				resource.Attribute{
					Name:        "ttl_seconds_after_finished",
					Description: `(Optional) ttlSecondsAfterFinished limits the lifetime of a Job that has finished execution (either Complete or Failed). If this field is set, ttlSecondsAfterFinished after the Job finishes, it is eligible to be automatically deleted. When the Job is being deleted, its lifecycle guarantees (e.g. finalizers) will be honored. If this field is unset, the Job won't be automatically deleted. If this field is set to zero, the Job becomes eligible to be deleted immediately after it finishes. ### ` + "`" + `selector` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "match_expressions",
					Description: `(Optional) A list of label selector requirements. The requirements are ANDed.`,
				},
				resource.Attribute{
					Name:        "match_labels",
					Description: `(Optional) A map of ` + "`" + `{key,value}` + "`" + ` pairs. A single ` + "`" + `{key,value}` + "`" + ` in the matchLabels map is equivalent to an element of matchExpressions, whose key field is "key", the operator is "In", and the values array contains only "value". The requirements are ANDed. ### ` + "`" + `template` + "`" + ` #### Arguments These arguments are the same as the for the ` + "`" + `spec` + "`" + ` block of a Pod. Please see the [Pod resource](pod.html#spec-1) for reference. ## Timeouts The following [Timeout](/docs/configuration/resources.html#operation-timeouts) configuration options are available for the ` + "`" + `kubernetes_job` + "`" + ` resource when used with ` + "`" + `wait_for_completion = true` + "`" + `:`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Default ` + "`" + `1 minute` + "`" + `) Used for creating a new job and waiting for a successful job completion.`,
				},
				resource.Attribute{
					Name:        "update",
					Description: `(Default ` + "`" + `1 minute` + "`" + `) Used for updating an existing job and waiting for a successful job completion. Note: - Kubernetes provider will treat update operations that change the Job spec resulting in the job re-run as "# forces replacement". In such cases, the ` + "`" + `create` + "`" + ` timeout value is used for both Create and Update operations. - ` + "`" + `wait_for_completion` + "`" + ` is not applicable during Delete operations; thus, there is no "delete" timeout value for Delete operation.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "kubernetes_limit_range",
			Category:         "Resources",
			ShortDescription: `Limit Range sets resource usage limits (e.g. memory, cpu, storage) for supported kinds of resources in a namespace.`,
			Description:      ``,
			Keywords: []string{
				"limit",
				"range",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "metadata",
					Description: `(Required) Standard limit range's metadata. For more info see [Kubernetes reference](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata)`,
				},
				resource.Attribute{
					Name:        "spec",
					Description: `(Optional) Spec defines the limits enforced. For more info see [Kubernetes reference](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#spec-and-status) ## Nested Blocks ### ` + "`" + `spec` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "limit",
					Description: `(Optional) The list of limits that are enforced. ### ` + "`" + `limit` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "default",
					Description: `(Optional) Default resource requirement limit value by resource name if resource limit is omitted.`,
				},
				resource.Attribute{
					Name:        "default_request",
					Description: `(Optional) The default resource requirement request value by resource name if resource request is omitted.`,
				},
				resource.Attribute{
					Name:        "max",
					Description: `(Optional) Max usage constraints on this kind by resource name.`,
				},
				resource.Attribute{
					Name:        "max_limit_request_ratio",
					Description: `(Optional) The named resource must have a request and limit that are both non-zero where limit divided by request is less than or equal to the enumerated value; this represents the max burst for the named resource.`,
				},
				resource.Attribute{
					Name:        "min",
					Description: `(Optional) Min usage constraints on this kind by resource name.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) Type of resource that this limit applies to. e.g. ` + "`" + `Pod` + "`" + `, ` + "`" + `Container` + "`" + ` or ` + "`" + `PersistentVolumeClaim` + "`" + ` ### ` + "`" + `metadata` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "annotations",
					Description: `(Optional) An unstructured key value map stored with the limit range that may be used to store arbitrary metadata. ~> By default, the provider ignores any annotations whose key names end with`,
				},
				resource.Attribute{
					Name:        "generate_name",
					Description: `(Optional) Prefix, used by the server, to generate a unique name ONLY IF the ` + "`" + `name` + "`" + ` field has not been provided. This value will also be combined with a unique suffix. For more info see [Kubernetes reference](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#idempotency)`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) Map of string keys and values that can be used to organize and categorize (scope and select) the limit range. May match selectors of replication controllers and services. ~> By default, the provider ignores any labels whose key names end with`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the limit range, must be unique. Cannot be updated. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/identifiers#names)`,
				},
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) Namespace defines the space within which name of the limit range must be unique. #### Attributes`,
				},
				resource.Attribute{
					Name:        "generation",
					Description: `A sequence number representing a specific generation of the desired state.`,
				},
				resource.Attribute{
					Name:        "resource_version",
					Description: `An opaque value that represents the internal version of this limit range that can be used by clients to determine when limit range has changed. For more info see [Kubernetes reference](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency)`,
				},
				resource.Attribute{
					Name:        "uid",
					Description: `The unique in time and space value for this limit range. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/identifiers#uids) ## Import Limit Range can be imported using its namespace and name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import kubernetes_limit_range.example default/terraform-example ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "kubernetes_namespace",
			Category:         "Resources",
			ShortDescription: `Kubernetes supports multiple virtual clusters backed by the same physical cluster. These virtual clusters are called namespaces.`,
			Description:      ``,
			Keywords: []string{
				"namespace",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "metadata",
					Description: `(Required) Standard namespace's [metadata](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata). ### Timeouts ` + "`" + `kubernetes_namespace` + "`" + ` provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `delete` + "`" + ` - Default ` + "`" + `5 minutes` + "`" + ` ## Nested Blocks ### ` + "`" + `metadata` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "annotations",
					Description: `(Optional) An unstructured key value map stored with the namespace that may be used to store arbitrary metadata. ~> By default, the provider ignores any annotations whose key names end with`,
				},
				resource.Attribute{
					Name:        "generate_name",
					Description: `(Optional) Prefix, used by the server, to generate a unique name ONLY IF the ` + "`" + `name` + "`" + ` field has not been provided. This value will also be combined with a unique suffix. Read more about [name idempotency](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#idempotency).`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) Map of string keys and values that can be used to organize and categorize (scope and select) namespaces. May match selectors of replication controllers and services. ~> By default, the provider ignores any labels whose key names end with`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the namespace, must be unique. Cannot be updated. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/identifiers#names) #### Attributes`,
				},
				resource.Attribute{
					Name:        "generation",
					Description: `A sequence number representing a specific generation of the desired state.`,
				},
				resource.Attribute{
					Name:        "resource_version",
					Description: `An opaque value that represents the internal version of this namespace that can be used by clients to determine when namespaces have changed. Read more about [concurrency control and consistency](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency).`,
				},
				resource.Attribute{
					Name:        "uid",
					Description: `The unique in time and space value for this namespace. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/identifiers#uids) ## Import Namespaces can be imported using their name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import kubernetes_namespace.n terraform-example-namespace ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "kubernetes_network_policy",
			Category:         "Resources",
			ShortDescription: `Kubernetes supports network policies to specify how groups of pods are allowed to communicate with each other and with other network endpoints. NetworkPolicy resources use labels to select pods and define rules which specify what traffic is allowed to the selected pods.`,
			Description:      ``,
			Keywords: []string{
				"network",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "metadata",
					Description: `(Required) Standard network policy's [metadata](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata). ## Nested Blocks ### ` + "`" + `metadata` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "annotations",
					Description: `(Optional) An unstructured key value map stored with the network policy that may be used to store arbitrary metadata. ~> By default, the provider ignores any annotations whose key names end with`,
				},
				resource.Attribute{
					Name:        "generate_name",
					Description: `(Optional) Prefix, used by the server, to generate a unique name ONLY IF the ` + "`" + `name` + "`" + ` field has not been provided. This value will also be combined with a unique suffix. Read more about [name idempotency](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#idempotency).`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) Map of string keys and values that can be used to organize and categorize (scope and select) network policies. May match selectors of replication controllers and services. ~> By default, the provider ignores any labels whose key names end with`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the network policy, must be unique. Cannot be updated. For more info: http://kubernetes.io/docs/user-guide/identifiers#names #### Attributes`,
				},
				resource.Attribute{
					Name:        "generation",
					Description: `A sequence number representing a specific generation of the desired state.`,
				},
				resource.Attribute{
					Name:        "resource_version",
					Description: `An opaque value that represents the internal version of this network policy that can be used by clients to determine when network policies have changed. Read more about [concurrency control and consistency](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency).`,
				},
				resource.Attribute{
					Name:        "uid",
					Description: `The unique in time and space value for this network policy. For more info: http://kubernetes.io/docs/user-guide/identifiers#uids ### ` + "`" + `spec` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "egress",
					Description: `(Optional) List of egress rules to be applied to the selected pods. Outgoing traffic is allowed if there are no NetworkPolicies selecting the pod (and cluster policy otherwise allows the traffic), OR if the traffic matches at least one egress rule across all of the NetworkPolicy objects whose podSelector matches the pod. If this block is empty then this NetworkPolicy allows all outgoing traffic. If this block is omitted then this NetworkPolicy does not allow any outgoing traffic (and serves solely to ensure that the pods it selects are isolated by default).`,
				},
				resource.Attribute{
					Name:        "ingress",
					Description: `(Optional) List of ingress rules to be applied to the selected pods. Traffic is allowed to a pod if there are no NetworkPolicies selecting the pod (and cluster policy otherwise allows the traffic), OR if the traffic source is the pod's local node, OR if the traffic matches at least one ingress rule across all of the NetworkPolicy objects whose podSelector matches the pod. If this block is empty then this NetworkPolicy allows all incoming traffic. If this block is omitted then this NetworkPolicy does not allow any incoming traffic (and serves solely to ensure that the pods it selects are isolated by default).`,
				},
				resource.Attribute{
					Name:        "pod_selector",
					Description: `(Required) Selects the pods to which this NetworkPolicy object applies. The array of ingress rules is applied to any pods selected by this field. Multiple network policies can select the same set of pods. In this case, the ingress rules for each are combined additively. This field is NOT optional and follows standard label selector semantics. An empty podSelector matches all pods in this namespace.`,
				},
				resource.Attribute{
					Name:        "from",
					Description: `(Optional) List of sources which should be able to access the pods selected for this rule. Items in this list are combined using a logical OR operation. If this field is empty or missing, this rule matches all sources (traffic not restricted by source). If this field is present and contains at least on item, this rule allows traffic only if the traffic matches at least one item in the from list.`,
				},
				resource.Attribute{
					Name:        "ports",
					Description: `(Optional) List of ports which should be made accessible on the pods selected for this rule. Each item in this list is combined using a logical OR. If this field is empty or missing, this rule matches all ports (traffic not restricted by port). If this field is present and contains at least one item, then this rule allows traffic only if the traffic matches at least one port in the list. ### ` + "`" + `egress` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "to",
					Description: `(Optional) List of destinations for outgoing traffic of pods selected for this rule. Items in this list are combined using a logical OR operation. If this field is empty or missing, this rule matches all destinations (traffic not restricted by destination). If this field is present and contains at least one item, this rule allows traffic only if the traffic matches at least one item in the to list.`,
				},
				resource.Attribute{
					Name:        "ports",
					Description: `(Optional) List of destination ports for outgoing traffic. Each item in this list is combined using a logical OR. If this field is empty or missing, this rule matches all ports (traffic not restricted by port). If this field is present and contains at least one item, then this rule allows traffic only if the traffic matches at least one port in the list. ### ` + "`" + `from` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "namespace_selector",
					Description: `(Optional) Selects Namespaces using cluster scoped-labels. This matches all pods in all namespaces selected by this label selector. This field follows standard label selector semantics. If present but empty, this selector selects all namespaces.`,
				},
				resource.Attribute{
					Name:        "pod_selector",
					Description: `(Optional) This is a label selector which selects Pods in this namespace. This field follows standard label selector semantics. If present but empty, this selector selects all pods in this namespace. ### ` + "`" + `ports` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional) The port on the given protocol. This can either be a numerical or named port on a pod. If this field is not provided, this matches all port names and numbers.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Optional) The protocol (TCP or UDP) which traffic must match. If not specified, this field defaults to TCP. ### ` + "`" + `to` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "ip_block",
					Description: `(Optional) IPBlock defines policy on a particular IPBlock`,
				},
				resource.Attribute{
					Name:        "namespace_selector",
					Description: `(Optional) Selects Namespaces using cluster scoped-labels. This matches all pods in all namespaces selected by this label selector. This field follows standard label selector semantics. If present but empty, this selector selects all namespaces.`,
				},
				resource.Attribute{
					Name:        "pod_selector",
					Description: `(Optional) This is a label selector which selects Pods in this namespace. This field follows standard label selector semantics. If present but empty, this selector selects all pods in this namespace. ### ` + "`" + `ip_block` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "cidr",
					Description: `(Optional) CIDR is a string representing the IP Block Valid examples are "192.168.1.1/24"`,
				},
				resource.Attribute{
					Name:        "except",
					Description: `(Optional) Except is a slice of CIDRs that should not be included within an IP Block. Valid examples are "192.168.1.1/24". Except values will be rejected if they are outside the CIDR range. ### ` + "`" + `namespace_selector` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "match_expressions",
					Description: `(Optional) A list of label selector requirements. The requirements are ANDed.`,
				},
				resource.Attribute{
					Name:        "match_labels",
					Description: `(Optional) A map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of ` + "`" + `match_expressions` + "`" + `, whose key field is "key", the operator is "In", and the values array contains only "value". The requirements are ANDed. ### ` + "`" + `pod_selector` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "match_expressions",
					Description: `(Optional) A list of label selector requirements. The requirements are ANDed.`,
				},
				resource.Attribute{
					Name:        "match_labels",
					Description: `(Optional) A map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of ` + "`" + `match_expressions` + "`" + `, whose key field is "key", the operator is "In", and the values array contains only "value". The requirements are ANDed. ### ` + "`" + `match_expressions` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Optional) The label key that the selector applies to.`,
				},
				resource.Attribute{
					Name:        "operator",
					Description: `(Optional) A key's relationship to a set of values. Valid operators ard ` + "`" + `In` + "`" + `, ` + "`" + `NotIn` + "`" + `, ` + "`" + `Exists` + "`" + ` and ` + "`" + `DoesNotExist` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Optional) An array of string values. If the operator is ` + "`" + `In` + "`" + ` or ` + "`" + `NotIn` + "`" + `, the values array must be non-empty. If the operator is ` + "`" + `Exists` + "`" + ` or ` + "`" + `DoesNotExist` + "`" + `, the values array must be empty. This array is replaced during a strategic merge patch. ## Import Network policies can be imported using their identifier consisting of ` + "`" + `<namespace-name>/<network-policy-name>` + "`" + `, e.g.: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import kubernetes_network_policy.example default/terraform-example-network-policy ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "kubernetes_persistent_volume",
			Category:         "Resources",
			ShortDescription: `A Persistent Volume (PV) is a piece of networked storage in the cluster that has been provisioned by an administrator.`,
			Description:      ``,
			Keywords: []string{
				"persistent",
				"volume",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "metadata",
					Description: `(Required) Standard persistent volume's metadata. For more info see [Kubernetes reference](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata)`,
				},
				resource.Attribute{
					Name:        "spec",
					Description: `(Required) Spec of the persistent volume owned by the cluster. See below. ## Nested Blocks ### ` + "`" + `spec` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "access_modes",
					Description: `(Required) Contains all ways the volume can be mounted. Valid values are ` + "`" + `ReadWriteOnce` + "`" + `, ` + "`" + `ReadOnlyMany` + "`" + `, ` + "`" + `ReadWriteMany` + "`" + `. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/persistent-volumes#access-modes)`,
				},
				resource.Attribute{
					Name:        "capacity",
					Description: `(Required) A description of the persistent volume's resources and capacity. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/persistent-volumes#capacity)`,
				},
				resource.Attribute{
					Name:        "node_affinity",
					Description: `(Optional) NodeAffinity defines constraints that limit what nodes this volume can be accessed from. This field influences the scheduling of pods that use this volume.`,
				},
				resource.Attribute{
					Name:        "persistent_volume_reclaim_policy",
					Description: `(Optional) What happens to a persistent volume when released from its claim. Valid options are Retain (default), Delete and Recycle. Recycling must be supported by the volume plugin underlying this persistent volume. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/persistent-volumes#recycling-policy)`,
				},
				resource.Attribute{
					Name:        "persistent_volume_source",
					Description: `(Required) The specification of a persistent volume.`,
				},
				resource.Attribute{
					Name:        "storage_class_name",
					Description: `(Optional) The name of the persistent volume's storage class. For more info see [Kubernetes reference](https://kubernetes.io/docs/concepts/storage/persistent-volumes/#class)`,
				},
				resource.Attribute{
					Name:        "mount_options",
					Description: `(Optional) A Kubernetes administrator can specify additional mount options for when a Persistent Volume is mounted on a node. ~> Not all Persistent Volume types support mount options. For more info see [Kubernetes reference](https://kubernetes.io/docs/concepts/storage/persistent-volumes/#mount-options)`,
				},
				resource.Attribute{
					Name:        "volume_mode",
					Description: `(Optional) Defines if a volume is used with a formatted filesystem or to remain in raw block state. Possible values are ` + "`" + `Block` + "`" + ` and ` + "`" + `Filesystem` + "`" + `. Default value is ` + "`" + `Filesystem` + "`" + `. For more info see [Kubernetes reference](https://kubernetes.io/docs/concepts/storage/persistent-volumes/#volume-mode) ### ` + "`" + `node_affinity` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "required",
					Description: `(Optional) Required specifies hard node constraints that must be met. ### ` + "`" + `required` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "node_selector_term",
					Description: `(Required) A list of node selector terms. The terms are ORed. ### ` + "`" + `node_selector_term` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "match_expressions",
					Description: `(Optional) A list of node selector requirements by node's labels.`,
				},
				resource.Attribute{
					Name:        "match_fields",
					Description: `(Optional) A list of node selector requirements by node's fields. ### ` + "`" + `match_expressions` + "`" + ` and ` + "`" + `match_fields` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) The label key that the selector applies to.`,
				},
				resource.Attribute{
					Name:        "operator",
					Description: `(Required) Represents a key's relationship to a set of values. Valid operators are ` + "`" + `In` + "`" + `, ` + "`" + `NotIn` + "`" + `, ` + "`" + `Exists` + "`" + `, ` + "`" + `DoesNotExist` + "`" + `. ` + "`" + `Gt` + "`" + `, and ` + "`" + `Lt` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Optional) An array of string values. If the operator is ` + "`" + `In` + "`" + ` or ` + "`" + `NotIn` + "`" + `, the values array must be non-empty. If the operator is ` + "`" + `Exists` + "`" + ` or ` + "`" + `DoesNotExist` + "`" + `, the values array must be empty. If the operator is ` + "`" + `Gt` + "`" + ` or ` + "`" + `Lt` + "`" + `, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch. ### ` + "`" + `persistent_volume_source` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "aws_elastic_block_store",
					Description: `(Optional) Represents an AWS Disk resource that is attached to a kubelet's host machine and then exposed to the pod. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/volumes#awselasticblockstore)`,
				},
				resource.Attribute{
					Name:        "azure_disk",
					Description: `(Optional) Represents an Azure Data Disk mount on the host and bind mount to the pod.`,
				},
				resource.Attribute{
					Name:        "azure_file",
					Description: `(Optional) Represents an Azure File Service mount on the host and bind mount to the pod.`,
				},
				resource.Attribute{
					Name:        "ceph_fs",
					Description: `(Optional) Represents a Ceph FS mount on the host that shares a pod's lifetime.`,
				},
				resource.Attribute{
					Name:        "cinder",
					Description: `(Optional) Represents a cinder volume attached and mounted on kubelets host machine. For more info see https://github.com/kubernetes/examples/tree/master/mysql-cinder-pd#mysql-installation-with-cinder-volume-plugin.`,
				},
				resource.Attribute{
					Name:        "csi",
					Description: `(Optional) CSI represents storage that is handled by an external CSI driver. For more info see [Kubernetes reference](https://kubernetes.io/docs/concepts/storage/volumes/#csi).`,
				},
				resource.Attribute{
					Name:        "fc",
					Description: `(Optional) Represents a Fibre Channel resource that is attached to a kubelet's host machine and then exposed to the pod.`,
				},
				resource.Attribute{
					Name:        "flex_volume",
					Description: `(Optional) Represents a generic volume resource that is provisioned/attached using an exec based plugin. This is an alpha feature and may change in future.`,
				},
				resource.Attribute{
					Name:        "flocker",
					Description: `(Optional) Represents a Flocker volume attached to a kubelet's host machine and exposed to the pod for its usage. This depends on the Flocker control service being running.`,
				},
				resource.Attribute{
					Name:        "gce_persistent_disk",
					Description: `(Optional) Represents a GCE Disk resource that is attached to a kubelet's host machine and then exposed to the pod. Provisioned by an admin. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/volumes#gcepersistentdisk).`,
				},
				resource.Attribute{
					Name:        "glusterfs",
					Description: `(Optional) Represents a Glusterfs volume that is attached to a host and exposed to the pod. Provisioned by an admin. For more info see https://github.com/kubernetes/examples/tree/master/volumes/glusterfs#glusterfs.`,
				},
				resource.Attribute{
					Name:        "host_path",
					Description: `(Optional) Represents a directory on the host. Provisioned by a developer or tester. This is useful for single-node development and testing only! On-host storage is not supported in any way and WILL NOT WORK in a multi-node cluster. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/volumes#hostpath)`,
				},
				resource.Attribute{
					Name:        "iscsi",
					Description: `(Optional) Represents an ISCSI Disk resource that is attached to a kubelet's host machine and then exposed to the pod. Provisioned by an admin.`,
				},
				resource.Attribute{
					Name:        "local",
					Description: `(Optional) Represents a local storage volume on the host. Provisioned by an admin. For more info see [Kubernetes reference](https://kubernetes.io/docs/concepts/storage/volumes/#local).`,
				},
				resource.Attribute{
					Name:        "nfs",
					Description: `(Optional) Represents an NFS mount on the host. Provisioned by an admin. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/volumes#nfs).`,
				},
				resource.Attribute{
					Name:        "photon_persistent_disk",
					Description: `(Optional) Represents a PhotonController persistent disk attached and mounted on kubelets host machine.`,
				},
				resource.Attribute{
					Name:        "quobyte",
					Description: `(Optional) Quobyte represents a Quobyte mount on the host that shares a pod's lifetime.`,
				},
				resource.Attribute{
					Name:        "rbd",
					Description: `(Optional) Represents a Rados Block Device mount on the host that shares a pod's lifetime. For more info see https://kubernetes.io/docs/concepts/storage/volumes/#rbd.`,
				},
				resource.Attribute{
					Name:        "vsphere_volume",
					Description: `(Optional) Represents a vSphere volume attached and mounted on kubelets host machine. ### ` + "`" + `aws_elastic_block_store` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "fs_type",
					Description: `(Optional) Filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/volumes#awselasticblockstore)`,
				},
				resource.Attribute{
					Name:        "partition",
					Description: `(Optional) The partition in the volume that you want to mount. If omitted, the default is to mount by volume name. Examples: For volume /dev/sda1, you specify the partition as "1". Similarly, the volume partition for /dev/sda is "0" (or you can leave the property empty).`,
				},
				resource.Attribute{
					Name:        "read_only",
					Description: `(Optional) Whether to set the read-only property in VolumeMounts to "true". If omitted, the default is "false". For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/volumes#awselasticblockstore)`,
				},
				resource.Attribute{
					Name:        "volume_id",
					Description: `(Required) Unique ID of the persistent disk resource in AWS (Amazon EBS volume). For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/volumes#awselasticblockstore) ### ` + "`" + `azure_disk` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "caching_mode",
					Description: `(Required) Host Caching mode: None, Read Only, Read Write.`,
				},
				resource.Attribute{
					Name:        "data_disk_uri",
					Description: `(Required) The URI the data disk in the blob storage OR the resource ID of an Azure managed data disk if ` + "`" + `kind` + "`" + ` is ` + "`" + `Managed` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "disk_name",
					Description: `(Required) The Name of the data disk in the blob storage OR the name of an Azure managed data disk if ` + "`" + `kind` + "`" + ` is ` + "`" + `Managed` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "fs_type",
					Description: `(Optional) Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.`,
				},
				resource.Attribute{
					Name:        "read_only",
					Description: `(Optional) Whether to force the read-only setting in VolumeMounts. Defaults to false (read/write).`,
				},
				resource.Attribute{
					Name:        "kind",
					Description: `(Optional) The type for the data disk. Expected values: ` + "`" + `Shared` + "`" + `, ` + "`" + `Dedicated` + "`" + `, ` + "`" + `Managed` + "`" + `. Defaults to ` + "`" + `Shared` + "`" + `. ### ` + "`" + `azure_file` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "read_only",
					Description: `(Optional) Whether to force the read-only setting in VolumeMounts. Defaults to false (read/write).`,
				},
				resource.Attribute{
					Name:        "secret_name",
					Description: `(Required) The name of secret that contains Azure Storage Account Name and Key.`,
				},
				resource.Attribute{
					Name:        "secret_namespace",
					Description: `(Optional) The namespace of the secret that contains Azure Storage Account Name and Key. For Kubernetes up to 1.18.x the default is the same as the Pod. For Kubernetes 1.19.x and later the default is \"default\" namespace.`,
				},
				resource.Attribute{
					Name:        "share_name",
					Description: `(Required) Share Name ### ` + "`" + `ceph_fs` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "monitors",
					Description: `(Required) Monitors is a collection of Ceph monitors. For more info see https://github.com/kubernetes/examples/tree/master/volumes/cephfs/#how-to-use-it.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Optional) Used as the mounted root, rather than the full Ceph tree, default is /.`,
				},
				resource.Attribute{
					Name:        "read_only",
					Description: `(Optional) Whether to force the read-only setting in VolumeMounts. Defaults to ` + "`" + `false` + "`" + ` (read/write). For more info see https://github.com/kubernetes/examples/tree/master/volumes/cephfs/#how-to-use-it.`,
				},
				resource.Attribute{
					Name:        "secret_file",
					Description: `(Optional) The path to key ring for User, default is /etc/ceph/user.secret. For more info see https://github.com/kubernetes/examples/tree/master/volumes/cephfs/#how-to-use-it.`,
				},
				resource.Attribute{
					Name:        "secret_ref",
					Description: `(Optional) Reference to the authentication secret for User, default is empty. sFor more info see https://github.com/kubernetes/examples/tree/master/volumes/cephfs/#how-to-use-it. see [secret_ref](#secret_ref) for more details.`,
				},
				resource.Attribute{
					Name:        "user",
					Description: `(Optional) User is the rados user name, default is admin. For more info see https://github.com/kubernetes/examples/tree/master/volumes/cephfs/#how-to-use-it. ### ` + "`" + `cinder` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "fs_type",
					Description: `(Optional) Filesystem type to mount. Must be a filesystem type supported by the host operating system. Examples: "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified. For more info see https://github.com/kubernetes/examples/blob/master/mysql-cinder-pd/README.md#mysql-installation-with-cinder-volume-plugin.`,
				},
				resource.Attribute{
					Name:        "read_only",
					Description: `(Optional) Whether to force the read-only setting in VolumeMounts. Defaults to false (read/write). For more info see https://github.com/kubernetes/examples/blob/master/mysql-cinder-pd/README.md#mysql-installation-with-cinder-volume-plugin.`,
				},
				resource.Attribute{
					Name:        "volume_id",
					Description: `(Required) Volume ID used to identify the volume in Cinder. For more info see https://github.com/kubernetes/examples/blob/master/mysql-cinder-pd/README.md#mysql-installation-with-cinder-volume-plugin. ### ` + "`" + `csi` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "driver",
					Description: `(Required) the name of the volume driver to use. For more info see [Kubernetes reference](https://kubernetes.io/docs/concepts/storage/volumes/#csi).`,
				},
				resource.Attribute{
					Name:        "volume_handle",
					Description: `(Required) A map that specifies static properties of a volume. For more info see [Kubernetes reference](https://kubernetes.io/docs/concepts/storage/volumes/#csi).`,
				},
				resource.Attribute{
					Name:        "volume_attributes",
					Description: `(Optional) Attributes of the volume to publish.`,
				},
				resource.Attribute{
					Name:        "fs_type",
					Description: `(Optional) Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. ` + "`" + `ext4` + "`" + `, ` + "`" + `xfs` + "`" + `, ` + "`" + `ntfs` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "read_only",
					Description: `(Optional) Whether to set the read-only property in VolumeMounts to ` + "`" + `true` + "`" + `. If omitted, the default is ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "controller_publish_secret_ref",
					Description: `(Optional) A reference to the secret object containing sensitive information to pass to the CSI driver to complete the CSI ControllerPublishVolume and ControllerUnpublishVolume calls. see [secret_ref](#secret_ref) for more details.`,
				},
				resource.Attribute{
					Name:        "node_stage_secret_ref",
					Description: `(Optional) A reference to the secret object containing sensitive information to pass to the CSI driver to complete the CSI NodeStageVolume and NodeStageVolume and NodeUnstageVolume calls. see [secret_ref](#secret_ref) for more details.`,
				},
				resource.Attribute{
					Name:        "node_publish_secret_ref",
					Description: `(Optional) A reference to the secret object containing sensitive information to pass to the CSI driver to complete the CSI NodePublishVolume and NodeUnpublishVolume calls. see [secret_ref](#secret_ref) for more details.`,
				},
				resource.Attribute{
					Name:        "controller_expand_secret_ref",
					Description: `(Optional) A reference to the secret object containing sensitive information to pass to the CSI driver to complete the CSI ControllerExpandVolume call. see [secret_ref](#secret_ref) for more details. ### ` + "`" + `fc` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "fs_type",
					Description: `(Optional) Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.`,
				},
				resource.Attribute{
					Name:        "lun",
					Description: `(Required) FC target lun number`,
				},
				resource.Attribute{
					Name:        "read_only",
					Description: `(Optional) Whether to force the read-only setting in VolumeMounts. Defaults to false (read/write).`,
				},
				resource.Attribute{
					Name:        "target_ww_ns",
					Description: `(Required) FC target worldwide names (WWNs) ### ` + "`" + `flex_volume` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "driver",
					Description: `(Required) Driver is the name of the driver to use for this volume.`,
				},
				resource.Attribute{
					Name:        "fs_type",
					Description: `(Optional) Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs", "ntfs". The default filesystem depends on FlexVolume script.`,
				},
				resource.Attribute{
					Name:        "options",
					Description: `(Optional) Extra command options if any.`,
				},
				resource.Attribute{
					Name:        "read_only",
					Description: `(Optional) Whether to force the ReadOnly setting in VolumeMounts. Defaults to false (read/write).`,
				},
				resource.Attribute{
					Name:        "secret_ref",
					Description: `(Optional) Reference to the secret object containing sensitive information to pass to the plugin scripts. This may be empty if no secret object is specified. If the secret object contains more than one secret, all secrets are passed to the plugin scripts. see [secret_ref](#secret_ref) for more details. ### ` + "`" + `flocker` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "dataset_name",
					Description: `(Optional) Name of the dataset stored as metadata -> name on the dataset for Flocker should be considered as deprecated`,
				},
				resource.Attribute{
					Name:        "dataset_uuid",
					Description: `(Optional) UUID of the dataset. This is unique identifier of a Flocker dataset ### ` + "`" + `gce_persistent_disk` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "fs_type",
					Description: `(Optional) Filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/volumes#gcepersistentdisk)`,
				},
				resource.Attribute{
					Name:        "partition",
					Description: `(Optional) The partition in the volume that you want to mount. If omitted, the default is to mount by volume name. Examples: For volume /dev/sda1, you specify the partition as "1". Similarly, the volume partition for /dev/sda is "0" (or you can leave the property empty). For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/volumes#gcepersistentdisk)`,
				},
				resource.Attribute{
					Name:        "pd_name",
					Description: `(Required) Unique name of the PD resource in GCE. Used to identify the disk in GCE. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/volumes#gcepersistentdisk)`,
				},
				resource.Attribute{
					Name:        "read_only",
					Description: `(Optional) Whether to force the ReadOnly setting in VolumeMounts. Defaults to false. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/volumes#gcepersistentdisk) ### ` + "`" + `glusterfs` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "endpoints_name",
					Description: `(Required) The endpoint name that details Glusterfs topology. For more info see https://github.com/kubernetes/examples/tree/master/volumes/glusterfs#create-a-pod.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Required) The Glusterfs volume path. For more info see https://github.com/kubernetes/examples/tree/master/volumes/glusterfs#create-a-pod.`,
				},
				resource.Attribute{
					Name:        "read_only",
					Description: `(Optional) Whether to force the Glusterfs volume to be mounted with read-only permissions. Defaults to false. For more info see https://github.com/kubernetes/examples/tree/master/volumes/glusterfs#create-a-pod. ### ` + "`" + `host_path` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Optional) Path of the directory on the host. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/volumes#hostpath)`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) Type for HostPath volume. Defaults to "". For more info see [Kubernetes reference](https://kubernetes.io/docs/concepts/storage/volumes#hostpath) ### ` + "`" + `iscsi` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "fs_type",
					Description: `(Optional) Filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/volumes#iscsi)`,
				},
				resource.Attribute{
					Name:        "iqn",
					Description: `(Required) Target iSCSI Qualified Name.`,
				},
				resource.Attribute{
					Name:        "iscsi_interface",
					Description: `(Optional) iSCSI interface name that uses an iSCSI transport. Defaults to 'default' (tcp).`,
				},
				resource.Attribute{
					Name:        "lun",
					Description: `(Optional) iSCSI target lun number.`,
				},
				resource.Attribute{
					Name:        "read_only",
					Description: `(Optional) Whether to force the read-only setting in VolumeMounts. Defaults to false.`,
				},
				resource.Attribute{
					Name:        "target_portal",
					Description: `(Required) iSCSI target portal. The portal is either an IP or ip_addr:port if the port is other than default (typically TCP ports 860 and 3260). ### ` + "`" + `local` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Optional) Path of the directory on the host. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/volumes#local) ### ` + "`" + `metadata` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "annotations",
					Description: `(Optional) An unstructured key value map stored with the persistent volume that may be used to store arbitrary metadata. ~> By default, the provider ignores any annotations whose key names end with`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) Map of string keys and values that can be used to organize and categorize (scope and select) the persistent volume. May match selectors of replication controllers and services. ~> By default, the provider ignores any labels whose key names end with`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the persistent volume, must be unique. Cannot be updated. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/identifiers#names) #### Attributes`,
				},
				resource.Attribute{
					Name:        "generation",
					Description: `A sequence number representing a specific generation of the desired state.`,
				},
				resource.Attribute{
					Name:        "resource_version",
					Description: `An opaque value that represents the internal version of this persistent volume that can be used by clients to determine when persistent volume has changed. For more info see [Kubernetes reference](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency)`,
				},
				resource.Attribute{
					Name:        "uid",
					Description: `The unique in time and space value for this persistent volume. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/identifiers#uids) ### ` + "`" + `nfs` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Required) Path that is exported by the NFS server. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/volumes#nfs)`,
				},
				resource.Attribute{
					Name:        "read_only",
					Description: `(Optional) Whether to force the NFS export to be mounted with read-only permissions. Defaults to false. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/volumes#nfs)`,
				},
				resource.Attribute{
					Name:        "server",
					Description: `(Required) Server is the hostname or IP address of the NFS server. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/volumes#nfs) ### ` + "`" + `photon_persistent_disk` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "fs_type",
					Description: `(Optional) Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.`,
				},
				resource.Attribute{
					Name:        "pd_id",
					Description: `(Required) ID that identifies Photon Controller persistent disk ### ` + "`" + `quobyte` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "group",
					Description: `(Optional) Group to map volume access to Default is no group`,
				},
				resource.Attribute{
					Name:        "read_only",
					Description: `(Optional) Whether to force the Quobyte volume to be mounted with read-only permissions. Defaults to false.`,
				},
				resource.Attribute{
					Name:        "registry",
					Description: `(Required) Registry represents a single or multiple Quobyte Registry services specified as a string as host:port pair (multiple entries are separated with commas) which acts as the central registry for volumes`,
				},
				resource.Attribute{
					Name:        "user",
					Description: `(Optional) User to map volume access to Defaults to serivceaccount user`,
				},
				resource.Attribute{
					Name:        "volume",
					Description: `(Required) Volume is a string that references an already created Quobyte volume by name. ### ` + "`" + `rbd` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "ceph_monitors",
					Description: `(Required) A collection of Ceph monitors. For more info see https://github.com/kubernetes/examples/tree/master/volumes/cephfs/#how-to-use-it.`,
				},
				resource.Attribute{
					Name:        "fs_type",
					Description: `(Optional) Filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/volumes#rbd)`,
				},
				resource.Attribute{
					Name:        "keyring",
					Description: `(Optional) Keyring is the path to key ring for RBDUser. Default is /etc/ceph/keyring. For more info see https://github.com/kubernetes/examples/tree/master/volumes/cephfs/#how-to-use-it.`,
				},
				resource.Attribute{
					Name:        "rados_user",
					Description: `(Optional) The rados user name. Default is admin. For more info see https://github.com/kubernetes/examples/tree/master/volumes/cephfs/#how-to-use-it.`,
				},
				resource.Attribute{
					Name:        "rbd_image",
					Description: `(Required) The rados image name. For more info see https://github.com/kubernetes/examples/tree/master/volumes/cephfs/#how-to-use-it.`,
				},
				resource.Attribute{
					Name:        "rbd_pool",
					Description: `(Optional) The rados pool name. Default is rbd. For more info see https://github.com/kubernetes/examples/tree/master/volumes/cephfs/#how-to-use-it.`,
				},
				resource.Attribute{
					Name:        "read_only",
					Description: `(Optional) Whether to force the read-only setting in VolumeMounts. Defaults to false. For more info see https://github.com/kubernetes/examples/tree/master/volumes/cephfs/#how-to-use-it.`,
				},
				resource.Attribute{
					Name:        "secret_ref",
					Description: `(Optional) Name of the authentication secret for RBDUser. If provided overrides keyring. Default is nil. For more info see https://github.com/kubernetes/examples/tree/master/volumes/cephfs/#how-to-use-it. see [secret_ref](#secret_ref) for more details. ### ` + "`" + `secret_ref` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the referent. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/identifiers#names)`,
				},
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) The Namespace of the referent secret. ### ` + "`" + `vsphere_volume` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "fs_type",
					Description: `(Optional) Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.`,
				},
				resource.Attribute{
					Name:        "volume_path",
					Description: `(Required) Path that identifies vSphere volume vmdk ## Import Persistent Volume can be imported using its name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import kubernetes_persistent_volume.example terraform-example ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "kubernetes_persistent_volume_claim",
			Category:         "Resources",
			ShortDescription: `This resource allows the user to request for and claim to a persistent volume.`,
			Description:      ``,
			Keywords: []string{
				"persistent",
				"volume",
				"claim",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "metadata",
					Description: `(Required) Standard persistent volume claim's metadata. For more info see [Kubernetes reference](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata)`,
				},
				resource.Attribute{
					Name:        "spec",
					Description: `(Required) Spec defines the desired characteristics of a volume requested by a pod author. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/persistent-volumes#persistentvolumeclaims)`,
				},
				resource.Attribute{
					Name:        "wait_until_bound",
					Description: `(Optional) Whether to wait for the claim to reach ` + "`" + `Bound` + "`" + ` state (to find volume in which to claim the space) ## Nested Blocks ### ` + "`" + `metadata` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "annotations",
					Description: `(Optional) An unstructured key value map stored with the persistent volume claim that may be used to store arbitrary metadata. ~> By default, the provider ignores any annotations whose key names end with`,
				},
				resource.Attribute{
					Name:        "generate_name",
					Description: `(Optional) Prefix, used by the server, to generate a unique name ONLY IF the ` + "`" + `name` + "`" + ` field has not been provided. This value will also be combined with a unique suffix. For more info see [Kubernetes reference](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#idempotency)`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) Map of string keys and values that can be used to organize and categorize (scope and select) the persistent volume claim. May match selectors of replication controllers and services. ~> By default, the provider ignores any labels whose key names end with`,
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
					Description: `The unique in time and space value for this persistent volume claim. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/identifiers#uids) ### ` + "`" + `spec` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "access_modes",
					Description: `(Required) A set of the desired access modes the volume should have. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/persistent-volumes#access-modes-1)`,
				},
				resource.Attribute{
					Name:        "resources",
					Description: `(Required) A list of the minimum resources the volume should have. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/persistent-volumes#resources)`,
				},
				resource.Attribute{
					Name:        "selector",
					Description: `(Optional) A label query over volumes to consider for binding.`,
				},
				resource.Attribute{
					Name:        "volume_name",
					Description: `(Optional) The binding reference to the PersistentVolume backing this claim.`,
				},
				resource.Attribute{
					Name:        "storage_class_name",
					Description: `(Optional) Name of the storage class requested by the claim ### ` + "`" + `match_expressions` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Optional) The label key that the selector applies to.`,
				},
				resource.Attribute{
					Name:        "operator",
					Description: `(Optional) A key's relationship to a set of values. Valid operators ard ` + "`" + `In` + "`" + `, ` + "`" + `NotIn` + "`" + `, ` + "`" + `Exists` + "`" + ` and ` + "`" + `DoesNotExist` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Optional) An array of string values. If the operator is ` + "`" + `In` + "`" + ` or ` + "`" + `NotIn` + "`" + `, the values array must be non-empty. If the operator is ` + "`" + `Exists` + "`" + ` or ` + "`" + `DoesNotExist` + "`" + `, the values array must be empty. This array is replaced during a strategic merge patch. ### ` + "`" + `resources` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "limits",
					Description: `(Optional) Map describing the maximum amount of compute resources allowed. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/compute-resources)/`,
				},
				resource.Attribute{
					Name:        "requests",
					Description: `(Optional) Map describing the minimum amount of compute resources required. If this is omitted for a container, it defaults to ` + "`" + `limits` + "`" + ` if that is explicitly specified, otherwise to an implementation-defined value. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/compute-resources)/ ### ` + "`" + `selector` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "match_expressions",
					Description: `(Optional) A list of label selector requirements. The requirements are ANDed.`,
				},
				resource.Attribute{
					Name:        "match_labels",
					Description: `(Optional) A map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of ` + "`" + `match_expressions` + "`" + `, whose key field is "key", the operator is "In", and the values array contains only "value". The requirements are ANDed. ## Import Persistent Volume Claim can be imported using its namespace and name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import kubernetes_persistent_volume_claim.example default/example-name ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "kubernetes_pod",
			Category:         "Resources",
			ShortDescription: `A pod is a group of one or more containers, the shared storage for those containers, and options about how to run the containers. Pods are always co-located and co-scheduled, and run in a shared context.`,
			Description:      ``,
			Keywords: []string{
				"pod",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "metadata",
					Description: `(Required) Standard pod's metadata. For more info see [Kubernetes reference](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata)`,
				},
				resource.Attribute{
					Name:        "spec",
					Description: `(Required) Spec of the pod owned by the cluster ## Nested Blocks ### ` + "`" + `metadata` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "annotations",
					Description: `(Optional) An unstructured key value map stored with the pod that may be used to store arbitrary metadata. ~> By default, the provider ignores any annotations whose key names end with`,
				},
				resource.Attribute{
					Name:        "generate_name",
					Description: `(Optional) Prefix, used by the server, to generate a unique name ONLY IF the ` + "`" + `name` + "`" + ` field has not been provided. This value will also be combined with a unique suffix. For more info see [Kubernetes reference](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#idempotency)`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) Map of string keys and values that can be used to organize and categorize (scope and select) the pod. May match selectors of replication controllers and services. ~> By default, the provider ignores any labels whose key names end with`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the pod, must be unique. Cannot be updated. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/identifiers#names)`,
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
					Description: `The unique in time and space value for this pod. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/identifiers#uids) ### ` + "`" + `spec` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "affinity",
					Description: `(Optional) A group of affinity scheduling rules. If specified, the pod will be dispatched by specified scheduler. If not specified, the pod will be dispatched by default scheduler.`,
				},
				resource.Attribute{
					Name:        "active_deadline_seconds",
					Description: `(Optional) Optional duration in seconds the pod may be active on the node relative to StartTime before the system will actively try to mark it failed and kill associated containers. Value must be a positive integer.`,
				},
				resource.Attribute{
					Name:        "automount_service_account_token",
					Description: `(Optional) Indicates whether a service account token should be automatically mounted. Defaults to ` + "`" + `true` + "`" + ` for Pods.`,
				},
				resource.Attribute{
					Name:        "container",
					Description: `(Optional) List of containers belonging to the pod. Containers cannot currently be added or removed. There must be at least one container in a Pod. Cannot be updated. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/containers)`,
				},
				resource.Attribute{
					Name:        "init_container",
					Description: `(Optional) List of init containers belonging to the pod. Init containers always run to completion and each must complete successfully before the next is started. For more info see [Kubernetes reference](https://kubernetes.io/docs/concepts/workloads/pods/init-containers)/`,
				},
				resource.Attribute{
					Name:        "dns_policy",
					Description: `(Optional) Set DNS policy for containers within the pod. Valid values are 'ClusterFirstWithHostNet', 'ClusterFirst', 'Default' or 'None'. DNS parameters given in DNSConfig will be merged with the policy selected with DNSPolicy. To have DNS options set along with hostNetwork, you have to specify DNS policy explicitly to 'ClusterFirstWithHostNet'. Optional: Defaults to 'ClusterFirst', see [Kubernetes reference](https://kubernetes.io/docs/concepts/services-networking/dns-pod-service/#pod-s-dns-policy).`,
				},
				resource.Attribute{
					Name:        "dns_config",
					Description: `(Optional) Specifies the DNS parameters of a pod. Parameters specified here will be merged to the generated DNS configuration based on DNSPolicy. Defaults to empty. See ` + "`" + `dns_config` + "`" + ` block definition below.`,
				},
				resource.Attribute{
					Name:        "enable_service_links",
					Description: `(Optional) Enables generating environment variables for service discovery. Optional: Defaults to true. For more info see [Kubernetes reference](https://kubernetes.io/docs/concepts/services-networking/connect-applications-service/#accessing-the-service).`,
				},
				resource.Attribute{
					Name:        "host_aliases",
					Description: `(Optional) List of hosts and IPs that will be injected into the pod's hosts file if specified. Optional: Defaults to empty. See ` + "`" + `host_aliases` + "`" + ` block definition below.`,
				},
				resource.Attribute{
					Name:        "host_ipc",
					Description: `(Optional) Use the host's ipc namespace. Optional: Defaults to false.`,
				},
				resource.Attribute{
					Name:        "host_network",
					Description: `(Optional) Host networking requested for this pod. Use the host's network namespace. If this option is set, the ports that will be used must be specified.`,
				},
				resource.Attribute{
					Name:        "host_pid",
					Description: `(Optional) Use the host's pid namespace.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `(Optional) Specifies the hostname of the Pod If not specified, the pod's hostname will be set to a system-defined value.`,
				},
				resource.Attribute{
					Name:        "image_pull_secrets",
					Description: `(Optional) ImagePullSecrets is an optional list of references to secrets in the same namespace to use for pulling any of the images used by this PodSpec. If specified, these secrets will be passed to individual puller implementations for them to use. For example, in the case of docker, only DockerConfig type secrets are honored. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/images#specifying-imagepullsecrets-on-a-pod)`,
				},
				resource.Attribute{
					Name:        "node_name",
					Description: `(Optional) NodeName is a request to schedule this pod onto a specific node. If it is non-empty, the scheduler simply schedules this pod onto that node, assuming that it fits resource requirements.`,
				},
				resource.Attribute{
					Name:        "node_selector",
					Description: `(Optional) NodeSelector is a selector which must be true for the pod to fit on a node. Selector which must match a node's labels for the pod to be scheduled on that node. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/node-selection).`,
				},
				resource.Attribute{
					Name:        "priority_class_name",
					Description: `(Optional) If specified, indicates the pod's priority. 'system-node-critical' and 'system-cluster-critical' are two special keywords which indicate the highest priorities with the formerer being the highest priority. Any other name must be defined by creating a PriorityClass object with that name. If not specified, the pod priority will be default or zero if there is no default.`,
				},
				resource.Attribute{
					Name:        "restart_policy",
					Description: `(Optional) Restart policy for all containers within the pod. One of Always, OnFailure, Never. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/pod-states#restartpolicy).`,
				},
				resource.Attribute{
					Name:        "security_context",
					Description: `(Optional) SecurityContext holds pod-level security attributes and common container settings. Optional: Defaults to empty`,
				},
				resource.Attribute{
					Name:        "service_account_name",
					Description: `(Optional) ServiceAccountName is the name of the ServiceAccount to use to run this pod. For more info see https://kubernetes.io/docs/tasks/configure-pod-container/configure-service-account/.`,
				},
				resource.Attribute{
					Name:        "share_process_namespace",
					Description: `(Optional) Share a single process namespace between all of the containers in a pod. When this is set containers will be able to view and signal processes from other containers in the same pod, and the first process in each container will not be assigned PID 1. HostPID and ShareProcessNamespace cannot both be set.`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `(Optional) If specified, the fully qualified Pod hostname will be "...svc.". If not specified, the pod will not have a domainname at all..`,
				},
				resource.Attribute{
					Name:        "termination_grace_period_seconds",
					Description: `(Optional) Optional duration in seconds the pod needs to terminate gracefully. May be decreased in delete request. Value must be non-negative integer. The value zero indicates delete immediately. If this value is nil, the default grace period will be used instead. The grace period is the duration in seconds after the processes running in the pod are sent a termination signal and the time when the processes are forcibly halted with a kill signal. Set this value longer than the expected cleanup time for your process.`,
				},
				resource.Attribute{
					Name:        "toleration",
					Description: `(Optional) Optional pod node tolerations. For more info see [Kubernetes reference](https://kubernetes.io/docs/concepts/configuration/taint-and-toleration/)`,
				},
				resource.Attribute{
					Name:        "topology_spread_constraint",
					Description: `(Optional) Describes how a group of pods ought to spread across topology domains. Scheduler will schedule pods in a way which abides by the constraints. For more info see [Kubernetes reference](https://kubernetes.io/docs/concepts/workloads/pods/pod-topology-spread-constraints/)`,
				},
				resource.Attribute{
					Name:        "volume",
					Description: `(Optional) List of volumes that can be mounted by containers belonging to the pod. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/volumes)`,
				},
				resource.Attribute{
					Name:        "readiness_gate",
					Description: `(Optional) If specified, all readiness gates will be evaluated for pod readiness. A pod is ready when all its containers are ready AND all conditions specified in the readiness gates have status equal to "True". [More info](https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/#pod-readiness-gate) ### ` + "`" + `affinity` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "node_affinity",
					Description: `(Optional) Node affinity scheduling rules for the pod. For more info see [Kubernetes reference](https://kubernetes.io/docs/concepts/configuration/assign-pod-node/#node-affinity-beta-feature)`,
				},
				resource.Attribute{
					Name:        "pod_affinity",
					Description: `(Optional) Inter-pod topological affinity. rules that specify that certain pods should be placed in the same topological domain (e.g. same node, same rack, same zone, same power domain, etc.) For more info see [Kubernetes reference](https://kubernetes.io/docs/concepts/configuration/assign-pod-node/#inter-pod-affinity-and-anti-affinity-beta-feature)`,
				},
				resource.Attribute{
					Name:        "pod_anti_affinity",
					Description: `(Optional) Inter-pod topological affinity. rules that specify that certain pods should be placed in the same topological domain (e.g. same node, same rack, same zone, same power domain, etc.) For more info see [Kubernetes reference](https://kubernetes.io/docs/concepts/configuration/assign-pod-node/#inter-pod-affinity-and-anti-affinity-beta-feature) ### ` + "`" + `node_affinity` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "required_during_scheduling_ignored_during_execution",
					Description: `(Optional) If the affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to an update), the system may or may not try to eventually evict the pod from its node.`,
				},
				resource.Attribute{
					Name:        "preferred_during_scheduling_ignored_during_execution",
					Description: `(Optional) The scheduler will prefer to schedule pods to nodes that satisfy the affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. ### ` + "`" + `required_during_scheduling_ignored_during_execution` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "node_selector_term",
					Description: `(Required) A list of node selector terms. The terms are ORed. ## ` + "`" + `node_selector_term` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "match_expressions",
					Description: `(Optional) A list of node selector requirements by node's labels.`,
				},
				resource.Attribute{
					Name:        "match_fields",
					Description: `(Optional) A list of node selector requirements by node's fields. ### ` + "`" + `match_expressions` + "`" + ` / ` + "`" + `match_fields` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) The label key that the selector applies to.`,
				},
				resource.Attribute{
					Name:        "operator",
					Description: `(Required) Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Optional) An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. ### ` + "`" + `preferred_during_scheduling_ignored_during_execution` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "preference",
					Description: `(Required) A node selector term, associated with the corresponding weight.`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `(Required) Weight associated with matching the corresponding nodeSelectorTerm, in the range 1-100. ### ` + "`" + `preference` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "match_expressions",
					Description: `(Optional) A list of node selector requirements by node's labels.`,
				},
				resource.Attribute{
					Name:        "match_fields",
					Description: `(Optional) A list of node selector requirements by node's fields. ## ` + "`" + `match_expressions` + "`" + ` / ` + "`" + `match_fields` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) The label key that the selector applies to.`,
				},
				resource.Attribute{
					Name:        "operator",
					Description: `(Required) Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Optional) An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. ### ` + "`" + `pod_affinity` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "required_during_scheduling_ignored_during_execution",
					Description: `(Optional) If the affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to a pod label update), the system may or may not try to eventually evict the pod from its node.`,
				},
				resource.Attribute{
					Name:        "preferred_during_scheduling_ignored_during_execution",
					Description: `(Optional) The scheduler will prefer to schedule pods to nodes that satisfy the affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. ### ` + "`" + `pod_anti_affinity` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "required_during_scheduling_ignored_during_execution",
					Description: `(Optional) If the anti-affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the anti-affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to a pod label update), the system may or may not try to eventually evict the pod from its node.`,
				},
				resource.Attribute{
					Name:        "preferred_during_scheduling_ignored_during_execution",
					Description: `(Optional) The scheduler will prefer to schedule pods to nodes that satisfy the anti-affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. ### ` + "`" + `required_during_scheduling_ignored_during_execution` + "`" + ` (pod_affinity_term) #### Arguments`,
				},
				resource.Attribute{
					Name:        "label_selector",
					Description: `(Optional) A label query over a set of resources, in this case pods.`,
				},
				resource.Attribute{
					Name:        "namespaces",
					Description: `(Optional) Specifies which namespaces the ` + "`" + `label_selector` + "`" + ` applies to (matches against). Null or empty list means "this pod's namespace"`,
				},
				resource.Attribute{
					Name:        "topology_key",
					Description: `(Optional) This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the ` + "`" + `label_selector` + "`" + ` in the specified namespaces, where co-located is defined as running on a node whose value of the label with key ` + "`" + `topology_key` + "`" + ` matches that of any node on which any of the selected pods is running. Empty ` + "`" + `topology_key` + "`" + ` is not allowed. ### ` + "`" + `preferred_during_scheduling_ignored_during_execution` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "pod_affinity_term",
					Description: `(Required) A pod affinity term, associated with the corresponding weight.`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `(Required) Weight associated with matching the corresponding ` + "`" + `pod_affinity_term` + "`" + `, in the range 1-100. ### ` + "`" + `container` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "args",
					Description: `(Optional) Arguments to the entrypoint. The docker image's CMD is used if this is not provided. Variable references $(VAR_NAME) are expanded using the container's environment. If a variable cannot be resolved, the reference in the input string will be unchanged. The $(VAR_NAME) syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped references will never be expanded, regardless of whether the variable exists or not. Cannot be updated. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/containers#containers-and-commands)`,
				},
				resource.Attribute{
					Name:        "command",
					Description: `(Optional) Entrypoint array. Not executed within a shell. The docker image's ENTRYPOINT is used if this is not provided. Variable references $(VAR_NAME) are expanded using the container's environment. If a variable cannot be resolved, the reference in the input string will be unchanged. The $(VAR_NAME) syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped references will never be expanded, regardless of whether the variable exists or not. Cannot be updated. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/containers#containers-and-commands)`,
				},
				resource.Attribute{
					Name:        "env",
					Description: `(Optional) Block of string name and value pairs to set in the container's environment. May be declared multiple times. Cannot be updated.`,
				},
				resource.Attribute{
					Name:        "env_from",
					Description: `(Optional) List of sources to populate environment variables in the container. The keys defined within a source must be a C_IDENTIFIER. All invalid keys will be reported as an event when the container is starting. When a key exists in multiple sources, the value associated with the last source will take precedence. Values defined by an Env with a duplicate key will take precedence. Cannot be updated.`,
				},
				resource.Attribute{
					Name:        "image",
					Description: `(Optional) Docker image name. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/images)`,
				},
				resource.Attribute{
					Name:        "image_pull_policy",
					Description: `(Optional) Image pull policy. One of Always, Never, IfNotPresent. Defaults to Always if :latest tag is specified, or IfNotPresent otherwise. Cannot be updated. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/images#updating-images)`,
				},
				resource.Attribute{
					Name:        "lifecycle",
					Description: `(Optional) Actions that the management system should take in response to container lifecycle events`,
				},
				resource.Attribute{
					Name:        "liveness_probe",
					Description: `(Optional) Periodic probe of container liveness. Container will be restarted if the probe fails. Cannot be updated. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/pod-states#container-probes)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the container specified as a DNS_LABEL. Each container in a pod must have a unique name (DNS_LABEL). Cannot be updated.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional) Block(s) of [port](#port)s to expose on the container. Exposing a port here gives the system additional information about the network connections a container uses, but is primarily informational. Not specifying a port here DOES NOT prevent that port from being exposed. Any port which is listening on the default "0.0.0.0" address inside a container will be accessible from the network. May be used multiple times. Cannot be updated.`,
				},
				resource.Attribute{
					Name:        "readiness_probe",
					Description: `(Optional) Periodic probe of container service readiness. Container will be removed from service endpoints if the probe fails. Cannot be updated. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/pod-states#container-probes)`,
				},
				resource.Attribute{
					Name:        "resources",
					Description: `(Optional) Compute Resources required by this container. Cannot be updated. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/persistent-volumes#resources)`,
				},
				resource.Attribute{
					Name:        "security_context",
					Description: `(Optional) Security options the pod should run with. For more info see https://kubernetes.io/docs/tasks/configure-pod-container/security-context/.`,
				},
				resource.Attribute{
					Name:        "startup_probe",
					Description: `(Optional) StartupProbe indicates that the Pod has successfully initialized. If specified, no other probes are executed until this completes successfully. If this probe fails, the Pod will be restarted, just as if the livenessProbe failed. This can be used to provide different probe parameters at the beginning of a Pod's lifecycle, when it might take a long time to load data or warm a cache, than during steady-state operation. This cannot be updated. For more info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes`,
				},
				resource.Attribute{
					Name:        "stdin",
					Description: `(Optional) Whether this container should allocate a buffer for stdin in the container runtime. If this is not set, reads from stdin in the container will always result in EOF.`,
				},
				resource.Attribute{
					Name:        "stdin_once",
					Description: `(Optional) Whether the container runtime should close the stdin channel after it has been opened by a single attach. When stdin is true the stdin stream will remain open across multiple attach sessions. If stdinOnce is set to true, stdin is opened on container start, is empty until the first client attaches to stdin, and then remains open and accepts data until the client disconnects, at which time stdin is closed and remains closed until the container is restarted. If this flag is false, a container processes that reads from stdin will never receive an EOF.`,
				},
				resource.Attribute{
					Name:        "termination_message_path",
					Description: `(Optional) Optional: Path at which the file to which the container's termination message will be written is mounted into the container's filesystem. Message written is intended to be brief final status, such as an assertion failure message. Defaults to /dev/termination-log. Cannot be updated.`,
				},
				resource.Attribute{
					Name:        "termination_message_policy",
					Description: `(Optional): Indicate how the termination message should be populated. File will use the contents of terminationMessagePath to populate the container status message on both success and failure. FallbackToLogsOnError will use the last chunk of container log output if the termination message file is empty and the container exited with an error. The log output is limited to 2048 bytes or 80 lines, whichever is smaller. Defaults to File. Cannot be updated.`,
				},
				resource.Attribute{
					Name:        "tty",
					Description: `(Optional) Whether this container should allocate a TTY for itself`,
				},
				resource.Attribute{
					Name:        "volume_mount",
					Description: `(Optional) Pod volumes to mount into the container's filesystem. Cannot be updated.`,
				},
				resource.Attribute{
					Name:        "working_dir",
					Description: `(Optional) Container's working directory. If not specified, the container runtime's default will be used, which might be configured in the container image. Cannot be updated. ### ` + "`" + `aws_elastic_block_store` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "fs_type",
					Description: `(Optional) Filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/volumes#awselasticblockstore)`,
				},
				resource.Attribute{
					Name:        "partition",
					Description: `(Optional) The partition in the volume that you want to mount. If omitted, the default is to mount by volume name. Examples: For volume /dev/sda1, you specify the partition as "1". Similarly, the volume partition for /dev/sda is "0" (or you can leave the property empty).`,
				},
				resource.Attribute{
					Name:        "read_only",
					Description: `(Optional) Whether to set the read-only property in VolumeMounts to "true". If omitted, the default is "false". For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/volumes#awselasticblockstore)`,
				},
				resource.Attribute{
					Name:        "volume_id",
					Description: `(Required) Unique ID of the persistent disk resource in AWS (Amazon EBS volume). For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/volumes#awselasticblockstore) ### ` + "`" + `azure_disk` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "caching_mode",
					Description: `(Required) Host Caching mode: None, Read Only, Read Write.`,
				},
				resource.Attribute{
					Name:        "data_disk_uri",
					Description: `(Required) The URI the data disk in the blob storage`,
				},
				resource.Attribute{
					Name:        "disk_name",
					Description: `(Required) The Name of the data disk in the blob storage`,
				},
				resource.Attribute{
					Name:        "fs_type",
					Description: `(Optional) Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.`,
				},
				resource.Attribute{
					Name:        "read_only",
					Description: `(Optional) Whether to force the read-only setting in VolumeMounts. Defaults to false (read/write). ### ` + "`" + `azure_file` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "read_only",
					Description: `(Optional) Whether to force the read-only setting in VolumeMounts. Defaults to false (read/write).`,
				},
				resource.Attribute{
					Name:        "secret_name",
					Description: `(Required) The name of secret that contains Azure Storage Account Name and Key`,
				},
				resource.Attribute{
					Name:        "share_name",
					Description: `(Required) Share Name ### ` + "`" + `capabilities` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "add",
					Description: `(Optional) Added capabilities`,
				},
				resource.Attribute{
					Name:        "drop",
					Description: `(Optional) Removed capabilities ### ` + "`" + `ceph_fs` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "monitors",
					Description: `(Required) Monitors is a collection of Ceph monitors. For more info see https://github.com/kubernetes/examples/tree/master/volumes/cephfs/#how-to-use-it.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Optional) Used as the mounted root, rather than the full Ceph tree, default is /.`,
				},
				resource.Attribute{
					Name:        "read_only",
					Description: `(Optional) Whether to force the read-only setting in VolumeMounts. Defaults to ` + "`" + `false` + "`" + ` (read/write). For more info see https://github.com/kubernetes/examples/tree/master/volumes/cephfs/#how-to-use-it.`,
				},
				resource.Attribute{
					Name:        "secret_file",
					Description: `(Optional) The path to key ring for User, default is /etc/ceph/user.secret. For more info see https://github.com/kubernetes/examples/tree/master/volumes/cephfs/#how-to-use-it.`,
				},
				resource.Attribute{
					Name:        "secret_ref",
					Description: `(Optional) Reference to the authentication secret for User, default is empty. For more info see https://github.com/kubernetes/examples/tree/master/volumes/cephfs/#how-to-use-it.`,
				},
				resource.Attribute{
					Name:        "user",
					Description: `(Optional) User is the rados user name, default is admin. For more info see https://github.com/kubernetes/examples/tree/master/volumes/cephfs/#how-to-use-it. ### ` + "`" + `cinder` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "fs_type",
					Description: `(Optional) Filesystem type to mount. Must be a filesystem type supported by the host operating system. Examples: "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified. For more info see https://github.com/kubernetes/examples/blob/master/mysql-cinder-pd/README.md#mysql-installation-with-cinder-volume-plugin.`,
				},
				resource.Attribute{
					Name:        "read_only",
					Description: `(Optional) Whether to force the read-only setting in VolumeMounts. Defaults to false (read/write). For more info see https://github.com/kubernetes/examples/blob/master/mysql-cinder-pd/README.md#mysql-installation-with-cinder-volume-plugin.`,
				},
				resource.Attribute{
					Name:        "volume_id",
					Description: `(Required) Volume ID used to identify the volume in Cinder. For more info see https://github.com/kubernetes/examples/blob/master/mysql-cinder-pd/README.md#mysql-installation-with-cinder-volume-plugin. ### ` + "`" + `config_map` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "default_mode",
					Description: `(Optional) Optional: mode bits to use on created files by default. Must be a value between 0 and 0777. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.`,
				},
				resource.Attribute{
					Name:        "items",
					Description: `(Optional) If unspecified, each key-value pair in the Data field of the referenced ConfigMap will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the ConfigMap, the volume setup will error unless it is marked ` + "`" + `optional` + "`" + `. Paths must be relative and may not contain the '..' path or start with '..'.`,
				},
				resource.Attribute{
					Name:        "optional",
					Description: `(Optional) Specify whether the ConfigMap or its keys must be defined.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the referent. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/identifiers#names) ### ` + "`" + `config_map_ref` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the referent. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/identifiers#names)`,
				},
				resource.Attribute{
					Name:        "optional",
					Description: `(Optional) Specify whether the ConfigMap must be defined ### ` + "`" + `config_map_key_ref` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Optional) The key to select.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the referent. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/identifiers#names)`,
				},
				resource.Attribute{
					Name:        "optional",
					Description: `(Optional) Specify whether the Secret or its key must be defined ### ` + "`" + `dns_config` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "nameservers",
					Description: `(Optional) A list of DNS name server IP addresses specified as strings. This will be appended to the base nameservers generated from DNSPolicy. Duplicated nameservers will be removed. Optional: Defaults to empty.`,
				},
				resource.Attribute{
					Name:        "option",
					Description: `(Optional) A list of DNS resolver options specified as blocks with ` + "`" + `name` + "`" + `/` + "`" + `value` + "`" + ` pairs. This will be merged with the base options generated from DNSPolicy. Duplicated entries will be removed. Resolution options given in Options will override those that appear in the base DNSPolicy. Optional: Defaults to empty.`,
				},
				resource.Attribute{
					Name:        "searches",
					Description: `(Optional) A list of DNS search domains for host-name lookup specified as strings. This will be appended to the base search paths generated from DNSPolicy. Duplicated search paths will be removed. Optional: Defaults to empty. The ` + "`" + `option` + "`" + ` block supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the option.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Optional) Value of the option. Optional: Defaults to empty. ### ` + "`" + `downward_api` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "default_mode",
					Description: `(Optional) Optional: mode bits to use on created files by default. Must be a value between 0 and 0777. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.`,
				},
				resource.Attribute{
					Name:        "items",
					Description: `(Optional) If unspecified, each key-value pair in the Data field of the referenced ConfigMap will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the ConfigMap, the volume setup will error. Paths must be relative and may not contain the '..' path or start with '..'. ### ` + "`" + `empty_dir` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "medium",
					Description: `(Optional) What type of storage medium should back this directory. The default is "" which means to use the node's default medium. Must be an empty string (default) or Memory. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/volumes#emptydir)`,
				},
				resource.Attribute{
					Name:        "size_limit",
					Description: `(Optional) Total amount of local storage required for this EmptyDir volume. For more info see [Kubernetes reference](https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/#resource-units-in-kubernetes) and [Kubernetes Quantity type](https://pkg.go.dev/k8s.io/apimachinery/pkg/api/resource?tab=doc#Quantity). ### ` + "`" + `env` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the environment variable. Must be a C_IDENTIFIER`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Optional) Variable references $(VAR_NAME) are expanded using the previous defined environment variables in the container and any service environment variables. If a variable cannot be resolved, the reference in the input string will be unchanged. The $(VAR_NAME) syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped references will never be expanded, regardless of whether the variable exists or not. Defaults to "".`,
				},
				resource.Attribute{
					Name:        "value_from",
					Description: `(Optional) Source for the environment variable's value ### ` + "`" + `env_from` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "config_map_ref",
					Description: `(Optional) The ConfigMap to select from`,
				},
				resource.Attribute{
					Name:        "prefix",
					Description: `(Optional) An optional identifer to prepend to each key in the ConfigMap. Must be a C_IDENTIFIER..`,
				},
				resource.Attribute{
					Name:        "secret_ref",
					Description: `(Optional) The Secret to select from ### ` + "`" + `exec` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "command",
					Description: `(Optional) Command is the command line to execute inside the container, the working directory for the command is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy. ### ` + "`" + `fc` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "fs_type",
					Description: `(Optional) Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.`,
				},
				resource.Attribute{
					Name:        "lun",
					Description: `(Required) FC target lun number`,
				},
				resource.Attribute{
					Name:        "read_only",
					Description: `(Optional) Whether to force the read-only setting in VolumeMounts. Defaults to false (read/write).`,
				},
				resource.Attribute{
					Name:        "target_ww_ns",
					Description: `(Required) FC target worldwide names (WWNs) ### ` + "`" + `field_ref` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "api_version",
					Description: `(Optional) Version of the schema the FieldPath is written in terms of, defaults to "v1".`,
				},
				resource.Attribute{
					Name:        "field_path",
					Description: `(Optional) Path of the field to select in the specified API version ### ` + "`" + `flex_volume` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "driver",
					Description: `(Required) Driver is the name of the driver to use for this volume.`,
				},
				resource.Attribute{
					Name:        "fs_type",
					Description: `(Optional) Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs", "ntfs". The default filesystem depends on FlexVolume script.`,
				},
				resource.Attribute{
					Name:        "options",
					Description: `(Optional) Extra command options if any.`,
				},
				resource.Attribute{
					Name:        "read_only",
					Description: `(Optional) Whether to force the ReadOnly setting in VolumeMounts. Defaults to false (read/write).`,
				},
				resource.Attribute{
					Name:        "secret_ref",
					Description: `(Optional) Reference to the secret object containing sensitive information to pass to the plugin scripts. This may be empty if no secret object is specified. If the secret object contains more than one secret, all secrets are passed to the plugin scripts. ### ` + "`" + `flocker` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "dataset_name",
					Description: `(Optional) Name of the dataset stored as metadata -> name on the dataset for Flocker should be considered as deprecated`,
				},
				resource.Attribute{
					Name:        "dataset_uuid",
					Description: `(Optional) UUID of the dataset. This is unique identifier of a Flocker dataset ### ` + "`" + `gce_persistent_disk` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "fs_type",
					Description: `(Optional) Filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/volumes#gcepersistentdisk)`,
				},
				resource.Attribute{
					Name:        "partition",
					Description: `(Optional) The partition in the volume that you want to mount. If omitted, the default is to mount by volume name. Examples: For volume /dev/sda1, you specify the partition as "1". Similarly, the volume partition for /dev/sda is "0" (or you can leave the property empty). For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/volumes#gcepersistentdisk)`,
				},
				resource.Attribute{
					Name:        "pd_name",
					Description: `(Required) Unique name of the PD resource in GCE. Used to identify the disk in GCE. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/volumes#gcepersistentdisk)`,
				},
				resource.Attribute{
					Name:        "read_only",
					Description: `(Optional) Whether to force the ReadOnly setting in VolumeMounts. Defaults to false. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/volumes#gcepersistentdisk) ### ` + "`" + `git_repo` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "directory",
					Description: `(Optional) Target directory name. Must not contain or start with '..'. If '.' is supplied, the volume directory will be the git repository. Otherwise, if specified, the volume will contain the git repository in the subdirectory with the given name.`,
				},
				resource.Attribute{
					Name:        "repository",
					Description: `(Optional) Repository URL`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `(Optional) Commit hash for the specified revision. ### ` + "`" + `glusterfs` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "endpoints_name",
					Description: `(Required) The endpoint name that details Glusterfs topology. For more info see https://github.com/kubernetes/examples/tree/master/volumes/glusterfs#create-a-pod.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Required) The Glusterfs volume path. For more info see https://github.com/kubernetes/examples/tree/master/volumes/glusterfs#create-a-pod.`,
				},
				resource.Attribute{
					Name:        "read_only",
					Description: `(Optional) Whether to force the Glusterfs volume to be mounted with read-only permissions. Defaults to false. For more info see https://github.com/kubernetes/examples/tree/master/volumes/glusterfs#create-a-pod. ### ` + "`" + `host_aliases` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "hostnames",
					Description: `(Required) Array of hostnames for the IP address.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `(Required) IP address of the host file entry. ### ` + "`" + `host_path` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Optional) Path of the directory on the host. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/volumes#hostpath)`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) Type for HostPath volume. Defaults to "". For more info see [Kubernetes reference](https://kubernetes.io/docs/concepts/storage/volumes#hostpath) ### ` + "`" + `http_get` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `(Optional) Host name to connect to, defaults to the pod IP. You probably want to set "Host" in httpHeaders instead.`,
				},
				resource.Attribute{
					Name:        "http_header",
					Description: `(Optional) Scheme to use for connecting to the host.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Optional) Path to access on the HTTP server.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional) Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.`,
				},
				resource.Attribute{
					Name:        "scheme",
					Description: `(Optional) Scheme to use for connecting to the host. ### ` + "`" + `http_header` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The header field name`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Optional) The header field value ### ` + "`" + `image_pull_secrets` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the referent. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/identifiers#names) ### ` + "`" + `iscsi` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "fs_type",
					Description: `(Optional) Filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/volumes#iscsi)`,
				},
				resource.Attribute{
					Name:        "iqn",
					Description: `(Required) Target iSCSI Qualified Name.`,
				},
				resource.Attribute{
					Name:        "iscsi_interface",
					Description: `(Optional) iSCSI interface name that uses an iSCSI transport. Defaults to 'default' (tcp).`,
				},
				resource.Attribute{
					Name:        "lun",
					Description: `(Optional) iSCSI target lun number.`,
				},
				resource.Attribute{
					Name:        "read_only",
					Description: `(Optional) Whether to force the read-only setting in VolumeMounts. Defaults to false.`,
				},
				resource.Attribute{
					Name:        "target_portal",
					Description: `(Required) iSCSI target portal. The portal is either an IP or ip_addr:port if the port is other than default (typically TCP ports 860 and 3260). ### ` + "`" + `items` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Optional) The key to project.`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `(Optional) Optional: mode bits to use on this file, must be a value between 0 and 0777. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Optional) The relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'. ### ` + "`" + `lifecycle` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "post_start",
					Description: `(Optional) post_start is called immediately after a container is created. If the handler fails, the container is terminated and restarted according to its restart policy. Other management of the container blocks until the hook completes. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/container-environment#hook-details)`,
				},
				resource.Attribute{
					Name:        "pre_stop",
					Description: `(Optional) pre_stop is called immediately before a container is terminated. The container is terminated after the handler completes. The reason for termination is passed to the handler. Regardless of the outcome of the handler, the container is eventually terminated. Other management of the container blocks until the hook completes. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/container-environment#hook-details) ### ` + "`" + `liveness_probe` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "exec",
					Description: `(Optional) exec specifies the action to take.`,
				},
				resource.Attribute{
					Name:        "failure_threshold",
					Description: `(Optional) Minimum consecutive failures for the probe to be considered failed after having succeeded.`,
				},
				resource.Attribute{
					Name:        "http_get",
					Description: `(Optional) Specifies the http request to perform.`,
				},
				resource.Attribute{
					Name:        "initial_delay_seconds",
					Description: `(Optional) Number of seconds after the container has started before liveness probes are initiated. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/pod-states#container-probes)`,
				},
				resource.Attribute{
					Name:        "period_seconds",
					Description: `(Optional) How often (in seconds) to perform the probe`,
				},
				resource.Attribute{
					Name:        "success_threshold",
					Description: `(Optional) Minimum consecutive successes for the probe to be considered successful after having failed.`,
				},
				resource.Attribute{
					Name:        "tcp_socket",
					Description: `(Optional) TCPSocket specifies an action involving a TCP port. TCP hooks not yet supported`,
				},
				resource.Attribute{
					Name:        "timeout_seconds",
					Description: `(Optional) Number of seconds after which the probe times out. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/pod-states#container-probes) ### ` + "`" + `nfs` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Required) Path that is exported by the NFS server. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/volumes#nfs)`,
				},
				resource.Attribute{
					Name:        "read_only",
					Description: `(Optional) Whether to force the NFS export to be mounted with read-only permissions. Defaults to false. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/volumes#nfs)`,
				},
				resource.Attribute{
					Name:        "server",
					Description: `(Required) Server is the hostname or IP address of the NFS server. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/volumes#nfs) ### ` + "`" + `persistent_volume_claim` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "claim_name",
					Description: `(Optional) ClaimName is the name of a PersistentVolumeClaim in the same`,
				},
				resource.Attribute{
					Name:        "read_only",
					Description: `(Optional) Will force the ReadOnly setting in VolumeMounts. ### ` + "`" + `photon_persistent_disk` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "fs_type",
					Description: `(Optional) Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.`,
				},
				resource.Attribute{
					Name:        "pd_id",
					Description: `(Required) ID that identifies Photon Controller persistent disk ### ` + "`" + `port` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "container_port",
					Description: `(Required) Number of port to expose on the pod's IP address. This must be a valid port number, 0 < x < 65536.`,
				},
				resource.Attribute{
					Name:        "host_ip",
					Description: `(Optional) What host IP to bind the external port to.`,
				},
				resource.Attribute{
					Name:        "host_port",
					Description: `(Optional) Number of port to expose on the host. If specified, this must be a valid port number, 0 < x < 65536. If HostNetwork is specified, this must match ContainerPort. Most containers do not need this.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) If specified, this must be an IANA_SVC_NAME and unique within the pod. Each named port in a pod must have a unique name. Name for the port that can be referred to by services`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Optional) Protocol for port. Must be UDP or TCP. Defaults to "TCP". ### ` + "`" + `post_start` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "exec",
					Description: `(Optional) exec specifies the action to take.`,
				},
				resource.Attribute{
					Name:        "http_get",
					Description: `(Optional) Specifies the http request to perform.`,
				},
				resource.Attribute{
					Name:        "tcp_socket",
					Description: `(Optional) TCPSocket specifies an action involving a TCP port. TCP hooks not yet supported ### ` + "`" + `pre_stop` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "exec",
					Description: `(Optional) exec specifies the action to take.`,
				},
				resource.Attribute{
					Name:        "http_get",
					Description: `(Optional) Specifies the http request to perform.`,
				},
				resource.Attribute{
					Name:        "tcp_socket",
					Description: `(Optional) TCPSocket specifies an action involving a TCP port. TCP hooks not yet supported ### ` + "`" + `quobyte` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "group",
					Description: `(Optional) Group to map volume access to Default is no group`,
				},
				resource.Attribute{
					Name:        "read_only",
					Description: `(Optional) Whether to force the Quobyte volume to be mounted with read-only permissions. Defaults to false.`,
				},
				resource.Attribute{
					Name:        "registry",
					Description: `(Required) Registry represents a single or multiple Quobyte Registry services specified as a string as host:port pair (multiple entries are separated with commas) which acts as the central registry for volumes`,
				},
				resource.Attribute{
					Name:        "user",
					Description: `(Optional) User to map volume access to Defaults to serivceaccount user`,
				},
				resource.Attribute{
					Name:        "volume",
					Description: `(Required) Volume is a string that references an already created Quobyte volume by name. ### ` + "`" + `rbd` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "ceph_monitors",
					Description: `(Required) A collection of Ceph monitors. For more info see https://github.com/kubernetes/examples/tree/master/volumes/rbd/#how-to-use-it.`,
				},
				resource.Attribute{
					Name:        "fs_type",
					Description: `(Optional) Filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/volumes#rbd)`,
				},
				resource.Attribute{
					Name:        "keyring",
					Description: `(Optional) Keyring is the path to key ring for RBDUser. Default is /etc/ceph/keyring. For more info see https://github.com/kubernetes/examples/tree/master/volumes/rbd/#how-to-use-it.`,
				},
				resource.Attribute{
					Name:        "rados_user",
					Description: `(Optional) The rados user name. Default is admin. For more info see https://github.com/kubernetes/examples/tree/master/volumes/rbd/#how-to-use-it.`,
				},
				resource.Attribute{
					Name:        "rbd_image",
					Description: `(Required) The rados image name. For more info see https://github.com/kubernetes/examples/tree/master/volumes/rbd/#how-to-use-it.`,
				},
				resource.Attribute{
					Name:        "rbd_pool",
					Description: `(Optional) The rados pool name. Default is rbd. For more info see https://github.com/kubernetes/examples/tree/master/volumes/rbd/#how-to-use-it.`,
				},
				resource.Attribute{
					Name:        "read_only",
					Description: `(Optional) Whether to force the read-only setting in VolumeMounts. Defaults to false. For more info see https://github.com/kubernetes/examples/tree/master/volumes/rbd/#how-to-use-it.`,
				},
				resource.Attribute{
					Name:        "secret_ref",
					Description: `(Optional) Name of the authentication secret for RBDUser. If provided overrides keyring. Default is nil. For more info see https://github.com/kubernetes/examples/tree/master/volumes/rbd/#how-to-use-it. ### ` + "`" + `readiness_probe` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "exec",
					Description: `(Optional) exec specifies the action to take.`,
				},
				resource.Attribute{
					Name:        "failure_threshold",
					Description: `(Optional) Minimum consecutive failures for the probe to be considered failed after having succeeded.`,
				},
				resource.Attribute{
					Name:        "http_get",
					Description: `(Optional) Specifies the http request to perform.`,
				},
				resource.Attribute{
					Name:        "initial_delay_seconds",
					Description: `(Optional) Number of seconds after the container has started before liveness probes are initiated. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/pod-states#container-probes)`,
				},
				resource.Attribute{
					Name:        "period_seconds",
					Description: `(Optional) How often (in seconds) to perform the probe`,
				},
				resource.Attribute{
					Name:        "success_threshold",
					Description: `(Optional) Minimum consecutive successes for the probe to be considered successful after having failed.`,
				},
				resource.Attribute{
					Name:        "tcp_socket",
					Description: `(Optional) TCPSocket specifies an action involving a TCP port. TCP hooks not yet supported`,
				},
				resource.Attribute{
					Name:        "timeout_seconds",
					Description: `(Optional) Number of seconds after which the probe times out. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/pod-states#container-probes) ### ` + "`" + `resources` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "limits",
					Description: `(Optional) Describes the maximum amount of compute resources allowed. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/compute-resources)/`,
				},
				resource.Attribute{
					Name:        "requests",
					Description: `(Optional) Describes the minimum amount of compute resources required. ### ` + "`" + `resource_field_ref` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "container_name",
					Description: `(Optional) The name of the container`,
				},
				resource.Attribute{
					Name:        "resource",
					Description: `(Required) Resource to select`,
				},
				resource.Attribute{
					Name:        "divisor",
					Description: `(Optional) Specifies the output format of the exposed resources, defaults to "1". ### ` + "`" + `se_linux_options` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "level",
					Description: `(Optional) Level is SELinux level label that applies to the container.`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `(Optional) Role is a SELinux role label that applies to the container.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) Type is a SELinux type label that applies to the container.`,
				},
				resource.Attribute{
					Name:        "user",
					Description: `(Optional) User is a SELinux user label that applies to the container. ### ` + "`" + `secret` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "default_mode",
					Description: `(Optional) Mode bits to use on created files by default. Must be a value between 0 and 0777. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.`,
				},
				resource.Attribute{
					Name:        "items",
					Description: `(Optional) List of Secret Items to project into the volume. See ` + "`" + `items` + "`" + ` block definition below. If unspecified, each key-value pair in the Data field of the referenced Secret will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the Secret, the volume setup will error unless it is marked ` + "`" + `optional` + "`" + `. Paths must be relative and may not contain the '..' path or start with '..'.`,
				},
				resource.Attribute{
					Name:        "optional",
					Description: `(Optional) Specify whether the Secret or its keys must be defined.`,
				},
				resource.Attribute{
					Name:        "secret_name",
					Description: `(Optional) Name of the secret in the pod's namespace to use. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/volumes#secrets) The ` + "`" + `items` + "`" + ` block supports the following:`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) The key to project.`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `(Optional) Mode bits to use on this file, must be a value between 0 and 0777. If not specified, the volume defaultMode will be used.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Required) The relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'. ### ` + "`" + `secret_ref` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the referent. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/identifiers#names)`,
				},
				resource.Attribute{
					Name:        "optional",
					Description: `(Optional) Specify whether the Secret must be defined ### ` + "`" + `secret_key_ref` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Optional) The key of the secret to select from. Must be a valid secret key.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the referent. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/identifiers#names)`,
				},
				resource.Attribute{
					Name:        "optional",
					Description: `(Optional) Specify whether the Secret or its key must be defined ### ` + "`" + `secret_ref` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the referent. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/identifiers#names) ### container ` + "`" + `security_context` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "allow_privilege_escalation",
					Description: `(Optional) AllowPrivilegeEscalation controls whether a process can gain more privileges than its parent process. This bool directly controls if the no_new_privs flag will be set on the container process. AllowPrivilegeEscalation is true always when the container is: 1) run as Privileged 2) has CAP_SYS_ADMIN`,
				},
				resource.Attribute{
					Name:        "capabilities",
					Description: `(Optional) The capabilities to add/drop when running containers. Defaults to the default set of capabilities granted by the container runtime.`,
				},
				resource.Attribute{
					Name:        "privileged",
					Description: `(Optional) Run container in privileged mode. Processes in privileged containers are essentially equivalent to root on the host. Defaults to false.`,
				},
				resource.Attribute{
					Name:        "read_only_root_filesystem",
					Description: `(Optional) Whether this container has a read-only root filesystem. Default is false.`,
				},
				resource.Attribute{
					Name:        "run_as_group",
					Description: `(Optional) The GID to run the entrypoint of the container process. Uses runtime default if unset. May also be set in PodSecurityContext. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.`,
				},
				resource.Attribute{
					Name:        "run_as_non_root",
					Description: `(Optional) Indicates that the container must run as a non-root user. If true, the Kubelet will validate the image at runtime to ensure that it does not run as UID 0 (root) and fail to start the container if it does. If unset or false, no such validation will be performed. May also be set in PodSecurityContext. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.`,
				},
				resource.Attribute{
					Name:        "run_as_user",
					Description: `(Optional) The UID to run the entrypoint of the container process. Defaults to user specified in image metadata if unspecified. May also be set in PodSecurityContext. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.`,
				},
				resource.Attribute{
					Name:        "se_linux_options",
					Description: `(Optional) The SELinux context to be applied to the container. If unspecified, the container runtime will allocate a random SELinux context for each container. May also be set in PodSecurityContext. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.`,
				},
				resource.Attribute{
					Name:        "sysctl",
					Description: `(Optional) holds a list of namespaced sysctls used for the pod. see [Sysctl](#sysctl) block. See [official docs](https://kubernetes.io/docs/tasks/administer-cluster/sysctl-cluster/) for more details. ##### Sysctl`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of a property to set.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) Value of a property to set. ### ` + "`" + `capabilities` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "add",
					Description: `(Optional) A list of added capabilities.`,
				},
				resource.Attribute{
					Name:        "drop",
					Description: `(Optional) A list of removed capabilities. ### pod ` + "`" + `security_context` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "fs_group",
					Description: `(Optional) A special supplemental group that applies to all containers in a pod. Some volume types allow the Kubelet to change the ownership of that volume to be owned by the pod: 1. The owning GID will be the FSGroup 2. The setgid bit is set (new files created in the volume will be owned by FSGroup) 3. The permission bits are OR'd with rw-rw---- If unset, the Kubelet will not modify the ownership and permissions of any volume.`,
				},
				resource.Attribute{
					Name:        "run_as_group",
					Description: `(Optional) The GID to run the entrypoint of the container process. Uses runtime default if unset. May also be set in SecurityContext. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence for that container.`,
				},
				resource.Attribute{
					Name:        "run_as_non_root",
					Description: `(Optional) Indicates that the container must run as a non-root user. If true, the Kubelet will validate the image at runtime to ensure that it does not run as UID 0 (root) and fail to start the container if it does. If unset or false, no such validation will be performed. May also be set in SecurityContext. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.`,
				},
				resource.Attribute{
					Name:        "run_as_user",
					Description: `(Optional) The UID to run the entrypoint of the container process. Defaults to user specified in image metadata if unspecified. May also be set in SecurityContext. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence for that container.`,
				},
				resource.Attribute{
					Name:        "se_linux_options",
					Description: `(Optional) The SELinux context to be applied to all containers. If unspecified, the container runtime will allocate a random SELinux context for each container. May also be set in SecurityContext. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence for that container.`,
				},
				resource.Attribute{
					Name:        "supplemental_groups",
					Description: `(Optional) A list of groups applied to the first process run in each container, in addition to the container's primary GID. If unspecified, no groups will be added to any container. ### ` + "`" + `tcp_socket` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Required) Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME. ### ` + "`" + `toleration` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "effect",
					Description: `(Optional) Effect indicates the taint effect to match. Empty means match all taint effects. When specified, allowed values are NoSchedule, PreferNoSchedule and NoExecute.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Optional) Key is the taint key that the toleration applies to. Empty means match all taint keys. If the key is empty, operator must be Exists; this combination means to match all values and all keys.`,
				},
				resource.Attribute{
					Name:        "operator",
					Description: `(Optional) Operator represents a key's relationship to the value. Valid operators are Exists and Equal. Defaults to Equal. Exists is equivalent to wildcard for value, so that a pod can tolerate all taints of a particular category.`,
				},
				resource.Attribute{
					Name:        "toleration_seconds",
					Description: `(Optional) TolerationSeconds represents the period of time the toleration (which must be of effect NoExecute, otherwise this field is ignored) tolerates the taint. By default, it is not set, which means tolerate the taint forever (do not evict). Zero and negative values will be treated as 0 (evict immediately) by the system.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Optional) Value is the taint value the toleration matches to. If the operator is Exists, the value should be empty, otherwise just a regular string. ### ` + "`" + `topology_spread_constraint` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "max_skew",
					Description: `(Optional) Describes the degree to which pods may be unevenly distributed. Default value is ` + "`" + `1` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "topology_key",
					Description: `(Optional) The key of node labels. Nodes that have a label with this key and identical values are considered to be in the same topology.`,
				},
				resource.Attribute{
					Name:        "when_unsatisfiable",
					Description: `(Optional) Indicates how to deal with a pod if it doesn't satisfy the spread constraint. Valid values are ` + "`" + `DoNotSchedule` + "`" + ` and ` + "`" + `ScheduleAnyway` + "`" + `. Default value is ` + "`" + `DoNotSchedule` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "label_selector",
					Description: `(Optional) A label query over a set of resources, in this case pods. ### ` + "`" + `value_from` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "config_map_key_ref",
					Description: `(Optional) Selects a key of a ConfigMap.`,
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
					Description: `(Optional) Selects a key of a secret in the pod's namespace. ### ` + "`" + `projected` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "default_mode",
					Description: `(Optional) Mode bits used to set permissions on created files by default. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.`,
				},
				resource.Attribute{
					Name:        "sources",
					Description: `(Required) List of volume projection sources ### ` + "`" + `sources` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "config_map",
					Description: `(Optional) Adapts a ConfigMap into a projected volume. The contents of the target ConfigMap's Data field will be presented in a projected volume as files using the keys in the Data field as the file names, unless the items element is populated with specific mappings of keys to paths. Note that this is identical to a configmap volume source without the default mode.`,
				},
				resource.Attribute{
					Name:        "downward_api",
					Description: `(Optional) Represents downward API info for projecting into a projected volume. Note that this is identical to a downward_api volume source without the default mode.`,
				},
				resource.Attribute{
					Name:        "secret",
					Description: `(Optional) Adapts a secret into a projected volume. The contents of the target Secret's Data field will be presented in a projected volume as files using the keys in the Data field as the file names. Note that this is identical to a secret volume source without the default mode.`,
				},
				resource.Attribute{
					Name:        "service_account_token",
					Description: `(Optional) Represents a projected service account token volume. This projection can be used to insert a service account token into the pods runtime filesystem for use against APIs (Kubernetes API Server or otherwise). ### ` + "`" + `service_account_token` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "audience",
					Description: `(Optional) Audience is the intended audience of the token. A recipient of a token must identify itself with an identifier specified in the audience of the token, and otherwise should reject the token. The audience defaults to the identifier of the apiserver.`,
				},
				resource.Attribute{
					Name:        "expiration_seconds",
					Description: `(Optional) The requested duration of validity of the service account token. As the token approaches expiration, the kubelet volume plugin will proactively rotate the service account token. The kubelet will start trying to rotate the token if the token is older than 80 percent of its time to live or if the token is older than 24 hours.Defaults to 1 hour and must be at least 10 minutes.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Required) Path is the path relative to the mount point of the file to project the token into. ### ` + "`" + `volume` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "aws_elastic_block_store",
					Description: `(Optional) Represents an AWS Disk resource that is attached to a kubelet's host machine and then exposed to the pod. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/volumes#awselasticblockstore)`,
				},
				resource.Attribute{
					Name:        "azure_disk",
					Description: `(Optional) Represents an Azure Data Disk mount on the host and bind mount to the pod.`,
				},
				resource.Attribute{
					Name:        "azure_file",
					Description: `(Optional) Represents an Azure File Service mount on the host and bind mount to the pod.`,
				},
				resource.Attribute{
					Name:        "ceph_fs",
					Description: `(Optional) Represents a Ceph FS mount on the host that shares a pod's lifetime`,
				},
				resource.Attribute{
					Name:        "cinder",
					Description: `(Optional) Represents a cinder volume attached and mounted on kubelets host machine. For more info see https://github.com/kubernetes/examples/blob/master/mysql-cinder-pd/README.md#mysql-installation-with-cinder-volume-plugin.`,
				},
				resource.Attribute{
					Name:        "config_map",
					Description: `(Optional) ConfigMap represents a configMap that should populate this volume`,
				},
				resource.Attribute{
					Name:        "downward_api",
					Description: `(Optional) DownwardAPI represents downward API about the pod that should populate this volume`,
				},
				resource.Attribute{
					Name:        "empty_dir",
					Description: `(Optional) EmptyDir represents a temporary directory that shares a pod's lifetime. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/volumes#emptydir)`,
				},
				resource.Attribute{
					Name:        "fc",
					Description: `(Optional) Represents a Fibre Channel resource that is attached to a kubelet's host machine and then exposed to the pod.`,
				},
				resource.Attribute{
					Name:        "flex_volume",
					Description: `(Optional) Represents a generic volume resource that is provisioned/attached using an exec based plugin. This is an alpha feature and may change in future.`,
				},
				resource.Attribute{
					Name:        "flocker",
					Description: `(Optional) Represents a Flocker volume attached to a kubelet's host machine and exposed to the pod for its usage. This depends on the Flocker control service being running`,
				},
				resource.Attribute{
					Name:        "gce_persistent_disk",
					Description: `(Optional) Represents a GCE Disk resource that is attached to a kubelet's host machine and then exposed to the pod. Provisioned by an admin. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/volumes#gcepersistentdisk)`,
				},
				resource.Attribute{
					Name:        "git_repo",
					Description: `(Optional) GitRepo represents a git repository at a particular revision.`,
				},
				resource.Attribute{
					Name:        "glusterfs",
					Description: `(Optional) Represents a Glusterfs volume that is attached to a host and exposed to the pod. Provisioned by an admin. For more info see https://github.com/kubernetes/examples/tree/master/volumes/glusterfs#glusterfs.`,
				},
				resource.Attribute{
					Name:        "host_path",
					Description: `(Optional) Represents a directory on the host. Provisioned by a developer or tester. This is useful for single-node development and testing only! On-host storage is not supported in any way and WILL NOT WORK in a multi-node cluster. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/volumes#hostpath)`,
				},
				resource.Attribute{
					Name:        "iscsi",
					Description: `(Optional) Represents an ISCSI Disk resource that is attached to a kubelet's host machine and then exposed to the pod. Provisioned by an admin.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Volume's name. Must be a DNS_LABEL and unique within the pod. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/identifiers#names)`,
				},
				resource.Attribute{
					Name:        "nfs",
					Description: `(Optional) Represents an NFS mount on the host. Provisioned by an admin. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/volumes#nfs)`,
				},
				resource.Attribute{
					Name:        "persistent_volume_claim",
					Description: `(Optional) The specification of a persistent volume.`,
				},
				resource.Attribute{
					Name:        "photon_persistent_disk",
					Description: `(Optional) Represents a PhotonController persistent disk attached and mounted on kubelets host machine`,
				},
				resource.Attribute{
					Name:        "quobyte",
					Description: `(Optional) Quobyte represents a Quobyte mount on the host that shares a pod's lifetime`,
				},
				resource.Attribute{
					Name:        "rbd",
					Description: `(Optional) Represents a Rados Block Device mount on the host that shares a pod's lifetime. For more info see https://kubernetes.io/docs/concepts/storage/volumes/#rbd.`,
				},
				resource.Attribute{
					Name:        "secret",
					Description: `(Optional) Secret represents a secret that should populate this volume. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/volumes#secrets)`,
				},
				resource.Attribute{
					Name:        "vsphere_volume",
					Description: `(Optional) Represents a vSphere volume attached and mounted on kubelets host machine ### ` + "`" + `volume_mount` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "mount_path",
					Description: `(Required) Path within the container at which the volume should be mounted. Must not contain ':'.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) This must match the Name of a Volume.`,
				},
				resource.Attribute{
					Name:        "read_only",
					Description: `(Optional) Mounted read-only if true, read-write otherwise (false or unspecified). Defaults to false.`,
				},
				resource.Attribute{
					Name:        "sub_path",
					Description: `(Optional) Path within the volume from which the container's volume should be mounted. Defaults to "" (volume's root).`,
				},
				resource.Attribute{
					Name:        "mount_propagation",
					Description: `(Optional) Mount propagation mode. Defaults to "None". For more info see [Kubernetes reference](https://kubernetes.io/docs/concepts/storage/volumes/#mount-propagation) ### ` + "`" + `vsphere_volume` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "fs_type",
					Description: `(Optional) Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.`,
				},
				resource.Attribute{
					Name:        "volume_path",
					Description: `(Required) Path that identifies vSphere volume vmdk ### ` + "`" + `readiness_gate` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "condition_type",
					Description: `(Required) refers to a condition in the pod's condition list with matching type. ## Timeouts The following [Timeout](/docs/configuration/resources.html#operation-timeouts) configuration options are available for the ` + "`" + `kubernetes_pod` + "`" + ` resource:`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Default ` + "`" + `5 minutes` + "`" + `) Used for Creating Pods.`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Default ` + "`" + `5 minutes` + "`" + `) Used for Destroying Pods. ## Import Pod can be imported using the namespace and name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import kubernetes_pod.example default/terraform-example ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "kubernetes_pod_disruption_budget",
			Category:         "Resources",
			ShortDescription: `A Pod Disruption Budget limits the number of pods of a replicated application that are down simultaneously from voluntary disruptions. For example, a quorum-based application would like to ensure that the number of replicas running is never brought below the number needed for a quorum. A web front end might want to ensure that the number of replicas serving load never falls below a certain percentage of the total.`,
			Description:      ``,
			Keywords: []string{
				"pod",
				"disruption",
				"budget",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "metadata",
					Description: `(Required) Standard resource's metadata. For more info: https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#spec-and-status`,
				},
				resource.Attribute{
					Name:        "spec",
					Description: `(Required) Spec defines the behavior of a Pod Disruption Budget. https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#spec-and-status ## Nested Blocks ### ` + "`" + `metadata` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "annotations",
					Description: `(Optional) An unstructured key value map stored with the resource that may be used to store arbitrary metadata. ~> By default, the provider ignores any annotations whose key names end with`,
				},
				resource.Attribute{
					Name:        "generate_name",
					Description: `(Optional) Prefix, used by the server, to generate a unique name ONLY IF the ` + "`" + `name` + "`" + ` field has not been provided. This value will also be combined with a unique suffix. Read more: https://github.com/kubernetes/community/blob/master/contributors/devel/api-conventions.md#idempotency`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) Map of string keys and values that can be used to organize and categorize (scope and select) the service. May match selectors of replication controllers and services. ~> By default, the provider ignores any labels whose key names end with`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the service, must be unique. Cannot be updated. For more info: http://kubernetes.io/docs/user-guide/identifiers#names`,
				},
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) Namespace defines the space within which name of the service must be unique. #### Attributes`,
				},
				resource.Attribute{
					Name:        "generation",
					Description: `A sequence number representing a specific generation of the desired state.`,
				},
				resource.Attribute{
					Name:        "resource_version",
					Description: `An opaque value that represents the internal version of this service that can be used by clients to determine when service has changed. Read more: https://github.com/kubernetes/community/blob/master/contributors/devel/api-conventions.md#concurrency-control-and-consistency`,
				},
				resource.Attribute{
					Name:        "uid",
					Description: `The unique in time and space value for this service. For more info: http://kubernetes.io/docs/user-guide/identifiers#uids ### ` + "`" + `spec` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "max_unavailable",
					Description: `(Optional) Specifies the number of pods from the selected set that can be unavailable after the eviction. It can be either an absolute number or a percentage. You can specify only one of max_unavailable and min_available in a single Pod Disruption Budget. max_unavailable can only be used to control the eviction of pods that have an associated controller managing them.`,
				},
				resource.Attribute{
					Name:        "min_available",
					Description: `(Optional) Specifies the number of pods from the selected set that must still be available after the eviction, even in the absence of the evicted pod. min_available can be either an absolute number or a percentage. You can specify only one of min_available and max_unavailable in a single Pod Disruption Budget. min_available can only be used to control the eviction of pods that have an associated controller managing them.`,
				},
				resource.Attribute{
					Name:        "selector",
					Description: `(Optional) A label query over controllers (Deployment, ReplicationController, ReplicaSet, or StatefulSet) that the Pod Disruption Budget should be applied to. For more info: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#label-selectors`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "kubernetes_pod_security_policy",
			Category:         "Resources",
			ShortDescription: `A Pod Security Policy is a cluster-level resource that controls security sensitive aspects of the pod specification.`,
			Description:      ``,
			Keywords: []string{
				"pod",
				"security",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "metadata",
					Description: `(Required) Standard Pod Security Policy's metadata. For more info see [Kubernetes reference](https://github.com/kubernetes/community/blob/e59e666e3464c7d4851136baa8835a311efdfb8e/contributors/devel/api-conventions.md#metadata)`,
				},
				resource.Attribute{
					Name:        "spec",
					Description: `(Required) Spec contains information for locating and communicating with a server. [Kubernetes reference](https://github.com/kubernetes/community/blob/e59e666e3464c7d4851136baa8835a311efdfb8e/contributors/devel/api-conventions.md#spec-and-status) ## Nested Blocks ### ` + "`" + `metadata` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "annotations",
					Description: `(Optional) An unstructured key value map stored with the Pod Security Policy that may be used to store arbitrary metadata. ~> By default, the provider ignores any annotations whose key names end with`,
				},
				resource.Attribute{
					Name:        "generate_name",
					Description: `(Optional) Prefix, used by the server, to generate a unique name ONLY IF the ` + "`" + `name` + "`" + ` field has not been provided. This value will also be combined with a unique suffix. For more info see [Kubernetes reference](https://github.com/kubernetes/community/blob/e59e666e3464c7d4851136baa8835a311efdfb8e/contributors/devel/api-conventions.md#idempotency)`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) Map of string keys and values that can be used to organize and categorize (scope and select) the Pod Security Policy. ~> By default, the provider ignores any labels whose key names end with`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the Pod Security Policy, must be unique. Cannot be updated. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/identifiers#names) #### Attributes`,
				},
				resource.Attribute{
					Name:        "generation",
					Description: `A sequence number representing a specific generation of the desired state.`,
				},
				resource.Attribute{
					Name:        "resource_version",
					Description: `An opaque value that represents the internal version of this Pod Security Policy that can be used by clients to determine when Pod Security Policy has changed. For more info see [Kubernetes reference](https://github.com/kubernetes/community/blob/e59e666e3464c7d4851136baa8835a311efdfb8e/contributors/devel/api-conventions.md#concurrency-control-and-consistency)`,
				},
				resource.Attribute{
					Name:        "uid",
					Description: `The unique in time and space value for this Pod Security Policy. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/identifiers#uids) ### ` + "`" + `spec` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "allow_privilege_escalation",
					Description: `(Optional) determines if a pod can request to allow privilege escalation. If unspecified, defaults to true.`,
				},
				resource.Attribute{
					Name:        "allowed_capabilities",
					Description: `(Optional) a list of capabilities that can be requested to add to the container. Capabilities in this field may be added at the pod author's discretion. You must not list a capability in both allowedCapabilities and requiredDropCapabilities.`,
				},
				resource.Attribute{
					Name:        "allowed_proc_mount_types",
					Description: `(Optional) a whitelist of allowed ProcMountTypes. Empty or nil indicates that only the DefaultProcMountType may be used. This requires the ProcMountType feature flag to be enabled. Possible values are ` + "`" + `"Default"` + "`" + ` or ` + "`" + `"Unmasked"` + "`" + ``,
				},
				resource.Attribute{
					Name:        "allowed_unsafe_sysctls",
					Description: `(Optional) a list of explicitly allowed unsafe sysctls, defaults to none. Each entry is either a plain sysctl name or ends in "`,
				},
				resource.Attribute{
					Name:        "default_add_capabilities",
					Description: `(Optional) the default set of capabilities that will be added to the container unless the pod spec specifically drops the capability. You may not list a capability in both defaultAddCapabilities and requiredDropCapabilities. Capabilities added here are implicitly allowed, and need not be included in the allowedCapabilities list.`,
				},
				resource.Attribute{
					Name:        "default_allow_privilege_escalation",
					Description: `(Optional) controls the default setting for whether a process can gain more privileges than its parent process.`,
				},
				resource.Attribute{
					Name:        "forbidden_sysctls",
					Description: `(Optional) forbiddenSysctls is a list of explicitly forbidden sysctls, defaults to none. Each entry is either a plain sysctl name or ends in "`,
				},
				resource.Attribute{
					Name:        "host_ipc",
					Description: `(Optional) determines if the policy allows the use of HostIPC in the pod spec.`,
				},
				resource.Attribute{
					Name:        "host_network",
					Description: `(Optional) determines if the policy allows the use of HostNetwork in the pod spec.`,
				},
				resource.Attribute{
					Name:        "host_pid",
					Description: `(Optional) determines if the policy allows the use of HostPID in the pod spec.`,
				},
				resource.Attribute{
					Name:        "host_ports",
					Description: `(Optional) determines which host port ranges are allowed to be exposed.`,
				},
				resource.Attribute{
					Name:        "privileged",
					Description: `(Optional) determines if a pod can request to be run as privileged.`,
				},
				resource.Attribute{
					Name:        "read_only_root_filesystem",
					Description: `(Optional) when set to true will force containers to run with a read only root file system. If the container specifically requests to run with a non-read only root file system the PSP should deny the pod. If set to false the container may run with a read only root file system if it wishes but it will not be forced to.`,
				},
				resource.Attribute{
					Name:        "required_drop_capabilities",
					Description: `(Optional) the capabilities that will be dropped from the container. These are required to be dropped and cannot be added.`,
				},
				resource.Attribute{
					Name:        "volumes",
					Description: `(Optional) a white list of allowed volume plugins. Empty indicates that no volumes may be used. To allow all volumes you may use '`,
				},
				resource.Attribute{
					Name:        "driver",
					Description: `(Required) the name of the Flexvolume driver. ### allowed_host_paths ### Arguments`,
				},
				resource.Attribute{
					Name:        "path_prefix",
					Description: `(Required) the path prefix that the host volume must match. It does not support ` + "`" + ``,
				},
				resource.Attribute{
					Name:        "read_only",
					Description: `(Optional) when set to true, will allow host volumes matching the pathPrefix only if all volume mounts are readOnly. ### ` + "`" + `fs_group` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "rule",
					Description: `(Required) the strategy that will dictate what FSGroup is used in the SecurityContext.`,
				},
				resource.Attribute{
					Name:        "rule",
					Description: `(Required) the strategy that will dictate the allowable RunAsUser values that may be set.`,
				},
				resource.Attribute{
					Name:        "rule",
					Description: `(Required) the strategy that will dictate the allowable RunAsGroup values that may be set.`,
				},
				resource.Attribute{
					Name:        "rule",
					Description: `(Required) the strategy that will dictate the allowable labels that may be set.`,
				},
				resource.Attribute{
					Name:        "se_linux_options",
					Description: `(Optional) required to run as; required for MustRunAs. For more info: https://kubernetes.io/docs/tasks/configure-pod-container/security-context/ ### ` + "`" + `supplemental_groups` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "rule",
					Description: `(Required) the strategy that will dictate what supplemental groups is used in the SecurityContext.`,
				},
				resource.Attribute{
					Name:        "min",
					Description: `(Required) the start of the range, inclusive.`,
				},
				resource.Attribute{
					Name:        "max",
					Description: `(Required) the end of the range, inclusive. ## Import Pod Security Policy can be imported using its name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import kubernetes_pod_security_policy.example terraform-example ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "kubernetes_priority_class",
			Category:         "Resources",
			ShortDescription: `A PriorityClass is a non-namespaced object that defines a mapping from a priority class name to the integer value of the priority.`,
			Description:      ``,
			Keywords: []string{
				"priority",
				"class",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "metadata",
					Description: `(Required) Standard resource quota's metadata. For more info see [Kubernetes reference](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata)`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required, Forces new resource) The value of this priority class. This is the actual priority that pods receive when they have the name of this class in their pod spec.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) An arbitrary string that usually provides guidelines on when this priority class should be used.`,
				},
				resource.Attribute{
					Name:        "global_default",
					Description: `(Optional) Boolean that specifies whether this PriorityClass should be considered as the default priority for pods that do not have any priority class. ## Nested Blocks ### ` + "`" + `metadata` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "annotations",
					Description: `(Optional) An unstructured key value map stored with the resource quota that may be used to store arbitrary metadata. ~> By default, the provider ignores any annotations whose key names end with`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) Map of string keys and values that can be used to organize and categorize (scope and select) the resource quota. May match selectors of replication controllers and services. ~> By default, the provider ignores any labels whose key names end with`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the resource quota, must be unique. Cannot be updated. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/identifiers#names) #### Attributes`,
				},
				resource.Attribute{
					Name:        "generation",
					Description: `A sequence number representing a specific generation of the desired state.`,
				},
				resource.Attribute{
					Name:        "resource_version",
					Description: `An opaque value that represents the internal version of this resource quota that can be used by clients to determine when resource quota has changed. For more info see [Kubernetes reference](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency)`,
				},
				resource.Attribute{
					Name:        "uid",
					Description: `The unique in time and space value for this resource quota. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/identifiers#uids) ## Import Priority Class can be imported using its name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import kubernetes_priority_class.example terraform-example ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "kubernetes_replication_controller",
			Category:         "Resources",
			ShortDescription: `A Replication Controller ensures that a specified number of pod “replicas” are running at any one time. In other words, a Replication Controller makes sure that a pod or homogeneous set of pods are always up and available. If there are too many pods, it will kill some. If there are too few, the Replication Controller will start more.`,
			Description:      ``,
			Keywords: []string{
				"replication",
				"controller",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "metadata",
					Description: `(Required) Standard replication controller's metadata. For more info see [Kubernetes reference](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata)`,
				},
				resource.Attribute{
					Name:        "spec",
					Description: `(Required) Spec defines the specification of the desired behavior of the replication controller. For more info see [Kubernetes reference](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#spec-and-status) ## Nested Blocks ### ` + "`" + `metadata` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "annotations",
					Description: `(Optional) An unstructured key value map stored with the replication controller that may be used to store arbitrary metadata. ~> By default, the provider ignores any annotations whose key names end with`,
				},
				resource.Attribute{
					Name:        "generate_name",
					Description: `(Optional) Prefix, used by the server, to generate a unique name ONLY IF the ` + "`" + `name` + "`" + ` field has not been provided. This value will also be combined with a unique suffix. For more info see [Kubernetes reference](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#idempotency)`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) Map of string keys and values that can be used to organize and categorize (scope and select) the replication controller. ~> By default, the provider ignores any labels whose key names end with`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the replication controller, must be unique. Cannot be updated. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/identifiers#names)`,
				},
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) Namespace defines the space within which name of the replication controller must be unique. #### Attributes`,
				},
				resource.Attribute{
					Name:        "generation",
					Description: `A sequence number representing a specific generation of the desired state.`,
				},
				resource.Attribute{
					Name:        "resource_version",
					Description: `An opaque value that represents the internal version of this replication controller that can be used by clients to determine when replication controller has changed. For more info see [Kubernetes reference](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency)`,
				},
				resource.Attribute{
					Name:        "uid",
					Description: `The unique in time and space value for this replication controller. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/identifiers#uids) ### ` + "`" + `spec` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "min_ready_seconds",
					Description: `(Optional) Minimum number of seconds for which a newly created pod should be ready without any of its container crashing, for it to be considered available. Defaults to 0 (pod will be considered available as soon as it is ready)`,
				},
				resource.Attribute{
					Name:        "replicas",
					Description: `(Optional) The number of desired replicas. Defaults to 1. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/replication-controller#what-is-a-replication-controller)`,
				},
				resource.Attribute{
					Name:        "selector",
					Description: `(Required) A label query over pods that should match the Replicas count. Label keys and values that must match in order to be controlled by this replication controller.`,
				},
				resource.Attribute{
					Name:        "template",
					Description: `(Required) Template is the object that describes the pod that will be created if insufficient replicas are detected. This takes precedence over a TemplateRef. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/replication-controller#pod-template) ## Nested Blocks ### ` + "`" + `spec.template` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `(Optional) Standard object's metadata. For more info see [Kubernetes reference](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata). While required by the kubernetes API, this field is marked as optional to allow the usage of the deprecated pod spec fields that were mistakenly placed directly under the ` + "`" + `template` + "`" + ` block.`,
				},
				resource.Attribute{
					Name:        "spec",
					Description: `(Optional) Specification of the desired behavior of the pod. For more info see [Kubernetes reference](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#spec-and-status) ~>`,
				},
				resource.Attribute{
					Name:        "annotations",
					Description: `(Optional) An unstructured key value map stored with the replication controller that may be used to store arbitrary metadata. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/annotations)`,
				},
				resource.Attribute{
					Name:        "generate_name",
					Description: `(Optional) Prefix, used by the server, to generate a unique name ONLY IF the ` + "`" + `name` + "`" + ` field has not been provided. This value will also be combined with a unique suffix. For more info see [Kubernetes reference](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#idempotency)`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) Map of string keys and values that can be used to organize and categorize (scope and select) the pods managed by this replication controller .`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the replication controller, must be unique. Cannot be updated. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/identifiers#names)`,
				},
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) Namespace defines the space within which name of the replication controller must be unique. ## Nested Blocks ### ` + "`" + `spec.template.spec` + "`" + ` #### Arguments These arguments are the same as the for the ` + "`" + `spec` + "`" + ` block of a Pod. Please see the [Pod resource](pod.html#spec-1) for reference. ## Timeouts The following [Timeout](/docs/configuration/resources.html#operation-timeouts) configuration options are available: - ` + "`" + `create` + "`" + ` - (Default ` + "`" + `10 minutes` + "`" + `) Used for creating new controller - ` + "`" + `update` + "`" + ` - (Default ` + "`" + `10 minutes` + "`" + `) Used for updating a controller - ` + "`" + `delete` + "`" + ` - (Default ` + "`" + `10 minutes` + "`" + `) Used for destroying a controller ## Import Replication Controller can be imported using the namespace and name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import kubernetes_replication_controller.example default/terraform-example ` + "`" + `` + "`" + `` + "`" + ` ~>`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "kubernetes_resource_quota",
			Category:         "Resources",
			ShortDescription: `A resource quota provides constraints that limit aggregate resource consumption per namespace. It can limit the quantity of objects that can be created in a namespace by type, as well as the total amount of compute resources that may be consumed by resources in that project.`,
			Description:      ``,
			Keywords: []string{
				"resource",
				"quota",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "metadata",
					Description: `(Required) Standard resource quota's metadata. For more info see [Kubernetes reference](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata)`,
				},
				resource.Attribute{
					Name:        "spec",
					Description: `(Optional) Spec defines the desired quota. [Kubernetes reference](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#spec-and-status) ## Nested Blocks ### ` + "`" + `metadata` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "annotations",
					Description: `(Optional) An unstructured key value map stored with the resource quota that may be used to store arbitrary metadata. ~> By default, the provider ignores any annotations whose key names end with`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) Map of string keys and values that can be used to organize and categorize (scope and select) the resource quota. May match selectors of replication controllers and services. ~> By default, the provider ignores any labels whose key names end with`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the resource quota, must be unique. Cannot be updated. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/identifiers#names)`,
				},
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) Namespace defines the space within which name of the resource quota must be unique. #### Attributes`,
				},
				resource.Attribute{
					Name:        "generation",
					Description: `A sequence number representing a specific generation of the desired state.`,
				},
				resource.Attribute{
					Name:        "resource_version",
					Description: `An opaque value that represents the internal version of this resource quota that can be used by clients to determine when resource quota has changed. For more info see [Kubernetes reference](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency)`,
				},
				resource.Attribute{
					Name:        "uid",
					Description: `The unique in time and space value for this resource quota. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/identifiers#uids) ### ` + "`" + `spec` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "hard",
					Description: `(Optional) The set of desired hard limits for each named resource. For more info see [Kubernetes reference](https://kubernetes.io/docs/concepts/policy/resource-quotas)`,
				},
				resource.Attribute{
					Name:        "scopes",
					Description: `(Optional) A collection of filters that must match each object tracked by a quota. If not specified, the quota matches all objects.`,
				},
				resource.Attribute{
					Name:        "scope_selector",
					Description: `(Optional) A collection of filters like scopes that must match each object tracked by a quota but expressed using ScopeSelectorOperator in combination with possible values. See ` + "`" + `scope_selector` + "`" + ` below for more details. #### ` + "`" + `scope_selector` + "`" + ` ##### Arguments`,
				},
				resource.Attribute{
					Name:        "match_expression",
					Description: `(Optional) A list of scope selector requirements by scope of the resources. See ` + "`" + `match_expression` + "`" + ` below for more details. ##### ` + "`" + `match_expression` + "`" + ` ###### Arguments`,
				},
				resource.Attribute{
					Name:        "scope_name",
					Description: `(Required) The name of the scope that the selector applies to. Valid values are ` + "`" + `Terminating` + "`" + `, ` + "`" + `NotTerminating` + "`" + `, ` + "`" + `BestEffort` + "`" + `, ` + "`" + `NotBestEffort` + "`" + `, and ` + "`" + `PriorityClass` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "operator",
					Description: `(Required) Represents a scope's relationship to a set of values. Valid operators are ` + "`" + `In` + "`" + `, ` + "`" + `NotIn` + "`" + `, ` + "`" + `Exists` + "`" + `, ` + "`" + `DoesNotExist` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Optional) A list of scope selector requirements by scope of the resources. ## Import Resource Quota can be imported using its namespace and name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import kubernetes_resource_quota.example default/terraform-example ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "kubernetes_role",
			Category:         "Resources",
			ShortDescription: `A role contains rules that represent a set of permissions. Permissions are purely additive (there are no “deny” rules).`,
			Description:      ``,
			Keywords: []string{
				"role",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "metadata",
					Description: `(Required) Standard role's metadata. For more info see [Kubernetes reference](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata)`,
				},
				resource.Attribute{
					Name:        "rule",
					Description: `(Required) List of rules that define the set of permissions for this role. For more info see [Kubernetes reference](https://kubernetes.io/docs/reference/access-authn-authz/rbac/) ## Nested Blocks ### ` + "`" + `metadata` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "annotations",
					Description: `(Optional) An unstructured key value map stored with the role that may be used to store arbitrary metadata. ~> By default, the provider ignores any annotations whose key names end with`,
				},
				resource.Attribute{
					Name:        "generate_name",
					Description: `(Optional) Prefix, used by the server, to generate a unique name ONLY IF the ` + "`" + `name` + "`" + ` field has not been provided. This value will also be combined with a unique suffix. For more info see [Kubernetes reference](hhttps://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#idempotency)`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) Map of string keys and values that can be used to organize and categorize (scope and select) the role.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the role, must be unique. Cannot be updated. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/identifiers#names)`,
				},
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) Namespace defines the space within which name of the role must be unique. #### Attributes`,
				},
				resource.Attribute{
					Name:        "generation",
					Description: `A sequence number representing a specific generation of the desired state.`,
				},
				resource.Attribute{
					Name:        "resource_version",
					Description: `An opaque value that represents the internal version of this role that can be used by clients to determine when role has changed. For more info see [Kubernetes reference](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency)`,
				},
				resource.Attribute{
					Name:        "uid",
					Description: `The unique in time and space value for this role. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/identifiers#uids) ### ` + "`" + `rule` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "api_groups",
					Description: `(Required) List of APIGroups that contains the resources.`,
				},
				resource.Attribute{
					Name:        "resources",
					Description: `(Required) List of resources that the rule applies to.`,
				},
				resource.Attribute{
					Name:        "resource_names",
					Description: `(Optional) White list of names that the rule applies to.`,
				},
				resource.Attribute{
					Name:        "verbs",
					Description: `(Required) List of Verbs that apply to ALL the ResourceKinds and AttributeRestrictions contained in this rule. ## Import Role can be imported using the namespace and name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import kubernetes_role.example default/terraform-example ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "kubernetes_role_binding",
			Category:         "Resources",
			ShortDescription: `A RoleBinding may be used to grant permission at the namespace level.`,
			Description:      ``,
			Keywords: []string{
				"role",
				"binding",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "metadata",
					Description: `(Required) Standard kubernetes metadata. For more info see [Kubernetes reference](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata)`,
				},
				resource.Attribute{
					Name:        "role_ref",
					Description: `(Required) The Role to bind Subjects to. For more info see [Kubernetes reference](https://kubernetes.io/docs/admin/authorization/rbac/#rolebinding-and-clusterrolebinding)`,
				},
				resource.Attribute{
					Name:        "subject",
					Description: `(Required) The Users, Groups, or ServiceAccounts to grand permissions to. For more info see [Kubernetes reference](https://kubernetes.io/docs/admin/authorization/rbac/#referring-to-subjects) ## Nested Blocks ### ` + "`" + `metadata` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "annotations",
					Description: `(Optional) An unstructured key value map stored with the role binding that may be used to store arbitrary metadata. ~> By default, the provider ignores any annotations whose key names end with`,
				},
				resource.Attribute{
					Name:        "generate_name",
					Description: `(Optional) Prefix, used by the server, to generate a unique name ONLY IF the ` + "`" + `name` + "`" + ` field has not been provided. This value will also be combined with a unique suffix. For more info see [Kubernetes reference](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#idempotency)`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) Map of string keys and values that can be used to organize and categorize (scope and select) the role binding. ~> By default, the provider ignores any labels whose key names end with`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the role binding, must be unique. Cannot be updated. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/identifiers#names)`,
				},
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) Namespace defines the space within which name of the role binding must be unique. #### Attributes`,
				},
				resource.Attribute{
					Name:        "generation",
					Description: `A sequence number representing a specific generation of the desired state.`,
				},
				resource.Attribute{
					Name:        "resource_version",
					Description: `An opaque value that represents the internal version of this object that can be used by clients to determine when the object has changed. For more info see [Kubernetes reference](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency)`,
				},
				resource.Attribute{
					Name:        "uid",
					Description: `The unique in time and space value for this role binding. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/identifiers#uids) ### ` + "`" + `role_ref` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of this Role to bind Subjects to.`,
				},
				resource.Attribute{
					Name:        "kind",
					Description: `(Required) The type of binding to use. This value must be present and defaults to ` + "`" + `Role` + "`" + ``,
				},
				resource.Attribute{
					Name:        "api_group",
					Description: `(Required) The API group to drive authorization decisions. This value must be and defaults to ` + "`" + `rbac.authorization.k8s.io` + "`" + ` ### ` + "`" + `subject` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of this Role to bind Subjects to.`,
				},
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) Namespace defines the namespace of the ServiceAccount to bind to. This value only applies to kind ` + "`" + `ServiceAccount` + "`" + ``,
				},
				resource.Attribute{
					Name:        "kind",
					Description: `(Required) The type of binding to use. This value must be ` + "`" + `ServiceAccount` + "`" + `, ` + "`" + `User` + "`" + ` or ` + "`" + `Group` + "`" + ``,
				},
				resource.Attribute{
					Name:        "api_group",
					Description: `(Required) The API group to drive authorization decisions. This value only applies to kind ` + "`" + `User` + "`" + ` and ` + "`" + `Group` + "`" + `. It must be ` + "`" + `rbac.authorization.k8s.io` + "`" + ` ## Import RoleBinding can be imported using the name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import kubernetes_role_binding.example default/terraform-name ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "kubernetes_secret",
			Category:         "Resources",
			ShortDescription: `The resource provides mechanisms to inject containers with sensitive information while keeping containers agnostic of Kubernetes.`,
			Description:      ``,
			Keywords: []string{
				"secret",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "data",
					Description: `(Optional) A map of the secret data.`,
				},
				resource.Attribute{
					Name:        "binary_data",
					Description: `(Optional) A map base64 encoded map of the secret data.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `(Required) Standard secret's metadata. For more info see [Kubernetes reference](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata)`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) The secret type. Defaults to ` + "`" + `Opaque` + "`" + `. For more info see [Kubernetes reference](https://github.com/kubernetes/community/blob/c7151dd8dd7e487e96e5ce34c6a416bb3b037609/contributors/design-proposals/auth/secrets.md#proposed-design) ## Nested Blocks ### ` + "`" + `metadata` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "annotations",
					Description: `(Optional) An unstructured key value map stored with the secret that may be used to store arbitrary metadata. ~> By default, the provider ignores any annotations whose key names end with`,
				},
				resource.Attribute{
					Name:        "generate_name",
					Description: `(Optional) Prefix, used by the server, to generate a unique name ONLY IF the ` + "`" + `name` + "`" + ` field has not been provided. This value will also be combined with a unique suffix. For more info see [Kubernetes reference](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#idempotency)`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) Map of string keys and values that can be used to organize and categorize (scope and select) the secret. May match selectors of replication controllers and services. ~> By default, the provider ignores any labels whose key names end with`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the secret, must be unique. Cannot be updated. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/identifiers#names)`,
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
					Description: `The unique in time and space value for this secret. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/identifiers#uids) ## Import Secret can be imported using its namespace and name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import kubernetes_secret.example default/my-secret ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "kubernetes_service",
			Category:         "Resources",
			ShortDescription: `A Service is an abstraction which defines a logical set of pods and a policy by which to access them - sometimes called a micro-service.`,
			Description:      ``,
			Keywords: []string{
				"service",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "metadata",
					Description: `(Required) Standard service's metadata. For more info see [Kubernetes reference](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata)`,
				},
				resource.Attribute{
					Name:        "spec",
					Description: `(Required) Spec defines the behavior of a service. [Kubernetes reference](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#spec-and-status)`,
				},
				resource.Attribute{
					Name:        "wait_for_load_balancer",
					Description: `(Optional) Terraform will wait for the load balancer to have at least 1 endpoint before considering the resource created. Defaults to ` + "`" + `true` + "`" + `. ## Nested Blocks ### ` + "`" + `metadata` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "annotations",
					Description: `(Optional) An unstructured key value map stored with the service that may be used to store arbitrary metadata. ~> By default, the provider ignores any annotations whose key names end with`,
				},
				resource.Attribute{
					Name:        "generate_name",
					Description: `(Optional) Prefix, used by the server, to generate a unique name ONLY IF the ` + "`" + `name` + "`" + ` field has not been provided. This value will also be combined with a unique suffix. For more info see [Kubernetes reference](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#idempotency)`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) Map of string keys and values that can be used to organize and categorize (scope and select) the service. May match selectors of replication controllers and services. ~> By default, the provider ignores any labels whose key names end with`,
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
					Name:        "generation",
					Description: `A sequence number representing a specific generation of the desired state.`,
				},
				resource.Attribute{
					Name:        "resource_version",
					Description: `An opaque value that represents the internal version of this service that can be used by clients to determine when service has changed. For more info see [Kubernetes reference](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency)`,
				},
				resource.Attribute{
					Name:        "uid",
					Description: `The unique in time and space value for this service. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/identifiers#uids) ### ` + "`" + `spec` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "cluster_ip",
					Description: `(Optional) The IP address of the service. It is usually assigned randomly by the master. If an address is specified manually and is not in use by others, it will be allocated to the service; otherwise, creation of the service will fail. ` + "`" + `None` + "`" + ` can be specified for headless services when proxying is not required. Ignored if type is ` + "`" + `ExternalName` + "`" + `. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/services#virtual-ips-and-service-proxies)`,
				},
				resource.Attribute{
					Name:        "external_ips",
					Description: `(Optional) A list of IP addresses for which nodes in the cluster will also accept traffic for this service. These IPs are not managed by Kubernetes. The user is responsible for ensuring that traffic arrives at a node with this IP. A common example is external load-balancers that are not part of the Kubernetes system.`,
				},
				resource.Attribute{
					Name:        "external_name",
					Description: `(Optional) The external reference that kubedns or equivalent will return as a CNAME record for this service. No proxying will be involved. Must be a valid DNS name and requires ` + "`" + `type` + "`" + ` to be ` + "`" + `ExternalName` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "external_traffic_policy",
					Description: `(Optional) Denotes if this Service desires to route external traffic to node-local or cluster-wide endpoints. ` + "`" + `Local` + "`" + ` preserves the client source IP and avoids a second hop for LoadBalancer and Nodeport type services, but risks potentially imbalanced traffic spreading. ` + "`" + `Cluster` + "`" + ` obscures the client source IP and may cause a second hop to another node, but should have good overall load-spreading. For more info: https://kubernetes.io/docs/tutorials/services/source-ip/`,
				},
				resource.Attribute{
					Name:        "load_balancer_ip",
					Description: `(Optional) Only applies to ` + "`" + `type = LoadBalancer` + "`" + `. LoadBalancer will get created with the IP specified in this field. This feature depends on whether the underlying cloud-provider supports specifying this field when a load balancer is created. This field will be ignored if the cloud-provider does not support the feature.`,
				},
				resource.Attribute{
					Name:        "load_balancer_source_ranges",
					Description: `(Optional) If specified and supported by the platform, this will restrict traffic through the cloud-provider load-balancer will be restricted to the specified client IPs. This field will be ignored if the cloud-provider does not support the feature. For more info see [Kubernetes reference](https://kubernetes.io/docs/tasks/access-application-cluster/create-external-load-balancer/).`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Required) The list of ports that are exposed by this service. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/services#virtual-ips-and-service-proxies)`,
				},
				resource.Attribute{
					Name:        "publish_not_ready_addresses",
					Description: `(Optional) When set to true, indicates that DNS implementations must publish the ` + "`" + `notReadyAddresses` + "`" + ` of subsets for the Endpoints associated with the Service. The default value is ` + "`" + `false` + "`" + `. The primary use case for setting this field is to use a StatefulSet's Headless Service to propagate ` + "`" + `SRV` + "`" + ` records for its Pods without respect to their readiness for purpose of peer discovery.`,
				},
				resource.Attribute{
					Name:        "selector",
					Description: `(Optional) Route service traffic to pods with label keys and values matching this selector. Only applies to types ` + "`" + `ClusterIP` + "`" + `, ` + "`" + `NodePort` + "`" + `, and ` + "`" + `LoadBalancer` + "`" + `. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/services#overview)`,
				},
				resource.Attribute{
					Name:        "session_affinity",
					Description: `(Optional) Used to maintain session affinity. Supports ` + "`" + `ClientIP` + "`" + ` and ` + "`" + `None` + "`" + `. Defaults to ` + "`" + `None` + "`" + `. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/services#virtual-ips-and-service-proxies)`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) Determines how the service is exposed. Defaults to ` + "`" + `ClusterIP` + "`" + `. Valid options are ` + "`" + `ExternalName` + "`" + `, ` + "`" + `ClusterIP` + "`" + `, ` + "`" + `NodePort` + "`" + `, and ` + "`" + `LoadBalancer` + "`" + `. ` + "`" + `ExternalName` + "`" + ` maps to the specified ` + "`" + `external_name` + "`" + `. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/services#overview)`,
				},
				resource.Attribute{
					Name:        "health_check_node_port",
					Description: `(Optional) Specifies the Healthcheck NodePort for the service. Only effects when type is set to ` + "`" + `LoadBalancer` + "`" + ` and external_traffic_policy is set to ` + "`" + `Local` + "`" + `. ### ` + "`" + `port` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of this port within the service. All ports within the service must have unique names. Optional if only one ServicePort is defined on this service.`,
				},
				resource.Attribute{
					Name:        "node_port",
					Description: `(Optional) The port on each node on which this service is exposed when ` + "`" + `type` + "`" + ` is ` + "`" + `NodePort` + "`" + ` or ` + "`" + `LoadBalancer` + "`" + `. Usually assigned by the system. If specified, it will be allocated to the service if unused or else creation of the service will fail. Default is to auto-allocate a port if the ` + "`" + `type` + "`" + ` of this service requires one. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/services#type--nodeport)`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Required) The port that will be exposed by this service.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Optional) The IP protocol for this port. Supports ` + "`" + `TCP` + "`" + ` and ` + "`" + `UDP` + "`" + `. Default is ` + "`" + `TCP` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "target_port",
					Description: `(Optional) Number or name of the port to access on the pods targeted by the service. Number must be in the range 1 to 65535. This field is ignored for services with ` + "`" + `cluster_ip = "None"` + "`" + `. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/services#defining-a-service) ## Attributes`,
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
					Description: `Hostname is set for load-balancer ingress points that are DNS based (typically AWS load-balancers). ### Timeouts ` + "`" + `kubernetes_service` + "`" + ` provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default ` + "`" + `10 minutes` + "`" + ` ## Import Service can be imported using its namespace and name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import kubernetes_service.example default/terraform-name ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "kubernetes_service_account",
			Category:         "Resources",
			ShortDescription: `A service account provides an identity for processes that run in a Pod.`,
			Description:      ``,
			Keywords: []string{
				"service",
				"account",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "metadata",
					Description: `(Required) Standard service account's metadata. For more info see [Kubernetes reference](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata)`,
				},
				resource.Attribute{
					Name:        "image_pull_secret",
					Description: `(Optional) A list of references to secrets in the same namespace to use for pulling any images in pods that reference this Service Account. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/secrets#manually-specifying-an-imagepullsecret)`,
				},
				resource.Attribute{
					Name:        "secret",
					Description: `(Optional) A list of secrets allowed to be used by pods running using this Service Account. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/secrets)`,
				},
				resource.Attribute{
					Name:        "automount_service_account_token",
					Description: `(Optional) Boolean, ` + "`" + `true` + "`" + ` to enable automatic mounting of the service account token. Defaults to ` + "`" + `true` + "`" + `. ## Nested Blocks ### ` + "`" + `metadata` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "annotations",
					Description: `(Optional) An unstructured key value map stored with the service account that may be used to store arbitrary metadata. ~> By default, the provider ignores any annotations whose key names end with`,
				},
				resource.Attribute{
					Name:        "generate_name",
					Description: `(Optional) Prefix, used by the server, to generate a unique name ONLY IF the ` + "`" + `name` + "`" + ` field has not been provided. This value will also be combined with a unique suffix. For more info see [Kubernetes reference](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#idempotency)`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) Map of string keys and values that can be used to organize and categorize (scope and select) the service account. May match selectors of replication controllers and services. ~> By default, the provider ignores any labels whose key names end with`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the service account, must be unique. Cannot be updated. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/identifiers#names)`,
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
					Description: `The unique in time and space value for this service account. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/identifiers#uids) ### ` + "`" + `image_pull_secret` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the referent. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/identifiers#names) ### ` + "`" + `secret` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the referent. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/identifiers#names) ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "default_secret_name",
					Description: `Name of the default secret, containing service account token, created & managed by the service. ## Import Service account can be imported using the namespace and name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import kubernetes_service_account.example default/terraform-example ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "default_secret_name",
					Description: `Name of the default secret, containing service account token, created & managed by the service. ## Import Service account can be imported using the namespace and name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import kubernetes_service_account.example default/terraform-example ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "kubernetes_stateful_set",
			Category:         "Resources",
			ShortDescription: `StatefulSet is a Kubernetes resource used to manage stateful applications.`,
			Description:      ``,
			Keywords: []string{
				"stateful",
				"set",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "metadata",
					Description: `(Required) Standard Kubernetes object metadata. For more info see [Kubernetes reference](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata)`,
				},
				resource.Attribute{
					Name:        "spec",
					Description: `(Required) Spec defines the specification of the desired behavior of the stateful set. For more info see [Kubernetes reference](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#spec-and-status)`,
				},
				resource.Attribute{
					Name:        "wait_for_rollout",
					Description: `(Optional) Wait for the StatefulSet to finish rolling out. Defaults to ` + "`" + `true` + "`" + `. ## Nested Blocks ### ` + "`" + `metadata` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "annotations",
					Description: `(Optional) An unstructured key value map stored with the stateful set that may be used to store arbitrary metadata. ~> By default, the provider ignores any annotations whose key names end with`,
				},
				resource.Attribute{
					Name:        "generate_name",
					Description: `(Optional) Prefix, used by the server, to generate a unique name ONLY IF the ` + "`" + `name` + "`" + ` field has not been provided. This value will also be combined with a unique suffix. For more info see [Kubernetes reference](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#idempotency)`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) Map of string keys and values that can be used to organize and categorize (scope and select) the stateful set.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the stateful set, must be unique. Cannot be updated. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/identifiers#names)`,
				},
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) Namespace defines the space within which name of the stateful set must be unique. #### Attributes`,
				},
				resource.Attribute{
					Name:        "generation",
					Description: `A sequence number representing a specific generation of the desired state.`,
				},
				resource.Attribute{
					Name:        "resource_version",
					Description: `An opaque value that represents the internal version of this stateful set that can be used by clients to determine when stateful set has changed. For more info see [Kubernetes reference](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency)`,
				},
				resource.Attribute{
					Name:        "uid",
					Description: `The unique in time and space value for this stateful set. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/identifiers#uids) ### ` + "`" + `spec` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "pod_management_policy",
					Description: `(Optional) podManagementPolicy controls how pods are created during initial scale up, when replacing pods on nodes, or when scaling down. The default policy is ` + "`" + `OrderedReady` + "`" + `, where pods are created in increasing order (pod-0, then pod-1, etc) and the controller will wait until each pod is ready before continuing. When scaling down, the pods are removed in the opposite order. The alternative policy is ` + "`" + `Parallel` + "`" + ` which will create pods in parallel to match the desired scale without waiting, and on scale down will delete all pods at once.`,
				},
				resource.Attribute{
					Name:        "replicas",
					Description: `(Optional) The desired number of replicas of the given Template. These are replicas in the sense that they are instantiations of the same Template, but individual replicas also have a consistent identity. If unspecified, defaults to 1. This attribute is a string to be able to distinguish between explicit zero and not specified.`,
				},
				resource.Attribute{
					Name:        "revision_history_limit",
					Description: `(Optional) The maximum number of revisions that will be maintained in the StatefulSet's revision history. The revision history consists of all revisions not represented by a currently applied StatefulSetSpec version. The default value is 10.`,
				},
				resource.Attribute{
					Name:        "selector",
					Description: `(Required) A label query over pods that should match the replica count.`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `(Required) The name of the service that governs this StatefulSet. This service must exist before the StatefulSet, and is responsible for the network identity of the set. Pods get DNS/hostnames that follow the pattern: pod-specific-string.serviceName.default.svc.cluster.local where "pod-specific-string" is managed by the StatefulSet controller.`,
				},
				resource.Attribute{
					Name:        "template",
					Description: `(Required) The object that describes the pod that will be created if insufficient replicas are detected. Each pod stamped out by the StatefulSet will fulfill this Template, but have a unique identity from the rest of the StatefulSet.`,
				},
				resource.Attribute{
					Name:        "update_strategy",
					Description: `(Optional) Indicates the StatefulSet update strategy that will be employed to update Pods in the StatefulSet when a revision is made to Template.`,
				},
				resource.Attribute{
					Name:        "volume_claim_template",
					Description: `(Optional) A list of volume claims that pods are allowed to reference. A claim in this list takes precedence over any volumes in the template, with the same name.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `(Required) Standard object's metadata. For more info see [Kubernetes reference](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata).`,
				},
				resource.Attribute{
					Name:        "spec",
					Description: `(Optional) Specification of the desired behavior of the pod. For more info see [Kubernetes reference](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#spec-and-status). ## Nested Blocks ### ` + "`" + `spec.template.spec` + "`" + ` #### Arguments These arguments are the same as the for the ` + "`" + `spec` + "`" + ` block of a Pod. Please see the [Pod resource](pod.html#spec-1) for reference. ## Nested Blocks ### ` + "`" + `spec.update_strategy` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) Indicates the type of the StatefulSetUpdateStrategy. There are two valid update strategies, RollingUpdate and OnDelete. Default is ` + "`" + `RollingUpdate` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "rolling_update",
					Description: `(Optional) The RollingUpdate update strategy will update all Pods in a StatefulSet, in reverse ordinal order, while respecting the StatefulSet guarantees. ### ` + "`" + `spec.update_strategy.rolling_update` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "partition",
					Description: `(Optional) Indicates the ordinal at which the StatefulSet should be partitioned. You can perform a phased roll out (e.g. a linear, geometric, or exponential roll out) using a partitioned rolling update in a similar manner to how you rolled out a canary. To perform a phased roll out, set the partition to the ordinal at which you want the controller to pause the update. By setting the partition to 0, you allow the StatefulSet controller to continue the update process. Default value is ` + "`" + `0` + "`" + `. ## Nested Blocks ### ` + "`" + `spec.volume_claim_template` + "`" + ` One or more ` + "`" + `volume_claim_template` + "`" + ` blocks can be specified. #### Arguments Each takes the same attibutes as a ` + "`" + `kubernetes_persistent_volume_claim` + "`" + ` resource. Please see its [documentation](persistent_volume_claim.html#argument-reference) for reference. ## Timeouts The following [Timeout](/docs/configuration/resources.html#operation-timeouts) configuration options are available for the ` + "`" + `kubernetes_stateful_set` + "`" + ` resource:`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Default ` + "`" + `10 minutes` + "`" + `) Used for creating new StatefulSet`,
				},
				resource.Attribute{
					Name:        "update",
					Description: `(Default ` + "`" + `10 minutes` + "`" + `) Used for updating a StatefulSet`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Default ` + "`" + `10 minutes` + "`" + `) Used for destroying a StatefulSet ## Import kubernetes_stateful_set can be imported using its namespace and name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import kubernetes_stateful_set.example default/terraform-example ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "kubernetes_storage_class",
			Category:         "Resources",
			ShortDescription: `Storage class is the foundation of dynamic provisioning, allowing cluster administrators to define abstractions for the underlying storage platform.`,
			Description:      ``,
			Keywords: []string{
				"storage",
				"class",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "metadata",
					Description: `(Required) Standard storage class's metadata. For more info see [Kubernetes reference](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata)`,
				},
				resource.Attribute{
					Name:        "parameters",
					Description: `(Optional) The parameters for the provisioner that should create volumes of this storage class. Read more about [available parameters](https://kubernetes.io/docs/concepts/storage/storage-classes/#parameters).`,
				},
				resource.Attribute{
					Name:        "storage_provisioner",
					Description: `(Required) Indicates the type of the provisioner`,
				},
				resource.Attribute{
					Name:        "reclaim_policy",
					Description: `(Optional) Indicates the reclaim policy to use. If no reclaimPolicy is specified when a StorageClass object is created, it will default to Delete.`,
				},
				resource.Attribute{
					Name:        "volume_binding_mode",
					Description: `(Optional) Indicates when volume binding and dynamic provisioning should occur.`,
				},
				resource.Attribute{
					Name:        "allow_volume_expansion",
					Description: `(Optional) Indicates whether the storage class allow volume expand, default true.`,
				},
				resource.Attribute{
					Name:        "mount_options",
					Description: `(Optional) Persistent Volumes that are dynamically created by a storage class will have the mount options specified.`,
				},
				resource.Attribute{
					Name:        "allowed_topologies",
					Description: `(Optional) Restrict the node topologies where volumes can be dynamically provisioned. See [allowed_topologies](#allowed_topologies) ## Nested Blocks ### ` + "`" + `metadata` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "annotations",
					Description: `(Optional) An unstructured key value map stored with the storage class that may be used to store arbitrary metadata. ~> By default, the provider ignores any annotations whose key names end with`,
				},
				resource.Attribute{
					Name:        "generate_name",
					Description: `(Optional) Prefix, used by the server, to generate a unique name ONLY IF the ` + "`" + `name` + "`" + ` field has not been provided. This value will also be combined with a unique suffix. For more info see [Kubernetes reference](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#idempotency)`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) Map of string keys and values that can be used to organize and categorize (scope and select) the storage class. May match selectors of replication controllers and services. ~> By default, the provider ignores any labels whose key names end with`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the storage class, must be unique. Cannot be updated. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/identifiers#names) ### ` + "`" + `allowed_topologies` + "`" + ` ￼ #### Arguments ￼`,
				},
				resource.Attribute{
					Name:        "match_label_expressions",
					Description: `(Optional) A list of topology selector requirements by labels. See [match_label_expressions](#match_label_expressions) ### ` + "`" + `match_label_expressions` + "`" + ` ￼ #### Arguments ￼`,
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
					Description: `The unique in time and space value for this storage class. For more info see [Kubernetes reference](http://kubernetes.io/docs/user-guide/identifiers#uids) ## Import kubernetes_storage_class can be imported using its name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import kubernetes_storage_class.example terraform-example ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"kubernetes_api_service":                 0,
		"kubernetes_certificate_signing_request": 1,
		"kubernetes_cluster_role":                2,
		"kubernetes_cluster_role_binding":        3,
		"kubernetes_config_map":                  4,
		"kubernetes_cron_job":                    5,
		"kubernetes_csi_driver":                  6,
		"kubernetes_daemonset":                   7,
		"kubernetes_deployment":                  8,
		"kubernetes_endpoints":                   9,
		"kubernetes_horizontal_pod_autoscaler":   10,
		"kubernetes_ingress":                     11,
		"kubernetes_job":                         12,
		"kubernetes_limit_range":                 13,
		"kubernetes_namespace":                   14,
		"kubernetes_network_policy":              15,
		"kubernetes_persistent_volume":           16,
		"kubernetes_persistent_volume_claim":     17,
		"kubernetes_pod":                         18,
		"kubernetes_pod_disruption_budget":       19,
		"kubernetes_pod_security_policy":         20,
		"kubernetes_priority_class":              21,
		"kubernetes_replication_controller":      22,
		"kubernetes_resource_quota":              23,
		"kubernetes_role":                        24,
		"kubernetes_role_binding":                25,
		"kubernetes_secret":                      26,
		"kubernetes_service":                     27,
		"kubernetes_service_account":             28,
		"kubernetes_stateful_set":                29,
		"kubernetes_storage_class":               30,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
