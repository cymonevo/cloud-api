package provider

import (
	"sync"

	"github.com/cymonevo/cloud-api/internal/render"
	"github.com/cymonevo/cloud-api/internal/router"
)

var (
	appRouter     router.Router
	syncAppRouter sync.Once

	renderEngine     render.Client
	syncRenderEngine sync.Once
)

func GetRouter() router.Router {
	if appRouter == nil {
		syncAppRouter.Do(func() {
			appRouter = router.New(GetRenderEngine())
		})
	}
	return appRouter
}

func GetRenderEngine() render.Client {
	if renderEngine == nil {
		syncRenderEngine.Do(func() {
			renderEngine = render.New(render.Config{
				TemplatePath: "files/",
			})
		})
	}
	return renderEngine
}
