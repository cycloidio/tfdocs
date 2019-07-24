package fastly

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "fastly_service_v1",
			Category:         "Resources",
			ShortDescription: `Provides an Fastly Service`,
			Description:      ``,
			Keywords: []string{
				"service",
				"v1",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "activate",
					Description: `(Optional) Conditionally prevents the Service from being activated. The apply step will continue to create a new draft version but will not activate it if this is set to false. Default true.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The unique name for the Service to create.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `(Optional) Description field for the service. Default ` + "`" + `Managed by Terraform` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "version_comment",
					Description: `(Optional) Description field for the version.`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `(Required) A set of Domain names to serve as entry points for your Service. Defined below.`,
				},
				resource.Attribute{
					Name:        "backend",
					Description: `(Optional) A set of Backends to service requests from your Domains. Defined below. Backends must be defined in this argument, or defined in the ` + "`" + `vcl` + "`" + ` argument below`,
				},
				resource.Attribute{
					Name:        "condition",
					Description: `(Optional) A set of conditions to add logic to any basic configuration object in this service. Defined below.`,
				},
				resource.Attribute{
					Name:        "cache_setting",
					Description: `(Optional) A set of Cache Settings, allowing you to override`,
				},
				resource.Attribute{
					Name:        "director",
					Description: `(Optional) A director to allow more control over balancing traffic over backends. when an item is not to be cached based on an above ` + "`" + `condition` + "`" + `. Defined below`,
				},
				resource.Attribute{
					Name:        "gzip",
					Description: `(Required) A set of gzip rules to control automatic gzipping of content. Defined below.`,
				},
				resource.Attribute{
					Name:        "header",
					Description: `(Optional) A set of Headers to manipulate for each request. Defined below.`,
				},
				resource.Attribute{
					Name:        "healthcheck",
					Description: `(Optional) Automated healthchecks on the cache that can change how fastly interacts with the cache based on its health.`,
				},
				resource.Attribute{
					Name:        "default_host",
					Description: `(Optional) The default hostname.`,
				},
				resource.Attribute{
					Name:        "default_ttl",
					Description: `(Optional) The default Time-to-live (TTL) for requests.`,
				},
				resource.Attribute{
					Name:        "force_destroy",
					Description: `(Optional) Services that are active cannot be destroyed. In order to destroy the Service, set ` + "`" + `force_destroy` + "`" + ` to ` + "`" + `true` + "`" + `. Default ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "request_setting",
					Description: `(Optional) A set of Request modifiers. Defined below`,
				},
				resource.Attribute{
					Name:        "s3logging",
					Description: `(Optional) A set of S3 Buckets to send streaming logs too. Defined below.`,
				},
				resource.Attribute{
					Name:        "papertrail",
					Description: `(Optional) A Papertrail endpoint to send streaming logs too. Defined below.`,
				},
				resource.Attribute{
					Name:        "sumologic",
					Description: `(Optional) A Sumologic endpoint to send streaming logs too. Defined below.`,
				},
				resource.Attribute{
					Name:        "gcslogging",
					Description: `(Optional) A gcs endpoint to send streaming logs too. Defined below.`,
				},
				resource.Attribute{
					Name:        "bigquerylogging",
					Description: `(Optional) A BigQuery endpoint to send streaming logs too. Defined below.`,
				},
				resource.Attribute{
					Name:        "syslog",
					Description: `(Optional) A syslog endpoint to send streaming logs too. Defined below.`,
				},
				resource.Attribute{
					Name:        "logentries",
					Description: `(Optional) A logentries endpoint to send streaming logs too. Defined below.`,
				},
				resource.Attribute{
					Name:        "splunk",
					Description: `(Optional) A Splunk endpoint to send streaming logs too. Defined below.`,
				},
				resource.Attribute{
					Name:        "blobstoragelogging",
					Description: `(Optional) An Azure Blob Storage endpoint to send streaming logs too. Defined below.`,
				},
				resource.Attribute{
					Name:        "response_object",
					Description: `(Optional) Allows you to create synthetic responses that exist entirely on the varnish machine. Useful for creating error or maintenance pages that exists outside the scope of your datacenter. Best when used with Condition objects.`,
				},
				resource.Attribute{
					Name:        "snippet",
					Description: `(Optional) A set of custom, "regular" (non-dynamic) VCL Snippet configuration blocks. Defined below.`,
				},
				resource.Attribute{
					Name:        "vcl",
					Description: `(Optional) A set of custom VCL configuration blocks. The ability to upload custom VCL code is not enabled by default for new Fastly accounts (see the [Fastly documentation](https://docs.fastly.com/guides/vcl/uploading-custom-vcl) for details). The ` + "`" + `domain` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The domain to which this Service will respond.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `(Optional) An optional comment about the Domain. The ` + "`" + `backend` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required, string) Name for this Backend. Must be unique to this Service.`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `(Required, string) An IPv4, hostname, or IPv6 address for the Backend.`,
				},
				resource.Attribute{
					Name:        "auto_loadbalance",
					Description: `(Optional, boolean) Denotes if this Backend should be included in the pool of backends that requests are load balanced against. Default ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "between_bytes_timeout",
					Description: `(Optional) How long to wait between bytes in milliseconds. Default ` + "`" + `10000` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "connect_timeout",
					Description: `(Optional) How long to wait for a timeout in milliseconds. Default ` + "`" + `1000` + "`" + ``,
				},
				resource.Attribute{
					Name:        "error_threshold",
					Description: `(Optional) Number of errors to allow before the Backend is marked as down. Default ` + "`" + `0` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "first_byte_timeout",
					Description: `(Optional) How long to wait for the first bytes in milliseconds. Default ` + "`" + `15000` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "max_conn",
					Description: `(Optional) Maximum number of connections for this Backend. Default ` + "`" + `200` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional) The port number on which the Backend responds. Default ` + "`" + `80` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "override_host",
					Description: `(Optional) The hostname to override the Host header.`,
				},
				resource.Attribute{
					Name:        "request_condition",
					Description: `(Optional, string) Name of already defined ` + "`" + `condition` + "`" + `, which if met, will select this backend during a request.`,
				},
				resource.Attribute{
					Name:        "use_ssl",
					Description: `(Optional) Whether or not to use SSL to reach the backend. Default ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "max_tls_version",
					Description: `(Optional) Maximum allowed TLS version on SSL connections to this backend.`,
				},
				resource.Attribute{
					Name:        "min_tls_version",
					Description: `(Optional) Minimum allowed TLS version on SSL connections to this backend.`,
				},
				resource.Attribute{
					Name:        "ssl_ciphers",
					Description: `(Optional) Comma separated list of OpenSSL Ciphers to try when negotiating to the backend.`,
				},
				resource.Attribute{
					Name:        "ssl_ca_cert",
					Description: `(Optional) CA certificate attached to origin.`,
				},
				resource.Attribute{
					Name:        "ssl_client_cert",
					Description: `(Optional) Client certificate attached to origin. Used when connecting to the backend.`,
				},
				resource.Attribute{
					Name:        "ssl_client_key",
					Description: `(Optional) Client key attached to origin. Used when connecting to the backend.`,
				},
				resource.Attribute{
					Name:        "ssl_check_cert",
					Description: `(Optional) Be strict about checking SSL certs. Default ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ssl_hostname",
					Description: `(Optional, deprecated by Fastly) Used for both SNI during the TLS handshake and to validate the cert.`,
				},
				resource.Attribute{
					Name:        "ssl_cert_hostname",
					Description: `(Optional) Overrides ssl_hostname, but only for cert verification. Does not affect SNI at all.`,
				},
				resource.Attribute{
					Name:        "ssl_sni_hostname",
					Description: `(Optional) Overrides ssl_hostname, but only for SNI in the handshake. Does not affect cert validation at all.`,
				},
				resource.Attribute{
					Name:        "shield",
					Description: `(Optional) The POP of the shield designated to reduce inbound load.`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `(Optional) The [portion of traffic](https://docs.fastly.com/guides/performance-tuning/load-balancing-configuration.html#how-weight-affects-load-balancing) to send to this Backend. Each Backend receives ` + "`" + `weight / total` + "`" + ` of the traffic. Default ` + "`" + `100` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "healthcheck",
					Description: `(Optional) Name of a defined ` + "`" + `healthcheck` + "`" + ` to assign to this backend. The ` + "`" + `condition` + "`" + ` block supports allows you to add logic to any basic configuration object in a service. See Fastly's documentation ["About Conditions"](https://docs.fastly.com/guides/conditions/about-conditions) for more detailed information on using Conditions. The Condition ` + "`" + `name` + "`" + ` can be used in the ` + "`" + `request_condition` + "`" + `, ` + "`" + `response_condition` + "`" + `, or ` + "`" + `cache_condition` + "`" + ` attributes of other block settings.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The unique name for the condition.`,
				},
				resource.Attribute{
					Name:        "statement",
					Description: `(Required) The statement used to determine if the condition is met.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) Type of condition, either ` + "`" + `REQUEST` + "`" + ` (req), ` + "`" + `RESPONSE` + "`" + ` (req, resp), or ` + "`" + `CACHE` + "`" + ` (req, beresp).`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `(Optional) A number used to determine the order in which multiple conditions execute. Lower numbers execute first. Default ` + "`" + `10` + "`" + `. The ` + "`" + `director` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Unique name for this Director.`,
				},
				resource.Attribute{
					Name:        "backends",
					Description: `(Required) Names of defined backends to map the director to. Example: ` + "`" + `[ "origin1", "origin2" ]` + "`" + ``,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `(Optional) An optional comment about the Director.`,
				},
				resource.Attribute{
					Name:        "shield",
					Description: `(Optional) Selected POP to serve as a "shield" for origin servers.`,
				},
				resource.Attribute{
					Name:        "capacity",
					Description: `(Optional) Load balancing weight for the backends. Default ` + "`" + `100` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "quorum",
					Description: `(Optional) Percentage of capacity that needs to be up for the director itself to be considered up. Default ` + "`" + `75` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) Type of load balance group to use. Integer, 1 to 4. Values: ` + "`" + `1` + "`" + ` (random), ` + "`" + `3` + "`" + ` (hash), ` + "`" + `4` + "`" + ` (client). Default ` + "`" + `1` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "retries",
					Description: `(Optional) How many backends to search if it fails. Default ` + "`" + `5` + "`" + `. The ` + "`" + `cache_setting` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Unique name for this Cache Setting.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Optional) One of ` + "`" + `cache` + "`" + `, ` + "`" + `pass` + "`" + `, or ` + "`" + `restart` + "`" + `, as defined on Fastly's documentation under ["Caching action descriptions"](https://docs.fastly.com/guides/performance-tuning/controlling-caching#caching-action-descriptions).`,
				},
				resource.Attribute{
					Name:        "cache_condition",
					Description: `(Optional) Name of already defined ` + "`" + `condition` + "`" + ` used to test whether this settings object should be used. This ` + "`" + `condition` + "`" + ` must be of type ` + "`" + `CACHE` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "stale_ttl",
					Description: `(Optional) Max "Time To Live" for stale (unreachable) objects.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `(Optional) The Time-To-Live (TTL) for the object. The ` + "`" + `gzip` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A unique name.`,
				},
				resource.Attribute{
					Name:        "content_types",
					Description: `(Optional) The content-type for each type of content you wish to have dynamically gzip'ed. Example: ` + "`" + `["text/html", "text/css"]` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "extensions",
					Description: `(Optional) File extensions for each file type to dynamically gzip. Example: ` + "`" + `["css", "js"]` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "cache_condition",
					Description: `(Optional) Name of already defined ` + "`" + `condition` + "`" + ` controlling when this gzip configuration applies. This ` + "`" + `condition` + "`" + ` must be of type ` + "`" + `CACHE` + "`" + `. For detailed information about Conditionals, see [Fastly's Documentation on Conditionals][fastly-conditionals]. The ` + "`" + `Header` + "`" + ` block supports adding, removing, or modifying Request and Response headers. See Fastly's documentation on [Adding or modifying headers on HTTP requests and responses](https://docs.fastly.com/guides/basic-configuration/adding-or-modifying-headers-on-http-requests-and-responses#field-description-table) for more detailed information on any of the properties below.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Unique name for this header attribute.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Required) The Header manipulation action to take; must be one of ` + "`" + `set` + "`" + `, ` + "`" + `append` + "`" + `, ` + "`" + `delete` + "`" + `, ` + "`" + `regex` + "`" + `, or ` + "`" + `regex_repeat` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) The Request type on which to apply the selected Action; must be one of ` + "`" + `request` + "`" + `, ` + "`" + `fetch` + "`" + `, ` + "`" + `cache` + "`" + ` or ` + "`" + `response` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "destination",
					Description: `(Required) The name of the header that is going to be affected by the Action.`,
				},
				resource.Attribute{
					Name:        "ignore_if_set",
					Description: `(Optional) Do not add the header if it is already present. (Only applies to the ` + "`" + `set` + "`" + ` action.). Default ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "source",
					Description: `(Optional) Variable to be used as a source for the header content. (Does not apply to the ` + "`" + `delete` + "`" + ` action.)`,
				},
				resource.Attribute{
					Name:        "regex",
					Description: `(Optional) Regular expression to use (Only applies to the ` + "`" + `regex` + "`" + ` and ` + "`" + `regex_repeat` + "`" + ` actions.)`,
				},
				resource.Attribute{
					Name:        "substitution",
					Description: `(Optional) Value to substitute in place of regular expression. (Only applies to the ` + "`" + `regex` + "`" + ` and ` + "`" + `regex_repeat` + "`" + ` actions.)`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `(Optional) Lower priorities execute first. Default: ` + "`" + `100` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "request_condition",
					Description: `(Optional) Name of already defined ` + "`" + `condition` + "`" + ` to apply. This ` + "`" + `condition` + "`" + ` must be of type ` + "`" + `REQUEST` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "cache_condition",
					Description: `(Optional) Name of already defined ` + "`" + `condition` + "`" + ` to apply. This ` + "`" + `condition` + "`" + ` must be of type ` + "`" + `CACHE` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "response_condition",
					Description: `(Optional) Name of already defined ` + "`" + `condition` + "`" + ` to apply. This ` + "`" + `condition` + "`" + ` must be of type ` + "`" + `RESPONSE` + "`" + `. For detailed information about Conditionals, see [Fastly's Documentation on Conditionals][fastly-conditionals]. The ` + "`" + `healthcheck` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A unique name to identify this Healthcheck.`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `(Required) The Host header to send for this Healthcheck.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Required) The path to check.`,
				},
				resource.Attribute{
					Name:        "check_interval",
					Description: `(Optional) How often to run the Healthcheck in milliseconds. Default ` + "`" + `5000` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "expected_response",
					Description: `(Optional) The status code expected from the host. Default ` + "`" + `200` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "http_version",
					Description: `(Optional) Whether to use version 1.0 or 1.1 HTTP. Default ` + "`" + `1.1` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "initial",
					Description: `(Optional) When loading a config, the initial number of probes to be seen as OK. Default ` + "`" + `2` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "method",
					Description: `(Optional) Which HTTP method to use. Default ` + "`" + `HEAD` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "threshold",
					Description: `(Optional) How many Healthchecks must succeed to be considered healthy. Default ` + "`" + `3` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `(Optional) Timeout in milliseconds. Default ` + "`" + `500` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "window",
					Description: `(Optional) The number of most recent Healthcheck queries to keep for this Healthcheck. Default ` + "`" + `5` + "`" + `. The ` + "`" + `request_setting` + "`" + ` block allow you to customize Fastly's request handling, by defining behavior that should change based on a predefined ` + "`" + `condition` + "`" + `:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The domain for this request setting.`,
				},
				resource.Attribute{
					Name:        "request_condition",
					Description: `(Optional) Name of already defined ` + "`" + `condition` + "`" + ` to determine if this request setting should be applied.`,
				},
				resource.Attribute{
					Name:        "max_stale_age",
					Description: `(Optional) How old an object is allowed to be to serve ` + "`" + `stale-if-error` + "`" + ` or ` + "`" + `stale-while-revalidate` + "`" + `, in seconds.`,
				},
				resource.Attribute{
					Name:        "force_miss",
					Description: `(Optional) Force a cache miss for the request. If specified, can be ` + "`" + `true` + "`" + ` or ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "force_ssl",
					Description: `(Optional) Forces the request to use SSL (Redirects a non-SSL request to SSL).`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Optional) Allows you to terminate request handling and immediately perform an action. When set it can be ` + "`" + `lookup` + "`" + ` or ` + "`" + `pass` + "`" + ` (Ignore the cache completely).`,
				},
				resource.Attribute{
					Name:        "bypass_busy_wait",
					Description: `(Optional) Disable collapsed forwarding, so you don't wait for other objects to origin.`,
				},
				resource.Attribute{
					Name:        "hash_keys",
					Description: `(Optional) Comma separated list of varnish request object fields that should be in the hash key.`,
				},
				resource.Attribute{
					Name:        "xff",
					Description: `(Optional) X-Forwarded-For, should be ` + "`" + `clear` + "`" + `, ` + "`" + `leave` + "`" + `, ` + "`" + `append` + "`" + `, ` + "`" + `append_all` + "`" + `, or ` + "`" + `overwrite` + "`" + `. Default ` + "`" + `append` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "timer_support",
					Description: `(Optional) Injects the X-Timer info into the request for viewing origin fetch durations.`,
				},
				resource.Attribute{
					Name:        "geo_headers",
					Description: `(Optional) Injects Fastly-Geo-Country, Fastly-Geo-City, and Fastly-Geo-Region into the request headers.`,
				},
				resource.Attribute{
					Name:        "default_host",
					Description: `(Optional) Sets the host header. The ` + "`" + `s3logging` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A unique name to identify this S3 Logging Bucket.`,
				},
				resource.Attribute{
					Name:        "bucket_name",
					Description: `(Required) The name of the bucket in which to store the logs.`,
				},
				resource.Attribute{
					Name:        "s3_access_key",
					Description: `(Required) AWS Access Key of an account with the required permissions to post logs. It is`,
				},
				resource.Attribute{
					Name:        "s3_secret_key",
					Description: `(Required) AWS Secret Key of an account with the required permissions to post logs. It is`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Optional) Path to store the files. Must end with a trailing slash. If this field is left empty, the files will be saved in the bucket's root path.`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `(Optional) If you created the S3 bucket outside of ` + "`" + `us-east-1` + "`" + `, then specify the corresponding bucket endpoint. Example: ` + "`" + `s3-us-west-2.amazonaws.com` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "period",
					Description: `(Optional) How frequently the logs should be transferred, in seconds. Default ` + "`" + `3600` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "gzip_level",
					Description: `(Optional) Level of GZIP compression, from ` + "`" + `0-9` + "`" + `. ` + "`" + `0` + "`" + ` is no compression. ` + "`" + `1` + "`" + ` is fastest and least compressed, ` + "`" + `9` + "`" + ` is slowest and most compressed. Default ` + "`" + `0` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "format",
					Description: `(Optional) Apache-style string or VCL variables to use for log formatting. Defaults to Apache Common Log format (` + "`" + `%h %l %u %t %r %>s` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "format_version",
					Description: `(Optional) The version of the custom logging format used for the configured endpoint. Can be either 1 (the default, version 1 log format) or 2 (the version 2 log format).`,
				},
				resource.Attribute{
					Name:        "message_type",
					Description: `(Optional) How the message should be formatted; one of: ` + "`" + `classic` + "`" + `, ` + "`" + `loggly` + "`" + `, ` + "`" + `logplex` + "`" + ` or ` + "`" + `blank` + "`" + `. Default ` + "`" + `classic` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "timestamp_format",
					Description: `(Optional) ` + "`" + `strftime` + "`" + ` specified timestamp formatting (default ` + "`" + `%Y-%m-%dT%H:%M:%S.000` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "redundancy",
					Description: `(Optional) The S3 redundancy level. Should be formatted; one of: ` + "`" + `standard` + "`" + `, ` + "`" + `reduced_redundancy` + "`" + ` or null. Default ` + "`" + `null` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "response_condition",
					Description: `(Optional) Name of already defined ` + "`" + `condition` + "`" + ` to apply. This ` + "`" + `condition` + "`" + ` must be of type ` + "`" + `RESPONSE` + "`" + `. For detailed information about Conditionals, see [Fastly's Documentation on Conditionals][fastly-conditionals].`,
				},
				resource.Attribute{
					Name:        "placement",
					Description: `(Optional) Where in the generated VCL the logging call should be placed; one of: ` + "`" + `none` + "`" + ` or ` + "`" + `waf_debug` + "`" + `. The ` + "`" + `papertrail` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A unique name to identify this Papertrail endpoint.`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `(Required) The address of the Papertrail endpoint.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Required) The port associated with the address where the Papertrail endpoint can be accessed.`,
				},
				resource.Attribute{
					Name:        "format",
					Description: `(Optional) Apache-style string or VCL variables to use for log formatting. Defaults to Apache Common Log format (` + "`" + `%h %l %u %t %r %>s` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "response_condition",
					Description: `(Optional) Name of already defined ` + "`" + `condition` + "`" + ` to apply. This ` + "`" + `condition` + "`" + ` must be of type ` + "`" + `RESPONSE` + "`" + `. For detailed information about Conditionals, see [Fastly's Documentation on Conditionals][fastly-conditionals].`,
				},
				resource.Attribute{
					Name:        "placement",
					Description: `(Optional) Where in the generated VCL the logging call should be placed; one of: ` + "`" + `none` + "`" + ` or ` + "`" + `waf_debug` + "`" + `. The ` + "`" + `sumologic` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A unique name to identify this Sumologic endpoint.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Required) The URL to Sumologic collector endpoint`,
				},
				resource.Attribute{
					Name:        "format",
					Description: `(Optional) Apache-style string or VCL variables to use for log formatting. Defaults to Apache Common Log format (` + "`" + `%h %l %u %t %r %>s` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "format_version",
					Description: `(Optional) The version of the custom logging format used for the configured endpoint. Can be either 1 (the default, version 1 log format) or 2 (the version 2 log format).`,
				},
				resource.Attribute{
					Name:        "response_condition",
					Description: `(Optional) Name of already defined ` + "`" + `condition` + "`" + ` to apply. This ` + "`" + `condition` + "`" + ` must be of type ` + "`" + `RESPONSE` + "`" + `. For detailed information about Conditionals, see [Fastly's Documentation on Conditionals][fastly-conditionals].`,
				},
				resource.Attribute{
					Name:        "message_type",
					Description: `(Optional) How the message should be formatted; one of: ` + "`" + `classic` + "`" + `, ` + "`" + `loggly` + "`" + `, ` + "`" + `logplex` + "`" + ` or ` + "`" + `blank` + "`" + `. Default ` + "`" + `classic` + "`" + `. See [Fastly's Documentation on Sumologic][fastly-sumologic]`,
				},
				resource.Attribute{
					Name:        "placement",
					Description: `(Optional) Where in the generated VCL the logging call should be placed; one of: ` + "`" + `none` + "`" + ` or ` + "`" + `waf_debug` + "`" + `. The ` + "`" + `gcslogging` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A unique name to identify this GCS endpoint.`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `(Required) The email address associated with the target GCS bucket on your account. You may optionally provide this secret via an environment variable, ` + "`" + `FASTLY_GCS_EMAIL` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "bucket_name",
					Description: `(Required) The name of the bucket in which to store the logs.`,
				},
				resource.Attribute{
					Name:        "secret_key",
					Description: `(Required) The secret key associated with the target gcs bucket on your account. You may optionally provide this secret via an environment variable, ` + "`" + `FASTLY_GCS_SECRET_KEY` + "`" + `. A typical format for the key is PEM format, containing actual newline characters where required.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Optional) Path to store the files. Must end with a trailing slash. If this field is left empty, the files will be saved in the bucket's root path.`,
				},
				resource.Attribute{
					Name:        "period",
					Description: `(Optional) How frequently the logs should be transferred, in seconds. Default ` + "`" + `3600` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "gzip_level",
					Description: `(Optional) Level of GZIP compression, from ` + "`" + `0-9` + "`" + `. ` + "`" + `0` + "`" + ` is no compression. ` + "`" + `1` + "`" + ` is fastest and least compressed, ` + "`" + `9` + "`" + ` is slowest and most compressed. Default ` + "`" + `0` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "format",
					Description: `(Optional) Apache-style string or VCL variables to use for log formatting. Defaults to Apache Common Log format (` + "`" + `%h %l %u %t %r %>s` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "response_condition",
					Description: `(Optional) Name of already defined ` + "`" + `condition` + "`" + ` to apply. This ` + "`" + `condition` + "`" + ` must be of type ` + "`" + `RESPONSE` + "`" + `. For detailed information about Conditionals, see [Fastly's Documentation on Conditionals][fastly-conditionals].`,
				},
				resource.Attribute{
					Name:        "message_type",
					Description: `(Optional) How the message should be formatted; one of: ` + "`" + `classic` + "`" + `, ` + "`" + `loggly` + "`" + `, ` + "`" + `logplex` + "`" + ` or ` + "`" + `blank` + "`" + `. Default ` + "`" + `classic` + "`" + `. [Fastly Documentation](https://docs.fastly.com/api/logging#logging_gcs)`,
				},
				resource.Attribute{
					Name:        "placement",
					Description: `(Optional) Where in the generated VCL the logging call should be placed; one of: ` + "`" + `none` + "`" + ` or ` + "`" + `waf_debug` + "`" + `. The ` + "`" + `bigquerylogging` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A unique name to identify this BigQuery logging endpoint.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The ID of your GCP project.`,
				},
				resource.Attribute{
					Name:        "dataset",
					Description: `(Required) The ID of your BigQuery dataset.`,
				},
				resource.Attribute{
					Name:        "table",
					Description: `(Required) The ID of your BigQuery table.`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `(Optional) The email for the service account with write access to your BigQuery dataset. If not provided, this will be pulled from a ` + "`" + `FASTLY_BQ_EMAIL` + "`" + ` environment variable.`,
				},
				resource.Attribute{
					Name:        "secret_key",
					Description: `(Optional) The secret key associated with the sservice account that has write access to your BigQuery table. If not provided, this will be pulled from the ` + "`" + `FASTLY_BQ_SECRET_KEY` + "`" + ` environment variable. Typical format for this is a private key in a string with newlines.`,
				},
				resource.Attribute{
					Name:        "format",
					Description: `(Optional) Apache style log formatting. Must produce JSON that matches the schema of your BigQuery table.`,
				},
				resource.Attribute{
					Name:        "response_condition",
					Description: `(Optional) Name of already defined ` + "`" + `condition` + "`" + ` to apply. This ` + "`" + `condition` + "`" + ` must be of type ` + "`" + `RESPONSE` + "`" + `. For detailed information about Conditionals, see [Fastly's Documentation on Conditionals][fastly-conditionals].`,
				},
				resource.Attribute{
					Name:        "template",
					Description: `(Optional) Big query table name suffix template. If set will be interpreted as a strftime compatible string and used as the [Template Suffix for your table](https://cloud.google.com/bigquery/streaming-data-into-bigquery#template-tables).`,
				},
				resource.Attribute{
					Name:        "placement",
					Description: `(Optional) Where in the generated VCL the logging call should be placed; one of: ` + "`" + `none` + "`" + ` or ` + "`" + `waf_debug` + "`" + `. The ` + "`" + `syslog` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A unique name to identify this Syslog endpoint.`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `(Required) A hostname or IPv4 address of the Syslog endpoint.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional) The port associated with the address where the Syslog endpoint can be accessed. Default ` + "`" + `514` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "format",
					Description: `(Optional) Apache-style string or VCL variables to use for log formatting. Defaults to Apache Common Log format (%h %l %u %t %r %>s)`,
				},
				resource.Attribute{
					Name:        "format_version",
					Description: `(Optional) The version of the custom logging format used for the configured endpoint. Can be either 1 (the default, version 1 log format) or 2 (the version 2 log format).`,
				},
				resource.Attribute{
					Name:        "token",
					Description: `(Optional) Whether to prepend each message with a specific token.`,
				},
				resource.Attribute{
					Name:        "use_tls",
					Description: `(Optional) Whether to use TLS for secure logging. Default ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "tls_hostname",
					Description: `(Optional) Used during the TLS handshake to validate the certificate.`,
				},
				resource.Attribute{
					Name:        "tls_ca_cert",
					Description: `(Optional) A secure certificate to authenticate the server with.`,
				},
				resource.Attribute{
					Name:        "response_condition",
					Description: `(Optional) Name of already defined ` + "`" + `condition` + "`" + ` to apply. This ` + "`" + `condition` + "`" + ` must be of type ` + "`" + `RESPONSE` + "`" + `. For detailed information about Conditionals, see [Fastly's Documentation on Conditionals][fastly-conditionals].`,
				},
				resource.Attribute{
					Name:        "message_type",
					Description: `(Optional) How the message should be formatted; one of: ` + "`" + `classic` + "`" + `, ` + "`" + `loggly` + "`" + `, ` + "`" + `logplex` + "`" + ` or ` + "`" + `blank` + "`" + `. Default ` + "`" + `classic` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "placement",
					Description: `(Optional) Where in the generated VCL the logging call should be placed; one of: ` + "`" + `none` + "`" + ` or ` + "`" + `waf_debug` + "`" + `. The ` + "`" + `logentries` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A unique name to identify this GCS endpoint.`,
				},
				resource.Attribute{
					Name:        "token",
					Description: `(Required) Logentries Token to be used for authentication (https://logentries.com/doc/input-token/).`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional) The port number configured in Logentries to send logs to. Defaults to ` + "`" + `20000` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "use_tls",
					Description: `(Optional) Whether to use TLS for secure logging. Defaults to ` + "`" + `true` + "`" + ``,
				},
				resource.Attribute{
					Name:        "format",
					Description: `(Optional) Apache-style string or VCL variables to use for log formatting. Defaults to Apache Common Log format (` + "`" + `%h %l %u %t %r %>s` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "format_version",
					Description: `(Optional) The version of the custom logging format used for the configured endpoint. Can be either 1 (the default, version 1 log format) or 2 (the version 2 log format).`,
				},
				resource.Attribute{
					Name:        "response_condition",
					Description: `(Optional) Name of already defined ` + "`" + `condition` + "`" + ` to apply. This ` + "`" + `condition` + "`" + ` must be of type ` + "`" + `RESPONSE` + "`" + `. For detailed information about Conditionals, see [Fastly's Documentation on Conditionals][fastly-conditionals].`,
				},
				resource.Attribute{
					Name:        "placement",
					Description: `(Optional) Where in the generated VCL the logging call should be placed; one of: ` + "`" + `none` + "`" + ` or ` + "`" + `waf_debug` + "`" + `. The ` + "`" + `splunk` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A unique name to identify the Splunk endpoint.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Required) The Splunk URL to stream logs to.`,
				},
				resource.Attribute{
					Name:        "token",
					Description: `(Required) The Splunk token to be used for authentication.`,
				},
				resource.Attribute{
					Name:        "format",
					Description: `(Optional) Apache-style string or VCL variables to use for log formatting. Default ` + "`" + `%h %l %u %t \"%r\" %>s %b` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "format_version",
					Description: `(Optional) The version of the custom logging format used for the configured endpoint. Can be either ` + "`" + `1` + "`" + ` or ` + "`" + `2` + "`" + `. The logging call gets placed by default in ` + "`" + `vcl_log` + "`" + ` if ` + "`" + `format_version` + "`" + ` is set to ` + "`" + `2` + "`" + ` and in ` + "`" + `vcl_deliver` + "`" + ` if ` + "`" + `format_version` + "`" + ` is set to ` + "`" + `1` + "`" + `. Default ` + "`" + `2` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "placement",
					Description: `(Optional) Where in the generated VCL the logging call should be placed, overriding any ` + "`" + `format_version` + "`" + ` default. Can be either ` + "`" + `none` + "`" + ` or ` + "`" + `waf_debug` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "response_condition",
					Description: `(Optional) The name of the ` + "`" + `condition` + "`" + ` to apply. If empty, always execute. The ` + "`" + `blobstoragelogging` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A unique name to identify the Azure Blob Storage endpoint.`,
				},
				resource.Attribute{
					Name:        "account_name",
					Description: `(Required) The unique Azure Blob Storage namespace in which your data objects are stored.`,
				},
				resource.Attribute{
					Name:        "container",
					Description: `(Required) The name of the Azure Blob Storage container in which to store logs.`,
				},
				resource.Attribute{
					Name:        "sas_token",
					Description: `(Required) The Azure shared access signature providing write access to the blob service objects. Be sure to update your token before it expires or the logging functionality will not work.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Optional) The path to upload logs to. Must end with a trailing slash. If this field is left empty, the files will be saved in the container's root path.`,
				},
				resource.Attribute{
					Name:        "period",
					Description: `(Optional) How frequently the logs should be transferred in seconds. Default ` + "`" + `3600` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "timestamp_format",
					Description: `(Optional) ` + "`" + `strftime` + "`" + ` specified timestamp formatting. Default ` + "`" + `%Y-%m-%dT%H:%M:%S.000` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "gzip_level",
					Description: `(Optional) Level of GZIP compression from ` + "`" + `0` + "`" + `to ` + "`" + `9` + "`" + `. ` + "`" + `0` + "`" + ` means no compression. ` + "`" + `1` + "`" + ` is the fastest and the least compressed version, ` + "`" + `9` + "`" + ` is the slowest and the most compressed version. Default ` + "`" + `0` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "public_key",
					Description: `(Optional) A PGP public key that Fastly will use to encrypt your log files before writing them to disk.`,
				},
				resource.Attribute{
					Name:        "format",
					Description: `(Optional) Apache-style string or VCL variables to use for log formatting. Default ` + "`" + `%h %l %u %t \"%r\" %>s %b` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "format_version",
					Description: `(Optional) The version of the custom logging format used for the configured endpoint. Can be either ` + "`" + `1` + "`" + ` or ` + "`" + `2` + "`" + `. The logging call gets placed by default in ` + "`" + `vcl_log` + "`" + ` if ` + "`" + `format_version` + "`" + ` is set to ` + "`" + `2` + "`" + ` and in ` + "`" + `vcl_deliver` + "`" + ` if ` + "`" + `format_version` + "`" + ` is set to ` + "`" + `1` + "`" + `. Default ` + "`" + `2` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "message_type",
					Description: `(Optional) How the message should be formatted. Can be either ` + "`" + `classic` + "`" + `, ` + "`" + `loggly` + "`" + `, ` + "`" + `logplex` + "`" + ` or ` + "`" + `blank` + "`" + `. Default ` + "`" + `classic` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "placement",
					Description: `(Optional) Where in the generated VCL the logging call should be placed, overriding any ` + "`" + `format_version` + "`" + ` default. Can be either ` + "`" + `none` + "`" + ` or ` + "`" + `waf_debug` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "response_condition",
					Description: `(Optional) The name of the ` + "`" + `condition` + "`" + ` to apply. If empty, always execute. The ` + "`" + `response_object` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A unique name to identify this Response Object.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Optional) The HTTP Status Code. Default ` + "`" + `200` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "response",
					Description: `(Optional) The HTTP Response. Default ` + "`" + `Ok` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "content",
					Description: `(Optional) The content to deliver for the response object.`,
				},
				resource.Attribute{
					Name:        "content_type",
					Description: `(Optional) The MIME type of the content.`,
				},
				resource.Attribute{
					Name:        "request_condition",
					Description: `(Optional) Name of already defined ` + "`" + `condition` + "`" + ` to be checked during the request phase. If the condition passes then this object will be delivered. This ` + "`" + `condition` + "`" + ` must be of type ` + "`" + `REQUEST` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "cache_condition",
					Description: `(Optional) Name of already defined ` + "`" + `condition` + "`" + ` to check after we have retrieved an object. If the condition passes then deliver this Request Object instead. This ` + "`" + `condition` + "`" + ` must be of type ` + "`" + `CACHE` + "`" + `. For detailed information about Conditionals, see [Fastly's Documentation on Conditionals][fastly-conditionals]. The ` + "`" + `snippet` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A unique name for the VCL Snippet configuration block.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) The location in generated VCL where the snippet should be placed (can be one of ` + "`" + `init` + "`" + `, ` + "`" + `recv` + "`" + `, ` + "`" + `hit` + "`" + `, ` + "`" + `miss` + "`" + `, ` + "`" + `pass` + "`" + `, ` + "`" + `fetch` + "`" + `, ` + "`" + `error` + "`" + `, ` + "`" + `deliver` + "`" + `, ` + "`" + `log` + "`" + ` or ` + "`" + `none` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `(Optional) Priority determines the ordering for multiple snippets. Lower numbers execute first. Defaults to ` + "`" + `100` + "`" + `. The ` + "`" + `vcl` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A unique name for this configuration block.`,
				},
				resource.Attribute{
					Name:        "content",
					Description: `(Required) The custom VCL code to upload.`,
				},
				resource.Attribute{
					Name:        "main",
					Description: `(Optional) If ` + "`" + `true` + "`" + `, use this block as the main configuration. If ` + "`" + `false` + "`" + `, use this block as an includable library. Only a single VCL block can be marked as the main block. Default is ` + "`" + `false` + "`" + `. ## Attributes Reference In addition to the arguments listed above, the following attributes are exported:`,
				},
			},
			Attributes: []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"fastly_service_v1": 0,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
