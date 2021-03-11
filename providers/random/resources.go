package random

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "random_id",
			Category:         "Resources",
			ShortDescription: `The resource random_id generates random numbers that are intended to be used as unique identifiers for other resources. This resource does use a cryptographic random number generator in order to minimize the chance of collisions, making the results of this resource when a 16-byte identifier is requested of equivalent uniqueness to a type-4 UUID. This resource can be used in conjunction with resources that have the create_before_destroy lifecycle flag set to avoid conflicts with unique names during the brief period where both the old and new resources exist concurrently.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "random_integer",
			Category:         "Resources",
			ShortDescription: `The resource random_integer generates random values from a given range, described by the min and max attributes of a given resource. This resource can be used in conjunction with resources that have the create_before_destroy lifecycle flag set, to avoid conflicts with unique names during the brief period where both the old and new resources exist concurrently.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "random_password",
			Category:         "Resources",
			ShortDescription: `Identical to random_string string.html with the exception that the result is treated as sensitive and, thus, not displayed in console output. Read more about sensitive data handling in the Terraform documentation https://www.terraform.io/docs/language/state/sensitive-data.html. This resource does use a cryptographic random number generator.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "random_pet",
			Category:         "Resources",
			ShortDescription: `The resource random_pet generates random pet names that are intended to be used as unique identifiers for other resources. This resource can be used in conjunction with resources that have the create_before_destroy lifecycle flag set, to avoid conflicts with unique names during the brief period where both the old and new resources exist concurrently.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "random_shuffle",
			Category:         "Resources",
			ShortDescription: `The resource random_shuffle generates a random permutation of a list of strings given as an argument.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "random_string",
			Category:         "Resources",
			ShortDescription: `The resource random_string generates a random permutation of alphanumeric characters and optionally special characters. This resource does use a cryptographic random number generator. Historically this resource's intended usage has been ambiguous as the original example used it in a password. For backwards compatibility it will continue to exist. For unique ids please use random_id id.html, for sensitive random values please use random_password password.html.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "random_uuid",
			Category:         "Resources",
			ShortDescription: `The resource random_uuid generates random uuid string that is intended to be used as unique identifiers for other resources. This resource uses hashicorp/go-uuid https://github.com/hashicorp/go-uuid to generate a UUID-formatted string for use with services needed a unique string identifier.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"random_id":       0,
		"random_integer":  1,
		"random_password": 2,
		"random_pet":      3,
		"random_shuffle":  4,
		"random_string":   5,
		"random_uuid":     6,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
