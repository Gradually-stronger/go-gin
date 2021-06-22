package middleware

import (
	"go-gin/internal/app/ginplus"

	"github.com/gin-gonic/gin"
)

// MenuMiddleware 菜单中间件
func MenuMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(ginplus.MenuIDKey, c.GetHeader("MenuID"))
	}
}
