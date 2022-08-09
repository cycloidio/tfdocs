package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go/format"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/cycloidio/tfdocs/resource"
)

const (
	iconFile  = "icons.json"
	aliasFile = "aliases.json"
)

var (
	erbEmbed         = regexp.MustCompile("<%=?[^>]*>")
	repositoriesPath = "repositories"
	providersPath    = "providers"
	assetsPath       = "assets"
	aliasesPath      = "aliases"
)

var reProvider = regexp.MustCompile("^[a-z]+$")

func generate() error {
	de, err := ioutil.ReadDir(repositoriesPath)
	if err != nil {
		return fmt.Errorf("could not read %q %w", repositoriesPath, err)
	}

	oldp := make([]string, 0, 0)
	newp := make([]string, 0, 0)
	unkp := make([]string, 0, 0)
	newformat := 0
	oldformat := 0
	for _, d := range de {
		if !d.IsDir() {
			continue
		}
		provider := d.Name()
		// Providers that do not follow the
		// rule of having the name only have
		// letters will e ignored for now
		if !reProvider.MatchString(provider) {
			continue
		}

		if _, err := os.Stat(filepath.Join(repositoriesPath, provider, "docs", "index.md")); err == nil {
			newGenerate(provider)
			newp = append(newp, provider)
			newformat++
		} else if _, err := os.Stat(filepath.Join(repositoriesPath, provider, "website", fmt.Sprintf("%s.erb", provider))); err == nil {
			oldGenerate(provider)
			oldp = append(oldp, provider)
			oldformat++
		} else {
			unkGenerate(provider)
			unkp = append(unkp, provider)
		}
	}
	total := len(de)
	fmt.Printf("New: %d (%.2f%%)\n", newformat, (float64(newformat)/float64(total))*100)
	fmt.Printf("Old: %d (%.2f%%)\n", oldformat, (float64(oldformat)/float64(total))*100)
	fmt.Printf("Unknown: %d (%.2f%%)\n", total-(oldformat+newformat), (float64(total-(oldformat+newformat))/float64(total))*100)
	fmt.Printf("Total: %d\n", total)

	fmt.Println("New")
	fmt.Println(newp)
	fmt.Println("--")
	fmt.Println("Old")
	fmt.Println(oldp)
	fmt.Println("--")
	fmt.Println("Unk")
	fmt.Println(unkp)

	return nil
}

func readJSONFile(path string, obj interface{}) error {
	b, err := ioutil.ReadFile(path)
	if !os.IsNotExist(err) {
		err = json.Unmarshal(b, &obj)
		if err != nil {
			return err
		}
	}
	return nil
}

func unkGenerate(provider string) {
	var icons map[string]string
	err := readJSONFile(filepath.Join(assetsPath, provider, iconFile), &icons)
	if err != nil {
		panic(err)
	}

	var aliases = make(map[string]map[string][]string)
	err = readJSONFile(filepath.Join(aliasesPath, provider, aliasFile), &aliases)
	if err != nil {
		panic(err)
	}

	docsPath := filepath.Join(repositoriesPath, provider, "website", "docs")
	for _, t := range []string{"r", "d"} {
		files, err := os.ReadDir(filepath.Join(docsPath, t))
		// If it does not have the type we ignore it
		if os.IsNotExist(err) {
			continue
		}

		// Create the Provider folder
		err = os.MkdirAll(filepath.Join(providersPath, provider), os.ModePerm)
		if err != nil && !os.IsExist(err) {
			panic(err)
		}

		// Create the Output file
		out, err := os.OpenFile(filepath.Join(providersPath, provider, fmt.Sprintf("%s.go", t)), os.O_APPEND|os.O_TRUNC|os.O_RDWR|os.O_CREATE, 0644)
		if err != nil {
			panic(err)
		}
		defer out.Close()

		resources := make([]resource.Resource, 0)
		for _, f := range files {
			if f.IsDir() || (path.Ext(f.Name()) != ".markdown" && path.Ext(f.Name()) != ".md") {
				continue
			}
			body, err := ioutil.ReadFile(filepath.Join(docsPath, t, f.Name()))
			if err != nil {
				panic(err)
			}

			rt := strings.Split(f.Name(), ".")[0]
			// The filename is the resource without the Provider
			if !strings.HasPrefix(f.Name(), provider) {
				rt = fmt.Sprintf("%s_%s", provider, rt)
			}

			r := resource.Resource{
				Type:             rt,
				ShortDescription: getShortDescription(body),
				Description:      getDescription(body),
				Arguments:        getArguments(body),
				Attributes:       getAttributes(body),
			}
			if t == "d" {
				r.Category = "Data Sources"
			} else {
				r.Category = "Resources"
			}
			cat := getCategory(body)
			if cat != "" && t != "d" {
				r.Category = cat
				r.Keywords = categoryAndTypeToKeywords(provider, cat, rt)
			}
			if icons != nil {
				if ic, ok := icons[rt]; ok {
					r.Icon = ic
				} else {
					icons[rt] = ""
				}
			}
			resources = append(resources, r)

			if nt, ok := aliases[t][rt]; ok {
				for _, t := range nt {
					r.Type = t
					resources = append(resources, r)
				}
			}
		}

		buff := &bytes.Buffer{}

		td := TemplateData{
			Resources: resources,
			Type:      t,
			Provider:  provider,
		}

		err = resourceTmpl.Execute(buff, td)
		if err != nil {
			panic(err)
		}

		b, err := format.Source(buff.Bytes())
		if err != nil {
			buff = &bytes.Buffer{}
			log.Printf("ERROR: could not import for provider %q and type %q setting empty, error: %s", provider, t, err)
			td.Resources = nil
			err = resourceTmpl.Execute(buff, td)
			if err != nil {
				panic(err)
			}
			b, err = format.Source(buff.Bytes())
			if err != nil {
				fmt.Println(provider)
				panic(err)
			}
		}

		io.Copy(out, bytes.NewBuffer(b))
	}

	// Update the list of icons with the potential new resources
	if icons != nil {
		f, err := os.OpenFile(filepath.Join(assetsPath, provider, iconFile), os.O_RDWR|os.O_TRUNC, 0755)
		if err != nil {
			log.Fatal(err)
		}
		e := json.NewEncoder(f)
		e.SetEscapeHTML(false)
		e.SetIndent("", " ")
		err = e.Encode(icons)
		if err != nil {
			log.Fatal(err)
		}
		f.Close()
	}
}

func newGenerate(provider string) {
	var icons map[string]string
	err := readJSONFile(filepath.Join(assetsPath, provider, iconFile), &icons)
	if err != nil {
		panic(err)
	}

	var aliases = make(map[string]map[string][]string)
	err = readJSONFile(filepath.Join(aliasesPath, provider, aliasFile), &aliases)
	if err != nil {
		panic(err)
	}

	docsPath := filepath.Join(repositoriesPath, provider, "docs")
	for _, t := range []string{"data-sources", "resources"} {
		files, err := os.ReadDir(filepath.Join(docsPath, t))
		// If it does not have the type we ignore it
		if os.IsNotExist(err) {
			continue
		}

		// Create the Provider folder
		err = os.MkdirAll(filepath.Join(providersPath, provider), os.ModePerm)
		if err != nil && !os.IsExist(err) {
			panic(err)
		}

		// Create the Output file
		out, err := os.OpenFile(filepath.Join(providersPath, provider, fmt.Sprintf("%s.go", t)), os.O_APPEND|os.O_TRUNC|os.O_RDWR|os.O_CREATE, 0644)
		if err != nil {
			panic(err)
		}
		defer out.Close()

		resources := make([]resource.Resource, 0)
		for _, f := range files {
			if f.IsDir() || path.Ext(f.Name()) != ".md" {
				continue
			}
			body, err := ioutil.ReadFile(filepath.Join(docsPath, t, f.Name()))
			if err != nil {
				panic(err)
			}

			// The filename is the resource without the Provider
			rt := fmt.Sprintf("%s_%s", provider, strings.Split(f.Name(), ".")[0])

			r := resource.Resource{
				Type:             rt,
				ShortDescription: getShortDescription(body),
				Description:      getDescription(body),
				Arguments:        getArguments(body),
				Attributes:       getAttributes(body),
			}
			if t == "data-sources" {
				r.Category = "Data Sources"
			} else {
				r.Category = "Resources"
			}
			cat := getCategory(body)
			if cat != "" {
				r.Category = cat
				r.Keywords = categoryAndTypeToKeywords(provider, cat, rt)
			}
			if icons != nil {
				if ic, ok := icons[rt]; ok {
					r.Icon = ic
				} else {
					icons[rt] = ""
				}
			}
			resources = append(resources, r)

			if nt, ok := aliases[t][rt]; ok {
				for _, t := range nt {
					r.Type = t
					resources = append(resources, r)
				}
			}
		}

		buff := &bytes.Buffer{}

		td := TemplateData{
			Resources: resources,
			Type:      t,
			Provider:  provider,
		}

		err = resourceTmpl.Execute(buff, td)
		if err != nil {
			panic(err)
		}

		b, err := format.Source(buff.Bytes())
		if err != nil {
			buff = &bytes.Buffer{}
			log.Printf("ERROR: could not import for provider %q and type %q setting empty", provider, t)
			td.Resources = nil
			err = resourceTmpl.Execute(buff, td)
			if err != nil {
				panic(err)
			}
			b, err = format.Source(buff.Bytes())
			if err != nil {
				fmt.Println(provider)
				panic(err)
			}
		}

		io.Copy(out, bytes.NewBuffer(b))
	}

	// Update the list of icons with the potential new resources
	if icons != nil {
		f, err := os.OpenFile(filepath.Join(assetsPath, provider, iconFile), os.O_RDWR|os.O_TRUNC, 0755)
		if err != nil {
			log.Fatal(err)
		}
		e := json.NewEncoder(f)
		e.SetEscapeHTML(false)
		e.SetIndent("", " ")
		err = e.Encode(icons)
		if err != nil {
			log.Fatal(err)
		}
		f.Close()
	}

}

func oldGenerate(provider string) {
	// avi has all the Resource names as 'ActionGroupConfig' instead of 'avi_actiongroupconfig'
	// as any other provider would have
	if provider == "avi" {
		return
	}
	var icons map[string]string
	err := readJSONFile(filepath.Join(assetsPath, provider, iconFile), &icons)
	if err != nil {
		panic(err)
	}

	var aliases = make(map[string]map[string][]string)
	err = readJSONFile(filepath.Join(aliasesPath, provider, aliasFile), &aliases)
	if err != nil {
		panic(err)
	}

	docsPath := filepath.Join(repositoriesPath, provider, "website", "docs")
	sidebarPath := filepath.Join(repositoriesPath, provider, "website", fmt.Sprintf("%s.erb", provider))
	secondSidebarPath := filepath.Join(repositoriesPath, provider, "website", "layout.erb")

	// categories has as key the resource name and as value the category
	// they belong to
	// The DataSources are not here
	categories := make(map[string]string)

	// The type it's hard to guess.
	// * On the .markdown, the `Resource: ` sometimes has with or without aws_, and some
	// times it does not match
	// * The `page_title` the smae
	// * File name does not match always
	// The only valid solution is to read the sidebar and get the URL which is the file
	// and use the Text to set as type
	fileType := make(map[string]string)

	body, err := ioutil.ReadFile(sidebarPath)
	if err != nil {
		if strings.Contains(err.Error(), "no such file or directory") {
			body, err = ioutil.ReadFile(secondSidebarPath)
			if err != nil {
				panic(err)
			}
		} else {
			panic(err)
		}
	}

	body = erbEmbed.ReplaceAll(body, nil)

	doc, err := goquery.NewDocumentFromReader(bytes.NewBuffer(body))
	if err != nil {
		panic(err)
	}

	doc.Find(".nav.docs-sidenav > li").Each(func(_ int, s *goquery.Selection) {
		as := s.Find("a")
		cat := as.First().Text()

		// If they have more than one 'ul' it means that it has
		// the Category as first lvl and then the actual differenciation
		// between Resource or DataSource in 2nd lvl
		ulsLen := s.Find("ul").Length()
		if ulsLen > 1 {
			s.ChildrenFiltered("ul").ChildrenFiltered("li").Each(func(_ int, s *goquery.Selection) {
				as = s.Find("a")
				resourceType := as.First().Text()

				isResource := true
				if strings.Contains(resourceType, "Data Sources") || strings.Contains(resourceType, "Data Resources") {
					isResource = false
				}
				as.Slice(1, goquery.ToEnd).Each(func(i int, s *goquery.Selection) {
					parseATag(s, provider, cat, isResource, fileType, categories)
				})
			})
		} else if ulsLen == 1 {
			as.Slice(1, goquery.ToEnd).Each(func(i int, s *goquery.Selection) {
				isResource := true
				if strings.Contains(cat, "Data Sources") || strings.Contains(cat, "Data Resources") {
					isResource = false
				}
				parseATag(s, provider, cat, isResource, fileType, categories)
			})
		}
	})

	for _, t := range []string{"r", "d"} {
		files, err := ioutil.ReadDir(filepath.Join(docsPath, t))
		if err != nil {
			// This means that the provides does not have 'r' or 'd'
			if strings.Contains(err.Error(), "no such file or directory") {
				continue
			}
			panic(err)
		}

		err = os.MkdirAll(filepath.Join(providersPath, provider), os.ModePerm)
		if err != nil && !os.IsExist(err) {
			panic(err)
		}
		out, err := os.OpenFile(filepath.Join(providersPath, provider, fmt.Sprintf("%s.go", t)), os.O_APPEND|os.O_TRUNC|os.O_RDWR|os.O_CREATE, 0644)
		if err != nil {
			panic(err)
		}
		defer out.Close()

		resources := make([]resource.Resource, 0)
		for _, f := range files {
			if f.IsDir() || f.Name() == ".keep" {
				continue
			}
			body, err := ioutil.ReadFile(filepath.Join(docsPath, t, f.Name()))
			if err != nil {
				panic(err)
			}

			fileName := fmt.Sprintf("%s_%s", t, strings.Split(f.Name(), ".")[0])
			rt, ok := fileType[fileName]
			if !ok {
				// This means that the sidebar has no link to this specific file
				// which means it cannot be seen on the docs
				log.Printf("the provider %q file %q of %q has no type, name used was %q", provider, f.Name(), t, fileName)
				continue
			}

			r := resource.Resource{
				Type:             rt,
				ShortDescription: getShortDescription(body),
				Description:      getDescription(body),
				Arguments:        getArguments(body),
				Attributes:       getAttributes(body),
			}
			if t == "d" {
				r.Category = "Data Sources"
			} else {
				c, ok := categories[rt]
				if !ok {
					log.Printf("the provider %q the resource %q has no category", provider, rt)
					continue
				}
				r.Category = c
				r.Keywords = categoryAndTypeToKeywords(provider, c, rt)
			}
			if icons != nil {
				if ic, ok := icons[rt]; ok {
					r.Icon = ic
				} else {
					icons[rt] = ""
				}
			}
			resources = append(resources, r)

			if nt, ok := aliases[t][rt]; ok {
				for _, t := range nt {
					r.Type = t
					resources = append(resources, r)
				}
			}
		}

		buff := &bytes.Buffer{}

		td := TemplateData{
			Resources: resources,
			Type:      t,
			Provider:  provider,
		}

		err = resourceTmpl.Execute(buff, td)
		if err != nil {
			panic(err)
		}

		b, err := format.Source(buff.Bytes())
		if err != nil {
			buff = &bytes.Buffer{}
			log.Printf("ERROR: could not import for provider %q and type %q setting empty", provider, t)
			td.Resources = nil
			err = resourceTmpl.Execute(buff, td)
			if err != nil {
				panic(err)
			}
			b, err = format.Source(buff.Bytes())
			if err != nil {
				panic(err)
			}
		}

		io.Copy(out, bytes.NewBuffer(b))

	}
	// Update the list of icons with the potential new resources
	if icons != nil {
		f, err := os.OpenFile(filepath.Join(assetsPath, provider, iconFile), os.O_RDWR|os.O_TRUNC, 0755)
		if err != nil {
			log.Fatal(err)
		}
		e := json.NewEncoder(f)
		e.SetEscapeHTML(false)
		e.SetIndent("", " ")
		err = e.Encode(icons)
		if err != nil {
			log.Fatal(err)
		}
		f.Close()
	}
}

// parseATag parses the 'a' tag into a resource and sets it where it should go
func parseATag(s *goquery.Selection, provider, cat string, isResource bool, fileType, categories map[string]string) {
	text := s.Text()
	// Set the href of the link with the actual type of it
	href, _ := s.Attr("href")
	paths := strings.Split(href, "/")
	fileName := strings.Split(paths[len(paths)-1], ".")[0]

	// We want to ignore the 'Data Sources'
	// as they all are the same category
	if !isResource {
		fileType[fmt.Sprintf("d_%s", fileName)] = text
		return
	}

	fileType[fmt.Sprintf("r_%s", fileName)] = text
	if v, ok := categories[text]; ok {
		log.Printf("WARNING: the provider %q with resource %q found with category %q and also on %q", provider, text, v, cat)
		cat += " " + v
	}
	categories[text] = cat
}

//var categoryReplacer = strings.NewReplacer("(", " ", ")", " ", "/", " ")
var categoryReplacer = regexp.MustCompile(`\W`)

func categoryAndTypeToKeywords(provider, c, rt string) []string {
	cws := strings.Split(
		strings.ToLower(categoryReplacer.ReplaceAllString(c, " ")), " ")
	tws := strings.Split(rt, "_")

	res := make([]string, 0, len(cws))
	ws := make(map[string]struct{})
	for _, w := range append(cws, tws...) {
		// The word it's already there
		if _, ok := ws[w]; ok {
			continue
		}
		if w == "resources" || w == provider || w == "" {
			continue
		}
		res = append(res, w)
		ws[w] = struct{}{}
	}

	return res
}

var resourceTypeRe = regexp.MustCompile(`(?ms)page_title:\s"[\w\d\-_]+:\s([\w\d\-_]+)"`)

func getResourceType(b []byte) string {
	res := resourceTypeRe.FindSubmatch(b)
	if len(res) != 2 {
		return ""
	}
	return standardizeSpaces(string(res[1]))
}

var categoryRe = regexp.MustCompile(`(?ms)subcategory:\s"([^"]+)"`)

func getCategory(b []byte) string {
	res := categoryRe.FindSubmatch(b)
	if len(res) != 2 {
		return ""
	}
	return standardizeSpaces(string(res[1]))
}

var shortDescriptionRe = regexp.MustCompile(`(?ms)description:\s(?:\|-)?([^#]+)+---`)

func getShortDescription(b []byte) string {
	res := shortDescriptionRe.FindSubmatch(b)
	if len(res) != 2 {
		return ""
	}
	return escapeRawStringQuotes(standardizeSpaces(string(res[1])))
}

var descriptionRe = regexp.MustCompile(`(?ms)#\s(?:Resource|Data\sSource):\s[\w\d-_]+(.+?)##[^#]+`)

func getDescription(b []byte) string {
	res := descriptionRe.FindSubmatch(b)
	if len(res) != 2 {
		return ""
	}
	return escapeRawStringQuotes(string(res[1]))
}

var argumentsBodyRe = regexp.MustCompile(`(?ms)##\sArgument\sReference(.+)(?:##[^#]+|$)+?`)

var argumentRe = regexp.MustCompile("(?ms)\\*\\s`([^`]+)`\\s-\\s([^*]+)")

func getArguments(b []byte) []resource.Attribute {
	res := argumentsBodyRe.FindSubmatch(b)
	if len(res) != 2 {
		return nil
	}
	atts := argumentRe.FindAllSubmatch(res[1], -1)
	attributes := make([]resource.Attribute, 0, len(atts))
	for _, at := range atts {
		attributes = append(attributes, resource.Attribute{
			Name:        removeQuotes(standardizeSpaces(string(at[1]))),
			Description: escapeRawStringQuotes(standardizeSpaces(string(at[2]))),
		})
	}
	return attributes
}

var attributesRe = regexp.MustCompile(`(?ms)##\sAttributes?\sReference(.+)(?:##[^#]+|$)+?`)

func getAttributes(b []byte) []resource.Attribute {
	res := attributesRe.FindSubmatch(b)
	if len(res) != 2 {
		return nil
	}
	atts := argumentRe.FindAllSubmatch(res[1], -1)
	attributes := make([]resource.Attribute, 0, len(atts))
	for _, at := range atts {
		attributes = append(attributes, resource.Attribute{
			Name:        removeQuotes(standardizeSpaces(string(at[1]))),
			Description: escapeRawStringQuotes(standardizeSpaces(string(at[2]))),
		})
	}
	return attributes
}

func standardizeSpaces(s string) string {
	return strings.Join(strings.Fields(s), " ")
}

func removeQuotes(s string) string {
	return strings.Replace(s, `"`, "", -1)
}

func escapeRawStringQuotes(s string) string {
	return strings.Replace(s, "`", "`+\"`\"+`", -1)
}
