package project

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "project_project",
			Category:         "Resources",
			ShortDescription: `Provides an Artifactory project resource. This can be used to create and manage Artifactory project, maintain users/groups/roles/repos. Repository Configuration After the project configuration is applied, the repository's attributes project_key and project_environments would be updated with the project's data. This will generate a state drift in the next Terraform plan/apply for the repository resource. To avoid this, apply lifecycle.ignore_changes: ` + "`" + `` + "`" + `` + "`" + `hcl resource "artifactorylocalmavenrepository" "mymaven_releases" { key = "my-maven-releases" ... lifecycle { ignore_changes = [ project_environments, project_key ] } } ` + "`" + `` + "`" + ` ~>We strongly recommend using the repos attribute to manage the list of repositories. See below for additional details.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"project_project": 0,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
