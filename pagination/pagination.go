package pagination

import (
	b64 "encoding/base64"
	"errors"

	"github.com/well-informed/wellinformed/graph/model"
)

//go:generate go run generator/pageables_generator.go

func BuildPage(first int, after *string, nodes []*model.Node) (*model.Connection, error) {
	edges := nodesToEdges(nodes)
	if after != nil {
		for i := 0; i < len(edges); i++ {
			if *after == edges[i].Cursor {
				if i+1 == len(edges) {
					return nil, errors.New("cursor not found in list")
				} else if i+first+1 > len(edges) {
					edges = edges[i+1:]
				} else {
					edges = edges[i+1 : i+first+1]
				}
				break
			}
		}
	} else if first < len(edges) {
		edges = edges[:first]
	}
	info := &model.PageInfo{
		HasPreviousPage: len(edges) > 0 && after != nil,
		HasNextPage:     len(edges) > first,
		StartCursor:     edges[0].Cursor,
		EndCursor:       edges[len(edges)-1].Cursor,
	}
	return &model.Connection{
		Edges:    edges,
		PageInfo: info,
	}, nil
}

func nodesToEdges(nodes []*model.Node) []*model.Edge {
	edges := make([]*model.Edge, 0)
	for _, node := range nodes {
		edges = append(edges, &model.Edge{
			Node:   node,
			Cursor: b64.StdEncoding.EncodeToString([]byte(string(node.ID))),
		})
	}
	return edges
}
