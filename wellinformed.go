package wellinformed

import "github.com/well-informed/wellinformed/graph/model"

type Persistor interface {
	InsertSrcRSSFeed(model.SrcRSSFeed) error
}
