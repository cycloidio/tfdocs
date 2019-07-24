package ultradns

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "ultradns_dirpool",
			Category:         "Resources",
			ShortDescription: `Provides an UltraDNS Directional Controller pool resource.`,
			Description:      ``,
			Keywords: []string{
				"dirpool",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "zone",
					Description: `(Required) The domain to add the record to`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the record - ` + "`" + `type` + "`" + ` - (Required) The Record Type of the record`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Required) Description of the Traffic Controller pool. Valid values are strings less than 256 characters.`,
				},
				resource.Attribute{
					Name:        "rdata",
					Description: `(Required) a list of Record Data blocks, one for each member in the pool. Record Data documented below.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `(Optional) The TTL of the record. Default: ` + "`" + `3600` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "conflict_resolve",
					Description: `(Optional) String. Valid: ` + "`" + `"GEO"` + "`" + ` or ` + "`" + `"IP"` + "`" + `. Default: ` + "`" + `"GEO"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "no_response",
					Description: `(Optional) a single Record Data block, without any ` + "`" + `host` + "`" + ` attribute. Record Data documented below. Record Data blocks support the following:`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `(Required in ` + "`" + `rdata` + "`" + `, absent in ` + "`" + `no_response` + "`" + `) IPv4 address or CNAME for the pool member. - ` + "`" + `all_non_configured` + "`" + ` - (Optional) Boolean. Default: ` + "`" + `false` + "`" + `. - ` + "`" + `geo_info` + "`" + ` - (Optional) a single Geo Info block. Geo Info documented below. - ` + "`" + `ip_info` + "`" + ` - (Optional) a single IP Info block. IP Info documented below. Geo Info blocks support the following: - ` + "`" + `name` + "`" + ` - (Optional) String. - ` + "`" + `is_account_level` + "`" + ` - (Optional) Boolean. Default: ` + "`" + `false` + "`" + `. - ` + "`" + `codes` + "`" + ` - (Optional) Set of geo code strings. Shorthand codes are expanded. IP Info blocks support the following: - ` + "`" + `name` + "`" + ` - (Optional) String. - ` + "`" + `is_account_level` + "`" + ` - (Optional) Boolean. Default: ` + "`" + `false` + "`" + `. - ` + "`" + `ips` + "`" + ` - (Optional) Set of IP blocks. IP Info documented below. IP blocks support the following: - ` + "`" + `start` + "`" + ` - (Optional) String. IP Address. Must be paired with ` + "`" + `end` + "`" + `. Conflicts with ` + "`" + `cidr` + "`" + ` or ` + "`" + `address` + "`" + `. - ` + "`" + `end` + "`" + ` - (Optional) String. IP Address. Must be paired with ` + "`" + `start` + "`" + `. - ` + "`" + `cidr` + "`" + ` - (Optional) String. - ` + "`" + `address` + "`" + ` - (Optional) String. IP Address. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The record ID`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `The FQDN of the record`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The record ID`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `The FQDN of the record`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ultradns_probe_http",
			Category:         "Resources",
			ShortDescription: `Provides an UltraDNS HTTP probe`,
			Description:      ``,
			Keywords: []string{
				"probe",
				"http",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "zone",
					Description: `(Required) The domain of the pool to probe.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the pool to probe. - ` + "`" + `pool_record` + "`" + ` - (optional) IP address or domain. If provided, a record-level probe is created, otherwise a pool-level probe is created. - ` + "`" + `agents` + "`" + ` - (Required) List of locations that will be used for probing. One or more values must be specified. Valid values are ` + "`" + `"NEW_YORK"` + "`" + `, ` + "`" + `"PALO_ALTO"` + "`" + `, ` + "`" + `"DALLAS"` + "`" + ` & ` + "`" + `"AMSTERDAM"` + "`" + `. - ` + "`" + `threshold` + "`" + ` - (Required) Number of agents that must agree for a probe state to be changed. - ` + "`" + `http_probe` + "`" + ` - (Required) an HTTP Probe block. - ` + "`" + `interval` + "`" + ` - (Optional) Length of time between probes in minutes. Valid values are ` + "`" + `"HALF_MINUTE"` + "`" + `, ` + "`" + `"ONE_MINUTE"` + "`" + `, ` + "`" + `"TWO_MINUTES"` + "`" + `, ` + "`" + `"FIVE_MINUTES"` + "`" + `, ` + "`" + `"TEN_MINUTES"` + "`" + ` & ` + "`" + `"FIFTEEN_MINUTE"` + "`" + `. Default: ` + "`" + `"FIVE_MINUTES"` + "`" + `. HTTP Probe block - ` + "`" + `transaction` + "`" + ` - (Optional) One or more Transaction blocks. - ` + "`" + `total_limits` + "`" + ` - (Optional) A Limit block, but with no ` + "`" + `name` + "`" + ` attribute. Transaction block - ` + "`" + `method` + "`" + ` - (Required) HTTP method. Valid values are` + "`" + `"GET"` + "`" + `, ` + "`" + `"POST"` + "`" + `. - ` + "`" + `url` + "`" + ` - (Required) URL to probe. - ` + "`" + `transmitted_data` + "`" + ` - (Optional) Data to send to URL. - ` + "`" + `follow_redirects` + "`" + ` - (Optional) Whether to follow redirects. - ` + "`" + `limit` + "`" + ` - (Required) One or more Limit blocks. Only one limit block may exist for each name. Limit block - ` + "`" + `name` + "`" + ` - (Required) Kind of limit. Valid values are ` + "`" + `"lossPercent"` + "`" + `, ` + "`" + `"total"` + "`" + `, ` + "`" + `"average"` + "`" + `, ` + "`" + `"run"` + "`" + ` & ` + "`" + `"avgRun"` + "`" + `. - ` + "`" + `warning` + "`" + ` - (Optional) Amount to trigger a warning. - ` + "`" + `critical` + "`" + ` - (Optional) Amount to trigger a critical. - ` + "`" + `fail` + "`" + ` - (Optional) Amount to trigger a failure.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ultradns_probe_ping",
			Category:         "Resources",
			ShortDescription: `Provides an UltraDNS Ping Probe`,
			Description:      ``,
			Keywords: []string{
				"probe",
				"ping",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "zone",
					Description: `(Required) The domain of the pool to probe.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the pool to probe. - ` + "`" + `pool_record` + "`" + ` - (optional) IP address or domain. If provided, a record-level probe is created, otherwise a pool-level probe is created. - ` + "`" + `agents` + "`" + ` - (Required) List of locations that will be used for probing. One or more values must be specified. Valid values are ` + "`" + `"NEW_YORK"` + "`" + `, ` + "`" + `"PALO_ALTO"` + "`" + `, ` + "`" + `"DALLAS"` + "`" + ` & ` + "`" + `"AMSTERDAM"` + "`" + `. - ` + "`" + `threshold` + "`" + ` - (Required) Number of agents that must agree for a probe state to be changed. - ` + "`" + `ping_probe` + "`" + ` - (Required) a Ping Probe block. - ` + "`" + `interval` + "`" + ` - (Optional) Length of time between probes in minutes. Valid values are ` + "`" + `"HALF_MINUTE"` + "`" + `, ` + "`" + `"ONE_MINUTE"` + "`" + `, ` + "`" + `"TWO_MINUTES"` + "`" + `, ` + "`" + `"FIVE_MINUTES"` + "`" + `, ` + "`" + `"TEN_MINUTES"` + "`" + ` & ` + "`" + `"FIFTEEN_MINUTE"` + "`" + `. Default: ` + "`" + `"FIVE_MINUTES"` + "`" + `. Ping Probe block - ` + "`" + `packets` + "`" + ` - (Optional) Number of ICMP packets to send. Default ` + "`" + `3` + "`" + `. - ` + "`" + `packet_size` + "`" + ` - (Optional) Size of packets in bytes. Default ` + "`" + `56` + "`" + `. - ` + "`" + `limit` + "`" + ` - (Required) One or more Limit blocks. Only one limit block may exist for each name. Limit block - ` + "`" + `name` + "`" + ` - (Required) Kind of limit. Valid values are ` + "`" + `"lossPercent"` + "`" + `, ` + "`" + `"total"` + "`" + `, ` + "`" + `"average"` + "`" + `, ` + "`" + `"run"` + "`" + ` & ` + "`" + `"avgRun"` + "`" + `. - ` + "`" + `warning` + "`" + ` - (Optional) Amount to trigger a warning. - ` + "`" + `critical` + "`" + ` - (Optional) Amount to trigger a critical. - ` + "`" + `fail` + "`" + ` - (Optional) Amount to trigger a failure.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ultradns_record",
			Category:         "Resources",
			ShortDescription: `Provides an UltraDNS record resource.`,
			Description:      ``,
			Keywords: []string{
				"record",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "zone",
					Description: `(Required) The domain to add the record to`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the record`,
				},
				resource.Attribute{
					Name:        "rdata",
					Description: `(Required) An array containing the values of the record`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) The type of the record`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `(Optional) The TTL of the record ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The record ID`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the record`,
				},
				resource.Attribute{
					Name:        "rdata",
					Description: `An array containing the values of the record`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of the record`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `The TTL of the record`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `The domain of the record`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `The FQDN of the record`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The record ID`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the record`,
				},
				resource.Attribute{
					Name:        "rdata",
					Description: `An array containing the values of the record`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of the record`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `The TTL of the record`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `The domain of the record`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `The FQDN of the record`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ultradns_tcpool",
			Category:         "Resources",
			ShortDescription: `Provides an UltraDNS Traffic Controller pool resource.`,
			Description:      ``,
			Keywords: []string{
				"tcpool",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "zone",
					Description: `(Required) The domain to add the record to`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the record`,
				},
				resource.Attribute{
					Name:        "rdata",
					Description: `(Required) a list of rdata blocks, one for each member in the pool. Record Data documented below.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Required) Description of the Traffic Controller pool. Valid values are strings less than 256 characters.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `(Optional) The TTL of the record. Default: ` + "`" + `3600` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "run_probes",
					Description: `(Optional) Boolean to run probes for this pool. Default: ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "act_on_probes",
					Description: `(Optional) Boolean to enable and disable pool records when probes are run. Default: ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "max_to_lb",
					Description: `(Optional) Determines the number of records to balance between. Valid values are integers ` + "`" + `0` + "`" + ` - ` + "`" + `len(rdata)` + "`" + `. Default: ` + "`" + `0` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "backup_record_rdata",
					Description: `(Optional) IPv4 address or CNAME for the backup record. Default: ` + "`" + `nil` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "backup_record_failover_delay",
					Description: `(Optional) Time in minutes that Traffic Controller waits after detecting that the pool record has failed before activating primary records. Valid values are integers ` + "`" + `0` + "`" + ` - ` + "`" + `30` + "`" + `. Default: ` + "`" + `0` + "`" + `. Record Data blocks support the following:`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `(Required) IPv4 address or CNAME for the pool member.`,
				},
				resource.Attribute{
					Name:        "failover_delay",
					Description: `(Optional) Time in minutes that Traffic Controller waits after detecting that the pool record has failed before activating secondary records. ` + "`" + `0` + "`" + ` will activate the secondary records immediately. Integer. Range: ` + "`" + `0` + "`" + ` - ` + "`" + `30` + "`" + `. Default: ` + "`" + `0` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `(Optional) Indicates the serving preference for this pool record. Valid values are integers ` + "`" + `1` + "`" + ` or greater. Default: ` + "`" + `1` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "run_probes",
					Description: `(Optional) Whether probes are run for this pool record. Boolean. Default: ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `(Optional) Current state of the pool record. String. Must be one of ` + "`" + `"NORMAL"` + "`" + `, ` + "`" + `"ACTIVE"` + "`" + `, or ` + "`" + `"INACTIVE"` + "`" + `. Default: ` + "`" + `"NORMAL"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "threshold",
					Description: `(Optional) How many probes must agree before the record state is changed. Valid values are integers ` + "`" + `1` + "`" + ` - ` + "`" + `len(probes)` + "`" + `. Default: ` + "`" + `1` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `(Optional) Traffic load to send to each server in the Traffic Controller pool. Valid values are integers ` + "`" + `2` + "`" + ` - ` + "`" + `100` + "`" + `. Default: ` + "`" + `2` + "`" + ` ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The record ID`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `The FQDN of the record`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The record ID`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `The FQDN of the record`,
				},
			},
		},
	}

	resourcesMap = map[string]int{

		"ultradns_dirpool":    0,
		"ultradns_probe_http": 1,
		"ultradns_probe_ping": 2,
		"ultradns_record":     3,
		"ultradns_tcpool":     4,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
