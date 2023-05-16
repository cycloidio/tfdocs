package lyvecloud

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "lyvecloud_s3_bucket",
			Category:         "S3",
			ShortDescription: `Provides details about a specific S3 bucket`,
			Description:      ``,
			Keywords: []string{
				"s3",
				"bucket",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "bucket",
					Description: `(Required) The name of the bucket. ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The name of the bucket.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The Lyve Cloud region this bucket resides in.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The name of the bucket.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The Lyve Cloud region this bucket resides in.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "lyvecloud_s3_object",
			Category:         "S3",
			ShortDescription: `Provides metadata and optionally content of an S3 object`,
			Description:      ``,
			Keywords: []string{
				"s3",
				"object",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "bucket",
					Description: `(Required) The name of the bucket to read the object from.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) The full path to the object inside the bucket.`,
				},
				resource.Attribute{
					Name:        "version_id",
					Description: `(Optional) Specific version ID of the object returned (defaults to latest version). ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
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
					Description: `[ETag](https://en.wikipedia.org/wiki/HTTP_ETag) generated for the object (an MD5 sum of the object content in case it's not encrypted).`,
				},
				resource.Attribute{
					Name:        "last_modified",
					Description: `Last modified date of the object in RFC1123 format (e.g., ` + "`" + `Mon, 02 Jan 2006 15:04:05 MST` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `A map of metadata stored with the object in S3.`,
				},
				resource.Attribute{
					Name:        "object_lock_mode",
					Description: `The object lock retention mode currently in place for this object.`,
				},
				resource.Attribute{
					Name:        "object_lock_retain_until_date",
					Description: `The date and time when this object's object lock will expire.`,
				},
				resource.Attribute{
					Name:        "version_id",
					Description: `The latest version ID of the object returned.`,
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
					Description: `[ETag](https://en.wikipedia.org/wiki/HTTP_ETag) generated for the object (an MD5 sum of the object content in case it's not encrypted).`,
				},
				resource.Attribute{
					Name:        "last_modified",
					Description: `Last modified date of the object in RFC1123 format (e.g., ` + "`" + `Mon, 02 Jan 2006 15:04:05 MST` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `A map of metadata stored with the object in S3.`,
				},
				resource.Attribute{
					Name:        "object_lock_mode",
					Description: `The object lock retention mode currently in place for this object.`,
				},
				resource.Attribute{
					Name:        "object_lock_retain_until_date",
					Description: `The date and time when this object's object lock will expire.`,
				},
				resource.Attribute{
					Name:        "version_id",
					Description: `The latest version ID of the object returned.`,
				},
			},
		},
	}

	dataSourcesMap = map[string]int{

		"lyvecloud_s3_bucket": 0,
		"lyvecloud_s3_object": 1,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
