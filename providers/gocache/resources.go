package gocache

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "gocache_gocache_domain",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "domain",
					Description: `(Required) The domain you want to add.`,
				},
				resource.Attribute{
					Name:        "cache_ttl",
					Description: `(Optional) The time after the CDN cache will be expired. Possible values are ` + "`" + `300` + "`" + `, ` + "`" + `600` + "`" + `, ` + "`" + `900` + "`" + `, ` + "`" + `1800` + "`" + `, ` + "`" + `3600` + "`" + `, ` + "`" + `7200` + "`" + `, ` + "`" + `14400` + "`" + `, ` + "`" + `86400` + "`" + `, ` + "`" + `172800` + "`" + `, ` + "`" + `604800` + "`" + `. Default ` + "`" + `86400` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "deploy_mode",
					Description: `(Optional) Deactivates cache for debbugging purpose. Default ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "smart_cache_status",
					Description: `(Optional) Activates Smart Cache, a template of cache rules for dinamic cache on popular Content Management Systems. Default ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "smart_cache_template",
					Description: `(Optional) Selects the Smart Cache template. Possible values are ` + "`" + `custom` + "`" + `, ` + "`" + `wordpress` + "`" + `, ` + "`" + `magento` + "`" + `, ` + "`" + `joomla` + "`" + `. Default ` + "`" + `custom` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "smart_cache_ttl",
					Description: `(Optional) The time after the cache from Smart Cache will be expired. Possible values are ` + "`" + `300` + "`" + `, ` + "`" + `600` + "`" + `, ` + "`" + `900` + "`" + `, ` + "`" + `1800` + "`" + `, ` + "`" + `3600` + "`" + `, ` + "`" + `7200` + "`" + `, ` + "`" + `14400` + "`" + `, ` + "`" + `86400` + "`" + `, ` + "`" + `172800` + "`" + `, ` + "`" + `604800` + "`" + `. Default ` + "`" + `14400` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "dns_mode",
					Description: `(Optional) Choose ` + "`" + `ns` + "`" + ` to use the GoCache name servers or choose ` + "`" + `cname` + "`" + ` if you are goint to maintain your current name servers. Default ` + "`" + `ns` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "gzip_status",
					Description: `(Optional) Activates the GZIP compression on assets server by the CDN. Default ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "expires_ttl",
					Description: `(Optional) The define the ` + "`" + `expires` + "`" + ` header that will be sent with assets requests to determines the time they will be cached in browsers. Possible values are ` + "`" + `-1` + "`" + `, ` + "`" + `10` + "`" + `, ` + "`" + `30` + "`" + `, ` + "`" + `60` + "`" + `, ` + "`" + `300` + "`" + `, ` + "`" + `600` + "`" + `, ` + "`" + `900` + "`" + `, ` + "`" + `1800` + "`" + `, ` + "`" + `3600` + "`" + `, ` + "`" + `7200` + "`" + `, ` + "`" + `14400` + "`" + `, ` + "`" + `43200` + "`" + `, ` + "`" + `86400` + "`" + `, ` + "`" + `172800` + "`" + `, ` + "`" + `345600` + "`" + `, ` + "`" + `604800` + "`" + `, ` + "`" + `1296000` + "`" + `, ` + "`" + `2592000` + "`" + `, ` + "`" + `15552000` + "`" + `, ` + "`" + `31536000` + "`" + `. Default ` + "`" + `14400` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ignore_vary",
					Description: `(Optional) Ignores the ` + "`" + `vary` + "`" + ` header sent by the origin server. Default ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ignore_cache_control",
					Description: `(Optional) Ignores the ` + "`" + `cache-control` + "`" + ` header sent by the origin server. Default ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ignore_expires",
					Description: `(Optional) Ignores the ` + "`" + `expires` + "`" + ` header sent by the origin server. Default ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ssl_mode",
					Description: `(Optional) The ` + "`" + `full` + "`" + ` value indicates that when a client sends an https request, GoCache will comunite with origin server on port 443 and the ` + "`" + `partial` + "`" + ` value indicates all requests sent to origin server by GoCache will be on port 80. Default ` + "`" + `full` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "cache_301",
					Description: `(Optional) Defines if responses with status 301 will be cache. Default ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "cache_302",
					Description: `(Optional) Defines if responses with status 302 will be cache. Default ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "cache_404",
					Description: `(Optional) Defines if responses with status 404 will be cache. Default ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "header_device_type",
					Description: `(Optional) If activated, a header ` + "`" + `GoCache-Device-Type` + "`" + ` will be sent to origin server containing the device type that made the request. Default ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "header_geoip_continent",
					Description: `(Optional) If activated, a header ` + "`" + `GoCache-GeoIP-Continent` + "`" + ` will be sent to origin server containing the continent where the request was from in ISO-ALPHA-2 notation. Default ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "header_geoip_country",
					Description: `(Optional) If activated, a header ` + "`" + `GoCache-GeoIP-Country` + "`" + ` will be sent to origin server containing the country where the request was from in ISO-ALPHA-2 notation. Default ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "header_geoip_org",
					Description: `(Optional) If activated, a header ` + "`" + `GoCache-GeoIP-Org` + "`" + ` will be sent to origin server containing the Autonomous System Number (ASN) and organization responsible for the IP that made the request. Default ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "caching_behavior",
					Description: `(Optional) When ` + "`" + `default` + "`" + `, each querystring on an URL will generate another cache, when ` + "`" + `ignore` + "`" + ` all querystrings in the same URL will be served by the same cache. Default ` + "`" + `default` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "waf_status",
					Description: `(Optional) Enables the WAF when status in ` + "`" + `true` + "`" + `. Default ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "waf_level",
					Description: `(Optional) Defines the sensitivity level of WAF. Possible values are ` + "`" + `low` + "`" + `, ` + "`" + `medium` + "`" + `, ` + "`" + `high` + "`" + `. Default ` + "`" + `low` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "waf_mode",
					Description: `(Optional) Defines the default action on a request when it is detected by WAF. When ` + "`" + `simulate` + "`" + `, the request is allowed but logged in security events. When ` + "`" + `challenge` + "`" + `, a captcha challenge will be sent to the client. Finally, when 'block', the request is blocked with status 403. Default ` + "`" + `simulate` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "expires_bypass_sec",
					Description: `(Optional) Sets the expiration time of the cookie sent for clients that was approved on challenge. Possible values are ` + "`" + `-1` + "`" + `, ` + "`" + `300` + "`" + `, ` + "`" + `600` + "`" + `, ` + "`" + `900` + "`" + `, ` + "`" + `1800` + "`" + `, ` + "`" + `3600` + "`" + `, ` + "`" + `7200` + "`" + `, ` + "`" + `14400` + "`" + `, ` + "`" + `43200` + "`" + `, ` + "`" + `86400` + "`" + `, ` + "`" + `172800` + "`" + `, ` + "`" + `345600` + "`" + `, ` + "`" + `604800` + "`" + `, ` + "`" + `1296000` + "`" + `, ` + "`" + `2592000` + "`" + `, ` + "`" + `15552000` + "`" + `, ` + "`" + `31536000` + "`" + `. Default ` + "`" + `-1` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "tls10",
					Description: `(Optional) Enables the TLS 1.0 protocol. Default ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "tls11",
					Description: `(Optional) Enables the TLS 1.1 protocol. Default ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "image_optimize",
					Description: `(Optional) Enables the image optimizer. This option needs to be enabled for all image optization options bellow work. See the https://www.gocache.com.br/planos/ page to know the princing of this feature. Default ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "image_optimize_webp",
					Description: `(Optional) Enables the image conversion to WEBP for browser that support the format. Default ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "image_optimize_progressive",
					Description: `(Optional) Enables the image conversion to progressive JPEG for browsers that support the format. Convertion to WEBP takes precedence. Default ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "image_optimize_metadata",
					Description: `(Optional) Removes unnecessary metadata from the images. Default ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "image_optimize_level",
					Description: `(Optional) Sets the compression level on the image optimization proccess. Choose ` + "`" + `0` + "`" + ` to maintain the image quality, ` + "`" + `90` + "`" + ` for low compression, ` + "`" + `75` + "`" + ` for medium and ` + "`" + `65` + "`" + ` for high. Default ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "rate_limit",
					Description: `(Optional) Enables the rate limiting. See the https://www.gocache.com.br/planos/ page to know the princing of this feature. Default ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "rate_limit_ignore_good_bots",
					Description: `(Optional) Ignores good bots in rate limiting. Default ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "rate_limit_ignore_static_content",
					Description: `(Optional) Ignores static content in rate limiting. Default ` + "`" + `true` + "`" + `. ## Import Domain resource can be imported using the domain address. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import gocache_domain.example example.com.br ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "gocache_gocache_domain_dnssec",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "domain",
					Description: `(Required) The domain to enable DNSSEC. ## Attribute Reference The following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "digest",
					Description: `Digest used for DNSSEC.`,
				},
				resource.Attribute{
					Name:        "digest_type",
					Description: `Digest type id used for DNSSEC.`,
				},
				resource.Attribute{
					Name:        "digest_type_name",
					Description: `Hash algorithm name represented by ` + "`" + `digest_type` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "algorithm",
					Description: `DNSSEC algorithm used.`,
				},
				resource.Attribute{
					Name:        "public_key",
					Description: `Public Key used for DNSSEC.`,
				},
				resource.Attribute{
					Name:        "key_tag",
					Description: `Key tag used for DNSSEC.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "digest",
					Description: `Digest used for DNSSEC.`,
				},
				resource.Attribute{
					Name:        "digest_type",
					Description: `Digest type id used for DNSSEC.`,
				},
				resource.Attribute{
					Name:        "digest_type_name",
					Description: `Hash algorithm name represented by ` + "`" + `digest_type` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "algorithm",
					Description: `DNSSEC algorithm used.`,
				},
				resource.Attribute{
					Name:        "public_key",
					Description: `Public Key used for DNSSEC.`,
				},
				resource.Attribute{
					Name:        "key_tag",
					Description: `Key tag used for DNSSEC.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "gocache_gocache_record",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "domain",
					Description: `(Required) The domain to add the record to.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The subdomain that the record will set. To declare an ` + "`" + `A` + "`" + ` record for the root domain on ` + "`" + `ns` + "`" + ` mode, set ` + "`" + `@` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "content",
					Description: `(Required) The value the record returns. When ` + "`" + `dns_mode` + "`" + ` is ` + "`" + `ns` + "`" + ` and record ` + "`" + `type` + "`" + ` is not ` + "`" + `TXT` + "`" + ` or ` + "`" + `dns_mode` + "`" + ` is ` + "`" + `cname` + "`" + `, it indicates a server address where the service of that subdomain is hosted. ` + "`" + `A` + "`" + ` type records only accepts IPs while ` + "`" + `CNAME` + "`" + ` and ` + "`" + `SRV` + "`" + ` type records only accepts hostnames. ` + "`" + `TXT` + "`" + ` type records is used to returns generic text. ## Specific arguments for ` + "`" + `dns` + "`" + ` mode This section lists specific arguments when ` + "`" + `dns_mode` + "`" + ` is set to ` + "`" + `ns` + "`" + ``,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) The DNS record type. Possible values are ` + "`" + `A` + "`" + `, ` + "`" + `CNAME` + "`" + `, ` + "`" + `NS` + "`" + `, ` + "`" + `TXT` + "`" + `, ` + "`" + `MX` + "`" + `, ` + "`" + `SRV` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `(Required) The DNS record type. Possible values are ` + "`" + `60` + "`" + `, ` + "`" + `300` + "`" + `, ` + "`" + `600` + "`" + `, ` + "`" + `900` + "`" + `, ` + "`" + `1800` + "`" + `, ` + "`" + `3600` + "`" + `, ` + "`" + `7200` + "`" + `, ` + "`" + `14400` + "`" + `, ` + "`" + `43200` + "`" + `, ` + "`" + `86400` + "`" + `. ### Arguments required by ` + "`" + `A` + "`" + ` and ` + "`" + `CNAME` + "`" + ` records The following argument is accepted and required only by ` + "`" + `A` + "`" + ` and ` + "`" + `CNAME` + "`" + ` records`,
				},
				resource.Attribute{
					Name:        "proxied",
					Description: `(Required) When set to ` + "`" + `true` + "`" + `, the subdomain traffic is proxied through GoCache, enabling the most of the services. When set to ` + "`" + `false` + "`" + `, the subdomain traffic is not delivered by GoCache. ### Optional arguments accepted by ` + "`" + `MX` + "`" + ` and ` + "`" + `SRV` + "`" + ` records`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `(Optional) When two or more records are declared for the same subdomain, the records with the lowest value is prioritized ### Arguments required by ` + "`" + `SRV` + "`" + ` records`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Required) The port number the service configured by this record listen`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `(Required) When two or more records are declared for the same subdomain, with the same ` + "`" + `priority` + "`" + ` argument value or without it, the traffic is distributed proportionaly based on the value of this field ## Attribute Reference The following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "record_id",
					Description: `The record ID in GoCache ## Import Record resource can be imported using a composite ID formed of domain and its record ID, as returned from [API](https://docs.gocache.com.br/api/#api-Websites_e_DNS-GetDNS) ` + "`" + `` + "`" + `` + "`" + ` $ terraform import gocache_record.www example.com.br/99990 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "record_id",
					Description: `The record ID in GoCache ## Import Record resource can be imported using a composite ID formed of domain and its record ID, as returned from [API](https://docs.gocache.com.br/api/#api-Websites_e_DNS-GetDNS) ` + "`" + `` + "`" + `` + "`" + ` $ terraform import gocache_record.www example.com.br/99990 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "gocache_gocache_smart_rules_firewall",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "domain",
					Description: `(Required) The domain where you will apply the policies.`,
				},
				resource.Attribute{
					Name:        "smart_rule",
					Description: `(Required) Each ` + "`" + `smart_rule` + "`" + ` block sets a rule. More than one block is allowed. List from top to bottom in order of precedence. The topmost block will have precedence 0. See below the structure of each block. ### Smart Rule Level Arguments`,
				},
				resource.Attribute{
					Name:        "match",
					Description: `(Required) Request criteria needed to trigger the rule action. At least one option is needed. See the options in [match level arguments](match-level-arguments).`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Required) Actions applied to a request when the criteria is matched. At least one option is needed. See the options in [action level arguments](action-level-arguments).`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `(Optional) Rule metadata. See the options in [metadata level arguments](metadata-level-arguments). ### Match Level Arguments Some fields can use regex patterns. To know more about regex, see the documentation in https://docs.gocache.com.br/smart_rules/smart_rules-regexp/.`,
				},
				resource.Attribute{
					Name:        "request_method",
					Description: `(Optional) List of HTTP request methods matched by the rule. Possible values ` + "`" + `GET` + "`" + `, ` + "`" + `POST` + "`" + `, ` + "`" + `PUT` + "`" + `, ` + "`" + `DELETE` + "`" + `, ` + "`" + `HEAD` + "`" + `, ` + "`" + `OPTIONS` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "scheme",
					Description: `(Optional) List of HTTP request scheme matched by the rule. Possible values ` + "`" + `http` + "`" + `, ` + "`" + `https` + "`" + ` or ` + "`" + `http`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `(Optional) HTTP request hostname matched by the rule. For all hostnames, set ` + "`" + ``,
				},
				resource.Attribute{
					Name:        "request_uri",
					Description: `(Optional) HTTP request URI pattern matched by the rule. You can use regex in this field.`,
				},
				resource.Attribute{
					Name:        "http_user_agent",
					Description: `(Optional) Request user agent matched by the rule. You can use regex in this field.`,
				},
				resource.Attribute{
					Name:        "cookie",
					Description: `(Optional) List of cequest cookie names matched by the rule. This field supports wildcard.`,
				},
				resource.Attribute{
					Name:        "cookie_content",
					Description: `(Optional) Cookie values for a cookie matched by the rule. You need to declare cookie and its value separeted by ` + "`" + `=` + "`" + `, like ` + "`" + `wp_test=yes` + "`" + `. You can use regex in the value.`,
				},
				resource.Attribute{
					Name:        "http_referer",
					Description: `(Optional) HTTP request referer matched by the rule.You can use regex in this field.`,
				},
				resource.Attribute{
					Name:        "remote_address",
					Description: `(Optional) Request client IP ou IP ranges matched by the rule. Only ` + "`" + `/32` + "`" + `, ` + "`" + `/24` + "`" + ` and ` + "`" + `/16` + "`" + ` ranges are allowed.`,
				},
				resource.Attribute{
					Name:        "http_version",
					Description: `(Optional) List of request HTTP version matched by the rule. Possible values are ` + "`" + `HTTP/1.0` + "`" + `,` + "`" + `HTTP/1.1` + "`" + `,` + "`" + `HTTP/2.0` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "header",
					Description: `(Optional) HTTP headers matched by the rule. Header name and value are separated by ` + "`" + `:` + "`" + `. You can use regex in the value.`,
				},
				resource.Attribute{
					Name:        "origin_country",
					Description: `(Optional) List of origin countries matched by the rule in the ISO format, like ` + "`" + `BR` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "origin_continent",
					Description: `(Optional) List of origin continents matched by the rule in the ISO format. Possible values ` + "`" + `SA` + "`" + `, ` + "`" + `NA` + "`" + `, ` + "`" + `OC` + "`" + `, ` + "`" + `EU` + "`" + `, ` + "`" + `AF` + "`" + `, ` + "`" + `AS` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "device_type",
					Description: `(Optional) list of device types based on user agent matched by the rule. Possible values ` + "`" + `mobile` + "`" + `,` + "`" + `desktop` + "`" + `,` + "`" + `bot` + "`" + `,` + "`" + `na` + "`" + ``,
				},
				resource.Attribute{
					Name:        "bots",
					Description: `(Optional) If the bot is legitimate. Possible values ` + "`" + `known` + "`" + `, ` + "`" + `others` + "`" + ` ### Action Level Arguments`,
				},
				resource.Attribute{
					Name:        "firewall_action",
					Description: `(Required) When ` + "`" + `block` + "`" + `, the matched request are denied with status ` + "`" + `403` + "`" + `, when ` + "`" + `challenge` + "`" + `, a captcha challenge is returned and when ` + "`" + `accept` + "`" + `, the request is allowed, even when WAF blocks it. ### Metadata Level Arguments`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Optional) If ` + "`" + `true` + "`" + `, the rule is activated.`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `(Optional) Smart Rule description ## Import Domain resource can be imported using the domain address. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import gocache_smart_rules_firewall.mypolicies example.com.br ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "gocache_gocache_smart_rules_rewrite",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "domain",
					Description: `(Required) The domain where you will apply the policies.`,
				},
				resource.Attribute{
					Name:        "smart_rule",
					Description: `(Required) Each ` + "`" + `smart_rule` + "`" + ` block sets a rule. More than one block is allowed. List from top to bottom in order of precedence. The topmost block will have precedence 0. See below the structure of each block. ### Smart Rule Level Arguments`,
				},
				resource.Attribute{
					Name:        "match",
					Description: `(Required) Request criteria needed to trigger the rule action. At least one option is needed. See the options in [match level arguments](match-level-arguments).`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Required) Actions applied to a request when the criteria is matched. At least one option is needed. See the options in [action level arguments](action-level-arguments).`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `(Optional) Rule metadata. See the options in [metadata level arguments](metadata-level-arguments). ### Match Level Arguments Some fields can use regex patterns. To know more about regex, see the documentation in https://docs.gocache.com.br/smart_rules/smart_rules-regexp/.`,
				},
				resource.Attribute{
					Name:        "request_method",
					Description: `(Optional) List of HTTP request methods matched by the rule. Possible values ` + "`" + `GET` + "`" + `, ` + "`" + `POST` + "`" + `, ` + "`" + `PUT` + "`" + `, ` + "`" + `DELETE` + "`" + `, ` + "`" + `HEAD` + "`" + `, ` + "`" + `OPTIONS` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "scheme",
					Description: `(Optional) List of HTTP request scheme matched by the rule. Possible values ` + "`" + `http` + "`" + `, ` + "`" + `https` + "`" + ` or ` + "`" + `http`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `(Optional) HTTP request hostname matched by the rule. For all hostnames, set ` + "`" + ``,
				},
				resource.Attribute{
					Name:        "request_uri",
					Description: `(Optional) HTTP request URI pattern matched by the rule. You can use regex in this field.`,
				},
				resource.Attribute{
					Name:        "http_user_agent",
					Description: `(Optional) Request user agent matched by the rule. You can use regex in this field.`,
				},
				resource.Attribute{
					Name:        "cookie",
					Description: `(Optional) List of cequest cookie names matched by the rule. This field supports wildcard.`,
				},
				resource.Attribute{
					Name:        "cookie_content",
					Description: `(Optional) Cookie values for a cookie matched by the rule. You need to declare cookie and its value separeted by ` + "`" + `=` + "`" + `, like ` + "`" + `wp_test=yes` + "`" + `. You can use regex in the value.`,
				},
				resource.Attribute{
					Name:        "http_referer",
					Description: `(Optional) HTTP request referer matched by the rule.You can use regex in this field.`,
				},
				resource.Attribute{
					Name:        "remote_address",
					Description: `(Optional) Request client IP ou IP ranges matched by the rule. Only ` + "`" + `/32` + "`" + `, ` + "`" + `/24` + "`" + ` and ` + "`" + `/16` + "`" + ` ranges are allowed.`,
				},
				resource.Attribute{
					Name:        "http_version",
					Description: `(Optional) List of request HTTP version matched by the rule. Possible values are ` + "`" + `HTTP/1.0` + "`" + `,` + "`" + `HTTP/1.1` + "`" + `,` + "`" + `HTTP/2.0` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "header",
					Description: `(Optional) HTTP headers matched by the rule. Header name and value are separated by ` + "`" + `:` + "`" + `. You can use regex in the value.`,
				},
				resource.Attribute{
					Name:        "origin_country",
					Description: `(Optional) List of origin countries matched by the rule in the ISO format, like ` + "`" + `BR` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "origin_continent",
					Description: `(Optional) List of origin continents matched by the rule in the ISO format. Possible values ` + "`" + `SA` + "`" + `, ` + "`" + `NA` + "`" + `, ` + "`" + `OC` + "`" + `, ` + "`" + `EU` + "`" + `, ` + "`" + `AF` + "`" + `, ` + "`" + `AS` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "device_type",
					Description: `(Optional) list of device types based on user agent matched by the rule. Possible values ` + "`" + `mobile` + "`" + `,` + "`" + `desktop` + "`" + `,` + "`" + `bot` + "`" + `,` + "`" + `na` + "`" + ``,
				},
				resource.Attribute{
					Name:        "bots",
					Description: `(Optional) If the bot is legitimate. Possible values ` + "`" + `known` + "`" + `, ` + "`" + `others` + "`" + ` ### Action Level Arguments`,
				},
				resource.Attribute{
					Name:        "redirect_type",
					Description: `(Required) Status code returned on redirect. Possible values ` + "`" + `301` + "`" + `,` + "`" + `302` + "`" + ``,
				},
				resource.Attribute{
					Name:        "redirect_to",
					Description: `(Required) URL where the request needs to be redirected. ### Metadata Level Arguments`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Optional) If ` + "`" + `true` + "`" + `, the rule is activated.`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `(Optional) Smart Rule description ## Import Domain resource can be imported using the domain address. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import gocache_smart_rules_rewrite.mypolicies example.com.br ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "gocache_gocache_smart_rules_settings",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "domain",
					Description: `(Required) The domain where you will apply the policies.`,
				},
				resource.Attribute{
					Name:        "smart_rule",
					Description: `(Required) Each ` + "`" + `smart_rule` + "`" + ` block sets a rule. More than one block is allowed. List from top to bottom in order of precedence. The topmost block will have precedence 0. See below the structure of each block. ### Smart Rule Level Arguments`,
				},
				resource.Attribute{
					Name:        "match",
					Description: `(Required) Request criteria needed to trigger the rule action. At least one option is needed. See the options in [match level arguments](match-level-arguments).`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Required) Actions applied to a request when the criteria is matched. At least one option is needed. See the options in [action level arguments](action-level-arguments).`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `(Optional) Rule metadata. See the options in [metadata level arguments](metadata-level-arguments). ### Match Level Arguments Some fields can use regex patterns. To know more about regex, see the documentation in https://docs.gocache.com.br/smart_rules/smart_rules-regexp/.`,
				},
				resource.Attribute{
					Name:        "request_method",
					Description: `(Optional) List of HTTP request methods matched by the rule. Possible values ` + "`" + `GET` + "`" + `, ` + "`" + `POST` + "`" + `, ` + "`" + `PUT` + "`" + `, ` + "`" + `DELETE` + "`" + `, ` + "`" + `HEAD` + "`" + `, ` + "`" + `OPTIONS` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "scheme",
					Description: `(Optional) List of HTTP request scheme matched by the rule. Possible values ` + "`" + `http` + "`" + `, ` + "`" + `https` + "`" + ` or ` + "`" + `http`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `(Optional) HTTP request hostname matched by the rule. For all hostnames, set ` + "`" + ``,
				},
				resource.Attribute{
					Name:        "request_uri",
					Description: `(Optional) HTTP request URI pattern matched by the rule. You can use regex in this field.`,
				},
				resource.Attribute{
					Name:        "http_user_agent",
					Description: `(Optional) Request user agent matched by the rule. You can use regex in this field.`,
				},
				resource.Attribute{
					Name:        "cookie",
					Description: `(Optional) List of cequest cookie names matched by the rule. This field supports wildcard.`,
				},
				resource.Attribute{
					Name:        "cookie_content",
					Description: `(Optional) Cookie values for a cookie matched by the rule. You need to declare cookie and its value separeted by ` + "`" + `=` + "`" + `, like ` + "`" + `wp_test=yes` + "`" + `. You can use regex in the value.`,
				},
				resource.Attribute{
					Name:        "http_referer",
					Description: `(Optional) HTTP request referer matched by the rule.You can use regex in this field.`,
				},
				resource.Attribute{
					Name:        "remote_address",
					Description: `(Optional) Request client IP ou IP ranges matched by the rule. Only ` + "`" + `/32` + "`" + `, ` + "`" + `/24` + "`" + ` and ` + "`" + `/16` + "`" + ` ranges are allowed.`,
				},
				resource.Attribute{
					Name:        "http_version",
					Description: `(Optional) List of request HTTP version matched by the rule. Possible values are ` + "`" + `HTTP/1.0` + "`" + `,` + "`" + `HTTP/1.1` + "`" + `,` + "`" + `HTTP/2.0` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "header",
					Description: `(Optional) HTTP headers matched by the rule. Header name and value are separated by ` + "`" + `:` + "`" + `. You can use regex in the value.`,
				},
				resource.Attribute{
					Name:        "origin_country",
					Description: `(Optional) List of origin countries matched by the rule in the ISO format, like ` + "`" + `BR` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "origin_continent",
					Description: `(Optional) List of origin continents matched by the rule in the ISO format. Possible values ` + "`" + `SA` + "`" + `, ` + "`" + `NA` + "`" + `, ` + "`" + `OC` + "`" + `, ` + "`" + `EU` + "`" + `, ` + "`" + `AF` + "`" + `, ` + "`" + `AS` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "device_type",
					Description: `(Optional) list of device types based on user agent matched by the rule. Possible values ` + "`" + `mobile` + "`" + `,` + "`" + `desktop` + "`" + `,` + "`" + `bot` + "`" + `,` + "`" + `na` + "`" + ``,
				},
				resource.Attribute{
					Name:        "bots",
					Description: `(Optional) If the bot is legitimate. Possible values ` + "`" + `known` + "`" + `, ` + "`" + `others` + "`" + ` ### Action Level Arguments`,
				},
				resource.Attribute{
					Name:        "set_host",
					Description: `(Optional) Overrides the header ` + "`" + `Host` + "`" + ` when requesting the origin.`,
				},
				resource.Attribute{
					Name:        "backend",
					Description: `(Optional) Sets a different backend as origin for the request. It can be an IP address or a hostname.`,
				},
				resource.Attribute{
					Name:        "set_uri",
					Description: `(Optional) Overrides the URI when requesting the origin. If there is one or more wildcard in the ` + "`" + `request_uri` + "`" + ` match criteria, you can use variables like ` + "`" + `$1` + "`" + ` to pass the respective string sequence on the backend request.`,
				},
				resource.Attribute{
					Name:        "custom_cache_key",
					Description: `(Optional) Customizes cache key. You can use variables like ` + "`" + `$cookie_NAME` + "`" + ` to group cached requests by specific cookie values, ` + "`" + `$http_NAME` + "`" + `, to group by specific header values, ` + "`" + `$1` + "`" + ` to group by URI patterns respective to ` + "`" + `request_uri` + "`" + ` wildcards, ` + "`" + `$geoip2_data_country_code` + "`" + ` to group by request's country and ` + "`" + `$geoip2_data_continent_code` + "`" + ` to group by request's continent.`,
				},
				resource.Attribute{
					Name:        "expires_ttl",
					Description: `(Optional) Overrides the ` + "`" + `expires` + "`" + ` header that will be sent with assets requests. Possible values are ` + "`" + `-1` + "`" + `, ` + "`" + `10` + "`" + `, ` + "`" + `30` + "`" + `, ` + "`" + `60` + "`" + `, ` + "`" + `300` + "`" + `, ` + "`" + `600` + "`" + `, ` + "`" + `900` + "`" + `, ` + "`" + `1800` + "`" + `, ` + "`" + `3600` + "`" + `, ` + "`" + `7200` + "`" + `, ` + "`" + `14400` + "`" + `, ` + "`" + `43200` + "`" + `, ` + "`" + `86400` + "`" + `, ` + "`" + `172800` + "`" + `, ` + "`" + `345600` + "`" + `, ` + "`" + `604800` + "`" + `, ` + "`" + `1296000` + "`" + `, ` + "`" + `2592000` + "`" + `, ` + "`" + `15552000` + "`" + `, ` + "`" + `31536000` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "caching_behavior",
					Description: `(Optional) Overrides the caching behaviour when there are different querystrings on a matched URL. Possible values are ` + "`" + `default` + "`" + ` and ` + "`" + `ignore_query_string` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "cache_mode",
					Description: `(Optional) When ` + "`" + `off` + "`" + ` no cache are going to be applied, when ` + "`" + `default` + "`" + `, only static assets will be cached, when ` + "`" + `full` + "`" + `, all requests will be cached.`,
				},
				resource.Attribute{
					Name:        "cache_ttl",
					Description: `(Optional) Overrides the time after the CDN cache will be expired. Possible values are ` + "`" + `300` + "`" + `, ` + "`" + `600` + "`" + `, ` + "`" + `900` + "`" + `, ` + "`" + `1800` + "`" + `, ` + "`" + `3600` + "`" + `, ` + "`" + `7200` + "`" + `, ` + "`" + `14400` + "`" + `, ` + "`" + `86400` + "`" + `, ` + "`" + `172800` + "`" + `, ` + "`" + `604800` + "`" + `. Default ` + "`" + `86400` + "`" + `..`,
				},
				resource.Attribute{
					Name:        "cors",
					Description: `(Optional) Defines Cross-Origin headers for the matched requests, based on the address informed.`,
				},
				resource.Attribute{
					Name:        "ssl_mode",
					Description: `(Optional) The ` + "`" + `full` + "`" + ` value indicates that when a client sends an https request, GoCache will comunite with origin server on port 443 and the ` + "`" + `partial` + "`" + ` value indicates all requests sent to origin server by GoCache will be on port 80..`,
				},
				resource.Attribute{
					Name:        "hide_header",
					Description: `(Optional) Omits a response header informed.`,
				},
				resource.Attribute{
					Name:        "encoding_header",
					Description: `(Optional) Defines the Accept-Encoding header.`,
				},
				resource.Attribute{
					Name:        "set_request_headers",
					Description: `(Optional) Map object of request headers to be added on the upstream request, where header names are the keys. You can use variables like ` + "`" + `$cookie_NAME` + "`" + ` to group cached requests by specific cookie values, ` + "`" + `$http_NAME` + "`" + `, to group by specific header values, ` + "`" + `$geoip2_data_country_code` + "`" + ` to group by request's country and ` + "`" + `$geoip2_data_continent_code` + "`" + ` to group by request's continent.`,
				},
				resource.Attribute{
					Name:        "set_response_headers",
					Description: `(Optional) Map object of response headers to be added on the response, where header names are the keys.`,
				},
				resource.Attribute{
					Name:        "signed_url_key",
					Description: `(Optional) Defines a signed URL key to request S3 protected assets.`,
				},
				resource.Attribute{
					Name:        "signed_url_type",
					Description: `(Optional) Defines a signed URL type to request S3 protected assets. Possible values ` + "`" + `s3qs` + "`" + `, ` + "`" + `off` + "`" + ``,
				},
				resource.Attribute{
					Name:        "cache_301",
					Description: `(Optional) Defines if responses with status 301 will be cache.`,
				},
				resource.Attribute{
					Name:        "cache_302",
					Description: `(Optional) Defines if responses with status 302 will be cache..`,
				},
				resource.Attribute{
					Name:        "cache_404",
					Description: `(Optional) Defines if responses with status 404 will be cache..`,
				},
				resource.Attribute{
					Name:        "gzip_status",
					Description: `(Optional) Activates the GZIP compression on assets server by the CDN.`,
				},
				resource.Attribute{
					Name:        "ignore_cache_control",
					Description: `(Optional) Ignores the ` + "`" + `cache-control` + "`" + ` header sent by the origin server.`,
				},
				resource.Attribute{
					Name:        "ignore_expires",
					Description: `(Optional) Ignores the ` + "`" + `expires` + "`" + ` header sent by the origin server.`,
				},
				resource.Attribute{
					Name:        "ignore_vary",
					Description: `(Optional) Ignores the ` + "`" + `vary` + "`" + ` header sent by the origin server.`,
				},
				resource.Attribute{
					Name:        "waf_status",
					Description: `(Optional) Enables the WAF when status in ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "waf_mode",
					Description: `(Optional) Defines the default action on a request when it is detected by WAF. When ` + "`" + `simulate` + "`" + `, the request is allowed but logged in security events. When ` + "`" + `challenge` + "`" + `, a captcha challenge will be sent to the client. Finally, when 'block', the request is blocked with status 403.`,
				},
				resource.Attribute{
					Name:        "waf_level",
					Description: `(Optional) Defines the sensitivity level of WAF. Possible values are ` + "`" + `low` + "`" + `, ` + "`" + `medium` + "`" + `, ` + "`" + `high` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ratelimit_status",
					Description: `(Optional) Ignores the request in rate limiting.`,
				},
				resource.Attribute{
					Name:        "image_optimize",
					Description: `(Optional) Enables the image optimizer. This option needs to be enabled for all image optization options bellow work. See the https://www.gocache.com.br/planos/ page to know the princing of this feature.`,
				},
				resource.Attribute{
					Name:        "image_optimize_webp",
					Description: `(Optional) Enables the image conversion to WEBP for browser that support the format.`,
				},
				resource.Attribute{
					Name:        "image_optimize_progressive",
					Description: `(Optional) Enables the image conversion to progressive JPEG for browsers that support the format. Convertion to WEBP takes precedence.`,
				},
				resource.Attribute{
					Name:        "image_optimize_metadata",
					Description: `(Optional) Removes unnecessary metadata from the images.`,
				},
				resource.Attribute{
					Name:        "image_optimize_level",
					Description: `(Optional) Sets the compression level on the image optimization proccess. Choose ` + "`" + `0` + "`" + ` to maintain the image quality, ` + "`" + `90` + "`" + ` for low compression, ` + "`" + `75` + "`" + ` for medium and ` + "`" + `65` + "`" + ` for high.`,
				},
				resource.Attribute{
					Name:        "waf_rule_action",
					Description: `(Optional) Customizes on or more WAF rules behaviour. Possible behaviours are ` + "`" + `DISABLE` + "`" + `, ` + "`" + `DEFAULT` + "`" + `, ` + "`" + `BLOCK` + "`" + `, ` + "`" + `CHALLENGE` + "`" + `, ` + "`" + `SIMULATE` + "`" + `. Send the behaviour along with all rules you want this behaviour, separated by ` + "`" + `:` + "`" + `, like ` + "`" + `CHALLENGE:gocache-v1/90015,gocache-v1/41`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Optional) If ` + "`" + `true` + "`" + `, the rule is activated.`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `(Optional) Smart Rule description ## Import Domain resource can be imported using the domain address. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import gocache_smart_rules_settings.mypolicies example.com.br ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "gocache_gocache_ssl_certificate",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "domain",
					Description: `(Required) The domain to add the certificate to.`,
				},
				resource.Attribute{
					Name:        "private_key",
					Description: `(Required) The write-only private key in PEM format. Note: This property is sensitive and will not be displayed in the plan.`,
				},
				resource.Attribute{
					Name:        "certificate",
					Description: `(Required) The certificate in PEM format. Note: This property is sensitive and will not be displayed in the plan.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) Choose ` + "`" + `true` + "`" + ` to enable the certificate on the domain after the upload. Default ` + "`" + `false` + "`" + `. ## Attribute Reference The following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "expiration_date",
					Description: `The RFC3339 timestamp of when the certificate will be expired.`,
				},
				resource.Attribute{
					Name:        "issue_date",
					Description: `The RFC3339 timestamp of when the certificate was created.`,
				},
				resource.Attribute{
					Name:        "cn",
					Description: `The certificate Common Name.`,
				},
				resource.Attribute{
					Name:        "san",
					Description: `The certicate Subject Alternative Names.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "expiration_date",
					Description: `The RFC3339 timestamp of when the certificate will be expired.`,
				},
				resource.Attribute{
					Name:        "issue_date",
					Description: `The RFC3339 timestamp of when the certificate was created.`,
				},
				resource.Attribute{
					Name:        "cn",
					Description: `The certificate Common Name.`,
				},
				resource.Attribute{
					Name:        "san",
					Description: `The certicate Subject Alternative Names.`,
				},
			},
		},
	}

	resourcesMap = map[string]int{

		"gocache_gocache_domain":               0,
		"gocache_gocache_domain_dnssec":        1,
		"gocache_gocache_record":               2,
		"gocache_gocache_smart_rules_firewall": 3,
		"gocache_gocache_smart_rules_rewrite":  4,
		"gocache_gocache_smart_rules_settings": 5,
		"gocache_gocache_ssl_certificate":      6,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
