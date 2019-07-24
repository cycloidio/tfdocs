package aws

import (
	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*Resource{

		&resource.Resource{
			Name:             "",
			Type:             "dme_record",
			Category:         "Resources",
			ShortDescription: `Provides a DNSMadeEasy record resource.`,
			Description:      ``,
			Keywords: []string{
				"record",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "domainid",
					Description: `(String, Required) The domain id to add the record to`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the record ` + "`" + `type` + "`" + ` - (Required) The type of`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `(Integer, Optional) The TTL of the record ` + "`" + `gtdLocation` + "`" + ` - (String, Optional) The GTD Location of the record on Global Traffic Director enabled domains; Unless GTD is enabled this should either be omitted or set to "DEFAULT" Additional arguments are listed below under DNS Record Types. ## DNS Record Types The type of record being created affects the interpretation of the ` + "`" + `value` + "`" + ` argument; also, some additional arguments are required for some record types. http://help.dnsmadeeasy.com/tutorials/managed-dns/ has more information. #### A Record`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the record`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of the record`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the record ` + "`" + `type` + "`" + ` (see below)`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `The TTL of the record`,
				},
				resource.Attribute{
					Name:        "gtdLocation",
					Description: `The GTD Location of the record on GTD enabled domains Additional fields may also be exported by some record types - see DNS Record Types. #### Record Type Examples Following are examples of using each of the record types. ` + "`" + `` + "`" + `` + "`" + `hcl # Provide your API and Secret Keys, and whether the sandbox # is being used (defaults to false) provider "dme" { akey = "aaaaaa1a-11a1-1aa1-a101-11a1a11aa1aa" skey = "11a0a11a-a1a1-111a-a11a-a11110a11111" usesandbox = true } # A Record resource "dme_record" "testa" { domainid = "123456" name = "testa" type = "A" value = "1.1.1.1" ttl = 1000 gtdLocation = "DEFAULT" } # CNAME record resource "dme_record" "testcname" { domainid = "123456" name = "testcname" type = "CNAME" value = "foo" ttl = 1000 } # ANAME record resource "dme_record" "testaname" { domainid = "123456" name = "testaname" type = "ANAME" value = "foo" ttl = 1000 } # MX record resource "dme_record" "testmx" { domainid = "123456" name = "testmx" type = "MX" value = "foo" mxLevel = 10 ttl = 1000 } # HTTPRED resource "dme_record" "testhttpred" { domainid = "123456" name = "testhttpred" type = "HTTPRED" value = "https://github.com/soniah/terraform-provider-dme" hardLink = true redirectType = "Hidden Frame Masked" title = "An Example" keywords = "terraform example" description = "This is a description" ttl = 2000 } # TXT record resource "dme_record" "testtxt" { domainid = "123456" name = "testtxt" type = "TXT" value = "foo" ttl = 1000 } # SPF record resource "dme_record" "testspf" { domainid = "123456" name = "testspf" type = "SPF" value = "foo" ttl = 1000 } # PTR record resource "dme_record" "testptr" { domainid = "123456" name = "testptr" type = "PTR" value = "foo" ttl = 1000 } # NS record resource "dme_record" "testns" { domainid = "123456" name = "testns" type = "NS" value = "foo" ttl = 1000 } # AAAA record resource "dme_record" "testaaaa" { domainid = "123456" name = "testaaaa" type = "AAAA" value = "FE80::0202:B3FF:FE1E:8329" ttl = 1000 } # SRV record resource "dme_record" "testsrv" { domainid = "123456" name = "testsrv" type = "SRV" value = "foo" priority = 10 weight = 20 port = 30 ttl = 1000 } ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `The name of the record`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of the record`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the record ` + "`" + `type` + "`" + ` (see below)`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `The TTL of the record`,
				},
				resource.Attribute{
					Name:        "gtdLocation",
					Description: `The GTD Location of the record on GTD enabled domains Additional fields may also be exported by some record types - see DNS Record Types. #### Record Type Examples Following are examples of using each of the record types. ` + "`" + `` + "`" + `` + "`" + `hcl # Provide your API and Secret Keys, and whether the sandbox # is being used (defaults to false) provider "dme" { akey = "aaaaaa1a-11a1-1aa1-a101-11a1a11aa1aa" skey = "11a0a11a-a1a1-111a-a11a-a11110a11111" usesandbox = true } # A Record resource "dme_record" "testa" { domainid = "123456" name = "testa" type = "A" value = "1.1.1.1" ttl = 1000 gtdLocation = "DEFAULT" } # CNAME record resource "dme_record" "testcname" { domainid = "123456" name = "testcname" type = "CNAME" value = "foo" ttl = 1000 } # ANAME record resource "dme_record" "testaname" { domainid = "123456" name = "testaname" type = "ANAME" value = "foo" ttl = 1000 } # MX record resource "dme_record" "testmx" { domainid = "123456" name = "testmx" type = "MX" value = "foo" mxLevel = 10 ttl = 1000 } # HTTPRED resource "dme_record" "testhttpred" { domainid = "123456" name = "testhttpred" type = "HTTPRED" value = "https://github.com/soniah/terraform-provider-dme" hardLink = true redirectType = "Hidden Frame Masked" title = "An Example" keywords = "terraform example" description = "This is a description" ttl = 2000 } # TXT record resource "dme_record" "testtxt" { domainid = "123456" name = "testtxt" type = "TXT" value = "foo" ttl = 1000 } # SPF record resource "dme_record" "testspf" { domainid = "123456" name = "testspf" type = "SPF" value = "foo" ttl = 1000 } # PTR record resource "dme_record" "testptr" { domainid = "123456" name = "testptr" type = "PTR" value = "foo" ttl = 1000 } # NS record resource "dme_record" "testns" { domainid = "123456" name = "testns" type = "NS" value = "foo" ttl = 1000 } # AAAA record resource "dme_record" "testaaaa" { domainid = "123456" name = "testaaaa" type = "AAAA" value = "FE80::0202:B3FF:FE1E:8329" ttl = 1000 } # SRV record resource "dme_record" "testsrv" { domainid = "123456" name = "testsrv" type = "SRV" value = "foo" priority = 10 weight = 20 port = 30 ttl = 1000 } ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
	}

	resourcesMap = map[string]int{

		"dme_record": 0,
	}
)

func GetResource(r string) (*resouce.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs]
}
