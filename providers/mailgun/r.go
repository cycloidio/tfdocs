package mailgun

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "mailgun_domain",
			Category:         "Resources",
			ShortDescription: `Provides a Mailgun App resource. This can be used to create and manage applications on Mailgun.`,
			Description:      ``,
			Keywords: []string{
				"domain",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The domain to add to Mailgun`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region where domain will be created. Default value is ` + "`" + `us` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "smtp_password",
					Description: `(Optional) Password for SMTP authentication`,
				},
				resource.Attribute{
					Name:        "spam_action",
					Description: `(Optional) ` + "`" + `disabled` + "`" + ` or ` + "`" + `tag` + "`" + ` Disable, no spam filtering will occur for inbound messages. Tag, messages will be tagged with a spam header.`,
				},
				resource.Attribute{
					Name:        "wildcard",
					Description: `(Optional) Boolean that determines whether the domain will accept email for sub-domains. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the domain.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The name of the region.`,
				},
				resource.Attribute{
					Name:        "smtp_login",
					Description: `The login email for the SMTP server.`,
				},
				resource.Attribute{
					Name:        "smtp_password",
					Description: `The password to the SMTP server.`,
				},
				resource.Attribute{
					Name:        "wildcard",
					Description: `Whether or not the domain will accept email for sub-domains.`,
				},
				resource.Attribute{
					Name:        "spam_action",
					Description: `The spam filtering setting.`,
				},
				resource.Attribute{
					Name:        "receiving_records",
					Description: `A list of DNS records for receiving validation.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `The priority of the record.`,
				},
				resource.Attribute{
					Name:        "record_type",
					Description: `The record type.`,
				},
				resource.Attribute{
					Name:        "valid",
					Description: `` + "`" + `"valid"` + "`" + ` if the record is valid.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the record.`,
				},
				resource.Attribute{
					Name:        "sending_records",
					Description: `A list of DNS records for sending validation.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the record.`,
				},
				resource.Attribute{
					Name:        "record_type",
					Description: `The record type.`,
				},
				resource.Attribute{
					Name:        "valid",
					Description: `` + "`" + `"valid"` + "`" + ` if the record is valid.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the record. ## Import Domains can be imported using ` + "`" + `region:domain_name` + "`" + ` via ` + "`" + `import` + "`" + ` command. Region has to be chosen from ` + "`" + `eu` + "`" + ` or ` + "`" + `us` + "`" + ` (when no selection ` + "`" + `us` + "`" + ` is applied). ` + "`" + `` + "`" + `` + "`" + `hcl terraform import mailgun_domain.test us:example.domain.com ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `The name of the domain.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The name of the region.`,
				},
				resource.Attribute{
					Name:        "smtp_login",
					Description: `The login email for the SMTP server.`,
				},
				resource.Attribute{
					Name:        "smtp_password",
					Description: `The password to the SMTP server.`,
				},
				resource.Attribute{
					Name:        "wildcard",
					Description: `Whether or not the domain will accept email for sub-domains.`,
				},
				resource.Attribute{
					Name:        "spam_action",
					Description: `The spam filtering setting.`,
				},
				resource.Attribute{
					Name:        "receiving_records",
					Description: `A list of DNS records for receiving validation.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `The priority of the record.`,
				},
				resource.Attribute{
					Name:        "record_type",
					Description: `The record type.`,
				},
				resource.Attribute{
					Name:        "valid",
					Description: `` + "`" + `"valid"` + "`" + ` if the record is valid.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the record.`,
				},
				resource.Attribute{
					Name:        "sending_records",
					Description: `A list of DNS records for sending validation.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the record.`,
				},
				resource.Attribute{
					Name:        "record_type",
					Description: `The record type.`,
				},
				resource.Attribute{
					Name:        "valid",
					Description: `` + "`" + `"valid"` + "`" + ` if the record is valid.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the record. ## Import Domains can be imported using ` + "`" + `region:domain_name` + "`" + ` via ` + "`" + `import` + "`" + ` command. Region has to be chosen from ` + "`" + `eu` + "`" + ` or ` + "`" + `us` + "`" + ` (when no selection ` + "`" + `us` + "`" + ` is applied). ` + "`" + `` + "`" + `` + "`" + `hcl terraform import mailgun_domain.test us:example.domain.com ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mailgun_route",
			Category:         "Resources",
			ShortDescription: `Provides a Mailgun App resource. This can be used to create and manage applications on Mailgun.`,
			Description:      ``,
			Keywords: []string{
				"route",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "priority",
					Description: `(Required) Smaller number indicates higher priority. Higher priority routes are handled first.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "expression",
					Description: `(Required) A filter expression like ` + "`" + `match_recipient('.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Required) Route action. This action is executed when the expression evaluates to True. Example: ` + "`" + `forward("alice@example.com")` + "`" + ` You can pass multiple ` + "`" + `action` + "`" + ` parameters.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region where domain will be created. Default value is ` + "`" + `us` + "`" + `. ## Import Routes can be imported using ` + "`" + `ROUTE_ID` + "`" + ` and ` + "`" + `region` + "`" + ` via ` + "`" + `import` + "`" + ` command. Route ID can be found on Mailgun portal in section ` + "`" + `Receiving/Routes` + "`" + `. Region has to be chosen from ` + "`" + `eu` + "`" + ` or ` + "`" + `us` + "`" + ` (when no selection ` + "`" + `us` + "`" + ` is applied). ` + "`" + `` + "`" + `` + "`" + `hcl terraform import mailgun_route.test eu:123456789 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"mailgun_domain": 0,
		"mailgun_route":  1,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
