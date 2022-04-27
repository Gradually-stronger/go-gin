package web

import (
	"github.com/casbin/casbin"
	"github.com/gin-gonic/gin"
	"go-gin/internal/app/middleware"
	"go-gin/internal/app/routers/web/ctl"
	"go-gin/pkg/auth"
	"go.uber.org/dig"
)

// RegisterRouter 注册/web路由
func RegisterRouter(app *gin.Engine, container *dig.Container) error {
	err := ctl.Inject(container)
	if err != nil {
		return err
	}

	return container.Invoke(func(
		a auth.Auther,
		cDemo *ctl.Demo,
		cLogin *ctl.Login,
		e *casbin.Enforcer,
		cUser *ctl.User,
	) error {
		g := app.Group("/web")

		// 用户身份授权
		g.Use(middleware.UserAuthMiddleware(
			a,
			middleware.AllowMethodAndPathPrefixSkipper(
				middleware.JoinRouter("GET", "/web/v1/pub/login"),
				middleware.JoinRouter("POST", "/web/v1/pub/login"),
			),
		))
		// casbin权限校验中间件
		g.Use(middleware.CasbinMiddleware(e,
			middleware.AllowMethodAndPathPrefixSkipper(
				middleware.JoinRouter("GET", "/web/v1/pub"),
				middleware.JoinRouter("POST", "/web/v1/pub"),
			),
		))
		// 请求频率限制中间件
		g.Use(middleware.RateLimiterMiddleware())

		v1 := g.Group("/v1")
		{
			// 注册/web/v1/demos
			v1.POST("/demos", cDemo.Create)
			pub := v1.Group("/pub")
			{
				// 注册/web/v1/pub/login
				pub.GET("/login/captchaid", cLogin.GetCaptcha)
				pub.GET("/login/captcha", cLogin.ResCaptcha)
				pub.POST("/login", cLogin.Login)

			}
			v1.POST("/users", cUser.Create)
			v1.GET("/users", cUser.Get)
		}
		return nil
	})
}
