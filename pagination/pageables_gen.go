// Code generated by go generate; DO NOT EDIT.
// This file was generated from "pagination/generator/pageables_generator.go" at
// 2020-08-27 10:12:17.670604 -0500 CDT m=+0.001897113
// using types listed in "pagination/generator/pageables.yml"

package pagination

import (
	b64 "encoding/base64"
	"errors"

	"github.com/well-informed/wellinformed/graph/model"
)

func BuildInteractionPage(first int, after *string, list []*model.Interaction) (*model.InteractionConnection, error) {
	edges := __Interaction__listToEdges(list)
	info := &model.InteractionPageInfo{
		HasPreviousPage: len(edges) > 0 && after != nil,
		HasNextPage:     len(edges) > first,
	}
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
	info.StartCursor = edges[0].Cursor
	info.EndCursor = edges[len(edges)-1].Cursor
	return &model.InteractionConnection{
		Edges:    edges,
		PageInfo: info,
	}, nil
}

func __Interaction__listToEdges(list []*model.Interaction) []*model.InteractionEdge {
	edges := make([]*model.InteractionEdge, 0)
	for _, item := range list {
		edges = append(edges, &model.InteractionEdge{
			Node:   item,
			Cursor: b64.StdEncoding.EncodeToString([]byte(string(item.ID))),
		})
	}
	return edges
}

func BuildContentItemPage(first int, after *string, list []*model.ContentItem) (*model.ContentItemConnection, error) {
	edges := __ContentItem__listToEdges(list)
	info := &model.ContentItemPageInfo{
		HasPreviousPage: len(edges) > 0 && after != nil,
		HasNextPage:     len(edges) > first,
	}
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
	info.StartCursor = edges[0].Cursor
	info.EndCursor = edges[len(edges)-1].Cursor
	return &model.ContentItemConnection{
		Edges:    edges,
		PageInfo: info,
	}, nil
}

func __ContentItem__listToEdges(list []*model.ContentItem) []*model.ContentItemEdge {
	edges := make([]*model.ContentItemEdge, 0)
	for _, item := range list {
		edges = append(edges, &model.ContentItemEdge{
			Node:   item,
			Cursor: b64.StdEncoding.EncodeToString([]byte(string(item.ID))),
		})
	}
	return edges
}

func BuildSrcRSSFeedPage(first int, after *string, list []*model.SrcRSSFeed) (*model.SrcRSSFeedConnection, error) {
	edges := __SrcRSSFeed__listToEdges(list)
	info := &model.SrcRSSFeedPageInfo{
		HasPreviousPage: len(edges) > 0 && after != nil,
		HasNextPage:     len(edges) > first,
	}
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
	info.StartCursor = edges[0].Cursor
	info.EndCursor = edges[len(edges)-1].Cursor
	return &model.SrcRSSFeedConnection{
		Edges:    edges,
		PageInfo: info,
	}, nil
}

func __SrcRSSFeed__listToEdges(list []*model.SrcRSSFeed) []*model.SrcRSSFeedEdge {
	edges := make([]*model.SrcRSSFeedEdge, 0)
	for _, item := range list {
		edges = append(edges, &model.SrcRSSFeedEdge{
			Node:   item,
			Cursor: b64.StdEncoding.EncodeToString([]byte(string(item.ID))),
		})
	}
	return edges
}

func BuildUserSubscriptionPage(first int, after *string, list []*model.UserSubscription) (*model.UserSubscriptionConnection, error) {
	edges := __UserSubscription__listToEdges(list)
	info := &model.UserSubscriptionPageInfo{
		HasPreviousPage: len(edges) > 0 && after != nil,
		HasNextPage:     len(edges) > first,
	}
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
	info.StartCursor = edges[0].Cursor
	info.EndCursor = edges[len(edges)-1].Cursor
	return &model.UserSubscriptionConnection{
		Edges:    edges,
		PageInfo: info,
	}, nil
}

func __UserSubscription__listToEdges(list []*model.UserSubscription) []*model.UserSubscriptionEdge {
	edges := make([]*model.UserSubscriptionEdge, 0)
	for _, item := range list {
		edges = append(edges, &model.UserSubscriptionEdge{
			Node:   item,
			Cursor: b64.StdEncoding.EncodeToString([]byte(string(item.ID))),
		})
	}
	return edges
}

