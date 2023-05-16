package couchbasecapella

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "couchbasecapella_bucket",
			Category:         "Resources",
			ShortDescription: `Create, edit and delete Buckets in a Couchbase Capella Cluster.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "couchbasecapella_database_user",
			Category:         "Resources",
			ShortDescription: `Create, edit and delete Database Users for a Couchbase Capella Cluster.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "couchbasecapella_hosted_cluster",
			Category:         "Resources",
			ShortDescription: `Create, edit and delete Hosted Clusters in Couchbase Capella.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "couchbasecapella_project",
			Category:         "Resources",
			ShortDescription: `Create, edit and delete Projects in Couchbase Capella.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "couchbasecapella_vpc_cluster",
			Category:         "Resources",
			ShortDescription: `Create and delete VPC Clusters in Couchbase Capella.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"couchbasecapella_bucket":         0,
		"couchbasecapella_database_user":  1,
		"couchbasecapella_hosted_cluster": 2,
		"couchbasecapella_project":        3,
		"couchbasecapella_vpc_cluster":    4,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
