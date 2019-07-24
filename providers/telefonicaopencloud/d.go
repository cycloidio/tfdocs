package telefonicaopencloud

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "telefonicaopencloud_dns_zone_v2",
			Category:         "Data Sources",
			ShortDescription: `Get information on an TelefonicaOpenCloud DNS Zone.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 DNS client. A DNS client is needed to retrieve zone ids. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the zone.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A description of the zone.`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `(Optional) The email contact for the zone record.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Optional) The zone's status.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `(Optional) The time to live (TTL) of the zone.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) The type of the zone. Can either be ` + "`" + `PRIMARY` + "`" + ` or ` + "`" + `SECONDARY` + "`" + `. ## Attributes Reference ` + "`" + `id` + "`" + ` is set to the ID of the found zone. In addition, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "attributes",
					Description: `Attributes of the DNS Service scheduler.`,
				},
				resource.Attribute{
					Name:        "masters",
					Description: `An array of master DNS servers. When ` + "`" + `type` + "`" + ` is ` + "`" + `SECONDARY` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `The time the zone was created.`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `The time the zone was last updated.`,
				},
				resource.Attribute{
					Name:        "transferred_at",
					Description: `The time the zone was transferred.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `The version of the zone.`,
				},
				resource.Attribute{
					Name:        "serial",
					Description: `The serial number of the zone.`,
				},
				resource.Attribute{
					Name:        "pool_id",
					Description: `The ID of the pool hosting the zone.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `The project ID that owns the zone.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "attributes",
					Description: `Attributes of the DNS Service scheduler.`,
				},
				resource.Attribute{
					Name:        "masters",
					Description: `An array of master DNS servers. When ` + "`" + `type` + "`" + ` is ` + "`" + `SECONDARY` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `The time the zone was created.`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `The time the zone was last updated.`,
				},
				resource.Attribute{
					Name:        "transferred_at",
					Description: `The time the zone was transferred.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `The version of the zone.`,
				},
				resource.Attribute{
					Name:        "serial",
					Description: `The serial number of the zone.`,
				},
				resource.Attribute{
					Name:        "pool_id",
					Description: `The ID of the pool hosting the zone.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `The project ID that owns the zone.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "telefonicaopencloud_networking_network_v2",
			Category:         "Data Sources",
			ShortDescription: `Get information on an TelefonicaOpenCloud Network.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 Neutron client. A Neutron client is needed to retrieve networks ids. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used.`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `(Optional) The ID of the network.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the network.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Optional) The status of the network.`,
				},
				resource.Attribute{
					Name:        "matching_subnet_cidr",
					Description: `(Optional) The CIDR of a subnet within the network.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Optional) The owner of the network. ## Attributes Reference ` + "`" + `id` + "`" + ` is set to the ID of the found network. In addition, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `(Optional) The administrative state of the network.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "shared",
					Description: `(Optional) Specifies whether the network resource can be accessed by any tenant or not.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `(Optional) The administrative state of the network.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "shared",
					Description: `(Optional) Specifies whether the network resource can be accessed by any tenant or not.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "telefonicaopencloud_networking_secgroup_v2",
			Category:         "Data Sources",
			ShortDescription: `Get information on an TelefonicaOpenCloud Security Group.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 Neutron client. A Neutron client is needed to retrieve security groups ids. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used.`,
				},
				resource.Attribute{
					Name:        "secgroup_id",
					Description: `(Optional) The ID of the security group.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the security group.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Optional) The owner of the security group. ## Attributes Reference ` + "`" + `id` + "`" + ` is set to the ID of the found security group. In addition, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "telefonicaopencloud_networking_subnet_v2",
			Category:         "Data Sources",
			ShortDescription: `Get information on an TelefonicaOpenCloud Subnet.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 Neutron client. A Neutron client is needed to retrieve subnet ids. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the subnet.`,
				},
				resource.Attribute{
					Name:        "dhcp_enabled",
					Description: `(Optional) If the subnet has DHCP enabled.`,
				},
				resource.Attribute{
					Name:        "dhcp_disabled",
					Description: `(Optional) If the subnet has DHCP disabled.`,
				},
				resource.Attribute{
					Name:        "ip_version",
					Description: `(Optional) The IP version of the subnet (either 4 or 6).`,
				},
				resource.Attribute{
					Name:        "gateway_ip",
					Description: `(Optional) The IP of the subnet's gateway.`,
				},
				resource.Attribute{
					Name:        "cidr",
					Description: `(Optional) The CIDR of the subnet.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Optional) The ID of the subnet.`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `(Optional) The ID of the network the subnet belongs to.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Optional) The owner of the subnet. ## Attributes Reference ` + "`" + `id` + "`" + ` is set to the ID of the found subnet. In addition, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "allocation_pools",
					Description: `Allocation pools of the subnet.`,
				},
				resource.Attribute{
					Name:        "enable_dhcp",
					Description: `Whether the subnet has DHCP enabled or not.`,
				},
				resource.Attribute{
					Name:        "dns_nameservers",
					Description: `DNS Nameservers of the subnet.`,
				},
				resource.Attribute{
					Name:        "host_routes",
					Description: `Host Routes of the subnet.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "allocation_pools",
					Description: `Allocation pools of the subnet.`,
				},
				resource.Attribute{
					Name:        "enable_dhcp",
					Description: `Whether the subnet has DHCP enabled or not.`,
				},
				resource.Attribute{
					Name:        "dns_nameservers",
					Description: `DNS Nameservers of the subnet.`,
				},
				resource.Attribute{
					Name:        "host_routes",
					Description: `Host Routes of the subnet.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "telefonicaopencloud_s3_bucket_object",
			Category:         "Data Sources",
			ShortDescription: `Provides metadata and optionally content of an S3 object`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "bucket",
					Description: `(Required) The name of the bucket to read the object from`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) The full path to the object inside the bucket`,
				},
				resource.Attribute{
					Name:        "version_id",
					Description: `(Optional) Specific version ID of the object returned (defaults to latest version) ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "body",
					Description: `Object data (see`,
				},
				resource.Attribute{
					Name:        "cache_control",
					Description: `Specifies caching behavior along the request/reply chain.`,
				},
				resource.Attribute{
					Name:        "content_disposition",
					Description: `Specifies presentational information for the object.`,
				},
				resource.Attribute{
					Name:        "content_encoding",
					Description: `Specifies what content encodings have been applied to the object and thus what decoding mechanisms must be applied to obtain the media-type referenced by the Content-Type header field.`,
				},
				resource.Attribute{
					Name:        "content_language",
					Description: `The language the content is in.`,
				},
				resource.Attribute{
					Name:        "content_length",
					Description: `Size of the body in bytes.`,
				},
				resource.Attribute{
					Name:        "content_type",
					Description: `A standard MIME type describing the format of the object data.`,
				},
				resource.Attribute{
					Name:        "etag",
					Description: `[ETag](https://en.wikipedia.org/wiki/HTTP_ETag) generated for the object (an MD5 sum of the object content in case it's not encrypted)`,
				},
				resource.Attribute{
					Name:        "expiration",
					Description: `If the object expiration is configured (see [object lifecycle management](http://docs.telefonicaopencloud.amazon.com/AmazonS3/latest/dev/object-lifecycle-mgmt.html)), the field includes this header. It includes the expiry-date and rule-id key value pairs providing object expiration information. The value of the rule-id is URL encoded.`,
				},
				resource.Attribute{
					Name:        "expires",
					Description: `The date and time at which the object is no longer cacheable.`,
				},
				resource.Attribute{
					Name:        "last_modified",
					Description: `Last modified date of the object in RFC1123 format (e.g. ` + "`" + `Mon, 02 Jan 2006 15:04:05 MST` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `A map of metadata stored with the object in S3`,
				},
				resource.Attribute{
					Name:        "server_side_encryption",
					Description: `If the object is stored using server-side encryption (KMS or Amazon S3-managed encryption key), this field includes the chosen encryption and algorithm used.`,
				},
				resource.Attribute{
					Name:        "sse_kms_key_id",
					Description: `If present, specifies the ID of the Key Management Service (KMS) master encryption key that was used for the object.`,
				},
				resource.Attribute{
					Name:        "storage_class",
					Description: `[Storage class](http://docs.aws.amazon.com/AmazonS3/latest/dev/storage-class-intro.html) information of the object. Available for all objects except for ` + "`" + `Standard` + "`" + ` storage class objects.`,
				},
				resource.Attribute{
					Name:        "version_id",
					Description: `The latest version ID of the object returned.`,
				},
				resource.Attribute{
					Name:        "website_redirect_location",
					Description: `If the bucket is configured as a website, redirects requests for this object to another object in the same bucket or to an external URL. Amazon S3 stores the value of this header in the object metadata.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "body",
					Description: `Object data (see`,
				},
				resource.Attribute{
					Name:        "cache_control",
					Description: `Specifies caching behavior along the request/reply chain.`,
				},
				resource.Attribute{
					Name:        "content_disposition",
					Description: `Specifies presentational information for the object.`,
				},
				resource.Attribute{
					Name:        "content_encoding",
					Description: `Specifies what content encodings have been applied to the object and thus what decoding mechanisms must be applied to obtain the media-type referenced by the Content-Type header field.`,
				},
				resource.Attribute{
					Name:        "content_language",
					Description: `The language the content is in.`,
				},
				resource.Attribute{
					Name:        "content_length",
					Description: `Size of the body in bytes.`,
				},
				resource.Attribute{
					Name:        "content_type",
					Description: `A standard MIME type describing the format of the object data.`,
				},
				resource.Attribute{
					Name:        "etag",
					Description: `[ETag](https://en.wikipedia.org/wiki/HTTP_ETag) generated for the object (an MD5 sum of the object content in case it's not encrypted)`,
				},
				resource.Attribute{
					Name:        "expiration",
					Description: `If the object expiration is configured (see [object lifecycle management](http://docs.telefonicaopencloud.amazon.com/AmazonS3/latest/dev/object-lifecycle-mgmt.html)), the field includes this header. It includes the expiry-date and rule-id key value pairs providing object expiration information. The value of the rule-id is URL encoded.`,
				},
				resource.Attribute{
					Name:        "expires",
					Description: `The date and time at which the object is no longer cacheable.`,
				},
				resource.Attribute{
					Name:        "last_modified",
					Description: `Last modified date of the object in RFC1123 format (e.g. ` + "`" + `Mon, 02 Jan 2006 15:04:05 MST` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `A map of metadata stored with the object in S3`,
				},
				resource.Attribute{
					Name:        "server_side_encryption",
					Description: `If the object is stored using server-side encryption (KMS or Amazon S3-managed encryption key), this field includes the chosen encryption and algorithm used.`,
				},
				resource.Attribute{
					Name:        "sse_kms_key_id",
					Description: `If present, specifies the ID of the Key Management Service (KMS) master encryption key that was used for the object.`,
				},
				resource.Attribute{
					Name:        "storage_class",
					Description: `[Storage class](http://docs.aws.amazon.com/AmazonS3/latest/dev/storage-class-intro.html) information of the object. Available for all objects except for ` + "`" + `Standard` + "`" + ` storage class objects.`,
				},
				resource.Attribute{
					Name:        "version_id",
					Description: `The latest version ID of the object returned.`,
				},
				resource.Attribute{
					Name:        "website_redirect_location",
					Description: `If the bucket is configured as a website, redirects requests for this object to another object in the same bucket or to an external URL. Amazon S3 stores the value of this header in the object metadata.`,
				},
			},
		},
	}

	dataSourcesMap = map[string]int{

		"telefonicaopencloud_dns_zone_v2":            0,
		"telefonicaopencloud_networking_network_v2":  1,
		"telefonicaopencloud_networking_secgroup_v2": 2,
		"telefonicaopencloud_networking_subnet_v2":   3,
		"telefonicaopencloud_s3_bucket_object":       4,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
