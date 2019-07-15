package main

import (
	"bytes"
	"fmt"
	"go/format"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/cycloidio/tfdocs/resource"
)

func main() {
	docsPath := filepath.Join("providers", "aws", "website", "docs")
	sidebarPath := filepath.Join("providers", "aws", "website", "aws.erb")

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
		panic(err)
	}

	doc, err := goquery.NewDocumentFromReader(bytes.NewBuffer(body))
	if err != nil {
		panic(err)
	}

	doc.Find(".nav.docs-sidenav > li").Each(func(_ int, s *goquery.Selection) {
		var cat string
		s.Find("a").Each(func(i int, s *goquery.Selection) {
			text := s.Text()

			if i == 0 {
				cat = text
				return
			}

			// Set the href of the link with the actual type of it
			href, _ := s.Attr("href")
			paths := strings.Split(href, "/")
			fileName := strings.Split(paths[len(paths)-1], ".")[0]

			// We want to ignore the 'Data Sources'
			// as they all are the same category
			if cat == "Data Sources" {
				fileType[fmt.Sprintf("d_%s", fileName)] = text
				return
			}

			fileType[fmt.Sprintf("r_%s", fileName)] = text
			if v, ok := categories[text]; ok {
				log.Fatalf("the resource %q found with category %q and also on %q", text, v, cat)
			}
			categories[text] = cat
		})
	})

	//spew.Dump(fileType)
	//panic("")

	for _, t := range []string{"r", "d"} {
		files, err := ioutil.ReadDir(filepath.Join(docsPath, t))
		if err != nil {
			panic(err)
		}

		out, err := os.OpenFile(filepath.Join("aws", fmt.Sprintf("%s.go", t)), os.O_APPEND|os.O_TRUNC|os.O_RDWR|os.O_CREATE, 0755)
		if err != nil {
			panic(err)
		}
		defer out.Close()

		resources := make([]resource.Resource, 0)
		for _, f := range files {
			body, err := ioutil.ReadFile(filepath.Join(docsPath, t, f.Name()))
			if err != nil {
				panic(err)
			}

			fileName := fmt.Sprintf("%s_%s", t, strings.Split(f.Name(), ".")[0])
			rt, ok := fileType[fileName]
			if !ok {
				log.Printf("the file %q of %q has no type, name used was %q", f.Name(), t, fileName)
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
					log.Printf("the resource %q has no category", rt)
					continue
				}
				r.Category = c
				r.Keywords = categoryToKeywords(c)
			}
			resources = append(resources, r)
		}

		buff := &bytes.Buffer{}

		td := TemplateData{
			Resources: resources,
			Type:      t,
		}

		err = resourceTmpl.Execute(buff, td)
		if err != nil {
			panic(err)
		}

		b, err := format.Source(buff.Bytes())
		if err != nil {
			panic(err)
		}

		io.Copy(out, bytes.NewBuffer(b))
	}
}

var categoryReplacer = strings.NewReplacer("(", "", ")", "", "/", "")

func categoryToKeywords(c string) []string {
	ws := strings.Split(strings.ToLower(c), " ")

	res := make([]string, 0, len(ws))
	for _, w := range ws {
		if w == "resources" {
			continue
		}
		w = categoryReplacer.Replace(w)
		res = append(res, w)
	}

	return res
}

//var resourceTypeRe = regexp.MustCompile(`(?ms)#\s(?:Resource|Data\sSource):\s([\w\d-_]+)\n`)

var resourceTypeRe = regexp.MustCompile(`(?ms)page_title:\s"[\w\d\-_]+:\s([\w\d\-_]+)"`)

func getResourceType(b []byte) string {
	res := resourceTypeRe.FindSubmatch(b)
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
	return standardizeSpaces(string(res[1]))
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
			Name:        standardizeSpaces(string(at[1])),
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
			Name:        standardizeSpaces(string(at[1])),
			Description: escapeRawStringQuotes(standardizeSpaces(string(at[2]))),
		})
	}
	return attributes
}

func standardizeSpaces(s string) string {
	return strings.Join(strings.Fields(s), " ")
}

func escapeRawStringQuotes(s string) string {
	return strings.Replace(s, "`", "`+\"`\"+`", -1)
}
