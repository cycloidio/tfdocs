package fastly

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "fastly_service_acl_entries_v1",
			Category:         "Resources",
			ShortDescription: `Defines a set of Fastly ACL entries that can be used to populate a service ACL.`,
			Description:      ``,
			Keywords: []string{
				"service",
				"acl",
				"entries",
				"v1",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "service_id",
					Description: `(Required) The ID of the Service that the ACL belongs to`,
				},
				resource.Attribute{
					Name:        "acl_id",
					Description: `(Required) The ID of the ACL that the items belong to`,
				},
				resource.Attribute{
					Name:        "entry",
					Description: `(Optional) A Set ACL entries that are applied to the service. Defined below The ` + "`" + `entry` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `(Required, string) An IP address that is the focus for the ACL`,
				},
				resource.Attribute{
					Name:        "subnet",
					Description: `(Optional, string) An optional subnet mask applied to the IP address`,
				},
				resource.Attribute{
					Name:        "negated",
					Description: `(Optional, boolean) A boolean that will negate the match if true`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `(Optional, string) A personal freeform descriptive note ## Attributes Reference`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fastly_service_dictionary_items_v1",
			Category:         "Resources",
			ShortDescription: `Provides a grouping of Fastly dictionary items that can be applied to a service.`,
			Description:      ``,
			Keywords: []string{
				"service",
				"dictionary",
				"items",
				"v1",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "service_id",
					Description: `(Required) The ID of the service that the dictionary belongs to`,
				},
				resource.Attribute{
					Name:        "dictionary_id",
					Description: `(Required) The ID of the dictionary that the items belong to`,
				},
				resource.Attribute{
					Name:        "items",
					Description: `(Optional) A map representing an entry in the dictionary, (key/value) ## Attributes Reference`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fastly_service_dynamic_snippet_content_v1",
			Category:         "Resources",
			ShortDescription: `Provides a means to define blocks of VCL logic that is inserted into your service through Fastly dynamic snippets.`,
			Description:      ``,
			Keywords: []string{
				"service",
				"dynamic",
				"snippet",
				"content",
				"v1",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "service_id",
					Description: `(Required) The ID of the service that the dynamic snippet belongs to`,
				},
				resource.Attribute{
					Name:        "snippet_id",
					Description: `(Required) The ID of the dynamic snippet that the content belong to`,
				},
				resource.Attribute{
					Name:        "content",
					Description: `(Required) The VCL code that specifies exactly what the snippet does. ## Attributes Reference`,
				},
			},
			Attributes: []resource.Attribute{},
		},
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
					Description: `(Optional) Automated healthchecks on the cache that can change how Fastly interacts with the cache based on its health.`,
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
					Name:        "httpslogging",
					Description: `(Optional) An HTTPS endpoint to send streaming logs to. Defined below.`,
				},
				resource.Attribute{
					Name:        "logging_elasticsearch",
					Description: `(optional) An Elasticsearch endpoint to send streaming logs to. Defined below.`,
				},
				resource.Attribute{
					Name:        "logging_ftp",
					Description: `(Optional) An FTP endpoint to send streaming logs to. Defined below.`,
				},
				resource.Attribute{
					Name:        "logging_sftp",
					Description: `(Optional) An SFTP endpoint to send streaming logs to. Defined below.`,
				},
				resource.Attribute{
					Name:        "logging_datadog",
					Description: `(Optional) A Datadog endpoint to send streaming logs to. Defined below.`,
				},
				resource.Attribute{
					Name:        "logging_loggly",
					Description: `(Optional) A Loggly endpoint to send streaming logs to. Defined below.`,
				},
				resource.Attribute{
					Name:        "logging_newrelic",
					Description: `(Optional) A New Relic endpoint to send streaming logs to. Defined below.`,
				},
				resource.Attribute{
					Name:        "logging_scalyr",
					Description: `(Optional) A Scalyr endpoint to send streaming logs to. Defined below.`,
				},
				resource.Attribute{
					Name:        "logging_googlepubsub",
					Description: `(Optional) A Google Cloud Pub/Sub endpoint to send streaming logs to. Defined below.`,
				},
				resource.Attribute{
					Name:        "logging_kafka",
					Description: `(Optional) A Kafka endpoint to send streaming logs to. Defined below.`,
				},
				resource.Attribute{
					Name:        "logging_heroku",
					Description: `(Optional) A Heroku endpoint to send streaming logs to. Defined below.`,
				},
				resource.Attribute{
					Name:        "logging_honeycomb",
					Description: `(Optional) A Honeycomb endpoint to send streaming logs to. Defined below.`,
				},
				resource.Attribute{
					Name:        "logging_logshuttle",
					Description: `(Optional) A Log Shuttle endpoint to send streaming logs to. Defined below.`,
				},
				resource.Attribute{
					Name:        "logging_openstack",
					Description: `(Optional) An OpenStack endpoint to send streaming logs to. Defined below.`,
				},
				resource.Attribute{
					Name:        "logging_digitalocean",
					Description: `(Optional) A DigitalOcean Spaces endpoint to send streaming logs to. Defined below.`,
				},
				resource.Attribute{
					Name:        "logging_cloudfiles",
					Description: `(Optional) A Rackspace Cloud Files endpoint to send streaming logs to. Defined below.`,
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
					Name:        "dynamicsnippet",
					Description: `(Optional) A set of custom, "dynamic" VCL Snippet configuration blocks. Defined below.`,
				},
				resource.Attribute{
					Name:        "vcl",
					Description: `(Optional) A set of custom VCL configuration blocks. See the [Fastly documentation](https://docs.fastly.com/vcl/custom-vcl/uploading-custom-vcl/) for more information on using custom VCL.`,
				},
				resource.Attribute{
					Name:        "acl",
					Description: `(Optional) A set of ACL configuration blocks. Defined below.`,
				},
				resource.Attribute{
					Name:        "dictionary",
					Description: `(Optional) A set of dictionaries that allow the storing of key values pair for use within VCL functions. Defined below. The ` + "`" + `domain` + "`" + ` block supports:`,
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
					Description: `(Optional) The POP of the shield designated to reduce inbound load. Valid values for ` + "`" + `shield` + "`" + ` are included in the [` + "`" + `GET /datacenters` + "`" + `](https://developer.fastly.com/reference/api/utils/datacenter/) API response.`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `(Optional) The [portion of traffic](https://docs.fastly.com/en/guides/load-balancing-configuration#how-weight-affects-load-balancing) to send to this Backend. Each Backend receives ` + "`" + `weight / total` + "`" + ` of the traffic. Default ` + "`" + `100` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "healthcheck",
					Description: `(Optional) Name of a defined ` + "`" + `healthcheck` + "`" + ` to assign to this backend. The ` + "`" + `condition` + "`" + ` block supports allows you to add logic to any basic configuration object in a service. See Fastly's documentation ["About Conditions"](https://docs.fastly.com/en/guides/about-conditions) for more detailed information on using Conditions. The Condition ` + "`" + `name` + "`" + ` can be used in the ` + "`" + `request_condition` + "`" + `, ` + "`" + `response_condition` + "`" + `, or ` + "`" + `cache_condition` + "`" + ` attributes of other block settings.`,
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
					Description: `(Optional) Selected POP to serve as a "shield" for backends. Valid values for ` + "`" + `shield` + "`" + ` are included in the [` + "`" + `GET /datacenters` + "`" + `](https://developer.fastly.com/reference/api/utils/datacenter/) API response.`,
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
					Description: `(Optional) One of ` + "`" + `cache` + "`" + `, ` + "`" + `pass` + "`" + `, or ` + "`" + `restart` + "`" + `, as defined on Fastly's documentation under ["Caching action descriptions"](https://docs.fastly.com/en/guides/controlling-caching#caching-action-descriptions).`,
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
					Description: `(Optional) Name of already defined ` + "`" + `condition` + "`" + ` controlling when this gzip configuration applies. This ` + "`" + `condition` + "`" + ` must be of type ` + "`" + `CACHE` + "`" + `. For detailed information about Conditionals, see [Fastly's Documentation on Conditionals][fastly-conditionals]. The ` + "`" + `header` + "`" + ` block supports adding, removing, or modifying Request and Response headers. See Fastly's documentation on [Adding or modifying headers on HTTP requests and responses](https://docs.fastly.com/en/guides/adding-or-modifying-headers-on-http-requests-and-responses#field-description-table) for more detailed information on any of the properties below.`,
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
					Description: `(Required) Unique name to refer to this Request Setting.`,
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
					Description: `(Required) The unique name of the S3 logging endpoint.`,
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
					Name:        "public_key",
					Description: `(Optional) A PGP public key that Fastly will use to encrypt your log files before writing them to disk.`,
				},
				resource.Attribute{
					Name:        "period",
					Description: `(Optional) How frequently the logs should be transferred, in seconds. Default ` + "`" + `3600` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "gzip_level",
					Description: `(Optional) Level of Gzip compression, from ` + "`" + `0-9` + "`" + `. ` + "`" + `0` + "`" + ` is no compression. ` + "`" + `1` + "`" + ` is fastest and least compressed, ` + "`" + `9` + "`" + ` is slowest and most compressed. Default ` + "`" + `0` + "`" + `.`,
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
					Description: `(Optional) Level of Gzip compression, from ` + "`" + `0-9` + "`" + `. ` + "`" + `0` + "`" + ` is no compression. ` + "`" + `1` + "`" + ` is fastest and least compressed, ` + "`" + `9` + "`" + ` is slowest and most compressed. Default ` + "`" + `0` + "`" + `.`,
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
					Description: `(Optional) How the message should be formatted; one of: ` + "`" + `classic` + "`" + `, ` + "`" + `loggly` + "`" + `, ` + "`" + `logplex` + "`" + ` or ` + "`" + `blank` + "`" + `. Default ` + "`" + `classic` + "`" + `. [Fastly Documentation](https://developer.fastly.com/reference/api/logging/gcs/)`,
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
					Description: `(Optional) A secure certificate to authenticate the server with. Must be in PEM format. You can provide this certificate via an environment variable, ` + "`" + `FASTLY_SYSLOG_CA_CERT` + "`" + ``,
				},
				resource.Attribute{
					Name:        "tls_client_cert",
					Description: `(Optional) The client certificate used to make authenticated requests. Must be in PEM format. You can provide this certificate via an environment variable, ` + "`" + `FASTLY_SYSLOG_CLIENT_CERT` + "`" + ``,
				},
				resource.Attribute{
					Name:        "tls_client_key",
					Description: `(Optional) The client private key used to make authenticated requests. Must be in PEM format. You can provide this key via an environment variable, ` + "`" + `FASTLY_SYSLOG_CLIENT_KEY` + "`" + ``,
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
					Description: `(Optional) The name of the ` + "`" + `condition` + "`" + ` to apply. If empty, always execute.`,
				},
				resource.Attribute{
					Name:        "tls_hostname",
					Description: `(Optional) The hostname used to verify the server's certificate. It can either be the Common Name or a Subject Alternative Name (SAN).`,
				},
				resource.Attribute{
					Name:        "tls_ca_cert",
					Description: `(Optional) A secure certificate to authenticate the server with. Must be in PEM format. You can provide this certificate via an environment variable, ` + "`" + `FASTLY_SPLUNK_CA_CERT` + "`" + `. The ` + "`" + `blobstoragelogging` + "`" + ` block supports:`,
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
					Description: `(Optional) Level of Gzip compression from ` + "`" + `0` + "`" + `to ` + "`" + `9` + "`" + `. ` + "`" + `0` + "`" + ` means no compression. ` + "`" + `1` + "`" + ` is the fastest and the least compressed version, ` + "`" + `9` + "`" + ` is the slowest and the most compressed version. Default ` + "`" + `0` + "`" + `.`,
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
					Description: `(Optional) The name of the ` + "`" + `condition` + "`" + ` to apply. If empty, always execute. The ` + "`" + `httpslogging` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The unique name of the HTTPS logging endpoint.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Required) URL that log data will be sent to. Must use the https protocol.`,
				},
				resource.Attribute{
					Name:        "request_max_entries",
					Description: `(Optional) The maximum number of logs sent in one request.`,
				},
				resource.Attribute{
					Name:        "request_max_bytes",
					Description: `(Optional) The maximum number of bytes sent in one request.`,
				},
				resource.Attribute{
					Name:        "content_type",
					Description: `(Optional) Value of the ` + "`" + `Content-Type` + "`" + ` header sent with the request.`,
				},
				resource.Attribute{
					Name:        "header_name",
					Description: `(Optional) Custom header sent with the request.`,
				},
				resource.Attribute{
					Name:        "header_value",
					Description: `(Optional) Value of the custom header sent with the request.`,
				},
				resource.Attribute{
					Name:        "method",
					Description: `(Optional) HTTP method used for request. Can be either ` + "`" + `POST` + "`" + ` or ` + "`" + `PUT` + "`" + `. Default ` + "`" + `POST` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "json_format",
					Description: `Formats log entries as JSON. Can be either disabled (` + "`" + `0` + "`" + `), array of json (` + "`" + `1` + "`" + `), or newline delimited json (` + "`" + `2` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "tls_hostname",
					Description: `(Optional) Used during the TLS handshake to validate the certificate.`,
				},
				resource.Attribute{
					Name:        "tls_ca_cert",
					Description: `(Optional) A secure certificate to authenticate the server with. Must be in PEM format.`,
				},
				resource.Attribute{
					Name:        "tls_client_cert",
					Description: `(Optional) The client certificate used to make authenticated requests. Must be in PEM format.`,
				},
				resource.Attribute{
					Name:        "tls_client_key",
					Description: `(Optional) The client private key used to make authenticated requests. Must be in PEM format.`,
				},
				resource.Attribute{
					Name:        "format",
					Description: `(Optional) Apache-style string or VCL variables to use for log formatting.`,
				},
				resource.Attribute{
					Name:        "format_version",
					Description: `(Optional) The version of the custom logging format used for the configured endpoint. Can be either ` + "`" + `1` + "`" + ` or ` + "`" + `2` + "`" + `. The logging call gets placed by default in ` + "`" + `vcl_log` + "`" + ` if ` + "`" + `format_version` + "`" + ` is set to ` + "`" + `2` + "`" + ` and in ` + "`" + `vcl_deliver` + "`" + ` if ` + "`" + `format_version` + "`" + ` is set to ` + "`" + `1` + "`" + `. Default ` + "`" + `2` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "message_type",
					Description: `How the message should be formatted; one of: ` + "`" + `classic` + "`" + `, ` + "`" + `loggly` + "`" + `, ` + "`" + `logplex` + "`" + ` or ` + "`" + `blank` + "`" + `. Default ` + "`" + `blank` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "placement",
					Description: `(Optional) Where in the generated VCL the logging call should be placed.`,
				},
				resource.Attribute{
					Name:        "response_condition",
					Description: `(Optional) The name of the ` + "`" + `condition` + "`" + ` to apply. If empty, always execute. The ` + "`" + `logging_elasticsearch` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The unique name of the Elasticsearch logging endpoint.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Required) The Elasticsearch URL to stream logs to.`,
				},
				resource.Attribute{
					Name:        "index",
					Description: `(Required) The name of the Elasticsearch index to send documents (logs) to.`,
				},
				resource.Attribute{
					Name:        "user",
					Description: `(Optional) BasicAuth username for Elasticsearch.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional) BasicAuth password for Elasticsearch.`,
				},
				resource.Attribute{
					Name:        "pipeline",
					Description: `(Optional) The ID of the Elasticsearch ingest pipeline to apply pre-process transformations to before indexing.`,
				},
				resource.Attribute{
					Name:        "request_max_bytes",
					Description: `(Optional) The maximum number of bytes sent in one request. Defaults to ` + "`" + `0` + "`" + ` for unbounded.`,
				},
				resource.Attribute{
					Name:        "request_max_entries",
					Description: `(Optional) The maximum number of logs sent in one request. Defaults to ` + "`" + `0` + "`" + ` for unbounded.`,
				},
				resource.Attribute{
					Name:        "tls_ca_cert",
					Description: `(Optional) A secure certificate to authenticate the server with. Must be in PEM format.`,
				},
				resource.Attribute{
					Name:        "tls_client_cert",
					Description: `(Optional) The client certificate used to make authenticated requests. Must be in PEM format.`,
				},
				resource.Attribute{
					Name:        "tls_client_key",
					Description: `(Optional) The client private key used to make authenticated requests. Must be in PEM format.`,
				},
				resource.Attribute{
					Name:        "tls_hostname",
					Description: `(Optional) The hostname used to verify the server's certificate. It can either be the Common Name (CN) or a Subject Alternative Name (SAN).`,
				},
				resource.Attribute{
					Name:        "format",
					Description: `(Optional) Apache-style string or VCL variables to use for log formatting.`,
				},
				resource.Attribute{
					Name:        "format_version",
					Description: `(Optional) The version of the custom logging format used for the configured endpoint. Can be either ` + "`" + `1` + "`" + ` or ` + "`" + `2` + "`" + `. (default: ` + "`" + `2` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "placement",
					Description: `(Optional) Where in the generated VCL the logging call should be placed.`,
				},
				resource.Attribute{
					Name:        "response_condition",
					Description: `(Optional) The name of the ` + "`" + `condition` + "`" + ` to apply. If empty, always execute. The ` + "`" + `logging_ftp` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The unique name of the FTP logging endpoint.`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `(Required) The FTP address to stream logs to.`,
				},
				resource.Attribute{
					Name:        "user",
					Description: `(Required) The username for the server (can be ` + "`" + `anonymous` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Required) The password for the server (for anonymous use an email address).`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Required) The path to upload log files to. If the path ends in ` + "`" + `/` + "`" + ` then it is treated as a directory.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional) The port number. Default: ` + "`" + `21` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "gzip_level",
					Description: `(Optional) Gzip Compression level. Default ` + "`" + `0` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "period",
					Description: `(Optional) How frequently the logs should be transferred, in seconds (Default 3600).`,
				},
				resource.Attribute{
					Name:        "public_key",
					Description: `(Optional) The PGP public key that Fastly will use to encrypt your log files before writing them to disk.`,
				},
				resource.Attribute{
					Name:        "timestamp_format",
					Description: `(Optional) specified timestamp formatting (default ` + "`" + `%Y-%m-%dT%H:%M:%S.000` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "format",
					Description: `(Optional) Apache-style string or VCL variables to use for log formatting.`,
				},
				resource.Attribute{
					Name:        "format_version",
					Description: `(Optional) The version of the custom logging format used for the configured endpoint. Can be either ` + "`" + `1` + "`" + ` or ` + "`" + `2` + "`" + `. (default: ` + "`" + `2` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "placement",
					Description: `(Optional) Where in the generated VCL the logging call should be placed.`,
				},
				resource.Attribute{
					Name:        "response_condition",
					Description: `(Optional) The name of the condition to apply. The ` + "`" + `logging_sftp` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The unique name of the SFTP logging endpoint.`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `(Required) The SFTP address to stream logs to.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Required) The path to upload log files to. If the path ends in / then it is treated as a directory.`,
				},
				resource.Attribute{
					Name:        "ssh_known_hosts",
					Description: `(Required) A list of host keys for all hosts we can connect to over SFTP.`,
				},
				resource.Attribute{
					Name:        "user",
					Description: `(Required) The username for the server.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional) The port the SFTP service listens on. (Default: ` + "`" + `22` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional) The password for the server. If both ` + "`" + `password` + "`" + ` and ` + "`" + `secret_key` + "`" + ` are passed, ` + "`" + `secret_key` + "`" + ` will be preferred.`,
				},
				resource.Attribute{
					Name:        "secret_key",
					Description: `(Optional) The SSH private key for the server. If both ` + "`" + `password` + "`" + ` and ` + "`" + `secret_key` + "`" + ` are passed, ` + "`" + `secret_key` + "`" + ` will be preferred.`,
				},
				resource.Attribute{
					Name:        "gzip_level",
					Description: `(Optional) What level of Gzip encoding to have when dumping logs (default 0, no compression).`,
				},
				resource.Attribute{
					Name:        "period",
					Description: `(Optional) How frequently log files are finalized so they can be available for reading (in seconds, default ` + "`" + `3600` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "placement",
					Description: `(Optional) Where in the generated VCL the logging call should be placed.`,
				},
				resource.Attribute{
					Name:        "public_key",
					Description: `(Optional) A PGP public key that Fastly will use to encrypt your log files before writing them to disk.`,
				},
				resource.Attribute{
					Name:        "format",
					Description: `(Optional) Apache-style string or VCL variables to use for log formatting.`,
				},
				resource.Attribute{
					Name:        "format_version",
					Description: `(Optional) The version of the custom logging format used for the configured endpoint. Can be either ` + "`" + `1` + "`" + ` or ` + "`" + `2` + "`" + `. (default: ` + "`" + `2` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "response_condition",
					Description: `(Optional) The name of the condition to apply.`,
				},
				resource.Attribute{
					Name:        "timestamp_format",
					Description: `(Optional) The strftime specified timestamp formatting (default ` + "`" + `%Y-%m-%dT%H:%M:%S.000` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "message_type",
					Description: `(Optional) How the message should be formatted. One of: classic (default), loggly, logplex or blank. The ` + "`" + `logging_datadog` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The unique name of the Datadog logging endpoint.`,
				},
				resource.Attribute{
					Name:        "token",
					Description: `(Required) The API key from your Datadog account.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region that log data will be sent to. One of ` + "`" + `US` + "`" + ` or ` + "`" + `EU` + "`" + `. Defaults to ` + "`" + `US` + "`" + ` if undefined.`,
				},
				resource.Attribute{
					Name:        "format",
					Description: `(Optional) Apache-style string or VCL variables to use for log formatting.`,
				},
				resource.Attribute{
					Name:        "format_version",
					Description: `(Optional) The version of the custom logging format used for the configured endpoint. Can be either ` + "`" + `1` + "`" + ` or ` + "`" + `2` + "`" + `. (default: ` + "`" + `2` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "placement",
					Description: `(Optional) Where in the generated VCL the logging call should be placed.`,
				},
				resource.Attribute{
					Name:        "response_condition",
					Description: `(Optional) The name of the condition to apply. The ` + "`" + `logging_loggly` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The unique name of the Loggly logging endpoint.`,
				},
				resource.Attribute{
					Name:        "token",
					Description: `(Required) The token to use for authentication (https://www.loggly.com/docs/customer-token-authentication-token/).`,
				},
				resource.Attribute{
					Name:        "format",
					Description: `(Optional) Apache-style string or VCL variables to use for log formatting.`,
				},
				resource.Attribute{
					Name:        "format_version",
					Description: `(Optional) The version of the custom logging format used for the configured endpoint. Can be either ` + "`" + `1` + "`" + ` or ` + "`" + `2` + "`" + `. (default: ` + "`" + `2` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "placement",
					Description: `(Optional) Where in the generated VCL the logging call should be placed. Can be ` + "`" + `none` + "`" + ` or ` + "`" + `waf_debug` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "response_condition",
					Description: `(Optional) The name of an existing condition in the configured endpoint, or leave blank to always execute. The ` + "`" + `logging_newrelic` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The unique name of the New Relic logging endpoint.`,
				},
				resource.Attribute{
					Name:        "token",
					Description: `(Required) The Insert API key from the Account page of your New Relic account.`,
				},
				resource.Attribute{
					Name:        "format",
					Description: `(Optional) Apache style log formatting. Your log must produce valid JSON that New Relic Logs can ingest.`,
				},
				resource.Attribute{
					Name:        "format_version",
					Description: `(Optional) The version of the custom logging format used for the configured endpoint. Can be either ` + "`" + `1` + "`" + ` or ` + "`" + `2` + "`" + `. (default: ` + "`" + `2` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "placement",
					Description: `(Optional) Where in the generated VCL the logging call should be placed.`,
				},
				resource.Attribute{
					Name:        "response_condition",
					Description: `(Optional) The name of the condition to apply. The ` + "`" + `logging_scalyr` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The unique name of the Scalyr logging endpoint.`,
				},
				resource.Attribute{
					Name:        "token",
					Description: `(Required) The token to use for authentication (https://www.scalyr.com/keys).`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region that log data will be sent to. One of US or EU. Defaults to US if undefined.`,
				},
				resource.Attribute{
					Name:        "format",
					Description: `(Optional) Apache-style string or VCL variables to use for log formatting.`,
				},
				resource.Attribute{
					Name:        "format_version",
					Description: `(Optional) The version of the custom logging format used for the configured endpoint. Can be either ` + "`" + `1` + "`" + ` or ` + "`" + `2` + "`" + `. The logging call gets placed by default in ` + "`" + `vcl_log` + "`" + ` if ` + "`" + `format_version` + "`" + ` is set to ` + "`" + `2` + "`" + ` and in ` + "`" + `vcl_deliver` + "`" + ` if ` + "`" + `format_version` + "`" + ` is set to ` + "`" + `1` + "`" + `. Default ` + "`" + `2` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "placement",
					Description: `(Optional) The name of an existing condition in the configured endpoint, or leave blank to always execute.`,
				},
				resource.Attribute{
					Name:        "response_condition",
					Description: `(Optional) The name of the ` + "`" + `condition` + "`" + ` to apply. If empty, always execute. The ` + "`" + `logging_googlepubsub` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The unique name of the Google Cloud Pub/Sub logging endpoint.`,
				},
				resource.Attribute{
					Name:        "user",
					Description: `(Required) Your Google Cloud Platform service account email address. The client_email field in your service account authentication JSON.`,
				},
				resource.Attribute{
					Name:        "secret_key",
					Description: `(Required) Your Google Cloud Platform account secret key. The private_key field in your service account authentication JSON.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The ID of your Google Cloud Platform project.`,
				},
				resource.Attribute{
					Name:        "topic",
					Description: `(Required) The Google Cloud Pub/Sub topic to which logs will be published.`,
				},
				resource.Attribute{
					Name:        "format",
					Description: `(Optional) Apache-style string or VCL variables to use for log formatting.`,
				},
				resource.Attribute{
					Name:        "format_version",
					Description: `(Optional) The version of the custom logging format used for the configured endpoint. Can be either ` + "`" + `1` + "`" + ` or ` + "`" + `2` + "`" + `. The logging call gets placed by default in ` + "`" + `vcl_log` + "`" + ` if ` + "`" + `format_version` + "`" + ` is set to ` + "`" + `2` + "`" + ` and in ` + "`" + `vcl_deliver` + "`" + ` if ` + "`" + `format_version` + "`" + ` is set to ` + "`" + `1` + "`" + `. Default ` + "`" + `2` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "placement",
					Description: `(Optional) The name of an existing condition in the configured endpoint, or leave blank to always execute.`,
				},
				resource.Attribute{
					Name:        "response_condition",
					Description: `(Optional) The name of the ` + "`" + `condition` + "`" + ` to apply. If empty, always execute. The ` + "`" + `logging_kafka` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The unique name of the Kafka logging endpoint.`,
				},
				resource.Attribute{
					Name:        "topic",
					Description: `(Required) The Kafka topic to send logs to.`,
				},
				resource.Attribute{
					Name:        "brokers",
					Description: `(Required) A comma-separated list of IP addresses or hostnames of Kafka brokers.`,
				},
				resource.Attribute{
					Name:        "compression_codec",
					Description: `(Optional) The codec used for compression of your logs. One of: gzip, snappy, lz4.`,
				},
				resource.Attribute{
					Name:        "required_acks",
					Description: `(Optional) The Number of acknowledgements a leader must receive before a write is considered successful. One of: 1 (default) One server needs to respond. 0 No servers need to respond. -1 Wait for all in-sync replicas to respond.`,
				},
				resource.Attribute{
					Name:        "use_tls",
					Description: `(Optional) Whether to use TLS for secure logging. Can be either true or false.`,
				},
				resource.Attribute{
					Name:        "tls_ca_cert",
					Description: `(Optional) A secure certificate to authenticate the server with. Must be in PEM format.`,
				},
				resource.Attribute{
					Name:        "tls_client_cert",
					Description: `(Optional) The client certificate used to make authenticated requests. Must be in PEM format.`,
				},
				resource.Attribute{
					Name:        "tls_client_key",
					Description: `(Optional) The client private key used to make authenticated requests. Must be in PEM format.`,
				},
				resource.Attribute{
					Name:        "tls_hostname",
					Description: `(Optional) The hostname used to verify the server's certificate. It can either be the Common Name or a Subject Alternative Name (SAN).`,
				},
				resource.Attribute{
					Name:        "format",
					Description: `(Optional) Apache-style string or VCL variables to use for log formatting.`,
				},
				resource.Attribute{
					Name:        "format_version",
					Description: `(Optional) The version of the custom logging format used for the configured endpoint. Can be either ` + "`" + `1` + "`" + ` or ` + "`" + `2` + "`" + `. The logging call gets placed by default in ` + "`" + `vcl_log` + "`" + ` if ` + "`" + `format_version` + "`" + ` is set to ` + "`" + `2` + "`" + ` and in ` + "`" + `vcl_deliver` + "`" + ` if ` + "`" + `format_version` + "`" + ` is set to ` + "`" + `1` + "`" + `. Default ` + "`" + `2` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "placement",
					Description: `(Optional) The name of an existing condition in the configured endpoint, or leave blank to always execute.`,
				},
				resource.Attribute{
					Name:        "response_condition",
					Description: `(Optional) The name of the ` + "`" + `condition` + "`" + ` to apply. If empty, always execute. The ` + "`" + `logging_heroku` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The unique name of the Heroku logging endpoint.`,
				},
				resource.Attribute{
					Name:        "token",
					Description: `(Required) The token to use for authentication (https://devcenter.heroku.com/articles/add-on-partner-log-integration).`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Required) The url to stream logs to.`,
				},
				resource.Attribute{
					Name:        "format",
					Description: `(Optional) Apache-style string or VCL variables to use for log formatting.`,
				},
				resource.Attribute{
					Name:        "format_version",
					Description: `(Optional) The version of the custom logging format used for the configured endpoint. Can be either ` + "`" + `1` + "`" + ` or ` + "`" + `2` + "`" + `. (default: ` + "`" + `2` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "placement",
					Description: `(Optional) Where in the generated VCL the logging call should be placed. Can be ` + "`" + `none` + "`" + ` or ` + "`" + `waf_debug` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "response_condition",
					Description: `(Optional) The name of an existing condition in the configured endpoint, or leave blank to always execute. The ` + "`" + `logging_honeycomb` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The unique name of the Honeycomb logging endpoint.`,
				},
				resource.Attribute{
					Name:        "dataset",
					Description: `(Required) The Honeycomb Dataset you want to log to.`,
				},
				resource.Attribute{
					Name:        "token",
					Description: `(Required) The Write Key from the Account page of your Honeycomb account.`,
				},
				resource.Attribute{
					Name:        "format",
					Description: `(Optional) Apache style log formatting. Your log must produce valid JSON that Honeycomb can ingest.`,
				},
				resource.Attribute{
					Name:        "format_version",
					Description: `(Optional) The version of the custom logging format used for the configured endpoint. Can be either ` + "`" + `1` + "`" + ` or ` + "`" + `2` + "`" + `. (default: ` + "`" + `2` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "placement",
					Description: `(Optional) Where in the generated VCL the logging call should be placed. Can be ` + "`" + `none` + "`" + ` or ` + "`" + `waf_debug` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "response_condition",
					Description: `(Optional) The name of an existing condition in the configured endpoint, or leave blank to always execute. The ` + "`" + `logging_logshuttle` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The unique name of the Logshuttle logging endpoint.`,
				},
				resource.Attribute{
					Name:        "token",
					Description: `(Required) The data authentication token associated with this endpoint.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Required) Your Log Shuttle endpoint url.`,
				},
				resource.Attribute{
					Name:        "format",
					Description: `(Optional) Apache style log formatting.`,
				},
				resource.Attribute{
					Name:        "format_version",
					Description: `(Optional) The version of the custom logging format used for the configured endpoint. Can be either ` + "`" + `1` + "`" + ` or ` + "`" + `2` + "`" + `. (default: ` + "`" + `2` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "placement",
					Description: `(Optional) Where in the generated VCL the logging call should be placed. Can be ` + "`" + `none` + "`" + ` or ` + "`" + `waf_debug` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "response_condition",
					Description: `(Optional) The name of an existing condition in the configured endpoint, or leave blank to always execute. The ` + "`" + `logging_openstack` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The unique name of the OpenStack logging endpoint.`,
				},
				resource.Attribute{
					Name:        "bucket_name",
					Description: `(Required) The name of your OpenStack container.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Required) Your OpenStack auth url.`,
				},
				resource.Attribute{
					Name:        "user",
					Description: `(Required) The username for your OpenStack account.`,
				},
				resource.Attribute{
					Name:        "access_key",
					Description: `(Required) Your OpenStack account access key.`,
				},
				resource.Attribute{
					Name:        "public_key",
					Description: `(Optional) A PGP public key that Fastly will use to encrypt your log files before writing them to disk.`,
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
					Description: `(Optional) What level of Gzip encoding to have when dumping logs (default 0, no compression).`,
				},
				resource.Attribute{
					Name:        "message_type",
					Description: `(Optional) How the message should be formatted; one of: ` + "`" + `classic` + "`" + `, ` + "`" + `loggly` + "`" + `, ` + "`" + `logplex` + "`" + ` or ` + "`" + `blank` + "`" + `. Default ` + "`" + `classic` + "`" + `. [Fastly Documentation](https://developer.fastly.com/reference/api/logging/gcs/)`,
				},
				resource.Attribute{
					Name:        "format",
					Description: `(Optional) Apache style log formatting.`,
				},
				resource.Attribute{
					Name:        "format_version",
					Description: `(Optional) The version of the custom logging format used for the configured endpoint. Can be either ` + "`" + `1` + "`" + ` or ` + "`" + `2` + "`" + `. The logging call gets placed by default in ` + "`" + `vcl_log` + "`" + ` if ` + "`" + `format_version` + "`" + ` is set to ` + "`" + `2` + "`" + ` and in ` + "`" + `vcl_deliver` + "`" + ` if ` + "`" + `format_version` + "`" + ` is set to ` + "`" + `1` + "`" + `. Default ` + "`" + `2` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "timestamp_format",
					Description: `(Optional) The strftime specified timestamp formatting (default ` + "`" + `%Y-%m-%dT%H:%M:%S.000` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "response_condition",
					Description: `(Optional) The name of an existing condition in the configured endpoint, or leave blank to always execute.`,
				},
				resource.Attribute{
					Name:        "placement",
					Description: `(Optional) Where in the generated VCL the logging call should be placed; one of: ` + "`" + `none` + "`" + ` or ` + "`" + `waf_debug` + "`" + `. The ` + "`" + `logging_digitalocean` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The unique name of the DigitalOcean Spaces logging endpoint.`,
				},
				resource.Attribute{
					Name:        "bucket_name",
					Description: `(Required) The name of the DigitalOcean Space.`,
				},
				resource.Attribute{
					Name:        "access_key",
					Description: `(Required) Your DigitalOcean Spaces account access key.`,
				},
				resource.Attribute{
					Name:        "secret_key",
					Description: `(Required) Your DigitalOcean Spaces account secret key.`,
				},
				resource.Attribute{
					Name:        "public_key",
					Description: `(Optional) A PGP public key that Fastly will use to encrypt your log files before writing them to disk.`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `(Optional) The domain of the DigitalOcean Spaces endpoint (default "nyc3.digitaloceanspaces.com").`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Optional) The path to upload logs to.`,
				},
				resource.Attribute{
					Name:        "period",
					Description: `(Optional) How frequently log files are finalized so they can be available for reading (in seconds, default 3600).`,
				},
				resource.Attribute{
					Name:        "timestamp_format",
					Description: `(Optional) The strftime specified timestamp formatting (default ` + "`" + `%Y-%m-%dT%H:%M:%S.000` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "gzip_level",
					Description: `(Optional) What level of Gzip encoding to have when dumping logs (default 0, no compression).`,
				},
				resource.Attribute{
					Name:        "format",
					Description: `(Optional) Apache-style string or VCL variables to use for log formatting.`,
				},
				resource.Attribute{
					Name:        "format_version",
					Description: `(Optional) The version of the custom logging format used for the configured endpoint. Can be either ` + "`" + `1` + "`" + ` or ` + "`" + `2` + "`" + `. (default: ` + "`" + `2` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "message_type",
					Description: `(Optional) How the message should be formatted. One of: classic (default), loggly, logplex or blank.`,
				},
				resource.Attribute{
					Name:        "placement",
					Description: `(Optional) Where in the generated VCL the logging call should be placed. Can be ` + "`" + `none` + "`" + ` or ` + "`" + `waf_debug` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "response_condition",
					Description: `(Optional) The name of an existing condition in the configured endpoint, or leave blank to always execute. The ` + "`" + `logging_cloudfiles` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The unique name of the Rackspace Cloud Files logging endpoint.`,
				},
				resource.Attribute{
					Name:        "user",
					Description: `(Required) The username for your Cloud Files account.`,
				},
				resource.Attribute{
					Name:        "bucket_name",
					Description: `(Required) The name of your Cloud Files container.`,
				},
				resource.Attribute{
					Name:        "access_key",
					Description: `(Required) Your Cloud File account access key.`,
				},
				resource.Attribute{
					Name:        "public_key",
					Description: `(Optional) The PGP public key that Fastly will use to encrypt your log files before writing them to disk.`,
				},
				resource.Attribute{
					Name:        "format",
					Description: `(Optional) Apache style log formatting.`,
				},
				resource.Attribute{
					Name:        "format_version",
					Description: `(Optional) The version of the custom logging format used for the configured endpoint. Can be either ` + "`" + `1` + "`" + ` or ` + "`" + `2` + "`" + `. (default: ` + "`" + `2` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "gzip_level",
					Description: `(Optional) What level of GZIP encoding to have when dumping logs (default 0, no compression).`,
				},
				resource.Attribute{
					Name:        "message_type",
					Description: `(Optional) How the message should be formatted. One of: classic (default), loggly, logplex or blank.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Optional) The path to upload logs to.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region to stream logs to. One of: DFW (Dallas), ORD (Chicago), IAD (Northern Virginia), LON (London), SYD (Sydney), HKG (Hong Kong).`,
				},
				resource.Attribute{
					Name:        "period",
					Description: `(Optional) How frequently log files are finalized so they can be available for reading (in seconds, default 3600).`,
				},
				resource.Attribute{
					Name:        "placement",
					Description: `(Optional) Where in the generated VCL the logging call should be placed. Can be ` + "`" + `none` + "`" + ` or ` + "`" + `waf_debug` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "response_condition",
					Description: `(Optional) The name of an existing condition in the configured endpoint, or leave blank to always execute.`,
				},
				resource.Attribute{
					Name:        "timestamp_format",
					Description: `(Optional) The strftime specified timestamp formatting (default ` + "`" + `%Y-%m-%dT%H:%M:%S.000` + "`" + `). The ` + "`" + `response_object` + "`" + ` block supports:`,
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
					Description: `(Required) A name that is unique across "regular" and "dynamic" VCL Snippet configuration blocks.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) The location in generated VCL where the snippet should be placed (can be one of ` + "`" + `init` + "`" + `, ` + "`" + `recv` + "`" + `, ` + "`" + `hit` + "`" + `, ` + "`" + `miss` + "`" + `, ` + "`" + `pass` + "`" + `, ` + "`" + `fetch` + "`" + `, ` + "`" + `error` + "`" + `, ` + "`" + `deliver` + "`" + `, ` + "`" + `log` + "`" + ` or ` + "`" + `none` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `(Optional) Priority determines the ordering for multiple snippets. Lower numbers execute first. Defaults to ` + "`" + `100` + "`" + `. The ` + "`" + `dynamicsnippet` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A name that is unique across "regular" and "dynamic" VCL Snippet configuration blocks.`,
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
					Description: `(Optional) If ` + "`" + `true` + "`" + `, use this block as the main configuration. If ` + "`" + `false` + "`" + `, use this block as an includable library. Only a single VCL block can be marked as the main block. Default is ` + "`" + `false` + "`" + `. The ` + "`" + `acl` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A unique name to identify this ACL. The ` + "`" + `dictionary` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A unique name to identify this dictionary.`,
				},
				resource.Attribute{
					Name:        "write_only",
					Description: `(Optional) If ` + "`" + `true` + "`" + `, the dictionary is a private dictionary, and items are not readable in the UI or via API. Default is ` + "`" + `false` + "`" + `. It is important to note that changing this attribute will delete and recreate the dictionary, discard the current items in the dictionary. Using a write-only/private dictionary should only be done if the items are managed outside of Terraform. ## Attributes Reference In addition to the arguments listed above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "cloned_version",
					Description: `The latest cloned version by the provider. The value gets only set after running ` + "`" + `terraform apply` + "`" + `. The ` + "`" + `dynamicsnippet` + "`" + ` block exports:`,
				},
				resource.Attribute{
					Name:        "snippet_id",
					Description: `The ID of the dynamic snippet. The ` + "`" + `acl` + "`" + ` block exports:`,
				},
				resource.Attribute{
					Name:        "acl_id",
					Description: `The ID of the ACL. The ` + "`" + `dictionary` + "`" + ` block exports:`,
				},
				resource.Attribute{
					Name:        "dictionary_id",
					Description: `The ID of the dictionary. [fastly-s3]: https://docs.fastly.com/en/guides/amazon-s3 [fastly-cname]: https://docs.fastly.com/en/guides/adding-cname-records [fastly-conditionals]: https://docs.fastly.com/en/guides/using-conditions [fastly-sumologic]: https://developer.fastly.com/reference/api/logging/sumologic/ [fastly-gcs]: https://developer.fastly.com/reference/api/logging/gcs/ ## Import Fastly Service can be imported using their service ID, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fastly_service_v1.demo xxxxxxxxxxxxxxxxxxxx ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "cloned_version",
					Description: `The latest cloned version by the provider. The value gets only set after running ` + "`" + `terraform apply` + "`" + `. The ` + "`" + `dynamicsnippet` + "`" + ` block exports:`,
				},
				resource.Attribute{
					Name:        "snippet_id",
					Description: `The ID of the dynamic snippet. The ` + "`" + `acl` + "`" + ` block exports:`,
				},
				resource.Attribute{
					Name:        "acl_id",
					Description: `The ID of the ACL. The ` + "`" + `dictionary` + "`" + ` block exports:`,
				},
				resource.Attribute{
					Name:        "dictionary_id",
					Description: `The ID of the dictionary. [fastly-s3]: https://docs.fastly.com/en/guides/amazon-s3 [fastly-cname]: https://docs.fastly.com/en/guides/adding-cname-records [fastly-conditionals]: https://docs.fastly.com/en/guides/using-conditions [fastly-sumologic]: https://developer.fastly.com/reference/api/logging/sumologic/ [fastly-gcs]: https://developer.fastly.com/reference/api/logging/gcs/ ## Import Fastly Service can be imported using their service ID, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import fastly_service_v1.demo xxxxxxxxxxxxxxxxxxxx ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fastly_user_v1",
			Category:         "Resources",
			ShortDescription: `Provides a Fastly User`,
			Description:      ``,
			Keywords: []string{
				"user",
				"v1",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "login",
					Description: `(Required, Forces new resource) The email address, which is the login name, of the User.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The real life name of the user.`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `(Optional) The role of this user. Can be ` + "`" + `user` + "`" + ` (the default), ` + "`" + `billing` + "`" + `, ` + "`" + `engineer` + "`" + `, or ` + "`" + `superuser` + "`" + `. For detailed information on the abilities granted to each role, see [Fastly's Documentation on User roles](https://docs.fastly.com/en/guides/configuring-user-roles-and-permissions#user-roles-and-what-they-can-do). ## Attributes Reference In addition to the arguments listed above, the following attributes are exported:`,
				},
			},
			Attributes: []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"fastly_service_acl_entries_v1":             0,
		"fastly_service_dictionary_items_v1":        1,
		"fastly_service_dynamic_snippet_content_v1": 2,
		"fastly_service_v1":                         3,
		"fastly_user_v1":                            4,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
