package aws

import (
	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*Resource{

		&resource.Resource{
			Name:             "",
			Type:             "circonus_check",
			Category:         "Resources",
			ShortDescription: `Manages a Circonus check.`,
			Description:      ``,
			Keywords: []string{
				"check",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "active",
					Description: `(Optional) Whether or not the check is enabled or not (default ` + "`" + `true` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "caql",
					Description: `(Optional) A [Circonus Analytics Query Language (CAQL)](https://login.circonus.com/user/docs/CAQL) check. See below for details on how to configure a ` + "`" + `caql` + "`" + ` check.`,
				},
				resource.Attribute{
					Name:        "cloudwatch",
					Description: `(Optional) A [CloudWatch check](https://login.circonus.com/user/docs/Data/CheckTypes/CloudWatch) check. See below for details on how to configure a ` + "`" + `cloudwatch` + "`" + ` check.`,
				},
				resource.Attribute{
					Name:        "collector",
					Description: `(Required) A collector ID. The collector(s) that are responsible for running a ` + "`" + `circonus_check` + "`" + `. The ` + "`" + `id` + "`" + ` can be the Circonus ID for a Circonus collector (a.k.a. "broker") running in the cloud or an enterprise collector running in your datacenter. One collection of metrics will be automatically created for each ` + "`" + `collector` + "`" + ` specified.`,
				},
				resource.Attribute{
					Name:        "consul",
					Description: `(Optional) A native Consul check. See below for details on how to configure a ` + "`" + `consul` + "`" + ` check.`,
				},
				resource.Attribute{
					Name:        "http",
					Description: `(Optional) A poll-based HTTP check. See below for details on how to configure the ` + "`" + `http` + "`" + ` check.`,
				},
				resource.Attribute{
					Name:        "httptrap",
					Description: `(Optional) An push-based HTTP check. This check method expects clients to send a specially crafted HTTP JSON payload. See below for details on how to configure the ` + "`" + `httptrap` + "`" + ` check.`,
				},
				resource.Attribute{
					Name:        "icmp_ping",
					Description: `(Optional) An ICMP ping check. See below for details on how to configure the ` + "`" + `icmp_ping` + "`" + ` check.`,
				},
				resource.Attribute{
					Name:        "json",
					Description: `(Optional) A JSON check. See below for details on how to configure the ` + "`" + `json` + "`" + ` check.`,
				},
				resource.Attribute{
					Name:        "metric",
					Description: `(Required) A list of one or more ` + "`" + `metric` + "`" + ` configurations. All metrics obtained from this check instance will be available as individual metric streams. See below for a list of supported ` + "`" + `metric` + "`" + ` attrbutes.`,
				},
				resource.Attribute{
					Name:        "metric_limit",
					Description: `(Optional) Setting a metric limit will tell the Circonus backend to periodically look at the check to see if there are additional metrics the collector has seen that we should collect. It will not reactivate metrics previously collected and then marked as inactive. Values are ` + "`" + `0` + "`" + ` to disable, ` + "`" + `-1` + "`" + ` to enable all metrics or ` + "`" + `N+` + "`" + ` to collect up to the value ` + "`" + `N` + "`" + ` (both ` + "`" + `-1` + "`" + ` and ` + "`" + `N+` + "`" + ` can not exceed other account restrictions).`,
				},
				resource.Attribute{
					Name:        "mysql",
					Description: `(Optional) A MySQL check. See below for details on how to configure the ` + "`" + `mysql` + "`" + ` check.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the check that will be displayed in the web interface.`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `(Optional) Notes about this check.`,
				},
				resource.Attribute{
					Name:        "period",
					Description: `(Optional) The period between each time the check is made in seconds.`,
				},
				resource.Attribute{
					Name:        "postgresql",
					Description: `(Optional) A PostgreSQL check. See below for details on how to configure the ` + "`" + `postgresql` + "`" + ` check.`,
				},
				resource.Attribute{
					Name:        "statsd",
					Description: `(Optional) A statsd check. See below for details on how to configure the ` + "`" + `statsd` + "`" + ` check.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A list of tags assigned to this check.`,
				},
				resource.Attribute{
					Name:        "target",
					Description: `(Required) A string containing the location of the thing being checked. This value changes based on the check type. For example, for an ` + "`" + `http` + "`" + ` check type this would be the URL you're checking. For a DNS check it would be the hostname you wanted to look up.`,
				},
				resource.Attribute{
					Name:        "tcp",
					Description: `(Optional) A TCP check. See below for details on how to configure the ` + "`" + `tcp` + "`" + ` check (includes TLS support).`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `(Optional) A floating point number representing the maximum number of seconds this check should wait for a result. Defaults to ` + "`" + `10.0` + "`" + `. ## Supported ` + "`" + `metric` + "`" + ` Attributes The following attributes are available within a ` + "`" + `metric` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "active",
					Description: `(Optional) Whether or not the metric is active or not. Defaults to ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the metric. A string containing freeform text.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A list of tags assigned to the metric.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) A string containing either ` + "`" + `numeric` + "`" + `, ` + "`" + `text` + "`" + `, ` + "`" + `histogram` + "`" + `, ` + "`" + `composite` + "`" + `, or ` + "`" + `caql` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "units",
					Description: `(Optional) The unit of measurement the metric represents (e.g., bytes, seconds, milliseconds). A string containing freeform text. ## Supported Check Types Circonus supports a variety of different checks. Each check type has its own set of options that must be configured. Each check type conflicts with every other check type (i.e. a ` + "`" + `circonus_check` + "`" + ` configured for a ` + "`" + `json` + "`" + ` check will conflict with all other check types, therefore a ` + "`" + `postgresql` + "`" + ` check must be a different ` + "`" + `circonus_check` + "`" + ` resource). ### ` + "`" + `caql` + "`" + ` Check Type Attributes`,
				},
				resource.Attribute{
					Name:        "query",
					Description: `(Required) The [CAQL Query](https://login.circonus.com/user/docs/caql_reference) to run. Available metrics depend on the payload returned in the ` + "`" + `caql` + "`" + ` check. See the [` + "`" + `caql` + "`" + ` check type](https://login.circonus.com/resources/api/calls/check_bundle) for additional details. ### ` + "`" + `cloudwatch` + "`" + ` Check Type Attributes`,
				},
				resource.Attribute{
					Name:        "api_key",
					Description: `(Required) The AWS access key. If this value is not explicitly set, this value is populated by the environment variable ` + "`" + `AWS_ACCESS_KEY_ID` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "api_secret",
					Description: `(Required) The AWS secret key. If this value is not explicitly set, this value is populated by the environment variable ` + "`" + `AWS_SECRET_ACCESS_KEY` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "dimmensions",
					Description: `(Required) A map of the CloudWatch dimmensions to include in the check.`,
				},
				resource.Attribute{
					Name:        "metric",
					Description: `(Required) A list of metric names to collect in this check.`,
				},
				resource.Attribute{
					Name:        "namespace",
					Description: `(Required) The namespace to pull parameters from.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Required) The AWS URL to pull from. This should be set to the region-specific endpoint (e.g. prefer ` + "`" + `https://monitoring.us-east-1.amazonaws.com` + "`" + ` over ` + "`" + `https://monitoring.amazonaws.com` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Optional) The version of the Cloudwatch API to use. Defaults to ` + "`" + `2010-08-01` + "`" + `. Available metrics depend on the payload returned in the ` + "`" + `cloudwatch` + "`" + ` check. See the [` + "`" + `cloudwatch` + "`" + ` check type](https://login.circonus.com/resources/api/calls/check_bundle) for additional details. The ` + "`" + `circonus_check` + "`" + ` ` + "`" + `period` + "`" + ` attribute must be set to either ` + "`" + `60s` + "`" + ` or ` + "`" + `300s` + "`" + ` for CloudWatch metrics. Example CloudWatch check (partial metrics collection): ` + "`" + `` + "`" + `` + "`" + `hcl variable "cloudwatch_rds_tags" { type = "list" default = [ "app:postgresql", "app:rds", "source:cloudwatch", ] } resource "circonus_check" "rds_metrics" { active = true name = "Terraform test: RDS Metrics via CloudWatch" notes = "Collect RDS metrics" period = "60s" collector { id = "/broker/1" } cloudwatch { dimmensions = { DBInstanceIdentifier = "my-db-name", } metric = [ "CPUUtilization", "DatabaseConnections", ] namespace = "AWS/RDS" url = "https://monitoring.us-east-1.amazonaws.com" } metric { name = "CPUUtilization" tags = [ "${var.cloudwatch_rds_tags}" ] type = "numeric" unit = "%" } metric { name = "DatabaseConnections" tags = [ "${var.cloudwatch_rds_tags}" ] type = "numeric" unit = "connections" } } ` + "`" + `` + "`" + `` + "`" + ` ### ` + "`" + `consul` + "`" + ` Check Type Attributes`,
				},
				resource.Attribute{
					Name:        "acl_token",
					Description: `(Optional) An ACL Token authenticate the API request. When an ACL Token is set, this value is transmitted as an HTTP Header in order to not show up in any logs. The default value is an empty string.`,
				},
				resource.Attribute{
					Name:        "allow_stale",
					Description: `(Optional) A boolean value that indicates whether or not this check should require the health information come from the Consul leader node. For scalability reasons, this value defaults to ` + "`" + `false` + "`" + `. See below for details on detecting the staleness of health information.`,
				},
				resource.Attribute{
					Name:        "ca_chain",
					Description: `(Optional) A path to a file containing all the certificate authorities that should be loaded to validate the remote certificate (required when ` + "`" + `http_addr` + "`" + ` is a TLS-enabled endpoint).`,
				},
				resource.Attribute{
					Name:        "certificate_file",
					Description: `(Optional) A path to a file containing the client certificate that will be presented to the remote server (required when ` + "`" + `http_addr` + "`" + ` is a TLS-enabled endpoint).`,
				},
				resource.Attribute{
					Name:        "check_blacklist",
					Description: `(Optional) A list of check names to exclude from the result of checks (i.e. no metrics will be generated by whose check name is in the ` + "`" + `check_blacklist` + "`" + `). This blacklist is applied to the ` + "`" + `node` + "`" + `, ` + "`" + `service` + "`" + `, and ` + "`" + `state` + "`" + ` check modes.`,
				},
				resource.Attribute{
					Name:        "ciphers",
					Description: `(Optional) A list of ciphers to be used in the TLS protocol (only used when ` + "`" + `http_addr` + "`" + ` is a TLS-enabled endpoint).`,
				},
				resource.Attribute{
					Name:        "dc",
					Description: `(Optional) Explicitly name the Consul datacenter to use. The default value is an empty string. When an empty value is specified, the Consul datacenter of the agent at the ` + "`" + `http_addr` + "`" + ` is implicitly used.`,
				},
				resource.Attribute{
					Name:        "headers",
					Description: `(Optional) A map of the HTTP headers to be sent when executing the check. NOTE: the ` + "`" + `headers` + "`" + ` attribute is processed last and will takes precidence over any other derived value that is transmitted as an HTTP header to Consul (i.e. it is possible to override the ` + "`" + `acl_token` + "`" + ` by setting a headers value).`,
				},
				resource.Attribute{
					Name:        "http_addr",
					Description: `(Optional) The Consul HTTP endpoint to to query for health information. The default value is ` + "`" + `http://consul.service.consul:8500` + "`" + `. The scheme must change from ` + "`" + `http` + "`" + ` to ` + "`" + `https` + "`" + ` when the endpoint has been TLS-enabled.`,
				},
				resource.Attribute{
					Name:        "key_file",
					Description: `(Optional) A path to a file containing key to be used in conjunction with the cilent certificate (required when ` + "`" + `http_addr` + "`" + ` is a TLS-enabled endpoint).`,
				},
				resource.Attribute{
					Name:        "node",
					Description: `(Optional) Check the health of this node. The value can be either a Consul Node ID (Consul Version >= 0.7.4) or Node Name. See also the ` + "`" + `service_blacklist` + "`" + `, ` + "`" + `node_blacklist` + "`" + `, and ` + "`" + `check_blacklist` + "`" + ` attributes. This attribute conflicts with the ` + "`" + `service` + "`" + ` and ` + "`" + `state` + "`" + ` attributes.`,
				},
				resource.Attribute{
					Name:        "node_blacklist",
					Description: `(Optional) A list of node IDs or node names to exclude from the results of checks (i.e. no metrics will be generated from nodes in the ` + "`" + `node_blacklist` + "`" + `). This blacklist is applied to the ` + "`" + `node` + "`" + `, ` + "`" + `service` + "`" + `, and ` + "`" + `state` + "`" + ` check modes.`,
				},
				resource.Attribute{
					Name:        "service",
					Description: `(Optional) Check the cluster-wide health of this named service. See also the ` + "`" + `service_blacklist` + "`" + `, ` + "`" + `node_blacklist` + "`" + `, and ` + "`" + `check_blacklist` + "`" + ` attributes. This attribute conflicts with the ` + "`" + `node` + "`" + ` and ` + "`" + `state` + "`" + ` attributes.`,
				},
				resource.Attribute{
					Name:        "service_blacklist",
					Description: `(Optional) A list of service names to exclude from the result of checks (i.e. no metrics will be generated by services whose service name is in the ` + "`" + `service_blacklist` + "`" + `). This blacklist is applied to the ` + "`" + `node` + "`" + `, ` + "`" + `service` + "`" + `, and ` + "`" + `state` + "`" + ` check modes.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `(Optional) A Circonus check to monitor Consul checks across the entire Consul cluster. This value may be either ` + "`" + `passing` + "`" + `, ` + "`" + `warning` + "`" + `, or ` + "`" + `critical` + "`" + `. This ` + "`" + `consul` + "`" + ` check mode is intended to act as the cluster check of last resort. This check type is useful when first starting and is intended to act as a check of last resort before transitioning to explicitly defined checks for individual services or nodes. The metrics returned from check will be sorted based on the ` + "`" + `CreateIndex` + "`" + ` of the entry in order to have a stable set of metrics in the array of returned values. See also the ` + "`" + `service_blacklist` + "`" + `, ` + "`" + `node_blacklist` + "`" + `, and ` + "`" + `check_blacklist` + "`" + ` attributes. This attribute conflicts with the ` + "`" + `node` + "`" + ` and ` + "`" + `state` + "`" + ` attributes. Available metrics depend on the consul check being performed (` + "`" + `node` + "`" + `, ` + "`" + `service` + "`" + `, or ` + "`" + `state` + "`" + `). In addition to the data avilable from the endpoints, the ` + "`" + `consul` + "`" + ` check also returns a set of metrics that are a variant of: ` + "`" + `{Num,Pct}{,Passing,Warning,Critical}{Checks,Nodes,Services}` + "`" + ` (see the ` + "`" + `GLOB_BRACE` + "`" + ` section of your local ` + "`" + `glob(3)` + "`" + ` documentation). Example Consul check (partial metrics collection): ` + "`" + `` + "`" + `` + "`" + `hcl resource "circonus_check" "consul_server" { active = true name = "%s" period = "60s" collector { # Collector ID must be an Enterprise broker able to reach the Consul agent # listed in ` + "`" + `http_addr` + "`" + `. id = "/broker/2110" } consul { service = "consul" # Other consul check modes: # node = "consul1" # state = "critical" } metric { name = "NumNodes" tags = [ "source:consul", "lifecycle:unittest" ] type = "numeric" } metric { name = "LastContact" tags = [ "source:consul", "lifecycle:unittest" ] type = "numeric" unit = "seconds" } metric { name = "Index" tags = [ "source:consul", "lifecycle:unittest" ] type = "numeric" unit = "transactions" } metric { name = "KnownLeader" tags = [ "source:consul", "lifecycle:unittest" ] type = "text" } tags = [ "source:consul", "lifecycle:unittest" ] } ` + "`" + `` + "`" + `` + "`" + ` ### ` + "`" + `http` + "`" + ` Check Type Attributes`,
				},
				resource.Attribute{
					Name:        "auth_method",
					Description: `(Optional) HTTP Authentication method to use. When set must be one of the values ` + "`" + `Basic` + "`" + `, ` + "`" + `Digest` + "`" + `, or ` + "`" + `Auto` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "auth_password",
					Description: `(Optional) The password to use during authentication.`,
				},
				resource.Attribute{
					Name:        "auth_user",
					Description: `(Optional) The user to authenticate as.`,
				},
				resource.Attribute{
					Name:        "body_regexp",
					Description: `(Optional) This regular expression is matched against the body of the response. If a match is not found, the check will be marked as "bad."`,
				},
				resource.Attribute{
					Name:        "ca_chain",
					Description: `(Optional) A path to a file containing all the certificate authorities that should be loaded to validate the remote certificate (for TLS checks).`,
				},
				resource.Attribute{
					Name:        "certificate_file",
					Description: `(Optional) A path to a file containing the client certificate that will be presented to the remote server (for TLS checks).`,
				},
				resource.Attribute{
					Name:        "ciphers",
					Description: `(Optional) A list of ciphers to be used in the TLS protocol (for HTTPS checks).`,
				},
				resource.Attribute{
					Name:        "code",
					Description: `(Optional) The HTTP code that is expected. If the code received does not match this regular expression, the check is marked as "bad."`,
				},
				resource.Attribute{
					Name:        "extract",
					Description: `(Optional) This regular expression is matched against the body of the response globally. The first capturing match is the key and the second capturing match is the value. Each key/value extracted is registered as a metric for the check.`,
				},
				resource.Attribute{
					Name:        "headers",
					Description: `(Optional) A map of the HTTP headers to be sent when executing the check.`,
				},
				resource.Attribute{
					Name:        "key_file",
					Description: `(Optional) A path to a file containing key to be used in conjunction with the cilent certificate (for TLS checks).`,
				},
				resource.Attribute{
					Name:        "method",
					Description: `(Optional) The HTTP Method to use. Defaults to ` + "`" + `GET` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "payload",
					Description: `(Optional) The information transferred as the payload of an HTTP request.`,
				},
				resource.Attribute{
					Name:        "read_limit",
					Description: `(Optional) Sets an approximate limit on the data read (` + "`" + `0` + "`" + ` means no limit). Default ` + "`" + `0` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "redirects",
					Description: `(Optional) The maximum number of HTTP ` + "`" + `Location` + "`" + ` header redirects to follow. Default ` + "`" + `0` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Required) The target for this ` + "`" + `json` + "`" + ` check. The ` + "`" + `url` + "`" + ` must include the scheme, host, port (optional), and path to use (e.g. ` + "`" + `https://app1.example.org/healthz` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Optional) The HTTP version to use. Defaults to ` + "`" + `1.1` + "`" + `. Available metrics include: ` + "`" + `body_match` + "`" + `, ` + "`" + `bytes` + "`" + `, ` + "`" + `cert_end` + "`" + `, ` + "`" + `cert_end_in` + "`" + `, ` + "`" + `cert_error` + "`" + `, ` + "`" + `cert_issuer` + "`" + `, ` + "`" + `cert_start` + "`" + `, ` + "`" + `cert_subject` + "`" + `, ` + "`" + `code` + "`" + `, ` + "`" + `duration` + "`" + `, ` + "`" + `truncated` + "`" + `, ` + "`" + `tt_connect` + "`" + `, and ` + "`" + `tt_firstbyte` + "`" + `. See the [` + "`" + `http` + "`" + ` check type](https://login.circonus.com/resources/api/calls/check_bundle) for additional details. ### ` + "`" + `httptrap` + "`" + ` Check Type Attributes`,
				},
				resource.Attribute{
					Name:        "async_metrics",
					Description: `(Optional) Boolean value specifies whether or not httptrap metrics are logged immediately or held until the status message is to be emitted. Default ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "secret",
					Description: `(Optional) Specify the secret with which metrics may be submitted. Available metrics depend on the payload returned in the ` + "`" + `httptrap` + "`" + ` doc. See the [` + "`" + `httptrap` + "`" + ` check type](https://login.circonus.com/resources/api/calls/check_bundle) for additional details. ### ` + "`" + `json` + "`" + ` Check Type Attributes`,
				},
				resource.Attribute{
					Name:        "auth_method",
					Description: `(Optional) HTTP Authentication method to use. When set must be one of the values ` + "`" + `Basic` + "`" + `, ` + "`" + `Digest` + "`" + `, or ` + "`" + `Auto` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "auth_password",
					Description: `(Optional) The password to use during authentication.`,
				},
				resource.Attribute{
					Name:        "auth_user",
					Description: `(Optional) The user to authenticate as.`,
				},
				resource.Attribute{
					Name:        "ca_chain",
					Description: `(Optional) A path to a file containing all the certificate authorities that should be loaded to validate the remote certificate (for TLS checks).`,
				},
				resource.Attribute{
					Name:        "certificate_file",
					Description: `(Optional) A path to a file containing the client certificate that will be presented to the remote server (for TLS checks).`,
				},
				resource.Attribute{
					Name:        "ciphers",
					Description: `(Optional) A list of ciphers to be used in the TLS protocol (for HTTPS checks).`,
				},
				resource.Attribute{
					Name:        "headers",
					Description: `(Optional) A map of the HTTP headers to be sent when executing the check.`,
				},
				resource.Attribute{
					Name:        "key_file",
					Description: `(Optional) A path to a file containing key to be used in conjunction with the cilent certificate (for TLS checks).`,
				},
				resource.Attribute{
					Name:        "method",
					Description: `(Optional) The HTTP Method to use. Defaults to ` + "`" + `GET` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional) The TCP Port number to use. Defaults to ` + "`" + `81` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "read_limit",
					Description: `(Optional) Sets an approximate limit on the data read (` + "`" + `0` + "`" + ` means no limit). Default ` + "`" + `0` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "redirects",
					Description: `(Optional) The maximum number of HTTP ` + "`" + `Location` + "`" + ` header redirects to follow. Default ` + "`" + `0` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Required) The target for this ` + "`" + `json` + "`" + ` check. The ` + "`" + `url` + "`" + ` must include the scheme, host, port (optional), and path to use (e.g. ` + "`" + `https://app1.example.org/healthz` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Optional) The HTTP version to use. Defaults to ` + "`" + `1.1` + "`" + `. Available metrics depend on the payload returned in the ` + "`" + `json` + "`" + ` doc. See the [` + "`" + `json` + "`" + ` check type](https://login.circonus.com/resources/api/calls/check_bundle) for additional details. ### ` + "`" + `icmp_ping` + "`" + ` Check Type Attributes The ` + "`" + `icmp_ping` + "`" + ` check requires the ` + "`" + `target` + "`" + ` top-level attribute to be set.`,
				},
				resource.Attribute{
					Name:        "availability",
					Description: `(Optional) The percentage of ping packets that must be returned for this measurement to be considered successful. Defaults to ` + "`" + `100.0` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "count",
					Description: `(Optional) The number of ICMP ping packets to send. Defaults to ` + "`" + `5` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "interval",
					Description: `(Optional) Interval between packets. Defaults to ` + "`" + `2s` + "`" + `. Available metrics include: ` + "`" + `available` + "`" + `, ` + "`" + `average` + "`" + `, ` + "`" + `count` + "`" + `, ` + "`" + `maximum` + "`" + `, and ` + "`" + `minimum` + "`" + `. See the [` + "`" + `ping_icmp` + "`" + ` check type](https://login.circonus.com/resources/api/calls/check_bundle) for additional details. ### ` + "`" + `mysql` + "`" + ` Check Type Attributes The ` + "`" + `mysql` + "`" + ` check requires the ` + "`" + `target` + "`" + ` top-level attribute to be set.`,
				},
				resource.Attribute{
					Name:        "dsn",
					Description: `(Required) The [MySQL DSN/connect string](https://github.com/go-sql-driver/mysql/blob/master/README.md) to use to talk to MySQL.`,
				},
				resource.Attribute{
					Name:        "query",
					Description: `(Required) The SQL query to execute. ### ` + "`" + `postgresql` + "`" + ` Check Type Attributes The ` + "`" + `postgresql` + "`" + ` check requires the ` + "`" + `target` + "`" + ` top-level attribute to be set.`,
				},
				resource.Attribute{
					Name:        "dsn",
					Description: `(Required) The [PostgreSQL DSN/connect string](https://www.postgresql.org/docs/current/static/libpq-connect.html) to use to talk to PostgreSQL.`,
				},
				resource.Attribute{
					Name:        "query",
					Description: `(Required) The SQL query to execute. Available metric names are dependent on the output of the ` + "`" + `query` + "`" + ` being run. ### ` + "`" + `statsd` + "`" + ` Check Type Attributes`,
				},
				resource.Attribute{
					Name:        "source_ip",
					Description: `(Required) Any statsd messages from this IP address (IPv4 or IPv6) will be associated with this check. Available metrics depend on the metrics sent to the ` + "`" + `statsd` + "`" + ` check. ### ` + "`" + `tcp` + "`" + ` Check Type Attributes`,
				},
				resource.Attribute{
					Name:        "banner_regexp",
					Description: `(Optional) This regular expression is matched against the response banner. If a match is not found, the check will be marked as bad.`,
				},
				resource.Attribute{
					Name:        "ca_chain",
					Description: `(Optional) A path to a file containing all the certificate authorities that should be loaded to validate the remote certificate (for TLS checks).`,
				},
				resource.Attribute{
					Name:        "certificate_file",
					Description: `(Optional) A path to a file containing the client certificate that will be presented to the remote server (for TLS checks).`,
				},
				resource.Attribute{
					Name:        "ciphers",
					Description: `(Optional) A list of ciphers to be used in the TLS protocol (for HTTPS checks).`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `(Required) Hostname or IP address of the host to connect to.`,
				},
				resource.Attribute{
					Name:        "key_file",
					Description: `(Optional) A path to a file containing key to be used in conjunction with the cilent certificate (for TLS checks).`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Required) Integer specifying the port on which the management interface can be reached.`,
				},
				resource.Attribute{
					Name:        "tls",
					Description: `(Optional) When enabled establish a TLS connection. Available metrics include: ` + "`" + `banner` + "`" + `, ` + "`" + `banner_match` + "`" + `, ` + "`" + `cert_end` + "`" + `, ` + "`" + `cert_end_in` + "`" + `, ` + "`" + `cert_error` + "`" + `, ` + "`" + `cert_issuer` + "`" + `, ` + "`" + `cert_start` + "`" + `, ` + "`" + `cert_subject` + "`" + `, ` + "`" + `duration` + "`" + `, ` + "`" + `tt_connect` + "`" + `, ` + "`" + `tt_firstbyte` + "`" + `. See the [` + "`" + `tcp` + "`" + ` check type](https://login.circonus.com/resources/api/calls/check_bundle) for additional details. Sample ` + "`" + `tcp` + "`" + ` check: ` + "`" + `` + "`" + `` + "`" + `hcl resource "circonus_check" "tcp_check" { name = "TCP and TLS check" notes = "Obtains the connect time and TTL for the TLS cert" period = "60s" collector { id = "/broker/1" } tcp { host = "127.0.0.1" port = 443 tls = true } metric { name = "cert_end_in" tags = [ "${var.tcp_check_tags}" ] type = "numeric" unit = "seconds" } metric { name = "tt_connect" tags = [ "${var.tcp_check_tags}" ] type = "numeric" unit = "miliseconds" } tags = [ "${var.tcp_check_tags}" ] } ` + "`" + `` + "`" + `` + "`" + ` ## Out Parameters`,
				},
				resource.Attribute{
					Name:        "check_by_collector",
					Description: `Maps the ID of the collector (` + "`" + `collector_id` + "`" + `, the map key) to the ` + "`" + `check_id` + "`" + ` (value) that is registered to a collector.`,
				},
				resource.Attribute{
					Name:        "check_id",
					Description: `If there is only one ` + "`" + `collector` + "`" + ` specified for the check, this value will be populated with the ` + "`" + `check_id` + "`" + `. If more than one ` + "`" + `collector` + "`" + ` is specified in the check, then this value will be an empty string. ` + "`" + `check_by_collector` + "`" + ` will always be populated.`,
				},
				resource.Attribute{
					Name:        "checks",
					Description: `List of ` + "`" + `check_id` + "`" + `s created by this ` + "`" + `circonus_check` + "`" + `. There is one element in this list per collector specified in the check.`,
				},
				resource.Attribute{
					Name:        "created",
					Description: `UNIX time at which this check was created.`,
				},
				resource.Attribute{
					Name:        "last_modified",
					Description: `UNIX time at which this check was last modified.`,
				},
				resource.Attribute{
					Name:        "last_modified_by",
					Description: `User ID in Circonus who modified this check last.`,
				},
				resource.Attribute{
					Name:        "reverse_connect_urls",
					Description: `Only relevant to Circonus support.`,
				},
				resource.Attribute{
					Name:        "uuids",
					Description: `List of Check ` + "`" + `uuid` + "`" + `s created by this ` + "`" + `circonus_check` + "`" + `. There is one element in this list per collector specified in the check. ## Import Example ` + "`" + `circonus_check` + "`" + ` supports importing resources. Supposing the following Terraform (and that the referenced [` + "`" + `circonus_metric` + "`" + `](metric.html) has already been imported): ` + "`" + `` + "`" + `` + "`" + `hcl provider "circonus" { alias = "b8fec159-f9e5-4fe6-ad2c-dc1ec6751586" } resource "circonus_metric" "used" { name = "_usage` + "`" + `0` + "`" + `_used" type = "numeric" } resource "circonus_check" "usage" { collector { id = "/broker/1" } json { url = "https://api.circonus.com/account/current" headers = { "Accept" = "application/json" "X-Circonus-App-Name" = "TerraformCheck" "X-Circonus-Auth-Token" = "${var.api_token}" } } metric { name = "${circonus_metric.used.name}" type = "${circonus_metric.used.type}" } } ` + "`" + `` + "`" + `` + "`" + ` It is possible to import a ` + "`" + `circonus_check` + "`" + ` resource with the following command: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import circonus_check.usage ID ` + "`" + `` + "`" + `` + "`" + ` Where ` + "`" + `ID` + "`" + ` is the ` + "`" + `_cid` + "`" + ` or Circonus ID of the Check Bundle (e.g. ` + "`" + `/check_bundle/12345` + "`" + `) and ` + "`" + `circonus_check.usage` + "`" + ` is the name of the resource whose state will be populated as a result of the command.`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "circonus_contact_group",
			Category:         "Resources",
			ShortDescription: `Manages a Circonus Contact Group.`,
			Description:      ``,
			Keywords: []string{
				"contact",
				"group",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "aggregation_window",
					Description: `(Optional) The aggregation window for batching up alert notifications.`,
				},
				resource.Attribute{
					Name:        "alert_option",
					Description: `(Optional) There is one ` + "`" + `alert_option` + "`" + ` per severity, where severity can be any number between 1 (high) and 5 (low). If configured, the alerting system will remind or escalate alerts to further contact groups if an alert sent to this contact group is not acknowledged or resolved. See below for details.`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `(Optional) Zero or more ` + "`" + `email` + "`" + ` attributes may be present to dispatch email to Circonus users by referencing their user ID, or by specifying an email address. See below for details on supported attributes.`,
				},
				resource.Attribute{
					Name:        "http",
					Description: `(Optional) Zero or more ` + "`" + `http` + "`" + ` attributes may be present to dispatch [Webhook/HTTP requests](https://login.circonus.com/user/docs/Alerting/ContactGroups#WebhookNotifications) by Circonus. See below for details on supported attributes.`,
				},
				resource.Attribute{
					Name:        "irc",
					Description: `(Optional) Zero or more ` + "`" + `irc` + "`" + ` attributes may be present to dispatch IRC notifications to users. See below for details on supported attributes.`,
				},
				resource.Attribute{
					Name:        "long_message",
					Description: `(Optional) The bulk of the message used in long form alert messages.`,
				},
				resource.Attribute{
					Name:        "long_subject",
					Description: `(Optional) The subject used in long form alert messages.`,
				},
				resource.Attribute{
					Name:        "long_summary",
					Description: `(Optional) The brief summary used in long form alert messages.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the contact group.`,
				},
				resource.Attribute{
					Name:        "pager_duty",
					Description: `(Optional) Zero or more ` + "`" + `pager_duty` + "`" + ` attributes may be present to dispatch to [Pager Duty teams](https://login.circonus.com/user/docs/Alerting/ContactGroups#PagerDutyOptions). See below for details on supported attributes.`,
				},
				resource.Attribute{
					Name:        "short_message",
					Description: `(Optional) The subject used in short form alert messages.`,
				},
				resource.Attribute{
					Name:        "short_summary",
					Description: `(Optional) The brief summary used in short form alert messages.`,
				},
				resource.Attribute{
					Name:        "slack",
					Description: `(Optional) Zero or more ` + "`" + `slack` + "`" + ` attributes may be present to dispatch to Slack teams. See below for details on supported attributes.`,
				},
				resource.Attribute{
					Name:        "sms",
					Description: `(Optional) Zero or more ` + "`" + `sms` + "`" + ` attributes may be present to dispatch SMS messages to Circonus users by referencing their user ID, or by specifying an SMS Phone Number. See below for details on supported attributes.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A list of tags attached to the Contact Group.`,
				},
				resource.Attribute{
					Name:        "victorops",
					Description: `(Optional) Zero or more ` + "`" + `victorops` + "`" + ` attributes may be present to dispatch to [VictorOps teams](https://login.circonus.com/user/docs/Alerting/ContactGroups#VictorOps). See below for details on supported attributes. ## Supported Contact Group ` + "`" + `alert_option` + "`" + ` Attributes`,
				},
				resource.Attribute{
					Name:        "escalate_after",
					Description: `(Optional) How long to wait before escalating an alert that is received at a given severity.`,
				},
				resource.Attribute{
					Name:        "escalate_to",
					Description: `(Optional) The Contact Group ID who will receive the escalation.`,
				},
				resource.Attribute{
					Name:        "reminder",
					Description: `(Optional) If specified, reminders will be sent after a user configurable number of minutes for open alerts.`,
				},
				resource.Attribute{
					Name:        "severity",
					Description: `(Required) An ` + "`" + `alert_option` + "`" + ` must be assigned to a given severity level. Valid severity levels range from 1 (highest severity) to 5 (lowest severity). ## Supported Contact Group ` + "`" + `email` + "`" + ` Attributes Either an ` + "`" + `address` + "`" + ` or ` + "`" + `user` + "`" + ` attribute is required.`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `(Optional) A well formed email address.`,
				},
				resource.Attribute{
					Name:        "user",
					Description: `(Optional) An email will be sent to the email address of record for the corresponding user ID (e.g. ` + "`" + `/user/1234` + "`" + `). A ` + "`" + `user` + "`" + `'s email address is automatically maintained and kept up to date by the recipient, whereas an ` + "`" + `address` + "`" + ` provides no automatic layer of indirection for keeping the information accurate (including LDAP and SAML-based authentication mechanisms). ## Supported Contact Group ` + "`" + `http` + "`" + ` Attributes`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `(Required) URL to send a webhook request to.`,
				},
				resource.Attribute{
					Name:        "format",
					Description: `(Optional) The payload of the request is a JSON-encoded payload when the ` + "`" + `format` + "`" + ` is set to ` + "`" + `json` + "`" + ` (the default). The alternate payload encoding is ` + "`" + `params` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "method",
					Description: `(Optional) The HTTP verb to use when making a request. Either ` + "`" + `GET` + "`" + ` or ` + "`" + `POST` + "`" + ` may be specified. The default verb is ` + "`" + `POST` + "`" + `. ## Supported Contact Group ` + "`" + `irc` + "`" + ` Attributes`,
				},
				resource.Attribute{
					Name:        "user",
					Description: `(Required) When a user has configured IRC on their user account, they will receive an IRC notification. ## Supported Contact Group ` + "`" + `pager_duty` + "`" + ` Attributes`,
				},
				resource.Attribute{
					Name:        "contact_group_fallback",
					Description: `(Optional) If there is a problem contacting PagerDuty, relay the notification automatically to the specified Contact Group (e.g. ` + "`" + `/contact_group/1234` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "service_key",
					Description: `(Required) The PagerDuty Service Key.`,
				},
				resource.Attribute{
					Name:        "webhook_url",
					Description: `(Required) The PagerDuty webhook URL that PagerDuty uses to notify Circonus of acknowledged actions. ## Supported Contact Group ` + "`" + `slack` + "`" + ` Attributes`,
				},
				resource.Attribute{
					Name:        "contact_group_fallback",
					Description: `(Optional) If there is a problem contacting Slack, relay the notification automatically to the specified Contact Group (e.g. ` + "`" + `/contact_group/1234` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "buttons",
					Description: `(Optional) Slack notifications can have acknowledgement buttons built into the notification message itself when enabled. Defaults to ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "channel",
					Description: `(Required) Specify what Slack channel Circonus should send alerts to.`,
				},
				resource.Attribute{
					Name:        "team",
					Description: `(Required) Specify what Slack team Circonus should look in for the aforementioned ` + "`" + `channel` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Optional) Specify the username Circonus should advertise itself as in Slack. Defaults to ` + "`" + `Circonus` + "`" + `. ## Supported Contact Group ` + "`" + `sms` + "`" + ` Attributes Either an ` + "`" + `address` + "`" + ` or ` + "`" + `user` + "`" + ` attribute is required.`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `(Optional) SMS Phone Number to send a short notification to.`,
				},
				resource.Attribute{
					Name:        "user",
					Description: `(Optional) An SMS page will be sent to the phone number of record for the corresponding user ID (e.g. ` + "`" + `/user/1234` + "`" + `). A ` + "`" + `user` + "`" + `'s phone number is automatically maintained and kept up to date by the recipient, whereas an ` + "`" + `address` + "`" + ` provides no automatic layer of indirection for keeping the information accurate (including LDAP and SAML-based authentication mechanisms). ## Supported Contact Group ` + "`" + `victorops` + "`" + ` Attributes`,
				},
				resource.Attribute{
					Name:        "contact_group_fallback",
					Description: `(Optional) If there is a problem contacting VictorOps, relay the notification automatically to the specified Contact Group (e.g. ` + "`" + `/contact_group/1234` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "api_key",
					Description: `(Required) The API Key for talking with VictorOps.`,
				},
				resource.Attribute{
					Name:        "critical",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "info",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "team",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "warning",
					Description: `(Required) ## Supported Contact Group ` + "`" + `xmpp` + "`" + ` Attributes Either an ` + "`" + `address` + "`" + ` or ` + "`" + `user` + "`" + ` attribute is required.`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `(Optional) XMPP address to send a short notification to.`,
				},
				resource.Attribute{
					Name:        "user",
					Description: `(Optional) An XMPP notification will be sent to the XMPP address of record for the corresponding user ID (e.g. ` + "`" + `/user/1234` + "`" + `). ## Import Example ` + "`" + `circonus_contact_group` + "`" + ` supports importing resources. Supposing the following Terraform: ` + "`" + `` + "`" + `` + "`" + `hcl provider "circonus" { alias = "b8fec159-f9e5-4fe6-ad2c-dc1ec6751586" } resource "circonus_contact_group" "myteam" { name = "My Team's Contact Group" email { address = "myteam@example.com" } slack { channel = "#myteam" team = "T024UT03C" } } ` + "`" + `` + "`" + `` + "`" + ` It is possible to import a ` + "`" + `circonus_contact_group` + "`" + ` resource with the following command: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import circonus_contact_group.myteam ID ` + "`" + `` + "`" + `` + "`" + ` Where ` + "`" + `ID` + "`" + ` is the ` + "`" + `_cid` + "`" + ` or Circonus ID of the Contact Group (e.g. ` + "`" + `/contact_group/12345` + "`" + `) and ` + "`" + `circonus_contact_group.myteam` + "`" + ` is the name of the resource whose state will be populated as a result of the command.`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "circonus_graph",
			Category:         "Resources",
			ShortDescription: `Manages a Circonus graph.`,
			Description:      ``,
			Keywords: []string{
				"graph",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of what the graph is for.`,
				},
				resource.Attribute{
					Name:        "graph_style",
					Description: `(Optional) How the graph should be rendered. Valid options are ` + "`" + `area` + "`" + ` or ` + "`" + `line` + "`" + ` (default).`,
				},
				resource.Attribute{
					Name:        "left",
					Description: `(Optional) A map of graph left axis options. Valid values in ` + "`" + `left` + "`" + ` include: ` + "`" + `logarithmic` + "`" + ` can be set to ` + "`" + `0` + "`" + ` (default) or ` + "`" + `1` + "`" + `; ` + "`" + `min` + "`" + ` is the ` + "`" + `min` + "`" + ` Y axis value on the left; and ` + "`" + `max` + "`" + ` is the Y axis max value on the left.`,
				},
				resource.Attribute{
					Name:        "line_style",
					Description: `(Optional) How the line should change between points. Can be either ` + "`" + `stepped` + "`" + ` (default) or ` + "`" + `interpolated` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The title of the graph.`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `(Optional) A place for storing notes about this graph.`,
				},
				resource.Attribute{
					Name:        "right",
					Description: `(Optional) A map of graph right axis options. Valid values in ` + "`" + `right` + "`" + ` include: ` + "`" + `logarithmic` + "`" + ` can be set to ` + "`" + `0` + "`" + ` (default) or ` + "`" + `1` + "`" + `; ` + "`" + `min` + "`" + ` is the ` + "`" + `min` + "`" + ` Y axis value on the right; and ` + "`" + `max` + "`" + ` is the Y axis max value on the right.`,
				},
				resource.Attribute{
					Name:        "metric",
					Description: `(Optional) A list of metric streams to graph. See below for options.`,
				},
				resource.Attribute{
					Name:        "metric_cluster",
					Description: `(Optional) A metric cluster to graph. See below for options.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A list of tags assigned to this graph. ## ` + "`" + `metric` + "`" + ` Configuration An individual metric stream is the underlying source of data points used for visualization in a graph. Either a ` + "`" + `caql` + "`" + ` attribute is required or a ` + "`" + `check` + "`" + ` and ` + "`" + `metric` + "`" + ` must be set. The ` + "`" + `metric` + "`" + ` attribute can have the following options set.`,
				},
				resource.Attribute{
					Name:        "active",
					Description: `(Optional) A boolean if the metric stream is enabled or not.`,
				},
				resource.Attribute{
					Name:        "alpha",
					Description: `(Optional) A floating point number between 0 and 1.`,
				},
				resource.Attribute{
					Name:        "axis",
					Description: `(Optional) The axis that the metric stream will use. Valid options are ` + "`" + `left` + "`" + ` (default) or ` + "`" + `right` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "caql",
					Description: `(Optional) A CAQL formula. Conflicts with the ` + "`" + `check` + "`" + ` and ` + "`" + `metric` + "`" + ` attributes.`,
				},
				resource.Attribute{
					Name:        "check",
					Description: `(Optional) The check that this metric stream belongs to.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `(Optional) A hex-encoded color of the line / area on the graph.`,
				},
				resource.Attribute{
					Name:        "formula",
					Description: `(Optional) Formula that should be aplied to both the values in the graph and the legend.`,
				},
				resource.Attribute{
					Name:        "legend_formula",
					Description: `(Optional) Formula that should be applied to values in the legend.`,
				},
				resource.Attribute{
					Name:        "function",
					Description: `(Optional) What derivative value, if any, should be used. Valid values are: ` + "`" + `gauge` + "`" + ` (default), ` + "`" + `derive` + "`" + `, and ` + "`" + `counter (_stddev)` + "`" + ``,
				},
				resource.Attribute{
					Name:        "metric_type",
					Description: `(Required) The type of the metric. Valid values are: ` + "`" + `numeric` + "`" + `, ` + "`" + `text` + "`" + `, ` + "`" + `histogram` + "`" + `, ` + "`" + `composite` + "`" + `, or ` + "`" + `caql` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) A name which will appear in the graph legend.`,
				},
				resource.Attribute{
					Name:        "metric_name",
					Description: `(Optional) The name of the metric stream within the check to graph.`,
				},
				resource.Attribute{
					Name:        "stack",
					Description: `(Optional) If this metric is to be stacked, which stack set does it belong to (starting at ` + "`" + `0` + "`" + `). ## ` + "`" + `metric_cluster` + "`" + ` Configuration A metric cluster selects multiple metric streams together dynamically using a query language and returns the set of matching metric streams as a single result set to the graph rendering engine.`,
				},
				resource.Attribute{
					Name:        "active",
					Description: `(Optional) A boolean if the metric cluster is enabled or not.`,
				},
				resource.Attribute{
					Name:        "aggregate",
					Description: `(Optional) The aggregate function to apply across this metric cluster to create a single value. Valid values are: ` + "`" + `none` + "`" + ` (default), ` + "`" + `min` + "`" + `, ` + "`" + `max` + "`" + `, ` + "`" + `sum` + "`" + `, ` + "`" + `mean` + "`" + `, or ` + "`" + `geometric_mean` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "axis",
					Description: `(Optional) The axis that the metric cluster will use. Valid options are ` + "`" + `left` + "`" + ` (default) or ` + "`" + `right` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `(Optional) A hex-encoded color of the line / area on the graph. This is a required attribute when ` + "`" + `aggregate` + "`" + ` is specified.`,
				},
				resource.Attribute{
					Name:        "group",
					Description: `(Optional) The ` + "`" + `metric_cluster` + "`" + ` that will provide datapoints for this graph.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) A name which will appear in the graph legend for this metric cluster. ## Import Example ` + "`" + `circonus_graph` + "`" + ` supports importing resources. Supposing the following Terraform (and that the referenced [` + "`" + `circonus_metric` + "`" + `](metric.html) and [` + "`" + `circonus_check` + "`" + `](check.html) have already been imported): ` + "`" + `` + "`" + `` + "`" + `text resource "circonus_graph" "icmp-graph" { name = "Test graph" graph_style = "line" line_style = "stepped" metric { check = "${circonus_check.api_latency.checks[0]}" metric_name = "maximum" metric_type = "numeric" name = "Maximum Latency" axis = "left" } } ` + "`" + `` + "`" + `` + "`" + ` It is possible to import a ` + "`" + `circonus_graph` + "`" + ` resource with the following command: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import circonus_graph.icmp-graph ID ` + "`" + `` + "`" + `` + "`" + ` Where ` + "`" + `ID` + "`" + ` is the ` + "`" + `_cid` + "`" + ` or Circonus ID of the graph (e.g. ` + "`" + `/graph/bd72aabc-90b9-4039-cc30-c9ab838c18f5` + "`" + `) and ` + "`" + `circonus_graph.icmp-graph` + "`" + ` is the name of the resource whose state will be populated as a result of the command.`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "circonus_metric",
			Category:         "Resources",
			ShortDescription: `Manages a Circonus metric.`,
			Description:      ``,
			Keywords: []string{
				"metric",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "active",
					Description: `(Optional) A boolean indicating if the metric is being filtered out at the ` + "`" + `circonus_check` + "`" + `'s collector(s) or not.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the metric. A ` + "`" + `name` + "`" + ` must be unique within a ` + "`" + `circonus_check` + "`" + ` and its meaning is ` + "`" + `circonus_check.type` + "`" + ` specific.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A list of tags assigned to the metric.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) The type of metric. This value must be present and can be one of the following values: ` + "`" + `numeric` + "`" + `, ` + "`" + `text` + "`" + `, ` + "`" + `histogram` + "`" + `, ` + "`" + `composite` + "`" + `, or ` + "`" + `caql` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "unit",
					Description: `(Optional) The unit of measurement for this ` + "`" + `circonus_metric` + "`" + `. ## Import Example ` + "`" + `circonus_metric` + "`" + ` supports importing resources. Supposing the following Terraform: ` + "`" + `` + "`" + `` + "`" + `hcl provider "circonus" { alias = "b8fec159-f9e5-4fe6-ad2c-dc1ec6751586" } resource "circonus_metric" "usage" { name = "_usage` + "`" + `0` + "`" + `_used" type = "numeric" unit = "qty" tags = { source = "circonus" } } ` + "`" + `` + "`" + `` + "`" + ` It is possible to import a ` + "`" + `circonus_metric` + "`" + ` resource with the following command: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import circonus_metric.usage ID ` + "`" + `` + "`" + `` + "`" + ` Where ` + "`" + `ID` + "`" + ` is a random, never before used UUID and ` + "`" + `circonus_metric.usage` + "`" + ` is the name of the resource whose state will be populated as a result of the command.`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "circonus_metric_cluster",
			Category:         "Resources",
			ShortDescription: `Manages a Circonus Metric Cluster.`,
			Description:      ``,
			Keywords: []string{
				"metric",
				"cluster",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A long-form description of the metric cluster.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the metric cluster. This name must be unique across all metric clusters in a given Circonus Account.`,
				},
				resource.Attribute{
					Name:        "query",
					Description: `(Required) One or more ` + "`" + `query` + "`" + ` attributes must be present. Each ` + "`" + `query` + "`" + ` must contain both a ` + "`" + `definition` + "`" + ` and a ` + "`" + `type` + "`" + `. See below for details on supported attributes.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A list of tags attached to the metric cluster. ## Supported Metric Cluster ` + "`" + `query` + "`" + ` Attributes`,
				},
				resource.Attribute{
					Name:        "definition",
					Description: `(Required) The definition of a metric cluster [query](https://login.circonus.com/resources/api/calls/metric_cluster).`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) The query type to execute per metric cluster. Valid query types are: ` + "`" + `average` + "`" + `, ` + "`" + `count` + "`" + `, ` + "`" + `counter` + "`" + `, ` + "`" + `counter2` + "`" + `, ` + "`" + `counter2_stddev` + "`" + `, ` + "`" + `counter_stddev` + "`" + `, ` + "`" + `derive` + "`" + `, ` + "`" + `derive2` + "`" + `, ` + "`" + `derive2_stddev` + "`" + `, ` + "`" + `derive_stddev` + "`" + `, ` + "`" + `histogram` + "`" + `, ` + "`" + `stddev` + "`" + `, ` + "`" + `text` + "`" + `. ## Out parameters`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Metric Cluster. ## Import Example ` + "`" + `circonus_metric_cluster` + "`" + ` supports importing resources. Supposing the following Terraform: ` + "`" + `` + "`" + `` + "`" + `hcl provider "circonus" { alias = "b8fec159-f9e5-4fe6-ad2c-dc1ec6751586" } resource "circonus_metric_cluster" "mymetriccluster" { name = "Metric Cluster for a particular metric in a job" query { definition = "`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "circonus_rule_set",
			Category:         "Resources",
			ShortDescription: `Manages a Circonus rule set.`,
			Description:      ``,
			Keywords: []string{
				"rule",
				"set",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "check",
					Description: `(Required) The Circonus ID that this Rule Set will use to search for a metric stream to alert on.`,
				},
				resource.Attribute{
					Name:        "if",
					Description: `(Required) One or more ordered predicate clauses that describe when Circonus should generate a notification. See below for details on the structure of an ` + "`" + `if` + "`" + ` configuration clause.`,
				},
				resource.Attribute{
					Name:        "link",
					Description: `(Optional) A link to external documentation (or anything else you feel is important) when a notification is sent. This value will show up in email alerts and the Circonus UI.`,
				},
				resource.Attribute{
					Name:        "metric_type",
					Description: `(Optional) The type of metric this rule set will operate on. Valid values are ` + "`" + `numeric` + "`" + ` (the default) and ` + "`" + `text` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `(Optional) Notes about this rule set.`,
				},
				resource.Attribute{
					Name:        "parent",
					Description: `(Optional) A Circonus Metric ID that, if specified and active with a severity 1 alert, will silence this rule set until all of the severity 1 alerts on the parent clear. This value must match the format ` + "`" + `${check_id}_${metric_name}` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "metric_name",
					Description: `(Required) The name of the metric stream within a given check that this rule set is active on.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A list of tags assigned to this rule set. ## ` + "`" + `if` + "`" + ` Configuration The ` + "`" + `if` + "`" + ` configuration block is an [ordered list of rules](https://login.circonus.com/user/docs/Alerting/Rules/Configure) that are evaluated in order, first to last. The first ` + "`" + `if` + "`" + ` condition to evaluate true shortcircuits all other ` + "`" + `if` + "`" + ` blocks in this rule set. An ` + "`" + `if` + "`" + ` block is also referred to as a "rule." It is advised that all high-severity rules are ordered before low-severity rules otherwise low-severity rules will mask notifications that should be delivered with a high-severity. ` + "`" + `if` + "`" + ` blocks are made up of two configuration blocks: ` + "`" + `value` + "`" + ` and ` + "`" + `then` + "`" + `. The ` + "`" + `value` + "`" + ` configuration block specifies the criteria underwhich the metric streams are evaluated. The ` + "`" + `then` + "`" + ` configuration block, optional, specifies what action to take. ### ` + "`" + `value` + "`" + ` Configuration A ` + "`" + `value` + "`" + ` block can have only one of several "predicate" attributes specified because they conflict with each other. The list of mutually exclusive predicates is dependent on the ` + "`" + `metric_type` + "`" + `. To evaluate multiple predicates, create multiple ` + "`" + `if` + "`" + ` configuration blocks in the proper order. #### ` + "`" + `numeric` + "`" + ` Predicates Metric types of type ` + "`" + `numeric` + "`" + ` support the following predicates. Only one of the following predicates may be specified at a time.`,
				},
				resource.Attribute{
					Name:        "absent",
					Description: `(Optional) If a metric has not been observed in this duration the rule will fire. When present, this duration is evaluated in terms of seconds.`,
				},
				resource.Attribute{
					Name:        "changed",
					Description: `(Optional) A boolean indicating this rule should fire when the value changes (e.g. ` + "`" + `n != n<sub>1</sub>` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "min_value",
					Description: `(Optional) When the value is less than this value, this rule will fire (e.g. ` + "`" + `n < ${min_value}` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "max_value",
					Description: `(Optional) When the value is greater than this value, this rule will fire (e.g. ` + "`" + `n > ${max_value}` + "`" + `). Additionally, a ` + "`" + `numeric` + "`" + ` check can also evaluate data based on a windowing function versus the last measured value in the metric stream. In order to have a rule evaluate on derived value from a window, include a nested ` + "`" + `over` + "`" + ` attribute inside of the ` + "`" + `value` + "`" + ` configuration block. An ` + "`" + `over` + "`" + ` attribute needs two attributes:`,
				},
				resource.Attribute{
					Name:        "last",
					Description: `(Optional) A duration for the sliding window. Default ` + "`" + `300s` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "using",
					Description: `(Optional) The window function to use over the ` + "`" + `last` + "`" + ` interval. Valid window functions include: ` + "`" + `average` + "`" + ` (the default), ` + "`" + `stddev` + "`" + `, ` + "`" + `derive` + "`" + `, ` + "`" + `derive_stddev` + "`" + `, ` + "`" + `counter` + "`" + `, ` + "`" + `counter_stddev` + "`" + `, ` + "`" + `derive_2` + "`" + `, ` + "`" + `derive_2_stddev` + "`" + `, ` + "`" + `counter_2` + "`" + `, and ` + "`" + `counter_2_stddev` + "`" + `. #### ` + "`" + `text` + "`" + ` Predicates Metric types of type ` + "`" + `text` + "`" + ` support the following predicates:`,
				},
				resource.Attribute{
					Name:        "absent",
					Description: `(Optional) If a metric has not been observed in this duration the rule will fire. When present, this duration is evaluated in terms of seconds.`,
				},
				resource.Attribute{
					Name:        "changed",
					Description: `(Optional) A boolean indicating this rule should fire when the last value in the metric stream changed from it's previous value (e.g. ` + "`" + `n != n-1` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "contains",
					Description: `(Optional) When the last value in the metric stream the value is less than this value, this rule will fire (e.g. ` + "`" + `strstr(n, ${contains}) != NULL` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "match",
					Description: `(Optional) When the last value in the metric stream value exactly matches this configured value, this rule will fire (e.g. ` + "`" + `strcmp(n, ${match}) == 0` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "not_contain",
					Description: `(Optional) When the last value in the metric stream does not match this configured value, this rule will fire (e.g. ` + "`" + `strstr(n, ${contains}) == NULL` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "not_match",
					Description: `(Optional) When the last value in the metric stream does not match this configured value, this rule will fire (e.g. ` + "`" + `strstr(n, ${not_match}) == NULL` + "`" + `). ### ` + "`" + `then` + "`" + ` Configuration A ` + "`" + `then` + "`" + ` block can have the following attributes:`,
				},
				resource.Attribute{
					Name:        "after",
					Description: `(Optional) Only execute this notification after waiting for this number of minutes. Defaults to immediately, or ` + "`" + `0m` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "notify",
					Description: `(Optional) A list of contact group IDs to notify when this rule is sends off a notification.`,
				},
				resource.Attribute{
					Name:        "severity",
					Description: `(Optional) The severity level of the notification. This can be set to any value between ` + "`" + `0` + "`" + ` and ` + "`" + `5` + "`" + `. Defaults to ` + "`" + `1` + "`" + `. ## Import Example ` + "`" + `circonus_rule_set` + "`" + ` supports importing resources. Supposing the following Terraform (and that the referenced [` + "`" + `circonus_metric` + "`" + `](metric.html) and [` + "`" + `circonus_check` + "`" + `](check.html) have already been imported): ` + "`" + `` + "`" + `` + "`" + `hcl resource "circonus_rule_set" "icmp-latency-alert" { check = "${circonus_check.api_latency.checks[0]}" metric_name = "maximum" if { value { absent = "600s" } then { notify = [ "${circonus_contact_group.test-trigger.id}" ] severity = 1 } } if { value { over { last = "120s" using = "average" } max_value = 0.5 # units are in miliseconds } then { notify = [ "${circonus_contact_group.test-trigger.id}" ] severity = 2 } } } ` + "`" + `` + "`" + `` + "`" + ` It is possible to import a ` + "`" + `circonus_rule_set` + "`" + ` resource with the following command: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import circonus_rule_set.icmp-latency-alert ID ` + "`" + `` + "`" + `` + "`" + ` Where ` + "`" + `ID` + "`" + ` is the ` + "`" + `_cid` + "`" + ` or Circonus ID of the Rule Set (e.g. ` + "`" + `/rule_set/201285_maximum` + "`" + `) and ` + "`" + `circonus_rule_set.icmp-latency-alert` + "`" + ` is the name of the resource whose state will be populated as a result of the command.`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "circonus_worksheet",
			Category:         "Resources",
			ShortDescription: `Manages a Circonus worksheet.`,
			Description:      ``,
			Keywords: []string{
				"worksheet",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "title",
					Description: `(Required) The title of the worksheet.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of what the worksheet is for.`,
				},
				resource.Attribute{
					Name:        "favourite",
					Description: `(Optional) Mark (star) this worksheet as a favorite. Default is ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `(Optional) A place to store notes about this worksheet.`,
				},
				resource.Attribute{
					Name:        "graphs",
					Description: `(Optional) A list of graphs that compose this worksheet.`,
				},
				resource.Attribute{
					Name:        "smart_queries",
					Description: `(Optional) The smart queries that will be displayed on this worksheet. See below for details on how to configure a ` + "`" + `smart_query` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A list of tags assigned to this worksheet. ### ` + "`" + `smart_queries` + "`" + ` Attributes ` + "`" + `smart_queries` + "`" + ` is a list of smart query objects. Each smart query object has the following required attributes:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name (heading) for the smart graph section in the worksheet.`,
				},
				resource.Attribute{
					Name:        "query",
					Description: `(Required) A search query that determines which graphs will be shown.. ## Import Example It is possible to import a ` + "`" + `circonus_worksheet` + "`" + ` resource with the following command: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import circonus_worksheet.icmp-latency ID ` + "`" + `` + "`" + `` + "`" + ` Where ` + "`" + `ID` + "`" + ` is the ` + "`" + `_cid` + "`" + ` or Circonus ID of the worksheet (e.g. ` + "`" + `worksheets/45640239-bb81-4ecb-81e6-b5c6015e5dd5` + "`" + `) and ` + "`" + `circonus_worksheet.icmp-latency` + "`" + ` is the name of the resource whose state will be populated as a result of the command.`,
				},
			},
			Attributes: []resource.Argument{},
		},
	}

	resourcesMap = map[string]int{

		"circonus_check":          0,
		"circonus_contact_group":  1,
		"circonus_graph":          2,
		"circonus_metric":         3,
		"circonus_metric_cluster": 4,
		"circonus_rule_set":       5,
		"circonus_worksheet":      6,
	}
)

func GetResource(r string) (*resouce.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs]
}
