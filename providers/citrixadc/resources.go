package citrixadc

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "citrixadc_appfwfieldtype",
			Category:         "Application Firewall",
			ShortDescription: ``,
			Description: `

The ` + "`" + `appfwfieldtype` + "`" + ` resource is used to create Application Firwall FieldType resource.

`,
			Keywords: []string{
				"application",
				"firewall",
				"appfwfieldtype",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `Name for the field type. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after the field type is added. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my field type" or 'my field type').`,
				},
				resource.Attribute{
					Name:        "regex",
					Description: `(Optional) PCRE - format regular expression defining the characters and length allowed for this field type.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `(Optional) Positive integer specifying the priority of the field type. A lower number specifies a higher priority. Field types are checked in the order of their priority numbers.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `(Optional) Comment describing the type of field that this field type is intended to match.`,
				},
				resource.Attribute{
					Name:        "nocharmaps",
					Description: `(Optional) will not show internal field types added as part of FieldFormat learn rules deployment. ## Attribute Reference In addition to the arguments, the following attributes are available:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of the ` + "`" + `appfwfieldtype` + "`" + `. It has the same value as the ` + "`" + `name` + "`" + ` attribute.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id of the ` + "`" + `appfwfieldtype` + "`" + `. It has the same value as the ` + "`" + `name` + "`" + ` attribute.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "citrixadc_appfwjsoncontenttype",
			Category:         "Application Firewall",
			ShortDescription: ``,
			Description: `

The ` + "`" + `appfwjsoncontenttype` + "`" + ` resource is used to create Application Firewall ContentType resource.

`,
			Keywords: []string{
				"application",
				"firewall",
				"appfwjsoncontenttype",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "jsoncontenttypevalue",
					Description: `Content type to be classified as JSON.`,
				},
				resource.Attribute{
					Name:        "isregex",
					Description: `(Optional) Is json content type a regular expression?. Possible values: [ REGEX, NOTREGEX ] ## Attribute Reference In addition to the arguments, the following attributes are available:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of the ` + "`" + `appfwjsoncontenttype` + "`" + `. It has the same value as the ` + "`" + `jsoncontenttypevalue` + "`" + ` attribute.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id of the ` + "`" + `appfwjsoncontenttype` + "`" + `. It has the same value as the ` + "`" + `jsoncontenttypevalue` + "`" + ` attribute.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "citrixadc_appfwpolicy",
			Category:         "Application Firewall",
			ShortDescription: ``,
			Description: `

The ` + "`" + `appfwpolicy` + "`" + ` resource is used to create Applicatin Firewall Policy resource.

`,
			Keywords: []string{
				"application",
				"firewall",
				"appfwpolicy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `Name for the policy. Must begin with a letter, number, or the underscore character \(_\), and must contain only letters, numbers, and the hyphen \(-\), period \(.\) pound \(\#\), space \( \), at (@), equals \(=\), colon \(:\), and underscore characters. Can be changed after the policy is created. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks \(for example, "my policy" or 'my policy'\).`,
				},
				resource.Attribute{
					Name:        "rule",
					Description: `(Optional) Name of the Citrix ADC named rule, or a Citrix ADC expression, that the policy uses to determine whether to filter the connection through the application firewall with the designated profile.`,
				},
				resource.Attribute{
					Name:        "profilename",
					Description: `(Optional) Name of the application firewall profile to use if the policy matches.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `(Optional) Any comments to preserve information about the policy for later reference.`,
				},
				resource.Attribute{
					Name:        "logaction",
					Description: `(Optional) Where to log information for connections that match this policy.`,
				},
				resource.Attribute{
					Name:        "newname",
					Description: `(Optional) New name for the policy. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my policy" or 'my policy'). ## Attribute Reference In addition to the arguments, the following attributes are available:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of the ` + "`" + `appfwpolicy` + "`" + `. It has the same value as the ` + "`" + `name` + "`" + ` attribute.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id of the ` + "`" + `appfwpolicy` + "`" + `. It has the same value as the ` + "`" + `name` + "`" + ` attribute.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "citrixadc_appfwpolicylabel",
			Category:         "Application Firewall",
			ShortDescription: ``,
			Description: `

The ` + "`" + `appfwpolicylabel` + "`" + ` resource is used to create Application Firewall Policy Label resource.

`,
			Keywords: []string{
				"application",
				"firewall",
				"appfwpolicylabel",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "labelname",
					Description: `Name for the policy label. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Can be changed after the policy label is created. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my policy label" or 'my policy label').`,
				},
				resource.Attribute{
					Name:        "policylabeltype",
					Description: `(Optional) Type of transformations allowed by the policies bound to the label. Always http_req for application firewall policy labels. Possible values: [ http_req ] ## Attribute Reference In addition to the arguments, the following attributes are available:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of the ` + "`" + `appfwpolicylabel` + "`" + `. It has the same value as the ` + "`" + `labelname` + "`" + ` attribute.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id of the ` + "`" + `appfwpolicylabel` + "`" + `. It has the same value as the ` + "`" + `labelname` + "`" + ` attribute.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "citrixadc_appfwprofile",
			Category:         "Application Firewall",
			ShortDescription: ``,
			Description: `

The ` + "`" + `appfwprofile` + "`" + ` resource is used to create Applicatin Firewall Profile.

`,
			Keywords: []string{
				"application",
				"firewall",
				"appfwprofile",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `Name for the profile. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.), pound (#), space ( ), at (@), equals (=), colon (:), and underscore (_) characters. Cannot be changed after the profile is added. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my profile" or 'my profile').`,
				},
				resource.Attribute{
					Name:        "defaults",
					Description: `(Optional) Default configuration to apply to the profile. Basic defaults are intended for standard content that requires little further configuration, such as static web site content. Advanced defaults are intended for specialized content that requires significant specialized configuration, such as heavily scripted or dynamic content. CLI users: When adding an application firewall profile, you can set either the defaults or the type, but not both. To set both options, create the profile by using the add appfw profile command, and then use the set appfw profile command to configure the other option. Possible values: [ basic, advanced ]`,
				},
				resource.Attribute{
					Name:        "starturlaction",
					Description: `(Optional) One or more Start URL actions. Available settings function as follows:`,
				},
				resource.Attribute{
					Name:        "contenttypeaction",
					Description: `(Optional) One or more Content-type actions. Available settings function as follows:`,
				},
				resource.Attribute{
					Name:        "inspectcontenttypes",
					Description: `(Optional) One or more InspectContentType lists.`,
				},
				resource.Attribute{
					Name:        "starturlclosure",
					Description: `(Optional) Toggle the state of Start URL Closure. Possible values: [ on, off ]`,
				},
				resource.Attribute{
					Name:        "denyurlaction",
					Description: `(Optional) One or more Deny URL actions. Available settings function as follows:`,
				},
				resource.Attribute{
					Name:        "refererheadercheck",
					Description: `(Optional) Enable validation of Referer headers. Referer validation ensures that a web form that a user sends to your web site originally came from your web site, not an outside attacker. Although this parameter is part of the Start URL check, referer validation protects against cross-site request forgery (CSRF) attacks, not Start URL attacks. Possible values: [ OFF, if_present, AlwaysExceptStartURLs, AlwaysExceptFirstRequest ]`,
				},
				resource.Attribute{
					Name:        "cookieconsistencyaction",
					Description: `(Optional) One or more Cookie Consistency actions. Available settings function as follows:`,
				},
				resource.Attribute{
					Name:        "cookietransforms",
					Description: `(Optional) Perform the specified type of cookie transformation. Available settings function as follows:`,
				},
				resource.Attribute{
					Name:        "cookieencryption",
					Description: `(Optional) Type of cookie encryption. Available settings function as follows:`,
				},
				resource.Attribute{
					Name:        "cookieproxying",
					Description: `(Optional) Cookie proxy setting. Available settings function as follows:`,
				},
				resource.Attribute{
					Name:        "addcookieflags",
					Description: `(Optional) Add the specified flags to cookies. Available settings function as follows:`,
				},
				resource.Attribute{
					Name:        "fieldconsistencyaction",
					Description: `(Optional) One or more Form Field Consistency actions. Available settings function as follows:`,
				},
				resource.Attribute{
					Name:        "csrftagaction",
					Description: `(Optional) One or more Cross-Site Request Forgery (CSRF) Tagging actions. Available settings function as follows:`,
				},
				resource.Attribute{
					Name:        "crosssitescriptingaction",
					Description: `(Optional) One or more Cross-Site Scripting (XSS) actions. Available settings function as follows:`,
				},
				resource.Attribute{
					Name:        "crosssitescriptingtransformunsafehtml",
					Description: `(Optional) Transform cross-site scripts. This setting configures the application firewall to disable dangerous HTML instead of blocking the request. CAUTION: Make sure that this parameter is set to ON if you are configuring any cross-site scripting transformations. If it is set to OFF, no cross-site scripting transformations are performed regardless of any other settings. Possible values: [ on, off ]`,
				},
				resource.Attribute{
					Name:        "crosssitescriptingcheckcompleteurls",
					Description: `(Optional) Check complete URLs for cross-site scripts, instead of just the query portions of URLs. Possible values: [ on, off ]`,
				},
				resource.Attribute{
					Name:        "sqlinjectionaction",
					Description: `(Optional) One or more HTML SQL Injection actions. Available settings function as follows:`,
				},
				resource.Attribute{
					Name:        "sqlinjectiontransformspecialchars",
					Description: `(Optional) Transform injected SQL code. This setting configures the application firewall to disable SQL special strings instead of blocking the request. Since most SQL servers require a special string to activate an SQL keyword, in most cases a request that contains injected SQL code is safe if special strings are disabled. CAUTION: Make sure that this parameter is set to ON if you are configuring any SQL injection transformations. If it is set to OFF, no SQL injection transformations are performed regardless of any other settings. Possible values: [ on, off ]`,
				},
				resource.Attribute{
					Name:        "sqlinjectiononlycheckfieldswithsqlchars",
					Description: `(Optional) Check only form fields that contain SQL special strings (characters) for injected SQL code. Most SQL servers require a special string to activate an SQL request, so SQL code without a special string is harmless to most SQL servers. Possible values: [ on, off ]`,
				},
				resource.Attribute{
					Name:        "sqlinjectiontype",
					Description: `(Optional) Available SQL injection types. -SQLSplChar : Checks for SQL Special Chars -SQLKeyword : Checks for SQL Keywords -SQLSplCharANDKeyword : Checks for both and blocks if both are found -SQLSplCharORKeyword : Checks for both and blocks if anyone is found. Possible values: [ SQLSplChar, SQLKeyword, SQLSplCharORKeyword, SQLSplCharANDKeyword ]`,
				},
				resource.Attribute{
					Name:        "sqlinjectionchecksqlwildchars",
					Description: `(Optional) Check for form fields that contain SQL wild chars . Possible values: [ on, off ]`,
				},
				resource.Attribute{
					Name:        "fieldformataction",
					Description: `(Optional) One or more Field Format actions. Available settings function as follows:`,
				},
				resource.Attribute{
					Name:        "defaultfieldformattype",
					Description: `(Optional) Designate a default field type to be applied to web form fields that do not have a field type explicitly assigned to them.`,
				},
				resource.Attribute{
					Name:        "defaultfieldformatminlength",
					Description: `(Optional) To disable the minimum and maximum length settings and allow data of any length to be entered into the field, set this parameter to zero (0).`,
				},
				resource.Attribute{
					Name:        "defaultfieldformatmaxlength",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "bufferoverflowaction",
					Description: `(Optional) One or more Buffer Overflow actions. Available settings function as follows:`,
				},
				resource.Attribute{
					Name:        "bufferoverflowmaxurllength",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "bufferoverflowmaxheaderlength",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "bufferoverflowmaxcookielength",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "creditcardaction",
					Description: `(Optional) One or more Credit Card actions. Available settings function as follows:`,
				},
				resource.Attribute{
					Name:        "creditcard",
					Description: `(Optional) Credit card types that the application firewall should protect. Possible values: [ none, visa, mastercard, discover, amex, jcb, dinersclub ]`,
				},
				resource.Attribute{
					Name:        "creditcardmaxallowed",
					Description: `(Optional) This parameter value is used by the block action. It represents the maximum number of credit card numbers that can appear on a web page served by your protected web sites. Pages that contain more credit card numbers are blocked.`,
				},
				resource.Attribute{
					Name:        "creditcardxout",
					Description: `(Optional) Mask any credit card number detected in a response by replacing each digit, except the digits in the final group, with the letter "X.". Possible values: [ on, off ]`,
				},
				resource.Attribute{
					Name:        "dosecurecreditcardlogging",
					Description: `(Optional) Setting this option logs credit card numbers in the response when the match is found. Possible values: [ on, off ]`,
				},
				resource.Attribute{
					Name:        "streaming",
					Description: `(Optional) Setting this option converts content-length form submission requests (requests with content-type "application/x-www-form-urlencoded" or "multipart/form-data") to chunked requests when atleast one of the following protections : SQL injection protection, XSS protection, form field consistency protection, starturl closure, CSRF tagging is enabled. Please make sure that the backend server accepts chunked requests before enabling this option. Possible values: [ on, off ]`,
				},
				resource.Attribute{
					Name:        "trace",
					Description: `(Optional) Toggle the state of trace. Possible values: [ on, off ]`,
				},
				resource.Attribute{
					Name:        "requestcontenttype",
					Description: `(Optional) Default Content-Type header for requests. A Content-Type header can contain 0-255 letters, numbers, and the hyphen (-) and underscore (_) characters.`,
				},
				resource.Attribute{
					Name:        "responsecontenttype",
					Description: `(Optional) Default Content-Type header for responses. A Content-Type header can contain 0-255 letters, numbers, and the hyphen (-) and underscore (_) characters.`,
				},
				resource.Attribute{
					Name:        "jsonerrorobject",
					Description: `(Optional) Name to the imported JSON Error Object to be set on application firewall profile. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks \(for example, "my JSON error object" or 'my JSON error object'\).`,
				},
				resource.Attribute{
					Name:        "jsondosaction",
					Description: `(Optional) One or more JSON Denial-of-Service (JsonDoS) actions. Available settings function as follows:`,
				},
				resource.Attribute{
					Name:        "jsonsqlinjectionaction",
					Description: `(Optional) One or more JSON SQL Injection actions. Available settings function as follows:`,
				},
				resource.Attribute{
					Name:        "jsonsqlinjectiontype",
					Description: `(Optional) Available SQL injection types. -SQLSplChar : Checks for SQL Special Chars -SQLKeyword : Checks for SQL Keywords -SQLSplCharANDKeyword : Checks for both and blocks if both are found -SQLSplCharORKeyword : Checks for both and blocks if anyone is found. Possible values: [ SQLSplChar, SQLKeyword, SQLSplCharORKeyword, SQLSplCharANDKeyword ]`,
				},
				resource.Attribute{
					Name:        "jsonxssaction",
					Description: `(Optional) One or more JSON Cross-Site Scripting actions. Available settings function as follows:`,
				},
				resource.Attribute{
					Name:        "xmldosaction",
					Description: `(Optional) One or more XML Denial-of-Service (XDoS) actions. Available settings function as follows:`,
				},
				resource.Attribute{
					Name:        "xmlformataction",
					Description: `(Optional) One or more XML Format actions. Available settings function as follows:`,
				},
				resource.Attribute{
					Name:        "xmlsqlinjectionaction",
					Description: `(Optional) One or more XML SQL Injection actions. Available settings function as follows:`,
				},
				resource.Attribute{
					Name:        "xmlsqlinjectiononlycheckfieldswithsqlchars",
					Description: `(Optional) Check only form fields that contain SQL special characters, which most SQL servers require before accepting an SQL command, for injected SQL. Possible values: [ on, off ]`,
				},
				resource.Attribute{
					Name:        "xmlsqlinjectiontype",
					Description: `(Optional) Available SQL injection types. -SQLSplChar : Checks for SQL Special Chars -SQLKeyword : Checks for SQL Keywords -SQLSplCharANDKeyword : Checks for both and blocks if both are found -SQLSplCharORKeyword : Checks for both and blocks if anyone is found. Possible values: [ SQLSplChar, SQLKeyword, SQLSplCharORKeyword, SQLSplCharANDKeyword ]`,
				},
				resource.Attribute{
					Name:        "xmlsqlinjectionchecksqlwildchars",
					Description: `(Optional) Check for form fields that contain SQL wild chars . Possible values: [ on, off ]`,
				},
				resource.Attribute{
					Name:        "xmlsqlinjectionparsecomments",
					Description: `(Optional) Parse comments in XML Data and exempt those sections of the request that are from the XML SQL Injection check. You must configure the type of comments that the application firewall is to detect and exempt from this security check. Available settings function as follows:`,
				},
				resource.Attribute{
					Name:        "xmlxssaction",
					Description: `(Optional) One or more XML Cross-Site Scripting actions. Available settings function as follows:`,
				},
				resource.Attribute{
					Name:        "xmlwsiaction",
					Description: `(Optional) One or more Web Services Interoperability (WSI) actions. Available settings function as follows:`,
				},
				resource.Attribute{
					Name:        "xmlattachmentaction",
					Description: `(Optional) One or more XML Attachment actions. Available settings function as follows:`,
				},
				resource.Attribute{
					Name:        "xmlvalidationaction",
					Description: `(Optional) One or more XML Validation actions. Available settings function as follows:`,
				},
				resource.Attribute{
					Name:        "xmlerrorobject",
					Description: `(Optional) Name to assign to the XML Error Object, which the application firewall displays when a user request is blocked. Must begin with a letter, number, or the underscore character \(_\), and must contain only letters, numbers, and the hyphen \(-\), period \(.\) pound \(\#\), space \( \), at (@), equals \(=\), colon \(:\), and underscore characters. Cannot be changed after the XML error object is added. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks \(for example, "my XML error object" or 'my XML error object'\).`,
				},
				resource.Attribute{
					Name:        "customsettings",
					Description: `(Optional) Object name for custom settings. This check is applicable to Profile Type: HTML, XML. .`,
				},
				resource.Attribute{
					Name:        "signatures",
					Description: `(Optional) Object name for signatures. This check is applicable to Profile Type: HTML, XML. .`,
				},
				resource.Attribute{
					Name:        "xmlsoapfaultaction",
					Description: `(Optional) One or more XML SOAP Fault Filtering actions. Available settings function as follows:`,
				},
				resource.Attribute{
					Name:        "usehtmlerrorobject",
					Description: `(Optional) Send an imported HTML Error object to a user when a request is blocked, instead of redirecting the user to the designated Error URL. Possible values: [ on, off ]`,
				},
				resource.Attribute{
					Name:        "errorurl",
					Description: `(Optional) URL that application firewall uses as the Error URL.`,
				},
				resource.Attribute{
					Name:        "htmlerrorobject",
					Description: `(Optional) Name to assign to the HTML Error Object. Must begin with a letter, number, or the underscore character \(_\), and must contain only letters, numbers, and the hyphen \(-\), period \(.\) pound \(\#\), space \( \), at (@), equals \(=\), colon \(:\), and underscore characters. Cannot be changed after the HTML error object is added. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks \(for example, "my HTML error object" or 'my HTML error object'\).`,
				},
				resource.Attribute{
					Name:        "logeverypolicyhit",
					Description: `(Optional) Log every profile match, regardless of security checks results. Possible values: [ on, off ]`,
				},
				resource.Attribute{
					Name:        "stripcomments",
					Description: `(Optional) Strip HTML comments. This check is applicable to Profile Type: HTML. . Possible values: [ on, off ]`,
				},
				resource.Attribute{
					Name:        "striphtmlcomments",
					Description: `(Optional) Strip HTML comments before forwarding a web page sent by a protected web site in response to a user request. Possible values: [ none, all, exclude_script_tag ]`,
				},
				resource.Attribute{
					Name:        "stripxmlcomments",
					Description: `(Optional) Strip XML comments before forwarding a web page sent by a protected web site in response to a user request. Possible values: [ none, all ]`,
				},
				resource.Attribute{
					Name:        "exemptclosureurlsfromsecuritychecks",
					Description: `(Optional) Exempt URLs that pass the Start URL closure check from SQL injection, cross-site script, field format and field consistency security checks at locations other than headers. Possible values: [ on, off ]`,
				},
				resource.Attribute{
					Name:        "defaultcharset",
					Description: `(Optional) Default character set for protected web pages. Web pages sent by your protected web sites in response to user requests are assigned this character set if the page does not already specify a character set. The character sets supported by the application firewall are:`,
				},
				resource.Attribute{
					Name:        "dynamiclearning",
					Description: `(Optional) One or more security checks. Available options are as follows:`,
				},
				resource.Attribute{
					Name:        "postbodylimit",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "postbodylimitsignature",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "fileuploadmaxnum",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "canonicalizehtmlresponse",
					Description: `(Optional) Perform HTML entity encoding for any special characters in responses sent by your protected web sites. Possible values: [ on, off ]`,
				},
				resource.Attribute{
					Name:        "enableformtagging",
					Description: `(Optional) Enable tagging of web form fields for use by the Form Field Consistency and CSRF Form Tagging checks. Possible values: [ on, off ]`,
				},
				resource.Attribute{
					Name:        "sessionlessfieldconsistency",
					Description: `(Optional) Perform sessionless Field Consistency Checks. Possible values: [ OFF, ON, postOnly ]`,
				},
				resource.Attribute{
					Name:        "sessionlessurlclosure",
					Description: `(Optional) Enable session less URL Closure Checks. This check is applicable to Profile Type: HTML. . Possible values: [ on, off ]`,
				},
				resource.Attribute{
					Name:        "semicolonfieldseparator",
					Description: `(Optional) Allow '; ' as a form field separator in URL queries and POST form bodies. . Possible values: [ on, off ]`,
				},
				resource.Attribute{
					Name:        "excludefileuploadfromchecks",
					Description: `(Optional) Exclude uploaded files from Form checks. Possible values: [ on, off ]`,
				},
				resource.Attribute{
					Name:        "sqlinjectionparsecomments",
					Description: `(Optional) Parse HTML comments and exempt them from the HTML SQL Injection check. You must specify the type of comments that the application firewall is to detect and exempt from this security check. Available settings function as follows:`,
				},
				resource.Attribute{
					Name:        "invalidpercenthandling",
					Description: `(Optional) Configure the method that the application firewall uses to handle percent-encoded names and values. Available settings function as follows:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) Application firewall profile type, which controls which security checks and settings are applied to content that is filtered with the profile. Available settings function as follows:`,
				},
				resource.Attribute{
					Name:        "checkrequestheaders",
					Description: `(Optional) Check request headers as well as web forms for injected SQL and cross-site scripts. Possible values: [ on, off ]`,
				},
				resource.Attribute{
					Name:        "optimizepartialreqs",
					Description: `(Optional) Optimize handle of HTTP partial requests i.e. those with range headers. Available settings are as follows:`,
				},
				resource.Attribute{
					Name:        "urldecoderequestcookies",
					Description: `(Optional) URL Decode request cookies before subjecting them to SQL and cross-site scripting checks. Possible values: [ on, off ]`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `(Optional) Any comments about the purpose of profile, or other useful information about the profile.`,
				},
				resource.Attribute{
					Name:        "percentdecoderecursively",
					Description: `(Optional) Configure whether the application firewall should use percentage recursive decoding. Possible values: [ on, off ]`,
				},
				resource.Attribute{
					Name:        "multipleheaderaction",
					Description: `(Optional) One or more multiple header actions. Available settings function as follows:`,
				},
				resource.Attribute{
					Name:        "rfcprofile",
					Description: `(Optional) Object name of the rfc profile.`,
				},
				resource.Attribute{
					Name:        "fileuploadtypesaction",
					Description: `(Optional) One or more file upload types actions. Available settings function as follows:`,
				},
				resource.Attribute{
					Name:        "verboseloglevel",
					Description: `(Optional) Detailed Logging Verbose Log Level. Possible values: [ pattern, patternPayload, patternPayloadHeader ]`,
				},
				resource.Attribute{
					Name:        "archivename",
					Description: `(Optional) Source for tar archive. ## Attribute Reference In addition to the arguments, the following attributes are available:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of the ` + "`" + `appfwprofile` + "`" + `. It has the same value as the ` + "`" + `name` + "`" + ` attribute.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id of the ` + "`" + `appfwprofile` + "`" + `. It has the same value as the ` + "`" + `name` + "`" + ` attribute.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "citrixadc_appfwprofile_cookieconsistency_binding",
			Category:         "Application Firewall",
			ShortDescription: ``,
			Description: `

The ` + "`" + `appfwprofile_cookieconsistency_binding` + "`" + ` resource is used to add bindings between appfwprofile and cookieconsistency relaxation rules.

`,
			Keywords: []string{
				"application",
				"firewall",
				"appfwprofile",
				"cookieconsistency",
				"binding",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `Name of the profile to which to bind an exemption or rule.`,
				},
				resource.Attribute{
					Name:        "cookieconsistency",
					Description: `The name of the cookie to be checked.`,
				},
				resource.Attribute{
					Name:        "isregex",
					Description: `(Optional) Is the cookie name a regular expression?. Possible values: [ REGEX, NOTREGEX ]`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `(Optional) Enabled. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `(Optional) Any comments about the purpose of profile, or other useful information about the profile.`,
				},
				resource.Attribute{
					Name:        "isautodeployed",
					Description: `(Optional) Is the rule auto deployed by dynamic profile ?. Possible values: [ AUTODEPLOYED, NOTAUTODEPLOYED ]`,
				},
				resource.Attribute{
					Name:        "alertonly",
					Description: `(Optional) Send SNMP alert?. Possible values: [ on, off ] ## Attribute Reference In addition to the arguments, the following attributes are available:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of the ` + "`" + `appfwprofile_cookieconsistency_binding` + "`" + `. It is the concatenation of the ` + "`" + `name` + "`" + ` and ` + "`" + `cookieconsistency` + "`" + ` attributes separated by a comma.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id of the ` + "`" + `appfwprofile_cookieconsistency_binding` + "`" + `. It is the concatenation of the ` + "`" + `name` + "`" + ` and ` + "`" + `cookieconsistency` + "`" + ` attributes separated by a comma.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "citrixadc_appfwprofile_crosssitescripting_binding",
			Category:         "Application Firewall",
			ShortDescription: ``,
			Description: `

The ` + "`" + `appfwprofile_crosssitescripting_binding` + "`" + ` resource is used to add binding between Applicatin Firewall Profile and CrossSiteScripting relaxation rule.

`,
			Keywords: []string{
				"application",
				"firewall",
				"appfwprofile",
				"crosssitescripting",
				"binding",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `Name of the profile to which to bind an exemption or rule.`,
				},
				resource.Attribute{
					Name:        "crosssitescripting",
					Description: `The web form field name.`,
				},
				resource.Attribute{
					Name:        "isregex_xss",
					Description: `Is the web form field name a regular expression?. Possible values: [ REGEX, NOTREGEX ]`,
				},
				resource.Attribute{
					Name:        "formactionurl_xss",
					Description: `The web form action URL.`,
				},
				resource.Attribute{
					Name:        "as_scan_location_xss",
					Description: `(Optional) Location of cross-site scripting exception - form field, header, cookie or URL. Possible values: [ FORMFIELD, HEADER, COOKIE, URL ]`,
				},
				resource.Attribute{
					Name:        "as_value_type_xss",
					Description: `(Optional) The web form value type. Possible values: [ Tag, Attribute, Pattern ]`,
				},
				resource.Attribute{
					Name:        "as_value_expr_xss",
					Description: `(Optional) The web form value expression.`,
				},
				resource.Attribute{
					Name:        "isvalueregex_xss",
					Description: `(Optional) Is the web form field value a regular expression?. Possible values: [ REGEX, NOTREGEX ]`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `(Optional) Enabled. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `(Optional) Any comments about the purpose of profile, or other useful information about the profile.`,
				},
				resource.Attribute{
					Name:        "isautodeployed",
					Description: `(Optional) Is the rule auto deployed by dynamic profile ?. Possible values: [ AUTODEPLOYED, NOTAUTODEPLOYED ]`,
				},
				resource.Attribute{
					Name:        "alertonly",
					Description: `(Optional) Send SNMP alert?. Possible values: [ on, off ] ## Attribute Reference In addition to the arguments, the following attributes are available:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of the ` + "`" + `appfwprofile_crosssitescripting_binding` + "`" + `.It is the concatenation of the ` + "`" + `name` + "`" + ` and ` + "`" + `crosssitescripting` + "`" + ` attributes separated by a comma.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id of the ` + "`" + `appfwprofile_crosssitescripting_binding` + "`" + `.It is the concatenation of the ` + "`" + `name` + "`" + ` and ` + "`" + `crosssitescripting` + "`" + ` attributes separated by a comma.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "citrixadc_appfwprofile_denyurl_binding",
			Category:         "Application Firewall",
			ShortDescription: ``,
			Description: `

The ` + "`" + `appfwprofile_denyurl_binding` + "`" + ` resource is used to add bindings between Application Firewall Profile and DenyURL relaxation rule.

`,
			Keywords: []string{
				"application",
				"firewall",
				"appfwprofile",
				"denyurl",
				"binding",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `Name of the profile to which to bind an exemption or rule.`,
				},
				resource.Attribute{
					Name:        "denyurl",
					Description: `A regular expression that designates a URL on the Deny URL list.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `(Optional) Enabled. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `(Optional) Any comments about the purpose of profile, or other useful information about the profile.`,
				},
				resource.Attribute{
					Name:        "isautodeployed",
					Description: `(Optional) Is the rule auto deployed by dynamic profile ?. Possible values: [ AUTODEPLOYED, NOTAUTODEPLOYED ]`,
				},
				resource.Attribute{
					Name:        "alertonly",
					Description: `(Optional) Send SNMP alert?. Possible values: [ on, off ] ## Attribute Reference In addition to the arguments, the following attributes are available:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of the ` + "`" + `appfwprofile_denyurl_binding` + "`" + `. It has the value of the string ` + "`" + `name,denyurl` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id of the ` + "`" + `appfwprofile_denyurl_binding` + "`" + `. It has the value of the string ` + "`" + `name,denyurl` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "citrixadc_appfwprofile_sqlinjection_binding",
			Category:         "Application Firewall",
			ShortDescription: ``,
			Description: `

The ` + "`" + `appfwprofile_sqlinjection_binding` + "`" + ` resource is used to add bindings between Applicatin Firewall Profile and HTML SQLInjection relaxation rule.

`,
			Keywords: []string{
				"application",
				"firewall",
				"appfwprofile",
				"sqlinjection",
				"binding",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `Name of the profile to which to bind an exemption or rule.`,
				},
				resource.Attribute{
					Name:        "sqlinjection",
					Description: `The web form field name.`,
				},
				resource.Attribute{
					Name:        "formactionurl_sql",
					Description: `The web form action URL.`,
				},
				resource.Attribute{
					Name:        "as_scan_location_sql",
					Description: `Location of SQL injection exception - form field, header or cookie. Possible values: [ FORMFIELD, HEADER, COOKIE ]`,
				},
				resource.Attribute{
					Name:        "isregex_sql",
					Description: `(Optional) Is the web form field name a regular expression?. Possible values: [ REGEX, NOTREGEX ]`,
				},
				resource.Attribute{
					Name:        "as_value_type_sql",
					Description: `(Optional) The web form value type. Possible values: [ Keyword, SpecialString, Wildchar ]`,
				},
				resource.Attribute{
					Name:        "as_value_expr_sql",
					Description: `(Optional) The web form value expression.`,
				},
				resource.Attribute{
					Name:        "isvalueregex_sql",
					Description: `(Optional) Is the web form field value a regular expression?. Possible values: [ REGEX, NOTREGEX ]`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `(Optional) Enabled. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `(Optional) Any comments about the purpose of profile, or other useful information about the profile.`,
				},
				resource.Attribute{
					Name:        "isautodeployed",
					Description: `(Optional) Is the rule auto deployed by dynamic profile ?. Possible values: [ AUTODEPLOYED, NOTAUTODEPLOYED ]`,
				},
				resource.Attribute{
					Name:        "alertonly",
					Description: `(Optional) Send SNMP alert?. Possible values: [ on, off ] ## Attribute Reference In addition to the arguments, the following attributes are available:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of the ` + "`" + `appfwprofile_sqlinjection_binding` + "`" + `. It is the concatenation of the ` + "`" + `name` + "`" + ` and ` + "`" + `sqlinjection` + "`" + ` attributes separated by a comma.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id of the ` + "`" + `appfwprofile_sqlinjection_binding` + "`" + `. It is the concatenation of the ` + "`" + `name` + "`" + ` and ` + "`" + `sqlinjection` + "`" + ` attributes separated by a comma.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "citrixadc_appfwprofile_starturl_binding",
			Category:         "Application Firewall",
			ShortDescription: ``,
			Description: `

The ` + "`" + `appfwprofile_starturl_binding` + "`" + ` resource is used to add bindings between Application Firewall Profile and StartURL relaxation rule.


`,
			Keywords: []string{
				"application",
				"firewall",
				"appfwprofile",
				"starturl",
				"binding",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `Name of the profile to which to bind an exemption or rule.`,
				},
				resource.Attribute{
					Name:        "starturl",
					Description: `A regular expression that designates a URL on the Start URL list.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `(Optional) Enabled. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `(Optional) Any comments about the purpose of profile, or other useful information about the profile.`,
				},
				resource.Attribute{
					Name:        "isautodeployed",
					Description: `(Optional) Is the rule auto deployed by dynamic profile ?. Possible values: [ AUTODEPLOYED, NOTAUTODEPLOYED ]`,
				},
				resource.Attribute{
					Name:        "alertonly",
					Description: `(Optional) Send SNMP alert?. Possible values: [ on, off ] ## Attribute Reference In addition to the arguments, the following attributes are available:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of the ` + "`" + `appfwprofile_starturl_binding` + "`" + `. It has the value of the string ` + "`" + `name,starturl` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id of the ` + "`" + `appfwprofile_starturl_binding` + "`" + `. It has the value of the string ` + "`" + `name,starturl` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "citrixadc_appfwxmlcontenttype",
			Category:         "Application Firewall",
			ShortDescription: ``,
			Description: `

The ` + "`" + `appfwxmlcontenttype` + "`" + ` resource is used to crate Application Firewall XML ContentType.

`,
			Keywords: []string{
				"application",
				"firewall",
				"appfwxmlcontenttype",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "xmlcontenttypevalue",
					Description: `Content type to be classified as XML.`,
				},
				resource.Attribute{
					Name:        "isregex",
					Description: `(Optional) Is field name a regular expression?. Possible values: [ REGEX, NOTREGEX ] ## Attribute Reference In addition to the arguments, the following attributes are available:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of the ` + "`" + `appfwxmlcontenttype` + "`" + `. It has the same value as the ` + "`" + `xmlcontenttypevalue` + "`" + ` attribute.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id of the ` + "`" + `appfwxmlcontenttype` + "`" + `. It has the same value as the ` + "`" + `xmlcontenttypevalue` + "`" + ` attribute.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "citrixadc_auditmessageaction",
			Category:         "Audit",
			ShortDescription: ``,
			Description: `

This resource is used to create audit message actions.


`,
			Keywords: []string{
				"audit",
				"auditmessageaction",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the audit message action.`,
				},
				resource.Attribute{
					Name:        "loglevel",
					Description: `(Optional) Audit log level, which specifies the severity level of the log message being generated.. The following loglevels are valid:`,
				},
				resource.Attribute{
					Name:        "stringbuilderexpr",
					Description: `(Optional) Default-syntax expression that defines the format and content of the log message.`,
				},
				resource.Attribute{
					Name:        "logtonewnslog",
					Description: `(Optional) Send the message to the new nslog. Possible values: [YES, NO].`,
				},
				resource.Attribute{
					Name:        "bypasssafetycheck",
					Description: `(Optional) Bypass the safety check and allow unsafe expressions. Possible values: [YES, NO]. ## Attribute Reference In addition to the arguments, the following attributes are available:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of the resource. It has the same value as the ` + "`" + `name` + "`" + ` attribute. ## Import A auditmessageaction can be imported using its name, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import citrixadc_auditmessageaction.tf_msgaction tf_msgaction ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id of the resource. It has the same value as the ` + "`" + `name` + "`" + ` attribute. ## Import A auditmessageaction can be imported using its name, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import citrixadc_auditmessageaction.tf_msgaction tf_msgaction ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "citrixadc_auditsyslogaction",
			Category:         "Audit",
			ShortDescription: ``,
			Description: `

This resource is used to create audit syslog actions.


`,
			Keywords: []string{
				"audit",
				"auditsyslogaction",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the audit syslog action.`,
				},
				resource.Attribute{
					Name:        "serverip",
					Description: `(Optional) IP address of the syslog server.`,
				},
				resource.Attribute{
					Name:        "serverdomainname",
					Description: `(Optional) SYSLOG server name as a FQDN. Mutually exclusive with serverIP/lbVserverName.`,
				},
				resource.Attribute{
					Name:        "domainresolveretry",
					Description: `(Optional) Time, in seconds, for which the Citrix ADC waits before sending another DNS query to resolve the host name of the syslog server if the last query failed.`,
				},
				resource.Attribute{
					Name:        "lbvservername",
					Description: `(Optional) Name of the LB vserver. Mutually exclusive with syslog serverIP/serverName.`,
				},
				resource.Attribute{
					Name:        "serverport",
					Description: `(Optional) Port on which the syslog server accepts connections.`,
				},
				resource.Attribute{
					Name:        "loglevel",
					Description: `(Optional) Audit log level, which specifies the types of events to log. Possible values: [ ALL, EMERGENCY, ALERT, CRITICAL, ERROR, WARNING, NOTICE, INFORMATIONAL, DEBUG, NONE ]`,
				},
				resource.Attribute{
					Name:        "dateformat",
					Description: `(Optional) Format of dates in the logs. Possible values: [ MMDDYYYY, DDMMYYYY, DDMMYYYY ]`,
				},
				resource.Attribute{
					Name:        "logfacility",
					Description: `(Optional) Facility value, as defined in RFC 3164, assigned to the log message. Possible values: [ LOCAL0, LOCAL1, LOCAL2, LOCAL3, LOCAL4, LOCAL5, LOCAL6, LOCAL7 ]`,
				},
				resource.Attribute{
					Name:        "tcp",
					Description: `(Optional) Log TCP messages. Possible values: [ NONE, ALL ]`,
				},
				resource.Attribute{
					Name:        "acl",
					Description: `(Optional) Log access control list (ACL) messages. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "timezone",
					Description: `(Optional) Time zone used for date and timestamps in the logs. Possible values: [ GMT\_TIME, LOCAL\_TIME ]`,
				},
				resource.Attribute{
					Name:        "userdefinedauditlog",
					Description: `(Optional) Log user-configurable log messages to syslog. Setting this parameter to NO causes auditing to ignore all user-configured message actions. Setting this parameter to YES causes auditing to log user-configured message actions that meet the other logging criteria. Possible values: [ YES, NO ]`,
				},
				resource.Attribute{
					Name:        "appflowexport",
					Description: `(Optional) Export log messages to AppFlow collectors. Appflow collectors are entities to which log messages can be sent so that some action can be performed on them. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "lsn",
					Description: `(Optional) Log lsn info. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "alg",
					Description: `(Optional) Log alg info. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "subscriberlog",
					Description: `(Optional) Log subscriber session event information. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "transport",
					Description: `(Optional) Transport type used to send auditlogs to syslog server. Possible values: [ TCP, UDP ]`,
				},
				resource.Attribute{
					Name:        "tcpprofilename",
					Description: `(Optional) Name of the TCP profile whose settings are to be applied to the audit server info to tune the TCP connection parameters.`,
				},
				resource.Attribute{
					Name:        "maxlogdatasizetohold",
					Description: `(Optional) Max size of log data that can be held in NSB chain of server info.`,
				},
				resource.Attribute{
					Name:        "dns",
					Description: `(Optional) Log DNS related syslog messages. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "contentinspectionlog",
					Description: `(Optional) Log Content Inspection event information. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "netprofile",
					Description: `(Optional) Name of the network profile. The SNIP configured in the network profile will be used as source IP while sending log messages.`,
				},
				resource.Attribute{
					Name:        "sslinterception",
					Description: `(Optional) Log SSL Interception event information. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "urlfiltering",
					Description: `(Optional) Log URL filtering event information. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "domainresolvenow",
					Description: `(Optional) Immediately send a DNS query to resolve the server's domain name. ## Attributes In addition to the arguments, the following attributes are available:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of the audit syslog action. It has the same value as the ` + "`" + `name` + "`" + ` attribute. ## Import An instance of the resource can be imported using its name, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import citrixadc_auditsyslogaction.tf_syslogaction tf_syslogaction ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "citrixadc_auditsyslogpolicy",
			Category:         "Audit",
			ShortDescription: ``,
			Description: `

The resource is used to create audit syslog policies


`,
			Keywords: []string{
				"audit",
				"auditsyslogpolicy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name for the policy.`,
				},
				resource.Attribute{
					Name:        "rule",
					Description: `(Optional) Name of the Citrix ADC named rule, or an expression, that defines the messages to be logged to the syslog server.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Optional) Syslog server action to perform when this policy matches traffic.`,
				},
				resource.Attribute{
					Name:        "globalbinding",
					Description: `(Optional) A single global binding block. Documented below. Global binding supports the following:`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `(Optional) The priority of the command policy.`,
				},
				resource.Attribute{
					Name:        "globalbindtype",
					Description: `(Optional) Possible values: [ SYSTEM\_GLOBAL, VPN\_GLOBAL, RNAT\_GLOBAL ]`,
				},
				resource.Attribute{
					Name:        "nextfactor",
					Description: `(Optional) On success invoke label. Applicable for advanced authentication policy binding.`,
				},
				resource.Attribute{
					Name:        "gotopriorityexpression",
					Description: `(Optional) Applicable only to advance authentication policy. Expression or other value specifying the next policy to be evaluated if the current policy evaluates to TRUE. Specify one of the following values:`,
				},
				resource.Attribute{
					Name:        "feature",
					Description: `(Optional) The feature to be checked while applying this config. ## Attribute Reference In addition to the arguments, the following attributes are available:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of the policy. It has the same value as the ` + "`" + `name` + "`" + ` attribute. ## Import An instance of the resource can be imported using its name, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import citrixadc_auditsyslogpolicy.tf_auditsyslogpolicy tf_auditsyslogpolicy ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id of the policy. It has the same value as the ` + "`" + `name` + "`" + ` attribute. ## Import An instance of the resource can be imported using its name, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import citrixadc_auditsyslogpolicy.tf_auditsyslogpolicy tf_auditsyslogpolicy ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "citrixadc_cluster",
			Category:         "Cluster",
			ShortDescription: ``,
			Description: `

The resource is used to create and modify Citrix ADC in cluster deployment.

If there is no cluster instantiated the resource will bootstrap
the first node so that the Cluster IP address is reachable
and then add cluster nodes according to the terraform execution plan.

If there already exists a cluster deployment with the Cluster IP address
reachable the resource will add, remove or modify cluster nodes
according to the terraform execution plan.

In both cases the Citrix ADC provider configuration should point to the
Cluster IP address.


`,
			Keywords: []string{
				"cluster",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "clip",
					Description: `(Required) Cluster IP address. It will be added on cluster bootstrap on the first cluster node.`,
				},
				resource.Attribute{
					Name:        "clid",
					Description: `(Required) Unique number that identifies the cluster.`,
				},
				resource.Attribute{
					Name:        "deadinterval",
					Description: `(Optional) Amount of time, in seconds, after which nodes that do not respond to the heartbeats are assumed to be down.If the value is less than 3 sec, set the ` + "`" + `hellointerval` + "`" + ` parameter to 200 msec.`,
				},
				resource.Attribute{
					Name:        "hellointerval",
					Description: `(Optional) Interval, in milliseconds, at which heartbeats are sent to each cluster node to check the health status.Set the value to 200 msec, if the ` + "`" + `deadinterval` + "`" + ` parameter is less than 3 sec.`,
				},
				resource.Attribute{
					Name:        "preemption",
					Description: `(Optional) Preempt a cluster node that is configured as a SPARE if an ACTIVE node becomes available. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "quorumtype",
					Description: `(Optional) Quorum Configuration Choices - "MAJORITY" (recommended) requires majority of nodes to be online for the cluster to be UP. "NONE" relaxes this requirement. Possible values: [ MAJORITY, NONE ]`,
				},
				resource.Attribute{
					Name:        "inc",
					Description: `(Optional) This option is required if the cluster nodes reside on different networks. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "processlocal",
					Description: `(Optional) By turning on this option packets destined to a service in a cluster will not under go any steering. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "retainconnectionsoncluster",
					Description: `(Optional) This option enables you to retain existing connections on a node joining a Cluster system or when a node is being configured for passive timeout. By default, this option is disabled. Possible values: [ YES, NO ]`,
				},
				resource.Attribute{
					Name:        "backplanebasedview",
					Description: `(Optional) View based on heartbeat only on bkplane interface. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "nodegroup",
					Description: `(Optional) The node group in a Cluster system used for transition from L2 to L3.`,
				},
				resource.Attribute{
					Name:        "bootstrap_poll_delay",
					Description: `(Optional) Time duration to wait before the first poll for the Cluster IP during first node bootstrap. Defaults to ` + "`" + `60s` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "bootstrap_poll_interval",
					Description: `(Optional) Time duration between subsequent polls for the Cluster IP during first node bootstrap. Defaults to ` + "`" + `60s` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "bootstrap_poll_timeout",
					Description: `(Optional) Time duration that defines http request timeout for each Cluster IP poll during first node bootstrap. Defaults to ` + "`" + `10s` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "bootstrap_total_timeout",
					Description: `(Optional) Time duration that defines the timeout for the whole operation of the first node bootstrap. If the node has not been added when this timeout expires the operation will fail. Defaults to ` + "`" + `10m` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "clip_migration_poll_delay",
					Description: `(Optional) Time duration to wait before the first poll for the Cluster IP during a Cluster IP address migration. Defaults to ` + "`" + `10s` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "clip_migration_poll_interval",
					Description: `(Optional) Time duration between subsequent polls for the Cluster IP during a Cluster IP address migration. Defaults to ` + "`" + `30s` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "clip_migration_poll_timeout",
					Description: `(Optional) Time duration that defines http request timeout for each Cluster IP poll during a Cluster IP address migration. Defaults to ` + "`" + `10s` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "clip_migration_total_timeout",
					Description: `(Optional) Time duration that defines the timeout for the whole operation of Cluster IP address migration. If the node has not been added when this timeout expires the operation will fail. Defaults to ` + "`" + `10m` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "node_add_poll_delay",
					Description: `(Optional) Time duration to wait before the first poll for the Cluster IP during the addition of a node to the cluster. Defaults to ` + "`" + `10s` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "node_add_poll_interval",
					Description: `(Optional) Time duration between subsequent polls for the Cluster IP during the addition of a node to the cluster. Defaults to ` + "`" + `30s` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "node_add_total_timeout",
					Description: `(Optional) Time duration that defines the timeout for the whole operation of the addition of a node to the cluster. If the node has not been added when this timeout expires the operation will fail. Defaults to ` + "`" + `10m` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "clusternode",
					Description: `(Required) Cluster node configuration blocks. Documented below.`,
				},
				resource.Attribute{
					Name:        "clusternodegroup",
					Description: `(Optional) One cluster nodegroup configuration block. Documented below. Cluster node supports the following:`,
				},
				resource.Attribute{
					Name:        "nodeid",
					Description: `(Required) Unique number that identifies the cluster node.`,
				},
				resource.Attribute{
					Name:        "ipaddress",
					Description: `(Required) Citrix ADC IP (NSIP) address of the appliance to add to the cluster. Must be an IPv4 address.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `(Optional) Admin state of the cluster node. Possible values: [ ACTIVE, SPARE, PASSIVE ]`,
				},
				resource.Attribute{
					Name:        "backplane",
					Description: `(Optional) Interface through which the node communicates with the other nodes in the cluster. Must be specified in the three-tuple form n/c/u, where n represents the node ID and c/u refers to the interface on the appliance.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `(Optional) Preference for selecting a node as the configuration coordinator. The node with the lowest priority value is selected as the configuration coordinator. When the current configuration coordinator goes down, the node with the next lowest priority is made the new configuration coordinator. When the original node comes back up, it will preempt the new configuration coordinator and take over as the configuration coordinator. -> When priority is not configured for any of the nodes or if multiple nodes have the same priority, the cluster elects one of the nodes as the configuration coordinator.`,
				},
				resource.Attribute{
					Name:        "nodegroup",
					Description: `(Optional) The default node group in a Cluster system.`,
				},
				resource.Attribute{
					Name:        "delay",
					Description: `(Optional) Applicable for Passive node and node becomes passive after this timeout (in minutes).`,
				},
				resource.Attribute{
					Name:        "tunnelmode",
					Description: `(Optional) To set the tunnel mode. Possible values: [ NONE, GRE, UDP ]`,
				},
				resource.Attribute{
					Name:        "clearnodegroupconfig",
					Description: `(Optional) Option to remove nodegroup config. Possible values: [ YES, NO ] -> When a standalone ADC instance joins a cluster the provider will issue a series of NITRO API calls against this particular node. The following arguments apply to the NITRO API client that will be used for this one time operation.`,
				},
				resource.Attribute{
					Name:        "endpoint",
					Description: `(Required) Defines the NITRO API endpoint prefix. Can use either ` + "`" + `http` + "`" + ` or ` + "`" + `https` + "`" + ` protocol.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Optional) Defines the username that will be used by the NITRO API for authentication. Defaults to the value of the same argument of the provider currently in effect.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Required) Defines the password that will be used by the NITRO API for authentication. Defaults to the value of the same argument of the provider currently in effect.`,
				},
				resource.Attribute{
					Name:        "insecure_skip_verify",
					Description: `(Optional) Boolean variable that defines if an error should be thrown if the target ADC's TLS certificate is not trusted. When ` + "`" + `true` + "`" + ` the error will be ignored. When ` + "`" + `false` + "`" + ` such an error will cause the failure of any provider operation. Defaults to ` + "`" + `false` + "`" + `. Cluster nodegroup supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the nodegroup. The name uniquely identifies the nodegroup on the cluster.`,
				},
				resource.Attribute{
					Name:        "strict",
					Description: `(Optional) Specifies whether cluster nodes, that are not part of the nodegroup, will be used as backup for the nodegroup. Possible values: [ YES, NO ]`,
				},
				resource.Attribute{
					Name:        "sticky",
					Description: `(Optional) Only one node can be bound to nodegroup with this option enabled. It specifies whether to prempt the traffic for the entities bound to nodegroup when owner node goes down and rejoins the cluster. Possible values: [ YES, NO ]`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `(Optional) State of the nodegroup. All the nodes binding to this nodegroup must have the same state. Possible values: [ ACTIVE, SPARE, PASSIVE ]`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of the cluster. It has the same value as the ` + "`" + `clid` + "`" + ` attribute.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id of the cluster. It has the same value as the ` + "`" + `clid` + "`" + ` attribute.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "citrixadc_clusterfiles_syncer",
			Category:         "Cluster",
			ShortDescription: ``,
			Description: `\_syncer

This resource is used to manually trigger the cluster files synchronization
operation.

It is the equivalent of running
` + "`" + `` + "`" + `` + "`" + `shell
sync cluster files
` + "`" + `` + "`" + `` + "`" + `
on the command prompt of the Cluster IP address.

By its nature this resource does not have a remote state to read or modify.

Any change in the local state will trigger the operation once more.



`,
			Keywords: []string{
				"cluster",
				"clusterfiles",
				"syncer",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "timestamp",
					Description: `(Required) A string representing the current time by convention. Its main purpose is to enable the user to trigger the operation again without having to manually taint the resource.`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `(Required) A list of values defining which directories and files will be synchronized. Possible values: [ all, bookmarks, ssl, htmlinjection, imports, misc, dns, krb, AAA, app\_catalog, all\_plus\_misc, all\_minus\_misc ] ## Attribute Reference In addition to the arguments, the following attributes are available:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of the clusterfiles\_syncer. It has the same value as the ` + "`" + `timestamp` + "`" + ` attribute.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id of the clusterfiles\_syncer. It has the same value as the ` + "`" + `timestamp` + "`" + ` attribute.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "citrixadc_cmppolicy",
			Category:         "CMP",
			ShortDescription: ``,
			Description: `

The cmppolicy resource is used to create compression policies.


`,
			Keywords: []string{
				"cmp",
				"cmppolicy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the HTTP compression policy. Must begin with an ASCII alphabetic or underscore (\_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Can be changed after the policy is created. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my cmp policy" or 'my cmp policy').`,
				},
				resource.Attribute{
					Name:        "rule",
					Description: `(Optional) Expression that determines which HTTP requests or responses match the compression policy. The following requirements apply only to the Citrix ADC CLI:`,
				},
				resource.Attribute{
					Name:        "resaction",
					Description: `(Optional) The built-in or user-defined compression action to apply to the response when the policy matches a request or response. ## Attribute Reference In addition to the arguments, the following attributes are available:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of the compression policy. It has the same value as the ` + "`" + `name` + "`" + ` attribute. ## Import A cmppolicy can be imported using its name, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import citrixadc_cmppolicy.tf_cmppolicy tf_cmppolicy ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id of the compression policy. It has the same value as the ` + "`" + `name` + "`" + ` attribute. ## Import A cmppolicy can be imported using its name, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import citrixadc_cmppolicy.tf_cmppolicy tf_cmppolicy ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "citrixadc_csaction",
			Category:         "Content Switching",
			ShortDescription: ``,
			Description: `

The csaction resource is used to create content switching actions.


`,
			Keywords: []string{
				"content",
				"switching",
				"csaction",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the content switching action.`,
				},
				resource.Attribute{
					Name:        "targetlbvserver",
					Description: `(Optional) Name of the load balancing virtual server to which the content is switched.`,
				},
				resource.Attribute{
					Name:        "targetvserver",
					Description: `(Optional) Name of the VPN, GSLB or Authentication virtual server to which the content is switched.`,
				},
				resource.Attribute{
					Name:        "targetvserverexpr",
					Description: `(Optional) Expression that evaluates to the target load balancing virtual server to which the content is switched.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `(Optional) Comment associated with this content switching action. ## Attribute Reference In addition to the arguments, the following attributes are available:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of the content switching action. It has the same value as the ` + "`" + `name` + "`" + ` attribute. ## Import A csaction can be imported using its name, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import citrixadc_csaction.tf_csaction tf_csaction ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id of the content switching action. It has the same value as the ` + "`" + `name` + "`" + ` attribute. ## Import A csaction can be imported using its name, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import citrixadc_csaction.tf_csaction tf_csaction ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "citrixadc_cspolicy",
			Category:         "Content Switching",
			ShortDescription: ``,
			Description: `

The resource is used to configure content switching policies.


`,
			Keywords: []string{
				"content",
				"switching",
				"cspolicy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "policyname",
					Description: `(Optional) Name of the content switching policy`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Optional) URL string that is matched with the URL of a request. Can contain a wildcard character. Specify the string value in the following format: [[prefix] [`,
				},
				resource.Attribute{
					Name:        "rule",
					Description: `(Optional) Expression, or name of a named expression, against which traffic is evaluated.`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `(Optional) The domain name. The string value can range to 63 characters.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Optional) Content switching action that names the target load balancing virtual server to which the traffic is switched.`,
				},
				resource.Attribute{
					Name:        "logaction",
					Description: `(Optional) The log action associated with the content switching policy.`,
				},
				resource.Attribute{
					Name:        "csvserver",
					Description: `(Required) Content switching vserver that this policy will bind to.`,
				},
				resource.Attribute{
					Name:        "targetlbvserver",
					Description: `(Optional) Targe load balancing vserver that will be used for the binding with the Content switching vserver.`,
				},
				resource.Attribute{
					Name:        "forcenew_id_set",
					Description: `(Optional) A list of terraform resource ids. Any change in the list of values will trigger the recreation of the cspolicy resource. Its main intent is to force the rebinding with the Content switching vserver defined in ` + "`" + `csvserver` + "`" + ` should it be deleted and recreated. The same applies for the LB vserver defined in ` + "`" + `targetlbvserver` + "`" + `. ## Attribute Reference In addition to the arguments, the following attributes are available:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of the cspolicy. It has the same value as the ` + "`" + `policyname` + "`" + ` attribute. ## Import An instance of the resource can be imported using its ` + "`" + `policyname` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import citrixadc_cspolicy.tf_cspolicy tf_cspolicy ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id of the cspolicy. It has the same value as the ` + "`" + `policyname` + "`" + ` attribute. ## Import An instance of the resource can be imported using its ` + "`" + `policyname` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import citrixadc_cspolicy.tf_cspolicy tf_cspolicy ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "citrixadc_csvserver",
			Category:         "Content Switching",
			ShortDescription: ``,
			Description: `

This resource is used to manage Content switching virtual servers.


`,
			Keywords: []string{
				"content",
				"switching",
				"csvserver",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name for the content switching virtual server.`,
				},
				resource.Attribute{
					Name:        "td",
					Description: `(Optional) Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.`,
				},
				resource.Attribute{
					Name:        "servicetype",
					Description: `(Optional) Protocol used by the virtual server. Possible values: [ HTTP, SSL, TCP, FTP, RTSP, SSL\_TCP, UDP, DNS, SIP\_UDP, SIP\_TCP, SIP\_SSL, ANY, RADIUS, RDP, MYSQL, MSSQL, DIAMETER, SSL\_DIAMETER, DNS\_TCP, ORACLE, SMPP, PROXY, MONGO, MONGO\_TLS ]`,
				},
				resource.Attribute{
					Name:        "ipv46",
					Description: `(Optional) IP address of the content switching virtual server.`,
				},
				resource.Attribute{
					Name:        "targettype",
					Description: `(Optional) Virtual server target type. Possible values: [ GSLB ]`,
				},
				resource.Attribute{
					Name:        "dnsrecordtype",
					Description: `(Optional) Possible values: [ A, AAAA, CNAME, NAPTR ]`,
				},
				resource.Attribute{
					Name:        "persistenceid",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "ippattern",
					Description: `(Optional) IP address pattern, in dotted decimal notation, for identifying packets to be accepted by the virtual server. The IP Mask parameter specifies which part of the destination IP address is matched against the pattern. Mutually exclusive with the IP Address parameter. For example, if the IP pattern assigned to the virtual server is 198.51.100.0 and the IP mask is 255.255.240.0 (a forward mask), the first 20 bits in the destination IP addresses are matched with the first 20 bits in the pattern. The virtual server accepts requests with IP addresses that range from 198.51.96.1 to 198.51.111.254. You can also use a pattern such as 0.0.2.2 and a mask such as 0.0.255.255 (a reverse mask). If a destination IP address matches more than one IP pattern, the pattern with the longest match is selected, and the associated virtual server processes the request. For example, if the virtual servers, vs1 and vs2, have the same IP pattern, 0.0.100.128, but different IP masks of 0.0.255.255 and 0.0.224.255, a destination IP address of 198.51.100.128 has the longest match with the IP pattern of vs1. If a destination IP address matches two or more virtual servers to the same extent, the request is processed by the virtual server whose port number matches the port number in the request.`,
				},
				resource.Attribute{
					Name:        "ipmask",
					Description: `(Optional) IP mask, in dotted decimal notation, for the IP Pattern parameter. Can have leading or trailing non-zero octets (for example, 255.255.240.0 or 0.0.255.255). Accordingly, the mask specifies whether the first n bits or the last n bits of the destination IP address in a client request are to be matched with the corresponding bits in the IP pattern. The former is called a forward mask. The latter is called a reverse mask.`,
				},
				resource.Attribute{
					Name:        "range",
					Description: `(Optional) Number of consecutive IP addresses, starting with the address specified by the IP Address parameter, to include in a range of addresses assigned to this virtual server.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional) Port number for content switching virtual server.`,
				},
				resource.Attribute{
					Name:        "ipset",
					Description: `(Optional) The list of IPv4/IPv6 addresses bound to ipset would form a part of listening service on the current cs vserver.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `(Optional) Initial state of the load balancing virtual server. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "stateupdate",
					Description: `(Optional) Enable state updates for a specific content switching virtual server. By default, the Content Switching virtual server is always UP, regardless of the state of the Load Balancing virtual servers bound to it. This parameter interacts with the global setting as follows: Global Level + Vserver Level = Result ENABLED + ENABLED = ENABLED ENABLED + DISABLED = ENABLED DISABLED + ENABLED = ENABLED DISABLED + DISABLED = DISABLED If you want to enable state updates for only some content switching virtual servers, be sure to disable the state update parameter. Possible values: [ ENABLED, DISABLED, UPDATEONBACKENDUPDATE ]`,
				},
				resource.Attribute{
					Name:        "cacheable",
					Description: `(Optional) Use this option to specify whether a virtual server, used for load balancing or content switching, routes requests to the cache redirection virtual server before sending it to the configured servers. Possible values: [ YES, NO ]`,
				},
				resource.Attribute{
					Name:        "redirecturl",
					Description: `(Optional) URL to which traffic is redirected if the virtual server becomes unavailable. The service type of the virtual server should be either HTTP or SSL. ~> Make sure that the domain in the URL does not match the domain specified for a content switching policy. If it does, requests are continuously redirected to the unavailable virtual server.`,
				},
				resource.Attribute{
					Name:        "clttimeout",
					Description: `(Optional) Idle time, in seconds, after which the client connection is terminated.`,
				},
				resource.Attribute{
					Name:        "precedence",
					Description: `(Optional) Type of precedence to use for both RULE-based and URL-based policies on the content switching virtual server. With the default (RULE) setting, incoming requests are evaluated against the rule-based content switching policies. If none of the rules match, the URL in the request is evaluated against the URL-based content switching policies. Possible values: [ RULE, URL ]`,
				},
				resource.Attribute{
					Name:        "casesensitive",
					Description: `(Optional) Consider case in URLs (for policies that use URLs instead of RULES). For example, with the ON setting, the URLs /a/1.html and /A/1.HTML are treated differently and can have different targets (set by content switching policies). With the OFF setting, /a/1.html and /A/1.HTML are switched to the same target. Possible values: [ ON, OFF ]`,
				},
				resource.Attribute{
					Name:        "somethod",
					Description: `(Optional) Type of spillover used to divert traffic to the backup virtual server when the primary virtual server reaches the spillover threshold. Connection spillover is based on the number of connections. Bandwidth spillover is based on the total Kbps of incoming and outgoing traffic. Possible values: [ CONNECTION, DYNAMICCONNECTION, BANDWIDTH, HEALTH, NONE ]`,
				},
				resource.Attribute{
					Name:        "sopersistence",
					Description: `(Optional) Maintain source-IP based persistence on primary and backup virtual servers. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "sopersistencetimeout",
					Description: `(Optional) Time-out value, in minutes, for spillover persistence.`,
				},
				resource.Attribute{
					Name:        "sothreshold",
					Description: `(Optional) Depending on the spillover method, the maximum number of connections or the maximum total bandwidth (Kbps) that a virtual server can handle before spillover occurs.`,
				},
				resource.Attribute{
					Name:        "sobackupaction",
					Description: `(Optional) Action to be performed if spillover is to take effect, but no backup chain to spillover is usable or exists. Possible values: [ DROP, ACCEPT, REDIRECT ]`,
				},
				resource.Attribute{
					Name:        "redirectportrewrite",
					Description: `(Optional) State of port rewrite while performing HTTP redirect.Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "downstateflush",
					Description: `(Optional) Flush all active transactions associated with a virtual server whose state transitions from UP to DOWN. Do not enable this option for applications that must complete their transactions. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "backupvserver",
					Description: `(Optional) Name of the backup virtual server that you are configuring.`,
				},
				resource.Attribute{
					Name:        "disableprimaryondown",
					Description: `(Optional) Continue forwarding the traffic to backup virtual server even after the primary server comes UP from the DOWN state. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "insertvserveripport",
					Description: `(Optional) Insert the virtual server's VIP address and port number in the request header. Available values function as follows: VIPADDR - Header contains the vserver's IP address and port number without any translation. OFF - The virtual IP and port header insertion option is disabled. V6TOV4MAPPING - Header contains the mapped IPv4 address corresponding to the IPv6 address of the vserver and the port number. An IPv6 address can be mapped to a user-specified IPv4 address using the set ns ip6 command. Possible values: [ OFF, VIPADDR, V6TOV4MAPPING ]`,
				},
				resource.Attribute{
					Name:        "vipheader",
					Description: `(Optional) Name of virtual server IP and port header, for use with the VServer IP Port Insertion parameter.`,
				},
				resource.Attribute{
					Name:        "rtspnat",
					Description: `(Optional) Enable network address translation (NAT) for real-time streaming protocol (RTSP) connections. Possible values: [ ON, OFF ]`,
				},
				resource.Attribute{
					Name:        "authenticationhost",
					Description: `(Optional) FQDN of the authentication virtual server. The service type of the virtual server should be either HTTP or SSL.`,
				},
				resource.Attribute{
					Name:        "listenpolicy",
					Description: `(Optional) String specifying the listen policy for the content switching virtual server. Can be either the name of an existing expression or an in-line expression.`,
				},
				resource.Attribute{
					Name:        "listenpriority",
					Description: `(Optional) Integer specifying the priority of the listen policy. A higher number specifies a lower priority. If a request matches the listen policies of more than one virtual server the virtual server whose listen policy has the highest priority (the lowest priority number) accepts the request.`,
				},
				resource.Attribute{
					Name:        "authn401",
					Description: `(Optional) Enable HTTP 401-response based authentication. Possible values: [ ON, OFF ]`,
				},
				resource.Attribute{
					Name:        "authnvsname",
					Description: `(Optional) Name of authentication virtual server that authenticates the incoming user requests to this content switching virtual server.`,
				},
				resource.Attribute{
					Name:        "push",
					Description: `(Optional) Process traffic with the push virtual server that is bound to this content switching virtual server (specified by the Push VServer parameter). The service type of the push virtual server should be either HTTP or SSL. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "pushvserver",
					Description: `(Optional) Name of the load balancing virtual server, of type PUSH or SSL_PUSH, to which the server pushes updates received on the client-facing load balancing virtual server.`,
				},
				resource.Attribute{
					Name:        "pushlabel",
					Description: `(Optional) Expression for extracting the label from the response received from server. This string can be either an existing rule name or an inline expression. The service type of the virtual server should be either HTTP or SSL.`,
				},
				resource.Attribute{
					Name:        "pushmulticlients",
					Description: `(Optional) Allow multiple Web 2.0 connections from the same client to connect to the virtual server and expect updates. Possible values: [ YES, NO ]`,
				},
				resource.Attribute{
					Name:        "httpprofilename",
					Description: `(Optional) Name of the HTTP profile containing HTTP configuration settings for the virtual server. The service type of the virtual server should be either HTTP or SSL.`,
				},
				resource.Attribute{
					Name:        "dbprofilename",
					Description: `(Optional) Name of the DB profile.`,
				},
				resource.Attribute{
					Name:        "oracleserverversion",
					Description: `(Optional) Oracle server version. Possible values: [ 10G, 11G ]`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `(Optional) Information about this virtual server.`,
				},
				resource.Attribute{
					Name:        "mssqlserverversion",
					Description: `(Optional) The version of the MSSQL server. Possible values: [ 70, 2000, 2000SP1, 2005, 2008, 2008R2, 2012, 2014 ]`,
				},
				resource.Attribute{
					Name:        "l2conn",
					Description: `(Optional) Use L2 Parameters to identify a connection. Possible values: [ ON, OFF ]`,
				},
				resource.Attribute{
					Name:        "mysqlprotocolversion",
					Description: `(Optional) The protocol version returned by the mysql vserver.`,
				},
				resource.Attribute{
					Name:        "mysqlserverversion",
					Description: `(Optional) The server version string returned by the mysql vserver.`,
				},
				resource.Attribute{
					Name:        "mysqlcharacterset",
					Description: `(Optional) The character set returned by the mysql vserver.`,
				},
				resource.Attribute{
					Name:        "mysqlservercapabilities",
					Description: `(Optional) The server capabilities returned by the mysql vserver.`,
				},
				resource.Attribute{
					Name:        "appflowlog",
					Description: `(Optional) Enable logging appflow flow information. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "netprofile",
					Description: `(Optional) The name of the network profile.`,
				},
				resource.Attribute{
					Name:        "icmpvsrresponse",
					Description: `(Optional) Can be active or passive. Possible values: [ PASSIVE, ACTIVE ]`,
				},
				resource.Attribute{
					Name:        "rhistate",
					Description: `(Optional) A host route is injected according to the setting on the virtual servers If set to PASSIVE on all the virtual servers that share the IP address, the appliance always injects the hostroute. If set to ACTIVE on all the virtual servers that share the IP address, the appliance injects even if one virtual server is UP. If set to ACTIVE on some virtual servers and PASSIVE on the others, the appliance, injects even if one virtual server set to ACTIVE is UP. Default value: PASSIVE Possible values: [ PASSIVE, ACTIVE ]`,
				},
				resource.Attribute{
					Name:        "authnprofile",
					Description: `(Optional) Name of the authentication profile to be used when authentication is turned on.`,
				},
				resource.Attribute{
					Name:        "dnsprofilename",
					Description: `(Optional) Name of the DNS profile to be associated with the VServer. DNS profile properties will applied to the transactions processed by a VServer. This parameter is valid only for DNS and DNS-TCP VServers.`,
				},
				resource.Attribute{
					Name:        "domainname",
					Description: `(Optional) Domain name for which to change the time to live (TTL) and/or backup service IP address.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "backupip",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "cookiedomain",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "cookietimeout",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "sitedomainttl",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "tcpprofilename",
					Description: `(Optional) Name of the TCP profile containing TCP configuration settings for the virtual server.`,
				},
				resource.Attribute{
					Name:        "ciphersuites",
					Description: `(Optional) List of ciphersuites to bind to the cs virtual server.`,
				},
				resource.Attribute{
					Name:        "ciphers",
					Description: `(Optional) List of ciphers to bind to the cs virtual server. Applicable only for cluster deployments of ADC.`,
				},
				resource.Attribute{
					Name:        "lbvserverbinding",
					Description: `(Optional) Name of lb vserver to bind to the cs vserver.`,
				},
				resource.Attribute{
					Name:        "sslprofile",
					Description: `(Optional) SSL profile to bind to cs vserver.`,
				},
				resource.Attribute{
					Name:        "snisslcertkeys",
					Description: `(Optional) List of SNI SSL certificates to bind to the cs vserver.`,
				},
				resource.Attribute{
					Name:        "sslcertkey",
					Description: `(Optional) SSL certificate to bind to this cs vserver.`,
				},
				resource.Attribute{
					Name:        "sslpolicybinding",
					Description: `(Optional) A block defining an ssl policy binding. See below for details. An ` + "`" + `sslpolicybinding` + "`" + ` block may contain the following attributes:`,
				},
				resource.Attribute{
					Name:        "policyname",
					Description: `(Optional) The name of the SSL policy binding.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `(Optional) The priority of the policies bound to this SSL service.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) Bind point to which to bind the policy. Possible Values: REQUEST, INTERCEPT_REQ and CLIENTHELLO_REQ. These bindpoints mean: 1. REQUEST: Policy evaluation will be done at appplication above SSL. This bindpoint is default and is used for actions based on clientauth and client cert. 2. INTERCEPT_REQ: Policy evaluation will be done during SSL handshake to decide whether to intercept or not. Actions allowed with this type are: INTERCEPT, BYPASS and RESET. 3. CLIENTHELLO_REQ: Policy evaluation will be done during handling of Client Hello Request. Action allowed with this type is: RESET, FORWARD and PICKCACERTGRP. Possible values: [ INTERCEPT_REQ, REQUEST, CLIENTHELLO_REQ ]`,
				},
				resource.Attribute{
					Name:        "gotopriorityexpression",
					Description: `(Optional) Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.`,
				},
				resource.Attribute{
					Name:        "invoke",
					Description: `(Optional) Invoke flag. This attribute is relevant only for ADVANCED policies.`,
				},
				resource.Attribute{
					Name:        "labeltype",
					Description: `(Optional) Type of policy label invocation. Possible values: [ vserver, service, policylabel ]`,
				},
				resource.Attribute{
					Name:        "labelname",
					Description: `(Optional) Name of the label to invoke if the current policy rule evaluates to TRUE. ## Attribute Reference In addition to the arguments, the following attributes are available:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of the csvserver. It has the same value as the ` + "`" + `name` + "`" + ` attribute. ## Import An instance of the resource can be imported using its name, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import citrixadc_csvserver.tf_csvserver tf_csvserver ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id of the csvserver. It has the same value as the ` + "`" + `name` + "`" + ` attribute. ## Import An instance of the resource can be imported using its name, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import citrixadc_csvserver.tf_csvserver tf_csvserver ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "citrixadc_csvserver_appfwpolicy_binding",
			Category:         "Content Switching",
			ShortDescription: ``,
			Description: `

The csvserver_appfwpolicy_binding resource is used to add csvserver_appfwpolicy_binding.

`,
			Keywords: []string{
				"content",
				"switching",
				"csvserver",
				"appfwpolicy",
				"binding",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `Name of the content switching virtual server to which the content switching policy applies.`,
				},
				resource.Attribute{
					Name:        "policyname",
					Description: `Policies bound to this vserver.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `(Optional) Priority for the policy.`,
				},
				resource.Attribute{
					Name:        "gotopriorityexpression",
					Description: `(Optional) Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.`,
				},
				resource.Attribute{
					Name:        "bindpoint",
					Description: `(Optional) The bindpoint to which the policy is bound. Possible values: [ REQUEST, RESPONSE, ICA_REQUEST, OTHERTCP_REQUEST ]`,
				},
				resource.Attribute{
					Name:        "invoke",
					Description: `(Optional) Invoke flag.`,
				},
				resource.Attribute{
					Name:        "labeltype",
					Description: `(Optional) The invocation type. Possible values: [ reqvserver, resvserver, policylabel ]`,
				},
				resource.Attribute{
					Name:        "labelname",
					Description: `(Optional) Name of the label invoked.`,
				},
				resource.Attribute{
					Name:        "targetlbvserver",
					Description: `(Optional) Name of the Load Balancing virtual server to which the content is switched, if policy rule is evaluated to be TRUE. Example: bind cs vs cs1 -policyname pol1 -priority 101 -targetLBVserver lb1 Note: Use this parameter only in case of Content Switching policy bind operations to a CS vserver. ## Attribute Reference In addition to the arguments, the following attributes are available:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of the csvserver_appfwpolicy_binding. It has the same value as the ` + "`" + `name` + "`" + ` attribute.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id of the csvserver_appfwpolicy_binding. It has the same value as the ` + "`" + `name` + "`" + ` attribute.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "citrixadc_csvserver_cmppolicy_binding",
			Category:         "Content Switching",
			ShortDescription: ``,
			Description: `\_cmppolicy\_binding

The csvserver\_cmppolicy\_binding resource is used to create bindings between a content switching vserver and compression policy.


`,
			Keywords: []string{
				"content",
				"switching",
				"csvserver",
				"cmppolicy",
				"binding",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "policyname",
					Description: `(Optional) Policies bound to this vserver.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `(Optional) Priority for the policy.`,
				},
				resource.Attribute{
					Name:        "gotopriorityexpression",
					Description: `(Optional) Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.`,
				},
				resource.Attribute{
					Name:        "bindpoint",
					Description: `(Optional) The bindpoint to which the policy is bound. Possible values: [ REQUEST, RESPONSE, ICA_REQUEST, OTHERTCP_REQUEST ]`,
				},
				resource.Attribute{
					Name:        "invoke",
					Description: `(Optional) Invoke flag.`,
				},
				resource.Attribute{
					Name:        "labeltype",
					Description: `(Optional) The invocation type. Possible values: [ reqvserver, resvserver, policylabel ]`,
				},
				resource.Attribute{
					Name:        "labelname",
					Description: `(Optional) Name of the label invoked.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the content switching virtual server to which the content switching policy applies.`,
				},
				resource.Attribute{
					Name:        "targetlbvserver",
					Description: `(Optional) Name of the Load Balancing virtual server to which the content is switched, if policy rule is evaluated to be TRUE. Example: bind cs vs cs1 -policyname pol1 -priority 101 -targetLBVserver lb1 Note: Use this parameter only in case of Content Switching policy bind operations to a CS vserver. ## Attribute Reference In addition to the arguments, the following attributes are available:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of the csvserver\_cmppolicy\_binding. It is the concatenation of the ` + "`" + `name` + "`" + ` and ` + "`" + `policyname` + "`" + ` attributes separated by a comma. ## Import A csvserver\_cmppolicy\_binding can be imported using its id e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import citrixadc_csvserver_cmppolicy_binding.tf_bind tf_csvserver,tf_cmppolicy ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id of the csvserver\_cmppolicy\_binding. It is the concatenation of the ` + "`" + `name` + "`" + ` and ` + "`" + `policyname` + "`" + ` attributes separated by a comma. ## Import A csvserver\_cmppolicy\_binding can be imported using its id e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import citrixadc_csvserver_cmppolicy_binding.tf_bind tf_csvserver,tf_cmppolicy ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "citrixadc_csvserver_cspolicy_binding",
			Category:         "Content Switching",
			ShortDescription: ``,
			Description: `\_cspolicy\_binding

The csvserver\_cspolicy\_binding resource is used to create bindings between a content switching virtual server and a content switcing policy.


`,
			Keywords: []string{
				"content",
				"switching",
				"csvserver",
				"cspolicy",
				"binding",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "policyname",
					Description: `(Optional) Policies bound to this vserver.`,
				},
				resource.Attribute{
					Name:        "targetlbvserver",
					Description: `(Optional) target vserver name.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `(Optional) Priority for the policy.`,
				},
				resource.Attribute{
					Name:        "gotopriorityexpression",
					Description: `(Optional) Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.`,
				},
				resource.Attribute{
					Name:        "bindpoint",
					Description: `(Optional) The bindpoint to which the policy is bound. Possible values: [ REQUEST, RESPONSE, ICA_REQUEST, OTHERTCP_REQUEST ]`,
				},
				resource.Attribute{
					Name:        "invoke",
					Description: `(Optional) Invoke flag.`,
				},
				resource.Attribute{
					Name:        "labeltype",
					Description: `(Optional) The invocation type. Possible values: [ reqvserver, resvserver, policylabel ]`,
				},
				resource.Attribute{
					Name:        "labelname",
					Description: `(Optional) Name of the label invoked.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the content switching virtual server to which the content switching policy applies. ## Attribute Reference In addition to the arguments, the following attributes are available:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of the csvserver\_cspolicy\_binding . It is the concatenation of the ` + "`" + `name` + "`" + ` and ` + "`" + `policyname` + "`" + ` attributes separated by a comma. ## Import A csvserver\_cspolicy\_binding can be imported using its id, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import citrixadc_csvserver_cspolicy_binding.tf_csvscspolbind tf_csvserver,tf_cspolicy ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id of the csvserver\_cspolicy\_binding . It is the concatenation of the ` + "`" + `name` + "`" + ` and ` + "`" + `policyname` + "`" + ` attributes separated by a comma. ## Import A csvserver\_cspolicy\_binding can be imported using its id, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import citrixadc_csvserver_cspolicy_binding.tf_csvscspolbind tf_csvserver,tf_cspolicy ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "citrixadc_csvserver_filterpolicy_binding",
			Category:         "Content Switching",
			ShortDescription: ``,
			Description: `\_filterpolicy\_binding

The csvserver\_filterpolicy\_binding resource is used to bind content switching virtual server with filter policies.


`,
			Keywords: []string{
				"content",
				"switching",
				"csvserver",
				"filterpolicy",
				"binding",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "policyname",
					Description: `(Required) Policies bound to this vserver.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `(Optional) Priority for the policy.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the content switching virtual server to which the content switching policy applies.`,
				},
				resource.Attribute{
					Name:        "targetlbvserver",
					Description: `(Optional) Name of the Load Balancing virtual server to which the content is switched, if policy rule is evaluated to be TRUE. Example: bind cs vs cs1 -policyname pol1 -priority 101 -targetLBVserver lb1 Note: Use this parameter only in case of Content Switching policy bind operations to a CS vserver.`,
				},
				resource.Attribute{
					Name:        "gotopriorityexpression",
					Description: `(Optional) Expression or other value specifying the next policy to be evaluated if the current policy evaluates to TRUE. Specify one of the following values:`,
				},
				resource.Attribute{
					Name:        "bindpoint",
					Description: `(Optional) For a rewrite policy, the bind point to which to bind the policy. Note: This parameter applies only to rewrite policies, because content switching policies are evaluated only at request time. Possible values: [ REQUEST, RESPONSE, ICA_REQUEST, OTHERTCP_REQUEST ]`,
				},
				resource.Attribute{
					Name:        "invoke",
					Description: `(Optional) Invoke a policy label if this policy's rule evaluates to TRUE (valid only for default-syntax policies such as application firewall, transform, integrated cache, rewrite, responder, and content switching).`,
				},
				resource.Attribute{
					Name:        "labeltype",
					Description: `(Optional) Type of label to be invoked. Possible values: [ reqvserver, resvserver, policylabel ]`,
				},
				resource.Attribute{
					Name:        "labelname",
					Description: `(Optional) Name of the label to be invoked. ## Attribute Reference In addition to the arguments, the following attributes are available:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of the csvserver\_filterpolicy\_binding. It is the concatenation of the ` + "`" + `name` + "`" + ` and ` + "`" + `policyname` + "`" + ` attributes separated by a comma. ## Import A csvserver\_filterpolicy\_binding can be imported using its id, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import citrixadc_csvserver_filterpolicy_binding.tf_bind tf_csvserver,tf_filterpolicy ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id of the csvserver\_filterpolicy\_binding. It is the concatenation of the ` + "`" + `name` + "`" + ` and ` + "`" + `policyname` + "`" + ` attributes separated by a comma. ## Import A csvserver\_filterpolicy\_binding can be imported using its id, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import citrixadc_csvserver_filterpolicy_binding.tf_bind tf_csvserver,tf_filterpolicy ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "citrixadc_csvserver_responderpolicy_binding",
			Category:         "Content Switching",
			ShortDescription: ``,
			Description: `\_responderpolicy\_binding

The csvserver\_responderpolicy\_binding resource is used to bind content switching virtual servers with responder policies


`,
			Keywords: []string{
				"content",
				"switching",
				"csvserver",
				"responderpolicy",
				"binding",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "policyname",
					Description: `(Required) Policies bound to this vserver.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `(Optional) Priority for the policy.`,
				},
				resource.Attribute{
					Name:        "gotopriorityexpression",
					Description: `(Optional) Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.`,
				},
				resource.Attribute{
					Name:        "invoke",
					Description: `(Optional) Invoke flag.`,
				},
				resource.Attribute{
					Name:        "labeltype",
					Description: `(Optional) The invocation type. Possible values: [ reqvserver, resvserver, policylabel ]`,
				},
				resource.Attribute{
					Name:        "labelname",
					Description: `(Optional) Name of the label invoked.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the content switching virtual server to which the content switching policy applies.`,
				},
				resource.Attribute{
					Name:        "targetlbvserver",
					Description: `(Optional) Name of the Load Balancing virtual server to which the content is switched, if policy rule is evaluated to be TRUE. Example: bind cs vs cs1 -policyname pol1 -priority 101 -targetLBVserver lb1 Note: Use this parameter only in case of Content Switching policy bind operations to a CS vserver.`,
				},
				resource.Attribute{
					Name:        "bindpoint",
					Description: `(Optional) For a rewrite policy, the bind point to which to bind the policy. Note: This parameter applies only to rewrite policies, because content switching policies are evaluated only at request time. Possible values: [ REQUEST, RESPONSE, ICA_REQUEST, OTHERTCP_REQUEST ] ## Attribute Reference In addition to the arguments, the following attributes are available:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of the csvserver\_responderpolicy\_binding. It is the concatenation of the ` + "`" + `name` + "`" + ` and ` + "`" + `policyname` + "`" + ` attributes separated by a comma. ## Import A csvserver\_responderpolicy\_binding can be imported using its id, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import citrixadc_csvserver_responderpolicy_binding.tf_bind tf_csvserver,tf_responder_policy ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id of the csvserver\_responderpolicy\_binding. It is the concatenation of the ` + "`" + `name` + "`" + ` and ` + "`" + `policyname` + "`" + ` attributes separated by a comma. ## Import A csvserver\_responderpolicy\_binding can be imported using its id, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import citrixadc_csvserver_responderpolicy_binding.tf_bind tf_csvserver,tf_responder_policy ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "citrixadc_csvserver_rewritepolicy_binding",
			Category:         "Content Switching",
			ShortDescription: ``,
			Description: `\_rewritepolicy\_binding

The csvserver\_rewritepolicy\_binding resource is used to bind content switching virtual servers to rewrite policies.


`,
			Keywords: []string{
				"content",
				"switching",
				"csvserver",
				"rewritepolicy",
				"binding",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "policyname",
					Description: `(Required) Policies bound to this vserver.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `(Optional) Priority for the policy.`,
				},
				resource.Attribute{
					Name:        "gotopriorityexpression",
					Description: `(Optional) Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.`,
				},
				resource.Attribute{
					Name:        "bindpoint",
					Description: `(Optional) The bindpoint to which the policy is bound. Possible values: [ REQUEST, RESPONSE, ICA_REQUEST, OTHERTCP_REQUEST ]`,
				},
				resource.Attribute{
					Name:        "invoke",
					Description: `(Optional) Invoke flag.`,
				},
				resource.Attribute{
					Name:        "labeltype",
					Description: `(Optional) The invocation type. Possible values: [ reqvserver, resvserver, policylabel ]`,
				},
				resource.Attribute{
					Name:        "labelname",
					Description: `(Optional) Name of the label invoked.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the content switching virtual server to which the content switching policy applies.`,
				},
				resource.Attribute{
					Name:        "targetlbvserver",
					Description: `(Optional) Name of the Load Balancing virtual server to which the content is switched, if policy rule is evaluated to be TRUE. Example: bind cs vs cs1 -policyname pol1 -priority 101 -targetLBVserver lb1 Note: Use this parameter only in case of Content Switching policy bind operations to a CS vserver. ## Attribute Reference In addition to the arguments, the following attributes are available:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of the csvserver\_rewritepolicy\_binding. It is the concatenation of ` + "`" + `name` + "`" + ` and ` + "`" + `policyname` + "`" + ` separated by a comma. ## Import A csvserver\_rewritepolicy\_binding can be imported using its id, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import citrixadc_csvserver_rewritepolicy_binding.tf_bind tf_csvserver,tf_rewrite_policy ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id of the csvserver\_rewritepolicy\_binding. It is the concatenation of ` + "`" + `name` + "`" + ` and ` + "`" + `policyname` + "`" + ` separated by a comma. ## Import A csvserver\_rewritepolicy\_binding can be imported using its id, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import citrixadc_csvserver_rewritepolicy_binding.tf_bind tf_csvserver,tf_rewrite_policy ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "citrixadc_csvserver_transformpolicy_binding",
			Category:         "Content Switching",
			ShortDescription: ``,
			Description: `\_transformpolicy\_binding

The csvserver\_transformpolicy\_binding resource is used to bind content swtching virtual servers with transform policies.


`,
			Keywords: []string{
				"content",
				"switching",
				"csvserver",
				"transformpolicy",
				"binding",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "policyname",
					Description: `(Required) Policies bound to this vserver.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `(Optional) Priority for the policy.`,
				},
				resource.Attribute{
					Name:        "gotopriorityexpression",
					Description: `(Optional) Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.`,
				},
				resource.Attribute{
					Name:        "bindpoint",
					Description: `(Optional) The bindpoint to which the policy is bound. Possible values: [ REQUEST, RESPONSE, ICA_REQUEST, OTHERTCP_REQUEST ]`,
				},
				resource.Attribute{
					Name:        "invoke",
					Description: `(Optional) Invoke flag.`,
				},
				resource.Attribute{
					Name:        "labeltype",
					Description: `(Optional) The invocation type. Possible values: [ reqvserver, resvserver, policylabel ]`,
				},
				resource.Attribute{
					Name:        "labelname",
					Description: `(Optional) Name of the label invoked.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the content switching virtual server to which the content switching policy applies.`,
				},
				resource.Attribute{
					Name:        "targetlbvserver",
					Description: `(Optional) Name of the Load Balancing virtual server to which the content is switched, if policy rule is evaluated to be TRUE. Example: bind cs vs cs1 -policyname pol1 -priority 101 -targetLBVserver lb1 Note: Use this parameter only in case of Content Switching policy bind operations to a CS vserver. ## Attribute Reference In addition to the arguments, the following attributes are available:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of the csvserver\_transformpolicy\_binding. It is the concatenation of the ` + "`" + `name` + "`" + ` and ` + "`" + `policyname` + "`" + ` attributes separated by a comma. ## Import A csvserver\_transformpolicy\_binding can be imported using its id, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import citrixadc_csvserver_transformpolicy_binding.tf_binding tf_csvserver,tf_trans_policy ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id of the csvserver\_transformpolicy\_binding. It is the concatenation of the ` + "`" + `name` + "`" + ` and ` + "`" + `policyname` + "`" + ` attributes separated by a comma. ## Import A csvserver\_transformpolicy\_binding can be imported using its id, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import citrixadc_csvserver_transformpolicy_binding.tf_binding tf_csvserver,tf_trans_policy ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "citrixadc_dnsnsrec",
			Category:         "DNS",
			ShortDescription: ``,
			Description: `

The dnsnsrec resource is used to create dns name server records.


`,
			Keywords: []string{
				"dns",
				"dnsnsrec",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "domain",
					Description: `(Optional) Domain name.`,
				},
				resource.Attribute{
					Name:        "nameserver",
					Description: `(Optional) Host name of the name server to add to the domain.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `(Optional) Time to Live (TTL), in seconds, for the record. TTL is the time for which the record must be cached by DNS proxies. The specified TTL is applied to all the resource records that are of the same record type and belong to the specified domain name. For example, if you add an address record, with a TTL of 36000, to the domain name example.com, the TTLs of all the address records of example.com are changed to 36000. If the TTL is not specified, the Citrix ADC uses either the DNS zone's minimum TTL or, if the SOA record is not available on the appliance, the default value of 3600. ## Attribute Reference In addition to the arguments, the following attributes are available:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of the dnsnsrec. It is the concatenation of the ` + "`" + `domain` + "`" + ` and ` + "`" + `nameserver` + "`" + ` attributes separated by a comma. ## Import A dnsnsrec can be imported using its id, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import citrixadc_dnsnsrec.tf_dnsnsrec www.test.com,192.168.1.100 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id of the dnsnsrec. It is the concatenation of the ` + "`" + `domain` + "`" + ` and ` + "`" + `nameserver` + "`" + ` attributes separated by a comma. ## Import A dnsnsrec can be imported using its id, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import citrixadc_dnsnsrec.tf_dnsnsrec www.test.com,192.168.1.100 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "citrixadc_dnssoarec",
			Category:         "DNS",
			ShortDescription: ``,
			Description: `

The dnssoarec resource is used to create dns SOA records.


`,
			Keywords: []string{
				"dns",
				"dnssoarec",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "domain",
					Description: `(Required) Domain name for which to add the SOA record.`,
				},
				resource.Attribute{
					Name:        "originserver",
					Description: `(Optional) Domain name of the name server that responds authoritatively for the domain.`,
				},
				resource.Attribute{
					Name:        "contact",
					Description: `(Optional) Email address of the contact to whom domain issues can be addressed. In the email address, replace the @ sign with a period (.). For example, enter domainadmin.example.com instead of domainadmin@example.com.`,
				},
				resource.Attribute{
					Name:        "serial",
					Description: `(Optional) The secondary server uses this parameter to determine whether it requires a zone transfer from the primary server.`,
				},
				resource.Attribute{
					Name:        "refresh",
					Description: `(Optional) Time, in seconds, for which a secondary server must wait between successive checks on the value of the serial number.`,
				},
				resource.Attribute{
					Name:        "retry",
					Description: `(Optional) Time, in seconds, between retries if a secondary server's attempt to contact the primary server for a zone refresh fails.`,
				},
				resource.Attribute{
					Name:        "expire",
					Description: `(Optional) Time, in seconds, after which the zone data on a secondary name server can no longer be considered authoritative because all refresh and retry attempts made during the period have failed. After the expiry period, the secondary server stops serving the zone. Typically one week. Not used by the primary server.`,
				},
				resource.Attribute{
					Name:        "minimum",
					Description: `(Optional) Default time to live (TTL) for all records in the zone. Can be overridden for individual records.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `(Optional) Time to Live (TTL), in seconds, for the record. TTL is the time for which the record must be cached by DNS proxies. The specified TTL is applied to all the resource records that are of the same record type and belong to the specified domain name. For example, if you add an address record, with a TTL of 36000, to the domain name example.com, the TTLs of all the address records of example.com are changed to 36000. If the TTL is not specified, the Citrix ADC uses either the DNS zone's minimum TTL or, if the SOA record is not available on the appliance, the default value of 3600.`,
				},
				resource.Attribute{
					Name:        "ecssubnet",
					Description: `(Optional) Subnet for which the cached SOA record need to be removed.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) Type of records to display. Available settings function as follows:`,
				},
				resource.Attribute{
					Name:        "nodeid",
					Description: `(Optional) Unique number that identifies the cluster node. ## Attribute Reference In addition to the arguments, the following attributes are available:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of the dnssoarec. It has the same value as the ` + "`" + `domain` + "`" + ` attribute. ## Import A dnssoarec can be imported using its domain, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import citrixadc_dnssoarec.tf_dnssoarec hello.com ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id of the dnssoarec. It has the same value as the ` + "`" + `domain` + "`" + ` attribute. ## Import A dnssoarec can be imported using its domain, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import citrixadc_dnssoarec.tf_dnssoarec hello.com ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "citrixadc_filterpolicy",
			Category:         "Filter",
			ShortDescription: ``,
			Description: `

The filterpolicy resource is used to create filter policies.


`,
			Keywords: []string{
				"filter",
				"filterpolicy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name for the filtering action. Must begin with a letter, number, or the underscore character (\_). Other characters allowed, after the first character, are the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), and colon (:) characters. Choose a name that helps identify the type of action. The name cannot be updated after the policy is created. CLI Users: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my policy" or 'my policy').`,
				},
				resource.Attribute{
					Name:        "rule",
					Description: `(Optional) Citrix ADC classic expression specifying the type of connections that match this policy.`,
				},
				resource.Attribute{
					Name:        "reqaction",
					Description: `(Optional) Name of the action to be performed on requests that match the policy. Cannot be specified if the rule includes condition to be evaluated for responses.`,
				},
				resource.Attribute{
					Name:        "resaction",
					Description: `(Optional) The action to be performed on the response. The string value can be a filter action created filter action or a built-in action. ## Attribute Reference In addition to the arguments, the following attributes are available:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of the filterpolicy. It has the same value as the ` + "`" + `name` + "`" + ` attribute. ## Import A filterpolicy can be imported using its name, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import citrixadc_filterpolicy.tf_filterpolicy tf_filterpolicy ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id of the filterpolicy. It has the same value as the ` + "`" + `name` + "`" + ` attribute. ## Import A filterpolicy can be imported using its name, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import citrixadc_filterpolicy.tf_filterpolicy tf_filterpolicy ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "citrixadc_gslbservice",
			Category:         "GSLB",
			ShortDescription: ``,
			Description: `

This resource is used to manage Global Server Load Balancing service.


`,
			Keywords: []string{
				"gslb",
				"gslbservice",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "servicename",
					Description: `(Optional) Name for the GSLB service.`,
				},
				resource.Attribute{
					Name:        "cnameentry",
					Description: `(Optional) Canonical name of the GSLB service. Used in CNAME-based GSLB.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `(Optional) IP address for the GSLB service. Should represent a load balancing, content switching, or VPN virtual server on the Citrix ADC, or the IP address of another load balancing device.`,
				},
				resource.Attribute{
					Name:        "servername",
					Description: `(Optional) Name of the server hosting the GSLB service.`,
				},
				resource.Attribute{
					Name:        "servicetype",
					Description: `(Optional) Type of service to create. Possible values: [ HTTP, FTP, TCP, UDP, SSL, SSL\_BRIDGE, SSL\_TCP, NNTP, ANY, SIP\_UDP, SIP\_TCP, SIP\_SSL, RADIUS, RDP, RTSP, MYSQL, MSSQL, ORACLE ]`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional) Port on which the load balancing entity represented by this GSLB service listens.`,
				},
				resource.Attribute{
					Name:        "publicip",
					Description: `(Optional) The public IP address that a NAT device translates to the GSLB service's private IP address. Optional.`,
				},
				resource.Attribute{
					Name:        "publicport",
					Description: `(Optional) The public port associated with the GSLB service's public IP address. The port is mapped to the service's private port number. Applicable to the local GSLB service.`,
				},
				resource.Attribute{
					Name:        "maxclient",
					Description: `(Optional) The maximum number of open connections that the service can support at any given time. A GSLB service whose connection count reaches the maximum is not considered when a GSLB decision is made, until the connection count drops below the maximum.`,
				},
				resource.Attribute{
					Name:        "healthmonitor",
					Description: `(Optional) Monitor the health of the GSLB service. Possible values: [ YES, NO ]`,
				},
				resource.Attribute{
					Name:        "sitename",
					Description: `(Optional) Name of the GSLB site to which the service belongs.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `(Optional) Enable or disable the service. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "cip",
					Description: `(Optional) In the request that is forwarded to the GSLB service, insert a header that stores the client's IP address. Client IP header insertion is used in connection-proxy based site persistence. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "cipheader",
					Description: `(Optional) Name for the HTTP header that stores the client's IP address. Used with the Client IP option. If client IP header insertion is enabled on the service and a name is not specified for the header, the Citrix ADC uses the name specified by the cipHeader parameter in the set ns param command or, in the GUI, the Client IP Header parameter in the Configure HTTP Parameters dialog box.`,
				},
				resource.Attribute{
					Name:        "sitepersistence",
					Description: `(Optional) Use cookie-based site persistence. Applicable only to HTTP and SSL GSLB services. Possible values: [ ConnectionProxy, HTTPRedirect, NONE ]`,
				},
				resource.Attribute{
					Name:        "cookietimeout",
					Description: `(Optional) Timeout value, in minutes, for the cookie, when cookie based site persistence is enabled.`,
				},
				resource.Attribute{
					Name:        "siteprefix",
					Description: `(Optional) The site's prefix string. When the service is bound to a GSLB virtual server, a GSLB site domain is generated internally for each bound service-domain pair by concatenating the site prefix of the service and the name of the domain. If the special string NONE is specified, the site-prefix string is unset. When implementing HTTP redirect site persistence, the Citrix ADC redirects GSLB requests to GSLB services by using their site domains.`,
				},
				resource.Attribute{
					Name:        "clttimeout",
					Description: `(Optional) Idle time, in seconds, after which a client connection is terminated. Applicable if connection proxy based site persistence is used.`,
				},
				resource.Attribute{
					Name:        "svrtimeout",
					Description: `(Optional) Idle time, in seconds, after which a server connection is terminated. Applicable if connection proxy based site persistence is used.`,
				},
				resource.Attribute{
					Name:        "maxbandwidth",
					Description: `(Optional) Integer specifying the maximum bandwidth allowed for the service. A GSLB service whose bandwidth reaches the maximum is not considered when a GSLB decision is made, until its bandwidth consumption drops below the maximum.`,
				},
				resource.Attribute{
					Name:        "downstateflush",
					Description: `(Optional) Flush all active transactions associated with the GSLB service when its state transitions from UP to DOWN. Do not enable this option for services that must complete their transactions. Applicable if connection proxy based site persistence is used. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "maxaaausers",
					Description: `(Optional) Maximum number of SSL VPN users that can be logged on concurrently to the VPN virtual server that is represented by this GSLB service. A GSLB service whose user count reaches the maximum is not considered when a GSLB decision is made, until the count drops below the maximum.`,
				},
				resource.Attribute{
					Name:        "monthreshold",
					Description: `(Optional) Monitoring threshold value for the GSLB service. If the sum of the weights of the monitors that are bound to this GSLB service and are in the UP state is not equal to or greater than this threshold value, the service is marked as DOWN.`,
				},
				resource.Attribute{
					Name:        "hashid",
					Description: `(Optional) Unique hash identifier for the GSLB service, used by hash based load balancing methods.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `(Optional) Any comments that you might want to associate with the GSLB service.`,
				},
				resource.Attribute{
					Name:        "appflowlog",
					Description: `(Optional) Enable logging appflow flow information. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "naptrreplacement",
					Description: `(Optional) The replacement domain name for this NAPTR.`,
				},
				resource.Attribute{
					Name:        "naptrorder",
					Description: `(Optional) An integer specifying the order in which the NAPTR records MUST be processed in order to accurately represent the ordered list of Rules. The ordering is from lowest to highest.`,
				},
				resource.Attribute{
					Name:        "naptrservices",
					Description: `(Optional) Service Parameters applicable to this delegation path.`,
				},
				resource.Attribute{
					Name:        "naptrdomainttl",
					Description: `(Optional) Modify the TTL of the internally created naptr domain.`,
				},
				resource.Attribute{
					Name:        "naptrpreference",
					Description: `(Optional) An integer specifying the preference of this NAPTR among NAPTR records having same order. lower the number, higher the preference.`,
				},
				resource.Attribute{
					Name:        "ipaddress",
					Description: `(Optional) The new IP address of the service.`,
				},
				resource.Attribute{
					Name:        "viewname",
					Description: `(Optional) Name of the DNS view of the service. A DNS view is used in global server load balancing (GSLB) to return a predetermined IP address to a specific group of clients, which are identified by using a DNS policy.`,
				},
				resource.Attribute{
					Name:        "viewip",
					Description: `(Optional) IP address to be used for the given view.`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `(Optional) Weight to assign to the monitor-service binding. A larger number specifies a greater weight. Contributes to the monitoring threshold, which determines the state of the service.`,
				},
				resource.Attribute{
					Name:        "monitornamesvc",
					Description: `(Optional) Name of the monitor to bind to the service.`,
				},
				resource.Attribute{
					Name:        "delay",
					Description: `(Optional) Delay to wait for service to be disabled.`,
				},
				resource.Attribute{
					Name:        "lbmonitorbinding",
					Description: `(Optional) A set of lb monitor blocks. Documented below Lb monitor supports the following:`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `(Optional) Weight to assign to the monitor-service binding. A larger number specifies a greater weight. Contributes to the monitoring threshold, which determines the state of the service.`,
				},
				resource.Attribute{
					Name:        "monitor_name",
					Description: `(Optional) Monitor name.`,
				},
				resource.Attribute{
					Name:        "monstate",
					Description: `(Optional) State of the monitor bound to gslb service. Possible values: [ ENABLED, DISABLED ] ## Attribute Reference In addition to the arguments, the following attributes are available:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of the glsbservice. It has the same value as the ` + "`" + `servicename` + "`" + ` attribute. ## Import An instance of the resource can be imported using its name, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import citrixadc_gslbservice.tf_gslbservice tf_gslbservice ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id of the glsbservice. It has the same value as the ` + "`" + `servicename` + "`" + ` attribute. ## Import An instance of the resource can be imported using its name, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import citrixadc_gslbservice.tf_gslbservice tf_gslbservice ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "citrixadc_gslbsite",
			Category:         "GSLB",
			ShortDescription: ``,
			Description: `

This resource is used to manage Global Service Load Balancing site.


`,
			Keywords: []string{
				"gslb",
				"gslbsite",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "sitename",
					Description: `(Optional) Name for the GSLB site.`,
				},
				resource.Attribute{
					Name:        "sitetype",
					Description: `(Optional) Type of site to create. If the type is not specified, the appliance automatically detects and sets the type on the basis of the IP address being assigned to the site. If the specified site IP address is owned by the appliance (for example, a MIP address or SNIP address), the site is a local site. Otherwise, it is a remote site. Possible values: [ REMOTE, LOCAL ]`,
				},
				resource.Attribute{
					Name:        "siteipaddress",
					Description: `(Optional) IP address for the GSLB site. The GSLB site uses this IP address to communicate with other GSLB sites. For a local site, use any IP address that is owned by the appliance (for example, a SNIP or MIP address, or the IP address of the ADNS service).`,
				},
				resource.Attribute{
					Name:        "publicip",
					Description: `(Optional) Public IP address for the local site. Required only if the appliance is deployed in a private address space and the site has a public IP address hosted on an external firewall or a NAT device.`,
				},
				resource.Attribute{
					Name:        "metricexchange",
					Description: `(Optional) Exchange metrics with other sites. Metrics are exchanged by using Metric Exchange Protocol (MEP). The appliances in the GSLB setup exchange health information once every second. If you disable metrics exchange, you can use only static load balancing methods (such as round robin, static proximity, or the hash-based methods), and if you disable metrics exchange when a dynamic load balancing method (such as least connection) is in operation, the appliance falls back to round robin. Also, if you disable metrics exchange, you must use a monitor to determine the state of GSLB services. Otherwise, the service is marked as DOWN. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "nwmetricexchange",
					Description: `(Optional) Exchange, with other GSLB sites, network metrics such as round-trip time (RTT), learned from communications with various local DNS (LDNS) servers used by clients. RTT information is used in the dynamic RTT load balancing method, and is exchanged every 5 seconds. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "sessionexchange",
					Description: `(Optional) Exchange persistent session entries with other GSLB sites every five seconds. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "triggermonitor",
					Description: `(Optional) Specify the conditions under which the GSLB service must be monitored by a monitor, if one is bound. Available settings function as follows:`,
				},
				resource.Attribute{
					Name:        "parentsite",
					Description: `(Optional) Parent site of the GSLB site, in a parent-child topology.`,
				},
				resource.Attribute{
					Name:        "clip",
					Description: `(Optional) Cluster IP address. Specify this parameter to connect to the remote cluster site for GSLB auto-sync. Note: The cluster IP address is defined when creating the cluster.`,
				},
				resource.Attribute{
					Name:        "publicclip",
					Description: `(Optional) IP address to be used to globally access the remote cluster when it is deployed behind a NAT. It can be same as the normal cluster IP address.`,
				},
				resource.Attribute{
					Name:        "naptrreplacementsuffix",
					Description: `(Optional) The naptr replacement suffix configured here will be used to construct the naptr replacement field in NAPTR record. ## Attribute Reference In addition to the arguments, the following attributes are available:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of the gslbsite. It has the same value as the ` + "`" + `sitename` + "`" + ` attribute. ## Import An instance of the resource can be imported using its sitename, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import citrixadc_gslbsite.tf_site_local tf_site_local ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id of the gslbsite. It has the same value as the ` + "`" + `sitename` + "`" + ` attribute. ## Import An instance of the resource can be imported using its sitename, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import citrixadc_gslbsite.tf_site_local tf_site_local ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "citrixadc_gslbvserver",
			Category:         "GSLB",
			ShortDescription: ``,
			Description: `

This resource is used to manage Global Service Load Balancing vserver.


`,
			Keywords: []string{
				"gslb",
				"gslbvserver",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name for the GSLB virtual server.`,
				},
				resource.Attribute{
					Name:        "servicetype",
					Description: `(Optional) Protocol used by services bound to the virtual server. Possible values: [ HTTP, FTP, TCP, UDP, SSL, SSL\_BRIDGE, SSL\_TCP, NNTP, ANY, SIP\_UDP, SIP\_TCP, SIP\_SSL, RADIUS, RDP, RTSP, MYSQL, MSSQL, ORACLE ]`,
				},
				resource.Attribute{
					Name:        "iptype",
					Description: `(Optional) The IP type for this GSLB vserver. Possible values: [ IPV4, IPV6 ]`,
				},
				resource.Attribute{
					Name:        "dnsrecordtype",
					Description: `(Optional) DNS record type to associate with the GSLB virtual server's domain name. Possible values: [ A, AAAA, CNAME, NAPTR ]`,
				},
				resource.Attribute{
					Name:        "lbmethod",
					Description: `(Optional) Load balancing method for the GSLB virtual server. Possible values: [ ROUNDROBIN, LEASTCONNECTION, LEASTRESPONSETIME, SOURCEIPHASH, LEASTBANDWIDTH, LEASTPACKETS, STATICPROXIMITY, RTT, CUSTOMLOAD, API ]`,
				},
				resource.Attribute{
					Name:        "backupsessiontimeout",
					Description: `(Optional) A non zero value enables the feature whose minimum value is 2 minutes. The feature can be disabled by setting the value to zero. The created session is in effect for a specific client per domain.`,
				},
				resource.Attribute{
					Name:        "backuplbmethod",
					Description: `(Optional) Backup load balancing method. Becomes operational if the primary load balancing method fails or cannot be used. Valid only if the primary method is based on either round-trip time (RTT) or static proximity. Possible values: [ ROUNDROBIN, LEASTCONNECTION, LEASTRESPONSETIME, SOURCEIPHASH, LEASTBANDWIDTH, LEASTPACKETS, STATICPROXIMITY, RTT, CUSTOMLOAD, API ]`,
				},
				resource.Attribute{
					Name:        "netmask",
					Description: `(Optional) IPv4 network mask for use in the SOURCEIPHASH load balancing method.`,
				},
				resource.Attribute{
					Name:        "v6netmasklen",
					Description: `(Optional) Number of bits to consider, in an IPv6 source IP address, for creating the hash that is required by the SOURCEIPHASH load balancing method.`,
				},
				resource.Attribute{
					Name:        "tolerance",
					Description: `(Optional) Site selection tolerance, in milliseconds, for implementing the RTT load balancing method. If a site's RTT deviates from the lowest RTT by more than the specified tolerance, the site is not considered when the Citrix ADC makes a GSLB decision. The appliance implements the round robin method of global server load balancing between sites whose RTT values are within the specified tolerance. If the tolerance is 0 (zero), the appliance always sends clients the IP address of the site with the lowest RTT.`,
				},
				resource.Attribute{
					Name:        "persistencetype",
					Description: `(Optional) Use source IP address based persistence for the virtual server. After the load balancing method selects a service for the first packet, the IP address received in response to the DNS query is used for subsequent requests from the same client. Possible values: [ SOURCEIP, NONE ]`,
				},
				resource.Attribute{
					Name:        "persistenceid",
					Description: `(Optional) The persistence ID for the GSLB virtual server. The ID is a positive integer that enables GSLB sites to identify the GSLB virtual server, and is required if source IP address based or spill over based persistence is enabled on the virtual server.`,
				},
				resource.Attribute{
					Name:        "persistmask",
					Description: `(Optional) The optional IPv4 network mask applied to IPv4 addresses to establish source IP address based persistence.`,
				},
				resource.Attribute{
					Name:        "v6persistmasklen",
					Description: `(Optional) Number of bits to consider in an IPv6 source IP address when creating source IP address based persistence sessions.`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `(Optional) Idle time, in minutes, after which a persistence entry is cleared.`,
				},
				resource.Attribute{
					Name:        "edr",
					Description: `(Optional) Send clients an empty DNS response when the GSLB virtual server is DOWN. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "ecs",
					Description: `(Optional) If enabled, respond with EDNS Client Subnet (ECS) option in the response for a DNS query with ECS. The ECS address will be used for persistence and spillover persistence (if enabled) instead of the LDNS address. Persistence mask is ignored if ECS is enabled. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "ecsaddrvalidation",
					Description: `(Optional) Validate if ECS address is a private or unroutable address and in such cases, use the LDNS IP. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "mir",
					Description: `(Optional) Include multiple IP addresses in the DNS responses sent to clients. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "disableprimaryondown",
					Description: `(Optional) Continue to direct traffic to the backup chain even after the primary GSLB virtual server returns to the UP state. Used when spillover is configured for the virtual server. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "dynamicweight",
					Description: `(Optional) Specify if the appliance should consider the service count, service weights, or ignore both when using weight-based load balancing methods. The state of the number of services bound to the virtual server help the appliance to select the service. Possible values : [ SERVICECOUNT, SERVICEWEIGHT, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `(Optional) State of the GSLB virtual server. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "considereffectivestate",
					Description: `(Optional) If the primary state of all bound GSLB services is DOWN, consider the effective states of all the GSLB services, obtained through the Metrics Exchange Protocol (MEP), when determining the state of the GSLB virtual server. To consider the effective state, set the parameter to STATE\_ONLY. To disregard the effective state, set the parameter to NONE. The effective state of a GSLB service is the ability of the corresponding virtual server to serve traffic. The effective state of the load balancing virtual server, which is transferred to the GSLB service, is UP even if only one virtual server in the backup chain of virtual servers is in the UP state. Possible values: [ NONE, STATE_ONLY ]`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `(Optional) Any comments that you might want to associate with the GSLB virtual server.`,
				},
				resource.Attribute{
					Name:        "somethod",
					Description: `(Optional) Type of threshold that, when exceeded, triggers spillover. Possible values: [ CONNECTION, DYNAMICCONNECTION, BANDWIDTH, HEALTH, NONE ]`,
				},
				resource.Attribute{
					Name:        "sopersistence",
					Description: `(Optional) If spillover occurs, maintain source IP address based persistence for both primary and backup GSLB virtual servers.Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "sopersistencetimeout",
					Description: `(Optional) Timeout for spillover persistence, in minutes.`,
				},
				resource.Attribute{
					Name:        "sothreshold",
					Description: `(Optional) Threshold at which spillover occurs. Specify an integer for the CONNECTION spillover method, a bandwidth value in kilobits per second for the BANDWIDTH method (do not enter the units), or a percentage for the HEALTH method (do not enter the percentage symbol).`,
				},
				resource.Attribute{
					Name:        "sobackupaction",
					Description: `(Optional) Action to be performed if spillover is to take effect, but no backup chain to spillover is usable or exists. Possible values: [ DROP, ACCEPT, REDIRECT ]`,
				},
				resource.Attribute{
					Name:        "appflowlog",
					Description: `(Optional) Enable logging appflow flow information. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "backupvserver",
					Description: `(Optional) Name of the backup GSLB virtual server to which the appliance should to forward requests if the status of the primary GSLB virtual server is down or exceeds its spillover threshold.`,
				},
				resource.Attribute{
					Name:        "servicename",
					Description: `(Optional) Name of the GSLB service for which to change the weight.`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `(Optional) Weight to assign to the GSLB service.`,
				},
				resource.Attribute{
					Name:        "domainname",
					Description: `(Optional) Domain name for which to change the time to live (TTL) and/or backup service IP address.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `(Optional) Time to live (TTL) for the domain.`,
				},
				resource.Attribute{
					Name:        "backupip",
					Description: `(Optional) The IP address of the backup service for the specified domain name. Used when all the services bound to the domain are down, or when the backup chain of virtual servers is down.`,
				},
				resource.Attribute{
					Name:        "cookiedomain",
					Description: `(Optional) The cookie domain for the GSLB site. Used when inserting the GSLB site cookie in the HTTP response.`,
				},
				resource.Attribute{
					Name:        "cookietimeout",
					Description: `(Optional) Timeout, in minutes, for the GSLB site cookie.`,
				},
				resource.Attribute{
					Name:        "sitedomainttl",
					Description: `(Optional) TTL, in seconds, for all internally created site domains (created when a site prefix is configured on a GSLB service) that are associated with this virtual server.`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `(Optional) A set of domain binding blocks. Documented below.`,
				},
				resource.Attribute{
					Name:        "service",
					Description: `(Optional) A set of GSLB service biding blocks. Documented below. A domain binding supports the following:`,
				},
				resource.Attribute{
					Name:        "backupipflag",
					Description: `(Optional) The IP address of the backup service for the specified domain name. Used when all the services bound to the domain are down, or when the backup chain of virtual servers is down.`,
				},
				resource.Attribute{
					Name:        "cookietimeout",
					Description: `(Optional) Timeout, in minutes, for the GSLB site cookie.`,
				},
				resource.Attribute{
					Name:        "backupip",
					Description: `(Optional) The IP address of the backup service for the specified domain name. Used when all the services bound to the domain are down, or when the backup chain of virtual servers is down.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `(Optional) Time to live (TTL) for the domain.`,
				},
				resource.Attribute{
					Name:        "domainname",
					Description: `(Optional) Domain name for which to change the time to live (TTL) and/or backup service IP address.`,
				},
				resource.Attribute{
					Name:        "sitedomainttl",
					Description: `(Optional) TTL, in seconds, for all internally created site domains (created when a site prefix is configured on a GSLB service) that are associated with this virtual server.`,
				},
				resource.Attribute{
					Name:        "cookiedomain",
					Description: `(Optional) The cookie domain for the GSLB site. Used when inserting the GSLB site cookie in the HTTP response.`,
				},
				resource.Attribute{
					Name:        "cookiedomainflag",
					Description: `(Optional) The cookie domain for the GSLB site. Used when inserting the GSLB site cookie in the HTTP response.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the virtual server on which to perform the binding operation. A GSLB service binding supports the following:`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `(Optional) Weight to assign to the GSLB service.`,
				},
				resource.Attribute{
					Name:        "servicename",
					Description: `(Optional) Name of the GSLB service for which to change the weight.`,
				},
				resource.Attribute{
					Name:        "domainname",
					Description: `(Optional) Domain name for which to change the time to live (TTL) and/or backup service IP address. ## Attribute Reference In addition to the arguments, the following attributes are available:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of the gslbvserver. It has the same value as the ` + "`" + `name` + "`" + ` attribute. ## Import An instance of the resource can be imported using its name, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import citrixadc_gslbvserver.tf_gslbvserver GSLB_East_Coast_Vserver ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id of the gslbvserver. It has the same value as the ` + "`" + `name` + "`" + ` attribute. ## Import An instance of the resource can be imported using its name, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import citrixadc_gslbvserver.tf_gslbvserver GSLB_East_Coast_Vserver ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "citrixadc_inat",
			Category:         "Network",
			ShortDescription: ``,
			Description: `

The inat resource is used to create inbound nat resources.


`,
			Keywords: []string{
				"network",
				"inat",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name for the Inbound NAT (INAT) entry.`,
				},
				resource.Attribute{
					Name:        "publicip",
					Description: `(Optional) Public IP address of packets received on the Citrix ADC. Can be aNetScaler-owned VIP or VIP6 address.`,
				},
				resource.Attribute{
					Name:        "privateip",
					Description: `(Optional) IP address of the server to which the packet is sent by the Citrix ADC. Can be an IPv4 or IPv6 address.`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `(Optional) Stateless translation. Possible values: [ STATELESS ]`,
				},
				resource.Attribute{
					Name:        "tcpproxy",
					Description: `(Optional) Enable TCP proxy, which enables the Citrix ADC to optimize the RNAT TCP traffic by using Layer 4 features. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "ftp",
					Description: `(Optional) Enable the FTP protocol on the server for transferring files between the client and the server. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "tftp",
					Description: `(Optional) To enable/disable TFTP (Default DISABLED). Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "usip",
					Description: `(Optional) Enable the Citrix ADC to retain the source IP address of packets before sending the packets to the server. Possible values: [ on, off ]`,
				},
				resource.Attribute{
					Name:        "usnip",
					Description: `(Optional) Enable the Citrix ADC to use a SNIP address as the source IP address of packets before sending the packets to the server. Possible values: [ on, off ]`,
				},
				resource.Attribute{
					Name:        "proxyip",
					Description: `(Optional) Unique IP address used as the source IP address in packets sent to the server. Must be a MIP or SNIP address.`,
				},
				resource.Attribute{
					Name:        "useproxyport",
					Description: `(Optional) Enable the Citrix ADC to proxy the source port of packets before sending the packets to the server. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "td",
					Description: `(Optional) Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0. ## Attribute Reference In addition to the arguments, the following attributes are available:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of the inat resource. It has the same value as the ` + "`" + `name` + "`" + ` attribute. ## Import An inat can be imported using its name, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import citrixadc_inat.inat_resource ip4ip ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id of the inat resource. It has the same value as the ` + "`" + `name` + "`" + ` attribute. ## Import An inat can be imported using its name, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import citrixadc_inat.inat_resource ip4ip ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "citrixadc_installer",
			Category:         "Utility",
			ShortDescription: ``,
			Description: `

The installer resource is used to install updated version build onto a target ADC.


`,
			Keywords: []string{
				"utility",
				"installer",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "url",
					Description: `(Optional) Url for the build file. Must be in the following formats: http://[user]:[password]@host/path/to/file https://[user]:[password]@host/path/to/file sftp://[user]:[password]@host/path/to/file scp://[user]:[password]@host/path/to/file ftp://[user]:[password]@host/path/to/file file://path/to/file.`,
				},
				resource.Attribute{
					Name:        "y",
					Description: `(Optional) Do not prompt for yes/no before rebooting.`,
				},
				resource.Attribute{
					Name:        "l",
					Description: `(Optional) Use this flag to enable callhome.`,
				},
				resource.Attribute{
					Name:        "enhancedupgrade",
					Description: `(Optional) Use this flag for upgrading from/to enhancement mode.`,
				},
				resource.Attribute{
					Name:        "resizeswapvar",
					Description: `(Optional) Use this flag to change swap size on ONLY 64bit nCore/MCNS/VMPE systems NON-VPX systems.`,
				},
				resource.Attribute{
					Name:        "wait_until_reachable",
					Description: `(Optional) Boolean value to determine if the resource should wait for the ADC to become reachable after the build is installed.`,
				},
				resource.Attribute{
					Name:        "reachable_timeout",
					Description: `(Optional) Time period to wait untill the target ADC becomes reachable. Default value "10m"`,
				},
				resource.Attribute{
					Name:        "reachable_poll_delay",
					Description: `(Optional) Time delay before the first poll. Must be sufficiently large to allow the ADC to start reboot. Default value "60s".`,
				},
				resource.Attribute{
					Name:        "reachable_poll_interval",
					Description: `(Optional) Time interval between polls for reachability. Default value "60s".`,
				},
				resource.Attribute{
					Name:        "reachable_poll_timeout",
					Description: `(Optional) Time period to wait before the http poll request times out. Default value "20s". ## Attribute Reference In addition to the arguments, the following attributes are available:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of the installer. It is a random string prefixed with "tf-installer-".`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id of the installer. It is a random string prefixed with "tf-installer-".`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "citrixadc_interface",
			Category:         "Network",
			ShortDescription: ``,
			Description: `

The interface resource is used to create network interfaces.


`,
			Keywords: []string{
				"network",
				"interface",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "interface_id",
					Description: `(Required) Interface number, in C/U format, where C can take one of the following values: 0 - Indicates a management interface. 1 - Indicates a 1 Gbps port. 10 - Indicates a 10 Gbps port. LA - Indicates a link aggregation port. LO - Indicates a loop back port. U is a unique integer for representing an interface in a particular port group.`,
				},
				resource.Attribute{
					Name:        "speed",
					Description: `(Optional) Ethernet speed of the interface, in Mbps. Notes: If you set the speed as AUTO, the Citrix ADC attempts to auto-negotiate or auto-sense the link speed of the interface when it is UP. You must enable auto negotiation on the interface. If you set a speed other than AUTO, you must specify the same speed for the peer network device. Mismatched speed and duplex settings between the peer devices of a link lead to link errors, packet loss, and other errors. Some interfaces do not support certain speeds. If you specify an unsupported speed, an error message appears. Possible values: [ AUTO, 10, 100, 1000, 10000, 25000, 40000, 50000, 100000 ]`,
				},
				resource.Attribute{
					Name:        "duplex",
					Description: `(Optional) The duplex mode for the interface. Notes: If you set the duplex mode to AUTO, the Citrix ADC attempts to auto-negotiate the duplex mode of the interface when it is UP. You must enable auto negotiation on the interface. If you set a duplex mode other than AUTO, you must specify the same duplex mode for the peer network device. Mismatched speed and duplex settings between the peer devices of a link lead to link errors, packet loss, and other errors. Possible values: [ AUTO, HALF, FULL ]`,
				},
				resource.Attribute{
					Name:        "flowctl",
					Description: `(Optional) 802.3x flow control setting for the interface. The 802.3x specification does not define flow control for 10 Mbps and 100 Mbps speeds, but if a Gigabit Ethernet interface operates at those speeds, the flow control settings can be applied. The flow control setting that is finally applied to an interface depends on auto-negotiation. With the ON option, the peer negotiates the flow control, but the appliance then forces two-way flow control for the interface. Possible values: [ OFF, RX, TX, RXTX, ON ]`,
				},
				resource.Attribute{
					Name:        "autoneg",
					Description: `(Optional) Auto-negotiation state of the interface. With the ENABLED setting, the Citrix ADC auto-negotiates the speed and duplex settings with the peer network device on the link. The Citrix ADC appliance auto-negotiates the settings of only those parameters (speed or duplex mode) for which the value is set as AUTO. Possible values: [ DISABLED, ENABLED ]`,
				},
				resource.Attribute{
					Name:        "hamonitor",
					Description: `(Optional) In a High Availability (HA) configuration, monitor the interface for failure events. In an HA configuration, an interface that has HA MON enabled and is not bound to any Failover Interface Set (FIS), is a critical interface. Failure or disabling of any critical interface triggers HA failover. Possible values: [ on, off ]`,
				},
				resource.Attribute{
					Name:        "haheartbeat",
					Description: `(Optional) In a High Availability (HA) or Cluster configuration, configure the interface for sending heartbeats. In an HA or Cluster configuration, an interface that has HA Heartbeat disabled should not send the heartbeats. Possible values: [ on, off ]`,
				},
				resource.Attribute{
					Name:        "mtu",
					Description: `(Optional) The maximum transmission unit (MTU) is the largest packet size, measured in bytes excluding 14 bytes ethernet header and 4 bytes crc, that can be transmitted and received by this interface. Default value of MTU is 1500 on all the interface of Citrix ADC any value configured more than 1500 on the interface will make the interface as jumbo enabled. In case of cluster backplane interface MTU value will be changed to 1514 by default, user has to change the backplane interface value to maximum mtu configured on any of the interface in cluster system plus 14 bytes more for backplane interface if Jumbo is enabled on any of the interface in a cluster system. Changing the backplane will bring back the MTU of backplane interface to default value of 1500. If a channel is configured as backplane then the same holds true for channel as well as member interfaces.`,
				},
				resource.Attribute{
					Name:        "ringsize",
					Description: `(Optional) The receive ringsize of the interface. A higher number provides more number of buffers in handling incoming traffic.`,
				},
				resource.Attribute{
					Name:        "ringtype",
					Description: `(Optional) The receive ringtype of the interface (Fixed or Elastic). A fixed ring type pre-allocates configured number of buffers irrespective of traffic rate. In contrast, an elastic ring, expands and shrinks based on incoming traffic rate. Possible values: [ Elastic, Fixed ]`,
				},
				resource.Attribute{
					Name:        "tagall",
					Description: `(Optional) Add a four-byte 802.1q tag to every packet sent on this interface. The ON setting applies the tag for this interface's native VLAN. OFF applies the tag for all VLANs other than the native VLAN. Possible values: [ on, off ]`,
				},
				resource.Attribute{
					Name:        "trunk",
					Description: `(Optional) This argument is deprecated by tagall. Possible values: [ on, off ]`,
				},
				resource.Attribute{
					Name:        "trunkmode",
					Description: `(Optional) Accept and send 802.1q VLAN tagged packets, based on Allowed Vlan List of this interface. Possible values: [ on, off ]`,
				},
				resource.Attribute{
					Name:        "lacpmode",
					Description: `(Optional) Bind the interface to a LA channel created by the Link Aggregation control protocol (LACP). Available settings function as follows:`,
				},
				resource.Attribute{
					Name:        "lacpkey",
					Description: `(Optional) Integer identifying the LACP LA channel to which the interface is to be bound. For an LA channel of the Citrix ADC, this digit specifies the variable x of an LA channel in LA/x notation, where x can range from 1 to 8. For example, if you specify 3 as the LACP key for an LA channel, the interface is bound to the LA channel LA/3. For an LA channel of a cluster configuration, this digit specifies the variable y of a cluster LA channel in CLA/(y-4) notation, where y can range from 5 to 8. For example, if you specify 6 as the LACP key for a cluster LA channel, the interface is bound to the cluster LA channel CLA/2.`,
				},
				resource.Attribute{
					Name:        "lagtype",
					Description: `(Optional) Type of entity (Citrix ADC or cluster configuration) for which to create the channel. Possible values: [ NODE, CLUSTER ]`,
				},
				resource.Attribute{
					Name:        "lacppriority",
					Description: `(Optional) LACP port priority, expressed as an integer. The lower the number, the higher the priority. The Citrix ADC limits the number of interfaces in an LA channel to sixteen.`,
				},
				resource.Attribute{
					Name:        "lacptimeout",
					Description: `(Optional) Interval at which the Citrix ADC sends LACPDU messages to the peer device on the LA channel. Available settings function as follows: LONG - 30 seconds. SHORT - 1 second. Possible values: [ LONG, SHORT ]`,
				},
				resource.Attribute{
					Name:        "ifalias",
					Description: `(Optional) Alias name for the interface. Used only to enhance readability. To perform any operations, you have to specify the interface ID.`,
				},
				resource.Attribute{
					Name:        "throughput",
					Description: `(Optional) Low threshold value for the throughput of the interface, in Mbps. In an HA configuration, failover is triggered if the interface has HA MON enabled and the throughput is below the specified the threshold.`,
				},
				resource.Attribute{
					Name:        "linkredundancy",
					Description: `(Optional) Link Redundancy for Cluster LAG. Possible values: [ on, off ]`,
				},
				resource.Attribute{
					Name:        "bandwidthhigh",
					Description: `(Optional) High threshold value for the bandwidth usage of the interface, in Mbps. The Citrix ADC generates an SNMP trap message when the bandwidth usage of the interface is greater than or equal to the specified high threshold value.`,
				},
				resource.Attribute{
					Name:        "bandwidthnormal",
					Description: `(Optional) Normal threshold value for the bandwidth usage of the interface, in Mbps. When the bandwidth usage of the interface becomes less than or equal to the specified normal threshold after exceeding the high threshold, the Citrix ADC generates an SNMP trap message to indicate that the bandwidth usage has returned to normal.`,
				},
				resource.Attribute{
					Name:        "lldpmode",
					Description: `(Optional) Link Layer Discovery Protocol (LLDP) mode for an interface. The resultant LLDP mode of an interface depends on the LLDP mode configured at the global and the interface levels. Possible values: [ NONE, TRANSMITTER, RECEIVER, TRANSCEIVER ]`,
				},
				resource.Attribute{
					Name:        "lrsetpriority",
					Description: `(Optional) LRSET port priority, expressed as an integer ranging from 1 to 1024. The highest priority is 1. The Citrix ADC limits the number of interfaces in an LRSET to 8. Within a LRSET the highest LR Priority Interface is considered as the first candidate for the Active interface, if the interface is UP. ## Attribute Reference In addition to the arguments, the following attributes are available:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of the interface. It has the same value as the ` + "`" + `interface_id` + "`" + ` attribute. ## Import A interface can be imported using its interface\_id, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import citrixadc_interface.tf_interface 1/1 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id of the interface. It has the same value as the ` + "`" + `interface_id` + "`" + ` attribute. ## Import A interface can be imported using its interface\_id, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import citrixadc_interface.tf_interface 1/1 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "citrixadc_ipset",
			Category:         "Network",
			ShortDescription: ``,
			Description: `

The ipset resource is used to create ip set resources.


`,
			Keywords: []string{
				"network",
				"ipset",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name for the IP set. Cannot be changed after the IP set is created. Choose a name that helps identify the IP set.`,
				},
				resource.Attribute{
					Name:        "td",
					Description: `(Optional) Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.`,
				},
				resource.Attribute{
					Name:        "nsipbinding",
					Description: `(Optional) A set of ipv4 addresses that will be bound to this ipset.`,
				},
				resource.Attribute{
					Name:        "nsip6binding",
					Description: `(Optional) A set of ipv6 addresses that will be bound to this ipset. ## Attribute Reference In addition to the arguments, the following attributes are available:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of the ipset. It has the same value as the ` + "`" + `name` + "`" + ` attribute. ## Import A ipset can be imported using its name, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import citrixadc_ipset.tf_ipset tf_ipset ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id of the ipset. It has the same value as the ` + "`" + `name` + "`" + ` attribute. ## Import A ipset can be imported using its name, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import citrixadc_ipset.tf_ipset tf_ipset ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "citrixadc_lbmonitor",
			Category:         "Load Balancing",
			ShortDescription: ``,
			Description: `

The lbmonitor resource is used to create load balancing monitors.


`,
			Keywords: []string{
				"load",
				"balancing",
				"lbmonitor",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "monitorname",
					Description: `(Optional) Name for the monitor.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) Type of monitor that you want to create. Possible values: [ PING, TCP, HTTP, TCP-ECV, HTTP-ECV, UDP-ECV, DNS, FTP, LDNS-PING, LDNS-TCP, LDNS-DNS, RADIUS, USER, HTTP-INLINE, SIP-UDP, SIP-TCP, LOAD, FTP-EXTENDED, SMTP, SNMP, NNTP, MYSQL, MYSQL-ECV, MSSQL-ECV, ORACLE-ECV, LDAP, POP3, CITRIX-XML-SERVICE, CITRIX-WEB-INTERFACE, DNS-TCP, RTSP, ARP, CITRIX-AG, CITRIX-AAC-LOGINPAGE, CITRIX-AAC-LAS, CITRIX-XD-DDC, ND6, CITRIX-WI-EXTENDED, DIAMETER, RADIUS_ACCOUNTING, STOREFRONT, APPC, SMPP, CITRIX-XNC-ECV, CITRIX-XDM, CITRIX-STA-SERVICE, CITRIX-STA-SERVICE-NHOP ]`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Optional) Action to perform when the response to an inline monitor (a monitor of type HTTP-INLINE) indicates that the service is down. A service monitored by an inline monitor is considered DOWN if the response code is not one of the codes that have been specified for the Response Code parameter. Available settings function as follows: NONE - Do not take any action. However, the show service command and the show lb monitor command indicate the total number of responses that were checked and the number of consecutive error responses received after the last successful probe. LOG - Log the event in NSLOG or SYSLOG. DOWN - Mark the service as being down, and then do not direct any traffic to the service until the configured down time has expired. Persistent connections to the service are terminated as soon as the service is marked as DOWN. Also, log the event in NSLOG or SYSLOG. Possible values: [ NONE, LOG, DOWN ]`,
				},
				resource.Attribute{
					Name:        "respcode",
					Description: `(Optional) Response codes for which to mark the service as UP. For any other response code, the action performed depends on the monitor type. HTTP monitors and RADIUS monitors mark the service as DOWN, while HTTP-INLINE monitors perform the action indicated by the Action parameter.`,
				},
				resource.Attribute{
					Name:        "httprequest",
					Description: `(Optional) HTTP request to send to the server (for example, "HEAD /file.html").`,
				},
				resource.Attribute{
					Name:        "rtsprequest",
					Description: `(Optional) RTSP request to send to the server (for example, "OPTIONS \`,
				},
				resource.Attribute{
					Name:        "customheaders",
					Description: `(Optional) Custom header string to include in the monitoring probes.`,
				},
				resource.Attribute{
					Name:        "maxforwards",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "sipmethod",
					Description: `(Optional) SIP method to use for the query. Applicable only to monitors of type SIP-UDP. Possible values: [ OPTIONS, INVITE, REGISTER ]`,
				},
				resource.Attribute{
					Name:        "sipuri",
					Description: `(Optional) SIP URI string to send to the service (for example, sip:sip.test). Applicable only to monitors of type SIP-UDP.`,
				},
				resource.Attribute{
					Name:        "sipreguri",
					Description: `(Optional) SIP user to be registered. Applicable only if the monitor is of type SIP-UDP and the SIP Method parameter is set to REGISTER.`,
				},
				resource.Attribute{
					Name:        "send",
					Description: `(Optional) String to send to the service. Applicable to TCP-ECV, HTTP-ECV, and UDP-ECV monitors.`,
				},
				resource.Attribute{
					Name:        "recv",
					Description: `(Optional) String expected from the server for the service to be marked as UP. Applicable to TCP-ECV, HTTP-ECV, and UDP-ECV monitors.`,
				},
				resource.Attribute{
					Name:        "query",
					Description: `(Optional) Domain name to resolve as part of monitoring the DNS service (for example, example.com).`,
				},
				resource.Attribute{
					Name:        "querytype",
					Description: `(Optional) Type of DNS record for which to send monitoring queries. Set to Address for querying A records, AAAA for querying AAAA records, and Zone for querying the SOA record. Possible values: [ Address, Zone, AAAA ]`,
				},
				resource.Attribute{
					Name:        "scriptname",
					Description: `(Optional) Path and name of the script to execute. The script must be available on the Citrix ADC, in the /nsconfig/monitors/ directory.`,
				},
				resource.Attribute{
					Name:        "scriptargs",
					Description: `(Optional) String of arguments for the script. The string is copied verbatim into the request.`,
				},
				resource.Attribute{
					Name:        "dispatcherip",
					Description: `(Optional) IP address of the dispatcher to which to send the probe.`,
				},
				resource.Attribute{
					Name:        "dispatcherport",
					Description: `(Optional) Port number on which the dispatcher listens for the monitoring probe.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Optional) User name with which to probe the RADIUS, NNTP, FTP, FTP-EXTENDED, MYSQL, MSSQL, POP3, CITRIX-AG, CITRIX-XD-DDC, CITRIX-WI-EXTENDED, CITRIX-XNC or CITRIX-XDM server.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional) Password that is required for logging on to the RADIUS, NNTP, FTP, FTP-EXTENDED, MYSQL, MSSQL, POP3, CITRIX-AG, CITRIX-XD-DDC, CITRIX-WI-EXTENDED, CITRIX-XNC-ECV or CITRIX-XDM server. Used in conjunction with the user name specified for the User Name parameter.`,
				},
				resource.Attribute{
					Name:        "secondarypassword",
					Description: `(Optional) Secondary password that users might have to provide to log on to the Access Gateway server. Applicable to CITRIX-AG monitors.`,
				},
				resource.Attribute{
					Name:        "logonpointname",
					Description: `(Optional) Name of the logon point that is configured for the Citrix Access Gateway Advanced Access Control software. Required if you want to monitor the associated login page or Logon Agent. Applicable to CITRIX-AAC-LAS and CITRIX-AAC-LOGINPAGE monitors.`,
				},
				resource.Attribute{
					Name:        "lasversion",
					Description: `(Optional) Version number of the Citrix Advanced Access Control Logon Agent. Required by the CITRIX-AAC-LAS monitor.`,
				},
				resource.Attribute{
					Name:        "radkey",
					Description: `(Optional) Authentication key (shared secret text string) for RADIUS clients and servers to exchange. Applicable to monitors of type RADIUS and RADIUS_ACCOUNTING.`,
				},
				resource.Attribute{
					Name:        "radnasid",
					Description: `(Optional) NAS-Identifier to send in the Access-Request packet. Applicable to monitors of type RADIUS.`,
				},
				resource.Attribute{
					Name:        "radnasip",
					Description: `(Optional) Network Access Server (NAS) IP address to use as the source IP address when monitoring a RADIUS server. Applicable to monitors of type RADIUS and RADIUS_ACCOUNTING.`,
				},
				resource.Attribute{
					Name:        "radaccounttype",
					Description: `(Optional) Account Type to be used in Account Request Packet. Applicable to monitors of type RADIUS_ACCOUNTING.`,
				},
				resource.Attribute{
					Name:        "radframedip",
					Description: `(Optional) Source ip with which the packet will go out . Applicable to monitors of type RADIUS_ACCOUNTING.`,
				},
				resource.Attribute{
					Name:        "radapn",
					Description: `(Optional) Called Station Id to be used in Account Request Packet. Applicable to monitors of type RADIUS_ACCOUNTING.`,
				},
				resource.Attribute{
					Name:        "radmsisdn",
					Description: `(Optional) Calling Stations Id to be used in Account Request Packet. Applicable to monitors of type RADIUS_ACCOUNTING.`,
				},
				resource.Attribute{
					Name:        "radaccountsession",
					Description: `(Optional) Account Session ID to be used in Account Request Packet. Applicable to monitors of type RADIUS_ACCOUNTING.`,
				},
				resource.Attribute{
					Name:        "lrtm",
					Description: `(Optional) Calculate the least response times for bound services. If this parameter is not enabled, the appliance does not learn the response times of the bound services. Also used for LRTM load balancing. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "deviation",
					Description: `(Optional) Time value added to the learned average response time in dynamic response time monitoring (DRTM). When a deviation is specified, the appliance learns the average response time of bound services and adds the deviation to the average. The final value is then continually adjusted to accommodate response time variations over time. Specified in milliseconds, seconds, or minutes.`,
				},
				resource.Attribute{
					Name:        "units1",
					Description: `(Optional) Unit of measurement for the Deviation parameter. Cannot be changed after the monitor is created. Possible values: [ SEC, MSEC, MIN ]`,
				},
				resource.Attribute{
					Name:        "interval",
					Description: `(Optional) Time interval between two successive probes. Must be greater than the value of Response Time-out.`,
				},
				resource.Attribute{
					Name:        "units3",
					Description: `(Optional) monitor interval units. Possible values: [ SEC, MSEC, MIN ]`,
				},
				resource.Attribute{
					Name:        "resptimeout",
					Description: `(Optional) Amount of time for which the appliance must wait before it marks a probe as FAILED. Must be less than the value specified for the Interval parameter. Note: For UDP-ECV monitors for which a receive string is not configured, response timeout does not apply. For UDP-ECV monitors with no receive string, probe failure is indicated by an ICMP port unreachable error received from the service.`,
				},
				resource.Attribute{
					Name:        "units4",
					Description: `(Optional) monitor response timeout units. Possible values: [ SEC, MSEC, MIN ]`,
				},
				resource.Attribute{
					Name:        "resptimeoutthresh",
					Description: `(Optional) Response time threshold, specified as a percentage of the Response Time-out parameter. If the response to a monitor probe has not arrived when the threshold is reached, the appliance generates an SNMP trap called monRespTimeoutAboveThresh. After the response time returns to a value below the threshold, the appliance generates a monRespTimeoutBelowThresh SNMP trap. For the traps to be generated, the "MONITOR-RTO-THRESHOLD" alarm must also be enabled.`,
				},
				resource.Attribute{
					Name:        "retries",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "failureretries",
					Description: `(Optional) Number of retries that must fail, out of the number specified for the Retries parameter, for a service to be marked as DOWN. For example, if the Retries parameter is set to 10 and the Failure Retries parameter is set to 6, out of the ten probes sent, at least six probes must fail if the service is to be marked as DOWN. The default value of 0 means that all the retries must fail if the service is to be marked as DOWN.`,
				},
				resource.Attribute{
					Name:        "alertretries",
					Description: `(Optional) Number of consecutive probe failures after which the appliance generates an SNMP trap called monProbeFailed.`,
				},
				resource.Attribute{
					Name:        "successretries",
					Description: `(Optional) Number of consecutive successful probes required to transition a service's state from DOWN to UP.`,
				},
				resource.Attribute{
					Name:        "downtime",
					Description: `(Optional) Time duration for which to wait before probing a service that has been marked as DOWN. Expressed in milliseconds, seconds, or minutes.`,
				},
				resource.Attribute{
					Name:        "units2",
					Description: `(Optional) Unit of measurement for the Down Time parameter. Cannot be changed after the monitor is created. Possible values: [ SEC, MSEC, MIN ]`,
				},
				resource.Attribute{
					Name:        "destip",
					Description: `(Optional) IP address of the service to which to send probes. If the parameter is set to 0, the IP address of the server to which the monitor is bound is considered the destination IP address.`,
				},
				resource.Attribute{
					Name:        "destport",
					Description: `(Optional) TCP or UDP port to which to send the probe. If the parameter is set to 0, the port number of the service to which the monitor is bound is considered the destination port. For a monitor of type USER, however, the destination port is the port number that is included in the HTTP request sent to the dispatcher. Does not apply to monitors of type PING.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `(Optional) State of the monitor. The DISABLED setting disables not only the monitor being configured, but all monitors of the same type, until the parameter is set to ENABLED. If the monitor is bound to a service, the state of the monitor is not taken into account when the state of the service is determined. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "reverse",
					Description: `(Optional) Mark a service as DOWN, instead of UP, when probe criteria are satisfied, and as UP instead of DOWN when probe criteria are not satisfied. Possible values: [ YES, NO ]`,
				},
				resource.Attribute{
					Name:        "transparent",
					Description: `(Optional) The monitor is bound to a transparent device such as a firewall or router. The state of a transparent device depends on the responsiveness of the services behind it. If a transparent device is being monitored, a destination IP address must be specified. The probe is sent to the specified IP address by using the MAC address of the transparent device. Possible values: [ YES, NO ]`,
				},
				resource.Attribute{
					Name:        "iptunnel",
					Description: `(Optional) Send the monitoring probe to the service through an IP tunnel. A destination IP address must be specified. Possible values: [ YES, NO ]`,
				},
				resource.Attribute{
					Name:        "tos",
					Description: `(Optional) Probe the service by encoding the destination IP address in the IP TOS (6) bits. Possible values: [ YES, NO ]`,
				},
				resource.Attribute{
					Name:        "tosid",
					Description: `(Optional) The TOS ID of the specified destination IP. Applicable only when the TOS parameter is set.`,
				},
				resource.Attribute{
					Name:        "secure",
					Description: `(Optional) Use a secure SSL connection when monitoring a service. Applicable only to TCP based monitors. The secure option cannot be used with a CITRIX-AG monitor, because a CITRIX-AG monitor uses a secure connection by default. Possible values: [ YES, NO ]`,
				},
				resource.Attribute{
					Name:        "validatecred",
					Description: `(Optional) Validate the credentials of the Xen Desktop DDC server user. Applicable to monitors of type CITRIX-XD-DDC. Possible values: [ YES, NO ]`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `(Optional) Domain in which the XenDesktop Desktop Delivery Controller (DDC) servers or Web Interface servers are present. Required by CITRIX-XD-DDC and CITRIX-WI-EXTENDED monitors for logging on to the DDC servers and Web Interface servers, respectively.`,
				},
				resource.Attribute{
					Name:        "group",
					Description: `(Optional) Name of a newsgroup available on the NNTP service that is to be monitored. The appliance periodically generates an NNTP query for the name of the newsgroup and evaluates the response. If the newsgroup is found on the server, the service is marked as UP. If the newsgroup does not exist or if the search fails, the service is marked as DOWN. Applicable to NNTP monitors.`,
				},
				resource.Attribute{
					Name:        "filename",
					Description: `(Optional) Name of a file on the FTP server. The appliance monitors the FTP service by periodically checking the existence of the file on the server. Applicable to FTP-EXTENDED monitors.`,
				},
				resource.Attribute{
					Name:        "basedn",
					Description: `(Optional) The base distinguished name of the LDAP service, from where the LDAP server can begin the search for the attributes in the monitoring query. Required for LDAP service monitoring.`,
				},
				resource.Attribute{
					Name:        "binddn",
					Description: `(Optional) The distinguished name with which an LDAP monitor can perform the Bind operation on the LDAP server. Optional. Applicable to LDAP monitors.`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) Filter criteria for the LDAP query. Optional.`,
				},
				resource.Attribute{
					Name:        "attribute",
					Description: `(Optional) Attribute to evaluate when the LDAP server responds to the query. Success or failure of the monitoring probe depends on whether the attribute exists in the response. Optional.`,
				},
				resource.Attribute{
					Name:        "database",
					Description: `(Optional) Name of the database to connect to during authentication.`,
				},
				resource.Attribute{
					Name:        "oraclesid",
					Description: `(Optional) Name of the service identifier that is used to connect to the Oracle database during authentication.`,
				},
				resource.Attribute{
					Name:        "sqlquery",
					Description: `(Optional) SQL query for a MYSQL-ECV or MSSQL-ECV monitor. Sent to the database server after the server authenticates the connection.`,
				},
				resource.Attribute{
					Name:        "evalrule",
					Description: `(Optional) Expression that evaluates the database server's response to a MYSQL-ECV or MSSQL-ECV monitoring query. Must produce a Boolean result. The result determines the state of the server. If the expression returns TRUE, the probe succeeds. For example, if you want the appliance to evaluate the error message to determine the state of the server, use the rule MYSQL.RES.ROW(10) .TEXT_ELEM(2).EQ("MySQL").`,
				},
				resource.Attribute{
					Name:        "mssqlprotocolversion",
					Description: `(Optional) Version of MSSQL server that is to be monitored. Possible values: [ 70, 2000, 2000SP1, 2005, 2008, 2008R2, 2012, 2014 ]`,
				},
				resource.Attribute{
					Name:        "snmpoid",
					Description: `(Optional) SNMP OID for SNMP monitors.`,
				},
				resource.Attribute{
					Name:        "snmpcommunity",
					Description: `(Optional) Community name for SNMP monitors.`,
				},
				resource.Attribute{
					Name:        "snmpthreshold",
					Description: `(Optional) Threshold for SNMP monitors.`,
				},
				resource.Attribute{
					Name:        "snmpversion",
					Description: `(Optional) SNMP version to be used for SNMP monitors. Possible values: [ V1, V2 ]`,
				},
				resource.Attribute{
					Name:        "metrictable",
					Description: `(Optional) Metric table to which to bind metrics.`,
				},
				resource.Attribute{
					Name:        "application",
					Description: `(Optional) Name of the application used to determine the state of the service. Applicable to monitors of type CITRIX-XML-SERVICE.`,
				},
				resource.Attribute{
					Name:        "sitepath",
					Description: `(Optional) URL of the logon page. For monitors of type CITRIX-WEB-INTERFACE, to monitor a dynamic page under the site path, terminate the site path with a slash (/). Applicable to CITRIX-WEB-INTERFACE, CITRIX-WI-EXTENDED and CITRIX-XDM monitors.`,
				},
				resource.Attribute{
					Name:        "storename",
					Description: `(Optional) Store Name. For monitors of type STOREFRONT, STORENAME is an optional argument defining storefront service store name. Applicable to STOREFRONT monitors.`,
				},
				resource.Attribute{
					Name:        "storefrontacctservice",
					Description: `(Optional) Enable/Disable probing for Account Service. Applicable only to Store Front monitors. For multi-tenancy configuration users my skip account service. Possible values: [ YES, NO ]`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `(Optional) Hostname in the FQDN format (Example: porche.cars.org). Applicable to STOREFRONT monitors.`,
				},
				resource.Attribute{
					Name:        "netprofile",
					Description: `(Optional) Name of the network profile.`,
				},
				resource.Attribute{
					Name:        "originhost",
					Description: `(Optional) Origin-Host value for the Capabilities-Exchange-Request (CER) message to use for monitoring Diameter servers.`,
				},
				resource.Attribute{
					Name:        "originrealm",
					Description: `(Optional) Origin-Realm value for the Capabilities-Exchange-Request (CER) message to use for monitoring Diameter servers.`,
				},
				resource.Attribute{
					Name:        "hostipaddress",
					Description: `(Optional) Host-IP-Address value for the Capabilities-Exchange-Request (CER) message to use for monitoring Diameter servers. If Host-IP-Address is not specified, the appliance inserts the mapped IP (MIP) address or subnet IP (SNIP) address from which the CER request (the monitoring probe) is sent.`,
				},
				resource.Attribute{
					Name:        "vendorid",
					Description: `(Optional) Vendor-Id value for the Capabilities-Exchange-Request (CER) message to use for monitoring Diameter servers.`,
				},
				resource.Attribute{
					Name:        "productname",
					Description: `(Optional) Product-Name value for the Capabilities-Exchange-Request (CER) message to use for monitoring Diameter servers.`,
				},
				resource.Attribute{
					Name:        "firmwarerevision",
					Description: `(Optional) Firmware-Revision value for the Capabilities-Exchange-Request (CER) message to use for monitoring Diameter servers.`,
				},
				resource.Attribute{
					Name:        "inbandsecurityid",
					Description: `(Optional) Inband-Security-Id for the Capabilities-Exchange-Request (CER) message to use for monitoring Diameter servers. Possible values: [ NO_INBAND_SECURITY, TLS ]`,
				},
				resource.Attribute{
					Name:        "vendorspecificvendorid",
					Description: `(Optional) Vendor-Id to use in the Vendor-Specific-Application-Id grouped attribute-value pair (AVP) in the monitoring CER message. To specify Auth-Application-Id or Acct-Application-Id in Vendor-Specific-Application-Id, use vendorSpecificAuthApplicationIds or vendorSpecificAcctApplicationIds, respectively. Only one Vendor-Id is supported for all the Vendor-Specific-Application-Id AVPs in a CER monitoring message.`,
				},
				resource.Attribute{
					Name:        "kcdaccount",
					Description: `(Optional) KCD Account used by MSSQL monitor.`,
				},
				resource.Attribute{
					Name:        "storedb",
					Description: `(Optional) Store the database list populated with the responses to monitor probes. Used in database specific load balancing if MSSQL-ECV/MYSQL-ECV monitor is configured. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "storefrontcheckbackendservices",
					Description: `(Optional) This option will enable monitoring of services running on storefront server. Storefront services are monitored by probing to a Windows service that runs on the Storefront server and exposes details of which storefront services are running. Possible values: [ YES, NO ]`,
				},
				resource.Attribute{
					Name:        "trofscode",
					Description: `(Optional) Code expected when the server is under maintenance.`,
				},
				resource.Attribute{
					Name:        "trofsstring",
					Description: `(Optional) String expected from the server for the service to be marked as trofs. Applicable to HTTP-ECV/TCP-ECV monitors.`,
				},
				resource.Attribute{
					Name:        "sslprofile",
					Description: `(Optional) SSL Profile associated with the monitor.`,
				},
				resource.Attribute{
					Name:        "metric",
					Description: `(Optional) Metric name in the metric table, whose setting is changed. A value zero disables the metric and it will not be used for load calculation.`,
				},
				resource.Attribute{
					Name:        "metricthreshold",
					Description: `(Optional) Threshold to be used for that metric.`,
				},
				resource.Attribute{
					Name:        "metricweight",
					Description: `(Optional) The weight for the specified service metric with respect to others.`,
				},
				resource.Attribute{
					Name:        "servicename",
					Description: `(Optional) The name of the service to which the monitor is bound.`,
				},
				resource.Attribute{
					Name:        "servicegroupname",
					Description: `(Optional) The name of the service group to which the monitor is to be bound. ## Attribute Reference In addition to the arguments, the following attributes are available:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of the lbmonitor. It has the same value as the ` + "`" + `name` + "`" + ` attribute. ## Import A lbmonitor can be imported using its name, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import citrixadc_lbmonitor.tf_lbmonitor tf_lbmonitor ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id of the lbmonitor. It has the same value as the ` + "`" + `name` + "`" + ` attribute. ## Import A lbmonitor can be imported using its name, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import citrixadc_lbmonitor.tf_lbmonitor tf_lbmonitor ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "citrixadc_lbvserver",
			Category:         "Load Balancing",
			ShortDescription: ``,
			Description: `

The lbvserver resource is used to create load balancing virtual servers.


`,
			Keywords: []string{
				"load",
				"balancing",
				"lbvserver",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name for the virtual server.`,
				},
				resource.Attribute{
					Name:        "servicetype",
					Description: `(Optional) Protocol used by the service (also called the service type). Possible values: [ HTTP, FTP, TCP, UDP, SSL, SSL_BRIDGE, SSL_TCP, DTLS, NNTP, DNS, DHCPRA, ANY, SIP_UDP, SIP_TCP, SIP_SSL, DNS_TCP, RTSP, PUSH, SSL_PUSH, RADIUS, RDP, MYSQL, MSSQL, DIAMETER, SSL_DIAMETER, TFTP, ORACLE, SMPP, SYSLOGTCP, SYSLOGUDP, FIX, SSL_FIX, PROXY, USER_TCP, USER_SSL_TCP, QUIC, IPFIX, LOGSTREAM, MONGO, MONGO_TLS ]`,
				},
				resource.Attribute{
					Name:        "ipv46",
					Description: `(Optional) IPv4 or IPv6 address to assign to the virtual server.`,
				},
				resource.Attribute{
					Name:        "ippattern",
					Description: `(Optional) IP address pattern, in dotted decimal notation, for identifying packets to be accepted by the virtual server. The IP Mask parameter specifies which part of the destination IP address is matched against the pattern. Mutually exclusive with the IP Address parameter. For example, if the IP pattern assigned to the virtual server is 198.51.100.0 and the IP mask is 255.255.240.0 (a forward mask), the first 20 bits in the destination IP addresses are matched with the first 20 bits in the pattern. The virtual server accepts requests with IP addresses that range from 198.51.96.1 to 198.51.111.254. You can also use a pattern such as 0.0.2.2 and a mask such as 0.0.255.255 (a reverse mask). If a destination IP address matches more than one IP pattern, the pattern with the longest match is selected, and the associated virtual server processes the request. For example, if virtual servers vs1 and vs2 have the same IP pattern, 0.0.100.128, but different IP masks of 0.0.255.255 and 0.0.224.255, a destination IP address of 198.51.100.128 has the longest match with the IP pattern of vs1. If a destination IP address matches two or more virtual servers to the same extent, the request is processed by the virtual server whose port number matches the port number in the request.`,
				},
				resource.Attribute{
					Name:        "ipmask",
					Description: `(Optional) IP mask, in dotted decimal notation, for the IP Pattern parameter. Can have leading or trailing non-zero octets (for example, 255.255.240.0 or 0.0.255.255). Accordingly, the mask specifies whether the first n bits or the last n bits of the destination IP address in a client request are to be matched with the corresponding bits in the IP pattern. The former is called a forward mask. The latter is called a reverse mask.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional) Port number for the virtual server. Range 1 - 65535`,
				},
				resource.Attribute{
					Name:        "ipset",
					Description: `(Optional) The list of IPv4/IPv6 addresses bound to ipset would form a part of listening service on the current lb vserver.`,
				},
				resource.Attribute{
					Name:        "range",
					Description: `(Optional) Number of IP addresses that the appliance must generate and assign to the virtual server. The virtual server then functions as a network virtual server, accepting traffic on any of the generated IP addresses. The IP addresses are generated automatically, as follows:`,
				},
				resource.Attribute{
					Name:        "persistencetype",
					Description: `(Optional) Type of persistence for the virtual server. Available settings function as follows: SOURCEIP - Connections from the same client IP address belong to the same persistence session. COOKIEINSERT - Connections that have the same HTTP Cookie, inserted by a Set-Cookie directive from a server, belong to the same persistence session. SSLSESSION - Connections that have the same SSL Session ID belong to the same persistence session. CUSTOMSERVERID - Connections with the same server ID form part of the same session. For this persistence type, set the Server ID (CustomServerID) parameter for each service and configure the Rule parameter to identify the server ID in a request. RULE - All connections that match a user defined rule belong to the same persistence session. URLPASSIVE - Requests that have the same server ID in the URL query belong to the same persistence session. The server ID is the hexadecimal representation of the IP address and port of the service to which the request must be forwarded. This persistence type requires a rule to identify the server ID in the request. DESTIP - Connections to the same destination IP address belong to the same persistence session. SRCIPDESTIP - Connections that have the same source IP address and destination IP address belong to the same persistence session. CALLID - Connections that have the same CALL-ID SIP header belong to the same persistence session. RTSPSID - Connections that have the same RTSP Session ID belong to the same persistence session. FIXSESSION - Connections that have the same SenderCompID and TargetCompID values belong to the same persistence session. USERSESSION - Persistence session is created based on the persistence parameter value provided from an extension. Possible values: [ SOURCEIP, COOKIEINSERT, SSLSESSION, RULE, URLPASSIVE, CUSTOMSERVERID, DESTIP, SRCIPDESTIP, CALLID, RTSPSID, DIAMETER, FIXSESSION, USERSESSION, NONE ]`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `(Optional) Time period for which a persistence session is in effect.`,
				},
				resource.Attribute{
					Name:        "persistencebackup",
					Description: `(Optional) Backup persistence type for the virtual server. Becomes operational if the primary persistence mechanism fails. Possible values: [ SOURCEIP, NONE ]`,
				},
				resource.Attribute{
					Name:        "backuppersistencetimeout",
					Description: `(Optional) Time period for which backup persistence is in effect.`,
				},
				resource.Attribute{
					Name:        "lbmethod",
					Description: `(Optional) Load balancing method. The available settings function as follows:`,
				},
				resource.Attribute{
					Name:        "hashlength",
					Description: `(Optional) Number of bytes to consider for the hash value used in the URLHASH and DOMAINHASH load balancing methods.`,
				},
				resource.Attribute{
					Name:        "netmask",
					Description: `(Optional) IPv4 subnet mask to apply to the destination IP address or source IP address when the load balancing method is DESTINATIONIPHASH or SOURCEIPHASH.`,
				},
				resource.Attribute{
					Name:        "v6netmasklen",
					Description: `(Optional) Number of bits to consider in an IPv6 destination or source IP address, for creating the hash that is required by the DESTINATIONIPHASH and SOURCEIPHASH load balancing methods.`,
				},
				resource.Attribute{
					Name:        "backuplbmethod",
					Description: `(Optional) Backup load balancing method. Becomes operational if the primary load balancing me thod fails or cannot be used. Valid only if the primary method is based on static proximity. Possible values: [ ROUNDROBIN, LEASTCONNECTION, LEASTRESPONSETIME, SOURCEIPHASH, LEASTBANDWIDTH, LEASTPACKETS, CUSTOMLOAD ]`,
				},
				resource.Attribute{
					Name:        "cookiename",
					Description: `(Optional) Use this parameter to specify the cookie name for COOKIE peristence type. It specifies the name of cookie with a maximum of 32 characters. If not specified, cookie name is internally generated.`,
				},
				resource.Attribute{
					Name:        "rule",
					Description: `(Optional) Expression, or name of a named expression, against which traffic is evaluated. The following requirements apply only to the Citrix ADC CLI:`,
				},
				resource.Attribute{
					Name:        "listenpolicy",
					Description: `(Optional) Expression identifying traffic accepted by the virtual server. Can be either an expression (for example, CLIENT.IP.DST.IN_SUBNET(192.0.2.0/24) or the name of a named expression. In the above example, the virtual server accepts all requests whose destination IP address is in the 192.0.2.0/24 subnet.`,
				},
				resource.Attribute{
					Name:        "listenpriority",
					Description: `(Optional) Integer specifying the priority of the listen policy. A higher number specifies a lower priority. If a request matches the listen policies of more than one virtual server the virtual server whose listen policy has the highest priority (the lowest priority number) accepts the request.`,
				},
				resource.Attribute{
					Name:        "resrule",
					Description: `(Optional) Expression specifying which part of a server's response to use for creating rule based persistence sessions (persistence type RULE). Can be either an expression or the name of a named expression. Example: HTTP.RES.HEADER("setcookie").VALUE(0).TYPECAST_NVLIST_T('=',';').VALUE("server1").`,
				},
				resource.Attribute{
					Name:        "persistmask",
					Description: `(Optional) Persistence mask for IP based persistence types, for IPv4 virtual servers.`,
				},
				resource.Attribute{
					Name:        "v6persistmasklen",
					Description: `(Optional) Persistence mask for IP based persistence types, for IPv6 virtual servers.`,
				},
				resource.Attribute{
					Name:        "pq",
					Description: `(Optional) Use priority queuing on the virtual server. based persistence types, for IPv6 virtual servers. Possible values: [ on, off ]`,
				},
				resource.Attribute{
					Name:        "sc",
					Description: `(Optional) Use SureConnect on the virtual server. Possible values: [ on, off ]`,
				},
				resource.Attribute{
					Name:        "rtspnat",
					Description: `(Optional) Use network address translation (NAT) for RTSP data connections. Possible values: [ on, off ]`,
				},
				resource.Attribute{
					Name:        "m",
					Description: `(Optional) Redirection mode for load balancing. Available settings function as follows:`,
				},
				resource.Attribute{
					Name:        "tosid",
					Description: `(Optional) TOS ID of the virtual server. Applicable only when the load balancing redirection mode is set to TOS.`,
				},
				resource.Attribute{
					Name:        "datalength",
					Description: `(Optional) Length of the token to be extracted from the data segment of an incoming packet, for use in the token method of load balancing. The length of the token, specified in bytes, must not be greater than 24 KB. Applicable to virtual servers of type TCP.`,
				},
				resource.Attribute{
					Name:        "dataoffset",
					Description: `(Optional) Offset to be considered when extracting a token from the TCP payload. Applicable to virtual servers, of type TCP, using the token method of load balancing. Must be within the first 24 KB of the TCP payload.`,
				},
				resource.Attribute{
					Name:        "sessionless",
					Description: `(Optional) Perform load balancing on a per-packet basis, without establishing sessions. Recommended for load balancing of intrusion detection system (IDS) servers and scenarios involving direct server return (DSR), where session information is unnecessary. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "trofspersistence",
					Description: `(Optional) When value is ENABLED, Trofs persistence is honored. When value is DISABLED, Trofs persistence is not honored. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `(Optional) State of the load balancing virtual server. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "connfailover",
					Description: `(Optional) Mode in which the connection failover feature must operate for the virtual server. After a failover, established TCP connections and UDP packet flows are kept active and resumed on the secondary appliance. Clients remain connected to the same servers. Available settings function as follows:`,
				},
				resource.Attribute{
					Name:        "redirurl",
					Description: `(Optional) URL to which to redirect traffic if the virtual server becomes unavailable. WARNING! Make sure that the domain in the URL does not match the domain specified for a content switching policy. If it does, requests are continuously redirected to the unavailable virtual server.`,
				},
				resource.Attribute{
					Name:        "cacheable",
					Description: `(Optional) Route cacheable requests to a cache redirection virtual server. The load balancing virtual server can forward requests only to a transparent cache redirection virtual server that has an IP address and port combination of \`,
				},
				resource.Attribute{
					Name:        "clttimeout",
					Description: `(Optional) Idle time, in seconds, after which a client connection is terminated.`,
				},
				resource.Attribute{
					Name:        "somethod",
					Description: `(Optional) Type of threshold that, when exceeded, triggers spillover. Available settings function as follows:`,
				},
				resource.Attribute{
					Name:        "sopersistence",
					Description: `(Optional) If spillover occurs, maintain source IP address based persistence for both primary and backup virtual servers. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "sopersistencetimeout",
					Description: `(Optional) Timeout for spillover persistence, in minutes.`,
				},
				resource.Attribute{
					Name:        "healththreshold",
					Description: `(Optional) Threshold in percent of active services below which vserver state is made down. If this threshold is 0, vserver state will be up even if one bound service is up.`,
				},
				resource.Attribute{
					Name:        "sothreshold",
					Description: `(Optional) Threshold at which spillover occurs. Specify an integer for the CONNECTION spillover method, a bandwidth value in kilobits per second for the BANDWIDTH method (do not enter the units), or a percentage for the HEALTH method (do not enter the percentage symbol).`,
				},
				resource.Attribute{
					Name:        "sobackupaction",
					Description: `(Optional) Action to be performed if spillover is to take effect, but no backup chain to spillover is usable or exists. Possible values: [ DROP, ACCEPT, REDIRECT ]`,
				},
				resource.Attribute{
					Name:        "redirectportrewrite",
					Description: `(Optional) Rewrite the port and change the protocol to ensure successful HTTP redirects from services. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "downstateflush",
					Description: `(Optional) Flush all active transactions associated with a virtual server whose state transitions from UP to DOWN. Do not enable this option for applications that must complete their transactions. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "backupvserver",
					Description: `(Optional) Name of the backup virtual server to which to forward requests if the primary virtual server goes DOWN or reaches its spillover threshold.`,
				},
				resource.Attribute{
					Name:        "disableprimaryondown",
					Description: `(Optional) If the primary virtual server goes down, do not allow it to return to primary status until manually enabled. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "insertvserveripport",
					Description: `(Optional) Insert an HTTP header, whose value is the IP address and port number of the virtual server, before forwarding a request to the server. The format of the header is <vipHeader>: <virtual server IP address>\_<port number >, where vipHeader is the name that you specify for the header. If the virtual server has an IPv6 address, the address in the header is enclosed in brackets ([ and ]) to separate it from the port number. If you have mapped an IPv4 address to a virtual server's IPv6 address, the value of this parameter determines which IP address is inserted in the header, as follows:`,
				},
				resource.Attribute{
					Name:        "vipheader",
					Description: `(Optional) Name for the inserted header. The default name is vip-header.`,
				},
				resource.Attribute{
					Name:        "authenticationhost",
					Description: `(Optional) Fully qualified domain name (FQDN) of the authentication virtual server to which the user must be redirected for authentication. Make sure that the Authentication parameter is set to ENABLED.`,
				},
				resource.Attribute{
					Name:        "authentication",
					Description: `(Optional) Enable or disable user authentication. Possible values: [ on, off ]`,
				},
				resource.Attribute{
					Name:        "authn401",
					Description: `(Optional) Enable or disable user authentication with HTTP 401 responses. Possible values: [ on, off ]`,
				},
				resource.Attribute{
					Name:        "authnvsname",
					Description: `(Optional) Name of an authentication virtual server with which to authenticate users.`,
				},
				resource.Attribute{
					Name:        "push",
					Description: `(Optional) Process traffic with the push virtual server that is bound to this load balancing virtual server. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "pushvserver",
					Description: `(Optional) Name of the load balancing virtual server, of type PUSH or SSL_PUSH, to which the server pushes updates received on the load balancing virtual server that you are configuring.`,
				},
				resource.Attribute{
					Name:        "pushlabel",
					Description: `(Optional) Expression for extracting a label from the server's response. Can be either an expression or the name of a named expression.`,
				},
				resource.Attribute{
					Name:        "pushmulticlients",
					Description: `(Optional) Allow multiple Web 2.0 connections from the same client to connect to the virtual server and expect updates. Possible values: [ YES, NO ]`,
				},
				resource.Attribute{
					Name:        "tcpprofilename",
					Description: `(Optional) Name of the TCP profile whose settings are to be applied to the virtual server.`,
				},
				resource.Attribute{
					Name:        "httpprofilename",
					Description: `(Optional) Name of the HTTP profile whose settings are to be applied to the virtual server.`,
				},
				resource.Attribute{
					Name:        "dbprofilename",
					Description: `(Optional) Name of the DB profile whose settings are to be applied to the virtual server.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `(Optional) Any comments that you might want to associate with the virtual server.`,
				},
				resource.Attribute{
					Name:        "l2conn",
					Description: `(Optional) Use Layer 2 parameters (channel number, MAC address, and VLAN ID) in addition to the 4-tuple (<source IP>:<source port>::<destination IP>:<destination port>) that is used to identify a connection. Allows multiple TCP and non-TCP connections with the same 4-tuple to co-exist on the Citrix ADC. Possible values: [ on, off ]`,
				},
				resource.Attribute{
					Name:        "oracleserverversion",
					Description: `(Optional) Oracle server version. Possible values: [ 10G, 11G ]`,
				},
				resource.Attribute{
					Name:        "mssqlserverversion",
					Description: `(Optional) For a load balancing virtual server of type MSSQL, the Microsoft SQL Server version. Set this parameter if you expect some clients to run a version different from the version of the database. This setting provides compatibility between the client-side and server-side connections by ensuring that all communication conforms to the server's version. Possible values: [ 70, 2000, 2000SP1, 2005, 2008, 2008R2, 2012, 2014 ]`,
				},
				resource.Attribute{
					Name:        "mysqlprotocolversion",
					Description: `(Optional) MySQL protocol version that the virtual server advertises to clients.`,
				},
				resource.Attribute{
					Name:        "mysqlserverversion",
					Description: `(Optional) MySQL server version string that the virtual server advertises to clients.`,
				},
				resource.Attribute{
					Name:        "mysqlcharacterset",
					Description: `(Optional) Character set that the virtual server advertises to clients.`,
				},
				resource.Attribute{
					Name:        "mysqlservercapabilities",
					Description: `(Optional) Server capabilities that the virtual server advertises to clients.`,
				},
				resource.Attribute{
					Name:        "appflowlog",
					Description: `(Optional) Apply AppFlow logging to the virtual server. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "netprofile",
					Description: `(Optional) Name of the network profile to associate with the virtual server. If you set this parameter, the virtual server uses only the IP addresses in the network profile as source IP addresses when initiating connections with servers.`,
				},
				resource.Attribute{
					Name:        "icmpvsrresponse",
					Description: `(Optional) How the Citrix ADC responds to ping requests received for an IP address that is common to one or more virtual servers. Available settings function as follows:`,
				},
				resource.Attribute{
					Name:        "rhistate",
					Description: `(Optional) Route Health Injection (RHI) functionality of the NetSaler appliance for advertising the route of the VIP address associated with the virtual server. When Vserver RHI Level (RHI) parameter is set to VSVR\_CNTRLD, the following are different RHI behaviors for the VIP address on the basis of RHIstate (RHI STATE) settings on the virtual servers associated with the VIP address:`,
				},
				resource.Attribute{
					Name:        "newservicerequest",
					Description: `(Optional) Number of requests, or percentage of the load on existing services, by which to increase the load on a new service at each interval in slow-start mode. A non-zero value indicates that slow-start is applicable. A zero value indicates that the global RR startup parameter is applied. Changing the value to zero will cause services currently in slow start to take the full traffic as determined by the LB method. Subsequently, any new services added will use the global RR factor.`,
				},
				resource.Attribute{
					Name:        "newservicerequestunit",
					Description: `(Optional) Units in which to increment load at each interval in slow-start mode. Possible values: [ PER_SECOND, PERCENT ]`,
				},
				resource.Attribute{
					Name:        "newservicerequestincrementinterval",
					Description: `(Optional) Interval, in seconds, between successive increments in the load on a new service or a service whose state has just changed from DOWN to UP. A value of 0 (zero) specifies manual slow start.`,
				},
				resource.Attribute{
					Name:        "minautoscalemembers",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "maxautoscalemembers",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "skippersistency",
					Description: `(Optional) This argument decides the behavior incase the service which is selected from an existing persistence session has reached threshold. Possible values: [ Bypass, ReLb, None ]`,
				},
				resource.Attribute{
					Name:        "td",
					Description: `(Optional) Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.`,
				},
				resource.Attribute{
					Name:        "authnprofile",
					Description: `(Optional) Name of the authentication profile to be used when authentication is turned on.`,
				},
				resource.Attribute{
					Name:        "macmoderetainvlan",
					Description: `(Optional) This option is used to retain vlan information of incoming packet when macmode is enabled. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "dbslb",
					Description: `(Optional) Enable database specific load balancing for MySQL and MSSQL service types. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "dns64",
					Description: `(Optional) This argument is for enabling/disabling the dns64 on lbvserver. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "bypassaaaa",
					Description: `(Optional) If this option is enabled while resolving DNS64 query AAAA queries are not sent to back end dns server. Possible values: [ YES, NO ]`,
				},
				resource.Attribute{
					Name:        "recursionavailable",
					Description: `(Optional) When set to YES, this option causes the DNS replies from this vserver to have the RA bit turned on. Typically one would set this option to YES, when the vserver is load balancing a set of DNS servers thatsupport recursive queries. Possible values: [ YES, NO ]`,
				},
				resource.Attribute{
					Name:        "processlocal",
					Description: `(Optional) By turning on this option packets destined to a vserver in a cluster will not under go any steering. Turn this option for single packet request response mode or when the upstream device is performing a proper RSS for connection based distribution. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "dnsprofilename",
					Description: `(Optional) Name of the DNS profile to be associated with the VServer. DNS profile properties will be applied to the transactions processed by a VServer. This parameter is valid only for DNS and DNS-TCP VServers.`,
				},
				resource.Attribute{
					Name:        "lbprofilename",
					Description: `(Optional) Name of the LB profile which is associated to the vserver.`,
				},
				resource.Attribute{
					Name:        "redirectfromport",
					Description: `(Optional) Port number for the virtual server, from which we absorb the traffic for http redirect. Range 1 - 65535`,
				},
				resource.Attribute{
					Name:        "httpsredirecturl",
					Description: `(Optional) URL to which to redirect traffic if the traffic is recieved from redirect port.`,
				},
				resource.Attribute{
					Name:        "retainconnectionsoncluster",
					Description: `(Optional) This option enables you to retain existing connections on a node joining a Cluster system or when a node is being configured for passive timeout. By default, this option is disabled. Possible values: [ YES, NO ]`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `(Optional) Weight to assign to the specified service.`,
				},
				resource.Attribute{
					Name:        "servicename",
					Description: `(Optional) Service to bind to the virtual server.`,
				},
				resource.Attribute{
					Name:        "redirurlflags",
					Description: `(Optional) The redirect URL to be unset.`,
				},
				resource.Attribute{
					Name:        "newname",
					Description: `(Optional) New name for the virtual server.`,
				},
				resource.Attribute{
					Name:        "sslcertkey",
					Description: `(Optoinal) SSL certificate key to be bound to the lb vserver.`,
				},
				resource.Attribute{
					Name:        "ciphersuites",
					Description: `(Optional) A set of ciphersuites to be bound to the lb vserver.`,
				},
				resource.Attribute{
					Name:        "ciphers",
					Description: `(Optional) A set of ciphers to be bound to the lb vserver. This is only relevant for cluster deployments of ADC. In all other cases it is ignored.`,
				},
				resource.Attribute{
					Name:        "snisslcertkeys",
					Description: `(Optional) A set of SNI ssl certificates bound to the lb vserver.`,
				},
				resource.Attribute{
					Name:        "sslprofile",
					Description: `(Optional) A SSL profile to bind to this lb vserver.`,
				},
				resource.Attribute{
					Name:        "sslpolicybinding",
					Description: `(Optional) SSL policy bound to the lb vserver block. Documented below SSL policy binding supports the following:`,
				},
				resource.Attribute{
					Name:        "policyname",
					Description: `(Optional) The name of the SSL policy binding.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `(Optional) The priority of the policies bound to this SSL service.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) Bind point to which to bind the policy. Possible Values: REQUEST, INTERCEPT_REQ and CLIENTHELLO_REQ. These bindpoints mean: 1. REQUEST: Policy evaluation will be done at appplication above SSL. This bindpoint is default and is used for actions based on clientauth and client cert. 2. INTERCEPT_REQ: Policy evaluation will be done during SSL handshake to decide whether to intercept or not. Actions allowed with this type are: INTERCEPT, BYPASS and RESET. 3. CLIENTHELLO_REQ: Policy evaluation will be done during handling of Client Hello Request. Action allowed with this type is: RESET, FORWARD and PICKCACERTGRP. Possible values: [ INTERCEPT_REQ, REQUEST, CLIENTHELLO_REQ ]`,
				},
				resource.Attribute{
					Name:        "gotopriorityexpression",
					Description: `(Optional) Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.`,
				},
				resource.Attribute{
					Name:        "invoke",
					Description: `(Optional) Invoke flag. This attribute is relevant only for ADVANCED policies.`,
				},
				resource.Attribute{
					Name:        "labeltype",
					Description: `(Optional) Type of policy label invocation. Possible values: [ vserver, service, policylabel ]`,
				},
				resource.Attribute{
					Name:        "labelname",
					Description: `(Optional) Name of the label to invoke if the current policy rule evaluates to TRUE.`,
				},
				resource.Attribute{
					Name:        "quicbridgeprofilename",
					Description: `(Optional) Name of the QUIC Bridge profile whose settings are to be applied to the virtual server. Refer [quicbridgeprofile](./quicbridgeprofile.md) for usage. ## Attribute Reference In addition to the arguments, the following attributes are available:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of the lbvserver. It has the same value as the ` + "`" + `name` + "`" + ` attribute. ## Import A lbvserver can be imported using its name, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import citrixadc_lbvserver.tf_lbvserver tf_lbvserver ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id of the lbvserver. It has the same value as the ` + "`" + `name` + "`" + ` attribute. ## Import A lbvserver can be imported using its name, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import citrixadc_lbvserver.tf_lbvserver tf_lbvserver ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "citrixadc_lbvserver_appfwpolicy_binding",
			Category:         "Load Balancing",
			ShortDescription: ``,
			Description: `

The lbvserver_appfwpolicy_binding resource is used to add AppFw policies to lbvserver.

`,
			Keywords: []string{
				"load",
				"balancing",
				"lbvserver",
				"appfwpolicy",
				"binding",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `Name for the virtual server. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at sign (@), equal sign (=), and hyphen (-) characters. Can be changed after the virtual server is created. CLI Users: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my vserver" or 'my vserver'). .`,
				},
				resource.Attribute{
					Name:        "policyname",
					Description: `Name of the policy bound to the LB vserver.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `(Optional) Priority.`,
				},
				resource.Attribute{
					Name:        "gotopriorityexpression",
					Description: `(Optional) Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.`,
				},
				resource.Attribute{
					Name:        "bindpoint",
					Description: `(Optional) The bindpoint to which the policy is bound. Possible values: [ REQUEST, RESPONSE ]`,
				},
				resource.Attribute{
					Name:        "invoke",
					Description: `(Optional) Invoke policies bound to a virtual server or policy label.`,
				},
				resource.Attribute{
					Name:        "labeltype",
					Description: `(Optional) The invocation type. Possible values: [ reqvserver, resvserver, policylabel ]`,
				},
				resource.Attribute{
					Name:        "labelname",
					Description: `(Optional) Name of the label invoked. ## Attribute Reference In addition to the arguments, the following attributes are available:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of the lbvserver_appfwpolicy_binding. It has the same value as the ` + "`" + `name` + "`" + ` attribute.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id of the lbvserver_appfwpolicy_binding. It has the same value as the ` + "`" + `name` + "`" + ` attribute.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "citrixadc_lbvserver_cmppolicy_binding",
			Category:         "Load Balancing",
			ShortDescription: ``,
			Description: `\_cmppolicy\_binding

The lbvserver\_cmppolicy\_binding resource is used to bind load balancing virtual servers with compression policies.


`,
			Keywords: []string{
				"load",
				"balancing",
				"lbvserver",
				"cmppolicy",
				"binding",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "policyname",
					Description: `(Required) Name of the policy bound to the LB vserver.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `(Optional) Priority.`,
				},
				resource.Attribute{
					Name:        "gotopriorityexpression",
					Description: `(Optional) Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.`,
				},
				resource.Attribute{
					Name:        "bindpoint",
					Description: `(Optional) The bindpoint to which the policy is bound. Possible values: [ REQUEST, RESPONSE ]`,
				},
				resource.Attribute{
					Name:        "invoke",
					Description: `(Optional) Invoke policies bound to a virtual server or policy label.`,
				},
				resource.Attribute{
					Name:        "labeltype",
					Description: `(Optional) The invocation type. Possible values: [ reqvserver, resvserver, policylabel ]`,
				},
				resource.Attribute{
					Name:        "labelname",
					Description: `(Optional) Name of the label invoked.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name for the virtual server. Must begin with an ASCII alphanumeric or underscore (\_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at sign (@), equal sign (=), and hyphen (-) characters. Can be changed after the virtual server is created. CLI Users: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my vserver" or 'my vserver'). . ## Attribute Reference In addition to the arguments, the following attributes are available:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of the lbvserver\_cmppolicy\_binding. It is the concatenation of the ` + "`" + `name` + "`" + ` and ` + "`" + `policyname` + "`" + ` attributes separated by a comma. ## Import A lbvserver\_cmppolicy\_binding can be imported using its id, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import citrixadc_lbvserver_cmppolicy_binding.tf_bind tf_lbvserver,tf_cmppolicy ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id of the lbvserver\_cmppolicy\_binding. It is the concatenation of the ` + "`" + `name` + "`" + ` and ` + "`" + `policyname` + "`" + ` attributes separated by a comma. ## Import A lbvserver\_cmppolicy\_binding can be imported using its id, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import citrixadc_lbvserver_cmppolicy_binding.tf_bind tf_lbvserver,tf_cmppolicy ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "citrixadc_lbvserver_filterpolicy_binding",
			Category:         "Load Balancing",
			ShortDescription: ``,
			Description: `\_filterpolicy\_binding

The lbvserver\_filterpolicy\_binding resource is used to bind load balancing virtual servers to filter policies.


`,
			Keywords: []string{
				"load",
				"balancing",
				"lbvserver",
				"filterpolicy",
				"binding",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "policyname",
					Description: `(Required) Name of the policy bound to the LB vserver.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `(Optional) Priority.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name for the virtual server. Must begin with an ASCII alphanumeric or underscore (\_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at sign (@), equal sign (=), and hyphen (-) characters. Can be changed after the virtual server is created. CLI Users: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my vserver" or 'my vserver'). .`,
				},
				resource.Attribute{
					Name:        "gotopriorityexpression",
					Description: `(Optional) Expression or other value specifying the next policy to be evaluated if the current policy evaluates to TRUE. Specify one of the following values:`,
				},
				resource.Attribute{
					Name:        "bindpoint",
					Description: `(Optional) Bind point to which to bind the policy. Applicable only to compression, rewrite, videooptimization and cache policies. Possible values: [ REQUEST, RESPONSE ]`,
				},
				resource.Attribute{
					Name:        "invoke",
					Description: `(Optional) Invoke policies bound to a virtual server or policy label.`,
				},
				resource.Attribute{
					Name:        "labeltype",
					Description: `(Optional) Type of policy label to invoke. Applicable only to rewrite, videooptimization and cache policies. Available settings function as follows:`,
				},
				resource.Attribute{
					Name:        "labelname",
					Description: `(Optional) Name of the virtual server or user-defined policy label to invoke if the policy evaluates to TRUE. ## Attribute Reference In addition to the arguments, the following attributes are available:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of the lbvserver\_filterpolicy\_binding. It is the concatenation of the ` + "`" + `name` + "`" + ` and ` + "`" + `policyname` + "`" + ` attributes separated by a comma. ## Import A lbvserver\_filterpolicy\_binding can be imported using its id, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import citrixadc_lbvserver_filterpolicy_binding.tf_bind tf_lbvserver,tf_filterpolicy ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id of the lbvserver\_filterpolicy\_binding. It is the concatenation of the ` + "`" + `name` + "`" + ` and ` + "`" + `policyname` + "`" + ` attributes separated by a comma. ## Import A lbvserver\_filterpolicy\_binding can be imported using its id, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import citrixadc_lbvserver_filterpolicy_binding.tf_bind tf_lbvserver,tf_filterpolicy ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "citrixadc_lbvserver_responderpolicy_binding",
			Category:         "Load Balancing",
			ShortDescription: ``,
			Description: `\_responderpolicy\_binding

The lbvserver\_responderpolicy\_binding resource is used to bind load balancing virtual servers to responder policies.


`,
			Keywords: []string{
				"load",
				"balancing",
				"lbvserver",
				"responderpolicy",
				"binding",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "policyname",
					Description: `(Required) Name of the policy bound to the LB vserver.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `(Optional) Priority.`,
				},
				resource.Attribute{
					Name:        "gotopriorityexpression",
					Description: `(Optional) Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.`,
				},
				resource.Attribute{
					Name:        "invoke",
					Description: `(Optional) Invoke policies bound to a virtual server or policy label.`,
				},
				resource.Attribute{
					Name:        "labeltype",
					Description: `(Optional) The invocation type. Possible values: [ reqvserver, resvserver, policylabel ]`,
				},
				resource.Attribute{
					Name:        "labelname",
					Description: `(Optional) Name of the label invoked.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name for the virtual server. Must begin with an ASCII alphanumeric or underscore (\_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at sign (@), equal sign (=), and hyphen (-) characters. Can be changed after the virtual server is created. CLI Users: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my vserver" or 'my vserver'). .`,
				},
				resource.Attribute{
					Name:        "bindpoint",
					Description: `(Optional) Bind point to which to bind the policy. Applicable only to compression, rewrite, videooptimization and cache policies. Possible values: [ REQUEST, RESPONSE ] ## Attribute Reference In addition to the arguments, the following attributes are available:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of the lbvserver\_responderpolicy\_binding. It is the concatenation of the ` + "`" + `name` + "`" + ` and ` + "`" + `policyname` + "`" + ` attributes separated by a comma. ## Import A lbvserver\_responderpolicy\_binding can be imported using its id, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import citrixadc_lbvserver_responderpolicy_binding.tf_bind tf_lbvserver,tf_responder_policy ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id of the lbvserver\_responderpolicy\_binding. It is the concatenation of the ` + "`" + `name` + "`" + ` and ` + "`" + `policyname` + "`" + ` attributes separated by a comma. ## Import A lbvserver\_responderpolicy\_binding can be imported using its id, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import citrixadc_lbvserver_responderpolicy_binding.tf_bind tf_lbvserver,tf_responder_policy ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "citrixadc_lbvserver_rewritepolicy_binding",
			Category:         "Load Balancing",
			ShortDescription: ``,
			Description: `\_rewritepolicy\_binding

The lbvserver\_rewritepolicy\_binding resource is used to bind load balancing virtual servers to rewrite policies.


`,
			Keywords: []string{
				"load",
				"balancing",
				"lbvserver",
				"rewritepolicy",
				"binding",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "policyname",
					Description: `(Required) Name of the policy bound to the LB vserver.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `(Optional) Priority.`,
				},
				resource.Attribute{
					Name:        "gotopriorityexpression",
					Description: `(Optional) Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.`,
				},
				resource.Attribute{
					Name:        "bindpoint",
					Description: `(Optional) The bindpoint to which the policy is bound. Possible values: [ REQUEST, RESPONSE ]`,
				},
				resource.Attribute{
					Name:        "invoke",
					Description: `(Optional) Invoke policies bound to a virtual server or policy label.`,
				},
				resource.Attribute{
					Name:        "labeltype",
					Description: `(Optional) The invocation type. Possible values: [ reqvserver, resvserver, policylabel ]`,
				},
				resource.Attribute{
					Name:        "labelname",
					Description: `(Optional) Name of the label invoked.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name for the virtual server. Must begin with an ASCII alphanumeric or underscore (\_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at sign (@), equal sign (=), and hyphen (-) characters. Can be changed after the virtual server is created. CLI Users: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my vserver" or 'my vserver'). . ## Attribute Reference In addition to the arguments, the following attributes are available:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of the lbvserver\_rewritepolicy\_binding. It is the concatenation of the ` + "`" + `name` + "`" + ` and ` + "`" + `policyname` + "`" + ` attributes separated by a comma. ## Import A lbvserver\_rewritepolicy\_binding can be imported using its id, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import citrixadc_lbvserver_rewritepolicy_binding.tf_bind tf_lbvserver,tf_rewrite_policy ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id of the lbvserver\_rewritepolicy\_binding. It is the concatenation of the ` + "`" + `name` + "`" + ` and ` + "`" + `policyname` + "`" + ` attributes separated by a comma. ## Import A lbvserver\_rewritepolicy\_binding can be imported using its id, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import citrixadc_lbvserver_rewritepolicy_binding.tf_bind tf_lbvserver,tf_rewrite_policy ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "citrixadc_lbvserver_service_binding",
			Category:         "Load Balancing",
			ShortDescription: ``,
			Description: `\_service\_binding

The lbvserver\_service\_binding resource is used to bind load balancing virtual servers to services.


`,
			Keywords: []string{
				"load",
				"balancing",
				"lbvserver",
				"service",
				"binding",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "servicename",
					Description: `(Required) Service to bind to the virtual server.`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `(Optional) Weight to assign to the specified service.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name for the virtual server. Must begin with an ASCII alphanumeric or underscore (\_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at sign (@), equal sign (=), and hyphen (-) characters. Can be changed after the virtual server is created. CLI Users: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my vserver" or 'my vserver'). . ## Attribute Reference In addition to the arguments, the following attributes are available:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of the lbvserver\_service\_binding. It is the concatenation of the ` + "`" + `name` + "`" + ` and ` + "`" + `servicename` + "`" + ` attributes separated by a comma. ## Import A lbvserver\_service\_binding can be imported using its id, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import citrixadc_lbvserver_service_binding.tf_binding tf_lbvserver,tf_service ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id of the lbvserver\_service\_binding. It is the concatenation of the ` + "`" + `name` + "`" + ` and ` + "`" + `servicename` + "`" + ` attributes separated by a comma. ## Import A lbvserver\_service\_binding can be imported using its id, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import citrixadc_lbvserver_service_binding.tf_binding tf_lbvserver,tf_service ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "citrixadc_lbvserver_transformpolicy_binding",
			Category:         "Load Balancing",
			ShortDescription: ``,
			Description: `\_transformpolicy\_binding

The lbvserver\_transformpolicy\_binding resource is used to bind load balancing virtual servers to transform policies.


`,
			Keywords: []string{
				"load",
				"balancing",
				"lbvserver",
				"transformpolicy",
				"binding",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "policyname",
					Description: `(Required) Name of the policy bound to the LB vserver.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `(Optional) Priority.`,
				},
				resource.Attribute{
					Name:        "gotopriorityexpression",
					Description: `(Optional) Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.`,
				},
				resource.Attribute{
					Name:        "bindpoint",
					Description: `(Optional) The bindpoint to which the policy is bound. Possible values: [ REQUEST, RESPONSE ]`,
				},
				resource.Attribute{
					Name:        "invoke",
					Description: `(Optional) Invoke policies bound to a virtual server or policy label.`,
				},
				resource.Attribute{
					Name:        "labeltype",
					Description: `(Optional) The invocation type. Possible values: [ reqvserver, resvserver, policylabel ]`,
				},
				resource.Attribute{
					Name:        "labelname",
					Description: `(Optional) Name of the label invoked.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name for the virtual server. Must begin with an ASCII alphanumeric or underscore (\_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at sign (@), equal sign (=), and hyphen (-) characters. Can be changed after the virtual server is created. CLI Users: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my vserver" or 'my vserver'). . ## Attribute Reference In addition to the arguments, the following attributes are available:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of the lbvserver\_transformpolicy\_binding. It is the concatenation of the ` + "`" + `name` + "`" + ` and ` + "`" + `policyname` + "`" + ` attributes separated by a comma. ## Import A lbvserver\_transformpolicy\_binding can be imported using its id, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import citrixadc_lbvserver_transformpolicy_binding.tf_binding tf_lbvserver,tf_trans_policy ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id of the lbvserver\_transformpolicy\_binding. It is the concatenation of the ` + "`" + `name` + "`" + ` and ` + "`" + `policyname` + "`" + ` attributes separated by a comma. ## Import A lbvserver\_transformpolicy\_binding can be imported using its id, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import citrixadc_lbvserver_transformpolicy_binding.tf_binding tf_lbvserver,tf_trans_policy ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "citrixadc_linkset",
			Category:         "Network",
			ShortDescription: ``,
			Description: `

The linkset resource is used to create Link sets.


`,
			Keywords: []string{
				"network",
				"linkset",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "linkset_id",
					Description: `(Required) Unique identifier for the linkset. Must be of the form LS/x, where x can be an integer from 1 to 32.`,
				},
				resource.Attribute{
					Name:        "interfacebinding",
					Description: `(Optional) A set of interfaces that are bound to the linkset. ## Attribute Reference In addition to the arguments, the following attributes are available:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of the linkset. It has the same value as the ` + "`" + `linkset_id` + "`" + ` attribute. ## Import A linkset can be imported using its ` + "`" + `linskset_id` + "`" + ` attribute, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import citrixadc_linkset.tf_linkset LS/1 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id of the linkset. It has the same value as the ` + "`" + `linkset_id` + "`" + ` attribute. ## Import A linkset can be imported using its ` + "`" + `linskset_id` + "`" + ` attribute, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import citrixadc_linkset.tf_linkset LS/1 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "citrixadc_netprofile",
			Category:         "Network",
			ShortDescription: ``,
			Description: `

The netprofile resource is used to create nework profiles.


`,
			Keywords: []string{
				"network",
				"netprofile",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name for the net profile. Cannot be changed after the profile is created. Choose a name that helps identify the net profile.`,
				},
				resource.Attribute{
					Name:        "td",
					Description: `(Optional) Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.`,
				},
				resource.Attribute{
					Name:        "srcip",
					Description: `(Optional) IP address or the name of an IP set.`,
				},
				resource.Attribute{
					Name:        "srcippersistency",
					Description: `(Optional) When the net profile is associated with a virtual server or its bound services, this option enables the Citrix ADC to use the same address, specified in the net profile, to communicate to servers for all sessions initiated from a particular client to the virtual server. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "overridelsn",
					Description: `(Optional) USNIP/USIP settings override LSN settings for configured service/virtual server traffic.. . Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "mbf",
					Description: `(Optional) Response will be sent using learnt info if enabled. When creating a netprofile, if you do not set this parameter, the netprofile inherits the global MBF setting (available in the enable ns mode and disable ns mode CLI commands, or in the System > Settings > Configure modes > Configure Modes dialog box). However, you can override this setting after you create the netprofile. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "proxyprotocol",
					Description: `(Optional) Proxy Protocol Action (Enabled/Disabled). Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "proxyprotocoltxversion",
					Description: `(Optional) Proxy Protocol Version (V1/V2). Possible values: [ V1, V2 ]`,
				},
				resource.Attribute{
					Name:        "proxyprotocolaftertlshandshake",
					Description: `(Optional) ## Attribute Reference In addition to the arguments, the following attributes are available:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of the netprofile. It has the same value as the ` + "`" + `name` + "`" + ` attribute. ## Import A netprofile can be imported using its name, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import citrixadc_netprofile.tf_netprofile tf_netprofile ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id of the netprofile. It has the same value as the ` + "`" + `name` + "`" + ` attribute. ## Import A netprofile can be imported using its name, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import citrixadc_netprofile.tf_netprofile tf_netprofile ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "citrixadc_nsacl",
			Category:         "NS",
			ShortDescription: ``,
			Description: `

The nsacl resource is used to create ADC acls.


`,
			Keywords: []string{
				"ns",
				"nsacl",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "aclname",
					Description: `(Optional) Name for the extended ACL rule.`,
				},
				resource.Attribute{
					Name:        "aclaction",
					Description: `(Optional) Action to perform on incoming IPv4 packets that match the extended ACL rule. Available settings function as follows:`,
				},
				resource.Attribute{
					Name:        "td",
					Description: `(Optional) Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.`,
				},
				resource.Attribute{
					Name:        "srcipop",
					Description: `(Optional) Either the equals (=) or does not equal (!=) logical operator. Possible values: [ =, !=, EQ, NEQ ]`,
				},
				resource.Attribute{
					Name:        "srcipval",
					Description: `(Optional) IP address or range of IP addresses to match against the source IP address of an incoming IPv4 packet. In the command line interface, separate the range with a hyphen. For example:10.102.29.30-10.102.29.189.`,
				},
				resource.Attribute{
					Name:        "srcportop",
					Description: `(Optional) Either the equals (=) or does not equal (!=) logical operator. Possible values: [ =, !=, EQ, NEQ ]`,
				},
				resource.Attribute{
					Name:        "srcportval",
					Description: `(Optional) Port number or range of port numbers to match against the source port number of an incoming IPv4 packet. In the command line interface, separate the range with a hyphen. For example: 40-90.`,
				},
				resource.Attribute{
					Name:        "destipop",
					Description: `(Optional) Either the equals (=) or does not equal (!=) logical operator. Possible values: [ =, !=, EQ, NEQ ]`,
				},
				resource.Attribute{
					Name:        "destipval",
					Description: `(Optional) IP address or range of IP addresses to match against the destination IP address of an incoming IPv4 packet. In the command line interface, separate the range with a hyphen. For example: 10.102.29.30-10.102.29.189.`,
				},
				resource.Attribute{
					Name:        "destportop",
					Description: `(Optional) Either the equals (=) or does not equal (!=) logical operator. Possible values: [ =, !=, EQ, NEQ ]`,
				},
				resource.Attribute{
					Name:        "destportval",
					Description: `(Optional) Port number or range of port numbers to match against the destination port number of an incoming IPv4 packet. In the command line interface, separate the range with a hyphen. For example: 40-90. Note: The destination port can be specified only for TCP and UDP protocols.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `(Optional) Number of seconds, in multiples of four, after which the extended ACL rule expires. If you do not want the extended ACL rule to expire, do not specify a TTL value.`,
				},
				resource.Attribute{
					Name:        "srcmac",
					Description: `(Optional) MAC address to match against the source MAC address of an incoming IPv4 packet.`,
				},
				resource.Attribute{
					Name:        "srcmacmask",
					Description: `(Optional) Used to define range of Source MAC address. It takes string of 0 and 1, 0s are for exact match and 1s for wildcard. For matching first 3 bytes of MAC address, srcMacMask value "000000111111". .`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Optional) Protocol to match against the protocol of an incoming IPv4 packet. Possible values: [ ICMP, IGMP, TCP, EGP, IGP, ARGUS, UDP, RDP, RSVP, EIGRP, L2TP, ISIS, GGP, IPoverIP, ST, CBT, BBN-RCC-M, NVP-II, PUP, EMCON, XNET, CHAOS, MUX, DCN-MEAS, HMP, PRM, XNS-IDP, TRUNK-1, TRUNK-2, LEAF-1, LEAF-2, IRTP, ISO-TP4, NETBLT, MFE-NSP, MERIT-INP, SEP, 3PC, IDPR, XTP, DDP, IDPR-CMTP, TP++, IL, IPv6, SDRP, IPv6-Route, IPv6-Frag, IDRP, GRE, MHRP, BNA, ESP, AH, I-NLSP, SWIPE, NARP, MOBILE, TLSP, SKIP, ICMPV6, IPv6-NoNx, IPv6-Opts, Any-Host-Internal-Protocol, CFTP, Any-Local-Network, SAT-EXPAK, KRYPTOLAN, RVD, IPPC, Any-Distributed-File-System, TFTP, VISA, IPCV, CPNX, CPHB, WSN, PVP, BR-SAT-MO, SUN-ND, WB-MON, WB-EXPAK, ISO-IP, VMTP, SECURE-VM, VINES, TTP, NSFNET-IG, DGP, TCF, OSPFIGP, Sprite-RP, LARP, MTP, AX.25, IPIP, MICP, SCC-SP, ETHERIP, Any-Private-Encryption-Scheme, GMTP, IFMP, PNNI, PIM, ARIS, SCPS, QNX, A/N, IPComp, SNP, Compaq-Pe, IPX-in-IP, VRRP, PGM, Any-0-Hop-Protocol, ENCAP, DDX, IATP, STP, SRP, UTI, SMP, SM, PTP, FIRE, CRTP, CRUDP, SSCOPMCE, IPLT, SPS, PIPE, SCTP, FC, RSVP-E2E-IGNORE, Mobility-Header, UDPLite ]`,
				},
				resource.Attribute{
					Name:        "protocolnumber",
					Description: `(Optional) Protocol to match against the protocol of an incoming IPv4 packet.`,
				},
				resource.Attribute{
					Name:        "vlan",
					Description: `(Optional) ID of the VLAN. The Citrix ADC applies the ACL rule only to the incoming packets of the specified VLAN. If you do not specify a VLAN ID, the appliance applies the ACL rule to the incoming packets on all VLANs.`,
				},
				resource.Attribute{
					Name:        "vxlan",
					Description: `(Optional) ID of the VXLAN. The Citrix ADC applies the ACL rule only to the incoming packets of the specified VXLAN. If you do not specify a VXLAN ID, the appliance applies the ACL rule to the incoming packets on all VXLANs.`,
				},
				resource.Attribute{
					Name:        "interface",
					Description: `(Optional) ID of an interface. The Citrix ADC applies the ACL rule only to the incoming packets from the specified interface. If you do not specify any value, the appliance applies the ACL rule to the incoming packets of all interfaces.`,
				},
				resource.Attribute{
					Name:        "established",
					Description: `(Optional) Allow only incoming TCP packets that have the ACK or RST bit set, if the action set for the ACL rule is ALLOW and these packets match the other conditions in the ACL rule.`,
				},
				resource.Attribute{
					Name:        "icmptype",
					Description: `(Optional) ICMP Message type to match against the message type of an incoming ICMP packet. For example, to block DESTINATION UNREACHABLE messages, you must specify 3 as the ICMP type. Note: This parameter can be specified only for the ICMP protocol.`,
				},
				resource.Attribute{
					Name:        "icmpcode",
					Description: `(Optional) Code of a particular ICMP message type to match against the ICMP code of an incoming ICMP packet. For example, to block DESTINATION HOST UNREACHABLE messages, specify 3 as the ICMP type and 1 as the ICMP code. If you set this parameter, you must set the ICMP Type parameter.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `(Optional) Priority for the extended ACL rule that determines the order in which it is evaluated relative to the other extended ACL rules. If you do not specify priorities while creating extended ACL rules, the ACL rules are evaluated in the order in which they are created.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `(Optional) Enable or disable the extended ACL rule. After you apply the extended ACL rules, the Citrix ADC compares incoming packets against the enabled extended ACL rules. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "logstate",
					Description: `(Optional) Enable or disable logging of events related to the extended ACL rule. The log messages are stored in the configured syslog or auditlog server. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "ratelimit",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "dfdhash",
					Description: `(Optional) Specifies the type hashmethod to be applied, to steer the packet to the FP of the packet. Possible values: [ SIP-SPORT-DIP-DPORT, SIP, DIP, SIP-DIP, SIP-SPORT, DIP-DPORT ]`,
				},
				resource.Attribute{
					Name:        "stateful",
					Description: `(Optional) If stateful option is enabled, transparent sessions are created for the traffic hitting this ACL and not hitting any other features like LB, INAT etc. . Possible values: [ YES, NO ] ## Attribute Reference In addition to the arguments, the following attributes are available:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of the nsacl. It has the same value as the ` + "`" + `name` + "`" + ` attribute. ## Import A nsacl can be imported using its ` + "`" + `aclname` + "`" + ` attribute, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import citrixadc_nsacl.tf_nsacl tf_nsacl ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id of the nsacl. It has the same value as the ` + "`" + `name` + "`" + ` attribute. ## Import A nsacl can be imported using its ` + "`" + `aclname` + "`" + ` attribute, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import citrixadc_nsacl.tf_nsacl tf_nsacl ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "citrixadc_nsacls",
			Category:         "NS",
			ShortDescription: ``,
			Description: `

The nsacls resource is used to create a set of ADC acls.


`,
			Keywords: []string{
				"ns",
				"nsacls",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "aclsname",
					Description: `(Optional) Name for the acls resource.`,
				},
				resource.Attribute{
					Name:        "acl",
					Description: `(Optional) A set of block defining acls. The acl block supports the following:`,
				},
				resource.Attribute{
					Name:        "aclname",
					Description: `(Required) Name for the extended ACL rule.`,
				},
				resource.Attribute{
					Name:        "aclaction",
					Description: `(Optional) Action to perform on incoming IPv4 packets that match the extended ACL rule. Available settings function as follows:`,
				},
				resource.Attribute{
					Name:        "td",
					Description: `(Optional) Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.`,
				},
				resource.Attribute{
					Name:        "srcipop",
					Description: `(Optional) Either the equals (=) or does not equal (!=) logical operator. Possible values: [ =, !=, EQ, NEQ ]`,
				},
				resource.Attribute{
					Name:        "srcipval",
					Description: `(Optional) IP address or range of IP addresses to match against the source IP address of an incoming IPv4 packet. In the command line interface, separate the range with a hyphen. For example:10.102.29.30-10.102.29.189.`,
				},
				resource.Attribute{
					Name:        "srcportop",
					Description: `(Optional) Either the equals (=) or does not equal (!=) logical operator. Possible values: [ =, !=, EQ, NEQ ]`,
				},
				resource.Attribute{
					Name:        "srcportval",
					Description: `(Optional) Port number or range of port numbers to match against the source port number of an incoming IPv4 packet. In the command line interface, separate the range with a hyphen. For example: 40-90.`,
				},
				resource.Attribute{
					Name:        "destipop",
					Description: `(Optional) Either the equals (=) or does not equal (!=) logical operator. Possible values: [ =, !=, EQ, NEQ ]`,
				},
				resource.Attribute{
					Name:        "destipval",
					Description: `(Optional) IP address or range of IP addresses to match against the destination IP address of an incoming IPv4 packet. In the command line interface, separate the range with a hyphen. For example: 10.102.29.30-10.102.29.189.`,
				},
				resource.Attribute{
					Name:        "destportop",
					Description: `(Optional) Either the equals (=) or does not equal (!=) logical operator. Possible values: [ =, !=, EQ, NEQ ]`,
				},
				resource.Attribute{
					Name:        "destportval",
					Description: `(Optional) Port number or range of port numbers to match against the destination port number of an incoming IPv4 packet. In the command line interface, separate the range with a hyphen. For example: 40-90. Note: The destination port can be specified only for TCP and UDP protocols.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `(Optional) Number of seconds, in multiples of four, after which the extended ACL rule expires. If you do not want the extended ACL rule to expire, do not specify a TTL value.`,
				},
				resource.Attribute{
					Name:        "srcmac",
					Description: `(Optional) MAC address to match against the source MAC address of an incoming IPv4 packet.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Optional) Protocol to match against the protocol of an incoming IPv4 packet. Possible values: [ ICMP, IGMP, TCP, EGP, IGP, ARGUS, UDP, RDP, RSVP, EIGRP, L2TP, ISIS, GGP, IPoverIP, ST, CBT, BBN-RCC-M, NVP-II, PUP, EMCON, XNET, CHAOS, MUX, DCN-MEAS, HMP, PRM, XNS-IDP, TRUNK-1, TRUNK-2, LEAF-1, LEAF-2, IRTP, ISO-TP4, NETBLT, MFE-NSP, MERIT-INP, SEP, 3PC, IDPR, XTP, DDP, IDPR-CMTP, TP++, IL, IPv6, SDRP, IPv6-Route, IPv6-Frag, IDRP, GRE, MHRP, BNA, ESP, AH, I-NLSP, SWIPE, NARP, MOBILE, TLSP, SKIP, ICMPV6, IPv6-NoNx, IPv6-Opts, Any-Host-Internal-Protocol, CFTP, Any-Local-Network, SAT-EXPAK, KRYPTOLAN, RVD, IPPC, Any-Distributed-File-System, TFTP, VISA, IPCV, CPNX, CPHB, WSN, PVP, BR-SAT-MO, SUN-ND, WB-MON, WB-EXPAK, ISO-IP, VMTP, SECURE-VM, VINES, TTP, NSFNET-IG, DGP, TCF, OSPFIGP, Sprite-RP, LARP, MTP, AX.25, IPIP, MICP, SCC-SP, ETHERIP, Any-Private-Encryption-Scheme, GMTP, IFMP, PNNI, PIM, ARIS, SCPS, QNX, A/N, IPComp, SNP, Compaq-Pe, IPX-in-IP, VRRP, PGM, Any-0-Hop-Protocol, ENCAP, DDX, IATP, STP, SRP, UTI, SMP, SM, PTP, FIRE, CRTP, CRUDP, SSCOPMCE, IPLT, SPS, PIPE, SCTP, FC, RSVP-E2E-IGNORE, Mobility-Header, UDPLite ]`,
				},
				resource.Attribute{
					Name:        "protocolnumber",
					Description: `(Optional) Protocol to match against the protocol of an incoming IPv4 packet.`,
				},
				resource.Attribute{
					Name:        "vlan",
					Description: `(Optional) ID of the VLAN. The Citrix ADC applies the ACL rule only to the incoming packets of the specified VLAN. If you do not specify a VLAN ID, the appliance applies the ACL rule to the incoming packets on all VLANs.`,
				},
				resource.Attribute{
					Name:        "interface",
					Description: `(Optional) ID of an interface. The Citrix ADC applies the ACL rule only to the incoming packets from the specified interface. If you do not specify any value, the appliance applies the ACL rule to the incoming packets of all interfaces.`,
				},
				resource.Attribute{
					Name:        "established",
					Description: `(Optional) Allow only incoming TCP packets that have the ACK or RST bit set, if the action set for the ACL rule is ALLOW and these packets match the other conditions in the ACL rule.`,
				},
				resource.Attribute{
					Name:        "icmptype",
					Description: `(Optional) ICMP Message type to match against the message type of an incoming ICMP packet. For example, to block DESTINATION UNREACHABLE messages, you must specify 3 as the ICMP type. Note: This parameter can be specified only for the ICMP protocol.`,
				},
				resource.Attribute{
					Name:        "icmpcode",
					Description: `(Optional) Code of a particular ICMP message type to match against the ICMP code of an incoming ICMP packet. For example, to block DESTINATION HOST UNREACHABLE messages, specify 3 as the ICMP type and 1 as the ICMP code. If you set this parameter, you must set the ICMP Type parameter.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `(Optional) Priority for the extended ACL rule that determines the order in which it is evaluated relative to the other extended ACL rules. If you do not specify priorities while creating extended ACL rules, the ACL rules are evaluated in the order in which they are created.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `(Optional) Enable or disable the extended ACL rule. After you apply the extended ACL rules, the Citrix ADC compares incoming packets against the enabled extended ACL rules. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "logstate",
					Description: `(Optional) Enable or disable logging of events related to the extended ACL rule. The log messages are stored in the configured syslog or auditlog server. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "ratelimit",
					Description: `(Optional) ## Attribute Reference In addition to the arguments, the following attributes are available:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of the nsacls. It has the same value as the ` + "`" + `aclsname` + "`" + ` attribute.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id of the nsacls. It has the same value as the ` + "`" + `aclsname` + "`" + ` attribute.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "citrixadc_nscapacity",
			Category:         "NS",
			ShortDescription: ``,
			Description: `

The nscapacity resource is used to apply licenses to a target ADC from a license server.


`,
			Keywords: []string{
				"ns",
				"nscapacity",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "bandwidth",
					Description: `(Optional) System bandwidth limit.`,
				},
				resource.Attribute{
					Name:        "platform",
					Description: `(Optional) appliance platform type. Possible values: [ VS10, VE10, VP10, VS25, VE25, VP25, VS200, VE200, VP200, VS1000, VE1000, VP1000, VS3000, VE3000, VP3000, VS5000, VE5000, VP5000, VS8000, VE8000, VP8000, VS10000, VE10000, VP10000, VS15000, VE15000, VP15000, VS25000, VE25000, VP25000, VS40000, VE40000, VP40000, VS100000, VE100000, VP100000, CP1000 ]`,
				},
				resource.Attribute{
					Name:        "vcpu",
					Description: `(Optional) licensed using vcpu pool.`,
				},
				resource.Attribute{
					Name:        "edition",
					Description: `(Optional) Product edition. Possible values: [ Standard, Enterprise, Platinum ]`,
				},
				resource.Attribute{
					Name:        "unit",
					Description: `(Optional) Bandwidth unit. Possible values: [ Gbps, Mbps ]`,
				},
				resource.Attribute{
					Name:        "nodeid",
					Description: `(Optional) Unique number that identifies the cluster node.`,
				},
				resource.Attribute{
					Name:        "reboot_timeout",
					Description: `(Optional) Time duration to wait for ADC to be available after reboot.`,
				},
				resource.Attribute{
					Name:        "poll_delay",
					Description: `(Optional) Time duration to wait for first poll after ADC reboot.`,
				},
				resource.Attribute{
					Name:        "poll_interval",
					Description: `(Optional) Time duration between polls after reboot.`,
				},
				resource.Attribute{
					Name:        "poll_timeout",
					Description: `(Optional) Time duration for each poll to timeout. ## Attribute Reference In addition to the arguments, the following attributes are available:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of the nscapacity. It is a unique string prefixed with ` + "`" + `tf-nscapacity-` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id of the nscapacity. It is a unique string prefixed with ` + "`" + `tf-nscapacity-` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "citrixadc_nsconfig_clear",
			Category:         "NS",
			ShortDescription: ``,
			Description: `

The nsconfig_clear resource is used to apply the clear operation for ns config.


`,
			Keywords: []string{
				"ns",
				"nsconfig",
				"clear",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "force",
					Description: `(Required) Configurations will be cleared without prompting for confirmation.`,
				},
				resource.Attribute{
					Name:        "level",
					Description: `(Required) Types of configurations to be cleared.`,
				},
				resource.Attribute{
					Name:        "rbaconfig",
					Description: `(Required) RBA configurations and TACACS policies bound to system global will not be cleared if RBA is set to NO.This option is applicable only for BASIC level of clear configuration.Default is YES, which will clear rba configurations. Possible values: [ YES, NO ]`,
				},
				resource.Attribute{
					Name:        "timestamp",
					Description: `(Required) the timestamp of the operation. Can be any string. Used to force the operation again if all other attributes have remained the same. ## Attribute Reference In addition to the arguments, the following attributes are available:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of the nsconfig_clear. It has the same value as the ` + "`" + `timestamp` + "`" + ` attribute.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id of the nsconfig_clear. It has the same value as the ` + "`" + `timestamp` + "`" + ` attribute.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "citrixadc_nsconfig_save",
			Category:         "NS",
			ShortDescription: ``,
			Description: `

The nsconfig_save resource is used to apply the save operation for ns config.


`,
			Keywords: []string{
				"ns",
				"nsconfig",
				"save",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "all",
					Description: `(Optional) Use this option to do saveconfig for all partitions.`,
				},
				resource.Attribute{
					Name:        "timestamp",
					Description: `(Required) the timestamp of the operation. Can be any string. Used to force the operation again if all other attributes have remained the same.`,
				},
				resource.Attribute{
					Name:        "concurrent_save_ok",
					Description: `(Optional) Boolean value signifying if a concurrent save error should be toleratted. When set to ` + "`" + `true` + "`" + ` a process of retries will take place waiting for the resource to return no error.`,
				},
				resource.Attribute{
					Name:        "concurrent_save_retries",
					Description: `(Optional) Number of retries after which we throw an error for the concurrent save error code.`,
				},
				resource.Attribute{
					Name:        "concurrent_save_timeout",
					Description: `(Optional) Time period after which we throw an error for the concurrent save error code.`,
				},
				resource.Attribute{
					Name:        "concurrent_save_interval",
					Description: `(Optional) Time period between tries to save the resource when processing the save error workflow. ## Attribute Reference In addition to the arguments, the following attributes are available:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of the nsconfig_save. It has the same value as the ` + "`" + `timestamp` + "`" + ` attribute.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id of the nsconfig_save. It has the same value as the ` + "`" + `timestamp` + "`" + ` attribute.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "citrixadc_nsconfig_update",
			Category:         "NS",
			ShortDescription: ``,
			Description: `

The nsconfig_update resource is used to apply the update operation for ns config.


`,
			Keywords: []string{
				"ns",
				"nsconfig",
				"update",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ipaddress",
					Description: `(Optional) IP address of the Citrix ADC. Commonly referred to as NSIP address. This parameter is mandatory to bring up the appliance.`,
				},
				resource.Attribute{
					Name:        "netmask",
					Description: `(Optional) Netmask corresponding to the IP address. This parameter is mandatory to bring up the appliance.`,
				},
				resource.Attribute{
					Name:        "nsvlan",
					Description: `(Optional) VLAN (NSVLAN) for the subnet on which the IP address resides.`,
				},
				resource.Attribute{
					Name:        "ifnum",
					Description: `(Optional) Interfaces of the appliances that must be bound to the NSVLAN.`,
				},
				resource.Attribute{
					Name:        "tagged",
					Description: `(Optional) Specifies that the interfaces will be added as 802.1q tagged interfaces. Packets sent on these interface on this VLAN will have an additional 4-byte 802.1q tag which identifies the VLAN. To use 802.1q tagging, the switch connected to the appliance's interfaces must also be configured for tagging. Possible values: [ YES, NO ] ## Attribute Reference In addition to the arguments, the following attributes are available:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of the nsconfig_update. It is a random string prefixed with "tf-nsconfig-update"`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id of the nsconfig_update. It is a random string prefixed with "tf-nsconfig-update"`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "citrixadc_nsfeature",
			Category:         "NS",
			ShortDescription: ``,
			Description: `

The nsfeature resource is used to enable or disable ADC features.


`,
			Keywords: []string{
				"ns",
				"nsfeature",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "wl",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "sp",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "lb",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "cs",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "cr",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "cmp",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "pq",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "ssl",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "gslb",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "hdosp",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "cf",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "ic",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "sslvpn",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "aaa",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "ospf",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "rip",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "bgp",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "rewrite",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "ipv6pt",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "appfw",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "responder",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "htmlinjection",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "push",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "appflow",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "cloudbridge",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "isis",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "ch",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "appqoe",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "contentaccelerator",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "rise",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "feo",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "lsn",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "rdpproxy",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "rep",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "urlfiltering",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "videooptimization",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "forwardproxy",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "sslinterception",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "adaptivetcp",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "cqa",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "ci",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "bot",
					Description: `(Optional) ## Attribute Reference In addition to the arguments, the following attributes are available:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of the nsfeature resource. It is a random string prefixed with "tf-nsfeature-"`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id of the nsfeature resource. It is a random string prefixed with "tf-nsfeature-"`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "citrixadc_nshttpprofile",
			Category:         "NS",
			ShortDescription: ``,
			Description: `

The nshttpprofile resource is used to create HTTP profiles.


`,
			Keywords: []string{
				"ns",
				"nshttpprofile",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name for an HTTP profile.The name of a HTTP profile cannot be changed after it is created.`,
				},
				resource.Attribute{
					Name:        "dropinvalreqs",
					Description: `(Optional) Drop invalid HTTP requests or responses. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "markhttp09inval",
					Description: `(Optional) Mark HTTP/0.9 requests as invalid. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "markconnreqinval",
					Description: `(Optional) Mark CONNECT requests as invalid. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "marktracereqinval",
					Description: `(Optional) Mark TRACE requests as invalid. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "cmponpush",
					Description: `(Optional) Start data compression on receiving a TCP packet with PUSH flag set. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "conmultiplex",
					Description: `(Optional) Reuse server connections for requests from more than one client connections. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "maxreusepool",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "dropextracrlf",
					Description: `(Optional) Drop any extra 'CR' and 'LF' characters present after the header. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "incomphdrdelay",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "websocket",
					Description: `(Optional) HTTP connection to be upgraded to a web socket connection. Once upgraded, Citrix ADC does not process Layer 7 traffic on this connection. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "rtsptunnel",
					Description: `(Optional) Allow RTSP tunnel in HTTP. Once application/x-rtsp-tunnelled is seen in Accept or Content-Type header, Citrix ADC does not process Layer 7 traffic on this connection. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "reqtimeout",
					Description: `(Optional) Time, in seconds, within which the HTTP request must complete. If the request does not complete within this time, the specified request timeout action is executed. Zero disables the timeout.`,
				},
				resource.Attribute{
					Name:        "adpttimeout",
					Description: `(Optional) Adapts the configured request timeout based on flow conditions. The timeout is increased or decreased internally and applied on the flow. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "reqtimeoutaction",
					Description: `(Optional) Action to take when the HTTP request does not complete within the specified request timeout duration. You can configure the following actions:`,
				},
				resource.Attribute{
					Name:        "dropextradata",
					Description: `(Optional) Drop any extra data when server sends more data than the specified content-length. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "weblog",
					Description: `(Optional) Enable or disable web logging. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "clientiphdrexpr",
					Description: `(Optional) Name of the header that contains the real client IP address.`,
				},
				resource.Attribute{
					Name:        "maxreq",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "persistentetag",
					Description: `(Optional) Generate the persistent Citrix ADC specific ETag for the HTTP response with ETag header. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "spdy",
					Description: `(Optional) Enable SPDYv2 or SPDYv3 or both over SSL vserver. SSL will advertise SPDY support either during NPN Handshake or when client will advertises SPDY support during ALPN handshake. Both SPDY versions are enabled when this parameter is set to ENABLED. Possible values: [ DISABLED, ENABLED, V2, V3 ]`,
				},
				resource.Attribute{
					Name:        "http2",
					Description: `(Optional) Choose whether to enable support for HTTP/2. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "http2direct",
					Description: `(Optional) Choose whether to enable support for Direct HTTP/2. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "http2strictcipher",
					Description: `(Optional) Choose whether to enable strict HTTP/2 cipher selection. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "altsvc",
					Description: `(Optional) Choose whether to enable support for Alternative Service. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "reusepooltimeout",
					Description: `(Optional) Idle timeout (in seconds) for server connections in re-use pool. Connections in the re-use pool are flushed, if they remain idle for the configured timeout.`,
				},
				resource.Attribute{
					Name:        "maxheaderlen",
					Description: `(Optional) Number of bytes to be queued to look for complete header before returning error. If complete header is not obtained after queuing these many bytes, request will be marked as invalid and no L7 processing will be done for that TCP connection.`,
				},
				resource.Attribute{
					Name:        "minreusepool",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "http2maxheaderlistsize",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "http2maxframesize",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "http2maxconcurrentstreams",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "http2initialwindowsize",
					Description: `(Optional) Initial window size for stream level flow control, in bytes.`,
				},
				resource.Attribute{
					Name:        "http2headertablesize",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "http2minseverconn",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "apdexcltresptimethreshold",
					Description: `(Optional) This option sets the satisfactory threshold (T) for client response time in milliseconds to be used for APDEX calculations. This means a transaction responding in less than this threshold is considered satisfactory. Transaction responding between T and 4\`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of the nshttpprofile. It has the same value as the ` + "`" + `name` + "`" + ` attribute. ## Import A nshttpprofile can be imported using its name, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import citrixadc_nshttpprofile.tf_httpprofile tf_httpprofile ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id of the nshttpprofile. It has the same value as the ` + "`" + `name` + "`" + ` attribute. ## Import A nshttpprofile can be imported using its name, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import citrixadc_nshttpprofile.tf_httpprofile tf_httpprofile ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "citrixadc_nsip",
			Category:         "NS",
			ShortDescription: ``,
			Description: `

The nsip resource is used to create nsip, snip and vip, ipv4 addresses for the ADC.


`,
			Keywords: []string{
				"ns",
				"nsip",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ipaddress",
					Description: `(Optional) IPv4 address to create on the Citrix ADC. Cannot be changed after the IP address is created.`,
				},
				resource.Attribute{
					Name:        "netmask",
					Description: `(Optional) Subnet mask associated with the IP address.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) Type of the IP address to create on the Citrix ADC. Cannot be changed after the IP address is created. The following are the different types of Citrix ADC owned IP addresses:`,
				},
				resource.Attribute{
					Name:        "arp",
					Description: `(Optional) Respond to ARP requests for this IP address. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "icmp",
					Description: `(Optional) Respond to ICMP requests for this IP address. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "vserver",
					Description: `(Optional) Use this option to set (enable or disable) the virtual server attribute for this IP address. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "telnet",
					Description: `(Optional) Allow Telnet access to this IP address. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "ftp",
					Description: `(Optional) Allow File Transfer Protocol (FTP) access to this IP address. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "gui",
					Description: `(Optional) Allow graphical user interface (GUI) access to this IP address. Possible values: [ ENABLED, SECUREONLY, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "ssh",
					Description: `(Optional) Allow secure shell (SSH) access to this IP address. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "snmp",
					Description: `(Optional) Allow Simple Network Management Protocol (SNMP) access to this IP address. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "mgmtaccess",
					Description: `(Optional) Allow access to management applications on this IP address. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "restrictaccess",
					Description: `(Optional) Block access to nonmanagement applications on this IP. This option is applicable for MIPs, SNIPs, and NSIP, and is disabled by default. Nonmanagement applications can run on the underlying Citrix ADC Free BSD operating system. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "dynamicrouting",
					Description: `(Optional) Allow dynamic routing on this IP address. Specific to Subnet IP (SNIP) address. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "decrementttl",
					Description: `(Optional) Decrement TTL by 1 when ENABLED.This setting is applicable only for UDP traffic. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "ospf",
					Description: `(Optional) Use this option to enable or disable OSPF on this IP address for the entity. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "bgp",
					Description: `(Optional) Use this option to enable or disable BGP on this IP address for the entity. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "rip",
					Description: `(Optional) Use this option to enable or disable RIP on this IP address for the entity. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "hostroute",
					Description: `(Optional) Option to push the VIP to ZebOS routing table for Kernel route redistribution through dynamic routing protocols. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "advertiseondefaultpartition",
					Description: `(Optional) Advertise VIPs from Shared VLAN on Default Partition. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "networkroute",
					Description: `(Optional) Option to push the SNIP subnet to ZebOS routing table for Kernel route redistribution through dynamic routing protocol. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `(Optional) Tag value for the network/host route associated with this IP.`,
				},
				resource.Attribute{
					Name:        "hostrtgw",
					Description: `(Optional) IP address of the gateway of the route for this VIP address.`,
				},
				resource.Attribute{
					Name:        "metric",
					Description: `(Optional) Integer value to add to or subtract from the cost of the route advertised for the VIP address.`,
				},
				resource.Attribute{
					Name:        "vserverrhilevel",
					Description: `(Optional) Advertise the route for the Virtual IP (VIP) address on the basis of the state of the virtual servers associated with that VIP.`,
				},
				resource.Attribute{
					Name:        "ospflsatype",
					Description: `(Optional) Type of LSAs to be used by the OSPF protocol, running on the Citrix ADC, for advertising the route for this VIP address. Possible values: [ TYPE1, TYPE5 ]`,
				},
				resource.Attribute{
					Name:        "ospfarea",
					Description: `(Optional) ID of the area in which the type1 link-state advertisements (LSAs) are to be advertised for this virtual IP (VIP) address by the OSPF protocol running on the Citrix ADC. When this parameter is not set, the VIP is advertised on all areas.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `(Optional) Enable or disable the IP address. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "vrid",
					Description: `(Optional) A positive integer that uniquely identifies a VMAC address for binding to this VIP address. This binding is used to set up Citrix ADCs in an active-active configuration using VRRP.`,
				},
				resource.Attribute{
					Name:        "icmpresponse",
					Description: `(Optional) Respond to ICMP requests for a Virtual IP (VIP) address on the basis of the states of the virtual servers associated with that VIP. Available settings function as follows:`,
				},
				resource.Attribute{
					Name:        "ownernode",
					Description: `(Optional) The owner node in a Cluster for this IP address. Owner node can vary from 0 to 31. If ownernode is not specified then the IP is treated as Striped IP.`,
				},
				resource.Attribute{
					Name:        "arpresponse",
					Description: `(Optional) Respond to ARP requests for a Virtual IP (VIP) address on the basis of the states of the virtual servers associated with that VIP. Available settings function as follows:`,
				},
				resource.Attribute{
					Name:        "ownerdownresponse",
					Description: `(Optional) in cluster system, if the owner node is down, whether should it respond to icmp/arp. Possible values: [ YES, NO ]`,
				},
				resource.Attribute{
					Name:        "td",
					Description: `(Optional) Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0. TD id 4095 is used reserved for LSN use .`,
				},
				resource.Attribute{
					Name:        "vserverrhimode",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "mptcpadvertise",
					Description: `(Optional) ## Attribute Reference In addition to the arguments, the following attributes are available:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of the nsip. It has the same value as the ` + "`" + `ipaddress` + "`" + ` attribute. ## Import A nsip can be imported using its name, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import citrixadc_nsip.tf_nsip 192.168.2.55 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id of the nsip. It has the same value as the ` + "`" + `ipaddress` + "`" + ` attribute. ## Import A nsip can be imported using its name, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import citrixadc_nsip.tf_nsip 192.168.2.55 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "citrixadc_nsip6",
			Category:         "NS",
			ShortDescription: ``,
			Description: `

The nsip6 resource is used to create nsip, snip and vip ipv6 addresses for ADC.


`,
			Keywords: []string{
				"ns",
				"nsip6",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ipv6address",
					Description: `(Optional) IPv6 address to create on the Citrix ADC.`,
				},
				resource.Attribute{
					Name:        "scope",
					Description: `(Optional) Scope of the IPv6 address to be created. Cannot be changed after the IP address is created. Possible values: [ global, link-local ]`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) Type of IP address to be created on the Citrix ADC. Cannot be changed after the IP address is created. Possible values: [ NSIP, VIP, SNIP, GSLBsiteIP, ADNSsvcIP, RADIUSListenersvcIP, CLIP ]`,
				},
				resource.Attribute{
					Name:        "vlan",
					Description: `(Optional) The VLAN number.`,
				},
				resource.Attribute{
					Name:        "nd",
					Description: `(Optional) Respond to Neighbor Discovery (ND) requests for this IP address. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "icmp",
					Description: `(Optional) Respond to ICMP requests for this IP address. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "vserver",
					Description: `(Optional) Enable or disable the state of all the virtual servers associated with this VIP6 address. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "telnet",
					Description: `(Optional) Allow Telnet access to this IP address. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "ftp",
					Description: `(Optional) Allow File Transfer Protocol (FTP) access to this IP address. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "gui",
					Description: `(Optional) Allow graphical user interface (GUI) access to this IP address. Possible values: [ ENABLED, SECUREONLY, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "ssh",
					Description: `(Optional) Allow secure Shell (SSH) access to this IP address. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "snmp",
					Description: `(Optional) Allow Simple Network Management Protocol (SNMP) access to this IP address. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "mgmtaccess",
					Description: `(Optional) Allow access to management applications on this IP address. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "restrictaccess",
					Description: `(Optional) Block access to nonmanagement applications on this IP address. This option is applicable forMIP6s, SNIP6s, and NSIP6s, and is disabled by default. Nonmanagement applications can run on the underlying Citrix ADC Free BSD operating system. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "dynamicrouting",
					Description: `(Optional) Allow dynamic routing on this IP address. Specific to Subnet IPv6 (SNIP6) address. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "decrementhoplimit",
					Description: `(Optional) Decrement Hop Limit by 1 when ENABLED.This setting is applicable only for UDP traffic. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "hostroute",
					Description: `(Optional) Option to push the VIP6 to ZebOS routing table for Kernel route redistribution through dynamic routing protocols. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "advertiseondefaultpartition",
					Description: `(Optional) Advertise VIPs from Shared VLAN on Default Partition. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "networkroute",
					Description: `(Optional) Option to push the SNIP6 subnet to ZebOS routing table for Kernel route redistribution through dynamic routing protocol. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `(Optional) Tag value for the network/host route associated with this IP.`,
				},
				resource.Attribute{
					Name:        "ip6hostrtgw",
					Description: `(Optional) IPv6 address of the gateway for the route. If Gateway is not set, VIP uses :: as the gateway.`,
				},
				resource.Attribute{
					Name:        "metric",
					Description: `(Optional) Integer value to add to or subtract from the cost of the route advertised for the VIP6 address.`,
				},
				resource.Attribute{
					Name:        "vserverrhilevel",
					Description: `(Optional) Advertise or do not advertise the route for the Virtual IP (VIP6) address on the basis of the state of the virtual servers associated with that VIP6.`,
				},
				resource.Attribute{
					Name:        "ospf6lsatype",
					Description: `(Optional) Type of LSAs to be used by the IPv6 OSPF protocol, running on the Citrix ADC, for advertising the route for the VIP6 address. Possible values: [ INTRA_AREA, EXTERNAL ]`,
				},
				resource.Attribute{
					Name:        "ospfarea",
					Description: `(Optional) ID of the area in which the Intra-Area-Prefix LSAs are to be advertised for the VIP6 address by the IPv6 OSPF protocol running on the Citrix ADC. When ospfArea is not set, VIP6 is advertised on all areas.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `(Optional) Enable or disable the IP address. Possible values: [ DISABLED, ENABLED ]`,
				},
				resource.Attribute{
					Name:        "map",
					Description: `(Optional) Mapped IPV4 address for the IPV6 address.`,
				},
				resource.Attribute{
					Name:        "vrid6",
					Description: `(Optional) A positive integer that uniquely identifies a VMAC address for binding to this VIP address. This binding is used to set up Citrix ADCs in an active-active configuration using VRRP.`,
				},
				resource.Attribute{
					Name:        "ownernode",
					Description: `(Optional) ID of the cluster node for which you are adding the IP address. Must be used if you want the IP address to be active only on the specific node. Can be configured only through the cluster IP address. Cannot be changed after the IP address is created.`,
				},
				resource.Attribute{
					Name:        "ownerdownresponse",
					Description: `(Optional) in cluster system, if the owner node is down, whether should it respond to icmp/arp. Possible values: [ YES, NO ]`,
				},
				resource.Attribute{
					Name:        "td",
					Description: `(Optional) Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.`,
				},
				resource.Attribute{
					Name:        "mptcpadvertise",
					Description: `(Optional) ## Attribute Reference In addition to the arguments, the following attributes are available:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of the nsip6. It has the same value as the ` + "`" + `ipv6address` + "`" + ` attribute. ## Import A nsip6 can be imported using its ipv6address, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import citrixadc_nsip6.tf_nsip6 2002:db8:100::ff/64 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id of the nsip6. It has the same value as the ` + "`" + `ipv6address` + "`" + ` attribute. ## Import A nsip6 can be imported using its ipv6address, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import citrixadc_nsip6.tf_nsip6 2002:db8:100::ff/64 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "citrixadc_nslicense",
			Category:         "NS",
			ShortDescription: ``,
			Description: `

The nslicense resource is used to upload and apply a license file to the target ADC.


`,
			Keywords: []string{
				"ns",
				"nslicense",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "license_file",
					Description: `(Required) License file to upload.`,
				},
				resource.Attribute{
					Name:        "ssh_host",
					Description: `(Optional) The ssh host ip address that will be used for the sftp transfer of the license file.`,
				},
				resource.Attribute{
					Name:        "ssh_username",
					Description: `(Optional) The user name for the ssh connection.`,
				},
				resource.Attribute{
					Name:        "ssh_password",
					Description: `(Optional) The password for the ssh connection.`,
				},
				resource.Attribute{
					Name:        "ssh_port",
					Description: `(Optional) The port for the ssh connection.`,
				},
				resource.Attribute{
					Name:        "ssh_host_pubkey",
					Description: `(Optional) The ADC public ssh host key.`,
				},
				resource.Attribute{
					Name:        "reboot",
					Description: `(Optional) Set this to true to reboot and wait for the ADC to become responsive.`,
				},
				resource.Attribute{
					Name:        "poll_delay",
					Description: `(Optional) Time to wait before the first poll after reboot. Defaults to "60s".`,
				},
				resource.Attribute{
					Name:        "poll_interval",
					Description: `(Optional) Interval between polls. Defaults to "60s".`,
				},
				resource.Attribute{
					Name:        "poll_timeout",
					Description: `(Optional) Timeout for a poll attempt. Default to "10s". ## Attribute Reference In addition to the arguments, the following attributes are available:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of the nslicense. It has the same value as the ` + "`" + `license_file` + "`" + ` attribute.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id of the nslicense. It has the same value as the ` + "`" + `license_file` + "`" + ` attribute.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "citrixadc_nslicenseserver",
			Category:         "NS",
			ShortDescription: ``,
			Description: `

The nslicenseserver resource is used to create license server entry in ADC.


`,
			Keywords: []string{
				"ns",
				"nslicenseserver",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "servername",
					Description: `(Required) Fully qualified domain name of the License server.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional) License server port.`,
				},
				resource.Attribute{
					Name:        "forceupdateip",
					Description: `(Optional) If this flag is used while adding the licenseserver, existing config will be overwritten. Use this flag only if you are sure that the new licenseserver has the required capacity.`,
				},
				resource.Attribute{
					Name:        "nodeid",
					Description: `(Optional) Unique number that identifies the cluster node. ## Attribute Reference In addition to the arguments, the following attributes are available:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of the nslicenseserver. It has the same value as the ` + "`" + `servername` + "`" + ` attribute.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id of the nslicenseserver. It has the same value as the ` + "`" + `servername` + "`" + ` attribute.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "citrixadc_nsparam",
			Category:         "NS",
			ShortDescription: ``,
			Description: `

The nsparam resource is used to set parameters to the target ADC.


`,
			Keywords: []string{
				"ns",
				"nsparam",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "maxconn",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "maxreq",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "cip",
					Description: `(Optional) Enable or disable the insertion of the actual client IP address into the HTTP header request passed from the client to one, some, or all servers attached to the system. The passed address can then be accessed through a minor modification to the server.`,
				},
				resource.Attribute{
					Name:        "cipheader",
					Description: `(Optional) Text that will be used as the client IP address header.`,
				},
				resource.Attribute{
					Name:        "cookieversion",
					Description: `(Optional) Version of the cookie inserted by the system. Possible values: [ 0, 1 ]`,
				},
				resource.Attribute{
					Name:        "securecookie",
					Description: `(Optional) Enable or disable secure flag for persistence cookie. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "pmtumin",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "pmtutimeout",
					Description: `(Optional) Interval, in minutes, for flushing the PMTU entries.`,
				},
				resource.Attribute{
					Name:        "ftpportrange",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "crportrange",
					Description: `(Optional) Port range for cache redirection services.`,
				},
				resource.Attribute{
					Name:        "timezone",
					Description: `(Optional) Time zone for the Citrix ADC. Name of the time zone should be specified as argument.`,
				},
				resource.Attribute{
					Name:        "grantquotamaxclient",
					Description: `(Optional) Percentage of shared quota to be granted at a time for maxClient.`,
				},
				resource.Attribute{
					Name:        "exclusivequotamaxclient",
					Description: `(Optional) Percentage of maxClient to be given to PEs.`,
				},
				resource.Attribute{
					Name:        "grantquotaspillover",
					Description: `(Optional) Percentage of shared quota to be granted at a time for spillover.`,
				},
				resource.Attribute{
					Name:        "exclusivequotaspillover",
					Description: `(Optional) Percentage of maximum limit to be given to PEs.`,
				},
				resource.Attribute{
					Name:        "useproxyport",
					Description: `(Optional) Enable/Disable use_proxy_port setting. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "internaluserlogin",
					Description: `(Optional) Enables/disables the internal user from logging in to the appliance. Before disabling internal user login, you must have key-based authentication set up on the appliance. The file name for the key pair must be "ns_comm_key". Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "aftpallowrandomsourceport",
					Description: `(Optional) Allow the FTP server to come from a random source port for active FTP data connections. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "tcpcip",
					Description: `(Optional) Enable or disable the insertion of the client TCP/IP header in TCP payload passed from the client to one, some, or all servers attached to the system. The passed address can then be accessed through a minor modification to the server. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "servicepathingressvlan",
					Description: `(Optional) VLAN on which the subscriber traffic arrives on the appliance.`,
				},
				resource.Attribute{
					Name:        "mgmthttpport",
					Description: `(Optional) This allow the configuration of management HTTP port.`,
				},
				resource.Attribute{
					Name:        "mgmthttpsport",
					Description: `(Optional) This allows the configuration of management HTTPS port.`,
				},
				resource.Attribute{
					Name:        "proxyprotocol",
					Description: `(Optional) Disable/Enable v1 or v2 proxy protocol header for client info insertion. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "advancedanalyticsstats",
					Description: `(Optional) Disable/Enable advanace analytics stats. Possible values: [ ENABLED, DISABLED ] ## Attribute Reference In addition to the arguments, the following attributes are available:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of the nsparam. It is a random string prefixed with "tf-nsparam-"`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id of the nsparam. It is a random string prefixed with "tf-nsparam-"`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "citrixadc_nsrpcnode",
			Category:         "NS",
			ShortDescription: ``,
			Description: `

The nsrpcnode resource is used to manage rpc nodes.


`,
			Keywords: []string{
				"ns",
				"nsrpcnode",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ipaddress",
					Description: `(Required) IP address of the node. This has to be in the same subnet as the NSIP address.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional) Password to be used in authentication with the peer system node.`,
				},
				resource.Attribute{
					Name:        "srcip",
					Description: `(Optional) Source IP address to be used to communicate with the peer system node. The default value is 0, which means that the appliance uses the NSIP address as the source IP address.`,
				},
				resource.Attribute{
					Name:        "secure",
					Description: `(Optional) State of the channel when talking to the node. Possible values: [ YES, NO ] ## Attribute Reference In addition to the arguments, the following attributes are available:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of the nsrpcnode. It has the same value as the ` + "`" + `ipaddress` + "`" + ` attribute. ## Import A nsrpcnode can be imported using its ipaddress, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import citrixadc_nsrpcnode.tf_nsrpcnode 10.78.60.201 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id of the nsrpcnode. It has the same value as the ` + "`" + `ipaddress` + "`" + ` attribute. ## Import A nsrpcnode can be imported using its ipaddress, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import citrixadc_nsrpcnode.tf_nsrpcnode 10.78.60.201 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "citrixadc_nstcpparam",
			Category:         "NS",
			ShortDescription: ``,
			Description: `

The nstcpparam resource is used to update the ADC tcp parameters.


`,
			Keywords: []string{
				"ns",
				"nstcpparam",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ws",
					Description: `(Optional) Enable or disable window scaling. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "wsval",
					Description: `(Optional) Factor used to calculate the new window size. This argument is needed only when the window scaling is enabled.`,
				},
				resource.Attribute{
					Name:        "sack",
					Description: `(Optional) Enable or disable Selective ACKnowledgement (SACK). Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "learnvsvrmss",
					Description: `(Optional) Enable or disable maximum segment size (MSS) learning for virtual servers. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "maxburst",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "initialcwnd",
					Description: `(Optional) Initial maximum upper limit on the number of TCP packets that can be outstanding on the TCP link to the server.`,
				},
				resource.Attribute{
					Name:        "recvbuffsize",
					Description: `(Optional) TCP Receive buffer size.`,
				},
				resource.Attribute{
					Name:        "delayedack",
					Description: `(Optional) Timeout for TCP delayed ACK, in milliseconds.`,
				},
				resource.Attribute{
					Name:        "downstaterst",
					Description: `(Optional) Flag to switch on RST on down services. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "nagle",
					Description: `(Optional) Enable or disable the Nagle algorithm on TCP connections. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "limitedpersist",
					Description: `(Optional) Limit the number of persist (zero window) probes. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "oooqsize",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "ackonpush",
					Description: `(Optional) Send immediate positive acknowledgement (ACK) on receipt of TCP packets with PUSH flag. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "maxpktpermss",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "pktperretx",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "minrto",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "slowstartincr",
					Description: `(Optional) Multiplier that determines the rate at which slow start increases the size of the TCP transmission window after each acknowledgement of successful transmission.`,
				},
				resource.Attribute{
					Name:        "maxdynserverprobes",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "synholdfastgiveup",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "maxsynholdperprobe",
					Description: `(Optional) Limit the number of client connections (SYN) waiting for status of single probe. Any new SYN packets will be dropped.`,
				},
				resource.Attribute{
					Name:        "maxsynhold",
					Description: `(Optional) Limit the number of client connections (SYN) waiting for status of probe system wide. Any new SYN packets will be dropped.`,
				},
				resource.Attribute{
					Name:        "msslearninterval",
					Description: `(Optional) Duration, in seconds, to sample the Maximum Segment Size (MSS) of the services. The Citrix ADC determines the best MSS to set for the virtual server based on this sampling. The argument to enable maximum segment size (MSS) for virtual servers must be enabled.`,
				},
				resource.Attribute{
					Name:        "msslearndelay",
					Description: `(Optional) Frequency, in seconds, at which the virtual servers learn the Maximum segment size (MSS) from the services. The argument to enable maximum segment size (MSS) for virtual servers must be enabled.`,
				},
				resource.Attribute{
					Name:        "maxtimewaitconn",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "kaprobeupdatelastactivity",
					Description: `(Optional) Update last activity for KA probes. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "maxsynackretx",
					Description: `(Optional) When 'syncookie' is disabled in the TCP profile that is bound to the virtual server or service, and the number of TCP SYN+ACK retransmission by Citrix ADC for that virtual server or service crosses this threshold, the Citrix ADC responds by using the TCP SYN-Cookie mechanism.`,
				},
				resource.Attribute{
					Name:        "synattackdetection",
					Description: `(Optional) Detect TCP SYN packet flood and send an SNMP trap. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "connflushifnomem",
					Description: `(Optional) Flush an existing connection if no memory can be obtained for new connection. HALF_CLOSED_AND_IDLE: Flush a connection that is closed by us but not by peer, or failing that, a connection that is past configured idle time. New connection fails if no such connection can be found. FIFO: If no half-closed or idle connection can be found, flush the oldest non-management connection, even if it is active. New connection fails if the oldest few connections are management connections. Note: If you enable this setting, you should also consider lowering the zombie timeout and half-close timeout, while setting the Citrix ADC timeout. See Also: connFlushThres argument below. Possible values: [ NONE , HALFCLOSED_AND_IDLE, FIFO ]`,
				},
				resource.Attribute{
					Name:        "connflushthres",
					Description: `(Optional) Flush an existing connection (as configured through -connFlushIfNoMem FIFO) if the system has more than specified number of connections, and a new connection is to be established. Note: This value may be rounded down to be a whole multiple of the number of packet engines running.`,
				},
				resource.Attribute{
					Name:        "mptcpconcloseonpassivesf",
					Description: `(Optional) Accept DATA_FIN/FAST_CLOSE on passive subflow. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "mptcpchecksum",
					Description: `(Optional) Use MPTCP DSS checksum. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "mptcpsftimeout",
					Description: `(Optional) The timeout value in seconds for idle mptcp subflows. If this timeout is not set, idle subflows are cleared after cltTimeout of vserver.`,
				},
				resource.Attribute{
					Name:        "mptcpsfreplacetimeout",
					Description: `(Optional) The minimum idle time value in seconds for idle mptcp subflows after which the sublow is replaced by new incoming subflow if maximum subflow limit is reached. The priority for replacement is given to those subflow without any transaction.`,
				},
				resource.Attribute{
					Name:        "mptcpmaxsf",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "mptcpmaxpendingsf",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "mptcppendingjointhreshold",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "mptcprtostoswitchsf",
					Description: `(Optional) Number of RTO's at subflow level, after which MPCTP should start using other subflow.`,
				},
				resource.Attribute{
					Name:        "mptcpusebackupondss",
					Description: `(Optional) When enabled, if NS receives a DSS on a backup subflow, NS will start using that subflow to send data. And if disabled, NS will continue to transmit on current chosen subflow. In case there is some error on a subflow (like RTO's/RST etc.) then NS can choose a backup subflow irrespective of this tunable. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "tcpmaxretries",
					Description: `(Optional) Number of RTO's after which a connection should be freed.`,
				},
				resource.Attribute{
					Name:        "mptcpimmediatesfcloseonfin",
					Description: `(Optional) Allow subflows to close immediately on FIN before the DATA_FIN exchange is completed at mptcp level. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "mptcpclosemptcpsessiononlastsfclose",
					Description: `(Optional) Allow to send DATA FIN or FAST CLOSE on mptcp connection while sending FIN or RST on the last subflow. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "tcpfastopencookietimeout",
					Description: `(Optional) Timeout in seconds after which a new TFO Key is computed for generating TFO Cookie. If zero, the same key is used always. If timeout is less than 120seconds, NS defaults to 120seconds timeout.`,
				},
				resource.Attribute{
					Name:        "autosyncookietimeout",
					Description: `(Optional) Timeout for the server to function in syncookie mode after the synattack. This is valid if TCP syncookie is disabled on the profile and server acts in non syncookie mode by default.`,
				},
				resource.Attribute{
					Name:        "tcpfintimeout",
					Description: `(Optional) The amount of time in seconds, after which a TCP connnection in the TCP TIME-WAIT state is flushed. ## Attribute Reference In addition to the arguments, the following attributes are available:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of the nstcpparam. It is a unique string prefixed with "tf-nstcpparam-"`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id of the nstcpparam. It is a unique string prefixed with "tf-nstcpparam-"`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "citrixadc_nstcpprofile",
			Category:         "NS",
			ShortDescription: ``,
			Description: `

The nstcpprofile resource is used to manage the TCP profile in the target ADC.


`,
			Keywords: []string{
				"ns",
				"nstcpprofile",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name for a TCP profile. The name of a TCP profile cannot be changed after it is created.`,
				},
				resource.Attribute{
					Name:        "ws",
					Description: `(Optional) Enable or disable window scaling. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "sack",
					Description: `(Optional) Enable or disable Selective ACKnowledgement (SACK). Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "wsval",
					Description: `(Optional) Factor used to calculate the new window size. This argument is needed only when window scaling is enabled.`,
				},
				resource.Attribute{
					Name:        "nagle",
					Description: `(Optional) Enable or disable the Nagle algorithm on TCP connections. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "ackonpush",
					Description: `(Optional) Send immediate positive acknowledgement (ACK) on receipt of TCP packets with PUSH flag. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "mss",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "maxburst",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "initialcwnd",
					Description: `(Optional) Initial maximum upper limit on the number of TCP packets that can be outstanding on the TCP link to the server.`,
				},
				resource.Attribute{
					Name:        "delayedack",
					Description: `(Optional) Timeout for TCP delayed ACK, in milliseconds.`,
				},
				resource.Attribute{
					Name:        "oooqsize",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "maxpktpermss",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "pktperretx",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "minrto",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "slowstartincr",
					Description: `(Optional) Multiplier that determines the rate at which slow start increases the size of the TCP transmission window after each acknowledgement of successful transmission.`,
				},
				resource.Attribute{
					Name:        "buffersize",
					Description: `(Optional) TCP buffering size, in bytes.`,
				},
				resource.Attribute{
					Name:        "syncookie",
					Description: `(Optional) Enable or disable the SYNCOOKIE mechanism for TCP handshake with clients. Disabling SYNCOOKIE prevents SYN attack protection on the Citrix ADC. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "kaprobeupdatelastactivity",
					Description: `(Optional) Update last activity for the connection after receiving keep-alive (KA) probes. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "flavor",
					Description: `(Optional) Set TCP congestion control algorithm. Possible values: [ Default, Westwood, BIC, CUBIC, Nile ]`,
				},
				resource.Attribute{
					Name:        "dynamicreceivebuffering",
					Description: `(Optional) Enable or disable dynamic receive buffering. When enabled, allows the receive buffer to be adjusted dynamically based on memory and network conditions. Note: The buffer size argument must be set for dynamic adjustments to take place. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "ka",
					Description: `(Optional) Send periodic TCP keep-alive (KA) probes to check if peer is still up. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "kaconnidletime",
					Description: `(Optional) Duration, in seconds, for the connection to be idle, before sending a keep-alive (KA) probe.`,
				},
				resource.Attribute{
					Name:        "kamaxprobes",
					Description: `(Optional) Number of keep-alive (KA) probes to be sent when not acknowledged, before assuming the peer to be down.`,
				},
				resource.Attribute{
					Name:        "kaprobeinterval",
					Description: `(Optional) Time interval, in seconds, before the next keep-alive (KA) probe, if the peer does not respond.`,
				},
				resource.Attribute{
					Name:        "sendbuffsize",
					Description: `(Optional) TCP Send Buffer Size.`,
				},
				resource.Attribute{
					Name:        "mptcp",
					Description: `(Optional) Enable or disable Multipath TCP. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "establishclientconn",
					Description: `(Optional) Establishing Client Client connection on First data/ Final-ACK / Automatic. Possible values: [ AUTOMATIC, CONN_ESTABLISHED, ON_FIRST_DATA ]`,
				},
				resource.Attribute{
					Name:        "tcpsegoffload",
					Description: `(Optional) Offload TCP segmentation to the NIC. If set to AUTOMATIC, TCP segmentation will be offloaded to the NIC, if the NIC supports it. Possible values: [ AUTOMATIC, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "rstwindowattenuate",
					Description: `(Optional) Enable or disable RST window attenuation to protect against spoofing. When enabled, will reply with corrective ACK when a sequence number is invalid. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "rstmaxack",
					Description: `(Optional) Enable or disable acceptance of RST that is out of window yet echoes highest ACK sequence number. Useful only in proxy mode. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "spoofsyndrop",
					Description: `(Optional) Enable or disable drop of invalid SYN packets to protect against spoofing. When disabled, established connections will be reset when a SYN packet is received. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "ecn",
					Description: `(Optional) Enable or disable TCP Explicit Congestion Notification. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "mptcpdropdataonpreestsf",
					Description: `(Optional) Enable or disable silently dropping the data on Pre-Established subflow. When enabled, DSS data packets are dropped silently instead of dropping the connection when data is received on pre established subflow. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "mptcpfastopen",
					Description: `(Optional) Enable or disable Multipath TCP fastopen. When enabled, DSS data packets are accepted before receiving the third ack of SYN handshake. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "mptcpsessiontimeout",
					Description: `(Optional) MPTCP session timeout in seconds. If this value is not set, idle MPTCP sessions are flushed after vserver's client idle timeout.`,
				},
				resource.Attribute{
					Name:        "timestamp",
					Description: `(Optional) Enable or Disable TCP Timestamp option (RFC 1323). Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "dsack",
					Description: `(Optional) Enable or disable DSACK. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "ackaggregation",
					Description: `(Optional) Enable or disable ACK Aggregation. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "frto",
					Description: `(Optional) Enable or disable FRTO (Forward RTO-Recovery). Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "maxcwnd",
					Description: `(Optional) TCP Maximum Congestion Window.`,
				},
				resource.Attribute{
					Name:        "fack",
					Description: `(Optional) Enable or disable FACK (Forward ACK). Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "tcpmode",
					Description: `(Optional) TCP Optimization modes TRANSPARENT / ENDPOINT. Possible values: [ TRANSPARENT, ENDPOINT ]`,
				},
				resource.Attribute{
					Name:        "tcpfastopen",
					Description: `(Optional) Enable or disable TCP Fastopen. When enabled, NS can receive or send Data in SYN or SYN-ACK packets. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "hystart",
					Description: `(Optional) Enable or disable CUBIC Hystart. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "dupackthresh",
					Description: `(Optional) TCP dupack threshold.`,
				},
				resource.Attribute{
					Name:        "burstratecontrol",
					Description: `(Optional) TCP Burst Rate Control DISABLED/FIXED/DYNAMIC. FIXED requires a TCP rate to be set. Possible values: [ DISABLED, FIXED, DYNAMIC ]`,
				},
				resource.Attribute{
					Name:        "tcprate",
					Description: `(Optional) TCP connection payload send rate in Kb/s.`,
				},
				resource.Attribute{
					Name:        "rateqmax",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "drophalfclosedconnontimeout",
					Description: `(Optional) Silently drop tcp half closed connections on idle timeout. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "dropestconnontimeout",
					Description: `(Optional) Silently drop tcp established connections on idle timeout. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "applyadaptivetcp",
					Description: `(Optional) Apply Adaptive TCP optimizations. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "tcpfastopencookiesize",
					Description: `(Optional) TCP FastOpen Cookie size. This accepts only even numbers. Odd number is trimmed down to nearest even number.`,
				},
				resource.Attribute{
					Name:        "taillossprobe",
					Description: `(Optional) TCP tail loss probe optimizations. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "clientiptcpoption",
					Description: `(Optional) Client IP in TCP options. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "clientiptcpoptionnumber",
					Description: `(Optional) ClientIP TCP Option number.`,
				},
				resource.Attribute{
					Name:        "mpcapablecbit",
					Description: `(Optional) ## Attribute Reference In addition to the arguments, the following attributes are available:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of the nstcpprofile. It has the same value as the ` + "`" + `name` + "`" + ` attribute. ## Import A nstcpprofile can be imported using its name, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import citrixadc_nstcpprofile.tf_nsprofile tf_nsprofile ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id of the nstcpprofile. It has the same value as the ` + "`" + `name` + "`" + ` attribute. ## Import A nstcpprofile can be imported using its name, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import citrixadc_nstcpprofile.tf_nsprofile tf_nsprofile ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "citrixadc_nsvpxparam",
			Category:         "NS",
			ShortDescription: ``,
			Description: `

The nsvpxparam resource is used to set ns vpx parameters.


`,
			Keywords: []string{
				"ns",
				"nsvpxparam",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "masterclockcpu1",
					Description: `(Optional) This setting applicable in virtual appliances, to move master clock source cpu from management cpu cpu0 to cpu1 ie PE0.`,
				},
				resource.Attribute{
					Name:        "cpuyield",
					Description: `(Optional) This setting applicable in virtual appliances, is to affect the cpu yield(relinquishing the cpu resources) in any hypervised environment.`,
				},
				resource.Attribute{
					Name:        "ownernode",
					Description: `(Optional) ID of the cluster node for which you are setting the cpuyield. It can be configured only through the cluster IP address. ## Attribute Reference In addition to the arguments, the following attributes are available:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of the nsvpxparam. It is a unique string prefixed with "tf-nsvpxparam-".`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id of the nsvpxparam. It is a unique string prefixed with "tf-nsvpxparam-".`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "citrixadc_password_resetter",
			Category:         "Utility",
			ShortDescription: ``,
			Description: `

This resource is used to perform the default password reset operation.


`,
			Keywords: []string{
				"utility",
				"password",
				"resetter",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "username",
					Description: `(Required) User name for the operation.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Required) The default password.`,
				},
				resource.Attribute{
					Name:        "new_password",
					Description: `(Required) The new password ## Attribute Reference In addition to the arguments, the following attributes are available:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of the password_resetter. It is a random string prefixed with "tf-password-resetter-".`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id of the password_resetter. It is a random string prefixed with "tf-password-resetter-".`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "citrixadc_pinger",
			Category:         "Utility",
			ShortDescription: ``,
			Description: `

The ping resource is used to send ping requests from the target ADC.


`,
			Keywords: []string{
				"utility",
				"pinger",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "c",
					Description: `(Optional) Number of packets to send. The default value is infinite. For Nitro API, defalut value is taken as 5.`,
				},
				resource.Attribute{
					Name:        "i",
					Description: `(Optional) Waiting time, in seconds. The default value is 1 second.`,
				},
				resource.Attribute{
					Name:        "n",
					Description: `(Optional) Numeric output only. No name resolution.`,
				},
				resource.Attribute{
					Name:        "p",
					Description: `(Optional) Pattern to fill in packets. Can be up to 16 bytes, useful for diagnosing data-dependent problems.`,
				},
				resource.Attribute{
					Name:        "q",
					Description: `(Optional) Quiet output. Only the summary is printed. For Nitro API, this flag is set by default.`,
				},
				resource.Attribute{
					Name:        "s",
					Description: `(Optional) Data size, in bytes. The default value is 56.`,
				},
				resource.Attribute{
					Name:        "t",
					Description: `(Optional) Time-out, in seconds, before ping exits.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `(Optional) Address of host to ping.`,
				},
				resource.Attribute{
					Name:        "forcenew_id_set",
					Description: `(Optional) A set of terraform resource ids. Any change in the set will trigger the execution of the pinger. ## Attribute Reference In addition to the arguments, the following attributes are available:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of the pinger. It is a random string prefixed with "tf-pinger-".`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id of the pinger. It is a random string prefixed with "tf-pinger-".`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "citrixadc_policydataset",
			Category:         "Policy",
			ShortDescription: ``,
			Description: `

The policydataset resource is used to create policy data sets.


`,
			Keywords: []string{
				"policy",
				"policydataset",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the dataset. Must not exceed 127 characters.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) Type of value to bind to the dataset. Possible values: [ ipv4, number, ipv6, ulong, double, mac ]`,
				},
				resource.Attribute{
					Name:        "indextype",
					Description: `(Optional) Index type. Possible values: [ Auto-generated, User-defined ]`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `(Optional) Any comments to preserve information about this dataset or a data bound to this dataset. ## Attribute Reference In addition to the arguments, the following attributes are available:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of the policydataset. It has the same value as the ` + "`" + `name` + "`" + ` attribute. ## Import A policydataset can be imported using its name, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import citrixadc_policydataset.tf_dataset tf_dataset ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id of the policydataset. It has the same value as the ` + "`" + `name` + "`" + ` attribute. ## Import A policydataset can be imported using its name, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import citrixadc_policydataset.tf_dataset tf_dataset ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "citrixadc_policydataset_value_binding",
			Category:         "Policy",
			ShortDescription: ``,
			Description: `

The policydataset_value_binding resource is used to add values to a policydataset.


`,
			Keywords: []string{
				"policy",
				"policydataset",
				"value",
				"binding",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "value",
					Description: `(Optional) Value of the specified type that is associated with the dataset.`,
				},
				resource.Attribute{
					Name:        "index",
					Description: `(Optional) The index of the value (ipv4, ipv6, number) associated with the set.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `(Optional) Any comments to preserve information about this dataset or a data bound to this dataset.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the dataset to which to bind the value. ## Attribute Reference In addition to the arguments, the following attributes are available:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of the policydataset_value_binding. Its value is the concatenation of the ` + "`" + `name` + "`" + ` and ` + "`" + `value` + "`" + ` attributes separated by a comma: ` + "`" + `<name>,<value>` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id of the policydataset_value_binding. Its value is the concatenation of the ` + "`" + `name` + "`" + ` and ` + "`" + `value` + "`" + ` attributes separated by a comma: ` + "`" + `<name>,<value>` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "citrixadc_policyexpression",
			Category:         "Policy",
			ShortDescription: ``,
			Description: `

The policyexpression resource is used to create policy expressions.


`,
			Keywords: []string{
				"policy",
				"policyexpression",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Unique name for the expression. Not case sensitive. Must not begin with 're' or 'xp' or be a word reserved for use as an expression qualifier prefix (such as HTTP) or enumeration value (such as ASCII). Must not be the name of an existing named expression, pattern set, dataset, stringmap, or HTTP callout.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Optional) Expression string. For example: http.req.body(100).contains("this").`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description for the expression.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `(Optional) Any comments associated with the expression. Displayed upon viewing the policy expression.`,
				},
				resource.Attribute{
					Name:        "clientsecuritymessage",
					Description: `(Optional) Message to display if the expression fails. Allowed for classic end-point check expressions only. ## Attribute Reference In addition to the arguments, the following attributes are available:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of the policyexpression. It has the same value as the ` + "`" + `name` + "`" + ` attribute. ## Import A policyexpression can be imported using its name, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import citrixadc_policyexpression.tf_policyexrpession tf_policyexrpession ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id of the policyexpression. It has the same value as the ` + "`" + `name` + "`" + ` attribute. ## Import A policyexpression can be imported using its name, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import citrixadc_policyexpression.tf_policyexrpession tf_policyexrpession ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "citrixadc_policypatset",
			Category:         "Policy",
			ShortDescription: ``,
			Description: `

The policypatset resource is used to create policy pattern sets.


`,
			Keywords: []string{
				"policy",
				"policypatset",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Unique name of the pattern set. Not case sensitive. Must begin with an ASCII letter or underscore (\_) character and must contain only alphanumeric and underscore characters. Must not be the name of an existing named expression, pattern set, dataset, string map, or HTTP callout.`,
				},
				resource.Attribute{
					Name:        "indextype",
					Description: `(Optional) Index type. Possible values: [ Auto-generated, User-defined ]`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `(Optional) Any comments to preserve information about this patset or a pattern bound to this patset. ## Attribute Reference In addition to the arguments, the following attributes are available:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of the policypatset. It has the same value as the ` + "`" + `name` + "`" + ` attribute.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id of the policypatset. It has the same value as the ` + "`" + `name` + "`" + ` attribute.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "citrixadc_policypatset_pattern_binding",
			Category:         "Policy",
			ShortDescription: ``,
			Description: `\_pattern\_biding

The policypatset\_pattern\_biding resource is used to bind patterns to a policy pattern set.


`,
			Keywords: []string{
				"policy",
				"policypatset",
				"pattern",
				"binding",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "string",
					Description: `(Required) String of characters that constitutes a pattern. For more information about the characters that can be used, refer to the character set parameter. Note: Minimum length for pattern sets used in rewrite actions of type REPLACE\_ALL, DELETE\_ALL, INSERT\_AFTER\_ALL, and INSERT\_BEFORE\_ALL, is three characters.`,
				},
				resource.Attribute{
					Name:        "feature",
					Description: `(Optional) The feature to be checked while applying this config.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the pattern set to which to bind the string.`,
				},
				resource.Attribute{
					Name:        "charset",
					Description: `(Optional) Character set associated with the characters in the string. Note: UTF-8 characters can be entered directly (if the UI supports it) or can be encoded as a sequence of hexadecimal bytes '\xNN'. For example, the UTF-8 character 'ü' can be encoded as '\xC3\xBC'.`,
				},
				resource.Attribute{
					Name:        "index",
					Description: `(Optional) The index of the string associated with the patset.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `(Optional) Any comments to preserve information about this patset or a pattern bound to this patset. ## Attribute Reference In addition to the arguments, the following attributes are available:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of the policypatset. It is the concatenation of the ` + "`" + `name` + "`" + ` and ` + "`" + `string` + "`" + ` attributes separated by a comma.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id of the policypatset. It is the concatenation of the ` + "`" + `name` + "`" + ` and ` + "`" + `string` + "`" + ` attributes separated by a comma.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "citrixadc_policystringmap",
			Category:         "Policy",
			ShortDescription: ``,
			Description: `

The policystringmap resource is used to create policy string maps.


`,
			Keywords: []string{
				"policy",
				"policystringmap",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Unique name for the string map. Not case sensitive. Must begin with an ASCII letter or underscore (\_) character, and must consist only of ASCII alphanumeric or underscore characters. Must not begin with 're' or 'xp' or be a word reserved for use as an expression qualifier prefix (such as HTTP) or enumeration value (such as ASCII). Must not be the name of an existing named expression, pattern set, dataset, string map, or HTTP callout.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `(Optional) Comments associated with the string map or key-value pair bound to this string map. ## Attribute Reference In addition to the arguments, the following attributes are available:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of the policystringmap. It has the same value as the ` + "`" + `name` + "`" + ` attribute.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id of the policystringmap. It has the same value as the ` + "`" + `name` + "`" + ` attribute.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "citrixadc_policystringmap_pattern_binding",
			Category:         "Policy",
			ShortDescription: ``,
			Description: `\_pattern\_binding

The policystringmap\_pattern\_binding resource is used to bind patters to a stringmap.


`,
			Keywords: []string{
				"policy",
				"policystringmap",
				"pattern",
				"binding",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the string map to which to bind the key-value pair.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) Character string constituting the value associated with the key. This value is returned when processed data matches the associated key. Refer to the key parameter for details of the value character set.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) Character string constituting the key to be bound to the string map. The key is matched against the data processed by the operation that uses the string map. The default character set is ASCII. UTF-8 characters can be included if the character set is UTF-8. UTF-8 characters can be entered directly (if the UI supports it) or can be encoded as a sequence of hexadecimal bytes '\xNN'. For example, the UTF-8 character 'ü' can be encoded as '\xC3\xBC'. ## Attribute Reference In addition to the arguments, the following attributes are available:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of the policystringmap. It is the concatenation of the ` + "`" + `name` + "`" + ` and ` + "`" + `key` + "`" + ` attributes separated by a comma.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id of the policystringmap. It is the concatenation of the ` + "`" + `name` + "`" + ` and ` + "`" + `key` + "`" + ` attributes separated by a comma.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "citrixadc_quicbridgeprofile",
			Category:         "QuickBridge",
			ShortDescription: ``,
			Description: `

The ` + "`" + `quicbridgeprofile` + "`" + ` resource is used to create Citrix ADC Quick Bridge Profiles.

`,
			Keywords: []string{
				"quickbridge",
				"quicbridgeprofile",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `Name for the QUIC profile. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals sign (=), and hyphen (-) characters. Cannot be changed after the profile is created.`,
				},
				resource.Attribute{
					Name:        "routingalgorithm",
					Description: `(Optional) Routing algorithm to generate routable connection IDs.`,
				},
				resource.Attribute{
					Name:        "serveridlength",
					Description: `(Optional) Length of serverid to encode/decode server information. ## Attribute Reference In addition to the arguments, the following attributes are available:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of the quicbridgeprofile. It has the same value as the ` + "`" + `name` + "`" + ` attribute.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id of the quicbridgeprofile. It has the same value as the ` + "`" + `name` + "`" + ` attribute.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "citrixadc_rebooter",
			Category:         "Utility",
			ShortDescription: ``,
			Description: `

The rebooter resource is used to reboot the target ADC.


`,
			Keywords: []string{
				"utility",
				"rebooter",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "warm",
					Description: `(Required) Restarts the Citrix ADC software without rebooting the underlying operating system. The session terminates and you must log on to the appliance after it has restarted. Note: This argument is required only for nCore appliances. Classic appliances ignore this argument.`,
				},
				resource.Attribute{
					Name:        "timestamp",
					Description: `(Required) Time when reboot happened. Used to force the operation again in case all other attributes remain the same.`,
				},
				resource.Attribute{
					Name:        "wait_until_reachable",
					Description: `(Required) Boolean flag. If set to true the resource will wait for the ADC to become reachable.`,
				},
				resource.Attribute{
					Name:        "reachable_timeout",
					Description: `(Optional) Time period to wait for the ADC to become reachable. Defaults to "10m".`,
				},
				resource.Attribute{
					Name:        "reachable_poll_delay",
					Description: `(Optional) Time period to wait before first poll after reboot. Defaults to "60s".`,
				},
				resource.Attribute{
					Name:        "reachable_poll_interval",
					Description: `(Optional) Time interval between polls. Defaults to "60s".`,
				},
				resource.Attribute{
					Name:        "reachable_poll_timeout",
					Description: `(Optional) Time period to wait for each poll request. Defaults to "20s". ## Attribute Reference In addition to the arguments, the following attributes are available:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of the rebooter. It is a unique string prefixed with "tf-rebooter-"`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id of the rebooter. It is a unique string prefixed with "tf-rebooter-"`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "citrixadc_responderaction",
			Category:         "Responder",
			ShortDescription: ``,
			Description: `

The responderaction resource is used to create responder actions.


`,
			Keywords: []string{
				"responder",
				"responderaction",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name for the responder action.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) Type of responder action. Available settings function as follows:`,
				},
				resource.Attribute{
					Name:        "target",
					Description: `(Optional) Expression specifying what to respond with. Typically a URL for redirect policies or a default-syntax expression. In addition to Citrix ADC default-syntax expressions that refer to information in the request, a stringbuilder expression can contain text and HTML, and simple escape codes that define new lines and paragraphs. Enclose each stringbuilder expression element (either a Citrix ADC default-syntax expression or a string) in double quotation marks. Use the plus (+) character to join the elements. Examples: 1) Respondwith expression that sends an HTTP 1.1 200 OK response: "HTTP/1.1 200 OK\r\n\r\n" 2) Redirect expression that redirects user to the specified web host and appends the request URL to the redirect. "http://backupsite2.com" + HTTP.REQ.URL 3) Respondwith expression that sends an HTTP 1.1 404 Not Found response with the request URL included in the response: "HTTP/1.1 404 Not Found\r\n\r\n"+ "HTTP.REQ.URL.HTTP_URL_SAFE" + "does not exist on the web server." The following requirement applies only to the Citrix ADC CLI: Enclose the entire expression in single quotation marks. (Citrix ADC expression elements should be included inside the single quotation marks for the entire expression, but do not need to be enclosed in double quotation marks.).`,
				},
				resource.Attribute{
					Name:        "htmlpage",
					Description: `(Optional) For respondwithhtmlpage policies, name of the HTML page object to use as the response. You must first import the page object.`,
				},
				resource.Attribute{
					Name:        "bypasssafetycheck",
					Description: `(Optional) Bypass the safety check, allowing potentially unsafe expressions. An unsafe expression in a response is one that contains references to request elements that might not be present in all requests. If a response refers to a missing request element, an empty string is used instead. Possible values: [ YES, NO ]`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `(Optional) Comment. Any type of information about this responder action.`,
				},
				resource.Attribute{
					Name:        "responsestatuscode",
					Description: `(Optional) HTTP response status code, for example 200, 302, 404, etc. The default value for the redirect action type is 302 and for respondwithhtmlpage is 200.`,
				},
				resource.Attribute{
					Name:        "reasonphrase",
					Description: `(Optional) Expression specifying the reason phrase of the HTTP response. The reason phrase may be a string literal with quotes or a PI expression. For example: "Invalid URL: " + HTTP.REQ.URL. ## Attribute Reference In addition to the arguments, the following attributes are available:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of the responderaction. It has the same value as the ` + "`" + `name` + "`" + ` attribute. ## Import A responderaction can be imported using its name, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import citrixadc_responderaction.tf_responderaction tf_responderaction ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id of the responderaction. It has the same value as the ` + "`" + `name` + "`" + ` attribute. ## Import A responderaction can be imported using its name, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import citrixadc_responderaction.tf_responderaction tf_responderaction ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "citrixadc_responderpolicy",
			Category:         "Responder",
			ShortDescription: ``,
			Description: `

The responderpolicy resource is used to create responder policies.


`,
			Keywords: []string{
				"responder",
				"responderpolicy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name for the responder policy.`,
				},
				resource.Attribute{
					Name:        "rule",
					Description: `(Optional) Expression that the policy uses to determine whether to respond to the specified request.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Optional) Name of the responder action to perform if the request matches this responder policy. There are also some built-in actions which can be used. These are:`,
				},
				resource.Attribute{
					Name:        "undefaction",
					Description: `(Optional) Action to perform if the result of policy evaluation is undefined (UNDEF). An UNDEF event indicates an internal error condition. Only the above built-in actions can be used.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `(Optional) Any type of information about this responder policy.`,
				},
				resource.Attribute{
					Name:        "logaction",
					Description: `(Optional) Name of the messagelog action to use for requests that match this policy.`,
				},
				resource.Attribute{
					Name:        "appflowaction",
					Description: `(Optional) AppFlow action to invoke for requests that match this policy.`,
				},
				resource.Attribute{
					Name:        "globalbinding",
					Description: `(Optional) A global binding block. Documented below.`,
				},
				resource.Attribute{
					Name:        "lbvserverbinding",
					Description: `(Optional) A lbvserver binding block. Documented below.`,
				},
				resource.Attribute{
					Name:        "csvserverbinding",
					Description: `(Optional) A csvserver binding block. Documented below. A global binding block supports the following:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) Specifies the bind point whose policies you want to display. Available settings function as follows:`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `(Optional) Specifies the priority of the policy.`,
				},
				resource.Attribute{
					Name:        "gotopriorityexpression",
					Description: `(Optional) Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.`,
				},
				resource.Attribute{
					Name:        "invoke",
					Description: `(Optional) If the current policy evaluates to TRUE, terminate evaluation of policies bound to the current policy label, and then forward the request to the specified virtual server or evaluate the specified policy label.`,
				},
				resource.Attribute{
					Name:        "labeltype",
					Description: `(Optional) Type of invocation, Available settings function as follows:`,
				},
				resource.Attribute{
					Name:        "labelname",
					Description: `(Optional) Name of the policy label to invoke. If the current policy evaluates to TRUE, the invoke parameter is set, and Label Type is policylabel. A lbvserver binding block supports the following:`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `(Optional) Priority.`,
				},
				resource.Attribute{
					Name:        "gotopriorityexpression",
					Description: `(Optional) Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.`,
				},
				resource.Attribute{
					Name:        "invoke",
					Description: `(Optional) Invoke policies bound to a virtual server or policy label.`,
				},
				resource.Attribute{
					Name:        "labeltype",
					Description: `(Optional) The invocation type. Possible values: [ reqvserver, resvserver, policylabel ]`,
				},
				resource.Attribute{
					Name:        "labelname",
					Description: `(Optional) Name of the label invoked.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name for the virtual server.`,
				},
				resource.Attribute{
					Name:        "bindpoint",
					Description: `(Optional) Bind point to which to bind the policy. Applicable only to compression, rewrite, videooptimization and cache policies. Possible values: [ REQUEST, RESPONSE ]`,
				},
				resource.Attribute{
					Name:        "policyname",
					Description: `(Optional) Name of the policy. A csvserver binding block supports the following:`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `(Optional) Priority for the policy.`,
				},
				resource.Attribute{
					Name:        "gotopriorityexpression",
					Description: `(Optional) Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.`,
				},
				resource.Attribute{
					Name:        "invoke",
					Description: `(Optional) Invoke flag.`,
				},
				resource.Attribute{
					Name:        "labeltype",
					Description: `(Optional) The invocation type. Possible values: [ reqvserver, resvserver, policylabel ]`,
				},
				resource.Attribute{
					Name:        "labelname",
					Description: `(Optional) Name of the label invoked.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the content switching virtual server to which the content switching policy applies.`,
				},
				resource.Attribute{
					Name:        "targetlbvserver",
					Description: `(Optional) Name of the Load Balancing virtual server to which the content is switched, if policy rule is evaluated to be TRUE. Example: bind cs vs cs1 -policyname pol1 -priority 101 -targetLBVserver lb1 Note: Use this parameter only in case of Content Switching policy bind operations to a CS vserver.`,
				},
				resource.Attribute{
					Name:        "bindpoint",
					Description: `(Optional) For a rewrite policy, the bind point to which to bind the policy. Note: This parameter applies only to rewrite policies, because content switching policies are evaluated only at request time. Possible values: [ REQUEST, RESPONSE, ICA_REQUEST, OTHERTCP_REQUEST ]`,
				},
				resource.Attribute{
					Name:        "policyname",
					Description: `(Optional) Name of the policy. ## Attribute Reference In addition to the arguments, the following attributes are available:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of the responderpolicy. It has the same value as the ` + "`" + `name` + "`" + ` attribute. ## Import A responderpolicy can be imported using its name, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import citrixadc_responderpolicy.tf_responder_policy tf_responder_policy ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id of the responderpolicy. It has the same value as the ` + "`" + `name` + "`" + ` attribute. ## Import A responderpolicy can be imported using its name, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import citrixadc_responderpolicy.tf_responder_policy tf_responder_policy ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "citrixadc_responderpolicylabel",
			Category:         "Responder",
			ShortDescription: ``,
			Description: `

The responderpolicylabel resource is used to create responder policy labels.


`,
			Keywords: []string{
				"responder",
				"responderpolicylabel",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "labelname",
					Description: `(Optional) Name for the responder policy label.`,
				},
				resource.Attribute{
					Name:        "policylabeltype",
					Description: `(Optional) Type of responses sent by the policies bound to this policy label. Types are:`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `(Optional) Any comments to preserve information about this responder policy label. ## Attribute Reference In addition to the arguments, the following attributes are available:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of the responderpolicylabel. It has the same value as the ` + "`" + `labelname` + "`" + ` attribute. ## Import A responderpolicylabel can be imported using its labelname, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import citrixadc_responderpolicylabel.tf_responder_policylabel tf_responder_policylabel ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id of the responderpolicylabel. It has the same value as the ` + "`" + `labelname` + "`" + ` attribute. ## Import A responderpolicylabel can be imported using its labelname, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import citrixadc_responderpolicylabel.tf_responder_policylabel tf_responder_policylabel ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "citrixadc_rewriteaction",
			Category:         "Rewrite",
			ShortDescription: ``,
			Description: `

The rewriteaction resource is used to create rewrite actions.


`,
			Keywords: []string{
				"rewrite",
				"rewriteaction",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name for the user-defined rewrite action.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) Type of user-defined rewrite action. The information that you provide for, and the effect of, each type are as follows::`,
				},
				resource.Attribute{
					Name:        "target",
					Description: `(Optional) Expression that specifies which part of the request or response to rewrite.`,
				},
				resource.Attribute{
					Name:        "stringbuilderexpr",
					Description: `(Optional) Expression that specifies the content to insert into the request or response at the specified location, or that replaces the specified string.`,
				},
				resource.Attribute{
					Name:        "pattern",
					Description: `(Optional) DEPRECATED in favor of -search: Pattern that is used to match multiple strings in the request or response. The pattern may be a string literal (without quotes) or a PCRE-format regular expression with a delimiter that consists of any printable ASCII non-alphanumeric character except for the underscore (\_) and space ( ) that is not otherwise used in the expression. Example: re~https?://|HTTPS?://~ The preceding regular expression can use the tilde (~) as the delimiter because that character does not appear in the regular expression itself. Used in the INSERT_BEFORE_ALL, INSERT_AFTER_ALL, REPLACE_ALL, and DELETE_ALL action types.`,
				},
				resource.Attribute{
					Name:        "search",
					Description: `(Optional) Search facility that is used to match multiple strings in the request or response. Used in the INSERT_BEFORE_ALL, INSERT_AFTER_ALL, REPLACE_ALL, and DELETE_ALL action types. The following search types are supported:`,
				},
				resource.Attribute{
					Name:        "bypasssafetycheck",
					Description: `(Optional) Bypass the safety check and allow unsafe expressions. An unsafe expression is one that contains references to message elements that might not be present in all messages. If an expression refers to a missing request element, an empty string is used instead. Possible values: [ YES, NO ]`,
				},
				resource.Attribute{
					Name:        "refinesearch",
					Description: `(Optional) Specify additional criteria to refine the results of the search. Always starts with the "extend(m,n)" operation, where 'm' specifies number of bytes to the left of selected data and 'n' specifies number of bytes to the right of selected data to extend the selected area. You can use refineSearch only on body expressions, and for the INSERT_BEFORE_ALL, INSERT_AFTER_ALL, REPLACE_ALL, and DELETE_ALL action types. Example: -refineSearch 'EXTEND(10, 20).REGEX_SELECT(re~0x[0-9a-zA-Z]+~).`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `(Optional) Comment. Can be used to preserve information about this rewrite action. ## Attribute Reference In addition to the arguments, the following attributes are available:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of the rewriteaction. It has the same value as the ` + "`" + `name` + "`" + ` attribute. ## Import A rewriteaction can be imported using its name, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import citrixadc_rewriteaction.tf_rewrite_action tf_rewrite_action ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id of the rewriteaction. It has the same value as the ` + "`" + `name` + "`" + ` attribute. ## Import A rewriteaction can be imported using its name, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import citrixadc_rewriteaction.tf_rewrite_action tf_rewrite_action ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "citrixadc_rewritepolicy",
			Category:         "Rewrite",
			ShortDescription: ``,
			Description: `

The rewritepolicy resource is used to create rewrite policies.


`,
			Keywords: []string{
				"rewrite",
				"rewritepolicy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name for the rewrite policy.`,
				},
				resource.Attribute{
					Name:        "rule",
					Description: `(Optional) Expression against which traffic is evaluated. The following requirements apply only to the Citrix ADC CLI:`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Optional) Name of the rewrite action to perform if the request or response matches this rewrite policy. There are also some built-in actions which can be used. These are:`,
				},
				resource.Attribute{
					Name:        "undefaction",
					Description: `(Optional) Action to perform if the result of policy evaluation is undefined (UNDEF). An UNDEF event indicates an internal error condition. Only the above built-in actions can be used.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `(Optional) Any comments to preserve information about this rewrite policy.`,
				},
				resource.Attribute{
					Name:        "logaction",
					Description: `(Optional) Name of messagelog action to use when a request matches this policy.`,
				},
				resource.Attribute{
					Name:        "globalbinding",
					Description: `(Optional) A global binding block, documented below.`,
				},
				resource.Attribute{
					Name:        "lbvserverbinding",
					Description: `(Optional) A lbvserver binding block, documented below.`,
				},
				resource.Attribute{
					Name:        "csvserverbinding",
					Description: `(Optional) A csvserver binding block, documented below. A global binding supports the following:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) The bindpoint to which to policy is bound. Possible values: [ REQ_OVERRIDE, REQ_DEFAULT, RES_OVERRIDE, RES_DEFAULT, OTHERTCP_REQ_OVERRIDE, OTHERTCP_REQ_DEFAULT, OTHERTCP_RES_OVERRIDE, OTHERTCP_RES_DEFAULT, SIPUDP_REQ_OVERRIDE, SIPUDP_REQ_DEFAULT, SIPUDP_RES_OVERRIDE, SIPUDP_RES_DEFAULT, SIPTCP_REQ_OVERRIDE, SIPTCP_REQ_DEFAULT, SIPTCP_RES_OVERRIDE, SIPTCP_RES_DEFAULT, DIAMETER_REQ_OVERRIDE, DIAMETER_REQ_DEFAULT, DIAMETER_RES_OVERRIDE, DIAMETER_RES_DEFAULT, RADIUS_REQ_OVERRIDE, RADIUS_REQ_DEFAULT, RADIUS_RES_OVERRIDE, RADIUS_RES_DEFAULT, DNS_REQ_OVERRIDE, DNS_REQ_DEFAULT, DNS_RES_OVERRIDE, DNS_RES_DEFAULT ]`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `(Optional) Specifies the priority of the policy.`,
				},
				resource.Attribute{
					Name:        "gotopriorityexpression",
					Description: `(Optional) Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.`,
				},
				resource.Attribute{
					Name:        "invoke",
					Description: `(Optional) Terminate evaluation of policies bound to the current policy label, and then forward the request to the specified virtual server or evaluate the specified policy label.`,
				},
				resource.Attribute{
					Name:        "labeltype",
					Description: `(Optional) Type of invocation. Available settings function as follows:`,
				},
				resource.Attribute{
					Name:        "labelname",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "globalbindtype",
					Description: `(Optional) . Possible values: [ SYSTEM_GLOBAL, VPN_GLOBAL, RNAT_GLOBAL ]`,
				},
				resource.Attribute{
					Name:        "policyname",
					Description: `(Optional) Policy name. A lbvserver binding supports the following:`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `(Optional) Priority.`,
				},
				resource.Attribute{
					Name:        "gotopriorityexpression",
					Description: `(Optional) Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.`,
				},
				resource.Attribute{
					Name:        "bindpoint",
					Description: `(Optional) The bindpoint to which the policy is bound. Possible values: [ REQUEST, RESPONSE ]`,
				},
				resource.Attribute{
					Name:        "invoke",
					Description: `(Optional) Invoke policies bound to a virtual server or policy label.`,
				},
				resource.Attribute{
					Name:        "labeltype",
					Description: `(Optional) The invocation type. Possible values: [ reqvserver, resvserver, policylabel ]`,
				},
				resource.Attribute{
					Name:        "labelname",
					Description: `(Optional) Name of the label invoked.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name for the virtual server. A csvserver binding supports the following:`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `(Optional) Priority for the policy.`,
				},
				resource.Attribute{
					Name:        "gotopriorityexpression",
					Description: `(Optional) Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.`,
				},
				resource.Attribute{
					Name:        "bindpoint",
					Description: `(Optional) The bindpoint to which the policy is bound. Possible values: [ REQUEST, RESPONSE, ICA_REQUEST, OTHERTCP_REQUEST ]`,
				},
				resource.Attribute{
					Name:        "invoke",
					Description: `(Optional) Invoke flag.`,
				},
				resource.Attribute{
					Name:        "labeltype",
					Description: `(Optional) The invocation type. Possible values: [ reqvserver, resvserver, policylabel ]`,
				},
				resource.Attribute{
					Name:        "labelname",
					Description: `(Optional) Name of the label invoked.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the content switching virtual server to which the content switching policy applies.`,
				},
				resource.Attribute{
					Name:        "targetlbvserver",
					Description: `(Optional) Name of the Load Balancing virtual server to which the content is switched, if policy rule is evaluated to be TRUE. Example: bind cs vs cs1 -policyname pol1 -priority 101 -targetLBVserver lb1 Note: Use this parameter only in case of Content Switching policy bind operations to a CS vserver. ## Attribute Reference In addition to the arguments, the following attributes are available:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of the rewritepolicy. It has the same value as the ` + "`" + `name` + "`" + ` attribute. ## Import A rewritepolicy can be imported using its name, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import citrixadc_rewritepolicy.tf_rewritepolicy tf_rewritepolicy ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id of the rewritepolicy. It has the same value as the ` + "`" + `name` + "`" + ` attribute. ## Import A rewritepolicy can be imported using its name, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import citrixadc_rewritepolicy.tf_rewritepolicy tf_rewritepolicy ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "citrixadc_rewritepolicylabel",
			Category:         "Rewrite",
			ShortDescription: ``,
			Description: `

The rewritepolicylabel resource is used to create rewrite policy labels.


`,
			Keywords: []string{
				"rewrite",
				"rewritepolicylabel",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "labelname",
					Description: `(Optional) Name for the rewrite policy label.`,
				},
				resource.Attribute{
					Name:        "transform",
					Description: `(Optional) Types of transformations allowed by the policies bound to the label. For Rewrite, the following types are supported:`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `(Optional) Any comments to preserve information about this rewrite policy label. ## Attribute Reference In addition to the arguments, the following attributes are available:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of the rewritepolicylabel. It has the same value as the ` + "`" + `labelname` + "`" + ` attribute. ## Import A rewritepolicylabel can be imported using its labelname, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import citrixadc_rewritepolicylabel.tf_rewritepolicylabel tf_rewritepolicylabel ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id of the rewritepolicylabel. It has the same value as the ` + "`" + `labelname` + "`" + ` attribute. ## Import A rewritepolicylabel can be imported using its labelname, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import citrixadc_rewritepolicylabel.tf_rewritepolicylabel tf_rewritepolicylabel ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "citrixadc_rnat",
			Category:         "Network",
			ShortDescription: ``,
			Description: `

The rnat resource is used to create reverse nat rules.


`,
			Keywords: []string{
				"network",
				"rnat",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "rnatsname",
					Description: `(Optional) the name for the rnat rules.`,
				},
				resource.Attribute{
					Name:        "rnat",
					Description: `(Optional) blocks of rnat rules. Documented below. A rnat block supports the following:`,
				},
				resource.Attribute{
					Name:        "network",
					Description: `(Optional) The network address defined for the RNAT entry.`,
				},
				resource.Attribute{
					Name:        "netmask",
					Description: `(Optional) The subnet mask for the network address.`,
				},
				resource.Attribute{
					Name:        "aclname",
					Description: `(Optional) An extended ACL defined for the RNAT entry.`,
				},
				resource.Attribute{
					Name:        "td",
					Description: `(Optional) Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.`,
				},
				resource.Attribute{
					Name:        "redirectport",
					Description: `(Optional) Port number to which the IPv4 packets are redirected. Applicable to TCP and UDP protocols.`,
				},
				resource.Attribute{
					Name:        "natip",
					Description: `(Optional) Any NetScaler-owned IPv4 address except the NSIP address. The NetScaler appliance replaces the source IP addresses of server-generated packets with the IP address specified. The IP address must be a public NetScaler-owned IP address. If you specify multiple addresses for this field, NATIP selection uses the round robin algorithm for each session. By specifying a range of IP addresses, you can specify all NetScaler-owned IP addresses, except the NSIP, that fall within the specified range.`,
				},
				resource.Attribute{
					Name:        "natip2",
					Description: `(Optional) ## Attribute Reference In addition to the arguments, the following attributes are available:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of the rnat rules. It has the same value as the ` + "`" + `rnatsname` + "`" + ` attribute.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id of the rnat rules. It has the same value as the ` + "`" + `rnatsname` + "`" + ` attribute.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "citrixadc_route",
			Category:         "Network",
			ShortDescription: ``,
			Description: `

The route resource is used to create routing rules.


`,
			Keywords: []string{
				"network",
				"route",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "network",
					Description: `(Optional) IPv4 network address for which to add a route entry in the routing table of the Citrix ADC.`,
				},
				resource.Attribute{
					Name:        "netmask",
					Description: `(Optional) The subnet mask associated with the network address.`,
				},
				resource.Attribute{
					Name:        "gateway",
					Description: `(Optional) IP address of the gateway for this route. Can be either the IP address of the gateway, or can be null to specify a null interface route.`,
				},
				resource.Attribute{
					Name:        "vlan",
					Description: `(Optional) VLAN as the gateway for this route.`,
				},
				resource.Attribute{
					Name:        "cost",
					Description: `(Optional) Positive integer used by the routing algorithms to determine preference for using this route. The lower the cost, the higher the preference.`,
				},
				resource.Attribute{
					Name:        "td",
					Description: `(Optional) Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.`,
				},
				resource.Attribute{
					Name:        "distance",
					Description: `(Optional) Administrative distance of this route, which determines the preference of this route over other routes, with same destination, from different routing protocols. A lower value is preferred.`,
				},
				resource.Attribute{
					Name:        "cost1",
					Description: `(Optional) The cost of a route is used to compare routes of the same type. The route having the lowest cost is the most preferred route. Possible values: 0 through 65535. Default: 0.`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `(Optional) Positive integer used by the routing algorithms to determine preference for this route over others of equal cost. The lower the weight, the higher the preference.`,
				},
				resource.Attribute{
					Name:        "advertise",
					Description: `(Optional) Advertise this route. Possible values: [ DISABLED, ENABLED ]`,
				},
				resource.Attribute{
					Name:        "msr",
					Description: `(Optional) Monitor this route using a monitor of type ARP or PING. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "monitor",
					Description: `(Optional) Name of the monitor, of type ARP or PING, configured on the Citrix ADC to monitor this route.`,
				},
				resource.Attribute{
					Name:        "ownergroup",
					Description: `(Optional) The owner node group in a Cluster for this route. If owner node group is not specified then the route is treated as Striped route.`,
				},
				resource.Attribute{
					Name:        "routetype",
					Description: `(Optional) Protocol used by routes that you want to remove from the routing table of the Citrix ADC. Possible values: [ CONNECTED, STATIC, DYNAMIC, OSPF, ISIS, RIP, BGP ]`,
				},
				resource.Attribute{
					Name:        "detail",
					Description: `(Optional) Display a detailed view. ## Attribute Reference In addition to the arguments, the following attributes are available:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of the route. It is the conatenation of the ` + "`" + `network` + "`" + `, ` + "`" + `netmask` + "`" + ` and ` + "`" + `gateway` + "`" + ` attributes. ## Import A route can be imported using its id value, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import citrixadc_route.tf_route 100.0.100.0__255.255.255.0__100.0.1.1 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id of the route. It is the conatenation of the ` + "`" + `network` + "`" + `, ` + "`" + `netmask` + "`" + ` and ` + "`" + `gateway` + "`" + ` attributes. ## Import A route can be imported using its id value, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import citrixadc_route.tf_route 100.0.100.0__255.255.255.0__100.0.1.1 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "citrixadc_routerdynamicrouting",
			Category:         "Network",
			ShortDescription: ``,
			Description: `

The routerdynamcirouting resource is used to create dynamic routing entries.


`,
			Keywords: []string{
				"network",
				"routerdynamicrouting",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "commandlines",
					Description: `(Optional) A list of commands to be executed.`,
				},
				resource.Attribute{
					Name:        "nodeid",
					Description: `(Optional) Unique number that identifies the cluster node. ## Attribute Reference In addition to the arguments, the following attributes are available:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of the routerdynamcirouting. It is a random string prefixed with "tf-routerdynamicrouting-"`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id of the routerdynamcirouting. It is a random string prefixed with "tf-routerdynamicrouting-"`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "citrixadc_server",
			Category:         "Basic",
			ShortDescription: ``,
			Description: `

The server resource is used to create servers.


`,
			Keywords: []string{
				"basic",
				"server",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name for the server.`,
				},
				resource.Attribute{
					Name:        "ipaddress",
					Description: `(Optional) IPv4 or IPv6 address of the server. If you create an IP address based server, you can specify the name of the server, instead of its IP address, when creating a service. Note: If you do not create a server entry, the server IP address that you enter when you create a service becomes the name of the server.`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `(Optional) Domain name of the server. For a domain based configuration, you must create the server first.`,
				},
				resource.Attribute{
					Name:        "translationip",
					Description: `(Optional) IP address used to transform the server's DNS-resolved IP address.`,
				},
				resource.Attribute{
					Name:        "translationmask",
					Description: `(Optional) The netmask of the translation ip.`,
				},
				resource.Attribute{
					Name:        "domainresolveretry",
					Description: `(Optional) Time, in seconds, for which the Citrix ADC must wait, after DNS resolution fails, before sending the next DNS query to resolve the domain name.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `(Optional) Initial state of the server. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "ipv6address",
					Description: `(Optional) Support IPv6 addressing mode. If you configure a server with the IPv6 addressing mode, you cannot use the server in the IPv4 addressing mode. Possible values: [ YES, NO ]`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `(Optional) Any information about the server.`,
				},
				resource.Attribute{
					Name:        "td",
					Description: `(Optional) Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.`,
				},
				resource.Attribute{
					Name:        "querytype",
					Description: `(Optional) Specify the type of DNS resolution to be done on the configured domain to get the backend services. Valid query types are A, AAAA and SRV with A being the default querytype. The type of DNS resolution done on the domains in SRV records is inherited from ipv6 argument. Possible values: [ A, AAAA, SRV ]`,
				},
				resource.Attribute{
					Name:        "domainresolvenow",
					Description: `(Optional) Immediately send a DNS query to resolve the server's domain name.`,
				},
				resource.Attribute{
					Name:        "delay",
					Description: `(Optional) Time, in seconds, after which all the services configured on the server are disabled.`,
				},
				resource.Attribute{
					Name:        "graceful",
					Description: `(Optional) Shut down gracefully, without accepting any new connections, and disabling each service when all of its connections are closed. Possible values: [ YES, NO ]`,
				},
				resource.Attribute{
					Name:        "internal",
					Description: `(Optional) Display names of the servers that have been created for internal use. ## Attribute Reference In addition to the arguments, the following attributes are available:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of the server. It has the same value as the ` + "`" + `name` + "`" + ` attribute. ## Import A server can be imported using its name, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import citrixadc_server.tf_server tf_server ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id of the server. It has the same value as the ` + "`" + `name` + "`" + ` attribute. ## Import A server can be imported using its name, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import citrixadc_server.tf_server tf_server ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "citrixadc_service",
			Category:         "Basic",
			ShortDescription: ``,
			Description: `

The service resource is used to create services.


`,
			Keywords: []string{
				"basic",
				"service",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name for the service.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `(Optional) IP to assign to the service.`,
				},
				resource.Attribute{
					Name:        "servername",
					Description: `(Optional) Name of the server that hosts the service.`,
				},
				resource.Attribute{
					Name:        "servicetype",
					Description: `(Optional) Protocol in which data is exchanged with the service. Possible values: [ HTTP, FTP, TCP, UDP, SSL, SSL_BRIDGE, SSL_TCP, DTLS, NNTP, RPCSVR, DNS, ADNS, SNMP, RTSP, DHCPRA, ANY, SIP_UDP, SIP_TCP, SIP_SSL, DNS_TCP, ADNS_TCP, MYSQL, MSSQL, ORACLE, MONGO, MONGO_TLS, RADIUS, RADIUSListener, RDP, DIAMETER, SSL_DIAMETER, TFTP, SMPP, PPTP, GRE, SYSLOGTCP, SYSLOGUDP, FIX, SSL_FIX, USER_TCP, USER_SSL_TCP, QUIC, IPFIX, LOGSTREAM, LOGSTREAM_SSL ]`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional) Port number of the service. Range 1 - 65535`,
				},
				resource.Attribute{
					Name:        "cleartextport",
					Description: `(Optional) Port to which clear text data must be sent after the appliance decrypts incoming SSL traffic. Applicable to transparent SSL services.`,
				},
				resource.Attribute{
					Name:        "cachetype",
					Description: `(Optional) Cache type supported by the cache server. Possible values: [ TRANSPARENT, REVERSE, FORWARD ]`,
				},
				resource.Attribute{
					Name:        "maxclient",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "healthmonitor",
					Description: `(Optional) Monitor the health of this service. Available settings function as follows: YES - Send probes to check the health of the service. NO - Do not send probes to check the health of the service. With the NO option, the appliance shows the service as UP at all times. Possible values: [ YES, NO ]`,
				},
				resource.Attribute{
					Name:        "maxreq",
					Description: `(Optional) Note: Connection requests beyond this value are rejected.`,
				},
				resource.Attribute{
					Name:        "cacheable",
					Description: `(Optional) Use the transparent cache redirection virtual server to forward requests to the cache server. Note: Do not specify this parameter if you set the Cache Type parameter. Possible values: [ YES, NO ]`,
				},
				resource.Attribute{
					Name:        "cip",
					Description: `(Optional) Before forwarding a request to the service, insert an HTTP header with the client's IPv4 or IPv6 address as its value. Used if the server needs the client's IP address for security, accounting, or other purposes, and setting the Use Source IP parameter is not a viable option. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "cipheader",
					Description: `(Optional) Name for the HTTP header whose value must be set to the IP address of the client. Used with the Client IP parameter. If you set the Client IP parameter, and you do not specify a name for the header, the appliance uses the header name specified for the global Client IP Header parameter (the cipHeader parameter in the set ns param CLI command or the Client IP Header parameter in the Configure HTTP Parameters dialog box at System > Settings > Change HTTP parameters). If the global Client IP Header parameter is not specified, the appliance inserts a header with the name "client-ip.".`,
				},
				resource.Attribute{
					Name:        "usip",
					Description: `(Optional) Use the client's IP address as the source IP address when initiating a connection to the server. When creating a service, if you do not set this parameter, the service inherits the global Use Source IP setting (available in the enable ns mode and disable ns mode CLI commands, or in the System > Settings > Configure modes > Configure Modes dialog box). However, you can override this setting after you create the service. Possible values: [ YES, NO ]`,
				},
				resource.Attribute{
					Name:        "pathmonitor",
					Description: `(Optional) Path monitoring for clustering. Possible values: [ YES, NO ]`,
				},
				resource.Attribute{
					Name:        "pathmonitorindv",
					Description: `(Optional) Individual Path monitoring decisions. Possible values: [ YES, NO ]`,
				},
				resource.Attribute{
					Name:        "useproxyport",
					Description: `(Optional) Use the proxy port as the source port when initiating connections with the server. With the NO setting, the client-side connection port is used as the source port for the server-side connection. Note: This parameter is available only when the Use Source IP (USIP) parameter is set to YES. Possible values: [ YES, NO ]`,
				},
				resource.Attribute{
					Name:        "sc",
					Description: `(Optional) State of SureConnect for the service. Possible values: [ on, off ]`,
				},
				resource.Attribute{
					Name:        "sp",
					Description: `(Optional) Enable surge protection for the service. Possible values: [ on, off ]`,
				},
				resource.Attribute{
					Name:        "rtspsessionidremap",
					Description: `(Optional) Enable RTSP session ID mapping for the service. Possible values: [ on, off ]`,
				},
				resource.Attribute{
					Name:        "clttimeout",
					Description: `(Optional) Time, in seconds, after which to terminate an idle client connection.`,
				},
				resource.Attribute{
					Name:        "svrtimeout",
					Description: `(Optional) Time, in seconds, after which to terminate an idle server connection.`,
				},
				resource.Attribute{
					Name:        "customserverid",
					Description: `(Optional) Unique identifier for the service. Used when the persistency type for the virtual server is set to Custom Server ID.`,
				},
				resource.Attribute{
					Name:        "serverid",
					Description: `(Optional) The identifier for the service. This is used when the persistency type is set to Custom Server ID.`,
				},
				resource.Attribute{
					Name:        "cka",
					Description: `(Optional) Enable client keep-alive for the service. Possible values: [ YES, NO ]`,
				},
				resource.Attribute{
					Name:        "tcpb",
					Description: `(Optional) Enable TCP buffering for the service. Possible values: [ YES, NO ]`,
				},
				resource.Attribute{
					Name:        "cmp",
					Description: `(Optional) Enable compression for the service. Possible values: [ YES, NO ]`,
				},
				resource.Attribute{
					Name:        "maxbandwidth",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "monthreshold",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `(Optional) Initial state of the service. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "downstateflush",
					Description: `(Optional) Flush all active transactions associated with a service whose state transitions from UP to DOWN. Do not enable this option for applications that must complete their transactions. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "tcpprofilename",
					Description: `(Optional) Name of the TCP profile that contains TCP configuration settings for the service.`,
				},
				resource.Attribute{
					Name:        "httpprofilename",
					Description: `(Optional) Name of the HTTP profile that contains HTTP configuration settings for the service.`,
				},
				resource.Attribute{
					Name:        "contentinspectionprofilename",
					Description: `(Optional) Name of the ContentInspection profile that contains IPS/IDS communication related setting for the service.`,
				},
				resource.Attribute{
					Name:        "hashid",
					Description: `(Optional) A numerical identifier that can be used by hash based load balancing methods. Must be unique for each service.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `(Optional) Any information about the service.`,
				},
				resource.Attribute{
					Name:        "netprofile",
					Description: `(Optional) Network profile to use for the service.`,
				},
				resource.Attribute{
					Name:        "td",
					Description: `(Optional) Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.`,
				},
				resource.Attribute{
					Name:        "processlocal",
					Description: `(Optional) By turning on this option packets destined to a service in a cluster will not under go any steering. Turn this option for single packet request response mode or when the upstream device is performing a proper RSS for connection based distribution. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "dnsprofilename",
					Description: `(Optional) Name of the DNS profile to be associated with the service. DNS profile properties will applied to the transactions processed by a service. This parameter is valid only for ADNS and ADNS-TCP services.`,
				},
				resource.Attribute{
					Name:        "monconnectionclose",
					Description: `(Optional) Close monitoring connections by sending the service a connection termination message with the specified bit set. Possible values: [ RESET, FIN ]`,
				},
				resource.Attribute{
					Name:        "ipaddress",
					Description: `(Optional) The new IP address of the service.`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `(Optional) Weight to assign to the monitor-service binding. When a monitor is UP, the weight assigned to its binding with the service determines how much the monitor contributes toward keeping the health of the service above the value configured for the Monitor Threshold parameter.`,
				},
				resource.Attribute{
					Name:        "monitornamesvc",
					Description: `(Optional) Name of the monitor bound to the specified service.`,
				},
				resource.Attribute{
					Name:        "delay",
					Description: `(Optional) Time, in seconds, allocated to the Citrix ADC for a graceful shutdown of the service. During this period, new requests are sent to the service only for clients who already have persistent sessions on the appliance. Requests from new clients are load balanced among other available services. After the delay time expires, no requests are sent to the service, and the service is marked as unavailable (OUT OF SERVICE).`,
				},
				resource.Attribute{
					Name:        "graceful",
					Description: `(Optional) Shut down gracefully, not accepting any new connections, and disabling the service when all of its connections are closed. Possible values: [ YES, NO ]`,
				},
				resource.Attribute{
					Name:        "internal",
					Description: `(Optional) Display only dynamically learned services.`,
				},
				resource.Attribute{
					Name:        "lbvserver",
					Description: `(Optional) The lb vserver to attach the service to.`,
				},
				resource.Attribute{
					Name:        "lbmonitor",
					Description: `(Optional) The lb monitor to attach the service to.`,
				},
				resource.Attribute{
					Name:        "snienable",
					Description: `(Optional) State of the Server Name Indication (SNI) feature on the virtual server and service-based offload. SNI helps to enable SSL encryption on multiple domains on a single virtual server or service if the domains are controlled by the same organization and share the same second-level domain name. For example, \`,
				},
				resource.Attribute{
					Name:        "commonname",
					Description: `(Optional) Name to be checked against the CommonName (CN) field in the server certificate bound to the SSL service.`,
				},
				resource.Attribute{
					Name:        "all",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "riseapbrstatsmsgcode",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "accessdown",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "appflowlog",
					Description: `(Optional) ## Attribute Reference In addition to the arguments, the following attributes are available:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of the service. It has the same value as the ` + "`" + `name` + "`" + ` attribute. ## Import A service can be imported using its name, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import citrixadc_service.tf_service tf_service ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id of the service. It has the same value as the ` + "`" + `name` + "`" + ` attribute. ## Import A service can be imported using its name, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import citrixadc_service.tf_service tf_service ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "citrixadc_servicegroup",
			Category:         "Basic",
			ShortDescription: ``,
			Description: `

The servicegroup resource is used to create servicegroups.


`,
			Keywords: []string{
				"basic",
				"servicegroup",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "servicegroupname",
					Description: `(Optional) Name of the service group.`,
				},
				resource.Attribute{
					Name:        "servicetype",
					Description: `(Optional) Protocol used to exchange data with the service. Possible values: [ HTTP, FTP, TCP, UDP, SSL, SSL_BRIDGE, SSL_TCP, DTLS, NNTP, RPCSVR, DNS, ADNS, SNMP, RTSP, DHCPRA, ANY, SIP_UDP, SIP_TCP, SIP_SSL, DNS_TCP, ADNS_TCP, MYSQL, MSSQL, ORACLE, MONGO, MONGO_TLS, RADIUS, RADIUSListener, RDP, DIAMETER, SSL_DIAMETER, TFTP, SMPP, PPTP, GRE, SYSLOGTCP, SYSLOGUDP, FIX, SSL_FIX, USER_TCP, USER_SSL_TCP, QUIC, IPFIX, LOGSTREAM, LOGSTREAM_SSL ]`,
				},
				resource.Attribute{
					Name:        "cachetype",
					Description: `(Optional) Cache type supported by the cache server. Possible values: [ TRANSPARENT, REVERSE, FORWARD ]`,
				},
				resource.Attribute{
					Name:        "td",
					Description: `(Optional) Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.`,
				},
				resource.Attribute{
					Name:        "maxclient",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "maxreq",
					Description: `(Optional) Note: Connection requests beyond this value are rejected.`,
				},
				resource.Attribute{
					Name:        "cacheable",
					Description: `(Optional) Use the transparent cache redirection virtual server to forward the request to the cache server. Note: Do not set this parameter if you set the Cache Type. Possible values: [ YES, NO ]`,
				},
				resource.Attribute{
					Name:        "cip",
					Description: `(Optional) Insert the Client IP header in requests forwarded to the service. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "cipheader",
					Description: `(Optional) Name of the HTTP header whose value must be set to the IP address of the client. Used with the Client IP parameter. If client IP insertion is enabled, and the client IP header is not specified, the value of Client IP Header parameter or the value set by the set ns config command is used as client's IP header name.`,
				},
				resource.Attribute{
					Name:        "usip",
					Description: `(Optional) Use client's IP address as the source IP address when initiating connection to the server. With the NO setting, which is the default, a mapped IP (MIP) address or subnet IP (SNIP) address is used as the source IP address to initiate server side connections. Possible values: [ YES, NO ]`,
				},
				resource.Attribute{
					Name:        "pathmonitor",
					Description: `(Optional) Path monitoring for clustering. Possible values: [ YES, NO ]`,
				},
				resource.Attribute{
					Name:        "pathmonitorindv",
					Description: `(Optional) Individual Path monitoring decisions. Possible values: [ YES, NO ]`,
				},
				resource.Attribute{
					Name:        "useproxyport",
					Description: `(Optional) Use the proxy port as the source port when initiating connections with the server. With the NO setting, the client-side connection port is used as the source port for the server-side connection. Note: This parameter is available only when the Use Source IP (USIP) parameter is set to YES. Possible values: [ YES, NO ]`,
				},
				resource.Attribute{
					Name:        "healthmonitor",
					Description: `(Optional) Monitor the health of this service. Available settings function as follows: YES - Send probes to check the health of the service. NO - Do not send probes to check the health of the service. With the NO option, the appliance shows the service as UP at all times. Possible values: [ YES, NO ]`,
				},
				resource.Attribute{
					Name:        "sc",
					Description: `(Optional) State of the SureConnect feature for the service group. Possible values: [ on, off ]`,
				},
				resource.Attribute{
					Name:        "sp",
					Description: `(Optional) Enable surge protection for the service group. Possible values: [ on, off ]`,
				},
				resource.Attribute{
					Name:        "rtspsessionidremap",
					Description: `(Optional) Enable RTSP session ID mapping for the service group. Possible values: [ on, off ]`,
				},
				resource.Attribute{
					Name:        "clttimeout",
					Description: `(Optional) Time, in seconds, after which to terminate an idle client connection.`,
				},
				resource.Attribute{
					Name:        "svrtimeout",
					Description: `(Optional) Time, in seconds, after which to terminate an idle server connection.`,
				},
				resource.Attribute{
					Name:        "cka",
					Description: `(Optional) Enable client keep-alive for the service group. Possible values: [ YES, NO ]`,
				},
				resource.Attribute{
					Name:        "tcpb",
					Description: `(Optional) Enable TCP buffering for the service group. Possible values: [ YES, NO ]`,
				},
				resource.Attribute{
					Name:        "cmp",
					Description: `(Optional) Enable compression for the specified service. Possible values: [ YES, NO ]`,
				},
				resource.Attribute{
					Name:        "maxbandwidth",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "monthreshold",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `(Optional) Initial state of the service group. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "downstateflush",
					Description: `(Optional) Flush all active transactions associated with all the services in the service group whose state transitions from UP to DOWN. Do not enable this option for applications that must complete their transactions. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "tcpprofilename",
					Description: `(Optional) Name of the TCP profile that contains TCP configuration settings for the service group.`,
				},
				resource.Attribute{
					Name:        "httpprofilename",
					Description: `(Optional) Name of the HTTP profile that contains HTTP configuration settings for the service group.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `(Optional) Any information about the service group.`,
				},
				resource.Attribute{
					Name:        "appflowlog",
					Description: `(Optional) Enable logging of AppFlow information for the specified service group. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "netprofile",
					Description: `(Optional) Network profile for the service group.`,
				},
				resource.Attribute{
					Name:        "autoscale",
					Description: `(Optional) Auto scale option for a servicegroup. Possible values: [ DISABLED, DNS, POLICY, CLOUD, API ]`,
				},
				resource.Attribute{
					Name:        "memberport",
					Description: `(Optional) member port.`,
				},
				resource.Attribute{
					Name:        "autodisablegraceful",
					Description: `(Optional) Indicates graceful shutdown of the service. System will wait for all outstanding connections to this service to be closed before disabling the service. Possible values: [ YES, NO ]`,
				},
				resource.Attribute{
					Name:        "autodisabledelay",
					Description: `(Optional) The time allowed (in seconds) for a graceful shutdown. During this period, new connections or requests will continue to be sent to this service for clients who already have a persistent session on the system. Connections or requests from fresh or new clients who do not yet have a persistence sessions on the system will not be sent to the service. Instead, they will be load balanced among other available services. After the delay time expires, no new requests or connections will be sent to the service.`,
				},
				resource.Attribute{
					Name:        "monconnectionclose",
					Description: `(Optional) Close monitoring connections by sending the service a connection termination message with the specified bit set. Possible values: [ RESET, FIN ]`,
				},
				resource.Attribute{
					Name:        "servername",
					Description: `(Optional) Name of the server to which to bind the service group.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional) Server port number. Range 1 - 65535`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `(Optional) Weight to assign to the servers in the service group. Specifies the capacity of the servers relative to the other servers in the load balancing configuration. The higher the weight, the higher the percentage of requests sent to the service.`,
				},
				resource.Attribute{
					Name:        "customserverid",
					Description: `(Optional) The identifier for this IP:Port pair. Used when the persistency type is set to Custom Server ID.`,
				},
				resource.Attribute{
					Name:        "serverid",
					Description: `(Optional) The identifier for the service. This is used when the persistency type is set to Custom Server ID.`,
				},
				resource.Attribute{
					Name:        "hashid",
					Description: `(Optional) The hash identifier for the service. This must be unique for each service. This parameter is used by hash based load balancing methods.`,
				},
				resource.Attribute{
					Name:        "nameserver",
					Description: `(Optional) Specify the nameserver to which the query for bound domain needs to be sent. If not specified, use the global nameserver.`,
				},
				resource.Attribute{
					Name:        "dbsttl",
					Description: `(Optional) Specify the TTL for DNS record for domain based service.The default value of ttl is 0 which indicates to use the TTL received in DNS response for monitors.`,
				},
				resource.Attribute{
					Name:        "monitornamesvc",
					Description: `(Optional) Name of the monitor bound to the service group. Used to assign a weight to the monitor.`,
				},
				resource.Attribute{
					Name:        "dupweight",
					Description: `(Optional) weight of the monitor that is bound to servicegroup.`,
				},
				resource.Attribute{
					Name:        "delay",
					Description: `(Optional) Time, in seconds, allocated for a shutdown of the services in the service group. During this period, new requests are sent to the service only for clients who already have persistent sessions on the appliance. Requests from new clients are load balanced among other available services. After the delay time expires, no requests are sent to the service, and the service is marked as unavailable (OUT OF SERVICE).`,
				},
				resource.Attribute{
					Name:        "graceful",
					Description: `(Optional) Wait for all existing connections to the service to terminate before shutting down the service. Possible values: [ YES, NO ]`,
				},
				resource.Attribute{
					Name:        "includemembers",
					Description: `(Optional) Display the members of the listed service groups in addition to their settings. Can be specified when no service group name is provided in the command. In that case, the details displayed for each service group are identical to the details displayed when a service group name is provided, except that bound monitors are not displayed.`,
				},
				resource.Attribute{
					Name:        "lbvservers",
					Description: `(Optional) List of lbvservers to bind the servicegroup to.`,
				},
				resource.Attribute{
					Name:        "servicegroupmembers_by_servername",
					Description: `(Optional) list of service members bindings by service name. e.g. ` + "`" + `["service1:80:1", "service2:80:1"]` + "`" + ``,
				},
				resource.Attribute{
					Name:        "servicegroupmembers",
					Description: `(Optional) list of members bindings by server ip address. e.g.` + "`" + `["172.20.0.9:80:10", "172.20.0.10:80:10"]`,
				},
				resource.Attribute{
					Name:        "lbmonitor",
					Description: `(Optional) lbmonitor to bind the servicegroup to.`,
				},
				resource.Attribute{
					Name:        "riseapbrstatsmsgcode",
					Description: `(Optional) ## Attribute Reference In addition to the arguments, the following attributes are available:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of the servicegroup. It has the same value as the ` + "`" + `servicegroupname` + "`" + ` attribute. ## Import A servicegroup can be imported using its name, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import citrixadc_servicegroup.tf_servicegroup tf_servicegroup ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id of the servicegroup. It has the same value as the ` + "`" + `servicegroupname` + "`" + ` attribute. ## Import A servicegroup can be imported using its name, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import citrixadc_servicegroup.tf_servicegroup tf_servicegroup ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "citrixadc_servicegroup_lbmonitor_binding",
			Category:         "Basic",
			ShortDescription: ``,
			Description: `

The servicegroup_lbmonitor_binding resource is used to bind servicegroups to load balancing monitors.

~> If you are using this resource to bind lbmonitors to a servicegroup,
do not define the ` + "`" + `lbmonitor` + "`" + ` attribute in the servicegroup resource.

`,
			Keywords: []string{
				"basic",
				"servicegroup",
				"lbmonitor",
				"binding",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "monitorname",
					Description: `(Optional) Monitor name.`,
				},
				resource.Attribute{
					Name:        "monstate",
					Description: `(Optional) Monitor state. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `(Optional) Weight to assign to the servers in the service group. Specifies the capacity of the servers relative to the other servers in the load balancing configuration. The higher the weight, the higher the percentage of requests sent to the service.`,
				},
				resource.Attribute{
					Name:        "passive",
					Description: `(Optional) Indicates if load monitor is passive. A passive load monitor does not remove service from LB decision when threshold is breached.`,
				},
				resource.Attribute{
					Name:        "servicegroupname",
					Description: `(Optional) Name of the service group.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional) Port number of the service. Each service must have a unique port number. Range 1 - 65535`,
				},
				resource.Attribute{
					Name:        "customserverid",
					Description: `(Optional) Unique service identifier. Used when the persistency type for the virtual server is set to Custom Server ID.`,
				},
				resource.Attribute{
					Name:        "serverid",
					Description: `(Optional) The identifier for the service. This is used when the persistency type is set to Custom Server ID.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `(Optional) Initial state of the service after binding. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "hashid",
					Description: `(Optional) Unique numerical identifier used by hash based load balancing methods to identify a service.`,
				},
				resource.Attribute{
					Name:        "nameserver",
					Description: `(Optional) Specify the nameserver to which the query for bound domain needs to be sent. If not specified, use the global nameserver.`,
				},
				resource.Attribute{
					Name:        "dbsttl",
					Description: `(Optional) Specify the TTL for DNS record for domain based service.The default value of ttl is 0 which indicates to use the TTL received in DNS response for monitors. ## Attribute Reference In addition to the arguments, the following attributes are available:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of the servicegroup_lbmonitor_binding. It is the concatenation of the ` + "`" + `servicegroupname` + "`" + ` and ` + "`" + `monitorname` + "`" + ` attributes ## Import A servicegroup_lbmonitor_binding can be imported using its id, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import citrixadc_servicegroup_lbmonitor_binding.tf_binding tf_csaction tf_servicegroup,tf_monitorname ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id of the servicegroup_lbmonitor_binding. It is the concatenation of the ` + "`" + `servicegroupname` + "`" + ` and ` + "`" + `monitorname` + "`" + ` attributes ## Import A servicegroup_lbmonitor_binding can be imported using its id, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import citrixadc_servicegroup_lbmonitor_binding.tf_binding tf_csaction tf_servicegroup,tf_monitorname ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "citrixadc_servicegroup_servicegroupmember_binding",
			Category:         "Basic",
			ShortDescription: ``,
			Description: `\_servicegroupmember\_binding

The servicegroup\_servicegroupmember\_binding resource is used to bind service members to servicegroups.


`,
			Keywords: []string{
				"basic",
				"servicegroup",
				"servicegroupmember",
				"binding",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ip",
					Description: `(Optional) IP Address.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional) Server port number. Range 1 - 65535`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `(Optional) Weight to assign to the servers in the service group. Specifies the capacity of the servers relative to the other servers in the load balancing configuration. The higher the weight, the higher the percentage of requests sent to the service.`,
				},
				resource.Attribute{
					Name:        "servername",
					Description: `(Optional) Name of the server to which to bind the service group.`,
				},
				resource.Attribute{
					Name:        "customserverid",
					Description: `(Optional) The identifier for this IP:Port pair. Used when the persistency type is set to Custom Server ID.`,
				},
				resource.Attribute{
					Name:        "serverid",
					Description: `(Optional) The identifier for the service. This is used when the persistency type is set to Custom Server ID.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `(Optional) Initial state of the service group. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "hashid",
					Description: `(Optional) The hash identifier for the service. This must be unique for each service. This parameter is used by hash based load balancing methods.`,
				},
				resource.Attribute{
					Name:        "nameserver",
					Description: `(Optional) Specify the nameserver to which the query for bound domain needs to be sent. If not specified, use the global nameserver.`,
				},
				resource.Attribute{
					Name:        "dbsttl",
					Description: `(Optional) Specify the TTL for DNS record for domain based service.The default value of ttl is 0 which indicates to use the TTL received in DNS response for monitors.`,
				},
				resource.Attribute{
					Name:        "servicegroupname",
					Description: `(Required) Name of the service group. ## Attribute Reference In addition to the arguments, the following attributes are available:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of the servicegroup\_servicegroupmember\_binding. It is the concatenation of three components separated by comma. First component is the ` + "`" + `servicegroupname` + "`" + `. Second component is the ` + "`" + `ip` + "`" + ` or the ` + "`" + `servername` + "`" + ` attribute. Last optional component is the ` + "`" + `port` + "`" + ` attribute.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id of the servicegroup\_servicegroupmember\_binding. It is the concatenation of three components separated by comma. First component is the ` + "`" + `servicegroupname` + "`" + `. Second component is the ` + "`" + `ip` + "`" + ` or the ` + "`" + `servername` + "`" + ` attribute. Last optional component is the ` + "`" + `port` + "`" + ` attribute.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "citrixadc_sslaction",
			Category:         "SSL",
			ShortDescription: ``,
			Description: `

The sslaction resource is used to create ssl actions.


`,
			Keywords: []string{
				"ssl",
				"sslaction",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name for the SSL action.`,
				},
				resource.Attribute{
					Name:        "clientauth",
					Description: `(Optional) Perform client certificate authentication. Possible values: [ DOCLIENTAUTH, NOCLIENTAUTH ]`,
				},
				resource.Attribute{
					Name:        "clientcertverification",
					Description: `(Optional) Client certificate verification is mandatory or optional. Possible values: [ Mandatory, Optional ]`,
				},
				resource.Attribute{
					Name:        "ssllogprofile",
					Description: `(Optional) The name of the ssllogprofile.`,
				},
				resource.Attribute{
					Name:        "clientcert",
					Description: `(Optional) Insert the entire client certificate into the HTTP header of the request being sent to the web server. The certificate is inserted in ASCII (PEM) format. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "certheader",
					Description: `(Optional) Name of the header into which to insert the client certificate.`,
				},
				resource.Attribute{
					Name:        "clientcertserialnumber",
					Description: `(Optional) Insert the entire client serial number into the HTTP header of the request being sent to the web server. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "certserialheader",
					Description: `(Optional) Name of the header into which to insert the client serial number.`,
				},
				resource.Attribute{
					Name:        "clientcertsubject",
					Description: `(Optional) Insert the client certificate subject, also known as the distinguished name (DN), into the HTTP header of the request being sent to the web server. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "certsubjectheader",
					Description: `(Optional) Name of the header into which to insert the client certificate subject.`,
				},
				resource.Attribute{
					Name:        "clientcerthash",
					Description: `(Optional) Insert the certificate's signature into the HTTP header of the request being sent to the web server. The signature is the value extracted directly from the X.509 certificate signature field. All X.509 certificates contain a signature field. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "certhashheader",
					Description: `(Optional) Name of the header into which to insert the client certificate signature (hash).`,
				},
				resource.Attribute{
					Name:        "clientcertfingerprint",
					Description: `(Optional) Insert the certificate's fingerprint into the HTTP header of the request being sent to the web server. The fingerprint is derived by computing the specified hash value (SHA256, for example) of the DER-encoding of the client certificate. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "certfingerprintheader",
					Description: `(Optional) Name of the header into which to insert the client certificate fingerprint.`,
				},
				resource.Attribute{
					Name:        "certfingerprintdigest",
					Description: `(Optional) Digest algorithm used to compute the fingerprint of the client certificate. Possible values: [ SHA1, SHA224, SHA256, SHA384, SHA512 ]`,
				},
				resource.Attribute{
					Name:        "clientcertissuer",
					Description: `(Optional) Insert the certificate issuer details into the HTTP header of the request being sent to the web server. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "certissuerheader",
					Description: `(Optional) Name of the header into which to insert the client certificate issuer details.`,
				},
				resource.Attribute{
					Name:        "sessionid",
					Description: `(Optional) Insert the SSL session ID into the HTTP header of the request being sent to the web server. Every SSL connection that the client and the Citrix ADC share has a unique ID that identifies the specific connection. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "sessionidheader",
					Description: `(Optional) Name of the header into which to insert the Session ID.`,
				},
				resource.Attribute{
					Name:        "cipher",
					Description: `(Optional) Insert the cipher suite that the client and the Citrix ADC negotiated for the SSL session into the HTTP header of the request being sent to the web server. The appliance inserts the cipher-suite name, SSL protocol, export or non-export string, and cipher strength bit, depending on the type of browser connecting to the SSL virtual server or service (for example, Cipher-Suite: RC4- MD5 SSLv3 Non-Export 128-bit). Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "cipherheader",
					Description: `(Optional) Name of the header into which to insert the name of the cipher suite.`,
				},
				resource.Attribute{
					Name:        "clientcertnotbefore",
					Description: `(Optional) Insert the date from which the certificate is valid into the HTTP header of the request being sent to the web server. Every certificate is configured with the date and time from which it is valid. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "certnotbeforeheader",
					Description: `(Optional) Name of the header into which to insert the date and time from which the certificate is valid.`,
				},
				resource.Attribute{
					Name:        "clientcertnotafter",
					Description: `(Optional) Insert the date of expiry of the certificate into the HTTP header of the request being sent to the web server. Every certificate is configured with the date and time at which the certificate expires. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "certnotafterheader",
					Description: `(Optional) Name of the header into which to insert the certificate's expiry date.`,
				},
				resource.Attribute{
					Name:        "owasupport",
					Description: `(Optional) If the appliance is in front of an Outlook Web Access (OWA) server, insert a special header field, FRONT-END-HTTPS: ON, into the HTTP requests going to the OWA server. This header communicates to the server that the transaction is HTTPS and not HTTP. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "forward",
					Description: `(Optional) This action takes an argument a vserver name, to this vserver one will be able to forward all the packets.`,
				},
				resource.Attribute{
					Name:        "cacertgrpname",
					Description: `(Optional) This action will allow to pick CA(s) from the specific CA group, to verify the client certificate. ## Attribute Reference In addition to the arguments, the following attributes are available:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of the sslaction. It has the same value as the ` + "`" + `name` + "`" + ` attribute. ## Import A sslaction can be imported using its name, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import citrixadc_sslaction.tf_sslaction tf_sslaction ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id of the sslaction. It has the same value as the ` + "`" + `name` + "`" + ` attribute. ## Import A sslaction can be imported using its name, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import citrixadc_sslaction.tf_sslaction tf_sslaction ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "citrixadc_sslcertkey",
			Category:         "SSL",
			ShortDescription: ``,
			Description: `

The sslcertkey resource is used to create TLS certificate keys.


`,
			Keywords: []string{
				"ssl",
				"sslcertkey",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "certkey",
					Description: `(Optional) Name for the certificate and private-key pair.`,
				},
				resource.Attribute{
					Name:        "cert",
					Description: `(Optional) Name of and, optionally, path to the X509 certificate file that is used to form the certificate-key pair. The certificate file should be present on the appliance's hard-disk drive or solid-state drive. Storing a certificate in any location other than the default might cause inconsistency in a high availability setup. /nsconfig/ssl/ is the default path.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Optional) Name of and, optionally, path to the private-key file that is used to form the certificate-key pair. The certificate file should be present on the appliance's hard-disk drive or solid-state drive. Storing a certificate in any location other than the default might cause inconsistency in a high availability setup. /nsconfig/ssl/ is the default path.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional) Passphrase that was used to encrypt the private-key. Use this option to load encrypted private-keys in PEM format.`,
				},
				resource.Attribute{
					Name:        "fipskey",
					Description: `(Optional) Name of the FIPS key that was created inside the Hardware Security Module (HSM) of a FIPS appliance, or a key that was imported into the HSM.`,
				},
				resource.Attribute{
					Name:        "hsmkey",
					Description: `(Optional) Name of the HSM key that was created in the External Hardware Security Module (HSM) of a FIPS appliance.`,
				},
				resource.Attribute{
					Name:        "inform",
					Description: `(Optional) Input format of the certificate and the private-key files. The three formats supported by the appliance are: PEM - Privacy Enhanced Mail DER - Distinguished Encoding Rule PFX - Personal Information Exchange. Possible values: [ DER, PEM, PFX ]`,
				},
				resource.Attribute{
					Name:        "passplain",
					Description: `(Optional) Pass phrase used to encrypt the private-key. Required when adding an encrypted private-key in PEM format.`,
				},
				resource.Attribute{
					Name:        "expirymonitor",
					Description: `(Optional) Issue an alert when the certificate is about to expire. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "notificationperiod",
					Description: `(Optional) Time, in number of days, before certificate expiration, at which to generate an alert that the certificate is about to expire.`,
				},
				resource.Attribute{
					Name:        "bundle",
					Description: `(Optional) Parse the certificate chain as a single file after linking the server certificate to its issuer's certificate within the file. Possible values: [ YES, NO ]`,
				},
				resource.Attribute{
					Name:        "linkcertkeyname",
					Description: `(Optional) Name of the Certificate Authority certificate-key pair to which to link a certificate-key pair.`,
				},
				resource.Attribute{
					Name:        "nodomaincheck",
					Description: `(Optional) Override the check for matching domain names during a certificate update operation.`,
				},
				resource.Attribute{
					Name:        "ocspstaplingcache",
					Description: `(Optional) Clear cached ocspStapling response in certkey. ## Attribute Reference In addition to the arguments, the following attributes are available:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of the sslcertkey. It has the same value as the ` + "`" + `certkey` + "`" + ` attribute. ## Import A sslcertkey can be imported using its certkey, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import citrixadc_sslcertkey.tf_sslcertkey tf_sslcertkey ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id of the sslcertkey. It has the same value as the ` + "`" + `certkey` + "`" + ` attribute. ## Import A sslcertkey can be imported using its certkey, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import citrixadc_sslcertkey.tf_sslcertkey tf_sslcertkey ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "citrixadc_sslcipher",
			Category:         "SSL",
			ShortDescription: ``,
			Description: `

The sslcipher resource is used to create ssl ciphers.


`,
			Keywords: []string{
				"ssl",
				"sslcipher",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ciphergroupname",
					Description: `(Required) Name of the cipher group to be created.`,
				},
				resource.Attribute{
					Name:        "ciphersuitebinding",
					Description: `(Required) A set of ciphersuites bound to this cipher group. Any change to this set will recreate the whole cipher group. Attributes documented below. A ciphersuitebinding supports the following:`,
				},
				resource.Attribute{
					Name:        "ciphername",
					Description: `(Required) Cipher name.`,
				},
				resource.Attribute{
					Name:        "cipherpriority",
					Description: `(Optional) This indicates priority assigned to the particular cipher. ## Attribute Reference In addition to the arguments, the following attributes are available:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of the sslcipher. It has the same value as the ` + "`" + `ciphergroupname` + "`" + ` attribute. ## Import A sslcipher can be imported using its ciphergroupname, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import citrixadc_sslcipher.tf_sslcipher tf_sslcipher ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id of the sslcipher. It has the same value as the ` + "`" + `ciphergroupname` + "`" + ` attribute. ## Import A sslcipher can be imported using its ciphergroupname, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import citrixadc_sslcipher.tf_sslcipher tf_sslcipher ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "citrixadc_ssldhparam",
			Category:         "SSL",
			ShortDescription: ``,
			Description: `

The ssldhparam resource is used to configure ssl DH parameters.


`,
			Keywords: []string{
				"ssl",
				"ssldhparam",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "dhfile",
					Description: `(Required) Name of and, optionally, path to the DH key file. /nsconfig/ssl/ is the default path.`,
				},
				resource.Attribute{
					Name:        "bits",
					Description: `(Required) Size, in bits, of the DH key being generated.`,
				},
				resource.Attribute{
					Name:        "gen",
					Description: `(Optional) Random number required for generating the DH key. Required as part of the DH key generation algorithm. Possible values: [ 2, 5 ] ## Attribute Reference In addition to the arguments, the following attributes are available:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of the ssldhparam. It has the same value as the ` + "`" + `dhfile` + "`" + ` attribute.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id of the ssldhparam. It has the same value as the ` + "`" + `dhfile` + "`" + ` attribute.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "citrixadc_sslparameter",
			Category:         "SSL",
			ShortDescription: ``,
			Description: `

The sslparameter resource is used to update the ADC SSL parameters.


`,
			Keywords: []string{
				"ssl",
				"sslparameter",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "quantumsize",
					Description: `(Optional) Amount of data to collect before the data is pushed to the crypto hardware for encryption. For large downloads, a larger quantum size better utilizes the crypto resources. Possible values: [ 4096, 8192, 16384 ]`,
				},
				resource.Attribute{
					Name:        "crlmemorysizemb",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "strictcachecks",
					Description: `(Optional) Enable strict CA certificate checks on the appliance. Possible values: [ YES, NO ]`,
				},
				resource.Attribute{
					Name:        "ssltriggertimeout",
					Description: `(Optional) Time, in milliseconds, after which encryption is triggered for transactions that are not tracked on the Citrix ADC because their length is not known. There can be a delay of up to 10ms from the specified timeout value before the packet is pushed into the queue.`,
				},
				resource.Attribute{
					Name:        "sendclosenotify",
					Description: `(Optional) Send an SSL Close-Notify message to the client at the end of a transaction. Possible values: [ YES, NO ]`,
				},
				resource.Attribute{
					Name:        "encrypttriggerpktcount",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "denysslreneg",
					Description: `(Optional) Deny renegotiation in specified circumstances. Available settings function as follows:`,
				},
				resource.Attribute{
					Name:        "insertionencoding",
					Description: `(Optional) Encoding method used to insert the subject or issuer's name in HTTP requests to servers. Possible values: [ Unicode, UTF-8 ]`,
				},
				resource.Attribute{
					Name:        "ocspcachesize",
					Description: `(Optional) Size, per packet engine, in megabytes, of the OCSP cache. A maximum of 10% of the packet engine memory can be assigned. Because the maximum allowed packet engine memory is 4GB, the maximum value that can be assigned to the OCSP cache is approximately 410 MB.`,
				},
				resource.Attribute{
					Name:        "pushflag",
					Description: `(Optional) Insert PUSH flag into decrypted, encrypted, or all records. If the PUSH flag is set to a value other than 0, the buffered records are forwarded on the basis of the value of the PUSH flag. Available settings function as follows: 0 - Auto (PUSH flag is not set.) 1 - Insert PUSH flag into every decrypted record. 2 -Insert PUSH flag into every encrypted record. 3 - Insert PUSH flag into every decrypted and encrypted record.`,
				},
				resource.Attribute{
					Name:        "dropreqwithnohostheader",
					Description: `(Optional) Host header check for SNI enabled sessions. If this check is enabled and the HTTP request does not contain the host header for SNI enabled sessions(i.e vserver or profile bound to vserver has SNI enabled and 'Client Hello' arrived with SNI extension), the request is dropped. Possible values: [ YES, NO ]`,
				},
				resource.Attribute{
					Name:        "snihttphostmatch",
					Description: `(Optional) Controls how the HTTP 'Host' header value is validated. These checks are performed only if the session is SNI enabled (i.e when vserver or profile bound to vserver has SNI enabled and 'Client Hello' arrived with SNI extension) and HTTP request contains 'Host' header. Available settings function as follows: CERT - Request is forwarded if the 'Host' value is covered by the certificate used to establish this SSL session. Note: 'CERT' matching mode cannot be applied in TLS 1.3 connections established by resuming from a previous TLS 1.3 session. On these connections, 'STRICT' matching mode will be used instead. STRICT - Request is forwarded only if value of 'Host' header in HTTP is identical to the 'Server name' value passed in 'Client Hello' of the SSL connection. NO - No validation is performed on the HTTP 'Host' header value. Possible values: [ NO, CERT, STRICT ]`,
				},
				resource.Attribute{
					Name:        "pushenctriggertimeout",
					Description: `(Optional) PUSH encryption trigger timeout value. The timeout value is applied only if you set the Push Encryption Trigger parameter to Timer in the SSL virtual server settings.`,
				},
				resource.Attribute{
					Name:        "cryptodevdisablelimit",
					Description: `(Optional) Limit to the number of disabled SSL chips after which the ADC restarts. A value of zero implies that the ADC does not automatically restart.`,
				},
				resource.Attribute{
					Name:        "undefactioncontrol",
					Description: `(Optional) Name of the undefined built-in control action: CLIENTAUTH, NOCLIENTAUTH, NOOP, RESET, or DROP.`,
				},
				resource.Attribute{
					Name:        "undefactiondata",
					Description: `(Optional) Name of the undefined built-in data action: NOOP, RESET or DROP.`,
				},
				resource.Attribute{
					Name:        "defaultprofile",
					Description: `(Optional) Global parameter used to enable default profile feature. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "softwarecryptothreshold",
					Description: `(Optional) Citrix ADC CPU utilization threshold (in percentage) beyond which crypto operations are not done in software. A value of zero implies that CPU is not utilized for doing crypto in software.`,
				},
				resource.Attribute{
					Name:        "hybridfipsmode",
					Description: `(Optional) When this mode is enabled, system will use additional crypto hardware to accelerate symmetric crypto operations. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "sslierrorcache",
					Description: `(Optional) Enable or disable dynamically learning and caching the learned information to make the subsequent interception or bypass decision. When enabled, NS does the lookup of this cached data to do early bypass. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "sslimaxerrorcachemem",
					Description: `(Optional) Specify the maximum memory that can be used for caching the learned data. This memory is used as a LRU cache so that the old entries gets replaced with new entry once the set memory limit is fully utilised. A value of 0 decides the limit automatically.`,
				},
				resource.Attribute{
					Name:        "insertcertspace",
					Description: `(Optional) To insert space between lines in the certificate header of request. Possible values: [ YES, NO ]`,
				},
				resource.Attribute{
					Name:        "ndcppcompliancecertcheck",
					Description: `(Optional) Applies when the Citrix ADC appliance acts as a client (back-end connection). Settings apply as follows: YES - During certificate verification, ignore the common name if SAN is present in the certificate. NO - Do not ignore common name. Possible values: [ YES, NO ]`,
				},
				resource.Attribute{
					Name:        "heterogeneoussslhw",
					Description: `(Optional) To support both cavium and coleto based platforms in cluster environment, this mode has to be enabled. Possible values: [ ENABLED, DISABLED ] ## Attribute Reference In addition to the arguments, the following attributes are available:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of the sslparameter. It is a unique string prefixed with "tf-sslparameter-"`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id of the sslparameter. It is a unique string prefixed with "tf-sslparameter-"`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "citrixadc_sslpolicy",
			Category:         "SSL",
			ShortDescription: ``,
			Description: `

The sslpolicy resource is used to create SSL policies.


`,
			Keywords: []string{
				"ssl",
				"sslpolicy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name for the new SSL policy.`,
				},
				resource.Attribute{
					Name:        "rule",
					Description: `(Optional) Expression, against which traffic is evaluated. The following requirements apply only to the Citrix ADC CLI:`,
				},
				resource.Attribute{
					Name:        "reqaction",
					Description: `(Optional) The name of the action to be performed on the request. Refer to 'add ssl action' command to add a new action. Builtin actions like NOOP, RESET, DROP, CLIENTAUTH and NOCLIENTAUTH are also allowed.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Optional) Name of the built-in or user-defined action to perform on the request. Available built-in actions are NOOP, RESET, DROP, CLIENTAUTH, NOCLIENTAUTH, INTERCEPT AND BYPASS.`,
				},
				resource.Attribute{
					Name:        "undefaction",
					Description: `(Optional) Name of the action to be performed when the result of rule evaluation is undefined. Possible values for control policies: CLIENTAUTH, NOCLIENTAUTH, NOOP, RESET, DROP. Possible values for data policies: NOOP, RESET, DROP and BYPASS.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `(Optional) Any comments associated with this policy. ## Attribute Reference In addition to the arguments, the following attributes are available:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of the sslpolicy. It has the same value as the ` + "`" + `name` + "`" + ` attribute. ## Import A sslpolicy can be imported using its name, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import citrixadc_sslpolicy.tf_sslpolicy tf_sslpolicy ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id of the sslpolicy. It has the same value as the ` + "`" + `name` + "`" + ` attribute. ## Import A sslpolicy can be imported using its name, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import citrixadc_sslpolicy.tf_sslpolicy tf_sslpolicy ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "citrixadc_sslprofile",
			Category:         "SSL",
			ShortDescription: ``,
			Description: `

The sslprofile resource is used to create SSL profiles.


`,
			Keywords: []string{
				"ssl",
				"sslprofile",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name for the SSL profile.`,
				},
				resource.Attribute{
					Name:        "sslprofiletype",
					Description: `(Optional) Type of profile. Front end profiles apply to the entity that receives requests from a client. Backend profiles apply to the entity that sends client requests to a server. Possible values: [ BackEnd, FrontEnd ]`,
				},
				resource.Attribute{
					Name:        "ssllogprofile",
					Description: `(Optional) The name of the ssllogprofile.`,
				},
				resource.Attribute{
					Name:        "dhcount",
					Description: `(Optional) Number of interactions, between the client and the Citrix ADC, after which the DH private-public pair is regenerated. A value of zero (0) specifies infinite use (no refresh). This parameter is not applicable when configuring a backend profile. Allowed DH count values are 0 and >= 500.`,
				},
				resource.Attribute{
					Name:        "dh",
					Description: `(Optional) State of Diffie-Hellman (DH) key exchange. This parameter is not applicable when configuring a backend profile. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "dhfile",
					Description: `(Optional) The file name and path for the DH parameter.`,
				},
				resource.Attribute{
					Name:        "ersa",
					Description: `(Optional) State of Ephemeral RSA (eRSA) key exchange. Ephemeral RSA allows clients that support only export ciphers to communicate with the secure server even if the server certificate does not support export clients. The ephemeral RSA key is automatically generated when you bind an export cipher to an SSL or TCP-based SSL virtual server or service. When you remove the export cipher, the eRSA key is not deleted. It is reused at a later date when another export cipher is bound to an SSL or TCP-based SSL virtual server or service. The eRSA key is deleted when the appliance restarts. This parameter is not applicable when configuring a backend profile. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "ersacount",
					Description: `(Optional) The refresh count for the re-generation of RSA public-key and private-key pair.`,
				},
				resource.Attribute{
					Name:        "sessreuse",
					Description: `(Optional) State of session reuse. Establishing the initial handshake requires CPU-intensive public key encryption operations. With the ENABLED setting, session key exchange is avoided for session resumption requests received from the client. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "sesstimeout",
					Description: `(Optional) The Session timeout value in seconds.`,
				},
				resource.Attribute{
					Name:        "cipherredirect",
					Description: `(Optional) State of Cipher Redirect. If this parameter is set to ENABLED, you can configure an SSL virtual server or service to display meaningful error messages if the SSL handshake fails because of a cipher mismatch between the virtual server or service and the client. This parameter is not applicable when configuring a backend profile. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "cipherurl",
					Description: `(Optional) The redirect URL to be used with the Cipher Redirect feature.`,
				},
				resource.Attribute{
					Name:        "clientauth",
					Description: `(Optional) State of client authentication. In service-based SSL offload, the service terminates the SSL handshake if the SSL client does not provide a valid certificate. This parameter is not applicable when configuring a backend profile. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "clientcert",
					Description: `(Optional) The rule for client certificate requirement in client authentication. Possible values: [ Mandatory, Optional ]`,
				},
				resource.Attribute{
					Name:        "dhkeyexpsizelimit",
					Description: `(Optional) This option enables the use of NIST recommended (NIST Special Publication 800-56A) bit size for private-key size. For example, for DH params of size 2048bit, the private-key size recommended is 224bits. This is rounded-up to 256bits. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "sslredirect",
					Description: `(Optional) State of HTTPS redirects for the SSL service. For an SSL session, if the client browser receives a redirect message, the browser tries to connect to the new location. However, the secure SSL session breaks if the object has moved from a secure site (https://) to an unsecure site (http://). Typically, a warning message appears on the screen, prompting the user to continue or disconnect. If SSL Redirect is ENABLED, the redirect message is automatically converted from http:// to https:// and the SSL session does not break. This parameter is not applicable when configuring a backend profile. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "redirectportrewrite",
					Description: `(Optional) State of the port rewrite while performing HTTPS redirect. If this parameter is set to ENABLED, and the URL from the server does not contain the standard port, the port is rewritten to the standard. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "ssl3",
					Description: `(Optional) State of SSLv3 protocol support for the SSL profile. Note: On platforms with SSL acceleration chips, if the SSL chip does not support SSLv3, this parameter cannot be set to ENABLED. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "tls1",
					Description: `(Optional) State of TLSv1.0 protocol support for the SSL profile. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "tls11",
					Description: `(Optional) State of TLSv1.1 protocol support for the SSL profile. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "tls12",
					Description: `(Optional) State of TLSv1.2 protocol support for the SSL profile. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "tls13",
					Description: `(Optional) State of TLSv1.3 protocol support for the SSL profile. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "snienable",
					Description: `(Optional) State of the Server Name Indication (SNI) feature on the virtual server and service-based offload. SNI helps to enable SSL encryption on multiple domains on a single virtual server or service if the domains are controlled by the same organization and share the same second-level domain name. For example, \`,
				},
				resource.Attribute{
					Name:        "ocspstapling",
					Description: `(Optional) State of OCSP stapling support on the SSL virtual server. Supported only if the protocol used is higher than SSLv3. Possible values: ENABLED: The appliance sends a request to the OCSP responder to check the status of the server certificate and caches the response for the specified time. If the response is valid at the time of SSL handshake with the client, the OCSP-based server certificate status is sent to the client during the handshake. DISABLED: The appliance does not check the status of the server certificate. . Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "serverauth",
					Description: `(Optional) State of server authentication support for the SSL Backend profile. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "commonname",
					Description: `(Optional) Name to be checked against the CommonName (CN) field in the server certificate bound to the SSL server.`,
				},
				resource.Attribute{
					Name:        "pushenctrigger",
					Description: `(Optional) Trigger encryption on the basis of the PUSH flag value. Available settings function as follows:`,
				},
				resource.Attribute{
					Name:        "sendclosenotify",
					Description: `(Optional) Enable sending SSL Close-Notify at the end of a transaction. Possible values: [ YES, NO ]`,
				},
				resource.Attribute{
					Name:        "cleartextport",
					Description: `(Optional) Port on which clear-text data is sent by the appliance to the server. Do not specify this parameter for SSL offloading with end-to-end encryption. Range 1 - 65535`,
				},
				resource.Attribute{
					Name:        "insertionencoding",
					Description: `(Optional) Encoding method used to insert the subject or issuer's name in HTTP requests to servers. Possible values: [ Unicode, UTF-8 ]`,
				},
				resource.Attribute{
					Name:        "denysslreneg",
					Description: `(Optional) Deny renegotiation in specified circumstances. Available settings function as follows:`,
				},
				resource.Attribute{
					Name:        "quantumsize",
					Description: `(Optional) Amount of data to collect before the data is pushed to the crypto hardware for encryption. For large downloads, a larger quantum size better utilizes the crypto resources. Possible values: [ 4096, 8192, 16384 ]`,
				},
				resource.Attribute{
					Name:        "strictcachecks",
					Description: `(Optional) Enable strict CA certificate checks on the appliance. Possible values: [ YES, NO ]`,
				},
				resource.Attribute{
					Name:        "encrypttriggerpktcount",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "pushflag",
					Description: `(Optional) Insert PUSH flag into decrypted, encrypted, or all records. If the PUSH flag is set to a value other than 0, the buffered records are forwarded on the basis of the value of the PUSH flag. Available settings function as follows: 0 - Auto (PUSH flag is not set.) 1 - Insert PUSH flag into every decrypted record. 2 -Insert PUSH flag into every encrypted record. 3 - Insert PUSH flag into every decrypted and encrypted record.`,
				},
				resource.Attribute{
					Name:        "dropreqwithnohostheader",
					Description: `(Optional) Host header check for SNI enabled sessions. If this check is enabled and the HTTP request does not contain the host header for SNI enabled sessions(i.e vserver or profile bound to vserver has SNI enabled and 'Client Hello' arrived with SNI extension), the request is dropped. Possible values: [ YES, NO ]`,
				},
				resource.Attribute{
					Name:        "snihttphostmatch",
					Description: `(Optional) Controls how the HTTP 'Host' header value is validated. These checks are performed only if the session is SNI enabled (i.e when vserver or profile bound to vserver has SNI enabled and 'Client Hello' arrived with SNI extension) and HTTP request contains 'Host' header. Available settings function as follows: CERT - Request is forwarded if the 'Host' value is covered by the certificate used to establish this SSL session. Note: 'CERT' matching mode cannot be applied in TLS 1.3 connections established by resuming from a previous TLS 1.3 session. On these connections, 'STRICT' matching mode will be used instead. STRICT - Request is forwarded only if value of 'Host' header in HTTP is identical to the 'Server name' value passed in 'Client Hello' of the SSL connection. NO - No validation is performed on the HTTP 'Host' header value. Possible values: [ NO, CERT, STRICT ]`,
				},
				resource.Attribute{
					Name:        "pushenctriggertimeout",
					Description: `(Optional) PUSH encryption trigger timeout value. The timeout value is applied only if you set the Push Encryption Trigger parameter to Timer in the SSL virtual server settings.`,
				},
				resource.Attribute{
					Name:        "ssltriggertimeout",
					Description: `(Optional) Time, in milliseconds, after which encryption is triggered for transactions that are not tracked on the Citrix ADC because their length is not known. There can be a delay of up to 10ms from the specified timeout value before the packet is pushed into the queue.`,
				},
				resource.Attribute{
					Name:        "clientauthuseboundcachain",
					Description: `(Optional) Certficates bound on the VIP are used for validating the client cert. Certficates came along with client cert are not used for validating the client cert. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "sslinterception",
					Description: `(Optional) Enable or disable transparent interception of SSL sessions. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "sslireneg",
					Description: `(Optional) Enable or disable triggering the client renegotiation when renegotiation request is received from the origin server. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "ssliocspcheck",
					Description: `(Optional) Enable or disable OCSP check for origin server certificate. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "sslimaxsessperserver",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "sessionticket",
					Description: `(Optional) This option enables the use of session tickets, as per the RFC 5077. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "sessionticketlifetime",
					Description: `(Optional) This option sets the life time of session tickets issued by NS in secs.`,
				},
				resource.Attribute{
					Name:        "sessionticketkeyrefresh",
					Description: `(Optional) This option enables the use of session tickets, as per the RFC 5077. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "sessionticketkeydata",
					Description: `(Optional) Session ticket enc/dec key , admin can set it.`,
				},
				resource.Attribute{
					Name:        "sessionkeylifetime",
					Description: `(Optional) This option sets the life time of symm key used to generate session tickets issued by NS in secs.`,
				},
				resource.Attribute{
					Name:        "prevsessionkeylifetime",
					Description: `(Optional) This option sets the life time of symm key used to generate session tickets issued by NS in secs.`,
				},
				resource.Attribute{
					Name:        "hsts",
					Description: `(Optional) State of HSTS protocol support for the SSL profile. Using HSTS, a server can enforce the use of an HTTPS connection for all communication with a client. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "maxage",
					Description: `(Optional) Set the maximum time, in seconds, in the strict transport security (STS) header during which the client must send only HTTPS requests to the server.`,
				},
				resource.Attribute{
					Name:        "includesubdomains",
					Description: `(Optional) Enable HSTS for subdomains. If set to Yes, a client must send only HTTPS requests for subdomains. Possible values: [ YES, NO ]`,
				},
				resource.Attribute{
					Name:        "preload",
					Description: `(Optional) Flag indicates the consent of the site owner to have their domain preloaded. Possible values: [ YES, NO ]`,
				},
				resource.Attribute{
					Name:        "skipclientcertpolicycheck",
					Description: `(Optional) This flag controls the processing of X509 certificate policies. If this option is Enabled, then the policy check in Client authentication will be skipped. This option can be used only when Client Authentication is Enabled and ClientCert is set to Mandatory. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "zerorttearlydata",
					Description: `(Optional) State of TLS 1.3 0-RTT early data support for the SSL Virtual Server. This setting only has an effect if resumption is enabled, as early data cannot be sent along with an initial handshake. Early application data has significantly different security properties - in particular there is no guarantee that the data cannot be replayed. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "tls13sessionticketsperauthcontext",
					Description: `(Optional) Number of tickets the SSL Virtual Server will issue anytime TLS 1.3 is negotiated, ticket-based resumption is enabled, and either (1) a handshake completes or (2) post-handhsake client auth completes. This value can be increased to enable clients to open multiple parallel connections using a fresh ticket for each connection. No tickets are sent if resumption is disabled.`,
				},
				resource.Attribute{
					Name:        "dhekeyexchangewithpsk",
					Description: `(Optional) Whether or not the SSL Virtual Server will require a DHE key exchange to occur when a PSK is accepted during a TLS 1.3 resumption handshake. A DHE key exchange ensures forward secrecy even in the event that ticket keys are compromised, at the expense of an additional round trip and resources required to carry out the DHE key exchange. If disabled, a DHE key exchange will be performed when a PSK is accepted but only if requested by the client. If enabled, the server will require a DHE key exchange when a PSK is accepted regardless of whether the client supports combined PSK-DHE key exchange. This setting only has an effect when resumption is enabled. Possible values: [ YES, NO ]`,
				},
				resource.Attribute{
					Name:        "alpnprotocol",
					Description: `(Optional) Protocol to negotiate with client and then send as part of the ALPN extension in the server hello message. Possible values are HTTP1.1, HTTP2 and NONE. Default is none i.e. ALPN extension will not be sent. This parameter is relevant only if ssl connection is handled by the virtual server of type SSL_TCP. This parameter has no effect if TLSv1.3 is negotiated. Possible values: [ NONE, HTTP1.1, HTTP2 ]`,
				},
				resource.Attribute{
					Name:        "ciphername",
					Description: `(Optional) The cipher group/alias/individual cipher configuration.`,
				},
				resource.Attribute{
					Name:        "cipherpriority",
					Description: `(Optional) cipher priority.`,
				},
				resource.Attribute{
					Name:        "strictsigdigestcheck",
					Description: `(Optional) Parameter indicating to check whether peer entity certificate during TLS1.2 handshake is signed with one of signature-hash combination supported by Citrix ADC. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "ecccurvebindings",
					Description: `(Optional) A set of ECC curve names to be bound to this SSL profile.`,
				},
				resource.Attribute{
					Name:        "cipherbindings",
					Description: `(Optional) A set of ciphersuite bindings to be bound to this SSL profile. Documented below. A cipherbindings block supports the following:`,
				},
				resource.Attribute{
					Name:        "ciphername",
					Description: `(Required) Cipher name.`,
				},
				resource.Attribute{
					Name:        "cipherpriority",
					Description: `(Optional) This indicates priority assigned to the particular cipher. ## Attribute Reference In addition to the arguments, the following attributes are available:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of the sslprofile. It has the same value as the ` + "`" + `name` + "`" + ` attribute. ## Import A sslprofile can be imported using its name, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import citrixadc_sslprofile.tf_sslprofile tf_sslprofile ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id of the sslprofile. It has the same value as the ` + "`" + `name` + "`" + ` attribute. ## Import A sslprofile can be imported using its name, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import citrixadc_sslprofile.tf_sslprofile tf_sslprofile ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "citrixadc_sslprofile_sslcipher_binding",
			Category:         "SSL",
			ShortDescription: ``,
			Description: `

The sslprofile_sslcipher_binding resource is used to create bindings between sslprofiles and sslciphers.

~> If you are using this resource to bind sslciphers to a sslprofile
do not define the ` + "`" + `cipherbindings` + "`" + ` attribute in the sslprofile resource.


`,
			Keywords: []string{
				"ssl",
				"sslprofile",
				"sslcipher",
				"binding",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cipherpriority",
					Description: `(Optional) cipher priority.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the SSL profile.`,
				},
				resource.Attribute{
					Name:        "ciphername",
					Description: `(Optional) Name of the cipher. ## Attribute Reference In addition to the arguments, the following attributes are available:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of the sslprofile_sslcipher_binding. It has is the conatenation of the ` + "`" + `name` + "`" + ` and ` + "`" + `ciphername` + "`" + ` attributes. ## Import A sslprofile_sslcipher_binding can be imported using its id, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import citrixadc_sslprofile_sslcipher_binding.tf_binding tf_sslprofile,HIGH ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id of the sslprofile_sslcipher_binding. It has is the conatenation of the ` + "`" + `name` + "`" + ` and ` + "`" + `ciphername` + "`" + ` attributes. ## Import A sslprofile_sslcipher_binding can be imported using its id, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import citrixadc_sslprofile_sslcipher_binding.tf_binding tf_sslprofile,HIGH ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "citrixadc_sslvserver_sslcertkey_binding",
			Category:         "SSL",
			ShortDescription: ``,
			Description: `\_sslcertkey\_binding

The sslvserver\_sslcertkey\_binding resource is used to bind ssl certificates to SSL vservers.


`,
			Keywords: []string{
				"ssl",
				"sslvserver",
				"sslcertkey",
				"binding",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "certkeyname",
					Description: `(Required) The name of the certificate key pair binding.`,
				},
				resource.Attribute{
					Name:        "crlcheck",
					Description: `(Optional) The state of the CRL check parameter. (Mandatory/Optional). Possible values: [ Mandatory, Optional ]`,
				},
				resource.Attribute{
					Name:        "ocspcheck",
					Description: `(Optional) The state of the OCSP check parameter. (Mandatory/Optional). Possible values: [ Mandatory, Optional ]`,
				},
				resource.Attribute{
					Name:        "ca",
					Description: `(Optional) CA certificate.`,
				},
				resource.Attribute{
					Name:        "snicert",
					Description: `(Optional) The name of the CertKey. Use this option to bind Certkey(s) which will be used in SNI processing.`,
				},
				resource.Attribute{
					Name:        "skipcaname",
					Description: `(Optional) The flag is used to indicate whether this particular CA certificate's CA_Name needs to be sent to the SSL client while requesting for client certificate in a SSL handshake.`,
				},
				resource.Attribute{
					Name:        "vservername",
					Description: `(Required) Name of the SSL virtual server. ## Attribute Reference In addition to the arguments, the following attributes are available:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of the sslvserver\_sslcertkey\_binding. It is the concatenation of the ` + "`" + `vservername` + "`" + ` and ` + "`" + `certkeyname` + "`" + ` attributes separated by a comma. ## Import A sslvserver\_sslcertkey\_binding can be imported using its id, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import citrixadc_sslvserver_sslcertkey_binding.tf_binding tf_lbvserver,tf_sslcertkey ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id of the sslvserver\_sslcertkey\_binding. It is the concatenation of the ` + "`" + `vservername` + "`" + ` and ` + "`" + `certkeyname` + "`" + ` attributes separated by a comma. ## Import A sslvserver\_sslcertkey\_binding can be imported using its id, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import citrixadc_sslvserver_sslcertkey_binding.tf_binding tf_lbvserver,tf_sslcertkey ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "citrixadc_sslvserver_sslpolicy_binding",
			Category:         "SSL",
			ShortDescription: ``,
			Description: `

The sslvserver_sslpolicy_binding resource is used to create bindings between sslvservers and sslpolicies.


`,
			Keywords: []string{
				"ssl",
				"sslvserver",
				"sslpolicy",
				"binding",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "policyname",
					Description: `(Required) The name of the SSL policy binding.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `(Optional) The priority of the policies bound to this SSL service.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) Bind point to which to bind the policy. Possible Values: REQUEST, INTERCEPT_REQ and CLIENTHELLO_REQ. These bindpoints mean: 1. REQUEST: Policy evaluation will be done at appplication above SSL. This bindpoint is default and is used for actions based on clientauth and client cert. 2. INTERCEPT_REQ: Policy evaluation will be done during SSL handshake to decide whether to intercept or not. Actions allowed with this type are: INTERCEPT, BYPASS and RESET. 3. CLIENTHELLO_REQ: Policy evaluation will be done during handling of Client Hello Request. Action allowed with this type is: RESET, FORWARD and PICKCACERTGRP. Possible values: [ INTERCEPT_REQ, REQUEST, CLIENTHELLO_REQ ]`,
				},
				resource.Attribute{
					Name:        "gotopriorityexpression",
					Description: `(Optional) Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.`,
				},
				resource.Attribute{
					Name:        "invoke",
					Description: `(Optional) Invoke flag. This attribute is relevant only for ADVANCED policies.`,
				},
				resource.Attribute{
					Name:        "labeltype",
					Description: `(Optional) Type of policy label invocation. Possible values: [ vserver, service, policylabel ]`,
				},
				resource.Attribute{
					Name:        "labelname",
					Description: `(Optional) Name of the label to invoke if the current policy rule evaluates to TRUE.`,
				},
				resource.Attribute{
					Name:        "vservername",
					Description: `(Required) Name of the SSL virtual server. ## Attribute Reference In addition to the arguments, the following attributes are available:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of the sslvserver_sslpolicy_binding. It is the concatenation of the ` + "`" + `vservername` + "`" + ` and ` + "`" + `policyname` + "`" + ` attributes. ## Import A sslvserver_sslpolicy_binding can be imported using its id, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import citrixadc_sslvserver_sslpolicy_binding.tf_binding tf_lbvserver,tf_sslpolicy ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id of the sslvserver_sslpolicy_binding. It is the concatenation of the ` + "`" + `vservername` + "`" + ` and ` + "`" + `policyname` + "`" + ` attributes. ## Import A sslvserver_sslpolicy_binding can be imported using its id, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import citrixadc_sslvserver_sslpolicy_binding.tf_binding tf_lbvserver,tf_sslpolicy ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "citrixadc_systemcmdpolicy",
			Category:         "System",
			ShortDescription: ``,
			Description: `

The systemcmdpolicy resource is used to create command policies.


`,
			Keywords: []string{
				"system",
				"systemcmdpolicy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "policyname",
					Description: `(Optional) Name for a command policy.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Optional) Action to perform when a request matches the policy. Possible values: [ ALLOW, DENY ]`,
				},
				resource.Attribute{
					Name:        "cmdspec",
					Description: `(Optional) Regular expression specifying the data that matches the policy. ## Attribute Reference In addition to the arguments, the following attributes are available:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of the systemcmdpolicy. It has the same value as the ` + "`" + `policyname` + "`" + ` attribute. ## Import A systemcmdpolicy can be imported using its policyname, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import citrixadc_systemcmdpolicy.tf_policy tf_policy ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id of the systemcmdpolicy. It has the same value as the ` + "`" + `policyname` + "`" + ` attribute. ## Import A systemcmdpolicy can be imported using its policyname, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import citrixadc_systemcmdpolicy.tf_policy tf_policy ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "citrixadc_systemextramgmtcpu",
			Category:         "System",
			ShortDescription: ``,
			Description: `

The systemextramgmtcpu resource is used to enable or disable the extra management cpu on the target ADC.


`,
			Keywords: []string{
				"system",
				"systemextramgmtcpu",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "enabled",
					Description: `(Required) Boolean value indicating the desired state of the extra management CPU. Set to ` + "`" + `true` + "`" + ` to enable it.`,
				},
				resource.Attribute{
					Name:        "reboot",
					Description: `(Optional) Boolean value. Set to ` + "`" + `true` + "`" + ` to reboot after the application of the extra management cpu.`,
				},
				resource.Attribute{
					Name:        "reachable_timeout",
					Description: `(Optional) Total time to wait after reboot. Defaults to "10m".`,
				},
				resource.Attribute{
					Name:        "reachable_poll_delay",
					Description: `(Optional) Time to wait before the first poll after reboot. Defaults to "60s".`,
				},
				resource.Attribute{
					Name:        "reachable_poll_interval",
					Description: `(Optional) Interval between polls. Defaults to "60s".`,
				},
				resource.Attribute{
					Name:        "reachable_poll_timeout",
					Description: `(Optional) Timeout for a poll attempt. Default to "20s". ## Attribute Reference In addition to the arguments, the following attributes are available:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of the systemextramgmtcpu. It is a random string prefixed with ` + "`" + `tf-systemextramgmtcpu-` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id of the systemextramgmtcpu. It is a random string prefixed with ` + "`" + `tf-systemextramgmtcpu-` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "citrixadc_systemfile",
			Category:         "System",
			ShortDescription: ``,
			Description: `

The systemfile resource is used to upload files to the target ADC.


`,
			Keywords: []string{
				"system",
				"systemfile",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filename",
					Description: `(Optional) Name of the file. It should not include filepath.`,
				},
				resource.Attribute{
					Name:        "filecontent",
					Description: `(Optional) File contents.`,
				},
				resource.Attribute{
					Name:        "filelocation",
					Description: `(Optional) Location of the file on Citrix ADC.`,
				},
				resource.Attribute{
					Name:        "fileencoding",
					Description: `(Optional) Encoding type of the file content. Defaults to ` + "`" + `BASE64` + "`" + `. ## Attribute Reference In addition to the arguments, the following attributes are available:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of the systemfile. It is the fullpath of the system file. ## Import A systemfile can be imported using its full path, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import citrixadc_systemfile.tf_file /var/tmp/hello.txt ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id of the systemfile. It is the fullpath of the system file. ## Import A systemfile can be imported using its full path, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import citrixadc_systemfile.tf_file /var/tmp/hello.txt ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "citrixadc_systemgroup",
			Category:         "System",
			ShortDescription: ``,
			Description: `

The systemgroup resource is used to create user groups.


`,
			Keywords: []string{
				"system",
				"systemgroup",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "groupname",
					Description: `(Optional) Name for the group.`,
				},
				resource.Attribute{
					Name:        "promptstring",
					Description: `(Optional) String to display at the command-line prompt. Can consist of letters, numbers, hyphen (-), period (.), hash (#), space ( ), at (@), equal (=), colon (:), underscore (\_), and the following variables:`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `(Optional) CLI session inactivity timeout, in seconds. If Restrictedtimeout argument of system parameter is enabled, Timeout can have values in the range [300-86400] seconds.If Restrictedtimeout argument of system parameter is disabled, Timeout can have values in the range [0, 10-100000000] seconds. Default value is 900 seconds.`,
				},
				resource.Attribute{
					Name:        "systemusers",
					Description: `(Optional) A set of user names to bind to this group.`,
				},
				resource.Attribute{
					Name:        "cmdpolicybinding",
					Description: `(Optional) A set of command policies to bing to this group. Attributes are detailed below In a command policy block the following attributes are allowed:`,
				},
				resource.Attribute{
					Name:        "policyname",
					Description: `(Optional) Name of the policy to bind.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `(Optional) Priority for the biding. ## Attribute Reference In addition to the arguments, the following attributes are available:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of the systemgroup. It has the same value as the ` + "`" + `groupname` + "`" + ` attribute. ## Import A systemgroup can be imported using its groupname, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import citrixadc_systemgroup.tf_systemgroup tf_systemgroup ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id of the systemgroup. It has the same value as the ` + "`" + `groupname` + "`" + ` attribute. ## Import A systemgroup can be imported using its groupname, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import citrixadc_systemgroup.tf_systemgroup tf_systemgroup ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "citrixadc_systemuser",
			Category:         "System",
			ShortDescription: ``,
			Description: `

The systemuser resource is used to create users for the target ADC.


`,
			Keywords: []string{
				"system",
				"systemuser",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "username",
					Description: `(Optional) Name for a user.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional) Password for the system user. Can include any ASCII character.`,
				},
				resource.Attribute{
					Name:        "externalauth",
					Description: `(Optional) Whether to use external authentication servers for the system user authentication or not. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "promptstring",
					Description: `(Optional) String to display at the command-line prompt. Can consist of letters, numbers, hyphen (-), period (.), hash (#), space ( ), at (@), equal (=), colon (:), underscore (\_), and the following variables:`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `(Optional) CLI session inactivity timeout, in seconds. If Restrictedtimeout argument of system parameter is enabled, Timeout can have values in the range [300-86400] seconds. If Restrictedtimeout argument of system parameter is disabled, Timeout can have values in the range [0, 10-100000000] seconds. Default value is 900 seconds.`,
				},
				resource.Attribute{
					Name:        "logging",
					Description: `(Optional) Users logging privilege. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "maxsession",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "hashedpassword",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "cmdpolicybinding",
					Description: `(Optional) A set of blocks binding systemcommandpolicies to the systemuser. See below for details. A ` + "`" + `cmdpolicybinding` + "`" + ` block can contain the following attributes`,
				},
				resource.Attribute{
					Name:        "policyname",
					Description: `(Optional) The name of command policy.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `(Optional) The priority of the policy. ## Attribute Reference In addition to the arguments, the following attributes are available:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of the systemuser. It has the same value as the ` + "`" + `username` + "`" + ` attribute. ## Import A systemuser can be imported using its name, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import citrixadc_systemuser.tf_user tf_user ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id of the systemuser. It has the same value as the ` + "`" + `username` + "`" + ` attribute. ## Import A systemuser can be imported using its name, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import citrixadc_systemuser.tf_user tf_user ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "citrixadc_transformaction",
			Category:         "Transform",
			ShortDescription: ``,
			Description: `

The transformaction resource is used to create transform actions.


`,
			Keywords: []string{
				"transform",
				"transformaction",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name for the URL transformation action. Must begin with a letter, number, or the underscore character (\_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after the URL Transformation action is added.`,
				},
				resource.Attribute{
					Name:        "profilename",
					Description: `(Optional) Name of the URL Transformation profile with which to associate this action.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `(Optional) Positive integer specifying the priority of the action within the profile. A lower number specifies a higher priority. Must be unique within the list of actions bound to the profile. Policies are evaluated in the order of their priority numbers, and the first policy that matches is applied.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `(Optional) Enable or disable this action. Possible values: [ ENABLED, DISABLED ]`,
				},
				resource.Attribute{
					Name:        "requrlfrom",
					Description: `(Optional) PCRE-format regular expression that describes the request URL pattern to be transformed.`,
				},
				resource.Attribute{
					Name:        "requrlinto",
					Description: `(Optional) PCRE-format regular expression that describes the transformation to be performed on URLs that match the reqUrlFrom pattern.`,
				},
				resource.Attribute{
					Name:        "resurlfrom",
					Description: `(Optional) PCRE-format regular expression that describes the response URL pattern to be transformed.`,
				},
				resource.Attribute{
					Name:        "resurlinto",
					Description: `(Optional) PCRE-format regular expression that describes the transformation to be performed on URLs that match the resUrlFrom pattern.`,
				},
				resource.Attribute{
					Name:        "cookiedomainfrom",
					Description: `(Optional) Pattern that matches the domain to be transformed in Set-Cookie headers.`,
				},
				resource.Attribute{
					Name:        "cookiedomaininto",
					Description: `(Optional) PCRE-format regular expression that describes the transformation to be performed on cookie domains that match the cookieDomainFrom pattern. NOTE: The cookie domain to be transformed is extracted from the request.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `(Optional) Any comments to preserve information about this URL Transformation action. ## Attribute Reference In addition to the arguments, the following attributes are available:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of the transformaction. It has the same value as the ` + "`" + `name` + "`" + ` attribute. ## Import A transformaction can be imported using its name, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import citrixadc_transformaction.tf_trans_action tf_trans_action ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id of the transformaction. It has the same value as the ` + "`" + `name` + "`" + ` attribute. ## Import A transformaction can be imported using its name, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import citrixadc_transformaction.tf_trans_action tf_trans_action ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "citrixadc_transformpolicy",
			Category:         "Transform",
			ShortDescription: ``,
			Description: `

The transformpolicy resource is used to create transform policies


`,
			Keywords: []string{
				"transform",
				"transformpolicy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name for the URL Transformation policy. Must begin with a letter, number, or the underscore character (\_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Can be changed after the URL Transformation policy is added.`,
				},
				resource.Attribute{
					Name:        "rule",
					Description: `(Optional) Expression, or name of a named expression, against which to evaluate traffic. The following requirements apply only to the Citrix ADC CLI:`,
				},
				resource.Attribute{
					Name:        "profilename",
					Description: `(Optional) Name of the URL Transformation profile to use to transform requests and responses that match the policy.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `(Optional) Any comments to preserve information about this URL Transformation policy.`,
				},
				resource.Attribute{
					Name:        "logaction",
					Description: `(Optional) Log server to use to log connections that match this policy. ## Attribute Reference In addition to the arguments, the following attributes are available:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of the transformpolicy. It has the same value as the ` + "`" + `name` + "`" + ` attribute. ## Import A transformpolicy can be imported using its name, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import citrixadc_transformpolicy.tf_trans_policy tf_trans_policy ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id of the transformpolicy. It has the same value as the ` + "`" + `name` + "`" + ` attribute. ## Import A transformpolicy can be imported using its name, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import citrixadc_transformpolicy.tf_trans_policy tf_trans_policy ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "citrixadc_transformprofile",
			Category:         "Transform",
			ShortDescription: ``,
			Description: `

The transformprofile resource is used to create transform profiles.


`,
			Keywords: []string{
				"transform",
				"transformprofile",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name for the URL transformation profile. Must begin with a letter, number, or the underscore character (\_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after the URL transformation profile is added.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) Type of transformation. Always URL for URL Transformation profiles. Possible values: [ URL ]`,
				},
				resource.Attribute{
					Name:        "onlytransformabsurlinbody",
					Description: `(Optional) In the HTTP body, transform only absolute URLs. Relative URLs are ignored. Possible values: [ on, off ]`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `(Optional) Any comments to preserve information about this URL Transformation profile. ## Attribute Reference In addition to the arguments, the following attributes are available:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of the transformprofile. It has the same value as the ` + "`" + `name` + "`" + ` attribute. ## Import A transformprofile can be imported using its name, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import citrixadc_transformprofile.tf_trans_profile tf_trans_profile ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id of the transformprofile. It has the same value as the ` + "`" + `name` + "`" + ` attribute. ## Import A transformprofile can be imported using its name, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import citrixadc_transformprofile.tf_trans_profile tf_trans_profile ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
	}

	resourcesMap = map[string]int{

		"citrixadc_appfwfieldtype":                          0,
		"citrixadc_appfwjsoncontenttype":                    1,
		"citrixadc_appfwpolicy":                             2,
		"citrixadc_appfwpolicylabel":                        3,
		"citrixadc_appfwprofile":                            4,
		"citrixadc_appfwprofile_cookieconsistency_binding":  5,
		"citrixadc_appfwprofile_crosssitescripting_binding": 6,
		"citrixadc_appfwprofile_denyurl_binding":            7,
		"citrixadc_appfwprofile_sqlinjection_binding":       8,
		"citrixadc_appfwprofile_starturl_binding":           9,
		"citrixadc_appfwxmlcontenttype":                     10,
		"citrixadc_auditmessageaction":                      11,
		"citrixadc_auditsyslogaction":                       12,
		"citrixadc_auditsyslogpolicy":                       13,
		"citrixadc_cluster":                                 14,
		"citrixadc_clusterfiles_syncer":                     15,
		"citrixadc_cmppolicy":                               16,
		"citrixadc_csaction":                                17,
		"citrixadc_cspolicy":                                18,
		"citrixadc_csvserver":                               19,
		"citrixadc_csvserver_appfwpolicy_binding":           20,
		"citrixadc_csvserver_cmppolicy_binding":             21,
		"citrixadc_csvserver_cspolicy_binding":              22,
		"citrixadc_csvserver_filterpolicy_binding":          23,
		"citrixadc_csvserver_responderpolicy_binding":       24,
		"citrixadc_csvserver_rewritepolicy_binding":         25,
		"citrixadc_csvserver_transformpolicy_binding":       26,
		"citrixadc_dnsnsrec":                                27,
		"citrixadc_dnssoarec":                               28,
		"citrixadc_filterpolicy":                            29,
		"citrixadc_gslbservice":                             30,
		"citrixadc_gslbsite":                                31,
		"citrixadc_gslbvserver":                             32,
		"citrixadc_inat":                                    33,
		"citrixadc_installer":                               34,
		"citrixadc_interface":                               35,
		"citrixadc_ipset":                                   36,
		"citrixadc_lbmonitor":                               37,
		"citrixadc_lbvserver":                               38,
		"citrixadc_lbvserver_appfwpolicy_binding":           39,
		"citrixadc_lbvserver_cmppolicy_binding":             40,
		"citrixadc_lbvserver_filterpolicy_binding":          41,
		"citrixadc_lbvserver_responderpolicy_binding":       42,
		"citrixadc_lbvserver_rewritepolicy_binding":         43,
		"citrixadc_lbvserver_service_binding":               44,
		"citrixadc_lbvserver_transformpolicy_binding":       45,
		"citrixadc_linkset":                                 46,
		"citrixadc_netprofile":                              47,
		"citrixadc_nsacl":                                   48,
		"citrixadc_nsacls":                                  49,
		"citrixadc_nscapacity":                              50,
		"citrixadc_nsconfig_clear":                          51,
		"citrixadc_nsconfig_save":                           52,
		"citrixadc_nsconfig_update":                         53,
		"citrixadc_nsfeature":                               54,
		"citrixadc_nshttpprofile":                           55,
		"citrixadc_nsip":                                    56,
		"citrixadc_nsip6":                                   57,
		"citrixadc_nslicense":                               58,
		"citrixadc_nslicenseserver":                         59,
		"citrixadc_nsparam":                                 60,
		"citrixadc_nsrpcnode":                               61,
		"citrixadc_nstcpparam":                              62,
		"citrixadc_nstcpprofile":                            63,
		"citrixadc_nsvpxparam":                              64,
		"citrixadc_password_resetter":                       65,
		"citrixadc_pinger":                                  66,
		"citrixadc_policydataset":                           67,
		"citrixadc_policydataset_value_binding":             68,
		"citrixadc_policyexpression":                        69,
		"citrixadc_policypatset":                            70,
		"citrixadc_policypatset_pattern_binding":            71,
		"citrixadc_policystringmap":                         72,
		"citrixadc_policystringmap_pattern_binding":         73,
		"citrixadc_quicbridgeprofile":                       74,
		"citrixadc_rebooter":                                75,
		"citrixadc_responderaction":                         76,
		"citrixadc_responderpolicy":                         77,
		"citrixadc_responderpolicylabel":                    78,
		"citrixadc_rewriteaction":                           79,
		"citrixadc_rewritepolicy":                           80,
		"citrixadc_rewritepolicylabel":                      81,
		"citrixadc_rnat":                                    82,
		"citrixadc_route":                                   83,
		"citrixadc_routerdynamicrouting":                    84,
		"citrixadc_server":                                  85,
		"citrixadc_service":                                 86,
		"citrixadc_servicegroup":                            87,
		"citrixadc_servicegroup_lbmonitor_binding":          88,
		"citrixadc_servicegroup_servicegroupmember_binding": 89,
		"citrixadc_sslaction":                               90,
		"citrixadc_sslcertkey":                              91,
		"citrixadc_sslcipher":                               92,
		"citrixadc_ssldhparam":                              93,
		"citrixadc_sslparameter":                            94,
		"citrixadc_sslpolicy":                               95,
		"citrixadc_sslprofile":                              96,
		"citrixadc_sslprofile_sslcipher_binding":            97,
		"citrixadc_sslvserver_sslcertkey_binding":           98,
		"citrixadc_sslvserver_sslpolicy_binding":            99,
		"citrixadc_systemcmdpolicy":                         100,
		"citrixadc_systemextramgmtcpu":                      101,
		"citrixadc_systemfile":                              102,
		"citrixadc_systemgroup":                             103,
		"citrixadc_systemuser":                              104,
		"citrixadc_transformaction":                         105,
		"citrixadc_transformpolicy":                         106,
		"citrixadc_transformprofile":                        107,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
