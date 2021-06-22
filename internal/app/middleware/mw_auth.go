package middleware

import (
	"go-gin/internal/app/config"
	"go-gin/internal/app/errors"
	"go-gin/internal/app/ginplus"
	"go-gin/pkg/auth"

	"github.com/gin-gonic/gin"
)

// UserAuthMiddleware 用户授权中间件
func UserAuthMiddleware(a auth.Auther, skipper ...SkipperFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		var userID string
		if t := ginplus.GetToken(c); t != "" {
			id, err := a.ParseUserID(t)
			if err != nil {
				if err == auth.ErrInvalidToken {
					ginplus.ResError(c, errors.ErrNoPerm)
					return
				}
				ginplus.ResError(c, errors.WithStack(err))
				return
			}
			userID = id
		}

		if userID != "" {
			c.Set(ginplus.UserIDKey, userID)
		}

		if len(skipper) > 0 && skipper[0](c) {
			c.Next()
			return
		}

		if userID == "" {
			if config.GetGlobalConfig().RunMode == "debug" {
				c.Set(ginplus.UserIDKey, config.GetGlobalConfig().Root.UserName)
				c.Next()
				return
			}
			ginplus.ResError(c, errors.ErrNoPerm)
		}
	}
}
