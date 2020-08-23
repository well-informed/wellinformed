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
	b64 "encoding/base64"
	"errors"

	"github.com/well-informed/wellinformed/graph/model"
)
{{print ""}}
{{- range .Types }}
func Build{{.}}Page(first int, after *string, list []*model.{{.}}) (*model.{{.}}Connection, error) {
	edges := __{{.}}__listToEdges(list)
	if after != nil {
		for i := 0; i < len(edges); i++ {
			if *after == edges[i].Cursor {
				nextIdx := i + 1
				if nextIdx == len(edges) {
					return nil, errors.New("cursor not found in list")
				} else if nextIdx+first > len(edges) {
					edges = edges[nextIdx:]
				} else {
					edges = edges[nextIdx : nextIdx+first]
				}
				break
			}
		}
	} else if first < len(edges) {
		edges = edges[:first]
	}
	info := &model.{{.}}PageInfo{
		HasPreviousPage: len(edges) > 0 && after != nil,
		HasNextPage:     len(edges) > first,
		StartCursor:     edges[0].Cursor,
		EndCursor:       edges[len(edges)-1].Cursor,
	}
	return &model.{{.}}Connection{
		Edges:    edges,
		PageInfo: info,
	}, nil
}
{{print ""}}
func __{{.}}__listToEdges(list []*model.{{.}}) []*model.{{.}}Edge {
	edges := make([]*model.{{.}}Edge, 0)
	for _, item := range list {
		edges = append(edges, &model.{{.}}Edge{
			Node:   item,
			Cursor: b64.StdEncoding.EncodeToString([]byte(string(item.ID))),
		})
	}
	return edges
}
{{print ""}}
{{- end }}
`))
