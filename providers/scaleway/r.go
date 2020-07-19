package scaleway

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "scaleway_scaleway_account_ssh_key",
			Category:         "Account Resources",
			ShortDescription: `Manages Scaleway user SSH keys.`,
			Description:      ``,
			Keywords: []string{
				"account",
				"ssh",
				"key",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "scaleway_baremetal_server",
			Category:         "Baremetal Resources",
			ShortDescription: `Manages Scaleway Compute Baremetal servers.`,
			Description:      ``,
			Keywords: []string{
				"baremetal",
				"server",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "scaleway_instance_ip",
			Category:         "Instance Resources",
			ShortDescription: `Manages Scaleway Compute Instance IPs.`,
			Description:      ``,
			Keywords: []string{
				"instance",
				"ip",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "instance_ip_reverse_dns",
			Category:         "Instance Resources",
			ShortDescription: `Manages Scaleway Compute Instance IPs' reverse DNS.`,
			Description:      ``,
			Keywords: []string{
				"instance",
				"ip",
				"reverse",
				"dns",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "scaleway_instance_placement_group",
			Category:         "Instance Resources",
			ShortDescription: `Manages Scaleway Compute Instance Placement Groups.`,
			Description:      ``,
			Keywords: []string{
				"instance",
				"placement",
				"group",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "scaleway_instance_security_group",
			Category:         "Instance Resources",
			ShortDescription: `Manages Scaleway Compute Instance security groups.`,
			Description:      ``,
			Keywords: []string{
				"instance",
				"security",
				"group",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "scaleway_instance_security_group_rules",
			Category:         "Instance Resources",
			ShortDescription: `Manages Scaleway Compute Instance security group rules.`,
			Description:      ``,
			Keywords: []string{
				"instance",
				"security",
				"group",
				"rules",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "scaleway_instance_server",
			Category:         "Instance Resources",
			ShortDescription: `Manages Scaleway Compute Instance servers.`,
			Description:      ``,
			Keywords: []string{
				"instance",
				"server",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "scaleway_instance_volume",
			Category:         "Instance Resources",
			ShortDescription: `Manages Scaleway Compute Instance Volumes.`,
			Description:      ``,
			Keywords: []string{
				"instance",
				"volume",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "scaleway_ip",
			Category:         "Deprecated Resources",
			ShortDescription: `Manages Scaleway IPs.`,
			Description:      ``,
			Keywords: []string{
				"deprecated",
				"ip",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "server",
					Description: `(Optional) ID of server to associate IP with`,
				},
				resource.Attribute{
					Name:        "reverse",
					Description: `(Deprecated) Please us the scaleway_ip_reverse_dns resource instead. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the new resource`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `IP of the new resource`,
				},
				resource.Attribute{
					Name:        "server",
					Description: `ID of the associated server resource`,
				},
				resource.Attribute{
					Name:        "reverse",
					Description: `reverse DNS setting of the IP resource ## Import Instances can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import scaleway_ip.jump_host 5faef9cd-ea9b-4a63-9171-9e26bec03dbc ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the new resource`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `IP of the new resource`,
				},
				resource.Attribute{
					Name:        "server",
					Description: `ID of the associated server resource`,
				},
				resource.Attribute{
					Name:        "reverse",
					Description: `reverse DNS setting of the IP resource ## Import Instances can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import scaleway_ip.jump_host 5faef9cd-ea9b-4a63-9171-9e26bec03dbc ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "scaleway_k8s_cluster_beta",
			Category:         "Kubernetes Resources (Beta)",
			ShortDescription: `Manages Scaleway Kubernetes clusters.`,
			Description:      ``,
			Keywords: []string{
				"kubernetes",
				"beta",
				"k8s",
				"cluster",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "scaleway_k8s_pool_beta",
			Category:         "Kubernetes Resources (Beta)",
			ShortDescription: `Manages Scaleway Kubernetes cluster pools.`,
			Description:      ``,
			Keywords: []string{
				"kubernetes",
				"beta",
				"k8s",
				"pool",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "scaleway_lb_backend_beta",
			Category:         "Load-Balancer Resources (Beta)",
			ShortDescription: `Manages Scaleway Load-Balancer Backends.`,
			Description:      ``,
			Keywords: []string{
				"load",
				"balancer",
				"beta",
				"lb",
				"backend",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "scaleway_lb_beta",
			Category:         "Load-Balancer Resources (Beta)",
			ShortDescription: `Manages Scaleway Load-Balancers.`,
			Description:      ``,
			Keywords: []string{
				"load",
				"balancer",
				"beta",
				"lb",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "scaleway_lb_certificate_beta",
			Category:         "Load-Balancer Resources (Beta)",
			ShortDescription: `Manages Scaleway Load-Balancer Certificates.`,
			Description:      ``,
			Keywords: []string{
				"load",
				"balancer",
				"beta",
				"lb",
				"certificate",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "scaleway_lb_frontend_beta",
			Category:         "Load-Balancer Resources (Beta)",
			ShortDescription: `Manages Scaleway Load-Balancer Frontends.`,
			Description:      ``,
			Keywords: []string{
				"load",
				"balancer",
				"beta",
				"lb",
				"frontend",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "scaleway_lb_ip_beta",
			Category:         "Load-Balancer Resources (Beta)",
			ShortDescription: `Manages Scaleway Load-Balancers IPs.`,
			Description:      ``,
			Keywords: []string{
				"load",
				"balancer",
				"beta",
				"lb",
				"ip",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "scaleway_object_bucket",
			Category:         "Object Storage Resources",
			ShortDescription: `Manages Scaleway object storage buckets.`,
			Description:      ``,
			Keywords: []string{
				"object",
				"storage",
				"bucket",
			},
			Arguments: []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The unique name of the bucket. ## Import Buckets can be imported using the ` + "`" + `{region}/{bucketName}` + "`" + ` identifier, e.g. ` + "`" + `` + "`" + `` + "`" + `bash $ terraform import scaleway_object_bucket.some_bucket fr-par/some-bucket ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "scaleway_rdb_instance_beta",
			Category:         "Database Resources (Beta)",
			ShortDescription: `Manages Scaleway Database Instances.`,
			Description:      ``,
			Keywords: []string{
				"database",
				"beta",
				"rdb",
				"instance",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "scaleway_registry_namespace_beta",
			Category:         "Container Registry Resources (Beta)",
			ShortDescription: `Manages Scaleway Container Registries.`,
			Description:      ``,
			Keywords: []string{
				"container",
				"registry",
				"beta",
				"namespace",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "scaleway_security_group",
			Category:         "Deprecated Resources",
			ShortDescription: `Manages Scaleway security groups.`,
			Description:      ``,
			Keywords: []string{
				"deprecated",
				"security",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of security group`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Required) description of security group`,
				},
				resource.Attribute{
					Name:        "enable_default_security",
					Description: `(Optional) default: true. Add default security group rules`,
				},
				resource.Attribute{
					Name:        "stateful",
					Description: `(Optional) default: false. Mark the security group as stateful. Note that stateful security groups can not be associated with bare metal servers`,
				},
				resource.Attribute{
					Name:        "inbound_default_policy",
					Description: `(Optional) default policy for inbound traffic. Can be one of accept or drop`,
				},
				resource.Attribute{
					Name:        "outbound_default_policy",
					Description: `(Optional) default policy for outbound traffic. Can be one of accept or drop Field ` + "`" + `name` + "`" + `, ` + "`" + `description` + "`" + ` are editable. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `id of the new resource ## Import Instances can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import scaleway_security_group.test 5faef9cd-ea9b-4a63-9171-9e26bec03dbc ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `id of the new resource ## Import Instances can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import scaleway_security_group.test 5faef9cd-ea9b-4a63-9171-9e26bec03dbc ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "scaleway_security_group_rule",
			Category:         "Deprecated Resources",
			ShortDescription: `Manages Scaleway security group rules.`,
			Description:      ``,
			Keywords: []string{
				"deprecated",
				"security",
				"group",
				"rule",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "security_group",
					Description: `(Required) the security group which should be associated with this rule`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Required) action of rule (` + "`" + `accept` + "`" + `, ` + "`" + `drop` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "direction",
					Description: `(Required) direction of rule (` + "`" + `inbound` + "`" + `, ` + "`" + `outbound` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "ip_range",
					Description: `(Required) ip_range of rule`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Required) protocol of rule (` + "`" + `ICMP` + "`" + `, ` + "`" + `TCP` + "`" + `, ` + "`" + `UDP` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional) port of the rule Fields ` + "`" + `action` + "`" + `, ` + "`" + `direction` + "`" + `, ` + "`" + `ip_range` + "`" + `, ` + "`" + `protocol` + "`" + `, ` + "`" + `port` + "`" + ` are editable. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `id of the new resource`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `id of the new resource`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "scaleway_server",
			Category:         "Deprecated Resources",
			ShortDescription: `Manages Scaleway servers.`,
			Description:      ``,
			Keywords: []string{
				"deprecated",
				"server",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of server`,
				},
				resource.Attribute{
					Name:        "image",
					Description: `(Required) base image of server`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) type of server`,
				},
				resource.Attribute{
					Name:        "bootscript",
					Description: `(Optional) server bootscript`,
				},
				resource.Attribute{
					Name:        "boot_type",
					Description: `(Optional) the boot mechanism for this server. Possible values include ` + "`" + `local` + "`" + ` and ` + "`" + `bootscript` + "`" + ``,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) list of tags for server`,
				},
				resource.Attribute{
					Name:        "enable_ipv6",
					Description: `(Optional) enable ipv6`,
				},
				resource.Attribute{
					Name:        "dynamic_ip_required",
					Description: `(Optional) make server publicly available`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `(Optional) set a public ip previously created (a real ip is expected here, not its resource id)`,
				},
				resource.Attribute{
					Name:        "security_group",
					Description: `(Optional) assign security group to server`,
				},
				resource.Attribute{
					Name:        "volume",
					Description: `(Optional) attach additional volumes to your instance (see below)`,
				},
				resource.Attribute{
					Name:        "public_ipv6",
					Description: `(Read Only) if ` + "`" + `enable_ipv6` + "`" + ` is set this contains the ipv6 address of your instance`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `(Optional) allows you to define the desired state of your server. Valid values include (` + "`" + `stopped` + "`" + `, ` + "`" + `running` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "cloudinit",
					Description: `(Optional) allows you to define cloudinit script for this server`,
				},
				resource.Attribute{
					Name:        "state_detail",
					Description: `(Read Only) contains details from the scaleway API the state of your instance Field ` + "`" + `name` + "`" + `, ` + "`" + `type` + "`" + `, ` + "`" + `tags` + "`" + `, ` + "`" + `dynamic_ip_required` + "`" + `, ` + "`" + `security_group` + "`" + ` are editable. ## Volume You can attach additional volumes to your instance, which will share the lifetime of your ` + "`" + `scaleway_server` + "`" + ` resource. ~>`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) The type of volume. Can be ` + "`" + `"l_ssd"` + "`" + ``,
				},
				resource.Attribute{
					Name:        "size_in_gb",
					Description: `(Required) The size of the volume in gigabytes. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `id of the new resource`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `private ip of the new resource`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `public ip of the new resource ## Import Instances can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import scaleway_server.web 5faef9cd-ea9b-4a63-9171-9e26bec03dbc ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `id of the new resource`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `private ip of the new resource`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `public ip of the new resource ## Import Instances can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import scaleway_server.web 5faef9cd-ea9b-4a63-9171-9e26bec03dbc ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "scaleway_ssh_key",
			Category:         "Deprecated Resources",
			ShortDescription: `Manages Scaleway user SSH keys.`,
			Description:      ``,
			Keywords: []string{
				"deprecated",
				"ssh",
				"key",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "scaleway_token",
			Category:         "Deprecated Resources",
			ShortDescription: `Manages Scaleway Tokens.`,
			Description:      ``,
			Keywords: []string{
				"deprecated",
				"token",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "expires",
					Description: `(Optional) Define if the token should automatically expire or not`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `(Optional) Scaleway account email. Defaults to registered account`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional) Scaleway account password. Required for cross-account token management`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Token description ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Token ID - can be used to access scaleway API`,
				},
				resource.Attribute{
					Name:        "access_key",
					Description: `Token Access Key`,
				},
				resource.Attribute{
					Name:        "secret_key",
					Description: `Token Secret Key`,
				},
				resource.Attribute{
					Name:        "creation_ip",
					Description: `IP used to create the token`,
				},
				resource.Attribute{
					Name:        "expiration_date",
					Description: `Expiration date of token, if expiration is requested ## Import Instances can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import scaleway_token.karls_token 5faef9cd-ea9b-4a63-9171-9e26bec03dbc ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Token ID - can be used to access scaleway API`,
				},
				resource.Attribute{
					Name:        "access_key",
					Description: `Token Access Key`,
				},
				resource.Attribute{
					Name:        "secret_key",
					Description: `Token Secret Key`,
				},
				resource.Attribute{
					Name:        "creation_ip",
					Description: `IP used to create the token`,
				},
				resource.Attribute{
					Name:        "expiration_date",
					Description: `Expiration date of token, if expiration is requested ## Import Instances can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import scaleway_token.karls_token 5faef9cd-ea9b-4a63-9171-9e26bec03dbc ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "scaleway_user_data",
			Category:         "Deprecated Resources",
			ShortDescription: `Manages Scaleway Server UserData.`,
			Description:      ``,
			Keywords: []string{
				"deprecated",
				"user",
				"data",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "server",
					Description: `(Required) ID of server to associate the user data with`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) The key of the user data object`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) The value of the user data object ## Import Instances can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import scaleway_user_data.gcp userdata-<server-id>-<key> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "scaleway_volume",
			Category:         "Deprecated Resources",
			ShortDescription: `Manages Scaleway Volumes.`,
			Description:      ``,
			Keywords: []string{
				"deprecated",
				"volume",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of volume`,
				},
				resource.Attribute{
					Name:        "size_in_gb",
					Description: `(Required) size of the volume in GB`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) type of volume ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `id of the new resource`,
				},
				resource.Attribute{
					Name:        "server",
					Description: `(Read Only) the ` + "`" + `scaleway_server` + "`" + ` instance which has this volume mounted right now ## Import Instances can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import scaleway_volume.test 5faef9cd-ea9b-4a63-9171-9e26bec03dbc ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `id of the new resource`,
				},
				resource.Attribute{
					Name:        "server",
					Description: `(Read Only) the ` + "`" + `scaleway_server` + "`" + ` instance which has this volume mounted right now ## Import Instances can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import scaleway_volume.test 5faef9cd-ea9b-4a63-9171-9e26bec03dbc ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "scaleway_volume_attachment",
			Category:         "Deprecated Resources",
			ShortDescription: `Manages Scaleway Volume attachments for servers.`,
			Description:      ``,
			Keywords: []string{
				"deprecated",
				"volume",
				"attachment",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "server",
					Description: `(Required) id of the server`,
				},
				resource.Attribute{
					Name:        "volume",
					Description: `(Required) id of the volume to be attached ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `id of the new resource`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `id of the new resource`,
				},
			},
		},
	}

	resourcesMap = map[string]int{

		"scaleway_scaleway_account_ssh_key":      0,
		"scaleway_baremetal_server":              1,
		"scaleway_instance_ip":                   2,
		"instance_ip_reverse_dns":                3,
		"scaleway_instance_placement_group":      4,
		"scaleway_instance_security_group":       5,
		"scaleway_instance_security_group_rules": 6,
		"scaleway_instance_server":               7,
		"scaleway_instance_volume":               8,
		"scaleway_ip":                            9,
		"scaleway_k8s_cluster_beta":              10,
		"scaleway_k8s_pool_beta":                 11,
		"scaleway_lb_backend_beta":               12,
		"scaleway_lb_beta":                       13,
		"scaleway_lb_certificate_beta":           14,
		"scaleway_lb_frontend_beta":              15,
		"scaleway_lb_ip_beta":                    16,
		"scaleway_object_bucket":                 17,
		"scaleway_rdb_instance_beta":             18,
		"scaleway_registry_namespace_beta":       19,
		"scaleway_security_group":                20,
		"scaleway_security_group_rule":           21,
		"scaleway_server":                        22,
		"scaleway_ssh_key":                       23,
		"scaleway_token":                         24,
		"scaleway_user_data":                     25,
		"scaleway_volume":                        26,
		"scaleway_volume_attachment":             27,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
