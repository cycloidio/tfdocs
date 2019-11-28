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
			ShortDescription: `Generates a random identifier.`,
			Description:      ``,
			Keywords: []string{
				"id",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "byte_length",
					Description: `(Required) The number of random bytes to produce. The minimum value is 1, which produces eight bits of randomness.`,
				},
				resource.Attribute{
					Name:        "keepers",
					Description: `(Optional) Arbitrary map of values that, when changed, will trigger a new id to be generated. See [the main provider documentation](../index.html) for more information.`,
				},
				resource.Attribute{
					Name:        "prefix",
					Description: `(Optional) Arbitrary string to prefix the output value with. This string is supplied as-is, meaning it is not guaranteed to be URL-safe or base64 encoded. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "b64_url",
					Description: `The generated id presented in base64, using the URL-friendly character set: case-sensitive letters, digits and the characters ` + "`" + `_` + "`" + ` and ` + "`" + `-` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "b64_std",
					Description: `The generated id presented in base64 without additional transformations.`,
				},
				resource.Attribute{
					Name:        "hex",
					Description: `The generated id presented in padded hexadecimal digits. This result will always be twice as long as the requested byte length.`,
				},
				resource.Attribute{
					Name:        "dec",
					Description: `The generated id presented in non-padded decimal digits. ## Import Random Ids can be imported using the ` + "`" + `b64_url` + "`" + ` with an optional ` + "`" + `prefix` + "`" + `. This can be used to replace a config value with a value interpolated from the random provider without experiencing diffs. Example with no prefix: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import random_id.server p-9hUg ` + "`" + `` + "`" + `` + "`" + ` Example with prefix (prefix is separated by a ` + "`" + `,` + "`" + `): ` + "`" + `` + "`" + `` + "`" + ` $ terraform import random_id.server my-prefix-,p-9hUg ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "b64_url",
					Description: `The generated id presented in base64, using the URL-friendly character set: case-sensitive letters, digits and the characters ` + "`" + `_` + "`" + ` and ` + "`" + `-` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "b64_std",
					Description: `The generated id presented in base64 without additional transformations.`,
				},
				resource.Attribute{
					Name:        "hex",
					Description: `The generated id presented in padded hexadecimal digits. This result will always be twice as long as the requested byte length.`,
				},
				resource.Attribute{
					Name:        "dec",
					Description: `The generated id presented in non-padded decimal digits. ## Import Random Ids can be imported using the ` + "`" + `b64_url` + "`" + ` with an optional ` + "`" + `prefix` + "`" + `. This can be used to replace a config value with a value interpolated from the random provider without experiencing diffs. Example with no prefix: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import random_id.server p-9hUg ` + "`" + `` + "`" + `` + "`" + ` Example with prefix (prefix is separated by a ` + "`" + `,` + "`" + `): ` + "`" + `` + "`" + `` + "`" + ` $ terraform import random_id.server my-prefix-,p-9hUg ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "random_integer",
			Category:         "Resources",
			ShortDescription: `Generates a random integer values.`,
			Description:      ``,
			Keywords: []string{
				"integer",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "min",
					Description: `(int) The minimum inclusive value of the range.`,
				},
				resource.Attribute{
					Name:        "max",
					Description: `(int) The maximum inclusive value of the range.`,
				},
				resource.Attribute{
					Name:        "keepers",
					Description: `(Optional) Arbitrary map of values that, when changed, will trigger a new id to be generated. See [the main provider documentation](../index.html) for more information.`,
				},
				resource.Attribute{
					Name:        "seed",
					Description: `(Optional) A custom seed to always produce the same value. ## Attribute Reference The following attributes are supported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(string) An internal id.`,
				},
				resource.Attribute{
					Name:        "result",
					Description: `(int) The random Integer result. ## Import Random integers can be imported using the ` + "`" + `result` + "`" + `, ` + "`" + `min` + "`" + `, and ` + "`" + `max` + "`" + `, with an optional ` + "`" + `seed` + "`" + `. This can be used to replace a config value with a value interpolated from the random provider without experiencing diffs. Example (values are separated by a ` + "`" + `,` + "`" + `): ` + "`" + `` + "`" + `` + "`" + ` $ terraform import random_integer.priority 15390,1,50000 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(string) An internal id.`,
				},
				resource.Attribute{
					Name:        "result",
					Description: `(int) The random Integer result. ## Import Random integers can be imported using the ` + "`" + `result` + "`" + `, ` + "`" + `min` + "`" + `, and ` + "`" + `max` + "`" + `, with an optional ` + "`" + `seed` + "`" + `. This can be used to replace a config value with a value interpolated from the random provider without experiencing diffs. Example (values are separated by a ` + "`" + `,` + "`" + `): ` + "`" + `` + "`" + `` + "`" + ` $ terraform import random_integer.priority 15390,1,50000 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "random_password",
			Category:         "Resources",
			ShortDescription: `Produces a random string of a length using alphanumeric characters and optionally special characters. The result will not be displayed to console.`,
			Description:      ``,
			Keywords: []string{
				"password",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "random_pet",
			Category:         "Resources",
			ShortDescription: `Generates a random pet.`,
			Description:      ``,
			Keywords: []string{
				"pet",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "keepers",
					Description: `(Optional) Arbitrary map of values that, when changed, will trigger a new id to be generated. See [the main provider documentation](../index.html) for more information.`,
				},
				resource.Attribute{
					Name:        "length",
					Description: `(Optional) The length (in words) of the pet name.`,
				},
				resource.Attribute{
					Name:        "prefix",
					Description: `(Optional) A string to prefix the name with.`,
				},
				resource.Attribute{
					Name:        "separator",
					Description: `(Optional) The character to separate words in the pet name. ## Attribute Reference The following attributes are supported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(string) The random pet name`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(string) The random pet name`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "random_shuffle",
			Category:         "Resources",
			ShortDescription: `Produces a random permutation of a given list.`,
			Description:      ``,
			Keywords: []string{
				"shuffle",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "input",
					Description: `(Required) The list of strings to shuffle.`,
				},
				resource.Attribute{
					Name:        "result_count",
					Description: `(Optional) The number of results to return. Defaults to the number of items in the ` + "`" + `input` + "`" + ` list. If fewer items are requested, some elements will be excluded from the result. If more items are requested, items will be repeated in the result but not more frequently than the number of items in the input list.`,
				},
				resource.Attribute{
					Name:        "keepers",
					Description: `(Optional) Arbitrary map of values that, when changed, will trigger a new id to be generated. See [the main provider documentation](../index.html) for more information.`,
				},
				resource.Attribute{
					Name:        "seed",
					Description: `(Optional) Arbitrary string with which to seed the random number generator, in order to produce less-volatile permutations of the list.`,
				},
				resource.Attribute{
					Name:        "result",
					Description: `Random permutation of the list of strings given in ` + "`" + `input` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "result",
					Description: `Random permutation of the list of strings given in ` + "`" + `input` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "random_string",
			Category:         "Resources",
			ShortDescription: `Produces a random string of a length using alphanumeric characters and optionally special characters.`,
			Description:      ``,
			Keywords: []string{
				"string",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "length",
					Description: `(Required) The length of the string desired`,
				},
				resource.Attribute{
					Name:        "upper",
					Description: `(Optional) (default true) Include uppercase alphabet characters in random string.`,
				},
				resource.Attribute{
					Name:        "min_upper",
					Description: `(Optional) (default 0) Minimum number of uppercase alphabet characters in random string.`,
				},
				resource.Attribute{
					Name:        "lower",
					Description: `(Optional) (default true) Include lowercase alphabet characters in random string.`,
				},
				resource.Attribute{
					Name:        "min_lower",
					Description: `(Optional) (default 0) Minimum number of lowercase alphabet characters in random string.`,
				},
				resource.Attribute{
					Name:        "number",
					Description: `(Optional) (default true) Include numeric characters in random string.`,
				},
				resource.Attribute{
					Name:        "min_numeric",
					Description: `(Optional) (default 0) Minimum number of numeric characters in random string.`,
				},
				resource.Attribute{
					Name:        "special",
					Description: `(Optional) (default true) Include special characters in random string. These are ` + "`" + `!@#$%&`,
				},
				resource.Attribute{
					Name:        "min_special",
					Description: `(Optional) (default 0) Minimum number of special characters in random string.`,
				},
				resource.Attribute{
					Name:        "override_special",
					Description: `(Optional) Supply your own list of special characters to use for string generation. This overrides the default character list in the special argument. The special argument must still be set to true for any overwritten characters to be used in generation.`,
				},
				resource.Attribute{
					Name:        "keepers",
					Description: `(Optional) Arbitrary map of values that, when changed, will trigger a new id to be generated. See [the main provider documentation](../index.html) for more information. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "result",
					Description: `Random string generated.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "result",
					Description: `Random string generated.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "random_uuid",
			Category:         "Resources",
			ShortDescription: `Generates a random identifier.`,
			Description:      ``,
			Keywords: []string{
				"uuid",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "keepers",
					Description: `(Optional) Arbitrary map of values that, when changed, will trigger a new uuid to be generated. See [the main provider documentation](../index.html) for more information. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "result",
					Description: `The generated uuid presented in string format. ## Import Random UUID's can be imported. This can be used to replace a config value with a value interpolated from the random provider without experiencing diffs. Example: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import random_uuid.main aabbccdd-eeff-0011-2233-445566778899 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "result",
					Description: `The generated uuid presented in string format. ## Import Random UUID's can be imported. This can be used to replace a config value with a value interpolated from the random provider without experiencing diffs. Example: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import random_uuid.main aabbccdd-eeff-0011-2233-445566778899 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
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
