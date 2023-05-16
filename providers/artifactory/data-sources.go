package artifactory

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "artifactory_artifactory_file",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "repository",
					Description: `(Required) Name of the repository where the file is stored.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Required) The path to the file within the repository.`,
				},
				resource.Attribute{
					Name:        "output_path",
					Description: `(Required) The local path the file should be downloaded to.`,
				},
				resource.Attribute{
					Name:        "force_overwrite",
					Description: `(Optional) If set to true, an existing file in the output_path will be overwritten. Default: ` + "`" + `false` + "`" + ``,
				},
				resource.Attribute{
					Name:        "path_is_aliased",
					Description: `(Optional) If set to ` + "`" + `true` + "`" + `, the provider will get the artifact directly from Artifactory without attempting to resolve it or verify it and will delegate this to artifactory if the file exists. More details in the [official documentation](https://www.jfrog.com/confluence/display/JFROG/Artifactory+REST+API#ArtifactoryRESTAPI-RetrieveLatestArtifact) ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "created",
					Description: `The time & date when the file was created.`,
				},
				resource.Attribute{
					Name:        "created_by",
					Description: `The user who created the file.`,
				},
				resource.Attribute{
					Name:        "last_modified",
					Description: `The time & date when the file was last modified.`,
				},
				resource.Attribute{
					Name:        "modified_by",
					Description: `The user who last modified the file.`,
				},
				resource.Attribute{
					Name:        "last_updated",
					Description: `The time & date when the file was last updated.`,
				},
				resource.Attribute{
					Name:        "mimetype",
					Description: `The mimetype of the file.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The size of the file.`,
				},
				resource.Attribute{
					Name:        "download_uri",
					Description: `The URI that can be used to download the file.`,
				},
				resource.Attribute{
					Name:        "md5",
					Description: `MD5 checksum of the file.`,
				},
				resource.Attribute{
					Name:        "sha1",
					Description: `SHA1 checksum of the file.`,
				},
				resource.Attribute{
					Name:        "sha256",
					Description: `SHA256 checksum of the file.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "created",
					Description: `The time & date when the file was created.`,
				},
				resource.Attribute{
					Name:        "created_by",
					Description: `The user who created the file.`,
				},
				resource.Attribute{
					Name:        "last_modified",
					Description: `The time & date when the file was last modified.`,
				},
				resource.Attribute{
					Name:        "modified_by",
					Description: `The user who last modified the file.`,
				},
				resource.Attribute{
					Name:        "last_updated",
					Description: `The time & date when the file was last updated.`,
				},
				resource.Attribute{
					Name:        "mimetype",
					Description: `The mimetype of the file.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The size of the file.`,
				},
				resource.Attribute{
					Name:        "download_uri",
					Description: `The URI that can be used to download the file.`,
				},
				resource.Attribute{
					Name:        "md5",
					Description: `MD5 checksum of the file.`,
				},
				resource.Attribute{
					Name:        "sha1",
					Description: `SHA1 checksum of the file.`,
				},
				resource.Attribute{
					Name:        "sha256",
					Description: `SHA256 checksum of the file.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_artifactory_fileinfo",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "repository",
					Description: `(Required) Name of the repository where the file is stored.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Required) The path to the file within the repository. ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "created",
					Description: `The time & date when the file was created.`,
				},
				resource.Attribute{
					Name:        "created_by",
					Description: `The user who created the file.`,
				},
				resource.Attribute{
					Name:        "last_modified",
					Description: `The time & date when the file was last modified.`,
				},
				resource.Attribute{
					Name:        "modified_by",
					Description: `The user who last modified the file.`,
				},
				resource.Attribute{
					Name:        "last_updated",
					Description: `The time & date when the file was last updated.`,
				},
				resource.Attribute{
					Name:        "mimetype",
					Description: `The mimetype of the file.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The size of the file.`,
				},
				resource.Attribute{
					Name:        "download_uri",
					Description: `The URI that can be used to download the file.`,
				},
				resource.Attribute{
					Name:        "md5",
					Description: `MD5 checksum of the file.`,
				},
				resource.Attribute{
					Name:        "sha1",
					Description: `SHA1 checksum of the file.`,
				},
				resource.Attribute{
					Name:        "sha256",
					Description: `SHA256 checksum of the file.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "created",
					Description: `The time & date when the file was created.`,
				},
				resource.Attribute{
					Name:        "created_by",
					Description: `The user who created the file.`,
				},
				resource.Attribute{
					Name:        "last_modified",
					Description: `The time & date when the file was last modified.`,
				},
				resource.Attribute{
					Name:        "modified_by",
					Description: `The user who last modified the file.`,
				},
				resource.Attribute{
					Name:        "last_updated",
					Description: `The time & date when the file was last updated.`,
				},
				resource.Attribute{
					Name:        "mimetype",
					Description: `The mimetype of the file.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The size of the file.`,
				},
				resource.Attribute{
					Name:        "download_uri",
					Description: `The URI that can be used to download the file.`,
				},
				resource.Attribute{
					Name:        "md5",
					Description: `MD5 checksum of the file.`,
				},
				resource.Attribute{
					Name:        "sha1",
					Description: `SHA1 checksum of the file.`,
				},
				resource.Attribute{
					Name:        "sha256",
					Description: `SHA256 checksum of the file.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_federated_alpine_repository",
			Category:         "Federated Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"federated",
				"repositories",
				"alpine",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) the identity key of the repo. ## Attribute Reference The following attributes are supported, along with the [list of attributes from the local Alpine repository](local_alpine_repository.md):`,
				},
				resource.Attribute{
					Name:        "member",
					Description: `The list of Federated members and must contain this repository URL (configured base URL ` + "`" + `/artifactory/` + "`" + ` + repo ` + "`" + `key` + "`" + `). Note that each of the federated members will need to have a base URL set. Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository) to set up Federated repositories correctly.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `Full URL to ending with the repository name.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Represents the active state of the federated member. It is supported to change the enabled status of my own member. The config will be updated on the other federated members automatically.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "member",
					Description: `The list of Federated members and must contain this repository URL (configured base URL ` + "`" + `/artifactory/` + "`" + ` + repo ` + "`" + `key` + "`" + `). Note that each of the federated members will need to have a base URL set. Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository) to set up Federated repositories correctly.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `Full URL to ending with the repository name.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Represents the active state of the federated member. It is supported to change the enabled status of my own member. The config will be updated on the other federated members automatically.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_federated_bower_repository",
			Category:         "Federated Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"federated",
				"repositories",
				"bower",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) the identity key of the repo. ## Attribute Reference The following attributes are supported, along with the [common list of arguments from the local repositories](local.md):`,
				},
				resource.Attribute{
					Name:        "member",
					Description: `The list of Federated members and must contain this repository URL (configured base URL ` + "`" + `/artifactory/` + "`" + ` + repo ` + "`" + `key` + "`" + `). Note that each of the federated members will need to have a base URL set. Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository) to set up Federated repositories correctly.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `Full URL to ending with the repository name.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Represents the active state of the federated member. It is supported to change the enabled status of my own member. The config will be updated on the other federated members automatically.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "member",
					Description: `The list of Federated members and must contain this repository URL (configured base URL ` + "`" + `/artifactory/` + "`" + ` + repo ` + "`" + `key` + "`" + `). Note that each of the federated members will need to have a base URL set. Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository) to set up Federated repositories correctly.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `Full URL to ending with the repository name.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Represents the active state of the federated member. It is supported to change the enabled status of my own member. The config will be updated on the other federated members automatically.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_federated_cargo_repository",
			Category:         "Federated Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"federated",
				"repositories",
				"cargo",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) the identity key of the repo. ## Attribute Reference The following attributes are supported, along with the [list of attributes from the local Cargo repository](local_cargo_repository.md):`,
				},
				resource.Attribute{
					Name:        "member",
					Description: `The list of Federated members and must contain this repository URL (configured base URL ` + "`" + `/artifactory/` + "`" + ` + repo ` + "`" + `key` + "`" + `). Note that each of the federated members will need to have a base URL set. Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository) to set up Federated repositories correctly.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `Full URL to ending with the repository name.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Represents the active state of the federated member. It is supported to change the enabled status of my own member. The config will be updated on the other federated members automatically.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "member",
					Description: `The list of Federated members and must contain this repository URL (configured base URL ` + "`" + `/artifactory/` + "`" + ` + repo ` + "`" + `key` + "`" + `). Note that each of the federated members will need to have a base URL set. Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository) to set up Federated repositories correctly.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `Full URL to ending with the repository name.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Represents the active state of the federated member. It is supported to change the enabled status of my own member. The config will be updated on the other federated members automatically.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_federated_chef_repository",
			Category:         "Federated Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"federated",
				"repositories",
				"chef",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) the identity key of the repo. ## Attribute Reference The following attributes are supported, along with the [common list of arguments from the local repositories](local.md):`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `the identity key of the repo.`,
				},
				resource.Attribute{
					Name:        "member",
					Description: `The list of Federated members and must contain this repository URL (configured base URL ` + "`" + `/artifactory/` + "`" + ` + repo ` + "`" + `key` + "`" + `). Note that each of the federated members will need to have a base URL set. Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository) to set up Federated repositories correctly.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `Full URL to ending with the repository name.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Represents the active state of the federated member. It is supported to change the enabled status of my own member. The config will be updated on the other federated members automatically.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `the identity key of the repo.`,
				},
				resource.Attribute{
					Name:        "member",
					Description: `The list of Federated members and must contain this repository URL (configured base URL ` + "`" + `/artifactory/` + "`" + ` + repo ` + "`" + `key` + "`" + `). Note that each of the federated members will need to have a base URL set. Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository) to set up Federated repositories correctly.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `Full URL to ending with the repository name.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Represents the active state of the federated member. It is supported to change the enabled status of my own member. The config will be updated on the other federated members automatically.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_federated_cocoapods_repository",
			Category:         "Federated Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"federated",
				"repositories",
				"cocoapods",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) the identity key of the repo. ## Attribute Reference The following attributes are supported, along with the [common list of arguments from the local repositories](local.md):`,
				},
				resource.Attribute{
					Name:        "member",
					Description: `The list of Federated members and must contain this repository URL (configured base URL ` + "`" + `/artifactory/` + "`" + ` + repo ` + "`" + `key` + "`" + `). Note that each of the federated members will need to have a base URL set. Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository) to set up Federated repositories correctly.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `Full URL to ending with the repository name.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Represents the active state of the federated member. It is supported to change the enabled status of my own member. The config will be updated on the other federated members automatically.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "member",
					Description: `The list of Federated members and must contain this repository URL (configured base URL ` + "`" + `/artifactory/` + "`" + ` + repo ` + "`" + `key` + "`" + `). Note that each of the federated members will need to have a base URL set. Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository) to set up Federated repositories correctly.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `Full URL to ending with the repository name.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Represents the active state of the federated member. It is supported to change the enabled status of my own member. The config will be updated on the other federated members automatically.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_federated_composer_repository",
			Category:         "Federated Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"federated",
				"repositories",
				"composer",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) the identity key of the repo. ## Attribute Reference The following attributes are supported, along with the [common list of arguments from the local repositories](local.md):`,
				},
				resource.Attribute{
					Name:        "member",
					Description: `The list of Federated members and must contain this repository URL (configured base URL ` + "`" + `/artifactory/` + "`" + ` + repo ` + "`" + `key` + "`" + `). Note that each of the federated members will need to have a base URL set. Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository) to set up Federated repositories correctly.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `Full URL to ending with the repository name.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Represents the active state of the federated member. It is supported to change the enabled status of my own member. The config will be updated on the other federated members automatically.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "member",
					Description: `The list of Federated members and must contain this repository URL (configured base URL ` + "`" + `/artifactory/` + "`" + ` + repo ` + "`" + `key` + "`" + `). Note that each of the federated members will need to have a base URL set. Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository) to set up Federated repositories correctly.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `Full URL to ending with the repository name.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Represents the active state of the federated member. It is supported to change the enabled status of my own member. The config will be updated on the other federated members automatically.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_federated_conan_repository",
			Category:         "Federated Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"federated",
				"repositories",
				"conan",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) the identity key of the repo. ## Attribute Reference The following attribute are supported, along with the [common list of arguments from the local repositories](local.md):`,
				},
				resource.Attribute{
					Name:        "member",
					Description: `The list of Federated members and must contain this repository URL (configured base URL ` + "`" + `/artifactory/` + "`" + ` + repo ` + "`" + `key` + "`" + `). Note that each of the federated members will need to have a base URL set. Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository) to set up Federated repositories correctly.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `Full URL to ending with the repository name.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Represents the active state of the federated member. It is supported to change the enabled status of my own member. The config will be updated on the other federated members automatically.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "member",
					Description: `The list of Federated members and must contain this repository URL (configured base URL ` + "`" + `/artifactory/` + "`" + ` + repo ` + "`" + `key` + "`" + `). Note that each of the federated members will need to have a base URL set. Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository) to set up Federated repositories correctly.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `Full URL to ending with the repository name.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Represents the active state of the federated member. It is supported to change the enabled status of my own member. The config will be updated on the other federated members automatically.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_federated_conda_repository",
			Category:         "Federated Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"federated",
				"repositories",
				"conda",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) the identity key of the repo. ## Attribute Reference The following attribute are supported, along with the [common list of arguments from the local repositories](local.md):`,
				},
				resource.Attribute{
					Name:        "member",
					Description: `The list of Federated members and must contain this repository URL (configured base URL ` + "`" + `/artifactory/` + "`" + ` + repo ` + "`" + `key` + "`" + `). Note that each of the federated members will need to have a base URL set. Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository) to set up Federated repositories correctly.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `Full URL to ending with the repository name.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Represents the active state of the federated member. It is supported to change the enabled status of my own member. The config will be updated on the other federated members automatically.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "member",
					Description: `The list of Federated members and must contain this repository URL (configured base URL ` + "`" + `/artifactory/` + "`" + ` + repo ` + "`" + `key` + "`" + `). Note that each of the federated members will need to have a base URL set. Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository) to set up Federated repositories correctly.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `Full URL to ending with the repository name.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Represents the active state of the federated member. It is supported to change the enabled status of my own member. The config will be updated on the other federated members automatically.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_federated_cran_repository",
			Category:         "Federated Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"federated",
				"repositories",
				"cran",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) the identity key of the repo. ## Attribute Reference The following attributes are supported, along with the [common list of arguments from the local repositories](local.md):`,
				},
				resource.Attribute{
					Name:        "member",
					Description: `The list of Federated members and must contain this repository URL (configured base URL ` + "`" + `/artifactory/` + "`" + ` + repo ` + "`" + `key` + "`" + `). Note that each of the federated members will need to have a base URL set. Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository) to set up Federated repositories correctly.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `Full URL to ending with the repository name.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Represents the active state of the federated member. It is supported to change the enabled status of my own member. The config will be updated on the other federated members automatically.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "member",
					Description: `The list of Federated members and must contain this repository URL (configured base URL ` + "`" + `/artifactory/` + "`" + ` + repo ` + "`" + `key` + "`" + `). Note that each of the federated members will need to have a base URL set. Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository) to set up Federated repositories correctly.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `Full URL to ending with the repository name.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Represents the active state of the federated member. It is supported to change the enabled status of my own member. The config will be updated on the other federated members automatically.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_federated_debian_repository",
			Category:         "Federated Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"federated",
				"repositories",
				"debian",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) the identity key of the repo. ## Attribute Reference The following attributes are supported, along with the [list of attributes from the local Debian repository](local_debian_repository.md):`,
				},
				resource.Attribute{
					Name:        "member",
					Description: `The list of Federated members and must contain this repository URL (configured base URL ` + "`" + `/artifactory/` + "`" + ` + repo ` + "`" + `key` + "`" + `). Note that each of the federated members will need to have a base URL set. Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository) to set up Federated repositories correctly.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `Full URL to ending with the repository name.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Represents the active state of the federated member. It is supported to change the enabled status of my own member. The config will be updated on the other federated members automatically.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "member",
					Description: `The list of Federated members and must contain this repository URL (configured base URL ` + "`" + `/artifactory/` + "`" + ` + repo ` + "`" + `key` + "`" + `). Note that each of the federated members will need to have a base URL set. Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository) to set up Federated repositories correctly.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `Full URL to ending with the repository name.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Represents the active state of the federated member. It is supported to change the enabled status of my own member. The config will be updated on the other federated members automatically.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_federated_docker_v1_repository",
			Category:         "Federated Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"federated",
				"repositories",
				"docker",
				"v1",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) the identity key of the repo. ## Attribute Reference The following attributes are supported, along with the [list of attributes from the local Docker V1 repository](local_docker_v1_repository.md):`,
				},
				resource.Attribute{
					Name:        "member",
					Description: `The list of Federated members and must contain this repository URL (configured base URL ` + "`" + `/artifactory/` + "`" + ` + repo ` + "`" + `key` + "`" + `). Note that each of the federated members will need to have a base URL set. Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository) to set up Federated repositories correctly.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `Full URL to ending with the repository name.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Represents the active state of the federated member. It is supported to change the enabled status of my own member. The config will be updated on the other federated members automatically.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "member",
					Description: `The list of Federated members and must contain this repository URL (configured base URL ` + "`" + `/artifactory/` + "`" + ` + repo ` + "`" + `key` + "`" + `). Note that each of the federated members will need to have a base URL set. Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository) to set up Federated repositories correctly.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `Full URL to ending with the repository name.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Represents the active state of the federated member. It is supported to change the enabled status of my own member. The config will be updated on the other federated members automatically.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_federated_docker_v2_repository",
			Category:         "Federated Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"federated",
				"repositories",
				"docker",
				"v2",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) the identity key of the repo. ## Attribute Reference The following attributes are supported, along with the [list of attributes from the local Docker V2 repository](local_docker_v2_repository.md):`,
				},
				resource.Attribute{
					Name:        "member",
					Description: `The list of Federated members and must contain this repository URL (configured base URL ` + "`" + `/artifactory/` + "`" + ` + repo ` + "`" + `key` + "`" + `). Note that each of the federated members will need to have a base URL set. Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository) to set up Federated repositories correctly.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `Full URL to ending with the repository name.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Represents the active state of the federated member. It is supported to change the enabled status of my own member. The config will be updated on the other federated members automatically.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "member",
					Description: `The list of Federated members and must contain this repository URL (configured base URL ` + "`" + `/artifactory/` + "`" + ` + repo ` + "`" + `key` + "`" + `). Note that each of the federated members will need to have a base URL set. Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository) to set up Federated repositories correctly.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `Full URL to ending with the repository name.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Represents the active state of the federated member. It is supported to change the enabled status of my own member. The config will be updated on the other federated members automatically.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_federated_gems_repository",
			Category:         "Federated Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"federated",
				"repositories",
				"gems",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) the identity key of the repo. ## Attribute Reference The following attribute are supported, along with the [common list of arguments from the local repositories](local.md):`,
				},
				resource.Attribute{
					Name:        "member",
					Description: `The list of Federated members and must contain this repository URL (configured base URL ` + "`" + `/artifactory/` + "`" + ` + repo ` + "`" + `key` + "`" + `). Note that each of the federated members will need to have a base URL set. Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository) to set up Federated repositories correctly.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `Full URL to ending with the repository name.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Represents the active state of the federated member. It is supported to change the enabled status of my own member. The config will be updated on the other federated members automatically.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "member",
					Description: `The list of Federated members and must contain this repository URL (configured base URL ` + "`" + `/artifactory/` + "`" + ` + repo ` + "`" + `key` + "`" + `). Note that each of the federated members will need to have a base URL set. Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository) to set up Federated repositories correctly.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `Full URL to ending with the repository name.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Represents the active state of the federated member. It is supported to change the enabled status of my own member. The config will be updated on the other federated members automatically.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_federated_generic_repository",
			Category:         "Federated Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"federated",
				"repositories",
				"generic",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) the identity key of the repo. ## Attribute Reference The following attribute are supported, along with the [common list of arguments from the local repositories](local.md):`,
				},
				resource.Attribute{
					Name:        "member",
					Description: `The list of Federated members and must contain this repository URL (configured base URL ` + "`" + `/artifactory/` + "`" + ` + repo ` + "`" + `key` + "`" + `). Note that each of the federated members will need to have a base URL set. Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository) to set up Federated repositories correctly.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `Full URL to ending with the repository name.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Represents the active state of the federated member. It is supported to change the enabled status of my own member. The config will be updated on the other federated members automatically.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "member",
					Description: `The list of Federated members and must contain this repository URL (configured base URL ` + "`" + `/artifactory/` + "`" + ` + repo ` + "`" + `key` + "`" + `). Note that each of the federated members will need to have a base URL set. Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository) to set up Federated repositories correctly.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `Full URL to ending with the repository name.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Represents the active state of the federated member. It is supported to change the enabled status of my own member. The config will be updated on the other federated members automatically.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_federated_gitlfs_repository",
			Category:         "Federated Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"federated",
				"repositories",
				"gitlfs",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) the identity key of the repo. ## Attribute Reference The following attribute are supported, along with the [common list of arguments from the local repositories](local.md):`,
				},
				resource.Attribute{
					Name:        "member",
					Description: `The list of Federated members and must contain this repository URL (configured base URL ` + "`" + `/artifactory/` + "`" + ` + repo ` + "`" + `key` + "`" + `). Note that each of the federated members will need to have a base URL set. Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository) to set up Federated repositories correctly.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `Full URL to ending with the repository name.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Represents the active state of the federated member. It is supported to change the enabled status of my own member. The config will be updated on the other federated members automatically.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "member",
					Description: `The list of Federated members and must contain this repository URL (configured base URL ` + "`" + `/artifactory/` + "`" + ` + repo ` + "`" + `key` + "`" + `). Note that each of the federated members will need to have a base URL set. Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository) to set up Federated repositories correctly.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `Full URL to ending with the repository name.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Represents the active state of the federated member. It is supported to change the enabled status of my own member. The config will be updated on the other federated members automatically.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_federated_go_repository",
			Category:         "Federated Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"federated",
				"repositories",
				"go",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) the identity key of the repo. ## Attribute Reference The following arguments are supported, along with the [common list of arguments from the local repositories](local.md):`,
				},
				resource.Attribute{
					Name:        "member",
					Description: `The list of Federated members and must contain this repository URL (configured base URL ` + "`" + `/artifactory/` + "`" + ` + repo ` + "`" + `key` + "`" + `). Note that each of the federated members will need to have a base URL set. Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository) to set up Federated repositories correctly.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `Full URL to ending with the repository name.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Represents the active state of the federated member. It is supported to change the enabled status of my own member. The config will be updated on the other federated members automatically.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "member",
					Description: `The list of Federated members and must contain this repository URL (configured base URL ` + "`" + `/artifactory/` + "`" + ` + repo ` + "`" + `key` + "`" + `). Note that each of the federated members will need to have a base URL set. Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository) to set up Federated repositories correctly.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `Full URL to ending with the repository name.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Represents the active state of the federated member. It is supported to change the enabled status of my own member. The config will be updated on the other federated members automatically.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_federated_gradle_repository",
			Category:         "Federated Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"federated",
				"repositories",
				"gradle",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) the identity key of the repo. ## Attribute Reference The following attributes are supported, along with the [list of attributes from the local Gradle repository](local_gradle_repository.md):`,
				},
				resource.Attribute{
					Name:        "member",
					Description: `The list of Federated members and must contain this repository URL (configured base URL ` + "`" + `/artifactory/` + "`" + ` + repo ` + "`" + `key` + "`" + `). Note that each of the federated members will need to have a base URL set. Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository) to set up Federated repositories correctly.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `Full URL to ending with the repository name.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Represents the active state of the federated member. It is supported to change the enabled status of my own member. The config will be updated on the other federated members automatically.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "member",
					Description: `The list of Federated members and must contain this repository URL (configured base URL ` + "`" + `/artifactory/` + "`" + ` + repo ` + "`" + `key` + "`" + `). Note that each of the federated members will need to have a base URL set. Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository) to set up Federated repositories correctly.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `Full URL to ending with the repository name.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Represents the active state of the federated member. It is supported to change the enabled status of my own member. The config will be updated on the other federated members automatically.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_federated_helm_repository",
			Category:         "Federated Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"federated",
				"repositories",
				"helm",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) the identity key of the repo. ## Attribute Reference The following attributes are supported, along with the [common list of arguments from the local repositories](local.md):`,
				},
				resource.Attribute{
					Name:        "member",
					Description: `The list of Federated members and must contain this repository URL (configured base URL ` + "`" + `/artifactory/` + "`" + ` + repo ` + "`" + `key` + "`" + `). Note that each of the federated members will need to have a base URL set. Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository) to set up Federated repositories correctly.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `Full URL to ending with the repository name.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Represents the active state of the federated member. It is supported to change the enabled status of my own member. The config will be updated on the other federated members automatically.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "member",
					Description: `The list of Federated members and must contain this repository URL (configured base URL ` + "`" + `/artifactory/` + "`" + ` + repo ` + "`" + `key` + "`" + `). Note that each of the federated members will need to have a base URL set. Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository) to set up Federated repositories correctly.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `Full URL to ending with the repository name.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Represents the active state of the federated member. It is supported to change the enabled status of my own member. The config will be updated on the other federated members automatically.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_federated_ivy_repository",
			Category:         "Federated Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"federated",
				"repositories",
				"ivy",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) the identity key of the repo. ## Attribute Reference The following attributes are supported, along with the [list of attributes from the local Ivy repository](local_ivy_repository.md):`,
				},
				resource.Attribute{
					Name:        "member",
					Description: `The list of Federated members and must contain this repository URL (configured base URL ` + "`" + `/artifactory/` + "`" + ` + repo ` + "`" + `key` + "`" + `). Note that each of the federated members will need to have a base URL set. Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository) to set up Federated repositories correctly.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `Full URL to ending with the repository name.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Represents the active state of the federated member. It is supported to change the enabled status of my own member. The config will be updated on the other federated members automatically.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "member",
					Description: `The list of Federated members and must contain this repository URL (configured base URL ` + "`" + `/artifactory/` + "`" + ` + repo ` + "`" + `key` + "`" + `). Note that each of the federated members will need to have a base URL set. Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository) to set up Federated repositories correctly.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `Full URL to ending with the repository name.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Represents the active state of the federated member. It is supported to change the enabled status of my own member. The config will be updated on the other federated members automatically.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_federated_maven_repository",
			Category:         "Federated Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"federated",
				"repositories",
				"maven",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) the identity key of the repo. ## Attribute Reference The following attributes are supported, along with the [list of attributes from the local Maven repository](local_maven_repository.md):`,
				},
				resource.Attribute{
					Name:        "member",
					Description: `The list of Federated members and must contain this repository URL (configured base URL ` + "`" + `/artifactory/` + "`" + ` + repo ` + "`" + `key` + "`" + `). Note that each of the federated members will need to have a base URL set. Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository) to set up Federated repositories correctly.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `Full URL to ending with the repository name.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Represents the active state of the federated member. It is supported to change the enabled status of my own member. The config will be updated on the other federated members automatically.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "member",
					Description: `The list of Federated members and must contain this repository URL (configured base URL ` + "`" + `/artifactory/` + "`" + ` + repo ` + "`" + `key` + "`" + `). Note that each of the federated members will need to have a base URL set. Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository) to set up Federated repositories correctly.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `Full URL to ending with the repository name.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Represents the active state of the federated member. It is supported to change the enabled status of my own member. The config will be updated on the other federated members automatically.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_federated_npm_repository",
			Category:         "Federated Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"federated",
				"repositories",
				"npm",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) the identity key of the repo. ## Attribute Reference The following attribute are supported, along with the [common list of arguments from the local repositories](local.md):`,
				},
				resource.Attribute{
					Name:        "member",
					Description: `The list of Federated members and must contain this repository URL (configured base URL ` + "`" + `/artifactory/` + "`" + ` + repo ` + "`" + `key` + "`" + `). Note that each of the federated members will need to have a base URL set. Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository) to set up Federated repositories correctly.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `Full URL to ending with the repository name.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Represents the active state of the federated member. It is supported to change the enabled status of my own member. The config will be updated on the other federated members automatically.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "member",
					Description: `The list of Federated members and must contain this repository URL (configured base URL ` + "`" + `/artifactory/` + "`" + ` + repo ` + "`" + `key` + "`" + `). Note that each of the federated members will need to have a base URL set. Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository) to set up Federated repositories correctly.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `Full URL to ending with the repository name.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Represents the active state of the federated member. It is supported to change the enabled status of my own member. The config will be updated on the other federated members automatically.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_federated_nuget_repository",
			Category:         "Federated Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"federated",
				"repositories",
				"nuget",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) the identity key of the repo. ## Attribute Reference The following attributes are supported, along with the [list of attributes from the local Nuget repository](local_nuget_repository.md):`,
				},
				resource.Attribute{
					Name:        "member",
					Description: `The list of Federated members and must contain this repository URL (configured base URL ` + "`" + `/artifactory/` + "`" + ` + repo ` + "`" + `key` + "`" + `). Note that each of the federated members will need to have a base URL set. Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository) to set up Federated repositories correctly.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `Full URL to ending with the repository name.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Represents the active state of the federated member. It is supported to change the enabled status of my own member. The config will be updated on the other federated members automatically.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "member",
					Description: `The list of Federated members and must contain this repository URL (configured base URL ` + "`" + `/artifactory/` + "`" + ` + repo ` + "`" + `key` + "`" + `). Note that each of the federated members will need to have a base URL set. Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository) to set up Federated repositories correctly.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `Full URL to ending with the repository name.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Represents the active state of the federated member. It is supported to change the enabled status of my own member. The config will be updated on the other federated members automatically.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_federated_opkg_repository",
			Category:         "Federated Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"federated",
				"repositories",
				"opkg",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) the identity key of the repo. ## Attribute Reference The following attribute are supported, along with the [common list of arguments from the local repositories](local.md):`,
				},
				resource.Attribute{
					Name:        "member",
					Description: `The list of Federated members and must contain this repository URL (configured base URL ` + "`" + `/artifactory/` + "`" + ` + repo ` + "`" + `key` + "`" + `). Note that each of the federated members will need to have a base URL set. Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository) to set up Federated repositories correctly.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `Full URL to ending with the repository name.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Represents the active state of the federated member. It is supported to change the enabled status of my own member. The config will be updated on the other federated members automatically.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "member",
					Description: `The list of Federated members and must contain this repository URL (configured base URL ` + "`" + `/artifactory/` + "`" + ` + repo ` + "`" + `key` + "`" + `). Note that each of the federated members will need to have a base URL set. Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository) to set up Federated repositories correctly.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `Full URL to ending with the repository name.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Represents the active state of the federated member. It is supported to change the enabled status of my own member. The config will be updated on the other federated members automatically.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_federated_puppet_repository",
			Category:         "Federated Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"federated",
				"repositories",
				"puppet",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) the identity key of the repo. ## Attribute Reference The following attribute are supported, along with the [common list of arguments from the local repositories](local.md):`,
				},
				resource.Attribute{
					Name:        "member",
					Description: `The list of Federated members and must contain this repository URL (configured base URL ` + "`" + `/artifactory/` + "`" + ` + repo ` + "`" + `key` + "`" + `). Note that each of the federated members will need to have a base URL set. Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository) to set up Federated repositories correctly.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `Full URL to ending with the repository name.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Represents the active state of the federated member. It is supported to change the enabled status of my own member. The config will be updated on the other federated members automatically.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "member",
					Description: `The list of Federated members and must contain this repository URL (configured base URL ` + "`" + `/artifactory/` + "`" + ` + repo ` + "`" + `key` + "`" + `). Note that each of the federated members will need to have a base URL set. Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository) to set up Federated repositories correctly.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `Full URL to ending with the repository name.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Represents the active state of the federated member. It is supported to change the enabled status of my own member. The config will be updated on the other federated members automatically.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_federated_pypi_repository",
			Category:         "Federated Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"federated",
				"repositories",
				"pypi",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) the identity key of the repo. ## Attribute Reference The following attribute are supported, along with the [common list of arguments from the local repositories](local.md):`,
				},
				resource.Attribute{
					Name:        "member",
					Description: `The list of Federated members and must contain this repository URL (configured base URL ` + "`" + `/artifactory/` + "`" + ` + repo ` + "`" + `key` + "`" + `). Note that each of the federated members will need to have a base URL set. Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository) to set up Federated repositories correctly.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `Full URL to ending with the repository name.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Represents the active state of the federated member. It is supported to change the enabled status of my own member. The config will be updated on the other federated members automatically.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "member",
					Description: `The list of Federated members and must contain this repository URL (configured base URL ` + "`" + `/artifactory/` + "`" + ` + repo ` + "`" + `key` + "`" + `). Note that each of the federated members will need to have a base URL set. Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository) to set up Federated repositories correctly.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `Full URL to ending with the repository name.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Represents the active state of the federated member. It is supported to change the enabled status of my own member. The config will be updated on the other federated members automatically.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_federated_rpm_repository",
			Category:         "Federated Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"federated",
				"repositories",
				"rpm",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) the identity key of the repo. ## Attribute Reference The following attributes are supported, along with the [list of attributes from the local RPM repository](local_rpm_repository.md):`,
				},
				resource.Attribute{
					Name:        "member",
					Description: `The list of Federated members and must contain this repository URL (configured base URL ` + "`" + `/artifactory/` + "`" + ` + repo ` + "`" + `key` + "`" + `). Note that each of the federated members will need to have a base URL set. Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository) to set up Federated repositories correctly.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `Full URL to ending with the repository name.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Represents the active state of the federated member. It is supported to change the enabled status of my own member. The config will be updated on the other federated members automatically.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "member",
					Description: `The list of Federated members and must contain this repository URL (configured base URL ` + "`" + `/artifactory/` + "`" + ` + repo ` + "`" + `key` + "`" + `). Note that each of the federated members will need to have a base URL set. Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository) to set up Federated repositories correctly.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `Full URL to ending with the repository name.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Represents the active state of the federated member. It is supported to change the enabled status of my own member. The config will be updated on the other federated members automatically.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_federated_sbt_repository",
			Category:         "Federated Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"federated",
				"repositories",
				"sbt",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) the identity key of the repo. ## Attribute Reference The following attributes are supported, along with the [list of attributes from the local SBT repository](local_sbt_repository.md):`,
				},
				resource.Attribute{
					Name:        "member",
					Description: `The list of Federated members and must contain this repository URL (configured base URL ` + "`" + `/artifactory/` + "`" + ` + repo ` + "`" + `key` + "`" + `). Note that each of the federated members will need to have a base URL set. Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository) to set up Federated repositories correctly.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `Full URL to ending with the repository name.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Represents the active state of the federated member. It is supported to change the enabled status of my own member. The config will be updated on the other federated members automatically.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "member",
					Description: `The list of Federated members and must contain this repository URL (configured base URL ` + "`" + `/artifactory/` + "`" + ` + repo ` + "`" + `key` + "`" + `). Note that each of the federated members will need to have a base URL set. Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository) to set up Federated repositories correctly.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `Full URL to ending with the repository name.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Represents the active state of the federated member. It is supported to change the enabled status of my own member. The config will be updated on the other federated members automatically.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_federated_swift_repository",
			Category:         "Federated Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"federated",
				"repositories",
				"swift",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) the identity key of the repo. ## Attribute Reference The following arguments are supported, along with the [common list of arguments from the local repositories](local.md):`,
				},
				resource.Attribute{
					Name:        "member",
					Description: `The list of Federated members and must contain this repository URL (configured base URL ` + "`" + `/artifactory/` + "`" + ` + repo ` + "`" + `key` + "`" + `). Note that each of the federated members will need to have a base URL set. Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository) to set up Federated repositories correctly.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `Full URL to ending with the repository name.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Represents the active state of the federated member. It is supported to change the enabled status of my own member. The config will be updated on the other federated members automatically.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "member",
					Description: `The list of Federated members and must contain this repository URL (configured base URL ` + "`" + `/artifactory/` + "`" + ` + repo ` + "`" + `key` + "`" + `). Note that each of the federated members will need to have a base URL set. Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository) to set up Federated repositories correctly.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `Full URL to ending with the repository name.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Represents the active state of the federated member. It is supported to change the enabled status of my own member. The config will be updated on the other federated members automatically.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_federated_terraform_module_repository",
			Category:         "Federated Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"federated",
				"repositories",
				"terraform",
				"module",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) the identity key of the repo. ## Attribute Reference The following attributes are supported, along with the [list of attributes from the local Terraform Module repository](local_terraform_module_repository.md):`,
				},
				resource.Attribute{
					Name:        "member",
					Description: `The list of Federated members and must contain this repository URL (configured base URL ` + "`" + `/artifactory/` + "`" + ` + repo ` + "`" + `key` + "`" + `). Note that each of the federated members will need to have a base URL set. Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository) to set up Federated repositories correctly.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `Full URL to ending with the repository name.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Represents the active state of the federated member. It is supported to change the enabled status of my own member. The config will be updated on the other federated members automatically.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "member",
					Description: `The list of Federated members and must contain this repository URL (configured base URL ` + "`" + `/artifactory/` + "`" + ` + repo ` + "`" + `key` + "`" + `). Note that each of the federated members will need to have a base URL set. Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository) to set up Federated repositories correctly.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `Full URL to ending with the repository name.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Represents the active state of the federated member. It is supported to change the enabled status of my own member. The config will be updated on the other federated members automatically.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_federated_terraform_provider_repository",
			Category:         "Federated Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"federated",
				"repositories",
				"terraform",
				"provider",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) the identity key of the repo. ## Attribute Reference The following attributes are supported, along with the [list of attributes from the local Terraform Provider repository](local_terraform_provider_repository.md):`,
				},
				resource.Attribute{
					Name:        "member",
					Description: `The list of Federated members and must contain this repository URL (configured base URL ` + "`" + `/artifactory/` + "`" + ` + repo ` + "`" + `key` + "`" + `). Note that each of the federated members will need to have a base URL set. Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository) to set up Federated repositories correctly.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `Full URL to ending with the repository name.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Represents the active state of the federated member. It is supported to change the enabled status of my own member. The config will be updated on the other federated members automatically.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "member",
					Description: `The list of Federated members and must contain this repository URL (configured base URL ` + "`" + `/artifactory/` + "`" + ` + repo ` + "`" + `key` + "`" + `). Note that each of the federated members will need to have a base URL set. Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository) to set up Federated repositories correctly.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `Full URL to ending with the repository name.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Represents the active state of the federated member. It is supported to change the enabled status of my own member. The config will be updated on the other federated members automatically.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_federated_vagrant_repository",
			Category:         "Federated Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"federated",
				"repositories",
				"vagrant",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) the identity key of the repo. ## Attribute Reference The following attributes are supported, along with the [common list of arguments from the local repositories](local.md):`,
				},
				resource.Attribute{
					Name:        "member",
					Description: `The list of Federated members and must contain this repository URL (configured base URL ` + "`" + `/artifactory/` + "`" + ` + repo ` + "`" + `key` + "`" + `). Note that each of the federated members will need to have a base URL set. Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository) to set up Federated repositories correctly.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `Full URL to ending with the repository name.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Represents the active state of the federated member. It is supported to change the enabled status of my own member. The config will be updated on the other federated members automatically.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "member",
					Description: `The list of Federated members and must contain this repository URL (configured base URL ` + "`" + `/artifactory/` + "`" + ` + repo ` + "`" + `key` + "`" + `). Note that each of the federated members will need to have a base URL set. Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository) to set up Federated repositories correctly.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `Full URL to ending with the repository name.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Represents the active state of the federated member. It is supported to change the enabled status of my own member. The config will be updated on the other federated members automatically.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_group",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the group.`,
				},
				resource.Attribute{
					Name:        "include_users",
					Description: `(Optional) Determines if the group's associated user list will return as an attribute. Default is ` + "`" + `false` + "`" + `. ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_local",
			Category:         "Local Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"local",
				"repositories",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) Name of the repository. ## Attribute Reference In addition to all arguments above, the following attributes are exported for all local repositories:`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `A mandatory identifier for the repository that must be unique. It cannot begin with a number or contain spaces or special characters.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the repository`,
				},
				resource.Attribute{
					Name:        "project_key",
					Description: `Project key for assigning this repository to. Will be 2 - 20 lowercase alphanumeric and hyphen characters. When assigning repository to a project, repository key must be prefixed with project key, separated by a dash. We don't recommend using this attribute to assign the repository to the project. Use the ` + "`" + `repos` + "`" + ` attribute in Project provider to manage the list of repositories.`,
				},
				resource.Attribute{
					Name:        "project_environments",
					Description: `Project environment for assigning this repository to. Allow values: ` + "`" + `DEV` + "`" + ` or ` + "`" + `PROD` + "`" + `. The attribute should only be used if the repository is already assigned to the existing project. If not, the attribute will be ignored by Artifactory, but will remain in the Terraform state, which will create state drift during the update.`,
				},
				resource.Attribute{
					Name:        "includes_pattern",
					Description: `List of artifact patterns to include when evaluating artifact requests in the form of x/y/`,
				},
				resource.Attribute{
					Name:        "excludes_pattern",
					Description: `List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/`,
				},
				resource.Attribute{
					Name:        "repo_layout_ref",
					Description: `Sets the layout that the repository should use for storing and identifying modules. A recommended layout that corresponds to the package type defined is suggested, and index packages uploaded and calculate metadata accordingly.`,
				},
				resource.Attribute{
					Name:        "blacked_out",
					Description: `(Default: ` + "`" + `false` + "`" + `) When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.`,
				},
				resource.Attribute{
					Name:        "xray_index",
					Description: `(Default: ` + "`" + `false` + "`" + `) Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via Xray settings.`,
				},
				resource.Attribute{
					Name:        "priority_resolution",
					Description: `(Default: ` + "`" + `false` + "`" + `) Setting repositories with priority will cause metadata to be merged only from repositories set with this field`,
				},
				resource.Attribute{
					Name:        "property_sets",
					Description: `List of property set names`,
				},
				resource.Attribute{
					Name:        "archive_browsing_enabled",
					Description: `When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and therefore requires strict content moderation to prevent malicious users from uploading content that may compromise security (e.g., cross-site scripting attacks).`,
				},
				resource.Attribute{
					Name:        "download_direct",
					Description: `When set, download requests to this repository will redirect the client to download the artifact directly from the cloud storage provider. Available in Enterprise+ and Edge licenses only.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `A mandatory identifier for the repository that must be unique. It cannot begin with a number or contain spaces or special characters.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the repository`,
				},
				resource.Attribute{
					Name:        "project_key",
					Description: `Project key for assigning this repository to. Will be 2 - 20 lowercase alphanumeric and hyphen characters. When assigning repository to a project, repository key must be prefixed with project key, separated by a dash. We don't recommend using this attribute to assign the repository to the project. Use the ` + "`" + `repos` + "`" + ` attribute in Project provider to manage the list of repositories.`,
				},
				resource.Attribute{
					Name:        "project_environments",
					Description: `Project environment for assigning this repository to. Allow values: ` + "`" + `DEV` + "`" + ` or ` + "`" + `PROD` + "`" + `. The attribute should only be used if the repository is already assigned to the existing project. If not, the attribute will be ignored by Artifactory, but will remain in the Terraform state, which will create state drift during the update.`,
				},
				resource.Attribute{
					Name:        "includes_pattern",
					Description: `List of artifact patterns to include when evaluating artifact requests in the form of x/y/`,
				},
				resource.Attribute{
					Name:        "excludes_pattern",
					Description: `List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/`,
				},
				resource.Attribute{
					Name:        "repo_layout_ref",
					Description: `Sets the layout that the repository should use for storing and identifying modules. A recommended layout that corresponds to the package type defined is suggested, and index packages uploaded and calculate metadata accordingly.`,
				},
				resource.Attribute{
					Name:        "blacked_out",
					Description: `(Default: ` + "`" + `false` + "`" + `) When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.`,
				},
				resource.Attribute{
					Name:        "xray_index",
					Description: `(Default: ` + "`" + `false` + "`" + `) Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via Xray settings.`,
				},
				resource.Attribute{
					Name:        "priority_resolution",
					Description: `(Default: ` + "`" + `false` + "`" + `) Setting repositories with priority will cause metadata to be merged only from repositories set with this field`,
				},
				resource.Attribute{
					Name:        "property_sets",
					Description: `List of property set names`,
				},
				resource.Attribute{
					Name:        "archive_browsing_enabled",
					Description: `When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and therefore requires strict content moderation to prevent malicious users from uploading content that may compromise security (e.g., cross-site scripting attacks).`,
				},
				resource.Attribute{
					Name:        "download_direct",
					Description: `When set, download requests to this repository will redirect the client to download the artifact directly from the cloud storage provider. Available in Enterprise+ and Edge licenses only.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_local_alpine_repository",
			Category:         "Local Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"local",
				"repositories",
				"alpine",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) the identity key of the repo. ## Attribute Reference The following attributes are supported, along with the [common list of attributes for the local repositories](local.md):`,
				},
				resource.Attribute{
					Name:        "primary_keypair_ref",
					Description: `The RSA key to be used to sign alpine indices.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "primary_keypair_ref",
					Description: `The RSA key to be used to sign alpine indices.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_local_bower_repository",
			Category:         "Local Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"local",
				"repositories",
				"bower",
				"repository",
			},
			Arguments: []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) the identity key of the repo.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `(Optional)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_local_cargo_repository",
			Category:         "Local Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"local",
				"repositories",
				"cargo",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) the identity key of the repo. ## Attribute Reference The following attributes are supported, along with the [common list of attributes for the local repositories](local.md):`,
				},
				resource.Attribute{
					Name:        "anonymous_access",
					Description: `Cargo client does not send credentials when performing download and search for crates. Enable this to allow anonymous access to these resources (only), note that this will override the security anonymous access option. Default value is ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "enable_sparse_index",
					Description: `Enable internal index support based on Cargo sparse index specifications, instead of the default git index. Default value is ` + "`" + `false` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "anonymous_access",
					Description: `Cargo client does not send credentials when performing download and search for crates. Enable this to allow anonymous access to these resources (only), note that this will override the security anonymous access option. Default value is ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "enable_sparse_index",
					Description: `Enable internal index support based on Cargo sparse index specifications, instead of the default git index. Default value is ` + "`" + `false` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_local_chef_repository",
			Category:         "Local Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"local",
				"repositories",
				"chef",
				"repository",
			},
			Arguments: []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) the identity key of the repo.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `(Optional)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_local_cocoapods_repository",
			Category:         "Local Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"local",
				"repositories",
				"cocoapods",
				"repository",
			},
			Arguments: []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `the identity key of the repo.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_local_composer_repository",
			Category:         "Local Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"local",
				"repositories",
				"composer",
				"repository",
			},
			Arguments: []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `the identity key of the repo.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_local_conan_repository",
			Category:         "Local Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"local",
				"repositories",
				"conan",
				"repository",
			},
			Arguments: []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `the identity key of the repo.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_local_conda_repository",
			Category:         "Local Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"local",
				"repositories",
				"conda",
				"repository",
			},
			Arguments: []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `the identity key of the repo.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_local_cran_repository",
			Category:         "Local Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"local",
				"repositories",
				"cran",
				"repository",
			},
			Arguments: []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `the identity key of the repo.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_local_debian_repository",
			Category:         "Local Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"local",
				"repositories",
				"debian",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) the identity key of the repo. ## Attribute Reference The following attributes are supported, along with the [common list of attributes for the local repositories](local.md):`,
				},
				resource.Attribute{
					Name:        "primary_keypair_ref",
					Description: `The primary RSA key to be used to sign packages.`,
				},
				resource.Attribute{
					Name:        "secondary_keypair_ref",
					Description: `The secondary RSA key to be used to sign packages.`,
				},
				resource.Attribute{
					Name:        "index_compression_formats",
					Description: `The options are Bzip2 (.bz2 extension) (default), LZMA (.lzma extension) and XZ (.xz extension).`,
				},
				resource.Attribute{
					Name:        "trivial_layout",
					Description: `When set, the repository will use the deprecated trivial layout.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "primary_keypair_ref",
					Description: `The primary RSA key to be used to sign packages.`,
				},
				resource.Attribute{
					Name:        "secondary_keypair_ref",
					Description: `The secondary RSA key to be used to sign packages.`,
				},
				resource.Attribute{
					Name:        "index_compression_formats",
					Description: `The options are Bzip2 (.bz2 extension) (default), LZMA (.lzma extension) and XZ (.xz extension).`,
				},
				resource.Attribute{
					Name:        "trivial_layout",
					Description: `When set, the repository will use the deprecated trivial layout.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_local_docker_v1_repository",
			Category:         "Local Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"local",
				"repositories",
				"docker",
				"v1",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) the identity key of the repo. ## Attribute Reference The following attributes are supported, along with the [common list of attributes for the local repositories](local.md):`,
				},
				resource.Attribute{
					Name:        "block_pushing_schema1",
					Description: `When set, Artifactory will block the pushing of Docker images with manifest v2 schema 1 to this repository.`,
				},
				resource.Attribute{
					Name:        "tag_retention",
					Description: `If greater than 1, overwritten tags will be saved by their digest, up to the set up number. This only applies to manifest V2.`,
				},
				resource.Attribute{
					Name:        "max_unique_tags",
					Description: `The maximum number of unique tags of a single Docker image to store in this repository. Once the number tags for an image exceeds this setting, older tags are removed. A value of 0 (default) indicates there is no limit. This only applies to manifest v2.`,
				},
				resource.Attribute{
					Name:        "api_version",
					Description: `The Docker API version in use.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "block_pushing_schema1",
					Description: `When set, Artifactory will block the pushing of Docker images with manifest v2 schema 1 to this repository.`,
				},
				resource.Attribute{
					Name:        "tag_retention",
					Description: `If greater than 1, overwritten tags will be saved by their digest, up to the set up number. This only applies to manifest V2.`,
				},
				resource.Attribute{
					Name:        "max_unique_tags",
					Description: `The maximum number of unique tags of a single Docker image to store in this repository. Once the number tags for an image exceeds this setting, older tags are removed. A value of 0 (default) indicates there is no limit. This only applies to manifest v2.`,
				},
				resource.Attribute{
					Name:        "api_version",
					Description: `The Docker API version in use.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_local_docker_v2_repository",
			Category:         "Local Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"local",
				"repositories",
				"docker",
				"v2",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) the identity key of the repo. ## Attribute Reference The following attributes are supported, along with the [common list of attributes for the local repositories](local.md):`,
				},
				resource.Attribute{
					Name:        "block_pushing_schema1",
					Description: `When set, Artifactory will block the pushing of Docker images with manifest v2 schema 1 to this repository.`,
				},
				resource.Attribute{
					Name:        "tag_retention",
					Description: `If greater than 1, overwritten tags will be saved by their digest, up to the set up number. This only applies to manifest V2.`,
				},
				resource.Attribute{
					Name:        "max_unique_tags",
					Description: `The maximum number of unique tags of a single Docker image to store in this repository. Once the number tags for an image exceeds this setting, older tags are removed. A value of 0 (default) indicates there is no limit. This only applies to manifest v2.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "block_pushing_schema1",
					Description: `When set, Artifactory will block the pushing of Docker images with manifest v2 schema 1 to this repository.`,
				},
				resource.Attribute{
					Name:        "tag_retention",
					Description: `If greater than 1, overwritten tags will be saved by their digest, up to the set up number. This only applies to manifest V2.`,
				},
				resource.Attribute{
					Name:        "max_unique_tags",
					Description: `The maximum number of unique tags of a single Docker image to store in this repository. Once the number tags for an image exceeds this setting, older tags are removed. A value of 0 (default) indicates there is no limit. This only applies to manifest v2.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_local_gems_repository",
			Category:         "Local Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"local",
				"repositories",
				"gems",
				"repository",
			},
			Arguments: []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `the identity key of the repo.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_local_generic_repository",
			Category:         "Local Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"local",
				"repositories",
				"generic",
				"repository",
			},
			Arguments: []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `the identity key of the repo.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_local_gitlfs_repository",
			Category:         "Local Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"local",
				"repositories",
				"gitlfs",
				"repository",
			},
			Arguments: []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `the identity key of the repo.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_local_go_repository",
			Category:         "Local Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"local",
				"repositories",
				"go",
				"repository",
			},
			Arguments: []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `the identity key of the repo.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_local_gradle_repository",
			Category:         "Local Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"local",
				"repositories",
				"gradle",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) the identity key of the repo. ## Attribute Reference The following attributes are supported, along with the [common list of attributes for the local repositories](local.md):`,
				},
				resource.Attribute{
					Name:        "checksum_policy_type",
					Description: `Checksum policy determines how Artifactory behaves when a client checksum for a deployed resource is missing or conflicts with the locally calculated checksum (bad checksum). The options are ` + "`" + `client-checksums` + "`" + ` and ` + "`" + `generated-checksums` + "`" + `. For more details, please refer to [Checksum Policy](https://www.jfrog.com/confluence/display/JFROG/Local+Repositories#LocalRepositories-ChecksumPolicy) .`,
				},
				resource.Attribute{
					Name:        "snapshot_version_behavior",
					Description: `Specifies the naming convention for Maven SNAPSHOT versions. The options are -`,
				},
				resource.Attribute{
					Name:        "max_unique_snapshots",
					Description: `The maximum number of unique snapshots of a single artifact to store. Once the number of snapshots exceeds this setting, older versions are removed. A value of 0 (default) indicates there is no limit, and unique snapshots are not cleaned up.`,
				},
				resource.Attribute{
					Name:        "handle_releases",
					Description: `If set, Artifactory allows you to deploy release artifacts into this repository. Default is ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "handle_snapshots",
					Description: `If set, Artifactory allows you to deploy snapshot artifacts into this repository. Default is ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "suppress_pom_consistency_checks",
					Description: `By default, Artifactory keeps your repositories healthy by refusing POMs with incorrect coordinates (path). If the groupId:artifactId:version information inside the POM does not match the deployed path, Artifactory rejects the deployment with a "409 Conflict" error. You can disable this behavior by setting the Suppress POM Consistency Checks checkbox. True by default for Gradle repository.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "checksum_policy_type",
					Description: `Checksum policy determines how Artifactory behaves when a client checksum for a deployed resource is missing or conflicts with the locally calculated checksum (bad checksum). The options are ` + "`" + `client-checksums` + "`" + ` and ` + "`" + `generated-checksums` + "`" + `. For more details, please refer to [Checksum Policy](https://www.jfrog.com/confluence/display/JFROG/Local+Repositories#LocalRepositories-ChecksumPolicy) .`,
				},
				resource.Attribute{
					Name:        "snapshot_version_behavior",
					Description: `Specifies the naming convention for Maven SNAPSHOT versions. The options are -`,
				},
				resource.Attribute{
					Name:        "max_unique_snapshots",
					Description: `The maximum number of unique snapshots of a single artifact to store. Once the number of snapshots exceeds this setting, older versions are removed. A value of 0 (default) indicates there is no limit, and unique snapshots are not cleaned up.`,
				},
				resource.Attribute{
					Name:        "handle_releases",
					Description: `If set, Artifactory allows you to deploy release artifacts into this repository. Default is ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "handle_snapshots",
					Description: `If set, Artifactory allows you to deploy snapshot artifacts into this repository. Default is ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "suppress_pom_consistency_checks",
					Description: `By default, Artifactory keeps your repositories healthy by refusing POMs with incorrect coordinates (path). If the groupId:artifactId:version information inside the POM does not match the deployed path, Artifactory rejects the deployment with a "409 Conflict" error. You can disable this behavior by setting the Suppress POM Consistency Checks checkbox. True by default for Gradle repository.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_local_helm_repository",
			Category:         "Local Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"local",
				"repositories",
				"helm",
				"repository",
			},
			Arguments: []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `the identity key of the repo.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_local_ivy_repository",
			Category:         "Local Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"local",
				"repositories",
				"ivy",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) the identity key of the repo. ## Attribute Reference There are no unique attributes for , along with the [common list of attributes for the local repositories](local.md):`,
				},
				resource.Attribute{
					Name:        "checksum_policy_type",
					Description: `Checksum policy determines how Artifactory behaves when a client checksum for a deployed resource is missing or conflicts with the locally calculated checksum (bad checksum). The options are ` + "`" + `client-checksums` + "`" + ` and ` + "`" + `generated-checksums` + "`" + `. For more details, please refer to [Checksum Policy](https://www.jfrog.com/confluence/display/JFROG/Local+Repositories#LocalRepositories-ChecksumPolicy) .`,
				},
				resource.Attribute{
					Name:        "snapshot_version_behavior",
					Description: `Specifies the naming convention for Maven SNAPSHOT versions. The options are -`,
				},
				resource.Attribute{
					Name:        "max_unique_snapshots",
					Description: `The maximum number of unique snapshots of a single artifact to store. Once the number of snapshots exceeds this setting, older versions are removed. A value of 0 (default) indicates there is no limit, and unique snapshots are not cleaned up.`,
				},
				resource.Attribute{
					Name:        "handle_releases",
					Description: `If set, Artifactory allows you to deploy release artifacts into this repository. Default is ` + "`" + `true` + "`" + ` .`,
				},
				resource.Attribute{
					Name:        "handle_snapshots",
					Description: `If set, Artifactory allows you to deploy snapshot artifacts into this repository. Default is ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "suppress_pom_consistency_checks",
					Description: `By default, Artifactory keeps your repositories healthy by refusing POMs with incorrect coordinates (path). If the groupId:artifactId:version information inside the POM does not match the deployed path, Artifactory rejects the deployment with a "409 Conflict" error. You can disable this behavior by setting the Suppress POM Consistency Checks checkbox. True by default for Gradle repository.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "checksum_policy_type",
					Description: `Checksum policy determines how Artifactory behaves when a client checksum for a deployed resource is missing or conflicts with the locally calculated checksum (bad checksum). The options are ` + "`" + `client-checksums` + "`" + ` and ` + "`" + `generated-checksums` + "`" + `. For more details, please refer to [Checksum Policy](https://www.jfrog.com/confluence/display/JFROG/Local+Repositories#LocalRepositories-ChecksumPolicy) .`,
				},
				resource.Attribute{
					Name:        "snapshot_version_behavior",
					Description: `Specifies the naming convention for Maven SNAPSHOT versions. The options are -`,
				},
				resource.Attribute{
					Name:        "max_unique_snapshots",
					Description: `The maximum number of unique snapshots of a single artifact to store. Once the number of snapshots exceeds this setting, older versions are removed. A value of 0 (default) indicates there is no limit, and unique snapshots are not cleaned up.`,
				},
				resource.Attribute{
					Name:        "handle_releases",
					Description: `If set, Artifactory allows you to deploy release artifacts into this repository. Default is ` + "`" + `true` + "`" + ` .`,
				},
				resource.Attribute{
					Name:        "handle_snapshots",
					Description: `If set, Artifactory allows you to deploy snapshot artifacts into this repository. Default is ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "suppress_pom_consistency_checks",
					Description: `By default, Artifactory keeps your repositories healthy by refusing POMs with incorrect coordinates (path). If the groupId:artifactId:version information inside the POM does not match the deployed path, Artifactory rejects the deployment with a "409 Conflict" error. You can disable this behavior by setting the Suppress POM Consistency Checks checkbox. True by default for Gradle repository.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_local_maven_repository",
			Category:         "Local Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"local",
				"repositories",
				"maven",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) the identity key of the repo. ## Attribute Reference The following attributes are supported, along with the [common list of attributes for the local repositories](local.md):`,
				},
				resource.Attribute{
					Name:        "checksum_policy_type",
					Description: `Checksum policy determines how Artifactory behaves when a client checksum for a deployed resource is missing or conflicts with the locally calculated checksum (bad checksum). The options are: - ` + "`" + `client-checksums` + "`" + ` - ` + "`" + `server-generated-checksums` + "`" + `. For more details, please refer to [Checksum Policy](https://www.jfrog.com/confluence/display/JFROG/Local+Repositories#LocalRepositories-ChecksumPolicy) .`,
				},
				resource.Attribute{
					Name:        "snapshot_version_behavior",
					Description: `Specifies the naming convention for Maven SNAPSHOT versions. The options are -`,
				},
				resource.Attribute{
					Name:        "max_unique_snapshots",
					Description: `The maximum number of unique snapshots of a single artifact to store. Once the number of snapshots exceeds this setting, older versions are removed. A value of 0 (default) indicates there is no limit, and unique snapshots are not cleaned up.`,
				},
				resource.Attribute{
					Name:        "handle_releases",
					Description: `If set, Artifactory allows you to deploy release artifacts into this repository. Default is ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "handle_snapshots",
					Description: `If set, Artifactory allows you to deploy snapshot artifacts into this repository. Default is ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "suppress_pom_consistency_checks",
					Description: `By default, Artifactory keeps your repositories healthy by refusing POMs with incorrect coordinates (path). If the groupId:artifactId:version information inside the POM does not match the deployed path, Artifactory rejects the deployment with a "409 Conflict" error. You can disable this behavior by setting the Suppress POM Consistency Checks checkbox. False by default for Maven repository.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "checksum_policy_type",
					Description: `Checksum policy determines how Artifactory behaves when a client checksum for a deployed resource is missing or conflicts with the locally calculated checksum (bad checksum). The options are: - ` + "`" + `client-checksums` + "`" + ` - ` + "`" + `server-generated-checksums` + "`" + `. For more details, please refer to [Checksum Policy](https://www.jfrog.com/confluence/display/JFROG/Local+Repositories#LocalRepositories-ChecksumPolicy) .`,
				},
				resource.Attribute{
					Name:        "snapshot_version_behavior",
					Description: `Specifies the naming convention for Maven SNAPSHOT versions. The options are -`,
				},
				resource.Attribute{
					Name:        "max_unique_snapshots",
					Description: `The maximum number of unique snapshots of a single artifact to store. Once the number of snapshots exceeds this setting, older versions are removed. A value of 0 (default) indicates there is no limit, and unique snapshots are not cleaned up.`,
				},
				resource.Attribute{
					Name:        "handle_releases",
					Description: `If set, Artifactory allows you to deploy release artifacts into this repository. Default is ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "handle_snapshots",
					Description: `If set, Artifactory allows you to deploy snapshot artifacts into this repository. Default is ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "suppress_pom_consistency_checks",
					Description: `By default, Artifactory keeps your repositories healthy by refusing POMs with incorrect coordinates (path). If the groupId:artifactId:version information inside the POM does not match the deployed path, Artifactory rejects the deployment with a "409 Conflict" error. You can disable this behavior by setting the Suppress POM Consistency Checks checkbox. False by default for Maven repository.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_local_npm_repository",
			Category:         "Local Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"local",
				"repositories",
				"npm",
				"repository",
			},
			Arguments: []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `the identity key of the repo.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_local_nuget_repository",
			Category:         "Local Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"local",
				"repositories",
				"nuget",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) the identity key of the repo. ## Attribute Reference The following attributes are supported, along with the [common list of attributes for the local repositories](local.md):`,
				},
				resource.Attribute{
					Name:        "max_unique_snapshots",
					Description: `The maximum number of unique snapshots of a single artifact to store Once the number of snapshots exceeds this setting, older versions are removed A value of 0 (default) indicates there is no limit, and unique snapshots are not cleaned up.`,
				},
				resource.Attribute{
					Name:        "force_nuget_authentication",
					Description: `Force basic authentication credentials in order to use this repository. Default is ` + "`" + `false` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "max_unique_snapshots",
					Description: `The maximum number of unique snapshots of a single artifact to store Once the number of snapshots exceeds this setting, older versions are removed A value of 0 (default) indicates there is no limit, and unique snapshots are not cleaned up.`,
				},
				resource.Attribute{
					Name:        "force_nuget_authentication",
					Description: `Force basic authentication credentials in order to use this repository. Default is ` + "`" + `false` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_local_opkg_repository",
			Category:         "Local Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"local",
				"repositories",
				"opkg",
				"repository",
			},
			Arguments: []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `the identity key of the repo.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_local_pub_repository",
			Category:         "Local Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"local",
				"repositories",
				"pub",
				"repository",
			},
			Arguments: []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `the identity key of the repo.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_local_puppet_repository",
			Category:         "Local Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"local",
				"repositories",
				"puppet",
				"repository",
			},
			Arguments: []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `the identity key of the repo.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_local_pypi_repository",
			Category:         "Local Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"local",
				"repositories",
				"pypi",
				"repository",
			},
			Arguments: []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `the identity key of the repo.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_local_rpm_repository",
			Category:         "Local Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"local",
				"repositories",
				"rpm",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) the identity key of the repo. ## Attribute Reference The following attributes are supported, along with the [common list of attributes for the local repositories](local.md):`,
				},
				resource.Attribute{
					Name:        "yum_root_depth",
					Description: `The depth, relative to the repository's root folder, where RPM metadata is created. This is useful when your repository contains multiple RPM repositories under parallel hierarchies. For example, if your RPMs are stored under 'fedora/linux/$releasever/$basearch', specify a depth of 4. Once the number of snapshots exceeds this setting, older versions are removed. A value of 0 (default) indicates there is no limit, and unique snapshots are not cleaned up.`,
				},
				resource.Attribute{
					Name:        "calculate_yum_metadata",
					Description: `Default: ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "enable_file_lists_indexing",
					Description: `Default: ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "yum_group_file_names",
					Description: `A comma separated list of XML file names containing RPM group component definitions. Artifactory includes the group definitions as part of the calculated RPM metadata, as well as automatically generating a gzipped version of the group files, if required. Default is empty string.`,
				},
				resource.Attribute{
					Name:        "primary_keypair_ref",
					Description: `The primary GPG key to be used to sign packages.`,
				},
				resource.Attribute{
					Name:        "secondary_keypair_ref",
					Description: `The secondary GPG key to be used to sign packages.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "yum_root_depth",
					Description: `The depth, relative to the repository's root folder, where RPM metadata is created. This is useful when your repository contains multiple RPM repositories under parallel hierarchies. For example, if your RPMs are stored under 'fedora/linux/$releasever/$basearch', specify a depth of 4. Once the number of snapshots exceeds this setting, older versions are removed. A value of 0 (default) indicates there is no limit, and unique snapshots are not cleaned up.`,
				},
				resource.Attribute{
					Name:        "calculate_yum_metadata",
					Description: `Default: ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "enable_file_lists_indexing",
					Description: `Default: ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "yum_group_file_names",
					Description: `A comma separated list of XML file names containing RPM group component definitions. Artifactory includes the group definitions as part of the calculated RPM metadata, as well as automatically generating a gzipped version of the group files, if required. Default is empty string.`,
				},
				resource.Attribute{
					Name:        "primary_keypair_ref",
					Description: `The primary GPG key to be used to sign packages.`,
				},
				resource.Attribute{
					Name:        "secondary_keypair_ref",
					Description: `The secondary GPG key to be used to sign packages.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_local_sbt_repository",
			Category:         "Local Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"local",
				"repositories",
				"sbt",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) the identity key of the repo. ## Attribute Reference There are no unique attributes for , along with the [common list of attributes for the local repositories](local.md):`,
				},
				resource.Attribute{
					Name:        "checksum_policy_type",
					Description: `Checksum policy determines how Artifactory behaves when a client checksum for a deployed resource is missing or conflicts with the locally calculated checksum (bad checksum). The options are ` + "`" + `client-checksums` + "`" + ` and ` + "`" + `generated-checksums` + "`" + `. For more details, please refer to [Checksum Policy](https://www.jfrog.com/confluence/display/JFROG/Local+Repositories#LocalRepositories-ChecksumPolicy) .`,
				},
				resource.Attribute{
					Name:        "snapshot_version_behavior",
					Description: `Specifies the naming convention for Maven SNAPSHOT versions. The options are -`,
				},
				resource.Attribute{
					Name:        "max_unique_snapshots",
					Description: `The maximum number of unique snapshots of a single artifact to store. Once the number of snapshots exceeds this setting, older versions are removed. A value of 0 (default) indicates there is no limit, and unique snapshots are not cleaned up.`,
				},
				resource.Attribute{
					Name:        "handle_releases",
					Description: `If set, Artifactory allows you to deploy release artifacts into this repository. Default is ` + "`" + `true` + "`" + ` .`,
				},
				resource.Attribute{
					Name:        "handle_snapshots",
					Description: `If set, Artifactory allows you to deploy snapshot artifacts into this repository. Default is ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "suppress_pom_consistency_checks",
					Description: `By default, Artifactory keeps your repositories healthy by refusing POMs with incorrect coordinates (path). If the groupId:artifactId:version information inside the POM does not match the deployed path, Artifactory rejects the deployment with a "409 Conflict" error. You can disable this behavior by setting the Suppress POM Consistency Checks checkbox. True by default for Gradle repository.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "checksum_policy_type",
					Description: `Checksum policy determines how Artifactory behaves when a client checksum for a deployed resource is missing or conflicts with the locally calculated checksum (bad checksum). The options are ` + "`" + `client-checksums` + "`" + ` and ` + "`" + `generated-checksums` + "`" + `. For more details, please refer to [Checksum Policy](https://www.jfrog.com/confluence/display/JFROG/Local+Repositories#LocalRepositories-ChecksumPolicy) .`,
				},
				resource.Attribute{
					Name:        "snapshot_version_behavior",
					Description: `Specifies the naming convention for Maven SNAPSHOT versions. The options are -`,
				},
				resource.Attribute{
					Name:        "max_unique_snapshots",
					Description: `The maximum number of unique snapshots of a single artifact to store. Once the number of snapshots exceeds this setting, older versions are removed. A value of 0 (default) indicates there is no limit, and unique snapshots are not cleaned up.`,
				},
				resource.Attribute{
					Name:        "handle_releases",
					Description: `If set, Artifactory allows you to deploy release artifacts into this repository. Default is ` + "`" + `true` + "`" + ` .`,
				},
				resource.Attribute{
					Name:        "handle_snapshots",
					Description: `If set, Artifactory allows you to deploy snapshot artifacts into this repository. Default is ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "suppress_pom_consistency_checks",
					Description: `By default, Artifactory keeps your repositories healthy by refusing POMs with incorrect coordinates (path). If the groupId:artifactId:version information inside the POM does not match the deployed path, Artifactory rejects the deployment with a "409 Conflict" error. You can disable this behavior by setting the Suppress POM Consistency Checks checkbox. True by default for Gradle repository.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_local_swift_repository",
			Category:         "Local Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"local",
				"repositories",
				"swift",
				"repository",
			},
			Arguments: []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `the identity key of the repo.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_local_terraform_module_repository",
			Category:         "Local Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"local",
				"repositories",
				"terraform",
				"module",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `the identity key of the repo. ## Attribute Reference Attributes have a one to one mapping with the [JFrog API](https://www.jfrog.com/confluence/display/RTF/Repository+Configuration+JSON). The following arguments are supported, along with the [common list of arguments for the local repositories](local.md):`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_local_terraform_provider_repository",
			Category:         "Local Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"local",
				"repositories",
				"terraform",
				"provider",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) the identity key of the repo. ## Attribute Reference Attributes have a one to one mapping with the [JFrog API](https://www.jfrog.com/confluence/display/RTF/Repository+Configuration+JSON). The following attributes are supported, along with the [common list of arguments for the local repositories](local.md):`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `(Optional)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `(Optional)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_local_terraformbackend_repository",
			Category:         "Local Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"local",
				"repositories",
				"terraformbackend",
				"repository",
			},
			Arguments: []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `the identity key of the repo.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_local_vagrant_repository",
			Category:         "Local Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"local",
				"repositories",
				"vagrant",
				"repository",
			},
			Arguments: []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `the identity key of the repo.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_permission_target",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the permission target. ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "repo",
					Description: `Repository permission configuration.`,
				},
				resource.Attribute{
					Name:        "includes_pattern",
					Description: `Pattern of artifacts to include.`,
				},
				resource.Attribute{
					Name:        "excludes_pattern",
					Description: `Pattern of artifacts to exclude.`,
				},
				resource.Attribute{
					Name:        "repositories",
					Description: `List of repositories this permission target is applicable for. You can specify the name ` + "`" + `ANY` + "`" + ` in the repositories section in order to apply to all repositories, ` + "`" + `ANY REMOTE` + "`" + ` for all remote repositories and ` + "`" + `ANY LOCAL` + "`" + ` for all local repositories. The default value will be ` + "`" + `[]` + "`" + ` if nothing is specified.`,
				},
				resource.Attribute{
					Name:        "actions",
					Description: ``,
				},
				resource.Attribute{
					Name:        "users",
					Description: `Users this permission target applies for.`,
				},
				resource.Attribute{
					Name:        "groups",
					Description: `Groups this permission applies for.`,
				},
				resource.Attribute{
					Name:        "build",
					Description: `Same as repo but for artifactory-build-info permissions.`,
				},
				resource.Attribute{
					Name:        "release_bundle",
					Description: `Same as repo but for release-bundles permissions.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "repo",
					Description: `Repository permission configuration.`,
				},
				resource.Attribute{
					Name:        "includes_pattern",
					Description: `Pattern of artifacts to include.`,
				},
				resource.Attribute{
					Name:        "excludes_pattern",
					Description: `Pattern of artifacts to exclude.`,
				},
				resource.Attribute{
					Name:        "repositories",
					Description: `List of repositories this permission target is applicable for. You can specify the name ` + "`" + `ANY` + "`" + ` in the repositories section in order to apply to all repositories, ` + "`" + `ANY REMOTE` + "`" + ` for all remote repositories and ` + "`" + `ANY LOCAL` + "`" + ` for all local repositories. The default value will be ` + "`" + `[]` + "`" + ` if nothing is specified.`,
				},
				resource.Attribute{
					Name:        "actions",
					Description: ``,
				},
				resource.Attribute{
					Name:        "users",
					Description: `Users this permission target applies for.`,
				},
				resource.Attribute{
					Name:        "groups",
					Description: `Groups this permission applies for.`,
				},
				resource.Attribute{
					Name:        "build",
					Description: `Same as repo but for artifactory-build-info permissions.`,
				},
				resource.Attribute{
					Name:        "release_bundle",
					Description: `Same as repo but for release-bundles permissions.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_remote_alpine_repository",
			Category:         "Remote Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"remote",
				"repositories",
				"alpine",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) the identity key of the repo. ## Attribute Reference The [common list of attributes for the remote repositories](../resources/remote.md) is supported.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_remote_bower_repository",
			Category:         "Remote Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"remote",
				"repositories",
				"bower",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) the identity key of the repo. ## Attribute Reference The following attributes are supported, along with the [common list of attributes for the remote repositories](../resources/remote.md):`,
				},
				resource.Attribute{
					Name:        "bower_registry_url",
					Description: `(Optional) Proxy remote Bower repository. Default value is ` + "`" + `https://registry.bower.io` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "bower_registry_url",
					Description: `(Optional) Proxy remote Bower repository. Default value is ` + "`" + `https://registry.bower.io` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_remote_cargo_repository",
			Category:         "Remote Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"remote",
				"repositories",
				"cargo",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) the identity key of the repo. ## Attribute Reference The following attributes are supported, along with the [common list of attributes for the remote repositories](../resources/remote.md):`,
				},
				resource.Attribute{
					Name:        "anonymous_access",
					Description: `(Required) Cargo client does not send credentials when performing download and search for crates. Enable this to allow anonymous access to these resources (only), note that this will override the security anonymous access option. Default value is ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "enable_sparse_index",
					Description: `(Optional) Enable internal index support based on Cargo sparse index specifications, instead of the default git index. Default value is ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "git_registry_url",
					Description: `(Optional) This is the index url, expected to be a git repository. Default value is ` + "`" + `https://github.com/rust-lang/crates.io-index` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "anonymous_access",
					Description: `(Required) Cargo client does not send credentials when performing download and search for crates. Enable this to allow anonymous access to these resources (only), note that this will override the security anonymous access option. Default value is ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "enable_sparse_index",
					Description: `(Optional) Enable internal index support based on Cargo sparse index specifications, instead of the default git index. Default value is ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "git_registry_url",
					Description: `(Optional) This is the index url, expected to be a git repository. Default value is ` + "`" + `https://github.com/rust-lang/crates.io-index` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_remote_chef_repository",
			Category:         "Remote Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"remote",
				"repositories",
				"chef",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) the identity key of the repo. ## Attribute Reference The [common list of attributes for the remote repositories](../resources/remote.md) is supported.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_remote_cocoapods_repository",
			Category:         "Remote Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"remote",
				"repositories",
				"cocoapods",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) the identity key of the repo. ## Attribute Reference The following attributes are supported, along with the [common list of attributes for the remote repositories](../resources/remote.md):`,
				},
				resource.Attribute{
					Name:        "vcs_git_provider",
					Description: `(Optional) Artifactory supports proxying the following Git providers out-of-the-box: GitHub or a remote Artifactory instance. Default value is ` + "`" + `GITHUB` + "`" + `. Possible values are: ` + "`" + `GITHUB` + "`" + `, ` + "`" + `BITBUCKET` + "`" + `, ` + "`" + `OLDSTASH` + "`" + `, ` + "`" + `STASH` + "`" + `, ` + "`" + `ARTIFACTORY` + "`" + `, ` + "`" + `CUSTOM` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "vcs_git_download_url",
					Description: `(Optional) This attribute is used when vcs_git_provider is set to ` + "`" + `CUSTOM` + "`" + `. Provided URL will be used as proxy.`,
				},
				resource.Attribute{
					Name:        "pods_specs_repo_url",
					Description: `(Optional) Proxy remote CocoaPods Specs repositories. Default value is ` + "`" + `https://github.com/CocoaPods/Specs` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "vcs_git_provider",
					Description: `(Optional) Artifactory supports proxying the following Git providers out-of-the-box: GitHub or a remote Artifactory instance. Default value is ` + "`" + `GITHUB` + "`" + `. Possible values are: ` + "`" + `GITHUB` + "`" + `, ` + "`" + `BITBUCKET` + "`" + `, ` + "`" + `OLDSTASH` + "`" + `, ` + "`" + `STASH` + "`" + `, ` + "`" + `ARTIFACTORY` + "`" + `, ` + "`" + `CUSTOM` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "vcs_git_download_url",
					Description: `(Optional) This attribute is used when vcs_git_provider is set to ` + "`" + `CUSTOM` + "`" + `. Provided URL will be used as proxy.`,
				},
				resource.Attribute{
					Name:        "pods_specs_repo_url",
					Description: `(Optional) Proxy remote CocoaPods Specs repositories. Default value is ` + "`" + `https://github.com/CocoaPods/Specs` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_remote_composer_repository",
			Category:         "Remote Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"remote",
				"repositories",
				"composer",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) the identity key of the repo. ## Attribute Reference The following attributes are supported, along with the [common list of attributes for the remote repositories](../resources/remote.md):`,
				},
				resource.Attribute{
					Name:        "vcs_git_provider",
					Description: `(Optional) Artifactory supports proxying the following Git providers out-of-the-box: GitHub or a remote Artifactory instance. Default value is ` + "`" + `GITHUB` + "`" + `. Possible values are: ` + "`" + `GITHUB` + "`" + `, ` + "`" + `BITBUCKET` + "`" + `, ` + "`" + `OLDSTASH` + "`" + `, ` + "`" + `STASH` + "`" + `, ` + "`" + `ARTIFACTORY` + "`" + `, ` + "`" + `CUSTOM` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "vcs_git_download_url",
					Description: `(Optional) This attribute is used when vcs_git_provider is set to ` + "`" + `CUSTOM` + "`" + `. Provided URL will be used as proxy.`,
				},
				resource.Attribute{
					Name:        "composer_registry_url",
					Description: `(Optional) Proxy remote Composer repository. Default value is ` + "`" + `https://packagist.org` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "vcs_git_provider",
					Description: `(Optional) Artifactory supports proxying the following Git providers out-of-the-box: GitHub or a remote Artifactory instance. Default value is ` + "`" + `GITHUB` + "`" + `. Possible values are: ` + "`" + `GITHUB` + "`" + `, ` + "`" + `BITBUCKET` + "`" + `, ` + "`" + `OLDSTASH` + "`" + `, ` + "`" + `STASH` + "`" + `, ` + "`" + `ARTIFACTORY` + "`" + `, ` + "`" + `CUSTOM` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "vcs_git_download_url",
					Description: `(Optional) This attribute is used when vcs_git_provider is set to ` + "`" + `CUSTOM` + "`" + `. Provided URL will be used as proxy.`,
				},
				resource.Attribute{
					Name:        "composer_registry_url",
					Description: `(Optional) Proxy remote Composer repository. Default value is ` + "`" + `https://packagist.org` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_remote_conan_repository",
			Category:         "Remote Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"remote",
				"repositories",
				"conan",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) the identity key of the repo. ## Attribute Reference The following attributes are supported, along with the [common list of attributes for the remote repositories](../resources/remote.md):`,
				},
				resource.Attribute{
					Name:        "force_conan_authentication",
					Description: `(Optional) Force basic authentication credentials in order to use this repository. Default value is ` + "`" + `false` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "force_conan_authentication",
					Description: `(Optional) Force basic authentication credentials in order to use this repository. Default value is ` + "`" + `false` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_remote_conda_repository",
			Category:         "Remote Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"remote",
				"repositories",
				"conda",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) the identity key of the repo. ## Attribute Reference The [common list of attributes for the remote repositories](../resources/remote.md) is supported.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_remote_cran_repository",
			Category:         "Remote Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"remote",
				"repositories",
				"cran",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) the identity key of the repo. ## Attribute Reference The [common list of attributes for the remote repositories](../resources/remote.md) is supported.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_remote_debian_repository",
			Category:         "Remote Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"remote",
				"repositories",
				"debian",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) the identity key of the repo. ## Attribute Reference The [common list of attributes for the remote repositories](../resources/remote.md) is supported.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_remote_docker_repository",
			Category:         "Remote Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"remote",
				"repositories",
				"docker",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) the identity key of the repo. ## Attribute Reference The following attributes are supported, along with the [common list of attributes for the remote repositories](../resources/remote.md):`,
				},
				resource.Attribute{
					Name:        "block_pushing_schema1",
					Description: `(Optional) When set, Artifactory will block the pulling of Docker images with manifest v2 schema 1 from the remote repository (i.e. the upstream). It will be possible to pull images with manifest v2 schema 1 that exist in the cache.`,
				},
				resource.Attribute{
					Name:        "enable_token_authentication",
					Description: `(Optional) Enable token (Bearer) based authentication.`,
				},
				resource.Attribute{
					Name:        "external_dependencies_enabled",
					Description: `(Optional) Also known as 'Foreign Layers Caching' on the UI.`,
				},
				resource.Attribute{
					Name:        "external_dependencies_patterns",
					Description: `(Optional) An allow list of Ant-style path patterns that determine which remote VCS roots Artifactory will follow to download remote modules from, when presented with 'go-import' meta tags in the remote repository response. By default, this is set to ` + "`" + `[`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "block_pushing_schema1",
					Description: `(Optional) When set, Artifactory will block the pulling of Docker images with manifest v2 schema 1 from the remote repository (i.e. the upstream). It will be possible to pull images with manifest v2 schema 1 that exist in the cache.`,
				},
				resource.Attribute{
					Name:        "enable_token_authentication",
					Description: `(Optional) Enable token (Bearer) based authentication.`,
				},
				resource.Attribute{
					Name:        "external_dependencies_enabled",
					Description: `(Optional) Also known as 'Foreign Layers Caching' on the UI.`,
				},
				resource.Attribute{
					Name:        "external_dependencies_patterns",
					Description: `(Optional) An allow list of Ant-style path patterns that determine which remote VCS roots Artifactory will follow to download remote modules from, when presented with 'go-import' meta tags in the remote repository response. By default, this is set to ` + "`" + `[`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_remote_gems_repository",
			Category:         "Remote Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"remote",
				"repositories",
				"gems",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) the identity key of the repo. ## Attribute Reference The [common list of attributes for the remote repositories](../resources/remote.md) is supported.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_remote_generic_repository",
			Category:         "Remote Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"remote",
				"repositories",
				"generic",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) the identity key of the repo. ## Attribute Reference The following attributes are supported, along with the [common list of attributes for the remote repositories](../resources/remote.md):`,
				},
				resource.Attribute{
					Name:        "propagate_query_params",
					Description: `(Optional, Default: ` + "`" + `false` + "`" + `) When set, if query params are included in the request to Artifactory, they will be passed on to the remote repository.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "propagate_query_params",
					Description: `(Optional, Default: ` + "`" + `false` + "`" + `) When set, if query params are included in the request to Artifactory, they will be passed on to the remote repository.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_remote_gitlfs_repository",
			Category:         "Remote Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"remote",
				"repositories",
				"gitlfs",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) the identity key of the repo. ## Attribute Reference The [common list of attributes for the remote repositories](../resources/remote.md) is supported.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_remote_go_repository",
			Category:         "Remote Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"remote",
				"repositories",
				"go",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) the identity key of the repo. ## Attribute Reference The following attributes are supported, along with the [common list of attributes for the remote repositories](../resources/remote.md):`,
				},
				resource.Attribute{
					Name:        "vcs_git_provider",
					Description: `(Optional) Artifactory supports proxying the following Git providers out-of-the-box: GitHub or a remote Artifactory instance. Default value is ` + "`" + `ARTIFACTORY` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "vcs_git_provider",
					Description: `(Optional) Artifactory supports proxying the following Git providers out-of-the-box: GitHub or a remote Artifactory instance. Default value is ` + "`" + `ARTIFACTORY` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_remote_gradle_repository",
			Category:         "Remote Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"remote",
				"repositories",
				"gradle",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) the identity key of the repo. ## Attribute Reference The following attributes are supported, along with the [common list of attributes for the remote repositories](../resources/remote.md):`,
				},
				resource.Attribute{
					Name:        "fetch_jars_eagerly",
					Description: `(Optional, Default: ` + "`" + `false` + "`" + `) When set, if a POM is requested, Artifactory attempts to fetch the corresponding jar in the background. This will accelerate first access time to the jar when it is subsequently requested.`,
				},
				resource.Attribute{
					Name:        "fetch_sources_eagerly",
					Description: `(Optional, Default: ` + "`" + `false` + "`" + `) When set, if a binaries jar is requested, Artifactory attempts to fetch the corresponding source jar in the background. This will accelerate first access time to the source jar when it is subsequently requested.`,
				},
				resource.Attribute{
					Name:        "handle_releases",
					Description: `(Optional, Default: ` + "`" + `true` + "`" + `) If set, Artifactory allows you to deploy release artifacts into this repository.`,
				},
				resource.Attribute{
					Name:        "handle_snapshots",
					Description: `(Optional, Default: ` + "`" + `true` + "`" + `) If set, Artifactory allows you to deploy snapshot artifacts into this repository.`,
				},
				resource.Attribute{
					Name:        "suppress_pom_consistency_checks",
					Description: `(Optional, Default: ` + "`" + `true` + "`" + `) - By default, the system keeps your repositories healthy by refusing POMs with incorrect coordinates (path). If the groupId:artifactId:version information inside the POM does not match the deployed path, Artifactory rejects the deployment with a "409 Conflict" error. You can disable this behavior by setting this attribute to ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "reject_invalid_jars",
					Description: `(Optional, Default: ` + "`" + `false` + "`" + `) Reject the caching of jar files that are found to be invalid. For example, pseudo jars retrieved behind a "captive portal".`,
				},
				resource.Attribute{
					Name:        "remote_repo_checksum_policy_type",
					Description: `(Optional, Default: ` + "`" + `generate-if-absent` + "`" + `) Checking the Checksum effectively verifies the integrity of a deployed resource. The Checksum Policy determines how the system behaves when a client checksum for a remote resource is missing or conflicts with the locally calculated checksum. Available policies are ` + "`" + `generate-if-absent` + "`" + `, ` + "`" + `fail` + "`" + `, ` + "`" + `ignore-and-generate` + "`" + `, and ` + "`" + `pass-thru` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "fetch_jars_eagerly",
					Description: `(Optional, Default: ` + "`" + `false` + "`" + `) When set, if a POM is requested, Artifactory attempts to fetch the corresponding jar in the background. This will accelerate first access time to the jar when it is subsequently requested.`,
				},
				resource.Attribute{
					Name:        "fetch_sources_eagerly",
					Description: `(Optional, Default: ` + "`" + `false` + "`" + `) When set, if a binaries jar is requested, Artifactory attempts to fetch the corresponding source jar in the background. This will accelerate first access time to the source jar when it is subsequently requested.`,
				},
				resource.Attribute{
					Name:        "handle_releases",
					Description: `(Optional, Default: ` + "`" + `true` + "`" + `) If set, Artifactory allows you to deploy release artifacts into this repository.`,
				},
				resource.Attribute{
					Name:        "handle_snapshots",
					Description: `(Optional, Default: ` + "`" + `true` + "`" + `) If set, Artifactory allows you to deploy snapshot artifacts into this repository.`,
				},
				resource.Attribute{
					Name:        "suppress_pom_consistency_checks",
					Description: `(Optional, Default: ` + "`" + `true` + "`" + `) - By default, the system keeps your repositories healthy by refusing POMs with incorrect coordinates (path). If the groupId:artifactId:version information inside the POM does not match the deployed path, Artifactory rejects the deployment with a "409 Conflict" error. You can disable this behavior by setting this attribute to ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "reject_invalid_jars",
					Description: `(Optional, Default: ` + "`" + `false` + "`" + `) Reject the caching of jar files that are found to be invalid. For example, pseudo jars retrieved behind a "captive portal".`,
				},
				resource.Attribute{
					Name:        "remote_repo_checksum_policy_type",
					Description: `(Optional, Default: ` + "`" + `generate-if-absent` + "`" + `) Checking the Checksum effectively verifies the integrity of a deployed resource. The Checksum Policy determines how the system behaves when a client checksum for a remote resource is missing or conflicts with the locally calculated checksum. Available policies are ` + "`" + `generate-if-absent` + "`" + `, ` + "`" + `fail` + "`" + `, ` + "`" + `ignore-and-generate` + "`" + `, and ` + "`" + `pass-thru` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_remote_helm_repository",
			Category:         "Remote Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"remote",
				"repositories",
				"helm",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) the identity key of the repo. ## Attribute Reference The following attributes are supported, along with the [common list of attributes for the remote repositories](../resources/remote.md):`,
				},
				resource.Attribute{
					Name:        "helm_charts_base_url",
					Description: `(Optional) No documentation is available. Hopefully you know what this means.`,
				},
				resource.Attribute{
					Name:        "external_dependencies_enabled",
					Description: `(Optional) When set, external dependencies are rewritten. ` + "`" + `External Dependency Rewrite` + "`" + ` in the UI.`,
				},
				resource.Attribute{
					Name:        "external_dependencies_patterns",
					Description: `(Optional) An allow list of Ant-style path patterns that determine which remote VCS roots Artifactory will follow to download remote modules from, when presented with 'go-import' meta tags in the remote repository response. By default, this is set to ` + "`" + `[`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "helm_charts_base_url",
					Description: `(Optional) No documentation is available. Hopefully you know what this means.`,
				},
				resource.Attribute{
					Name:        "external_dependencies_enabled",
					Description: `(Optional) When set, external dependencies are rewritten. ` + "`" + `External Dependency Rewrite` + "`" + ` in the UI.`,
				},
				resource.Attribute{
					Name:        "external_dependencies_patterns",
					Description: `(Optional) An allow list of Ant-style path patterns that determine which remote VCS roots Artifactory will follow to download remote modules from, when presented with 'go-import' meta tags in the remote repository response. By default, this is set to ` + "`" + `[`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_remote_ivy_repository",
			Category:         "Remote Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"remote",
				"repositories",
				"ivy",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) the identity key of the repo. ## Attribute Reference The following attributes are supported, along with the [common list of attributes for the remote repositories](../resources/remote.md):`,
				},
				resource.Attribute{
					Name:        "fetch_jars_eagerly",
					Description: `(Optional, Default: ` + "`" + `false` + "`" + `) When set, if a POM is requested, Artifactory attempts to fetch the corresponding jar in the background. This will accelerate first access time to the jar when it is subsequently requested.`,
				},
				resource.Attribute{
					Name:        "fetch_sources_eagerly",
					Description: `(Optional, Default: ` + "`" + `false` + "`" + `) When set, if a binaries jar is requested, Artifactory attempts to fetch the corresponding source jar in the background. This will accelerate first access time to the source jar when it is subsequently requested.`,
				},
				resource.Attribute{
					Name:        "handle_releases",
					Description: `(Optional, Default: ` + "`" + `true` + "`" + `) If set, Artifactory allows you to deploy release artifacts into this repository.`,
				},
				resource.Attribute{
					Name:        "handle_snapshots",
					Description: `(Optional, Default: ` + "`" + `true` + "`" + `) If set, Artifactory allows you to deploy snapshot artifacts into this repository.`,
				},
				resource.Attribute{
					Name:        "suppress_pom_consistency_checks",
					Description: `(Optional, Default: ` + "`" + `true` + "`" + `) - By default, the system keeps your repositories healthy by refusing POMs with incorrect coordinates (path). If the groupId:artifactId:version information inside the POM does not match the deployed path, Artifactory rejects the deployment with a "409 Conflict" error. You can disable this behavior by setting this attribute to ` + "`" + `true` + "`" + `'`,
				},
				resource.Attribute{
					Name:        "reject_invalid_jars",
					Description: `(Optional, Default: ` + "`" + `false` + "`" + `) Reject the caching of jar files that are found to be invalid. For example, pseudo jars retrieved behind a "captive portal".`,
				},
				resource.Attribute{
					Name:        "remote_repo_checksum_policy_type",
					Description: `(Optional, Default: ` + "`" + `generate-if-absent` + "`" + `) Checking the Checksum effectively verifies the integrity of a deployed resource. The Checksum Policy determines how the system behaves when a client checksum for a remote resource is missing or conflicts with the locally calculated checksum. Available policies are ` + "`" + `generate-if-absent` + "`" + `, ` + "`" + `fail` + "`" + `, ` + "`" + `ignore-and-generate` + "`" + `, and ` + "`" + `pass-thru` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "fetch_jars_eagerly",
					Description: `(Optional, Default: ` + "`" + `false` + "`" + `) When set, if a POM is requested, Artifactory attempts to fetch the corresponding jar in the background. This will accelerate first access time to the jar when it is subsequently requested.`,
				},
				resource.Attribute{
					Name:        "fetch_sources_eagerly",
					Description: `(Optional, Default: ` + "`" + `false` + "`" + `) When set, if a binaries jar is requested, Artifactory attempts to fetch the corresponding source jar in the background. This will accelerate first access time to the source jar when it is subsequently requested.`,
				},
				resource.Attribute{
					Name:        "handle_releases",
					Description: `(Optional, Default: ` + "`" + `true` + "`" + `) If set, Artifactory allows you to deploy release artifacts into this repository.`,
				},
				resource.Attribute{
					Name:        "handle_snapshots",
					Description: `(Optional, Default: ` + "`" + `true` + "`" + `) If set, Artifactory allows you to deploy snapshot artifacts into this repository.`,
				},
				resource.Attribute{
					Name:        "suppress_pom_consistency_checks",
					Description: `(Optional, Default: ` + "`" + `true` + "`" + `) - By default, the system keeps your repositories healthy by refusing POMs with incorrect coordinates (path). If the groupId:artifactId:version information inside the POM does not match the deployed path, Artifactory rejects the deployment with a "409 Conflict" error. You can disable this behavior by setting this attribute to ` + "`" + `true` + "`" + `'`,
				},
				resource.Attribute{
					Name:        "reject_invalid_jars",
					Description: `(Optional, Default: ` + "`" + `false` + "`" + `) Reject the caching of jar files that are found to be invalid. For example, pseudo jars retrieved behind a "captive portal".`,
				},
				resource.Attribute{
					Name:        "remote_repo_checksum_policy_type",
					Description: `(Optional, Default: ` + "`" + `generate-if-absent` + "`" + `) Checking the Checksum effectively verifies the integrity of a deployed resource. The Checksum Policy determines how the system behaves when a client checksum for a remote resource is missing or conflicts with the locally calculated checksum. Available policies are ` + "`" + `generate-if-absent` + "`" + `, ` + "`" + `fail` + "`" + `, ` + "`" + `ignore-and-generate` + "`" + `, and ` + "`" + `pass-thru` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_remote_maven_repository",
			Category:         "Remote Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"remote",
				"repositories",
				"maven",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) the identity key of the repo. ## Attribute Reference The following attributes are supported, along with the [common list of attributes for the remote repositories](../resources/remote.md):`,
				},
				resource.Attribute{
					Name:        "fetch_jars_eagerly",
					Description: `(Optional, Default: ` + "`" + `false` + "`" + `) When set, if a POM is requested, Artifactory attempts to fetch the corresponding jar in the background. This will accelerate first access time to the jar when it is subsequently requested.`,
				},
				resource.Attribute{
					Name:        "fetch_sources_eagerly",
					Description: `(Optional, Default: ` + "`" + `false` + "`" + `) When set, if a binaries jar is requested, Artifactory attempts to fetch the corresponding source jar in the background. This will accelerate first access time to the source jar when it is subsequently requested.`,
				},
				resource.Attribute{
					Name:        "handle_releases",
					Description: `(Optional, Default: ` + "`" + `true` + "`" + `) If set, Artifactory allows you to deploy release artifacts into this repository.`,
				},
				resource.Attribute{
					Name:        "handle_snapshots",
					Description: `(Optional, Default: ` + "`" + `true` + "`" + `) If set, Artifactory allows you to deploy snapshot artifacts into this repository.`,
				},
				resource.Attribute{
					Name:        "suppress_pom_consistency_checks",
					Description: `(Optional, Default: ` + "`" + `true` + "`" + `) By default, the system keeps your repositories healthy by refusing POMs with incorrect coordinates (path). If the groupId:artifactId:version information inside the POM does not match the deployed path, Artifactory rejects the deployment with a "409 Conflict" error. You can disable this behavior by setting this attribute to ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "reject_invalid_jars",
					Description: `(Optional, Default: ` + "`" + `false` + "`" + `) Reject the caching of jar files that are found to be invalid. For example, pseudo jars retrieved behind a "captive portal".`,
				},
				resource.Attribute{
					Name:        "remote_repo_checksum_policy_type",
					Description: `(Optional, Default: ` + "`" + `generate-if-absent` + "`" + `) Checking the Checksum effectively verifies the integrity of a deployed resource. The Checksum Policy determines how the system behaves when a client checksum for a remote resource is missing or conflicts with the locally calculated checksum. Available policies are ` + "`" + `generate-if-absent` + "`" + `, ` + "`" + `fail` + "`" + `, ` + "`" + `ignore-and-generate` + "`" + `, and ` + "`" + `pass-thru` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "metadata_retrieval_timeout_secs",
					Description: `(Optional, Default: 60) This value refers to the number of seconds to cache metadata files before checking for newer versions on remote server. A value of 0 indicates no caching. Cannot be larger than ` + "`" + `retrieval_cache_period_seconds` + "`" + ` attribute.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "fetch_jars_eagerly",
					Description: `(Optional, Default: ` + "`" + `false` + "`" + `) When set, if a POM is requested, Artifactory attempts to fetch the corresponding jar in the background. This will accelerate first access time to the jar when it is subsequently requested.`,
				},
				resource.Attribute{
					Name:        "fetch_sources_eagerly",
					Description: `(Optional, Default: ` + "`" + `false` + "`" + `) When set, if a binaries jar is requested, Artifactory attempts to fetch the corresponding source jar in the background. This will accelerate first access time to the source jar when it is subsequently requested.`,
				},
				resource.Attribute{
					Name:        "handle_releases",
					Description: `(Optional, Default: ` + "`" + `true` + "`" + `) If set, Artifactory allows you to deploy release artifacts into this repository.`,
				},
				resource.Attribute{
					Name:        "handle_snapshots",
					Description: `(Optional, Default: ` + "`" + `true` + "`" + `) If set, Artifactory allows you to deploy snapshot artifacts into this repository.`,
				},
				resource.Attribute{
					Name:        "suppress_pom_consistency_checks",
					Description: `(Optional, Default: ` + "`" + `true` + "`" + `) By default, the system keeps your repositories healthy by refusing POMs with incorrect coordinates (path). If the groupId:artifactId:version information inside the POM does not match the deployed path, Artifactory rejects the deployment with a "409 Conflict" error. You can disable this behavior by setting this attribute to ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "reject_invalid_jars",
					Description: `(Optional, Default: ` + "`" + `false` + "`" + `) Reject the caching of jar files that are found to be invalid. For example, pseudo jars retrieved behind a "captive portal".`,
				},
				resource.Attribute{
					Name:        "remote_repo_checksum_policy_type",
					Description: `(Optional, Default: ` + "`" + `generate-if-absent` + "`" + `) Checking the Checksum effectively verifies the integrity of a deployed resource. The Checksum Policy determines how the system behaves when a client checksum for a remote resource is missing or conflicts with the locally calculated checksum. Available policies are ` + "`" + `generate-if-absent` + "`" + `, ` + "`" + `fail` + "`" + `, ` + "`" + `ignore-and-generate` + "`" + `, and ` + "`" + `pass-thru` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "metadata_retrieval_timeout_secs",
					Description: `(Optional, Default: 60) This value refers to the number of seconds to cache metadata files before checking for newer versions on remote server. A value of 0 indicates no caching. Cannot be larger than ` + "`" + `retrieval_cache_period_seconds` + "`" + ` attribute.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_remote_npm_repository",
			Category:         "Remote Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"remote",
				"repositories",
				"npm",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) the identity key of the repo. ## Attribute Reference The [common list of attributes for the remote repositories](../resources/remote.md) is supported.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_remote_nuget_repository",
			Category:         "Remote Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"remote",
				"repositories",
				"nuget",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) the identity key of the repo. ## Attribute Reference The following attributes are supported, along with the [common list of attributes for the remote repositories](../resources/remote.md):`,
				},
				resource.Attribute{
					Name:        "feed_context_path",
					Description: `(Optional) When proxying a remote NuGet repository, customize feed resource location using this attribute. Default value is ` + "`" + `api/v2` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "download_context_path",
					Description: `(Optional) The context path prefix through which NuGet downloads are served. For example, the NuGet Gallery download URL is ` + "`" + `https://nuget.org/api/v2/package` + "`" + `, so the repository URL should be configured as ` + "`" + `https://nuget.org` + "`" + ` and the download context path should be configured as ` + "`" + `api/v2/package` + "`" + `. Default value is ` + "`" + `api/v2/package` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "v3_feed_url",
					Description: `(Optional) The URL to the NuGet v3 feed. Default value is ` + "`" + `https://api.nuget.org/v3/index.json` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "force_nuget_authentication",
					Description: `(Optional) Force basic authentication credentials in order to use this repository. Default value is ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "symbol_server_url",
					Description: `(Optional) NuGet symbol server URL. Default value is ` + "`" + `https://symbols.nuget.org/download/symbols` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "feed_context_path",
					Description: `(Optional) When proxying a remote NuGet repository, customize feed resource location using this attribute. Default value is ` + "`" + `api/v2` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "download_context_path",
					Description: `(Optional) The context path prefix through which NuGet downloads are served. For example, the NuGet Gallery download URL is ` + "`" + `https://nuget.org/api/v2/package` + "`" + `, so the repository URL should be configured as ` + "`" + `https://nuget.org` + "`" + ` and the download context path should be configured as ` + "`" + `api/v2/package` + "`" + `. Default value is ` + "`" + `api/v2/package` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "v3_feed_url",
					Description: `(Optional) The URL to the NuGet v3 feed. Default value is ` + "`" + `https://api.nuget.org/v3/index.json` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "force_nuget_authentication",
					Description: `(Optional) Force basic authentication credentials in order to use this repository. Default value is ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "symbol_server_url",
					Description: `(Optional) NuGet symbol server URL. Default value is ` + "`" + `https://symbols.nuget.org/download/symbols` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_remote_opkg_repository",
			Category:         "Remote Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"remote",
				"repositories",
				"opkg",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) the identity key of the repo. ## Attribute Reference The [common list of attributes for the remote repositories](../resources/remote.md) is supported.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_remote_p2_repository",
			Category:         "Remote Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"remote",
				"repositories",
				"p2",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) the identity key of the repo. ## Attribute Reference The [common list of attributes for the remote repositories](../resources/remote.md) is supported.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_remote_pub_repository",
			Category:         "Remote Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"remote",
				"repositories",
				"pub",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) the identity key of the repo. ## Attribute Reference The [common list of attributes for the remote repositories](../resources/remote.md) is supported.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_remote_puppet_repository",
			Category:         "Remote Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"remote",
				"repositories",
				"puppet",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) the identity key of the repo. ## Attribute Reference The [common list of attributes for the remote repositories](../resources/remote.md) is supported.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_remote_pypi_repository",
			Category:         "Remote Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"remote",
				"repositories",
				"pypi",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) the identity key of the repo. ## Attribute Reference The following attributes are supported, along with the [common list of attributes for the remote repositories](../resources/remote.md):`,
				},
				resource.Attribute{
					Name:        "pypi_registry_url",
					Description: `(Optional) To configure the remote repo to proxy public external PyPI repository, or a PyPI repository hosted on another Artifactory server. See JFrog Pypi documentation [here](https://www.jfrog.com/confluence/display/JFROG/PyPI+Repositories) for the usage details. Default value is ` + "`" + `https://pypi.org` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "pypi_repository_suffix",
					Description: `(Optional) Usually should be left as a default for ` + "`" + `simple` + "`" + `, unless the remote is a PyPI server that has custom registry suffix, like +simple in DevPI. Default value is ` + "`" + `simple` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "pypi_registry_url",
					Description: `(Optional) To configure the remote repo to proxy public external PyPI repository, or a PyPI repository hosted on another Artifactory server. See JFrog Pypi documentation [here](https://www.jfrog.com/confluence/display/JFROG/PyPI+Repositories) for the usage details. Default value is ` + "`" + `https://pypi.org` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "pypi_repository_suffix",
					Description: `(Optional) Usually should be left as a default for ` + "`" + `simple` + "`" + `, unless the remote is a PyPI server that has custom registry suffix, like +simple in DevPI. Default value is ` + "`" + `simple` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_remote_rpm_repository",
			Category:         "Remote Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"remote",
				"repositories",
				"rpm",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) the identity key of the repo. ## Attribute Reference The [common list of attributes for the remote repositories](../resources/remote.md) is supported.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_remote_sbt_repository",
			Category:         "Remote Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"remote",
				"repositories",
				"sbt",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) the identity key of the repo. ## Attribute Reference The following attributes are supported, along with the [common list of attributes for the remote repositories](../resources/remote.md):`,
				},
				resource.Attribute{
					Name:        "fetch_jars_eagerly",
					Description: `(Optional, Default: ` + "`" + `false` + "`" + `) When set, if a POM is requested, Artifactory attempts to fetch the corresponding jar in the background. This will accelerate first access time to the jar when it is subsequently requested.`,
				},
				resource.Attribute{
					Name:        "fetch_sources_eagerly",
					Description: `(Optional, Default: ` + "`" + `false` + "`" + `) - When set, if a binaries jar is requested, Artifactory attempts to fetch the corresponding source jar in the background. This will accelerate first access time to the source jar when it is subsequently requested.`,
				},
				resource.Attribute{
					Name:        "handle_releases",
					Description: `(Optional, Default: ` + "`" + `true` + "`" + `) If set, Artifactory allows you to deploy release artifacts into this repository.`,
				},
				resource.Attribute{
					Name:        "handle_snapshots",
					Description: `(Optional, Default: ` + "`" + `true` + "`" + `) If set, Artifactory allows you to deploy snapshot artifacts into this repository.`,
				},
				resource.Attribute{
					Name:        "suppress_pom_consistency_checks",
					Description: `(Optional, Default: ` + "`" + `true` + "`" + `) By default, the system keeps your repositories healthy by refusing POMs with incorrect coordinates (path). If the groupId:artifactId:version information inside the POM does not match the deployed path, Artifactory rejects the deployment with a "409 Conflict" error. You can disable this behavior by setting this attribute to ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "reject_invalid_jars",
					Description: `(Optional, Default: ` + "`" + `false` + "`" + `) Reject the caching of jar files that are found to be invalid. For example, pseudo jars retrieved behind a "captive portal".`,
				},
				resource.Attribute{
					Name:        "remote_repo_checksum_policy_type",
					Description: `(Optional, Default: ` + "`" + `generate-if-absent` + "`" + `) Checking the Checksum effectively verifies the integrity of a deployed resource. The Checksum Policy determines how the system behaves when a client checksum for a remote resource is missing or conflicts with the locally calculated checksum. Available policies are ` + "`" + `generate-if-absent` + "`" + `, ` + "`" + `fail` + "`" + `, ` + "`" + `ignore-and-generate` + "`" + `, and ` + "`" + `pass-thru` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "fetch_jars_eagerly",
					Description: `(Optional, Default: ` + "`" + `false` + "`" + `) When set, if a POM is requested, Artifactory attempts to fetch the corresponding jar in the background. This will accelerate first access time to the jar when it is subsequently requested.`,
				},
				resource.Attribute{
					Name:        "fetch_sources_eagerly",
					Description: `(Optional, Default: ` + "`" + `false` + "`" + `) - When set, if a binaries jar is requested, Artifactory attempts to fetch the corresponding source jar in the background. This will accelerate first access time to the source jar when it is subsequently requested.`,
				},
				resource.Attribute{
					Name:        "handle_releases",
					Description: `(Optional, Default: ` + "`" + `true` + "`" + `) If set, Artifactory allows you to deploy release artifacts into this repository.`,
				},
				resource.Attribute{
					Name:        "handle_snapshots",
					Description: `(Optional, Default: ` + "`" + `true` + "`" + `) If set, Artifactory allows you to deploy snapshot artifacts into this repository.`,
				},
				resource.Attribute{
					Name:        "suppress_pom_consistency_checks",
					Description: `(Optional, Default: ` + "`" + `true` + "`" + `) By default, the system keeps your repositories healthy by refusing POMs with incorrect coordinates (path). If the groupId:artifactId:version information inside the POM does not match the deployed path, Artifactory rejects the deployment with a "409 Conflict" error. You can disable this behavior by setting this attribute to ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "reject_invalid_jars",
					Description: `(Optional, Default: ` + "`" + `false` + "`" + `) Reject the caching of jar files that are found to be invalid. For example, pseudo jars retrieved behind a "captive portal".`,
				},
				resource.Attribute{
					Name:        "remote_repo_checksum_policy_type",
					Description: `(Optional, Default: ` + "`" + `generate-if-absent` + "`" + `) Checking the Checksum effectively verifies the integrity of a deployed resource. The Checksum Policy determines how the system behaves when a client checksum for a remote resource is missing or conflicts with the locally calculated checksum. Available policies are ` + "`" + `generate-if-absent` + "`" + `, ` + "`" + `fail` + "`" + `, ` + "`" + `ignore-and-generate` + "`" + `, and ` + "`" + `pass-thru` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_remote_swift_repository",
			Category:         "Remote Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"remote",
				"repositories",
				"swift",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) the identity key of the repo. ## Attribute Reference The [common list of attributes for the remote repositories](../resources/remote.md) is supported.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_remote_terraform_repository",
			Category:         "Remote Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"remote",
				"repositories",
				"terraform",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) the identity key of the repo. ## Attribute Reference The following attributes are supported, along with the [common list of attributes for the remote repositories](../resources/remote.md):`,
				},
				resource.Attribute{
					Name:        "terraform_registry_url",
					Description: `(Optional) The base URL of the registry API. When using Smart Remote Repositories, set the URL to ` + "`" + `<base_Artifactory_URL>/api/terraform/repokey` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "terraform_providers_url",
					Description: `(Optional) The base URL of the Provider's storage API. When using Smart remote repositories, set the URL to ` + "`" + `<base_Artifactory_URL>/api/terraform/repokey/providers` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "terraform_registry_url",
					Description: `(Optional) The base URL of the registry API. When using Smart Remote Repositories, set the URL to ` + "`" + `<base_Artifactory_URL>/api/terraform/repokey` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "terraform_providers_url",
					Description: `(Optional) The base URL of the Provider's storage API. When using Smart remote repositories, set the URL to ` + "`" + `<base_Artifactory_URL>/api/terraform/repokey/providers` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_remote_vcs_repository",
			Category:         "Remote Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"remote",
				"repositories",
				"vcs",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) the identity key of the repo. ## Attribute Reference The following attributes are supported, along with the [common list of attributes for the remote repositories](../resources/remote.md):`,
				},
				resource.Attribute{
					Name:        "vcs_git_provider",
					Description: `(Optional) Artifactory supports proxying the following Git providers out-of-the-box: GitHub, Bitbucket, Stash, a remote Artifactory instance or a custom Git repository. Allowed values are: ` + "`" + `GITHUB` + "`" + `, ` + "`" + `BITBUCKET` + "`" + `, ` + "`" + `OLDSTASH` + "`" + `, ` + "`" + `STASH` + "`" + `, ` + "`" + `ARTIFACTORY` + "`" + `, ` + "`" + `CUSTOM` + "`" + `. Default value is ` + "`" + `GITHUB` + "`" + ``,
				},
				resource.Attribute{
					Name:        "vcs_git_download_url",
					Description: `(Optional) This attribute is used when vcs_git_provider is set to ` + "`" + `CUSTOM` + "`" + `. Provided URL will be used as proxy.`,
				},
				resource.Attribute{
					Name:        "max_unique_snapshots",
					Description: `(Optional) The maximum number of unique snapshots of a single artifact to store. Once the number of snapshots exceeds this setting, older versions are removed. A value of 0 (default) indicates there is no limit, and unique snapshots are not cleaned up.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "vcs_git_provider",
					Description: `(Optional) Artifactory supports proxying the following Git providers out-of-the-box: GitHub, Bitbucket, Stash, a remote Artifactory instance or a custom Git repository. Allowed values are: ` + "`" + `GITHUB` + "`" + `, ` + "`" + `BITBUCKET` + "`" + `, ` + "`" + `OLDSTASH` + "`" + `, ` + "`" + `STASH` + "`" + `, ` + "`" + `ARTIFACTORY` + "`" + `, ` + "`" + `CUSTOM` + "`" + `. Default value is ` + "`" + `GITHUB` + "`" + ``,
				},
				resource.Attribute{
					Name:        "vcs_git_download_url",
					Description: `(Optional) This attribute is used when vcs_git_provider is set to ` + "`" + `CUSTOM` + "`" + `. Provided URL will be used as proxy.`,
				},
				resource.Attribute{
					Name:        "max_unique_snapshots",
					Description: `(Optional) The maximum number of unique snapshots of a single artifact to store. Once the number of snapshots exceeds this setting, older versions are removed. A value of 0 (default) indicates there is no limit, and unique snapshots are not cleaned up.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_user",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the user. ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `Email for user.`,
				},
				resource.Attribute{
					Name:        "admin",
					Description: `When enabled, this user is an administrator with all the ensuing privileges. Default value is ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "profile_updatable",
					Description: `When set, this user can update his profile details (except for the password. Only an administrator can update the password). Default value is ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "disable_ui_access",
					Description: `When set, this user can only access Artifactory through the REST API. This option cannot be set if the user has Admin privileges. Default value is ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "internal_password_disabled",
					Description: `When set, disables the fallback of using an internal password when external authentication (such as LDAP) is enabled.`,
				},
				resource.Attribute{
					Name:        "groups",
					Description: `List of groups this user is a part of.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "email",
					Description: `Email for user.`,
				},
				resource.Attribute{
					Name:        "admin",
					Description: `When enabled, this user is an administrator with all the ensuing privileges. Default value is ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "profile_updatable",
					Description: `When set, this user can update his profile details (except for the password. Only an administrator can update the password). Default value is ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "disable_ui_access",
					Description: `When set, this user can only access Artifactory through the REST API. This option cannot be set if the user has Admin privileges. Default value is ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "internal_password_disabled",
					Description: `When set, disables the fallback of using an internal password when external authentication (such as LDAP) is enabled.`,
				},
				resource.Attribute{
					Name:        "groups",
					Description: `List of groups this user is a part of.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_virtual_alpine_repository",
			Category:         "Virtual Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"virtual",
				"repositories",
				"alpine",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) the identity key of the repo. ## Attribute Reference The following attributes are supported, along with the [common list of attributes for the remote repositories](../resources/virtual.md):`,
				},
				resource.Attribute{
					Name:        "primary_keypair_ref",
					Description: `(Optional) Primary keypair used to sign artifacts. Default value is empty.`,
				},
				resource.Attribute{
					Name:        "retrieval_cache_period_seconds",
					Description: `(Optional, Default: ` + "`" + `7200` + "`" + `) This value refers to the number of seconds to cache metadata files before checking for newer versions on aggregated repositories. A value of 0 indicates no caching.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "primary_keypair_ref",
					Description: `(Optional) Primary keypair used to sign artifacts. Default value is empty.`,
				},
				resource.Attribute{
					Name:        "retrieval_cache_period_seconds",
					Description: `(Optional, Default: ` + "`" + `7200` + "`" + `) This value refers to the number of seconds to cache metadata files before checking for newer versions on aggregated repositories. A value of 0 indicates no caching.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_virtual_bower_repository",
			Category:         "Virtual Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"virtual",
				"repositories",
				"bower",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) the identity key of the repo. ## Attribute Reference The following attributes are supported, along with the [common list of arguments for the virtual repositories](../resources/virtual.md):`,
				},
				resource.Attribute{
					Name:        "external_dependencies_enabled",
					Description: `(Optional) When set, external dependencies are rewritten. Default value is false.`,
				},
				resource.Attribute{
					Name:        "external_dependencies_remote_repo",
					Description: `(Optional) The remote repository aggregated by this virtual repository in which the external dependency will be cached.`,
				},
				resource.Attribute{
					Name:        "external_dependencies_patterns",
					Description: `(Optional) An Allow List of Ant-style path expressions that specify where external dependencies may be downloaded from. By default, this is set to`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "external_dependencies_enabled",
					Description: `(Optional) When set, external dependencies are rewritten. Default value is false.`,
				},
				resource.Attribute{
					Name:        "external_dependencies_remote_repo",
					Description: `(Optional) The remote repository aggregated by this virtual repository in which the external dependency will be cached.`,
				},
				resource.Attribute{
					Name:        "external_dependencies_patterns",
					Description: `(Optional) An Allow List of Ant-style path expressions that specify where external dependencies may be downloaded from. By default, this is set to`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_virtual_chef_repository",
			Category:         "Virtual Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"virtual",
				"repositories",
				"chef",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) the identity key of the repo. ## Attribute Reference The following attributes are supported, along with the [common list of arguments for the virtual repositories](../resources/virtual.md):`,
				},
				resource.Attribute{
					Name:        "retrieval_cache_period_seconds",
					Description: `(Optional, Default: ` + "`" + `7200` + "`" + `) This value refers to the number of seconds to cache metadata files before checking for newer versions on aggregated repositories. A value of 0 indicates no caching.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "retrieval_cache_period_seconds",
					Description: `(Optional, Default: ` + "`" + `7200` + "`" + `) This value refers to the number of seconds to cache metadata files before checking for newer versions on aggregated repositories. A value of 0 indicates no caching.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_virtual_composer_repository",
			Category:         "Virtual Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"virtual",
				"repositories",
				"composer",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) the identity key of the repo. ## Attribute Reference The [common list of attributes for the virtual repositories](../resources/virtual.md) is supported.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_virtual_conan_repository",
			Category:         "Virtual Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"virtual",
				"repositories",
				"conan",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) the identity key of the repo. ## Attribute Reference The following attributes are supported, along with the [common list of arguments for the virtual repositories](../resources/virtual.md):`,
				},
				resource.Attribute{
					Name:        "retrieval_cache_period_seconds",
					Description: `(Optional, Default: ` + "`" + `7200` + "`" + `) This value refers to the number of seconds to cache metadata files before checking for newer versions on aggregated repositories. A value of 0 indicates no caching.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "retrieval_cache_period_seconds",
					Description: `(Optional, Default: ` + "`" + `7200` + "`" + `) This value refers to the number of seconds to cache metadata files before checking for newer versions on aggregated repositories. A value of 0 indicates no caching.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_virtual_conda_repository",
			Category:         "Virtual Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"virtual",
				"repositories",
				"conda",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) the identity key of the repo. ## Attribute Reference The following attributes are supported, along with the [common list of arguments for the virtual repositories](../resources/virtual.md):`,
				},
				resource.Attribute{
					Name:        "retrieval_cache_period_seconds",
					Description: `(Optional, Default: ` + "`" + `7200` + "`" + `) This value refers to the number of seconds to cache metadata files before checking for newer versions on aggregated repositories. A value of 0 indicates no caching.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "retrieval_cache_period_seconds",
					Description: `(Optional, Default: ` + "`" + `7200` + "`" + `) This value refers to the number of seconds to cache metadata files before checking for newer versions on aggregated repositories. A value of 0 indicates no caching.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_virtual_cran_repository",
			Category:         "Virtual Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"virtual",
				"repositories",
				"cran",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) the identity key of the repo. ## Attribute Reference The following attributes are supported, along with the [common list of arguments for the virtual repositories](../resources/virtual.md):`,
				},
				resource.Attribute{
					Name:        "retrieval_cache_period_seconds",
					Description: `(Optional, Default: ` + "`" + `7200` + "`" + `) This value refers to the number of seconds to cache metadata files before checking for newer versions on aggregated repositories. A value of 0 indicates no caching.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "retrieval_cache_period_seconds",
					Description: `(Optional, Default: ` + "`" + `7200` + "`" + `) This value refers to the number of seconds to cache metadata files before checking for newer versions on aggregated repositories. A value of 0 indicates no caching.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_virtual_debian_repository",
			Category:         "Virtual Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"virtual",
				"repositories",
				"debian",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) the identity key of the repo. ## Attribute Reference The following attributes are supported, along with the [common list of arguments for the virtual repositories](../resources/virtual.md):`,
				},
				resource.Attribute{
					Name:        "retrieval_cache_period_seconds",
					Description: `(Optional, Default: ` + "`" + `7200` + "`" + `) This value refers to the number of seconds to cache metadata files before checking for newer versions on aggregated repositories. A value of 0 indicates no caching.`,
				},
				resource.Attribute{
					Name:        "primary_keypair_ref",
					Description: `(Optional) Primary keypair used to sign artifacts. Default is empty.`,
				},
				resource.Attribute{
					Name:        "secondary_keypair_ref",
					Description: `(Optional) Secondary keypair used to sign artifacts. Default is empty.`,
				},
				resource.Attribute{
					Name:        "optional_index_compression_formats",
					Description: `(Optional) Index file formats you would like to create in addition to the default Gzip (.gzip extension). Supported values are ` + "`" + `bz2` + "`" + `,` + "`" + `lzma` + "`" + ` and ` + "`" + `xz` + "`" + `. Default value is ` + "`" + `bz2` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "debian_default_architectures",
					Description: `(Optional) Specifying architectures will speed up Artifactory's initial metadata indexing process. The default architecture values are amd64 and i386.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "retrieval_cache_period_seconds",
					Description: `(Optional, Default: ` + "`" + `7200` + "`" + `) This value refers to the number of seconds to cache metadata files before checking for newer versions on aggregated repositories. A value of 0 indicates no caching.`,
				},
				resource.Attribute{
					Name:        "primary_keypair_ref",
					Description: `(Optional) Primary keypair used to sign artifacts. Default is empty.`,
				},
				resource.Attribute{
					Name:        "secondary_keypair_ref",
					Description: `(Optional) Secondary keypair used to sign artifacts. Default is empty.`,
				},
				resource.Attribute{
					Name:        "optional_index_compression_formats",
					Description: `(Optional) Index file formats you would like to create in addition to the default Gzip (.gzip extension). Supported values are ` + "`" + `bz2` + "`" + `,` + "`" + `lzma` + "`" + ` and ` + "`" + `xz` + "`" + `. Default value is ` + "`" + `bz2` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "debian_default_architectures",
					Description: `(Optional) Specifying architectures will speed up Artifactory's initial metadata indexing process. The default architecture values are amd64 and i386.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_virtual_docker_repository",
			Category:         "Virtual Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"virtual",
				"repositories",
				"docker",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) the identity key of the repo. ## Attribute Reference The following attributes are supported, along with the [common list of arguments for the virtual repositories](../resources/virtual.md):`,
				},
				resource.Attribute{
					Name:        "resolve_docker_tags_by_timestamp",
					Description: `(Optional) When enabled, in cases where the same Docker tag exists in two or more of the aggregated repositories, Artifactory will return the tag that has the latest timestamp. Default values is ` + "`" + `false` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "resolve_docker_tags_by_timestamp",
					Description: `(Optional) When enabled, in cases where the same Docker tag exists in two or more of the aggregated repositories, Artifactory will return the tag that has the latest timestamp. Default values is ` + "`" + `false` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_virtual_gems_repository",
			Category:         "Virtual Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"virtual",
				"repositories",
				"gems",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) the identity key of the repo. ## Attribute Reference The [common list of attributes for the virtual repositories](../resources/virtual.md) is supported.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_virtual_generic_repository",
			Category:         "Virtual Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"virtual",
				"repositories",
				"generic",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) the identity key of the repo. ## Attribute Reference The [common list of attributes for the virtual repositories](../resources/virtual.md) is supported.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_virtual_gitlfs_repository",
			Category:         "Virtual Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"virtual",
				"repositories",
				"gitlfs",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) the identity key of the repo. ## Attribute Reference The [common list of attributes for the virtual repositories](../resources/virtual.md) is supported.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_virtual_go_repository",
			Category:         "Virtual Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"virtual",
				"repositories",
				"go",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) the identity key of the repo. ## Attribute Reference The following attributes are supported, along with the [common list of arguments for the virtual repositories](../resources/virtual.md):`,
				},
				resource.Attribute{
					Name:        "external_dependencies_enabled",
					Description: `(Optional) Shorthand for "Enable 'go-import' Meta Tags" on the UI. This must be set to true in order to use the allow list. When checked (default), Artifactory will automatically follow remote VCS roots in 'go-import' meta tags to download remote modules.`,
				},
				resource.Attribute{
					Name:        "external_dependencies_patterns",
					Description: `(Optional) 'go-import' Allow List on the UI.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "external_dependencies_enabled",
					Description: `(Optional) Shorthand for "Enable 'go-import' Meta Tags" on the UI. This must be set to true in order to use the allow list. When checked (default), Artifactory will automatically follow remote VCS roots in 'go-import' meta tags to download remote modules.`,
				},
				resource.Attribute{
					Name:        "external_dependencies_patterns",
					Description: `(Optional) 'go-import' Allow List on the UI.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_virtual_gradle_repository",
			Category:         "Virtual Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"virtual",
				"repositories",
				"gradle",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) the identity key of the repo. ## Attribute Reference The following attributes are supported, along with the [common list of arguments for the virtual repositories](../resources/virtual.md):`,
				},
				resource.Attribute{
					Name:        "pom_repository_references_cleanup_policy",
					Description: `(Optional) - (1: discard_active_reference) Discard Active References - Removes repository elements that are declared directly under project or under a profile in the same POM that is activeByDefault. - (2: discard_any_reference) Discard Any References - Removes all repository elements regardless of whether they are included in an active profile or not. - (3: nothing) Nothing - Does not remove any repository elements declared in the POM.`,
				},
				resource.Attribute{
					Name:        "key_pair",
					Description: `(Optional) The keypair used to sign artifacts.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "pom_repository_references_cleanup_policy",
					Description: `(Optional) - (1: discard_active_reference) Discard Active References - Removes repository elements that are declared directly under project or under a profile in the same POM that is activeByDefault. - (2: discard_any_reference) Discard Any References - Removes all repository elements regardless of whether they are included in an active profile or not. - (3: nothing) Nothing - Does not remove any repository elements declared in the POM.`,
				},
				resource.Attribute{
					Name:        "key_pair",
					Description: `(Optional) The keypair used to sign artifacts.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_virtual_helm_repository",
			Category:         "Virtual Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"virtual",
				"repositories",
				"helm",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) the identity key of the repo. ## Attribute Reference The following attributes are supported, along with the [common list of arguments for the virtual repositories](../resources/virtual.md):`,
				},
				resource.Attribute{
					Name:        "retrieval_cache_period_seconds",
					Description: `(Optional, Default: ` + "`" + `7200` + "`" + `) This value refers to the number of seconds to cache metadata files before checking for newer versions on aggregated repositories. A value of 0 indicates no caching.`,
				},
				resource.Attribute{
					Name:        "use_namespaces",
					Description: `(Optional) From Artifactory 7.24.1 (SaaS Version), you can explicitly state a specific aggregated local or remote repository to fetch from a virtual by assigning namespaces to local and remote repositories. See the documentation [here](https://www.jfrog.com/confluence/display/JFROG/Kubernetes+Helm+Chart+Repositories#KubernetesHelmChartRepositories-NamespaceSupportforHelmVirtualRepositories). Default is ` + "`" + `false` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "retrieval_cache_period_seconds",
					Description: `(Optional, Default: ` + "`" + `7200` + "`" + `) This value refers to the number of seconds to cache metadata files before checking for newer versions on aggregated repositories. A value of 0 indicates no caching.`,
				},
				resource.Attribute{
					Name:        "use_namespaces",
					Description: `(Optional) From Artifactory 7.24.1 (SaaS Version), you can explicitly state a specific aggregated local or remote repository to fetch from a virtual by assigning namespaces to local and remote repositories. See the documentation [here](https://www.jfrog.com/confluence/display/JFROG/Kubernetes+Helm+Chart+Repositories#KubernetesHelmChartRepositories-NamespaceSupportforHelmVirtualRepositories). Default is ` + "`" + `false` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_virtual_ivy_repository",
			Category:         "Virtual Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"virtual",
				"repositories",
				"ivy",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) the identity key of the repo. ## Attribute Reference The following attributes are supported, along with the [common list of arguments for the virtual repositories](../resources/virtual.md):`,
				},
				resource.Attribute{
					Name:        "pom_repository_references_cleanup_policy",
					Description: `(Optional) - (1: discard_active_reference) Discard Active References - Removes repository elements that are declared directly under project or under a profile in the same POM that is activeByDefault. - (2: discard_any_reference) Discard Any References - Removes all repository elements regardless of whether they are included in an active profile or not. - (3: nothing) Nothing - Does not remove any repository elements declared in the POM.`,
				},
				resource.Attribute{
					Name:        "key_pair",
					Description: `(Optional) The keypair used to sign artifacts.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "pom_repository_references_cleanup_policy",
					Description: `(Optional) - (1: discard_active_reference) Discard Active References - Removes repository elements that are declared directly under project or under a profile in the same POM that is activeByDefault. - (2: discard_any_reference) Discard Any References - Removes all repository elements regardless of whether they are included in an active profile or not. - (3: nothing) Nothing - Does not remove any repository elements declared in the POM.`,
				},
				resource.Attribute{
					Name:        "key_pair",
					Description: `(Optional) The keypair used to sign artifacts.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_virtual_maven_repository",
			Category:         "Virtual Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"virtual",
				"repositories",
				"maven",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) the identity key of the repo. ## Attribute Reference The following attributes are supported, along with the [common list of arguments for the virtual repositories](../resources/virtual.md):`,
				},
				resource.Attribute{
					Name:        "pom_repository_references_cleanup_policy",
					Description: `(Optional) One of: ` + "`" + `"discard_active_reference", "discard_any_reference", "nothing"` + "`" + ``,
				},
				resource.Attribute{
					Name:        "force_maven_authentication",
					Description: `(Optional) Forces authentication when fetching from remote repos.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "pom_repository_references_cleanup_policy",
					Description: `(Optional) One of: ` + "`" + `"discard_active_reference", "discard_any_reference", "nothing"` + "`" + ``,
				},
				resource.Attribute{
					Name:        "force_maven_authentication",
					Description: `(Optional) Forces authentication when fetching from remote repos.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_virtual_npm_repository",
			Category:         "Virtual Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"virtual",
				"repositories",
				"npm",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) the identity key of the repo. ## Attribute Reference The following attributes are supported, along with the [common list of arguments for the virtual repositories](../resources/virtual.md):`,
				},
				resource.Attribute{
					Name:        "retrieval_cache_period_seconds",
					Description: `(Optional, Default: ` + "`" + `7200` + "`" + `) This value refers to the number of seconds to cache metadata files before checking for newer versions on aggregated repositories. A value of 0 indicates no caching.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "retrieval_cache_period_seconds",
					Description: `(Optional, Default: ` + "`" + `7200` + "`" + `) This value refers to the number of seconds to cache metadata files before checking for newer versions on aggregated repositories. A value of 0 indicates no caching.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_virtual_nuget_repository",
			Category:         "Virtual Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"virtual",
				"repositories",
				"nuget",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) the identity key of the repo. ## Attribute Reference The following attributes are supported, along with the [common list of arguments for the virtual repositories](../resources/virtual.md):`,
				},
				resource.Attribute{
					Name:        "force_nuget_authentication",
					Description: `(Optional) If set, user authentication is required when accessing the repository. An anonymous request will display an HTTP 401 error. This is also enforced when aggregated repositories support anonymous requests. Default is ` + "`" + `false` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "force_nuget_authentication",
					Description: `(Optional) If set, user authentication is required when accessing the repository. An anonymous request will display an HTTP 401 error. This is also enforced when aggregated repositories support anonymous requests. Default is ` + "`" + `false` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_virtual_p2_repository",
			Category:         "Virtual Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"virtual",
				"repositories",
				"p2",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) the identity key of the repo. ## Attribute Reference The [common list of attributes for the virtual repositories](../resources/virtual.md) is supported.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_virtual_pub_repository",
			Category:         "Virtual Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"virtual",
				"repositories",
				"pub",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) the identity key of the repo. ## Attribute Reference The [common list of attributes for the virtual repositories](../resources/virtual.md) is supported.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_virtual_puppet_repository",
			Category:         "Virtual Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"virtual",
				"repositories",
				"puppet",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) the identity key of the repo. ## Attribute Reference The [common list of attributes for the virtual repositories](../resources/virtual.md) is supported.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_virtual_pypi_repository",
			Category:         "Virtual Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"virtual",
				"repositories",
				"pypi",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) the identity key of the repo. ## Attribute Reference The [common list of attributes for the virtual repositories](../resources/virtual.md) is supported.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_virtual_rpm_repository",
			Category:         "Virtual Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"virtual",
				"repositories",
				"rpm",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) the identity key of the repo. ## Attribute Reference The following attributes are supported, along with the [common list of arguments for the virtual repositories](../resources/virtual.md):`,
				},
				resource.Attribute{
					Name:        "primary_keypair_ref",
					Description: `(Optional) The primary GPG key to be used to sign packages.`,
				},
				resource.Attribute{
					Name:        "secondary_keypair_ref",
					Description: `(Optional) The secondary GPG key to be used to sign packages.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "primary_keypair_ref",
					Description: `(Optional) The primary GPG key to be used to sign packages.`,
				},
				resource.Attribute{
					Name:        "secondary_keypair_ref",
					Description: `(Optional) The secondary GPG key to be used to sign packages.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_virtual_sbt_repository",
			Category:         "Virtual Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"virtual",
				"repositories",
				"sbt",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) the identity key of the repo. ## Attribute Reference The following attributes are supported, along with the [common list of arguments for the virtual repositories](../resources/virtual.md):`,
				},
				resource.Attribute{
					Name:        "pom_repository_references_cleanup_policy",
					Description: `(Optional) - (1: discard_active_reference) Discard Active References - Removes repository elements that are declared directly under project or under a profile in the same POM that is activeByDefault. - (2: discard_any_reference) Discard Any References - Removes all repository elements regardless of whether they are included in an active profile or not. - (3: nothing) Nothing - Does not remove any repository elements declared in the POM.`,
				},
				resource.Attribute{
					Name:        "key_pair",
					Description: `(Optional) The keypair used to sign artifacts.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "pom_repository_references_cleanup_policy",
					Description: `(Optional) - (1: discard_active_reference) Discard Active References - Removes repository elements that are declared directly under project or under a profile in the same POM that is activeByDefault. - (2: discard_any_reference) Discard Any References - Removes all repository elements regardless of whether they are included in an active profile or not. - (3: nothing) Nothing - Does not remove any repository elements declared in the POM.`,
				},
				resource.Attribute{
					Name:        "key_pair",
					Description: `(Optional) The keypair used to sign artifacts.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_virtual_swift_repository",
			Category:         "Virtual Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"virtual",
				"repositories",
				"swift",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) the identity key of the repo. ## Attribute Reference The [common list of attributes for the virtual repositories](../resources/virtual.md) is supported.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "artifactory_virtual_terraform_repository",
			Category:         "Virtual Repositories",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"virtual",
				"repositories",
				"terraform",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) the identity key of the repo. ## Attribute Reference The [common list of attributes for the virtual repositories](../resources/virtual.md) is supported.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
	}

	dataSourcesMap = map[string]int{

		"artifactory_artifactory_file":                        0,
		"artifactory_artifactory_fileinfo":                    1,
		"artifactory_federated_alpine_repository":             2,
		"artifactory_federated_bower_repository":              3,
		"artifactory_federated_cargo_repository":              4,
		"artifactory_federated_chef_repository":               5,
		"artifactory_federated_cocoapods_repository":          6,
		"artifactory_federated_composer_repository":           7,
		"artifactory_federated_conan_repository":              8,
		"artifactory_federated_conda_repository":              9,
		"artifactory_federated_cran_repository":               10,
		"artifactory_federated_debian_repository":             11,
		"artifactory_federated_docker_v1_repository":          12,
		"artifactory_federated_docker_v2_repository":          13,
		"artifactory_federated_gems_repository":               14,
		"artifactory_federated_generic_repository":            15,
		"artifactory_federated_gitlfs_repository":             16,
		"artifactory_federated_go_repository":                 17,
		"artifactory_federated_gradle_repository":             18,
		"artifactory_federated_helm_repository":               19,
		"artifactory_federated_ivy_repository":                20,
		"artifactory_federated_maven_repository":              21,
		"artifactory_federated_npm_repository":                22,
		"artifactory_federated_nuget_repository":              23,
		"artifactory_federated_opkg_repository":               24,
		"artifactory_federated_puppet_repository":             25,
		"artifactory_federated_pypi_repository":               26,
		"artifactory_federated_rpm_repository":                27,
		"artifactory_federated_sbt_repository":                28,
		"artifactory_federated_swift_repository":              29,
		"artifactory_federated_terraform_module_repository":   30,
		"artifactory_federated_terraform_provider_repository": 31,
		"artifactory_federated_vagrant_repository":            32,
		"artifactory_group":                                   33,
		"artifactory_local":                                   34,
		"artifactory_local_alpine_repository":                 35,
		"artifactory_local_bower_repository":                  36,
		"artifactory_local_cargo_repository":                  37,
		"artifactory_local_chef_repository":                   38,
		"artifactory_local_cocoapods_repository":              39,
		"artifactory_local_composer_repository":               40,
		"artifactory_local_conan_repository":                  41,
		"artifactory_local_conda_repository":                  42,
		"artifactory_local_cran_repository":                   43,
		"artifactory_local_debian_repository":                 44,
		"artifactory_local_docker_v1_repository":              45,
		"artifactory_local_docker_v2_repository":              46,
		"artifactory_local_gems_repository":                   47,
		"artifactory_local_generic_repository":                48,
		"artifactory_local_gitlfs_repository":                 49,
		"artifactory_local_go_repository":                     50,
		"artifactory_local_gradle_repository":                 51,
		"artifactory_local_helm_repository":                   52,
		"artifactory_local_ivy_repository":                    53,
		"artifactory_local_maven_repository":                  54,
		"artifactory_local_npm_repository":                    55,
		"artifactory_local_nuget_repository":                  56,
		"artifactory_local_opkg_repository":                   57,
		"artifactory_local_pub_repository":                    58,
		"artifactory_local_puppet_repository":                 59,
		"artifactory_local_pypi_repository":                   60,
		"artifactory_local_rpm_repository":                    61,
		"artifactory_local_sbt_repository":                    62,
		"artifactory_local_swift_repository":                  63,
		"artifactory_local_terraform_module_repository":       64,
		"artifactory_local_terraform_provider_repository":     65,
		"artifactory_local_terraformbackend_repository":       66,
		"artifactory_local_vagrant_repository":                67,
		"artifactory_permission_target":                       68,
		"artifactory_remote_alpine_repository":                69,
		"artifactory_remote_bower_repository":                 70,
		"artifactory_remote_cargo_repository":                 71,
		"artifactory_remote_chef_repository":                  72,
		"artifactory_remote_cocoapods_repository":             73,
		"artifactory_remote_composer_repository":              74,
		"artifactory_remote_conan_repository":                 75,
		"artifactory_remote_conda_repository":                 76,
		"artifactory_remote_cran_repository":                  77,
		"artifactory_remote_debian_repository":                78,
		"artifactory_remote_docker_repository":                79,
		"artifactory_remote_gems_repository":                  80,
		"artifactory_remote_generic_repository":               81,
		"artifactory_remote_gitlfs_repository":                82,
		"artifactory_remote_go_repository":                    83,
		"artifactory_remote_gradle_repository":                84,
		"artifactory_remote_helm_repository":                  85,
		"artifactory_remote_ivy_repository":                   86,
		"artifactory_remote_maven_repository":                 87,
		"artifactory_remote_npm_repository":                   88,
		"artifactory_remote_nuget_repository":                 89,
		"artifactory_remote_opkg_repository":                  90,
		"artifactory_remote_p2_repository":                    91,
		"artifactory_remote_pub_repository":                   92,
		"artifactory_remote_puppet_repository":                93,
		"artifactory_remote_pypi_repository":                  94,
		"artifactory_remote_rpm_repository":                   95,
		"artifactory_remote_sbt_repository":                   96,
		"artifactory_remote_swift_repository":                 97,
		"artifactory_remote_terraform_repository":             98,
		"artifactory_remote_vcs_repository":                   99,
		"artifactory_user":                                    100,
		"artifactory_virtual_alpine_repository":               101,
		"artifactory_virtual_bower_repository":                102,
		"artifactory_virtual_chef_repository":                 103,
		"artifactory_virtual_composer_repository":             104,
		"artifactory_virtual_conan_repository":                105,
		"artifactory_virtual_conda_repository":                106,
		"artifactory_virtual_cran_repository":                 107,
		"artifactory_virtual_debian_repository":               108,
		"artifactory_virtual_docker_repository":               109,
		"artifactory_virtual_gems_repository":                 110,
		"artifactory_virtual_generic_repository":              111,
		"artifactory_virtual_gitlfs_repository":               112,
		"artifactory_virtual_go_repository":                   113,
		"artifactory_virtual_gradle_repository":               114,
		"artifactory_virtual_helm_repository":                 115,
		"artifactory_virtual_ivy_repository":                  116,
		"artifactory_virtual_maven_repository":                117,
		"artifactory_virtual_npm_repository":                  118,
		"artifactory_virtual_nuget_repository":                119,
		"artifactory_virtual_p2_repository":                   120,
		"artifactory_virtual_pub_repository":                  121,
		"artifactory_virtual_puppet_repository":               122,
		"artifactory_virtual_pypi_repository":                 123,
		"artifactory_virtual_rpm_repository":                  124,
		"artifactory_virtual_sbt_repository":                  125,
		"artifactory_virtual_swift_repository":                126,
		"artifactory_virtual_terraform_repository":            127,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
