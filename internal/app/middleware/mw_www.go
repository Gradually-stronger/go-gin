package middleware

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

// WWWMiddleware 静态站点中间件
func WWWMiddleware(root string, skipper ...SkipperFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		if len(skipper) > 0 && skipper[0](c) {
			c.Next()
			return
		}

		p := c.Request.URL.Path

		fpath := filepath.Join(root, filepath.FromSlash(p))
		_, err := os.Stat(fpath)
		if err != nil && os.IsNotExist(err) {
			fpath = filepath.Join(root, "index.html")
			// 只禁止index.html文件的缓存
			resHeader := c.Writer.Header()
			resHeader.Set("Cache-Control", "private, no-store, no-cache, must-revalidate, proxy-revalidate")
			http.ServeFile(c.Writer, c.Request, fpath)
		} else {
			if p == "index.html" || p == "/" {
				resHeader := c.Writer.Header()
				resHeader.Set("Cache-Control", "private, no-store, no-cache, must-revalidate, proxy-revalidate")
			}
			http.ServeFile(c.Writer, c.Request, fpath)
		}
		c.Abort()
	}
}
