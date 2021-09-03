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
					Name:        "backend",
					Description: `(Optional) The unique path this backend should be mounted at. Must not begin or end with a ` + "`" + `/` + "`" + `. Defaults to ` + "`" + `ad` + "`" + `.`,
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
					Description: `(Optional) Text to insert the password into, ex. "customPrefix{{PASSWORD}}customSuffix". This setting is deprecated and should instead use ` + "`" + `password_policy` + "`" + `.`,
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
					Description: `(Optional) The desired length of passwords that Vault generates. This setting is deprecated and should instead use ` + "`" + `password_policy` + "`" + `.`,
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
			Type:             "vault_ad_secret_backend_role",
			Category:         "Resources",
			ShortDescription: `Creates a role on the Active Directory Secret Backend for Vault.`,
			Description:      ``,
			Keywords: []string{
				"ad",
				"secret",
				"backend",
				"role",
			},
			Arguments: []resource.Attribute{
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
					Description: `Timestamp of the last password set by Vault. ## Import AD secret backend roles can be imported using the ` + "`" + `path` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_ad_secret_backend_role.role ad/roles/bob ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "last_vault_rotation",
					Description: `Timestamp of the last password rotation by Vault.`,
				},
				resource.Attribute{
					Name:        "password_last_set",
					Description: `Timestamp of the last password set by Vault. ## Import AD secret backend roles can be imported using the ` + "`" + `path` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_ad_secret_backend_role.role ad/roles/bob ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `(Optional) The [period](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls), if any, in number of seconds to set on the token.`,
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
					Description: `(Optional) The [period](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls), if any, in number of seconds to set on the token.`,
				},
				resource.Attribute{
					Name:        "token_type",
					Description: `(Optional) The type of token that should be generated. Can be ` + "`" + `service` + "`" + `, ` + "`" + `batch` + "`" + `, or ` + "`" + `default` + "`" + ` to use the mount's tuned default (which unless changed will be ` + "`" + `service` + "`" + ` tokens). For token store roles, there are two additional possibilities: ` + "`" + `default-service` + "`" + ` and ` + "`" + `default-batch` + "`" + ` which specify the type to return unless the client requests a different type at generation time. ### Deprecated Arguments`,
				},
				resource.Attribute{
					Name:        "bound_cidr_list",
					Description: `(Optional; Deprecated, use ` + "`" + `secret_id_bound_cidrs` + "`" + ` instead) If set, specifies blocks of IP addresses which can perform the login operation. These arguments are deprecated since Vault 1.2 in favour of the common token arguments documented above.`,
				},
				resource.Attribute{
					Name:        "policies",
					Description: `(Optional; Deprecated, use ` + "`" + `token_policies` + "`" + ` instead if you are running Vault >= 1.2) An array of strings specifying the policies to be set on tokens issued using this role.`,
				},
				resource.Attribute{
					Name:        "period",
					Description: `(Optional; Deprecated, use ` + "`" + `token_period` + "`" + ` instead if you are running Vault >= 1.2) If set, indicates that the token generated using this role should never expire. The token should be renewed within the duration specified by this value. At each renewal, the token's TTL will be set to the value of this field. Specified in seconds. ## Attributes Reference No additional attributes are exported by this resource. ## Import AppRole authentication backend roles can be imported using the ` + "`" + `path` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_approle_auth_backend_role.example auth/approle/role/test-role ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `(Optional) If set, the SecretID response will be [response-wrapped](https://www.vaultproject.io/docs/concepts/response-wrapping) and available for the duration specified. Only a single unwrapping of the token is allowed. ## Attributes Reference In addition to the fields above, the following attributes are exported:`,
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
					Name:        "type",
					Description: `(Required) The name of the auth method type`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Optional) The path to mount the auth method — this defaults to the name of the type`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A description of the auth method`,
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
					Description: `The accessor for this auth method ### Deprecated Arguments These arguments are deprecated since version 1.8 of the provider in favour of the ` + "`" + `tune` + "`" + ` block arguments documented above.`,
				},
				resource.Attribute{
					Name:        "default_lease_ttl_seconds",
					Description: `(Optional; Deprecated, use ` + "`" + `tune.default_lease_ttl` + "`" + ` if you are using Vault provider version >= 1.8) The default lease duration in seconds.`,
				},
				resource.Attribute{
					Name:        "max_lease_ttl_seconds",
					Description: `(Optional; Deprecated, use ` + "`" + `tune.max_lease_ttl` + "`" + ` if you are using Vault provider version >= 1.8) The maximum lease duration in seconds.`,
				},
				resource.Attribute{
					Name:        "listing_visibility",
					Description: `(Optional; Deprecated, use ` + "`" + `tune.listing_visibility` + "`" + ` if you are using Vault provider version >= 1.8) Speficies whether to show this mount in the UI-specific listing endpoint. ## Import Auth methods can be imported using the ` + "`" + `path` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_auth_backend.example github ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "accessor",
					Description: `The accessor for this auth method ### Deprecated Arguments These arguments are deprecated since version 1.8 of the provider in favour of the ` + "`" + `tune` + "`" + ` block arguments documented above.`,
				},
				resource.Attribute{
					Name:        "default_lease_ttl_seconds",
					Description: `(Optional; Deprecated, use ` + "`" + `tune.default_lease_ttl` + "`" + ` if you are using Vault provider version >= 1.8) The default lease duration in seconds.`,
				},
				resource.Attribute{
					Name:        "max_lease_ttl_seconds",
					Description: `(Optional; Deprecated, use ` + "`" + `tune.max_lease_ttl` + "`" + ` if you are using Vault provider version >= 1.8) The maximum lease duration in seconds.`,
				},
				resource.Attribute{
					Name:        "listing_visibility",
					Description: `(Optional; Deprecated, use ` + "`" + `tune.listing_visibility` + "`" + ` if you are using Vault provider version >= 1.8) Speficies whether to show this mount in the UI-specific listing endpoint. ## Import Auth methods can be imported using the ` + "`" + `path` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_auth_backend.example github ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `(Optional) The [period](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls), if any, in number of seconds to set on the token.`,
				},
				resource.Attribute{
					Name:        "token_type",
					Description: `(Optional) The type of token that should be generated. Can be ` + "`" + `service` + "`" + `, ` + "`" + `batch` + "`" + `, or ` + "`" + `default` + "`" + ` to use the mount's tuned default (which unless changed will be ` + "`" + `service` + "`" + ` tokens). For token store roles, there are two additional possibilities: ` + "`" + `default-service` + "`" + ` and ` + "`" + `default-batch` + "`" + ` which specify the type to return unless the client requests a different type at generation time. ### Deprecated Arguments These arguments are deprecated since Vault 1.2 in favour of the common token arguments documented above.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `(Optional; Deprecated, use ` + "`" + `token_ttl` + "`" + ` instead if you are running Vault >= 1.2) The TTL period of tokens issued using this role, provided as a number of seconds.`,
				},
				resource.Attribute{
					Name:        "max_ttl",
					Description: `(Optional; Deprecated, use ` + "`" + `token_max_ttl` + "`" + ` instead if you are running Vault >= 1.2) The maximum allowed lifetime of tokens issued using this role, provided as a number of seconds.`,
				},
				resource.Attribute{
					Name:        "policies",
					Description: `(Optional; Deprecated, use ` + "`" + `token_policies` + "`" + ` instead if you are running Vault >= 1.2) An array of strings specifying the policies to be set on tokens issued using this role.`,
				},
				resource.Attribute{
					Name:        "period",
					Description: `(Optional; Deprecated, use ` + "`" + `token_period` + "`" + ` instead if you are running Vault >= 1.2) If set, indicates that the token generated using this role should never expire. The token should be renewed within the duration specified by this value. At each renewal, the token's TTL will be set to the value of this field. Specified in seconds. ## Attributes Reference No additional attributes are exported by this resource. ## Import AWS auth backend roles can be imported using ` + "`" + `auth/` + "`" + `, the ` + "`" + `backend` + "`" + ` path, ` + "`" + `/role/` + "`" + `, and the ` + "`" + `role` + "`" + ` name e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_aws_auth_backend_role.example auth/aws/role/test-role ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
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
					Name:        "backend",
					Description: `(Required) The path the AWS auth backend being configured was mounted at.`,
				},
				resource.Attribute{
					Name:        "safety_buffer",
					Description: `(Oprtional) The amount of extra time that must have passed beyond the roletag expiration, before it is removed from the backend storage. Defaults to 259,200 seconds, or 72 hours.`,
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
					Description: `(Optional) Specifies a custom HTTP STS endpoint to use. ## Attributes Reference No additional attributes are exported by this resource. ## Import AWS secret backends can be imported using the ` + "`" + `path` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_aws_secret_backend.aws aws ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `(Optional) The max allowed TTL in seconds for STS credentials (credentials TTL are capped to ` + "`" + `max_sts_ttl` + "`" + `). Valid only when ` + "`" + `credential_type` + "`" + ` is one of ` + "`" + `assumed_role` + "`" + ` or ` + "`" + `federation_token` + "`" + `. ## Attributes Reference No additional attributes are exported by this resource. ## Import AWS secret backend roles can be imported using the ` + "`" + `path` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_aws_secret_backend_role.role aws/roles/deploy ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `(Optional) If set, defines a constraint on the virtual machiness that can perform the login operation that they be associated with the resource group that matches the value specified by this field.`,
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
					Description: `(Optional) The [period](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls), if any, in number of seconds to set on the token.`,
				},
				resource.Attribute{
					Name:        "token_type",
					Description: `(Optional) The type of token that should be generated. Can be ` + "`" + `service` + "`" + `, ` + "`" + `batch` + "`" + `, or ` + "`" + `default` + "`" + ` to use the mount's tuned default (which unless changed will be ` + "`" + `service` + "`" + ` tokens). For token store roles, there are two additional possibilities: ` + "`" + `default-service` + "`" + ` and ` + "`" + `default-batch` + "`" + ` which specify the type to return unless the client requests a different type at generation time. ### Deprecated Arguments These arguments are deprecated since Vault 1.2 in favour of the common token arguments documented above.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `(Optional; Deprecated, use ` + "`" + `token_ttl` + "`" + ` instead if you are running Vault >= 1.2) The TTL period of tokens issued using this role, provided as a number of seconds.`,
				},
				resource.Attribute{
					Name:        "max_ttl",
					Description: `(Optional; Deprecated, use ` + "`" + `token_max_ttl` + "`" + ` instead if you are running Vault >= 1.2) The maximum allowed lifetime of tokens issued using this role, provided as a number of seconds.`,
				},
				resource.Attribute{
					Name:        "policies",
					Description: `(Optional; Deprecated, use ` + "`" + `token_policies` + "`" + ` instead if you are running Vault >= 1.2) An array of strings specifying the policies to be set on tokens issued using this role.`,
				},
				resource.Attribute{
					Name:        "period",
					Description: `(Optional; Deprecated, use ` + "`" + `token_period` + "`" + ` instead if you are running Vault >= 1.2) If set, indicates that the token generated using this role should never expire. The token should be renewed within the duration specified by this value. At each renewal, the token's TTL will be set to the value of this field. Specified in seconds. ## Attributes Reference No additional attributes are exported by this resource. ## Import Azure auth backend roles can be imported using ` + "`" + `auth/` + "`" + `, the ` + "`" + `backend` + "`" + ` path, ` + "`" + `/role/` + "`" + `, and the ` + "`" + `role` + "`" + ` name e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_azure_auth_backend_role.example auth/azure/role/test-role ` + "`" + `` + "`" + `` + "`" + ``,
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
					Name:        "allowed_organization_units",
					Description: `(Optional) Allowed organization units for authenticated client certificates`,
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
					Description: `(Optional) The number of times issued tokens can be used. A value of 0 means unlimited uses.`,
				},
				resource.Attribute{
					Name:        "token_num_uses",
					Description: `(Optional) The [period](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls), if any, in number of seconds to set on the token.`,
				},
				resource.Attribute{
					Name:        "token_type",
					Description: `(Optional) The type of token that should be generated. Can be ` + "`" + `service` + "`" + `, ` + "`" + `batch` + "`" + `, or ` + "`" + `default` + "`" + ` to use the mount's tuned default (which unless changed will be ` + "`" + `service` + "`" + ` tokens). For token store roles, there are two additional possibilities: ` + "`" + `default-service` + "`" + ` and ` + "`" + `default-batch` + "`" + ` which specify the type to return unless the client requests a different type at generation time. ### Deprecated Arguments These arguments are deprecated since Vault 1.2 in favour of the common token arguments documented above.`,
				},
				resource.Attribute{
					Name:        "bound_cidrs",
					Description: `(Optional; Deprecated, use ` + "`" + `token_bound_cidrs` + "`" + ` instead if you are running Vault >= 1.2) Restriction usage of the certificates to client IPs falling within the range of the specified CIDRs`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `(Optional; Deprecated, use ` + "`" + `token_ttl` + "`" + ` instead if you are running Vault >= 1.2) The TTL period of tokens issued using this role, provided as a number of seconds.`,
				},
				resource.Attribute{
					Name:        "max_ttl",
					Description: `(Optional; Deprecated, use ` + "`" + `token_max_ttl` + "`" + ` instead if you are running Vault >= 1.2) The maximum allowed lifetime of tokens issued using this role, provided as a number of seconds.`,
				},
				resource.Attribute{
					Name:        "policies",
					Description: `(Optional; Deprecated, use ` + "`" + `token_policies` + "`" + ` instead if you are running Vault >= 1.2) An array of strings specifying the policies to be set on tokens issued using this role.`,
				},
				resource.Attribute{
					Name:        "period",
					Description: `(Optional; Deprecated, use ` + "`" + `token_period` + "`" + ` instead if you are running Vault >= 1.2) If set, indicates that the token generated using this role should never expire. The token should be renewed within the duration specified by this value. At each renewal, the token's TTL will be set to the value of this field. Specified in seconds. For more details on the usage of each argument consult the [Vault Cert API documentation](https://www.vaultproject.io/api-docs/auth/cert). ## Attribute Reference No additional attributes are exposed by this resource.`,
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
					Name:        "token",
					Description: `(Required) The Consul management token this backend should use to issue new tokens. ~>`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Optional) The unique location this backend should be mounted at. Must not begin or end with a ` + "`" + `/` + "`" + `. Defaults to ` + "`" + `consul` + "`" + `.`,
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
					Name:        "path",
					Description: `(Optional) The unique name of an existing Consul secrets backend mount. Must not begin or end with a ` + "`" + `/` + "`" + `.`,
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
					Description: `(Required) The list of Consul ACL policies to associate with these roles.`,
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
					Description: `(Optional) Indicates that the token should not be replicated globally and instead be local to the current datacenter. ## Attributes Reference No additional attributes are exported by this resource. ## Import Consul secret backend roles can be imported using the ` + "`" + `backend` + "`" + `, ` + "`" + `/roles/` + "`" + `, and the ` + "`" + `name` + "`" + ` e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_consul_secret_backend_role.example consul/roles/my-role ` + "`" + `` + "`" + `` + "`" + ``,
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
					Name:        "name",
					Description: `(Required) A unique name to give the database connection.`,
				},
				resource.Attribute{
					Name:        "backend",
					Description: `(Required) The unique name of the Vault mount to configure.`,
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
					Description: `(Optional) A nested block containing configuration options for Snowflake connections. Exactly one of the nested blocks of configuration options must be supplied. ### Cassandra Configuration Options`,
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
					Description: `(Optional) The number of seconds to use as a connection timeout. ### MongoDB Configuration Options`,
				},
				resource.Attribute{
					Name:        "connection_url",
					Description: `(Required) A URL containing connection information. See the [Vault docs](https://www.vaultproject.io/api-docs/secret/databases/mongodb.html#sample-payload) for an example.`,
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
					Description: `(Optional) The maximum number of seconds to keep a connection alive for. ### MSSQL Configuration Options`,
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
					Description: `(Optional) For Vault v1.7+. The template to use for username generation. See the [Vault docs](https://www.vaultproject.io/docs/concepts/username-templating) ### MySQL Configuration Options`,
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
					Description: `(Required) The password to be used in the connection. ### Snowflake Configuration Options`,
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
					Name:        "credentials",
					Description: `A JSON string containing the contents of a GCP credentials file. If this value is empty, Vault will try to use Application Default Credentials from the machine on which the Vault server is running.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Optional) The path to mount the auth method — this defaults to 'gcp'.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A description of the auth method.`,
				},
				resource.Attribute{
					Name:        "local",
					Description: `(Optional) Specifies if the auth method is local only. For more details on the usage of each argument consult the [Vault GCP API documentation](https://www.vaultproject.io/api-docs/auth/gcp#configure). ## Attribute Reference In addition to the fields above, the following attributes are also exposed:`,
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
					Name:        "role",
					Description: `(Required) Name of the GCP role`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) Type of GCP authentication role (either ` + "`" + `gce` + "`" + ` or ` + "`" + `iam` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Optional; Deprecated, use ` + "`" + `bound_projects` + "`" + ` instead) GCP Project that the role exists within`,
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
					Description: `(Optional) The [period](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls), if any, in number of seconds to set on the token.`,
				},
				resource.Attribute{
					Name:        "token_type",
					Description: `(Optional) The type of token that should be generated. Can be ` + "`" + `service` + "`" + `, ` + "`" + `batch` + "`" + `, or ` + "`" + `default` + "`" + ` to use the mount's tuned default (which unless changed will be ` + "`" + `service` + "`" + ` tokens). For token store roles, there are two additional possibilities: ` + "`" + `default-service` + "`" + ` and ` + "`" + `default-batch` + "`" + ` which specify the type to return unless the client requests a different type at generation time. ### Deprecated Arguments These arguments are deprecated since Vault 1.2 in favour of the common token arguments documented above.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `(Optional; Deprecated, use ` + "`" + `token_ttl` + "`" + ` instead if you are running Vault >= 1.2) The TTL period of tokens issued using this role, provided as a number of seconds.`,
				},
				resource.Attribute{
					Name:        "max_ttl",
					Description: `(Optional; Deprecated, use ` + "`" + `token_max_ttl` + "`" + ` instead if you are running Vault >= 1.2) The maximum allowed lifetime of tokens issued using this role, provided as a number of seconds.`,
				},
				resource.Attribute{
					Name:        "policies",
					Description: `(Optional; Deprecated, use ` + "`" + `token_policies` + "`" + ` instead if you are running Vault >= 1.2) An array of strings specifying the policies to be set on tokens issued using this role.`,
				},
				resource.Attribute{
					Name:        "period",
					Description: `(Optional; Deprecated, use ` + "`" + `token_period` + "`" + ` instead if you are running Vault >= 1.2) If set, indicates that the token generated using this role should never expire. The token should be renewed within the duration specified by this value. At each renewal, the token's TTL will be set to the value of this field. Specified in seconds. ## Attribute Reference No additional attributes are exposed by this resource. ## Import GCP authentication roles can be imported using the ` + "`" + `path` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_gcp_auth_backend_role.my_role auth/gcp/role/my_role ` + "`" + `` + "`" + `` + "`" + ``,
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
					Name:        "credentials",
					Description: `(Optional) The GCP service account credentials in JSON format. ~>`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Optional) The unique path this backend should be mounted at. Must not begin or end with a ` + "`" + `/` + "`" + `. Defaults to ` + "`" + `gcp` + "`" + `.`,
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
					Name:        "path",
					Description: `(Required) The full logical path at which to write the given data. To write data into the "generic" secret backend mounted in Vault by default, this should be prefixed with ` + "`" + `secret/` + "`" + `. Writing to other backends with this resource is possible; consult each backend's documentation to see which endpoints support the ` + "`" + `PUT` + "`" + ` and ` + "`" + `DELETE` + "`" + ` methods.`,
				},
				resource.Attribute{
					Name:        "data_json",
					Description: `(Required) String containing a JSON-encoded object that will be written as the secret data at the given path.`,
				},
				resource.Attribute{
					Name:        "allow_read",
					Description: `(Optional, Deprecated) True/false. Set this to true if your vault authentication is able to read the data, this allows the resource to be compared and updated. Defaults to false.`,
				},
				resource.Attribute{
					Name:        "disable_read",
					Description: `(Optional) True/false. Set this to true if your vault authentication is not able to read the data. Setting this to ` + "`" + `true` + "`" + ` will break drift detection. Defaults to false. ## Required Vault Capabilities Use of this resource requires the ` + "`" + `create` + "`" + ` or ` + "`" + `update` + "`" + ` capability (depending on whether the resource already exists) on the given path, the ` + "`" + `delete` + "`" + ` capability if the resource is removed from configuration, and the ` + "`" + `read` + "`" + ` capability for drift detection (by default). ### Drift Detection This resource does not necessarily need to`,
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
			ShortDescription: `Manages Github Auth mounts in Vault.`,
			Description:      ``,
			Keywords: []string{
				"github",
				"auth",
				"backend",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "path",
					Description: `(Optional) Path where the auth backend is mounted. Defaults to ` + "`" + `auth/github` + "`" + ` if not specified.`,
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
					Description: `(Optional) The number of times issued tokens can be used. A value of 0 means unlimited uses.`,
				},
				resource.Attribute{
					Name:        "token_num_uses",
					Description: `(Optional) The [period](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls), if any, in number of seconds to set on the token.`,
				},
				resource.Attribute{
					Name:        "token_type",
					Description: `(Optional) The type of token that should be generated. Can be ` + "`" + `service` + "`" + `, ` + "`" + `batch` + "`" + `, or ` + "`" + `default` + "`" + ` to use the mount's tuned default (which unless changed will be ` + "`" + `service` + "`" + ` tokens). For token store roles, there are two additional possibilities: ` + "`" + `default-service` + "`" + ` and ` + "`" + `default-batch` + "`" + ` which specify the type to return unless the client requests a different type at generation time. ### Deprecated Arguments These arguments are deprecated since Vault 1.2 in favour of the common token arguments documented above.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `(Optional; Deprecated, use ` + "`" + `token_ttl` + "`" + ` instead if you are running Vault >= 1.2) The TTL period of tokens issued using this role. This must be a valid [duration string](https://golang.org/pkg/time/#ParseDuration).`,
				},
				resource.Attribute{
					Name:        "max_ttl",
					Description: `(Optional; Deprecated, use ` + "`" + `token_max_ttl` + "`" + ` instead if you are running Vault >= 1.2) The maximum allowed lifetime of tokens issued using this role. This must be a valid [duration string](https://golang.org/pkg/time/#ParseDuration). ## Import Github authentication mounts can be imported using the ` + "`" + `path` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_github_auth_backend.example github ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `(Optional) The number of times issued tokens can be used. A value of 0 means unlimited uses.`,
				},
				resource.Attribute{
					Name:        "token_num_uses",
					Description: `(Optional) The [period](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls), if any, in number of seconds to set on the token.`,
				},
				resource.Attribute{
					Name:        "token_type",
					Description: `(Optional) The type of token that should be generated. Can be ` + "`" + `service` + "`" + `, ` + "`" + `batch` + "`" + `, or ` + "`" + `default` + "`" + ` to use the mount's tuned default (which unless changed will be ` + "`" + `service` + "`" + ` tokens). For token store roles, there are two additional possibilities: ` + "`" + `default-service` + "`" + ` and ` + "`" + `default-batch` + "`" + ` which specify the type to return unless the client requests a different type at generation time. ### Deprecated Arguments These arguments are deprecated since Vault 1.2 in favour of the common token arguments documented above.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `(Optional; Deprecated, use ` + "`" + `token_ttl` + "`" + ` instead if you are running Vault >= 1.2) The TTL period of tokens issued using this role. This must be a valid [duration string](https://golang.org/pkg/time/#ParseDuration).`,
				},
				resource.Attribute{
					Name:        "max_ttl",
					Description: `(Optional; Deprecated, use ` + "`" + `token_max_ttl` + "`" + ` instead if you are running Vault >= 1.2) The maximum allowed lifetime of tokens issued using this role. This must be a valid [duration string](https://golang.org/pkg/time/#ParseDuration). ## Import Github authentication mounts can be imported using the ` + "`" + `path` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_github_auth_backend.example github ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `(Optional) ` + "`" + `false` + "`" + ` by default. If set to ` + "`" + `true` + "`" + `, this resource will ignore any Entity IDs returned from Vault or specified in the resource. You can use [` + "`" + `vault_identity_group_member_entity_ids` + "`" + `](identity_group_member_entity_ids.html) to manage Entity IDs for this group in a decoupled manner. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
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
					Name:        "issuer",
					Description: `(Optional) Issuer URL to be used in the iss claim of the token. If not set, Vault's ` + "`" + `api_addr` + "`" + ` will be used. The issuer is a case sensitive URL using the https scheme that contains scheme, host, and optionally, port number and path components, but no query or fragment components. ## Attributes Reference No additional attributes are exposed by this resource.`,
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
					Name:        "path",
					Description: `(Required) Path to mount the JWT/OIDC auth backend`,
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
					Name:        "role_name",
					Description: `(Required) The name of the role.`,
				},
				resource.Attribute{
					Name:        "role_type",
					Description: `(Optional) Type of role, either "oidc" (default) or "jwt".`,
				},
				resource.Attribute{
					Name:        "bound_audiences",
					Description: `(Required for roles of type ` + "`" + `jwt` + "`" + `, optional for roles of type ` + "`" + `oidc` + "`" + `) List of ` + "`" + `aud` + "`" + ` claims to match against. Any match is sufficient.`,
				},
				resource.Attribute{
					Name:        "user_claim",
					Description: `(Required) The claim to use to uniquely identify the user; this will be used as the name for the Identity entity alias created due to a successful login.`,
				},
				resource.Attribute{
					Name:        "bound_subject",
					Description: `(Optional) If set, requires that the ` + "`" + `sub` + "`" + ` claim matches this value.`,
				},
				resource.Attribute{
					Name:        "bound_claims",
					Description: `(Optional) If set, a map of claims/values to match against. The expected value may be a single string or a list of strings.`,
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
					Name:        "groups_claim_delimiter_pattern",
					Description: `(Optional; Deprecated. This field has been removed since Vault 1.1. If the groups claim is not at the top level, it can now be specified as a [JSONPointer](https://tools.ietf.org/html/rfc6901).) A pattern of delimiters used to allow the groups_claim to live outside of the top-level JWT structure. For instance, a groups_claim of meta/user.name/groups with this field set to // will expect nested structures named meta, user.name, and groups. If this field was set to /./ the groups information would expect to be via nested structures of meta, user, name, and groups.`,
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
					Description: `(Optional) Log received OIDC tokens and claims when debug-level logging is active. Not recommended in production since sensitive information may be present in OIDC responses. ### Common Token Arguments These arguments are common across several Authentication Token resources since Vault 1.2.`,
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
					Description: `(Optional) The [period](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls), if any, in number of seconds to set on the token.`,
				},
				resource.Attribute{
					Name:        "token_type",
					Description: `(Optional) The type of token that should be generated. Can be ` + "`" + `service` + "`" + `, ` + "`" + `batch` + "`" + `, or ` + "`" + `default` + "`" + ` to use the mount's tuned default (which unless changed will be ` + "`" + `service` + "`" + ` tokens). For token store roles, there are two additional possibilities: ` + "`" + `default-service` + "`" + ` and ` + "`" + `default-batch` + "`" + ` which specify the type to return unless the client requests a different type at generation time. ### Deprecated Arguments These arguments are deprecated since Vault 1.2 in favour of the common token arguments documented above.`,
				},
				resource.Attribute{
					Name:        "num_uses",
					Description: `(Optional; Deprecated, use ` + "`" + `token_num_uses` + "`" + ` instead if you are running Vault >= 1.2) If set, puts a use-count limitation on the issued token.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `(Optional; Deprecated, use ` + "`" + `token_ttl` + "`" + ` instead if you are running Vault >= 1.2) The TTL period of tokens issued using this role, provided as a number of seconds.`,
				},
				resource.Attribute{
					Name:        "max_ttl",
					Description: `(Optional; Deprecated, use ` + "`" + `token_max_ttl` + "`" + ` instead if you are running Vault >= 1.2) The maximum allowed lifetime of tokens issued using this role, provided as a number of seconds.`,
				},
				resource.Attribute{
					Name:        "policies",
					Description: `(Optional; Deprecated, use ` + "`" + `token_policies` + "`" + ` instead if you are running Vault >= 1.2) An array of strings specifying the policies to be set on tokens issued using this role.`,
				},
				resource.Attribute{
					Name:        "period",
					Description: `(Optional; Deprecated, use ` + "`" + `token_period` + "`" + ` instead if you are running Vault >= 1.2) If set, indicates that the token generated using this role should never expire. The token should be renewed within the duration specified by this value. At each renewal, the token's TTL will be set to the value of this field. Specified in seconds.`,
				},
				resource.Attribute{
					Name:        "bound_cidrs",
					Description: `(Optional; Deprecated, use ` + "`" + `token_bound_cidrs` + "`" + ` instead if you are running Vault >= 1.2) If set, a list of CIDRs valid as the source address for login requests. This value is also encoded into any resulting token. ## Attributes Reference No additional attributes are exported by this resource. ## Import JWT authentication backend roles can be imported using the ` + "`" + `path` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_jwt_auth_backend_role.example auth/jwt/role/test-role ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `Optional JWT issuer. If no issuer is specified, ` + "`" + `kubernetes.io/serviceaccount` + "`" + ` will be used as the default issuer.`,
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
					Description: `(Optional) Audience claim to verify in the JWT. ### Common Token Arguments These arguments are common across several Authentication Token resources since Vault 1.2.`,
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
					Description: `(Optional) The [period](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls), if any, in number of seconds to set on the token.`,
				},
				resource.Attribute{
					Name:        "token_type",
					Description: `(Optional) The type of token that should be generated. Can be ` + "`" + `service` + "`" + `, ` + "`" + `batch` + "`" + `, or ` + "`" + `default` + "`" + ` to use the mount's tuned default (which unless changed will be ` + "`" + `service` + "`" + ` tokens). For token store roles, there are two additional possibilities: ` + "`" + `default-service` + "`" + ` and ` + "`" + `default-batch` + "`" + ` which specify the type to return unless the client requests a different type at generation time. ### Deprecated Arguments These arguments are deprecated since Vault 1.2 in favour of the common token arguments documented above.`,
				},
				resource.Attribute{
					Name:        "num_uses",
					Description: `(Optional; Deprecated, use ` + "`" + `token_num_uses` + "`" + ` instead if you are running Vault >= 1.2) If set, puts a use-count limitation on the issued token.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `(Optional; Deprecated, use ` + "`" + `token_ttl` + "`" + ` instead if you are running Vault >= 1.2) The TTL period of tokens issued using this role, provided as a number of seconds.`,
				},
				resource.Attribute{
					Name:        "max_ttl",
					Description: `(Optional; Deprecated, use ` + "`" + `token_max_ttl` + "`" + ` instead if you are running Vault >= 1.2) The maximum allowed lifetime of tokens issued using this role, provided as a number of seconds.`,
				},
				resource.Attribute{
					Name:        "policies",
					Description: `(Optional; Deprecated, use ` + "`" + `token_policies` + "`" + ` instead if you are running Vault >= 1.2) An array of strings specifying the policies to be set on tokens issued using this role.`,
				},
				resource.Attribute{
					Name:        "period",
					Description: `(Optional; Deprecated, use ` + "`" + `token_period` + "`" + ` instead if you are running Vault >= 1.2) If set, indicates that the token generated using this role should never expire. The token should be renewed within the duration specified by this value. At each renewal, the token's TTL will be set to the value of this field. Specified in seconds.`,
				},
				resource.Attribute{
					Name:        "bound_cidrs",
					Description: `(Optional; Deprecated, use ` + "`" + `token_bound_cidrs` + "`" + ` instead if you are running Vault >= 1.2) If set, a list of CIDRs valid as the source address for login requests. This value is also encoded into any resulting token. ## Attributes Reference No additional attributes are exported by this resource. ## Import Kubernetes auth backend role can be imported using the ` + "`" + `path` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_kubernetes_auth_backend_role.foo auth/kubernetes/role/foo ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
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
					Name:        "url",
					Description: `(Required) The URL of the LDAP server`,
				},
				resource.Attribute{
					Name:        "starttls",
					Description: `(Optional) Control use of TLS when conecting to LDAP`,
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
					Name:        "use_token_groups",
					Description: `(Optional) Use the Active Directory tokenGroups constructed attribute of the user to find the group memberships`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Optional) Path to mount the LDAP auth backend under`,
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
					Description: `(Optional) The number of times issued tokens can be used. A value of 0 means unlimited uses.`,
				},
				resource.Attribute{
					Name:        "token_num_uses",
					Description: `(Optional) The [period](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls), if any, in number of seconds to set on the token.`,
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
			Type:             "vault_mfa_duo",
			Category:         "Resources",
			ShortDescription: `Managing the MFA Duo method configuration`,
			Description:      ``,
			Keywords: []string{
				"mfa",
				"duo",
			},
			Arguments:  []resource.Attribute{},
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
					Description: `(Optional) Boolean flag that can be explicitly set to true to enable the secrets engine to access Vault's external entropy source ## Attributes Reference In addition to the fields above, the following attributes are exported:`,
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
					Name:        "path",
					Description: `(Required) The path of the namespace. Must not have a trailing ` + "`" + `/` + "`" + ` ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the namespace.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the namespace.`,
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
					Name:        "path",
					Description: `(Required) Path to mount the Okta auth backend`,
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
			Type:             "vault_pki_secret_backend",
			Category:         "Resources",
			ShortDescription: `Creates an PKI secret backend for Vault.`,
			Description:      ``,
			Keywords: []string{
				"pki",
				"secret",
				"backend",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "path",
					Description: `(Required) The unique path this backend should be mounted at. Must not begin or end with a ` + "`" + `/` + "`" + `.`,
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
					Description: `(Optional) The maximum TTL that can be requested for credentials issued by this backend. ## Attributes Reference No additional attributes are exported by this resource. ## Import PKI secret backends can be imported using the ` + "`" + `path` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_pki_secret_backend.pki pki ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `(Optional) Specifies the URL values for the OCSP Servers field. ## Attributes Reference No additional attributes are exported by this resource.`,
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
					Name:        "backend",
					Description: `(Required) The path the PKI secret backend is mounted at, with no leading or trailing ` + "`" + `/` + "`" + `s.`,
				},
				resource.Attribute{
					Name:        "expiry",
					Description: `(Optional) Specifies the time until expiration.`,
				},
				resource.Attribute{
					Name:        "disable",
					Description: `(Optional) Disables or enables CRL building. ## Attributes Reference No additional attributes are exported by this resource.`,
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
					Name:        "backend",
					Description: `(Required) The PKI secret backend the resource belongs to.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) Type of intermediate to create. Must be either \"exported\" or \"internal\"`,
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
					Description: `(Optional) The postal code ## Attributes Reference In addition to the fields above, the following attributes are exported:`,
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
					Name:        "backend",
					Description: `(Required) The PKI secret backend the resource belongs to.`,
				},
				resource.Attribute{
					Name:        "certificate",
					Description: `(Required) The certificate ## Attributes Reference No additional attributes are exported by this resource.`,
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
					Name:        "backend",
					Description: `(Required) The path the PKI secret backend is mounted at, with no leading or trailing ` + "`" + `/` + "`" + `s.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name to identify this role within the backend. Must be unique within the backend.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `(Optional) The TTL`,
				},
				resource.Attribute{
					Name:        "max_ttl",
					Description: `(Optional) The maximum TTL`,
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
					Description: `(Optional) The type of generated keys`,
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
					Description: `(Optional) Specify the list of allowed policies IODs`,
				},
				resource.Attribute{
					Name:        "basic_constraints_valid_for_non_ca",
					Description: `(Optional) Flag to mark basic constraints valid when issuing non-CA certificates`,
				},
				resource.Attribute{
					Name:        "not_before_duration",
					Description: `(Optional) Specifies the duration by which to backdate the NotBefore property. ## Attributes Reference No additional attributes are exported by this resource. ## Import PKI secret backend roles can be imported using the ` + "`" + `path` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_pki_secret_backend_role.role pki/roles/my_role ` + "`" + `` + "`" + `` + "`" + ``,
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
					Name:        "backend",
					Description: `(Required) The PKI secret backend the resource belongs to.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) Type of intermediate to create. Must be either \"exported\" or \"internal\"`,
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
					Description: `(Optional) The postal code ## Attributes Reference In addition to the fields above, the following attributes are exported:`,
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
					Name:        "serial",
					Description: `The serial`,
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
					Name:        "serial",
					Description: `The serial`,
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
					Description: `(Optional) The postal code ## Attributes Reference In addition to the fields above, the following attributes are exported:`,
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
					Name:        "serial",
					Description: `The serial`,
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
					Name:        "serial",
					Description: `The serial`,
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
					Description: `(Optional) List of alterative URIs`,
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
					Name:        "serial",
					Description: `The serial`,
				},
				resource.Attribute{
					Name:        "expiration",
					Description: `The expiration date of the certificate in unix epoch format`,
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
					Name:        "serial",
					Description: `The serial`,
				},
				resource.Attribute{
					Name:        "expiration",
					Description: `The expiration date of the certificate in unix epoch format`,
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
					Name:        "name",
					Description: `(Required) The name of the policy`,
				},
				resource.Attribute{
					Name:        "policy",
					Description: `(Required) String containing a Vault policy ## Attributes Reference No additional attributes are exported by this resource. ## Import Policies can be imported using the ` + "`" + `name` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_policy.example dev-team ` + "`" + `` + "`" + `` + "`" + ``,
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
					Name:        "name",
					Description: `(Required) Name of the rate limit quota`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Optional) Path of the mount or namespace to apply the quota. A blank path configures a global rate limit quota. For example ` + "`" + `namespace1/` + "`" + ` adds a quota to a full namespace, ` + "`" + `namespace1/auth/userpass` + "`" + ` adds a ` + "`" + `quota` + "`" + ` to ` + "`" + `userpass` + "`" + ` in ` + "`" + `namespace1` + "`" + `. Updating this field on an existing quota can have "moving" effects. For example, updating ` + "`" + `auth/userpass` + "`" + ` to ` + "`" + `namespace1/auth/userpass` + "`" + ` moves this quota from being a global mount quota to a namespace specific mount quota.`,
				},
				resource.Attribute{
					Name:        "rate",
					Description: `(Required) The maximum number of requests at any given second to be allowed by the quota rule. The ` + "`" + `rate` + "`" + ` must be positive. ## Attributes Reference No additional attributes are exported by this resource. ## Import Rate limit quotas can be imported using their names ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_quota_rate_limit.global global ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `(Optional) Specifies whether to verify connection URI, username, and password. Defaults to ` + "`" + `true` + "`" + `. ~>`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Optional) The unique path this backend should be mounted at. Must not begin or end with a ` + "`" + `/` + "`" + `. Defaults to ` + "`" + `rabbitmq` + "`" + `.`,
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
					Name:        "vhosts",
					Description: `(Optional) Specifies a map of virtual hosts to permissions. ## Attributes Reference No additional attributes are exported by this resource. ## Import RabbitMQ secret backend roles can be imported using the ` + "`" + `path` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_rabbitmq_secret_backend_role.role rabbitmq/roles/deploy ` + "`" + `` + "`" + `` + "`" + ``,
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
					Name:        "allowed_user_key_lengths",
					Description: `(Optional) Specifies a map of ssh key types and their expected sizes which are allowed to be signed by the CA type.`,
				},
				resource.Attribute{
					Name:        "max_ttl",
					Description: `(Optional) Specifies the maximum Time To Live value.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `(Optional) Specifies the Time To Live value. ## Attributes Reference No additional attributes are exposed by this resource. ## Import SSH secret backend roles can be imported using the ` + "`" + `path` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_ssh_secret_backend_role.foo ssh/roles/my-role ` + "`" + `` + "`" + `` + "`" + ``,
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
					Name:        "pgp_key",
					Description: `(Optional) The PGP key with which the ` + "`" + `client_token` + "`" + ` will be encrypted. The key must be provided using either a base64 encoded non-armored PGP key, or a keybase username in the form ` + "`" + `keybase:somebody` + "`" + `. The token won't be renewed automatically by the provider and ` + "`" + `client_token` + "`" + ` will be empty.`,
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
					Description: `String containing the client token if stored in present file`,
				},
				resource.Attribute{
					Name:        "encrypted_client_token",
					Description: `String containing the client token encrypted with the given ` + "`" + `pgp_key` + "`" + ` if stored in present file ## Import Tokens can be imported using its ` + "`" + `id` + "`" + ` as accessor id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_token.example <accessor_id> ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `String containing the client token if stored in present file`,
				},
				resource.Attribute{
					Name:        "encrypted_client_token",
					Description: `String containing the client token encrypted with the given ` + "`" + `pgp_key` + "`" + ` if stored in present file ## Import Tokens can be imported using its ` + "`" + `id` + "`" + ` as accessor id, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_token.example <accessor_id> ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `(Optional) The [period](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls), if any, in number of seconds to set on the token.`,
				},
				resource.Attribute{
					Name:        "token_type",
					Description: `(Optional) The type of token that should be generated. Can be ` + "`" + `service` + "`" + `, ` + "`" + `batch` + "`" + `, or ` + "`" + `default` + "`" + ` to use the mount's tuned default (which unless changed will be ` + "`" + `service` + "`" + ` tokens). For token store roles, there are two additional possibilities: ` + "`" + `default-service` + "`" + ` and ` + "`" + `default-batch` + "`" + ` which specify the type to return unless the client requests a different type at generation time. ### Deprecated Arguments These arguments are deprecated since Vault 1.2 in favour of the common token arguments documented above.`,
				},
				resource.Attribute{
					Name:        "period",
					Description: `(Optional; Deprecated, use ` + "`" + `token_period` + "`" + ` instead if you are running Vault >= 1.2) If set, indicates that the token generated using this role should never expire. The token should be renewed within the duration specified by this value. At each renewal, the token's TTL will be set to the value of this field. Specified in seconds.`,
				},
				resource.Attribute{
					Name:        "bound_cidrs",
					Description: `(Optional; Deprecated, use ` + "`" + `token_bound_cidrs` + "`" + ` instead if you are running Vault >= 1.2) If set, a list of CIDRs valid as the source address for login requests. This value is also encoded into any resulting token. ## Attributes Reference No additional attributes are exported by this resource. ## Import Token auth backend roles can be imported with ` + "`" + `auth/token/roles/` + "`" + ` followed by the ` + "`" + `role_name` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_token_auth_backend_role.example auth/token/roles/my-role ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `(Optional) Minimum key version to use for encryption ## Attributes Reference`,
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
					Description: `Whether or not the key supports signing, based on key type. ## Import Transit secret backend keys can be imported using the ` + "`" + `path` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_transit_secret_backend_key.key transit/keys/my_key ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `Whether or not the key supports signing, based on key type. ## Import Transit secret backend keys can be imported using the ` + "`" + `path` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import vault_transit_secret_backend_key.key transit/keys/my_key ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
	}

	resourcesMap = map[string]int{

		"vault_ad_secret_backend":                   0,
		"vault_ad_secret_backend_library":           1,
		"vault_ad_secret_backend_role":              2,
		"vault_alicloud_auth_backend_role":          3,
		"vault_approle_auth_backend_login":          4,
		"vault_approle_auth_backend_role":           5,
		"vault_approle_auth_backend_role_secret_id": 6,
		"vault_audit":                                        7,
		"vault_auth_backend":                                 8,
		"vault_aws_auth_backend_cert":                        9,
		"vault_aws_auth_backend_client":                      10,
		"vault_aws_auth_backend_identity_whitelist":          11,
		"vault_aws_auth_backend_login":                       12,
		"vault_aws_auth_backend_role":                        13,
		"vault_aws_auth_backend_role_tag":                    14,
		"vault_aws_auth_backend_roletag_blacklist":           15,
		"vault_aws_auth_backend_sts_role":                    16,
		"vault_aws_secret_backend":                           17,
		"vault_aws_secret_backend_role":                      18,
		"vault_azure_auth_backend_config":                    19,
		"vault_azure_auth_backend_role":                      20,
		"vault_azure_secret_backend":                         21,
		"vault_azure_secret_backend_role":                    22,
		"vault_cert_auth_backend_role":                       23,
		"vault_consul_secret_backend":                        24,
		"vault_consul_secret_backend_role":                   25,
		"vault_database_secret_backend_connection":           26,
		"vault_database_secret_backend_role":                 27,
		"vault_database_secret_backend_static_role":          28,
		"vault_egp_policy":                                   29,
		"vault_gcp_auth_backend":                             30,
		"vault_gcp_auth_backend_role":                        31,
		"vault_gcp_secret_backend":                           32,
		"vault_gcp_secret_roleset":                           33,
		"vault_gcp_secret_static_account":                    34,
		"vault_generic_endpoint":                             35,
		"vault_generic_secret":                               36,
		"vault_github_auth_backend":                          37,
		"vault_github_team":                                  38,
		"vault_github_user":                                  39,
		"vault_identity_entity":                              40,
		"vault_identity_entity_alias":                        41,
		"vault_identity_entity_policies":                     42,
		"vault_identity_group":                               43,
		"vault_identity_group_alias":                         44,
		"vault_identity_group_member_entity_ids":             45,
		"vault_identity_group_policies":                      46,
		"vault_identity_oidc":                                47,
		"vault_identity_oidc_key":                            48,
		"vault_identity_oidc_key_allowed_client_id":          49,
		"vault_identity_oidc_role":                           50,
		"vault_jwt_auth_backend":                             51,
		"vault_jwt_auth_backend_role":                        52,
		"vault_kubernetes_auth_backend_config":               53,
		"vault_kubernetes_auth_backend_role":                 54,
		"vault_ldap_auth_backend":                            55,
		"vault_ldap_auth_backend_group":                      56,
		"vault_ldap_auth_backend_user":                       57,
		"vault_mfa_duo":                                      58,
		"vault_mount":                                        59,
		"vault_namespace":                                    60,
		"vault_okta_auth_backend":                            61,
		"vault_okta_auth_backend_group":                      62,
		"vault_okta_auth_backend_user":                       63,
		"vault_pki_secret_backend":                           64,
		"vault_pki_secret_backend_cert":                      65,
		"vault_pki_secret_backend_config_ca":                 66,
		"vault_pki_secret_backend_config_urls":               67,
		"vault_pki_secret_backend_crl_config":                68,
		"vault_pki_secret_backend_intermediate_cert_request": 69,
		"vault_pki_secret_backend_intermediate_set_signed":   70,
		"vault_pki_secret_backend_role":                      71,
		"vault_pki_secret_backend_root_cert":                 72,
		"vault_pki_secret_backend_root_sign_intermediate":    73,
		"vault_pki_secret_backend_sign":                      74,
		"vault_policy":                                       75,
		"vault_quota_lease_count":                            76,
		"vault_quota_rate_limit":                             77,
		"vault_rabbitmq_secret_backend":                      78,
		"vault_rabbitmq_secret_backend_role":                 79,
		"vault_rgp_policy":                                   80,
		"vault_ssh_secret_backend_ca":                        81,
		"vault_ssh_secret_backend_role":                      82,
		"vault_token":                                        83,
		"vault_token_auth_backend_role":                      84,
		"vault_transit_secret_backend_key":                   85,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
