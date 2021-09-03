package constellix

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "constellix_dns_check",
			Category:         "Data Sources",
			ShortDescription: `Data source for DNS check resource`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of resource. Name should be unique. ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of the resource. Name should be unique.`,
				},
				resource.Attribute{
					Name:        "fqdn",
					Description: `(Required) A website address. It can be set only once`,
				},
				resource.Attribute{
					Name:        "resolver",
					Description: `(Required) A website address. It can be set only once`,
				},
				resource.Attribute{
					Name:        "check_sites",
					Description: `(Required) Site ids to check.`,
				},
				resource.Attribute{
					Name:        "notification_groups",
					Description: `(Optional) List of group IDs for the notification group of DNS Check.`,
				},
				resource.Attribute{
					Name:        "interval",
					Description: `(Optional) Check Interval. Allowed values are ` + "`" + `THIRTYSECONDS` + "`" + `, ` + "`" + `ONEMINUTE` + "`" + `, ` + "`" + `TWOMINUTES` + "`" + `, ` + "`" + `THREEMINUTES` + "`" + `, ` + "`" + `FOURMINUTES` + "`" + `, ` + "`" + `FIVEMINUTES` + "`" + `, ` + "`" + `TENMINUTES` + "`" + `, ` + "`" + `THIRTYMINUTES` + "`" + `, ` + "`" + `HALFDAY` + "`" + ` and ` + "`" + `DAY` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "interval_policy",
					Description: `(Optional) Agent Interval Run Policy. It specifies whether you want to run checks from one location or all. Allowed values are ` + "`" + `PARALLEL` + "`" + `, ` + "`" + `ONCEPERSITE` + "`" + ` and ` + "`" + `ONCEPERREGION` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "verification_policy",
					Description: `(Optional) Specifies how the check should be validated. Allowed values are ` + "`" + `SIMPLE` + "`" + ` and ` + "`" + `MAJORITY` + "`" + `. This parameter will only work with the ` + "`" + `interval_policy` + "`" + ` set to ` + "`" + `PARALLEL` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "expected_response",
					Description: `(Optional) Ip Address where DNS provided in the FQDN should resolved to in ideal conditions.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of the resource. Name should be unique.`,
				},
				resource.Attribute{
					Name:        "fqdn",
					Description: `(Required) A website address. It can be set only once`,
				},
				resource.Attribute{
					Name:        "resolver",
					Description: `(Required) A website address. It can be set only once`,
				},
				resource.Attribute{
					Name:        "check_sites",
					Description: `(Required) Site ids to check.`,
				},
				resource.Attribute{
					Name:        "notification_groups",
					Description: `(Optional) List of group IDs for the notification group of DNS Check.`,
				},
				resource.Attribute{
					Name:        "interval",
					Description: `(Optional) Check Interval. Allowed values are ` + "`" + `THIRTYSECONDS` + "`" + `, ` + "`" + `ONEMINUTE` + "`" + `, ` + "`" + `TWOMINUTES` + "`" + `, ` + "`" + `THREEMINUTES` + "`" + `, ` + "`" + `FOURMINUTES` + "`" + `, ` + "`" + `FIVEMINUTES` + "`" + `, ` + "`" + `TENMINUTES` + "`" + `, ` + "`" + `THIRTYMINUTES` + "`" + `, ` + "`" + `HALFDAY` + "`" + ` and ` + "`" + `DAY` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "interval_policy",
					Description: `(Optional) Agent Interval Run Policy. It specifies whether you want to run checks from one location or all. Allowed values are ` + "`" + `PARALLEL` + "`" + `, ` + "`" + `ONCEPERSITE` + "`" + ` and ` + "`" + `ONCEPERREGION` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "verification_policy",
					Description: `(Optional) Specifies how the check should be validated. Allowed values are ` + "`" + `SIMPLE` + "`" + ` and ` + "`" + `MAJORITY` + "`" + `. This parameter will only work with the ` + "`" + `interval_policy` + "`" + ` set to ` + "`" + `PARALLEL` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "expected_response",
					Description: `(Optional) Ip Address where DNS provided in the FQDN should resolved to in ideal conditions.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "constellix_domain",
			Category:         "Data Sources",
			ShortDescription: `Data source for Constellix domain`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the domain. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "has_gtd_regions",
					Description: `(Optional) GTD Region status of the domain. The Default value is false.`,
				},
				resource.Attribute{
					Name:        "has_geoip",
					Description: `(Optional) GTD Region status of the domain. The Default value is false.`,
				},
				resource.Attribute{
					Name:        "vanity_nameserver",
					Description: `(Optional) vanity nameserver of domain.`,
				},
				resource.Attribute{
					Name:        "nameserver_group",
					Description: `(Optional) Shows the nameserver group of domain. The Default nameserverGroup is 1.`,
				},
				resource.Attribute{
					Name:        "note",
					Description: `(Optional) Notes while creating the domain. The maximum length will be 1000 characters.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Id of tags applied on domain. The default value is empty.`,
				},
				resource.Attribute{
					Name:        "soa",
					Description: `(Optional) Object`,
				},
				resource.Attribute{
					Name:        "soa.primary_nameserver",
					Description: `(Optional) The Default value of SOA Primary Nameserver is "ns0.constellix.com.". However, it is possible to create a custom SOA record with differing values if required.`,
				},
				resource.Attribute{
					Name:        "soa.email",
					Description: `(Optional) An Email Address specifies the mailbox of the person responsible for this zone. The default value is "dns.constellix.com."`,
				},
				resource.Attribute{
					Name:        "soa.ttl",
					Description: `(Optional) The number of seconds that this SOA record will be cached in other resolving name servers. The Default value is "86400".TTL must be in between 0 and 2147483647`,
				},
				resource.Attribute{
					Name:        "soa.refresh",
					Description: `(Optional) The time interval (in seconds) before the zone should be refreshed. The recommended value – 86400 (24 Hours). The default value is 43200 (12 hours)`,
				},
				resource.Attribute{
					Name:        "soa.serial",
					Description: `(Optional) The starting serial number for the version of the zone. If the SOA record is applied to a domain that is already created (and thus already has a starting serial number), the existing serial number will be incremented by one. e.g 2015010196`,
				},
				resource.Attribute{
					Name:        "soa.retry",
					Description: `(Optional) The time interval (in seconds) before a failed refresh should be retried. Recommended value – 7200 (2 Hours). The default value is 1 hour`,
				},
				resource.Attribute{
					Name:        "soa.expire",
					Description: `(Optional) The time internal (in seconds) that specifies the upper limit on the time internally that can elapse before the zone is no longer authoritative. This is when the secondary name servers will expire if they are unable to refresh. Recommended value – up to 1209600`,
				},
				resource.Attribute{
					Name:        "soa.negcache",
					Description: `(Optional) The amount of time a record not found is cached. Recommended values can vary, between 180 and 172800 (3 min – 2 days). The default value is 180`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "has_gtd_regions",
					Description: `(Optional) GTD Region status of the domain. The Default value is false.`,
				},
				resource.Attribute{
					Name:        "has_geoip",
					Description: `(Optional) GTD Region status of the domain. The Default value is false.`,
				},
				resource.Attribute{
					Name:        "vanity_nameserver",
					Description: `(Optional) vanity nameserver of domain.`,
				},
				resource.Attribute{
					Name:        "nameserver_group",
					Description: `(Optional) Shows the nameserver group of domain. The Default nameserverGroup is 1.`,
				},
				resource.Attribute{
					Name:        "note",
					Description: `(Optional) Notes while creating the domain. The maximum length will be 1000 characters.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Id of tags applied on domain. The default value is empty.`,
				},
				resource.Attribute{
					Name:        "soa",
					Description: `(Optional) Object`,
				},
				resource.Attribute{
					Name:        "soa.primary_nameserver",
					Description: `(Optional) The Default value of SOA Primary Nameserver is "ns0.constellix.com.". However, it is possible to create a custom SOA record with differing values if required.`,
				},
				resource.Attribute{
					Name:        "soa.email",
					Description: `(Optional) An Email Address specifies the mailbox of the person responsible for this zone. The default value is "dns.constellix.com."`,
				},
				resource.Attribute{
					Name:        "soa.ttl",
					Description: `(Optional) The number of seconds that this SOA record will be cached in other resolving name servers. The Default value is "86400".TTL must be in between 0 and 2147483647`,
				},
				resource.Attribute{
					Name:        "soa.refresh",
					Description: `(Optional) The time interval (in seconds) before the zone should be refreshed. The recommended value – 86400 (24 Hours). The default value is 43200 (12 hours)`,
				},
				resource.Attribute{
					Name:        "soa.serial",
					Description: `(Optional) The starting serial number for the version of the zone. If the SOA record is applied to a domain that is already created (and thus already has a starting serial number), the existing serial number will be incremented by one. e.g 2015010196`,
				},
				resource.Attribute{
					Name:        "soa.retry",
					Description: `(Optional) The time interval (in seconds) before a failed refresh should be retried. Recommended value – 7200 (2 Hours). The default value is 1 hour`,
				},
				resource.Attribute{
					Name:        "soa.expire",
					Description: `(Optional) The time internal (in seconds) that specifies the upper limit on the time internally that can elapse before the zone is no longer authoritative. This is when the secondary name servers will expire if they are unable to refresh. Recommended value – up to 1209600`,
				},
				resource.Attribute{
					Name:        "soa.negcache",
					Description: `(Optional) The amount of time a record not found is cached. Recommended values can vary, between 180 and 172800 (3 min – 2 days). The default value is 180`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "constellix_http_check",
			Category:         "Data Sources",
			ShortDescription: `Data source for HTTP check resource`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of resource. Name should be unique. ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of the resource. Name should be unique.`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `(Optional) Host for the resource, for example "constellix.com". It can be set only once.`,
				},
				resource.Attribute{
					Name:        "ip_version",
					Description: `(Optional) Specifies the version of IP. It can be set only once.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional) Specifies the port number.`,
				},
				resource.Attribute{
					Name:        "protocol_type",
					Description: `(Optional) Specifies upper layer protocol like HTTP, HTTPs, etc.`,
				},
				resource.Attribute{
					Name:        "check_sites",
					Description: `(Optional) Site ids to check.`,
				},
				resource.Attribute{
					Name:        "notification_groups",
					Description: `(Optional) List of group IDs for the notification group of HTTP Check.`,
				},
				resource.Attribute{
					Name:        "interval",
					Description: `(Optional) Check Interval. Allowed values are ` + "`" + `THIRTYSECONDS` + "`" + `, ` + "`" + `ONEMINUTE` + "`" + `, ` + "`" + `TWOMINUTES` + "`" + `, ` + "`" + `THREEMINUTES` + "`" + `, ` + "`" + `FOURMINUTES` + "`" + `, ` + "`" + `FIVEMINUTES` + "`" + `, ` + "`" + `TENMINUTES` + "`" + `, ` + "`" + `THIRTYMINUTES` + "`" + `, ` + "`" + `HALFDAY` + "`" + ` and ` + "`" + `DAY` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "interval_policy",
					Description: `(Optional) Agent Interval Run Policy. It specifies whether you want to run checks from one location or all. Allowed values are ` + "`" + `PARALLEL` + "`" + `, ` + "`" + `ONCEPERSITE` + "`" + ` and ` + "`" + `ONCEPERREGION` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "verification_policy",
					Description: `(Optional) Specifies how the check should be validated. Allowed values are ` + "`" + `SIMPLE` + "`" + ` and ` + "`" + `MAJORITY` + "`" + `. This parameter will only work with the ` + "`" + `interval_policy` + "`" + ` set to ` + "`" + `PARALLEL` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "fqdn",
					Description: `(Optional) Fully qualified domain name of the URL should be checked.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Optional) In case of multi-page site, which path should be checked.`,
				},
				resource.Attribute{
					Name:        "search_string",
					Description: `(Optional) String to search in the first 2KB of resonse received.`,
				},
				resource.Attribute{
					Name:        "expected_status_code",
					Description: `(Optional) Expected HTTP status code for this check.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of the resource. Name should be unique.`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `(Optional) Host for the resource, for example "constellix.com". It can be set only once.`,
				},
				resource.Attribute{
					Name:        "ip_version",
					Description: `(Optional) Specifies the version of IP. It can be set only once.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional) Specifies the port number.`,
				},
				resource.Attribute{
					Name:        "protocol_type",
					Description: `(Optional) Specifies upper layer protocol like HTTP, HTTPs, etc.`,
				},
				resource.Attribute{
					Name:        "check_sites",
					Description: `(Optional) Site ids to check.`,
				},
				resource.Attribute{
					Name:        "notification_groups",
					Description: `(Optional) List of group IDs for the notification group of HTTP Check.`,
				},
				resource.Attribute{
					Name:        "interval",
					Description: `(Optional) Check Interval. Allowed values are ` + "`" + `THIRTYSECONDS` + "`" + `, ` + "`" + `ONEMINUTE` + "`" + `, ` + "`" + `TWOMINUTES` + "`" + `, ` + "`" + `THREEMINUTES` + "`" + `, ` + "`" + `FOURMINUTES` + "`" + `, ` + "`" + `FIVEMINUTES` + "`" + `, ` + "`" + `TENMINUTES` + "`" + `, ` + "`" + `THIRTYMINUTES` + "`" + `, ` + "`" + `HALFDAY` + "`" + ` and ` + "`" + `DAY` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "interval_policy",
					Description: `(Optional) Agent Interval Run Policy. It specifies whether you want to run checks from one location or all. Allowed values are ` + "`" + `PARALLEL` + "`" + `, ` + "`" + `ONCEPERSITE` + "`" + ` and ` + "`" + `ONCEPERREGION` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "verification_policy",
					Description: `(Optional) Specifies how the check should be validated. Allowed values are ` + "`" + `SIMPLE` + "`" + ` and ` + "`" + `MAJORITY` + "`" + `. This parameter will only work with the ` + "`" + `interval_policy` + "`" + ` set to ` + "`" + `PARALLEL` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "fqdn",
					Description: `(Optional) Fully qualified domain name of the URL should be checked.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Optional) In case of multi-page site, which path should be checked.`,
				},
				resource.Attribute{
					Name:        "search_string",
					Description: `(Optional) String to search in the first 2KB of resonse received.`,
				},
				resource.Attribute{
					Name:        "expected_status_code",
					Description: `(Optional) Expected HTTP status code for this check.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "constellix_tags",
			Category:         "Data Sources",
			ShortDescription: `Data source for Tags`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of record. Name should be unique. ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of record. Name should be unique.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of record. Name should be unique.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "constellix_tcp_check",
			Category:         "Data Sources",
			ShortDescription: `Data source for TCP check resource`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of resource. Name should be unique. ## Attribute Reference ##`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of the resource. Name should be unique.`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `(Optional) Host for the resource, for example "constellix.com". It can be set only once.`,
				},
				resource.Attribute{
					Name:        "ip_version",
					Description: `(Optional) Specifies the version of IP. It can be set only once.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional) Specifies the port number.`,
				},
				resource.Attribute{
					Name:        "check_sites",
					Description: `(Optional) Site ids to check.`,
				},
				resource.Attribute{
					Name:        "notification_groups",
					Description: `(Optional) List of group IDs for the notification group of TCP Check.`,
				},
				resource.Attribute{
					Name:        "interval",
					Description: `(Optional) Check Interval. Allowed values are ` + "`" + `THIRTYSECONDS` + "`" + `, ` + "`" + `ONEMINUTE` + "`" + `, ` + "`" + `TWOMINUTES` + "`" + `, ` + "`" + `THREEMINUTES` + "`" + `, ` + "`" + `FOURMINUTES` + "`" + `, ` + "`" + `FIVEMINUTES` + "`" + `, ` + "`" + `TENMINUTES` + "`" + `, ` + "`" + `THIRTYMINUTES` + "`" + `, ` + "`" + `HALFDAY` + "`" + ` and ` + "`" + `DAY` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "interval_policy",
					Description: `(Optional) Agent Interval Run Policy. It specifies whether you want to run checks from one location or all. Allowed values are ` + "`" + `PARALLEL` + "`" + `, ` + "`" + `ONCEPERSITE` + "`" + ` and ` + "`" + `ONCEPERREGION` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "verification_policy",
					Description: `(Optional) Specifies how the check should be validated. Allowed values are ` + "`" + `SIMPLE` + "`" + ` and ` + "`" + `MAJORITY` + "`" + `. This parameter will only work with the ` + "`" + `interval_policy` + "`" + ` set to ` + "`" + `PARALLEL` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "string_to_send",
					Description: `(Optional) String to send along with the check. It can be any parameter to the endpoint.`,
				},
				resource.Attribute{
					Name:        "string_to_receive",
					Description: `(Optional) String which should be received as a result of TCP check.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of the resource. Name should be unique.`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `(Optional) Host for the resource, for example "constellix.com". It can be set only once.`,
				},
				resource.Attribute{
					Name:        "ip_version",
					Description: `(Optional) Specifies the version of IP. It can be set only once.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional) Specifies the port number.`,
				},
				resource.Attribute{
					Name:        "check_sites",
					Description: `(Optional) Site ids to check.`,
				},
				resource.Attribute{
					Name:        "notification_groups",
					Description: `(Optional) List of group IDs for the notification group of TCP Check.`,
				},
				resource.Attribute{
					Name:        "interval",
					Description: `(Optional) Check Interval. Allowed values are ` + "`" + `THIRTYSECONDS` + "`" + `, ` + "`" + `ONEMINUTE` + "`" + `, ` + "`" + `TWOMINUTES` + "`" + `, ` + "`" + `THREEMINUTES` + "`" + `, ` + "`" + `FOURMINUTES` + "`" + `, ` + "`" + `FIVEMINUTES` + "`" + `, ` + "`" + `TENMINUTES` + "`" + `, ` + "`" + `THIRTYMINUTES` + "`" + `, ` + "`" + `HALFDAY` + "`" + ` and ` + "`" + `DAY` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "interval_policy",
					Description: `(Optional) Agent Interval Run Policy. It specifies whether you want to run checks from one location or all. Allowed values are ` + "`" + `PARALLEL` + "`" + `, ` + "`" + `ONCEPERSITE` + "`" + ` and ` + "`" + `ONCEPERREGION` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "verification_policy",
					Description: `(Optional) Specifies how the check should be validated. Allowed values are ` + "`" + `SIMPLE` + "`" + ` and ` + "`" + `MAJORITY` + "`" + `. This parameter will only work with the ` + "`" + `interval_policy` + "`" + ` set to ` + "`" + `PARALLEL` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "string_to_send",
					Description: `(Optional) String to send along with the check. It can be any parameter to the endpoint.`,
				},
				resource.Attribute{
					Name:        "string_to_receive",
					Description: `(Optional) String which should be received as a result of TCP check.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "constellix_template",
			Category:         "Data Sources",
			ShortDescription: `Data source for the DNS record templates.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Template names. e.g "sampletemplate". ## Attributes Reference ##`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `(Optional) Id of domain to be applied.`,
				},
				resource.Attribute{
					Name:        "has_gtd_region",
					Description: `(Optional) Enable/Disable GTD Region of the domain. The Default value is false.`,
				},
				resource.Attribute{
					Name:        "has_geoip",
					Description: `(Optional) Enable/Disable GEO IP. The Default value is false.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Optional) System generated template history version.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "domain",
					Description: `(Optional) Id of domain to be applied.`,
				},
				resource.Attribute{
					Name:        "has_gtd_region",
					Description: `(Optional) Enable/Disable GTD Region of the domain. The Default value is false.`,
				},
				resource.Attribute{
					Name:        "has_geoip",
					Description: `(Optional) Enable/Disable GEO IP. The Default value is false.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Optional) System generated template history version.`,
				},
			},
		},
	}

	dataSourcesMap = map[string]int{

		"constellix_dns_check":  0,
		"constellix_domain":     1,
		"constellix_http_check": 2,
		"constellix_tags":       3,
		"constellix_tcp_check":  4,
		"constellix_template":   5,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
