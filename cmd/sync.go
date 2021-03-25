package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/go-git/go-git/v5"
)

// sync will pull all the Providers from TF Registry and create
// the corresponding directories and also pull them to update them
func sync() error {
	providersURL := "https://registry.terraform.io/v2/providers?filter[tier]=official,partner&page[number]=%d&page[size]=100"
	page := 1
	var finish bool
	exportedProviders := map[string]Provider{
		"flexibleengine": Provider{
			Attributes: ProviderAttibutes{
				Source: "https://github.com/FlexibleEngineCloud/terraform-provider-flexibleengine",
			},
		},
		"openstack": Provider{
			Attributes: ProviderAttibutes{
				Source: "https://github.com/terraform-provider-openstack/terraform-provider-openstack",
			},
		},
	}

	for !finish {
		url := fmt.Sprintf(providersURL, page)
		res, err := http.Get(url)
		if err != nil {
			return fmt.Errorf("could not http.Get: %w", err)
		}
		defer res.Body.Close()

		if res.StatusCode != http.StatusOK {
			return fmt.Errorf("Wrong HTTP status from Providers %d with url %s", res.StatusCode, url)
		}

		var pvs ProvidersResponse
		err = json.NewDecoder(res.Body).Decode(&pvs)
		if res.StatusCode != http.StatusOK {
			return fmt.Errorf("Failed to decode body %w", err)
		}

		for _, p := range pvs.Data {
			exportedProviders[p.Attributes.Name] = p
		}

		// 0 means it has no value
		if pvs.Meta.Pagination.NextPage != 0 {
			page = pvs.Meta.Pagination.NextPage
		} else {
			finish = true
		}
	}

	deletedProviders := make(map[string]struct{})
	addedProviders := make(map[string]struct{})
	currentProviders := make(map[string]struct{})

	// New we have all the Providers we want to document
	// on providers
	de, err := os.ReadDir(repositoriesPath)
	if err != nil {
		return fmt.Errorf("could not read %q %w", repositoriesPath, err)
	}

	for _, d := range de {
		if d.IsDir() {
			// If a local ones is not on the exported list it means
			// it has been deleted
			if _, ok := exportedProviders[d.Name()]; !ok {
				deletedProviders[d.Name()] = struct{}{}
				continue
			}

			currentProviders[d.Name()] = struct{}{}
		}
	}

	for k := range exportedProviders {
		if _, ok := currentProviders[k]; !ok {
			addedProviders[k] = struct{}{}
		}
	}

	// We then delete the Providers that need it
	deletedCount := 0
	for k := range deletedProviders {
		deletedCount++
		fmt.Println("Deleting [%d/%d] %s\n", deletedCount, len(deletedProviders), k)
		os.RemoveAll(filepath.Join(repositoriesPath, k))
		os.RemoveAll(filepath.Join(providersPath, k))
	}

	// Now we first clone the new ones
	// which are on addedProviders
	addedCount := 0
	for k := range addedProviders {
		addedCount++
		p := exportedProviders[k]
		fmt.Printf("Cloning [%d/%d] %s\n", addedCount, len(addedProviders), k)
		_, err = git.PlainClone(filepath.Join(repositoriesPath, k), false, &git.CloneOptions{
			URL:               p.Attributes.Source,
			RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
		})
		if err != nil {
			return fmt.Errorf("failed to clone %q %q: %w", k, p.Attributes.Source, err)
		}
	}

	// Then we pull the existing ones
	// which are on currentProviders
	currentCount := 0
	for k := range currentProviders {
		currentCount++

		fmt.Printf("Pulling [%d/%d] %s\n", currentCount, len(currentProviders), k)

		path := filepath.Join(repositoriesPath, k)
		r, err := git.PlainOpen(path)
		if err != nil {
			return fmt.Errorf("failed to open %q %q: %w", k, path, err)
		}

		w, err := r.Worktree()
		if err != nil {
			return fmt.Errorf("failed to operate worktree %q %q: %w", k, path, err)
		}

		err = w.Pull(&git.PullOptions{})
		if err != nil && !errors.Is(err, git.NoErrAlreadyUpToDate) {
			fmt.Printf("%+v\n", err)
			return fmt.Errorf("failed to pull %q %q: %w", k, path, err)
		}
	}

	return nil
}

type ProvidersResponse struct {
	Data  []Provider `json:"data"`
	Links struct {
		First string      `json:"first"`
		Last  string      `json:"last"`
		Next  string      `json:"next"`
		Prev  interface{} `json:"prev"`
	} `json:"links"`
	Meta struct {
		Pagination struct {
			PageSize    int `json:"page-size"`
			CurrentPage int `json:"current-page"`
			NextPage    int `json:"next-page"`
			PrevPage    int `json:"prev-page"`
			TotalPages  int `json:"total-pages"`
			TotalCount  int `json:"total-count"`
		} `json:"pagination"`
	} `json:"meta"`
}

type Provider struct {
	Type       string            `json:"type"`
	ID         string            `json:"id"`
	Attributes ProviderAttibutes `json:"attributes"`
	Links      struct {
		Self string `json:"self"`
	} `json:"links"`
}

type ProviderAttibutes struct {
	Alias         string `json:"alias"`
	Description   string `json:"description"`
	Downloads     int    `json:"downloads"`
	Featured      bool   `json:"featured"`
	FullName      string `json:"full-name"`
	Name          string `json:"name"`
	Namespace     string `json:"namespace"`
	OwnerName     string `json:"owner-name"`
	RobotsNoindex bool   `json:"robots-noindex"`
	Source        string `json:"source"`
	Tier          string `json:"tier"`
	Unlisted      bool   `json:"unlisted"`
	Warning       string `json:"warning"`
}
