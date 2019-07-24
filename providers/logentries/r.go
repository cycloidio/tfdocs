package aws

import (
	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*Resource{

		&resource.Resource{
			Name:             "",
			Type:             "logentries_log",
			Category:         "Resources",
			ShortDescription: `Creates a Logentries log.`,
			Description:      ``,
			Keywords: []string{
				"log",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "logset_id",
					Description: `(Required) The id of the ` + "`" + `logentries_logset` + "`" + ` resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the log. The name should be short and descriptive. For example, Apache Access, Hadoop Namenode.`,
				},
				resource.Attribute{
					Name:        "filename",
					Description: `(Optional) the filename of the log.`,
				},
				resource.Attribute{
					Name:        "retention_period",
					Description: `(Optional, default ` + "`" + `ACCOUNT_DEFAULT` + "`" + `) The retention period (` + "`" + `1W` + "`" + `, ` + "`" + `2W` + "`" + `, ` + "`" + `1M` + "`" + `, ` + "`" + `2M` + "`" + `, ` + "`" + `6M` + "`" + `, ` + "`" + `1Y` + "`" + `, ` + "`" + `2Y` + "`" + `, ` + "`" + `UNLIMITED` + "`" + `, ` + "`" + `ACCOUNT_DEFAULT` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "source",
					Description: `(Optional, default ` + "`" + `token` + "`" + `) The log source (` + "`" + `token` + "`" + `, ` + "`" + `syslog` + "`" + `, ` + "`" + `agent` + "`" + `, ` + "`" + `api` + "`" + `). Review the Logentries [log inputs documentation](https://docs.logentries.com/docs/) for more information.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) The log type. See the Logentries [log type documentation](https://logentries.com/doc/log-types/) for more information. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "token",
					Description: `If the log ` + "`" + `source` + "`" + ` is ` + "`" + `token` + "`" + `, this value holds the generated log token that is used by logging clients. See the Logentries [token-based input documentation](https://logentries.com/doc/input-token/) for more information.`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "token",
					Description: `If the log ` + "`" + `source` + "`" + ` is ` + "`" + `token` + "`" + `, this value holds the generated log token that is used by logging clients. See the Logentries [token-based input documentation](https://logentries.com/doc/input-token/) for more information.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "logentries_logset",
			Category:         "Resources",
			ShortDescription: `Creates a Logentries logset.`,
			Description:      ``,
			Keywords: []string{
				"logset",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The log set name, which should be short and descriptive. For example, www, db1.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `(Optional, default "nonlocation") A location is for your convenience only. You can specify a DNS entry such as web.example.com, IP address or arbitrary comment.`,
				},
			},
			Attributes: []resource.Argument{},
		},
	}

	resourcesMap = map[string]int{

		"logentries_log":    0,
		"logentries_logset": 1,
	}
)

func GetResource(r string) (*resouce.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs]
}
