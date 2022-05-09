package elasticstack

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "elasticstack_elasticsearch_ingest_processor_append",
			Category:         "Ingest",
			ShortDescription: `Helper data source to create a processor which appends one or more values to an existing array if the field already exists and it is an array.`,
			Description: `

Helper data source to which can be used to create a processor to append one or more values to an existing array if the field already exists and it is an array.
Converts a scalar to an array and appends one or more values to it if the field exists and it is a scalar. Creates an array containing the provided values if the field doesn’t exist.

See: https://www.elastic.co/guide/en/elasticsearch/reference/current/append-processor.html

`,
			Keywords: []string{
				"ingest",
				"elasticsearch",
				"processor",
				"append",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "elasticstack_elasticsearch_ingest_processor_bytes",
			Category:         "Ingest",
			ShortDescription: `Helper data source to create a processor which converts a human readable byte value (e.g. 1kb) to its value in bytes (e.g. 1024).`,
			Description: `

Helper data source to which can be used to create a processor to convert a human readable byte value (e.g. 1kb) to its value in bytes (e.g. 1024). If the field is an array of strings, all members of the array will be converted.

Supported human readable units are "b", "kb", "mb", "gb", "tb", "pb" case insensitive. An error will occur if the field is not a supported format or resultant value exceeds 2^63.

See: https://www.elastic.co/guide/en/elasticsearch/reference/current/bytes-processor.html

`,
			Keywords: []string{
				"ingest",
				"elasticsearch",
				"processor",
				"bytes",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "elasticstack_elasticsearch_ingest_processor_circle",
			Category:         "Ingest",
			ShortDescription: `Helper data source to create a processor which converts circle definitions of shapes to regular polygons which approximate them.`,
			Description: `

Helper data source to which can be used to create a processor to convert circle definitions of shapes to regular polygons which approximate them.

See: https://www.elastic.co/guide/en/elasticsearch/reference/current/ingest-circle-processor.html

`,
			Keywords: []string{
				"ingest",
				"elasticsearch",
				"processor",
				"circle",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "elasticstack_elasticsearch_ingest_processor_community_id",
			Category:         "Ingest",
			ShortDescription: `Helper data source to create a processor which computes the Community ID for network flow data as defined in the Community ID Specification.`,
			Description: `

Helper data source to which can be used to create a processor to compute the Community ID for network flow data as defined in the [Community ID Specification](https://github.com/corelight/community-id-spec). 
You can use a community ID to correlate network events related to a single flow.

The community ID processor reads network flow data from related [Elastic Common Schema (ECS)](https://www.elastic.co/guide/en/ecs/1.12) fields by default. If you use the ECS, no configuration is required.

See: https://www.elastic.co/guide/en/elasticsearch/reference/current/community-id-processor.html

`,
			Keywords: []string{
				"ingest",
				"elasticsearch",
				"processor",
				"community",
				"id",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "elasticstack_elasticsearch_ingest_processor_convert",
			Category:         "Ingest",
			ShortDescription: `Helper data source to create a processor which converts a field in the currently ingested document to a different type, such as converting a string to an integer.`,
			Description: `

Helper data source to which can be used to convert a field in the currently ingested document to a different type, such as converting a string to an integer. If the field value is an array, all members will be converted.

The supported types include: ` + "`" + `integer` + "`" + `, ` + "`" + `long` + "`" + `, ` + "`" + `float` + "`" + `, ` + "`" + `double` + "`" + `, ` + "`" + `string` + "`" + `, ` + "`" + `boolean` + "`" + `, ` + "`" + `ip` + "`" + `, and ` + "`" + `auto` + "`" + `.

Specifying ` + "`" + `boolean` + "`" + ` will set the field to true if its string value is equal to true (ignore case), to false if its string value is equal to false (ignore case), or it will throw an exception otherwise.

Specifying ` + "`" + `ip` + "`" + ` will set the target field to the value of ` + "`" + `field` + "`" + ` if it contains a valid IPv4 or IPv6 address that can be indexed into an IP field type.

Specifying ` + "`" + `auto` + "`" + ` will attempt to convert the string-valued ` + "`" + `field` + "`" + ` into the closest non-string, non-IP type. For example, a field whose value is "true" will be converted to its respective boolean type: true. Do note that float takes precedence of double in auto. A value of "242.15" will "automatically" be converted to 242.15 of type ` + "`" + `float` + "`" + `. If a provided field cannot be appropriately converted, the processor will still process successfully and leave the field value as-is. In such a case, ` + "`" + `target_field` + "`" + ` will be updated with the unconverted field value.

See: https://www.elastic.co/guide/en/elasticsearch/reference/current/convert-processor.html

`,
			Keywords: []string{
				"ingest",
				"elasticsearch",
				"processor",
				"convert",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "elasticstack_elasticsearch_ingest_processor_csv",
			Category:         "Ingest",
			ShortDescription: `Helper data source to create a processor which extracts fields from CSV line out of a single text field within a document.`,
			Description: `

Helper data source to which can be used to extract fields from CSV line out of a single text field within a document. Any empty field in CSV will be skipped.

See: https://www.elastic.co/guide/en/elasticsearch/reference/current/csv-processor.html

`,
			Keywords: []string{
				"ingest",
				"elasticsearch",
				"processor",
				"csv",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "elasticstack_elasticsearch_ingest_processor_date",
			Category:         "Ingest",
			ShortDescription: `Helper data source to create a processor which parses dates from fields, and then uses the date or timestamp as the timestamp for the document.`,
			Description: `

Helper data source to which can be used to parse dates from fields, and then uses the date or timestamp as the timestamp for the document. 
By default, the date processor adds the parsed date as a new field called ` + "`" + `@timestamp` + "`" + `. You can specify a different field by setting the ` + "`" + `target_field` + "`" + ` configuration parameter. Multiple date formats are supported as part of the same date processor definition. They will be used sequentially to attempt parsing the date field, in the same order they were defined as part of the processor definition.

See: https://www.elastic.co/guide/en/elasticsearch/reference/current/date-processor.html

`,
			Keywords: []string{
				"ingest",
				"elasticsearch",
				"processor",
				"date",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "elasticstack_elasticsearch_ingest_processor_date_index_name",
			Category:         "Ingest",
			ShortDescription: `Helper data source to create a processor which helps to point documents to the right time based index based on a date or timestamp field in a document by using the date math index name support.`,
			Description: `

The purpose of this processor is to point documents to the right time based index based on a date or timestamp field in a document by using the date math index name support.

The processor sets the _index metadata field with a date math index name expression based on the provided index name prefix, a date or timestamp field in the documents being processed and the provided date rounding.

First, this processor fetches the date or timestamp from a field in the document being processed. Optionally, date formatting can be configured on how the field’s value should be parsed into a date. Then this date, the provided index name prefix and the provided date rounding get formatted into a date math index name expression. Also here optionally date formatting can be specified on how the date should be formatted into a date math index name expression.

See: https://www.elastic.co/guide/en/elasticsearch/reference/current/date-index-name-processor.html

`,
			Keywords: []string{
				"ingest",
				"elasticsearch",
				"processor",
				"date",
				"index",
				"name",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "elasticstack_elasticsearch_ingest_processor_dissect",
			Category:         "Ingest",
			ShortDescription: `Helper data source to create a processor which extracts structured fields out of a single text field within a document.`,
			Description: `

Similar to the Grok Processor, dissect also extracts structured fields out of a single text field within a document. However unlike the Grok Processor, dissect does not use Regular Expressions. This allows dissect’s syntax to be simple and for some cases faster than the Grok Processor.

Dissect matches a single text field against a defined pattern.


See: https://www.elastic.co/guide/en/elasticsearch/reference/current/dissect-processor.html

`,
			Keywords: []string{
				"ingest",
				"elasticsearch",
				"processor",
				"dissect",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "elasticstack_elasticsearch_ingest_processor_dot_expander",
			Category:         "Ingest",
			ShortDescription: `Helper data source to create a processor which expands a field with dots into an object field.`,
			Description: `

Expands a field with dots into an object field. This processor allows fields with dots in the name to be accessible by other processors in the pipeline. Otherwise these fields can’t be accessed by any processor.

See: elastic.co/guide/en/elasticsearch/reference/current/dot-expand-processor.html


`,
			Keywords: []string{
				"ingest",
				"elasticsearch",
				"processor",
				"dot",
				"expander",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "elasticstack_elasticsearch_ingest_processor_drop",
			Category:         "Ingest",
			ShortDescription: `Helper data source to create a processor which drops the document without raising any errors.`,
			Description: `

Drops the document without raising any errors. This is useful to prevent the document from getting indexed based on some condition.

See: https://www.elastic.co/guide/en/elasticsearch/reference/current/drop-processor.html


`,
			Keywords: []string{
				"ingest",
				"elasticsearch",
				"processor",
				"drop",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "elasticstack_elasticsearch_ingest_processor_enrich",
			Category:         "Ingest",
			ShortDescription: `Helper data source to create a processor which enriches documents with data from another index.`,
			Description: `

The enrich processor can enrich documents with data from another index. See enrich data section for more information about how to set this up.

See: https://www.elastic.co/guide/en/elasticsearch/reference/current/ingest-enriching-data.html and https://www.elastic.co/guide/en/elasticsearch/reference/current/enrich-processor.html

`,
			Keywords: []string{
				"ingest",
				"elasticsearch",
				"processor",
				"enrich",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "elasticstack_elasticsearch_ingest_processor_fail",
			Category:         "Ingest",
			ShortDescription: `Helper data source to create a processor which raises an exception.`,
			Description: `

Raises an exception. This is useful for when you expect a pipeline to fail and want to relay a specific message to the requester.

See: https://www.elastic.co/guide/en/elasticsearch/reference/current/fail-processor.html


`,
			Keywords: []string{
				"ingest",
				"elasticsearch",
				"processor",
				"fail",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "elasticstack_elasticsearch_ingest_processor_fingerprint",
			Category:         "Ingest",
			ShortDescription: `Helper data source to create a processor which computes a hash of the document’s content.`,
			Description: `

Computes a hash of the document’s content. You can use this hash for content fingerprinting.

See: https://www.elastic.co/guide/en/elasticsearch/reference/current/fingerprint-processor.html

`,
			Keywords: []string{
				"ingest",
				"elasticsearch",
				"processor",
				"fingerprint",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "elasticstack_elasticsearch_ingest_processor_foreach",
			Category:         "Ingest",
			ShortDescription: `Helper data source to create a processor which runs an ingest processor on each element of an array or object.`,
			Description: `

Runs an ingest processor on each element of an array or object.

All ingest processors can run on array or object elements. However, if the number of elements is unknown, it can be cumbersome to process each one in the same way.

The ` + "`" + `foreach` + "`" + ` processor lets you specify a ` + "`" + `field` + "`" + ` containing array or object values and a ` + "`" + `processor` + "`" + ` to run on each element in the field.

See: https://www.elastic.co/guide/en/elasticsearch/reference/current/foreach-processor.html


#`,
			Keywords: []string{
				"ingest",
				"elasticsearch",
				"processor",
				"foreach",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "elasticstack_elasticsearch_ingest_processor_geoip",
			Category:         "Ingest",
			ShortDescription: `Helper data source to create a processor which adds information about the geographical location of an IPv4 or IPv6 address.`,
			Description: `

The geoip processor adds information about the geographical location of an IPv4 or IPv6 address.

By default, the processor uses the GeoLite2 City, GeoLite2 Country, and GeoLite2 ASN GeoIP2 databases from MaxMind, shared under the CC BY-SA 4.0 license. Elasticsearch automatically downloads updates for these databases from the Elastic GeoIP endpoint: https://geoip.elastic.co/v1/database. To get download statistics for these updates, use the GeoIP stats API.

If your cluster can’t connect to the Elastic GeoIP endpoint or you want to manage your own updates, [see Manage your own GeoIP2 database updates](https://www.elastic.co/guide/en/elasticsearch/reference/current/geoip-processor.html#manage-geoip-database-updates).

If Elasticsearch can’t connect to the endpoint for 30 days all updated databases will become invalid. Elasticsearch will stop enriching documents with geoip data and will add tags: ["_geoip_expired_database"] field instead.


See: https://www.elastic.co/guide/en/elasticsearch/reference/current/geoip-processor.html


`,
			Keywords: []string{
				"ingest",
				"elasticsearch",
				"processor",
				"geoip",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "elasticstack_elasticsearch_ingest_processor_grok",
			Category:         "Ingest",
			ShortDescription: `Helper data source to create a processor which extracts structured fields out of a single text field within a document.`,
			Description: `

Extracts structured fields out of a single text field within a document. You choose which field to extract matched fields from, as well as the grok pattern you expect will match. A grok pattern is like a regular expression that supports aliased expressions that can be reused.

This processor comes packaged with many [reusable patterns](https://github.com/elastic/elasticsearch/blob/master/libs/grok/src/main/resources/patterns).

If you need help building patterns to match your logs, you will find the [Grok Debugger](https://www.elastic.co/guide/en/kibana/master/xpack-grokdebugger.html) tool quite useful! [The Grok Constructor](https://grokconstructor.appspot.com/) is also a useful tool.


See: https://www.elastic.co/guide/en/elasticsearch/reference/current/grok-processor.html


`,
			Keywords: []string{
				"ingest",
				"elasticsearch",
				"processor",
				"grok",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "elasticstack_elasticsearch_ingest_processor_gsub",
			Category:         "Ingest",
			ShortDescription: `Helper data source to create a processor which converts a string field by applying a regular expression and a replacement.`,
			Description: `

Converts a string field by applying a regular expression and a replacement. If the field is an array of string, all members of the array will be converted. If any non-string values are encountered, the processor will throw an exception.

See: https://www.elastic.co/guide/en/elasticsearch/reference/current/gsub-processor.html


`,
			Keywords: []string{
				"ingest",
				"elasticsearch",
				"processor",
				"gsub",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "elasticstack_elasticsearch_ingest_processor_html_strip",
			Category:         "Ingest",
			ShortDescription: `Helper data source to create a processor which removes HTML tags from the field.`,
			Description: `

Removes HTML tags from the field. If the field is an array of strings, HTML tags will be removed from all members of the array.

See: templates/data-sources/elasticsearch_ingest_processor_html_strip.md.tmpl


`,
			Keywords: []string{
				"ingest",
				"elasticsearch",
				"processor",
				"html",
				"strip",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "elasticstack_elasticsearch_ingest_processor_join",
			Category:         "Ingest",
			ShortDescription: `Helper data source to create a processor which joins each element of an array into a single string using a separator character between each element.`,
			Description: `

Joins each element of an array into a single string using a separator character between each element. Throws an error when the field is not an array.

See: https://www.elastic.co/guide/en/elasticsearch/reference/current/join-processor.html


`,
			Keywords: []string{
				"ingest",
				"elasticsearch",
				"processor",
				"join",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "elasticstack_elasticsearch_ingest_processor_json",
			Category:         "Ingest",
			ShortDescription: `Helper data source to create a processor which converts a JSON string into a structured JSON object.`,
			Description: `

Converts a JSON string into a structured JSON object.

See: https://www.elastic.co/guide/en/elasticsearch/reference/current/json-processor.html

`,
			Keywords: []string{
				"ingest",
				"elasticsearch",
				"processor",
				"json",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "elasticstack_elasticsearch_ingest_processor_kv",
			Category:         "Ingest",
			ShortDescription: `Helper data source to create a processor which helps automatically parse messages (or specific event fields) which are of the ` + "`" + `foo=bar` + "`" + ` variety.`,
			Description: `

This processor helps automatically parse messages (or specific event fields) which are of the ` + "`" + `foo=bar` + "`" + ` variety.

See: https://www.elastic.co/guide/en/elasticsearch/reference/current/kv-processor.html


`,
			Keywords: []string{
				"ingest",
				"elasticsearch",
				"processor",
				"kv",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "elasticstack_elasticsearch_ingest_processor_lowercase",
			Category:         "Ingest",
			ShortDescription: `Helper data source to create a processor which converts a string to its lowercase equivalent.`,
			Description: `

Converts a string to its lowercase equivalent. If the field is an array of strings, all members of the array will be converted.

See: https://www.elastic.co/guide/en/elasticsearch/reference/current/lowercase-processor.html


`,
			Keywords: []string{
				"ingest",
				"elasticsearch",
				"processor",
				"lowercase",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "elasticstack_elasticsearch_ingest_processor_network_direction",
			Category:         "Ingest",
			ShortDescription: `Helper data source to create a processor which calculates the network direction given a source IP address, destination IP address, and a list of internal networks.`,
			Description: `

Calculates the network direction given a source IP address, destination IP address, and a list of internal networks.

The network direction processor reads IP addresses from Elastic Common Schema (ECS) fields by default. If you use the ECS, only the ` + "`" + `internal_networks` + "`" + ` option must be specified.


One of either ` + "`" + `internal_networks` + "`" + ` or ` + "`" + `internal_networks_field` + "`" + ` must be specified. If ` + "`" + `internal_networks_field` + "`" + ` is specified, it follows the behavior specified by ` + "`" + `ignore_missing` + "`" + `.

#`,
			Keywords: []string{
				"ingest",
				"elasticsearch",
				"processor",
				"network",
				"direction",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "elasticstack_elasticsearch_ingest_processor_pipeline",
			Category:         "Ingest",
			ShortDescription: `Helper data source to create a processor which executes another pipeline.`,
			Description: `

Executes another pipeline.

The name of the current pipeline can be accessed from the ` + "`" + `_ingest.pipeline` + "`" + ` ingest metadata key.

See: https://www.elastic.co/guide/en/elasticsearch/reference/current/pipeline-processor.html


`,
			Keywords: []string{
				"ingest",
				"elasticsearch",
				"processor",
				"pipeline",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "elasticstack_elasticsearch_ingest_processor_registered_domain",
			Category:         "Ingest",
			ShortDescription: `Helper data source to create a processor which Extracts the registered domain, sub-domain, and top-level domain from a fully qualified domain name.`,
			Description: `

Extracts the registered domain (also known as the effective top-level domain or eTLD), sub-domain, and top-level domain from a fully qualified domain name (FQDN). Uses the registered domains defined in the Mozilla Public Suffix List.

See: https://www.elastic.co/guide/en/elasticsearch/reference/current/registered-domain-processor.html


`,
			Keywords: []string{
				"ingest",
				"elasticsearch",
				"processor",
				"registered",
				"domain",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "elasticstack_elasticsearch_ingest_processor_remove",
			Category:         "Ingest",
			ShortDescription: `Helper data source to create a processor which removes existing fields.`,
			Description: `

Removes existing fields. If one field doesn’t exist, an exception will be thrown.

See: https://www.elastic.co/guide/en/elasticsearch/reference/current/remove-processor.html


`,
			Keywords: []string{
				"ingest",
				"elasticsearch",
				"processor",
				"remove",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "elasticstack_elasticsearch_ingest_processor_rename",
			Category:         "Ingest",
			ShortDescription: `Helper data source to create a processor which renames an existing field.`,
			Description: `

Renames an existing field. If the field doesn’t exist or the new name is already used, an exception will be thrown.

See: https://www.elastic.co/guide/en/elasticsearch/reference/current/rename-processor.html


`,
			Keywords: []string{
				"ingest",
				"elasticsearch",
				"processor",
				"rename",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "elasticstack_elasticsearch_ingest_processor_script",
			Category:         "Ingest",
			ShortDescription: `Helper data source to create a processor which runs an inline or stored script on incoming documents.`,
			Description: `

Runs an inline or stored script on incoming documents. The script runs in the ingest context.

The script processor uses the script cache to avoid recompiling the script for each incoming document. To improve performance, ensure the script cache is properly sized before using a script processor in production.

See: https://www.elastic.co/guide/en/elasticsearch/reference/current/script-processor.html

#`,
			Keywords: []string{
				"ingest",
				"elasticsearch",
				"processor",
				"script",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "elasticstack_elasticsearch_ingest_processor_set",
			Category:         "Ingest",
			ShortDescription: `Helper data source to create a processor which sets one field and associates it with the specified value.`,
			Description: `

Sets one field and associates it with the specified value. If the field already exists, its value will be replaced with the provided one.

See: https://www.elastic.co/guide/en/elasticsearch/reference/current/set-processor.html


`,
			Keywords: []string{
				"ingest",
				"elasticsearch",
				"processor",
				"set",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "elasticstack_elasticsearch_ingest_processor_set_security_user",
			Category:         "Ingest",
			ShortDescription: `Helper data source to create a processor which sets user-related details from the current authenticated user to the current document by pre-processing the ingest.`,
			Description: `

Sets user-related details (such as ` + "`" + `username` + "`" + `, ` + "`" + `roles` + "`" + `, ` + "`" + `email` + "`" + `, ` + "`" + `full_name` + "`" + `, ` + "`" + `metadata` + "`" + `, ` + "`" + `api_key` + "`" + `, ` + "`" + `realm` + "`" + ` and ` + "`" + `authentication_typ` + "`" + `e) from the current authenticated user to the current document by pre-processing the ingest. The ` + "`" + `api_key` + "`" + ` property exists only if the user authenticates with an API key. It is an object containing the id, name and metadata (if it exists and is non-empty) fields of the API key. The realm property is also an object with two fields, name and type. When using API key authentication, the realm property refers to the realm from which the API key is created. The ` + "`" + `authentication_type property` + "`" + ` is a string that can take value from ` + "`" + `REALM` + "`" + `, ` + "`" + `API_KEY` + "`" + `, ` + "`" + `TOKEN` + "`" + ` and ` + "`" + `ANONYMOUS` + "`" + `.

See: https://www.elastic.co/guide/en/elasticsearch/reference/current/ingest-node-set-security-user-processor.html


`,
			Keywords: []string{
				"ingest",
				"elasticsearch",
				"processor",
				"set",
				"security",
				"user",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "elasticstack_elasticsearch_ingest_processor_sort",
			Category:         "Ingest",
			ShortDescription: `Helper data source to create a processor which sorts the elements of an array ascending or descending.`,
			Description: `

Sorts the elements of an array ascending or descending. Homogeneous arrays of numbers will be sorted numerically, while arrays of strings or heterogeneous arrays of strings + numbers will be sorted lexicographically. Throws an error when the field is not an array.

See: https://www.elastic.co/guide/en/elasticsearch/reference/current/sort-processor.html


`,
			Keywords: []string{
				"ingest",
				"elasticsearch",
				"processor",
				"sort",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "elasticstack_elasticsearch_ingest_processor_split",
			Category:         "Ingest",
			ShortDescription: `Helper data source to create a processor which splits a field into an array using a separator character.`,
			Description: `

Splits a field into an array using a separator character. Only works on string fields.

See: https://www.elastic.co/guide/en/elasticsearch/reference/current/split-processor.html


`,
			Keywords: []string{
				"ingest",
				"elasticsearch",
				"processor",
				"split",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "elasticstack_elasticsearch_ingest_processor_trim",
			Category:         "Ingest",
			ShortDescription: `Helper data source to create a processor which trims whitespace from field.`,
			Description: `

Trims whitespace from field. If the field is an array of strings, all members of the array will be trimmed.

**NOTE:** This only works on leading and trailing whitespace.

See: https://www.elastic.co/guide/en/elasticsearch/reference/current/trim-processor.html


`,
			Keywords: []string{
				"ingest",
				"elasticsearch",
				"processor",
				"trim",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "elasticstack_elasticsearch_ingest_processor_uppercase",
			Category:         "Ingest",
			ShortDescription: `Helper data source to create a processor which converts a string to its uppercase equivalent.`,
			Description: `

Converts a string to its uppercase equivalent. If the field is an array of strings, all members of the array will be converted.

See: https://www.elastic.co/guide/en/elasticsearch/reference/current/uppercase-processor.html


`,
			Keywords: []string{
				"ingest",
				"elasticsearch",
				"processor",
				"uppercase",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "elasticstack_elasticsearch_ingest_processor_uri_parts",
			Category:         "Ingest",
			ShortDescription: `Helper data source to create a processor which parses a Uniform Resource Identifier (URI) string and extracts its components as an object.`,
			Description: `

Parses a Uniform Resource Identifier (URI) string and extracts its components as an object. This URI object includes properties for the URI’s domain, path, fragment, port, query, scheme, user info, username, and password.

See: https://www.elastic.co/guide/en/elasticsearch/reference/current/uri-parts-processor.html


`,
			Keywords: []string{
				"ingest",
				"elasticsearch",
				"processor",
				"uri",
				"parts",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "elasticstack_elasticsearch_ingest_processor_urldecode",
			Category:         "Ingest",
			ShortDescription: `Helper data source to create a processor which URL-decodes a string.`,
			Description: `

URL-decodes a string. If the field is an array of strings, all members of the array will be decoded.

See: https://www.elastic.co/guide/en/elasticsearch/reference/current/urldecode-processor.html


`,
			Keywords: []string{
				"ingest",
				"elasticsearch",
				"processor",
				"urldecode",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "elasticstack_elasticsearch_ingest_processor_user_agent",
			Category:         "Ingest",
			ShortDescription: `Helper data source to create a processor which extracts details from the user agent string a browser sends with its web requests.`,
			Description: `

The ` + "`" + `user_agent` + "`" + ` processor extracts details from the user agent string a browser sends with its web requests. This processor adds this information by default under the ` + "`" + `user_agent` + "`" + ` field.

The ingest-user-agent module ships by default with the regexes.yaml made available by uap-java with an Apache 2.0 license. For more details see https://github.com/ua-parser/uap-core.


See: https://www.elastic.co/guide/en/elasticsearch/reference/current/user-agent-processor.html


`,
			Keywords: []string{
				"ingest",
				"elasticsearch",
				"processor",
				"user",
				"agent",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "elasticstack_elasticsearch_security_user",
			Category:         "Security",
			ShortDescription: `Gets information about Elasticsearch user.`,
			Description: `

Use this data source to get information about existing Elasticsearch user. See, https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-get-user.html".

`,
			Keywords: []string{
				"security",
				"elasticsearch",
				"user",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "elasticstack_elasticsearch_snapshot_repository",
			Category:         "Snapshot",
			ShortDescription: `Gets information about the registered snapshot repositories.`,
			Description: `

This data source provides the information about the registered snaphosts repositories

`,
			Keywords: []string{
				"snapshot",
				"elasticsearch",
				"repository",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
	}

	dataSourcesMap = map[string]int{

		"elasticstack_elasticsearch_ingest_processor_append":            0,
		"elasticstack_elasticsearch_ingest_processor_bytes":             1,
		"elasticstack_elasticsearch_ingest_processor_circle":            2,
		"elasticstack_elasticsearch_ingest_processor_community_id":      3,
		"elasticstack_elasticsearch_ingest_processor_convert":           4,
		"elasticstack_elasticsearch_ingest_processor_csv":               5,
		"elasticstack_elasticsearch_ingest_processor_date":              6,
		"elasticstack_elasticsearch_ingest_processor_date_index_name":   7,
		"elasticstack_elasticsearch_ingest_processor_dissect":           8,
		"elasticstack_elasticsearch_ingest_processor_dot_expander":      9,
		"elasticstack_elasticsearch_ingest_processor_drop":              10,
		"elasticstack_elasticsearch_ingest_processor_enrich":            11,
		"elasticstack_elasticsearch_ingest_processor_fail":              12,
		"elasticstack_elasticsearch_ingest_processor_fingerprint":       13,
		"elasticstack_elasticsearch_ingest_processor_foreach":           14,
		"elasticstack_elasticsearch_ingest_processor_geoip":             15,
		"elasticstack_elasticsearch_ingest_processor_grok":              16,
		"elasticstack_elasticsearch_ingest_processor_gsub":              17,
		"elasticstack_elasticsearch_ingest_processor_html_strip":        18,
		"elasticstack_elasticsearch_ingest_processor_join":              19,
		"elasticstack_elasticsearch_ingest_processor_json":              20,
		"elasticstack_elasticsearch_ingest_processor_kv":                21,
		"elasticstack_elasticsearch_ingest_processor_lowercase":         22,
		"elasticstack_elasticsearch_ingest_processor_network_direction": 23,
		"elasticstack_elasticsearch_ingest_processor_pipeline":          24,
		"elasticstack_elasticsearch_ingest_processor_registered_domain": 25,
		"elasticstack_elasticsearch_ingest_processor_remove":            26,
		"elasticstack_elasticsearch_ingest_processor_rename":            27,
		"elasticstack_elasticsearch_ingest_processor_script":            28,
		"elasticstack_elasticsearch_ingest_processor_set":               29,
		"elasticstack_elasticsearch_ingest_processor_set_security_user": 30,
		"elasticstack_elasticsearch_ingest_processor_sort":              31,
		"elasticstack_elasticsearch_ingest_processor_split":             32,
		"elasticstack_elasticsearch_ingest_processor_trim":              33,
		"elasticstack_elasticsearch_ingest_processor_uppercase":         34,
		"elasticstack_elasticsearch_ingest_processor_uri_parts":         35,
		"elasticstack_elasticsearch_ingest_processor_urldecode":         36,
		"elasticstack_elasticsearch_ingest_processor_user_agent":        37,
		"elasticstack_elasticsearch_security_user":                      38,
		"elasticstack_elasticsearch_snapshot_repository":                39,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
