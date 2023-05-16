package vault

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "vault_ad_secret_backend",
			Category:         "Resources",
			ShortDescription: `Creates an Active Directory secret backend for Vault.`,
			Description:      ``,
			Keywords: []string{
				"ad",
				"secret",
				"backend",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) The namespace to provision the resource in. The value should not contain leading or trailing forward slashes. The ` + "`" + `namespace` + "`" + ` is always relative to the provider's configured [namespace](/docs/providers/vault#namespace).`,
				},
				resource.Attribute{
					Name:        "backend",
					Description: `(Optional) The unique path this backend should be mounted at. Must not begin or end with a ` + "`" + `/` + "`" + `. Defaults to ` + "`" + `ad` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "disable_remount",
					Description: `(Optional) If set, opts out of mount migration on path updates. See here for more info on [Mount Migration](https://www.vaultproject.io/docs/concepts/mount-migration)`,
				},
				resource.Attribute{
					Name:        "anonymous_group_search",
					Description: `(Optional) Use anonymous binds when performing LDAP group searches (if true the initial credentials will still be used for the initial connection test).`,
				},
				resource.Attribute{
					Name:        "binddn",
					Description: `(Required) Distinguished name of object to bind when performing user and group search.`,
				},
				resource.Attribute{
					Name:        "bindpass",
					Description: `(Required) Password to use along with binddn when performing user search.`,
				},
				resource.Attribute{
					Name:        "case_sensitive_names",
					Description: `(Optional) If set, user and group names assigned to policies within the backend will be case sensitive. Otherwise, names will be normalized to lower case.`,
				},
				resource.Attribute{
					Name:        "certificate",
					Description: `(Optional) CA certificate to use when verifying LDAP server certificate, must be x509 PEM encoded.`,
				},
				resource.Attribute{
					Name:        "client_tls_cert",
					Description: `(Optional) Client certificate to provide to the LDAP server, must be x509 PEM encoded.`,
				},
				resource.Attribute{
					Name:        "client_tls_key",
					Description: `(Optional) Client certificate key to provide to the LDAP server, must be x509 PEM encoded.`,
				},
				resource.Attribute{
					Name:        "default_lease_ttl_seconds",
					Description: `(Optional) Default lease duration for secrets in seconds.`,
				},
				resource.Attribute{
					Name:        "deny_null_bind",
					Description: `(Optional) Denies an unauthenticated LDAP bind request if the user's password is empty; defaults to true.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Human-friendly description of the mount for the Active Directory backend.`,
				},
				resource.Attribute{
					Name:        "discoverdn",
					Description: `(Optional) Use anonymous bind to discover the bind Distinguished Name of a user.`,
				},
				resource.Attribute{
					Name:        "formatter",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "groupattr",
					Description: `(Optional) LDAP attribute to follow on objects returned by <groupfilter> in order to enumerate user group membership. Examples: ` + "`" + `cn` + "`" + ` or ` + "`" + `memberOf` + "`" + `, etc. Defaults to ` + "`" + `cn` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "groupdn",
					Description: `(Optional) LDAP search base to use for group membership search (eg: ou=Groups,dc=example,dc=org).`,
				},
				resource.Attribute{
					Name:        "groupfilter",
					Description: `(Optional) Go template for querying group membership of user (optional) The template can access the following context variables: UserDN, Username. Defaults to ` + "`" + `(|(memberUid={{.Username}})(member={{.UserDN}})(uniqueMember={{.UserDN}}))` + "`" + ``,
				},
				resource.Attribute{
					Name:        "insecure_tls",
					Description: `(Optional) Skip LDAP server SSL Certificate verification. This is not recommended for production. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "last_rotation_tolerance",
					Description: `(Optional) The number of seconds after a Vault rotation where, if Active Directory shows a later rotation, it should be considered out-of-band`,
				},
				resource.Attribute{
					Name:        "length",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "local",
					Description: `(Optional) Mark the secrets engine as local-only. Local engines are not replicated or removed by replication.Tolerance duration to use when checking the last rotation time.`,
				},
				resource.Attribute{
					Name:        "max_lease_ttl_seconds",
					Description: `(Optional) Maximum possible lease duration for secrets in seconds.`,
				},
				resource.Attribute{
					Name:        "max_ttl",
					Description: `(Optional) In seconds, the maximum password time-to-live.`,
				},
				resource.Attribute{
					Name:        "password_policy",
					Description: `(Optional) Name of the password policy to use to generate passwords.`,
				},
				resource.Attribute{
					Name:        "request_timeout",
					Description: `(Optional) Timeout, in seconds, for the connection when making requests against the server before returning back an error.`,
				},
				resource.Attribute{
					Name:        "starttls",
					Description: `(Optional) Issue a StartTLS command after establishing unencrypted connection.`,
				},
				resource.Attribute{
					Name:        "tls_max_version",
					Description: `(Optional) Maximum TLS version to use. Accepted values are ` + "`" + `tls10` + "`" + `, ` + "`" + `tls11` + "`" + `, ` + "`" + `tls12` + "`" + ` or ` + "`" + `tls13` + "`" + `. Defaults to ` + "`" + `tls12` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "tls_min_version",
					Description: `(Optional) Minimum TLS version to use. Accepted values are ` + "`" + `tls10` + "`" + `, ` + "`" + `tls11` + "`" + `, ` + "`" + `tls12` + "`" + ` or ` + "`" + `tls13` + "`" + `. Defaults to ` + "`" + `tls12` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `(Optional) In seconds, the default password time-to-live.`,
				},
				resource.Attribute{
					Name:        "upndomain",
					Description: `(Optional) Enables userPrincipalDomain login with [username]@UPNDomain.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Required) LDAP URL to connect to. Multiple URLs can be specified by concatenating them with commas; they will be tried in-order. Defaults to ` + "`" + `ldap://127.0.0.1` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "use_pre111_group_cn_behavior",
					Description: `(Optional) In Vault 1.1.1 a fix for handling group CN values of different cases unfortunately introduced a regression that could cause previously defined groups to not be found due to a change in the resulting name. If set true, the pre-1.1.1 behavior for matching group CNs will be used. This is only needed in some upgrade scenarios for backwards compatibility. It is enabled by default if the config is upgraded but disabled by default on new configurations.`,
				},
				resource.Attribute{
					Name:        "use_token_groups",
					Description: `(Optional) If true, use the Active Directory tokenGroups constructed attribute of the user to find the group memberships. This will find all security groups including nested ones.`,
				},
				resource.Attribute{
					Name:        "userattr",
					Description: `(Optional) Attribute used when searching users. Defaults to ` + "`" + `cn` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "userdn",
					Description: `(Optional) LDAP domain to use for users (eg: ou=People,dc=example,dc=org)` + "`" + `. ## Attributes Reference No additional attributes are exported by this resource. ## Import AD secret backend can be imported using the ` + "`" + `backend` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_ad_secret_backend.ad ad ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_ad_secret_backend_library",
			Category:         "Resources",
			ShortDescription: `Creates a library on the Active Directory Secret Backend for Vault.`,
			Description:      ``,
			Keywords: []string{
				"ad",
				"secret",
				"backend",
				"library",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) The namespace to provision the resource in. The value should not contain leading or trailing forward slashes. The ` + "`" + `namespace` + "`" + ` is always relative to the provider's configured [namespace](/docs/providers/vault#namespace).`,
				},
				resource.Attribute{
					Name:        "backend",
					Description: `(Required) The path the AD secret backend is mounted at, with no leading or trailing ` + "`" + `/` + "`" + `s.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name to identify this set of service accounts. Must be unique within the backend.`,
				},
				resource.Attribute{
					Name:        "service_account_names",
					Description: `(Required) Specifies the slice of service accounts mapped to this set.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `(Optional) The password time-to-live in seconds. Defaults to the configuration ttl if not provided.`,
				},
				resource.Attribute{
					Name:        "max_ttl",
					Description: `(Optional) The maximum password time-to-live in seconds. Defaults to the configuration max_ttl if not provided. ## Import AD secret backend libraries can be imported using the ` + "`" + `path` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_ad_secret_backend_library.role ad/library/bob ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_ad_secret_role",
			Category:         "Resources",
			ShortDescription: `Creates a role on the Active Directory Secret Backend for Vault.`,
			Description:      ``,
			Keywords: []string{
				"ad",
				"secret",
				"role",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) The namespace to provision the resource in. The value should not contain leading or trailing forward slashes. The ` + "`" + `namespace` + "`" + ` is always relative to the provider's configured [namespace](/docs/providers/vault#namespace).`,
				},
				resource.Attribute{
					Name:        "backend",
					Description: `(Required) The path the AD secret backend is mounted at, with no leading or trailing ` + "`" + `/` + "`" + `s.`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `(Required) The name to identify this role within the backend. Must be unique within the backend.`,
				},
				resource.Attribute{
					Name:        "service_account_name",
					Description: `(Required) Specifies the name of the Active Directory service account mapped to this role.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `(Optional) The password time-to-live in seconds. Defaults to the configuration ttl if not provided. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "last_vault_rotation",
					Description: `Timestamp of the last password rotation by Vault.`,
				},
				resource.Attribute{
					Name:        "password_last_set",
					Description: `Timestamp of the last password set by Vault. ## Import AD secret backend roles can be imported using the ` + "`" + `path` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_ad_secret_role.role ad/roles/bob ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "last_vault_rotation",
					Description: `Timestamp of the last password rotation by Vault.`,
				},
				resource.Attribute{
					Name:        "password_last_set",
					Description: `Timestamp of the last password set by Vault. ## Import AD secret backend roles can be imported using the ` + "`" + `path` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_ad_secret_role.role ad/roles/bob ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_alicloud_auth_backend_role",
			Category:         "Resources",
			ShortDescription: `Managing roles in an AliCloud auth backend in Vault`,
			Description:      ``,
			Keywords: []string{
				"alicloud",
				"auth",
				"backend",
				"role",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) The namespace to provision the resource in. The value should not contain leading or trailing forward slashes. The ` + "`" + `namespace` + "`" + ` is always relative to the provider's configured [namespace](../index.html#namespace).`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `(Required; Forces new resource) Name of the role. Must correspond with the name of the role reflected in the arn.`,
				},
				resource.Attribute{
					Name:        "arn",
					Description: `(Required) The role's arn.`,
				},
				resource.Attribute{
					Name:        "backend",
					Description: `(Optional; Forces new resource) Path to the mounted AliCloud auth backend. Defaults to ` + "`" + `alicloud` + "`" + ` For more details on the usage of each argument consult the [Vault AliCloud API documentation](https://www.vaultproject.io/api-docs/auth/alicloud). ### Common Token Arguments These arguments are common across several Authentication Token resources since Vault 1.2.`,
				},
				resource.Attribute{
					Name:        "token_ttl",
					Description: `(Optional) The incremental lifetime for generated tokens in number of seconds. Its current value will be referenced at renewal time.`,
				},
				resource.Attribute{
					Name:        "token_max_ttl",
					Description: `(Optional) The maximum lifetime for generated tokens in number of seconds. Its current value will be referenced at renewal time.`,
				},
				resource.Attribute{
					Name:        "token_period",
					Description: `(Optional) If set, indicates that the token generated using this role should never expire. The token should be renewed within the duration specified by this value. At each renewal, the token's TTL will be set to the value of this field. Specified in seconds.`,
				},
				resource.Attribute{
					Name:        "token_policies",
					Description: `(Optional) List of policies to encode onto generated tokens. Depending on the auth method, this list may be supplemented by user/group/other values.`,
				},
				resource.Attribute{
					Name:        "token_bound_cidrs",
					Description: `(Optional) List of CIDR blocks; if set, specifies blocks of IP addresses which can authenticate successfully, and ties the resulting token to these blocks as well.`,
				},
				resource.Attribute{
					Name:        "token_explicit_max_ttl",
					Description: `(Optional) If set, will encode an [explicit max TTL](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls) onto the token in number of seconds. This is a hard cap even if ` + "`" + `token_ttl` + "`" + ` and ` + "`" + `token_max_ttl` + "`" + ` would otherwise allow a renewal.`,
				},
				resource.Attribute{
					Name:        "token_no_default_policy",
					Description: `(Optional) If set, the default policy will not be set on generated tokens; otherwise it will be added to the policies set in token_policies.`,
				},
				resource.Attribute{
					Name:        "token_num_uses",
					Description: `(Optional) The [maximum number](https://www.vaultproject.io/api-docs/auth/alicloud#token_num_uses) of times a generated token may be used (within its lifetime); 0 means unlimited.`,
				},
				resource.Attribute{
					Name:        "token_type",
					Description: `(Optional) The type of token that should be generated. Can be ` + "`" + `service` + "`" + `, ` + "`" + `batch` + "`" + `, or ` + "`" + `default` + "`" + ` to use the mount's tuned default (which unless changed will be ` + "`" + `service` + "`" + ` tokens). For token store roles, there are two additional possibilities: ` + "`" + `default-service` + "`" + ` and ` + "`" + `default-batch` + "`" + ` which specify the type to return unless the client requests a different type at generation time. ## Attribute Reference No additional attributes are exposed by this resource. ## Import Alicloud authentication roles can be imported using the ` + "`" + `path` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_alicloud_auth_backend_role.my_role auth/alicloud/role/my_role ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_approle_auth_backend_login",
			Category:         "Resources",
			ShortDescription: `Log into Vault using the AppRole auth backend.`,
			Description:      ``,
			Keywords: []string{
				"approle",
				"auth",
				"backend",
				"login",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) The namespace to provision the resource in. The value should not contain leading or trailing forward slashes. The ` + "`" + `namespace` + "`" + ` is always relative to the provider's configured [namespace](/docs/providers/vault#namespace).`,
				},
				resource.Attribute{
					Name:        "role_id",
					Description: `(Required) The ID of the role to log in with.`,
				},
				resource.Attribute{
					Name:        "secret_id",
					Description: `(Optional) The secret ID of the role to log in with. Required unless ` + "`" + `bind_secret_id` + "`" + ` is set to false on the role.`,
				},
				resource.Attribute{
					Name:        "backend",
					Description: `The unique path of the Vault backend to log in with. ## Attributes Reference In addition to the fields above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "policies",
					Description: `A list of policies applied to the token.`,
				},
				resource.Attribute{
					Name:        "renewable",
					Description: `Whether the token is renewable or not.`,
				},
				resource.Attribute{
					Name:        "lease_duration",
					Description: `How long the token is valid for, in seconds.`,
				},
				resource.Attribute{
					Name:        "lease_started",
					Description: `The date and time the lease started, in RFC 3339 format.`,
				},
				resource.Attribute{
					Name:        "accessor",
					Description: `The accessor for the token.`,
				},
				resource.Attribute{
					Name:        "client_token",
					Description: `The Vault token created.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `The metadata associated with the token.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "policies",
					Description: `A list of policies applied to the token.`,
				},
				resource.Attribute{
					Name:        "renewable",
					Description: `Whether the token is renewable or not.`,
				},
				resource.Attribute{
					Name:        "lease_duration",
					Description: `How long the token is valid for, in seconds.`,
				},
				resource.Attribute{
					Name:        "lease_started",
					Description: `The date and time the lease started, in RFC 3339 format.`,
				},
				resource.Attribute{
					Name:        "accessor",
					Description: `The accessor for the token.`,
				},
				resource.Attribute{
					Name:        "client_token",
					Description: `The Vault token created.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `The metadata associated with the token.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_approle_auth_backend_role",
			Category:         "Resources",
			ShortDescription: `Manages AppRole auth backend roles in Vault.`,
			Description:      ``,
			Keywords: []string{
				"approle",
				"auth",
				"backend",
				"role",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) The namespace to provision the resource in. The value should not contain leading or trailing forward slashes. The ` + "`" + `namespace` + "`" + ` is always relative to the provider's configured [namespace](/docs/providers/vault#namespace).`,
				},
				resource.Attribute{
					Name:        "role_name",
					Description: `(Required) The name of the role.`,
				},
				resource.Attribute{
					Name:        "role_id",
					Description: `(Optional) The RoleID of this role. If not specified, one will be auto-generated.`,
				},
				resource.Attribute{
					Name:        "bind_secret_id",
					Description: `(Optional) Whether or not to require ` + "`" + `secret_id` + "`" + ` to be presented when logging in using this AppRole. Defaults to ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "secret_id_bound_cidrs",
					Description: `(Optional) If set, specifies blocks of IP addresses which can perform the login operation.`,
				},
				resource.Attribute{
					Name:        "secret_id_num_uses",
					Description: `(Optional) The number of times any particular SecretID can be used to fetch a token from this AppRole, after which the SecretID will expire. A value of zero will allow unlimited uses.`,
				},
				resource.Attribute{
					Name:        "secret_id_ttl",
					Description: `(Optional) The number of seconds after which any SecretID expires.`,
				},
				resource.Attribute{
					Name:        "backend",
					Description: `(Optional) The unique name of the auth backend to configure. Defaults to ` + "`" + `approle` + "`" + `. ### Common Token Arguments These arguments are common across several Authentication Token resources since Vault 1.2.`,
				},
				resource.Attribute{
					Name:        "token_ttl",
					Description: `(Optional) The incremental lifetime for generated tokens in number of seconds. Its current value will be referenced at renewal time.`,
				},
				resource.Attribute{
					Name:        "token_max_ttl",
					Description: `(Optional) The maximum lifetime for generated tokens in number of seconds. Its current value will be referenced at renewal time.`,
				},
				resource.Attribute{
					Name:        "token_period",
					Description: `(Optional) If set, indicates that the token generated using this role should never expire. The token should be renewed within the duration specified by this value. At each renewal, the token's TTL will be set to the value of this field. Specified in seconds.`,
				},
				resource.Attribute{
					Name:        "token_policies",
					Description: `(Optional) List of policies to encode onto generated tokens. Depending on the auth method, this list may be supplemented by user/group/other values.`,
				},
				resource.Attribute{
					Name:        "token_bound_cidrs",
					Description: `(Optional) List of CIDR blocks; if set, specifies blocks of IP addresses which can authenticate successfully, and ties the resulting token to these blocks as well.`,
				},
				resource.Attribute{
					Name:        "token_explicit_max_ttl",
					Description: `(Optional) If set, will encode an [explicit max TTL](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls) onto the token in number of seconds. This is a hard cap even if ` + "`" + `token_ttl` + "`" + ` and ` + "`" + `token_max_ttl` + "`" + ` would otherwise allow a renewal.`,
				},
				resource.Attribute{
					Name:        "token_no_default_policy",
					Description: `(Optional) If set, the default policy will not be set on generated tokens; otherwise it will be added to the policies set in token_policies.`,
				},
				resource.Attribute{
					Name:        "token_num_uses",
					Description: `(Optional) The [maximum number](https://www.vaultproject.io/api-docs/auth/approle#token_num_uses) of times a generated token may be used (within its lifetime); 0 means unlimited.`,
				},
				resource.Attribute{
					Name:        "token_type",
					Description: `(Optional) The type of token that should be generated. Can be ` + "`" + `service` + "`" + `, ` + "`" + `batch` + "`" + `, or ` + "`" + `default` + "`" + ` to use the mount's tuned default (which unless changed will be ` + "`" + `service` + "`" + ` tokens). For token store roles, there are two additional possibilities: ` + "`" + `default-service` + "`" + ` and ` + "`" + `default-batch` + "`" + ` which specify the type to return unless the client requests a different type at generation time. ## Attributes Reference No additional attributes are exported by this resource. ## Import AppRole authentication backend roles can be imported using the ` + "`" + `path` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_approle_auth_backend_role.example auth/approle/role/test-role ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_approle_auth_backend_role_secret_id",
			Category:         "Resources",
			ShortDescription: `Manages AppRole auth backend role SecretIDs in Vault.`,
			Description:      ``,
			Keywords: []string{
				"approle",
				"auth",
				"backend",
				"role",
				"secret",
				"id",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) The namespace to provision the resource in. The value should not contain leading or trailing forward slashes. The ` + "`" + `namespace` + "`" + ` is always relative to the provider's configured [namespace](/docs/providers/vault#namespace).`,
				},
				resource.Attribute{
					Name:        "role_name",
					Description: `(Required) The name of the role to create the SecretID for.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `(Optional) A JSON-encoded string containing metadata in key-value pairs to be set on tokens issued with this SecretID.`,
				},
				resource.Attribute{
					Name:        "cidr_list",
					Description: `(Optional) If set, specifies blocks of IP addresses which can perform the login operation using this SecretID.`,
				},
				resource.Attribute{
					Name:        "secret_id",
					Description: `(Optional) The SecretID to be created. If set, uses "Push" mode. Defaults to Vault auto-generating SecretIDs.`,
				},
				resource.Attribute{
					Name:        "wrapping_ttl",
					Description: `(Optional) If set, the SecretID response will be [response-wrapped](https://www.vaultproject.io/docs/concepts/response-wrapping) and available for the duration specified. Only a single unwrapping of the token is allowed.`,
				},
				resource.Attribute{
					Name:        "with_wrapped_accessor",
					Description: `(Optional) Set to ` + "`" + `true` + "`" + ` to use the wrapped secret-id accessor as the resource ID. If ` + "`" + `false` + "`" + ` (default value), a fresh secret ID will be regenerated whenever the wrapping token is expired or invalidated through unwrapping. ## Attributes Reference In addition to the fields above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "accessor",
					Description: `The unique ID for this SecretID that can be safely logged.`,
				},
				resource.Attribute{
					Name:        "wrapping_accessor",
					Description: `The unique ID for the response-wrapped SecretID that can be safely logged.`,
				},
				resource.Attribute{
					Name:        "wrapping_token",
					Description: `The token used to retrieve a response-wrapped SecretID.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "accessor",
					Description: `The unique ID for this SecretID that can be safely logged.`,
				},
				resource.Attribute{
					Name:        "wrapping_accessor",
					Description: `The unique ID for the response-wrapped SecretID that can be safely logged.`,
				},
				resource.Attribute{
					Name:        "wrapping_token",
					Description: `The token used to retrieve a response-wrapped SecretID.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_audit",
			Category:         "Resources",
			ShortDescription: `Writes audit backends for Vault`,
			Description:      ``,
			Keywords: []string{
				"audit",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) The namespace to provision the resource in. The value should not contain leading or trailing forward slashes. The ` + "`" + `namespace` + "`" + ` is always relative to the provider's configured [namespace](/docs/providers/vault#namespace).`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) Type of the audit device, such as 'file'.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(optional) The path to mount the audit device. This defaults to the type.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Human-friendly description of the audit device.`,
				},
				resource.Attribute{
					Name:        "local",
					Description: `(Optional) Specifies if the audit device is a local only. Local audit devices are not replicated nor (if a secondary) removed by replication.`,
				},
				resource.Attribute{
					Name:        "options",
					Description: `(Required) Configuration options to pass to the audit device itself. For a reference of the device types and their options, consult the [Vault documentation.](https://www.vaultproject.io/docs/audit/index.html) ## Attributes Reference No additional attributes are exported by this resource. ## Import Audit devices can be imported using the ` + "`" + `path` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_audit.test syslog ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_audit_request_header",
			Category:         "Resources",
			ShortDescription: `Manages audited request headers in Vault`,
			Description:      ``,
			Keywords: []string{
				"audit",
				"request",
				"header",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the request header to audit.`,
				},
				resource.Attribute{
					Name:        "hmac",
					Description: `(Optional) Whether this header's value should be HMAC'd in the audit logs. ## Attributes Reference No additional attributes are exported by this resource.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_auth_backend",
			Category:         "Resources",
			ShortDescription: `Writes auth methods for Vault`,
			Description:      ``,
			Keywords: []string{
				"auth",
				"backend",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) The namespace to provision the resource in. The value should not contain leading or trailing forward slashes. The ` + "`" + `namespace` + "`" + ` is always relative to the provider's configured [namespace](/docs/providers/vault#namespace).`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) The name of the auth method type.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Optional) The path to mount the auth method — this defaults to the name of the type.`,
				},
				resource.Attribute{
					Name:        "disable_remount",
					Description: `(Optional) If set, opts out of mount migration on path updates. See here for more info on [Mount Migration](https://www.vaultproject.io/docs/concepts/mount-migration)`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A description of the auth method.`,
				},
				resource.Attribute{
					Name:        "local",
					Description: `(Optional) Specifies if the auth method is local only.`,
				},
				resource.Attribute{
					Name:        "tune",
					Description: `(Optional) Extra configuration block. Structure is documented below. The ` + "`" + `tune` + "`" + ` block is used to tune the auth backend:`,
				},
				resource.Attribute{
					Name:        "default_lease_ttl",
					Description: `(Optional) Specifies the default time-to-live. If set, this overrides the global default. Must be a valid [duration string](https://golang.org/pkg/time/#ParseDuration)`,
				},
				resource.Attribute{
					Name:        "max_lease_ttl",
					Description: `(Optional) Specifies the maximum time-to-live. If set, this overrides the global default. Must be a valid [duration string](https://golang.org/pkg/time/#ParseDuration)`,
				},
				resource.Attribute{
					Name:        "audit_non_hmac_response_keys",
					Description: `(Optional) Specifies the list of keys that will not be HMAC'd by audit devices in the response data object.`,
				},
				resource.Attribute{
					Name:        "audit_non_hmac_request_keys",
					Description: `(Optional) Specifies the list of keys that will not be HMAC'd by audit devices in the request data object.`,
				},
				resource.Attribute{
					Name:        "listing_visibility",
					Description: `(Optional) Specifies whether to show this mount in the UI-specific listing endpoint. Valid values are "unauth" or "hidden".`,
				},
				resource.Attribute{
					Name:        "passthrough_request_headers",
					Description: `(Optional) List of headers to whitelist and pass from the request to the backend.`,
				},
				resource.Attribute{
					Name:        "allowed_response_headers",
					Description: `(Optional) List of headers to whitelist and allowing a plugin to include them in the response.`,
				},
				resource.Attribute{
					Name:        "token_type",
					Description: `(Optional) Specifies the type of tokens that should be returned by the mount. Valid values are "default-service", "default-batch", "service", "batch". ## Attributes Reference In addition to the fields above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "accessor",
					Description: `The accessor for this auth method ## Import Auth methods can be imported using the ` + "`" + `path` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_auth_backend.example github ` + "`" + `` + "`" + `` + "`" + ` ## Tutorials Refer to the following tutorials for additional usage examples: - [Codify Management of Vault Enterprise Using Terraform](https://learn.hashicorp.com/tutorials/vault/codify-mgmt-enterprise) - [Codify Management of Vault Using Terraform](https://learn.hashicorp.com/tutorials/vault/codify-mgmt-oss)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "accessor",
					Description: `The accessor for this auth method ## Import Auth methods can be imported using the ` + "`" + `path` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_auth_backend.example github ` + "`" + `` + "`" + `` + "`" + ` ## Tutorials Refer to the following tutorials for additional usage examples: - [Codify Management of Vault Enterprise Using Terraform](https://learn.hashicorp.com/tutorials/vault/codify-mgmt-enterprise) - [Codify Management of Vault Using Terraform](https://learn.hashicorp.com/tutorials/vault/codify-mgmt-oss)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_aws_auth_backend_cert",
			Category:         "Resources",
			ShortDescription: `Manages a certificate for an AWS Auth Backend in Vault.`,
			Description:      ``,
			Keywords: []string{
				"aws",
				"auth",
				"backend",
				"cert",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) The namespace to provision the resource in. The value should not contain leading or trailing forward slashes. The ` + "`" + `namespace` + "`" + ` is always relative to the provider's configured [namespace](/docs/providers/vault#namespace).`,
				},
				resource.Attribute{
					Name:        "cert_name",
					Description: `(Required) The name of the certificate.`,
				},
				resource.Attribute{
					Name:        "aws_public_cert",
					Description: `(Required) The Base64 encoded AWS Public key required to verify PKCS7 signature of the EC2 instance metadata. You can find this key in the [AWS documentation](http://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-identity-documents.html).`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) Either "pkcs7" or "identity", indicating the type of document which can be verified using the given certificate. Defaults to "pkcs7".`,
				},
				resource.Attribute{
					Name:        "backend",
					Description: `(Optional) The path the AWS auth backend being configured was mounted at. Defaults to ` + "`" + `aws` + "`" + `. ## Attributes Reference No additional attributes are exported by this resource. ## Import AWS auth backend certificates can be imported using ` + "`" + `auth/` + "`" + `, the ` + "`" + `backend` + "`" + ` path, ` + "`" + `/config/certificate/` + "`" + `, and the ` + "`" + `cert_name` + "`" + ` e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_aws_auth_backend_cert.example auth/aws/config/certificate/my-cert ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_aws_auth_backend_client",
			Category:         "Resources",
			ShortDescription: `Configures the client used by an AWS Auth Backend in Vault.`,
			Description:      ``,
			Keywords: []string{
				"aws",
				"auth",
				"backend",
				"client",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) The namespace to provision the resource in. The value should not contain leading or trailing forward slashes. The ` + "`" + `namespace` + "`" + ` is always relative to the provider's configured [namespace](/docs/providers/vault#namespace).`,
				},
				resource.Attribute{
					Name:        "backend",
					Description: `(Optional) The path the AWS auth backend being configured was mounted at. Defaults to ` + "`" + `aws` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "access_key",
					Description: `(Optional) The AWS access key that Vault should use for the auth backend.`,
				},
				resource.Attribute{
					Name:        "secret_key",
					Description: `(Optional) The AWS secret key that Vault should use for the auth backend.`,
				},
				resource.Attribute{
					Name:        "ec2_endpoint",
					Description: `(Optional) Override the URL Vault uses when making EC2 API calls.`,
				},
				resource.Attribute{
					Name:        "iam_endpoint",
					Description: `(Optional) Override the URL Vault uses when making IAM API calls.`,
				},
				resource.Attribute{
					Name:        "sts_endpoint",
					Description: `(Optional) Override the URL Vault uses when making STS API calls.`,
				},
				resource.Attribute{
					Name:        "sts_region",
					Description: `(Optional) Override the default region when making STS API calls. The ` + "`" + `sts_endpoint` + "`" + ` argument must be set when using ` + "`" + `sts_region` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "iam_server_id_header_value",
					Description: `(Optional) The value to require in the ` + "`" + `X-Vault-AWS-IAM-Server-ID` + "`" + ` header as part of ` + "`" + `GetCallerIdentity` + "`" + ` requests that are used in the IAM auth method. ## Attributes Reference No additional attributes are exported by this resource. ## Import AWS auth backend clients can be imported using ` + "`" + `auth/` + "`" + `, the ` + "`" + `backend` + "`" + ` path, and ` + "`" + `/config/client` + "`" + ` e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_aws_auth_backend_client.example auth/aws/config/client ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_aws_auth_backend_config_identity",
			Category:         "Resources",
			ShortDescription: `Manages AWS auth backend identity configuration in Vault.`,
			Description:      ``,
			Keywords: []string{
				"aws",
				"auth",
				"backend",
				"config",
				"identity",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) The namespace to provision the resource in. The value should not contain leading or trailing forward slashes. The ` + "`" + `namespace` + "`" + ` is always relative to the provider's configured [namespace](/docs/providers/vault#namespace).`,
				},
				resource.Attribute{
					Name:        "iam_alias",
					Description: `(Optional) How to generate the identity alias when using the iam auth method. Valid choices are ` + "`" + `role_id` + "`" + `, ` + "`" + `unique_id` + "`" + `, and ` + "`" + `full_arn` + "`" + `. Defaults to ` + "`" + `role_id` + "`" + ``,
				},
				resource.Attribute{
					Name:        "iam_metadata",
					Description: `(Optional) The metadata to include on the token returned by the ` + "`" + `login` + "`" + ` endpoint. This metadata will be added to both audit logs, and on the ` + "`" + `iam_alias` + "`" + ``,
				},
				resource.Attribute{
					Name:        "ec2_alias",
					Description: `(Optional) How to generate the identity alias when using the ec2 auth method. Valid choices are ` + "`" + `role_id` + "`" + `, ` + "`" + `instance_id` + "`" + `, and ` + "`" + `image_id` + "`" + `. Defaults to ` + "`" + `role_id` + "`" + ``,
				},
				resource.Attribute{
					Name:        "ec2_metadata",
					Description: `(Optional) The metadata to include on the token returned by the ` + "`" + `login` + "`" + ` endpoint. This metadata will be added to both audit logs, and on the ` + "`" + `ec2_alias` + "`" + ` ## Attributes Reference No additional attributes are exported by this resource. ## Import AWS auth backend identity config can be imported using ` + "`" + `auth/` + "`" + `, the ` + "`" + `backend` + "`" + ` path, and ` + "`" + `/config/identity` + "`" + ` e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_aws_auth_backend_role.example auth/aws/config/identity ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_aws_auth_backend_identity_whitelist",
			Category:         "Resources",
			ShortDescription: `Configures the periodic tidying operation of the whitelisted identity entries.`,
			Description:      ``,
			Keywords: []string{
				"aws",
				"auth",
				"backend",
				"identity",
				"whitelist",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) The namespace to provision the resource in. The value should not contain leading or trailing forward slashes. The ` + "`" + `namespace` + "`" + ` is always relative to the provider's configured [namespace](/docs/providers/vault#namespace).`,
				},
				resource.Attribute{
					Name:        "backend",
					Description: `(Optional) The path of the AWS backend being configured.`,
				},
				resource.Attribute{
					Name:        "safety_buffer",
					Description: `(Optional) The amount of extra time, in minutes, that must have passed beyond the roletag expiration, before it is removed from the backend storage.`,
				},
				resource.Attribute{
					Name:        "disable_periodic_tidy",
					Description: `(Optional) If set to true, disables the periodic tidying of the identity-whitelist entries. ## Attributes Reference No additional attributes are exported by this resource. ## Import AWS auth backend identity whitelists can be imported using ` + "`" + `auth/` + "`" + `, the ` + "`" + `backend` + "`" + ` path, and ` + "`" + `/config/tidy/identity-whitelist` + "`" + ` e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_aws_auth_backend_identity_whitelist.example auth/aws/config/tidy/identity-whitelist ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_aws_auth_backend_login",
			Category:         "Resources",
			ShortDescription: `Manages Vault tokens acquired using the AWS auth backend.`,
			Description:      ``,
			Keywords: []string{
				"aws",
				"auth",
				"backend",
				"login",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) The namespace to provision the resource in. The value should not contain leading or trailing forward slashes. The ` + "`" + `namespace` + "`" + ` is always relative to the provider's configured [namespace](/docs/providers/vault#namespace).`,
				},
				resource.Attribute{
					Name:        "backend",
					Description: `(Optional) The unique name of the AWS auth backend. Defaults to 'aws'.`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `(Optional) The name of the AWS auth backend role to create tokens against.`,
				},
				resource.Attribute{
					Name:        "identity",
					Description: `(Optional) The base64-encoded EC2 instance identity document to authenticate with. Can be retrieved from the EC2 metadata server.`,
				},
				resource.Attribute{
					Name:        "signature",
					Description: `(Optional) The base64-encoded SHA256 RSA signature of the instance identity document to authenticate with, with all newline characters removed. Can be retrieved from the EC2 metadata server.`,
				},
				resource.Attribute{
					Name:        "pkcs7",
					Description: `(Optional) The PKCS#7 signature of the identity document to authenticate with, with all newline characters removed. Can be retrieved from the EC2 metadata server.`,
				},
				resource.Attribute{
					Name:        "nonce",
					Description: `(Optional) The unique nonce to be used for login requests. Can be set to a user-specified value, or will contain the server-generated value once a token is issued. EC2 instances can only acquire a single token until the whitelist is tidied again unless they keep track of this nonce.`,
				},
				resource.Attribute{
					Name:        "iam_http_request_method",
					Description: `(Optional) The HTTP method used in the signed IAM request.`,
				},
				resource.Attribute{
					Name:        "iam_request_url",
					Description: `(Optional) The base64-encoded HTTP URL used in the signed request.`,
				},
				resource.Attribute{
					Name:        "iam_request_body",
					Description: `(Optional) The base64-encoded body of the signed request.`,
				},
				resource.Attribute{
					Name:        "iam_request_headers",
					Description: `(Optional) The base64-encoded, JSON serialized representation of the GetCallerIdentity HTTP request headers. ## Attributes Reference In addition to the fields above, the following attributes are also exposed:`,
				},
				resource.Attribute{
					Name:        "lease_duration",
					Description: `The duration in seconds the token will be valid, relative to the time in ` + "`" + `lease_start_time` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "lease_start_time",
					Description: `The approximate time at which the token was created, using the clock of the system where Terraform was running.`,
				},
				resource.Attribute{
					Name:        "renewable",
					Description: `Set to true if the token can be extended through renewal.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `A map of information returned by the Vault server about the authentication used to generate this token.`,
				},
				resource.Attribute{
					Name:        "auth_type",
					Description: `The authentication type used to generate this token.`,
				},
				resource.Attribute{
					Name:        "policies",
					Description: `The Vault policies assigned to this token.`,
				},
				resource.Attribute{
					Name:        "accessor",
					Description: `The token's accessor.`,
				},
				resource.Attribute{
					Name:        "client_token",
					Description: `The token returned by Vault.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "lease_duration",
					Description: `The duration in seconds the token will be valid, relative to the time in ` + "`" + `lease_start_time` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "lease_start_time",
					Description: `The approximate time at which the token was created, using the clock of the system where Terraform was running.`,
				},
				resource.Attribute{
					Name:        "renewable",
					Description: `Set to true if the token can be extended through renewal.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `A map of information returned by the Vault server about the authentication used to generate this token.`,
				},
				resource.Attribute{
					Name:        "auth_type",
					Description: `The authentication type used to generate this token.`,
				},
				resource.Attribute{
					Name:        "policies",
					Description: `The Vault policies assigned to this token.`,
				},
				resource.Attribute{
					Name:        "accessor",
					Description: `The token's accessor.`,
				},
				resource.Attribute{
					Name:        "client_token",
					Description: `The token returned by Vault.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_aws_auth_backend_role",
			Category:         "Resources",
			ShortDescription: `Manages AWS auth backend roles in Vault.`,
			Description:      ``,
			Keywords: []string{
				"aws",
				"auth",
				"backend",
				"role",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) The namespace to provision the resource in. The value should not contain leading or trailing forward slashes. The ` + "`" + `namespace` + "`" + ` is always relative to the provider's configured [namespace](/docs/providers/vault#namespace).`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `(Required) The name of the role.`,
				},
				resource.Attribute{
					Name:        "auth_type",
					Description: `(Optional) The auth type permitted for this role. Valid choices are ` + "`" + `ec2` + "`" + ` and ` + "`" + `iam` + "`" + `. Defaults to ` + "`" + `iam` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "bound_ami_ids",
					Description: `(Optional) If set, defines a constraint on the EC2 instances that can perform the login operation that they should be using the AMI ID specified by this field. ` + "`" + `auth_type` + "`" + ` must be set to ` + "`" + `ec2` + "`" + ` or ` + "`" + `inferred_entity_type` + "`" + ` must be set to ` + "`" + `ec2_instance` + "`" + ` to use this constraint.`,
				},
				resource.Attribute{
					Name:        "bound_account_ids",
					Description: `(Optional) If set, defines a constraint on the EC2 instances that can perform the login operation that they should be using the account ID specified by this field. ` + "`" + `auth_type` + "`" + ` must be set to ` + "`" + `ec2` + "`" + ` or ` + "`" + `inferred_entity_type` + "`" + ` must be set to ` + "`" + `ec2_instance` + "`" + ` to use this constraint.`,
				},
				resource.Attribute{
					Name:        "bound_regions",
					Description: `(Optional) If set, defines a constraint on the EC2 instances that can perform the login operation that the region in their identity document must match the one specified by this field. ` + "`" + `auth_type` + "`" + ` must be set to ` + "`" + `ec2` + "`" + ` or ` + "`" + `inferred_entity_type` + "`" + ` must be set to ` + "`" + `ec2_instance` + "`" + ` to use this constraint.`,
				},
				resource.Attribute{
					Name:        "bound_vpc_ids",
					Description: `(Optional) If set, defines a constraint on the EC2 instances that can perform the login operation that they be associated with the VPC ID that matches the value specified by this field. ` + "`" + `auth_type` + "`" + ` must be set to ` + "`" + `ec2` + "`" + ` or ` + "`" + `inferred_entity_type` + "`" + ` must be set to ` + "`" + `ec2_instance` + "`" + ` to use this constraint.`,
				},
				resource.Attribute{
					Name:        "bound_subnet_ids",
					Description: `(Optional) If set, defines a constraint on the EC2 instances that can perform the login operation that they be associated with the subnet ID that matches the value specified by this field. ` + "`" + `auth_type` + "`" + ` must be set to ` + "`" + `ec2` + "`" + ` or ` + "`" + `inferred_entity_type` + "`" + ` must be set to ` + "`" + `ec2_instance` + "`" + ` to use this constraint.`,
				},
				resource.Attribute{
					Name:        "bound_iam_role_arns",
					Description: `(Optional) If set, defines a constraint on the EC2 instances that can perform the login operation that they must match the IAM role ARN specified by this field. ` + "`" + `auth_type` + "`" + ` must be set to ` + "`" + `ec2` + "`" + ` or ` + "`" + `inferred_entity_type` + "`" + ` must be set to ` + "`" + `ec2_instance` + "`" + ` to use this constraint.`,
				},
				resource.Attribute{
					Name:        "bound_iam_instance_profile_arns",
					Description: `(Optional) If set, defines a constraint on the EC2 instances that can perform the login operation that they must be associated with an IAM instance profile ARN which has a prefix that matches the value specified by this field. The value is prefix-matched as though it were a glob ending in ` + "`" + ``,
				},
				resource.Attribute{
					Name:        "role_tag",
					Description: `(Optional) If set, enable role tags for this role. The value set for this field should be the key of the tag on the EC2 instance. ` + "`" + `auth_type` + "`" + ` must be set to ` + "`" + `ec2` + "`" + ` or ` + "`" + `inferred_entity_type` + "`" + ` must be set to ` + "`" + `ec2_instance` + "`" + ` to use this constraint.`,
				},
				resource.Attribute{
					Name:        "bound_iam_principal_arns",
					Description: `(Optional) If set, defines the IAM principal that must be authenticated when ` + "`" + `auth_type` + "`" + ` is set to ` + "`" + `iam` + "`" + `. Wildcards are supported at the end of the ARN.`,
				},
				resource.Attribute{
					Name:        "inferred_entity_type",
					Description: `(Optional) If set, instructs Vault to turn on inferencing. The only valid value is ` + "`" + `ec2_instance` + "`" + `, which instructs Vault to infer that the role comes from an EC2 instance in an IAM instance profile. This only applies when ` + "`" + `auth_type` + "`" + ` is set to ` + "`" + `iam` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "inferred_aws_region",
					Description: `(Optional) When ` + "`" + `inferred_entity_type` + "`" + ` is set, this is the region to search for the inferred entities. Required if ` + "`" + `inferred_entity_type` + "`" + ` is set. This only applies when ` + "`" + `auth_type` + "`" + ` is set to ` + "`" + `iam` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "resolve_aws_unique_ids",
					Description: `(Optional, Forces new resource) Only valid when ` + "`" + `auth_type` + "`" + ` is ` + "`" + `iam` + "`" + `. If set to ` + "`" + `true` + "`" + `, the ` + "`" + `bound_iam_principal_arns` + "`" + ` are resolved to [AWS Unique IDs](http://docs.aws.amazon.com/IAM/latest/UserGuide/reference_identifiers.html#identifiers-unique-ids) for the bound principal ARN. This field is ignored when a ` + "`" + `bound_iam_principal_arn` + "`" + ` ends in a wildcard. Resolving to unique IDs more closely mimics the behavior of AWS services in that if an IAM user or role is deleted and a new one is recreated with the same name, those new users or roles won't get access to roles in Vault that were permissioned to the prior principals of the same name. Defaults to ` + "`" + `true` + "`" + `. Once set to ` + "`" + `true` + "`" + `, this cannot be changed to ` + "`" + `false` + "`" + ` without recreating the role.`,
				},
				resource.Attribute{
					Name:        "allow_instance_migration",
					Description: `(Optional) If set to ` + "`" + `true` + "`" + `, allows migration of the underlying instance where the client resides.`,
				},
				resource.Attribute{
					Name:        "disallow_reauthentication",
					Description: `(Optional) IF set to ` + "`" + `true` + "`" + `, only allows a single token to be granted per instance ID. This can only be set when ` + "`" + `auth_type` + "`" + ` is set to ` + "`" + `ec2` + "`" + `. ### Common Token Arguments These arguments are common across several Authentication Token resources since Vault 1.2.`,
				},
				resource.Attribute{
					Name:        "token_ttl",
					Description: `(Optional) The incremental lifetime for generated tokens in number of seconds. Its current value will be referenced at renewal time.`,
				},
				resource.Attribute{
					Name:        "token_max_ttl",
					Description: `(Optional) The maximum lifetime for generated tokens in number of seconds. Its current value will be referenced at renewal time.`,
				},
				resource.Attribute{
					Name:        "token_period",
					Description: `(Optional) If set, indicates that the token generated using this role should never expire. The token should be renewed within the duration specified by this value. At each renewal, the token's TTL will be set to the value of this field. Specified in seconds.`,
				},
				resource.Attribute{
					Name:        "token_policies",
					Description: `(Optional) List of policies to encode onto generated tokens. Depending on the auth method, this list may be supplemented by user/group/other values.`,
				},
				resource.Attribute{
					Name:        "token_bound_cidrs",
					Description: `(Optional) List of CIDR blocks; if set, specifies blocks of IP addresses which can authenticate successfully, and ties the resulting token to these blocks as well.`,
				},
				resource.Attribute{
					Name:        "token_explicit_max_ttl",
					Description: `(Optional) If set, will encode an [explicit max TTL](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls) onto the token in number of seconds. This is a hard cap even if ` + "`" + `token_ttl` + "`" + ` and ` + "`" + `token_max_ttl` + "`" + ` would otherwise allow a renewal.`,
				},
				resource.Attribute{
					Name:        "token_no_default_policy",
					Description: `(Optional) If set, the default policy will not be set on generated tokens; otherwise it will be added to the policies set in token_policies.`,
				},
				resource.Attribute{
					Name:        "token_num_uses",
					Description: `(Optional) The [maximum number](https://www.vaultproject.io/api-docs/auth/aws#token_num_uses) of times a generated token may be used (within its lifetime); 0 means unlimited.`,
				},
				resource.Attribute{
					Name:        "token_type",
					Description: `(Optional) The type of token that should be generated. Can be ` + "`" + `service` + "`" + `, ` + "`" + `batch` + "`" + `, or ` + "`" + `default` + "`" + ` to use the mount's tuned default (which unless changed will be ` + "`" + `service` + "`" + ` tokens). For token store roles, there are two additional possibilities: ` + "`" + `default-service` + "`" + ` and ` + "`" + `default-batch` + "`" + ` which specify the type to return unless the client requests a different type at generation time. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "role_id",
					Description: `The Vault generated role ID. ## Import AWS auth backend roles can be imported using ` + "`" + `auth/` + "`" + `, the ` + "`" + `backend` + "`" + ` path, ` + "`" + `/role/` + "`" + `, and the ` + "`" + `role` + "`" + ` name e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_aws_auth_backend_role.example auth/aws/role/test-role ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "role_id",
					Description: `The Vault generated role ID. ## Import AWS auth backend roles can be imported using ` + "`" + `auth/` + "`" + `, the ` + "`" + `backend` + "`" + ` path, ` + "`" + `/role/` + "`" + `, and the ` + "`" + `role` + "`" + ` name e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_aws_auth_backend_role.example auth/aws/role/test-role ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_aws_auth_backend_role_tag",
			Category:         "Resources",
			ShortDescription: `Reads role tags from a Vault AWS auth backend.`,
			Description:      ``,
			Keywords: []string{
				"aws",
				"auth",
				"backend",
				"role",
				"tag",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) The namespace to provision the resource in. The value should not contain leading or trailing forward slashes. The ` + "`" + `namespace` + "`" + ` is always relative to the provider's configured [namespace](/docs/providers/vault#namespace).`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `(Required) The name of the AWS auth backend role to read role tags from, with no leading or trailing ` + "`" + `/` + "`" + `s.`,
				},
				resource.Attribute{
					Name:        "backend",
					Description: `(Optional) The path to the AWS auth backend to read role tags from, with no leading or trailing ` + "`" + `/` + "`" + `s. Defaults to "aws".`,
				},
				resource.Attribute{
					Name:        "policies",
					Description: `(Optional) The policies to be associated with the tag. Must be a subset of the policies associated with the role.`,
				},
				resource.Attribute{
					Name:        "max_ttl",
					Description: `(Optional) The maximum TTL of the tokens issued using this role.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Optional) Instance ID for which this tag is intended for. If set, the created tag can only be used by the instance with the given ID.`,
				},
				resource.Attribute{
					Name:        "allow_instance_migration",
					Description: `(Optional) If set, allows migration of the underlying instances where the client resides. Use with caution.`,
				},
				resource.Attribute{
					Name:        "disallow_reauthentication",
					Description: `(Optional) If set, only allows a single token to be granted per instance ID. ## Attributes Reference In addition to the arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "tag_key",
					Description: `The key of the role tag.`,
				},
				resource.Attribute{
					Name:        "tag_value",
					Description: `The value to set the role key.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "tag_key",
					Description: `The key of the role tag.`,
				},
				resource.Attribute{
					Name:        "tag_value",
					Description: `The value to set the role key.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_aws_auth_backend_roletag_blacklist",
			Category:         "Resources",
			ShortDescription: `Configures the periodic tidying operation of the blacklisted role tag entries.`,
			Description:      ``,
			Keywords: []string{
				"aws",
				"auth",
				"backend",
				"roletag",
				"blacklist",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) The namespace to provision the resource in. The value should not contain leading or trailing forward slashes. The ` + "`" + `namespace` + "`" + ` is always relative to the provider's configured [namespace](/docs/providers/vault#namespace).`,
				},
				resource.Attribute{
					Name:        "backend",
					Description: `(Required) The path the AWS auth backend being configured was mounted at.`,
				},
				resource.Attribute{
					Name:        "safety_buffer",
					Description: `(Optional) The amount of extra time that must have passed beyond the roletag expiration, before it is removed from the backend storage. Defaults to 259,200 seconds, or 72 hours.`,
				},
				resource.Attribute{
					Name:        "disable_periodic_tidy",
					Description: `(Optional) If set to true, disables the periodic tidying of the roletag blacklist entries. Defaults to false. ## Attributes Reference No additional attributes are exported by this resource.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_aws_auth_backend_sts_role",
			Category:         "Resources",
			ShortDescription: `Configures an STS role in the Vault AWS Auth backend.`,
			Description:      ``,
			Keywords: []string{
				"aws",
				"auth",
				"backend",
				"sts",
				"role",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) The namespace to provision the resource in. The value should not contain leading or trailing forward slashes. The ` + "`" + `namespace` + "`" + ` is always relative to the provider's configured [namespace](/docs/providers/vault#namespace).`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `(Optional) The AWS account ID to configure the STS role for.`,
				},
				resource.Attribute{
					Name:        "sts_role",
					Description: `(Optional) The STS role to assume when verifying requests made by EC2 instances in the account specified by ` + "`" + `account_id` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "backend",
					Description: `(Optional) The path the AWS auth backend being configured was mounted at. Defaults to ` + "`" + `aws` + "`" + `. ## Attributes Reference No additional attributes are exported by this resource. ## Import AWS auth backend STS roles can be imported using ` + "`" + `auth/` + "`" + `, the ` + "`" + `backend` + "`" + ` path, ` + "`" + `/config/sts/` + "`" + `, and the ` + "`" + `account_id` + "`" + ` e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_aws_auth_backend_sts_role.example auth/aws/config/sts/1234567890 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_aws_secret_backend",
			Category:         "Resources",
			ShortDescription: `Creates an AWS secret backend for Vault.`,
			Description:      ``,
			Keywords: []string{
				"aws",
				"secret",
				"backend",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) The namespace to provision the resource in. The value should not contain leading or trailing forward slashes. The ` + "`" + `namespace` + "`" + ` is always relative to the provider's configured [namespace](/docs/providers/vault#namespace).`,
				},
				resource.Attribute{
					Name:        "access_key",
					Description: `(Optional) The AWS Access Key ID this backend should use to issue new credentials. Vault uses the official AWS SDK to authenticate, and thus can also use standard AWS environment credentials, shared file credentials or IAM role/ECS task credentials.`,
				},
				resource.Attribute{
					Name:        "secret_key",
					Description: `(Optional) The AWS Secret Key this backend should use to issue new credentials. Vault uses the official AWS SDK to authenticate, and thus can also use standard AWS environment credentials, shared file credentials or IAM role/ECS task credentials. ~>`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The AWS region for API calls. Defaults to ` + "`" + `us-east-1` + "`" + `. ~>`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Optional) The unique path this backend should be mounted at. Must not begin or end with a ` + "`" + `/` + "`" + `. Defaults to ` + "`" + `aws` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "disable_remount",
					Description: `(Optional) If set, opts out of mount migration on path updates. See here for more info on [Mount Migration](https://www.vaultproject.io/docs/concepts/mount-migration)`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A human-friendly description for this backend.`,
				},
				resource.Attribute{
					Name:        "default_lease_ttl_seconds",
					Description: `(Optional) The default TTL for credentials issued by this backend.`,
				},
				resource.Attribute{
					Name:        "max_lease_ttl_seconds",
					Description: `(Optional) The maximum TTL that can be requested for credentials issued by this backend.`,
				},
				resource.Attribute{
					Name:        "iam_endpoint",
					Description: `(Optional) Specifies a custom HTTP IAM endpoint to use.`,
				},
				resource.Attribute{
					Name:        "sts_endpoint",
					Description: `(Optional) Specifies a custom HTTP STS endpoint to use.`,
				},
				resource.Attribute{
					Name:        "username_template",
					Description: `(Optional) Template describing how dynamic usernames are generated. The username template is used to generate both IAM usernames (capped at 64 characters) and STS usernames (capped at 32 characters). If no template is provided the field defaults to the template: ` + "`" + `` + "`" + `` + "`" + ` {{ if (eq .Type "STS") }} {{ printf "vault-%s-%s" (unix_time) (random 20) | truncate 32 }} {{ else }} {{ printf "vault-%s-%s-%s" (printf "%s-%s" (.DisplayName) (.PolicyName) | truncate 42) (unix_time) (random 20) | truncate 64 }} {{ end }} ` + "`" + `` + "`" + `` + "`" + ` ## Attributes Reference No additional attributes are exported by this resource. ## Import AWS secret backends can be imported using the ` + "`" + `path` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_aws_secret_backend.aws aws ` + "`" + `` + "`" + `` + "`" + ` ## Tutorials Refer to the [Inject Secrets into Terraform Using the Vault Provider](https://learn.hashicorp.com/tutorials/terraform/secrets-vault) tutorial for a step-by-step usage example.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_aws_secret_backend_role",
			Category:         "Resources",
			ShortDescription: `Creates a role on an AWS Secret Backend for Vault.`,
			Description:      ``,
			Keywords: []string{
				"aws",
				"secret",
				"backend",
				"role",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) The namespace to provision the resource in. The value should not contain leading or trailing forward slashes. The ` + "`" + `namespace` + "`" + ` is always relative to the provider's configured [namespace](/docs/providers/vault#namespace).`,
				},
				resource.Attribute{
					Name:        "backend",
					Description: `(Required) The path the AWS secret backend is mounted at, with no leading or trailing ` + "`" + `/` + "`" + `s.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name to identify this role within the backend. Must be unique within the backend.`,
				},
				resource.Attribute{
					Name:        "credential_type",
					Description: `(Required) Specifies the type of credential to be used when retrieving credentials from the role. Must be one of ` + "`" + `iam_user` + "`" + `, ` + "`" + `assumed_role` + "`" + `, or ` + "`" + `federation_token` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "role_arns",
					Description: `(Optional) Specifies the ARNs of the AWS roles this Vault role is allowed to assume. Required when ` + "`" + `credential_type` + "`" + ` is ` + "`" + `assumed_role` + "`" + ` and prohibited otherwise.`,
				},
				resource.Attribute{
					Name:        "policy_arns",
					Description: `(Optional) Specifies a list of AWS managed policy ARNs. The behavior depends on the credential type. With ` + "`" + `iam_user` + "`" + `, the policies will be attached to IAM users when they are requested. With ` + "`" + `assumed_role` + "`" + ` and ` + "`" + `federation_token` + "`" + `, the policy ARNs will act as a filter on what the credentials can do, similar to ` + "`" + `policy_document` + "`" + `. When ` + "`" + `credential_type` + "`" + ` is ` + "`" + `iam_user` + "`" + ` or ` + "`" + `federation_token` + "`" + `, at least one of ` + "`" + `policy_document` + "`" + ` or ` + "`" + `policy_arns` + "`" + ` must be specified.`,
				},
				resource.Attribute{
					Name:        "policy_document",
					Description: `(Optional) The IAM policy document for the role. The behavior depends on the credential type. With ` + "`" + `iam_user` + "`" + `, the policy document will be attached to the IAM user generated and augment the permissions the IAM user has. With ` + "`" + `assumed_role` + "`" + ` and ` + "`" + `federation_token` + "`" + `, the policy document will act as a filter on what the credentials can do, similar to ` + "`" + `policy_arns` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "default_sts_ttl",
					Description: `(Optional) The default TTL in seconds for STS credentials. When a TTL is not specified when STS credentials are requested, and a default TTL is specified on the role, then this default TTL will be used. Valid only when ` + "`" + `credential_type` + "`" + ` is one of ` + "`" + `assumed_role` + "`" + ` or ` + "`" + `federation_token` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "max_sts_ttl",
					Description: `(Optional) The max allowed TTL in seconds for STS credentials (credentials TTL are capped to ` + "`" + `max_sts_ttl` + "`" + `). Valid only when ` + "`" + `credential_type` + "`" + ` is one of ` + "`" + `assumed_role` + "`" + ` or ` + "`" + `federation_token` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "user_path",
					Description: `(Optional) The path for the user name. Valid only when ` + "`" + `credential_type` + "`" + ` is ` + "`" + `iam_user` + "`" + `. Default is ` + "`" + `/` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "permissions_boundary_arn",
					Description: `(Optional) The ARN of the AWS Permissions Boundary to attach to IAM users created in the role. Valid only when ` + "`" + `credential_type` + "`" + ` is ` + "`" + `iam_user` + "`" + `. If not specified, then no permissions boundary policy will be attached. ## Attributes Reference No additional attributes are exported by this resource. ## Import AWS secret backend roles can be imported using the ` + "`" + `path` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_aws_secret_backend_role.role aws/roles/deploy ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_azure_auth_backend_config",
			Category:         "Resources",
			ShortDescription: `Configures the Azure Auth Backend in Vault.`,
			Description:      ``,
			Keywords: []string{
				"azure",
				"auth",
				"backend",
				"config",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) The namespace to provision the resource in. The value should not contain leading or trailing forward slashes. The ` + "`" + `namespace` + "`" + ` is always relative to the provider's configured [namespace](/docs/providers/vault#namespace).`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Required) The tenant id for the Azure Active Directory organization.`,
				},
				resource.Attribute{
					Name:        "resource",
					Description: `(Required) The configured URL for the application registered in Azure Active Directory.`,
				},
				resource.Attribute{
					Name:        "backend",
					Description: `(Optional) The path the Azure auth backend being configured was mounted at. Defaults to ` + "`" + `azure` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "client_id",
					Description: `(Optional) The client id for credentials to query the Azure APIs. Currently read permissions to query compute resources are required.`,
				},
				resource.Attribute{
					Name:        "client_secret",
					Description: `(Optional) The client secret for credentials to query the Azure APIs.`,
				},
				resource.Attribute{
					Name:        "environment",
					Description: `(Optional) The Azure cloud environment. Valid values: AzurePublicCloud, AzureUSGovernmentCloud, AzureChinaCloud, AzureGermanCloud. Defaults to ` + "`" + `AzurePublicCloud` + "`" + `. ## Attributes Reference No additional attributes are exported by this resource. ## Import Azure auth backends can be imported using ` + "`" + `auth/` + "`" + `, the ` + "`" + `backend` + "`" + ` path, and ` + "`" + `/config` + "`" + ` e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_azure_auth_backend_config.example auth/azure/config ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_azure_auth_backend_role",
			Category:         "Resources",
			ShortDescription: `Manages Azure auth backend roles in Vault.`,
			Description:      ``,
			Keywords: []string{
				"azure",
				"auth",
				"backend",
				"role",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) The namespace to provision the resource in. The value should not contain leading or trailing forward slashes. The ` + "`" + `namespace` + "`" + ` is always relative to the provider's configured [namespace](/docs/providers/vault#namespace).`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `(Required) The name of the role.`,
				},
				resource.Attribute{
					Name:        "bound_service_principal_ids",
					Description: `(Optional) If set, defines a constraint on the service principals that can perform the login operation that they should be possess the ids specified by this field.`,
				},
				resource.Attribute{
					Name:        "bound_group_ids",
					Description: `(Optional) If set, defines a constraint on the groups that can perform the login operation that they should be using the group ID specified by this field.`,
				},
				resource.Attribute{
					Name:        "bound_locations",
					Description: `(Optional) If set, defines a constraint on the virtual machines that can perform the login operation that the location in their identity document must match the one specified by this field.`,
				},
				resource.Attribute{
					Name:        "bound_subscription_ids",
					Description: `(Optional) If set, defines a constraint on the subscriptions that can perform the login operation to ones which matches the value specified by this field.`,
				},
				resource.Attribute{
					Name:        "bound_resource_groups",
					Description: `(Optional) If set, defines a constraint on the virtual machines that can perform the login operation that they be associated with the resource group that matches the value specified by this field.`,
				},
				resource.Attribute{
					Name:        "bound_scale_sets",
					Description: `(Optional) If set, defines a constraint on the virtual machines that can perform the login operation that they must match the scale set specified by this field. ### Common Token Arguments These arguments are common across several Authentication Token resources since Vault 1.2.`,
				},
				resource.Attribute{
					Name:        "token_ttl",
					Description: `(Optional) The incremental lifetime for generated tokens in number of seconds. Its current value will be referenced at renewal time.`,
				},
				resource.Attribute{
					Name:        "token_max_ttl",
					Description: `(Optional) The maximum lifetime for generated tokens in number of seconds. Its current value will be referenced at renewal time.`,
				},
				resource.Attribute{
					Name:        "token_period",
					Description: `(Optional) If set, indicates that the token generated using this role should never expire. The token should be renewed within the duration specified by this value. At each renewal, the token's TTL will be set to the value of this field. Specified in seconds.`,
				},
				resource.Attribute{
					Name:        "token_policies",
					Description: `(Optional) List of policies to encode onto generated tokens. Depending on the auth method, this list may be supplemented by user/group/other values.`,
				},
				resource.Attribute{
					Name:        "token_bound_cidrs",
					Description: `(Optional) List of CIDR blocks; if set, specifies blocks of IP addresses which can authenticate successfully, and ties the resulting token to these blocks as well.`,
				},
				resource.Attribute{
					Name:        "token_explicit_max_ttl",
					Description: `(Optional) If set, will encode an [explicit max TTL](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls) onto the token in number of seconds. This is a hard cap even if ` + "`" + `token_ttl` + "`" + ` and ` + "`" + `token_max_ttl` + "`" + ` would otherwise allow a renewal.`,
				},
				resource.Attribute{
					Name:        "token_no_default_policy",
					Description: `(Optional) If set, the default policy will not be set on generated tokens; otherwise it will be added to the policies set in token_policies.`,
				},
				resource.Attribute{
					Name:        "token_num_uses",
					Description: `(Optional) The [maximum number](https://www.vaultproject.io/api-docs/azure#token_num_uses) of times a generated token may be used (within its lifetime); 0 means unlimited.`,
				},
				resource.Attribute{
					Name:        "token_type",
					Description: `(Optional) The type of token that should be generated. Can be ` + "`" + `service` + "`" + `, ` + "`" + `batch` + "`" + `, or ` + "`" + `default` + "`" + ` to use the mount's tuned default (which unless changed will be ` + "`" + `service` + "`" + ` tokens). For token store roles, there are two additional possibilities: ` + "`" + `default-service` + "`" + ` and ` + "`" + `default-batch` + "`" + ` which specify the type to return unless the client requests a different type at generation time. ## Attributes Reference No additional attributes are exported by this resource. ## Import Azure auth backend roles can be imported using ` + "`" + `auth/` + "`" + `, the ` + "`" + `backend` + "`" + ` path, ` + "`" + `/role/` + "`" + `, and the ` + "`" + `role` + "`" + ` name e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_azure_auth_backend_role.example auth/azure/role/test-role ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_azure_secret_backend",
			Category:         "Resources",
			ShortDescription: `Creates an azure secret backend for Vault.`,
			Description:      ``,
			Keywords: []string{
				"azure",
				"secret",
				"backend",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_azure_secret_backend_role",
			Category:         "Resources",
			ShortDescription: `Creates an azure secret backend role for Vault.`,
			Description:      ``,
			Keywords: []string{
				"azure",
				"secret",
				"backend",
				"role",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) The namespace to provision the resource in. The value should not contain leading or trailing forward slashes. The ` + "`" + `namespace` + "`" + ` is always relative to the provider's configured [namespace](/docs/providers/vault#namespace).`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `(Required) Name of the Azure role`,
				},
				resource.Attribute{
					Name:        "backend",
					Description: `Path to the mounted Azure auth backend`,
				},
				resource.Attribute{
					Name:        "azure_groups",
					Description: `List of Azure groups to be assigned to the generated service principal.`,
				},
				resource.Attribute{
					Name:        "azure_roles",
					Description: `List of Azure roles to be assigned to the generated service principal.`,
				},
				resource.Attribute{
					Name:        "application_object_id",
					Description: `Application Object ID for an existing service principal that will be used instead of creating dynamic service principals. If present, ` + "`" + `azure_roles` + "`" + ` will be ignored.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_cert_auth_backend_role",
			Category:         "Resources",
			ShortDescription: `Managing roles in an Cert auth backend in Vault`,
			Description:      ``,
			Keywords: []string{
				"cert",
				"auth",
				"backend",
				"role",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) The namespace to provision the resource in. The value should not contain leading or trailing forward slashes. The ` + "`" + `namespace` + "`" + ` is always relative to the provider's configured [namespace](/docs/providers/vault#namespace).`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the role`,
				},
				resource.Attribute{
					Name:        "certificate",
					Description: `(Required) CA certificate used to validate client certificates`,
				},
				resource.Attribute{
					Name:        "allowed_names",
					Description: `(Optional) Allowed subject names for authenticated client certificates`,
				},
				resource.Attribute{
					Name:        "allowed_common_names",
					Description: `(Optional) Allowed the common names for authenticated client certificates`,
				},
				resource.Attribute{
					Name:        "allowed_dns_sans",
					Description: `(Optional) Allowed alternative dns names for authenticated client certificates`,
				},
				resource.Attribute{
					Name:        "allowed_email_sans",
					Description: `(Optional) Allowed emails for authenticated client certificates`,
				},
				resource.Attribute{
					Name:        "allowed_uri_sans",
					Description: `(Optional) Allowed URIs for authenticated client certificates`,
				},
				resource.Attribute{
					Name:        "allowed_organizational_units",
					Description: `(Optional) Allowed organization units for authenticated client certificates.`,
				},
				resource.Attribute{
					Name:        "required_extensions",
					Description: `(Optional) TLS extensions required on client certificates`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) The name to display on tokens issued under this role.`,
				},
				resource.Attribute{
					Name:        "backend",
					Description: `(Optional) Path to the mounted Cert auth backend ### Common Token Arguments These arguments are common across several Authentication Token resources since Vault 1.2.`,
				},
				resource.Attribute{
					Name:        "token_ttl",
					Description: `(Optional) The incremental lifetime for generated tokens in number of seconds. Its current value will be referenced at renewal time.`,
				},
				resource.Attribute{
					Name:        "token_max_ttl",
					Description: `(Optional) The maximum lifetime for generated tokens in number of seconds. Its current value will be referenced at renewal time.`,
				},
				resource.Attribute{
					Name:        "token_period",
					Description: `(Optional) If set, indicates that the token generated using this role should never expire. The token should be renewed within the duration specified by this value. At each renewal, the token's TTL will be set to the value of this field. Specified in seconds.`,
				},
				resource.Attribute{
					Name:        "token_policies",
					Description: `(Optional) List of policies to encode onto generated tokens. Depending on the auth method, this list may be supplemented by user/group/other values.`,
				},
				resource.Attribute{
					Name:        "token_bound_cidrs",
					Description: `(Optional) List of CIDR blocks; if set, specifies blocks of IP addresses which can authenticate successfully, and ties the resulting token to these blocks as well.`,
				},
				resource.Attribute{
					Name:        "token_explicit_max_ttl",
					Description: `(Optional) If set, will encode an [explicit max TTL](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls) onto the token in number of seconds. This is a hard cap even if ` + "`" + `token_ttl` + "`" + ` and ` + "`" + `token_max_ttl` + "`" + ` would otherwise allow a renewal.`,
				},
				resource.Attribute{
					Name:        "token_no_default_policy",
					Description: `(Optional) If set, the default policy will not be set on generated tokens; otherwise it will be added to the policies set in token_policies.`,
				},
				resource.Attribute{
					Name:        "token_num_uses",
					Description: `(Optional) The [maximum number](https://www.vaultproject.io/api-docs/auth/cert#token_num_uses) of times a generated token may be used (within its lifetime); 0 means unlimited.`,
				},
				resource.Attribute{
					Name:        "token_type",
					Description: `(Optional) The type of token that should be generated. Can be ` + "`" + `service` + "`" + `, ` + "`" + `batch` + "`" + `, or ` + "`" + `default` + "`" + ` to use the mount's tuned default (which unless changed will be ` + "`" + `service` + "`" + ` tokens). For token store roles, there are two additional possibilities: ` + "`" + `default-service` + "`" + ` and ` + "`" + `default-batch` + "`" + ` which specify the type to return unless the client requests a different type at generation time. For more details on the usage of each argument consult the [Vault Cert API documentation](https://www.vaultproject.io/api-docs/auth/cert). ## Attribute Reference No additional attributes are exposed by this resource.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_consul_secret_backend",
			Category:         "Resources",
			ShortDescription: `Creates a Consul secret backend for Vault.`,
			Description:      ``,
			Keywords: []string{
				"consul",
				"secret",
				"backend",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) The namespace to provision the resource in. The value should not contain leading or trailing forward slashes. The ` + "`" + `namespace` + "`" + ` is always relative to the provider's configured [namespace](/docs/providers/vault#namespace).`,
				},
				resource.Attribute{
					Name:        "token",
					Description: `(Optional) The Consul management token this backend should use to issue new tokens. This field is required when ` + "`" + `bootstrap` + "`" + ` is false. ~>`,
				},
				resource.Attribute{
					Name:        "bootstrap",
					Description: `(Optional) Denotes that the resource is used to bootstrap the Consul ACL system. ~>`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Optional) The unique location this backend should be mounted at. Must not begin or end with a ` + "`" + `/` + "`" + `. Defaults to ` + "`" + `consul` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "disable_remount",
					Description: `(Optional) If set, opts out of mount migration on path updates. See here for more info on [Mount Migration](https://www.vaultproject.io/docs/concepts/mount-migration)`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A human-friendly description for this backend.`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `(Required) Specifies the address of the Consul instance, provided as "host:port" like "127.0.0.1:8500".`,
				},
				resource.Attribute{
					Name:        "scheme",
					Description: `(Optional) Specifies the URL scheme to use. Defaults to ` + "`" + `http` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ca_cert",
					Description: `(Optional) CA certificate to use when verifying Consul server certificate, must be x509 PEM encoded.`,
				},
				resource.Attribute{
					Name:        "client_cert",
					Description: `(Optional) Client certificate used for Consul's TLS communication, must be x509 PEM encoded and if this is set you need to also set client_key.`,
				},
				resource.Attribute{
					Name:        "client_key",
					Description: `(Optional) Client key used for Consul's TLS communication, must be x509 PEM encoded and if this is set you need to also set client_cert.`,
				},
				resource.Attribute{
					Name:        "default_lease_ttl_seconds",
					Description: `(Optional) The default TTL for credentials issued by this backend.`,
				},
				resource.Attribute{
					Name:        "max_lease_ttl_seconds",
					Description: `(Optional) The maximum TTL that can be requested for credentials issued by this backend.`,
				},
				resource.Attribute{
					Name:        "local",
					Description: `(Optional) Specifies if the secret backend is local only. ## Attributes Reference No additional attributes are exported by this resource. ## Import Consul secret backends can be imported using the ` + "`" + `path` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_consul_secret_backend.example consul ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_consul_secret_backend_role",
			Category:         "Resources",
			ShortDescription: `Manages a Consul secrets role for a Consul secrets engine in Vault.`,
			Description:      ``,
			Keywords: []string{
				"consul",
				"secret",
				"backend",
				"role",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) The namespace to provision the resource in. The value should not contain leading or trailing forward slashes. The ` + "`" + `namespace` + "`" + ` is always relative to the provider's configured [namespace](/docs/providers/vault#namespace).`,
				},
				resource.Attribute{
					Name:        "backend",
					Description: `(Optional) The unique name of an existing Consul secrets backend mount. Must not begin or end with a ` + "`" + `/` + "`" + `. One of ` + "`" + `path` + "`" + ` or ` + "`" + `backend` + "`" + ` is required.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the Consul secrets engine role to create.`,
				},
				resource.Attribute{
					Name:        "policies",
					Description: `(Optional) The list of Consul ACL policies to associate with these roles.`,
				},
				resource.Attribute{
					Name:        "consul_policies",
					Description: `(Optional)<sup><a href="#note-about-required-arguments">SEE NOTE</a></sup> The list of Consul ACL policies to associate with these roles.`,
				},
				resource.Attribute{
					Name:        "consul_roles",
					Description: `(Optional)<sup><a href="#note-about-required-arguments">SEE NOTE</a></sup> Set of Consul roles to attach to the token. Applicable for Vault 1.10+ with Consul 1.5+.`,
				},
				resource.Attribute{
					Name:        "service_identities",
					Description: `(Optional)<sup><a href="#note-about-required-arguments">SEE NOTE</a></sup> Set of Consul service identities to attach to the token. Applicable for Vault 1.11+ with Consul 1.5+.`,
				},
				resource.Attribute{
					Name:        "node_identities",
					Description: `(Optional)<sup><a href="#note-about-required-arguments">SEE NOTE</a></sup> Set of Consul node identities to attach to the token. Applicable for Vault 1.11+ with Consul 1.8+.`,
				},
				resource.Attribute{
					Name:        "consul_namespace",
					Description: `(Optional) The Consul namespace that the token will be created in. Applicable for Vault 1.10+ and Consul 1.7+".`,
				},
				resource.Attribute{
					Name:        "partition",
					Description: `(Optional) The admin partition that the token will be created in. Applicable for Vault 1.10+ and Consul 1.11+".`,
				},
				resource.Attribute{
					Name:        "max_ttl",
					Description: `(Optional) Maximum TTL for leases associated with this role, in seconds.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `(Optional) Specifies the TTL for this role.`,
				},
				resource.Attribute{
					Name:        "token_type",
					Description: `(Optional) Specifies the type of token to create when using this role. Valid values are "client" or "management".`,
				},
				resource.Attribute{
					Name:        "local",
					Description: `(Optional) Indicates that the token should not be replicated globally and instead be local to the current datacenter. ## Attributes Reference No additional attributes are exported by this resource. ## Import Consul secret backend roles can be imported using the ` + "`" + `backend` + "`" + `, ` + "`" + `/roles/` + "`" + `, and the ` + "`" + `name` + "`" + ` e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_consul_secret_backend_role.example consul/roles/my-role ` + "`" + `` + "`" + `` + "`" + ` ## Note About Required Arguments`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_database_secret_backend_connection",
			Category:         "Resources",
			ShortDescription: `Configures a database secret backend connection for Vault.`,
			Description:      ``,
			Keywords: []string{
				"database",
				"secret",
				"backend",
				"connection",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) The namespace to provision the resource in. The value should not contain leading or trailing forward slashes. The ` + "`" + `namespace` + "`" + ` is always relative to the provider's configured [namespace](../index.html#namespace).`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A unique name to give the database connection.`,
				},
				resource.Attribute{
					Name:        "backend",
					Description: `(Required) The unique name of the Vault mount to configure.`,
				},
				resource.Attribute{
					Name:        "plugin_name",
					Description: `(Optional) Specifies the name of the plugin to use.`,
				},
				resource.Attribute{
					Name:        "verify_connection",
					Description: `(Optional) Whether the connection should be verified on initial configuration or not.`,
				},
				resource.Attribute{
					Name:        "allowed_roles",
					Description: `(Optional) A list of roles that are allowed to use this connection.`,
				},
				resource.Attribute{
					Name:        "root_rotation_statements",
					Description: `(Optional) A list of database statements to be executed to rotate the root user's credentials.`,
				},
				resource.Attribute{
					Name:        "data",
					Description: `(Optional) A map of sensitive data to pass to the endpoint. Useful for templated connection strings.`,
				},
				resource.Attribute{
					Name:        "cassandra",
					Description: `(Optional) A nested block containing configuration options for Cassandra connections.`,
				},
				resource.Attribute{
					Name:        "couchbase",
					Description: `(Optional) A nested block containing configuration options for Couchbase connections.`,
				},
				resource.Attribute{
					Name:        "mongodb",
					Description: `(Optional) A nested block containing configuration options for MongoDB connections.`,
				},
				resource.Attribute{
					Name:        "mongodbatlas",
					Description: `(Optional) A nested block containing configuration options for MongoDB Atlas connections.`,
				},
				resource.Attribute{
					Name:        "hana",
					Description: `(Optional) A nested block containing configuration options for SAP HanaDB connections.`,
				},
				resource.Attribute{
					Name:        "mssql",
					Description: `(Optional) A nested block containing configuration options for MSSQL connections.`,
				},
				resource.Attribute{
					Name:        "mysql",
					Description: `(Optional) A nested block containing configuration options for MySQL connections.`,
				},
				resource.Attribute{
					Name:        "mysql_rds",
					Description: `(Optional) A nested block containing configuration options for RDS MySQL connections.`,
				},
				resource.Attribute{
					Name:        "mysql_aurora",
					Description: `(Optional) A nested block containing configuration options for Aurora MySQL connections.`,
				},
				resource.Attribute{
					Name:        "mysql_legacy",
					Description: `(Optional) A nested block containing configuration options for legacy MySQL connections.`,
				},
				resource.Attribute{
					Name:        "postgresql",
					Description: `(Optional) A nested block containing configuration options for PostgreSQL connections.`,
				},
				resource.Attribute{
					Name:        "oracle",
					Description: `(Optional) A nested block containing configuration options for Oracle connections.`,
				},
				resource.Attribute{
					Name:        "elasticsearch",
					Description: `(Optional) A nested block containing configuration options for Elasticsearch connections.`,
				},
				resource.Attribute{
					Name:        "snowflake",
					Description: `(Optional) A nested block containing configuration options for Snowflake connections.`,
				},
				resource.Attribute{
					Name:        "influxdb",
					Description: `(Optional) A nested block containing configuration options for InfluxDB connections.`,
				},
				resource.Attribute{
					Name:        "redis",
					Description: `(Optional) A nested block containing configuration options for Redis connections.`,
				},
				resource.Attribute{
					Name:        "redis_elasticache",
					Description: `(Optional) A nested block containing configuration options for Redis ElastiCache connections. Exactly one of the nested blocks of configuration options must be supplied. ### Cassandra Configuration Options`,
				},
				resource.Attribute{
					Name:        "hosts",
					Description: `(Required) The hosts to connect to.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Required) The username to authenticate with.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Required) The password to authenticate with.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional) The default port to connect to if no port is specified as part of the host.`,
				},
				resource.Attribute{
					Name:        "tls",
					Description: `(Optional) Whether to use TLS when connecting to Cassandra.`,
				},
				resource.Attribute{
					Name:        "insecure_tls",
					Description: `(Optional) Whether to skip verification of the server certificate when using TLS.`,
				},
				resource.Attribute{
					Name:        "pem_bundle",
					Description: `(Optional) Concatenated PEM blocks configuring the certificate chain.`,
				},
				resource.Attribute{
					Name:        "pem_json",
					Description: `(Optional) A JSON structure configuring the certificate chain.`,
				},
				resource.Attribute{
					Name:        "protocol_version",
					Description: `(Optional) The CQL protocol version to use.`,
				},
				resource.Attribute{
					Name:        "connect_timeout",
					Description: `(Optional) The number of seconds to use as a connection timeout. ### Couchbase Configuration Options`,
				},
				resource.Attribute{
					Name:        "hosts",
					Description: `(Required) A set of Couchbase URIs to connect to. Must use ` + "`" + `couchbases://` + "`" + ` scheme if ` + "`" + `tls` + "`" + ` is ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Required) Specifies the username for Vault to use.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Required) Specifies the password corresponding to the given username.`,
				},
				resource.Attribute{
					Name:        "tls",
					Description: `(Optional) Whether to use TLS when connecting to Couchbase.`,
				},
				resource.Attribute{
					Name:        "insecure_tls",
					Description: `(Optional) Whether to skip verification of the server certificate when using TLS.`,
				},
				resource.Attribute{
					Name:        "base64_pem",
					Description: `(Optional) Required if ` + "`" + `tls` + "`" + ` is ` + "`" + `true` + "`" + `. Specifies the certificate authority of the Couchbase server, as a PEM certificate that has been base64 encoded.`,
				},
				resource.Attribute{
					Name:        "bucket_name",
					Description: `(Optional) Required for Couchbase versions prior to 6.5.0. This is only used to verify vault's connection to the server.`,
				},
				resource.Attribute{
					Name:        "username_template",
					Description: `(Optional) Template describing how dynamic usernames are generated. ### InfluxDB Configuration Options`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `(Required) The host to connect to.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Required) The username to authenticate with.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Required) The password to authenticate with.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional) The default port to connect to if no port is specified as part of the host.`,
				},
				resource.Attribute{
					Name:        "tls",
					Description: `(Optional) Whether to use TLS when connecting to Cassandra.`,
				},
				resource.Attribute{
					Name:        "insecure_tls",
					Description: `(Optional) Whether to skip verification of the server certificate when using TLS.`,
				},
				resource.Attribute{
					Name:        "pem_bundle",
					Description: `(Optional) Concatenated PEM blocks configuring the certificate chain.`,
				},
				resource.Attribute{
					Name:        "pem_json",
					Description: `(Optional) A JSON structure configuring the certificate chain.`,
				},
				resource.Attribute{
					Name:        "username_template",
					Description: `(Optional) Template describing how dynamic usernames are generated.`,
				},
				resource.Attribute{
					Name:        "connect_timeout",
					Description: `(Optional) The number of seconds to use as a connection timeout. ### Redis Configuration Options`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `(Required) The host to connect to.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Required) The username to authenticate with.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Required) The password to authenticate with.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Required) The default port to connect to if no port is specified as part of the host.`,
				},
				resource.Attribute{
					Name:        "tls",
					Description: `(Optional) Whether to use TLS when connecting to Redis.`,
				},
				resource.Attribute{
					Name:        "insecure_tls",
					Description: `(Optional) Whether to skip verification of the server certificate when using TLS.`,
				},
				resource.Attribute{
					Name:        "ca_cert",
					Description: `(Optional) The contents of a PEM-encoded CA cert file to use to verify the Redis server's identity. ### Redis ElastiCache Configuration Options`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Required) The url to connect to including the port; e.g. master.my-cluster.xxxxxx.use1.cache.amazonaws.com:6379.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Optional) The AWS access key id to authenticate with. If omitted Vault tries to infer from the credential provider chain instead.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional) The AWS secret access key to authenticate with. If omitted Vault tries to infer from the credential provider chain instead.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region where the ElastiCache cluster is hosted. If omitted Vault tries to infer from the environment instead. ### MongoDB Configuration Options`,
				},
				resource.Attribute{
					Name:        "connection_url",
					Description: `(Required) A URL containing connection information. See the [Vault docs](https://www.vaultproject.io/api-docs/secret/databases/mongodb.html#sample-payload) for an example.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Optional) The root credential username used in the connection URL.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional) The root credential password used in the connection URL.`,
				},
				resource.Attribute{
					Name:        "username_template",
					Description: `(Optional) For Vault v1.7+. The template to use for username generation. See the [Vault docs](https://www.vaultproject.io/docs/concepts/username-templating) ### MongoDB Atlas Configuration Options`,
				},
				resource.Attribute{
					Name:        "public_key",
					Description: `(Required) The Public Programmatic API Key used to authenticate with the MongoDB Atlas API.`,
				},
				resource.Attribute{
					Name:        "private_key",
					Description: `(Required) The Private Programmatic API Key used to connect with MongoDB Atlas API.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The Project ID the Database User should be created within. ### SAP HanaDB Configuration Options`,
				},
				resource.Attribute{
					Name:        "connection_url",
					Description: `(Required) A URL containing connection information. See the [Vault docs](https://www.vaultproject.io/api-docs/secret/databases/hanadb.html#sample-payload) for an example.`,
				},
				resource.Attribute{
					Name:        "max_open_connections",
					Description: `(Optional) The maximum number of open connections to use.`,
				},
				resource.Attribute{
					Name:        "max_idle_connections",
					Description: `(Optional) The maximum number of idle connections to maintain.`,
				},
				resource.Attribute{
					Name:        "max_connection_lifetime",
					Description: `(Optional) The maximum number of seconds to keep a connection alive for.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Optional) The root credential username used in the connection URL.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional) The root credential password used in the connection URL.`,
				},
				resource.Attribute{
					Name:        "disable_escaping",
					Description: `(Optional) Disable special character escaping in username and password. ### MSSQL Configuration Options`,
				},
				resource.Attribute{
					Name:        "connection_url",
					Description: `(Required) A URL containing connection information. See the [Vault docs](https://www.vaultproject.io/api-docs/secret/databases/mssql.html#sample-payload) for an example.`,
				},
				resource.Attribute{
					Name:        "max_open_connections",
					Description: `(Optional) The maximum number of open connections to use.`,
				},
				resource.Attribute{
					Name:        "max_idle_connections",
					Description: `(Optional) The maximum number of idle connections to maintain.`,
				},
				resource.Attribute{
					Name:        "max_connection_lifetime",
					Description: `(Optional) The maximum number of seconds to keep a connection alive for.`,
				},
				resource.Attribute{
					Name:        "username_template",
					Description: `(Optional) For Vault v1.7+. The template to use for username generation. See the [Vault docs](https://www.vaultproject.io/docs/concepts/username-templating)`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Optional) The root credential username used in the connection URL.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional) The root credential password used in the connection URL.`,
				},
				resource.Attribute{
					Name:        "disable_escaping",
					Description: `(Optional) Disable special character escaping in username and password.`,
				},
				resource.Attribute{
					Name:        "contained_db",
					Description: `(Optional bool: false) For Vault v1.9+. Set to true when the target is a Contained Database, e.g. AzureSQL. See the [Vault docs](https://www.vaultproject.io/api/secret/databases/mssql#contained_db) ### MySQL Configuration Options`,
				},
				resource.Attribute{
					Name:        "connection_url",
					Description: `(Required) A URL containing connection information. See the [Vault docs](https://www.vaultproject.io/api-docs/secret/databases/mysql-maria.html#sample-payload) for an example.`,
				},
				resource.Attribute{
					Name:        "max_open_connections",
					Description: `(Optional) The maximum number of open connections to use.`,
				},
				resource.Attribute{
					Name:        "max_idle_connections",
					Description: `(Optional) The maximum number of idle connections to maintain.`,
				},
				resource.Attribute{
					Name:        "max_connection_lifetime",
					Description: `(Optional) The maximum number of seconds to keep a connection alive for.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Optional) The root credential username used in the connection URL.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional) The root credential password used in the connection URL.`,
				},
				resource.Attribute{
					Name:        "tls_certificate_key",
					Description: `(Optional) x509 certificate for connecting to the database. This must be a PEM encoded version of the private key and the certificate combined.`,
				},
				resource.Attribute{
					Name:        "tls_ca",
					Description: `(Optional) x509 CA file for validating the certificate presented by the MySQL server. Must be PEM encoded.`,
				},
				resource.Attribute{
					Name:        "username_template",
					Description: `(Optional) For Vault v1.7+. The template to use for username generation. See the [Vault docs](https://www.vaultproject.io/docs/concepts/username-templating) ### PostgreSQL Configuration Options`,
				},
				resource.Attribute{
					Name:        "connection_url",
					Description: `(Required) A URL containing connection information. See the [Vault docs](https://www.vaultproject.io/api-docs/secret/databases/postgresql.html#sample-payload) for an example.`,
				},
				resource.Attribute{
					Name:        "max_open_connections",
					Description: `(Optional) The maximum number of open connections to use.`,
				},
				resource.Attribute{
					Name:        "max_idle_connections",
					Description: `(Optional) The maximum number of idle connections to maintain.`,
				},
				resource.Attribute{
					Name:        "max_connection_lifetime",
					Description: `(Optional) The maximum number of seconds to keep a connection alive for.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Optional) The root credential username used in the connection URL.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional) The root credential password used in the connection URL.`,
				},
				resource.Attribute{
					Name:        "disable_escaping",
					Description: `(Optional) Disable special character escaping in username and password.`,
				},
				resource.Attribute{
					Name:        "username_template",
					Description: `(Optional) For Vault v1.7+. The template to use for username generation. See the [Vault docs](https://www.vaultproject.io/docs/concepts/username-templating) ### Oracle Configuration Options`,
				},
				resource.Attribute{
					Name:        "connection_url",
					Description: `(Required) A URL containing connection information. See the [Vault docs](https://www.vaultproject.io/api-docs/secret/databases/oracle.html#sample-payload) for an example.`,
				},
				resource.Attribute{
					Name:        "max_open_connections",
					Description: `(Optional) The maximum number of open connections to use.`,
				},
				resource.Attribute{
					Name:        "max_idle_connections",
					Description: `(Optional) The maximum number of idle connections to maintain.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Optional) The root credential username used in the connection URL.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional) The root credential password used in the connection URL.`,
				},
				resource.Attribute{
					Name:        "max_connection_lifetime",
					Description: `(Optional) The maximum number of seconds to keep a connection alive for.`,
				},
				resource.Attribute{
					Name:        "username_template",
					Description: `(Optional) For Vault v1.7+. The template to use for username generation. See the [Vault docs](https://www.vaultproject.io/docs/concepts/username-templating) ### Elasticsearch Configuration Options`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Required) The URL for Elasticsearch's API. https requires certificate by trusted CA if used.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Required) The username to be used in the connection.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Required) The password to be used in the connection.`,
				},
				resource.Attribute{
					Name:        "ca_cert",
					Description: `(Optional) The path to a PEM-encoded CA cert file to use to verify the Elasticsearch server's identity.`,
				},
				resource.Attribute{
					Name:        "ca_path",
					Description: `(Optional) The path to a directory of PEM-encoded CA cert files to use to verify the Elasticsearch server's identity.`,
				},
				resource.Attribute{
					Name:        "client_cert",
					Description: `(Optional) The path to the certificate for the Elasticsearch client to present for communication.`,
				},
				resource.Attribute{
					Name:        "client_key",
					Description: `(Optional) The path to the key for the Elasticsearch client to use for communication.`,
				},
				resource.Attribute{
					Name:        "tls_server_name",
					Description: `(Optional) This, if set, is used to set the SNI host when connecting via TLS.`,
				},
				resource.Attribute{
					Name:        "insecure",
					Description: `(Optional) Whether to disable certificate verification.`,
				},
				resource.Attribute{
					Name:        "username_template",
					Description: `(Optional) For Vault v1.7+. The template to use for username generation. See [Vault docs](https://www.vaultproject.io/docs/concepts/username-templating) for more details. ### Snowflake Configuration Options`,
				},
				resource.Attribute{
					Name:        "connection_url",
					Description: `(Required) A URL containing connection information. See the [Vault docs](https://www.vaultproject.io/api-docs/secret/databases/snowflake#sample-payload) for an example.`,
				},
				resource.Attribute{
					Name:        "max_open_connections",
					Description: `(Optional) The maximum number of open connections to use.`,
				},
				resource.Attribute{
					Name:        "max_idle_connections",
					Description: `(Optional) The maximum number of idle connections to maintain.`,
				},
				resource.Attribute{
					Name:        "max_connection_lifetime",
					Description: `(Optional) The maximum number of seconds to keep a connection alive for.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Optional) The username to be used in the connection (the account admin level).`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional) The password to be used in the connection.`,
				},
				resource.Attribute{
					Name:        "username_template",
					Description: `(Optional) - [Template](https://www.vaultproject.io/docs/concepts/username-templating) describing how dynamic usernames are generated. ### Redshift Configuration Options`,
				},
				resource.Attribute{
					Name:        "connection_url",
					Description: `(Required) Specifies the Redshift DSN. See the [Vault docs](https://www.vaultproject.io/api-docs/secret/databases/redshift#sample-payload) for an example.`,
				},
				resource.Attribute{
					Name:        "max_open_connections",
					Description: `(Optional) The maximum number of open connections to the database.`,
				},
				resource.Attribute{
					Name:        "max_idle_connections",
					Description: `(Optional) The maximum number of idle connections to the database.`,
				},
				resource.Attribute{
					Name:        "max_connection_lifetime",
					Description: `(Optional) The maximum amount of time a connection may be reused.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Optional) The root credential username used in the connection URL.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional) The root credential password used in the connection URL.`,
				},
				resource.Attribute{
					Name:        "disable_escaping",
					Description: `(Optional) Disable special character escaping in username and password.`,
				},
				resource.Attribute{
					Name:        "username_template",
					Description: `(Optional) - [Template](https://www.vaultproject.io/docs/concepts/username-templating) describing how dynamic usernames are generated. ## Attributes Reference No additional attributes are exported by this resource. ## Import Database secret backend connections can be imported using the ` + "`" + `backend` + "`" + `, ` + "`" + `/config/` + "`" + `, and the ` + "`" + `name` + "`" + ` e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_database_secret_backend_connection.example postgres/config/postgres ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_database_secret_backend_role",
			Category:         "Resources",
			ShortDescription: `Configures a database secret backend role for Vault.`,
			Description:      ``,
			Keywords: []string{
				"database",
				"secret",
				"backend",
				"role",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) The namespace to provision the resource in. The value should not contain leading or trailing forward slashes. The ` + "`" + `namespace` + "`" + ` is always relative to the provider's configured [namespace](../index.html#namespace).`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A unique name to give the role.`,
				},
				resource.Attribute{
					Name:        "backend",
					Description: `(Required) The unique name of the Vault mount to configure.`,
				},
				resource.Attribute{
					Name:        "db_name",
					Description: `(Required) The unique name of the database connection to use for the role.`,
				},
				resource.Attribute{
					Name:        "creation_statements",
					Description: `(Required) The database statements to execute when creating a user.`,
				},
				resource.Attribute{
					Name:        "revocation_statements",
					Description: `(Optional) The database statements to execute when revoking a user.`,
				},
				resource.Attribute{
					Name:        "rollback_statements",
					Description: `(Optional) The database statements to execute when rolling back creation due to an error.`,
				},
				resource.Attribute{
					Name:        "renew_statements",
					Description: `(Optional) The database statements to execute when renewing a user.`,
				},
				resource.Attribute{
					Name:        "default_ttl",
					Description: `(Optional) The default number of seconds for leases for this role.`,
				},
				resource.Attribute{
					Name:        "max_ttl",
					Description: `(Optional) The maximum number of seconds for leases for this role. ## Attributes Reference No additional attributes are exported by this resource. ## Import Database secret backend roles can be imported using the ` + "`" + `backend` + "`" + `, ` + "`" + `/roles/` + "`" + `, and the ` + "`" + `name` + "`" + ` e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_database_secret_backend_role.example postgres/roles/my-role ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_database_secret_backend_static_role",
			Category:         "Resources",
			ShortDescription: `Configures a database secret backend static role for Vault.`,
			Description:      ``,
			Keywords: []string{
				"database",
				"secret",
				"backend",
				"static",
				"role",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) The namespace to provision the resource in. The value should not contain leading or trailing forward slashes. The ` + "`" + `namespace` + "`" + ` is always relative to the provider's configured [namespace](../index.html#namespace).`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A unique name to give the static role.`,
				},
				resource.Attribute{
					Name:        "backend",
					Description: `(Required) The unique name of the Vault mount to configure.`,
				},
				resource.Attribute{
					Name:        "db_name",
					Description: `(Required) The unique name of the database connection to use for the static role.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Required) The database username that this static role corresponds to.`,
				},
				resource.Attribute{
					Name:        "rotation_period",
					Description: `(Required) The amount of time Vault should wait before rotating the password, in seconds.`,
				},
				resource.Attribute{
					Name:        "rotation_statements",
					Description: `(Optional) Database statements to execute to rotate the password for the configured database user. ## Attributes Reference No additional attributes are exported by this resource. ## Import Database secret backend static roles can be imported using the ` + "`" + `backend` + "`" + `, ` + "`" + `/static-roles/` + "`" + `, and the ` + "`" + `name` + "`" + ` e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_database_secret_backend_static_role.example postgres/static-roles/my-role ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_database_secrets_mount",
			Category:         "Resources",
			ShortDescription: `Configures any number of database secrets engines under a single mount resource`,
			Description:      ``,
			Keywords: []string{
				"database",
				"secrets",
				"mount",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "path",
					Description: `(Required) Where the secret backend will be mounted`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Human-friendly description of the mount`,
				},
				resource.Attribute{
					Name:        "default_lease_ttl_seconds",
					Description: `(Optional) Default lease duration for tokens and secrets in seconds`,
				},
				resource.Attribute{
					Name:        "max_lease_ttl_seconds",
					Description: `(Optional) Maximum possible lease duration for tokens and secrets in seconds`,
				},
				resource.Attribute{
					Name:        "audit_non_hmac_response_keys",
					Description: `(Optional) Specifies the list of keys that will not be HMAC'd by audit devices in the response data object.`,
				},
				resource.Attribute{
					Name:        "audit_non_hmac_request_keys",
					Description: `(Optional) Specifies the list of keys that will not be HMAC'd by audit devices in the request data object.`,
				},
				resource.Attribute{
					Name:        "local",
					Description: `(Optional) Boolean flag that can be explicitly set to true to enforce local mount in HA environment`,
				},
				resource.Attribute{
					Name:        "options",
					Description: `(Optional) Specifies mount type specific options that are passed to the backend`,
				},
				resource.Attribute{
					Name:        "seal_wrap",
					Description: `(Optional) Boolean flag that can be explicitly set to true to enable seal wrapping for the mount, causing values stored by the mount to be wrapped by the seal's encryption capability`,
				},
				resource.Attribute{
					Name:        "external_entropy_access",
					Description: `(Optional) Boolean flag that can be explicitly set to true to enable the secrets engine to access Vault's external entropy source`,
				},
				resource.Attribute{
					Name:        "allowed_managed_keys",
					Description: `(Optional) Set of managed key registry entry names that the mount in question is allowed to access The following arguments are common to all database engines:`,
				},
				resource.Attribute{
					Name:        "plugin_name",
					Description: `(Optional) Specifies the name of the plugin to use.`,
				},
				resource.Attribute{
					Name:        "verify_connection",
					Description: `(Optional) Whether the connection should be verified on initial configuration or not.`,
				},
				resource.Attribute{
					Name:        "allowed_roles",
					Description: `(Optional) A list of roles that are allowed to use this connection.`,
				},
				resource.Attribute{
					Name:        "root_rotation_statements",
					Description: `(Optional) A list of database statements to be executed to rotate the root user's credentials.`,
				},
				resource.Attribute{
					Name:        "data",
					Description: `(Optional) A map of sensitive data to pass to the endpoint. Useful for templated connection strings. Supported list of database secrets engines that can be configured:`,
				},
				resource.Attribute{
					Name:        "cassandra",
					Description: `(Optional) A nested block containing configuration options for Cassandra connections.`,
				},
				resource.Attribute{
					Name:        "couchbase",
					Description: `(Optional) A nested block containing configuration options for Couchbase connections.`,
				},
				resource.Attribute{
					Name:        "elasticsearch",
					Description: `(Optional) A nested block containing configuration options for Elasticsearch connections.`,
				},
				resource.Attribute{
					Name:        "hana",
					Description: `(Optional) A nested block containing configuration options for SAP HanaDB connections.`,
				},
				resource.Attribute{
					Name:        "mongodb",
					Description: `(Optional) A nested block containing configuration options for MongoDB connections.`,
				},
				resource.Attribute{
					Name:        "mongodbatlas",
					Description: `(Optional) A nested block containing configuration options for MongoDB Atlas connections.`,
				},
				resource.Attribute{
					Name:        "mssql",
					Description: `(Optional) A nested block containing configuration options for MSSQL connections.`,
				},
				resource.Attribute{
					Name:        "mysql",
					Description: `(Optional) A nested block containing configuration options for MySQL connections.`,
				},
				resource.Attribute{
					Name:        "mysql_rds",
					Description: `(Optional) A nested block containing configuration options for RDS MySQL connections.`,
				},
				resource.Attribute{
					Name:        "mysql_aurora",
					Description: `(Optional) A nested block containing configuration options for Aurora MySQL connections.`,
				},
				resource.Attribute{
					Name:        "mysql_legacy",
					Description: `(Optional) A nested block containing configuration options for legacy MySQL connections.`,
				},
				resource.Attribute{
					Name:        "oracle",
					Description: `(Optional) A nested block containing configuration options for Oracle connections.`,
				},
				resource.Attribute{
					Name:        "postgresql",
					Description: `(Optional) A nested block containing configuration options for PostgreSQL connections.`,
				},
				resource.Attribute{
					Name:        "redshift",
					Description: `(Optional) A nested block containing configuration options for AWS Redshift connections.`,
				},
				resource.Attribute{
					Name:        "snowflake",
					Description: `(Optional) A nested block containing configuration options for Snowflake connections.`,
				},
				resource.Attribute{
					Name:        "influxdb",
					Description: `(Optional) A nested block containing configuration options for InfluxDB connections.`,
				},
				resource.Attribute{
					Name:        "redis",
					Description: `(Optional) A nested block containing configuration options for Redis connections.`,
				},
				resource.Attribute{
					Name:        "redis_elasticache",
					Description: `(Optional) A nested block containing configuration options for Redis ElastiCache connections.`,
				},
				resource.Attribute{
					Name:        "hosts",
					Description: `(Required) The hosts to connect to.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Required) The username to authenticate with.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Required) The password to authenticate with.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional) The default port to connect to if no port is specified as part of the host.`,
				},
				resource.Attribute{
					Name:        "tls",
					Description: `(Optional) Whether to use TLS when connecting to Cassandra.`,
				},
				resource.Attribute{
					Name:        "insecure_tls",
					Description: `(Optional) Whether to skip verification of the server certificate when using TLS.`,
				},
				resource.Attribute{
					Name:        "pem_bundle",
					Description: `(Optional) Concatenated PEM blocks configuring the certificate chain.`,
				},
				resource.Attribute{
					Name:        "pem_json",
					Description: `(Optional) A JSON structure configuring the certificate chain.`,
				},
				resource.Attribute{
					Name:        "protocol_version",
					Description: `(Optional) The CQL protocol version to use.`,
				},
				resource.Attribute{
					Name:        "connect_timeout",
					Description: `(Optional) The number of seconds to use as a connection timeout. ### Couchbase Configuration Options`,
				},
				resource.Attribute{
					Name:        "hosts",
					Description: `(Required) A set of Couchbase URIs to connect to. Must use ` + "`" + `couchbases://` + "`" + ` scheme if ` + "`" + `tls` + "`" + ` is ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Required) Specifies the username for Vault to use.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Required) Specifies the password corresponding to the given username.`,
				},
				resource.Attribute{
					Name:        "tls",
					Description: `(Optional) Whether to use TLS when connecting to Couchbase.`,
				},
				resource.Attribute{
					Name:        "insecure_tls",
					Description: `(Optional) Whether to skip verification of the server certificate when using TLS.`,
				},
				resource.Attribute{
					Name:        "base64_pem",
					Description: `(Optional) Required if ` + "`" + `tls` + "`" + ` is ` + "`" + `true` + "`" + `. Specifies the certificate authority of the Couchbase server, as a PEM certificate that has been base64 encoded.`,
				},
				resource.Attribute{
					Name:        "bucket_name",
					Description: `(Optional) Required for Couchbase versions prior to 6.5.0. This is only used to verify vault's connection to the server.`,
				},
				resource.Attribute{
					Name:        "username_template",
					Description: `(Optional) Template describing how dynamic usernames are generated. ### Elasticsearch Configuration Options`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Required) The URL for Elasticsearch's API. https requires certificate by trusted CA if used.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Required) The username to be used in the connection.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Required) The password to be used in the connection.`,
				},
				resource.Attribute{
					Name:        "ca_cert",
					Description: `(Optional) The path to a PEM-encoded CA cert file to use to verify the Elasticsearch server's identity.`,
				},
				resource.Attribute{
					Name:        "ca_path",
					Description: `(Optional) The path to a directory of PEM-encoded CA cert files to use to verify the Elasticsearch server's identity.`,
				},
				resource.Attribute{
					Name:        "client_cert",
					Description: `(Optional) The path to the certificate for the Elasticsearch client to present for communication.`,
				},
				resource.Attribute{
					Name:        "client_key",
					Description: `(Optional) The path to the key for the Elasticsearch client to use for communication.`,
				},
				resource.Attribute{
					Name:        "tls_server_name",
					Description: `(Optional) This, if set, is used to set the SNI host when connecting via TLS.`,
				},
				resource.Attribute{
					Name:        "insecure",
					Description: `(Optional) Whether to disable certificate verification.`,
				},
				resource.Attribute{
					Name:        "username_template",
					Description: `(Optional) For Vault v1.7+. The template to use for username generation. See [Vault docs](https://www.vaultproject.io/docs/concepts/username-templating) for more details. ### InfluxDB Configuration Options`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `(Required) The host to connect to.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Required) The username to authenticate with.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Required) The password to authenticate with.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional) The default port to connect to if no port is specified as part of the host.`,
				},
				resource.Attribute{
					Name:        "tls",
					Description: `(Optional) Whether to use TLS when connecting to Cassandra.`,
				},
				resource.Attribute{
					Name:        "insecure_tls",
					Description: `(Optional) Whether to skip verification of the server certificate when using TLS.`,
				},
				resource.Attribute{
					Name:        "pem_bundle",
					Description: `(Optional) Concatenated PEM blocks configuring the certificate chain.`,
				},
				resource.Attribute{
					Name:        "pem_json",
					Description: `(Optional) A JSON structure configuring the certificate chain.`,
				},
				resource.Attribute{
					Name:        "username_template",
					Description: `(Optional) Template describing how dynamic usernames are generated.`,
				},
				resource.Attribute{
					Name:        "connect_timeout",
					Description: `(Optional) The number of seconds to use as a connection timeout. ### MongoDB Configuration Options`,
				},
				resource.Attribute{
					Name:        "connection_url",
					Description: `(Required) A URL containing connection information. See [Vault docs](https://www.vaultproject.io/api-docs/secret/databases/mongodb.html#sample-payload)`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Optional) The root credential username used in the connection URL.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional) The root credential password used in the connection URL.`,
				},
				resource.Attribute{
					Name:        "username_template",
					Description: `(Optional) For Vault v1.7+. The template to use for username generation. See [Vault docs](https://www.vaultproject.io/docs/concepts/username-templating) ### MongoDB Atlas Configuration Options`,
				},
				resource.Attribute{
					Name:        "public_key",
					Description: `(Required) The Public Programmatic API Key used to authenticate with the MongoDB Atlas API.`,
				},
				resource.Attribute{
					Name:        "private_key",
					Description: `(Required) The Private Programmatic API Key used to connect with MongoDB Atlas API.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The Project ID the Database User should be created within. ### SAP HanaDB Configuration Options`,
				},
				resource.Attribute{
					Name:        "connection_url",
					Description: `(Required) A URL containing connection information. See [Vault docs](https://www.vaultproject.io/api-docs/secret/databases/hanadb.html#sample-payload)`,
				},
				resource.Attribute{
					Name:        "max_open_connections",
					Description: `(Optional) The maximum number of open connections to use.`,
				},
				resource.Attribute{
					Name:        "max_idle_connections",
					Description: `(Optional) The maximum number of idle connections to maintain.`,
				},
				resource.Attribute{
					Name:        "max_connection_lifetime",
					Description: `(Optional) The maximum number of seconds to keep a connection alive for.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Optional) The root credential username used in the connection URL.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional) The root credential password used in the connection URL.`,
				},
				resource.Attribute{
					Name:        "disable_escaping",
					Description: `(Optional) Disable special character escaping in username and password. ### MSSQL Configuration Options`,
				},
				resource.Attribute{
					Name:        "connection_url",
					Description: `(Required) A URL containing connection information. See [Vault docs](https://www.vaultproject.io/api-docs/secret/databases/mssql.html#sample-payload)`,
				},
				resource.Attribute{
					Name:        "max_open_connections",
					Description: `(Optional) The maximum number of open connections to use.`,
				},
				resource.Attribute{
					Name:        "max_idle_connections",
					Description: `(Optional) The maximum number of idle connections to maintain.`,
				},
				resource.Attribute{
					Name:        "max_connection_lifetime",
					Description: `(Optional) The maximum number of seconds to keep a connection alive for.`,
				},
				resource.Attribute{
					Name:        "username_template",
					Description: `(Optional) For Vault v1.7+. The template to use for username generation. See [Vault docs](https://www.vaultproject.io/docs/concepts/username-templating)`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Optional) The root credential username used in the connection URL.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional) The root credential password used in the connection URL.`,
				},
				resource.Attribute{
					Name:        "disable_escaping",
					Description: `(Optional) Disable special character escaping in username and password.`,
				},
				resource.Attribute{
					Name:        "contained_db",
					Description: `(Optional bool: false) For Vault v1.9+. Set to true when the target is a Contained Database, e.g. AzureSQL. See [Vault docs](https://www.vaultproject.io/api/secret/databases/mssql#contained_db) ### MySQL Configuration Options`,
				},
				resource.Attribute{
					Name:        "connection_url",
					Description: `(Required) A URL containing connection information. See [Vault docs](https://www.vaultproject.io/api-docs/secret/databases/mysql-maria.html#sample-payload)`,
				},
				resource.Attribute{
					Name:        "max_open_connections",
					Description: `(Optional) The maximum number of open connections to use.`,
				},
				resource.Attribute{
					Name:        "max_idle_connections",
					Description: `(Optional) The maximum number of idle connections to maintain.`,
				},
				resource.Attribute{
					Name:        "max_connection_lifetime",
					Description: `(Optional) The maximum number of seconds to keep a connection alive for.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Optional) The root credential username used in the connection URL.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional) The root credential password used in the connection URL.`,
				},
				resource.Attribute{
					Name:        "tls_certificate_key",
					Description: `(Optional) x509 certificate for connecting to the database. This must be a PEM encoded version of the private key and the certificate combined.`,
				},
				resource.Attribute{
					Name:        "tls_ca",
					Description: `(Optional) x509 CA file for validating the certificate presented by the MySQL server. Must be PEM encoded.`,
				},
				resource.Attribute{
					Name:        "username_template",
					Description: `(Optional) For Vault v1.7+. The template to use for username generation. See [Vault docs](https://www.vaultproject.io/docs/concepts/username-templating) ### Oracle Configuration Options`,
				},
				resource.Attribute{
					Name:        "connection_url",
					Description: `(Required) A URL containing connection information. See [Vault docs](https://www.vaultproject.io/api-docs/secret/databases/oracle.html#sample-payload)`,
				},
				resource.Attribute{
					Name:        "max_open_connections",
					Description: `(Optional) The maximum number of open connections to use.`,
				},
				resource.Attribute{
					Name:        "max_idle_connections",
					Description: `(Optional) The maximum number of idle connections to maintain.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Optional) The root credential username used in the connection URL.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional) The root credential password used in the connection URL.`,
				},
				resource.Attribute{
					Name:        "max_connection_lifetime",
					Description: `(Optional) The maximum number of seconds to keep a connection alive for.`,
				},
				resource.Attribute{
					Name:        "username_template",
					Description: `(Optional) For Vault v1.7+. The template to use for username generation. See [Vault docs](https://www.vaultproject.io/docs/concepts/username-templating) ### PostgreSQL Configuration Options`,
				},
				resource.Attribute{
					Name:        "connection_url",
					Description: `(Required) A URL containing connection information. See [Vault docs](https://www.vaultproject.io/api-docs/secret/databases/postgresql.html#sample-payload)`,
				},
				resource.Attribute{
					Name:        "max_open_connections",
					Description: `(Optional) The maximum number of open connections to use.`,
				},
				resource.Attribute{
					Name:        "max_idle_connections",
					Description: `(Optional) The maximum number of idle connections to maintain.`,
				},
				resource.Attribute{
					Name:        "max_connection_lifetime",
					Description: `(Optional) The maximum number of seconds to keep a connection alive for.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Optional) The root credential username used in the connection URL.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional) The root credential password used in the connection URL.`,
				},
				resource.Attribute{
					Name:        "disable_escaping",
					Description: `(Optional) Disable special character escaping in username and password.`,
				},
				resource.Attribute{
					Name:        "username_template",
					Description: `(Optional) For Vault v1.7+. The template to use for username generation. See [Vault docs](https://www.vaultproject.io/docs/concepts/username-templating) ### Redis Configuration Options`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `(Required) The host to connect to.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Required) The username to authenticate with.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Required) The password to authenticate with.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional) The default port to connect to if no port is specified as part of the host.`,
				},
				resource.Attribute{
					Name:        "tls",
					Description: `(Optional) Whether to use TLS when connecting to Redis.`,
				},
				resource.Attribute{
					Name:        "insecure_tls",
					Description: `(Optional) Whether to skip verification of the server certificate when using TLS.`,
				},
				resource.Attribute{
					Name:        "ca_cert",
					Description: `(Optional) The contents of a PEM-encoded CA cert file to use to verify the Redis server's identity. ### Redis ElastiCache Configuration Options`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Required) The configuration endpoint for the ElastiCache cluster to connect to.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Optional) The AWS access key id to use to talk to ElastiCache. If omitted the credentials chain provider is used instead.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional) The AWS secret key id to use to talk to ElastiCache. If omitted the credentials chain provider is used instead.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The AWS region where the ElastiCache cluster is hosted. If omitted the plugin tries to infer the region from the environment. ### AWS Redshift Configuration Options`,
				},
				resource.Attribute{
					Name:        "connection_url",
					Description: `(Required) Specifies the Redshift DSN. See [Vault docs](https://www.vaultproject.io/api-docs/secret/databases/redshift#sample-payload)`,
				},
				resource.Attribute{
					Name:        "max_open_connections",
					Description: `(Optional) The maximum number of open connections to the database.`,
				},
				resource.Attribute{
					Name:        "max_idle_connections",
					Description: `(Optional) The maximum number of idle connections to the database.`,
				},
				resource.Attribute{
					Name:        "max_connection_lifetime",
					Description: `(Optional) The maximum amount of time a connection may be reused.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Optional) The root credential username used in the connection URL.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional) The root credential password used in the connection URL.`,
				},
				resource.Attribute{
					Name:        "disable_escaping",
					Description: `(Optional) Disable special character escaping in username and password.`,
				},
				resource.Attribute{
					Name:        "username_template",
					Description: `(Optional) - [Template](https://www.vaultproject.io/docs/concepts/username-templating) describing how dynamic usernames are generated. ### Snowflake Configuration Options`,
				},
				resource.Attribute{
					Name:        "connection_url",
					Description: `(Required) A URL containing connection information. See [Vault docs](https://www.vaultproject.io/api-docs/secret/databases/snowflake#sample-payload)`,
				},
				resource.Attribute{
					Name:        "max_open_connections",
					Description: `(Optional) The maximum number of open connections to use.`,
				},
				resource.Attribute{
					Name:        "max_idle_connections",
					Description: `(Optional) The maximum number of idle connections to maintain.`,
				},
				resource.Attribute{
					Name:        "max_connection_lifetime",
					Description: `(Optional) The maximum number of seconds to keep a connection alive for.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Optional) The username to be used in the connection (the account admin level).`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional) The password to be used in the connection.`,
				},
				resource.Attribute{
					Name:        "username_template",
					Description: `(Optional) - [Template](https://www.vaultproject.io/docs/concepts/username-templating) describing how dynamic usernames are generated. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "engine_count",
					Description: `The total number of database secrets engines configured. ## Import Database secret backend connections can be imported using the ` + "`" + `path` + "`" + ` e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_database_secrets_mount.db db ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "engine_count",
					Description: `The total number of database secrets engines configured. ## Import Database secret backend connections can be imported using the ` + "`" + `path` + "`" + ` e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_database_secrets_mount.db db ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_egp_policy",
			Category:         "Resources",
			ShortDescription: `Writes Sentinel endpoint governing policies for Vault`,
			Description:      ``,
			Keywords: []string{
				"egp",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) The namespace to provision the resource in. The value should not contain leading or trailing forward slashes. The ` + "`" + `namespace` + "`" + ` is always relative to the provider's configured [namespace](/docs/providers/vault#namespace).`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the policy`,
				},
				resource.Attribute{
					Name:        "paths",
					Description: `(Required) List of paths to which the policy will be applied to`,
				},
				resource.Attribute{
					Name:        "enforcement_level",
					Description: `(Required) Enforcement level of Sentinel policy. Can be either ` + "`" + `advisory` + "`" + ` or ` + "`" + `soft-mandatory` + "`" + ` or ` + "`" + `hard-mandatory` + "`" + ``,
				},
				resource.Attribute{
					Name:        "policy",
					Description: `(Required) String containing a Sentinel policy ## Attributes Reference No additional attributes are exported by this resource.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_gcp_auth_backend",
			Category:         "Resources",
			ShortDescription: `Managing roles in an GCP auth backend in Vault`,
			Description:      ``,
			Keywords: []string{
				"gcp",
				"auth",
				"backend",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) The namespace to provision the resource in. The value should not contain leading or trailing forward slashes. The ` + "`" + `namespace` + "`" + ` is always relative to the provider's configured [namespace](/docs/providers/vault#namespace).`,
				},
				resource.Attribute{
					Name:        "credentials",
					Description: `A JSON string containing the contents of a GCP credentials file. If this value is empty, Vault will try to use Application Default Credentials from the machine on which the Vault server is running.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Optional) The path to mount the auth method — this defaults to 'gcp'.`,
				},
				resource.Attribute{
					Name:        "disable_remount",
					Description: `(Optional) If set, opts out of mount migration on path updates. See here for more info on [Mount Migration](https://www.vaultproject.io/docs/concepts/mount-migration)`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A description of the auth method.`,
				},
				resource.Attribute{
					Name:        "local",
					Description: `(Optional) Specifies if the auth method is local only.`,
				},
				resource.Attribute{
					Name:        "custom_endpoint",
					Description: `(Optional) Specifies overrides to [service endpoints](https://cloud.google.com/apis/design/glossary#api_service_endpoint) used when making API requests. This allows specific requests made during authentication to target alternative service endpoints for use in [Private Google Access](https://cloud.google.com/vpc/docs/configure-private-google-access) environments. Requires Vault 1.11+. Overrides are set at the subdomain level using the following keys: - ` + "`" + `api` + "`" + ` - Replaces the service endpoint used in API requests to ` + "`" + `https://www.googleapis.com` + "`" + `. - ` + "`" + `iam` + "`" + ` - Replaces the service endpoint used in API requests to ` + "`" + `https://iam.googleapis.com` + "`" + `. - ` + "`" + `crm` + "`" + ` - Replaces the service endpoint used in API requests to ` + "`" + `https://cloudresourcemanager.googleapis.com` + "`" + `. - ` + "`" + `compute` + "`" + ` - Replaces the service endpoint used in API requests to ` + "`" + `https://compute.googleapis.com` + "`" + `. The endpoint value provided for a given key has the form of ` + "`" + `scheme://host:port` + "`" + `. The ` + "`" + `scheme://` + "`" + ` and ` + "`" + `:port` + "`" + ` portions of the endpoint value are optional. For more details on the usage of each argument consult the [Vault GCP API documentation](https://www.vaultproject.io/api-docs/auth/gcp#configure). ## Attribute Reference In addition to the fields above, the following attributes are also exposed:`,
				},
				resource.Attribute{
					Name:        "client_id",
					Description: `The Client ID of the credentials`,
				},
				resource.Attribute{
					Name:        "private_key_id",
					Description: `The ID of the private key from the credentials`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `The GCP Project ID`,
				},
				resource.Attribute{
					Name:        "client_email",
					Description: `The clients email associated with the credentials ## Import GCP authentication backends can be imported using the backend name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_gcp_auth_backend.gcp gcp ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "client_id",
					Description: `The Client ID of the credentials`,
				},
				resource.Attribute{
					Name:        "private_key_id",
					Description: `The ID of the private key from the credentials`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `The GCP Project ID`,
				},
				resource.Attribute{
					Name:        "client_email",
					Description: `The clients email associated with the credentials ## Import GCP authentication backends can be imported using the backend name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_gcp_auth_backend.gcp gcp ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_gcp_auth_backend_role",
			Category:         "Resources",
			ShortDescription: `Managing roles in an GCP auth backend in Vault`,
			Description:      ``,
			Keywords: []string{
				"gcp",
				"auth",
				"backend",
				"role",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) The namespace to provision the resource in. The value should not contain leading or trailing forward slashes. The ` + "`" + `namespace` + "`" + ` is always relative to the provider's configured [namespace](/docs/providers/vault#namespace).`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `(Required) Name of the GCP role`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) Type of GCP authentication role (either ` + "`" + `gce` + "`" + ` or ` + "`" + `iam` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "bound_projects",
					Description: `(Optional) An array of GCP project IDs. Only entities belonging to this project can authenticate under the role.`,
				},
				resource.Attribute{
					Name:        "backend",
					Description: `(Optional) Path to the mounted GCP auth backend`,
				},
				resource.Attribute{
					Name:        "bound_service_accounts",
					Description: `(Optional) GCP Service Accounts allowed to issue tokens under this role. (Note:`,
				},
				resource.Attribute{
					Name:        "max_jwt_exp",
					Description: `(Optional) The number of seconds past the time of authentication that the login param JWT must expire within. For example, if a user attempts to login with a token that expires within an hour and this is set to 15 minutes, Vault will return an error prompting the user to create a new signed JWT with a shorter ` + "`" + `exp` + "`" + `. The GCE metadata tokens currently do not allow the ` + "`" + `exp` + "`" + ` claim to be customized.`,
				},
				resource.Attribute{
					Name:        "allow_gce_inference",
					Description: `(Optional) A flag to determine if this role should allow GCE instances to authenticate by inferring service accounts from the GCE identity metadata token. ### ` + "`" + `gce` + "`" + `-only Parameters The following parameters are only valid when the role is of type ` + "`" + `"gce"` + "`" + `:`,
				},
				resource.Attribute{
					Name:        "bound_zones",
					Description: `(Optional) The list of zones that a GCE instance must belong to in order to be authenticated. If bound_instance_groups is provided, it is assumed to be a zonal group and the group must belong to this zone.`,
				},
				resource.Attribute{
					Name:        "bound_regions",
					Description: `(Optional) The list of regions that a GCE instance must belong to in order to be authenticated. If bound_instance_groups is provided, it is assumed to be a regional group and the group must belong to this region. If bound_zones are provided, this attribute is ignored.`,
				},
				resource.Attribute{
					Name:        "bound_instance_groups",
					Description: `(Optional) The instance groups that an authorized instance must belong to in order to be authenticated. If specified, either ` + "`" + `bound_zones` + "`" + ` or ` + "`" + `bound_regions` + "`" + ` must be set too.`,
				},
				resource.Attribute{
					Name:        "bound_labels",
					Description: `(Optional) A comma-separated list of GCP labels formatted as ` + "`" + `"key:value"` + "`" + ` strings that must be set on authorized GCE instances. Because GCP labels are not currently ACL'd, we recommend that this be used in conjunction with other restrictions.`,
				},
				resource.Attribute{
					Name:        "bound_projects",
					Description: `(Optional) GCP Projects that the role exists within For more details on the usage of each argument consult the [Vault GCP API documentation](https://www.vaultproject.io/api-docs/auth/gcp). ### Common Token Arguments These arguments are common across several Authentication Token resources since Vault 1.2.`,
				},
				resource.Attribute{
					Name:        "token_ttl",
					Description: `(Optional) The incremental lifetime for generated tokens in number of seconds. Its current value will be referenced at renewal time.`,
				},
				resource.Attribute{
					Name:        "token_max_ttl",
					Description: `(Optional) The maximum lifetime for generated tokens in number of seconds. Its current value will be referenced at renewal time.`,
				},
				resource.Attribute{
					Name:        "token_period",
					Description: `(Optional) If set, indicates that the token generated using this role should never expire. The token should be renewed within the duration specified by this value. At each renewal, the token's TTL will be set to the value of this field. Specified in seconds.`,
				},
				resource.Attribute{
					Name:        "token_policies",
					Description: `(Optional) List of policies to encode onto generated tokens. Depending on the auth method, this list may be supplemented by user/group/other values.`,
				},
				resource.Attribute{
					Name:        "token_bound_cidrs",
					Description: `(Optional) List of CIDR blocks; if set, specifies blocks of IP addresses which can authenticate successfully, and ties the resulting token to these blocks as well.`,
				},
				resource.Attribute{
					Name:        "token_explicit_max_ttl",
					Description: `(Optional) If set, will encode an [explicit max TTL](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls) onto the token in number of seconds. This is a hard cap even if ` + "`" + `token_ttl` + "`" + ` and ` + "`" + `token_max_ttl` + "`" + ` would otherwise allow a renewal.`,
				},
				resource.Attribute{
					Name:        "token_no_default_policy",
					Description: `(Optional) If set, the default policy will not be set on generated tokens; otherwise it will be added to the policies set in token_policies.`,
				},
				resource.Attribute{
					Name:        "token_num_uses",
					Description: `(Optional) The [maximum number](https://www.vaultproject.io/api-docs/gcp#token_num_uses) of times a generated token may be used (within its lifetime); 0 means unlimited.`,
				},
				resource.Attribute{
					Name:        "token_type",
					Description: `(Optional) The type of token that should be generated. Can be ` + "`" + `service` + "`" + `, ` + "`" + `batch` + "`" + `, or ` + "`" + `default` + "`" + ` to use the mount's tuned default (which unless changed will be ` + "`" + `service` + "`" + ` tokens). For token store roles, there are two additional possibilities: ` + "`" + `default-service` + "`" + ` and ` + "`" + `default-batch` + "`" + ` which specify the type to return unless the client requests a different type at generation time. ## Attribute Reference No additional attributes are exposed by this resource. ## Import GCP authentication roles can be imported using the ` + "`" + `path` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_gcp_auth_backend_role.my_role auth/gcp/role/my_role ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_gcp_secret_backend",
			Category:         "Resources",
			ShortDescription: `Creates an GCP secret backend for Vault.`,
			Description:      ``,
			Keywords: []string{
				"gcp",
				"secret",
				"backend",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) The namespace to provision the resource in. The value should not contain leading or trailing forward slashes. The ` + "`" + `namespace` + "`" + ` is always relative to the provider's configured [namespace](/docs/providers/vault#namespace).`,
				},
				resource.Attribute{
					Name:        "credentials",
					Description: `(Optional) The GCP service account credentials in JSON format. ~>`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Optional) The unique path this backend should be mounted at. Must not begin or end with a ` + "`" + `/` + "`" + `. Defaults to ` + "`" + `gcp` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "disable_remount",
					Description: `(Optional) If set, opts out of mount migration on path updates. See here for more info on [Mount Migration](https://www.vaultproject.io/docs/concepts/mount-migration)`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A human-friendly description for this backend.`,
				},
				resource.Attribute{
					Name:        "default_lease_ttl_seconds",
					Description: `(Optional) The default TTL for credentials issued by this backend. Defaults to '0'.`,
				},
				resource.Attribute{
					Name:        "max_lease_ttl_seconds",
					Description: `(Optional) The maximum TTL that can be requested for credentials issued by this backend. Defaults to '0'.`,
				},
				resource.Attribute{
					Name:        "local",
					Description: `(Optional) Boolean flag that can be explicitly set to true to enforce local mount in HA environment ## Attributes Reference No additional attributes are exported by this resource.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_gcp_secret_roleset",
			Category:         "Resources",
			ShortDescription: `Creates a Roleset for the GCP Secret Backend for Vault.`,
			Description:      ``,
			Keywords: []string{
				"gcp",
				"secret",
				"roleset",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) The namespace to provision the resource in. The value should not contain leading or trailing forward slashes. The ` + "`" + `namespace` + "`" + ` is always relative to the provider's configured [namespace](/docs/providers/vault#namespace).`,
				},
				resource.Attribute{
					Name:        "backend",
					Description: `(Required, Forces new resource) Path where the GCP Secrets Engine is mounted`,
				},
				resource.Attribute{
					Name:        "roleset",
					Description: `(Required, Forces new resource) Name of the Roleset to create`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Required, Forces new resource) Name of the GCP project that this roleset's service account will belong to.`,
				},
				resource.Attribute{
					Name:        "secret_type",
					Description: `(Optional, Forces new resource) Type of secret generated for this role set. Accepted values: ` + "`" + `access_token` + "`" + `, ` + "`" + `service_account_key` + "`" + `. Defaults to ` + "`" + `access_token` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "token_scopes",
					Description: `(Optional, Required for ` + "`" + `secret_type = "access_token"` + "`" + `) List of OAuth scopes to assign to ` + "`" + `access_token` + "`" + ` secrets generated under this role set (` + "`" + `access_token` + "`" + ` role sets only).`,
				},
				resource.Attribute{
					Name:        "binding",
					Description: `(Required) Bindings to create for this roleset. This can be specified multiple times for multiple bindings. Structure is documented below. The ` + "`" + `binding` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "resource",
					Description: `(Required) Resource or resource path for which IAM policy information will be bound. The resource path may be specified in a few different [formats](https://www.vaultproject.io/docs/secrets/gcp/index.html#roleset-bindings).`,
				},
				resource.Attribute{
					Name:        "roles",
					Description: `(Required) List of [GCP IAM roles](https://cloud.google.com/iam/docs/understanding-roles) for the resource. ## Attributes Reference In addition to the fields above, the following attributes are also exposed:`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_gcp_secret_static_account",
			Category:         "Resources",
			ShortDescription: `Creates a Static Account for the GCP Secret Backend for Vault.`,
			Description:      ``,
			Keywords: []string{
				"gcp",
				"secret",
				"static",
				"account",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) The namespace to provision the resource in. The value should not contain leading or trailing forward slashes. The ` + "`" + `namespace` + "`" + ` is always relative to the provider's configured [namespace](/docs/providers/vault#namespace).`,
				},
				resource.Attribute{
					Name:        "backend",
					Description: `(Required, Forces new resource) Path where the GCP Secrets Engine is mounted`,
				},
				resource.Attribute{
					Name:        "static_account",
					Description: `(Required, Forces new resource) Name of the Static Account to create`,
				},
				resource.Attribute{
					Name:        "service_account_email",
					Description: `(Required, Forces new resource) Email of the GCP service account to manage.`,
				},
				resource.Attribute{
					Name:        "secret_type",
					Description: `(Optional, Forces new resource) Type of secret generated for this static account. Accepted values: ` + "`" + `access_token` + "`" + `, ` + "`" + `service_account_key` + "`" + `. Defaults to ` + "`" + `access_token` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "token_scopes",
					Description: `(Optional, Required for ` + "`" + `secret_type = "access_token"` + "`" + `) List of OAuth scopes to assign to ` + "`" + `access_token` + "`" + ` secrets generated under this static account (` + "`" + `access_token` + "`" + ` static accounts only).`,
				},
				resource.Attribute{
					Name:        "binding",
					Description: `(Optional) Bindings to create for this static account. This can be specified multiple times for multiple bindings. Structure is documented below. The ` + "`" + `binding` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "resource",
					Description: `(Required) Resource or resource path for which IAM policy information will be bound. The resource path may be specified in a few different [formats](https://www.vaultproject.io/docs/secrets/gcp/index.html#bindings).`,
				},
				resource.Attribute{
					Name:        "roles",
					Description: `(Required) List of [GCP IAM roles](https://cloud.google.com/iam/docs/understanding-roles) for the resource. ## Attributes Reference In addition to the fields above, the following attributes are also exposed:`,
				},
				resource.Attribute{
					Name:        "service_account_project",
					Description: `Project the service account belongs to. ## Import A static account can be imported using its Vault Path. For example, referencing the example above, ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_gcp_secret_static_account.static_account gcp/static-account/project_viewer ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "service_account_project",
					Description: `Project the service account belongs to. ## Import A static account can be imported using its Vault Path. For example, referencing the example above, ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_gcp_secret_static_account.static_account gcp/static-account/project_viewer ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_generic_endpoint",
			Category:         "Resources",
			ShortDescription: `Writes arbitrary data to a given path in Vault`,
			Description:      ``,
			Keywords: []string{
				"generic",
				"endpoint",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) The namespace to provision the resource in. The value should not contain leading or trailing forward slashes. The ` + "`" + `namespace` + "`" + ` is always relative to the provider's configured [namespace](/docs/providers/vault#namespace).`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Required) The full logical path at which to write the given data. Consult each backend's documentation to see which endpoints support the ` + "`" + `PUT` + "`" + ` methods and to determine whether they also support ` + "`" + `DELETE` + "`" + ` and ` + "`" + `GET` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "data_json",
					Description: `(Required) String containing a JSON-encoded object that will be written to the given path as the secret data.`,
				},
				resource.Attribute{
					Name:        "disable_read",
					Description: `(Optional) True/false. Set this to true if your vault authentication is not able to read the data or if the endpoint does not support the ` + "`" + `GET` + "`" + ` method. Setting this to ` + "`" + `true` + "`" + ` will break drift detection. You should set this to ` + "`" + `true` + "`" + ` for endpoints that are write-only. Defaults to false.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_generic_secret",
			Category:         "Resources",
			ShortDescription: `Writes arbitrary data to a given path in Vault`,
			Description:      ``,
			Keywords: []string{
				"generic",
				"secret",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) The namespace to provision the resource in. The value should not contain leading or trailing forward slashes. The ` + "`" + `namespace` + "`" + ` is always relative to the provider's configured [namespace](/docs/providers/vault#namespace).`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Required) The full logical path at which to write the given data. To write data into the "generic" secret backend mounted in Vault by default, this should be prefixed with ` + "`" + `secret/` + "`" + `. Writing to other backends with this resource is possible; consult each backend's documentation to see which endpoints support the ` + "`" + `PUT` + "`" + ` and ` + "`" + `DELETE` + "`" + ` methods.`,
				},
				resource.Attribute{
					Name:        "data_json",
					Description: `(Required) String containing a JSON-encoded object that will be written as the secret data at the given path.`,
				},
				resource.Attribute{
					Name:        "disable_read",
					Description: `(Optional) true/false. Set this to true if your vault authentication is not able to read the data. Setting this to ` + "`" + `true` + "`" + ` will break drift detection. Defaults to false.`,
				},
				resource.Attribute{
					Name:        "delete_all_versions",
					Description: `(Optional) true/false. Only applicable for kv-v2 stores. If set to ` + "`" + `true` + "`" + `, permanently deletes all versions for the specified key. The default behavior is to only delete the latest version of the secret. ## Required Vault Capabilities Use of this resource requires the ` + "`" + `create` + "`" + ` or ` + "`" + `update` + "`" + ` capability (depending on whether the resource already exists) on the given path, the ` + "`" + `delete` + "`" + ` capability if the resource is removed from configuration, and the ` + "`" + `read` + "`" + ` capability for drift detection (by default). ### Drift Detection This resource does not necessarily need to`,
				},
				resource.Attribute{
					Name:        "data",
					Description: `A mapping whose keys are the top-level data keys returned from Vault and whose values are the corresponding values. This map can only represent string data, so any non-string values returned from Vault are serialized as JSON. ## Import Generic secrets can be imported using the ` + "`" + `path` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_generic_secret.example secret/foo ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "data",
					Description: `A mapping whose keys are the top-level data keys returned from Vault and whose values are the corresponding values. This map can only represent string data, so any non-string values returned from Vault are serialized as JSON. ## Import Generic secrets can be imported using the ` + "`" + `path` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_generic_secret.example secret/foo ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_github_auth_backend",
			Category:         "Resources",
			ShortDescription: `Manages GitHub Auth mounts in Vault.`,
			Description:      ``,
			Keywords: []string{
				"github",
				"auth",
				"backend",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) The namespace to provision the resource in. The value should not contain leading or trailing forward slashes. The ` + "`" + `namespace` + "`" + ` is always relative to the provider's configured [namespace](/docs/providers/vault#namespace).`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Optional) Path where the auth backend is mounted. Defaults to ` + "`" + `auth/github` + "`" + ` if not specified.`,
				},
				resource.Attribute{
					Name:        "disable_remount",
					Description: `(Optional) If set, opts out of mount migration on path updates. See here for more info on [Mount Migration](https://www.vaultproject.io/docs/concepts/mount-migration)`,
				},
				resource.Attribute{
					Name:        "organization",
					Description: `(Required) The organization configured users must be part of.`,
				},
				resource.Attribute{
					Name:        "base_url",
					Description: `(Optional) The API endpoint to use. Useful if you are running GitHub Enterprise or an API-compatible authentication server.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Specifies the description of the mount. This overrides the current stored value, if any.`,
				},
				resource.Attribute{
					Name:        "tune",
					Description: `(Optional) Extra configuration block. Structure is documented below. The ` + "`" + `tune` + "`" + ` block is used to tune the auth backend:`,
				},
				resource.Attribute{
					Name:        "default_lease_ttl",
					Description: `(Optional) Specifies the default time-to-live. If set, this overrides the global default. Must be a valid [duration string](https://golang.org/pkg/time/#ParseDuration)`,
				},
				resource.Attribute{
					Name:        "max_lease_ttl",
					Description: `(Optional) Specifies the maximum time-to-live. If set, this overrides the global default. Must be a valid [duration string](https://golang.org/pkg/time/#ParseDuration)`,
				},
				resource.Attribute{
					Name:        "audit_non_hmac_response_keys",
					Description: `(Optional) Specifies the list of keys that will not be HMAC'd by audit devices in the response data object.`,
				},
				resource.Attribute{
					Name:        "audit_non_hmac_request_keys",
					Description: `(Optional) Specifies the list of keys that will not be HMAC'd by audit devices in the request data object.`,
				},
				resource.Attribute{
					Name:        "listing_visibility",
					Description: `(Optional) Specifies whether to show this mount in the UI-specific listing endpoint. Valid values are "unauth" or "hidden".`,
				},
				resource.Attribute{
					Name:        "passthrough_request_headers",
					Description: `(Optional) List of headers to whitelist and pass from the request to the backend.`,
				},
				resource.Attribute{
					Name:        "allowed_response_headers",
					Description: `(Optional) List of headers to whitelist and allowing a plugin to include them in the response.`,
				},
				resource.Attribute{
					Name:        "token_type",
					Description: `(Optional) Specifies the type of tokens that should be returned by the mount. Valid values are "default-service", "default-batch", "service", "batch". ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "accessor",
					Description: `The mount accessor related to the auth mount. It is useful for integration with [Identity Secrets Engine](https://www.vaultproject.io/docs/secrets/identity/index.html). ### Common Token Arguments These arguments are common across several Authentication Token resources since Vault 1.2.`,
				},
				resource.Attribute{
					Name:        "token_ttl",
					Description: `(Optional) The incremental lifetime for generated tokens in number of seconds. Its current value will be referenced at renewal time.`,
				},
				resource.Attribute{
					Name:        "token_max_ttl",
					Description: `(Optional) The maximum lifetime for generated tokens in number of seconds. Its current value will be referenced at renewal time.`,
				},
				resource.Attribute{
					Name:        "token_period",
					Description: `(Optional) If set, indicates that the token generated using this role should never expire. The token should be renewed within the duration specified by this value. At each renewal, the token's TTL will be set to the value of this field. Specified in seconds.`,
				},
				resource.Attribute{
					Name:        "token_policies",
					Description: `(Optional) List of policies to encode onto generated tokens. Depending on the auth method, this list may be supplemented by user/group/other values.`,
				},
				resource.Attribute{
					Name:        "token_bound_cidrs",
					Description: `(Optional) List of CIDR blocks; if set, specifies blocks of IP addresses which can authenticate successfully, and ties the resulting token to these blocks as well.`,
				},
				resource.Attribute{
					Name:        "token_explicit_max_ttl",
					Description: `(Optional) If set, will encode an [explicit max TTL](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls) onto the token in number of seconds. This is a hard cap even if ` + "`" + `token_ttl` + "`" + ` and ` + "`" + `token_max_ttl` + "`" + ` would otherwise allow a renewal.`,
				},
				resource.Attribute{
					Name:        "token_no_default_policy",
					Description: `(Optional) If set, the default policy will not be set on generated tokens; otherwise it will be added to the policies set in token_policies.`,
				},
				resource.Attribute{
					Name:        "token_num_uses",
					Description: `(Optional) The [maximum number](https://www.vaultproject.io/api-docs/github#token_num_uses) of times a generated token may be used (within its lifetime); 0 means unlimited.`,
				},
				resource.Attribute{
					Name:        "token_type",
					Description: `(Optional) The type of token that should be generated. Can be ` + "`" + `service` + "`" + `, ` + "`" + `batch` + "`" + `, or ` + "`" + `default` + "`" + ` to use the mount's tuned default (which unless changed will be ` + "`" + `service` + "`" + ` tokens). For token store roles, there are two additional possibilities: ` + "`" + `default-service` + "`" + ` and ` + "`" + `default-batch` + "`" + ` which specify the type to return unless the client requests a different type at generation time. ## Import GitHub authentication mounts can be imported using the ` + "`" + `path` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_github_auth_backend.example github ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "accessor",
					Description: `The mount accessor related to the auth mount. It is useful for integration with [Identity Secrets Engine](https://www.vaultproject.io/docs/secrets/identity/index.html). ### Common Token Arguments These arguments are common across several Authentication Token resources since Vault 1.2.`,
				},
				resource.Attribute{
					Name:        "token_ttl",
					Description: `(Optional) The incremental lifetime for generated tokens in number of seconds. Its current value will be referenced at renewal time.`,
				},
				resource.Attribute{
					Name:        "token_max_ttl",
					Description: `(Optional) The maximum lifetime for generated tokens in number of seconds. Its current value will be referenced at renewal time.`,
				},
				resource.Attribute{
					Name:        "token_period",
					Description: `(Optional) If set, indicates that the token generated using this role should never expire. The token should be renewed within the duration specified by this value. At each renewal, the token's TTL will be set to the value of this field. Specified in seconds.`,
				},
				resource.Attribute{
					Name:        "token_policies",
					Description: `(Optional) List of policies to encode onto generated tokens. Depending on the auth method, this list may be supplemented by user/group/other values.`,
				},
				resource.Attribute{
					Name:        "token_bound_cidrs",
					Description: `(Optional) List of CIDR blocks; if set, specifies blocks of IP addresses which can authenticate successfully, and ties the resulting token to these blocks as well.`,
				},
				resource.Attribute{
					Name:        "token_explicit_max_ttl",
					Description: `(Optional) If set, will encode an [explicit max TTL](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls) onto the token in number of seconds. This is a hard cap even if ` + "`" + `token_ttl` + "`" + ` and ` + "`" + `token_max_ttl` + "`" + ` would otherwise allow a renewal.`,
				},
				resource.Attribute{
					Name:        "token_no_default_policy",
					Description: `(Optional) If set, the default policy will not be set on generated tokens; otherwise it will be added to the policies set in token_policies.`,
				},
				resource.Attribute{
					Name:        "token_num_uses",
					Description: `(Optional) The [maximum number](https://www.vaultproject.io/api-docs/github#token_num_uses) of times a generated token may be used (within its lifetime); 0 means unlimited.`,
				},
				resource.Attribute{
					Name:        "token_type",
					Description: `(Optional) The type of token that should be generated. Can be ` + "`" + `service` + "`" + `, ` + "`" + `batch` + "`" + `, or ` + "`" + `default` + "`" + ` to use the mount's tuned default (which unless changed will be ` + "`" + `service` + "`" + ` tokens). For token store roles, there are two additional possibilities: ` + "`" + `default-service` + "`" + ` and ` + "`" + `default-batch` + "`" + ` which specify the type to return unless the client requests a different type at generation time. ## Import GitHub authentication mounts can be imported using the ` + "`" + `path` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_github_auth_backend.example github ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_github_team",
			Category:         "Resources",
			ShortDescription: `Manages Team mappings for Github Auth backend mounts in Vault.`,
			Description:      ``,
			Keywords: []string{
				"github",
				"team",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) The namespace to provision the resource in. The value should not contain leading or trailing forward slashes. The ` + "`" + `namespace` + "`" + ` is always relative to the provider's configured [namespace](/docs/providers/vault#namespace).`,
				},
				resource.Attribute{
					Name:        "backend",
					Description: `(Required) Path where the github auth backend is mounted. Defaults to ` + "`" + `github` + "`" + ` if not specified.`,
				},
				resource.Attribute{
					Name:        "team",
					Description: `(Required) GitHub team name in "slugified" format, for example: Terraform Developers -> ` + "`" + `terraform-developers` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "policies",
					Description: `(Optional) An array of strings specifying the policies to be set on tokens issued using this role. ## Attributes Reference No additional attributes are exported by this resource. ## Import Github team mappings can be imported using the ` + "`" + `path` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_github_team.tf_devs auth/github/map/teams/terraform-developers ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_github_user",
			Category:         "Resources",
			ShortDescription: `Manages User mappings for Github Auth backend mounts in Vault.`,
			Description:      ``,
			Keywords: []string{
				"github",
				"user",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) The namespace to provision the resource in. The value should not contain leading or trailing forward slashes. The ` + "`" + `namespace` + "`" + ` is always relative to the provider's configured [namespace](/docs/providers/vault#namespace).`,
				},
				resource.Attribute{
					Name:        "backend",
					Description: `(Required) Path where the github auth backend is mounted. Defaults to ` + "`" + `github` + "`" + ` if not specified.`,
				},
				resource.Attribute{
					Name:        "user",
					Description: `(Required) GitHub user name.`,
				},
				resource.Attribute{
					Name:        "policies",
					Description: `(Optional) An array of strings specifying the policies to be set on tokens issued using this role. ## Attributes Reference No additional attributes are exported by this resource. ## Import Github user mappings can be imported using the ` + "`" + `path` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_github_user.tf_user auth/github/map/users/john.doe ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_identity_entity",
			Category:         "Resources",
			ShortDescription: `Creates an Identity Entity for Vault.`,
			Description:      ``,
			Keywords: []string{
				"identity",
				"entity",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) The namespace to provision the resource in. The value should not contain leading or trailing forward slashes. The ` + "`" + `namespace` + "`" + ` is always relative to the provider's configured [namespace](/docs/providers/vault#namespace).`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the identity entity to create.`,
				},
				resource.Attribute{
					Name:        "policies",
					Description: `(Optional) A list of policies to apply to the entity.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `(Optional) A Map of additional metadata to associate with the user.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `(Optional) True/false Is this entity currently disabled. Defaults to ` + "`" + `false` + "`" + ``,
				},
				resource.Attribute{
					Name:        "external_policies",
					Description: `(Optional) ` + "`" + `false` + "`" + ` by default. If set to ` + "`" + `true` + "`" + `, this resource will ignore any policies return from Vault or specified in the resource. You can use [` + "`" + `vault_identity_entity_policies` + "`" + `](identity_entity_policies.html) to manage policies for this entity in a decoupled manner. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ` + "`" + `id` + "`" + ` of the created entity. ## Import Identity entity can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_identity_entity.test "ae6f8ued-0f1a-9f6b-2915-1a2be20dc053" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ` + "`" + `id` + "`" + ` of the created entity. ## Import Identity entity can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_identity_entity.test "ae6f8ued-0f1a-9f6b-2915-1a2be20dc053" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_identity_entity_alias",
			Category:         "Resources",
			ShortDescription: `Creates an Identity Entity Alias for Vault.`,
			Description:      ``,
			Keywords: []string{
				"identity",
				"entity",
				"alias",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) The namespace to provision the resource in. The value should not contain leading or trailing forward slashes. The ` + "`" + `namespace` + "`" + ` is always relative to the provider's configured [namespace](/docs/providers/vault#namespace).`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the alias. Name should be the identifier of the client in the authentication source. For example, if the alias belongs to userpass backend, the name should be a valid username within userpass backend. If alias belongs to GitHub, it should be the GitHub username.`,
				},
				resource.Attribute{
					Name:        "mount_accessor",
					Description: `(Required) Accessor of the mount to which the alias should belong to.`,
				},
				resource.Attribute{
					Name:        "canonical_id",
					Description: `(Required) Entity ID to which this alias belongs to. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the entity alias. ## Import Identity entity alias can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_identity_entity_alias.test "3856fb4d-3c91-dcaf-2401-68f446796bfb" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the entity alias. ## Import Identity entity alias can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_identity_entity_alias.test "3856fb4d-3c91-dcaf-2401-68f446796bfb" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_identity_entity_policies",
			Category:         "Resources",
			ShortDescription: `Manages policies for an Identity Entity for Vault.`,
			Description:      ``,
			Keywords: []string{
				"identity",
				"entity",
				"policies",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) The namespace to provision the resource in. The value should not contain leading or trailing forward slashes. The ` + "`" + `namespace` + "`" + ` is always relative to the provider's configured [namespace](/docs/providers/vault#namespace).`,
				},
				resource.Attribute{
					Name:        "policies",
					Description: `(Required) List of policies to assign to the entity`,
				},
				resource.Attribute{
					Name:        "entity_id",
					Description: `(Required) Entity ID to assign policies to.`,
				},
				resource.Attribute{
					Name:        "exclusive",
					Description: `(Optional) Defaults to ` + "`" + `true` + "`" + `. If ` + "`" + `true` + "`" + `, this resource will take exclusive control of the policies assigned to the entity and will set it equal to what is specified in the resource. If set to ` + "`" + `false` + "`" + `, this resource will simply ensure that the policies specified in the resource are present in the entity. When destroying the resource, the resource will ensure that the policies specified in the resource are removed. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "entity_name",
					Description: `The name of the entity that are assigned the policies.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "entity_name",
					Description: `The name of the entity that are assigned the policies.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_identity_group",
			Category:         "Resources",
			ShortDescription: `Creates an Identity Group for Vault.`,
			Description:      ``,
			Keywords: []string{
				"identity",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) The namespace to provision the resource in. The value should not contain leading or trailing forward slashes. The ` + "`" + `namespace` + "`" + ` is always relative to the provider's configured [namespace](/docs/providers/vault#namespace).`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required, Forces new resource) Name of the identity group to create.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional, Forces new resource) Type of the group, internal or external. Defaults to ` + "`" + `internal` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "policies",
					Description: `(Optional) A list of policies to apply to the group.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `(Optional) A Map of additional metadata to associate with the group.`,
				},
				resource.Attribute{
					Name:        "member_group_ids",
					Description: `(Optional) A list of Group IDs to be assigned as group members. Not allowed on ` + "`" + `external` + "`" + ` groups.`,
				},
				resource.Attribute{
					Name:        "member_entity_ids",
					Description: `(Optional) A list of Entity IDs to be assigned as group members. Not allowed on ` + "`" + `external` + "`" + ` groups.`,
				},
				resource.Attribute{
					Name:        "external_policies",
					Description: `(Optional) ` + "`" + `false` + "`" + ` by default. If set to ` + "`" + `true` + "`" + `, this resource will ignore any policies returned from Vault or specified in the resource. You can use [` + "`" + `vault_identity_group_policies` + "`" + `](identity_group_policies.html) to manage policies for this group in a decoupled manner.`,
				},
				resource.Attribute{
					Name:        "external_member_entity_ids",
					Description: `(Optional) ` + "`" + `false` + "`" + ` by default. If set to ` + "`" + `true` + "`" + `, this resource will ignore any Entity IDs returned from Vault or specified in the resource. You can use [` + "`" + `vault_identity_group_member_entity_ids` + "`" + `](identity_group_member_entity_ids.html) to manage Entity IDs for this group in a decoupled manner.`,
				},
				resource.Attribute{
					Name:        "external_member_group_ids",
					Description: `(Optional) ` + "`" + `false` + "`" + ` by default. If set to ` + "`" + `true` + "`" + `, this resource will ignore any Group IDs returned from Vault or specified in the resource. You can use [` + "`" + `vault_identity_group_member_group_ids` + "`" + `](identity_group_member_group_ids.html) to manage Group IDs for this group in a decoupled manner. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ` + "`" + `id` + "`" + ` of the created group. ## Import Identity group can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_identity_group.test 'fcbf1efb-2b69-4209-bed8-811e3475dad3' ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ` + "`" + `id` + "`" + ` of the created group. ## Import Identity group can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_identity_group.test 'fcbf1efb-2b69-4209-bed8-811e3475dad3' ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_identity_group_alias",
			Category:         "Resources",
			ShortDescription: `Creates an Identity Group Alias for Vault.`,
			Description:      ``,
			Keywords: []string{
				"identity",
				"group",
				"alias",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) The namespace to provision the resource in. The value should not contain leading or trailing forward slashes. The ` + "`" + `namespace` + "`" + ` is always relative to the provider's configured [namespace](/docs/providers/vault#namespace).`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required, Forces new resource) Name of the group alias to create.`,
				},
				resource.Attribute{
					Name:        "mount_accessor",
					Description: `(Required) Mount accessor of the authentication backend to which this alias belongs to.`,
				},
				resource.Attribute{
					Name:        "canonical_id",
					Description: `(Required) ID of the group to which this is an alias. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ` + "`" + `id` + "`" + ` of the created group alias. ## Import The group alias can be imported with the group alias ` + "`" + `id` + "`" + `, for example: ` + "`" + `` + "`" + `` + "`" + `shell $ terraform import vault_identity_group_alias.group-alias id ` + "`" + `` + "`" + `` + "`" + ` Group aliases can also be imported using the UUID of the alias record, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import vault_identity_group_alias.alias_name 63104e20-88e4-11eb-8d04-cf7ac9d60157 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ` + "`" + `id` + "`" + ` of the created group alias. ## Import The group alias can be imported with the group alias ` + "`" + `id` + "`" + `, for example: ` + "`" + `` + "`" + `` + "`" + `shell $ terraform import vault_identity_group_alias.group-alias id ` + "`" + `` + "`" + `` + "`" + ` Group aliases can also be imported using the UUID of the alias record, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import vault_identity_group_alias.alias_name 63104e20-88e4-11eb-8d04-cf7ac9d60157 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_identity_group_member_entity_ids",
			Category:         "Resources",
			ShortDescription: `Manages member entities for an Identity Group for Vault.`,
			Description:      ``,
			Keywords: []string{
				"identity",
				"group",
				"member",
				"entity",
				"ids",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) The namespace to provision the resource in. The value should not contain leading or trailing forward slashes. The ` + "`" + `namespace` + "`" + ` is always relative to the provider's configured [namespace](/docs/providers/vault#namespace).`,
				},
				resource.Attribute{
					Name:        "member_entity_ids",
					Description: `(Required) List of member entities that belong to the group`,
				},
				resource.Attribute{
					Name:        "group_id",
					Description: `(Required) Group ID to assign member entities to.`,
				},
				resource.Attribute{
					Name:        "exclusive",
					Description: `(Optional) Defaults to ` + "`" + `true` + "`" + `. If ` + "`" + `true` + "`" + `, this resource will take exclusive control of the member entities that belong to the group and will set it equal to what is specified in the resource. If set to ` + "`" + `false` + "`" + `, this resource will simply ensure that the member entities specified in the resource are present in the group. When destroying the resource, the resource will ensure that the member entities specified in the resource are removed. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "group_name",
					Description: `The name of the group that are assigned the member entities.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "group_name",
					Description: `The name of the group that are assigned the member entities.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_identity_group_member_group_ids",
			Category:         "Resources",
			ShortDescription: `Manages member groups for an Identity Group for Vault.`,
			Description:      ``,
			Keywords: []string{
				"identity",
				"group",
				"member",
				"ids",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) The namespace to provision the resource in. The value should not contain leading or trailing forward slashes. The ` + "`" + `namespace` + "`" + ` is always relative to the provider's configured [namespace](/docs/providers/vault#namespace).`,
				},
				resource.Attribute{
					Name:        "member_group_ids",
					Description: `(Required) List of member groups that belong to the group`,
				},
				resource.Attribute{
					Name:        "group_id",
					Description: `(Required) Group ID to assign member entities to.`,
				},
				resource.Attribute{
					Name:        "exclusive",
					Description: `(Optional) Defaults to ` + "`" + `true` + "`" + `. If ` + "`" + `true` + "`" + `, this resource will take exclusive control of the member groups that belong to the group and will set it equal to what is specified in the resource. If set to ` + "`" + `false` + "`" + `, this resource will simply ensure that the member groups specified in the resource are present in the group. When destroying the resource, the resource will ensure that the member groups specified in the resource are removed. ## Attributes Reference No additional attributes are exported by this resource.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_identity_group_policies",
			Category:         "Resources",
			ShortDescription: `Manages policies for an Identity Group for Vault.`,
			Description:      ``,
			Keywords: []string{
				"identity",
				"group",
				"policies",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) The namespace to provision the resource in. The value should not contain leading or trailing forward slashes. The ` + "`" + `namespace` + "`" + ` is always relative to the provider's configured [namespace](/docs/providers/vault#namespace).`,
				},
				resource.Attribute{
					Name:        "policies",
					Description: `(Required) List of policies to assign to the group`,
				},
				resource.Attribute{
					Name:        "group_id",
					Description: `(Required) Group ID to assign policies to.`,
				},
				resource.Attribute{
					Name:        "exclusive",
					Description: `(Optional) Defaults to ` + "`" + `true` + "`" + `. If ` + "`" + `true` + "`" + `, this resource will take exclusive control of the policies assigned to the group and will set it equal to what is specified in the resource. If set to ` + "`" + `false` + "`" + `, this resource will simply ensure that the policies specified in the resource are present in the group. When destroying the resource, the resource will ensure that the policies specified in the resource are removed. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "group_name",
					Description: `The name of the group that are assigned the policies.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "group_name",
					Description: `The name of the group that are assigned the policies.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_identity_mfa_duo",
			Category:         "Resources",
			ShortDescription: `Resource for configuring the duo MFA method.`,
			Description:      ``,
			Keywords: []string{
				"identity",
				"mfa",
				"duo",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "api_hostname",
					Description: `(Required) API hostname for Duo`,
				},
				resource.Attribute{
					Name:        "integration_key",
					Description: `(Required) Integration key for Duo`,
				},
				resource.Attribute{
					Name:        "secret_key",
					Description: `(Required) Secret key for Duo`,
				},
				resource.Attribute{
					Name:        "mount_accessor",
					Description: `(Optional) Mount accessor.`,
				},
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) Target namespace. (requires Enterprise)`,
				},
				resource.Attribute{
					Name:        "push_info",
					Description: `(Optional) Push information for Duo.`,
				},
				resource.Attribute{
					Name:        "use_passcode",
					Description: `(Optional) Require passcode upon MFA validation.`,
				},
				resource.Attribute{
					Name:        "username_format",
					Description: `(Optional) A template string for mapping Identity names to MFA methods.`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `(Optional) Resource UUID. ## Attributes Reference In addition to the fields above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "method_id",
					Description: `Method ID.`,
				},
				resource.Attribute{
					Name:        "namespace_id",
					Description: `Method's namespace ID.`,
				},
				resource.Attribute{
					Name:        "namespace_path",
					Description: `Method's namespace path.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `MFA type. ## Import Resource can be imported using its ` + "`" + `uuid` + "`" + ` field, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_identity_mfa_duo.example 0d89c36a-4ff5-4d70-8749-bb6a5598aeec ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "method_id",
					Description: `Method ID.`,
				},
				resource.Attribute{
					Name:        "namespace_id",
					Description: `Method's namespace ID.`,
				},
				resource.Attribute{
					Name:        "namespace_path",
					Description: `Method's namespace path.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `MFA type. ## Import Resource can be imported using its ` + "`" + `uuid` + "`" + ` field, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_identity_mfa_duo.example 0d89c36a-4ff5-4d70-8749-bb6a5598aeec ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_identity_mfa_login_enforcement",
			Category:         "Resources",
			ShortDescription: `Resource for configuring MFA login-enforcement`,
			Description:      ``,
			Keywords: []string{
				"identity",
				"mfa",
				"login",
				"enforcement",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "mfa_method_ids",
					Description: `(Required) Set of MFA method UUIDs.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Login enforcement name.`,
				},
				resource.Attribute{
					Name:        "auth_method_accessors",
					Description: `(Optional) Set of auth method accessor IDs.`,
				},
				resource.Attribute{
					Name:        "auth_method_types",
					Description: `(Optional) Set of auth method types.`,
				},
				resource.Attribute{
					Name:        "identity_entity_ids",
					Description: `(Optional) Set of identity entity IDs.`,
				},
				resource.Attribute{
					Name:        "identity_group_ids",
					Description: `(Optional) Set of identity group IDs.`,
				},
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) Target namespace. (requires Enterprise)`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `(Optional) Resource UUID. ## Attributes Reference In addition to the fields above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "namespace_id",
					Description: `Method's namespace ID.`,
				},
				resource.Attribute{
					Name:        "namespace_path",
					Description: `Method's namespace path. ## Import Resource can be imported using its ` + "`" + `name` + "`" + ` field, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_identity_mfa_login_enforcement.example default ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "namespace_id",
					Description: `Method's namespace ID.`,
				},
				resource.Attribute{
					Name:        "namespace_path",
					Description: `Method's namespace path. ## Import Resource can be imported using its ` + "`" + `name` + "`" + ` field, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_identity_mfa_login_enforcement.example default ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_identity_mfa_okta",
			Category:         "Resources",
			ShortDescription: `Resource for configuring the okta MFA method.`,
			Description:      ``,
			Keywords: []string{
				"identity",
				"mfa",
				"okta",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "api_token",
					Description: `(Required) Okta API token.`,
				},
				resource.Attribute{
					Name:        "org_name",
					Description: `(Required) Name of the organization to be used in the Okta API.`,
				},
				resource.Attribute{
					Name:        "base_url",
					Description: `(Optional) The base domain to use for API requests.`,
				},
				resource.Attribute{
					Name:        "mount_accessor",
					Description: `(Optional) Mount accessor.`,
				},
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) Target namespace. (requires Enterprise)`,
				},
				resource.Attribute{
					Name:        "primary_email",
					Description: `(Optional) Only match the primary email for the account.`,
				},
				resource.Attribute{
					Name:        "username_format",
					Description: `(Optional) A template string for mapping Identity names to MFA methods.`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `(Optional) Resource UUID. ## Attributes Reference In addition to the fields above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "method_id",
					Description: `Method ID.`,
				},
				resource.Attribute{
					Name:        "namespace_id",
					Description: `Method's namespace ID.`,
				},
				resource.Attribute{
					Name:        "namespace_path",
					Description: `Method's namespace path.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `MFA type. ## Import Resource can be imported using its ` + "`" + `uuid` + "`" + ` field, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_identity_mfa_okta.example 0d89c36a-4ff5-4d70-8749-bb6a5598aeec ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "method_id",
					Description: `Method ID.`,
				},
				resource.Attribute{
					Name:        "namespace_id",
					Description: `Method's namespace ID.`,
				},
				resource.Attribute{
					Name:        "namespace_path",
					Description: `Method's namespace path.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `MFA type. ## Import Resource can be imported using its ` + "`" + `uuid` + "`" + ` field, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_identity_mfa_okta.example 0d89c36a-4ff5-4d70-8749-bb6a5598aeec ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_identity_mfa_pingid",
			Category:         "Resources",
			ShortDescription: `Resource for configuring the pingid MFA method.`,
			Description:      ``,
			Keywords: []string{
				"identity",
				"mfa",
				"pingid",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "settings_file_base64",
					Description: `(Required) A base64-encoded third-party settings contents as retrieved from PingID's configuration page.`,
				},
				resource.Attribute{
					Name:        "admin_url",
					Description: `(Optional) The admin URL, derived from "settings_file_base64"`,
				},
				resource.Attribute{
					Name:        "authenticator_url",
					Description: `(Optional) A unique identifier of the organization, derived from "settings_file_base64"`,
				},
				resource.Attribute{
					Name:        "idp_url",
					Description: `(Optional) The IDP URL, derived from "settings_file_base64"`,
				},
				resource.Attribute{
					Name:        "mount_accessor",
					Description: `(Optional) Mount accessor.`,
				},
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) Target namespace. (requires Enterprise)`,
				},
				resource.Attribute{
					Name:        "org_alias",
					Description: `(Optional) The name of the PingID client organization, derived from "settings_file_base64"`,
				},
				resource.Attribute{
					Name:        "use_signature",
					Description: `(Optional) Use signature value, derived from "settings_file_base64"`,
				},
				resource.Attribute{
					Name:        "username_format",
					Description: `(Optional) A template string for mapping Identity names to MFA methods.`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `(Optional) Resource UUID. ## Attributes Reference In addition to the fields above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "method_id",
					Description: `Method ID.`,
				},
				resource.Attribute{
					Name:        "namespace_id",
					Description: `Method's namespace ID.`,
				},
				resource.Attribute{
					Name:        "namespace_path",
					Description: `Method's namespace path.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `MFA type. ## Import Resource can be imported using its ` + "`" + `uuid` + "`" + ` field, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_identity_mfa_pingid.example 0d89c36a-4ff5-4d70-8749-bb6a5598aeec ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "method_id",
					Description: `Method ID.`,
				},
				resource.Attribute{
					Name:        "namespace_id",
					Description: `Method's namespace ID.`,
				},
				resource.Attribute{
					Name:        "namespace_path",
					Description: `Method's namespace path.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `MFA type. ## Import Resource can be imported using its ` + "`" + `uuid` + "`" + ` field, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_identity_mfa_pingid.example 0d89c36a-4ff5-4d70-8749-bb6a5598aeec ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_identity_mfa_totp",
			Category:         "Resources",
			ShortDescription: `Resource for configuring the totp MFA method.`,
			Description:      ``,
			Keywords: []string{
				"identity",
				"mfa",
				"totp",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "issuer",
					Description: `(Required) The name of the key's issuing organization.`,
				},
				resource.Attribute{
					Name:        "algorithm",
					Description: `(Optional) Specifies the hashing algorithm used to generate the TOTP code. Options include SHA1, SHA256, SHA512.`,
				},
				resource.Attribute{
					Name:        "digits",
					Description: `(Optional) The number of digits in the generated TOTP token. This value can either be 6 or 8`,
				},
				resource.Attribute{
					Name:        "key_size",
					Description: `(Optional) Specifies the size in bytes of the generated key.`,
				},
				resource.Attribute{
					Name:        "max_validation_attempts",
					Description: `(Optional) The maximum number of consecutive failed validation attempts allowed.`,
				},
				resource.Attribute{
					Name:        "mount_accessor",
					Description: `(Optional) Mount accessor.`,
				},
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) Target namespace. (requires Enterprise)`,
				},
				resource.Attribute{
					Name:        "period",
					Description: `(Optional) The length of time in seconds used to generate a counter for the TOTP token calculation.`,
				},
				resource.Attribute{
					Name:        "qr_size",
					Description: `(Optional) The pixel size of the generated square QR code.`,
				},
				resource.Attribute{
					Name:        "skew",
					Description: `(Optional) The number of delay periods that are allowed when validating a TOTP token. This value can either be 0 or 1.`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `(Optional) Resource UUID. ## Attributes Reference In addition to the fields above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "method_id",
					Description: `Method ID.`,
				},
				resource.Attribute{
					Name:        "namespace_id",
					Description: `Method's namespace ID.`,
				},
				resource.Attribute{
					Name:        "namespace_path",
					Description: `Method's namespace path.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `MFA type. ## Import Resource can be imported using its ` + "`" + `uuid` + "`" + ` field, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_identity_mfa_totp.example 0d89c36a-4ff5-4d70-8749-bb6a5598aeec ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "method_id",
					Description: `Method ID.`,
				},
				resource.Attribute{
					Name:        "namespace_id",
					Description: `Method's namespace ID.`,
				},
				resource.Attribute{
					Name:        "namespace_path",
					Description: `Method's namespace path.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `MFA type. ## Import Resource can be imported using its ` + "`" + `uuid` + "`" + ` field, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_identity_mfa_totp.example 0d89c36a-4ff5-4d70-8749-bb6a5598aeec ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_identity_oidc",
			Category:         "Resources",
			ShortDescription: `Configure the Identity Tokens Backend for Vault`,
			Description:      ``,
			Keywords: []string{
				"identity",
				"oidc",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) The namespace to provision the resource in. The value should not contain leading or trailing forward slashes. The ` + "`" + `namespace` + "`" + ` is always relative to the provider's configured [namespace](/docs/providers/vault#namespace).`,
				},
				resource.Attribute{
					Name:        "issuer",
					Description: `(Optional) Issuer URL to be used in the iss claim of the token. If not set, Vault's ` + "`" + `api_addr` + "`" + ` will be used. The issuer is a case sensitive URL using the https scheme that contains scheme, host, and optionally, port number and path components, but no query or fragment components. ## Attributes Reference No additional attributes are exposed by this resource.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_identity_oidc_assignment",
			Category:         "Resources",
			ShortDescription: `Provision OIDC Assignments in Vault.`,
			Description:      ``,
			Keywords: []string{
				"identity",
				"oidc",
				"assignment",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) The namespace to provision the resource in. The value should not contain leading or trailing forward slashes. The ` + "`" + `namespace` + "`" + ` is always relative to the provider's configured [namespace](/docs/providers/vault#namespace).`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the assignment.`,
				},
				resource.Attribute{
					Name:        "entity_ids",
					Description: `(Optional) A set of Vault entity IDs.`,
				},
				resource.Attribute{
					Name:        "group_ids",
					Description: `(Optional) A set of Vault group IDs. ## Attributes Reference No additional attributes are exported by this resource. ## Import OIDC Assignments can be imported using the ` + "`" + `name` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_identity_oidc_assignment.default assignment ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_identity_oidc_client",
			Category:         "Resources",
			ShortDescription: `Provision OIDC Clients in Vault.`,
			Description:      ``,
			Keywords: []string{
				"identity",
				"oidc",
				"client",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) The namespace to provision the resource in. The value should not contain leading or trailing forward slashes. The ` + "`" + `namespace` + "`" + ` is always relative to the provider's configured [namespace](/docs/providers/vault#namespace).`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the client.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Optional) A reference to a named key resource in Vault. This cannot be modified after creation. If not provided, the ` + "`" + `default` + "`" + ` key is used.`,
				},
				resource.Attribute{
					Name:        "redirect_uris",
					Description: `(Optional) Redirection URI values used by the client. One of these values must exactly match the ` + "`" + `redirect_uri` + "`" + ` parameter value used in each authentication request.`,
				},
				resource.Attribute{
					Name:        "assignments",
					Description: `(Optional) A list of assignment resources associated with the client.`,
				},
				resource.Attribute{
					Name:        "id_token_ttl",
					Description: `(Optional) The time-to-live for ID tokens obtained by the client. The value should be less than the ` + "`" + `verification_ttl` + "`" + ` on the key.`,
				},
				resource.Attribute{
					Name:        "access_token_ttl",
					Description: `(Optional) The time-to-live for access tokens obtained by the client.`,
				},
				resource.Attribute{
					Name:        "client_type",
					Description: `(Optional) The client type based on its ability to maintain confidentiality of credentials. The following client types are supported: ` + "`" + `confidential` + "`" + `, ` + "`" + `public` + "`" + `. Defaults to ` + "`" + `confidential` + "`" + `. ## Attributes Reference No additional attributes are exported by this resource. ## Import OIDC Clients can be imported using the ` + "`" + `name` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_identity_oidc_client.test my-app ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_identity_oidc_key",
			Category:         "Resources",
			ShortDescription: `Creates an Identity OIDC Named Key for Vault`,
			Description:      ``,
			Keywords: []string{
				"identity",
				"oidc",
				"key",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) The namespace to provision the resource in. The value should not contain leading or trailing forward slashes. The ` + "`" + `namespace` + "`" + ` is always relative to the provider's configured [namespace](/docs/providers/vault#namespace).`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required; Forces new resource) Name of the OIDC Key to create.`,
				},
				resource.Attribute{
					Name:        "rotation_period",
					Description: `(Optional) How often to generate a new signing key in number of seconds`,
				},
				resource.Attribute{
					Name:        "verification_ttl",
					Description: `(Optional) "Controls how long the public portion of a signing key will be available for verification after being rotated in seconds.`,
				},
				resource.Attribute{
					Name:        "algorithm",
					Description: `(Optional) Signing algorithm to use. Signing algorithm to use. Allowed values are: RS256 (default), RS384, RS512, ES256, ES384, ES512, EdDSA.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The name of the created key. ## Import The key can be imported with the key name, for example: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_identity_oidc_key.key key ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The name of the created key. ## Import The key can be imported with the key name, for example: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_identity_oidc_key.key key ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_identity_oidc_key_allowed_client_id",
			Category:         "Resources",
			ShortDescription: `Allows an Identity OIDC Role to use an OIDC Named key.`,
			Description:      ``,
			Keywords: []string{
				"identity",
				"oidc",
				"key",
				"allowed",
				"client",
				"id",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) The namespace to provision the resource in. The value should not contain leading or trailing forward slashes. The ` + "`" + `namespace` + "`" + ` is always relative to the provider's configured [namespace](/docs/providers/vault#namespace).`,
				},
				resource.Attribute{
					Name:        "key_name",
					Description: `(Required; Forces new resource) Name of the OIDC Key allow the Client ID.`,
				},
				resource.Attribute{
					Name:        "allowed_client_id",
					Description: `(Required; Forces new resource) Client ID to allow usage with the OIDC named key`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_identity_oidc_provider",
			Category:         "Resources",
			ShortDescription: `Provision OIDC Providers in Vault.`,
			Description:      ``,
			Keywords: []string{
				"identity",
				"oidc",
				"provider",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) The namespace to provision the resource in. The value should not contain leading or trailing forward slashes. The ` + "`" + `namespace` + "`" + ` is always relative to the provider's configured [namespace](/docs/providers/vault#namespace).`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the provider.`,
				},
				resource.Attribute{
					Name:        "https_enabled",
					Description: `(Optional) Set to true if the issuer endpoint uses HTTPS.`,
				},
				resource.Attribute{
					Name:        "issuer_host",
					Description: `(Optional) The host for the issuer. Can be either host or host:port.`,
				},
				resource.Attribute{
					Name:        "allowed_client_ids",
					Description: `(Optional) The client IDs that are permitted to use the provider. If empty, no clients are allowed. If ` + "`" + ``,
				},
				resource.Attribute{
					Name:        "scopes_supported",
					Description: `(Optional) The scopes available for requesting on the provider. ## Attributes Reference In addition to the arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "issuer",
					Description: `Specifies what will be used as the ` + "`" + `scheme://host:port` + "`" + ` component for the ` + "`" + `iss` + "`" + ` claim of ID tokens. This value is computed using the ` + "`" + `issuer_host` + "`" + ` and ` + "`" + `https_enabled` + "`" + ` fields. ## Import OIDC Providers can be imported using the ` + "`" + `name` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_identity_oidc_provider.test my-provider ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "issuer",
					Description: `Specifies what will be used as the ` + "`" + `scheme://host:port` + "`" + ` component for the ` + "`" + `iss` + "`" + ` claim of ID tokens. This value is computed using the ` + "`" + `issuer_host` + "`" + ` and ` + "`" + `https_enabled` + "`" + ` fields. ## Import OIDC Providers can be imported using the ` + "`" + `name` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_identity_oidc_provider.test my-provider ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_identity_oidc_role",
			Category:         "Resources",
			ShortDescription: `Creates an Identity OIDC Role for Vault`,
			Description:      ``,
			Keywords: []string{
				"identity",
				"oidc",
				"role",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) The namespace to provision the resource in. The value should not contain leading or trailing forward slashes. The ` + "`" + `namespace` + "`" + ` is always relative to the provider's configured [namespace](/docs/providers/vault#namespace).`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required; Forces new resource) Name of the OIDC Role to create.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required; Forces new resource) A configured named key, the key must already exist before tokens can be issued.`,
				},
				resource.Attribute{
					Name:        "template",
					Description: `(Optional) The template string to use for generating tokens. This may be in string-ified JSON or base64 format. See the [documentation](https://www.vaultproject.io/docs/secrets/identity/index.html#token-contents-and-templates) for the template format.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `(Optional) TTL of the tokens generated against the role in number of seconds.`,
				},
				resource.Attribute{
					Name:        "client_id",
					Description: `(Optional) The value that will be included in the ` + "`" + `aud` + "`" + ` field of all the OIDC identity tokens issued by this role ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The name of the created role. ## Import The key can be imported with the role name, for example: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_identity_oidc_role.role role ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The name of the created role. ## Import The key can be imported with the role name, for example: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_identity_oidc_role.role role ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_identity_oidc_scope",
			Category:         "Resources",
			ShortDescription: `Provision OIDC Scopes in Vault.`,
			Description:      ``,
			Keywords: []string{
				"identity",
				"oidc",
				"scope",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) The namespace to provision the resource in. The value should not contain leading or trailing forward slashes. The ` + "`" + `namespace` + "`" + ` is always relative to the provider's configured [namespace](/docs/providers/vault#namespace).`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the scope. The ` + "`" + `openid` + "`" + ` scope name is reserved.`,
				},
				resource.Attribute{
					Name:        "template",
					Description: `(Optional) The template string for the scope. This may be provided as escaped JSON or base64 encoded JSON.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A description of the scope. ## Attributes Reference No additional attributes are exported by this resource. ## Import OIDC Scopes can be imported using the ` + "`" + `name` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_identity_oidc_scope.groups groups ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_jwt_auth_backend",
			Category:         "Resources",
			ShortDescription: `Managing JWT/OIDC auth backends in Vault`,
			Description:      ``,
			Keywords: []string{
				"jwt",
				"auth",
				"backend",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) The namespace to provision the resource in. The value should not contain leading or trailing forward slashes. The ` + "`" + `namespace` + "`" + ` is always relative to the provider's configured [namespace](/docs/providers/vault#namespace).`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Required) Path to mount the JWT/OIDC auth backend`,
				},
				resource.Attribute{
					Name:        "disable_remount",
					Description: `(Optional) If set, opts out of mount migration on path updates. See here for more info on [Mount Migration](https://www.vaultproject.io/docs/concepts/mount-migration)`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) Type of auth backend. Should be one of ` + "`" + `jwt` + "`" + ` or ` + "`" + `oidc` + "`" + `. Default - ` + "`" + `jwt` + "`" + ``,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description of the auth backend`,
				},
				resource.Attribute{
					Name:        "oidc_discovery_url",
					Description: `(Optional) The OIDC Discovery URL, without any .well-known component (base path). Cannot be used in combination with ` + "`" + `jwt_validation_pubkeys` + "`" + ``,
				},
				resource.Attribute{
					Name:        "oidc_discovery_ca_pem",
					Description: `(Optional) The CA certificate or chain of certificates, in PEM format, to use to validate connections to the OIDC Discovery URL. If not set, system certificates are used`,
				},
				resource.Attribute{
					Name:        "oidc_client_id",
					Description: `(Optional) Client ID used for OIDC backends`,
				},
				resource.Attribute{
					Name:        "oidc_client_secret",
					Description: `(Optional) Client Secret used for OIDC backends`,
				},
				resource.Attribute{
					Name:        "oidc_response_mode",
					Description: `(Optional) The response mode to be used in the OAuth2 request. Allowed values are ` + "`" + `query` + "`" + ` and ` + "`" + `form_post` + "`" + `. Defaults to ` + "`" + `query` + "`" + `. If using Vault namespaces, and ` + "`" + `oidc_response_mode` + "`" + ` is ` + "`" + `form_post` + "`" + `, then ` + "`" + `namespace_in_state` + "`" + ` should be set to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "oidc_response_types",
					Description: `(Optional) List of response types to request. Allowed values are 'code' and 'id_token'. Defaults to ` + "`" + `["code"]` + "`" + `. Note: ` + "`" + `id_token` + "`" + ` may only be used if ` + "`" + `oidc_response_mode` + "`" + ` is set to ` + "`" + `form_post` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "jwks_url",
					Description: `(Optional) JWKS URL to use to authenticate signatures. Cannot be used with "oidc_discovery_url" or "jwt_validation_pubkeys".`,
				},
				resource.Attribute{
					Name:        "jwks_ca_pem",
					Description: `(Optional) The CA certificate or chain of certificates, in PEM format, to use to validate connections to the JWKS URL. If not set, system certificates are used.`,
				},
				resource.Attribute{
					Name:        "jwt_validation_pubkeys",
					Description: `(Optional) A list of PEM-encoded public keys to use to authenticate signatures locally. Cannot be used in combination with ` + "`" + `oidc_discovery_url` + "`" + ``,
				},
				resource.Attribute{
					Name:        "bound_issuer",
					Description: `(Optional) The value against which to match the iss claim in a JWT`,
				},
				resource.Attribute{
					Name:        "jwt_supported_algs",
					Description: `(Optional) A list of supported signing algorithms. Vault 1.1.0 defaults to [RS256] but future or past versions of Vault may differ`,
				},
				resource.Attribute{
					Name:        "default_role",
					Description: `(Optional) The default role to use if none is provided during login`,
				},
				resource.Attribute{
					Name:        "provider_config",
					Description: `(Optional) Provider specific handling configuration. All values may be strings, and the provider will convert to the appropriate type when configuring Vault.`,
				},
				resource.Attribute{
					Name:        "local",
					Description: `(Optional) Specifies if the auth method is local only.`,
				},
				resource.Attribute{
					Name:        "namespace_in_state",
					Description: `(Optional) Pass namespace in the OIDC state parameter instead of as a separate query parameter. With this setting, the allowed redirect URL(s) in Vault and on the provider side should not contain a namespace query parameter. This means only one redirect URL entry needs to be maintained on the OIDC provider side for all vault namespaces that will be authenticating against it. Defaults to true for new configs`,
				},
				resource.Attribute{
					Name:        "default_lease_ttl",
					Description: `(Optional) Specifies the default time-to-live. If set, this overrides the global default. Must be a valid [duration string](https://golang.org/pkg/time/#ParseDuration)`,
				},
				resource.Attribute{
					Name:        "max_lease_ttl",
					Description: `(Optional) Specifies the maximum time-to-live. If set, this overrides the global default. Must be a valid [duration string](https://golang.org/pkg/time/#ParseDuration)`,
				},
				resource.Attribute{
					Name:        "audit_non_hmac_response_keys",
					Description: `(Optional) Specifies the list of keys that will not be HMAC'd by audit devices in the response data object.`,
				},
				resource.Attribute{
					Name:        "audit_non_hmac_request_keys",
					Description: `(Optional) Specifies the list of keys that will not be HMAC'd by audit devices in the request data object.`,
				},
				resource.Attribute{
					Name:        "listing_visibility",
					Description: `(Optional) Specifies whether to show this mount in the UI-specific listing endpoint. Valid values are "unauth" or "hidden".`,
				},
				resource.Attribute{
					Name:        "passthrough_request_headers",
					Description: `(Optional) List of headers to whitelist and pass from the request to the backend.`,
				},
				resource.Attribute{
					Name:        "allowed_response_headers",
					Description: `(Optional) List of headers to whitelist and allowing a plugin to include them in the response.`,
				},
				resource.Attribute{
					Name:        "token_type",
					Description: `(Optional) Specifies the type of tokens that should be returned by the mount. Valid values are "default-service", "default-batch", "service", "batch". ## Attributes Reference In addition to the fields above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "accessor",
					Description: `The accessor for this auth method ## Import JWT auth backend can be imported using the ` + "`" + `type` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_jwt_auth_backend.oidc oidc ` + "`" + `` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_jwt_auth_backend.jwt jwt ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "accessor",
					Description: `The accessor for this auth method ## Import JWT auth backend can be imported using the ` + "`" + `type` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_jwt_auth_backend.oidc oidc ` + "`" + `` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_jwt_auth_backend.jwt jwt ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_jwt_auth_backend_role",
			Category:         "Resources",
			ShortDescription: `Manages JWT/OIDC auth backend roles in Vault.`,
			Description:      ``,
			Keywords: []string{
				"jwt",
				"auth",
				"backend",
				"role",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) The namespace to provision the resource in. The value should not contain leading or trailing forward slashes. The ` + "`" + `namespace` + "`" + ` is always relative to the provider's configured [namespace](/docs/providers/vault#namespace).`,
				},
				resource.Attribute{
					Name:        "role_name",
					Description: `(Required) The name of the role.`,
				},
				resource.Attribute{
					Name:        "role_type",
					Description: `(Optional) Type of role, either "oidc" (default) or "jwt".`,
				},
				resource.Attribute{
					Name:        "bound_audiences",
					Description: `(For "jwt" roles, at least one of ` + "`" + `bound_audiences` + "`" + `, ` + "`" + `bound_subject` + "`" + `, ` + "`" + `bound_claims` + "`" + ` or ` + "`" + `token_bound_cidrs` + "`" + ` is required. Optional for "oidc" roles.) List of ` + "`" + `aud` + "`" + ` claims to match against. Any match is sufficient.`,
				},
				resource.Attribute{
					Name:        "user_claim",
					Description: `(Required) The claim to use to uniquely identify the user; this will be used as the name for the Identity entity alias created due to a successful login.`,
				},
				resource.Attribute{
					Name:        "user_claim_json_pointer",
					Description: `(Optional) Specifies if the ` + "`" + `user_claim` + "`" + ` value uses [JSON pointer](https://www.vaultproject.io/docs/auth/jwt#claim-specifications-and-json-pointer) syntax for referencing claims. By default, the ` + "`" + `user_claim` + "`" + ` value will not use JSON pointer. Requires Vault 1.11+.`,
				},
				resource.Attribute{
					Name:        "bound_subject",
					Description: `(Optional) If set, requires that the ` + "`" + `sub` + "`" + ` claim matches this value.`,
				},
				resource.Attribute{
					Name:        "bound_claims",
					Description: `(Optional) If set, a map of claims to values to match against. A claim's value must be a string, which may contain one value or multiple comma-separated values, e.g. ` + "`" + `"red"` + "`" + ` or ` + "`" + `"red,green,blue"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "bound_claims_type",
					Description: `(Optional) How to interpret values in the claims/values map (` + "`" + `bound_claims` + "`" + `): can be either ` + "`" + `string` + "`" + ` (exact match) or ` + "`" + `glob` + "`" + ` (wildcard match). Requires Vault 1.4.0 or above.`,
				},
				resource.Attribute{
					Name:        "claim_mappings",
					Description: `(Optional) If set, a map of claims (keys) to be copied to specified metadata fields (values).`,
				},
				resource.Attribute{
					Name:        "oidc_scopes",
					Description: `(Optional) If set, a list of OIDC scopes to be used with an OIDC role. The standard scope "openid" is automatically included and need not be specified.`,
				},
				resource.Attribute{
					Name:        "groups_claim",
					Description: `(Optional) The claim to use to uniquely identify the set of groups to which the user belongs; this will be used as the names for the Identity group aliases created due to a successful login. The claim value must be a list of strings.`,
				},
				resource.Attribute{
					Name:        "backend",
					Description: `(Optional) The unique name of the auth backend to configure. Defaults to ` + "`" + `jwt` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "allowed_redirect_uris",
					Description: `(Optional) The list of allowed values for redirect_uri during OIDC logins. Required for OIDC roles`,
				},
				resource.Attribute{
					Name:        "clock_skew_leeway",
					Description: `(Optional) The amount of leeway to add to all claims to account for clock skew, in seconds. Defaults to ` + "`" + `60` + "`" + ` seconds if set to ` + "`" + `0` + "`" + ` and can be disabled if set to ` + "`" + `-1` + "`" + `. Only applicable with "jwt" roles.`,
				},
				resource.Attribute{
					Name:        "expiration_leeway",
					Description: `(Optional) The amount of leeway to add to expiration (` + "`" + `exp` + "`" + `) claims to account for clock skew, in seconds. Defaults to ` + "`" + `60` + "`" + ` seconds if set to ` + "`" + `0` + "`" + ` and can be disabled if set to ` + "`" + `-1` + "`" + `. Only applicable with "jwt" roles.`,
				},
				resource.Attribute{
					Name:        "not_before_leeway",
					Description: `(Optional) The amount of leeway to add to not before (` + "`" + `nbf` + "`" + `) claims to account for clock skew, in seconds. Defaults to ` + "`" + `60` + "`" + ` seconds if set to ` + "`" + `0` + "`" + ` and can be disabled if set to ` + "`" + `-1` + "`" + `. Only applicable with "jwt" roles.`,
				},
				resource.Attribute{
					Name:        "verbose_oidc_logging",
					Description: `(Optional) Log received OIDC tokens and claims when debug-level logging is active. Not recommended in production since sensitive information may be present in OIDC responses.`,
				},
				resource.Attribute{
					Name:        "max_age",
					Description: `(Optional) Specifies the allowable elapsed time in seconds since the last time the user was actively authenticated with the OIDC provider. ### Common Token Arguments These arguments are common across several Authentication Token resources since Vault 1.2.`,
				},
				resource.Attribute{
					Name:        "token_ttl",
					Description: `(Optional) The incremental lifetime for generated tokens in number of seconds. Its current value will be referenced at renewal time.`,
				},
				resource.Attribute{
					Name:        "token_max_ttl",
					Description: `(Optional) The maximum lifetime for generated tokens in number of seconds. Its current value will be referenced at renewal time.`,
				},
				resource.Attribute{
					Name:        "token_period",
					Description: `(Optional) If set, indicates that the token generated using this role should never expire. The token should be renewed within the duration specified by this value. At each renewal, the token's TTL will be set to the value of this field. Specified in seconds.`,
				},
				resource.Attribute{
					Name:        "token_policies",
					Description: `(Optional) List of policies to encode onto generated tokens. Depending on the auth method, this list may be supplemented by user/group/other values.`,
				},
				resource.Attribute{
					Name:        "token_bound_cidrs",
					Description: `(Optional) List of CIDR blocks; if set, specifies blocks of IP addresses which can authenticate successfully, and ties the resulting token to these blocks as well.`,
				},
				resource.Attribute{
					Name:        "token_explicit_max_ttl",
					Description: `(Optional) If set, will encode an [explicit max TTL](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls) onto the token in number of seconds. This is a hard cap even if ` + "`" + `token_ttl` + "`" + ` and ` + "`" + `token_max_ttl` + "`" + ` would otherwise allow a renewal.`,
				},
				resource.Attribute{
					Name:        "token_no_default_policy",
					Description: `(Optional) If set, the default policy will not be set on generated tokens; otherwise it will be added to the policies set in token_policies.`,
				},
				resource.Attribute{
					Name:        "token_num_uses",
					Description: `(Optional) The [maximum number](https://www.vaultproject.io/api-docs/jwt#token_num_uses) of times a generated token may be used (within its lifetime); 0 means unlimited.`,
				},
				resource.Attribute{
					Name:        "token_type",
					Description: `(Optional) The type of token that should be generated. Can be ` + "`" + `service` + "`" + `, ` + "`" + `batch` + "`" + `, or ` + "`" + `default` + "`" + ` to use the mount's tuned default (which unless changed will be ` + "`" + `service` + "`" + ` tokens). For token store roles, there are two additional possibilities: ` + "`" + `default-service` + "`" + ` and ` + "`" + `default-batch` + "`" + ` which specify the type to return unless the client requests a different type at generation time. ## Attributes Reference No additional attributes are exported by this resource. ## Import JWT authentication backend roles can be imported using the ` + "`" + `path` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_jwt_auth_backend_role.example auth/jwt/role/test-role ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_kmip_secret_backend",
			Category:         "Resources",
			ShortDescription: `Provision KMIP Secret backends in Vault.`,
			Description:      ``,
			Keywords: []string{
				"kmip",
				"secret",
				"backend",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) The namespace to provision the resource in. The value should not contain leading or trailing forward slashes. The ` + "`" + `namespace` + "`" + ` is always relative to the provider's configured [namespace](/docs/providers/vault#namespace).`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Required) The unique path this backend should be mounted at. Must not begin or end with a ` + "`" + `/` + "`" + `. Defaults to ` + "`" + `kmip` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "disable_remount",
					Description: `(Optional) If set, opts out of mount migration on path updates. See here for more info on [Mount Migration](https://www.vaultproject.io/docs/concepts/mount-migration)`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A human-friendly description for this backend.`,
				},
				resource.Attribute{
					Name:        "listen_addrs",
					Description: `(Optional) Addresses the KMIP server should listen on (` + "`" + `host:port` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "server_hostnames",
					Description: `(Optional) Hostnames to include in the server's TLS certificate as SAN DNS names. The first will be used as the common name (CN).`,
				},
				resource.Attribute{
					Name:        "server_ips",
					Description: `(Optional) IPs to include in the server's TLS certificate as SAN IP addresses.`,
				},
				resource.Attribute{
					Name:        "tls_ca_key_type",
					Description: `(Optional) CA key type, rsa or ec.`,
				},
				resource.Attribute{
					Name:        "tls_ca_key_bits",
					Description: `(Optional) CA key bits, valid values depend on key type.`,
				},
				resource.Attribute{
					Name:        "tls_min_version",
					Description: `(Optional) Minimum TLS version to accept.`,
				},
				resource.Attribute{
					Name:        "default_tls_client_key_type",
					Description: `(Optional) Client certificate key type, ` + "`" + `rsa` + "`" + ` or ` + "`" + `ec` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "default_tls_client_key_bits",
					Description: `(Optional) Client certificate key bits, valid values depend on key type.`,
				},
				resource.Attribute{
					Name:        "default_tls_client_key_type",
					Description: `(Optional) Client certificate key type, ` + "`" + `rsa` + "`" + ` or ` + "`" + `ec` + "`" + `. ## Attributes Reference No additional attributes are exported by this resource. ## Import KMIP Secret backend can be imported using the ` + "`" + `path` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_kmip_secret_backend.default kmip ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_kmip_secret_role",
			Category:         "Resources",
			ShortDescription: `Provision KMIP Secret roles in Vault.`,
			Description:      ``,
			Keywords: []string{
				"kmip",
				"secret",
				"role",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) The namespace to provision the resource in. The value should not contain leading or trailing forward slashes. The ` + "`" + `namespace` + "`" + ` is always relative to the provider's configured [namespace](/docs/providers/vault#namespace).`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Required) The unique path this backend should be mounted at. Must not begin or end with a ` + "`" + `/` + "`" + `. Defaults to ` + "`" + `kmip` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "scope",
					Description: `(Required) Name of the scope.`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `(Required) Name of the role.`,
				},
				resource.Attribute{
					Name:        "tls_client_key_type",
					Description: `(Optional) Client certificate key type, ` + "`" + `rsa` + "`" + ` or ` + "`" + `ec` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "tls_client_key_bits",
					Description: `(Optional) Client certificate key bits, valid values depend on key type.`,
				},
				resource.Attribute{
					Name:        "tls_client_ttl",
					Description: `(Optional) Client certificate TTL in seconds.`,
				},
				resource.Attribute{
					Name:        "operation_activate",
					Description: `(Optional) Grant permission to use the KMIP Activate operation.`,
				},
				resource.Attribute{
					Name:        "operation_add_attribute",
					Description: `(Optional) Grant permission to use the KMIP Add Attribute operation.`,
				},
				resource.Attribute{
					Name:        "operation_all",
					Description: `(Optional) Grant all permissions to this role. May not be specified with any other ` + "`" + `operation_`,
				},
				resource.Attribute{
					Name:        "operation_create",
					Description: `(Optional) Grant permission to use the KMIP Create operation.`,
				},
				resource.Attribute{
					Name:        "operation_destroy",
					Description: `(Optional) Grant permission to use the KMIP Destroy operation.`,
				},
				resource.Attribute{
					Name:        "operation_discover_versions",
					Description: `(Optional) Grant permission to use the KMIP Discover Version operation.`,
				},
				resource.Attribute{
					Name:        "operation_get",
					Description: `(Optional) Grant permission to use the KMIP Get operation.`,
				},
				resource.Attribute{
					Name:        "operation_get_attribute_list",
					Description: `(Optional) Grant permission to use the KMIP Get Atrribute List operation.`,
				},
				resource.Attribute{
					Name:        "operation_get_attributes",
					Description: `(Optional) Grant permission to use the KMIP Get Atrributes operation.`,
				},
				resource.Attribute{
					Name:        "operation_locate",
					Description: `(Optional) Grant permission to use the KMIP Get Locate operation.`,
				},
				resource.Attribute{
					Name:        "operation_none",
					Description: `(Optional) Remove all permissions from this role. May not be specified with any other ` + "`" + `operation_`,
				},
				resource.Attribute{
					Name:        "operation_register",
					Description: `(Optional) Grant permission to use the KMIP Register operation.`,
				},
				resource.Attribute{
					Name:        "operation_rekey",
					Description: `(Optional) Grant permission to use the KMIP Rekey operation.`,
				},
				resource.Attribute{
					Name:        "operation_revoke",
					Description: `(Optional) Grant permission to use the KMIP Revoke operation. ## Attributes Reference No additional attributes are exported by this resource. ## Import KMIP Secret role can be imported using the ` + "`" + `path` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_kmip_secret_role.admin kmip ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_kmip_secret_scope",
			Category:         "Resources",
			ShortDescription: `Provision KMIP Secret scopes in Vault.`,
			Description:      ``,
			Keywords: []string{
				"kmip",
				"secret",
				"scope",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) The namespace to provision the resource in. The value should not contain leading or trailing forward slashes. The ` + "`" + `namespace` + "`" + ` is always relative to the provider's configured [namespace](/docs/providers/vault#namespace).`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Required) The unique path this backend should be mounted at. Must not begin or end with a ` + "`" + `/` + "`" + `. Defaults to ` + "`" + `kmip` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "scope",
					Description: `(Required) Name of the scope.`,
				},
				resource.Attribute{
					Name:        "force",
					Description: `(Optional) Boolean field to force deletion even if there are managed objects in the scope. ## Attributes Reference No additional attributes are exported by this resource. ## Import KMIP Secret scope can be imported using the ` + "`" + `path` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_kmip_secret_scope.dev kmip ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_kubernetes_auth_backend_config",
			Category:         "Resources",
			ShortDescription: `Manages Kubernetes auth backend configs in Vault.`,
			Description:      ``,
			Keywords: []string{
				"kubernetes",
				"auth",
				"backend",
				"config",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) The namespace to provision the resource in. The value should not contain leading or trailing forward slashes. The ` + "`" + `namespace` + "`" + ` is always relative to the provider's configured [namespace](../index.html#namespace).`,
				},
				resource.Attribute{
					Name:        "kubernetes_host",
					Description: `(Required) Host must be a host string, a host:port pair, or a URL to the base of the Kubernetes API server.`,
				},
				resource.Attribute{
					Name:        "kubernetes_ca_cert",
					Description: `(Optional) PEM encoded CA cert for use by the TLS client used to talk with the Kubernetes API.`,
				},
				resource.Attribute{
					Name:        "token_reviewer_jwt",
					Description: `(Optional) A service account JWT used to access the TokenReview API to validate other JWTs during login. If not set the JWT used for login will be used to access the API.`,
				},
				resource.Attribute{
					Name:        "pem_keys",
					Description: `(Optional) List of PEM-formatted public keys or certificates used to verify the signatures of Kubernetes service account JWTs. If a certificate is given, its public key will be extracted. Not every installation of Kubernetes exposes these keys.`,
				},
				resource.Attribute{
					Name:        "issuer",
					Description: `(Optional) JWT issuer. If no issuer is specified, ` + "`" + `kubernetes.io/serviceaccount` + "`" + ` will be used as the default issuer.`,
				},
				resource.Attribute{
					Name:        "disable_iss_validation",
					Description: `(Optional) Disable JWT issuer validation. Allows to skip ISS validation. Requires Vault ` + "`" + `v1.5.4+` + "`" + ` or Vault auth kubernetes plugin ` + "`" + `v0.7.1+` + "`" + ``,
				},
				resource.Attribute{
					Name:        "disable_local_ca_jwt",
					Description: `(Optional) Disable defaulting to the local CA cert and service account JWT when running in a Kubernetes pod. Requires Vault ` + "`" + `v1.5.4+` + "`" + ` or Vault auth kubernetes plugin ` + "`" + `v0.7.1+` + "`" + ` ## Attributes Reference No additional attributes are exported by this resource. ## Import Kubernetes authentication backend can be imported using the ` + "`" + `path` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_kubernetes_auth_backend_config.config auth/kubernetes/config ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_kubernetes_auth_backend_role",
			Category:         "Resources",
			ShortDescription: `Manages Kubernetes auth backend roles in Vault.`,
			Description:      ``,
			Keywords: []string{
				"kubernetes",
				"auth",
				"backend",
				"role",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) The namespace to provision the resource in. The value should not contain leading or trailing forward slashes. The ` + "`" + `namespace` + "`" + ` is always relative to the provider's configured [namespace](/docs/providers/vault#namespace).`,
				},
				resource.Attribute{
					Name:        "role_name",
					Description: `(Required) Name of the role.`,
				},
				resource.Attribute{
					Name:        "bound_service_account_names",
					Description: `(Required) List of service account names able to access this role. If set to ` + "`" + `["`,
				},
				resource.Attribute{
					Name:        "bound_service_account_namespaces",
					Description: `(Required) List of namespaces allowed to access this role. If set to ` + "`" + `["`,
				},
				resource.Attribute{
					Name:        "backend",
					Description: `(Optional) Unique name of the kubernetes backend to configure.`,
				},
				resource.Attribute{
					Name:        "audience",
					Description: `(Optional) Audience claim to verify in the JWT. ~> Please see [alias_name_source](https://www.vaultproject.io/api-docs/auth/kubernetes#alias_name_source) before setting this to something other its default value. There are`,
				},
				resource.Attribute{
					Name:        "alias_name_source",
					Description: `(Optional, default: ` + "`" + `serviceaccount_uid` + "`" + `) Configures how identity aliases are generated. Valid choices are: ` + "`" + `serviceaccount_uid` + "`" + `, ` + "`" + `serviceaccount_name` + "`" + `. (vault-1.9+) ### Common Token Arguments These arguments are common across several Authentication Token resources since Vault 1.2. Its current value will be referenced at renewal time.`,
				},
				resource.Attribute{
					Name:        "token_max_ttl",
					Description: `(Optional) The maximum lifetime for generated tokens in number of seconds. Its current value will be referenced at renewal time.`,
				},
				resource.Attribute{
					Name:        "token_period",
					Description: `(Optional) If set, indicates that the token generated using this role should never expire. The token should be renewed within the duration specified by this value. At each renewal, the token's TTL will be set to the value of this field. Specified in seconds.`,
				},
				resource.Attribute{
					Name:        "token_policies",
					Description: `(Optional) List of policies to encode onto generated tokens. Depending on the auth method, this list may be supplemented by user/group/other values.`,
				},
				resource.Attribute{
					Name:        "token_bound_cidrs",
					Description: `(Optional) List of CIDR blocks; if set, specifies blocks of IP addresses which can authenticate successfully, and ties the resulting token to these blocks as well.`,
				},
				resource.Attribute{
					Name:        "token_explicit_max_ttl",
					Description: `(Optional) If set, will encode an [explicit max TTL](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls) onto the token in number of seconds. This is a hard cap even if ` + "`" + `token_ttl` + "`" + ` and ` + "`" + `token_max_ttl` + "`" + ` would otherwise allow a renewal.`,
				},
				resource.Attribute{
					Name:        "token_no_default_policy",
					Description: `(Optional) If set, the default policy will not be set on generated tokens; otherwise it will be added to the policies set in token_policies.`,
				},
				resource.Attribute{
					Name:        "token_num_uses",
					Description: `(Optional) The [maximum number](https://www.vaultproject.io/api-docs/kubernetes#token_num_uses) of times a generated token may be used (within its lifetime); 0 means unlimited.`,
				},
				resource.Attribute{
					Name:        "token_type",
					Description: `(Optional) The type of token that should be generated. Can be ` + "`" + `service` + "`" + `, ` + "`" + `batch` + "`" + `, or ` + "`" + `default` + "`" + ` to use the mount's tuned default (which unless changed will be ` + "`" + `service` + "`" + ` tokens). For token store roles, there are two additional possibilities: ` + "`" + `default-service` + "`" + ` and ` + "`" + `default-batch` + "`" + ` which specify the type to return unless the client requests a different type at generation time. ## Attributes Reference No additional attributes are exported by this resource. ## Import Kubernetes auth backend role can be imported using the ` + "`" + `path` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_kubernetes_auth_backend_role.foo auth/kubernetes/role/foo ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_kubernetes_secret_backend",
			Category:         "Resources",
			ShortDescription: `Creates a Kubernetes Secrets Engine in Vault.`,
			Description:      ``,
			Keywords: []string{
				"kubernetes",
				"secret",
				"backend",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) The namespace to provision the resource in. The value should not contain leading or trailing forward slashes. The ` + "`" + `namespace` + "`" + ` is always relative to the provider's configured [namespace](/docs/providers/vault#namespace).`,
				},
				resource.Attribute{
					Name:        "kubernetes_host",
					Description: `(Optional) The Kubernetes API URL to connect to. Required if the standard pod environment variables ` + "`" + `KUBERNETES_SERVICE_HOST` + "`" + ` or ` + "`" + `KUBERNETES_SERVICE_PORT` + "`" + ` are not set on the host that Vault is running on.`,
				},
				resource.Attribute{
					Name:        "kubernetes_ca_cert",
					Description: `(Optional) A PEM-encoded CA certificate used by the secrets engine to verify the Kubernetes API server certificate. Defaults to the local pod’s CA if Vault is running in Kubernetes. Otherwise, defaults to the root CA set where Vault is running.`,
				},
				resource.Attribute{
					Name:        "service_account_jwt",
					Description: `(Optional) The JSON web token of the service account used by the secrets engine to manage Kubernetes credentials. Defaults to the local pod’s JWT if Vault is running in Kubernetes.`,
				},
				resource.Attribute{
					Name:        "disable_local_ca_jwt",
					Description: `(Optional) Disable defaulting to the local CA certificate and service account JWT when Vault is running in a Kubernetes pod. ## Attributes Reference No additional attributes are exported by this resource. ## Import The Kubernetes secret backend can be imported using its ` + "`" + `path` + "`" + ` e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_kubernetes_secret_backend.config kubernetes ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_kubernetes_secret_backend_role",
			Category:         "Resources",
			ShortDescription: `Creates a role for the Kubernetes Secrets Engine in Vault.`,
			Description:      ``,
			Keywords: []string{
				"kubernetes",
				"secret",
				"backend",
				"role",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) The namespace to provision the resource in. The value should not contain leading or trailing forward slashes. The ` + "`" + `namespace` + "`" + ` is always relative to the provider's configured [namespace](/docs/providers/vault#namespace).`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the role.`,
				},
				resource.Attribute{
					Name:        "backend",
					Description: `(Required) The path of the Kubernetes Secrets Engine backend mount to create the role in.`,
				},
				resource.Attribute{
					Name:        "allowed_kubernetes_namespaces",
					Description: `(Required) The list of Kubernetes namespaces this role can generate credentials for. If set to ` + "`" + ``,
				},
				resource.Attribute{
					Name:        "token_max_ttl",
					Description: `(Optional) The maximum TTL for generated Kubernetes tokens in seconds.`,
				},
				resource.Attribute{
					Name:        "token_default_ttl",
					Description: `(Optional) The default TTL for generated Kubernetes tokens in seconds.`,
				},
				resource.Attribute{
					Name:        "service_account_name",
					Description: `(Optional) The pre-existing service account to generate tokens for. Mutually exclusive with ` + "`" + `kubernetes_role_name` + "`" + ` and ` + "`" + `generated_role_rules` + "`" + `. If set, only a Kubernetes token will be created when credentials are requested.`,
				},
				resource.Attribute{
					Name:        "kubernetes_role_name",
					Description: `(Optional) The pre-existing Role or ClusterRole to bind a generated service account to. Mutually exclusive with ` + "`" + `service_account_name` + "`" + ` and ` + "`" + `generated_role_rules` + "`" + `. If set, Kubernetes token, service account, and role binding objects will be created when credentials are requested.`,
				},
				resource.Attribute{
					Name:        "kubernetes_role_type",
					Description: `(Optional) Specifies whether the Kubernetes role is a Role or ClusterRole.`,
				},
				resource.Attribute{
					Name:        "generated_role_rules",
					Description: `(Optional) The Role or ClusterRole rules to use when generating a role. Accepts either JSON or YAML formatted rules. Mutually exclusive with ` + "`" + `service_account_name` + "`" + ` and ` + "`" + `kubernetes_role_name` + "`" + `. If set, the entire chain of Kubernetes objects will be generated when credentials are requested.`,
				},
				resource.Attribute{
					Name:        "name_template",
					Description: `(Optional) The name template to use when generating service accounts, roles and role bindings. If unset, a default template is used.`,
				},
				resource.Attribute{
					Name:        "extra_annotations",
					Description: `(Optional) Additional annotations to apply to all generated Kubernetes objects.`,
				},
				resource.Attribute{
					Name:        "extra_labels",
					Description: `(Optional) Additional labels to apply to all generated Kubernetes objects. This resource also directly accepts all [vault_mount](mount.html.md) fields. ## Attributes Reference No additional attributes are exported by this resource. ## Import The Kubernetes secret backend role can be imported using the full path to the role of the form: ` + "`" + `<backend_path>/roles/<role_name>` + "`" + ` e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_kubernetes_secret_backend_role.example kubernetes kubernetes/roles/example-role ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_kv_secret",
			Category:         "Resources",
			ShortDescription: `Writes a KV-V1 secret to a given path in Vault`,
			Description:      ``,
			Keywords: []string{
				"kv",
				"secret",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) The namespace to provision the resource in. The value should not contain leading or trailing forward slashes. The ` + "`" + `namespace` + "`" + ` is always relative to the provider's configured [namespace](/docs/providers/vault#namespace).`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Required) Full path of the KV-V1 secret.`,
				},
				resource.Attribute{
					Name:        "data_json",
					Description: `(Required) JSON-encoded string that will be written as the secret data at the given path. ## Required Vault Capabilities Use of this resource requires the ` + "`" + `create` + "`" + ` or ` + "`" + `update` + "`" + ` capability (depending on whether the resource already exists) on the given path, the ` + "`" + `delete` + "`" + ` capability if the resource is removed from configuration, and the ` + "`" + `read` + "`" + ` capability for drift detection (by default). ## Attributes Reference The following attributes are exported in addition to the above:`,
				},
				resource.Attribute{
					Name:        "data",
					Description: `A mapping whose keys are the top-level data keys returned from Vault and whose values are the corresponding values. This map can only represent string data, so any non-string values returned from Vault are serialized as JSON. ## Import KV-V1 secrets can be imported using the ` + "`" + `path` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_kv_secret.secret kvv1/secret ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "data",
					Description: `A mapping whose keys are the top-level data keys returned from Vault and whose values are the corresponding values. This map can only represent string data, so any non-string values returned from Vault are serialized as JSON. ## Import KV-V1 secrets can be imported using the ` + "`" + `path` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_kv_secret.secret kvv1/secret ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_kv_secret_backend_v2",
			Category:         "Resources",
			ShortDescription: `Configures KV-V2 backend level settings that are applied to every key in the key-value store.`,
			Description:      ``,
			Keywords: []string{
				"kv",
				"secret",
				"backend",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) The namespace to provision the resource in. The value should not contain leading or trailing forward slashes. The ` + "`" + `namespace` + "`" + ` is always relative to the provider's configured [namespace](/docs/providers/vault#namespace).`,
				},
				resource.Attribute{
					Name:        "mount",
					Description: `(Required) Path where KV-V2 engine is mounted.`,
				},
				resource.Attribute{
					Name:        "max_versions",
					Description: `(Optional) The number of versions to keep per key.`,
				},
				resource.Attribute{
					Name:        "cas_required",
					Description: `(Optional) If true, all keys will require the cas parameter to be set on all write requests.`,
				},
				resource.Attribute{
					Name:        "delete_version_after",
					Description: `(Optional) If set, specifies the length of time before a version is deleted. Accepts duration in integer seconds. ## Required Vault Capabilities Use of this resource requires the ` + "`" + `create` + "`" + ` or ` + "`" + `update` + "`" + ` capability (depending on whether the resource already exists) on the given path, the ` + "`" + `delete` + "`" + ` capability if the resource is removed from configuration, and the ` + "`" + `read` + "`" + ` capability for drift detection (by default). ## Attributes Reference No additional attributes are exported by this resource. ## Import The KV-V2 secret backend can be imported using its unique ID, the ` + "`" + `${mount}/config` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_kv_secret_backend_v2.example kvv2/config ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_kv_secret_v2",
			Category:         "Resources",
			ShortDescription: `Writes a KV-V2 secret to a given path in Vault`,
			Description:      ``,
			Keywords: []string{
				"kv",
				"secret",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) The namespace to provision the resource in. The value should not contain leading or trailing forward slashes. The ` + "`" + `namespace` + "`" + ` is always relative to the provider's configured [namespace](/docs/providers/vault#namespace).`,
				},
				resource.Attribute{
					Name:        "mount",
					Description: `(Required) Path where KV-V2 engine is mounted.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Full name of the secret. For a nested secret the name is the nested path excluding the mount and data prefix. For example, for a secret at ` + "`" + `kvv2/data/foo/bar/baz` + "`" + ` the name is ` + "`" + `foo/bar/baz` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "cas",
					Description: `(Optional) This flag is required if ` + "`" + `cas_required` + "`" + ` is set to true on either the secret or the engine's config. In order for a write operation to be successful, cas must be set to the current version of the secret.`,
				},
				resource.Attribute{
					Name:        "options",
					Description: `(Optional) An object that holds option settings.`,
				},
				resource.Attribute{
					Name:        "disable_read",
					Description: `(Optional) If set to true, disables reading secret from Vault; note: drift won't be detected.`,
				},
				resource.Attribute{
					Name:        "delete_all_versions",
					Description: `(Optional) If set to true, permanently deletes all versions for the specified key.`,
				},
				resource.Attribute{
					Name:        "data_json",
					Description: `(Required) JSON-encoded string that will be written as the secret data at the given path.`,
				},
				resource.Attribute{
					Name:        "custom_metadata",
					Description: `(Optional) A nested block that allows configuring metadata for the KV secret. Refer to the [Configuration Options](#custom-metadata-configuration-options) for more info. ## Required Vault Capabilities Use of this resource requires the ` + "`" + `create` + "`" + ` or ` + "`" + `update` + "`" + ` capability (depending on whether the resource already exists) on the given path, the ` + "`" + `delete` + "`" + ` capability if the resource is removed from configuration, and the ` + "`" + `read` + "`" + ` capability for drift detection (by default). ### Custom Metadata Configuration Options`,
				},
				resource.Attribute{
					Name:        "max_versions",
					Description: `(Optional) The number of versions to keep per key.`,
				},
				resource.Attribute{
					Name:        "cas_required",
					Description: `(Optional) If true, all keys will require the cas parameter to be set on all write requests.`,
				},
				resource.Attribute{
					Name:        "delete_version_after",
					Description: `(Optional) If set, specifies the length of time before a version is deleted. Accepts duration in integer seconds.`,
				},
				resource.Attribute{
					Name:        "data",
					Description: `(Optional) A string to string map describing the secret. ## Attributes Reference The following attributes are exported in addition to the above:`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `Full path where the KV-V2 secret will be written.`,
				},
				resource.Attribute{
					Name:        "data",
					Description: `A mapping whose keys are the top-level data keys returned from Vault and whose values are the corresponding values. This map can only represent string data, so any non-string values returned from Vault are serialized as JSON.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `Metadata associated with this secret read from Vault. ## Import KV-V2 secrets can be imported using the ` + "`" + `path` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_kv_secret_v2.example kvv2/data/secret ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "path",
					Description: `Full path where the KV-V2 secret will be written.`,
				},
				resource.Attribute{
					Name:        "data",
					Description: `A mapping whose keys are the top-level data keys returned from Vault and whose values are the corresponding values. This map can only represent string data, so any non-string values returned from Vault are serialized as JSON.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `Metadata associated with this secret read from Vault. ## Import KV-V2 secrets can be imported using the ` + "`" + `path` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_kv_secret_v2.example kvv2/data/secret ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_ldap_auth_backend",
			Category:         "Resources",
			ShortDescription: `Managing LDAP auth backends in Vault`,
			Description:      ``,
			Keywords: []string{
				"ldap",
				"auth",
				"backend",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) The namespace to provision the resource in. The value should not contain leading or trailing forward slashes. The ` + "`" + `namespace` + "`" + ` is always relative to the provider's configured [namespace](/docs/providers/vault#namespace).`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Required) The URL of the LDAP server`,
				},
				resource.Attribute{
					Name:        "starttls",
					Description: `(Optional) Control use of TLS when conecting to LDAP`,
				},
				resource.Attribute{
					Name:        "case_sensitive_names",
					Description: `(Optional) Control case senstivity of objects fetched from LDAP, this is used for object matching in vault`,
				},
				resource.Attribute{
					Name:        "tls_min_version",
					Description: `(Optional) Minimum acceptable version of TLS`,
				},
				resource.Attribute{
					Name:        "tls_max_version",
					Description: `(Optional) Maximum acceptable version of TLS`,
				},
				resource.Attribute{
					Name:        "insecure_tls",
					Description: `(Optional) Control whether or TLS certificates must be validated`,
				},
				resource.Attribute{
					Name:        "certificate",
					Description: `(Optional) Trusted CA to validate TLS certificate`,
				},
				resource.Attribute{
					Name:        "binddn",
					Description: `(Optional) DN of object to bind when performing user search`,
				},
				resource.Attribute{
					Name:        "bindpass",
					Description: `(Optional) Password to use with ` + "`" + `binddn` + "`" + ` when performing user search`,
				},
				resource.Attribute{
					Name:        "userdn",
					Description: `(Optional) Base DN under which to perform user search`,
				},
				resource.Attribute{
					Name:        "userattr",
					Description: `(Optional) Attribute on user object matching username passed in`,
				},
				resource.Attribute{
					Name:        "userfilter",
					Description: `(Optional) LDAP user search filter`,
				},
				resource.Attribute{
					Name:        "upndomain",
					Description: `(Optional) The userPrincipalDomain used to construct UPN string`,
				},
				resource.Attribute{
					Name:        "groupfilter",
					Description: `(Optional) Go template used to construct group membership query`,
				},
				resource.Attribute{
					Name:        "groupdn",
					Description: `(Optional) Base DN under which to perform group search`,
				},
				resource.Attribute{
					Name:        "groupattr",
					Description: `(Optional) LDAP attribute to follow on objects returned by groupfilter`,
				},
				resource.Attribute{
					Name:        "username_as_alias",
					Description: `(Optional) Force the auth method to use the username passed by the user as the alias name.`,
				},
				resource.Attribute{
					Name:        "use_token_groups",
					Description: `(Optional) Use the Active Directory tokenGroups constructed attribute of the user to find the group memberships`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Optional) Path to mount the LDAP auth backend under`,
				},
				resource.Attribute{
					Name:        "disable_remount",
					Description: `(Optional) If set, opts out of mount migration on path updates. See here for more info on [Mount Migration](https://www.vaultproject.io/docs/concepts/mount-migration)`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description for the LDAP auth backend mount`,
				},
				resource.Attribute{
					Name:        "local",
					Description: `(Optional) Specifies if the auth method is local only. ### Common Token Arguments These arguments are common across several Authentication Token resources since Vault 1.2.`,
				},
				resource.Attribute{
					Name:        "token_ttl",
					Description: `(Optional) The incremental lifetime for generated tokens in number of seconds. Its current value will be referenced at renewal time.`,
				},
				resource.Attribute{
					Name:        "token_max_ttl",
					Description: `(Optional) The maximum lifetime for generated tokens in number of seconds. Its current value will be referenced at renewal time.`,
				},
				resource.Attribute{
					Name:        "token_period",
					Description: `(Optional) If set, indicates that the token generated using this role should never expire. The token should be renewed within the duration specified by this value. At each renewal, the token's TTL will be set to the value of this field. Specified in seconds.`,
				},
				resource.Attribute{
					Name:        "token_policies",
					Description: `(Optional) List of policies to encode onto generated tokens. Depending on the auth method, this list may be supplemented by user/group/other values.`,
				},
				resource.Attribute{
					Name:        "token_bound_cidrs",
					Description: `(Optional) List of CIDR blocks; if set, specifies blocks of IP addresses which can authenticate successfully, and ties the resulting token to these blocks as well.`,
				},
				resource.Attribute{
					Name:        "token_explicit_max_ttl",
					Description: `(Optional) If set, will encode an [explicit max TTL](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls) onto the token in number of seconds. This is a hard cap even if ` + "`" + `token_ttl` + "`" + ` and ` + "`" + `token_max_ttl` + "`" + ` would otherwise allow a renewal.`,
				},
				resource.Attribute{
					Name:        "token_no_default_policy",
					Description: `(Optional) If set, the default policy will not be set on generated tokens; otherwise it will be added to the policies set in token_policies.`,
				},
				resource.Attribute{
					Name:        "token_num_uses",
					Description: `(Optional) The [maximum number](https://www.vaultproject.io/api-docs/ldap#token_num_uses) of times a generated token may be used (within its lifetime); 0 means unlimited.`,
				},
				resource.Attribute{
					Name:        "token_type",
					Description: `(Optional) The type of token that should be generated. Can be ` + "`" + `service` + "`" + `, ` + "`" + `batch` + "`" + `, or ` + "`" + `default` + "`" + ` to use the mount's tuned default (which unless changed will be ` + "`" + `service` + "`" + ` tokens). For token store roles, there are two additional possibilities: ` + "`" + `default-service` + "`" + ` and ` + "`" + `default-batch` + "`" + ` which specify the type to return unless the client requests a different type at generation time. For more details on the usage of each argument consult the [Vault LDAP API documentation](https://www.vaultproject.io/api-docs/auth/ldap). ~>`,
				},
				resource.Attribute{
					Name:        "accessor",
					Description: `The accessor for this auth mount. ## Import LDAP authentication backends can be imported using the ` + "`" + `path` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_ldap_auth_backend.ldap ldap ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "accessor",
					Description: `The accessor for this auth mount. ## Import LDAP authentication backends can be imported using the ` + "`" + `path` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_ldap_auth_backend.ldap ldap ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_ldap_auth_backend_group",
			Category:         "Resources",
			ShortDescription: `Managing groups in an LDAP auth backend in Vault`,
			Description:      ``,
			Keywords: []string{
				"ldap",
				"auth",
				"backend",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) The namespace to provision the resource in. The value should not contain leading or trailing forward slashes. The ` + "`" + `namespace` + "`" + ` is always relative to the provider's configured [namespace](/docs/providers/vault#namespace).`,
				},
				resource.Attribute{
					Name:        "groupname",
					Description: `(Required) The LDAP groupname`,
				},
				resource.Attribute{
					Name:        "policies",
					Description: `(Optional) Policies which should be granted to members of the group`,
				},
				resource.Attribute{
					Name:        "backend",
					Description: `(Optional) Path to the authentication backend For more details on the usage of each argument consult the [Vault LDAP API documentation](https://www.vaultproject.io/api-docs/auth/ldap). ## Attribute Reference No additional attributes are exposed by this resource. ## Import LDAP authentication backend groups can be imported using the ` + "`" + `path` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_ldap_auth_backend_group.foo auth/ldap/groups/foo ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_ldap_auth_backend_user",
			Category:         "Resources",
			ShortDescription: `Managing users in an LDAP auth backend in Vault`,
			Description:      ``,
			Keywords: []string{
				"ldap",
				"auth",
				"backend",
				"user",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) The namespace to provision the resource in. The value should not contain leading or trailing forward slashes. The ` + "`" + `namespace` + "`" + ` is always relative to the provider's configured [namespace](/docs/providers/vault#namespace).`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Required) The LDAP username`,
				},
				resource.Attribute{
					Name:        "policies",
					Description: `(Optional) Policies which should be granted to user`,
				},
				resource.Attribute{
					Name:        "groups",
					Description: `(Optional) Override LDAP groups which should be granted to user`,
				},
				resource.Attribute{
					Name:        "backend",
					Description: `(Optional) Path to the authentication backend For more details on the usage of each argument consult the [Vault LDAP API documentation](https://www.vaultproject.io/api-docs/auth/ldap). ## Attribute Reference No additional attributes are exposed by this resource. ## Import LDAP authentication backend users can be imported using the ` + "`" + `path` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_ldap_auth_backend_user.foo auth/ldap/users/foo ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_managed_keys",
			Category:         "Resources",
			ShortDescription: `Configures Managed Keys in Vault`,
			Description:      ``,
			Keywords: []string{
				"managed",
				"keys",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) The namespace to provision the resource in. The value should not contain leading or trailing forward slashes. The ` + "`" + `namespace` + "`" + ` is always relative to the provider's configured [namespace](../index.html#namespace).`,
				},
				resource.Attribute{
					Name:        "allow_generate_key",
					Description: `(Optional) If no existing key can be found in the referenced backend, instructs Vault to generate a key within the backend.`,
				},
				resource.Attribute{
					Name:        "allow_replace_key",
					Description: `(Optional) Controls the ability for Vault to replace through generation or importing a key into the configured backend even if a key is present, if set to ` + "`" + `false` + "`" + ` those operations are forbidden if a key exists.`,
				},
				resource.Attribute{
					Name:        "allow_store_key",
					Description: `(Optional) Controls the ability for Vault to import a key to the configured backend, if ` + "`" + `false` + "`" + `, those operations will be forbidden.`,
				},
				resource.Attribute{
					Name:        "any_mount",
					Description: `(Optional) If ` + "`" + `true` + "`" + `, allows usage from any mount point within the namespace. ### AWS Parameters`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A unique lowercase name that serves as identifying the key.`,
				},
				resource.Attribute{
					Name:        "access_key",
					Description: `(Required) The AWS access key to use.`,
				},
				resource.Attribute{
					Name:        "secret_key",
					Description: `(Required) The AWS access key to use.`,
				},
				resource.Attribute{
					Name:        "key_bits",
					Description: `(Required) The size in bits for an RSA key.`,
				},
				resource.Attribute{
					Name:        "key_type",
					Description: `(Required) The type of key to use.`,
				},
				resource.Attribute{
					Name:        "kms_key",
					Description: `(Required) An identifier for the key.`,
				},
				resource.Attribute{
					Name:        "curve",
					Description: `(Optional) The curve to use for an ECDSA key. Used when ` + "`" + `key_type` + "`" + ` is ` + "`" + `ECDSA` + "`" + `. Required if ` + "`" + `allow_generate_key` + "`" + ` is ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "endpoint",
					Description: `(Optional) Used to specify a custom AWS endpoint.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The AWS region where the keys are stored (or will be stored). ### Azure Parameters`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A unique lowercase name that serves as identifying the key.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Required) The tenant id for the Azure Active Directory organization.`,
				},
				resource.Attribute{
					Name:        "client_id",
					Description: `(Required) The client id for credentials to query the Azure APIs.`,
				},
				resource.Attribute{
					Name:        "client_secret",
					Description: `(Required) The client secret for credentials to query the Azure APIs.`,
				},
				resource.Attribute{
					Name:        "vault_name",
					Description: `(Required) The Key Vault vault to use for encryption and decryption.`,
				},
				resource.Attribute{
					Name:        "key_name",
					Description: `(Required) The Key Vault key to use for encryption and decryption.`,
				},
				resource.Attribute{
					Name:        "key_type",
					Description: `(Required) The type of key to use.`,
				},
				resource.Attribute{
					Name:        "environment",
					Description: `(Optional) The Azure Cloud environment API endpoints to use.`,
				},
				resource.Attribute{
					Name:        "resource",
					Description: `(Optional) The Azure Key Vault resource's DNS Suffix to connect to.`,
				},
				resource.Attribute{
					Name:        "key_bits",
					Description: `(Optional) The size in bits for an RSA key. This field is required when ` + "`" + `key_type` + "`" + ` is ` + "`" + `RSA` + "`" + ` or when ` + "`" + `allow_generate_key` + "`" + ` is ` + "`" + `true` + "`" + ` ### PKCS Parameters`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A unique lowercase name that serves as identifying the key.`,
				},
				resource.Attribute{
					Name:        "library",
					Description: `(Required) The name of the kms_library stanza to use from Vault's config to lookup the local library path.`,
				},
				resource.Attribute{
					Name:        "key_label",
					Description: `(Required) The label of the key to use.`,
				},
				resource.Attribute{
					Name:        "key_id",
					Description: `(Required) The id of a PKCS#11 key to use.`,
				},
				resource.Attribute{
					Name:        "mechanism",
					Description: `(Required) The encryption/decryption mechanism to use, specified as a hexadecimal (prefixed by 0x) string.`,
				},
				resource.Attribute{
					Name:        "pin",
					Description: `(Required) The PIN for login.`,
				},
				resource.Attribute{
					Name:        "slot",
					Description: `(Optional) The slot number to use, specified as a string in a decimal format (e.g. ` + "`" + `2305843009213693953` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "token_label",
					Description: `(Optional) The slot token label to use.`,
				},
				resource.Attribute{
					Name:        "curve",
					Description: `(Optional) Supplies the curve value when using the ` + "`" + `CKM_ECDSA` + "`" + ` mechanism. Required if ` + "`" + `allow_generate_key` + "`" + ` is ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "key_bits",
					Description: `(Optional) Supplies the size in bits of the key when using ` + "`" + `CKM_RSA_PKCS_PSS` + "`" + `, ` + "`" + `CKM_RSA_PKCS_OAEP` + "`" + ` or ` + "`" + `CKM_RSA_PKCS` + "`" + ` as a value for ` + "`" + `mechanism` + "`" + `. Required if ` + "`" + `allow_generate_key` + "`" + ` is ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "force_rw_session",
					Description: `(Optional) Force all operations to open up a read-write session to the HSM. ## Import Mounts can be imported using the ` + "`" + `id` + "`" + ` of ` + "`" + `default` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_managed_keys.keys default ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_mfa_duo",
			Category:         "Resources",
			ShortDescription: `Managing the MFA Duo method configuration`,
			Description:      ``,
			Keywords: []string{
				"mfa",
				"duo",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) The namespace to provision the resource in. The value should not contain leading or trailing forward slashes. The ` + "`" + `namespace` + "`" + ` is always relative to the provider's configured [namespace](/docs/providers/vault#namespace).`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_mfa_okta",
			Category:         "Resources",
			ShortDescription: `Managing the MFA Okta method configuration`,
			Description:      ``,
			Keywords: []string{
				"mfa",
				"okta",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) The namespace to provision the resource in. The value should not contain leading or trailing forward slashes. The ` + "`" + `namespace` + "`" + ` is always relative to the provider's configured [namespace](/docs/providers/vault#namespace).`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_mfa_pingid",
			Category:         "Resources",
			ShortDescription: `Managing the MFA PingID method configuration`,
			Description:      ``,
			Keywords: []string{
				"mfa",
				"pingid",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) The namespace to provision the resource in. The value should not contain leading or trailing forward slashes. The ` + "`" + `namespace` + "`" + ` is always relative to the provider's configured [namespace](/docs/providers/vault#namespace).`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_mfa_totp",
			Category:         "Resources",
			ShortDescription: `Managing the MFA TOTP method configuration`,
			Description:      ``,
			Keywords: []string{
				"mfa",
				"totp",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) The namespace to provision the resource in. The value should not contain leading or trailing forward slashes. The ` + "`" + `namespace` + "`" + ` is always relative to the provider's configured [namespace](/docs/providers/vault#namespace).`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_mongodbatlas_secret_backend",
			Category:         "Resources",
			ShortDescription: `Creates a MongoDB Atlas secret backend for Vault.`,
			Description:      ``,
			Keywords: []string{
				"mongodbatlas",
				"secret",
				"backend",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) The namespace to provision the resource in. The value should not contain leading or trailing forward slashes. The ` + "`" + `namespace` + "`" + ` is always relative to the provider's configured [namespace](/docs/providers/vault#namespace).`,
				},
				resource.Attribute{
					Name:        "mount",
					Description: `(Required) Path where the MongoDB Atlas Secrets Engine is mounted.`,
				},
				resource.Attribute{
					Name:        "private_key",
					Description: `(Required) Specifies the Private API Key used to authenticate with the MongoDB Atlas API.`,
				},
				resource.Attribute{
					Name:        "public_key",
					Description: `(Required) Specifies the Public API Key used to authenticate with the MongoDB Atlas API. ## Attributes Reference No additional attributes are exported by this resource. ## Import MongoDB Atlas secret backends can be imported using the ` + "`" + `${mount}/config` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_mongodbatlas_secret_backend.config mongodbatlas/config ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_mongodbatlas_secret_role",
			Category:         "Resources",
			ShortDescription: `Creates a role for the MongoDB Atlas Secret Engine in Vault.`,
			Description:      ``,
			Keywords: []string{
				"mongodbatlas",
				"secret",
				"role",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) The namespace to provision the resource in. The value should not contain leading or trailing forward slashes. The ` + "`" + `namespace` + "`" + ` is always relative to the provider's configured [namespace](/docs/providers/vault#namespace).`,
				},
				resource.Attribute{
					Name:        "mount",
					Description: `(Required) Path where the MongoDB Atlas Secrets Engine is mounted.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the role.`,
				},
				resource.Attribute{
					Name:        "organization_id",
					Description: `(Optional) Unique identifier for the organization to which the target API Key belongs. Required if ` + "`" + `project_id` + "`" + ` is not set.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Optional) Unique identifier for the project to which the target API Key belongs. Required if ` + "`" + `organization_id is` + "`" + ` not set.`,
				},
				resource.Attribute{
					Name:        "roles",
					Description: `(Required) List of roles that the API Key needs to have.`,
				},
				resource.Attribute{
					Name:        "ip_addresses",
					Description: `(Optional) IP address to be added to the whitelist for the API key.`,
				},
				resource.Attribute{
					Name:        "cidr_blocks",
					Description: `(Optional) Whitelist entry in CIDR notation to be added for the API key.`,
				},
				resource.Attribute{
					Name:        "project_roles",
					Description: `(Optional) Roles assigned when an org API key is assigned to a project API key.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `(Optional) Duration in seconds after which the issued credential should expire.`,
				},
				resource.Attribute{
					Name:        "max_ttl",
					Description: `(Optional) The maximum allowed lifetime of credentials issued using this role. ## Attributes Reference No additional attributes are exported by this resource. ## Import The MongoDB Atlas secret role can be imported using the full path to the role of the form: ` + "`" + `<mount_path>/roles/<role_name>` + "`" + ` e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_mongodbatlas_secret_role.example mongodbatlas/roles/example-role ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_mount",
			Category:         "Resources",
			ShortDescription: `Managing the mounting of secret backends in Vault`,
			Description:      ``,
			Keywords: []string{
				"mount",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) The namespace to provision the resource in. The value should not contain leading or trailing forward slashes. The ` + "`" + `namespace` + "`" + ` is always relative to the provider's configured [namespace](/docs/providers/vault#namespace).`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Required) Where the secret backend will be mounted`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) Type of the backend, such as "aws"`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Human-friendly description of the mount`,
				},
				resource.Attribute{
					Name:        "default_lease_ttl_seconds",
					Description: `(Optional) Default lease duration for tokens and secrets in seconds`,
				},
				resource.Attribute{
					Name:        "max_lease_ttl_seconds",
					Description: `(Optional) Maximum possible lease duration for tokens and secrets in seconds`,
				},
				resource.Attribute{
					Name:        "audit_non_hmac_response_keys",
					Description: `(Optional) Specifies the list of keys that will not be HMAC'd by audit devices in the response data object.`,
				},
				resource.Attribute{
					Name:        "audit_non_hmac_request_keys",
					Description: `(Optional) Specifies the list of keys that will not be HMAC'd by audit devices in the request data object.`,
				},
				resource.Attribute{
					Name:        "local",
					Description: `(Optional) Boolean flag that can be explicitly set to true to enforce local mount in HA environment`,
				},
				resource.Attribute{
					Name:        "options",
					Description: `(Optional) Specifies mount type specific options that are passed to the backend`,
				},
				resource.Attribute{
					Name:        "seal_wrap",
					Description: `(Optional) Boolean flag that can be explicitly set to true to enable seal wrapping for the mount, causing values stored by the mount to be wrapped by the seal's encryption capability`,
				},
				resource.Attribute{
					Name:        "external_entropy_access",
					Description: `(Optional) Boolean flag that can be explicitly set to true to enable the secrets engine to access Vault's external entropy source`,
				},
				resource.Attribute{
					Name:        "allowed_managed_keys",
					Description: `(Optional) Set of managed key registry entry names that the mount in question is allowed to access ## Attributes Reference In addition to the fields above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "accessor",
					Description: `The accessor for this mount. ## Import Mounts can be imported using the ` + "`" + `path` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_mount.example dummy ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "accessor",
					Description: `The accessor for this mount. ## Import Mounts can be imported using the ` + "`" + `path` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_mount.example dummy ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_namespace",
			Category:         "Resources",
			ShortDescription: `Writes namespaces for Vault`,
			Description:      ``,
			Keywords: []string{
				"namespace",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) The namespace to provision the resource in. The value should not contain leading or trailing forward slashes. The ` + "`" + `namespace` + "`" + ` is always relative to the provider's configured [namespace](/docs/providers/vault#namespace).`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Required) The path of the namespace. Must not have a trailing ` + "`" + `/` + "`" + ` ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the namespace.`,
				},
				resource.Attribute{
					Name:        "path_fq",
					Description: `The fully qualified path to the namespace. Useful when provisioning resources in a child ` + "`" + `namespace` + "`" + `. ## Import Namespaces can be imported using its ` + "`" + `name` + "`" + ` as accessor id ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_namespace.example <name> ` + "`" + `` + "`" + `` + "`" + ` If the declared resource is imported and intends to support namespaces using a provider alias, then the name is relative to the namespace path. ` + "`" + `` + "`" + `` + "`" + ` provider "vault" { # Configuration options namespace = "example" alias = "example" } resource vault_namespace "example2" { provider = vault.example } $ terraform import vault_namespace.example2 example2 $ terraform state show vault_namespace.example2 # vault_namespace.example2 resource "vault_namespace" "example2" { id = "example/example2/" namespace_id = <known after import> path = "example2" } ` + "`" + `` + "`" + `` + "`" + ` ## Tutorials Refer to the [Codify Management of Vault Enterprise Using Terraform](https://learn.hashicorp.com/tutorials/vault/codify-mgmt-enterprise) tutorial for additional examples using Vault namespaces.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the namespace.`,
				},
				resource.Attribute{
					Name:        "path_fq",
					Description: `The fully qualified path to the namespace. Useful when provisioning resources in a child ` + "`" + `namespace` + "`" + `. ## Import Namespaces can be imported using its ` + "`" + `name` + "`" + ` as accessor id ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_namespace.example <name> ` + "`" + `` + "`" + `` + "`" + ` If the declared resource is imported and intends to support namespaces using a provider alias, then the name is relative to the namespace path. ` + "`" + `` + "`" + `` + "`" + ` provider "vault" { # Configuration options namespace = "example" alias = "example" } resource vault_namespace "example2" { provider = vault.example } $ terraform import vault_namespace.example2 example2 $ terraform state show vault_namespace.example2 # vault_namespace.example2 resource "vault_namespace" "example2" { id = "example/example2/" namespace_id = <known after import> path = "example2" } ` + "`" + `` + "`" + `` + "`" + ` ## Tutorials Refer to the [Codify Management of Vault Enterprise Using Terraform](https://learn.hashicorp.com/tutorials/vault/codify-mgmt-enterprise) tutorial for additional examples using Vault namespaces.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_okta_auth_backend",
			Category:         "Resources",
			ShortDescription: `Managing Okta auth backends in Vault`,
			Description:      ``,
			Keywords: []string{
				"okta",
				"auth",
				"backend",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) The namespace to provision the resource in. The value should not contain leading or trailing forward slashes. The ` + "`" + `namespace` + "`" + ` is always relative to the provider's configured [namespace](/docs/providers/vault#namespace).`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Optional) Path to mount the Okta auth backend. Default to path ` + "`" + `okta` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "disable_remount",
					Description: `(Optional) If set, opts out of mount migration on path updates. See here for more info on [Mount Migration](https://www.vaultproject.io/docs/concepts/mount-migration)`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description of the auth backend`,
				},
				resource.Attribute{
					Name:        "organization",
					Description: `(Required) The Okta organization. This will be the first part of the url ` + "`" + `https://XXX.okta.com` + "`" + ``,
				},
				resource.Attribute{
					Name:        "token",
					Description: `(Optional) The Okta API token. This is required to query Okta for user group membership. If this is not supplied only locally configured groups will be enabled.`,
				},
				resource.Attribute{
					Name:        "base_url",
					Description: `(Optional) The Okta url. Examples: oktapreview.com, okta.com`,
				},
				resource.Attribute{
					Name:        "bypass_okta_mfa",
					Description: `(Optional) When true, requests by Okta for a MFA check will be bypassed. This also disallows certain status checks on the account, such as whether the password is expired.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `(Optional) Duration after which authentication will be expired. [See the documentation for info on valid duration formats](https://golang.org/pkg/time/#ParseDuration).`,
				},
				resource.Attribute{
					Name:        "max_ttl",
					Description: `(Optional) Maximum duration after which authentication will be expired [See the documentation for info on valid duration formats](https://golang.org/pkg/time/#ParseDuration).`,
				},
				resource.Attribute{
					Name:        "group",
					Description: `(Optional) Associate Okta groups with policies within Vault. [See below for more details](#okta-group).`,
				},
				resource.Attribute{
					Name:        "user",
					Description: `(Optional) Associate Okta users with groups or policies within Vault. [See below for more details](#okta-user). ### Okta Group`,
				},
				resource.Attribute{
					Name:        "group_name",
					Description: `(Required) Name of the group within the Okta`,
				},
				resource.Attribute{
					Name:        "policies",
					Description: `(Optional) Vault policies to associate with this group ### Okta User`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Required Optional) Name of the user within Okta`,
				},
				resource.Attribute{
					Name:        "groups",
					Description: `(Optional) List of Okta groups to associate with this user`,
				},
				resource.Attribute{
					Name:        "policies",
					Description: `(Optional) List of Vault policies to associate with this user ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "accessor",
					Description: `The mount accessor related to the auth mount. It is useful for integration with [Identity Secrets Engine](https://www.vaultproject.io/docs/secrets/identity/index.html). ## Import Okta authentication backends can be imported using its ` + "`" + `path` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_okta_auth_backend.example okta ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "accessor",
					Description: `The mount accessor related to the auth mount. It is useful for integration with [Identity Secrets Engine](https://www.vaultproject.io/docs/secrets/identity/index.html). ## Import Okta authentication backends can be imported using its ` + "`" + `path` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_okta_auth_backend.example okta ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_okta_auth_backend_group",
			Category:         "Resources",
			ShortDescription: `Managing groups in an Okta auth backend in Vault`,
			Description:      ``,
			Keywords: []string{
				"okta",
				"auth",
				"backend",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) The namespace to provision the resource in. The value should not contain leading or trailing forward slashes. The ` + "`" + `namespace` + "`" + ` is always relative to the provider's configured [namespace](/docs/providers/vault#namespace).`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Required) The path where the Okta auth backend is mounted`,
				},
				resource.Attribute{
					Name:        "group_name",
					Description: `(Required) Name of the group within the Okta`,
				},
				resource.Attribute{
					Name:        "policies",
					Description: `(Optional) Vault policies to associate with this group ## Attributes Reference No additional attributes are exposed by this resource. ## Import Okta authentication backend groups can be imported using the format ` + "`" + `backend/groupName` + "`" + ` e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_okta_auth_backend_group.foo okta/foo ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_okta_auth_backend_user",
			Category:         "Resources",
			ShortDescription: `Managing users in an Okta auth backend in Vault`,
			Description:      ``,
			Keywords: []string{
				"okta",
				"auth",
				"backend",
				"user",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) The namespace to provision the resource in. The value should not contain leading or trailing forward slashes. The ` + "`" + `namespace` + "`" + ` is always relative to the provider's configured [namespace](/docs/providers/vault#namespace).`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Required) The path where the Okta auth backend is mounted`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Required Optional) Name of the user within Okta`,
				},
				resource.Attribute{
					Name:        "groups",
					Description: `(Optional) List of Okta groups to associate with this user`,
				},
				resource.Attribute{
					Name:        "policies",
					Description: `(Optional) List of Vault policies to associate with this user ## Attributes Reference No additional attributes are exposed by this resource.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_pki_secret_backend_cert",
			Category:         "Resources",
			ShortDescription: `Generate an PKI certificate.`,
			Description:      ``,
			Keywords: []string{
				"pki",
				"secret",
				"backend",
				"cert",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) The namespace to provision the resource in. The value should not contain leading or trailing forward slashes. The ` + "`" + `namespace` + "`" + ` is always relative to the provider's configured [namespace](/docs/providers/vault#namespace).`,
				},
				resource.Attribute{
					Name:        "backend",
					Description: `(Required) The PKI secret backend the resource belongs to.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the role to create the certificate against`,
				},
				resource.Attribute{
					Name:        "common_name",
					Description: `(Required) CN of certificate to create`,
				},
				resource.Attribute{
					Name:        "alt_names",
					Description: `(Optional) List of alternative names`,
				},
				resource.Attribute{
					Name:        "ip_sans",
					Description: `(Optional) List of alternative IPs`,
				},
				resource.Attribute{
					Name:        "uri_sans",
					Description: `(Optional) List of alternative URIs`,
				},
				resource.Attribute{
					Name:        "other_sans",
					Description: `(Optional) List of other SANs`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `(Optional) Time to live`,
				},
				resource.Attribute{
					Name:        "format",
					Description: `(Optional) The format of data`,
				},
				resource.Attribute{
					Name:        "private_key_format",
					Description: `(Optional) The private key format`,
				},
				resource.Attribute{
					Name:        "exclude_cn_from_sans",
					Description: `(Optional) Flag to exclude CN from SANs`,
				},
				resource.Attribute{
					Name:        "min_seconds_remaining",
					Description: `(Optional) Generate a new certificate when the expiration is within this number of seconds, default is 604800 (7 days)`,
				},
				resource.Attribute{
					Name:        "auto_renew",
					Description: `(Optional) If set to ` + "`" + `true` + "`" + `, certs will be renewed if the expiration is within ` + "`" + `min_seconds_remaining` + "`" + `. Default ` + "`" + `false` + "`" + ``,
				},
				resource.Attribute{
					Name:        "revoke",
					Description: `If set to ` + "`" + `true` + "`" + `, the certificate will be revoked on resource destruction. ## Attributes Reference In addition to the fields above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "certificate",
					Description: `The certificate`,
				},
				resource.Attribute{
					Name:        "issuing_ca",
					Description: `The issuing CA`,
				},
				resource.Attribute{
					Name:        "ca_chain",
					Description: `The CA chain`,
				},
				resource.Attribute{
					Name:        "private_key",
					Description: `The private key`,
				},
				resource.Attribute{
					Name:        "private_key_type",
					Description: `The private key type`,
				},
				resource.Attribute{
					Name:        "serial_number",
					Description: `The serial number`,
				},
				resource.Attribute{
					Name:        "expiration",
					Description: `The expiration date of the certificate in unix epoch format`,
				},
				resource.Attribute{
					Name:        "renew_pending",
					Description: `` + "`" + `true` + "`" + ` if the current time (during refresh) is after the start of the early renewal window declared by ` + "`" + `min_seconds_remaining` + "`" + `, and ` + "`" + `false` + "`" + ` otherwise; if ` + "`" + `auto_renew` + "`" + ` is set to ` + "`" + `true` + "`" + ` then the provider will plan to replace the certificate once renewal is pending.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "certificate",
					Description: `The certificate`,
				},
				resource.Attribute{
					Name:        "issuing_ca",
					Description: `The issuing CA`,
				},
				resource.Attribute{
					Name:        "ca_chain",
					Description: `The CA chain`,
				},
				resource.Attribute{
					Name:        "private_key",
					Description: `The private key`,
				},
				resource.Attribute{
					Name:        "private_key_type",
					Description: `The private key type`,
				},
				resource.Attribute{
					Name:        "serial_number",
					Description: `The serial number`,
				},
				resource.Attribute{
					Name:        "expiration",
					Description: `The expiration date of the certificate in unix epoch format`,
				},
				resource.Attribute{
					Name:        "renew_pending",
					Description: `` + "`" + `true` + "`" + ` if the current time (during refresh) is after the start of the early renewal window declared by ` + "`" + `min_seconds_remaining` + "`" + `, and ` + "`" + `false` + "`" + ` otherwise; if ` + "`" + `auto_renew` + "`" + ` is set to ` + "`" + `true` + "`" + ` then the provider will plan to replace the certificate once renewal is pending.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_pki_secret_backend_config_ca",
			Category:         "Resources",
			ShortDescription: `Submit the CA information to PKI.`,
			Description:      ``,
			Keywords: []string{
				"pki",
				"secret",
				"backend",
				"config",
				"ca",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) The namespace to provision the resource in. The value should not contain leading or trailing forward slashes. The ` + "`" + `namespace` + "`" + ` is always relative to the provider's configured [namespace](/docs/providers/vault#namespace).`,
				},
				resource.Attribute{
					Name:        "backend",
					Description: `(Required) The PKI secret backend the resource belongs to.`,
				},
				resource.Attribute{
					Name:        "pem_bundle",
					Description: `(Required) The key and certificate PEM bundle ## Attributes Reference No additional attributes are exported by this resource.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_pki_secret_backend_config_urls",
			Category:         "Resources",
			ShortDescription: `Sets the config URL's on an PKI Secret Backend for Vault.`,
			Description:      ``,
			Keywords: []string{
				"pki",
				"secret",
				"backend",
				"config",
				"urls",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) The namespace to provision the resource in. The value should not contain leading or trailing forward slashes. The ` + "`" + `namespace` + "`" + ` is always relative to the provider's configured [namespace](/docs/providers/vault#namespace).`,
				},
				resource.Attribute{
					Name:        "backend",
					Description: `(Required) The path the PKI secret backend is mounted at, with no leading or trailing ` + "`" + `/` + "`" + `s.`,
				},
				resource.Attribute{
					Name:        "issuing_certificates",
					Description: `(Optional) Specifies the URL values for the Issuing Certificate field.`,
				},
				resource.Attribute{
					Name:        "crl_distribution_points",
					Description: `(Optional) Specifies the URL values for the CRL Distribution Points field.`,
				},
				resource.Attribute{
					Name:        "ocsp_servers",
					Description: `(Optional) Specifies the URL values for the OCSP Servers field. ## Attributes Reference No additional attributes are exported by this resource. ## Import The PKI config URLs can be imported using the resource's ` + "`" + `id` + "`" + `. In the case of the example above the ` + "`" + `id` + "`" + ` would be ` + "`" + `pki-root/config/urls` + "`" + `, where the ` + "`" + `pki-root` + "`" + ` component is the resource's ` + "`" + `backend` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_pki_secret_backend_config_urls.example pki-root/config/urls ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_pki_secret_backend_crl_config",
			Category:         "Resources",
			ShortDescription: `Sets the CRL config on an PKI Secret Backend for Vault.`,
			Description:      ``,
			Keywords: []string{
				"pki",
				"secret",
				"backend",
				"crl",
				"config",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) The namespace to provision the resource in. The value should not contain leading or trailing forward slashes. The ` + "`" + `namespace` + "`" + ` is always relative to the provider's configured [namespace](/docs/providers/vault#namespace).`,
				},
				resource.Attribute{
					Name:        "backend",
					Description: `(Required) The path the PKI secret backend is mounted at, with no leading or trailing ` + "`" + `/` + "`" + `s.`,
				},
				resource.Attribute{
					Name:        "expiry",
					Description: `(Optional) Specifies the time until expiration.`,
				},
				resource.Attribute{
					Name:        "disable",
					Description: `(Optional) Disables or enables CRL building.`,
				},
				resource.Attribute{
					Name:        "ocsp_disable",
					Description: `(Optional) Disables the OCSP responder in Vault.`,
				},
				resource.Attribute{
					Name:        "ocsp_expiry",
					Description: `(Optional) The amount of time an OCSP response can be cached for, useful for OCSP stapling refresh durations.`,
				},
				resource.Attribute{
					Name:        "auto_rebuild",
					Description: `(Optional) Enables periodic rebuilding of the CRL upon expiry.`,
				},
				resource.Attribute{
					Name:        "auto_rebuild_grace_period",
					Description: `(Optional) Grace period before CRL expiry to attempt rebuild of CRL.`,
				},
				resource.Attribute{
					Name:        "enable_delta",
					Description: `(Optional) Enables building of delta CRLs with up-to-date revocation information, augmenting the last complete CRL.`,
				},
				resource.Attribute{
					Name:        "delta_rebuild_interval",
					Description: `(Optional) Interval to check for new revocations on, to regenerate the delta CRL.`,
				},
				resource.Attribute{
					Name:        "cross_cluster_revocation",
					Description: `(Optional) Enable cross-cluster revocation request queues.`,
				},
				resource.Attribute{
					Name:        "unified_crl",
					Description: `(Optional) Enables unified CRL and OCSP building.`,
				},
				resource.Attribute{
					Name:        "unified_crl_on_existing_paths",
					Description: `(Optional) Enables serving the unified CRL and OCSP on the existing, previously cluster-local paths.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_pki_secret_backend_intermediate_cert_request",
			Category:         "Resources",
			ShortDescription: `Generate a new private key and a CSR for signing the PKI.`,
			Description:      ``,
			Keywords: []string{
				"pki",
				"secret",
				"backend",
				"intermediate",
				"cert",
				"request",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) The namespace to provision the resource in. The value should not contain leading or trailing forward slashes. The ` + "`" + `namespace` + "`" + ` is always relative to the provider's configured [namespace](/docs/providers/vault#namespace).`,
				},
				resource.Attribute{
					Name:        "backend",
					Description: `(Required) The PKI secret backend the resource belongs to.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) Type of intermediate to create. Must be either \"exported\" or \"internal\" or \"kms\"`,
				},
				resource.Attribute{
					Name:        "common_name",
					Description: `(Required) CN of intermediate to create`,
				},
				resource.Attribute{
					Name:        "alt_names",
					Description: `(Optional) List of alternative names`,
				},
				resource.Attribute{
					Name:        "ip_sans",
					Description: `(Optional) List of alternative IPs`,
				},
				resource.Attribute{
					Name:        "uri_sans",
					Description: `(Optional) List of alternative URIs`,
				},
				resource.Attribute{
					Name:        "other_sans",
					Description: `(Optional) List of other SANs`,
				},
				resource.Attribute{
					Name:        "format",
					Description: `(Optional) The format of data`,
				},
				resource.Attribute{
					Name:        "private_key_format",
					Description: `(Optional) The private key format`,
				},
				resource.Attribute{
					Name:        "key_type",
					Description: `(Optional) The desired key type`,
				},
				resource.Attribute{
					Name:        "key_bits",
					Description: `(Optional) The number of bits to use`,
				},
				resource.Attribute{
					Name:        "exclude_cn_from_sans",
					Description: `(Optional) Flag to exclude CN from SANs`,
				},
				resource.Attribute{
					Name:        "ou",
					Description: `(Optional) The organization unit`,
				},
				resource.Attribute{
					Name:        "organization",
					Description: `(Optional) The organization`,
				},
				resource.Attribute{
					Name:        "country",
					Description: `(Optional) The country`,
				},
				resource.Attribute{
					Name:        "locality",
					Description: `(Optional) The locality`,
				},
				resource.Attribute{
					Name:        "province",
					Description: `(Optional) The province`,
				},
				resource.Attribute{
					Name:        "street_address",
					Description: `(Optional) The street address`,
				},
				resource.Attribute{
					Name:        "postal_code",
					Description: `(Optional) The postal code`,
				},
				resource.Attribute{
					Name:        "managed_key_name",
					Description: `(Optional) The name of the previously configured managed key. This field is required if ` + "`" + `type` + "`" + ` is ` + "`" + `kms` + "`" + ` and it conflicts with ` + "`" + `managed_key_id` + "`" + ``,
				},
				resource.Attribute{
					Name:        "managed_key_id",
					Description: `(Optional) The ID of the previously configured managed key. This field is required if ` + "`" + `type` + "`" + ` is ` + "`" + `kms` + "`" + ` and it conflicts with ` + "`" + `managed_key_name` + "`" + ``,
				},
				resource.Attribute{
					Name:        "add_basic_constraints",
					Description: `(Optional) Adds a Basic Constraints extension with 'CA: true'. Only needed as a workaround in some compatibility scenarios with Active Directory Certificate Services ## Attributes Reference In addition to the fields above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "csr",
					Description: `The CSR`,
				},
				resource.Attribute{
					Name:        "private_key",
					Description: `The private key`,
				},
				resource.Attribute{
					Name:        "private_key_type",
					Description: `The private key type`,
				},
				resource.Attribute{
					Name:        "serial_number",
					Description: `The serial number`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "csr",
					Description: `The CSR`,
				},
				resource.Attribute{
					Name:        "private_key",
					Description: `The private key`,
				},
				resource.Attribute{
					Name:        "private_key_type",
					Description: `The private key type`,
				},
				resource.Attribute{
					Name:        "serial_number",
					Description: `The serial number`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_pki_secret_backend_intermediate_set_signed",
			Category:         "Resources",
			ShortDescription: `Submit the PKI CA certificate.`,
			Description:      ``,
			Keywords: []string{
				"pki",
				"secret",
				"backend",
				"intermediate",
				"set",
				"signed",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) The namespace to provision the resource in. The value should not contain leading or trailing forward slashes. The ` + "`" + `namespace` + "`" + ` is always relative to the provider's configured [namespace](/docs/providers/vault#namespace).`,
				},
				resource.Attribute{
					Name:        "backend",
					Description: `(Required) The PKI secret backend the resource belongs to.`,
				},
				resource.Attribute{
					Name:        "certificate",
					Description: `(Required) Specifies the PEM encoded certificate. May optionally append additional CA certificates to populate the whole chain, which will then enable returning the full chain from issue and sign operations. ## Attributes Reference No additional attributes are exported by this resource.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_pki_secret_backend_role",
			Category:         "Resources",
			ShortDescription: `Create a role on an PKI Secret Backend for Vault.`,
			Description:      ``,
			Keywords: []string{
				"pki",
				"secret",
				"backend",
				"role",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) The namespace to provision the resource in. The value should not contain leading or trailing forward slashes. The ` + "`" + `namespace` + "`" + ` is always relative to the provider's configured [namespace](/docs/providers/vault#namespace).`,
				},
				resource.Attribute{
					Name:        "backend",
					Description: `(Required) The path the PKI secret backend is mounted at, with no leading or trailing ` + "`" + `/` + "`" + `s.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name to identify this role within the backend. Must be unique within the backend.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `(Optional, integer) The TTL, in seconds, for any certificate issued against this role.`,
				},
				resource.Attribute{
					Name:        "max_ttl",
					Description: `(Optional, integer) The maximum lease TTL, in seconds, for the role.`,
				},
				resource.Attribute{
					Name:        "allow_localhost",
					Description: `(Optional) Flag to allow certificates for localhost`,
				},
				resource.Attribute{
					Name:        "allowed_domains",
					Description: `(Optional) List of allowed domains for certificates`,
				},
				resource.Attribute{
					Name:        "allowed_domains_template",
					Description: `(Optional) Flag, if set, ` + "`" + `allowed_domains` + "`" + ` can be specified using identity template expressions such as ` + "`" + `{{identity.entity.aliases.<mount accessor>.name}}` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "allow_bare_domains",
					Description: `(Optional) Flag to allow certificates matching the actual domain`,
				},
				resource.Attribute{
					Name:        "allow_subdomains",
					Description: `(Optional) Flag to allow certificates matching subdomains`,
				},
				resource.Attribute{
					Name:        "allow_glob_domains",
					Description: `(Optional) Flag to allow names containing glob patterns.`,
				},
				resource.Attribute{
					Name:        "allow_any_name",
					Description: `(Optional) Flag to allow any name`,
				},
				resource.Attribute{
					Name:        "enforce_hostnames",
					Description: `(Optional) Flag to allow only valid host names`,
				},
				resource.Attribute{
					Name:        "allow_ip_sans",
					Description: `(Optional) Flag to allow IP SANs`,
				},
				resource.Attribute{
					Name:        "allowed_uri_sans",
					Description: `(Optional) Defines allowed URI SANs`,
				},
				resource.Attribute{
					Name:        "allowed_other_sans",
					Description: `(Optional) Defines allowed custom SANs`,
				},
				resource.Attribute{
					Name:        "server_flag",
					Description: `(Optional) Flag to specify certificates for server use`,
				},
				resource.Attribute{
					Name:        "client_flag",
					Description: `(Optional) Flag to specify certificates for client use`,
				},
				resource.Attribute{
					Name:        "code_signing_flag",
					Description: `(Optional) Flag to specify certificates for code signing use`,
				},
				resource.Attribute{
					Name:        "email_protection_flag",
					Description: `(Optional) Flag to specify certificates for email protection use`,
				},
				resource.Attribute{
					Name:        "key_type",
					Description: `(Optional) The generated key type, choices: ` + "`" + `rsa` + "`" + `, ` + "`" + `ec` + "`" + `, ` + "`" + `ed25519` + "`" + `, ` + "`" + `any` + "`" + ` Defaults to ` + "`" + `rsa` + "`" + ``,
				},
				resource.Attribute{
					Name:        "key_bits",
					Description: `(Optional) The number of bits of generated keys`,
				},
				resource.Attribute{
					Name:        "key_usage",
					Description: `(Optional) Specify the allowed key usage constraint on issued certificates`,
				},
				resource.Attribute{
					Name:        "ext_key_usage",
					Description: `(Optional) Specify the allowed extended key usage constraint on issued certificates`,
				},
				resource.Attribute{
					Name:        "use_csr_common_name",
					Description: `(Optional) Flag to use the CN in the CSR`,
				},
				resource.Attribute{
					Name:        "use_csr_sans",
					Description: `(Optional) Flag to use the SANs in the CSR`,
				},
				resource.Attribute{
					Name:        "ou",
					Description: `(Optional) The organization unit of generated certificates`,
				},
				resource.Attribute{
					Name:        "organization",
					Description: `(Optional) The organization of generated certificates`,
				},
				resource.Attribute{
					Name:        "country",
					Description: `(Optional) The country of generated certificates`,
				},
				resource.Attribute{
					Name:        "locality",
					Description: `(Optional) The locality of generated certificates`,
				},
				resource.Attribute{
					Name:        "province",
					Description: `(Optional) The province of generated certificates`,
				},
				resource.Attribute{
					Name:        "street_address",
					Description: `(Optional) The street address of generated certificates`,
				},
				resource.Attribute{
					Name:        "postal_code",
					Description: `(Optional) The postal code of generated certificates`,
				},
				resource.Attribute{
					Name:        "generate_lease",
					Description: `(Optional) Flag to generate leases with certificates`,
				},
				resource.Attribute{
					Name:        "no_store",
					Description: `(Optional) Flag to not store certificates in the storage backend`,
				},
				resource.Attribute{
					Name:        "require_cn",
					Description: `(Optional) Flag to force CN usage`,
				},
				resource.Attribute{
					Name:        "policy_identifiers",
					Description: `(Optional) Specify the list of allowed policies OIDs. Use with Vault 1.10 or before. For Vault 1.11+, use ` + "`" + `policy_identifier` + "`" + ` blocks instead`,
				},
				resource.Attribute{
					Name:        "policy_identifier",
					Description: `(Optional) (Vault 1.11+ only) A block for specifying policy identifers. The ` + "`" + `policy_identifier` + "`" + ` block can be repeated, and supports the following arguments: - ` + "`" + `oid` + "`" + ` - (Required) The OID for the policy identifier - ` + "`" + `notice` + "`" + ` - (Optional) A notice for the policy identifier - ` + "`" + `cps` + "`" + ` - (Optional) The URL of the CPS for the policy identifier Example usage: ` + "`" + `` + "`" + `` + "`" + `hcl resource "vault_mount" "pki" { path = "pki" type = "pki" default_lease_ttl_seconds = 3600 max_lease_ttl_seconds = 86400 } resource "vault_pki_secret_backend_role" "role" { backend = vault_mount.pki.path name = "my_role" ttl = 3600 allow_ip_sans = true key_type = "rsa" key_bits = 4096 allowed_domains = ["example.com", "my.domain"] allow_subdomains = true policy_identifier { oid = "1.3.6.1.4.1.7.8" notice= "I am a user Notice" } policy_identifier { oid = "1.3.6.1.4.1.44947.1.2.4" cps ="https://example.com" } } ` + "`" + `` + "`" + `` + "`" + ``,
				},
				resource.Attribute{
					Name:        "basic_constraints_valid_for_non_ca",
					Description: `(Optional) Flag to mark basic constraints valid when issuing non-CA certificates`,
				},
				resource.Attribute{
					Name:        "not_before_duration",
					Description: `(Optional) Specifies the duration by which to backdate the NotBefore property.`,
				},
				resource.Attribute{
					Name:        "allowed_serial_numbers",
					Description: `(Optional) An array of allowed serial numbers to put in Subject ## Attributes Reference No additional attributes are exported by this resource. ## Import PKI secret backend roles can be imported using the ` + "`" + `path` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_pki_secret_backend_role.role pki/roles/my_role ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_pki_secret_backend_root_cert",
			Category:         "Resources",
			ShortDescription: `Generate root.`,
			Description:      ``,
			Keywords: []string{
				"pki",
				"secret",
				"backend",
				"root",
				"cert",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) The namespace to provision the resource in. The value should not contain leading or trailing forward slashes. The ` + "`" + `namespace` + "`" + ` is always relative to the provider's configured [namespace](/docs/providers/vault#namespace).`,
				},
				resource.Attribute{
					Name:        "backend",
					Description: `(Required) The PKI secret backend the resource belongs to.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) Type of intermediate to create. Must be either \"exported\", \"internal\" or \"kms\"`,
				},
				resource.Attribute{
					Name:        "common_name",
					Description: `(Required) CN of intermediate to create`,
				},
				resource.Attribute{
					Name:        "alt_names",
					Description: `(Optional) List of alternative names`,
				},
				resource.Attribute{
					Name:        "ip_sans",
					Description: `(Optional) List of alternative IPs`,
				},
				resource.Attribute{
					Name:        "uri_sans",
					Description: `(Optional) List of alternative URIs`,
				},
				resource.Attribute{
					Name:        "other_sans",
					Description: `(Optional) List of other SANs`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `(Optional) Time to live`,
				},
				resource.Attribute{
					Name:        "format",
					Description: `(Optional) The format of data`,
				},
				resource.Attribute{
					Name:        "private_key_format",
					Description: `(Optional) The private key format`,
				},
				resource.Attribute{
					Name:        "key_type",
					Description: `(Optional) The desired key type`,
				},
				resource.Attribute{
					Name:        "key_bits",
					Description: `(Optional) The number of bits to use`,
				},
				resource.Attribute{
					Name:        "max_path_length",
					Description: `(Optional) The maximum path length to encode in the generated certificate`,
				},
				resource.Attribute{
					Name:        "exclude_cn_from_sans",
					Description: `(Optional) Flag to exclude CN from SANs`,
				},
				resource.Attribute{
					Name:        "permitted_dns_domains",
					Description: `(Optional) List of domains for which certificates are allowed to be issued`,
				},
				resource.Attribute{
					Name:        "ou",
					Description: `(Optional) The organization unit`,
				},
				resource.Attribute{
					Name:        "organization",
					Description: `(Optional) The organization`,
				},
				resource.Attribute{
					Name:        "country",
					Description: `(Optional) The country`,
				},
				resource.Attribute{
					Name:        "locality",
					Description: `(Optional) The locality`,
				},
				resource.Attribute{
					Name:        "province",
					Description: `(Optional) The province`,
				},
				resource.Attribute{
					Name:        "street_address",
					Description: `(Optional) The street address`,
				},
				resource.Attribute{
					Name:        "postal_code",
					Description: `(Optional) The postal code`,
				},
				resource.Attribute{
					Name:        "managed_key_name",
					Description: `(Optional) The name of the previously configured managed key. This field is required if ` + "`" + `type` + "`" + ` is ` + "`" + `kms` + "`" + ` and it conflicts with ` + "`" + `managed_key_id` + "`" + ``,
				},
				resource.Attribute{
					Name:        "managed_key_id",
					Description: `(Optional) The ID of the previously configured managed key. This field is required if ` + "`" + `type` + "`" + ` is ` + "`" + `kms` + "`" + ` and it conflicts with ` + "`" + `managed_key_name` + "`" + ` ## Attributes Reference In addition to the fields above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "certificate",
					Description: `The certificate.`,
				},
				resource.Attribute{
					Name:        "issuing_ca",
					Description: `The issuing CA certificate.`,
				},
				resource.Attribute{
					Name:        "serial",
					Description: `Deprecated, use ` + "`" + `serial_number` + "`" + ` instead.`,
				},
				resource.Attribute{
					Name:        "serial_number",
					Description: `The certificate's serial number, hex formatted.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "certificate",
					Description: `The certificate.`,
				},
				resource.Attribute{
					Name:        "issuing_ca",
					Description: `The issuing CA certificate.`,
				},
				resource.Attribute{
					Name:        "serial",
					Description: `Deprecated, use ` + "`" + `serial_number` + "`" + ` instead.`,
				},
				resource.Attribute{
					Name:        "serial_number",
					Description: `The certificate's serial number, hex formatted.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_pki_secret_backend_root_sign_intermediate",
			Category:         "Resources",
			ShortDescription: `Signs intermediate certificate.`,
			Description:      ``,
			Keywords: []string{
				"pki",
				"secret",
				"backend",
				"root",
				"sign",
				"intermediate",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) The namespace to provision the resource in. The value should not contain leading or trailing forward slashes. The ` + "`" + `namespace` + "`" + ` is always relative to the provider's configured [namespace](/docs/providers/vault#namespace).`,
				},
				resource.Attribute{
					Name:        "backend",
					Description: `(Required) The PKI secret backend the resource belongs to.`,
				},
				resource.Attribute{
					Name:        "csr",
					Description: `(Required) The CSR`,
				},
				resource.Attribute{
					Name:        "common_name",
					Description: `(Required) CN of intermediate to create`,
				},
				resource.Attribute{
					Name:        "alt_names",
					Description: `(Optional) List of alternative names`,
				},
				resource.Attribute{
					Name:        "ip_sans",
					Description: `(Optional) List of alternative IPs`,
				},
				resource.Attribute{
					Name:        "uri_sans",
					Description: `(Optional) List of alternative URIs`,
				},
				resource.Attribute{
					Name:        "other_sans",
					Description: `(Optional) List of other SANs`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `(Optional) Time to live`,
				},
				resource.Attribute{
					Name:        "format",
					Description: `(Optional) The format of data`,
				},
				resource.Attribute{
					Name:        "max_path_length",
					Description: `(Optional) The maximum path length to encode in the generated certificate`,
				},
				resource.Attribute{
					Name:        "exclude_cn_from_sans",
					Description: `(Optional) Flag to exclude CN from SANs`,
				},
				resource.Attribute{
					Name:        "use_csr_values",
					Description: `(Optional) Preserve CSR values`,
				},
				resource.Attribute{
					Name:        "permitted_dns_domains",
					Description: `(Optional) List of domains for which certificates are allowed to be issued`,
				},
				resource.Attribute{
					Name:        "ou",
					Description: `(Optional) The organization unit`,
				},
				resource.Attribute{
					Name:        "organization",
					Description: `(Optional) The organization`,
				},
				resource.Attribute{
					Name:        "country",
					Description: `(Optional) The country`,
				},
				resource.Attribute{
					Name:        "locality",
					Description: `(Optional) The locality`,
				},
				resource.Attribute{
					Name:        "province",
					Description: `(Optional) The province`,
				},
				resource.Attribute{
					Name:        "street_address",
					Description: `(Optional) The street address`,
				},
				resource.Attribute{
					Name:        "postal_code",
					Description: `(Optional) The postal code`,
				},
				resource.Attribute{
					Name:        "revoke",
					Description: `If set to ` + "`" + `true` + "`" + `, the certificate will be revoked on resource destruction. ## Attributes Reference In addition to the fields above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "certificate",
					Description: `The intermediate CA certificate in the ` + "`" + `format` + "`" + ` specified.`,
				},
				resource.Attribute{
					Name:        "issuing_ca",
					Description: `The issuing CA certificate in the ` + "`" + `format` + "`" + ` specified.`,
				},
				resource.Attribute{
					Name:        "ca_chain",
					Description: `A list of the issuing and intermediate CA certificates in the ` + "`" + `format` + "`" + ` specified.`,
				},
				resource.Attribute{
					Name:        "certificate_bundle",
					Description: `The concatenation of the intermediate CA and the issuing CA certificates (PEM encoded). Requires the ` + "`" + `format` + "`" + ` to be set to any of: pem, pem_bundle. The value will be empty for all other formats.`,
				},
				resource.Attribute{
					Name:        "serial_number",
					Description: `The certificate's serial number, hex formatted. ## Deprecations`,
				},
				resource.Attribute{
					Name:        "serial",
					Description: `Use ` + "`" + `serial_number` + "`" + ` instead.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "certificate",
					Description: `The intermediate CA certificate in the ` + "`" + `format` + "`" + ` specified.`,
				},
				resource.Attribute{
					Name:        "issuing_ca",
					Description: `The issuing CA certificate in the ` + "`" + `format` + "`" + ` specified.`,
				},
				resource.Attribute{
					Name:        "ca_chain",
					Description: `A list of the issuing and intermediate CA certificates in the ` + "`" + `format` + "`" + ` specified.`,
				},
				resource.Attribute{
					Name:        "certificate_bundle",
					Description: `The concatenation of the intermediate CA and the issuing CA certificates (PEM encoded). Requires the ` + "`" + `format` + "`" + ` to be set to any of: pem, pem_bundle. The value will be empty for all other formats.`,
				},
				resource.Attribute{
					Name:        "serial_number",
					Description: `The certificate's serial number, hex formatted. ## Deprecations`,
				},
				resource.Attribute{
					Name:        "serial",
					Description: `Use ` + "`" + `serial_number` + "`" + ` instead.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_pki_secret_backend_sign",
			Category:         "Resources",
			ShortDescription: `Sign a new certificate based on the CSR by the PKI.`,
			Description:      ``,
			Keywords: []string{
				"pki",
				"secret",
				"backend",
				"sign",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) The namespace to provision the resource in. The value should not contain leading or trailing forward slashes. The ` + "`" + `namespace` + "`" + ` is always relative to the provider's configured [namespace](/docs/providers/vault#namespace).`,
				},
				resource.Attribute{
					Name:        "backend",
					Description: `(Required) The PKI secret backend the resource belongs to.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the role to create the certificate against`,
				},
				resource.Attribute{
					Name:        "csr",
					Description: `(Required) The CSR`,
				},
				resource.Attribute{
					Name:        "common_name",
					Description: `(Required) CN of certificate to create`,
				},
				resource.Attribute{
					Name:        "alt_names",
					Description: `(Optional) List of alternative names`,
				},
				resource.Attribute{
					Name:        "other_sans",
					Description: `(Optional) List of other SANs`,
				},
				resource.Attribute{
					Name:        "ip_sans",
					Description: `(Optional) List of alternative IPs`,
				},
				resource.Attribute{
					Name:        "uri_sans",
					Description: `(Optional) List of alternative URIs`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `(Optional) Time to live`,
				},
				resource.Attribute{
					Name:        "format",
					Description: `(Optional) The format of data`,
				},
				resource.Attribute{
					Name:        "exclude_cn_from_sans",
					Description: `(Optional) Flag to exclude CN from SANs`,
				},
				resource.Attribute{
					Name:        "min_seconds_remaining",
					Description: `(Optional) Generate a new certificate when the expiration is within this number of seconds, default is 604800 (7 days)`,
				},
				resource.Attribute{
					Name:        "auto_renew",
					Description: `(Optional) If set to ` + "`" + `true` + "`" + `, certs will be renewed if the expiration is within ` + "`" + `min_seconds_remaining` + "`" + `. Default ` + "`" + `false` + "`" + ` ## Attributes Reference In addition to the fields above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "certificate",
					Description: `The certificate`,
				},
				resource.Attribute{
					Name:        "issuing_ca",
					Description: `The issuing CA`,
				},
				resource.Attribute{
					Name:        "ca_chain",
					Description: `The CA chain`,
				},
				resource.Attribute{
					Name:        "serial_number",
					Description: `The certificate's serial number, hex formatted.`,
				},
				resource.Attribute{
					Name:        "expiration",
					Description: `The expiration date of the certificate in unix epoch format`,
				},
				resource.Attribute{
					Name:        "renew_pending",
					Description: `` + "`" + `true` + "`" + ` if the current time (during refresh) is after the start of the early renewal window declared by ` + "`" + `min_seconds_remaining` + "`" + `, and ` + "`" + `false` + "`" + ` otherwise; if ` + "`" + `auto_renew` + "`" + ` is set to ` + "`" + `true` + "`" + ` then the provider will plan to replace the certificate once renewal is pending. ## Deprecations`,
				},
				resource.Attribute{
					Name:        "serial",
					Description: `Use ` + "`" + `serial_number` + "`" + ` instead.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "certificate",
					Description: `The certificate`,
				},
				resource.Attribute{
					Name:        "issuing_ca",
					Description: `The issuing CA`,
				},
				resource.Attribute{
					Name:        "ca_chain",
					Description: `The CA chain`,
				},
				resource.Attribute{
					Name:        "serial_number",
					Description: `The certificate's serial number, hex formatted.`,
				},
				resource.Attribute{
					Name:        "expiration",
					Description: `The expiration date of the certificate in unix epoch format`,
				},
				resource.Attribute{
					Name:        "renew_pending",
					Description: `` + "`" + `true` + "`" + ` if the current time (during refresh) is after the start of the early renewal window declared by ` + "`" + `min_seconds_remaining` + "`" + `, and ` + "`" + `false` + "`" + ` otherwise; if ` + "`" + `auto_renew` + "`" + ` is set to ` + "`" + `true` + "`" + ` then the provider will plan to replace the certificate once renewal is pending. ## Deprecations`,
				},
				resource.Attribute{
					Name:        "serial",
					Description: `Use ` + "`" + `serial_number` + "`" + ` instead.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_policy",
			Category:         "Resources",
			ShortDescription: `Writes arbitrary policies for Vault`,
			Description:      ``,
			Keywords: []string{
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) The namespace to provision the resource in. The value should not contain leading or trailing forward slashes. The ` + "`" + `namespace` + "`" + ` is always relative to the provider's configured [namespace](/docs/providers/vault#namespace).`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the policy`,
				},
				resource.Attribute{
					Name:        "policy",
					Description: `(Required) String containing a Vault policy ## Attributes Reference No additional attributes are exported by this resource. ## Import Policies can be imported using the ` + "`" + `name` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_policy.example dev-team ` + "`" + `` + "`" + `` + "`" + ` ## Tutorials Refer to the following tutorials for additional usage examples: - [Codify Management of Vault Enterprise Using Terraform](https://learn.hashicorp.com/tutorials/vault/codify-mgmt-enterprise) - [Codify Management of Vault Using Terraform](https://learn.hashicorp.com/tutorials/vault/codify-mgmt-oss)`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_quota_lease_count",
			Category:         "Resources",
			ShortDescription: `Manage Lease Count Quota`,
			Description:      ``,
			Keywords: []string{
				"quota",
				"lease",
				"count",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) The namespace to provision the resource in. The value should not contain leading or trailing forward slashes. The ` + "`" + `namespace` + "`" + ` is always relative to the provider's configured [namespace](../index.html#namespace).`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the rate limit quota`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Optional) Path of the mount or namespace to apply the quota. A blank path configures a global rate limit quota. For example ` + "`" + `namespace1/` + "`" + ` adds a quota to a full namespace, ` + "`" + `namespace1/auth/userpass` + "`" + ` adds a ` + "`" + `quota` + "`" + ` to ` + "`" + `userpass` + "`" + ` in ` + "`" + `namespace1` + "`" + `. Updating this field on an existing quota can have "moving" effects. For example, updating ` + "`" + `auth/userpass` + "`" + ` to ` + "`" + `namespace1/auth/userpass` + "`" + ` moves this quota from being a global mount quota to a namespace specific mount quota.`,
				},
				resource.Attribute{
					Name:        "max_leases",
					Description: `(Required) The maximum number of leases to be allowed by the quota rule. The ` + "`" + `max_leases` + "`" + ` must be positive. ## Attributes Reference No additional attributes are exported by this resource. ## Import Lease count quotas can be imported using their names ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_quota_lease_count.global global ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_quota_rate_limit",
			Category:         "Resources",
			ShortDescription: `Manage Rate Limit Quota`,
			Description:      ``,
			Keywords: []string{
				"quota",
				"rate",
				"limit",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) The namespace to provision the resource in. The value should not contain leading or trailing forward slashes. The ` + "`" + `namespace` + "`" + ` is always relative to the provider's configured [namespace](../index.html#namespace).`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the rate limit quota`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Optional) Path of the mount or namespace to apply the quota. A blank path configures a global rate limit quota. For example ` + "`" + `namespace1/` + "`" + ` adds a quota to a full namespace, ` + "`" + `namespace1/auth/userpass` + "`" + ` adds a ` + "`" + `quota` + "`" + ` to ` + "`" + `userpass` + "`" + ` in ` + "`" + `namespace1` + "`" + `. Updating this field on an existing quota can have "moving" effects. For example, updating ` + "`" + `auth/userpass` + "`" + ` to ` + "`" + `namespace1/auth/userpass` + "`" + ` moves this quota from being a global mount quota to a namespace specific mount quota.`,
				},
				resource.Attribute{
					Name:        "rate",
					Description: `(Required) The maximum number of requests at any given second to be allowed by the quota rule. The ` + "`" + `rate` + "`" + ` must be positive.`,
				},
				resource.Attribute{
					Name:        "interval",
					Description: `(Optional) The duration in seconds to enforce rate limiting for.`,
				},
				resource.Attribute{
					Name:        "block_interval",
					Description: `(Optional) If set, when a client reaches a rate limit threshold, the client will be prohibited from any further requests until after the 'block_interval' in seconds has elapsed. ## Attributes Reference No additional attributes are exported by this resource. ## Import Rate limit quotas can be imported using their names ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_quota_rate_limit.global global ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_rabbitmq_secret_backend",
			Category:         "Resources",
			ShortDescription: `Creates an RabbitMQ secret backend for Vault.`,
			Description:      ``,
			Keywords: []string{
				"rabbitmq",
				"secret",
				"backend",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) The namespace to provision the resource in. The value should not contain leading or trailing forward slashes. The ` + "`" + `namespace` + "`" + ` is always relative to the provider's configured [namespace](/docs/providers/vault#namespace).`,
				},
				resource.Attribute{
					Name:        "connection_uri",
					Description: `(Required) Specifies the RabbitMQ connection URI.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Required) Specifies the RabbitMQ management administrator username.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Required) Specifies the RabbitMQ management administrator password.`,
				},
				resource.Attribute{
					Name:        "verify_connection",
					Description: `(Optional) Specifies whether to verify connection URI, username, and password. Defaults to ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "password_policy",
					Description: `(Optional) Specifies a password policy to use when creating dynamic credentials. Defaults to generating an alphanumeric password if not set.`,
				},
				resource.Attribute{
					Name:        "username_template",
					Description: `(Optional) Template describing how dynamic usernames are generated. ~>`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Optional) The unique path this backend should be mounted at. Must not begin or end with a ` + "`" + `/` + "`" + `. Defaults to ` + "`" + `rabbitmq` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "disable_remount",
					Description: `(Optional) If set, opts out of mount migration on path updates. See here for more info on [Mount Migration](https://www.vaultproject.io/docs/concepts/mount-migration)`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A human-friendly description for this backend.`,
				},
				resource.Attribute{
					Name:        "default_lease_ttl_seconds",
					Description: `(Optional) The default TTL for credentials issued by this backend.`,
				},
				resource.Attribute{
					Name:        "max_lease_ttl_seconds",
					Description: `(Optional) The maximum TTL that can be requested for credentials issued by this backend. ## Attributes Reference No additional attributes are exported by this resource. ## Import RabbitMQ secret backends can be imported using the ` + "`" + `path` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_rabbitmq_secret_backend.rabbitmq rabbitmq ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_rabbitmq_secret_backend_role",
			Category:         "Resources",
			ShortDescription: `Creates a role on an RabbitMQ Secret Backend for Vault.`,
			Description:      ``,
			Keywords: []string{
				"rabbitmq",
				"secret",
				"backend",
				"role",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) The namespace to provision the resource in. The value should not contain leading or trailing forward slashes. The ` + "`" + `namespace` + "`" + ` is always relative to the provider's configured [namespace](/docs/providers/vault#namespace).`,
				},
				resource.Attribute{
					Name:        "backend",
					Description: `(Required) The path the RabbitMQ secret backend is mounted at, with no leading or trailing ` + "`" + `/` + "`" + `s.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name to identify this role within the backend. Must be unique within the backend.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Specifies a comma-separated RabbitMQ management tags.`,
				},
				resource.Attribute{
					Name:        "vhost",
					Description: `(Optional) Specifies a map of virtual hosts to permissions.`,
				},
				resource.Attribute{
					Name:        "vhost_topic",
					Description: `(Optional) Specifies a map of virtual hosts and exchanges to topic permissions. This option requires RabbitMQ 3.7.0 or later. ## Attributes Reference No additional attributes are exported by this resource. ## Import RabbitMQ secret backend roles can be imported using the ` + "`" + `path` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_rabbitmq_secret_backend_role.role rabbitmq/roles/deploy ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_rgp_policy",
			Category:         "Resources",
			ShortDescription: `Writes Sentinel role governing policies for Vault`,
			Description:      ``,
			Keywords: []string{
				"rgp",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) The namespace to provision the resource in. The value should not contain leading or trailing forward slashes. The ` + "`" + `namespace` + "`" + ` is always relative to the provider's configured [namespace](/docs/providers/vault#namespace).`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the policy`,
				},
				resource.Attribute{
					Name:        "enforcement_level",
					Description: `(Required) Enforcement level of Sentinel policy. Can be either ` + "`" + `advisory` + "`" + ` or ` + "`" + `soft-mandatory` + "`" + ` or ` + "`" + `hard-mandatory` + "`" + ``,
				},
				resource.Attribute{
					Name:        "policy",
					Description: `(Required) String containing a Sentinel policy ## Attributes Reference No additional attributes are exported by this resource.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_ssh_secret_backend_ca",
			Category:         "Resources",
			ShortDescription: `Managing CA information in an SSH secret backend in Vault`,
			Description:      ``,
			Keywords: []string{
				"ssh",
				"secret",
				"backend",
				"ca",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) The namespace to provision the resource in. The value should not contain leading or trailing forward slashes. The ` + "`" + `namespace` + "`" + ` is always relative to the provider's configured [namespace](/docs/providers/vault#namespace).`,
				},
				resource.Attribute{
					Name:        "backend",
					Description: `(Optional) The path where the SSH secret backend is mounted. Defaults to 'ssh'`,
				},
				resource.Attribute{
					Name:        "generate_signing_key",
					Description: `(Optional) Whether Vault should generate the signing key pair internally. Defaults to true`,
				},
				resource.Attribute{
					Name:        "public_key",
					Description: `(Optional) The public key part the SSH CA key pair; required if generate_signing_key is false.`,
				},
				resource.Attribute{
					Name:        "private_key",
					Description: `(Optional) The private key part the SSH CA key pair; required if generate_signing_key is false. ~>`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_ssh_secret_backend_role",
			Category:         "Resources",
			ShortDescription: `Managing roles in an SSH secret backend in Vault`,
			Description:      ``,
			Keywords: []string{
				"ssh",
				"secret",
				"backend",
				"role",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) The namespace to provision the resource in. The value should not contain leading or trailing forward slashes. The ` + "`" + `namespace` + "`" + ` is always relative to the provider's configured [namespace](/docs/providers/vault#namespace).`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Specifies the name of the role to create.`,
				},
				resource.Attribute{
					Name:        "backend",
					Description: `(Required) The path where the SSH secret backend is mounted.`,
				},
				resource.Attribute{
					Name:        "key_type",
					Description: `(Required) Specifies the type of credentials generated by this role. This can be either ` + "`" + `otp` + "`" + `, ` + "`" + `dynamic` + "`" + ` or ` + "`" + `ca` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "allow_bare_domains",
					Description: `(Optional) Specifies if host certificates that are requested are allowed to use the base domains listed in ` + "`" + `allowed_domains` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "allow_host_certificates",
					Description: `(Optional) Specifies if certificates are allowed to be signed for use as a 'host'.`,
				},
				resource.Attribute{
					Name:        "allow_subdomains",
					Description: `(Optional) Specifies if host certificates that are requested are allowed to be subdomains of those listed in ` + "`" + `allowed_domains` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "allow_user_certificates",
					Description: `(Optional) Specifies if certificates are allowed to be signed for use as a 'user'.`,
				},
				resource.Attribute{
					Name:        "allow_user_key_ids",
					Description: `(Optional) Specifies if users can override the key ID for a signed certificate with the ` + "`" + `key_id` + "`" + ` field.`,
				},
				resource.Attribute{
					Name:        "allowed_critical_options",
					Description: `(Optional) Specifies a comma-separated list of critical options that certificates can have when signed.`,
				},
				resource.Attribute{
					Name:        "allowed_domains",
					Description: `(Optional) The list of domains for which a client can request a host certificate.`,
				},
				resource.Attribute{
					Name:        "cidr_list",
					Description: `(Optional) The comma-separated string of CIDR blocks for which this role is applicable.`,
				},
				resource.Attribute{
					Name:        "allowed_extensions",
					Description: `(Optional) Specifies a comma-separated list of extensions that certificates can have when signed.`,
				},
				resource.Attribute{
					Name:        "default_extensions",
					Description: `(Optional) Specifies a map of extensions that certificates have when signed.`,
				},
				resource.Attribute{
					Name:        "default_critical_options",
					Description: `(Optional) Specifies a map of critical options that certificates have when signed.`,
				},
				resource.Attribute{
					Name:        "allowed_users_template",
					Description: `(Optional) Specifies if ` + "`" + `allowed_users` + "`" + ` can be declared using identity template policies. Non-templated users are also permitted.`,
				},
				resource.Attribute{
					Name:        "allowed_users",
					Description: `(Optional) Specifies a comma-separated list of usernames that are to be allowed, only if certain usernames are to be allowed.`,
				},
				resource.Attribute{
					Name:        "default_user_template",
					Description: `(Optional) If set, ` + "`" + `default_users` + "`" + ` can be specified using identity template values. A non-templated user is also permitted.`,
				},
				resource.Attribute{
					Name:        "default_user",
					Description: `(Optional) Specifies the default username for which a credential will be generated.`,
				},
				resource.Attribute{
					Name:        "key_id_format",
					Description: `(Optional) Specifies a custom format for the key id of a signed certificate.`,
				},
				resource.Attribute{
					Name:        "algorithm_signer",
					Description: `(Optional) When supplied, this value specifies a signing algorithm for the key. Possible values: ssh-rsa, rsa-sha2-256, rsa-sha2-512.`,
				},
				resource.Attribute{
					Name:        "allowed_user_key_config",
					Description: `(Optional) Set of configuration blocks to define allowed user key configuration, like key type and their lengths. Can be specified multiple times.`,
				},
				resource.Attribute{
					Name:        "allowed_user_key_lengths",
					Description: `(Optional) Specifies a map of ssh key types and their expected sizes which are allowed to be signed by the CA type.`,
				},
				resource.Attribute{
					Name:        "max_ttl",
					Description: `(Optional) Specifies the maximum Time To Live value.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `(Optional) Specifies the Time To Live value. ### Allowed User Key Configuration`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) The SSH public key type.`,
				},
				resource.Attribute{
					Name:        "lengths",
					Description: `(Required) A list of allowed key lengths as integers. For key types that do not support setting the length a value of ` + "`" + `[0]` + "`" + ` should be used. Setting multiple lengths is only supported on Vault 1.10+. For prior releases ` + "`" + `length` + "`" + ` must be set to a single element list. Example configuration blocks that might be included in the ` + "`" + `vault_ssh_secret_backend_role` + "`" + ` ` + "`" + `` + "`" + `` + "`" + `hcl allowed_user_key_config { type = "rsa" lengths = [2048, 4096] } allowed_user_key_config { type = "dss" lengths = [2048, 4096] } ` + "`" + `` + "`" + `` + "`" + ` ## Attributes Reference No additional attributes are exposed by this resource. ## Import SSH secret backend roles can be imported using the ` + "`" + `path` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_ssh_secret_backend_role.foo ssh/roles/my-role ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_token",
			Category:         "Resources",
			ShortDescription: `Writes token for Vault`,
			Description:      ``,
			Keywords: []string{
				"token",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) The namespace to provision the resource in. The value should not contain leading or trailing forward slashes. The ` + "`" + `namespace` + "`" + ` is always relative to the provider's configured [namespace](/docs/providers/vault#namespace).`,
				},
				resource.Attribute{
					Name:        "role_name",
					Description: `(Optional) The token role name`,
				},
				resource.Attribute{
					Name:        "policies",
					Description: `(Optional) List of policies to attach to this token`,
				},
				resource.Attribute{
					Name:        "no_parent",
					Description: `(Optional) Flag to create a token without parent`,
				},
				resource.Attribute{
					Name:        "no_default_policy",
					Description: `(Optional) Flag to not attach the default policy to this token`,
				},
				resource.Attribute{
					Name:        "renewable",
					Description: `(Optional) Flag to allow to renew this token`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `(Optional) The TTL period of this token`,
				},
				resource.Attribute{
					Name:        "explicit_max_ttl",
					Description: `(Optional) The explicit max TTL of this token`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) String containing the token display name`,
				},
				resource.Attribute{
					Name:        "num_uses",
					Description: `(Optional) The number of allowed uses of this token`,
				},
				resource.Attribute{
					Name:        "period",
					Description: `(Optional) The period of this token`,
				},
				resource.Attribute{
					Name:        "renew_min_lease",
					Description: `(Optional) The minimal lease to renew this token`,
				},
				resource.Attribute{
					Name:        "renew_increment",
					Description: `(Optional) The renew increment`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `(Optional) Metadata to be set on this token ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "lease_duration",
					Description: `String containing the token lease duration if present in state file`,
				},
				resource.Attribute{
					Name:        "lease_started",
					Description: `String containing the token lease started time if present in state file`,
				},
				resource.Attribute{
					Name:        "client_token",
					Description: `String containing the client token if stored in present file ## Import Tokens can be imported using its ` + "`" + `id` + "`" + ` as accessor id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_token.example <accessor_id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "lease_duration",
					Description: `String containing the token lease duration if present in state file`,
				},
				resource.Attribute{
					Name:        "lease_started",
					Description: `String containing the token lease started time if present in state file`,
				},
				resource.Attribute{
					Name:        "client_token",
					Description: `String containing the client token if stored in present file ## Import Tokens can be imported using its ` + "`" + `id` + "`" + ` as accessor id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_token.example <accessor_id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_token_auth_backend_role",
			Category:         "Resources",
			ShortDescription: `Manages Token auth backend roles in Vault.`,
			Description:      ``,
			Keywords: []string{
				"token",
				"auth",
				"backend",
				"role",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) The namespace to provision the resource in. The value should not contain leading or trailing forward slashes. The ` + "`" + `namespace` + "`" + ` is always relative to the provider's configured [namespace](/docs/providers/vault#namespace).`,
				},
				resource.Attribute{
					Name:        "role_name",
					Description: `(Required) The name of the role.`,
				},
				resource.Attribute{
					Name:        "token_ttl",
					Description: `(Optional) The incremental lifetime for generated tokens in number of seconds. Its current value will be referenced at renewal time.`,
				},
				resource.Attribute{
					Name:        "token_max_ttl",
					Description: `(Optional) The maximum lifetime for generated tokens in number of seconds. Its current value will be referenced at renewal time.`,
				},
				resource.Attribute{
					Name:        "token_period",
					Description: `(Optional) If set, indicates that the token generated using this role should never expire. The token should be renewed within the duration specified by this value. At each renewal, the token's TTL will be set to the value of this field. Specified in seconds.`,
				},
				resource.Attribute{
					Name:        "token_bound_cidrs",
					Description: `(Optional) List of CIDR blocks; if set, specifies blocks of IP addresses which can authenticate successfully, and ties the resulting token to these blocks as well.`,
				},
				resource.Attribute{
					Name:        "token_explicit_max_ttl",
					Description: `(Optional) If set, will encode an [explicit max TTL](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls) onto the token in number of seconds. This is a hard cap even if ` + "`" + `token_ttl` + "`" + ` and ` + "`" + `token_max_ttl` + "`" + ` would otherwise allow a renewal.`,
				},
				resource.Attribute{
					Name:        "token_no_default_policy",
					Description: `(Optional) If set, the default policy will not be set on generated tokens; otherwise it will be added to the policies set in token_policies.`,
				},
				resource.Attribute{
					Name:        "token_num_uses",
					Description: `(Optional) The [maximum number](https://www.vaultproject.io/api-docs/token#token_num_uses) of times a generated token may be used (within its lifetime); 0 means unlimited.`,
				},
				resource.Attribute{
					Name:        "token_type",
					Description: `(Optional) The type of token that should be generated. Can be ` + "`" + `service` + "`" + `, ` + "`" + `batch` + "`" + `, or ` + "`" + `default` + "`" + ` to use the mount's tuned default (which unless changed will be ` + "`" + `service` + "`" + ` tokens). For token store roles, there are two additional possibilities: ` + "`" + `default-service` + "`" + ` and ` + "`" + `default-batch` + "`" + ` which specify the type to return unless the client requests a different type at generation time. ## Attributes Reference No additional attributes are exported by this resource. ## Import Token auth backend roles can be imported with ` + "`" + `auth/token/roles/` + "`" + ` followed by the ` + "`" + `role_name` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_token_auth_backend_role.example auth/token/roles/my-role ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_transform_alphabet",
			Category:         "Resources",
			ShortDescription: `"/transform/alphabet/{name}"`,
			Description:      ``,
			Keywords: []string{
				"transform",
				"alphabet",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) The namespace to provision the resource in. The value should not contain leading or trailing forward slashes. The ` + "`" + `namespace` + "`" + ` is always relative to the provider's configured [namespace](/docs/providers/vault#namespace).`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Required) Path to where the back-end is mounted within Vault.`,
				},
				resource.Attribute{
					Name:        "alphabet",
					Description: `(Optional) A string of characters that contains the alphabet set.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the alphabet.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_transform_role",
			Category:         "Resources",
			ShortDescription: `"/transform/role/{name}"`,
			Description:      ``,
			Keywords: []string{
				"transform",
				"role",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) The namespace to provision the resource in. The value should not contain leading or trailing forward slashes. The ` + "`" + `namespace` + "`" + ` is always relative to the provider's configured [namespace](/docs/providers/vault#namespace).`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Required) Path to where the back-end is mounted within Vault.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the role.`,
				},
				resource.Attribute{
					Name:        "transformations",
					Description: `(Optional) A comma separated string or slice of transformations to use.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_transform_template",
			Category:         "Resources",
			ShortDescription: `"/transform/template/{name}"`,
			Description:      ``,
			Keywords: []string{
				"transform",
				"template",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) The namespace to provision the resource in. The value should not contain leading or trailing forward slashes. The ` + "`" + `namespace` + "`" + ` is always relative to the provider's configured [namespace](/docs/providers/vault#namespace).`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Required) Path to where the back-end is mounted within Vault.`,
				},
				resource.Attribute{
					Name:        "alphabet",
					Description: `(Optional) The alphabet to use for this template. This is only used during FPE transformations.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the template.`,
				},
				resource.Attribute{
					Name:        "pattern",
					Description: `(Optional) The pattern used for matching. Currently, only regular expression pattern is supported.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) The pattern type to use for match detection. Currently, only regex is supported.`,
				},
				resource.Attribute{
					Name:        "encode_format",
					Description: `(Optional) - The regular expression template used to format encoded values. (requires Vault Enterprise 1.9+)`,
				},
				resource.Attribute{
					Name:        "decode_formats",
					Description: `(Optional) - Optional mapping of name to regular expression template, used to customize the decoded output. (requires Vault Enterprise 1.9+)`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_transform_transformation",
			Category:         "Resources",
			ShortDescription: `"/transform/transformation/{name}"`,
			Description:      ``,
			Keywords: []string{
				"transform",
				"transformation",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) The namespace to provision the resource in. The value should not contain leading or trailing forward slashes. The ` + "`" + `namespace` + "`" + ` is always relative to the provider's configured [namespace](/docs/providers/vault#namespace).`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Required) Path to where the back-end is mounted within Vault.`,
				},
				resource.Attribute{
					Name:        "allowed_roles",
					Description: `(Optional) The set of roles allowed to perform this transformation.`,
				},
				resource.Attribute{
					Name:        "masking_character",
					Description: `(Optional) The character used to replace data when in masking mode`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the transformation.`,
				},
				resource.Attribute{
					Name:        "template",
					Description: `(Optional) The name of the template to use.`,
				},
				resource.Attribute{
					Name:        "templates",
					Description: `(Optional) Templates configured for transformation.`,
				},
				resource.Attribute{
					Name:        "tweak_source",
					Description: `(Optional) The source of where the tweak value comes from. Only valid when in FPE mode.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) The type of transformation to perform.`,
				},
				resource.Attribute{
					Name:        "deletion_allowed",
					Description: `(Optional) If true, this transform can be deleted. Otherwise, deletion is blocked while this value remains false. Default: ` + "`" + `false` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vault_transit_secret_backend_key",
			Category:         "Resources",
			ShortDescription: `Create an Encryption Keyring on a Transit Secret Backend for Vault.`,
			Description:      ``,
			Keywords: []string{
				"transit",
				"secret",
				"backend",
				"key",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) The namespace to provision the resource in. The value should not contain leading or trailing forward slashes. The ` + "`" + `namespace` + "`" + ` is always relative to the provider's configured [namespace](/docs/providers/vault#namespace).`,
				},
				resource.Attribute{
					Name:        "backend",
					Description: `(Required) The path the transit secret backend is mounted at, with no leading or trailing ` + "`" + `/` + "`" + `s.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name to identify this key within the backend. Must be unique within the backend.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) Specifies the type of key to create. The currently-supported types are: ` + "`" + `aes128-gcm96` + "`" + `, ` + "`" + `aes256-gcm96` + "`" + ` (default), ` + "`" + `chacha20-poly1305` + "`" + `, ` + "`" + `ed25519` + "`" + `, ` + "`" + `ecdsa-p256` + "`" + `, ` + "`" + `ecdsa-p384` + "`" + `, ` + "`" + `ecdsa-p521` + "`" + `, ` + "`" + `rsa-2048` + "`" + `, ` + "`" + `rsa-3072` + "`" + ` and ` + "`" + `rsa-4096` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "deletion_allowed",
					Description: `(Optional) Specifies if the keyring is allowed to be deleted. Must be set to 'true' before terraform will be able to destroy keys.`,
				},
				resource.Attribute{
					Name:        "derived",
					Description: `(Optional) Specifies if key derivation is to be used. If enabled, all encrypt/decrypt requests to this key must provide a context which is used for key derivation.`,
				},
				resource.Attribute{
					Name:        "convergent_encryption",
					Description: `(Optional) Whether or not to support convergent encryption, where the same plaintext creates the same ciphertext. This requires ` + "`" + `derived` + "`" + ` to be set to ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "exportable",
					Description: `(Optional) Enables keys to be exportable. This allows for all valid private keys in the keyring to be exported. Once set, this cannot be disabled.`,
				},
				resource.Attribute{
					Name:        "allow_plaintext_backup",
					Description: `(Optional) Enables taking backup of entire keyring in the plaintext format. Once set, this cannot be disabled.`,
				},
				resource.Attribute{
					Name:        "min_decryption_version",
					Description: `(Optional) Minimum key version to use for decryption.`,
				},
				resource.Attribute{
					Name:        "min_encryption_version",
					Description: `(Optional) Minimum key version to use for encryption`,
				},
				resource.Attribute{
					Name:        "auto_rotate_period",
					Description: `(Optional) Amount of time the key should live before being automatically rotated. A value of 0 disables automatic rotation for the key. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "keys",
					Description: `List of key versions in the keyring. This attribute is zero-indexed and will contain a map of values depending on the ` + "`" + `type` + "`" + ` of the encryption key.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of keychain`,
				},
				resource.Attribute{
					Name:        "creation_time",
					Description: `ISO 8601 format timestamp indicating when the key version was created`,
				},
				resource.Attribute{
					Name:        "public_key",
					Description: `This is the base64-encoded public key for use outside of Vault.`,
				},
				resource.Attribute{
					Name:        "latest_version",
					Description: `Latest key version available. This value is 1-indexed, so if ` + "`" + `latest_version` + "`" + ` is ` + "`" + `1` + "`" + `, then the key's information can be referenced from ` + "`" + `keys` + "`" + ` by selecting element ` + "`" + `0` + "`" + ``,
				},
				resource.Attribute{
					Name:        "min_available_version",
					Description: `Minimum key version available for use. If keys have been archived by increasing ` + "`" + `min_decryption_version` + "`" + `, this attribute will reflect that change.`,
				},
				resource.Attribute{
					Name:        "supports_encryption",
					Description: `Whether or not the key supports encryption, based on key type.`,
				},
				resource.Attribute{
					Name:        "supports_decryption",
					Description: `Whether or not the key supports decryption, based on key type.`,
				},
				resource.Attribute{
					Name:        "supports_derivation",
					Description: `Whether or not the key supports derivation, based on key type.`,
				},
				resource.Attribute{
					Name:        "supports_signing",
					Description: `Whether or not the key supports signing, based on key type. ## Deprecations`,
				},
				resource.Attribute{
					Name:        "auto_rotate_interval",
					Description: `Replaced by ` + "`" + `auto_rotate_period` + "`" + `. ## Import Transit secret backend keys can be imported using the ` + "`" + `path` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_transit_secret_backend_key.key transit/keys/my_key ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "keys",
					Description: `List of key versions in the keyring. This attribute is zero-indexed and will contain a map of values depending on the ` + "`" + `type` + "`" + ` of the encryption key.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of keychain`,
				},
				resource.Attribute{
					Name:        "creation_time",
					Description: `ISO 8601 format timestamp indicating when the key version was created`,
				},
				resource.Attribute{
					Name:        "public_key",
					Description: `This is the base64-encoded public key for use outside of Vault.`,
				},
				resource.Attribute{
					Name:        "latest_version",
					Description: `Latest key version available. This value is 1-indexed, so if ` + "`" + `latest_version` + "`" + ` is ` + "`" + `1` + "`" + `, then the key's information can be referenced from ` + "`" + `keys` + "`" + ` by selecting element ` + "`" + `0` + "`" + ``,
				},
				resource.Attribute{
					Name:        "min_available_version",
					Description: `Minimum key version available for use. If keys have been archived by increasing ` + "`" + `min_decryption_version` + "`" + `, this attribute will reflect that change.`,
				},
				resource.Attribute{
					Name:        "supports_encryption",
					Description: `Whether or not the key supports encryption, based on key type.`,
				},
				resource.Attribute{
					Name:        "supports_decryption",
					Description: `Whether or not the key supports decryption, based on key type.`,
				},
				resource.Attribute{
					Name:        "supports_derivation",
					Description: `Whether or not the key supports derivation, based on key type.`,
				},
				resource.Attribute{
					Name:        "supports_signing",
					Description: `Whether or not the key supports signing, based on key type. ## Deprecations`,
				},
				resource.Attribute{
					Name:        "auto_rotate_interval",
					Description: `Replaced by ` + "`" + `auto_rotate_period` + "`" + `. ## Import Transit secret backend keys can be imported using the ` + "`" + `path` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_transit_secret_backend_key.key transit/keys/my_key ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
	}

	resourcesMap = map[string]int{

		"vault_ad_secret_backend":                            0,
		"vault_ad_secret_backend_library":                    1,
		"vault_ad_secret_role":                               2,
		"vault_alicloud_auth_backend_role":                   3,
		"vault_approle_auth_backend_login":                   4,
		"vault_approle_auth_backend_role":                    5,
		"vault_approle_auth_backend_role_secret_id":          6,
		"vault_audit":                                        7,
		"vault_audit_request_header":                         8,
		"vault_auth_backend":                                 9,
		"vault_aws_auth_backend_cert":                        10,
		"vault_aws_auth_backend_client":                      11,
		"vault_aws_auth_backend_config_identity":             12,
		"vault_aws_auth_backend_identity_whitelist":          13,
		"vault_aws_auth_backend_login":                       14,
		"vault_aws_auth_backend_role":                        15,
		"vault_aws_auth_backend_role_tag":                    16,
		"vault_aws_auth_backend_roletag_blacklist":           17,
		"vault_aws_auth_backend_sts_role":                    18,
		"vault_aws_secret_backend":                           19,
		"vault_aws_secret_backend_role":                      20,
		"vault_azure_auth_backend_config":                    21,
		"vault_azure_auth_backend_role":                      22,
		"vault_azure_secret_backend":                         23,
		"vault_azure_secret_backend_role":                    24,
		"vault_cert_auth_backend_role":                       25,
		"vault_consul_secret_backend":                        26,
		"vault_consul_secret_backend_role":                   27,
		"vault_database_secret_backend_connection":           28,
		"vault_database_secret_backend_role":                 29,
		"vault_database_secret_backend_static_role":          30,
		"vault_database_secrets_mount":                       31,
		"vault_egp_policy":                                   32,
		"vault_gcp_auth_backend":                             33,
		"vault_gcp_auth_backend_role":                        34,
		"vault_gcp_secret_backend":                           35,
		"vault_gcp_secret_roleset":                           36,
		"vault_gcp_secret_static_account":                    37,
		"vault_generic_endpoint":                             38,
		"vault_generic_secret":                               39,
		"vault_github_auth_backend":                          40,
		"vault_github_team":                                  41,
		"vault_github_user":                                  42,
		"vault_identity_entity":                              43,
		"vault_identity_entity_alias":                        44,
		"vault_identity_entity_policies":                     45,
		"vault_identity_group":                               46,
		"vault_identity_group_alias":                         47,
		"vault_identity_group_member_entity_ids":             48,
		"vault_identity_group_member_group_ids":              49,
		"vault_identity_group_policies":                      50,
		"vault_identity_mfa_duo":                             51,
		"vault_identity_mfa_login_enforcement":               52,
		"vault_identity_mfa_okta":                            53,
		"vault_identity_mfa_pingid":                          54,
		"vault_identity_mfa_totp":                            55,
		"vault_identity_oidc":                                56,
		"vault_identity_oidc_assignment":                     57,
		"vault_identity_oidc_client":                         58,
		"vault_identity_oidc_key":                            59,
		"vault_identity_oidc_key_allowed_client_id":          60,
		"vault_identity_oidc_provider":                       61,
		"vault_identity_oidc_role":                           62,
		"vault_identity_oidc_scope":                          63,
		"vault_jwt_auth_backend":                             64,
		"vault_jwt_auth_backend_role":                        65,
		"vault_kmip_secret_backend":                          66,
		"vault_kmip_secret_role":                             67,
		"vault_kmip_secret_scope":                            68,
		"vault_kubernetes_auth_backend_config":               69,
		"vault_kubernetes_auth_backend_role":                 70,
		"vault_kubernetes_secret_backend":                    71,
		"vault_kubernetes_secret_backend_role":               72,
		"vault_kv_secret":                                    73,
		"vault_kv_secret_backend_v2":                         74,
		"vault_kv_secret_v2":                                 75,
		"vault_ldap_auth_backend":                            76,
		"vault_ldap_auth_backend_group":                      77,
		"vault_ldap_auth_backend_user":                       78,
		"vault_managed_keys":                                 79,
		"vault_mfa_duo":                                      80,
		"vault_mfa_okta":                                     81,
		"vault_mfa_pingid":                                   82,
		"vault_mfa_totp":                                     83,
		"vault_mongodbatlas_secret_backend":                  84,
		"vault_mongodbatlas_secret_role":                     85,
		"vault_mount":                                        86,
		"vault_namespace":                                    87,
		"vault_okta_auth_backend":                            88,
		"vault_okta_auth_backend_group":                      89,
		"vault_okta_auth_backend_user":                       90,
		"vault_pki_secret_backend_cert":                      91,
		"vault_pki_secret_backend_config_ca":                 92,
		"vault_pki_secret_backend_config_urls":               93,
		"vault_pki_secret_backend_crl_config":                94,
		"vault_pki_secret_backend_intermediate_cert_request": 95,
		"vault_pki_secret_backend_intermediate_set_signed":   96,
		"vault_pki_secret_backend_role":                      97,
		"vault_pki_secret_backend_root_cert":                 98,
		"vault_pki_secret_backend_root_sign_intermediate":    99,
		"vault_pki_secret_backend_sign":                      100,
		"vault_policy":                                       101,
		"vault_quota_lease_count":                            102,
		"vault_quota_rate_limit":                             103,
		"vault_rabbitmq_secret_backend":                      104,
		"vault_rabbitmq_secret_backend_role":                 105,
		"vault_rgp_policy":                                   106,
		"vault_ssh_secret_backend_ca":                        107,
		"vault_ssh_secret_backend_role":                      108,
		"vault_token":                                        109,
		"vault_token_auth_backend_role":                      110,
		"vault_transform_alphabet":                           111,
		"vault_transform_role":                               112,
		"vault_transform_template":                           113,
		"vault_transform_transformation":                     114,
		"vault_transit_secret_backend_key":                   115,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
