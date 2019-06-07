package main

import (
	"bytes"
	"fmt"
	"go/format"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/cycloidio/tfdocs/resource"
)

func main() {
	docsPath := filepath.Join("providers", "aws", "website", "docs")

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
			resources = append(resources, resource.Resource{
				Type:             trimResourceType(f.Name()),
				ShortDescription: getShortDescription(body),
				Description:      getDescription(body),
				Arguments:        getArguments(body),
				Attributes:       getAttributes(body),
			})
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

func trimResourceType(n string) string {
	return strings.Split(n, ".")[0]
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
