package rke

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "rke_cluster",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "delay_on_creation",
					Description: `(Optional) RKE k8s cluster delay on creation (int)`,
				},
				resource.Attribute{
					Name:        "disable_port_check",
					Description: `(Optional) Enable/Disable RKE k8s cluster port checking. Default ` + "`" + `false` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "addon_job_timeout",
					Description: `(Optional) RKE k8s cluster addon deployment timeout in seconds for status check (int)`,
				},
				resource.Attribute{
					Name:        "addons",
					Description: `(Optional) RKE k8s cluster user addons YAML manifest to be deployed (string)`,
				},
				resource.Attribute{
					Name:        "addons_include",
					Description: `(Optional) RKE k8s cluster user addons YAML manifest urls or paths to be deployed (list)`,
				},
				resource.Attribute{
					Name:        "authentication",
					Description: `(Optional) RKE k8s cluster authentication configuration (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "authorization",
					Description: `(Optional) RKE k8s cluster authorization mode configuration (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "bastion_host",
					Description: `(Optional) RKE k8s cluster bastion Host configuration (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "cert_dir",
					Description: `(Optional) Specify a certificate dir path (string)`,
				},
				resource.Attribute{
					Name:        "cloud_provider",
					Description: `(Optional) RKE k8s cluster cloud provider configuration [rke-cloud-providers](https://rancher.com/docs/rke/latest/en/config-options/cloud-providers/) (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "cluster_name",
					Description: `(Optional) RKE k8s cluster name used in the kube config (string)`,
				},
				resource.Attribute{
					Name:        "cluster_yaml",
					Description: `(Optional) RKE k8s cluster config yaml encoded. Provider arguments take precedence over this one (string)`,
				},
				resource.Attribute{
					Name:        "custom_certs",
					Description: `(Optional) Use custom certificates from a cert dir (string)`,
				},
				resource.Attribute{
					Name:        "dind",
					Description: `(Optional/Experimental) Deploy RKE cluster on a dind environment. Default: ` + "`" + `false` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "dind_storage_driver",
					Description: `(Optional/Experimental) DinD RKE cluster storage driver (string)`,
				},
				resource.Attribute{
					Name:        "dind_dns_server",
					Description: `(Optional/Experimental) DinD RKE cluster dns (string)`,
				},
				resource.Attribute{
					Name:        "dns",
					Description: `(Optional) RKE k8s cluster DNS Config (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "enable_cri_dockerd",
					Description: `(Optional) Enable/Disable CRI dockerd for kubelet. Default ` + "`" + `false` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "ignore_docker_version",
					Description: `(Optional) Enable/Disable RKE k8s cluster strict docker version checking. Default ` + "`" + `false` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "ingress",
					Description: `(Optional) RKE k8s cluster ingress controller configuration (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "kubernetes_version",
					Description: `(Optional) K8s version to deploy. If kubernetes image is specified, image version takes precedence. Default: ` + "`" + `rke default` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "monitoring",
					Description: `(Optional) RKE k8s cluster monitoring Config (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "network",
					Description: `(Optional) RKE k8s cluster network configuration (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "nodes",
					Description: `(Optional) RKE k8s cluster nodes (list)`,
				},
				resource.Attribute{
					Name:        "prefix_path",
					Description: `(Optional/Computed) RKE k8s directory path (string)`,
				},
				resource.Attribute{
					Name:        "private_registries",
					Description: `(Optional/Computed) RKE k8s cluster private docker registries (list)`,
				},
				resource.Attribute{
					Name:        "restore",
					Description: `(Optional/Computed) RKE k8s cluster restore configuration (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "rotate_certificates",
					Description: `(Optional) RKE k8s cluster rotate certificates configuration (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "services",
					Description: `(Optional) RKE k8s cluster services (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "services_etcd",
					Description: `(DEPRECATED) Use services.etcd instead (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "services_kube_api",
					Description: `(DEPRECATED) Use services.kube_api instead (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "services_kube_controller",
					Description: `(DEPRECATED) Use services.kube_controller instead (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "services_kubelet",
					Description: `(DEPRECATED) Use services.kubelet instead (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "services_kubeproxy",
					Description: `(DEPRECATED) Use services.kubeproxy instead (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "services_scheduler",
					Description: `(DEPRECATED) Use services.scheduler instead (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "ssh_agent_auth",
					Description: `(Optional/Computed) SSH Agent Auth enable (bool)`,
				},
				resource.Attribute{
					Name:        "ssh_cert_path",
					Description: `(Optional) SSH Certificate Path (string)`,
				},
				resource.Attribute{
					Name:        "ssh_key_path",
					Description: `(Optional) SSH Private Key Path (string)`,
				},
				resource.Attribute{
					Name:        "system_images",
					Description: `(Optional) RKE k8s cluster system images list (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "update_only",
					Description: `(Optional) Skip idempotent deployment of control and etcd plane. Default ` + "`" + `false` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "upgrade_strategy",
					Description: `(Optional) RKE k8s cluster upgrade strategy (list maxitems:1) ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Computed) The ID of the resource (string)`,
				},
				resource.Attribute{
					Name:        "ca_crt",
					Description: `(Computed/Sensitive) RKE k8s cluster CA certificate (string)`,
				},
				resource.Attribute{
					Name:        "client_cert",
					Description: `(Computed/Sensitive) RKE k8s cluster client certificate (string)`,
				},
				resource.Attribute{
					Name:        "client_key",
					Description: `(Computed/Sensitive) RKE k8s cluster client key (string)`,
				},
				resource.Attribute{
					Name:        "rke_state",
					Description: `(Computed/Sensitive) RKE k8s cluster state (string)`,
				},
				resource.Attribute{
					Name:        "kube_config_yaml",
					Description: `(Computed/Sensitive) RKE k8s cluster kube config yaml (string)`,
				},
				resource.Attribute{
					Name:        "internal_kube_config_yaml",
					Description: `(Computed/Sensitive) RKE k8s cluster internal kube config yaml (string)`,
				},
				resource.Attribute{
					Name:        "rke_cluster_yaml",
					Description: `(Computed/Sensitive) RKE k8s cluster config yaml (string)`,
				},
				resource.Attribute{
					Name:        "certificates",
					Description: `(Computed/Sensitive) RKE k8s cluster certificates (string)`,
				},
				resource.Attribute{
					Name:        "kube_admin_user",
					Description: `(Computed) RKE k8s cluster admin user (string)`,
				},
				resource.Attribute{
					Name:        "api_server_url",
					Description: `(Computed) RKE k8s cluster api server url (string)`,
				},
				resource.Attribute{
					Name:        "cluster_domain",
					Description: `(Computed) RKE k8s cluster domain (string)`,
				},
				resource.Attribute{
					Name:        "cluster_cidr",
					Description: `(Computed) RKE k8s cluster cidr (string)`,
				},
				resource.Attribute{
					Name:        "cluster_dns_server",
					Description: `(Computed) RKE k8s cluster dns server (string)`,
				},
				resource.Attribute{
					Name:        "control_plane_hosts",
					Description: `(Computed) RKE k8s cluster control plane nodes (list)`,
				},
				resource.Attribute{
					Name:        "etcd_hosts",
					Description: `(Computed) RKE k8s cluster etcd nodes (list)`,
				},
				resource.Attribute{
					Name:        "inactive_hosts",
					Description: `(Computed) RKE k8s cluster inactive nodes (list)`,
				},
				resource.Attribute{
					Name:        "worker_hosts",
					Description: `(Computed) RKE k8s cluster worker nodes (list)`,
				},
				resource.Attribute{
					Name:        "running_system_images",
					Description: `(Computed) RKE k8s cluster running system images list (list) ## Nested blocks ### ` + "`" + `authentication` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "sans",
					Description: `(Optional/Computed) List of additional hostnames and IPs to include in the api server PKI cert (list)`,
				},
				resource.Attribute{
					Name:        "strategy",
					Description: `(Optional) Authentication strategy that will be used in RKE k8s cluster. Default: ` + "`" + `x509` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "webhook",
					Description: `(Optional/Computed) Webhook configuration options (list maxitem: 1) #### ` + "`" + `webhook` + "`" + ` ##### Arguments`,
				},
				resource.Attribute{
					Name:        "config_file",
					Description: `(Optional) Multiline string that represent a custom webhook config file (string)`,
				},
				resource.Attribute{
					Name:        "cache_timeout",
					Description: `(Optional) Controls how long to cache authentication decisions (string) ### ` + "`" + `authorization` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `(Optional) RKE mode for authorization. ` + "`" + `rbac` + "`" + ` and ` + "`" + `none` + "`" + ` modes are available. Default ` + "`" + `rbac` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "options",
					Description: `(Optional/Computed) RKE options for authorization (map) ### ` + "`" + `bastion_host` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `(Required) Address of Bastion Host (string)`,
				},
				resource.Attribute{
					Name:        "user",
					Description: `(Required) SSH User to Bastion Host (string)`,
				},
				resource.Attribute{
					Name:        "ignore_proxy_env_vars",
					Description: `(Optional) Ignore proxy env vars at Bastion Host? Default: ` + "`" + `false` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional) SSH Port of Bastion Host. Default ` + "`" + `22` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "ssh_agent_auth",
					Description: `(Optional/Computed) SSH Agent Auth enable (bool)`,
				},
				resource.Attribute{
					Name:        "ssh_cert",
					Description: `(Optional/Sensitive) SSH Certificate Key (string)`,
				},
				resource.Attribute{
					Name:        "ssh_cert_path",
					Description: `(Optional) SSH Certificate Key Path (string)`,
				},
				resource.Attribute{
					Name:        "ssh_key",
					Description: `(Optional/Sensitive) SSH Private Key (string)`,
				},
				resource.Attribute{
					Name:        "ssh_key_path",
					Description: `(Optional) SSH Private Key Path (string) ### ` + "`" + `cloud_provider` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Cloud Provider name. ` + "`" + `aws` + "`" + `, ` + "`" + `azure` + "`" + `, ` + "`" + `custom` + "`" + `, ` + "`" + `external` + "`" + `, ` + "`" + `openstack` + "`" + `, ` + "`" + `vsphere` + "`" + ` are supported (string)`,
				},
				resource.Attribute{
					Name:        "aws_cloud_config",
					Description: `(DEPRECATED) Use aws_cloud_provider instead`,
				},
				resource.Attribute{
					Name:        "aws_cloud_provider",
					Description: `(Optional) AWS Cloud Provider config [rke-aws-cloud-provider](https://rancher.com/docs/rke/latest/en/config-options/cloud-providers/aws/) (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "azure_cloud_config",
					Description: `(DEPRECATED) Use azure_cloud_provider instead`,
				},
				resource.Attribute{
					Name:        "azure_cloud_provider",
					Description: `(Optional) Azure Cloud Provider config [rke-azure-cloud-provider](https://rancher.com/docs/rke/latest/en/config-options/cloud-providers/azure/) (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "custom_cloud_config",
					Description: `(DEPRECATED) Use custom_cloud_provider instead`,
				},
				resource.Attribute{
					Name:        "custom_cloud_provider",
					Description: `(Optional) Custom Cloud Provider config (string)`,
				},
				resource.Attribute{
					Name:        "openstack_cloud_config",
					Description: `(DEPRECATED) Use openstack_cloud_provider instead`,
				},
				resource.Attribute{
					Name:        "openstack_cloud_provider",
					Description: `(Optional/Computed) Openstack Cloud Provider config [rke-openstack-cloud-provider](https://rancher.com/docs/rke/latest/en/config-options/cloud-providers/openstack/) (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "vsphere_cloud_config",
					Description: `(DEPRECATED) Use vsphere_cloud_provider instead`,
				},
				resource.Attribute{
					Name:        "vsphere_cloud_provider",
					Description: `(Optional/Computed) Vsphere Cloud Provider config [rke-vsphere-cloud-provider](https://rancher.com/docs/rke/latest/en/config-options/cloud-providers/vsphere/) Extra argument ` + "`" + `name` + "`" + ` is required on ` + "`" + `virtual_center` + "`" + ` configuration. (list maxitems:1) #### ` + "`" + `aws_cloud_provider` + "`" + ` ##### Arguments`,
				},
				resource.Attribute{
					Name:        "global",
					Description: `(Optional) (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "service_override",
					Description: `(Optional) (list) ##### ` + "`" + `global` + "`" + ` ###### Arguments`,
				},
				resource.Attribute{
					Name:        "disable_security_group_ingress",
					Description: `(Optional) Disables the automatic ingress creation. Default ` + "`" + `false` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "disable_strict_zone_check",
					Description: `(Optional) Setting this to true will disable the check and provide a warning that the check was skipped. Default ` + "`" + `false` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "elb_security_group",
					Description: `(Optional) Use these ELB security groups instead create new (string)`,
				},
				resource.Attribute{
					Name:        "kubernetes_cluster_id",
					Description: `(Optional) The cluster id we'll use to identify our cluster resources (string)`,
				},
				resource.Attribute{
					Name:        "kubernetes_cluster_tag",
					Description: `(Optional) Legacy cluster id we'll use to identify our cluster resources (string)`,
				},
				resource.Attribute{
					Name:        "role_arn",
					Description: `(Optional) IAM role to assume when interaction with AWS APIs (string)`,
				},
				resource.Attribute{
					Name:        "route_table_id",
					Description: `(Optional) Enables using a specific RouteTable (string)`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Optional) Enables using a specific subnet to use for ELB's (string)`,
				},
				resource.Attribute{
					Name:        "vpc",
					Description: `(Optional) The AWS VPC flag enables the possibility to run the master components on a different aws account, on a different cloud provider or on-premises. If the flag is set also the KubernetesClusterTag must be provided (string)`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `(Optional) The AWS zone (string) ##### ` + "`" + `service_override` + "`" + ` ###### Arguments`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(DEPRECATED) Use service instead`,
				},
				resource.Attribute{
					Name:        "service",
					Description: `(Required) (string)`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) (string)`,
				},
				resource.Attribute{
					Name:        "signing_method",
					Description: `(Optional/Computed) (string)`,
				},
				resource.Attribute{
					Name:        "signing_name",
					Description: `(Optional) (string)`,
				},
				resource.Attribute{
					Name:        "signing_region",
					Description: `(Optional) (string)`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Optional) (string) #### ` + "`" + `azure_cloud_provider` + "`" + ` ##### Arguments`,
				},
				resource.Attribute{
					Name:        "aad_client_id",
					Description: `(Required/Sensitive) (string)`,
				},
				resource.Attribute{
					Name:        "aad_client_secret",
					Description: `(Required/Sensitive) (string)`,
				},
				resource.Attribute{
					Name:        "subscription_id",
					Description: `(Required/Sensitive) (string)`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Required/Sensitive) (string)`,
				},
				resource.Attribute{
					Name:        "aad_client_cert_password",
					Description: `(Optional/Computed/Sensitive) (string)`,
				},
				resource.Attribute{
					Name:        "aad_client_cert_path",
					Description: `(Optional/Computed) (string)`,
				},
				resource.Attribute{
					Name:        "cloud",
					Description: `(Optional/Computed) (string)`,
				},
				resource.Attribute{
					Name:        "cloud_provider_backoff",
					Description: `(Optional/Computed) (bool)`,
				},
				resource.Attribute{
					Name:        "cloud_provider_backoff_duration",
					Description: `(Optional/Computed) (int)`,
				},
				resource.Attribute{
					Name:        "cloud_provider_backoff_exponent",
					Description: `(Optional/Computed) (int)`,
				},
				resource.Attribute{
					Name:        "cloud_provider_backoff_jitter",
					Description: `(Optional/Computed) (int)`,
				},
				resource.Attribute{
					Name:        "cloud_provider_backoff_retries",
					Description: `(Optional/Computed) (int)`,
				},
				resource.Attribute{
					Name:        "cloud_provider_rate_limit",
					Description: `(Optional/Computed) (bool)`,
				},
				resource.Attribute{
					Name:        "cloud_provider_rate_limit_bucket",
					Description: `(Optional/Computed) (int)`,
				},
				resource.Attribute{
					Name:        "cloud_provider_rate_limit_qps",
					Description: `(Optional/Computed) (int)`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `(Optional/Computed) (string)`,
				},
				resource.Attribute{
					Name:        "maximum_load_balancer_rule_count",
					Description: `(Optional/Computed) (int)`,
				},
				resource.Attribute{
					Name:        "primary_availability_set_name",
					Description: `(Optional/Computed) (string)`,
				},
				resource.Attribute{
					Name:        "primary_scale_set_name",
					Description: `(Optional/Computed) (string)`,
				},
				resource.Attribute{
					Name:        "resource_group",
					Description: `(Optional/Computed) (string)`,
				},
				resource.Attribute{
					Name:        "route_table_name",
					Description: `(Optional/Computed) (string)`,
				},
				resource.Attribute{
					Name:        "security_group_name",
					Description: `(Optional/Computed) (string)`,
				},
				resource.Attribute{
					Name:        "subnet_name",
					Description: `(Optional/Computed) (string)`,
				},
				resource.Attribute{
					Name:        "use_instance_metadata",
					Description: `(Optional/Computed) (bool)`,
				},
				resource.Attribute{
					Name:        "use_managed_identity_extension",
					Description: `(Optional/Computed) (bool)`,
				},
				resource.Attribute{
					Name:        "vm_type",
					Description: `(Optional/Computed) (string)`,
				},
				resource.Attribute{
					Name:        "vnet_name",
					Description: `(Optional/Computed) (string)`,
				},
				resource.Attribute{
					Name:        "vnet_resource_group",
					Description: `(Optional/Computed) (string) #### ` + "`" + `openstack_cloud_provider` + "`" + ` ##### Arguments`,
				},
				resource.Attribute{
					Name:        "global",
					Description: `(Required) (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "block_storage",
					Description: `(Optional/Computed) (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "load_balancer",
					Description: `(Optional/Computed) (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `(Optional/Computed) (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "route",
					Description: `(Optional/Computed) (list maxitems:1) ##### ` + "`" + `global` + "`" + ` ###### Arguments`,
				},
				resource.Attribute{
					Name:        "auth_url",
					Description: `(Required) (string)`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Required/Sensitive) (string)`,
				},
				resource.Attribute{
					Name:        "ca_file",
					Description: `(Optional) (string)`,
				},
				resource.Attribute{
					Name:        "domain_id",
					Description: `(Optional/Sensitive) Required if ` + "`" + `domain_name` + "`" + ` not provided. (string)`,
				},
				resource.Attribute{
					Name:        "domain_name",
					Description: `(Optional) Required if ` + "`" + `domain_id` + "`" + ` not provided. (string)`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) (string)`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Optional/Sensitive) Required if ` + "`" + `tenant_name` + "`" + ` not provided. (string)`,
				},
				resource.Attribute{
					Name:        "tenant_name",
					Description: `(Optional) Required if ` + "`" + `tenant_id` + "`" + ` not provided. (string)`,
				},
				resource.Attribute{
					Name:        "trust_id",
					Description: `(Optional/Sensitive) (string)`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Optional) Required if ` + "`" + `user_id` + "`" + ` not provided. (string)`,
				},
				resource.Attribute{
					Name:        "user_id",
					Description: `(Optional/Sensitive) Required if ` + "`" + `username` + "`" + ` not provided. (string) ##### ` + "`" + `block_storage` + "`" + ` ###### Arguments`,
				},
				resource.Attribute{
					Name:        "bs_version",
					Description: `(Optional) (string)`,
				},
				resource.Attribute{
					Name:        "ignore_volume_az",
					Description: `(Optional) (string)`,
				},
				resource.Attribute{
					Name:        "trust_device_path",
					Description: `(Optional) (string) ##### ` + "`" + `load_balancer` + "`" + ` ###### Arguments`,
				},
				resource.Attribute{
					Name:        "create_monitor",
					Description: `(Optional) (bool)`,
				},
				resource.Attribute{
					Name:        "floating_network_id",
					Description: `(Optional) (string)`,
				},
				resource.Attribute{
					Name:        "lb_method",
					Description: `(Optional) (string)`,
				},
				resource.Attribute{
					Name:        "lb_provider",
					Description: `(Optional) (string)`,
				},
				resource.Attribute{
					Name:        "lb_version",
					Description: `(Optional) (string)`,
				},
				resource.Attribute{
					Name:        "manage_security_groups",
					Description: `(Optional) (bool)`,
				},
				resource.Attribute{
					Name:        "monitor_delay",
					Description: `(Optional) (string)`,
				},
				resource.Attribute{
					Name:        "monitor_max_retries",
					Description: `(Optional) (int)`,
				},
				resource.Attribute{
					Name:        "monitor_timeout",
					Description: `(Optional) (string)`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Optional) (string)`,
				},
				resource.Attribute{
					Name:        "use_octavia",
					Description: `(Optional) (bool) ##### ` + "`" + `metadata` + "`" + ` ###### Arguments`,
				},
				resource.Attribute{
					Name:        "request_timeout",
					Description: `(Optional) (int)`,
				},
				resource.Attribute{
					Name:        "search_order",
					Description: `(Optional) (string) ##### ` + "`" + `route` + "`" + ` ###### Arguments`,
				},
				resource.Attribute{
					Name:        "router_id",
					Description: `(Optional) (string) #### ` + "`" + `vsphere_cloud_provider` + "`" + ` ##### Arguments`,
				},
				resource.Attribute{
					Name:        "virtual_center",
					Description: `(Required) (List)`,
				},
				resource.Attribute{
					Name:        "workspace",
					Description: `(Required) (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "disk",
					Description: `(Optional/Computed) (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "global",
					Description: `(Optional/Computed) (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "network",
					Description: `(Optional/Computed) (list maxitems:1) ##### ` + "`" + `virtual_center` + "`" + ` ###### Arguments`,
				},
				resource.Attribute{
					Name:        "datacenters",
					Description: `(Required) (string)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of virtualcenter config for Vsphere Cloud Provider config (string)`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Required/Sensitive) (string)`,
				},
				resource.Attribute{
					Name:        "user",
					Description: `(Required/Sensitive) (string)`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional) (string)`,
				},
				resource.Attribute{
					Name:        "soap_roundtrip_count",
					Description: `(Optional) (int) ##### ` + "`" + `workspace` + "`" + ` ###### Arguments`,
				},
				resource.Attribute{
					Name:        "datacenter",
					Description: `(Required) (string)`,
				},
				resource.Attribute{
					Name:        "server",
					Description: `(Required) (string)`,
				},
				resource.Attribute{
					Name:        "default_datastore",
					Description: `(Optional) (string)`,
				},
				resource.Attribute{
					Name:        "folder",
					Description: `(Optional) (string)`,
				},
				resource.Attribute{
					Name:        "resourcepool_path",
					Description: `(Optional) (string) ##### ` + "`" + `disk` + "`" + ` ###### Arguments`,
				},
				resource.Attribute{
					Name:        "scsi_controller_type",
					Description: `(Optional) (string) ##### ` + "`" + `global` + "`" + ` ###### Arguments`,
				},
				resource.Attribute{
					Name:        "datacenters",
					Description: `(Optional/Computed) (string)`,
				},
				resource.Attribute{
					Name:        "datastore",
					Description: `(Optional) (string)`,
				},
				resource.Attribute{
					Name:        "insecure_flag",
					Description: `(Optional) (bool)`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional/Sensitive) (string)`,
				},
				resource.Attribute{
					Name:        "user",
					Description: `(Optional/Sensitive) (string)`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional) (string)`,
				},
				resource.Attribute{
					Name:        "soap_roundtrip_count",
					Description: `(Optional) (int)`,
				},
				resource.Attribute{
					Name:        "working_dir",
					Description: `(Optional) (string)`,
				},
				resource.Attribute{
					Name:        "vm_name",
					Description: `(Optional) (string)`,
				},
				resource.Attribute{
					Name:        "vm_uuid",
					Description: `(Optional) (string) ##### ` + "`" + `network` + "`" + ` ###### Arguments`,
				},
				resource.Attribute{
					Name:        "public_network",
					Description: `(Optional) (string) ### ` + "`" + `dns` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "nodelocal",
					Description: `(Optional) Nodelocal dns config (list Maxitem: 1)`,
				},
				resource.Attribute{
					Name:        "node_selector",
					Description: `(Optional) Node selector key pair (map)`,
				},
				resource.Attribute{
					Name:        "provider",
					Description: `(Optional) DNS provider. ` + "`" + `kube-dns` + "`" + `, ` + "`" + `coredns` + "`" + ` (default), and ` + "`" + `none` + "`" + ` are supported (string)`,
				},
				resource.Attribute{
					Name:        "reverse_cidrs",
					Description: `(Optional) Reverse CIDRs (list)`,
				},
				resource.Attribute{
					Name:        "upstream_nameservers",
					Description: `(Optional) Upstream nameservers (list) #### ` + "`" + `nodelocal` + "`" + ` ##### Arguments`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `(required) Nodelocal dns ip address (string)`,
				},
				resource.Attribute{
					Name:        "node_selector",
					Description: `(Optional) Node selector key pair (map) ### ` + "`" + `ingress` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "dns_policy",
					Description: `(Optional) Ingress controller DNS policy. ` + "`" + `ClusterFirstWithHostNet` + "`" + `, ` + "`" + `ClusterFirst` + "`" + `, ` + "`" + `Default` + "`" + `, and ` + "`" + `None` + "`" + ` are supported. [K8S dns Policy](https://kubernetes.io/docs/concepts/services-networking/dns-pod-service/#pod-s-dns-policy) (string)`,
				},
				resource.Attribute{
					Name:        "extra_args",
					Description: `(Optional) Extra arguments for the ingress controller (map)`,
				},
				resource.Attribute{
					Name:        "http_port",
					Description: `(Optional) Ingress controller http port (int)`,
				},
				resource.Attribute{
					Name:        "https_port",
					Description: `(Optional) Ingress controller https port (int)`,
				},
				resource.Attribute{
					Name:        "network_mode",
					Description: `(Optional) Networt mode for the ingress controller. ` + "`" + `hostNetwork` + "`" + `, ` + "`" + `hostPort` + "`" + ` and ` + "`" + `none` + "`" + ` are supported (string)`,
				},
				resource.Attribute{
					Name:        "node_selector",
					Description: `(Optional) Node selector key pair (map)`,
				},
				resource.Attribute{
					Name:        "options",
					Description: `(Optional) Ingress controller options (map)`,
				},
				resource.Attribute{
					Name:        "provider",
					Description: `(Optional) Ingress controller provider. ` + "`" + `nginx` + "`" + ` (default), and ` + "`" + `none` + "`" + ` are supported (string) ### ` + "`" + `monitoring` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "node_selector",
					Description: `(Optional) Node selector key pair (map)`,
				},
				resource.Attribute{
					Name:        "options",
					Description: `(Optional) Monitoring options (map)`,
				},
				resource.Attribute{
					Name:        "provider",
					Description: `(Optional/Computed) Monitoring provider (string) ### ` + "`" + `network` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "calico_network_provider",
					Description: `(Optional) Calico network provider config (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "canal_network_provider",
					Description: `(Optional) Canal network provider config (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "flannel_network_provider",
					Description: `(Optional) Flannel network provider config (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "weave_network_provider",
					Description: `(Optional) Weave network provider config (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "aci_network_provider",
					Description: `(Optional) Aci network provider config (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "mtu",
					Description: `(Optional) Network provider MTU. Default ` + "`" + `0` + "`" + ` (int)`,
				},
				resource.Attribute{
					Name:        "options",
					Description: `(Optional/Computed) Network provider options (map)`,
				},
				resource.Attribute{
					Name:        "plugin",
					Description: `(Optional) Network provider plugin. ` + "`" + `calico` + "`" + `, ` + "`" + `canal` + "`" + ` (default), ` + "`" + `flannel` + "`" + `, ` + "`" + `none` + "`" + ` and ` + "`" + `weave` + "`" + ` are supported. (string) #### ` + "`" + `calico_network_provider` + "`" + ` ##### Arguments`,
				},
				resource.Attribute{
					Name:        "cloud_provider",
					Description: `(Optional/Computed) Calico cloud provider (string) #### ` + "`" + `canal_network_provider` + "`" + ` ##### Arguments`,
				},
				resource.Attribute{
					Name:        "iface",
					Description: `(Optional/Computed) Canal network interface (string) #### ` + "`" + `flannel_network_provider` + "`" + ` ##### Arguments`,
				},
				resource.Attribute{
					Name:        "iface",
					Description: `(Optional/Computed) Flannel network interface (string) #### ` + "`" + `weave_network_provider` + "`" + ` ##### Arguments`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional/Computed) Weave password (string) #### ` + "`" + `aci_network_provider` + "`" + ` ##### Arguments`,
				},
				resource.Attribute{
					Name:        "system_id",
					Description: `(Required) Unique suffix for all cluster related objects in aci (string)`,
				},
				resource.Attribute{
					Name:        "apic_hosts",
					Description: `(Required) Ip address for apic hosts (list)`,
				},
				resource.Attribute{
					Name:        "token",
					Description: `(Required/Sensitive) UUID for this version of the input configuration (string)`,
				},
				resource.Attribute{
					Name:        "apic_user_name",
					Description: `(Required) User name for aci apic (string)`,
				},
				resource.Attribute{
					Name:        "apic_user_key",
					Description: `(Required/Sensitive) Base64 encoded private key for aci apic user (string)`,
				},
				resource.Attribute{
					Name:        "apic_user_crt",
					Description: `(Required/Sensitive) Base64 encoded certificate for aci apic user (string)`,
				},
				resource.Attribute{
					Name:        "encap_type",
					Description: `(Required) One of the supported encap types for aci(vlan/vxlan) (string)`,
				},
				resource.Attribute{
					Name:        "mcast_range_start",
					Description: `(Required) Mcast range start address for endpoint groups on aci (string)`,
				},
				resource.Attribute{
					Name:        "mcast_range_end",
					Description: `(Required) Mcast range end address for endpoint groups on aci (string)`,
				},
				resource.Attribute{
					Name:        "aep",
					Description: `(Required) Attachment entity profile name on aci (string)`,
				},
				resource.Attribute{
					Name:        "vrf_name",
					Description: `(Required) VRF Name on aci (string)`,
				},
				resource.Attribute{
					Name:        "vrf_tenant",
					Description: `(Required) Tenant for VRF on aci (string)`,
				},
				resource.Attribute{
					Name:        "l3out",
					Description: `(Required) L3Out on aci (string)`,
				},
				resource.Attribute{
					Name:        "node_subnet",
					Description: `(Required) Kubernetes node address subnet (string)`,
				},
				resource.Attribute{
					Name:        "l3out_external_networks",
					Description: `(Required) L3out external networks on aci (list)`,
				},
				resource.Attribute{
					Name:        "extern_dynamic",
					Description: `(Required) Subnet to use for dynamic external IPs on aci (string)`,
				},
				resource.Attribute{
					Name:        "extern_static",
					Description: `(Required) Subnet to use for static external IPs on aci (string)`,
				},
				resource.Attribute{
					Name:        "node_svc_subnet",
					Description: `(Required) Subnet to use for service graph endpoints on aci (string)`,
				},
				resource.Attribute{
					Name:        "kube_api_vlan",
					Description: `(Required) Vlan for node network on aci (string)`,
				},
				resource.Attribute{
					Name:        "service_vlan",
					Description: `(Required) Vlan for service graph nodes on aci (string)`,
				},
				resource.Attribute{
					Name:        "infra_vlan",
					Description: `(Required) Vlan for infra network on aci (string)`,
				},
				resource.Attribute{
					Name:        "snat_port_range_start",
					Description: `(Optional) Port start range for Source Network Address Translation on aci (string)`,
				},
				resource.Attribute{
					Name:        "snat_port_range_end",
					Description: `(Optional) Port end range for Source Network Address Translation on aci (string)`,
				},
				resource.Attribute{
					Name:        "snat_ports_per_node",
					Description: `(Optional) Ports per node for Source Network Address Translation on aci (string) ### ` + "`" + `nodes` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `(Required) Address ip for node (string)`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `(Required) Node roles in k8s cluster. ` + "`" + `controlplane` + "`" + `, ` + "`" + `etcd` + "`" + ` and ` + "`" + `worker` + "`" + ` are supported. (list)`,
				},
				resource.Attribute{
					Name:        "user",
					Description: `(Required/Sensitive) SSH user that will be used by RKE (string)`,
				},
				resource.Attribute{
					Name:        "docker_socket",
					Description: `(Optional) Docker socket on the node that will be used in tunneling (string)`,
				},
				resource.Attribute{
					Name:        "hostname_override",
					Description: `(Optional) Hostname override for node (string)`,
				},
				resource.Attribute{
					Name:        "internal_address",
					Description: `(Optional) Internal address that will be used for components communication (string)`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) Node labels (map)`,
				},
				resource.Attribute{
					Name:        "node_name",
					Description: `(Optional) Name of the host provisioned via docker machine (string)`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional) Port used for SSH communication (string)`,
				},
				resource.Attribute{
					Name:        "ssh_agent_auth",
					Description: `(Optional/Computed) SSH Agent Auth enable (bool)`,
				},
				resource.Attribute{
					Name:        "ssh_cert",
					Description: `(Optional/Sensitive) SSH Certificate (string)`,
				},
				resource.Attribute{
					Name:        "ssh_cert_path",
					Description: `(Optional) SSH Certificate path (string)`,
				},
				resource.Attribute{
					Name:        "ssh_key",
					Description: `(Optional/Sensitive) SSH Private Key (string)`,
				},
				resource.Attribute{
					Name:        "ssh_key_path",
					Description: `(Optional) SSH Private Key path (string)`,
				},
				resource.Attribute{
					Name:        "taints",
					Description: `(Optional) Node taints (list) #### ` + "`" + `taint` + "`" + ` ##### Arguments`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) Taint key (string)`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) Taint value (string)`,
				},
				resource.Attribute{
					Name:        "effect",
					Description: `(Optional) Taint effect. ` + "`" + `NoExecute` + "`" + `, ` + "`" + `NoSchedule` + "`" + ` (default) and ` + "`" + `PreferNoSchedule` + "`" + ` are supported (string) ### ` + "`" + `private_registries` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Required) Registry URL (string)`,
				},
				resource.Attribute{
					Name:        "is_default",
					Description: `(Optional) Set as default registry. Default ` + "`" + `false` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional/Sensitive) Registry password (string)`,
				},
				resource.Attribute{
					Name:        "user",
					Description: `(Optional/Sensitive) Registry user (string) ### ` + "`" + `restore` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "restore",
					Description: `(Optional) Restore cluster. Default ` + "`" + `false` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "snapshot_name",
					Description: `(Optional) Snapshot name (string) ### ` + "`" + `rotate_certificates` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "ca_certificates",
					Description: `(Optional) Rotate CA Certificates. Default ` + "`" + `false` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "services",
					Description: `(Optional) Services to rotate their certs. ` + "`" + `etcd` + "`" + `, ` + "`" + `kubelet` + "`" + `, ` + "`" + `kube-apiserver` + "`" + `, ` + "`" + `kube-proxy` + "`" + `, ` + "`" + `kube-scheduler` + "`" + ` and ` + "`" + `kube-controller-manager` + "`" + ` are supported (list) ### ` + "`" + `services` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "etcd",
					Description: `(Optional/Computed) Etcd options for RKE services (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "kube_api",
					Description: `(Optional/Computed) Kube API options for RKE services (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "kube_controller",
					Description: `(Optional/Computed) Kube Controller options for RKE services (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "kubelet",
					Description: `(Optional/Computed) Kubelet options for RKE services (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "kubeproxy",
					Description: `(Optional/Computed) Kubeproxy options for RKE services (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "scheduler",
					Description: `(Optional/Computed) Scheduler options for RKE services (list maxitems:1) #### ` + "`" + `etcd` + "`" + ` ##### Arguments`,
				},
				resource.Attribute{
					Name:        "backup_config",
					Description: `(Optional/Computed) Backup options for etcd service. Just for Rancher v2.2.x (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "ca_cert",
					Description: `(Optional/Computed/Sensitive) TLS CA certificate for etcd service (string)`,
				},
				resource.Attribute{
					Name:        "cert",
					Description: `(Optional/Computed/Sensitive) TLS certificate for etcd service (string)`,
				},
				resource.Attribute{
					Name:        "creation",
					Description: `(Optional/Computed) Creation option for etcd service (string)`,
				},
				resource.Attribute{
					Name:        "external_urls",
					Description: `(Optional/Computed) External urls for etcd service (list)`,
				},
				resource.Attribute{
					Name:        "extra_args",
					Description: `(Optional/Computed) Extra arguments for etcd service (map)`,
				},
				resource.Attribute{
					Name:        "extra_binds",
					Description: `(Optional/Computed) Extra binds for etcd service (list)`,
				},
				resource.Attribute{
					Name:        "extra_env",
					Description: `(Optional/Computed) Extra environment for etcd service (list)`,
				},
				resource.Attribute{
					Name:        "gid",
					Description: `(Optional) Etcd service GID. Default: ` + "`" + `0` + "`" + `. For Rancher v2.3.x or above (int)`,
				},
				resource.Attribute{
					Name:        "image",
					Description: `(Optional/Computed) Docker image for etcd service (string)`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Optional/Computed/Sensitive) TLS key for etcd service (string)`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Optional/Computed) Path for etcd service (string)`,
				},
				resource.Attribute{
					Name:        "retention",
					Description: `(Optional/Computed) Retention option for etcd service (string)`,
				},
				resource.Attribute{
					Name:        "snapshot",
					Description: `(Optional) Snapshot option for etcd service. Default ` + "`" + `true` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "uid",
					Description: `(Optional) Etcd service UID. Default: ` + "`" + `0` + "`" + `. For Rancher v2.3.x or above (int) ##### ` + "`" + `backup_config` + "`" + ` ###### Arguments`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) Enable etcd backup. Default ` + "`" + `true` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "interval_hours",
					Description: `(Optional) Interval hours for etcd backup. Default ` + "`" + `12` + "`" + ` (int)`,
				},
				resource.Attribute{
					Name:        "retention",
					Description: `(Optional) Retention for etcd backup. Default ` + "`" + `6` + "`" + ` (int)`,
				},
				resource.Attribute{
					Name:        "s3_backup_config",
					Description: `(Optional) S3 config options for etcd backup (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "safe_timestamp",
					Description: `(Optional) Safe timestamp for etcd backup. Default: ` + "`" + `false` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `(Optional/Computed) Timeout in seconds for etcd backup. Default: ` + "`" + `300` + "`" + `. Just for RKE v1.2.6 and above (int) ###### ` + "`" + `s3_backup_config` + "`" + ` ###### Arguments`,
				},
				resource.Attribute{
					Name:        "access_key",
					Description: `(Optional/Sensitive) Access key for S3 service (string)`,
				},
				resource.Attribute{
					Name:        "bucket_name",
					Description: `(Optional) Bucket name for S3 service (string)`,
				},
				resource.Attribute{
					Name:        "custom_ca",
					Description: `(Optional/Sensitive) Base64 encoded custom CA for S3 service. Use filebase64(<FILE>) for encoding file. Available from Rancher v2.2.5 (string)`,
				},
				resource.Attribute{
					Name:        "endpoint",
					Description: `(Optional) Endpoint for S3 service (string)`,
				},
				resource.Attribute{
					Name:        "folder",
					Description: `(Optional) Folder for S3 service. Available from Rancher v2.2.7 (string)`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) Region for S3 service (string)`,
				},
				resource.Attribute{
					Name:        "secret_key",
					Description: `(Optional/Sensitive) Secret key for S3 service (string) #### ` + "`" + `kube_api` + "`" + ` ##### Arguments`,
				},
				resource.Attribute{
					Name:        "always_pull_images",
					Description: `(Optional/Computed) Enable [AlwaysPullImages](https://kubernetes.io/docs/reference/access-authn-authz/admission-controllers/#alwayspullimages) Admission controller plugin. [Rancher docs](https://rancher.com/docs/rke/latest/en/config-options/services/#kubernetes-api-server-options) (bool)`,
				},
				resource.Attribute{
					Name:        "audit_log",
					Description: `(Optional/Computed) K8s audit log configuration. (list maxitem: 1)`,
				},
				resource.Attribute{
					Name:        "event_rate_limit",
					Description: `(Optional) K8s event rate limit configuration. (list maxitem: 1)`,
				},
				resource.Attribute{
					Name:        "extra_args",
					Description: `(Optional/Computed) Extra arguments for kube API service (map)`,
				},
				resource.Attribute{
					Name:        "extra_binds",
					Description: `(Optional/Computed) Extra binds for kube API service (list)`,
				},
				resource.Attribute{
					Name:        "extra_env",
					Description: `(Optional/Computed) Extra environment for kube API service (list)`,
				},
				resource.Attribute{
					Name:        "image",
					Description: `(Optional/Computed) Docker image for kube API service (string)`,
				},
				resource.Attribute{
					Name:        "pod_security_policy",
					Description: `(Optional/Computed) Pod Security Policy option for kube API service (bool)`,
				},
				resource.Attribute{
					Name:        "secrets_encryption_config",
					Description: `(Optional) [Encrypt k8s secret data configration](https://rancher.com/docs/rke/latest/en/config-options/secrets-encryption/). (list maxitem: 1)`,
				},
				resource.Attribute{
					Name:        "service_cluster_ip_range",
					Description: `(Optional/Computed) Service Cluster IP Range option for kube API service (string)`,
				},
				resource.Attribute{
					Name:        "service_node_port_range",
					Description: `(Optional/Computed) Service Node Port Range option for kube API service (string) ##### ` + "`" + `audit_log` + "`" + ` ###### Arguments`,
				},
				resource.Attribute{
					Name:        "configuration",
					Description: `(Optional/Computed) Audit log configuration. (list maxtiem: 1)`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional/Computed) Enable audit log (bool) ###### ` + "`" + `configuration` + "`" + ` ###### Arguments`,
				},
				resource.Attribute{
					Name:        "format",
					Description: `(Optional/Computed) Audit log format (string)`,
				},
				resource.Attribute{
					Name:        "max_age",
					Description: `(Optional/Computed) Audit log max age (int)`,
				},
				resource.Attribute{
					Name:        "max_backup",
					Description: `(Optional/Computed) Audit log max backup. Default: ` + "`" + `10` + "`" + ` (int)`,
				},
				resource.Attribute{
					Name:        "max_size",
					Description: `(Optional/Computed) Audit log max size. Default: ` + "`" + `100` + "`" + ` (int)`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Optional/Computed) Audit log path. Default: ` + "`" + `/var/log/kube-audit/audit-log.json` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "policy",
					Description: `(Optional/Computed) Audit policy json encoded definition. ` + "`" + `"apiVersion"` + "`" + ` and ` + "`" + `"kind":"Policy","rules"` + "`" + ` fields are required in the json. Ex. ` + "`" + `jsonencode({"apiVersion":"audit.k8s.io/v1","kind":"Policy","rules":[{"level":"RequestResponse","resources":[{"group":"","resources":["pods"]}]}]})` + "`" + ` [More info](https://rancher.com/docs/rke/latest/en/config-options/audit-log/) (string) ##### ` + "`" + `event_rate_limit` + "`" + ` ###### Arguments`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional/Computed) Enable event rate limit (bool)`,
				},
				resource.Attribute{
					Name:        "configuration",
					Description: `(Optional) Event rate limit yaml encoded configuration. ` + "`" + `"apiVersion"` + "`" + ` and ` + "`" + `"kind":"Configuration"` + "`" + ` fields are required in the yaml. Ex. ` + "`" + `apiVersion: eventratelimit.admission.k8s.io/v1alpha1\nkind: Configuration\nlimits:\n- type: Server\n burst: 30000\n qps: 6000\n` + "`" + ` [More info](https://rancher.com/docs/rke/latest/en/config-options/rate-limiting/) (string) ##### ` + "`" + `secrets_encryption_config` + "`" + ` ###### Arguments`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional/Computed) Enable secrets encryption (bool)`,
				},
				resource.Attribute{
					Name:        "custom_config",
					Description: `(Optional) Secrets encryption yaml encoded custom configuration. ` + "`" + `"apiVersion"` + "`" + ` and ` + "`" + `"kind":"EncryptionConfiguration"` + "`" + ` fields are required in the yaml. Ex. ` + "`" + `apiVersion: apiserver.config.k8s.io/v1\nkind: EncryptionConfiguration\nresources:\n- resources:\n - secrets\n providers:\n - aescbc:\n keys:\n - name: k-fw5hn\n secret: RTczRjFDODMwQzAyMDVBREU4NDJBMUZFNDhCNzM5N0I=\n identity: {}\n` + "`" + ` [More info](https://rancher.com/docs/rke/latest/en/config-options/secrets-encryption/) (string) #### ` + "`" + `kube_controller` + "`" + ` ##### Arguments`,
				},
				resource.Attribute{
					Name:        "cluster_cidr",
					Description: `(Optional/Computed) Cluster CIDR option for kube controller service (string)`,
				},
				resource.Attribute{
					Name:        "extra_args",
					Description: `(Optional/Computed) Extra arguments for kube controller service (map)`,
				},
				resource.Attribute{
					Name:        "extra_binds",
					Description: `(Optional/Computed) Extra binds for kube controller service (list)`,
				},
				resource.Attribute{
					Name:        "extra_env",
					Description: `(Optional/Computed) Extra environment for kube controller service (list)`,
				},
				resource.Attribute{
					Name:        "image",
					Description: `(Optional/Computed) Docker image for kube controller service (string)`,
				},
				resource.Attribute{
					Name:        "service_cluster_ip_range",
					Description: `(Optional/Computed) Service Cluster ip Range option for kube controller service (string) #### ` + "`" + `kubelet` + "`" + ` ##### Arguments`,
				},
				resource.Attribute{
					Name:        "cluster_dns_server",
					Description: `(Optional/Computed) Cluster DNS Server option for kubelet service (string)`,
				},
				resource.Attribute{
					Name:        "cluster_domain",
					Description: `(Optional) Cluster Domain option for kubelet service. Default ` + "`" + `cluster.local` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "extra_args",
					Description: `(Optional/Computed) Extra arguments for kubelet service (map)`,
				},
				resource.Attribute{
					Name:        "extra_binds",
					Description: `(Optional/Computed) Extra binds for kubelet service (list)`,
				},
				resource.Attribute{
					Name:        "extra_env",
					Description: `(Optional/Computed) Extra environment for kubelet service (list)`,
				},
				resource.Attribute{
					Name:        "fail_swap_on",
					Description: `(Optional/Computed) Enable or disable failing when swap on is not supported (bool)`,
				},
				resource.Attribute{
					Name:        "image",
					Description: `(Optional/Computed) Docker image for kubelet service (string)`,
				},
				resource.Attribute{
					Name:        "infra_container_image",
					Description: `(Optional/Computed) Infra container image for kubelet service (string) #### ` + "`" + `kubeproxy` + "`" + ` ##### Arguments`,
				},
				resource.Attribute{
					Name:        "extra_args",
					Description: `(Optional/Computed) Extra arguments for kubeproxy service (map)`,
				},
				resource.Attribute{
					Name:        "extra_binds",
					Description: `(Optional/Computed) Extra binds for kubeproxy service (list)`,
				},
				resource.Attribute{
					Name:        "extra_env",
					Description: `(Optional/Computed) Extra environment for kubeproxy service (list)`,
				},
				resource.Attribute{
					Name:        "image",
					Description: `(Optional/Computed) Docker image for kubeproxy service (string) #### ` + "`" + `scheduler` + "`" + ` ##### Arguments`,
				},
				resource.Attribute{
					Name:        "extra_args",
					Description: `(Optional/Computed) Extra arguments for scheduler service (map)`,
				},
				resource.Attribute{
					Name:        "extra_binds",
					Description: `(Optional/Computed) Extra binds for scheduler service (list)`,
				},
				resource.Attribute{
					Name:        "extra_env",
					Description: `(Optional/Computed) Extra environment for scheduler service (list)`,
				},
				resource.Attribute{
					Name:        "image",
					Description: `(Optional/Computed) Docker image for scheduler service (string) ### ` + "`" + `system_images` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "etcd",
					Description: `(Optional) Docker image for etcd (string)`,
				},
				resource.Attribute{
					Name:        "alpine",
					Description: `(Optional) Docker image for alpine (string)`,
				},
				resource.Attribute{
					Name:        "nginx_proxy",
					Description: `(Optional) Docker image for nginx_proxy (string)`,
				},
				resource.Attribute{
					Name:        "cert_downloader",
					Description: `(Optional) Docker image for cert_downloader (string)`,
				},
				resource.Attribute{
					Name:        "kubernetes_services_sidecar",
					Description: `(Optional) Docker image for kubernetes_services_sidecar (string)`,
				},
				resource.Attribute{
					Name:        "kube_dns",
					Description: `(Optional) Docker image for kube_dns (string)`,
				},
				resource.Attribute{
					Name:        "dnsmasq",
					Description: `(Optional) Docker image for dnsmasq (string)`,
				},
				resource.Attribute{
					Name:        "kube_dns_sidecar",
					Description: `(Optional) Docker image for kube_dns_sidecar (string)`,
				},
				resource.Attribute{
					Name:        "kube_dns_autoscaler",
					Description: `(Optional) Docker image for kube_dns_autoscaler (string)`,
				},
				resource.Attribute{
					Name:        "coredns",
					Description: `(Optional) Docker image for coredns (string)`,
				},
				resource.Attribute{
					Name:        "coredns_autoscaler",
					Description: `(Optional) Docker image for coredns_autoscaler (string)`,
				},
				resource.Attribute{
					Name:        "nodelocal",
					Description: `(Optional) Docker image for nodelocal (string)`,
				},
				resource.Attribute{
					Name:        "kubernetes",
					Description: `(Optional) Docker image for kubernetes (string)`,
				},
				resource.Attribute{
					Name:        "flannel",
					Description: `(Optional) Docker image for flannel (string)`,
				},
				resource.Attribute{
					Name:        "flannel_cni",
					Description: `(Optional) Docker image for flannel_cni (string)`,
				},
				resource.Attribute{
					Name:        "calico_node",
					Description: `(Optional) Docker image for calico_node (string)`,
				},
				resource.Attribute{
					Name:        "calico_cni",
					Description: `(Optional) Docker image for calico_cni (string)`,
				},
				resource.Attribute{
					Name:        "calico_controllers",
					Description: `(Optional) Docker image for calico_controllers (string)`,
				},
				resource.Attribute{
					Name:        "calico_ctl",
					Description: `(Optional) Docker image for calico_ctl (string)`,
				},
				resource.Attribute{
					Name:        "calico_flex_vol",
					Description: `(Optional) Docker image for calico_flex_vol (string)`,
				},
				resource.Attribute{
					Name:        "canal_node",
					Description: `(Optional) Docker image for canal_node (string)`,
				},
				resource.Attribute{
					Name:        "canal_cni",
					Description: `(Optional) Docker image for canal_cni (string)`,
				},
				resource.Attribute{
					Name:        "canal_flannel",
					Description: `(Optional) Docker image for canal_flannel (string)`,
				},
				resource.Attribute{
					Name:        "canal_flex_vol",
					Description: `(Optional) Docker image for canal_flex_vol (string)`,
				},
				resource.Attribute{
					Name:        "weave_node",
					Description: `(Optional) Docker image for weave_node (string)`,
				},
				resource.Attribute{
					Name:        "weave_cni",
					Description: `(Optional) Docker image for weave_cni (string)`,
				},
				resource.Attribute{
					Name:        "pod_infra_container",
					Description: `(Optional) Docker image for pod_infra_container (string)`,
				},
				resource.Attribute{
					Name:        "ingress",
					Description: `(Optional) Docker image for ingress (string)`,
				},
				resource.Attribute{
					Name:        "ingress_backend",
					Description: `(Optional) Docker image for ingress_backend (string)`,
				},
				resource.Attribute{
					Name:        "metrics_server",
					Description: `(Optional) Docker image for metrics_server (string)`,
				},
				resource.Attribute{
					Name:        "windows_pod_infra_container",
					Description: `(Optional) Docker image for windows_pod_infra_container (string)`,
				},
				resource.Attribute{
					Name:        "aci_cni_deploy_container",
					Description: `(Optional) Docker image for aci_cni_deploy_container (string)`,
				},
				resource.Attribute{
					Name:        "aci_host_container",
					Description: `(Optional) Docker image for aci_host_container (string)`,
				},
				resource.Attribute{
					Name:        "aci_opflex_container",
					Description: `(Optional) Docker image for aci_opflex_container (string)`,
				},
				resource.Attribute{
					Name:        "aci_mcast_container",
					Description: `(Optional) Docker image for aci_mcast_container (string)`,
				},
				resource.Attribute{
					Name:        "aci_ovs_container",
					Description: `(Optional) Docker image for aci_ovs_container (string)`,
				},
				resource.Attribute{
					Name:        "aci_controller_container",
					Description: `(Optional) Docker image for aci_controller_container (string) ### ` + "`" + `upgrade_strategy` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "drain",
					Description: `(Optional/Computed) RKE drain nodes (bool)`,
				},
				resource.Attribute{
					Name:        "drain_input",
					Description: `(Optional/Computed) RKE drain node input (list Maxitems: 1)`,
				},
				resource.Attribute{
					Name:        "max_unavailable_controlplane",
					Description: `(Optional/Computed) RKE max unavailable controlplane nodes (string)`,
				},
				resource.Attribute{
					Name:        "max_unavailable_worker",
					Description: `(Optional/Computed) RKE max unavailable worker nodes (string) #### ` + "`" + `drain_input` + "`" + ` ##### Arguments`,
				},
				resource.Attribute{
					Name:        "delete_local_data",
					Description: `(Optional/Computed) Delete RKE node local data (bool)`,
				},
				resource.Attribute{
					Name:        "force",
					Description: `(Optional/Computed) Force RKE node drain (bool)`,
				},
				resource.Attribute{
					Name:        "grace_period",
					Description: `(Optional/Computed) RKE node drain grace period (int)`,
				},
				resource.Attribute{
					Name:        "ignore_daemon_sets",
					Description: `(Optional/Computed) Ignore RKE daemon sets (bool)`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `(Optional/Computed) RKE node drain timeout (int) ## Timeouts ` + "`" + `rke_cluster` + "`" + ` provides the following [Timeouts](https://www.terraform.io/docs/configuration/resources.html#operation-timeouts) configuration options: - ` + "`" + `create` + "`" + ` - (Default ` + "`" + `30 minutes` + "`" + `) Used for creating clusters. - ` + "`" + `update` + "`" + ` - (Default ` + "`" + `30 minutes` + "`" + `) Used for cluster modifications. - ` + "`" + `delete` + "`" + ` - (Default ` + "`" + `30 minutes` + "`" + `) Used for deleting clusters. ## Import rke_cluster can be imported using the RKE cluster config and state files as ID in the format ` + "`" + `<cluster_config_file>:<rke_state_file>` + "`" + ` ` + "`" + `` + "`" + `` + "`" + ` $ terraform import rke_cluster.foo &lt;cluster_config_file&gt;:&lt;rke_state_file&gt; ` + "`" + `` + "`" + `` + "`" + ` As experimental feature, dind rke_cluster can be also imported adding ` + "`" + `dind` + "`" + ` as 3rd import parameter ` + "`" + `<cluster_config_file>:<rke_state_file>:dind` + "`" + ` ` + "`" + `` + "`" + `` + "`" + ` $ terraform import rke_cluster.foo &lt;cluster_config_file&gt;:&lt;rke_state_file&gt;:dind ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Computed) The ID of the resource (string)`,
				},
				resource.Attribute{
					Name:        "ca_crt",
					Description: `(Computed/Sensitive) RKE k8s cluster CA certificate (string)`,
				},
				resource.Attribute{
					Name:        "client_cert",
					Description: `(Computed/Sensitive) RKE k8s cluster client certificate (string)`,
				},
				resource.Attribute{
					Name:        "client_key",
					Description: `(Computed/Sensitive) RKE k8s cluster client key (string)`,
				},
				resource.Attribute{
					Name:        "rke_state",
					Description: `(Computed/Sensitive) RKE k8s cluster state (string)`,
				},
				resource.Attribute{
					Name:        "kube_config_yaml",
					Description: `(Computed/Sensitive) RKE k8s cluster kube config yaml (string)`,
				},
				resource.Attribute{
					Name:        "internal_kube_config_yaml",
					Description: `(Computed/Sensitive) RKE k8s cluster internal kube config yaml (string)`,
				},
				resource.Attribute{
					Name:        "rke_cluster_yaml",
					Description: `(Computed/Sensitive) RKE k8s cluster config yaml (string)`,
				},
				resource.Attribute{
					Name:        "certificates",
					Description: `(Computed/Sensitive) RKE k8s cluster certificates (string)`,
				},
				resource.Attribute{
					Name:        "kube_admin_user",
					Description: `(Computed) RKE k8s cluster admin user (string)`,
				},
				resource.Attribute{
					Name:        "api_server_url",
					Description: `(Computed) RKE k8s cluster api server url (string)`,
				},
				resource.Attribute{
					Name:        "cluster_domain",
					Description: `(Computed) RKE k8s cluster domain (string)`,
				},
				resource.Attribute{
					Name:        "cluster_cidr",
					Description: `(Computed) RKE k8s cluster cidr (string)`,
				},
				resource.Attribute{
					Name:        "cluster_dns_server",
					Description: `(Computed) RKE k8s cluster dns server (string)`,
				},
				resource.Attribute{
					Name:        "control_plane_hosts",
					Description: `(Computed) RKE k8s cluster control plane nodes (list)`,
				},
				resource.Attribute{
					Name:        "etcd_hosts",
					Description: `(Computed) RKE k8s cluster etcd nodes (list)`,
				},
				resource.Attribute{
					Name:        "inactive_hosts",
					Description: `(Computed) RKE k8s cluster inactive nodes (list)`,
				},
				resource.Attribute{
					Name:        "worker_hosts",
					Description: `(Computed) RKE k8s cluster worker nodes (list)`,
				},
				resource.Attribute{
					Name:        "running_system_images",
					Description: `(Computed) RKE k8s cluster running system images list (list) ## Nested blocks ### ` + "`" + `authentication` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "sans",
					Description: `(Optional/Computed) List of additional hostnames and IPs to include in the api server PKI cert (list)`,
				},
				resource.Attribute{
					Name:        "strategy",
					Description: `(Optional) Authentication strategy that will be used in RKE k8s cluster. Default: ` + "`" + `x509` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "webhook",
					Description: `(Optional/Computed) Webhook configuration options (list maxitem: 1) #### ` + "`" + `webhook` + "`" + ` ##### Arguments`,
				},
				resource.Attribute{
					Name:        "config_file",
					Description: `(Optional) Multiline string that represent a custom webhook config file (string)`,
				},
				resource.Attribute{
					Name:        "cache_timeout",
					Description: `(Optional) Controls how long to cache authentication decisions (string) ### ` + "`" + `authorization` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `(Optional) RKE mode for authorization. ` + "`" + `rbac` + "`" + ` and ` + "`" + `none` + "`" + ` modes are available. Default ` + "`" + `rbac` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "options",
					Description: `(Optional/Computed) RKE options for authorization (map) ### ` + "`" + `bastion_host` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `(Required) Address of Bastion Host (string)`,
				},
				resource.Attribute{
					Name:        "user",
					Description: `(Required) SSH User to Bastion Host (string)`,
				},
				resource.Attribute{
					Name:        "ignore_proxy_env_vars",
					Description: `(Optional) Ignore proxy env vars at Bastion Host? Default: ` + "`" + `false` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional) SSH Port of Bastion Host. Default ` + "`" + `22` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "ssh_agent_auth",
					Description: `(Optional/Computed) SSH Agent Auth enable (bool)`,
				},
				resource.Attribute{
					Name:        "ssh_cert",
					Description: `(Optional/Sensitive) SSH Certificate Key (string)`,
				},
				resource.Attribute{
					Name:        "ssh_cert_path",
					Description: `(Optional) SSH Certificate Key Path (string)`,
				},
				resource.Attribute{
					Name:        "ssh_key",
					Description: `(Optional/Sensitive) SSH Private Key (string)`,
				},
				resource.Attribute{
					Name:        "ssh_key_path",
					Description: `(Optional) SSH Private Key Path (string) ### ` + "`" + `cloud_provider` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Cloud Provider name. ` + "`" + `aws` + "`" + `, ` + "`" + `azure` + "`" + `, ` + "`" + `custom` + "`" + `, ` + "`" + `external` + "`" + `, ` + "`" + `openstack` + "`" + `, ` + "`" + `vsphere` + "`" + ` are supported (string)`,
				},
				resource.Attribute{
					Name:        "aws_cloud_config",
					Description: `(DEPRECATED) Use aws_cloud_provider instead`,
				},
				resource.Attribute{
					Name:        "aws_cloud_provider",
					Description: `(Optional) AWS Cloud Provider config [rke-aws-cloud-provider](https://rancher.com/docs/rke/latest/en/config-options/cloud-providers/aws/) (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "azure_cloud_config",
					Description: `(DEPRECATED) Use azure_cloud_provider instead`,
				},
				resource.Attribute{
					Name:        "azure_cloud_provider",
					Description: `(Optional) Azure Cloud Provider config [rke-azure-cloud-provider](https://rancher.com/docs/rke/latest/en/config-options/cloud-providers/azure/) (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "custom_cloud_config",
					Description: `(DEPRECATED) Use custom_cloud_provider instead`,
				},
				resource.Attribute{
					Name:        "custom_cloud_provider",
					Description: `(Optional) Custom Cloud Provider config (string)`,
				},
				resource.Attribute{
					Name:        "openstack_cloud_config",
					Description: `(DEPRECATED) Use openstack_cloud_provider instead`,
				},
				resource.Attribute{
					Name:        "openstack_cloud_provider",
					Description: `(Optional/Computed) Openstack Cloud Provider config [rke-openstack-cloud-provider](https://rancher.com/docs/rke/latest/en/config-options/cloud-providers/openstack/) (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "vsphere_cloud_config",
					Description: `(DEPRECATED) Use vsphere_cloud_provider instead`,
				},
				resource.Attribute{
					Name:        "vsphere_cloud_provider",
					Description: `(Optional/Computed) Vsphere Cloud Provider config [rke-vsphere-cloud-provider](https://rancher.com/docs/rke/latest/en/config-options/cloud-providers/vsphere/) Extra argument ` + "`" + `name` + "`" + ` is required on ` + "`" + `virtual_center` + "`" + ` configuration. (list maxitems:1) #### ` + "`" + `aws_cloud_provider` + "`" + ` ##### Arguments`,
				},
				resource.Attribute{
					Name:        "global",
					Description: `(Optional) (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "service_override",
					Description: `(Optional) (list) ##### ` + "`" + `global` + "`" + ` ###### Arguments`,
				},
				resource.Attribute{
					Name:        "disable_security_group_ingress",
					Description: `(Optional) Disables the automatic ingress creation. Default ` + "`" + `false` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "disable_strict_zone_check",
					Description: `(Optional) Setting this to true will disable the check and provide a warning that the check was skipped. Default ` + "`" + `false` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "elb_security_group",
					Description: `(Optional) Use these ELB security groups instead create new (string)`,
				},
				resource.Attribute{
					Name:        "kubernetes_cluster_id",
					Description: `(Optional) The cluster id we'll use to identify our cluster resources (string)`,
				},
				resource.Attribute{
					Name:        "kubernetes_cluster_tag",
					Description: `(Optional) Legacy cluster id we'll use to identify our cluster resources (string)`,
				},
				resource.Attribute{
					Name:        "role_arn",
					Description: `(Optional) IAM role to assume when interaction with AWS APIs (string)`,
				},
				resource.Attribute{
					Name:        "route_table_id",
					Description: `(Optional) Enables using a specific RouteTable (string)`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Optional) Enables using a specific subnet to use for ELB's (string)`,
				},
				resource.Attribute{
					Name:        "vpc",
					Description: `(Optional) The AWS VPC flag enables the possibility to run the master components on a different aws account, on a different cloud provider or on-premises. If the flag is set also the KubernetesClusterTag must be provided (string)`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `(Optional) The AWS zone (string) ##### ` + "`" + `service_override` + "`" + ` ###### Arguments`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(DEPRECATED) Use service instead`,
				},
				resource.Attribute{
					Name:        "service",
					Description: `(Required) (string)`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) (string)`,
				},
				resource.Attribute{
					Name:        "signing_method",
					Description: `(Optional/Computed) (string)`,
				},
				resource.Attribute{
					Name:        "signing_name",
					Description: `(Optional) (string)`,
				},
				resource.Attribute{
					Name:        "signing_region",
					Description: `(Optional) (string)`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Optional) (string) #### ` + "`" + `azure_cloud_provider` + "`" + ` ##### Arguments`,
				},
				resource.Attribute{
					Name:        "aad_client_id",
					Description: `(Required/Sensitive) (string)`,
				},
				resource.Attribute{
					Name:        "aad_client_secret",
					Description: `(Required/Sensitive) (string)`,
				},
				resource.Attribute{
					Name:        "subscription_id",
					Description: `(Required/Sensitive) (string)`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Required/Sensitive) (string)`,
				},
				resource.Attribute{
					Name:        "aad_client_cert_password",
					Description: `(Optional/Computed/Sensitive) (string)`,
				},
				resource.Attribute{
					Name:        "aad_client_cert_path",
					Description: `(Optional/Computed) (string)`,
				},
				resource.Attribute{
					Name:        "cloud",
					Description: `(Optional/Computed) (string)`,
				},
				resource.Attribute{
					Name:        "cloud_provider_backoff",
					Description: `(Optional/Computed) (bool)`,
				},
				resource.Attribute{
					Name:        "cloud_provider_backoff_duration",
					Description: `(Optional/Computed) (int)`,
				},
				resource.Attribute{
					Name:        "cloud_provider_backoff_exponent",
					Description: `(Optional/Computed) (int)`,
				},
				resource.Attribute{
					Name:        "cloud_provider_backoff_jitter",
					Description: `(Optional/Computed) (int)`,
				},
				resource.Attribute{
					Name:        "cloud_provider_backoff_retries",
					Description: `(Optional/Computed) (int)`,
				},
				resource.Attribute{
					Name:        "cloud_provider_rate_limit",
					Description: `(Optional/Computed) (bool)`,
				},
				resource.Attribute{
					Name:        "cloud_provider_rate_limit_bucket",
					Description: `(Optional/Computed) (int)`,
				},
				resource.Attribute{
					Name:        "cloud_provider_rate_limit_qps",
					Description: `(Optional/Computed) (int)`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `(Optional/Computed) (string)`,
				},
				resource.Attribute{
					Name:        "maximum_load_balancer_rule_count",
					Description: `(Optional/Computed) (int)`,
				},
				resource.Attribute{
					Name:        "primary_availability_set_name",
					Description: `(Optional/Computed) (string)`,
				},
				resource.Attribute{
					Name:        "primary_scale_set_name",
					Description: `(Optional/Computed) (string)`,
				},
				resource.Attribute{
					Name:        "resource_group",
					Description: `(Optional/Computed) (string)`,
				},
				resource.Attribute{
					Name:        "route_table_name",
					Description: `(Optional/Computed) (string)`,
				},
				resource.Attribute{
					Name:        "security_group_name",
					Description: `(Optional/Computed) (string)`,
				},
				resource.Attribute{
					Name:        "subnet_name",
					Description: `(Optional/Computed) (string)`,
				},
				resource.Attribute{
					Name:        "use_instance_metadata",
					Description: `(Optional/Computed) (bool)`,
				},
				resource.Attribute{
					Name:        "use_managed_identity_extension",
					Description: `(Optional/Computed) (bool)`,
				},
				resource.Attribute{
					Name:        "vm_type",
					Description: `(Optional/Computed) (string)`,
				},
				resource.Attribute{
					Name:        "vnet_name",
					Description: `(Optional/Computed) (string)`,
				},
				resource.Attribute{
					Name:        "vnet_resource_group",
					Description: `(Optional/Computed) (string) #### ` + "`" + `openstack_cloud_provider` + "`" + ` ##### Arguments`,
				},
				resource.Attribute{
					Name:        "global",
					Description: `(Required) (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "block_storage",
					Description: `(Optional/Computed) (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "load_balancer",
					Description: `(Optional/Computed) (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `(Optional/Computed) (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "route",
					Description: `(Optional/Computed) (list maxitems:1) ##### ` + "`" + `global` + "`" + ` ###### Arguments`,
				},
				resource.Attribute{
					Name:        "auth_url",
					Description: `(Required) (string)`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Required/Sensitive) (string)`,
				},
				resource.Attribute{
					Name:        "ca_file",
					Description: `(Optional) (string)`,
				},
				resource.Attribute{
					Name:        "domain_id",
					Description: `(Optional/Sensitive) Required if ` + "`" + `domain_name` + "`" + ` not provided. (string)`,
				},
				resource.Attribute{
					Name:        "domain_name",
					Description: `(Optional) Required if ` + "`" + `domain_id` + "`" + ` not provided. (string)`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) (string)`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Optional/Sensitive) Required if ` + "`" + `tenant_name` + "`" + ` not provided. (string)`,
				},
				resource.Attribute{
					Name:        "tenant_name",
					Description: `(Optional) Required if ` + "`" + `tenant_id` + "`" + ` not provided. (string)`,
				},
				resource.Attribute{
					Name:        "trust_id",
					Description: `(Optional/Sensitive) (string)`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Optional) Required if ` + "`" + `user_id` + "`" + ` not provided. (string)`,
				},
				resource.Attribute{
					Name:        "user_id",
					Description: `(Optional/Sensitive) Required if ` + "`" + `username` + "`" + ` not provided. (string) ##### ` + "`" + `block_storage` + "`" + ` ###### Arguments`,
				},
				resource.Attribute{
					Name:        "bs_version",
					Description: `(Optional) (string)`,
				},
				resource.Attribute{
					Name:        "ignore_volume_az",
					Description: `(Optional) (string)`,
				},
				resource.Attribute{
					Name:        "trust_device_path",
					Description: `(Optional) (string) ##### ` + "`" + `load_balancer` + "`" + ` ###### Arguments`,
				},
				resource.Attribute{
					Name:        "create_monitor",
					Description: `(Optional) (bool)`,
				},
				resource.Attribute{
					Name:        "floating_network_id",
					Description: `(Optional) (string)`,
				},
				resource.Attribute{
					Name:        "lb_method",
					Description: `(Optional) (string)`,
				},
				resource.Attribute{
					Name:        "lb_provider",
					Description: `(Optional) (string)`,
				},
				resource.Attribute{
					Name:        "lb_version",
					Description: `(Optional) (string)`,
				},
				resource.Attribute{
					Name:        "manage_security_groups",
					Description: `(Optional) (bool)`,
				},
				resource.Attribute{
					Name:        "monitor_delay",
					Description: `(Optional) (string)`,
				},
				resource.Attribute{
					Name:        "monitor_max_retries",
					Description: `(Optional) (int)`,
				},
				resource.Attribute{
					Name:        "monitor_timeout",
					Description: `(Optional) (string)`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Optional) (string)`,
				},
				resource.Attribute{
					Name:        "use_octavia",
					Description: `(Optional) (bool) ##### ` + "`" + `metadata` + "`" + ` ###### Arguments`,
				},
				resource.Attribute{
					Name:        "request_timeout",
					Description: `(Optional) (int)`,
				},
				resource.Attribute{
					Name:        "search_order",
					Description: `(Optional) (string) ##### ` + "`" + `route` + "`" + ` ###### Arguments`,
				},
				resource.Attribute{
					Name:        "router_id",
					Description: `(Optional) (string) #### ` + "`" + `vsphere_cloud_provider` + "`" + ` ##### Arguments`,
				},
				resource.Attribute{
					Name:        "virtual_center",
					Description: `(Required) (List)`,
				},
				resource.Attribute{
					Name:        "workspace",
					Description: `(Required) (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "disk",
					Description: `(Optional/Computed) (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "global",
					Description: `(Optional/Computed) (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "network",
					Description: `(Optional/Computed) (list maxitems:1) ##### ` + "`" + `virtual_center` + "`" + ` ###### Arguments`,
				},
				resource.Attribute{
					Name:        "datacenters",
					Description: `(Required) (string)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of virtualcenter config for Vsphere Cloud Provider config (string)`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Required/Sensitive) (string)`,
				},
				resource.Attribute{
					Name:        "user",
					Description: `(Required/Sensitive) (string)`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional) (string)`,
				},
				resource.Attribute{
					Name:        "soap_roundtrip_count",
					Description: `(Optional) (int) ##### ` + "`" + `workspace` + "`" + ` ###### Arguments`,
				},
				resource.Attribute{
					Name:        "datacenter",
					Description: `(Required) (string)`,
				},
				resource.Attribute{
					Name:        "server",
					Description: `(Required) (string)`,
				},
				resource.Attribute{
					Name:        "default_datastore",
					Description: `(Optional) (string)`,
				},
				resource.Attribute{
					Name:        "folder",
					Description: `(Optional) (string)`,
				},
				resource.Attribute{
					Name:        "resourcepool_path",
					Description: `(Optional) (string) ##### ` + "`" + `disk` + "`" + ` ###### Arguments`,
				},
				resource.Attribute{
					Name:        "scsi_controller_type",
					Description: `(Optional) (string) ##### ` + "`" + `global` + "`" + ` ###### Arguments`,
				},
				resource.Attribute{
					Name:        "datacenters",
					Description: `(Optional/Computed) (string)`,
				},
				resource.Attribute{
					Name:        "datastore",
					Description: `(Optional) (string)`,
				},
				resource.Attribute{
					Name:        "insecure_flag",
					Description: `(Optional) (bool)`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional/Sensitive) (string)`,
				},
				resource.Attribute{
					Name:        "user",
					Description: `(Optional/Sensitive) (string)`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional) (string)`,
				},
				resource.Attribute{
					Name:        "soap_roundtrip_count",
					Description: `(Optional) (int)`,
				},
				resource.Attribute{
					Name:        "working_dir",
					Description: `(Optional) (string)`,
				},
				resource.Attribute{
					Name:        "vm_name",
					Description: `(Optional) (string)`,
				},
				resource.Attribute{
					Name:        "vm_uuid",
					Description: `(Optional) (string) ##### ` + "`" + `network` + "`" + ` ###### Arguments`,
				},
				resource.Attribute{
					Name:        "public_network",
					Description: `(Optional) (string) ### ` + "`" + `dns` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "nodelocal",
					Description: `(Optional) Nodelocal dns config (list Maxitem: 1)`,
				},
				resource.Attribute{
					Name:        "node_selector",
					Description: `(Optional) Node selector key pair (map)`,
				},
				resource.Attribute{
					Name:        "provider",
					Description: `(Optional) DNS provider. ` + "`" + `kube-dns` + "`" + `, ` + "`" + `coredns` + "`" + ` (default), and ` + "`" + `none` + "`" + ` are supported (string)`,
				},
				resource.Attribute{
					Name:        "reverse_cidrs",
					Description: `(Optional) Reverse CIDRs (list)`,
				},
				resource.Attribute{
					Name:        "upstream_nameservers",
					Description: `(Optional) Upstream nameservers (list) #### ` + "`" + `nodelocal` + "`" + ` ##### Arguments`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `(required) Nodelocal dns ip address (string)`,
				},
				resource.Attribute{
					Name:        "node_selector",
					Description: `(Optional) Node selector key pair (map) ### ` + "`" + `ingress` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "dns_policy",
					Description: `(Optional) Ingress controller DNS policy. ` + "`" + `ClusterFirstWithHostNet` + "`" + `, ` + "`" + `ClusterFirst` + "`" + `, ` + "`" + `Default` + "`" + `, and ` + "`" + `None` + "`" + ` are supported. [K8S dns Policy](https://kubernetes.io/docs/concepts/services-networking/dns-pod-service/#pod-s-dns-policy) (string)`,
				},
				resource.Attribute{
					Name:        "extra_args",
					Description: `(Optional) Extra arguments for the ingress controller (map)`,
				},
				resource.Attribute{
					Name:        "http_port",
					Description: `(Optional) Ingress controller http port (int)`,
				},
				resource.Attribute{
					Name:        "https_port",
					Description: `(Optional) Ingress controller https port (int)`,
				},
				resource.Attribute{
					Name:        "network_mode",
					Description: `(Optional) Networt mode for the ingress controller. ` + "`" + `hostNetwork` + "`" + `, ` + "`" + `hostPort` + "`" + ` and ` + "`" + `none` + "`" + ` are supported (string)`,
				},
				resource.Attribute{
					Name:        "node_selector",
					Description: `(Optional) Node selector key pair (map)`,
				},
				resource.Attribute{
					Name:        "options",
					Description: `(Optional) Ingress controller options (map)`,
				},
				resource.Attribute{
					Name:        "provider",
					Description: `(Optional) Ingress controller provider. ` + "`" + `nginx` + "`" + ` (default), and ` + "`" + `none` + "`" + ` are supported (string) ### ` + "`" + `monitoring` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "node_selector",
					Description: `(Optional) Node selector key pair (map)`,
				},
				resource.Attribute{
					Name:        "options",
					Description: `(Optional) Monitoring options (map)`,
				},
				resource.Attribute{
					Name:        "provider",
					Description: `(Optional/Computed) Monitoring provider (string) ### ` + "`" + `network` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "calico_network_provider",
					Description: `(Optional) Calico network provider config (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "canal_network_provider",
					Description: `(Optional) Canal network provider config (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "flannel_network_provider",
					Description: `(Optional) Flannel network provider config (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "weave_network_provider",
					Description: `(Optional) Weave network provider config (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "aci_network_provider",
					Description: `(Optional) Aci network provider config (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "mtu",
					Description: `(Optional) Network provider MTU. Default ` + "`" + `0` + "`" + ` (int)`,
				},
				resource.Attribute{
					Name:        "options",
					Description: `(Optional/Computed) Network provider options (map)`,
				},
				resource.Attribute{
					Name:        "plugin",
					Description: `(Optional) Network provider plugin. ` + "`" + `calico` + "`" + `, ` + "`" + `canal` + "`" + ` (default), ` + "`" + `flannel` + "`" + `, ` + "`" + `none` + "`" + ` and ` + "`" + `weave` + "`" + ` are supported. (string) #### ` + "`" + `calico_network_provider` + "`" + ` ##### Arguments`,
				},
				resource.Attribute{
					Name:        "cloud_provider",
					Description: `(Optional/Computed) Calico cloud provider (string) #### ` + "`" + `canal_network_provider` + "`" + ` ##### Arguments`,
				},
				resource.Attribute{
					Name:        "iface",
					Description: `(Optional/Computed) Canal network interface (string) #### ` + "`" + `flannel_network_provider` + "`" + ` ##### Arguments`,
				},
				resource.Attribute{
					Name:        "iface",
					Description: `(Optional/Computed) Flannel network interface (string) #### ` + "`" + `weave_network_provider` + "`" + ` ##### Arguments`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional/Computed) Weave password (string) #### ` + "`" + `aci_network_provider` + "`" + ` ##### Arguments`,
				},
				resource.Attribute{
					Name:        "system_id",
					Description: `(Required) Unique suffix for all cluster related objects in aci (string)`,
				},
				resource.Attribute{
					Name:        "apic_hosts",
					Description: `(Required) Ip address for apic hosts (list)`,
				},
				resource.Attribute{
					Name:        "token",
					Description: `(Required/Sensitive) UUID for this version of the input configuration (string)`,
				},
				resource.Attribute{
					Name:        "apic_user_name",
					Description: `(Required) User name for aci apic (string)`,
				},
				resource.Attribute{
					Name:        "apic_user_key",
					Description: `(Required/Sensitive) Base64 encoded private key for aci apic user (string)`,
				},
				resource.Attribute{
					Name:        "apic_user_crt",
					Description: `(Required/Sensitive) Base64 encoded certificate for aci apic user (string)`,
				},
				resource.Attribute{
					Name:        "encap_type",
					Description: `(Required) One of the supported encap types for aci(vlan/vxlan) (string)`,
				},
				resource.Attribute{
					Name:        "mcast_range_start",
					Description: `(Required) Mcast range start address for endpoint groups on aci (string)`,
				},
				resource.Attribute{
					Name:        "mcast_range_end",
					Description: `(Required) Mcast range end address for endpoint groups on aci (string)`,
				},
				resource.Attribute{
					Name:        "aep",
					Description: `(Required) Attachment entity profile name on aci (string)`,
				},
				resource.Attribute{
					Name:        "vrf_name",
					Description: `(Required) VRF Name on aci (string)`,
				},
				resource.Attribute{
					Name:        "vrf_tenant",
					Description: `(Required) Tenant for VRF on aci (string)`,
				},
				resource.Attribute{
					Name:        "l3out",
					Description: `(Required) L3Out on aci (string)`,
				},
				resource.Attribute{
					Name:        "node_subnet",
					Description: `(Required) Kubernetes node address subnet (string)`,
				},
				resource.Attribute{
					Name:        "l3out_external_networks",
					Description: `(Required) L3out external networks on aci (list)`,
				},
				resource.Attribute{
					Name:        "extern_dynamic",
					Description: `(Required) Subnet to use for dynamic external IPs on aci (string)`,
				},
				resource.Attribute{
					Name:        "extern_static",
					Description: `(Required) Subnet to use for static external IPs on aci (string)`,
				},
				resource.Attribute{
					Name:        "node_svc_subnet",
					Description: `(Required) Subnet to use for service graph endpoints on aci (string)`,
				},
				resource.Attribute{
					Name:        "kube_api_vlan",
					Description: `(Required) Vlan for node network on aci (string)`,
				},
				resource.Attribute{
					Name:        "service_vlan",
					Description: `(Required) Vlan for service graph nodes on aci (string)`,
				},
				resource.Attribute{
					Name:        "infra_vlan",
					Description: `(Required) Vlan for infra network on aci (string)`,
				},
				resource.Attribute{
					Name:        "snat_port_range_start",
					Description: `(Optional) Port start range for Source Network Address Translation on aci (string)`,
				},
				resource.Attribute{
					Name:        "snat_port_range_end",
					Description: `(Optional) Port end range for Source Network Address Translation on aci (string)`,
				},
				resource.Attribute{
					Name:        "snat_ports_per_node",
					Description: `(Optional) Ports per node for Source Network Address Translation on aci (string) ### ` + "`" + `nodes` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `(Required) Address ip for node (string)`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `(Required) Node roles in k8s cluster. ` + "`" + `controlplane` + "`" + `, ` + "`" + `etcd` + "`" + ` and ` + "`" + `worker` + "`" + ` are supported. (list)`,
				},
				resource.Attribute{
					Name:        "user",
					Description: `(Required/Sensitive) SSH user that will be used by RKE (string)`,
				},
				resource.Attribute{
					Name:        "docker_socket",
					Description: `(Optional) Docker socket on the node that will be used in tunneling (string)`,
				},
				resource.Attribute{
					Name:        "hostname_override",
					Description: `(Optional) Hostname override for node (string)`,
				},
				resource.Attribute{
					Name:        "internal_address",
					Description: `(Optional) Internal address that will be used for components communication (string)`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) Node labels (map)`,
				},
				resource.Attribute{
					Name:        "node_name",
					Description: `(Optional) Name of the host provisioned via docker machine (string)`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional) Port used for SSH communication (string)`,
				},
				resource.Attribute{
					Name:        "ssh_agent_auth",
					Description: `(Optional/Computed) SSH Agent Auth enable (bool)`,
				},
				resource.Attribute{
					Name:        "ssh_cert",
					Description: `(Optional/Sensitive) SSH Certificate (string)`,
				},
				resource.Attribute{
					Name:        "ssh_cert_path",
					Description: `(Optional) SSH Certificate path (string)`,
				},
				resource.Attribute{
					Name:        "ssh_key",
					Description: `(Optional/Sensitive) SSH Private Key (string)`,
				},
				resource.Attribute{
					Name:        "ssh_key_path",
					Description: `(Optional) SSH Private Key path (string)`,
				},
				resource.Attribute{
					Name:        "taints",
					Description: `(Optional) Node taints (list) #### ` + "`" + `taint` + "`" + ` ##### Arguments`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) Taint key (string)`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) Taint value (string)`,
				},
				resource.Attribute{
					Name:        "effect",
					Description: `(Optional) Taint effect. ` + "`" + `NoExecute` + "`" + `, ` + "`" + `NoSchedule` + "`" + ` (default) and ` + "`" + `PreferNoSchedule` + "`" + ` are supported (string) ### ` + "`" + `private_registries` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Required) Registry URL (string)`,
				},
				resource.Attribute{
					Name:        "is_default",
					Description: `(Optional) Set as default registry. Default ` + "`" + `false` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional/Sensitive) Registry password (string)`,
				},
				resource.Attribute{
					Name:        "user",
					Description: `(Optional/Sensitive) Registry user (string) ### ` + "`" + `restore` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "restore",
					Description: `(Optional) Restore cluster. Default ` + "`" + `false` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "snapshot_name",
					Description: `(Optional) Snapshot name (string) ### ` + "`" + `rotate_certificates` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "ca_certificates",
					Description: `(Optional) Rotate CA Certificates. Default ` + "`" + `false` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "services",
					Description: `(Optional) Services to rotate their certs. ` + "`" + `etcd` + "`" + `, ` + "`" + `kubelet` + "`" + `, ` + "`" + `kube-apiserver` + "`" + `, ` + "`" + `kube-proxy` + "`" + `, ` + "`" + `kube-scheduler` + "`" + ` and ` + "`" + `kube-controller-manager` + "`" + ` are supported (list) ### ` + "`" + `services` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "etcd",
					Description: `(Optional/Computed) Etcd options for RKE services (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "kube_api",
					Description: `(Optional/Computed) Kube API options for RKE services (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "kube_controller",
					Description: `(Optional/Computed) Kube Controller options for RKE services (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "kubelet",
					Description: `(Optional/Computed) Kubelet options for RKE services (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "kubeproxy",
					Description: `(Optional/Computed) Kubeproxy options for RKE services (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "scheduler",
					Description: `(Optional/Computed) Scheduler options for RKE services (list maxitems:1) #### ` + "`" + `etcd` + "`" + ` ##### Arguments`,
				},
				resource.Attribute{
					Name:        "backup_config",
					Description: `(Optional/Computed) Backup options for etcd service. Just for Rancher v2.2.x (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "ca_cert",
					Description: `(Optional/Computed/Sensitive) TLS CA certificate for etcd service (string)`,
				},
				resource.Attribute{
					Name:        "cert",
					Description: `(Optional/Computed/Sensitive) TLS certificate for etcd service (string)`,
				},
				resource.Attribute{
					Name:        "creation",
					Description: `(Optional/Computed) Creation option for etcd service (string)`,
				},
				resource.Attribute{
					Name:        "external_urls",
					Description: `(Optional/Computed) External urls for etcd service (list)`,
				},
				resource.Attribute{
					Name:        "extra_args",
					Description: `(Optional/Computed) Extra arguments for etcd service (map)`,
				},
				resource.Attribute{
					Name:        "extra_binds",
					Description: `(Optional/Computed) Extra binds for etcd service (list)`,
				},
				resource.Attribute{
					Name:        "extra_env",
					Description: `(Optional/Computed) Extra environment for etcd service (list)`,
				},
				resource.Attribute{
					Name:        "gid",
					Description: `(Optional) Etcd service GID. Default: ` + "`" + `0` + "`" + `. For Rancher v2.3.x or above (int)`,
				},
				resource.Attribute{
					Name:        "image",
					Description: `(Optional/Computed) Docker image for etcd service (string)`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Optional/Computed/Sensitive) TLS key for etcd service (string)`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Optional/Computed) Path for etcd service (string)`,
				},
				resource.Attribute{
					Name:        "retention",
					Description: `(Optional/Computed) Retention option for etcd service (string)`,
				},
				resource.Attribute{
					Name:        "snapshot",
					Description: `(Optional) Snapshot option for etcd service. Default ` + "`" + `true` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "uid",
					Description: `(Optional) Etcd service UID. Default: ` + "`" + `0` + "`" + `. For Rancher v2.3.x or above (int) ##### ` + "`" + `backup_config` + "`" + ` ###### Arguments`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) Enable etcd backup. Default ` + "`" + `true` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "interval_hours",
					Description: `(Optional) Interval hours for etcd backup. Default ` + "`" + `12` + "`" + ` (int)`,
				},
				resource.Attribute{
					Name:        "retention",
					Description: `(Optional) Retention for etcd backup. Default ` + "`" + `6` + "`" + ` (int)`,
				},
				resource.Attribute{
					Name:        "s3_backup_config",
					Description: `(Optional) S3 config options for etcd backup (list maxitems:1)`,
				},
				resource.Attribute{
					Name:        "safe_timestamp",
					Description: `(Optional) Safe timestamp for etcd backup. Default: ` + "`" + `false` + "`" + ` (bool)`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `(Optional/Computed) Timeout in seconds for etcd backup. Default: ` + "`" + `300` + "`" + `. Just for RKE v1.2.6 and above (int) ###### ` + "`" + `s3_backup_config` + "`" + ` ###### Arguments`,
				},
				resource.Attribute{
					Name:        "access_key",
					Description: `(Optional/Sensitive) Access key for S3 service (string)`,
				},
				resource.Attribute{
					Name:        "bucket_name",
					Description: `(Optional) Bucket name for S3 service (string)`,
				},
				resource.Attribute{
					Name:        "custom_ca",
					Description: `(Optional/Sensitive) Base64 encoded custom CA for S3 service. Use filebase64(<FILE>) for encoding file. Available from Rancher v2.2.5 (string)`,
				},
				resource.Attribute{
					Name:        "endpoint",
					Description: `(Optional) Endpoint for S3 service (string)`,
				},
				resource.Attribute{
					Name:        "folder",
					Description: `(Optional) Folder for S3 service. Available from Rancher v2.2.7 (string)`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) Region for S3 service (string)`,
				},
				resource.Attribute{
					Name:        "secret_key",
					Description: `(Optional/Sensitive) Secret key for S3 service (string) #### ` + "`" + `kube_api` + "`" + ` ##### Arguments`,
				},
				resource.Attribute{
					Name:        "always_pull_images",
					Description: `(Optional/Computed) Enable [AlwaysPullImages](https://kubernetes.io/docs/reference/access-authn-authz/admission-controllers/#alwayspullimages) Admission controller plugin. [Rancher docs](https://rancher.com/docs/rke/latest/en/config-options/services/#kubernetes-api-server-options) (bool)`,
				},
				resource.Attribute{
					Name:        "audit_log",
					Description: `(Optional/Computed) K8s audit log configuration. (list maxitem: 1)`,
				},
				resource.Attribute{
					Name:        "event_rate_limit",
					Description: `(Optional) K8s event rate limit configuration. (list maxitem: 1)`,
				},
				resource.Attribute{
					Name:        "extra_args",
					Description: `(Optional/Computed) Extra arguments for kube API service (map)`,
				},
				resource.Attribute{
					Name:        "extra_binds",
					Description: `(Optional/Computed) Extra binds for kube API service (list)`,
				},
				resource.Attribute{
					Name:        "extra_env",
					Description: `(Optional/Computed) Extra environment for kube API service (list)`,
				},
				resource.Attribute{
					Name:        "image",
					Description: `(Optional/Computed) Docker image for kube API service (string)`,
				},
				resource.Attribute{
					Name:        "pod_security_policy",
					Description: `(Optional/Computed) Pod Security Policy option for kube API service (bool)`,
				},
				resource.Attribute{
					Name:        "secrets_encryption_config",
					Description: `(Optional) [Encrypt k8s secret data configration](https://rancher.com/docs/rke/latest/en/config-options/secrets-encryption/). (list maxitem: 1)`,
				},
				resource.Attribute{
					Name:        "service_cluster_ip_range",
					Description: `(Optional/Computed) Service Cluster IP Range option for kube API service (string)`,
				},
				resource.Attribute{
					Name:        "service_node_port_range",
					Description: `(Optional/Computed) Service Node Port Range option for kube API service (string) ##### ` + "`" + `audit_log` + "`" + ` ###### Arguments`,
				},
				resource.Attribute{
					Name:        "configuration",
					Description: `(Optional/Computed) Audit log configuration. (list maxtiem: 1)`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional/Computed) Enable audit log (bool) ###### ` + "`" + `configuration` + "`" + ` ###### Arguments`,
				},
				resource.Attribute{
					Name:        "format",
					Description: `(Optional/Computed) Audit log format (string)`,
				},
				resource.Attribute{
					Name:        "max_age",
					Description: `(Optional/Computed) Audit log max age (int)`,
				},
				resource.Attribute{
					Name:        "max_backup",
					Description: `(Optional/Computed) Audit log max backup. Default: ` + "`" + `10` + "`" + ` (int)`,
				},
				resource.Attribute{
					Name:        "max_size",
					Description: `(Optional/Computed) Audit log max size. Default: ` + "`" + `100` + "`" + ` (int)`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Optional/Computed) Audit log path. Default: ` + "`" + `/var/log/kube-audit/audit-log.json` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "policy",
					Description: `(Optional/Computed) Audit policy json encoded definition. ` + "`" + `"apiVersion"` + "`" + ` and ` + "`" + `"kind":"Policy","rules"` + "`" + ` fields are required in the json. Ex. ` + "`" + `jsonencode({"apiVersion":"audit.k8s.io/v1","kind":"Policy","rules":[{"level":"RequestResponse","resources":[{"group":"","resources":["pods"]}]}]})` + "`" + ` [More info](https://rancher.com/docs/rke/latest/en/config-options/audit-log/) (string) ##### ` + "`" + `event_rate_limit` + "`" + ` ###### Arguments`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional/Computed) Enable event rate limit (bool)`,
				},
				resource.Attribute{
					Name:        "configuration",
					Description: `(Optional) Event rate limit yaml encoded configuration. ` + "`" + `"apiVersion"` + "`" + ` and ` + "`" + `"kind":"Configuration"` + "`" + ` fields are required in the yaml. Ex. ` + "`" + `apiVersion: eventratelimit.admission.k8s.io/v1alpha1\nkind: Configuration\nlimits:\n- type: Server\n burst: 30000\n qps: 6000\n` + "`" + ` [More info](https://rancher.com/docs/rke/latest/en/config-options/rate-limiting/) (string) ##### ` + "`" + `secrets_encryption_config` + "`" + ` ###### Arguments`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional/Computed) Enable secrets encryption (bool)`,
				},
				resource.Attribute{
					Name:        "custom_config",
					Description: `(Optional) Secrets encryption yaml encoded custom configuration. ` + "`" + `"apiVersion"` + "`" + ` and ` + "`" + `"kind":"EncryptionConfiguration"` + "`" + ` fields are required in the yaml. Ex. ` + "`" + `apiVersion: apiserver.config.k8s.io/v1\nkind: EncryptionConfiguration\nresources:\n- resources:\n - secrets\n providers:\n - aescbc:\n keys:\n - name: k-fw5hn\n secret: RTczRjFDODMwQzAyMDVBREU4NDJBMUZFNDhCNzM5N0I=\n identity: {}\n` + "`" + ` [More info](https://rancher.com/docs/rke/latest/en/config-options/secrets-encryption/) (string) #### ` + "`" + `kube_controller` + "`" + ` ##### Arguments`,
				},
				resource.Attribute{
					Name:        "cluster_cidr",
					Description: `(Optional/Computed) Cluster CIDR option for kube controller service (string)`,
				},
				resource.Attribute{
					Name:        "extra_args",
					Description: `(Optional/Computed) Extra arguments for kube controller service (map)`,
				},
				resource.Attribute{
					Name:        "extra_binds",
					Description: `(Optional/Computed) Extra binds for kube controller service (list)`,
				},
				resource.Attribute{
					Name:        "extra_env",
					Description: `(Optional/Computed) Extra environment for kube controller service (list)`,
				},
				resource.Attribute{
					Name:        "image",
					Description: `(Optional/Computed) Docker image for kube controller service (string)`,
				},
				resource.Attribute{
					Name:        "service_cluster_ip_range",
					Description: `(Optional/Computed) Service Cluster ip Range option for kube controller service (string) #### ` + "`" + `kubelet` + "`" + ` ##### Arguments`,
				},
				resource.Attribute{
					Name:        "cluster_dns_server",
					Description: `(Optional/Computed) Cluster DNS Server option for kubelet service (string)`,
				},
				resource.Attribute{
					Name:        "cluster_domain",
					Description: `(Optional) Cluster Domain option for kubelet service. Default ` + "`" + `cluster.local` + "`" + ` (string)`,
				},
				resource.Attribute{
					Name:        "extra_args",
					Description: `(Optional/Computed) Extra arguments for kubelet service (map)`,
				},
				resource.Attribute{
					Name:        "extra_binds",
					Description: `(Optional/Computed) Extra binds for kubelet service (list)`,
				},
				resource.Attribute{
					Name:        "extra_env",
					Description: `(Optional/Computed) Extra environment for kubelet service (list)`,
				},
				resource.Attribute{
					Name:        "fail_swap_on",
					Description: `(Optional/Computed) Enable or disable failing when swap on is not supported (bool)`,
				},
				resource.Attribute{
					Name:        "image",
					Description: `(Optional/Computed) Docker image for kubelet service (string)`,
				},
				resource.Attribute{
					Name:        "infra_container_image",
					Description: `(Optional/Computed) Infra container image for kubelet service (string) #### ` + "`" + `kubeproxy` + "`" + ` ##### Arguments`,
				},
				resource.Attribute{
					Name:        "extra_args",
					Description: `(Optional/Computed) Extra arguments for kubeproxy service (map)`,
				},
				resource.Attribute{
					Name:        "extra_binds",
					Description: `(Optional/Computed) Extra binds for kubeproxy service (list)`,
				},
				resource.Attribute{
					Name:        "extra_env",
					Description: `(Optional/Computed) Extra environment for kubeproxy service (list)`,
				},
				resource.Attribute{
					Name:        "image",
					Description: `(Optional/Computed) Docker image for kubeproxy service (string) #### ` + "`" + `scheduler` + "`" + ` ##### Arguments`,
				},
				resource.Attribute{
					Name:        "extra_args",
					Description: `(Optional/Computed) Extra arguments for scheduler service (map)`,
				},
				resource.Attribute{
					Name:        "extra_binds",
					Description: `(Optional/Computed) Extra binds for scheduler service (list)`,
				},
				resource.Attribute{
					Name:        "extra_env",
					Description: `(Optional/Computed) Extra environment for scheduler service (list)`,
				},
				resource.Attribute{
					Name:        "image",
					Description: `(Optional/Computed) Docker image for scheduler service (string) ### ` + "`" + `system_images` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "etcd",
					Description: `(Optional) Docker image for etcd (string)`,
				},
				resource.Attribute{
					Name:        "alpine",
					Description: `(Optional) Docker image for alpine (string)`,
				},
				resource.Attribute{
					Name:        "nginx_proxy",
					Description: `(Optional) Docker image for nginx_proxy (string)`,
				},
				resource.Attribute{
					Name:        "cert_downloader",
					Description: `(Optional) Docker image for cert_downloader (string)`,
				},
				resource.Attribute{
					Name:        "kubernetes_services_sidecar",
					Description: `(Optional) Docker image for kubernetes_services_sidecar (string)`,
				},
				resource.Attribute{
					Name:        "kube_dns",
					Description: `(Optional) Docker image for kube_dns (string)`,
				},
				resource.Attribute{
					Name:        "dnsmasq",
					Description: `(Optional) Docker image for dnsmasq (string)`,
				},
				resource.Attribute{
					Name:        "kube_dns_sidecar",
					Description: `(Optional) Docker image for kube_dns_sidecar (string)`,
				},
				resource.Attribute{
					Name:        "kube_dns_autoscaler",
					Description: `(Optional) Docker image for kube_dns_autoscaler (string)`,
				},
				resource.Attribute{
					Name:        "coredns",
					Description: `(Optional) Docker image for coredns (string)`,
				},
				resource.Attribute{
					Name:        "coredns_autoscaler",
					Description: `(Optional) Docker image for coredns_autoscaler (string)`,
				},
				resource.Attribute{
					Name:        "nodelocal",
					Description: `(Optional) Docker image for nodelocal (string)`,
				},
				resource.Attribute{
					Name:        "kubernetes",
					Description: `(Optional) Docker image for kubernetes (string)`,
				},
				resource.Attribute{
					Name:        "flannel",
					Description: `(Optional) Docker image for flannel (string)`,
				},
				resource.Attribute{
					Name:        "flannel_cni",
					Description: `(Optional) Docker image for flannel_cni (string)`,
				},
				resource.Attribute{
					Name:        "calico_node",
					Description: `(Optional) Docker image for calico_node (string)`,
				},
				resource.Attribute{
					Name:        "calico_cni",
					Description: `(Optional) Docker image for calico_cni (string)`,
				},
				resource.Attribute{
					Name:        "calico_controllers",
					Description: `(Optional) Docker image for calico_controllers (string)`,
				},
				resource.Attribute{
					Name:        "calico_ctl",
					Description: `(Optional) Docker image for calico_ctl (string)`,
				},
				resource.Attribute{
					Name:        "calico_flex_vol",
					Description: `(Optional) Docker image for calico_flex_vol (string)`,
				},
				resource.Attribute{
					Name:        "canal_node",
					Description: `(Optional) Docker image for canal_node (string)`,
				},
				resource.Attribute{
					Name:        "canal_cni",
					Description: `(Optional) Docker image for canal_cni (string)`,
				},
				resource.Attribute{
					Name:        "canal_flannel",
					Description: `(Optional) Docker image for canal_flannel (string)`,
				},
				resource.Attribute{
					Name:        "canal_flex_vol",
					Description: `(Optional) Docker image for canal_flex_vol (string)`,
				},
				resource.Attribute{
					Name:        "weave_node",
					Description: `(Optional) Docker image for weave_node (string)`,
				},
				resource.Attribute{
					Name:        "weave_cni",
					Description: `(Optional) Docker image for weave_cni (string)`,
				},
				resource.Attribute{
					Name:        "pod_infra_container",
					Description: `(Optional) Docker image for pod_infra_container (string)`,
				},
				resource.Attribute{
					Name:        "ingress",
					Description: `(Optional) Docker image for ingress (string)`,
				},
				resource.Attribute{
					Name:        "ingress_backend",
					Description: `(Optional) Docker image for ingress_backend (string)`,
				},
				resource.Attribute{
					Name:        "metrics_server",
					Description: `(Optional) Docker image for metrics_server (string)`,
				},
				resource.Attribute{
					Name:        "windows_pod_infra_container",
					Description: `(Optional) Docker image for windows_pod_infra_container (string)`,
				},
				resource.Attribute{
					Name:        "aci_cni_deploy_container",
					Description: `(Optional) Docker image for aci_cni_deploy_container (string)`,
				},
				resource.Attribute{
					Name:        "aci_host_container",
					Description: `(Optional) Docker image for aci_host_container (string)`,
				},
				resource.Attribute{
					Name:        "aci_opflex_container",
					Description: `(Optional) Docker image for aci_opflex_container (string)`,
				},
				resource.Attribute{
					Name:        "aci_mcast_container",
					Description: `(Optional) Docker image for aci_mcast_container (string)`,
				},
				resource.Attribute{
					Name:        "aci_ovs_container",
					Description: `(Optional) Docker image for aci_ovs_container (string)`,
				},
				resource.Attribute{
					Name:        "aci_controller_container",
					Description: `(Optional) Docker image for aci_controller_container (string) ### ` + "`" + `upgrade_strategy` + "`" + ` #### Arguments`,
				},
				resource.Attribute{
					Name:        "drain",
					Description: `(Optional/Computed) RKE drain nodes (bool)`,
				},
				resource.Attribute{
					Name:        "drain_input",
					Description: `(Optional/Computed) RKE drain node input (list Maxitems: 1)`,
				},
				resource.Attribute{
					Name:        "max_unavailable_controlplane",
					Description: `(Optional/Computed) RKE max unavailable controlplane nodes (string)`,
				},
				resource.Attribute{
					Name:        "max_unavailable_worker",
					Description: `(Optional/Computed) RKE max unavailable worker nodes (string) #### ` + "`" + `drain_input` + "`" + ` ##### Arguments`,
				},
				resource.Attribute{
					Name:        "delete_local_data",
					Description: `(Optional/Computed) Delete RKE node local data (bool)`,
				},
				resource.Attribute{
					Name:        "force",
					Description: `(Optional/Computed) Force RKE node drain (bool)`,
				},
				resource.Attribute{
					Name:        "grace_period",
					Description: `(Optional/Computed) RKE node drain grace period (int)`,
				},
				resource.Attribute{
					Name:        "ignore_daemon_sets",
					Description: `(Optional/Computed) Ignore RKE daemon sets (bool)`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `(Optional/Computed) RKE node drain timeout (int) ## Timeouts ` + "`" + `rke_cluster` + "`" + ` provides the following [Timeouts](https://www.terraform.io/docs/configuration/resources.html#operation-timeouts) configuration options: - ` + "`" + `create` + "`" + ` - (Default ` + "`" + `30 minutes` + "`" + `) Used for creating clusters. - ` + "`" + `update` + "`" + ` - (Default ` + "`" + `30 minutes` + "`" + `) Used for cluster modifications. - ` + "`" + `delete` + "`" + ` - (Default ` + "`" + `30 minutes` + "`" + `) Used for deleting clusters. ## Import rke_cluster can be imported using the RKE cluster config and state files as ID in the format ` + "`" + `<cluster_config_file>:<rke_state_file>` + "`" + ` ` + "`" + `` + "`" + `` + "`" + ` $ terraform import rke_cluster.foo &lt;cluster_config_file&gt;:&lt;rke_state_file&gt; ` + "`" + `` + "`" + `` + "`" + ` As experimental feature, dind rke_cluster can be also imported adding ` + "`" + `dind` + "`" + ` as 3rd import parameter ` + "`" + `<cluster_config_file>:<rke_state_file>:dind` + "`" + ` ` + "`" + `` + "`" + `` + "`" + ` $ terraform import rke_cluster.foo &lt;cluster_config_file&gt;:&lt;rke_state_file&gt;:dind ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
	}

	resourcesMap = map[string]int{

		"rke_cluster": 0,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
