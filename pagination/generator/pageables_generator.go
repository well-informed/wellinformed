// +build ignore

// This program generates pageables.go. It can be invoked by running "go generate"
package main

import (
	"io/ioutil"
	"os"
	"text/template"
	"time"

	log "github.com/sirupsen/logrus"
	yaml "gopkg.in/yaml.v2"
)

type Pageables struct {
	Types []string `yaml:,flow`
}

func main() {
	path, err := os.Getwd()
	die(err)

	data, err := ioutil.ReadFile(path + "/generator/pageables.yml")
	die(err)

	pageables := Pageables{}
	die(yaml.Unmarshal(data, &pageables))

	f, err := os.Create(path + "/pageables_gen.go")
	die(err)
	defer f.Close()

	packageTemplate.Execute(f, struct {
		Timestamp time.Time
		Types     []string
	}{
		Timestamp: time.Now(),
		Types:     pageables.Types,
	})
}

func die(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

var packageTemplate = template.Must(template.New("").Parse(`// Code generated by go generate; DO NOT EDIT.
// This file was generated from "pagination/generator/pageables_generator.go" at
// {{ .Timestamp }}
// using types listed in "pagination/generator/pageables.yml"

package pagination
{{print ""}}
import (
	"github.com/well-informed/wellinformed/graph/model"
)
{{print ""}}
{{- range .Types }}
func {{.}}sToNodes(list []*model.{{.}}) []*model.Node {
	nodes := make([]*model.Node, 0)
	for _, item := range list {
		nodes = append(nodes, &model.Node{
			Value: item,
			ID:    item.ID,
		})
	}
	return nodes
}
{{print ""}}
{{- end }}
`))
