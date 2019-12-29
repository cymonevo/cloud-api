package provider

import (
	"sync"

	"github.com/cymonevo/cloud-api/handler"
)

var (
	articleHandler     handler.BaseHandler
	syncArticleHandler sync.Once
)

func GetArticleHandler() handler.BaseHandler {
	if articleHandler == nil {
		syncArticleHandler.Do(func() {
			//articleHandler = handler.NewArticleHandler(GetRouter(), GetArticleFactory())
			articleHandler = handler.NewArticleHandler(GetRouter(), nil)
		})
	}
	return articleHandler
}
