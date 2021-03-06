// +build ignore

package main

import (
	"bytes"
	"go/format"
	"log"
	"os"
	"sort"
	"text/template"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/terraform-providers/terraform-provider-aws/aws"
)

const filename = "resources.go"

type TemplateData struct {
	Resources []string
}

func main() {
	provider := aws.Provider().(*schema.Provider)
	resources := make([]string, 0)

	for name, resource := range provider.ResourcesMap {
		if _, ok := resource.Schema["tags"]; ok {
			resources = append(resources, name)
		}
	}

	sort.Strings(resources)

	tpl, err := template.New("tagged").Parse(templateBody)
	if err != nil {
		log.Fatalf("error parsing template: %v", err)
	}

	var buffer bytes.Buffer
	err = tpl.Execute(&buffer, &TemplateData{
		Resources: resources,
	})
	if err != nil {
		log.Fatalf("error executing template: %v", err)
	}

	formatted, err := format.Source(buffer.Bytes())
	if err != nil {
		log.Fatalf("error formatting generated file: %v", err)
	}

	f, err := os.Create(filename)
	if err != nil {
		log.Fatalf("error creating file (%s): %v", filename, err)
	}
	defer f.Close()

	_, err = f.Write(formatted)
	if err != nil {
		log.Fatalf("error writing to file (%s): %v", filename, err)
	}

}

const templateBody = `
// Code generated by generator/main.go; DO NOT EDIT.

package tags

var Resources = []string{
	{{- range .Resources }}
	"{{ . }}",
	{{- end }}
}
`
