package pagination

import (
	"testing"

	log "github.com/sirupsen/logrus"
	"github.com/well-informed/wellinformed/graph/model"
)

func TestContentItemsPaging(t *testing.T) {
	contentItems := []*model.ContentItem{{ID: 1}, {ID: 2}}

	list := ContentItems(contentItems)
	pageInfo, err := BuildPage(1, nil, list)
	if err != nil {
		t.Error("could not build page. err: ", err)
	}
	var edges []*model.ContentItemsEdge
	for i, v := range list {
		edge := &model.ContentItemsEdge{
			ContentItem: v,
			Cursor:      list.GetID(i),
		}
		edges = append(edges, edge)
	}
	//Need to figure out how to get a []*model.ContentItem back out. Think I could pass one in and have buildPage modify it in place?
	connection := &model.ContentItemsConnection{
		Edges:    edges,
		PageInfo: pageInfo,
	}
	log.Info("firstPage: PageInfo: %+v, Edges: %+v", connection.Edges, connection.PageInfo)

	//Build second page
	_, err := BuildPage(1, &connection.PageInfo.EndCursor, list)
	if err != nil {
		t.Error("could not build second page. err: ", err)
	}

	// nextConnection := &model.ContentItemsConnection{
	// 	Edges: nextPage,
	// }

}
