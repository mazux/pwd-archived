package infrastructure

import (
	"fmt"

	"github.com/MAZEN-Kenjrawi/pwd/internal/app/handler"
	"github.com/MAZEN-Kenjrawi/pwd/internal/app/query"
)

type QueryBus struct {
	getLoginHandler    handler.GetLoginHandler
	searchLoginHandler handler.SearchLoginHandler
}

func (b *QueryBus) Handle(q interface{}) (interface{}, error) {
	qType := getType(q)
	switch qType {
	case "GetLogin":
		return b.getLoginHandler.Handle(q.(query.GetLogin))
	case "query.ReplaceLogin":
		return b.searchLoginHandler.Handle(q.(query.SearchLogin))
	}

	return nil, fmt.Errorf("no handler found for query %s", qType)
}
