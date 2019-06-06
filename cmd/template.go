package main

import (
	"text/template"

	"github.com/cycloidio/tfdocs/resource"
)

const (
	rTmpl = `
		package aws

		import (
			"github.com/cycloidio/tfdocs/resource"
		)

		var (
			{{ if eq .Type "r" }}
				Resources = []Resource{
			{{ else }}
				DataSources = []Resource{
			{{ end }}
				{{ range .Resources }}
					resource.Resource{
						Name: "{{ .Name }}",
						Type: "aws_{{ .Type }}",
						Category: "{{ .Category }}",
						ShortDescription: "{{ .ShortDescription }}",
						Description: ` + "`{{ .Description }}`" + `,
						Arguments: []resource.Argument{
							{{ range .Arguments }}
								resource.Attribute{
									Name: "{{ .Name }}",
									Description: ` + "`{{ .Description }}`" + `,
								},
							{{ end }}
						},
						Attributes: []resource.Argument{
							{{ range .Attributes }}
								resource.Attribute{
									Name: "{{ .Name }}",
									Description: ` + "`{{ .Description }}`" + `,
								},
							{{ end }}
						},
					},
				{{ end }}
			}
		)
	`
)

var (
	resourceTmpl *template.Template
)

type TemplateData struct {
	Resources []resource.Resource
	Type      string
}

func init() {
	var err error

	resourceTmpl, err = template.New("test").Parse(rTmpl)
	if err != nil {
		panic(err)
	}
}
