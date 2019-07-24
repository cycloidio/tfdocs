package aws

import (
	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*Resource{

		&resource.Resource{
			Name:             "",
			Type:             "vault_approle_auth_backend_role_id",
			Category:         "Data Sources",
			ShortDescription: `Manages AppRole auth backend roles in Vault.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "role_name",
					Description: `(Required) The name of the role to retrieve the Role ID for.`,
				},
				resource.Attribute{
					Name:        "backend",
					Description: `(Optional) The unique name for the AppRole backend the role to retrieve a RoleID for resides in. Defaults to "approle". ## Attributes Reference In addition to the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "role_id",
					Description: `The RoleID of the role.`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "role_id",
					Description: `The RoleID of the role.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_aws_access_credentials",
			Category:         "Data Sources",
			ShortDescription: `Reads AWS credentials from an AWS secret backend in Vault`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "backend",
					Description: `(Required) The path to the AWS secret backend to read credentials from, with no leading or trailing ` + "`" + `/` + "`" + `s.`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `(Required) The name of the AWS secret backend role to read credentials from, with no leading or trailing ` + "`" + `/` + "`" + `s.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) The type of credentials to read. Defaults to ` + "`" + `"creds"` + "`" + `, which just returns an AWS Access Key ID and Secret Key. Can also be set to ` + "`" + `"sts"` + "`" + `, which will return a security token in addition to the keys. ## Attributes Reference In addition to the arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "access_key",
					Description: `The AWS Access Key ID returned by Vault.`,
				},
				resource.Attribute{
					Name:        "secret_key",
					Description: `The AWS Secret Key returned by Vault.`,
				},
				resource.Attribute{
					Name:        "security_token",
					Description: `The STS token returned by Vault, if any.`,
				},
				resource.Attribute{
					Name:        "lease_id",
					Description: `The lease identifier assigned by Vault.`,
				},
				resource.Attribute{
					Name:        "lease_duration",
					Description: `The duration of the secret lease, in seconds relative to the time the data was requested. Once this time has passed any plan generated with this data may fail to apply.`,
				},
				resource.Attribute{
					Name:        "lease_start_time",
					Description: `As a convenience, this records the current time on the computer where Terraform is running when the data is requested. This can be used to approximate the absolute time represented by ` + "`" + `lease_duration` + "`" + `, though users must allow for any clock drift and response latency relative to the Vault server.`,
				},
				resource.Attribute{
					Name:        "lease_renewable",
					Description: `` + "`" + `true` + "`" + ` if the lease can be renewed using Vault's ` + "`" + `sys/renew/{lease-id}` + "`" + ` endpoint. Terraform does not currently support lease renewal, and so it will request a new lease each time this data source is refreshed.`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "access_key",
					Description: `The AWS Access Key ID returned by Vault.`,
				},
				resource.Attribute{
					Name:        "secret_key",
					Description: `The AWS Secret Key returned by Vault.`,
				},
				resource.Attribute{
					Name:        "security_token",
					Description: `The STS token returned by Vault, if any.`,
				},
				resource.Attribute{
					Name:        "lease_id",
					Description: `The lease identifier assigned by Vault.`,
				},
				resource.Attribute{
					Name:        "lease_duration",
					Description: `The duration of the secret lease, in seconds relative to the time the data was requested. Once this time has passed any plan generated with this data may fail to apply.`,
				},
				resource.Attribute{
					Name:        "lease_start_time",
					Description: `As a convenience, this records the current time on the computer where Terraform is running when the data is requested. This can be used to approximate the absolute time represented by ` + "`" + `lease_duration` + "`" + `, though users must allow for any clock drift and response latency relative to the Vault server.`,
				},
				resource.Attribute{
					Name:        "lease_renewable",
					Description: `` + "`" + `true` + "`" + ` if the lease can be renewed using Vault's ` + "`" + `sys/renew/{lease-id}` + "`" + ` endpoint. Terraform does not currently support lease renewal, and so it will request a new lease each time this data source is refreshed.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_generic_secret",
			Category:         "Data Sources",
			ShortDescription: `Reads arbitrary data from a given path in Vault`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "path",
					Description: `(Required) The full logical path from which to request data. To read data from the "generic" secret backend mounted in Vault by default, this should be prefixed with ` + "`" + `secret/` + "`" + `. Reading from other backends with this data source is possible; consult each backend's documentation to see which endpoints support the ` + "`" + `GET` + "`" + ` method. ## Required Vault Capabilities Use of this resource requires the ` + "`" + `read` + "`" + ` capability on the given path. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "data_json",
					Description: `A string containing the full data payload retrieved from Vault, serialized in JSON format.`,
				},
				resource.Attribute{
					Name:        "data",
					Description: `A mapping whose keys are the top-level data keys returned from Vault and whose values are the corresponding values. This map can only represent string data, so any non-string values returned from Vault are serialized as JSON.`,
				},
				resource.Attribute{
					Name:        "lease_id",
					Description: `The lease identifier assigned by Vault, if any.`,
				},
				resource.Attribute{
					Name:        "lease_duration",
					Description: `The duration of the secret lease, in seconds relative to the time the data was requested. Once this time has passed any plan generated with this data may fail to apply.`,
				},
				resource.Attribute{
					Name:        "lease_start_time",
					Description: `As a convenience, this records the current time on the computer where Terraform is running when the data is requested. This can be used to approximate the absolute time represented by ` + "`" + `lease_duration` + "`" + `, though users must allow for any clock drift and response latency relative to to the Vault server.`,
				},
				resource.Attribute{
					Name:        "lease_renewable",
					Description: `` + "`" + `true` + "`" + ` if the lease can be renewed using Vault's ` + "`" + `sys/renew/{lease-id}` + "`" + ` endpoint. Terraform does not currently support lease renewal, and so it will request a new lease each time this data source is refreshed.`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "data_json",
					Description: `A string containing the full data payload retrieved from Vault, serialized in JSON format.`,
				},
				resource.Attribute{
					Name:        "data",
					Description: `A mapping whose keys are the top-level data keys returned from Vault and whose values are the corresponding values. This map can only represent string data, so any non-string values returned from Vault are serialized as JSON.`,
				},
				resource.Attribute{
					Name:        "lease_id",
					Description: `The lease identifier assigned by Vault, if any.`,
				},
				resource.Attribute{
					Name:        "lease_duration",
					Description: `The duration of the secret lease, in seconds relative to the time the data was requested. Once this time has passed any plan generated with this data may fail to apply.`,
				},
				resource.Attribute{
					Name:        "lease_start_time",
					Description: `As a convenience, this records the current time on the computer where Terraform is running when the data is requested. This can be used to approximate the absolute time represented by ` + "`" + `lease_duration` + "`" + `, though users must allow for any clock drift and response latency relative to to the Vault server.`,
				},
				resource.Attribute{
					Name:        "lease_renewable",
					Description: `` + "`" + `true` + "`" + ` if the lease can be renewed using Vault's ` + "`" + `sys/renew/{lease-id}` + "`" + ` endpoint. Terraform does not currently support lease renewal, and so it will request a new lease each time this data source is refreshed.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_kubernetes_auth_backend_config",
			Category:         "Data Sources",
			ShortDescription: `Manages Kubernetes auth backend configs in Vault.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "backend",
					Description: `(Optional) The unique name for the Kubernetes backend the config to retrieve Role attributes for resides in. Defaults to "kubernetes". ## Attributes Reference In addition to the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "kubernetes_host",
					Description: `Host must be a host string, a host:port pair, or a URL to the base of the Kubernetes API server.`,
				},
				resource.Attribute{
					Name:        "kubernetes_ca_cert",
					Description: `PEM encoded CA cert for use by the TLS client used to talk with the Kubernetes API.`,
				},
				resource.Attribute{
					Name:        "pem_keys",
					Description: `Optional list of PEM-formatted public keys or certificates used to verify the signatures of Kubernetes service account JWTs. If a certificate is given, its public key will be extracted. Not every installation of Kubernetes exposes these keys.`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "kubernetes_host",
					Description: `Host must be a host string, a host:port pair, or a URL to the base of the Kubernetes API server.`,
				},
				resource.Attribute{
					Name:        "kubernetes_ca_cert",
					Description: `PEM encoded CA cert for use by the TLS client used to talk with the Kubernetes API.`,
				},
				resource.Attribute{
					Name:        "pem_keys",
					Description: `Optional list of PEM-formatted public keys or certificates used to verify the signatures of Kubernetes service account JWTs. If a certificate is given, its public key will be extracted. Not every installation of Kubernetes exposes these keys.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_kubernetes_auth_backend_role",
			Category:         "Data Sources",
			ShortDescription: `Manages Kubernetes auth backend roles in Vault.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "role_name",
					Description: `(Required) The name of the role to retrieve the Role attributes for.`,
				},
				resource.Attribute{
					Name:        "backend",
					Description: `(Optional) The unique name for the Kubernetes backend the role to retrieve Role attributes for resides in. Defaults to "kubernetes". ## Attributes Reference In addition to the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "bound_cirs",
					Description: `List of CIDR blocks. If set, specifies the blocks of IP addresses which can perform the login operation.`,
				},
				resource.Attribute{
					Name:        "bound_service_account_names",
					Description: `List of service account names able to access this role. If set to "`,
				},
				resource.Attribute{
					Name:        "bound_service_account_namespaces",
					Description: `List of namespaces allowed to access this role. If set to "`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `The TTL period of tokens issued using this role in seconds.`,
				},
				resource.Attribute{
					Name:        "max_ttl",
					Description: `The maximum allowed lifetime of tokens issued in seconds using this role.`,
				},
				resource.Attribute{
					Name:        "num_uses",
					Description: `Number of times issued tokens can be used. Setting this to 0 or leaving it unset means unlimited uses.`,
				},
				resource.Attribute{
					Name:        "period",
					Description: `If set, indicates that the token generated using this role should never expire. The token should be renewed within the duration specified by this value. At each renewal, the token's TTL will be set to the value of this parameter.`,
				},
				resource.Attribute{
					Name:        "policies",
					Description: `Policies to be set on tokens issued using this role.`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "bound_cirs",
					Description: `List of CIDR blocks. If set, specifies the blocks of IP addresses which can perform the login operation.`,
				},
				resource.Attribute{
					Name:        "bound_service_account_names",
					Description: `List of service account names able to access this role. If set to "`,
				},
				resource.Attribute{
					Name:        "bound_service_account_namespaces",
					Description: `List of namespaces allowed to access this role. If set to "`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `The TTL period of tokens issued using this role in seconds.`,
				},
				resource.Attribute{
					Name:        "max_ttl",
					Description: `The maximum allowed lifetime of tokens issued in seconds using this role.`,
				},
				resource.Attribute{
					Name:        "num_uses",
					Description: `Number of times issued tokens can be used. Setting this to 0 or leaving it unset means unlimited uses.`,
				},
				resource.Attribute{
					Name:        "period",
					Description: `If set, indicates that the token generated using this role should never expire. The token should be renewed within the duration specified by this value. At each renewal, the token's TTL will be set to the value of this parameter.`,
				},
				resource.Attribute{
					Name:        "policies",
					Description: `Policies to be set on tokens issued using this role.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_policy_document",
			Category:         "Data Sources",
			ShortDescription: `Generates an Vault policy document in HCL format.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "path",
					Description: `(Required) A path in Vault that this rule applies to.`,
				},
				resource.Attribute{
					Name:        "capabilities",
					Description: `(Required) A list of capabilities that this rule apply to ` + "`" + `path` + "`" + `. For example, ["read", "write"].`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the rule. Will be added as a commend to rendered rule.`,
				},
				resource.Attribute{
					Name:        "required_parameters",
					Description: `(Optional) A list of parameters that must be specified.`,
				},
				resource.Attribute{
					Name:        "allowed_parameter",
					Description: `(Optional) Whitelists a list of keys and values that are permitted on the given path. See [Parameters](#Parameters) below.`,
				},
				resource.Attribute{
					Name:        "denied_parameter",
					Description: `(Optional) Blacklists a list of parameter and values. Any values specified here take precedence over ` + "`" + `allowed_parameter` + "`" + `. See [Parameters](#Parameters) below.`,
				},
				resource.Attribute{
					Name:        "min_wrapping_ttl",
					Description: `(Optional) The minimum allowed TTL that clients can specify for a wrapped response.`,
				},
				resource.Attribute{
					Name:        "max_wrapping_ttl",
					Description: `(Optional) The maximum allowed TTL that clients can specify for a wrapped response. ### Parameters Each of ` + "`" + ``,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) name of permitted or denied parameter.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) list of values what are permitted or denied by policy rule. ## Attributes Reference In addition to the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "hcl",
					Description: `The above arguments serialized as a standard Vault HCL policy document.`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "hcl",
					Description: `The above arguments serialized as a standard Vault HCL policy document.`,
				},
			},
		},
	}

	dataSourcesMap = map[string]Resource{

		"vault_approle_auth_backend_role_id":   0,
		"vault_aws_access_credentials":         1,
		"vault_generic_secret":                 2,
		"vault_kubernetes_auth_backend_config": 3,
		"vault_kubernetes_auth_backend_role":   4,
		"vault_policy_document":                5,
	}
)

func GetDataSource(r string) (*resouce.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs]
}
