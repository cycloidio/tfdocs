package ultradns

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "ultradns_dirpool",
			Category:         "DIR-Pool",
			ShortDescription: `Get information of Directional (DIR) pool records in UltraDNS.`,
			Description: `

Use this data source to get detailed information of Directional (DIR) pool records.

`,
			Keywords: []string{
				"dir",
				"pool",
				"dirpool",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "zone_name",
					Description: `(Required) (String) Name of the zone.`,
				},
				resource.Attribute{
					Name:        "owner_name",
					Description: `(Required) (String) The domain name of the owner of the RRSet. Can be either a fully qualified domain name (FQDN) or a relative domain name. If provided as a FQDN, it must be contained within the zone name's FQDN.`,
				},
				resource.Attribute{
					Name:        "record_type",
					Description: `(Required) (String) Must be formatted as a well-known resource record type (A), or the corresponding number for the type (1).<br/> Below are the supported resource record types with the corresponding number:<br/> ` + "`" + `A (1)` + "`" + ` ## Attributes Reference In addition to all of the arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "pool_description",
					Description: `(Computed) (String) Allows for an additional description of the Directional (DIR) pool.`,
				},
				resource.Attribute{
					Name:        "conflict_resolve",
					Description: `(Computed) (String) When there is a conflict between a matching GeoIP group and a matching SourceIP group, this will determine which should take precedence. This only applies to a mixed pool (contains both GeoIP and SourceIP data).`,
				},
				resource.Attribute{
					Name:        "ignore_ecs",
					Description: `(Computed) (Boolean) Whether to ignore the EDNSO (which is an extended label type allowing for greater DNS messaging size) Client Subnet data when available in the DNS request.</br> ` + "`" + `false` + "`" + `= EDNSO data will be used for IP directional routing.</br> ` + "`" + `true` + "`" + ` = EDNSO data will not be used and IP directional routing decisions will always use the IP address of the recursive server.`,
				},
				resource.Attribute{
					Name:        "no_response",
					Description: `(Computed) (Block Set, Max: 1) Nested block describing certain geographical territories and IP addresses that will not get a response if they try to access the directional pool. The structure of this block is described below.`,
				},
				resource.Attribute{
					Name:        "rdata_info",
					Description: `(Computed) (Block Set, Min: 1) List of nested blocks describing the pool records. The structure of this block is described below. ### Nested ` + "`" + `rdata_info` + "`" + ` block has the following structure:`,
				},
				resource.Attribute{
					Name:        "rdata",
					Description: `(Computed) (String) The IPv4/IPv6 address, CNAME, MX, TXT, or SRV format data.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Computed) (String) The type for the specified pool record.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `(Computed) (Integer) The time to live (in seconds) for the corresponding record in rdata. Must be a value between 0 and 2147483647, inclusive.`,
				},
				resource.Attribute{
					Name:        "all_non_configured",
					Description: `(Computed) (Boolean) Indicates whether or not the associated rdata is used for all of the non-configured geographical territories and SourceIP ranges. At most, one entry in rdataInfo can have this set to ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "geo_group_name",
					Description: `(Computed) (String) The name of the GeoIP group.`,
				},
				resource.Attribute{
					Name:        "geo_codes",
					Description: `(Computed) (String List) The codes for the geographical territories that make up this group.`,
				},
				resource.Attribute{
					Name:        "ip_group_name",
					Description: `(Computed) (String) The name of the SourceIP group.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `(Computed) (Block Set) List of nested blocks describing the IP addresses and IP ranges this SourceIP group contains. The structure of this block is described below. ### Nested ` + "`" + `no_response` + "`" + ` block has the following structure:`,
				},
				resource.Attribute{
					Name:        "all_non_configured",
					Description: `(Computed) (Boolean) Indicates whether or not “no response” is returned for all of the non-configured geographical territories and IP ranges. This can only be set to ` + "`" + `true` + "`" + ` if there is no entry for rdataInfo, with allNonConfigured is set to ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "geo_group_name",
					Description: `(Computed) (String) The name for the “no response” GeoIP group.`,
				},
				resource.Attribute{
					Name:        "geo_codes",
					Description: `(Computed) (String List) The codes for the geographical territories that make up the “no response” group.`,
				},
				resource.Attribute{
					Name:        "ip_group_name",
					Description: `(Computed) (String) The name of the “no response” SourceIP group.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `(Computed) (Block Set) List of nested blocks describing the IP addresses and IP range for the “no response” SourceIP group. The structure of this block is described below. ### Nested ` + "`" + `ip` + "`" + ` block has the following structure:`,
				},
				resource.Attribute{
					Name:        "start",
					Description: `(Computed) (String) The starting IP address (IPv4 or IPv6).`,
				},
				resource.Attribute{
					Name:        "end",
					Description: `(Computed) (String) The ending IP address (IPv4 or IPv6).`,
				},
				resource.Attribute{
					Name:        "cidr",
					Description: `(Computed) (String) The CIDR format (IPv4 or IPv6) for an IP address range.`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `(Computed) (String) A single IPv4 or IPv6 address.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "pool_description",
					Description: `(Computed) (String) Allows for an additional description of the Directional (DIR) pool.`,
				},
				resource.Attribute{
					Name:        "conflict_resolve",
					Description: `(Computed) (String) When there is a conflict between a matching GeoIP group and a matching SourceIP group, this will determine which should take precedence. This only applies to a mixed pool (contains both GeoIP and SourceIP data).`,
				},
				resource.Attribute{
					Name:        "ignore_ecs",
					Description: `(Computed) (Boolean) Whether to ignore the EDNSO (which is an extended label type allowing for greater DNS messaging size) Client Subnet data when available in the DNS request.</br> ` + "`" + `false` + "`" + `= EDNSO data will be used for IP directional routing.</br> ` + "`" + `true` + "`" + ` = EDNSO data will not be used and IP directional routing decisions will always use the IP address of the recursive server.`,
				},
				resource.Attribute{
					Name:        "no_response",
					Description: `(Computed) (Block Set, Max: 1) Nested block describing certain geographical territories and IP addresses that will not get a response if they try to access the directional pool. The structure of this block is described below.`,
				},
				resource.Attribute{
					Name:        "rdata_info",
					Description: `(Computed) (Block Set, Min: 1) List of nested blocks describing the pool records. The structure of this block is described below. ### Nested ` + "`" + `rdata_info` + "`" + ` block has the following structure:`,
				},
				resource.Attribute{
					Name:        "rdata",
					Description: `(Computed) (String) The IPv4/IPv6 address, CNAME, MX, TXT, or SRV format data.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Computed) (String) The type for the specified pool record.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `(Computed) (Integer) The time to live (in seconds) for the corresponding record in rdata. Must be a value between 0 and 2147483647, inclusive.`,
				},
				resource.Attribute{
					Name:        "all_non_configured",
					Description: `(Computed) (Boolean) Indicates whether or not the associated rdata is used for all of the non-configured geographical territories and SourceIP ranges. At most, one entry in rdataInfo can have this set to ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "geo_group_name",
					Description: `(Computed) (String) The name of the GeoIP group.`,
				},
				resource.Attribute{
					Name:        "geo_codes",
					Description: `(Computed) (String List) The codes for the geographical territories that make up this group.`,
				},
				resource.Attribute{
					Name:        "ip_group_name",
					Description: `(Computed) (String) The name of the SourceIP group.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `(Computed) (Block Set) List of nested blocks describing the IP addresses and IP ranges this SourceIP group contains. The structure of this block is described below. ### Nested ` + "`" + `no_response` + "`" + ` block has the following structure:`,
				},
				resource.Attribute{
					Name:        "all_non_configured",
					Description: `(Computed) (Boolean) Indicates whether or not “no response” is returned for all of the non-configured geographical territories and IP ranges. This can only be set to ` + "`" + `true` + "`" + ` if there is no entry for rdataInfo, with allNonConfigured is set to ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "geo_group_name",
					Description: `(Computed) (String) The name for the “no response” GeoIP group.`,
				},
				resource.Attribute{
					Name:        "geo_codes",
					Description: `(Computed) (String List) The codes for the geographical territories that make up the “no response” group.`,
				},
				resource.Attribute{
					Name:        "ip_group_name",
					Description: `(Computed) (String) The name of the “no response” SourceIP group.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `(Computed) (Block Set) List of nested blocks describing the IP addresses and IP range for the “no response” SourceIP group. The structure of this block is described below. ### Nested ` + "`" + `ip` + "`" + ` block has the following structure:`,
				},
				resource.Attribute{
					Name:        "start",
					Description: `(Computed) (String) The starting IP address (IPv4 or IPv6).`,
				},
				resource.Attribute{
					Name:        "end",
					Description: `(Computed) (String) The ending IP address (IPv4 or IPv6).`,
				},
				resource.Attribute{
					Name:        "cidr",
					Description: `(Computed) (String) The CIDR format (IPv4 or IPv6) for an IP address range.`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `(Computed) (String) A single IPv4 or IPv6 address.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ultradns_probedns",
			Category:         "DNS-Probe",
			ShortDescription: `Get information of DNS probe records in UltraDNS.`,
			Description: `

Use this data source to get detailed information of DNS probe records.

`,
			Keywords: []string{
				"dns",
				"probe",
				"probedns",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "zone_name",
					Description: `(Required) (String) Name of the zone.`,
				},
				resource.Attribute{
					Name:        "owner_name",
					Description: `(Required) (String) The domain name of the owner of the RRSet. Can be either a fully qualified domain name (FQDN) or a relative domain name. If provided as a FQDN, it must be contained within the zone name's FQDN. The following arguments are used to filter the probes:`,
				},
				resource.Attribute{
					Name:        "guid",
					Description: `(Optional) (String) The internal id for this probe. -> ` + "`" + `guid` + "`" + ` can be fetched using UltraDNS REST API.`,
				},
				resource.Attribute{
					Name:        "interval",
					Description: `(Optional) (String) Length of time between probes in minutes.`,
				},
				resource.Attribute{
					Name:        "agents",
					Description: `(Optional) (String List) Locations that will be used for probing. The exact list must be provided if probes need to be filtered using agents. Valid values are: ` + "`" + `ASIA` + "`" + `, ` + "`" + `CHINA` + "`" + `, ` + "`" + `EUROPE_EAST` + "`" + `, ` + "`" + `EUROPE_WEST` + "`" + `, ` + "`" + `NORTH_AMERICA_CENTRAL` + "`" + `, ` + "`" + `NORTH_AMERICA_EAST` + "`" + `, ` + "`" + `NORTH_AMERICA_WEST` + "`" + `, ` + "`" + `SOUTH_AMERICA` + "`" + `, ` + "`" + `NEW_YORK` + "`" + `, ` + "`" + `PALO_ALTO` + "`" + `, ` + "`" + `DALLAS` + "`" + `, and ` + "`" + `AMSTERDAM` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "threshold",
					Description: `(Optional) (Integer) The probe threshold value.`,
				},
				resource.Attribute{
					Name:        "pool_record",
					Description: `(Optional) (String) The pool record associated with this probe. -> 1) If ` + "`" + `guid` + "`" + ` is provided, the probe with that guid is returned, and other filter options are not considered.</br> 2) If there is a conflict between probes due to filter options other than ` + "`" + `guid` + "`" + `, the last created probe is returned.</br> 3) If no probe is found for the filter options, an error is returned. ## Attributes Reference In addition to all of the arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Computed) (String) The Port that should be used for DNS lookup.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Computed) (String) Select the record type that the probe will check for. Valid values are ` + "`" + `NULL` + "`" + `, ` + "`" + `AXFR` + "`" + `, or any Resource Record Type.`,
				},
				resource.Attribute{
					Name:        "query_name",
					Description: `(Computed) (String) The name that should be queried.`,
				},
				resource.Attribute{
					Name:        "tcp_only",
					Description: `(Computed) (Boolean) Indicates whether or not the probe should use TCP only, or first UDP then TCP.`,
				},
				resource.Attribute{
					Name:        "response",
					Description: `(Computed) (Block Set, Max:1) Nested block describing the strings to match the response that will generate a warning or failure. The structure of this block is described below.`,
				},
				resource.Attribute{
					Name:        "run_limit",
					Description: `(Computed) (Block Set, Max:1) Nested block describing how long the probe should run. The structure of this block follows the same structure as the [` + "`" + `limit` + "`" + `](#nested-limit-block-has-the-following-structure) block described below.`,
				},
				resource.Attribute{
					Name:        "avg_run_limit",
					Description: `(Computed) (Block Set, Max:1) Nested block describing the mean (average) run-time for the five most recent probes that have run on each agent. This is only used for Traffic Controller pools. The structure of this block follows the same structure as the [` + "`" + `limit` + "`" + `](#nested-limit-block-has-the-following-structure) block described below. ### Nested ` + "`" + `limit` + "`" + ` block has the following structure:`,
				},
				resource.Attribute{
					Name:        "warning",
					Description: `(Computed) (Integer) Indicates how long (in seconds) the DNS Probe should wait, before a warning is generated.`,
				},
				resource.Attribute{
					Name:        "critical",
					Description: `(Computed) (Integer) Indicates how long (in seconds) the DNS Probe should wait, before a critical warning is generated.`,
				},
				resource.Attribute{
					Name:        "fail",
					Description: `(Computed) (Integer) Indicates how long (in seconds) the DNS Probe should wait, before causing the probe to fail. ### Nested ` + "`" + `response` + "`" + ` block has the following structure:`,
				},
				resource.Attribute{
					Name:        "warning",
					Description: `(Computed) (String) Will match exactly for records containing a single field response (i.e., A, CNAME, DNAME, NS, MB, MD, MF, MG, MR, PTR), and partial matches for record types with multiple field response. Fields separated by spaces will be combined with the matches to trigger a warning.`,
				},
				resource.Attribute{
					Name:        "critical",
					Description: `(Computed) (String) Will match exactly for records containing a single field response (i.e., A, CNAME, DNAME, NS, MB, MD, MF, MG, MR, PTR), and partial matches for record types with multiple field response. Fields separated by spaces will be combined with the matches to trigger a critical warning.`,
				},
				resource.Attribute{
					Name:        "fail",
					Description: `(Computed) (String) Will match exactly for records containing a single field response (i.e., A, CNAME, DNAME, NS, MB, MD, MF, MG, MR, PTR), and partial matches for record types with multiple field response. Fields separated by spaces will be combined with the matches to trigger a failure. -> ` + "`" + `warning` + "`" + ` and ` + "`" + `critical` + "`" + ` are only used for Traffic Controller pools.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "port",
					Description: `(Computed) (String) The Port that should be used for DNS lookup.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Computed) (String) Select the record type that the probe will check for. Valid values are ` + "`" + `NULL` + "`" + `, ` + "`" + `AXFR` + "`" + `, or any Resource Record Type.`,
				},
				resource.Attribute{
					Name:        "query_name",
					Description: `(Computed) (String) The name that should be queried.`,
				},
				resource.Attribute{
					Name:        "tcp_only",
					Description: `(Computed) (Boolean) Indicates whether or not the probe should use TCP only, or first UDP then TCP.`,
				},
				resource.Attribute{
					Name:        "response",
					Description: `(Computed) (Block Set, Max:1) Nested block describing the strings to match the response that will generate a warning or failure. The structure of this block is described below.`,
				},
				resource.Attribute{
					Name:        "run_limit",
					Description: `(Computed) (Block Set, Max:1) Nested block describing how long the probe should run. The structure of this block follows the same structure as the [` + "`" + `limit` + "`" + `](#nested-limit-block-has-the-following-structure) block described below.`,
				},
				resource.Attribute{
					Name:        "avg_run_limit",
					Description: `(Computed) (Block Set, Max:1) Nested block describing the mean (average) run-time for the five most recent probes that have run on each agent. This is only used for Traffic Controller pools. The structure of this block follows the same structure as the [` + "`" + `limit` + "`" + `](#nested-limit-block-has-the-following-structure) block described below. ### Nested ` + "`" + `limit` + "`" + ` block has the following structure:`,
				},
				resource.Attribute{
					Name:        "warning",
					Description: `(Computed) (Integer) Indicates how long (in seconds) the DNS Probe should wait, before a warning is generated.`,
				},
				resource.Attribute{
					Name:        "critical",
					Description: `(Computed) (Integer) Indicates how long (in seconds) the DNS Probe should wait, before a critical warning is generated.`,
				},
				resource.Attribute{
					Name:        "fail",
					Description: `(Computed) (Integer) Indicates how long (in seconds) the DNS Probe should wait, before causing the probe to fail. ### Nested ` + "`" + `response` + "`" + ` block has the following structure:`,
				},
				resource.Attribute{
					Name:        "warning",
					Description: `(Computed) (String) Will match exactly for records containing a single field response (i.e., A, CNAME, DNAME, NS, MB, MD, MF, MG, MR, PTR), and partial matches for record types with multiple field response. Fields separated by spaces will be combined with the matches to trigger a warning.`,
				},
				resource.Attribute{
					Name:        "critical",
					Description: `(Computed) (String) Will match exactly for records containing a single field response (i.e., A, CNAME, DNAME, NS, MB, MD, MF, MG, MR, PTR), and partial matches for record types with multiple field response. Fields separated by spaces will be combined with the matches to trigger a critical warning.`,
				},
				resource.Attribute{
					Name:        "fail",
					Description: `(Computed) (String) Will match exactly for records containing a single field response (i.e., A, CNAME, DNAME, NS, MB, MD, MF, MG, MR, PTR), and partial matches for record types with multiple field response. Fields separated by spaces will be combined with the matches to trigger a failure. -> ` + "`" + `warning` + "`" + ` and ` + "`" + `critical` + "`" + ` are only used for Traffic Controller pools.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ultradns_probehttp",
			Category:         "HTTP-Probe",
			ShortDescription: `Get information of HTTP probe records in UltraDNS.`,
			Description: `

Use this data source to get detailed information of HTTP probe records.

`,
			Keywords: []string{
				"http",
				"probe",
				"probehttp",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "zone_name",
					Description: `(Required) (String) Name of the zone.`,
				},
				resource.Attribute{
					Name:        "owner_name",
					Description: `(Required) (String) The domain name of the owner of the RRSet. Can be either a fully qualified domain name (FQDN) or a relative domain name. If provided as a FQDN, it must be contained within the zone name's FQDN. The following arguments are used to filter the probes:`,
				},
				resource.Attribute{
					Name:        "guid",
					Description: `(Optional) (String) The internal id for this probe. -> ` + "`" + `guid` + "`" + ` can be fetched using UltraDNS REST API.`,
				},
				resource.Attribute{
					Name:        "interval",
					Description: `(Optional) (String) Length of time between probes in minutes.`,
				},
				resource.Attribute{
					Name:        "agents",
					Description: `(Optional) (String List) Locations that will be used for probing. The exact list must be provided if probes need to be filtered using agents. Valid values are: ` + "`" + `ASIA` + "`" + `, ` + "`" + `CHINA` + "`" + `, ` + "`" + `EUROPE_EAST` + "`" + `, ` + "`" + `EUROPE_WEST` + "`" + `, ` + "`" + `NORTH_AMERICA_CENTRAL` + "`" + `, ` + "`" + `NORTH_AMERICA_EAST` + "`" + `, ` + "`" + `NORTH_AMERICA_WEST` + "`" + `, ` + "`" + `SOUTH_AMERICA` + "`" + `, ` + "`" + `NEW_YORK` + "`" + `, ` + "`" + `PALO_ALTO` + "`" + `, ` + "`" + `DALLAS` + "`" + `, and ` + "`" + `AMSTERDAM` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "threshold",
					Description: `(Optional) (Integer) The probe threshold value.`,
				},
				resource.Attribute{
					Name:        "pool_record",
					Description: `(Optional) (String) The pool record associated with this probe. -> 1) If ` + "`" + `guid` + "`" + ` is provided, the probe with that guid is returned, and other filter options are not considered.</br> 2) If there is a conflict between probes due to filter options other than ` + "`" + `guid` + "`" + `, the last created probe is returned.</br> 3) If no probe is found for the filter options, an error is returned. ## Attributes Reference In addition to all of the arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "transaction",
					Description: `(Computed) (Block Set) List of nested blocks describing the http requests sent for a single probe. The structure of this block is described below.`,
				},
				resource.Attribute{
					Name:        "total_limit",
					Description: `(Computed) (Block Set) Nested block describing the total amount of time spent on all http transactions. The structure of this block follows the same structure as the [` + "`" + `limit` + "`" + `](#nested-limit-block-has-the-following-structure) block described below. ### Nested ` + "`" + `transaction` + "`" + ` block has the following structure:`,
				},
				resource.Attribute{
					Name:        "method",
					Description: `(Computed) (String) HTTP method.`,
				},
				resource.Attribute{
					Name:        "protocol_version",
					Description: `(Computed) (String) HTTP protocol version.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Computed) (String) The URL that will be probed.`,
				},
				resource.Attribute{
					Name:        "transmitted_data",
					Description: `(Computed) (String) The data to send to the URL.`,
				},
				resource.Attribute{
					Name:        "follow_redirects",
					Description: `(Computed) (Boolean) Indicates whether or not to follow redirects.`,
				},
				resource.Attribute{
					Name:        "expected_response",
					Description: `(Computed) (String) The Expected Response code for probes to be returned as Successful.`,
				},
				resource.Attribute{
					Name:        "search_string",
					Description: `(Computed) (Block Set) Nested block describing the strings required to be searched for a probe’s successful response. This does not search the status line or headers. The structure of this block is described below.`,
				},
				resource.Attribute{
					Name:        "connect_limit",
					Description: `(Computed) (Block Set) Nested block describing how long the probe stays connected to the resource. The structure of this block follows the same structure as the [` + "`" + `limit` + "`" + `](#nested-limit-block-has-the-following-structure) block described below.`,
				},
				resource.Attribute{
					Name:        "avg_connect_limit",
					Description: `(Computed) (Block Set) Nested block describing the mean (average) time to connect for the five most recent probes that have run on each agent. This is only used for Traffic Controller pools. The structure of this block follows the same structure as the [` + "`" + `limit` + "`" + `](#nested-limit-block-has-the-following-structure) block described below.`,
				},
				resource.Attribute{
					Name:        "run_limit",
					Description: `(Computed) (Block Set) Nested block describing how long the probe should run. The structure of this block follows the same structure as the [` + "`" + `limit` + "`" + `](#nested-limit-block-has-the-following-structure) block described below.`,
				},
				resource.Attribute{
					Name:        "avg_run_limit",
					Description: `(Computed) (Block Set) Nested block describing the mean (average) run-time for the five most recent probes that have run on each agent. This is only used for Traffic Controller pools. The structure of this block follows the same structure as the [` + "`" + `limit` + "`" + `](#nested-limit-block-has-the-following-structure) block described below. ### Nested ` + "`" + `limit` + "`" + ` block has the following structure:`,
				},
				resource.Attribute{
					Name:        "warning",
					Description: `(Computed) (Integer) Indicates how long (in seconds) the HTTP Transactional Probe should wait, before a warning is generated.`,
				},
				resource.Attribute{
					Name:        "critical",
					Description: `(Computed) (Integer) Indicates how long (in seconds) the HTTP Transactional Probe should wait, before a critical warning is generated.`,
				},
				resource.Attribute{
					Name:        "fail",
					Description: `(Computed) (Integer) Indicates how long the (in seconds) HTTP Transactional Probe should wait, before causing the probe to fail. ### Nested ` + "`" + `search_string` + "`" + ` block has the following structure:`,
				},
				resource.Attribute{
					Name:        "warning",
					Description: `(Computed) (String) If the probe does not find the search string within the response, or does not match it as a regular expression, a warning will be generated.`,
				},
				resource.Attribute{
					Name:        "critical",
					Description: `(Computed) (String) If the probe does not find the search string within the response, or does not match it as a regular expression, a critical warning will be generated.`,
				},
				resource.Attribute{
					Name:        "fail",
					Description: `(Computed) (String) If the probe does not find the search string within the response, or does not match it as a regular expression, the probe will fail. -> ` + "`" + `warning` + "`" + ` and ` + "`" + `critical` + "`" + ` are only used for Traffic Controller pools.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "transaction",
					Description: `(Computed) (Block Set) List of nested blocks describing the http requests sent for a single probe. The structure of this block is described below.`,
				},
				resource.Attribute{
					Name:        "total_limit",
					Description: `(Computed) (Block Set) Nested block describing the total amount of time spent on all http transactions. The structure of this block follows the same structure as the [` + "`" + `limit` + "`" + `](#nested-limit-block-has-the-following-structure) block described below. ### Nested ` + "`" + `transaction` + "`" + ` block has the following structure:`,
				},
				resource.Attribute{
					Name:        "method",
					Description: `(Computed) (String) HTTP method.`,
				},
				resource.Attribute{
					Name:        "protocol_version",
					Description: `(Computed) (String) HTTP protocol version.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Computed) (String) The URL that will be probed.`,
				},
				resource.Attribute{
					Name:        "transmitted_data",
					Description: `(Computed) (String) The data to send to the URL.`,
				},
				resource.Attribute{
					Name:        "follow_redirects",
					Description: `(Computed) (Boolean) Indicates whether or not to follow redirects.`,
				},
				resource.Attribute{
					Name:        "expected_response",
					Description: `(Computed) (String) The Expected Response code for probes to be returned as Successful.`,
				},
				resource.Attribute{
					Name:        "search_string",
					Description: `(Computed) (Block Set) Nested block describing the strings required to be searched for a probe’s successful response. This does not search the status line or headers. The structure of this block is described below.`,
				},
				resource.Attribute{
					Name:        "connect_limit",
					Description: `(Computed) (Block Set) Nested block describing how long the probe stays connected to the resource. The structure of this block follows the same structure as the [` + "`" + `limit` + "`" + `](#nested-limit-block-has-the-following-structure) block described below.`,
				},
				resource.Attribute{
					Name:        "avg_connect_limit",
					Description: `(Computed) (Block Set) Nested block describing the mean (average) time to connect for the five most recent probes that have run on each agent. This is only used for Traffic Controller pools. The structure of this block follows the same structure as the [` + "`" + `limit` + "`" + `](#nested-limit-block-has-the-following-structure) block described below.`,
				},
				resource.Attribute{
					Name:        "run_limit",
					Description: `(Computed) (Block Set) Nested block describing how long the probe should run. The structure of this block follows the same structure as the [` + "`" + `limit` + "`" + `](#nested-limit-block-has-the-following-structure) block described below.`,
				},
				resource.Attribute{
					Name:        "avg_run_limit",
					Description: `(Computed) (Block Set) Nested block describing the mean (average) run-time for the five most recent probes that have run on each agent. This is only used for Traffic Controller pools. The structure of this block follows the same structure as the [` + "`" + `limit` + "`" + `](#nested-limit-block-has-the-following-structure) block described below. ### Nested ` + "`" + `limit` + "`" + ` block has the following structure:`,
				},
				resource.Attribute{
					Name:        "warning",
					Description: `(Computed) (Integer) Indicates how long (in seconds) the HTTP Transactional Probe should wait, before a warning is generated.`,
				},
				resource.Attribute{
					Name:        "critical",
					Description: `(Computed) (Integer) Indicates how long (in seconds) the HTTP Transactional Probe should wait, before a critical warning is generated.`,
				},
				resource.Attribute{
					Name:        "fail",
					Description: `(Computed) (Integer) Indicates how long the (in seconds) HTTP Transactional Probe should wait, before causing the probe to fail. ### Nested ` + "`" + `search_string` + "`" + ` block has the following structure:`,
				},
				resource.Attribute{
					Name:        "warning",
					Description: `(Computed) (String) If the probe does not find the search string within the response, or does not match it as a regular expression, a warning will be generated.`,
				},
				resource.Attribute{
					Name:        "critical",
					Description: `(Computed) (String) If the probe does not find the search string within the response, or does not match it as a regular expression, a critical warning will be generated.`,
				},
				resource.Attribute{
					Name:        "fail",
					Description: `(Computed) (String) If the probe does not find the search string within the response, or does not match it as a regular expression, the probe will fail. -> ` + "`" + `warning` + "`" + ` and ` + "`" + `critical` + "`" + ` are only used for Traffic Controller pools.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ultradns_probeping",
			Category:         "PING-Probe",
			ShortDescription: `Get information of PING probe records in UltraDNS.`,
			Description: `

Use this data source to get detailed information of PING probe records.

`,
			Keywords: []string{
				"ping",
				"probe",
				"probeping",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "zone_name",
					Description: `(Required) (String) Name of the zone.`,
				},
				resource.Attribute{
					Name:        "owner_name",
					Description: `(Required) (String) The domain name of the owner of the RRSet. Can be either a fully qualified domain name (FQDN) or a relative domain name. If provided as a FQDN, it must be contained within the zone name's FQDN. The following arguments are used to filter the probes:`,
				},
				resource.Attribute{
					Name:        "guid",
					Description: `(Optional) (String) The internal id for this probe. -> ` + "`" + `guid` + "`" + ` can be fetched using UltraDNS REST API.`,
				},
				resource.Attribute{
					Name:        "interval",
					Description: `(Optional) (String) Length of time between probes in minutes.`,
				},
				resource.Attribute{
					Name:        "agents",
					Description: `(Optional) (String List) Locations that will be used for probing. The exact list must be provided if probes need to be filtered using agents. Valid values are: ` + "`" + `ASIA` + "`" + `, ` + "`" + `CHINA` + "`" + `, ` + "`" + `EUROPE_EAST` + "`" + `, ` + "`" + `EUROPE_WEST` + "`" + `, ` + "`" + `NORTH_AMERICA_CENTRAL` + "`" + `, ` + "`" + `NORTH_AMERICA_EAST` + "`" + `, ` + "`" + `NORTH_AMERICA_WEST` + "`" + `, ` + "`" + `SOUTH_AMERICA` + "`" + `, ` + "`" + `NEW_YORK` + "`" + `, ` + "`" + `PALO_ALTO` + "`" + `, ` + "`" + `DALLAS` + "`" + `, and ` + "`" + `AMSTERDAM` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "threshold",
					Description: `(Optional) (Integer) The probe threshold value.`,
				},
				resource.Attribute{
					Name:        "pool_record",
					Description: `(Optional) (String) The pool record associated with this probe. -> 1) If ` + "`" + `guid` + "`" + ` is provided, the probe with that guid is returned, and other filter options are not considered.</br> 2) If there is a conflict between probes due to filter options other than ` + "`" + `guid` + "`" + `, the last created probe is returned.</br> 3) If no probe is found for the filter options, an error is returned. ## Attributes Reference In addition to all of the arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "packets",
					Description: `(Computed) (String) Number of ICMP packets to send.`,
				},
				resource.Attribute{
					Name:        "packet_size",
					Description: `(Computed) (String) Size of packets in bytes.`,
				},
				resource.Attribute{
					Name:        "loss_percent_limit",
					Description: `(Computed) (Block Set, Max:1) Nested block describing the acceptable percentage of packets lost, which will in turn, generate either a warning or a failure. The structure of this block follows the same structure as the [` + "`" + `limit` + "`" + `](#nested-limit-block-has-the-following-structure) block described below.`,
				},
				resource.Attribute{
					Name:        "total_limit",
					Description: `(Computed) (Block Set, Max:1) Nested block describing how long the probe should run in total for all pings. The structure of this block follows the same structure as the [` + "`" + `limit` + "`" + `](#nested-limit-block-has-the-following-structure) block described below.`,
				},
				resource.Attribute{
					Name:        "average_limit",
					Description: `(Computed) (Block Set, Max:1) Nested block describing the mean (average) time to connect for the five most recent probes that have run on each agent. This is only used for Traffic Controller pools. The structure of this block follows the same structure as the [` + "`" + `limit` + "`" + `](#nested-limit-block-has-the-following-structure) block described below.`,
				},
				resource.Attribute{
					Name:        "run_limit",
					Description: `(Computed) (Block Set, Max:1) Nested block describing how long the probe should run. The structure of this block follows the same structure as the [` + "`" + `limit` + "`" + `](#nested-limit-block-has-the-following-structure) block described below.`,
				},
				resource.Attribute{
					Name:        "avg_run_limit",
					Description: `(Computed) (Block Set, Max:1) Nested block describing the mean (average) run-time for the five most recent probes that have run on each agent. This is only used for Traffic Controller pools. The structure of this block follows the same structure as the [` + "`" + `limit` + "`" + `](#nested-limit-block-has-the-following-structure) block described below. ### Nested ` + "`" + `limit` + "`" + ` block has the following structure:`,
				},
				resource.Attribute{
					Name:        "warning",
					Description: `(Computed) (Integer) Indicates how long (in seconds, or by percentage value) the PING Probe should wait, before a warning is generated.`,
				},
				resource.Attribute{
					Name:        "critical",
					Description: `(Computed) (Integer) Indicates how long (in seconds, or by percentage value) the PING Probe should wait, before a critical warning is generated.`,
				},
				resource.Attribute{
					Name:        "fail",
					Description: `(Computed) (Integer) Indicates how long (in seconds, or by percentage value) the PING Probe should wait, before causing the probe to fail. -> ` + "`" + `warning` + "`" + ` and ` + "`" + `critical` + "`" + ` are only used for Traffic Controller pools.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "packets",
					Description: `(Computed) (String) Number of ICMP packets to send.`,
				},
				resource.Attribute{
					Name:        "packet_size",
					Description: `(Computed) (String) Size of packets in bytes.`,
				},
				resource.Attribute{
					Name:        "loss_percent_limit",
					Description: `(Computed) (Block Set, Max:1) Nested block describing the acceptable percentage of packets lost, which will in turn, generate either a warning or a failure. The structure of this block follows the same structure as the [` + "`" + `limit` + "`" + `](#nested-limit-block-has-the-following-structure) block described below.`,
				},
				resource.Attribute{
					Name:        "total_limit",
					Description: `(Computed) (Block Set, Max:1) Nested block describing how long the probe should run in total for all pings. The structure of this block follows the same structure as the [` + "`" + `limit` + "`" + `](#nested-limit-block-has-the-following-structure) block described below.`,
				},
				resource.Attribute{
					Name:        "average_limit",
					Description: `(Computed) (Block Set, Max:1) Nested block describing the mean (average) time to connect for the five most recent probes that have run on each agent. This is only used for Traffic Controller pools. The structure of this block follows the same structure as the [` + "`" + `limit` + "`" + `](#nested-limit-block-has-the-following-structure) block described below.`,
				},
				resource.Attribute{
					Name:        "run_limit",
					Description: `(Computed) (Block Set, Max:1) Nested block describing how long the probe should run. The structure of this block follows the same structure as the [` + "`" + `limit` + "`" + `](#nested-limit-block-has-the-following-structure) block described below.`,
				},
				resource.Attribute{
					Name:        "avg_run_limit",
					Description: `(Computed) (Block Set, Max:1) Nested block describing the mean (average) run-time for the five most recent probes that have run on each agent. This is only used for Traffic Controller pools. The structure of this block follows the same structure as the [` + "`" + `limit` + "`" + `](#nested-limit-block-has-the-following-structure) block described below. ### Nested ` + "`" + `limit` + "`" + ` block has the following structure:`,
				},
				resource.Attribute{
					Name:        "warning",
					Description: `(Computed) (Integer) Indicates how long (in seconds, or by percentage value) the PING Probe should wait, before a warning is generated.`,
				},
				resource.Attribute{
					Name:        "critical",
					Description: `(Computed) (Integer) Indicates how long (in seconds, or by percentage value) the PING Probe should wait, before a critical warning is generated.`,
				},
				resource.Attribute{
					Name:        "fail",
					Description: `(Computed) (Integer) Indicates how long (in seconds, or by percentage value) the PING Probe should wait, before causing the probe to fail. -> ` + "`" + `warning` + "`" + ` and ` + "`" + `critical` + "`" + ` are only used for Traffic Controller pools.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ultradns_rdpool",
			Category:         "RD-Pool",
			ShortDescription: `Get information of Resource Distribution (RD) pool records in UltraDNS.`,
			Description: `

Use this data source to get detailed information of Resource Distribution (RD) pool records.

`,
			Keywords: []string{
				"rd",
				"pool",
				"rdpool",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "zone_name",
					Description: `(Required) (String) Name of the zone.`,
				},
				resource.Attribute{
					Name:        "owner_name",
					Description: `(Required) (String) The domain name of the owner of the RRSet. Can be either a fully qualified domain name (FQDN) or a relative domain name. If provided as a FQDN, it must be contained within the zone name's FQDN.`,
				},
				resource.Attribute{
					Name:        "record_type",
					Description: `(Required) (String) Must be formatted as the well-known resource record type (A or AAAA) or the corresponding number for the type (1 or 28).<br/> Below are the supported resource record types with the corresponding number:<br/> ` + "`" + `A (1)` + "`" + ` ` + "`" + `AAAA (28)` + "`" + ` ## Attributes Reference In addition to all of the arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `(Computed) (Integer) The time to live (in seconds) for the record. Must be a value between 0 and 2147483647, inclusive.`,
				},
				resource.Attribute{
					Name:        "record_data",
					Description: `(Computed) (String List) The list of IPv4 or IPv6 addresses.`,
				},
				resource.Attribute{
					Name:        "order",
					Description: `(Computed) (String) The order of the records will be returned in.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Computed) (String) An optional description of the RD pool.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ttl",
					Description: `(Computed) (Integer) The time to live (in seconds) for the record. Must be a value between 0 and 2147483647, inclusive.`,
				},
				resource.Attribute{
					Name:        "record_data",
					Description: `(Computed) (String List) The list of IPv4 or IPv6 addresses.`,
				},
				resource.Attribute{
					Name:        "order",
					Description: `(Computed) (String) The order of the records will be returned in.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Computed) (String) An optional description of the RD pool.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ultradns_record",
			Category:         "Record",
			ShortDescription: `Get information of standard DNS records in UltraDNS.`,
			Description: `

Use this data source to get detailed information of standard DNS records.

`,
			Keywords: []string{
				"record",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "zone_name",
					Description: `(Required) (String) Name of the zone.`,
				},
				resource.Attribute{
					Name:        "owner_name",
					Description: `(Required) (String) The domain name of the owner of the RRSet. Can be either a fully qualified domain name (FQDN) or a relative domain name. If provided as a FQDN, it must be contained within the zone name's FQDN.`,
				},
				resource.Attribute{
					Name:        "record_type",
					Description: `(Required) (String) Must be formatted as the well-known resource record type (A, AAAA, TXT, etc.) or the corresponding number for the type, between 1 and 65535.<br/> Below are the supported resource record types with the corresponding number:<br/> ` + "`" + `A (1)` + "`" + ` ` + "`" + `NS (2)` + "`" + ` ` + "`" + `CNAME (5)` + "`" + ` ` + "`" + `SOA (6)` + "`" + ` ` + "`" + `PTR (12)` + "`" + ` ` + "`" + `MX (15)` + "`" + ` ` + "`" + `TXT (16)` + "`" + ` ` + "`" + `AAAA (28)` + "`" + ` ` + "`" + `SRV (33)` + "`" + ` ` + "`" + `SSHFP (44)` + "`" + ` ` + "`" + `APEXALIAS (65282)` + "`" + ` ## Attributes Reference In addition to all of the arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `(Computed) (Integer) The time to live (in seconds) for the record. Must be a value between 0 and 2147483647, inclusive.`,
				},
				resource.Attribute{
					Name:        "record_data",
					Description: `(Computed) (String List) The data for the record displayed as the BIND presentation format for the specified RRTYPE.<br/> __Note__ In SOA records the serial number is ignored and removed from the data: ` + "`" + `["mname rname refresh retry expire minimum"]` + "`" + ` ` + "`" + `` + "`" + `` + "`" + `["ns.example.com admin@example.com 7200 3600 1209600 36000"]` + "`" + `` + "`" + `` + "`" + ` For a SRV record, the format of data is ` + "`" + `["priority weight port target"]` + "`" + ` ` + "`" + `` + "`" + `` + "`" + `["2 2 523 example.com."]` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ttl",
					Description: `(Computed) (Integer) The time to live (in seconds) for the record. Must be a value between 0 and 2147483647, inclusive.`,
				},
				resource.Attribute{
					Name:        "record_data",
					Description: `(Computed) (String List) The data for the record displayed as the BIND presentation format for the specified RRTYPE.<br/> __Note__ In SOA records the serial number is ignored and removed from the data: ` + "`" + `["mname rname refresh retry expire minimum"]` + "`" + ` ` + "`" + `` + "`" + `` + "`" + `["ns.example.com admin@example.com 7200 3600 1209600 36000"]` + "`" + `` + "`" + `` + "`" + ` For a SRV record, the format of data is ` + "`" + `["priority weight port target"]` + "`" + ` ` + "`" + `` + "`" + `` + "`" + `["2 2 523 example.com."]` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ultradns_sbpool",
			Category:         "SB-Pool",
			ShortDescription: `Get information of SiteBacker (SB) pool records in UltraDNS.`,
			Description: `

Use this data source to get detailed information of SiteBacker (SB) pool records.

`,
			Keywords: []string{
				"sb",
				"pool",
				"sbpool",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "zone_name",
					Description: `(Required) (String) Name of the zone.`,
				},
				resource.Attribute{
					Name:        "owner_name",
					Description: `(Required) (String) The domain name of the owner of the RRSet. Can be either a fully qualified domain name (FQDN) or a relative domain name. If provided as a FQDN, it must be contained within the zone name's FQDN.`,
				},
				resource.Attribute{
					Name:        "record_type",
					Description: `(Required) (String) Must be formatted as a well-known resource record type (A), or the corresponding number for the type (1).<br/> Below are the supported resource record types with the corresponding number:<br/> ` + "`" + `A (1)` + "`" + ` ## Attributes Reference In addition to all of the arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `(Computed) (Integer) The time to live (in seconds) for the record. Must be a value between 0 and 2147483647, inclusive.`,
				},
				resource.Attribute{
					Name:        "pool_description",
					Description: `(Computed) (String) An optional description of the SiteBacker (SB) field.`,
				},
				resource.Attribute{
					Name:        "run_probes",
					Description: `(Computed) (Boolean) Indicates whether or not the probes are run for this pool.`,
				},
				resource.Attribute{
					Name:        "act_on_probes",
					Description: `(Computed) (Boolean) Indicates whether or not pool records will be enabled (true) or disabled (false) when probes are run.`,
				},
				resource.Attribute{
					Name:        "order",
					Description: `(Computed) (String) Indicates the order of the records returned by the resolver for the SiteBacker pool.`,
				},
				resource.Attribute{
					Name:        "failure_threshold",
					Description: `(Computed) (Integer) The minimum number of records that must fail for a pool to be labeled 'FAILED'. If the number of failed records in the pool is greater than or equal to the 'Failure Threshold' value, the pool will be labeled FAILED.<br/> For example, a pool with six priority records, one all-fail record, and the Failure Threshold value is set to four (4). If four or more priority records are not available to serve, the pool will be labeled FAILED, and the all-fail record will be served.<br/>`,
				},
				resource.Attribute{
					Name:        "max_active",
					Description: `(Computed) (Integer) Specifies the maximum number of active servers in a pool and determines when SiteBacker takes backup servers offline.<br/> For example, consider a pool with six servers. Setting Max Active to 4 means SiteBacker takes two servers offline and sends the four active records in the answer.`,
				},
				resource.Attribute{
					Name:        "max_served",
					Description: `(Computed) (Integer) Determines the number of record answers for each query. This is typically All Active records, or a subset of Max Active.`,
				},
				resource.Attribute{
					Name:        "backup_record",
					Description: `(Computed) (Block Set) List of nested blocks describing the information of backup records for the SiteBacker pool. Specifies the records to be served if all other records fail. There can be one or more A records used as backup records, or a single CNAME record. The structure of this block is described below.`,
				},
				resource.Attribute{
					Name:        "rdata_info",
					Description: `(Computed) (Block Set) List of nested blocks describing the pool records. The structure of this block is described below.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Computed) (String) Current status of the serving record. Valid values are:</br> ` + "`" + `OK` + "`" + `- If the number of records serving is equal to the Max Active value, and all the active records are top priority records.</br> For example, if a pool has a Max Active of 1 and the Priority 1 record is serving.</br> ` + "`" + `WARNING` + "`" + ` – If the number of records serving is equal to the Max Active value, and the active records are not top priority records.</br> For example, if a pool has a Max Active of 1 and the Priority 1 records is not serving and the Priority 2 record is serving.</br> ` + "`" + `CRITICAL` + "`" + ` – If the number of records serving is less than the Max Active value, or the All Fail record is being served.</br> For example, if a pool has a Max Active of 2, and only one record is serving.</br> ` + "`" + `FAILED` + "`" + ` - If the FailureThreshold value is 0 or null, and no records are serving, and there is no All Fail record configured.</br>OR</br>If the number of priority records not available to serve equals or exceeds the FailureThreshold’s value.</br> For example, if the Failure Threshold value is 3, and there are 3 or more Priority Records that are not available to serve. ### Nested ` + "`" + `backup_record` + "`" + ` block has the following structure:`,
				},
				resource.Attribute{
					Name:        "rdata",
					Description: `(Computed) (String) The IPv4 address or CNAME for the backup record.`,
				},
				resource.Attribute{
					Name:        "failover_delay",
					Description: `(Computed) (Integer) Specifies the time, between 0 – 30 minutes, that SiteBacker waits after detecting that the pool record has failed, prior to activating the primary records. Default value set to 0.`,
				},
				resource.Attribute{
					Name:        "available_to_serve",
					Description: `(Computed) (Boolean) Indicates whether the pool's backup record is active and available to serve records. ### Nested ` + "`" + `rdata_info` + "`" + ` block has the following structure:`,
				},
				resource.Attribute{
					Name:        "rdata",
					Description: `(Computed) (String) The IPv4 address or CNAME.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `(Computed) (Integer) Indicates the serving preference for this pool record.`,
				},
				resource.Attribute{
					Name:        "threshold",
					Description: `(Computed) (Integer) Specifies how many probes must agree before the record state is changed.`,
				},
				resource.Attribute{
					Name:        "failover_delay",
					Description: `(Computed) (Integer) Specifies the time, between 0 – 30 minutes, that SiteBacker waits after detecting that the pool record has failed, prior to activating the secondary records. Default value set to 0.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `(Computed) (String) The current state of the pool record.`,
				},
				resource.Attribute{
					Name:        "run_probes",
					Description: `(Computed) (Boolean) Indicates whether or not probes are run for this pool record.`,
				},
				resource.Attribute{
					Name:        "available_to_serve",
					Description: `(Computed) (Boolean) Indicates whether the pool record is active and available to serve records.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ttl",
					Description: `(Computed) (Integer) The time to live (in seconds) for the record. Must be a value between 0 and 2147483647, inclusive.`,
				},
				resource.Attribute{
					Name:        "pool_description",
					Description: `(Computed) (String) An optional description of the SiteBacker (SB) field.`,
				},
				resource.Attribute{
					Name:        "run_probes",
					Description: `(Computed) (Boolean) Indicates whether or not the probes are run for this pool.`,
				},
				resource.Attribute{
					Name:        "act_on_probes",
					Description: `(Computed) (Boolean) Indicates whether or not pool records will be enabled (true) or disabled (false) when probes are run.`,
				},
				resource.Attribute{
					Name:        "order",
					Description: `(Computed) (String) Indicates the order of the records returned by the resolver for the SiteBacker pool.`,
				},
				resource.Attribute{
					Name:        "failure_threshold",
					Description: `(Computed) (Integer) The minimum number of records that must fail for a pool to be labeled 'FAILED'. If the number of failed records in the pool is greater than or equal to the 'Failure Threshold' value, the pool will be labeled FAILED.<br/> For example, a pool with six priority records, one all-fail record, and the Failure Threshold value is set to four (4). If four or more priority records are not available to serve, the pool will be labeled FAILED, and the all-fail record will be served.<br/>`,
				},
				resource.Attribute{
					Name:        "max_active",
					Description: `(Computed) (Integer) Specifies the maximum number of active servers in a pool and determines when SiteBacker takes backup servers offline.<br/> For example, consider a pool with six servers. Setting Max Active to 4 means SiteBacker takes two servers offline and sends the four active records in the answer.`,
				},
				resource.Attribute{
					Name:        "max_served",
					Description: `(Computed) (Integer) Determines the number of record answers for each query. This is typically All Active records, or a subset of Max Active.`,
				},
				resource.Attribute{
					Name:        "backup_record",
					Description: `(Computed) (Block Set) List of nested blocks describing the information of backup records for the SiteBacker pool. Specifies the records to be served if all other records fail. There can be one or more A records used as backup records, or a single CNAME record. The structure of this block is described below.`,
				},
				resource.Attribute{
					Name:        "rdata_info",
					Description: `(Computed) (Block Set) List of nested blocks describing the pool records. The structure of this block is described below.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Computed) (String) Current status of the serving record. Valid values are:</br> ` + "`" + `OK` + "`" + `- If the number of records serving is equal to the Max Active value, and all the active records are top priority records.</br> For example, if a pool has a Max Active of 1 and the Priority 1 record is serving.</br> ` + "`" + `WARNING` + "`" + ` – If the number of records serving is equal to the Max Active value, and the active records are not top priority records.</br> For example, if a pool has a Max Active of 1 and the Priority 1 records is not serving and the Priority 2 record is serving.</br> ` + "`" + `CRITICAL` + "`" + ` – If the number of records serving is less than the Max Active value, or the All Fail record is being served.</br> For example, if a pool has a Max Active of 2, and only one record is serving.</br> ` + "`" + `FAILED` + "`" + ` - If the FailureThreshold value is 0 or null, and no records are serving, and there is no All Fail record configured.</br>OR</br>If the number of priority records not available to serve equals or exceeds the FailureThreshold’s value.</br> For example, if the Failure Threshold value is 3, and there are 3 or more Priority Records that are not available to serve. ### Nested ` + "`" + `backup_record` + "`" + ` block has the following structure:`,
				},
				resource.Attribute{
					Name:        "rdata",
					Description: `(Computed) (String) The IPv4 address or CNAME for the backup record.`,
				},
				resource.Attribute{
					Name:        "failover_delay",
					Description: `(Computed) (Integer) Specifies the time, between 0 – 30 minutes, that SiteBacker waits after detecting that the pool record has failed, prior to activating the primary records. Default value set to 0.`,
				},
				resource.Attribute{
					Name:        "available_to_serve",
					Description: `(Computed) (Boolean) Indicates whether the pool's backup record is active and available to serve records. ### Nested ` + "`" + `rdata_info` + "`" + ` block has the following structure:`,
				},
				resource.Attribute{
					Name:        "rdata",
					Description: `(Computed) (String) The IPv4 address or CNAME.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `(Computed) (Integer) Indicates the serving preference for this pool record.`,
				},
				resource.Attribute{
					Name:        "threshold",
					Description: `(Computed) (Integer) Specifies how many probes must agree before the record state is changed.`,
				},
				resource.Attribute{
					Name:        "failover_delay",
					Description: `(Computed) (Integer) Specifies the time, between 0 – 30 minutes, that SiteBacker waits after detecting that the pool record has failed, prior to activating the secondary records. Default value set to 0.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `(Computed) (String) The current state of the pool record.`,
				},
				resource.Attribute{
					Name:        "run_probes",
					Description: `(Computed) (Boolean) Indicates whether or not probes are run for this pool record.`,
				},
				resource.Attribute{
					Name:        "available_to_serve",
					Description: `(Computed) (Boolean) Indicates whether the pool record is active and available to serve records.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ultradns_sfpool",
			Category:         "SF-Pool",
			ShortDescription: `Get information of Simple Monitor/Failover (SF) pool records in UltraDNS.`,
			Description: `

Use this data source to get detailed information of Simple Monitor/Failover (SF) pool records.

`,
			Keywords: []string{
				"sf",
				"pool",
				"sfpool",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "zone_name",
					Description: `(Required) (String) Name of the zone.`,
				},
				resource.Attribute{
					Name:        "owner_name",
					Description: `(Required) (String) The domain name of the owner of the RRSet. Can be either a fully qualified domain name (FQDN) or a relative domain name. If provided as a FQDN, it must be contained within the zone name's FQDN.`,
				},
				resource.Attribute{
					Name:        "record_type",
					Description: `(Required) (String) Must be formatted as a well-known resource record type (A or AAAA) or the corresponding number for the type (1 or 28).<br/> Below are the supported resource record types with the corresponding number:<br/> ` + "`" + `A (1)` + "`" + ` ` + "`" + `AAAA (28)` + "`" + ` ## Attributes Reference In addition to all of the arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `(Computed) (Integer) The time to live (in seconds) for the record. Must be a value between 0 and 2147483647, inclusive.`,
				},
				resource.Attribute{
					Name:        "record_data",
					Description: `(Computed) (String List) The list of IPv4 or IPv6 addresses.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Computed) (String) Current status of the serving record. Valid values are:</br> ` + "`" + `OK` + "`" + ` – Live record is being served.</br> ` + "`" + `CRITICAL` + "`" + ` – The backup record is being served due to the monitor detecting a failure.</br> ` + "`" + `MANUAL` + "`" + ` – The backup record is being served due to the user forcing the live record to be inactive.`,
				},
				resource.Attribute{
					Name:        "region_failure_sensitivity",
					Description: `(Computed) (String) Specifies the sensitivity to the number of regions that need to fail for the backup record to be made active.`,
				},
				resource.Attribute{
					Name:        "live_record_description",
					Description: `(Computed) (String) An optional description of the live record.`,
				},
				resource.Attribute{
					Name:        "pool_description",
					Description: `(Computed) (String) An optional description of the Simple Failover field.`,
				},
				resource.Attribute{
					Name:        "monitor",
					Description: `(Computed) (Block Set) Nested block describing the information for the monitor. The structure of this block is described below.`,
				},
				resource.Attribute{
					Name:        "backup_record",
					Description: `(Computed) (Block Set) Nested block describing the information for the backup record. The structure of this block is described below. ### Nested ` + "`" + `monitor` + "`" + ` block has the following structure:`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Computed) (String) Monitored URL. A full URL including the protocol, host, and URI. Valid protocols are HTTP and HTTPS.`,
				},
				resource.Attribute{
					Name:        "method",
					Description: `(Computed) (String) HTTP method used to connect to the monitored URL.`,
				},
				resource.Attribute{
					Name:        "transmitted_data",
					Description: `(Computed) (String) If a monitor is sending a POST, this is the data that is sent as the body of the request.`,
				},
				resource.Attribute{
					Name:        "search_string",
					Description: `(Computed) (String) A string that is checked against the returned data from the request. ### Nested ` + "`" + `backup_record` + "`" + ` block has the following structure:`,
				},
				resource.Attribute{
					Name:        "rdata",
					Description: `(Computed) (String) An IPv4 or IPv6 address as a backup record.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Computed) (String) An optional description for the backup record.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ttl",
					Description: `(Computed) (Integer) The time to live (in seconds) for the record. Must be a value between 0 and 2147483647, inclusive.`,
				},
				resource.Attribute{
					Name:        "record_data",
					Description: `(Computed) (String List) The list of IPv4 or IPv6 addresses.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Computed) (String) Current status of the serving record. Valid values are:</br> ` + "`" + `OK` + "`" + ` – Live record is being served.</br> ` + "`" + `CRITICAL` + "`" + ` – The backup record is being served due to the monitor detecting a failure.</br> ` + "`" + `MANUAL` + "`" + ` – The backup record is being served due to the user forcing the live record to be inactive.`,
				},
				resource.Attribute{
					Name:        "region_failure_sensitivity",
					Description: `(Computed) (String) Specifies the sensitivity to the number of regions that need to fail for the backup record to be made active.`,
				},
				resource.Attribute{
					Name:        "live_record_description",
					Description: `(Computed) (String) An optional description of the live record.`,
				},
				resource.Attribute{
					Name:        "pool_description",
					Description: `(Computed) (String) An optional description of the Simple Failover field.`,
				},
				resource.Attribute{
					Name:        "monitor",
					Description: `(Computed) (Block Set) Nested block describing the information for the monitor. The structure of this block is described below.`,
				},
				resource.Attribute{
					Name:        "backup_record",
					Description: `(Computed) (Block Set) Nested block describing the information for the backup record. The structure of this block is described below. ### Nested ` + "`" + `monitor` + "`" + ` block has the following structure:`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Computed) (String) Monitored URL. A full URL including the protocol, host, and URI. Valid protocols are HTTP and HTTPS.`,
				},
				resource.Attribute{
					Name:        "method",
					Description: `(Computed) (String) HTTP method used to connect to the monitored URL.`,
				},
				resource.Attribute{
					Name:        "transmitted_data",
					Description: `(Computed) (String) If a monitor is sending a POST, this is the data that is sent as the body of the request.`,
				},
				resource.Attribute{
					Name:        "search_string",
					Description: `(Computed) (String) A string that is checked against the returned data from the request. ### Nested ` + "`" + `backup_record` + "`" + ` block has the following structure:`,
				},
				resource.Attribute{
					Name:        "rdata",
					Description: `(Computed) (String) An IPv4 or IPv6 address as a backup record.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Computed) (String) An optional description for the backup record.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ultradns_slbpool",
			Category:         "SLB-Pool",
			ShortDescription: `Get information of Simple Load Balancing (SLB) pool records in UltraDNS.`,
			Description: `

Use this data source to get detailed information of Simple Load Balancing (SLB) pool records.

`,
			Keywords: []string{
				"slb",
				"pool",
				"slbpool",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "zone_name",
					Description: `(Required) (String) Name of the zone.`,
				},
				resource.Attribute{
					Name:        "owner_name",
					Description: `(Required) (String) The domain name of the owner of the RRSet. Can be either a fully qualified domain name (FQDN) or a relative domain name. If provided as a FQDN, it must be contained within the zone name's FQDN.`,
				},
				resource.Attribute{
					Name:        "record_type",
					Description: `(Required) (String) Must be formatted as a well-known resource record type (A or AAAA), or the corresponding number for the type (1 or 28).<br/> Below are the supported resource record types with the corresponding number:<br/> ` + "`" + `A (1)` + "`" + ` ` + "`" + `AAAA (28)` + "`" + ` ## Attributes Reference In addition to all of the arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `(Computed) (Integer) The time to live (in seconds) for the record. Must be a value between 0 and 2147483647, inclusive.`,
				},
				resource.Attribute{
					Name:        "region_failure_sensitivity",
					Description: `(Computed) (String) Specifies the sensitivity to the number of regions that need to fail for the backup record to be made active.`,
				},
				resource.Attribute{
					Name:        "response_method",
					Description: `(Required) (String) The method used to select which record is returned from the primary record pool. valid values are:</br> ` + "`" + `PRIORITY_HUNT` + "`" + ` – The sequence of the records listed in the primary record pool determines the priority. The first record listed is the highest priority record. Once a record starts being served, it will always be served until the probe detects a failure on this record, or, the record is set to FORCED_INACTIVE.</br> ` + "`" + `RANDOM` + "`" + ` – A random record is returned from the following set of primary records.</br> ` + "`" + `ROUND_ROBIN` + "`" + ` - A record is returned in (a round robin fashion) rotation, based upon the priority of the following active set of records.`,
				},
				resource.Attribute{
					Name:        "serving_preference",
					Description: `(Required) (String) It determines if records will be selected from the Primary Records pool , or, from the All Fail Record. Valid values are:</br> ` + "`" + `AUTO_SELECT` + "`" + `: Serving method switches from serving Primary Records, to All Fail Record based upon probe results, and the Forced State of the Primary Records.</br> ` + "`" + `SERVE_PRIMARY` + "`" + `: Only the Primary Records are served based upon the probe results and the Forced State of the Primary Records.</br> ` + "`" + `SERVE_ALL_FAIL` + "`" + `: Only the All Fail Record will be served, ignoring the probe results and the Forced State of the Primary Records.`,
				},
				resource.Attribute{
					Name:        "pool_description",
					Description: `(Computed) (String) An optional description of the Simple Load Balancing (SLB) field.`,
				},
				resource.Attribute{
					Name:        "monitor",
					Description: `(Computed) (Block Set) Nested block describing the information for the monitor. The structure of this block is described below.`,
				},
				resource.Attribute{
					Name:        "all_fail_record",
					Description: `(Required) (Block Set) Nested block describing the information for the backup record. The structure of this block is described below.`,
				},
				resource.Attribute{
					Name:        "rdata_info",
					Description: `(Required) (Block Set, Max: 5) Nested block describing the pool records. The structure of this block is described below.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Computed) (String) Current status of the serving record. Valid values are:</br> ` + "`" + `OK` + "`" + `- Priority record(s) are being served.</br> ` + "`" + `WARNING` + "`" + ` – One of the priority records is not being served due to the monitor detecting a failure, but there is still a priority record to be served.</br> ` + "`" + `CRITICAL` + "`" + ` – The backup All Fail record is being served due to the monitor detecting a failure. ### Nested ` + "`" + `monitor` + "`" + ` block has the following structure:`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Computed) (String) Monitored URL. A full URL including the protocol, host, and URI. Valid protocols are HTTP and HTTPS.`,
				},
				resource.Attribute{
					Name:        "method",
					Description: `(Computed) (String) HTTP method used to connect to the monitored URL.`,
				},
				resource.Attribute{
					Name:        "transmitted_data",
					Description: `(Computed) (String) If a monitor is sending a POST, the data that is sent as the body of the request.`,
				},
				resource.Attribute{
					Name:        "search_string",
					Description: `(Computed) (String) A string that is checked against the returned data from the request. ### Nested ` + "`" + `all_fail_record` + "`" + ` block has the following structure:`,
				},
				resource.Attribute{
					Name:        "rdata",
					Description: `(Computed) (String) An IPv4 address or IPv6 address as a backup record.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Computed) (String) An optional description for the backup record.`,
				},
				resource.Attribute{
					Name:        "serving",
					Description: `(Computed) (Boolean) Serving status of the AllFail Record. ### Nested ` + "`" + `rdata_info` + "`" + ` block has the following structure:`,
				},
				resource.Attribute{
					Name:        "rdata",
					Description: `(Computed) (String) An IPv4 or IPv6 address.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Computed) (String) An optional description of the record in the live pool.`,
				},
				resource.Attribute{
					Name:        "forced_state",
					Description: `(Computed) (String) The Forced State of the record that indicates whether the record needs to be: force served, forced to be inactive, or the force status not being considered (monitoring result decides the record state).`,
				},
				resource.Attribute{
					Name:        "probing_enabled",
					Description: `(Computed) (Boolean) Can be set at the record level to indicate whether probing is required (true) or not (false) for the given record.`,
				},
				resource.Attribute{
					Name:        "available_to_serve",
					Description: `(Computed) (Boolean) Indicates whether the record is available to be served (true) or not (false), based upon the probe results or the forced state of the record.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ttl",
					Description: `(Computed) (Integer) The time to live (in seconds) for the record. Must be a value between 0 and 2147483647, inclusive.`,
				},
				resource.Attribute{
					Name:        "region_failure_sensitivity",
					Description: `(Computed) (String) Specifies the sensitivity to the number of regions that need to fail for the backup record to be made active.`,
				},
				resource.Attribute{
					Name:        "response_method",
					Description: `(Required) (String) The method used to select which record is returned from the primary record pool. valid values are:</br> ` + "`" + `PRIORITY_HUNT` + "`" + ` – The sequence of the records listed in the primary record pool determines the priority. The first record listed is the highest priority record. Once a record starts being served, it will always be served until the probe detects a failure on this record, or, the record is set to FORCED_INACTIVE.</br> ` + "`" + `RANDOM` + "`" + ` – A random record is returned from the following set of primary records.</br> ` + "`" + `ROUND_ROBIN` + "`" + ` - A record is returned in (a round robin fashion) rotation, based upon the priority of the following active set of records.`,
				},
				resource.Attribute{
					Name:        "serving_preference",
					Description: `(Required) (String) It determines if records will be selected from the Primary Records pool , or, from the All Fail Record. Valid values are:</br> ` + "`" + `AUTO_SELECT` + "`" + `: Serving method switches from serving Primary Records, to All Fail Record based upon probe results, and the Forced State of the Primary Records.</br> ` + "`" + `SERVE_PRIMARY` + "`" + `: Only the Primary Records are served based upon the probe results and the Forced State of the Primary Records.</br> ` + "`" + `SERVE_ALL_FAIL` + "`" + `: Only the All Fail Record will be served, ignoring the probe results and the Forced State of the Primary Records.`,
				},
				resource.Attribute{
					Name:        "pool_description",
					Description: `(Computed) (String) An optional description of the Simple Load Balancing (SLB) field.`,
				},
				resource.Attribute{
					Name:        "monitor",
					Description: `(Computed) (Block Set) Nested block describing the information for the monitor. The structure of this block is described below.`,
				},
				resource.Attribute{
					Name:        "all_fail_record",
					Description: `(Required) (Block Set) Nested block describing the information for the backup record. The structure of this block is described below.`,
				},
				resource.Attribute{
					Name:        "rdata_info",
					Description: `(Required) (Block Set, Max: 5) Nested block describing the pool records. The structure of this block is described below.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Computed) (String) Current status of the serving record. Valid values are:</br> ` + "`" + `OK` + "`" + `- Priority record(s) are being served.</br> ` + "`" + `WARNING` + "`" + ` – One of the priority records is not being served due to the monitor detecting a failure, but there is still a priority record to be served.</br> ` + "`" + `CRITICAL` + "`" + ` – The backup All Fail record is being served due to the monitor detecting a failure. ### Nested ` + "`" + `monitor` + "`" + ` block has the following structure:`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Computed) (String) Monitored URL. A full URL including the protocol, host, and URI. Valid protocols are HTTP and HTTPS.`,
				},
				resource.Attribute{
					Name:        "method",
					Description: `(Computed) (String) HTTP method used to connect to the monitored URL.`,
				},
				resource.Attribute{
					Name:        "transmitted_data",
					Description: `(Computed) (String) If a monitor is sending a POST, the data that is sent as the body of the request.`,
				},
				resource.Attribute{
					Name:        "search_string",
					Description: `(Computed) (String) A string that is checked against the returned data from the request. ### Nested ` + "`" + `all_fail_record` + "`" + ` block has the following structure:`,
				},
				resource.Attribute{
					Name:        "rdata",
					Description: `(Computed) (String) An IPv4 address or IPv6 address as a backup record.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Computed) (String) An optional description for the backup record.`,
				},
				resource.Attribute{
					Name:        "serving",
					Description: `(Computed) (Boolean) Serving status of the AllFail Record. ### Nested ` + "`" + `rdata_info` + "`" + ` block has the following structure:`,
				},
				resource.Attribute{
					Name:        "rdata",
					Description: `(Computed) (String) An IPv4 or IPv6 address.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Computed) (String) An optional description of the record in the live pool.`,
				},
				resource.Attribute{
					Name:        "forced_state",
					Description: `(Computed) (String) The Forced State of the record that indicates whether the record needs to be: force served, forced to be inactive, or the force status not being considered (monitoring result decides the record state).`,
				},
				resource.Attribute{
					Name:        "probing_enabled",
					Description: `(Computed) (Boolean) Can be set at the record level to indicate whether probing is required (true) or not (false) for the given record.`,
				},
				resource.Attribute{
					Name:        "available_to_serve",
					Description: `(Computed) (Boolean) Indicates whether the record is available to be served (true) or not (false), based upon the probe results or the forced state of the record.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ultradns_tcpool",
			Category:         "TC-Pool",
			ShortDescription: `Get information of Traffic Controller (TC) pool records in UltraDNS.`,
			Description: `

Use this data source to get detailed information of Traffic Controller (TC) pool records.

`,
			Keywords: []string{
				"tc",
				"pool",
				"tcpool",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "zone_name",
					Description: `(Required) (String) Name of the zone.`,
				},
				resource.Attribute{
					Name:        "owner_name",
					Description: `(Required) (String) The domain name of the owner of the RRSet. Can be either a fully qualified domain name (FQDN) or a relative domain name. If provided as a FQDN, it must be contained within the zone name's FQDN.`,
				},
				resource.Attribute{
					Name:        "record_type",
					Description: `(Required) (String) Must be formatted as a well-known resource record type (A), or the corresponding number for the type (1).<br/> Below are the supported resource record types with the corresponding number:<br/> ` + "`" + `A (1)` + "`" + ` ## Attributes Reference In addition to all of the arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `(Computed) (Integer) The time to live (in seconds) for the record. Must be a value between 0 and 2147483647, inclusive.`,
				},
				resource.Attribute{
					Name:        "pool_description",
					Description: `(Computed) (String) An optional description of the Traffic Controller (TC) field.`,
				},
				resource.Attribute{
					Name:        "run_probes",
					Description: `(Computed) (Boolean) Indicates whether or not the probes are run for this pool.`,
				},
				resource.Attribute{
					Name:        "act_on_probes",
					Description: `(Computed) (Boolean) Indicates whether or not pool records will be enabled (true) or disabled (false) when probes are run.`,
				},
				resource.Attribute{
					Name:        "failure_threshold",
					Description: `(Computed) (Integer) The minimum number of records that must fail for a pool to be labeled 'FAILED'. If the number of failed records in the pool is greater than or equal to the 'Failure Threshold' value, the pool will be labeled FAILED.<br/> For example, a pool with six priority records, one all-fail record, and the Failure Threshold value is set to four (4). If four or more priority records are not available to serve, the pool will be labeled FAILED, and the all-fail record will be served.`,
				},
				resource.Attribute{
					Name:        "max_to_lb",
					Description: `(Computed) (Integer) Specifies the maximum number of active servers in a pool. The maximum value is the number of pool records.`,
				},
				resource.Attribute{
					Name:        "backup_record",
					Description: `(Computed) (Block Set, Max: 1) Nested block describing the information of the backup record for the Traffic Controller pool. The backup record is served if all other records fail. There can be one or more A records used as backup records, or a single CNAME record. The structure of this block is described below.`,
				},
				resource.Attribute{
					Name:        "rdata_info",
					Description: `(Computed) (Block Set) List of nested blocks describing the pool records. The structure of this block is described below.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Computed) (String) Current status of the serving record. Valid values are:</br> ` + "`" + `OK` + "`" + `- If the number of records serving is equal to the Max Active value, and all the active records are top priority records.</br> For example, if a pool has a Max Active of 1 and the Priority 1 record is serving.</br> ` + "`" + `WARNING` + "`" + ` – If the number of records serving is equal to the Max Active value, and the active records are not top priority records.</br> For example, if a pool has a Max Active of 1 and the Priority 1 records is not serving and the Priority 2 record is serving.</br> ` + "`" + `CRITICAL` + "`" + ` – If the number of records serving is less than the Max Active value, or the All Fail record is being served.</br> For example, if a pool has a Max Active of 2, and only one record is serving.</br> ` + "`" + `FAILED` + "`" + ` - If the FailureThreshold value is 0 or null, and no records are serving, and there is no All Fail record configured.</br>OR</br>If the number of priority records not available to serve equals or exceeds the FailureThreshold’s value.</br> For example, if the Failure Threshold value is 3, and there are 3 or more Priority Records that are not available to serve. ### Nested ` + "`" + `backup_record` + "`" + ` block has the following structure:`,
				},
				resource.Attribute{
					Name:        "rdata",
					Description: `(Computed) (String) The IPv4 address or CNAME for the backup record.`,
				},
				resource.Attribute{
					Name:        "failover_delay",
					Description: `(Computed) (Integer) Specifies the time, between 0 – 30 minutes, that the Traffic Controller waits after detecting that the pool record has failed, prior to activating the primary records. Default value set to 0.`,
				},
				resource.Attribute{
					Name:        "available_to_serve",
					Description: `(Computed) (Boolean) Indicates whether the pool's backup record is active and available to serve records. ### Nested ` + "`" + `rdata_info` + "`" + ` block has the following structure:`,
				},
				resource.Attribute{
					Name:        "rdata",
					Description: `(Computed) (String) The IPv4 address or CNAME.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `(Computed) (Integer) Indicates the serving preference for this pool record.`,
				},
				resource.Attribute{
					Name:        "threshold",
					Description: `(Computed) (Integer) Specifies how many probes must agree before the record state is changed.`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `(Optional) (Integer) Determines the traffic load to send to each server in the Traffic Controller pool. Even integers from 2 to 100.`,
				},
				resource.Attribute{
					Name:        "failover_delay",
					Description: `(Computed) (Integer) Specifies the time, between 0 – 30 minutes, that the Traffic Controller waits after detecting that the pool record has failed, prior to activating the secondary records. Default value set to 0.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `(Computed) (String) The current state of the pool record.`,
				},
				resource.Attribute{
					Name:        "run_probes",
					Description: `(Computed) (Boolean) Indicates whether or not probes are run for this pool record.`,
				},
				resource.Attribute{
					Name:        "available_to_serve",
					Description: `(Computed) (Boolean) Indicates whether the pool record is active and available to serve records.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ttl",
					Description: `(Computed) (Integer) The time to live (in seconds) for the record. Must be a value between 0 and 2147483647, inclusive.`,
				},
				resource.Attribute{
					Name:        "pool_description",
					Description: `(Computed) (String) An optional description of the Traffic Controller (TC) field.`,
				},
				resource.Attribute{
					Name:        "run_probes",
					Description: `(Computed) (Boolean) Indicates whether or not the probes are run for this pool.`,
				},
				resource.Attribute{
					Name:        "act_on_probes",
					Description: `(Computed) (Boolean) Indicates whether or not pool records will be enabled (true) or disabled (false) when probes are run.`,
				},
				resource.Attribute{
					Name:        "failure_threshold",
					Description: `(Computed) (Integer) The minimum number of records that must fail for a pool to be labeled 'FAILED'. If the number of failed records in the pool is greater than or equal to the 'Failure Threshold' value, the pool will be labeled FAILED.<br/> For example, a pool with six priority records, one all-fail record, and the Failure Threshold value is set to four (4). If four or more priority records are not available to serve, the pool will be labeled FAILED, and the all-fail record will be served.`,
				},
				resource.Attribute{
					Name:        "max_to_lb",
					Description: `(Computed) (Integer) Specifies the maximum number of active servers in a pool. The maximum value is the number of pool records.`,
				},
				resource.Attribute{
					Name:        "backup_record",
					Description: `(Computed) (Block Set, Max: 1) Nested block describing the information of the backup record for the Traffic Controller pool. The backup record is served if all other records fail. There can be one or more A records used as backup records, or a single CNAME record. The structure of this block is described below.`,
				},
				resource.Attribute{
					Name:        "rdata_info",
					Description: `(Computed) (Block Set) List of nested blocks describing the pool records. The structure of this block is described below.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Computed) (String) Current status of the serving record. Valid values are:</br> ` + "`" + `OK` + "`" + `- If the number of records serving is equal to the Max Active value, and all the active records are top priority records.</br> For example, if a pool has a Max Active of 1 and the Priority 1 record is serving.</br> ` + "`" + `WARNING` + "`" + ` – If the number of records serving is equal to the Max Active value, and the active records are not top priority records.</br> For example, if a pool has a Max Active of 1 and the Priority 1 records is not serving and the Priority 2 record is serving.</br> ` + "`" + `CRITICAL` + "`" + ` – If the number of records serving is less than the Max Active value, or the All Fail record is being served.</br> For example, if a pool has a Max Active of 2, and only one record is serving.</br> ` + "`" + `FAILED` + "`" + ` - If the FailureThreshold value is 0 or null, and no records are serving, and there is no All Fail record configured.</br>OR</br>If the number of priority records not available to serve equals or exceeds the FailureThreshold’s value.</br> For example, if the Failure Threshold value is 3, and there are 3 or more Priority Records that are not available to serve. ### Nested ` + "`" + `backup_record` + "`" + ` block has the following structure:`,
				},
				resource.Attribute{
					Name:        "rdata",
					Description: `(Computed) (String) The IPv4 address or CNAME for the backup record.`,
				},
				resource.Attribute{
					Name:        "failover_delay",
					Description: `(Computed) (Integer) Specifies the time, between 0 – 30 minutes, that the Traffic Controller waits after detecting that the pool record has failed, prior to activating the primary records. Default value set to 0.`,
				},
				resource.Attribute{
					Name:        "available_to_serve",
					Description: `(Computed) (Boolean) Indicates whether the pool's backup record is active and available to serve records. ### Nested ` + "`" + `rdata_info` + "`" + ` block has the following structure:`,
				},
				resource.Attribute{
					Name:        "rdata",
					Description: `(Computed) (String) The IPv4 address or CNAME.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `(Computed) (Integer) Indicates the serving preference for this pool record.`,
				},
				resource.Attribute{
					Name:        "threshold",
					Description: `(Computed) (Integer) Specifies how many probes must agree before the record state is changed.`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `(Optional) (Integer) Determines the traffic load to send to each server in the Traffic Controller pool. Even integers from 2 to 100.`,
				},
				resource.Attribute{
					Name:        "failover_delay",
					Description: `(Computed) (Integer) Specifies the time, between 0 – 30 minutes, that the Traffic Controller waits after detecting that the pool record has failed, prior to activating the secondary records. Default value set to 0.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `(Computed) (String) The current state of the pool record.`,
				},
				resource.Attribute{
					Name:        "run_probes",
					Description: `(Computed) (Boolean) Indicates whether or not probes are run for this pool record.`,
				},
				resource.Attribute{
					Name:        "available_to_serve",
					Description: `(Computed) (Boolean) Indicates whether the pool record is active and available to serve records.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ultradns_zone",
			Category:         "Zone",
			ShortDescription: `Get information of zones in UltraDNS.`,
			Description: `

Use this data source to get detailed information for your zones.

`,
			Keywords: []string{
				"zone",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) (String) Name of the zone. ## Attributes Reference In addition to the arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "account_name",
					Description: `(Computed) (String) Name of the account.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Computed) (String) Type of zone. Valid values are ` + "`" + `PRIMARY` + "`" + `, ` + "`" + `SECONDARY` + "`" + ` or ` + "`" + `ALIAS` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Computed) (String) Displays the status of the zone. Valid values are ` + "`" + `ACTIVE` + "`" + `, ` + "`" + `SUSPENDED` + "`" + ` or ` + "`" + `ERROR` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "dnssec_status",
					Description: `(Computed) (String) Whether or not the zone is signed with DNSSEC. Valid values are ` + "`" + `SIGNED` + "`" + ` or ` + "`" + `UNSIGNED` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "owner",
					Description: `(Computed) (String) Name of the user that created the zone.`,
				},
				resource.Attribute{
					Name:        "resource_record_count",
					Description: `(Computed) (Integer) Number of records in the zone.`,
				},
				resource.Attribute{
					Name:        "last_modified_time",
					Description: `(Computed) (String) The last date and time the zone was modified, represented in ISO8601 format(` + "`" + `yyyy-MM-ddTHH:mmZ` + "`" + `).<br/> Example: ` + "`" + `2021-12-07T11:25Z` + "`" + `. #### When ` + "`" + `type` + "`" + ` is "PRIMARY" the below attributes will be exported.`,
				},
				resource.Attribute{
					Name:        "inherit",
					Description: `(Computed) (String) Describes the inherited zone transfer values from the Account. Valid values are ` + "`" + `ALL` + "`" + `, ` + "`" + `NONE` + "`" + `, any combination of ` + "`" + `IP_RANGE` + "`" + `, ` + "`" + `NOTIFY_IP` + "`" + `, ` + "`" + `TSIG` + "`" + `. Multiple values are separated with a comma.<br/> Example: ` + "`" + `IP_RANGE, NOTIFY_IP` + "`" + ``,
				},
				resource.Attribute{
					Name:        "tsig",
					Description: `(Computed) (Block Set, Max: 1) Nested block describing the TSIG information for the primary zone. The structure of this block is described below.`,
				},
				resource.Attribute{
					Name:        "restrict_ip",
					Description: `(Computed) (Block Set) Nested block describing the list of IPv4 or IPv6 ranges that are allowed to transfer primary zones out using zone transfer protocol (AXFR/IXFR). The structure of this block is described below.`,
				},
				resource.Attribute{
					Name:        "notify_addresses",
					Description: `(Computed) (Block Set) Nested block describing the IPv4 Addresses that are notified when updates are made to the primary zone. The structure of this block is described below.`,
				},
				resource.Attribute{
					Name:        "registrar_info",
					Description: `(Computed) (Block Set) Nested block describing information about the name server configuration for this zone. The structure of this block is described below. #### When ` + "`" + `type` + "`" + ` is "SECONDARY" the below attributes will be exported.`,
				},
				resource.Attribute{
					Name:        "primary_name_server_1",
					Description: `(Computed) (Block Set) The structure of this block follows the same structure as the [` + "`" + `name_server` + "`" + `](#nested-name_server-block-has-the-following-structure) block described below. It is the info of primary name server.`,
				},
				resource.Attribute{
					Name:        "primary_name_server_2",
					Description: `(Computed) (Block Set) The structure of this block follows the same structure as the [` + "`" + `name_server` + "`" + `](#nested-name_server-block-has-the-following-structure) block described below. It is the info of first backup primary name server.`,
				},
				resource.Attribute{
					Name:        "primary_name_server_3",
					Description: `(Computed) (Block Set) The structure of this block follows the same structure as the [` + "`" + `name_server` + "`" + `](#nested-name_server-block-has-the-following-structure) block described below. It is the info of second backup primary name server.`,
				},
				resource.Attribute{
					Name:        "notification_email_address",
					Description: `(Computed) (String) The Notification Email for a secondary zone.`,
				},
				resource.Attribute{
					Name:        "transfer_status_details",
					Description: `(Computed) (Block Set) Nested block describing the zone transfer details. The structure of this block is described below. #### When ` + "`" + `type` + "`" + ` is "ALIAS" the below attributes will be exported.`,
				},
				resource.Attribute{
					Name:        "original_zone_name",
					Description: `(Computed) (String) The name of the zone being aliased. The existing zone must be owned by the same account as the new zone. ### Nested ` + "`" + `name_server` + "`" + ` block has the following structure:`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `(Required) (String) The IPv4 or IPv6 address of the primary name server for the source zone.`,
				},
				resource.Attribute{
					Name:        "tsig_key",
					Description: `(Optional) (String) If TSIG is enabled for this name server, the name of the TSIG key.`,
				},
				resource.Attribute{
					Name:        "tsig_key_value",
					Description: `(Optional) (String) If TSIG is enabled for this name server, the TSIG key’s value.`,
				},
				resource.Attribute{
					Name:        "tsig_algorithm",
					Description: `(Optional) (String) The hash algorithm used to generate the TSIG key. Valid values are ` + "`" + `hmac-md5` + "`" + `, ` + "`" + `hmac-sha1` + "`" + `, ` + "`" + `hmac-sha224` + "`" + `, ` + "`" + `hmac-sha256` + "`" + `, ` + "`" + `hmac-sha384` + "`" + `, ` + "`" + `hmac-sha512` + "`" + `. ### Nested ` + "`" + `tsig` + "`" + ` block has the following structure:`,
				},
				resource.Attribute{
					Name:        "tsig_key_name",
					Description: `(Required) (String) The name of the TSIG key for the zone.`,
				},
				resource.Attribute{
					Name:        "tsig_key_value",
					Description: `(Required) (String) The value of the TSIG key for the zone.`,
				},
				resource.Attribute{
					Name:        "tsig_algorithm",
					Description: `(Required) (String) The hash algorithm used to generate the TSIG key. Valid values are ` + "`" + `hmac-md5` + "`" + `, ` + "`" + `hmac-sha1` + "`" + `, ` + "`" + `hmac-sha224` + "`" + `, ` + "`" + `hmac-sha256` + "`" + `, ` + "`" + `hmac-sha384` + "`" + `, ` + "`" + `hmac-sha512` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) (String) A description of this key. ### Nested ` + "`" + `restrict_ip` + "`" + ` block has the following structure:`,
				},
				resource.Attribute{
					Name:        "start_ip",
					Description: `(Optional) (String) The start of the IPv4 or IPv6 range that is allowed to transfer this primary zone out using zone transfer protocol.`,
				},
				resource.Attribute{
					Name:        "end_ip",
					Description: `(Optional) (String) The end of the IPv4 or IPv6 range that is allowed to transfer this primary zone out using zone transfer protocol.`,
				},
				resource.Attribute{
					Name:        "cidr",
					Description: `(Optional) (String) The IP Address ranges specified in CIDR.`,
				},
				resource.Attribute{
					Name:        "single_ip",
					Description: `(Optional) (String) The IPv4 or IPv6 address that is allowed to transfer this primary zone out using zone transfer protocol.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `(Optional) (String) A description of this range of IP addresses. ### Nested ` + "`" + `notify_addresses` + "`" + ` block has the following structure:`,
				},
				resource.Attribute{
					Name:        "notify_address",
					Description: `(Required) (String) The IPv4 Address that is notified when the primary zone is updated.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) (String) A description of this address. ### Nested ` + "`" + `registrar_info` + "`" + ` block has the following structure:`,
				},
				resource.Attribute{
					Name:        "registrar",
					Description: `(Computed) (String) The name of the domain registrar.`,
				},
				resource.Attribute{
					Name:        "who_is_expiration",
					Description: `(Computed) (String) The date (` + "`" + `yyyy-MM-dd HH:mm:ss.S` + "`" + `) when the domain name registration expires.<br/> Example: ` + "`" + `2022-08-17 03:59:59.0` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "name_servers",
					Description: `(Computed) (Block Set) Nested block describing the name servers configuration of the zone. The structure of this block is described below. ### Nested ` + "`" + `registrar_info.name_servers` + "`" + ` block has the following structure:`,
				},
				resource.Attribute{
					Name:        "ok",
					Description: `(Computed) (List String) List of UltraDNS name servers that are configured for this domain.`,
				},
				resource.Attribute{
					Name:        "unknown",
					Description: `(Computed) (List String) List of name servers that are configured for this domain, but are not UltraDNS managed name servers.`,
				},
				resource.Attribute{
					Name:        "missing",
					Description: `(Computed) (List String) List of UltraDNS name servers that should be configured for this domain, but are not.`,
				},
				resource.Attribute{
					Name:        "incorrect",
					Description: `(Computed) (List String) List of any obsolete UltraDNS name servers that are still configured for this zone. ### Nested ` + "`" + `transfer_status_details` + "`" + ` block has the following structure:`,
				},
				resource.Attribute{
					Name:        "last_refresh",
					Description: `(Computed) (String) Displays the date (` + "`" + `MM/dd/yy HH:mm:ss tt vvv` + "`" + `) when the last transfer attempt or refresh was.<br/> Example: ` + "`" + `03/18/21 10:20:34 AM GMT` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "next_refresh",
					Description: `(Computed) (String) Displays the date (` + "`" + `MM/dd/yy HH:mm:ss tt vvv` + "`" + `) when the next transfer attempt or refresh is.<br/> Example: ` + "`" + `03/18/21 10:20:34 AM GMT` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "last_refresh_status",
					Description: `(Computed) (String) Displays the status of the last transfer that was attempted. Valid values are ` + "`" + `IN_PROGRESS` + "`" + `, ` + "`" + `FAILED` + "`" + `, ` + "`" + `SUCCESSFUL` + "`" + ``,
				},
				resource.Attribute{
					Name:        "last_refresh_status_message",
					Description: `(Computed) (String) Displays the last transfer’s status message. This is currently shown as failure reason.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "account_name",
					Description: `(Computed) (String) Name of the account.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Computed) (String) Type of zone. Valid values are ` + "`" + `PRIMARY` + "`" + `, ` + "`" + `SECONDARY` + "`" + ` or ` + "`" + `ALIAS` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Computed) (String) Displays the status of the zone. Valid values are ` + "`" + `ACTIVE` + "`" + `, ` + "`" + `SUSPENDED` + "`" + ` or ` + "`" + `ERROR` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "dnssec_status",
					Description: `(Computed) (String) Whether or not the zone is signed with DNSSEC. Valid values are ` + "`" + `SIGNED` + "`" + ` or ` + "`" + `UNSIGNED` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "owner",
					Description: `(Computed) (String) Name of the user that created the zone.`,
				},
				resource.Attribute{
					Name:        "resource_record_count",
					Description: `(Computed) (Integer) Number of records in the zone.`,
				},
				resource.Attribute{
					Name:        "last_modified_time",
					Description: `(Computed) (String) The last date and time the zone was modified, represented in ISO8601 format(` + "`" + `yyyy-MM-ddTHH:mmZ` + "`" + `).<br/> Example: ` + "`" + `2021-12-07T11:25Z` + "`" + `. #### When ` + "`" + `type` + "`" + ` is "PRIMARY" the below attributes will be exported.`,
				},
				resource.Attribute{
					Name:        "inherit",
					Description: `(Computed) (String) Describes the inherited zone transfer values from the Account. Valid values are ` + "`" + `ALL` + "`" + `, ` + "`" + `NONE` + "`" + `, any combination of ` + "`" + `IP_RANGE` + "`" + `, ` + "`" + `NOTIFY_IP` + "`" + `, ` + "`" + `TSIG` + "`" + `. Multiple values are separated with a comma.<br/> Example: ` + "`" + `IP_RANGE, NOTIFY_IP` + "`" + ``,
				},
				resource.Attribute{
					Name:        "tsig",
					Description: `(Computed) (Block Set, Max: 1) Nested block describing the TSIG information for the primary zone. The structure of this block is described below.`,
				},
				resource.Attribute{
					Name:        "restrict_ip",
					Description: `(Computed) (Block Set) Nested block describing the list of IPv4 or IPv6 ranges that are allowed to transfer primary zones out using zone transfer protocol (AXFR/IXFR). The structure of this block is described below.`,
				},
				resource.Attribute{
					Name:        "notify_addresses",
					Description: `(Computed) (Block Set) Nested block describing the IPv4 Addresses that are notified when updates are made to the primary zone. The structure of this block is described below.`,
				},
				resource.Attribute{
					Name:        "registrar_info",
					Description: `(Computed) (Block Set) Nested block describing information about the name server configuration for this zone. The structure of this block is described below. #### When ` + "`" + `type` + "`" + ` is "SECONDARY" the below attributes will be exported.`,
				},
				resource.Attribute{
					Name:        "primary_name_server_1",
					Description: `(Computed) (Block Set) The structure of this block follows the same structure as the [` + "`" + `name_server` + "`" + `](#nested-name_server-block-has-the-following-structure) block described below. It is the info of primary name server.`,
				},
				resource.Attribute{
					Name:        "primary_name_server_2",
					Description: `(Computed) (Block Set) The structure of this block follows the same structure as the [` + "`" + `name_server` + "`" + `](#nested-name_server-block-has-the-following-structure) block described below. It is the info of first backup primary name server.`,
				},
				resource.Attribute{
					Name:        "primary_name_server_3",
					Description: `(Computed) (Block Set) The structure of this block follows the same structure as the [` + "`" + `name_server` + "`" + `](#nested-name_server-block-has-the-following-structure) block described below. It is the info of second backup primary name server.`,
				},
				resource.Attribute{
					Name:        "notification_email_address",
					Description: `(Computed) (String) The Notification Email for a secondary zone.`,
				},
				resource.Attribute{
					Name:        "transfer_status_details",
					Description: `(Computed) (Block Set) Nested block describing the zone transfer details. The structure of this block is described below. #### When ` + "`" + `type` + "`" + ` is "ALIAS" the below attributes will be exported.`,
				},
				resource.Attribute{
					Name:        "original_zone_name",
					Description: `(Computed) (String) The name of the zone being aliased. The existing zone must be owned by the same account as the new zone. ### Nested ` + "`" + `name_server` + "`" + ` block has the following structure:`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `(Required) (String) The IPv4 or IPv6 address of the primary name server for the source zone.`,
				},
				resource.Attribute{
					Name:        "tsig_key",
					Description: `(Optional) (String) If TSIG is enabled for this name server, the name of the TSIG key.`,
				},
				resource.Attribute{
					Name:        "tsig_key_value",
					Description: `(Optional) (String) If TSIG is enabled for this name server, the TSIG key’s value.`,
				},
				resource.Attribute{
					Name:        "tsig_algorithm",
					Description: `(Optional) (String) The hash algorithm used to generate the TSIG key. Valid values are ` + "`" + `hmac-md5` + "`" + `, ` + "`" + `hmac-sha1` + "`" + `, ` + "`" + `hmac-sha224` + "`" + `, ` + "`" + `hmac-sha256` + "`" + `, ` + "`" + `hmac-sha384` + "`" + `, ` + "`" + `hmac-sha512` + "`" + `. ### Nested ` + "`" + `tsig` + "`" + ` block has the following structure:`,
				},
				resource.Attribute{
					Name:        "tsig_key_name",
					Description: `(Required) (String) The name of the TSIG key for the zone.`,
				},
				resource.Attribute{
					Name:        "tsig_key_value",
					Description: `(Required) (String) The value of the TSIG key for the zone.`,
				},
				resource.Attribute{
					Name:        "tsig_algorithm",
					Description: `(Required) (String) The hash algorithm used to generate the TSIG key. Valid values are ` + "`" + `hmac-md5` + "`" + `, ` + "`" + `hmac-sha1` + "`" + `, ` + "`" + `hmac-sha224` + "`" + `, ` + "`" + `hmac-sha256` + "`" + `, ` + "`" + `hmac-sha384` + "`" + `, ` + "`" + `hmac-sha512` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) (String) A description of this key. ### Nested ` + "`" + `restrict_ip` + "`" + ` block has the following structure:`,
				},
				resource.Attribute{
					Name:        "start_ip",
					Description: `(Optional) (String) The start of the IPv4 or IPv6 range that is allowed to transfer this primary zone out using zone transfer protocol.`,
				},
				resource.Attribute{
					Name:        "end_ip",
					Description: `(Optional) (String) The end of the IPv4 or IPv6 range that is allowed to transfer this primary zone out using zone transfer protocol.`,
				},
				resource.Attribute{
					Name:        "cidr",
					Description: `(Optional) (String) The IP Address ranges specified in CIDR.`,
				},
				resource.Attribute{
					Name:        "single_ip",
					Description: `(Optional) (String) The IPv4 or IPv6 address that is allowed to transfer this primary zone out using zone transfer protocol.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `(Optional) (String) A description of this range of IP addresses. ### Nested ` + "`" + `notify_addresses` + "`" + ` block has the following structure:`,
				},
				resource.Attribute{
					Name:        "notify_address",
					Description: `(Required) (String) The IPv4 Address that is notified when the primary zone is updated.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) (String) A description of this address. ### Nested ` + "`" + `registrar_info` + "`" + ` block has the following structure:`,
				},
				resource.Attribute{
					Name:        "registrar",
					Description: `(Computed) (String) The name of the domain registrar.`,
				},
				resource.Attribute{
					Name:        "who_is_expiration",
					Description: `(Computed) (String) The date (` + "`" + `yyyy-MM-dd HH:mm:ss.S` + "`" + `) when the domain name registration expires.<br/> Example: ` + "`" + `2022-08-17 03:59:59.0` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "name_servers",
					Description: `(Computed) (Block Set) Nested block describing the name servers configuration of the zone. The structure of this block is described below. ### Nested ` + "`" + `registrar_info.name_servers` + "`" + ` block has the following structure:`,
				},
				resource.Attribute{
					Name:        "ok",
					Description: `(Computed) (List String) List of UltraDNS name servers that are configured for this domain.`,
				},
				resource.Attribute{
					Name:        "unknown",
					Description: `(Computed) (List String) List of name servers that are configured for this domain, but are not UltraDNS managed name servers.`,
				},
				resource.Attribute{
					Name:        "missing",
					Description: `(Computed) (List String) List of UltraDNS name servers that should be configured for this domain, but are not.`,
				},
				resource.Attribute{
					Name:        "incorrect",
					Description: `(Computed) (List String) List of any obsolete UltraDNS name servers that are still configured for this zone. ### Nested ` + "`" + `transfer_status_details` + "`" + ` block has the following structure:`,
				},
				resource.Attribute{
					Name:        "last_refresh",
					Description: `(Computed) (String) Displays the date (` + "`" + `MM/dd/yy HH:mm:ss tt vvv` + "`" + `) when the last transfer attempt or refresh was.<br/> Example: ` + "`" + `03/18/21 10:20:34 AM GMT` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "next_refresh",
					Description: `(Computed) (String) Displays the date (` + "`" + `MM/dd/yy HH:mm:ss tt vvv` + "`" + `) when the next transfer attempt or refresh is.<br/> Example: ` + "`" + `03/18/21 10:20:34 AM GMT` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "last_refresh_status",
					Description: `(Computed) (String) Displays the status of the last transfer that was attempted. Valid values are ` + "`" + `IN_PROGRESS` + "`" + `, ` + "`" + `FAILED` + "`" + `, ` + "`" + `SUCCESSFUL` + "`" + ``,
				},
				resource.Attribute{
					Name:        "last_refresh_status_message",
					Description: `(Computed) (String) Displays the last transfer’s status message. This is currently shown as failure reason.`,
				},
			},
		},
	}

	dataSourcesMap = map[string]int{

		"ultradns_dirpool":   0,
		"ultradns_probedns":  1,
		"ultradns_probehttp": 2,
		"ultradns_probeping": 3,
		"ultradns_rdpool":    4,
		"ultradns_record":    5,
		"ultradns_sbpool":    6,
		"ultradns_sfpool":    7,
		"ultradns_slbpool":   8,
		"ultradns_tcpool":    9,
		"ultradns_zone":      10,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
