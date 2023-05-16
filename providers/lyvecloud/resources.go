package lyvecloud

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "lyvecloud_permission",
			Category:         "Account",
			ShortDescription: `Provides a permission resource.`,
			Description: `

Provides a permission resource. Based on Account API.

~> **NOTE:** Credentials for Account API must be provided to use this resource.

`,
			Keywords: []string{
				"account",
				"permission",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Specifies a unique permission name. The name allows only alphanumeric, '-', '_' or spaces Maximum length can be 128 characters. If omitted, Terraform will assign a random, unique name.`,
				},
				resource.Attribute{
					Name:        "name_prefix",
					Description: `(Optional) Creates a unique permission name beginning with the specified prefix. Conflicts with ` + "`" + `name` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Required) Description of the permission.`,
				},
				resource.Attribute{
					Name:        "actions",
					Description: `(Optional) Actions Enum: “all-operations”, “read-only”, or “write-only”. Must be set if permission is created with ` + "`" + `buckets` + "`" + `, ` + "`" + `bucket_prefix` + "`" + ` or ` + "`" + `all_buckets` + "`" + `. Conflicts with ` + "`" + `policy` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "buckets",
					Description: `(Optional) List (one or more) of existing bucket names. To list one or more existing buckets you can specify [“bucket1”, “bucket2”, and so on]. Conflicts with ` + "`" + `all_buckets` + "`" + `, ` + "`" + `bucket_prefix` + "`" + ` and ` + "`" + `policy` + "`" + `. Required with ` + "`" + `actions` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "all_buckets",
					Description: `(Optional) If set to ` + "`" + `true` + "`" + `, the permission is applied to all the existing and new buckets in the account. Required with ` + "`" + `actions` + "`" + `. Conflicts with ` + "`" + `buckets` + "`" + `, ` + "`" + `bucket_prefix` + "`" + ` and ` + "`" + `policy` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "bucket_prefix",
					Description: `(Optional) Specify the initial name of the bucket as a prefix to apply for permission. Required with ` + "`" + `actions` + "`" + `. Conflicts with ` + "`" + `buckets` + "`" + `, ` + "`" + `all_buckets` + "`" + ` and ` + "`" + `policy` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "policy",
					Description: `(Optional) specify a JSON file path compatible with the AWS IAM policy file or specify the JSON string as shown in the example above. Conflicts with ` + "`" + `buckets` + "`" + `, ` + "`" + `bucket_prefix` + "`" + `, ` + "`" + `all_buckets` + "`" + ` and ` + "`" + `actions` + "`" + `. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `A Permission ID that uniquely identifies each permission created in Lyve Cloud. Can be used to identify this permission when creating a service account.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The permission type: all-buckets/bucket-prefix/bucket-names/policy.`,
				},
				resource.Attribute{
					Name:        "ready_state",
					Description: `True if the permission is ready across all regions. ## Import Permission can be imported using the ` + "`" + `permission` + "`" + `, e.g., ` + "`" + `` + "`" + `` + "`" + ` $ terraform import lyvecloud_permission.permission permission-id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `A Permission ID that uniquely identifies each permission created in Lyve Cloud. Can be used to identify this permission when creating a service account.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The permission type: all-buckets/bucket-prefix/bucket-names/policy.`,
				},
				resource.Attribute{
					Name:        "ready_state",
					Description: `True if the permission is ready across all regions. ## Import Permission can be imported using the ` + "`" + `permission` + "`" + `, e.g., ` + "`" + `` + "`" + `` + "`" + ` $ terraform import lyvecloud_permission.permission permission-id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "lyvecloud_s3_bucket",
			Category:         "S3",
			ShortDescription: `Provides a S3 bucket resource.`,
			Description: `

Provides a S3 bucket resource.

`,
			Keywords: []string{
				"s3",
				"bucket",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "bucket",
					Description: `(Optional, Forces new resource) The name of the bucket. If omitted, Terraform will assign a random, unique name. Must be lowercase and less than or equal to 63 characters in length.`,
				},
				resource.Attribute{
					Name:        "bucket_prefix",
					Description: `(Optional, Forces new resource) Creates a unique bucket name beginning with the specified prefix. Conflicts with ` + "`" + `bucket` + "`" + `. Must be lowercase and less than or equal to 37 characters in length.`,
				},
				resource.Attribute{
					Name:        "force_destroy",
					Description: `(Optional, Default:` + "`" + `false` + "`" + `) A boolean that indicates all objects (including any [locked objects](https://help.lyvecloud.seagate.com/en/using-object-immutability.html)) should be deleted from the bucket so that the bucket can be destroyed without error. These objects are`,
				},
				resource.Attribute{
					Name:        "object_lock_enabled",
					Description: `(Optional, Default:` + "`" + `false` + "`" + `, Forces new resource) Indicates whether this bucket has an Object Lock configuration enabled. Valid values are ` + "`" + `true` + "`" + ` or ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A map of tags to assign to the bucket. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The name of the bucket.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The Lyve Cloud region this bucket resides in.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A map of tags assigned to the resource. ## Timeouts [Configuration options](https://www.terraform.io/docs/configuration/blocks/resources/syntax.html#operation-timeouts): - ` + "`" + `create` + "`" + ` - (Default ` + "`" + `20m` + "`" + `) - ` + "`" + `read` + "`" + ` - (Default ` + "`" + `20m` + "`" + `) - ` + "`" + `update` + "`" + ` - (Default ` + "`" + `20m` + "`" + `) - ` + "`" + `delete` + "`" + ` - (Default ` + "`" + `60m` + "`" + `) ## Import S3 bucket can be imported using the ` + "`" + `bucket` + "`" + `, e.g., ` + "`" + `` + "`" + `` + "`" + ` $ terraform import lyve_s3_bucket.bucket bucket-name ` + "`" + `` + "`" + `` + "`" + ``,
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
				resource.Attribute{
					Name:        "tags",
					Description: `A map of tags assigned to the resource. ## Timeouts [Configuration options](https://www.terraform.io/docs/configuration/blocks/resources/syntax.html#operation-timeouts): - ` + "`" + `create` + "`" + ` - (Default ` + "`" + `20m` + "`" + `) - ` + "`" + `read` + "`" + ` - (Default ` + "`" + `20m` + "`" + `) - ` + "`" + `update` + "`" + ` - (Default ` + "`" + `20m` + "`" + `) - ` + "`" + `delete` + "`" + ` - (Default ` + "`" + `60m` + "`" + `) ## Import S3 bucket can be imported using the ` + "`" + `bucket` + "`" + `, e.g., ` + "`" + `` + "`" + `` + "`" + ` $ terraform import lyve_s3_bucket.bucket bucket-name ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "lyvecloud_s3_bucket_object_lock_configuration",
			Category:         "S3",
			ShortDescription: `Provides an S3 bucket Object Lock configuration resource.`,
			Description: `

Provides an S3 bucket Object Lock configuration resource, known in the Lyve Cloud console as **Object Immutability**. For more information about Object Locking/Object Immutability, go to [Using object immutability](https://help.lyvecloud.seagate.com/en/using-object-immutability.html) in the Lyve Cloud Administrator's Guide.

~> **NOTE:** This resource **does not enable** Object Lock for **new** buckets. It configures a default retention period for objects placed in the specified bucket.
Thus, to **enable** Object Lock for a **new** bucket, see the [Using object lock configuration](s3_bucket.md#Using-object-lock-configuration) section in  the ` + "`" + `lyvecloud_s3_bucket` + "`" + ` resource or the [Object Lock configuration for a new bucket](#object-lock-configuration-for-a-new-bucket) example below.


`,
			Keywords: []string{
				"s3",
				"bucket",
				"object",
				"lock",
				"configuration",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "bucket",
					Description: `(Required, Forces new resource) The name of the bucket.`,
				},
				resource.Attribute{
					Name:        "object_lock_enabled",
					Description: `(Optional, Forces new resource) Indicates whether this bucket has an Object Lock configuration enabled. Defaults to ` + "`" + `Enabled` + "`" + `. Valid values: ` + "`" + `Enabled` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "rule",
					Description: `(Required) Configuration block for specifying the Object Lock rule for the specified object [detailed below](#rule). ### rule The ` + "`" + `rule` + "`" + ` configuration block supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "default_retention",
					Description: `(Required) A configuration block for specifying the default Object Lock retention settings for new objects placed in the specified bucket [detailed below](#default_retention). ### default_retention The ` + "`" + `default_retention` + "`" + ` configuration block supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "days",
					Description: `(Optional, Required if ` + "`" + `years` + "`" + ` is not specified) The number of days that you want to specify for the default retention period.`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `(Required) The default Object Lock retention mode you want to apply to new objects placed in the specified bucket. Valid values: ` + "`" + `COMPLIANCE` + "`" + `, ` + "`" + `GOVERNANCE` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "years",
					Description: `(Optional, Required if ` + "`" + `days` + "`" + ` is not specified) The number of years that you want to specify for the default retention period. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ` + "`" + `bucket` + "`" + `. ## Import S3 bucket Object Lock configuration can be imported using the following example command. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import lyvecloud_s3_bucket_object_lock_configuration.example bucket-name ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ` + "`" + `bucket` + "`" + `. ## Import S3 bucket Object Lock configuration can be imported using the following example command. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import lyvecloud_s3_bucket_object_lock_configuration.example bucket-name ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "lyvecloud_s3_object",
			Category:         "S3",
			ShortDescription: `Provides a S3 object resource.`,
			Description: `

Provides an S3 object resource.

`,
			Keywords: []string{
				"s3",
				"object",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "bucket",
					Description: `(Required) Name of the bucket to put the file in.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) Name of the object once it is in the bucket. The following arguments are optional:`,
				},
				resource.Attribute{
					Name:        "source",
					Description: `(Required) Path to a file that will be read and uploaded as raw bytes for the object content.`,
				},
				resource.Attribute{
					Name:        "cache_control",
					Description: `(Optional) Caching behavior along the request/reply chain Read [w3c cache_control](http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.9) for further details.`,
				},
				resource.Attribute{
					Name:        "content_base64",
					Description: `(Optional, conflicts with ` + "`" + `source` + "`" + ` and ` + "`" + `content` + "`" + `) Base64-encoded data that will be decoded and uploaded as raw bytes for the object content. This allows safely uploading non-UTF8 binary data, but is recommended only for small content such as the result of the ` + "`" + `gzipbase64` + "`" + ` function with small text strings. For larger objects, use ` + "`" + `source` + "`" + ` to stream the content from a disk file.`,
				},
				resource.Attribute{
					Name:        "content_disposition",
					Description: `(Optional) Presentational information for the object. Read [w3c content_disposition](http://www.w3.org/Protocols/rfc2616/rfc2616-sec19.html#sec19.5.1) for further information.`,
				},
				resource.Attribute{
					Name:        "content_encoding",
					Description: `(Optional) Content encodings that have been applied to the object and thus what decoding mechanisms must be applied to obtain the media-type referenced by the Content-Type header field.`,
				},
				resource.Attribute{
					Name:        "content_language",
					Description: `(Optional) Language the content is in e.g., en-US or en-GB.`,
				},
				resource.Attribute{
					Name:        "content_type",
					Description: `(Optional) Standard MIME type describing the format of the object data, e.g., application/octet-stream. All Valid MIME Types are valid for this input.`,
				},
				resource.Attribute{
					Name:        "content",
					Description: `(Optional, conflicts with ` + "`" + `source` + "`" + ` and ` + "`" + `content_base64` + "`" + `) Literal string value to use as the object content, which will be uploaded as UTF-8-encoded text.`,
				},
				resource.Attribute{
					Name:        "etag",
					Description: `(Optional) Triggers updates when the value changes. The only meaningful value is ` + "`" + `filemd5("path/to/file")` + "`" + ` (Terraform 0.11.12 or later) or ` + "`" + `${md5(file("path/to/file"))}` + "`" + ` (Terraform 0.11.11 or earlier).`,
				},
				resource.Attribute{
					Name:        "force_destroy",
					Description: `(Optional) Whether to allow the object to be deleted by removing any legal hold on any object version. Default is ` + "`" + `false` + "`" + `. This value should be set to ` + "`" + `true` + "`" + ` only if the bucket has S3 object lock enabled.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `(Optional) Map of keys/values to provision metadata (will be automatically prefixed by ` + "`" + `x-amz-meta-` + "`" + `, note that only lowercase label are currently supported by the AWS Go API).`,
				},
				resource.Attribute{
					Name:        "object_lock_mode",
					Description: `(Optional) Object lock retention mode that you want to apply to this object. Valid values are ` + "`" + `GOVERNANCE` + "`" + ` and ` + "`" + `COMPLIANCE` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "object_lock_retain_until_date",
					Description: `(Optional) Date and time, in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8), when this object's object lock will expire.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Map of tags to assign to the object. If no content is provided through ` + "`" + `source` + "`" + `, ` + "`" + `content` + "`" + ` or ` + "`" + `content_base64` + "`" + `, then the object will be empty.`,
				},
				resource.Attribute{
					Name:        "etag",
					Description: `ETag generated for the object.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `` + "`" + `key` + "`" + ` of the resource supplied above.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Map of tags assigned to the resource.`,
				},
				resource.Attribute{
					Name:        "version_id",
					Description: `Unique version ID value for the object, if bucket versioning is enabled. ## Import Objects can be imported using the ` + "`" + `id` + "`" + `. The ` + "`" + `id` + "`" + ` is the bucket name and the key together e.g., ` + "`" + `` + "`" + `` + "`" + ` $ terraform import lyvecloud_s3_object.object some-bucket-name/some/key.txt ` + "`" + `` + "`" + `` + "`" + ` Additionally, s3 url syntax can be used, e.g., ` + "`" + `` + "`" + `` + "`" + ` $ terraform import lyve_s3_object.object s3://some-bucket-name/some/key.txt ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "etag",
					Description: `ETag generated for the object.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `` + "`" + `key` + "`" + ` of the resource supplied above.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Map of tags assigned to the resource.`,
				},
				resource.Attribute{
					Name:        "version_id",
					Description: `Unique version ID value for the object, if bucket versioning is enabled. ## Import Objects can be imported using the ` + "`" + `id` + "`" + `. The ` + "`" + `id` + "`" + ` is the bucket name and the key together e.g., ` + "`" + `` + "`" + `` + "`" + ` $ terraform import lyvecloud_s3_object.object some-bucket-name/some/key.txt ` + "`" + `` + "`" + `` + "`" + ` Additionally, s3 url syntax can be used, e.g., ` + "`" + `` + "`" + `` + "`" + ` $ terraform import lyve_s3_object.object s3://some-bucket-name/some/key.txt ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "lyvecloud_s3_object_copy",
			Category:         "S3",
			ShortDescription: `Provides a resource for copying an S3 object.`,
			Description: `

Provides a resource for copying an S3 object.

`,
			Keywords: []string{
				"s3",
				"object",
				"copy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "bucket",
					Description: `(Required) Name of the bucket to put the file in.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) Name of the object once it is in the bucket.`,
				},
				resource.Attribute{
					Name:        "source",
					Description: `(Required) Specifies the source object for the copy operation. You specify the value in one of two formats. For objects not accessed through an access point, specify the name of the source bucket and the key of the source object, separated by a slash (` + "`" + `/` + "`" + `). For example, ` + "`" + `testbucket/test1.json` + "`" + `. The following arguments are optional:`,
				},
				resource.Attribute{
					Name:        "cache_control",
					Description: `(Optional) Specifies caching behavior along the request/reply chain Read [w3c cache_control](http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.9) for further details.`,
				},
				resource.Attribute{
					Name:        "content_disposition",
					Description: `(Optional) Specifies presentational information for the object. Read [w3c content_disposition](http://www.w3.org/Protocols/rfc2616/rfc2616-sec19.html#sec19.5.1) for further information.`,
				},
				resource.Attribute{
					Name:        "content_encoding",
					Description: `(Optional) Specifies what content encodings have been applied to the object and thus what decoding mechanisms must be applied to obtain the media-type referenced by the Content-Type header field. Read [w3c content encoding](http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.11) for further information.`,
				},
				resource.Attribute{
					Name:        "content_language",
					Description: `(Optional) Language the content is in e.g., en-US or en-GB.`,
				},
				resource.Attribute{
					Name:        "content_type",
					Description: `(Optional) Standard MIME type describing the format of the object data, e.g., ` + "`" + `application/octet-stream` + "`" + `. All Valid MIME Types are valid for this input.`,
				},
				resource.Attribute{
					Name:        "copy_if_match",
					Description: `(Optional) Copies the object if its entity tag (ETag) matches the specified tag.`,
				},
				resource.Attribute{
					Name:        "copy_if_modified_since",
					Description: `(Optional) Copies the object if it has been modified since the specified time, in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8).`,
				},
				resource.Attribute{
					Name:        "copy_if_none_match",
					Description: `(Optional) Copies the object if its entity tag (ETag) is different than the specified ETag.`,
				},
				resource.Attribute{
					Name:        "copy_if_unmodified_since",
					Description: `(Optional) Copies the object if it hasn't been modified since the specified time, in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8).`,
				},
				resource.Attribute{
					Name:        "force_destroy",
					Description: `(Optional) Allow the object to be deleted by removing any legal hold on any object version. Default is ` + "`" + `false` + "`" + `. This value should be set to ` + "`" + `true` + "`" + ` only if the bucket has S3 object lock enabled.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `(Optional) A map of keys/values to provision metadata (will be automatically prefixed by ` + "`" + `x-amz-meta-` + "`" + `, note that only lowercase label are currently supported by the AWS Go API).`,
				},
				resource.Attribute{
					Name:        "metadata_directive",
					Description: `(Optional) Specifies whether the metadata is copied from the source object or replaced with metadata provided in the request. Valid values are ` + "`" + `COPY` + "`" + ` and ` + "`" + `REPLACE` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "tagging_directive",
					Description: `(Optional) Specifies whether the object tag-set are copied from the source object or replaced with tag-set provided in the request. Valid values are ` + "`" + `COPY` + "`" + ` and ` + "`" + `REPLACE` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "object_lock_mode",
					Description: `(Optional) The object lock [retention mode](https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock-overview.html#object-lock-retention-modes) that you want to apply to this object. Valid values are ` + "`" + `GOVERNANCE` + "`" + ` and ` + "`" + `COMPLIANCE` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "object_lock_retain_until_date",
					Description: `(Optional) The date and time, in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8), when this object's object lock will [expire](https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock-overview.html#object-lock-retention-periods).`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A map of tags to assign to the object.`,
				},
				resource.Attribute{
					Name:        "etag",
					Description: `The ETag generated for the object (an MD5 sum of the object content).`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ` + "`" + `key` + "`" + ` of the resource supplied above.`,
				},
				resource.Attribute{
					Name:        "last_modified",
					Description: `Returns the date that the object was last modified, in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8).`,
				},
				resource.Attribute{
					Name:        "source_version_id",
					Description: `Version of the copied object in the source bucket.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A map of tags assigned to the resource.`,
				},
				resource.Attribute{
					Name:        "version_id",
					Description: `Version ID of the newly created copy.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "etag",
					Description: `The ETag generated for the object (an MD5 sum of the object content).`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ` + "`" + `key` + "`" + ` of the resource supplied above.`,
				},
				resource.Attribute{
					Name:        "last_modified",
					Description: `Returns the date that the object was last modified, in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8).`,
				},
				resource.Attribute{
					Name:        "source_version_id",
					Description: `Version of the copied object in the source bucket.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A map of tags assigned to the resource.`,
				},
				resource.Attribute{
					Name:        "version_id",
					Description: `Version ID of the newly created copy.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "lyvecloud_service_account",
			Category:         "Account",
			ShortDescription: `Provides a service account resource.`,
			Description: `

Provides a service account resource. Based on Account API.

~> **NOTE:** Credentials for Account API must be provided to use this resource.

`,
			Keywords: []string{
				"account",
				"service",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Specifies the unique Service Account name. The name allows only alphanumeric, '-', '_' or space.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the Service Account.`,
				},
				resource.Attribute{
					Name:        "permissions",
					Description: `(Required) Specify (one or more) unique values of permission-id. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `A Service Account ID that uniquely identifies each Service Account created in Lyve Cloud. Used to identify this Service Account when it is deleted.`,
				},
				resource.Attribute{
					Name:        "access_key",
					Description: `Access key to use when authenticating S3 API requests.`,
				},
				resource.Attribute{
					Name:        "secret",
					Description: `Access secret key to use when authenticating S3 API requests.`,
				},
				resource.Attribute{
					Name:        "ready_state",
					Description: `True if the service account is ready across all regions.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `State of the Service Account. It can be enabled or disabled. ## Import Service Account can be imported using the ` + "`" + `service account` + "`" + `, e.g., ` + "`" + `` + "`" + `` + "`" + ` $ terraform import lyvecloud_servcie_account.servcie-account servcie-account-id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `A Service Account ID that uniquely identifies each Service Account created in Lyve Cloud. Used to identify this Service Account when it is deleted.`,
				},
				resource.Attribute{
					Name:        "access_key",
					Description: `Access key to use when authenticating S3 API requests.`,
				},
				resource.Attribute{
					Name:        "secret",
					Description: `Access secret key to use when authenticating S3 API requests.`,
				},
				resource.Attribute{
					Name:        "ready_state",
					Description: `True if the service account is ready across all regions.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `State of the Service Account. It can be enabled or disabled. ## Import Service Account can be imported using the ` + "`" + `service account` + "`" + `, e.g., ` + "`" + `` + "`" + `` + "`" + ` $ terraform import lyvecloud_servcie_account.servcie-account servcie-account-id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
	}

	resourcesMap = map[string]int{

		"lyvecloud_permission":                          0,
		"lyvecloud_s3_bucket":                           1,
		"lyvecloud_s3_bucket_object_lock_configuration": 2,
		"lyvecloud_s3_object":                           3,
		"lyvecloud_s3_object_copy":                      4,
		"lyvecloud_service_account":                     5,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
