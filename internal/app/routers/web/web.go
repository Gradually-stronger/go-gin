package web

import (
	"go-gin/internal/app/routers/web/ctl"
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

// RegisterRouter 注册/web路由
func RegisterRouter(app *gin.Engine, container *dig.Container) error {
	err := ctl.Inject(container)
	if err != nil {
		return err
	}

	return container.Invoke(func(
		cDemo *ctl.Demo,
	) error {
		g := app.Group("/web")

		v1 := g.Group("/v1")
		{
			// 注册/web/v1/demos
			v1.GET("/demos", cDemo.Query)
		}
		return nil
	})
}
