package influxdb

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "influxdb_continuous_query",
			Category:         "Resources",
			ShortDescription: `The influxdb_continuous_query resource allows an InfluxDB continuous query to be managed.`,
			Description:      ``,
			Keywords: []string{
				"continuous",
				"query",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name for the continuous_query. This must be unique on the InfluxDB server.`,
				},
				resource.Attribute{
					Name:        "database",
					Description: `(Required) The database for the continuous_query. This must be an existing influxdb database.`,
				},
				resource.Attribute{
					Name:        "query",
					Description: `(Required) The query for the continuous_query.`,
				},
				resource.Attribute{
					Name:        "resample",
					Description: `(Optional) The body of the query's RESAMPLE clause. The format is detailed in the InfluxDB documentation. ## Attributes Reference This resource exports no further attributes.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "influxdb_database",
			Category:         "Resources",
			ShortDescription: `The influxdb_database resource allows an InfluxDB database to be created.`,
			Description:      ``,
			Keywords: []string{
				"database",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name for the database. This must be unique on the InfluxDB server.`,
				},
				resource.Attribute{
					Name:        "retention_policies",
					Description: `(Optional) A list of retention policies for specified database Each ` + "`" + `retention_policies` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the retention policy`,
				},
				resource.Attribute{
					Name:        "duration",
					Description: `(Required) The duration for retention policy, format of duration can be found at InfluxDB Documentation.`,
				},
				resource.Attribute{
					Name:        "replication",
					Description: `(Optional) Determines how many copies of data points are stored in a cluster. Not applicable for single node / Open Source version of InfluxDB. Default value of 1.`,
				},
				resource.Attribute{
					Name:        "shardgroupduration",
					Description: `(Optional) Determines how much time each shard group spans. How and why to modify can be found at InfluxDB Documentation.`,
				},
				resource.Attribute{
					Name:        "default",
					Description: `(Optional) Marks current retention policy as default. Default value is false. ## Attributes Reference This resource exports no further attributes.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "influxdb_user",
			Category:         "Resources",
			ShortDescription: `The influxdb_user resource allows an InfluxDB users to be managed.`,
			Description:      ``,
			Keywords: []string{
				"user",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name for the user.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Required) The password for the user.`,
				},
				resource.Attribute{
					Name:        "admin",
					Description: `(Optional) Mark the user as admin.`,
				},
				resource.Attribute{
					Name:        "grant",
					Description: `(Optional) A list of grants for non-admin users Each ` + "`" + `grant` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "database",
					Description: `(Required) The name of the database the privilege is associated with`,
				},
				resource.Attribute{
					Name:        "privilege",
					Description: `(Required) The privilege to grant (READ|WRITE|ALL) ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "admin",
					Description: `(Bool) indication if the user is an admin or not.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "admin",
					Description: `(Bool) indication if the user is an admin or not.`,
				},
			},
		},
	}

	resourcesMap = map[string]int{

		"influxdb_continuous_query": 0,
		"influxdb_database":         1,
		"influxdb_user":             2,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
