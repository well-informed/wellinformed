package database

import (
	b64 "encoding/base64"
	"errors"

	"github.com/well-informed/wellinformed/graph/model"
)

func buildPage(first int, after *string, nodes []*model.Pageable) (*model.Connection, error) {
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

func nodesToEdges(nodes []*model.Pageable) []*model.Edge {
	edges := make([]*model.Edge, 0)
	for _, node := range nodes {
		value := *node
		edges = append(edges, &model.Edge{
			Node:   value,
			Cursor: b64.StdEncoding.EncodeToString([]byte(string(value.GetID()))),
		})
	}
	return edges
}
