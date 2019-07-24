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
				Resources = []*Resource{
			{{ else }}
				DataSources = []*Resource{
			{{ end }}
				{{- range .Resources }}
					&resource.Resource{
						Name: "{{ .Name }}",
						Type: "{{ .Type }}",
						Category: "{{ .Category }}",
						ShortDescription: ` + "`{{ .ShortDescription }}`" + `,
						Description: ` + "`{{ .Description }}`" + `,
						Keywords: []string{
							{{- range .Keywords }}
							  "{{ . }}",
							{{- end }}
						},
						Arguments: []resource.Argument{
							{{- range .Arguments }}
								resource.Attribute{
									Name: "{{ .Name }}",
									Description: ` + "`{{ .Description }}`" + `,
								},
							{{- end }}
						},
						Attributes: []resource.Argument{
							{{- range .Attributes }}
								resource.Attribute{
									Name: "{{ .Name }}",
									Description: ` + "`{{ .Description }}`" + `,
								},
							{{- end }}
						},
					},
				{{- end }}
			}
			{{ if eq .Type "r" }}
				resourcesMap = map[string]int{
			{{ else }}
				dataSourcesMap = map[string]Resource{
			{{ end }}
				{{- range $index, $el := .Resources }}
					"{{ .Type }}": {{ $index }},
				{{- end }}
			}
		)

		{{ if eq .Type "r" }}
			func GetResource(r string) (*resouce.Resource, error) {
				rs, ok := resourcesMap[r]
				if !ok {
					return nil, fmt.Errorf("resource %q not found", r)
				}
				return Resources[rs]
			}
		{{ else }}
			func GetDataSource(r string) (*resouce.Resource, error) {
				rs, ok := dataSourcesMap[r]
				if !ok {
					return nil, fmt.Errorf("datasource %q not found", r)
				}
				return DataSources[rs]
			}
		{{ end }}
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
